// Command gh is a xo/ox version of `gh`.
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
		ox.Usage(`gh`, ``),
		ox.Banner(`Work seamlessly with GitHub from the command line.`),
		ox.Spec(`<command> <subcommand> [flags]`),
		ox.Example(`
  $ gh issue create
  $ gh repo clone cli/cli
  $ gh pr checkout 321`),
		ox.Sections("CORE COMMANDS", "GITHUB ACTIONS COMMANDS", "ALIAS COMMANDS", "ADDITIONAL COMMANDS"),
		ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
		ox.Sub(
			ox.Usage(`auth`, `Authenticate gh and git with GitHub`),
			ox.Banner(`Authenticate gh and git with GitHub`),
			ox.Spec(`<command> [flags]`),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`login`, `Log in to a GitHub account`),
				ox.Banner(`Authenticate with a GitHub host.

The default hostname is `+"`"+`github.com`+"`"+`. This can be overridden using the `+"`"+`--hostname`+"`"+`
flag.

The default authentication mode is a web-based browser flow. After completion, an
authentication token will be stored securely in the system credential store.
If a credential store is not found or there is an issue using it gh will fallback
to writing the token to a plain text file. See `+"`"+`gh auth status`+"`"+` for its
stored location.

Alternatively, use `+"`"+`--with-token`+"`"+` to pass in a personal access token (classic) on standard input.
The minimum required scopes for the token are: `+"`"+`repo`+"`"+`, `+"`"+`read:org`+"`"+`, and `+"`"+`gist`+"`"+`.
Take care when passing a fine-grained personal access token to `+"`"+`--with-token`+"`"+`
as the inherent scoping to certain resources may cause confusing behaviour when interacting with other
resources. Favour setting `+"`"+`GH_TOKEN`+"`"+` for fine-grained personal access token usage.

Alternatively, gh will use the authentication token found in environment variables.
This method is most suitable for "headless" use of gh such as in automation. See
`+"`"+`gh help environment`+"`"+` for more info.

To use gh in GitHub Actions, add `+"`"+`GH_TOKEN: ${{ github.token }}`+"`"+` to `+"`"+`env`+"`"+`.

The git protocol to use for git operations on this host can be set with `+"`"+`--git-protocol`+"`"+`,
or during the interactive prompting. Although login is for a single account on a host, setting
the git protocol will take effect for all users on the host.

Specifying `+"`"+`ssh`+"`"+` for the git protocol will detect existing SSH keys to upload,
prompting to create and upload a new key if one is not found. This can be skipped with
`+"`"+`--skip-ssh-key`+"`"+` flag.

For more information on OAuth scopes, see
<https://docs.github.com/en/developers/apps/building-oauth-apps/scopes-for-oauth-apps/>.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Start interactive setup
  $ gh auth login
  
  # Authenticate against <github.com> by reading the token from a file
  $ gh auth login --with-token < mytoken.txt
  
  # Authenticate with specific host
  $ gh auth login --hostname enterprise.internal`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`git-protocol`, `The protocol to use for git operations on this host: {ssh|https}`, ox.Short("p"), ox.Section(0)).
					String(`hostname`, `The hostname of the GitHub instance to authenticate with`, ox.Short("h"), ox.Section(0)).
					Bool(`insecure-storage`, `Save authentication credentials in plain text instead of credential store`, ox.Section(0)).
					Slice(`scopes`, `Additional authentication scopes to request`, ox.Short("s"), ox.Section(0)).
					Bool(`skip-ssh-key`, `Skip generate/upload SSH key prompt`, ox.Section(0)).
					Bool(`web`, `Open a browser to authenticate`, ox.Short("w"), ox.Section(0)).
					Bool(`with-token`, `Read token from standard input`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`logout`, `Log out of a GitHub account`),
				ox.Banner(`Remove authentication for a GitHub account.

This command removes the stored authentication configuration
for an account. The authentication configuration is only
removed locally.

This command does not revoke authentication tokens.

To revoke all authentication tokens generated by the GitHub CLI:

1. Visit <https://github.com/settings/applications>
2. Select the "GitHub CLI" application
3. Select "Revoke Access"
4. Select "I understand, revoke access"

Note: this procedure will revoke all authentication tokens ever
generated by the GitHub CLI across all your devices.

For more information about revoking OAuth application tokens, see:
<https://docs.github.com/en/apps/oauth-apps/using-oauth-apps/reviewing-your-authorized-oauth-apps>`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Select what host and account to log out of via a prompt
  $ gh auth logout
  
  # Log out of a specific host and specific account
  $ gh auth logout --hostname enterprise.internal --user monalisa`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`hostname`, `The hostname of the GitHub instance to log out of`, ox.Short("h"), ox.Section(0)).
					String(`user`, `The account to log out of`, ox.Short("u"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`refresh`, `Refresh stored authentication credentials`),
				ox.Banner(`Expand or fix the permission scopes for stored credentials for active account.

The `+"`"+`--scopes`+"`"+` flag accepts a comma separated list of scopes you want
your gh credentials to have. If no scopes are provided, the command
maintains previously added scopes.

The `+"`"+`--remove-scopes`+"`"+` flag accepts a comma separated list of scopes you
want to remove from your gh credentials. Scope removal is idempotent.
The minimum set of scopes (`+"`"+`repo`+"`"+`, `+"`"+`read:org`+"`"+`, and `+"`"+`gist`+"`"+`) cannot be removed.

The `+"`"+`--reset-scopes`+"`"+` flag resets the scopes for your gh credentials to
the default set of scopes for your auth flow.

If you have multiple accounts in `+"`"+`gh auth status`+"`"+` and want to refresh the credentials for an
inactive account, you will have to use `+"`"+`gh auth switch`+"`"+` to that account first before using
this command, and then switch back when you are done.

For more information on OAuth scopes, see
<https://docs.github.com/en/developers/apps/building-oauth-apps/scopes-for-oauth-apps/>.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Open a browser to add write:org and read:public_key scopes
  $ gh auth refresh --scopes write:org,read:public_key
  
  # Open a browser to ensure your authentication credentials have the correct minimum scopes
  $ gh auth refresh
  
  # Open a browser to idempotently remove the delete_repo scope
  $ gh auth refresh --remove-scopes delete_repo
  
  # Open a browser to re-authenticate with the default minimum scopes
  $ gh auth refresh --reset-scopes`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`hostname`, `The GitHub host to use for authentication`, ox.Short("h"), ox.Section(0)).
					Bool(`insecure-storage`, `Save authentication credentials in plain text instead of credential store`, ox.Section(0)).
					Slice(`remove-scopes`, `Authentication scopes to remove from gh`, ox.Short("r"), ox.Section(0)).
					Bool(`reset-scopes`, `Reset authentication scopes to the default minimum set of scopes`, ox.Section(0)).
					Slice(`scopes`, `Additional authentication scopes for gh to have`, ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`setup-git`, `Setup git with GitHub CLI`),
				ox.Banner(`This command configures `+"`"+`git`+"`"+` to use GitHub CLI as a credential helper.
For more information on git credential helpers please reference:
<https://git-scm.com/docs/gitcredentials>.

By default, GitHub CLI will be set as the credential helper for all authenticated hosts.
If there is no authenticated hosts the command fails with an error.

Alternatively, use the `+"`"+`--hostname`+"`"+` flag to specify a single host to be configured.
If the host is not authenticated with, the command fails with an error.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Configure git to use GitHub CLI as the credential helper for all authenticated hosts
  $ gh auth setup-git
  
  # Configure git to use GitHub CLI as the credential helper for enterprise.internal host
  $ gh auth setup-git --hostname enterprise.internal`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`force`, `Force setup even if the host is not known. Must be used in conjunction with --hostname`, ox.Spec(`--hostname`), ox.Short("f"), ox.Section(0)).
					String(`hostname`, `The hostname to configure git for`, ox.Short("h"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`status`, `Display active account and authentication state on each known GitHub host`),
				ox.Banner(`Display active account and authentication state on each known GitHub host.

For each host, the authentication state of each known account is tested and any issues are included in the output.
Each host section will indicate the active account, which will be used when targeting that host.
If an account on any host (or only the one given via `+"`"+`--hostname`+"`"+`) has authentication issues,
the command will exit with 1 and output to stderr.

To change the active account for a host, see `+"`"+`gh auth switch`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`active`, `Display the active account only`, ox.Short("a"), ox.Section(0)).
					String(`hostname`, `Check only a specific hostname's auth status`, ox.Short("h"), ox.Section(0)).
					Bool(`show-token`, `Display the auth token`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`switch`, `Switch active GitHub account`),
				ox.Banner(`Switch the active account for a GitHub host.

This command changes the authentication configuration that will
be used when running commands targeting the specified GitHub host.

If the specified host has two accounts, the active account will be switched
automatically. If there are more than two accounts, disambiguation will be
required either through the `+"`"+`--user`+"`"+` flag or an interactive prompt.

For a list of authenticated accounts you can run `+"`"+`gh auth status`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Select what host and account to switch to via a prompt
  $ gh auth switch
  
  # Switch the active account on a specific host to a specific user
  $ gh auth switch --hostname enterprise.internal --user monalisa`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`hostname`, `The hostname of the GitHub instance to switch account for`, ox.Short("h"), ox.Section(0)).
					String(`user`, `The account to switch to`, ox.Short("u"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`token`, `Print the authentication token gh uses for a hostname and account`),
				ox.Banner(`This command outputs the authentication token for an account on a given GitHub host.

Without the `+"`"+`--hostname`+"`"+` flag, the default host is chosen.

Without the `+"`"+`--user`+"`"+` flag, the active account for the host is chosen.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`hostname`, `The hostname of the GitHub instance authenticated with`, ox.Short("h"), ox.Section(0)).
					String(`user`, `The account to output the token for`, ox.Short("u"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`browse`, `Open repositories, issues, pull requests, and more in the browser`),
			ox.Banner(`Transition from the terminal to the web browser to view and interact with:

- Issues
- Pull requests
- Repository content
- Repository home page
- Repository settings`),
			ox.Spec(`[<number> | <path> | <commit-sha>] [flags]`),
			ox.Example(`
  # Open the home page of the current repository
  $ gh browse
  
  # Open the script directory of the current repository
  $ gh browse script/
  
  # Open issue or pull request 217
  $ gh browse 217
  
  # Open commit page
  $ gh browse 77507cd94ccafcf568f8560cfecde965fcfa63
  
  # Open repository settings
  $ gh browse --settings
  
  # Open main.go at line 312
  $ gh browse main.go:312
  
  # Open main.go with the repository at head of bug-fix branch
  $ gh browse main.go --branch bug-fix
  
  # Open main.go with the repository at commit 775007cd
  $ gh browse main.go --commit=77507cd94ccafcf568f8560cfecde965fcfa63`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`branch`, `Select another branch by passing in the branch name`, ox.Short("b"), ox.Section(0)).
				Map(`commit`, `Select another commit by passing in the commit SHA, default is the last commit`, ox.Spec(`string[="last"]`), ox.Short("c"), ox.Section(0)).
				Bool(`no-browser`, `Print destination URL instead of opening the browser`, ox.Short("n"), ox.Section(0)).
				Bool(`projects`, `Open repository projects`, ox.Short("p"), ox.Section(0)).
				Bool(`releases`, `Open repository releases`, ox.Short("r"), ox.Section(0)).
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)).
				Bool(`settings`, `Open repository settings`, ox.Short("s"), ox.Section(0)).
				Bool(`wiki`, `Open repository wiki`, ox.Short("w"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`codespace`, `Connect to and manage codespaces`),
			ox.Banner(`Connect to and manage codespaces`),
			ox.Spec(`[flags]`),
			ox.Aliases("cs"),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`code`, `Open a codespace in Visual Studio Code`),
				ox.Banner(`Open a codespace in Visual Studio Code`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`insiders`, `Use the insiders version of Visual Studio Code`, ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					Bool(`web`, `Use the web version of Visual Studio Code`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`cp`, `Copy files between local and remote file systems`),
				ox.Banner(`The `+"`"+`cp`+"`"+` command copies files between the local and remote file systems.

As with the UNIX `+"`"+`cp`+"`"+` command, the first argument specifies the source and the last
specifies the destination; additional sources may be specified after the first,
if the destination is a directory.

The `+"`"+`--recursive`+"`"+` flag is required if any source is a directory.

A `+"`"+`remote:`+"`"+` prefix on any file name argument indicates that it refers to
the file system of the remote (Codespace) machine. It is resolved relative
to the home directory of the remote user.

By default, remote file names are interpreted literally. With the `+"`"+`--expand`+"`"+` flag,
each such argument is treated in the manner of `+"`"+`scp`+"`"+`, as a Bash expression to
be evaluated on the remote machine, subject to expansion of tildes, braces, globs,
environment variables, and backticks. For security, do not use this flag with arguments
provided by untrusted users; see <https://lwn.net/Articles/835962/> for discussion.

By default, the `+"`"+`cp`+"`"+` command will create a public/private ssh key pair to authenticate with
the codespace inside the `+"`"+`~/.ssh directory`+"`"+`.`),
				ox.Spec(`[-e] [-r] [-- [<scp flags>...]] <sources>... <dest>`),
				ox.Example(`
  $ gh codespace cp -e README.md 'remote:/workspaces/$RepositoryName/'
  $ gh codespace cp -e 'remote:~/*.go' ./gofiles/
  $ gh codespace cp -e 'remote:/workspaces/myproj/go.{mod,sum}' ./gofiles/
  $ gh codespace cp -e -- -F ~/.ssh/codespaces_config 'remote:~/*.go' ./gofiles/`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`expand`, `Expand remote file names on remote shell`, ox.Short("e"), ox.Section(0)).
					String(`profile`, `Name of the SSH profile to use`, ox.Short("p"), ox.Section(0)).
					Bool(`recursive`, `Recursively copy directories`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`create`, `Create a codespace`),
				ox.Banner(`Create a codespace`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`branch`, `Repository branch`, ox.Short("b"), ox.Section(0)).
					Bool(`default-permissions`, `Do not prompt to accept additional permissions requested by the codespace`, ox.Section(0)).
					String(`devcontainer-path`, `Path to the devcontainer.json file to use when creating codespace`, ox.Section(0)).
					String(`display-name`, `Display name for the codespace (48 characters or less)`, ox.Short("d"), ox.Section(0)).
					Duration(`idle-timeout`, `Allowed inactivity before codespace is stopped, e.g. "10m", "1h"`, ox.Section(0)).
					String(`location`, `Location: {EastUs|SouthEastAsia|WestEurope|WestUs2} (determined automatically if not provided)`, ox.Short("l"), ox.Section(0)).
					String(`machine`, `Hardware specifications for the VM`, ox.Short("m"), ox.Section(0)).
					String(`repo`, `Repository name with owner: user/repo`, ox.Short("R"), ox.Section(0)).
					Duration(`retention-period`, `Allowed time after shutting down before the codespace is automatically deleted (maximum 30 days), e.g. "1h", "72h"`, ox.Section(0)).
					Bool(`status`, `Show status of post-create command and dotfiles`, ox.Short("s"), ox.Section(0)).
					Bool(`web`, `Create codespace from browser, cannot be used with --display-name, --idle-timeout, or --retention-period`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete codespaces`),
				ox.Banner(`Delete codespaces based on selection criteria.

All codespaces for the authenticated user can be deleted, as well as codespaces for a
specific repository. Alternatively, only codespaces older than N days can be deleted.

Organization administrators may delete any codespace billed to the organization.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Delete all codespaces`, ox.Section(0)).
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`days`, `Delete codespaces older than N days`, ox.Default("N"), ox.Section(0)).
					Bool(`force`, `Skip confirmation for codespaces that contain unsaved changes`, ox.Short("f"), ox.Section(0)).
					String(`org`, `The login handle of the organization (admin-only)`, ox.Spec(`login`), ox.Short("o"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					String(`user`, `The username to delete codespaces for (used with --org)`, ox.Spec(`username`), ox.Short("u"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit a codespace`),
				ox.Banner(`Edit a codespace`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					String(`display-name`, `Set the display name`, ox.Short("d"), ox.Section(0)).
					String(`machine`, `Set hardware specifications for the VM`, ox.Short("m"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`jupyter`, `Open a codespace in JupyterLab`),
				ox.Banner(`Open a codespace in JupyterLab`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List codespaces`),
				ox.Banner(`List codespaces of the authenticated user.

Alternatively, organization administrators may list all codespaces billed to the organization.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("codespace ls", "cs ls"),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of codespaces to list`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`org`, `The login handle of the organization to list codespaces for (admin-only)`, ox.Spec(`login`), ox.Short("o"), ox.Section(0)).
					String(`repo`, `Repository name with owner: user/repo`, ox.Short("R"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`user`, `The username to list codespaces for (used with --org)`, ox.Spec(`username`), ox.Short("u"), ox.Section(0)).
					Bool(`web`, `List codespaces in the web browser, cannot be used with --user or --org`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`logs`, `Access codespace logs`),
				ox.Banner(`Access codespace logs`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`follow`, `Tail and follow the logs`, ox.Short("f"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`ports`, `List ports in a codespace`),
				ox.Banner(`List ports in a codespace

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
				ox.Sub(
					ox.Usage(`forward`, `Forward ports`),
					ox.Banner(`Forward ports`),
					ox.Spec(`<remote-port>:<local-port>... [flags]`),
					ox.Help(ox.Sections(
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
						String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
						String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
				),
				ox.Sub(
					ox.Usage(`visibility`, `Change the visibility of the forwarded port`),
					ox.Banner(`Change the visibility of the forwarded port`),
					ox.Spec(`<port>:{public|private|org}... [flags]`),
					ox.Example(`
  $ gh codespace ports visibility 80:org 3000:private 8000:public`),
					ox.Help(ox.Sections(
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
						String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
						String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
				),
			),
			ox.Sub(
				ox.Usage(`rebuild`, `Rebuild a codespace`),
				ox.Banner(`Rebuilding recreates your codespace.

Your code and any current changes will be preserved. Your codespace will be rebuilt using
your working directory's dev container. A full rebuild also removes cached Docker images.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`full`, `Perform a full rebuild`, ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`ssh`, `SSH into a codespace`),
				ox.Banner(`The `+"`"+`ssh`+"`"+` command is used to SSH into a codespace. In its simplest form, you can
run `+"`"+`gh cs ssh`+"`"+`, select a codespace interactively, and connect.

The `+"`"+`ssh`+"`"+` command will automatically create a public/private ssh key pair in the
`+"`"+`~/.ssh`+"`"+` directory if you do not have an existing valid key pair. When selecting the
key pair to use, the preferred order is:

1. Key specified by `+"`"+`-i`+"`"+` in `+"`"+`<ssh-flags>`+"`"+`
2. Automatic key, if it already exists
3. First valid key pair in ssh config (according to `+"`"+`ssh -G`+"`"+`)
4. Automatic key, newly created

The `+"`"+`ssh`+"`"+` command also supports deeper integration with OpenSSH using a `+"`"+`--config`+"`"+`
option that generates per-codespace ssh configuration in OpenSSH format.
Including this configuration in your `+"`"+`~/.ssh/config`+"`"+` improves the user experience
of tools that integrate with OpenSSH, such as Bash/Zsh completion of ssh hostnames,
remote path completion for `+"`"+`scp/rsync/sshfs`+"`"+`, `+"`"+`git`+"`"+` ssh remotes, and so on.

Once that is set up (see the second example below), you can ssh to codespaces as
if they were ordinary remote hosts (using `+"`"+`ssh`+"`"+`, not `+"`"+`gh cs ssh`+"`"+`).

Note that the codespace you are connecting to must have an SSH server pre-installed.
If the docker image being used for the codespace does not have an SSH server,
install it in your `+"`"+`Dockerfile`+"`"+` or, for codespaces that use Debian-based images,
you can add the following to your `+"`"+`devcontainer.json`+"`"+`:

	"features": {
		"ghcr.io/devcontainers/features/sshd:1": {
			"version": "latest"
		}
	}`),
				ox.Spec(`[<flags>...] [-- <ssh-flags>...] [<command>]`),
				ox.Example(`
  $ gh codespace ssh
  
  $ gh codespace ssh --config > ~/.ssh/codespaces
  $ printf 'Match all\nInclude ~/.ssh/codespaces\n' >> ~/.ssh/config`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					Bool(`config`, `Write OpenSSH configuration to stdout`, ox.Section(0)).
					Bool(`debug`, `Log debug data to a file`, ox.Short("d"), ox.Section(0)).
					String(`debug-file`, `Path of the file log to`, ox.Section(0)).
					String(`profile`, `Name of the SSH profile to use`, ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					Int(`server-port`, `SSH server port number (0 => pick unused)`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`stop`, `Stop a running codespace`),
				ox.Banner(`Stop a running codespace`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					String(`org`, `The login handle of the organization (admin-only)`, ox.Spec(`login`), ox.Short("o"), ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					String(`user`, `The username to stop codespace for (used with --org)`, ox.Spec(`username`), ox.Short("u"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`view`, `View details about a codespace`),
				ox.Banner(`View details about a codespace

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Select a codespace from a list of all codespaces you own
  $ gh cs view
  
  # View the details of a specific codespace
  $ gh cs view -c codespace-name-12345
  
  # View the list of all available fields for a codespace
  $ gh cs view --json
  
  # View specific fields for a codespace
  $ gh cs view --json displayName,machineDisplayName,state`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`codespace`, `Name of the codespace`, ox.Short("c"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`repo`, `Filter codespace selection by repository name (user/repo)`, ox.Short("R"), ox.Section(0)).
					String(`repo-owner`, `Filter codespace selection by repository owner (username or org)`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`gist`, `Manage gists`),
			ox.Banner(`Work with GitHub gists.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`clone`, `Clone a gist locally`),
				ox.Banner(`Clone a GitHub gist locally.

A gist can be supplied as argument in either of the following formats:
- by ID, e.g. `+"`"+`5b0e0062eb8e9654adad7bb1d81cc75f`+"`"+`
- by URL, e.g. `+"`"+`https://gist.github.com/OWNER/5b0e0062eb8e9654adad7bb1d81cc75f`+"`"+`

Pass additional `+"`"+`git clone`+"`"+` flags by listing them after `+"`"+`--`+"`"+`.`),
				ox.Spec(`<gist> [<directory>] [-- <gitflags>...]`),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`create`, `Create a new gist`),
				ox.Banner(`Create a new GitHub gist with given contents.

Gists can be created from one or multiple files. Alternatively, pass `+"`"+`-`+"`"+` as
filename to read from standard input.

By default, gists are secret; use `+"`"+`--public`+"`"+` to make publicly listed ones.`),
				ox.Spec(`[<filename>... | <pattern>... | -] [flags]`),
				ox.Aliases("gist new"),
				ox.Example(`
  # Publish file 'hello.py' as a public gist
  $ gh gist create --public hello.py
  
  # Create a gist with a description
  $ gh gist create hello.py -d "my Hello-World program in Python"
  
  # Create a gist containing several files
  $ gh gist create hello.py world.py cool.txt
  
  # Create a gist containing several files using patterns
  $ gh gist create *.md *.txt artifact.*
  
  # Read from standard input to create a gist
  $ gh gist create -
  
  # Create a gist from output piped from another command
  $ cat cool.txt | gh gist create`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`desc`, `A description for this gist`, ox.Short("d"), ox.Section(0)).
					String(`filename`, `Provide a filename to be used when reading from standard input`, ox.Short("f"), ox.Section(0)).
					Bool(`public`, `List the gist publicly`, ox.Default("secret"), ox.Short("p"), ox.Section(0)).
					Bool(`web`, `Open the web browser with created gist`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a gist`),
				ox.Banner(`Delete a GitHub gist.

To delete a gist interactively, use `+"`"+`gh gist delete`+"`"+` with no arguments.

To delete a gist non-interactively, supply the gist id or url.`),
				ox.Spec(`{<id> | <url>} [flags]`),
				ox.Example(`
  # Delete a gist interactively
  $ gh gist delete
  
  # Delete a gist non-interactively
  $ gh gist delete 1234`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Confirm deletion without prompting`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit one of your gists`),
				ox.Banner(`Edit one of your gists`),
				ox.Spec(`{<id> | <url>} [<filename>] [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`add`, `Add a new file to the gist`, ox.Short("a"), ox.Section(0)).
					String(`desc`, `New description for the gist`, ox.Short("d"), ox.Section(0)).
					String(`filename`, `Select a file to edit`, ox.Short("f"), ox.Section(0)).
					String(`remove`, `Remove a file from the gist`, ox.Short("r"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List your gists`),
				ox.Banner(`List gists from your user account.

You can use a regular expression to filter the description, file names,
or even the content of files in the gist using `+"`"+`--filter`+"`"+`.

For supported regular expression syntax, see <https://pkg.go.dev/regexp/syntax>.

Use `+"`"+`--include-content`+"`"+` to include content of files, noting that
this will be slower and increase the rate limit used. Instead of printing a table,
code will be printed with highlights similar to `+"`"+`gh search code`+"`"+`:

	{{gist ID}} {{file name}}
	    {{description}}
	        {{matching lines from content}}

No highlights or other color is printed when output is redirected.`),
				ox.Spec(`[flags]`),
				ox.Aliases("gist ls"),
				ox.Example(`
  # List all secret gists from your user account
  $ gh gist list --secret
  
  # Find all gists from your user account mentioning "octo" anywhere
  $ gh gist list --filter octo --include-content`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`filter`, `Filter gists using a regular expression`, ox.Spec(`expression`), ox.Section(0)).
					Bool(`include-content`, `Include gists' file content when filtering`, ox.Section(0)).
					Int(`limit`, `Maximum number of gists to fetch`, ox.Default("10"), ox.Short("L"), ox.Section(0)).
					Bool(`public`, `Show only public gists`, ox.Section(0)).
					Bool(`secret`, `Show only secret gists`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`rename`, `Rename a file in a gist`),
				ox.Banner(`Rename a file in the given gist ID / URL.`),
				ox.Spec(`{<id> | <url>} <old-filename> <new-filename> [flags]`),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`view`, `View a gist`),
				ox.Banner(`View the given gist or select from recent gists.`),
				ox.Spec(`[<id> | <url>] [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`filename`, `Display a single file from the gist`, ox.Short("f"), ox.Section(0)).
					Bool(`files`, `List file names from the gist`, ox.Section(0)).
					Bool(`raw`, `Print raw instead of rendered gist contents`, ox.Short("r"), ox.Section(0)).
					Bool(`web`, `Open gist in the browser`, ox.Short("w"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`issue`, `Manage issues`),
			ox.Banner(`Work with GitHub issues.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh issue list
  $ gh issue create --label bug
  $ gh issue view 123 --web`),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`create`, `Create a new issue`),
				ox.Banner(`Create an issue on GitHub.

Adding an issue to projects requires authorization with the `+"`"+`project`+"`"+` scope.
To authorize, run `+"`"+`gh auth refresh -s project`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("issue new"),
				ox.Example(`
  $ gh issue create --title "I found a bug" --body "Nothing works"
  $ gh issue create --label "bug,help wanted"
  $ gh issue create --label bug --label "help wanted"
  $ gh issue create --assignee monalisa,hubot
  $ gh issue create --assignee "@me"
  $ gh issue create --project "Roadmap"
  $ gh issue create --template "Bug Report"`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`assignee`, `Assign people by their login. Use "@me" to self-assign.`, ox.Spec(`login`), ox.Short("a"), ox.Section(0)).
					String(`body`, `Supply a body. Will prompt for one otherwise.`, ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`editor`, `Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.`, ox.Short("e"), ox.Section(0)).
					String(`label`, `Add labels by name`, ox.Spec(`name`), ox.Short("l"), ox.Section(0)).
					String(`milestone`, `Add the issue to a milestone by name`, ox.Spec(`name`), ox.Short("m"), ox.Section(0)).
					String(`project`, `Add the issue to projects by title`, ox.Spec(`title`), ox.Short("p"), ox.Section(0)).
					String(`recover`, `Recover input from a failed run of create`, ox.Section(0)).
					String(`template`, `Template name to use as starting body text`, ox.Spec(`name`), ox.Short("T"), ox.Section(0)).
					String(`title`, `Supply a title. Will prompt for one otherwise.`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open the browser to create an issue`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List issues in a repository`),
				ox.Banner(`List issues in a GitHub repository. By default, this only lists open issues.

The search query syntax is documented here:
<https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("issue ls"),
				ox.Example(`
  $ gh issue list --label "bug" --label "help wanted"
  $ gh issue list --author monalisa
  $ gh issue list --assignee "@me"
  $ gh issue list --milestone "The big 1.0"
  $ gh issue list --search "error no:assignee sort:created-asc"
  $ gh issue list --state all`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Filter by GitHub App author`, ox.Section(0)).
					String(`assignee`, `Filter by assignee`, ox.Short("a"), ox.Section(0)).
					String(`author`, `Filter by author`, ox.Short("A"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Slice(`label`, `Filter by label`, ox.Short("l"), ox.Section(0)).
					Int(`limit`, `Maximum number of issues to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`mention`, `Filter by mention`, ox.Section(0)).
					String(`milestone`, `Filter by milestone number or title`, ox.Short("m"), ox.Section(0)).
					String(`search`, `Search issues with query`, ox.Spec(`query`), ox.Short("S"), ox.Section(0)).
					String(`state`, `Filter by state: {open|closed|all}`, ox.Default("open"), ox.Short("s"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `List issues in the web browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`status`, `Show status of relevant issues`),
				ox.Banner(`Show status of relevant issues

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`close`, `Close issue`),
				ox.Banner(`Close issue`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`comment`, `Leave a closing comment`, ox.Short("c"), ox.Section(0)).
					String(`reason`, `Reason for closing: {completed|not planned}`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`comment`, `Add a comment to an issue`),
				ox.Banner(`Add a comment to a GitHub issue.

Without the body text supplied through flags, the command will interactively
prompt for the comment text.`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Example(`
  $ gh issue comment 12 --body "Hi from GitHub CLI"`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`body`, `The comment body text`, ox.Default("text"), ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`create-if-none`, `Create a new comment if no comments are found. Can be used only with --edit-last`, ox.Section(0)).
					Bool(`delete-last`, `Delete the last comment of the current user`, ox.Section(0)).
					Bool(`edit-last`, `Edit the last comment of the current user`, ox.Section(0)).
					Bool(`editor`, `Skip prompts and open the text editor to write the body in`, ox.Short("e"), ox.Section(0)).
					Bool(`web`, `Open the web browser to write the comment`, ox.Short("w"), ox.Section(0)).
					Bool(`yes`, `Skip the delete confirmation prompt when --delete-last is provided`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete issue`),
				ox.Banner(`Delete issue`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Confirm deletion without prompting`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`develop`, `Manage linked branches for an issue`),
				ox.Banner(`Manage linked branches for an issue.

When using the `+"`"+`--base`+"`"+` flag, the new development branch will be created from the specified
remote branch. The new branch will be configured as the base branch for pull requests created using
`+"`"+`gh pr create`+"`"+`.`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Example(`
  # List branches for issue 123
  $ gh issue develop --list 123
  
  # List branches for issue 123 in repo cli/cli
  $ gh issue develop --list --repo cli/cli 123
  
  # Create a branch for issue 123 based on the my-feature branch
  $ gh issue develop 123 --base my-feature
  
  # Create a branch for issue 123 and checkout it out
  $ gh issue develop 123 --checkout
  
  # Create a branch in repo monalisa/cli for issue 123 in repo cli/cli
  $ gh issue develop 123 --repo cli/cli --branch-repo monalisa/cli`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`base`, `Name of the remote branch you want to make your new branch from`, ox.Short("b"), ox.Section(0)).
					String(`branch-repo`, `Name or URL of the repository where you want to create your new branch`, ox.Section(0)).
					Bool(`checkout`, `Checkout the branch after creating it`, ox.Short("c"), ox.Section(0)).
					Bool(`list`, `List linked branches for the issue`, ox.Short("l"), ox.Section(0)).
					String(`name`, `Name of the branch to create`, ox.Short("n"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit issues`),
				ox.Banner(`Edit one or more issues within the same repository.

Editing issues' projects requires authorization with the `+"`"+`project`+"`"+` scope.
To authorize, run `+"`"+`gh auth refresh -s project`+"`"+`.

The `+"`"+`--add-assignee`+"`"+` and `+"`"+`--remove-assignee`+"`"+` flags both support
the following special values:
- `+"`"+`@me`+"`"+`: assign or unassign yourself
- `+"`"+`@copilot`+"`"+`: assign or unassign Copilot (not supported on GitHub Enterprise Server)`),
				ox.Spec(`{<numbers> | <urls>} [flags]`),
				ox.Example(`
  $ gh issue edit 23 --title "I found a bug" --body "Nothing works"
  $ gh issue edit 23 --add-label "bug,help wanted" --remove-label "core"
  $ gh issue edit 23 --add-assignee "@me" --remove-assignee monalisa,hubot
  $ gh issue edit 23 --add-assignee "@copilot"
  $ gh issue edit 23 --add-project "Roadmap" --remove-project v1,v2
  $ gh issue edit 23 --milestone "Version 1"
  $ gh issue edit 23 --remove-milestone
  $ gh issue edit 23 --body-file body.txt
  $ gh issue edit 23 34 --add-label "help wanted"`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`add-assignee`, `Add assigned users by their login. Use "@me" to assign yourself.`, ox.Spec(`login`), ox.Section(0)).
					String(`add-label`, `Add labels by name`, ox.Spec(`name`), ox.Section(0)).
					String(`add-project`, `Add the issue to projects by title`, ox.Spec(`title`), ox.Section(0)).
					String(`body`, `Set the new body.`, ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					String(`milestone`, `Edit the milestone the issue belongs to by name`, ox.Spec(`name`), ox.Short("m"), ox.Section(0)).
					String(`remove-assignee`, `Remove assigned users by their login. Use "@me" to unassign yourself.`, ox.Spec(`login`), ox.Section(0)).
					String(`remove-label`, `Remove labels by name`, ox.Spec(`name`), ox.Section(0)).
					Bool(`remove-milestone`, `Remove the milestone association from the issue`, ox.Section(0)).
					String(`remove-project`, `Remove the issue from projects by title`, ox.Spec(`title`), ox.Section(0)).
					String(`title`, `Set the new title.`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`lock`, `Lock issue conversation`),
				ox.Banner(`Lock issue conversation`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`reason`, `Optional reason for locking conversation (off_topic, resolved, spam, too_heated).`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`pin`, `Pin a issue`),
				ox.Banner(`Pin an issue to a repository.

The issue can be specified by issue number or URL.`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Example(`
  # Pin an issue to the current repository
  $ gh issue pin 23
  
  # Pin an issue by URL
  $ gh issue pin https://github.com/owner/repo/issues/23
  
  # Pin an issue to specific repository
  $ gh issue pin 23 --repo owner/repo`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`reopen`, `Reopen issue`),
				ox.Banner(`Reopen issue`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`comment`, `Add a reopening comment`, ox.Short("c"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`transfer`, `Transfer issue to another repository`),
				ox.Banner(`Transfer issue to another repository`),
				ox.Spec(`{<number> | <url>} <destination-repo> [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`unlock`, `Unlock issue conversation`),
				ox.Banner(`Unlock issue conversation`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`unpin`, `Unpin a issue`),
				ox.Banner(`Unpin an issue from a repository.

The issue can be specified by issue number or URL.`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Example(`
  # Unpin issue from the current repository
  $ gh issue unpin 23
  
  # Unpin issue by URL
  $ gh issue unpin https://github.com/owner/repo/issues/23
  
  # Unpin an issue from specific repository
  $ gh issue unpin 23 --repo owner/repo`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`view`, `View an issue`),
				ox.Banner(`Display the title, body, and other information about an issue.

With `+"`"+`--web`+"`"+` flag, open the issue in a web browser instead.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`comments`, `View issue comments`, ox.Short("c"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open an issue in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`org`, `Manage organizations`),
			ox.Banner(`Work with GitHub organizations.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh org list`),
			ox.Sections("GENERAL COMMANDS"),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`list`, `List organizations for the authenticated user.`),
				ox.Banner(`List organizations for the authenticated user.`),
				ox.Spec(`[flags]`),
				ox.Aliases("org ls"),
				ox.Example(`
  # List the first 30 organizations
  $ gh org list
  
  # List more organizations
  $ gh org list --limit 100`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Int(`limit`, `Maximum number of organizations to list`, ox.Default("30"), ox.Short("L"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`pr`, `Manage pull requests`),
			ox.Banner(`Work with GitHub pull requests.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh pr checkout 353
  $ gh pr create --fill
  $ gh pr view --web`),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`create`, `Create a pull request`),
				ox.Banner(`Create a pull request on GitHub.

When the current branch isn't fully pushed to a git remote, a prompt will ask where
to push the branch and offer an option to fork the base repository. Use `+"`"+`--head`+"`"+` to
explicitly skip any forking or pushing behavior.

`+"`"+`--head`+"`"+` supports `+"`"+`<user>:<branch>`+"`"+` syntax to select a head repo owned by `+"`"+`<user>`+"`"+`.
Using an organization as the `+"`"+`<user>`+"`"+` is currently not supported.
For more information, see <https://github.com/cli/cli/issues/10093>

A prompt will also ask for the title and the body of the pull request. Use `+"`"+`--title`+"`"+` and
`+"`"+`--body`+"`"+` to skip this, or use `+"`"+`--fill`+"`"+` to autofill these values from git commits.
It's important to notice that if the `+"`"+`--title`+"`"+` and/or `+"`"+`--body`+"`"+` are also provided
alongside `+"`"+`--fill`+"`"+`, the values specified by `+"`"+`--title`+"`"+` and/or `+"`"+`--body`+"`"+` will
take precedence and overwrite any autofilled content.

The base branch for the created PR can be specified using the `+"`"+`--base`+"`"+` flag. If not provided,
the value of `+"`"+`gh-merge-base`+"`"+` git branch config will be used. If not configured, the repository's
default branch will be used. Run `+"`"+`git config branch.{current}.gh-merge-base {base}`+"`"+` to configure
the current branch to use the specified merge base.

Link an issue to the pull request by referencing the issue in the body of the pull
request. If the body text mentions `+"`"+`Fixes #123`+"`"+` or `+"`"+`Closes #123`+"`"+`, the referenced issue
will automatically get closed when the pull request gets merged.

By default, users with write access to the base repository can push new commits to the
head branch of the pull request. Disable this with `+"`"+`--no-maintainer-edit`+"`"+`.

Adding a pull request to projects requires authorization with the `+"`"+`project`+"`"+` scope.
To authorize, run `+"`"+`gh auth refresh -s project`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("pr new"),
				ox.Example(`
  $ gh pr create --title "The bug is fixed" --body "Everything works again"
  $ gh pr create --reviewer monalisa,hubot  --reviewer myorg/team-name
  $ gh pr create --project "Roadmap"
  $ gh pr create --base develop --head monalisa:feature
  $ gh pr create --template "pull_request_template.md"`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`assignee`, `Assign people by their login. Use "@me" to self-assign.`, ox.Spec(`login`), ox.Short("a"), ox.Section(0)).
					String(`base`, `The branch into which you want your code merged`, ox.Spec(`branch`), ox.Short("B"), ox.Section(0)).
					String(`body`, `Body for the pull request`, ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`draft`, `Mark pull request as a draft`, ox.Short("d"), ox.Section(0)).
					Bool(`dry-run`, `Print details instead of creating the PR. May still push git changes.`, ox.Section(0)).
					Bool(`editor`, `Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.`, ox.Short("e"), ox.Section(0)).
					Bool(`fill`, `Use commit info for title and body`, ox.Short("f"), ox.Section(0)).
					Bool(`fill-first`, `Use first commit info for title and body`, ox.Section(0)).
					Bool(`fill-verbose`, `Use commits msg+body for description`, ox.Section(0)).
					String(`head`, `The branch that contains commits for your pull request`, ox.Spec(`branch`), ox.Default("[current branch]"), ox.Short("H"), ox.Section(0)).
					String(`label`, `Add labels by name`, ox.Spec(`name`), ox.Short("l"), ox.Section(0)).
					String(`milestone`, `Add the pull request to a milestone by name`, ox.Spec(`name`), ox.Short("m"), ox.Section(0)).
					Bool(`no-maintainer-edit`, `Disable maintainer's ability to modify pull request`, ox.Section(0)).
					String(`project`, `Add the pull request to projects by title`, ox.Spec(`title`), ox.Short("p"), ox.Section(0)).
					String(`recover`, `Recover input from a failed run of create`, ox.Section(0)).
					String(`reviewer`, `Request reviews from people or teams by their handle`, ox.Spec(`handle`), ox.Short("r"), ox.Section(0)).
					String(`template`, `Template file to use as starting body text`, ox.Spec(`file`), ox.Short("T"), ox.Section(0)).
					String(`title`, `Title for the pull request`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open the web browser to create a pull request`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List pull requests in a repository`),
				ox.Banner(`List pull requests in a GitHub repository. By default, this only lists open PRs.

The search query syntax is documented here:
<https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("pr ls"),
				ox.Example(`
  # List PRs authored by you
  $ gh pr list --author "@me"
  
  # List only PRs with all of the given labels
  $ gh pr list --label bug --label "priority 1"
  
  # Filter PRs using search syntax
  $ gh pr list --search "status:success review:required"
  
  # Find a PR that introduced a given commit
  $ gh pr list --search "<SHA>" --state merged`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Filter by GitHub App author`, ox.Section(0)).
					String(`assignee`, `Filter by assignee`, ox.Short("a"), ox.Section(0)).
					String(`author`, `Filter by author`, ox.Short("A"), ox.Section(0)).
					String(`base`, `Filter by base branch`, ox.Short("B"), ox.Section(0)).
					Bool(`draft`, `Filter by draft state`, ox.Short("d"), ox.Section(0)).
					String(`head`, `Filter by head branch`, ox.Short("H"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Slice(`label`, `Filter by label`, ox.Short("l"), ox.Section(0)).
					Int(`limit`, `Maximum number of items to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`search`, `Search pull requests with query`, ox.Spec(`query`), ox.Short("S"), ox.Section(0)).
					String(`state`, `Filter by state: {open|closed|merged|all}`, ox.Default("open"), ox.Short("s"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `List pull requests in the web browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`status`, `Show status of relevant pull requests`),
				ox.Banner(`Show status of relevant pull requests.

The status shows a summary of pull requests that includes information such as
pull request number, title, CI checks, reviews, etc.

To see more details of CI checks, run `+"`"+`gh pr checks`+"`"+`.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`conflict-status`, `Display the merge conflict status of each pull request`, ox.Short("c"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`checkout`, `Check out a pull request in git`),
				ox.Banner(`Check out a pull request in git`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Example(`
  # Interactively select a PR from the 10 most recent to check out
  $ gh pr checkout
  
  # Checkout a specific PR
  $ gh pr checkout 32
  $ gh pr checkout https://github.com/OWNER/REPO/pull/32
  $ gh pr checkout feature`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`branch`, `Local branch name to use`, ox.Default("[the name of the head branch]"), ox.Short("b"), ox.Section(0)).
					Bool(`detach`, `Checkout PR with a detached HEAD`, ox.Section(0)).
					Bool(`force`, `Reset the existing local branch to the latest state of the pull request`, ox.Short("f"), ox.Section(0)).
					Bool(`recurse-submodules`, `Update all submodules after checkout`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`checks`, `Show CI status for a single pull request`),
				ox.Banner(`Show CI status for a single pull request.

Without an argument, the pull request that belongs to the current branch
is selected.

When the `+"`"+`--json`+"`"+` flag is used, it includes a `+"`"+`bucket`+"`"+` field, which categorizes
the `+"`"+`state`+"`"+` field into `+"`"+`pass`+"`"+`, `+"`"+`fail`+"`"+`, `+"`"+`pending`+"`"+`, `+"`"+`skipping`+"`"+`, or `+"`"+`cancel`+"`"+`.

Additional exit codes:
	8: Checks pending

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`fail-fast`, `Exit watch mode on first check failure`, ox.Section(0)).
					String(`interval`, `Refresh interval in seconds when using --watch flag`, ox.Spec(`--watch`), ox.Default("10"), ox.Short("i"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Bool(`required`, `Only show checks that are required`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`watch`, `Watch checks until they finish`, ox.Section(0)).
					Bool(`web`, `Open the web browser to show details about checks`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`close`, `Close a pull request`),
				ox.Banner(`Close a pull request`),
				ox.Spec(`{<number> | <url> | <branch>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`comment`, `Leave a closing comment`, ox.Short("c"), ox.Section(0)).
					Bool(`delete-branch`, `Delete the local and remote branch after close`, ox.Short("d"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`comment`, `Add a comment to a pull request`),
				ox.Banner(`Add a comment to a GitHub pull request.

Without the body text supplied through flags, the command will interactively
prompt for the comment text.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Example(`
  $ gh pr comment 13 --body "Hi from GitHub CLI"`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`body`, `The comment body text`, ox.Default("text"), ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`create-if-none`, `Create a new comment if no comments are found. Can be used only with --edit-last`, ox.Section(0)).
					Bool(`delete-last`, `Delete the last comment of the current user`, ox.Section(0)).
					Bool(`edit-last`, `Edit the last comment of the current user`, ox.Section(0)).
					Bool(`editor`, `Skip prompts and open the text editor to write the body in`, ox.Short("e"), ox.Section(0)).
					Bool(`web`, `Open the web browser to write the comment`, ox.Short("w"), ox.Section(0)).
					Bool(`yes`, `Skip the delete confirmation prompt when --delete-last is provided`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`diff`, `View changes in a pull request`),
				ox.Banner(`View changes in a pull request.

Without an argument, the pull request that belongs to the current branch
is selected.

With `+"`"+`--web`+"`"+` flag, open the pull request diff in a web browser instead.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`color`, `Use color in diff output: {always|never|auto}`, ox.Default("auto"), ox.Section(0)).
					Bool(`name-only`, `Display only names of changed files`, ox.Section(0)).
					Bool(`patch`, `Display diff in patch format`, ox.Section(0)).
					Bool(`web`, `Open the pull request diff in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit a pull request`),
				ox.Banner(`Edit a pull request.

Without an argument, the pull request that belongs to the current branch
is selected.

Editing a pull request's projects requires authorization with the `+"`"+`project`+"`"+` scope.
To authorize, run `+"`"+`gh auth refresh -s project`+"`"+`.

The `+"`"+`--add-assignee`+"`"+` and `+"`"+`--remove-assignee`+"`"+` flags both support
the following special values:
- `+"`"+`@me`+"`"+`: assign or unassign yourself
- `+"`"+`@copilot`+"`"+`: assign or unassign Copilot (not supported on GitHub Enterprise Server)

The `+"`"+`--add-reviewer`+"`"+` and `+"`"+`--remove-reviewer`+"`"+` flags do not support
these special values.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Example(`
  $ gh pr edit 23 --title "I found a bug" --body "Nothing works"
  $ gh pr edit 23 --add-label "bug,help wanted" --remove-label "core"
  $ gh pr edit 23 --add-reviewer monalisa,hubot  --remove-reviewer myorg/team-name
  $ gh pr edit 23 --add-assignee "@me" --remove-assignee monalisa,hubot
  $ gh pr edit 23 --add-assignee "@copilot"
  $ gh pr edit 23 --add-project "Roadmap" --remove-project v1,v2
  $ gh pr edit 23 --milestone "Version 1"
  $ gh pr edit 23 --remove-milestone`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`add-assignee`, `Add assigned users by their login. Use "@me" to assign yourself.`, ox.Spec(`login`), ox.Section(0)).
					String(`add-label`, `Add labels by name`, ox.Spec(`name`), ox.Section(0)).
					String(`add-project`, `Add the pull request to projects by title`, ox.Spec(`title`), ox.Section(0)).
					String(`add-reviewer`, `Add reviewers by their login.`, ox.Spec(`login`), ox.Section(0)).
					String(`base`, `Change the base branch for this pull request`, ox.Spec(`branch`), ox.Short("B"), ox.Section(0)).
					String(`body`, `Set the new body.`, ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					String(`milestone`, `Edit the milestone the pull request belongs to by name`, ox.Spec(`name`), ox.Short("m"), ox.Section(0)).
					String(`remove-assignee`, `Remove assigned users by their login. Use "@me" to unassign yourself.`, ox.Spec(`login`), ox.Section(0)).
					String(`remove-label`, `Remove labels by name`, ox.Spec(`name`), ox.Section(0)).
					Bool(`remove-milestone`, `Remove the milestone association from the pull request`, ox.Section(0)).
					String(`remove-project`, `Remove the pull request from projects by title`, ox.Spec(`title`), ox.Section(0)).
					String(`remove-reviewer`, `Remove reviewers by their login.`, ox.Spec(`login`), ox.Section(0)).
					String(`title`, `Set the new title.`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`lock`, `Lock pull request conversation`),
				ox.Banner(`Lock pull request conversation`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`reason`, `Optional reason for locking conversation (off_topic, resolved, spam, too_heated).`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`merge`, `Merge a pull request`),
				ox.Banner(`Merge a pull request on GitHub.

Without an argument, the pull request that belongs to the current branch
is selected.

When targeting a branch that requires a merge queue, no merge strategy is required.
If required checks have not yet passed, auto-merge will be enabled.
If required checks have passed, the pull request will be added to the merge queue.
To bypass a merge queue and merge directly, pass the `+"`"+`--admin`+"`"+` flag.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`admin`, `Use administrator privileges to merge a pull request that does not meet requirements`, ox.Section(0)).
					String(`author-email`, `Email text for merge commit author`, ox.Default("text"), ox.Short("A"), ox.Section(0)).
					Bool(`auto`, `Automatically merge only after necessary requirements are met`, ox.Section(0)).
					String(`body`, `Body text for the merge commit`, ox.Default("text"), ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`delete-branch`, `Delete the local and remote branch after merge`, ox.Short("d"), ox.Section(0)).
					Bool(`disable-auto`, `Disable auto-merge for this pull request`, ox.Section(0)).
					String(`match-head-commit`, `Commit SHA that the pull request head must match to allow merge`, ox.Spec(`SHA`), ox.Section(0)).
					Bool(`merge`, `Merge the commits with the base branch`, ox.Short("m"), ox.Section(0)).
					Bool(`rebase`, `Rebase the commits onto the base branch`, ox.Short("r"), ox.Section(0)).
					Bool(`squash`, `Squash the commits into one commit and merge it into the base branch`, ox.Short("s"), ox.Section(0)).
					String(`subject`, `Subject text for the merge commit`, ox.Default("text"), ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`ready`, `Mark a pull request as ready for review`),
				ox.Banner(`Mark a pull request as ready for review.

Without an argument, the pull request that belongs to the current branch
is marked as ready.

If supported by your plan, convert to draft with `+"`"+`--undo`+"`"+``),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`undo`, `Convert a pull request to "draft"`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`reopen`, `Reopen a pull request`),
				ox.Banner(`Reopen a pull request`),
				ox.Spec(`{<number> | <url> | <branch>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`comment`, `Add a reopening comment`, ox.Short("c"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`review`, `Add a review to a pull request`),
				ox.Banner(`Add a review to a pull request.

Without an argument, the pull request that belongs to the current branch is reviewed.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Example(`
  # Approve the pull request of the current branch
  $ gh pr review --approve
  
  # Leave a review comment for the current branch
  $ gh pr review --comment -b "interesting"
  
  # Add a review for a specific pull request
  $ gh pr review 123
  
  # Request changes on a specific pull request
  $ gh pr review 123 -r -b "needs more ASCII art"`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`approve`, `Approve pull request`, ox.Short("a"), ox.Section(0)).
					String(`body`, `Specify the body of a review`, ox.Short("b"), ox.Section(0)).
					String(`body-file`, `Read body text from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`comment`, `Comment on a pull request`, ox.Short("c"), ox.Section(0)).
					Bool(`request-changes`, `Request changes on a pull request`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`unlock`, `Unlock pull request conversation`),
				ox.Banner(`Unlock pull request conversation`),
				ox.Spec(`{<number> | <url>} [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`update-branch`, `Update a pull request branch`),
				ox.Banner(`Update a pull request branch with latest changes of the base branch.

Without an argument, the pull request that belongs to the current branch is selected.

The default behavior is to update with a merge commit (i.e., merging the base branch
into the PR's branch). To reconcile the changes with rebasing on top of the base
branch, the `+"`"+`--rebase`+"`"+` option should be provided.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Example(`
  $ gh pr update-branch 23
  $ gh pr update-branch 23 --rebase
  $ gh pr update-branch 23 --repo owner/repo`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`rebase`, `Update PR branch by rebasing on top of latest base branch`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`view`, `View a pull request`),
				ox.Banner(`Display the title, body, and other information about a pull request.

Without an argument, the pull request that belongs to the current branch
is displayed.

With `+"`"+`--web`+"`"+` flag, open the pull request in a web browser instead.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number> | <url> | <branch>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`comments`, `View pull request comments`, ox.Short("c"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open a pull request in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`project`, `Work with GitHub Projects.`),
			ox.Banner(`Work with GitHub Projects.

The minimum required scope for the token is: `+"`"+`project`+"`"+`.
You can verify your token scope by running `+"`"+`gh auth status`+"`"+` and
add the `+"`"+`project`+"`"+` scope by running `+"`"+`gh auth refresh -s project`+"`"+`.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh project create --owner monalisa --title "Roadmap"
  $ gh project view 1 --owner cli --web
  $ gh project field-list 1 --owner cli
  $ gh project item-list 1 --owner cli`),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`close`, `Close a project`),
				ox.Banner(`Close a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Close project "1" owned by monalisa
  $ gh project close 1 --owner monalisa
  
  # Reopen closed project "1" owned by github
  $ gh project close 1 --owner github --undo`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`undo`, `Reopen a closed project`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`copy`, `Copy a project`),
				ox.Banner(`Copy a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Copy project "1" owned by monalisa to github
  $ gh project copy 1 --source-owner monalisa --target-owner github --title "a new project"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`drafts`, `Include draft issues when copying`, ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`source-owner`, `Login of the source owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`target-owner`, `Login of the target owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`title`, `Title for the new project`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`create`, `Create a project`),
				ox.Banner(`Create a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Create a new project owned by login monalisa
  $ gh project create --owner monalisa --title "a new project"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`title`, `Title for the project`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a project`),
				ox.Banner(`Delete a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Delete the current user's project "1"
  $ gh project delete 1 --owner "@me"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit a project`),
				ox.Banner(`Edit a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Edit the title of monalisa's project "1"
  $ gh project edit 1 --owner monalisa --title "New title"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`description`, `New description of the project`, ox.Short("d"), ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`readme`, `New readme for the project`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`title`, `New title for the project`, ox.Section(0)).
					String(`visibility`, `Change project visibility: {PUBLIC|PRIVATE}`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`field-create`, `Create a field in a project`),
				ox.Banner(`Create a field in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Create a field in the current user's project "1"
  $ gh project field-create 1 --owner "@me" --name "new field" --data-type "text"
  
  # Create a field with three options to select from for owner monalisa
  $ gh project field-create 1 --owner monalisa --name "new field" --data-type "SINGLE_SELECT" --single-select-options "one,two,three"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`data-type`, `DataType of the new field.: {TEXT|SINGLE_SELECT|DATE|NUMBER}`, ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`name`, `Name of the new field`, ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					Slice(`single-select-options`, `Options for SINGLE_SELECT data type`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`field-delete`, `Delete a field in a project`),
				ox.Banner(`Delete a field in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`id`, `ID of the field to delete`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`field-list`, `List the fields in a project`),
				ox.Banner(`List the fields in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # List fields in the current user's project "1"
  $ gh project field-list 1 --owner "@me"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Int(`limit`, `Maximum number of fields to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-add`, `Add a pull request or an issue to a project`),
				ox.Banner(`Add a pull request or an issue to a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Add an item to monalisa's project "1"
  $ gh project item-add 1 --owner monalisa --url https://github.com/monalisa/myproject/issues/23`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`url`, `URL of the issue or pull request to add to the project`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-archive`, `Archive an item in a project`),
				ox.Banner(`Archive an item in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Archive an item in the current user's project "1"
  $ gh project item-archive 1 --owner "@me" --id <item-ID>`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`id`, `ID of the item to archive`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`undo`, `Unarchive an item`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-create`, `Create a draft issue item in a project`),
				ox.Banner(`Create a draft issue item in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Create a draft issue in the current user's project "1"
  $ gh project item-create 1 --owner "@me" --title "new item" --body "new item body"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`body`, `Body for the draft issue`, ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`title`, `Title for the draft issue`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-delete`, `Delete an item from a project by ID`),
				ox.Banner(`Delete an item from a project by ID

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Delete an item in the current user's project "1"
  $ gh project item-delete 1 --owner "@me" --id <item-id>`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`id`, `ID of the item to delete`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-edit`, `Edit an item in a project`),
				ox.Banner(`Edit either a draft issue or a project item. Both usages require the ID of the item to edit.

For non-draft issues, the ID of the project is also required, and only a single field value can be updated per invocation.

Remove project item field value using `+"`"+`--clear`+"`"+` flag.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Edit an item's text field value
  $ gh project item-edit --id <item-id> --field-id <field-id> --project-id <project-id> --text "new text"
  
  # Clear an item's field value
  $ gh project item-edit --id <item-id> --field-id <field-id> --project-id <project-id> --clear`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`body`, `Body of the draft issue item`, ox.Section(0)).
					Bool(`clear`, `Remove field value`, ox.Section(0)).
					String(`date`, `Date value for the field (YYYY-MM-DD)`, ox.Section(0)).
					String(`field-id`, `ID of the field to update`, ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`id`, `ID of the item to edit`, ox.Section(0)).
					String(`iteration-id`, `ID of the iteration value to set on the field`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Float32(`number`, `Number value for the field`, ox.Spec(`float`), ox.Section(0)).
					String(`project-id`, `ID of the project to which the field belongs to`, ox.Section(0)).
					String(`single-select-option-id`, `ID of the single select option value to set on the field`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`text`, `Text value for the field`, ox.Section(0)).
					String(`title`, `Title of the draft issue item`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`item-list`, `List the items in a project`),
				ox.Banner(`List the items in a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # List the items in the current users's project "1"
  $ gh project item-list 1 --owner "@me"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Int(`limit`, `Maximum number of items to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`link`, `Link a project to a repository or a team`),
				ox.Banner(`Link a project to a repository or a team`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Link monalisa's project 1 to her repository "my_repo"
  $ gh project link 1 --owner monalisa --repo my_repo
  
  # Link monalisa's organization's project 1 to her team "my_team"
  $ gh project link 1 --owner my_organization --team my_team
  
  # Link monalisa's project 1 to the repository of current directory if neither --repo nor --team is specified
  $ gh project link 1`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`repo`, `The repository to be linked to this project`, ox.Short("R"), ox.Section(0)).
					String(`team`, `The team to be linked to this project`, ox.Short("T"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List the projects for an owner`),
				ox.Banner(`List the projects for an owner

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("project ls"),
				ox.Example(`
  # List the current user's projects
  $ gh project list
  
  # List the projects for org github including closed projects
  $ gh project list --owner github --closed`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`closed`, `Include closed projects`, ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Int(`limit`, `Maximum number of projects to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`owner`, `Login of the owner`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open projects list in the browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`mark-template`, `Mark a project as a template`),
				ox.Banner(`Mark a project as a template

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Mark the github org's project "1" as a template
  $ gh project mark-template 1 --owner "github"
  
  # Unmark the github org's project "1" as a template
  $ gh project mark-template 1 --owner "github" --undo`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the org owner.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`undo`, `Unmark the project as a template.`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`unlink`, `Unlink a project from a repository or a team`),
				ox.Banner(`Unlink a project from a repository or a team`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # Unlink monalisa's project 1 from her repository "my_repo"
  $ gh project unlink 1 --owner monalisa --repo my_repo
  
  # Unlink monalisa's organization's project 1 from her team "my_team"
  $ gh project unlink 1 --owner my_organization --team my_team
  
  # Unlink monalisa's project 1 from the repository of current directory if neither --repo nor --team is specified
  $ gh project unlink 1`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`repo`, `The repository to be unlinked from this project`, ox.Short("R"), ox.Section(0)).
					String(`team`, `The team to be unlinked from this project`, ox.Short("T"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`view`, `View a project`),
				ox.Banner(`View a project

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<number>] [flags]`),
				ox.Example(`
  # View the current user's project "1"
  $ gh project view 1
  
  # Open user monalisa's project "1" in the browser
  $ gh project view 1 --owner monalisa --web`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					String(`owner`, `Login of the owner. Use "@me" for the current user.`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open a project in the browser`, ox.Short("w"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`release`, `Manage releases`),
			ox.Banner(`Manage releases`),
			ox.Spec(`<command> [flags]`),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`create`, `Create a new release`),
				ox.Banner(`Create a new GitHub Release for a repository.

A list of asset files may be given to upload to the new release. To define a
display label for an asset, append text starting with `+"`"+`#`+"`"+` after the file name.

If a matching git tag does not yet exist, one will automatically get created
from the latest state of the default branch.
Use `+"`"+`--target`+"`"+` to point to a different branch or commit for the automatic tag creation.
Use `+"`"+`--verify-tag`+"`"+` to abort the release if the tag doesn't already exist.
To fetch the new tag locally after the release, do `+"`"+`git fetch --tags origin`+"`"+`.

To create a release from an annotated git tag, first create one locally with
git, push the tag to GitHub, then run this command.
Use `+"`"+`--notes-from-tag`+"`"+` to automatically generate the release notes
from the annotated git tag.

When using automatically generated release notes, a release title will also be automatically
generated unless a title was explicitly passed. Additional release notes can be prepended to
automatically generated notes by using the `+"`"+`--notes`+"`"+` flag.

By default, the release is created even if there are no new commits since the last release.
This may result in the same or duplicate release which may not be desirable in some cases.
Use `+"`"+`--fail-on-no-commits`+"`"+` to fail if no new commits are available. This flag has no
effect if there are no existing releases or this is the very first release.`),
				ox.Spec(`[<tag>] [<filename>... | <pattern>...]`),
				ox.Aliases("release new"),
				ox.Example(`
  # Interactively create a release
  $ gh release create
  
  # Interactively create a release from specific tag
  $ gh release create v1.2.3
  
  # Non-interactively create a release
  $ gh release create v1.2.3 --notes "bugfix release"
  
  # Use automatically generated release notes
  $ gh release create v1.2.3 --generate-notes
  
  # Use release notes from a file
  $ gh release create v1.2.3 -F release-notes.md
  
  # Use annotated tag notes
  $ gh release create v1.2.3 --notes-from-tag
  
  # Don't mark the release as latest
  $ gh release create v1.2.3 --latest=false
  
  # Upload all tarballs in a directory as release assets
  $ gh release create v1.2.3 ./dist/*.tgz
  
  # Upload a release asset with a display label
  $ gh release create v1.2.3 '/path/to/asset.zip#My display label'
  
  # Create a release and start a discussion
  $ gh release create v1.2.3 --discussion-category "General"
  
  # Create a release only if there are new commits available since the last release
  $ gh release create v1.2.3 --fail-on-no-commits`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`discussion-category`, `Start a discussion in the specified category`, ox.Section(0)).
					Bool(`draft`, `Save the release as a draft instead of publishing it`, ox.Short("d"), ox.Section(0)).
					Bool(`fail-on-no-commits`, `Fail if there are no commits since the last release (no impact on the first release)`, ox.Section(0)).
					Bool(`generate-notes`, `Automatically generate title and notes for the release`, ox.Section(0)).
					Bool(`latest`, `Mark this release as "Latest" (default [automatic based on date and version]). --latest=false to explicitly NOT set as latest`, ox.Section(0)).
					String(`notes`, `Release notes`, ox.Short("n"), ox.Section(0)).
					String(`notes-file`, `Read release notes from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`notes-from-tag`, `Automatically generate notes from annotated tag`, ox.Section(0)).
					String(`notes-start-tag`, `Tag to use as the starting point for generating release notes`, ox.Section(0)).
					Bool(`prerelease`, `Mark the release as a prerelease`, ox.Short("p"), ox.Section(0)).
					String(`target`, `Target branch or full commit SHA`, ox.Spec(`branch`), ox.Default("[main branch]"), ox.Section(0)).
					String(`title`, `Release title`, ox.Short("t"), ox.Section(0)).
					Bool(`verify-tag`, `Abort in case the git tag doesn't already exist in the remote repository`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List releases in a repository`),
				ox.Banner(`List releases in a repository

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("release ls"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`exclude-drafts`, `Exclude draft releases`, ox.Section(0)).
					Bool(`exclude-pre-releases`, `Exclude pre-releases`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of items to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`order`, `Order of releases returned: {asc|desc}`, ox.Default("desc"), ox.Short("O"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a release`),
				ox.Banner(`Delete a release`),
				ox.Spec(`<tag> [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`cleanup-tag`, `Delete the specified tag in addition to its release`, ox.Section(0)).
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`delete-asset`, `Delete an asset from a release`),
				ox.Banner(`Delete an asset from a release`),
				ox.Spec(`<tag> <asset-name> [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`download`, `Download release assets`),
				ox.Banner(`Download assets from a GitHub release.

Without an explicit tag name argument, assets are downloaded from the
latest release in the project. In this case, `+"`"+`--pattern`+"`"+` or `+"`"+`--archive`+"`"+`
is required.`),
				ox.Spec(`[<tag>] [flags]`),
				ox.Example(`
  # Download all assets from a specific release
  $ gh release download v1.2.3
  
  # Download only Debian packages for the latest release
  $ gh release download --pattern '*.deb'
  
  # Specify multiple file patterns
  $ gh release download -p '*.deb' -p '*.rpm'
  
  # Download the archive of the source code for a release
  $ gh release download v1.2.3 --archive=zip`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`archive`, `Download the source code archive in the specified format (zip or tar.gz)`, ox.Spec(`format`), ox.Short("A"), ox.Section(0)).
					Bool(`clobber`, `Overwrite existing files of the same name`, ox.Section(0)).
					String(`dir`, `The directory to download files into`, ox.Spec(`directory`), ox.Default("."), ox.Short("D"), ox.Section(0)).
					String(`output`, `The file to write a single asset to (use "-" to write to standard output)`, ox.Spec(`file`), ox.Short("O"), ox.Section(0)).
					Array(`pattern`, `Download only assets that match a glob pattern`, ox.Short("p"), ox.Section(0)).
					Bool(`skip-existing`, `Skip downloading when files of the same name exist`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit a release`),
				ox.Banner(`Edit a release`),
				ox.Spec(`<tag>`),
				ox.Example(`
  # Publish a release that was previously a draft
  $ gh release edit v1.0 --draft=false
  
  # Update the release notes from the content of a file
  $ gh release edit v1.0 --notes-file /path/to/release_notes.md`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`discussion-category`, `Start a discussion in the specified category when publishing a draft`, ox.Section(0)).
					Bool(`draft`, `Save the release as a draft instead of publishing it`, ox.Section(0)).
					Bool(`latest`, `Explicitly mark the release as "Latest"`, ox.Section(0)).
					String(`notes`, `Release notes`, ox.Short("n"), ox.Section(0)).
					String(`notes-file`, `Read release notes from file (use "-" to read from standard input)`, ox.Spec(`file`), ox.Short("F"), ox.Section(0)).
					Bool(`prerelease`, `Mark the release as a prerelease`, ox.Section(0)).
					String(`tag`, `The name of the tag`, ox.Section(0)).
					String(`target`, `Target branch or full commit SHA`, ox.Spec(`branch`), ox.Default("[main branch]"), ox.Section(0)).
					String(`title`, `Release title`, ox.Short("t"), ox.Section(0)).
					Bool(`verify-tag`, `Abort in case the git tag doesn't already exist in the remote repository`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`upload`, `Upload assets to a release`),
				ox.Banner(`Upload asset files to a GitHub Release.

To define a display label for an asset, append text starting with `+"`"+`#`+"`"+` after the
file name.`),
				ox.Spec(`<tag> <files>... [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`clobber`, `Overwrite existing assets of the same name`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`view`, `View information about a release`),
				ox.Banner(`View information about a GitHub Release.

Without an explicit tag name argument, the latest release in the project
is shown.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<tag>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open the release in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`repo`, `Manage repositories`),
			ox.Banner(`Work with GitHub repositories.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh repo create
  $ gh repo clone cli/cli
  $ gh repo view --web`),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`create`, `Create a new repository`),
				ox.Banner(`Create a new GitHub repository.

To create a repository interactively, use `+"`"+`gh repo create`+"`"+` with no arguments.

To create a remote repository non-interactively, supply the repository name and one of `+"`"+`--public`+"`"+`, `+"`"+`--private`+"`"+`, or `+"`"+`--internal`+"`"+`.
Pass `+"`"+`--clone`+"`"+` to clone the new repository locally.

If the `+"`"+`OWNER/`+"`"+` portion of the `+"`"+`OWNER/REPO`+"`"+` name argument is omitted, it
defaults to the name of the authenticating user.

To create a remote repository from an existing local repository, specify the source directory with `+"`"+`--source`+"`"+`.
By default, the remote repository name will be the name of the source directory.

Pass `+"`"+`--push`+"`"+` to push any local commits to the new repository. If the repo is bare, this will mirror all refs.

For language or platform .gitignore templates to use with `+"`"+`--gitignore`+"`"+`, <https://github.com/github/gitignore>.

For license keywords to use with `+"`"+`--license`+"`"+`, run `+"`"+`gh repo license list`+"`"+` or visit <https://choosealicense.com>.

The repo is created with the configured repository default branch, see <https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-user-account-settings/managing-the-default-branch-name-for-your-repositories>.`),
				ox.Spec(`[<name>] [flags]`),
				ox.Aliases("repo new"),
				ox.Example(`
  # Create a repository interactively
  $ gh repo create
  
  # Create a new remote repository and clone it locally
  $ gh repo create my-project --public --clone
  
  # Create a new remote repository in a different organization
  $ gh repo create my-org/my-project --public
  
  # Create a remote repository from the current directory
  $ gh repo create my-project --private --source=. --remote=upstream`),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`add-readme`, `Add a README file to the new repository`, ox.Section(0)).
					Bool(`clone`, `Clone the new repository to the current directory`, ox.Short("c"), ox.Section(0)).
					String(`description`, `Description of the repository`, ox.Short("d"), ox.Section(0)).
					Bool(`disable-issues`, `Disable issues in the new repository`, ox.Section(0)).
					Bool(`disable-wiki`, `Disable wiki in the new repository`, ox.Section(0)).
					String(`gitignore`, `Specify a gitignore template for the repository`, ox.Short("g"), ox.Section(0)).
					URL(`homepage`, `Repository home page URL`, ox.Short("h"), ox.Section(0)).
					Bool(`include-all-branches`, `Include all branches from template repository`, ox.Section(0)).
					Bool(`internal`, `Make the new repository internal`, ox.Section(0)).
					String(`license`, `Specify an Open Source License for the repository`, ox.Short("l"), ox.Section(0)).
					Bool(`private`, `Make the new repository private`, ox.Section(0)).
					Bool(`public`, `Make the new repository public`, ox.Section(0)).
					Bool(`push`, `Push local commits to the new repository`, ox.Section(0)).
					String(`remote`, `Specify remote name for the new repository`, ox.Short("r"), ox.Section(0)).
					String(`source`, `Specify path to local repository to use as source`, ox.Short("s"), ox.Section(0)).
					String(`team`, `The name of the organization team to be granted access`, ox.Spec(`name`), ox.Short("t"), ox.Section(0)).
					String(`template`, `Make the new repository based on a template repository`, ox.Spec(`repository`), ox.Short("p"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List repositories owned by user or organization`),
				ox.Banner(`List repositories owned by a user or organization.

Note that the list will only include repositories owned by the provided argument,
and the `+"`"+`--fork`+"`"+` or `+"`"+`--source`+"`"+` flags will not traverse ownership boundaries. For example,
when listing the forks in an organization, the output would not include those owned by individual users.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<owner>] [flags]`),
				ox.Aliases("repo ls"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`archived`, `Show only archived repositories`, ox.Section(0)).
					Bool(`fork`, `Show only forks`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`language`, `Filter by primary coding language`, ox.Short("l"), ox.Section(0)).
					Int(`limit`, `Maximum number of repositories to list`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Bool(`no-archived`, `Omit archived repositories`, ox.Section(0)).
					Bool(`source`, `Show only non-forks`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Slice(`topic`, `Filter by topic`, ox.Section(0)).
					String(`visibility`, `Filter by repository visibility: {public|private|internal}`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`archive`, `Archive a repository`),
				ox.Banner(`Archive a GitHub repository.

With no argument, archives the current repository.`),
				ox.Spec(`[<repository>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`autolink`, `Manage autolink references`),
				ox.Banner(`Autolinks link issues, pull requests, commit messages, and release descriptions to external third-party services.

Autolinks require `+"`"+`admin`+"`"+` role to view or manage.

For more information, see <https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/managing-repository-settings/configuring-autolinks-to-reference-external-resources>`),
				ox.Spec(`<command> [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
				ox.Sub(
					ox.Usage(`create`, `Create a new autolink reference`),
					ox.Banner(`Create a new autolink reference for a repository.

The `+"`"+`keyPrefix`+"`"+` argument specifies the prefix that will generate a link when it is appended by certain characters.

The `+"`"+`urlTemplate`+"`"+` argument specifies the target URL that will be generated when the keyPrefix is found, which
must contain `+"`"+`<num>`+"`"+` variable for the reference number.

By default, autolinks are alphanumeric with `+"`"+`--numeric`+"`"+` flag used to create a numeric autolink.

The `+"`"+`<num>`+"`"+` variable behavior differs depending on whether the autolink is alphanumeric or numeric:

- alphanumeric: matches `+"`"+`A-Z`+"`"+` (case insensitive), `+"`"+`0-9`+"`"+`, and `+"`"+`-`+"`"+`
- numeric: matches `+"`"+`0-9`+"`"+`

If the template contains multiple instances of `+"`"+`<num>`+"`"+`, only the first will be replaced.`),
					ox.Spec(`<keyPrefix> <urlTemplate> [flags]`),
					ox.Aliases("repo autolink new"),
					ox.Example(`
  # Create an alphanumeric autolink to example.com for the key prefix "TICKET-".
  # Generates https://example.com/TICKET?query=123abc from "TICKET-123abc".
  $ gh repo autolink create TICKET- "https://example.com/TICKET?query=<num>"
  
  # Create a numeric autolink to example.com for the key prefix "STORY-".
  # Generates https://example.com/STORY?id=123 from "STORY-123".
  $ gh repo autolink create STORY- "https://example.com/STORY?id=<num>" --numeric`),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						Bool(`numeric`, `Mark autolink as numeric`, ox.Short("n"), ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
				ox.Sub(
					ox.Usage(`delete`, `Delete an autolink reference`),
					ox.Banner(`Delete an autolink reference for a repository.`),
					ox.Spec(`<id> [flags]`),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						Bool(`yes`, `Confirm deletion without prompting`, ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
				ox.Sub(
					ox.Usage(`list`, `List autolink references for a GitHub repository`),
					ox.Banner(`Gets all autolink references that are configured for a repository.

Information about autolinks is only available to repository administrators.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
					ox.Spec(`[flags]`),
					ox.Aliases("repo autolink ls"),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
						Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
						String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
						Bool(`web`, `List autolink references in the web browser`, ox.Short("w"), ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
				ox.Sub(
					ox.Usage(`view`, `View an autolink reference`),
					ox.Banner(`View an autolink reference for a repository.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
					ox.Spec(`<id> [flags]`),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
						Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
						String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
			),
			ox.Sub(
				ox.Usage(`clone`, `Clone a repository locally`),
				ox.Banner(`Clone a GitHub repository locally. Pass additional `+"`"+`git clone`+"`"+` flags by listing
them after `+"`"+`--`+"`"+`.

If the `+"`"+`OWNER/`+"`"+` portion of the `+"`"+`OWNER/REPO`+"`"+` repository argument is omitted, it
defaults to the name of the authenticating user.

When a protocol scheme is not provided in the repository argument, the `+"`"+`git_protocol`+"`"+` will be
chosen from your configuration, which can be checked via `+"`"+`gh config get git_protocol`+"`"+`. If the protocol
scheme is provided, the repository will be cloned using the specified protocol.

If the repository is a fork, its parent repository will be added as an additional
git remote called `+"`"+`upstream`+"`"+`. The remote name can be configured using `+"`"+`--upstream-remote-name`+"`"+`.
The `+"`"+`--upstream-remote-name`+"`"+` option supports an `+"`"+`@owner`+"`"+` value which will name
the remote after the owner of the parent repository.

If the repository is a fork, its parent repository will be set as the default remote repository.`),
				ox.Spec(`<repository> [<directory>] [-- <gitflags>...]`),
				ox.Example(`
  # Clone a repository from a specific org
  $ gh repo clone cli/cli
  
  # Clone a repository from your own account
  $ gh repo clone myrepo
  
  # Clone a repo, overriding git protocol configuration
  $ gh repo clone https://github.com/cli/cli
  $ gh repo clone git@github.com:cli/cli.git
  
  # Clone a repository to a custom directory
  $ gh repo clone cli/cli workspace/cli
  
  # Clone a repository with additional git clone flags
  $ gh repo clone cli/cli -- --depth=1`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`upstream-remote-name`, `Upstream remote name when cloning a fork`, ox.Default("upstream"), ox.Short("u"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a repository`),
				ox.Banner(`Delete a GitHub repository.

With no argument, deletes the current repository. Otherwise, deletes the specified repository.

Deletion requires authorization with the `+"`"+`delete_repo`+"`"+` scope.
To authorize, run `+"`"+`gh auth refresh -s delete_repo`+"`"+``),
				ox.Spec(`[<repository>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Confirm deletion without prompting`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`deploy-key`, `Manage deploy keys in a repository`),
				ox.Banner(`Manage deploy keys in a repository`),
				ox.Spec(`<command> [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
				ox.Sub(
					ox.Usage(`add`, `Add a deploy key to a GitHub repository`),
					ox.Banner(`Add a deploy key to a GitHub repository.

Note that any key added by gh will be associated with the current authentication token.
If you de-authorize the GitHub CLI app or authentication token from your account, any
deploy keys added by GitHub CLI will be removed as well.`),
					ox.Spec(`<key-file> [flags]`),
					ox.Example(`
  # Generate a passwordless SSH key and add it as a deploy key to a repository
  $ ssh-keygen -t ed25519 -C "my description" -N "" -f ~/.ssh/gh-test
  $ gh repo deploy-key add ~/.ssh/gh-test.pub`),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						Bool(`allow-write`, `Allow write access for the key`, ox.Short("w"), ox.Section(0)).
						String(`title`, `Title of the new key`, ox.Short("t"), ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
				ox.Sub(
					ox.Usage(`delete`, `Delete a deploy key from a GitHub repository`),
					ox.Banner(`Delete a deploy key from a GitHub repository`),
					ox.Spec(`<key-id> [flags]`),
					ox.Help(ox.Sections(
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
				),
				ox.Sub(
					ox.Usage(`list`, `List deploy keys in a GitHub repository`),
					ox.Banner(`List deploy keys in a GitHub repository

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
					ox.Spec(`[flags]`),
					ox.Aliases("repo deploy-key ls"),
					ox.Help(ox.Sections(
						"FLAGS",
						"INHERITED FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
						Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
						String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
						String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
				),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit repository settings`),
				ox.Banner(`Edit repository settings.

To toggle a setting off, use the `+"`"+`--<flag>=false`+"`"+` syntax.

Changing repository visibility can have unexpected consequences including but not limited to:

- Losing stars and watchers, affecting repository ranking
- Detaching public forks from the network
- Disabling push rulesets
- Allowing access to GitHub Actions history and logs

When the `+"`"+`--visibility`+"`"+` flag is used, `+"`"+`--accept-visibility-change-consequences`+"`"+` flag is required.

For information on all the potential consequences, see <https://gh.io/setting-repository-visibility>.`),
				ox.Spec(`[<repository>] [flags]`),
				ox.Example(`
  # Enable issues and wiki
  $ gh repo edit --enable-issues --enable-wiki
  
  # Disable projects
  $ gh repo edit --enable-projects=false`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`accept-visibility-change-consequences`, `Accept the consequences of changing the repository visibility`, ox.Section(0)).
					Slice(`add-topic`, `Add repository topic`, ox.Section(0)).
					Bool(`allow-forking`, `Allow forking of an organization repository`, ox.Section(0)).
					Bool(`allow-update-branch`, `Allow a pull request head branch that is behind its base branch to be updated`, ox.Section(0)).
					String(`default-branch`, `Set the default branch name for the repository`, ox.Spec(`name`), ox.Section(0)).
					Bool(`delete-branch-on-merge`, `Delete head branch when pull requests are merged`, ox.Section(0)).
					String(`description`, `Description of the repository`, ox.Short("d"), ox.Section(0)).
					Bool(`enable-advanced-security`, `Enable advanced security in the repository`, ox.Section(0)).
					Bool(`enable-auto-merge`, `Enable auto-merge functionality`, ox.Section(0)).
					Bool(`enable-discussions`, `Enable discussions in the repository`, ox.Section(0)).
					Bool(`enable-issues`, `Enable issues in the repository`, ox.Section(0)).
					Bool(`enable-merge-commit`, `Enable merging pull requests via merge commit`, ox.Section(0)).
					Bool(`enable-projects`, `Enable projects in the repository`, ox.Section(0)).
					Bool(`enable-rebase-merge`, `Enable merging pull requests via rebase`, ox.Section(0)).
					Bool(`enable-secret-scanning`, `Enable secret scanning in the repository`, ox.Section(0)).
					Bool(`enable-secret-scanning-push-protection`, `Enable secret scanning push protection in the repository. Secret scanning must be enabled first`, ox.Section(0)).
					Bool(`enable-squash-merge`, `Enable merging pull requests via squashed commit`, ox.Section(0)).
					Bool(`enable-wiki`, `Enable wiki in the repository`, ox.Section(0)).
					URL(`homepage`, `Repository home page URL`, ox.Short("h"), ox.Section(0)).
					Slice(`remove-topic`, `Remove repository topic`, ox.Section(0)).
					Bool(`template`, `Make the repository available as a template repository`, ox.Section(0)).
					String(`visibility`, `Change the visibility of the repository to {public,private,internal}`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`fork`, `Create a fork of a repository`),
				ox.Banner(`Create a fork of a repository.

With no argument, creates a fork of the current repository. Otherwise, forks
the specified repository.

By default, the new fork is set to be your `+"`"+`origin`+"`"+` remote and any existing
origin remote is renamed to `+"`"+`upstream`+"`"+`. To alter this behavior, you can set
a name for the new fork's remote with `+"`"+`--remote-name`+"`"+`.

The `+"`"+`upstream`+"`"+` remote will be set as the default remote repository.

Additional `+"`"+`git clone`+"`"+` flags can be passed after `+"`"+`--`+"`"+`.`),
				ox.Spec(`[<repository>] [-- <gitflags>...] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`clone`, `Clone the fork`, ox.Section(0)).
					Bool(`default-branch-only`, `Only include the default branch in the fork`, ox.Section(0)).
					String(`fork-name`, `Rename the forked repository`, ox.Section(0)).
					String(`org`, `Create the fork in an organization`, ox.Section(0)).
					Bool(`remote`, `Add a git remote for the fork`, ox.Section(0)).
					String(`remote-name`, `Specify the name for the new remote`, ox.Default("origin"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`gitignore`, `List and view available repository gitignore templates`),
				ox.Banner(`List and view available repository gitignore templates`),
				ox.Spec(`<command> [flags]`),
				ox.Section(1),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Sub(
					ox.Usage(`list`, `List available repository gitignore templates`),
					ox.Banner(`List available repository gitignore templates`),
					ox.Spec(`[flags]`),
					ox.Aliases("repo gitignore ls"),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				),
				ox.Sub(
					ox.Usage(`view`, `View an available repository gitignore template`),
					ox.Banner(`View an available repository `+"`"+`.gitignore`+"`"+` template.

`+"`"+`<template>`+"`"+` is a case-sensitive `+"`"+`.gitignore`+"`"+` template name.

For a list of available templates, run `+"`"+`gh repo gitignore list`+"`"+`.`),
					ox.Spec(`<template> [flags]`),
					ox.Example(`
  # View the Go gitignore template
  $ gh repo gitignore view Go
  
  # View the Python gitignore template
  $ gh repo gitignore view Python
  
  # Create a new .gitignore file using the Go template
  $ gh repo gitignore view Go > .gitignore
  
  # Create a new .gitignore file using the Python template
  $ gh repo gitignore view Python > .gitignore`),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				),
			),
			ox.Sub(
				ox.Usage(`license`, `Explore repository licenses`),
				ox.Banner(`Explore repository licenses`),
				ox.Spec(`<command> [flags]`),
				ox.Section(1),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Sub(
					ox.Usage(`list`, `List common repository licenses`),
					ox.Banner(`List common repository licenses.

For even more licenses, visit <https://choosealicense.com/appendix>`),
					ox.Spec(`[flags]`),
					ox.Aliases("repo license ls"),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				),
				ox.Sub(
					ox.Usage(`view`, `View a specific repository license`),
					ox.Banner(`View a specific repository license by license key or SPDX ID.

Run `+"`"+`gh repo license list`+"`"+` to see available commonly used licenses. For even more licenses, visit <https://choosealicense.com/appendix>.`),
					ox.Spec(`{<license-key> | <spdx-id>} [flags]`),
					ox.Example(`
  # View the MIT license from SPDX ID
  $ gh repo license view MIT
  
  # View the MIT license from license key
  $ gh repo license view mit
  
  # View the GNU AGPL-3.0 license from SPDX ID
  $ gh repo license view AGPL-3.0
  
  # View the GNU AGPL-3.0 license from license key
  $ gh repo license view agpl-3.0
  
  # Create a LICENSE.md with the MIT license
  $ gh repo license view MIT > LICENSE.md`),
					ox.Help(ox.Sections(
						"FLAGS",
					)),
					ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
					ox.Flags().
						Bool(`web`, `Open https://choosealicense.com/ in the browser`, ox.Short("w"), ox.Section(0)),
				),
			),
			ox.Sub(
				ox.Usage(`rename`, `Rename a repository`),
				ox.Banner(`Rename a GitHub repository.

`+"`"+`<new-name>`+"`"+` is the desired repository name without the owner.

By default, the current repository is renamed. Otherwise, the repository specified
with `+"`"+`--repo`+"`"+` is renamed.

To transfer repository ownership to another user account or organization,
you must follow additional steps on <github.com>.

For more information on transferring repository ownership, see:
<https://docs.github.com/en/repositories/creating-and-managing-repositories/transferring-a-repository>`),
				ox.Spec(`[<new-name>] [flags]`),
				ox.Example(`
  # Rename the current repository (foo/bar -> foo/baz)
  $ gh repo rename baz
  
  # Rename the specified repository (qux/quux -> qux/baz)
  $ gh repo rename -R qux/quux baz`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)).
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`set-default`, `Configure default repository for this directory`),
				ox.Banner(`This command sets the default remote repository to use when querying the
GitHub API for the locally cloned repository.

gh uses the default repository for things like:

 - viewing and creating pull requests
 - viewing and creating issues
 - viewing and creating releases
 - working with GitHub Actions

### NOTE: gh does not use the default repository for managing repository and environment secrets.`),
				ox.Spec(`[<repository>] [flags]`),
				ox.Example(`
  # Interactively select a default repository
  $ gh repo set-default
  
  # Set a repository explicitly
  $ gh repo set-default owner/repo
  
  # View the current default repository
  $ gh repo set-default --view
  
  # Show more repository options in the interactive picker
  $ git remote add newrepo https://github.com/owner/repo
  $ gh repo set-default`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`unset`, `Unset the current default repository`, ox.Short("u"), ox.Section(0)).
					Bool(`view`, `View the current default repository`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`sync`, `Sync a repository`),
				ox.Banner(`Sync destination repository from source repository. Syncing uses the default branch
of the source repository to update the matching branch on the destination
repository so they are equal. A fast forward update will be used except when the
`+"`"+`--force`+"`"+` flag is specified, then the two branches will
be synced using a hard reset.

Without an argument, the local repository is selected as the destination repository.

The source repository is the parent of the destination repository by default.
This can be overridden with the `+"`"+`--source`+"`"+` flag.`),
				ox.Spec(`[<destination-repository>] [flags]`),
				ox.Example(`
  # Sync local repository from remote parent
  $ gh repo sync
  
  # Sync local repository from remote parent on specific branch
  $ gh repo sync --branch v1
  
  # Sync remote fork from its parent
  $ gh repo sync owner/cli-fork
  
  # Sync remote repository from another remote repository
  $ gh repo sync owner/repo --source owner2/repo2`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`branch`, `Branch to sync`, ox.Default("[default branch]"), ox.Short("b"), ox.Section(0)).
					Bool(`force`, `Hard reset the branch of the destination repository to match the source repository`, ox.Section(0)).
					String(`source`, `Source repository`, ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`unarchive`, `Unarchive a repository`),
				ox.Banner(`Unarchive a GitHub repository.

With no argument, unarchives the current repository.`),
				ox.Spec(`[<repository>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`view`, `View a repository`),
				ox.Banner(`Display the description and the README of a GitHub repository.

With no argument, the repository for the current directory is displayed.

With `+"`"+`--web`+"`"+`, open the repository in a web browser instead.

With `+"`"+`--branch`+"`"+`, view a specific branch of the repository.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<repository>] [flags]`),
				ox.Section(1),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`branch`, `View a specific branch of the repository`, ox.Short("b"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open a repository in the browser`, ox.Short("w"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`cache`, `Manage GitHub Actions caches`),
			ox.Banner(`Work with GitHub Actions caches.`),
			ox.Spec(`<command> [flags]`),
			ox.Example(`
  $ gh cache list
  $ gh cache delete --all`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`delete`, `Delete GitHub Actions caches`),
				ox.Banner(`Delete GitHub Actions caches.

Deletion requires authorization with the `+"`"+`repo`+"`"+` scope.`),
				ox.Spec(`[<cache-id> | <cache-key> | --all] [flags]`),
				ox.Example(`
  # Delete a cache by id
  $ gh cache delete 1234
  
  # Delete a cache by key
  $ gh cache delete cache-key
  
  # Delete a cache by id in a specific repo
  $ gh cache delete 1234 --repo cli/cli
  
  # Delete all caches (exit code 1 on no caches)
  $ gh cache delete --all
  
  # Delete all caches (exit code 0 on no caches)
  $ gh cache delete --all --succeed-on-no-caches`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Delete all caches`, ox.Short("a"), ox.Section(0)).
					String(`succeed-on-no-caches`, `Return exit code 0 if no caches found. Must be used in conjunction with --all`, ox.Spec(`--all`), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List GitHub Actions caches`),
				ox.Banner(`List GitHub Actions caches

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("cache ls"),
				ox.Example(`
  # List caches for current repository
  $ gh cache list
  
  # List caches for specific repository
  $ gh cache list --repo cli/cli
  
  # List caches sorted by least recently accessed
  $ gh cache list --sort last_accessed_at --order asc
  
  # List caches that have keys matching a prefix (or that match exactly)
  $ gh cache list --key key-prefix
  
  # List caches for a specific branch, replace <branch-name> with the actual branch name
  $ gh cache list --ref refs/heads/<branch-name>
  
  # List caches for a specific pull request, replace <pr-number> with the actual pull request number
  $ gh cache list --ref refs/pull/<pr-number>/merge`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`key`, `Filter by cache key prefix`, ox.Short("k"), ox.Section(0)).
					Int(`limit`, `Maximum number of caches to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`order`, `Order of caches returned: {asc|desc}`, ox.Default("desc"), ox.Short("O"), ox.Section(0)).
					String(`ref`, `Filter by ref, formatted as refs/heads/<branch name> or refs/pull/<number>/merge`, ox.Short("r"), ox.Section(0)).
					String(`sort`, `Sort fetched caches: {created_at|last_accessed_at|size_in_bytes}`, ox.Default("last_accessed_at"), ox.Short("S"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`run`, `View details about workflow runs`),
			ox.Banner(`List, view, and watch recent workflow runs from GitHub Actions.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`cancel`, `Cancel a workflow run`),
				ox.Banner(`Cancel a workflow run`),
				ox.Spec(`[<run-id>] [flags]`),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a workflow run`),
				ox.Banner(`Delete a workflow run`),
				ox.Spec(`[<run-id>] [flags]`),
				ox.Example(`
  # Interactively select a run to delete
  $ gh run delete
  
  # Delete a specific run
  $ gh run delete 12345`),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`download`, `Download artifacts generated by a workflow run`),
				ox.Banner(`Download artifacts generated by a GitHub Actions workflow run.

The contents of each artifact will be extracted under separate directories based on
the artifact name. If only a single artifact is specified, it will be extracted into
the current directory.

By default, this command downloads the latest artifact created and uploaded through
GitHub Actions. Because workflows can delete or overwrite artifacts, `+"`"+`<run-id>`+"`"+`
must be used to select an artifact from a specific workflow run.`),
				ox.Spec(`[<run-id>] [flags]`),
				ox.Example(`
  # Download all artifacts generated by a workflow run
  $ gh run download <run-id>
  
  # Download a specific artifact within a run
  $ gh run download <run-id> -n <name>
  
  # Download specific artifacts across all runs in a repository
  $ gh run download -n <name1> -n <name2>
  
  # Select artifacts to download interactively
  $ gh run download`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`dir`, `The directory to download artifacts into`, ox.Default("."), ox.Short("D"), ox.Section(0)).
					Array(`name`, `Download artifacts that match any of the given names`, ox.Short("n"), ox.Section(0)).
					Array(`pattern`, `Download artifacts that match a glob pattern`, ox.Short("p"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List recent workflow runs`),
				ox.Banner(`List recent workflow runs.

Note that providing the `+"`"+`workflow_name`+"`"+` to the `+"`"+`-w`+"`"+` flag will not fetch disabled workflows.
Also pass the `+"`"+`-a`+"`"+` flag to fetch disabled workflow runs using the `+"`"+`workflow_name`+"`"+` and the `+"`"+`-w`+"`"+` flag.

Runs created by organization and enterprise ruleset workflows will not display a workflow name due to GitHub API limitations.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("run ls"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Include disabled workflows`, ox.Short("a"), ox.Section(0)).
					String(`branch`, `Filter runs by branch`, ox.Short("b"), ox.Section(0)).
					String(`commit`, `Filter runs by the SHA of the commit`, ox.Spec(`SHA`), ox.Short("c"), ox.Section(0)).
					Date(`created`, `Filter runs by the date it was created`, ox.Section(0)).
					String(`event`, `Filter runs by which event triggered the run`, ox.Spec(`event`), ox.Short("e"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of runs to fetch`, ox.Default("20"), ox.Short("L"), ox.Section(0)).
					String(`status`, `Filter runs by status: {queued|completed|in_progress|requested|waiting|pending|action_required|cancelled|failure|neutral|skipped|stale|startup_failure|success|timed_out}`, ox.Short("s"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`user`, `Filter runs by user who triggered the run`, ox.Short("u"), ox.Section(0)).
					String(`workflow`, `Filter runs by workflow`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`rerun`, `Rerun a run`),
				ox.Banner(`Rerun an entire run, only failed jobs, or a specific job from a run.

Note that due to historical reasons, the `+"`"+`--job`+"`"+` flag may not take what you expect.
Specifically, when navigating to a job in the browser, the URL looks like this:
`+"`"+`https://github.com/<owner>/<repo>/actions/runs/<run-id>/jobs/<number>`+"`"+`.

However, this `+"`"+`<number>`+"`"+` should not be used with the `+"`"+`--job`+"`"+` flag and will result in the
API returning `+"`"+`404 NOT FOUND`+"`"+`. Instead, you can get the correct job IDs using the following command:

	gh run view <run-id> --json jobs --jq '.jobs[] | {name, databaseId}'

You will need to use databaseId field for triggering job re-runs.`),
				ox.Spec(`[<run-id>] [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`debug`, `Rerun with debug logging`, ox.Short("d"), ox.Section(0)).
					Bool(`failed`, `Rerun only failed jobs, including dependencies`, ox.Section(0)).
					String(`job`, `Rerun a specific job ID from a run, including dependencies`, ox.Short("j"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`view`, `View a summary of a workflow run`),
				ox.Banner(`View a summary of a workflow run.

This command does not support authenticating via fine grained PATs
as it is not currently possible to create a PAT with the `+"`"+`checks:read`+"`"+` permission.

Due to platform limitations, `+"`"+`gh`+"`"+` may not always be able to associate log lines with a
particular step in a job. In this case, the step name in the log output will be replaced with
`+"`"+`UNKNOWN STEP`+"`"+`.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<run-id>] [flags]`),
				ox.Example(`
  # Interactively select a run to view, optionally selecting a single job
  $ gh run view
  
  # View a specific run
  $ gh run view 12345
  
  # View a specific run with specific attempt number
  $ gh run view 12345 --attempt 3
  
  # View a specific job within a run
  $ gh run view --job 456789
  
  # View the full log for a specific job
  $ gh run view --log --job 456789
  
  # Exit non-zero if a run failed
  $ gh run view 0451 --exit-status && echo "run pending or passed"`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Uint(`attempt`, `The attempt number of the workflow run`, ox.Short("a"), ox.Section(0)).
					Bool(`exit-status`, `Exit with non-zero status if run failed`, ox.Section(0)).
					String(`job`, `View a specific job ID from a run`, ox.Short("j"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Bool(`log`, `View full log for either a run or specific job`, ox.Section(0)).
					Bool(`log-failed`, `View the log for any failed steps in a run or specific job`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`verbose`, `Show job steps`, ox.Short("v"), ox.Section(0)).
					Bool(`web`, `Open run in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`watch`, `Watch a run until it completes, showing its progress`),
				ox.Banner(`Watch a run until it completes, showing its progress.

This command does not support authenticating via fine grained PATs
as it is not currently possible to create a PAT with the `+"`"+`checks:read`+"`"+` permission.`),
				ox.Spec(`<run-id> [flags]`),
				ox.Example(`
  # Watch a run until it's done
  $ gh run watch
  
  # Run some other command when the run is finished
  $ gh run watch && notify-send 'run is done!'`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`exit-status`, `Exit with non-zero status if run fails`, ox.Section(0)).
					Int(`interval`, `Refresh interval in seconds`, ox.Default("3"), ox.Short("i"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`workflow`, `View details about GitHub Actions workflows`),
			ox.Banner(`List, view, and run workflows in GitHub Actions.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`disable`, `Disable a workflow`),
				ox.Banner(`Disable a workflow, preventing it from running or showing up when listing workflows.`),
				ox.Spec(`[<workflow-id> | <workflow-name>] [flags]`),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`enable`, `Enable a workflow`),
				ox.Banner(`Enable a workflow, allowing it to be run and show up when listing workflows.`),
				ox.Spec(`[<workflow-id> | <workflow-name>] [flags]`),
				ox.Help(ox.Sections(
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List workflows`),
				ox.Banner(`List workflow files, hiding disabled workflows by default.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("workflow ls"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Include disabled workflows`, ox.Short("a"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of workflows to fetch`, ox.Default("50"), ox.Short("L"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`run`, `Run a workflow by creating a workflow_dispatch event`),
				ox.Banner(`Create a `+"`"+`workflow_dispatch`+"`"+` event for a given workflow.

This command will trigger GitHub Actions to run a given workflow file. The given workflow file must
support an `+"`"+`on.workflow_dispatch`+"`"+` trigger in order to be run in this way.

If the workflow file supports inputs, they can be specified in a few ways:

- Interactively
- Via `+"`"+`-f/--raw-field`+"`"+` or `+"`"+`-F/--field`+"`"+` flags
- As JSON, via standard input`),
				ox.Spec(`[<workflow-id> | <workflow-name>] [flags]`),
				ox.Example(`
  # Have gh prompt you for what workflow you'd like to run and interactively collect inputs
  $ gh workflow run
  
  # Run the workflow file 'triage.yml' at the remote's default branch
  $ gh workflow run triage.yml
  
  # Run the workflow file 'triage.yml' at a specified ref
  $ gh workflow run triage.yml --ref my-branch
  
  # Run the workflow file 'triage.yml' with command line inputs
  $ gh workflow run triage.yml -f name=scully -f greeting=hello
  
  # Run the workflow file 'triage.yml' with JSON via standard input
  $ echo '{"name":"scully", "greeting":"hello"}' | gh workflow run triage.yml --json`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Map(`field`, `Add a string parameter in key=value format, respecting @ syntax (see "gh help api").`, ox.Spec(`key=value`), ox.Short("F"), ox.Section(0)).
					Bool(`json`, `Read workflow inputs as JSON via STDIN`, ox.Section(0)).
					Map(`raw-field`, `Add a string parameter in key=value format`, ox.Spec(`key=value`), ox.Short("f"), ox.Section(0)).
					String(`ref`, `Branch or tag name which contains the version of the workflow file you'd like to run`, ox.Short("r"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`view`, `View the summary of a workflow`),
				ox.Banner(`View the summary of a workflow`),
				ox.Spec(`[<workflow-id> | <workflow-name> | <filename>] [flags]`),
				ox.Example(`
  # Interactively select a workflow to view
  $ gh workflow view
  
  # View a specific workflow
  $ gh workflow view 0451`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`ref`, `The branch or tag name which contains the version of the workflow file you'd like to view`, ox.Short("r"), ox.Section(0)).
					Bool(`web`, `Open workflow in the browser`, ox.Short("w"), ox.Section(0)).
					Bool(`yaml`, `View the workflow yaml file`, ox.Short("y"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`co`, `Alias for "pr checkout"`),
			ox.Banner(`Check out a pull request in git`),
			ox.Spec(`gh pr checkout [<number> | <url> | <branch>] [flags]`),
			ox.Example(`
  # Interactively select a PR from the 10 most recent to check out
  $ gh pr checkout
  
  # Checkout a specific PR
  $ gh pr checkout 32
  $ gh pr checkout https://github.com/OWNER/REPO/pull/32
  $ gh pr checkout feature`),
			ox.Section(2),
			ox.Help(ox.Sections(
				"FLAGS",
				"INHERITED FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`branch`, `Local branch name to use`, ox.Default("[the name of the head branch]"), ox.Short("b"), ox.Section(0)).
				Bool(`detach`, `Checkout PR with a detached HEAD`, ox.Section(0)).
				Bool(`force`, `Reset the existing local branch to the latest state of the pull request`, ox.Short("f"), ox.Section(0)).
				Bool(`recurse-submodules`, `Update all submodules after checkout`, ox.Section(0)).
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
		),
		ox.Sub(
			ox.Usage(`alias`, `Create command shortcuts`),
			ox.Banner(`Aliases can be used to make shortcuts for gh commands or to compose multiple commands.

Run `+"`"+`gh help alias set`+"`"+` to learn more.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`delete`, `Delete set aliases`),
				ox.Banner(`Delete set aliases`),
				ox.Spec(`{<alias> | --all} [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Delete all aliases`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`import`, `Import aliases from a YAML file`),
				ox.Banner(`Import aliases from the contents of a YAML file.

Aliases should be defined as a map in YAML, where the keys represent aliases and
the values represent the corresponding expansions. An example file should look like
the following:

    bugs: issue list --label=bug
    igrep: '!gh issue list --label="$1" | grep "$2"'
    features: |-
        issue list
        --label=enhancement

Use `+"`"+`-`+"`"+` to read aliases (in YAML format) from standard input.

The output from `+"`"+`gh alias list`+"`"+` can be used to produce a YAML file
containing your aliases, which you can use to import them from one machine to
another. Run `+"`"+`gh help alias list`+"`"+` to learn more.`),
				ox.Spec(`[<filename> | -] [flags]`),
				ox.Example(`
  # Import aliases from a file
  $ gh alias import aliases.yml
  
  # Import aliases from standard input
  $ gh alias import -`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`clobber`, `Overwrite existing aliases of the same name`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List your aliases`),
				ox.Banner(`This command prints out all of the aliases gh is configured to use.`),
				ox.Spec(`[flags]`),
				ox.Aliases("alias ls"),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`set`, `Create a shortcut for a gh command`),
				ox.Banner(`Define a word that will expand to a full gh command when invoked.

The expansion may specify additional arguments and flags. If the expansion includes
positional placeholders such as `+"`"+`$1`+"`"+`, extra arguments that follow the alias will be
inserted appropriately. Otherwise, extra arguments will be appended to the expanded
command.

Use `+"`"+`-`+"`"+` as expansion argument to read the expansion string from standard input. This
is useful to avoid quoting issues when defining expansions.

If the expansion starts with `+"`"+`!`+"`"+` or if `+"`"+`--shell`+"`"+` was given, the expansion is a shell
expression that will be evaluated through the `+"`"+`sh`+"`"+` interpreter when the alias is
invoked. This allows for chaining multiple commands via piping and redirection.`),
				ox.Spec(`<alias> <expansion> [flags]`),
				ox.Example(`
  # Note: Command Prompt on Windows requires using double quotes for arguments
  $ gh alias set pv 'pr view'
  $ gh pv -w 123  #=> gh pr view -w 123
  
  $ gh alias set bugs 'issue list --label=bugs'
  $ gh bugs
  
  $ gh alias set homework 'issue list --assignee @me'
  $ gh homework
  
  $ gh alias set 'issue mine' 'issue list --mention @me'
  $ gh issue mine
  
  $ gh alias set epicsBy 'issue list --author="$1" --label="epic"'
  $ gh epicsBy vilmibm  #=> gh issue list --author="vilmibm" --label="epic"
  
  $ gh alias set --shell igrep 'gh issue list --label="$1" | grep "$2"'
  $ gh igrep epic foo  #=> gh issue list --label="epic" | grep "foo"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`clobber`, `Overwrite existing aliases of the same name`, ox.Section(0)).
					Bool(`shell`, `Declare an alias to be passed through a shell interpreter`, ox.Short("s"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`api`, `Make an authenticated GitHub API request`),
			ox.Banner(`Makes an authenticated HTTP request to the GitHub API and prints the response.

The endpoint argument should either be a path of a GitHub API v3 endpoint, or
`+"`"+`graphql`+"`"+` to access the GitHub API v4.

Placeholder values `+"`"+`{owner}`+"`"+`, `+"`"+`{repo}`+"`"+`, and `+"`"+`{branch}`+"`"+` in the endpoint
argument will get replaced with values from the repository of the current
directory or the repository specified in the `+"`"+`GH_REPO`+"`"+` environment variable.
Note that in some shells, for example PowerShell, you may need to enclose
any value that contains `+"`"+`{...}`+"`"+` in quotes to prevent the shell from
applying special meaning to curly braces.

The default HTTP request method is `+"`"+`GET`+"`"+` normally and `+"`"+`POST`+"`"+` if any parameters
were added. Override the method with `+"`"+`--method`+"`"+`.

Pass one or more `+"`"+`-f/--raw-field`+"`"+` values in `+"`"+`key=value`+"`"+` format to add static string
parameters to the request payload. To add non-string or placeholder-determined values, see
`+"`"+`-F/--field`+"`"+` below. Note that adding request parameters will automatically switch the
request method to `+"`"+`POST`+"`"+`. To send the parameters as a `+"`"+`GET`+"`"+` query string instead, use
`+"`"+`--method GET`+"`"+`.

The `+"`"+`-F/--field`+"`"+` flag has magic type conversion based on the format of the value:

- literal values `+"`"+`true`+"`"+`, `+"`"+`false`+"`"+`, `+"`"+`null`+"`"+`, and integer numbers get converted to
  appropriate JSON types;
- placeholder values `+"`"+`{owner}`+"`"+`, `+"`"+`{repo}`+"`"+`, and `+"`"+`{branch}`+"`"+` get populated with values
  from the repository of the current directory;
- if the value starts with `+"`"+`@`+"`"+`, the rest of the value is interpreted as a
  filename to read the value from. Pass `+"`"+`-`+"`"+` to read from standard input.

For GraphQL requests, all fields other than `+"`"+`query`+"`"+` and `+"`"+`operationName`+"`"+` are
interpreted as GraphQL variables.

To pass nested parameters in the request payload, use `+"`"+`key[subkey]=value`+"`"+` syntax when
declaring fields. To pass nested values as arrays, declare multiple fields with the
syntax `+"`"+`key[]=value1`+"`"+`, `+"`"+`key[]=value2`+"`"+`. To pass an empty array, use `+"`"+`key[]`+"`"+` without a
value.

To pass pre-constructed JSON or payloads in other formats, a request body may be read
from file specified by `+"`"+`--input`+"`"+`. Use `+"`"+`-`+"`"+` to read from standard input. When passing the
request body this way, any parameters specified via field flags are added to the query
string of the endpoint URL.

In `+"`"+`--paginate`+"`"+` mode, all pages of results will sequentially be requested until
there are no more pages of results. For GraphQL requests, this requires that the
original query accepts an `+"`"+`$endCursor: String`+"`"+` variable and that it fetches the
`+"`"+`pageInfo{ hasNextPage, endCursor }`+"`"+` set of fields from a collection. Each page is a separate
JSON array or object. Pass `+"`"+`--slurp`+"`"+` to wrap all pages of JSON arrays or objects
into an outer JSON array.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
			ox.Spec(`<endpoint> [flags]`),
			ox.Example(`
  # List releases in the current repository
  $ gh api repos/{owner}/{repo}/releases
  
  # Post an issue comment
  $ gh api repos/{owner}/{repo}/issues/123/comments -f body='Hi from CLI'
  
  # Post nested parameter read from a file
  $ gh api gists -F 'files[myfile.txt][content]=@myfile.txt'
  
  # Add parameters to a GET request
  $ gh api -X GET search/issues -f q='repo:cli/cli is:open remote'
  
  # Set a custom HTTP header
  $ gh api -H 'Accept: application/vnd.github.v3.raw+json' ...
  
  # Opt into GitHub API previews
  $ gh api --preview baptiste,nebula ...
  
  # Print only specific fields from the response
  $ gh api repos/{owner}/{repo}/issues --jq '.[].title'
  
  # Use a template for the output
  $ gh api repos/{owner}/{repo}/issues --template \
    '{{range .}}{{.title}} ({{.labels | pluck "name" | join ", " | color "yellow"}}){{"\n"}}{{end}}'
  
  # Update allowed values of the "environment" custom property in a deeply nested array
  $ gh api -X PATCH /orgs/{org}/properties/schema \
     -F 'properties[][property_name]=environment' \
     -F 'properties[][default_value]=production' \
     -F 'properties[][allowed_values][]=staging' \
     -F 'properties[][allowed_values][]=production'
  
  # List releases with GraphQL
  $ gh api graphql -F owner='{owner}' -F name='{repo}' -f query='
    query($name: String!, $owner: String!) {
      repository(owner: $owner, name: $name) {
        releases(last: 3) {
          nodes { tagName }
        }
      }
    }
  '
  
  # List all repositories for a user
  $ gh api graphql --paginate -f query='
    query($endCursor: String) {
      viewer {
        repositories(first: 100, after: $endCursor) {
          nodes { nameWithOwner }
          pageInfo {
            hasNextPage
            endCursor
          }
        }
      }
    }
  '
  
  # Get the percentage of forks for the current user
  $ gh api graphql --paginate --slurp -f query='
    query($endCursor: String) {
      viewer {
        repositories(first: 100, after: $endCursor) {
          nodes { isFork }
          pageInfo {
            hasNextPage
            endCursor
          }
        }
      }
    }
  ' | jq 'def count(e): reduce e as $_ (0;.+1);
  [.[].data.viewer.repositories.nodes[]] as $r | count(select($r[].isFork))/count($r[])'`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				Duration(`cache`, `Cache the response, e.g. "3600s", "60m", "1h"`, ox.Section(0)).
				Map(`field`, `Add a typed parameter in key=value format`, ox.Spec(`key=value`), ox.Short("F"), ox.Section(0)).
				String(`header`, `Add a HTTP request header in key:value format`, ox.Spec(`key:value`), ox.Short("H"), ox.Section(0)).
				String(`hostname`, `The GitHub hostname for the request`, ox.Default("github.com"), ox.Section(0)).
				Bool(`include`, `Include HTTP response status line and headers in the output`, ox.Short("i"), ox.Section(0)).
				String(`input`, `The file to use as body for the HTTP request (use "-" to read from standard input)`, ox.Spec(`file`), ox.Section(0)).
				String(`jq`, `Query to select values from the response using jq syntax`, ox.Short("q"), ox.Section(0)).
				String(`method`, `The HTTP method for the request`, ox.Default("GET"), ox.Short("X"), ox.Section(0)).
				Bool(`paginate`, `Make additional HTTP requests to fetch all pages of results`, ox.Section(0)).
				Slice(`preview`, `GitHub API preview names to request (without the "-preview" suffix)`, ox.Short("p"), ox.Section(0)).
				Map(`raw-field`, `Add a string parameter in key=value format`, ox.Spec(`key=value`), ox.Short("f"), ox.Section(0)).
				Bool(`silent`, `Do not print the response body`, ox.Section(0)).
				Bool(`slurp`, `Use with "--paginate" to return an array of all pages of either JSON arrays or objects`, ox.Section(0)).
				String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
				Bool(`verbose`, `Include full HTTP request and response in the output`, ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`attestation`, `Work with artifact attestations`),
			ox.Banner(`Download and verify artifact attestations.`),
			ox.Spec(`[subcommand] [flags]`),
			ox.Aliases("at"),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`download`, `Download an artifact's attestations for offline use`),
				ox.Banner(`### NOTE: This feature is currently in public preview, and subject to change.

Download attestations associated with an artifact for offline use.

The command requires either:
* a file path to an artifact, or
* a container image URI (e.g. `+"`"+`oci://<image-uri>`+"`"+`)
  * (note that if you provide an OCI URL, you must already be authenticated with
its container registry)

In addition, the command requires either:
* the `+"`"+`--repo`+"`"+` flag (e.g. --repo github/example).
* the `+"`"+`--owner`+"`"+` flag (e.g. --owner github), or

The `+"`"+`--repo`+"`"+` flag value must match the name of the GitHub repository
that the artifact is linked with.

The `+"`"+`--owner`+"`"+` flag value must match the name of the GitHub organization
that the artifact's linked repository belongs to.

Any associated bundle(s) will be written to a file in the
current directory named after the artifact's digest. For example, if the
digest is "sha256:1234", the file will be named "sha256:1234.jsonl".

Colons are special characters on Windows and cannot be used in
file names. To accommodate, a dash will be used to separate the algorithm
from the digest in the attestations file name. For example, if the digest
is "sha256:1234", the file will be named "sha256-1234.jsonl".`),
				ox.Spec(`[<file-path> | oci://<image-uri>] [--owner | --repo] [flags]`),
				ox.Example(`
  # Download attestations for a local artifact linked with an organization
  $ gh attestation download example.bin -o github
  
  # Download attestations for a local artifact linked with a repository
  $ gh attestation download example.bin -R github/example
  
  # Download attestations for an OCI image linked with an organization
  $ gh attestation download oci://example.com/foo/bar:latest -o github`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`digest-alg`, `The algorithm used to compute a digest of the artifact: {sha256|sha512}`, ox.Default("sha256"), ox.Short("d"), ox.Section(0)).
					String(`hostname`, `Configure host to use`, ox.Section(0)).
					Int(`limit`, `Maximum number of attestations to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`owner`, `GitHub organization to scope attestation lookup by`, ox.Short("o"), ox.Section(0)).
					String(`predicate-type`, `Filter attestations by provided predicate type`, ox.Section(0)).
					String(`repo`, `Repository name in the format <owner>/<repo>`, ox.Short("R"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`trusted-root`, `Output trusted_root.jsonl contents, likely for offline verification`),
				ox.Banner(`### NOTE: This feature is currently in public preview, and subject to change.

Output contents for a trusted_root.jsonl file, likely for offline verification.

When using `+"`"+`gh attestation verify`+"`"+`, if your machine is on the internet,
this will happen automatically. But to do offline verification, you need to
supply a trusted root file with `+"`"+`--custom-trusted-root`+"`"+`; this command
will help you fetch a `+"`"+`trusted_root.jsonl`+"`"+` file for that purpose.

You can call this command without any flags to get a trusted root file covering
the Sigstore Public Good Instance as well as GitHub's Sigstore instance.

Otherwise you can use `+"`"+`--tuf-url`+"`"+` to specify the URL of a custom TUF
repository mirror, and `+"`"+`--tuf-root`+"`"+` should be the path to the
`+"`"+`root.json`+"`"+` file that you securely obtained out-of-band.

If you just want to verify the integrity of your local TUF repository, and don't
want the contents of a trusted_root.jsonl file, use `+"`"+`--verify-only`+"`"+`.`),
				ox.Spec(`[--tuf-url <url> --tuf-root <file-path>] [--verify-only] [flags]`),
				ox.Example(`
  # Get a trusted_root.jsonl for both Sigstore Public Good and GitHub's instance
  $ gh attestation trusted-root`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`hostname`, `Configure host to use`, ox.Section(0)).
					String(`tuf-root`, `Path to the TUF root.json file on disk`, ox.Section(0)).
					String(`tuf-url`, `URL to the TUF repository mirror`, ox.Section(0)).
					Bool(`verify-only`, `Don't output trusted_root.jsonl contents`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`verify`, `Verify an artifact's integrity using attestations`),
				ox.Banner(`Verify the integrity and provenance of an artifact using its associated
cryptographically signed attestations.

## Understanding Verification

An attestation is a claim (i.e. a provenance statement) made by an actor
(i.e. a GitHub Actions workflow) regarding a subject (i.e. an artifact).

In order to verify an attestation, you must provide an artifact and validate:
* the identity of the actor that produced the attestation
* the expected attestation predicate type (the nature of the claim)

By default, this command enforces the `+"`"+`https://slsa.dev/provenance/v1`+"`"+`
predicate type. To verify other attestation predicate types use the
`+"`"+`--predicate-type`+"`"+` flag.

The "actor identity" consists of:
* the repository or the repository owner the artifact is linked with
* the Actions workflow that produced the attestation (a.k.a the
  signer workflow)

This identity is then validated against the attestation's certificate's
SourceRepository, SourceRepositoryOwner, and SubjectAlternativeName
(SAN) fields, among others.

It is up to you to decide how precisely you want to enforce this identity.

At a minimum, this command requires either:
* the `+"`"+`--owner`+"`"+` flag (e.g. --owner github), or
* the `+"`"+`--repo`+"`"+` flag (e.g. --repo github/example)

The more precisely you specify the identity, the more control you will
have over the security guarantees offered by the verification process.

Ideally, the path of the signer workflow is also validated using the
`+"`"+`--signer-workflow`+"`"+` or `+"`"+`--cert-identity`+"`"+` flags.

Please note: if your attestation was generated via a reusable workflow then
that reusable workflow is the signer whose identity needs to be validated.
In this situation, you must use either the `+"`"+`--signer-workflow`+"`"+` or
the `+"`"+`--signer-repo`+"`"+` flag.

For more options, see the other available flags.

## Loading Artifacts And Attestations

To specify the artifact, this command requires:
* a file path to an artifact, or
* a container image URI (e.g. `+"`"+`oci://<image-uri>`+"`"+`)
  * (note that if you provide an OCI URL, you must already be authenticated with
its container registry)

By default, this command will attempt to fetch relevant attestations via the
GitHub API using the values provided to `+"`"+`--owner`+"`"+` or  `+"`"+`--repo`+"`"+`.

To instead fetch attestations from your artifact's OCI registry, use the
`+"`"+`--bundle-from-oci`+"`"+` flag.

For offline verification using attestations stored on disk (c.f. the download command)
provide a path to the `+"`"+`--bundle`+"`"+` flag.

## Additional Policy Enforcement

Given the `+"`"+`--format=json`+"`"+` flag, upon successful verification this
command will output a JSON array containing one entry per verified attestation.

This output can then be used for additional policy enforcement, i.e. by being
piped into a policy engine.

Each object in the array contains two properties:
* an `+"`"+`attestation`+"`"+` object, which contains the bundle that was verified
* a `+"`"+`verificationResult`+"`"+` object, which is a parsed representation of the
  contents of the bundle that was verified.

Within the `+"`"+`verificationResult`+"`"+` object you will find:
* `+"`"+`signature.certificate`+"`"+`, which is a parsed representation of the X.509
  certificate embedded in the attestation,
* `+"`"+`verifiedTimestamps`+"`"+`, an array of objects denoting when the attestation
  was witnessed by a transparency log or a timestamp authority
* `+"`"+`statement`+"`"+`, which contains the `+"`"+`subject`+"`"+` array referencing artifacts,
  the `+"`"+`predicateType`+"`"+` field, and the `+"`"+`predicate`+"`"+` object which contains
  additional, often user-controllable, metadata

IMPORTANT: please note that only the `+"`"+`signature.certificate`+"`"+` and the
`+"`"+`verifiedTimestamps`+"`"+` properties contain values that cannot be
manipulated by the workflow that originated the attestation.

When dealing with attestations created within GitHub Actions, the contents of
`+"`"+`signature.certificate`+"`"+` are populated directly from the OpenID Connect
token that GitHub has generated. The contents of the `+"`"+`verifiedTimestamps`+"`"+`
array are populated from the signed timestamps originating from either a
transparency log or a timestamp authority – and likewise cannot be forged by users.

When designing policy enforcement using this output, special care must be taken
when examining the contents of the `+"`"+`statement.predicate`+"`"+` property:
should an attacker gain access to your workflow's execution context, they
could then falsify the contents of the `+"`"+`statement.predicate`+"`"+`.

To mitigate this attack vector, consider using a "trusted builder": when generating
an artifact, have the build and attestation signing occur within a reusable workflow
whose execution cannot be influenced by input provided through the caller workflow.

See above re: `+"`"+`--signer-workflow`+"`"+`.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<file-path> | oci://<image-uri>] [--owner | --repo] [flags]`),
				ox.Example(`
  # Verify an artifact linked with a repository
  $ gh attestation verify example.bin --repo github/example
  
  # Verify an artifact linked with an organization
  $ gh attestation verify example.bin --owner github
  
  # Verify an artifact and output the full verification result
  $ gh attestation verify example.bin --owner github --format json
  
  # Verify an OCI image using attestations stored on disk
  $ gh attestation verify oci://<image-uri> --owner github --bundle sha256:foo.jsonl
  
  # Verify an artifact signed with a reusable workflow
  $ gh attestation verify example.bin --owner github --signer-repo actions/example`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`bundle`, `Path to bundle on disk, either a single bundle in a JSON file or a JSON lines file with multiple bundles`, ox.Short("b"), ox.Section(0)).
					Bool(`bundle-from-oci`, `When verifying an OCI image, fetch the attestation bundle from the OCI registry instead of from GitHub`, ox.Section(0)).
					String(`cert-identity`, `Enforce that the certificate's SubjectAlternativeName matches the provided value exactly`, ox.Section(0)).
					String(`cert-identity-regex`, `Enforce that the certificate's SubjectAlternativeName matches the provided regex`, ox.Short("i"), ox.Section(0)).
					String(`cert-oidc-issuer`, `Enforce that the issuer of the OIDC token matches the provided value`, ox.Default("https://to$USER.actions.githubusercontent.com"), ox.Section(0)).
					String(`custom-trusted-root`, `Path to a trusted_root.jsonl file; likely for offline verification`, ox.Section(0)).
					Bool(`deny-self-hosted-runners`, `Fail verification for attestations generated on self-hosted runners`, ox.Section(0)).
					String(`digest-alg`, `The algorithm used to compute a digest of the artifact: {sha256|sha512}`, ox.Default("sha256"), ox.Short("d"), ox.Section(0)).
					String(`format`, `Output format: {json}`, ox.Section(0)).
					String(`hostname`, `Configure host to use`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Int(`limit`, `Maximum number of attestations to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Bool(`no-public-good`, `Do not verify attestations signed with Sigstore public good instance`, ox.Section(0)).
					String(`owner`, `GitHub organization to scope attestation lookup by`, ox.Short("o"), ox.Section(0)).
					String(`predicate-type`, `Enforce that verified attestations' predicate type matches the provided value`, ox.Default("https://slsa.dev/provenance/v1"), ox.Section(0)).
					String(`repo`, `Repository name in the format <owner>/<repo>`, ox.Short("R"), ox.Section(0)).
					String(`signer-digest`, `Enforce that the digest associated with the signer workflow matches the provided value`, ox.Section(0)).
					String(`signer-repo`, `Enforce that the workflow that signed the attestation's repository matches the provided value (<owner>/<repo>)`, ox.Section(0)).
					String(`signer-workflow`, `Enforce that the workflow that signed the attestation matches the provided value ([host/]<owner>/<repo>/<path>/<to>/<workflow>)`, ox.Section(0)).
					String(`source-digest`, `Enforce that the digest associated with the source repository matches the provided value`, ox.Section(0)).
					String(`source-ref`, `Enforce that the git ref associated with the source repository matches the provided value`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`config`, `Manage configuration for gh`),
			ox.Banner(`Display or change configuration settings for gh.

Current respected settings:

- `+"`"+`git_protocol`+"`"+`: the protocol to use for git clone and push operations {https|ssh} (default https)
- `+"`"+`editor`+"`"+`: the text editor program to use for authoring text
- `+"`"+`prompt`+"`"+`: toggle interactive prompting in the terminal {enabled|disabled} (default enabled)
- `+"`"+`prefer_editor_prompt`+"`"+`: toggle preference for editor-based interactive prompting in the terminal {enabled|disabled} (default disabled)
- `+"`"+`pager`+"`"+`: the terminal pager program to send standard output to
- `+"`"+`http_unix_socket`+"`"+`: the path to a Unix socket through which to make an HTTP connection
- `+"`"+`browser`+"`"+`: the web browser to use for opening URLs
- `+"`"+`color_labels`+"`"+`: whether to display labels using their RGB hex color codes in terminals that support truecolor {enabled|disabled} (default disabled)
- `+"`"+`accessible_colors`+"`"+`: whether customizable, 4-bit accessible colors should be used {enabled|disabled} (default disabled)
- `+"`"+`accessible_prompter`+"`"+`: whether an accessible prompter should be used {enabled|disabled} (default disabled)
- `+"`"+`spinner`+"`"+`: whether to use a animated spinner as a progress indicator {enabled|disabled} (default enabled)`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`clear-cache`, `Clear the cli cache`),
				ox.Banner(`Clear the cli cache`),
				ox.Spec(`[flags]`),
				ox.Example(`
  # Clear the cli cache
  $ gh config clear-cache`),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`get`, `Print the value of a given configuration key`),
				ox.Banner(`Print the value of a given configuration key`),
				ox.Spec(`<key> [flags]`),
				ox.Example(`
  $ gh config get git_protocol`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`host`, `Get per-host setting`, ox.Short("h"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `Print a list of configuration keys and values`),
				ox.Banner(`Print a list of configuration keys and values`),
				ox.Spec(`[flags]`),
				ox.Aliases("config ls"),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`host`, `Get per-host configuration`, ox.Short("h"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`set`, `Update configuration with a value for the given key`),
				ox.Banner(`Update configuration with a value for the given key`),
				ox.Spec(`<key> <value> [flags]`),
				ox.Example(`
  $ gh config set editor vim
  $ gh config set editor "code --wait"
  $ gh config set git_protocol ssh --host github.com
  $ gh config set prompt disabled`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`host`, `Set per-host setting`, ox.Short("h"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`extension`, `Manage gh extensions`),
			ox.Banner(`GitHub CLI extensions are repositories that provide additional gh commands.

The name of the extension repository must start with `+"`"+`gh-`+"`"+` and it must contain an
executable of the same name. All arguments passed to the `+"`"+`gh <extname>`+"`"+` invocation
will be forwarded to the `+"`"+`gh-<extname>`+"`"+` executable of the extension.

An extension cannot override any of the core gh commands. If an extension name conflicts
with a core gh command, you can use `+"`"+`gh extension exec <extname>`+"`"+`.

When an extension is executed, gh will check for new versions once every 24 hours and display
an upgrade notice. See `+"`"+`gh help environment`+"`"+` for information on disabling extension notices.

For the list of available extensions, see <https://github.com/topics/gh-extension>.`),
			ox.Spec(`[flags]`),
			ox.Aliases("extensions", "ext"),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`browse`, `Enter a UI for browsing, adding, and removing extensions`),
				ox.Banner(`This command will take over your terminal and run a fully interactive
interface for browsing, adding, and removing gh extensions. A terminal
width greater than 100 columns is recommended.

To learn how to control this interface, press `+"`"+`?`+"`"+` after running to see
the help text.

Press `+"`"+`q`+"`"+` to quit.

Running this command with `+"`"+`--single-column`+"`"+` should make this command
more intelligible for users who rely on assistive technology like screen
readers or high zoom.

For a more traditional way to discover extensions, see:

	gh ext search

along with `+"`"+`gh ext install`+"`"+`, `+"`"+`gh ext remove`+"`"+`, and `+"`"+`gh repo view`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`debug`, `Log to /tmp/extBrowse-*`, ox.Section(0)).
					Bool(`single-column`, `Render TUI with only one column of text`, ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`create`, `Create a new extension`),
				ox.Banner(`Create a new extension`),
				ox.Spec(`[<name>] [flags]`),
				ox.Example(`
  # Use interactively
  $ gh extension create
  
  # Create a script-based extension
  $ gh extension create foobar
  
  # Create a Go extension
  $ gh extension create --precompiled=go foobar
  
  # Create a non-Go precompiled extension
  $ gh extension create --precompiled=other foobar`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`precompiled`, `Create a precompiled extension. Possible values: go, other`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`exec`, `Execute an installed extension`),
				ox.Banner(`Execute an extension using the short name. For example, if the extension repository is
`+"`"+`owner/gh-extension`+"`"+`, you should pass `+"`"+`extension`+"`"+`. You can use this command when
the short name conflicts with a core gh command.

All arguments after the extension name will be forwarded to the executable
of the extension.`),
				ox.Spec(`<name> [args] [flags]`),
				ox.Example(`
  # Execute a label extension instead of the core gh label command
  $ gh extension exec label`),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`install`, `Install a gh extension from a repository`),
				ox.Banner(`Install a GitHub CLI extension from a GitHub or local repository.

For GitHub repositories, the repository argument can be specified in
`+"`"+`OWNER/REPO`+"`"+` format or as a full repository URL.
The URL format is useful when the repository is not hosted on <github.com>.

For remote repositories, the GitHub CLI first looks for the release artifacts assuming
that it's a binary extension i.e. prebuilt binaries provided as part of the release.
In the absence of a release, the repository itself is cloned assuming that it's a
script extension i.e. prebuilt executable or script exists on its root.

The `+"`"+`--pin`+"`"+` flag may be used to specify a tag or commit for binary and script
extensions respectively, the latest version is used otherwise.

For local repositories, often used while developing extensions, use `+"`"+`.`+"`"+` as the
value of the repository argument. Note the following:

- After installing an extension from a locally cloned repository, the GitHub CLI will
manage this extension as a symbolic link (or equivalent mechanism on Windows) pointing
to an executable file with the same name as the repository in the repository's root.
For example, if the repository is named `+"`"+`gh-foobar`+"`"+`, the symbolic link will point
to `+"`"+`gh-foobar`+"`"+` in the extension repository's root.
- When executing the extension, the GitHub CLI will run the executable file found
by following the symbolic link. If no executable file is found, the extension
will fail to execute.
- If the extension is precompiled, the executable file must be built manually and placed
in the repository's root.

For the list of available extensions, see <https://github.com/topics/gh-extension>.`),
				ox.Spec(`<repository> [flags]`),
				ox.Example(`
  # Install an extension from a remote repository hosted on GitHub
  $ gh extension install owner/gh-extension
  
  # Install an extension from a remote repository via full URL
  $ gh extension install https://my.ghes.com/owner/gh-extension
  
  # Install an extension from a local repository in the current working directory
  $ gh extension install .`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`force`, `Force upgrade extension, or ignore if latest already installed`, ox.Section(0)).
					String(`pin`, `Pin extension to a release tag or commit ref`, ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `List installed extension commands`),
				ox.Banner(`List installed extension commands`),
				ox.Spec(`[flags]`),
				ox.Aliases("ext ls", "extension ls", "extensions ls"),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`remove`, `Remove an installed extension`),
				ox.Banner(`Remove an installed extension`),
				ox.Spec(`<name> [flags]`),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
			ox.Sub(
				ox.Usage(`search`, `Search extensions to the GitHub CLI`),
				ox.Banner(`Search for gh extensions.

With no arguments, this command prints out the first 30 extensions
available to install sorted by number of stars. More extensions can
be fetched by specifying a higher limit with the `+"`"+`--limit`+"`"+` flag.

When connected to a terminal, this command prints out three columns.
The first has a ✓ if the extension is already installed locally. The
second is the full name of the extension repository in `+"`"+`OWNER/REPO`+"`"+`
format. The third is the extension's description.

When not connected to a terminal, the ✓ character is rendered as the
word "installed" but otherwise the order and content of the columns
are the same.

This command behaves similarly to `+"`"+`gh search repos`+"`"+` but does not
support as many search qualifiers. For a finer grained search of
extensions, try using:

	gh search repos --topic "gh-extension"

and adding qualifiers as needed. See `+"`"+`gh help search repos`+"`"+` to learn
more about repository search.

For listing just the extensions that are already installed locally,
see:

	gh ext list

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<query>] [flags]`),
				ox.Example(`
  # List the first 30 extensions sorted by star count, descending
  $ gh ext search
  
  # List more extensions
  $ gh ext search --limit 300
  
  # List extensions matching the term "branch"
  $ gh ext search branch
  
  # List extensions owned by organization "github"
  $ gh ext search --owner github
  
  # List extensions, sorting by recently updated, ascending
  $ gh ext search --sort updated --order asc
  
  # List extensions, filtering by license
  $ gh ext search --license MIT
  
  # Open search results in the browser
  $ gh ext search -w`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Slice(`license`, `Filter based on license type`, ox.Section(0)).
					Int(`limit`, `Maximum number of extensions to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`order`, `Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}`, ox.Default("desc"), ox.Section(0)).
					Slice(`owner`, `Filter on owner`, ox.Section(0)).
					String(`sort`, `Sort fetched repositories: {forks|help-wanted-issues|stars|updated}`, ox.Default("best-match"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`upgrade`, `Upgrade installed extensions`),
				ox.Banner(`Upgrade installed extensions`),
				ox.Spec(`{<name> | --all} [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`all`, `Upgrade all extensions`, ox.Section(0)).
					Bool(`dry-run`, `Only display upgrades`, ox.Section(0)).
					Bool(`force`, `Force upgrade extension`, ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`gpg-key`, `Manage GPG keys`),
			ox.Banner(`Manage GPG keys registered with your GitHub account.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`add`, `Add a GPG key to your GitHub account`),
				ox.Banner(`Add a GPG key to your GitHub account`),
				ox.Spec(`[<key-file>] [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`title`, `Title for the new key`, ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a GPG key from your GitHub account`),
				ox.Banner(`Delete a GPG key from your GitHub account`),
				ox.Spec(`<key-id> [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `Lists GPG keys in your GitHub account`),
				ox.Banner(`Lists GPG keys in your GitHub account`),
				ox.Spec(`[flags]`),
				ox.Aliases("gpg-key ls"),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
		),
		ox.Sub(
			ox.Usage(`label`, `Manage labels`),
			ox.Banner(`Work with GitHub labels.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`clone`, `Clones labels from one repository to another`),
				ox.Banner(`Clones labels from a source repository to a destination repository on GitHub.
By default, the destination repository is the current repository.

All labels from the source repository will be copied to the destination
repository. Labels in the destination repository that are not in the source
repository will not be deleted or modified.

Labels from the source repository that already exist in the destination
repository will be skipped. You can overwrite existing labels in the
destination repository using the `+"`"+`--force`+"`"+` flag.`),
				ox.Spec(`<source-repository> [flags]`),
				ox.Example(`
  # Clone and overwrite labels from cli/cli repository into the current repository
  $ gh label clone cli/cli --force
  
  # Clone labels from cli/cli repository into a octocat/cli repository
  $ gh label clone cli/cli --repo octocat/cli`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`force`, `Overwrite labels in the destination repository`, ox.Short("f"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`create`, `Create a new label`),
				ox.Banner(`Create a new label on GitHub, or update an existing one with `+"`"+`--force`+"`"+`.

Must specify name for the label. The description and color are optional.
If a color isn't provided, a random one will be chosen.

The label color needs to be 6 character hex value.`),
				ox.Spec(`<name> [flags]`),
				ox.Example(`
  # Create new bug label
  $ gh label create bug --description "Something isn't working" --color E99695`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`color`, `Color of the label`, ox.Short("c"), ox.Section(0)).
					String(`description`, `Description of the label`, ox.Short("d"), ox.Section(0)).
					Bool(`force`, `Update the label color and description if label already exists`, ox.Short("f"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete a label from a repository`),
				ox.Banner(`Delete a label from a repository`),
				ox.Spec(`<name> [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Confirm deletion without prompting`, ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`edit`, `Edit a label`),
				ox.Banner(`Update a label on GitHub.

A label can be renamed using the `+"`"+`--name`+"`"+` flag.

The label color needs to be 6 character hex value.`),
				ox.Spec(`<name> [flags]`),
				ox.Example(`
  # Update the color of the bug label
  $ gh label edit bug --color FF0000
  
  # Rename and edit the description of the bug label
  $ gh label edit bug --name big-bug --description "Bigger than normal bug"`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`color`, `Color of the label`, ox.Short("c"), ox.Section(0)).
					String(`description`, `Description of the label`, ox.Short("d"), ox.Section(0)).
					String(`name`, `New name of the label`, ox.Short("n"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List labels in a repository`),
				ox.Banner(`Display labels in a GitHub repository.

When using the `+"`"+`--search`+"`"+` flag results are sorted by best match of the query.
This behavior cannot be configured with the `+"`"+`--order`+"`"+` or `+"`"+`--sort`+"`"+` flags.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("label ls"),
				ox.Example(`
  # Sort labels by name
  $ gh label list --sort name
  
  # Find labels with "bug" in the name or description
  $ gh label list --search bug`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of labels to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`order`, `Order of labels returned: {asc|desc}`, ox.Default("asc"), ox.Section(0)).
					String(`search`, `Search label names and descriptions`, ox.Short("S"), ox.Section(0)).
					String(`sort`, `Sort fetched labels: {created|name}`, ox.Default("created"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `List labels in the web browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`ruleset`, `View info about repo rulesets`),
			ox.Banner(`Repository rulesets are a way to define a set of rules that apply to a repository.
These commands allow you to view information about them.`),
			ox.Spec(`<command> [flags]`),
			ox.Aliases("rs"),
			ox.Example(`
  $ gh ruleset list
  $ gh ruleset view --repo OWNER/REPO --web
  $ gh ruleset check branch-name`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`check`, `View rules that would apply to a given branch`),
				ox.Banner(`View information about GitHub rules that apply to a given branch.

The provided branch name does not need to exist; rules will be displayed that would apply
to a branch with that name. All rules are returned regardless of where they are configured.

If no branch name is provided, then the current branch will be used.

The `+"`"+`--default`+"`"+` flag can be used to view rules that apply to the default branch of the
repository.`),
				ox.Spec(`[<branch>] [flags]`),
				ox.Example(`
  # View all rules that apply to the current branch
  $ gh ruleset check
  
  # View all rules that apply to a branch named "my-branch" in a different repository
  $ gh ruleset check my-branch --repo owner/repo
  
  # View all rules that apply to the default branch in a different repository
  $ gh ruleset check --default --repo owner/repo
  
  # View a ruleset configured in a different repository or any of its parents
  $ gh ruleset view 23 --repo owner/repo
  
  # View an organization-level ruleset
  $ gh ruleset view 23 --org my-org`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`default`, `Check rules on default branch`, ox.Section(0)).
					Bool(`web`, `Open the branch rules page in a web browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List rulesets for a repository or organization`),
				ox.Banner(`List GitHub rulesets for a repository or organization.

If no options are provided, the current repository's rulesets are listed. You can query a different
repository's rulesets by using the `+"`"+`--repo`+"`"+` flag. You can also use the `+"`"+`--org`+"`"+` flag to list rulesets
configured for the provided organization.

Use the `+"`"+`--parents`+"`"+` flag to control whether rulesets configured at higher levels that also apply to the provided
repository or organization should be returned. The default is `+"`"+`true`+"`"+`.

Your access token must have the `+"`"+`admin:org`+"`"+` scope to use the `+"`"+`--org`+"`"+` flag, which can be granted by running `+"`"+`gh auth refresh -s admin:org`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("rs ls", "ruleset ls"),
				ox.Example(`
  # List rulesets in the current repository
  $ gh ruleset list
  
  # List rulesets in a different repository, including those configured at higher levels
  $ gh ruleset list --repo owner/repo --parents
  
  # List rulesets in an organization
  $ gh ruleset list --org org-name`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Int(`limit`, `Maximum number of rulesets to list`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					String(`org`, `List organization-wide rulesets for the provided organization`, ox.Short("o"), ox.Section(0)).
					Bool(`parents`, `Whether to include rulesets configured at higher levels that also apply`, ox.Default("true"), ox.Short("p"), ox.Section(0)).
					Bool(`web`, `Open the list of rulesets in the web browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`view`, `View information about a ruleset`),
				ox.Banner(`View information about a GitHub ruleset.

If no ID is provided, an interactive prompt will be used to choose
the ruleset to view.

Use the `+"`"+`--parents`+"`"+` flag to control whether rulesets configured at higher
levels that also apply to the provided repository or organization should
be returned. The default is `+"`"+`true`+"`"+`.`),
				ox.Spec(`[<ruleset-id>] [flags]`),
				ox.Example(`
  # Interactively choose a ruleset to view from all rulesets that apply to the current repository
  $ gh ruleset view
  
  # Interactively choose a ruleset to view from only rulesets configured in the current repository
  $ gh ruleset view --no-parents
  
  # View a ruleset configured in the current repository or any of its parents
  $ gh ruleset view 43
  
  # View a ruleset configured in a different repository or any of its parents
  $ gh ruleset view 23 --repo owner/repo
  
  # View an organization-level ruleset
  $ gh ruleset view 23 --org my-org`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`org`, `Organization name if the provided ID is an organization-level ruleset`, ox.Short("o"), ox.Section(0)).
					Bool(`parents`, `Whether to include rulesets configured at higher levels that also apply`, ox.Default("true"), ox.Short("p"), ox.Section(0)).
					Bool(`web`, `Open the ruleset in the browser`, ox.Short("w"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`search`, `Search for repositories, issues, and pull requests`),
			ox.Banner(`Search across all of GitHub.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`code`, `Search within code`),
				ox.Banner(`Search within code in GitHub repositories.

The search syntax is documented at:
<https://docs.github.com/search-github/searching-on-github/searching-code>

Note that these search results are powered by what is now a legacy GitHub code search engine.
The results might not match what is seen on <github.com>, and new features like regex search
are not yet available via the GitHub API.

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`<query> [flags]`),
				ox.Example(`
  # Search code matching "react" and "lifecycle"
  $ gh search code react lifecycle
  
  # Search code matching "error handling"
  $ gh search code "error handling"
  
  # Search code matching "deque" in Python files
  $ gh search code deque --language=python
  
  # Search code matching "cli" in repositories owned by microsoft organization
  $ gh search code cli --owner=microsoft
  
  # Search code matching "panic" in the GitHub CLI repository
  $ gh search code panic --repo cli/cli
  
  # Search code matching keyword "lint" in package.json files
  $ gh search code lint --filename package.json`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`extension`, `Filter on file extension`, ox.Section(0)).
					String(`filename`, `Filter on filename`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`language`, `Filter results by language`, ox.Section(0)).
					Int(`limit`, `Maximum number of code results to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Slice(`match`, `Restrict search to file contents or file path: {file|path}`, ox.Section(0)).
					Slice(`owner`, `Filter on owner`, ox.Section(0)).
					Slice(`repo`, `Filter on repository`, ox.Short("R"), ox.Section(0)).
					String(`size`, `Filter on size range, in kilobytes`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`commits`, `Search for commits`),
				ox.Banner(`Search for commits on GitHub.

The command supports constructing queries using the GitHub search syntax,
using the parameter and qualifier flags, or a combination of the two.

GitHub search syntax is documented at:
<https://docs.github.com/search-github/searching-on-github/searching-commits>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<query>] [flags]`),
				ox.Example(`
  # Search commits matching set of keywords "readme" and "typo"
  $ gh search commits readme typo
  
  # Search commits matching phrase "bug fix"
  $ gh search commits "bug fix"
  
  # Search commits committed by user "monalisa"
  $ gh search commits --committer=monalisa
  
  # Search commits authored by users with name "Jane Doe"
  $ gh search commits --author-name="Jane Doe"
  
  # Search commits matching hash "8dd03144ffdc6c0d486d6b705f9c7fba871ee7c3"
  $ gh search commits --hash=8dd03144ffdc6c0d486d6b705f9c7fba871ee7c3
  
  # Search commits authored before February 1st, 2022
  $ gh search commits --author-date="<2022-02-01"`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`author`, `Filter by author`, ox.Section(0)).
					Date(`author-date`, `Filter based on authored date`, ox.Section(0)).
					String(`author-email`, `Filter on author email`, ox.Section(0)).
					String(`author-name`, `Filter on author name`, ox.Section(0)).
					String(`committer`, `Filter by committer`, ox.Section(0)).
					Date(`committer-date`, `Filter based on committed date`, ox.Section(0)).
					String(`committer-email`, `Filter on committer email`, ox.Section(0)).
					String(`committer-name`, `Filter on committer name`, ox.Section(0)).
					String(`hash`, `Filter by commit hash`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Int(`limit`, `Maximum number of commits to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Bool(`merge`, `Filter on merge commits`, ox.Section(0)).
					String(`order`, `Order of commits returned, ignored unless '--sort' flag is specified: {asc|desc}`, ox.Default("desc"), ox.Section(0)).
					Slice(`owner`, `Filter on repository owner`, ox.Section(0)).
					String(`parent`, `Filter by parent hash`, ox.Section(0)).
					Slice(`repo`, `Filter on repository`, ox.Short("R"), ox.Section(0)).
					String(`sort`, `Sort fetched commits: {author-date|committer-date}`, ox.Default("best-match"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`tree`, `Filter by tree hash`, ox.Section(0)).
					Slice(`visibility`, `Filter based on repository visibility: {public|private|internal}`, ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`issues`, `Search for issues`),
				ox.Banner(`Search for issues on GitHub.

The command supports constructing queries using the GitHub search syntax,
using the parameter and qualifier flags, or a combination of the two.

GitHub search syntax is documented at:
<https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<query>] [flags]`),
				ox.Example(`
  # Search issues matching set of keywords "readme" and "typo"
  $ gh search issues readme typo
  
  # Search issues matching phrase "broken feature"
  $ gh search issues "broken feature"
  
  # Search issues and pull requests in cli organization
  $ gh search issues --include-prs --owner=cli
  
  # Search open issues assigned to yourself
  $ gh search issues --assignee=@me --state=open
  
  # Search issues with numerous comments
  $ gh search issues --comments=">100"
  
  # Search issues without label "bug"
  $ gh search issues -- -label:bug
  
  # Search issues only from un-archived repositories (default is all repositories)
  $ gh search issues --owner github --archived=false`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Filter by GitHub App author`, ox.Section(0)).
					Bool(`archived`, `Filter based on the repository archived state {true|false}`, ox.Section(0)).
					String(`assignee`, `Filter by assignee`, ox.Section(0)).
					String(`author`, `Filter by author`, ox.Section(0)).
					Date(`closed`, `Filter on closed at date`, ox.Section(0)).
					String(`commenter`, `Filter based on comments by user`, ox.Spec(`user`), ox.Section(0)).
					Int(`comments`, `Filter on number of comments`, ox.Spec(`number`), ox.Section(0)).
					Date(`created`, `Filter based on created at date`, ox.Section(0)).
					Bool(`include-prs`, `Include pull requests in results`, ox.Section(0)).
					Int(`interactions`, `Filter on number of reactions and comments`, ox.Spec(`number`), ox.Section(0)).
					String(`involves`, `Filter based on involvement of user`, ox.Spec(`user`), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Slice(`label`, `Filter on label`, ox.Section(0)).
					String(`language`, `Filter based on the coding language`, ox.Section(0)).
					Int(`limit`, `Maximum number of results to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Bool(`locked`, `Filter on locked conversation status`, ox.Section(0)).
					Slice(`match`, `Restrict search to specific field of issue: {title|body|comments}`, ox.Section(0)).
					String(`mentions`, `Filter based on user mentions`, ox.Spec(`user`), ox.Section(0)).
					String(`milestone`, `Filter by milestone title`, ox.Spec(`title`), ox.Section(0)).
					Bool(`no-assignee`, `Filter on missing assignee`, ox.Section(0)).
					Bool(`no-label`, `Filter on missing label`, ox.Section(0)).
					Bool(`no-milestone`, `Filter on missing milestone`, ox.Section(0)).
					Bool(`no-project`, `Filter on missing project`, ox.Section(0)).
					String(`order`, `Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}`, ox.Default("desc"), ox.Section(0)).
					Slice(`owner`, `Filter on repository owner`, ox.Section(0)).
					String(`project`, `Filter on project board owner/number`, ox.Spec(`owner/number`), ox.Section(0)).
					Int(`reactions`, `Filter on number of reactions`, ox.Spec(`number`), ox.Section(0)).
					Slice(`repo`, `Filter on repository`, ox.Short("R"), ox.Section(0)).
					String(`sort`, `Sort fetched results: {comments|created|interactions|reactions|reactions-+1|reactions--1|reactions-heart|reactions-smile|reactions-tada|reactions-thinking_face|updated}`, ox.Default("best-match"), ox.Section(0)).
					String(`state`, `Filter based on state: {open|closed}`, ox.Section(0)).
					String(`team-mentions`, `Filter based on team mentions`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Date(`updated`, `Filter on last updated at date`, ox.Section(0)).
					Slice(`visibility`, `Filter based on repository visibility: {public|private|internal}`, ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`prs`, `Search for pull requests`),
				ox.Banner(`Search for pull requests on GitHub.

The command supports constructing queries using the GitHub search syntax,
using the parameter and qualifier flags, or a combination of the two.

GitHub search syntax is documented at:
<https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<query>] [flags]`),
				ox.Example(`
  # Search pull requests matching set of keywords "fix" and "bug"
  $ gh search prs fix bug
  
  # Search draft pull requests in cli repository
  $ gh search prs --repo=cli/cli --draft
  
  # Search open pull requests requesting your review
  $ gh search prs --review-requested=@me --state=open
  
  # Search merged pull requests assigned to yourself
  $ gh search prs --assignee=@me --merged
  
  # Search pull requests with numerous reactions
  $ gh search prs --reactions=">100"
  
  # Search pull requests without label "bug"
  $ gh search prs -- -label:bug
  
  # Search pull requests only from un-archived repositories (default is all repositories)
  $ gh search prs --owner github --archived=false`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Filter by GitHub App author`, ox.Section(0)).
					Bool(`archived`, `Filter based on the repository archived state {true|false}`, ox.Section(0)).
					String(`assignee`, `Filter by assignee`, ox.Section(0)).
					String(`author`, `Filter by author`, ox.Section(0)).
					String(`base`, `Filter on base branch name`, ox.Short("B"), ox.Section(0)).
					String(`checks`, `Filter based on status of the checks: {pending|success|failure}`, ox.Section(0)).
					Date(`closed`, `Filter on closed at date`, ox.Section(0)).
					String(`commenter`, `Filter based on comments by user`, ox.Spec(`user`), ox.Section(0)).
					Int(`comments`, `Filter on number of comments`, ox.Spec(`number`), ox.Section(0)).
					Date(`created`, `Filter based on created at date`, ox.Section(0)).
					Bool(`draft`, `Filter based on draft state`, ox.Section(0)).
					String(`head`, `Filter on head branch name`, ox.Short("H"), ox.Section(0)).
					Int(`interactions`, `Filter on number of reactions and comments`, ox.Spec(`number`), ox.Section(0)).
					String(`involves`, `Filter based on involvement of user`, ox.Spec(`user`), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					Slice(`label`, `Filter on label`, ox.Section(0)).
					String(`language`, `Filter based on the coding language`, ox.Section(0)).
					Int(`limit`, `Maximum number of results to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Bool(`locked`, `Filter on locked conversation status`, ox.Section(0)).
					Slice(`match`, `Restrict search to specific field of issue: {title|body|comments}`, ox.Section(0)).
					String(`mentions`, `Filter based on user mentions`, ox.Spec(`user`), ox.Section(0)).
					Bool(`merged`, `Filter based on merged state`, ox.Section(0)).
					Date(`merged-at`, `Filter on merged at date`, ox.Section(0)).
					String(`milestone`, `Filter by milestone title`, ox.Spec(`title`), ox.Section(0)).
					Bool(`no-assignee`, `Filter on missing assignee`, ox.Section(0)).
					Bool(`no-label`, `Filter on missing label`, ox.Section(0)).
					Bool(`no-milestone`, `Filter on missing milestone`, ox.Section(0)).
					Bool(`no-project`, `Filter on missing project`, ox.Section(0)).
					String(`order`, `Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}`, ox.Default("desc"), ox.Section(0)).
					Slice(`owner`, `Filter on repository owner`, ox.Section(0)).
					String(`project`, `Filter on project board owner/number`, ox.Spec(`owner/number`), ox.Section(0)).
					Int(`reactions`, `Filter on number of reactions`, ox.Spec(`number`), ox.Section(0)).
					Slice(`repo`, `Filter on repository`, ox.Short("R"), ox.Section(0)).
					String(`review`, `Filter based on review status: {none|required|approved|changes_requested}`, ox.Section(0)).
					String(`review-requested`, `Filter on user or team requested to review`, ox.Spec(`user`), ox.Section(0)).
					String(`reviewed-by`, `Filter on user who reviewed`, ox.Spec(`user`), ox.Section(0)).
					String(`sort`, `Sort fetched results: {comments|reactions|reactions-+1|reactions--1|reactions-smile|reactions-thinking_face|reactions-heart|reactions-tada|interactions|created|updated}`, ox.Default("best-match"), ox.Section(0)).
					String(`state`, `Filter based on state: {open|closed}`, ox.Section(0)).
					String(`team-mentions`, `Filter based on team mentions`, ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Date(`updated`, `Filter on last updated at date`, ox.Section(0)).
					Slice(`visibility`, `Filter based on repository visibility: {public|private|internal}`, ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`repos`, `Search for repositories`),
				ox.Banner(`Search for repositories on GitHub.

The command supports constructing queries using the GitHub search syntax,
using the parameter and qualifier flags, or a combination of the two.

GitHub search syntax is documented at:
<https://docs.github.com/search-github/searching-on-github/searching-for-repositories>

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[<query>] [flags]`),
				ox.Example(`
  # Search repositories matching set of keywords "cli" and "shell"
  $ gh search repos cli shell
  
  # Search repositories matching phrase "vim plugin"
  $ gh search repos "vim plugin"
  
  # Search repositories public repos in the microsoft organization
  $ gh search repos --owner=microsoft --visibility=public
  
  # Search repositories with a set of topics
  $ gh search repos --topic=unix,terminal
  
  # Search repositories by coding language and number of good first issues
  $ gh search repos --language=go --good-first-issues=">=10"
  
  # Search repositories without topic "linux"
  $ gh search repos -- -topic:linux
  
  # Search repositories excluding archived repositories
  $ gh search repos --archived=false`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`archived`, `Filter based on the repository archived state {true|false}`, ox.Section(0)).
					Date(`created`, `Filter based on created at date`, ox.Section(0)).
					Int(`followers`, `Filter based on number of followers`, ox.Spec(`number`), ox.Section(0)).
					Int(`forks`, `Filter on number of forks`, ox.Spec(`number`), ox.Section(0)).
					Int(`good-first-issues`, `Filter on number of issues with the 'good first issue' label`, ox.Spec(`number`), ox.Section(0)).
					Int(`help-wanted-issues`, `Filter on number of issues with the 'help wanted' label`, ox.Spec(`number`), ox.Section(0)).
					String(`include-forks`, `Include forks in fetched repositories: {false|true|only}`, ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`language`, `Filter based on the coding language`, ox.Section(0)).
					Slice(`license`, `Filter based on license type`, ox.Section(0)).
					Int(`limit`, `Maximum number of repositories to fetch`, ox.Default("30"), ox.Short("L"), ox.Section(0)).
					Slice(`match`, `Restrict search to specific field of repository: {name|description|readme}`, ox.Section(0)).
					Int(`number-topics`, `Filter on number of topics`, ox.Spec(`number`), ox.Section(0)).
					String(`order`, `Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}`, ox.Default("desc"), ox.Section(0)).
					Slice(`owner`, `Filter on owner`, ox.Section(0)).
					String(`size`, `Filter on a size range, in kilobytes`, ox.Section(0)).
					String(`sort`, `Sort fetched repositories: {forks|help-wanted-issues|stars|updated}`, ox.Default("best-match"), ox.Section(0)).
					Int(`stars`, `Filter on number of stars`, ox.Spec(`number`), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Slice(`topic`, `Filter on topic`, ox.Section(0)).
					Date(`updated`, `Filter on last updated at date`, ox.Section(0)).
					Slice(`visibility`, `Filter based on visibility: {public|private|internal}`, ox.Section(0)).
					Bool(`web`, `Open the search query in the web browser`, ox.Short("w"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`secret`, `Manage GitHub secrets`),
			ox.Banner(`Secrets can be set at the repository, or organization level for use in
GitHub Actions or Dependabot. User, organization, and repository secrets can be set for
use in GitHub Codespaces. Environment secrets can be set for use in
GitHub Actions. Run `+"`"+`gh help secret set`+"`"+` to learn how to get started.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`delete`, `Delete secrets`),
				ox.Banner(`Delete a secret on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization
- user: available to Codespaces for your user`),
				ox.Spec(`<secret-name> [flags]`),
				ox.Aliases("secret remove"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Delete a secret for a specific application: {actions|codespaces|dependabot}`, ox.Short("a"), ox.Section(0)).
					String(`env`, `Delete a secret for an environment`, ox.Short("e"), ox.Section(0)).
					String(`org`, `Delete a secret for an organization`, ox.Short("o"), ox.Section(0)).
					Bool(`user`, `Delete a secret for your user`, ox.Short("u"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List secrets`),
				ox.Banner(`List secrets on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization
- user: available to Codespaces for your user

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("secret ls"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `List secrets for a specific application: {actions|codespaces|dependabot}`, ox.Short("a"), ox.Section(0)).
					String(`env`, `List secrets for an environment`, ox.Short("e"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`org`, `List secrets for an organization`, ox.Short("o"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					Bool(`user`, `List a secret for your user`, ox.Short("u"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`set`, `Create or update secrets`),
				ox.Banner(`Set a value for a secret on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization
- user: available to Codespaces for your user

Organization and user secrets can optionally be restricted to only be available to
specific repositories.

Secret values are locally encrypted before being sent to GitHub.`),
				ox.Spec(`<secret-name> [flags]`),
				ox.Example(`
  # Paste secret value for the current repository in an interactive prompt
  $ gh secret set MYSECRET
  
  # Read secret value from an environment variable
  $ gh secret set MYSECRET --body "$ENV_VALUE"
  
  # Set secret for a specific remote repository
  $ gh secret set MYSECRET --repo origin/repo --body "$ENV_VALUE"
  
  # Read secret value from a file
  $ gh secret set MYSECRET < myfile.txt
  
  # Set secret for a deployment environment in the current repository
  $ gh secret set MYSECRET --env myenvironment
  
  # Set organization-level secret visible to both public and private repositories
  $ gh secret set MYSECRET --org myOrg --visibility all
  
  # Set organization-level secret visible to specific repositories
  $ gh secret set MYSECRET --org myOrg --repos repo1,repo2,repo3
  
  # Set user-level secret for Codespaces
  $ gh secret set MYSECRET --user
  
  # Set repository-level secret for Dependabot
  $ gh secret set MYSECRET --app dependabot
  
  # Set multiple secrets imported from the ".env" file
  $ gh secret set -f .env
  
  # Set multiple secrets from stdin
  $ gh secret set -f - < myfile.txt`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`app`, `Set the application for a secret: {actions|codespaces|dependabot}`, ox.Short("a"), ox.Section(0)).
					String(`body`, `The value for the secret (reads from standard input if not specified)`, ox.Short("b"), ox.Section(0)).
					String(`env`, `Set deployment environment secret`, ox.Spec(`environment`), ox.Short("e"), ox.Section(0)).
					String(`env-file`, `Load secret names and values from a dotenv-formatted file`, ox.Spec(`file`), ox.Short("f"), ox.Section(0)).
					Bool(`no-store`, `Print the encrypted, base64-encoded value instead of storing it on GitHub`, ox.Section(0)).
					String(`org`, `Set organization secret`, ox.Spec(`organization`), ox.Short("o"), ox.Section(0)).
					String(`repos`, `List of repositories that can access an organization or user secret`, ox.Spec(`repositories`), ox.Short("r"), ox.Section(0)).
					Bool(`user`, `Set a secret for your user`, ox.Short("u"), ox.Section(0)).
					String(`visibility`, `Set visibility for an organization secret: {all|private|selected}`, ox.Default("private"), ox.Short("v"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
		ox.Sub(
			ox.Usage(`ssh-key`, `Manage SSH keys`),
			ox.Banner(`Manage SSH keys registered with your GitHub account.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Sub(
				ox.Usage(`add`, `Add an SSH key to your GitHub account`),
				ox.Banner(`Add an SSH key to your GitHub account`),
				ox.Spec(`[<key-file>] [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`title`, `Title for the new key`, ox.Short("t"), ox.Section(0)).
					String(`type`, `Type of the ssh key: {authentication|signing}`, ox.Default("authentication"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`delete`, `Delete an SSH key from your GitHub account`),
				ox.Banner(`Delete an SSH key from your GitHub account`),
				ox.Spec(`<id> [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					Bool(`yes`, `Skip the confirmation prompt`, ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `Lists SSH keys in your GitHub account`),
				ox.Banner(`Lists SSH keys in your GitHub account`),
				ox.Spec(`[flags]`),
				ox.Aliases("ssh-key ls"),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			),
		),
		ox.Sub(
			ox.Usage(`status`, `Print information about relevant issues, pull requests, and notifications across repositories`),
			ox.Banner(`The status command prints information about your work on GitHub across all the repositories you're subscribed to, including:

- Assigned Issues
- Assigned Pull Requests
- Review Requests
- Mentions
- Repository Activity (new issues/pull requests, comments)`),
			ox.Spec(`[flags]`),
			ox.Example(`
  $ gh status -e cli/cli -e cli/go-gh # Exclude multiple repositories
  $ gh status -o cli # Limit results to a single organization`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				Slice(`exclude`, `Comma separated list of repos to exclude in owner/name format`, ox.Short("e"), ox.Section(0)).
				String(`org`, `Report status within an organization`, ox.Short("o"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`variable`, `Manage GitHub Actions variables`),
			ox.Banner(`Variables can be set at the repository, environment or organization level for use in
GitHub Actions or Dependabot. Run `+"`"+`gh help variable set`+"`"+` to learn how to get started.`),
			ox.Spec(`<command> [flags]`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"FLAGS",
			)),
			ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
			ox.Flags().
				String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`delete`, `Delete variables`),
				ox.Banner(`Delete a variable on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs or Dependabot within an organization`),
				ox.Spec(`<variable-name> [flags]`),
				ox.Aliases("variable remove"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`env`, `Delete a variable for an environment`, ox.Short("e"), ox.Section(0)).
					String(`org`, `Delete a variable for an organization`, ox.Short("o"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`get`, `Get variables`),
				ox.Banner(`Get a variable on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs or Dependabot within an organization

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`<variable-name> [flags]`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`env`, `Get a variable for an environment`, ox.Short("e"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`org`, `Get a variable for an organization`, ox.Short("o"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`list`, `List variables`),
				ox.Banner(`List variables on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs or Dependabot within an organization

For more information about output formatting flags, see `+"`"+`gh help formatting`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("variable ls"),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`env`, `List variables for an environment`, ox.Short("e"), ox.Section(0)).
					String(`jq`, `Filter JSON output using a jq expression`, ox.Spec(`expression`), ox.Short("q"), ox.Section(0)).
					Slice(`json`, `Output JSON with the specified fields`, ox.Section(0)).
					String(`org`, `List variables for an organization`, ox.Short("o"), ox.Section(0)).
					String(`template`, `Format JSON output using a Go template; see "gh help formatting"`, ox.Short("t"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
			ox.Sub(
				ox.Usage(`set`, `Create or update variables`),
				ox.Banner(`Set a value for a variable on one of the following levels:
- repository (default): available to GitHub Actions runs or Dependabot in a repository
- environment: available to GitHub Actions runs for a deployment environment in a repository
- organization: available to GitHub Actions runs or Dependabot within an organization

Organization variable can optionally be restricted to only be available to
specific repositories.`),
				ox.Spec(`<variable-name> [flags]`),
				ox.Example(`
  # Add variable value for the current repository in an interactive prompt
  $ gh variable set MYVARIABLE
  
  # Read variable value from an environment variable
  $ gh variable set MYVARIABLE --body "$ENV_VALUE"
  
  # Read variable value from a file
  $ gh variable set MYVARIABLE < myfile.txt
  
  # Set variable for a deployment environment in the current repository
  $ gh variable set MYVARIABLE --env myenvironment
  
  # Set organization-level variable visible to both public and private repositories
  $ gh variable set MYVARIABLE --org myOrg --visibility all
  
  # Set organization-level variable visible to specific repositories
  $ gh variable set MYVARIABLE --org myOrg --repos repo1,repo2,repo3
  
  # Set multiple variables imported from the ".env" file
  $ gh variable set -f .env`),
				ox.Help(ox.Sections(
					"FLAGS",
					"INHERITED FLAGS",
				)),
				ox.Footer(`Use `+"`"+`gh <command> <subcommand> --help`+"`"+` for more information about a command.
  Read the manual at https://cli.github.com/manual
  Learn about exit codes using `+"`"+`gh help exit-codes`+"`"+`
  Learn about accessibility experiences using `+"`"+`gh help accessibility`+"`"+``),
				ox.Flags().
					String(`body`, `The value for the variable (reads from standard input if not specified)`, ox.Short("b"), ox.Section(0)).
					String(`env`, `Set deployment environment variable`, ox.Spec(`environment`), ox.Short("e"), ox.Section(0)).
					String(`env-file`, `Load variable names and values from a dotenv-formatted file`, ox.Spec(`file`), ox.Short("f"), ox.Section(0)).
					String(`org`, `Set organization variable`, ox.Spec(`organization`), ox.Short("o"), ox.Section(0)).
					String(`repos`, `List of repositories that can access an organization variable`, ox.Spec(`repositories`), ox.Short("r"), ox.Section(0)).
					String(`visibility`, `Set visibility for an organization variable: {all|private|selected}`, ox.Default("private"), ox.Short("v"), ox.Section(0)).
					String(`repo`, `Select another repository using the [HOST/]OWNER/REPO format`, ox.Spec(`[HOST/]OWNER/REPO`), ox.Short("R"), ox.Section(1)),
			),
		),
	)
}
