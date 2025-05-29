// Package strcase provides methods to convert CamelCase to and from snake_case.
//
// Correctly recognizes common (Go idiomatic) initialisms (HTTP, XML, etc) with
// the ability to define/override/add initialisms.
package strcase

import (
	"regexp"
	"strings"
	"unicode"
)

// Defaults is the set of default (common) initialisms used by the package
// level conversions funcs.
var Defaults *Initialisms

func init() {
	// initialize common default initialisms
	var err error
	if Defaults, err = NewDefaults(); err != nil {
		panic(err)
	}
}

// CamelToSnake converts name from camel case ("AnIdentifier") to snake case
// ("an_identifier").
func CamelToSnake(name string) string {
	return Defaults.CamelToSnake(name)
}

// CamelToSnakeIdentifier converts name from camel case to a snake case
// identifier.
func CamelToSnakeIdentifier(name string) string {
	return Defaults.CamelToSnakeIdentifier(name)
}

// SnakeToCamel converts name to CamelCase.
func SnakeToCamel(name string) string {
	return Defaults.SnakeToCamel(name)
}

// SnakeToCamelIdentifier converts name to its CamelCase identifier (first
// letter is capitalized).
func SnakeToCamelIdentifier(name string) string {
	return Defaults.SnakeToCamelIdentifier(name)
}

// ForceCamelIdentifier forces name to its CamelCase specific to Go
// ("AnIdentifier").
func ForceCamelIdentifier(name string) string {
	return Defaults.ForceCamelIdentifier(name)
}

// ForceLowerCamelIdentifier forces the first portion of an identifier to be
// lower case ("anIdentifier").
func ForceLowerCamelIdentifier(name string) string {
	return Defaults.ForceLowerCamelIdentifier(name)
}

// IsInitialism returns true when s is a registered default initialism.
func IsInitialism(s string) bool {
	return Defaults.Is(s)
}

// ToIdentifier cleans s so that it is usable as an identifier.
//
// Substitutes invalid characters with an underscore, removing leading
// numbers/underscores and trailing underscores.
//
// Additionally collapses multiple underscores to a single underscore.
//
// Makes no changes to case.
func ToIdentifier(s string) string {
	return toIdent(s, '_')
}

// ToSnake cleans s to snake_case.
//
// Substitutes invalid characters with an underscore, removing leading
// numbers/underscores and trailing underscores.
//
// Additionally collapses multiple underscores to a single underscore.
//
// Converts entire string to lower case.
func ToSnake(s string) string {
	return strings.ToLower(toIdent(s, '_'))
}

// ToKebab changes s to kebab-case.
//
// Substitutes invalid characters with a hyphen, removing leading
// numbers/hyphens and trailing hyphens.
//
// Additionally collapses multiple hyphens to a single hyphen.
//
// Converts entire string to lower case.
func ToKebab(s string) string {
	return strings.ToLower(toIdent(s, '-'))
}

// CommonInitialisms returns the set of common initialisms.
//
// Originally built from the list in golang.org/x/lint @ 738671d.
//
// Note: golang.org/x/lint has since been deprecated, and some additional
// initialisms have since been added.
func CommonInitialisms() []string {
	return []string{
		"ACL",
		"API",
		"ASCII",
		"CPU",
		"CSS",
		"CSV",
		"DNS",
		"EOF",
		"GPU",
		"GUID",
		"HTML",
		"HTTP",
		"HTTPS",
		"ID",
		"IP",
		"JSON",
		"LHS",
		"QPS",
		"RAM",
		"RHS",
		"RPC",
		"SLA",
		"SMTP",
		"SQL",
		"SSH",
		"TCP",
		"TLS",
		"TSV",
		"TTL",
		"UDP",
		"UI",
		"UID",
		"URI",
		"URL",
		"UTC",
		"UTF8",
		"UUID",
		"VM",
		"XML",
		"XMPP",
		"XSRF",
		"XSS",
		"YAML",
	}
}

// CommonPlurals returns initialisms that have a common plural of s.
func CommonPlurals() []string {
	return []string{
		"ACL",
		"API",
		"CPU",
		"CSV",
		"GPU",
		"GUID",
		"ID",
		"IP",
		"TSV",
		"UID",
		"UID",
		"URI",
		"URL",
		"UUID",
		"VM",
	}
}

// toIdent converts s to a identifier.
func toIdent(s string, repl rune) string {
	// replace bad chars with c
	s = sub(strings.TrimSpace(s), repl)
	// compact multiple c to single c
	s = regexp.MustCompile(`\Q`+string(repl)+`\E{2,}`).ReplaceAllString(s, string(repl))
	// remove leading numbers and c
	s = strings.TrimLeftFunc(s, func(r rune) bool {
		return unicode.IsNumber(r) || r == repl
	})
	// remove trailing c
	s = strings.TrimRightFunc(s, func(r rune) bool {
		return r == repl
	})
	return s
}

// sub substitutes all non-valid identifier characters in s with repl.
func sub(s string, repl rune) string {
	r := []rune(s)
	for i, ch := range r {
		if !isIdentifierChar(ch, repl) {
			r[i] = repl
		}
	}
	return string(r)
}

// isIdentifierChar determines if ch is a valid character for a Go identifier.
//
// See: go/src/go/scanner/scanner.go.
func isIdentifierChar(ch, repl rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == repl || ch >= 0x80 && unicode.IsLetter(ch) ||
		'0' <= ch && ch <= '9' || ch >= 0x80 && unicode.IsDigit(ch)
}
