// Command rclone is a xo/ox version of `rclone`.
//
// Generated from _examples/gen.go.
package main

import (
	"context"

	"github.com/xo/ox"
)

func main() {
	ox.RunContext(
		context.Background(),
		ox.Defaults(),
		ox.Usage(`rclone`, ``),
		ox.Banner(`
Rclone syncs files to and from cloud storage providers as well as
mounting them, listing them in lots of different ways.

See the home page (https://rclone.org/) for installation, usage,
documentation, changelog and configuration walkthroughs.`),
		ox.Spec(`[flags]`),
		ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		ox.Sub(
			ox.Usage(`authorize`, `Remote authorization.`),
			ox.Banner(`
Remote authorization. Used to authorize a remote or headless
rclone from a machine with a browser - use as instructed by
rclone config.

Use the --auth-no-open-browser to prevent rclone to open auth
link in default browser automatically.`),
			ox.Spec(`[flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`auth-no-open-browser`, `Do not automatically open auth link in default browser`),
		),
		ox.Sub(
			ox.Usage(`backend`, `Run a backend-specific command.`),
			ox.Banner(`
This runs a backend-specific command. The commands themselves (except
for "help" and "features") are defined by the backends and you should
see the backend docs for definitions.

You can discover what commands a backend implements by using

    rclone backend help remote:
    rclone backend help <backendname>

You can also discover information about the backend using (see
[operations/fsinfo](/rc/#operations-fsinfo) in the remote control docs
for more info).

    rclone backend features remote:

Pass options to the backend command with -o. This should be key=value or key, e.g.:

    rclone backend stats remote:path stats -o format=json -o long

Pass arguments to the backend by placing them on the end of the line

    rclone backend cleanup remote:path file1 file2 file3

Note to run these commands on a running backend then see
[backend/command](/rc/#backend-command) in the rc docs.`),
			ox.Spec(`<command> remote:path [opts] <args> [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`json`, `Always output in JSON format`).
				Array(`option`, `Option in the form name=value or name`, ox.Short("o")),
		),
		ox.Sub(
			ox.Usage(`bisync`, `Perform bidirectonal synchronization between two paths.`),
			ox.Banner(`Perform bidirectonal synchronization between two paths.

[Bisync](https://rclone.org/bisync/) provides a
bidirectional cloud sync solution in rclone.
It retains the Path1 and Path2 filesystem listings from the prior run.
On each successive run it will:
- list files on Path1 and Path2, and check for changes on each side.
  Changes include `+"`"+`New`+"`"+`, `+"`"+`Newer`+"`"+`, `+"`"+`Older`+"`"+`, and `+"`"+`Deleted`+"`"+` files.
- Propagate changes on Path1 to Path2, and vice-versa.

See [full bisync description](https://rclone.org/bisync/) for details.`),
			ox.Spec(`remote1:path1 remote2:path2 [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`check-access`, `Ensure expected RCLONE_TEST files are found on both Path1 and Path2 filesystems, else abort.`).
				String(`check-filename`, `Filename for --check-access (default: RCLONE_TEST)`).
				String(`check-sync`, `Controls comparison of final listings: true|false|only (default: true)`, ox.Default("true")).
				String(`filters-file`, `Read filtering patterns from a file`).
				Bool(`force`, `Bypass --max-delete safety check and run the sync. Consider using with --verbose`).
				Bool(`localtime`, `Use local time in listings (default: UTC)`).
				Bool(`no-cleanup`, `Retain working files (useful for troubleshooting and testing).`).
				Bool(`remove-empty-dirs`, `Remove empty directories at the final cleanup step.`).
				Bool(`resync`, `Performs the resync run. Path1 files may overwrite Path2 versions. Consider using --verbose or --dry-run first.`, ox.Short("1")).
				String(`workdir`, `Use custom working dir - useful for testing. (default: /home/ken/.cache/rclone/bisync)`),
		),
		ox.Sub(
			ox.Usage(`cat`, `Concatenates any files and sends them to stdout.`),
			ox.Banner(`
rclone cat sends any files to standard output.

You can use it like this to output a single file

    rclone cat remote:path/to/file

Or like this to output any file in dir or its subdirectories.

    rclone cat remote:path/to/dir

Or like this to output any .txt files in dir or its subdirectories.

    rclone --include "*.txt" cat remote:path/to/dir

Use the `+"`"+`--head`+"`"+` flag to print characters only at the start, `+"`"+`--tail`+"`"+` for
the end and `+"`"+`--offset`+"`"+` and `+"`"+`--count`+"`"+` to print a section in the middle.
Note that if offset is negative it will count from the end, so
`+"`"+`--offset -1 --count 1`+"`"+` is equivalent to `+"`"+`--tail 1`+"`"+`.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Int(`count`, `Only print N characters`, ox.Default("-1")).
				Bool(`discard`, `Discard the output instead of printing`).
				Int(`head`, `Only print the first N characters`).
				Int(`offset`, `Start printing at offset N (or from end if -ve)`).
				Int(`tail`, `Only print the last N characters`),
		),
		ox.Sub(
			ox.Usage(`check`, `Checks the files in the source and destination match.`),
			ox.Banner(`
Checks the files in the source and destination match.  It compares
sizes and hashes (MD5 or SHA1) and logs a report of files that don't
match.  It doesn't alter the source or destination.

If you supply the `+"`"+`--size-only`+"`"+` flag, it will only compare the sizes not
the hashes as well.  Use this for a quick check.

If you supply the `+"`"+`--download`+"`"+` flag, it will download the data from
both remotes and check them against each other on the fly.  This can
be useful for remotes that don't support hashes or if you really want
to check all the data.

If you supply the `+"`"+`--checkfile HASH`+"`"+` flag with a valid hash name,
the `+"`"+`source:path`+"`"+` must point to a text file in the SUM format.

If you supply the `+"`"+`--one-way`+"`"+` flag, it will only check that files in
the source match the files in the destination, not the other way
around. This means that extra files in the destination that are not in
the source will not be detected.

The `+"`"+`--differ`+"`"+`, `+"`"+`--missing-on-dst`+"`"+`, `+"`"+`--missing-on-src`+"`"+`, `+"`"+`--match`+"`"+`
and `+"`"+`--error`+"`"+` flags write paths, one per line, to the file name (or
stdout if it is `+"`"+`-`+"`"+`) supplied. What they write is described in the
help below. For example `+"`"+`--differ`+"`"+` will write all paths which are
present on both the source and destination but different.

The `+"`"+`--combined`+"`"+` flag will write a file (or stdout) which contains all
file paths with a symbol and then a space and then the path to tell
you what happened to it. These are reminiscent of diff files.

- `+"`"+`= path`+"`"+` means path was found in source and destination and was identical
- `+"`"+`- path`+"`"+` means path was missing on the source, so only in the destination
- `+"`"+`+ path`+"`"+` means path was missing on the destination, so only in the source
- `+"`"+`* path`+"`"+` means path was present in source and destination but different.
- `+"`"+`! path`+"`"+` means there was an error reading or hashing the source or dest.`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				String(`checkfile`, `Treat source:path as a SUM file with hashes of given type`, ox.Short("C")).
				String(`combined`, `Make a combined report of changes to this file`).
				String(`differ`, `Report all non-matching files to this file`).
				Bool(`download`, `Check by downloading rather than with hash`).
				String(`error`, `Report all files with errors (hashing or reading) to this file`).
				String(`match`, `Report all matching files to this file`).
				String(`missing-on-dst`, `Report all files missing from the destination to this file`).
				String(`missing-on-src`, `Report all files missing from the source to this file`).
				Bool(`one-way`, `Check one way only, source files must exist on remote`),
		),
		ox.Sub(
			ox.Usage(`checksum`, `Checks the files in the source against a SUM file.`),
			ox.Banner(`
Checks that hashsums of source files match the SUM file.
It compares hashes (MD5, SHA1, etc) and logs a report of files which
don't match.  It doesn't alter the file system.

If you supply the `+"`"+`--download`+"`"+` flag, it will download the data from remote
and calculate the contents hash on the fly.  This can be useful for remotes
that don't support hashes or if you really want to check all the data.

Note that hash values in the SUM file are treated as case insensitive.

If you supply the `+"`"+`--one-way`+"`"+` flag, it will only check that files in
the source match the files in the destination, not the other way
around. This means that extra files in the destination that are not in
the source will not be detected.

The `+"`"+`--differ`+"`"+`, `+"`"+`--missing-on-dst`+"`"+`, `+"`"+`--missing-on-src`+"`"+`, `+"`"+`--match`+"`"+`
and `+"`"+`--error`+"`"+` flags write paths, one per line, to the file name (or
stdout if it is `+"`"+`-`+"`"+`) supplied. What they write is described in the
help below. For example `+"`"+`--differ`+"`"+` will write all paths which are
present on both the source and destination but different.

The `+"`"+`--combined`+"`"+` flag will write a file (or stdout) which contains all
file paths with a symbol and then a space and then the path to tell
you what happened to it. These are reminiscent of diff files.

- `+"`"+`= path`+"`"+` means path was found in source and destination and was identical
- `+"`"+`- path`+"`"+` means path was missing on the source, so only in the destination
- `+"`"+`+ path`+"`"+` means path was missing on the destination, so only in the source
- `+"`"+`* path`+"`"+` means path was present in source and destination but different.
- `+"`"+`! path`+"`"+` means there was an error reading or hashing the source or dest.`),
			ox.Spec(`<hash> sumfile src:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				String(`combined`, `Make a combined report of changes to this file`).
				String(`differ`, `Report all non-matching files to this file`).
				Bool(`download`, `Check by hashing the contents`).
				String(`error`, `Report all files with errors (hashing or reading) to this file`).
				String(`match`, `Report all matching files to this file`).
				String(`missing-on-dst`, `Report all files missing from the destination to this file`).
				String(`missing-on-src`, `Report all files missing from the source to this file`).
				Bool(`one-way`, `Check one way only, source files must exist on remote`),
		),
		ox.Sub(
			ox.Usage(`cleanup`, `Clean up the remote if possible.`),
			ox.Banner(`
Clean up the remote if possible.  Empty the trash or delete old file
versions. Not supported by all remotes.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`config`, `Enter an interactive configuration session.`),
			ox.Banner(`Enter an interactive configuration session where you can setup new
remotes and manage existing ones. You may also set or remove a
password to protect your configuration.`),
			ox.Spec(`[flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Sub(
				ox.Usage(`create`, `Create a new remote with name, type and options.`),
				ox.Banner(`
Create a new remote of `+"`"+`name`+"`"+` with `+"`"+`type`+"`"+` and options.  The options
should be passed in pairs of `+"`"+`key`+"`"+` `+"`"+`value`+"`"+` or as `+"`"+`key=value`+"`"+`.

For example, to make a swift remote of name myremote using auto config
you would do:

    rclone config create myremote swift env_auth true
    rclone config create myremote swift env_auth=true

So for example if you wanted to configure a Google Drive remote but
using remote authorization you would do this:

    rclone config create mydrive drive config_is_local=false

Note that if the config process would normally ask a question the
default is taken (unless `+"`"+`--non-interactive`+"`"+` is used).  Each time
that happens rclone will print or DEBUG a message saying how to
affect the value taken.

If any of the parameters passed is a password field, then rclone will
automatically obscure them if they aren't already obscured before
putting them in the config file.

**NB** If the password parameter is 22 characters or longer and
consists only of base64 characters then rclone can get confused about
whether the password is already obscured or not and put unobscured
passwords into the config file. If you want to be 100% certain that
the passwords get obscured then use the `+"`"+`--obscure`+"`"+` flag, or if you
are 100% certain you are already passing obscured passwords then use
`+"`"+`--no-obscure`+"`"+`.  You can also set obscured passwords using the
`+"`"+`rclone config password`+"`"+` command.

The flag `+"`"+`--non-interactive`+"`"+` is for use by applications that wish to
configure rclone themeselves, rather than using rclone's text based
configuration questions. If this flag is set, and rclone needs to ask
the user a question, a JSON blob will be returned with the question in
it.

This will look something like (some irrelevant detail removed):

`+"`"+``+"`"+``+"`"+`
{
    "State": "*oauth-islocal,teamdrive,,",
    "Option": {
        "Name": "config_is_local",
        "Help": "Use auto config?\n * Say Y if not sure\n * Say N if you are working on a remote or headless machine\n",
        "Default": true,
        "Examples": [
            {
                "Value": "true",
                "Help": "Yes"
            },
            {
                "Value": "false",
                "Help": "No"
            }
        ],
        "Required": false,
        "IsPassword": false,
        "Type": "bool",
        "Exclusive": true,
    },
    "Error": "",
}
`+"`"+``+"`"+``+"`"+`

The format of `+"`"+`Option`+"`"+` is the same as returned by `+"`"+`rclone config
providers`+"`"+`. The question should be asked to the user and returned to
rclone as the `+"`"+`--result`+"`"+` option along with the `+"`"+`--state`+"`"+` parameter.

The keys of `+"`"+`Option`+"`"+` are used as follows:

- `+"`"+`Name`+"`"+` - name of variable - show to user
- `+"`"+`Help`+"`"+` - help text. Hard wrapped at 80 chars. Any URLs should be clicky.
- `+"`"+`Default`+"`"+` - default value - return this if the user just wants the default.
- `+"`"+`Examples`+"`"+` - the user should be able to choose one of these
- `+"`"+`Required`+"`"+` - the value should be non-empty
- `+"`"+`IsPassword`+"`"+` - the value is a password and should be edited as such
- `+"`"+`Type`+"`"+` - type of value, eg `+"`"+`bool`+"`"+`, `+"`"+`string`+"`"+`, `+"`"+`int`+"`"+` and others
- `+"`"+`Exclusive`+"`"+` - if set no free-form entry allowed only the `+"`"+`Examples`+"`"+`
- Irrelevant keys `+"`"+`Provider`+"`"+`, `+"`"+`ShortOpt`+"`"+`, `+"`"+`Hide`+"`"+`, `+"`"+`NoPrefix`+"`"+`, `+"`"+`Advanced`+"`"+`

If `+"`"+`Error`+"`"+` is set then it should be shown to the user at the same
time as the question.

    rclone config update name --continue --state "*oauth-islocal,teamdrive,," --result "true"

Note that when using `+"`"+`--continue`+"`"+` all passwords should be passed in
the clear (not obscured). Any default config values should be passed
in with each invocation of `+"`"+`--continue`+"`"+`.

At the end of the non interactive process, rclone will return a result
with `+"`"+`State`+"`"+` as empty string.

If `+"`"+`--all`+"`"+` is passed then rclone will ask all the config questions,
not just the post config questions. Any parameters are used as
defaults for questions as usual.

Note that `+"`"+`bin/config.py`+"`"+` in the rclone source implements this protocol
as a readable demonstration.`),
				ox.Spec(`name type [key value]* [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Bool(`all`, `Ask the full set of config questions`).
					Bool(`continue`, `Continue the configuration process with an answer`).
					Bool(`no-obscure`, `Force any passwords not to be obscured`).
					Bool(`non-interactive`, `Don't interact with user and return questions`).
					Bool(`obscure`, `Force any passwords to be obscured`).
					String(`result`, `Result - use with --continue`).
					String(`state`, `State - use with --continue`),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete an existing remote.`),
				ox.Banner(`Delete an existing remote.`),
				ox.Spec(`name [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`disconnect`, `Disconnects user from remote`),
				ox.Banner(`
This disconnects the remote: passed in to the cloud storage system.

This normally means revoking the oauth token.

To reconnect use "rclone config reconnect".`),
				ox.Spec(`remote: [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`dump`, `Dump the config file as JSON.`),
				ox.Banner(`Dump the config file as JSON.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`file`, `Show path of configuration file in use.`),
				ox.Banner(`Show path of configuration file in use.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`password`, `Update password in an existing remote.`),
				ox.Banner(`
Update an existing remote's password. The password
should be passed in pairs of `+"`"+`key`+"`"+` `+"`"+`password`+"`"+` or as `+"`"+`key=password`+"`"+`.
The `+"`"+`password`+"`"+` should be passed in in clear (unobscured).

For example, to set password of a remote of name myremote you would do:

    rclone config password myremote fieldname mypassword
    rclone config password myremote fieldname=mypassword

This command is obsolete now that "config update" and "config create"
both support obscuring passwords directly.`),
				ox.Spec(`name [key value]+ [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`paths`, `Show paths used for configuration, cache, temp etc.`),
				ox.Banner(`Show paths used for configuration, cache, temp etc.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`providers`, `List in JSON format all the providers and options.`),
				ox.Banner(`List in JSON format all the providers and options.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`reconnect`, `Re-authenticates user with remote.`),
				ox.Banner(`
This reconnects remote: passed in to the cloud storage system.

To disconnect the remote use "rclone config disconnect".

This normally means going through the interactive oauth flow again.`),
				ox.Spec(`remote: [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`show`, `Print (decrypted) config file, or the config for a single remote.`),
				ox.Banner(`Print (decrypted) config file, or the config for a single remote.`),
				ox.Spec(`[<remote>] [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`touch`, `Ensure configuration file exists.`),
				ox.Banner(`Ensure configuration file exists.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`update`, `Update options in an existing remote.`),
				ox.Banner(`
Update an existing remote's options. The options should be passed in
pairs of `+"`"+`key`+"`"+` `+"`"+`value`+"`"+` or as `+"`"+`key=value`+"`"+`.

For example, to update the env_auth field of a remote of name myremote
you would do:

    rclone config update myremote env_auth true
    rclone config update myremote env_auth=true

If the remote uses OAuth the token will be updated, if you don't
require this add an extra parameter thus:

    rclone config update myremote env_auth=true config_refresh_token=false

Note that if the config process would normally ask a question the
default is taken (unless `+"`"+`--non-interactive`+"`"+` is used).  Each time
that happens rclone will print or DEBUG a message saying how to
affect the value taken.

If any of the parameters passed is a password field, then rclone will
automatically obscure them if they aren't already obscured before
putting them in the config file.

**NB** If the password parameter is 22 characters or longer and
consists only of base64 characters then rclone can get confused about
whether the password is already obscured or not and put unobscured
passwords into the config file. If you want to be 100% certain that
the passwords get obscured then use the `+"`"+`--obscure`+"`"+` flag, or if you
are 100% certain you are already passing obscured passwords then use
`+"`"+`--no-obscure`+"`"+`.  You can also set obscured passwords using the
`+"`"+`rclone config password`+"`"+` command.

The flag `+"`"+`--non-interactive`+"`"+` is for use by applications that wish to
configure rclone themeselves, rather than using rclone's text based
configuration questions. If this flag is set, and rclone needs to ask
the user a question, a JSON blob will be returned with the question in
it.

This will look something like (some irrelevant detail removed):

`+"`"+``+"`"+``+"`"+`
{
    "State": "*oauth-islocal,teamdrive,,",
    "Option": {
        "Name": "config_is_local",
        "Help": "Use auto config?\n * Say Y if not sure\n * Say N if you are working on a remote or headless machine\n",
        "Default": true,
        "Examples": [
            {
                "Value": "true",
                "Help": "Yes"
            },
            {
                "Value": "false",
                "Help": "No"
            }
        ],
        "Required": false,
        "IsPassword": false,
        "Type": "bool",
        "Exclusive": true,
    },
    "Error": "",
}
`+"`"+``+"`"+``+"`"+`

The format of `+"`"+`Option`+"`"+` is the same as returned by `+"`"+`rclone config
providers`+"`"+`. The question should be asked to the user and returned to
rclone as the `+"`"+`--result`+"`"+` option along with the `+"`"+`--state`+"`"+` parameter.

The keys of `+"`"+`Option`+"`"+` are used as follows:

- `+"`"+`Name`+"`"+` - name of variable - show to user
- `+"`"+`Help`+"`"+` - help text. Hard wrapped at 80 chars. Any URLs should be clicky.
- `+"`"+`Default`+"`"+` - default value - return this if the user just wants the default.
- `+"`"+`Examples`+"`"+` - the user should be able to choose one of these
- `+"`"+`Required`+"`"+` - the value should be non-empty
- `+"`"+`IsPassword`+"`"+` - the value is a password and should be edited as such
- `+"`"+`Type`+"`"+` - type of value, eg `+"`"+`bool`+"`"+`, `+"`"+`string`+"`"+`, `+"`"+`int`+"`"+` and others
- `+"`"+`Exclusive`+"`"+` - if set no free-form entry allowed only the `+"`"+`Examples`+"`"+`
- Irrelevant keys `+"`"+`Provider`+"`"+`, `+"`"+`ShortOpt`+"`"+`, `+"`"+`Hide`+"`"+`, `+"`"+`NoPrefix`+"`"+`, `+"`"+`Advanced`+"`"+`

If `+"`"+`Error`+"`"+` is set then it should be shown to the user at the same
time as the question.

    rclone config update name --continue --state "*oauth-islocal,teamdrive,," --result "true"

Note that when using `+"`"+`--continue`+"`"+` all passwords should be passed in
the clear (not obscured). Any default config values should be passed
in with each invocation of `+"`"+`--continue`+"`"+`.

At the end of the non interactive process, rclone will return a result
with `+"`"+`State`+"`"+` as empty string.

If `+"`"+`--all`+"`"+` is passed then rclone will ask all the config questions,
not just the post config questions. Any parameters are used as
defaults for questions as usual.

Note that `+"`"+`bin/config.py`+"`"+` in the rclone source implements this protocol
as a readable demonstration.`),
				ox.Spec(`name [key value]+ [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Bool(`all`, `Ask the full set of config questions`).
					Bool(`continue`, `Continue the configuration process with an answer`).
					Bool(`no-obscure`, `Force any passwords not to be obscured`).
					Bool(`non-interactive`, `Don't interact with user and return questions`).
					Bool(`obscure`, `Force any passwords to be obscured`).
					String(`result`, `Result - use with --continue`).
					String(`state`, `State - use with --continue`),
			),
			ox.Sub(
				ox.Usage(`userinfo`, `Prints info about logged in user of remote.`),
				ox.Banner(`
This prints the details of the person logged in to the cloud storage
system.`),
				ox.Spec(`remote: [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Bool(`json`, `Format output as JSON`),
			),
		),
		ox.Sub(
			ox.Usage(`copy`, `Copy files from source to dest, skipping identical files.`),
			ox.Banner(`
Copy the source to the destination.  Does not transfer files that are
identical on source and destination, testing by size and modification
time or MD5SUM.  Doesn't delete files from the destination.

Note that it is always the contents of the directory that is synced,
not the directory so when source:path is a directory, it's the
contents of source:path that are copied, not the directory name and
contents.

If dest:path doesn't exist, it is created and the source:path contents
go there.

For example

    rclone copy source:sourcepath dest:destpath

Let's say there are two files in sourcepath

    sourcepath/one.txt
    sourcepath/two.txt

This copies them to

    destpath/one.txt
    destpath/two.txt

Not to

    destpath/sourcepath/one.txt
    destpath/sourcepath/two.txt

If you are familiar with `+"`"+`rsync`+"`"+`, rclone always works as if you had
written a trailing `+"`"+`/`+"`"+` - meaning "copy the contents of this directory".
This applies to all commands and whether you are talking about the
source or destination.

See the [--no-traverse](/docs/#no-traverse) option for controlling
whether rclone lists the destination directory or not.  Supplying this
option when copying a small number of files into a large destination
can speed transfers up greatly.

For example, if you have many files in /path/to/src but only a few of
them change every day, you can copy all the files which have changed
recently very efficiently like this:

    rclone copy --max-age 24h --no-traverse /path/to/src remote:

**Note**: Use the `+"`"+`-P`+"`"+`/`+"`"+`--progress`+"`"+` flag to view real-time transfer statistics.

**Note**: Use the `+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag to test without copying anything.`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`create-empty-src-dirs`, `Create empty source dirs on destination after copy`),
		),
		ox.Sub(
			ox.Usage(`copyto`, `Copy files from source to dest, skipping identical files.`),
			ox.Banner(`
If source:path is a file or directory then it copies it to a file or
directory named dest:path.

This can be used to upload single files to other than their current
name.  If the source is a directory then it acts exactly like the copy
command.

So

    rclone copyto src dst

where src and dst are rclone paths, either remote:path or
/path/to/local or C:\windows\path\if\on\windows.

This will:

    if src is file
        copy it to dst, overwriting an existing file if it exists
    if src is directory
        copy it to dst, overwriting existing files if they exist
        see copy command for full details

This doesn't transfer files that are identical on src and dst, testing
by size and modification time or MD5SUM.  It doesn't delete files from
the destination.

**Note**: Use the `+"`"+`-P`+"`"+`/`+"`"+`--progress`+"`"+` flag to view real-time transfer statistics`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`copyurl`, `Copy url content to dest.`),
			ox.Banner(`
Download a URL's content and copy it to the destination without saving
it in temporary storage.

Setting `+"`"+`--auto-filename`+"`"+` will cause the file name to be retrieved from
the URL (after any redirections) and used in the destination
path. With `+"`"+`--print-filename`+"`"+` in addition, the resulting file name will
be printed.

Setting `+"`"+`--no-clobber`+"`"+` will prevent overwriting file on the 
destination if there is one with the same name.

Setting `+"`"+`--stdout`+"`"+` or making the output file name `+"`"+`-`+"`"+`
will cause the output to be written to standard output.`),
			ox.Spec(`https://example.com dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`auto-filename`, `Get the file name from the URL and use it for destination file path`, ox.Short("a")).
				Bool(`no-clobber`, `Prevent overwriting file with same name`).
				Bool(`print-filename`, `Print the resulting name from --auto-filename`, ox.Short("p")).
				Bool(`stdout`, `Write the output to stdout rather than a file`),
		),
		ox.Sub(
			ox.Usage(`cryptcheck`, `Cryptcheck checks the integrity of a crypted remote.`),
			ox.Banner(`
rclone cryptcheck checks a remote against a crypted remote.  This is
the equivalent of running rclone check, but able to check the
checksums of the crypted remote.

For it to work the underlying remote of the cryptedremote must support
some kind of checksum.

It works by reading the nonce from each file on the cryptedremote: and
using that to encrypt each file on the remote:.  It then checks the
checksum of the underlying file on the cryptedremote: against the
checksum of the file it has just encrypted.

Use it like this

    rclone cryptcheck /path/to/files encryptedremote:path

You can use it like this also, but that will involve downloading all
the files in remote:path.

    rclone cryptcheck remote:path encryptedremote:path

After it has run it will log the status of the encryptedremote:.

If you supply the `+"`"+`--one-way`+"`"+` flag, it will only check that files in
the source match the files in the destination, not the other way
around. This means that extra files in the destination that are not in
the source will not be detected.

The `+"`"+`--differ`+"`"+`, `+"`"+`--missing-on-dst`+"`"+`, `+"`"+`--missing-on-src`+"`"+`, `+"`"+`--match`+"`"+`
and `+"`"+`--error`+"`"+` flags write paths, one per line, to the file name (or
stdout if it is `+"`"+`-`+"`"+`) supplied. What they write is described in the
help below. For example `+"`"+`--differ`+"`"+` will write all paths which are
present on both the source and destination but different.

The `+"`"+`--combined`+"`"+` flag will write a file (or stdout) which contains all
file paths with a symbol and then a space and then the path to tell
you what happened to it. These are reminiscent of diff files.

- `+"`"+`= path`+"`"+` means path was found in source and destination and was identical
- `+"`"+`- path`+"`"+` means path was missing on the source, so only in the destination
- `+"`"+`+ path`+"`"+` means path was missing on the destination, so only in the source
- `+"`"+`* path`+"`"+` means path was present in source and destination but different.
- `+"`"+`! path`+"`"+` means there was an error reading or hashing the source or dest.`),
			ox.Spec(`remote:path cryptedremote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				String(`combined`, `Make a combined report of changes to this file`).
				String(`differ`, `Report all non-matching files to this file`).
				String(`error`, `Report all files with errors (hashing or reading) to this file`).
				String(`match`, `Report all matching files to this file`).
				String(`missing-on-dst`, `Report all files missing from the destination to this file`).
				String(`missing-on-src`, `Report all files missing from the source to this file`).
				Bool(`one-way`, `Check one way only, source files must exist on remote`),
		),
		ox.Sub(
			ox.Usage(`cryptdecode`, `Cryptdecode returns unencrypted file names.`),
			ox.Banner(`
rclone cryptdecode returns unencrypted file names when provided with
a list of encrypted file names. List limit is 10 items.

If you supply the --reverse flag, it will return encrypted file names.

use it like this

	rclone cryptdecode encryptedremote: encryptedfilename1 encryptedfilename2

	rclone cryptdecode --reverse encryptedremote: filename1 filename2

Another way to accomplish this is by using the `+"`"+`rclone backend encode`+"`"+` (or `+"`"+`decode`+"`"+`)command.
See the documentation on the `+"`"+`crypt`+"`"+` overlay for more info.`),
			ox.Spec(`encryptedremote: encryptedfilename [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`reverse`, `Reverse cryptdecode, encrypts filenames`),
		),
		ox.Sub(
			ox.Usage(`dedupe`, `Interactively find duplicate filenames and delete/rename them.`),
			ox.Banner(`

By default `+"`"+`dedupe`+"`"+` interactively finds files with duplicate
names and offers to delete all but one or rename them to be
different. This is known as deduping by name.

Deduping by name is only useful with a small group of backends (e.g. Google Drive,
Opendrive) that can have duplicate file names. It can be run on wrapping backends
(e.g. crypt) if they wrap a backend which supports duplicate file
names.

However if --by-hash is passed in then dedupe will find files with
duplicate hashes instead which will work on any backend which supports
at least one hash. This can be used to find files with duplicate
content. This is known as deduping by hash.

If deduping by name, first rclone will merge directories with the same
name.  It will do this iteratively until all the identically named
directories have been merged.

Next, if deduping by name, for every group of duplicate file names /
hashes, it will delete all but one identical file it finds without
confirmation.  This means that for most duplicated files the `+"`"+`dedupe`+"`"+` command will not be interactive.

`+"`"+`dedupe`+"`"+` considers files to be identical if they have the
same file path and the same hash. If the backend does not support hashes (e.g. crypt wrapping
Google Drive) then they will never be found to be identical. If you
use the `+"`"+`--size-only`+"`"+` flag then files will be considered
identical if they have the same size (any hash will be ignored). This
can be useful on crypt backends which do not support hashes.

Next rclone will resolve the remaining duplicates. Exactly which
action is taken depends on the dedupe mode. By default, rclone will
interactively query the user for each one.

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.

Here is an example run.

Before - with duplicates

    $ rclone lsl drive:dupes
      6048320 2016-03-05 16:23:16.798000000 one.txt
      6048320 2016-03-05 16:23:11.775000000 one.txt
       564374 2016-03-05 16:23:06.731000000 one.txt
      6048320 2016-03-05 16:18:26.092000000 one.txt
      6048320 2016-03-05 16:22:46.185000000 two.txt
      1744073 2016-03-05 16:22:38.104000000 two.txt
       564374 2016-03-05 16:22:52.118000000 two.txt

Now the `+"`"+`dedupe`+"`"+` session

    $ rclone dedupe drive:dupes
    2016/03/05 16:24:37 Google drive root 'dupes': Looking for duplicates using interactive mode.
    one.txt: Found 4 files with duplicate names
    one.txt: Deleting 2/3 identical duplicates (MD5 "1eedaa9fe86fd4b8632e2ac549403b36")
    one.txt: 2 duplicates remain
      1:      6048320 bytes, 2016-03-05 16:23:16.798000000, MD5 1eedaa9fe86fd4b8632e2ac549403b36
      2:       564374 bytes, 2016-03-05 16:23:06.731000000, MD5 7594e7dc9fc28f727c42ee3e0749de81
    s) Skip and do nothing
    k) Keep just one (choose which in next step)
    r) Rename all to be different (by changing file.jpg to file-1.jpg)
    s/k/r> k
    Enter the number of the file to keep> 1
    one.txt: Deleted 1 extra copies
    two.txt: Found 3 files with duplicate names
    two.txt: 3 duplicates remain
      1:       564374 bytes, 2016-03-05 16:22:52.118000000, MD5 7594e7dc9fc28f727c42ee3e0749de81
      2:      6048320 bytes, 2016-03-05 16:22:46.185000000, MD5 1eedaa9fe86fd4b8632e2ac549403b36
      3:      1744073 bytes, 2016-03-05 16:22:38.104000000, MD5 851957f7fb6f0bc4ce76be966d336802
    s) Skip and do nothing
    k) Keep just one (choose which in next step)
    r) Rename all to be different (by changing file.jpg to file-1.jpg)
    s/k/r> r
    two-1.txt: renamed from: two.txt
    two-2.txt: renamed from: two.txt
    two-3.txt: renamed from: two.txt

The result being

    $ rclone lsl drive:dupes
      6048320 2016-03-05 16:23:16.798000000 one.txt
       564374 2016-03-05 16:22:52.118000000 two-1.txt
      6048320 2016-03-05 16:22:46.185000000 two-2.txt
      1744073 2016-03-05 16:22:38.104000000 two-3.txt

Dedupe can be run non interactively using the `+"`"+`--dedupe-mode`+"`"+` flag or by using an extra parameter with the same value

  * `+"`"+`--dedupe-mode interactive`+"`"+` - interactive as above.
  * `+"`"+`--dedupe-mode skip`+"`"+` - removes identical files then skips anything left.
  * `+"`"+`--dedupe-mode first`+"`"+` - removes identical files then keeps the first one.
  * `+"`"+`--dedupe-mode newest`+"`"+` - removes identical files then keeps the newest one.
  * `+"`"+`--dedupe-mode oldest`+"`"+` - removes identical files then keeps the oldest one.
  * `+"`"+`--dedupe-mode largest`+"`"+` - removes identical files then keeps the largest one.
  * `+"`"+`--dedupe-mode smallest`+"`"+` - removes identical files then keeps the smallest one.
  * `+"`"+`--dedupe-mode rename`+"`"+` - removes identical files then renames the rest to be different.
  * `+"`"+`--dedupe-mode list`+"`"+` - lists duplicate dirs and files only and changes nothing.

For example, to rename all the identically named photos in your Google Photos directory, do

    rclone dedupe --dedupe-mode rename "drive:Google Photos"

Or

    rclone dedupe rename "drive:Google Photos"`),
			ox.Spec(`[mode] remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`by-hash`, `Find identical hashes rather than names`).
				String(`dedupe-mode`, `Dedupe mode interactive|skip|first|newest|oldest|largest|smallest|rename`, ox.Default("interactive")),
		),
		ox.Sub(
			ox.Usage(`delete`, `Remove the files in path.`),
			ox.Banner(`
Remove the files in path.  Unlike `+"`"+`purge`+"`"+` it obeys include/exclude
filters so can be used to selectively delete files.

`+"`"+`rclone delete`+"`"+` only deletes files but leaves the directory structure
alone. If you want to delete a directory and all of its contents use
the `+"`"+`purge`+"`"+` command.

If you supply the `+"`"+`--rmdirs`+"`"+` flag, it will remove all empty directories along with it.
You can also use the separate command `+"`"+`rmdir`+"`"+` or `+"`"+`rmdirs`+"`"+` to
delete empty directories only.

For example, to delete all files bigger than 100 MiB, you may first want to
check what would be deleted (use either):

    rclone --min-size 100M lsl remote:path
    rclone --dry-run --min-size 100M delete remote:path

Then proceed with the actual delete:

    rclone --min-size 100M delete remote:path

That reads "delete everything with a minimum size of 100 MiB", hence
delete all files bigger than 100 MiB.

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`rmdirs`, `rmdirs removes empty directories but leaves root intact`),
		),
		ox.Sub(
			ox.Usage(`deletefile`, `Remove a single file from remote.`),
			ox.Banner(`
Remove a single file from remote.  Unlike `+"`"+`delete`+"`"+` it cannot be used to
remove a directory and it doesn't obey include/exclude filters - if the specified file exists,
it will always be removed.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`genautocomplete`, `Output completion script for a given shell.`),
			ox.Banner(`
Generates a shell completion script for rclone.
Run with --help to list the supported shells.`),
			ox.Spec(`[command]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Sub(
				ox.Usage(`bash`, `Output bash completion script for rclone.`),
				ox.Banner(`
Generates a bash shell autocompletion script for rclone.

This writes to /etc/bash_completion.d/rclone by default so will
probably need to be run with sudo or as root, e.g.

    sudo rclone genautocomplete bash

Logout and login again to use the autocompletion scripts, or source
them directly

    . /etc/bash_completion

If you supply a command line argument the script will be written
there.

If output_file is "-", then the output will be written to stdout.`),
				ox.Spec(`[output_file] [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`fish`, `Output fish completion script for rclone.`),
				ox.Banner(`
Generates a fish autocompletion script for rclone.

This writes to /etc/fish/completions/rclone.fish by default so will
probably need to be run with sudo or as root, e.g.

    sudo rclone genautocomplete fish

Logout and login again to use the autocompletion scripts, or source
them directly

    . /etc/fish/completions/rclone.fish

If you supply a command line argument the script will be written
there.

If output_file is "-", then the output will be written to stdout.`),
				ox.Spec(`[output_file] [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`zsh`, `Output zsh completion script for rclone.`),
				ox.Banner(`
Generates a zsh autocompletion script for rclone.

This writes to /usr/share/zsh/vendor-completions/_rclone by default so will
probably need to be run with sudo or as root, e.g.

    sudo rclone genautocomplete zsh

Logout and login again to use the autocompletion scripts, or source
them directly

    autoload -U compinit && compinit

If you supply a command line argument the script will be written
there.

If output_file is "-", then the output will be written to stdout.`),
				ox.Spec(`[output_file] [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
		),
		ox.Sub(
			ox.Usage(`gendocs`, `Output markdown docs for rclone to the directory supplied.`),
			ox.Banner(`
This produces markdown docs for the rclone commands to the directory
supplied.  These are in a format suitable for hugo to render into the
rclone.org website.`),
			ox.Spec(`output_directory [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`hashsum`, `Produces a hashsum file for all the objects in the path.`),
			ox.Banner(`
Produces a hash file for all the objects in the path using the hash
named.  The output is in the same format as the standard
md5sum/sha1sum tool.

By default, the hash is requested from the remote.  If the hash is
not supported by the remote, no hash will be returned.  With the
download flag, the file will be downloaded from the remote and
hashed locally enabling any hash for any remote.

This command can also hash data received on standard input (stdin),
by not passing a remote:path, or by passing a hyphen as remote:path
when there is data to read (if not, the hypen will be treated literaly,
as a relative path).

Run without a hash to see the list of all supported hashes, e.g.

    $ rclone hashsum
    Supported hashes are:
      * md5
      * sha1
      * whirlpool
      * crc32
      * sha256
      * dropbox
      * mailru
      * quickxor

Then

    $ rclone hashsum MD5 remote:path

Note that hash names are case insensitive and values are output in lower case.`),
			ox.Spec(`<hash> remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`base64`, `Output base64 encoded hashsum`).
				String(`checkfile`, `Validate hashes against a given SUM file instead of printing them`, ox.Short("C")).
				Bool(`download`, `Download the file and hash it locally; if this flag is not specified, the hash is requested from the remote`).
				String(`output-file`, `Output hashsums to a file rather than the terminal`),
		),
		ox.Sub(
			ox.Usage(`link`, `Generate public link to file/folder.`),
			ox.Banner(`rclone link will create, retrieve or remove a public link to the given
file or folder.

    rclone link remote:path/to/file
    rclone link remote:path/to/folder/
    rclone link --unlink remote:path/to/folder/
    rclone link --expire 1d remote:path/to/file

If you supply the --expire flag, it will set the expiration time
otherwise it will use the default (100 years). **Note** not all
backends support the --expire flag - if the backend doesn't support it
then the link returned won't expire.

Use the --unlink flag to remove existing public links to the file or
folder. **Note** not all backends support "--unlink" flag - those that
don't will just ignore it.

If successful, the last line of the output will contain the
link. Exact capabilities depend on the remote, but the link will
always by default be created with the least constraints – e.g. no
expiry, no password protection, accessible without account.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				String(`expire`, `The amount of time that the link will be valid`, ox.Spec(`Duration`), ox.Default("off")).
				Bool(`unlink`, `Remove existing public link to file/folder`),
		),
		ox.Sub(
			ox.Usage(`listremotes`, `List all the remotes in the config file.`),
			ox.Banner(`
rclone listremotes lists all the available remotes from the config file.

When uses with the -l flag it lists the types too.`),
			ox.Spec(`[flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`long`, `Show the type as well as names`),
		),
		ox.Sub(
			ox.Usage(`ls`, `List the objects in the path with size and path.`),
			ox.Banner(`
Lists the objects in the source path to standard output in a human
readable format with size and path. Recurses by default.

Eg

    $ rclone ls swift:bucket
        60295 bevajer5jef
        90613 canole
        94467 diwogej7
        37600 fubuwic


Any of the filtering options can be applied to this command.

There are several related list commands

  * `+"`"+`ls`+"`"+` to list size and path of objects only
  * `+"`"+`lsl`+"`"+` to list modification time, size and path of objects only
  * `+"`"+`lsd`+"`"+` to list directories only
  * `+"`"+`lsf`+"`"+` to list objects and directories in easy to parse format
  * `+"`"+`lsjson`+"`"+` to list objects and directories in JSON format

`+"`"+`ls`+"`"+`,`+"`"+`lsl`+"`"+`,`+"`"+`lsd`+"`"+` are designed to be human-readable.
`+"`"+`lsf`+"`"+` is designed to be human and machine-readable.
`+"`"+`lsjson`+"`"+` is designed to be machine-readable.

Note that `+"`"+`ls`+"`"+` and `+"`"+`lsl`+"`"+` recurse by default - use `+"`"+`--max-depth 1`+"`"+` to stop the recursion.

The other list commands `+"`"+`lsd`+"`"+`,`+"`"+`lsf`+"`"+`,`+"`"+`lsjson`+"`"+` do not recurse by default - use `+"`"+`-R`+"`"+` to make them recurse.

Listing a non-existent directory will produce an error except for
remotes which can't have empty directories (e.g. s3, swift, or gcs -
the bucket-based remotes).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`lsd`, `List all directories/containers/buckets in the path.`),
			ox.Banner(`
Lists the directories in the source path to standard output. Does not
recurse by default.  Use the -R flag to recurse.

This command lists the total size of the directory (if known, -1 if
not), the modification time (if known, the current time if not), the
number of objects in the directory (if known, -1 if not) and the name
of the directory, Eg

    $ rclone lsd swift:
          494000 2018-04-26 08:43:20     10000 10000files
              65 2018-04-26 08:43:20         1 1File

Or

    $ rclone lsd drive:test
              -1 2016-10-17 17:41:53        -1 1000files
              -1 2017-01-03 14:40:54        -1 2500files
              -1 2017-07-08 14:39:28        -1 4000files

If you just want the directory names use "rclone lsf --dirs-only".


Any of the filtering options can be applied to this command.

There are several related list commands

  * `+"`"+`ls`+"`"+` to list size and path of objects only
  * `+"`"+`lsl`+"`"+` to list modification time, size and path of objects only
  * `+"`"+`lsd`+"`"+` to list directories only
  * `+"`"+`lsf`+"`"+` to list objects and directories in easy to parse format
  * `+"`"+`lsjson`+"`"+` to list objects and directories in JSON format

`+"`"+`ls`+"`"+`,`+"`"+`lsl`+"`"+`,`+"`"+`lsd`+"`"+` are designed to be human-readable.
`+"`"+`lsf`+"`"+` is designed to be human and machine-readable.
`+"`"+`lsjson`+"`"+` is designed to be machine-readable.

Note that `+"`"+`ls`+"`"+` and `+"`"+`lsl`+"`"+` recurse by default - use `+"`"+`--max-depth 1`+"`"+` to stop the recursion.

The other list commands `+"`"+`lsd`+"`"+`,`+"`"+`lsf`+"`"+`,`+"`"+`lsjson`+"`"+` do not recurse by default - use `+"`"+`-R`+"`"+` to make them recurse.

Listing a non-existent directory will produce an error except for
remotes which can't have empty directories (e.g. s3, swift, or gcs -
the bucket-based remotes).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`recursive`, `Recurse into the listing`, ox.Short("R")),
		),
		ox.Sub(
			ox.Usage(`lsf`, `List directories and objects in remote:path formatted for parsing.`),
			ox.Banner(`
List the contents of the source path (directories and objects) to
standard output in a form which is easy to parse by scripts.  By
default this will just be the names of the objects and directories,
one per line.  The directories will have a / suffix.

Eg

    $ rclone lsf swift:bucket
    bevajer5jef
    canole
    diwogej7
    ferejej3gux/
    fubuwic

Use the --format option to control what gets listed.  By default this
is just the path, but you can use these parameters to control the
output:

    p - path
    s - size
    t - modification time
    h - hash
    i - ID of object
    o - Original ID of underlying object
    m - MimeType of object if known
    e - encrypted name
    T - tier of storage if known, e.g. "Hot" or "Cool"

So if you wanted the path, size and modification time, you would use
--format "pst", or maybe --format "tsp" to put the path last.

Eg

    $ rclone lsf  --format "tsp" swift:bucket
    2016-06-25 18:55:41;60295;bevajer5jef
    2016-06-25 18:55:43;90613;canole
    2016-06-25 18:55:43;94467;diwogej7
    2018-04-26 08:50:45;0;ferejej3gux/
    2016-06-25 18:55:40;37600;fubuwic

If you specify "h" in the format you will get the MD5 hash by default,
use the "--hash" flag to change which hash you want.  Note that this
can be returned as an empty string if it isn't available on the object
(and for directories), "ERROR" if there was an error reading it from
the object and "UNSUPPORTED" if that object does not support that hash
type.

For example, to emulate the md5sum command you can use

    rclone lsf -R --hash MD5 --format hp --separator "  " --files-only .

Eg

    $ rclone lsf -R --hash MD5 --format hp --separator "  " --files-only swift:bucket
    7908e352297f0f530b84a756f188baa3  bevajer5jef
    cd65ac234e6fea5925974a51cdd865cc  canole
    03b5341b4f234b9d984d03ad076bae91  diwogej7
    8fd37c3810dd660778137ac3a66cc06d  fubuwic
    99713e14a4c4ff553acaf1930fad985b  gixacuh7ku

(Though "rclone md5sum ." is an easier way of typing this.)

By default the separator is ";" this can be changed with the
--separator flag.  Note that separators aren't escaped in the path so
putting it last is a good strategy.

Eg

    $ rclone lsf  --separator "," --format "tshp" swift:bucket
    2016-06-25 18:55:41,60295,7908e352297f0f530b84a756f188baa3,bevajer5jef
    2016-06-25 18:55:43,90613,cd65ac234e6fea5925974a51cdd865cc,canole
    2016-06-25 18:55:43,94467,03b5341b4f234b9d984d03ad076bae91,diwogej7
    2018-04-26 08:52:53,0,,ferejej3gux/
    2016-06-25 18:55:40,37600,8fd37c3810dd660778137ac3a66cc06d,fubuwic

You can output in CSV standard format.  This will escape things in "
if they contain ,

Eg

    $ rclone lsf --csv --files-only --format ps remote:path
    test.log,22355
    test.sh,449
    "this file contains a comma, in the file name.txt",6

Note that the --absolute parameter is useful for making lists of files
to pass to an rclone copy with the --files-from-raw flag.

For example, to find all the files modified within one day and copy
those only (without traversing the whole directory structure):

    rclone lsf --absolute --files-only --max-age 1d /path/to/local > new_files
    rclone copy --files-from-raw new_files /path/to/local remote:path


Any of the filtering options can be applied to this command.

There are several related list commands

  * `+"`"+`ls`+"`"+` to list size and path of objects only
  * `+"`"+`lsl`+"`"+` to list modification time, size and path of objects only
  * `+"`"+`lsd`+"`"+` to list directories only
  * `+"`"+`lsf`+"`"+` to list objects and directories in easy to parse format
  * `+"`"+`lsjson`+"`"+` to list objects and directories in JSON format

`+"`"+`ls`+"`"+`,`+"`"+`lsl`+"`"+`,`+"`"+`lsd`+"`"+` are designed to be human-readable.
`+"`"+`lsf`+"`"+` is designed to be human and machine-readable.
`+"`"+`lsjson`+"`"+` is designed to be machine-readable.

Note that `+"`"+`ls`+"`"+` and `+"`"+`lsl`+"`"+` recurse by default - use `+"`"+`--max-depth 1`+"`"+` to stop the recursion.

The other list commands `+"`"+`lsd`+"`"+`,`+"`"+`lsf`+"`"+`,`+"`"+`lsjson`+"`"+` do not recurse by default - use `+"`"+`-R`+"`"+` to make them recurse.

Listing a non-existent directory will produce an error except for
remotes which can't have empty directories (e.g. s3, swift, or gcs -
the bucket-based remotes).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`absolute`, `Put a leading / in front of path names`).
				Bool(`csv`, `Output in CSV format`).
				Bool(`dir-slash`, `Append a slash to directory names`, ox.Default("true"), ox.Short("d")).
				Bool(`dirs-only`, `Only list directories`).
				Bool(`files-only`, `Only list files`).
				String(`format`, `Output format - see  help for details`, ox.Default("p"), ox.Short("F")).
				String(`hash`, `Use this hash when h is used in the format MD5|SHA-1|DropboxHash`, ox.Spec(`h`), ox.Default("md5")).
				Bool(`recursive`, `Recurse into the listing`, ox.Short("R")).
				String(`separator`, `Separator for the items in the format`, ox.Default(";"), ox.Short("s")),
		),
		ox.Sub(
			ox.Usage(`lsjson`, `List directories and objects in the path in JSON format.`),
			ox.Banner(`List directories and objects in the path in JSON format.

The output is an array of Items, where each Item looks like this

   {
      "Hashes" : {
         "SHA-1" : "f572d396fae9206628714fb2ce00f72e94f2258f",
         "MD5" : "b1946ac92492d2347c6235b4d2611184",
         "DropboxHash" : "ecb65bb98f9d905b70458986c39fcbad7715e5f2fcc3b1f07767d7c83e2438cc"
      },
      "ID": "y2djkhiujf83u33",
      "OrigID": "UYOJVTUW00Q1RzTDA",
      "IsBucket" : false,
      "IsDir" : false,
      "MimeType" : "application/octet-stream",
      "ModTime" : "2017-05-31T16:15:57.034468261+01:00",
      "Name" : "file.txt",
      "Encrypted" : "v0qpsdq8anpci8n929v3uu9338",
      "EncryptedPath" : "kja9098349023498/v0qpsdq8anpci8n929v3uu9338",
      "Path" : "full/path/goes/here/file.txt",
      "Size" : 6,
      "Tier" : "hot",
   }

If --hash is not specified the Hashes property won't be emitted. The
types of hash can be specified with the --hash-type parameter (which
may be repeated). If --hash-type is set then it implies --hash.

If --no-modtime is specified then ModTime will be blank. This can
speed things up on remotes where reading the ModTime takes an extra
request (e.g. s3, swift).

If --no-mimetype is specified then MimeType will be blank. This can
speed things up on remotes where reading the MimeType takes an extra
request (e.g. s3, swift).

If --encrypted is not specified the Encrypted won't be emitted.

If --dirs-only is not specified files in addition to directories are
returned

If --files-only is not specified directories in addition to the files
will be returned.

if --stat is set then a single JSON blob will be returned about the
item pointed to. This will return an error if the item isn't found.
However on bucket based backends (like s3, gcs, b2, azureblob etc) if
the item isn't found it will return an empty directory as it isn't
possible to tell empty directories from missing directories there.

The Path field will only show folders below the remote path being listed.
If "remote:path" contains the file "subfolder/file.txt", the Path for "file.txt"
will be "subfolder/file.txt", not "remote:path/subfolder/file.txt".
When used without --recursive the Path will always be the same as Name.

If the directory is a bucket in a bucket-based backend, then
"IsBucket" will be set to true. This key won't be present unless it is
"true".

The time is in RFC3339 format with up to nanosecond precision.  The
number of decimal digits in the seconds will depend on the precision
that the remote can hold the times, so if times are accurate to the
nearest millisecond (e.g. Google Drive) then 3 digits will always be
shown ("2017-05-31T16:15:57.034+01:00") whereas if the times are
accurate to the nearest second (Dropbox, Box, WebDav, etc.) no digits
will be shown ("2017-05-31T16:15:57+01:00").

The whole output can be processed as a JSON blob, or alternatively it
can be processed line by line as each item is written one to a line.

Any of the filtering options can be applied to this command.

There are several related list commands

  * `+"`"+`ls`+"`"+` to list size and path of objects only
  * `+"`"+`lsl`+"`"+` to list modification time, size and path of objects only
  * `+"`"+`lsd`+"`"+` to list directories only
  * `+"`"+`lsf`+"`"+` to list objects and directories in easy to parse format
  * `+"`"+`lsjson`+"`"+` to list objects and directories in JSON format

`+"`"+`ls`+"`"+`,`+"`"+`lsl`+"`"+`,`+"`"+`lsd`+"`"+` are designed to be human-readable.
`+"`"+`lsf`+"`"+` is designed to be human and machine-readable.
`+"`"+`lsjson`+"`"+` is designed to be machine-readable.

Note that `+"`"+`ls`+"`"+` and `+"`"+`lsl`+"`"+` recurse by default - use `+"`"+`--max-depth 1`+"`"+` to stop the recursion.

The other list commands `+"`"+`lsd`+"`"+`,`+"`"+`lsf`+"`"+`,`+"`"+`lsjson`+"`"+` do not recurse by default - use `+"`"+`-R`+"`"+` to make them recurse.

Listing a non-existent directory will produce an error except for
remotes which can't have empty directories (e.g. s3, swift, or gcs -
the bucket-based remotes).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`dirs-only`, `Show only directories in the listing`).
				Bool(`encrypted`, `Show the encrypted names`, ox.Short("M")).
				Bool(`files-only`, `Show only files in the listing`).
				Bool(`hash`, `Include hashes in the output (may take longer)`).
				Array(`hash-type`, `Show only this hash type (may be repeated)`).
				Bool(`no-mimetype`, `Don't read the mime type (can speed things up)`).
				Bool(`no-modtime`, `Don't read the modification time (can speed things up)`).
				Bool(`original`, `Show the ID of the underlying Object`).
				Bool(`recursive`, `Recurse into the listing`, ox.Short("R")).
				Bool(`stat`, `Just return the info for the pointed to file`),
		),
		ox.Sub(
			ox.Usage(`lsl`, `List the objects in path with modification time, size and path.`),
			ox.Banner(`
Lists the objects in the source path to standard output in a human
readable format with modification time, size and path. Recurses by default.

Eg

    $ rclone lsl swift:bucket
        60295 2016-06-25 18:55:41.062626927 bevajer5jef
        90613 2016-06-25 18:55:43.302607074 canole
        94467 2016-06-25 18:55:43.046609333 diwogej7
        37600 2016-06-25 18:55:40.814629136 fubuwic


Any of the filtering options can be applied to this command.

There are several related list commands

  * `+"`"+`ls`+"`"+` to list size and path of objects only
  * `+"`"+`lsl`+"`"+` to list modification time, size and path of objects only
  * `+"`"+`lsd`+"`"+` to list directories only
  * `+"`"+`lsf`+"`"+` to list objects and directories in easy to parse format
  * `+"`"+`lsjson`+"`"+` to list objects and directories in JSON format

`+"`"+`ls`+"`"+`,`+"`"+`lsl`+"`"+`,`+"`"+`lsd`+"`"+` are designed to be human-readable.
`+"`"+`lsf`+"`"+` is designed to be human and machine-readable.
`+"`"+`lsjson`+"`"+` is designed to be machine-readable.

Note that `+"`"+`ls`+"`"+` and `+"`"+`lsl`+"`"+` recurse by default - use `+"`"+`--max-depth 1`+"`"+` to stop the recursion.

The other list commands `+"`"+`lsd`+"`"+`,`+"`"+`lsf`+"`"+`,`+"`"+`lsjson`+"`"+` do not recurse by default - use `+"`"+`-R`+"`"+` to make them recurse.

Listing a non-existent directory will produce an error except for
remotes which can't have empty directories (e.g. s3, swift, or gcs -
the bucket-based remotes).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`md5sum`, `Produces an md5sum file for all the objects in the path.`),
			ox.Banner(`
Produces an md5sum file for all the objects in the path.  This
is in the same format as the standard md5sum tool produces.

By default, the hash is requested from the remote.  If MD5 is
not supported by the remote, no hash will be returned.  With the
download flag, the file will be downloaded from the remote and
hashed locally enabling MD5 for any remote.

This command can also hash data received on standard input (stdin),
by not passing a remote:path, or by passing a hyphen as remote:path
when there is data to read (if not, the hypen will be treated literaly,
as a relative path).`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`base64`, `Output base64 encoded hashsum`).
				String(`checkfile`, `Validate hashes against a given SUM file instead of printing them`, ox.Short("C")).
				Bool(`download`, `Download the file and hash it locally; if this flag is not specified, the hash is requested from the remote`).
				String(`output-file`, `Output hashsums to a file rather than the terminal`),
		),
		ox.Sub(
			ox.Usage(`mkdir`, `Make the path if it doesn't already exist.`),
			ox.Banner(`Make the path if it doesn't already exist.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`mount`, `Mount the remote as file system on a mountpoint.`),
			ox.Banner(`
rclone mount allows Linux, FreeBSD, macOS and Windows to
mount any of Rclone's cloud storage systems as a file system with
FUSE.

First set up your remote using `+"`"+`rclone config`+"`"+`.  Check it works with `+"`"+`rclone ls`+"`"+` etc.

On Linux and macOS, you can run mount in either foreground or background (aka
daemon) mode. Mount runs in foreground mode by default. Use the `+"`"+`--daemon`+"`"+` flag
to force background mode. On Windows you can run mount in foreground only,
the flag is ignored.

In background mode rclone acts as a generic Unix mount program: the main
program starts, spawns background rclone process to setup and maintain the
mount, waits until success or timeout and exits with appropriate code
(killing the child process if it fails).

On Linux/macOS/FreeBSD start the mount like this, where `+"`"+`/path/to/local/mount`+"`"+`
is an **empty** **existing** directory:

    rclone mount remote:path/to/files /path/to/local/mount

On Windows you can start a mount in different ways. See [below](#mounting-modes-on-windows)
for details. If foreground mount is used interactively from a console window,
rclone will serve the mount and occupy the console so another window should be
used to work with the mount until rclone is interrupted e.g. by pressing Ctrl-C.

The following examples will mount to an automatically assigned drive,
to specific drive letter `+"`"+`X:`+"`"+`, to path `+"`"+`C:\path\parent\mount`+"`"+`
(where parent directory or drive must exist, and mount must **not** exist,
and is not supported when [mounting as a network drive](#mounting-modes-on-windows)), and
the last example will mount as network share `+"`"+`\\cloud\remote`+"`"+` and map it to an
automatically assigned drive:

    rclone mount remote:path/to/files *
    rclone mount remote:path/to/files X:
    rclone mount remote:path/to/files C:\path\parent\mount
    rclone mount remote:path/to/files \\cloud\remote

When the program ends while in foreground mode, either via Ctrl+C or receiving
a SIGINT or SIGTERM signal, the mount should be automatically stopped.

When running in background mode the user will have to stop the mount manually:

    # Linux
    fusermount -u /path/to/local/mount
    # OS X
    umount /path/to/local/mount

The umount operation can fail, for example when the mountpoint is busy.
When that happens, it is the user's responsibility to stop the mount manually.

The size of the mounted file system will be set according to information retrieved
from the remote, the same as returned by the [rclone about](https://rclone.org/commands/rclone_about/)
command. Remotes with unlimited storage may report the used size only,
then an additional 1 PiB of free space is assumed. If the remote does not
[support](https://rclone.org/overview/#optional-features) the about feature
at all, then 1 PiB is set as both the total and the free size.

### Installing on Windows

To run rclone mount on Windows, you will need to
download and install [WinFsp](http://www.secfs.net/winfsp/).

[WinFsp](https://github.com/billziss-gh/winfsp) is an open-source
Windows File System Proxy which makes it easy to write user space file
systems for Windows.  It provides a FUSE emulation layer which rclone
uses combination with [cgofuse](https://github.com/billziss-gh/cgofuse).
Both of these packages are by Bill Zissimopoulos who was very helpful
during the implementation of rclone mount for Windows.

#### Mounting modes on windows

Unlike other operating systems, Microsoft Windows provides a different filesystem
type for network and fixed drives. It optimises access on the assumption fixed
disk drives are fast and reliable, while network drives have relatively high latency
and less reliability. Some settings can also be differentiated between the two types,
for example that Windows Explorer should just display icons and not create preview
thumbnails for image and video files on network drives.

In most cases, rclone will mount the remote as a normal, fixed disk drive by default.
However, you can also choose to mount it as a remote network drive, often described
as a network share. If you mount an rclone remote using the default, fixed drive mode
and experience unexpected program errors, freezes or other issues, consider mounting
as a network drive instead.

When mounting as a fixed disk drive you can either mount to an unused drive letter,
or to a path representing a **non-existent** subdirectory of an **existing** parent
directory or drive. Using the special value `+"`"+`*`+"`"+` will tell rclone to
automatically assign the next available drive letter, starting with Z: and moving backward.
Examples:

    rclone mount remote:path/to/files *
    rclone mount remote:path/to/files X:
    rclone mount remote:path/to/files C:\path\parent\mount
    rclone mount remote:path/to/files X:

Option `+"`"+`--volname`+"`"+` can be used to set a custom volume name for the mounted
file system. The default is to use the remote name and path.

To mount as network drive, you can add option `+"`"+`--network-mode`+"`"+`
to your mount command. Mounting to a directory path is not supported in
this mode, it is a limitation Windows imposes on junctions, so the remote must always
be mounted to a drive letter.

    rclone mount remote:path/to/files X: --network-mode

A volume name specified with `+"`"+`--volname`+"`"+` will be used to create the network share path.
A complete UNC path, such as `+"`"+`\\cloud\remote`+"`"+`, optionally with path
`+"`"+`\\cloud\remote\madeup\path`+"`"+`, will be used as is. Any other
string will be used as the share part, after a default prefix `+"`"+`\\server\`+"`"+`.
If no volume name is specified then `+"`"+`\\server\share`+"`"+` will be used.
You must make sure the volume name is unique when you are mounting more than one drive,
or else the mount command will fail. The share name will treated as the volume label for
the mapped drive, shown in Windows Explorer etc, while the complete
`+"`"+`\\server\share`+"`"+` will be reported as the remote UNC path by
`+"`"+`net use`+"`"+` etc, just like a normal network drive mapping.

If you specify a full network share UNC path with `+"`"+`--volname`+"`"+`, this will implicitely
set the `+"`"+`--network-mode`+"`"+` option, so the following two examples have same result:

    rclone mount remote:path/to/files X: --network-mode
    rclone mount remote:path/to/files X: --volname \\server\share

You may also specify the network share UNC path as the mountpoint itself. Then rclone
will automatically assign a drive letter, same as with `+"`"+`*`+"`"+` and use that as
mountpoint, and instead use the UNC path specified as the volume name, as if it were
specified with the `+"`"+`--volname`+"`"+` option. This will also implicitely set
the `+"`"+`--network-mode`+"`"+` option. This means the following two examples have same result:

    rclone mount remote:path/to/files \\cloud\remote
    rclone mount remote:path/to/files * --volname \\cloud\remote

There is yet another way to enable network mode, and to set the share path,
and that is to pass the "native" libfuse/WinFsp option directly:
`+"`"+`--fuse-flag --VolumePrefix=\server\share`+"`"+`. Note that the path
must be with just a single backslash prefix in this case.


*Note:* In previous versions of rclone this was the only supported method.

[Read more about drive mapping](https://en.wikipedia.org/wiki/Drive_mapping)

See also [Limitations](#limitations) section below.

#### Windows filesystem permissions

The FUSE emulation layer on Windows must convert between the POSIX-based
permission model used in FUSE, and the permission model used in Windows,
based on access-control lists (ACL).

The mounted filesystem will normally get three entries in its access-control list (ACL),
representing permissions for the POSIX permission scopes: Owner, group and others.
By default, the owner and group will be taken from the current user, and the built-in
group "Everyone" will be used to represent others. The user/group can be customized
with FUSE options "UserName" and "GroupName",
e.g. `+"`"+`-o UserName=user123 -o GroupName="Authenticated Users"`+"`"+`.
The permissions on each entry will be set according to [options](#options)
`+"`"+`--dir-perms`+"`"+` and `+"`"+`--file-perms`+"`"+`, which takes a value in traditional
[numeric notation](https://en.wikipedia.org/wiki/File-system_permissions#Numeric_notation).

The default permissions corresponds to `+"`"+`--file-perms 0666 --dir-perms 0777`+"`"+`,
i.e. read and write permissions to everyone. This means you will not be able
to start any programs from the the mount. To be able to do that you must add
execute permissions, e.g. `+"`"+`--file-perms 0777 --dir-perms 0777`+"`"+` to add it
to everyone. If the program needs to write files, chances are you will have
to enable [VFS File Caching](#vfs-file-caching) as well (see also [limitations](#limitations)).

Note that the mapping of permissions is not always trivial, and the result
you see in Windows Explorer may not be exactly like you expected.
For example, when setting a value that includes write access, this will be
mapped to individual permissions "write attributes", "write data" and "append data",
but not "write extended attributes". Windows will then show this as basic
permission "Special" instead of "Write", because "Write" includes the
"write extended attributes" permission.

If you set POSIX permissions for only allowing access to the owner, using
`+"`"+`--file-perms 0600 --dir-perms 0700`+"`"+`, the user group and the built-in "Everyone"
group will still be given some special permissions, such as "read attributes"
and "read permissions", in Windows. This is done for compatibility reasons,
e.g. to allow users without additional permissions to be able to read basic
metadata about files like in UNIX. One case that may arise is that other programs
(incorrectly) interprets this as the file being accessible by everyone. For example
an SSH client may warn about "unprotected private key file".

WinFsp 2021 (version 1.9) introduces a new FUSE option "FileSecurity",
that allows the complete specification of file security descriptors using
[SDDL](https://docs.microsoft.com/en-us/windows/win32/secauthz/security-descriptor-string-format).
With this you can work around issues such as the mentioned "unprotected private key file"
by specifying `+"`"+`-o FileSecurity="D:P(A;;FA;;;OW)"`+"`"+`, for file all access (FA) to the owner (OW).

#### Windows caveats

Drives created as Administrator are not visible to other accounts,
not even an account that was elevated to Administrator with the
User Account Control (UAC) feature. A result of this is that if you mount
to a drive letter from a Command Prompt run as Administrator, and then try
to access the same drive from Windows Explorer (which does not run as
Administrator), you will not be able to see the mounted drive.

If you don't need to access the drive from applications running with
administrative privileges, the easiest way around this is to always
create the mount from a non-elevated command prompt.

To make mapped drives available to the user account that created them
regardless if elevated or not, there is a special Windows setting called
[linked connections](https://docs.microsoft.com/en-us/troubleshoot/windows-client/networking/mapped-drives-not-available-from-elevated-command#detail-to-configure-the-enablelinkedconnections-registry-entry)
that can be enabled.

It is also possible to make a drive mount available to everyone on the system,
by running the process creating it as the built-in SYSTEM account.
There are several ways to do this: One is to use the command-line
utility [PsExec](https://docs.microsoft.com/en-us/sysinternals/downloads/psexec),
from Microsoft's Sysinternals suite, which has option `+"`"+`-s`+"`"+` to start
processes as the SYSTEM account. Another alternative is to run the mount
command from a Windows Scheduled Task, or a Windows Service, configured
to run as the SYSTEM account. A third alternative is to use the
[WinFsp.Launcher infrastructure](https://github.com/billziss-gh/winfsp/wiki/WinFsp-Service-Architecture)).
Note that when running rclone as another user, it will not use
the configuration file from your profile unless you tell it to
with the [`+"`"+`--config`+"`"+`](https://rclone.org/docs/#config-config-file) option.
Read more in the [install documentation](https://rclone.org/install/).

Note that mapping to a directory path, instead of a drive letter,
does not suffer from the same limitations.

### Limitations

Without the use of `+"`"+`--vfs-cache-mode`+"`"+` this can only write files
sequentially, it can only seek when reading.  This means that many
applications won't work with their files on an rclone mount without
`+"`"+`--vfs-cache-mode writes`+"`"+` or `+"`"+`--vfs-cache-mode full`+"`"+`.
See the [VFS File Caching](#vfs-file-caching) section for more info.

The bucket-based remotes (e.g. Swift, S3, Google Compute Storage, B2,
Hubic) do not support the concept of empty directories, so empty
directories will have a tendency to disappear once they fall out of
the directory cache.

When `+"`"+`rclone mount`+"`"+` is invoked on Unix with `+"`"+`--daemon`+"`"+` flag, the main rclone
program will wait for the background mount to become ready or until the timeout
specified by the `+"`"+`--daemon-wait`+"`"+` flag. On Linux it can check mount status using
ProcFS so the flag in fact sets **maximum** time to wait, while the real wait
can be less. On macOS / BSD the time to wait is constant and the check is
performed only at the end. We advise you to set wait time on macOS reasonably.

Only supported on Linux, FreeBSD, OS X and Windows at the moment.

### rclone mount vs rclone sync/copy

File systems expect things to be 100% reliable, whereas cloud storage
systems are a long way from 100% reliable. The rclone sync/copy
commands cope with this with lots of retries.  However rclone mount
can't use retries in the same way without making local copies of the
uploads. Look at the [VFS File Caching](#vfs-file-caching)
for solutions to make mount more reliable.

### Attribute caching

You can use the flag `+"`"+`--attr-timeout`+"`"+` to set the time the kernel caches
the attributes (size, modification time, etc.) for directory entries.

The default is `+"`"+`1s`+"`"+` which caches files just long enough to avoid
too many callbacks to rclone from the kernel.

In theory 0s should be the correct value for filesystems which can
change outside the control of the kernel. However this causes quite a
few problems such as
[rclone using too much memory](https://github.com/rclone/rclone/issues/2157),
[rclone not serving files to samba](https://forum.rclone.org/t/rclone-1-39-vs-1-40-mount-issue/5112)
and [excessive time listing directories](https://github.com/rclone/rclone/issues/2095#issuecomment-371141147).

The kernel can cache the info about a file for the time given by
`+"`"+`--attr-timeout`+"`"+`. You may see corruption if the remote file changes
length during this window.  It will show up as either a truncated file
or a file with garbage on the end.  With `+"`"+`--attr-timeout 1s`+"`"+` this is
very unlikely but not impossible.  The higher you set `+"`"+`--attr-timeout`+"`"+`
the more likely it is.  The default setting of "1s" is the lowest
setting which mitigates the problems above.

If you set it higher (`+"`"+`10s`+"`"+` or `+"`"+`1m`+"`"+` say) then the kernel will call
back to rclone less often making it more efficient, however there is
more chance of the corruption issue above.

If files don't change on the remote outside of the control of rclone
then there is no chance of corruption.

This is the same as setting the attr_timeout option in mount.fuse.

### Filters

Note that all the rclone filters can be used to select a subset of the
files to be visible in the mount.

### systemd

When running rclone mount as a systemd service, it is possible
to use Type=notify. In this case the service will enter the started state
after the mountpoint has been successfully set up.
Units having the rclone mount service specified as a requirement
will see all files and folders immediately in this mode.

Note that systemd runs mount units without any environment variables including
`+"`"+`PATH`+"`"+` or `+"`"+`HOME`+"`"+`. This means that tilde (`+"`"+`~`+"`"+`) expansion will not work
and you should provide `+"`"+`--config`+"`"+` and `+"`"+`--cache-dir`+"`"+` explicitly as absolute
paths via rclone arguments.
Since mounting requires the `+"`"+`fusermount`+"`"+` program, rclone will use the fallback
PATH of `+"`"+`/bin:/usr/bin`+"`"+` in this scenario. Please ensure that `+"`"+`fusermount`+"`"+`
is present on this PATH.

### Rclone as Unix mount helper

The core Unix program `+"`"+`/bin/mount`+"`"+` normally takes the `+"`"+`-t FSTYPE`+"`"+` argument
then runs the `+"`"+`/sbin/mount.FSTYPE`+"`"+` helper program passing it mount options
as `+"`"+`-o key=val,...`+"`"+` or `+"`"+`--opt=...`+"`"+`. Automount (classic or systemd) behaves
in a similar way.

rclone by default expects GNU-style flags `+"`"+`--key val`+"`"+`. To run it as a mount
helper you should symlink rclone binary to `+"`"+`/sbin/mount.rclone`+"`"+` and optionally
`+"`"+`/usr/bin/rclonefs`+"`"+`, e.g. `+"`"+`ln -s /usr/bin/rclone /sbin/mount.rclone`+"`"+`.
rclone will detect it and translate command-line arguments appropriately.

Now you can run classic mounts like this:
`+"`"+``+"`"+``+"`"+`
mount sftp1:subdir /mnt/data -t rclone -o vfs_cache_mode=writes,sftp_key_file=/path/to/pem
`+"`"+``+"`"+``+"`"+`

or create systemd mount units:
`+"`"+``+"`"+``+"`"+`
# /etc/systemd/system/mnt-data.mount
[Unit]
After=network-online.target
[Mount]
Type=rclone
What=sftp1:subdir
Where=/mnt/data
Options=rw,allow_other,args2env,vfs-cache-mode=writes,config=/etc/rclone.conf,cache-dir=/var/rclone
`+"`"+``+"`"+``+"`"+`

optionally accompanied by systemd automount unit
`+"`"+``+"`"+``+"`"+`
# /etc/systemd/system/mnt-data.automount
[Unit]
After=network-online.target
Before=remote-fs.target
[Automount]
Where=/mnt/data
TimeoutIdleSec=600
[Install]
WantedBy=multi-user.target
`+"`"+``+"`"+``+"`"+`

or add in `+"`"+`/etc/fstab`+"`"+` a line like
`+"`"+``+"`"+``+"`"+`
sftp1:subdir /mnt/data rclone rw,noauto,nofail,_netdev,x-systemd.automount,args2env,vfs_cache_mode=writes,config=/etc/rclone.conf,cache_dir=/var/cache/rclone 0 0
`+"`"+``+"`"+``+"`"+`

or use classic Automountd.
Remember to provide explicit `+"`"+`config=...,cache-dir=...`+"`"+` as a workaround for
mount units being run without `+"`"+`HOME`+"`"+`.

Rclone in the mount helper mode will split `+"`"+`-o`+"`"+` argument(s) by comma, replace `+"`"+`_`+"`"+`
by `+"`"+`-`+"`"+` and prepend `+"`"+`--`+"`"+` to get the command-line flags. Options containing commas
or spaces can be wrapped in single or double quotes. Any inner quotes inside outer
quotes of the same type should be doubled.

Mount option syntax includes a few extra options treated specially:

- `+"`"+`env.NAME=VALUE`+"`"+` will set an environment variable for the mount process.
  This helps with Automountd and Systemd.mount which don't allow setting
  custom environment for mount helpers.
  Typically you will use `+"`"+`env.HTTPS_PROXY=proxy.host:3128`+"`"+` or `+"`"+`env.HOME=/root`+"`"+`
- `+"`"+`command=cmount`+"`"+` can be used to run `+"`"+`cmount`+"`"+` or any other rclone command
  rather than the default `+"`"+`mount`+"`"+`.
- `+"`"+`args2env`+"`"+` will pass mount options to the mount helper running in background
  via environment variables instead of command line arguments. This allows to
  hide secrets from such commands as `+"`"+`ps`+"`"+` or `+"`"+`pgrep`+"`"+`.
- `+"`"+`vv...`+"`"+` will be transformed into appropriate `+"`"+`--verbose=N`+"`"+`
- standard mount options like `+"`"+`x-systemd.automount`+"`"+`, `+"`"+`_netdev`+"`"+`, `+"`"+`nosuid`+"`"+` and alike
  are intended only for Automountd and ignored by rclone.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.`),
			ox.Spec(`remote:path /path/to/mountpoint [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`allow-non-empty`, `Allow mounting over a non-empty directory (not supported on Windows)`).
				Bool(`allow-other`, `Allow access to other users (not supported on Windows)`).
				Bool(`allow-root`, `Allow access to root user (not supported on Windows)`).
				Bool(`async-read`, `Use asynchronous reads (not supported on Windows)`, ox.Default("true")).
				Duration(`attr-timeout`, `Time for which file/directory attributes are cached`, ox.Default("1s")).
				Bool(`daemon`, `Run mount in background and exit parent process (as background output is suppressed, use --log-file with --log-format=pid,... to monitor) (not supported on Windows)`).
				Duration(`daemon-timeout`, `Time limit for rclone to respond to kernel (not supported on Windows)`).
				Duration(`daemon-wait`, `Time to wait for ready mount from daemon (maximum time on Linux, constant sleep time on OSX/BSD) (not supported on Windows)`, ox.Default("1m0s")).
				Bool(`debug-fuse`, `Debug the FUSE internals - needs -v`).
				Bool(`default-permissions`, `Makes kernel enforce access control based on the file mode (not supported on Windows)`).
				String(`devname`, `Set the device name - default is remote:path`).
				Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
				String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
				String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
				Array(`fuse-flag`, `Flags or arguments to be passed direct to libfuse/WinFsp (repeat if required)`).
				Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
				String(`max-read-ahead`, `The number of bytes that can be prefetched for sequential reads (not supported on Windows)`, ox.Spec(`SizeSuffix`), ox.Default("128Ki")).
				Bool(`network-mode`, `Mount as remote network drive, instead of fixed disk drive (supported on Windows only)`).
				Bool(`no-checksum`, `Don't compare checksums on up/download`).
				Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
				Bool(`no-seek`, `Don't allow seeking in files`).
				Bool(`noappledouble`, `Ignore Apple Double (._) and .DS_Store files (supported on OSX only)`, ox.Default("true")).
				Bool(`noapplexattr`, `Ignore all "com.apple.*" extended attributes (supported on OSX only)`).
				Array(`option`, `Option for libfuse/WinFsp (repeat if required)`, ox.Short("o")).
				Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
				Bool(`read-only`, `Mount read-only`).
				Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
				Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
				Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
				String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
				String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
				Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
				Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
				Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
				String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
				String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
				String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
				Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
				String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
				Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
				Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")).
				String(`volname`, `Set the volume name (supported on Windows and OSX only)`).
				Bool(`write-back-cache`, `Makes kernel buffer writes before sending them to rclone (without this, writethrough caching is used) (not supported on Windows)`),
		),
		ox.Sub(
			ox.Usage(`move`, `Move files from source to dest.`),
			ox.Banner(`
Moves the contents of the source directory to the destination
directory. Rclone will error if the source and destination overlap and
the remote does not support a server-side directory move operation.

If no filters are in use and if possible this will server-side move
`+"`"+`source:path`+"`"+` into `+"`"+`dest:path`+"`"+`. After this `+"`"+`source:path`+"`"+` will no
longer exist.

Otherwise for each file in `+"`"+`source:path`+"`"+` selected by the filters (if
any) this will move it into `+"`"+`dest:path`+"`"+`.  If possible a server-side
move will be used, otherwise it will copy it (server-side if possible)
into `+"`"+`dest:path`+"`"+` then delete the original (if no errors on copy) in
`+"`"+`source:path`+"`"+`.

If you want to delete empty source directories after move, use the --delete-empty-src-dirs flag.

See the [--no-traverse](/docs/#no-traverse) option for controlling
whether rclone lists the destination directory or not.  Supplying this
option when moving a small number of files into a large destination
can speed transfers up greatly.

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.

**Note**: Use the `+"`"+`-P`+"`"+`/`+"`"+`--progress`+"`"+` flag to view real-time transfer statistics.`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`create-empty-src-dirs`, `Create empty source dirs on destination after move`).
				Bool(`delete-empty-src-dirs`, `Delete empty source dirs after move`),
		),
		ox.Sub(
			ox.Usage(`moveto`, `Move file or directory from source to dest.`),
			ox.Banner(`
If source:path is a file or directory then it moves it to a file or
directory named dest:path.

This can be used to rename files or upload single files to other than
their existing name.  If the source is a directory then it acts exactly
like the move command.

So

    rclone moveto src dst

where src and dst are rclone paths, either remote:path or
/path/to/local or C:\windows\path\if\on\windows.

This will:

    if src is file
        move it to dst, overwriting an existing file if it exists
    if src is directory
        move it to dst, overwriting existing files if they exist
        see move command for full details

This doesn't transfer files that are identical on src and dst, testing
by size and modification time or MD5SUM.  src will be deleted on
successful transfer.

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.

**Note**: Use the `+"`"+`-P`+"`"+`/`+"`"+`--progress`+"`"+` flag to view real-time transfer statistics.`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`ncdu`, `Explore a remote with a text based user interface.`),
			ox.Banner(`
This displays a text based user interface allowing the navigation of a
remote. It is most useful for answering the question - "What is using
all my disk space?".

{{< asciinema 157793 >}}

To make the user interface it first scans the entire remote given and
builds an in memory representation.  rclone ncdu can be used during
this scanning phase and you will see it building up the directory
structure as it goes along.

Here are the keys - press '?' to toggle the help on and off

     ↑,↓ or k,j to Move
     →,l to enter
     ←,h to return
     c toggle counts
     g toggle graph
     a toggle average size in directory
     u toggle human-readable format
     n,s,C,A sort by name,size,count,average size
     d delete file/directory
     y copy current path to clipboard
     Y display current path
     ^L refresh screen
     ? to toggle help on and off
     q/ESC/c-C to quit

This an homage to the [ncdu tool](https://dev.yorhel.nl/ncdu) but for
rclone remotes.  It is missing lots of features at the moment
but is useful as it stands.

Note that it might take some time to delete big files/folders. The
UI won't respond in the meantime since the deletion is done synchronously.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`obscure`, `Obscure password for use in the rclone config file.`),
			ox.Banner(`In the rclone config file, human-readable passwords are
obscured. Obscuring them is done by encrypting them and writing them
out in base64. This is **not** a secure way of encrypting these
passwords as rclone can decrypt them - it is to prevent "eyedropping"
- namely someone seeing a password in the rclone config file by
accident.

Many equally important things (like access tokens) are not obscured in
the config file. However it is very hard to shoulder surf a 64
character hex token.

This command can also accept a password through STDIN instead of an
argument by passing a hyphen as an argument. This will use the first
line of STDIN as the password not including the trailing newline.

echo "secretpassword" | rclone obscure -

If there is no data on STDIN to read, rclone obscure will default to
obfuscating the hyphen itself.

If you want to encrypt the config file then please use config file
encryption - see [rclone config](/commands/rclone_config/) for more
info.`),
			ox.Spec(`password [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`purge`, `Remove the path and all of its contents.`),
			ox.Banner(`
Remove the path and all of its contents.  Note that this does not obey
include/exclude filters - everything will be removed.  Use the `+"`"+`delete`+"`"+`
command if you want to selectively delete files. To delete empty directories only,
use command `+"`"+`rmdir`+"`"+` or `+"`"+`rmdirs`+"`"+`.

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`rc`, `Run a command against a running rclone.`),
			ox.Banner(`

This runs a command against a running rclone.  Use the --url flag to
specify an non default URL to connect on.  This can be either a
":port" which is taken to mean "http://localhost:port" or a
"host:port" which is taken to mean "http://host:port"

A username and password can be passed in with --user and --pass.

Note that --rc-addr, --rc-user, --rc-pass will be read also for --url,
--user, --pass.

Arguments should be passed in as parameter=value.

The result will be returned as a JSON object by default.

The --json parameter can be used to pass in a JSON blob as an input
instead of key=value arguments.  This is the only way of passing in
more complicated values.

The -o/--opt option can be used to set a key "opt" with key, value
options in the form "-o key=value" or "-o key". It can be repeated as
many times as required. This is useful for rc commands which take the
"opt" parameter which by convention is a dictionary of strings.

    -o key=value -o key2

Will place this in the "opt" value

    {"key":"value", "key2","")


The -a/--arg option can be used to set strings in the "arg" value. It
can be repeated as many times as required. This is useful for rc
commands which take the "arg" parameter which by convention is a list
of strings.

    -a value -a value2

Will place this in the "arg" value

    ["value", "value2"]

Use --loopback to connect to the rclone instance running "rclone rc".
This is very useful for testing commands without having to run an
rclone rc server, e.g.:

    rclone rc --loopback operations/about fs=/

Use "rclone rc" to see a list of all possible commands.`),
			ox.Spec(`commands parameter [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Array(`arg`, `Argument placed in the "arg" array`, ox.Short("a")).
				String(`json`, `Input JSON - use instead of key=value args`).
				Bool(`loopback`, `If set connect to this rclone instance not via HTTP`).
				Bool(`no-output`, `If set, don't output the JSON result`).
				Array(`opt`, `Option in the form name=value or name placed in the "opt" array`, ox.Short("o")).
				String(`pass`, `Password to use to connect to rclone remote control`).
				String(`url`, `URL to connect to rclone remote control`, ox.Default("http://localhost:5572/")).
				String(`user`, `Username to use to rclone remote control`),
		),
		ox.Sub(
			ox.Usage(`rcat`, `Copies standard input to file on remote.`),
			ox.Banner(`
rclone rcat reads from standard input (stdin) and copies it to a
single remote file.

    echo "hello world" | rclone rcat remote:path/to/file
    ffmpeg - | rclone rcat remote:path/to/file

If the remote file already exists, it will be overwritten.

rcat will try to upload small files in a single request, which is
usually more efficient than the streaming/chunked upload endpoints,
which use multiple requests. Exact behaviour depends on the remote.
What is considered a small file may be set through
`+"`"+`--streaming-upload-cutoff`+"`"+`. Uploading only starts after
the cutoff is reached or if the file ends before that. The data
must fit into RAM. The cutoff needs to be small enough to adhere
the limits of your remote, please see there. Generally speaking,
setting this cutoff too high will decrease your performance.

Use the |--size| flag to preallocate the file in advance at the remote end
and actually stream it, even if remote backend doesn't support streaming.

|--size| should be the exact size of the input stream in bytes. If the
size of the stream is different in length to the |--size| passed in
then the transfer will likely fail.

Note that the upload can also not be retried because the data is
not kept around until the upload succeeds. If you need to transfer
a lot of data, you're better off caching locally and then
`+"`"+`rclone move`+"`"+` it to the destination.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Int(`size`, `File size hint to preallocate`, ox.Default("-1")),
		),
		ox.Sub(
			ox.Usage(`rcd`, `Run rclone listening to remote control commands only.`),
			ox.Banner(`
This runs rclone so that it only listens to remote control commands.

This is useful if you are controlling rclone via the rc API.

If you pass in a path to a directory, rclone will serve that directory
for GET requests on the URL passed in.  It will also open the URL in
the browser when rclone is run.

See the [rc documentation](/rc/) for more info on the rc flags.`),
			ox.Spec(`<path to files to serve>* [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`rmdir`, `Remove the empty directory at path.`),
			ox.Banner(`
This removes empty directory given by path. Will not remove the path if it
has any objects in it, not even empty subdirectories. Use
command `+"`"+`rmdirs`+"`"+` (or `+"`"+`delete`+"`"+` with option `+"`"+`--rmdirs`+"`"+`)
to do that.

To delete a path and any objects in it, use `+"`"+`purge`+"`"+` command.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`rmdirs`, `Remove empty directories under the path.`),
			ox.Banner(`
This recursively removes any empty directories (including directories
that only contain empty directories), that it finds under the path.
The root path itself will also be removed if it is empty, unless
you supply the `+"`"+`--leave-root`+"`"+` flag.

Use command `+"`"+`rmdir`+"`"+` to delete just the empty directory
given by path, not recurse.

This is useful for tidying up remotes that rclone has left a lot of
empty directories in. For example the `+"`"+`delete`+"`"+` command will
delete files but leave the directory structure (unless used with
option `+"`"+`--rmdirs`+"`"+`).

To delete a path and any objects in it, use `+"`"+`purge`+"`"+` command.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`leave-root`, `Do not remove root directory if empty`),
		),
		ox.Sub(
			ox.Usage(`selfupdate`, `Update the rclone binary.`),
			ox.Banner(`
This command downloads the latest release of rclone and replaces
the currently running binary. The download is verified with a hashsum
and cryptographically signed signature.

If used without flags (or with implied `+"`"+`--stable`+"`"+` flag), this command
will install the latest stable release. However, some issues may be fixed
(or features added) only in the latest beta release. In such cases you should
run the command with the `+"`"+`--beta`+"`"+` flag, i.e. `+"`"+`rclone selfupdate --beta`+"`"+`.
You can check in advance what version would be installed by adding the
`+"`"+`--check`+"`"+` flag, then repeat the command without it when you are satisfied.

Sometimes the rclone team may recommend you a concrete beta or stable
rclone release to troubleshoot your issue or add a bleeding edge feature.
The `+"`"+`--version VER`+"`"+` flag, if given, will update to the concrete version
instead of the latest one. If you omit micro version from `+"`"+`VER`+"`"+` (for
example `+"`"+`1.53`+"`"+`), the latest matching micro version will be used.

Upon successful update rclone will print a message that contains a previous
version number. You will need it if you later decide to revert your update
for some reason. Then you'll have to note the previous version and run the
following command: `+"`"+`rclone selfupdate [--beta] OLDVER`+"`"+`.
If the old version contains only dots and digits (for example `+"`"+`v1.54.0`+"`"+`)
then it's a stable release so you won't need the `+"`"+`--beta`+"`"+` flag. Beta releases
have an additional information similar to `+"`"+`v1.54.0-beta.5111.06f1c0c61`+"`"+`.
(if you are a developer and use a locally built rclone, the version number
will end with `+"`"+`-DEV`+"`"+`, you will have to rebuild it as it obviously can't
be distributed).

If you previously installed rclone via a package manager, the package may
include local documentation or configure services. You may wish to update
with the flag `+"`"+`--package deb`+"`"+` or `+"`"+`--package rpm`+"`"+` (whichever is correct for
your OS) to update these too. This command with the default `+"`"+`--package zip`+"`"+`
will update only the rclone executable so the local manual may become
inaccurate after it.

The `+"`"+`rclone mount`+"`"+` command (https://rclone.org/commands/rclone_mount/) may
or may not support extended FUSE options depending on the build and OS.
`+"`"+`selfupdate`+"`"+` will refuse to update if the capability would be discarded.

Note: Windows forbids deletion of a currently running executable so this
command will rename the old executable to 'rclone.old.exe' upon success.

Please note that this command was not available before rclone version 1.55.
If it fails for you with the message `+"`"+`unknown command "selfupdate"`+"`"+` then
you will need to update manually following the install instructions located
at https://rclone.org/install/`),
			ox.Spec(`[flags]`),
			ox.Aliases("self-update"),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`beta`, `Install beta release`).
				Bool(`check`, `Check for latest release, do not download`).
				String(`output`, `Save the downloaded binary at a given path (default: replace running binary)`).
				String(`package`, `Package format: zip|deb|rpm (default: zip)`).
				Bool(`stable`, `Install stable release (this is the default)`),
		),
		ox.Sub(
			ox.Usage(`serve`, `Serve a remote over a protocol.`),
			ox.Banner(`rclone serve is used to serve a remote over a given protocol. This
command requires the use of a subcommand to specify the protocol, e.g.

    rclone serve http remote:

Each subcommand has its own options which you can see in their help.`),
			ox.Spec(`<protocol> [opts] <remote> [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Sub(
				ox.Usage(`dlna`, `Serve remote:path over DLNA`),
				ox.Banner(`rclone serve dlna is a DLNA media server for media stored in an rclone remote. Many
devices, such as the Xbox and PlayStation, can automatically discover this server in the LAN
and play audio/video from it. VLC is also supported. Service discovery uses UDP multicast
packets (SSDP) and will thus only work on LANs.

Rclone will list all files present in the remote, without filtering based on media formats or
file extensions. Additionally, there is no media transcoding support. This means that some
players might show files that they are not able to play back correctly.


### Server options

Use `+"`"+`--addr`+"`"+` to specify which IP address and port the server should
listen on, e.g. `+"`"+`--addr 1.2.3.4:8000`+"`"+` or `+"`"+`--addr :8080`+"`"+` to listen to all
IPs.

Use `+"`"+`--name`+"`"+` to choose the friendly server name, which is by
default "rclone (hostname)".

Use `+"`"+`--log-trace`+"`"+` in conjunction with `+"`"+`-vv`+"`"+` to enable additional debug
logging of all UPNP traffic.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `The ip:port or :port to bind the DLNA http server to`, ox.Default(":7879")).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Bool(`log-trace`, `Enable trace logging of SOAP traffic`).
					String(`name`, `Name of DLNA server`).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					Bool(`read-only`, `Mount read-only`).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")),
			),
			ox.Sub(
				ox.Usage(`docker`, `Serve any remote on docker's volume plugin API.`),
				ox.Banner(`
This command implements the Docker volume plugin API allowing docker to use
rclone as a data storage mechanism for various cloud providers.
rclone provides [docker volume plugin](/docker) based on it.

To create a docker plugin, one must create a Unix or TCP socket that Docker
will look for when you use the plugin and then it listens for commands from
docker daemon and runs the corresponding code when necessary.
Docker plugins can run as a managed plugin under control of the docker daemon
or as an independent native service. For testing, you can just run it directly
from the command line, for example:
`+"`"+``+"`"+``+"`"+`
sudo rclone serve docker --base-dir /tmp/rclone-volumes --socket-addr localhost:8787 -vv
`+"`"+``+"`"+``+"`"+`

Running `+"`"+`rclone serve docker`+"`"+` will create the said socket, listening for
commands from Docker to create the necessary Volumes. Normally you need not
give the `+"`"+`--socket-addr`+"`"+` flag. The API will listen on the unix domain socket
at `+"`"+`/run/docker/plugins/rclone.sock`+"`"+`. In the example above rclone will create
a TCP socket and a small file `+"`"+`/etc/docker/plugins/rclone.spec`+"`"+` containing
the socket address. We use `+"`"+`sudo`+"`"+` because both paths are writeable only by
the root user.

If you later decide to change listening socket, the docker daemon must be
restarted to reconnect to `+"`"+`/run/docker/plugins/rclone.sock`+"`"+`
or parse new `+"`"+`/etc/docker/plugins/rclone.spec`+"`"+`. Until you restart, any
volume related docker commands will timeout trying to access the old socket.
Running directly is supported on **Linux only**, not on Windows or MacOS.
This is not a problem with managed plugin mode described in details
in the [full documentation](https://rclone.org/docker).

The command will create volume mounts under the path given by `+"`"+`--base-dir`+"`"+`
(by default `+"`"+`/var/lib/docker-volumes/rclone`+"`"+` available only to root)
and maintain the JSON formatted file `+"`"+`docker-plugin.state`+"`"+` in the rclone cache
directory with book-keeping records of created and mounted volumes.

All mount and VFS options are submitted by the docker daemon via API, but
you can also provide defaults on the command line as well as set path to the
config file and cache directory or adjust logging verbosity.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.`),
				ox.Spec(`[flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Bool(`allow-non-empty`, `Allow mounting over a non-empty directory (not supported on Windows)`).
					Bool(`allow-other`, `Allow access to other users (not supported on Windows)`).
					Bool(`allow-root`, `Allow access to root user (not supported on Windows)`).
					Bool(`async-read`, `Use asynchronous reads (not supported on Windows)`, ox.Default("true")).
					Duration(`attr-timeout`, `Time for which file/directory attributes are cached`, ox.Default("1s")).
					String(`base-dir`, `Base directory for volumes`, ox.Default("/var/lib/docker-volumes/$APPNAME")).
					Bool(`daemon`, `Run mount in background and exit parent process (as background output is suppressed, use --log-file with --log-format=pid,... to monitor) (not supported on Windows)`).
					Duration(`daemon-timeout`, `Time limit for rclone to respond to kernel (not supported on Windows)`).
					Duration(`daemon-wait`, `Time to wait for ready mount from daemon (maximum time on Linux, constant sleep time on OSX/BSD) (not supported on Windows)`, ox.Default("1m0s")).
					Bool(`debug-fuse`, `Debug the FUSE internals - needs -v`).
					Bool(`default-permissions`, `Makes kernel enforce access control based on the file mode (not supported on Windows)`).
					String(`devname`, `Set the device name - default is remote:path`).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Bool(`forget-state`, `Skip restoring previous state`).
					Array(`fuse-flag`, `Flags or arguments to be passed direct to libfuse/WinFsp (repeat if required)`).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					String(`max-read-ahead`, `The number of bytes that can be prefetched for sequential reads (not supported on Windows)`, ox.Spec(`SizeSuffix`), ox.Default("128Ki")).
					Bool(`network-mode`, `Mount as remote network drive, instead of fixed disk drive (supported on Windows only)`).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					Bool(`no-spec`, `Do not write spec file`).
					Bool(`noappledouble`, `Ignore Apple Double (._) and .DS_Store files (supported on OSX only)`, ox.Default("true")).
					Bool(`noapplexattr`, `Ignore all "com.apple.*" extended attributes (supported on OSX only)`).
					Array(`option`, `Option for libfuse/WinFsp (repeat if required)`, ox.Short("o")).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					Bool(`read-only`, `Mount read-only`).
					String(`socket-addr`, `Address <host:port> or absolute path (default: /run/docker/plugins/rclone.sock)`).
					Int(`socket-gid`, `GID for unix socket (default: current process GID)`, ox.Default("1000")).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")).
					String(`volname`, `Set the volume name (supported on Windows and OSX only)`).
					Bool(`write-back-cache`, `Makes kernel buffer writes before sending them to rclone (without this, writethrough caching is used) (not supported on Windows)`),
			),
			ox.Sub(
				ox.Usage(`ftp`, `Serve remote:path over FTP.`),
				ox.Banner(`
rclone serve ftp implements a basic ftp server to serve the
remote over FTP protocol. This can be viewed with a ftp client
or you can make a remote of type ftp to read and write it.

### Server options

Use --addr to specify which IP address and port the server should
listen on, e.g. --addr 1.2.3.4:8000 or --addr :8080 to listen to all
IPs.  By default it only listens on localhost.  You can use port
:0 to let the OS choose an available port.

If you set --addr to listen on a public or LAN accessible IP address
then using Authentication is advised - see the next section for info.

#### Authentication

By default this will serve files without needing a login.

You can set a single username and password with the --user and --pass flags.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.

### Auth Proxy

If you supply the parameter `+"`"+`--auth-proxy /path/to/program`+"`"+` then
rclone will use that program to generate backends on the fly which
then are used to authenticate incoming requests.  This uses a simple
JSON based protocol with input on STDIN and output on STDOUT.

**PLEASE NOTE:** `+"`"+`--auth-proxy`+"`"+` and `+"`"+`--authorized-keys`+"`"+` cannot be used
together, if `+"`"+`--auth-proxy`+"`"+` is set the authorized keys option will be
ignored.

There is an example program
[bin/test_proxy.py](https://github.com/rclone/rclone/blob/master/test_proxy.py)
in the rclone source code.

The program's job is to take a `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+` on the input and turn
those into the config for a backend on STDOUT in JSON format.  This
config will have any default parameters for the backend added, but it
won't use configuration from environment variables or command line
options - it is the job of the proxy program to make a complete
config.

This config generated must have this extra parameter
- `+"`"+`_root`+"`"+` - root to use for the backend

And it may have this parameter
- `+"`"+`_obscure`+"`"+` - comma separated strings for parameters to obscure

If password authentication was used by the client, input to the proxy
process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"pass": "mypassword"
}
`+"`"+``+"`"+``+"`"+`

If public-key authentication was used by the client, input to the
proxy process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDuwESFdAe14hVS6omeyX7edc...JQdf"
}
`+"`"+``+"`"+``+"`"+`

And as an example return this on STDOUT

`+"`"+``+"`"+``+"`"+`
{
	"type": "sftp",
	"_root": "",
	"_obscure": "pass",
	"user": "me",
	"pass": "mypassword",
	"host": "sftp.example.com"
}
`+"`"+``+"`"+``+"`"+`

This would mean that an SFTP backend would be created on the fly for
the `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+`/`+"`"+`public_key`+"`"+` returned in the output to the host given.  Note
that since `+"`"+`_obscure`+"`"+` is set to `+"`"+`pass`+"`"+`, rclone will obscure the `+"`"+`pass`+"`"+`
parameter before creating the backend (which is required for sftp
backends).

The program can manipulate the supplied `+"`"+`user`+"`"+` in any way, for example
to make proxy to many different sftp backends, you could make the
`+"`"+`user`+"`"+` be `+"`"+`user@example.com`+"`"+` and then set the `+"`"+`host`+"`"+` to `+"`"+`example.com`+"`"+`
in the output and the user to `+"`"+`user`+"`"+`. For security you'd probably want
to restrict the `+"`"+`host`+"`"+` to a limited list.

Note that an internal cache is keyed on `+"`"+`user`+"`"+` so only use that for
configuration, don't use `+"`"+`pass`+"`"+` or `+"`"+`public_key`+"`"+`.  This also means that if a user's
password or public-key is changed the cache will need to expire (which takes 5 mins)
before it takes effect.

This can be used to build general purpose proxies to any kind of
backend that rclone supports.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `IPaddress:Port or :Port to bind server to`, ox.Default("localhost:2121")).
					String(`auth-proxy`, `A program to use to create the backend from the auth`).
					String(`cert`, `TLS PEM key (concatenation of certificate and CA certificate)`).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					String(`key`, `TLS PEM Private key`).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					String(`pass`, `Password for authentication (empty value allow every password)`).
					String(`passive-port`, `Passive port range to use`, ox.Default("30000-32000")).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					String(`public-ip`, `Public IP address to advertise for passive connections`).
					Bool(`read-only`, `Mount read-only`).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					String(`user`, `User name for authentication`, ox.Default("anonymous")).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")),
			),
			ox.Sub(
				ox.Usage(`http`, `Serve the remote over HTTP.`),
				ox.Banner(`rclone serve http implements a basic web server to serve the remote
over HTTP.  This can be viewed in a web browser or you can make a
remote of type http read from it.

You can use the filter flags (e.g. --include, --exclude) to control what
is served.

The server will log errors.  Use -v to see access logs.

--bwlimit will be respected for file transfers.  Use --stats to
control the stats printing.

### Server options

Use --addr to specify which IP address and port the server should
listen on, eg --addr 1.2.3.4:8000 or --addr :8080 to listen to all
IPs.  By default it only listens on localhost.  You can use port
:0 to let the OS choose an available port.

If you set --addr to listen on a public or LAN accessible IP address
then using Authentication is advised - see the next section for info.

--server-read-timeout and --server-write-timeout can be used to
control the timeouts on the server.  Note that this is the total time
for a transfer.

--max-header-bytes controls the maximum number of bytes the server will
accept in the HTTP header.

--baseurl controls the URL prefix that rclone serves from.  By default
rclone will serve from the root.  If you used --baseurl "/rclone" then
rclone would serve from a URL starting with "/rclone/".  This is
useful if you wish to proxy rclone serve.  Rclone automatically
inserts leading and trailing "/" on --baseurl, so --baseurl "rclone",
--baseurl "/rclone" and --baseurl "/rclone/" are all treated
identically.

#### SSL/TLS

By default this will serve over http.  If you want you can serve over
https.  You will need to supply the --cert and --key flags.  If you
wish to do client side certificate validation then you will need to
supply --client-ca also.

--cert should be a either a PEM encoded certificate or a concatenation
of that with the CA certificate.  --key should be the PEM encoded
private key and --client-ca should be the PEM encoded client
certificate authority certificate.

#### Template

--template allows a user to specify a custom markup template for http
and webdav serve functions.  The server exports the following markup
to be used within the template to server pages:

| Parameter   | Description |
| :---------- | :---------- |
| .Name       | The full path of a file/directory. |
| .Title      | Directory listing of .Name |
| .Sort       | The current sort used.  This is changeable via ?sort= parameter |
|             | Sort Options: namedirfirst,name,size,time (default namedirfirst) |
| .Order      | The current ordering used.  This is changeable via ?order= parameter |
|             | Order Options: asc,desc (default asc) |
| .Query      | Currently unused. |
| .Breadcrumb | Allows for creating a relative navigation |
|-- .Link     | The relative to the root link of the Text. |
|-- .Text     | The Name of the directory. |
| .Entries    | Information about a specific file/directory. |
|-- .URL      | The 'url' of an entry.  |
|-- .Leaf     | Currently same as 'URL' but intended to be 'just' the name. |
|-- .IsDir    | Boolean for if an entry is a directory or not. |
|-- .Size     | Size in Bytes of the entry. |
|-- .ModTime  | The UTC timestamp of an entry. |

#### Authentication

By default this will serve files without needing a login.

You can either use an htpasswd file which can take lots of users, or
set a single username and password with the --user and --pass flags.

Use --htpasswd /path/to/htpasswd to provide an htpasswd file.  This is
in standard apache format and supports MD5, SHA1 and BCrypt for basic
authentication.  Bcrypt is recommended.

To create an htpasswd file:

    touch htpasswd
    htpasswd -B htpasswd user
    htpasswd -B htpasswd anotherUser

The password file can be updated while rclone is running.

Use --realm to set the authentication realm.

Use --salt to change the password hashing salt from the default.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `IPaddress:Port or :Port to bind server to`, ox.Default("127.0.0.1:8080")).
					String(`baseurl`, `Prefix for URLs - leave blank for root`).
					String(`cert`, `SSL PEM key (concatenation of certificate and CA certificate)`).
					String(`client-ca`, `Client certificate authority to verify clients with`).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					String(`htpasswd`, `A htpasswd file - if not provided no authentication is done`).
					String(`key`, `SSL PEM Private key`).
					Int(`max-header-bytes`, `Maximum size of request header`, ox.Default("4096")).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					String(`pass`, `Password for authentication`).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					Bool(`read-only`, `Mount read-only`).
					String(`realm`, `Realm for authentication`).
					String(`salt`, `Password hashing salt`, ox.Default("dlPL2MqE")).
					Duration(`server-read-timeout`, `Timeout for server reading data`, ox.Default("1h0m0s")).
					Duration(`server-write-timeout`, `Timeout for server writing data`, ox.Default("1h0m0s")).
					String(`template`, `User-specified template`).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					String(`user`, `User name for authentication`).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")),
			),
			ox.Sub(
				ox.Usage(`restic`, `Serve the remote for restic's REST API.`),
				ox.Banner(`rclone serve restic implements restic's REST backend API
over HTTP.  This allows restic to use rclone as a data storage
mechanism for cloud providers that restic does not support directly.

[Restic](https://restic.net/) is a command-line program for doing
backups.

The server will log errors.  Use -v to see access logs.

--bwlimit will be respected for file transfers.  Use --stats to
control the stats printing.

### Setting up rclone for use by restic ###

First [set up a remote for your chosen cloud provider](/docs/#configure).

Once you have set up the remote, check it is working with, for example
"rclone lsd remote:".  You may have called the remote something other
than "remote:" - just substitute whatever you called it in the
following instructions.

Now start the rclone restic server

    rclone serve restic -v remote:backup

Where you can replace "backup" in the above by whatever path in the
remote you wish to use.

By default this will serve on "localhost:8080" you can change this
with use of the "--addr" flag.

You might wish to start this server on boot.

Adding --cache-objects=false will cause rclone to stop caching objects
returned from the List call. Caching is normally desirable as it speeds
up downloading objects, saves transactions and uses very little memory.

### Setting up restic to use rclone ###

Now you can [follow the restic
instructions](http://restic.readthedocs.io/en/latest/030_preparing_a_new_repo.html#rest-server)
on setting up restic.

Note that you will need restic 0.8.2 or later to interoperate with
rclone.

For the example above you will want to use "http://localhost:8080/" as
the URL for the REST server.

For example:

    $ export RESTIC_REPOSITORY=rest:http://localhost:8080/
    $ export RESTIC_PASSWORD=yourpassword
    $ restic init
    created restic backend 8b1a4b56ae at rest:http://localhost:8080/

    Please note that knowledge of your password is required to access
    the repository. Losing your password means that your data is
    irrecoverably lost.
    $ restic backup /path/to/files/to/backup
    scan [/path/to/files/to/backup]
    scanned 189 directories, 312 files in 0:00
    [0:00] 100.00%  38.128 MiB / 38.128 MiB  501 / 501 items  0 errors  ETA 0:00
    duration: 0:00
    snapshot 45c8fdd8 saved

#### Multiple repositories ####

Note that you can use the endpoint to host multiple repositories.  Do
this by adding a directory name or path after the URL.  Note that
these **must** end with /.  Eg

    $ export RESTIC_REPOSITORY=rest:http://localhost:8080/user1repo/
    # backup user1 stuff
    $ export RESTIC_REPOSITORY=rest:http://localhost:8080/user2repo/
    # backup user2 stuff

#### Private repositories ####

The "--private-repos" flag can be used to limit users to repositories starting
with a path of `+"`"+`/<username>/`+"`"+`.

### Server options

Use --addr to specify which IP address and port the server should
listen on, e.g. --addr 1.2.3.4:8000 or --addr :8080 to listen to all
IPs.  By default it only listens on localhost.  You can use port
:0 to let the OS choose an available port.

If you set --addr to listen on a public or LAN accessible IP address
then using Authentication is advised - see the next section for info.

--server-read-timeout and --server-write-timeout can be used to
control the timeouts on the server.  Note that this is the total time
for a transfer.

--max-header-bytes controls the maximum number of bytes the server will
accept in the HTTP header.

--baseurl controls the URL prefix that rclone serves from.  By default
rclone will serve from the root.  If you used --baseurl "/rclone" then
rclone would serve from a URL starting with "/rclone/".  This is
useful if you wish to proxy rclone serve.  Rclone automatically
inserts leading and trailing "/" on --baseurl, so --baseurl "rclone",
--baseurl "/rclone" and --baseurl "/rclone/" are all treated
identically.

--template allows a user to specify a custom markup template for http
and webdav serve functions.  The server exports the following markup
to be used within the template to server pages:

| Parameter   | Description |
| :---------- | :---------- |
| .Name       | The full path of a file/directory. |
| .Title      | Directory listing of .Name |
| .Sort       | The current sort used.  This is changeable via ?sort= parameter |
|             | Sort Options: namedirfirst,name,size,time (default namedirfirst) |
| .Order      | The current ordering used.  This is changeable via ?order= parameter |
|             | Order Options: asc,desc (default asc) |
| .Query      | Currently unused. |
| .Breadcrumb | Allows for creating a relative navigation |
|-- .Link     | The relative to the root link of the Text. |
|-- .Text     | The Name of the directory. |
| .Entries    | Information about a specific file/directory. |
|-- .URL      | The 'url' of an entry.  |
|-- .Leaf     | Currently same as 'URL' but intended to be 'just' the name. |
|-- .IsDir    | Boolean for if an entry is a directory or not. |
|-- .Size     | Size in Bytes of the entry. |
|-- .ModTime  | The UTC timestamp of an entry. |

#### Authentication

By default this will serve files without needing a login.

You can either use an htpasswd file which can take lots of users, or
set a single username and password with the --user and --pass flags.

Use --htpasswd /path/to/htpasswd to provide an htpasswd file.  This is
in standard apache format and supports MD5, SHA1 and BCrypt for basic
authentication.  Bcrypt is recommended.

To create an htpasswd file:

    touch htpasswd
    htpasswd -B htpasswd user
    htpasswd -B htpasswd anotherUser

The password file can be updated while rclone is running.

Use --realm to set the authentication realm.

#### SSL/TLS

By default this will serve over http.  If you want you can serve over
https.  You will need to supply the --cert and --key flags.  If you
wish to do client side certificate validation then you will need to
supply --client-ca also.

--cert should be either a PEM encoded certificate or a concatenation
of that with the CA certificate.  --key should be the PEM encoded
private key and --client-ca should be the PEM encoded client
certificate authority certificate.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `IPaddress:Port or :Port to bind server to`, ox.Default("localhost:8080")).
					Bool(`append-only`, `Disallow deletion of repository data`).
					String(`baseurl`, `Prefix for URLs - leave blank for root`).
					Bool(`cache-objects`, `Cache listed objects`, ox.Default("true")).
					String(`cert`, `SSL PEM key (concatenation of certificate and CA certificate)`).
					String(`client-ca`, `Client certificate authority to verify clients with`).
					String(`htpasswd`, `htpasswd file - if not provided no authentication is done`).
					String(`key`, `SSL PEM Private key`).
					Int(`max-header-bytes`, `Maximum size of request header`, ox.Default("4096")).
					String(`pass`, `Password for authentication`).
					Bool(`private-repos`, `Users can only access their private repo`).
					String(`realm`, `Realm for authentication`, ox.Default("$APPNAME")).
					Duration(`server-read-timeout`, `Timeout for server reading data`, ox.Default("1h0m0s")).
					Duration(`server-write-timeout`, `Timeout for server writing data`, ox.Default("1h0m0s")).
					Bool(`stdio`, `Run an HTTP2 server on stdin/stdout`).
					String(`template`, `User-specified template`).
					String(`user`, `User name for authentication`),
			),
			ox.Sub(
				ox.Usage(`sftp`, `Serve the remote over SFTP.`),
				ox.Banner(`rclone serve sftp implements an SFTP server to serve the remote
over SFTP.  This can be used with an SFTP client or you can make a
remote of type sftp to use with it.

You can use the filter flags (e.g. --include, --exclude) to control what
is served.

The server will log errors.  Use -v to see access logs.

--bwlimit will be respected for file transfers.  Use --stats to
control the stats printing.

You must provide some means of authentication, either with --user/--pass,
an authorized keys file (specify location with --authorized-keys - the
default is the same as ssh), an --auth-proxy, or set the --no-auth flag for no
authentication when logging in.

Note that this also implements a small number of shell commands so
that it can provide md5sum/sha1sum/df information for the rclone sftp
backend.  This means that is can support SHA1SUMs, MD5SUMs and the
about command when paired with the rclone sftp backend.

If you don't supply a host --key then rclone will generate rsa, ecdsa
and ed25519 variants, and cache them for later use in rclone's cache
directory (see "rclone help flags cache-dir") in the "serve-sftp"
directory.

By default the server binds to localhost:2022 - if you want it to be
reachable externally then supply "--addr :2022" for example.

Note that the default of "--vfs-cache-mode off" is fine for the rclone
sftp backend, but it may not be with other SFTP clients.

If --stdio is specified, rclone will serve SFTP over stdio, which can
be used with sshd via ~/.ssh/authorized_keys, for example:

    restrict,command="rclone serve sftp --stdio ./photos" ssh-rsa ...

On the client you need to set "--transfers 1" when using --stdio.
Otherwise multiple instances of the rclone server are started by OpenSSH
which can lead to "corrupted on transfer" errors. This is the case because
the client chooses indiscriminately which server to send commands to while
the servers all have different views of the state of the filing system.

The "restrict" in authorized_keys prevents SHA1SUMs and MD5SUMs from beeing
used. Omitting "restrict" and using --sftp-path-override to enable
checksumming is possible but less secure and you could use the SFTP server
provided by OpenSSH in this case.


### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.

### Auth Proxy

If you supply the parameter `+"`"+`--auth-proxy /path/to/program`+"`"+` then
rclone will use that program to generate backends on the fly which
then are used to authenticate incoming requests.  This uses a simple
JSON based protocol with input on STDIN and output on STDOUT.

**PLEASE NOTE:** `+"`"+`--auth-proxy`+"`"+` and `+"`"+`--authorized-keys`+"`"+` cannot be used
together, if `+"`"+`--auth-proxy`+"`"+` is set the authorized keys option will be
ignored.

There is an example program
[bin/test_proxy.py](https://github.com/rclone/rclone/blob/master/test_proxy.py)
in the rclone source code.

The program's job is to take a `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+` on the input and turn
those into the config for a backend on STDOUT in JSON format.  This
config will have any default parameters for the backend added, but it
won't use configuration from environment variables or command line
options - it is the job of the proxy program to make a complete
config.

This config generated must have this extra parameter
- `+"`"+`_root`+"`"+` - root to use for the backend

And it may have this parameter
- `+"`"+`_obscure`+"`"+` - comma separated strings for parameters to obscure

If password authentication was used by the client, input to the proxy
process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"pass": "mypassword"
}
`+"`"+``+"`"+``+"`"+`

If public-key authentication was used by the client, input to the
proxy process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDuwESFdAe14hVS6omeyX7edc...JQdf"
}
`+"`"+``+"`"+``+"`"+`

And as an example return this on STDOUT

`+"`"+``+"`"+``+"`"+`
{
	"type": "sftp",
	"_root": "",
	"_obscure": "pass",
	"user": "me",
	"pass": "mypassword",
	"host": "sftp.example.com"
}
`+"`"+``+"`"+``+"`"+`

This would mean that an SFTP backend would be created on the fly for
the `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+`/`+"`"+`public_key`+"`"+` returned in the output to the host given.  Note
that since `+"`"+`_obscure`+"`"+` is set to `+"`"+`pass`+"`"+`, rclone will obscure the `+"`"+`pass`+"`"+`
parameter before creating the backend (which is required for sftp
backends).

The program can manipulate the supplied `+"`"+`user`+"`"+` in any way, for example
to make proxy to many different sftp backends, you could make the
`+"`"+`user`+"`"+` be `+"`"+`user@example.com`+"`"+` and then set the `+"`"+`host`+"`"+` to `+"`"+`example.com`+"`"+`
in the output and the user to `+"`"+`user`+"`"+`. For security you'd probably want
to restrict the `+"`"+`host`+"`"+` to a limited list.

Note that an internal cache is keyed on `+"`"+`user`+"`"+` so only use that for
configuration, don't use `+"`"+`pass`+"`"+` or `+"`"+`public_key`+"`"+`.  This also means that if a user's
password or public-key is changed the cache will need to expire (which takes 5 mins)
before it takes effect.

This can be used to build general purpose proxies to any kind of
backend that rclone supports.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `IPaddress:Port or :Port to bind server to`, ox.Default("localhost:2022")).
					String(`auth-proxy`, `A program to use to create the backend from the auth`).
					String(`authorized-keys`, `Authorized keys file`, ox.Default("~/.ssh/authorized_keys")).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Array(`key`, `SSH private host key file (Can be multi-valued, leave blank to auto generate)`).
					Bool(`no-auth`, `Allow connections with no authentication if set`).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					String(`pass`, `Password for authentication`).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					Bool(`read-only`, `Mount read-only`).
					Bool(`stdio`, `Run an sftp server on run stdin/stdout`).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					String(`user`, `User name for authentication`).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")),
			),
			ox.Sub(
				ox.Usage(`webdav`, `Serve remote:path over webdav.`),
				ox.Banner(`
rclone serve webdav implements a basic webdav server to serve the
remote over HTTP via the webdav protocol. This can be viewed with a
webdav client, through a web browser, or you can make a remote of
type webdav to read and write it.

### Webdav options

#### --etag-hash 

This controls the ETag header.  Without this flag the ETag will be
based on the ModTime and Size of the object.

If this flag is set to "auto" then rclone will choose the first
supported hash on the backend or you can use a named hash such as
"MD5" or "SHA-1".

Use "rclone hashsum" to see the full list.


### Server options

Use --addr to specify which IP address and port the server should
listen on, e.g. --addr 1.2.3.4:8000 or --addr :8080 to listen to all
IPs.  By default it only listens on localhost.  You can use port
:0 to let the OS choose an available port.

If you set --addr to listen on a public or LAN accessible IP address
then using Authentication is advised - see the next section for info.

--server-read-timeout and --server-write-timeout can be used to
control the timeouts on the server.  Note that this is the total time
for a transfer.

--max-header-bytes controls the maximum number of bytes the server will
accept in the HTTP header.

--baseurl controls the URL prefix that rclone serves from.  By default
rclone will serve from the root.  If you used --baseurl "/rclone" then
rclone would serve from a URL starting with "/rclone/".  This is
useful if you wish to proxy rclone serve.  Rclone automatically
inserts leading and trailing "/" on --baseurl, so --baseurl "rclone",
--baseurl "/rclone" and --baseurl "/rclone/" are all treated
identically.

--template allows a user to specify a custom markup template for http
and webdav serve functions.  The server exports the following markup
to be used within the template to server pages:

| Parameter   | Description |
| :---------- | :---------- |
| .Name       | The full path of a file/directory. |
| .Title      | Directory listing of .Name |
| .Sort       | The current sort used.  This is changeable via ?sort= parameter |
|             | Sort Options: namedirfirst,name,size,time (default namedirfirst) |
| .Order      | The current ordering used.  This is changeable via ?order= parameter |
|             | Order Options: asc,desc (default asc) |
| .Query      | Currently unused. |
| .Breadcrumb | Allows for creating a relative navigation |
|-- .Link     | The relative to the root link of the Text. |
|-- .Text     | The Name of the directory. |
| .Entries    | Information about a specific file/directory. |
|-- .URL      | The 'url' of an entry.  |
|-- .Leaf     | Currently same as 'URL' but intended to be 'just' the name. |
|-- .IsDir    | Boolean for if an entry is a directory or not. |
|-- .Size     | Size in Bytes of the entry. |
|-- .ModTime  | The UTC timestamp of an entry. |

#### Authentication

By default this will serve files without needing a login.

You can either use an htpasswd file which can take lots of users, or
set a single username and password with the --user and --pass flags.

Use --htpasswd /path/to/htpasswd to provide an htpasswd file.  This is
in standard apache format and supports MD5, SHA1 and BCrypt for basic
authentication.  Bcrypt is recommended.

To create an htpasswd file:

    touch htpasswd
    htpasswd -B htpasswd user
    htpasswd -B htpasswd anotherUser

The password file can be updated while rclone is running.

Use --realm to set the authentication realm.

#### SSL/TLS

By default this will serve over http.  If you want you can serve over
https.  You will need to supply the --cert and --key flags.  If you
wish to do client side certificate validation then you will need to
supply --client-ca also.

--cert should be either a PEM encoded certificate or a concatenation
of that with the CA certificate.  --key should be the PEM encoded
private key and --client-ca should be the PEM encoded client
certificate authority certificate.

### VFS - Virtual File System

This command uses the VFS layer. This adapts the cloud storage objects
that rclone uses into something which looks much more like a disk
filing system.

Cloud storage objects have lots of properties which aren't like disk
files - you can't extend them or write to the middle of them, so the
VFS layer has to deal with that. Because there is no one right way of
doing this there are various options explained below.

The VFS layer also implements a directory cache - this caches info
about files and directories (but not the data) in memory.

### VFS Directory Cache

Using the `+"`"+`--dir-cache-time`+"`"+` flag, you can control how long a
directory should be considered up to date and not refreshed from the
backend. Changes made through the mount will appear immediately or
invalidate the cache.

    --dir-cache-time duration   Time to cache directory entries for (default 5m0s)
    --poll-interval duration    Time to wait between polling for changes. Must be smaller than dir-cache-time. Only on supported remotes. Set to 0 to disable (default 1m0s)

However, changes made directly on the cloud storage by the web
interface or a different copy of rclone will only be picked up once
the directory cache expires if the backend configured does not support
polling for changes. If the backend supports polling, changes will be
picked up within the polling interval.

You can send a `+"`"+`SIGHUP`+"`"+` signal to rclone for it to flush all
directory caches, regardless of how old they are.  Assuming only one
rclone instance is running, you can reset the cache like this:

    kill -SIGHUP $(pidof rclone)

If you configure rclone with a [remote control](/rc) then you can use
rclone rc to flush the whole directory cache:

    rclone rc vfs/forget

Or individual files or directories:

    rclone rc vfs/forget file=path/to/file dir=path/to/dir

### VFS File Buffering

The `+"`"+`--buffer-size`+"`"+` flag determines the amount of memory,
that will be used to buffer data in advance.

Each open file will try to keep the specified amount of data in memory
at all times. The buffered data is bound to one open file and won't be
shared.

This flag is a upper limit for the used memory per open file.  The
buffer will only use memory for data that is downloaded but not not
yet read. If the buffer is empty, only a small amount of memory will
be used.

The maximum memory used by rclone for buffering can be up to
`+"`"+`--buffer-size * open files`+"`"+`.

### VFS File Caching

These flags control the VFS file caching options. File caching is
necessary to make the VFS layer appear compatible with a normal file
system. It can be disabled at the cost of some compatibility.

For example you'll need to enable VFS caching if you want to read and
write simultaneously to a file.  See below for more details.

Note that the VFS cache is separate from the cache backend and you may
find that you need one or the other or both.

    --cache-dir string                   Directory rclone will use for caching.
    --vfs-cache-mode CacheMode           Cache mode off|minimal|writes|full (default off)
    --vfs-cache-max-age duration         Max age of objects in the cache (default 1h0m0s)
    --vfs-cache-max-size SizeSuffix      Max total size of objects in the cache (default off)
    --vfs-cache-poll-interval duration   Interval to poll the cache for stale objects (default 1m0s)
    --vfs-write-back duration            Time to writeback files after last use when using cache (default 5s)

If run with `+"`"+`-vv`+"`"+` rclone will print the location of the file cache.  The
files are stored in the user cache file area which is OS dependent but
can be controlled with `+"`"+`--cache-dir`+"`"+` or setting the appropriate
environment variable.

The cache has 4 different modes selected by `+"`"+`--vfs-cache-mode`+"`"+`.
The higher the cache mode the more compatible rclone becomes at the
cost of using disk space.

Note that files are written back to the remote only when they are
closed and if they haven't been accessed for `+"`"+`--vfs-write-back`+"`"+`
seconds. If rclone is quit or dies with files that haven't been
uploaded, these will be uploaded next time rclone is run with the same
flags.

If using `+"`"+`--vfs-cache-max-size`+"`"+` note that the cache may exceed this size
for two reasons.  Firstly because it is only checked every
`+"`"+`--vfs-cache-poll-interval`+"`"+`.  Secondly because open files cannot be
evicted from the cache.

You **should not** run two copies of rclone using the same VFS cache
with the same or overlapping remotes if using `+"`"+`--vfs-cache-mode > off`+"`"+`.
This can potentially cause data corruption if you do. You can work
around this by giving each rclone its own cache hierarchy with
`+"`"+`--cache-dir`+"`"+`. You don't need to worry about this if the remotes in
use don't overlap.

#### --vfs-cache-mode off

In this mode (the default) the cache will read directly from the remote and write
directly to the remote without caching anything on disk.

This will mean some operations are not possible

  * Files can't be opened for both read AND write
  * Files opened for write can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files open for read with O_TRUNC will be opened write only
  * Files open for write only will behave as if O_TRUNC was supplied
  * Open modes O_APPEND, O_TRUNC are ignored
  * If an upload fails it can't be retried

#### --vfs-cache-mode minimal

This is very similar to "off" except that files opened for read AND
write will be buffered to disk.  This means that files opened for
write will be a lot more compatible, but uses the minimal disk space.

These operations are not possible

  * Files opened for write only can't be seeked
  * Existing files opened for write must have O_TRUNC set
  * Files opened for write only will ignore O_APPEND, O_TRUNC
  * If an upload fails it can't be retried

#### --vfs-cache-mode writes

In this mode files opened for read only are still read directly from
the remote, write only and read/write files are buffered to disk
first.

This mode should support all normal file system operations.

If an upload fails it will be retried at exponentially increasing
intervals up to 1 minute.

#### --vfs-cache-mode full

In this mode all reads and writes are buffered to and from disk. When
data is read from the remote this is buffered to disk as well.

In this mode the files in the cache will be sparse files and rclone
will keep track of which bits of the files it has downloaded.

So if an application only reads the starts of each file, then rclone
will only buffer the start of the file. These files will appear to be
their full size in the cache, but they will be sparse files with only
the data that has been downloaded present in them.

This mode should support all normal file system operations and is
otherwise identical to `+"`"+`--vfs-cache-mode`+"`"+` writes.

When reading a file rclone will read `+"`"+`--buffer-size`+"`"+` plus
`+"`"+`--vfs-read-ahead`+"`"+` bytes ahead.  The `+"`"+`--buffer-size`+"`"+` is buffered in memory
whereas the `+"`"+`--vfs-read-ahead`+"`"+` is buffered on disk.

When using this mode it is recommended that `+"`"+`--buffer-size`+"`"+` is not set
too large and `+"`"+`--vfs-read-ahead`+"`"+` is set large if required.

**IMPORTANT** not all file systems support sparse files. In particular
FAT/exFAT do not. Rclone will perform very badly if the cache
directory is on a filesystem which doesn't support sparse files and it
will log an ERROR message if one is detected.

#### Fingerprinting

Various parts of the VFS use fingerprinting to see if a local file
copy has changed relative to a remote file. Fingerprints are made
from:

- size
- modification time
- hash

where available on an object.

On some backends some of these attributes are slow to read (they take
an extra API call per object, or extra work per object).

For example `+"`"+`hash`+"`"+` is slow with the `+"`"+`local`+"`"+` and `+"`"+`sftp`+"`"+` backends as
they have to read the entire file and hash it, and `+"`"+`modtime`+"`"+` is slow
with the `+"`"+`s3`+"`"+`, `+"`"+`swift`+"`"+`, `+"`"+`ftp`+"`"+` and `+"`"+`qinqstor`+"`"+` backends because they
need to do an extra API call to fetch it.

If you use the `+"`"+`--vfs-fast-fingerprint`+"`"+` flag then rclone will not
include the slow operations in the fingerprint. This makes the
fingerprinting less accurate but much faster and will improve the
opening time of cached files.

If you are running a vfs cache over `+"`"+`local`+"`"+`, `+"`"+`s3`+"`"+` or `+"`"+`swift`+"`"+` backends
then using this flag is recommended.

Note that if you change the value of this flag, the fingerprints of
the files in the cache may be invalidated and the files will need to
be downloaded again.

### VFS Chunked Reading

When rclone reads files from a remote it reads them in chunks. This
means that rather than requesting the whole file rclone reads the
chunk specified.  This can reduce the used download quota for some
remotes by requesting only chunks from the remote that are actually
read, at the cost of an increased number of requests.

These flags control the chunking:

    --vfs-read-chunk-size SizeSuffix        Read the source objects in chunks (default 128M)
    --vfs-read-chunk-size-limit SizeSuffix  Max chunk doubling size (default off)

Rclone will start reading a chunk of size `+"`"+`--vfs-read-chunk-size`+"`"+`,
and then double the size for each read. When `+"`"+`--vfs-read-chunk-size-limit`+"`"+` is
specified, and greater than `+"`"+`--vfs-read-chunk-size`+"`"+`, the chunk size for each
open file will get doubled only until the specified value is reached. If the
value is "off", which is the default, the limit is disabled and the chunk size
will grow indefinitely.

With `+"`"+`--vfs-read-chunk-size 100M`+"`"+` and `+"`"+`--vfs-read-chunk-size-limit 0`+"`"+`
the following parts will be downloaded: 0-100M, 100M-200M, 200M-300M, 300M-400M and so on.
When `+"`"+`--vfs-read-chunk-size-limit 500M`+"`"+` is specified, the result would be
0-100M, 100M-300M, 300M-700M, 700M-1200M, 1200M-1700M and so on.

Setting `+"`"+`--vfs-read-chunk-size`+"`"+` to `+"`"+`0`+"`"+` or "off" disables chunked reading.

### VFS Performance

These flags may be used to enable/disable features of the VFS for
performance or other reasons. See also the [chunked reading](#vfs-chunked-reading)
feature.

In particular S3 and Swift benefit hugely from the `+"`"+`--no-modtime`+"`"+` flag
(or use `+"`"+`--use-server-modtime`+"`"+` for a slightly different effect) as each
read of the modification time takes a transaction.

    --no-checksum     Don't compare checksums on up/download.
    --no-modtime      Don't read/write the modification time (can speed things up).
    --no-seek         Don't allow seeking in files.
    --read-only       Mount read-only.

Sometimes rclone is delivered reads or writes out of order. Rather
than seeking rclone will wait a short time for the in sequence read or
write to come in. These flags only come into effect when not using an
on disk cache file.

    --vfs-read-wait duration   Time to wait for in-sequence read before seeking (default 20ms)
    --vfs-write-wait duration  Time to wait for in-sequence write before giving error (default 1s)

When using VFS write caching (`+"`"+`--vfs-cache-mode`+"`"+` with value writes or full),
the global flag `+"`"+`--transfers`+"`"+` can be set to adjust the number of parallel uploads of
modified files from cache (the related global flag `+"`"+`--checkers`+"`"+` have no effect on mount).

    --transfers int  Number of file transfers to run in parallel (default 4)

### VFS Case Sensitivity

Linux file systems are case-sensitive: two files can differ only
by case, and the exact case must be used when opening a file.

File systems in modern Windows are case-insensitive but case-preserving:
although existing files can be opened using any case, the exact case used
to create the file is preserved and available for programs to query.
It is not allowed for two files in the same directory to differ only by case.

Usually file systems on macOS are case-insensitive. It is possible to make macOS
file systems case-sensitive but that is not the default.

The `+"`"+`--vfs-case-insensitive`+"`"+` mount flag controls how rclone handles these
two cases. If its value is "false", rclone passes file names to the mounted
file system as-is. If the flag is "true" (or appears without a value on
command line), rclone may perform a "fixup" as explained below.

The user may specify a file name to open/delete/rename/etc with a case
different than what is stored on mounted file system. If an argument refers
to an existing file with exactly the same name, then the case of the existing
file on the disk will be used. However, if a file name with exactly the same
name is not found but a name differing only by case exists, rclone will
transparently fixup the name. This fixup happens only when an existing file
is requested. Case sensitivity of file names created anew by rclone is
controlled by an underlying mounted file system.

Note that case sensitivity of the operating system running rclone (the target)
may differ from case sensitivity of a file system mounted by rclone (the source).
The flag controls whether "fixup" is performed to satisfy the target.

If the flag is not provided on the command line, then its default value depends
on the operating system where rclone runs: "true" on Windows and macOS, "false"
otherwise. If the flag is provided without a value, then it is "true".

### Alternate report of used bytes

Some backends, most notably S3, do not report the amount of bytes used.
If you need this information to be available when running `+"`"+`df`+"`"+` on the
filesystem, then pass the flag `+"`"+`--vfs-used-is-size`+"`"+` to rclone.
With this flag set, instead of relying on the backend to report this
information, rclone will scan the whole remote similar to `+"`"+`rclone size`+"`"+`
and compute the total used space itself.

_WARNING._ Contrary to `+"`"+`rclone size`+"`"+`, this flag ignores filters so that the
result is accurate. However, this is very inefficient and may cost lots of API
calls resulting in extra charges. Use it as a last resort and only with caching.

### Auth Proxy

If you supply the parameter `+"`"+`--auth-proxy /path/to/program`+"`"+` then
rclone will use that program to generate backends on the fly which
then are used to authenticate incoming requests.  This uses a simple
JSON based protocol with input on STDIN and output on STDOUT.

**PLEASE NOTE:** `+"`"+`--auth-proxy`+"`"+` and `+"`"+`--authorized-keys`+"`"+` cannot be used
together, if `+"`"+`--auth-proxy`+"`"+` is set the authorized keys option will be
ignored.

There is an example program
[bin/test_proxy.py](https://github.com/rclone/rclone/blob/master/test_proxy.py)
in the rclone source code.

The program's job is to take a `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+` on the input and turn
those into the config for a backend on STDOUT in JSON format.  This
config will have any default parameters for the backend added, but it
won't use configuration from environment variables or command line
options - it is the job of the proxy program to make a complete
config.

This config generated must have this extra parameter
- `+"`"+`_root`+"`"+` - root to use for the backend

And it may have this parameter
- `+"`"+`_obscure`+"`"+` - comma separated strings for parameters to obscure

If password authentication was used by the client, input to the proxy
process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"pass": "mypassword"
}
`+"`"+``+"`"+``+"`"+`

If public-key authentication was used by the client, input to the
proxy process (on STDIN) would look similar to this:

`+"`"+``+"`"+``+"`"+`
{
	"user": "me",
	"public_key": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDuwESFdAe14hVS6omeyX7edc...JQdf"
}
`+"`"+``+"`"+``+"`"+`

And as an example return this on STDOUT

`+"`"+``+"`"+``+"`"+`
{
	"type": "sftp",
	"_root": "",
	"_obscure": "pass",
	"user": "me",
	"pass": "mypassword",
	"host": "sftp.example.com"
}
`+"`"+``+"`"+``+"`"+`

This would mean that an SFTP backend would be created on the fly for
the `+"`"+`user`+"`"+` and `+"`"+`pass`+"`"+`/`+"`"+`public_key`+"`"+` returned in the output to the host given.  Note
that since `+"`"+`_obscure`+"`"+` is set to `+"`"+`pass`+"`"+`, rclone will obscure the `+"`"+`pass`+"`"+`
parameter before creating the backend (which is required for sftp
backends).

The program can manipulate the supplied `+"`"+`user`+"`"+` in any way, for example
to make proxy to many different sftp backends, you could make the
`+"`"+`user`+"`"+` be `+"`"+`user@example.com`+"`"+` and then set the `+"`"+`host`+"`"+` to `+"`"+`example.com`+"`"+`
in the output and the user to `+"`"+`user`+"`"+`. For security you'd probably want
to restrict the `+"`"+`host`+"`"+` to a limited list.

Note that an internal cache is keyed on `+"`"+`user`+"`"+` so only use that for
configuration, don't use `+"`"+`pass`+"`"+` or `+"`"+`public_key`+"`"+`.  This also means that if a user's
password or public-key is changed the cache will need to expire (which takes 5 mins)
before it takes effect.

This can be used to build general purpose proxies to any kind of
backend that rclone supports.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					String(`addr`, `IPaddress:Port or :Port to bind server to`, ox.Default("localhost:8080")).
					String(`auth-proxy`, `A program to use to create the backend from the auth`).
					String(`baseurl`, `Prefix for URLs - leave blank for root`).
					String(`cert`, `SSL PEM key (concatenation of certificate and CA certificate)`).
					String(`client-ca`, `Client certificate authority to verify clients with`).
					Duration(`dir-cache-time`, `Time to cache directory entries for`, ox.Default("5m0s")).
					String(`dir-perms`, `Directory permissions`, ox.Spec(`FileMode`), ox.Default("0777")).
					Bool(`disable-dir-list`, `Disable HTML directory list on GET request for a directory`).
					String(`etag-hash`, `Which hash to use for the ETag, or auto or blank for off`).
					String(`file-perms`, `File permissions`, ox.Spec(`FileMode`), ox.Default("0666")).
					Uint32(`gid`, `Override the gid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					String(`htpasswd`, `htpasswd file - if not provided no authentication is done`).
					String(`key`, `SSL PEM Private key`).
					Int(`max-header-bytes`, `Maximum size of request header`, ox.Default("4096")).
					Bool(`no-checksum`, `Don't compare checksums on up/download`).
					Bool(`no-modtime`, `Don't read/write the modification time (can speed things up)`).
					Bool(`no-seek`, `Don't allow seeking in files`).
					String(`pass`, `Password for authentication`).
					Duration(`poll-interval`, `Time to wait between polling for changes, must be smaller than dir-cache-time and only on supported remotes (set 0 to disable)`, ox.Default("1m0s")).
					Bool(`read-only`, `Mount read-only`).
					String(`realm`, `Realm for authentication`, ox.Default("$APPNAME")).
					Duration(`server-read-timeout`, `Timeout for server reading data`, ox.Default("1h0m0s")).
					Duration(`server-write-timeout`, `Timeout for server writing data`, ox.Default("1h0m0s")).
					String(`template`, `User-specified template`).
					Uint32(`uid`, `Override the uid field set by the filesystem (not supported on Windows)`, ox.Default("1000")).
					Int(`umask`, `Override the permission bits set by the filesystem (not supported on Windows)`, ox.Default("18")).
					String(`user`, `User name for authentication`).
					Duration(`vfs-cache-max-age`, `Max age of objects in the cache`, ox.Default("1h0m0s")).
					String(`vfs-cache-max-size`, `Max total size of objects in the cache`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					String(`vfs-cache-mode`, `Cache mode off|minimal|writes|full`, ox.Spec(`CacheMode`), ox.Default("off")).
					Duration(`vfs-cache-poll-interval`, `Interval to poll the cache for stale objects`, ox.Default("1m0s")).
					Bool(`vfs-case-insensitive`, `If a file name not found, find a case insensitive match`).
					Bool(`vfs-fast-fingerprint`, `Use fast (less accurate) fingerprints for change detection`).
					String(`vfs-read-ahead`, `Extra read ahead over --buffer-size when using cache-mode full`, ox.Spec(`SizeSuffix`)).
					String(`vfs-read-chunk-size`, `Read the source objects in chunks`, ox.Spec(`SizeSuffix`), ox.Default("128Mi")).
					String(`vfs-read-chunk-size-limit`, `If greater than --vfs-read-chunk-size, double the chunk size after each chunk read, until the limit is reached ('off' is unlimited)`, ox.Spec(`SizeSuffix`), ox.Default("off")).
					Duration(`vfs-read-wait`, `Time to wait for in-sequence read before seeking`, ox.Default("20ms")).
					String(`vfs-used-is-size`, `size           Use the rclone size algorithm for Used size`, ox.Spec(`rclone`)).
					Duration(`vfs-write-back`, `Time to writeback files after last use when using cache`, ox.Default("5s")).
					Duration(`vfs-write-wait`, `Time to wait for in-sequence write before giving error`, ox.Default("1s")),
			),
		),
		ox.Sub(
			ox.Usage(`settier`, `Changes storage class/tier of objects in remote.`),
			ox.Banner(`
rclone settier changes storage tier or class at remote if supported.
Few cloud storage services provides different storage classes on objects,
for example AWS S3 and Glacier, Azure Blob storage - Hot, Cool and Archive,
Google Cloud Storage, Regional Storage, Nearline, Coldline etc.

Note that, certain tier changes make objects not available to access immediately.
For example tiering to archive in azure blob storage makes objects in frozen state,
user can restore by setting tier to Hot/Cool, similarly S3 to Glacier makes object
inaccessible.true

You can use it to tier single object

    rclone settier Cool remote:path/file

Or use rclone filters to set tier on only specific files

	rclone --include "*.txt" settier Hot remote:path/dir

Or just provide remote directory and all files in directory will be tiered

    rclone settier tier remote:path/dir`),
			ox.Spec(`tier remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
		),
		ox.Sub(
			ox.Usage(`sha1sum`, `Produces an sha1sum file for all the objects in the path.`),
			ox.Banner(`
Produces an sha1sum file for all the objects in the path.  This
is in the same format as the standard sha1sum tool produces.

By default, the hash is requested from the remote.  If SHA-1 is
not supported by the remote, no hash will be returned.  With the
download flag, the file will be downloaded from the remote and
hashed locally enabling SHA-1 for any remote.

This command can also hash data received on standard input (stdin),
by not passing a remote:path, or by passing a hyphen as remote:path
when there is data to read (if not, the hypen will be treated literaly,
as a relative path).

This command can also hash data received on STDIN, if not passing
a remote:path.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`base64`, `Output base64 encoded hashsum`).
				String(`checkfile`, `Validate hashes against a given SUM file instead of printing them`, ox.Short("C")).
				Bool(`download`, `Download the file and hash it locally; if this flag is not specified, the hash is requested from the remote`).
				String(`output-file`, `Output hashsums to a file rather than the terminal`),
		),
		ox.Sub(
			ox.Usage(`size`, `Prints the total size and number of objects in remote:path.`),
			ox.Banner(`Prints the total size and number of objects in remote:path.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`json`, `Format output as JSON`),
		),
		ox.Sub(
			ox.Usage(`sync`, `Make source and dest identical, modifying destination only.`),
			ox.Banner(`
Sync the source to the destination, changing the destination
only.  Doesn't transfer files that are identical on source and
destination, testing by size and modification time or MD5SUM.
Destination is updated to match source, including deleting files
if necessary (except duplicate objects, see below).

**Important**: Since this can cause data loss, test first with the
`+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+`/`+"`"+`-i`+"`"+` flag.

    rclone sync -i SOURCE remote:DESTINATION

Note that files in the destination won't be deleted if there were any
errors at any point.  Duplicate objects (files with the same name, on
those providers that support it) are also not yet handled.

It is always the contents of the directory that is synced, not the
directory so when source:path is a directory, it's the contents of
source:path that are copied, not the directory name and contents.  See
extended explanation in the `+"`"+`copy`+"`"+` command above if unsure.

If dest:path doesn't exist, it is created and the source:path contents
go there.

**Note**: Use the `+"`"+`-P`+"`"+`/`+"`"+`--progress`+"`"+` flag to view real-time transfer statistics

**Note**: Use the `+"`"+`rclone dedupe`+"`"+` command to deal with "Duplicate object/directory found in source/destination - ignoring" errors.
See [this forum post](https://forum.rclone.org/t/sync-not-clearing-duplicates/14372) for more info.`),
			ox.Spec(`source:path dest:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`create-empty-src-dirs`, `Create empty source dirs on destination after sync`),
		),
		ox.Sub(
			ox.Usage(`test`, `Run a test command`),
			ox.Banner(`Rclone test is used to run test commands.

Select which test comand you want with the subcommand, eg

    rclone test memory remote:

Each subcommand has its own options which you can see in their help.

**NB** Be careful running these commands, they may do strange things
so reading their documentation first is recommended.`),
			ox.Spec(`[command]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Sub(
				ox.Usage(`changenotify`, `Log any change notify requests for the remote passed in.`),
				ox.Banner(`Log any change notify requests for the remote passed in.`),
				ox.Spec(`remote: [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Duration(`poll-interval`, `Time to wait between polling for changes`, ox.Default("10s")),
			),
			ox.Sub(
				ox.Usage(`histogram`, `Makes a histogram of file name characters.`),
				ox.Banner(`This command outputs JSON which shows the histogram of characters used
in filenames in the remote:path specified.

The data doesn't contain any identifying information but is useful for
the rclone developers when developing filename compression.`),
				ox.Spec(`[remote:path] [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
			ox.Sub(
				ox.Usage(`info`, `Discovers file name or other limitations for paths.`),
				ox.Banner(`rclone info discovers what filenames and upload methods are possible
to write to the paths passed in and how long they can be.  It can take some
time.  It will write test files into the remote:path passed in.  It outputs
a bit of go code for each one.

**NB** this can create undeletable files and other hazards - use with care`),
				ox.Spec(`[remote:path]+ [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Bool(`all`, `Run all tests`).
					Bool(`check-control`, `Check control characters`).
					Bool(`check-length`, `Check max filename length`).
					Bool(`check-normalization`, `Check UTF-8 Normalization`).
					Bool(`check-streaming`, `Check uploads with indeterminate file size`).
					Duration(`upload-wait`, `Wait after writing a file`).
					String(`write-json`, `Write results to file`),
			),
			ox.Sub(
				ox.Usage(`makefiles`, `Make a random file hierarchy in a directory`),
				ox.Banner(`Make a random file hierarchy in a directory`),
				ox.Spec(`<dir> [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
				ox.Flags().
					Int(`files`, `Number of files to create`, ox.Default("1000")).
					Int(`files-per-directory`, `Average number of files per directory`, ox.Default("10")).
					String(`max-file-size`, `Maximum size of files to create`, ox.Spec(`SizeSuffix`), ox.Default("100")).
					Int(`max-name-length`, `Maximum size of file names`, ox.Default("12")).
					String(`min-file-size`, `Minimum size of file to create`, ox.Spec(`SizeSuffix`)).
					Int(`min-name-length`, `Minimum size of file names`, ox.Default("4")).
					Int(`seed`, `Seed for the random number generator (0 for random)`, ox.Default("1")),
			),
			ox.Sub(
				ox.Usage(`memory`, `Load all the objects at remote:path into memory and report memory stats.`),
				ox.Banner(`Load all the objects at remote:path into memory and report memory stats.`),
				ox.Spec(`remote:path [flags]`),
				ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			),
		),
		ox.Sub(
			ox.Usage(`touch`, `Create new file or change file modification time.`),
			ox.Banner(`
Set the modification time on file(s) as specified by remote:path to
have the current time.

If remote:path does not exist then a zero sized file will be created,
unless `+"`"+`--no-create`+"`"+` or `+"`"+`--recursive`+"`"+` is provided.

If `+"`"+`--recursive`+"`"+` is used then recursively sets the modification
time on all existing files that is found under the path. Filters are supported,
and you can test with the `+"`"+`--dry-run`+"`"+` or the `+"`"+`--interactive`+"`"+` flag.

If `+"`"+`--timestamp`+"`"+` is used then sets the modification time to that
time instead of the current time. Times may be specified as one of:

- 'YYMMDD' - e.g. 17.10.30
- 'YYYY-MM-DDTHH:MM:SS' - e.g. 2006-01-02T15:04:05
- 'YYYY-MM-DDTHH:MM:SS.SSS' - e.g. 2006-01-02T15:04:05.123456789

Note that value of `+"`"+`--timestamp`+"`"+` is in UTC. If you want local time
then add the `+"`"+`--localtime`+"`"+` flag.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`localtime`, `Use localtime for timestamp, not UTC`).
				Bool(`no-create`, `Do not create the file if it does not exist (implied with --recursive)`, ox.Short("C")).
				Bool(`recursive`, `Recursively touch all files`, ox.Short("R")).
				String(`timestamp`, `Use specified time instead of the current time of day`, ox.Short("t")),
		),
		ox.Sub(
			ox.Usage(`tree`, `List the contents of the remote in a tree like fashion.`),
			ox.Banner(`
rclone tree lists the contents of a remote in a similar way to the
unix tree command.

For example

    $ rclone tree remote:path
    /
    ├── file1
    ├── file2
    ├── file3
    └── subdir
        ├── file4
        └── file5

    1 directories, 5 files

You can use any of the filtering options with the tree command (e.g.
--include and --exclude).  You can also use --fast-list.

The tree command has many options for controlling the listing which
are compatible with the tree command.  Note that not all of them have
short options as they conflict with rclone's short options.`),
			ox.Spec(`remote:path [flags]`),
			ox.Footer(`Use "rclone [command] --help" for more information about a command.
Use "rclone help flags" for to see the global flags.
Use "rclone help backends" for a list of supported services.`),
			ox.Flags().
				Bool(`all`, `All files are listed (list . files too)`, ox.Short("a")).
				Bool(`color`, `Turn colorization on always`, ox.Short("C")).
				Bool(`dirs-only`, `List directories only`, ox.Short("d")).
				Bool(`dirsfirst`, `List directories before files (-U disables)`).
				Bool(`full-path`, `Print the full path prefix for each file`).
				Int(`level`, `Descend only level directories deep`).
				Bool(`modtime`, `Print the date of last modification.`, ox.Short("D")).
				Bool(`noindent`, `Don't print indentation lines`).
				Bool(`noreport`, `Turn off file/directory count at end of tree listing`).
				String(`output`, `Output to file instead of stdout`, ox.Short("o")).
				Bool(`protections`, `Print the protections for each file.`, ox.Short("p")).
				Bool(`quote`, `Quote filenames with double quotes.`, ox.Short("Q")).
				Bool(`size`, `Print the size in bytes of each file.`, ox.Short("s")).
				String(`sort`, `Select sort: name,version,size,mtime,ctime`).
				Bool(`sort-ctime`, `Sort files by last status change time`).
				Bool(`sort-modtime`, `Sort files by last modification time`, ox.Short("t")).
				Bool(`sort-reverse`, `Reverse the order of the sort`, ox.Short("r")).
				Bool(`unsorted`, `Leave files unsorted`, ox.Short("U")),
		),
	)
}
