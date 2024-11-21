//go:build !tinygo

package ox

import (
	"os"
	"reflect"
)

func overflowComplex(v reflect.Value, c complex128) bool {
	return v.OverflowComplex(c)
}

/*
func userCacheDir() (string, error) {
	return os.UserCacheDir()
}
*/

func userConfigDir() (string, error) {
	return os.UserConfigDir()
}

/*
func userHomeDir() (string, error) {
	return os.UserHomeDir()
}
*/
