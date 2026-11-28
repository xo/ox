package ox

import (
	"bytes"
	"context"
	"errors"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestParsePortRange(t *testing.T) {
	tests := []struct {
		s   string
		exp PortRange
		err error
	}{
		{"80", PortRange{80, 80}, nil},
		{"1", PortRange{1, 1}, nil},
		{"65535", PortRange{65535, 65535}, nil},
		{"100-200", PortRange{100, 200}, nil},
		{"100:200", PortRange{100, 200}, nil},
		{"200-200", PortRange{200, 200}, nil},
		{"1-65535", PortRange{1, 65535}, nil},
		{"0080", PortRange{80, 80}, nil},
		{"", PortRange{}, ErrInvalidPortRange},
		{"0", PortRange{}, ErrInvalidPortRange},
		{"0-10", PortRange{}, ErrInvalidPortRange},
		{"10-0", PortRange{}, ErrInvalidPortRange},
		{"65536", PortRange{}, strconv.ErrRange},
		{"1-65536", PortRange{}, strconv.ErrRange},
		{"99999999999999999999", PortRange{}, strconv.ErrRange},
		{"200-100", PortRange{}, ErrInvalidPortRange},
		{"+80", PortRange{}, ErrInvalidPortRange},
		{"-80", PortRange{}, ErrInvalidPortRange},
		{"80-", PortRange{}, ErrInvalidPortRange},
		{":80", PortRange{}, ErrInvalidPortRange},
		{"80:", PortRange{}, ErrInvalidPortRange},
		{"-", PortRange{}, ErrInvalidPortRange},
		{" 80", PortRange{}, ErrInvalidPortRange},
		{"80 ", PortRange{}, ErrInvalidPortRange},
		{"100 - 200", PortRange{}, ErrInvalidPortRange},
		{"100-200-300", PortRange{}, ErrInvalidPortRange},
		{"100:200-300", PortRange{}, ErrInvalidPortRange},
		{"100::200", PortRange{}, ErrInvalidPortRange},
		{"1,2", PortRange{}, ErrInvalidPortRange},
		{"0x50", PortRange{}, ErrInvalidPortRange},
		{"http", PortRange{}, ErrInvalidPortRange},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			r, err := ParsePortRange(test.s)
			switch {
			case test.err != nil && !errors.Is(err, test.err):
				t.Fatalf("expected error %v, got: %v", test.err, err)
			case test.err != nil && !errors.Is(err, ErrInvalidPortRange):
				t.Fatalf("expected error %v, got: %v", ErrInvalidPortRange, err)
			case test.err == nil && err != nil:
				t.Fatalf("expected no error, got: %v", err)
			case r != test.exp:
				t.Errorf("expected %#v, got: %#v", test.exp, r)
			}
		})
	}
}

func TestPortRangeText(t *testing.T) {
	tests := []struct {
		r   PortRange
		exp string
	}{
		{PortRange{}, ""},
		{PortRange{80, 80}, "80"},
		{PortRange{100, 200}, "100-200"},
		{PortRange{1, 65535}, "1-65535"},
		{PortRange{65535, 65535}, "65535"},
	}
	for _, test := range tests {
		t.Run(test.exp, func(t *testing.T) {
			if s := test.r.String(); s != test.exp {
				t.Errorf("expected String %q, got: %q", test.exp, s)
			}
			b, err := test.r.MarshalText()
			if err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if string(b) != test.exp {
				t.Errorf("expected MarshalText %q, got: %q", test.exp, b)
			}
			// the text reads back to the same range
			r := PortRange{1, 2}
			if err := r.UnmarshalText(b); err != nil {
				t.Fatalf("expected no error, got: %v", err)
			}
			if r != test.r {
				t.Errorf("expected %#v, got: %#v", test.r, r)
			}
		})
	}
}

func TestPortRangeUnmarshalTextError(t *testing.T) {
	r := PortRange{1, 2}
	if err := r.UnmarshalText([]byte("100-")); !errors.Is(err, ErrInvalidPortRange) {
		t.Fatalf("expected error %v, got: %v", ErrInvalidPortRange, err)
	}
	if exp := (PortRange{1, 2}); r != exp {
		t.Errorf("expected %#v to stay, got: %#v", exp, r)
	}
}

func TestPortRangeContains(t *testing.T) {
	tests := []struct {
		r    PortRange
		port uint16
		exp  bool
	}{
		{PortRange{80, 80}, 80, true},
		{PortRange{80, 80}, 81, false},
		{PortRange{100, 200}, 99, false},
		{PortRange{100, 200}, 100, true},
		{PortRange{100, 200}, 150, true},
		{PortRange{100, 200}, 200, true},
		{PortRange{100, 200}, 201, false},
		{PortRange{1, 65535}, 65535, true},
		{PortRange{1, 65535}, 0, false},
		{PortRange{}, 0, false},
		{PortRange{0, 10}, 0, false},
		{PortRange{200, 100}, 150, false},
	}
	for _, test := range tests {
		t.Run(test.r.String()+"/"+strconv.Itoa(int(test.port)), func(t *testing.T) {
			if b := test.r.Contains(test.port); b != test.exp {
				t.Errorf("expected %t, got: %t", test.exp, b)
			}
		})
	}
}

func TestPortRangePorts(t *testing.T) {
	tests := []struct {
		name string
		r    PortRange
		exp  []uint16
	}{
		{"one", PortRange{80, 80}, []uint16{80}},
		{"three", PortRange{100, 102}, []uint16{100, 101, 102}},
		{"last", PortRange{65535, 65535}, []uint16{65535}},
		{"end", PortRange{65533, 65535}, []uint16{65533, 65534, 65535}},
		{"zero", PortRange{}, nil},
		{"reversed", PortRange{200, 100}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ports := slices.Collect(test.r.Ports()); !slices.Equal(ports, test.exp) {
				t.Errorf("expected %v, got: %v", test.exp, ports)
			}
		})
	}
	t.Run("all", func(t *testing.T) {
		n := 0
		for port := range (PortRange{1, 65535}).Ports() {
			if n++; int(port) != n {
				t.Fatalf("expected port %d, got: %d", n, port)
			}
		}
		if n != 65535 {
			t.Errorf("expected 65535 ports, got: %d", n)
		}
	})
	t.Run("break", func(t *testing.T) {
		var ports []uint16
		for port := range (PortRange{10, 20}).Ports() {
			if port == 12 {
				break
			}
			ports = append(ports, port)
		}
		if exp := []uint16{10, 11}; !slices.Equal(ports, exp) {
			t.Errorf("expected %v, got: %v", exp, ports)
		}
	})
}

func TestTypeNewPortRange(t *testing.T) {
	testTypeNew(t, []typeTest{
		{
			PortRangeT, []test{
				{"", ""},
				{"80", "80"},
				{"100-200", "100-200"},
				{"100:200", "100-200"},
				{"0", ErrInvalidPortRange},
				{"200-100", ErrInvalidPortRange},
				{"100-", ErrInvalidValue},
				{"65536", ErrInvalidValue},
			},
		},
	})
}

func TestPortRangeTypeLookup(t *testing.T) {
	var v struct {
		R PortRange   `ox:"a range"`
		P *PortRange  `ox:"a pointer"`
		S []PortRange `ox:"a slice"`
	}
	flags, err := FlagsFrom(&v)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	tests := []struct {
		name string
		typ  Type
	}{
		{"value", flags[0].Type},
		{"pointer", flags[1].Type},
		{"slice", flags[2].Elem},
		{"typeRef", typeRef(new(PortRange))},
		{"typeType", typeType[*PortRange]()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.typ != PortRangeT {
				t.Errorf("expected %s, got: %s", PortRangeT, test.typ)
			}
		})
	}
}

func TestPortRangeFlag(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		ports PortRange
		allow []PortRange
		err   error
	}{
		{"default", nil, PortRange{8000, 8099}, nil, nil},
		{"one", []string{"--ports", "80"}, PortRange{80, 80}, nil, nil},
		{"colon", []string{"--ports", "100:200"}, PortRange{100, 200}, nil, nil},
		{"slice", []string{"--allow", "22,80-90", "--allow", "443"}, PortRange{8000, 8099}, []PortRange{{22, 22}, {80, 90}, {443, 443}}, nil},
		{"zero", []string{"--ports", "0"}, PortRange{}, nil, ErrInvalidPortRange},
		{"reversed", []string{"--ports", "200-100"}, PortRange{}, nil, ErrInvalidPortRange},
		{"slice error", []string{"--allow", "22,0"}, PortRange{}, nil, ErrInvalidPortRange},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var ports PortRange
			var got Vars
			err := RunError(
				context.Background(),
				Usage("app", ""),
				Flags().
					PortRange("ports", "the ports", Default("8000-8099"), Bind(&ports)).
					Slice("allow", "the allowed ports", PortRangeT),
				Exec(func(ctx context.Context) {
					c, _ := Ctx(ctx)
					got = c.Vars
				}),
				Args(test.args...),
			)
			switch {
			case test.err != nil && (!errors.Is(err, test.err) || !errors.Is(err, ErrInvalidValue)):
				t.Fatalf("expected errors %v and %v, got: %v", test.err, ErrInvalidValue, err)
			case test.err != nil:
				return
			case err != nil:
				t.Fatalf("expected no error, got: %v", err)
			}
			if r := To[PortRange](got["ports"]); r != test.ports {
				t.Errorf("expected %v, got: %v", test.ports, r)
			}
			if ports != test.ports {
				t.Errorf("expected bound %v, got: %v", test.ports, ports)
			}
			if test.allow == nil {
				return
			}
			if s := ToSlice[PortRange](got["allow"]); !slices.Equal(s, test.allow) {
				t.Errorf("expected %v, got: %v", test.allow, s)
			}
		})
	}
}

func TestPortRangeFrom(t *testing.T) {
	var v struct {
		Ports PortRange            `ox:"the ports,default:8000-8099"`
		Allow []PortRange          `ox:"the allowed ports"`
		Admin *PortRange           `ox:"the admin ports"`
		Peers map[string]PortRange `ox:"the ports of each peer"`
	}
	err := RunError(
		context.Background(),
		Usage("app", ""),
		From(&v),
		Exec(func() {}),
		Args("--allow", "22,80-90", "--admin", "9000:9010", "--peers", "a=1-2,b=3"),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if exp := (PortRange{8000, 8099}); v.Ports != exp {
		t.Errorf("expected %v, got: %v", exp, v.Ports)
	}
	if exp := []PortRange{{22, 22}, {80, 90}}; !slices.Equal(v.Allow, exp) {
		t.Errorf("expected %v, got: %v", exp, v.Allow)
	}
	if exp := (PortRange{9000, 9010}); v.Admin == nil || *v.Admin != exp {
		t.Errorf("expected %v, got: %v", exp, v.Admin)
	}
	if exp := map[string]PortRange{"a": {1, 2}, "b": {3, 3}}; len(v.Peers) != 2 || v.Peers["a"] != exp["a"] || v.Peers["b"] != exp["b"] {
		t.Errorf("expected %v, got: %v", exp, v.Peers)
	}
}

func TestPortRangeHelp(t *testing.T) {
	var buf bytes.Buffer
	err := RunError(
		context.Background(),
		Usage("app", ""),
		Help(),
		Pipe(nil, &buf, &buf),
		Flags().
			PortRange("ports", "the ports", Default("8000-8099")).
			Slice("allow", "the allowed ports", PortRangeT),
		Args("--help"),
	)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for _, exp := range []string{
		"--ports portrange  the ports (default: 8000-8099)",
		"--allow portrange  the allowed ports\n",
	} {
		if !strings.Contains(buf.String(), exp) {
			t.Errorf("expected help to contain %q, got:\n%s", exp, buf.String())
		}
	}
}
