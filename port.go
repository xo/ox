package ox

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

// PortRange is a range of TCP or UDP ports, from Lo to Hi, inclusive. A
// valid range has 1 <= Lo <= Hi. The zero PortRange is not valid, and it
// holds no port. See [ParsePortRange] for the text that it reads.
type PortRange struct {
	Lo, Hi uint16
}

// ParsePortRange parses a port, such as 80, or a range of ports, such as
// 100-200 or 100:200. Each port is 1 to 65535, and the first port is not
// greater than the last. ParsePortRange refuses a sign, a space and an empty
// port, such as 100- or :200.
func ParsePortRange(s string) (PortRange, error) {
	lo, hi, ok := strings.Cut(s, "-")
	if !ok {
		lo, hi, ok = strings.Cut(s, ":")
	}
	if !ok {
		hi = lo
	}
	a, err := parsePort(s, lo)
	if err != nil {
		return PortRange{}, err
	}
	b, err := parsePort(s, hi)
	if err != nil {
		return PortRange{}, err
	}
	if b < a {
		return PortRange{}, fmt.Errorf("%w %q: the first port is greater than the last", ErrInvalidPortRange, s)
	}
	return PortRange{Lo: a, Hi: b}, nil
}

// parsePort parses one port of the range s.
func parsePort(s, port string) (uint16, error) {
	if port == "" {
		return 0, fmt.Errorf("%w %q: missing port", ErrInvalidPortRange, s)
	}
	// strconv.ParseUint reads only digits, so it refuses a sign and a space
	for i := range len(port) {
		if port[i] < '0' || '9' < port[i] {
			return 0, fmt.Errorf("%w %q: %q is not a port", ErrInvalidPortRange, s, port)
		}
	}
	n, err := strconv.ParseUint(port, 10, 16)
	switch {
	case err != nil:
		return 0, fmt.Errorf("%w %q: %w", ErrInvalidPortRange, s, err)
	case n == 0:
		return 0, fmt.Errorf("%w %q: port 0 is not valid", ErrInvalidPortRange, s)
	}
	return uint16(n), nil
}

// IsValid returns true when 1 <= Lo <= Hi.
func (r PortRange) IsValid() bool {
	return r.Lo != 0 && r.Lo <= r.Hi
}

// Contains returns true when the range is valid and holds port.
func (r PortRange) Contains(port uint16) bool {
	return r.IsValid() && r.Lo <= port && port <= r.Hi
}

// Ports returns an iterator over each port of the range, in order. A range
// that is not valid has no port.
func (r PortRange) Ports() iter.Seq[uint16] {
	return func(yield func(uint16) bool) {
		if !r.IsValid() {
			return
		}
		// the loop ends at Hi before the increment, so that a range that
		// ends at 65535 does not wrap to 0
		for port := r.Lo; ; port++ {
			if !yield(port) || port == r.Hi {
				return
			}
		}
	}
}

// String satisfies the [fmt.Stringer] interface. It writes 100-200, or 80
// when Lo and Hi are the same. The zero PortRange writes nothing.
func (r PortRange) String() string {
	switch {
	case r == PortRange{}:
		return ""
	case r.Lo == r.Hi:
		return strconv.FormatUint(uint64(r.Lo), 10)
	}
	return strconv.FormatUint(uint64(r.Lo), 10) + "-" + strconv.FormatUint(uint64(r.Hi), 10)
}

// MarshalText satisfies the [encoding.TextMarshaler] interface. It writes
// the text of [PortRange.String].
func (r PortRange) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

// UnmarshalText satisfies the [encoding.TextUnmarshaler] interface. It reads
// the text of [ParsePortRange]. Empty text gives the zero PortRange.
func (r *PortRange) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*r = PortRange{}
		return nil
	}
	v, err := ParsePortRange(string(text))
	if err != nil {
		return err
	}
	*r = v
	return nil
}
