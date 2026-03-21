package ox

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
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
		"##": func(s, v string) (string, error) {
			return "", ErrNotImplemented
		},
		// suffix trim
		"%": func(s, v string) (string, error) {
			return strings.TrimSuffix(s, v), nil
		},
		"%%": func(s, v string) (string, error) {
			return "", ErrNotImplemented
		},
		"^": func(s, _ string) (string, error) {
			return s, nil
		},
		"^^": func(s, _ string) (string, error) {
			return "", ErrNotImplemented
		},
		// replace one
		"/": func(s, v string) (string, error) {
			v, t, ok := strings.Cut(v, "/")
			if !ok {
				return "", fmt.Errorf("missing /")
			}
			re, err := regexp.Compile(v)
			if err != nil {
				return "", err
			}
			r := []rune(s)
			if m := re.FindStringIndex(s); m != nil {
				r = slices.Replace(r, m[0], m[1], []rune(t)...)
			}
			return string(r), nil
		},
		// replace all
		"//": func(s, v string) (string, error) {
			v, t, ok := strings.Cut(v, "/")
			if !ok {
				return "", fmt.Errorf("missing /")
			}
			re, err := regexp.Compile(v)
			if err != nil {
				return "", err
			}
			return re.ReplaceAllString(s, t), nil
		},
		// substring
		":": func(s, v string) (string, error) {
			start, n := 0, len(s)
			end, ok := parseInt(v)
			switch {
			case !ok, n < end:
				return "", fmt.Errorf("invalid index %q", v)
			case end < 0:
				end = len(s) + end
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
		//
		":=": func(s, v string) (string, error) {
			return "", ErrNotImplemented
		},
		":?": func(s, v string) (string, error) {
			return "", ErrNotImplemented
		},
		":+": func(s, v string) (string, error) {
			return "", ErrNotImplemented
		},
		// apply filter
		"|": func(s, v string) (string, error) {
			f, ok := DefaultFilters[v]
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
	// DefaultFilters are default filters used by [InterpolateVar].
	DefaultFilters = map[string]func(string) string{
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
		"camel": strcase.ForceCamelIdentifier,
		"snake": strcase.ToSnake,
		"ident": strcase.ToIdentifier,
		"kebab": strcase.ToKebab,
		"trim":  strings.TrimSpace,
		"md5": func(s string) string {
			b := md5.Sum([]byte(s))
			return string(b[:])
		},
		"sha1": func(s string) string {
			b := sha1.Sum([]byte(s))
			return string(b[:])
		},
		"sha256": func(s string) string {
			b := sha256.Sum256([]byte(s))
			return string(b[:])
		},
		"sha512": func(s string) string {
			b := sha512.Sum512([]byte(s))
			return string(b[:])
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
//	md5 - md5 hash
//	sha1 - sha1 hash
//	sha256 - sha256 hash
//	sha512 - sha512 hash
//
// Built-in Bash style string variable operators (see [DefaultOps]):
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
//	${yaml::a.key:?yaml a.key not defined!}
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
	s, err := asString[string](v)
	if err != nil {
		return "", err
	}
	if transform && len(ops) != 0 {
		return apply(s, ops)
	}
	return s, nil
}

// apply applies the string ops to the string.
func apply(s, ops string) (string, error) {
	var op, str string
	for ops != "" {
		op, str, ops = opNext(ops)
		f, ok := DefaultOps[op]
		if !ok {
			return "", fmt.Errorf("%w %q", ErrInvalidOp, op)
		}
		var err error
		if s, err = f(s, str); err != nil {
			return "", err
		}
	}
	return s, nil
}

// opNext reads the next op from ops.
func opNext(ops string) (string, string, string) {
	op, ops := ops[0:1], ops[1:]
	if len(ops) != 0 && isOp(rune(ops[0])) {
		op += string(ops[0])
		ops = ops[1:]
	}
	if i := strings.IndexFunc(ops, isOp); i != -1 {
		return op, ops[:i], ops[i:]
	}
	return op, ops, ""
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
		case c != '\\', prev == '\\':
			v = append(v, r[i])
		default:
			prev = c
		}
	}
	return "", -1
}

// isOp returns true if r is an allowed op character.
func isOp(r rune) bool {
	switch r {
	case '#', '%', '^', ':', '/', '|', '?', '=', '+':
		return true
	}
	return false
}

// parseInt parses an int.
func parseInt(s string) (int, bool) {
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	return int(i), err == nil
}
