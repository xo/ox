// Command kubectl is a xo/ox version of `+kubectl`.
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
		ox.Banner("kubectl controls the Kubernetes cluster manager.\n\n Find more information at: https://kubernetes.io/docs/reference/kubectl/"),
		ox.Usage("kubectl", ""),
		ox.Spec("[flags] [options]"),
		ox.Sections("Basic Commands (Beginner)", "Basic Commands (Intermediate)", "Deploy Commands", "Cluster Management Commands", "Troubleshooting and Debugging Commands", "Advanced Commands", "Settings Commands", "Other Commands"),
		ox.Footer("Use \"kubectl <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
		ox.Sub(
			ox.Banner("Create a resource from a file or from stdin.\n\n JSON and YAML formats are accepted."),
			ox.Usage("create", "Create a resource from a file or from stdin"),
			ox.Spec("-f FILENAME [options]"),
			ox.Example("\n  # Create a pod using the data in pod.json\n  kubectl create -f ./pod.json\n  \n  # Create a pod based on the JSON passed into stdin\n  cat pod.json | kubectl create -f -\n  \n  # Edit the data in registry.yaml in JSON then create the resource using the edited data\n  kubectl create -f registry.yaml --edit -o json"),
			ox.Section(0),
			ox.Footer("Use \"kubectl create <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Create a cluster role."),
				ox.Usage("clusterrole", "Create a cluster role"),
				ox.Spec("NAME --verb=verb --resource=resource.group [--resource-name=resourcename] [--dry-run=server|client|none] [options]"),
				ox.Example("\n  # Create a cluster role named \"pod-reader\" that allows user to perform \"get\", \"watch\" and \"list\" on pods\n  kubectl create clusterrole pod-reader --verb=get,list,watch --resource=pods\n  \n  # Create a cluster role named \"pod-reader\" with ResourceName specified\n  kubectl create clusterrole pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod\n  \n  # Create a cluster role named \"foo\" with API Group specified\n  kubectl create clusterrole foo --verb=get,list,watch --resource=rs.apps\n  \n  # Create a cluster role named \"foo\" with SubResource specified\n  kubectl create clusterrole foo --verb=get,list,watch --resource=pods,pods/status\n  \n  # Create a cluster role name \"foo\" with NonResourceURL specified\n  kubectl create clusterrole \"foo\" --verb=get --non-resource-url=/logs/*\n  \n  # Create a cluster role name \"monitoring\" with AggregationRule specified\n  kubectl create clusterrole monitoring --aggregation-rule=\"rbac.example.com/aggregate-to-monitoring=true\""),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("aggregation-rule", "An aggregation label selector for combining ClusterRoles.").
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Slice("non-resource-url", "A partial url that user should have access to.", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Slice("resource", "Resource that the rule applies to", ox.Elem(ox.StringT)).
					Slice("resource-name", "Resource in the white list that the rule applies to, repeat this flag for multiple items", ox.Elem(ox.StringT)).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
					Slice("verb", "Verb that applies to the resources contained in the rule", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Create a cluster role binding for a particular cluster role."),
				ox.Usage("clusterrolebinding", "Create a cluster role binding for a particular cluster role"),
				ox.Spec("NAME --clusterrole=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]"),
				ox.Example("\n  # Create a cluster role binding for user1, user2, and group1 using the cluster-admin cluster role\n  kubectl create clusterrolebinding cluster-admin --clusterrole=cluster-admin --user=user1 --user=user2 --group=group1"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("clusterrole", "ClusterRole this ClusterRoleBinding should reference").
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Slice("group", "Groups to bind to the clusterrole. The flag can be repeated to add multiple groups.", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Slice("serviceaccount", "Service accounts to bind to the clusterrole, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.", ox.Elem(ox.StringT)).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					Slice("user", "Usernames to bind to the clusterrole. The flag can be repeated to add multiple users.", ox.Elem(ox.StringT)).
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a config map based on a file, directory, or specified literal value.\n\n A single config map may package one or more key/value pairs.\n\n When creating a config map based on a file, the key will default to the basename of the file, and the value will default to the file content.  If the basename is an invalid key, you may specify an alternate key.\n\n When creating a config map based on a directory, each file whose basename is a valid key in the directory will be packaged into the config map.  Any directory entries except regular files are ignored (e.g. subdirectories, symlinks, devices, pipes, etc)."),
				ox.Usage("configmap", "Create a config map from a local file, directory or literal value"),
				ox.Spec("NAME [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none] [options]"),
				ox.Aliases("configmap", "cm"),
				ox.Example("\n  # Create a new config map named my-config based on folder bar\n  kubectl create configmap my-config --from-file=path/to/bar\n  \n  # Create a new config map named my-config with specified keys instead of file basenames on disk\n  kubectl create configmap my-config --from-file=key1=/path/to/bar/file1.txt --from-file=key2=/path/to/bar/file2.txt\n  \n  # Create a new config map named my-config with key1=config1 and key2=config2\n  kubectl create configmap my-config --from-literal=key1=config1 --from-literal=key2=config2\n  \n  # Create a new config map named my-config from the key=value pairs in the file\n  kubectl create configmap my-config --from-file=path/to/bar\n  \n  # Create a new config map named my-config from an env file\n  kubectl create configmap my-config --from-env-file=path/to/foo.env --from-env-file=path/to/bar.env"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Bool("append-hash", "Append a hash of the configmap to its name.", ox.Spec("false")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Slice("from-env-file", "Specify the path to a file to read lines of key=val pairs to create a configmap.", ox.Elem(ox.StringT)).
					Slice("from-file", "Key file can be specified using its file path, in which case file basename will be used as configmap key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid configmap key.", ox.Elem(ox.StringT)).
					Slice("from-literal", "Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a cron job with the specified name."),
				ox.Usage("cronjob", "Create a cron job with the specified name"),
				ox.Spec("NAME --image=image --schedule='0/5 * * * ?' -- [COMMAND] [args...] [flags] [options]"),
				ox.Aliases("cronjob", "cj"),
				ox.Example("\n  # Create a cron job\n  kubectl create cronjob my-job --image=busybox --schedule=\"*/1 * * * *\"\n  \n  # Create a cron job with a command\n  kubectl create cronjob my-job --image=busybox --schedule=\"*/1 * * * *\" -- date"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("image", "Image name to run.").
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					String("restart", "job's restart policy. supported values: OnFailure, Never").
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					String("schedule", "A schedule in the Cron format the job should be run with.").
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a deployment with the specified name."),
				ox.Usage("deployment", "Create a deployment with the specified name"),
				ox.Spec("NAME --image=image -- [COMMAND] [args...] [options]"),
				ox.Aliases("deployment", "deploy"),
				ox.Example("\n  # Create a deployment named my-dep that runs the busybox image\n  kubectl create deployment my-dep --image=busybox\n  \n  # Create a deployment with a command\n  kubectl create deployment my-dep --image=busybox -- date\n  \n  # Create a deployment named my-dep that runs the nginx image with 3 replicas\n  kubectl create deployment my-dep --image=nginx --replicas=3\n  \n  # Create a deployment named my-dep that runs the busybox image and expose port 5701\n  kubectl create deployment my-dep --image=busybox --port=5701\n  \n  # Create a deployment named my-dep that runs multiple containers\n  kubectl create deployment my-dep --image=busybox:latest --image=ubuntu:latest --image=nginx"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Slice("image", "Image names to run. A deployment can have multiple images set for multi-container pod.", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Int("port", "The containerPort that this deployment exposes.", ox.Default("-1")).
					Int("replicas", "Number of replicas to create. Default is 1.", ox.Default("1"), ox.Short("r")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create an ingress with the specified name."),
				ox.Usage("ingress", "Create an ingress with the specified name"),
				ox.Spec("NAME --rule=host/path=service:port[,tls[=secret]]  [options]"),
				ox.Aliases("ingress", "ing"),
				ox.Example("\n  # Create a single ingress called 'simple' that directs requests to foo.com/bar to svc\n  # svc1:8080 with a TLS secret \"my-cert\"\n  kubectl create ingress simple --rule=\"foo.com/bar=svc1:8080,tls=my-cert\"\n  \n  # Create a catch all ingress of \"/path\" pointing to service svc:port and Ingress Class as \"otheringress\"\n  kubectl create ingress catch-all --class=otheringress --rule=\"/path=svc:port\"\n  \n  # Create an ingress with two annotations: ingress.annotation1 and ingress.annotations2\n  kubectl create ingress annotated --class=default --rule=\"foo.com/bar=svc:port\" \\\n  --annotation ingress.annotation1=foo \\\n  --annotation ingress.annotation2=bla\n  \n  # Create an ingress with the same host and multiple paths\n  kubectl create ingress multipath --class=default \\\n  --rule=\"foo.com/=svc:port\" \\\n  --rule=\"foo.com/admin/=svcadmin:portadmin\"\n  \n  # Create an ingress with multiple hosts and the pathType as Prefix\n  kubectl create ingress ingress1 --class=default \\\n  --rule=\"foo.com/path*=svc:8080\" \\\n  --rule=\"bar.com/admin*=svc2:http\"\n  \n  # Create an ingress with TLS enabled using the default ingress certificate and different path types\n  kubectl create ingress ingtls --class=default \\\n  --rule=\"foo.com/=svc:https,tls\" \\\n  --rule=\"foo.com/path/subpath*=othersvc:8080\"\n  \n  # Create an ingress with TLS enabled using a specific secret and pathType as Prefix\n  kubectl create ingress ingsecret --class=default \\\n  --rule=\"foo.com/*=svc:8080,tls=secret1\"\n  \n  # Create an ingress with a default backend\n  kubectl create ingress ingdefault --class=default \\\n  --default-backend=defaultsvc:http \\\n  --rule=\"foo.com/*=svc:8080,tls=secret1\""),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("annotation", "Annotation to insert in the ingress object, in the format annotation=value", ox.Elem(ox.StringT)).
					String("class", "Ingress Class to be used").
					String("default-backend", "Default service for backend, in format of svcname:port").
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Slice("rule", "Rule in format host/path=service:port[,tls=secretname]. Paths containing the leading character '*' are considered pathType=Prefix. tls argument is optional.", ox.Elem(ox.StringT)).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a job with the specified name."),
				ox.Usage("job", "Create a job with the specified name"),
				ox.Spec("NAME --image=image [--from=cronjob/name] -- [COMMAND] [args...] [options]"),
				ox.Example("\n  # Create a job\n  kubectl create job my-job --image=busybox\n  \n  # Create a job with a command\n  kubectl create job my-job --image=busybox -- date\n  \n  # Create a job from a cron job named \"a-cronjob\"\n  kubectl create job test-job --from=cronjob/a-cronjob"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("from", "The name of the resource to create a Job from (only cronjob is supported).").
					String("image", "Image name to run.").
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a namespace with the specified name."),
				ox.Usage("namespace", "Create a namespace with the specified name"),
				ox.Spec("NAME [--dry-run=server|client|none] [options]"),
				ox.Aliases("namespace", "ns"),
				ox.Example("\n  # Create a new namespace named my-namespace\n  kubectl create namespace my-namespace"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a pod disruption budget with the specified name, selector, and desired minimum available pods."),
				ox.Usage("poddisruptionbudget", "Create a pod disruption budget with the specified name"),
				ox.Spec("NAME --selector=SELECTOR --min-available=N [--dry-run=server|client|none] [options]"),
				ox.Aliases("poddisruptionbudget", "pdb"),
				ox.Example("\n  # Create a pod disruption budget named my-pdb that will select all pods with the app=rails label\n  # and require at least one of them being available at any point in time\n  kubectl create poddisruptionbudget my-pdb --selector=app=rails --min-available=1\n  \n  # Create a pod disruption budget named my-pdb that will select all pods with the app=nginx label\n  # and require at least half of the pods selected to be available at any point in time\n  kubectl create pdb my-pdb --selector=app=nginx --min-available=50%"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("max-unavailable", "The maximum number or percentage of unavailable pods this budget requires.").
					String("min-available", "The minimum number or percentage of available pods this budget requires.").
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					String("selector", "A label selector to use for this budget. Only equality-based selector requirements are supported.").
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a priority class with the specified name, value, globalDefault and description."),
				ox.Usage("priorityclass", "Create a priority class with the specified name"),
				ox.Spec("NAME --value=VALUE --global-default=BOOL [--dry-run=server|client|none] [options]"),
				ox.Aliases("priorityclass", "pc"),
				ox.Example("\n  # Create a priority class named high-priority\n  kubectl create priorityclass high-priority --value=1000 --description=\"high priority\"\n  \n  # Create a priority class named default-priority that is considered as the global default priority\n  kubectl create priorityclass default-priority --value=1000 --global-default=true --description=\"default priority\"\n  \n  # Create a priority class named high-priority that cannot preempt pods with lower priority\n  kubectl create priorityclass high-priority --value=1000 --description=\"high priority\" --preemption-policy=\"Never\""),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("description", "description is an arbitrary string that usually provides guidelines on when this priority class should be used.").
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Bool("global-default", "global-default specifies whether this PriorityClass should be considered as the default priority.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					String("preemption-policy", "preemption-policy is the policy for preempting pods with lower priority.", ox.Default("PreemptLowerPriority")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
					Int("value", "the value of this priority class.", ox.Default("0")),
			), ox.Sub(
				ox.Banner("Create a resource quota with the specified name, hard limits, and optional scopes."),
				ox.Usage("quota", "Create a quota with the specified name"),
				ox.Spec("NAME [--hard=key1=value1,key2=value2] [--scopes=Scope1,Scope2] [--dry-run=server|client|none] [options]"),
				ox.Aliases("quota", "resourcequota"),
				ox.Example("\n  # Create a new resource quota named my-quota\n  kubectl create quota my-quota --hard=cpu=1,memory=1G,pods=2,services=3,replicationcontrollers=2,resourcequotas=1,secrets=5,persistentvolumeclaims=10\n  \n  # Create a new resource quota named best-effort\n  kubectl create quota best-effort --hard=pods=100 --scopes=BestEffort"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("hard", "A comma-delimited set of resource=quantity pairs that define a hard limit.").
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					String("scopes", "A comma-delimited set of quota scopes that must all match each object tracked by the quota.").
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a role with single rule."),
				ox.Usage("role", "Create a role with single rule"),
				ox.Spec("NAME --verb=verb --resource=resource.group/subresource [--resource-name=resourcename] [--dry-run=server|client|none] [options]"),
				ox.Example("\n  # Create a role named \"pod-reader\" that allows user to perform \"get\", \"watch\" and \"list\" on pods\n  kubectl create role pod-reader --verb=get --verb=list --verb=watch --resource=pods\n  \n  # Create a role named \"pod-reader\" with ResourceName specified\n  kubectl create role pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod\n  \n  # Create a role named \"foo\" with API Group specified\n  kubectl create role foo --verb=get,list,watch --resource=rs.apps\n  \n  # Create a role named \"foo\" with SubResource specified\n  kubectl create role foo --verb=get,list,watch --resource=pods,pods/status"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Slice("resource", "Resource that the rule applies to", ox.Elem(ox.StringT)).
					Slice("resource-name", "Resource in the white list that the rule applies to, repeat this flag for multiple items", ox.Elem(ox.StringT)).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
					Slice("verb", "Verb that applies to the resources contained in the rule", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Create a role binding for a particular role or cluster role."),
				ox.Usage("rolebinding", "Create a role binding for a particular role or cluster role"),
				ox.Spec("NAME --clusterrole=NAME|--role=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]"),
				ox.Example("\n  # Create a role binding for user1, user2, and group1 using the admin cluster role\n  kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1\n  \n  # Create a role binding for service account monitoring:sa-dev using the admin role\n  kubectl create rolebinding admin-binding --role=admin --serviceaccount=monitoring:sa-dev"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("clusterrole", "ClusterRole this RoleBinding should reference").
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					Slice("group", "Groups to bind to the role. The flag can be repeated to add multiple groups.", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					String("role", "Role this RoleBinding should reference").
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Slice("serviceaccount", "Service accounts to bind to the role, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.", ox.Elem(ox.StringT)).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					Slice("user", "Usernames to bind to the role. The flag can be repeated to add multiple users.", ox.Elem(ox.StringT)).
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Create a secret with specified type.\n\n A docker-registry type secret is for accessing a container registry.\n\n A generic type secret indicate an Opaque secret type.\n\n A tls type secret holds TLS certificate and its associated key."),
				ox.Usage("secret", "Create a secret using a specified subcommand"),
				ox.Spec("(docker-registry | generic | tls) [options]"),
				ox.Footer("Use \"kubectl create secret <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Sub(
					ox.Banner("Create a new secret for use with Docker registries.\n        \n        Dockercfg secrets are used to authenticate against Docker registries.\n        \n        When using the Docker command line to push images, you can authenticate to a given registry by running:\n        '$ docker login DOCKER_REGISTRY_SERVER --username=DOCKER_USER --password=DOCKER_PASSWORD --email=DOCKER_EMAIL'.\n        \n That produces a ~/.dockercfg file that is used by subsequent 'docker push' and 'docker pull' commands to authenticate to the registry. The email address is optional.\n\n        When creating applications, you may have a Docker registry that requires authentication.  In order for the\n        nodes to pull images on your behalf, they must have the credentials.  You can provide this information\n        by creating a dockercfg secret and attaching it to your service account."),
					ox.Usage("docker-registry", "Create a secret for use with a Docker registry"),
					ox.Spec("NAME --docker-username=user --docker-password=password --docker-email=email [--docker-server=string] [--from-file=[key=]source] [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # If you do not already have a .dockercfg file, create a dockercfg secret directly\n  kubectl create secret docker-registry my-secret --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-password=DOCKER_PASSWORD --docker-email=DOCKER_EMAIL\n  \n  # Create a new secret named my-secret from ~/.docker/config.json\n  kubectl create secret docker-registry my-secret --from-file=path/to/.docker/config.json"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						Bool("append-hash", "Append a hash of the secret to its name.", ox.Spec("false")).
						String("docker-email", "Email for Docker registry").
						String("docker-password", "Password for Docker registry authentication").
						String("docker-server", "Server location for Docker registry", ox.Default("https://index.docker.io/v1/")).
						String("docker-username", "Username for Docker registry authentication").
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						Slice("from-file", "Key files can be specified using their file path, in which case a default name of .dockerconfigjson will be given to them, or optionally with a name and file path, in which case the given name will be used. Specifying a directory will iterate each named file in the directory that is a valid secret key. For this command, the key should always be .dockerconfigjson.", ox.Elem(ox.StringT)).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				), ox.Sub(
					ox.Banner("Create a secret based on a file, directory, or specified literal value.\n\n A single secret may package one or more key/value pairs.\n\n When creating a secret based on a file, the key will default to the basename of the file, and the value will default to the file content. If the basename is an invalid key or you wish to chose your own, you may specify an alternate key.\n\n When creating a secret based on a directory, each file whose basename is a valid key in the directory will be packaged into the secret. Any directory entries except regular files are ignored (e.g. subdirectories, symlinks, devices, pipes, etc)."),
					ox.Usage("generic", "Create a secret from a local file, directory, or literal value"),
					ox.Spec("NAME [--type=string] [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new secret named my-secret with keys for each file in folder bar\n  kubectl create secret generic my-secret --from-file=path/to/bar\n  \n  # Create a new secret named my-secret with specified keys instead of names on disk\n  kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa --from-file=ssh-publickey=path/to/id_rsa.pub\n  \n  # Create a new secret named my-secret with key1=supersecret and key2=topsecret\n  kubectl create secret generic my-secret --from-literal=key1=supersecret --from-literal=key2=topsecret\n  \n  # Create a new secret named my-secret using a combination of a file and a literal\n  kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa --from-literal=passphrase=topsecret\n  \n  # Create a new secret named my-secret from env files\n  kubectl create secret generic my-secret --from-env-file=path/to/foo.env --from-env-file=path/to/bar.env"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						Bool("append-hash", "Append a hash of the secret to its name.", ox.Spec("false")).
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						Slice("from-env-file", "Specify the path to a file to read lines of key=val pairs to create a secret.", ox.Elem(ox.StringT)).
						Slice("from-file", "Key files can be specified using their file path, in which case a default name will be given to them, or optionally with a name and file path, in which case the given name will be used.  Specifying a directory will iterate each named file in the directory that is a valid secret key.", ox.Elem(ox.StringT)).
						Slice("from-literal", "Specify a key and literal value to insert in secret (i.e. mykey=somevalue)", ox.Elem(ox.StringT)).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("type", "The type of secret to create").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				), ox.Sub(
					ox.Banner("Create a TLS secret from the given public/private key pair.\n\n The public/private key pair must exist beforehand. The public key certificate must be .PEM encoded and match the given private key."),
					ox.Usage("tls", "Create a TLS secret"),
					ox.Spec("NAME --cert=path/to/cert/file --key=path/to/key/file [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new TLS secret named tls-secret with the given key pair\n  kubectl create secret tls tls-secret --cert=path/to/tls.crt --key=path/to/tls.key"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						Bool("append-hash", "Append a hash of the secret to its name.", ox.Spec("false")).
						String("cert", "Path to PEM encoded public key certificate.").
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						String("key", "Path to private key associated with given certificate.").
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				)), ox.Sub(
				ox.Banner("Create a service using a specified subcommand."),
				ox.Usage("service", "Create a service using a specified subcommand"),
				ox.Spec("[flags] [options]"),
				ox.Aliases("service", "svc"),
				ox.Footer("Use \"kubectl create service <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Sub(
					ox.Banner("Create a ClusterIP service with the specified name."),
					ox.Usage("clusterip", "Create a ClusterIP service"),
					ox.Spec("NAME [--tcp=<port>:<targetPort>] [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new ClusterIP service named my-cs\n  kubectl create service clusterip my-cs --tcp=5678:8080\n  \n  # Create a new ClusterIP service named my-cs (in headless mode)\n  kubectl create service clusterip my-cs --clusterip=\"None\""),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						String("clusterip", "Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing).").
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						Slice("tcp", "Port pairs can be specified as '<port>:<targetPort>'.", ox.Elem(ox.StringT)).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				), ox.Sub(
					ox.Banner("Create an ExternalName service with the specified name.\n\n ExternalName service references to an external DNS address instead of only pods, which will allow application authors to reference services that exist off platform, on other clusters, or locally."),
					ox.Usage("externalname", "Create an ExternalName service"),
					ox.Spec("NAME --external-name external.name [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new ExternalName service named my-ns\n  kubectl create service externalname my-ns --external-name bar.com"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("external-name", "External name of service").
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						Slice("tcp", "Port pairs can be specified as '<port>:<targetPort>'.", ox.Elem(ox.StringT)).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				), ox.Sub(
					ox.Banner("Create a LoadBalancer service with the specified name."),
					ox.Usage("loadbalancer", "Create a LoadBalancer service"),
					ox.Spec("NAME [--tcp=port:targetPort] [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new LoadBalancer service named my-lbs\n  kubectl create service loadbalancer my-lbs --tcp=5678:8080"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						Slice("tcp", "Port pairs can be specified as '<port>:<targetPort>'.", ox.Elem(ox.StringT)).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				), ox.Sub(
					ox.Banner("Create a NodePort service with the specified name."),
					ox.Usage("nodeport", "Create a NodePort service"),
					ox.Spec("NAME [--tcp=port:targetPort] [--dry-run=server|client|none] [options]"),
					ox.Example("\n  # Create a new NodePort service named my-ns\n  kubectl create service nodeport my-ns --tcp=5678:8080"),
					ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
					ox.Flags().
						Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
						String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
						String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
						Int("node-port", "Port used to expose the service on each node in a cluster.", ox.Default("0")).
						String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
						Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
						Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
						Slice("tcp", "Port pairs can be specified as '<port>:<targetPort>'.", ox.Elem(ox.StringT)).
						String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
						String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
				)), ox.Sub(
				ox.Banner("Create a service account with the specified name."),
				ox.Usage("serviceaccount", "Create a service account with the specified name"),
				ox.Spec("NAME [--dry-run=server|client|none] [options]"),
				ox.Aliases("serviceaccount", "sa"),
				ox.Example("\n  # Create a new service account named my-service-account\n  kubectl create serviceaccount my-service-account"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
			), ox.Sub(
				ox.Banner("Request a service account token."),
				ox.Usage("token", "Request a service account token"),
				ox.Spec("SERVICE_ACCOUNT_NAME [options]"),
				ox.Example("\n  # Request a token to authenticate to the kube-apiserver as the service account \"myapp\" in the current namespace\n  kubectl create token myapp\n  \n  # Request a token for a service account in a custom namespace\n  kubectl create token myapp --namespace myns\n  \n  # Request a token with a custom expiration\n  kubectl create token myapp --duration 10m\n  \n  # Request a token with a custom audience\n  kubectl create token myapp --audience https://example.com\n  \n  # Request a token bound to an instance of a Secret object\n  kubectl create token myapp --bound-object-kind Secret --bound-object-name mysecret\n  \n  # Request a token bound to an instance of a Secret object with a specific UID\n  kubectl create token myapp --bound-object-kind Secret --bound-object-name mysecret --bound-object-uid 0d4691ed-659b-4935-a832-355f77ee47cc"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("audience", "Audience of the requested token. If unset, defaults to requesting a token for use with the Kubernetes API server. May be repeated to request a token valid for multiple audiences.", ox.Elem(ox.StringT)).
					String("bound-object-kind", "Kind of an object to bind the token to. Supported kinds are Node, Pod, Secret. If set, --bound-object-name must be provided.").
					String("bound-object-name", "Name of an object to bind the token to. The token will expire when the object is deleted. Requires --bound-object-kind.").
					String("bound-object-uid", "UID of an object to bind the token to. Requires --bound-object-kind and --bound-object-name. If unset, the UID of the existing object is used.").
					Duration("duration", "Requested lifetime of the issued token. If not set or if set to 0, the lifetime will be determined by the server automatically. The server may return a token with a longer or shorter lifetime.", ox.Default("0s")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				Bool("edit", "Edit the API resource before creating", ox.Spec("false")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-create")).
				Slice("filename", "Filename, directory, or URL to files to use to create the resource", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				String("raw", "Raw URI to POST to the server.  Uses the transport specified by the kubeconfig file.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
				Bool("windows-line-endings", "Only relevant if --edit=true. Defaults to the line ending native to your platform.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Expose a resource as a new Kubernetes service.\n\n Looks up a deployment, service, replica set, replication controller or pod by name and uses the selector for that resource as the selector for a new service on the specified port. A deployment or replica set will be exposed as a service only if its selector is convertible to a selector that service supports, i.e. when the selector contains only the matchLabels component. Note that if no port is specified via --port and the exposed resource has multiple ports, all will be re-used by the new service. Also if no labels are specified, the new service will re-use the labels from the resource it exposes.\n\n Possible resources include (case insensitive):\n\n pod (po), service (svc), replicationcontroller (rc), deployment (deploy), replicaset (rs)"),
			ox.Usage("expose", "Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service"),
			ox.Spec("(-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type] [options]"),
			ox.Example("\n  # Create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000\n  kubectl expose rc nginx --port=80 --target-port=8000\n  \n  # Create a service for a replication controller identified by type and name specified in \"nginx-controller.yaml\", which serves on port 80 and connects to the containers on port 8000\n  kubectl expose -f nginx-controller.yaml --port=80 --target-port=8000\n  \n  # Create a service for a pod valid-pod, which serves on port 444 with the name \"frontend\"\n  kubectl expose pod valid-pod --port=444 --name=frontend\n  \n  # Create a second service based on the above service, exposing the container port 8443 as port 443 with the name \"nginx-https\"\n  kubectl expose service nginx --port=443 --target-port=8443 --name=nginx-https\n  \n  # Create a service for a replicated streaming application on port 4100 balancing UDP traffic and named 'video-stream'.\n  kubectl expose rc streamer --port=4100 --protocol=UDP --name=video-stream\n  \n  # Create a service for a replicated nginx using replica set, which serves on port 80 and connects to the containers on port 8000\n  kubectl expose rs nginx --port=80 --target-port=8000\n  \n  # Create a service for an nginx deployment, which serves on port 80 and connects to the containers on port 8000\n  kubectl expose deployment nginx --port=80 --target-port=8000"),
			ox.Section(0),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("cluster-ip", "ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service.").
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("external-ip", "Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.").
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-expose")).
				Slice("filename", "Filename, directory, or URL to files identifying the resource to expose a service", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("labels", "Labels to apply to the service created by this call.", ox.Short("l")).
				String("load-balancer-ip", "IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider specific).").
				String("name", "The name for the newly created object.").
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				String("override-type", "The method used to override the generated object: json, merge, or strategic.", ox.Default("merge")).
				String("overrides", "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.").
				String("port", "The port that the service should serve on. Copied from the resource being exposed, if unspecified").
				String("protocol", "The network protocol for the service to be created. Default is 'TCP'.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				String("selector", "A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.)").
				String("session-affinity", "If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'").
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("target-port", "Name or number for the port on the container that the service should direct traffic to. Optional.").
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				String("type", "Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'."),
		), ox.Sub(
			ox.Banner("Create and run a particular image in a pod."),
			ox.Usage("run", "Run a particular image on the cluster"),
			ox.Spec("NAME --image=image [--env=\"key=value\"] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...] [options]"),
			ox.Example("\n  # Start a nginx pod\n  kubectl run nginx --image=nginx\n  \n  # Start a hazelcast pod and let the container expose port 5701\n  kubectl run hazelcast --image=hazelcast/hazelcast --port=5701\n  \n  # Start a hazelcast pod and set environment variables \"DNS_DOMAIN=cluster\" and \"POD_NAMESPACE=default\" in the container\n  kubectl run hazelcast --image=hazelcast/hazelcast --env=\"DNS_DOMAIN=cluster\" --env=\"POD_NAMESPACE=default\"\n  \n  # Start a hazelcast pod and set labels \"app=hazelcast\" and \"env=prod\" in the container\n  kubectl run hazelcast --image=hazelcast/hazelcast --labels=\"app=hazelcast,env=prod\"\n  \n  # Dry run; print the corresponding API objects without creating them\n  kubectl run nginx --image=nginx --dry-run=client\n  \n  # Start a nginx pod, but overload the spec with a partial set of values parsed from JSON\n  kubectl run nginx --image=nginx --overrides='{ \"apiVersion\": \"v1\", \"spec\": { ... } }'\n  \n  # Start a busybox pod and keep it in the foreground, don't restart it if it exits\n  kubectl run -i -t busybox --image=busybox --restart=Never\n  \n  # Start the nginx pod using the default command, but use custom arguments (arg1 .. argN) for that command\n  kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN>\n  \n  # Start the nginx pod using a different command and custom arguments\n  kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN>"),
			ox.Section(0),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				Slice("annotations", "Annotations to apply to the pod.", ox.Elem(ox.StringT)).
				Bool("attach", "If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned.", ox.Spec("false")).
				String("cascade", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.", ox.Default("background")).
				Bool("command", "If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default.", ox.Spec("false")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				Slice("env", "Environment variables to set in the container.", ox.Elem(ox.StringT)).
				Bool("expose", "If true, create a ClusterIP service associated with the pod.  Requires `--port`.", ox.Spec("false")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-run")).
				Slice("filename", "to use to replace the resource.", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("force", "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.", ox.Spec("false")).
				Int("grace-period", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).", ox.Default("-1")).
				String("image", "The image for the container to run.").
				String("image-pull-policy", "The image pull policy for the container.  If left empty, this value will not be specified by the client and defaulted by the server.").
				String("kustomize", "Process a kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("labels", "Comma separated labels to apply to the pod. Will override previous values.", ox.Short("l")).
				Bool("leave-stdin-open", "If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				String("override-type", "The method used to override the generated object: json, merge, or strategic.", ox.Default("merge")).
				String("overrides", "An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.").
				Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("1m0s")).
				String("port", "The port that this container exposes.").
				Bool("privileged", "If true, run the container in privileged mode.", ox.Spec("false")).
				Bool("quiet", "If true, suppress prompt messages.", ox.Spec("false"), ox.Short("q")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("restart", "The restart policy for this Pod.  Legal values [Always, OnFailure, Never].", ox.Default("Always")).
				Bool("rm", "If true, delete the pod after it exits.  Only valid when attaching to the container, e.g. with '--attach' or with '-i/--stdin'.", ox.Spec("false")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				Bool("stdin", "Keep stdin open on the container in the pod, even if nothing is attached.", ox.Spec("false"), ox.Short("i")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Duration("timeout", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object", ox.Default("0s")).
				Bool("tty", "Allocate a TTY for the container in the pod.", ox.Spec("false"), ox.Short("t")).
				Bool("wait", "If true, wait for resources to be gone before returning. This waits for finalizers.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Configure application resources.\n\n These commands help you make changes to existing application resources."),
			ox.Usage("set", "Set specific features on objects"),
			ox.Spec("SUBCOMMAND [options]"),
			ox.Section(0),
			ox.Footer("Use \"kubectl set <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Update environment variables on a pod template.\n\n List environment variable definitions in one or more pods, pod templates. Add, update, or remove container environment variable definitions in one or more pod templates (within replication controllers or deployment configurations). View or modify the environment variable definitions on all containers in the specified pods or pod templates, or just those that match a wildcard.\n\n If \"--env -\" is passed, environment variables can be read from STDIN using the standard env syntax.\n\n Possible resources include (case insensitive):\n\n        pod (po), replicationcontroller (rc), deployment (deploy), daemonset (ds), statefulset (sts), cronjob (cj), replicaset (rs)"),
				ox.Usage("env", "Update environment variables on a pod template"),
				ox.Spec("RESOURCE/NAME KEY_1=VAL_1 ... KEY_N=VAL_N [options]"),
				ox.Example("\n  # Update deployment 'registry' with a new environment variable\n  kubectl set env deployment/registry STORAGE_DIR=/local\n  \n  # List the environment variables defined on a deployments 'sample-build'\n  kubectl set env deployment/sample-build --list\n  \n  # List the environment variables defined on all pods\n  kubectl set env pods --all --list\n  \n  # Output modified deployment in YAML, and does not alter the object on the server\n  kubectl set env deployment/sample-build STORAGE_DIR=/data -o yaml\n  \n  # Update all containers in all replication controllers in the project to have ENV=prod\n  kubectl set env rc --all ENV=prod\n  \n  # Import environment from a secret\n  kubectl set env --from=secret/mysecret deployment/myapp\n  \n  # Import environment from a config map with a prefix\n  kubectl set env --from=configmap/myconfigmap --prefix=MYSQL_ deployment/myapp\n  \n  # Import specific keys from a config map\n  kubectl set env --keys=my-example-key --from=configmap/myconfigmap deployment/myapp\n  \n  # Remove the environment variable ENV from container 'c1' in all deployment configs\n  kubectl set env deployments --all --containers=\"c1\" ENV-\n  \n  # Remove the environment variable ENV from a deployment definition on disk and\n  # update the deployment config on the server\n  kubectl set env -f deploy.json ENV-\n  \n  # Set some of the local shell environment into a deployment config on the server\n  env | grep RAILS_ | kubectl set env -e - deployment/registry"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "If true, select all resources in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("containers", "The names of containers in the selected pod templates to change - may use wildcards", ox.Default("*"), ox.Short("c")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					Slice("env", "Specify a key-value pair for an environment variable to set into each container.", ox.Elem(ox.StringT), ox.Short("e")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "Filename, directory, or URL to files the resource to update the env", ox.Elem(ox.StringT), ox.Short("f")).
					String("from", "The name of a resource from which to inject environment variables").
					Slice("keys", "Comma-separated list of keys to import from specified resource", ox.Elem(ox.StringT)).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					Bool("list", "If true, display the environment and any changes in the standard format. this flag will removed when we have kubectl view env.", ox.Spec("false")).
					Bool("local", "If true, set env will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("overwrite", "If true, allow environment to be overwritten, otherwise reject updates that overwrite existing environment.", ox.Spec("true")).
					String("prefix", "Prefix to append to variable names").
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("resolve", "If true, show secret or configmap references when listing variables", ox.Spec("false")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Update existing container image(s) of resources.\n\n Possible resources include (case insensitive):\n\n        pod (po), replicationcontroller (rc), deployment (deploy), daemonset (ds), statefulset (sts), cronjob (cj), replicaset (rs)"),
				ox.Usage("image", "Update the image of a pod template"),
				ox.Spec("(-f FILENAME | TYPE NAME) CONTAINER_NAME_1=CONTAINER_IMAGE_1 ... CONTAINER_NAME_N=CONTAINER_IMAGE_N [options]"),
				ox.Example("\n  # Set a deployment's nginx container image to 'nginx:1.9.1', and its busybox container image to 'busybox'\n  kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1\n  \n  # Update all deployments' and rc's nginx container's image to 'nginx:1.9.1'\n  kubectl set image deployments,rc nginx=nginx:1.9.1 --all\n  \n  # Update image of all containers of daemonset abc to 'nginx:1.9.1'\n  kubectl set image daemonset abc *=nginx:1.9.1\n  \n  # Print result (in yaml format) of updating nginx container image from local file, without hitting the server\n  kubectl set image -f path/to/file.yaml nginx=nginx:1.9.1 --local -o yaml"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources, in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					Bool("local", "If true, set image will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Specify compute resource requirements (CPU, memory) for any resource that defines a pod template.  If a pod is successfully scheduled, it is guaranteed the amount of resource requested, but may burst up to its specified limits.\n\n For each compute resource, if a limit is specified and a request is omitted, the request will default to the limit.\n\n Possible resources include (case insensitive): Use \"kubectl api-resources\" for a complete list of supported resources.."),
				ox.Usage("resources", "Update resource requests/limits on objects with pod templates"),
				ox.Spec("(-f FILENAME | TYPE NAME)  ([--limits=LIMITS & --requests=REQUESTS] [options]"),
				ox.Example("\n  # Set a deployments nginx container cpu limits to \"200m\" and memory to \"512Mi\"\n  kubectl set resources deployment nginx -c=nginx --limits=cpu=200m,memory=512Mi\n  \n  # Set the resource request and limits for all containers in nginx\n  kubectl set resources deployment nginx --limits=cpu=200m,memory=512Mi --requests=cpu=100m,memory=256Mi\n  \n  # Remove the resource requests for resources on containers in nginx\n  kubectl set resources deployment nginx --limits=cpu=0,memory=0 --requests=cpu=0,memory=0\n  \n  # Print the result (in yaml format) of updating nginx container limits from a local, without hitting the server\n  kubectl set resources -f path/to/file.yaml --limits=cpu=200m,memory=512Mi --local -o yaml"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources, in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("containers", "The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards", ox.Default("*"), ox.Short("c")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("limits", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.").
					Bool("local", "If true, set resources will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("requests", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.").
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Set the selector on a resource. Note that the new selector will overwrite the old selector if the resource had one prior to the invocation of 'set selector'.\n\n A selector must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters. If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used. Note: currently selectors can only be set on Service objects."),
				ox.Usage("selector", "Set the selector on a resource"),
				ox.Spec("(-f FILENAME | TYPE NAME) EXPRESSIONS [--resource-version=version] [options]"),
				ox.Example("\n  # Set the labels and selector before creating a deployment/service pair\n  kubectl create service clusterip my-svc --clusterip=\"None\" -o yaml --dry-run=client | kubectl set selector --local -f - 'environment=qa' -o yaml | kubectl create -f -\n  kubectl create deployment my-dep -o yaml --dry-run=client | kubectl label --local -f - environment=qa -o yaml | kubectl create -f -"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "identifying the resource.", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("local", "If true, annotation will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("true"), ox.Short("R")).
					String("resource-version", "If non-empty, the selectors update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.").
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Update the service account of pod template resources.\n\n Possible resources (case insensitive) can be:\n\n replicationcontroller (rc), deployment (deploy), daemonset (ds), job, replicaset (rs), statefulset"),
				ox.Usage("serviceaccount", "Update the service account of a resource"),
				ox.Spec("(-f FILENAME | TYPE NAME) SERVICE_ACCOUNT [options]"),
				ox.Aliases("serviceaccount", "sa"),
				ox.Example("\n  # Set deployment nginx-deployment's service account to serviceaccount1\n  kubectl set serviceaccount deployment nginx-deployment serviceaccount1\n  \n  # Print the result (in YAML format) of updated nginx deployment with the service account from local file, without hitting the API server\n  kubectl set sa -f nginx-deployment.yaml serviceaccount1 --local --dry-run=client -o yaml"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources, in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					Bool("local", "If true, set serviceaccount will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Update the user, group, or service account in a role binding or cluster role binding."),
				ox.Usage("subject", "Update the user, group, or service account in a role binding or cluster role binding"),
				ox.Spec("(-f FILENAME | TYPE NAME) [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]"),
				ox.Example("\n  # Update a cluster role binding for serviceaccount1\n  kubectl set subject clusterrolebinding admin --serviceaccount=namespace:serviceaccount1\n  \n  # Update a role binding for user1, user2, and group1\n  kubectl set subject rolebinding admin --user=user1 --user=user2 --group=group1\n  \n  # Print the result (in YAML format) of updating rolebinding subjects from a local, without hitting the server\n  kubectl create rolebinding admin --role=admin --user=admin -o yaml --dry-run=client | kubectl set subject --local -f - --user=foo -o yaml"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources, in the namespace of the specified resource types", ox.Spec("false")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-set")).
					Slice("filename", "Filename, directory, or URL to files the resource to update the subjects", ox.Elem(ox.StringT), ox.Short("f")).
					Slice("group", "Groups to bind to the role", ox.Elem(ox.StringT)).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					Bool("local", "If true, set subject will NOT contact api-server but run locally.", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Slice("serviceaccount", "Service accounts to bind to the role", ox.Elem(ox.StringT)).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					Slice("user", "Usernames to bind to the role", ox.Elem(ox.StringT)),
			)), ox.Sub(
			ox.Banner("Describe fields and structure of various resources.\n\n This command describes the fields associated with each supported API resource. Fields are identified via a simple JSONPath identifier:\n\n        <type>.<fieldName>[.<fieldName>]\n        \n Information about each field is retrieved from the server in OpenAPI format.\n\nUse \"kubectl api-resources\" for a complete list of supported resources."),
			ox.Usage("explain", "Get documentation for a resource"),
			ox.Spec("TYPE [--recursive=FALSE|TRUE] [--api-version=api-version-group] [--output=plaintext|plaintext-openapiv2] [options]"),
			ox.Example("\n  # Get the documentation of the resource and its fields\n  kubectl explain pods\n  \n  # Get all the fields in the resource\n  kubectl explain pods --recursive\n  \n  # Get the explanation for deployment in supported api versions\n  kubectl explain deployments --api-version=apps/v1\n  \n  # Get the documentation of a specific field of a resource\n  kubectl explain pods.spec.containers\n  \n  # Get the documentation of resources in different format\n  kubectl explain deployment --output=plaintext-openapiv2"),
			ox.Section(1),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("api-version", "Use given api-version (group/version) of the resource.").
				String("output", "Format in which to render the schema. Valid values are: (plaintext, plaintext-openapiv2).", ox.Default("plaintext")).
				Bool("recursive", "When true, print the name of all the fields recursively. Otherwise, print the available fields with their description.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Display one or many resources.\n\n Prints a table of the most important information about the specified resources. You can filter the list using a label selector and the --selector flag. If the desired resource type is namespaced you will only see results in the current namespace if you don't specify any namespace.\n\n By specifying the output as 'template' and providing a Go template as the value of the --template flag, you can filter the attributes of the fetched resources.\n\nUse \"kubectl api-resources\" for a complete list of supported resources."),
			ox.Usage("get", "Display one or many resources"),
			ox.Spec("[(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns|custom-columns-file|wide] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags] [options]"),
			ox.Example("\n  # List all pods in ps output format\n  kubectl get pods\n  \n  # List all pods in ps output format with more information (such as node name)\n  kubectl get pods -o wide\n  \n  # List a single replication controller with specified NAME in ps output format\n  kubectl get replicationcontroller web\n  \n  # List deployments in JSON output format, in the \"v1\" version of the \"apps\" API group\n  kubectl get deployments.v1.apps -o json\n  \n  # List a single pod in JSON output format\n  kubectl get -o json pod web-pod-13je7\n  \n  # List a pod identified by type and name specified in \"pod.yaml\" in JSON output format\n  kubectl get -f pod.yaml -o json\n  \n  # List resources from a directory with kustomization.yaml - e.g. dir/kustomization.yaml\n  kubectl get -k dir/\n  \n  # Return only the phase value of the specified pod\n  kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}\n  \n  # List resource information in custom columns\n  kubectl get pod test-pod -o custom-columns=CONTAINER:.spec.containers[0].name,IMAGE:.spec.containers[0].image\n  \n  # List all replication controllers and services together in ps output format\n  kubectl get rc,services\n  \n  # List one or more resources by their type and names\n  kubectl get rc/web service/frontend pods/web-pod-13je7\n  \n  # List the 'status' subresource for a single pod\n  kubectl get pod web-pod-13je7 --subresource status\n  \n  # List all deployments in namespace 'backend'\n  kubectl get deployments.apps --namespace backend\n  \n  # List all pods existing in all namespaces\n  kubectl get pods --all-namespaces"),
			ox.Section(1),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				Int("chunk-size", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.", ox.Default("500")).
				String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
				Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("ignore-not-found", "If the requested object does not exist the command will return exit code 0.", ox.Spec("false")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Slice("label-columns", "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2...", ox.Elem(ox.StringT), ox.Short("L")).
				Bool("no-headers", "When using the default or custom-column output format, don't print headers (default print headers).", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file, custom-columns, custom-columns-file, wide). See custom columns [https://kubernetes.io/docs/reference/kubectl/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/].", ox.Short("o")).
				Bool("output-watch-events", "Output watch event objects when --watch or --watch-only is used. Existing objects are output as initial ADDED events.", ox.Spec("false")).
				String("raw", "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("server-print", "If true, have the server return the appropriate table output. Supports extension APIs and CRDs.", ox.Spec("true")).
				Bool("show-kind", "If present, list the resource type for the requested object(s).", ox.Spec("false")).
				Bool("show-labels", "When printing, show all labels as the last column", ox.Spec("false"), ox.Default("hide labels column")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("sort-by", "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string.").
				String("subresource", "If specified, gets the subresource of the requested object. Must be one of [status scale]. This flag is beta and may change in the future.").
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Bool("watch", "After listing/getting the requested object, watch for changes.", ox.Spec("false"), ox.Short("w")).
				Bool("watch-only", "Watch for changes to the requested object(s), without listing/getting first.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Edit a resource from the default editor.\n\n The edit command allows you to directly edit any API resource you can retrieve via the command-line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows. When attempting to open the editor, it will first attempt to use the shell that has been defined in the 'SHELL' environment variable. If this is not defined, the default shell will be used, which is '/bin/bash' for Linux or 'cmd' for Windows.\n\n You can edit multiple objects, although changes are applied one at a time. The command accepts file names as well as command-line arguments, although the files you point to must be previously saved versions of resources.\n\n Editing is done with the API version used to fetch the resource. To edit using a specific API version, fully-qualify the resource, version, and group.\n\n The default format is YAML. To edit in JSON, specify \"-o json\".\n\n The flag --windows-line-endings can be used to force Windows line endings, otherwise the default for your operating system will be used.\n\n In the event an error occurs while updating, a temporary file will be created on disk that contains your unapplied changes. The most common error when updating a resource is another editor changing the resource on the server. When this occurs, you will have to apply your changes to the newer version of the resource, or update your temporary saved copy to include the latest resource version."),
			ox.Usage("edit", "Edit a resource on the server"),
			ox.Spec("(RESOURCE/NAME | -f FILENAME) [options]"),
			ox.Example("\n  # Edit the service named 'registry'\n  kubectl edit svc/registry\n  \n  # Use an alternative editor\n  KUBE_EDITOR=\"nano\" kubectl edit svc/registry\n  \n  # Edit the job 'myjob' in JSON using the v1 API format\n  kubectl edit job.v1.batch/myjob -o json\n  \n  # Edit the deployment 'mydeployment' in YAML and save the modified config in its annotation\n  kubectl edit deployment/mydeployment -o yaml --save-config\n  \n  # Edit the 'status' subresource for the 'mydeployment' deployment\n  kubectl edit deployment mydeployment --subresource='status'"),
			ox.Section(1),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-edit")).
				Slice("filename", "Filename, directory, or URL to files to use to edit the resource", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("output-patch", "Output the patch if the resource is edited.", ox.Spec("false")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("subresource", "If specified, edit will operate on the subresource of the requested object. Must be one of [status]. This flag is beta and may change in the future.").
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
				Bool("windows-line-endings", "Defaults to the line ending native to your platform.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Delete resources by file names, stdin, resources and names, or by resources and label selector.\n\n JSON and YAML formats are accepted. Only one type of argument may be specified: file names, resources and names, or resources and label selector.\n\n Some resources, such as pods, support graceful deletion. These resources define a default period before they are forcibly terminated (the grace period) but you may override that value with the --grace-period flag, or pass --now to set a grace-period of 1. Because these resources often represent entities in the cluster, deletion may not be acknowledged immediately. If the node hosting a pod is down or cannot reach the API server, termination may take significantly longer than the grace period. To force delete a resource, you must specify the --force flag. Note: only a subset of resources support graceful deletion. In absence of the support, the --grace-period flag is ignored.\n\n IMPORTANT: Force deleting pods does not wait for confirmation that the pod's processes have been terminated, which can leave those processes running until the node detects the deletion and completes graceful deletion. If your processes use shared storage or talk to a remote API and depend on the name of the pod to identify themselves, force deleting those pods may result in multiple processes running on different machines using the same identification which may lead to data corruption or inconsistency. Only force delete pods when you are sure the pod is terminated, or if your application can tolerate multiple copies of the same pod running at once. Also, if you force delete pods, the scheduler may place new pods on those nodes before the node has released those resources and causing those pods to be evicted immediately.\n\n Note that the delete command does NOT do resource version checks, so if someone submits an update to a resource right when you submit a delete, their update will be lost along with the rest of the resource.\n\n After a CustomResourceDefinition is deleted, invalidation of discovery cache may take up to 6 hours. If you don't want to wait, you might want to run \"kubectl api-resources\" to refresh the discovery cache."),
			ox.Usage("delete", "Delete resources by file names, stdin, resources and names, or by resources and label selector"),
			ox.Spec("([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)]) [options]"),
			ox.Example("\n  # Delete a pod using the type and name specified in pod.json\n  kubectl delete -f ./pod.json\n  \n  # Delete resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml\n  kubectl delete -k dir\n  \n  # Delete resources from all files that end with '.json'\n  kubectl delete -f '*.json'\n  \n  # Delete a pod based on the type and name in the JSON passed into stdin\n  cat pod.json | kubectl delete -f -\n  \n  # Delete pods and services with same names \"baz\" and \"foo\"\n  kubectl delete pod,service baz foo\n  \n  # Delete pods and services with label name=myLabel\n  kubectl delete pods,services -l name=myLabel\n  \n  # Delete a pod with minimal delay\n  kubectl delete pod foo --now\n  \n  # Force delete a pod on a dead node\n  kubectl delete pod foo --force\n  \n  # Delete all pods\n  kubectl delete pods --all"),
			ox.Section(1),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Delete all resources, in the namespace of the specified resource types.", ox.Spec("false")).
				Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
				String("cascade", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.", ox.Default("background")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
				Slice("filename", "containing the resource to delete.", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("force", "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.", ox.Spec("false")).
				Int("grace-period", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).", ox.Default("-1")).
				Bool("ignore-not-found", "Treat \"resource not found\" as a successful delete. Defaults to \"true\" when --all is specified.", ox.Spec("false")).
				Bool("interactive", "If true, delete resource only when user confirms.", ox.Spec("false"), ox.Short("i")).
				String("kustomize", "Process a kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("now", "If true, resources are signaled for immediate shutdown (same as --grace-period=1).", ox.Spec("false")).
				String("output", "Output mode. Use \"-o name\" for shorter output (resource/name).", ox.Short("o")).
				String("raw", "Raw URI to DELETE to the server.  Uses the transport specified by the kubeconfig file.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Duration("timeout", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object", ox.Default("0s")).
				Bool("wait", "If true, wait for resources to be gone before returning. This waits for finalizers.", ox.Spec("true")),
		), ox.Sub(
			ox.Banner("Manage the rollout of one or many resources.\n        \n Valid resource types include:\n\n  *  deployments\n  *  daemonsets\n  *  statefulsets"),
			ox.Usage("rollout", "Manage the rollout of a resource"),
			ox.Spec("SUBCOMMAND [options]"),
			ox.Example("\n  # Rollback to the previous deployment\n  kubectl rollout undo deployment/abc\n  \n  # Check the rollout status of a daemonset\n  kubectl rollout status daemonset/foo\n  \n  # Restart a deployment\n  kubectl rollout restart deployment/abc\n  \n  # Restart deployments with the 'app=nginx' label\n  kubectl rollout restart deployment --selector=app=nginx"),
			ox.Section(2),
			ox.Footer("Use \"kubectl rollout <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("View previous rollout revisions and configurations."),
				ox.Usage("history", "View rollout history"),
				ox.Spec("(TYPE NAME | TYPE/NAME) [flags] [options]"),
				ox.Example("\n  # View the rollout history of a deployment\n  kubectl rollout history deployment/abc\n  \n  # View the details of daemonset revision 3\n  kubectl rollout history daemonset/abc --revision=3"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Int("revision", "See the details, including podTemplate of the revision specified", ox.Default("0")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Mark the provided resource as paused.\n\n Paused resources will not be reconciled by a controller. Use \"kubectl rollout resume\" to resume a paused resource. Currently only deployments support being paused."),
				ox.Usage("pause", "Mark the provided resource as paused"),
				ox.Spec("RESOURCE [options]"),
				ox.Example("\n  # Mark the nginx deployment as paused\n  # Any current state of the deployment will continue its function; new updates\n  # to the deployment will not have an effect as long as the deployment is paused\n  kubectl rollout pause deployment/nginx"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-rollout")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Restart a resource.\n\n        Resource rollout will be restarted."),
				ox.Usage("restart", "Restart a resource"),
				ox.Spec("RESOURCE [options]"),
				ox.Example("\n  # Restart all deployments in the test-namespace namespace\n  kubectl rollout restart deployment -n test-namespace\n  \n  # Restart a deployment\n  kubectl rollout restart deployment/nginx\n  \n  # Restart a daemon set\n  kubectl rollout restart daemonset/abc\n  \n  # Restart deployments with the app=nginx label\n  kubectl rollout restart deployment --selector=app=nginx"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-rollout")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Resume a paused resource.\n\n Paused resources will not be reconciled by a controller. By resuming a resource, we allow it to be reconciled again. Currently only deployments support being resumed."),
				ox.Usage("resume", "Resume a paused resource"),
				ox.Spec("RESOURCE [options]"),
				ox.Example("\n  # Resume an already paused deployment\n  kubectl rollout resume deployment/nginx"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-rollout")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Show the status of the rollout.\n\n By default 'rollout status' will watch the status of the latest rollout until it's done. If you don't want to wait for the rollout to finish then you can use --watch=false. Note that if a new rollout starts in-between, then 'rollout status' will continue watching the latest revision. If you want to pin to a specific revision and abort if it is rolled over by another revision, use --revision=N where N is the revision you need to watch for."),
				ox.Usage("status", "Show the status of the rollout"),
				ox.Spec("(TYPE NAME | TYPE/NAME) [flags] [options]"),
				ox.Example("\n  # Watch the rollout status of a deployment\n  kubectl rollout status deployment/nginx"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Int("revision", "Pin to a specific revision for showing its status. Defaults to 0 (last revision).", ox.Default("0")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Duration("timeout", "The length of time to wait before ending watch, zero means never. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).", ox.Default("0s")).
					Bool("watch", "Watch the status of the rollout until it's done.", ox.Spec("true"), ox.Short("w")),
			), ox.Sub(
				ox.Banner("Roll back to a previous rollout."),
				ox.Usage("undo", "Undo a previous rollout"),
				ox.Spec("(TYPE NAME | TYPE/NAME) [flags] [options]"),
				ox.Example("\n  # Roll back to the previous deployment\n  kubectl rollout undo deployment/abc\n  \n  # Roll back to daemonset revision 3\n  kubectl rollout undo daemonset/abc --to-revision=3\n  \n  # Roll back to the previous deployment with dry-run\n  kubectl rollout undo --dry-run=server deployment/abc"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to get from a server.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					Int("to-revision", "The revision to rollback to. Default to 0 (last revision).", ox.Default("0")),
			)), ox.Sub(
			ox.Banner("Set a new size for a deployment, replica set, replication controller, or stateful set.\n\n Scale also allows users to specify one or more preconditions for the scale action.\n\n If --current-replicas or --resource-version is specified, it is validated before the scale is attempted, and it is guaranteed that the precondition holds true when the scale is sent to the server."),
			ox.Usage("scale", "Set a new size for a deployment, replica set, or replication controller"),
			ox.Spec("[--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME) [options]"),
			ox.Example("\n  # Scale a replica set named 'foo' to 3\n  kubectl scale --replicas=3 rs/foo\n  \n  # Scale a resource identified by type and name specified in \"foo.yaml\" to 3\n  kubectl scale --replicas=3 -f foo.yaml\n  \n  # If the deployment named mysql's current size is 2, scale mysql to 3\n  kubectl scale --current-replicas=2 --replicas=3 deployment/mysql\n  \n  # Scale multiple replication controllers\n  kubectl scale --replicas=5 rc/example1 rc/example2 rc/example3\n  \n  # Scale stateful set named 'web' to 3\n  kubectl scale --replicas=3 statefulset/web"),
			ox.Section(2),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Select all resources in the namespace of the specified resource types", ox.Spec("false")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				Int("current-replicas", "Precondition for current size. Requires that the current size of the resource match this value in order to scale. -1 (default) for no condition.", ox.Default("-1")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				Slice("filename", "Filename, directory, or URL to files identifying the resource to set a new size", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Int("replicas", "The new desired number of replicas. Required.", ox.Default("0")).
				String("resource-version", "Precondition for resource version. Requires that the current resource version match this value in order to scale.").
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Duration("timeout", "The length of time to wait before giving up on a scale operation, zero means don't wait. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).", ox.Default("0s")),
		), ox.Sub(
			ox.Banner("Creates an autoscaler that automatically chooses and sets the number of pods that run in a Kubernetes cluster.\n\n Looks up a deployment, replica set, stateful set, or replication controller by name and creates an autoscaler that uses the given resource as a reference. An autoscaler can automatically increase or decrease number of pods deployed within the system as needed."),
			ox.Usage("autoscale", "Auto-scale a deployment, replica set, stateful set, or replication controller"),
			ox.Spec("(-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [options]"),
			ox.Example("\n  # Auto scale a deployment \"foo\", with the number of pods between 2 and 10, no target CPU utilization specified so a default autoscaling policy will be used\n  kubectl autoscale deployment foo --min=2 --max=10\n  \n  # Auto scale a replication controller \"foo\", with the number of pods between 1 and 5, target CPU utilization at 80%\n  kubectl autoscale rc foo --max=5 --cpu-percent=80"),
			ox.Section(2),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				Int("cpu-percent", "The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used.", ox.Default("-1")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-autoscale")).
				Slice("filename", "Filename, directory, or URL to files identifying the resource to autoscale.", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Int("max", "The upper limit for the number of pods that can be set by the autoscaler. Required.", ox.Default("-1")).
				Int("min", "The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value.", ox.Default("-1")).
				String("name", "The name for the newly created object. If not specified, the name of the input resource will be used.").
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
		), ox.Sub(
			ox.Banner("Modify certificate resources."),
			ox.Usage("certificate", "Modify certificate resources"),
			ox.Spec("SUBCOMMAND [options]"),
			ox.Section(3),
			ox.Footer("Use \"kubectl certificate <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Approve a certificate signing request.\n\n kubectl certificate approve allows a cluster admin to approve a certificate signing request (CSR). This action tells a certificate signing controller to issue a certificate to the requester with the attributes requested in the CSR.\n\n SECURITY NOTICE: Depending on the requested attributes, the issued certificate can potentially grant a requester access to cluster resources or to authenticate as a requested identity. Before approving a CSR, ensure you understand what the signed certificate can do."),
				ox.Usage("approve", "Approve a certificate signing request"),
				ox.Spec("(-f FILENAME | NAME) [options]"),
				ox.Example("\n  # Approve CSR 'csr-sqgzp'\n  kubectl certificate approve csr-sqgzp"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to update", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("force", "Update the CSR even if it is already approved.", ox.Spec("false")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Deny a certificate signing request.\n\n kubectl certificate deny allows a cluster admin to deny a certificate signing request (CSR). This action tells a certificate signing controller to not to issue a certificate to the requester."),
				ox.Usage("deny", "Deny a certificate signing request"),
				ox.Spec("(-f FILENAME | NAME) [options]"),
				ox.Example("\n  # Deny CSR 'csr-sqgzp'\n  kubectl certificate deny csr-sqgzp"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to update", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("force", "Update the CSR even if it is already denied.", ox.Spec("false")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			)), ox.Sub(
			ox.Banner("Display addresses of the control plane and services with label kubernetes.io/cluster-service=true. To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'."),
			ox.Usage("cluster-info", "Display cluster information"),
			ox.Spec("[flags] [options]"),
			ox.Example("\n  # Print the address of the control plane and cluster services\n  kubectl cluster-info"),
			ox.Section(3),
			ox.Footer("Use \"kubectl cluster-info <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Dump cluster information out suitable for debugging and diagnosing cluster problems.  By default, dumps everything to stdout. You can optionally specify a directory with --output-directory.  If you specify a directory, Kubernetes will build a set of files in that directory.  By default, only dumps things in the current namespace and 'kube-system' namespace, but you can switch to a different namespace with the --namespaces flag, or specify --all-namespaces to dump all namespaces.\n\n The command also dumps the logs of all of the pods in the cluster; these logs are dumped into different directories based on namespace and pod name."),
				ox.Usage("dump", "Dump relevant information for debugging and diagnosis"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # Dump current cluster state to stdout\n  kubectl cluster-info dump\n  \n  # Dump current cluster state to /path/to/cluster-state\n  kubectl cluster-info dump --output-directory=/path/to/cluster-state\n  \n  # Dump all namespaces to stdout\n  kubectl cluster-info dump --all-namespaces\n  \n  # Dump a set of namespaces to /path/to/cluster-state\n  kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all-namespaces", "If true, dump all namespaces.  If true, --namespaces is ignored.", ox.Spec("false"), ox.Short("A")).
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Slice("namespaces", "A comma separated list of namespaces to dump.", ox.Elem(ox.StringT)).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Default("json"), ox.Short("o")).
					String("output-directory", "Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory").
					Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("20s")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			)), ox.Sub(
			ox.Banner("Display resource (CPU/memory) usage.\n\n The top command allows you to see the resource consumption for nodes or pods.\n\n This command requires Metrics Server to be correctly configured and working on the server."),
			ox.Usage("top", "Display resource (CPU/memory) usage"),
			ox.Spec("[flags] [options]"),
			ox.Section(3),
			ox.Footer("Use \"kubectl top <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Display resource (CPU/memory) usage of nodes.\n\n The top-node command allows you to see the resource consumption of nodes."),
				ox.Usage("node", "Display resource (CPU/memory) usage of nodes"),
				ox.Spec("[NAME | -l label] [options]"),
				ox.Aliases("node", "nodes", "no"),
				ox.Example("\n  # Show metrics for all nodes\n  kubectl top node\n  \n  # Show metrics for a given node\n  kubectl top node NODE_NAME"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("no-headers", "If present, print output without headers", ox.Spec("false")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					Bool("show-capacity", "Print node resources based on Capacity instead of Allocatable(default) of the nodes.", ox.Spec("false")).
					String("sort-by", "If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'.").
					Bool("use-protocol-buffers", "Enables using protocol-buffers to access Metrics API.", ox.Spec("true")),
			), ox.Sub(
				ox.Banner("Display resource (CPU/memory) usage of pods.\n\n The 'top pod' command allows you to see the resource consumption of pods.\n\n Due to the metrics pipeline delay, they may be unavailable for a few minutes since pod creation."),
				ox.Usage("pod", "Display resource (CPU/memory) usage of pods"),
				ox.Spec("[NAME | -l label] [options]"),
				ox.Aliases("pod", "pods", "po"),
				ox.Example("\n  # Show metrics for all pods in the default namespace\n  kubectl top pod\n  \n  # Show metrics for all pods in the given namespace\n  kubectl top pod --namespace=NAMESPACE\n  \n  # Show metrics for a given pod and its containers\n  kubectl top pod POD_NAME --containers\n  \n  # Show metrics for the pods defined by label name=myLabel\n  kubectl top pod -l name=myLabel"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
					Bool("containers", "If present, print usage of containers within a pod.", ox.Spec("false")).
					String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
					Bool("no-headers", "If present, print output without headers.", ox.Spec("false")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
					String("sort-by", "If non-empty, sort pods list using specified field. The field can be either 'cpu' or 'memory'.").
					Bool("sum", "Print the sum of the resource usage", ox.Spec("false")).
					Bool("use-protocol-buffers", "Enables using protocol-buffers to access Metrics API.", ox.Spec("true")),
			)), ox.Sub(
			ox.Banner("Mark node as unschedulable."),
			ox.Usage("cordon", "Mark node as unschedulable"),
			ox.Spec("NODE [options]"),
			ox.Example("\n  # Mark node \"foo\" as unschedulable\n  kubectl cordon foo"),
			ox.Section(3),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Mark node as schedulable."),
			ox.Usage("uncordon", "Mark node as schedulable"),
			ox.Spec("NODE [options]"),
			ox.Example("\n  # Mark node \"foo\" as schedulable\n  kubectl uncordon foo"),
			ox.Section(3),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Drain node in preparation for maintenance.\n\n The given node will be marked unschedulable to prevent new pods from arriving. 'drain' evicts the pods if the API server supports https://kubernetes.io/docs/concepts/workloads/pods/disruptions/ eviction https://kubernetes.io/docs/concepts/workloads/pods/disruptions/ . Otherwise, it will use normal DELETE to delete the pods. The 'drain' evicts or deletes all pods except mirror pods (which cannot be deleted through the API server).  If there are daemon set-managed pods, drain will not proceed without --ignore-daemonsets, and regardless it will not delete any daemon set-managed pods, because those pods would be immediately replaced by the daemon set controller, which ignores unschedulable markings.  If there are any pods that are neither mirror pods nor managed by a replication controller, replica set, daemon set, stateful set, or job, then drain will not delete any pods unless you use --force.  --force will also allow deletion to proceed if the managing resource of one or more pods is missing.\n\n 'drain' waits for graceful termination. You should not operate on the machine until the command completes.\n\n When you are ready to put the node back into service, use kubectl uncordon, which will make the node schedulable again.\n\nhttps://kubernetes.io/images/docs/kubectl_drain.svg Workflowhttps://kubernetes.io/images/docs/kubectl_drain.svg"),
			ox.Usage("drain", "Drain node in preparation for maintenance"),
			ox.Spec("NODE [options]"),
			ox.Example("\n  # Drain node \"foo\", even if there are pods not managed by a replication controller, replica set, job, daemon set, or stateful set on it\n  kubectl drain foo --force\n  \n  # As above, but abort if there are pods not managed by a replication controller, replica set, job, daemon set, or stateful set, and use a grace period of 15 minutes\n  kubectl drain foo --grace-period=900"),
			ox.Section(3),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Int("chunk-size", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.", ox.Default("500")).
				Bool("delete-emptydir-data", "Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained).", ox.Spec("false")).
				Bool("disable-eviction", "Force drain to use delete, even if eviction is supported. This will bypass checking PodDisruptionBudgets, use with caution.", ox.Spec("false")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				Bool("force", "Continue even if there are pods that do not declare a controller.", ox.Spec("false")).
				Int("grace-period", "Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used.", ox.Default("-1")).
				Bool("ignore-daemonsets", "Ignore DaemonSet-managed pods.", ox.Spec("false")).
				String("pod-selector", "Label selector to filter pods on the node").
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Int("skip-wait-for-delete-timeout", "If pod DeletionTimestamp older than N seconds, skip waiting for the pod.  Seconds must be greater than 0 to skip.", ox.Default("0")).
				Duration("timeout", "The length of time to wait before giving up, zero means infinite", ox.Default("0s")),
		), ox.Sub(
			ox.Banner("Update the taints on one or more nodes.\n\n  *  A taint consists of a key, value, and effect. As an argument here, it is expressed as key=value:effect.\n  *  The key must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 253 characters.\n  *  Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app.\n  *  The value is optional. If given, it must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters.\n  *  The effect must be NoSchedule, PreferNoSchedule or NoExecute.\n  *  Currently taint can only apply to node."),
			ox.Usage("taint", "Update the taints on one or more nodes"),
			ox.Spec("NODE NAME KEY_1=VAL_1:TAINT_EFFECT_1 ... KEY_N=VAL_N:TAINT_EFFECT_N [options]"),
			ox.Example("\n  # Update node 'foo' with a taint with key 'dedicated' and value 'special-user' and effect 'NoSchedule'\n  # If a taint with that key and effect already exists, its value is replaced as specified\n  kubectl taint nodes foo dedicated=special-user:NoSchedule\n  \n  # Remove from node 'foo' the taint with key 'dedicated' and effect 'NoSchedule' if one exists\n  kubectl taint nodes foo dedicated:NoSchedule-\n  \n  # Remove from node 'foo' all the taints with key 'dedicated'\n  kubectl taint nodes foo dedicated-\n  \n  # Add a taint with key 'dedicated' on nodes having label myLabel=X\n  kubectl taint node -l myLabel=X  dedicated=foo:PreferNoSchedule\n  \n  # Add to node 'foo' a taint with key 'bar' and no value\n  kubectl taint nodes foo bar:NoSchedule"),
			ox.Section(3),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Select all nodes in the cluster", ox.Spec("false")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-taint")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("overwrite", "If true, allow taints to be overwritten, otherwise reject taint updates that overwrite existing taints.", ox.Spec("false")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")),
		), ox.Sub(
			ox.Banner("Show details of a specific resource or group of resources.\n\n Print a detailed description of the selected resources, including related resources such as events or controllers. You may select a single object by name, all objects of that type, provide a name prefix, or label selector. For example:\n\n        $ kubectl describe TYPE NAME_PREFIX\n        \n will first check for an exact match on TYPE and NAME_PREFIX. If no such resource exists, it will output details for every resource that has a name prefixed with NAME_PREFIX.\n\nUse \"kubectl api-resources\" for a complete list of supported resources."),
			ox.Usage("describe", "Show details of a specific resource or group of resources"),
			ox.Spec("(-f FILENAME | TYPE [NAME_PREFIX | -l label] | TYPE/NAME) [options]"),
			ox.Example("\n  # Describe a node\n  kubectl describe nodes kubernetes-node-emt8.c.myproject.internal\n  \n  # Describe a pod\n  kubectl describe pods/nginx\n  \n  # Describe a pod identified by type and name in \"pod.json\"\n  kubectl describe -f pod.json\n  \n  # Describe all pods\n  kubectl describe pods\n  \n  # Describe pods by label name=myLabel\n  kubectl describe pods -l name=myLabel\n  \n  # Describe all pods managed by the 'frontend' replication controller\n  # (rc-created pods get the name of the rc as a prefix in the pod name)\n  kubectl describe pods frontend"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
				Int("chunk-size", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.", ox.Default("500")).
				Slice("filename", "Filename, directory, or URL to files containing the resource to describe", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-events", "If true, display events related to the described object.", ox.Spec("true")),
		), ox.Sub(
			ox.Banner("Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name is optional."),
			ox.Usage("logs", "Print the logs for a container in a pod"),
			ox.Spec("[-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [options]"),
			ox.Example("\n  # Return snapshot logs from pod nginx with only one container\n  kubectl logs nginx\n  \n  # Return snapshot logs from pod nginx with multi containers\n  kubectl logs nginx --all-containers=true\n  \n  # Return snapshot logs from all containers in pods defined by label app=nginx\n  kubectl logs -l app=nginx --all-containers=true\n  \n  # Return snapshot of previous terminated ruby container logs from pod web-1\n  kubectl logs -p -c ruby web-1\n  \n  # Begin streaming the logs of the ruby container in pod web-1\n  kubectl logs -f -c ruby web-1\n  \n  # Begin streaming the logs from all containers in pods defined by label app=nginx\n  kubectl logs -f -l app=nginx --all-containers=true\n  \n  # Display only the most recent 20 lines of output in pod nginx\n  kubectl logs --tail=20 nginx\n  \n  # Show all logs from pod nginx written in the last hour\n  kubectl logs --since=1h nginx\n  \n  # Show logs from a kubelet with an expired serving certificate\n  kubectl logs --insecure-skip-tls-verify-backend nginx\n  \n  # Return snapshot logs from first container of a job named hello\n  kubectl logs job/hello\n  \n  # Return snapshot logs from container nginx-1 of a deployment named nginx\n  kubectl logs deployment/nginx -c nginx-1"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all-containers", "Get all containers' logs in the pod(s).", ox.Spec("false")).
				Bool("all-pods", "Get logs from all pod(s). Sets prefix to true.", ox.Spec("false")).
				String("container", "Print the logs of this container", ox.Short("c")).
				Bool("follow", "Specify if the logs should be streamed.", ox.Spec("false"), ox.Short("f")).
				Bool("ignore-errors", "If watching / following pod logs, allow for any errors that occur to be non-fatal", ox.Spec("false")).
				Bool("insecure-skip-tls-verify-backend", "Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker could provide invalid log content back. You might want to use this if your kubelet serving certificates have expired.", ox.Spec("false")).
				Int("limit-bytes", "Maximum bytes of logs to return. Defaults to no limit.", ox.Default("0")).
				Int("max-log-requests", "Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5.", ox.Default("5")).
				Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("20s")).
				Bool("prefix", "Prefix each log line with the log source (pod name and container name)", ox.Spec("false")).
				Bool("previous", "If true, print the logs for the previous instance of the container in a pod if it exists.", ox.Spec("false"), ox.Short("p")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Duration("since", "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.", ox.Default("0s")).
				String("since-time", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.").
				Int("tail", "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.", ox.Default("-1")).
				Bool("timestamps", "Include timestamps on each line in the log output", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Attach to a process that is already running inside an existing container."),
			ox.Usage("attach", "Attach to a running container"),
			ox.Spec("(POD | TYPE/NAME) -c CONTAINER [options]"),
			ox.Example("\n  # Get output from running pod mypod; use the 'kubectl.kubernetes.io/default-container' annotation\n  # for selecting the container to be attached or the first container in the pod will be chosen\n  kubectl attach mypod\n  \n  # Get output from ruby-container from pod mypod\n  kubectl attach mypod -c ruby-container\n  \n  # Switch to raw terminal mode; sends stdin to 'bash' in ruby-container from pod mypod\n  # and sends stdout/stderr from 'bash' back to the client\n  kubectl attach mypod -c ruby-container -i -t\n  \n  # Get output from the first pod of a replica set named nginx\n  kubectl attach rs/nginx"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("container", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen", ox.Short("c")).
				Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("1m0s")).
				Bool("quiet", "Only print output from the remote session", ox.Spec("false"), ox.Short("q")).
				Bool("stdin", "Pass stdin to the container", ox.Spec("false"), ox.Short("i")).
				Bool("tty", "Stdin is a TTY", ox.Spec("false"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Execute a command in a container."),
			ox.Usage("exec", "Execute a command in a container"),
			ox.Spec("(POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...] [options]"),
			ox.Example("\n  # Get output from running the 'date' command from pod mypod, using the first container by default\n  kubectl exec mypod -- date\n  \n  # Get output from running the 'date' command in ruby-container from pod mypod\n  kubectl exec mypod -c ruby-container -- date\n  \n  # Switch to raw terminal mode; sends stdin to 'bash' in ruby-container from pod mypod\n  # and sends stdout/stderr from 'bash' back to the client\n  kubectl exec mypod -c ruby-container -i -t -- bash -il\n  \n  # List contents of /usr from the first container of pod mypod and sort by modification time\n  # If the command you want to execute in the pod has any flags in common (e.g. -i),\n  # you must use two dashes (--) to separate your command's flags/arguments\n  # Also note, do not surround your command and its flags/arguments with quotes\n  # unless that is how you would execute it normally (i.e., do ls -t /usr, not \"ls -t /usr\")\n  kubectl exec mypod -i -t -- ls -t /usr\n  \n  # Get output from running 'date' command from the first pod of the deployment mydeployment, using the first container by default\n  kubectl exec deploy/mydeployment -- date\n  \n  # Get output from running 'date' command from the first pod of the service myservice, using the first container by default\n  kubectl exec svc/myservice -- date"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("container", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen", ox.Short("c")).
				Slice("filename", "to use to exec into the resource", ox.Elem(ox.StringT), ox.Short("f")).
				Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("1m0s")).
				Bool("quiet", "Only print output from the remote session", ox.Spec("false"), ox.Short("q")).
				Bool("stdin", "Pass stdin to the container", ox.Spec("false"), ox.Short("i")).
				Bool("tty", "Stdin is a TTY", ox.Spec("false"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Forward one or more local ports to a pod.\n\n Use resource type/name such as deployment/mydeployment to select a pod. Resource type defaults to 'pod' if omitted.\n\n If there are multiple pods matching the criteria, a pod will be selected automatically. The forwarding session ends when the selected pod terminates, and a rerun of the command is needed to resume forwarding."),
			ox.Usage("port-forward", "Forward one or more local ports to a pod"),
			ox.Spec("TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N]"),
			ox.Example("\n  # Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in the pod\n  kubectl port-forward pod/mypod 5000 6000\n  \n  # Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in a pod selected by the deployment\n  kubectl port-forward deployment/mydeployment 5000 6000\n  \n  # Listen on port 8443 locally, forwarding to the targetPort of the service's port named \"https\" in a pod selected by the service\n  kubectl port-forward service/myservice 8443:https\n  \n  # Listen on port 8888 locally, forwarding to 5000 in the pod\n  kubectl port-forward pod/mypod 8888:5000\n  \n  # Listen on port 8888 on all addresses, forwarding to 5000 in the pod\n  kubectl port-forward --address 0.0.0.0 pod/mypod 8888:5000\n  \n  # Listen on port 8888 on localhost and selected IP, forwarding to 5000 in the pod\n  kubectl port-forward --address localhost,10.19.21.23 pod/mypod 8888:5000\n  \n  # Listen on a random port locally, forwarding to 5000 in the pod\n  kubectl port-forward pod/mypod :5000"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Slice("address", "Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When localhost is supplied, kubectl will try to bind on both 127.0.0.1 and ::1 and will fail if neither of these addresses are available to bind.", ox.Elem(ox.StringT), ox.Default("localhost")).
				Duration("pod-running-timeout", "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running", ox.Default("1m0s")),
		), ox.Sub(
			ox.Banner("Creates a proxy server or application-level gateway between localhost and the Kubernetes API server. It also allows serving static content over specified HTTP path. All incoming data enters through one port and gets forwarded to the remote Kubernetes API server port, except for the path matching the static content path."),
			ox.Usage("proxy", "Run a proxy to the Kubernetes API server"),
			ox.Spec("[--port=PORT] [--www=static-dir] [--www-prefix=prefix] [--api-prefix=prefix] [options]"),
			ox.Example("\n  # To proxy all of the Kubernetes API and nothing else\n  kubectl proxy --api-prefix=/\n  \n  # To proxy only part of the Kubernetes API and also some static files\n  # You can get pods info with 'curl localhost:8001/api/v1/pods'\n  kubectl proxy --www=/my/files --www-prefix=/static/ --api-prefix=/api/\n  \n  # To proxy the entire Kubernetes API at a different root\n  # You can get pods info with 'curl localhost:8001/custom/api/v1/pods'\n  kubectl proxy --api-prefix=/custom/\n  \n  # Run a proxy to the Kubernetes API server on port 8011, serving static content from ./local/www/\n  kubectl proxy --port=8011 --www=./local/www/\n  \n  # Run a proxy to the Kubernetes API server on an arbitrary local port\n  # The chosen port for the server will be output to stdout\n  kubectl proxy --port=0\n  \n  # Run a proxy to the Kubernetes API server, changing the API prefix to k8s-api\n  # This makes e.g. the pods API available at localhost:8001/k8s-api/v1/pods/\n  kubectl proxy --api-prefix=/k8s-api"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("accept-hosts", "Regular expression for hosts that the proxy should accept.", ox.Spec("regexp"), ox.Default("^localhost$,^127\\.0\\.0\\.1$,^\\[::1\\]$")).
				String("accept-paths", "Regular expression for paths that the proxy should accept.", ox.Spec("regexp"), ox.Default("^.*")).
				Addr("address", "The IP address on which to serve on.", ox.Default("127.0.0.1")).
				String("api-prefix", "Prefix to serve the proxied API under.", ox.Spec("path"), ox.Default("/")).
				Bool("append-server-path", "If true, enables automatic path appending of the kube context server path to each request.", ox.Spec("false")).
				Bool("disable-filter", "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port.", ox.Spec("false")).
				Duration("keepalive", "keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keepalive.", ox.Default("0s")).
				Int("port", "The port on which to run the proxy. Set to 0 to pick a random port.", ox.Default("8001"), ox.Short("p")).
				String("reject-methods", "Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT,PATCH').", ox.Spec("regexp"), ox.Default("^$")).
				String("reject-paths", "Regular expression for paths that the proxy should reject. Paths specified here will be rejected even accepted by --accept-paths.", ox.Spec("regexp"), ox.Default("^/api/.*/pods/.*/exec,^/api/.*/pods/.*/attach")).
				String("unix-socket", "Unix socket on which to run the proxy.", ox.Short("u")).
				String("www", "Also serve static files from the given directory under the specified prefix.", ox.Short("w")).
				String("www-prefix", "Prefix to serve static files under, if static file directory is specified.", ox.Spec("path"), ox.Default("/static/"), ox.Short("P")),
		), ox.Sub(
			ox.Banner("Copy files and directories to and from containers."),
			ox.Usage("cp", "Copy files and directories to and from containers"),
			ox.Spec("<file-spec-src> <file-spec-dest> [options]"),
			ox.Example("\n  # !!!Important Note!!!\n  # Requires that the 'tar' binary is present in your container\n  # image.  If 'tar' is not present, 'kubectl cp' will fail.\n  #\n  # For advanced use cases, such as symlinks, wildcard expansion or\n  # file mode preservation, consider using 'kubectl exec'.\n  \n  # Copy /tmp/foo local file to /tmp/bar in a remote pod in namespace <some-namespace>\n  tar cf - /tmp/foo | kubectl exec -i -n <some-namespace> <some-pod> -- tar xf - -C /tmp/bar\n  \n  # Copy /tmp/foo from a remote pod to /tmp/bar locally\n  kubectl exec -n <some-namespace> <some-pod> -- tar cf - /tmp/foo | tar xf - -C /tmp/bar\n  \n  # Copy /tmp/foo_dir local directory to /tmp/bar_dir in a remote pod in the default namespace\n  kubectl cp /tmp/foo_dir <some-pod>:/tmp/bar_dir\n  \n  # Copy /tmp/foo local file to /tmp/bar in a remote pod in a specific container\n  kubectl cp /tmp/foo <some-pod>:/tmp/bar -c <specific-container>\n  \n  # Copy /tmp/foo local file to /tmp/bar in a remote pod in namespace <some-namespace>\n  kubectl cp /tmp/foo <some-namespace>/<some-pod>:/tmp/bar\n  \n  # Copy /tmp/foo from a remote pod to /tmp/bar locally\n  kubectl cp <some-namespace>/<some-pod>:/tmp/foo /tmp/bar"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("container", "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen", ox.Short("c")).
				Bool("no-preserve", "The copied file/directory's ownership and permissions will not be preserved in the container", ox.Spec("false")).
				Int("retries", "Set number of retries to complete a copy operation from a container. Specify 0 to disable or any negative value for infinite retrying. The default is 0 (no retry).", ox.Default("0")),
		), ox.Sub(
			ox.Banner("Inspect authorization."),
			ox.Usage("auth", "Inspect authorization"),
			ox.Spec("[flags] [options]"),
			ox.Section(4),
			ox.Footer("Use \"kubectl auth <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Check whether an action is allowed.\n\n VERB is a logical Kubernetes API verb like 'get', 'list', 'watch', 'delete', etc. TYPE is a Kubernetes resource. Shortcuts and groups will be resolved. NONRESOURCEURL is a partial URL that starts with \"/\". NAME is the name of a particular Kubernetes resource. This command pairs nicely with impersonation. See --as global flag."),
				ox.Usage("can-i", "Check whether an action is allowed"),
				ox.Spec("VERB [TYPE | TYPE/NAME | NONRESOURCEURL] [options]"),
				ox.Example("\n  # Check to see if I can create pods in any namespace\n  kubectl auth can-i create pods --all-namespaces\n  \n  # Check to see if I can list deployments in my current namespace\n  kubectl auth can-i list deployments.apps\n  \n  # Check to see if service account \"foo\" of namespace \"dev\" can list pods in the namespace \"prod\"\n  # You must be allowed to use impersonation for the global option \"--as\"\n  kubectl auth can-i list pods --as=system:serviceaccount:dev:foo -n prod\n  \n  # Check to see if I can do everything in my current namespace (\"*\" means all)\n  kubectl auth can-i '*' '*'\n  \n  # Check to see if I can get the job named \"bar\" in namespace \"foo\"\n  kubectl auth can-i list jobs.batch/bar -n foo\n  \n  # Check to see if I can read pod logs\n  kubectl auth can-i get pods --subresource=log\n  \n  # Check to see if I can access the URL /logs/\n  kubectl auth can-i get /logs/\n  \n  # Check to see if I can approve certificates.k8s.io\n  kubectl auth can-i approve certificates.k8s.io\n  \n  # List all allowed actions in namespace \"foo\"\n  kubectl auth can-i --list --namespace=foo"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all-namespaces", "If true, check the specified action in all namespaces.", ox.Spec("false"), ox.Short("A")).
					Bool("list", "If true, prints all allowed actions.", ox.Spec("false")).
					Bool("no-headers", "If true, prints allowed actions without headers", ox.Spec("false")).
					Bool("quiet", "If true, suppress output and just return the exit code.", ox.Spec("false"), ox.Short("q")).
					String("subresource", "SubResource such as pod/log or deployment/scale"),
			), ox.Sub(
				ox.Banner("Reconciles rules for RBAC role, role binding, cluster role, and cluster role binding objects.\n\n Missing objects are created, and the containing namespace is created for namespaced objects, if required.\n\n Existing roles are updated to include the permissions in the input objects, and remove extra permissions if --remove-extra-permissions is specified.\n\n Existing bindings are updated to include the subjects in the input objects, and remove extra subjects if --remove-extra-subjects is specified.\n\n This is preferred to 'apply' for RBAC resources so that semantically-aware merging of rules and subjects is done."),
				ox.Usage("reconcile", "Reconciles rules for RBAC role, role binding, cluster role, and cluster role binding objects"),
				ox.Spec("-f FILENAME [options]"),
				ox.Example("\n  # Reconcile RBAC resources from a file\n  kubectl auth reconcile -f my-rbac-rules.yaml"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					Slice("filename", "Filename, directory, or URL to files identifying the resource to reconcile.", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("remove-extra-permissions", "If true, removes extra permissions added to roles", ox.Spec("false")).
					Bool("remove-extra-subjects", "If true, removes extra subjects added to rolebindings", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("Check who you are and your attributes (groups, extra).\n\n        This command is helpful to get yourself aware of the current user attributes,\n        especially when dynamic authentication, e.g., token webhook, auth proxy, or OIDC provider,\n        is enabled in the Kuberne"),
				ox.Usage("whoami", "Experimental: Check self subject attributes"),
				ox.Spec("[options]"),
				ox.Example("\n  # Get your subject attributes\n  kubectl auth whoami\n  \n  # Get your subject attributes in JSON format\n  kubectl auth whoami -o json"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			)), ox.Sub(
			ox.Banner("Debug cluster resources using interactive debugging containers.\n\n 'debug' provides automation for common debugging tasks for cluster objects identified by resource and name. Pods will be used by default if no resource is specified.\n\n The action taken by 'debug' varies depending on what resource is specified. Supported actions include:\n\n  *  Workload: Create a copy of an existing pod with certain attributes changed, for example changing the image tag to a new version.\n  *  Workload: Add an ephemeral container to an already running pod, for example to add debugging utilities without restarting the pod.\n  *  Node: Create a new pod that runs in the node's host namespaces and can access the node's filesystem."),
			ox.Usage("debug", "Create debugging sessions for troubleshooting workloads and nodes"),
			ox.Spec("(POD | TYPE[[.VERSION].GROUP]/NAME) [ -- COMMAND [args...] ] [options]"),
			ox.Example("\n  # Create an interactive debugging session in pod mypod and immediately attach to it.\n  kubectl debug mypod -it --image=busybox\n  \n  # Create an interactive debugging session for the pod in the file pod.yaml and immediately attach to it.\n  # (requires the EphemeralContainers feature to be enabled in the cluster)\n  kubectl debug -f pod.yaml -it --image=busybox\n  \n  # Create a debug container named debugger using a custom automated debugging image.\n  kubectl debug --image=myproj/debug-tools -c debugger mypod\n  \n  # Create a copy of mypod adding a debug container and attach to it\n  kubectl debug mypod -it --image=busybox --copy-to=my-debugger\n  \n  # Create a copy of mypod changing the command of mycontainer\n  kubectl debug mypod -it --copy-to=my-debugger --container=mycontainer -- sh\n  \n  # Create a copy of mypod changing all container images to busybox\n  kubectl debug mypod --copy-to=my-debugger --set-image=*=busybox\n  \n  # Create a copy of mypod adding a debug container and changing container images\n  kubectl debug mypod -it --copy-to=my-debugger --image=debian --set-image=app=app:debug,sidecar=sidecar:debug\n  \n  # Create an interactive debugging session on a node and immediately attach to it.\n  # The container will run in the host namespaces and the host's filesystem will be mounted at /host\n  kubectl debug node/mynode -it --image=busybox"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("arguments-only", "If specified, everything after -- will be passed to the new container as Args instead of Command.", ox.Spec("false")).
				Bool("attach", "If true, wait for the container to start running, and then attach as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true.", ox.Spec("false")).
				String("container", "Container name to use for debug container.", ox.Short("c")).
				String("copy-to", "Create a copy of the target Pod with this name.").
				String("custom", "Path to a JSON or YAML file containing a partial container spec to customize built-in debug profiles.").
				Slice("env", "Environment variables to set in the container.", ox.Elem(ox.StringT)).
				Slice("filename", "identifying the resource to debug", ox.Elem(ox.StringT), ox.Short("f")).
				String("image", "Container image to use for debug container.").
				String("image-pull-policy", "The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server.").
				Bool("keep-annotations", "If true, keep the original pod annotations.(This flag only works when used with '--copy-to')", ox.Spec("false")).
				Bool("keep-init-containers", "Run the init containers for the pod. Defaults to true.(This flag only works when used with '--copy-to')", ox.Spec("true")).
				Bool("keep-labels", "If true, keep the original pod labels.(This flag only works when used with '--copy-to')", ox.Spec("false")).
				Bool("keep-liveness", "If true, keep the original pod liveness probes.(This flag only works when used with '--copy-to')", ox.Spec("false")).
				Bool("keep-readiness", "If true, keep the original pod readiness probes.(This flag only works when used with '--copy-to')", ox.Spec("false")).
				Bool("keep-startup", "If true, keep the original startup probes.(This flag only works when used with '--copy-to')", ox.Spec("false")).
				String("profile", "Options are \"legacy\", \"general\", \"baseline\", \"netadmin\", \"restricted\" or \"sysadmin\".", ox.Default("legacy")).
				Bool("quiet", "If true, suppress informational messages.", ox.Spec("false"), ox.Short("q")).
				Bool("replace", "When used with '--copy-to', delete the original Pod.", ox.Spec("false")).
				Bool("same-node", "When used with '--copy-to', schedule the copy of target Pod on the same node.", ox.Spec("false")).
				Slice("set-image", "When used with '--copy-to', a list of name=image pairs for changing container images, similar to how 'kubectl set image' works.", ox.Elem(ox.StringT)).
				Bool("share-processes", "When used with '--copy-to', enable process namespace sharing in the copy.", ox.Spec("true")).
				Bool("stdin", "Keep stdin open on the container(s) in the pod, even if nothing is attached.", ox.Spec("false"), ox.Short("i")).
				String("target", "When using an ephemeral container, target processes in this container name.").
				Bool("tty", "Allocate a TTY for the debugging container.", ox.Spec("false"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Display events.\n\n Prints a table of the most important information about events. You can request events for a namespace, for all namespace, or filtered to only those pertaining to a specified resource."),
			ox.Usage("events", "List events"),
			ox.Spec("[(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file] [--for TYPE/NAME] [--watch] [--types=Normal,Warning] [options]"),
			ox.Example("\n  # List recent events in the default namespace\n  kubectl events\n  \n  # List recent events in all namespaces\n  kubectl events --all-namespaces\n  \n  # List recent events for the specified pod, then wait for more events and list them as they arrive\n  kubectl events --for pod/web-pod-13je7 --watch\n  \n  # List recent events in YAML format\n  kubectl events -oyaml\n  \n  # List recent only events of type 'Warning' or 'Normal'\n  kubectl events --types=Warning,Normal"),
			ox.Section(4),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				Int("chunk-size", "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.", ox.Default("500")).
				String("for", "Filter events to only those pertaining to the specified resource.").
				Bool("no-headers", "When using the default output format, don't print headers.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Slice("types", "Output only events of given types.", ox.Elem(ox.StringT)).
				Bool("watch", "After listing the requested events, watch for more events.", ox.Spec("false"), ox.Short("w")),
		), ox.Sub(
			ox.Banner("Diff configurations specified by file name or stdin between the current online configuration, and the configuration as it would be if applied.\n\n The output is always YAML.\n\n KUBECTL_EXTERNAL_DIFF environment variable can be used to select your own diff command. Users can use external commands with params too, example: KUBECTL_EXTERNAL_DIFF=\"colordiff -N -u\"\n\n By default, the \"diff\" command available in your path will be run with the \"-u\" (unified diff) and \"-N\" (treat absent files as empty) options.\n\n Exit status: 0 No differences were found. 1 Differences were found. >1 Kubectl or diff failed with an error.\n\n Note: KUBECTL_EXTERNAL_DIFF, if used, is expected to follow that convention."),
			ox.Usage("diff", "Diff the live version against a would-be applied version"),
			ox.Spec("-f FILENAME [options]"),
			ox.Example("\n  # Diff resources included in pod.json\n  kubectl diff -f pod.json\n  \n  # Diff file read from stdin\n  cat service.yaml | kubectl diff -f -"),
			ox.Section(5),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Int("concurrency", "Number of objects to process in parallel when diffing against the live version. Larger number = faster, but more memory, I/O and CPU over that shorter period of time.", ox.Default("1")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-client-side-apply")).
				Slice("filename", "Filename, directory, or URL to files contains the configuration to diff", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("force-conflicts", "If true, server-side apply will force the changes against conflicts.", ox.Spec("false")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("prune", "Include resources that would be deleted by pruning. Can be used with -l and default shows all resources would be pruned", ox.Spec("false")).
				Slice("prune-allowlist", "Overwrite the default allowlist with <group/version/kind> for --prune", ox.Elem(ox.StringT)).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("server-side", "If true, apply runs in the server instead of the client.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, include managed fields in the diff.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Apply a configuration to a resource by file name or stdin. The resource name must be specified. This resource will be created if it doesn't exist yet. To use 'apply', always create the resource initially with either 'apply' or 'create --save-config'.\n\n JSON and YAML formats are accepted.\n\n Alpha Disclaimer: the --prune functionality is not yet complete. Do not use unless you are aware of what the current state is. See https://issues.k8s.io/34274."),
			ox.Usage("apply", "Apply a configuration to a resource by file name or stdin"),
			ox.Spec("(-f FILENAME | -k DIRECTORY) [options]"),
			ox.Example("\n  # Apply the configuration in pod.json to a pod\n  kubectl apply -f ./pod.json\n  \n  # Apply resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml\n  kubectl apply -k dir/\n  \n  # Apply the JSON passed into stdin to a pod\n  cat pod.json | kubectl apply -f -\n  \n  # Apply the configuration from all files that end with '.json'\n  kubectl apply -f '*.json'\n  \n  # Note: --prune is still in Alpha\n  # Apply the configuration in manifest.yaml that matches label app=nginx and delete all other resources that are not in the file and match label app=nginx\n  kubectl apply --prune -f manifest.yaml -l app=nginx\n  \n  # Apply the configuration in manifest.yaml and delete all the other config maps that are not in the file\n  kubectl apply --prune -f manifest.yaml --all --prune-allowlist=core/v1/ConfigMap"),
			ox.Section(5),
			ox.Footer("Use \"kubectl apply <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Edit the latest last-applied-configuration annotations of resources from the default editor.\n\n The edit-last-applied command allows you to directly edit any API resource you can retrieve via the command-line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows. You can edit multiple objects, although changes are applied one at a time. The command accepts file names as well as command-line arguments, although the files you point to must be previously saved versions of resources.\n\n The default format is YAML. To edit in JSON, specify \"-o json\".\n\n The flag --windows-line-endings can be used to force Windows line endings, otherwise the default for your operating system will be used.\n\n In the event an error occurs while updating, a temporary file will be created on disk that contains your unapplied changes. The most common error when updating a resource is another editor changing the resource on the server. When this occurs, you will have to apply your changes to the newer version of the resource, or update your temporary saved copy to include the latest resource version."),
				ox.Usage("edit-last-applied", "Edit latest last-applied-configuration annotations of a resource/object"),
				ox.Spec("(RESOURCE/NAME | -f FILENAME) [options]"),
				ox.Example("\n  # Edit the last-applied-configuration annotations by type/name in YAML\n  kubectl apply edit-last-applied deployment/nginx\n  \n  # Edit the last-applied-configuration annotations by file in JSON\n  kubectl apply edit-last-applied -f deploy.yaml -o json"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-client-side-apply")).
					Slice("filename", "Filename, directory, or URL to files to use to edit the resource", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
					String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
					Bool("windows-line-endings", "Defaults to the line ending native to your platform.", ox.Spec("false")),
			), ox.Sub(
				ox.Banner("Set the latest last-applied-configuration annotations by setting it to match the contents of a file. This results in the last-applied-configuration being updated as though 'kubectl apply -f<file> ' was run, without updating any other parts of the object."),
				ox.Usage("set-last-applied", "Set the last-applied-configuration annotation on a live object to match the contents of a file"),
				ox.Spec("-f FILENAME [options]"),
				ox.Example("\n  # Set the last-applied-configuration of a resource to match the contents of a file\n  kubectl apply set-last-applied -f deploy.yaml\n  \n  # Execute set-last-applied against each configuration file in a directory\n  kubectl apply set-last-applied -f path/\n  \n  # Set the last-applied-configuration of a resource to match the contents of a file; will create the annotation if it does not already exist\n  kubectl apply set-last-applied -f deploy.yaml --create-annotation=true"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Bool("create-annotation", "Will create 'last-applied-configuration' annotations if current objects doesn't have one", ox.Spec("false")).
					String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
					Slice("filename", "Filename, directory, or URL to files that contains the last-applied-configuration annotations", ox.Elem(ox.StringT), ox.Short("f")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			), ox.Sub(
				ox.Banner("View the latest last-applied-configuration annotations by type/name or file.\n\n The default output will be printed to stdout in YAML format. You can use the -o option to change the output format."),
				ox.Usage("view-last-applied", "View the latest last-applied-configuration annotations of a resource/object"),
				ox.Spec("(TYPE [NAME | -l label] | TYPE/NAME | -f FILENAME) [options]"),
				ox.Example("\n  # View the last-applied-configuration annotations by type/name in YAML\n  kubectl apply view-last-applied deployment/nginx\n  \n  # View the last-applied-configuration annotations by file in JSON\n  kubectl apply view-last-applied -f deploy.yaml -o json"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("all", "Select all resources in the namespace of the specified resource types", ox.Spec("false")).
					Slice("filename", "Filename, directory, or URL to files that contains the last-applied-configuration annotations", ox.Elem(ox.StringT), ox.Short("f")).
					String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
					String("output", "Output format. Must be one of (yaml, json)", ox.Default("yaml"), ox.Short("o")).
					Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
					String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")),
			), ox.Flags().
				Bool("all", "Select all resources in the namespace of the specified resource types.", ox.Spec("false")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("cascade", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.", ox.Default("background")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-client-side-apply")).
				Slice("filename", "The files that contain the configurations to apply.", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("force", "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.", ox.Spec("false")).
				Bool("force-conflicts", "If true, server-side apply will force the changes against conflicts.", ox.Spec("false")).
				Int("grace-period", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).", ox.Default("-1")).
				String("kustomize", "Process a kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("openapi-patch", "If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types.", ox.Spec("true")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("overwrite", "Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration", ox.Spec("true")).
				Bool("prune", "Automatically delete resource objects, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all.", ox.Spec("false")).
				Slice("prune-allowlist", "Overwrite the default allowlist with <group/version/kind> for --prune", ox.Elem(ox.StringT)).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("server-side", "If true, apply runs in the server instead of the client.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Duration("timeout", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object", ox.Default("0s")).
				String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
				Bool("wait", "If true, wait for resources to be gone before returning. This waits for finalizers.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Update fields of a resource using strategic merge patch, a JSON merge patch, or a JSON patch.\n\n JSON and YAML formats are accepted.\n\n Note: Strategic merge patch is not supported for custom resources."),
			ox.Usage("patch", "Update fields of a resource"),
			ox.Spec("(-f FILENAME | TYPE NAME) [-p PATCH|--patch-file FILE] [options]"),
			ox.Example("\n  # Partially update a node using a strategic merge patch, specifying the patch as JSON\n  kubectl patch node k8s-node-1 -p '{\"spec\":{\"unschedulable\":true}}'\n  \n  # Partially update a node using a strategic merge patch, specifying the patch as YAML\n  kubectl patch node k8s-node-1 -p $'spec:\\n unschedulable: true'\n  \n  # Partially update a node identified by the type and name specified in \"node.json\" using strategic merge patch\n  kubectl patch -f node.json -p '{\"spec\":{\"unschedulable\":true}}'\n  \n  # Update a container's image; spec.containers[*].name is required because it's a merge key\n  kubectl patch pod valid-pod -p '{\"spec\":{\"containers\":[{\"name\":\"kubernetes-serve-hostname\",\"image\":\"new image\"}]}}'\n  \n  # Update a container's image using a JSON patch with positional arrays\n  kubectl patch pod valid-pod --type='json' -p='[{\"op\": \"replace\", \"path\": \"/spec/containers/0/image\", \"value\":\"new image\"}]'\n  \n  # Update a deployment's replicas through the 'scale' subresource using a merge patch\n  kubectl patch deployment nginx-deployment --subresource='scale' --type='merge' -p '{\"spec\":{\"replicas\":2}}'"),
			ox.Section(5),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-patch")).
				Slice("filename", "Filename, directory, or URL to files identifying the resource to update", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("local", "If true, patch will operate on the content of the file, not the server-side resource.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				String("patch", "The patch to be applied to the resource JSON file.", ox.Short("p")).
				String("patch-file", "A file containing a patch to be applied to the resource.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("subresource", "If specified, patch will operate on the subresource of the requested object. Must be one of [status scale]. This flag is beta and may change in the future.").
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				String("type", "The type of patch being provided; one of [json merge strategic]", ox.Default("strategic")),
		), ox.Sub(
			ox.Banner("Replace a resource by file name or stdin.\n\n JSON and YAML formats are accepted. If replacing an existing resource, the complete resource spec must be provided. This can be obtained by\n\n        $ kubectl get TYPE NAME -o yaml"),
			ox.Usage("replace", "Replace a resource by file name or stdin"),
			ox.Spec("-f FILENAME [options]"),
			ox.Example("\n  # Replace a pod using the data in pod.json\n  kubectl replace -f ./pod.json\n  \n  # Replace a pod based on the JSON passed into stdin\n  cat pod.json | kubectl replace -f -\n  \n  # Update a single-container pod's image version (tag) to v4\n  kubectl get pod mypod -o yaml | sed 's/\\(image: myimage\\):.*$/\\1:v4/' | kubectl replace -f -\n  \n  # Force replace, delete and then re-create the resource\n  kubectl replace --force -f ./pod.json"),
			ox.Section(5),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("cascade", "Must be \"background\", \"orphan\", or \"foreground\". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.", ox.Default("background")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-replace")).
				Slice("filename", "The files that contain the configurations to replace.", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("force", "If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.", ox.Spec("false")).
				Int("grace-period", "Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).", ox.Default("-1")).
				String("kustomize", "Process a kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				String("raw", "Raw URI to PUT to the server.  Uses the transport specified by the kubeconfig file.").
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				Bool("save-config", "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.", ox.Spec("false")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("subresource", "If specified, replace will operate on the subresource of the requested object. Must be one of [status scale]. This flag is beta and may change in the future.").
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Duration("timeout", "The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object", ox.Default("0s")).
				String("validate", "Must be one of: strict (or true), warn, ignore (or false). \t\t\"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \t\t\"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \t\t\"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.", ox.Default("strict")).
				Bool("wait", "If true, wait for resources to be gone before returning. This waits for finalizers.", ox.Spec("false")),
		), ox.Sub(
			ox.Banner("Wait for a specific condition on one or many resources.\n\n The command takes multiple resources and waits until the specified condition is seen in the Status field of every given resource.\n\n Alternatively, the command can wait for the given set of resources to be deleted by providing the \"delete\" keyword as the value to the --for flag.\n\n A successful message will be printed to stdout indicating when the specified condition has been met. You can use -o option to change to output"),
			ox.Usage("wait", "Experimental: Wait for a specific condition on one or many resources"),
			ox.Spec("([-f FILENAME] | resource.group/resource.name | resource.group [(-l label | --all)]) [--for=delete|--for condition=available|--for=jsonpath='{}'[=value]] [options]"),
			ox.Example("\n  # Wait for the pod \"busybox1\" to contain the status condition of type \"Ready\"\n  kubectl wait --for=condition=Ready pod/busybox1\n  \n  # The default value of status condition is true; you can wait for other targets after an equal delimiter (compared after Unicode simple case folding, which is a more general form of case-insensitivity)\n  kubectl wait --for=condition=Ready=false pod/busybox1\n  \n  # Wait for the pod \"busybox1\" to contain the status phase to be \"Running\"\n  kubectl wait --for=jsonpath='{.status.phase}'=Running pod/busybox1\n  \n  # Wait for pod \"busybox1\" to be Ready\n  kubectl wait --for='jsonpath={.status.conditions[?(@.type==\"Ready\")].status}=True' pod/busybox1\n  \n  # Wait for the service \"loadbalancer\" to have ingress\n  kubectl wait --for=jsonpath='{.status.loadBalancer.ingress}' service/loadbalancer\n  \n  # Wait for the pod \"busybox1\" to be deleted, with a timeout of 60s, after having issued the \"delete\" command\n  kubectl delete pod/busybox1\n  kubectl wait --for=delete pod/busybox1 --timeout=60s"),
			ox.Section(5),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Select all resources in the namespace of the specified resource types", ox.Spec("false")).
				Bool("all-namespaces", "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.", ox.Spec("false"), ox.Short("A")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
				Slice("filename", "identifying the resource.", ox.Elem(ox.StringT), ox.Short("f")).
				String("for", "The condition to wait on: [delete|condition=condition-name[=condition-value]|jsonpath='{JSONPath expression}'=[JSONPath value]]. The default condition-value is true.  Condition values are compared after Unicode simple case folding, which is a more general form of case-insensitivity.").
				Bool("local", "If true, annotation will NOT contact api-server but run locally.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("true"), ox.Short("R")).
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].").
				Duration("timeout", "The length of time to wait before giving up.  Zero means check once and don't wait, negative means wait for a week.", ox.Default("30s")),
		), ox.Sub(
			ox.Banner("Build a set of KRM resources using a 'kustomization.yaml' file. The DIR argument must be a path to a directory containing 'kustomization.yaml', or a git repository URL with a path suffix specifying same with respect to the repository root. If DIR is omitted, '.' is assumed."),
			ox.Usage("kustomize", "Build a kustomization target from a directory or URL"),
			ox.Spec("DIR [flags] [options]"),
			ox.Example("\n  # Build the current working directory\n  kubectl kustomize\n  \n  # Build some shared configuration directory\n  kubectl kustomize /home/config/production\n  \n  # Build from github\n  kubectl kustomize https://github.com/kubernetes-sigs/kustomize.git/examples/helloWorld?ref=v1.0.6"),
			ox.Section(5),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("as-current-user", "use the uid and gid of the command executor to run the function in the container", ox.Spec("false")).
				Bool("enable-alpha-plugins", "enable kustomize plugins", ox.Spec("false")).
				Bool("enable-helm", "Enable use of the Helm chart inflator generator.", ox.Spec("false")).
				Slice("env", "a list of environment variables to be used by functions", ox.Elem(ox.StringT), ox.Short("e")).
				Slice("helm-api-versions", "Kubernetes api versions used by Helm for Capabilities.APIVersions", ox.Elem(ox.StringT)).
				String("helm-command", "helm command (path to executable)", ox.Default("helm")).
				String("helm-kube-version", "Kubernetes version used by Helm for Capabilities.KubeVersion").
				String("load-restrictor", "if set to 'LoadRestrictionsNone', local kustomizations may load files from outside their root. This does, however, break the relocatability of the kustomization.", ox.Default("LoadRestrictionsRootOnly")).
				Slice("mount", "a list of storage options read from the filesystem", ox.Elem(ox.StringT)).
				Bool("network", "enable network access for functions that declare it", ox.Spec("false")).
				String("network-name", "the docker network to run the container in", ox.Default("bridge")).
				String("output", "If specified, write output to this path.", ox.Short("o")),
		), ox.Sub(
			ox.Banner("Update the labels on a resource.\n\n  *  A label key and value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters each.\n  *  Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app.\n  *  If --overwrite is true, then existing labels can be overwritten, otherwise attempting to overwrite a label will result in an error.\n  *  If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used."),
			ox.Usage("label", "Update the labels on a resource"),
			ox.Spec("[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]"),
			ox.Example("\n  # Update pod 'foo' with the label 'unhealthy' and the value 'true'\n  kubectl label pods foo unhealthy=true\n  \n  # Update pod 'foo' with the label 'status' and the value 'unhealthy', overwriting any existing value\n  kubectl label --overwrite pods foo status=unhealthy\n  \n  # Update all pods in the namespace\n  kubectl label pods --all status=unhealthy\n  \n  # Update a pod identified by the type and name in \"pod.json\"\n  kubectl label -f pod.json status=unhealthy\n  \n  # Update pod 'foo' only if the resource is unchanged from version 1\n  kubectl label pods foo status=unhealthy --resource-version=1\n  \n  # Update pod 'foo' by removing a label named 'bar' if it exists\n  # Does not require the --overwrite flag\n  kubectl label pods foo bar-"),
			ox.Section(6),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Select all resources, in the namespace of the specified resource types", ox.Spec("false")).
				Bool("all-namespaces", "If true, check the specified action in all namespaces.", ox.Spec("false"), ox.Short("A")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-label")).
				String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
				Slice("filename", "Filename, directory, or URL to files identifying the resource to update the labels", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("list", "If true, display the labels for a given resource.", ox.Spec("false")).
				Bool("local", "If true, label will NOT contact api-server but run locally.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("overwrite", "If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labels.", ox.Spec("false")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("resource-version", "If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.").
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
		), ox.Sub(
			ox.Banner("Update the annotations on one or more resources.\n\n All Kubernetes objects support the ability to store additional data with the object as annotations. Annotations are key/value pairs that can be larger than labels and include arbitrary string values such as structured JSON. Tools and system extensions may use annotations to store their own data.\n\n Attempting to set an annotation that already exists will fail unless --overwrite is set. If --resource-version is specified and does not match the current resource version on the server the command will fail.\n\nUse \"kubectl api-resources\" for a complete list of supported resources."),
			ox.Usage("annotate", "Update the annotations on a resource"),
			ox.Spec("[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]"),
			ox.Example("\n  # Update pod 'foo' with the annotation 'description' and the value 'my frontend'\n  # If the same annotation is set multiple times, only the last value will be applied\n  kubectl annotate pods foo description='my frontend'\n  \n  # Update a pod identified by type and name in \"pod.json\"\n  kubectl annotate -f pod.json description='my frontend'\n  \n  # Update pod 'foo' with the annotation 'description' and the value 'my frontend running nginx', overwriting any existing value\n  kubectl annotate --overwrite pods foo description='my frontend running nginx'\n  \n  # Update all pods in the namespace\n  kubectl annotate pods --all description='my frontend running nginx'\n  \n  # Update pod 'foo' only if the resource is unchanged from version 1\n  kubectl annotate pods foo description='my frontend running nginx' --resource-version=1\n  \n  # Update pod 'foo' by removing an annotation named 'description' if it exists\n  # Does not require the --overwrite flag\n  kubectl annotate pods foo description-"),
			ox.Section(6),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				Bool("all", "Select all resources, in the namespace of the specified resource types.", ox.Spec("false")).
				Bool("all-namespaces", "If true, check the specified action in all namespaces.", ox.Spec("false"), ox.Short("A")).
				Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
				String("dry-run", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.", ox.Default("none")).
				String("field-manager", "Name of the manager used to track field ownership.", ox.Default("$APPNAME-annotate")).
				String("field-selector", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.").
				Slice("filename", "Filename, directory, or URL to files identifying the resource to update the annotation", ox.Elem(ox.StringT), ox.Short("f")).
				String("kustomize", "Process the kustomization directory. This flag can't be used together with -f or -R.", ox.Short("k")).
				Bool("list", "If true, display the annotations for a given resource.", ox.Spec("false")).
				Bool("local", "If true, annotation will NOT contact api-server but run locally.", ox.Spec("false")).
				String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Short("o")).
				Bool("overwrite", "If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite existing annotations.", ox.Spec("false")).
				Bool("recursive", "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.", ox.Spec("false"), ox.Short("R")).
				String("resource-version", "If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.").
				String("selector", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.", ox.Short("l")).
				Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
				String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
		), ox.Sub(
			ox.Banner("Print the supported API resources on the server."),
			ox.Usage("api-resources", "Print the supported API resources on the server"),
			ox.Spec("[flags] [options]"),
			ox.Example("\n  # Print the supported API resources\n  kubectl api-resources\n  \n  # Print the supported API resources with more information\n  kubectl api-resources -o wide\n  \n  # Print the supported API resources sorted by a column\n  kubectl api-resources --sort-by=name\n  \n  # Print the supported namespaced resources\n  kubectl api-resources --namespaced=true\n  \n  # Print the supported non-namespaced resources\n  kubectl api-resources --namespaced=false\n  \n  # Print the supported API resources with a specific APIGroup\n  kubectl api-resources --api-group=rbac.authorization.k8s.io"),
			ox.Section(7),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Flags().
				String("api-group", "Limit to resources in the specified API group.").
				Bool("cached", "Use the cached list of resources if available.", ox.Spec("false")).
				Slice("categories", "Limit to resources that belong to the specified categories.", ox.Elem(ox.StringT)).
				Bool("namespaced", "If false, non-namespaced resources will be returned, otherwise returning namespaced resources by default.", ox.Spec("true")).
				Bool("no-headers", "When using the default or custom-column output format, don't print headers (default print headers).", ox.Spec("false")).
				String("output", "Output format. One of: (wide, name).", ox.Short("o")).
				String("sort-by", "If non-empty, sort list of resources using specified field. The field can be either 'name' or 'kind'.").
				Slice("verbs", "Limit to resources that support the specified verbs.", ox.Elem(ox.StringT)),
		), ox.Sub(
			ox.Banner("Print the supported API versions on the server, in the form of \"group/version\"."),
			ox.Usage("api-versions", "Print the supported API versions on the server, in the form of \"group/version\""),
			ox.Spec("[options]"),
			ox.Example("\n  # Print the supported API versions\n  kubectl api-versions"),
			ox.Section(7),
			ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
		), ox.Sub(
			ox.Banner("Modify kubeconfig files using subcommands like \"kubectl config set current-context my-context\".\n\n The loading order follows these rules:\n\n  1.  If the --kubeconfig flag is set, then only that file is loaded. The flag may only be set once and no merging takes place.\n  2.  If $KUBECONFIG environment variable is set, then it is used as a list of paths (normal path delimiting rules for your system). These paths are merged. When a value is modified, it is modified in the file that defines the stanza. When a value is created, it is created in the first file that exists. If no files in the chain exist, then it creates the last file in the list.\n  3.  Otherwise, ${HOME}/.kube/config is used and no merging takes place."),
			ox.Usage("config", "Modify kubeconfig files"),
			ox.Spec("SUBCOMMAND [options]"),
			ox.Section(7),
			ox.Footer("Use \"kubectl config <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("Display the current-context."),
				ox.Usage("current-context", "Display the current-context"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # Display the current-context\n  kubectl config current-context"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Delete the specified cluster from the kubeconfig."),
				ox.Usage("delete-cluster", "Delete the specified cluster from the kubeconfig"),
				ox.Spec("NAME [options]"),
				ox.Example("\n  # Delete the minikube cluster\n  kubectl config delete-cluster minikube"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Delete the specified context from the kubeconfig."),
				ox.Usage("delete-context", "Delete the specified context from the kubeconfig"),
				ox.Spec("NAME [options]"),
				ox.Example("\n  # Delete the context for the minikube cluster\n  kubectl config delete-context minikube"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Delete the specified user from the kubeconfig."),
				ox.Usage("delete-user", "Delete the specified user from the kubeconfig"),
				ox.Spec("NAME [options]"),
				ox.Example("\n  # Delete the minikube user\n  kubectl config delete-user minikube"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Display clusters defined in the kubeconfig."),
				ox.Usage("get-clusters", "Display clusters defined in the kubeconfig"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # List the clusters that kubectl knows about\n  kubectl config get-clusters"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Display one or many contexts from the kubeconfig file."),
				ox.Usage("get-contexts", "Describe one or many contexts"),
				ox.Spec("[(-o|--output=)name)] [options]"),
				ox.Example("\n  # List all the contexts in your kubeconfig file\n  kubectl config get-contexts\n  \n  # Describe one context in your kubeconfig file\n  kubectl config get-contexts my-context"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("no-headers", "When using the default or custom-column output format, don't print headers (default print headers).", ox.Spec("false")).
					String("output", "Output format. One of: (name).", ox.Short("o")),
			), ox.Sub(
				ox.Banner("Display users defined in the kubeconfig."),
				ox.Usage("get-users", "Display users defined in the kubeconfig"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # List the users that kubectl knows about\n  kubectl config get-users"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Renames a context from the kubeconfig file.\n\n CONTEXT_NAME is the context name that you want to change.\n\n NEW_NAME is the new name you want to set.\n\n Note: If the context being renamed is the 'current-context', this field will also be updated."),
				ox.Usage("rename-context", "Rename a context from the kubeconfig file"),
				ox.Spec("CONTEXT_NAME NEW_NAME [options]"),
				ox.Example("\n  # Rename the context 'old-name' to 'new-name' in your kubeconfig file\n  kubectl config rename-context old-name new-name"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Set an individual value in a kubeconfig file.\n\n PROPERTY_NAME is a dot delimited name where each token represents either an attribute name or a map key.  Map keys may not contain dots.\n\n PROPERTY_VALUE is the new value you want to set. Binary fields such as 'certificate-authority-data' expect a base64 encoded string unless the --set-raw-bytes flag is used.\n\n Specifying an attribute name that already exists will merge new fields on top of existing values."),
				ox.Usage("set", "Set an individual value in a kubeconfig file"),
				ox.Spec("PROPERTY_NAME PROPERTY_VALUE [options]"),
				ox.Example("\n  # Set the server field on the my-cluster cluster to https://1.2.3.4\n  kubectl config set clusters.my-cluster.server https://1.2.3.4\n  \n  # Set the certificate-authority-data field on the my-cluster cluster\n  kubectl config set clusters.my-cluster.certificate-authority-data $(echo \"cert_data_here\" | base64 -i -)\n  \n  # Set the cluster field in the my-context context to my-cluster\n  kubectl config set contexts.my-context.cluster my-cluster\n  \n  # Set the client-key-data field in the cluster-admin user using --set-raw-bytes option\n  kubectl config set users.cluster-admin.client-key-data cert_data_here --set-raw-bytes=true"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("set-raw-bytes", "When writing a []byte PROPERTY_VALUE, write the given string directly without base64 decoding.", ox.Spec("false")),
			), ox.Sub(
				ox.Banner("Set a cluster entry in kubeconfig.\n\n Specifying a name that already exists will merge new fields on top of existing values for those fields."),
				ox.Usage("set-cluster", "Set a cluster entry in kubeconfig"),
				ox.Spec("NAME [--server=server] [--certificate-authority=path/to/certificate/authority] [--insecure-skip-tls-verify=true] [--tls-server-name=example.com] [options]"),
				ox.Example("\n  # Set only the server field on the e2e cluster entry without touching other values\n  kubectl config set-cluster e2e --server=https://1.2.3.4\n  \n  # Embed certificate authority data for the e2e cluster entry\n  kubectl config set-cluster e2e --embed-certs --certificate-authority=~/.kube/e2e/kubernetes.ca.crt\n  \n  # Disable cert checking for the e2e cluster entry\n  kubectl config set-cluster e2e --insecure-skip-tls-verify=true\n  \n  # Set the custom TLS server name to use for validation for the e2e cluster entry\n  kubectl config set-cluster e2e --tls-server-name=my-cluster-name\n  \n  # Set the proxy URL for the e2e cluster entry\n  kubectl config set-cluster e2e --proxy-url=https://1.2.3.4"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					String("certificate-authority", "Path to certificate-authority file for the cluster entry in kubeconfig").
					Bool("embed-certs", "embed-certs for the cluster entry in kubeconfig", ox.Spec("false")).
					Bool("insecure-skip-tls-verify", "insecure-skip-tls-verify for the cluster entry in kubeconfig", ox.Spec("false")).
					String("proxy-url", "proxy-url for the cluster entry in kubeconfig").
					String("server", "server for the cluster entry in kubeconfig").
					String("tls-server-name", "tls-server-name for the cluster entry in kubeconfig"),
			), ox.Sub(
				ox.Banner("Set a context entry in kubeconfig.\n\n Specifying a name that already exists will merge new fields on top of existing values for those fields."),
				ox.Usage("set-context", "Set a context entry in kubeconfig"),
				ox.Spec("[NAME | --current] [--cluster=cluster_nickname] [--user=user_nickname] [--namespace=namespace] [options]"),
				ox.Example("\n  # Set the user field on the gce context entry without touching other values\n  kubectl config set-context gce --user=cluster-admin"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					String("cluster", "cluster for the context entry in kubeconfig").
					Bool("current", "Modify the current context", ox.Spec("false")).
					String("namespace", "namespace for the context entry in kubeconfig").
					String("user", "user for the context entry in kubeconfig"),
			), ox.Sub(
				ox.Banner("Set a user entry in kubeconfig.\n\n Specifying a name that already exists will merge new fields on top of existing values.\n\n        Client-certificate flags:\n        --client-certificate=certfile --client-key=keyfile\n        \n        Bearer token flags:\n        --token=bearer_token\n        \n        Basic auth flags:\n        --username=basic_user --password=basic_password\n        \n Bearer token and basic auth are mutually exclusive."),
				ox.Usage("set-credentials", "Set a user entry in kubeconfig"),
				ox.Spec("NAME [--client-certificate=path/to/certfile] [--client-key=path/to/keyfile] [--token=bearer_token] [--username=basic_user] [--password=basic_password] [--auth-provider=provider_name] [--auth-provider-arg=key=value] [--exec-command=exec_command] [--exec-api-version=exec_api_version] [--exec-arg=arg] [--exec-env=key=value] [options]"),
				ox.Example("\n  # Set only the \"client-key\" field on the \"cluster-admin\"\n  # entry, without touching other values\n  kubectl config set-credentials cluster-admin --client-key=~/.kube/admin.key\n  \n  # Set basic auth for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --username=admin --password=uXFGweU9l35qcif\n  \n  # Embed client certificate data in the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --client-certificate=~/.kube/admin.crt --embed-certs=true\n  \n  # Enable the Google Compute Platform auth provider for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --auth-provider=gcp\n  \n  # Enable the OpenID Connect auth provider for the \"cluster-admin\" entry with additional arguments\n  kubectl config set-credentials cluster-admin --auth-provider=oidc --auth-provider-arg=client-id=foo --auth-provider-arg=client-secret=bar\n  \n  # Remove the \"client-secret\" config value for the OpenID Connect auth provider for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --auth-provider=oidc --auth-provider-arg=client-secret-\n  \n  # Enable new exec auth plugin for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --exec-command=/path/to/the/executable --exec-api-version=client.authentication.k8s.io/v1beta1\n  \n  # Enable new exec auth plugin for the \"cluster-admin\" entry with interactive mode\n  kubectl config set-credentials cluster-admin --exec-command=/path/to/the/executable --exec-api-version=client.authentication.k8s.io/v1beta1 --exec-interactive-mode=Never\n  \n  # Define new exec auth plugin arguments for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --exec-arg=arg1 --exec-arg=arg2\n  \n  # Create or update exec auth plugin environment variables for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --exec-env=key1=val1 --exec-env=key2=val2\n  \n  # Remove exec auth plugin environment variables for the \"cluster-admin\" entry\n  kubectl config set-credentials cluster-admin --exec-env=var-to-remove-"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					String("auth-provider", "Auth provider for the user entry in kubeconfig").
					Slice("auth-provider-arg", "'key=value' arguments for the auth provider", ox.Elem(ox.StringT)).
					String("client-certificate", "Path to client-certificate file for the user entry in kubeconfig").
					String("client-key", "Path to client-key file for the user entry in kubeconfig").
					Bool("embed-certs", "Embed client cert/key for the user entry in kubeconfig", ox.Spec("false")).
					String("exec-api-version", "API version of the exec credential plugin for the user entry in kubeconfig").
					Slice("exec-arg", "New arguments for the exec credential plugin command for the user entry in kubeconfig", ox.Elem(ox.StringT)).
					String("exec-command", "Command for the exec credential plugin for the user entry in kubeconfig").
					Slice("exec-env", "'key=value' environment values for the exec credential plugin", ox.Elem(ox.StringT)).
					String("exec-interactive-mode", "InteractiveMode of the exec credentials plugin for the user entry in kubeconfig").
					Bool("exec-provide-cluster-info", "ProvideClusterInfo of the exec credentials plugin for the user entry in kubeconfig", ox.Spec("false")).
					String("password", "password for the user entry in kubeconfig").
					String("token", "token for the user entry in kubeconfig").
					String("username", "username for the user entry in kubeconfig"),
			), ox.Sub(
				ox.Banner("Unset an individual value in a kubeconfig file.\n\n PROPERTY_NAME is a dot delimited name where each token represents either an attribute name or a map key.  Map keys may not contain dots."),
				ox.Usage("unset", "Unset an individual value in a kubeconfig file"),
				ox.Spec("PROPERTY_NAME [options]"),
				ox.Example("\n  # Unset the current-context\n  kubectl config unset current-context\n  \n  # Unset namespace in foo context\n  kubectl config unset contexts.foo.namespace"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Set the current-context in a kubeconfig file."),
				ox.Usage("use-context", "Set the current-context in a kubeconfig file"),
				ox.Spec("CONTEXT_NAME [options]"),
				ox.Aliases("use-context", "use"),
				ox.Example("\n  # Use the context for the minikube cluster\n  kubectl config use-context minikube"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			), ox.Sub(
				ox.Banner("Display merged kubeconfig settings or a specified kubeconfig file.\n\n You can use --output jsonpath={...} to extract specific values using a jsonpath expression."),
				ox.Usage("view", "Display merged kubeconfig settings or a specified kubeconfig file"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # Show merged kubeconfig settings\n  kubectl config view\n  \n  # Show merged kubeconfig settings, raw certificate data, and exposed secrets\n  kubectl config view --raw\n  \n  # Get the password for the e2e user\n  kubectl config view -o jsonpath='{.users[?(@.name == \"e2e\")].user.password}'"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("allow-missing-template-keys", "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.", ox.Spec("true")).
					Bool("flatten", "Flatten the resulting kubeconfig file into self-contained output (useful for creating portable kubeconfig files)", ox.Spec("false")).
					Bool("merge", "Merge the full hierarchy of kubeconfig files", ox.Spec("true")).
					Bool("minify", "Remove all information not used by current-context from the output", ox.Spec("false")).
					String("output", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).", ox.Default("yaml"), ox.Short("o")).
					Bool("raw", "Display raw byte data and sensitive data", ox.Spec("false")).
					Bool("show-managed-fields", "If true, keep the managedFields when printing objects in JSON or YAML format.", ox.Spec("false")).
					String("template", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."),
			)), ox.Sub(
			ox.Banner("Provides utilities for interacting with plugins.\n\n Plugins provide extended functionality that is not part of the major command-line distribution. Please refer to the documentation and examples for more information about how write your own plugins.\n\n The easiest way to discover and install plugins is via the kubernetes sub-project krew: [krew.sigs.k8s.io]. To install krew, visit https://krew.sigs.k8s.io/docs/user-guide/setup/install"),
			ox.Usage("plugin", "Provides utilities for interacting with plugins"),
			ox.Spec("[flags] [options]"),
			ox.Example("\n  # List all available plugins\n  kubectl plugin list\n  \n  # List only binary names of available plugins without paths\n  kubectl plugin list --name-only"),
			ox.Section(7),
			ox.Footer("Use \"kubectl plugin <command> --help\" for more information about a given command.\nUse \"kubectl options\" for a list of global command-line options (applies to all commands)."),
			ox.Sub(
				ox.Banner("List all available plugin files on a user's PATH. To see plugins binary names without the full path use --name-only flag.\n\n Available plugin files are those that are: - executable - anywhere on the user's PATH - begin with \"kubectl-\""),
				ox.Usage("list", "List all visible plugin executables on a user's PATH"),
				ox.Spec("[flags] [options]"),
				ox.Example("\n  # List all available plugins\n  kubectl plugin list\n  \n  # List only binary names of available plugins without paths\n  kubectl plugin list --name-only"),
				ox.Footer("Use \"kubectl options\" for a list of global command-line options (applies to all commands)."),
				ox.Flags().
					Bool("name-only", "If true, display only the binary name of each plugin, rather than its full path", ox.Spec("false")),
			)),
	)
}
