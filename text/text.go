// Package text contains the text strings for xo/ox.
package text

var (
	// ErrorMessage is the 'error: ' message.
	ErrorMessage = "error: %v\n"
	// SuggestionErrorMessage is the suggestion error message.
	SuggestionErrorMessage = `%s %q for %q`
	// SuggestionErrorDetails is the suggestion error details.
	SuggestionErrorDetails = "\nDid you mean this?\n  %s\n\n"
	// Usage is the `Usage:` section name.
	Usage = `Usage`
	// Aliases is the `Aliases:` section name.
	Aliases = `Aliases`
	// Commands is the `Available Commands` section name.
	Commands = `Available Commands`
	// Flags is the `Flags:` section name.
	Flags = `Flags`
	// GlobalFlags is the `Global Flags:` section name.
	GlobalFlags = `Global Flags`
	// VersionCommandName is the version command name.
	VersionCommandName = `version`
	// VersionCommandDesc is the version command description.
	VersionCommandDesc = `show version for any command`
	// VersionFlagName is the version flag name.
	VersionFlagName = `version`
	// VersionFlagDesc is the version flag description.
	VersionFlagDesc = `show version, then exit`
	// VersionFlagShort is the version short flag name.
	VersionFlagShort = `v`
	// HelpFlagName is the help command/flag name.
	HelpFlagName = `help`
	// HelpFlagDesc is the help flag description.
	HelpFlagDesc = `show help, then exit`
	// HelpFlagShort is the help short flag name.
	HelpFlagShort = `h`
	// HelpCommandName is the help command name.
	HelpCommandName = `help`
	// HelpCommandDesc is the help command description.
	HelpCommandDesc = `show help for any command`
	// CommandFlagsSpec is the default spec for command flags.
	CommandFlagsSpec = `[flags]`
	// CommandSubSpec is the default spec for sub command names.
	CommandSubSpec = `[command]`
	// CommandArgsSpec is the default spec for command args.
	CommandArgsSpec = `[args]`
	// Footer is the default footer for commands with sub commands.
	Footer = `Use "%s [command] --help" for more information about a command.`
	// FlagSpecSpacer is the spacer between the flag and the displayed flag
	// spec.
	FlagSpecSpacer = ` `
	// FlagDefault is the default flag text.
	FlagDefault = `(default: %s)`
	// CompFlagName is the default completion flag name.
	CompFlagName = `completion-script-%s`
	// CompCommandName is the default completion command name.
	CompCommandName = `completion`
	// CompDesc is the default completion command description.
	CompDesc = `generate completion script for %s`
	// CompCommandBanner is the defalut completion command banner.
	CompCommandBanner = `Generate %s completion script for %s.`
	// CompCommandAShellDesc is the default `a shell` completion command
	// description.
	CompCommandAShellDesc = `a shell`
)
