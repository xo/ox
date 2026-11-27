package ox

import (
	"context"
	"fmt"
	"io"
	"iter"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/xo/ox/text"
)

// Command is a command.
type Command struct {
	// Parent is the command's parent.
	Parent *Command
	// Exec is the command's exec func.
	Exec ExecFunc
	// Before is the func that runs before the exec func of the command, and
	// of each sub command below it. See [Before].
	Before ExecFunc
	// After is the func that runs after the exec func of the command, and of
	// each sub command below it. See [After].
	After ExecFunc
	// Name is the command's name.
	Name string
	// Usage is the command's usage.
	Usage string
	// Aliases are the command's aliases.
	Aliases []string
	// Suggested are the command's suggested names.
	Suggested []string
	// Flags are the command's flags.
	Flags *FlagSet
	// Commands are the command's sub commands.
	Commands []*Command
	// Args are the command's argument validation func's.
	Args []func([]string) error
	// OnErr indicates whether to continue, panic, or to error when
	// flag/argument parsing fails.
	OnErr OnErr
	// Help is the command's help emitter.
	Help io.WriterTo
	// Comp enables completion for the command.
	Comp bool
	// Section is the help section.
	Section int
	// Hidden indicates the command is hidden from help output.
	Hidden bool
	// Deprecated indicates the command is deprecated.
	Deprecated bool
	// Special is the special value.
	Special string
	// Version is the version string.
	Version string
	// ConfigFiles are the config files to read, in order. Each path is
	// interpolated. Only the root command reads them.
	ConfigFiles []string
	// ConfigRequired makes it an error when no config file exists.
	ConfigRequired bool
	// EnvPrefix is the prefix of the environment variables that fill the
	// flags, such as "APP" for APP_LOG_LEVEL.
	EnvPrefix string
	// ValidArgs are the allowed values of the positional arguments, from
	// [ValidArgs]. Completion offers them.
	ValidArgs []string
	// CompFunc is the func that completes the positional arguments. See
	// [CompFunc].
	CompFunc func(ctx *Context, toComplete string) ([]Completion, CompDirective, error)
	// FlagGroups are the rules between the flags of the command, from
	// [OnlyOneOf], [AllOrNoneOf] and [AtLeastOneOf].
	FlagGroups []FlagGroup
	// StopAtArg makes each argument after the first positional argument of
	// the command a positional argument, even one that starts with `-`. See
	// [StopAtArg].
	StopAtArg bool
	// Abbrev lets the command line give a long flag by a unique prefix of
	// its name. A sub command inherits it. See [Abbrev].
	Abbrev bool
	// Positional are the command's positional arguments, in order. See
	// [Arg].
	Positional []*Argument

	// maxArgs is the maximum of [ValidArgs]. A value < 0 means no maximum.
	maxArgs int
	// noArgs is true when [ValidArgs] allows no argument.
	noArgs bool
}

// NewCommand creates a new command.
func NewCommand(opts ...Option) (*Command, error) {
	cmd := &Command{
		Section: -1,
	}
	if err := Apply(cmd, opts...); err != nil {
		return nil, err
	}
	switch noName := cmd.Name == ""; {
	case noName && cmd.Parent == nil:
		cmd.Name = filepath.Base(os.Args[0])
	case noName:
		return nil, fmt.Errorf("command: %w", ErrUsageNotSet)
	}
	if err := ApplyPost(cmd, opts...); err != nil {
		return nil, err
	}
	if err := cmd.Flags.checkNames(); err != nil {
		return nil, fmt.Errorf("command %s: %w", cmd.Name, err)
	}
	if err := cmd.checkArgNames(); err != nil {
		return nil, fmt.Errorf("command %s: %w", cmd.Name, err)
	}
	return cmd, nil
}

// Sub creates a sub command.
func (cmd *Command) Sub(opts ...Option) error {
	sub, err := NewCommand(prepend(opts, Parent(cmd))...)
	if err != nil {
		return err
	}
	cmd.Commands = append(cmd.Commands, sub)
	return nil
}

// Tree returns the parents for the command.
func (cmd *Command) Tree() []string {
	var v []string
	for c := cmd; c != nil; c = c.Parent {
		v = append(v, c.Name)
	}
	slices.Reverse(v)
	return v
}

// Path returns the execution path for the command.
func (cmd *Command) Path() []string {
	return cmd.Tree()[1:]
}

// RootName returns the root name for the command.
func (cmd *Command) RootName() string {
	return cmd.Tree()[0]
}

// Suggest returns a [SuggestionError] if there is a matching sub command
// within the command's help distance.
func (cmd *Command) Suggest(args ...string) error {
	if len(args) == 0 {
		return nil
	}
	r, maxDist := []rune(foldCase(args[0])), maxDist(cmd)
	type suggest struct {
		cmd  *Command
		dist int
	}
	var suggested []suggest
	for _, c := range cmd.Commands {
		for _, name := range append(prepend(c.Aliases, c.Name), c.Suggested...) {
			if dist := Ldist(r, []rune(foldCase(name))); dist <= maxDist {
				suggested = append(suggested, suggest{
					cmd:  c,
					dist: dist,
				})
			}
		}
	}
	if len(suggested) != 0 {
		sort.Slice(suggested, func(i, j int) bool {
			return suggested[i].dist < suggested[j].dist
		})
		return NewSuggestionError(cmd, args[0], suggested[0].cmd)
	}
	return fmt.Errorf(text.SuggestionErrorMessage, ErrUnknownCommand, args[0], cmd.Name)
}

// Lookup returns the furthest matching command in the command tree.
func (cmd *Command) Lookup(args ...string) *Command {
	c := cmd
	for _, arg := range args {
		if next := c.Command(arg); next != nil {
			c = next
			continue
		}
		break
	}
	return c
}

// Command returns the sub command with the name.
func (cmd *Command) Command(name string) *Command {
	for _, c := range cmd.Commands {
		if c.Name == name || slices.Contains(c.Aliases, name) {
			return c
		}
	}
	return nil
}

// CommandSpecial returns the sub command with the special value.
func (cmd *Command) CommandSpecial(special string) *Command {
	for _, c := range cmd.Commands {
		if c.Special == special {
			return c
		}
	}
	return nil
}

// Flag finds a flag from the command or the command's parents. It does not
// find a local flag of a parent. See [Flag.Local].
func (cmd *Command) Flag(name string, parents, short bool) *Flag {
	for c := cmd; c != nil; c = c.Parent {
		if c.Flags != nil {
			if i := slices.IndexFunc(c.Flags.Flags, func(g *Flag) bool {
				return (c == cmd || !g.Local) &&
					(g.Name == name && !short ||
						g.Short == name && short ||
						slices.ContainsFunc(g.aliases(true), func(s string) bool {
							return s == name && len(s) == 1 == short
						}))
			}); i != -1 {
				return c.Flags.Flags[i]
			}
		}
		if !parents {
			break
		}
	}
	return nil
}

// FlagSpecial returns the flag with a special value.
func (cmd *Command) FlagSpecial(special string) *Flag {
	if cmd.Flags != nil && len(cmd.Flags.Flags) != 0 {
		for _, g := range cmd.Flags.Flags {
			if g.Special == special {
				return g
			}
		}
	}
	return nil
}

// SubSpecial returns the sub command with a special value.
func (cmd *Command) SubSpecial(special string) *Command {
	for _, c := range cmd.Commands {
		if c.Special == special {
			return c
		}
	}
	return nil
}

// Validate validates the passed args, with the funcs of [Command.Args], then
// with the count of [Command.Positional].
func (cmd *Command) Validate(args []string) error {
	// validate args
	for _, f := range cmd.Args {
		if err := f(args); err != nil {
			return newCommandError(cmd.Name, err)
		}
	}
	if err := cmd.checkArgCount(len(args)); err != nil {
		return newCommandError(cmd.Name, err)
	}
	return nil
}

// HelpContext adds the context to the command's help.
func (cmd *Command) HelpContext(ctx *Context) io.WriterTo {
	help := cmd.Help
	if help == nil {
		help, _ = NewCommandHelp(cmd)
	}
	if v, ok := help.(interface{ SetContext(*Context) }); ok {
		v.SetContext(ctx)
	}
	return help
}

// UsageString returns the help of the command, as `--help` writes it. It
// expands no default, because it has no [Context]. Write
// [Command.HelpContext] to a [strings.Builder] to expand them.
func (cmd *Command) UsageString() string {
	var sb strings.Builder
	_, _ = cmd.WriteTo(&sb)
	return sb.String()
}

// WriteTo satisfies the [io.WriterTo] interface.
func (cmd *Command) WriteTo(w io.Writer) (int64, error) {
	help := cmd.Help
	if help == nil {
		var err error
		if help, err = NewCommandHelp(cmd); err != nil {
			return 0, err
		}
	}
	return help.WriteTo(w)
}

// WalkFlags returns an iterator for all flags on the command and its parents.
// It skips a local flag of a parent, because the command does not inherit
// it. See [Flag.Local].
func (cmd *Command) WalkFlags(hidden, deprecated bool) iter.Seq[*Flag] {
	return func(yield func(*Flag) bool) {
		for c := cmd; c != nil; c = c.Parent {
			if c.Flags == nil {
				continue
			}
			for _, g := range c.Flags.Flags {
				switch {
				case !hidden && g.Hidden, !deprecated && g.Deprecated, c != cmd && g.Local:
					continue
				}
				if !yield(g) {
					return
				}
			}
		}
	}
}

// CompCommands returns command completions for the command.
func (cmd *Command) CompCommands(name string, hidden, deprecated bool) ([]Completion, CompDirective) {
	lower, m := foldCase(name), make(map[string]bool)
	var comps []Completion
loop:
	for _, c := range cmd.Commands {
		switch {
		case m[c.Name], !hidden && c.Hidden, !deprecated && c.Deprecated:
			continue
		}
		for _, s := range prepend(c.Aliases, c.Name) {
			if strings.HasPrefix(foldCase(s), lower) {
				comps = append(comps, NewCompletion(c.Name, c.Usage, 0))
				m[c.Name] = true
				continue loop
			}
		}
	}
	if maxDist := maxDist(cmd); 0 < maxDist && len(comps) == 0 {
		r := []rune(lower)
	loop2:
		for _, c := range cmd.Commands {
			switch {
			case m[c.Name], !hidden && c.Hidden, !deprecated && c.Deprecated:
				continue
			}
			for _, s := range prepend(c.Aliases, c.Name) {
				if dist := Ldist([]rune(foldCase(s)), r); dist <= maxDist {
					comps = append(comps, NewCompletion(c.Name, c.Usage, dist))
					m[c.Name] = true
					continue loop2
				}
			}
		}
		sort.Slice(comps, func(i, j int) bool {
			return comps[i].Dist < comps[j].Dist
		})
		if len(comps) != 0 {
			d := comps[0].Dist
			comps = slices.DeleteFunc(comps, func(c Completion) bool {
				return d < c.Dist
			})
		}
	}
	if len(comps) != 0 {
		return comps, CompNoFileComp | CompKeepOrder
	}
	return nil, CompDefault
}

// CompFlags returns flag completions for the command. It offers the name,
// the aliases and the hidden aliases of each flag, and never a deprecated
// alias.
func (cmd *Command) CompFlags(name string, hidden, deprecated, short bool) ([]Completion, CompDirective) {
	var comps []Completion
	lower, m := foldCase(name), make(map[string]bool)
	for g := range cmd.WalkFlags(hidden, deprecated) {
		if !short || len(name) == 0 {
			for _, s := range g.compNames(false) {
				switch long := "--" + s; {
				case m[long]:
					continue
				case strings.HasPrefix(foldCase(s), lower):
					comps = append(comps, NewCompletion(long, g.Usage, 0))
					m[long] = true
				}
			}
		}
		if short {
			for _, s := range g.compNames(true) {
				switch shortstr := "-" + s; {
				case m[shortstr]:
					continue
				case strings.HasPrefix(s, name):
					comps = append(comps, NewCompletion(shortstr, g.Usage, 0))
					m[shortstr] = true
				}
			}
		}
	}
	// distance match
	if maxDist := maxDist(cmd); 0 < maxDist && len(comps) == 0 && len(name) != 1 {
		r := []rune(lower)
		for g := range cmd.WalkFlags(hidden, deprecated) {
			for _, s := range g.compNames(false) {
				long := "--" + s
				if m[long] {
					continue
				}
				if dist := Ldist([]rune(foldCase(s)), r); dist <= maxDist {
					comps = append(comps, NewCompletion(long, g.Usage, dist))
					m[long] = true
				}
			}
		}
		sort.Slice(comps, func(i, j int) bool {
			return comps[i].Dist < comps[j].Dist
		})
		if len(comps) != 0 {
			d := comps[0].Dist
			comps = slices.DeleteFunc(comps, func(c Completion) bool {
				return d < c.Dist
			})
		}
	}
	if len(comps) != 0 {
		return comps, CompNoFileComp | CompKeepOrder
	}
	return nil, CompDefault
}

// longFlag returns the long flag of the command with the name. When no flag
// has the name, and [Command.Abbrev] is on for the command, it returns the
// one flag with a long name that starts with name. A prefix matches the long
// names that completion offers: the name, the aliases and the hidden aliases
// of each flag that is not hidden and not deprecated. The error wraps
// [ErrUnknownFlag], or [ErrAmbiguousFlag] when the prefix matches more than
// one flag.
func (cmd *Command) longFlag(name string) (*Flag, error) {
	if g := cmd.Flag(name, true, false); g != nil {
		return g, nil
	}
	if name == "" || !cmd.abbrev() {
		return nil, ErrUnknownFlag
	}
	var flags []*Flag
	var names []string
	seen := make(map[string]bool)
	for g := range cmd.WalkFlags(true, true) {
		for _, s := range prepend(g.aliases(true), g.Name) {
			switch {
			case len(s) == 1 && s != g.Name, seen[s]:
				// a short alias, or the name of a nearer flag
				continue
			}
			seen[s] = true
			switch {
			case g.Hidden, g.Deprecated, slices.Contains(g.DeprecatedAliases, s), !strings.HasPrefix(s, name):
				continue
			}
			names = append(names, "--"+s)
			if !slices.Contains(flags, g) {
				flags = append(flags, g)
			}
		}
	}
	switch len(flags) {
	case 0:
		return nil, ErrUnknownFlag
	case 1:
		return flags[0], nil
	}
	slices.Sort(names)
	return nil, fmt.Errorf("%w: %s", ErrAmbiguousFlag, strings.Join(names, ", "))
}

// abbrev returns true when [Command.Abbrev] is on for the command or for a
// parent.
func (cmd *Command) abbrev() bool {
	for c := cmd; c != nil; c = c.Parent {
		if c.Abbrev {
			return true
		}
	}
	return false
}

// inheritedFlags returns the flags that the command inherits from its
// parents, from the nearest parent first. A local flag is not inherited. A
// flag with the name of a nearer flag is not inherited, because the nearer
// flag hides it.
func (cmd *Command) inheritedFlags() []*Flag {
	var flags []*Flag
	seen := make(map[string]bool)
	for c := cmd; c != nil; c = c.Parent {
		if c.Flags == nil {
			continue
		}
		for _, g := range c.Flags.Flags {
			switch {
			case c != cmd && g.Local, seen[g.Name]:
				continue
			}
			seen[g.Name] = true
			if c != cmd {
				flags = append(flags, g)
			}
		}
	}
	return flags
}

// checkArgCount returns an error when n args do not fit the positional
// arguments of the command: [ErrMissingArgument] for too few, and
// [ErrInvalidArgCount] for too many.
func (cmd *Command) checkArgCount(n int) error {
	if len(cmd.Positional) == 0 {
		return nil
	}
	for i, a := range cmd.Positional {
		if n <= i && !a.Optional {
			return fmt.Errorf("%w: %s", ErrMissingArgument, a.Flag.Name)
		}
	}
	if m := len(cmd.Positional); m < n && !cmd.Positional[m-1].Variadic() {
		return fmt.Errorf("%w: takes max %d arg(s)", ErrInvalidArgCount, m)
	}
	return nil
}

// addArgs adds positional arguments to the command. It returns an error when
// an argument follows a variadic argument, or when a required argument
// follows an optional argument.
func (cmd *Command) addArgs(args ...*Argument) error {
	for _, a := range args {
		if n := len(cmd.Positional); n != 0 {
			switch prev := cmd.Positional[n-1]; {
			case prev.Variadic():
				return fmt.Errorf("%w: %s follows the variadic argument %s", ErrInvalidArg, a.Flag.Name, prev.Flag.Name)
			case prev.Optional && !a.Optional:
				return fmt.Errorf("%w: the required argument %s follows the optional argument %s", ErrInvalidArg, a.Flag.Name, prev.Flag.Name)
			}
		}
		cmd.Positional = append(cmd.Positional, a)
	}
	return nil
}

// checkArgNames returns an error when two positional arguments of the
// command have the same name, or when an argument has the name of a flag of
// the command. An argument and a flag keep their values in the same
// [Vars].
func (cmd *Command) checkArgNames() error {
	seen := make(map[string]bool)
	for _, a := range cmd.Positional {
		name := a.Flag.Name
		switch {
		case seen[name]:
			return fmt.Errorf("%w: two arguments are named %s", ErrInvalidArg, name)
		case cmd.Flag(name, false, false) != nil:
			return fmt.Errorf("%w: the argument %s has the name of a flag", ErrInvalidArg, name)
		}
		seen[name] = true
	}
	return nil
}

// argAt returns the positional argument at index i, or the last argument
// when it is variadic and i is past it. It returns nil when the command
// takes no argument at i.
func (cmd *Command) argAt(i int) *Argument {
	switch n := len(cmd.Positional); {
	case i < n:
		return cmd.Positional[i]
	case n != 0 && cmd.Positional[n-1].Variadic():
		return cmd.Positional[n-1]
	}
	return nil
}

// FlagSet is a set of command-line flag definitions.
type FlagSet struct {
	Flags []*Flag
}

// Flags creates a new flag set from the options.
func Flags() *FlagSet {
	return new(FlagSet)
}

// Option returns a [CommandOption] for the flag set.
func (fs *FlagSet) Option() option {
	return option{
		name: "FlagSet",
		cmd: func(cmd *Command) error {
			cmd.Flags = fs
			return nil
		},
	}
}

// With adds the flags to the flag set.
func (fs *FlagSet) With(flags ...*Flag) *FlagSet {
	if fs == nil {
		fs = Flags()
	}
	fs.Flags = append(fs.Flags, flags...)
	return fs
}

// Var adds a variable to the flag set.
func (fs *FlagSet) Var(name, usage string, opts ...Option) *FlagSet {
	if fs == nil {
		fs = Flags()
	}
	g, err := NewFlag(name, usage, opts...)
	if err != nil {
		panic(err)
	}
	fs.Flags = append(fs.Flags, g)
	return fs
}

// String adds a string variable to the flag set.
func (fs *FlagSet) String(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(StringT))...)
}

// Bytes adds a []byte variable to the flag set.
func (fs *FlagSet) Bytes(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(BytesT))...)
}

// Base64 adds a base64 encoded []byte variable to the flag set.
func (fs *FlagSet) Base64(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Base64T))...)
}

// Hex adds a hex encoded []byte variable to the flag set.
func (fs *FlagSet) Hex(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(HexT))...)
}

// Bool adds a bool variable to the flag set.
func (fs *FlagSet) Bool(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(BoolT))...)
}

// Byte adds a byte variable to the flag set.
func (fs *FlagSet) Byte(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(ByteT))...)
}

// Rune adds a rune variable to the flag set.
func (fs *FlagSet) Rune(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(RuneT))...)
}

// Int64 adds a int64 variable to the flag set.
func (fs *FlagSet) Int64(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Int64T))...)
}

// Int32 adds a int32 variable to the flag set.
func (fs *FlagSet) Int32(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Int32T))...)
}

// Int16 adds a int16 variable to the flag set.
func (fs *FlagSet) Int16(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Int16T))...)
}

// Int8 adds a int8 variable to the flag set.
func (fs *FlagSet) Int8(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Int8T))...)
}

// Int adds a int variable to the flag set.
func (fs *FlagSet) Int(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(IntT))...)
}

// Uint64 adds a uint64 variable to the flag set.
func (fs *FlagSet) Uint64(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Uint64T))...)
}

// Uint32 adds a uint32 variable to the flag set.
func (fs *FlagSet) Uint32(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Uint32T))...)
}

// Uint16 adds a uint16 variable to the flag set.
func (fs *FlagSet) Uint16(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Uint16T))...)
}

// Uint8 adds a uint8 variable to the flag set.
func (fs *FlagSet) Uint8(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Uint8T))...)
}

// Uint adds a uint variable to the flag set.
func (fs *FlagSet) Uint(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(UintT))...)
}

// Float64 adds a float64 variable to the flag set.
func (fs *FlagSet) Float64(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Float64T))...)
}

// Float32 adds a float32 variable to the flag set.
func (fs *FlagSet) Float32(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Float32T))...)
}

// Complex128 adds a complex128 variable to the flag set.
func (fs *FlagSet) Complex128(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Complex128T))...)
}

// Complex64 adds a complex64 variable to the flag set.
func (fs *FlagSet) Complex64(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(Complex64T))...)
}

// Timestamp adds a [time.Time] variable to the flag set in the expected format
// of [time.RFC3339].
func (fs *FlagSet) Timestamp(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(TimestampT))...)
}

// DateTime adds a [time.Time] variable to the flag set in the expected format
// of [time.DateTime].
func (fs *FlagSet) DateTime(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(DateTimeT))...)
}

// Date adds a [time.Time] variable to the flag set in the expected format of
// [time.DateOnly].
func (fs *FlagSet) Date(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(DateT))...)
}

// Duration adds a [time.Duration] variable to the flag set.
func (fs *FlagSet) Duration(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(DurationT))...)
}

// Size adds a [Size] variable to the flag set.
func (fs *FlagSet) Size(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(SizeT))...)
}

// Rate adds a [Rate] variable to the flag set.
func (fs *FlagSet) Rate(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(RateT))...)
}

// Path adds a path variable to the flag set. See [PathT].
func (fs *FlagSet) Path(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(PathT))...)
}

// File adds a variable to the flag set for a file that exists. See [FileT].
func (fs *FlagSet) File(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(FileT))...)
}

// Dir adds a variable to the flag set for a directory that exists. See
// [DirT].
func (fs *FlagSet) Dir(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(DirT))...)
}

// Count adds a count variable to the flag set.
func (fs *FlagSet) Count(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(CountT))...)
}

// BigInt adds a [math/big.Int] variable to the flag set.
func (fs *FlagSet) BigInt(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(BigIntT))...)
}

// BigFloat adds a [math/big.Float] variable to the flag set.
func (fs *FlagSet) BigFloat(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(BigFloatT))...)
}

// BigRat adds a [math/big.Rat] variable to the flag set.
func (fs *FlagSet) BigRat(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(BigRatT))...)
}

// Addr adds a [netip.Addr] variable to the flag set.
func (fs *FlagSet) Addr(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(AddrT))...)
}

// AddrPort adds a [netip.AddrPort] variable to the flag set.
func (fs *FlagSet) AddrPort(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(AddrPortT))...)
}

// CIDR adds a [netip.Prefix] variable to the flag set.
func (fs *FlagSet) CIDR(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(CIDRT))...)
}

// Regexp adds a [regexp.Regexp] variable to the flag set.
func (fs *FlagSet) Regexp(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(RegexpT))...)
}

// URL adds a [url.URL] variable to the flag set.
func (fs *FlagSet) URL(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(URLT))...)
}

// UUID adds a uuid variable to the flag set.
func (fs *FlagSet) UUID(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(UUIDT))...)
}

// Color adds a color variable to the flag set.
func (fs *FlagSet) Color(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(ColorT))...)
}

// Glob adds a glob variable to the flag set.
func (fs *FlagSet) Glob(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(GlobT))...)
}

// Slice adds a slice variable to the flag set.
func (fs *FlagSet) Slice(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(SliceT))...)
}

// Array adds a array variable to the flag set.
func (fs *FlagSet) Array(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(ArrayT))...)
}

// Map adds a map variable to the flag set.
func (fs *FlagSet) Map(name, usage string, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(MapT))...)
}

// Hook adds a hook to the flag set.
func (fs *FlagSet) Hook(name, usage string, f any, opts ...Option) *FlagSet {
	return fs.Var(name, usage, prepend(opts, Option(HookT), NoArg(true, ""), Default(f))...)
}

// checkNames returns an error when two flags of the set have the same long
// or short name, counting their aliases.
func (fs *FlagSet) checkNames() error {
	if fs == nil {
		return nil
	}
	seen := make(map[string]*Flag)
	for _, g := range fs.Flags {
		// the name is long, the short name is short, and an alias of one
		// letter is short, as Command.Flag finds them
		names := []string{"--" + g.Name}
		if g.Short != "" {
			names = append(names, "-"+g.Short)
		}
		for _, s := range g.aliases(true) {
			if len([]rune(s)) == 1 {
				names = append(names, "-"+s)
			} else {
				names = append(names, "--"+s)
			}
		}
		for _, name := range names {
			if prev, ok := seen[name]; ok && prev != g {
				return fmt.Errorf("%w: %s", ErrDuplicateFlag, name)
			}
			seen[name] = g
		}
	}
	return nil
}

// Flag is a command-line flag variable definition.
type Flag struct {
	// Type is the flag's [Type].
	Type Type
	// MapKey is the flag's map key type when the flag is a [MapT].
	MapKey Type
	// Elem is the flag's element type when the flag is a [SliceT], [ArrayT] or [MapT].
	Elem Type
	// Name is the flag's long name (`--arg`).
	Name string
	// Usage is the flag's usage.
	Usage string
	// Short is the flag's short, single letter name (`-a`).
	Short string
	// Aliases are the flag's long and short aliases.
	Aliases []string
	// Spec is the flag's help spec.
	Spec string
	// Def is the flag's default value.
	Def any
	// NoDefault hides the "(default: ...)" text of the flag in help.
	NoDefault bool
	// Optional shows the flag's spec in help as an optional value, such as
	// `--help[=options]`. See [OptionalValue].
	Optional bool
	// HelpValues are the values of the flag that help writes on lines of
	// their own, such as `--help=commands`. See [HelpValue].
	HelpValues []HelpLine
	// NoArg when true, indicates that the flag does not require an argument.
	NoArg bool
	// NoArgDef is the flag's default value when no argument was passed for the
	// flag.
	NoArgDef any
	// Binds are the bound variables that will be additionally set when the
	// flag is encountered on the command-line.
	Binds []Binder
	// Keys are the flag's config look up keys.
	Keys map[string]string
	// Section is the flag's help section.
	Section int
	// Hidden indicates the flag is hidden from help output.
	Hidden bool
	// Deprecated indicates the flag is deprecated.
	Deprecated bool
	// Local keeps the flag to its command. A sub command does not inherit
	// it, so the command line of a sub command refuses it. See [Local].
	Local bool
	// Split is a split separator, to split values for slices/arrays/maps.
	Split string
	// Special is the flag's special value.
	Special string
	// Valid are the flag's validators.
	Valid []func(any) error
	// Allowed are the allowed values.
	Allowed []any
	// EnvVar is the environment variable that fills the flag, whatever the
	// prefix of the root command.
	EnvVar string
	// NoEnv keeps the flag out of the environment.
	NoEnv bool
	// NoConfig keeps the flag out of the config files.
	NoConfig bool
	// ConfigFile indicates that the flag names the config file.
	ConfigFile bool
	// HiddenAliases are the flag's long and short aliases that help does not
	// write. Completion offers them.
	HiddenAliases []string
	// DeprecatedAliases are the flag's long and short aliases that neither
	// help nor completion shows.
	DeprecatedAliases []string
	// Abs makes the value of a path flag absolute. See [Abs].
	Abs bool
	// Real resolves each symbolic link in the value of a path flag. See
	// [Real].
	Real bool
	// Home shows a path flag's default under the home directory as `~` in
	// help. See [Home].
	Home bool
	// Exts are the file extensions that file completion offers for a path
	// flag. See [Ext].
	Exts []string
	// NotExist makes the value of a path flag a path that must not exist.
	// See [NotExist].
	NotExist bool
	// CompFunc is the func that completes the flag's value. See [CompFunc].
	CompFunc func(ctx *Context, toComplete string) ([]Completion, CompDirective, error)
	// Required makes the parse fail when the flag gets no value from the
	// command line, the environment or a config file. See [Required].
	Required bool

	// groups are the groups of the flag from its struct tag, such as
	// `onlyoneof:format`.
	groups []groupTag
}

// Rule is a rule between the flags of a [FlagGroup].
type Rule int

// Rules.
const (
	// RuleOnlyOneOf allows one flag of the group at most. See [OnlyOneOf].
	RuleOnlyOneOf Rule = iota
	// RuleAllOrNoneOf needs each flag of the group, or none of them. See
	// [AllOrNoneOf].
	RuleAllOrNoneOf
	// RuleAtLeastOneOf needs one flag of the group at least. See
	// [AtLeastOneOf].
	RuleAtLeastOneOf
)

// FlagGroup is a group of flags of a command, and the rule between them. A
// flag of the group counts when it gets a value from the command line, the
// environment or a config file. A default does not count.
type FlagGroup struct {
	// Rule is the rule between the flags.
	Rule Rule
	// Names are the long names of the flags.
	Names []string
}

// check returns an error when the flags of the group break its rule. isSet
// tells whether a flag has a value.
func (grp FlagGroup) check(isSet func(string) bool) error {
	var set, unset []string
	for _, name := range grp.Names {
		if isSet(name) {
			set = append(set, name)
		} else {
			unset = append(unset, name)
		}
	}
	switch {
	case grp.Rule == RuleOnlyOneOf && len(set) > 1:
		return fmt.Errorf("%w: %s", ErrOnlyOneOf, flagNames(set))
	case grp.Rule == RuleAllOrNoneOf && len(set) != 0 && len(unset) != 0:
		return fmt.Errorf("%w: %s (missing %s)", ErrAllOrNoneOf, flagNames(grp.Names), flagNames(unset))
	case grp.Rule == RuleAtLeastOneOf && len(set) == 0:
		return fmt.Errorf("%w: %s", ErrAtLeastOneOf, flagNames(grp.Names))
	}
	return nil
}

// groupTag is the group of a flag from a struct tag, such as
// `onlyoneof:format`.
type groupTag struct {
	rule Rule
	name string
}

// addTagGroups adds to cmd a [FlagGroup] for each group that the struct tags
// of flags name. The flags with the same rule and the same group name are
// one group.
func addTagGroups(cmd *Command, flags []*Flag) {
	index := make(map[groupTag]int)
	for _, g := range flags {
		for _, tag := range g.groups {
			i, ok := index[tag]
			if !ok {
				i = len(cmd.FlagGroups)
				index[tag] = i
				cmd.FlagGroups = append(cmd.FlagGroups, FlagGroup{Rule: tag.rule})
			}
			cmd.FlagGroups[i].Names = append(cmd.FlagGroups[i].Names, g.Name)
		}
	}
}

// flagNames returns the names as the command line gives them, such as
// `--json, --yaml`.
func flagNames(names []string) string {
	var sb strings.Builder
	for i, name := range names {
		if i != 0 {
			_, _ = sb.WriteString(", ")
		}
		_, _ = sb.WriteString("--")
		_, _ = sb.WriteString(name)
	}
	return sb.String()
}

// HelpLine is a name and its usage, which help writes on a line of its own,
// such as a value of a flag from [HelpValue], or a positional argument from
// [ArgUsage].
type HelpLine struct {
	// Name is the name, such as `commands` for `--help=commands`, or `DSN`.
	Name string
	// Usage is the usage of the name.
	Usage string
}

// NewFlag creates a new command-line flag.
func NewFlag(name, usage string, opts ...Option) (*Flag, error) {
	g := &Flag{
		Type:    StringT,
		MapKey:  StringT,
		Elem:    StringT,
		Name:    name,
		Usage:   usage,
		Section: -1,
	}
	if err := Apply(g, opts...); err != nil {
		return nil, err
	}
	if extra, ok := typeFlagOpts[g.Type]; ok {
		if err := Apply(g, extra...); err != nil {
			return nil, err
		}
	}
	switch {
	case g.Name == "":
		return nil, ErrInvalidFlagName
	case g.NoArg && g.NoArgDef == nil:
		return nil, fmt.Errorf("flag %s: %w: NoArgDef cannot be nil", g.Name, ErrInvalidValue)
	case g.Type == HookT && g.Def == nil:
		return nil, fmt.Errorf("flag %s: %w: Def cannot be nil", g.Name, ErrInvalidValue)
	}
	return g, nil
}

// Argument is a positional argument of a command, from [Arg].
type Argument struct {
	// Flag holds the name, the usage, the type, the default, the validators,
	// the binds and the completion of the argument, as for a flag.
	Flag *Flag
	// Optional is true when the command runs without the argument. An
	// argument with a default is optional.
	Optional bool
}

// NewArgument creates a positional argument. It applies each [ArgOption],
// such as [Optional], to the argument, and each other option, such as a
// [Type], [Valid] or [Bind], to its [Argument.Flag], as [NewFlag] does.
func NewArgument(name, usage string, opts ...Option) (*Argument, error) {
	a := new(Argument)
	var flagOpts []Option
	for _, o := range opts {
		if opt := o.Option(); opt.arg != nil {
			if err := opt.apply(a); err != nil {
				return nil, fmt.Errorf("argument %s: %w", name, err)
			}
			continue
		}
		flagOpts = append(flagOpts, o)
	}
	g, err := NewFlag(name, usage, flagOpts...)
	if err != nil {
		return nil, fmt.Errorf("argument %s: %w", name, err)
	}
	a.Flag, a.Optional = g, a.Optional || g.Def != nil
	return a, nil
}

// Variadic returns true when the argument takes each argument that is left,
// which is when its type is a [SliceT] or an [ArrayT].
func (a *Argument) Variadic() bool {
	return a.Flag.Type == SliceT || a.Flag.Type == ArrayT
}

// SpecString returns the spec of the argument in the usage line of help,
// such as `file`, `[file]`, `file...` or `[file...]`. The name is the
// [Flag.Spec] of the argument, or its name.
func (a *Argument) SpecString() string {
	s := a.name()
	if a.Variadic() {
		s += text.CommandArgsMore
	}
	if a.Optional {
		return fmt.Sprintf(text.CommandArgOptionalSpec, s)
	}
	return fmt.Sprintf(text.CommandArgSpec, s)
}

// name returns the name of the argument that help writes: its [Flag.Spec],
// or its name.
func (a *Argument) name() string {
	if a.Flag.Spec != "" {
		return a.Flag.Spec
	}
	return a.Flag.Name
}

// elem returns the type of each value of the argument: the element type of a
// variadic argument, or the type of the argument.
func (a *Argument) elem() Type {
	if a.Variadic() {
		return a.Flag.Elem
	}
	return a.Flag.Type
}

// FlagsFrom creates flags for a value of type *struct using reflection. Builds
// flags for all exported fields with a `ox` tag, and a non-empty description.
//
// A `ox` tag starts with the flag's description, followed by one or more
// options separated by `,`. If a flag description must contain a comma (`,`)
// it can be escaped with a double backslash (`\\`). The flag's name is
// determined by the result of calling [DefaultFlagNameMapper], or it can be
// set with the `name:` option (see below).
//
// Use the [CommandOption] [From] to add flags when creating a command with
// [Run]/[RunContext]/[Sub]/[NewCommand].
//
// Example:
//
//	args := struct{
//		MyFlag      string   `ox:"my flag,short:f,default:$USER"`
//		MyVerbosity int      `ox:"my verbosity,type:count,short:v"`
//		MyURL       *url.URL `ox:"my url,set:MyURLSet"`
//		MyURLSet    bool
//		MyOtherFlag string   `ox:"a long\\, long description,short:F"`
//		MyFloat     float64  `ox:"my float,hidden,name:MYF"`
//	}{}
//
//	ox.FlagsFrom(&args)
//
// Recognized options:
//
//	type - sets the flag's field type
//	mapkey - sets the flag's map key type
//	elem - sets the flag's slice/array/map element type
//	name - sets the flag's name
//	short - sets the flag's short (single character) name
//	alias - adds a alias to the flag
//	aliases - adds multiple aliases, separated by `|` to the flag
//	hiddenaliases - adds aliases, separated by `|`, that help does not write
//	deprecatedaliases - adds aliases, separated by `|`, that help and completion do not show
//	spec - sets the flag's use spec
//	default - sets the flag's default value
//	nodefault - hides the flag's default value in help
//	noarg - sets the flag as requiring no argument, and the default value for the flag when toggled
//	key - sets the flag's lookup config key
//	hook - sets the flag's special value to `hook:<type>`, and can be used to hook [Defaults]'s flags
//	section - sets the flag's section
//	hidden - marks the flag as hidden
//	deprecated - marks the flag as deprecated
//	local - keeps the flag to its command, so that a sub command does not inherit it
//	split - set a split separator for slice/array/map values
//	set - binds the flag's set value to a bool field in the *struct of the name
//	required - marks the flag as required, see [Required]
//	onlyoneof - adds the flag to a group of the name, see [OnlyOneOf]
//	allornoneof - adds the flag to a group of the name, see [AllOrNoneOf]
//	atleastoneof - adds the flag to a group of the name, see [AtLeastOneOf]
//	arg - makes the field a positional argument and not a flag, in the order of the fields
//	optional - makes a positional argument optional
//
// [From] and [App] add each positional argument to the command, as [Arg]
// does. FlagsFrom leaves them out.
//
// The `default:` option will be processed by [Context.Interpolate] when the
// command's flags are populated.
//
// FlagsFrom keeps the group of the `onlyoneof`, `allornoneof` and
// `atleastoneof` options on each flag. [From] and [App] add the groups to the
// command as [Command.FlagGroups].
//
// The tag name (`ox`) can be changed by setting the [DefaultStructTagName] variable
// if necessary.
func FlagsFrom[T *E, E any](val T) ([]*Flag, error) {
	flags, _, err := appendFlags(nil, nil, reflect.ValueOf(val), nil)
	return flags, err
}

// New creates a new value for the flag's type.
func (g *Flag) New(ctx *Context) (Value, error) {
	switch g.Type {
	case SliceT:
		return NewSlice(g.Elem), nil
	case ArrayT:
		return NewArray(g.Elem), nil
	case MapT:
		return NewMap(g.MapKey, g.Elem)
	case HookT:
		return newHook(ctx, g.Def)
	}
	return g.Type.New()
}

// SpecString returns the spec string for the flag. A hook flag shows its
// [Flag.Spec], and shows only its name when it has no spec.
func (g *Flag) SpecString() string {
	return g.Name + g.specArg()
}

// Validate checks a value against the flag's validator funcs.
func (g *Flag) Validate(v any) error {
	for _, f := range g.Valid {
		if err := f(v); err != nil {
			if g.Allowed != nil {
				s := g.Formatter()
				return fmt.Errorf("%w: "+s+": allowed values: %s", ErrInvalidValue, v, allowed(s, g.Allowed))
			}
			return fmt.Errorf("%w: "+g.Formatter()+" (%w)", ErrInvalidValue, v, err)
		}
	}
	return nil
}

// Formatter returns the value format string for the flag.
func (g *Flag) Formatter() string {
	switch g.Type {
	case SliceT, ArrayT, MapT:
		return g.Elem.Formatter()
	}
	return g.Type.Formatter()
}

// specArg returns the part of the spec string after the flag's name.
func (g *Flag) specArg() string {
	switch {
	case g.Optional && g.Spec != "":
		// an optional value must follow an =, so the spacer does not apply
		return "[=" + g.Spec + "]"
	case g.Type == HookT && g.Spec != "" && !g.NoArg:
		// a hook that takes a value, such as usql's -v NAME=VALUE; a hook
		// that takes none, such as --help, has no spec to show even when the
		// command's Spec reached it
		return text.FlagSpecSpacer + g.Spec
	case g.NoArg || g.Type == HookT:
		return ""
	case g.Spec != "":
		return text.FlagSpecSpacer + g.Spec
	case g.Type == SliceT, g.Type == ArrayT:
		return text.FlagSpecSpacer + g.Elem.String()
	case g.Type == MapT:
		return text.FlagSpecSpacer + g.MapKey.String() + "=" + g.Elem.String()
	}
	return text.FlagSpecSpacer + g.Type.String()
}

// aliases returns the flag's aliases and hidden aliases. When deprecated is
// true, it adds the deprecated aliases.
func (g *Flag) aliases(deprecated bool) []string {
	v := slices.Concat(g.Aliases, g.HiddenAliases)
	if deprecated {
		v = append(v, g.DeprecatedAliases...)
	}
	return v
}

// compNames returns the flag's long names, or its short names, that
// completion offers, without the leading dashes.
func (g *Flag) compNames(short bool) []string {
	var v []string
	for _, s := range prepend(g.aliases(false), g.Name, g.Short) {
		if s != "" && utf8.RuneCountInString(s) == 1 == short && !slices.Contains(v, s) {
			v = append(v, s)
		}
	}
	return v
}

// ExecType is the interface for func's that can be used with [Run],
// [NewCommand], and [Exec].
type ExecType interface {
	func(context.Context, []string) error |
		func(context.Context, []string) |
		func(context.Context) error |
		func(context.Context) |
		func([]string) error |
		func([]string) |
		func() error |
		func()
}

// ExecFunc wraps a exec func.
type ExecFunc func(context.Context, []string) error

// NewExec creates a [ExecFunc] func.
func NewExec[T ExecType](f T) (ExecFunc, error) {
	var v any = f
	switch f := v.(type) {
	case func(context.Context, []string) error:
		return f, nil
	case func(context.Context, []string):
		return func(ctx context.Context, args []string) error {
			f(ctx, args)
			return nil
		}, nil
	case func(context.Context) error:
		return func(ctx context.Context, _ []string) error {
			return f(ctx)
		}, nil
	case func(context.Context):
		return func(ctx context.Context, _ []string) error {
			f(ctx)
			return nil
		}, nil
	case func([]string) error:
		return func(_ context.Context, args []string) error {
			return f(args)
		}, nil
	case func([]string):
		return func(_ context.Context, args []string) error {
			f(args)
			return nil
		}, nil
	case func() error:
		return func(context.Context, []string) error {
			return f()
		}, nil
	case func():
		return func(context.Context, []string) error {
			f()
			return nil
		}, nil
	}
	return nil, fmt.Errorf("%w: invalid exec func %T", ErrInvalidType, f)
}

// appendFlags builds and appends flags for value v to flags, and the
// positional arguments of the fields with the tag option `arg` to args.
func appendFlags(flags []*Flag, args []*Argument, v reflect.Value, parents []string) ([]*Flag, []*Argument, error) {
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return nil, nil, fmt.Errorf("%w: %s is not a *struct", ErrInvalidType, v.Type())
	}
	v = v.Elem()
	typ := v.Type()
	for i := range typ.NumField() {
		// check field is exported
		f := typ.Field(i)
		tag, ok := f.Tag.Lookup(DefaultStructTagName)
		if !ok {
			continue
		}
		if r := []rune(f.Name); !unicode.IsUpper(r[0]) {
			return nil, nil, fmt.Errorf("%w: unexported field `%s` has tag `%s`", ErrInvalidType, f.Name, DefaultStructTagName)
		}
		field := v.Field(i)
		tags := SplitBy(tag, ',')
		isField := isField(field)
		switch name, kind := tags[0], f.Type.Kind(); {
		case !isField && kind == reflect.Pointer && f.Type.Elem().Kind() == reflect.Struct: // *struct
			switch name {
			case "-":
				continue
			case "":
				name = f.Name
			}
			if field.IsNil() {
				field.Set(reflect.New(f.Type.Elem()))
			}
			var err error
			if flags, args, err = appendFlags(flags, args, field, append(parents, name)); err != nil {
				return nil, nil, err
			}
		case !isField && kind == reflect.Struct: // struct
			switch name {
			case "-":
				continue
			case "":
				name = f.Name
			}
			var err error
			if flags, args, err = appendFlags(flags, args, field.Addr(), append(parents, name)); err != nil {
				return nil, nil, err
			}
		default:
			switch name {
			case "-":
				continue
			case "":
				// the doc comment of the field, from RegisterDoc
				if tags[0] = fieldDoc(typ, f.Name); tags[0] == "" {
					continue
				}
			}
			g, a, err := buildFlag(v, i, tags, parents)
			switch {
			case err != nil:
				return nil, nil, err
			case a != nil:
				args = append(args, a)
			default:
				flags = append(flags, g)
			}
		}
	}
	return flags, args, nil
}

// buildFlag builds a flag for the field i of v, or a positional argument
// when the tag has the option `arg`.
func buildFlag(v reflect.Value, i int, tag []string, parents []string) (*Flag, *Argument, error) {
	f := v.Type().Field(i)
	// flag name
	name := f.Name
	if len(parents) != 0 {
		name = strings.Join(parents, "_") + "_" + name
	}
	// the options of a positional argument
	var arg, optional bool
	tags := make([]string, 0, len(tag)-1)
	for _, s := range tag[1:] {
		switch s {
		case "arg":
			arg = true
		case "optional":
			optional = true
		default:
			tags = append(tags, s)
		}
	}
	if optional && !arg {
		return nil, nil, fmt.Errorf("field %s: %w: optional without arg", f.Name, ErrInvalidValue)
	}
	// flag opts
	opts, err := buildFlagOpts(v, v.Field(i).Addr(), tags)
	if err != nil {
		return nil, nil, fmt.Errorf("field %s: %w", f.Name, err)
	}
	if arg {
		if optional {
			opts = append(opts, Optional())
		}
		a, err := NewArgument(DefaultFlagNameMapper(name), tag[0], opts...)
		if err != nil {
			return nil, nil, fmt.Errorf("field %s: %w", f.Name, err)
		}
		return nil, a, nil
	}
	// create flag
	g, err := NewFlag(DefaultFlagNameMapper(name), tag[0], opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("field %s: cannot create flag: %w", f.Name, err)
	}
	return g, nil, nil
}

// buildFlagOpts builds flag options for value from the passed struct tags.
func buildFlagOpts(parent, value reflect.Value, tags []string) ([]Option, error) {
	var set *bool
	typ, mapKey, elem, err := defaultType(value.Elem().Type())
	if err != nil {
		return nil, fmt.Errorf("%s: %w", value.Type().String(), err)
	}
	opts := []Option{typ, MapKey(mapKey), Elem(elem)}
	if typ == SliceT || typ == MapT {
		opts = append(opts, Split(","))
	}
	for _, opt := range tags {
		switch key, val, _ := strings.Cut(opt, ":"); key {
		case "type":
			opts = append(opts, Type(val))
		case "mapkey":
			opts = append(opts, MapKey(Type(val)))
		case "elem":
			opts = append(opts, Elem(Type(val)))
		case "name":
			opts = append(opts, Name(val))
		case "short":
			opts = append(opts, Short(val))
		case "alias":
			opts = append(opts, Aliases(val))
		case "aliases":
			opts = append(opts, Aliases(strings.Split(val, "|")...))
		case "hiddenaliases":
			opts = append(opts, HiddenAliases(strings.Split(val, "|")...))
		case "deprecatedaliases":
			opts = append(opts, DeprecatedAliases(strings.Split(val, "|")...))
		case "spec":
			opts = append(opts, Spec(val))
		case "default":
			opts = append(opts, Default(val))
		case "nodefault":
			opts = append(opts, NoDefault())
		case "noarg":
			opts = append(opts, NoArg(true, val))
		case "key":
			typ, key, _ := strings.Cut(val, "|")
			if key == "" {
				typ, key = key, typ
			}
			opts = append(opts, Key(typ, key))
		case "hook":
			opts = append(opts, Special("hook:"+val))
		case "section":
			section, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("%w: section %s: %w", ErrInvalidConversion, val, err)
			}
			opts = append(opts, Section(section))
		case "config":
			opts = append(opts, ConfigFlag())
		case "env":
			opts = append(opts, EnvVar(val))
		case "noenv":
			opts = append(opts, NoEnv())
		case "noconfig":
			opts = append(opts, NoConfig())
		case "hidden":
			opts = append(opts, Hidden(true))
		case "deprecated":
			opts = append(opts, Deprecated(true))
		case "local":
			opts = append(opts, Local())
		case "split":
			opts = append(opts, Split(val))
		case "abs":
			opts = append(opts, Abs())
		case "real":
			opts = append(opts, Real())
		case "home":
			opts = append(opts, Home())
		case "ext":
			opts = append(opts, Ext(strings.Split(val, "|")...))
		case "notexist":
			opts = append(opts, NotExist())
		case "required":
			opts = append(opts, Required())
		case "onlyoneof":
			opts = append(opts, groupOption(RuleOnlyOneOf, val))
		case "allornoneof":
			opts = append(opts, groupOption(RuleAllOrNoneOf, val))
		case "atleastoneof":
			opts = append(opts, groupOption(RuleAtLeastOneOf, val))
		case "set":
			if set, err = setField(parent, val); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("%w: %q", ErrUnknownTagOption, key)
		}
	}
	// prepend bind to options
	return prepend(opts, BindRef(value, set)), nil
}

// isField returns true if v is not a struct or the struct contains at least
// one `ox` tag.
func isField(v reflect.Value) bool {
	typ := v.Type()
	kind := typ.Kind()
	if kind == reflect.Pointer {
		typ = typ.Elem()
		kind = typ.Kind()
	}
	if kind != reflect.Struct {
		return false
	}
	for field := range fields(typ) {
		if _, ok := field.Tag.Lookup(DefaultStructTagName); ok {
			return false
		}
	}
	return true
}

// setField returns the pointer to the bool for name.
func setField(value reflect.Value, name string) (*bool, error) {
	if r := []rune(name); !unicode.IsUpper(r[0]) {
		return nil, fmt.Errorf("%w: set: %q is not an exported field", ErrInvalidType, name)
	}
	refType := value.Type()
	for i := range refType.NumField() {
		switch f := refType.Field(i); {
		case f.Name != name:
			continue
		case f.Type.Kind() != reflect.Bool:
			return nil, fmt.Errorf("%w: set: field %q is not a bool", ErrInvalidType, name)
		}
		return value.Field(i).Addr().Interface().(*bool), nil
	}
	return nil, fmt.Errorf("%w: set: field %q was not found", ErrInvalidType, name)
}

// newCommandError creates a command error.
func newCommandError(name string, err error) error {
	return fmt.Errorf("command %s: %w", name, err)
}

// newFlagError creates a flag error. name is the flag as the command line
// gave it, such as --d or -d.
func newFlagError(name string, err error) error {
	return fmt.Errorf("%s: %w", name, err)
}

// SplitBy splits a string by cut, skipping runes escaped with `\`.
func SplitBy(str string, cut rune) []string {
	r, v := []rune(str), make([][]rune, 1)
	var c, next rune
	for i, j, n := 0, 0, len(r); i < n; i++ {
		c = r[i]
		switch next = peek(r, i+1, n); {
		case r[i] == '\\' && next != 0:
			i, c = i+1, next
		case r[i] == cut:
			j, v = j+1, append(v, nil)
			continue
		}
		v[j] = append(v[j], c)
	}
	s := make([]string, len(v))
	for i, r := range v {
		s[i] = string(r)
	}
	return s
}

// allowed formats a flag's allowed values.
func allowed(s string, allowed []any) string {
	var sb strings.Builder
	for i, v := range allowed {
		if i != 0 {
			_, _ = sb.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&sb, s, v)
	}
	return sb.String()
}

// peek peeks at the next item in v.
func peek[S []T, T any](v S, i, n int) T {
	if i < n {
		return v[i]
	}
	var z T
	return z
}

// prev returns the n-1 element in s.
func prev(s []string, n int) string {
	if 0 < n {
		return s[n-1]
	}
	return ""
}

// maxDist gets the maximum distance from the command's help, or returns
// [DefaultMaxDist].
func maxDist(cmd *Command) int {
	if help, ok := cmd.Help.(*CommandHelp); ok && help.MaxDist != 0 {
		return max(help.MaxDist, 0)
	}
	return DefaultMaxDist
}

// foldCase returns s in lower case when [DefaultFoldCase] is true.
func foldCase(s string) string {
	if DefaultFoldCase {
		return strings.ToLower(s)
	}
	return s
}

// fields is an iterator for reflect struct fields. Cribbed from the go1.26.1
// source tree.
//
// TODO: remove this after tinygo supports go1.26+, and change fields() call to
// TODO: t.Fields().
func fields(t reflect.Type) iter.Seq[reflect.StructField] {
	if t.Kind() != reflect.Struct {
		panic("reflect: Fields of non-struct type " + t.String())
	}
	return func(yield func(reflect.StructField) bool) {
		for i := range t.NumField() {
			if !yield(t.Field(i)) {
				return
			}
		}
	}
}
