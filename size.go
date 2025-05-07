package ox

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ParseSize parses a byte size string.
func ParseSize(s string) (Size, error) {
	m := sizeRE.FindStringSubmatch(s)
	if m == nil {
		return 0, fmt.Errorf("%w %q", ErrInvalidSize, s)
	}
	f, err := strconv.ParseFloat(m[2], 64)
	switch {
	case err != nil:
		return 0, fmt.Errorf("%w: %w", ErrInvalidSize, err)
	case m[1] == "-":
		f = -f
	}
	sz, err := parseSize(m[3])
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidSize, err)
	}
	return Size(f * float64(sz)), nil
}

// ParseRate parses a byte rate string.
func ParseRate(s string) (Rate, error) {
	unit := time.Second
	if i := strings.LastIndexByte(s, '/'); i != -1 {
		switch strings.ToLower(s[i+1:]) {
		case "ns":
			unit = time.Nanosecond
		case "us", "µs", "μs": // U+00B5 = micro symbol, U+03BC = Greek letter mu
			unit = time.Microsecond
		case "ms":
			unit = time.Millisecond
		case "s":
			unit = time.Second
		case "m":
			unit = time.Minute
		case "h":
			unit = time.Hour
		default:
			return Rate{}, fmt.Errorf("%w %q", ErrInvalidRate, s)
		}
		s = s[:i]
	}
	sz, err := ParseSize(s)
	if err != nil {
		return Rate{}, err
	}
	return Rate{sz, unit}, nil
}

// AppendSize appends the formatted size to b. A precision below -1 will format
// the value to the fixed precision then trims any trailing '.' / '0'. Formats
// the value based on the verb (see below). When space is true, a space will be
// appended between the formatted size and the suffix.
//
// Supported verbs:
//
//	d/D - size in bytes (ex: 12345678)
//	f/F - size in best fitting *B/*iB (ex: 999 B, 1.1 KiB) and always with a space
//	z/Z - size in best fitting *B/*iB
//	k/K - size in KB/KiB (ex: 0.9 kB, 2.3 KiB)
//	m/M - size in MB/MiB (ex: 1.2345 MB)
//	g/G - size in GB/GiB (ex: 1 GiB)
//	t/T - size in TB/TiB (ex: 4.5 TiB)
//	b/P - size in PB/PiB (ex: 4.5 PiB) -- must use b, as p is reserved for pointers
//	s/S - same as f/F
//	v/V - same as f/F
func AppendSize(b []byte, size int64, verb rune, prec int, space bool) []byte {
	neg, sz, unit := "", float64(size), "B"
	switch verb {
	case 'd', 'D':
		return strconv.AppendInt(b, size, 10)
	case 'k':
		sz, unit = sz/KiB, "KiB"
	case 'K':
		sz, unit = sz/KB, "kB"
	case 'm':
		sz, unit = sz/MiB, "MiB"
	case 'M':
		sz, unit = sz/MB, "MB"
	case 'g':
		sz, unit = sz/GiB, "GiB"
	case 'G':
		sz, unit = sz/GB, "GB"
	case 't':
		sz, unit = sz/TiB, "TiB"
	case 'T':
		sz, unit = sz/TB, "TB"
	case 'b':
		sz, unit = sz/PiB, "PiB"
	case 'P':
		sz, unit = sz/PB, "PB"
	case 'z':
		neg, sz, unit = bestSize(size, true)
	case 'Z':
		neg, sz, unit = bestSize(size, false)
	case 'f', 's', 'v':
		neg, sz, unit = bestSize(size, true)
		prec, space = DefaultSizePrec, true
	case 'F', 'S', 'V':
		neg, sz, unit = bestSize(size, false)
		prec, space = DefaultSizePrec, true
	default:
		return append(b, "%"+string(verb)+"(error=unknown size verb)"...)
	}
	aprec := prec
	if aprec < -1 {
		aprec = -aprec
	}
	b = append(b, neg...)
	b = strconv.AppendFloat(b, sz, 'f', aprec, 64)
	// trim right {.,0}
	if prec < -1 {
		b = bytes.TrimRightFunc(b, func(r rune) bool {
			return r == '0'
		})
		b = bytes.TrimRightFunc(b, func(r rune) bool {
			return r == '.'
		})
	}
	if space {
		b = append(b, ' ')
	}
	return append(b, unit...)
}

func bestSize(size int64, iec bool) (string, float64, string) {
	n, units, suffix := int64(KB), "kMGTPE", "B"
	if iec {
		n, units, suffix = KiB, "KMGTPE", "iB"
	}
	var neg string
	if size < 0 {
		neg, size = "-", -size
	}
	if size < n {
		return neg, float64(size), "B"
	}
	e, d := 0, n
	for i := size / n; n <= i; i /= n {
		d *= n
		e++
	}
	return neg, float64(size) / float64(d), string(units[e]) + suffix
}

// FormatSize formats a byte size.
func FormatSize(size int64, verb rune, prec int, space bool) string {
	return string(AppendSize(make([]byte, 0, 28), size, verb, prec, space))
}

// AppendRate appends the formatted rate to b.
func AppendRate(b []byte, rate Rate, verb rune, prec int, space bool) []byte {
	return append(append(AppendSize(b, int64(rate.Size), verb, prec, space), '/'), unitString(rate.Unit)...)
}

// FormatRate formats a byte rate.
func FormatRate(rate Rate, verb rune, prec int, space bool) string {
	return string(AppendRate(make([]byte, 0, 31), rate, verb, prec, space))
}

// Size is a byte size.
type Size int64

// NewSize creates a byte size.
func NewSize[T inti | uinti](size T) Size {
	return Size(size)
}

// Format satisfies the [fmt.Formatter] interface. See [AppendSize] for
// recognized verbs.
func (size Size) Format(f fmt.State, verb rune) {
	prec := DefaultSizePrec
	if p, ok := f.Precision(); ok {
		prec = p
	}
	_, _ = f.Write(AppendSize(make([]byte, 0, 28), int64(size), verb, prec, f.Flag(' ')))
}

// MarshalText satisfies the [BinaryMarshalUnmarshaler] interface.
func (size Size) MarshalText() ([]byte, error) {
	return AppendSize(make([]byte, 0, 28), int64(size), 'z', -2, true), nil
}

// UnmarshalText satisfies the [BinaryMarshalUnmarshaler] interface.
func (size *Size) UnmarshalText(b []byte) error {
	i, err := ParseSize(string(b))
	if err != nil {
		return err
	}
	*size = Size(i)
	return nil
}

// Rate is a byte rate.
type Rate struct {
	Size Size
	Unit time.Duration
}

// NewRate creates a byte rate.
func NewRate[T inti | uinti](size T, unit time.Duration) Rate {
	return Rate{
		Size: NewSize(size),
		Unit: unit,
	}
}

// Format satisfies the [fmt.Formatter] interface. See [AppendRate] for
// recognized verbs.
func (rate Rate) Format(f fmt.State, verb rune) {
	prec := DefaultSizePrec
	if p, ok := f.Precision(); ok {
		prec = p
	}
	_, _ = f.Write(AppendRate(make([]byte, 0, 31), rate, verb, prec, f.Flag(' ')))
}

// Int64 returns the bytes as an int64.
func (rate Rate) Int64() int64 {
	return int64(rate.Size)
}

// UnmarshalText satisfies the [BinaryMarshalUnmarshaler] interface.
func (rate *Rate) UnmarshalText(b []byte) error {
	var err error
	*rate, err = ParseRate(string(b))
	return err
}

// MarshalText satisfies the [BinaryMarshalUnmarshaler] interface.
func (rate *Rate) MarshalText() ([]byte, error) {
	return AppendRate(make([]byte, 0, 31), *rate, 'z', -2, true), nil
}

// Byte sizes.
const (
	B   = 1
	KB  = 1_000
	MB  = 1_000_000
	GB  = 1_000_000_000
	TB  = 1_000_000_000_000
	PB  = 1_000_000_000_000_000
	EB  = 1_000_000_000_000_000_000
	KiB = 1_024
	MiB = 1_048_576
	GiB = 1_073_741_824
	TiB = 1_099_511_627_776
	PiB = 1_125_899_906_842_624
	EiB = 1_152_921_504_606_846_976
)

// parseSize returns the byte size of s.
func parseSize(s string) (int64, error) {
	switch strings.ToLower(s) {
	case "", "b":
		return B, nil
	case "kb":
		return KB, nil
	case "mb":
		return MB, nil
	case "gb":
		return GB, nil
	case "tb":
		return TB, nil
	case "pb":
		return PB, nil
	case "eb":
		return EB, nil
	case "kib":
		return KiB, nil
	case "mib":
		return MiB, nil
	case "gib":
		return GiB, nil
	case "tib":
		return TiB, nil
	case "pib":
		return PiB, nil
	case "eib":
		return EiB, nil
	}
	return 0, fmt.Errorf("%w %q", ErrUnknownSize, s)
}

// unitString returns the string for a time unit (duration).
func unitString(unit time.Duration) string {
	switch {
	case unit == 0, unit > time.Hour:
		return unitString(DefaultRateUnit)
	case unit > time.Minute:
		return "h"
	case unit > time.Second:
		return "m"
	case unit > time.Millisecond:
		return "s"
	case unit > time.Microsecond:
		return "ms"
	case unit > time.Nanosecond:
		return "µs"
	}
	return "ns"
}

// sizeRE matches sizes.
var sizeRE = regexp.MustCompile(`(?i)^([-+])?([0-9]+(?:\.[0-9]*)?)(?: ?([a-z]+))?$`)
