// Package text contains the text strings for xo/ox.
package text

var (
	// ErrorMessage is the 'error: ' message.
	ErrorMessage = "error: %v\n"
	// SuggestionErrorMessage is the suggestion error message.
	SuggestionErrorMessage = `%s %q for %q`
	// SuggestionErrorDetails is the suggestion error details.
	SuggestionErrorDetails = "\nDid you mean this?\n  %s\n\n"
	// CommandFlagsSpec is the spec for command flags.
	CommandFlagsSpec = `[flags]`
	// CommandSubSpec is the spec for sub command names.
	CommandSubSpec = `[command]`
	// CommandArgsSpec is the spec for command args.
	CommandArgsSpec = `[args]`
	// FlagSpecSpacer is the spacer between the flag and the displayed flag
	// spec.
	FlagSpecSpacer = ` `
	// FlagDefault is the flag `(default: ...)` text.
	FlagDefault = `(default: %s)`

	// Usage is the `Usage:` section name.
	Usage = `Usage`
	// Aliases is the `Aliases:` section name.
	Aliases = `Aliases`
	// Examples is the `Examples:` section name.
	Examples = `Examples`
	// Commands is the `Available Commands:` section name.
	Commands = `Available Commands`
	// GlobalFlags is the `Global Flags:` section name.
	GlobalFlags = `Global Flags`
	// Flags is the `Flags:` section name.
	Flags = `Flags`

	// VersionCommandName is the `version` command name.
	VersionCommandName = `version`
	// VersionCommandUsage is the `version` command description.
	VersionCommandUsage = `show %s version information`
	// VersionCommandBanner is the `version` command banner.
	VersionCommandBanner = `Show %s version information.`
	// VersionFlagName is the `--version` flag name.
	VersionFlagName = `version`
	// VersionFlagUsage is the `--version` flag description.
	VersionFlagUsage = `show version, then exit`
	// VersionFlagShort is the `--version` short flag name.
	VersionFlagShort = `v`

	// HelpCommandName is the `help` command name.
	HelpCommandName = `help`
	// HelpCommandUsage is the `help` command description.
	HelpCommandUsage = `show help for any command`
	// HelpCommandBanner is the `help` command banner.
	HelpCommandBanner = "Help provides help for any %[1]s command.\n\nSimply type %[1]s help [path to command] for full details."
	// HelpFlagName is the `--help` flag name.
	HelpFlagName = `help`
	// HelpFlagUsage is the `--help` flag description.
	HelpFlagUsage = `show help, then exit`
	// HelpFlagShort is the `--help` short flag name.
	HelpFlagShort = `h`

	// CompCommandName is the `completion` command name.
	CompCommandName = `completion`
	// CompCommandUsage is the `completion` command description.
	CompCommandUsage = `generate completion script for %s`
	// CompCommandBanner is the `completion` command banner.
	CompCommandBanner = "Generate %s completion script for %s.\n\nSee each sub-command's help for details on using the generated completion script."
	// CompFlagName is the `--completion-script-<type>` flag name.
	CompFlagName = `completion-script-%s`
	// CompFlagUsage is the `--completion-script-<type>` flag description.
	CompFlagUsage = `generate completion script for %s`
	// CompCommandFlagNoDescriptionsName is the `completion` command
	// `no-descriptions` flag name.
	CompCommandFlagNoDescriptionsName = `no-descriptions`
	// CompCommandFlagNoDescriptionsUsage is the `completion` command
	// `no-descriptions` flag description.
	CompCommandFlagNoDescriptionsUsage = `disable completion descriptions`
	// CompCommandAnyShellDesc is the `completion` command `a specified shell`
	// description.
	CompCommandAnyShellDesc = `a specified shell`

	// Footer is the default footer for commands with sub commands.
	Footer = `Use "%s [command] --help" for more information about a command.`
)
