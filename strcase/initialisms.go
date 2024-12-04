package strcase

import (
	"fmt"
	"strings"
	"unicode"
)

// Initialisms is a set of initialisms.
type Initialisms struct {
	known map[string]string
	post  map[string]string
	max   int
}

// New creates a new set of initialisms.
func New(initialisms ...string) (*Initialisms, error) {
	ini := &Initialisms{
		known: make(map[string]string),
		post:  make(map[string]string),
	}
	if err := ini.Add(initialisms...); err != nil {
		return nil, err
	}
	return ini, nil
}

// NewDefaultInitialisms creates a set of known, common initialisms.
func NewDefaultInitialisms() (*Initialisms, error) {
	ini, err := New(CommonInitialisms()...)
	if err != nil {
		return nil, err
	}
	var pairs []string
	for _, s := range CommonPlurals() {
		pairs = append(pairs, s+"S", s+"s")
	}
	if err := ini.Post(pairs...); err != nil {
		return nil, err
	}
	return ini, nil
}

// Add adds a known initialisms.
func (ini *Initialisms) Add(initialisms ...string) error {
	for _, s := range initialisms {
		s = strings.ToUpper(s)
		if len(s) < 2 {
			return fmt.Errorf("invalid initialism %q", s)
		}
		ini.known[s], ini.max = s, max(ini.max, len(s))
	}
	return nil
}

// Post adds a key, value pair to the initialisms and post map.
func (ini *Initialisms) Post(pairs ...string) error {
	if len(pairs)%2 != 0 {
		return fmt.Errorf("invalid pairs length %d", len(pairs))
	}
	for i := 0; i < len(pairs); i += 2 {
		s := strings.ToUpper(pairs[i])
		if s != strings.ToUpper(pairs[i+1]) {
			return fmt.Errorf("invalid pair %q, %q", pairs[i], pairs[i+1])
		}
		ini.known[s] = pairs[i+1]
		ini.post[pairs[i+1]] = pairs[i+1]
		ini.max = max(ini.max, len(s))
	}
	return nil
}

// CamelToSnake converts name from camel case ("AnIdentifier") to snake case
// ("an_identifier").
func (ini *Initialisms) CamelToSnake(name string) string {
	if name == "" {
		return ""
	}
	var s string
	var wasUpper, wasLetter, wasIsm, isUpper, isLetter bool
	for i, r, next := 0, []rune(name), ""; i < len(r); i, s = i+len(next), s+next {
		isUpper, isLetter = unicode.IsUpper(r[i]), unicode.IsLetter(r[i])
		// append _ when last was not upper and not letter
		if (wasLetter && isUpper) || (wasIsm && isLetter) {
			s += "_"
		}
		// determine next to append to r
		if ism := ini.Peek(r[i:]); ism != "" && (!wasUpper || wasIsm) {
			next = ism
		} else {
			next = string(r[i])
		}
		// save for next iteration
		wasIsm, wasUpper, wasLetter = 1 < len(next), isUpper, isLetter
	}
	return strings.ToLower(s)
}

// CamelToSnakeIdentifier converts name from camel case to a snake case
// identifier.
func (ini *Initialisms) CamelToSnakeIdentifier(name string) string {
	return ToIdentifier(ini.CamelToSnake(name))
}

// SnakeToCamel converts name to CamelCase.
func (ini *Initialisms) SnakeToCamel(name string) string {
	var s string
	for _, word := range strings.Split(name, "_") {
		if word == "" {
			continue
		}
		if u, ok := ini.known[strings.ToUpper(word)]; ok {
			s += u
		} else {
			s += strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return s
}

// SnakeToCamelIdentifier converts name to its CamelCase identifier (first
// letter is capitalized).
func (ini *Initialisms) SnakeToCamelIdentifier(name string) string {
	return ini.SnakeToCamel(ToIdentifier(name))
}

// ForceCamelIdentifier forces name to its CamelCase specific to Go
// ("AnIdentifier").
func (ini *Initialisms) ForceCamelIdentifier(name string) string {
	if name == "" {
		return ""
	}
	return ini.SnakeToCamelIdentifier(ini.CamelToSnake(name))
}

// ForceLowerCamelIdentifier forces the first portion of an identifier to be
// lower case ("anIdentifier").
func (ini *Initialisms) ForceLowerCamelIdentifier(name string) string {
	if name == "" {
		return ""
	}
	name = ini.CamelToSnake(name)
	first := strings.Split(name, "_")[0]
	name = ini.SnakeToCamelIdentifier(name)
	return strings.ToLower(first) + name[len(first):]
}

// Peek returns the next longest possible initialism in v.
func (ini *Initialisms) Peek(r []rune) string {
	// do no work
	if len(r) < 2 {
		return ""
	}
	// peek next few runes, up to max length of the largest known initialism
	var i int
	for n := min(len(r), ini.max); i < n && unicode.IsLetter(r[i]); i++ {
	}
	// bail if not enough letters
	if i < 2 {
		return ""
	}
	// determine if common initialism
	var k string
	for i = min(ini.max, i+1, len(r)); i >= 2; i-- {
		k = string(r[:i])
		if s, ok := ini.known[k]; ok {
			return s
		}
		if s, ok := ini.post[k]; ok {
			return s
		}
	}
	return ""
}

// Is indicates whether or not s is a registered initialism.
func (ini *Initialisms) Is(s string) bool {
	_, ok := ini.known[strings.ToUpper(s)]
	return ok
}
