package ox

import (
	"context"
	"errors"
	"net/netip"
	"net/url"
	"testing"
)

func TestTypeNew(t *testing.T) {
	for _, tt := range typeTests(t) {
		for _, test := range tt.tests {
			t.Run(tt.typ.String()+"/"+toString(test.s), func(t *testing.T) {
				expErr, ok := test.exp.(error)
				switch v, err := tt.typ.New(); {
				case err != nil:
					t.Fatalf("expected no error, got: %v", err)
				default:
					switch err := v.Set(context.Background(), toString(test.s)); {
					case err != nil && ok && !errors.Is(err, expErr):
						t.Errorf("expected error %v, got: %v", expErr, err)
					case err != nil && !ok:
						t.Errorf("expected no error, got: %v", err)
					case err == nil && ok:
						t.Errorf("expected error %v, got: nil", expErr)
					case !ok:
						t.Logf("type: %T", v)
						s, err := v.Get()
						if err != nil {
							t.Fatalf("expected no error, got: %v", err)
						}
						t.Logf("val: %s", s)
						if exp := toString(test.exp); s != exp {
							t.Errorf("expected %q, got: %q", exp, s)
						}
					}
				}
			})
		}
	}
}

func typeTests(t *testing.T) []typeTest {
	t.Helper()
	return []typeTest{
		{
			StringT, []test{
				{"", ""},
				{"a", "a"},
				{15, "15"},
				{int64(20), "20"},
				{[]rune("foo"), "foo"},
				{[]byte("bar"), "bar"},
				{float64(0.0), "0"},
			},
		},
		{
			BoolT, []test{
				{"", false},
				{"true", true},
				{"t", true},
				{"T", true},
				{"foo", ErrInvalidValue},
			},
		},
		{
			IntT, []test{
				{"", int(0)},
				{"0", int(0)},
				{0, int(0)},
				{21551, int(21551)},
				{float64(1.0), int(1)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Float64T, []test{
				{"", float64(0.0)},
				{"0.0", float64(0.0)},
				{"79.99", float64(79.99)},
				{float64(57.33), float64(57.33)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			URLT, []test{
				{"", mustURL(t, "")},
				{"https://www.google.com", mustURL(t, "https://www.google.com")},
				{"file:test", mustURL(t, "file:test")},
			},
		},
		{
			AddrT,
			[]test{
				{"", nil},
				{"0.0.0.0", mustAddr(t, "0.0.0.0")},
				{"127.0.0.1", mustAddr(t, "127.0.0.1")},
				{"::ffff:192.168.140.255", mustAddr(t, "::ffff:192.168.140.255")},
				{"foo", ErrInvalidValue},
			},
		},
		{
			AddrPortT,
			[]test{
				{"1.2.3.4:80", mustAddrPort(t, "1.2.3.4:80")},
				{"[::]:80", mustAddrPort(t, "[::]:80")},
				{"[1::CAFE]:80", mustAddrPort(t, "[1::cafe]:80")},
				{"[1::CAFE%en0]:80", mustAddrPort(t, "[1::cafe%en0]:80")},
				{"[::FFFF:192.168.140.255]:80", mustAddrPort(t, "[::ffff:192.168.140.255]:80")},
				{"[::FFFF:192.168.140.255%en0]:80", mustAddrPort(t, "[::ffff:192.168.140.255%en0]:80")},
				{"foo", ErrInvalidValue},
			},
		},
		{
			CIDRT,
			[]test{
				{"1.2.3.4/24", mustPrefix(t, "1.2.3.4/24")},
				{"fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118", mustPrefix(t, "fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118")},
				{"::ffff:c000:0280/96", mustPrefix(t, "::ffff:192.0.2.128/96")},
				{"::ffff:192.168.140.255/8", mustPrefix(t, "::ffff:192.168.140.255/8")},
				{"1.2.3.4/24", mustPrefix(t, "1.2.3.4/24")},
				{"foo", ErrInvalidValue},
			},
		},
	}
}

func mustURL(t *testing.T, s string) *url.URL {
	t.Helper()
	u, err := url.Parse(s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return u
}

func mustAddr(t *testing.T, s string) netip.Addr {
	t.Helper()
	a, err := netip.ParseAddr(s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return a
}

func mustAddrPort(t *testing.T, s string) netip.AddrPort {
	t.Helper()
	a, err := netip.ParseAddrPort(s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return a
}

func mustPrefix(t *testing.T, s string) netip.Prefix {
	t.Helper()
	a, err := netip.ParsePrefix(s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return a
}

type typeTest struct {
	typ   Type
	tests []test
}

type test struct {
	s   any
	exp any
}
