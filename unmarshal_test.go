package ox

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// shade reads its own text, with UnmarshalText on its pointer. Nothing
// registers it.
type shade struct {
	name string
}

func (sh *shade) UnmarshalText(b []byte) error {
	switch s := string(b); s {
	case "light", "dark":
		sh.name = s
		return nil
	}
	return fmt.Errorf("unknown shade %q", b)
}

func (sh shade) MarshalText() ([]byte, error) {
	return []byte(sh.name), nil
}

// loudness is an int that reads a name, as slog.Level does.
type loudness int

func (l *loudness) UnmarshalText(b []byte) error {
	switch string(b) {
	case "quiet":
		*l = 1
	case "loud":
		*l = 9
	default:
		return fmt.Errorf("unknown loudness %q", b)
	}
	return nil
}

// blob reads its value with UnmarshalBinary.
type blob struct {
	b []byte
}

func (bl *blob) UnmarshalBinary(b []byte) error {
	bl.b = append([]byte("blob:"), b...)
	return nil
}

func (bl *blob) MarshalBinary() ([]byte, error) {
	return bl.b, nil
}

// tags is a flag.Value that keeps each value.
type tags []string

func (tg *tags) Set(s string) error {
	*tg = append(*tg, s)
	return nil
}

func (tg *tags) String() string {
	return strings.Join(*tg, "+")
}

// reader has only UnmarshalText.
type reader struct {
	s string
}

func (r *reader) UnmarshalText(b []byte) error {
	r.s = "read:" + string(b)
	return nil
}

// URL has the name of a registered type, so it gets the name of its package.
type URL struct {
	s string
}

func (u *URL) UnmarshalText(b []byte) error {
	u.s = string(b)
	return nil
}

// TestUnmarshalTypes holds that a type that reads its own text works without
// registration, as a field, as the value of a pointer field, with Bind and
// with TypeFor, and that help writes its name in lower case.
func TestUnmarshalTypes(t *testing.T) {
	var args struct {
		Shade    shade             `ox:"a shade"`
		ShadePtr *shade            `ox:"a shade pointer"`
		Shades   []shade           `ox:"shades"`
		ShadeMap map[string]*shade `ox:"a map of shades"`
		Loud     loudness          `ox:"a loudness"`
		Blob     blob              `ox:"a blob"`
		Tags     tags              `ox:"tags"`
		Reader   *reader           `ox:"a reader"`
		URL      URL               `ox:"a url of its own"`
	}
	var bound shade
	var vars Vars
	// each call makes a new flag set, because From adds to the flag set
	opts := func() []Option {
		return []Option{
			Usage("app", "an app"),
			Flags().
				Var("bound", "a bound shade", Bind(&bound)).
				Var("typed", "a typed loudness", TypeFor[loudness]()),
			From(&args),
			Exec(func(ctx context.Context) {
				if c, ok := Ctx(ctx); ok {
					vars = c.Vars
				}
			}),
		}
	}
	err := RunError(context.Background(), append(opts(), Args(
		"--shade=dark", "--shade-ptr=light",
		"--shades=light", "--shades=dark",
		"--shade-map=a=dark",
		"--loud=loud", "--blob=x", "--tags=a", "--tags=b",
		"--reader=r", "--url=u",
		"--bound=light", "--typed=quiet",
	))...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		name string
		v    any
		exp  any
	}{
		{"shade", args.Shade, shade{"dark"}},
		{"shade ptr", *args.ShadePtr, shade{"light"}},
		{"shades", args.Shades, []shade{{"light"}, {"dark"}}},
		{"shade map", *args.ShadeMap["a"], shade{"dark"}},
		{"loud", args.Loud, loudness(9)},
		{"blob", string(args.Blob.b), "blob:x"},
		{"tags", args.Tags, tags{"a", "b"}},
		{"reader", *args.Reader, reader{"read:r"}},
		{"url", args.URL, URL{"u"}},
		{"bound", bound, shade{"light"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.v, test.exp) {
				t.Errorf("expected %#v, got: %#v", test.exp, test.v)
			}
		})
	}
	// the values read with As
	for _, test := range []struct {
		name string
		f    func(Value) (any, error)
		exp  any
	}{
		{"shade", func(v Value) (any, error) { return As[shade](v) }, shade{"dark"}},
		{"shade", func(v Value) (any, error) { return As[*shade](v) }, &shade{"dark"}},
		{"typed", func(v Value) (any, error) { return As[loudness](v) }, loudness(1)},
		{"tags", func(v Value) (any, error) { return As[string](v) }, "a+b"},
	} {
		t.Run("as "+test.name, func(t *testing.T) {
			v, err := test.f(vars[test.name])
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if !reflect.DeepEqual(v, test.exp) {
				t.Errorf("expected %#v, got: %#v", test.exp, v)
			}
		})
	}
	// help writes the name of each type in lower case
	cmd, err := NewCommand(opts()...)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	var buf bytes.Buffer
	help, err := NewCommandHelp(cmd)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	help.AddFlags(&buf)
	for _, exp := range []string{
		"--shade shade ", "--shade-ptr shade ", "--loud loudness ", "--blob blob ",
		"--tags tags ", "--reader reader ", "--url ox.url ", "--bound string ", "--typed loudness ",
	} {
		if !strings.Contains(buf.String(), exp) {
			t.Errorf("expected help to hold %q, got:\n%s", exp, buf.String())
		}
	}
}

// TestUnmarshalErrors holds that a value that the type refuses fails in the
// parse, and that a type that reads no text is refused.
func TestUnmarshalErrors(t *testing.T) {
	var args struct {
		Shade shade `ox:"a shade"`
	}
	err := RunError(context.Background(), Usage("app", "an app"), From(&args), Exec(func() {}), Args("--shade=grey"))
	if !errors.Is(err, ErrInvalidValue) {
		t.Errorf("expected %v, got: %v", ErrInvalidValue, err)
	}
	var plain struct {
		V struct{ s string } `ox:"a struct that reads no text"`
	}
	if _, err := NewCommand(Usage("app", "an app"), From(&plain)); !errors.Is(err, ErrInvalidType) {
		t.Errorf("expected %v, got: %v", ErrInvalidType, err)
	}
	if _, err := NewFlag("v", "v", TypeFor[struct{ s string }]()); !errors.Is(err, ErrInvalidType) {
		t.Errorf("expected %v, got: %v", ErrInvalidType, err)
	}
}

// TestUnmarshalAs holds that As converts a string to a type that reads its
// own text, without registration.
func TestUnmarshalAs(t *testing.T) {
	// mood reads no text
	type mood int
	val := NewVal[string]()
	v, err := val()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if err := v.Set("dark"); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	sh, err := As[shade](v)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if sh.name != "dark" {
		t.Errorf("expected dark, got: %q", sh.name)
	}
	if _, err := As[mood](v); !errors.Is(err, ErrInvalidConversion) {
		t.Errorf("expected %v, got: %v", ErrInvalidConversion, err)
	}
}
