package ox

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

// DefaultLoader is the default config loader.
var DefaultLoader = func(ctx *Context, typ, key string) (any, bool, error) {
	// forced type, only check specific
	if typ != "" {
		if loader, ok := loaders[strings.ToLower(typ)]; ok {
			return loader(ctx, key)
		}
		return nil, false, ErrUnknownLoader
	}
	// check each loader
	for _, typ := range loaderOrder {
		if loader, ok := loaders[typ]; ok {
			switch v, ok, err := loader(ctx, key); {
			case errors.Is(err, ErrUnknownKey):
				continue
			case err != nil:
				return nil, false, err
			case ok:
				return v, true, nil
			}
		}
	}
	return nil, false, ErrUnknownKey
}

// DefaultKeyLoader is the default config key loader.
//
// The following keys are recognized:
//
//	$PWD - the current working directory, taken from [os.Getwd] (ex: /some/path)
//	$HOME - the current user's home directory, taken from [os.UserHomeDir] (ex: ~/)
//	$USER - the current user's user name, taken from [user.Current] (ex: user)
//	$APPNAME - the root command's name (ex: appName)
//	$APPVERSION - the root command's version (ex: v0.1.0)
//	$CONFIG - the current user's config directory, taken from [os.UserConfigDir] (ex: ~/.config)
//	$APPCONFIG - the current user's config directory, with the root command's name added as a subdir (ex: ~/.config/appName)
//	$CACHE - the current user's cache directory, taken from [os.UserCacheDir] (ex: ~/.cache)
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir (ex: ~/.cache/appName)
//	$NUMCPU - the value of [runtime.NumCPU] (ex: 4)
//	$NUMCPU2 - the value of [runtime.NumCPU], plus 2 (ex: 6)
//	$NUMCPU2X - the value of [runtime.NumCPU], times 2 (ex: 8)
//	$ARCH - the value of [runtime.GOARCH] (ex: amd64)
//	$OS - the value of [runtime.GOOS] (ex: windows)
var DefaultKeyLoader = func(ctx *Context, key string) (any, bool, error) {
	var f func() (string, error)
	switch key {
	case "PWD":
		f = DefaultGetwd
	case "HOME":
		f = userHomeDir
	case "USER":
		f = func() (string, error) {
			u, err := user.Current()
			if err != nil {
				return "", err
			}
			return u.Username, nil
		}
	case "APPNAME":
		return ctx.Root.Name, true, nil
	case "APPVERSION":
		return ctx.Root.Version, true, nil
	case "CONFIG":
		f = userConfigDir
	case "APPCONFIG":
		f = func() (string, error) {
			dir, err := userConfigDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "CACHE":
		f = userCacheDir
	case "APPCACHE":
		f = func() (string, error) {
			dir, err := userCacheDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "NUMCPU":
		return strconv.Itoa(runtime.NumCPU()), true, nil
	case "NUMCPU2":
		return strconv.Itoa(runtime.NumCPU() + 2), true, nil
	case "NUMCPU2X":
		return strconv.Itoa(runtime.NumCPU() * 2), true, nil
	case "ARCH":
		return runtime.GOARCH, true, nil
	case "OS":
		return runtime.GOOS, true, nil
	default:
		return nil, false, ErrUnknownKey
	}
	s, err := f()
	if err != nil {
		return "", false, err
	}
	return s, true, nil
}

// DefaultEnvLoader is the default environment config key loader.
var DefaultEnvLoader = func(_ *Context, key string) (any, bool, error) {
	s, ok := os.LookupEnv(key)
	return s, ok, nil
}

var (
	// loaders are config loaders.
	loaders map[string]func(*Context, string) (any, bool, error)
	// loaderOrder is the order that loaders were registered.
	loaderOrder []string
)

func init() {
	loaders = make(map[string]func(*Context, string) (any, bool, error))
	RegisterConfigLoader("", DefaultKeyLoader)
	RegisterConfigLoader("env", DefaultEnvLoader)
}

// RegisterConfigLoader registers a config file type.
func RegisterConfigLoader(typ string, f func(*Context, string) (any, bool, error)) {
	typ = strings.ToLower(typ)
	loaders[typ] = f
	if !slices.Contains(loaderOrder, typ) {
		loaderOrder = append(loaderOrder, typ)
	}
}

// RegisterLoaderOrder sets the loader order.
func RegisterLoaderOrder(order ...string) {
	loaderOrder = loaderOrder[:0]
	for _, typ := range order {
		if typ = strings.ToLower(typ); !slices.Contains(loaderOrder, typ) {
			loaderOrder = append(loaderOrder, typ)
		}
	}
}
