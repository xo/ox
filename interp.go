package ox

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"github.com/xo/ox/strcase"
)

var (
	// DefaultLookup retrieves keys from the [Context]'s config loaders.
	DefaultLookup = func(ctx *Context, typ, key string) (any, error) {
		if ctx.Override != nil {
			switch v, ok, err := ctx.Override(typ, key); {
			case err != nil:
				return nil, err
			case ok:
				return v, nil
			}
		}
		if ctx.Loader != nil {
			switch v, ok, err := ctx.Loader(ctx, typ, key); {
			case err != nil:
				return nil, err
			case ok:
				return v, nil
			}
		}
		return nil, nil
	}
	// DefaultInterpolate is the default interpolation func.
	DefaultInterpolate = InterpolateVar
	// DefaultOps are default ops used by [InterpolateVar].
	DefaultOps = map[string]func(string, string) (string, error){
		// prefix trim
		"#": func(s, v string) (string, error) {
			return strings.TrimPrefix(s, v), nil
		},
		// suffix trim
		"%": func(s, v string) (string, error) {
			return strings.TrimSuffix(s, v), nil
		},
		// replace one
		"/": func(s, v string) (string, error) {
			str, rep, ok := strings.Cut(v, "/")
			if !ok {
				return "", fmt.Errorf("%w %q: missing /", ErrInvalidOp, v)
			}
			re, err := regexp.Compile(str)
			if err != nil {
				return "", fmt.Errorf("%w invalid regexp %q: %w", ErrInvalidOp, str, err)
			}
			if m := re.FindStringIndex(s); m != nil {
				return string(slices.Replace([]byte(s), m[0], m[1], []byte(rep)...)), nil
			}
			return s, nil
		},
		// replace all
		"//": func(s, v string) (string, error) {
			str, rep, ok := strings.Cut(v, "/")
			if !ok {
				return "", fmt.Errorf("%w %q: missing /", ErrInvalidOp, v)
			}
			re, err := regexp.Compile(str)
			if err != nil {
				return "", fmt.Errorf("%w invalid regexp %q: %w", ErrInvalidOp, str, err)
			}
			return re.ReplaceAllString(s, rep), nil
		},
		// substring
		":": func(s, idx string) (string, error) {
			start, end, single, v := 0, len(s), true, idx
			if i := strings.Index(v, ":"); i != -1 {
				switch i, ok := parseInt(v[i+1:]); {
				case !ok:
					return "", fmt.Errorf("%w index %q", ErrInvalidOp, idx)
				case 0 <= i:
					end = i
				default:
					end = end + i
				}
				single, v = false, v[:i]
			}
			switch i, ok := parseInt(v); {
			case !ok:
				return "", fmt.Errorf("%w index %q", ErrInvalidOp, idx)
			case !single && 0 <= i:
				start = i
			case !single:
				start = len(s) + i
			case single && 0 <= i:
				start = i
			case single:
				start = len(s) + i
			}
			switch { // indices check
			case len(s) < end, end < start, start < 0:
				return "", fmt.Errorf("%w index %q", ErrInvalidOp, idx)
			}
			return s[start:end], nil
		},
		// default
		":-": func(s, v string) (string, error) {
			if s == "" {
				return v, nil
			}
			return s, nil
		},
		// apply filter
		"|": func(s, v string) (string, error) {
			f, ok := DefaultFilters[strings.TrimSpace(v)]
			if !ok {
				return "", fmt.Errorf("%w: %q", ErrInvalidFilter, v)
			}
			return f(s), nil
		},
		// default
		"||": func(s, v string) (string, error) {
			if s == "" {
				return v, nil
			}
			return s, nil
		},
	}
	// DefaultOpFormatFunc is the default op format func used by [InterpolateVar].
	DefaultOpFormatFunc = func(format string, v any) (any, error) {
		return fmt.Sprintf(format, v), nil
	}
	// DefaultFilters are default filters used by [InterpolateVar].
	DefaultFilters = map[string]func(string) string{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"camel": strcase.ForceCamelIdentifier,
		"snake": strcase.ToSnake,
		"ident": strcase.ToIdentifier,
		"kebab": strcase.ToKebab,
		"trim":  strings.TrimSpace,
		"base":  filepath.Base,
		"dir":   filepath.Dir,
		"md5": func(s string) string {
			b := md5.Sum([]byte(s))
			return fmt.Sprintf("%x", string(b[:]))
		},
		"sha1": func(s string) string {
			b := sha1.Sum([]byte(s))
			return fmt.Sprintf("%x", string(b[:]))
		},
		"sha256": func(s string) string {
			b := sha256.Sum256([]byte(s))
			return fmt.Sprintf("%x", string(b[:]))
		},
		"sha512": func(s string) string {
			b := sha512.Sum512([]byte(s))
			return fmt.Sprintf("%x", string(b[:]))
		},
	}
)

// InterpolateVar interpolates variables based on registered config loaders
// defined on the [Context].
//
// Interplates variables of the form:
//
//	$KEY
//	${KEY}
//	${KEY|filter}
//	${KEY||default|filter}
//	$TYPE{KEY}
//	$TYPE{KEY|filter}
//	$TYPE{KEY||default|filter}
//	${TYPE::KEY}
//	${TYPE::KEY|filter}
//	${TYPE::KEY||default|filter}
//	${KEY%str}
//	${KEY%%str}
//	${KEY^str}
//	${KEY^^str}
//	${KEY/str/rep}
//	${KEY//str/rep}
//	${KEY:offset}
//	${KEY:offset:length}
//	${KEY^}
//	${KEY^^}
//	${KEY,}
//	${KEY,,}
//
// Where the above are:
//
//	KEY - the key name
//	TYPE - the config loader type
//	default - a default value
//	filter - a filter defined in [DefaultFilters]
//	str - a string
//	rep - a replacement string
//	offset - a offset into the string; negative values are allowed
//	length - a length from the offset
//
// See [DefaultKeyLoader] for the list of specially recognized keys.
//
// Built-in filters (see [DefaultFilters]):
//
//	lower - lower case
//	upper - upper case
//	camel - camel case
//	snake - snake case
//	ident - ident case
//	kebab - kebab case
//	trim - trim space
//	base - file path base
//	dir - file path dir
//	md5 - md5 hash
//	sha1 - sha1 hash
//	sha256 - sha256 hash
//	sha512 - sha512 hash
//
// Built-in Bash-like string variable operators (see [DefaultOps]):
//
//	#, ##, %, %%, /, //, ^, ^^, :
//
// Interpolation examples:
//
//	$APPNAME
//	$ENV{MY_VALUE|lower||default_value}
//	${env::MY_VALUE|lower||default_value}
//	$YAML{a.key.path|upper}
//	${yaml::a.key.path||default_value|upper}
//	${VERSION#v}
//	${MY_VALUE//foo/bar||default_value|upper}
//	${env::MY_VALUE%foo|upper}
func InterpolateVar(ctx *Context, v any) (any, error) {
	str, ok := v.(string)
	if !ok || str == "" {
		return v, nil
	}
	var sb strings.Builder
	r := []rune(str)
	for i, n := 0, len(r); i < n; i++ {
		switch {
		case r[i] == '\\':
			if c := peek(r, i+1, n); c != 0 {
				_, _ = sb.WriteRune(c)
			}
			i++
			continue
		case r[i] != '$':
			_, _ = sb.WriteRune(r[i])
			continue
		}
		// read var
		typ, key, transform, end := readVar(r, i, n)
		// fmt.Fprintf(os.Stderr, "type: %q key: %q, transform: %t end: %d\n", typ, key, transform, end)
		if end == -1 || (typ == "" && key == "") {
			_ = sb.WriteByte('$')
			continue
		}
		// fmt.Fprintf(os.Stderr, "captured: %q type: %q key: %q trans: %t\n", string(r[i:end]), typ, key, transform)
		switch s, err := lookup(ctx, typ, key, transform); {
		case err == nil:
			_, _ = sb.WriteString(s)
		default:
			_, _ = sb.WriteString(string(r[i:end]))
		}
		i = end - 1
	}
	return sb.String(), nil
}

// lookup retrieves a type/key from the context, and transforms it if
// applicable.
func lookup(ctx *Context, typ, key string, transform bool) (string, error) {
	var ops string
	if i := strings.IndexFunc(key, isOp); i != -1 && transform {
		key, ops = key[:i], key[i:]
	}
	v, err := ctx.Lookup(ctx, typ, key)
	switch {
	case errors.Is(err, ErrUnknownKey), v == nil:
		v = ""
	case err != nil:
		return "", err
	}
	if transform && len(ops) != 0 {
		return apply(v, ops)
	}
	s, err := asString[string](v)
	if err != nil {
		return "", err
	}
	return s, nil
}

// apply applies the string ops to the string.
func apply(v any, ops string) (string, error) {
	var op, str string
	for ops != "" {
		op, str, ops = readOp(ops)
		// fmt.Fprintf(os.Stderr, "  op: %q str: %q ops: %q\n", op, str, ops)
		switch f, ok := DefaultOps[op]; {
		case op == "@":
			var err error
			if v, err = DefaultOpFormatFunc(str, v); err != nil {
				return "", err
			}
		case !ok && len(op) == 2:
			if f, ok = DefaultOps[op[:1]]; !ok {
				return "", fmt.Errorf("%w: %w %q", ErrNotImplemented, ErrInvalidOp, op)
			}
			op, str = op[:1], op[1:]+str
		case !ok:
			return "", fmt.Errorf("%w: %w %q", ErrNotImplemented, ErrInvalidOp, op)
		default:
			s, err := asString[string](v)
			if err != nil {
				return "", err
			}
			if v, err = f(s, str); err != nil {
				return "", err
			}
		}
	}
	s, err := asString[string](v)
	if err != nil {
		return "", err
	}
	return s, nil
}

// readVar reads a var in r.
func readVar(r []rune, i, n int) (string, string, bool, int) {
	// fmt.Fprintf(os.Stderr, "readvar: %q\n", string(r[i:n]))
	var typ, key string
	c := peek(r, i+1, n)
	switch {
	case 'A' <= c && c <= 'Z':
		typ, i = readType(r, i+1, n)
		c = peek(r, i, n)
	case c == '{':
		i++
	}
	bracket := false
	if c == '{' {
		// open bracket
		if key, i = readBracket(r, i+1, n); i == -1 {
			return "", "", false, -1
		}
		bracket = true
	}
	// if empty type, cut type at :: if within {} -- ie, ${typ::key}
	if typ == "" && key != "" && bracket {
		if s, after, ok := strings.Cut(key, "::"); ok {
			typ, key = s, after
		}
	}
	switch {
	case typ == "" && key == "":
		return "", "", false, n
	case key == "":
		typ, key, bracket = "", typ, false
	}
	return typ, key, bracket, min(i, n)
}

// readType reads a variable type between '$' and '{'.
func readType(r []rune, i, n int) (string, int) {
	start := i
	for ; i < n; i++ {
		if r[i] != '_' && !unicode.IsUpper(r[i]) && !unicode.IsDigit(r[i]) {
			break
		}
	}
	return string(r[start:i]), i
}

// readBracket reads until the first unescaped $ or }.
func readBracket(r []rune, i, n int) (string, int) {
	var v []rune
	var c, prev rune
	for ; i < n; i++ {
		switch c = r[i]; {
		case prev != '\\' && c == '$':
			return "", -1
		case prev != '\\' && c == '}':
			return string(v), i + 1
		case prev == '\\':
			prev = c
			fallthrough
		case c != '\\':
			v = append(v, r[i])
		default:
			prev = c
		}
	}
	return "", -1
}

// readOp reads the next op from ops.
func readOp(ops string) (string, string, string) {
	op, ops := ops[0:1], ops[1:]
	f := rune(op[0])
	if len(ops) != 0 && isOpSecond(rune(ops[0]), f) {
		op += string(ops[0])
		ops = ops[1:]
	}
	if f != '%' && f != '#' && f != '^' && f != '@' {
		if i := opIndex(ops, f); i != -1 {
			return op, ops[:i], ops[i:]
		}
	}
	return op, ops, ""
}

// opIndex returns the index for the next op. For '/' and ':' continues to the
// next op.
func opIndex(ops string, op rune) int {
	if i := strings.IndexFunc(ops, isOp); i != -1 {
		switch c := rune(ops[i]); {
		case
			op == '/' && op == c,
			op == ':' && op == c:
			if j := strings.IndexFunc(ops[i+1:], isOp); j != -1 {
				return i + j + 1
			}
		default:
			return i
		}
	}
	return -1
}

// isOp returns true if r is an allowed op character.
func isOp(r rune) bool {
	switch r {
	case '#', '%', '^', ':', '/', '|', '@':
		return true
	}
	return false
}

// isOpSecond returns true if r is an allowed second op character.
func isOpSecond(r, f rune) bool {
	if (f == '#' || f == '%' || f == '^' || f == '@') && f != r {
		return false
	}
	switch r {
	case '#', '%', '^', ':', '/', '|', '@', '?', '=', '+', '-':
		return true
	}
	return false
}

// parseInt parses an int.
func parseInt(s string) (int, bool) {
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return int(i), err == nil
}
