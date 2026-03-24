package ox

import (
	"net/netip"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/xo/ox/strcase"
)

func TestInterpolateVar(t *testing.T) {
	root := &Command{
		Name: "MyApp",
	}
	home, user := os.Getenv("HOME"), os.Getenv("USER")
	pwd, err := DefaultGetwd()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	configDir, err := userConfigDir()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	cacheDir, err := userCacheDir()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	user, pwd, configDir, cacheDir = user, pwd, configDir, cacheDir
	tests := []struct {
		v   string
		exp string
	}{
		{``, ``},
		{`a`, `a`},
		{`\\`, `\`},
		{`\{`, `{`},
		{`\{\}`, `{}`},
		{` \{\} `, ` {} `},
		{` {} `, ` {} `},
		{`A`, `A`},
		{`$`, `$`},
		{`$$`, `$$`},
		{`${}`, `${}`},
		{` $$ `, ` $$ `},
		{`foo`, `foo`},
		{`$ABC`, ``},
		{`$A`, longstr},
		{`$A{}`, longstr},
		{`${A}`, longstr},
		{`${A|upper}`, strings.ToUpper(longstr)},
		{` ${A} `, ` ` + longstr + ` `},
		{` {${A}} `, ` {` + longstr + `} `},
		{` {$A\}} `, ` {` + longstr + `}} `},
		{` {${A|upper}} `, ` {` + strings.ToUpper(longstr) + `} `},
		{` [${A|upper}] `, ` [` + strings.ToUpper(longstr) + `] `},
		{` /${A|upper|lower}/ `, ` /` + longstr + `/ `},
		{` {${A\}} `, ` { `},  // bad substitution
		{` {${A\\}} `, ` { `}, // bad substitution
		{` {${\\}} `, ` { `},  // bad substitution
		{`$PWD`, pwd},
		{`$HOME`, home},
		{`$USER`, user},
		{`$APPNAME`, root.Name},
		{`$CONFIG`, configDir},
		{`$APPCONFIG`, filepath.Join(configDir, root.Name)},
		{`$CACHE`, cacheDir},
		{`$APPCACHE`, filepath.Join(cacheDir, root.Name)},
		{`$NUMCPU`, strconv.Itoa(runtime.NumCPU())},
		{`$NUMCPU2`, strconv.Itoa(runtime.NumCPU() + 2)},
		{`$NUMCPU2X`, strconv.Itoa(runtime.NumCPU() * 2)},
		{`$ARCH`, runtime.GOARCH},
		{`$OS`, runtime.GOOS},
		{`$ENV{HOME}`, home},
		{`$ENV{NOT_DEFINED}`, ``},
		{`$MY_OVERRIDE`, `override`},
		{`$MY_OVERRIDE{foo}`, `bar`},
		{`$MY_OVERRIDE{foo|upper}`, `BAR`},
		{`${MY_OVERRIDE|upper}`, `OVERRIDE`},
		{`${MY_OVERRIDE::e||blah|upper}`, `BLAH`},
		{`$HOME/$USER/$APPNAME`, home + `/` + user + `/` + root.Name},
		{` $HOME $USER $APPNAME `, ` ` + home + ` ` + user + ` ` + root.Name + ` `},
		{`$HOME$USER$APPNAME`, home + user + root.Name},
		{`$HOME$USER${APPNAME}`, home + user + root.Name},
		{`$HOME${USER}${APPNAME}`, home + user + root.Name},
		{`$HOME${USER${APPNAME}`, home + "${USER" + root.Name},
		{`$HOME${USER$APPNAME`, home + "${USER" + root.Name},
		{`//$OS//`, `//` + runtime.GOOS + `//`},
		{`//${OS}//`, `//` + runtime.GOOS + `//`},
		{`${HOME|upper}${USER|upper}${APPNAME|upper}`, strings.ToUpper(home) + strings.ToUpper(user) + strings.ToUpper(root.Name)},
		{`// $HOME/.${APPNAME|lower}rc //`, `// ` + home + `/.` + strings.ToLower(root.Name) + `rc //`},
		{`${UNDEFINED||a}`, `a`},
		{`$ENV{HOME||a}`, home},
		{`${env::HOME||a}`, home},
		{`${env::__undef__||a}`, `a`},
		{`${env::__undef__||a|upper}`, `A`},
		{`${env::__undef__||a||b}`, `a`},
		{`${env::__undef__||a||b|upper}`, `A`},
		{`$$FOO`, `$`},
		{`$$HOME`, `$` + home},
		{`$$HOME_`, `$`},
		{`$$ENV{HOME}`, `$` + home},
		{`$$ENV{HOME}_`, `$` + home + `_`},
		{`$MO_999_`, `COW GOES MOO`},
		{`${MO_999_}`, `COW GOES MOO`},
		{`${_MO_999_}`, ``},
		{`${_MO_999_`, `${_MO_999_`},
		{`${MO_999_|lower}`, `cow goes moo`},
		// filters
		{`${undef||a|sha1}`, `86f7e437faa5a7fce15d1ddcb9eaeaea377667b8`},
		{`${undef||a|sha1|md5}`, `77de54ccf56eb6f7dbf99e4d3be949ab`},
		{`${undef||a|sha1|md5|sha256}`, `03fe963498c652c72bf590b3baa2a327ac8976989a6f024994f923a221ffc931`},
		{`${undef||a|sha1|md5|sha512}`, `1e7ecf724efcf303623dde1b992b9dc107cb9e77cff0acd8d1bd238713be182282e84daa511a6c3b6f083662c5395bfa8793319dd0c6d8fef2999a532a64cb5e`},
		{`${HOME|kebab}`, strcase.ToKebab(home)},
		{`${A|camel}`, strcase.ForceCamelIdentifier(longstr)},
		{`${A|snake}`, strcase.ToSnake(longstr)},
		{`${A|ident}`, strcase.ToIdentifier(longstr)},
		{`${A|kebab}`, strcase.ToKebab(longstr)},
		// bash-like
		{`${HOME:-foo}`, home},
		{`${ABC:-foo}`, `foo`},
		{`${ABC:-||bar}`, `bar`},
		{`${HOME:1}`, home[1:]},
		{`${HOME: 1 }`, home[1:]},
		{`${HOME: 1 : 2 }`, home[1:2]},
		{`${HOME:-1}`, home}, // :- matches the default operator
		{`${HOME: -1}`, home[len(home)-1:]},
		{`${HOME: -2}`, home[len(home)-2:]},
		{`${HOME:1:-2}`, home[1 : len(home)-2]},
		{`${HOME:1:-2|upper}`, strings.ToUpper(home[1 : len(home)-2])},
		{`${HOME|base}`, filepath.Base(home)},
		{`${HOME| base }`, filepath.Base(home)},
		{`${HOME|dir}`, filepath.Dir(home)},
		{`${HOME| dir }`, filepath.Dir(home)},
		{`${VERSION#v}`, `0.1.0`},
		{`${HOME#/home}`, strings.TrimPrefix(home, "/home")},
		{`${HOME%/` + user + `}`, strings.TrimSuffix(home, "/"+user)},
		{`${A/a/|trim}`, `really long long string`},
		{`${A#a really}`, ` long long string`},
		{`${A|upper#A REALLY }`, `LONG LONG STRING`},
		{`${A|upper|lower|kebab%string}`, `a-really-long-long-`},
		{`${A|upper|lower|kebab%-string}`, `a-really-long-long`},
		{`${A/long/foo|upper|lower}`, `a really foo long string`},
		{`${A//long/foo|upper}`, `A REALLY FOO FOO STRING`},
		{`${A/long/\{foo\}}`, `a really {foo} long string`},
		{`${A/long/\{\{foo\}\}}`, `a really {{foo}} long string`},
		{`${a.string}`, `foo`},
		{`$TEST{a.string|upper}`, `FOO`},
		{`${test::a.string|upper}`, `FOO`},
		{`${test::a.string|upper#F}`, `OO`},
		{`${test::a.string|upper#F |lower}`, `FOO`}, // prefix is 'F |lower'
		{`${test::a.string|upper%O}`, `FO`},
		{`${test::a.string|upper%O|lower}`, `FOO`}, // suffix is 'O|lower'
		{`${a.rate}`, `1 MiB/s`},
		{`${a.rate@%d}`, `1048576/s`},
		{`${a.rate@%e}`, `1Mi/s`},
		{`${a.addrport}`, `2.3.4.5:6789`},
	}
	m := make(map[string]bool)
	for i, test := range tests {
		if m[test.v] {
			t.Fatalf("test %d: already included %q", i, test.v)
		}
		m[test.v] = true
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Logf("test: %q exp: %q", test.v, test.exp)
			ctx := &Context{
				Root:   root,
				Lookup: DefaultLookup,
				Loader: DefaultLoader,
				Override: func(typ, key string) (any, bool, error) {
					switch {
					case typ == "MY_OVERRIDE", key == "MY_OVERRIDE":
						switch strings.ToLower(key) {
						case "foo":
							return "bar", true, nil
						case "e":
							return "", true, nil
						}
						return "override", true, nil
					case typ == "MO_999_", key == "MO_999_":
						return "COW GOES MOO", true, nil
					case typ == "A", key == "A":
						return longstr, true, nil
					case key == "VERSION":
						return "v0.1.0", true, nil
					}
					return "", false, nil
				},
			}
			v, err := InterpolateVar(ctx, test.v)
			if err != nil {
				t.Fatalf("expected error no error, got: %v", err)
			}
			t.Logf("v: `%v` (%T)", v, v)
			s, ok := v.(string)
			if !ok {
				t.Fatalf("expected string, got: %T", v)
			}
			if s != test.exp {
				t.Errorf("expected %q, got: %q", test.exp, s)
			}
		})
	}
}

func init() {
	RegisterConfigLoader("test", func(ctx *Context, key string) (any, bool, error) {
		switch key {
		case "a.int":
			return 10, true, nil
		case "a.string":
			return "foo", true, nil
		case "a.rate":
			return NewRate(1*MiB, time.Second), true, nil
		case "a.addrport":
			return netip.MustParseAddrPort("2.3.4.5:6789"), true, nil
		}
		return nil, false, ErrUnknownKey
	})
}

const longstr = "a really long long string"
