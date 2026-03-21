// Command omnictl is a xo/ox version of `omnictl`.
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
		ox.Usage(`omnictl`, ``),
		ox.Banner(`A CLI for accessing Omni API.`),
		ox.Spec(`[command]`),
		ox.Footer(`Use "omnictl [command] --help" for more information about a command.`),
		ox.Flags().
			String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`).
			Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`).
			String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`),
		ox.Sub( // omnictl apply
			ox.Usage(`apply`, `Create or update resource using YAML file as an input`),
			ox.Banner(`Create or update resource using YAML file as an input`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`dry-run`, `Dry run, implies verbose`, ox.Short("d")).
				String(`file`, `Resource file to load and apply`, ox.Short("f")).
				Bool(`verbose`, `Verbose output`, ox.Short("v")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl audit-log
			ox.Usage(`audit-log`, `Read audit log from Omni`),
			ox.Banner(`Read audit log from Omni`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl cluster
			ox.Usage(`cluster`, `Cluster-related subcommands.`),
			ox.Banner(`Commands to destroy clusters and manage cluster templates.`),
			ox.Spec(`[command]`),
			ox.Aliases("c"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "omnictl cluster [command] --help" for more information about a command.`),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			ox.Sub( // omnictl cluster delete
				ox.Usage(`delete`, `Delete all cluster resources.`),
				ox.Banner(`Delete all resources related to the cluster. The command waits for the cluster to be fully destroyed.`),
				ox.Spec(`cluster-name [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`destroy-disconnected-machines`, `removes all disconnected machines which are part of the cluster from Omni`).
					Bool(`dry-run`, `dry run`, ox.Short("d")).
					Bool(`verbose`, `verbose output (show diff for each resource)`, ox.Short("v")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl cluster kubernetes
				ox.Usage(`kubernetes`, `Cluster Kubernetes management subcommands.`),
				ox.Banner(`Commands to render, validate, manage cluster templates.`),
				ox.Spec(`[command]`),
				ox.Aliases("k"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "omnictl cluster kubernetes [command] --help" for more information about a command.`),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				ox.Sub( // omnictl cluster kubernetes manifest-sync
					ox.Usage(`manifest-sync`, `Sync Kubernetes bootstrap manifests from Talos controlplane nodes to Kubernetes API.`),
					ox.Banner(`Sync Kubernetes bootstrap manifests from Talos controlplane nodes to Kubernetes API.
Bootstrap manifests might be updated with Talos version update, Kubernetes upgrade, and config patching.
Talos never updates or deletes Kubernetes manifests, so this command fills the gap to keep manifests up-to-date.`),
					ox.Spec(`cluster-name [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`dry-run`, `don't actually sync manifests, just print what would be done`, ox.Default("true")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster kubernetes upgrade-pre-checks
					ox.Usage(`upgrade-pre-checks`, `Run Kubernetes upgrade pre-checks for the cluster.`),
					ox.Banner(`Verify that upgrading Kubernetes version is available for the cluster: version compatibility, deprecated APIs, etc.`),
					ox.Spec(`cluster-name [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`to`, `target Kubernetes version for the planned upgrade`).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
			),
			ox.Sub( // omnictl cluster machine
				ox.Usage(`machine`, `Machine related commands.`),
				ox.Banner(`Commands to manage cluster machines.`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "omnictl cluster machine [command] --help" for more information about a command.`),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				ox.Sub( // omnictl cluster machine delete
					ox.Usage(`delete`, `Delete the machine from the cluster`),
					ox.Banner(`Delete the machine from the cluster. The command waits for the machine to be fully deleted.`),
					ox.Spec(`machine-id [flags]`),
					ox.Aliases("rm", "destroy"),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`force`, `force destroy the machine`, ox.Short("f")).
						Duration(`timeout`, `timeout for the machine deletion`, ox.Default("5m0s"), ox.Short("t")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster machine lock
					ox.Usage(`lock`, `Lock the machine`),
					ox.Banner(`When locked, no config updates, upgrades and downgrades will be performed on the machine.`),
					ox.Spec(`machine-id [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster machine unlock
					ox.Usage(`unlock`, `Unlock the machine`),
					ox.Banner(`Removes locked annotation from the machine.`),
					ox.Spec(`machine-id [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
			),
			ox.Sub( // omnictl cluster status
				ox.Usage(`status`, `Show cluster status, wait for the cluster to be ready.`),
				ox.Banner(`Shows current cluster status, if the terminal supports it, watch the status as it updates. The command waits for the cluster to be ready by default.`),
				ox.Spec(`cluster-name [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`quiet`, `suppress output`, ox.Short("q")).
					Duration(`wait`, `wait timeout, if zero, report current status and exit`, ox.Default("5m0s"), ox.Short("w")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl cluster template
				ox.Usage(`template`, `Cluster template management subcommands.`),
				ox.Banner(`Commands to render, validate, manage cluster templates.`),
				ox.Spec(`[command]`),
				ox.Aliases("t"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "omnictl cluster template [command] --help" for more information about a command.`),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				ox.Sub( // omnictl cluster template delete
					ox.Usage(`delete`, `Delete all cluster template resources from Omni.`),
					ox.Banner(`Delete all resources related to the cluster template. This command requires API access.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`destroy-disconnected-machines`, `removes all disconnected machines which are part of the cluster from Omni`).
						Bool(`dry-run`, `dry run`, ox.Short("d")).
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						Bool(`verbose`, `verbose output (show diff for each resource)`, ox.Short("v")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template diff
					ox.Usage(`diff`, `Show diff in resources if the template is synced.`),
					ox.Banner(`Query existing resources for the cluster and compare them with the resources generated from the template. This command requires API access.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template export
					ox.Usage(`export`, `Export a cluster template from an existing cluster on Omni.`),
					ox.Banner(`Export a cluster template from an existing cluster on Omni. This command requires API access.`),
					ox.Spec(`cluster-name [flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cluster`, `cluster name`, ox.Short("c")).
						Bool(`force`, `overwrite output file if it exists`, ox.Short("f")).
						String(`output`, `output file (default: stdout)`, ox.Short("o")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template render
					ox.Usage(`render`, `Render a cluster template to a set of resources.`),
					ox.Banner(`Validate template contents, convert to resources and output resources to stdout as YAML. This command is offline (doesn't access API).`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template status
					ox.Usage(`status`, `Show template cluster status, wait for the cluster to be ready.`),
					ox.Banner(`Shows current cluster status, if the terminal supports it, watch the status as it updates. The command waits for the cluster to be ready by default.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						Bool(`quiet`, `suppress output`, ox.Short("q")).
						Duration(`wait`, `wait timeout, if zero, report current status and exit`, ox.Default("5m0s"), ox.Short("w")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template sync
					ox.Usage(`sync`, `Apply template to the Omni.`),
					ox.Banner(`Query existing resources for the cluster and compare them with the resources generated from the template, create/update/delete resources as needed. This command requires API access.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						Bool(`dry-run`, `dry run`, ox.Short("d")).
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						Bool(`verbose`, `verbose output (show diff for each resource)`, ox.Short("v")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
				ox.Sub( // omnictl cluster template validate
					ox.Usage(`validate`, `Validate a cluster template.`),
					ox.Banner(`Validate that template contains valid structures, and there are no other warnings. This command is offline (doesn't access API).`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`file`, `path to the cluster template file.`, ox.Short("f")).
						String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
						Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
						String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
				),
			),
		),
		ox.Sub( // omnictl config
			ox.Usage(`config`, `Manage the client configuration file (omniconfig)`),
			ox.Banner(`Manage the client configuration file (omniconfig)`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "omnictl config [command] --help" for more information about a command.`),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			ox.Sub( // omnictl config add
				ox.Usage(`add`, `Add a new context`),
				ox.Banner(`Add a new context`),
				ox.Spec(`<context> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`identity`, `identity to use for authentication`).
					String(`url`, `URL of the server`, ox.Default("grpc://127.0.0.1:8080")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config context
				ox.Usage(`context`, `Set the current context`),
				ox.Banner(`Set the current context`),
				ox.Spec(`<context> [flags]`),
				ox.Aliases("use-context"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config contexts
				ox.Usage(`contexts`, `List defined contexts`),
				ox.Banner(`List defined contexts`),
				ox.Spec(`[flags]`),
				ox.Aliases("get-contexts"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config identity
				ox.Usage(`identity`, `Set the auth identity for the current context`),
				ox.Banner(`Set the auth identity for the current context`),
				ox.Spec(`<identity> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config info
				ox.Usage(`info`, `Show information about the current context`),
				ox.Banner(`Show information about the current context`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config merge
				ox.Usage(`merge`, `Merge additional contexts from another client configuration file`),
				ox.Banner(`Contexts with the same name are renamed while merging configs.`),
				ox.Spec(`<from> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config new
				ox.Usage(`new`, `Generate a new client configuration file`),
				ox.Banner(`Generate a new client configuration file`),
				ox.Spec(`[<path>] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`identity`, `identity to use for authentication`).
					String(`url`, `URL of the server`, ox.Default("grpc://127.0.0.1:8080")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl config url
				ox.Usage(`url`, `Set the URL for the current context`),
				ox.Banner(`Set the URL for the current context`),
				ox.Spec(`<url> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
		),
		ox.Sub( // omnictl delete
			ox.Usage(`delete`, `Delete a specific resource by ID or all resources of the type.`),
			ox.Banner(`Similar to 'kubectl delete', 'omnictl delete' initiates resource deletion and waits for the operation to complete.`),
			ox.Spec(`<type> [<id>] [flags]`),
			ox.Aliases("d"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`all`, `Delete all resources of the type.`).
				String(`namespace`, `The resource namespace.`, ox.Default("default"), ox.Short("n")).
				String(`selector`, `Selector (label query) to filter on, supports '=' and '==' (e.g. -l key1=value1,key2=value2)`, ox.Short("l")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl download
			ox.Usage(`download`, `Download installer media`),
			ox.Banner(`This command downloads installer media from the server

It accepts one argument, which is the name of the image to download. Name can be one of the following:

     * iso - downloads the latest ISO image
     * AWS AMI (amd64), Vultr (arm64), Raspberry Pi 4 Model B - full image name
     * oracle, aws, vmware - platform name
     * rpi_generic, rockpi_4c, rock64 - board name

To get the full list of available images, look at the output of the following command:
    omnictl get installationmedia -o yaml

The download command tries to match the passed string in this order:

    * name
    * profile

By default it will download amd64 image if there are multiple images available for the same name.

For example, to download the latest ISO image for arm64, run:

    omnictl download iso --arch amd64

To download the same ISO with two extensions added, the --extensions argument gets repeated to produce a stringArray:

    omnictl download iso --arch amd64 --extensions intel-ucode --extensions qemu-guest-agent

To download the latest Vultr image, run:

    omnictl download "vultr"

To download the latest Radxa ROCK PI 4 image, run:

    omnictl download "rpi_generic"`),
			ox.Spec(`<image name> [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`arch`, `Image architecture to download (amd64, arm64)`, ox.Default("amd64")).
				Slice(`extensions`, `Generate installation media with extensions pre-installed`).
				Array(`extra-kernel-args`, `Add extra kernel args to the generated installation media`).
				Slice(`initial-labels`, `Bake initial labels into the generated installation media`).
				String(`output`, `Output file or directory, defaults to current working directory`, ox.Default(".")).
				Bool(`pxe`, `Print PXE URL and exit`).
				Bool(`secureboot`, `Download SecureBoot enabled installation media`).
				String(`talos-version`, `Talos version to be used in the generated installation media`, ox.Default("1.10.1")).
				Bool(`use-siderolink-grpc-tunnel`, `Configure Talos to use the SideroLink (WireGuard) gRPC tunnel over HTTP2 for Omni management traffic, instead of UDP. Note that this will add overhead to the traffic.`).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl get
			ox.Usage(`get`, `Get a specific resource or list of resources.`),
			ox.Banner(`Similar to 'kubectl get', 'omnictl get' returns a set of resources from the OS.
To get a list of all available resource definitions, issue 'omnictl get rd'`),
			ox.Spec(`<type> [<id>] [flags]`),
			ox.Aliases("g"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`id-match-regexp`, `Match resource ID against a regular expression.`).
				String(`namespace`, `The resource namespace.`, ox.Default("default"), ox.Short("n")).
				String(`output`, `Output format (json, table, yaml, jsonpath).`, ox.Default("table"), ox.Short("o")).
				String(`selector`, `Selector (label query) to filter on, supports '=' and '==' (e.g. -l key1=value1,key2=value2)`, ox.Short("l")).
				Bool(`watch`, `Watch the resource state.`, ox.Short("w")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl infraprovider
			ox.Usage(`infraprovider`, `Manage infra providers`),
			ox.Banner(`Manage infra providers`),
			ox.Spec(`[command]`),
			ox.Aliases("ip"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "omnictl infraprovider [command] --help" for more information about a command.`),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			ox.Sub( // omnictl infraprovider create
				ox.Usage(`create`, `Create an infra provider`),
				ox.Banner(`Create an infra provider`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("c"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Duration(`ttl`, `TTL for the infra provider service account key`, ox.Default("8760h0m0s"), ox.Short("t")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl infraprovider delete
				ox.Usage(`delete`, `Delete an infra provider`),
				ox.Banner(`Delete an infra provider`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("d"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl infraprovider list
				ox.Usage(`list`, `List infra providers`),
				ox.Banner(`List infra providers`),
				ox.Spec(`[flags]`),
				ox.Aliases("l"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl infraprovider renewkey
				ox.Usage(`renewkey`, `Renew an infra provider service account by registering a new public key to it`),
				ox.Banner(`Renew an infra provider service account by registering a new public key to it`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("r"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Duration(`ttl`, `TTL for the infra provider service account key`, ox.Default("8760h0m0s"), ox.Short("t")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
		),
		ox.Sub( // omnictl kubeconfig
			ox.Usage(`kubeconfig`, `Download the admin kubeconfig of a cluster`),
			ox.Banner(`Download the admin kubeconfig of a cluster.
If merge flag is defined, config will be merged with ~/.kube/config or [local-path] if specified.
Otherwise kubeconfig will be written to PWD or [local-path] if specified.`),
			ox.Spec(`[local-path] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`break-glass`, `get kubeconfig that allows accessing nodes bypasing Omni (if enabled for the account)`).
				String(`cluster`, `cluster to use`, ox.Short("c")).
				Bool(`force`, `force overwrite of kubeconfig if already present, force overwrite on kubeconfig merge`, ox.Short("f")).
				String(`force-context-name`, `force context name for kubeconfig merge`).
				String(`grant-type`, `Authorization grant type to use. One of (auto|authcode|authcode-keyboard)`).
				Slice(`groups`, `group to be used in the service account token (groups). only used when --service-account is set to true`, ox.Default("[system:masters]")).
				Bool(`merge`, `merge with existing kubeconfig`, ox.Default("true"), ox.Short("m")).
				Bool(`service-account`, `create a service account type kubeconfig instead of a OIDC-authenticated user type`).
				Duration(`ttl`, `ttl for the service account token. only used when --service-account is set to true`, ox.Default("8760h0m0s")).
				String(`user`, `user to be used in the service account token (sub). required when --service-account is set to true`).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl machine-logs
			ox.Usage(`machine-logs`, `Get logs for a machine`),
			ox.Banner(`Get logs for a provided machine id`),
			ox.Spec(`machineID [flags]`),
			ox.Aliases("l"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`follow`, `specify if the logs should be streamed`, ox.Short("f")).
				String(`log-format`, `log format (raw, omni, dmesg) to display (default is to display in raw format)`, ox.Default("raw")).
				Int32(`tail`, `lines of log file to display (default is to show from the beginning)`, ox.Default("-1")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl serviceaccount
			ox.Usage(`serviceaccount`, `Manage service accounts`),
			ox.Banner(`Manage service accounts`),
			ox.Spec(`[command]`),
			ox.Aliases("sa"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "omnictl serviceaccount [command] --help" for more information about a command.`),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			ox.Sub( // omnictl serviceaccount create
				ox.Usage(`create`, `Create a service account`),
				ox.Banner(`Create a service account`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("c"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`role`, `role of the service account. only used when --use-user-role=false`, ox.Short("r")).
					Duration(`ttl`, `TTL for the service account key`, ox.Default("8760h0m0s"), ox.Short("t")).
					Bool(`use-user-role`, `use the role of the creating user. if true, --role is ignored`, ox.Default("true"), ox.Short("u")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl serviceaccount destroy
				ox.Usage(`destroy`, `Destroy a service account`),
				ox.Banner(`Destroy a service account`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("d"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl serviceaccount list
				ox.Usage(`list`, `List service accounts`),
				ox.Banner(`List service accounts`),
				ox.Spec(`[flags]`),
				ox.Aliases("l"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl serviceaccount renew
				ox.Usage(`renew`, `Renew a service account by registering a new public key to it`),
				ox.Banner(`Renew a service account by registering a new public key to it`),
				ox.Spec(`<name> [flags]`),
				ox.Aliases("r"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Duration(`ttl`, `TTL for the service account key`, ox.Default("8760h0m0s"), ox.Short("t")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
		),
		ox.Sub( // omnictl support
			ox.Usage(`support`, `Download the support bundle for a cluster`),
			ox.Banner(`The command collects all non-sensitive information for the cluster from the Omni state.`),
			ox.Spec(`[local-path] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `cluster to use`, ox.Short("c")).
				String(`output`, `support bundle output`, ox.Default("support.zip"), ox.Short("O")).
				Bool(`verbose`, `verbose output`, ox.Short("v")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl talosconfig
			ox.Usage(`talosconfig`, `Download an admin talosconfig.`),
			ox.Banner(`Download the generic admin talosconfig of the Omni instance or the admin talosconfig of a cluster.
Generic talosconfig can be used with any machine, including those in maintenance mode.
If merge flag is defined, config will be merged with ~/.talos/config or [local-path] if specified.
Otherwise talosconfig will be written to PWD or [local-path] if specified.`),
			ox.Spec(`[local-path] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`break-glass`, `get operator talosconfig that allows bypassing Omni (if enabled for the account)`).
				String(`cluster`, `cluster to use. If omitted, download the generic talosconfig for the Omni instance.`, ox.Short("c")).
				Bool(`force`, `force overwrite of talosconfig if already present`, ox.Short("f")).
				Bool(`merge`, `merge with existing talosconfig`, ox.Default("true"), ox.Short("m")).
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
		),
		ox.Sub( // omnictl user
			ox.Usage(`user`, `User-related subcommands.`),
			ox.Banner(`Commands to manage users.`),
			ox.Spec(`[command]`),
			ox.Aliases("u"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "omnictl user [command] --help" for more information about a command.`),
			ox.Flags().
				String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
				Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
				String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			ox.Sub( // omnictl user create
				ox.Usage(`create`, `Create a user.`),
				ox.Banner(`Create a user with the specified email.`),
				ox.Spec(`[email] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`role`, `Role to use for the user creation`, ox.Short("r")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl user delete
				ox.Usage(`delete`, `Delete users.`),
				ox.Banner(`Delete users with the specified emails.`),
				ox.Spec(`[email1 email2] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl user list
				ox.Usage(`list`, `List all users.`),
				ox.Banner(`List all existing users on the Omni instance.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
			ox.Sub( // omnictl user set-role
				ox.Usage(`set-role`, `Update the role of the user.`),
				ox.Banner(`Update the user role.`),
				ox.Spec(`[email] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`role`, `Role to use`, ox.Short("r")).
					String(`context`, `The context to be used. Defaults to the selected context in the omniconfig file.`, ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `Skip TLS verification for the Omni GRPC and HTTP API endpoints.`, ox.Section(0)).
					String(`omniconfig`, `The path to the omni configuration file. Defaults to 'OMNICONFIG' env variable if set, otherwise the config directory according to the XDG specification.`, ox.Section(0)),
			),
		),
	)
}
