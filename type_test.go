package kobra

import (
	"errors"
	"net/netip"
	"net/url"
	"testing"
)

func TestTypeNew(t *testing.T) {
	for _, tt := range typeTests(t) {
		for _, test := range tt.tests {
			t.Run(tt.typ.String()+"/"+toString(test.v), func(t *testing.T) {
				expErr, ok := test.exp.(error)
				switch v, err := tt.typ.New(); {
				case err != nil && ok && !errors.Is(err, expErr):
					t.Errorf("expected error %v, got: %v", expErr, err)
				case err != nil && !ok:
					t.Errorf("expected no error, got: %v", err)
				default:
					t.Logf("type: %T", v)
					s, err := v.Get()
					if err != nil {
						t.Fatalf("expected no error, got: %v", err)
					}
					t.Logf("val: %s", s)
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
				{float64(0.0), "0.0"},
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
				{"https://google.com", mustURL(t, "https://www.google.com")},
				{"file:test", mustURL(t, "file:test")},
			},
		},
		{
			AddrT, []test{
				{"", nil},
				{"0.0.0.0", mustAddr(t, "0.0.0.0")},
				{"127.0.0.1", mustAddr(t, "127.0.0.1")},
				{"::ffff:192.168.140.255", mustAddr(t, "::ffff:192.168.140.255")},
			},
		},
	}
}

func mustURL(t *testing.T, urlstr string) *url.URL {
	t.Helper()
	u, err := url.Parse(urlstr)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return u
}

func mustAddr(t *testing.T, addrstr string) netip.Addr {
	t.Helper()
	a, err := netip.ParseAddr(addrstr)
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
	v   any
	exp any
}
