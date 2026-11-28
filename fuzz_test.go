package ox

import (
	"context"
	"io"
	"math/big"
	"strings"
	"testing"
)

// fuzzOpts returns the options of the command tree of FuzzCommandLine. The
// tree uses most features of the parse: short and long flags, a count, a
// slice, a map, an optional value with a func, allowed values, a local flag,
// plus forms, prefixes, negated bools, a sub command with bash's command
// line, sub commands with aliases, typed and variadic arguments, a stop at
// the first argument, and help, version and completion.
func fuzzOpts() []Option {
	return []Option{
		Usage("app", "an app"),
		Defaults(),
		Abbrev(),
		NegateBools(),
		Flags().
			Bool("verbose", "write more", Short("v")).
			Count("level", "the level", Short("l")).
			String("name", "the name", Short("n"), Default("x")).
			Int("port", "the port", Short("p"), Valid(80, 443, 8080)).
			Slice("tags", "the tags", Split(",")).
			Map("env", "the environment").
			Duration("timeout", "the timeout").
			Size("size", "the size").
			PortRange("ports", "the ports").
			String("local", "a local flag", Local()).
			Hook("errexit", "exit on an error", func(*Context) {}, Short("e"), AllowPlus()).
			Var("set-option", "set an option", HookT, Default(func(*Context, string) error { return nil }), Short("o"), AllowPlus()).
			Var("mode", "the mode", OptionalValue("MODE"), HelpValue("fast", "go fast", func(*Context) error {
				return nil
			})),
		Sub(
			Usage("run", "run a command"),
			StopAtArg(),
			Arg("cmd", "the command"),
			Arg("args", "its args", SliceT, Optional()),
			Flags().
				String("dir", "the folder", Short("C")),
			Exec(func(context.Context, []string) {}),
		),
		Sub(
			Usage("sh", "a shell with bash's command line"),
			PlusWords(),
			LetterClusters(),
			Flags().
				Var("opt", "set an option", HookT, Default(func(*Context, string) error { return nil }), Short("O"), AllowPlus(), OptionalAtEnd()).
				Hook("trace", "trace each command", func(*Context) {}, Short("x"), AllowPlus()),
			Exec(func(context.Context, []string) {}),
		),
		Sub(
			Usage("get", "get a thing"),
			Aliases("g"),
			Arg("n", "a number", IntT),
			Arg("rest", "the rest", Optional()),
			Flags().
				Bool("all", "get each thing", Short("a")),
			Exec(func() {}),
		),
		Exec(func() {}),
	}
}

// FuzzCommandLine parses, checks and runs any command line with the tree of
// fuzzOpts. A command line is the fuzz text, split at each newline. The parse
// can fail, and it must never panic. Help, version and completion write to
// io.Discard, and the exit func does nothing.
func FuzzCommandLine(f *testing.F) {
	for _, s := range []string{
		"",
		"-v",
		"-vvv\n-l\n--name=a",
		"--port\n443\n--tags=a,b\n--tags\nc",
		"--env\na=b\n--env=c=d\n--timeout\n1m30s\n--size=1.5GiB\n--ports=80:90",
		"--no-verbose\n--verb\n--mode\n--mode=fast",
		"run\nls\n-la\n--color",
		"run\n-C\n/tmp\nls",
		"g\n3\nx\n-a",
		"get\nnot-a-number",
		"--help",
		"get\n--help",
		"--version",
		"help\nget",
		"__complete\n--",
		"__complete\nget\n",
		"__complete\n--mode=",
		"completion\nbash",
		"--\n-v\n--",
		"-",
		"--=x",
		"-nvalue",
		"--port=99999999999999999999",
		"--loc=x",
		"+e",
		"-e\n+e\n+o\nerrexit\n-o\nerrexit",
		"+eo\nnounset",
		"+oerrexit",
		"+o",
		"+e=1",
		"+ev",
		"+%Y",
		"+",
		"__complete\n+",
		"__complete\n+o\n",
		"sh\n-Ox\nfoo",
		"sh\n+Ox\nfoo",
		"sh\n+\n-x",
		"sh\n+z",
		"sh\n-O",
		"sh\n+xO",
		"sh\n-x=",
		"__complete\nsh\n-Ox\n",
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > 1<<12 {
			return
		}
		var args []string
		if s != "" {
			args = strings.Split(s, "\n")
		}
		ctx, err := NewContext(append(fuzzOpts(), Args(args...))...)
		if err != nil {
			t.Fatalf("expected no error building the tree, got: %v", err)
		}
		ctx.Stdout, ctx.Stderr, ctx.Exit = io.Discard, io.Discard, func(int) {}
		if err := ctx.Parse(); err != nil {
			return
		}
		if ctx.Exec == nil {
			t.Fatalf("expected a command after a parse of %q", args)
		}
		if err := ctx.Validate(); err != nil {
			return
		}
		_ = ctx.Run(context.Background())
	})
}

// fuzzTypes are the types of FuzzTypeValue.
var fuzzTypes = []Type{
	StringT, BytesT, RunesT, BoolT, ByteT, RuneT,
	Int64T, Int32T, Int16T, Int8T, IntT,
	Uint64T, Uint32T, Uint16T, Uint8T, UintT,
	Float64T, Float32T, Complex128T, Complex64T,
	SizeT, RateT, DurationT,
	TimestampT, DateTimeT, DateT, TimeT,
	BigIntT, BigFloatT, BigRatT,
	AddrT, AddrPortT, CIDRT, RegexpT, URLT, PortRangeT,
	Base64T, HexT,
}

// FuzzTypeValue sets any text on a value of each type. Set can refuse the
// text, and it must never panic. When it takes the text, Get must not fail,
// and a new value must take the text of Get and give the same text, so that
// help and config output read back.
func FuzzTypeValue(f *testing.F) {
	for _, s := range []string{
		"", "0", "-1", "1.5", "1e9", "true", "x", "é", "1 GiB", "1536Mi", "2KiB/s",
		"1h2m", "2026-11-28", "2026-11-28 10:30:00", "10:30:00", "2026-11-28T10:30:00Z",
		"1/3", "127.0.0.1", "[::1]:80", "10.0.0.0/8", "a+b", "https://x/y?z", "80-90",
		"aGVsbG8=", "68656c6c6f", "(1+2i)",
	} {
		for i := range fuzzTypes {
			f.Add(uint8(i), s)
		}
	}
	f.Fuzz(func(t *testing.T, i uint8, s string) {
		typ := fuzzTypes[int(i)%len(fuzzTypes)]
		v, err := typ.New()
		if err != nil {
			t.Fatalf("%s: expected no error, got: %v", typ, err)
		}
		if err := v.Set(s); err != nil {
			return
		}
		// Go writes a big.Float with a huge exponent, such as 8p97070970,
		// in about a minute, in MarshalText of math/big. The parse is fast,
		// so ox has no fault, and the fuzz does not wait for the text.
		if f, ok := v.Val().(*big.Float); ok && f != nil {
			if exp := f.MantExp(nil); exp < -1<<20 || 1<<20 < exp {
				return
			}
		}
		got, err := v.Get()
		if err != nil {
			t.Fatalf("%s: Set(%q) worked, and Get failed: %v", typ, s, err)
		}
		w, err := typ.New()
		if err != nil {
			t.Fatalf("%s: expected no error, got: %v", typ, err)
		}
		if err := w.Set(got); err != nil {
			t.Fatalf("%s: Set(%q) worked, and Set of its text %q failed: %v", typ, s, got, err)
		}
		again, err := w.Get()
		if err != nil {
			t.Fatalf("%s: Get of %q failed: %v", typ, got, err)
		}
		if again != got {
			t.Fatalf("%s: Set(%q) gives %q, which reads back as %q", typ, s, got, again)
		}
	})
}

// FuzzParseSize parses any text as a size. A size that parses must read back
// the same from its exact form, with the verbs q and Q (D40).
func FuzzParseSize(f *testing.F) {
	for _, s := range []string{"0", "1", "1536", "1.5GiB", "6e9", "1E3", "1Ki", "6G", "-2 MiB", "9007199254740993", "8EiB"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		size, err := ParseSize(s)
		if err != nil {
			return
		}
		for _, verb := range []rune{'q', 'Q'} {
			text := FormatSize(int64(size), verb, -2, false)
			back, err := ParseSize(text)
			if err != nil {
				t.Fatalf("ParseSize(%q) = %d, and its %%%c form %q does not parse: %v", s, size, verb, text, err)
			}
			if back != size {
				t.Fatalf("ParseSize(%q) = %d, and its %%%c form %q reads back as %d", s, size, verb, text, back)
			}
		}
	})
}

// FuzzParsePortRange parses any text as a port range. A range that parses
// must be valid, and must read back the same from its text.
func FuzzParsePortRange(f *testing.F) {
	for _, s := range []string{"", "80", "1-65535", "100:200", "0", "65536", "200-100", "1-", "-1", "80,90"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		r, err := ParsePortRange(s)
		if err != nil || s == "" {
			return
		}
		if !r.IsValid() {
			t.Fatalf("ParsePortRange(%q) = %v, which is not valid", s, r)
		}
		back, err := ParsePortRange(r.String())
		if err != nil || back != r {
			t.Fatalf("ParsePortRange(%q) = %v, and its text %q reads back as %v, %v", s, r, r.String(), back, err)
		}
	})
}
