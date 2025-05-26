// Command kubectl is a xo/ox version of `kubectl`.
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
		ox.Usage(`kubectl`, ``),
		ox.Banner(`kubectl controls the Kubernetes cluster manager.

 Find more information at: https://kubernetes.io/docs/reference/kubectl/`),
		ox.Spec(`[flags] [options]`),
		ox.Sections("Basic Commands (Beginner)", "Basic Commands (Intermediate)", "Deploy Commands", "Cluster Management Commands", "Troubleshooting and Debugging Commands", "Advanced Commands", "Settings Commands", "Other Commands"),
		ox.Footer(`Use "kubectl <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
		ox.Sub( // kubectl create
			ox.Usage(`create`, `Create a resource from a file or from stdin`),
			ox.Banner(`Create a resource from a file or from stdin.

 JSON and YAML formats are accepted.`),
			ox.Spec(`-f FILENAME [options]`),
			ox.Example(`
  # Create a pod using the data in pod.json
  kubectl create -f ./pod.json
  
  # Create a pod based on the JSON passed into stdin
  cat pod.json | kubectl create -f -
  
  # Edit the data in registry.yaml in JSON then create the resource using the edited data
  kubectl create -f registry.yaml --edit -o json`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl create <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				Bool(`edit`, `Edit the API resource before creating`, ox.Spec(`false`), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files to use to create the resource`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				String(`raw`, `Raw URI to POST to the server.  Uses the transport specified by the kubeconfig file.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
				Bool(`windows-line-endings`, `Only relevant if --edit=true. Defaults to the line ending native to your platform.`, ox.Spec(`false`), ox.Section(0)),
			ox.Sub( // kubectl create clusterrole
				ox.Usage(`clusterrole`, `Create a cluster role`),
				ox.Banner(`Create a cluster role.`),
				ox.Spec(`NAME --verb=verb --resource=resource.group [--resource-name=resourcename] [--dry-run=server|client|none] [options]`),
				ox.Example(`
  # Create a cluster role named "pod-reader" that allows user to perform "get", "watch" and "list" on pods
  kubectl create clusterrole pod-reader --verb=get,list,watch --resource=pods
  
  # Create a cluster role named "pod-reader" with ResourceName specified
  kubectl create clusterrole pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod
  
  # Create a cluster role named "foo" with API Group specified
  kubectl create clusterrole foo --verb=get,list,watch --resource=rs.apps
  
  # Create a cluster role named "foo" with SubResource specified
  kubectl create clusterrole foo --verb=get,list,watch --resource=pods,pods/status
  
  # Create a cluster role name "foo" with NonResourceURL specified
  kubectl create clusterrole "foo" --verb=get --non-resource-url=/logs/*
  
  # Create a cluster role name "monitoring" with AggregationRule specified
  kubectl create clusterrole monitoring --aggregation-rule="rbac.example.com/aggregate-to-monitoring=true"`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`aggregation-rule`, `An aggregation label selector for combining ClusterRoles.`, ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Slice(`non-resource-url`, `A partial url that user should have access to.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Slice(`resource`, `Resource that the rule applies to`, ox.Section(0)).
					Slice(`resource-name`, `Resource in the white list that the rule applies to, repeat this flag for multiple items`, ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
					Slice(`verb`, `Verb that applies to the resources contained in the rule`, ox.Section(0)),
			),
			ox.Sub( // kubectl create clusterrolebinding
				ox.Usage(`clusterrolebinding`, `Create a cluster role binding for a particular cluster role`),
				ox.Banner(`Create a cluster role binding for a particular cluster role.`),
				ox.Spec(`NAME --clusterrole=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]`),
				ox.Example(`
  # Create a cluster role binding for user1, user2, and group1 using the cluster-admin cluster role
  kubectl create clusterrolebinding cluster-admin --clusterrole=cluster-admin --user=user1 --user=user2 --group=group1`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`clusterrole`, `ClusterRole this ClusterRoleBinding should reference`, ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Slice(`group`, `Groups to bind to the clusterrole. The flag can be repeated to add multiple groups.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Slice(`serviceaccount`, `Service accounts to bind to the clusterrole, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					Slice(`user`, `Usernames to bind to the clusterrole. The flag can be repeated to add multiple users.`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create configmap
				ox.Usage(`configmap`, `Create a config map from a local file, directory or literal value`),
				ox.Banner(`Create a config map based on a file, directory, or specified literal value.

 A single config map may package one or more key/value pairs.

 When creating a config map based on a file, the key will default to the basename of the file, and the value will default to the file content.  If the basename is an invalid key, you may specify an alternate key.

 When creating a config map based on a directory, each file whose basename is a valid key in the directory will be packaged into the config map.  Any directory entries except regular files are ignored (e.g. subdirectories, symlinks, devices, pipes, etc).`),
				ox.Spec(`NAME [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none] [options]`),
				ox.Aliases("cm"),
				ox.Example(`
  # Create a new config map named my-config based on folder bar
  kubectl create configmap my-config --from-file=path/to/bar
  
  # Create a new config map named my-config with specified keys instead of file basenames on disk
  kubectl create configmap my-config --from-file=key1=/path/to/bar/file1.txt --from-file=key2=/path/to/bar/file2.txt
  
  # Create a new config map named my-config with key1=config1 and key2=config2
  kubectl create configmap my-config --from-literal=key1=config1 --from-literal=key2=config2
  
  # Create a new config map named my-config from the key=value pairs in the file
  kubectl create configmap my-config --from-file=path/to/bar
  
  # Create a new config map named my-config from an env file
  kubectl create configmap my-config --from-env-file=path/to/foo.env --from-env-file=path/to/bar.env`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Bool(`append-hash`, `Append a hash of the configmap to its name.`, ox.Spec(`false`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Slice(`from-env-file`, `Specify the path to a file to read lines of key=val pairs to create a configmap.`, ox.Section(0)).
					Slice(`from-file`, `Key file can be specified using its file path, in which case file basename will be used as configmap key, or optionally with a key and file path, in which case the given key will be used.  Specifying a directory will iterate each named file in the directory whose basename is a valid configmap key.`, ox.Section(0)).
					Slice(`from-literal`, `Specify a key and literal value to insert in configmap (i.e. mykey=somevalue)`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create cronjob
				ox.Usage(`cronjob`, `Create a cron job with the specified name`),
				ox.Banner(`Create a cron job with the specified name.`),
				ox.Spec(`NAME --image=image --schedule='0/5 * * * ?' -- [COMMAND] [args...] [flags] [options]`),
				ox.Aliases("cj"),
				ox.Example(`
  # Create a cron job
  kubectl create cronjob my-job --image=busybox --schedule="*/1 * * * *"
  
  # Create a cron job with a command
  kubectl create cronjob my-job --image=busybox --schedule="*/1 * * * *" -- date`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`image`, `Image name to run.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					String(`restart`, `job's restart policy. supported values: OnFailure, Never`, ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					String(`schedule`, `A schedule in the Cron format the job should be run with.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create deployment
				ox.Usage(`deployment`, `Create a deployment with the specified name`),
				ox.Banner(`Create a deployment with the specified name.`),
				ox.Spec(`NAME --image=image -- [COMMAND] [args...] [options]`),
				ox.Aliases("deploy"),
				ox.Example(`
  # Create a deployment named my-dep that runs the busybox image
  kubectl create deployment my-dep --image=busybox
  
  # Create a deployment with a command
  kubectl create deployment my-dep --image=busybox -- date
  
  # Create a deployment named my-dep that runs the nginx image with 3 replicas
  kubectl create deployment my-dep --image=nginx --replicas=3
  
  # Create a deployment named my-dep that runs the busybox image and expose port 5701
  kubectl create deployment my-dep --image=busybox --port=5701
  
  # Create a deployment named my-dep that runs multiple containers
  kubectl create deployment my-dep --image=busybox:latest --image=ubuntu:latest --image=nginx`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Slice(`image`, `Image names to run. A deployment can have multiple images set for multi-container pod.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Int(`port`, `The containerPort that this deployment exposes.`, ox.Default("-1"), ox.Section(0)).
					Int(`replicas`, `Number of replicas to create. Default is 1.`, ox.Default("1"), ox.Short("r"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create ingress
				ox.Usage(`ingress`, `Create an ingress with the specified name`),
				ox.Banner(`Create an ingress with the specified name.`),
				ox.Spec(`NAME --rule=host/path=service:port[,tls[=secret]]  [options]`),
				ox.Aliases("ing"),
				ox.Example(`
  # Create a single ingress called 'simple' that directs requests to foo.com/bar to svc
  # svc1:8080 with a TLS secret "my-cert"
  kubectl create ingress simple --rule="foo.com/bar=svc1:8080,tls=my-cert"
  
  # Create a catch all ingress of "/path" pointing to service svc:port and Ingress Class as "otheringress"
  kubectl create ingress catch-all --class=otheringress --rule="/path=svc:port"
  
  # Create an ingress with two annotations: ingress.annotation1 and ingress.annotations2
  kubectl create ingress annotated --class=default --rule="foo.com/bar=svc:port" \
  --annotation ingress.annotation1=foo \
  --annotation ingress.annotation2=bla
  
  # Create an ingress with the same host and multiple paths
  kubectl create ingress multipath --class=default \
  --rule="foo.com/=svc:port" \
  --rule="foo.com/admin/=svcadmin:portadmin"
  
  # Create an ingress with multiple hosts and the pathType as Prefix
  kubectl create ingress ingress1 --class=default \
  --rule="foo.com/path*=svc:8080" \
  --rule="bar.com/admin*=svc2:http"
  
  # Create an ingress with TLS enabled using the default ingress certificate and different path types
  kubectl create ingress ingtls --class=default \
  --rule="foo.com/=svc:https,tls" \
  --rule="foo.com/path/subpath*=othersvc:8080"
  
  # Create an ingress with TLS enabled using a specific secret and pathType as Prefix
  kubectl create ingress ingsecret --class=default \
  --rule="foo.com/*=svc:8080,tls=secret1"
  
  # Create an ingress with a default backend
  kubectl create ingress ingdefault --class=default \
  --default-backend=defaultsvc:http \
  --rule="foo.com/*=svc:8080,tls=secret1"`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`annotation`, `Annotation to insert in the ingress object, in the format annotation=value`, ox.Section(0)).
					String(`class`, `Ingress Class to be used`, ox.Section(0)).
					String(`default-backend`, `Default service for backend, in format of svcname:port`, ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Slice(`rule`, `Rule in format host/path=service:port[,tls=secretname]. Paths containing the leading character '*' are considered pathType=Prefix. tls argument is optional.`, ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create job
				ox.Usage(`job`, `Create a job with the specified name`),
				ox.Banner(`Create a job with the specified name.`),
				ox.Spec(`NAME --image=image [--from=cronjob/name] -- [COMMAND] [args...] [options]`),
				ox.Example(`
  # Create a job
  kubectl create job my-job --image=busybox
  
  # Create a job with a command
  kubectl create job my-job --image=busybox -- date
  
  # Create a job from a cron job named "a-cronjob"
  kubectl create job test-job --from=cronjob/a-cronjob`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`from`, `The name of the resource to create a Job from (only cronjob is supported).`, ox.Section(0)).
					String(`image`, `Image name to run.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create namespace
				ox.Usage(`namespace`, `Create a namespace with the specified name`),
				ox.Banner(`Create a namespace with the specified name.`),
				ox.Spec(`NAME [--dry-run=server|client|none] [options]`),
				ox.Aliases("ns"),
				ox.Example(`
  # Create a new namespace named my-namespace
  kubectl create namespace my-namespace`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create poddisruptionbudget
				ox.Usage(`poddisruptionbudget`, `Create a pod disruption budget with the specified name`),
				ox.Banner(`Create a pod disruption budget with the specified name, selector, and desired minimum available pods.`),
				ox.Spec(`NAME --selector=SELECTOR --min-available=N [--dry-run=server|client|none] [options]`),
				ox.Aliases("pdb"),
				ox.Example(`
  # Create a pod disruption budget named my-pdb that will select all pods with the app=rails label
  # and require at least one of them being available at any point in time
  kubectl create poddisruptionbudget my-pdb --selector=app=rails --min-available=1
  
  # Create a pod disruption budget named my-pdb that will select all pods with the app=nginx label
  # and require at least half of the pods selected to be available at any point in time
  kubectl create pdb my-pdb --selector=app=nginx --min-available=50%`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`max-unavailable`, `The maximum number or percentage of unavailable pods this budget requires.`, ox.Section(0)).
					String(`min-available`, `The minimum number or percentage of available pods this budget requires.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					String(`selector`, `A label selector to use for this budget. Only equality-based selector requirements are supported.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create priorityclass
				ox.Usage(`priorityclass`, `Create a priority class with the specified name`),
				ox.Banner(`Create a priority class with the specified name, value, globalDefault and description.`),
				ox.Spec(`NAME --value=VALUE --global-default=BOOL [--dry-run=server|client|none] [options]`),
				ox.Aliases("pc"),
				ox.Example(`
  # Create a priority class named high-priority
  kubectl create priorityclass high-priority --value=1000 --description="high priority"
  
  # Create a priority class named default-priority that is considered as the global default priority
  kubectl create priorityclass default-priority --value=1000 --global-default=true --description="default priority"
  
  # Create a priority class named high-priority that cannot preempt pods with lower priority
  kubectl create priorityclass high-priority --value=1000 --description="high priority" --preemption-policy="Never"`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`description`, `description is an arbitrary string that usually provides guidelines on when this priority class should be used.`, ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Bool(`global-default`, `global-default specifies whether this PriorityClass should be considered as the default priority.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					String(`preemption-policy`, `preemption-policy is the policy for preempting pods with lower priority.`, ox.Default("PreemptLowerPriority"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
					Int(`value`, `the value of this priority class.`, ox.Default("0"), ox.Section(0)),
			),
			ox.Sub( // kubectl create quota
				ox.Usage(`quota`, `Create a quota with the specified name`),
				ox.Banner(`Create a resource quota with the specified name, hard limits, and optional scopes.`),
				ox.Spec(`NAME [--hard=key1=value1,key2=value2] [--scopes=Scope1,Scope2] [--dry-run=server|client|none] [options]`),
				ox.Aliases("resourcequota"),
				ox.Example(`
  # Create a new resource quota named my-quota
  kubectl create quota my-quota --hard=cpu=1,memory=1G,pods=2,services=3,replicationcontrollers=2,resourcequotas=1,secrets=5,persistentvolumeclaims=10
  
  # Create a new resource quota named best-effort
  kubectl create quota best-effort --hard=pods=100 --scopes=BestEffort`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`hard`, `A comma-delimited set of resource=quantity pairs that define a hard limit.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					String(`scopes`, `A comma-delimited set of quota scopes that must all match each object tracked by the quota.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create role
				ox.Usage(`role`, `Create a role with single rule`),
				ox.Banner(`Create a role with single rule.`),
				ox.Spec(`NAME --verb=verb --resource=resource.group/subresource [--resource-name=resourcename] [--dry-run=server|client|none] [options]`),
				ox.Example(`
  # Create a role named "pod-reader" that allows user to perform "get", "watch" and "list" on pods
  kubectl create role pod-reader --verb=get --verb=list --verb=watch --resource=pods
  
  # Create a role named "pod-reader" with ResourceName specified
  kubectl create role pod-reader --verb=get --resource=pods --resource-name=readablepod --resource-name=anotherpod
  
  # Create a role named "foo" with API Group specified
  kubectl create role foo --verb=get,list,watch --resource=rs.apps
  
  # Create a role named "foo" with SubResource specified
  kubectl create role foo --verb=get,list,watch --resource=pods,pods/status`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Slice(`resource`, `Resource that the rule applies to`, ox.Section(0)).
					Slice(`resource-name`, `Resource in the white list that the rule applies to, repeat this flag for multiple items`, ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
					Slice(`verb`, `Verb that applies to the resources contained in the rule`, ox.Section(0)),
			),
			ox.Sub( // kubectl create rolebinding
				ox.Usage(`rolebinding`, `Create a role binding for a particular role or cluster role`),
				ox.Banner(`Create a role binding for a particular role or cluster role.`),
				ox.Spec(`NAME --clusterrole=NAME|--role=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]`),
				ox.Example(`
  # Create a role binding for user1, user2, and group1 using the admin cluster role
  kubectl create rolebinding admin --clusterrole=admin --user=user1 --user=user2 --group=group1
  
  # Create a role binding for service account monitoring:sa-dev using the admin role
  kubectl create rolebinding admin-binding --role=admin --serviceaccount=monitoring:sa-dev`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`clusterrole`, `ClusterRole this RoleBinding should reference`, ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					Slice(`group`, `Groups to bind to the role. The flag can be repeated to add multiple groups.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					String(`role`, `Role this RoleBinding should reference`, ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Slice(`serviceaccount`, `Service accounts to bind to the role, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					Slice(`user`, `Usernames to bind to the role. The flag can be repeated to add multiple users.`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create secret
				ox.Usage(`secret`, `Create a secret using a specified subcommand`),
				ox.Banner(`Create a secret with specified type.

 A docker-registry type secret is for accessing a container registry.

 A generic type secret indicate an Opaque secret type.

 A tls type secret holds TLS certificate and its associated key.`),
				ox.Spec(`(docker-registry | generic | tls) [options]`),
				ox.Footer(`Use "kubectl create secret <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Sub( // kubectl create secret docker-registry
					ox.Usage(`docker-registry`, `Create a secret for use with a Docker registry`),
					ox.Banner(`Create a new secret for use with Docker registries.
        
        Dockercfg secrets are used to authenticate against Docker registries.
        
        When using the Docker command line to push images, you can authenticate to a given registry by running:
        '$ docker login DOCKER_REGISTRY_SERVER --username=DOCKER_USER --password=DOCKER_PASSWORD --email=DOCKER_EMAIL'.
        
 That produces a ~/.dockercfg file that is used by subsequent 'docker push' and 'docker pull' commands to authenticate to the registry. The email address is optional.

        When creating applications, you may have a Docker registry that requires authentication.  In order for the
        nodes to pull images on your behalf, they must have the credentials.  You can provide this information
        by creating a dockercfg secret and attaching it to your service account.`),
					ox.Spec(`NAME --docker-username=user --docker-password=password --docker-email=email [--docker-server=string] [--from-file=[key=]source] [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # If you do not already have a .dockercfg file, create a dockercfg secret directly
  kubectl create secret docker-registry my-secret --docker-server=DOCKER_REGISTRY_SERVER --docker-username=DOCKER_USER --docker-password=DOCKER_PASSWORD --docker-email=DOCKER_EMAIL
  
  # Create a new secret named my-secret from ~/.docker/config.json
  kubectl create secret docker-registry my-secret --from-file=path/to/.docker/config.json`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						Bool(`append-hash`, `Append a hash of the secret to its name.`, ox.Spec(`false`), ox.Section(0)).
						String(`docker-email`, `Email for Docker registry`, ox.Section(0)).
						String(`docker-password`, `Password for Docker registry authentication`, ox.Section(0)).
						String(`docker-server`, `Server location for Docker registry`, ox.Default("https://index.docker.io/v1/"), ox.Section(0)).
						String(`docker-username`, `Username for Docker registry authentication`, ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						Slice(`from-file`, `Key files can be specified using their file path, in which case a default name of .dockerconfigjson will be given to them, or optionally with a name and file path, in which case the given name will be used. Specifying a directory will iterate each named file in the directory that is a valid secret key. For this command, the key should always be .dockerconfigjson.`, ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
				ox.Sub( // kubectl create secret generic
					ox.Usage(`generic`, `Create a secret from a local file, directory, or literal value`),
					ox.Banner(`Create a secret based on a file, directory, or specified literal value.

 A single secret may package one or more key/value pairs.

 When creating a secret based on a file, the key will default to the basename of the file, and the value will default to the file content. If the basename is an invalid key or you wish to chose your own, you may specify an alternate key.

 When creating a secret based on a directory, each file whose basename is a valid key in the directory will be packaged into the secret. Any directory entries except regular files are ignored (e.g. subdirectories, symlinks, devices, pipes, etc).`),
					ox.Spec(`NAME [--type=string] [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new secret named my-secret with keys for each file in folder bar
  kubectl create secret generic my-secret --from-file=path/to/bar
  
  # Create a new secret named my-secret with specified keys instead of names on disk
  kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa --from-file=ssh-publickey=path/to/id_rsa.pub
  
  # Create a new secret named my-secret with key1=supersecret and key2=topsecret
  kubectl create secret generic my-secret --from-literal=key1=supersecret --from-literal=key2=topsecret
  
  # Create a new secret named my-secret using a combination of a file and a literal
  kubectl create secret generic my-secret --from-file=ssh-privatekey=path/to/id_rsa --from-literal=passphrase=topsecret
  
  # Create a new secret named my-secret from env files
  kubectl create secret generic my-secret --from-env-file=path/to/foo.env --from-env-file=path/to/bar.env`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						Bool(`append-hash`, `Append a hash of the secret to its name.`, ox.Spec(`false`), ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						Slice(`from-env-file`, `Specify the path to a file to read lines of key=val pairs to create a secret.`, ox.Section(0)).
						Slice(`from-file`, `Key files can be specified using their file path, in which case a default name will be given to them, or optionally with a name and file path, in which case the given name will be used.  Specifying a directory will iterate each named file in the directory that is a valid secret key.`, ox.Section(0)).
						Slice(`from-literal`, `Specify a key and literal value to insert in secret (i.e. mykey=somevalue)`, ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`type`, `The type of secret to create`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
				ox.Sub( // kubectl create secret tls
					ox.Usage(`tls`, `Create a TLS secret`),
					ox.Banner(`Create a TLS secret from the given public/private key pair.

 The public/private key pair must exist beforehand. The public key certificate must be .PEM encoded and match the given private key.`),
					ox.Spec(`NAME --cert=path/to/cert/file --key=path/to/key/file [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new TLS secret named tls-secret with the given key pair
  kubectl create secret tls tls-secret --cert=path/to/tls.crt --key=path/to/tls.key`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						Bool(`append-hash`, `Append a hash of the secret to its name.`, ox.Spec(`false`), ox.Section(0)).
						String(`cert`, `Path to PEM encoded public key certificate.`, ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						String(`key`, `Path to private key associated with given certificate.`, ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
			),
			ox.Sub( // kubectl create service
				ox.Usage(`service`, `Create a service using a specified subcommand`),
				ox.Banner(`Create a service using a specified subcommand.`),
				ox.Spec(`[flags] [options]`),
				ox.Aliases("svc"),
				ox.Footer(`Use "kubectl create service <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Sub( // kubectl create service clusterip
					ox.Usage(`clusterip`, `Create a ClusterIP service`),
					ox.Banner(`Create a ClusterIP service with the specified name.`),
					ox.Spec(`NAME [--tcp=<port>:<targetPort>] [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new ClusterIP service named my-cs
  kubectl create service clusterip my-cs --tcp=5678:8080
  
  # Create a new ClusterIP service named my-cs (in headless mode)
  kubectl create service clusterip my-cs --clusterip="None"`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						String(`clusterip`, `Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing).`, ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						Slice(`tcp`, `Port pairs can be specified as '<port>:<targetPort>'.`, ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
				ox.Sub( // kubectl create service externalname
					ox.Usage(`externalname`, `Create an ExternalName service`),
					ox.Banner(`Create an ExternalName service with the specified name.

 ExternalName service references to an external DNS address instead of only pods, which will allow application authors to reference services that exist off platform, on other clusters, or locally.`),
					ox.Spec(`NAME --external-name external.name [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new ExternalName service named my-ns
  kubectl create service externalname my-ns --external-name bar.com`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`external-name`, `External name of service`, ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						Slice(`tcp`, `Port pairs can be specified as '<port>:<targetPort>'.`, ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
				ox.Sub( // kubectl create service loadbalancer
					ox.Usage(`loadbalancer`, `Create a LoadBalancer service`),
					ox.Banner(`Create a LoadBalancer service with the specified name.`),
					ox.Spec(`NAME [--tcp=port:targetPort] [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new LoadBalancer service named my-lbs
  kubectl create service loadbalancer my-lbs --tcp=5678:8080`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						Slice(`tcp`, `Port pairs can be specified as '<port>:<targetPort>'.`, ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
				ox.Sub( // kubectl create service nodeport
					ox.Usage(`nodeport`, `Create a NodePort service`),
					ox.Banner(`Create a NodePort service with the specified name.`),
					ox.Spec(`NAME [--tcp=port:targetPort] [--dry-run=server|client|none] [options]`),
					ox.Example(`
  # Create a new NodePort service named my-ns
  kubectl create service nodeport my-ns --tcp=5678:8080`),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
					ox.Flags().
						Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
						String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
						String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
						Int(`node-port`, `Port used to expose the service on each node in a cluster.`, ox.Default("0"), ox.Section(0)).
						String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
						Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
						Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
						Slice(`tcp`, `Port pairs can be specified as '<port>:<targetPort>'.`, ox.Section(0)).
						String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
						String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
				),
			),
			ox.Sub( // kubectl create serviceaccount
				ox.Usage(`serviceaccount`, `Create a service account with the specified name`),
				ox.Banner(`Create a service account with the specified name.`),
				ox.Spec(`NAME [--dry-run=server|client|none] [options]`),
				ox.Aliases("sa"),
				ox.Example(`
  # Create a new service account named my-service-account
  kubectl create serviceaccount my-service-account`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-create"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
			),
			ox.Sub( // kubectl create token
				ox.Usage(`token`, `Request a service account token`),
				ox.Banner(`Request a service account token.`),
				ox.Spec(`SERVICE_ACCOUNT_NAME [options]`),
				ox.Example(`
  # Request a token to authenticate to the kube-apiserver as the service account "myapp" in the current namespace
  kubectl create token myapp
  
  # Request a token for a service account in a custom namespace
  kubectl create token myapp --namespace myns
  
  # Request a token with a custom expiration
  kubectl create token myapp --duration 10m
  
  # Request a token with a custom audience
  kubectl create token myapp --audience https://example.com
  
  # Request a token bound to an instance of a Secret object
  kubectl create token myapp --bound-object-kind Secret --bound-object-name mysecret
  
  # Request a token bound to an instance of a Secret object with a specific UID
  kubectl create token myapp --bound-object-kind Secret --bound-object-name mysecret --bound-object-uid 0d4691ed-659b-4935-a832-355f77ee47cc`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`audience`, `Audience of the requested token. If unset, defaults to requesting a token for use with the Kubernetes API server. May be repeated to request a token valid for multiple audiences.`, ox.Section(0)).
					String(`bound-object-kind`, `Kind of an object to bind the token to. Supported kinds are Node, Pod, Secret. If set, --bound-object-name must be provided.`, ox.Section(0)).
					String(`bound-object-name`, `Name of an object to bind the token to. The token will expire when the object is deleted. Requires --bound-object-kind.`, ox.Section(0)).
					String(`bound-object-uid`, `UID of an object to bind the token to. Requires --bound-object-kind and --bound-object-name. If unset, the UID of the existing object is used.`, ox.Section(0)).
					Duration(`duration`, `Requested lifetime of the issued token. If not set or if set to 0, the lifetime will be determined by the server automatically. The server may return a token with a longer or shorter lifetime.`, ox.Default("0s"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl expose
			ox.Usage(`expose`, `Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service`),
			ox.Banner(`Expose a resource as a new Kubernetes service.

 Looks up a deployment, service, replica set, replication controller or pod by name and uses the selector for that resource as the selector for a new service on the specified port. A deployment or replica set will be exposed as a service only if its selector is convertible to a selector that service supports, i.e. when the selector contains only the matchLabels component. Note that if no port is specified via --port and the exposed resource has multiple ports, all will be re-used by the new service. Also if no labels are specified, the new service will re-use the labels from the resource it exposes.

 Possible resources include (case insensitive):

 pod (po), service (svc), replicationcontroller (rc), deployment (deploy), replicaset (rs)`),
			ox.Spec(`(-f FILENAME | TYPE NAME) [--port=port] [--protocol=TCP|UDP|SCTP] [--target-port=number-or-name] [--name=name] [--external-ip=external-ip-of-service] [--type=type] [options]`),
			ox.Example(`
  # Create a service for a replicated nginx, which serves on port 80 and connects to the containers on port 8000
  kubectl expose rc nginx --port=80 --target-port=8000
  
  # Create a service for a replication controller identified by type and name specified in "nginx-controller.yaml", which serves on port 80 and connects to the containers on port 8000
  kubectl expose -f nginx-controller.yaml --port=80 --target-port=8000
  
  # Create a service for a pod valid-pod, which serves on port 444 with the name "frontend"
  kubectl expose pod valid-pod --port=444 --name=frontend
  
  # Create a second service based on the above service, exposing the container port 8443 as port 443 with the name "nginx-https"
  kubectl expose service nginx --port=443 --target-port=8443 --name=nginx-https
  
  # Create a service for a replicated streaming application on port 4100 balancing UDP traffic and named 'video-stream'.
  kubectl expose rc streamer --port=4100 --protocol=UDP --name=video-stream
  
  # Create a service for a replicated nginx using replica set, which serves on port 80 and connects to the containers on port 8000
  kubectl expose rs nginx --port=80 --target-port=8000
  
  # Create a service for an nginx deployment, which serves on port 80 and connects to the containers on port 8000
  kubectl expose deployment nginx --port=80 --target-port=8000`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`cluster-ip`, `ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service.`, ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`external-ip`, `Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP.`, ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-expose"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to expose a service`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`labels`, `Labels to apply to the service created by this call.`, ox.Short("l"), ox.Section(0)).
				String(`load-balancer-ip`, `IP to assign to the LoadBalancer. If empty, an ephemeral IP will be created and used (cloud-provider specific).`, ox.Section(0)).
				String(`name`, `The name for the newly created object.`, ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				String(`override-type`, `The method used to override the generated object: json, merge, or strategic.`, ox.Default("merge"), ox.Section(0)).
				String(`overrides`, `An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.`, ox.Section(0)).
				String(`port`, `The port that the service should serve on. Copied from the resource being exposed, if unspecified`, ox.Section(0)).
				String(`protocol`, `The network protocol for the service to be created. Default is 'TCP'.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				String(`selector`, `A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.)`, ox.Section(0)).
				String(`session-affinity`, `If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP'`, ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`target-port`, `Name or number for the port on the container that the service should direct traffic to. Optional.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				String(`type`, `Type for this service: ClusterIP, NodePort, LoadBalancer, or ExternalName. Default is 'ClusterIP'.`, ox.Section(0)),
		),
		ox.Sub( // kubectl run
			ox.Usage(`run`, `Run a particular image on the cluster`),
			ox.Banner(`Create and run a particular image in a pod.`),
			ox.Spec(`NAME --image=image [--env="key=value"] [--port=port] [--dry-run=server|client] [--overrides=inline-json] [--command] -- [COMMAND] [args...] [options]`),
			ox.Example(`
  # Start a nginx pod
  kubectl run nginx --image=nginx
  
  # Start a hazelcast pod and let the container expose port 5701
  kubectl run hazelcast --image=hazelcast/hazelcast --port=5701
  
  # Start a hazelcast pod and set environment variables "DNS_DOMAIN=cluster" and "POD_NAMESPACE=default" in the container
  kubectl run hazelcast --image=hazelcast/hazelcast --env="DNS_DOMAIN=cluster" --env="POD_NAMESPACE=default"
  
  # Start a hazelcast pod and set labels "app=hazelcast" and "env=prod" in the container
  kubectl run hazelcast --image=hazelcast/hazelcast --labels="app=hazelcast,env=prod"
  
  # Dry run; print the corresponding API objects without creating them
  kubectl run nginx --image=nginx --dry-run=client
  
  # Start a nginx pod, but overload the spec with a partial set of values parsed from JSON
  kubectl run nginx --image=nginx --overrides='{ "apiVersion": "v1", "spec": { ... } }'
  
  # Start a busybox pod and keep it in the foreground, don't restart it if it exits
  kubectl run -i -t busybox --image=busybox --restart=Never
  
  # Start the nginx pod using the default command, but use custom arguments (arg1 .. argN) for that command
  kubectl run nginx --image=nginx -- <arg1> <arg2> ... <argN>
  
  # Start the nginx pod using a different command and custom arguments
  kubectl run nginx --image=nginx --command -- <cmd> <arg1> ... <argN>`),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				Slice(`annotations`, `Annotations to apply to the pod.`, ox.Section(0)).
				Bool(`attach`, `If true, wait for the Pod to start running, and then attach to the Pod as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true. With '--restart=Never' the exit code of the container process is returned.`, ox.Spec(`false`), ox.Section(0)).
				String(`cascade`, `Must be "background", "orphan", or "foreground". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.`, ox.Default("background"), ox.Section(0)).
				Bool(`command`, `If true and extra arguments are present, use them as the 'command' field in the container, rather than the 'args' field which is the default.`, ox.Spec(`false`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				Slice(`env`, `Environment variables to set in the container.`, ox.Section(0)).
				Bool(`expose`, `If true, create a ClusterIP service associated with the pod.  Requires `+"`"+`--port`+"`"+`.`, ox.Spec(`false`), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-run"), ox.Section(0)).
				Slice(`filename`, `to use to replace the resource.`, ox.Short("f"), ox.Section(0)).
				Bool(`force`, `If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.`, ox.Spec(`false`), ox.Section(0)).
				Int(`grace-period`, `Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).`, ox.Default("-1"), ox.Section(0)).
				String(`image`, `The image for the container to run.`, ox.Section(0)).
				String(`image-pull-policy`, `The image pull policy for the container.  If left empty, this value will not be specified by the client and defaulted by the server.`, ox.Section(0)).
				String(`kustomize`, `Process a kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`labels`, `Comma separated labels to apply to the pod. Will override previous values.`, ox.Short("l"), ox.Section(0)).
				Bool(`leave-stdin-open`, `If the pod is started in interactive mode or with stdin, leave stdin open after the first attach completes. By default, stdin will be closed after the first attach completes.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				String(`override-type`, `The method used to override the generated object: json, merge, or strategic.`, ox.Default("merge"), ox.Section(0)).
				String(`overrides`, `An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field.`, ox.Section(0)).
				Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("1m0s"), ox.Section(0)).
				String(`port`, `The port that this container exposes.`, ox.Section(0)).
				Bool(`privileged`, `If true, run the container in privileged mode.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`quiet`, `If true, suppress prompt messages.`, ox.Spec(`false`), ox.Short("q"), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`restart`, `The restart policy for this Pod.  Legal values [Always, OnFailure, Never].`, ox.Default("Always"), ox.Section(0)).
				Bool(`rm`, `If true, delete the pod after it exits.  Only valid when attaching to the container, e.g. with '--attach' or with '-i/--stdin'.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`stdin`, `Keep stdin open on the container in the pod, even if nothing is attached.`, ox.Spec(`false`), ox.Short("i"), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object`, ox.Default("0s"), ox.Section(0)).
				Bool(`tty`, `Allocate a TTY for the container in the pod.`, ox.Spec(`false`), ox.Short("t"), ox.Section(0)).
				Bool(`wait`, `If true, wait for resources to be gone before returning. This waits for finalizers.`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl set
			ox.Usage(`set`, `Set specific features on objects`),
			ox.Banner(`Configure application resources.

 These commands help you make changes to existing application resources.`),
			ox.Spec(`SUBCOMMAND [options]`),
			ox.Section(0),
			ox.Footer(`Use "kubectl set <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl set env
				ox.Usage(`env`, `Update environment variables on a pod template`),
				ox.Banner(`Update environment variables on a pod template.

 List environment variable definitions in one or more pods, pod templates. Add, update, or remove container environment variable definitions in one or more pod templates (within replication controllers or deployment configurations). View or modify the environment variable definitions on all containers in the specified pods or pod templates, or just those that match a wildcard.

 If "--env -" is passed, environment variables can be read from STDIN using the standard env syntax.

 Possible resources include (case insensitive):

        pod (po), replicationcontroller (rc), deployment (deploy), daemonset (ds), statefulset (sts), cronjob (cj), replicaset (rs)`),
				ox.Spec(`RESOURCE/NAME KEY_1=VAL_1 ... KEY_N=VAL_N [options]`),
				ox.Example(`
  # Update deployment 'registry' with a new environment variable
  kubectl set env deployment/registry STORAGE_DIR=/local
  
  # List the environment variables defined on a deployments 'sample-build'
  kubectl set env deployment/sample-build --list
  
  # List the environment variables defined on all pods
  kubectl set env pods --all --list
  
  # Output modified deployment in YAML, and does not alter the object on the server
  kubectl set env deployment/sample-build STORAGE_DIR=/data -o yaml
  
  # Update all containers in all replication controllers in the project to have ENV=prod
  kubectl set env rc --all ENV=prod
  
  # Import environment from a secret
  kubectl set env --from=secret/mysecret deployment/myapp
  
  # Import environment from a config map with a prefix
  kubectl set env --from=configmap/myconfigmap --prefix=MYSQL_ deployment/myapp
  
  # Import specific keys from a config map
  kubectl set env --keys=my-example-key --from=configmap/myconfigmap deployment/myapp
  
  # Remove the environment variable ENV from container 'c1' in all deployment configs
  kubectl set env deployments --all --containers="c1" ENV-
  
  # Remove the environment variable ENV from a deployment definition on disk and
  # update the deployment config on the server
  kubectl set env -f deploy.json ENV-
  
  # Set some of the local shell environment into a deployment config on the server
  env | grep RAILS_ | kubectl set env -e - deployment/registry`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `If true, select all resources in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`containers`, `The names of containers in the selected pod templates to change - may use wildcards`, ox.Default("*"), ox.Short("c"), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					Slice(`env`, `Specify a key-value pair for an environment variable to set into each container.`, ox.Short("e"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files the resource to update the env`, ox.Short("f"), ox.Section(0)).
					String(`from`, `The name of a resource from which to inject environment variables`, ox.Section(0)).
					Slice(`keys`, `Comma-separated list of keys to import from specified resource`, ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					Bool(`list`, `If true, display the environment and any changes in the standard format. this flag will removed when we have kubectl view env.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`local`, `If true, set env will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`overwrite`, `If true, allow environment to be overwritten, otherwise reject updates that overwrite existing environment.`, ox.Spec(`true`), ox.Section(0)).
					String(`prefix`, `Prefix to append to variable names`, ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`resolve`, `If true, show secret or configmap references when listing variables`, ox.Spec(`false`), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl set image
				ox.Usage(`image`, `Update the image of a pod template`),
				ox.Banner(`Update existing container image(s) of resources.

 Possible resources include (case insensitive):

        pod (po), replicationcontroller (rc), deployment (deploy), daemonset (ds), statefulset (sts), cronjob (cj), replicaset (rs)`),
				ox.Spec(`(-f FILENAME | TYPE NAME) CONTAINER_NAME_1=CONTAINER_IMAGE_1 ... CONTAINER_NAME_N=CONTAINER_IMAGE_N [options]`),
				ox.Example(`
  # Set a deployment's nginx container image to 'nginx:1.9.1', and its busybox container image to 'busybox'
  kubectl set image deployment/nginx busybox=busybox nginx=nginx:1.9.1
  
  # Update all deployments' and rc's nginx container's image to 'nginx:1.9.1'
  kubectl set image deployments,rc nginx=nginx:1.9.1 --all
  
  # Update image of all containers of daemonset abc to 'nginx:1.9.1'
  kubectl set image daemonset abc *=nginx:1.9.1
  
  # Print result (in yaml format) of updating nginx container image from local file, without hitting the server
  kubectl set image -f path/to/file.yaml nginx=nginx:1.9.1 --local -o yaml`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources, in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					Bool(`local`, `If true, set image will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl set resources
				ox.Usage(`resources`, `Update resource requests/limits on objects with pod templates`),
				ox.Banner(`Specify compute resource requirements (CPU, memory) for any resource that defines a pod template.  If a pod is successfully scheduled, it is guaranteed the amount of resource requested, but may burst up to its specified limits.

 For each compute resource, if a limit is specified and a request is omitted, the request will default to the limit.

 Possible resources include (case insensitive): Use "kubectl api-resources" for a complete list of supported resources..`),
				ox.Spec(`(-f FILENAME | TYPE NAME)  ([--limits=LIMITS & --requests=REQUESTS] [options]`),
				ox.Example(`
  # Set a deployments nginx container cpu limits to "200m" and memory to "512Mi"
  kubectl set resources deployment nginx -c=nginx --limits=cpu=200m,memory=512Mi
  
  # Set the resource request and limits for all containers in nginx
  kubectl set resources deployment nginx --limits=cpu=200m,memory=512Mi --requests=cpu=100m,memory=256Mi
  
  # Remove the resource requests for resources on containers in nginx
  kubectl set resources deployment nginx --limits=cpu=0,memory=0 --requests=cpu=0,memory=0
  
  # Print the result (in yaml format) of updating nginx container limits from a local, without hitting the server
  kubectl set resources -f path/to/file.yaml --limits=cpu=200m,memory=512Mi --local -o yaml`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources, in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`containers`, `The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards`, ox.Default("*"), ox.Short("c"), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`limits`, `The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.`, ox.Section(0)).
					Bool(`local`, `If true, set resources will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`requests`, `The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.`, ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl set selector
				ox.Usage(`selector`, `Set the selector on a resource`),
				ox.Banner(`Set the selector on a resource. Note that the new selector will overwrite the old selector if the resource had one prior to the invocation of 'set selector'.

 A selector must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters. If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used. Note: currently selectors can only be set on Service objects.`),
				ox.Spec(`(-f FILENAME | TYPE NAME) EXPRESSIONS [--resource-version=version] [options]`),
				ox.Example(`
  # Set the labels and selector before creating a deployment/service pair
  kubectl create service clusterip my-svc --clusterip="None" -o yaml --dry-run=client | kubectl set selector --local -f - 'environment=qa' -o yaml | kubectl create -f -
  kubectl create deployment my-dep -o yaml --dry-run=client | kubectl label --local -f - environment=qa -o yaml | kubectl create -f -`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `identifying the resource.`, ox.Short("f"), ox.Section(0)).
					Bool(`local`, `If true, annotation will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`true`), ox.Short("R"), ox.Section(0)).
					String(`resource-version`, `If non-empty, the selectors update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl set serviceaccount
				ox.Usage(`serviceaccount`, `Update the service account of a resource`),
				ox.Banner(`Update the service account of pod template resources.

 Possible resources (case insensitive) can be:

 replicationcontroller (rc), deployment (deploy), daemonset (ds), job, replicaset (rs), statefulset`),
				ox.Spec(`(-f FILENAME | TYPE NAME) SERVICE_ACCOUNT [options]`),
				ox.Aliases("sa"),
				ox.Example(`
  # Set deployment nginx-deployment's service account to serviceaccount1
  kubectl set serviceaccount deployment nginx-deployment serviceaccount1
  
  # Print the result (in YAML format) of updated nginx deployment with the service account from local file, without hitting the API server
  kubectl set sa -f nginx-deployment.yaml serviceaccount1 --local --dry-run=client -o yaml`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources, in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					Bool(`local`, `If true, set serviceaccount will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl set subject
				ox.Usage(`subject`, `Update the user, group, or service account in a role binding or cluster role binding`),
				ox.Banner(`Update the user, group, or service account in a role binding or cluster role binding.`),
				ox.Spec(`(-f FILENAME | TYPE NAME) [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none] [options]`),
				ox.Example(`
  # Update a cluster role binding for serviceaccount1
  kubectl set subject clusterrolebinding admin --serviceaccount=namespace:serviceaccount1
  
  # Update a role binding for user1, user2, and group1
  kubectl set subject rolebinding admin --user=user1 --user=user2 --group=group1
  
  # Print the result (in YAML format) of updating rolebinding subjects from a local, without hitting the server
  kubectl create rolebinding admin --role=admin --user=admin -o yaml --dry-run=client | kubectl set subject --local -f - --user=foo -o yaml`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources, in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-set"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files the resource to update the subjects`, ox.Short("f"), ox.Section(0)).
					Slice(`group`, `Groups to bind to the role`, ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					Bool(`local`, `If true, set subject will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Slice(`serviceaccount`, `Service accounts to bind to the role`, ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					Slice(`user`, `Usernames to bind to the role`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl explain
			ox.Usage(`explain`, `Get documentation for a resource`),
			ox.Banner(`Describe fields and structure of various resources.

 This command describes the fields associated with each supported API resource. Fields are identified via a simple JSONPath identifier:

        <type>.<fieldName>[.<fieldName>]
        
 Information about each field is retrieved from the server in OpenAPI format.

Use "kubectl api-resources" for a complete list of supported resources.`),
			ox.Spec(`TYPE [--recursive=FALSE|TRUE] [--api-version=api-version-group] [-o|--output=plaintext|plaintext-openapiv2] [options]`),
			ox.Example(`
  # Get the documentation of the resource and its fields
  kubectl explain pods
  
  # Get all the fields in the resource
  kubectl explain pods --recursive
  
  # Get the explanation for deployment in supported api versions
  kubectl explain deployments --api-version=apps/v1
  
  # Get the documentation of a specific field of a resource
  kubectl explain pods.spec.containers
  
  # Get the documentation of resources in different format
  kubectl explain deployment --output=plaintext-openapiv2`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`api-version`, `Get different explanations for particular API version (API group/version)`, ox.Section(0)).
				String(`output`, `Format in which to render the schema (plaintext, plaintext-openapiv2)`, ox.Default("plaintext"), ox.Section(0)).
				Bool(`recursive`, `Print the fields of fields (Currently only 1 level deep)`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl get
			ox.Usage(`get`, `Display one or many resources`),
			ox.Banner(`Display one or many resources.

 Prints a table of the most important information about the specified resources. You can filter the list using a label selector and the --selector flag. If the desired resource type is namespaced you will only see results in the current namespace if you don't specify any namespace.

 By specifying the output as 'template' and providing a Go template as the value of the --template flag, you can filter the attributes of the fetched resources.

Use "kubectl api-resources" for a complete list of supported resources.`),
			ox.Spec(`[(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns|custom-columns-file|wide] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags] [options]`),
			ox.Example(`
  # List all pods in ps output format
  kubectl get pods
  
  # List all pods in ps output format with more information (such as node name)
  kubectl get pods -o wide
  
  # List a single replication controller with specified NAME in ps output format
  kubectl get replicationcontroller web
  
  # List deployments in JSON output format, in the "v1" version of the "apps" API group
  kubectl get deployments.v1.apps -o json
  
  # List a single pod in JSON output format
  kubectl get -o json pod web-pod-13je7
  
  # List a pod identified by type and name specified in "pod.yaml" in JSON output format
  kubectl get -f pod.yaml -o json
  
  # List resources from a directory with kustomization.yaml - e.g. dir/kustomization.yaml
  kubectl get -k dir/
  
  # Return only the phase value of the specified pod
  kubectl get -o template pod/web-pod-13je7 --template={{.status.phase}}
  
  # List resource information in custom columns
  kubectl get pod test-pod -o custom-columns=CONTAINER:.spec.containers[0].name,IMAGE:.spec.containers[0].image
  
  # List all replication controllers and services together in ps output format
  kubectl get rc,services
  
  # List one or more resources by their type and names
  kubectl get rc/web service/frontend pods/web-pod-13je7
  
  # List the 'status' subresource for a single pod
  kubectl get pod web-pod-13je7 --subresource status
  
  # List all deployments in namespace 'backend'
  kubectl get deployments.apps --namespace backend
  
  # List all pods existing in all namespaces
  kubectl get pods --all-namespaces`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				Int(`chunk-size`, `Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.`, ox.Default("500"), ox.Section(0)).
				String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
				Bool(`ignore-not-found`, `If the requested object does not exist the command will return exit code 0.`, ox.Spec(`false`), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Slice(`label-columns`, `Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2...`, ox.Short("L"), ox.Section(0)).
				Bool(`no-headers`, `When using the default or custom-column output format, don't print headers (default print headers).`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file, custom-columns, custom-columns-file, wide). See custom columns [https://kubernetes.io/docs/reference/kubectl/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/].`, ox.Short("o"), ox.Section(0)).
				Bool(`output-watch-events`, `Output watch event objects when --watch or --watch-only is used. Existing objects are output as initial ADDED events.`, ox.Spec(`false`), ox.Section(0)).
				String(`raw`, `Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`server-print`, `If true, have the server return the appropriate table output. Supports extension APIs and CRDs.`, ox.Spec(`true`), ox.Section(0)).
				Bool(`show-kind`, `If present, list the resource type for the requested object(s).`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-labels`, `When printing, show all labels as the last column`, ox.Spec(`false`), ox.Default("hide labels column"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`sort-by`, `If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string.`, ox.Section(0)).
				String(`subresource`, `If specified, gets the subresource of the requested object.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Bool(`watch`, `After listing/getting the requested object, watch for changes.`, ox.Spec(`false`), ox.Short("w"), ox.Section(0)).
				Bool(`watch-only`, `Watch for changes to the requested object(s), without listing/getting first.`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl edit
			ox.Usage(`edit`, `Edit a resource on the server`),
			ox.Banner(`Edit a resource from the default editor.

 The edit command allows you to directly edit any API resource you can retrieve via the command-line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows. When attempting to open the editor, it will first attempt to use the shell that has been defined in the 'SHELL' environment variable. If this is not defined, the default shell will be used, which is '/bin/bash' for Linux or 'cmd' for Windows.

 You can edit multiple objects, although changes are applied one at a time. The command accepts file names as well as command-line arguments, although the files you point to must be previously saved versions of resources.

 Editing is done with the API version used to fetch the resource. To edit using a specific API version, fully-qualify the resource, version, and group.

 The default format is YAML. To edit in JSON, specify "-o json".

 The flag --windows-line-endings can be used to force Windows line endings, otherwise the default for your operating system will be used.

 In the event an error occurs while updating, a temporary file will be created on disk that contains your unapplied changes. The most common error when updating a resource is another editor changing the resource on the server. When this occurs, you will have to apply your changes to the newer version of the resource, or update your temporary saved copy to include the latest resource version.`),
			ox.Spec(`(RESOURCE/NAME | -f FILENAME) [options]`),
			ox.Example(`
  # Edit the service named 'registry'
  kubectl edit svc/registry
  
  # Use an alternative editor
  KUBE_EDITOR="nano" kubectl edit svc/registry
  
  # Edit the job 'myjob' in JSON using the v1 API format
  kubectl edit job.v1.batch/myjob -o json
  
  # Edit the deployment 'mydeployment' in YAML and save the modified config in its annotation
  kubectl edit deployment/mydeployment -o yaml --save-config
  
  # Edit the 'status' subresource for the 'mydeployment' deployment
  kubectl edit deployment mydeployment --subresource='status'`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-edit"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files to use to edit the resource`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`output-patch`, `Output the patch if the resource is edited.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`subresource`, `If specified, edit will operate on the subresource of the requested object.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
				Bool(`windows-line-endings`, `Defaults to the line ending native to your platform.`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl delete
			ox.Usage(`delete`, `Delete resources by file names, stdin, resources and names, or by resources and label selector`),
			ox.Banner(`Delete resources by file names, stdin, resources and names, or by resources and label selector.

 JSON and YAML formats are accepted. Only one type of argument may be specified: file names, resources and names, or resources and label selector.

 Some resources, such as pods, support graceful deletion. These resources define a default period before they are forcibly terminated (the grace period) but you may override that value with the --grace-period flag, or pass --now to set a grace-period of 1. Because these resources often represent entities in the cluster, deletion may not be acknowledged immediately. If the node hosting a pod is down or cannot reach the API server, termination may take significantly longer than the grace period. To force delete a resource, you must specify the --force flag. Note: only a subset of resources support graceful deletion. In absence of the support, the --grace-period flag is ignored.

 IMPORTANT: Force deleting pods does not wait for confirmation that the pod's processes have been terminated, which can leave those processes running until the node detects the deletion and completes graceful deletion. If your processes use shared storage or talk to a remote API and depend on the name of the pod to identify themselves, force deleting those pods may result in multiple processes running on different machines using the same identification which may lead to data corruption or inconsistency. Only force delete pods when you are sure the pod is terminated, or if your application can tolerate multiple copies of the same pod running at once. Also, if you force delete pods, the scheduler may place new pods on those nodes before the node has released those resources and causing those pods to be evicted immediately.

 Note that the delete command does NOT do resource version checks, so if someone submits an update to a resource right when you submit a delete, their update will be lost along with the rest of the resource.

 After a CustomResourceDefinition is deleted, invalidation of discovery cache may take up to 6 hours. If you don't want to wait, you might want to run "kubectl api-resources" to refresh the discovery cache.`),
			ox.Spec(`([-f FILENAME] | [-k DIRECTORY] | TYPE [(NAME | -l label | --all)]) [options]`),
			ox.Example(`
  # Delete a pod using the type and name specified in pod.json
  kubectl delete -f ./pod.json
  
  # Delete resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml
  kubectl delete -k dir
  
  # Delete resources from all files that end with '.json'
  kubectl delete -f '*.json'
  
  # Delete a pod based on the type and name in the JSON passed into stdin
  cat pod.json | kubectl delete -f -
  
  # Delete pods and services with same names "baz" and "foo"
  kubectl delete pod,service baz foo
  
  # Delete pods and services with label name=myLabel
  kubectl delete pods,services -l name=myLabel
  
  # Delete a pod with minimal delay
  kubectl delete pod foo --now
  
  # Force delete a pod on a dead node
  kubectl delete pod foo --force
  
  # Delete all pods
  kubectl delete pods --all
  
  # Delete all pods only if the user confirms the deletion
  kubectl delete pods --all --interactive`),
			ox.Section(1),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Delete all resources, in the namespace of the specified resource types.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				String(`cascade`, `Must be "background", "orphan", or "foreground". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.`, ox.Default("background"), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
				Slice(`filename`, `containing the resource to delete.`, ox.Short("f"), ox.Section(0)).
				Bool(`force`, `If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.`, ox.Spec(`false`), ox.Section(0)).
				Int(`grace-period`, `Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).`, ox.Default("-1"), ox.Section(0)).
				Bool(`ignore-not-found`, `Treat "resource not found" as a successful delete. Defaults to "true" when --all is specified.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`interactive`, `If true, delete resource only when user confirms.`, ox.Spec(`false`), ox.Short("i"), ox.Section(0)).
				String(`kustomize`, `Process a kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`now`, `If true, resources are signaled for immediate shutdown (same as --grace-period=1).`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output mode. Use "-o name" for shorter output (resource/name).`, ox.Short("o"), ox.Section(0)).
				String(`raw`, `Raw URI to DELETE to the server.  Uses the transport specified by the kubeconfig file.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object`, ox.Default("0s"), ox.Section(0)).
				Bool(`wait`, `If true, wait for resources to be gone before returning. This waits for finalizers.`, ox.Spec(`true`), ox.Section(0)),
		),
		ox.Sub( // kubectl rollout
			ox.Usage(`rollout`, `Manage the rollout of a resource`),
			ox.Banner(`Manage the rollout of one or many resources.
        
 Valid resource types include:

  *  deployments
  *  daemonsets
  *  statefulsets`),
			ox.Spec(`SUBCOMMAND [options]`),
			ox.Example(`
  # Rollback to the previous deployment
  kubectl rollout undo deployment/abc
  
  # Check the rollout status of a daemonset
  kubectl rollout status daemonset/foo
  
  # Restart a deployment
  kubectl rollout restart deployment/abc
  
  # Restart deployments with the 'app=nginx' label
  kubectl rollout restart deployment --selector=app=nginx`),
			ox.Section(2),
			ox.Footer(`Use "kubectl rollout <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl rollout history
				ox.Usage(`history`, `View rollout history`),
				ox.Banner(`View previous rollout revisions and configurations.`),
				ox.Spec(`(TYPE NAME | TYPE/NAME) [flags] [options]`),
				ox.Example(`
  # View the rollout history of a deployment
  kubectl rollout history deployment/abc
  
  # View the details of daemonset revision 3
  kubectl rollout history daemonset/abc --revision=3`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Int(`revision`, `See the details, including podTemplate of the revision specified`, ox.Default("0"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl rollout pause
				ox.Usage(`pause`, `Mark the provided resource as paused`),
				ox.Banner(`Mark the provided resource as paused.

 Paused resources will not be reconciled by a controller. Use "kubectl rollout resume" to resume a paused resource. Currently only deployments support being paused.`),
				ox.Spec(`RESOURCE [options]`),
				ox.Example(`
  # Mark the nginx deployment as paused
  # Any current state of the deployment will continue its function; new updates
  # to the deployment will not have an effect as long as the deployment is paused
  kubectl rollout pause deployment/nginx`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-rollout"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl rollout restart
				ox.Usage(`restart`, `Restart a resource`),
				ox.Banner(`Restart a resource.

        Resource rollout will be restarted.`),
				ox.Spec(`RESOURCE [options]`),
				ox.Example(`
  # Restart all deployments in the test-namespace namespace
  kubectl rollout restart deployment -n test-namespace
  
  # Restart a deployment
  kubectl rollout restart deployment/nginx
  
  # Restart a daemon set
  kubectl rollout restart daemonset/abc
  
  # Restart deployments with the app=nginx label
  kubectl rollout restart deployment --selector=app=nginx`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-rollout"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl rollout resume
				ox.Usage(`resume`, `Resume a paused resource`),
				ox.Banner(`Resume a paused resource.

 Paused resources will not be reconciled by a controller. By resuming a resource, we allow it to be reconciled again. Currently only deployments support being resumed.`),
				ox.Spec(`RESOURCE [options]`),
				ox.Example(`
  # Resume an already paused deployment
  kubectl rollout resume deployment/nginx`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-rollout"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl rollout status
				ox.Usage(`status`, `Show the status of the rollout`),
				ox.Banner(`Show the status of the rollout.

 By default 'rollout status' will watch the status of the latest rollout until it's done. If you don't want to wait for the rollout to finish then you can use --watch=false. Note that if a new rollout starts in-between, then 'rollout status' will continue watching the latest revision. If you want to pin to a specific revision and abort if it is rolled over by another revision, use --revision=N where N is the revision you need to watch for.`),
				ox.Spec(`(TYPE NAME | TYPE/NAME) [flags] [options]`),
				ox.Example(`
  # Watch the rollout status of a deployment
  kubectl rollout status deployment/nginx`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Int(`revision`, `Pin to a specific revision for showing its status. Defaults to 0 (last revision).`, ox.Default("0"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Duration(`timeout`, `The length of time to wait before ending watch, zero means never. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).`, ox.Default("0s"), ox.Section(0)).
					Bool(`watch`, `Watch the status of the rollout until it's done.`, ox.Spec(`true`), ox.Short("w"), ox.Section(0)),
			),
			ox.Sub( // kubectl rollout undo
				ox.Usage(`undo`, `Undo a previous rollout`),
				ox.Banner(`Roll back to a previous rollout.`),
				ox.Spec(`(TYPE NAME | TYPE/NAME) [flags] [options]`),
				ox.Example(`
  # Roll back to the previous deployment
  kubectl rollout undo deployment/abc
  
  # Roll back to daemonset revision 3
  kubectl rollout undo daemonset/abc --to-revision=3
  
  # Roll back to the previous deployment with dry-run
  kubectl rollout undo --dry-run=server deployment/abc`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to get from a server.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					Int(`to-revision`, `The revision to rollback to. Default to 0 (last revision).`, ox.Default("0"), ox.Section(0)),
			),
		),
		ox.Sub( // kubectl scale
			ox.Usage(`scale`, `Set a new size for a deployment, replica set, or replication controller`),
			ox.Banner(`Set a new size for a deployment, replica set, replication controller, or stateful set.

 Scale also allows users to specify one or more preconditions for the scale action.

 If --current-replicas or --resource-version is specified, it is validated before the scale is attempted, and it is guaranteed that the precondition holds true when the scale is sent to the server.`),
			ox.Spec(`[--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME) [options]`),
			ox.Example(`
  # Scale a replica set named 'foo' to 3
  kubectl scale --replicas=3 rs/foo
  
  # Scale a resource identified by type and name specified in "foo.yaml" to 3
  kubectl scale --replicas=3 -f foo.yaml
  
  # If the deployment named mysql's current size is 2, scale mysql to 3
  kubectl scale --current-replicas=2 --replicas=3 deployment/mysql
  
  # Scale multiple replication controllers
  kubectl scale --replicas=5 rc/example1 rc/example2 rc/example3
  
  # Scale stateful set named 'web' to 3
  kubectl scale --replicas=3 statefulset/web`),
			ox.Section(2),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all resources in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				Int(`current-replicas`, `Precondition for current size. Requires that the current size of the resource match this value in order to scale. -1 (default) for no condition.`, ox.Default("-1"), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to set a new size`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Int(`replicas`, `The new desired number of replicas. Required.`, ox.Default("0"), ox.Section(0)).
				String(`resource-version`, `Precondition for resource version. Requires that the current resource version match this value in order to scale.`, ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up on a scale operation, zero means don't wait. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).`, ox.Default("0s"), ox.Section(0)),
		),
		ox.Sub( // kubectl autoscale
			ox.Usage(`autoscale`, `Auto-scale a deployment, replica set, stateful set, or replication controller`),
			ox.Banner(`Creates an autoscaler that automatically chooses and sets the number of pods that run in a Kubernetes cluster. The command will attempt to use the autoscaling/v2 API first, in case of an error, it will fall back to autoscaling/v1 API.

 Looks up a deployment, replica set, stateful set, or replication controller by name and creates an autoscaler that uses the given resource as a reference. An autoscaler can automatically increase or decrease number of pods deployed within the system as needed.`),
			ox.Spec(`(-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [options]`),
			ox.Example(`
  # Auto scale a deployment "foo", with the number of pods between 2 and 10, no target CPU utilization specified so a default autoscaling policy will be used
  kubectl autoscale deployment foo --min=2 --max=10
  
  # Auto scale a replication controller "foo", with the number of pods between 1 and 5, target CPU utilization at 80%
  kubectl autoscale rc foo --max=5 --cpu-percent=80`),
			ox.Section(2),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				Int(`cpu-percent`, `The target average CPU utilization (represented as a percent of requested CPU) over all the pods. If it's not specified or negative, a default autoscaling policy will be used.`, ox.Default("-1"), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-autoscale"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to autoscale.`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Int(`max`, `The upper limit for the number of pods that can be set by the autoscaler. Required.`, ox.Default("-1"), ox.Section(0)).
				Int(`min`, `The lower limit for the number of pods that can be set by the autoscaler. If it's not specified or negative, the server will apply a default value.`, ox.Default("-1"), ox.Section(0)).
				String(`name`, `The name for the newly created object. If not specified, the name of the input resource will be used.`, ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
		),
		ox.Sub( // kubectl certificate
			ox.Usage(`certificate`, `Modify certificate resources`),
			ox.Banner(`Modify certificate resources.`),
			ox.Spec(`SUBCOMMAND [options]`),
			ox.Section(3),
			ox.Footer(`Use "kubectl certificate <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl certificate approve
				ox.Usage(`approve`, `Approve a certificate signing request`),
				ox.Banner(`Approve a certificate signing request.

 kubectl certificate approve allows a cluster admin to approve a certificate signing request (CSR). This action tells a certificate signing controller to issue a certificate to the requester with the attributes requested in the CSR.

 SECURITY NOTICE: Depending on the requested attributes, the issued certificate can potentially grant a requester access to cluster resources or to authenticate as a requested identity. Before approving a CSR, ensure you understand what the signed certificate can do.`),
				ox.Spec(`(-f FILENAME | NAME) [options]`),
				ox.Example(`
  # Approve CSR 'csr-sqgzp'
  kubectl certificate approve csr-sqgzp`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to update`, ox.Short("f"), ox.Section(0)).
					Bool(`force`, `Update the CSR even if it is already approved.`, ox.Spec(`false`), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl certificate deny
				ox.Usage(`deny`, `Deny a certificate signing request`),
				ox.Banner(`Deny a certificate signing request.

 kubectl certificate deny allows a cluster admin to deny a certificate signing request (CSR). This action tells a certificate signing controller to not to issue a certificate to the requester.`),
				ox.Spec(`(-f FILENAME | NAME) [options]`),
				ox.Example(`
  # Deny CSR 'csr-sqgzp'
  kubectl certificate deny csr-sqgzp`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to update`, ox.Short("f"), ox.Section(0)).
					Bool(`force`, `Update the CSR even if it is already denied.`, ox.Spec(`false`), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl cluster-info
			ox.Usage(`cluster-info`, `Display cluster information`),
			ox.Banner(`Display addresses of the control plane and services with label kubernetes.io/cluster-service=true. To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.`),
			ox.Spec(`[flags] [options]`),
			ox.Example(`
  # Print the address of the control plane and cluster services
  kubectl cluster-info`),
			ox.Section(3),
			ox.Footer(`Use "kubectl cluster-info <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl cluster-info dump
				ox.Usage(`dump`, `Dump relevant information for debugging and diagnosis`),
				ox.Banner(`Dump cluster information out suitable for debugging and diagnosing cluster problems.  By default, dumps everything to stdout. You can optionally specify a directory with --output-directory.  If you specify a directory, Kubernetes will build a set of files in that directory.  By default, only dumps things in the current namespace and 'kube-system' namespace, but you can switch to a different namespace with the --namespaces flag, or specify --all-namespaces to dump all namespaces.

 The command also dumps the logs of all of the pods in the cluster; these logs are dumped into different directories based on namespace and pod name.`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # Dump current cluster state to stdout
  kubectl cluster-info dump
  
  # Dump current cluster state to /path/to/cluster-state
  kubectl cluster-info dump --output-directory=/path/to/cluster-state
  
  # Dump all namespaces to stdout
  kubectl cluster-info dump --all-namespaces
  
  # Dump a set of namespaces to /path/to/cluster-state
  kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all-namespaces`, `If true, dump all namespaces.  If true, --namespaces is ignored.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Slice(`namespaces`, `A comma separated list of namespaces to dump.`, ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Default("json"), ox.Short("o"), ox.Section(0)).
					String(`output-directory`, `Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory`, ox.Section(0)).
					Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("20s"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl top
			ox.Usage(`top`, `Display resource (CPU/memory) usage`),
			ox.Banner(`Display resource (CPU/memory) usage.

 The top command allows you to see the resource consumption for nodes or pods.

 This command requires Metrics Server to be correctly configured and working on the server.`),
			ox.Spec(`[flags] [options]`),
			ox.Section(3),
			ox.Footer(`Use "kubectl top <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl top node
				ox.Usage(`node`, `Display resource (CPU/memory) usage of nodes`),
				ox.Banner(`Display resource (CPU/memory) usage of nodes.

 The top-node command allows you to see the resource consumption of nodes.`),
				ox.Spec(`[NAME | -l label] [options]`),
				ox.Aliases("nodes", "no"),
				ox.Example(`
  # Show metrics for all nodes
  kubectl top node
  
  # Show metrics for a given node
  kubectl top node NODE_NAME`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`no-headers`, `If present, print output without headers`, ox.Spec(`false`), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					Bool(`show-capacity`, `Print node resources based on Capacity instead of Allocatable(default) of the nodes.`, ox.Spec(`false`), ox.Section(0)).
					String(`sort-by`, `If non-empty, sort nodes list using specified field. The field can be either 'cpu' or 'memory'.`, ox.Section(0)).
					Bool(`use-protocol-buffers`, `Enables using protocol-buffers to access Metrics API.`, ox.Spec(`true`), ox.Section(0)),
			),
			ox.Sub( // kubectl top pod
				ox.Usage(`pod`, `Display resource (CPU/memory) usage of pods`),
				ox.Banner(`Display resource (CPU/memory) usage of pods.

 The 'top pod' command allows you to see the resource consumption of pods.

 Due to the metrics pipeline delay, they may be unavailable for a few minutes since pod creation.`),
				ox.Spec(`[NAME | -l label] [options]`),
				ox.Aliases("pods", "po"),
				ox.Example(`
  # Show metrics for all pods in the default namespace
  kubectl top pod
  
  # Show metrics for all pods in the given namespace
  kubectl top pod --namespace=NAMESPACE
  
  # Show metrics for a given pod and its containers
  kubectl top pod POD_NAME --containers
  
  # Show metrics for the pods defined by label name=myLabel
  kubectl top pod -l name=myLabel`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
					Bool(`containers`, `If present, print usage of containers within a pod.`, ox.Spec(`false`), ox.Section(0)).
					String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
					Bool(`no-headers`, `If present, print output without headers.`, ox.Spec(`false`), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
					String(`sort-by`, `If non-empty, sort pods list using specified field. The field can be either 'cpu' or 'memory'.`, ox.Section(0)).
					Bool(`sum`, `Print the sum of the resource usage`, ox.Spec(`false`), ox.Section(0)).
					Bool(`use-protocol-buffers`, `Enables using protocol-buffers to access Metrics API.`, ox.Spec(`true`), ox.Section(0)),
			),
		),
		ox.Sub( // kubectl cordon
			ox.Usage(`cordon`, `Mark node as unschedulable`),
			ox.Banner(`Mark node as unschedulable.`),
			ox.Spec(`NODE [options]`),
			ox.Example(`
  # Mark node "foo" as unschedulable
  kubectl cordon foo`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)),
		),
		ox.Sub( // kubectl uncordon
			ox.Usage(`uncordon`, `Mark node as schedulable`),
			ox.Banner(`Mark node as schedulable.`),
			ox.Spec(`NODE [options]`),
			ox.Example(`
  # Mark node "foo" as schedulable
  kubectl uncordon foo`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)),
		),
		ox.Sub( // kubectl drain
			ox.Usage(`drain`, `Drain node in preparation for maintenance`),
			ox.Banner(`Drain node in preparation for maintenance.

 The given node will be marked unschedulable to prevent new pods from arriving. 'drain' evicts the pods if the API server supports https://kubernetes.io/docs/concepts/workloads/pods/disruptions/ eviction https://kubernetes.io/docs/concepts/workloads/pods/disruptions/ . Otherwise, it will use normal DELETE to delete the pods. The 'drain' evicts or deletes all pods except mirror pods (which cannot be deleted through the API server).  If there are daemon set-managed pods, drain will not proceed without --ignore-daemonsets, and regardless it will not delete any daemon set-managed pods, because those pods would be immediately replaced by the daemon set controller, which ignores unschedulable markings.  If there are any pods that are neither mirror pods nor managed by a replication controller, replica set, daemon set, stateful set, or job, then drain will not delete any pods unless you use --force.  --force will also allow deletion to proceed if the managing resource of one or more pods is missing.

 'drain' waits for graceful termination. You should not operate on the machine until the command completes.

 When you are ready to put the node back into service, use kubectl uncordon, which will make the node schedulable again.

https://kubernetes.io/images/docs/kubectl_drain.svg Workflowhttps://kubernetes.io/images/docs/kubectl_drain.svg`),
			ox.Spec(`NODE [options]`),
			ox.Example(`
  # Drain node "foo", even if there are pods not managed by a replication controller, replica set, job, daemon set, or stateful set on it
  kubectl drain foo --force
  
  # As above, but abort if there are pods not managed by a replication controller, replica set, job, daemon set, or stateful set, and use a grace period of 15 minutes
  kubectl drain foo --grace-period=900`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Int(`chunk-size`, `Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.`, ox.Default("500"), ox.Section(0)).
				Bool(`delete-emptydir-data`, `Continue even if there are pods using emptyDir (local data that will be deleted when the node is drained).`, ox.Spec(`false`), ox.Section(0)).
				Bool(`disable-eviction`, `Force drain to use delete, even if eviction is supported. This will bypass checking PodDisruptionBudgets, use with caution.`, ox.Spec(`false`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				Bool(`force`, `Continue even if there are pods that do not declare a controller.`, ox.Spec(`false`), ox.Section(0)).
				Int(`grace-period`, `Period of time in seconds given to each pod to terminate gracefully. If negative, the default value specified in the pod will be used.`, ox.Default("-1"), ox.Section(0)).
				Bool(`ignore-daemonsets`, `Ignore DaemonSet-managed pods.`, ox.Spec(`false`), ox.Section(0)).
				String(`pod-selector`, `Label selector to filter pods on the node`, ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Int(`skip-wait-for-delete-timeout`, `If pod DeletionTimestamp older than N seconds, skip waiting for the pod.  Seconds must be greater than 0 to skip.`, ox.Default("0"), ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up, zero means infinite`, ox.Default("0s"), ox.Section(0)),
		),
		ox.Sub( // kubectl taint
			ox.Usage(`taint`, `Update the taints on one or more nodes`),
			ox.Banner(`Update the taints on one or more nodes.

  *  A taint consists of a key, value, and effect. As an argument here, it is expressed as key=value:effect.
  *  The key must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 253 characters.
  *  Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app.
  *  The value is optional. If given, it must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters.
  *  The effect must be NoSchedule, PreferNoSchedule or NoExecute.
  *  Currently taint can only apply to node.`),
			ox.Spec(`NODE NAME KEY_1=VAL_1:TAINT_EFFECT_1 ... KEY_N=VAL_N:TAINT_EFFECT_N [options]`),
			ox.Example(`
  # Update node 'foo' with a taint with key 'dedicated' and value 'special-user' and effect 'NoSchedule'
  # If a taint with that key and effect already exists, its value is replaced as specified
  kubectl taint nodes foo dedicated=special-user:NoSchedule
  
  # Remove from node 'foo' the taint with key 'dedicated' and effect 'NoSchedule' if one exists
  kubectl taint nodes foo dedicated:NoSchedule-
  
  # Remove from node 'foo' all the taints with key 'dedicated'
  kubectl taint nodes foo dedicated-
  
  # Add a taint with key 'dedicated' on nodes having label myLabel=X
  kubectl taint node -l myLabel=X  dedicated=foo:PreferNoSchedule
  
  # Add to node 'foo' a taint with key 'bar' and no value
  kubectl taint nodes foo bar:NoSchedule`),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all nodes in the cluster`, ox.Spec(`false`), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-taint"), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`overwrite`, `If true, allow taints to be overwritten, otherwise reject taint updates that overwrite existing taints.`, ox.Spec(`false`), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)),
		),
		ox.Sub( // kubectl describe
			ox.Usage(`describe`, `Show details of a specific resource or group of resources`),
			ox.Banner(`Show details of a specific resource or group of resources.

 Print a detailed description of the selected resources, including related resources such as events or controllers. You may select a single object by name, all objects of that type, provide a name prefix, or label selector. For example:

        $ kubectl describe TYPE NAME_PREFIX
        
 will first check for an exact match on TYPE and NAME_PREFIX. If no such resource exists, it will output details for every resource that has a name prefixed with NAME_PREFIX.

Use "kubectl api-resources" for a complete list of supported resources.`),
			ox.Spec(`(-f FILENAME | TYPE [NAME_PREFIX | -l label] | TYPE/NAME) [options]`),
			ox.Example(`
  # Describe a node
  kubectl describe nodes kubernetes-node-emt8.c.myproject.internal
  
  # Describe a pod
  kubectl describe pods/nginx
  
  # Describe a pod identified by type and name in "pod.json"
  kubectl describe -f pod.json
  
  # Describe all pods
  kubectl describe pods
  
  # Describe pods by label name=myLabel
  kubectl describe pods -l name=myLabel
  
  # Describe all pods managed by the 'frontend' replication controller
  # (rc-created pods get the name of the rc as a prefix in the pod name)
  kubectl describe pods frontend`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Int(`chunk-size`, `Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.`, ox.Default("500"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files containing the resource to describe`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-events`, `If true, display events related to the described object.`, ox.Spec(`true`), ox.Section(0)),
		),
		ox.Sub( // kubectl logs
			ox.Usage(`logs`, `Print the logs for a container in a pod`),
			ox.Banner(`Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name is optional.`),
			ox.Spec(`[-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [options]`),
			ox.Example(`
  # Return snapshot logs from pod nginx with only one container
  kubectl logs nginx
  
  # Return snapshot logs from pod nginx, prefixing each line with the source pod and container name
  kubectl logs nginx --prefix
  
  # Return snapshot logs from pod nginx, limiting output to 500 bytes
  kubectl logs nginx --limit-bytes=500
  
  # Return snapshot logs from pod nginx, waiting up to 20 seconds for it to start running.
  kubectl logs nginx --pod-running-timeout=20s
  
  # Return snapshot logs from pod nginx with multi containers
  kubectl logs nginx --all-containers=true
  
  # Return snapshot logs from all pods in the deployment nginx
  kubectl logs deployment/nginx --all-pods=true
  
  # Return snapshot logs from all containers in pods defined by label app=nginx
  kubectl logs -l app=nginx --all-containers=true
  
  # Return snapshot logs from all pods defined by label app=nginx, limiting concurrent log requests to 10 pods
  kubectl logs -l app=nginx --max-log-requests=10
  
  # Return snapshot of previous terminated ruby container logs from pod web-1
  kubectl logs -p -c ruby web-1
  
  # Begin streaming the logs from pod nginx, continuing even if errors occur
  kubectl logs nginx -f --ignore-errors=true
  
  # Begin streaming the logs of the ruby container in pod web-1
  kubectl logs -f -c ruby web-1
  
  # Begin streaming the logs from all containers in pods defined by label app=nginx
  kubectl logs -f -l app=nginx --all-containers=true
  
  # Display only the most recent 20 lines of output in pod nginx
  kubectl logs --tail=20 nginx
  
  # Show all logs from pod nginx written in the last hour
  kubectl logs --since=1h nginx
  
  # Show all logs with timestamps from pod nginx starting from August 30, 2024, at 06:00:00 UTC
  kubectl logs nginx --since-time=2024-08-30T06:00:00Z --timestamps=true
  
  # Show logs from a kubelet with an expired serving certificate
  kubectl logs --insecure-skip-tls-verify-backend nginx
  
  # Return snapshot logs from first container of a job named hello
  kubectl logs job/hello
  
  # Return snapshot logs from container nginx-1 of a deployment named nginx
  kubectl logs deployment/nginx -c nginx-1`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all-containers`, `Get all containers' logs in the pod(s).`, ox.Spec(`false`), ox.Section(0)).
				Bool(`all-pods`, `Get logs from all pod(s). Sets prefix to true.`, ox.Spec(`false`), ox.Section(0)).
				String(`container`, `Print the logs of this container`, ox.Short("c"), ox.Section(0)).
				Bool(`follow`, `Specify if the logs should be streamed.`, ox.Spec(`false`), ox.Short("f"), ox.Section(0)).
				Bool(`ignore-errors`, `If watching / following pod logs, allow for any errors that occur to be non-fatal`, ox.Spec(`false`), ox.Section(0)).
				Bool(`insecure-skip-tls-verify-backend`, `Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker could provide invalid log content back. You might want to use this if your kubelet serving certificates have expired.`, ox.Spec(`false`), ox.Section(0)).
				Int(`limit-bytes`, `Maximum bytes of logs to return. Defaults to no limit.`, ox.Default("0"), ox.Section(0)).
				Int(`max-log-requests`, `Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5.`, ox.Default("5"), ox.Section(0)).
				Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("20s"), ox.Section(0)).
				Bool(`prefix`, `Prefix each log line with the log source (pod name and container name)`, ox.Spec(`false`), ox.Section(0)).
				Bool(`previous`, `If true, print the logs for the previous instance of the container in a pod if it exists.`, ox.Spec(`false`), ox.Short("p"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Duration(`since`, `Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.`, ox.Default("0s"), ox.Section(0)).
				String(`since-time`, `Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.`, ox.Section(0)).
				Int(`tail`, `Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.`, ox.Default("-1"), ox.Section(0)).
				Bool(`timestamps`, `Include timestamps on each line in the log output`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl attach
			ox.Usage(`attach`, `Attach to a running container`),
			ox.Banner(`Attach to a process that is already running inside an existing container.`),
			ox.Spec(`(POD | TYPE/NAME) -c CONTAINER [options]`),
			ox.Example(`
  # Get output from running pod mypod; use the 'kubectl.kubernetes.io/default-container' annotation
  # for selecting the container to be attached or the first container in the pod will be chosen
  kubectl attach mypod
  
  # Get output from ruby-container from pod mypod
  kubectl attach mypod -c ruby-container
  
  # Switch to raw terminal mode; sends stdin to 'bash' in ruby-container from pod mypod
  # and sends stdout/stderr from 'bash' back to the client
  kubectl attach mypod -c ruby-container -i -t
  
  # Get output from the first pod of a replica set named nginx
  kubectl attach rs/nginx`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`container`, `Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen`, ox.Short("c"), ox.Section(0)).
				Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("1m0s"), ox.Section(0)).
				Bool(`quiet`, `Only print output from the remote session`, ox.Spec(`false`), ox.Short("q"), ox.Section(0)).
				Bool(`stdin`, `Pass stdin to the container`, ox.Spec(`false`), ox.Short("i"), ox.Section(0)).
				Bool(`tty`, `Stdin is a TTY`, ox.Spec(`false`), ox.Short("t"), ox.Section(0)),
		),
		ox.Sub( // kubectl exec
			ox.Usage(`exec`, `Execute a command in a container`),
			ox.Banner(`Execute a command in a container.`),
			ox.Spec(`(POD | TYPE/NAME) [-c CONTAINER] [flags] -- COMMAND [args...] [options]`),
			ox.Example(`
  # Get output from running the 'date' command from pod mypod, using the first container by default
  kubectl exec mypod -- date
  
  # Get output from running the 'date' command in ruby-container from pod mypod
  kubectl exec mypod -c ruby-container -- date
  
  # Switch to raw terminal mode; sends stdin to 'bash' in ruby-container from pod mypod
  # and sends stdout/stderr from 'bash' back to the client
  kubectl exec mypod -c ruby-container -i -t -- bash -il
  
  # List contents of /usr from the first container of pod mypod and sort by modification time
  # If the command you want to execute in the pod has any flags in common (e.g. -i),
  # you must use two dashes (--) to separate your command's flags/arguments
  # Also note, do not surround your command and its flags/arguments with quotes
  # unless that is how you would execute it normally (i.e., do ls -t /usr, not "ls -t /usr")
  kubectl exec mypod -i -t -- ls -t /usr
  
  # Get output from running 'date' command from the first pod of the deployment mydeployment, using the first container by default
  kubectl exec deploy/mydeployment -- date
  
  # Get output from running 'date' command from the first pod of the service myservice, using the first container by default
  kubectl exec svc/myservice -- date`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`container`, `Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen`, ox.Short("c"), ox.Section(0)).
				Slice(`filename`, `to use to exec into the resource`, ox.Short("f"), ox.Section(0)).
				Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("1m0s"), ox.Section(0)).
				Bool(`quiet`, `Only print output from the remote session`, ox.Spec(`false`), ox.Short("q"), ox.Section(0)).
				Bool(`stdin`, `Pass stdin to the container`, ox.Spec(`false`), ox.Short("i"), ox.Section(0)).
				Bool(`tty`, `Stdin is a TTY`, ox.Spec(`false`), ox.Short("t"), ox.Section(0)),
		),
		ox.Sub( // kubectl port-forward
			ox.Usage(`port-forward`, `Forward one or more local ports to a pod`),
			ox.Banner(`Forward one or more local ports to a pod.

 Use resource type/name such as deployment/mydeployment to select a pod. Resource type defaults to 'pod' if omitted.

 If there are multiple pods matching the criteria, a pod will be selected automatically. The forwarding session ends when the selected pod terminates, and a rerun of the command is needed to resume forwarding.`),
			ox.Spec(`TYPE/NAME [options] [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N]`),
			ox.Example(`
  # Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in the pod
  kubectl port-forward pod/mypod 5000 6000
  
  # Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in a pod selected by the deployment
  kubectl port-forward deployment/mydeployment 5000 6000
  
  # Listen on port 8443 locally, forwarding to the targetPort of the service's port named "https" in a pod selected by the service
  kubectl port-forward service/myservice 8443:https
  
  # Listen on port 8888 locally, forwarding to 5000 in the pod
  kubectl port-forward pod/mypod 8888:5000
  
  # Listen on port 8888 on all addresses, forwarding to 5000 in the pod
  kubectl port-forward --address 0.0.0.0 pod/mypod 8888:5000
  
  # Listen on port 8888 on localhost and selected IP, forwarding to 5000 in the pod
  kubectl port-forward --address localhost,10.19.21.23 pod/mypod 8888:5000
  
  # Listen on a random port locally, forwarding to 5000 in the pod
  kubectl port-forward pod/mypod :5000`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Slice(`address`, `Addresses to listen on (comma separated). Only accepts IP addresses or localhost as a value. When localhost is supplied, kubectl will try to bind on both 127.0.0.1 and ::1 and will fail if neither of these addresses are available to bind.`, ox.Default("localhost"), ox.Section(0)).
				Duration(`pod-running-timeout`, `The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running`, ox.Default("1m0s"), ox.Section(0)),
		),
		ox.Sub( // kubectl proxy
			ox.Usage(`proxy`, `Run a proxy to the Kubernetes API server`),
			ox.Banner(`Creates a proxy server or application-level gateway between localhost and the Kubernetes API server. It also allows serving static content over specified HTTP path. All incoming data enters through one port and gets forwarded to the remote Kubernetes API server port, except for the path matching the static content path.`),
			ox.Spec(`[--port=PORT] [--www=static-dir] [--www-prefix=prefix] [--api-prefix=prefix] [options]`),
			ox.Example(`
  # To proxy all of the Kubernetes API and nothing else
  kubectl proxy --api-prefix=/
  
  # To proxy only part of the Kubernetes API and also some static files
  # You can get pods info with 'curl localhost:8001/api/v1/pods'
  kubectl proxy --www=/my/files --www-prefix=/static/ --api-prefix=/api/
  
  # To proxy the entire Kubernetes API at a different root
  # You can get pods info with 'curl localhost:8001/custom/api/v1/pods'
  kubectl proxy --api-prefix=/custom/
  
  # Run a proxy to the Kubernetes API server on port 8011, serving static content from ./local/www/
  kubectl proxy --port=8011 --www=./local/www/
  
  # Run a proxy to the Kubernetes API server on an arbitrary local port
  # The chosen port for the server will be output to stdout
  kubectl proxy --port=0
  
  # Run a proxy to the Kubernetes API server, changing the API prefix to k8s-api
  # This makes e.g. the pods API available at localhost:8001/k8s-api/v1/pods/
  kubectl proxy --api-prefix=/k8s-api`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`accept-hosts`, `Regular expression for hosts that the proxy should accept.`, ox.Spec(`regexp`), ox.Default("^localhost$,^127\\.0\\.0\\.1$,^\\[::1\\]$"), ox.Section(0)).
				String(`accept-paths`, `Regular expression for paths that the proxy should accept.`, ox.Spec(`regexp`), ox.Default("^.*"), ox.Section(0)).
				Addr(`address`, `The IP address on which to serve on.`, ox.Default("127.0.0.1"), ox.Section(0)).
				String(`api-prefix`, `Prefix to serve the proxied API under.`, ox.Spec(`path`), ox.Default("/"), ox.Section(0)).
				Bool(`append-server-path`, `If true, enables automatic path appending of the kube context server path to each request.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`disable-filter`, `If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to XSRF attacks, when used with an accessible port.`, ox.Spec(`false`), ox.Section(0)).
				Duration(`keepalive`, `keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keepalive.`, ox.Default("0s"), ox.Section(0)).
				Int(`port`, `The port on which to run the proxy. Set to 0 to pick a random port.`, ox.Default("8001"), ox.Short("p"), ox.Section(0)).
				String(`reject-methods`, `Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT,PATCH').`, ox.Spec(`regexp`), ox.Default("^$"), ox.Section(0)).
				String(`reject-paths`, `Regular expression for paths that the proxy should reject. Paths specified here will be rejected even accepted by --accept-paths.`, ox.Spec(`regexp`), ox.Default("^/api/.*/pods/.*/exec,^/api/.*/pods/.*/attach"), ox.Section(0)).
				String(`unix-socket`, `Unix socket on which to run the proxy.`, ox.Short("u"), ox.Section(0)).
				String(`www`, `Also serve static files from the given directory under the specified prefix.`, ox.Short("w"), ox.Section(0)).
				String(`www-prefix`, `Prefix to serve static files under, if static file directory is specified.`, ox.Spec(`path`), ox.Default("/static/"), ox.Short("P"), ox.Section(0)),
		),
		ox.Sub( // kubectl cp
			ox.Usage(`cp`, `Copy files and directories to and from containers`),
			ox.Banner(`Copy files and directories to and from containers.`),
			ox.Spec(`<file-spec-src> <file-spec-dest> [options]`),
			ox.Example(`
  # !!!Important Note!!!
  # Requires that the 'tar' binary is present in your container
  # image.  If 'tar' is not present, 'kubectl cp' will fail.
  #
  # For advanced use cases, such as symlinks, wildcard expansion or
  # file mode preservation, consider using 'kubectl exec'.
  
  # Copy /tmp/foo local file to /tmp/bar in a remote pod in namespace <some-namespace>
  tar cf - /tmp/foo | kubectl exec -i -n <some-namespace> <some-pod> -- tar xf - -C /tmp/bar
  
  # Copy /tmp/foo from a remote pod to /tmp/bar locally
  kubectl exec -n <some-namespace> <some-pod> -- tar cf - /tmp/foo | tar xf - -C /tmp/bar
  
  # Copy /tmp/foo_dir local directory to /tmp/bar_dir in a remote pod in the default namespace
  kubectl cp /tmp/foo_dir <some-pod>:/tmp/bar_dir
  
  # Copy /tmp/foo local file to /tmp/bar in a remote pod in a specific container
  kubectl cp /tmp/foo <some-pod>:/tmp/bar -c <specific-container>
  
  # Copy /tmp/foo local file to /tmp/bar in a remote pod in namespace <some-namespace>
  kubectl cp /tmp/foo <some-namespace>/<some-pod>:/tmp/bar
  
  # Copy /tmp/foo from a remote pod to /tmp/bar locally
  kubectl cp <some-namespace>/<some-pod>:/tmp/foo /tmp/bar`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`container`, `Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen`, ox.Short("c"), ox.Section(0)).
				Bool(`no-preserve`, `The copied file/directory's ownership and permissions will not be preserved in the container`, ox.Spec(`false`), ox.Section(0)).
				Int(`retries`, `Set number of retries to complete a copy operation from a container. Specify 0 to disable or any negative value for infinite retrying. The default is 0 (no retry).`, ox.Default("0"), ox.Section(0)),
		),
		ox.Sub( // kubectl auth
			ox.Usage(`auth`, `Inspect authorization`),
			ox.Banner(`Inspect authorization.`),
			ox.Spec(`[flags] [options]`),
			ox.Section(4),
			ox.Footer(`Use "kubectl auth <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl auth can-i
				ox.Usage(`can-i`, `Check whether an action is allowed`),
				ox.Banner(`Check whether an action is allowed.

 VERB is a logical Kubernetes API verb like 'get', 'list', 'watch', 'delete', etc. TYPE is a Kubernetes resource. Shortcuts and groups will be resolved. NONRESOURCEURL is a partial URL that starts with "/". NAME is the name of a particular Kubernetes resource. This command pairs nicely with impersonation. See --as global flag.`),
				ox.Spec(`VERB [TYPE | TYPE/NAME | NONRESOURCEURL] [options]`),
				ox.Example(`
  # Check to see if I can create pods in any namespace
  kubectl auth can-i create pods --all-namespaces
  
  # Check to see if I can list deployments in my current namespace
  kubectl auth can-i list deployments.apps
  
  # Check to see if service account "foo" of namespace "dev" can list pods in the namespace "prod"
  # You must be allowed to use impersonation for the global option "--as"
  kubectl auth can-i list pods --as=system:serviceaccount:dev:foo -n prod
  
  # Check to see if I can do everything in my current namespace ("*" means all)
  kubectl auth can-i '*' '*'
  
  # Check to see if I can get the job named "bar" in namespace "foo"
  kubectl auth can-i list jobs.batch/bar -n foo
  
  # Check to see if I can read pod logs
  kubectl auth can-i get pods --subresource=log
  
  # Check to see if I can access the URL /logs/
  kubectl auth can-i get /logs/
  
  # Check to see if I can approve certificates.k8s.io
  kubectl auth can-i approve certificates.k8s.io
  
  # List all allowed actions in namespace "foo"
  kubectl auth can-i --list --namespace=foo`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all-namespaces`, `If true, check the specified action in all namespaces.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
					Bool(`list`, `If true, prints all allowed actions.`, ox.Spec(`false`), ox.Section(0)).
					Bool(`no-headers`, `If true, prints allowed actions without headers`, ox.Spec(`false`), ox.Section(0)).
					Bool(`quiet`, `If true, suppress output and just return the exit code.`, ox.Spec(`false`), ox.Short("q"), ox.Section(0)).
					String(`subresource`, `SubResource such as pod/log or deployment/scale`, ox.Section(0)),
			),
			ox.Sub( // kubectl auth reconcile
				ox.Usage(`reconcile`, `Reconciles rules for RBAC role, role binding, cluster role, and cluster role binding objects`),
				ox.Banner(`Reconciles rules for RBAC role, role binding, cluster role, and cluster role binding objects.

 Missing objects are created, and the containing namespace is created for namespaced objects, if required.

 Existing roles are updated to include the permissions in the input objects, and remove extra permissions if --remove-extra-permissions is specified.

 Existing bindings are updated to include the subjects in the input objects, and remove extra subjects if --remove-extra-subjects is specified.

 This is preferred to 'apply' for RBAC resources so that semantically-aware merging of rules and subjects is done.`),
				ox.Spec(`-f FILENAME [options]`),
				ox.Example(`
  # Reconcile RBAC resources from a file
  kubectl auth reconcile -f my-rbac-rules.yaml`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files identifying the resource to reconcile.`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`remove-extra-permissions`, `If true, removes extra permissions added to roles`, ox.Spec(`false`), ox.Section(0)).
					Bool(`remove-extra-subjects`, `If true, removes extra subjects added to rolebindings`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl auth whoami
				ox.Usage(`whoami`, `Experimental: Check self subject attributes`),
				ox.Banner(`Check who you are and your attributes (groups, extra).

        This command is helpful to get yourself aware of the current user attributes,
        especially when dynamic authentication, e.g., token webhook, auth proxy, or OIDC provider,
        is enabled in the Kuberne`),
				ox.Spec(`[options]`),
				ox.Example(`
  # Get your subject attributes
  kubectl auth whoami
  
  # Get your subject attributes in JSON format
  kubectl auth whoami -o json`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl debug
			ox.Usage(`debug`, `Create debugging sessions for troubleshooting workloads and nodes`),
			ox.Banner(`Debug cluster resources using interactive debugging containers.

 'debug' provides automation for common debugging tasks for cluster objects identified by resource and name. Pods will be used by default if no resource is specified.

 The action taken by 'debug' varies depending on what resource is specified. Supported actions include:

  *  Workload: Create a copy of an existing pod with certain attributes changed, for example changing the image tag to a new version.
  *  Workload: Add an ephemeral container to an already running pod, for example to add debugging utilities without restarting the pod.
  *  Node: Create a new pod that runs in the node's host namespaces and can access the node's filesystem.

 Note: When a non-root user is configured for the entire target Pod, some capabilities granted by debug profile may not work.`),
			ox.Spec(`(POD | TYPE[[.VERSION].GROUP]/NAME) [ -- COMMAND [args...] ] [options]`),
			ox.Example(`
  # Create an interactive debugging session in pod mypod and immediately attach to it.
  kubectl debug mypod -it --image=busybox
  
  # Create an interactive debugging session for the pod in the file pod.yaml and immediately attach to it.
  # (requires the EphemeralContainers feature to be enabled in the cluster)
  kubectl debug -f pod.yaml -it --image=busybox
  
  # Create a debug container named debugger using a custom automated debugging image.
  kubectl debug --image=myproj/debug-tools -c debugger mypod
  
  # Create a copy of mypod adding a debug container and attach to it
  kubectl debug mypod -it --image=busybox --copy-to=my-debugger
  
  # Create a copy of mypod changing the command of mycontainer
  kubectl debug mypod -it --copy-to=my-debugger --container=mycontainer -- sh
  
  # Create a copy of mypod changing all container images to busybox
  kubectl debug mypod --copy-to=my-debugger --set-image=*=busybox
  
  # Create a copy of mypod adding a debug container and changing container images
  kubectl debug mypod -it --copy-to=my-debugger --image=debian --set-image=app=app:debug,sidecar=sidecar:debug
  
  # Create an interactive debugging session on a node and immediately attach to it.
  # The container will run in the host namespaces and the host's filesystem will be mounted at /host
  kubectl debug node/mynode -it --image=busybox`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`arguments-only`, `If specified, everything after -- will be passed to the new container as Args instead of Command.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`attach`, `If true, wait for the container to start running, and then attach as if 'kubectl attach ...' were called.  Default false, unless '-i/--stdin' is set, in which case the default is true.`, ox.Spec(`false`), ox.Section(0)).
				String(`container`, `Container name to use for debug container.`, ox.Short("c"), ox.Section(0)).
				String(`copy-to`, `Create a copy of the target Pod with this name.`, ox.Section(0)).
				String(`custom`, `Path to a JSON or YAML file containing a partial container spec to customize built-in debug profiles.`, ox.Section(0)).
				Slice(`env`, `Environment variables to set in the container.`, ox.Section(0)).
				Slice(`filename`, `identifying the resource to debug`, ox.Short("f"), ox.Section(0)).
				String(`image`, `Container image to use for debug container.`, ox.Section(0)).
				String(`image-pull-policy`, `The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server.`, ox.Section(0)).
				Bool(`keep-annotations`, `If true, keep the original pod annotations.(This flag only works when used with '--copy-to')`, ox.Spec(`false`), ox.Section(0)).
				Bool(`keep-init-containers`, `Run the init containers for the pod. Defaults to true.(This flag only works when used with '--copy-to')`, ox.Spec(`true`), ox.Section(0)).
				Bool(`keep-labels`, `If true, keep the original pod labels.(This flag only works when used with '--copy-to')`, ox.Spec(`false`), ox.Section(0)).
				Bool(`keep-liveness`, `If true, keep the original pod liveness probes.(This flag only works when used with '--copy-to')`, ox.Spec(`false`), ox.Section(0)).
				Bool(`keep-readiness`, `If true, keep the original pod readiness probes.(This flag only works when used with '--copy-to')`, ox.Spec(`false`), ox.Section(0)).
				Bool(`keep-startup`, `If true, keep the original startup probes.(This flag only works when used with '--copy-to')`, ox.Spec(`false`), ox.Section(0)).
				String(`profile`, `Options are "legacy", "general", "baseline", "netadmin", "restricted" or "sysadmin".`, ox.Default("legacy"), ox.Section(0)).
				Bool(`quiet`, `If true, suppress informational messages.`, ox.Spec(`false`), ox.Short("q"), ox.Section(0)).
				Bool(`replace`, `When used with '--copy-to', delete the original Pod.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`same-node`, `When used with '--copy-to', schedule the copy of target Pod on the same node.`, ox.Spec(`false`), ox.Section(0)).
				Slice(`set-image`, `When used with '--copy-to', a list of name=image pairs for changing container images, similar to how 'kubectl set image' works.`, ox.Section(0)).
				Bool(`share-processes`, `When used with '--copy-to', enable process namespace sharing in the copy.`, ox.Spec(`true`), ox.Section(0)).
				Bool(`stdin`, `Keep stdin open on the container(s) in the pod, even if nothing is attached.`, ox.Spec(`false`), ox.Short("i"), ox.Section(0)).
				String(`target`, `When using an ephemeral container, target processes in this container name.`, ox.Section(0)).
				Bool(`tty`, `Allocate a TTY for the debugging container.`, ox.Spec(`false`), ox.Short("t"), ox.Section(0)),
		),
		ox.Sub( // kubectl events
			ox.Usage(`events`, `List events`),
			ox.Banner(`Display events.

 Prints a table of the most important information about events. You can request events for a namespace, for all namespace, or filtered to only those pertaining to a specified resource.`),
			ox.Spec(`[(-o|--output=)json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file] [--for TYPE/NAME] [--watch] [--types=Normal,Warning] [options]`),
			ox.Example(`
  # List recent events in the default namespace
  kubectl events
  
  # List recent events in all namespaces
  kubectl events --all-namespaces
  
  # List recent events for the specified pod, then wait for more events and list them as they arrive
  kubectl events --for pod/web-pod-13je7 --watch
  
  # List recent events in YAML format
  kubectl events -oyaml
  
  # List recent only events of type 'Warning' or 'Normal'
  kubectl events --types=Warning,Normal`),
			ox.Section(4),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				Int(`chunk-size`, `Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.`, ox.Default("500"), ox.Section(0)).
				String(`for`, `Filter events to only those pertaining to the specified resource.`, ox.Section(0)).
				Bool(`no-headers`, `When using the default output format, don't print headers.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Slice(`types`, `Output only events of given types.`, ox.Section(0)).
				Bool(`watch`, `After listing the requested events, watch for more events.`, ox.Spec(`false`), ox.Short("w"), ox.Section(0)),
		),
		ox.Sub( // kubectl diff
			ox.Usage(`diff`, `Diff the live version against a would-be applied version`),
			ox.Banner(`Diff configurations specified by file name or stdin between the current online configuration, and the configuration as it would be if applied.

 The output is always YAML.

 KUBECTL_EXTERNAL_DIFF environment variable can be used to select your own diff command. Users can use external commands with params too, example: KUBECTL_EXTERNAL_DIFF="colordiff -N -u"

 By default, the "diff" command available in your path will be run with the "-u" (unified diff) and "-N" (treat absent files as empty) options.

 Exit status: 0 No differences were found. 1 Differences were found. >1 Kubectl or diff failed with an error.

 Note: KUBECTL_EXTERNAL_DIFF, if used, is expected to follow that convention.`),
			ox.Spec(`-f FILENAME [options]`),
			ox.Example(`
  # Diff resources included in pod.json
  kubectl diff -f pod.json
  
  # Diff file read from stdin
  cat service.yaml | kubectl diff -f -`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Int(`concurrency`, `Number of objects to process in parallel when diffing against the live version. Larger number = faster, but more memory, I/O and CPU over that shorter period of time.`, ox.Default("1"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-client-side-apply"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files contains the configuration to diff`, ox.Short("f"), ox.Section(0)).
				Bool(`force-conflicts`, `If true, server-side apply will force the changes against conflicts.`, ox.Spec(`false`), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`prune`, `Include resources that would be deleted by pruning. Can be used with -l and default shows all resources would be pruned`, ox.Spec(`false`), ox.Section(0)).
				Slice(`prune-allowlist`, `Overwrite the default allowlist with <group/version/kind> for --prune`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`server-side`, `If true, apply runs in the server instead of the client.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, include managed fields in the diff.`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl apply
			ox.Usage(`apply`, `Apply a configuration to a resource by file name or stdin`),
			ox.Banner(`Apply a configuration to a resource by file name or stdin. The resource name must be specified. This resource will be created if it doesn't exist yet. To use 'apply', always create the resource initially with either 'apply' or 'create --save-config'.

 JSON and YAML formats are accepted.

 Alpha Disclaimer: the --prune functionality is not yet complete. Do not use unless you are aware of what the current state is. See https://issues.k8s.io/34274.`),
			ox.Spec(`(-f FILENAME | -k DIRECTORY) [options]`),
			ox.Example(`
  # Apply the configuration in pod.json to a pod
  kubectl apply -f ./pod.json
  
  # Apply resources from a directory containing kustomization.yaml - e.g. dir/kustomization.yaml
  kubectl apply -k dir/
  
  # Apply the JSON passed into stdin to a pod
  cat pod.json | kubectl apply -f -
  
  # Apply the configuration from all files that end with '.json'
  kubectl apply -f '*.json'
  
  # Note: --prune is still in Alpha
  # Apply the configuration in manifest.yaml that matches label app=nginx and delete all other resources that are not in the file and match label app=nginx
  kubectl apply --prune -f manifest.yaml -l app=nginx
  
  # Apply the configuration in manifest.yaml and delete all the other config maps that are not in the file
  kubectl apply --prune -f manifest.yaml --all --prune-allowlist=core/v1/ConfigMap`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl apply <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all resources in the namespace of the specified resource types.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`cascade`, `Must be "background", "orphan", or "foreground". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.`, ox.Default("background"), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-client-side-apply"), ox.Section(0)).
				Slice(`filename`, `The files that contain the configurations to apply.`, ox.Short("f"), ox.Section(0)).
				Bool(`force`, `If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`force-conflicts`, `If true, server-side apply will force the changes against conflicts.`, ox.Spec(`false`), ox.Section(0)).
				Int(`grace-period`, `Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).`, ox.Default("-1"), ox.Section(0)).
				String(`kustomize`, `Process a kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`openapi-patch`, `If true, use openapi to calculate diff when the openapi presents and the resource can be found in the openapi spec. Otherwise, fall back to use baked-in types.`, ox.Spec(`true`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`overwrite`, `Automatically resolve conflicts between the modified and live configuration by using values from the modified configuration`, ox.Spec(`true`), ox.Section(0)).
				Bool(`prune`, `Automatically delete resource objects, that do not appear in the configs and are created by either apply or create --save-config. Should be used with either -l or --all.`, ox.Spec(`false`), ox.Section(0)).
				Slice(`prune-allowlist`, `Overwrite the default allowlist with <group/version/kind> for --prune`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`server-side`, `If true, apply runs in the server instead of the client.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`subresource`, `If specified, apply will operate on the subresource of the requested object.  Only allowed when using --server-side.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object`, ox.Default("0s"), ox.Section(0)).
				String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
				Bool(`wait`, `If true, wait for resources to be gone before returning. This waits for finalizers.`, ox.Spec(`false`), ox.Section(0)),
			ox.Sub( // kubectl apply edit-last-applied
				ox.Usage(`edit-last-applied`, `Edit latest last-applied-configuration annotations of a resource/object`),
				ox.Banner(`Edit the latest last-applied-configuration annotations of resources from the default editor.

 The edit-last-applied command allows you to directly edit any API resource you can retrieve via the command-line tools. It will open the editor defined by your KUBE_EDITOR, or EDITOR environment variables, or fall back to 'vi' for Linux or 'notepad' for Windows. You can edit multiple objects, although changes are applied one at a time. The command accepts file names as well as command-line arguments, although the files you point to must be previously saved versions of resources.

 The default format is YAML. To edit in JSON, specify "-o json".

 The flag --windows-line-endings can be used to force Windows line endings, otherwise the default for your operating system will be used.

 In the event an error occurs while updating, a temporary file will be created on disk that contains your unapplied changes. The most common error when updating a resource is another editor changing the resource on the server. When this occurs, you will have to apply your changes to the newer version of the resource, or update your temporary saved copy to include the latest resource version.`),
				ox.Spec(`(RESOURCE/NAME | -f FILENAME) [options]`),
				ox.Example(`
  # Edit the last-applied-configuration annotations by type/name in YAML
  kubectl apply edit-last-applied deployment/nginx
  
  # Edit the last-applied-configuration annotations by file in JSON
  kubectl apply edit-last-applied -f deploy.yaml -o json`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-client-side-apply"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files to use to edit the resource`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
					String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
					Bool(`windows-line-endings`, `Defaults to the line ending native to your platform.`, ox.Spec(`false`), ox.Section(0)),
			),
			ox.Sub( // kubectl apply set-last-applied
				ox.Usage(`set-last-applied`, `Set the last-applied-configuration annotation on a live object to match the contents of a file`),
				ox.Banner(`Set the latest last-applied-configuration annotations by setting it to match the contents of a file. This results in the last-applied-configuration being updated as though 'kubectl apply -f<file> ' was run, without updating any other parts of the object.`),
				ox.Spec(`-f FILENAME [options]`),
				ox.Example(`
  # Set the last-applied-configuration of a resource to match the contents of a file
  kubectl apply set-last-applied -f deploy.yaml
  
  # Execute set-last-applied against each configuration file in a directory
  kubectl apply set-last-applied -f path/
  
  # Set the last-applied-configuration of a resource to match the contents of a file; will create the annotation if it does not already exist
  kubectl apply set-last-applied -f deploy.yaml --create-annotation=true`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Bool(`create-annotation`, `Will create 'last-applied-configuration' annotations if current objects doesn't have one`, ox.Spec(`false`), ox.Section(0)).
					String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files that contains the last-applied-configuration annotations`, ox.Short("f"), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
			ox.Sub( // kubectl apply view-last-applied
				ox.Usage(`view-last-applied`, `View the latest last-applied-configuration annotations of a resource/object`),
				ox.Banner(`View the latest last-applied-configuration annotations by type/name or file.

 The default output will be printed to stdout in YAML format. You can use the -o option to change the output format.`),
				ox.Spec(`(TYPE [NAME | -l label] | TYPE/NAME | -f FILENAME) [options]`),
				ox.Example(`
  # View the last-applied-configuration annotations by type/name in YAML
  kubectl apply view-last-applied deployment/nginx
  
  # View the last-applied-configuration annotations by file in JSON
  kubectl apply view-last-applied -f deploy.yaml -o json`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`all`, `Select all resources in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
					Slice(`filename`, `Filename, directory, or URL to files that contains the last-applied-configuration annotations`, ox.Short("f"), ox.Section(0)).
					String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
					String(`output`, `Output format. Must be one of (yaml, json)`, ox.Default("yaml"), ox.Short("o"), ox.Section(0)).
					Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
					String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)),
			),
		),
		ox.Sub( // kubectl patch
			ox.Usage(`patch`, `Update fields of a resource`),
			ox.Banner(`Update fields of a resource using strategic merge patch, a JSON merge patch, or a JSON patch.

 JSON and YAML formats are accepted.

 Note: Strategic merge patch is not supported for custom resources.`),
			ox.Spec(`(-f FILENAME | TYPE NAME) [-p PATCH|--patch-file FILE] [options]`),
			ox.Example(`
  # Partially update a node using a strategic merge patch, specifying the patch as JSON
  kubectl patch node k8s-node-1 -p '{"spec":{"unschedulable":true}}'
  
  # Partially update a node using a strategic merge patch, specifying the patch as YAML
  kubectl patch node k8s-node-1 -p $'spec:\n unschedulable: true'
  
  # Partially update a node identified by the type and name specified in "node.json" using strategic merge patch
  kubectl patch -f node.json -p '{"spec":{"unschedulable":true}}'
  
  # Update a container's image; spec.containers[*].name is required because it's a merge key
  kubectl patch pod valid-pod -p '{"spec":{"containers":[{"name":"kubernetes-serve-hostname","image":"new image"}]}}'
  
  # Update a container's image using a JSON patch with positional arrays
  kubectl patch pod valid-pod --type='json' -p='[{"op": "replace", "path": "/spec/containers/0/image", "value":"new image"}]'
  
  # Update a deployment's replicas through the 'scale' subresource using a merge patch
  kubectl patch deployment nginx-deployment --subresource='scale' --type='merge' -p '{"spec":{"replicas":2}}'`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-patch"), ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to update`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`local`, `If true, patch will operate on the content of the file, not the server-side resource.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				String(`patch`, `The patch to be applied to the resource JSON file.`, ox.Short("p"), ox.Section(0)).
				String(`patch-file`, `A file containing a patch to be applied to the resource.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`subresource`, `If specified, patch will operate on the subresource of the requested object.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				String(`type`, `The type of patch being provided; one of [json merge strategic]`, ox.Default("strategic"), ox.Section(0)),
		),
		ox.Sub( // kubectl replace
			ox.Usage(`replace`, `Replace a resource by file name or stdin`),
			ox.Banner(`Replace a resource by file name or stdin.

 JSON and YAML formats are accepted. If replacing an existing resource, the complete resource spec must be provided. This can be obtained by

        $ kubectl get TYPE NAME -o yaml`),
			ox.Spec(`-f FILENAME [options]`),
			ox.Example(`
  # Replace a pod using the data in pod.json
  kubectl replace -f ./pod.json
  
  # Replace a pod based on the JSON passed into stdin
  cat pod.json | kubectl replace -f -
  
  # Update a single-container pod's image version (tag) to v4
  kubectl get pod mypod -o yaml | sed 's/\(image: myimage\):.*$/\1:v4/' | kubectl replace -f -
  
  # Force replace, delete and then re-create the resource
  kubectl replace --force -f ./pod.json`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`cascade`, `Must be "background", "orphan", or "foreground". Selects the deletion cascading strategy for the dependents (e.g. Pods created by a ReplicationController). Defaults to background.`, ox.Default("background"), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-replace"), ox.Section(0)).
				Slice(`filename`, `The files that contain the configurations to replace.`, ox.Short("f"), ox.Section(0)).
				Bool(`force`, `If true, immediately remove resources from API and bypass graceful deletion. Note that immediate deletion of some resources may result in inconsistency or data loss and requires confirmation.`, ox.Spec(`false`), ox.Section(0)).
				Int(`grace-period`, `Period of time in seconds given to the resource to terminate gracefully. Ignored if negative. Set to 1 for immediate shutdown. Can only be set to 0 when --force is true (force deletion).`, ox.Default("-1"), ox.Section(0)).
				String(`kustomize`, `Process a kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				String(`raw`, `Raw URI to PUT to the server.  Uses the transport specified by the kubeconfig file.`, ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				Bool(`save-config`, `If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`subresource`, `If specified, replace will operate on the subresource of the requested object.`, ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up on a delete, zero means determine a timeout from the size of the object`, ox.Default("0s"), ox.Section(0)).
				String(`validate`, `Must be one of: strict (or true), warn, ignore (or false). "true" or "strict" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. "warn" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as "ignore" otherwise. "false" or "ignore" will not perform any schema validation, silently dropping any unknown or duplicate fields.`, ox.Default("strict"), ox.Section(0)).
				Bool(`wait`, `If true, wait for resources to be gone before returning. This waits for finalizers.`, ox.Spec(`false`), ox.Section(0)),
		),
		ox.Sub( // kubectl wait
			ox.Usage(`wait`, `Experimental: Wait for a specific condition on one or many resources`),
			ox.Banner(`Wait for a specific condition on one or many resources.

 The command takes multiple resources and waits until the specified condition is seen in the Status field of every given resource.

 Alternatively, the command can wait for the given set of resources to be created or deleted by providing the "create" or "delete" keyword as the value to the --for flag.

 A successful message will be printed to stdout indicating when the specified condition has been met. You can use -o option to change to output`),
			ox.Spec(`([-f FILENAME] | resource.group/resource.name | resource.group [(-l label | --all)]) [--for=create|--for=delete|--for condition=available|--for=jsonpath='{}'[=value]] [options]`),
			ox.Example(`
  # Wait for the pod "busybox1" to contain the status condition of type "Ready"
  kubectl wait --for=condition=Ready pod/busybox1
  
  # The default value of status condition is true; you can wait for other targets after an equal delimiter (compared after Unicode simple case folding, which is a more general form of case-insensitivity)
  kubectl wait --for=condition=Ready=false pod/busybox1
  
  # Wait for the pod "busybox1" to contain the status phase to be "Running"
  kubectl wait --for=jsonpath='{.status.phase}'=Running pod/busybox1
  
  # Wait for pod "busybox1" to be Ready
  kubectl wait --for='jsonpath={.status.conditions[?(@.type=="Ready")].status}=True' pod/busybox1
  
  # Wait for the service "loadbalancer" to have ingress
  kubectl wait --for=jsonpath='{.status.loadBalancer.ingress}' service/loadbalancer
  
  # Wait for the secret "busybox1" to be created, with a timeout of 30s
  kubectl create secret generic busybox1
  kubectl wait --for=create secret/busybox1 --timeout=30s
  
  # Wait for the pod "busybox1" to be deleted, with a timeout of 60s, after having issued the "delete" command
  kubectl delete pod/busybox1
  kubectl wait --for=delete pod/busybox1 --timeout=60s`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all resources in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
				Bool(`all-namespaces`, `If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
				Slice(`filename`, `identifying the resource.`, ox.Short("f"), ox.Section(0)).
				String(`for`, `The condition to wait on: [create|delete|condition=condition-name[=condition-value]|jsonpath='{JSONPath expression}'=[JSONPath value]]. The default condition-value is true. Condition values are compared after Unicode simple case folding, which is a more general form of case-insensitivity.`, ox.Section(0)).
				Bool(`local`, `If true, annotation will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`true`), ox.Short("R"), ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)).
				Duration(`timeout`, `The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.`, ox.Default("30s"), ox.Section(0)),
		),
		ox.Sub( // kubectl kustomize
			ox.Usage(`kustomize`, `Build a kustomization target from a directory or URL`),
			ox.Banner(`Build a set of KRM resources using a 'kustomization.yaml' file. The DIR argument must be a path to a directory containing 'kustomization.yaml', or a git repository URL with a path suffix specifying same with respect to the repository root. If DIR is omitted, '.' is assumed.`),
			ox.Spec(`DIR [flags] [options]`),
			ox.Example(`
  # Build the current working directory
  kubectl kustomize
  
  # Build some shared configuration directory
  kubectl kustomize /home/config/production
  
  # Build from github
  kubectl kustomize https://github.com/kubernetes-sigs/kustomize.git/examples/helloWorld?ref=v1.0.6`),
			ox.Section(5),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`as-current-user`, `use the uid and gid of the command executor to run the function in the container`, ox.Spec(`false`), ox.Section(0)).
				Bool(`enable-alpha-plugins`, `enable kustomize plugins`, ox.Spec(`false`), ox.Section(0)).
				Bool(`enable-helm`, `Enable use of the Helm chart inflator generator.`, ox.Spec(`false`), ox.Section(0)).
				Slice(`env`, `a list of environment variables to be used by functions`, ox.Short("e"), ox.Section(0)).
				Slice(`helm-api-versions`, `Kubernetes api versions used by Helm for Capabilities.APIVersions`, ox.Section(0)).
				String(`helm-command`, `helm command (path to executable)`, ox.Default("helm"), ox.Section(0)).
				Bool(`helm-debug`, `Enable debug output from the Helm chart inflator generator.`, ox.Spec(`false`), ox.Section(0)).
				String(`helm-kube-version`, `Kubernetes version used by Helm for Capabilities.KubeVersion`, ox.Section(0)).
				String(`load-restrictor`, `if set to 'LoadRestrictionsNone', local kustomizations may load files from outside their root. This does, however, break the relocatability of the kustomization.`, ox.Default("LoadRestrictionsRootOnly"), ox.Section(0)).
				Slice(`mount`, `a list of storage options read from the filesystem`, ox.Section(0)).
				Bool(`network`, `enable network access for functions that declare it`, ox.Spec(`false`), ox.Section(0)).
				String(`network-name`, `the docker network to run the container in`, ox.Default("bridge"), ox.Section(0)).
				String(`output`, `If specified, write output to this path.`, ox.Short("o"), ox.Section(0)),
		),
		ox.Sub( // kubectl label
			ox.Usage(`label`, `Update the labels on a resource`),
			ox.Banner(`Update the labels on a resource.

  *  A label key and value must begin with a letter or number, and may contain letters, numbers, hyphens, dots, and underscores, up to 63 characters each.
  *  Optionally, the key can begin with a DNS subdomain prefix and a single '/', like example.com/my-app.
  *  If --overwrite is true, then existing labels can be overwritten, otherwise attempting to overwrite a label will result in an error.
  *  If --resource-version is specified, then updates will use this resource version, otherwise the existing resource-version will be used.`),
			ox.Spec(`[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]`),
			ox.Example(`
  # Update pod 'foo' with the label 'unhealthy' and the value 'true'
  kubectl label pods foo unhealthy=true
  
  # Update pod 'foo' with the label 'status' and the value 'unhealthy', overwriting any existing value
  kubectl label --overwrite pods foo status=unhealthy
  
  # Update all pods in the namespace
  kubectl label pods --all status=unhealthy
  
  # Update a pod identified by the type and name in "pod.json"
  kubectl label -f pod.json status=unhealthy
  
  # Update pod 'foo' only if the resource is unchanged from version 1
  kubectl label pods foo status=unhealthy --resource-version=1
  
  # Update pod 'foo' by removing a label named 'bar' if it exists
  # Does not require the --overwrite flag
  kubectl label pods foo bar-`),
			ox.Section(6),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all resources, in the namespace of the specified resource types`, ox.Spec(`false`), ox.Section(0)).
				Bool(`all-namespaces`, `If true, check the specified action in all namespaces.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-label"), ox.Section(0)).
				String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to update the labels`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`list`, `If true, display the labels for a given resource.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`local`, `If true, label will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`overwrite`, `If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labels.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`resource-version`, `If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.`, ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
		),
		ox.Sub( // kubectl annotate
			ox.Usage(`annotate`, `Update the annotations on a resource`),
			ox.Banner(`Update the annotations on one or more resources.

 All Kubernetes objects support the ability to store additional data with the object as annotations. Annotations are key/value pairs that can be larger than labels and include arbitrary string values such as structured JSON. Tools and system extensions may use annotations to store their own data.

 Attempting to set an annotation that already exists will fail unless --overwrite is set. If --resource-version is specified and does not match the current resource version on the server the command will fail.

Use "kubectl api-resources" for a complete list of supported resources.`),
			ox.Spec(`[--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]`),
			ox.Example(`
  # Update pod 'foo' with the annotation 'description' and the value 'my frontend'
  # If the same annotation is set multiple times, only the last value will be applied
  kubectl annotate pods foo description='my frontend'
  
  # Update a pod identified by type and name in "pod.json"
  kubectl annotate -f pod.json description='my frontend'
  
  # Update pod 'foo' with the annotation 'description' and the value 'my frontend running nginx', overwriting any existing value
  kubectl annotate --overwrite pods foo description='my frontend running nginx'
  
  # Update all pods in the namespace
  kubectl annotate pods --all description='my frontend running nginx'
  
  # Update pod 'foo' only if the resource is unchanged from version 1
  kubectl annotate pods foo description='my frontend running nginx' --resource-version=1
  
  # Update pod 'foo' by removing an annotation named 'description' if it exists
  # Does not require the --overwrite flag
  kubectl annotate pods foo description-`),
			ox.Section(6),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				Bool(`all`, `Select all resources, in the namespace of the specified resource types.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`all-namespaces`, `If true, check the specified action in all namespaces.`, ox.Spec(`false`), ox.Short("A"), ox.Section(0)).
				Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
				String(`dry-run`, `Must be "none", "server", or "client". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.`, ox.Default("none"), ox.Section(0)).
				String(`field-manager`, `Name of the manager used to track field ownership.`, ox.Default("$APPNAME-annotate"), ox.Section(0)).
				String(`field-selector`, `Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.`, ox.Section(0)).
				Slice(`filename`, `Filename, directory, or URL to files identifying the resource to update the annotation`, ox.Short("f"), ox.Section(0)).
				String(`kustomize`, `Process the kustomization directory. This flag can't be used together with -f or -R.`, ox.Short("k"), ox.Section(0)).
				Bool(`list`, `If true, display the annotations for a given resource.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`local`, `If true, annotation will NOT contact api-server but run locally.`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Short("o"), ox.Section(0)).
				Bool(`overwrite`, `If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite existing annotations.`, ox.Spec(`false`), ox.Section(0)).
				Bool(`recursive`, `Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.`, ox.Spec(`false`), ox.Short("R"), ox.Section(0)).
				String(`resource-version`, `If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.`, ox.Section(0)).
				String(`selector`, `Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.`, ox.Short("l"), ox.Section(0)).
				Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
				String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
		),
		ox.Sub( // kubectl api-resources
			ox.Usage(`api-resources`, `Print the supported API resources on the server`),
			ox.Banner(`Print the supported API resources on the server.`),
			ox.Spec(`[flags] [options]`),
			ox.Example(`
  # Print the supported API resources
  kubectl api-resources
  
  # Print the supported API resources with more information
  kubectl api-resources -o wide
  
  # Print the supported API resources sorted by a column
  kubectl api-resources --sort-by=name
  
  # Print the supported namespaced resources
  kubectl api-resources --namespaced=true
  
  # Print the supported non-namespaced resources
  kubectl api-resources --namespaced=false
  
  # Print the supported API resources with a specific APIGroup
  kubectl api-resources --api-group=rbac.authorization.k8s.io`),
			ox.Section(7),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Flags().
				String(`api-group`, `Limit to resources in the specified API group.`, ox.Section(0)).
				Bool(`cached`, `Use the cached list of resources if available.`, ox.Spec(`false`), ox.Section(0)).
				Slice(`categories`, `Limit to resources that belong to the specified categories.`, ox.Section(0)).
				Bool(`namespaced`, `If false, non-namespaced resources will be returned, otherwise returning namespaced resources by default.`, ox.Spec(`true`), ox.Section(0)).
				Bool(`no-headers`, `When using the default or custom-column output format, don't print headers (default print headers).`, ox.Spec(`false`), ox.Section(0)).
				String(`output`, `Output format. One of: (wide, name).`, ox.Short("o"), ox.Section(0)).
				String(`sort-by`, `If non-empty, sort list of resources using specified field. The field can be either 'name' or 'kind'.`, ox.Section(0)).
				Slice(`verbs`, `Limit to resources that support the specified verbs.`, ox.Section(0)),
		),
		ox.Sub( // kubectl api-versions
			ox.Usage(`api-versions`, `Print the supported API versions on the server, in the form of "group/version"`),
			ox.Banner(`Print the supported API versions on the server, in the form of "group/version".`),
			ox.Spec(`[options]`),
			ox.Example(`
  # Print the supported API versions
  kubectl api-versions`),
			ox.Section(7),
			ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
		),
		ox.Sub( // kubectl config
			ox.Usage(`config`, `Modify kubeconfig files`),
			ox.Banner(`Modify kubeconfig files using subcommands like "kubectl config set current-context my-context".

 The loading order follows these rules:

  1.  If the --kubeconfig flag is set, then only that file is loaded. The flag may only be set once and no merging takes place.
  2.  If $KUBECONFIG environment variable is set, then it is used as a list of paths (normal path delimiting rules for your system). These paths are merged. When a value is modified, it is modified in the file that defines the stanza. When a value is created, it is created in the first file that exists. If no files in the chain exist, then it creates the last file in the list.
  3.  Otherwise, ${HOME}/.kube/config is used and no merging takes place.`),
			ox.Spec(`SUBCOMMAND [options]`),
			ox.Section(7),
			ox.Footer(`Use "kubectl config <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl config current-context
				ox.Usage(`current-context`, `Display the current-context`),
				ox.Banner(`Display the current-context.`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # Display the current-context
  kubectl config current-context`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config delete-cluster
				ox.Usage(`delete-cluster`, `Delete the specified cluster from the kubeconfig`),
				ox.Banner(`Delete the specified cluster from the kubeconfig.`),
				ox.Spec(`NAME [options]`),
				ox.Example(`
  # Delete the minikube cluster
  kubectl config delete-cluster minikube`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config delete-context
				ox.Usage(`delete-context`, `Delete the specified context from the kubeconfig`),
				ox.Banner(`Delete the specified context from the kubeconfig.`),
				ox.Spec(`NAME [options]`),
				ox.Example(`
  # Delete the context for the minikube cluster
  kubectl config delete-context minikube`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config delete-user
				ox.Usage(`delete-user`, `Delete the specified user from the kubeconfig`),
				ox.Banner(`Delete the specified user from the kubeconfig.`),
				ox.Spec(`NAME [options]`),
				ox.Example(`
  # Delete the minikube user
  kubectl config delete-user minikube`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config get-clusters
				ox.Usage(`get-clusters`, `Display clusters defined in the kubeconfig`),
				ox.Banner(`Display clusters defined in the kubeconfig.`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # List the clusters that kubectl knows about
  kubectl config get-clusters`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config get-contexts
				ox.Usage(`get-contexts`, `Describe one or many contexts`),
				ox.Banner(`Display one or many contexts from the kubeconfig file.`),
				ox.Spec(`[(-o|--output=)name)] [options]`),
				ox.Example(`
  # List all the contexts in your kubeconfig file
  kubectl config get-contexts
  
  # Describe one context in your kubeconfig file
  kubectl config get-contexts my-context`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`no-headers`, `When using the default or custom-column output format, don't print headers (default print headers).`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (name).`, ox.Short("o"), ox.Section(0)),
			),
			ox.Sub( // kubectl config get-users
				ox.Usage(`get-users`, `Display users defined in the kubeconfig`),
				ox.Banner(`Display users defined in the kubeconfig.`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # List the users that kubectl knows about
  kubectl config get-users`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config rename-context
				ox.Usage(`rename-context`, `Rename a context from the kubeconfig file`),
				ox.Banner(`Renames a context from the kubeconfig file.

 CONTEXT_NAME is the context name that you want to change.

 NEW_NAME is the new name you want to set.

 Note: If the context being renamed is the 'current-context', this field will also be updated.`),
				ox.Spec(`CONTEXT_NAME NEW_NAME [options]`),
				ox.Example(`
  # Rename the context 'old-name' to 'new-name' in your kubeconfig file
  kubectl config rename-context old-name new-name`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config set
				ox.Usage(`set`, `Set an individual value in a kubeconfig file`),
				ox.Banner(`Set an individual value in a kubeconfig file.

 PROPERTY_NAME is a dot delimited name where each token represents either an attribute name or a map key.  Map keys may not contain dots.

 PROPERTY_VALUE is the new value you want to set. Binary fields such as 'certificate-authority-data' expect a base64 encoded string unless the --set-raw-bytes flag is used.

 Specifying an attribute name that already exists will merge new fields on top of existing values.`),
				ox.Spec(`PROPERTY_NAME PROPERTY_VALUE [options]`),
				ox.Example(`
  # Set the server field on the my-cluster cluster to https://1.2.3.4
  kubectl config set clusters.my-cluster.server https://1.2.3.4
  
  # Set the certificate-authority-data field on the my-cluster cluster
  kubectl config set clusters.my-cluster.certificate-authority-data $(echo "cert_data_here" | base64 -i -)
  
  # Set the cluster field in the my-context context to my-cluster
  kubectl config set contexts.my-context.cluster my-cluster
  
  # Set the client-key-data field in the cluster-admin user using --set-raw-bytes option
  kubectl config set users.cluster-admin.client-key-data cert_data_here --set-raw-bytes=true`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`set-raw-bytes`, `When writing a []byte PROPERTY_VALUE, write the given string directly without base64 decoding.`, ox.Spec(`false`), ox.Section(0)),
			),
			ox.Sub( // kubectl config set-cluster
				ox.Usage(`set-cluster`, `Set a cluster entry in kubeconfig`),
				ox.Banner(`Set a cluster entry in kubeconfig.

 Specifying a name that already exists will merge new fields on top of existing values for those fields.`),
				ox.Spec(`NAME [--server=server] [--certificate-authority=path/to/certificate/authority] [--insecure-skip-tls-verify=true] [--tls-server-name=example.com] [options]`),
				ox.Example(`
  # Set only the server field on the e2e cluster entry without touching other values
  kubectl config set-cluster e2e --server=https://1.2.3.4
  
  # Embed certificate authority data for the e2e cluster entry
  kubectl config set-cluster e2e --embed-certs --certificate-authority=~/.kube/e2e/kubernetes.ca.crt
  
  # Disable cert checking for the e2e cluster entry
  kubectl config set-cluster e2e --insecure-skip-tls-verify=true
  
  # Set the custom TLS server name to use for validation for the e2e cluster entry
  kubectl config set-cluster e2e --tls-server-name=my-cluster-name
  
  # Set the proxy URL for the e2e cluster entry
  kubectl config set-cluster e2e --proxy-url=https://1.2.3.4`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					String(`certificate-authority`, `Path to certificate-authority file for the cluster entry in kubeconfig`, ox.Section(0)).
					Bool(`embed-certs`, `embed-certs for the cluster entry in kubeconfig`, ox.Spec(`false`), ox.Section(0)).
					Bool(`insecure-skip-tls-verify`, `insecure-skip-tls-verify for the cluster entry in kubeconfig`, ox.Spec(`false`), ox.Section(0)).
					String(`proxy-url`, `proxy-url for the cluster entry in kubeconfig`, ox.Section(0)).
					String(`server`, `server for the cluster entry in kubeconfig`, ox.Section(0)).
					String(`tls-server-name`, `tls-server-name for the cluster entry in kubeconfig`, ox.Section(0)),
			),
			ox.Sub( // kubectl config set-context
				ox.Usage(`set-context`, `Set a context entry in kubeconfig`),
				ox.Banner(`Set a context entry in kubeconfig.

 Specifying a name that already exists will merge new fields on top of existing values for those fields.`),
				ox.Spec(`[NAME | --current] [--cluster=cluster_nickname] [--user=user_nickname] [--namespace=namespace] [options]`),
				ox.Example(`
  # Set the user field on the gce context entry without touching other values
  kubectl config set-context gce --user=cluster-admin`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					String(`cluster`, `cluster for the context entry in kubeconfig`, ox.Section(0)).
					Bool(`current`, `Modify the current context`, ox.Spec(`false`), ox.Section(0)).
					String(`namespace`, `namespace for the context entry in kubeconfig`, ox.Section(0)).
					String(`user`, `user for the context entry in kubeconfig`, ox.Section(0)),
			),
			ox.Sub( // kubectl config set-credentials
				ox.Usage(`set-credentials`, `Set a user entry in kubeconfig`),
				ox.Banner(`Set a user entry in kubeconfig.

 Specifying a name that already exists will merge new fields on top of existing values.

        Client-certificate flags:
        --client-certificate=certfile --client-key=keyfile
        
        Bearer token flags:
        --token=bearer_token
        
        Basic auth flags:
        --username=basic_user --password=basic_password
        
 Bearer token and basic auth are mutually exclusive.`),
				ox.Spec(`NAME [--client-certificate=path/to/certfile] [--client-key=path/to/keyfile] [--token=bearer_token] [--username=basic_user] [--password=basic_password] [--auth-provider=provider_name] [--auth-provider-arg=key=value] [--exec-command=exec_command] [--exec-api-version=exec_api_version] [--exec-arg=arg] [--exec-env=key=value] [options]`),
				ox.Example(`
  # Set only the "client-key" field on the "cluster-admin"
  # entry, without touching other values
  kubectl config set-credentials cluster-admin --client-key=~/.kube/admin.key
  
  # Set basic auth for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --username=admin --password=uXFGweU9l35qcif
  
  # Embed client certificate data in the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --client-certificate=~/.kube/admin.crt --embed-certs=true
  
  # Enable the Google Compute Platform auth provider for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --auth-provider=gcp
  
  # Enable the OpenID Connect auth provider for the "cluster-admin" entry with additional arguments
  kubectl config set-credentials cluster-admin --auth-provider=oidc --auth-provider-arg=client-id=foo --auth-provider-arg=client-secret=bar
  
  # Remove the "client-secret" config value for the OpenID Connect auth provider for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --auth-provider=oidc --auth-provider-arg=client-secret-
  
  # Enable new exec auth plugin for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --exec-command=/path/to/the/executable --exec-api-version=client.authentication.k8s.io/v1beta1
  
  # Enable new exec auth plugin for the "cluster-admin" entry with interactive mode
  kubectl config set-credentials cluster-admin --exec-command=/path/to/the/executable --exec-api-version=client.authentication.k8s.io/v1beta1 --exec-interactive-mode=Never
  
  # Define new exec auth plugin arguments for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --exec-arg=arg1 --exec-arg=arg2
  
  # Create or update exec auth plugin environment variables for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --exec-env=key1=val1 --exec-env=key2=val2
  
  # Remove exec auth plugin environment variables for the "cluster-admin" entry
  kubectl config set-credentials cluster-admin --exec-env=var-to-remove-`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					String(`auth-provider`, `Auth provider for the user entry in kubeconfig`, ox.Section(0)).
					Slice(`auth-provider-arg`, `'key=value' arguments for the auth provider`, ox.Section(0)).
					String(`client-certificate`, `Path to client-certificate file for the user entry in kubeconfig`, ox.Section(0)).
					String(`client-key`, `Path to client-key file for the user entry in kubeconfig`, ox.Section(0)).
					Bool(`embed-certs`, `Embed client cert/key for the user entry in kubeconfig`, ox.Spec(`false`), ox.Section(0)).
					String(`exec-api-version`, `API version of the exec credential plugin for the user entry in kubeconfig`, ox.Section(0)).
					Slice(`exec-arg`, `New arguments for the exec credential plugin command for the user entry in kubeconfig`, ox.Section(0)).
					String(`exec-command`, `Command for the exec credential plugin for the user entry in kubeconfig`, ox.Section(0)).
					Slice(`exec-env`, `'key=value' environment values for the exec credential plugin`, ox.Section(0)).
					String(`exec-interactive-mode`, `InteractiveMode of the exec credentials plugin for the user entry in kubeconfig`, ox.Section(0)).
					Bool(`exec-provide-cluster-info`, `ProvideClusterInfo of the exec credentials plugin for the user entry in kubeconfig`, ox.Spec(`false`), ox.Section(0)).
					String(`password`, `password for the user entry in kubeconfig`, ox.Section(0)).
					String(`token`, `token for the user entry in kubeconfig`, ox.Section(0)).
					String(`username`, `username for the user entry in kubeconfig`, ox.Section(0)),
			),
			ox.Sub( // kubectl config unset
				ox.Usage(`unset`, `Unset an individual value in a kubeconfig file`),
				ox.Banner(`Unset an individual value in a kubeconfig file.

 PROPERTY_NAME is a dot delimited name where each token represents either an attribute name or a map key.  Map keys may not contain dots.`),
				ox.Spec(`PROPERTY_NAME [options]`),
				ox.Example(`
  # Unset the current-context
  kubectl config unset current-context
  
  # Unset namespace in foo context
  kubectl config unset contexts.foo.namespace`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config use-context
				ox.Usage(`use-context`, `Set the current-context in a kubeconfig file`),
				ox.Banner(`Set the current-context in a kubeconfig file.`),
				ox.Spec(`CONTEXT_NAME [options]`),
				ox.Aliases("use"),
				ox.Example(`
  # Use the context for the minikube cluster
  kubectl config use-context minikube`),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			),
			ox.Sub( // kubectl config view
				ox.Usage(`view`, `Display merged kubeconfig settings or a specified kubeconfig file`),
				ox.Banner(`Display merged kubeconfig settings or a specified kubeconfig file.

 You can use --output jsonpath={...} to extract specific values using a jsonpath expression.`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # Show merged kubeconfig settings
  kubectl config view
  
  # Show merged kubeconfig settings, raw certificate data, and exposed secrets
  kubectl config view --raw
  
  # Get the password for the e2e user
  kubectl config view -o jsonpath='{.users[?(@.name == "e2e")].user.password}'`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`allow-missing-template-keys`, `If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.`, ox.Spec(`true`), ox.Section(0)).
					Bool(`flatten`, `Flatten the resulting kubeconfig file into self-contained output (useful for creating portable kubeconfig files)`, ox.Spec(`false`), ox.Section(0)).
					Bool(`merge`, `Merge the full hierarchy of kubeconfig files`, ox.Spec(`true`), ox.Section(0)).
					Bool(`minify`, `Remove all information not used by current-context from the output`, ox.Spec(`false`), ox.Section(0)).
					String(`output`, `Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).`, ox.Default("yaml"), ox.Short("o"), ox.Section(0)).
					Bool(`raw`, `Display raw byte data and sensitive data`, ox.Spec(`false`), ox.Section(0)).
					Bool(`show-managed-fields`, `If true, keep the managedFields when printing objects in JSON or YAML format.`, ox.Spec(`false`), ox.Section(0)).
					String(`template`, `Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].`, ox.Section(0)),
			),
		),
		ox.Sub( // kubectl plugin
			ox.Usage(`plugin`, `Provides utilities for interacting with plugins`),
			ox.Banner(`Provides utilities for interacting with plugins.

 Plugins provide extended functionality that is not part of the major command-line distribution. Please refer to the documentation and examples for more information about how write your own plugins.

 The easiest way to discover and install plugins is via the kubernetes sub-project krew: [krew.sigs.k8s.io]. To install krew, visit https://krew.sigs.k8s.io/docs/user-guide/setup/install`),
			ox.Spec(`[flags] [options]`),
			ox.Example(`
  # List all available plugins
  kubectl plugin list
  
  # List only binary names of available plugins without paths
  kubectl plugin list --name-only`),
			ox.Section(7),
			ox.Footer(`Use "kubectl plugin <command> --help" for more information about a given command.
Use "kubectl options" for a list of global command-line options (applies to all commands).`),
			ox.Sub( // kubectl plugin list
				ox.Usage(`list`, `List all visible plugin executables on a user's PATH`),
				ox.Banner(`List all available plugin files on a user's PATH. To see plugins binary names without the full path use --name-only flag.

 Available plugin files are those that are: - executable - anywhere on the user's PATH - begin with "kubectl-"`),
				ox.Spec(`[flags] [options]`),
				ox.Example(`
  # List all available plugins
  kubectl plugin list
  
  # List only binary names of available plugins without paths
  kubectl plugin list --name-only`),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Footer(`Use "kubectl options" for a list of global command-line options (applies to all commands).`),
				ox.Flags().
					Bool(`name-only`, `If true, display only the binary name of each plugin, rather than its full path`, ox.Spec(`false`), ox.Section(0)),
			),
		),
	)
}
