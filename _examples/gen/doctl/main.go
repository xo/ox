// Command doctl is a xo/ox version of `doctl`.
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
		ox.Usage(`doctl`, ``),
		ox.Banner(`doctl is a command line interface (CLI) for the DigitalOcean API.`),
		ox.Spec(`[command]`),
		ox.Sections("Manage DigitalOcean Resources", "Configure doctl", "View Billing"),
		ox.Footer(`Use "doctl [command] --help" for more information about a command.`),
		ox.Flags().
			String(`access-token`, `API V2 access token`, ox.Short("t")).
			String(`api-url`, `Override default API endpoint`, ox.Short("u")).
			String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c")).
			String(`context`, `Specify a custom authentication context name`).
			Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5")).
			Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false")).
			String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o")).
			Bool(`trace`, `Show a log of network activity while performing a command`).
			Bool(`verbose`, `Enable verbose output`, ox.Short("v")),
		ox.Sub( // doctl 1-click
			ox.Usage(`1-click`, `Display commands that pertain to 1-click applications`),
			ox.Banner(`The commands under `+"`"+`doctl 1-click`+"`"+` are for interacting with DigitalOcean 1-Click applications.`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl 1-click [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl 1-click list
				ox.Usage(`list`, `Retrieve a list of 1-Click applications`),
				ox.Banner(`Use this command to retrieve a list of 1-Click applications. You can narrow it by type, current types: kubernetes, droplet`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example retrieves a list of 1-Click applications available for Droplets: doctl 1-click list --type droplet`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: SLUG, `+"`"+`TYPE`+"`"+`.`, ox.Spec(`SLUG`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`type`, `The 1-Click type. Valid types are one of the following: kubernetes, droplet`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl account
			ox.Usage(`account`, `Display commands that retrieve account details`),
			ox.Banner(`The subcommands of `+"`"+`doctl account`+"`"+` retrieve information about DigitalOcean accounts.

For example, `+"`"+`doctl account get`+"`"+` retrieves account profile details, and `+"`"+`doctl account ratelimit`+"`"+` retrieves API usage details.`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl account [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl account get
				ox.Usage(`get`, `Retrieve account profile details`),
				ox.Banner(`Retrieve the following details from your account profile:

- Email address
- Team 
- Account Droplet limit
- Email verification status
- UUID for the account
- Account status (active or disabled).`),
				ox.Spec(`[flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves email addresses associated with the account: doctl account get --format Email`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Email, `+"`"+`Team`+"`"+`, `+"`"+`DropletLimit`+"`"+`, `+"`"+`EmailVerified`+"`"+`, `+"`"+`UUID`+"`"+`, `+"`"+`Status`+"`"+`.`, ox.Default("Email")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl account ratelimit
				ox.Usage(`ratelimit`, `Retrieve your API usage and the remaining quota`),
				ox.Banner(`Retrieve the following details about your account's API usage:

- The current limit on your account for API calls (default is 5,000 per hour per OAuth token)
- The number of API calls you have made in the last hour
- When the API call count resets to zero, which happens hourly

Note that these details are per OAuth token and are bound to the token you used when calling `+"`"+`doctl auth init`+"`"+` at setup time.`),
				ox.Spec(`[flags]`),
				ox.Aliases("rl"),
				ox.Example(`
The following example retrieves the number of API calls you have left for the hour: doctl account ratelimit --format Remaining`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Limit, `+"`"+`Remaining`+"`"+`, `+"`"+`Reset`+"`"+`.`, ox.Default("Limit")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl apps
			ox.Usage(`apps`, `Displays commands for working with apps`),
			ox.Banner(`The subcommands of `+"`"+`doctl app`+"`"+` manage your App Platform apps. For documentation on app specs, see the [app spec reference](https://www.digitalocean.com/docs/app-platform/concepts/app-spec).`),
			ox.Spec(`[command]`),
			ox.Aliases("app", "a"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl apps [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl apps console
				ox.Usage(`console`, `Starts a console session`),
				ox.Banner(`Instantiates a console session for a component of an app.`),
				ox.Spec(`<app id> <component name> [flags]`),
				ox.Aliases("cs"),
				ox.Example(`
The following example initiates a console session for the app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` and the component `+"`"+`web`+"`"+`: doctl apps console f81d4fae-7dec-11d0-a765-00a0c91e6bf6 web`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`deployment`, `Starts a console session for a specific deployment ID. Defaults to current deployment.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps create
				ox.Usage(`create`, `Create an app`),
				ox.Banner(`Create an app with the given app spec.`),
				ox.Spec(`[flags]`),
				ox.Aliases("c"),
				ox.Example(`
The following example creates an app in a project named `+"`"+`example-project`+"`"+` using an app spec located in a directory called `+"`"+`/src/your-app.yaml`+"`"+`. Additionally, the command returns the new app's ID, ingress information, and creation date: doctl apps create --spec src/your-app.yaml --format ID,DefaultIngress,Created`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Name`+"`"+`, `+"`"+`DefaultIngress`+"`"+`, `+"`"+`ActiveDeployment.ID`+"`"+`, `+"`"+`InProgressDeployment.ID`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`project-id`, `The ID of the project to assign the created app and resources to. If not provided, the default project will be used.`).
					String(`spec`, `Path to an app spec in JSON or YAML format. Set to "-" to read from stdin. (required)`).
					Bool(`update-sources`, `Boolean that specifies whether, on update, the app should also update its source code`).
					Bool(`upsert`, `Boolean that specifies whether the app should be updated if it already exists`).
					Bool(`wait`, `Boolean that specifies whether to wait for an app to complete before returning control to the terminal`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps create-deployment
				ox.Usage(`create-deployment`, `Creates a deployment`),
				ox.Banner(`Deploys the app with the latest changes from your repository.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("cd"),
				ox.Example(`
The following example creates a deployment for an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`. Additionally, the command returns the app's ID and status: doctl apps create-deployment f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --format ID,Status`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force-rebuild`, `Force a re-build even if a previous build is eligible for reuse.`).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Cause`+"`"+`, `+"`"+`Progress`+"`"+`, `+"`"+`Phase`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					Bool(`wait`, `Boolean that specifies whether to wait for the deployment to complete before allowing further terminal input. This can be helpful for scripting.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps delete
				ox.Usage(`delete`, `Deletes an app`),
				ox.Banner(`Deletes the specified app.

This permanently deletes the app and all of its associated deployments.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("d", "rm"),
				ox.Example(`
The following example deletes an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl apps delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `Delete the App without a confirmation prompt`, ox.Short("f")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps dev
				ox.Usage(`dev`, `[BETA] Display commands for working with App Platform local development.`),
				ox.Banner(`[BETA] Display commands for working with App Platform local development.

  To get started, run `+"`"+`doctl app dev build`+"`"+`.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl apps dev [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl apps dev build
					ox.Usage(`build`, `Build an app component`),
					ox.Banner(`[BETA] Build an app component locally.

  The component name is optional unless running non-interactively.

  All command line flags as optional. You may specify flags to be applied to the current build
  or use the command `+"`"+`doctl app dev config`+"`"+` to permanently configure default values.`),
					ox.Spec(`[component name]`),
					ox.Aliases("b"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`app`, `An optional existing app ID. If specified, the app spec will be fetched from the given app.`).
						String(`build-command`, `An optional build command override for local development.`).
						String(`env-file`, `An optional path to a .env file with overrides for values of app spec environment variables.`).
						Bool(`no-cache`, `Set to disable build caching.`).
						String(`registry`, `An optional registry name to tag built container images with.`).
						String(`spec`, `An optional path to an app spec in JSON or YAML format. Default: .do/app.yaml.`).
						Duration(`timeout`, `An optional timeout duration for the build. Valid time units are "s", "m", "h". Example: 15m30s`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl apps dev config
					ox.Usage(`config`, `Displays commands for working with App Platform local development configuration settings.`),
					ox.Banner(`[BETA] Displays commands for working with App Platform local development configuration settings.

Configuration Format:

	timeout: An optional timeout duration for the build. Valid time units are 's', 'm', 'h'. Example: 15m30s
	app: ID of an App Platform App to load the AppSpec from.
	spec: Path to an AppSpec to load for builds.
	registry: An optional registry name used to tag built container images.
	no_cache: Boolean set to disable build caching.
	components:
	  # Per-component configuration
	  component-name: 
	    build_command: Custom build command override for a given component.
	    env_file: Path to an env file to override envs for a given component.`),
					ox.Spec(`[command]`),
					ox.Aliases("c"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl apps dev config [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl apps dev config set
						ox.Usage(`set`, `Set a value in the local development configuration settings.`),
						ox.Banner(`Set a value in the local development configuration settings.

KEY is the name of a configuration option, for example: spec=/path/to/app.yaml
Nested component KEYs can also be set, for example: components.my-component.build_command="go build ."

Multiple KEY=VALUE pairs may be specified separated by a space.

Configuration Format:

	timeout: An optional timeout duration for the build. Valid time units are 's', 'm', 'h'. Example: 15m30s
	app: ID of an App Platform App to load the AppSpec from.
	spec: Path to an AppSpec to load for builds.
	registry: An optional registry name used to tag built container images.
	no_cache: Boolean set to disable build caching.
	components:
	  # Per-component configuration
	  component-name: 
	    build_command: Custom build command override for a given component.
	    env_file: Path to an env file to override envs for a given component.`),
						ox.Spec(`KEY=VALUE... [flags]`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`dev-config`, `Path to the app dev config.`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl apps dev config unset
						ox.Usage(`unset`, `Unset a value in the local development configuration settings.`),
						ox.Banner(`Unset a value in the local development configuration settings.

KEY is the name of a configuration option to unset, for example: spec
Nested component KEYs can also be unset, for example: components.my-component.build_command

Multiple KEYs may be specified separated by a space.

Configuration Format:

	timeout: An optional timeout duration for the build. Valid time units are 's', 'm', 'h'. Example: 15m30s
	app: ID of an App Platform App to load the AppSpec from.
	spec: Path to an AppSpec to load for builds.
	registry: An optional registry name used to tag built container images.
	no_cache: Boolean set to disable build caching.
	components:
	  # Per-component configuration
	  component-name: 
	    build_command: Custom build command override for a given component.
	    env_file: Path to an env file to override envs for a given component.`),
						ox.Spec(`KEY... [flags]`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`dev-config`, `Path to the app dev config.`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
			),
			ox.Sub( // doctl apps get
				ox.Usage(`get`, `Get an app`),
				ox.Banner(`Get an app with the provided id.

Only basic information is included with the text output format. For complete app details including its app spec, use the JSON format.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("g"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Name`+"`"+`, `+"`"+`DefaultIngress`+"`"+`, `+"`"+`ActiveDeployment.ID`+"`"+`, `+"`"+`InProgressDeployment.ID`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps get-deployment
				ox.Usage(`get-deployment`, `Get a deployment`),
				ox.Banner(`Gets information about a specific deployment for the given app, including when the app updated and what triggered the deployment (Cause).

Only basic information is included with the text output format. For complete app details including an updated app spec, use the `+"`"+`--output`+"`"+` global flag and specify the JSON format.`),
				ox.Spec(`<app id> <deployment id> [flags]`),
				ox.Aliases("gd"),
				ox.Example(`
The following example gets information about a deployment with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+` for an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`. Additionally, the command returns the deployment's ID, status, and cause: doctl apps get-deployment f81d4fae-7dec-11d0-a765-00a0c91e6bf6 418b7972-fc67-41ea-ab4b-6f9477c4f7d8 --format ID,Status,Cause`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Cause`+"`"+`, `+"`"+`Progress`+"`"+`, `+"`"+`Phase`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps list
				ox.Usage(`list`, `Lists all apps`),
				ox.Banner(`Lists all apps associated with your account, including their ID, spec name, creation date, and other information.

Only basic information is included with the text output format. For complete app details including an updated app spec, use the `+"`"+`--output`+"`"+` global flag and specify the JSON format.`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following lists all apps in your account, but returns just their ID and creation date: doctl apps list --format ID,Created`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Name`+"`"+`, `+"`"+`DefaultIngress`+"`"+`, `+"`"+`ActiveDeployment.ID`+"`"+`, `+"`"+`InProgressDeployment.ID`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					Bool(`with-projects`, `Boolean that specifies whether project ids should be fetched along with listed apps`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps list-alerts
				ox.Usage(`list-alerts`, `Lists alerts on an app`),
				ox.Banner(`Lists all alerts associated to an app and its component, such as deployment failures and domain failures.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("la"),
				ox.Example(`
The following example lists all alerts associated to an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to specifically return the alert ID, trigger, and rule: doctl apps list-alerts f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --format ID,Trigger,Spec.Rule`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Rule`+"`"+`, `+"`"+`Trigger`+"`"+`, `+"`"+`ComponentName`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`SlackWebhooks`+"`"+`, `+"`"+`Spec.Disabled`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps list-buildpacks
				ox.Usage(`list-buildpacks`, `Lists buildpacks`),
				ox.Banner(`Lists all buildpacks available on App Platform`),
				ox.Spec(`[flags]`),
				ox.Example(`
The following example lists all buildpacks available on App Platform and uses the `+"`"+`--format`+"`"+` flag to specifically return the buildpack ID and version: doctl apps list-buildpacks --format ID,Version`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`ID`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`Documentation`+"`"+`.`, ox.Default("Name")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps list-deployments
				ox.Usage(`list-deployments`, `List all deployments`),
				ox.Banner(`List all deployments for an app.

Only basic information is included with the text output format. For complete app details including the app specs, use the JSON format.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("lsd"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Cause`+"`"+`, `+"`"+`Progress`+"`"+`, `+"`"+`Phase`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps list-regions
				ox.Usage(`list-regions`, `Lists available App Platform regions`),
				ox.Banner(`Lists all regions supported by App Platform, including details about their current availability.`),
				ox.Spec(`[flags]`),
				ox.Example(`
The following example lists all regions supported by App Platform, including details about their current availability: doctl apps list-regions --format DataCenters,Disabled,Reason`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Slug, `+"`"+`Label`+"`"+`, `+"`"+`Continent`+"`"+`, `+"`"+`DataCenters`+"`"+`, `+"`"+`Disabled`+"`"+`, `+"`"+`Reason`+"`"+`, `+"`"+`Default`+"`"+`.`, ox.Default("Slug")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps logs
				ox.Usage(`logs`, `Retrieves logs`),
				ox.Banner(`Retrieves component logs for a deployment of an app.

Three types of logs are supported and can be specified with the --type flag:
- build
- deploy
- run
- run_restarted 

For more information about logs, see [How to View Logs](https://www.digitalocean.com/docs/app-platform/how-to/view-logs/).`),
				ox.Spec(`<app name or id> <component name (defaults to all components)> [flags]`),
				ox.Aliases("l"),
				ox.Example(`
The following example retrieves the build logs for the app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` and the component `+"`"+`web`+"`"+`: doctl apps logs f81d4fae-7dec-11d0-a765-00a0c91e6bf6 web --type build`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`deployment`, `Retrieves logs for a specific deployment ID. Defaults to current deployment.`).
					Bool(`follow`, `Returns logs as they are emitted by the app.`, ox.Short("f")).
					Bool(`no-prefix`, `Removes the prefix from logs. Useful for JSON structured logs`).
					Int(`tail`, `Specifies the number of lines to show from the end of the log.`, ox.Default("-1")).
					String(`type`, `Retrieves logs for a specific log type. Defaults to run logs.`, ox.Default("run")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps propose
				ox.Usage(`propose`, `Proposes an app spec`),
				ox.Banner(`Reviews and validates an app specification for a new or existing app. The request returns some information about the proposed app, including app cost and upgrade cost. If an existing app ID is specified, the app spec is treated as a proposed update to the existing app.

Only basic information is included with the text output format. For complete app details including an updated app spec, use the `+"`"+`--output`+"`"+` global flag and specify the JSON format.`),
				ox.Spec(`[flags]`),
				ox.Aliases("p"),
				ox.Example(`
The following example proposes an app spec from the file directory `+"`"+`src/your-app.yaml`+"`"+` for a new app: doctl apps propose --spec src/your-app.yaml`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`app`, `An optional existing app ID. If specified, App Platform treats the spec as a proposed update to the existing app.`).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Name`+"`"+`, `+"`"+`DefaultIngress`+"`"+`, `+"`"+`ActiveDeployment.ID`+"`"+`, `+"`"+`InProgressDeployment.ID`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`spec`, `Path to an app spec in JSON or YAML format. For more information about app specs, see the [app spec reference](https://www.digitalocean.com/docs/app-platform/concepts/app-spec) (required)`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps restart
				ox.Usage(`restart`, `Restarts an app`),
				ox.Banner(`Restarts the specified app or some of its components.`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("r"),
				ox.Example(`
The following example restarts an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`. Additionally, the command returns the app's ID and status: doctl apps restart f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --format ID,Status`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Slice(`components`, `The components to restart. If not provided, all components are restarted.`).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Cause`+"`"+`, `+"`"+`Progress`+"`"+`, `+"`"+`Phase`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					Bool(`wait`, `Boolean that specifies whether to wait for the restart to complete before allowing further terminal input. This can be helpful for scripting.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps spec
				ox.Usage(`spec`, `Display commands for working with app specs`),
				ox.Banner(`The subcommands of `+"`"+`doctl app spec`+"`"+` manage your app specs.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl apps spec [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl apps spec get
					ox.Usage(`get`, `Retrieve an application's spec`),
					ox.Banner(`Use this command to retrieve the latest spec of an app.

Optionally, pass a deployment ID to get the spec of that specific deployment.`),
					ox.Spec(`<app id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`deployment`, `optional: a deployment ID`).
						String(`format`, `the format to output the spec in; either "yaml" or "json"`, ox.Default("yaml")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl apps spec validate
					ox.Usage(`validate`, `Validate an application spec`),
					ox.Banner(`Use this command to check whether a given app spec (YAML or JSON) is valid.

You may pass - as the filename to read from stdin.`),
					ox.Spec(`<spec file> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`schema-only`, `Only validate the spec schema and not the correctness of the spec.`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl apps tier
				ox.Usage(`tier`, `Display commands for working with app tiers`),
				ox.Banner(`The subcommands of `+"`"+`doctl app tier`+"`"+` retrieve information about app tiers.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl apps tier [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl apps tier instance-size
					ox.Usage(`instance-size`, `Display commands for working with app instance sizes`),
					ox.Banner(`The subcommands of `+"`"+`doctl app tier instance-size`+"`"+` retrieve information about app instance sizes.`),
					ox.Spec(`[command]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl apps tier instance-size [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl apps tier instance-size get
						ox.Usage(`get`, `Retrieve an app instance size`),
						ox.Banner(`Use this command to retrieve information about a specific app instance size.`),
						ox.Spec(`<instance size slug> [flags]`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl apps tier instance-size list
						ox.Usage(`list`, `List all app instance sizes`),
						ox.Banner(`Use this command to list all the available app instance sizes.`),
						ox.Spec(`[flags]`),
						ox.Aliases("ls"),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
			),
			ox.Sub( // doctl apps update
				ox.Usage(`update`, `Updates an app`),
				ox.Banner(`Updates the specified app with the given app spec. For more information about app specs, see the [app spec reference](https://www.digitalocean.com/docs/app-platform/concepts/app-spec)`),
				ox.Spec(`<app id> [flags]`),
				ox.Aliases("u"),
				ox.Example(`
The following example updates an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` using an app spec located in a directory called `+"`"+`/src/your-app.yaml`+"`"+`. Additionally, the command returns the updated app's ID, ingress information, and creation date: doctl apps update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --spec src/your-app.yaml --format ID,DefaultIngress,Created`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Name`+"`"+`, `+"`"+`DefaultIngress`+"`"+`, `+"`"+`ActiveDeployment.ID`+"`"+`, `+"`"+`InProgressDeployment.ID`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`spec`, `Path to an app spec in JSON or YAML format. Set to "-" to read from stdin. (required)`).
					Bool(`update-sources`, `Boolean that specifies whether the app should also update its source code`).
					Bool(`wait`, `Boolean that specifies whether to wait for an app to complete updating before allowing further terminal input. This can be helpful for scripting.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps update-alert-destinations
				ox.Usage(`update-alert-destinations`, `Updates alert destinations`),
				ox.Banner(`Updates alert destinations`),
				ox.Spec(`<app id> <alert id> [flags]`),
				ox.Aliases("uad"),
				ox.Example(`
The following example updates the alert destinations for an app with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` and the alert ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl apps update-alert-destinations f81d4fae-7dec-11d0-a765-00a0c91e6bf6 f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --alert-destinations src/your-alert-destinations.yaml`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`app-alert-destinations`, `Path to an alert destinations file in JSON or YAML format.`).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Spec.Rule`+"`"+`, `+"`"+`Trigger`+"`"+`, `+"`"+`ComponentName`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`SlackWebhooks`+"`"+`, `+"`"+`Spec.Disabled`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl apps upgrade-buildpack
				ox.Usage(`upgrade-buildpack`, `Upgrades app's buildpack`),
				ox.Banner(`Upgrades an app's buildpack. For more information about buildpacks, see the [buildpack reference](https://docs.digitalocean.com/products/app-platform/reference/buildpacks/)`),
				ox.Spec(`<app id> [flags]`),
				ox.Example(`
The following example upgrades an app's buildpack with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to the latest available version: doctl apps upgrade-buildpack f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --buildpack f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`buildpack`, `The ID of the buildpack to upgrade to. Use the list-buildpacks command to list available buildpacks. (required)`).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Cause`+"`"+`, `+"`"+`Progress`+"`"+`, `+"`"+`Phase`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`.`, ox.Spec(`ID`)).
					Int(`major-version`, `The major version to upgrade to. If empty, the buildpack upgrades to the latest available version.`).
					Bool(`no-header`, `Return raw data with no headers`).
					Bool(`trigger-deployment`, `Specifies whether to trigger a new deployment to apply the upgrade.`, ox.Default("true")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl compute
			ox.Usage(`compute`, `Display commands that manage infrastructure`),
			ox.Banner(`The subcommands under `+"`"+`doctl compute`+"`"+` are for managing DigitalOcean resources.`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl compute [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl compute action
				ox.Usage(`action`, `Display commands for retrieving resource action history`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute action`+"`"+` retrieve the history of actions taken on your resources.

You can retrieve information for a specific action by adding the action's ID as an argument. For example, while `+"`"+`doctl compute action list`+"`"+` lists all of the actions taken on all of the resources in your account, `+"`"+`doctl compute action get <action-id>`+"`"+` retrieves details for a specific action. Additionally, you can use `+"`"+`--action-type`+"`"+` flag to filter the list of actions by type. For example, `+"`"+`doctl compute action list --action-type power_on`+"`"+` lists all of the actions that powered on a resource.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute action [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute action get
					ox.Usage(`get`, `Retrieve details about a specific action`),
					ox.Banner(`Retrieve the following details about a specific action taken on one of your resources:

- The action ID
- The action status (`+"`"+`pending`+"`"+`, `+"`"+`completed`+"`"+`, etc)
- The action type, such as: `+"`"+`create`+"`"+`, `+"`"+`destroy`+"`"+`, `+"`"+`power_cycle`+"`"+`, `+"`"+`power_off`+"`"+`, `+"`"+`power_on`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`migrate`+"`"+`, `+"`"+`attach_volume`+"`"+`
- The Date/Time when the action started, in RFC3339 format
- The Date/Time when the action completed, in RFC3339 format
- The resource ID of the resource upon which the action was taken
- The resource type (Droplet, backend)
- The region in which the action took place (nyc3, sfo2, etc)`),
					ox.Spec(`<action-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the action's ID, status, and resource type of the action with ID 123456: doctl compute action get 123456 --format ID,Status,ResourceType`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute action list
					ox.Usage(`list`, `Retrieve a  list of all recent actions taken on your resources`),
					ox.Banner(`Retrieve a list of all actions taken on your resources. The following details are provided:

- The action ID
- The action status (`+"`"+`pending`+"`"+`, `+"`"+`completed`+"`"+`, etc)
- The action type, such as: `+"`"+`create`+"`"+`, `+"`"+`destroy`+"`"+`, `+"`"+`power_cycle`+"`"+`, `+"`"+`power_off`+"`"+`, `+"`"+`power_on`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`migrate`+"`"+`, `+"`"+`attach_volume`+"`"+`
- The Date/Time when the action started, in RFC3339 format
- The Date/Time when the action completed, in RFC3339 format
- The resource ID of the resource upon which the action was taken
- The resource type (Droplet, backend)
- The region in which the action took place (nyc3, sfo2, etc)`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following command retrieves a list of all the destroy actions taken on the account after October 12, 2022 at 12:00:01 AM UTC, and displays the action ID and region: doctl compute action list --action-type destroy --after 2022-10-12T00:00:01.00Z --format ID,Region`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`action-type`, `Filter by action type, such as create or `+"`"+`destroy`+"`"+``, ox.Default("create")).
						String(`after`, `Filter actions taken after a specified date, in RFC3339 format.`).
						String(`before`, `Filter actions taken after a specified date, in RFC3339 format.`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `Filter by a specified datacenter region, such as nyc`, ox.Default("nyc")).
						String(`resource-type`, `Filter by action resource type, such as droplet`, ox.Default("droplet")).
						String(`status`, `Filter by action status, such as completed or `+"`"+`in-progress`+"`"+`.`, ox.Default("completed")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute action wait
					ox.Usage(`wait`, `Block thread until an action completes`),
					ox.Banner(`Block the current thread, returning when an action completes.

For example, if you find an action when calling `+"`"+`doctl compute action list`+"`"+` that has a status of `+"`"+`in-progress`+"`"+`, you can note the action ID and call `+"`"+`doctl compute action wait <action-id>`+"`"+`, and doctl will appear to "hang" until the action has completed. This can be useful for scripting purposes.`),
					ox.Spec(`<action-id> [flags]`),
					ox.Aliases("w"),
					ox.Example(`
The following example waits for the action `+"`"+`123456`+"`"+` to complete before allowing further commands to execute: doctl compute action wait 123456`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Int(`poll-timeout`, `Re-poll time in seconds`, ox.Default("5")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute cdn
				ox.Usage(`cdn`, `Display commands that manage CDNs`),
				ox.Banner(`The subcommands of `+"`"+`doctl compute cdn`+"`"+` are for managing Content Delivery Networks (CDNs).

Content hosted in DigitalOcean's object storage solution, Spaces, can optionally be served by our globally distributed CDNs. This allows you to deliver content to users based on their geographic location.

To use a custom subdomain to access the CDN endpoint, provide the ID of a DigitalOcean-managed TLS certificate and the fully qualified domain name (FQDN) for the custom subdomain.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute cdn [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute cdn create
					ox.Usage(`create`, `Create a CDN`),
					ox.Banner(`Creates a Content Delivery Network (CDN) on the origin server you specify and automatically generates an endpoint. You can also use a custom subdomain you own to create an additional endpoint, which must be secured with SSL.

The Time To Live (TTL) value is the length of time in seconds that a file is cached by the CDN before being refreshed. If a request to access a file occurs after the TTL has expired, the CDN delivers the file by requesting it directly from the origin URL, re-caching the file, and resetting the TTL.`),
					ox.Spec(`<cdn-origin> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a CDN for the custom domain `+"`"+`cdn.example.com `+"`"+` using a DigitalOcean Spaces origin endpoint and SSL certificate ID for the custom domain: doctl compute cdn create https://tester-two.blr1.digitaloceanspaces.com --domain cdn.example.com --certificate-id f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`certificate-id`, `Specify a certificate ID for the custom domain`).
						String(`domain`, `Specify a custom domain to use with the CDN`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Origin`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`CustomDomain`+"`"+`, `+"`"+`CertificateID`+"`"+`, `+"`"+`CreatedAt`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Int(`ttl`, `The "Time To Live" (TTL) value for cached content, in seconds`, ox.Default("3600")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute cdn delete
					ox.Usage(`delete`, `Delete a CDN`),
					ox.Banner(`Deletes the CDN specified by the ID.

You can retrieve a list of CDN IDs by calling `+"`"+`doctl compute cdn list`+"`"+``),
					ox.Spec(`<cdn-id> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes a CDN with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+`: doctl compute cdn delete 418b7972-fc67-41ea-ab4b-6f9477c4f7d8`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the specified CDN without prompting for confirmation`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute cdn flush
					ox.Usage(`flush`, `Flush the cache of a CDN`),
					ox.Banner(`Flushes the cache of a Content Delivery Network (CDN), which:

- purges all copies of the files in the cache
- re-caches the files
- retrieves files from the origin server for any requests that hit the CDN endpoint until all the files are re-cached

This ensures that recently updated files on the origin server are immediately available via the CDN.

To purge specific files, you can use the `+"`"+`--files`+"`"+` flag and supply a path to the file in the Spaces bucket. The path may be for a single file or may contain a wildcard (`+"`"+`*`+"`"+`) to recursively purge all files under a directory. When only a wildcard is provided, or no path is provided, all cached files will be purged.
Examples:		
 doctl compute cdn flush <cdn-id>  --files /path/to/assets/*
 doctl compute cdn flush <cdn-id>  --files "/path/to/file.one, /path/to/file.two"
 doctl compute cdn flush <cdn-id>  --files /path/to/file.one --files /path/to/file.two
 doctl compute cdn flush <cdn-id>  --files *`),
					ox.Spec(`<cdn-id> [flags]`),
					ox.Aliases("fc"),
					ox.Example(`
The following example flushes the cache of the `+"`"+`/path/to/assets`+"`"+` directory in a CDN: doctl compute cdn flush 418b7972-fc67-41ea-ab4b-6f9477c4f7d8 --files /path/to/assets/*`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`files`, `cdn files`, ox.Default("[*]")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute cdn get
					ox.Usage(`get`, `Retrieve details about a specific CDN`),
					ox.Banner(`Lists the following details for the specified Content Delivery Network (CDNs):

- The ID for the CDN, in UUID format
- The fully qualified domain name (FQDN) for the origin server, which provides the content to the CDN. Currently, only Spaces are supported with CDNs.
- The fully qualified domain name (FQDN) of the endpoint from which the CDN-backed content is served.
- The "Time To Live" (TTL) value for cached content, in seconds. The default is 3,600 (one hour).
- An optional custom subdomain when the CDN can be accessed
- The ID of a DigitalOcean-managed TLS certificate used for SSL when a custom subdomain is provided.
- The date and time when the CDN was created, in ISO8601 date/time format

The Time To Live (TTL) value is the length of time in seconds that a file is cached by the CDN before being refreshed. If a request to access a file occurs after the TTL has expired, the CDN delivers the file by requesting it directly from the origin URL, re-caching the file, and resetting the TTL.`),
					ox.Spec(`<cdn-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the origin endpoint, CDN endpoint, and certificate ID for a CDN with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+`: doctl compute cdn get 418b7972-fc67-41ea-ab4b-6f9477c4f7d8 --format ID,Origin,Endpoint,CertificateID`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Origin`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`CustomDomain`+"`"+`, `+"`"+`CertificateID`+"`"+`, `+"`"+`CreatedAt`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute cdn list
					ox.Usage(`list`, `List CDNs that have already been created`),
					ox.Banner(`Retrieves a list of your existing Content Delivery Networks (CDNs) and their following details:

- The ID for the CDN, in UUID format
- The fully qualified domain name (FQDN) for the origin server, which provides the content to the CDN. Currently, only Spaces are supported with CDNs.
- The fully qualified domain name (FQDN) of the endpoint from which the CDN-backed content is served.
- The "Time To Live" (TTL) value for cached content, in seconds. The default is 3,600 (one hour).
- An optional custom subdomain when the CDN can be accessed
- The ID of a DigitalOcean-managed TLS certificate used for SSL when a custom subdomain is provided.
- The date and time when the CDN was created, in ISO8601 date/time format`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of CDNs for your account. The command uses the `+"`"+`--format`+"`"+` flag to only return each CDN`+"`"+`'`+"`"+`s origin endpoint, CDN endpoint, and certificate ID: doctl compute cdn list --format ID,Origin,Endpoint,CertificateID`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Origin`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`CustomDomain`+"`"+`, `+"`"+`CertificateID`+"`"+`, `+"`"+`CreatedAt`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute cdn update
					ox.Usage(`update`, `Update the configuration for a CDN`),
					ox.Banner(`Updates the configuration details of an existing Content Delivery Network (CDN).`),
					ox.Spec(`<cdn-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates the TTL for a CDN with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+` to 600 seconds: doctl compute cdn update 418b7972-fc67-41ea-ab4b-6f9477c4f7d8 --ttl 600`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`certificate-id`, `Specify a certificate ID for the custom domain`).
						String(`domain`, `Specify a custom domain to use with the CDN`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Origin`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`CustomDomain`+"`"+`, `+"`"+`CertificateID`+"`"+`, `+"`"+`CreatedAt`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Int(`ttl`, `The "Time To Live" (TTL) value for cached content, in seconds`, ox.Default("3600")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute certificate
				ox.Usage(`certificate`, `Display commands that manage SSL certificates and private keys`),
				ox.Banner(`The subcommands of `+"`"+`doctl compute certificate`+"`"+` allow you to store and manage your SSL certificates, private keys, and certificate paths.

Once a certificate has been stored, it is assigned a unique certificate ID that can be referenced in your doctl and API workflows.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute certificate [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute certificate create
					ox.Usage(`create`, `Create a new certificate`),
					ox.Banner(`Creates a new Let's Encrypt certificate or adds an existing custom certificate to your team. There are two supported certificate types: Let's Encrypt certificates, and custom certificates.

Let's Encrypt certificates are free, auto-renewed and managed for you by DigitalOcean.

To create a Let's Encrypt certificate, you need to add the domain(s) to your account at using the DigitalOcean control panel, or via `+"`"+`doctl compute domain create`+"`"+`, then provide a certificate name and a comma-separated list of the domain names you'd like to associate with the certificate:

	doctl compute certificate create --type lets_encrypt --name mycert --dns-names example.org

To upload a custom certificate, you need to provide a certificate name, the path to the certificate, the path to the certificate's private key, and the path to the certificate chain, all in PEM format:

	doctl compute certificate create --type custom --name mycert --leaf-certificate-path cert.pem --certificate-chain-path fullchain.pem --private-key-path privkey.pem`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`certificate-chain-path`, `The path on your local machine to a full PEM-formatted trust chain between the certificate authority's certificate and your domain's SSL certificate.`).
						Slice(`dns-names`, `Comma-separated list of domains for which the certificate will be issued. The domains must be managed using DigitalOcean's DNS.`).
						String(`leaf-certificate-path`, `The path on your local machine to a PEM-formatted public SSL certificate.`).
						String(`name`, `A user-specified name for the certificate. (required)`).
						String(`private-key-path`, `The path on your local machine to a PEM-formatted private-key corresponding to the SSL certificate.`).
						String(`type`, `The type of certificate, custom or `+"`"+`lets_encrypt`+"`"+`.`, ox.Default("custom")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute certificate delete
					ox.Usage(`delete`, `Delete the specified certificate`),
					ox.Banner(`Deletes the specified certificate.

Use `+"`"+`doctl compute certificate list`+"`"+` to see all available certificates associated with your account.`),
					ox.Spec(`<id> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes the certificate with the ID  `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute certificate delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the certificate without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute certificate get
					ox.Usage(`get`, `Retrieve details about a certificate`),
					ox.Banner(`This command retrieves the following details about a certificate:

- The certificate ID
- The name you gave the certificate
- A comma-separated list of domain names associated with the certificate
- The SHA-1 fingerprint of the certificate
- The certificate's expiration date, in ISO8601 date/time format
- The certificate's creation date, in ISO8601 date/time format
- The certificate type (`+"`"+`custom`+"`"+` or `+"`"+`lets_encrypt`+"`"+`)
- The certificate state (`+"`"+`pending`+"`"+`, `+"`"+`verified`+"`"+`, or `+"`"+`error`+"`"+`)`),
					ox.Spec(`<id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the ID, name, and domains associated with a certificate: doctl compute certificate get f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --format ID,Name,DNSNames`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`DNSNames`+"`"+`, `+"`"+`SHA1Fingerprint`+"`"+`, `+"`"+`NotAfter`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`State`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute certificate list
					ox.Usage(`list`, `Retrieve list of the account's stored certificates`),
					ox.Banner(`This command retrieves a list of all certificates associated with the account. The following details are shown for each certificate:

- The certificate ID
- The name you gave the certificate
- A comma-separated list of domain names associated with the certificate
- The SHA-1 fingerprint of the certificate
- The certificate's expiration date, in ISO8601 date/time format
- The certificate's creation date, in ISO8601 date/time format
- The certificate type (`+"`"+`custom`+"`"+` or `+"`"+`lets_encrypt`+"`"+`)
- The certificate state (`+"`"+`pending`+"`"+`, `+"`"+`verified`+"`"+`, or `+"`"+`error`+"`"+`)`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of all certificates associated with your account and uses the `+"`"+`--format`+"`"+` flag return only the IDs, names, and the domains associated with each ticket: doctl compute certificate list --format ID,Name,DNSNames`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`DNSNames`+"`"+`, `+"`"+`SHA1Fingerprint`+"`"+`, `+"`"+`NotAfter`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`State`+"`"+`.`, ox.Spec(`ID`)).
						String(`name`, `Filter certificates by the specified name`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute domain
				ox.Usage(`domain`, `Display commands that manage domains`),
				ox.Banner(`Use the subcommands of `+"`"+`doctl compute domain`+"`"+` to manage domains you have purchased from a domain name registrar that you are managing through the DigitalOcean DNS interface.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute domain [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute domain create
					ox.Usage(`create`, `Add a domain to your account`),
					ox.Banner(`Adds a domain to your account that you can assign to Droplets, load balancers, and other resources.`),
					ox.Spec(`<domain> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following command creates a domain named example.com and adds an A record to the domain: doctl compute domain create example.com --ip-address 198.51.100.215`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Domain, `+"`"+`TTL`+"`"+`.`, ox.Default("Domain")).
						String(`ip-address`, `Creates an A record for a IPv4 address`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute domain delete
					ox.Usage(`delete`, `Permanently delete a domain from your account`),
					ox.Banner(`Permanently deletes a domain from your account. You cannot undo this command once done.`),
					ox.Spec(`<domain> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following command deletes the domain example.com: doctl compute domain delete example.com`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the domain without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute domain get
					ox.Usage(`get`, `Retrieve information about a domain`),
					ox.Banner(`Retrieves information about a domain on your account.`),
					ox.Spec(`<domain> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following command retrieves information about the domain example.com: doctl compute domain get example.com`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Domain, `+"`"+`TTL`+"`"+`.`, ox.Default("Domain")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute domain list
					ox.Usage(`list`, `List all domains on your account`),
					ox.Banner(`Retrieves a list of domains on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following command lists all domains on your account: doctl compute domain list`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Domain, `+"`"+`TTL`+"`"+`.`, ox.Default("Domain")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute domain records
					ox.Usage(`records`, `Manage DNS records`),
					ox.Banner(`Use the subcommands of `+"`"+`doctl compute domain records`+"`"+` to manage the DNS records for your domains.`),
					ox.Spec(`[command]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl compute domain records [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl compute domain records create
						ox.Usage(`create`, `Create a DNS record`),
						ox.Banner(`Create DNS records for a domain.`),
						ox.Spec(`<domain> [flags]`),
						ox.Aliases("c"),
						ox.Example(`
The following command creates an A record for the domain example.com: doctl compute domain records create example.com --record-type A --record-name example.com --record-data 198.51.100.215`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Type`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Data`+"`"+`, `+"`"+`Priority`+"`"+`, `+"`"+`Port`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`Weight`+"`"+`, `+"`"+`Flags`+"`"+`, `+"`"+`Tag`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`record-data`, `The record's data. This value varies depending on record type.`).
							Int(`record-flags`, `The flag value of a CAA record. A valid is an unsigned integer between 0-255.`).
							String(`record-name`, `The host name, alias, or service being defined by the record`).
							Int(`record-port`, `The port value for an SRV record`).
							Int(`record-priority`, `The priority for an MX or SRV record`).
							String(`record-tag`, `The parameter tag for CAA records. Valid values are issue, `+"`"+`issuewild`+"`"+`, or `+"`"+`iodef`+"`"+``, ox.Default("issue")).
							Int(`record-ttl`, `The record's Time To Live (TTL) value, in seconds`, ox.Default("1800")).
							String(`record-type`, `The type of DNS record. Valid values are: A, `+"`"+`AAAA`+"`"+`, `+"`"+`CAA`+"`"+`, `+"`"+`CNAME`+"`"+`, `+"`"+`MX`+"`"+`, `+"`"+`NS`+"`"+`, `+"`"+`SOA`+"`"+`, `+"`"+`SRV`+"`"+`, and `+"`"+`TXT`+"`"+`.`, ox.Default("A")).
							Int(`record-weight`, `The weight value for an SRV record`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl compute domain records delete
						ox.Usage(`delete`, `Delete a DNS record`),
						ox.Banner(`Deletes DNS records for a domain.`),
						ox.Spec(`<domain> <record-id>... [flags]`),
						ox.Aliases("d", "rm"),
						ox.Example(`
The following command deletes a DNS record with the ID `+"`"+`98858421`+"`"+` from the domain `+"`"+`example.com`+"`"+`: doctl compute domain records delete example.com 98858421`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`force`, `Delete record without confirmation prompt`, ox.Short("f")).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl compute domain records list
						ox.Usage(`list`, `List the DNS records for a domain`),
						ox.Banner(`Lists the DNS records for a domain.`),
						ox.Spec(`<domain> [flags]`),
						ox.Aliases("ls"),
						ox.Example(`
The following command lists the DNS records for the domain example.com. The command also uses the `+"`"+`--format`+"`"+` flag to only return each record's ID, type, and TTL: doctl compute domain records list example.com --format ID,Type,TTL`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Type`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Data`+"`"+`, `+"`"+`Priority`+"`"+`, `+"`"+`Port`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`Weight`+"`"+`, `+"`"+`Flags`+"`"+`, `+"`"+`Tag`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl compute domain records update
						ox.Usage(`update`, `Update a DNS record`),
						ox.Banner(`Updates or changes the properties of DNS records for a domain.`),
						ox.Spec(`<domain> [flags]`),
						ox.Aliases("u"),
						ox.Example(`
The following command updates the record with the ID `+"`"+`98858421`+"`"+` for the domain `+"`"+`example.com`+"`"+`: doctl compute domain records update example.com --record-id 98858421 --record-name example.com --record-data 198.51.100.215`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Type`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Data`+"`"+`, `+"`"+`Priority`+"`"+`, `+"`"+`Port`+"`"+`, `+"`"+`TTL`+"`"+`, `+"`"+`Weight`+"`"+`, `+"`"+`Flags`+"`"+`, `+"`"+`Tag`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`record-data`, `The record's data. This value varies depending on record type.`).
							Int(`record-flags`, `The flag value of a CAA record. A valid is an unsigned integer between 0-255.`).
							Int(`record-id`, `The record's ID`).
							String(`record-name`, `The host name, alias, or service being defined by the record`).
							Int(`record-port`, `The port value for an SRV record`).
							Int(`record-priority`, `The priority for an MX or SRV record`).
							String(`record-tag`, `The parameter tag for CAA records. Valid values are issue, `+"`"+`issuewild`+"`"+`, or `+"`"+`iodef`+"`"+``, ox.Default("issue")).
							Int(`record-ttl`, `The record's Time To Live (TTL) value, in seconds`).
							String(`record-type`, `The type of DNS record`).
							Int(`record-weight`, `The weight value for an SRV record`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
			),
			ox.Sub( // doctl compute droplet
				ox.Usage(`droplet`, `Manage virtual machines (Droplets)`),
				ox.Banner(`A Droplet is a DigitalOcean virtual machine. Use the subcommands of `+"`"+`doctl compute droplet`+"`"+` to create, manage, and retrieve information about Droplets.`),
				ox.Spec(`[command]`),
				ox.Aliases("d"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute droplet [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute droplet 1-click
					ox.Usage(`1-click`, `Display commands that pertain to Droplet 1-click applications`),
					ox.Banner(`The commands under `+"`"+`doctl compute droplet 1-click`+"`"+` are for interacting with DigitalOcean Droplet 1-Click applications.`),
					ox.Spec(`[command]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl compute droplet 1-click [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl compute droplet 1-click list
						ox.Usage(`list`, `Retrieve a list of Droplet 1-Click applications`),
						ox.Banner(`Retrieves a list of Droplet 1-Click application slugs.

You can use 1-click slugs to create Droplets by using them as the argument for the `+"`"+`--image`+"`"+` flag in the `+"`"+`doctl compute droplet create`+"`"+` command. For example, the following command creates a Droplet with an Openblocks installation on it: `+"`"+`doctl compute droplet create example-droplet --image openblocks --size s-2vcpu-2gb --region nyc1`+"`"+``),
						ox.Spec(`[flags]`),
						ox.Aliases("ls"),
						ox.Example(`
The following example retrieves a list of 1-clicks for Droplets: doctl compute droplet 1-click list`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: SLUG, `+"`"+`TYPE`+"`"+`.`, ox.Spec(`SLUG`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl compute droplet actions
					ox.Usage(`actions`, `List Droplet actions`),
					ox.Banner(`Retrieves a list of previous actions taken on the Droplet, such as reboots, resizes, and snapshots actions.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("a"),
					ox.Example(`
The following example retrieves a list of actions taken on a Droplet with the ID `+"`"+`386734086`+"`"+`. Additionally, the command uses the `+"`"+`--format`+"`"+` flag to return only the ID, status, and type of action: doctl compute droplet actions 386734086 --format ID,Status,Type`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet backup-policies
					ox.Usage(`backup-policies`, `Display commands for Droplet's backup policies.`),
					ox.Banner(`The commands under `+"`"+`doctl compute droplet backup-policies`+"`"+` are for displaying the commands for Droplet's backup policies.`),
					ox.Spec(`[command]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl compute droplet backup-policies [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl compute droplet backup-policies get
						ox.Usage(`get`, `Get droplet's backup policy`),
						ox.Banner(`Retrieves a backup policy of a Droplet.`),
						ox.Spec(`<droplet-id> [flags]`),
						ox.Example(`
The following example retrieves a backup policy for a Droplet with the ID `+"`"+`386734086`+"`"+`. The command also uses the `+"`"+`--format`+"`"+` flag to only return the Droplet's id and backup policy plan: doctl compute droplet backup-policies get 386734086 --format DropletID,BackupPolicyPlan`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: DropletID, `+"`"+`BackupEnabled`+"`"+`, `+"`"+`BackupPolicyPlan`+"`"+`, `+"`"+`BackupPolicyWeekday`+"`"+`, `+"`"+`BackupPolicyHour`+"`"+`, `+"`"+`BackupPolicyWindowLengthHours`+"`"+`, `+"`"+`BackupPolicyRetentionPeriodDays`+"`"+`, `+"`"+`NextBackupWindowStart`+"`"+`, `+"`"+`NextBackupWindowEnd`+"`"+`.`, ox.Default("DropletID")).
							Bool(`no-header`, `Return raw data with no headers`).
							Bool(`template`, `Go template format. Sample values: `+"`"+`{{.DropletID}}`+"`"+`, `+"`"+`{{.BackupEnabled}}`+"`"+`, `+"`"+`{{.BackupPolicy.Plan}}`+"`"+`, `+"`"+`{{.BackupPolicy.Weekday}}`+"`"+`, `+"`"+`{{.BackupPolicy.Hour}}`+"`"+`, `+"`"+`{{.BackupPolicy.Plan}}`+"`"+`, `+"`"+`{{.BackupPolicy.WindowLengthHours}}`+"`"+`, `+"`"+`{{.BackupPolicy.RetentionPeriodDays}}`+"`"+`, `+"`"+`{{.NextBackupWindow.Start}}`+"`"+`, `+"`"+`{{.NextBackupWindow.End}}`+"`"+``).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl compute droplet backup-policies list
						ox.Usage(`list`, `List backup policies for all Droplets`),
						ox.Banner(`List droplet backup policies for all existing Droplets.`),
						ox.Spec(`[flags]`),
						ox.Aliases("ls"),
						ox.Example(`
The following example list backup policies for all existing Droplets: doctl compute droplet backup-policies list`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl compute droplet backup-policies list-supported
						ox.Usage(`list-supported`, `List of all supported droplet backup policies`),
						ox.Banner(`List of all supported droplet backup policies.`),
						ox.Spec(`[flags]`),
						ox.Example(`
The following example list all supported backup policies for Droplets: doctl compute droplet backup-policies list-supported`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl compute droplet backups
					ox.Usage(`backups`, `List Droplet backups`),
					ox.Banner(`Lists backup images for a Droplet, including each image's slug and ID.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("b"),
					ox.Example(`
The following example retrieves a list of backups for a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet backups 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet create
					ox.Usage(`create`, `Create a new Droplet`),
					ox.Banner(`Creates a new Droplet on your account. The command requires values for the `+"`"+`--size`+"`"+`, and `+"`"+`--image`+"`"+` flags.

To retrieve a list of size slugs, use the `+"`"+`doctl compute size list`+"`"+` command. To retrieve a list of image slugs, use the `+"`"+`doctl compute image list`+"`"+` command.

If you do not specify a region, the Droplet is created in the default region for your account. If you do not specify any SSH keys, we email a temporary password to your account's email address.`),
					ox.Spec(`<droplet-name>... [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a Droplet named `+"`"+`example-droplet`+"`"+` with a two vCPUs, two GiB of RAM, and 20 GBs of disk space. The Droplet is created in the `+"`"+`nyc1`+"`"+` region and is based on the `+"`"+`ubuntu-20-04-x64`+"`"+` image. Additionally, the command uses the `+"`"+`--user-data`+"`"+` flag to run a Bash script the first time the Droplet boots up: doctl compute droplet create example-droplet --size s-2vcpu-2gb --image ubuntu-20-04-x64 --region nyc1 --user-data $'#!/bin/bash\n touch /root/example.txt; sudo apt update;sudo snap install doctl'`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`backup-policy-hour`, `Backup policy hour.`).
						String(`backup-policy-plan`, `Backup policy frequency plan.`).
						String(`backup-policy-weekday`, `Backup policy weekday.`).
						Bool(`droplet-agent`, `Specifies whether or not the Droplet monitoring agent should be installed. By default, the agent is installed on new Droplets but installation errors are ignored. Set --droplet-agent=false to prevent installation. Set to `+"`"+`true`+"`"+` to make installation errors fatal.`).
						Bool(`enable-backups`, `Enables backups for the Droplet. By default, backups are created on a daily basis.`).
						Bool(`enable-ipv6`, `Enables IPv6 support and assigns an IPv6 address to the Droplet`).
						Bool(`enable-monitoring`, `Installs the DigitalOcean agent for additional monitoring`).
						Bool(`enable-private-networking`, `Enables private networking for the Droplet by provisioning it inside of your account's default VPC for the region`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`PublicIPv4`+"`"+`, `+"`"+`PrivateIPv4`+"`"+`, `+"`"+`PublicIPv6`+"`"+`, `+"`"+`Memory`+"`"+`, `+"`"+`VCPUs`+"`"+`, `+"`"+`Disk`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Image`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Features`+"`"+`, `+"`"+`Volumes`+"`"+`.`, ox.Spec(`ID`)).
						String(`image`, `An ID or slug specifying the image to use to create the Droplet, such as ubuntu-20-04-x64. Use the commands under `+"`"+`doctl compute image`+"`"+` to find additional images. (required)`, ox.Default("ubuntu-20-04-x64")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`project-id`, `The UUID of the project to assign the Droplet to`).
						String(`region`, `A slug specifying the region to create the Droplet in, such as `+"`"+`nyc1`+"`"+`. Use the `+"`"+`doctl compute region list`+"`"+` command for a list of valid regions.`, ox.Default("slug")).
						String(`size`, `A slug indicating the Droplet's number of vCPUs, RAM, and disk size. For example, `+"`"+`s-1vcpu-1gb`+"`"+` specifies a Droplet with one vCPU and 1 GiB of RAM. The disk size is defined by the slug's plan. Run `+"`"+`doctl compute size list`+"`"+` for a list of valid size slugs and their disk sizes. (required)`, ox.Default("slug")).
						Slice(`ssh-keys`, `A list of SSH key IDs or fingerprints to embed in the Droplet's root account upon creation`).
						String(`tag-name`, `Applies a tag to the Droplet`).
						Slice(`tag-names`, `Applies a list of tags to the Droplet`).
						String(`user-data`, `A shell script to run on the Droplet's first boot`).
						String(`user-data-file`, `The path to a file containing a shell script or Cloud-init YAML file to run on the Droplet's first boot. Example: path/to/file.yaml`, ox.Spec(`path/to/file.yaml`)).
						Slice(`volumes`, `A list of block storage volume IDs to attach to the Droplet`).
						String(`vpc-uuid`, `The UUID of a non-default VPC to create the Droplet in`).
						Bool(`wait`, `Instructs the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet delete
					ox.Usage(`delete`, `Permanently delete a Droplet`),
					ox.Banner(`Permanently deletes a Droplet. This is irreversible.`),
					ox.Spec(`<droplet-id|droplet-name>... [flags]`),
					ox.Aliases("d", "del", "rm"),
					ox.Example(`
The following example deletes a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet delete 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the Droplet without a confirmation prompt`, ox.Short("f")).
						String(`tag-name`, `Tag name`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet get
					ox.Usage(`get`, `Retrieve information about a Droplet`),
					ox.Banner(`Retrieves information about a Droplet, including:

- The Droplet's ID
- The Droplet's name
- The Droplet's public IPv4 address
- The Droplet's private IPv4 address
- The Droplet's IPv6 address
- The memory size of the Droplet in MB
- The number of vCPUs on the Droplet
- The size of the Droplet's disk in GB
- The Droplet's data center region
- The image the Droplet was created from
- The status of the Droplet. Possible values: `+"`"+`new`+"`"+`, `+"`"+`active`+"`"+`, `+"`"+`off`+"`"+`, or `+"`"+`archive`+"`"+`
- The tags assigned to the Droplet
- A list of features enabled for the Droplet, such as `+"`"+`backups`+"`"+`, `+"`"+`ipv6`+"`"+`, `+"`"+`monitoring`+"`"+`, and `+"`"+`private_networking`+"`"+`
- The IDs of block storage volumes attached to the Droplet`),
					ox.Spec(`<droplet-id|droplet-name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about a Droplet with the ID `+"`"+`386734086`+"`"+`. The command also uses the `+"`"+`--format`+"`"+` flag to only return the Droplet's name, ID, and public IPv4 address: doctl compute droplet get 386734086 --format Name,ID,PublicIPv4`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`PublicIPv4`+"`"+`, `+"`"+`PrivateIPv4`+"`"+`, `+"`"+`PublicIPv6`+"`"+`, `+"`"+`Memory`+"`"+`, `+"`"+`VCPUs`+"`"+`, `+"`"+`Disk`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Image`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Features`+"`"+`, `+"`"+`Volumes`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`template`, `Go template format. Sample values: {{.ID}}, `+"`"+`{{.Name}}`+"`"+`, `+"`"+`{{.Memory}}`+"`"+`, `+"`"+`{{.Region.Name}}`+"`"+`, `+"`"+`{{.Image}}`+"`"+`, `+"`"+`{{.Tags}}`+"`"+``, ox.Default("{{.ID}}")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet kernels
					ox.Usage(`kernels`, `List available Droplet kernels`),
					ox.Banner(`Retrieves a list of all kernels available to a Droplet. This command is only available for Droplets with externally managed kernels. All Droplets created after March 2017 have internally managed kernels by default.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("k"),
					ox.Example(`
The following example retrieves a list of available kernels for a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet kernels 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Version`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet list
					ox.Usage(`list`, `List Droplets on your account`),
					ox.Banner(`Retrieves a list of Droplets on your account, including the following information about each:

- The Droplet's ID
- The Droplet's name
- The Droplet's public IPv4 address
- The Droplet's private IPv4 address
- The Droplet's IPv6 address
- The memory size of the Droplet in MB
- The number of vCPUs on the Droplet
- The size of the Droplet's disk in GB
- The Droplet's data center region
- The image the Droplet was created from
- The status of the Droplet. Possible values: `+"`"+`new`+"`"+`, `+"`"+`active`+"`"+`, `+"`"+`off`+"`"+`, or `+"`"+`archive`+"`"+`
- The tags assigned to the Droplet
- A list of features enabled for the Droplet, such as `+"`"+`backups`+"`"+`, `+"`"+`ipv6`+"`"+`, `+"`"+`monitoring`+"`"+`, and `+"`"+`private_networking`+"`"+`
- The IDs of block storage volumes attached to the Droplet`),
					ox.Spec(`[GLOB] [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of all Droplets in the `+"`"+`nyc1`+"`"+` region: doctl compute droplet list --region nyc1`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`PublicIPv4`+"`"+`, `+"`"+`PrivateIPv4`+"`"+`, `+"`"+`PublicIPv6`+"`"+`, `+"`"+`Memory`+"`"+`, `+"`"+`VCPUs`+"`"+`, `+"`"+`Disk`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Image`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Features`+"`"+`, `+"`"+`Volumes`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`gpus`, `List GPU Droplets only. By default, only non-GPU Droplets are returned.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `Retrieves a list of Droplets in a specified region`).
						String(`tag-name`, `Retrieves a list of Droplets with the specified tag name`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet neighbors
					ox.Usage(`neighbors`, `List a Droplet's neighbors on your account`),
					ox.Banner(`Lists your Droplets that are on the same physical hardware, including the following details:

- The Droplet's ID
- The Droplet's name
- The Droplet's public IPv4 address
- The Droplet's private IPv4 address
- The Droplet's IPv6 address
- The memory size of the Droplet in MB
- The number of vCPUs on the Droplet
- The size of the Droplet's disk in GB
- The Droplet's data center region
- The image the Droplet was created from
- The status of the Droplet. Possible values: `+"`"+`new`+"`"+`, `+"`"+`active`+"`"+`, `+"`"+`off`+"`"+`, or `+"`"+`archive`+"`"+`
- The tags assigned to the Droplet
- A list of features enabled for the Droplet, such as `+"`"+`backups`+"`"+`, `+"`"+`ipv6`+"`"+`, `+"`"+`monitoring`+"`"+`, and `+"`"+`private_networking`+"`"+`
- The IDs of block storage volumes attached to the Droplet`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("n"),
					ox.Example(`
The following example retrieves a list of Droplets that are on the same physical hardware as the Droplet with the ID `+"`"+`386734086`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only each Droplet's ID, name and public IPv4 address: doctl compute droplet neighbors 386734086 --format ID,Name,PublicIPv4`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`PublicIPv4`+"`"+`, `+"`"+`PrivateIPv4`+"`"+`, `+"`"+`PublicIPv6`+"`"+`, `+"`"+`Memory`+"`"+`, `+"`"+`VCPUs`+"`"+`, `+"`"+`Disk`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Image`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Features`+"`"+`, `+"`"+`Volumes`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet snapshots
					ox.Usage(`snapshots`, `List all snapshots for a Droplet`),
					ox.Banner(`Retrieves a list of snapshots created from this Droplet.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("s"),
					ox.Example(`
The following example retrieves a list of snapshots for a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet snapshots 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet tag
					ox.Usage(`tag`, `Add a tag to a Droplet`),
					ox.Banner(`Applies a tag to a Droplet. Specify the tag with the `+"`"+`--tag-name`+"`"+` flag.`),
					ox.Spec(`<droplet-id|droplet-name> [flags]`),
					ox.Example(`
The following example applies the tag `+"`"+`frontend`+"`"+` to a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet tag 386734086 --tag-name frontend`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`tag-name`, `the tag name apply to the Droplet. You can use a new or existing tag. (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet untag
					ox.Usage(`untag`, `Remove a tag from a Droplet`),
					ox.Banner(`Removes a tag from a Droplet. Specify the tag with the `+"`"+`--tag-name`+"`"+` flag.`),
					ox.Spec(`<droplet-id|droplet-name> [flags]`),
					ox.Example(`
The following example removes the tag `+"`"+`frontend`+"`"+` from a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet untag 386734086 --tag-name frontend`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`tag-name`, `The tag name to remove from Droplet`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute droplet-action
				ox.Usage(`droplet-action`, `Display Droplet action commands`),
				ox.Banner(`Use the subcommands of `+"`"+`doctl compute droplet-action`+"`"+` to perform actions on Droplets.

You can use Droplet actions to perform tasks on a Droplet, such as rebooting, resizing, or snapshotting it.`),
				ox.Spec(`[command]`),
				ox.Aliases("da"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute droplet-action [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute droplet-action change-backup-policy
					ox.Usage(`change-backup-policy`, `Change backup policy on a Droplet`),
					ox.Banner(`Changes backup policy for a Droplet with enabled backups.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example changes backup policy on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action change-backup-policy 386734086 --backup-policy-plan weekly --backup-policy-weekday SUN --backup-policy-hour 4`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`backup-policy-hour`, `Backup policy hour.`).
						String(`backup-policy-plan`, `Backup policy frequency plan. (required)`).
						String(`backup-policy-weekday`, `Backup policy weekday.`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Wait for action to complete`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action change-kernel
					ox.Usage(`change-kernel`, `Change a Droplet's kernel`),
					ox.Banner(`Changes a Droplet's kernel. This is only available for externally managed kernels. All Droplets created after 17 March 2017 have internally managed kernels by default.
		
Use the `+"`"+`doctl compute droplet kernels <droplet-id>`+"`"+` command to retrieve a list of kernels for the Droplet.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Int(`kernel-id`, `Kernel ID (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action disable-backups
					ox.Usage(`disable-backups`, `Disable backups on a Droplet`),
					ox.Banner(`Disables backups on a Droplet. This does not delete existing backups.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example disables backups on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action disable-backups 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action enable-backups
					ox.Usage(`enable-backups`, `Enable backups on a Droplet`),
					ox.Banner(`Enables backups on a Droplet. This automatically creates and stores a disk image of the Droplet. By default, backups happen daily.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example enables backups on a Droplet with the ID `+"`"+`386734086`+"`"+` with a backup policy flag: doctl compute droplet-action enable-backups 386734086 --backup-policy-plan weekly --backup-policy-weekday SUN --backup-policy-hour 4`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`backup-policy-hour`, `Backup policy hour.`).
						String(`backup-policy-plan`, `Backup policy frequency plan.`).
						String(`backup-policy-weekday`, `Backup policy weekday.`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Wait for action to complete`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action enable-ipv6
					ox.Usage(`enable-ipv6`, `Enable IPv6 on a Droplet`),
					ox.Banner(`Enables IPv6 networking on a Droplet. When executed, we automatically assign an IPv6 address to the Droplet. 

The Droplet may require additional network configuration to properly use the new IPv6 address. For more information, see: https://docs.digitalocean.com/products/networking/ipv6/how-to/enable`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example enables IPv6 on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action enable-ipv6 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action enable-private-networking
					ox.Usage(`enable-private-networking`, `Enable private networking on a Droplet`),
					ox.Banner(`Enables VPC networking on a Droplet. This command adds a private IPv4 address to the Droplet that other resources inside the Droplet's VPC network can access. The Droplet is placed in the default VPC network for the region it resides in.

All Droplets created after 1 October 2020 are provided a private IP address and placed into a VPC network by default. You can use this command to enable private networking on a Droplet that was created before 1 October 2020 and was not already in a VPC network.

Once you have manually enabled private networking for a Droplet, the Droplet requires additional internal network configuration for it to become accessible through the VPC network. For more information, see: https://docs.digitalocean.com/products/networking/vpc/how-to/enable`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example enables private networking on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action enable-private-networking 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action get
					ox.Usage(`get`, `Retrieve a specific Droplet action`),
					ox.Banner(`Retrieves information about an action performed on a Droplet, including its status, type, and completion time.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about an action, with the ID `+"`"+`1978716488`+"`"+`, performed on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action get 1978716488 --action-id 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`action-id`, `Action ID (required)`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action password-reset
					ox.Usage(`password-reset`, `Reset the root password for a Droplet`),
					ox.Banner(`Initiates a root password reset on a Droplet. We provide a new password for the Droplet via the accounts email address. The password must be changed after first use. 

This also powercycles the Droplet.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example resets the root password for a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action password-reset 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action power-cycle
					ox.Usage(`power-cycle`, `Powercycle a Droplet`),
					ox.Banner(`Powercycles a Droplet. A powercycle action is similar to pushing the reset button on a physical machine.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example powercycles a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action power-cycle 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action power-off
					ox.Usage(`power-off`, `Power off a Droplet`),
					ox.Banner(`Use this command to power off a Droplet.
		
A `+"`"+`power_off`+"`"+` event is a hard shutdown and should only be used if the shutdown action is not successful. It is similar to cutting the power on a server and could lead to complications.

Droplets that are powered off are still billable. To stop incurring charges on a Droplet, destroy it.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example powers off a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action power-off 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action power-on
					ox.Usage(`power-on`, `Power on a Droplet`),
					ox.Banner(`Powers on a Droplet. This is similar to pressing the power button on a physical machine.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example powers on a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action power-on 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action reboot
					ox.Usage(`reboot`, `Reboot a Droplet`),
					ox.Banner(`Reboots a Droplet. A reboot action is an attempt to reboot the Droplet in a graceful way, similar to using the reboot command from the Droplet's console.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example reboots a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action reboot 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action rebuild
					ox.Usage(`rebuild`, `Rebuild a Droplet`),
					ox.Banner(`Rebuilds a Droplet from an image, such as an Ubuntu base image or a backup image of the Droplet. Set the image attribute to an image ID or slug.

To retrieve a list of images on your account, use the `+"`"+`doctl compute image list`+"`"+` command. To retrieve a list of base images, use the `+"`"+`doctl compute image list-distribution`+"`"+` command.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example rebuilds a Droplet with the ID `+"`"+`386734086`+"`"+` from the image with the ID `+"`"+`146288445`+"`"+`: doctl compute droplet-action rebuild 386734086 --image 146288445`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						String(`image`, `An image ID or slug (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action rename
					ox.Usage(`rename`, `Rename a Droplet`),
					ox.Banner(`Renames a Droplet. When using a Fully Qualified Domain Name (FQDN) this also updates the Droplet's pointer (PTR) record.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example renames a Droplet with the ID `+"`"+`386734086`+"`"+` to `+"`"+`example.com`+"`"+` an FQDN: doctl compute droplet-action rename 386734086 --droplet-name example.com`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-name`, `The new name for the Droplet (required)`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action resize
					ox.Usage(`resize`, `Resize a Droplet`),
					ox.Banner(`Resizes a Droplet to a different plan.

By default, this command only increases or decreases the CPU and RAM of the Droplet, not its disk size. Unlike increasing disk size, you can reverse this action.

To also increase the Droplet's disk size, choose a size slug with the desired amount of vCPUs, RAM, and disk space and then set the `+"`"+`--resize-disk`+"`"+` flag to `+"`"+`true`+"`"+`. This is a permanent change and cannot be reversed as a Droplet's disk size cannot be decreased.

For a list of size slugs, use the `+"`"+`doctl compute size list`+"`"+` command.

This command automatically powers off the Droplet before resizing it.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example resizes a Droplet with the ID `+"`"+`386734086`+"`"+` to a Droplet with two CPUs, two GiB of RAM, and 60 GBs of disk space. The 60 GBs of disk space is the defined amount for the `+"`"+`s-2vcpu-2gb`+"`"+` plan: doctl compute droplet-action resize 386734086 --size s-2vcpu-2gb --resize-disk=true`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`resize-disk`, `Resize the Droplet's disk size in addition to its RAM and CPUs`).
						String(`size`, `A slug indicating the new size for the Droplet, for example `+"`"+`s-2vcpu-2gb`+"`"+`. Run `+"`"+`doctl compute size list`+"`"+` for a list of valid sizes. (required)`, ox.Default("slug")).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action restore
					ox.Usage(`restore`, `Restore a Droplet from a backup`),
					ox.Banner(`Restores a Droplet from a backup image. You must pass an image ID that is a backup of the current Droplet instance. The operation leaves any embedded SSH keys intact.
		
		To retrieve a list of backup images, use the `+"`"+`doctl compute image list`+"`"+` command.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example restores a Droplet with the ID `+"`"+`386734086`+"`"+` from a backup image with the ID `+"`"+`146288445`+"`"+`: doctl compute droplet-action restore 386734086 --image-id 146288445`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Int(`image-id`, `The ID of the image to restore the Droplet from (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action shutdown
					ox.Usage(`shutdown`, `Shut down a Droplet`),
					ox.Banner(`Shuts down a Droplet. 
		
A shutdown action is an attempt to shutdown the Droplet in a graceful way, similar to using the shutdown command from the Droplet's console. Since a shutdown command can fail, this action guarantees that the command is issued, not that it succeeds. The preferred way to turn off a Droplet is to attempt a shutdown, with a reasonable timeout, followed by a `+"`"+`doctl compute droplet-action power_off`+"`"+` action to ensure the Droplet is off.
		
Droplets that are powered off are still billable. To stop incurring charges on a Droplet, destroy it.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Example(`
The following example shuts down a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute droplet-action shutdown 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute droplet-action snapshot
					ox.Usage(`snapshot`, `Take a Droplet snapshot`),
					ox.Banner(`Takes a snapshot of a Droplet. Snapshots are complete disk images that contain all of the data on a Droplet at the time of the snapshot. This can be useful for restoring and rebuilding Droplets.
		
We recommend that you power off the Droplet before taking a snapshot to ensure data consistency.`),
					ox.Spec(`<droplet-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`snapshot-name`, `The snapshot's name (required)`).
						Bool(`wait`, `Instruct the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute firewall
				ox.Usage(`firewall`, `Display commands to manage cloud firewalls`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute firewall`+"`"+` manage DigitalOcean cloud firewalls.

Cloud firewalls allow you to restrict network access to and from a Droplet by defining which ports accept inbound or outbound connections. With these commands, you can list, create, or delete Cloud firewalls, as well as modify access rules. 

Note: Cloud firewalls are not internal Droplet firewalls on Droplets, such as UFW or FirewallD.

A firewall's `+"`"+`inbound_rules`+"`"+` and `+"`"+`outbound_rules`+"`"+` attributes contain arrays of objects as their values. These objects contain the standard attributes of their associated types, which can be found below.

Inbound access rules specify the protocol (TCP, UDP, or ICMP), ports, and sources for inbound traffic that will be allowed through the Firewall to the target Droplets. The `+"`"+`ports`+"`"+` attribute may contain a single port, a range of ports (e.g. `+"`"+`8000-9000`+"`"+`), or `+"`"+`all`+"`"+` to allow traffic on all ports for the specified protocol. The `+"`"+`sources`+"`"+` attribute will contain an object specifying a whitelist of sources from which traffic will be accepted.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute firewall [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute firewall add-droplets
					ox.Usage(`add-droplets`, `Add Droplets to a cloud firewall`),
					ox.Banner(`Assigns Droplets to a cloud firewall on your account.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example assigns two Droplets to the cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall add-droplets f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --droplet-ids "386734086,391669331"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of Droplet IDs to place behind the cloud firewall, for example: 386734086,391669331`, ox.Spec(`386734086,391669331`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall add-rules
					ox.Usage(`add-rules`, `Add inbound or outbound rules to a cloud firewall`),
					ox.Banner(`Add inbound or outbound rules to a cloud firewall.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example adds an inbound rule and an outbound rule to a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall add-rules f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --inbound-rules "protocol:tcp,ports:22,droplet_id:386734086" --outbound-rules "protocol:tcp,ports:22,address:0.0.0.0/0"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`inbound-rules`, `A comma-separated key-value list that defines an inbound rule. The rule must define a communication protocol, a port number, and a traffic source location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that resources can only receive TCP traffic on port 22 from addresses in the specified CIDR: protocol:tcp,ports:22,address:192.0.2.0/24.`, ox.Default("protocol:tcp,ports:22,address:192.0.2.0/24")).
						String(`outbound-rules`, `A comma-separate key-value list that defines an outbound rule. The rule must define a communication protocol, a port number, and a destination location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that the firewall only allows traffic to be sent to port 22 of any IPv4 address on the internet: protocol:tcp,ports:22,address:0.0.0.0/0.`, ox.Default("protocol:tcp,ports:22,address:0.0.0.0/0")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall add-tags
					ox.Usage(`add-tags`, `Add tags to a cloud firewall`),
					ox.Banner(`Add tags to a cloud firewall. This adds all assets using that tag to the firewall.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example adds two tags to a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall add-tags f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --tag-names "frontend,backend"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`tag-names`, `A comma-separated list of existing tags, for example: frontend,backend,env:prod. Droplets with these tags will be placed behind the cloud firewall`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall create
					ox.Usage(`create`, `Create a new cloud firewall`),
					ox.Banner(`Creates a cloud firewall. This command must contain at least one inbound or outbound access rule.`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a cloud firewall named `+"`"+`example-firewall`+"`"+` that contains an inbound rule and an outbound rule and applies them to the specified Droplets: doctl compute firewall create --name "example-firewall" --inbound-rules "protocol:tcp,ports:22,droplet_id:386734086" --outbound-rules "protocol:tcp,ports:22,address:0.0.0.0/0" --droplet-ids "386734086,391669331" --tag-names "frontend,backend,k8s:f81d4fae-7dec-11d0-a765-00a0c91e6bf6"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of Droplet IDs to place behind the cloud firewall, for example: 386734086,391669331`, ox.Spec(`386734086,391669331`)).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`InboundRules`+"`"+`, `+"`"+`OutboundRules`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`PendingChanges`+"`"+`.`, ox.Spec(`ID`)).
						String(`inbound-rules`, `A comma-separated key-value list that defines an inbound rule. The rule must define a communication protocol, a port number, and a traffic source location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that resources can only receive TCP traffic on port 22 from addresses in the specified CIDR: protocol:tcp,ports:22,address:192.0.2.0/24.`, ox.Default("protocol:tcp,ports:22,address:192.0.2.0/24")).
						String(`name`, `The firewall's name (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`outbound-rules`, `A comma-separate key-value list that defines an outbound rule. The rule must define a communication protocol, a port number, and a destination location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that the firewall only allows traffic to be sent to port 22 of any IPv4 address on the internet: protocol:tcp,ports:22,address:0.0.0.0/0.`, ox.Default("protocol:tcp,ports:22,address:0.0.0.0/0")).
						Slice(`tag-names`, `A comma-separated list of existing tags, for example: frontend,backend,env:prod. Droplets with these tags will be placed behind the cloud firewall`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall delete
					ox.Usage(`delete`, `Permanently delete a cloud firewall`),
					ox.Banner(`Permanently deletes a cloud firewall. This is irreversible, but does not delete any Droplets assigned to the cloud firewall.`),
					ox.Spec(`<firewall-id>... [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the firewall without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall get
					ox.Usage(`get`, `Retrieve information about a cloud firewall`),
					ox.Banner(`Retrieves information about an existing cloud firewall, including:

- The firewall's UUID
- The firewall's name
- The status of the firewall. Possible values: `+"`"+`waiting`+"`"+`, `+"`"+`succeeded`+"`"+`, `+"`"+`failed`+"`"+`.
- The firewall's creation date, in ISO8601 combined date and time format.
- Any pending changes to the firewall. Possible values:`+"`"+`droplet_id`+"`"+`, `+"`"+`removing`+"`"+`, `+"`"+`status`+"`"+`.
	  When empty, all changes have been successfully applied.
- The inbound rules for the firewall
- The outbound rules for the firewall
- The IDs of Droplets assigned to the firewall
- The tags of Droplets assigned to the firewall`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about the cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`InboundRules`+"`"+`, `+"`"+`OutboundRules`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`PendingChanges`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall list
					ox.Usage(`list`, `List the cloud firewalls on your account`),
					ox.Banner(`Retrieves a list of cloud firewalls on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all cloud firewalls on your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, name and inbound rules for each firewall: doctl compute firewall list --format ID,Name,InboundRules`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`InboundRules`+"`"+`, `+"`"+`OutboundRules`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`PendingChanges`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall list-by-droplet
					ox.Usage(`list-by-droplet`, `List firewalls by Droplet`),
					ox.Banner(`Lists the cloud firewalls assigned to a Droplet.`),
					ox.Spec(`<droplet_id> [flags]`),
					ox.Example(`
The following example lists all cloud firewalls assigned to the Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute firewall list-by-droplet 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`InboundRules`+"`"+`, `+"`"+`OutboundRules`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`PendingChanges`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall remove-droplets
					ox.Usage(`remove-droplets`, `Remove Droplets from a cloud firewall`),
					ox.Banner(`Removes Droplets from a cloud firewall.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example removes two Droplets from a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall remove-droplets f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --droplet-ids "386734086,391669331"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of Droplet IDs to place behind the cloud firewall, for example: 386734086,391669331`, ox.Spec(`386734086,391669331`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall remove-rules
					ox.Usage(`remove-rules`, `Remove inbound or outbound rules from a cloud firewall`),
					ox.Banner(`Remove inbound or outbound rules from a cloud firewall.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example removes an inbound rule and an outbound rule from a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall remove-rules f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --inbound-rules "protocol:tcp,ports:22,droplet_id:386734086" --outbound-rules "protocol:tcp,ports:22,address:0.0.0.0/0"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`inbound-rules`, `A comma-separated key-value list that defines an inbound rule. The rule must define a communication protocol, a port number, and a traffic source location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that resources can only receive TCP traffic on port 22 from addresses in the specified CIDR: protocol:tcp,ports:22,address:192.0.2.0/24.`, ox.Default("protocol:tcp,ports:22,address:192.0.2.0/24")).
						String(`outbound-rules`, `A comma-separate key-value list that defines an outbound rule. The rule must define a communication protocol, a port number, and a destination location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that the firewall only allows traffic to be sent to port 22 of any IPv4 address on the internet: protocol:tcp,ports:22,address:0.0.0.0/0.`, ox.Default("protocol:tcp,ports:22,address:0.0.0.0/0")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall remove-tags
					ox.Usage(`remove-tags`, `Remove tags from a cloud firewall`),
					ox.Banner(`Removes tags from a cloud firewall. This removes all assets using that tag from the firewall.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Example(`
The following example removes two tags from a cloud firewall with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute firewall remove-tags f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --tag-names "frontend,backend"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`tag-names`, `A comma-separated list of existing tags, for example: frontend,backend,env:prod. Droplets with these tags will be placed behind the cloud firewall`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute firewall update
					ox.Usage(`update`, `Update a cloud firewall's configuration`),
					ox.Banner(`Updates the configuration of an existing cloud firewall. The request should contain a full representation of the firewall, including existing attributes. Any attributes that are not provided are reset to their default values.`),
					ox.Spec(`<firewall-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates a cloud firewall named `+"`"+`example-firewall`+"`"+` that contains an inbound rule and an outbound rule and applies them to the specified Droplet: doctl compute firewall update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name "example-firewall" --inbound-rules "protocol:tcp,ports:22,droplet_id:386734086" --outbound-rules "protocol:tcp,ports:22,address:0.0.0.0/0" --droplet-ids "386734086,391669331"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of Droplet IDs to place behind the cloud firewall, for example: 386734086,391669331`, ox.Spec(`386734086,391669331`)).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`InboundRules`+"`"+`, `+"`"+`OutboundRules`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`PendingChanges`+"`"+`.`, ox.Spec(`ID`)).
						String(`inbound-rules`, `A comma-separated key-value list that defines an inbound rule. The rule must define a communication protocol, a port number, and a traffic source location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that resources can only receive TCP traffic on port 22 from addresses in the specified CIDR: protocol:tcp,ports:22,address:192.0.2.0/24.`, ox.Default("protocol:tcp,ports:22,address:192.0.2.0/24")).
						String(`name`, `The firewall's name (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`outbound-rules`, `A comma-separate key-value list that defines an outbound rule. The rule must define a communication protocol, a port number, and a destination location, such as a Droplet ID, IP address, or a tag. For example, the following rule defines that the firewall only allows traffic to be sent to port 22 of any IPv4 address on the internet: protocol:tcp,ports:22,address:0.0.0.0/0.`, ox.Default("protocol:tcp,ports:22,address:0.0.0.0/0")).
						Slice(`tag-names`, `A comma-separated list of existing tags, for example: frontend,backend,env:prod. Droplets with these tags will be placed behind the cloud firewall`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute image
				ox.Usage(`image`, `Display commands to manage images`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute image`+"`"+` manage images. A DigitalOcean image can be used to create a Droplet.

Currently, there are five types of images: snapshots, backups, custom images, distributions, and one-click application.

- Snapshots provide a full copy of an existing Droplet instance taken on demand.
- Backups are similar to snapshots but are created automatically at regular intervals when enabled for a Droplet.
- Custom images are Linux-based virtual machine images that you may upload for use on DigitalOcean. We support the following formats: raw, qcow2, vhdx, vdi, or vmdk.
- Distributions are the public Linux distributions that are available to be used as a base to create Droplets.
- Applications, or one-click apps, are distributions pre-configured with additional software, such as WordPress, Django, or Flask.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute image [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute image create
					ox.Usage(`create`, `Create custom image`),
					ox.Banner(`Creates an image in your DigitalOcean account. Specify a URL to download the image from and the region to store the image in. You can add additional metadata to the image using the optional flags.`),
					ox.Spec(`<image-name> [flags]`),
					ox.Example(`
The following example creates a custom image named `+"`"+`Example Image`+"`"+` from a URL and stores it in the `+"`"+`nyc1`+"`"+` region: doctl compute image create "Example Image" --image-url "https://example.com/image.iso" --region nyc1`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`image-description`, `An optional description of the image`).
						String(`image-distribution`, `A custom image distribution slug to apply to the image`, ox.Default("Unknown")).
						String(`image-url`, `The URL to retrieve the image from (required)`).
						String(`region`, `The slug of the region you want to store the image in. For a list of region slugs, use the `+"`"+`doctl compute region list`+"`"+` command. (required)`, ox.Default("slug")).
						Slice(`tag-names`, `A list of tag names to apply to the image`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image delete
					ox.Usage(`delete`, `Permanently delete an image from your account`),
					ox.Banner(`Permanently deletes an image from your account. This is irreversible.`),
					ox.Spec(`<image-id> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes an image with the ID `+"`"+`386734086`+"`"+`: doctl compute image delete 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Force image delete`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image get
					ox.Usage(`get`, `Retrieve information about an image`),
					ox.Banner(`Returns the following information about the specified image:

- The image's ID
- The image's name
- The type of image. Possible values: `+"`"+`snapshot`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`custom`+"`"+`.
- The distribution of the image. For custom images, this is user defined.
- The image's slug. This is a unique string that identifies each DigitalOcean-provided public image. These can be used to reference a public image as an alternative to the numeric ID.
- Whether the image is public or not. An public image is available to all accounts. A private image is only accessible from your account. This is boolean value, true or false.
- The minimum Droplet disk size required for a Droplet to use this image, in GB.`),
					ox.Spec(`<image-id|image-slug> [flags]`),
					ox.Example(`
The following example retrieves information about an image with the ID `+"`"+`386734086`+"`"+`: doctl compute image get 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image list
					ox.Usage(`list`, `List images on your account`),
					ox.Banner(`Lists all private images on your account. To list public images, use the `+"`"+`--public`+"`"+` flag. This command returns the following information about each image:

- The image's ID
- The image's name
- The type of image. Possible values: `+"`"+`snapshot`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`custom`+"`"+`.
- The distribution of the image. For custom images, this is user defined.
- The image's slug. This is a unique string that identifies each DigitalOcean-provided public image. These can be used to reference a public image as an alternative to the numeric ID.
- Whether the image is public or not. An public image is available to all accounts. A private image is only accessible from your account. This is boolean value, true or false.
- The minimum Droplet disk size required for a Droplet to use this image, in GB.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all private images on your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, distribution, and slug for each image: doctl compute image list --format ID,Distribution,Slug`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`public`, `Lists public images`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image list-application
					ox.Usage(`list-application`, `List available One-Click Apps`),
					ox.Banner(`Lists all public one-click apps that are currently available on the DigitalOcean Marketplace. This command returns the following information about each image:

- The image's ID
- The image's name
- The type of image. Possible values: `+"`"+`snapshot`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`custom`+"`"+`.
- The distribution of the image. For custom images, this is user defined.
- The image's slug. This is a unique string that identifies each DigitalOcean-provided public image. These can be used to reference a public image as an alternative to the numeric ID.
- Whether the image is public or not. An public image is available to all accounts. A private image is only accessible from your account. This is boolean value, true or false.
- The minimum Droplet disk size required for a Droplet to use this image, in GB.`),
					ox.Spec(`[flags]`),
					ox.Example(`
The following example lists all public One-Click Apps available from DigitalOcean and uses the `+"`"+`--format`+"`"+` flag to return only the ID, name, distribution, and slug for each image: doctl compute image list-application --format ID,Name,Distribution,Slug`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`public`, `Lists public images`, ox.Default("true")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image list-distribution
					ox.Usage(`list-distribution`, `List available distribution images`),
					ox.Banner(`Lists the distribution images available from DigitalOcean. This command returns the following information about each image:

- The image's ID
- The image's name
- The type of image. Possible values: `+"`"+`snapshot`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`custom`+"`"+`.
- The distribution of the image. For custom images, this is user defined.
- The image's slug. This is a unique string that identifies each DigitalOcean-provided public image. These can be used to reference a public image as an alternative to the numeric ID.
- Whether the image is public or not. An public image is available to all accounts. A private image is only accessible from your account. This is boolean value, true or false.
- The minimum Droplet disk size required for a Droplet to use this image, in GB.`),
					ox.Spec(`[flags]`),
					ox.Example(`
The following example lists all public distribution images available from DigitalOcean and uses the `+"`"+`--format`+"`"+` flag to return only the ID, distribution, and slug for each image: doctl compute image list-distribution --format ID,Distribution,Slug`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`public`, `Lists public images`, ox.Default("true")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image list-user
					ox.Usage(`list-user`, `List user-created images`),
					ox.Banner(`Use this command to list user-created images, such as snapshots or custom images that you have uploaded to your account. This command returns the following information about each image:

- The image's ID
- The image's name
- The type of image. Possible values: `+"`"+`snapshot`+"`"+`, `+"`"+`backup`+"`"+`, `+"`"+`custom`+"`"+`.
- The distribution of the image. For custom images, this is user defined.
- The image's slug. This is a unique string that identifies each DigitalOcean-provided public image. These can be used to reference a public image as an alternative to the numeric ID.
- Whether the image is public or not. An public image is available to all accounts. A private image is only accessible from your account. This is boolean value, true or false.
- The minimum Droplet disk size required for a Droplet to use this image, in GB.`),
					ox.Spec(`[flags]`),
					ox.Example(`
The following example lists all user-created images on your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, name, distribution, and slug for each image: doctl compute image list-user --format ID,Name,Distribution,Slug`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						Bool(`public`, `Lists public images`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image update
					ox.Usage(`update`, `Update an image's metadata`),
					ox.Banner(`Updates an image's metadata, including its name, description, and distribution.`),
					ox.Spec(`<image-id> [flags]`),
					ox.Example(`
The following example updates the name of an image with the ID `+"`"+`386734086`+"`"+` to `+"`"+`New Image Name`+"`"+`: doctl compute image update 386734086 --name "Example Image Name"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Distribution`+"`"+`, `+"`"+`Slug`+"`"+`, `+"`"+`Public`+"`"+`, `+"`"+`MinDisk`+"`"+`.`, ox.Spec(`ID`)).
						String(`image-name`, `The name of the image to update (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute image-action
				ox.Usage(`image-action`, `Display commands to perform actions on images`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute image-action`+"`"+` can be used to perform actions on images.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute image-action [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute image-action get
					ox.Usage(`get`, `Retrieve the status of an image action`),
					ox.Banner(`Retrieves the status of an image action, including the following details:

- The unique ID used to identify and reference an image action
- The status of the image action. Possible values: `+"`"+`in-progress`+"`"+`, `+"`"+`completed`+"`"+`, `+"`"+`errored`+"`"+`.
- When the action was initiated, in ISO8601 combined date and time format
- When the action was completed, in ISO8601 combined date and time format
- The ID of the resource that the action was taken on
- The type of resource that the action was taken on
- The region where the action occurred
- The region's slug`),
					ox.Spec(`<image-id> [flags]`),
					ox.Example(`
The following example retrieves the details for an image-action with ID 191669331 take on an image with the ID 386734086: doctl compute image-action get 386734086 --action-id 191669331`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`action-id`, `action id (required)`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute image-action transfer
					ox.Usage(`transfer`, `Transfer an image to another datacenter region`),
					ox.Banner(`Transfers an image to a different datacenter region. Also outputs the following details:

- The unique ID used to identify and reference an image action
- The status of the image action. Possible values: `+"`"+`in-progress`+"`"+`, `+"`"+`completed`+"`"+`, `+"`"+`errored`+"`"+`.
- When the action was initiated, in ISO8601 combined date and time format
- When the action was completed, in ISO8601 combined date and time format
- The ID of the resource that the action was taken on
- The type of resource that the action was taken on
- The region where the action occurred
- The region's slug`),
					ox.Spec(`<image-id> [flags]`),
					ox.Example(`
The following example transfers an image with the ID 386734086 to the region with the slug nyc3: doctl compute image-action transfer 386734086 --region nyc3`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `The target region to transfer the image to (required)`).
						Bool(`wait`, `Instructs the terminal to wait for the action to complete before returning access to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute load-balancer
				ox.Usage(`load-balancer`, `Display commands to manage load balancers`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute load-balancer`+"`"+` manage your load balancers.

With the load-balancer command, you can list, create, or delete load balancers, and manage their configuration details.`),
				ox.Spec(`[command]`),
				ox.Aliases("lb"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute load-balancer [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute load-balancer add-droplets
					ox.Usage(`add-droplets`, `Add Droplets to a load balancer`),
					ox.Banner(`Use this command to add Droplets to a load balancer.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of IDs of Droplet to add to the load balancer, example value: 12,33`, ox.Spec(`12,33`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer add-forwarding-rules
					ox.Usage(`add-forwarding-rules`, `Add forwarding rules to a load balancer`),
					ox.Banner(`Use this command to add forwarding rules to a load balancer, specified with the `+"`"+`--forwarding-rules`+"`"+` flag. Valid rules include:
- `+"`"+`entry_protocol`+"`"+`: The entry protocol used for traffic to the load balancer. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`entry_port`+"`"+`: The entry port used for incoming traffic to the load balancer.
- `+"`"+`target_protocol`+"`"+`: The target protocol used for traffic from the load balancer to the backend Droplets. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`target_port`+"`"+`: The target port used to send traffic from the load balancer to the backend Droplets.
- `+"`"+`certificate_id`+"`"+`: The ID of the TLS certificate used for SSL termination, if enabled. Can be obtained with `+"`"+`doctl certificate list`+"`"+`
- `+"`"+`tls_passthrough`+"`"+`: Whether SSL passthrough is enabled on the load balancer.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`forwarding-rules`, `A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306.`, ox.Spec(`entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer create
					ox.Usage(`create`, `Create a new load balancer`),
					ox.Banner(`Use this command to create a new load balancer on your account. Valid forwarding rules are:
- `+"`"+`entry_protocol`+"`"+`: The entry protocol used for traffic to the load balancer. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`entry_port`+"`"+`: The entry port used for incoming traffic to the load balancer.
- `+"`"+`target_protocol`+"`"+`: The target protocol used for traffic from the load balancer to the backend Droplets. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`target_port`+"`"+`: The target port used to send traffic from the load balancer to the backend Droplets.
- `+"`"+`certificate_id`+"`"+`: The ID of the TLS certificate used for SSL termination, if enabled. Can be obtained with `+"`"+`doctl certificate list`+"`"+`
- `+"`"+`tls_passthrough`+"`"+`: Whether SSL passthrough is enabled on the load balancer.`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`algorithm`, `This field has been deprecated. You can no longer specify an algorithm for load balancers.`, ox.Default("round_robin")).
						String(`allow-list`, `A comma-separated list of ALLOW rules for the load balancer, e.g.: ip:1.2.3.4,cidr:1.2.0.0/16`, ox.Spec(`ip:1.2.3.4,cidr:1.2.0.0/16`)).
						String(`deny-list`, `A comma-separated list of DENY rules for the load balancer, e.g.: ip:1.2.3.4,cidr:1.2.0.0/16`, ox.Spec(`ip:1.2.3.4,cidr:1.2.0.0/16`)).
						Bool(`disable-lets-encrypt-dns-records`, `disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer`).
						String(`domains`, `is_managed:true certificate_id:test-cert-id-1                                                                             A comma-separated list of domains required to ingress traffic to a global load balancer, e.g.: name:test-domain-1 is_managed:true certificate_id:test-cert-id-1`, ox.Spec(`name:test-domain-1`)).
						String(`droplet-ids`, `A comma-separated list of Droplet IDs to add to the load balancer, e.g.: 12,33`, ox.Spec(`12,33`)).
						Bool(`enable-backend-keepalive`, `enable keepalive connections to backend target droplets`).
						Bool(`enable-proxy-protocol`, `enable proxy protocol`).
						String(`forwarding-rules`, `A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306.`, ox.Spec(`entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`)).
						String(`glb-cdn-settings`, `CDN cache settings global load balancer, e.g.: is_enabled:true`, ox.Spec(`is_enabled:true`)).
						String(`glb-settings`, `Target protocol and port settings for ingressing traffic to a global load balancer, e.g.: target_protocol:http,target_port:80`, ox.Spec(`target_protocol:http,target_port:80`)).
						String(`health-check`, `A comma-separated list of key-value pairs representing recent health check results, e.g.: protocol:http,port:80,path:/index.html,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3`, ox.Spec(`protocol:http,port:80,path:/index.html,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3`)).
						Int(`http-idle-timeout-seconds`, `HTTP idle timeout that configures the idle timeout for http connections on the load balancer`).
						String(`name`, `The load balancer's name (required)`).
						String(`network`, `The type of network the load balancer is accessible from, e.g.: EXTERNAL or `+"`"+`INTERNAL`+"`"+``, ox.Default("EXTERNAL")).
						String(`project-id`, `Indicates which project to associate the Load Balancer with. If not specified, the Load Balancer will be placed in your default project.`).
						Bool(`redirect-http-to-https`, `Redirects HTTP requests to the load balancer on port 80 to HTTPS on port 443`).
						String(`region`, `The load balancer's region, e.g.: nyc1`, ox.Default("nyc1")).
						String(`size`, `The load balancer's size, e.g.: lb-small. Only one of size and size-unit should be used`, ox.Default("lb-small")).
						Int(`size-unit`, `The load balancer's size, e.g.: 1. Only one of size-unit and size should be used`).
						String(`sticky-sessions`, `cookie_name:DO-LB, cookie_ttl_seconds:5                                                                                A comma-separated list of key-value pairs representing a list of active sessions, e.g.: type:cookies, cookie_name:DO-LB, cookie_ttl_seconds:5`, ox.Spec(`type:cookies,`)).
						String(`tag-name`, `The name of a tag. All Droplets with this tag applied will be assigned to the load balancer.`).
						Slice(`target-lb-ids`, `A comma-separated list of Load Balancer IDs to add as target to the global load balancer`).
						String(`type`, `The type of load balancer, e.g.: REGIONAL or `+"`"+`GLOBAL`+"`"+``, ox.Default("REGIONAL")).
						String(`vpc-uuid`, `The UUID of the VPC to create the load balancer in`).
						Bool(`wait`, `Boolean that specifies whether to wait for a load balancer to complete before returning control to the terminal`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer delete
					ox.Usage(`delete`, `Permanently delete a load balancer`),
					ox.Banner(`Use this command to permanently delete the specified load balancer. This is irreversible.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the load balancer without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer get
					ox.Usage(`get`, `Retrieve a load balancer`),
					ox.Banner(`Use this command to retrieve information about a load balancer instance, including:

- The load balancer's ID
- The load balancer's name
- The load balancer's IP address
- The current state of the load balancer. This can be `+"`"+`new`+"`"+`, `+"`"+`active`+"`"+`, or `+"`"+`errored`+"`"+`.
- The load balancer's creation date, in ISO8601 combined date and time format.
- The load balancer's forwarding rules. See `+"`"+`doctl compute load-balancer add-forwarding-rules --help`+"`"+` for a list.
- The `+"`"+`health_check`+"`"+` settings for the load balancer.
- The `+"`"+`sticky_sessions`+"`"+` settings for the load balancer.
- The datacenter region the load balancer is located in.
- The Droplet tag corresponding to the Droplets assigned to the load balancer.
- The IDs of the Droplets assigned to the load balancer.
- Whether HTTP request to the load balancer on port 80 will be redirected to HTTPS on port 443.
- Whether the PROXY protocol is in use on the load balancer.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Aliases("g"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`IP`+"`"+`, `+"`"+`IPv6`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`SizeUnit`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Tag`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`RedirectHttpToHttps`+"`"+`, `+"`"+`StickySessions`+"`"+`, `+"`"+`HealthCheck`+"`"+`, `+"`"+`ForwardingRules`+"`"+`, `+"`"+`Firewall`+"`"+`, `+"`"+`DisableLetsEncryptDNSRecords`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer list
					ox.Usage(`list`, `List load balancers`),
					ox.Banner(`Use this command to get a list of the load balancers on your account, including the following information for each:

- The load balancer's ID
- The load balancer's name
- The load balancer's IP address
- The current state of the load balancer. This can be `+"`"+`new`+"`"+`, `+"`"+`active`+"`"+`, or `+"`"+`errored`+"`"+`.
- The load balancer's creation date, in ISO8601 combined date and time format.
- The load balancer's forwarding rules. See `+"`"+`doctl compute load-balancer add-forwarding-rules --help`+"`"+` for a list.
- The `+"`"+`health_check`+"`"+` settings for the load balancer.
- The `+"`"+`sticky_sessions`+"`"+` settings for the load balancer.
- The datacenter region the load balancer is located in.
- The Droplet tag corresponding to the Droplets assigned to the load balancer.
- The IDs of the Droplets assigned to the load balancer.
- Whether HTTP request to the load balancer on port 80 will be redirected to HTTPS on port 443.
- Whether the PROXY protocol is in use on the load balancer.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`IP`+"`"+`, `+"`"+`IPv6`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`SizeUnit`+"`"+`, `+"`"+`VPCUUID`+"`"+`, `+"`"+`Tag`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`RedirectHttpToHttps`+"`"+`, `+"`"+`StickySessions`+"`"+`, `+"`"+`HealthCheck`+"`"+`, `+"`"+`ForwardingRules`+"`"+`, `+"`"+`Firewall`+"`"+`, `+"`"+`DisableLetsEncryptDNSRecords`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer purge-cache
					ox.Usage(`purge-cache`, `Purges CDN cache for a global load balancer`),
					ox.Banner(`Use this command to purge the CDN cache for specified global load balancer.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Purge the global load balancer CDN cache without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer remove-droplets
					ox.Usage(`remove-droplets`, `Remove Droplets from a load balancer`),
					ox.Banner(`Use this command to remove Droplets from a load balancer. This command does not destroy any Droplets.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-ids`, `A comma-separated list of IDs of Droplets to remove from the load balancer, example value: 12,33`, ox.Spec(`12,33`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer remove-forwarding-rules
					ox.Usage(`remove-forwarding-rules`, `Remove forwarding rules from a load balancer`),
					ox.Banner(`Use this command to remove forwarding rules from a load balancer, specified with the `+"`"+`--forwarding-rules`+"`"+` flag. Valid rules include:
- `+"`"+`entry_protocol`+"`"+`: The entry protocol used for traffic to the load balancer. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`entry_port`+"`"+`: The entry port used for incoming traffic to the load balancer.
- `+"`"+`target_protocol`+"`"+`: The target protocol used for traffic from the load balancer to the backend Droplets. Possible values are: `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`, `+"`"+`http2`+"`"+`, `+"`"+`http3`+"`"+`, `+"`"+`tcp`+"`"+`, or `+"`"+`udp`+"`"+`.
- `+"`"+`target_port`+"`"+`: The target port used to send traffic from the load balancer to the backend Droplets.
- `+"`"+`certificate_id`+"`"+`: The ID of the TLS certificate used for SSL termination, if enabled. Can be obtained with `+"`"+`doctl certificate list`+"`"+`
- `+"`"+`tls_passthrough`+"`"+`: Whether SSL passthrough is enabled on the load balancer.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`forwarding-rules`, `A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306.`, ox.Spec(`entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute load-balancer update
					ox.Usage(`update`, `Update a load balancer's configuration`),
					ox.Banner(`Use this command to update the configuration of a specified load balancer. Using all applicable flags, the command should contain a full representation of the load balancer including existing attributes, such as the load balancer's name, region, forwarding rules, and Droplet IDs. Any attribute that is not provided is reset to its default value.`),
					ox.Spec(`<load-balancer-id> [flags]`),
					ox.Aliases("u"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`algorithm`, `This field has been deprecated. You can no longer specify an algorithm for load balancers.`, ox.Default("round_robin")).
						String(`allow-list`, `A comma-separated list of ALLOW rules for the load balancer, e.g.: ip:1.2.3.4,cidr:1.2.0.0/16`, ox.Spec(`ip:1.2.3.4,cidr:1.2.0.0/16`)).
						String(`deny-list`, `A comma-separated list of DENY rules for the load balancer, e.g.: ip:1.2.3.4,cidr:1.2.0.0/16`, ox.Spec(`ip:1.2.3.4,cidr:1.2.0.0/16`)).
						Bool(`disable-lets-encrypt-dns-records`, `disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer`).
						String(`domains`, `is_managed:true certificate_id:test-cert-id-1                                                                                   A comma-separated list of domains required to ingress traffic to a global load balancer, e.g.: name:test-domain-1 is_managed:true certificate_id:test-cert-id-1`, ox.Spec(`name:test-domain-1`)).
						String(`droplet-ids`, `A comma-separated list of Droplet IDs, e.g.: 215,378`, ox.Spec(`215,378`)).
						Bool(`enable-backend-keepalive`, `enable keepalive connections to backend target droplets`).
						Bool(`enable-proxy-protocol`, `enable proxy protocol`).
						String(`forwarding-rules`, `A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306.`, ox.Spec(`entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`)).
						String(`glb-cdn-settings`, `CDN cache settings global load balancer, e.g.: is_enabled:true`, ox.Spec(`is_enabled:true`)).
						String(`glb-settings`, `Target protocol and port settings for ingressing traffic to a global load balancer, e.g.: target_protocol:http,target_port:80`, ox.Spec(`target_protocol:http,target_port:80`)).
						String(`health-check`, `port:80, path:/index.html, check_interval_seconds:10, response_timeout_seconds:5, healthy_threshold:5, unhealthy_threshold:3   A comma-separated list of key-value pairs representing recent health check results, e.g.: protocol:http, port:80, path:/index.html, check_interval_seconds:10, response_timeout_seconds:5, healthy_threshold:5, unhealthy_threshold:3`, ox.Spec(`protocol:http,`)).
						String(`name`, `The load balancer's name`).
						String(`project-id`, `Indicates which project to associate the Load Balancer with. If not specified, the Load Balancer will be placed in your default project.`).
						Bool(`redirect-http-to-https`, `Flag to redirect HTTP requests to the load balancer on port 80 to HTTPS on port 443`).
						String(`region`, `The load balancer's region, e.g.: nyc1`, ox.Default("nyc1")).
						String(`size`, `The load balancer's size, e.g.: lb-small. Only one of size and size-unit should be used`, ox.Default("lb-small")).
						Int(`size-unit`, `The load balancer's size, e.g.: 1. Only one of size-unit and size should be used`).
						String(`sticky-sessions`, `cookie_name:DO-LB, cookie_ttl_seconds:5                                                                                      A comma-separated list of key-value pairs representing a list of active sessions, e.g.: type:cookies, cookie_name:DO-LB, cookie_ttl_seconds:5`, ox.Spec(`type:cookies,`)).
						String(`tag-name`, `Assigns Droplets with the specified tag to the load balancer`).
						Slice(`target-lb-ids`, `A comma-separated list of Load Balancer IDs to add as target to the global load balancer`).
						String(`vpc-uuid`, `The UUID of the VPC to create the load balancer in`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute region
				ox.Usage(`region`, `Display commands to list datacenter regions`),
				ox.Banner(`The subcommands of `+"`"+`doctl compute region`+"`"+` retrieve information about DigitalOcean datacenter regions.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute region [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute region list
					ox.Usage(`list`, `Retrieves a list of datacenter regions`),
					ox.Banner(`List DigitalOcean datacenter regions displaying their name, slug, and availability.

Use the slugs displayed by this command to specify regions in other commands.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of regions and uses the --format flag to return only the slug for each region: doctl compute region list --format Slug`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Slug, `+"`"+`Name`+"`"+`, `+"`"+`Available`+"`"+`.`, ox.Default("Slug")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute reserved-ip
				ox.Usage(`reserved-ip`, `Display commands to manage reserved IP addresses`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute reserved-ip`+"`"+` manage reserved IP addresses.
Reserved IPs are publicly-accessible static IP addresses that you can to one of your Droplets. They can be used to create highly available setups or other configurations requiring movable addresses. Reserved IPs are bound to the regions they are created in.`),
				ox.Spec(`[command]`),
				ox.Aliases("fip", "floating-ip", "floating-ips", "reserved-ips"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute reserved-ip [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute reserved-ip create
					ox.Usage(`create`, `Create a new reserved IP address`),
					ox.Banner(`Creates a new reserved IP address.

Reserved IP addresses can either be assigned to Droplets or held in the region they were created in on your account, but because of the IPv4 address shortage, unassigned reserved IP addresses incur charges.`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a reserved IP address in the `+"`"+`nyc1`+"`"+` region and assigns it to a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute reserved-ip create --region nyc1 --droplet-id 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`droplet-id`, `The ID of the Droplet to assign the reserved IP to. Cannot be used with the --region flag.`, ox.Spec(`--region`)).
						String(`format`, `Columns for output in a comma-separated list. Possible values: IP, `+"`"+`Region`+"`"+`, `+"`"+`DropletID`+"`"+`, `+"`"+`DropletName`+"`"+`, `+"`"+`ProjectID`+"`"+`.`, ox.Spec(`IP`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`project-id`, `The ID of the project to assign the IP address. When excluded, the address is assigned to your default project. When using the --droplet-id flag, it is assigned to the project containing the Droplet.`, ox.Spec(`--droplet-id`)).
						String(`region`, `The region where to create the reserved IP address. Cannot be used with the --droplet-id flag.`, ox.Spec(`--droplet-id`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute reserved-ip delete
					ox.Usage(`delete`, `Permanently delete a reserved IP address`),
					ox.Banner(`Permanently deletes a reserved IP address. This is irreversible.`),
					ox.Spec(`<reserved-ip> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes the reserved IP address `+"`"+`203.0.113.25`+"`"+`: doctl compute reserved-ip delete 203.0.113.25`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the reserved IP address without confirmation`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute reserved-ip get
					ox.Usage(`get`, `Retrieve information about a reserved IP address`),
					ox.Banner(`Retrieves detailed information about a reserved IP address, including its region and the ID of the Droplet its assigned to.`),
					ox.Spec(`<reserved-ip> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about the reserved IP address `+"`"+`203.0.113.25`+"`"+`: doctl compute reserved-ip get 203.0.113.25`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: IP, `+"`"+`Region`+"`"+`, `+"`"+`DropletID`+"`"+`, `+"`"+`DropletName`+"`"+`, `+"`"+`ProjectID`+"`"+`.`, ox.Spec(`IP`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute reserved-ip list
					ox.Usage(`list`, `List all reserved IP addresses on your account`),
					ox.Banner(`Retrieves a list of all the reserved IP addresses on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all reserved IP addresses in the `+"`"+`nyc1`+"`"+` region: doctl compute reserved-ip list --region nyc1`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: IP, `+"`"+`Region`+"`"+`, `+"`"+`DropletID`+"`"+`, `+"`"+`DropletName`+"`"+`, `+"`"+`ProjectID`+"`"+`.`, ox.Spec(`IP`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `Retrieves a list of reserved IP addresses in the specified region`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute reserved-ip-action
				ox.Usage(`reserved-ip-action`, `Display commands to associate reserved IP addresses with Droplets`),
				ox.Banner(`Reserved IP actions are commands that are used to manage DigitalOcean reserved IP addresses.`),
				ox.Spec(`[command]`),
				ox.Aliases("fipa", "floating-ip-action", "floating-ip-actions", "reserved-ip-actions"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute reserved-ip-action [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute reserved-ip-action assign
					ox.Usage(`assign`, `Assign a reserved IP address to a Droplet`),
					ox.Banner(`Assigns a reserved IP address to the specified Droplet.`),
					ox.Spec(`<reserved-ip> <droplet-id> [flags]`),
					ox.Example(`
The following example assigns the reserved IP address `+"`"+`203.0.113.25`+"`"+` to a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute reserved-ip-action assign 203.0.113.25 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute reserved-ip-action get
					ox.Usage(`get`, `Retrieve the status of a reserved IP action`),
					ox.Banner(`Retrieves the status of a reserved IP action. Outputs the following information:

- The unique numeric ID used to identify and reference a reserved IP action
- The status of the reserved IP action. Possible values: "in-progress", "completed", "errored"
- When the action was initiated, in ISO8601 combined date and time format
- When the action was completed, in ISO8601 combined date and time format
- The ID of the resource that the action is associated with
- The type of resource that the action is associated with
- The region where the action occurred
- The slug for the region where the action occurred`),
					ox.Spec(`<reserved-ip> <action-id> [flags]`),
					ox.Example(`
The following example retrieves the status of an action, that has the ID `+"`"+`191669331`+"`"+`, that was taken on the reserved IP address `+"`"+`203.0.113.25`+"`"+`: doctl compute reserved-ip-action get 203.0.113.25 191669331`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute reserved-ip-action unassign
					ox.Usage(`unassign`, `Unassign a reserved IP address from a Droplet`),
					ox.Banner(`Unassigns a reserved IP address from a Droplet. Due to a shortage on IPv4 addresses, unassigned reserved IP addresses remain available on your account but accumulate charges for not being assigned.`),
					ox.Spec(`<reserved-ip> [flags]`),
					ox.Example(`
The following example unassigns the reserved IP address `+"`"+`203.0.113.25`+"`"+` from a resource: doctl compute reserved-ip-action unassign 203.0.113.25`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute size
				ox.Usage(`size`, `List available Droplet sizes`),
				ox.Banner(`The subcommands of `+"`"+`doctl compute size`+"`"+` retrieve information about Droplet sizes.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute size [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute size list
					ox.Usage(`list`, `List available Droplet sizes`),
					ox.Banner(`Retrieves a list of slug identifiers, RAM amounts, vCPU counts, disk sizes, and pricing details for each Droplet size.

Use these slugs to specify the size of Droplet in other commands, such as `+"`"+`doctl compute droplet create <droplet-name> --size <size-slug>`+"`"+`.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of Droplet sizes and uses the --format flag to return only the slug for each size and its monthly price: doctl compute size list --format Slug,PriceMonthly`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Slug, `+"`"+`Description`+"`"+`, `+"`"+`Memory`+"`"+`, `+"`"+`VCPUs`+"`"+`, `+"`"+`Disk`+"`"+`, `+"`"+`PriceMonthly`+"`"+`, `+"`"+`PriceHourly`+"`"+`.`, ox.Default("Slug")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute snapshot
				ox.Usage(`snapshot`, `Access and manage snapshots`),
				ox.Banner(`The subcommands of `+"`"+`doctl compute snapshot`+"`"+` allow you to manage and retrieve information about Droplet and block storage volume snapshots.`),
				ox.Spec(`[command]`),
				ox.Aliases("s"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute snapshot [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute snapshot delete
					ox.Usage(`delete`, `Delete a snapshot of a Droplet or volume`),
					ox.Banner(`Deletes the specified snapshot or volume. This is irreversible.`),
					ox.Spec(`<snapshot-id>... [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes a Droplet snapshot with ID `+"`"+`386734086`+"`"+`: doctl compute snapshot delete 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the snapshot without confirmation`, ox.Short("f")).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`ResourceId`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`MinDiskSize`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute snapshot get
					ox.Usage(`get`, `Retrieve a Droplet or volume snapshot`),
					ox.Banner(`Retrieves information about a Droplet or block storage volume snapshot, including:

- The snapshot's ID
- The snapshot's name
- The date and time when the snapshot was created
- The slugs of the datacenter regions in which the snapshot is available
- The type of resource the snapshot was made from (either from Droplet or volume) and its ID
- The minimum size required for a Droplet or volume to use this snapshot, in GB
- The compressed, billable size of the snapshot`),
					ox.Spec(`<snapshot-id>... [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about a Droplet snapshot with ID `+"`"+`386734086`+"`"+`: doctl compute snapshot get 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`ResourceId`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`MinDiskSize`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute snapshot list
					ox.Usage(`list`, `List Droplet and volume snapshots`),
					ox.Banner(`Retrieves a list of snapshots and their information, including:

- The snapshot's ID
- The snapshot's name
- The date and time when the snapshot was created
- The slugs of the datacenter regions in which the snapshot is available
- The type of resource the snapshot was made from (either from Droplet or volume) and its ID
- The minimum size required for a Droplet or volume to use this snapshot, in GB
- The compressed, billable size of the snapshot`),
					ox.Spec(`[glob] [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all Droplet snapshots in the `+"`"+`nyc1`+"`"+` region and uses the `+"`"+`--format`+"`"+` flag to return only name, ID, and resource type for each snapshot: doctl compute snapshot list --resource droplet --region nyc1 --format Name,ID,ResourceType`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`ResourceId`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`MinDiskSize`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `Filters by regional availability`).
						String(`resource`, `Filters by resource type (droplet or `+"`"+`volume`+"`"+`)`, ox.Default("droplet")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute ssh
				ox.Usage(`ssh`, `Access a Droplet using SSH`),
				ox.Banner(`Access a Droplet using SSH by providing its ID or name.

You may specify the user to login with by passing the `+"`"+`--ssh-user`+"`"+` flag. To access the Droplet on a non-default port, use the `+"`"+`--ssh-port`+"`"+` flag. By default, the connection will be made to the Droplet's public IP address. In order access it using its private IP address, use the `+"`"+`--ssh-private-ip`+"`"+` flag.`),
				ox.Spec(`<droplet-id|name> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`ssh-agent-forwarding`, `Enable SSH agent forwarding`).
					String(`ssh-command`, `Command to execute on Droplet`).
					String(`ssh-key-path`, `Path to SSH private key`, ox.Default("$HOME/.ssh/id_rsa")).
					Int(`ssh-port`, `The remote port sshd is running on`, ox.Default("22")).
					Bool(`ssh-private-ip`, `SSH to Droplet's private IP address`).
					Int(`ssh-retry-max`, `Max number of retries for a successful SSH connection to a Droplet`, ox.Default("is 0")).
					String(`ssh-user`, `SSH user for connection`, ox.Default("root")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl compute ssh-key
				ox.Usage(`ssh-key`, `Display commands to manage SSH keys on your account`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute ssh-key`+"`"+` manage the SSH keys on your account.

DigitalOcean allows you to add SSH public keys to the interface so that you can embed your public key into a Droplet at the time of creation. Only the public key is required to take advantage of this functionality. Note that this command does not add, delete, or otherwise modify any ssh keys that may be on existing Droplets.`),
				ox.Spec(`[command]`),
				ox.Aliases("k"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute ssh-key [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute ssh-key create
					ox.Usage(`create`, `Create a new SSH key on your account`),
					ox.Banner(`Use this command to add a new SSH key to your account.

Specify a `+"`"+`<key-name>`+"`"+` for the key, and set the `+"`"+`--public-key`+"`"+` flag to a string with the contents of the key.

Note that creating a key will not add it to any Droplets.`),
					ox.Spec(`<key-name> [flags]`),
					ox.Aliases("c"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`FingerPrint`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`public-key`, `Key contents (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute ssh-key delete
					ox.Usage(`delete`, `Permanently delete an SSH key from your account`),
					ox.Banner(`Use this command to permanently delete an SSH key from your account.

Note that this does not delete an SSH key from any Droplets.`),
					ox.Spec(`<key-id|key-fingerprint> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the key without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute ssh-key get
					ox.Usage(`get`, `Retrieve information about an SSH key on your account`),
					ox.Banner(`Use this command to get the id, fingerprint, public_key, and name of a specific SSH key on your account.`),
					ox.Spec(`<key-id|key-fingerprint> [flags]`),
					ox.Aliases("g"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`FingerPrint`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute ssh-key import
					ox.Usage(`import`, `Import an SSH key from your computer to your account`),
					ox.Banner(`Use this command to add a new SSH key to your account, using a local public key file.

Note that importing a key to your account will not add it to any Droplets`),
					ox.Spec(`<key-name> [flags]`),
					ox.Aliases("i"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`FingerPrint`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`public-key-file`, `Public key file (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute ssh-key list
					ox.Usage(`list`, `List all SSH keys on your account`),
					ox.Banner(`Use this command to list the id, fingerprint, public_key, and name of all SSH keys on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`FingerPrint`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute ssh-key update
					ox.Usage(`update`, `Update an SSH key's name`),
					ox.Banner(`Use this command to update the name of an SSH key.`),
					ox.Spec(`<key-id|key-fingerprint> [flags]`),
					ox.Aliases("u"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`FingerPrint`+"`"+`.`, ox.Spec(`ID`)).
						String(`key-name`, `Key name (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute tag
				ox.Usage(`tag`, `Display commands to manage tags`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute tag`+"`"+` manage the tags on your account.

Tags are labels that you can apply to resources to better organize them and more efficiently take actions on them. For example, if you have a group of Droplets that you want to place behind the same set of cloud firewall rules, you can tag those Droplets with a common tag and then apply the firewall rules to all Droplets with that tag.

You can tag Droplets, images, volumes, volume snapshots, and database clusters.

Tags have two attributes: a user defined name attribute and an embedded
resources attribute with information about resources that have been tagged.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute tag [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute tag apply
					ox.Usage(`apply`, `Apply a tag to resources`),
					ox.Banner(`Tag one or more resources. You can tag Droplets, images, volumes, volume snapshots, and database clusters.
	
Resources must be specified as Uniform Resource Names (URNs) and has the following syntax: `+"`"+`do:<resource_type>:<identifier>`+"`"+`.`),
					ox.Spec(`<tag-name> --resource=<urn> [--resource=<urn> ...] [flags]`),
					ox.Example(`
The following example tags two Droplet with the tag named `+"`"+`web`+"`"+`: doctl compute tag apply web --resource=do:droplet:386734086,do:droplet:191669331`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`resource`, `The resource to tag in URN format (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute tag create
					ox.Usage(`create`, `Create a tag`),
					ox.Banner(`Creates a new tag that you can apply to resources.`),
					ox.Spec(`<tag-name> [flags]`),
					ox.Example(`
The following example creates a tag name `+"`"+`web`+"`"+`: doctl compute tag create web`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute tag delete
					ox.Usage(`delete`, `Delete a tag`),
					ox.Banner(`Deletes a tag from your account.

Deleting a tag also removes the tag from all the resources that had been tagged with it.`),
					ox.Spec(`<tag-name>... [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes the tag named `+"`"+`web`+"`"+`: doctl compute tag delete web`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete tag without confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute tag get
					ox.Usage(`get`, `Retrieve information about a tag`),
					ox.Banner(`Retrieves the number of resources using the tag.`),
					ox.Spec(`<tag-name> [flags]`),
					ox.Example(`
The following example retrieves information about the tag named `+"`"+`web`+"`"+`: doctl compute tag get web`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`DropletCount`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute tag list
					ox.Usage(`list`, `List all tags`),
					ox.Banner(`Retrieves a list of all the tags in your account and how many resources are using each tag.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`DropletCount`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute tag remove
					ox.Usage(`remove`, `Remove a tag from resources`),
					ox.Banner(`Removes a tag from one or more resources. Resources must be specified as Uniform Resource Names (URNs) and has the following syntax: `+"`"+`do:<resource_type>:<identifier>`+"`"+`.`),
					ox.Spec(`<tag-name> --resource=<urn> [--resource=<urn> ...] [flags]`),
					ox.Example(`
The following example removes the tag named `+"`"+`web`+"`"+` from two Droplets: doctl compute tag remove web --resource=do:droplet:386734086,do:droplet:191669331`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`resource`, `The resource to untag in URN format (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute volume
				ox.Usage(`volume`, `Display commands to manage block storage volumes`),
				ox.Banner(`The sub-commands of `+"`"+`doctl compute volume`+"`"+` manage your block storage volumes.

Block storage volumes provide expanded storage capacity for your Droplets, ranging in size from 1GiB to 16TiB.

Volumes function as raw block devices, meaning they appear to the operating system as locally attached storage which can be formatted using any filesystem supported by the OS. They can be moved between Droplets located in the same region as the volume.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute volume [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute volume create
					ox.Usage(`create`, `Create a block storage volume`),
					ox.Banner(`Creates a block storage volume on your account.

You can use flags to specify the volume size, region, description, filesystem type, tags, and to create a volume from an existing volume snapshot.

Use the `+"`"+`doctl compute volume-action attach <volume-id> <droplet-id>`+"`"+` command to attach a new volume to a Droplet.`),
					ox.Spec(`<volume-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a 4TiB volume named `+"`"+`example-volume`+"`"+` in the `+"`"+`nyc1`+"`"+` region. The command also applies two tags to the volume: doctl compute volume create example-volume --region nyc1 --size 4TiB --tag frontend,backend`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`desc`, `A description of the volume`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Filesystem Type`+"`"+`, `+"`"+`Filesystem Label`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						String(`fs-label`, `The volume's filesystem label`).
						String(`fs-type`, `The volume's filesystem type: ext4 or xfs. If not specified, the volume is left unformatted`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `The volume's region. Not compatible with the --snapshot flag`, ox.Spec(`--snapshot`)).
						String(`size`, `Volume size (required)`, ox.Default("4TiB")).
						String(`snapshot`, `Creates a volume from the specified snapshot ID. Not compatible with the --region flag`, ox.Spec(`--region`)).
						String(`tag`, `frontend   A comma-separated list of tags to apply to the volume. For example, --tag frontend or `+"`"+`--tag frontend,backend`+"`"+``, ox.Spec(`--tag`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume delete
					ox.Usage(`delete`, `Delete a block storage volume`),
					ox.Banner(`Deletes a block storage volume by ID, destroying all of its data and removing it from your account. This is irreversible.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute volume delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the volume without prompting for confirmation`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume get
					ox.Usage(`get`, `Retrieve an existing block storage volume`),
					ox.Banner(`Retrieves information about a block storage volume.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute volume get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Filesystem Type`+"`"+`, `+"`"+`Filesystem Label`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume list
					ox.Usage(`list`, `List block storage volumes by ID`),
					ox.Banner(`Lists all of the block storage volumes on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of volumes on your account in the `+"`"+`nyc1`+"`"+` region. The command also uses the `+"`"+`--format`+"`"+` flag to return only the name and size of each volume: doctl compute volume list --region nyc1 --format Name,Size`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Filesystem Type`+"`"+`, `+"`"+`Filesystem Label`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`region`, `Filter's volumes by the specified region`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume snapshot
					ox.Usage(`snapshot`, `Create a block storage volume snapshot`),
					ox.Banner(`Creates a snapshot of a block storage volume by ID.

You can use a block storage volume snapshot ID as a flag with `+"`"+`doctl volume create`+"`"+` to create a new block storage volume with the same data as the volume the snapshot was taken from.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Aliases("s"),
					ox.Example(`
The following example creates a snapshot of a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute volume snapshot f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --snapshot-name example-snapshot --tag frontend,backend`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Filesystem Type`+"`"+`, `+"`"+`Filesystem Label`+"`"+`, `+"`"+`DropletIDs`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`snapshot-desc`, `A description of the snapshot`).
						String(`snapshot-name`, `The snapshot name (required)`).
						String(`tag`, `frontend     A comma-separate list of tags to apply to the snapshot. For example, --tag frontend or `+"`"+`--tag frontend,backend`+"`"+``, ox.Spec(`--tag`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl compute volume-action
				ox.Usage(`volume-action`, `Display commands to perform actions on a volume`),
				ox.Banner(`Block storage volume action commands allow you to attach, detach, and resize existing volumes.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl compute volume-action [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl compute volume-action attach
					ox.Usage(`attach`, `Attach a volume to a Droplet`),
					ox.Banner(`Attaches a block storage volume to a Droplet.

You can only attach one Droplet to a volume at a time. However, you can attach up to fifteen different volumes to a Droplet at a time.

When you attach a pre-formatted volume to Ubuntu, Debian, Fedora, Fedora Atomic, and CentOS Droplets created on or after April 26, 2018, the volume automatically mounts. On older Droplets, additional configuration is required. Visit https://docs.digitalocean.com/products/volumes/how-to/mount/ for details`),
					ox.Spec(`<volume-id> <droplet-id> [flags]`),
					ox.Aliases("a"),
					ox.Example(`
The following example attaches a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute volume-action attach f81d4fae-7dec-11d0-a765-00a0c91e6bf6 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`wait`, `Instructs the terminal to wait for the volume to attach before returning control to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume-action detach
					ox.Usage(`detach`, `Detach a volume from a Droplet`),
					ox.Banner(`Detaches a block storage volume from a Droplet.`),
					ox.Spec(`<volume-id> <droplet-id> [flags]`),
					ox.Aliases("d"),
					ox.Example(`
The following example detaches a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` from a Droplet with the ID `+"`"+`386734086`+"`"+`: doctl compute volume-action detach f81d4fae-7dec-11d0-a765-00a0c91e6bf6 386734086`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`wait`, `Instructs the terminal to wait for the volume to detach before returning control to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume-action detach-by-droplet-id
					ox.Usage(`detach-by-droplet-id`, `(Deprecated) Detach a volume. Use `+"`"+`detach`+"`"+` instead.`),
					ox.Banner(`This command detaches a volume. This command is deprecated. Use `+"`"+`doctl compute volume-action detach`+"`"+` instead.`),
					ox.Spec(`<volume-id> <droplet-id> [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume-action get
					ox.Usage(`get`, `Retrieve the status of a volume action`),
					ox.Banner(`Retrieves the status of a volume action, including the following details:

- The unique numeric ID used to identify and reference a volume action.
- The status of the volume action. Possible values: `+"`"+`in-progress`+"`"+`, `+"`"+`completed`+"`"+`, `+"`"+`errored`+"`"+`.
- When the action was initiated, in ISO8601 combined date and time format
- When the action was completed, in ISO8601 combined date and time format
- The resource ID, which is a unique identifier for the resource that the action is associated with.
- The type of resource that the action is associated with.
- The region where the action occurred.
- The slug for the region where the action occurred.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Example(`
The following example retrieves the status of an action taken on a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl compute volume-action get f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --action-id 191669331`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Int(`action-id`, `action id (required)`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume-action list
					ox.Usage(`list`, `Retrieve a list of actions taken on a volume`),
					ox.Banner(`Retrieves a list of actions taken on a volume. The following details are provided:

- The unique numeric ID used to identify and reference a volume action.
- The status of the volume action. Possible values: `+"`"+`in-progress`+"`"+`, `+"`"+`completed`+"`"+`, `+"`"+`errored`+"`"+`.
- When the action was initiated, in ISO8601 combined date and time format
- When the action was completed, in ISO8601 combined date and time format
- The resource ID, which is a unique identifier for the resource that the action is associated with.
- The type of resource that the action is associated with.
- The region where the action occurred.
- The slug for the region where the action occurred.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of actions taken on a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`. The command also uses the `+"`"+`--format`+"`"+` flag to return only the resource ID and status for each action listed: doctl compute volume-action list f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --format ResourceID,Status`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Status`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`StartedAt`+"`"+`, `+"`"+`CompletedAt`+"`"+`, `+"`"+`ResourceID`+"`"+`, `+"`"+`ResourceType`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl compute volume-action resize
					ox.Usage(`resize`, `Resize the disk of a volume`),
					ox.Banner(`Resizes a block storage volume.

Volumes may only be resized upwards. The maximum size for a volume is 16TiB.`),
					ox.Spec(`<volume-id> [flags]`),
					ox.Aliases("r"),
					ox.Example(`
The following example resizes a volume with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to 120 GiB in the `+"`"+`nyc1`+"`"+` region: doctl compute volume-action resize f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --size 120 --region nyc1`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`region`, `The volume's current region (required)`).
						Int(`size`, `The volume's new size, in GiB (required)`).
						Bool(`wait`, `Instructs the terminal to wait for the volume to complete resizing before returning control to the user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
		),
		ox.Sub( // doctl databases
			ox.Usage(`databases`, `Display commands that manage databases`),
			ox.Banner(`The commands under `+"`"+`doctl databases`+"`"+` are for managing your MySQL, Redis, PostgreSQL, MongoDB, Kafka and Opensearch database services.`),
			ox.Spec(`[command]`),
			ox.Aliases("db", "dbs", "d", "database"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl databases [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl databases backups
				ox.Usage(`backups`, `List database cluster backups`),
				ox.Banner(`Retrieves a list of backups created for the specified database cluster.

The list contains the size in GB, and the date and time the backup was created.`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("bu"),
				ox.Example(`
The following example retrieves a list of backups for a database cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl databases backups f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Size, `+"`"+`Created`+"`"+`.`, ox.Default("Size")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases configuration
				ox.Usage(`configuration`, `View the configuration of a database cluster given its ID and Engine`),
				ox.Banner(`The subcommands of `+"`"+`doctl databases configuration`+"`"+` are used to view a database cluster's configuration.`),
				ox.Spec(`[command]`),
				ox.Aliases("cfg", "config"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases configuration [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases configuration get
					ox.Usage(`get`, `Get a database cluster's configuration`),
					ox.Banner(`Retrieves the advanced configuration for the specified cluster, including its backup settings, temporary file limit, and session timeout values.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("g"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`engine`, `The engine of the database you want to get the configuration for. (required)`, ox.Short("e")).
						String(`format`, `Columns for output in a comma-separated list. Possible values: key, `+"`"+`value`+"`"+`.`, ox.Default("key")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases configuration update
					ox.Usage(`update`, `Update a database cluster's configuration`),
					ox.Banner(`Updates the specified database cluster's advanced configuration. Using this command, you can update various settings like backup times, temporary file limits, and session timeouts. Available settings vary by database engine.

This command functions as a PATCH request, meaning that only the specified fields are updated. If a field is not specified, it will not be changed. The settings are passed using the `+"`"+`--config-json`+"`"+` flag, which takes a JSON object as its value.

For a full list of available fields, see the API documentation: https://docs.digitalocean.com/reference/api/api-reference/#operation/databases_patch_config`),
					ox.Spec(`<db-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following command updates a MySQL database's time zone: doctl databases configuration update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --engine mysql --config-json '{"default_time_zone":"Africa/Maputo"}'`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`config-json`, `the desired configuration of the database cluster you want to update (required)`, ox.Default("{}")).
						String(`engine`, `the engine of the database you want to update the configuration for (required)`, ox.Short("e")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases connection
				ox.Usage(`connection`, `Retrieve connection details for a database cluster`),
				ox.Banner(`Retrieves the following connection details for a database cluster:

- A connection string for the database cluster
- The default database name
- The fully-qualified domain name of the publicly-connectable host
- The port on which the database is listening for connections
- The default username
- The randomly-generated password for the default username
- A boolean value indicating if the connection should be made over SSL

While you can use these connection details, you can manually update the connection string's parameters to change how you connect to the database, such using a private hostname, custom username, or a different database.`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("conn"),
				ox.Example(`
The following example retrieves the connection details for a database cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl databases connection f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: URI, `+"`"+`Database`+"`"+`, `+"`"+`Host`+"`"+`, `+"`"+`Port`+"`"+`, `+"`"+`User`+"`"+`, `+"`"+`Password`+"`"+`, `+"`"+`SSL`+"`"+`.`, ox.Default("URI")).
					Bool(`no-header`, `Return raw data with no headers`).
					Bool(`private`, `Returns connection details that use the database's VPC network connection.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases create
				ox.Usage(`create`, `Create a database cluster`),
				ox.Banner(`Creates a database cluster with the specified name.

You can customize the configuration using the listed flags, all of which are optional. Without any flags set, the command creates a single-node, single-CPU PostgreSQL database cluster.`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("c"),
				ox.Example(`
The following example creates a database cluster named `+"`"+`example-database`+"`"+` in the `+"`"+`nyc1`+"`"+` region with a single  1 GB node: doctl databases create example-database --region nyc1 --size db-s-1vcpu-1gb --num-nodes 1`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`engine`, `The database's engine. Possible values are: pg, `+"`"+`mysql`+"`"+`, `+"`"+`redis`+"`"+`, `+"`"+`mongodb`+"`"+`, `+"`"+`kafka`+"`"+` and `+"`"+`opensearch`+"`"+`.`, ox.Default("pg")).
					Int(`num-nodes`, `The number of nodes in the database cluster. Valid values are 1-3. In addition to the primary node, up to two standby nodes may be added for high availability.`, ox.Default("1")).
					String(`private-network-uuid`, `The UUID of a VPC to create the database cluster in. The command uses the region's default VPC if excluded.`).
					String(`region`, `The data center region where the database cluster resides, such as nyc1 or `+"`"+`sfo2`+"`"+`.`, ox.Default("nyc1")).
					String(`restore-from-cluster-name`, `The name of an existing database cluster to restore from.`).
					String(`restore-from-timestamp`, `The timestamp of an existing database cluster backup in UTC combined date and time format (2006-01-02 15:04:05 +0000 UTC). The most recent backup is used if excluded.`).
					String(`size`, `The size of the nodes in the database cluster, for example db-s-1vcpu-1gb indicates a 1 CPU, 1GB node. For a list of available size slugs, visit: https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases`, ox.Default("db-s-1vcpu-1gb")).
					Int(`storage-size-mib`, `The amount of disk space allocated to the cluster. Applicable for PostgreSQL and MySQL clusters. Each plan size has a default value but can be increased in increments up to a maximum amount. For ranges, visit: https://www.digitalocean.com/pricing/managed-databases`).
					Slice(`tag`, `A comma-separated list of tags to apply to the database cluster.`).
					Bool(`wait`, `A boolean value that specifies whether to wait for the database cluster to be provisioned before returning control to the terminal.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases db
				ox.Usage(`db`, `Display commands for managing individual databases within a cluster`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases db`+"`"+` are for managing specific databases that are served by a database cluster.

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases db [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases db create
					ox.Usage(`create`, `Create a database within a cluster`),
					ox.Banner(`Creates a database with the specified name in the specified database cluster.

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <database-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a database named `+"`"+`example-db`+"`"+` in a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases db create ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-db`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases db delete
					ox.Usage(`delete`, `Delete the specified database from the cluster`),
					ox.Banner(`Deletes the specified database from the specified database cluster.

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list

You can get a list of existing databases that are hosted within a cluster by calling:

	doctl databases db list <cluster-id>`),
					ox.Spec(`<database-cluster-id> <database-name> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes a database named `+"`"+`example-db`+"`"+` in a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases db delete ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-db`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the database without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases db get
					ox.Usage(`get`, `Retrieve the name of a database within a cluster`),
					ox.Banner(`Retrieves the name of the specified database hosted in the specified database cluster.

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list

You can get a list of existing databases that are hosted within a cluster by calling:

	doctl databases db list <cluster-id>`),
					ox.Spec(`<database-cluster-id> <database-name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the name of a database in a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` and the name `+"`"+`example-db`+"`"+`: doctl databases db get ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-db`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases db list
					ox.Usage(`list`, `Retrieve a list of databases within a cluster`),
					ox.Banner(`Retrieves a list of databases being hosted in the specified database cluster.

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of databases in a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases db list ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases delete
				ox.Usage(`delete`, `Delete a database cluster`),
				ox.Banner(`Deletes the database cluster with the specified ID.

To retrieve a list of your database clusters and their IDs, use `+"`"+`doctl databases list`+"`"+`.`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("rm"),
				ox.Example(`
The following example deletes the database cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl databases delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `Delete the database cluster without a confirmation prompt`, ox.Short("f")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases events
				ox.Usage(`events`, `Display commands for listing database cluster events`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases events`+"`"+` are for listing database cluster events.

You can get a list of database events by calling:

	doctl databases events list <cluster-id>`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases events [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases events list
					ox.Usage(`list`, `List your database cluster events`),
					ox.Banner(`Retrieves a list of database clusters events:

You can get a list of database events by calling:

	doctl databases events list <cluster-id>`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of databases events in a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases events list ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`ServiceName`+"`"+`, `+"`"+`EventType`+"`"+`, `+"`"+`CreateTime`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases firewalls
				ox.Usage(`firewalls`, `Display commands to manage firewall rules (called`+"`"+`trusted sources`+"`"+` in the control panel) for database clusters`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases firewalls`+"`"+` enable the management of firewalls for database clusters.`),
				ox.Spec(`[command]`),
				ox.Aliases("fw"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases firewalls [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases firewalls append
					ox.Usage(`append`, `Add a database firewall rule to a given database`),
					ox.Banner(`
Appends a single rule to the existing firewall rules of the specified database.

This command requires the `+"`"+`--rule`+"`"+` flag specifying the resource or resources allowed to access the database cluster. The rule passed to the `+"`"+`--rule`+"`"+` flag must be in a <type>:<value> format where:
	- `+"`"+`type`+"`"+` is the type of resource that the firewall rule allows to access the database cluster. Possible values are:  `+"`"+`droplet`+"`"+`, `+"`"+`k8s", `+"`"+`ip_addr`+"`"+`, `+"`"+`tag`+"`"+`, `+"`"+`app`+"`"+`
	- `+"`"+`value`+"`"+` is either the ID of a specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.`),
					ox.Spec(`<database-cluster-id> --rule <type>:<value> [flags]`),
					ox.Aliases("a"),
					ox.Example(`
The following example appends a firewall rule to a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` that allows any resources with the `+"`"+`example-tag`+"`"+` to access the database: doctl databases firewalls append ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --rule tag:example-tag`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`rule`, `(required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases firewalls list
					ox.Usage(`list`, `Retrieve a list of firewall rules for a given database`),
					ox.Banner(`
This command lists the following details for each firewall rule in a given database:

	- The UUID of the firewall rule
	- The UUID of the cluster for which the rule is applied
	- The type of resource that the firewall rule allows to access the database cluster. Possible values are: `+"`"+`droplet`+"`"+`, `+"`"+`k8s`+"`"+`, `+"`"+`ip_addr`+"`"+`, `+"`"+`tag`+"`"+`, `+"`"+`app`+"`"+`
	- The value, which specifies the resource or resources allowed to access the database cluster. Possible values are either the ID of the specific resource, the name of a tag applied to a group of resources, or an IP address
	- When the firewall rule was created, in ISO8601 date/time format
	

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of firewall rules for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases firewalls list ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases firewalls remove
					ox.Usage(`remove`, `Remove a firewall rule for a given database`),
					ox.Banner(`
Removes single rule from the list of firewall rules for a specified database. You can retrieve a firewall rule's UUIDs by calling:

	doctl database firewalls list <database-cluster-id>`),
					ox.Spec(`<database-cluster-id> --uuid <firerule-uuid> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example removes a firewall rule with the UUID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` from a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases firewalls remove ca9f591d-f38h-5555-a0ef-1c02d1d1e35 f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`uuid`, `(required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases firewalls replace
					ox.Usage(`replace`, `Replaces the firewall rules for a given database. The rules passed to the `+"`"+`--rules`+"`"+` flag replace the firewall rules previously assigned to the database,`),
					ox.Banner(`
Replace the firewall rules for a specified database. This command requires the `+"`"+`--rule`+"`"+` flag.

You can configure multiple rules for the firewall by passing additional arguments in a comma-separated list with the `+"`"+`--rule`+"`"+` flag. Each rule passed using the `+"`"+`--rule`+"`"+` flag must be in a `+"`"+`<type>:<value>`+"`"+` format where:
	 `+"`"+`type`+"`"+` is the type of resource that the firewall rule allows to access the database cluster. Possible values are:  `+"`"+`droplet`+"`"+`, `+"`"+`k8s`+"`"+`, `+"`"+`ip_addr`+"`"+`, `+"`"+`tag`+"`"+`, `+"`"+`app`+"`"+`
	- `+"`"+`value`+"`"+` is either the ID of a specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.`),
					ox.Spec(`<database-cluster-id> --rules type:value [--rule type:value] [flags]`),
					ox.Aliases("r"),
					ox.Example(`
The following example replaces the firewall rules for a database cluster, with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`, with rules that allow a specific Droplet, a specific IP address, and any resources with the `+"`"+`example-tag`+"`"+` to access the database: doctl databases firewalls replace ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --rules droplet:f81d4fae-7dec-11d0-a765-00a0c91e6bf6,ip_addr:192.168.1.1,tag:example-tag`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`rule`, `A comma-separated list of firewall rules, in type:value format. (required)`, ox.Spec(`type:value`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases fork
				ox.Usage(`fork`, `Create a new database cluster by forking an existing database cluster.`),
				ox.Banner(`Creates a new database cluster from an existing cluster. The forked database contains all of the data from the original database at the time the fork is created.`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("f"),
				ox.Example(`
The following example forks a database cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to create a new database cluster. The command also uses the `+"`"+`--restore-from-timestamp`+"`"+` flag to specifically fork the database from a cluster backup that was created on 2023 November 7: doctl databases fork new-db-cluster --restore-from-cluster-id f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --restore-from-timestamp 2023-11-07 12:34:56 +0000 UTC`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`restore-from-cluster-id`, `The ID of an existing database cluster from which the new database will be forked from (required)`).
					String(`restore-from-timestamp`, `The timestamp of an existing database cluster backup in UTC combined date and time format (2006-01-02 15:04:05 +0000 UTC). The most recent backup is used if excluded.`).
					Bool(`wait`, `A boolean that specifies whether to wait for a database to complete before returning control to the terminal`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases get
				ox.Usage(`get`, `Get details for a database cluster`),
				ox.Banner(`Retrieves the following details about the specified database cluster: 

- The database ID, in UUID format
- The name you gave the database cluster
- The database engine. Possible values: `+"`"+`redis`+"`"+`, `+"`"+`pg`+"`"+`, `+"`"+`mysql`+"`"+` , `+"`"+`mongodb`+"`"+`, `+"`"+`kafka`+"`"+`, `+"`"+`opensearch`+"`"+`
- The engine version, such as `+"`"+`14`+"`"+` for PostgreSQL version 14
- The number of nodes in the database cluster
- The region the database cluster resides in, such as `+"`"+`sfo2`+"`"+`, `+"`"+`nyc1`+"`"+`
- The current status of the database cluster, such as `+"`"+`online`+"`"+`
- The size of the machine running the database instance, such as `+"`"+`db-s-1vcpu-1gb`+"`"+`)
- A connection string for the database cluster
- The date and time when the database cluster was created

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves the details for a database cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the database's ID, engine, and engine version: doctl databases get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Engine`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`NumNodes`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`URI`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`StorageMib`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases get-ca
				ox.Usage(`get-ca`, `Provides the CA certificate for a DigitalOcean database`),
				ox.Banner(`Retrieves a database certificate`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("gc"),
				ox.Example(`
Retrieves the database certificate for the cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl databases get-ca f81d4fae-7dec-11d0-a765-00a0c91e6bf6
With the `+"`"+`-o json flag`+"`"+`, the certificate to connect to the database is base64 encoded. To decode it: `+"`"+`doctl databases get-ca <database-cluster-id> -o json | jq -r .certificate | base64 --decode`+"`"+``),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Certificate.`, ox.Spec(`Certificate`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases indexes
				ox.Usage(`indexes`, `Display commands to manage indexes for opensearch clusters`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases indexes`+"`"+` enable the management of indexes for opensearch clusters`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases indexes [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases indexes delete
					ox.Usage(`delete`, `Deletes an opensearch index by index name`),
					ox.Banner(`Deletes an opensearch index by index name`),
					ox.Spec(`<database-uuid> <index-name> [flags]`),
					ox.Aliases("rm"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the opensearch index without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases indexes list
					ox.Usage(`list`, `Retrieve a list of indexes for a given opensearch cluster`),
					ox.Banner(`
This command lists the following details for each index in an opensearch cluster:

	- The Name of the index.
	- The Status of the index.
	- The Health of the index.
	- The Number of Shards in the index.
	- The Number of Replicas in the index.
	- The Number of Documents in the index.
	- The Size of the index.`),
					ox.Spec(`<database-uuid> [flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Name   Columns for output in a comma-separated list. Possible values: Index Name, `+"`"+`Status`+"`"+`, `+"`"+`Health`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Docs`+"`"+`, `+"`"+`Create At`+"`"+`, `+"`"+`Number of Shards`+"`"+`, `+"`"+`Number of Replica`+"`"+`.`, ox.Default("Index")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases list
				ox.Usage(`list`, `List your database clusters`),
				ox.Banner(`Retrieves a list of database clusters and their following details:

- The database ID, in UUID format
- The name you gave the database cluster
- The database engine. Possible values: `+"`"+`redis`+"`"+`, `+"`"+`pg`+"`"+`, `+"`"+`mysql`+"`"+` , `+"`"+`mongodb`+"`"+`, `+"`"+`kafka`+"`"+`, `+"`"+`opensearch`+"`"+`
- The engine version, such as `+"`"+`14`+"`"+` for PostgreSQL version 14
- The number of nodes in the database cluster
- The region the database cluster resides in, such as `+"`"+`sfo2`+"`"+`, `+"`"+`nyc1`+"`"+`
- The current status of the database cluster, such as `+"`"+`online`+"`"+`
- The size of the machine running the database instance, such as `+"`"+`db-s-1vcpu-1gb`+"`"+`)`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example lists all database associated with your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, engine, and engine version of each database: doctl databases list --format ID,Engine,Version`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Engine`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`NumNodes`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`URI`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`StorageMib`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases maintenance-window
				ox.Usage(`maintenance-window`, `Display commands for scheduling automatic maintenance on your database cluster`),
				ox.Banner(`The `+"`"+`doctl databases maintenance-window`+"`"+` commands allow you to schedule, and check the schedule of, maintenance windows for your databases.

Maintenance windows are hour-long blocks of time during which DigitalOcean performs automatic maintenance on databases every week. During this time, health checks, security updates, version upgrades, and more are performed.

To install an update outside of a maintenance window, use the `+"`"+`doctl databases maintenance-window install`+"`"+` command.`),
				ox.Spec(`[command]`),
				ox.Aliases("maintenance", "mw", "main"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases maintenance-window [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases maintenance-window get
					ox.Usage(`get`, `Retrieve details about a database cluster's maintenance windows`),
					ox.Banner(`Retrieves the following information on currently-scheduled maintenance windows for the specified database cluster:

- The day of the week the maintenance window occurs
- The hour in UTC when maintenance updates will be applied, in 24 hour format, such as "16:00"
- A boolean representing whether maintenance updates are currently pending

To see a list of your databases and their IDs, run `+"`"+`doctl databases list`+"`"+`.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the maintenance window for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases maintenance-window ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Day, `+"`"+`Hour`+"`"+`, `+"`"+`Pending`+"`"+`.`, ox.Default("Day")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases maintenance-window install
					ox.Usage(`install`, `Start installation of updates immediately`),
					ox.Banner(`Starts the installation of updates for the specified database cluster immediately outside of a maintenance window.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("i"),
					ox.Example(`
The following example starts installation of updates for your databases with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases maintenance-window install ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases maintenance-window update
					ox.Usage(`update`, `Update the maintenance window for a database cluster`),
					ox.Banner(`Updates the maintenance window for the specified database cluster.

Maintenance windows are hour-long blocks of time during which DigitalOcean performs automatic maintenance on databases every week. During this time, health checks, security updates, version upgrades, and more are performed.

To change the maintenance window for your database cluster, specify a day of the week and an hour of that day during which you would prefer such maintenance would occur.

To see a list of your databases and their IDs, run `+"`"+`doctl databases list`+"`"+`.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates the maintenance window for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases maintenance-window update ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --day tuesday --hour 16:00`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`day`, `The day of the week the maintenance window occurs, for example: 'tuesday') (required)`).
						String(`hour`, `The hour when maintenance updates are applied, in UTC 24-hour format. Example: '16:00') (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases migrate
				ox.Usage(`migrate`, `Migrate a database cluster to a new region`),
				ox.Banner(`Migrates the specified database cluster to a new region.`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("m"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`private-network-uuid`, `The UUID of a VPC network to create the database cluster in. The command uses the region's default VPC network if not specified.`).
					String(`region`, `The region to which the database cluster should be migrated, such as sfo2 or `+"`"+`nyc3`+"`"+`. (required)`, ox.Default("sfo2")).
					Bool(`wait`, `A boolean value that specifies whether to wait for the database migration to complete before returning control to the terminal.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases options
				ox.Usage(`options`, `Display available database options (regions, version, layouts, etc.) for all available database engines`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases options`+"`"+` retrieve configuration options for databases, such as available engines, engine versions and their equivalent slugs.`),
				ox.Spec(`[command]`),
				ox.Aliases("o"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases options [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases options engines
					ox.Usage(`engines`, `Retrieves a list of the available database engines`),
					ox.Banner(`Lists the available database engines for DigitalOcean Managed Databases.`),
					ox.Spec(`[flags]`),
					ox.Aliases("eng"),
					ox.Example(`
The following example retrieves a list of the available database engines: doctl databases options engines`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases options regions
					ox.Usage(`regions`, `Retrieves a list of the available regions for a given database engine`),
					ox.Banner(`Lists the available regions for a given database engine. Some engines may not be available in certain regions.`),
					ox.Spec(`[flags]`),
					ox.Aliases("r"),
					ox.Example(`
The following example retrieves a list of the available regions for the PostgreSQL engine: doctl databases options regions --engine pg`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`engine`, `The database engine. Possible values:  mysql,  `+"`"+`pg`+"`"+`,  `+"`"+`redis`+"`"+`,  `+"`"+`kafka`+"`"+`, `+"`"+`opensearch`+"`"+`,  `+"`"+`mongodb`+"`"+``, ox.Default("mysql")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases options slugs
					ox.Usage(`slugs`, `Retrieves a list of the available slugs for a given database engine`),
					ox.Banner(`Lists the available slugs for a given database engine.`),
					ox.Spec(`[flags]`),
					ox.Aliases("s"),
					ox.Example(`
The following example retrieves a list of the available slugs for the PostgreSQL engine: doctl databases options slugs --engine pg`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`engine`, `The database engine. Possible values:  mysql,  `+"`"+`pg`+"`"+`,  `+"`"+`redis`+"`"+`,  `+"`"+`kafka`+"`"+`,  `+"`"+`opensearch`+"`"+`, `+"`"+`mongodb`+"`"+` (required)`, ox.Default("mysql")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases options versions
					ox.Usage(`versions`, `Retrieves a list of the available versions for a given database engine`),
					ox.Banner(`Lists the available versions for a given database engine.`),
					ox.Spec(`[flags]`),
					ox.Aliases("v"),
					ox.Example(`
The following example retrieves a list of the available versions for the PostgreSQL engine: doctl databases options versions --engine pg`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`engine`, `The database engine. Possible values:  mysql,  `+"`"+`pg`+"`"+`,  `+"`"+`redis`+"`"+`,  `+"`"+`kafka`+"`"+`,  `+"`"+`opensearch`+"`"+`, `+"`"+`mongodb`+"`"+``, ox.Default("mysql")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases pool
				ox.Usage(`pool`, `Display commands for managing connection pools`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases pool`+"`"+` manage connection pools for your database cluster.

A connection pool may be useful if your database:

- Typically handles a large number of idle connections
- Has wide variability in the possible number of connections at any given time
- Drops connections due to max connection limits
- Experiences performance issues due to high CPU usage`),
				ox.Spec(`[command]`),
				ox.Aliases("p"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases pool [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases pool create
					ox.Usage(`create`, `Create a connection pool for a database cluster`),
					ox.Banner(`Creates a connection pool for the specified database cluster.

In addition to the pool's name, you must also use flags to specify the pool's target database, its size, and a database user that the pool uses to authenticate. If you do not specify a user, the field is set to inbound user. An example call would be:

The pool size is the minimum number of connections the pool can handle. The maximum pool size varies based on the size of the cluster.

There’s no perfect formula to determine how large your pool should be, but there are a few good guidelines to keep in mind:

- A large pool stresses your database at similar levels as that number of clients would alone.
- A pool that’s much smaller than the number of clients communicating with the database can act as a bottleneck, reducing the rate when your database receives and responds to transactions.

We recommend starting with a pool size of about half your available connections and adjusting later based on performance. If you see slow query responses, check the CPU usage on the database’s Overview tab. We recommend decreasing your pool size if CPU usage is high, and increasing your pool size if it’s low.

You can get a list of existing connection pools by calling:

	doctl databases pool list <database-cluster-id>

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <pool-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a connection pool named `+"`"+`example-pool`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`. The command uses the `+"`"+`--size`+"`"+` flag to set the pool size to 10 and sets the user to the database's default user: doctl databases pool create ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-pool --size 10`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`db`, `The name of the specific database within the database cluster (required)`).
						String(`mode`, `The pool mode for the connection pool, such as session, `+"`"+`transaction`+"`"+`, and `+"`"+`statement`+"`"+``, ox.Default("transaction")).
						Int(`size`, `pool size (required)`).
						String(`user`, `The username for the database user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases pool delete
					ox.Usage(`delete`, `Delete a connection pool for a database`),
					ox.Banner(`Deletes the specified connection pool for the specified database cluster.

You can get a list of existing connection pools by calling:

	doctl databases pool list <database-cluster-id>

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <pool-name> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes a connection pool named `+"`"+`example-pool`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases pool delete ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-pool`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the connection pool without confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases pool get
					ox.Usage(`get`, `Retrieve information about a database connection pool`),
					ox.Banner(`This command retrieves the following information about the specified connection pool for the specified database cluster:

- The database user that the connection pool uses. When excluded, all connections to the database use the inbound user.
- The connection pool's name
- The connection pool's size
- The database within the cluster that the connection pool connects to
- The pool mode for the connection pool. Possible values: `+"`"+`session`+"`"+`, `+"`"+`transaction`+"`"+`, or `+"`"+`statement`+"`"+`
- A connection string for the connection pool

You can get a list of existing connection pools by calling:

	doctl databases pool list <database-cluster-id>

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <pool-name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the details for a connection pool named `+"`"+`example-pool`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the pool's name and connection string: doctl databases pool get ca9f591d-fb58-5555-a0ef-1c02d1d1e352 example-pool --format Name,URI`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: User, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Database`+"`"+`, `+"`"+`Mode`+"`"+`, `+"`"+`URI`+"`"+`.`, ox.Default("User")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases pool list
					ox.Usage(`list`, `List connection pools for a database cluster`),
					ox.Banner(`Lists the existing connection pools for the specified database. The command returns the following details about each connection pool:

- The database user that the connection pool uses. When excluded, all connections to the database use the inbound user.
- The connection pool's name
- The connection pool's size
- The database within the cluster that the connection pool connects to
- The pool mode for the connection pool. Possible values: `+"`"+`session`+"`"+`, `+"`"+`transaction`+"`"+`, or `+"`"+`statement`+"`"+`
- A connection string for the connection pool`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists the connection pools for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only each pool's name and connection string: doctl databases pool list ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --format Name,URI`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: User, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Database`+"`"+`, `+"`"+`Mode`+"`"+`, `+"`"+`URI`+"`"+`.`, ox.Default("User")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases pool update
					ox.Usage(`update`, `Update a connection pool for a database`),
					ox.Banner(`Updates the specified connection pool for the specified database cluster.

You can get a list of existing connection pools by calling:

	doctl databases pool list <database-cluster-id>

You can get a list of existing database clusters and their IDs by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <pool-name> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates a connection pool named `+"`"+`example-pool`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`. The command uses the `+"`"+`--size`+"`"+` flag to set the pool size to 10 and sets the user to the database's default user: doctl databases pool update ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-pool --size 10`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`db`, `The name of the specific database within the database cluster`).
						String(`mode`, `The pool mode for the connection pool, such as session, `+"`"+`transaction`+"`"+`, and `+"`"+`statement`+"`"+``, ox.Default("transaction")).
						Int(`size`, `pool size`).
						String(`user`, `The username for the database user`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases replica
				ox.Usage(`replica`, `Display commands to manage read-only database replicas`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases replica`+"`"+` allow you to manage read-only replicas associated with a database cluster.

In addition to primary nodes in a database cluster, you can create up to 2 read-only replica nodes (also referred to as "standby nodes") to maintain high availability.`),
				ox.Spec(`[command]`),
				ox.Aliases("rep", "r"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases replica [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases replica connection
					ox.Usage(`connection`, `Retrieve information for connecting to a read-only database replica`),
					ox.Banner(`Retrieves information for connecting to the specified read-only database replica in the specified database cluster

This command requires that you pass in the replica's name, which you can retrieve by querying a database ID:

	doctl databases replica list <database-cluster-id>

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <replica-name> [flags]`),
					ox.Aliases("conn"),
					ox.Example(`
The following example retrieves the connection details for a read-only replica named `+"`"+`example-replica`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases replica connection get ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-replica`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases replica create
					ox.Usage(`create`, `Create a read-only database replica`),
					ox.Banner(`Creates a read-only database replica for the specified database cluster, giving it the specified name.

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <replica-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a read-only replica named `+"`"+`example-replica`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases replica create ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-replica --size db-s-1vcpu-1gb`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`private-network-uuid`, `The UUID of a VPC to create the replica in; the default VPC for the region will be used if excluded.`).
						String(`region`, `Specifies the region in which to create the replica, such as nyc3 or `+"`"+`sfo2`+"`"+`.`, ox.Default("nyc1")).
						String(`size`, `Specifies the machine size for the replica, such as db-s-1vcpu-1gb. Must be the same size or larger than the primary database cluster.`, ox.Default("db-s-1vcpu-1gb")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases replica delete
					ox.Usage(`delete`, `Delete a read-only database replica`),
					ox.Banner(`Deletes the specified read-only replica for the specified database cluster.

This command requires that you pass in the replica's name, which you can retrieve by querying a database ID:

	doctl databases replica list <database-cluster-id>

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <replica-name> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes a read-only replica named `+"`"+`example-replica`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases replica delete ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-replica`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the replica without a confirmation prompt.`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases replica get
					ox.Usage(`get`, `Retrieve information about a read-only database replica`),
					ox.Banner(`Gets the following details for the specified read-only replica of the specified database cluster:

- The name of the replica
- The information required to connect to the read-only replica
- The region where the database cluster is located, such as `+"`"+`nyc3`+"`"+` or `+"`"+`sfo2`+"`"+`
- The status of the replica. Possible values: `+"`"+`creating`+"`"+`, `+"`"+`forking`+"`"+`, `+"`"+`active`+"`"+`
- When the read-only replica was created, in ISO8601 date/time format

This command requires that you pass in the replica's name, which you can retrieve by querying a database ID:

	doctl databases replica list <database-cluster-id>

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <replica-name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the details for a read-only replica named `+"`"+`example-replica`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases replica get ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-replica`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`ID`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`URI`+"`"+`, `+"`"+`Created`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases replica list
					ox.Usage(`list`, `Retrieve list of read-only database replicas`),
					ox.Banner(`Lists the following details for read-only replicas for the specified database cluster.

- The replica's name
- The region where the database cluster is located, such as `+"`"+`nyc3`+"`"+`, `+"`"+`sfo2`+"`"+`
- The replica's status. Possible values: `+"`"+`forking`+"`"+` and `+"`"+`active`+"`"+`


This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of read-only replicas for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the ID and URI for each replica: doctl databases replica list ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --format ID,URI`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`ID`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`URI`+"`"+`, `+"`"+`Created`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases replica promote
					ox.Usage(`promote`, `Promote a read-only database replica to become a primary cluster`),
					ox.Banner(`Promotes a read-only database replica to become its own independent primary cluster. Promoted replicas no longer stay in sync with primary cluster they were forked from.

This command requires that you pass in the replica's name, which you can retrieve by querying a database ID:

	doctl databases replica list <database-cluster-id>

This command requires the ID of a database cluster, which you can retrieve by calling:

	doctl databases list`),
					ox.Spec(`<database-cluster-id> <replica-name> [flags]`),
					ox.Aliases("p"),
					ox.Example(`
The following example promotes a read-only replica named `+"`"+`example-replica`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases replica promote ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-replica`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases resize
				ox.Usage(`resize`, `Resize a database cluster`),
				ox.Banner(`Resizes the specified database cluster.

You must specify the desired number of nodes and size of the nodes. For example:

	doctl databases resize ca9f591d-9999-5555-a0ef-1c02d1d1e352 --num-nodes 2 --size db-s-16vcpu-64gb

Database nodes cannot be resized to smaller sizes due to the risk of data loss.

For PostgreSQL and MySQL clusters, you can also provide a disk size in MiB to scale the storage up to 15 TB, depending on your plan. You cannot reduce the storage size of a cluster.`),
				ox.Spec(`<database-cluster-id> [flags]`),
				ox.Aliases("rs"),
				ox.Example(`
The following example resizes a PostgreSQL or MySQL database to have two nodes, 16 vCPUs, 64 GB of memory, and 2048 GiB of storage space: doctl databases resize ca9f591d-9999-5555-a0ef-1c02d1d1e352 --num-nodes 2 --size db-s-16vcpu-64gb --storage-size-mib 2048000 --wait true`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`num-nodes`, `The number of nodes in the database cluster. Valid values are 1-3. In addition to the primary node, up to two standby nodes may be added for high availability. (required)`).
					String(`size`, `The size of the nodes in the database cluster, for example db-s-1vcpu-1gb indicates a 1 CPU, 1GB node. For a list of available size slugs, visit: https://docs.digitalocean.com/reference/api/api-reference/#tag/Databases (required)`, ox.Default("db-s-1vcpu-1gb")).
					Int(`storage-size-mib`, `The amount of disk space allocated to the cluster. Applicable for PostgreSQL and MySQL clusters. Each plan size has a default value but can be increased in increments up to a maximum amount. For ranges, visit: https://www.digitalocean.com/pricing/managed-databases`).
					Bool(`wait`, `Boolean that specifies whether to wait for the resize to complete before returning control to the terminal`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl databases sql-mode
				ox.Usage(`sql-mode`, `Display commands to configure a MySQL database cluster's SQL modes`),
				ox.Banner(`The subcommands of `+"`"+`doctl databases sql-mode`+"`"+` are used to view and configure a MySQL database cluster's global SQL modes. Global SQL modes affect the SQL syntax MySQL supports and the data validation checks it performs.`),
				ox.Spec(`[command]`),
				ox.Aliases("sm"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases sql-mode [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases sql-mode get
					ox.Usage(`get`, `Get a MySQL database cluster's SQL modes`),
					ox.Banner(`Displays the configured SQL modes for the specified MySQL database cluster.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the SQL modes for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases sql-mode get ca9f591d-f38h-5555-a0ef-1c02d1d1e35`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases sql-mode set
					ox.Usage(`set`, `Set a MySQL database cluster's SQL modes`),
					ox.Banner(`This command configures the SQL modes for the specified MySQL database cluster. The SQL modes should be provided as a space separated list.

This command replaces the existing SQL mode configuration completely. Include all of the current values when adding a new one.`),
					ox.Spec(`<database-cluster-id> <sql-mode-1> ... <sql-mode-n> [flags]`),
					ox.Aliases("s"),
					ox.Example(`
The following example sets the SQL mode ALLOW_INVALID_DATES for an existing database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`. The cluster already has the modes `+"`"+`NO_ZERO_DATE`+"`"+`, `+"`"+`NO_ZERO_IN_DATE`+"`"+`, `+"`"+`STRICT_ALL_TABLES`+"`"+` set, but they must be included in the command to avoid being overwritten by the additional mode: doctl databases sql-mode set ca9f591d-f38h-5555-a0ef-1c02d1d1e35 NO_ZERO_DATE NO_ZERO_IN_DATE STRICT_ALL_TABLES ALLOW_INVALID_DATES`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases topics
				ox.Usage(`topics`, `Display commands to manage topics for kafka database clusters`),
				ox.Banner(`The subcommands under `+"`"+`doctl databases topics`+"`"+` enable the management of topics for kafka database clusters`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases topics [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases topics create
					ox.Usage(`create`, `Creates a topic for a given kafka database`),
					ox.Banner(`This command creates a kafka topic for the specified kafka database cluster, giving it the specified name. Example: doctl databases topics create <database-uuid> <topic-name> --replication_factor 2 --partition_count 4`),
					ox.Spec(`<database-uuid> <topic-name> [flags]`),
					ox.Aliases("c"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cleanup-policy`, `Specifies the retention policy to use on log segments: Possible values are 'delete', 'compact_delete', 'compact'`, ox.Default("delete")).
						String(`compression-type`, `Specifies the compression type for a kafka topic: Possible values are 'producer', 'gzip', 'snappy', 'Iz4', 'zstd', 'uncompressed'`, ox.Default("producer")).
						String(`delete-retention-ms`, `Specifies how long (in ms) to retain delete tombstone markers for topics`).
						String(`file-delete-delay-ms`, `Specifies the minimum time (in ms) to wait before deleting a file from the filesystem`).
						String(`flush-messages`, `Specifies the maximum number of messages to accumulate on a log partition before messages are flushed to disk`).
						String(`flush-ms`, `Specifies the maximum time (in ms) that a message is kept in memory before being flushed to disk`).
						String(`interval-index-bytes`, `Specifies the number of bytes between entries being added into the offset index`).
						String(`max-compaction-lag-ms`, `Specifies the maximum time (in ms) that a message will remain uncompacted. This is only applicable if the logs have compaction enabled`).
						String(`max-message-bytes`, `Specifies the largest record batch (in bytes) that can be sent to the server. This is calculated after compression, if compression is enabled`).
						Bool(`message-down-conversion-enable`, `Specifies whether down-conversion of message formats is enabled to satisfy consumer requests`, ox.Default("true")).
						String(`message-format-version`, `Specifies the message format version used by the broker to append messages to the logs. By setting a format version, all existing messages on disk must be smaller or equal to the specified version`).
						String(`message-timestamp-type`, `Specifies whether to use the create time or log append time as the timestamp on a message`).
						String(`min-cleanable-dirty-ratio`, `Specifies the frequenty of log compaction (if enabled) in relation to duplicates present in the logs. For example, 0.5 would mean at most half of the log could be duplicates before compaction would begin`).
						String(`min-compaction-lag-ms`, `Specifies the minimum time (in ms) that a message will remain uncompacted. This is only applicable if the logs have compaction enabled`).
						String(`min-insync-replicas`, `Specifies the minimum number of replicas that must ACK a write for it to be considered successful`).
						Int(`partition-count`, `Specifies the number of partitions available for the topic`, ox.Default("1")).
						Bool(`preallocate`, `Specifies whether a file should be preallocated on disk when creating a new log segment`).
						Int(`replication-factor`, `Specifies the number of nodes to replicate data across the kafka cluster`, ox.Default("2")).
						String(`retention-bytes`, `Specifies the maximum size (in bytes) before deleting messages. '-1' indicates that there is no limit`).
						String(`retention-ms`, `Specifies the maximum time (in ms) to store a message before deleting it. '-1' indicates that there is no limit`).
						String(`segment-bytes`, `Specifies the maximum size (in bytes) of a single log file`).
						String(`segment-jitter-ms`, `Specifies the maximum time (in ms) for random jitter that is subtracted from the scheduled segment roll time to avoid thundering herd problems`).
						String(`segment-ms`, `Specifies the maximum time (in ms) to wait to force a log roll if the segment file isn't full. After this period, the log will be forced to roll`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases topics delete
					ox.Usage(`delete`, `Deletes a kafka topic by topic name`),
					ox.Banner(`Deletes a kafka topic by topic name`),
					ox.Spec(`<database-uuid> <topic-name> [flags]`),
					ox.Aliases("rm"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the kafka topic without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases topics get
					ox.Usage(`get`, `Retrieve the configuration for a given kafka topic`),
					ox.Banner(`
This command lists the following details for a given topic in a kafka database cluster:

	- The Name of the topic.
	- The Partitions of the topic - the number of partitions in the topics
	- The Replication Factor of the topic - number of brokers the topic's partitions are replicated across.
	- Additional advanced configuration for the topic.

The details of the topic are listed in key/value pairs`),
					ox.Spec(`<database-uuid> <topic-name> [flags]`),
					ox.Aliases("g"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: key, `+"`"+`value`+"`"+`.`, ox.Default("key")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases topics list
					ox.Usage(`list`, `Retrieve a list of topics for a given kafka database`),
					ox.Banner(`
This command lists the following details for each topic in a kafka database cluster:

	- The Name of the topic.
	- The State of the topic.
	- The Replication Factor of the topic - number of brokers the topic's partitions are replicated across.`),
					ox.Spec(`<database-uuid> [flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`State`+"`"+`, `+"`"+`ReplicationFactor`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases topics partitions
					ox.Usage(`partitions`, `Retrieve the partitions for a given kafka topic`),
					ox.Banner(`
This command lists the following details for each partition of a given topic in a kafka database cluster:

	- The Id - identifier of the topic partition.
	- The Size - size of the topic partition, in bytes.
	- The InSyncReplicas - number of brokers that are in sync with the partition leader.
	- The EarliestOffset - earliest offset read amongst all consumers of the partition.`),
					ox.Spec(`<database-id> <topic-name> [flags]`),
					ox.Aliases("p"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases topics update
					ox.Usage(`update`, `Updates a topic for a given kafka database`),
					ox.Banner(`This command updates a kafka topic for the specified kafka database cluster. Example: doctl databases topics update <database-uuid> <topic-name>`),
					ox.Spec(`<database-uuid> <topic-name> [flags]`),
					ox.Aliases("u"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cleanup-policy`, `Specifies the retention policy to use on log segments: Possible values are 'delete', 'compact_delete', 'compact'`, ox.Default("delete")).
						String(`compression-type`, `Specifies the compression type for a kafka topic: Possible values are 'producer', 'gzip', 'snappy', 'Iz4', 'zstd', 'uncompressed'`, ox.Default("producer")).
						String(`delete-retention-ms`, `Specifies how long (in ms) to retain delete tombstone markers for topics`).
						String(`file-delete-delay-ms`, `Specifies the minimum time (in ms) to wait before deleting a file from the filesystem`).
						String(`flush-messages`, `Specifies the maximum number of messages to accumulate on a log partition before messages are flushed to disk`).
						String(`flush-ms`, `Specifies the maximum time (in ms) that a message is kept in memory before being flushed to disk`).
						String(`interval-index-bytes`, `Specifies the number of bytes between entries being added into the offset index`).
						String(`max-compaction-lag-ms`, `Specifies the maximum time (in ms) that a message will remain uncompacted. This is only applicable if the logs have compaction enabled`).
						String(`max-message-bytes`, `Specifies the largest record batch (in bytes) that can be sent to the server. This is calculated after compression, if compression is enabled`).
						Bool(`message-down-conversion-enable`, `Specifies whether down-conversion of message formats is enabled to satisfy consumer requests`, ox.Default("true")).
						String(`message-format-version`, `Specifies the message format version used by the broker to append messages to the logs. By setting a format version, all existing messages on disk must be smaller or equal to the specified version`).
						String(`message-timestamp-type`, `Specifies whether to use the create time or log append time as the timestamp on a message`).
						String(`min-cleanable-dirty-ratio`, `Specifies the frequenty of log compaction (if enabled) in relation to duplicates present in the logs. For example, 0.5 would mean at most half of the log could be duplicates before compaction would begin`).
						String(`min-compaction-lag-ms`, `Specifies the minimum time (in ms) that a message will remain uncompacted. This is only applicable if the logs have compaction enabled`).
						String(`min-insync-replicas`, `Specifies the minimum number of replicas that must ACK a write for it to be considered successful`).
						Int(`partition-count`, `Specifies the number of partitions available for the topic`, ox.Default("1")).
						Bool(`preallocate`, `Specifies whether a file should be preallocated on disk when creating a new log segment`).
						Int(`replication-factor`, `Specifies the number of nodes to replicate data across the kafka cluster`, ox.Default("2")).
						String(`retention-bytes`, `Specifies the maximum size (in bytes) before deleting messages. '-1' indicates that there is no limit`).
						String(`retention-ms`, `Specifies the maximum time (in ms) to store a message before deleting it. '-1' indicates that there is no limit`).
						String(`segment-bytes`, `Specifies the maximum size (in bytes) of a single log file`).
						String(`segment-jitter-ms`, `Specifies the maximum time (in ms) for random jitter that is subtracted from the scheduled segment roll time to avoid thundering herd problems`).
						String(`segment-ms`, `Specifies the maximum time (in ms) to wait to force a log roll if the segment file isn't full. After this period, the log will be forced to roll`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl databases user
				ox.Usage(`user`, `Display commands for managing database users`),
				ox.Banner(`The commands under `+"`"+`doctl databases user`+"`"+` allow you to view details for, and create, database users.

Database user accounts are scoped to one database cluster, to which they have full admin access, and are given an automatically-generated password.`),
				ox.Spec(`[command]`),
				ox.Aliases("u"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl databases user [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl databases user create
					ox.Usage(`create`, `Create a database user`),
					ox.Banner(`Creates a new user for a database. New users are given a role of `+"`"+`normal`+"`"+` and are given an automatically-generated password.

To retrieve a list of your databases and their IDs, call `+"`"+`doctl databases list`+"`"+`.`),
					ox.Spec(`<database-cluster-id> <user-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a new user with the username `+"`"+`example-user`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases user create ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-user`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`acl`, `A comma-separated list of kafka ACL rules, in topic:permission format.`, ox.Spec(`topic:permission`)).
						String(`mysql-auth-plugin`, `Sets authorization plugin for a MySQL user. Possible values: caching_sha2_password or `+"`"+`mysql_native_password`+"`"+``, ox.Default("caching_sha2_password")).
						String(`opensearch-acl`, `A comma-separated list of OpenSearch ACL rules, in index:permission format.`, ox.Spec(`index:permission`)).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases user delete
					ox.Usage(`delete`, `Delete a database user`),
					ox.Banner(`Deletes the specified database user.

To retrieve a list of your databases and their IDs, call `+"`"+`doctl databases list`+"`"+`.`),
					ox.Spec(`<database-cluster-id> <user-id> [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes the user with the username `+"`"+`example-user`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+`: doctl databases user delete ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-user`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the user without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases user get
					ox.Usage(`get`, `Retrieve details about a database user`),
					ox.Banner(`Retrieves the following details about the specified user:

- The username for the user
- The password for the user
- The user's role, either "primary" or "normal"

Primary user accounts are created by DigitalOcean at database cluster creation time and can't be deleted. You can create additional users with a "normal" role. Both have administrative privileges on the database cluster.

To retrieve a list of your databases and their IDs, call `+"`"+`doctl databases list`+"`"+`.

To retrieve a list of database users for a database cluster, call `+"`"+`doctl databases user list <database-cluster-id>`+"`"+`.`),
					ox.Spec(`<database-cluster-id> <user-name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the details for the user with the username `+"`"+`example-user`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the user's name and role: doctl databases user get ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-user --format Name,Role`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`Role`+"`"+`, `+"`"+`Password`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases user list
					ox.Usage(`list`, `Retrieve list of database users`),
					ox.Banner(`Retrieves a list of users for the specified database with the following details:

- The username for the user
- The password for the user
- The user's role, either "primary" or "normal"

Primary user accounts are created by DigitalOcean at database cluster creation time and can't be deleted. You can create additional users with a "normal" role. Both have administrative privileges on the database cluster.

To retrieve a list of your databases and their IDs, call `+"`"+`doctl databases list`+"`"+`.`),
					ox.Spec(`<database-cluster-id> [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of users for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` and uses the `+"`"+`--format flag`+"`"+` to return only the name and role for each each user: doctl databases user list ca9f591d-f38h-5555-a0ef-1c02d1d1e35 --format Name,Role`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`Role`+"`"+`, `+"`"+`Password`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl databases user reset
					ox.Usage(`reset`, `Resets a user's auth`),
					ox.Banner(`Resets the auth password or the MySQL authorization plugin for a given user and returns the user's new credentials. When resetting MySQL auth, valid values for `+"`"+`<new-auth-mode>`+"`"+` are `+"`"+`caching_sha2_password`+"`"+` and `+"`"+`mysql_native_password`+"`"+`.`),
					ox.Spec(`<database-cluster-id> <user-name> <new-auth-mode> [flags]`),
					ox.Aliases("rs"),
					ox.Example(`
The following example resets the auth plugin for the user with the username `+"`"+`example-user`+"`"+` for a database cluster with the ID `+"`"+`ca9f591d-f38h-5555-a0ef-1c02d1d1e35`+"`"+` to `+"`"+`mysql_native_password`+"`"+`: doctl databases user reset ca9f591d-f38h-5555-a0ef-1c02d1d1e35 example-user mysql_native_password`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
		),
		ox.Sub( // doctl kubernetes
			ox.Usage(`kubernetes`, `Displays commands to manage Kubernetes clusters and configurations`),
			ox.Banner(`The commands under `+"`"+`doctl kubernetes`+"`"+` are for managing Kubernetes clusters and viewing configuration options relating to clusters.

A typical workflow is to use `+"`"+`doctl kubernetes cluster create`+"`"+` to create the cluster on DigitalOcean's infrastructure, then call `+"`"+`doctl kubernetes cluster kubeconfig`+"`"+` to configure `+"`"+`kubectl`+"`"+` to connect to the cluster. You are then able to use `+"`"+`kubectl`+"`"+` to create and manage workloads.

The commands under `+"`"+`doctl kubernetes options`+"`"+` retrieve values used while creating clusters, such as the list of regions where cluster creation is supported.`),
			ox.Spec(`[command]`),
			ox.Aliases("kube", "k8s", "k"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl kubernetes [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl kubernetes 1-click
				ox.Usage(`1-click`, `Display commands that pertain to kubernetes 1-click applications`),
				ox.Banner(`The commands under `+"`"+`doctl kubernetes 1-click`+"`"+` are for managing DigitalOcean Kubernetes 1-Click applications.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl kubernetes 1-click [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl kubernetes 1-click install
					ox.Usage(`install`, `Install 1-click apps on a Kubernetes cluster`),
					ox.Banner(`Installs 1-click applications on a Kubernetes cluster. Use the `+"`"+`--1-click`+"`"+` flag to specify one or multiple pieces of software to install on the cluster.`),
					ox.Spec(`<cluster-id> [flags]`),
					ox.Aliases("in"),
					ox.Example(`
The following example installs Loki and Netdata on a Kubernetes cluster with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl kubernetes 1-click install f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --1-clicks loki,netdata`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`1-clicks`, `moon,loki,netdata   The 1-clicks application to install on the cluster. Multiple 1-clicks can be added simultaneously, for example: --1-clicks moon,loki,netdata`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: SLUG, `+"`"+`TYPE`+"`"+`.`, ox.Spec(`SLUG`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes 1-click list
					ox.Usage(`list`, `Retrieve a list of Kubernetes 1-Click applications`),
					ox.Banner(`Retrieves a list of Kubernetes 1-Click applications you can install on a Kubernetes cluster.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all available 1-click apps for Kubernetes: doctl kubernetes 1-click list`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: SLUG, `+"`"+`TYPE`+"`"+`.`, ox.Spec(`SLUG`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl kubernetes cluster
				ox.Usage(`cluster`, `Display commands for managing Kubernetes clusters`),
				ox.Banner(`The commands under `+"`"+`doctl kubernetes cluster`+"`"+` are for the management of Kubernetes clusters.

A typical workflow is to use `+"`"+`doctl kubernetes cluster create`+"`"+` to create the cluster on DigitalOcean's infrastructure, then call `+"`"+`doctl kubernetes cluster kubeconfig`+"`"+` to configure `+"`"+`kubectl`+"`"+` to connect to the cluster. You are then able to use `+"`"+`kubectl`+"`"+` to create and manage workloads.`),
				ox.Spec(`[command]`),
				ox.Aliases("clusters", "c"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl kubernetes cluster [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl kubernetes cluster create
					ox.Usage(`create`, `Create a Kubernetes cluster`),
					ox.Banner(`
Creates a Kubernetes cluster given the specified options and using the specified name. Before creating the cluster, you can use `+"`"+`doctl kubernetes options`+"`"+` to see possible values for the various configuration flags.

If no configuration flags are used, a three-node cluster with a single node pool is created in the `+"`"+`nyc1`+"`"+` region, using the latest Kubernetes version.

After creating a cluster, a configuration context is added to kubectl and made active so that you can begin managing your new cluster immediately.`),
					ox.Spec(`<name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a cluster named `+"`"+`example-cluster`+"`"+` in the `+"`"+`nyc1`+"`"+` region with a node pool, using Kubernetes version `+"`"+`1.28.2-do.0`+"`"+`: doctl kubernetes cluster create example-cluster --region nyc1 --version 1.28.2-do.0 --maintenance-window saturday=02:00 --node-pool "name=example-pool;size=s-2vcpu-2gb;count=5;tag=web;tag=frontend;label=key1=value1;label=key2=label2;taint=key1=value1:NoSchedule;taint=key2:NoExecute"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`1-clicks`, `A comma-separated list of 1-click applications to install on the Kubernetes cluster. Use the `+"`"+`doctl kubernetes 1-click list`+"`"+` command for a list of available 1-click applications.`).
						Bool(`auto-upgrade`, `Enables automatic upgrades to new patch releases during the cluster's maintenance window. Defaults to false. To enable automatic upgrade, supply `+"`"+`--auto-upgrade=true`+"`"+`.`, ox.Spec(`false`)).
						CIDR(`cluster-subnet`, `The CIDR block to use for the pod network. Must be a valid CIDR block. Defaults to 10.244.0.0/16. If left empty/default the cluster will be created with a virtual network. If a custom one is provided, the cluster will be created as vpc-native cluster. VPC-native CIDR blocks cannot overlap within an account.`, ox.Default("10.244.0.0/16")).
						Slice(`control-plane-firewall-allowed-addresses`, `A comma-separated list of allowed addresses that can access the control plane.`).
						Int(`count`, `The number of nodes in the default node pool (incompatible with --node-pool)`, ox.Default("3")).
						Bool(`enable-control-plane-firewall`, `Creates the cluster with control plane firewall enabled. Defaults to false. To enable the control plane firewall, supply --enable-control-plane-firewall=true.`).
						Bool(`ha`, `Creates the cluster with a highly-available control plane. Defaults to false. To enable the HA control plane, supply --ha=true.`).
						String(`maintenance-window`, `Sets the beginning of the schedule for the four hour maintenance window for the cluster. The syntax format is: `+"`"+`day=HH:MM`+"`"+`, where time is in UTC. Day can be: `+"`"+`any`+"`"+`, `+"`"+`monday`+"`"+`, `+"`"+`tuesday`+"`"+`, `+"`"+`wednesday`+"`"+`, `+"`"+`thursday`+"`"+`, `+"`"+`friday`+"`"+`, `+"`"+`saturday`+"`"+`, `+"`"+`sunday`+"`"+`.`, ox.Spec(`schedule`), ox.Default("any=00:00")).
						Slice(`node-pool`, `A comma-separated list of node-pools, defined using semicolon-separated configuration values and surrounded by quotes (incompatible with --size and --count).`).
						String(`region`, `A slug indicating which region to create the cluster in. Use the `+"`"+`doctl kubernetes options regions`+"`"+` command for a list of options (required)`, ox.Default("nyc1")).
						CIDR(`service-subnet`, `The CIDR block to use for the service network. Must be a valid CIDR block. Defaults to 10.245.0.0/16. If left empty/default the cluster will be created with a virtual network. If a custom one is provided, the cluster will be created as vpc-native cluster. VPC-native CIDR blocks cannot overlap within an account.`, ox.Default("10.245.0.0/16")).
						Bool(`set-current-context`, `Sets the current kubectl context to that of the new cluster`, ox.Default("true")).
						String(`size`, `The machine size to use when creating nodes in the default node pool (incompatible with --node-pool). Use the `+"`"+`doctl kubernetes options sizes`+"`"+` command for a list of possible values.`, ox.Spec(`size`), ox.Default("s-1vcpu-2gb")).
						Bool(`surge-upgrade`, `Enables surge-upgrade for the cluster`, ox.Default("true")).
						Slice(`tag`, `A comma-separated list of tags to apply to the cluster, in addition to the default tags of `+"`"+`k8s`+"`"+` and `+"`"+`k8s:$K8S_CLUSTER_ID`+"`"+`.`).
						Bool(`update-kubeconfig`, `Adds a configuration context for the new cluster to your kubectl`, ox.Default("true")).
						String(`vpc-uuid`, `The UUID of a VPC network to create the cluster in. Must be the UUID of a valid VPC in the same region specified for the cluster. If a VPC is not specified, the cluster is placed in the default VPC network for the region.`).
						Bool(`wait`, `Instructs the terminal to wait for the action to complete before returning control to the user`, ox.Default("true")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster delete
					ox.Usage(`delete`, `Delete Kubernetes clusters`),
					ox.Banner(`
Deletes the specified Kubernetes clusters and the Droplets associated with them. To delete all other DigitalOcean resources created during the operation of the clusters, such as load balancers, volumes or volume snapshots, use the `+"`"+`--dangerous`+"`"+` flag.`),
					ox.Spec(`<id|name>... [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster delete example-cluster`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`dangerous`, `Deletes the cluster's associated resources like load balancers, volumes and volume snapshots`).
						Bool(`force`, `Deletes the cluster without a confirmation prompt`, ox.Short("f")).
						Bool(`update-kubeconfig`, `Remove the deleted cluster from your kubeconfig`, ox.Default("true")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster delete-selective
					ox.Usage(`delete-selective`, `Delete a Kubernetes cluster and selectively delete resources associated with it`),
					ox.Banner(`
Deletes the specified Kubernetes cluster and Droplets associated with it. It also deletes the specified associated resources. Associated resources can be load balancers, volumes and volume snapshots.`),
					ox.Spec(`<id|name> [flags]`),
					ox.Aliases("ds"),
					ox.Example(`
The following example deletes a cluster named `+"`"+`example-cluster`+"`"+` and selectively deletes the specified load balancers and volumes associated with the cluster: doctl kubernetes cluster delete-selective example-cluster --volume-list "386734086,example-volume" --load-balancer-list "191669331,example-load-balancer"`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes the cluster without a confirmation prompt`, ox.Short("f")).
						Slice(`load-balancers`, `A comma-separated list of load balancer IDs or names to delete`).
						Slice(`snapshots`, `A comma-separated list of volume snapshot IDs or names to delete`).
						Bool(`update-kubeconfig`, `Removes the deleted cluster from your kubeconfig`, ox.Default("true")).
						Slice(`volumes`, `A comma-separated list of volume IDs or names to delete`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster get
					ox.Usage(`get`, `Retrieve details about a Kubernetes cluster`),
					ox.Banner(`
Retrieves the following details about a Kubernetes cluster: 

- A unique ID for the cluster
- A human-readable name for the cluster
- The slug identifying the region where the Kubernetes cluster is located
- The slug identifying the cluster's Kubernetes version. If set to a minor version, the latest patch version for that minor version is returned. For example, if the cluster is set to  `+"`"+`1.14`+"`"+`, the command would return `+"`"+`1.14.6-do.1`+"`"+`. If it is set to `+"`"+`latest`+"`"+`, the latest published version is used.
- A boolean value indicating whether the cluster automatically upgrades to new patch releases during its maintenance window.
- An object containing a "state" attribute whose value is set to a string indicating the current status of the node. Potential values: `+"`"+`running`+"`"+`, `+"`"+`provisioning`+"`"+`, `+"`"+`errored`+"`"+`.
- The base URL of the cluster's Kubernetes API server
- The public IPv4 address of the cluster's Kubernetes API server
- The range of IP addresses in the overlay network of the Kubernetes cluster, in CIDR notation
- The range of assignable IP addresses for services running in the Kubernetes cluster, in CIDR notation
- An array of tags applied to the Kubernetes cluster. All clusters are automatically tagged `+"`"+`k8s`+"`"+` and `+"`"+`k8s:$K8S_CLUSTER_ID`+"`"+`.
- When the Kubernetes cluster was created, in ISO8601 combined date and time format
- When the Kubernetes cluster was last updated, in ISO8601 combined date and time format
- A list of node pools available inside the cluster`),
					ox.Spec(`<id|name> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieve details about a Kubernetes cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster get example-cluster`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`AutoUpgrade`+"`"+`, `+"`"+`HAControlPlane`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`IPv4`+"`"+`, `+"`"+`ClusterSubnet`+"`"+`, `+"`"+`ServiceSubnet`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`, `+"`"+`NodePools`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster get-upgrades
					ox.Usage(`get-upgrades`, `Retrieve a list of available Kubernetes version upgrades`),
					ox.Banner(`
Retrieves a list of slugs representing Kubernetes upgrade versions you can use to upgrade the cluster. To upgrade your cluster, use the `+"`"+`doctl kubernetes cluster upgrade`+"`"+` command.`),
					ox.Spec(`<id|name> [flags]`),
					ox.Aliases("gu"),
					ox.Example(`
The following example retrieves a list of available Kubernetes version upgrades for a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster get-upgrades example-cluster`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster kubeconfig
					ox.Usage(`kubeconfig`, `Display commands for managing your local kubeconfig`),
					ox.Banner(`The commands under `+"`"+`doctl kubernetes cluster kubeconfig`+"`"+` are used to manage Kubernetes cluster credentials on your local machine. The credentials are used as authentication contexts with `+"`"+`kubectl`+"`"+`, the Kubernetes command-line interface.`),
					ox.Spec(`[command]`),
					ox.Aliases("kubecfg", "k8scfg", "config", "cfg"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl kubernetes cluster kubeconfig [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl kubernetes cluster kubeconfig remove
						ox.Usage(`remove`, `Remove a cluster's credentials from your local kubeconfig`),
						ox.Banner(`
This command removes the specified cluster's credentials from your local kubeconfig. After running this command, you cannot use `+"`"+`kubectl`+"`"+` to interact with your cluster.`),
						ox.Spec(`<cluster-id|cluster-name> [flags]`),
						ox.Aliases("d", "rm"),
						ox.Example(`
The following example removes the credentials for a cluster named `+"`"+`example-cluster`+"`"+` from your local kubeconfig: doctl kubernetes cluster kubeconfig remove example-cluster`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster kubeconfig save
						ox.Usage(`save`, `Save a cluster's credentials to your local kubeconfig`),
						ox.Banner(`
Adds the credentials for the specified cluster to your local kubeconfig. After this, your kubectl installation can directly manage the specified cluster.`),
						ox.Spec(`<cluster-id|cluster-name> [flags]`),
						ox.Aliases("s"),
						ox.Example(`
The following example saves the credentials for a cluster named `+"`"+`example-cluster`+"`"+` to your local kubeconfig: doctl kubernetes cluster kubeconfig save example-cluster`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`alias`, `An alias for the cluster context name. Defaults to 'do-[region]-[cluster-name]'`).
							Int(`expiry-seconds`, `The length of time the cluster credentials are valid for, in seconds. By default, the credentials are automatically renewed as needed.`).
							Bool(`set-current-context`, `Sets the current kubectl context to that of the newest cluster in your account`, ox.Default("true")).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster kubeconfig show
						ox.Usage(`show`, `Show a Kubernetes cluster's kubeconfig YAML`),
						ox.Banner(`
Returns the raw YAML for the specified cluster's kubeconfig.`),
						ox.Spec(`<cluster-id|cluster-name> [flags]`),
						ox.Aliases("p", "g"),
						ox.Example(`
The following example shows the kubeconfig YAML for a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster kubeconfig show example-cluster`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Int(`expiry-seconds`, `The length of time the cluster credentials are valid for, in seconds. By default, the credentials expire after seven days.`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl kubernetes cluster list
					ox.Usage(`list`, `Retrieve the list of Kubernetes clusters for your account`),
					ox.Banner(`
Retrieves the following details about all Kubernetes clusters that are on your account:

- A unique ID for the cluster
- A human-readable name for the cluster
- The slug identifying the region where the Kubernetes cluster is located
- The slug identifying the cluster's Kubernetes version. If set to a minor version, the latest patch version for that minor version is returned. For example, if the cluster is set to  `+"`"+`1.14`+"`"+`, the command would return `+"`"+`1.14.6-do.1`+"`"+`. If it is set to `+"`"+`latest`+"`"+`, the latest published version is used.
- A boolean value indicating whether the cluster automatically upgrades to new patch releases during its maintenance window.
- An object containing a "state" attribute whose value is set to a string indicating the current status of the node. Potential values: `+"`"+`running`+"`"+`, `+"`"+`provisioning`+"`"+`, `+"`"+`errored`+"`"+`.- A list of node pools available inside the cluster`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves the list of Kubernetes clusters for your account and uses the `+"`"+`--format`+"`"+` flag to return only the name and endpoint for each cluster: doctl kubernetes cluster list --format Name,Endpoint`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`AutoUpgrade`+"`"+`, `+"`"+`HAControlPlane`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Endpoint`+"`"+`, `+"`"+`IPv4`+"`"+`, `+"`"+`ClusterSubnet`+"`"+`, `+"`"+`ServiceSubnet`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Updated`+"`"+`, `+"`"+`NodePools`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster list-associated-resources
					ox.Usage(`list-associated-resources`, `Retrieve DigitalOcean resources associated with a Kubernetes cluster`),
					ox.Banner(`
Retrieves the following details for a Kubernetes cluster:
- Volume IDs for volumes created by the DigitalOcean CSI driver
- Volume snapshot IDs for volume snapshots created by the DigitalOcean CSI driver
- Load balancer IDs for load balancers managed by the Kubernetes cluster`),
					ox.Spec(`<id|name> [flags]`),
					ox.Aliases("ar"),
					ox.Example(`
The following example retrieves the associated resources for a cluster named `+"`"+`example-cluster`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the associated volumes: doctl kubernetes cluster list-associated-resources example-cluster --format Volumes`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Volumes, `+"`"+`VolumeSnapshots`+"`"+`, `+"`"+`LoadBalancers`+"`"+`.`, ox.Default("Volumes")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster node-pool
					ox.Usage(`node-pool`, `Display commands for managing node pools`),
					ox.Banner(`The commands under `+"`"+`node-pool`+"`"+` are for managing Kubernetes cluster's node pools. You can use these commands to create or delete node pools, enable autoscaling, and more.`),
					ox.Spec(`[command]`),
					ox.Aliases("node-pools", "nodepool", "nodepools", "pool", "pools", "np", "p"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl kubernetes cluster node-pool [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl kubernetes cluster node-pool create
						ox.Usage(`create`, `Create a new node pool for a cluster`),
						ox.Banner(`
Creates a new node pool for the specified cluster. The command requires values for the `+"`"+`--name`+"`"+`, `+"`"+`--size`+"`"+`, and `+"`"+`--count`+"`"+` flags to create a node pool. You can also specify that you'd like to enable autoscaling and set minimum and maximum node poll sizes.`),
						ox.Spec(`<cluster-id|cluster-name> [flags]`),
						ox.Aliases("c"),
						ox.Example(`
The following example creates a node pool named `+"`"+`example-pool`+"`"+` in a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster node-pool create example-cluster --name example-pool --size s-1vcpu-2gb --count 3 --taint "key1=value1:NoSchedule" --taint "key2:NoExecute"`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`auto-scale`, `Enables auto-scaling on the node pool`).
							Int(`count`, `The number of nodes in the node pool (required)`).
							Slice(`label`, `A label in key=value notation to apply to the node pool. Repeat this flag to specify additional labels. An existing label is removed from the node pool if it is not specified by any flag.`).
							Int(`max-nodes`, `The maximum number of nodes in the node pool when autoscaling is enabled`).
							Int(`min-nodes`, `The minimum number of nodes in the node pool when autoscaling (auto_scale) is enabled. If auto_scale is set to false, the default value will be 0. Scale-to-zero is not supported.`).
							String(`name`, `The name of the node pool (required)`).
							String(`size`, `The size of the nodes in the node pool. Use the `+"`"+`doctl kubernetes options sizes`+"`"+` command for a list of possible values. (required)`, ox.Spec(`size`)).
							Slice(`tag`, `A tag to apply to the node pool. Repeat this flag to specify additional tags. An existing tag is removed from the node pool if it is not specified by any flag.`).
							Map(`taint`, `Taint in key=value:effect notation to apply to the node pool. Repeat this flag to specify additional taints. Set to an empty string ("") to clear all taints. An existing taint is removed from the node pool if it is not specified by any flag.`, ox.Spec(`key=value:effect`)).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool delete
						ox.Usage(`delete`, `Delete a node pool`),
						ox.Banner(`Deletes a node pool in a cluster, which also removes all the nodes inside that pool. You cannot reverse this action.`),
						ox.Spec(`<cluster-id|cluster-name> <pool-id|pool-name> [flags]`),
						ox.Aliases("d", "rm"),
						ox.Example(`
The following example deletes a node pool named `+"`"+`example-pool`+"`"+` in a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster node-pool delete example-cluster example-pool`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`force`, `Deletes node pool without a confirmation prompt`, ox.Short("f")).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool delete-node
						ox.Usage(`delete-node`, `Delete a node`),
						ox.Banner(`
Deletes a node in the specified node pool. By default, this deletion happens gracefully and Kubernetes drains the node of any pods before deleting it.`),
						ox.Spec(`<cluster-id|cluster-name> <pool-id|pool-name> <node-id> [flags]`),
						ox.Example(`
The following example deletes a node named `+"`"+`example-node`+"`"+` in a node pool named `+"`"+`example-pool`+"`"+`: doctl kubernetes cluster node-pool delete-node example-cluster example-pool example-node`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`force`, `Deletes the node without a confirmation prompt`, ox.Short("f")).
							Bool(`skip-drain`, `Skips draining the node before deletion`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool get
						ox.Usage(`get`, `Retrieve information about a cluster's node pool`),
						ox.Banner(`
Retrieves information about the specified node pool in the specified cluster, including:

- The node pool ID
- The slug indicating the machine size of the nodes, such as `+"`"+`s-1vcpu-2gb`+"`"+`
- The number of nodes in the pool
- The tags applied to the node pool
- The names of the nodes

Specifying `+"`"+`--output=json`+"`"+` when calling this command returns additional information about the individual nodes in the response, such as their IDs, status, creation time, and update time.`),
						ox.Spec(`<cluster-id|cluster-name> <pool-id|pool-name> [flags]`),
						ox.Aliases("g"),
						ox.Example(`
The following example retrieves information about a node pool named `+"`"+`example-pool`+"`"+` in a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster node-pool get example-cluster example-pool`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Count`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Labels`+"`"+`, `+"`"+`Taints`+"`"+`, `+"`"+`Nodes`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool list
						ox.Usage(`list`, `List a cluster's node pools`),
						ox.Banner(`
Retrieves information about the specified cluster's node pools, including:

- The node pool ID
- The slug indicating the machine size of the nodes, such as `+"`"+`s-1vcpu-2gb`+"`"+`
- The number of nodes in the pool
- The tags applied to the node pool
- The names of the nodes

Specifying `+"`"+`--output=json`+"`"+` when calling this command returns additional information about the individual nodes in the response, such as their IDs, status, creation time, and update time.`),
						ox.Spec(`<cluster-id|cluster-name> [flags]`),
						ox.Aliases("ls"),
						ox.Example(`
The following example retrieves information about all node pools in a cluster named `+"`"+`example-cluster`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to only return the ID, name, and nodes for each pool: doctl kubernetes cluster node-pool list example-cluster --format ID,Name,Nodes`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Size`+"`"+`, `+"`"+`Count`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Labels`+"`"+`, `+"`"+`Taints`+"`"+`, `+"`"+`Nodes`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool replace-node
						ox.Usage(`replace-node`, `Replace node with a new one`),
						ox.Banner(`
Deletes the specified node in the specified node pool, and then creates a new node in its place. This is useful if you suspect a node has entered an undesired state. By default, the deletion happens gracefully and Kubernetes drains the node of any pods before deleting it.`),
						ox.Spec(`<cluster-id|cluster-name> <pool-id|pool-name> <node-id> [flags]`),
						ox.Example(`
The following example replaces a node named `+"`"+`example-node`+"`"+` in a node pool named `+"`"+`example-pool`+"`"+`: doctl kubernetes cluster node-pool replace-node example-cluster example-pool example-node`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`force`, `Replaces node without confirmation prompt`, ox.Short("f")).
							Bool(`skip-drain`, `Skips draining the node before replacement`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster node-pool update
						ox.Usage(`update`, `Update an existing node pool in a cluster`),
						ox.Banner(`Updates a node pool in a cluster. You can update any value for which there is a flag.`),
						ox.Spec(`<cluster-id|cluster-name> <pool-id|pool-name> [flags]`),
						ox.Aliases("u"),
						ox.Example(`
The following example updates a node pool named `+"`"+`example-pool`+"`"+` in a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster node-pool update example-cluster example-pool --count 5 --taint "key1=value1:NoSchedule" --taint "key2:NoExecute"`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							Bool(`auto-scale`, `Enables auto-scaling on the node pool`).
							Int(`count`, `The number of nodes in the node pool`).
							Map(`label`, `A label in key=value notation to apply to the node pool. Repeat this flag to specify additional labels. Existing labels are removed from the node pool if they are not specified in the updated value.`, ox.Spec(`key=value`)).
							Int(`max-nodes`, `The maximum number of nodes in the node pool when autoscaling is enabled`).
							Int(`min-nodes`, `The minimum number of nodes in the node pool when autoscaling is enabled`).
							String(`name`, `The name of the node pool`).
							Slice(`tag`, `A tag to apply to the node pool. Repeat this flag to specify additional tags. An existing tag is removed from the node pool if it is not specified by any flag.`).
							Map(`taint`, `Taint in key=value:effect notation to apply to the node pool. Repeat this flag to specify additional taints. Set to an empty string ("") to clear all taints. An existing taint is removed from the node pool if it is not specified by any flag.`, ox.Spec(`key=value:effect`)).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl kubernetes cluster registry
					ox.Usage(`registry`, `Display commands for integrating clusters with docr`),
					ox.Banner(`The commands under `+"`"+`registry`+"`"+` are for managing DOCR integration with Kubernetes clusters. You can use these commands to add or remove registry from one or more clusters.`),
					ox.Spec(`[command]`),
					ox.Aliases("reg"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl kubernetes cluster registry [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl kubernetes cluster registry add
						ox.Usage(`add`, `Add container registry support to Kubernetes clusters`),
						ox.Banner(`
Adds container registry support to the specified Kubernetes cluster(s).`),
						ox.Spec(`<cluster-id|cluster-name> <cluster-id|cluster-name> [flags]`),
						ox.Aliases("a"),
						ox.Example(`
The following example adds container registry support to a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster registry add example-cluster`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl kubernetes cluster registry remove
						ox.Usage(`remove`, `Remove container registry support from Kubernetes clusters`),
						ox.Banner(`
Removes container registry support from the specified Kubernetes cluster(s).`),
						ox.Spec(`<cluster-id|cluster-name> <cluster-id|cluster-name> [flags]`),
						ox.Aliases("rm"),
						ox.Example(`
The following example removes container registry support from a cluster named `+"`"+`example-cluster`+"`"+`: doctl kubernetes cluster registry remove example-cluster`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl kubernetes cluster update
					ox.Usage(`update`, `Update a Kubernetes cluster's configuration`),
					ox.Banner(`
Updates the configuration values for a Kubernetes cluster. The cluster must be referred to by its name or ID. Use the `+"`"+`doctl kubernetes cluster list`+"`"+` command to get a list of clusters on your account.`),
					ox.Spec(`<id|name> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates a cluster named `+"`"+`example-cluster`+"`"+` to enable automatic upgrades and sets the maintenance window to `+"`"+`saturday=02:00`+"`"+`: doctl kubernetes cluster update example-cluster --auto-upgrade --maintenance-window saturday=02:00`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`auto-upgrade`, `Indicates whether the cluster automatically upgrades to new patch releases during its maintenance window. To enable automatic upgrade, use --auto-upgrade=true.`, ox.Spec(`--auto-upgrade=true`)).
						String(`cluster-name`, `Specifies a new cluster name`).
						Slice(`control-plane-firewall-allowed-addresses`, `A comma-separated list of allowed addresses that can access the control plane.`).
						Bool(`enable-control-plane-firewall`, `Creates the cluster with control plane firewall enabled. Defaults to false. To enable the control plane firewall, supply --enable-control-plane-firewall=true.`).
						Bool(`ha`, `Enables the highly-available control plane for the cluster`).
						String(`maintenance-window`, `Sets the beginning of the four hour maintenance window for the cluster. Syntax is in the format: 'day=HH:MM', where time is in UTC. Day can be: any, `+"`"+`monday`+"`"+`, `+"`"+`tuesday`+"`"+`, `+"`"+`wednesday`+"`"+`, `+"`"+`thursday`+"`"+`, `+"`"+`friday`+"`"+`, `+"`"+`saturday`+"`"+`, `+"`"+`sunday`+"`"+`.`, ox.Default("any=00:00")).
						Bool(`set-current-context`, `Sets the current kubectl context to that of the new cluster`, ox.Default("true")).
						Bool(`surge-upgrade`, `Enables surge-upgrade for the cluster`).
						Slice(`tag`, `A comma-separated list of tags to apply to the cluster. This removes other  user-provided tags from the cluster if they are not specified in this argument.`).
						Bool(`update-kubeconfig`, `Updates the cluster in your kubeconfig`, ox.Default("true")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes cluster upgrade
					ox.Usage(`upgrade`, `Upgrades a cluster to a new Kubernetes version`),
					ox.Banner(`

Upgrades a Kubernetes cluster. By default, this upgrades the cluster to the latest available release, but you can also specify any version listed for your cluster by using `+"`"+`doctl k8s get-upgrades`+"`"+`.`),
					ox.Spec(`<id|name> [flags]`),
					ox.Example(`
The following example upgrades a cluster named `+"`"+`example-cluster`+"`"+` to version 1.28.2: doctl kubernetes cluster upgrade example-cluster --version 1.28.2-do.0`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl kubernetes options
				ox.Usage(`options`, `List possible option values for use inside Kubernetes commands`),
				ox.Banner(`The `+"`"+`options`+"`"+` commands are used to enumerate values for use with `+"`"+`doctl`+"`"+`'s Kubernetes commands. This is useful in certain cases where flags only accept input that is from a list of possible values, such as Kubernetes versions, datacenter regions, and machine sizes.`),
				ox.Spec(`[command]`),
				ox.Aliases("opts", "o"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl kubernetes options [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl kubernetes options regions
					ox.Usage(`regions`, `Lists regions that support DigitalOcean Kubernetes clusters`),
					ox.Banner(`Lists regions that support DigitalOcean Kubernetes clusters`),
					ox.Spec(`[flags]`),
					ox.Aliases("r"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes options sizes
					ox.Usage(`sizes`, `Lists machine sizes that you can use in a DigitalOcean Kubernetes cluster`),
					ox.Banner(`Lists machine sizes that you can use in a DigitalOcean Kubernetes cluster`),
					ox.Spec(`[flags]`),
					ox.Aliases("s"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl kubernetes options versions
					ox.Usage(`versions`, `Lists Kubernetes versions that you can use with DigitalOcean clusters`),
					ox.Banner(`Lists Kubernetes versions that you can use with DigitalOcean clusters`),
					ox.Spec(`[flags]`),
					ox.Aliases("v"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
		),
		ox.Sub( // doctl monitoring
			ox.Usage(`monitoring`, `[Beta] Display commands to manage monitoring`),
			ox.Banner(`The sub-commands of `+"`"+`doctl monitoring`+"`"+` manage the monitoring on your account.

You can create alert policies to monitor the resource consumption of your Droplets, and uptime checks to monitor the availability of your websites and services`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl monitoring [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl monitoring alert
				ox.Usage(`alert`, `Display commands for managing alert policies`),
				ox.Banner(`The commands under `+"`"+`doctl monitoring alert`+"`"+` are for managing alert policies.

You can apply alert policies to resources in order to receive alerts on resource consumption. 
			
If you'd like to alert on the uptime of a specific URL or IP address, use `+"`"+`doctl monitoring uptime alert`+"`"+` instead`),
				ox.Spec(`[command]`),
				ox.Aliases("alerts", "a"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl monitoring alert [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl monitoring alert create
					ox.Usage(`create`, `Create an alert policy`),
					ox.Banner(`Creates a new alert policy. You can create policies that monitor various metrics of your Droplets and send you alerts when a metric exceeds a specified threshold.
	
	For example, you can create a policy that monitors a Droplet's CPU usage and triggers an alert when the Droplet's CPU usage exceeds more than 80% for more than 10 minutes.
	
	For a full list of policy types you can set up, see our API documentation: https://docs.digitalocean.com/reference/api/api-reference/#operation/monitoring_create_alertPolicy`),
					ox.Spec(`[flags]`),
					ox.Example(`
The following example creates an alert policy that sends an email to `+"`"+`admin@example.com`+"`"+` whenever the memory usage on the listed Droplets (entities) exceeds 80% for more than five minutes: doctl monitoring alert create --type "v1/insights/droplet/memory_utilization_percent" --compare GreaterThan --value 80 --window 5m --entities 386734086,191669331 --emails admin@example.com`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`compare`, `The comparator of the alert policy. Possible values: GreaterThan or `+"`"+`LessThan`+"`"+``, ox.Default("GreaterThan")).
						String(`description`, `A description of the alert policy`).
						Slice(`emails`, `Email address to send alerts to`).
						Bool(`enabled`, `Enables the alert policy`, ox.Default("true")).
						Slice(`entities`, `Resources to apply the alert against, such as a Droplet ID.`).
						String(`slack-channels`, `A Slack channel to send alerts to. For example, production-alerts`, ox.Default("production-alerts")).
						String(`slack-urls`, `A Slack webhook URL to send alerts to, for example, https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ.`, ox.Spec(`https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ`)).
						Slice(`tags`, `Tags to apply the alert policy to. If set to a tag, all Droplet with that tag are monitored by the alert policy.`).
						String(`type`, `The type of alert policy. For example,v1/insights/droplet/memory_utilization_percent alerts on the percent of memory utilization. For a full list of alert types, see https://docs.digitalocean.com/reference/api/api-reference/#operation/monitoring_create_alertPolicy`, ox.Spec(`v1/insights/droplet/memory_utilization_percent`)).
						String(`value`, `The threshold value of the alert policy to compare against. For example, if the alert policy is of type DropletCPUUtilizationPercent and the value is set to `+"`"+`80`+"`"+`, an alert is triggered if the Droplet's CPU usage exceeds 80% for the specified window.`, ox.Default("DropletCPUUtilizationPercent")).
						Duration(`window`, `The amount of time the resource must exceed the threshold value before an alert is triggered. Possible values: 5m, `+"`"+`10m`+"`"+`, `+"`"+`30m`+"`"+`, or `+"`"+`1h`+"`"+``, ox.Default("5m")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring alert delete
					ox.Usage(`delete`, `Delete an alert policy`),
					ox.Banner(`Deletes an alert policy.`),
					ox.Spec(`<alert-policy-uuid>... [flags]`),
					ox.Aliases("rm"),
					ox.Example(`
The following example deletes an alert policy with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring alert delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete an alert policy without a confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring alert get
					ox.Usage(`get`, `Retrieve information about an alert policy`),
					ox.Banner(`Retrieves an alert policy and its configuration.`),
					ox.Spec(`<alert-policy-uuid> [flags]`),
					ox.Example(`
The following example retrieves information about a policy with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring alert get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						UUID(`format`, `Columns for output in a comma-separated list. Possible values: UUID, `+"`"+`Type`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Compare`+"`"+`, `+"`"+`Value`+"`"+`, `+"`"+`Window`+"`"+`, `+"`"+`Entities`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`, `+"`"+`Enabled`+"`"+`.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring alert list
					ox.Usage(`list`, `List all alert policies`),
					ox.Banner(`Retrieves a list of all the alert policies in your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all alert policies in your account: doctl monitoring alert list`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						UUID(`format`, `Columns for output in a comma-separated list. Possible values: UUID, `+"`"+`Type`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Compare`+"`"+`, `+"`"+`Value`+"`"+`, `+"`"+`Window`+"`"+`, `+"`"+`Entities`+"`"+`, `+"`"+`Tags`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`, `+"`"+`Enabled`+"`"+`.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring alert update
					ox.Usage(`update`, `Update an alert policy`),
					ox.Banner(`Updates an existing alert policy.`),
					ox.Spec(`<alert-policy-uuid>... [flags]`),
					ox.Example(`
The following example updates an alert policy's details: doctl monitoring alert update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --type "v1/insights/droplet/memory_utilization_percent" --compare GreaterThan --value 80 --window 10m --entities 386734086,191669331 --emails admin@example.com`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`compare`, `The comparator of the alert policy. Either GreaterThan or `+"`"+`LessThan`+"`"+``, ox.Default("GreaterThan")).
						String(`description`, `A description of the alert policy.`).
						Slice(`emails`, `Email addresses to send alerts to`).
						Bool(`enabled`, `Whether the alert policy is enabled.`, ox.Default("true")).
						Slice(`entities`, `Resources to apply the policy to`).
						String(`slack-channels`, `A Slack channel to send alerts to, for example, production-alerts. Must be used with `+"`"+`--slack-url`+"`"+`.`, ox.Default("production-alerts")).
						String(`slack-urls`, `A Slack webhook URL to send alerts to, for example, https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ.`, ox.Spec(`https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ`)).
						Slice(`tags`, `Tags to apply the alert against`).
						String(`type`, `The type of alert policy. For example,v1/insights/droplet/memory_utilization_percent alerts on the percent of memory utilization. For a full list of alert types, see https://docs.digitalocean.com/reference/api/api-reference/#operation/monitoring_create_alertPolicy`, ox.Spec(`v1/insights/droplet/memory_utilization_percent`)).
						Int(`value`, `The value of the alert policy to compare against.`).
						String(`window`, `The window to apply the alert policy conditions against.`, ox.Default("5m")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl monitoring uptime
				ox.Usage(`uptime`, `Display commands to manage uptime checks`),
				ox.Banner(`The sub-commands of `+"`"+`doctl uptime`+"`"+` manage your uptime checks.

DigitalOcean Uptime Checks provide the ability to monitor your endpoints from around the world,
and alert you when they're slow, unavailable, or SSL certificates are expiring.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl monitoring uptime [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl monitoring uptime alert
					ox.Usage(`alert`, `Display commands to manage uptime alerts`),
					ox.Banner(`The sub-commands of `+"`"+`doctl monitoring uptime alert`+"`"+` manage your uptime alerts.

DigitalOcean Uptime Alerts provide the ability to monitor your endpoints from around the world,
and alert you when they're slow, unavailable, or SSL certificates are expiring.

In order to set up uptime alerts, you must first set up an uptime check. Uptime checks monitor and track the status of an endpoint while alerts notify you of status changes based on the thresholds you set.`),
					ox.Spec(`[command]`),
					ox.Aliases("alerts"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Footer(`Use "doctl monitoring uptime alert [command] --help" for more information about a command.`),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					ox.Sub( // doctl monitoring uptime alert create
						ox.Usage(`create`, `Create an uptime alert`),
						ox.Banner(`Creates an alert policy for an uptime check.
	
	You can create an alert policy based on the following metrics: 
	
	- `+"`"+`latency`+"`"+`: Alerts on the response latency. `+"`"+`--threshold`+"`"+` value is an integer representing milliseconds.
	- `+"`"+`down`+"`"+`: Alerts on whether the endpoints registers as down from any of the configured regions. No `+"`"+`--threshold`+"`"+` value is necessary.
	- `+"`"+`down_global`+"`"+`: Alerts on a target registering as down globally. No `+"`"+`--threshold`+"`"+` value is necessary.
	- `+"`"+`ssl_expiry`+"`"+`: Alerts on an SSL certificate expiring within the set threshold of days. `+"`"+`--threshold`+"`"+` value is an integer representing days.`),
						ox.Spec(`<uptime-check-id> [flags]`),
						ox.Aliases("c"),
						ox.Example(`
The following example creates an alert for an uptime check with an ID of f81d4fae-7dec-11d0-a765-00a0c91e6bf6. The alert triggers if the endpoint's latency exceed 500ms for more than two minutes: doctl monitoring uptime alert create f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name "Example Alert" --type latency --threshold 100 --comparison greater_than --period 2m --emails "admin@example.com"`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`comparison`, `A comparison operator used against the alert's threshold. Possible values: greater_than or `+"`"+`less_than`+"`"+``, ox.Default("greater_than")).
							Slice(`emails`, `Emails addresses to send alerts. The addresses must be associated with your DigitalOcean account`).
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Threshold`+"`"+`, `+"`"+`Comparison`+"`"+`, `+"`"+`Period`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`.`, ox.Spec(`ID`)).
							String(`name`, `The name of the alert (required)`).
							Bool(`no-header`, `Return raw data with no headers`).
							Duration(`period`, `The period of time the threshold must be exceeded to trigger an alert. Possible values: 2m, `+"`"+`3m`+"`"+`, `+"`"+`5m`+"`"+`, `+"`"+`10m`+"`"+`, `+"`"+`15m`+"`"+`, `+"`"+`30m`+"`"+`, `+"`"+`1h`+"`"+` (required)`, ox.Default("2m")).
							String(`slack-channels`, `Slack channels to send alerts to, for example, production-alerts. Must be used with the `+"`"+`--slack-urls`+"`"+` flag.`, ox.Default("production-alerts")).
							String(`slack-urls`, `A Slack webhook URL to send alerts to, for example, https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ.`, ox.Spec(`https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ`)).
							Int(`threshold`, `The threshold at which to trigger the alert. The specific threshold is dependent on the alert type.`).
							String(`type`, `The metric to alert on. Possible values: latency, `+"`"+`down`+"`"+`, `+"`"+`down_global`+"`"+`, `+"`"+`ssl_expiry`+"`"+` (required)`, ox.Default("latency")).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl monitoring uptime alert delete
						ox.Usage(`delete`, `Delete an uptime alert`),
						ox.Banner(`Deletes an uptime check on your account by ID.`),
						ox.Spec(`<uptime-check-id> <uptime-alert-id> [flags]`),
						ox.Aliases("d", "del", "rm"),
						ox.Example(`
The following example deletes an alert policy with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+` that is a policy of an uptime check with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring uptime alert delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6 418b7972-fc67-41ea-ab4b-6f9477c4f7d8`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl monitoring uptime alert get
						ox.Usage(`get`, `Get uptime alert`),
						ox.Banner(`Retrieves information about an uptime alert policy.`),
						ox.Spec(`<uptime-check-id> <uptime-alert-id> [flags]`),
						ox.Aliases("g"),
						ox.Example(`
The following example retrieves the configuration for an alert policy with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+` that is a policy of an uptime check with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring uptime alert get f81d4fae-7dec-11d0-a765-00a0c91e6bf6 418b7972-fc67-41ea-ab4b-6f9477c4f7d8`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Threshold`+"`"+`, `+"`"+`Comparison`+"`"+`, `+"`"+`Period`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl monitoring uptime alert list
						ox.Usage(`list`, `List uptime alerts`),
						ox.Banner(`Retrieves a list of the alert policies for an uptime check.`),
						ox.Spec(`<uptime-check-id> [flags]`),
						ox.Aliases("ls"),
						ox.Example(`
The following example lists the alert policies for an uptime check with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring uptime alert list f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Threshold`+"`"+`, `+"`"+`Comparison`+"`"+`, `+"`"+`Period`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`.`, ox.Spec(`ID`)).
							Bool(`no-header`, `Return raw data with no headers`).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
					ox.Sub( // doctl monitoring uptime alert update
						ox.Usage(`update`, `Update an uptime alert`),
						ox.Banner(`Updates an uptime alert configuration.`),
						ox.Spec(`<uptime-check-id> <uptime-alert-id> [flags]`),
						ox.Aliases("u"),
						ox.Example(`
The following example updates an alert policy with the ID `+"`"+`418b7972-fc67-41ea-ab4b-6f9477c4f7d8`+"`"+` that is a policy of an uptime check with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring uptime alert update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 418b7972-fc67-41ea-ab4b-6f9477c4f7d8 --name "Example Alert" --type latency --threshold 100 --comparison greater_than --period 2m --emails "admin@example.com"`),
						ox.Help(ox.Sections(
							"Global Flags",
						)),
						ox.Flags().
							String(`comparison`, `A comparison operator used against the alert's threshold. Possible values: greater_than or `+"`"+`less_than`+"`"+``, ox.Default("greater_than")).
							Slice(`emails`, `Emails addresses to send alerts. The addresses must be associated with your DigitalOcean account`).
							String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Threshold`+"`"+`, `+"`"+`Comparison`+"`"+`, `+"`"+`Period`+"`"+`, `+"`"+`Emails`+"`"+`, `+"`"+`Slack Channels`+"`"+`.`, ox.Spec(`ID`)).
							String(`name`, `The name of the alert (required)`).
							Bool(`no-header`, `Return raw data with no headers`).
							Duration(`period`, `The period of time the threshold must be exceeded to trigger an alert. Possible values: 2m, `+"`"+`3m`+"`"+`, `+"`"+`5m`+"`"+`, `+"`"+`10m`+"`"+`, `+"`"+`15m`+"`"+`, `+"`"+`30m`+"`"+`, `+"`"+`1h`+"`"+` (required)`, ox.Default("2m")).
							String(`slack-channels`, `Slack channels to send uptime alerts to, for example, production-alerts. Must be used with the `+"`"+`--slack-urls`+"`"+` flag.`, ox.Default("production-alerts")).
							String(`slack-urls`, `A Slack webhook URL to send alerts to, for example, https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ.`, ox.Spec(`https://hooks.slack.com/services/T1234567/AAAAAAAA/ZZZZZZ`)).
							Int(`threshold`, `The threshold at which to trigger the alert. The specific threshold is dependent on the alert type.`).
							String(`type`, `The metric to alert on. Possible values: latency, `+"`"+`down`+"`"+`, `+"`"+`down_global`+"`"+`, `+"`"+`ssl_expiry`+"`"+` (required)`, ox.Default("latency")).
							String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
							String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
							String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
							String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
							Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
							Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
							String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
							Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
							Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
					),
				),
				ox.Sub( // doctl monitoring uptime create
					ox.Usage(`create`, `Create an uptime check`),
					ox.Banner(`Creates an uptime check on your account. Uptime checks monitor any endpoint that is accessible over HTTP, HTTPS, ping (ICMP).
	
	You can use this check to set up an alert policy using the `+"`"+`doctl monitoring uptime alert`+"`"+` commands.`),
					ox.Spec(`<uptime-check-name> [flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates an uptime check that monitors the URL, `+"`"+`example.com`+"`"+` from the eastern and western regions of the Unites States: doctl monitoring uptime create --target https://example.com --type https --regions us_east,us_west --enabled true`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`enabled`, `Whether or not the uptime check is enabled. Defaults to true`, ox.Default("true")).
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Target`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`Enabled`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`regions`, `A comma-separated list of regions to monitor the target from. Possible values: us_east, `+"`"+`us_west`+"`"+`, `+"`"+`eu_west`+"`"+`, `+"`"+`se_asia`+"`"+`. Defaults to `+"`"+`us_east`+"`"+``, ox.Default("[us_east]")).
						String(`target`, `A valid URL to monitor (required)`).
						String(`type`, `The protocol to use to monitor the target URL. Possible values: ping, `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`. Defaults to either `+"`"+`http`+"`"+` or `+"`"+`https`+"`"+`, depending on the URL target provided`, ox.Default("ping")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring uptime delete
					ox.Usage(`delete`, `Delete an uptime check`),
					ox.Banner(`Deletes an uptime check on your account.`),
					ox.Spec(`<uptime-check-id> [flags]`),
					ox.Aliases("d", "del", "rm"),
					ox.Example(`
The following example deletes an uptime check with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl monitoring uptime delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring uptime get
					ox.Usage(`get`, `Get an uptime check`),
					ox.Banner(`Retrieves information about an uptime check on your account.`),
					ox.Spec(`<uptime-check-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves the ID, name, and target of an uptime check: doctl monitoring uptime get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Target`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`Enabled`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring uptime list
					ox.Usage(`list`, `List uptime checks`),
					ox.Banner(`Retrieves a list of all of the uptime checks on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example retrieves a list of all of the uptime checks on your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, name, and target of each check: doctl monitoring uptime list --format ID,Name,Target`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Target`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`Enabled`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl monitoring uptime update
					ox.Usage(`update`, `Update an uptime check`),
					ox.Banner(`Updates an uptime check on your account.

All of these flags are required. Uptime checks cannot be disabled via `+"`"+`doctl`+"`"+`. You can only disable them using the control panel or the public API.`),
					ox.Spec(`<uptime-check-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates the name, target, type, and regions of an uptime check: doctl monitoring uptime update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name example --target https://example.com --type https --regions us_east,us_west`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`Type`+"`"+`, `+"`"+`Target`+"`"+`, `+"`"+`Regions`+"`"+`, `+"`"+`Enabled`+"`"+`.`, ox.Spec(`ID`)).
						String(`name`, `A name for the check (required)`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`regions`, `A comma-separated list of regions to monitor the target from. Possible values: us_east, `+"`"+`us_west`+"`"+`, `+"`"+`eu_west`+"`"+`, `+"`"+`se_asia`+"`"+`. Defaults to `+"`"+`us_east`+"`"+` (required)`, ox.Default("[us_east]")).
						String(`target`, `A valid URL to monitor (required)`).
						String(`type`, `The protocol to use to monitor the target URL. Possible values: ping, `+"`"+`http`+"`"+`, `+"`"+`https`+"`"+`. Defaults to either `+"`"+`http`+"`"+` or `+"`"+`https`+"`"+`, depending on the URL target provided (required)`, ox.Default("ping")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
		),
		ox.Sub( // doctl projects
			ox.Usage(`projects`, `Manage projects and assign resources to them`),
			ox.Banner(`The subcommands of `+"`"+`doctl projects`+"`"+` allow you to create, manage, and assign resources to your projects.

Projects allow you to organize your DigitalOcean resources (like Droplets, Spaces, load balancers, domains, and floating IPs) into groups that fit the way you work. You can create projects that align with the applications, environments, and clients that you host on DigitalOcean.`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl projects [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl projects create
				ox.Usage(`create`, `Create a new project`),
				ox.Banner(`Creates a new project in your account.

Projects allow you to organize your DigitalOcean resources (like Droplets, Spaces, load balancers, domains, and floating IPs) into groups that fit the way you work. You can create projects that align with the applications, environments, and clients that you host on DigitalOcean.`),
				ox.Spec(`[flags]`),
				ox.Aliases("c"),
				ox.Example(`
The following example creates a project named `+"`"+`Example Project`+"`"+` with the purpose "Frontend development": doctl projects create --name "Example Project" --purpose "Frontend development"`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`description`, `A description of the project`).
					String(`environment`, `The environment in which your project resides. Possible enum values: `+"`"+`Development`+"`"+`, `+"`"+`Staging`+"`"+`, `+"`"+`Production`+"`"+`.`, ox.Spec(`enum`)).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`OwnerUUID`+"`"+`, `+"`"+`OwnerID`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Purpose`+"`"+`, `+"`"+`Environment`+"`"+`, `+"`"+`IsDefault`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`.`, ox.Spec(`ID`)).
					String(`name`, `The project's name (required)`).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`purpose`, `The project's purpose (required)`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl projects delete
				ox.Usage(`delete`, `Delete the specified project`),
				ox.Banner(`Deletes a project. To be deleted, a project must not have any resources assigned to it.`),
				ox.Spec(`<project-id> [<project-id> ...] [flags]`),
				ox.Aliases("d", "rm"),
				ox.Example(`
The following example deletes the project with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl projects delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `Deletes the project without confirmation`, ox.Short("f")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl projects get
				ox.Usage(`get`, `Retrieve details for a specific project`),
				ox.Banner(`Retrieves the following details for an existing project (use `+"`"+`default`+"`"+` as the <project-id> to retrieve details about your default project):

- The project's ID, in UUID format
- The project owner's account UUID
- The name of the project
- The project's description
- The project's specified purpose
- The project's environment (Development, Staging, or Production)
- A boolean indicating whether it is you default project
- The date and time when the project was created
- The date and time when the project was last updated`),
				ox.Spec(`<project-id> [flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves details for a project with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl projects get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`OwnerUUID`+"`"+`, `+"`"+`OwnerID`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Purpose`+"`"+`, `+"`"+`Environment`+"`"+`, `+"`"+`IsDefault`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl projects list
				ox.Usage(`list`, `List existing projects`),
				ox.Banner(`List details for your DigitalOcean projects, including:

- The project's ID, in UUID format
- The project owner's account UUID
- The name of the project
- The project's description
- The project's specified purpose
- The project's environment (Development, Staging, or Production)
- A boolean indicating whether it is you default project
- The date and time when the project was created
- The date and time when the project was last updated`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example retrieves a list of projects on your account and uses the `+"`"+`--format`+"`"+` flag to return only the ID, name, and purpose of each project: doctl projects list --format ID,Name,Purpose`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`OwnerUUID`+"`"+`, `+"`"+`OwnerID`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Purpose`+"`"+`, `+"`"+`Environment`+"`"+`, `+"`"+`IsDefault`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl projects resources
				ox.Usage(`resources`, `Manage resources assigned to a project`),
				ox.Banner(`The subcommands of `+"`"+`doctl projects resources`+"`"+` allow you to list and assign resources to your projects.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl projects resources [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl projects resources assign
					ox.Usage(`assign`, `Assign one or more resources to a project`),
					ox.Banner(`Assign one or more resources to a project by specifying the resource's uniform resource name ("URN").

A valid URN has the format: `+"`"+`do:resource_type:resource_id`+"`"+`. For example:

  - `+"`"+`do:droplet:4126873`+"`"+`
  - `+"`"+`do:volume:6fc4c277-ea5c-448a-93cd-dd496cfef71f`+"`"+`
  - `+"`"+`do:app:be5aab85-851b-4cab-b2ed-98d5a63ba4e8`+"`"+``),
					ox.Spec(`<project-id> --resource=<urn> [--resource=<urn> ...] [flags]`),
					ox.Aliases("a"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Slice(`resource`, `URNs specifying resources to assign to the project`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl projects resources get
					ox.Usage(`get`, `Retrieve a resource by its URN`),
					ox.Banner(`Retrieve information about a resource by specifying its uniform resource name ("URN"). Currently, only Droplets, floating IPs, load balancers, domains, volumes, and App Platform apps are supported.

A valid URN has the format: `+"`"+`do:resource_type:resource_id`+"`"+`. For example:

  - `+"`"+`do:droplet:4126873`+"`"+`
  - `+"`"+`do:volume:6fc4c277-ea5c-448a-93cd-dd496cfef71f`+"`"+`
  - `+"`"+`do:app:be5aab85-851b-4cab-b2ed-98d5a63ba4e8`+"`"+``),
					ox.Spec(`<urn> [flags]`),
					ox.Aliases("g"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: URN, `+"`"+`AssignedAt`+"`"+`, `+"`"+`Status`+"`"+`.`, ox.Default("URN")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl projects resources list
					ox.Usage(`list`, `List resources assigned to a project`),
					ox.Banner(`List all of the resources assigned to the specified project displaying their uniform resource names ("URNs").`),
					ox.Spec(`<project-id> [flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: URN, `+"`"+`AssignedAt`+"`"+`, `+"`"+`Status`+"`"+`.`, ox.Default("URN")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl projects update
				ox.Usage(`update`, `Update an existing project`),
				ox.Banner(`Updates information about an existing project. Use `+"`"+`default`+"`"+` as the <project-id> to update your default project.`),
				ox.Spec(`<project-id> [flags]`),
				ox.Aliases("u"),
				ox.Example(`
The following example updates the project with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to have the name `+"`"+`API Project`+"`"+` and the purpose "Backend development": doctl projects update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name "API Project" --purpose "Backend development"`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`description`, `A description of the project`).
					String(`environment`, `The environment in which your project resides. Possible enum values: `+"`"+`Development`+"`"+`, `+"`"+`Staging`+"`"+`, or `+"`"+`Production`+"`"+``, ox.Spec(`enum`)).
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`OwnerUUID`+"`"+`, `+"`"+`OwnerID`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Purpose`+"`"+`, `+"`"+`Environment`+"`"+`, `+"`"+`IsDefault`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`is_default`, `Sets the specified project as your default project`).
					String(`name`, `The project's name`).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`purpose`, `The project's purpose`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl registry
			ox.Usage(`registry`, `Display commands for working with container registries`),
			ox.Banner(`The subcommands of `+"`"+`doctl registry`+"`"+` create, manage, and allow access to your private container registry.`),
			ox.Spec(`[command]`),
			ox.Aliases("reg", "r"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl registry [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl registry create
				ox.Usage(`create`, `Create a private container registry`),
				ox.Banner(`Creates a new private container registry with the provided name.`),
				ox.Spec(`<registry-name> [flags]`),
				ox.Example(`
The following example creates a registry named `+"`"+`example-registry`+"`"+` in the NYC3 region: doctl registry create example-registry --region=nyc3`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`region`, `A slug indicating which datacenter region the registry reside in. For a list of supported region slugs, use the `+"`"+`doctl registry options available-regions`+"`"+` command`, ox.Default("slug")).
					String(`subscription-tier`, `registry options subscription-tiers   Subscription tier for the new registry. For a list of possible values, use the doctl registry options subscription-tiers command. (required)`, ox.Spec(`doctl`), ox.Default("basic")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry delete
				ox.Usage(`delete`, `Delete a container registry`),
				ox.Banner(`Permanently deletes a private container registry and all of its contents. This is irreversible.`),
				ox.Spec(`[flags]`),
				ox.Aliases("d", "del", "rm"),
				ox.Example(`
The following example deletes a registry named `+"`"+`example-registry`+"`"+`: doctl registry delete example-registry`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `Force deletes the registry without confirmation prompt`, ox.Short("f")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry docker-config
				ox.Usage(`docker-config`, `Generate a Docker auth configuration for a registry`),
				ox.Banner(`Outputs a JSON-formatted Docker configuration that you can use to configure a Docker client to authenticate with your private container registry. This configuration is useful for configuring third-party tools that need access to your registry. For configuring your local Docker client use `+"`"+`doctl registry login`+"`"+` instead, as it preserves the configuration of any other registries you have authenticated to.

By default this command generates read-only credentials. Use the `+"`"+`--read-write`+"`"+` flag to generate credentials that can push. The configuration produced by this command contains a DigitalOcean API token that can be used to access your account and should be treated as sensitive information.`),
				ox.Spec(`[flags]`),
				ox.Aliases("config"),
				ox.Example(`
The following example generates a Docker configuration for a registry named `+"`"+`example-registry`+"`"+` and uses the `+"`"+`--expiry-seconds`+"`"+` to set the credentials to expire after one day: doctl registry docker-config example-registry --expiry-seconds=86400`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`expiry-seconds`, `The length of time the registry credentials are valid for, in seconds. By default, the credentials do not expire.`).
					Bool(`read-write`, `Generates credentials that can push to your registry`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry garbage-collection
				ox.Usage(`garbage-collection`, `Display commands for garbage collection for a container registry`),
				ox.Banner(`The subcommands of `+"`"+`doctl registry garbage-collection`+"`"+` start a garbage collection, retrieve or cancel a currently-active garbage collection, or list past garbage collections for a specified registry.`),
				ox.Spec(`[command]`),
				ox.Aliases("gc", "g"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl registry garbage-collection [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl registry garbage-collection cancel
					ox.Usage(`cancel`, `Cancel the currently-active garbage collection for a container registry`),
					ox.Banner(`Cancels the currently-active garbage collection for a container registry.`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example cancels the currently-active garbage collection for a registry: doctl registry garbage-collection cancel`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry garbage-collection get-active
					ox.Usage(`get-active`, `Retrieve information about the currently-active garbage collection for a container registry`),
					ox.Banner(`Retrieves the currently-active garbage collection for a container registry, if any active garbage collection exists. Information included about the registry includes:
  - UUID
  - Status
  - Registry name
  - Create time
  - Updated at time
  - Blobs deleted
  - Freed bytes`),
					ox.Spec(`[flags]`),
					ox.Aliases("ga", "g"),
					ox.Example(`
The following example retrieves the currently-active garbage collection for a registry: doctl registry garbage-collection get-active`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						UUID(`format`, `Columns for output in a comma-separated list. Possible values: UUID, `+"`"+`RegistryName`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`, `+"`"+`BlobsDeleted`+"`"+`, `+"`"+`FreedBytes`+"`"+`.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry garbage-collection list
					ox.Usage(`list`, `Retrieve information about past garbage collections for a container registry`),
					ox.Banner(`Retrieves a list of past garbage collections for a registry. Information about each garbage collection includes:
  - UUID
  - Status
  - Registry name
  - Create time
  - Updated at time
  - Blobs deleted
  - Freed bytes`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls", "l"),
					ox.Example(`
The following example retrieves a list of past garbage collections for a registry: doctl registry garbage-collection list`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						UUID(`format`, `Columns for output in a comma-separated list. Possible values: UUID, `+"`"+`RegistryName`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`, `+"`"+`BlobsDeleted`+"`"+`, `+"`"+`FreedBytes`+"`"+`.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry garbage-collection start
					ox.Usage(`start`, `Start garbage collection for a container registry`),
					ox.Banner(`Starts a garbage collection on a container registry. You can only have one active garbage collection at a time for a given registry.`),
					ox.Spec(`[flags]`),
					ox.Aliases("s"),
					ox.Example(`
The following example starts a garbage collection on a registry named `+"`"+`example-registry`+"`"+`: doctl registry garbage-collection start example-registry`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`exclude-unreferenced-blobs`, `Exclude unreferenced blobs from garbage collection.`).
						Bool(`force`, `Run garbage collection without confirmation prompt`, ox.Short("f")).
						UUID(`format`, `Columns for output in a comma-separated list. Possible values: UUID, `+"`"+`RegistryName`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`CreatedAt`+"`"+`, `+"`"+`UpdatedAt`+"`"+`, `+"`"+`BlobsDeleted`+"`"+`, `+"`"+`FreedBytes`+"`"+`.`).
						Bool(`include-untagged-manifests`, `Include untagged manifests in garbage collection.`).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl registry get
				ox.Usage(`get`, `Retrieve details about a container registry`),
				ox.Banner(`Retrieves details about a private container registry, including its name and the endpoint used to access it.`),
				ox.Spec(`[flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves details about a registry named `+"`"+`example-registry`+"`"+`: doctl registry get example-registry`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`Endpoint`+"`"+`, `+"`"+`Region`+"`"+`.`, ox.Default("Name")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry kubernetes-manifest
				ox.Usage(`kubernetes-manifest`, `Generate a Kubernetes secret manifest for a registry.`),
				ox.Banner(`Outputs a YAML-formatted Kubernetes secret manifest that can be used to grant a Kubernetes cluster pull access to your private container registry.

By default, the secret manifest is applied to all the namespaces for the Kubernetes cluster using the DOSecret operator. The DOSecret operator is available on clusters running version 1.15.12-do.2 or greater. For older clusters, or to restrict the secret to a specific namespace, use the `+"`"+`--namespace`+"`"+` flag.

You can redirect the command's output to a file to save the manifest for later use or pipe it directly to `+"`"+`kubectl`+"`"+` to create the secret in your cluster:

    doctl registry kubernetes-manifest | kubectl apply -f -`),
				ox.Spec(`[flags]`),
				ox.Aliases("k8s"),
				ox.Example(`
The following example generates a secret manifest for a registry named `+"`"+`example-registry`+"`"+` and applies it to the `+"`"+`kube-system`+"`"+` namespace: doctl registry kubernetes-manifest example-registry --namespace=kube-system`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`name`, `The secret's name. Defaults to the registry name prefixed with "registry-"`).
					String(`namespace`, `The Kubernetes namespace to hold the secret`, ox.Default("kube-system")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry login
				ox.Usage(`login`, `Log in Docker to a container registry`),
				ox.Banner(`Logs Docker into Container Registry making pull and push commands to your private container registry authenticated.`),
				ox.Spec(`[flags]`),
				ox.Example(`
The following example logs Docker into a registry and provides Docker with read-only credentials: doctl registry login --read-only=true`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`expiry-seconds`, `The length of time the registry credentials are valid for, in seconds. By default, the credentials expire after 30 days.`).
					Bool(`never-expire`, `Sets the DigitalOcean API token generated by the login command to never expire. By default, this is set to false.`).
					Bool(`read-only`, `Sets the DigitalOcean API token generated by the login command to read-only, causing any push operations to fail. By default, the API token is read-write.`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry logout
				ox.Usage(`logout`, `Log out Docker from a container registry`),
				ox.Banner(`This command logs Docker out of the private container registry, revoking access to it.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`authorization-server-endpoint`, `The endpoint of the OAuth authorization server used to revoke credentials on logout.`, ox.Default("https://cloud.digitalocean.com/v1/oauth/revoke")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl registry options
				ox.Usage(`options`, `List available container registry options`),
				ox.Banner(`This command lists options available when creating or updating a container registry.`),
				ox.Spec(`[command]`),
				ox.Aliases("opts", "o"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl registry options [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl registry options available-regions
					ox.Usage(`available-regions`, `Lists available container registry regions`),
					ox.Banner(`Lists available container registry regions`),
					ox.Spec(`[flags]`),
					ox.Aliases("regions"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry options subscription-tiers
					ox.Usage(`subscription-tiers`, `Lists available container registry subscription tiers`),
					ox.Banner(`Lists available container registry subscription tiers`),
					ox.Spec(`[flags]`),
					ox.Aliases("tiers"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl registry repository
				ox.Usage(`repository`, `Display commands for working with repositories in a container registry`),
				ox.Banner(`The subcommands of `+"`"+`doctl registry repository`+"`"+` allow you to manage various properties of your repository.`),
				ox.Spec(`[command]`),
				ox.Aliases("repo", "r"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl registry repository [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl registry repository delete-manifest
					ox.Usage(`delete-manifest`, `Delete one or more container repository manifests by digest`),
					ox.Banner(`Permanently deletes one or more repository manifests by digest.`),
					ox.Spec(`<repository> <manifest-digest>... [flags]`),
					ox.Aliases("dm"),
					ox.Example(`
The following example deletes a manifest with digest `+"`"+`sha256:1234567890abcdef`+"`"+` from a repository named `+"`"+`example-repository`+"`"+` in a registry named `+"`"+`example-registry`+"`"+`: doctl registry repository delete-manifest example-registry/example-repository sha256:a67c20e45178d90cbe686575719bd81f279b06842dc77521690e292c1eea685`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Deletes manifest without confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry repository delete-tag
					ox.Usage(`delete-tag`, `Delete one or more container repository tags`),
					ox.Banner(`Permanently deletes one or more repository tags.`),
					ox.Spec(`<repository> <tag>... [flags]`),
					ox.Aliases("dt"),
					ox.Example(`
The following example deletes a tag named `+"`"+`web`+"`"+` from a repository named `+"`"+`example-repository`+"`"+` in a registry named `+"`"+`example-registry`+"`"+`: doctl registry repository delete-tag example-registry/example-repository web`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete tag without confirmation prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry repository list-manifests
					ox.Usage(`list-manifests`, `List manifests for a repository in a container registry`),
					ox.Banner(`Retrieves information about manifests in a repository, including:
  - The manifest digest
  - The compressed size
  - The uncompressed size
  - The last updated timestamp
  - The manifest tags
  - The manifest blobs (available in detailed output only)`),
					ox.Spec(`<repository> [flags]`),
					ox.Aliases("lm"),
					ox.Example(`
The following example lists manifests in a repository named `+"`"+`example-repository`+"`"+`. The command also uses the `+"`"+`--format`+"`"+` flag to return only the digest and update time for each manifest: doctl registry repository list-manifests example-repository --format Digest,UpdatedAt`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Digest, `+"`"+`CompressedSizeBytes`+"`"+`, `+"`"+`SizeBytes`+"`"+`, `+"`"+`UpdatedAt`+"`"+`, `+"`"+`Tags`+"`"+`.`, ox.Default("Digest")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry repository list-tags
					ox.Usage(`list-tags`, `List tags for a repository in a container registry`),
					ox.Banner(`Retrieves information about tags in a repository, including:
  - The tag name
  - The compressed size
  - The manifest digest
  - The last updated timestamp`),
					ox.Spec(`<repository> [flags]`),
					ox.Aliases("lt"),
					ox.Example(`
The following example lists tags in a repository named `+"`"+`example-repository`+"`"+` in a registry named `+"`"+`example-registry`+"`"+`. The command also uses the `+"`"+`--format`+"`"+` flag to return only the tag name and manifest digest for each tag: doctl registry repository list-tags example-registry/example-repository --format Tag,ManifestDigest`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Tag, `+"`"+`CompressedSizeBytes`+"`"+`, `+"`"+`UpdatedAt`+"`"+`, `+"`"+`ManifestDigest`+"`"+`.`, ox.Default("Tag")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl registry repository list-v2
					ox.Usage(`list-v2`, `List repositories for a container registry`),
					ox.Banner(`Retrieves information about repositories in a registry, including:
  - The repository name
  - The latest manifest of the repository
  - The latest manifest's latest tag, if any
  - The number of tags in the repository
  - The number of manifests in the repository`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls2"),
					ox.Example(`
The following example lists repositories in a registry named `+"`"+`example-registry`+"`"+` and uses the `+"`"+`--format`+"`"+` flag to return only the name and update time of each repository: doctl registry repository list-v2 example-registry --format Name,UpdatedAt`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Name, `+"`"+`LatestTag`+"`"+`, `+"`"+`TagCount`+"`"+`, `+"`"+`UpdatedAt`+"`"+`.`, ox.Default("Name")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
		),
		ox.Sub( // doctl serverless
			ox.Usage(`serverless`, `Develop, test, and deploy serverless functions`),
			ox.Banner(`The `+"`"+`doctl serverless`+"`"+` commands provide an environment for developing, testing, and deploying serverless functions.
One or more local file system areas are employed, along with one or more 'functions namespaces' in the cloud.
A one-time install of the serverless software is needed (use `+"`"+`doctl serverless install`+"`"+` to install the software,
then `+"`"+`doctl serverless connect`+"`"+` to connect to a functions namespace associated with your account).
Other `+"`"+`doctl serverless`+"`"+` commands are used to develop, test, and deploy.`),
			ox.Spec(`[command]`),
			ox.Aliases("sandbox", "sbx", "sls"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl serverless [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl serverless activations
				ox.Usage(`activations`, `Retrieve activation records`),
				ox.Banner(`The subcommands of `+"`"+`doctl serverless activations`+"`"+` retrieve results, logs, or complete
activation records of functions deployed to your functions namespace.`),
				ox.Spec(`[command]`),
				ox.Aliases("activation", "actv"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl serverless activations [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl serverless activations get
					ox.Usage(`get`, `Retrieve information about an activation.`),
					ox.Banner(`Retrieve the activation record for a previously invoked function. You can limit output to the result
or the logs.  The `+"`"+`doctl serverless activation logs`+"`"+` command has additional advanced capabilities for retrieving
logs.`),
					ox.Spec(`[<activationId>] [flags]`),
					ox.Example(`
The following example retrieves the results for the most recent activation of a function named `+"`"+`yourFunction`+"`"+`: doctl serverless activations get --function yourFunction --last --result`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`function`, `Retrieve activations for a specific function.`, ox.Short("f")).
						Bool(`last`, `Retrieve the most recent activation (default). Does not return activations for web-invoked functions.`, ox.Short("l")).
						Bool(`logs`, `Retrieve only the logs, stripped of time stamps and stream identifier.`, ox.Short("g")).
						Bool(`quiet`, `Suppress the last activation information header.`, ox.Short("q")).
						Bool(`result`, `Retrieve only the resulting output of a function.`, ox.Short("r")).
						Int(`skip`, `Exclude a specified number of activations from the returned list, starting with the most recent.`, ox.Short("s")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless activations list
					ox.Usage(`list`, `Lists activations for which records exist.`),
					ox.Banner(`Use `+"`"+`doctl serverless activations list`+"`"+` to list the activation records that are present in the cloud for previously
invoked functions.`),
					ox.Spec(`[<function_name>] [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists all of the activations for a function named `+"`"+`yourFunction`+"`"+` since January 1, 2023: doctl serverless activations list --function yourFunction --since 1672549200000`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`count`, `Return only the total number of activations.`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: Datetime, `+"`"+`Status`+"`"+`, `+"`"+`Kind`+"`"+`, `+"`"+`Version`+"`"+`, `+"`"+`ActivationID`+"`"+`, `+"`"+`Start`+"`"+`, `+"`"+`Wait`+"`"+`, `+"`"+`Duration`+"`"+`, `+"`"+`Function`+"`"+`.`, ox.Default("Datetime")).
						Bool(`full`, `Include the full activation description.`, ox.Short("f")).
						Int(`limit`, `Limit the number of activations returned to the specified amount. Default: 30, Maximum: 200`, ox.Default("30"), ox.Short("l")).
						Bool(`no-header`, `Return raw data with no headers`).
						Int(`since`, `Retrieve activations invoked after the specified date-time, in UNIX timestamp format measured in milliseconds.`).
						Int(`skip`, `Exclude a specified number of activations from the returned list, starting with the most recent.`, ox.Short("s")).
						Int(`upto`, `Retrieve activations invoked before the specified date-time; in UNIX timestamp format measured in milliseconds.`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless activations logs
					ox.Usage(`logs`, `Retrieve the logs for an activation.`),
					ox.Banner(`Use `+"`"+`doctl serverless activations logs`+"`"+` to retrieve the logs portion of one or more activation records
with various options, such as selecting by package or function, and optionally watching continuously
for new arrivals.`),
					ox.Spec(`[<activationId>] [flags]`),
					ox.Example(`
The following example retrieves the logs for the most recent activation of a function named `+"`"+`yourFunction`+"`"+`: doctl serverless activations logs --function yourFunction --last`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`follow`, `Continuously return log information.`).
						String(`function`, `Retrieve the logs for a specific function.`, ox.Short("f")).
						Int(`limit`, `Limit the number of logs returned to the specified amount, up to 200.`, ox.Default("1"), ox.Short("n")).
						String(`package`, `Retrieve the logs for a specific package.`, ox.Short("p")).
						Bool(`strip`, `Retrieves only the first line of output in the log, stripped of time stamps.`, ox.Short("r")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless activations result
					ox.Usage(`result`, `Retrieve the output for an activation.`),
					ox.Banner(`Retrieve just the results portion
of one or more activation records.`),
					ox.Spec(`[<activationId>] [flags]`),
					ox.Example(`
The following example retrieves the results for the most recent activation of a function named `+"`"+`yourFunction`+"`"+`: doctl serverless activations result --function yourFunction --last`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`function`, `Retrieve the results for a specific function.`, ox.Short("f")).
						Bool(`last`, `Retrieve the most recent activation result (default).`, ox.Short("l")).
						Int(`limit`, `Limit the number of results return to the specified number. (default 30, max 200)`, ox.Default("1"), ox.Short("n")).
						Bool(`quiet`, `Suppress last activation information header.`, ox.Short("q")).
						Int(`skip`, `Exclude a specified number of activations from the returned list, starting with the most recent.`, ox.Short("s")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl serverless connect
				ox.Usage(`connect`, `Connects local serverless support to a functions namespace`),
				ox.Banner(`This command connects `+"`"+`doctl serverless`+"`"+` support to a functions namespace of your choice.
The optional argument should be a (complete or partial) match to a namespace label or id.
If there is no argument, all namespaces are matched.  If the result is exactly one namespace,
you are connected to it.  If there are multiple namespaces, you have an opportunity to choose
the one you want from a dialog.  Use `+"`"+`doctl serverless namespaces`+"`"+` to create, delete, and
list your namespaces.`),
				ox.Spec(`[<hint>] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless deploy
				ox.Usage(`deploy`, `Deploy a functions project to your functions namespace`),
				ox.Banner(`At any time you can use `+"`"+`doctl serverless deploy`+"`"+` to upload the contents of a functions project in your file system for
testing in your serverless namespace.  The project must be organized in the fashion expected by an App Platform Functions
component.  The `+"`"+`doctl serverless init`+"`"+` command will create a properly organized directory for you to work in.`),
				ox.Spec(`<directory> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`apihost`, `API host to use`).
					String(`auth`, `OpenWhisk auth token to use`).
					String(`build-env`, `Path to build-time environment file`).
					String(`env`, `Path to runtime environment file`).
					String(`exclude`, `Functions and/or packages to exclude`).
					String(`include`, `Functions and/or packages to include`).
					Bool(`incremental`, `Deploy only changes since last deploy`).
					Bool(`insecure`, `Ignore SSL Certificates`).
					Bool(`remote-build`, `Run builds remotely`).
					Bool(`verbose-build`, `Display build details`).
					Bool(`verbose-zip`, `Display start/end of zipping phase for each function`).
					Bool(`yarn`, `Use yarn instead of npm for node builds`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless functions
				ox.Usage(`functions`, `Work with the functions in your namespace`),
				ox.Banner(`The subcommands of `+"`"+`doctl serverless functions`+"`"+` manage your functions namespace, such as listing, reviewing, and invoking your functions.`),
				ox.Spec(`[command]`),
				ox.Aliases("function", "fn"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl serverless functions [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl serverless functions get
					ox.Usage(`get`, `Retrieve the metadata or code of a deployed function`),
					ox.Banner(`Retrieves the code or metadata of a deployed function.`),
					ox.Spec(`<functionName> [flags]`),
					ox.Example(`
The following example retrieves the code for a function named "example/helloWorld" and saves it to a file named `+"`"+`local-helloWorld.py`+"`"+`: doctl serverless functions get example/helloWorld --code --save-as local-helloWorld.py`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`code`, `Retrieves the functions code. This does not work if the function is saved as a zip file.`).
						Bool(`save`, `Saves the function's code to a local file`).
						String(`save-as`, `Saves the file as the specified name`).
						String(`save-env`, `Saves the function's environment variables to a local file as key-value pairs`, ox.Short("E")).
						String(`save-env-json`, `Saves the function's environment variables to a local file as JSON`, ox.Short("J")).
						Bool(`url`, `Retrieves function URL`, ox.Short("r")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless functions invoke
					ox.Usage(`invoke`, `Invokes a function`),
					ox.Banner(`Invokes a function in your functions namespace.
You can provide inputs and inspect outputs.`),
					ox.Spec(`<functionName> [flags]`),
					ox.Example(`
The following example invokes a function named "example/helloWorld" with the parameters `+"`"+`name:John,place:NY`+"`"+`: doctl serverless functions invoke example/helloWorld --param name:John,place:NY`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`full`, `Waits for the function to complete and then outputs the function's response along with its complete activation record. The record contains log information, duration time, and other information about the function's execution.`, ox.Short("f")).
						Bool(`no-wait`, `Asynchronously invokes the function and does not wait for the result to be returned. An activation ID is returned in the response, instead.`, ox.Short("n")).
						String(`param`, `Key-value pairs of input parameters. For example, if your function takes two parameters called name and `+"`"+`place`+"`"+`, you would provide them as `+"`"+`name:John,place:NY`+"`"+`.`, ox.Spec(`name`), ox.Short("p")).
						String(`param-file`, `A path to a file containing parameter values in JSON format, such as path/to/file.json.`, ox.Spec(`path/to/file.json`), ox.Short("P")).
						Bool(`web`, `Invokes the function as a web function and displays the result in your browser`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless functions list
					ox.Usage(`list`, `Lists the functions in your functions namespace`),
					ox.Banner(`Lists the functions in your functions namespace.`),
					ox.Spec(`[<packageName>] [flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists the three most recently updated functions in the `+"`"+`example`+"`"+` package: doctl serverless functions list example --limit 3`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`count`, `Returns only the total number of functions in the namespace`).
						String(`format`, `Columns for output in a comma-separated list. Possible values: Update, `+"`"+`Version`+"`"+`, `+"`"+`Runtime`+"`"+`, `+"`"+`Function`+"`"+`.`, ox.Default("Update")).
						String(`limit`, `Returns the specified number of functions in the result, starting with the most recently updated function.`, ox.Short("l")).
						Bool(`name`, `Sorts results by name in alphabetical order`, ox.Short("n")).
						Bool(`name-sort`, `Sorts results by name in alphabetical order`).
						Bool(`no-header`, `Return raw data with no headers`).
						Int(`skip`, `Excludes the specified number of functions from the result, starting with the most recently updated function. For example, if you specify 2, the most recently updated function and the function updated before that are excluded from the result.`, ox.Default("2"), ox.Short("s")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl serverless get-metadata
				ox.Usage(`get-metadata`, `Obtain metadata of a functions project`),
				ox.Banner(`The `+"`"+`doctl serverless get-metadata`+"`"+` command produces a JSON structure that summarizes the contents of a functions
project (a directory you have designated for functions development).  This can be useful for feeding into other tools.`),
				ox.Spec(`<directory> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`env`, `Path to environment file`).
					String(`exclude`, `Functions or packages to exclude`).
					String(`include`, `Functions or packages to include`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless init
				ox.Usage(`init`, `Initialize a 'functions project' directory in your local file system`),
				ox.Banner(`The `+"`"+`doctl serverless init`+"`"+` command specifies a directory in your file system which will hold functions and
supporting artifacts while you're developing them.  This 'functions project' can be uploaded to your functions namespace for testing.
Later, after the functions project is committed to a `+"`"+`git`+"`"+` repository, you can create an app, or an app component, from it.

Type `+"`"+`doctl serverless status --languages`+"`"+` for a list of supported languages.  Use one of the displayed keywords
to choose your sample language for `+"`"+`doctl serverless init`+"`"+`.`),
				ox.Spec(`<path> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`language`, `Language for the initial sample code`, ox.Default("javascript"), ox.Short("l")).
					Bool(`overwrite`, `Clears and reuses an existing directory`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless install
				ox.Usage(`install`, `Installs the serverless support`),
				ox.Banner(`This command installs additional software under `+"`"+`doctl`+"`"+` needed to make the other serverless commands work.
The install operation is long-running, and a network connection is required.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless namespaces
				ox.Usage(`namespaces`, `Manage your functions namespaces`),
				ox.Banner(`Functions namespaces (in the cloud) contain the result of deploying packages and functions with `+"`"+`doctl serverless deploy`+"`"+`.
The subcommands of `+"`"+`doctl serverless namespaces`+"`"+` are used to manage multiple functions namespaces within your account.
Use `+"`"+`doctl serverless connect`+"`"+` with an explicit argument to connect to a specific namespace.  You are connected to one namespace at a time.`),
				ox.Spec(`[command]`),
				ox.Aliases("namespace", "ns"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl serverless namespaces [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl serverless namespaces create
					ox.Usage(`create`, `Creates a namespace`),
					ox.Banner(``+"`"+`Use `+"`"+`doctl serverless namespaces create`+"`"+` to create a new functions namespace.
Both a region and a label must be specified.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`label`, `the label for the namespace (required)`, ox.Short("l")).
						Bool(`no-connect`, `don't immediately connect to the created namespace`, ox.Short("n")).
						String(`region`, `the region for the namespace (required)`, ox.Short("r")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless namespaces delete
					ox.Usage(`delete`, `Deletes a namespace`),
					ox.Banner(`Use `+"`"+`doctl serverless namespaces delete`+"`"+` to delete a functions namespace.
The full label or full id of the namespace is required as an argument.
You are prompted for confirmation unless `+"`"+`--force`+"`"+` is specified.`),
					ox.Spec(`<namespaceIdOrLabel> [flags]`),
					ox.Aliases("rm"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Just do it, omitting confirmatory prompt`, ox.Short("f")).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless namespaces list
					ox.Usage(`list`, `Lists your namespaces`),
					ox.Banner(`Use `+"`"+`doctl serverless namespaces list`+"`"+` to list your functions namespaces.`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: Label, `+"`"+`Region`+"`"+`, `+"`"+`ID`+"`"+`, `+"`"+`Host`+"`"+`.`, ox.Default("Label")).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl serverless namespaces list-regions
					ox.Usage(`list-regions`, `Lists the accepted 'region' values`),
					ox.Banner(`Use `+"`"+`doctl serverless namespaces list-regions`+"`"+` to list the values that are accepted
in the `+"`"+`--region`+"`"+` flag of `+"`"+`doctl serverless namespaces create`+"`"+`.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl serverless status
				ox.Usage(`status`, `Provide information about serverless support`),
				ox.Banner(`This command reports the status of serverless support and some details concerning its connected functions namespace.
With the `+"`"+`--languages flag, it will report the supported languages.
With the `+"`"+`--version flag, it will show just version information about the serverless component`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`languages`, `show available languages (if connected to the cloud)`, ox.Short("l")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless undeploy
				ox.Usage(`undeploy`, `Removes resources from your functions namespace`),
				ox.Banner(`This command removes functions, entire packages, or all functions and packages, from your function
namespace.  In general, deploying new content does not remove old content although it may overwrite it.
Use `+"`"+`doctl serverless undeploy`+"`"+` to effect removal.  The command accepts a list of functions or packages.
Functions should be listed in `+"`"+`pkgName/fnName`+"`"+` form, or `+"`"+`fnName`+"`"+` for a function not in any package.
The `+"`"+`--packages`+"`"+` flag causes arguments without slash separators to be interpreted as packages, in which case
the entire packages are removed.`),
				ox.Spec(`[<package|function>...] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`all`, `remove all packages and functions`).
					Bool(`packages`, `interpret simple name arguments as packages`, ox.Short("p")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless uninstall
				ox.Usage(`uninstall`, `Removes the serverless support`),
				ox.Banner(`Removes serverless support from `+"`"+`doctl`+"`"+``),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless upgrade
				ox.Usage(`upgrade`, `Upgrades serverless support to match this version of doctl`),
				ox.Banner(`This command upgrades the serverless support software under `+"`"+`doctl`+"`"+` by installing over the existing version.
The install operation is long-running, and a network connection is required.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl serverless watch
				ox.Usage(`watch`, `Watch a functions project directory, deploying incrementally on change`),
				ox.Banner(`Type `+"`"+`doctl serverless watch <directory>`+"`"+` in a separate terminal window.  It will run until interrupted.
It will watch the directory (which should be one you initialized for serverless development) and will deploy
the contents to the cloud incrementally as it detects changes.`),
				ox.Spec(`<directory> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`apihost`, `API host to use`).
					String(`auth`, `OpenWhisk auth token to use`).
					String(`build-env`, `Path to build-time environment file`).
					String(`env`, `Path to runtime environment file`).
					String(`exclude`, `Functions and/or packages to exclude`).
					String(`include`, `Functions and/or packages to include`).
					Bool(`insecure`, `Ignore SSL Certificates`).
					Bool(`remote-build`, `Run builds remotely`).
					Bool(`verbose-build`, `Display build details`).
					Bool(`verbose-zip`, `Display start/end of zipping phase for each function`).
					Bool(`yarn`, `Use yarn instead of npm for node builds`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl vpcs
			ox.Usage(`vpcs`, `Display commands that manage VPCs`),
			ox.Banner(`The commands under `+"`"+`doctl vpcs`+"`"+` are for managing your VPC networks.

With the VPC commands, you can list, create, or delete VPCs, and manage their configuration details.`),
			ox.Spec(`[command]`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl vpcs [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl vpcs create
				ox.Usage(`create`, `Create a new VPC network`),
				ox.Banner(`Use this command to create a new VPC network on your account.`),
				ox.Spec(`[flags]`),
				ox.Aliases("c"),
				ox.Example(`
The following example creates a VPC network named `+"`"+`example-vpc`+"`"+` in the `+"`"+`nyc1`+"`"+` region: doctl vpcs create --name example-vpc --region nyc1`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`description`, `A description of the VPC network`).
					CIDR(`ip-range`, `The range of IP addresses in the VPC network, in CIDR notation, such as 10.116.0.0/20. If not specified, we generate a range for you.`, ox.Default("10.116.0.0/20")).
					String(`name`, `The VPC network's name (required)`).
					String(`region`, `The VPC network's region slug, such as nyc1 (required)`, ox.Default("nyc1")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl vpcs delete
				ox.Usage(`delete`, `Permanently delete a VPC network`),
				ox.Banner(`Permanently deletes the specified VPC. This is irreversible.
		
		You cannot delete VPCs that are default networks for a region. To delete a default VPC network, make another VPC network the default for the region using the `+"`"+`doctl vpcs update <vpc-network-id> --default=true`+"`"+` command, and then delete the target VPC network.`),
				ox.Spec(`<vpc-id> [flags]`),
				ox.Aliases("d", "rm"),
				ox.Example(`
The following example deletes the VPC network with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl vpcs delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `Delete the VPC without a confirmation prompt`, ox.Short("f")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl vpcs get
				ox.Usage(`get`, `Retrieve a VPC network`),
				ox.Banner(`Retrieve information about a VPC network, including:

- The VPC network's ID
- The uniform resource name (URN) for the VPC network
- The VPC network's name
- The VPC network's description
- The range of IP addresses in the VPC network, in CIDR notation
- The datacenter region slug the VPC network is located in
- The VPC network's default boolean value indicating whether or not it is the default one for the region
- The VPC network's creation date, in ISO8601 combined date and time format`),
				ox.Spec(`<vpc-id> [flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves information about a VPC network with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl vpcs get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`URN`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`IPRange`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Default`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl vpcs list
				ox.Usage(`list`, `List VPC networks`),
				ox.Banner(`Retrieves a list of the VPCs on your account, including the following information for each:

- The VPC network's ID
- The uniform resource name (URN) for the VPC network
- The VPC network's name
- The VPC network's description
- The range of IP addresses in the VPC network, in CIDR notation
- The datacenter region slug the VPC network is located in
- The VPC network's default boolean value indicating whether or not it is the default one for the region
- The VPC network's creation date, in ISO8601 combined date and time format`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example lists the VPCs on your account and uses the --format flag to return only the name, IP range, and region for each VPC network: doctl vpcs list --format Name,IPRange,Region`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`URN`+"`"+`, `+"`"+`Name`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`IPRange`+"`"+`, `+"`"+`Region`+"`"+`, `+"`"+`Created`+"`"+`, `+"`"+`Default`+"`"+`.`, ox.Spec(`ID`)).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl vpcs peerings
				ox.Usage(`peerings`, `Display commands that manage VPC Peerings`),
				ox.Banner(`The commands under `+"`"+`doctl vpcs peerings`+"`"+` are for managing your VPC Peerings.
With the VPC Peerings commands, you can get, list, create, update, or delete VPC Peerings, and manage their configuration details.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "doctl vpcs peerings [command] --help" for more information about a command.`),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				ox.Sub( // doctl vpcs peerings create
					ox.Usage(`create`, `Create a new VPC Peering`),
					ox.Banner(`Use this command to create a new VPC Peering on your account.`),
					ox.Spec(`[flags]`),
					ox.Aliases("c"),
					ox.Example(`
The following example creates a VPC Peering named `+"`"+`example-peering-name`+"`"+` : doctl vpcs peerings create example-peering-name --vpc-ids f81d4fae-7dec-11d0-a765-00a0c91e6bf6,3f900b61-30d7-40d8-9711-8c5d6264b268`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`vpc-ids`, `Peering VPC IDs should be comma separated (required)`).
						Bool(`wait`, `Boolean that specifies whether to wait for a VPC Peering creation to complete before returning control to the terminal`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl vpcs peerings delete
					ox.Usage(`delete`, `Permanently delete a VPC Peering`),
					ox.Banner(`Permanently deletes the specified VPC Peering. This is irreversible.`),
					ox.Spec(`<peering-id> [flags]`),
					ox.Aliases("d", "rm"),
					ox.Example(`
The following example deletes the VPC Peering with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl vpcs peerings delete f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `Delete the VPC Peering without any confirmation prompt`, ox.Short("f")).
						Bool(`wait`, `Boolean that specifies whether to wait for a VPC Peering deletion to complete before returning control to the terminal`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl vpcs peerings get
					ox.Usage(`get`, `Retrieves a VPC Peering`),
					ox.Banner(`Retrieves information about a VPC Peering, including:
- The VPC Peering ID
- The VPC Peering Name
- The Peered VPC network IDs
- The VPC Peering Status
- The VPC Peering creation date, in ISO8601 combined date and time format`),
					ox.Spec(`<peering-id> [flags]`),
					ox.Aliases("g"),
					ox.Example(`
The following example retrieves information about a VPC Peering with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl vpcs peerings get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`VPCIDs`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl vpcs peerings list
					ox.Usage(`list`, `List VPC Peerings`),
					ox.Banner(`Retrieves a list of the VPC Peerings on your account, including the following information for each:
- The VPC Peering ID
- The VPC Peering Name
- The Peered VPC network IDs
- The VPC Peering Status
- The VPC Peering creation date, in ISO8601 combined date and time format`),
					ox.Spec(`[flags]`),
					ox.Aliases("ls"),
					ox.Example(`
The following example lists the VPC Peerings on your account : doctl vpcs peerings list --format Name,VPCIDs`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`format`, `Columns for output in a comma-separated list. Possible values: ID, `+"`"+`Name`+"`"+`, `+"`"+`VPCIDs`+"`"+`, `+"`"+`Status`+"`"+`, `+"`"+`Created`+"`"+`.`, ox.Spec(`ID`)).
						Bool(`no-header`, `Return raw data with no headers`).
						String(`vpc-id`, `VPC ID`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
				ox.Sub( // doctl vpcs peerings update
					ox.Usage(`update`, `Update a VPC Peering's name`),
					ox.Banner(`Use this command to update the name of a VPC Peering`),
					ox.Spec(`<peering-id> [flags]`),
					ox.Aliases("u"),
					ox.Example(`
The following example updates the name of a VPC Peering with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to `+"`"+`new-name`+"`"+`: doctl vpcs peerings update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name new-name`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`name`, `The VPC Peering's name (required)`).
						String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
						String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
						String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
						String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
						Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
						Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
						String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
						Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
						Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
				),
			),
			ox.Sub( // doctl vpcs update
				ox.Usage(`update`, `Update a VPC network's configuration`),
				ox.Banner(`Updates a VPC network's configuration. You can update its name, description, and default state.`),
				ox.Spec(`<vpc-id> [flags]`),
				ox.Aliases("u"),
				ox.Example(`
The following example updates the name of a VPC network with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to `+"`"+`new-name`+"`"+`: doctl vpcs update f81d4fae-7dec-11d0-a765-00a0c91e6bf6 --name new-name --default=true`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`default`, `A boolean value indicating whether or not the VPC network is the default one for the region`).
					String(`description`, `The VPC network's description`).
					String(`name`, `The VPC network's name`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl auth
			ox.Usage(`auth`, `Display commands for authenticating doctl with an account`),
			ox.Banner(`The `+"`"+`doctl auth`+"`"+` commands allow you to authenticate doctl for use with your DigitalOcean account using tokens that you generate in the control panel at https://cloud.digitalocean.com/account/api/tokens.

If you work with a just one account, call `+"`"+`doctl auth init`+"`"+` and supply the token when prompted. This creates an authentication context named `+"`"+`default`+"`"+`.

To switch between multiple DigitalOcean accounts, including team accounts, create named contexts using `+"`"+`doctl auth init --context <name>`+"`"+`, then providing the applicable token when prompted. This saves the token under the name you provide. To switch between contexts, use `+"`"+`doctl auth switch --context <name>`+"`"+`.

To remove accounts from the configuration file, run `+"`"+`doctl auth remove --context <name>`+"`"+`. This removes the token under the name you provide.`),
			ox.Spec(`[command]`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl auth [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl auth init
				ox.Usage(`init`, `Initialize doctl to use a specific account`),
				ox.Banner(`This command allows you to initialize doctl with a token that allows it to query and manage your account details and resources.

The command requires and API token to authenticate, which you can generate in the control panel at https://cloud.digitalocean.com/account/api/tokens.

The `+"`"+`--context`+"`"+` flag allows you to add authentication for multiple accounts and then switch between them as needed. Provide a case-sensitive name for the context and then enter the API token you want use for that context when prompted. You can switch authentication contexts using `+"`"+`doctl auth switch`+"`"+`, which re-initializes doctl. You can also provide the `+"`"+`--context`+"`"+` flag when using any doctl command to specify the auth context for that command. This enables you to use multiple DigitalOcean accounts with doctl, or tokens that have different authentication scopes.

If the `+"`"+`--context`+"`"+` flag is not specified, doctl creates a default authentication context named `+"`"+`default`+"`"+`.

You can use doctl without initializing it by adding the `+"`"+`--access-token`+"`"+` flag to each command and providing an API token as the argument.`),
				ox.Spec(`[flags]`),
				ox.Example(`
The following example initializes doctl with a token for a single account with the context `+"`"+`your-team`+"`"+`: doctl auth init --context your-team`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`token-validation-server`, `The server used to validate a token`, ox.Default("https://cloud.digitalocean.com")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl auth list
				ox.Usage(`list`, `List available authentication contexts`),
				ox.Banner(`List named authentication contexts that you created with `+"`"+`doctl auth init`+"`"+`.

To switch between the contexts use `+"`"+`doctl auth switch --context <name>`+"`"+`, where `+"`"+`<name>`+"`"+` is one of the contexts listed.

To create new contexts, see the help for `+"`"+`doctl auth init`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example lists the available contexts with the `+"`"+`--format`+"`"+` flag: doctl auth list`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: text`, ox.Default("text")).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl auth remove
				ox.Usage(`remove`, `Remove authentication contexts`),
				ox.Banner(`This command allows you to remove authentication contexts you've already created.

To see a list of available authentication contexts, call `+"`"+`doctl auth list`+"`"+`.

For details on creating an authentication context, see the help for `+"`"+`doctl auth init`+"`"+`.`),
				ox.Spec(`--context <name> [flags]`),
				ox.Example(`
The following example removes the context `+"`"+`your-team`+"`"+`: doctl auth remove --context your-team`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl auth switch
				ox.Usage(`switch`, `Switch between authentication contexts`),
				ox.Banner(`This command allows you to switch between authentication contexts you've already created.

To see a list of available authentication contexts, call `+"`"+`doctl auth list`+"`"+`.

For details on creating an authentication context, see the help for `+"`"+`doctl auth init`+"`"+`.`),
				ox.Spec(`[flags]`),
				ox.Example(`
The following example switches to the context `+"`"+`your-team`+"`"+`: doctl auth switch --context your-team`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl balance
			ox.Usage(`balance`, `Display commands for retrieving your account balance`),
			ox.Banner(`The subcommands of `+"`"+`doctl balance`+"`"+` retrieve information about your account balance.`),
			ox.Spec(`[command]`),
			ox.Section(2),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl balance [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl balance get
				ox.Usage(`get`, `Retrieve your account balance`),
				ox.Banner(`This command retrieves the following details about your account balance:

- Your month-to-date balance including your account
  balance and month-to-date usage.
- Your current overall balance as of your most recent billing activity.
- Your usage in the current billing period.
- The time at which balances were most recently generated.`),
				ox.Spec(`[flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example uses the `+"`"+`--format`+"`"+` flag to display only your month-to-date balance: doctl balance get --format MonthToDateUsage`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: MonthToDateBalance, `+"`"+`AccountBalance`+"`"+`, `+"`"+`MonthToDateUsage`+"`"+`, `+"`"+`GeneratedAt`+"`"+`.`, ox.Default("MonthToDateBalance")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl billing-history
			ox.Usage(`billing-history`, `Display commands for retrieving your billing history`),
			ox.Banner(`The subcommands of `+"`"+`doctl billing-history`+"`"+` are used to access the billing history for your DigitalOcean account.`),
			ox.Spec(`[command]`),
			ox.Aliases("bh"),
			ox.Section(2),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl billing-history [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl billing-history list
				ox.Usage(`list`, `Retrieve a paginated billing history for a user`),
				ox.Banner(`This command retrieves the following details for each event in your billing history:
- The date of the event
- The type of billing event
- A description of the event
- The amount of the event in USD
- The invoice ID associated with the event, if applicable
- The invoice UUID associated with the event, if applicable`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example uses the `+"`"+`--format`+"`"+` flag to display only the date and description of each event in your billing history: doctl billing-history list --format Date,Description`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: Date, `+"`"+`Type`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`Amount`+"`"+`, `+"`"+`InvoiceID`+"`"+`, `+"`"+`InvoiceUUID`+"`"+`.`, ox.Default("Date")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
		ox.Sub( // doctl invoice
			ox.Usage(`invoice`, `Display commands for retrieving invoices for your account`),
			ox.Banner(`The subcommands of `+"`"+`doctl invoice`+"`"+` retrieve details about invoices for your account.`),
			ox.Spec(`[command]`),
			ox.Section(2),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "doctl invoice [command] --help" for more information about a command.`),
			ox.Flags().
				String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
				String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
				String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
				String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
				Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
				Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
				String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
				Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
				Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			ox.Sub( // doctl invoice csv
				ox.Usage(`csv`, `Downloads a CSV file of a specific invoice to you local machine`),
				ox.Banner(`Downloads a CSV-formatted file of a specific invoice to your local machine.

Use the `+"`"+`doctl invoice list`+"`"+` command to find the UUID of the invoice to retrieve.`),
				ox.Spec(`<invoice-uuid> <output-file.csv> [flags]`),
				ox.Aliases("c"),
				ox.Example(`
The following example downloads a CSV summary of an invoice with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to the file `+"`"+`invoice.csv`+"`"+`: doctl invoice csv f81d4fae-7dec-11d0-a765-00a0c91e6bf6 invoice.csv`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl invoice get
				ox.Usage(`get`, `Retrieve a list of all the items on an invoice`),
				ox.Banner(`Retrieves an itemized list of resources and their costs on the specified invoice, including each resource's:
- ID
- UUID (if applicable)
- Product name
- Description
- Group description
- Amount charged, in USD
- Duration of usage for the invoice period
- Duration unit of measurement, such as hours
- The start time of the invoice period, in ISO8601 combined date and time format
- The end time of the invoice period, in ISO8601 combined date and time format
- The project name the resource belongs to
- Category, such as "iaas"

Use the `+"`"+`doctl invoice list`+"`"+` command to find the UUID of the invoice to retrieve.`),
				ox.Spec(`<invoice-uuid> [flags]`),
				ox.Aliases("g"),
				ox.Example(`
The following example retrieves details about an invoice with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl invoice get f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ResourceID, `+"`"+`ResourceUUID`+"`"+`, `+"`"+`Product`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`GroupDescription`+"`"+`, `+"`"+`Amount`+"`"+`, `+"`"+`Duration`+"`"+`, `+"`"+`DurationUnit`+"`"+`, `+"`"+`StartTime`+"`"+`, `+"`"+`EndTime`+"`"+`, `+"`"+`ProjectName`+"`"+`, `+"`"+`Category`+"`"+`.`, ox.Default("ResourceID")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl invoice list
				ox.Usage(`list`, `List all of the invoices for your account`),
				ox.Banner(`Lists all of the invoices on your account including the UUID, amount in USD, and time period for each.`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Example(`
The following example lists all of the invoices on your account and uses the `+"`"+`--format`+"`"+` flag to only return the product name and the amount charged for it: doctl invoice list --format Product,Amount`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ResourceID, `+"`"+`ResourceUUID`+"`"+`, `+"`"+`Product`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`GroupDescription`+"`"+`, `+"`"+`Amount`+"`"+`, `+"`"+`Duration`+"`"+`, `+"`"+`DurationUnit`+"`"+`, `+"`"+`StartTime`+"`"+`, `+"`"+`EndTime`+"`"+`, `+"`"+`ProjectName`+"`"+`, `+"`"+`Category`+"`"+`.`, ox.Default("ResourceID")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl invoice pdf
				ox.Usage(`pdf`, `Downloads a PDF file of a specific invoice to your local machine`),
				ox.Banner(`This command downloads a PDF summary of a specific invoice to the provided location.

Use the `+"`"+`doctl invoice list`+"`"+` command to find the UUID of the invoice to retrieve.`),
				ox.Spec(`<invoice-uuid> <output-file.pdf> [flags]`),
				ox.Aliases("p"),
				ox.Example(`
The following example downloads a PDF summary of an invoice with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+` to the file `+"`"+`invoice.pdf`+"`"+`: doctl invoice pdf f81d4fae-7dec-11d0-a765-00a0c91e6bf6 invoice.pdf`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
			ox.Sub( // doctl invoice summary
				ox.Usage(`summary`, `Get a summary of an invoice`),
				ox.Banner(`Retrieves a summary of an invoice, including the following details:

- The invoice's UUID
- The year and month of the billing period
- The total amount of the invoice, in USD
- The name of the user associated with the invoice
- The company associated with the invoice
- The email address associated with the invoice
- The amount of product usage charges contributing to the invoice
- The amount of overage charges contributing to the invoice, such as bandwidth overages
- The amount of taxes contributing to the invoice
- The amount of any credits or other adjustments contributing to the invoice

Use the `+"`"+`doctl invoice list`+"`"+` command to find the UUID of the invoice to retrieve.`),
				ox.Spec(`<invoice-uuid> [flags]`),
				ox.Aliases("s"),
				ox.Example(`
The following example retrieves a summary of an invoice with the ID `+"`"+`f81d4fae-7dec-11d0-a765-00a0c91e6bf6`+"`"+`: doctl invoice summary f81d4fae-7dec-11d0-a765-00a0c91e6bf6`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`format`, `Columns for output in a comma-separated list. Possible values: ResourceID, `+"`"+`ResourceUUID`+"`"+`, `+"`"+`Product`+"`"+`, `+"`"+`Description`+"`"+`, `+"`"+`GroupDescription`+"`"+`, `+"`"+`Amount`+"`"+`, `+"`"+`Duration`+"`"+`, `+"`"+`DurationUnit`+"`"+`, `+"`"+`StartTime`+"`"+`, `+"`"+`EndTime`+"`"+`, `+"`"+`ProjectName`+"`"+`, `+"`"+`Category`+"`"+`.`, ox.Default("ResourceID")).
					Bool(`no-header`, `Return raw data with no headers`).
					String(`access-token`, `API V2 access token`, ox.Short("t"), ox.Section(0)).
					String(`api-url`, `Override default API endpoint`, ox.Short("u"), ox.Section(0)).
					String(`config`, `Specify a custom config file`, ox.Default("$APPCONFIG/config.yaml"), ox.Short("c"), ox.Section(0)).
					String(`context`, `Specify a custom authentication context name`, ox.Section(0)).
					Int(`http-retry-max`, `Set maximum number of retries for requests that fail with a 429 or 500-level error`, ox.Default("5"), ox.Section(0)).
					Bool(`interactive`, `Enable interactive behavior. Defaults to true if the terminal supports it`, ox.Default("false"), ox.Section(0)).
					String(`output`, `Desired output format [text|json]`, ox.Default("text"), ox.Short("o"), ox.Section(0)).
					Bool(`trace`, `Show a log of network activity while performing a command`, ox.Section(0)).
					Bool(`verbose`, `Enable verbose output`, ox.Short("v"), ox.Section(0)),
			),
		),
	)
}
