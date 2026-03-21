package ox

import (
	"errors"
	"net/netip"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
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
	tests := []struct {
		v   string
		exp string
		err error
	}{
		{``, ``, nil},
		{`a`, `a`, nil},
		{`\\`, `\`, nil},
		{`\{`, `{`, nil},
		{`\{\}`, `{}`, nil},
		{` \{\} `, ` {} `, nil},
		{` {} `, ` {} `, nil},
		{`A`, `A`, nil},
		{`$`, `$`, nil},
		{`$$`, `$$`, nil},
		{`${}`, `${}`, nil},
		{` $$ `, ` $$ `, nil},
		{`foo`, `foo`, nil},
		{`$ABC`, ``, nil},
		{`$A`, `a really long string`, nil},
		{`$A{}`, `a really long string`, nil},
		{`${A}`, `a really long string`, nil},
		{`${A|upper}`, `A REALLY LONG STRING`, nil},
		{` ${A} `, ` a really long string `, nil},
		{` {${A}} `, ` {a really long string} `, nil},
		{` {$A\}} `, ` {a really long string}} `, nil},
		{` {${A|upper}} `, ` {A REALLY LONG STRING} `, nil},
		{` [${A|upper}] `, ` [A REALLY LONG STRING] `, nil},
		{` /${A|upper|lower}/ `, ` /a really long string/ `, nil},
		{` {${A\}} `, ` {${A}} `, nil},
		{` {${A\\}} `, ` {${A\}} `, nil},
		{` {${\\}} `, ` {${\}} `, nil},
		{`$PWD`, pwd, nil},
		{`$HOME`, home, nil},
		{`$USER`, user, nil},
		{`$APPNAME`, root.Name, nil},
		{`$CONFIG`, configDir, nil},
		{`$APPCONFIG`, filepath.Join(configDir, root.Name), nil},
		{`$CACHE`, cacheDir, nil},
		{`$APPCACHE`, filepath.Join(cacheDir, root.Name), nil},
		{`$NUMCPU`, strconv.Itoa(runtime.NumCPU()), nil},
		{`$NUMCPU2`, strconv.Itoa(runtime.NumCPU() + 2), nil},
		{`$NUMCPU2X`, strconv.Itoa(runtime.NumCPU() * 2), nil},
		{`$ARCH`, runtime.GOARCH, nil},
		{`$OS`, runtime.GOOS, nil},
		{`$ENV{HOME}`, home, nil},
		{`$ENV{NOT_DEFINED}`, ``, nil},
		{`$MY_OVERRIDE`, `override`, nil},
		{`$MY_OVERRIDE{foo}`, `bar`, nil},
		{`$MY_OVERRIDE{foo|upper}`, `BAR`, nil},
		{`${MY_OVERRIDE|upper}`, `OVERRIDE`, nil},
		{`${MY_OVERRIDE::e||blah|upper}`, `BLAH`, nil},
		{`$HOME/$USER/$APPNAME`, home + `/` + user + `/` + root.Name, nil},
		{` $HOME $USER $APPNAME `, ` ` + home + ` ` + user + ` ` + root.Name + ` `, nil},
		{`$HOME$USER$APPNAME`, home + user + root.Name, nil},
		{`$HOME$USER${APPNAME}`, home + user + root.Name, nil},
		{`$HOME${USER}${APPNAME}`, home + user + root.Name, nil},
		{`$HOME${USER${APPNAME}`, home + "${USER" + root.Name, nil},
		{`$HOME${USER$APPNAME`, home + "${USER" + root.Name, nil},
		{`//$OS//`, `//` + runtime.GOOS + `//`, nil},
		{`//${OS}//`, `//` + runtime.GOOS + `//`, nil},
		{`${HOME|upper}${USER|upper}${APPNAME|upper}`, strings.ToUpper(home) + strings.ToUpper(user) + strings.ToUpper(root.Name), nil},
		{`// $HOME/.${APPNAME|lower}rc //`, `// ` + home + `/.` + strings.ToLower(root.Name) + `rc //`, nil},
		{`${UNDEFINED||a}`, `a`, nil},
		{`$ENV{HOME||a}`, home, nil},
		{`${env::HOME||a}`, home, nil},
		{`${env::__undef__||a}`, `a`, nil},
		{`${env::__undef__||a|upper}`, `A`, nil},
		{`${env::__undef__||a||b}`, `a`, nil},
		{`${env::__undef__||a||b|upper}`, `A`, nil},
		{`$$FOO`, `$`, nil},
		{`$$HOME`, `$` + home, nil},
		{`$$HOME_`, `$`, nil},
		{`$$ENV{HOME}`, `$` + home, nil},
		{`$$ENV{HOME}_`, `$` + home + `_`, nil},
		{`$MO_999_`, `COW GOES MOO`, nil},
		{`${MO_999_}`, `COW GOES MOO`, nil},
		{`${_MO_999_}`, ``, nil},
		{`${_MO_999_`, `${_MO_999_`, nil},
		{`${MO_999_|lower}`, `cow goes moo`, nil},
		// TODO: bash style interpolation
	}
	for i, test := range tests {
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
						return "a really long string", true, nil
					}
					return "", false, nil
				},
			}
			v, err := InterpolateVar(ctx, test.v)
			switch {
			case err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case err == nil && test.err != nil:
				t.Fatalf("expected error %v, got nil", test.err)
			}
			t.Logf("v: %v (%T)", v, v)
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
