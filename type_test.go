package ox

import (
	"errors"
	"net/netip"
	"net/url"
	"testing"
	"time"
)

func TestTypeNew(t *testing.T) {
	for _, tt := range typeTests(t) {
		for _, test := range tt.tests {
			t.Run(tt.typ.String()+"/"+toString[string](test.v), func(t *testing.T) {
				expErr, ok := test.exp.(error)
				switch v, err := tt.typ.New(); {
				case err != nil:
					t.Fatalf("expected no error, got: %v", err)
				default:
					switch err := v.Set(toString[string](test.v)); {
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
						if exp := toString[string](test.exp); s != exp {
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
			BytesT, []test{
				{"", []byte("")},
				{"a", []byte("a")},
				{15, []byte("15")},
				{int64(20), []byte("20")},
				{[]rune("foo"), []byte("foo")},
				{[]byte("bar"), []byte("bar")},
				{float64(0.0), []byte("0")},
				{float64(0.127), []byte("0.127")},
			},
		},
		{
			StringT, []test{
				{"", ""},
				{"a", "a"},
				{15, "15"},
				{int64(20), "20"},
				{[]rune("foo"), "foo"},
				{[]byte("bar"), "bar"},
				{float64(0.0), "0"},
				{float64(0.127), "0.127"},
			},
		},
		{
			RunesT, []test{
				{"", []rune("")},
				{"a", []rune("a")},
				{15, []rune("15")},
				{int64(20), []rune("20")},
				{[]rune("foo"), []rune("foo")},
				{[]byte("bar"), []rune("bar")},
				{float64(0.0), []rune("0")},
				{float64(0.127), []rune("0.127")},
			},
		},
		{
			Base64T, []test{
				{"", ""},
				{"Zm9v", "foo"},
				{"Zm9vCg==", "foo\n"},
				{"---", ErrInvalidValue},
			},
		},
		{
			HexT, []test{
				{"", ""},
				{"666f6f", "foo"},
				{"666f6f0a", "foo\n"},
				{"---", ErrInvalidValue},
			},
		},
		{
			BoolT, []test{
				{"", false},
				{"true", true},
				{"t", true},
				{"T", true},
				{"f", false},
				{"false", false},
				{"0", false},
				{true, true},
				{false, false},
				{"1", true},
				{float64(1.0), true},
				{uint(128), true},
				{int64(124), true},
				{int64(-124), false},
				{int64(0), false},
				{"foo", ErrInvalidValue},
			},
		},
		{
			ByteT, []test{
				{"", ""},
				{"a", "a"},
				{"😀", ErrInvalidValue},
				{"foo", ErrInvalidValue},
			},
		},
		{
			RuneT, []test{
				{"", ""},
				{"a", "a"},
				{"😀", "😀"},
				{"😀😀", ErrInvalidValue},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Int64T, []test{
				{"", int8(0)},
				{"0", int8(0)},
				{0, int8(0)},
				{21551, "21551"},
				{float64(1.0), int(1)},
				{"57", int8(57)},
				{"-10", int8(-10)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Int8T, []test{
				{"", int8(0)},
				{"0", int8(0)},
				{0, int8(0)},
				{21551, ErrInvalidValue},
				{float64(1.0), int(1)},
				{"57", int8(57)},
				{"-10", int8(-10)},
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
				{"57", int(57)},
				{"-10", int(-10)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Uint64T, []test{
				{"", uint(0)},
				{"0", uint(0)},
				{0, uint(0)},
				{21551, uint(21551)},
				{-25555, ErrInvalidValue},
				{float64(1.0), uint(1)},
				{"foo", ErrInvalidValue},
				{"-10", ErrInvalidValue},
			},
		},
		{
			Uint8T, []test{
				{"", uint(0)},
				{"0", uint(0)},
				{0, uint(0)},
				{21551, ErrInvalidValue},
				{-25555, ErrInvalidValue},
				{float64(1.0), uint(1)},
				{"foo", ErrInvalidValue},
				{"-10", ErrInvalidValue},
			},
		},
		{
			UintT, []test{
				{"", uint(0)},
				{"0", uint(0)},
				{0, uint(0)},
				{21551, uint(21551)},
				{-25555, ErrInvalidValue},
				{float64(1.0), uint(1)},
				{"foo", ErrInvalidValue},
				{"-10", ErrInvalidValue},
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
			Float32T, []test{
				{"", float32(0.0)},
				{"0.0", float32(0.0)},
				{"79.99", float32(79.99)},
				{float64(57.33), float32(57.33)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Complex128T, []test{
				{"", complex128(0.0)},
				{"0.0", complex128(0.0)},
				{"79.99", complex128(79.99)},
				{complex128(57.33), complex128(57.33)},
				{float64(54.33), complex128(54.33)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			Complex64T, []test{
				{"", complex64(0.0)},
				{"0.0", complex64(0.0)},
				{"79.99", complex64(79.99)},
				{complex128(57.33), complex64(57.33)},
				{float64(54.33), complex64(54.33)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			TimestampT, []test{
				{"", ""},
				{"2024-11-24T07:41:36+07:00", mustTime(t, "2024-11-24T07:41:36+07:00", time.RFC3339)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			DateTimeT, []test{
				{"", ""},
				{"2024-11-24 7:41:36", mustTime(t, "2024-11-24 7:41:36", time.DateTime)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			DateT, []test{
				{"", ""},
				{"2004-01-02", mustTime(t, "2004-01-02", time.DateOnly)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			TimeT, []test{
				{"", ""},
				{"7:41:36", mustTime(t, "7:41:36", time.TimeOnly)},
				{"foo", ErrInvalidValue},
			},
		},
		{
			DurationT, []test{
				{"", ""},
				{"0", time.Duration(0)},
				{"1", 1 * time.Second},
				{"1s", 1 * time.Second},
				{int(125), 125 * time.Second},
				{uint(65), "1m5s"},
				{float64(2.0), 2 * time.Second},
				{"foo", ErrInvalidValue},
			},
		},
		{
			CountT, []test{
				{"", 1},
				{"a", 1},
				{"foo", 1},
				{"bar", "1"},
			},
		},
		{
			AddrT,
			[]test{
				{"", ""},
				{"0.0.0.0", mustAddr(t, "0.0.0.0")},
				{"127.0.0.1", mustAddr(t, "127.0.0.1")},
				{"::ffff:192.168.140.255", mustAddr(t, "::ffff:192.168.140.255")},
				{"foo", ErrInvalidValue},
			},
		},
		{
			AddrPortT,
			[]test{
				{"", ""},
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
				{"", ""},
				{"1.2.3.4/24", mustPrefix(t, "1.2.3.4/24")},
				{"fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118", mustPrefix(t, "fd7a:115c:a1e0:ab12:4843:cd96:626b:430b/118")},
				{"::ffff:c000:0280/96", mustPrefix(t, "::ffff:192.0.2.128/96")},
				{"::ffff:192.168.140.255/8", mustPrefix(t, "::ffff:192.168.140.255/8")},
				{"1.2.3.4/24", mustPrefix(t, "1.2.3.4/24")},
				{"foo", ErrInvalidValue},
			},
		},
		{
			URLT, []test{
				{"", mustURL(t, "")},
				{"https://www.google.com", mustURL(t, "https://www.google.com")},
				{"file:test", mustURL(t, "file:test")},
				{":foo", ErrInvalidValue},
			},
		},
	}
}

func mustTime(t *testing.T, s, layout string) FormattedTime {
	t.Helper()
	v, err := time.Parse(layout, s)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	return FormattedTime{layout: layout, v: v}
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
	v   any
	exp any
}
