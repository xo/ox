package ox

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// strftimeLayout converts a strftime layout to a Go layout. See [Strftime]
// for the directives that it takes.
func strftimeLayout(layout string) (string, error) {
	if layout == "" {
		return "", fmt.Errorf("%w: empty layout", ErrInvalidLayout)
	}
	// each part is text, or the Go layout of a directive
	type part struct {
		s         string
		directive bool
		frac      bool
	}
	var parts []part
	var sb strings.Builder
	for i := 0; i < len(layout); i++ {
		if layout[i] != '%' {
			j := strings.IndexByte(layout[i:], '%')
			if j == -1 {
				j = len(layout) - i
			}
			parts = append(parts, part{s: layout[i : i+j]})
			_, _ = sb.WriteString(layout[i : i+j])
			i += j - 1
			continue
		}
		if i++; i == len(layout) {
			return "", fmt.Errorf("%w %q: a %% ends the layout", ErrInvalidLayout, layout)
		}
		s, err := strftimeDirective(layout, layout[i:], sb.String())
		if err != nil {
			return "", err
		}
		parts = append(parts, part{
			s:         s,
			directive: layout[i] != '%',
			frac:      layout[i] == 'f' || layout[i] == 'L',
		})
		_, _ = sb.WriteString(s)
	}
	goLayout := sb.String()
	// Go has no escape for text in a layout, so text such as 1 or Jan, or
	// text next to a directive, such as day after %a, can read as a part of
	// the time. Two times that differ in each field write the same with the
	// whole layout and with each part on its own only when Go reads each
	// part as it is.
	for _, t := range []time.Time{
		time.Date(2009, time.November, 10, 23, 4, 5, 123456789, time.FixedZone("AAA", 90*60)),
		time.Date(2017, time.March, 27, 8, 51, 39, 987654321, time.FixedZone("BBB", -(8*60+45)*60)),
	} {
		var exp strings.Builder
		for _, p := range parts {
			switch {
			case !p.directive:
				_, _ = exp.WriteString(p.s)
			case p.frac:
				// a fraction is a fraction only after a . or a ,
				_, _ = exp.WriteString(t.Format("." + p.s)[1:])
			default:
				_, _ = exp.WriteString(t.Format(p.s))
			}
		}
		if t.Format(goLayout) != exp.String() {
			return "", fmt.Errorf("%w %q: Go reads the text between the directives as a part of the time", ErrInvalidLayout, layout)
		}
	}
	return goLayout, nil
}

// strftimeDirective returns the Go layout of the directive at the start of
// s, which follows a %. prev is the Go layout before the directive.
func strftimeDirective(layout, s, prev string) (string, error) {
	switch s[0] {
	case 'Y':
		return "2006", nil
	case 'y':
		return "06", nil
	case 'm':
		return "01", nil
	case 'b':
		return "Jan", nil
	case 'B':
		return "January", nil
	case 'd':
		return "02", nil
	case 'e':
		return "_2", nil
	case 'j':
		return "002", nil
	case 'a':
		return "Mon", nil
	case 'A':
		return "Monday", nil
	case 'H':
		return "15", nil
	case 'I':
		return "03", nil
	case 'p':
		return "PM", nil
	case 'M':
		return "04", nil
	case 'S':
		return "05", nil
	case 'f', 'L':
		if !strings.HasSuffix(prev, ".") && !strings.HasSuffix(prev, ",") {
			return "", fmt.Errorf("%w %q: %%%c must follow a . or a ,", ErrInvalidLayout, layout, s[0])
		}
		if s[0] == 'f' {
			return "000000", nil
		}
		return "000", nil
	case 'z':
		return "-0700", nil
	case 'Z':
		return "MST", nil
	case 'T':
		return "15:04:05", nil
	case 'R':
		return "15:04", nil
	case 'D':
		return "01/02/06", nil
	case 'F':
		return "2006-01-02", nil
	case '%':
		return "%", nil
	}
	r, _ := utf8.DecodeRuneInString(s)
	return "", fmt.Errorf("%w %q: %%%c is not supported", ErrInvalidLayout, layout, r)
}
