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
	// Flags is the `Flags:` section name.
	Flags = `Flags`
	// GlobalFlags is the `Global Flags:` section name.
	GlobalFlags = `Global Flags`

	// VersionCommandName is the `version` command name.
	VersionCommandName = `version`
	// VersionCommandDesc is the `version` command description.
	VersionCommandDesc = `show version for any command`
	// VersionCommandBanner is the `version` command banner.
	VersionCommandBanner = `Shows the %s version and exits.`
	// VersionFlagName is the `--version` flag name.
	VersionFlagName = `version`
	// VersionFlagDesc is the `--version` flag description.
	VersionFlagDesc = `show version, then exit`
	// VersionFlagShort is the `--version` short flag name.
	VersionFlagShort = `v`

	// HelpCommandName is the `help` command name.
	HelpCommandName = `help`
	// HelpCommandDesc is the `help` command description.
	HelpCommandDesc = `show help for any command`
	// HelpCommandBanner is the `help` command banner.
	HelpCommandBanner = "Help provides help for any %[1]s command.\n\nSimply type %[1]s help [path to command] for full details."
	// HelpFlagName is the `--help` flag name.
	HelpFlagName = `help`
	// HelpFlagDesc is the `--help` flag description.
	HelpFlagDesc = `show help, then exit`
	// HelpFlagShort is the `--help` short flag name.
	HelpFlagShort = `h`

	// CompCommandName is the `completion` command name.
	CompCommandName = `completion`
	// CompCommandBanner is the `completion` command banner.
	CompCommandBanner = "Generate %s completion script for %s.\n\nSee each sub-command's help for details on using the generated completion script."
	// CompCommandSubBanner is the `completion <type>` command banner.
	CompCommandSubBanner = `Generate %s completion script for %s.`
	// CompFlagName is the `--completion-script-<type>` flag name.
	CompFlagName = `completion-script-%s`
	// CompFlagDesc is the `--completion-script-<type>` flag description.
	CompFlagDesc = `generate completion script for %s`
	// CompCommandFlagNoDescriptionsName is the `completion` command
	// `no-descriptions` flag name.
	CompCommandFlagNoDescriptionsName = `no-descriptions`
	// CompCommandFlagNoDescriptionsDesc is the `completion` command
	// `no-descriptions` flag description.
	CompCommandFlagNoDescriptionsDesc = `disable completion descriptions`
	// CompCommandAnyShellDesc is the default `a shell` completion command
	// description.
	CompCommandAnyShellDesc = `a specified shell`

	// Footer is the default footer for commands with sub commands.
	Footer = `Use "%s [command] --help" for more information about a command.`
)
