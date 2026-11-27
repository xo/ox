// Command cobraimport imports the completion scripts of cobra into comp/. Run
// it from the repository root:
//
//	go run ./_tools/cobraimport [version]
//
// The version is a cobra release, such as v1.10.2. It defaults to latest.
// docs/COBRA-IMPORT.md holds the steps that come before and after it.
//
// It is not in comp/, because ox embeds every file in comp/.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/xo/ox"
)

// scripts maps each cobra source file to the template in comp/ that it
// writes.
var scripts = []struct {
	src, dst string
}{
	{"bash_completionsV2.go", "bash.bash"},
	{"fish_completions.go", "fish.fish"},
	{"powershell_completions.go", "powershell.ps1"},
	{"zsh_completions.go", "zsh.zsh"},
}

// args maps each argument that cobra formats a script with to the verb that
// ox formats the template with. See DefaultCompWrite in ox.go.
var args = map[string]string{
	"name":                   "%[1]s",
	"nameForVar":             "%[2]s",
	"compCmd":                "%[3]s",
	"activeHelpEnvVar(name)": "%[4]s",
}

// edits are the changes that ox makes to the scripts of cobra. Each old text
// must occur exactly once in the converted script, or the import fails. If
// it fails, cobra changed that part of a script. Read the change, and then
// change the edit. docs/COBRA-IMPORT.md holds why each edit exists.
var edits = map[string][]struct {
	old, new string
}{
	"bash.bash": {
		{
			"# bash completion V2 for %-36[1]s -*- shell-script -*-\n",
			"# bash completion for %[1]s\n",
		},
		{
			"        # Only consider the completions that match\n" +
				"        IFS=$'\\n' read -ra COMPREPLY -d '' < <(IFS=$'\\n'; compgen -W \"${completions[*]}\" -- \"${cur}\")\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        IFS=$'\\n' read -ra COMPREPLY -d '' < <(IFS=$'\\n'; compgen -W \"${completions[*]}\" -- \"\")\n",
		},
		{
			"        # Only consider the completions that match what the user typed\n" +
				"        IFS=$'\\n' read -ra COMPREPLY -d '' < <(IFS=$'\\n'; compgen -W \"${completions[*]}\" -- \"${cur}\")\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        IFS=$'\\n' read -ra COMPREPLY -d '' < <(IFS=$'\\n'; compgen -W \"${completions[*]}\" -- \"\")\n",
		},
		{
			"        # Only consider the completions that match\n" +
				"        [[ $comp == \"$cur\"* ]] || continue\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        #[[ $comp == \"$cur\"* ]] || continue\n",
		},
	},
	"fish.fish": {
		{
			"\nfunction __%[2]s_debug\n",
			"# fish completion for %[1]s\n\nfunction __%[2]s_debug\n",
		},
		{
			"        set -l completions (string match -r -- \"^$prefix.*\" $__%[2]s_comp_results)\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        set -l completions (string match -r -- \"^.*\" $__%[2]s_comp_results)\n",
		},
	},
	"powershell.ps1": {
		{
			"# powershell completion for %-36[1]s -*- shell-script -*-\n",
			"# powershell completion for %[1]s\n",
		},
		{
			"        $_.Name -like \"$WordToComplete*\"\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        $_.Name -like \"*\"\n",
		},
	},
	"zsh.zsh": {
		{
			"# zsh completion for %-36[1]s -*- shell-script -*-\n",
			"# zsh completion for %[1]s\n",
		},
		{
			"        if eval _describe $keepOrder \"completions\" completions $flagPrefix $noSpace; then\n",
			"        # ox: offer every completion, so that a corrected name gets through\n" +
				"        if eval _describe $keepOrder \"completions\" completions $flagPrefix $noSpace -U; then\n",
		},
	},
}

// directives maps the name of each cobra directive to the ox directive.
var directives = map[string]ox.CompDirective{
	"ShellCompDirectiveError":         ox.CompError,
	"ShellCompDirectiveNoSpace":       ox.CompNoSpace,
	"ShellCompDirectiveNoFileComp":    ox.CompNoFileComp,
	"ShellCompDirectiveFilterFileExt": ox.CompFilterFileExt,
	"ShellCompDirectiveFilterDirs":    ox.CompFilterDirs,
	"ShellCompDirectiveKeepOrder":     ox.CompKeepOrder,
	"ShellCompDirectiveDefault":       ox.CompDefault,
}

func main() {
	out := flag.String("out", "comp", "directory to write the templates to")
	flag.Parse()
	version := "latest"
	if flag.NArg() > 0 {
		version = flag.Arg(0)
	}
	if err := run(version, *out); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// run downloads the cobra release and writes each template.
func run(version, out string) error {
	dir, version, err := download(version)
	if err != nil {
		return err
	}
	consts, err := stringConsts(dir)
	if err != nil {
		return err
	}
	for _, s := range scripts {
		tpl, err := convert(filepath.Join(dir, s.src), consts)
		if err != nil {
			return fmt.Errorf("converting %s: %w", s.src, err)
		}
		for _, e := range edits[s.dst] {
			if n := strings.Count(tpl, e.old); n != 1 {
				return fmt.Errorf("editing %s: found %d copies of %q, expected 1", s.dst, n, e.old)
			}
			tpl = strings.Replace(tpl, e.old, e.new, 1)
		}
		if err := os.WriteFile(filepath.Join(out, s.dst), []byte(tpl), 0o644); err != nil {
			return err
		}
	}
	fmt.Printf("imported cobra %s\n", version)
	return nil
}

// download downloads the cobra module and returns its directory and its
// version.
func download(version string) (string, string, error) {
	cmd := exec.Command("go", "mod", "download", "-json", "github.com/spf13/cobra@"+version)
	cmd.Stderr = os.Stderr
	buf, err := cmd.Output()
	if err != nil {
		return "", "", fmt.Errorf("downloading cobra %s: %w", version, err)
	}
	var mod struct {
		Dir     string `json:"Dir"`
		Version string `json:"Version"`
	}
	if err := json.Unmarshal(buf, &mod); err != nil {
		return "", "", fmt.Errorf("reading go mod download output: %w", err)
	}
	return mod.Dir, mod.Version, nil
}

// stringConsts returns the value of each string constant in the cobra
// package, such as activeHelpMarker.
func stringConsts(dir string) (map[string]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	consts := make(map[string]string)
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}
		f, err := parser.ParseFile(token.NewFileSet(), filepath.Join(dir, name), nil, 0)
		if err != nil {
			return nil, err
		}
		for _, decl := range f.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok || gen.Tok != token.CONST {
				continue
			}
			for _, spec := range gen.Specs {
				vs := spec.(*ast.ValueSpec)
				for i, v := range vs.Values {
					if s, ok := stringLit(v); ok && i < len(vs.Names) {
						consts[vs.Names[i].Name] = s
					}
				}
			}
		}
	}
	return consts, nil
}

// convert finds the script template in a cobra source file, and rewrites
// each of its verbs for ox.
func convert(path string, consts map[string]string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return "", err
	}
	var format string
	var exprs []string
	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok || format != "" {
			return format == ""
		}
		for i, a := range call.Args {
			// the script is the one format string that is longer than 2000
			// bytes
			if s, ok := stringLit(a); ok && len(s) > 2000 {
				format = s
				for _, x := range call.Args[i+1:] {
					var sb strings.Builder
					_ = printer.Fprint(&sb, fset, x)
					exprs = append(exprs, sb.String())
				}
				return false
			}
		}
		return true
	})
	if format == "" {
		return "", errors.New("no script template found")
	}
	return rewrite(format, exprs, consts)
}

// stringLit returns the value of a string literal, or of string literals
// joined with +.
func stringLit(x ast.Expr) (string, bool) {
	switch x := x.(type) {
	case *ast.BasicLit:
		if x.Kind != token.STRING {
			return "", false
		}
		s, err := strconv.Unquote(x.Value)
		return s, err == nil
	case *ast.BinaryExpr:
		if x.Op != token.ADD {
			return "", false
		}
		a, ok := stringLit(x.X)
		if !ok {
			return "", false
		}
		b, ok := stringLit(x.Y)
		return a + b, ok
	}
	return "", false
}

// rewrite replaces each verb in format with the ox verb or the literal value
// for its argument. It keeps each %% as it is, because the template is a
// format string in ox too.
func rewrite(format string, exprs []string, consts map[string]string) (string, error) {
	var sb strings.Builder
	next := 0
	for i := 0; i < len(format); i++ {
		c := format[i]
		if c != '%' {
			_ = sb.WriteByte(c)
			continue
		}
		if i+1 < len(format) && format[i+1] == '%' {
			_, _ = sb.WriteString("%%")
			i++
			continue
		}
		// read the flags, the width, the argument index and the verb
		j := i + 1
		for j < len(format) && strings.IndexByte("+-# 0123456789.", format[j]) != -1 {
			j++
		}
		arg := next
		if j < len(format) && format[j] == '[' {
			end := strings.IndexByte(format[j:], ']')
			if end == -1 {
				return "", fmt.Errorf("verb at byte %d has no closing ]", i)
			}
			n, err := strconv.Atoi(format[j+1 : j+end])
			if err != nil {
				return "", fmt.Errorf("verb at byte %d: %w", i, err)
			}
			arg, j = n-1, j+end+1
		}
		if j >= len(format) {
			return "", fmt.Errorf("verb at byte %d has no verb character", i)
		}
		verb, flags := format[j], format[i+1:j]
		if k := strings.IndexByte(flags, '['); k != -1 {
			flags = flags[:k]
		}
		if arg < 0 || len(exprs) <= arg {
			return "", fmt.Errorf("verb %q at byte %d names argument %d of %d", format[i:j+1], i, arg+1, len(exprs))
		}
		s, err := replacement(exprs[arg], verb, flags, consts)
		if err != nil {
			return "", fmt.Errorf("verb %q at byte %d: %w", format[i:j+1], i, err)
		}
		_, _ = sb.WriteString(s)
		next, i = arg+1, j
	}
	return sb.String(), nil
}

// replacement returns the text that replaces one verb.
func replacement(expr string, verb byte, flags string, consts map[string]string) (string, error) {
	if v, ok := args[expr]; ok {
		if verb != 's' {
			return "", fmt.Errorf("%s is formatted with %%%c and not %%s", expr, verb)
		}
		// keep a width, such as the one in %-36[1]s
		return "%" + flags + v[1:], nil
	}
	if d, ok := directives[expr]; ok {
		return fmt.Sprintf("%"+flags+string(verb), int(d)), nil
	}
	if s, ok := consts[expr]; ok {
		return strings.ReplaceAll(fmt.Sprintf("%"+flags+string(verb), s), "%", "%%"), nil
	}
	return "", fmt.Errorf("cobra formats the script with %s, and cobraimport has no mapping for it", expr)
}
