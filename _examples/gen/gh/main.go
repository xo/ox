// Command gh is a xo/ox version of `+gh`.
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
		ox.Banner("Work seamlessly with GitHub from the command line."),
		ox.Usage("gh", ""),
		ox.Spec("<command> <subcommand> [flags]"),
		ox.Example("\n  $ gh issue create\n  $ gh repo clone cli/cli\n  $ gh pr checkout 321"),
		ox.Sections("CORE COMMANDS", "GITHUB ACTIONS COMMANDS", "ALIAS COMMANDS", "ADDITIONAL COMMANDS"),
		ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
		ox.Sub(
			ox.Banner("Authenticate gh and git with GitHub"),
			ox.Usage("auth", "Authenticate gh and git with GitHub"),
			ox.Spec("<command> [flags]"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Authenticate with a GitHub host.\n\nThe default hostname is `github.com`. This can be overridden using the `--hostname`\nflag.\n\nThe default authentication mode is a web-based browser flow. After completion, an\nauthentication token will be stored securely in the system credential store.\nIf a credential store is not found or there is an issue using it gh will fallback\nto writing the token to a plain text file. See `gh auth status` for its\nstored location.\n\nAlternatively, use `--with-token` to pass in a token on standard input.\nThe minimum required scopes for the token are: `repo`, `read:org`, and `gist`.\n\nAlternatively, gh will use the authentication token found in environment variables.\nThis method is most suitable for \"headless\" use of gh such as in automation. See\n`gh help environment` for more info.\n\nTo use gh in GitHub Actions, add `GH_TOKEN: ${{ github.token }}` to `env`.\n\nThe git protocol to use for git operations on this host can be set with `--git-protocol`,\nor during the interactive prompting. Although login is for a single account on a host, setting\nthe git protocol will take effect for all users on the host.\n\nSpecifying `ssh` for the git protocol will detect existing SSH keys to upload,\nprompting to create and upload a new key if one is not found. This can be skipped with\n`--skip-ssh-key` flag."),
				ox.Usage("login", "Log in to a GitHub account"),
				ox.Spec("[flags]"),
				ox.Example("\n  # Start interactive setup\n  $ gh auth login\n  \n  # Authenticate against github.com by reading the token from a file\n  $ gh auth login --with-token < mytoken.txt\n  \n  # Authenticate with specific host\n  $ gh auth login --hostname enterprise.internal"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("git-protocol", "The protocol to use for git operations on this host: {ssh|https}", ox.Short("p")).
					String("hostname", "The hostname of the GitHub instance to authenticate with", ox.Short("h")).
					Bool("insecure-storage", "Save authentication credentials in plain text instead of credential store").
					Slice("scopes", "Additional authentication scopes to request", ox.Elem(ox.StringT), ox.Short("s")).
					Bool("skip-ssh-key", "Skip generate/upload SSH key prompt").
					Bool("web", "Open a browser to authenticate", ox.Short("w")).
					Bool("with-token", "Read token from standard input"),
			), ox.Sub(
				ox.Banner("Remove authentication for a GitHub account.\n\nThis command removes the stored authentication configuration\nfor an account. The authentication configuration is only\nremoved locally.\n\nThis command does not invalidate authentication tokens."),
				ox.Usage("logout", "Log out of a GitHub account"),
				ox.Spec("[flags]"),
				ox.Example("\n  # Select what host and account to log out of via a prompt\n  $ gh auth logout\n  \n  # Log out of a specific host and specific account\n  $ gh auth logout --hostname enterprise.internal --user monalisa"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("hostname", "The hostname of the GitHub instance to log out of", ox.Short("h")).
					String("user", "The account to log out of", ox.Short("u")),
			), ox.Sub(
				ox.Banner("Expand or fix the permission scopes for stored credentials for active account.\n\nThe `--scopes` flag accepts a comma separated list of scopes you want\nyour gh credentials to have. If no scopes are provided, the command\nmaintains previously added scopes.\n\nThe `--remove-scopes` flag accepts a comma separated list of scopes you\nwant to remove from your gh credentials. Scope removal is idempotent.\nThe minimum set of scopes (`repo`, `read:org`, and `gist`) cannot be removed.\n\nThe `--reset-scopes` flag resets the scopes for your gh credentials to\nthe default set of scopes for your auth flow.\n\nIf you have multiple accounts in `gh auth status` and want to refresh the credentials for an\ninactive account, you will have to use `gh auth switch` to that account first before using\nthis command, and then switch back when you are done."),
				ox.Usage("refresh", "Refresh stored authentication credentials"),
				ox.Spec("[flags]"),
				ox.Example("\n  $ gh auth refresh --scopes write:org,read:public_key\n  # => open a browser to add write:org and read:public_key scopes\n  \n  $ gh auth refresh\n  # => open a browser to ensure your authentication credentials have the correct minimum scopes\n  \n  $ gh auth refresh --remove-scopes delete_repo\n  # => open a browser to idempotently remove the delete_repo scope\n  \n  $ gh auth refresh --reset-scopes\n  # => open a browser to re-authenticate with the default minimum scopes"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("hostname", "The GitHub host to use for authentication", ox.Short("h")).
					Bool("insecure-storage", "Save authentication credentials in plain text instead of credential store").
					Slice("remove-scopes", "Authentication scopes to remove from gh", ox.Elem(ox.StringT), ox.Short("r")).
					Bool("reset-scopes", "Reset authentication scopes to the default minimum set of scopes").
					Slice("scopes", "Additional authentication scopes for gh to have", ox.Elem(ox.StringT), ox.Short("s")),
			), ox.Sub(
				ox.Banner("This command configures `git` to use GitHub CLI as a credential helper.\nFor more information on git credential helpers please reference:\n<https://git-scm.com/docs/gitcredentials>.\n\nBy default, GitHub CLI will be set as the credential helper for all authenticated hosts.\nIf there is no authenticated hosts the command fails with an error.\n\nAlternatively, use the `--hostname` flag to specify a single host to be configured.\nIf the host is not authenticated with, the command fails with an error."),
				ox.Usage("setup-git", "Setup git with GitHub CLI"),
				ox.Spec("[flags]"),
				ox.Example("\n  # Configure git to use GitHub CLI as the credential helper for all authenticated hosts\n  $ gh auth setup-git\n  \n  # Configure git to use GitHub CLI as the credential helper for enterprise.internal host\n  $ gh auth setup-git --hostname enterprise.internal"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("force", "Force setup even if the host is not known. Must be used in conjunction with --hostname", ox.Spec("--hostname"), ox.Short("f")).
					String("hostname", "The hostname to configure git for", ox.Short("h")),
			), ox.Sub(
				ox.Banner("Display active account and authentication state on each known GitHub host.\n\nFor each host, the authentication state of each known account is tested and any issues are included in the output.\nEach host section will indicate the active account, which will be used when targeting that host.\nIf an account on any host (or only the one given via `--hostname`) has authentication issues,\nthe command will exit with 1 and output to stderr.\n\nTo change the active account for a host, see `gh auth switch`."),
				ox.Usage("status", "Display active account and authentication state on each known GitHub host"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("active", "Display the active account only", ox.Short("a")).
					String("hostname", "Check only a specific hostname's auth status", ox.Short("h")).
					Bool("show-token", "Display the auth token", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Switch the active account for a GitHub host.\n\nThis command changes the authentication configuration that will\nbe used when running commands targeting the specified GitHub host.\n\nIf the specified host has two accounts, the active account will be switched\nautomatically. If there are more than two accounts, disambiguation will be\nrequired either through the `--user` flag or an interactive prompt.\n\nFor a list of authenticated accounts you can run `gh auth status`."),
				ox.Usage("switch", "Switch active GitHub account"),
				ox.Spec("[flags]"),
				ox.Example("\n  # Select what host and account to switch to via a prompt\n  $ gh auth switch\n  \n  # Switch the active account on a specific host to a specific user\n  $ gh auth switch --hostname enterprise.internal --user monalisa"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("hostname", "The hostname of the GitHub instance to switch account for", ox.Short("h")).
					String("user", "The account to switch to", ox.Short("u")),
			), ox.Sub(
				ox.Banner("This command outputs the authentication token for an account on a given GitHub host.\n\nWithout the `--hostname` flag, the default host is chosen.\n\nWithout the `--user` flag, the active account for the host is chosen."),
				ox.Usage("token", "Print the authentication token gh uses for a hostname and account"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("hostname", "The hostname of the GitHub instance authenticated with", ox.Short("h")).
					String("user", "The account to output the token for", ox.Short("u")),
			)), ox.Sub(
			ox.Banner("Open the GitHub repository in the web browser."),
			ox.Usage("browse", "Open the repository in the browser"),
			ox.Spec("[<number> | <path> | <commit-SHA>] [flags]"),
			ox.Example("\n  $ gh browse\n  #=> Open the home page of the current repository\n  \n  $ gh browse script/\n  #=> Open the script directory of the current repository\n  \n  $ gh browse 217\n  #=> Open issue or pull request 217\n  \n  $ gh browse 77507cd94ccafcf568f8560cfecde965fcfa63\n  #=> Open commit page\n  \n  $ gh browse --settings\n  #=> Open repository settings\n  \n  $ gh browse main.go:312\n  #=> Open main.go at line 312\n  \n  $ gh browse main.go --branch bug-fix\n  #=> Open main.go with the repository at head of bug-fix branch\n  \n  $ gh browse main.go --commit=77507cd94ccafcf568f8560cfecde965fcfa63\n  #=> Open main.go with the repository at commit 775007cd"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Flags().
				String("branch", "Select another branch by passing in the branch name", ox.Short("b")).
				Map("commit", "Select another commit by passing in the commit SHA, default is the last commit", ox.Spec("string[=\"last\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Short("c")).
				Bool("no-browser", "Print destination URL instead of opening the browser", ox.Short("n")).
				Bool("projects", "Open repository projects", ox.Short("p")).
				Bool("releases", "Open repository releases", ox.Short("r")).
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")).
				Bool("settings", "Open repository settings", ox.Short("s")).
				Bool("wiki", "Open repository wiki", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Connect to and manage codespaces"),
			ox.Usage("codespace", "Connect to and manage codespaces"),
			ox.Spec("[flags]"),
			ox.Aliases("gh cs"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Open a codespace in Visual Studio Code"),
				ox.Usage("code", "Open a codespace in Visual Studio Code"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("insiders", "Use the insiders version of Visual Studio Code").
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					Bool("web", "Use the web version of Visual Studio Code", ox.Short("w")),
			), ox.Sub(
				ox.Banner("The `cp` command copies files between the local and remote file systems.\n\nAs with the UNIX `cp` command, the first argument specifies the source and the last\nspecifies the destination; additional sources may be specified after the first,\nif the destination is a directory.\n\nThe `--recursive` flag is required if any source is a directory.\n\nA `remote:` prefix on any file name argument indicates that it refers to\nthe file system of the remote (Codespace) machine. It is resolved relative\nto the home directory of the remote user.\n\nBy default, remote file names are interpreted literally. With the `--expand` flag,\neach such argument is treated in the manner of `scp`, as a Bash expression to\nbe evaluated on the remote machine, subject to expansion of tildes, braces, globs,\nenvironment variables, and backticks. For security, do not use this flag with arguments\nprovided by untrusted users; see <https://lwn.net/Articles/835962/> for discussion.\n\nBy default, the `cp` command will create a public/private ssh key pair to authenticate with \nthe codespace inside the `~/.ssh directory`."),
				ox.Usage("cp", "Copy files between local and remote file systems"),
				ox.Spec("[-e] [-r] [-- [<scp flags>...]] <sources>... <dest>"),
				ox.Example("\n  $ gh codespace cp -e README.md 'remote:/workspaces/$RepositoryName/'\n  $ gh codespace cp -e 'remote:~/*.go' ./gofiles/\n  $ gh codespace cp -e 'remote:/workspaces/myproj/go.{mod,sum}' ./gofiles/\n  $ gh codespace cp -e -- -F ~/.ssh/codespaces_config 'remote:~/*.go' ./gofiles/"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("expand", "Expand remote file names on remote shell", ox.Short("e")).
					String("profile", "Name of the SSH profile to use", ox.Short("p")).
					Bool("recursive", "Recursively copy directories", ox.Short("r")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
			), ox.Sub(
				ox.Banner("Create a codespace"),
				ox.Usage("create", "Create a codespace"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("branch", "repository branch", ox.Short("b")).
					Bool("default-permissions", "do not prompt to accept additional permissions requested by the codespace").
					String("devcontainer-path", "path to the devcontainer.json file to use when creating codespace").
					String("display-name", "display name for the codespace (48 characters or less)", ox.Short("d")).
					Duration("idle-timeout", "allowed inactivity before codespace is stopped, e.g. \"10m\", \"1h\"").
					String("location", "location: {EastUs|SouthEastAsia|WestEurope|WestUs2} (determined automatically if not provided)", ox.Short("l")).
					String("machine", "hardware specifications for the VM", ox.Short("m")).
					String("repo", "repository name with owner: user/repo", ox.Short("R")).
					Duration("retention-period", "allowed time after shutting down before the codespace is automatically deleted (maximum 30 days), e.g. \"1h\", \"72h\"").
					Bool("status", "show status of post-create command and dotfiles", ox.Short("s")).
					Bool("web", "create codespace from browser, cannot be used with --display-name, --idle-timeout, or --retention-period", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Delete codespaces based on selection criteria.\n\nAll codespaces for the authenticated user can be deleted, as well as codespaces for a\nspecific repository. Alternatively, only codespaces older than N days can be deleted.\n\nOrganization administrators may delete any codespace billed to the organization."),
				ox.Usage("delete", "Delete codespaces"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Delete all codespaces").
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("days", "Delete codespaces older than N days", ox.Default("N")).
					Bool("force", "Skip confirmation for codespaces that contain unsaved changes", ox.Short("f")).
					String("org", "The login handle of the organization (admin-only)", ox.Spec("login"), ox.Short("o")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					String("user", "The username to delete codespaces for (used with --org)", ox.Spec("username"), ox.Short("u")),
			), ox.Sub(
				ox.Banner("Edit a codespace"),
				ox.Usage("edit", "Edit a codespace"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					String("display-name", "Set the display name", ox.Short("d")).
					String("machine", "Set hardware specifications for the VM", ox.Short("m")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
			), ox.Sub(
				ox.Banner("Open a codespace in JupyterLab"),
				ox.Usage("jupyter", "Open a codespace in JupyterLab"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
			), ox.Sub(
				ox.Banner("List codespaces of the authenticated user.\n\nAlternatively, organization administrators may list all codespaces billed to the organization.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List codespaces"),
				ox.Spec("[flags]"),
				ox.Aliases("gh codespace ls", "gh cs ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of codespaces to list", ox.Default("30"), ox.Short("L")).
					String("org", "The login handle of the organization to list codespaces for (admin-only)", ox.Spec("login"), ox.Short("o")).
					String("repo", "Repository name with owner: user/repo", ox.Short("R")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("user", "The username to list codespaces for (used with --org)", ox.Spec("username"), ox.Short("u")).
					Bool("web", "List codespaces in the web browser, cannot be used with --user or --org", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Access codespace logs"),
				ox.Usage("logs", "Access codespace logs"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("follow", "Tail and follow the logs", ox.Short("f")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
			), ox.Sub(
				ox.Banner("List ports in a codespace\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("ports", "List ports in a codespace"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Sub(
					ox.Banner("Forward ports"),
					ox.Usage("forward", "Forward ports"),
					ox.Spec("<remote-port>:<local-port>... [flags]"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						String("codespace", "Name of the codespace", ox.Short("c")).
						String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
						String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
				), ox.Sub(
					ox.Banner("Change the visibility of the forwarded port"),
					ox.Usage("visibility", "Change the visibility of the forwarded port"),
					ox.Spec("<port>:{public|private|org}... [flags]"),
					ox.Example("\n  gh codespace ports visibility 80:org 3000:private 8000:public"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						String("codespace", "Name of the codespace", ox.Short("c")).
						String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
						String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
				), ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Rebuilding recreates your codespace.\n\nYour code and any current changes will be preserved. Your codespace will be rebuilt using\nyour working directory's dev container. A full rebuild also removes cached Docker images."),
				ox.Usage("rebuild", "Rebuild a codespace"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("full", "perform a full rebuild").
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)"),
			), ox.Sub(
				ox.Banner("The `ssh` command is used to SSH into a codespace. In its simplest form, you can\nrun `gh cs ssh`, select a codespace interactively, and connect.\n\nThe `ssh` command will automatically create a public/private ssh key pair in the\n`~/.ssh` directory if you do not have an existing valid key pair. When selecting the\nkey pair to use, the preferred order is:\n\n1. Key specified by `-i` in `<ssh-flags>`\n2. Automatic key, if it already exists\n3. First valid key pair in ssh config (according to `ssh -G`)\n4. Automatic key, newly created\n\nThe `ssh` command also supports deeper integration with OpenSSH using a `--config`\noption that generates per-codespace ssh configuration in OpenSSH format.\nIncluding this configuration in your `~/.ssh/config` improves the user experience\nof tools that integrate with OpenSSH, such as Bash/Zsh completion of ssh hostnames,\nremote path completion for `scp/rsync/sshfs`, `git` ssh remotes, and so on.\n\nOnce that is set up (see the second example below), you can ssh to codespaces as\nif they were ordinary remote hosts (using `ssh`, not `gh cs ssh`).\n\nNote that the codespace you are connecting to must have an SSH server pre-installed.\nIf the docker image being used for the codespace does not have an SSH server,\ninstall it in your `Dockerfile` or, for codespaces that use Debian-based images,\nyou can add the following to your `devcontainer.json`:\n\n\t\"features\": {\n\t\t\"ghcr.io/devcontainers/features/sshd:1\": {\n\t\t\t\"version\": \"latest\"\n\t\t}\n\t}"),
				ox.Usage("ssh", "SSH into a codespace"),
				ox.Spec("[<flags>...] [-- <ssh-flags>...] [<command>]"),
				ox.Example("\n  $ gh codespace ssh\n  \n  $ gh codespace ssh --config > ~/.ssh/codespaces\n  $ printf 'Match all\\nInclude ~/.ssh/codespaces\\n' >> ~/.ssh/config"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					Bool("config", "Write OpenSSH configuration to stdout").
					Bool("debug", "Log debug data to a file", ox.Short("d")).
					String("debug-file", "Path of the file log to").
					String("profile", "Name of the SSH profile to use").
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					Int("server-port", "SSH server port number (0 => pick unused)"),
			), ox.Sub(
				ox.Banner("Stop a running codespace"),
				ox.Usage("stop", "Stop a running codespace"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					String("org", "The login handle of the organization (admin-only)", ox.Spec("login"), ox.Short("o")).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					String("user", "The username to stop codespace for (used with --org)", ox.Spec("username"), ox.Short("u")),
			), ox.Sub(
				ox.Banner("View details about a codespace\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View details about a codespace"),
				ox.Spec("[flags]"),
				ox.Example("\n  # select a codespace from a list of all codespaces you own\n  $ gh cs view\t\n  \n  # view the details of a specific codespace\n  $ gh cs view -c codespace-name-12345\n  \n  # view the list of all available fields for a codespace\n  $ gh cs view --json\n  \n  # view specific fields for a codespace\n  $ gh cs view --json displayName,machineDisplayName,state"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("codespace", "Name of the codespace", ox.Short("c")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("repo", "Filter codespace selection by repository name (user/repo)", ox.Short("R")).
					String("repo-owner", "Filter codespace selection by repository owner (username or org)").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			)), ox.Sub(
			ox.Banner("Work with GitHub gists."),
			ox.Usage("gist", "Manage gists"),
			ox.Spec("<command> [flags]"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Clone a GitHub gist locally.\n\nA gist can be supplied as argument in either of the following formats:\n- by ID, e.g. `5b0e0062eb8e9654adad7bb1d81cc75f`\n- by URL, e.g. `https://gist.github.com/OWNER/5b0e0062eb8e9654adad7bb1d81cc75f`\n\nPass additional `git clone` flags by listing them after `--`."),
				ox.Usage("clone", "Clone a gist locally"),
				ox.Spec("<gist> [<directory>] [-- <gitflags>...]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Create a new GitHub gist with given contents.\n\nGists can be created from one or multiple files. Alternatively, pass `-` as\nfile name to read from standard input.\n\nBy default, gists are secret; use `--public` to make publicly listed ones."),
				ox.Usage("create", "Create a new gist"),
				ox.Spec("[<filename>... | -] [flags]"),
				ox.Aliases("gh gist new"),
				ox.Example("\n  # publish file 'hello.py' as a public gist\n  $ gh gist create --public hello.py\n  \n  # create a gist with a description\n  $ gh gist create hello.py -d \"my Hello-World program in Python\"\n  \n  # create a gist containing several files\n  $ gh gist create hello.py world.py cool.txt\n  \n  # read from standard input to create a gist\n  $ gh gist create -\n  \n  # create a gist from output piped from another command\n  $ cat cool.txt | gh gist create"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("desc", "A description for this gist", ox.Short("d")).
					String("filename", "Provide a filename to be used when reading from standard input", ox.Short("f")).
					Bool("public", "List the gist publicly", ox.Default("secret"), ox.Short("p")).
					Bool("web", "Open the web browser with created gist", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Delete a gist"),
				ox.Usage("delete", "Delete a gist"),
				ox.Spec("{<id> | <url>} [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Edit one of your gists"),
				ox.Usage("edit", "Edit one of your gists"),
				ox.Spec("{<id> | <url>} [<filename>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("add", "Add a new file to the gist", ox.Short("a")).
					String("desc", "New description for the gist", ox.Short("d")).
					String("filename", "Select a file to edit", ox.Short("f")).
					String("remove", "Remove a file from the gist", ox.Short("r")),
			), ox.Sub(
				ox.Banner("List gists from your user account.\n\nYou can use a regular expression to filter the description, file names,\nor even the content of files in the gist using `--filter`.\n\nFor supported regular expression syntax, see <https://pkg.go.dev/regexp/syntax>.\n\nUse `--include-content` to include content of files, noting that\nthis will be slower and increase the rate limit used. Instead of printing a table,\ncode will be printed with highlights similar to `gh search code`:\n\n\t{{gist ID}} {{file name}}\n\t    {{description}}\n\t        {{matching lines from content}}\n\nNo highlights or other color is printed when output is redirected."),
				ox.Usage("list", "List your gists"),
				ox.Spec("[flags]"),
				ox.Aliases("gh gist ls"),
				ox.Example("\n  # list all secret gists from your user account\n  $ gh gist list --secret\n  \n  # find all gists from your user account mentioning \"octo\" anywhere\n  $ gh gist list --filter octo --include-content"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("filter", "Filter gists using a regular expression", ox.Spec("expression")).
					Bool("include-content", "Include gists' file content when filtering").
					Int("limit", "Maximum number of gists to fetch", ox.Default("10"), ox.Short("L")).
					Bool("public", "Show only public gists").
					Bool("secret", "Show only secret gists"),
			), ox.Sub(
				ox.Banner("Rename a file in the given gist ID / URL."),
				ox.Usage("rename", "Rename a file in a gist"),
				ox.Spec("{<id> | <url>} <oldFilename> <newFilename> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("View the given gist or select from recent gists."),
				ox.Usage("view", "View a gist"),
				ox.Spec("[<id> | <url>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("filename", "Display a single file from the gist", ox.Short("f")).
					Bool("files", "List file names from the gist").
					Bool("raw", "Print raw instead of rendered gist contents", ox.Short("r")).
					Bool("web", "Open gist in the browser", ox.Short("w")),
			)), ox.Sub(
			ox.Banner("Work with GitHub issues."),
			ox.Usage("issue", "Manage issues"),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh issue list\n  $ gh issue create --label bug\n  $ gh issue view 123 --web"),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Create an issue on GitHub.\n\nAdding an issue to projects requires authorization with the `project` scope.\nTo authorize, run `gh auth refresh -s project`."),
				ox.Usage("create", "Create a new issue"),
				ox.Spec("[flags]"),
				ox.Aliases("gh issue new"),
				ox.Example("\n  $ gh issue create --title \"I found a bug\" --body \"Nothing works\"\n  $ gh issue create --label \"bug,help wanted\"\n  $ gh issue create --label bug --label \"help wanted\"\n  $ gh issue create --assignee monalisa,hubot\n  $ gh issue create --assignee \"@me\"\n  $ gh issue create --project \"Roadmap\"\n  $ gh issue create --template \"bug_report.md\""),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("assignee", "Assign people by their login. Use \"@me\" to self-assign.", ox.Spec("login"), ox.Short("a")).
					String("body", "Supply a body. Will prompt for one otherwise.", ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("editor", "Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.", ox.Short("e")).
					String("label", "Add labels by name", ox.Spec("name"), ox.Short("l")).
					String("milestone", "Add the issue to a milestone by name", ox.Spec("name"), ox.Short("m")).
					String("project", "Add the issue to projects by title", ox.Spec("title"), ox.Short("p")).
					String("recover", "Recover input from a failed run of create").
					String("template", "Template file to use as starting body text", ox.Spec("file"), ox.Short("T")).
					String("title", "Supply a title. Will prompt for one otherwise.", ox.Short("t")).
					Bool("web", "Open the browser to create an issue", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List issues in a GitHub repository.\n\nThe search query syntax is documented here:\n<https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List issues in a repository"),
				ox.Spec("[flags]"),
				ox.Aliases("gh issue ls"),
				ox.Example("\n  $ gh issue list --label \"bug\" --label \"help wanted\"\n  $ gh issue list --author monalisa\n  $ gh issue list --assignee \"@me\"\n  $ gh issue list --milestone \"The big 1.0\"\n  $ gh issue list --search \"error no:assignee sort:created-asc\""),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Filter by GitHub App author").
					String("assignee", "Filter by assignee", ox.Short("a")).
					String("author", "Filter by author", ox.Short("A")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Slice("label", "Filter by label", ox.Elem(ox.StringT), ox.Short("l")).
					Int("limit", "Maximum number of issues to fetch", ox.Default("30"), ox.Short("L")).
					String("mention", "Filter by mention").
					String("milestone", "Filter by milestone number or title", ox.Short("m")).
					String("search", "Search issues with query", ox.Spec("query"), ox.Short("S")).
					String("state", "Filter by state: {open|closed|all}", ox.Default("open"), ox.Short("s")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "List issues in the web browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Show status of relevant issues\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("status", "Show status of relevant issues"),
				ox.Spec("[flags]"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Close issue"),
				ox.Usage("close", "Close issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("comment", "Leave a closing comment", ox.Short("c")).
					String("reason", "Reason for closing: {completed|not planned}", ox.Short("r")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Add a comment to a GitHub issue.\n\nWithout the body text supplied through flags, the command will interactively\nprompt for the comment text."),
				ox.Usage("comment", "Add a comment to an issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Example("\n  $ gh issue comment 12 --body \"Hi from GitHub CLI\""),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("body", "The comment body text", ox.Default("text"), ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("edit-last", "Edit the last comment of the same author").
					Bool("editor", "Skip prompts and open the text editor to write the body in", ox.Short("e")).
					Bool("web", "Open the web browser to write the comment", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Delete issue"),
				ox.Usage("delete", "Delete issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "confirm deletion without prompting").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Manage linked branches for an issue"),
				ox.Usage("develop", "Manage linked branches for an issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Example("\n  # List branches for issue 123\n  $ gh issue develop --list 123\n  \n  # List branches for issue 123 in repo cli/cli\n  $ gh issue develop --list --repo cli/cli 123\n  \n  # Create a branch for issue 123 based on the my-feature branch\n  $ gh issue develop 123 --base my-feature\n  \n  # Create a branch for issue 123 and checkout it out\n  $ gh issue develop 123 --checkout\n  \n  # Create a branch in repo monalisa/cli for issue 123 in repo cli/cli\n  $ gh issue develop 123 --repo cli/cli --branch-repo monalisa/cli"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("base", "Name of the remote branch you want to make your new branch from", ox.Short("b")).
					String("branch-repo", "Name or URL of the repository where you want to create your new branch").
					Bool("checkout", "Checkout the branch after creating it", ox.Short("c")).
					Bool("list", "List linked branches for the issue", ox.Short("l")).
					String("name", "Name of the branch to create", ox.Short("n")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Edit one or more issues within the same repository.\n\nEditing issues' projects requires authorization with the `project` scope.\nTo authorize, run `gh auth refresh -s project`."),
				ox.Usage("edit", "Edit issues"),
				ox.Spec("{<numbers> | <urls>} [flags]"),
				ox.Example("\n  $ gh issue edit 23 --title \"I found a bug\" --body \"Nothing works\"\n  $ gh issue edit 23 --add-label \"bug,help wanted\" --remove-label \"core\"\n  $ gh issue edit 23 --add-assignee \"@me\" --remove-assignee monalisa,hubot\n  $ gh issue edit 23 --add-project \"Roadmap\" --remove-project v1,v2\n  $ gh issue edit 23 --milestone \"Version 1\"\n  $ gh issue edit 23 --remove-milestone\n  $ gh issue edit 23 --body-file body.txt\n  $ gh issue edit 23 34 --add-label \"help wanted\""),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("add-assignee", "Add assigned users by their login. Use \"@me\" to assign yourself.", ox.Spec("login")).
					String("add-label", "Add labels by name", ox.Spec("name")).
					String("add-project", "Add the issue to projects by title", ox.Spec("title")).
					String("body", "Set the new body.", ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					String("milestone", "Edit the milestone the issue belongs to by name", ox.Spec("name"), ox.Short("m")).
					String("remove-assignee", "Remove assigned users by their login. Use \"@me\" to unassign yourself.", ox.Spec("login")).
					String("remove-label", "Remove labels by name", ox.Spec("name")).
					Bool("remove-milestone", "Remove the milestone association from the issue").
					String("remove-project", "Remove the issue from projects by title", ox.Spec("title")).
					String("title", "Set the new title.", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Lock issue conversation"),
				ox.Usage("lock", "Lock issue conversation"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("reason", "Optional reason for locking conversation (off_topic, resolved, spam, too_heated).", ox.Short("r")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Pin an issue to a repository.\n\nThe issue can be specified by issue number or URL."),
				ox.Usage("pin", "Pin a issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Example("\n  # Pin an issue to the current repository\n  $ gh issue pin 23\n  \n  # Pin an issue by URL\n  $ gh issue pin https://github.com/owner/repo/issues/23\n  \n  # Pin an issue to specific repository\n  $ gh issue pin 23 --repo owner/repo"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Reopen issue"),
				ox.Usage("reopen", "Reopen issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("comment", "Add a reopening comment", ox.Short("c")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Transfer issue to another repository"),
				ox.Usage("transfer", "Transfer issue to another repository"),
				ox.Spec("{<number> | <url>} <destination-repo> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Unlock issue conversation"),
				ox.Usage("unlock", "Unlock issue conversation"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Unpin an issue from a repository.\n\nThe issue can be specified by issue number or URL."),
				ox.Usage("unpin", "Unpin a issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Example("\n  # Unpin issue from the current repository\n  $ gh issue unpin 23\n  \n  # Unpin issue by URL\n  $ gh issue unpin https://github.com/owner/repo/issues/23\n  \n  # Unpin an issue from specific repository\n  $ gh issue unpin 23 --repo owner/repo"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Display the title, body, and other information about an issue.\n\nWith `--web` flag, open the issue in a web browser instead.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View an issue"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("comments", "View issue comments", ox.Short("c")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open an issue in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Work with GitHub organizations."),
			ox.Usage("org", "Manage organizations"),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh org list"),
			ox.Sections("GENERAL COMMANDS"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("List organizations for the authenticated user."),
				ox.Usage("list", "List organizations for the authenticated user."),
				ox.Spec("[flags]"),
				ox.Aliases("gh org ls"),
				ox.Example("\n  # List the first 30 organizations\n  $ gh org list\n  \n  # List more organizations\n  $ gh org list --limit 100"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Int("limit", "Maximum number of organizations to list", ox.Default("30"), ox.Short("L")),
			)), ox.Sub(
			ox.Banner("Work with GitHub pull requests."),
			ox.Usage("pr", "Manage pull requests"),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh pr checkout 353\n  $ gh pr create --fill\n  $ gh pr view --web"),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Create a pull request on GitHub.\n\nWhen the current branch isn't fully pushed to a git remote, a prompt will ask where\nto push the branch and offer an option to fork the base repository. Use `--head` to\nexplicitly skip any forking or pushing behavior.\n\nA prompt will also ask for the title and the body of the pull request. Use `--title` and\n`--body` to skip this, or use `--fill` to autofill these values from git commits.\nIt's important to notice that if the `--title` and/or `--body` are also provided\nalongside `--fill`, the values specified by `--title` and/or `--body` will\ntake precedence and overwrite any autofilled content.\n\nLink an issue to the pull request by referencing the issue in the body of the pull\nrequest. If the body text mentions `Fixes #123` or `Closes #123`, the referenced issue\nwill automatically get closed when the pull request gets merged.\n\nBy default, users with write access to the base repository can push new commits to the\nhead branch of the pull request. Disable this with `--no-maintainer-edit`.\n\nAdding a pull request to projects requires authorization with the `project` scope.\nTo authorize, run `gh auth refresh -s project`."),
				ox.Usage("create", "Create a pull request"),
				ox.Spec("[flags]"),
				ox.Aliases("gh pr new"),
				ox.Example("\n  $ gh pr create --title \"The bug is fixed\" --body \"Everything works again\"\n  $ gh pr create --reviewer monalisa,hubot  --reviewer myorg/team-name\n  $ gh pr create --project \"Roadmap\"\n  $ gh pr create --base develop --head monalisa:feature\n  $ gh pr create --template \"pull_request_template.md\""),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("assignee", "Assign people by their login. Use \"@me\" to self-assign.", ox.Spec("login"), ox.Short("a")).
					String("base", "The branch into which you want your code merged", ox.Spec("branch"), ox.Short("B")).
					String("body", "Body for the pull request", ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("draft", "Mark pull request as a draft", ox.Short("d")).
					Bool("dry-run", "Print details instead of creating the PR. May still push git changes.").
					Bool("editor", "Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.", ox.Short("e")).
					Bool("fill", "Use commit info for title and body", ox.Short("f")).
					Bool("fill-first", "Use first commit info for title and body").
					Bool("fill-verbose", "Use commits msg+body for description").
					String("head", "The branch that contains commits for your pull request", ox.Spec("branch"), ox.Default("[current branch]"), ox.Short("H")).
					String("label", "Add labels by name", ox.Spec("name"), ox.Short("l")).
					String("milestone", "Add the pull request to a milestone by name", ox.Spec("name"), ox.Short("m")).
					Bool("no-maintainer-edit", "Disable maintainer's ability to modify pull request").
					String("project", "Add the pull request to projects by title", ox.Spec("title"), ox.Short("p")).
					String("recover", "Recover input from a failed run of create").
					String("reviewer", "Request reviews from people or teams by their handle", ox.Spec("handle"), ox.Short("r")).
					String("template", "Template file to use as starting body text", ox.Spec("file"), ox.Short("T")).
					String("title", "Title for the pull request", ox.Short("t")).
					Bool("web", "Open the web browser to create a pull request", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List pull requests in a GitHub repository.\n\nThe search query syntax is documented here:\n<https://docs.github.com/en/search-github/searching-on-github/searching-issues-and-pull-requests>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List pull requests in a repository"),
				ox.Spec("[flags]"),
				ox.Aliases("gh pr ls"),
				ox.Example("\n  List PRs authored by you\n  $ gh pr list --author \"@me\"\n  \n  List only PRs with all of the given labels\n  $ gh pr list --label bug --label \"priority 1\"\n  \n  Filter PRs using search syntax\n  $ gh pr list --search \"status:success review:required\"\n  \n  Find a PR that introduced a given commit\n  $ gh pr list --search \"<SHA>\" --state merged"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Filter by GitHub App author").
					String("assignee", "Filter by assignee", ox.Short("a")).
					String("author", "Filter by author", ox.Short("A")).
					String("base", "Filter by base branch", ox.Short("B")).
					Bool("draft", "Filter by draft state", ox.Short("d")).
					String("head", "Filter by head branch", ox.Short("H")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Slice("label", "Filter by label", ox.Elem(ox.StringT), ox.Short("l")).
					Int("limit", "Maximum number of items to fetch", ox.Default("30"), ox.Short("L")).
					String("search", "Search pull requests with query", ox.Spec("query"), ox.Short("S")).
					String("state", "Filter by state: {open|closed|merged|all}", ox.Default("open"), ox.Short("s")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "List pull requests in the web browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Show status of relevant pull requests\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("status", "Show status of relevant pull requests"),
				ox.Spec("[flags]"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("conflict-status", "Display the merge conflict status of each pull request", ox.Short("c")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Check out a pull request in git"),
				ox.Usage("checkout", "Check out a pull request in git"),
				ox.Spec("{<number> | <url> | <branch>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("branch", "Local branch name to use", ox.Default("[the name of the head branch]"), ox.Short("b")).
					Bool("detach", "Checkout PR with a detached HEAD").
					Bool("force", "Reset the existing local branch to the latest state of the pull request", ox.Short("f")).
					Bool("recurse-submodules", "Update all submodules after checkout").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Show CI status for a single pull request.\n\nWithout an argument, the pull request that belongs to the current branch\nis selected.\n\nWhen the `--json` flag is used, it includes a `bucket` field, which categorizes\nthe `state` field into `pass`, `fail`, `pending`, `skipping`, or `cancel`.\n\nAdditional exit codes:\n\t8: Checks pending\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("checks", "Show CI status for a single pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("fail-fast", "Exit watch mode on first check failure").
					String("interval", "Refresh interval in seconds when using --watch flag", ox.Spec("--watch"), ox.Default("10"), ox.Short("i")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Bool("required", "Only show checks that are required").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("watch", "Watch checks until they finish").
					Bool("web", "Open the web browser to show details about checks", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Close a pull request"),
				ox.Usage("close", "Close a pull request"),
				ox.Spec("{<number> | <url> | <branch>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("comment", "Leave a closing comment", ox.Short("c")).
					Bool("delete-branch", "Delete the local and remote branch after close", ox.Short("d")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Add a comment to a GitHub pull request.\n\nWithout the body text supplied through flags, the command will interactively\nprompt for the comment text."),
				ox.Usage("comment", "Add a comment to a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Example("\n  $ gh pr comment 13 --body \"Hi from GitHub CLI\""),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("body", "The comment body text", ox.Default("text"), ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("edit-last", "Edit the last comment of the same author").
					Bool("editor", "Skip prompts and open the text editor to write the body in", ox.Short("e")).
					Bool("web", "Open the web browser to write the comment", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("View changes in a pull request.\n\nWithout an argument, the pull request that belongs to the current branch\nis selected.\n\nWith `--web` flag, open the pull request diff in a web browser instead."),
				ox.Usage("diff", "View changes in a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("color", "Use color in diff output: {always|never|auto}", ox.Default("auto")).
					Bool("name-only", "Display only names of changed files").
					Bool("patch", "Display diff in patch format").
					Bool("web", "Open the pull request diff in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Edit a pull request.\n\nWithout an argument, the pull request that belongs to the current branch\nis selected.\n\nEditing a pull request's projects requires authorization with the `project` scope.\nTo authorize, run `gh auth refresh -s project`."),
				ox.Usage("edit", "Edit a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Example("\n  $ gh pr edit 23 --title \"I found a bug\" --body \"Nothing works\"\n  $ gh pr edit 23 --add-label \"bug,help wanted\" --remove-label \"core\"\n  $ gh pr edit 23 --add-reviewer monalisa,hubot  --remove-reviewer myorg/team-name\n  $ gh pr edit 23 --add-assignee \"@me\" --remove-assignee monalisa,hubot\n  $ gh pr edit 23 --add-project \"Roadmap\" --remove-project v1,v2\n  $ gh pr edit 23 --milestone \"Version 1\"\n  $ gh pr edit 23 --remove-milestone"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("add-assignee", "Add assigned users by their login. Use \"@me\" to assign yourself.", ox.Spec("login")).
					String("add-label", "Add labels by name", ox.Spec("name")).
					String("add-project", "Add the pull request to projects by title", ox.Spec("title")).
					String("add-reviewer", "Add reviewers by their login.", ox.Spec("login")).
					String("base", "Change the base branch for this pull request", ox.Spec("branch"), ox.Short("B")).
					String("body", "Set the new body.", ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					String("milestone", "Edit the milestone the pull request belongs to by name", ox.Spec("name"), ox.Short("m")).
					String("remove-assignee", "Remove assigned users by their login. Use \"@me\" to unassign yourself.", ox.Spec("login")).
					String("remove-label", "Remove labels by name", ox.Spec("name")).
					Bool("remove-milestone", "Remove the milestone association from the pull request").
					String("remove-project", "Remove the pull request from projects by title", ox.Spec("title")).
					String("remove-reviewer", "Remove reviewers by their login.", ox.Spec("login")).
					String("title", "Set the new title.", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Lock pull request conversation"),
				ox.Usage("lock", "Lock pull request conversation"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("reason", "Optional reason for locking conversation (off_topic, resolved, spam, too_heated).", ox.Short("r")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Merge a pull request on GitHub.\n\nWithout an argument, the pull request that belongs to the current branch\nis selected.\n\nWhen targeting a branch that requires a merge queue, no merge strategy is required.\nIf required checks have not yet passed, auto-merge will be enabled.\nIf required checks have passed, the pull request will be added to the merge queue.\nTo bypass a merge queue and merge directly, pass the `--admin` flag."),
				ox.Usage("merge", "Merge a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("admin", "Use administrator privileges to merge a pull request that does not meet requirements").
					String("author-email", "Email text for merge commit author", ox.Default("text"), ox.Short("A")).
					Bool("auto", "Automatically merge only after necessary requirements are met").
					String("body", "Body text for the merge commit", ox.Default("text"), ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("delete-branch", "Delete the local and remote branch after merge", ox.Short("d")).
					Bool("disable-auto", "Disable auto-merge for this pull request").
					String("match-head-commit", "Commit SHA that the pull request head must match to allow merge", ox.Spec("SHA")).
					Bool("merge", "Merge the commits with the base branch", ox.Short("m")).
					Bool("rebase", "Rebase the commits onto the base branch", ox.Short("r")).
					Bool("squash", "Squash the commits into one commit and merge it into the base branch", ox.Short("s")).
					String("subject", "Subject text for the merge commit", ox.Default("text"), ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Mark a pull request as ready for review.\n\nWithout an argument, the pull request that belongs to the current branch\nis marked as ready.\n\nIf supported by your plan, convert to draft with `--undo`"),
				ox.Usage("ready", "Mark a pull request as ready for review"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("undo", "Convert a pull request to \"draft\"").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Reopen a pull request"),
				ox.Usage("reopen", "Reopen a pull request"),
				ox.Spec("{<number> | <url> | <branch>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("comment", "Add a reopening comment", ox.Short("c")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Add a review to a pull request.\n\nWithout an argument, the pull request that belongs to the current branch is reviewed."),
				ox.Usage("review", "Add a review to a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Example("\n  # approve the pull request of the current branch\n  $ gh pr review --approve\n  \n  # leave a review comment for the current branch\n  $ gh pr review --comment -b \"interesting\"\n  \n  # add a review for a specific pull request\n  $ gh pr review 123\n  \n  # request changes on a specific pull request\n  $ gh pr review 123 -r -b \"needs more ASCII art\""),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("approve", "Approve pull request", ox.Short("a")).
					String("body", "Specify the body of a review", ox.Short("b")).
					String("body-file", "Read body text from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("comment", "Comment on a pull request", ox.Short("c")).
					Bool("request-changes", "Request changes on a pull request", ox.Short("r")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Unlock pull request conversation"),
				ox.Usage("unlock", "Unlock pull request conversation"),
				ox.Spec("{<number> | <url>} [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Update a pull request branch with latest changes of the base branch.\n\nWithout an argument, the pull request that belongs to the current branch is selected.\n\nThe default behavior is to update with a merge commit (i.e., merging the base branch\ninto the the PR's branch). To reconcile the changes with rebasing on top of the base\nbranch, the `--rebase` option should be provided."),
				ox.Usage("update-branch", "Update a pull request branch"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Example("\n  $ gh pr update-branch 23\n  $ gh pr update-branch 23 --rebase\n  $ gh pr update-branch 23 --repo owner/repo"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("rebase", "Update PR branch by rebasing on top of latest base branch").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Display the title, body, and other information about a pull request.\n\nWithout an argument, the pull request that belongs to the current branch\nis displayed.\n\nWith `--web` flag, open the pull request in a web browser instead.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View a pull request"),
				ox.Spec("[<number> | <url> | <branch>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("comments", "View pull request comments", ox.Short("c")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open a pull request in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Work with GitHub Projects. Note that the token you are using must have 'project' scope, which is not set by default. You can verify your token scope by running 'gh auth status' and add the project scope by running 'gh auth refresh -s project'."),
			ox.Usage("project", "Work with GitHub Projects."),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh project create --owner monalisa --title \"Roadmap\"\n  $ gh project view 1 --owner cli --web\n  $ gh project field-list 1 --owner cli\n  $ gh project item-list 1 --owner cli"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Close a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("close", "Close a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # close project \"1\" owned by monalisa\n  gh project close 1 --owner monalisa\n  \n  # reopen closed project \"1\" owned by github\n  gh project close 1 --owner github --undo"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("undo", "Reopen a closed project"),
			), ox.Sub(
				ox.Banner("Copy a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("copy", "Copy a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # copy project \"1\" owned by monalisa to github\n  gh project copy 1 --source-owner monalisa --target-owner github --title \"a new project\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("drafts", "Include draft issues when copying").
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("source-owner", "Login of the source owner. Use \"@me\" for the current user.").
					String("target-owner", "Login of the target owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("title", "Title for the new project"),
			), ox.Sub(
				ox.Banner("Create a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("create", "Create a project"),
				ox.Spec("[flags]"),
				ox.Example("\n  # create a new project owned by login monalisa\n  gh project create --owner monalisa --title \"a new project\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("title", "Title for the project"),
			), ox.Sub(
				ox.Banner("Delete a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("delete", "Delete a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # delete the current user's project \"1\"\n  gh project delete 1 --owner \"@me\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Edit a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("edit", "Edit a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # edit the title of monalisa's project \"1\"\n  gh project edit 1 --owner monalisa --title \"New title\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("description", "New description of the project", ox.Short("d")).
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("readme", "New readme for the project").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("title", "New title for the project").
					String("visibility", "Change project visibility: {PUBLIC|PRIVATE}"),
			), ox.Sub(
				ox.Banner("Create a field in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("field-create", "Create a field in a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # create a field in the current user's project \"1\"\n  gh project field-create 1 --owner \"@me\" --name \"new field\" --data-type \"text\"\n  \n  # create a field with three options to select from for owner monalisa\n  gh project field-create 1 --owner monalisa --name \"new field\" --data-type \"SINGLE_SELECT\" --single-select-options \"one,two,three\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("data-type", "DataType of the new field.: {TEXT|SINGLE_SELECT|DATE|NUMBER}").
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("name", "Name of the new field").
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					Slice("single-select-options", "Options for SINGLE_SELECT data type", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Delete a field in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("field-delete", "Delete a field in a project"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("id", "ID of the field to delete").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("List the fields in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("field-list", "List the fields in a project"),
				ox.Spec("number [flags]"),
				ox.Example("\n  # list fields in the current user's project \"1\"\n  gh project field-list 1 --owner \"@me\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Int("limit", "Maximum number of fields to fetch", ox.Default("30"), ox.Short("L")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Add a pull request or an issue to a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-add", "Add a pull request or an issue to a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # add an item to monalisa's project \"1\"\n  gh project item-add 1 --owner monalisa --url https://github.com/monalisa/myproject/issues/23"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("url", "URL of the issue or pull request to add to the project"),
			), ox.Sub(
				ox.Banner("Archive an item in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-archive", "Archive an item in a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # archive an item in the current user's project \"1\"\n  gh project item-archive 1 --owner \"@me\" --id <item-ID>"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("id", "ID of the item to archive").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("undo", "Unarchive an item"),
			), ox.Sub(
				ox.Banner("Create a draft issue item in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-create", "Create a draft issue item in a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # create a draft issue in the current user's project \"1\"\n  gh project item-create 1 --owner \"@me\" --title \"new item\" --body \"new item body\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("body", "Body for the draft issue").
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("title", "Title for the draft issue"),
			), ox.Sub(
				ox.Banner("Delete an item from a project by ID\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-delete", "Delete an item from a project by ID"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # delete an item in the current user's project \"1\"\n  gh project item-delete 1 --owner \"@me\" --id <item-ID>"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("id", "ID of the item to delete").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Edit either a draft issue or a project item. Both usages require the ID of the item to edit.\n\nFor non-draft issues, the ID of the project is also required, and only a single field value can be updated per invocation.\n\nRemove project item field value using `--clear` flag.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-edit", "Edit an item in a project"),
				ox.Spec("[flags]"),
				ox.Example("\n  # edit an item's text field value\n  gh project item-edit --id <item-ID> --field-id <field-ID> --project-id <project-ID> --text \"new text\"\n  \n  # clear an item's field value\n  gh project item-edit --id <item-ID> --field-id <field-ID> --project-id <project-ID> --clear"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("body", "Body of the draft issue item").
					Bool("clear", "Remove field value").
					String("date", "Date value for the field (YYYY-MM-DD)").
					String("field-id", "ID of the field to update").
					String("format", "Output format: {json}").
					String("id", "ID of the item to edit").
					String("iteration-id", "ID of the iteration value to set on the field").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Float32("number", "Number value for the field").
					String("project-id", "ID of the project to which the field belongs to").
					String("single-select-option-id", "ID of the single select option value to set on the field").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("text", "Text value for the field").
					String("title", "Title of the draft issue item"),
			), ox.Sub(
				ox.Banner("List the items in a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("item-list", "List the items in a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # list the items in the current users's project \"1\"\n  gh project item-list 1 --owner \"@me\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Int("limit", "Maximum number of items to fetch", ox.Default("30"), ox.Short("L")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Link a project to a repository or a team"),
				ox.Usage("link", "Link a project to a repository or a team"),
				ox.Spec("[<number>] [flag] [flags]"),
				ox.Example("\n  # link monalisa's project 1 to her repository \"my_repo\"\n  gh project link 1 --owner monalisa --repo my_repo\n  \n  # link monalisa's organization's project 1 to her team \"my_team\"\n  gh project link 1 --owner my_organization --team my_team\n  \n  # link monalisa's project 1 to the repository of current directory if neither --repo nor --team is specified\n  gh project link 1"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("repo", "The repository to be linked to this project", ox.Short("R")).
					String("team", "The team to be linked to this project", ox.Short("T")),
			), ox.Sub(
				ox.Banner("List the projects for an owner\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List the projects for an owner"),
				ox.Spec("[flags]"),
				ox.Aliases("gh project ls"),
				ox.Example("\n  # list the current user's projects\n  gh project list\n  \n  # list the projects for org github including closed projects\n  gh project list --owner github --closed"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("closed", "Include closed projects").
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Int("limit", "Maximum number of projects to fetch", ox.Default("30"), ox.Short("L")).
					String("owner", "Login of the owner").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open projects list in the browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Mark a project as a template\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("mark-template", "Mark a project as a template"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # mark the github org's project \"1\" as a template\n  gh project mark-template 1 --owner \"github\"\n  \n  # unmark the github org's project \"1\" as a template\n  gh project mark-template 1 --owner \"github\" --undo"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the org owner.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("undo", "Unmark the project as a template."),
			), ox.Sub(
				ox.Banner("Unlink a project from a repository or a team"),
				ox.Usage("unlink", "Unlink a project from a repository or a team"),
				ox.Spec("[<number>] [flag] [flags]"),
				ox.Example("\n  # unlink monalisa's project 1 from her repository \"my_repo\"\n  gh project unlink 1 --owner monalisa --repo my_repo\n  \n  # unlink monalisa's organization's project 1 from her team \"my_team\"\n  gh project unlink 1 --owner my_organization --team my_team\n  \n  # unlink monalisa's project 1 from the repository of current directory if neither --repo nor --team is specified\n  gh project unlink 1"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("repo", "The repository to be unlinked from this project", ox.Short("R")).
					String("team", "The team to be unlinked from this project", ox.Short("T")),
			), ox.Sub(
				ox.Banner("View a project\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View a project"),
				ox.Spec("[<number>] [flags]"),
				ox.Example("\n  # view the current user's project \"1\"\n  gh project view 1\n  \n  # open user monalisa's project \"1\" in the browser\n  gh project view 1 --owner monalisa --web"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("format", "Output format: {json}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					String("owner", "Login of the owner. Use \"@me\" for the current user.").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open a project in the browser", ox.Short("w")),
			)), ox.Sub(
			ox.Banner("Manage releases"),
			ox.Usage("release", "Manage releases"),
			ox.Spec("<command> [flags]"),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Create a new GitHub Release for a repository.\n\nA list of asset files may be given to upload to the new release. To define a\ndisplay label for an asset, append text starting with `#` after the file name.\n\nIf a matching git tag does not yet exist, one will automatically get created\nfrom the latest state of the default branch.\nUse `--target` to point to a different branch or commit for the automatic tag creation.\nUse `--verify-tag` to abort the release if the tag doesn't already exist.\nTo fetch the new tag locally after the release, do `git fetch --tags origin`.\n\nTo create a release from an annotated git tag, first create one locally with\ngit, push the tag to GitHub, then run this command.\nUse `--notes-from-tag` to automatically generate the release notes\nfrom the annotated git tag.\n\nWhen using automatically generated release notes, a release title will also be automatically\ngenerated unless a title was explicitly passed. Additional release notes can be prepended to\nautomatically generated notes by using the `--notes` flag."),
				ox.Usage("create", "Create a new release"),
				ox.Spec("[<tag>] [<files>...]"),
				ox.Aliases("gh release new"),
				ox.Example("\n  Interactively create a release\n  $ gh release create\n  \n  Interactively create a release from specific tag\n  $ gh release create v1.2.3\n  \n  Non-interactively create a release\n  $ gh release create v1.2.3 --notes \"bugfix release\"\n  \n  Use automatically generated release notes\n  $ gh release create v1.2.3 --generate-notes\n  \n  Use release notes from a file\n  $ gh release create v1.2.3 -F release-notes.md\n  \n  Use annotated tag notes\n  $ gh release create v1.2.3 --notes-from-tag\n  \n  Don't mark the release as latest\n  $ gh release create v1.2.3 --latest=false \n  \n  Upload all tarballs in a directory as release assets\n  $ gh release create v1.2.3 ./dist/*.tgz\n  \n  Upload a release asset with a display label\n  $ gh release create v1.2.3 '/path/to/asset.zip#My display label'\n  \n  Create a release and start a discussion\n  $ gh release create v1.2.3 --discussion-category \"General\""),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("discussion-category", "Start a discussion in the specified category").
					Bool("draft", "Save the release as a draft instead of publishing it", ox.Short("d")).
					Bool("generate-notes", "Automatically generate title and notes for the release").
					Bool("latest", "Mark this release as \"Latest\" (default [automatic based on date and version]). --latest=false to explicitly NOT set as latest").
					String("notes", "Release notes", ox.Short("n")).
					String("notes-file", "Read release notes from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("notes-from-tag", "Automatically generate notes from annotated tag").
					String("notes-start-tag", "Tag to use as the starting point for generating release notes").
					Bool("prerelease", "Mark the release as a prerelease", ox.Short("p")).
					String("target", "Target branch or full commit SHA", ox.Spec("branch"), ox.Default("[main branch]")).
					String("title", "Release title", ox.Short("t")).
					Bool("verify-tag", "Abort in case the git tag doesn't already exist in the remote repository").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List releases in a repository\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List releases in a repository"),
				ox.Spec("[flags]"),
				ox.Aliases("gh release ls"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("exclude-drafts", "Exclude draft releases").
					Bool("exclude-pre-releases", "Exclude pre-releases").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of items to fetch", ox.Default("30"), ox.Short("L")).
					String("order", "Order of releases returned: {asc|desc}", ox.Default("desc"), ox.Short("O")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Delete a release"),
				ox.Usage("delete", "Delete a release"),
				ox.Spec("<tag> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("cleanup-tag", "Delete the specified tag in addition to its release").
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Delete an asset from a release"),
				ox.Usage("delete-asset", "Delete an asset from a release"),
				ox.Spec("<tag> <asset-name> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Download assets from a GitHub release.\n\nWithout an explicit tag name argument, assets are downloaded from the\nlatest release in the project. In this case, `--pattern` or `--archive`\nis required."),
				ox.Usage("download", "Download release assets"),
				ox.Spec("[<tag>] [flags]"),
				ox.Example("\n  # download all assets from a specific release\n  $ gh release download v1.2.3\n  \n  # download only Debian packages for the latest release\n  $ gh release download --pattern '*.deb'\n  \n  # specify multiple file patterns\n  $ gh release download -p '*.deb' -p '*.rpm'\n  \n  # download the archive of the source code for a release\n  $ gh release download v1.2.3 --archive=zip"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("archive", "Download the source code archive in the specified format (zip or tar.gz)", ox.Spec("format"), ox.Short("A")).
					Bool("clobber", "Overwrite existing files of the same name").
					String("dir", "The directory to download files into", ox.Spec("directory"), ox.Default("."), ox.Short("D")).
					String("output", "The file to write a single asset to (use \"-\" to write to standard output)", ox.Spec("file"), ox.Short("O")).
					Slice("pattern", "Download only assets that match a glob pattern", ox.Elem(ox.StringT), ox.Short("p")).
					Bool("skip-existing", "Skip downloading when files of the same name exist").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Edit a release"),
				ox.Usage("edit", "Edit a release"),
				ox.Spec("<tag>"),
				ox.Example("\n  Publish a release that was previously a draft\n  $ gh release edit v1.0 --draft=false\n  \n  Update the release notes from the content of a file\n  $ gh release edit v1.0 --notes-file /path/to/release_notes.md"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("discussion-category", "Start a discussion in the specified category when publishing a draft").
					Bool("draft", "Save the release as a draft instead of publishing it").
					Bool("latest", "Explicitly mark the release as \"Latest\"").
					String("notes", "Release notes", ox.Short("n")).
					String("notes-file", "Read release notes from file (use \"-\" to read from standard input)", ox.Spec("file"), ox.Short("F")).
					Bool("prerelease", "Mark the release as a prerelease").
					String("tag", "The name of the tag").
					String("target", "Target branch or full commit SHA", ox.Spec("branch"), ox.Default("[main branch]")).
					String("title", "Release title", ox.Short("t")).
					Bool("verify-tag", "Abort in case the git tag doesn't already exist in the remote repository").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Upload asset files to a GitHub Release.\n\nTo define a display label for an asset, append text starting with `#` after the\nfile name."),
				ox.Usage("upload", "Upload assets to a release"),
				ox.Spec("<tag> <files>... [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("clobber", "Overwrite existing assets of the same name").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("View information about a GitHub Release.\n\nWithout an explicit tag name argument, the latest release in the project\nis shown.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View information about a release"),
				ox.Spec("[<tag>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open the release in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Work with GitHub repositories."),
			ox.Usage("repo", "Manage repositories"),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh repo create\n  $ gh repo clone cli/cli\n  $ gh repo view --web"),
			ox.Sections("GENERAL COMMANDS", "TARGETED COMMANDS"),
			ox.Section(0),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Create a new GitHub repository.\n\nTo create a repository interactively, use `gh repo create` with no arguments.\n\nTo create a remote repository non-interactively, supply the repository name and one of `--public`, `--private`, or `--internal`.\nPass `--clone` to clone the new repository locally.\n\nIf the `OWNER/` portion of the `OWNER/REPO` name argument is omitted, it\ndefaults to the name of the authenticating user.\n\nTo create a remote repository from an existing local repository, specify the source directory with `--source`.\nBy default, the remote repository name will be the name of the source directory.\n\nPass `--push` to push any local commits to the new repository. If the repo is bare, this will mirror all refs.\n\nFor language or platform .gitignore templates to use with `--gitignore`, <https://github.com/github/gitignore>.\n\nFor license keywords to use with `--license`, run `gh repo license list` or visit <https://choosealicense.com>."),
				ox.Usage("create", "Create a new repository"),
				ox.Spec("[<name>] [flags]"),
				ox.Aliases("gh repo new"),
				ox.Example("\n  # create a repository interactively\n  gh repo create\n  \n  # create a new remote repository and clone it locally\n  gh repo create my-project --public --clone\n  \n  # create a new remote repository in a different organization\n  gh repo create my-org/my-project --public\n  \n  # create a remote repository from the current directory\n  gh repo create my-project --private --source=. --remote=upstream"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("add-readme", "Add a README file to the new repository").
					Bool("clone", "Clone the new repository to the current directory", ox.Short("c")).
					String("description", "Description of the repository", ox.Short("d")).
					Bool("disable-issues", "Disable issues in the new repository").
					Bool("disable-wiki", "Disable wiki in the new repository").
					String("gitignore", "Specify a gitignore template for the repository", ox.Short("g")).
					URL("homepage", "Repository home page URL", ox.Short("h")).
					Bool("include-all-branches", "Include all branches from template repository").
					Bool("internal", "Make the new repository internal").
					String("license", "Specify an Open Source License for the repository", ox.Short("l")).
					Bool("private", "Make the new repository private").
					Bool("public", "Make the new repository public").
					Bool("push", "Push local commits to the new repository").
					String("remote", "Specify remote name for the new repository", ox.Short("r")).
					String("source", "Specify path to local repository to use as source", ox.Short("s")).
					String("team", "The name of the organization team to be granted access", ox.Spec("name"), ox.Short("t")).
					String("template", "Make the new repository based on a template repository", ox.Spec("repository"), ox.Short("p")),
			), ox.Sub(
				ox.Banner("List repositories owned by a user or organization.\n\nNote that the list will only include repositories owned by the provided argument,\nand the `--fork` or `--source` flags will not traverse ownership boundaries. For example,\nwhen listing the forks in an organization, the output would not include those owned by individual users.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List repositories owned by user or organization"),
				ox.Spec("[<owner>] [flags]"),
				ox.Aliases("gh repo ls"),
				ox.Section(0),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("archived", "Show only archived repositories").
					Bool("fork", "Show only forks").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("language", "Filter by primary coding language", ox.Short("l")).
					Int("limit", "Maximum number of repositories to list", ox.Default("30"), ox.Short("L")).
					Bool("no-archived", "Omit archived repositories").
					Bool("source", "Show only non-forks").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Slice("topic", "Filter by topic", ox.Elem(ox.StringT)).
					String("visibility", "Filter by repository visibility: {public|private|internal}"),
			), ox.Sub(
				ox.Banner("Archive a GitHub repository.\n\nWith no argument, archives the current repository."),
				ox.Usage("archive", "Archive a repository"),
				ox.Spec("[<repository>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")),
			), ox.Sub(
				ox.Banner("Clone a GitHub repository locally. Pass additional `git clone` flags by listing\nthem after `--`.\n\nIf the `OWNER/` portion of the `OWNER/REPO` repository argument is omitted, it\ndefaults to the name of the authenticating user.\n\nWhen a protocol scheme is not provided in the repository argument, the `git_protocol` will be\nchosen from your configuration, which can be checked via `gh config get git_protocol`. If the protocol\nscheme is provided, the repository will be cloned using the specified protocol.\n\nIf the repository is a fork, its parent repository will be added as an additional\ngit remote called `upstream`. The remote name can be configured using `--upstream-remote-name`.\nThe `--upstream-remote-name` option supports an `@owner` value which will name\nthe remote after the owner of the parent repository.\n\nIf the repository is a fork, its parent repository will be set as the default remote repository."),
				ox.Usage("clone", "Clone a repository locally"),
				ox.Spec("<repository> [<directory>] [-- <gitflags>...]"),
				ox.Example("\n  # Clone a repository from a specific org\n  $ gh repo clone cli/cli\n  \n  # Clone a repository from your own account\n  $ gh repo clone myrepo\n  \n  # Clone a repo, overriding git protocol configuration\n  $ gh repo clone https://github.com/cli/cli\n  $ gh repo clone git@github.com:cli/cli.git\n  \n  # Clone a repository to a custom directory\n  $ gh repo clone cli/cli workspace/cli\n  \n  # Clone a repository with additional git clone flags\n  $ gh repo clone cli/cli -- --depth=1"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("upstream-remote-name", "Upstream remote name when cloning a fork", ox.Default("upstream"), ox.Short("u")),
			), ox.Sub(
				ox.Banner("Delete a GitHub repository.\n\nWith no argument, deletes the current repository. Otherwise, deletes the specified repository.\n\nDeletion requires authorization with the `delete_repo` scope. \nTo authorize, run `gh auth refresh -s delete_repo`"),
				ox.Usage("delete", "Delete a repository"),
				ox.Spec("[<repository>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "confirm deletion without prompting"),
			), ox.Sub(
				ox.Banner("Manage deploy keys in a repository"),
				ox.Usage("deploy-key", "Manage deploy keys in a repository"),
				ox.Spec("<command> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Sub(
					ox.Banner("Add a deploy key to a GitHub repository.\n\nNote that any key added by gh will be associated with the current authentication token.\nIf you de-authorize the GitHub CLI app or authentication token from your account, any\ndeploy keys added by GitHub CLI will be removed as well."),
					ox.Usage("add", "Add a deploy key to a GitHub repository"),
					ox.Spec("<key-file> [flags]"),
					ox.Example("\n  # generate a passwordless SSH key and add it as a deploy key to a repository\n  $ ssh-keygen -t ed25519 -C \"my description\" -N \"\" -f ~/.ssh/gh-test\n  $ gh repo deploy-key add ~/.ssh/gh-test.pub"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						Bool("allow-write", "Allow write access for the key", ox.Short("w")).
						String("title", "Title of the new key", ox.Short("t")).
						String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
				), ox.Sub(
					ox.Banner("Delete a deploy key from a GitHub repository"),
					ox.Usage("delete", "Delete a deploy key from a GitHub repository"),
					ox.Spec("<key-id> [flags]"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
				), ox.Sub(
					ox.Banner("List deploy keys in a GitHub repository\n\nFor more information about output formatting flags, see `gh help formatting`."),
					ox.Usage("list", "List deploy keys in a GitHub repository"),
					ox.Spec("[flags]"),
					ox.Aliases("gh repo deploy-key ls"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
						Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
						String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
						String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
				), ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Edit repository settings.\n\nTo toggle a setting off, use the `--<flag>=false` syntax.\n\nChanging repository visibility can have unexpected consequences including but not limited to:\n\n- Losing stars and watchers, affecting repository ranking\n- Detaching public forks from the network\n- Disabling push rulesets\n- Allowing access to GitHub Actions history and logs\n\nWhen the `--visibility` flag is used, `--accept-visibility-change-consequences` flag is required.\n\nFor information on all the potential consequences, see <https://gh.io/setting-repository-visibility>"),
				ox.Usage("edit", "Edit repository settings"),
				ox.Spec("[<repository>] [flags]"),
				ox.Example("\n  # enable issues and wiki\n  gh repo edit --enable-issues --enable-wiki\n  \n  # disable projects\n  gh repo edit --enable-projects=false"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("accept-visibility-change-consequences", "Accept the consequences of changing the repository visibility").
					Slice("add-topic", "Add repository topic", ox.Elem(ox.StringT)).
					Bool("allow-forking", "Allow forking of an organization repository").
					Bool("allow-update-branch", "Allow a pull request head branch that is behind its base branch to be updated").
					String("default-branch", "Set the default branch name for the repository", ox.Spec("name")).
					Bool("delete-branch-on-merge", "Delete head branch when pull requests are merged").
					String("description", "Description of the repository", ox.Short("d")).
					Bool("enable-auto-merge", "Enable auto-merge functionality").
					Bool("enable-discussions", "Enable discussions in the repository").
					Bool("enable-issues", "Enable issues in the repository").
					Bool("enable-merge-commit", "Enable merging pull requests via merge commit").
					Bool("enable-projects", "Enable projects in the repository").
					Bool("enable-rebase-merge", "Enable merging pull requests via rebase").
					Bool("enable-squash-merge", "Enable merging pull requests via squashed commit").
					Bool("enable-wiki", "Enable wiki in the repository").
					URL("homepage", "Repository home page URL", ox.Short("h")).
					Slice("remove-topic", "Remove repository topic", ox.Elem(ox.StringT)).
					Bool("template", "Make the repository available as a template repository").
					String("visibility", "Change the visibility of the repository to {public,private,internal}"),
			), ox.Sub(
				ox.Banner("Create a fork of a repository.\n\nWith no argument, creates a fork of the current repository. Otherwise, forks\nthe specified repository.\n\nBy default, the new fork is set to be your `origin` remote and any existing\norigin remote is renamed to `upstream`. To alter this behavior, you can set\na name for the new fork's remote with `--remote-name`.\n\nThe `upstream` remote will be set as the default remote repository.\n\nAdditional `git clone` flags can be passed after `--`."),
				ox.Usage("fork", "Create a fork of a repository"),
				ox.Spec("[<repository>] [-- <gitflags>...] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("clone", "Clone the fork").
					Bool("default-branch-only", "Only include the default branch in the fork").
					String("fork-name", "Rename the forked repository").
					String("org", "Create the fork in an organization").
					Bool("remote", "Add a git remote for the fork").
					String("remote-name", "Specify the name for the new remote", ox.Default("origin")),
			), ox.Sub(
				ox.Banner("List and view available repository gitignore templates"),
				ox.Usage("gitignore", "List and view available repository gitignore templates"),
				ox.Spec("<command> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Sub(
					ox.Banner("List available repository gitignore templates"),
					ox.Usage("list", "List available repository gitignore templates"),
					ox.Spec("[flags]"),
					ox.Aliases("gh repo gitignore ls"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				), ox.Sub(
					ox.Banner("View an available repository `.gitignore` template.\n\n`<template>` is a case-sensitive `.gitignore` template name.\n\nFor a list of available templates, run `gh repo gitignore list`."),
					ox.Usage("view", "View an available repository gitignore template"),
					ox.Spec("<template> [flags]"),
					ox.Example("\n  # View the Go gitignore template\n  gh repo gitignore view Go\n  \n  # View the Python gitignore template\n  gh repo gitignore view Python\n  \n  # Create a new .gitignore file using the Go template\n  gh repo gitignore view Go > .gitignore\n  \n  # Create a new .gitignore file using the Python template\n  gh repo gitignore view Python > .gitignore"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				)), ox.Sub(
				ox.Banner("Explore repository licenses"),
				ox.Usage("license", "Explore repository licenses"),
				ox.Spec("<command> [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Sub(
					ox.Banner("List common repository licenses.\n\nFor even more licenses, visit <https://choosealicense.com/appendix>"),
					ox.Usage("list", "List common repository licenses"),
					ox.Spec("[flags]"),
					ox.Aliases("gh repo license ls"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				), ox.Sub(
					ox.Banner("View a specific repository license by license key or SPDX ID.\n\nRun `gh repo license list` to see available commonly used licenses. For even more licenses, visit <https://choosealicense.com/appendix>."),
					ox.Usage("view", "View a specific repository license"),
					ox.Spec("{<license-key> | <SPDX-ID>} [flags]"),
					ox.Example("\n  # View the MIT license from SPDX ID\n  gh repo license view MIT\n  \n  # View the MIT license from license key\n  gh repo license view mit\n  \n  # View the GNU AGPL-3.0 license from SPDX ID\n  gh repo license view AGPL-3.0\n  \n  # View the GNU AGPL-3.0 license from license key\n  gh repo license view agpl-3.0\n  \n  # Create a LICENSE.md with the MIT license\n  gh repo license view MIT > LICENSE.md"),
					ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
					ox.Flags().
						Bool("web", "Open https://choosealicense.com/ in the browser", ox.Short("w")),
				)), ox.Sub(
				ox.Banner("Rename a GitHub repository.\n\nBy default, this renames the current repository; otherwise renames the specified repository."),
				ox.Usage("rename", "Rename a repository"),
				ox.Spec("[<new-name>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")).
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")),
			), ox.Sub(
				ox.Banner("This command sets the default remote repository to use when querying the\nGitHub API for the locally cloned repository.\n\ngh uses the default repository for things like:\n\n - viewing and creating pull requests\n - viewing and creating issues\n - viewing and creating releases\n - working with GitHub Actions\n - adding repository and environment secrets"),
				ox.Usage("set-default", "Configure default repository for this directory"),
				ox.Spec("[<repository>] [flags]"),
				ox.Example("\n  Interactively select a default repository:\n  $ gh repo set-default\n  \n  Set a repository explicitly:\n  $ gh repo set-default owner/repo\n  \n  View the current default repository:\n  $ gh repo set-default --view\n  \n  Show more repository options in the interactive picker:\n  $ git remote add newrepo https://github.com/owner/repo\n  $ gh repo set-default"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("unset", "unset the current default repository", ox.Short("u")).
					Bool("view", "view the current default repository", ox.Short("v")),
			), ox.Sub(
				ox.Banner("Sync destination repository from source repository. Syncing uses the default branch\nof the source repository to update the matching branch on the destination\nrepository so they are equal. A fast forward update will be used except when the\n`--force` flag is specified, then the two branches will\nbe synced using a hard reset.\n\nWithout an argument, the local repository is selected as the destination repository.\n\nThe source repository is the parent of the destination repository by default.\nThis can be overridden with the `--source` flag."),
				ox.Usage("sync", "Sync a repository"),
				ox.Spec("[<destination-repository>] [flags]"),
				ox.Example("\n  # Sync local repository from remote parent\n  $ gh repo sync\n  \n  # Sync local repository from remote parent on specific branch\n  $ gh repo sync --branch v1\n  \n  # Sync remote fork from its parent\n  $ gh repo sync owner/cli-fork\n  \n  # Sync remote repository from another remote repository\n  $ gh repo sync owner/repo --source owner2/repo2"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("branch", "Branch to sync", ox.Default("[default branch]"), ox.Short("b")).
					Bool("force", "Hard reset the branch of the destination repository to match the source repository").
					String("source", "Source repository", ox.Short("s")),
			), ox.Sub(
				ox.Banner("Unarchive a GitHub repository.\n\nWith no argument, unarchives the current repository."),
				ox.Usage("unarchive", "Unarchive a repository"),
				ox.Spec("[<repository>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")),
			), ox.Sub(
				ox.Banner("Display the description and the README of a GitHub repository.\n\nWith no argument, the repository for the current directory is displayed.\n\nWith `--web`, open the repository in a web browser instead.\n\nWith `--branch`, view a specific branch of the repository.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View a repository"),
				ox.Spec("[<repository>] [flags]"),
				ox.Section(1),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("branch", "View a specific branch of the repository", ox.Short("b")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open a repository in the browser", ox.Short("w")),
			)), ox.Sub(
			ox.Banner("Work with GitHub Actions caches."),
			ox.Usage("cache", "Manage GitHub Actions caches"),
			ox.Spec("<command> [flags]"),
			ox.Example("\n  $ gh cache list\n  $ gh cache delete --all"),
			ox.Section(1),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Delete GitHub Actions caches.\n\nDeletion requires authorization with the `repo` scope."),
				ox.Usage("delete", "Delete GitHub Actions caches"),
				ox.Spec("[<cache-id>| <cache-key> | --all] [flags]"),
				ox.Example("\n  # Delete a cache by id\n  $ gh cache delete 1234\n  \n  # Delete a cache by key\n  $ gh cache delete cache-key\n  \n  # Delete a cache by id in a specific repo\n  $ gh cache delete 1234 --repo cli/cli\n  \n  # Delete all caches\n  $ gh cache delete --all"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Delete all caches", ox.Short("a")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List GitHub Actions caches\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List GitHub Actions caches"),
				ox.Spec("[flags]"),
				ox.Aliases("gh cache ls"),
				ox.Example("\n  # List caches for current repository\n  $ gh cache list\n  \n  # List caches for specific repository\n  $ gh cache list --repo cli/cli\n  \n  # List caches sorted by least recently accessed\n  $ gh cache list --sort last_accessed_at --order asc\n  \n  # List caches that have keys matching a prefix (or that match exactly)\n  $ gh cache list --key key-prefix\n  \n  # To list caches for a specific branch, replace <branch-name> with the actual branch name\n  $ gh cache list --ref refs/heads/<branch-name>\n  \n  # To list caches for a specific pull request, replace <pr-number> with the actual pull request number\n  $ gh cache list --ref refs/pull/<pr-number>/merge"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("key", "Filter by cache key prefix", ox.Short("k")).
					Int("limit", "Maximum number of caches to fetch", ox.Default("30"), ox.Short("L")).
					String("order", "Order of caches returned: {asc|desc}", ox.Default("desc"), ox.Short("O")).
					String("ref", "Filter by ref, formatted as refs/heads/<branch name> or refs/pull/<number>/merge", ox.Short("r")).
					String("sort", "Sort fetched caches: {created_at|last_accessed_at|size_in_bytes}", ox.Default("last_accessed_at"), ox.Short("S")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("List, view, and watch recent workflow runs from GitHub Actions."),
			ox.Usage("run", "View details about workflow runs"),
			ox.Spec("<command> [flags]"),
			ox.Section(1),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Cancel a workflow run"),
				ox.Usage("cancel", "Cancel a workflow run"),
				ox.Spec("[<run-id>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Delete a workflow run"),
				ox.Usage("delete", "Delete a workflow run"),
				ox.Spec("[<run-id>] [flags]"),
				ox.Example("\n  # Interactively select a run to delete\n  $ gh run delete\n  \n  # Delete a specific run\n  $ gh run delete 12345"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Download artifacts generated by a GitHub Actions workflow run.\n\nThe contents of each artifact will be extracted under separate directories based on\nthe artifact name. If only a single artifact is specified, it will be extracted into\nthe current directory.\n\nBy default, this command downloads the latest artifact created and uploaded through\nGitHub Actions. Because workflows can delete or overwrite artifacts, `<run-id>`\nmust be used to select an artifact from a specific workflow run."),
				ox.Usage("download", "Download artifacts generated by a workflow run"),
				ox.Spec("[<run-id>] [flags]"),
				ox.Example("\n  # Download all artifacts generated by a workflow run\n  $ gh run download <run-id>\n  \n  # Download a specific artifact within a run\n  $ gh run download <run-id> -n <name>\n  \n  # Download specific artifacts across all runs in a repository\n  $ gh run download -n <name1> -n <name2>\n  \n  # Select artifacts to download interactively\n  $ gh run download"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("dir", "The directory to download artifacts into", ox.Default("."), ox.Short("D")).
					Slice("name", "Download artifacts that match any of the given names", ox.Elem(ox.StringT), ox.Short("n")).
					Slice("pattern", "Download artifacts that match a glob pattern", ox.Elem(ox.StringT), ox.Short("p")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List recent workflow runs.\n\nNote that providing the `workflow_name` to the `-w` flag will not fetch disabled workflows.\nAlso pass the `-a` flag to fetch disabled workflow runs using the `workflow_name` and the `-w` flag.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List recent workflow runs"),
				ox.Spec("[flags]"),
				ox.Aliases("gh run ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Include disabled workflows", ox.Short("a")).
					String("branch", "Filter runs by branch", ox.Short("b")).
					String("commit", "Filter runs by the SHA of the commit", ox.Spec("SHA"), ox.Short("c")).
					Date("created", "Filter runs by the date it was created").
					String("event", "Filter runs by which event triggered the run", ox.Spec("event"), ox.Short("e")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of runs to fetch", ox.Default("20"), ox.Short("L")).
					String("status", "Filter runs by status: {queued|completed|in_progress|requested|waiting|action_required|cancelled|failure|neutral|skipped|stale|startup_failure|success|timed_out}", ox.Short("s")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("user", "Filter runs by user who triggered the run", ox.Short("u")).
					String("workflow", "Filter runs by workflow", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Rerun an entire run, only failed jobs, or a specific job from a run.\n\nNote that due to historical reasons, the `--job` flag may not take what you expect.\nSpecifically, when navigating to a job in the browser, the URL looks like this:\n`https://github.com/<owner>/<repo>/actions/runs/<run-id>/jobs/<number>`.\n\nHowever, this `<number>` should not be used with the `--job` flag and will result in the\nAPI returning `404 NOT FOUND`. Instead, you can get the correct job IDs using the following command:\n\n\tgh run view <run-id> --json jobs --jq '.jobs[] | {name, databaseId}'\n\nYou will need to use databaseId field for triggering job re-runs."),
				ox.Usage("rerun", "Rerun a run"),
				ox.Spec("[<run-id>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("debug", "Rerun with debug logging", ox.Short("d")).
					Bool("failed", "Rerun only failed jobs, including dependencies").
					String("job", "Rerun a specific job ID from a run, including dependencies", ox.Short("j")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("View a summary of a workflow run.\n\nThis command does not support authenticating via fine grained PATs\nas it is not currently possible to create a PAT with the `checks:read` permission.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("view", "View a summary of a workflow run"),
				ox.Spec("[<run-id>] [flags]"),
				ox.Example("\n  # Interactively select a run to view, optionally selecting a single job\n  $ gh run view\n  \n  # View a specific run\n  $ gh run view 12345\n  \n  # View a specific run with specific attempt number\n  $ gh run view 12345 --attempt 3\n  \n  # View a specific job within a run\n  $ gh run view --job 456789\n  \n  # View the full log for a specific job\n  $ gh run view --log --job 456789\n  \n  # Exit non-zero if a run failed\n  $ gh run view 0451 --exit-status && echo \"run pending or passed\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Uint("attempt", "The attempt number of the workflow run", ox.Short("a")).
					Bool("exit-status", "Exit with non-zero status if run failed").
					String("job", "View a specific job ID from a run", ox.Short("j")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Bool("log", "View full log for either a run or specific job").
					Bool("log-failed", "View the log for any failed steps in a run or specific job").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("verbose", "Show job steps", ox.Short("v")).
					Bool("web", "Open run in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Watch a run until it completes, showing its progress.\n\nThis command does not support authenticating via fine grained PATs\nas it is not currently possible to create a PAT with the `checks:read` permission."),
				ox.Usage("watch", "Watch a run until it completes, showing its progress"),
				ox.Spec("<run-id> [flags]"),
				ox.Example("\n  # Watch a run until it's done\n  gh run watch\n  \n  # Run some other command when the run is finished\n  gh run watch && notify-send 'run is done!'"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("exit-status", "Exit with non-zero status if run fails").
					Int("interval", "Refresh interval in seconds", ox.Default("3"), ox.Short("i")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("List, view, and run workflows in GitHub Actions."),
			ox.Usage("workflow", "View details about GitHub Actions workflows"),
			ox.Spec("<command> [flags]"),
			ox.Section(1),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Disable a workflow, preventing it from running or showing up when listing workflows."),
				ox.Usage("disable", "Disable a workflow"),
				ox.Spec("[<workflow-id> | <workflow-name>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Enable a workflow, allowing it to be run and show up when listing workflows."),
				ox.Usage("enable", "Enable a workflow"),
				ox.Spec("[<workflow-id> | <workflow-name>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List workflow files, hiding disabled workflows by default.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List workflows"),
				ox.Spec("[flags]"),
				ox.Aliases("gh workflow ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Include disabled workflows", ox.Short("a")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of workflows to fetch", ox.Default("50"), ox.Short("L")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Create a `workflow_dispatch` event for a given workflow.\n\nThis command will trigger GitHub Actions to run a given workflow file. The given workflow file must\nsupport an `on.workflow_dispatch` trigger in order to be run in this way.\n\nIf the workflow file supports inputs, they can be specified in a few ways:\n\n- Interactively\n- Via `-f/--raw-field` or `-F/--field` flags\n- As JSON, via standard input"),
				ox.Usage("run", "Run a workflow by creating a workflow_dispatch event"),
				ox.Spec("[<workflow-id> | <workflow-name>] [flags]"),
				ox.Example("\n  # Have gh prompt you for what workflow you'd like to run and interactively collect inputs\n  $ gh workflow run\n  \n  # Run the workflow file 'triage.yml' at the remote's default branch\n  $ gh workflow run triage.yml\n  \n  # Run the workflow file 'triage.yml' at a specified ref\n  $ gh workflow run triage.yml --ref my-branch\n  \n  # Run the workflow file 'triage.yml' with command line inputs\n  $ gh workflow run triage.yml -f name=scully -f greeting=hello\n  \n  # Run the workflow file 'triage.yml' with JSON via standard input\n  $ echo '{\"name\":\"scully\", \"greeting\":\"hello\"}' | gh workflow run triage.yml --json"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Map("field", "Add a string parameter in key=value format, respecting @ syntax (see \"gh help api\").", ox.Spec("key=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Short("F")).
					Bool("json", "Read workflow inputs as JSON via STDIN").
					Map("raw-field", "Add a string parameter in key=value format", ox.Spec("key=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Short("f")).
					String("ref", "The branch or tag name which contains the version of the workflow file you'd like to run", ox.Short("r")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("View the summary of a workflow"),
				ox.Usage("view", "View the summary of a workflow"),
				ox.Spec("[<workflow-id> | <workflow-name> | <filename>] [flags]"),
				ox.Example("\n  # Interactively select a workflow to view\n  $ gh workflow view\n  \n  # View a specific workflow\n  $ gh workflow view 0451"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("ref", "The branch or tag name which contains the version of the workflow file you'd like to view", ox.Short("r")).
					Bool("web", "Open workflow in the browser", ox.Short("w")).
					Bool("yaml", "View the workflow yaml file", ox.Short("y")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Check out a pull request in git"),
			ox.Usage("co", "Alias for \"pr checkout\""),
			ox.Spec("gh pr checkout {<number> | <url> | <branch>} [flags]"),
			ox.Section(2),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Flags().
				String("branch", "Local branch name to use", ox.Default("[the name of the head branch]"), ox.Short("b")).
				Bool("detach", "Checkout PR with a detached HEAD").
				Bool("force", "Reset the existing local branch to the latest state of the pull request", ox.Short("f")).
				Bool("recurse-submodules", "Update all submodules after checkout").
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Aliases can be used to make shortcuts for gh commands or to compose multiple commands.\n\nRun `gh help alias set` to learn more."),
			ox.Usage("alias", "Create command shortcuts"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Delete set aliases"),
				ox.Usage("delete", "Delete set aliases"),
				ox.Spec("{<alias> | --all} [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Delete all aliases"),
			), ox.Sub(
				ox.Banner("Import aliases from the contents of a YAML file.\n\nAliases should be defined as a map in YAML, where the keys represent aliases and\nthe values represent the corresponding expansions. An example file should look like\nthe following:\n\n    bugs: issue list --label=bug\n    igrep: '!gh issue list --label=\"$1\" | grep \"$2\"'\n    features: |-\n        issue list\n        --label=enhancement\n\nUse `-` to read aliases (in YAML format) from standard input.\n\nThe output from `gh alias list` can be used to produce a YAML file\ncontaining your aliases, which you can use to import them from one machine to\nanother. Run `gh help alias list` to learn more."),
				ox.Usage("import", "Import aliases from a YAML file"),
				ox.Spec("[<filename> | -] [flags]"),
				ox.Example("\n  # Import aliases from a file\n  $ gh alias import aliases.yml\n  \n  # Import aliases from standard input\n  $ gh alias import -"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("clobber", "Overwrite existing aliases of the same name"),
			), ox.Sub(
				ox.Banner("This command prints out all of the aliases gh is configured to use."),
				ox.Usage("list", "List your aliases"),
				ox.Spec("[flags]"),
				ox.Aliases("gh alias ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Define a word that will expand to a full gh command when invoked.\n\nThe expansion may specify additional arguments and flags. If the expansion includes\npositional placeholders such as `$1`, extra arguments that follow the alias will be\ninserted appropriately. Otherwise, extra arguments will be appended to the expanded\ncommand.\n\nUse `-` as expansion argument to read the expansion string from standard input. This\nis useful to avoid quoting issues when defining expansions.\n\nIf the expansion starts with `!` or if `--shell` was given, the expansion is a shell\nexpression that will be evaluated through the `sh` interpreter when the alias is\ninvoked. This allows for chaining multiple commands via piping and redirection."),
				ox.Usage("set", "Create a shortcut for a gh command"),
				ox.Spec("<alias> <expansion> [flags]"),
				ox.Example("\n  # note: Command Prompt on Windows requires using double quotes for arguments\n  $ gh alias set pv 'pr view'\n  $ gh pv -w 123  #=> gh pr view -w 123\n  \n  $ gh alias set bugs 'issue list --label=bugs'\n  $ gh bugs\n  \n  $ gh alias set homework 'issue list --assignee @me'\n  $ gh homework\n  \n  $ gh alias set 'issue mine' 'issue list --mention @me'\n  $ gh issue mine\n  \n  $ gh alias set epicsBy 'issue list --author=\"$1\" --label=\"epic\"'\n  $ gh epicsBy vilmibm  #=> gh issue list --author=\"vilmibm\" --label=\"epic\"\n  \n  $ gh alias set --shell igrep 'gh issue list --label=\"$1\" | grep \"$2\"'\n  $ gh igrep epic foo  #=> gh issue list --label=\"epic\" | grep \"foo\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("clobber", "Overwrite existing aliases of the same name").
					Bool("shell", "Declare an alias to be passed through a shell interpreter", ox.Short("s")),
			)), ox.Sub(
			ox.Banner("Makes an authenticated HTTP request to the GitHub API and prints the response.\n\nThe endpoint argument should either be a path of a GitHub API v3 endpoint, or\n`graphql` to access the GitHub API v4.\n\nPlaceholder values `{owner}`, `{repo}`, and `{branch}` in the endpoint\nargument will get replaced with values from the repository of the current\ndirectory or the repository specified in the `GH_REPO` environment variable.\nNote that in some shells, for example PowerShell, you may need to enclose\nany value that contains `{...}` in quotes to prevent the shell from\napplying special meaning to curly braces.\n\nThe default HTTP request method is `GET` normally and `POST` if any parameters\nwere added. Override the method with `--method`.\n\nPass one or more `-f/--raw-field` values in `key=value` format to add static string\nparameters to the request payload. To add non-string or placeholder-determined values, see\n`-F/--field` below. Note that adding request parameters will automatically switch the\nrequest method to `POST`. To send the parameters as a `GET` query string instead, use\n`--method GET`.\n\nThe `-F/--field` flag has magic type conversion based on the format of the value:\n\n- literal values `true`, `false`, `null`, and integer numbers get converted to\n  appropriate JSON types;\n- placeholder values `{owner}`, `{repo}`, and `{branch}` get populated with values\n  from the repository of the current directory;\n- if the value starts with `@`, the rest of the value is interpreted as a\n  filename to read the value from. Pass `-` to read from standard input.\n\nFor GraphQL requests, all fields other than `query` and `operationName` are\ninterpreted as GraphQL variables.\n\nTo pass nested parameters in the request payload, use `key[subkey]=value` syntax when\ndeclaring fields. To pass nested values as arrays, declare multiple fields with the\nsyntax `key[]=value1`, `key[]=value2`. To pass an empty array, use `key[]` without a\nvalue.\n\nTo pass pre-constructed JSON or payloads in other formats, a request body may be read\nfrom file specified by `--input`. Use `-` to read from standard input. When passing the\nrequest body this way, any parameters specified via field flags are added to the query\nstring of the endpoint URL.\n\nIn `--paginate` mode, all pages of results will sequentially be requested until\nthere are no more pages of results. For GraphQL requests, this requires that the\noriginal query accepts an `$endCursor: String` variable and that it fetches the\n`pageInfo{ hasNextPage, endCursor }` set of fields from a collection. Each page is a separate\nJSON array or object. Pass `--slurp` to wrap all pages of JSON arrays or objects\ninto an outer JSON array.\n\nFor more information about output formatting flags, see `gh help formatting`."),
			ox.Usage("api", "Make an authenticated GitHub API request"),
			ox.Spec("<endpoint> [flags]"),
			ox.Example("\n  # list releases in the current repository\n  $ gh api repos/{owner}/{repo}/releases\n  \n  # post an issue comment\n  $ gh api repos/{owner}/{repo}/issues/123/comments -f body='Hi from CLI'\n  \n  # post nested parameter read from a file\n  $ gh api gists -F 'files[myfile.txt][content]=@myfile.txt'\n  \n  # add parameters to a GET request\n  $ gh api -X GET search/issues -f q='repo:cli/cli is:open remote'\n  \n  # set a custom HTTP header\n  $ gh api -H 'Accept: application/vnd.github.v3.raw+json' ...\n  \n  # opt into GitHub API previews\n  $ gh api --preview baptiste,nebula ...\n  \n  # print only specific fields from the response\n  $ gh api repos/{owner}/{repo}/issues --jq '.[].title'\n  \n  # use a template for the output\n  $ gh api repos/{owner}/{repo}/issues --template \\\n    '{{range .}}{{.title}} ({{.labels | pluck \"name\" | join \", \" | color \"yellow\"}}){{\"\\n\"}}{{end}}'\n  \n  # update allowed values of the \"environment\" custom property in a deeply nested array\n  gh api -X PATCH /orgs/{org}/properties/schema \\\n     -F 'properties[][property_name]=environment' \\\n     -F 'properties[][default_value]=production' \\\n     -F 'properties[][allowed_values][]=staging' \\\n     -F 'properties[][allowed_values][]=production'\n  \n  # list releases with GraphQL\n  $ gh api graphql -F owner='{owner}' -F name='{repo}' -f query='\n    query($name: String!, $owner: String!) {\n      repository(owner: $owner, name: $name) {\n        releases(last: 3) {\n          nodes { tagName }\n        }\n      }\n    }\n  '\n  \n  # list all repositories for a user\n  $ gh api graphql --paginate -f query='\n    query($endCursor: String) {\n      viewer {\n        repositories(first: 100, after: $endCursor) {\n          nodes { nameWithOwner }\n          pageInfo {\n            hasNextPage\n            endCursor\n          }\n        }\n      }\n    }\n  '\n  \n  # get the percentage of forks for the current user\n  $ gh api graphql --paginate --slurp -f query='\n    query($endCursor: String) {\n      viewer {\n        repositories(first: 100, after: $endCursor) {\n          nodes { isFork }\n          pageInfo {\n            hasNextPage\n            endCursor\n          }\n        }\n      }\n    }\n  ' | jq 'def count(e): reduce e as $_ (0;.+1);\n  [.[].data.viewer.repositories.nodes[]] as $r | count(select($r[].isFork))/count($r[])'"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Flags().
				Duration("cache", "Cache the response, e.g. \"3600s\", \"60m\", \"1h\"").
				Map("field", "Add a typed parameter in key=value format", ox.Spec("key=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Short("F")).
				String("header", "Add a HTTP request header in key:value format", ox.Spec("key:value"), ox.Short("H")).
				String("hostname", "The GitHub hostname for the request", ox.Default("github.com")).
				Bool("include", "Include HTTP response status line and headers in the output", ox.Short("i")).
				String("input", "The file to use as body for the HTTP request (use \"-\" to read from standard input)", ox.Spec("file")).
				String("jq", "Query to select values from the response using jq syntax", ox.Short("q")).
				String("method", "The HTTP method for the request", ox.Default("GET"), ox.Short("X")).
				Bool("paginate", "Make additional HTTP requests to fetch all pages of results").
				Slice("preview", "GitHub API preview names to request (without the \"-preview\" suffix)", ox.Elem(ox.StringT), ox.Short("p")).
				Map("raw-field", "Add a string parameter in key=value format", ox.Spec("key=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Short("f")).
				Bool("silent", "Do not print the response body").
				Bool("slurp", "Use with \"--paginate\" to return an array of all pages of either JSON arrays or objects").
				String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
				Bool("verbose", "Include full HTTP request and response in the output"),
		), ox.Sub(
			ox.Banner("Download and verify artifact attestations."),
			ox.Usage("attestation", "Work with artifact attestations"),
			ox.Spec("[subcommand] [flags]"),
			ox.Aliases("gh at"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("### NOTE: This feature is currently in public preview, and subject to change.\n\nDownload attestations associated with an artifact for offline use.\n\nThe command requires either:\n* a file path to an artifact, or\n* a container image URI (e.g. `oci://<image-uri>`)\n  * (note that if you provide an OCI URL, you must already be authenticated with\nits container registry)\n\nIn addition, the command requires either:\n* the `--repo` flag (e.g. --repo github/example).\n* the `--owner` flag (e.g. --owner github), or\n\nThe `--repo` flag value must match the name of the GitHub repository\nthat the artifact is linked with.\n\nThe `--owner` flag value must match the name of the GitHub organization\nthat the artifact's linked repository belongs to.\n\nAny associated bundle(s) will be written to a file in the\ncurrent directory named after the artifact's digest. For example, if the\ndigest is \"sha256:1234\", the file will be named \"sha256:1234.jsonl\"."),
				ox.Usage("download", "Download an artifact's attestations for offline use"),
				ox.Spec("[<file-path> | oci://<image-uri>] [--owner | --repo] [flags]"),
				ox.Example("\n  # Download attestations for a local artifact linked with an organization\n  $ gh attestation download example.bin -o github\n  \n  # Download attestations for a local artifact linked with a repository\n  $ gh attestation download example.bin -R github/example\n  \n  # Download attestations for an OCI image linked with an organization\n  $ gh attestation download oci://example.com/foo/bar:latest -o github"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("digest-alg", "The algorithm used to compute a digest of the artifact: {sha256|sha512}", ox.Default("sha256"), ox.Short("d")).
					String("hostname", "Configure host to use").
					Int("limit", "Maximum number of attestations to fetch", ox.Default("30"), ox.Short("L")).
					String("owner", "a GitHub organization to scope attestation lookup by", ox.Short("o")).
					String("predicate-type", "Filter attestations by provided predicate type").
					String("repo", "Repository name in the format <owner>/<repo>", ox.Short("R")),
			), ox.Sub(
				ox.Banner("### NOTE: This feature is currently in public preview, and subject to change.\n\nOutput contents for a trusted_root.jsonl file, likely for offline verification.\n\nWhen using `gh attestation verify`, if your machine is on the internet,\nthis will happen automatically. But to do offline verification, you need to\nsupply a trusted root file with `--custom-trusted-root`; this command\nwill help you fetch a `trusted_root.jsonl` file for that purpose.\n\nYou can call this command without any flags to get a trusted root file covering\nthe Sigstore Public Good Instance as well as GitHub's Sigstore instance.\n\nOtherwise you can use `--tuf-url` to specify the URL of a custom TUF\nrepository mirror, and `--tuf-root` should be the path to the\n`root.json` file that you securely obtained out-of-band.\n\nIf you just want to verify the integrity of your local TUF repository, and don't\nwant the contents of a trusted_root.jsonl file, use `--verify-only`."),
				ox.Usage("trusted-root", "Output trusted_root.jsonl contents, likely for offline verification"),
				ox.Spec("[--tuf-url <url> --tuf-root <file-path>] [--verify-only] [flags]"),
				ox.Example("\n  # Get a trusted_root.jsonl for both Sigstore Public Good and GitHub's instance\n  gh attestation trusted-root"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("hostname", "Configure host to use").
					String("tuf-root", "Path to the TUF root.json file on disk").
					String("tuf-url", "URL to the TUF repository mirror").
					Bool("verify-only", "Don't output trusted_root.jsonl contents"),
			), ox.Sub(
				ox.Banner("Verify the integrity and provenance of an artifact using its associated\ncryptographically signed attestations.\n\nIn order to verify an attestation, you must validate the identity of the Actions\nworkflow that produced the attestation (a.k.a. the signer workflow). Given this\nidentity, the verification process checks the signatures in the attestations,\nand confirms that the attestation refers to provided artifact.\n\nTo specify the artifact, the command requires:\n* a file path to an artifact, or\n* a container image URI (e.g. `oci://<image-uri>`)\n  * (note that if you provide an OCI URL, you must already be authenticated with\nits container registry)\n\nTo fetch the attestation, and validate the identity of the signer, the command\nrequires either:\n* the `--repo` flag (e.g. --repo github/example).\n* the `--owner` flag (e.g. --owner github), or\n\nThe `--repo` flag value must match the name of the GitHub repository\nthat the artifact is linked with.\n\nThe `--owner` flag value must match the name of the GitHub organization\nthat the artifact's linked repository belongs to.\n\nBy default, the verify command will:\n- only verify provenance attestations\n- attempt to fetch relevant attestations via the GitHub API.\n\nTo verify other types of attestations, use the `--predicate-type` flag.\n\nTo use your artifact's OCI registry instead of GitHub's API, use the\n`--bundle-from-oci` flag. For offline verification, using attestations\nstored on desk (c.f. the download command), provide a path to the `--bundle` flag.\n\nTo see the full results that are generated upon successful verification, i.e.\nfor use with a policy engine, provide the `--format=json` flag.\n\nThe signer workflow's identity is validated against the Subject Alternative Name (SAN)\nwithin the attestation certificate. Often, the signer workflow is the\nsame workflow that started the run and generated the attestation, and will be\nlocated inside your repository. For this reason, by default this command uses\neither the `--repo` or the `--owner` flag value to validate the SAN.\n\nHowever, sometimes the caller workflow is not the same workflow that\nperformed the signing. If your attestation was generated via a reusable\nworkflow, then that reusable workflow is the signer whose identity needs to be\nvalidated. In this situation, the signer workflow may or may not be located\ninside your `--repo` or `--owner`.\n\nWhen using reusable workflows, use the `--signer-repo`, `--signer-workflow`,\nor `--cert-identity` flags to validate the signer workflow's identity.\n\nFor more policy verification options, see the other available flags.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("verify", "Verify an artifact's integrity using attestations"),
				ox.Spec("[<file-path> | oci://<image-uri>] [--owner | --repo] [flags]"),
				ox.Example("\n  # Verify an artifact linked with a repository\n  $ gh attestation verify example.bin --repo github/example\n  \n  # Verify an artifact linked with an organization\n  $ gh attestation verify example.bin --owner github\n  \n  # Verify an artifact and output the full verification result\n  $ gh attestation verify example.bin --owner github --format json\n  \n  # Verify an OCI image using attestations stored on disk\n  $ gh attestation verify oci://<image-uri> --owner github --bundle sha256:foo.jsonl\n  \n  # Verify an artifact signed with a reusable workflow\n  $ gh attestation verify example.bin --owner github --signer-repo actions/example"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("bundle", "Path to bundle on disk, either a single bundle in a JSON file or a JSON lines file with multiple bundles", ox.Short("b")).
					Bool("bundle-from-oci", "When verifying an OCI image, fetch the attestation bundle from the OCI registry instead of from GitHub").
					String("cert-identity", "Enforce that the certificate's subject alternative name matches the provided value exactly").
					String("cert-identity-regex", "Enforce that the certificate's subject alternative name matches the provided regex", ox.Short("i")).
					String("cert-oidc-issuer", "Issuer of the OIDC token", ox.Default("https://to$USER.actions.githubusercontent.com")).
					String("custom-trusted-root", "Path to a trusted_root.jsonl file; likely for offline verification").
					Bool("deny-self-hosted-runners", "Fail verification for attestations generated on self-hosted runners").
					String("digest-alg", "The algorithm used to compute a digest of the artifact: {sha256|sha512}", ox.Default("sha256"), ox.Short("d")).
					String("format", "Output format: {json}").
					String("hostname", "Configure host to use").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Int("limit", "Maximum number of attestations to fetch", ox.Default("30"), ox.Short("L")).
					Bool("no-public-good", "Do not verify attestations signed with Sigstore public good instance").
					String("owner", "GitHub organization to scope attestation lookup by", ox.Short("o")).
					String("predicate-type", "Filter attestations by provided predicate type", ox.Default("https://slsa.dev/provenance/v1")).
					String("repo", "Repository name in the format <owner>/<repo>", ox.Short("R")).
					String("signer-repo", "Repository of reusable workflow that signed attestation in the format <owner>/<repo>").
					String("signer-workflow", "Workflow that signed attestation in the format [host/]<owner>/<repo>/<path>/<to>/<workflow>").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")),
			)), ox.Sub(
			ox.Banner("Display or change configuration settings for gh.\n\nCurrent respected settings:\n- `git_protocol`: the protocol to use for git clone and push operations {https|ssh} (default https)\n- `editor`: the text editor program to use for authoring text\n- `prompt`: toggle interactive prompting in the terminal {enabled|disabled} (default enabled)\n- `prefer_editor_prompt`: toggle preference for editor-based interactive prompting in the terminal {enabled|disabled} (default disabled)\n- `pager`: the terminal pager program to send standard output to\n- `http_unix_socket`: the path to a Unix socket through which to make an HTTP connection\n- `browser`: the web browser to use for opening URLs"),
			ox.Usage("config", "Manage configuration for gh"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Clear the cli cache"),
				ox.Usage("clear-cache", "Clear the cli cache"),
				ox.Spec("[flags]"),
				ox.Example("\n  # Clear the cli cache\n  $ gh config clear-cache"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Print the value of a given configuration key"),
				ox.Usage("get", "Print the value of a given configuration key"),
				ox.Spec("<key> [flags]"),
				ox.Example("\n  $ gh config get git_protocol\n  https"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("host", "Get per-host setting", ox.Short("h")),
			), ox.Sub(
				ox.Banner("Print a list of configuration keys and values"),
				ox.Usage("list", "Print a list of configuration keys and values"),
				ox.Spec("[flags]"),
				ox.Aliases("gh config ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("host", "Get per-host configuration", ox.Short("h")),
			), ox.Sub(
				ox.Banner("Update configuration with a value for the given key"),
				ox.Usage("set", "Update configuration with a value for the given key"),
				ox.Spec("<key> <value> [flags]"),
				ox.Example("\n  $ gh config set editor vim\n  $ gh config set editor \"code --wait\"\n  $ gh config set git_protocol ssh --host github.com\n  $ gh config set prompt disabled"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("host", "Set per-host setting", ox.Short("h")),
			)), ox.Sub(
			ox.Banner("GitHub CLI extensions are repositories that provide additional gh commands.\n\nThe name of the extension repository must start with `gh-` and it must contain an\nexecutable of the same name. All arguments passed to the `gh <extname>` invocation\nwill be forwarded to the `gh-<extname>` executable of the extension.\n\nAn extension cannot override any of the core gh commands. If an extension name conflicts\nwith a core gh command, you can use `gh extension exec <extname>`.\n\nFor the list of available extensions, see <https://github.com/topics/gh-extension>."),
			ox.Usage("extension", "Manage gh extensions"),
			ox.Spec("[flags]"),
			ox.Aliases("s", "gh ext"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("This command will take over your terminal and run a fully interactive\ninterface for browsing, adding, and removing gh extensions. A terminal\nwidth greater than 100 columns is recommended.\n\nTo learn how to control this interface, press `?` after running to see\nthe help text.\n\nPress `q` to quit.\n\nRunning this command with `--single-column` should make this command\nmore intelligible for users who rely on assistive technology like screen\nreaders or high zoom.\n\nFor a more traditional way to discover extensions, see:\n\n\tgh ext search\n\nalong with `gh ext install`, `gh ext remove`, and `gh repo view`."),
				ox.Usage("browse", "Enter a UI for browsing, adding, and removing extensions"),
				ox.Spec("[flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("debug", "log to /tmp/extBrowse-*").
					Bool("single-column", "Render TUI with only one column of text", ox.Short("s")),
			), ox.Sub(
				ox.Banner("Create a new extension"),
				ox.Usage("create", "Create a new extension"),
				ox.Spec("[<name>] [flags]"),
				ox.Example("\n  # Use interactively\n  gh extension create\n  \n  # Create a script-based extension\n  gh extension create foobar\n  \n  # Create a Go extension\n  gh extension create --precompiled=go foobar\n  \n  # Create a non-Go precompiled extension\n  gh extension create --precompiled=other foobar"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("precompiled", "Create a precompiled extension. Possible values: go, other"),
			), ox.Sub(
				ox.Banner("Execute an extension using the short name. For example, if the extension repository is\n`owner/gh-extension`, you should pass `extension`. You can use this command when\nthe short name conflicts with a core gh command.\n\nAll arguments after the extension name will be forwarded to the executable\nof the extension."),
				ox.Usage("exec", "Execute an installed extension"),
				ox.Spec("<name> [args] [flags]"),
				ox.Example("\n  # execute a label extension instead of the core gh label command\n  $ gh extension exec label"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Install a GitHub repository locally as a GitHub CLI extension.\n\nThe repository argument can be specified in `OWNER/REPO` format as well as a full URL.\nThe URL format is useful when the repository is not hosted on github.com.\n\nTo install an extension in development from the current directory, use `.` as the\nvalue of the repository argument.\n\nFor the list of available extensions, see <https://github.com/topics/gh-extension>."),
				ox.Usage("install", "Install a gh extension from a repository"),
				ox.Spec("<repository> [flags]"),
				ox.Example("\n  $ gh extension install owner/gh-extension\n  $ gh extension install https://git.example.com/owner/gh-extension\n  $ gh extension install ."),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("force", "force upgrade extension, or ignore if latest already installed").
					String("pin", "pin extension to a release tag or commit ref"),
			), ox.Sub(
				ox.Banner("List installed extension commands"),
				ox.Usage("list", "List installed extension commands"),
				ox.Spec("[flags]"),
				ox.Aliases("gh ext ls", "gh extension ls", "gh extensions ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Remove an installed extension"),
				ox.Usage("remove", "Remove an installed extension"),
				ox.Spec("<name> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			), ox.Sub(
				ox.Banner("Search for gh extensions.\n\nWith no arguments, this command prints out the first 30 extensions\navailable to install sorted by number of stars. More extensions can\nbe fetched by specifying a higher limit with the `--limit` flag.\n\nWhen connected to a terminal, this command prints out three columns.\nThe first has a ✓ if the extension is already installed locally. The\nsecond is the full name of the extension repository in `OWNER/REPO`\nformat. The third is the extension's description.\n\nWhen not connected to a terminal, the ✓ character is rendered as the\nword \"installed\" but otherwise the order and content of the columns\nare the same.\n\nThis command behaves similarly to `gh search repos` but does not\nsupport as many search qualifiers. For a finer grained search of\nextensions, try using:\n\n\tgh search repos --topic \"gh-extension\"\n\nand adding qualifiers as needed. See `gh help search repos` to learn\nmore about repository search.\n\nFor listing just the extensions that are already installed locally,\nsee:\n\n\tgh ext list\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("search", "Search extensions to the GitHub CLI"),
				ox.Spec("[<query>] [flags]"),
				ox.Example("\n  # List the first 30 extensions sorted by star count, descending\n  $ gh ext search\n  \n  # List more extensions\n  $ gh ext search --limit 300\n  \n  # List extensions matching the term \"branch\"\n  $ gh ext search branch\n  \n  # List extensions owned by organization \"github\"\n  $ gh ext search --owner github\n  \n  # List extensions, sorting by recently updated, ascending\n  $ gh ext search --sort updated --order asc\n  \n  # List extensions, filtering by license\n  $ gh ext search --license MIT\n  \n  # Open search results in the browser\n  $ gh ext search -w"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Slice("license", "Filter based on license type", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of extensions to fetch", ox.Default("30"), ox.Short("L")).
					String("order", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}", ox.Default("desc")).
					Slice("owner", "Filter on owner", ox.Elem(ox.StringT)).
					String("sort", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}", ox.Default("best-match")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Upgrade installed extensions"),
				ox.Usage("upgrade", "Upgrade installed extensions"),
				ox.Spec("{<name> | --all} [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("all", "Upgrade all extensions").
					Bool("dry-run", "Only display upgrades").
					Bool("force", "Force upgrade extension"),
			)), ox.Sub(
			ox.Banner("Manage GPG keys registered with your GitHub account."),
			ox.Usage("gpg-key", "Manage GPG keys"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Add a GPG key to your GitHub account"),
				ox.Usage("add", "Add a GPG key to your GitHub account"),
				ox.Spec("[<key-file>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("title", "Title for the new key", ox.Short("t")),
			), ox.Sub(
				ox.Banner("Delete a GPG key from your GitHub account"),
				ox.Usage("delete", "Delete a GPG key from your GitHub account"),
				ox.Spec("<key-id> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")),
			), ox.Sub(
				ox.Banner("Lists GPG keys in your GitHub account"),
				ox.Usage("list", "Lists GPG keys in your GitHub account"),
				ox.Spec("[flags]"),
				ox.Aliases("gh gpg-key ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			)), ox.Sub(
			ox.Banner("Work with GitHub labels."),
			ox.Usage("label", "Manage labels"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Clones labels from a source repository to a destination repository on GitHub.\nBy default, the destination repository is the current repository.\n\nAll labels from the source repository will be copied to the destination\nrepository. Labels in the destination repository that are not in the source\nrepository will not be deleted or modified.\n\nLabels from the source repository that already exist in the destination\nrepository will be skipped. You can overwrite existing labels in the\ndestination repository using the `--force` flag."),
				ox.Usage("clone", "Clones labels from one repository to another"),
				ox.Spec("<source-repository> [flags]"),
				ox.Example("\n  # clone and overwrite labels from cli/cli repository into the current repository\n  $ gh label clone cli/cli --force\n  \n  # clone labels from cli/cli repository into a octocat/cli repository\n  $ gh label clone cli/cli --repo octocat/cli"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("force", "Overwrite labels in the destination repository", ox.Short("f")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Create a new label on GitHub, or update an existing one with `--force`.\n\nMust specify name for the label. The description and color are optional.\nIf a color isn't provided, a random one will be chosen.\n\nThe label color needs to be 6 character hex value."),
				ox.Usage("create", "Create a new label"),
				ox.Spec("<name> [flags]"),
				ox.Example("\n  # create new bug label\n  $ gh label create bug --description \"Something isn't working\" --color E99695"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("color", "Color of the label", ox.Short("c")).
					String("description", "Description of the label", ox.Short("d")).
					Bool("force", "Update the label color and description if label already exists", ox.Short("f")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Delete a label from a repository"),
				ox.Usage("delete", "Delete a label from a repository"),
				ox.Spec("<name> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Confirm deletion without prompting").
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Update a label on GitHub.\n\nA label can be renamed using the `--name` flag.\n\nThe label color needs to be 6 character hex value."),
				ox.Usage("edit", "Edit a label"),
				ox.Spec("<name> [flags]"),
				ox.Example("\n  # update the color of the bug label\n  $ gh label edit bug --color FF0000\n  \n  # rename and edit the description of the bug label\n  $ gh label edit bug --name big-bug --description \"Bigger than normal bug\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("color", "Color of the label", ox.Short("c")).
					String("description", "Description of the label", ox.Short("d")).
					String("name", "New name of the label", ox.Short("n")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Display labels in a GitHub repository.\n\nWhen using the `--search` flag results are sorted by best match of the query.\nThis behavior cannot be configured with the `--order` or `--sort` flags.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List labels in a repository"),
				ox.Spec("[flags]"),
				ox.Aliases("gh label ls"),
				ox.Example("\n  # sort labels by name\n  $ gh label list --sort name\n  \n  # find labels with \"bug\" in the name or description\n  $ gh label list --search bug"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of labels to fetch", ox.Default("30"), ox.Short("L")).
					String("order", "Order of labels returned: {asc|desc}", ox.Default("asc")).
					String("search", "Search label names and descriptions", ox.Short("S")).
					String("sort", "Sort fetched labels: {created|name}", ox.Default("created")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "List labels in the web browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Repository rulesets are a way to define a set of rules that apply to a repository.\nThese commands allow you to view information about them."),
			ox.Usage("ruleset", "View info about repo rulesets"),
			ox.Spec("<command> [flags]"),
			ox.Aliases("gh rs"),
			ox.Example("\n  $ gh ruleset list\n  $ gh ruleset view --repo OWNER/REPO --web\n  $ gh ruleset check branch-name"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("View information about GitHub rules that apply to a given branch.\n\nThe provided branch name does not need to exist; rules will be displayed that would apply\nto a branch with that name. All rules are returned regardless of where they are configured.\n\nIf no branch name is provided, then the current branch will be used.\n\nThe `--default` flag can be used to view rules that apply to the default branch of the\nrepository."),
				ox.Usage("check", "View rules that would apply to a given branch"),
				ox.Spec("[<branch>] [flags]"),
				ox.Example("\n  # View all rules that apply to the current branch\n  $ gh ruleset check\n  \n  # View all rules that apply to a branch named \"my-branch\" in a different repository\n  $ gh ruleset check my-branch --repo owner/repo\n  \n  # View all rules that apply to the default branch in a different repository\n  $ gh ruleset check --default --repo owner/repo\n  \n  # View a ruleset configured in a different repository or any of its parents\n  $ gh ruleset view 23 --repo owner/repo\n  \n  # View an organization-level ruleset\n  $ gh ruleset view 23 --org my-org"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("default", "Check rules on default branch").
					Bool("web", "Open the branch rules page in a web browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List GitHub rulesets for a repository or organization.\n\nIf no options are provided, the current repository's rulesets are listed. You can query a different\nrepository's rulesets by using the `--repo` flag. You can also use the `--org` flag to list rulesets\nconfigured for the provided organization.\n\nUse the `--parents` flag to control whether rulesets configured at higher levels that also apply to the provided\nrepository or organization should be returned. The default is `true`.\n\nYour access token must have the `admin:org` scope to use the `--org` flag, which can be granted by running `gh auth refresh -s admin:org`."),
				ox.Usage("list", "List rulesets for a repository or organization"),
				ox.Spec("[flags]"),
				ox.Aliases("gh rs ls", "gh ruleset ls"),
				ox.Example("\n  # List rulesets in the current repository\n  $ gh ruleset list\n  \n  # List rulesets in a different repository, including those configured at higher levels\n  $ gh ruleset list --repo owner/repo --parents\n  \n  # List rulesets in an organization\n  $ gh ruleset list --org org-name"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Int("limit", "Maximum number of rulesets to list", ox.Default("30"), ox.Short("L")).
					String("org", "List organization-wide rulesets for the provided organization", ox.Short("o")).
					Bool("parents", "Whether to include rulesets configured at higher levels that also apply", ox.Default("true"), ox.Short("p")).
					Bool("web", "Open the list of rulesets in the web browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("View information about a GitHub ruleset.\n\nIf no ID is provided, an interactive prompt will be used to choose\nthe ruleset to view.\n\nUse the `--parents` flag to control whether rulesets configured at higher\nlevels that also apply to the provided repository or organization should\nbe returned. The default is `true`."),
				ox.Usage("view", "View information about a ruleset"),
				ox.Spec("[<ruleset-id>] [flags]"),
				ox.Example("\n  # Interactively choose a ruleset to view from all rulesets that apply to the current repository\n  $ gh ruleset view\n  \n  # Interactively choose a ruleset to view from only rulesets configured in the current repository\n  $ gh ruleset view --no-parents\n  \n  # View a ruleset configured in the current repository or any of its parents\n  $ gh ruleset view 43\n  \n  # View a ruleset configured in a different repository or any of its parents\n  $ gh ruleset view 23 --repo owner/repo\n  \n  # View an organization-level ruleset\n  $ gh ruleset view 23 --org my-org"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("org", "Organization name if the provided ID is an organization-level ruleset", ox.Short("o")).
					Bool("parents", "Whether to include rulesets configured at higher levels that also apply", ox.Default("true"), ox.Short("p")).
					Bool("web", "Open the ruleset in the browser", ox.Short("w")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Search across all of GitHub."),
			ox.Usage("search", "Search for repositories, issues, and pull requests"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Search within code in GitHub repositories.\n\nThe search syntax is documented at:\n<https://docs.github.com/search-github/searching-on-github/searching-code>\n\nNote that these search results are powered by what is now a legacy GitHub code search engine.\nThe results might not match what is seen on github.com, and new features like regex search\nare not yet available via the GitHub API.\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("code", "Search within code"),
				ox.Spec("<query> [flags]"),
				ox.Example("\n  # search code matching \"react\" and \"lifecycle\"\n  $ gh search code react lifecycle\n  \n  # search code matching \"error handling\" \n  $ gh search code \"error handling\"\n  \t\n  # search code matching \"deque\" in Python files\n  $ gh search code deque --language=python\n  \n  # search code matching \"cli\" in repositories owned by microsoft organization\n  $ gh search code cli --owner=microsoft\n  \n  # search code matching \"panic\" in the GitHub CLI repository\n  $ gh search code panic --repo cli/cli\n  \n  # search code matching keyword \"lint\" in package.json files\n  $ gh search code lint --filename package.json"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("extension", "Filter on file extension").
					String("filename", "Filter on filename").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("language", "Filter results by language").
					Int("limit", "Maximum number of code results to fetch", ox.Default("30"), ox.Short("L")).
					Slice("match", "Restrict search to file contents or file path: {file|path}", ox.Elem(ox.StringT)).
					Slice("owner", "Filter on owner", ox.Elem(ox.StringT)).
					Slice("repo", "Filter on repository", ox.Elem(ox.StringT), ox.Short("R")).
					String("size", "Filter on size range, in kilobytes").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Search for commits on GitHub.\n\nThe command supports constructing queries using the GitHub search syntax,\nusing the parameter and qualifier flags, or a combination of the two.\n\nGitHub search syntax is documented at:\n<https://docs.github.com/search-github/searching-on-github/searching-commits>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("commits", "Search for commits"),
				ox.Spec("[<query>] [flags]"),
				ox.Example("\n  # search commits matching set of keywords \"readme\" and \"typo\"\n  $ gh search commits readme typo\n  \n  # search commits matching phrase \"bug fix\"\n  $ gh search commits \"bug fix\"\n  \n  # search commits committed by user \"monalisa\"\n  $ gh search commits --committer=monalisa\n  \n  # search commits authored by users with name \"Jane Doe\"\n  $ gh search commits --author-name=\"Jane Doe\"\n  \n  # search commits matching hash \"8dd03144ffdc6c0d486d6b705f9c7fba871ee7c3\"\n  $ gh search commits --hash=8dd03144ffdc6c0d486d6b705f9c7fba871ee7c3\n  \n  # search commits authored before February 1st, 2022\n  $ gh search commits --author-date=\"<2022-02-01\""),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("author", "Filter by author").
					Date("author-date", "Filter based on authored date").
					String("author-email", "Filter on author email").
					String("author-name", "Filter on author name").
					String("committer", "Filter by committer").
					Date("committer-date", "Filter based on committed date").
					String("committer-email", "Filter on committer email").
					String("committer-name", "Filter on committer name").
					String("hash", "Filter by commit hash").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of commits to fetch", ox.Default("30"), ox.Short("L")).
					Bool("merge", "Filter on merge commits").
					String("order", "Order of commits returned, ignored unless '--sort' flag is specified: {asc|desc}", ox.Default("desc")).
					Slice("owner", "Filter on repository owner", ox.Elem(ox.StringT)).
					String("parent", "Filter by parent hash").
					Slice("repo", "Filter on repository", ox.Elem(ox.StringT), ox.Short("R")).
					String("sort", "Sort fetched commits: {author-date|committer-date}", ox.Default("best-match")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("tree", "Filter by tree hash").
					Slice("visibility", "Filter based on repository visibility: {public|private|internal}", ox.Elem(ox.StringT)).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Search for issues on GitHub.\n\nThe command supports constructing queries using the GitHub search syntax,\nusing the parameter and qualifier flags, or a combination of the two.\n\nGitHub search syntax is documented at:\n<https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("issues", "Search for issues"),
				ox.Spec("[<query>] [flags]"),
				ox.Example("\n  # search issues matching set of keywords \"readme\" and \"typo\"\n  $ gh search issues readme typo\n  \n  # search issues matching phrase \"broken feature\"\n  $ gh search issues \"broken feature\"\n  \n  # search issues and pull requests in cli organization\n  $ gh search issues --include-prs --owner=cli\n  \n  # search open issues assigned to yourself\n  $ gh search issues --assignee=@me --state=open\n  \n  # search issues with numerous comments\n  $ gh search issues --comments=\">100\"\n  \n  # search issues without label \"bug\"\n  $ gh search issues -- -label:bug\n  \n  # search issues only from un-archived repositories (default is all repositories)\n  $ gh search issues --owner github --archived=false"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Filter by GitHub App author").
					Bool("archived", "Filter based on the repository archived state {true|false}").
					String("assignee", "Filter by assignee").
					String("author", "Filter by author").
					Date("closed", "Filter on closed at date").
					String("commenter", "Filter based on comments by user", ox.Spec("user")).
					Int("comments", "Filter on number of comments", ox.Spec("number")).
					Date("created", "Filter based on created at date").
					Bool("include-prs", "Include pull requests in results").
					Int("interactions", "Filter on number of reactions and comments", ox.Spec("number")).
					String("involves", "Filter based on involvement of user", ox.Spec("user")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Slice("label", "Filter on label", ox.Elem(ox.StringT)).
					String("language", "Filter based on the coding language").
					Int("limit", "Maximum number of results to fetch", ox.Default("30"), ox.Short("L")).
					Bool("locked", "Filter on locked conversation status").
					Slice("match", "Restrict search to specific field of issue: {title|body|comments}", ox.Elem(ox.StringT)).
					String("mentions", "Filter based on user mentions", ox.Spec("user")).
					String("milestone", "Filter by milestone title", ox.Spec("title")).
					Bool("no-assignee", "Filter on missing assignee").
					Bool("no-label", "Filter on missing label").
					Bool("no-milestone", "Filter on missing milestone").
					Bool("no-project", "Filter on missing project").
					String("order", "Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}", ox.Default("desc")).
					Slice("owner", "Filter on repository owner", ox.Elem(ox.StringT)).
					String("project", "Filter on project board owner/number", ox.Spec("owner/number")).
					Int("reactions", "Filter on number of reactions", ox.Spec("number")).
					Slice("repo", "Filter on repository", ox.Elem(ox.StringT), ox.Short("R")).
					String("sort", "Sort fetched results: {comments|created|interactions|reactions|reactions-+1|reactions--1|reactions-heart|reactions-smile|reactions-tada|reactions-thinking_face|updated}", ox.Default("best-match")).
					String("state", "Filter based on state: {open|closed}").
					String("team-mentions", "Filter based on team mentions").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Date("updated", "Filter on last updated at date").
					Slice("visibility", "Filter based on repository visibility: {public|private|internal}", ox.Elem(ox.StringT)).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Search for pull requests on GitHub.\n\nThe command supports constructing queries using the GitHub search syntax,\nusing the parameter and qualifier flags, or a combination of the two.\n\nGitHub search syntax is documented at:\n<https://docs.github.com/search-github/searching-on-github/searching-issues-and-pull-requests>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("prs", "Search for pull requests"),
				ox.Spec("[<query>] [flags]"),
				ox.Example("\n  # search pull requests matching set of keywords \"fix\" and \"bug\"\n  $ gh search prs fix bug\n  \n  # search draft pull requests in cli repository\n  $ gh search prs --repo=cli/cli --draft\n  \n  # search open pull requests requesting your review\n  $ gh search prs --review-requested=@me --state=open\n  \n  # search merged pull requests assigned to yourself\n  $ gh search prs --assignee=@me --merged\n  \n  # search pull requests with numerous reactions\n  $ gh search prs --reactions=\">100\"\n  \n  # search pull requests without label \"bug\"\n  $ gh search prs -- -label:bug\n  \n  # search pull requests only from un-archived repositories (default is all repositories)\n  $ gh search prs --owner github --archived=false"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Filter by GitHub App author").
					Bool("archived", "Filter based on the repository archived state {true|false}").
					String("assignee", "Filter by assignee").
					String("author", "Filter by author").
					String("base", "Filter on base branch name", ox.Short("B")).
					String("checks", "Filter based on status of the checks: {pending|success|failure}").
					Date("closed", "Filter on closed at date").
					String("commenter", "Filter based on comments by user", ox.Spec("user")).
					Int("comments", "Filter on number of comments", ox.Spec("number")).
					Date("created", "Filter based on created at date").
					Bool("draft", "Filter based on draft state").
					String("head", "Filter on head branch name", ox.Short("H")).
					Int("interactions", "Filter on number of reactions and comments", ox.Spec("number")).
					String("involves", "Filter based on involvement of user", ox.Spec("user")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					Slice("label", "Filter on label", ox.Elem(ox.StringT)).
					String("language", "Filter based on the coding language").
					Int("limit", "Maximum number of results to fetch", ox.Default("30"), ox.Short("L")).
					Bool("locked", "Filter on locked conversation status").
					Slice("match", "Restrict search to specific field of issue: {title|body|comments}", ox.Elem(ox.StringT)).
					String("mentions", "Filter based on user mentions", ox.Spec("user")).
					Bool("merged", "Filter based on merged state").
					Date("merged-at", "Filter on merged at date").
					String("milestone", "Filter by milestone title", ox.Spec("title")).
					Bool("no-assignee", "Filter on missing assignee").
					Bool("no-label", "Filter on missing label").
					Bool("no-milestone", "Filter on missing milestone").
					Bool("no-project", "Filter on missing project").
					String("order", "Order of results returned, ignored unless '--sort' flag is specified: {asc|desc}", ox.Default("desc")).
					Slice("owner", "Filter on repository owner", ox.Elem(ox.StringT)).
					String("project", "Filter on project board owner/number", ox.Spec("owner/number")).
					Int("reactions", "Filter on number of reactions", ox.Spec("number")).
					Slice("repo", "Filter on repository", ox.Elem(ox.StringT), ox.Short("R")).
					String("review", "Filter based on review status: {none|required|approved|changes_requested}").
					String("review-requested", "Filter on user or team requested to review", ox.Spec("user")).
					String("reviewed-by", "Filter on user who reviewed", ox.Spec("user")).
					String("sort", "Sort fetched results: {comments|reactions|reactions-+1|reactions--1|reactions-smile|reactions-thinking_face|reactions-heart|reactions-tada|interactions|created|updated}", ox.Default("best-match")).
					String("state", "Filter based on state: {open|closed}").
					String("team-mentions", "Filter based on team mentions").
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Date("updated", "Filter on last updated at date").
					Slice("visibility", "Filter based on repository visibility: {public|private|internal}", ox.Elem(ox.StringT)).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Search for repositories on GitHub.\n\nThe command supports constructing queries using the GitHub search syntax,\nusing the parameter and qualifier flags, or a combination of the two.\n\nGitHub search syntax is documented at:\n<https://docs.github.com/search-github/searching-on-github/searching-for-repositories>\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("repos", "Search for repositories"),
				ox.Spec("[<query>] [flags]"),
				ox.Example("\n  # search repositories matching set of keywords \"cli\" and \"shell\"\n  $ gh search repos cli shell\n  \n  # search repositories matching phrase \"vim plugin\"\n  $ gh search repos \"vim plugin\"\n  \n  # search repositories public repos in the microsoft organization\n  $ gh search repos --owner=microsoft --visibility=public\n  \n  # search repositories with a set of topics\n  $ gh search repos --topic=unix,terminal\n  \n  # search repositories by coding language and number of good first issues\n  $ gh search repos --language=go --good-first-issues=\">=10\"\n  \n  # search repositories without topic \"linux\"\n  $ gh search repos -- -topic:linux\n  \n  # search repositories excluding archived repositories\n  $ gh search repos --archived=false"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("archived", "Filter based on the repository archived state {true|false}").
					Date("created", "Filter based on created at date").
					Int("followers", "Filter based on number of followers", ox.Spec("number")).
					Int("forks", "Filter on number of forks", ox.Spec("number")).
					Int("good-first-issues", "Filter on number of issues with the 'good first issue' label", ox.Spec("number")).
					Int("help-wanted-issues", "Filter on number of issues with the 'help wanted' label", ox.Spec("number")).
					String("include-forks", "Include forks in fetched repositories: {false|true|only}").
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("language", "Filter based on the coding language").
					Slice("license", "Filter based on license type", ox.Elem(ox.StringT)).
					Int("limit", "Maximum number of repositories to fetch", ox.Default("30"), ox.Short("L")).
					Slice("match", "Restrict search to specific field of repository: {name|description|readme}", ox.Elem(ox.StringT)).
					Int("number-topics", "Filter on number of topics", ox.Spec("number")).
					String("order", "Order of repositories returned, ignored unless '--sort' flag is specified: {asc|desc}", ox.Default("desc")).
					Slice("owner", "Filter on owner", ox.Elem(ox.StringT)).
					String("size", "Filter on a size range, in kilobytes").
					String("sort", "Sort fetched repositories: {forks|help-wanted-issues|stars|updated}", ox.Default("best-match")).
					Int("stars", "Filter on number of stars", ox.Spec("number")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Slice("topic", "Filter on topic", ox.Elem(ox.StringT)).
					Date("updated", "Filter on last updated at date").
					Slice("visibility", "Filter based on visibility: {public|private|internal}", ox.Elem(ox.StringT)).
					Bool("web", "Open the search query in the web browser", ox.Short("w")),
			)), ox.Sub(
			ox.Banner("Secrets can be set at the repository, or organization level for use in\nGitHub Actions or Dependabot. User, organization, and repository secrets can be set for\nuse in GitHub Codespaces. Environment secrets can be set for use in\nGitHub Actions. Run `gh help secret set` to learn how to get started."),
			ox.Usage("secret", "Manage GitHub secrets"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Delete a secret on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization\n- user: available to Codespaces for your user"),
				ox.Usage("delete", "Delete secrets"),
				ox.Spec("<secret-name> [flags]"),
				ox.Aliases("gh secret remove"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Delete a secret for a specific application: {actions|codespaces|dependabot}", ox.Short("a")).
					String("env", "Delete a secret for an environment", ox.Short("e")).
					String("org", "Delete a secret for an organization", ox.Short("o")).
					Bool("user", "Delete a secret for your user", ox.Short("u")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List secrets on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization\n- user: available to Codespaces for your user\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List secrets"),
				ox.Spec("[flags]"),
				ox.Aliases("gh secret ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "List secrets for a specific application: {actions|codespaces|dependabot}", ox.Short("a")).
					String("env", "List secrets for an environment", ox.Short("e")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("org", "List secrets for an organization", ox.Short("o")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					Bool("user", "List a secret for your user", ox.Short("u")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Set a value for a secret on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs, Dependabot, or Codespaces within an organization\n- user: available to Codespaces for your user\n\nOrganization and user secrets can optionally be restricted to only be available to\nspecific repositories.\n\nSecret values are locally encrypted before being sent to GitHub."),
				ox.Usage("set", "Create or update secrets"),
				ox.Spec("<secret-name> [flags]"),
				ox.Example("\n  # Paste secret value for the current repository in an interactive prompt\n  $ gh secret set MYSECRET\n  \n  # Read secret value from an environment variable\n  $ gh secret set MYSECRET --body \"$ENV_VALUE\"\n  \n  # Read secret value from a file\n  $ gh secret set MYSECRET < myfile.txt\n  \n  # Set secret for a deployment environment in the current repository\n  $ gh secret set MYSECRET --env myenvironment\n  \n  # Set organization-level secret visible to both public and private repositories\n  $ gh secret set MYSECRET --org myOrg --visibility all\n  \n  # Set organization-level secret visible to specific repositories\n  $ gh secret set MYSECRET --org myOrg --repos repo1,repo2,repo3\n  \n  # Set user-level secret for Codespaces\n  $ gh secret set MYSECRET --user\n  \n  # Set repository-level secret for Dependabot\n  $ gh secret set MYSECRET --app dependabot\n  \n  # Set multiple secrets imported from the \".env\" file\n  $ gh secret set -f .env\n  \n  # Set multiple secrets from stdin\n  $ gh secret set -f - < myfile.txt"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("app", "Set the application for a secret: {actions|codespaces|dependabot}", ox.Short("a")).
					String("body", "The value for the secret (reads from standard input if not specified)", ox.Short("b")).
					String("env", "Set deployment environment secret", ox.Spec("environment"), ox.Short("e")).
					String("env-file", "Load secret names and values from a dotenv-formatted file", ox.Spec("file"), ox.Short("f")).
					Bool("no-store", "Print the encrypted, base64-encoded value instead of storing it on GitHub").
					String("org", "Set organization secret", ox.Spec("organization"), ox.Short("o")).
					String("repos", "List of repositories that can access an organization or user secret", ox.Spec("repositories"), ox.Short("r")).
					Bool("user", "Set a secret for your user", ox.Short("u")).
					String("visibility", "Set visibility for an organization secret: {all|private|selected}", ox.Default("private"), ox.Short("v")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		), ox.Sub(
			ox.Banner("Manage SSH keys registered with your GitHub account."),
			ox.Usage("ssh-key", "Manage SSH keys"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Add an SSH key to your GitHub account"),
				ox.Usage("add", "Add an SSH key to your GitHub account"),
				ox.Spec("[<key-file>] [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("title", "Title for the new key", ox.Short("t")).
					String("type", "Type of the ssh key: {authentication|signing}", ox.Default("authentication")),
			), ox.Sub(
				ox.Banner("Delete an SSH key from your GitHub account"),
				ox.Usage("delete", "Delete an SSH key from your GitHub account"),
				ox.Spec("<id> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					Bool("yes", "Skip the confirmation prompt", ox.Short("y")),
			), ox.Sub(
				ox.Banner("Lists SSH keys in your GitHub account"),
				ox.Usage("list", "Lists SSH keys in your GitHub account"),
				ox.Spec("[flags]"),
				ox.Aliases("gh ssh-key ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			)), ox.Sub(
			ox.Banner("The status command prints information about your work on GitHub across all the repositories you're subscribed to, including:\n\n- Assigned Issues\n- Assigned Pull Requests\n- Review Requests\n- Mentions\n- Repository Activity (new issues/pull requests, comments)"),
			ox.Usage("status", "Print information about relevant issues, pull requests, and notifications across repositories"),
			ox.Spec("[flags]"),
			ox.Example("\n  $ gh status -e cli/cli -e cli/go-gh # Exclude multiple repositories\n  $ gh status -o cli # Limit results to a single organization"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Flags().
				Slice("exclude", "Comma separated list of repos to exclude in owner/name format", ox.Elem(ox.StringT), ox.Short("e")).
				String("org", "Report status within an organization", ox.Short("o")),
		), ox.Sub(
			ox.Banner("Variables can be set at the repository, environment or organization level for use in\nGitHub Actions or Dependabot. Run `gh help variable set` to learn how to get started."),
			ox.Usage("variable", "Manage GitHub Actions variables"),
			ox.Spec("<command> [flags]"),
			ox.Section(3),
			ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
			ox.Sub(
				ox.Banner("Delete a variable on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs or Dependabot within an organization"),
				ox.Usage("delete", "Delete variables"),
				ox.Spec("<variable-name> [flags]"),
				ox.Aliases("gh variable remove"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("env", "Delete a variable for an environment", ox.Short("e")).
					String("org", "Delete a variable for an organization", ox.Short("o")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Get a variable on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs or Dependabot within an organization\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("get", "Get variables"),
				ox.Spec("<variable-name> [flags]"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("env", "Get a variable for an environment", ox.Short("e")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("org", "Get a variable for an organization", ox.Short("o")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("List variables on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs or Dependabot within an organization\n\nFor more information about output formatting flags, see `gh help formatting`."),
				ox.Usage("list", "List variables"),
				ox.Spec("[flags]"),
				ox.Aliases("gh variable ls"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("env", "List variables for an environment", ox.Short("e")).
					String("jq", "Filter JSON output using a jq expression", ox.Spec("expression"), ox.Short("q")).
					Slice("json", "Output JSON with the specified fields", ox.Elem(ox.StringT)).
					String("org", "List variables for an organization", ox.Short("o")).
					String("template", "Format JSON output using a Go template; see \"gh help formatting\"", ox.Short("t")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Sub(
				ox.Banner("Set a value for a variable on one of the following levels:\n- repository (default): available to GitHub Actions runs or Dependabot in a repository\n- environment: available to GitHub Actions runs for a deployment environment in a repository\n- organization: available to GitHub Actions runs or Dependabot within an organization\n\nOrganization variable can optionally be restricted to only be available to\nspecific repositories."),
				ox.Usage("set", "Create or update variables"),
				ox.Spec("<variable-name> [flags]"),
				ox.Example("\n  # Add variable value for the current repository in an interactive prompt\n  $ gh variable set MYVARIABLE\n  \n  # Read variable value from an environment variable\n  $ gh variable set MYVARIABLE --body \"$ENV_VALUE\"\n  \n  # Read variable value from a file\n  $ gh variable set MYVARIABLE < myfile.txt\n  \n  # Set variable for a deployment environment in the current repository\n  $ gh variable set MYVARIABLE --env myenvironment\n  \n  # Set organization-level variable visible to both public and private repositories\n  $ gh variable set MYVARIABLE --org myOrg --visibility all\n  \n  # Set organization-level variable visible to specific repositories\n  $ gh variable set MYVARIABLE --org myOrg --repos repo1,repo2,repo3\n  \n  # Set multiple variables imported from the \".env\" file\n  $ gh variable set -f .env"),
				ox.Footer("\n  Use `gh <command> <subcommand> --help` for more information about a command.\n  Read the manual at https://cli.github.com/manual\n  Learn about exit codes using `gh help exit-codes`"),
				ox.Flags().
					String("body", "The value for the variable (reads from standard input if not specified)", ox.Short("b")).
					String("env", "Set deployment environment variable", ox.Spec("environment"), ox.Short("e")).
					String("env-file", "Load variable names and values from a dotenv-formatted file", ox.Spec("file"), ox.Short("f")).
					String("org", "Set organization variable", ox.Spec("organization"), ox.Short("o")).
					String("repos", "List of repositories that can access an organization variable", ox.Spec("repositories"), ox.Short("r")).
					String("visibility", "Set visibility for an organization variable: {all|private|selected}", ox.Default("private"), ox.Short("v")).
					String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
			), ox.Flags().
				String("repo", "Select another repository using the [HOST/]OWNER/REPO format", ox.Spec("[HOST/]OWNER/REPO"), ox.Short("R")),
		),
	)
}
