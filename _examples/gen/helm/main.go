// Command helm is a xo/ox version of `helm`.
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
		ox.Usage(`helm`, ``),
		ox.Banner(`The Kubernetes package manager

Common actions for Helm:

- helm search:    search for charts
- helm pull:      download a chart to your local directory to view
- helm install:   upload the chart to Kubernetes
- helm list:      list releases of charts

Environment variables:

| Name                               | Description                                                                                                |
|------------------------------------|------------------------------------------------------------------------------------------------------------|
| $HELM_CACHE_HOME                   | set an alternative location for storing cached files.                                                      |
| $HELM_CONFIG_HOME                  | set an alternative location for storing Helm configuration.                                                |
| $HELM_DATA_HOME                    | set an alternative location for storing Helm data.                                                         |
| $HELM_DEBUG                        | indicate whether or not Helm is running in Debug mode                                                      |
| $HELM_DRIVER                       | set the backend storage driver. Values are: configmap, secret, memory, sql.                                |
| $HELM_DRIVER_SQL_CONNECTION_STRING | set the connection string the SQL storage driver should use.                                               |
| $HELM_MAX_HISTORY                  | set the maximum number of helm release history.                                                            |
| $HELM_NAMESPACE                    | set the namespace used for the helm operations.                                                            |
| $HELM_NO_PLUGINS                   | disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.                                                 |
| $HELM_PLUGINS                      | set the path to the plugins directory                                                                      |
| $HELM_REGISTRY_CONFIG              | set the path to the registry config file.                                                                  |
| $HELM_REPOSITORY_CACHE             | set the path to the repository cache directory                                                             |
| $HELM_REPOSITORY_CONFIG            | set the path to the repositories file.                                                                     |
| $KUBECONFIG                        | set an alternative Kubernetes configuration file (default "~/.kube/config")                                |
| $HELM_KUBEAPISERVER                | set the Kubernetes API Server Endpoint for authentication                                                  |
| $HELM_KUBECAFILE                   | set the Kubernetes certificate authority file.                                                             |
| $HELM_KUBEASGROUPS                 | set the Groups to use for impersonation using a comma-separated list.                                      |
| $HELM_KUBEASUSER                   | set the Username to impersonate for the operation.                                                         |
| $HELM_KUBECONTEXT                  | set the name of the kubeconfig context.                                                                    |
| $HELM_KUBETOKEN                    | set the Bearer KubeToken used for authentication.                                                          |
| $HELM_KUBEINSECURE_SKIP_TLS_VERIFY | indicate if the Kubernetes API server's certificate validation should be skipped (insecure)                |
| $HELM_KUBETLS_SERVER_NAME          | set the server name used to validate the Kubernetes API server certificate                                 |
| $HELM_BURST_LIMIT                  | set the default burst limit in the case the server contains many CRDs (default 100, -1 to disable)         |
| $HELM_QPS                          | set the Queries Per Second in cases where a high number of calls exceed the option for higher burst values |

Helm stores cache, configuration, and data based on the following configuration order:

- If a HELM_*_HOME environment variable is set, it will be used
- Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used
- When no other location is set a default location will be used based on the operating system

By default, the default directories depend on the Operating System. The defaults are listed below:

| Operating System | Cache Path                | Configuration Path             | Data Path               |
|------------------|---------------------------|--------------------------------|-------------------------|
| Linux            | $HOME/.cache/helm         | $HOME/.config/helm             | $HOME/.local/share/helm |
| macOS            | $HOME/Library/Caches/helm | $HOME/Library/Preferences/helm | $HOME/Library/helm      |
| Windows          | %TEMP%\helm               | %APPDATA%\helm                 | %APPDATA%\helm          |`),
		ox.Spec(`[command]`),
		ox.Footer(`Use "helm [command] --help" for more information about a command.`),
		ox.Flags().
			Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100")).
			Bool(`debug`, `enable verbose output`).
			String(`kube-apiserver`, `the address and the port for the Kubernetes API server`).
			Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`).
			String(`kube-as-user`, `username to impersonate for the operation`).
			String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`).
			String(`kube-context`, `name of the kubeconfig context to use`).
			Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`).
			String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`).
			String(`kube-token`, `bearer token used for authentication`).
			String(`kubeconfig`, `path to the kubeconfig file`).
			String(`namespace`, `namespace scope for this request`, ox.Short("n")).
			Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`).
			String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json")).
			String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository")).
			String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml")),
		ox.Sub(
			ox.Usage(`create`, `create a new chart with the given name`),
			ox.Banner(`
This command creates a chart directory along with the common files and
directories used in a chart.

For example, 'helm create foo' will create a directory structure that looks
something like this:

    foo/
    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.
    ├── Chart.yaml    # Information about your chart
    ├── values.yaml   # The default values for your templates
    ├── charts/       # Charts that this chart depends on
    └── templates/    # The template files
        └── tests/    # The test files

'helm create' takes a path for an argument. If directories in the given path
do not exist, Helm will attempt to create them as it goes. If the given
destination exists and there are files in that directory, conflicting files
will be overwritten, but other files will be left alone.`),
			ox.Spec(`NAME [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`starter`, `the name or absolute path to Helm starter scaffold`, ox.Short("p")).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`dependency`, `manage a chart's dependencies`),
			ox.Banner(`
Manage the dependencies of a chart.

Helm charts store their dependencies in 'charts/'. For chart developers, it is
often easier to manage dependencies in 'Chart.yaml' which declares all
dependencies.

The dependency commands operate on that file, making it easy to synchronize
between the desired dependencies and the actual dependencies stored in the
'charts/' directory.

For example, this Chart.yaml declares two dependencies:

    # Chart.yaml
    dependencies:
    - name: nginx
      version: "1.2.3"
      repository: "https://example.com/charts"
    - name: memcached
      version: "3.2.1"
      repository: "https://another.example.com/charts"


The 'name' should be the name of a chart, where that name must match the name
in that chart's 'Chart.yaml' file.

The 'version' field should contain a semantic version or version range.

The 'repository' URL should point to a Chart Repository. Helm expects that by
appending '/index.yaml' to the URL, it should be able to retrieve the chart
repository's index. Note: 'repository' can be an alias. The alias must start
with 'alias:' or '@'.

Starting from 2.2.0, repository can be defined as the path to the directory of
the dependency charts stored locally. The path should start with a prefix of
"file://". For example,

    # Chart.yaml
    dependencies:
    - name: nginx
      version: "1.2.3"
      repository: "file://../dependency_chart/nginx"

If the dependency chart is retrieved locally, it is not required to have the
repository added to helm by "helm add repo". Version matching is also supported
for this case.`),
			ox.Spec(`[command]`),
			ox.Aliases("dep", "dependencies"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm dependency [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`build`, `rebuild the charts/ directory based on the Chart.lock file`),
				ox.Banner(`
Build out the charts/ directory from the Chart.lock file.

Build is used to reconstruct a chart's dependencies to the state specified in
the lock file. This will not re-negotiate dependencies, as 'helm dependency update'
does.

If no lock file is found, 'helm dependency build' will mirror the behavior
of 'helm dependency update'.`),
				ox.Spec(`CHART [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `keyring containing public keys`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					Bool(`skip-refresh`, `do not refresh the local repository cache`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the packages against signatures`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `list the dependencies for the given chart`),
				ox.Banner(`
List all of the dependencies declared in a chart.

This can take chart archives and chart directories as input. It will not alter
the contents of a chart.

This will produce an error if the chart cannot be loaded.`),
				ox.Spec(`CHART [flags]`),
				ox.Aliases("ls"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Uint(`max-col-width`, `maximum column width for output table`, ox.Default("80")).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`update`, `update charts/ based on the contents of Chart.yaml`),
				ox.Banner(`
Update the on-disk dependencies to mirror Chart.yaml.

This command verifies that the required charts, as expressed in 'Chart.yaml',
are present in 'charts/' and are at an acceptable version. It will pull down
the latest charts that satisfy the dependencies, and clean up old dependencies.

On successful update, this will generate a lock file that can be used to
rebuild the dependencies to an exact version.

Dependencies are not required to be represented in 'Chart.yaml'. For that
reason, an update command will not remove charts unless they are (a) present
in the Chart.yaml file, but (b) at the wrong version.`),
				ox.Spec(`CHART [flags]`),
				ox.Aliases("up"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `keyring containing public keys`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					Bool(`skip-refresh`, `do not refresh the local repository cache`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the packages against signatures`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`env`, `helm client environment information`),
			ox.Banner(`
Env prints out all the environment information in use by Helm.`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`get`, `download extended information of a named release`),
			ox.Banner(`
This command consists of multiple subcommands which can be used to
get extended information about the release, including:

- The values used to generate the release
- The generated manifest file
- The notes provided by the chart of the release
- The hooks associated with the release
- The metadata of the release`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm get [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`all`, `download all information for a named release`),
				ox.Banner(`
This command prints a human readable collection of information about the
notes, hooks, supplied values, and generated manifest file of the given release.`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`revision`, `get the named release with revision`).
					String(`template`, `go template for formatting the output, eg: {{.Release.Name}}`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`hooks`, `download all hooks for a named release`),
				ox.Banner(`
This command downloads hooks for a given release.

Hooks are formatted in YAML and separated by the YAML '---\n' separator.`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`revision`, `get the named release with revision`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`manifest`, `download the manifest for a named release`),
				ox.Banner(`
This command fetches the generated manifest for a given release.

A manifest is a YAML-encoded representation of the Kubernetes resources that
were generated from this release's chart(s). If a chart is dependent on other
charts, those resources will also be included in the manifest.`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`revision`, `get the named release with revision`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`metadata`, `This command fetches metadata for a given release`),
				ox.Banner(`This command fetches metadata for a given release`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
					Int(`revision`, `specify release revision`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`notes`, `download the notes for a named release`),
				ox.Banner(`
This command shows notes provided by the chart of a named release.`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`revision`, `get the named release with revision`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`values`, `download the values file for a named release`),
				ox.Banner(`
This command downloads a values file for a given release.`),
				ox.Spec(`RELEASE_NAME [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`all`, `dump all (computed) values`, ox.Short("a")).
					String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
					Int(`revision`, `get the named release with revision`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`history`, `fetch release history`),
			ox.Banner(`
History prints historical revisions for a given release.

A default maximum of 256 revisions will be returned. Setting '--max'
configures the maximum length of the revision list returned.

The historical release set is printed as a formatted table, e.g:

    $ helm history angry-bird
    REVISION    UPDATED                     STATUS          CHART             APP VERSION     DESCRIPTION
    1           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Initial install
    2           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Upgraded successfully
    3           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Rolled back to 2
    4           Mon Oct 3 10:15:13 2016     deployed        alpine-0.1.0      1.0             Upgraded successfully`),
			ox.Spec(`RELEASE_NAME [flags]`),
			ox.Aliases("hist"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Int(`max`, `maximum number of revision to include in history`, ox.Default("256")).
				String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`install`, `install a chart`),
			ox.Banner(`
This command installs a chart archive.

The install argument must be a chart reference, a path to a packaged chart,
a path to an unpacked chart directory or a URL.

To override values in a chart, use either the '--values' flag and pass in a file
or use the '--set' flag and pass configuration from the command line, to force
a string value use '--set-string'. You can use '--set-file' to set individual
values from a file when the value itself is too long for the command line
or is dynamically generated. You can also use '--set-json' to set json values
(scalars/objects/arrays) from the command line.

    $ helm install -f myvalues.yaml myredis ./redis

or

    $ helm install --set name=prod myredis ./redis

or

    $ helm install --set-string long_int=1234567890 myredis ./redis

or

    $ helm install --set-file my_script=dothings.sh myredis ./redis

or

    $ helm install --set-json 'master.sidecars=[{"name":"sidecar","image":"myImage","imagePullPolicy":"Always","ports":[{"name":"portname","containerPort":1234}]}]' myredis ./redis


You can specify the '--values'/'-f' flag multiple times. The priority will be given to the
last (right-most) file specified. For example, if both myvalues.yaml and override.yaml
contained a key called 'Test', the value set in override.yaml would take precedence:

    $ helm install -f myvalues.yaml -f override.yaml  myredis ./redis

You can specify the '--set' flag multiple times. The priority will be given to the
last (right-most) set specified. For example, if both 'bar' and 'newbar' values are
set for a key called 'foo', the 'newbar' value would take precedence:

    $ helm install --set foo=bar --set foo=newbar  myredis ./redis

Similarly, in the following example 'foo' is set to '["four"]':

    $ helm install --set-json='foo=["one", "two", "three"]' --set-json='foo=["four"]' myredis ./redis

And in the following example, 'foo' is set to '{"key1":"value1","key2":"bar"}':

    $ helm install --set-json='foo={"key1":"value1","key2":"value2"}' --set-json='foo.key2="bar"' myredis ./redis

To check the generated manifests of a release without installing the chart,
the --debug and --dry-run flags can be combined.

The --dry-run flag will output all generated chart manifests, including Secrets
which can contain sensitive values. To hide Kubernetes Secrets use the
--hide-secret flag. Please carefully consider how and when these flags are used.

If --verify is set, the chart MUST have a provenance file, and the provenance
file MUST pass all verification steps.

There are six different ways you can express the chart you want to install:

1. By chart reference: helm install mymaria example/mariadb
2. By path to a packaged chart: helm install mynginx ./nginx-1.2.3.tgz
3. By path to an unpacked chart directory: helm install mynginx ./nginx
4. By absolute URL: helm install mynginx https://example.com/charts/nginx-1.2.3.tgz
5. By chart reference and repo url: helm install --repo https://example.com/charts/ mynginx nginx
6. By OCI registries: helm install mynginx --version 1.2.3 oci://example.com/charts/nginx

CHART REFERENCES

A chart reference is a convenient way of referencing a chart in a chart repository.

When you use a chart reference with a repo prefix ('example/mariadb'), Helm will look in the local
configuration for a chart repository named 'example', and will then look for a
chart in that repository whose name is 'mariadb'. It will install the latest stable version of that chart
until you specify '--devel' flag to also include development version (alpha, beta, and release candidate releases), or
supply a version number with the '--version' flag.

To see the list of chart repositories, use 'helm repo list'. To search for
charts in a repository, use 'helm search'.`),
			ox.Spec(`[NAME] [CHART] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`atomic`, `if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used`).
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
				Bool(`create-namespace`, `create the release namespace if not present`).
				Bool(`dependency-update`, `update dependencies if they are missing before installing the chart`).
				String(`description`, `add a custom description`).
				Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
				Bool(`disable-openapi-validation`, `if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema`).
				Map(`dry-run`, `simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.`, ox.Spec(`string[="client"]`)).
				Bool(`enable-dns`, `enable DNS lookups when rendering templates`).
				Bool(`force`, `force resource updates through a replacement strategy`).
				Bool(`generate-name`, `generate the name (and omit the NAME parameter)`, ox.Short("g")).
				Bool(`hide-notes`, `if set, do not show notes in install output. Does not affect presence in chart metadata`).
				Bool(`hide-secret`, `hide Kubernetes Secrets when also using the --dry-run flag`).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
				String(`key-file`, `identify HTTPS client using this SSL key file`).
				String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				Map(`labels`, `Labels that would be added to release metadata. Should be divided by comma.`, ox.Default("[]"), ox.Short("l")).
				String(`name-template`, `specify template used to name the release`).
				Bool(`no-hooks`, `prevent hooks from running during install`).
				String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
				Bool(`pass-credentials`, `pass credentials to all domains`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
				String(`post-renderer`, `the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path`, ox.Spec(`postRendererString`)).
				String(`post-renderer-args`, `an argument to the post-renderer (can specify multiple)`, ox.Spec(`postRendererArgsSlice`), ox.Default("[]")).
				Bool(`render-subchart-notes`, `if set, render subchart notes along with the parent`).
				Bool(`replace`, `reuse the given name, only if that name is a deleted release which remains in the history. This is unsafe in production`).
				String(`repo`, `chart repository url where to locate the requested chart`).
				Array(`set`, `set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Array(`set-file`, `set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)`).
				Array(`set-json`, `set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)`).
				Array(`set-literal`, `set a literal STRING value on the command line`).
				Array(`set-string`, `set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Bool(`skip-crds`, `if set, no CRDs will be installed. By default, CRDs are installed if not already present`).
				Bool(`skip-schema-validation`, `if set, disables JSON schema validation`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				String(`username`, `chart repository username where to locate the requested chart`).
				Slice(`values`, `specify values in a YAML file or a URL (can specify multiple)`, ox.Short("f")).
				Bool(`verify`, `verify the package before using it`).
				Bool(`wait`, `if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout`).
				Bool(`wait-for-jobs`, `if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`lint`, `examine a chart for possible issues`),
			ox.Banner(`
This command takes a path to a chart and runs a series of tests to verify that
the chart is well-formed.

If the linter encounters things that will cause the chart to fail installation,
it will emit [ERROR] messages. If it encounters issues that break with convention
or recommendation, it will emit [WARNING] messages.`),
			ox.Spec(`PATH [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`kube-version`, `Kubernetes version used for capabilities and deprecation checks`).
				Bool(`quiet`, `print only warnings and errors`).
				Array(`set`, `set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Array(`set-file`, `set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)`).
				Array(`set-json`, `set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)`).
				Array(`set-literal`, `set a literal STRING value on the command line`).
				Array(`set-string`, `set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Bool(`skip-schema-validation`, `if set, disables JSON schema validation`).
				Bool(`strict`, `fail on lint warnings`).
				Slice(`values`, `specify values in a YAML file or a URL (can specify multiple)`, ox.Short("f")).
				Bool(`with-subcharts`, `lint dependent charts`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`list`, `list releases`),
			ox.Banner(`
This command lists all of the releases for a specified namespace (uses current namespace context if namespace not specified).

By default, it lists only releases that are deployed or failed. Flags like
'--uninstalled' and '--all' will alter this behavior. Such flags can be combined:
'--uninstalled --failed'.

By default, items are sorted alphabetically. Use the '-d' flag to sort by
release date.

If the --filter flag is provided, it will be treated as a filter. Filters are
regular expressions (Perl compatible) that are applied to the list of releases.
Only items that match the filter will be returned.

    $ helm list --filter 'ara[a-z]+'
    NAME                UPDATED                                  CHART
    maudlin-arachnid    2020-06-18 14:17:46.125134977 +0000 UTC  alpine-0.1.0

If no results are found, 'helm list' will exit 0, but with no output (or in
the case of no '-q' flag, only headers).

By default, up to 256 items may be returned. To limit this, use the '--max' flag.
Setting '--max' to 0 will not return all results. Rather, it will return the
server's default, which may be much higher than 256. Pairing the '--max'
flag with the '--offset' flag allows you to page through results.`),
			ox.Spec(`[flags]`),
			ox.Aliases("ls"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`all`, `show all releases without any filter applied`, ox.Short("a")).
				Bool(`all-namespaces`, `list releases across all namespaces`, ox.Short("A")).
				Bool(`date`, `sort by release date`, ox.Short("d")).
				Bool(`deployed`, `show deployed releases. If no other is specified, this will be automatically enabled`).
				Bool(`failed`, `show failed releases`).
				String(`filter`, `a regular expression (Perl compatible). Any releases that match the expression will be included in the results`, ox.Short("f")).
				Int(`max`, `maximum number of releases to fetch`, ox.Default("256"), ox.Short("m")).
				Bool(`no-headers`, `don't print headers when using the default output format`).
				Int(`offset`, `next release index in the list, used to offset from start value`).
				String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
				Bool(`pending`, `show pending releases`).
				Bool(`reverse`, `reverse the sort order`, ox.Short("r")).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Works only for secret(default) and configmap storage backends.`, ox.Short("l")).
				Bool(`short`, `output short (quiet) listing format`, ox.Short("q")).
				Bool(`superseded`, `show superseded releases`).
				String(`time-format`, `format time using golang time formatter. Example: --time-format "2006-01-02 15:04:05Z0700"`).
				Bool(`uninstalled`, `show uninstalled releases (if 'helm uninstall --keep-history' was used)`).
				Bool(`uninstalling`, `show releases that are currently being uninstalled`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`package`, `package a chart directory into a chart archive`),
			ox.Banner(`
This command packages a chart into a versioned chart archive file. If a path
is given, this will look at that path for a chart (which must contain a
Chart.yaml file) and then package that directory.

Versioned chart archives are used by Helm package repositories.

To sign a chart, use the '--sign' flag. In most cases, you should also
provide '--keyring path/to/secret/keys' and '--key keyname'.

  $ helm package --sign ./mychart --key mykey --keyring ~/.gnupg/secring.gpg

If '--keyring' is not specified, Helm usually defaults to the public keyring
unless your environment is otherwise configured.`),
			ox.Spec(`[CHART_PATH] [...] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`app-version`, `set the appVersion on the chart to this version`).
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
				Bool(`dependency-update`, `update dependencies from "Chart.yaml" to dir "charts/" before packaging`, ox.Short("u")).
				String(`destination`, `location to write the chart.`, ox.Default("."), ox.Short("d")).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
				String(`key`, `name of the key to use when signing. Used if --sign is true`).
				String(`key-file`, `identify HTTPS client using this SSL key file`).
				String(`keyring`, `location of a public keyring`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				String(`passphrase-file`, `location of a file which contains the passphrase for the signing key. Use "-" in order to read from stdin.`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
				Bool(`sign`, `use a PGP private key to sign this package`).
				String(`username`, `chart repository username where to locate the requested chart`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`plugin`, `install, list, or uninstall Helm plugins`),
			ox.Banner(`
Manage client-side Helm plugins.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm plugin [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`install`, `install a Helm plugin`),
				ox.Banner(`
This command allows you to install a plugin from a url to a VCS repo or a local path.`),
				ox.Spec(`[options] <path|url> [flags]`),
				ox.Aliases("add"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `list installed Helm plugins`),
				ox.Banner(`list installed Helm plugins`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`uninstall`, `uninstall one or more Helm plugins`),
				ox.Banner(`uninstall one or more Helm plugins`),
				ox.Spec(`<plugin>... [flags]`),
				ox.Aliases("rm", "remove"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`update`, `update one or more Helm plugins`),
				ox.Banner(`update one or more Helm plugins`),
				ox.Spec(`<plugin>... [flags]`),
				ox.Aliases("up"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`pull`, `download a chart from a repository and (optionally) unpack it in local directory`),
			ox.Banner(`
Retrieve a package from a package repository, and download it locally.

This is useful for fetching packages to inspect, modify, or repackage. It can
also be used to perform cryptographic verification of a chart without installing
the chart.

There are options for unpacking the chart after download. This will create a
directory for the chart and uncompress into that directory.

If the --verify flag is specified, the requested chart MUST have a provenance
file, and MUST pass the verification process. Failure in any part of this will
result in an error, and the chart will not be saved locally.`),
			ox.Spec(`[chart URL | repo/chartname] [...] [flags]`),
			ox.Aliases("fetch"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
				String(`destination`, `location to write the chart. If this and untardir are specified, untardir is appended to this`, ox.Default("."), ox.Short("d")).
				Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored.`).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
				String(`key-file`, `identify HTTPS client using this SSL key file`).
				String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				Bool(`pass-credentials`, `pass credentials to all domains`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
				Bool(`prov`, `fetch the provenance file, but don't perform verification`).
				String(`repo`, `chart repository url where to locate the requested chart`).
				Bool(`untar`, `if set to true, will untar the chart after downloading it`).
				String(`untardir`, `if untar is specified, this flag specifies the name of the directory into which the chart is expanded`, ox.Default(".")).
				String(`username`, `chart repository username where to locate the requested chart`).
				Bool(`verify`, `verify the package before using it`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`push`, `push a chart to remote`),
			ox.Banner(`
Upload a chart to a registry.

If the chart has an associated provenance file,
it will also be uploaded.`),
			ox.Spec(`[chart] [remote] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify registry client using this SSL certificate file`).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart upload`).
				String(`key-file`, `identify registry client using this SSL key file`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart upload`).
				String(`username`, `chart repository username where to locate the requested chart`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`registry`, `login to or logout from a registry`),
			ox.Banner(`
This command consists of multiple subcommands to interact with registries.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm registry [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`login`, `login to a registry`),
				ox.Banner(`
Authenticate to a remote registry.`),
				ox.Spec(`[host] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify registry client using this SSL certificate file`).
					Bool(`insecure`, `allow connections to TLS registry without certs`).
					String(`key-file`, `identify registry client using this SSL key file`).
					String(`password`, `registry password or identity token`, ox.Short("p")).
					Bool(`password-stdin`, `read password or identity token from stdin`).
					String(`username`, `registry username`, ox.Short("u")).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`logout`, `logout from a registry`),
				ox.Banner(`
Remove credentials stored for a remote registry.`),
				ox.Spec(`[host] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`repo`, `add, list, remove, update, and index chart repositories`),
			ox.Banner(`
This command consists of multiple subcommands to interact with chart repositories.

It can be used to add, remove, list, and index chart repositories.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm repo [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`add`, `add a chart repository`),
				ox.Banner(`add a chart repository`),
				ox.Spec(`[NAME] [URL] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`allow-deprecated-repos`, `by default, this command will not allow adding official repos that have been permanently deleted. This disables that behavior`).
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`force-update`, `replace (overwrite) the repo if it already exists`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the repository`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					Bool(`no-update`, `Ignored. Formerly, it would disabled forced updates. It is deprecated by force-update.`).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password`).
					Bool(`password-stdin`, `read chart repository password from stdin`).
					String(`username`, `chart repository username`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`index`, `generate an index file given a directory containing packaged charts`),
				ox.Banner(`
Read the current directory, generate an index file based on the charts found
and write the result to 'index.yaml' in the current directory.

This tool is used for creating an 'index.yaml' file for a chart repository. To
set an absolute URL to the charts, use '--url' flag.

To merge the generated index with an existing index file, use the '--merge'
flag. In this case, the charts found in the current directory will be merged
into the index passed in with --merge, with local charts taking priority over
existing charts.`),
				ox.Spec(`[DIR] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`json`, `output in JSON format`).
					String(`merge`, `merge the generated index into the given index`).
					String(`url`, `url of chart repository`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`list`, `list chart repositories`),
				ox.Banner(`list chart repositories`),
				ox.Spec(`[flags]`),
				ox.Aliases("ls"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`remove`, `remove one or more chart repositories`),
				ox.Banner(`remove one or more chart repositories`),
				ox.Spec(`[REPO1 [REPO2 ...]] [flags]`),
				ox.Aliases("rm"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`update`, `update information of available charts locally from chart repositories`),
				ox.Banner(`
Update gets the latest information about charts from the respective chart repositories.
Information is cached locally, where it is used by commands like 'helm search'.

You can optionally specify a list of repositories you want to update.
	$ helm repo update <repo_name> ...
To update all the repositories, use 'helm repo update'.`),
				ox.Spec(`[REPO1 [REPO2 ...]] [flags]`),
				ox.Aliases("up"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`fail-on-repo-update-fail`, `update fails if any of the repository updates fail`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`rollback`, `roll back a release to a previous revision`),
			ox.Banner(`
This command rolls back a release to a previous revision.

The first argument of the rollback command is the name of a release, and the
second is a revision (version) number. If this argument is omitted or set to
0, it will roll back to the previous release.

To see revision numbers, run 'helm history RELEASE'.`),
			ox.Spec(`<RELEASE> [REVISION] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`cleanup-on-fail`, `allow deletion of new resources created in this rollback when rollback fails`).
				Bool(`dry-run`, `simulate a rollback`).
				Bool(`force`, `force resource update through delete/recreate if needed`).
				Int(`history-max`, `limit the maximum number of revisions saved per release. Use 0 for no limit`, ox.Default("10")).
				Bool(`no-hooks`, `prevent hooks from running during rollback`).
				Bool(`recreate-pods`, `performs pods restart for the resource if applicable`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				Bool(`wait`, `if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout`).
				Bool(`wait-for-jobs`, `if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`search`, `search for a keyword in charts`),
			ox.Banner(`
Search provides the ability to search for Helm charts in the various places
they can be stored including the Artifact Hub and repositories you have added.
Use search subcommands to search different locations for charts.`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm search [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`hub`, `search for charts in the Artifact Hub or your own hub instance`),
				ox.Banner(`
Search for Helm charts in the Artifact Hub or your own hub instance.

Artifact Hub is a web-based application that enables finding, installing, and
publishing packages and configurations for CNCF projects, including publicly
available distributed charts Helm charts. It is a Cloud Native Computing
Foundation sandbox project. You can browse the hub at https://artifacthub.io/

The [KEYWORD] argument accepts either a keyword string, or quoted string of rich
query options. For rich query options documentation, see
https://artifacthub.github.io/hub/api/?urls.primaryName=Monocular%20compatible%20search%20API#/Monocular/get_api_chartsvc_v1_charts_search

Previous versions of Helm used an instance of Monocular as the default
'endpoint', so for backwards compatibility Artifact Hub is compatible with the
Monocular search API. Similarly, when setting the 'endpoint' flag, the specified
endpoint must also be implement a Monocular compatible search API endpoint.
Note that when specifying a Monocular instance as the 'endpoint', rich queries
are not supported. For API details, see https://github.com/helm/monocular`),
				ox.Spec(`[KEYWORD] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`endpoint`, `Hub instance to query for charts`, ox.Default("https://hub.$APPNAME.sh")).
					Bool(`fail-on-no-result`, `search fails if no results are found`).
					Bool(`list-repo-url`, `print charts repository URL`).
					Uint(`max-col-width`, `maximum column width for output table`, ox.Default("50")).
					String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`repo`, `search repositories for a keyword in charts`),
				ox.Banner(`
Search reads through all of the repositories configured on the system, and
looks for matches. Search of these repositories uses the metadata stored on
the system.

It will display the latest stable versions of the charts found. If you
specify the --devel flag, the output will include pre-release versions.
If you want to search using a version constraint, use --version.

Examples:

    # Search for stable release versions matching the keyword "nginx"
    $ helm search repo nginx

    # Search for release versions matching the keyword "nginx", including pre-release versions
    $ helm search repo nginx --devel

    # Search for the latest stable release for nginx-ingress with a major version of 1
    $ helm search repo nginx-ingress --version ^1.0.0

Repositories are managed with 'helm repo' commands.`),
				ox.Spec(`[keyword] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`devel`, `use development versions (alpha, beta, and release candidate releases), too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`fail-on-no-result`, `search fails if no results are found`).
					Uint(`max-col-width`, `maximum column width for output table`, ox.Default("50")).
					String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
					Bool(`regexp`, `use regular expressions for searching repositories you have added`, ox.Short("r")).
					Bool(`versions`, `show the long listing, with each version of each chart on its own line, for repositories you have added`, ox.Short("l")).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`show`, `show information of a chart`),
			ox.Banner(`
This command consists of multiple subcommands to display information about a chart`),
			ox.Spec(`[command]`),
			ox.Aliases("inspect"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "helm show [command] --help" for more information about a command.`),
			ox.Flags().
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			ox.Sub(
				ox.Usage(`all`, `show all information of the chart`),
				ox.Banner(`
This command inspects a chart (directory, file, or URL) and displays all its content
(values.yaml, Chart.yaml, README)`),
				ox.Spec(`[CHART] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					String(`repo`, `chart repository url where to locate the requested chart`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the package before using it`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`chart`, `show the chart's definition`),
				ox.Banner(`
This command inspects a chart (directory, file, or URL) and displays the contents
of the Chart.yaml file`),
				ox.Spec(`[CHART] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					String(`repo`, `chart repository url where to locate the requested chart`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the package before using it`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`crds`, `show the chart's CRDs`),
				ox.Banner(`
This command inspects a chart (directory, file, or URL) and displays the contents
of the CustomResourceDefinition files`),
				ox.Spec(`[CHART] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					String(`repo`, `chart repository url where to locate the requested chart`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the package before using it`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`readme`, `show the chart's README`),
				ox.Banner(`
This command inspects a chart (directory, file, or URL) and displays the contents
of the README file`),
				ox.Spec(`[CHART] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					String(`repo`, `chart repository url where to locate the requested chart`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the package before using it`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage(`values`, `show the chart's values`),
				ox.Banner(`
This command inspects a chart (directory, file, or URL) and displays the contents
of the values.yaml file`),
				ox.Spec(`[CHART] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
					String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
					Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
					Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
					String(`jsonpath`, `supply a JSONPath expression to filter the output`).
					String(`key-file`, `identify HTTPS client using this SSL key file`).
					String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool(`pass-credentials`, `pass credentials to all domains`).
					String(`password`, `chart repository password where to locate the requested chart`).
					Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
					String(`repo`, `chart repository url where to locate the requested chart`).
					String(`username`, `chart repository username where to locate the requested chart`).
					Bool(`verify`, `verify the package before using it`).
					Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
					Bool(`debug`, `enable verbose output`, ox.Section(0)).
					String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
					Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
					String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
					String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
					String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
					Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
					String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
					String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
					String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
					String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
					Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
					String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
					String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
					String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage(`status`, `display the status of the named release`),
			ox.Banner(`
This command shows the status of a named release.
The status consists of:
- last deployment time
- k8s namespace in which the release lives
- state of the release (can be: unknown, deployed, uninstalled, superseded, failed, uninstalling, pending-install, pending-upgrade or pending-rollback)
- revision of the release
- description of the release (can be completion message or error message)
- list of resources that this release consists of
- details on last test suite run, if applicable
- additional notes provided by the chart`),
			ox.Spec(`RELEASE_NAME [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
				Int(`revision`, `if set, display the status of the named release with revision`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`template`, `locally render templates`),
			ox.Banner(`
Render chart templates locally and display the output.

Any values that would normally be looked up or retrieved in-cluster will be
faked locally. Additionally, none of the server-side testing of chart validity
(e.g. whether an API is supported) is done.`),
			ox.Spec(`[NAME] [CHART] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Slice(`api-versions`, `Kubernetes api versions used for Capabilities.APIVersions`, ox.Short("a")).
				Bool(`atomic`, `if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used`).
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
				Bool(`create-namespace`, `create the release namespace if not present`).
				Bool(`dependency-update`, `update dependencies if they are missing before installing the chart`).
				String(`description`, `add a custom description`).
				Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
				Bool(`disable-openapi-validation`, `if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema`).
				Map(`dry-run`, `simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.`, ox.Spec(`string[="client"]`)).
				Bool(`enable-dns`, `enable DNS lookups when rendering templates`).
				Bool(`force`, `force resource updates through a replacement strategy`).
				Bool(`generate-name`, `generate the name (and omit the NAME parameter)`, ox.Short("g")).
				Bool(`hide-notes`, `if set, do not show notes in install output. Does not affect presence in chart metadata`).
				Bool(`include-crds`, `include CRDs in the templated output`).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
				Bool(`is-upgrade`, `set .Release.IsUpgrade instead of .Release.IsInstall`).
				String(`key-file`, `identify HTTPS client using this SSL key file`).
				String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				String(`kube-version`, `Kubernetes version used for Capabilities.KubeVersion`).
				Map(`labels`, `Labels that would be added to release metadata. Should be divided by comma.`, ox.Default("[]"), ox.Short("l")).
				String(`name-template`, `specify template used to name the release`).
				Bool(`no-hooks`, `prevent hooks from running during install`).
				String(`output-dir`, `writes the executed templates to files in output-dir instead of stdout`).
				Bool(`pass-credentials`, `pass credentials to all domains`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
				String(`post-renderer`, `the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path`, ox.Spec(`postRendererString`)).
				String(`post-renderer-args`, `an argument to the post-renderer (can specify multiple)`, ox.Spec(`postRendererArgsSlice`), ox.Default("[]")).
				Bool(`release-name`, `use release name in the output-dir path.`).
				Bool(`render-subchart-notes`, `if set, render subchart notes along with the parent`).
				Bool(`replace`, `reuse the given name, only if that name is a deleted release which remains in the history. This is unsafe in production`).
				String(`repo`, `chart repository url where to locate the requested chart`).
				Array(`set`, `set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Array(`set-file`, `set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)`).
				Array(`set-json`, `set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)`).
				Array(`set-literal`, `set a literal STRING value on the command line`).
				Array(`set-string`, `set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Array(`show-only`, `only show manifests rendered from the given templates`, ox.Short("s")).
				Bool(`skip-crds`, `if set, no CRDs will be installed. By default, CRDs are installed if not already present`).
				Bool(`skip-schema-validation`, `if set, disables JSON schema validation`).
				Bool(`skip-tests`, `skip tests from templated output`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				String(`username`, `chart repository username where to locate the requested chart`).
				Bool(`validate`, `validate your manifests against the Kubernetes cluster you are currently pointing at. This is the same validation performed on an install`).
				Slice(`values`, `specify values in a YAML file or a URL (can specify multiple)`, ox.Short("f")).
				Bool(`verify`, `verify the package before using it`).
				Bool(`wait`, `if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout`).
				Bool(`wait-for-jobs`, `if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`test`, `run tests for a release`),
			ox.Banner(`
The test command runs the tests for a release.

The argument this command takes is the name of a deployed release.
The tests to be run are defined in the chart that was installed.`),
			ox.Spec(`[RELEASE] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Slice(`filter`, `specify tests by attribute (currently "name") using attribute=value syntax or '!attribute=value' to exclude a test (can specify multiple or separate values with commas: name=test1,name=test2)`).
				Bool(`hide-notes`, `if set, do not show notes in test output. Does not affect presence in chart metadata`).
				Bool(`logs`, `dump the logs from test pods (this runs after all tests are complete, but before any cleanup)`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`uninstall`, `uninstall a release`),
			ox.Banner(`
This command takes a release name and uninstalls the release.

It removes all of the resources associated with the last release of the chart
as well as the release history, freeing it up for future use.

Use the '--dry-run' flag to see which releases will be uninstalled without actually
uninstalling them.`),
			ox.Spec(`RELEASE_NAME [...] [flags]`),
			ox.Aliases("del", "delete", "un"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cascade`, `Must be "background", "orphan", or "foreground". Selects the deletion cascading strategy for the dependents. Defaults to background.`, ox.Default("background")).
				String(`description`, `add a custom description`).
				Bool(`dry-run`, `simulate a uninstall`).
				Bool(`ignore-not-found`, `Treat "release not found" as a successful uninstall`).
				Bool(`keep-history`, `remove all associated resources and mark the release as deleted, but retain the release history`).
				Bool(`no-hooks`, `prevent hooks from running during uninstallation`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				Bool(`wait`, `if set, will wait until all the resources are deleted before returning. It will wait for as long as --timeout`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`upgrade`, `upgrade a release`),
			ox.Banner(`
This command upgrades a release to a new version of a chart.

The upgrade arguments must be a release and chart. The chart
argument can be either: a chart reference('example/mariadb'), a path to a chart directory,
a packaged chart, or a fully qualified URL. For chart references, the latest
version will be specified unless the '--version' flag is set.

To override values in a chart, use either the '--values' flag and pass in a file
or use the '--set' flag and pass configuration from the command line, to force string
values, use '--set-string'. You can use '--set-file' to set individual
values from a file when the value itself is too long for the command line
or is dynamically generated. You can also use '--set-json' to set json values
(scalars/objects/arrays) from the command line.

You can specify the '--values'/'-f' flag multiple times. The priority will be given to the
last (right-most) file specified. For example, if both myvalues.yaml and override.yaml
contained a key called 'Test', the value set in override.yaml would take precedence:

    $ helm upgrade -f myvalues.yaml -f override.yaml redis ./redis

You can specify the '--set' flag multiple times. The priority will be given to the
last (right-most) set specified. For example, if both 'bar' and 'newbar' values are
set for a key called 'foo', the 'newbar' value would take precedence:

    $ helm upgrade --set foo=bar --set foo=newbar redis ./redis

You can update the values for an existing release with this command as well via the
'--reuse-values' flag. The 'RELEASE' and 'CHART' arguments should be set to the original
parameters, and existing values will be merged with any values set via '--values'/'-f'
or '--set' flags. Priority is given to new values.

    $ helm upgrade --reuse-values --set foo=bar --set foo=newbar redis ./redis

The --dry-run flag will output all generated chart manifests, including Secrets
which can contain sensitive values. To hide Kubernetes Secrets use the
--hide-secret flag. Please carefully consider how and when these flags are used.`),
			ox.Spec(`[RELEASE] [CHART] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`atomic`, `if set, upgrade process rolls back changes made in case of failed upgrade. The --wait flag will be set automatically if --atomic is used`).
				String(`ca-file`, `verify certificates of HTTPS-enabled servers using this CA bundle`).
				String(`cert-file`, `identify HTTPS client using this SSL certificate file`).
				Bool(`cleanup-on-fail`, `allow deletion of new resources created in this upgrade when upgrade fails`).
				Bool(`create-namespace`, `if --install is set, create the release namespace if not present`).
				Bool(`dependency-update`, `update dependencies if they are missing before installing the chart`).
				String(`description`, `add a custom description`).
				Bool(`devel`, `use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored`).
				Bool(`disable-openapi-validation`, `if set, the upgrade process will not validate rendered templates against the Kubernetes OpenAPI Schema`).
				Map(`dry-run`, `simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.`, ox.Spec(`string[="client"]`)).
				Bool(`enable-dns`, `enable DNS lookups when rendering templates`).
				Bool(`force`, `force resource updates through a replacement strategy`).
				Bool(`hide-notes`, `if set, do not show notes in upgrade output. Does not affect presence in chart metadata`).
				Bool(`hide-secret`, `hide Kubernetes Secrets when also using the --dry-run flag`).
				Int(`history-max`, `limit the maximum number of revisions saved per release. Use 0 for no limit`, ox.Default("10")).
				Bool(`insecure-skip-tls-verify`, `skip tls certificate checks for the chart download`).
				Bool(`install`, `if a release by this name doesn't already exist, run an install`, ox.Short("i")).
				String(`key-file`, `identify HTTPS client using this SSL key file`).
				String(`keyring`, `location of public keys used for verification`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				Map(`labels`, `Labels that would be added to release metadata. Should be separated by comma. Original release labels will be merged with upgrade labels. You can unset label using null.`, ox.Default("[]"), ox.Short("l")).
				Bool(`no-hooks`, `disable pre/post upgrade hooks`).
				String(`output`, `prints the output in the specified format. Allowed values: table, json, yaml`, ox.Spec(`format`), ox.Default("table"), ox.Short("o")).
				Bool(`pass-credentials`, `pass credentials to all domains`).
				String(`password`, `chart repository password where to locate the requested chart`).
				Bool(`plain-http`, `use insecure HTTP connections for the chart download`).
				String(`post-renderer`, `the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path`, ox.Spec(`postRendererString`)).
				String(`post-renderer-args`, `an argument to the post-renderer (can specify multiple)`, ox.Spec(`postRendererArgsSlice`), ox.Default("[]")).
				Bool(`render-subchart-notes`, `if set, render subchart notes along with the parent`).
				String(`repo`, `chart repository url where to locate the requested chart`).
				Bool(`reset-then-reuse-values`, `when upgrading, reset the values to the ones built into the chart, apply the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' or '--reuse-values' is specified, this is ignored`).
				Bool(`reset-values`, `when upgrading, reset the values to the ones built into the chart`).
				Bool(`reuse-values`, `when upgrading, reuse the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' is specified, this is ignored`).
				Array(`set`, `set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Array(`set-file`, `set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)`).
				Array(`set-json`, `set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)`).
				Array(`set-literal`, `set a literal STRING value on the command line`).
				Array(`set-string`, `set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)`).
				Bool(`skip-crds`, `if set, no CRDs will be installed when an upgrade is performed with install flag enabled. By default, CRDs are installed if not already present, when an upgrade is performed with install flag enabled`).
				Bool(`skip-schema-validation`, `if set, disables JSON schema validation`).
				Duration(`timeout`, `time to wait for any individual Kubernetes operation (like Jobs for hooks)`, ox.Default("5m0s")).
				String(`username`, `chart repository username where to locate the requested chart`).
				Slice(`values`, `specify values in a YAML file or a URL (can specify multiple)`, ox.Short("f")).
				Bool(`verify`, `verify the package before using it`).
				Bool(`wait`, `if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout`).
				Bool(`wait-for-jobs`, `if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout`).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage(`verify`, `verify that a chart at the given path has been signed and is valid`),
			ox.Banner(`
Verify that the given chart has a valid provenance file.

Provenance files provide cryptographic verification that a chart has not been
tampered with, and was packaged by a trusted provider.

This command can be used to verify a local chart. Several other commands provide
'--verify' flags that run the same validation. To generate a signed package, use
the 'helm package --sign' command.`),
			ox.Spec(`PATH [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`keyring`, `keyring containing public keys`, ox.Default("$HOME/.gnupg/pubring.gpg")).
				Int(`burst-limit`, `client-side default throttling limit`, ox.Default("100"), ox.Section(0)).
				Bool(`debug`, `enable verbose output`, ox.Section(0)).
				String(`kube-apiserver`, `the address and the port for the Kubernetes API server`, ox.Section(0)).
				Array(`kube-as-group`, `group to impersonate for the operation, this flag can be repeated to specify multiple groups.`, ox.Section(0)).
				String(`kube-as-user`, `username to impersonate for the operation`, ox.Section(0)).
				String(`kube-ca-file`, `the certificate authority file for the Kubernetes API server connection`, ox.Section(0)).
				String(`kube-context`, `name of the kubeconfig context to use`, ox.Section(0)).
				Bool(`kube-insecure-skip-tls-verify`, `if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure`, ox.Section(0)).
				String(`kube-tls-server-name`, `server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used`, ox.Section(0)).
				String(`kube-token`, `bearer token used for authentication`, ox.Section(0)).
				String(`kubeconfig`, `path to the kubeconfig file`, ox.Section(0)).
				String(`namespace`, `namespace scope for this request`, ox.Short("n"), ox.Section(0)).
				Float32(`qps`, `queries per second used when communicating with the Kubernetes API, not including bursting`, ox.Section(0)).
				String(`registry-config`, `path to the registry config file`, ox.Default("$APPCONFIG/registry/config.json"), ox.Section(0)).
				String(`repository-cache`, `path to the directory containing cached repository indexes`, ox.Default("$APPCACHE/repository"), ox.Section(0)).
				String(`repository-config`, `path to the file containing repository names and URLs`, ox.Default("$APPCONFIG/repositories.yaml"), ox.Section(0)),
		),
	)
}
