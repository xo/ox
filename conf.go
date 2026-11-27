package ox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

// ConfigType is a config type.
type ConfigType string

// Config types.
const (
	AnyT  ConfigType = ""
	EnvT  ConfigType = "env"
	JSONT ConfigType = "json"
	YAMLT ConfigType = "yaml"
	HCLT  ConfigType = "hcl"
	TOMLT ConfigType = "toml"
)

// DefaultLoader is the default config loader.
var DefaultLoader = func(ctx *Context, typ ConfigType, key string) (any, bool, error) {
	// forced type, check only one
	if typ != "" {
		if loader, ok := loaders[typ]; ok {
			return loader(ctx, key)
		}
		return nil, false, ErrUnknownLoader
	}
	// check each loader
	for _, typ := range loaderOrder {
		if loader, ok := loaders[typ]; ok {
			switch v, ok, err := loader(ctx, key); {
			case errors.Is(err, ErrUnknownKey), errors.Is(err, ErrInvalidKey):
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
//	$STATE - the current user's state directory, taken from [UserStateDir] (ex: ~/.local/state)
//	$APPCACHE - the current user's cache directory, with the root command's name added as a subdir (ex: ~/.cache/appName)
//	$APPSTATE - the current user's state directory, with the root command's name added as a subdir (ex: ~/.local/state/appName)
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
	case "STATE":
		f = UserStateDir
	case "APPCACHE":
		f = func() (string, error) {
			dir, err := userCacheDir()
			if err != nil {
				return "", err
			}
			return filepath.Join(dir, ctx.Root.Name), nil
		}
	case "APPSTATE":
		f = func() (string, error) {
			dir, err := UserStateDir()
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
	if s, ok := os.LookupEnv(key); ok {
		return s, true, nil
	}
	return "", false, ErrUnknownKey
}

var (
	// loaders are config loaders.
	loaders map[ConfigType]func(*Context, string) (any, bool, error)
	// loaderOrder is the order that loaders were registered.
	loaderOrder []ConfigType
	// formats are the config file formats.
	formats map[ConfigType]configFormat
)

func init() {
	loaders = make(map[ConfigType]func(*Context, string) (any, bool, error))
	formats = make(map[ConfigType]configFormat)
	RegisterConfigLoader(AnyT, DefaultKeyLoader)
	RegisterConfigLoader(EnvT, DefaultEnvLoader)
	RegisterConfigFormat(JSONT, decodeJSON, "json")
}

// RegisterConfigLoader registers a config loader, which answers the keys of
// a config type in interpolation, such as ${env::HOME}. A config file format
// registers its loader with [RegisterConfigFormat].
func RegisterConfigLoader(typ ConfigType, loader func(*Context, string) (any, bool, error)) {
	loaders[typ] = loader
	if !slices.Contains(loaderOrder, typ) {
		loaderOrder = append(loaderOrder, typ)
	}
}

// RegisterConfigLoaderOrder sets the loader order.
func RegisterConfigLoaderOrder(order ...ConfigType) {
	loaderOrder = loaderOrder[:0]
	for _, typ := range order {
		if !slices.Contains(loaderOrder, typ) {
			loaderOrder = append(loaderOrder, typ)
		}
	}
}

// Config is one decoded config file.
type Config interface {
	// Lookup returns the value of a key, and whether the file has the key. A
	// key is a dotted path, such as "server.port". A format can accept a
	// syntax of its own as well, such as a YAML path that starts with "$.".
	Lookup(key string) (any, bool, error)
}

// MapConfig is a [Config] over a decoded tree of maps and slices.
type MapConfig map[string]any

// Lookup satisfies the [Config] interface. A key that the map has, as is,
// wins. Otherwise Lookup walks the tree one dotted part at a time, and a
// part that is a number indexes a slice.
func (m MapConfig) Lookup(key string) (any, bool, error) {
	if v, ok := m[key]; ok {
		return v, true, nil
	}
	var v any = map[string]any(m)
	for part := range strings.SplitSeq(key, ".") {
		switch z := v.(type) {
		case map[string]any:
			var ok bool
			if v, ok = z[part]; !ok {
				return nil, false, nil
			}
		case []any:
			i, err := strconv.Atoi(part)
			if err != nil || i < 0 || len(z) <= i {
				return nil, false, nil
			}
			v = z[i]
		default:
			return nil, false, nil
		}
	}
	return v, true, nil
}

// configFormat is a config file format.
type configFormat struct {
	decode func([]byte) (Config, error)
	exts   []string
}

// RegisterConfigFormat registers the decoder of a config file format, for
// the file extensions that follow, without the dot. It also registers a
// config loader for typ, so that ${typ::key} reads the files of the format.
// The root package registers [JSONT]. The yaml, toml and hcl modules
// register [YAMLT], [TOMLT] and [HCLT].
func RegisterConfigFormat(typ ConfigType, decode func([]byte) (Config, error), exts ...string) {
	formats[typ] = configFormat{
		decode: decode,
		exts:   exts,
	}
	RegisterConfigLoader(typ, func(ctx *Context, key string) (any, bool, error) {
		return ctx.lookupFile(typ, key)
	})
}

// decodeJSON decodes a JSON config file.
func decodeJSON(data []byte) (Config, error) {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var m map[string]any
	if err := dec.Decode(&m); err != nil {
		return nil, err
	}
	return MapConfig(m), nil
}

// configFile is a config file that was read.
type configFile struct {
	path string
	typ  ConfigType
	cfg  Config
}

// lookupFile looks up a key in the config files of typ, or in every config
// file when typ is [AnyT]. A later file wins over an earlier one.
func (ctx *Context) lookupFile(typ ConfigType, key string) (any, bool, error) {
	for _, f := range slices.Backward(ctx.files) {
		if typ != AnyT && f.typ != typ {
			continue
		}
		switch v, ok, err := f.cfg.Lookup(key); {
		case err != nil:
			return nil, false, fmt.Errorf("%s: %s: %w", f.path, key, err)
		case ok:
			return v, true, nil
		}
	}
	return nil, false, ErrUnknownKey
}

// readConfigs reads the config files at paths, in order. A file that does
// not exist is skipped, unless required is true and no file exists.
func (ctx *Context) readConfigs(paths []string, required bool) error {
	for _, path := range paths {
		f, err := readConfig(path)
		switch {
		case errors.Is(err, fs.ErrNotExist):
			continue
		case err != nil:
			return err
		}
		ctx.files = append(ctx.files, f)
	}
	if required && len(ctx.files) == 0 {
		return fmt.Errorf("%w: %s", ErrConfigNotFound, strings.Join(paths, ", "))
	}
	return nil
}

// readConfig reads and decodes one config file. The format comes from the
// extension of the file.
func readConfig(path string) (configFile, error) {
	ext := strings.TrimPrefix(filepath.Ext(path), ".")
	var typ ConfigType
	var format configFormat
	for t, f := range formats {
		if slices.Contains(f.exts, ext) {
			typ, format = t, f
			break
		}
	}
	data, err := os.ReadFile(path)
	switch {
	case err != nil:
		return configFile{}, err
	case format.decode == nil:
		return configFile{}, fmt.Errorf("%w: %s: no format for extension %q", ErrUnknownLoader, path, ext)
	}
	cfg, err := format.decode(data)
	if err != nil {
		return configFile{}, fmt.Errorf("%w: %s: %w", ErrInvalidConfig, path, err)
	}
	return configFile{
		path: path,
		typ:  typ,
		cfg:  cfg,
	}, nil
}

// Fill gives a value to each flag of cmd and of its parents that the command
// line did not set. Each flag takes its value from the first source that has
// one: the environment, the config files, then its default, which is
// expanded only when it wins. Fill first resolves the flag that names the
// config file, then reads the config files once.
func (ctx *Context) Fill(cmd *Command) error {
	if ctx.Vars == nil {
		ctx.Vars = make(Vars)
	}
	root, flags := cmd, scopeFlags(cmd)
	for root.Parent != nil {
		root = root.Parent
	}
	// resolve the config flag, which cannot read a config file
	var paths []string
	explicit := false
	if i := slices.IndexFunc(flags, func(g *Flag) bool { return g.ConfigFile }); i != -1 {
		g := flags[i]
		if err := ctx.fill(root, g); err != nil {
			return err
		}
		if v, ok := ctx.Vars[g.Name]; ok {
			if s, err := v.Get(); err == nil && s != "" {
				paths, explicit = []string{s}, true
			}
		}
	}
	if !explicit {
		for _, path := range root.ConfigFiles {
			v, err := ctx.Expand(path)
			if err != nil {
				return fmt.Errorf("config file %s: %w", path, err)
			}
			s, err := asString[string](v)
			if err != nil {
				return fmt.Errorf("config file %s: %w", path, err)
			}
			paths = append(paths, s)
		}
	}
	// read the config files once
	if !ctx.loaded {
		ctx.loaded = true
		if err := ctx.readConfigs(paths, explicit || root.ConfigRequired); err != nil {
			return err
		}
	}
	// fill the other flags
	for _, g := range flags {
		if g.ConfigFile {
			continue
		}
		if err := ctx.fill(root, g); err != nil {
			return err
		}
	}
	return nil
}

// fill gives a value to one flag that the command line did not set.
func (ctx *Context) fill(root *Command, g *Flag) error {
	if v, ok := ctx.Vars[g.Name]; ok && v.WasSet() || g.Special != "" {
		return nil
	}
	// a hook is filled only when the program binds it by name, because a
	// hook can have side effects
	hook := g.Type == HookT
	// environment
	if name := envName(root, g); name != "" && (!hook || g.EnvVar != "") {
		if s, ok := os.LookupEnv(name); ok {
			if err := ctx.Vars.fill(ctx, g, s); err != nil {
				return fmt.Errorf("--%s: $%s: %w", g.Name, name, err)
			}
			return nil
		}
	}
	// config files
	if !g.NoConfig && (!hook || len(g.Keys) != 0) {
		for _, f := range slices.Backward(ctx.files) {
			key := configKey(g, f.typ)
			switch v, ok, err := f.cfg.Lookup(key); {
			case err != nil:
				return fmt.Errorf("--%s: %s: %s: %w", g.Name, f.path, key, err)
			case ok:
				if err := ctx.Vars.fill(ctx, g, v); err != nil {
					return fmt.Errorf("--%s: %s: %s: %w", g.Name, f.path, key, err)
				}
				return nil
			}
		}
	}
	// default
	if hook || g.Def == nil {
		return nil
	}
	v, err := ctx.Expand(g.Def)
	if err != nil {
		return fmt.Errorf("--%s: cannot expand %v (%T): %w", g.Name, g.Def, g.Def, err)
	}
	if err := ctx.Vars.fill(ctx, g, v); err != nil {
		return fmt.Errorf("--%s: default %v: %w", g.Name, v, err)
	}
	return nil
}

// scopeFlags returns the flags of cmd and of its parents, with a flag of cmd
// before a flag of a parent that has the same name.
func scopeFlags(cmd *Command) []*Flag {
	var flags []*Flag
	seen := make(map[string]bool)
	for c := cmd; c != nil; c = c.Parent {
		if c.Flags == nil {
			continue
		}
		for _, g := range c.Flags.Flags {
			if !seen[g.Name] {
				flags, seen[g.Name] = append(flags, g), true
			}
		}
	}
	return flags
}

// envName returns the name of the environment variable of a flag, or "" when
// the flag has none.
func envName(root *Command, g *Flag) string {
	switch {
	case g.NoEnv:
		return ""
	case g.EnvVar != "":
		return g.EnvVar
	case root.EnvPrefix != "":
		return root.EnvPrefix + "_" + strings.ToUpper(strings.ReplaceAll(g.Name, "-", "_"))
	}
	return ""
}

// configKey returns the config key of a flag for a config type.
func configKey(g *Flag, typ ConfigType) string {
	if key, ok := g.Keys[string(typ)]; ok {
		return key
	}
	if key, ok := g.Keys[""]; ok {
		return key
	}
	return g.Name
}

// fillValues converts a value that is not from the command line to the
// strings to set, and says whether each string is split. A list sets each
// element whole, and a map sets each key and value, in key order.
func fillValues(g *Flag, val any) ([]string, bool, error) {
	if s, ok := val.(string); ok {
		return []string{s}, true, nil
	}
	switch v := reflect.ValueOf(val); v.Kind() {
	case reflect.Slice, reflect.Array:
		if g.Type != SliceT && g.Type != ArrayT {
			return nil, false, fmt.Errorf("%w: a list for a %s flag", ErrInvalidValue, g.Type)
		}
		strs := make([]string, v.Len())
		for i := range v.Len() {
			s, err := marshal[string](v.Index(i).Interface())
			if err != nil {
				return nil, false, fmt.Errorf("index %d: %w", i, err)
			}
			strs[i] = s
		}
		return strs, false, nil
	case reflect.Map:
		if g.Type != MapT {
			return nil, false, fmt.Errorf("%w: a map for a %s flag", ErrInvalidValue, g.Type)
		}
		m := make(map[string]string, v.Len())
		for iter := v.MapRange(); iter.Next(); {
			k, err := marshal[string](iter.Key().Interface())
			if err != nil {
				return nil, false, err
			}
			if m[k], err = marshal[string](iter.Value().Interface()); err != nil {
				return nil, false, fmt.Errorf("key %s: %w", k, err)
			}
		}
		strs := make([]string, 0, len(m))
		for _, k := range slices.Sorted(maps.Keys(m)) {
			strs = append(strs, k+"="+m[k])
		}
		return strs, false, nil
	}
	s, err := marshal[string](val)
	if err != nil {
		return nil, false, err
	}
	return []string{s}, true, nil
}
