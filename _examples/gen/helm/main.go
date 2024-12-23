// Command helm is a xo/ox version of `+helm`.
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
		ox.Banner("The Kubernetes package manager\n\nCommon actions for Helm:\n\n- helm search:    search for charts\n- helm pull:      download a chart to your local directory to view\n- helm install:   upload the chart to Kubernetes\n- helm list:      list releases of charts\n\nEnvironment variables:\n\n| Name                               | Description                                                                                                |\n|------------------------------------|------------------------------------------------------------------------------------------------------------|\n| $HELM_CACHE_HOME                   | set an alternative location for storing cached files.                                                      |\n| $HELM_CONFIG_HOME                  | set an alternative location for storing Helm configuration.                                                |\n| $HELM_DATA_HOME                    | set an alternative location for storing Helm data.                                                         |\n| $HELM_DEBUG                        | indicate whether or not Helm is running in Debug mode                                                      |\n| $HELM_DRIVER                       | set the backend storage driver. Values are: configmap, secret, memory, sql.                                |\n| $HELM_DRIVER_SQL_CONNECTION_STRING | set the connection string the SQL storage driver should use.                                               |\n| $HELM_MAX_HISTORY                  | set the maximum number of helm release history.                                                            |\n| $HELM_NAMESPACE                    | set the namespace used for the helm operations.                                                            |\n| $HELM_NO_PLUGINS                   | disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.                                                 |\n| $HELM_PLUGINS                      | set the path to the plugins directory                                                                      |\n| $HELM_REGISTRY_CONFIG              | set the path to the registry config file.                                                                  |\n| $HELM_REPOSITORY_CACHE             | set the path to the repository cache directory                                                             |\n| $HELM_REPOSITORY_CONFIG            | set the path to the repositories file.                                                                     |\n| $KUBECONFIG                        | set an alternative Kubernetes configuration file (default \"~/.kube/config\")                                |\n| $HELM_KUBEAPISERVER                | set the Kubernetes API Server Endpoint for authentication                                                  |\n| $HELM_KUBECAFILE                   | set the Kubernetes certificate authority file.                                                             |\n| $HELM_KUBEASGROUPS                 | set the Groups to use for impersonation using a comma-separated list.                                      |\n| $HELM_KUBEASUSER                   | set the Username to impersonate for the operation.                                                         |\n| $HELM_KUBECONTEXT                  | set the name of the kubeconfig context.                                                                    |\n| $HELM_KUBETOKEN                    | set the Bearer KubeToken used for authentication.                                                          |\n| $HELM_KUBEINSECURE_SKIP_TLS_VERIFY | indicate if the Kubernetes API server's certificate validation should be skipped (insecure)                |\n| $HELM_KUBETLS_SERVER_NAME          | set the server name used to validate the Kubernetes API server certificate                                 |\n| $HELM_BURST_LIMIT                  | set the default burst limit in the case the server contains many CRDs (default 100, -1 to disable)         |\n| $HELM_QPS                          | set the Queries Per Second in cases where a high number of calls exceed the option for higher burst values |\n\nHelm stores cache, configuration, and data based on the following configuration order:\n\n- If a HELM_*_HOME environment variable is set, it will be used\n- Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used\n- When no other location is set a default location will be used based on the operating system\n\nBy default, the default directories depend on the Operating System. The defaults are listed below:\n\n| Operating System | Cache Path                | Configuration Path             | Data Path               |\n|------------------|---------------------------|--------------------------------|-------------------------|\n| Linux            | $HOME/.cache/helm         | $HOME/.config/helm             | $HOME/.local/share/helm |\n| macOS            | $HOME/Library/Caches/helm | $HOME/Library/Preferences/helm | $HOME/Library/helm      |\n| Windows          | %TEMP%\\helm               | %APPDATA%\\helm                 | %APPDATA%\\helm          |"),
		ox.Usage("helm", ""),
		ox.Spec("[command]"),
		ox.Footer("Use \"helm [command] --help\" for more information about a command."),
		ox.Sub(
			ox.Banner("\nThis command creates a chart directory along with the common files and\ndirectories used in a chart.\n\nFor example, 'helm create foo' will create a directory structure that looks\nsomething like this:\n\n    foo/\n    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.\n    ├── Chart.yaml    # Information about your chart\n    ├── values.yaml   # The default values for your templates\n    ├── charts/       # Charts that this chart depends on\n    └── templates/    # The template files\n        └── tests/    # The test files\n\n'helm create' takes a path for an argument. If directories in the given path\ndo not exist, Helm will attempt to create them as it goes. If the given\ndestination exists and there are files in that directory, conflicting files\nwill be overwritten, but other files will be left alone."),
			ox.Usage("create", "create a new chart with the given name"),
			ox.Spec("NAME [flags]"),
			ox.Flags().
				String("starter", "the name or absolute path to Helm starter scaffold", ox.Short("p")).
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nManage the dependencies of a chart.\n\nHelm charts store their dependencies in 'charts/'. For chart developers, it is\noften easier to manage dependencies in 'Chart.yaml' which declares all\ndependencies.\n\nThe dependency commands operate on that file, making it easy to synchronize\nbetween the desired dependencies and the actual dependencies stored in the\n'charts/' directory.\n\nFor example, this Chart.yaml declares two dependencies:\n\n    # Chart.yaml\n    dependencies:\n    - name: nginx\n      version: \"1.2.3\"\n      repository: \"https://example.com/charts\"\n    - name: memcached\n      version: \"3.2.1\"\n      repository: \"https://another.example.com/charts\"\n\n\nThe 'name' should be the name of a chart, where that name must match the name\nin that chart's 'Chart.yaml' file.\n\nThe 'version' field should contain a semantic version or version range.\n\nThe 'repository' URL should point to a Chart Repository. Helm expects that by\nappending '/index.yaml' to the URL, it should be able to retrieve the chart\nrepository's index. Note: 'repository' can be an alias. The alias must start\nwith 'alias:' or '@'.\n\nStarting from 2.2.0, repository can be defined as the path to the directory of\nthe dependency charts stored locally. The path should start with a prefix of\n\"file://\". For example,\n\n    # Chart.yaml\n    dependencies:\n    - name: nginx\n      version: \"1.2.3\"\n      repository: \"file://../dependency_chart/nginx\"\n\nIf the dependency chart is retrieved locally, it is not required to have the\nrepository added to helm by \"helm add repo\". Version matching is also supported\nfor this case."),
			ox.Usage("dependency", "manage a chart's dependencies"),
			ox.Spec("[command]"),
			ox.Aliases("dependency", "dep", "dependencies"),
			ox.Footer("Use \"helm dependency [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nBuild out the charts/ directory from the Chart.lock file.\n\nBuild is used to reconstruct a chart's dependencies to the state specified in\nthe lock file. This will not re-negotiate dependencies, as 'helm dependency update'\ndoes.\n\nIf no lock file is found, 'helm dependency build' will mirror the behavior\nof 'helm dependency update'."),
				ox.Usage("build", "rebuild the charts/ directory based on the Chart.lock file"),
				ox.Spec("CHART [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "keyring containing public keys", ox.Default("$HOME/.gnupg/pubring.gpg")).
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					Bool("skip-refresh", "do not refresh the local repository cache").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the packages against signatures").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nList all of the dependencies declared in a chart.\n\nThis can take chart archives and chart directories as input. It will not alter\nthe contents of a chart.\n\nThis will produce an error if the chart cannot be loaded."),
				ox.Usage("list", "list the dependencies for the given chart"),
				ox.Spec("CHART [flags]"),
				ox.Aliases("list", "ls"),
				ox.Flags().
					Uint("max-col-width", "maximum column width for output table", ox.Default("80")).
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nUpdate the on-disk dependencies to mirror Chart.yaml.\n\nThis command verifies that the required charts, as expressed in 'Chart.yaml',\nare present in 'charts/' and are at an acceptable version. It will pull down\nthe latest charts that satisfy the dependencies, and clean up old dependencies.\n\nOn successful update, this will generate a lock file that can be used to\nrebuild the dependencies to an exact version.\n\nDependencies are not required to be represented in 'Chart.yaml'. For that\nreason, an update command will not remove charts unless they are (a) present\nin the Chart.yaml file, but (b) at the wrong version."),
				ox.Usage("update", "update charts/ based on the contents of Chart.yaml"),
				ox.Spec("CHART [flags]"),
				ox.Aliases("update", "up"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "keyring containing public keys", ox.Default("$HOME/.gnupg/pubring.gpg")).
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					Bool("skip-refresh", "do not refresh the local repository cache").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the packages against signatures").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nEnv prints out all the environment information in use by Helm."),
			ox.Usage("env", "helm client environment information"),
			ox.Spec("[flags]"),
			ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command consists of multiple subcommands which can be used to\nget extended information about the release, including:\n\n- The values used to generate the release\n- The generated manifest file\n- The notes provided by the chart of the release\n- The hooks associated with the release\n- The metadata of the release"),
			ox.Usage("get", "download extended information of a named release"),
			ox.Spec("[command]"),
			ox.Footer("Use \"helm get [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nThis command prints a human readable collection of information about the\nnotes, hooks, supplied values, and generated manifest file of the given release."),
				ox.Usage("all", "download all information for a named release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					Int("revision", "get the named release with revision").
					String("template", "go template for formatting the output, eg: {{.Release.Name}}").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command downloads hooks for a given release.\n\nHooks are formatted in YAML and separated by the YAML '---\\n' separator."),
				ox.Usage("hooks", "download all hooks for a named release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					Int("revision", "get the named release with revision").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command fetches the generated manifest for a given release.\n\nA manifest is a YAML-encoded representation of the Kubernetes resources that\nwere generated from this release's chart(s). If a chart is dependent on other\ncharts, those resources will also be included in the manifest."),
				ox.Usage("manifest", "download the manifest for a named release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					Int("revision", "get the named release with revision").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("This command fetches metadata for a given release"),
				ox.Usage("metadata", "This command fetches metadata for a given release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
					Int("revision", "specify release revision").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command shows notes provided by the chart of a named release."),
				ox.Usage("notes", "download the notes for a named release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					Int("revision", "get the named release with revision").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command downloads a values file for a given release."),
				ox.Usage("values", "download the values file for a named release"),
				ox.Spec("RELEASE_NAME [flags]"),
				ox.Flags().
					Bool("all", "dump all (computed) values", ox.Short("a")).
					String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
					Int("revision", "get the named release with revision").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nHistory prints historical revisions for a given release.\n\nA default maximum of 256 revisions will be returned. Setting '--max'\nconfigures the maximum length of the revision list returned.\n\nThe historical release set is printed as a formatted table, e.g:\n\n    $ helm history angry-bird\n    REVISION    UPDATED                     STATUS          CHART             APP VERSION     DESCRIPTION\n    1           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Initial install\n    2           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Upgraded successfully\n    3           Mon Oct 3 10:15:13 2016     superseded      alpine-0.1.0      1.0             Rolled back to 2\n    4           Mon Oct 3 10:15:13 2016     deployed        alpine-0.1.0      1.0             Upgraded successfully"),
			ox.Usage("history", "fetch release history"),
			ox.Spec("RELEASE_NAME [flags]"),
			ox.Aliases("history", "hist"),
			ox.Flags().
				Int("max", "maximum number of revision to include in history", ox.Default("256")).
				String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command installs a chart archive.\n\nThe install argument must be a chart reference, a path to a packaged chart,\na path to an unpacked chart directory or a URL.\n\nTo override values in a chart, use either the '--values' flag and pass in a file\nor use the '--set' flag and pass configuration from the command line, to force\na string value use '--set-string'. You can use '--set-file' to set individual\nvalues from a file when the value itself is too long for the command line\nor is dynamically generated. You can also use '--set-json' to set json values\n(scalars/objects/arrays) from the command line.\n\n    $ helm install -f myvalues.yaml myredis ./redis\n\nor\n\n    $ helm install --set name=prod myredis ./redis\n\nor\n\n    $ helm install --set-string long_int=1234567890 myredis ./redis\n\nor\n\n    $ helm install --set-file my_script=dothings.sh myredis ./redis\n\nor\n\n    $ helm install --set-json 'master.sidecars=[{\"name\":\"sidecar\",\"image\":\"myImage\",\"imagePullPolicy\":\"Always\",\"ports\":[{\"name\":\"portname\",\"containerPort\":1234}]}]' myredis ./redis\n\n\nYou can specify the '--values'/'-f' flag multiple times. The priority will be given to the\nlast (right-most) file specified. For example, if both myvalues.yaml and override.yaml\ncontained a key called 'Test', the value set in override.yaml would take precedence:\n\n    $ helm install -f myvalues.yaml -f override.yaml  myredis ./redis\n\nYou can specify the '--set' flag multiple times. The priority will be given to the\nlast (right-most) set specified. For example, if both 'bar' and 'newbar' values are\nset for a key called 'foo', the 'newbar' value would take precedence:\n\n    $ helm install --set foo=bar --set foo=newbar  myredis ./redis\n\nSimilarly, in the following example 'foo' is set to '[\"four\"]':\n\n    $ helm install --set-json='foo=[\"one\", \"two\", \"three\"]' --set-json='foo=[\"four\"]' myredis ./redis\n\nAnd in the following example, 'foo' is set to '{\"key1\":\"value1\",\"key2\":\"bar\"}':\n\n    $ helm install --set-json='foo={\"key1\":\"value1\",\"key2\":\"value2\"}' --set-json='foo.key2=\"bar\"' myredis ./redis\n\nTo check the generated manifests of a release without installing the chart,\nthe --debug and --dry-run flags can be combined.\n\nThe --dry-run flag will output all generated chart manifests, including Secrets\nwhich can contain sensitive values. To hide Kubernetes Secrets use the\n--hide-secret flag. Please carefully consider how and when these flags are used.\n\nIf --verify is set, the chart MUST have a provenance file, and the provenance\nfile MUST pass all verification steps.\n\nThere are six different ways you can express the chart you want to install:\n\n1. By chart reference: helm install mymaria example/mariadb\n2. By path to a packaged chart: helm install mynginx ./nginx-1.2.3.tgz\n3. By path to an unpacked chart directory: helm install mynginx ./nginx\n4. By absolute URL: helm install mynginx https://example.com/charts/nginx-1.2.3.tgz\n5. By chart reference and repo url: helm install --repo https://example.com/charts/ mynginx nginx\n6. By OCI registries: helm install mynginx --version 1.2.3 oci://example.com/charts/nginx\n\nCHART REFERENCES\n\nA chart reference is a convenient way of referencing a chart in a chart repository.\n\nWhen you use a chart reference with a repo prefix ('example/mariadb'), Helm will look in the local\nconfiguration for a chart repository named 'example', and will then look for a\nchart in that repository whose name is 'mariadb'. It will install the latest stable version of that chart\nuntil you specify '--devel' flag to also include development version (alpha, beta, and release candidate releases), or\nsupply a version number with the '--version' flag.\n\nTo see the list of chart repositories, use 'helm repo list'. To search for\ncharts in a repository, use 'helm search'."),
			ox.Usage("install", "install a chart"),
			ox.Spec("[NAME] [CHART] [flags]"),
			ox.Flags().
				Bool("atomic", "if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used").
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify HTTPS client using this SSL certificate file").
				Bool("create-namespace", "create the release namespace if not present").
				Bool("dependency-update", "update dependencies if they are missing before installing the chart").
				String("description", "add a custom description").
				Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
				Bool("disable-openapi-validation", "if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema").
				Map("dry-run", "simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.", ox.Spec("string[=\"client\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Bool("enable-dns", "enable DNS lookups when rendering templates").
				Bool("force", "force resource updates through a replacement strategy").
				Bool("generate-name", "generate the name (and omit the NAME parameter)", ox.Short("g")).
				Bool("hide-notes", "if set, do not show notes in install output. Does not affect presence in chart metadata").
				Bool("hide-secret", "hide Kubernetes Secrets when also using the --dry-run flag").
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
				String("key-file", "identify HTTPS client using this SSL key file").
				String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
				Map("labels", "Labels that would be added to release metadata. Should be divided by comma.", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("l")).
				String("name-template", "specify template used to name the release").
				Bool("no-hooks", "prevent hooks from running during install").
				String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
				Bool("pass-credentials", "pass credentials to all domains").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart download").
				String("post-renderer", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path", ox.Spec("postRendererString")).
				String("post-renderer-args", "an argument to the post-renderer (can specify multiple)", ox.Spec("postRendererArgsSlice"), ox.Default("[]")).
				Bool("render-subchart-notes", "if set, render subchart notes along with the parent").
				Bool("replace", "reuse the given name, only if that name is a deleted release which remains in the history. This is unsafe in production").
				String("repo", "chart repository url where to locate the requested chart").
				Slice("set", "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Slice("set-file", "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)", ox.Elem(ox.StringT)).
				Slice("set-json", "set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)", ox.Elem(ox.StringT)).
				Slice("set-literal", "set a literal STRING value on the command line", ox.Elem(ox.StringT)).
				Slice("set-string", "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Bool("skip-crds", "if set, no CRDs will be installed. By default, CRDs are installed if not already present").
				Bool("skip-schema-validation", "if set, disables JSON schema validation").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				String("username", "chart repository username where to locate the requested chart").
				Slice("values", "specify values in a YAML file or a URL (can specify multiple)", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("verify", "verify the package before using it").
				Bool("wait", "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout").
				Bool("wait-for-jobs", "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command takes a path to a chart and runs a series of tests to verify that\nthe chart is well-formed.\n\nIf the linter encounters things that will cause the chart to fail installation,\nit will emit [ERROR] messages. If it encounters issues that break with convention\nor recommendation, it will emit [WARNING] messages."),
			ox.Usage("lint", "examine a chart for possible issues"),
			ox.Spec("PATH [flags]"),
			ox.Flags().
				String("kube-version", "Kubernetes version used for capabilities and deprecation checks").
				Bool("quiet", "print only warnings and errors").
				Slice("set", "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Slice("set-file", "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)", ox.Elem(ox.StringT)).
				Slice("set-json", "set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)", ox.Elem(ox.StringT)).
				Slice("set-literal", "set a literal STRING value on the command line", ox.Elem(ox.StringT)).
				Slice("set-string", "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Bool("skip-schema-validation", "if set, disables JSON schema validation").
				Bool("strict", "fail on lint warnings").
				Slice("values", "specify values in a YAML file or a URL (can specify multiple)", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("with-subcharts", "lint dependent charts").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command lists all of the releases for a specified namespace (uses current namespace context if namespace not specified).\n\nBy default, it lists only releases that are deployed or failed. Flags like\n'--uninstalled' and '--all' will alter this behavior. Such flags can be combined:\n'--uninstalled --failed'.\n\nBy default, items are sorted alphabetically. Use the '-d' flag to sort by\nrelease date.\n\nIf the --filter flag is provided, it will be treated as a filter. Filters are\nregular expressions (Perl compatible) that are applied to the list of releases.\nOnly items that match the filter will be returned.\n\n    $ helm list --filter 'ara[a-z]+'\n    NAME                UPDATED                                  CHART\n    maudlin-arachnid    2020-06-18 14:17:46.125134977 +0000 UTC  alpine-0.1.0\n\nIf no results are found, 'helm list' will exit 0, but with no output (or in\nthe case of no '-q' flag, only headers).\n\nBy default, up to 256 items may be returned. To limit this, use the '--max' flag.\nSetting '--max' to 0 will not return all results. Rather, it will return the\nserver's default, which may be much higher than 256. Pairing the '--max'\nflag with the '--offset' flag allows you to page through results."),
			ox.Usage("list", "list releases"),
			ox.Spec("[flags]"),
			ox.Aliases("list", "ls"),
			ox.Flags().
				Bool("all", "show all releases without any filter applied", ox.Short("a")).
				Bool("all-namespaces", "list releases across all namespaces", ox.Short("A")).
				Bool("date", "sort by release date", ox.Short("d")).
				Bool("deployed", "show deployed releases. If no other is specified, this will be automatically enabled").
				Bool("failed", "show failed releases").
				String("filter", "a regular expression (Perl compatible). Any releases that match the expression will be included in the results", ox.Short("f")).
				Int("max", "maximum number of releases to fetch", ox.Default("256"), ox.Short("m")).
				Bool("no-headers", "don't print headers when using the default output format").
				Int("offset", "next release index in the list, used to offset from start value").
				String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
				Bool("pending", "show pending releases").
				Bool("reverse", "reverse the sort order", ox.Short("r")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Works only for secret(default) and configmap storage backends.", ox.Short("l")).
				Bool("short", "output short (quiet) listing format", ox.Short("q")).
				Bool("superseded", "show superseded releases").
				String("time-format", "format time using golang time formatter. Example: --time-format \"2006-01-02 15:04:05Z0700\"").
				Bool("uninstalled", "show uninstalled releases (if 'helm uninstall --keep-history' was used)").
				Bool("uninstalling", "show releases that are currently being uninstalled").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command packages a chart into a versioned chart archive file. If a path\nis given, this will look at that path for a chart (which must contain a\nChart.yaml file) and then package that directory.\n\nVersioned chart archives are used by Helm package repositories.\n\nTo sign a chart, use the '--sign' flag. In most cases, you should also\nprovide '--keyring path/to/secret/keys' and '--key keyname'.\n\n  $ helm package --sign ./mychart --key mykey --keyring ~/.gnupg/secring.gpg\n\nIf '--keyring' is not specified, Helm usually defaults to the public keyring\nunless your environment is otherwise configured."),
			ox.Usage("package", "package a chart directory into a chart archive"),
			ox.Spec("[CHART_PATH] [...] [flags]"),
			ox.Flags().
				String("app-version", "set the appVersion on the chart to this version").
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify HTTPS client using this SSL certificate file").
				Bool("dependency-update", "update dependencies from \"Chart.yaml\" to dir \"charts/\" before packaging", ox.Short("u")).
				String("destination", "location to write the chart.", ox.Default("."), ox.Short("d")).
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
				String("key", "name of the key to use when signing. Used if --sign is true").
				String("key-file", "identify HTTPS client using this SSL key file").
				String("keyring", "location of a public keyring", ox.Default("$HOME/.gnupg/pubring.gpg")).
				String("passphrase-file", "location of a file which contains the passphrase for the signing key. Use \"-\" in order to read from stdin.").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart download").
				Bool("sign", "use a PGP private key to sign this package").
				String("username", "chart repository username where to locate the requested chart").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nManage client-side Helm plugins."),
			ox.Usage("plugin", "install, list, or uninstall Helm plugins"),
			ox.Spec("[command]"),
			ox.Footer("Use \"helm plugin [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nThis command allows you to install a plugin from a url to a VCS repo or a local path."),
				ox.Usage("install", "install a Helm plugin"),
				ox.Spec("[options] <path|url> [flags]"),
				ox.Aliases("install", "add"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("list installed Helm plugins"),
				ox.Usage("list", "list installed Helm plugins"),
				ox.Spec("[flags]"),
				ox.Aliases("list", "ls"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("uninstall one or more Helm plugins"),
				ox.Usage("uninstall", "uninstall one or more Helm plugins"),
				ox.Spec("<plugin>... [flags]"),
				ox.Aliases("uninstall", "rm", "remove"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("update one or more Helm plugins"),
				ox.Usage("update", "update one or more Helm plugins"),
				ox.Spec("<plugin>... [flags]"),
				ox.Aliases("update", "up"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nRetrieve a package from a package repository, and download it locally.\n\nThis is useful for fetching packages to inspect, modify, or repackage. It can\nalso be used to perform cryptographic verification of a chart without installing\nthe chart.\n\nThere are options for unpacking the chart after download. This will create a\ndirectory for the chart and uncompress into that directory.\n\nIf the --verify flag is specified, the requested chart MUST have a provenance\nfile, and MUST pass the verification process. Failure in any part of this will\nresult in an error, and the chart will not be saved locally."),
			ox.Usage("pull", "download a chart from a repository and (optionally) unpack it in local directory"),
			ox.Spec("[chart URL | repo/chartname] [...] [flags]"),
			ox.Aliases("pull", "fetch"),
			ox.Flags().
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify HTTPS client using this SSL certificate file").
				String("destination", "location to write the chart. If this and untardir are specified, untardir is appended to this", ox.Default("."), ox.Short("d")).
				Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored.").
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
				String("key-file", "identify HTTPS client using this SSL key file").
				String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
				Bool("pass-credentials", "pass credentials to all domains").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart download").
				Bool("prov", "fetch the provenance file, but don't perform verification").
				String("repo", "chart repository url where to locate the requested chart").
				Bool("untar", "if set to true, will untar the chart after downloading it").
				String("untardir", "if untar is specified, this flag specifies the name of the directory into which the chart is expanded", ox.Default(".")).
				String("username", "chart repository username where to locate the requested chart").
				Bool("verify", "verify the package before using it").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nUpload a chart to a registry.\n\nIf the chart has an associated provenance file,\nit will also be uploaded."),
			ox.Usage("push", "push a chart to remote"),
			ox.Spec("[chart] [remote] [flags]"),
			ox.Flags().
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify registry client using this SSL certificate file").
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart upload").
				String("key-file", "identify registry client using this SSL key file").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart upload").
				String("username", "chart repository username where to locate the requested chart").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command consists of multiple subcommands to interact with registries."),
			ox.Usage("registry", "login to or logout from a registry"),
			ox.Spec("[command]"),
			ox.Footer("Use \"helm registry [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nAuthenticate to a remote registry."),
				ox.Usage("login", "login to a registry"),
				ox.Spec("[host] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify registry client using this SSL certificate file").
					Bool("insecure", "allow connections to TLS registry without certs").
					String("key-file", "identify registry client using this SSL key file").
					String("password", "registry password or identity token", ox.Short("p")).
					Bool("password-stdin", "read password or identity token from stdin").
					String("username", "registry username", ox.Short("u")).
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nRemove credentials stored for a remote registry."),
				ox.Usage("logout", "logout from a registry"),
				ox.Spec("[host] [flags]"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command consists of multiple subcommands to interact with chart repositories.\n\nIt can be used to add, remove, list, and index chart repositories."),
			ox.Usage("repo", "add, list, remove, update, and index chart repositories"),
			ox.Spec("[command]"),
			ox.Footer("Use \"helm repo [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("add a chart repository"),
				ox.Usage("add", "add a chart repository"),
				ox.Spec("[NAME] [URL] [flags]"),
				ox.Flags().
					Bool("allow-deprecated-repos", "by default, this command will not allow adding official repos that have been permanently deleted. This disables that behavior").
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("force-update", "replace (overwrite) the repo if it already exists").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the repository").
					String("key-file", "identify HTTPS client using this SSL key file").
					Bool("no-update", "Ignored. Formerly, it would disabled forced updates. It is deprecated by force-update.").
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password").
					Bool("password-stdin", "read chart repository password from stdin").
					String("username", "chart repository username").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nRead the current directory, generate an index file based on the charts found\nand write the result to 'index.yaml' in the current directory.\n\nThis tool is used for creating an 'index.yaml' file for a chart repository. To\nset an absolute URL to the charts, use '--url' flag.\n\nTo merge the generated index with an existing index file, use the '--merge'\nflag. In this case, the charts found in the current directory will be merged\ninto the index passed in with --merge, with local charts taking priority over\nexisting charts."),
				ox.Usage("index", "generate an index file given a directory containing packaged charts"),
				ox.Spec("[DIR] [flags]"),
				ox.Flags().
					Bool("json", "output in JSON format").
					String("merge", "merge the generated index into the given index").
					String("url", "url of chart repository").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("list chart repositories"),
				ox.Usage("list", "list chart repositories"),
				ox.Spec("[flags]"),
				ox.Aliases("list", "ls"),
				ox.Flags().
					String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("remove one or more chart repositories"),
				ox.Usage("remove", "remove one or more chart repositories"),
				ox.Spec("[REPO1 [REPO2 ...]] [flags]"),
				ox.Aliases("remove", "rm"),
				ox.Flags().
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nUpdate gets the latest information about charts from the respective chart repositories.\nInformation is cached locally, where it is used by commands like 'helm search'.\n\nYou can optionally specify a list of repositories you want to update.\n\t$ helm repo update <repo_name> ...\nTo update all the repositories, use 'helm repo update'."),
				ox.Usage("update", "update information of available charts locally from chart repositories"),
				ox.Spec("[REPO1 [REPO2 ...]] [flags]"),
				ox.Aliases("update", "up"),
				ox.Flags().
					Bool("fail-on-repo-update-fail", "update fails if any of the repository updates fail").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command rolls back a release to a previous revision.\n\nThe first argument of the rollback command is the name of a release, and the\nsecond is a revision (version) number. If this argument is omitted or set to\n0, it will roll back to the previous release.\n\nTo see revision numbers, run 'helm history RELEASE'."),
			ox.Usage("rollback", "roll back a release to a previous revision"),
			ox.Spec("<RELEASE> [REVISION] [flags]"),
			ox.Flags().
				Bool("cleanup-on-fail", "allow deletion of new resources created in this rollback when rollback fails").
				Bool("dry-run", "simulate a rollback").
				Bool("force", "force resource update through delete/recreate if needed").
				Int("history-max", "limit the maximum number of revisions saved per release. Use 0 for no limit", ox.Default("10")).
				Bool("no-hooks", "prevent hooks from running during rollback").
				Bool("recreate-pods", "performs pods restart for the resource if applicable").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				Bool("wait", "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout").
				Bool("wait-for-jobs", "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nSearch provides the ability to search for Helm charts in the various places\nthey can be stored including the Artifact Hub and repositories you have added.\nUse search subcommands to search different locations for charts."),
			ox.Usage("search", "search for a keyword in charts"),
			ox.Spec("[command]"),
			ox.Footer("Use \"helm search [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nSearch for Helm charts in the Artifact Hub or your own hub instance.\n\nArtifact Hub is a web-based application that enables finding, installing, and\npublishing packages and configurations for CNCF projects, including publicly\navailable distributed charts Helm charts. It is a Cloud Native Computing\nFoundation sandbox project. You can browse the hub at https://artifacthub.io/\n\nThe [KEYWORD] argument accepts either a keyword string, or quoted string of rich\nquery options. For rich query options documentation, see\nhttps://artifacthub.github.io/hub/api/?urls.primaryName=Monocular%20compatible%20search%20API#/Monocular/get_api_chartsvc_v1_charts_search\n\nPrevious versions of Helm used an instance of Monocular as the default\n'endpoint', so for backwards compatibility Artifact Hub is compatible with the\nMonocular search API. Similarly, when setting the 'endpoint' flag, the specified\nendpoint must also be implement a Monocular compatible search API endpoint.\nNote that when specifying a Monocular instance as the 'endpoint', rich queries\nare not supported. For API details, see https://github.com/helm/monocular"),
				ox.Usage("hub", "search for charts in the Artifact Hub or your own hub instance"),
				ox.Spec("[KEYWORD] [flags]"),
				ox.Flags().
					String("endpoint", "Hub instance to query for charts", ox.Default("https://hub.$APPNAME.sh")).
					Bool("fail-on-no-result", "search fails if no results are found").
					Bool("list-repo-url", "print charts repository URL").
					Uint("max-col-width", "maximum column width for output table", ox.Default("50")).
					String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nSearch reads through all of the repositories configured on the system, and\nlooks for matches. Search of these repositories uses the metadata stored on\nthe system.\n\nIt will display the latest stable versions of the charts found. If you\nspecify the --devel flag, the output will include pre-release versions.\nIf you want to search using a version constraint, use --version.\n\nExamples:\n\n    # Search for stable release versions matching the keyword \"nginx\"\n    $ helm search repo nginx\n\n    # Search for release versions matching the keyword \"nginx\", including pre-release versions\n    $ helm search repo nginx --devel\n\n    # Search for the latest stable release for nginx-ingress with a major version of 1\n    $ helm search repo nginx-ingress --version ^1.0.0\n\nRepositories are managed with 'helm repo' commands."),
				ox.Usage("repo", "search repositories for a keyword in charts"),
				ox.Spec("[keyword] [flags]"),
				ox.Flags().
					Bool("devel", "use development versions (alpha, beta, and release candidate releases), too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("fail-on-no-result", "search fails if no results are found").
					Uint("max-col-width", "maximum column width for output table", ox.Default("50")).
					String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
					Bool("regexp", "use regular expressions for searching repositories you have added", ox.Short("r")).
					Bool("versions", "show the long listing, with each version of each chart on its own line, for repositories you have added", ox.Short("l")).
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command consists of multiple subcommands to display information about a chart"),
			ox.Usage("show", "show information of a chart"),
			ox.Spec("[command]"),
			ox.Aliases("show", "inspect"),
			ox.Footer("Use \"helm show [command] --help\" for more information about a command."),
			ox.Sub(
				ox.Banner("\nThis command inspects a chart (directory, file, or URL) and displays all its content\n(values.yaml, Chart.yaml, README)"),
				ox.Usage("all", "show all information of the chart"),
				ox.Spec("[CHART] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					String("repo", "chart repository url where to locate the requested chart").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the package before using it").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command inspects a chart (directory, file, or URL) and displays the contents\nof the Chart.yaml file"),
				ox.Usage("chart", "show the chart's definition"),
				ox.Spec("[CHART] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					String("repo", "chart repository url where to locate the requested chart").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the package before using it").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command inspects a chart (directory, file, or URL) and displays the contents\nof the CustomResourceDefinition files"),
				ox.Usage("crds", "show the chart's CRDs"),
				ox.Spec("[CHART] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					String("repo", "chart repository url where to locate the requested chart").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the package before using it").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command inspects a chart (directory, file, or URL) and displays the contents\nof the README file"),
				ox.Usage("readme", "show the chart's README"),
				ox.Spec("[CHART] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					String("repo", "chart repository url where to locate the requested chart").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the package before using it").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Sub(
				ox.Banner("\nThis command inspects a chart (directory, file, or URL) and displays the contents\nof the values.yaml file"),
				ox.Usage("values", "show the chart's values"),
				ox.Spec("[CHART] [flags]"),
				ox.Flags().
					String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
					String("cert-file", "identify HTTPS client using this SSL certificate file").
					Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
					Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
					String("jsonpath", "supply a JSONPath expression to filter the output").
					String("key-file", "identify HTTPS client using this SSL key file").
					String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
					Bool("pass-credentials", "pass credentials to all domains").
					String("password", "chart repository password where to locate the requested chart").
					Bool("plain-http", "use insecure HTTP connections for the chart download").
					String("repo", "chart repository url where to locate the requested chart").
					String("username", "chart repository username where to locate the requested chart").
					Bool("verify", "verify the package before using it").
					Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
					Bool("debug", "enable verbose output").
					String("kube-apiserver", "the address and the port for the Kubernetes API server").
					Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
					String("kube-as-user", "username to impersonate for the operation").
					String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
					String("kube-context", "name of the kubeconfig context to use").
					Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
					String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
					String("kube-token", "bearer token used for authentication").
					String("kubeconfig", "path to the kubeconfig file").
					String("namespace", "namespace scope for this request", ox.Short("n")).
					Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
					String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
					String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
					String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
			), ox.Flags().
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command shows the status of a named release.\nThe status consists of:\n- last deployment time\n- k8s namespace in which the release lives\n- state of the release (can be: unknown, deployed, uninstalled, superseded, failed, uninstalling, pending-install, pending-upgrade or pending-rollback)\n- revision of the release\n- description of the release (can be completion message or error message)\n- list of resources that this release consists of\n- details on last test suite run, if applicable\n- additional notes provided by the chart"),
			ox.Usage("status", "display the status of the named release"),
			ox.Spec("RELEASE_NAME [flags]"),
			ox.Flags().
				String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
				Int("revision", "if set, display the status of the named release with revision").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nRender chart templates locally and display the output.\n\nAny values that would normally be looked up or retrieved in-cluster will be\nfaked locally. Additionally, none of the server-side testing of chart validity\n(e.g. whether an API is supported) is done."),
			ox.Usage("template", "locally render templates"),
			ox.Spec("[NAME] [CHART] [flags]"),
			ox.Flags().
				Slice("api-versions", "Kubernetes api versions used for Capabilities.APIVersions", ox.Elem(ox.StringT), ox.Short("a")).
				Bool("atomic", "if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used").
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify HTTPS client using this SSL certificate file").
				Bool("create-namespace", "create the release namespace if not present").
				Bool("dependency-update", "update dependencies if they are missing before installing the chart").
				String("description", "add a custom description").
				Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
				Bool("disable-openapi-validation", "if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema").
				Map("dry-run", "simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.", ox.Spec("string[=\"client\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Bool("enable-dns", "enable DNS lookups when rendering templates").
				Bool("force", "force resource updates through a replacement strategy").
				Bool("generate-name", "generate the name (and omit the NAME parameter)", ox.Short("g")).
				Bool("hide-notes", "if set, do not show notes in install output. Does not affect presence in chart metadata").
				Bool("include-crds", "include CRDs in the templated output").
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
				Bool("is-upgrade", "set .Release.IsUpgrade instead of .Release.IsInstall").
				String("key-file", "identify HTTPS client using this SSL key file").
				String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
				String("kube-version", "Kubernetes version used for Capabilities.KubeVersion").
				Map("labels", "Labels that would be added to release metadata. Should be divided by comma.", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("l")).
				String("name-template", "specify template used to name the release").
				Bool("no-hooks", "prevent hooks from running during install").
				String("output-dir", "writes the executed templates to files in output-dir instead of stdout").
				Bool("pass-credentials", "pass credentials to all domains").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart download").
				String("post-renderer", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path", ox.Spec("postRendererString")).
				String("post-renderer-args", "an argument to the post-renderer (can specify multiple)", ox.Spec("postRendererArgsSlice"), ox.Default("[]")).
				Bool("release-name", "use release name in the output-dir path.").
				Bool("render-subchart-notes", "if set, render subchart notes along with the parent").
				Bool("replace", "reuse the given name, only if that name is a deleted release which remains in the history. This is unsafe in production").
				String("repo", "chart repository url where to locate the requested chart").
				Slice("set", "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Slice("set-file", "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)", ox.Elem(ox.StringT)).
				Slice("set-json", "set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)", ox.Elem(ox.StringT)).
				Slice("set-literal", "set a literal STRING value on the command line", ox.Elem(ox.StringT)).
				Slice("set-string", "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Slice("show-only", "only show manifests rendered from the given templates", ox.Elem(ox.StringT), ox.Short("s")).
				Bool("skip-crds", "if set, no CRDs will be installed. By default, CRDs are installed if not already present").
				Bool("skip-schema-validation", "if set, disables JSON schema validation").
				Bool("skip-tests", "skip tests from templated output").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				String("username", "chart repository username where to locate the requested chart").
				Bool("validate", "validate your manifests against the Kubernetes cluster you are currently pointing at. This is the same validation performed on an install").
				Slice("values", "specify values in a YAML file or a URL (can specify multiple)", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("verify", "verify the package before using it").
				Bool("wait", "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout").
				Bool("wait-for-jobs", "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThe test command runs the tests for a release.\n\nThe argument this command takes is the name of a deployed release.\nThe tests to be run are defined in the chart that was installed."),
			ox.Usage("test", "run tests for a release"),
			ox.Spec("[RELEASE] [flags]"),
			ox.Flags().
				Slice("filter", "specify tests by attribute (currently \"name\") using attribute=value syntax or '!attribute=value' to exclude a test (can specify multiple or separate values with commas: name=test1,name=test2)", ox.Elem(ox.StringT)).
				Bool("hide-notes", "if set, do not show notes in test output. Does not affect presence in chart metadata").
				Bool("logs", "dump the logs from test pods (this runs after all tests are complete, but before any cleanup)").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command takes a release name and uninstalls the release.\n\nIt removes all of the resources associated with the last release of the chart\nas well as the release history, freeing it up for future use.\n\nUse the '--dry-run' flag to see which releases will be uninstalled without actually\nuninstalling them."),
			ox.Usage("uninstall", "uninstall a release"),
			ox.Spec("RELEASE_NAME [...] [flags]"),
			ox.Aliases("uninstall", "del", "delete", "un"),
			ox.Flags().
				String("cascade", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents. Defaults to background.", ox.Default("background")).
				String("description", "add a custom description").
				Bool("dry-run", "simulate a uninstall").
				Bool("ignore-not-found", "Treat \"release not found\" as a successful uninstall").
				Bool("keep-history", "remove all associated resources and mark the release as deleted, but retain the release history").
				Bool("no-hooks", "prevent hooks from running during uninstallation").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				Bool("wait", "if set, will wait until all the resources are deleted before returning. It will wait for as long as --timeout").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nThis command upgrades a release to a new version of a chart.\n\nThe upgrade arguments must be a release and chart. The chart\nargument can be either: a chart reference('example/mariadb'), a path to a chart directory,\na packaged chart, or a fully qualified URL. For chart references, the latest\nversion will be specified unless the '--version' flag is set.\n\nTo override values in a chart, use either the '--values' flag and pass in a file\nor use the '--set' flag and pass configuration from the command line, to force string\nvalues, use '--set-string'. You can use '--set-file' to set individual\nvalues from a file when the value itself is too long for the command line\nor is dynamically generated. You can also use '--set-json' to set json values\n(scalars/objects/arrays) from the command line.\n\nYou can specify the '--values'/'-f' flag multiple times. The priority will be given to the\nlast (right-most) file specified. For example, if both myvalues.yaml and override.yaml\ncontained a key called 'Test', the value set in override.yaml would take precedence:\n\n    $ helm upgrade -f myvalues.yaml -f override.yaml redis ./redis\n\nYou can specify the '--set' flag multiple times. The priority will be given to the\nlast (right-most) set specified. For example, if both 'bar' and 'newbar' values are\nset for a key called 'foo', the 'newbar' value would take precedence:\n\n    $ helm upgrade --set foo=bar --set foo=newbar redis ./redis\n\nYou can update the values for an existing release with this command as well via the\n'--reuse-values' flag. The 'RELEASE' and 'CHART' arguments should be set to the original\nparameters, and existing values will be merged with any values set via '--values'/'-f'\nor '--set' flags. Priority is given to new values.\n\n    $ helm upgrade --reuse-values --set foo=bar --set foo=newbar redis ./redis\n\nThe --dry-run flag will output all generated chart manifests, including Secrets\nwhich can contain sensitive values. To hide Kubernetes Secrets use the\n--hide-secret flag. Please carefully consider how and when these flags are used."),
			ox.Usage("upgrade", "upgrade a release"),
			ox.Spec("[RELEASE] [CHART] [flags]"),
			ox.Flags().
				Bool("atomic", "if set, upgrade process rolls back changes made in case of failed upgrade. The --wait flag will be set automatically if --atomic is used").
				String("ca-file", "verify certificates of HTTPS-enabled servers using this CA bundle").
				String("cert-file", "identify HTTPS client using this SSL certificate file").
				Bool("cleanup-on-fail", "allow deletion of new resources created in this upgrade when upgrade fails").
				Bool("create-namespace", "if --install is set, create the release namespace if not present").
				Bool("dependency-update", "update dependencies if they are missing before installing the chart").
				String("description", "add a custom description").
				Bool("devel", "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored").
				Bool("disable-openapi-validation", "if set, the upgrade process will not validate rendered templates against the Kubernetes OpenAPI Schema").
				Map("dry-run", "simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.", ox.Spec("string[=\"client\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Bool("enable-dns", "enable DNS lookups when rendering templates").
				Bool("force", "force resource updates through a replacement strategy").
				Bool("hide-notes", "if set, do not show notes in upgrade output. Does not affect presence in chart metadata").
				Bool("hide-secret", "hide Kubernetes Secrets when also using the --dry-run flag").
				Int("history-max", "limit the maximum number of revisions saved per release. Use 0 for no limit", ox.Default("10")).
				Bool("insecure-skip-tls-verify", "skip tls certificate checks for the chart download").
				Bool("install", "if a release by this name doesn't already exist, run an install", ox.Short("i")).
				String("key-file", "identify HTTPS client using this SSL key file").
				String("keyring", "location of public keys used for verification", ox.Default("$HOME/.gnupg/pubring.gpg")).
				Map("labels", "Labels that would be added to release metadata. Should be separated by comma. Original release labels will be merged with upgrade labels. You can unset label using null.", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("l")).
				Bool("no-hooks", "disable pre/post upgrade hooks").
				String("output", "prints the output in the specified format. Allowed values: table, json, yaml", ox.Spec("format"), ox.Default("table"), ox.Short("o")).
				Bool("pass-credentials", "pass credentials to all domains").
				String("password", "chart repository password where to locate the requested chart").
				Bool("plain-http", "use insecure HTTP connections for the chart download").
				String("post-renderer", "the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path", ox.Spec("postRendererString")).
				String("post-renderer-args", "an argument to the post-renderer (can specify multiple)", ox.Spec("postRendererArgsSlice"), ox.Default("[]")).
				Bool("render-subchart-notes", "if set, render subchart notes along with the parent").
				String("repo", "chart repository url where to locate the requested chart").
				Bool("reset-then-reuse-values", "when upgrading, reset the values to the ones built into the chart, apply the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' or '--reuse-values' is specified, this is ignored").
				Bool("reset-values", "when upgrading, reset the values to the ones built into the chart").
				Bool("reuse-values", "when upgrading, reuse the last release's values and merge in any overrides from the command line via --set and -f. If '--reset-values' is specified, this is ignored").
				Slice("set", "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Slice("set-file", "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)", ox.Elem(ox.StringT)).
				Slice("set-json", "set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)", ox.Elem(ox.StringT)).
				Slice("set-literal", "set a literal STRING value on the command line", ox.Elem(ox.StringT)).
				Slice("set-string", "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)", ox.Elem(ox.StringT)).
				Bool("skip-crds", "if set, no CRDs will be installed when an upgrade is performed with install flag enabled. By default, CRDs are installed if not already present, when an upgrade is performed with install flag enabled").
				Bool("skip-schema-validation", "if set, disables JSON schema validation").
				Duration("timeout", "time to wait for any individual Kubernetes operation (like Jobs for hooks)", ox.Default("5m0s")).
				String("username", "chart repository username where to locate the requested chart").
				Slice("values", "specify values in a YAML file or a URL (can specify multiple)", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("verify", "verify the package before using it").
				Bool("wait", "if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout").
				Bool("wait-for-jobs", "if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout").
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Sub(
			ox.Banner("\nVerify that the given chart has a valid provenance file.\n\nProvenance files provide cryptographic verification that a chart has not been\ntampered with, and was packaged by a trusted provider.\n\nThis command can be used to verify a local chart. Several other commands provide\n'--verify' flags that run the same validation. To generate a signed package, use\nthe 'helm package --sign' command."),
			ox.Usage("verify", "verify that a chart at the given path has been signed and is valid"),
			ox.Spec("PATH [flags]"),
			ox.Flags().
				String("keyring", "keyring containing public keys", ox.Default("$HOME/.gnupg/pubring.gpg")).
				Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
				Bool("debug", "enable verbose output").
				String("kube-apiserver", "the address and the port for the Kubernetes API server").
				Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
				String("kube-as-user", "username to impersonate for the operation").
				String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
				String("kube-context", "name of the kubeconfig context to use").
				Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
				String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
				String("kube-token", "bearer token used for authentication").
				String("kubeconfig", "path to the kubeconfig file").
				String("namespace", "namespace scope for this request", ox.Short("n")).
				Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
				String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
				String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
				String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
		), ox.Flags().
			Int("burst-limit", "client-side default throttling limit", ox.Default("100")).
			Bool("debug", "enable verbose output").
			String("kube-apiserver", "the address and the port for the Kubernetes API server").
			Slice("kube-as-group", "group to impersonate for the operation, this flag can be repeated to specify multiple groups.", ox.Elem(ox.StringT)).
			String("kube-as-user", "username to impersonate for the operation").
			String("kube-ca-file", "the certificate authority file for the Kubernetes API server connection").
			String("kube-context", "name of the kubeconfig context to use").
			Bool("kube-insecure-skip-tls-verify", "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure").
			String("kube-tls-server-name", "server name to use for Kubernetes API server certificate validation. If it is not provided, the hostname used to contact the server is used").
			String("kube-token", "bearer token used for authentication").
			String("kubeconfig", "path to the kubeconfig file").
			String("namespace", "namespace scope for this request", ox.Short("n")).
			Float32("qps", "queries per second used when communicating with the Kubernetes API, not including bursting").
			String("registry-config", "path to the registry config file", ox.Default("$APPCONFIG/registry/config.json")).
			String("repository-cache", "path to the directory containing cached repository indexes", ox.Default("$APPCACHE/repository")).
			String("repository-config", "path to the file containing repository names and URLs", ox.Default("$APPCONFIG/repositories.yaml")),
	)
}
