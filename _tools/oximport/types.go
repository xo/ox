package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/xo/ox"
)

// placeholderType returns the type of a value from its placeholder in GNU or
// psql help, such as FILENAME or N. An empty placeholder is a bool.
func placeholderType(spec string) (ox.Type, ox.Type) {
	switch {
	case spec == "":
		return ox.BoolT, ""
	case strings.Contains(spec, "="):
		// NAME=VALUE and VAR[=ARG]
		return ox.MapT, ox.StringT
	}
	switch strings.ToUpper(spec) {
	case "N", "NUM", "NUMBER", "COUNT":
		return ox.IntT, ""
	}
	return ox.StringT, ""
}

// pflagType returns the type, the element type and the placeholder for the
// word that pflag writes after a flag, such as int, strings or stringToString.
// A word that is not a pflag type is a placeholder that the program named in
// backquotes, and the value is a string.
func pflagType(word string) (ox.Type, ox.Type, string) {
	switch word {
	case "":
		return ox.BoolT, "", ""
	case "string":
		return ox.StringT, "", ""
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64":
		return ox.Type(word), "", ""
	case "duration":
		return ox.DurationT, "", ""
	case "count":
		return ox.CountT, "", ""
	case "ip":
		return ox.AddrT, "", ""
	case "ipNet":
		return ox.CIDRT, "", ""
	case "bytesHex":
		return ox.HexT, "", ""
	case "bytesBase64":
		return ox.Base64T, "", ""
	case "strings", "stringSlice":
		return ox.SliceT, ox.StringT, ""
	case "stringArray":
		return ox.ArrayT, ox.StringT, ""
	case "ints", "intSlice":
		return ox.SliceT, ox.IntT, ""
	case "int32Slice", "int64Slice", "uintSlice", "float32Slice", "float64Slice", "durationSlice", "boolSlice":
		return ox.SliceT, ox.Type(strings.TrimSuffix(word, "Slice")), ""
	case "uints":
		return ox.SliceT, ox.UintT, ""
	case "ipSlice":
		return ox.SliceT, ox.AddrT, ""
	case "stringToString":
		return ox.MapT, ox.StringT, ""
	case "stringToInt":
		return ox.MapT, ox.IntT, ""
	case "stringToInt64":
		return ox.MapT, ox.Int64T, ""
	}
	return ox.StringT, "", word
}

// goFlagType returns the type and the placeholder for the word that the flag
// package of Go writes after a flag, such as string or duration.
func goFlagType(word string) (ox.Type, string) {
	switch word {
	case "":
		return ox.BoolT, ""
	case "string":
		return ox.StringT, ""
	case "int", "int64", "uint", "uint64":
		return ox.Type(word), ""
	case "float":
		return ox.Float64T, ""
	case "duration":
		return ox.DurationT, ""
	}
	return ox.StringT, word
}

// kubectlType returns the type and the default for the default that kubectl
// writes after a flag, such as `=false`, `=[]` or `=”`.
func kubectlType(def string) (ox.Type, ox.Type, string) {
	switch {
	case def == "false", def == "true":
		return ox.BoolT, "", def
	case def == "[]":
		return ox.SliceT, ox.StringT, ""
	case strings.HasPrefix(def, "[") && strings.HasSuffix(def, "]"):
		return ox.SliceT, ox.StringT, def[1 : len(def)-1]
	case def == "''", def == `""`:
		return ox.StringT, "", ""
	case numRE.MatchString(def):
		return ox.IntT, "", def
	case isDuration(def):
		return ox.DurationT, "", def
	}
	return ox.StringT, "", unquote(def)
}

// numRE matches an integer.
var numRE = regexp.MustCompile(`^-?[0-9]+$`)

// isDuration returns true when s is a duration with a unit, such as 1m0s.
func isDuration(s string) bool {
	if numRE.MatchString(s) {
		return false
	}
	_, err := time.ParseDuration(s)
	return err == nil
}

// unquote removes the quotes around s, when it has them.
func unquote(s string) string {
	if len(s) >= 2 && (s[0] == '"' && s[len(s)-1] == '"') {
		if v, err := strconv.Unquote(s); err == nil {
			return v
		}
	}
	if len(s) >= 2 && s[0] == '\'' && s[len(s)-1] == '\'' {
		return s[1 : len(s)-1]
	}
	return s
}

// cutDefault removes the default that pflag or the flag package of Go writes
// at the end of a description, such as ` (default "open")`.
func cutDefault(desc string) (string, string) {
	const prefix = " (default "
	i := strings.LastIndex(desc, prefix)
	if i == -1 || !strings.HasSuffix(desc, ")") {
		return desc, ""
	}
	def := desc[i+len(prefix) : len(desc)-1]
	return desc[:i], unquote(def)
}

// cleanDefault returns def ready for [ox.Default]. It escapes each `\` and
// `$`, so that ox does not expand them, it replaces the home directory of the
// program with $HOME, and it removes the brackets around the default of a
// slice.
func cleanDefault(g *Flag, def, home string) string {
	if g.Type == ox.SliceT || g.Type == ox.ArrayT || g.Type == ox.MapT {
		def = strings.TrimSuffix(strings.TrimPrefix(def, "["), "]")
	}
	def = strings.NewReplacer(`\`, `\\`, `$`, `\$`).Replace(def)
	if home != "" && home != "/" {
		if def == home {
			return "$HOME"
		}
		def = strings.ReplaceAll(def, home+"/", "$HOME/")
	}
	return def
}
