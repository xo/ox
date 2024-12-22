// Command docker is a xo/ox version of `+docker`.
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
		ox.Usage("docker", ""),
		ox.Spec("[OPTIONS] COMMAND"),
		ox.Sections("Common Commands", "Management Commands", "Swarm Commands", "Commands"),
		ox.Footer("Run 'docker COMMAND --help' for more information on a command.\n\nFor more help on how to use Docker, head to https://docs.docker.com/go/guides/"),
		ox.Sub(
			ox.Usage("run", "Create and run a new container from an image"),
			ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
			ox.Aliases("docker container run", "docker run"),
			ox.Section(0),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping", ox.Elem(ox.StringT)).
				Map("annotation", "Add an annotation to the", ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
				Uint16("blkio-weight", "Block IO (relative weight),").
				Slice("blkio-weight-device", "Block IO weight (relative device", ox.Elem(ox.StringT)).
				Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
				Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
				String("cgroup-parent", "Optional parent cgroup for the").
				String("cgroupns", "Cgroup namespace to use").
				String("cidfile", "Write the container ID to the file").
				Int("cpu-count", "CPU count (Windows only)").
				Int("cpu-percent", "CPU percent (Windows only)").
				Int("cpu-period", "Limit CPU CFS (Completely Fair").
				Int("cpu-quota", "Limit CPU CFS (Completely Fair").
				Int("cpu-rt-period", "Limit CPU real-time period in").
				Int("cpu-rt-runtime", "Limit CPU real-time runtime in").
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
				String("cpuset-cpus", "CPUs in which to allow execution").
				String("cpuset-mems", "MEMs in which to allow execution").
				Bool("detach", "Run container in background and", ox.Short("d")).
				String("detach-keys", "Override the key sequence for").
				Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Elem(ox.StringT)).
				Slice("device-read-bps", "Limit read rate (bytes per", ox.Elem(ox.StringT)).
				Slice("device-read-iops", "Limit read rate (IO per second)", ox.Elem(ox.StringT)).
				Slice("device-write-bps", "Limit write rate (bytes per", ox.Elem(ox.StringT)).
				Slice("device-write-iops", "Limit write rate (IO per second)", ox.Elem(ox.StringT)).
				Bool("disable-content-trust", "Skip image verification (default").
				Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
				Slice("dns-option", "Set DNS options", ox.Elem(ox.StringT)).
				Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
				String("domainname", "Container NIS domain name").
				String("entrypoint", "Overwrite the default ENTRYPOINT").
				Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
				Slice("env-file", "Read in a file of environment", ox.Elem(ox.StringT)).
				Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
				String("gpus", "GPU devices to add to the", ox.Spec("gpu-request")).
				Slice("group-add", "Add additional groups to join", ox.Elem(ox.StringT)).
				String("health-cmd", "Command to run to check health").
				Duration("health-interval", "Time between running the check").
				Int("health-retries", "Consecutive failures needed to").
				Duration("health-start-interval", "Time between running the check").
				Duration("health-start-period", "Start period for the container").
				Duration("health-timeout", "Maximum time to allow one check").
				String("hostname", "Container host name", ox.Short("h")).
				Bool("init", "Run an init inside the container").
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
				Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT)).
				Uint("io-maxiops", "Maximum IOps limit for the").
				String("ip", "IPv4 address (e.g., 172.30.100.104)").
				String("ip6", "IPv6 address (e.g., 2001:db8::33)").
				String("ipc", "IPC mode to use").
				String("isolation", "Container isolation technology").
				Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT)).
				Slice("label", "Set meta data on a container", ox.Elem(ox.StringT), ox.Short("l")).
				Slice("label-file", "Read in a line delimited file of", ox.Elem(ox.StringT)).
				Slice("link", "Add link to another container", ox.Elem(ox.StringT)).
				Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Elem(ox.StringT)).
				String("log-driver", "Logging driver for the container").
				Slice("log-opt", "Log driver options", ox.Elem(ox.StringT)).
				String("mac-address", "Container MAC address (e.g.,").
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
				Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT)).
				Int("memory-swappiness", "Tune container memory swappiness").
				String("mount", "Attach a filesystem mount to the", ox.Spec("mount")).
				String("name", "Assign a name to the container").
				String("network", "Connect a container to a network", ox.Spec("network")).
				Slice("network-alias", "Add network-scoped alias for the", ox.Elem(ox.StringT)).
				Bool("no-healthcheck", "Disable any container-specified").
				Bool("oom-kill-disable", "Disable OOM Killer").
				Int("oom-score-adj", "Tune host's OOM preferences").
				String("pid", "PID namespace to use").
				Int("pids-limit", "Tune container pids limit (set").
				String("platform", "Set platform if server is").
				Bool("privileged", "Give extended privileges to this").
				Slice("publish", "Publish a container's port(s) to", ox.Elem(ox.StringT), ox.Short("p")).
				Bool("publish-all", "Publish all exposed ports to", ox.Short("P")).
				String("pull", "Pull image before running").
				Bool("quiet", "Suppress the pull output", ox.Short("q")).
				Bool("read-only", "Mount the container's root").
				String("restart", "Restart policy to apply when a").
				Bool("rm", "Automatically remove the").
				String("runtime", "Runtime to use for this container").
				Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
				Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT)).
				Bool("sig-proxy", "Proxy received signals to the").
				String("stop-signal", "Signal to stop the container").
				Int("stop-timeout", "Timeout (in seconds) to stop a").
				Slice("storage-opt", "Storage driver options for the", ox.Elem(ox.StringT)).
				Map("sysctl", "Sysctl options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
				Slice("tmpfs", "Mount a tmpfs directory", ox.Elem(ox.StringT)).
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")).
				String("user", "Username or UID (format:", ox.Short("u")).
				String("userns", "User namespace to use").
				String("uts", "UTS namespace to use").
				Slice("volume", "Bind mount a volume", ox.Elem(ox.StringT), ox.Short("v")).
				String("volume-driver", "Optional volume driver for the").
				Slice("volumes-from", "Mount volumes from the specified", ox.Elem(ox.StringT)).
				String("workdir", "Working directory inside the", ox.Short("w")),
		), ox.Sub(
			ox.Usage("exec", "Execute a command in a running container"),
			ox.Spec("[OPTIONS] CONTAINER COMMAND [ARG...]"),
			ox.Aliases("docker container exec", "docker exec"),
			ox.Section(0),
			ox.Flags().
				Bool("detach", "Detached mode: run command in the background", ox.Short("d")).
				String("detach-keys", "Override the key sequence for detaching a").
				Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
				Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
				Bool("privileged", "Give extended privileges to the command").
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
				String("user", "Username or UID (format:", ox.Short("u")).
				String("workdir", "Working directory inside the container", ox.Short("w")),
		), ox.Sub(
			ox.Usage("ps", "List containers"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("docker container ls", "docker container list", "docker container ps", "docker ps"),
			ox.Section(0),
			ox.Flags().
				Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a")).
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
				String("format", "Format output using a custom template:").
				Int("last", "Show n last created containers (includes all", ox.Short("n")).
				Bool("latest", "Show the latest created container (includes all", ox.Short("l")).
				Bool("no-trunc", "Don't truncate output").
				Bool("quiet", "Only display container IDs", ox.Short("q")).
				Bool("size", "Display total file sizes", ox.Short("s")),
		), ox.Sub(
			ox.Usage("build", "Build an image from a Dockerfile"),
			ox.Spec("[OPTIONS] PATH | URL | -"),
			ox.Aliases("docker image build", "docker build", "docker builder build"),
			ox.Section(0),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Elem(ox.StringT)).
				Slice("build-arg", "Set build-time variables", ox.Elem(ox.StringT)).
				Slice("cache-from", "Images to consider as cache sources", ox.Elem(ox.StringT)).
				String("cgroup-parent", "Set the parent cgroup for the \"RUN\"").
				Bool("compress", "Compress the build context using gzip").
				Int("cpu-period", "Limit the CPU CFS (Completely Fair").
				Int("cpu-quota", "Limit the CPU CFS (Completely Fair").
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)").
				Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
				String("file", "Name of the Dockerfile (Default is", ox.Short("f")).
				Bool("force-rm", "Always remove intermediate containers").
				String("iidfile", "Write the image ID to the file").
				String("isolation", "Container isolation technology").
				Slice("label", "Set metadata for an image", ox.Elem(ox.StringT)).
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
				Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT)).
				String("network", "Set the networking mode for the RUN").
				Bool("no-cache", "Do not use cache when building the image").
				String("platform", "Set platform if server is multi-platform").
				Bool("pull", "Always attempt to pull a newer version of").
				Bool("quiet", "Suppress the build output and print image", ox.Short("q")).
				Bool("rm", "Remove intermediate containers after a").
				Slice("security-opt", "Security options", ox.Elem(ox.StringT)).
				Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT)).
				Bool("squash", "Squash newly built layers into a single").
				Slice("tag", "Name and optionally a tag in the", ox.Elem(ox.StringT), ox.Short("t")).
				String("target", "Set the target build stage to build.").
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")),
		), ox.Sub(
			ox.Usage("pull", "Download an image from a registry"),
			ox.Spec("[OPTIONS] NAME[:TAG|@DIGEST]"),
			ox.Aliases("docker image pull", "docker pull"),
			ox.Section(0),
			ox.Flags().
				Bool("all-tags", "Download all tagged images in the repository", ox.Short("a")).
				Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
				String("platform", "Set platform if server is multi-platform").
				Bool("quiet", "Suppress verbose output", ox.Short("q")),
		), ox.Sub(
			ox.Usage("push", "Upload an image to a registry"),
			ox.Spec("[OPTIONS] NAME[:TAG]"),
			ox.Aliases("docker image push", "docker push"),
			ox.Section(0),
			ox.Flags().
				Bool("all-tags", "Push all tags of an image to the repository", ox.Short("a")).
				Bool("disable-content-trust", "Skip image signing", ox.Default("true")).
				String("platform", "Push a platform-specific manifest as a").
				Bool("quiet", "Suppress verbose output", ox.Short("q")),
		), ox.Sub(
			ox.Usage("images", "List images"),
			ox.Spec("[OPTIONS] [REPOSITORY[:TAG]]"),
			ox.Aliases("docker image ls", "docker image list", "docker images"),
			ox.Section(0),
			ox.Flags().
				Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a")).
				Bool("digests", "Show digests").
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
				String("format", "Format output using a custom template:").
				Bool("no-trunc", "Don't truncate output").
				Bool("quiet", "Only show image IDs", ox.Short("q")).
				Bool("tree", "List multi-platform images as a tree (EXPERIMENTAL)"),
		), ox.Sub(
			ox.Usage("login", "Authenticate to a registry"),
			ox.Spec("[OPTIONS] [SERVER]"),
			ox.Section(0),
			ox.Flags().
				String("password", "Password", ox.Short("p")).
				Bool("password-stdin", "Take the password from stdin").
				String("username", "Username", ox.Short("u")),
		), ox.Sub(
			ox.Usage("logout", "Log out from a registry"),
			ox.Spec("[SERVER]"),
			ox.Section(0),
			ox.Footer("Log out from a registry.\nIf no server is specified, the default is defined by the daemon."),
		), ox.Sub(
			ox.Usage("search", "Search Docker Hub for images"),
			ox.Spec("[OPTIONS] TERM"),
			ox.Section(0),
			ox.Flags().
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
				String("format", "Pretty-print search using a Go template").
				Int("limit", "Max number of search results").
				Bool("no-trunc", "Don't truncate output"),
		), ox.Sub(
			ox.Usage("info", "Display system-wide information"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("docker system info", "docker info"),
			ox.Section(0),
			ox.Flags().
				String("format", "Format output using a custom template:", ox.Short("f")),
		), ox.Sub(
			ox.Usage("builder", "Manage builds"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker builder COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("build", "Build an image from a Dockerfile"),
				ox.Spec("[OPTIONS] PATH | URL | -"),
				ox.Aliases("docker image build", "docker build", "docker builder build"),
				ox.Section(0),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Elem(ox.StringT)).
					Slice("build-arg", "Set build-time variables", ox.Elem(ox.StringT)).
					Slice("cache-from", "Images to consider as cache sources", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Set the parent cgroup for the \"RUN\"").
					Bool("compress", "Compress the build context using gzip").
					Int("cpu-period", "Limit the CPU CFS (Completely Fair").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair").
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)").
					Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
					String("file", "Name of the Dockerfile (Default is", ox.Short("f")).
					Bool("force-rm", "Always remove intermediate containers").
					String("iidfile", "Write the image ID to the file").
					String("isolation", "Container isolation technology").
					Slice("label", "Set metadata for an image", ox.Elem(ox.StringT)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
					Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT)).
					String("network", "Set the networking mode for the RUN").
					Bool("no-cache", "Do not use cache when building the image").
					String("platform", "Set platform if server is multi-platform").
					Bool("pull", "Always attempt to pull a newer version of").
					Bool("quiet", "Suppress the build output and print image", ox.Short("q")).
					Bool("rm", "Remove intermediate containers after a").
					Slice("security-opt", "Security options", ox.Elem(ox.StringT)).
					Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT)).
					Bool("squash", "Squash newly built layers into a single").
					Slice("tag", "Name and optionally a tag in the", ox.Elem(ox.StringT), ox.Short("t")).
					String("target", "Set the target build stage to build.").
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")),
			), ox.Sub(
				ox.Usage("prune", "Remove build cache"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Remove all unused build cache, not just", ox.Short("a")).
					String("filter", "Provide filter values (e.g. \"until=24h\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")).
					Slice("keep-storage", "Amount of disk space to keep for cache", ox.Elem(ox.UintT)),
			)), ox.Sub(
			ox.Usage("checkpoint", "Manage checkpoints"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker checkpoint COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a checkpoint from a running container"),
				ox.Spec("[OPTIONS] CONTAINER CHECKPOINT"),
				ox.Section(0),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory").
					Bool("leave-running", "Leave the container running after checkpoint"),
			), ox.Sub(
				ox.Usage("ls", "List checkpoints for a container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("docker checkpoint list"),
				ox.Section(0),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory"),
			), ox.Sub(
				ox.Usage("rm", "Remove a checkpoint"),
				ox.Spec("[OPTIONS] CONTAINER CHECKPOINT"),
				ox.Aliases("docker checkpoint remove"),
				ox.Section(0),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory"),
			)), ox.Sub(
			ox.Usage("container", "Manage containers"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker container COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("attach", "Attach local standard input, output, and error streams to a running container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("docker attach"),
				ox.Section(0),
				ox.Flags().
					String("detach-keys", "Override the key sequence for detaching a").
					Bool("no-stdin", "Do not attach STDIN").
					Bool("sig-proxy", "Proxy all received signals to the process"),
			), ox.Sub(
				ox.Usage("commit", "Create a new image from a container's changes"),
				ox.Spec("[OPTIONS] CONTAINER [REPOSITORY[:TAG]]"),
				ox.Aliases("docker commit"),
				ox.Section(0),
				ox.Flags().
					String("author", "Author (e.g., \"John Hannibal Smith", ox.Short("a")).
					Slice("change", "Apply Dockerfile instruction to the created image", ox.Elem(ox.StringT), ox.Short("c")).
					String("message", "Commit message", ox.Short("m")).
					Bool("pause", "Pause container during commit", ox.Default("true"), ox.Short("p")),
			), ox.Sub(
				ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
				ox.Spec("[OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-"),
				ox.Aliases("docker cp"),
				ox.Section(0),
				ox.Flags().
					Bool("archive", "Archive mode (copy all uid/gid information)", ox.Short("a")).
					Bool("follow-link", "Always follow symbol link in SRC_PATH", ox.Short("L")).
					Bool("quiet", "Suppress progress output during copy. Progress", ox.Short("q")),
			), ox.Sub(
				ox.Usage("create", "Create a new container"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Aliases("docker create"),
				ox.Section(0),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping", ox.Elem(ox.StringT)).
					Map("annotation", "Add an annotation to the", ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
					Uint16("blkio-weight", "Block IO (relative weight),").
					Slice("blkio-weight-device", "Block IO weight (relative device", ox.Elem(ox.StringT)).
					Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Optional parent cgroup for the").
					String("cgroupns", "Cgroup namespace to use").
					String("cidfile", "Write the container ID to the file").
					Int("cpu-count", "CPU count (Windows only)").
					Int("cpu-percent", "CPU percent (Windows only)").
					Int("cpu-period", "Limit CPU CFS (Completely Fair").
					Int("cpu-quota", "Limit CPU CFS (Completely Fair").
					Int("cpu-rt-period", "Limit CPU real-time period in").
					Int("cpu-rt-runtime", "Limit CPU real-time runtime in").
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
					String("cpuset-cpus", "CPUs in which to allow execution").
					String("cpuset-mems", "MEMs in which to allow execution").
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per", ox.Elem(ox.StringT)).
					Slice("device-read-iops", "Limit read rate (IO per second)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per", ox.Elem(ox.StringT)).
					Slice("device-write-iops", "Limit write rate (IO per second)", ox.Elem(ox.StringT)).
					Bool("disable-content-trust", "Skip image verification (default").
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("domainname", "Container NIS domain name").
					String("entrypoint", "Overwrite the default ENTRYPOINT").
					Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
					Slice("env-file", "Read in a file of environment", ox.Elem(ox.StringT)).
					Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
					String("gpus", "GPU devices to add to the", ox.Spec("gpu-request")).
					Slice("group-add", "Add additional groups to join", ox.Elem(ox.StringT)).
					String("health-cmd", "Command to run to check health").
					Duration("health-interval", "Time between running the check").
					Int("health-retries", "Consecutive failures needed to").
					Duration("health-start-interval", "Time between running the check").
					Duration("health-start-period", "Start period for the container").
					Duration("health-timeout", "Maximum time to allow one check").
					String("hostname", "Container host name", ox.Short("h")).
					Bool("init", "Run an init inside the container").
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
					Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT)).
					Uint("io-maxiops", "Maximum IOps limit for the").
					String("ip", "IPv4 address (e.g., 172.30.100.104)").
					String("ip6", "IPv6 address (e.g., 2001:db8::33)").
					String("ipc", "IPC mode to use").
					String("isolation", "Container isolation technology").
					Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT)).
					Slice("label", "Set meta data on a container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of", ox.Elem(ox.StringT)).
					Slice("link", "Add link to another container", ox.Elem(ox.StringT)).
					Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Elem(ox.StringT)).
					String("log-driver", "Logging driver for the container").
					Slice("log-opt", "Log driver options", ox.Elem(ox.StringT)).
					String("mac-address", "Container MAC address (e.g.,").
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
					Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT)).
					Int("memory-swappiness", "Tune container memory swappiness").
					String("mount", "Attach a filesystem mount to the", ox.Spec("mount")).
					String("name", "Assign a name to the container").
					String("network", "Connect a container to a network", ox.Spec("network")).
					Slice("network-alias", "Add network-scoped alias for the", ox.Elem(ox.StringT)).
					Bool("no-healthcheck", "Disable any container-specified").
					Bool("oom-kill-disable", "Disable OOM Killer").
					Int("oom-score-adj", "Tune host's OOM preferences").
					String("pid", "PID namespace to use").
					Int("pids-limit", "Tune container pids limit (set").
					String("platform", "Set platform if server is").
					Bool("privileged", "Give extended privileges to this").
					Slice("publish", "Publish a container's port(s) to", ox.Elem(ox.StringT), ox.Short("p")).
					Bool("publish-all", "Publish all exposed ports to", ox.Short("P")).
					String("pull", "Pull image before creating").
					Bool("quiet", "Suppress the pull output", ox.Short("q")).
					Bool("read-only", "Mount the container's root").
					String("restart", "Restart policy to apply when a").
					Bool("rm", "Automatically remove the").
					String("runtime", "Runtime to use for this container").
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT)).
					String("stop-signal", "Signal to stop the container").
					Int("stop-timeout", "Timeout (in seconds) to stop a").
					Slice("storage-opt", "Storage driver options for the", ox.Elem(ox.StringT)).
					Map("sysctl", "Sysctl options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
					Slice("tmpfs", "Mount a tmpfs directory", ox.Elem(ox.StringT)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")).
					String("user", "Username or UID (format:", ox.Short("u")).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					Slice("volume", "Bind mount a volume", ox.Elem(ox.StringT), ox.Short("v")).
					String("volume-driver", "Optional volume driver for the").
					Slice("volumes-from", "Mount volumes from the specified", ox.Elem(ox.StringT)).
					String("workdir", "Working directory inside the", ox.Short("w")),
			), ox.Sub(
				ox.Usage("diff", "Inspect changes to files or directories on a container's filesystem"),
				ox.Spec("CONTAINER"),
				ox.Aliases("docker diff"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("exec", "Execute a command in a running container"),
				ox.Spec("[OPTIONS] CONTAINER COMMAND [ARG...]"),
				ox.Aliases("docker exec"),
				ox.Section(0),
				ox.Flags().
					Bool("detach", "Detached mode: run command in the background", ox.Short("d")).
					String("detach-keys", "Override the key sequence for detaching a").
					Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
					Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
					Bool("privileged", "Give extended privileges to the command").
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
					String("user", "Username or UID (format:", ox.Short("u")).
					String("workdir", "Working directory inside the container", ox.Short("w")),
			), ox.Sub(
				ox.Usage("export", "Export a container's filesystem as a tar archive"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("docker export"),
				ox.Section(0),
				ox.Flags().
					String("output", "Write to a file, instead of STDOUT", ox.Short("o")),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("size", "Display total file sizes", ox.Short("s")),
			), ox.Sub(
				ox.Usage("kill", "Kill one or more running containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker kill"),
				ox.Section(0),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s")),
			), ox.Sub(
				ox.Usage("logs", "Fetch the logs of a container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("docker logs"),
				ox.Section(0),
				ox.Flags().
					Bool("details", "Show extra details provided to logs").
					Bool("follow", "Follow log output", ox.Short("f")).
					String("since", "Show logs since timestamp (e.g.").
					String("tail", "Number of lines to show from the end of the logs", ox.Short("n")).
					Bool("timestamps", "Show timestamps", ox.Short("t")).
					String("until", "Show logs before a timestamp (e.g."),
			), ox.Sub(
				ox.Usage("ls", "List containers"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker container list", "docker container ps", "docker ps"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a")).
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Int("last", "Show n last created containers (includes all", ox.Short("n")).
					Bool("latest", "Show the latest created container (includes all", ox.Short("l")).
					Bool("no-trunc", "Don't truncate output").
					Bool("quiet", "Only display container IDs", ox.Short("q")).
					Bool("size", "Display total file sizes", ox.Short("s")),
			), ox.Sub(
				ox.Usage("pause", "Pause all processes within one or more containers"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("docker pause"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("port", "List port mappings or a specific mapping for the container"),
				ox.Spec("CONTAINER [PRIVATE_PORT[/PROTO]]"),
				ox.Aliases("docker port"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("prune", "Remove all stopped containers"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Usage("rename", "Rename a container"),
				ox.Spec("CONTAINER NEW_NAME"),
				ox.Aliases("docker rename"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("restart", "Restart one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker restart"),
				ox.Section(0),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s")).
					Int("time", "Seconds to wait before killing the container", ox.Short("t")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker container remove", "docker rm"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the removal of a running container (uses SIGKILL)", ox.Short("f")).
					Bool("link", "Remove the specified link", ox.Short("l")).
					Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v")),
			), ox.Sub(
				ox.Usage("run", "Create and run a new container from an image"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Aliases("docker run"),
				ox.Section(0),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping", ox.Elem(ox.StringT)).
					Map("annotation", "Add an annotation to the", ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
					Uint16("blkio-weight", "Block IO (relative weight),").
					Slice("blkio-weight-device", "Block IO weight (relative device", ox.Elem(ox.StringT)).
					Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Optional parent cgroup for the").
					String("cgroupns", "Cgroup namespace to use").
					String("cidfile", "Write the container ID to the file").
					Int("cpu-count", "CPU count (Windows only)").
					Int("cpu-percent", "CPU percent (Windows only)").
					Int("cpu-period", "Limit CPU CFS (Completely Fair").
					Int("cpu-quota", "Limit CPU CFS (Completely Fair").
					Int("cpu-rt-period", "Limit CPU real-time period in").
					Int("cpu-rt-runtime", "Limit CPU real-time runtime in").
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
					String("cpuset-cpus", "CPUs in which to allow execution").
					String("cpuset-mems", "MEMs in which to allow execution").
					Bool("detach", "Run container in background and", ox.Short("d")).
					String("detach-keys", "Override the key sequence for").
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per", ox.Elem(ox.StringT)).
					Slice("device-read-iops", "Limit read rate (IO per second)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per", ox.Elem(ox.StringT)).
					Slice("device-write-iops", "Limit write rate (IO per second)", ox.Elem(ox.StringT)).
					Bool("disable-content-trust", "Skip image verification (default").
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("domainname", "Container NIS domain name").
					String("entrypoint", "Overwrite the default ENTRYPOINT").
					Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
					Slice("env-file", "Read in a file of environment", ox.Elem(ox.StringT)).
					Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
					String("gpus", "GPU devices to add to the", ox.Spec("gpu-request")).
					Slice("group-add", "Add additional groups to join", ox.Elem(ox.StringT)).
					String("health-cmd", "Command to run to check health").
					Duration("health-interval", "Time between running the check").
					Int("health-retries", "Consecutive failures needed to").
					Duration("health-start-interval", "Time between running the check").
					Duration("health-start-period", "Start period for the container").
					Duration("health-timeout", "Maximum time to allow one check").
					String("hostname", "Container host name", ox.Short("h")).
					Bool("init", "Run an init inside the container").
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
					Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT)).
					Uint("io-maxiops", "Maximum IOps limit for the").
					String("ip", "IPv4 address (e.g., 172.30.100.104)").
					String("ip6", "IPv6 address (e.g., 2001:db8::33)").
					String("ipc", "IPC mode to use").
					String("isolation", "Container isolation technology").
					Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT)).
					Slice("label", "Set meta data on a container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of", ox.Elem(ox.StringT)).
					Slice("link", "Add link to another container", ox.Elem(ox.StringT)).
					Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Elem(ox.StringT)).
					String("log-driver", "Logging driver for the container").
					Slice("log-opt", "Log driver options", ox.Elem(ox.StringT)).
					String("mac-address", "Container MAC address (e.g.,").
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
					Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT)).
					Int("memory-swappiness", "Tune container memory swappiness").
					String("mount", "Attach a filesystem mount to the", ox.Spec("mount")).
					String("name", "Assign a name to the container").
					String("network", "Connect a container to a network", ox.Spec("network")).
					Slice("network-alias", "Add network-scoped alias for the", ox.Elem(ox.StringT)).
					Bool("no-healthcheck", "Disable any container-specified").
					Bool("oom-kill-disable", "Disable OOM Killer").
					Int("oom-score-adj", "Tune host's OOM preferences").
					String("pid", "PID namespace to use").
					Int("pids-limit", "Tune container pids limit (set").
					String("platform", "Set platform if server is").
					Bool("privileged", "Give extended privileges to this").
					Slice("publish", "Publish a container's port(s) to", ox.Elem(ox.StringT), ox.Short("p")).
					Bool("publish-all", "Publish all exposed ports to", ox.Short("P")).
					String("pull", "Pull image before running").
					Bool("quiet", "Suppress the pull output", ox.Short("q")).
					Bool("read-only", "Mount the container's root").
					String("restart", "Restart policy to apply when a").
					Bool("rm", "Automatically remove the").
					String("runtime", "Runtime to use for this container").
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT)).
					Bool("sig-proxy", "Proxy received signals to the").
					String("stop-signal", "Signal to stop the container").
					Int("stop-timeout", "Timeout (in seconds) to stop a").
					Slice("storage-opt", "Storage driver options for the", ox.Elem(ox.StringT)).
					Map("sysctl", "Sysctl options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
					Slice("tmpfs", "Mount a tmpfs directory", ox.Elem(ox.StringT)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")).
					String("user", "Username or UID (format:", ox.Short("u")).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					Slice("volume", "Bind mount a volume", ox.Elem(ox.StringT), ox.Short("v")).
					String("volume-driver", "Optional volume driver for the").
					Slice("volumes-from", "Mount volumes from the specified", ox.Elem(ox.StringT)).
					String("workdir", "Working directory inside the", ox.Short("w")),
			), ox.Sub(
				ox.Usage("start", "Start one or more stopped containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker start"),
				ox.Section(0),
				ox.Flags().
					Bool("attach", "Attach STDOUT/STDERR and forward signals", ox.Short("a")).
					String("checkpoint", "Restore from this checkpoint").
					String("checkpoint-dir", "Use a custom checkpoint storage directory").
					String("detach-keys", "Override the key sequence for detaching a").
					Bool("interactive", "Attach container's STDIN", ox.Short("i")),
			), ox.Sub(
				ox.Usage("stats", "Display a live stream of container(s) resource usage statistics"),
				ox.Spec("[OPTIONS] [CONTAINER...]"),
				ox.Aliases("docker stats"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a")).
					String("format", "Format output using a custom template:").
					Bool("no-stream", "Disable streaming stats and only pull the first result").
					Bool("no-trunc", "Do not truncate output"),
			), ox.Sub(
				ox.Usage("stop", "Stop one or more running containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker stop"),
				ox.Section(0),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s")).
					Int("time", "Seconds to wait before killing the container", ox.Short("t")),
			), ox.Sub(
				ox.Usage("top", "Display the running processes of a container"),
				ox.Spec("CONTAINER [ps OPTIONS]"),
				ox.Aliases("docker top"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("unpause", "Unpause all processes within one or more containers"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("docker unpause"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("update", "Update configuration of one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("docker update"),
				ox.Section(0),
				ox.Flags().
					Uint16("blkio-weight", "Block IO (relative weight), between 10").
					Int("cpu-period", "Limit CPU CFS (Completely Fair").
					Int("cpu-quota", "Limit CPU CFS (Completely Fair").
					Int("cpu-rt-period", "Limit the CPU real-time period in").
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in").
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)").
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
					Slice("memory-swap", "Swap limit equal to memory plus swap:", ox.Elem(ox.UintT)).
					Int("pids-limit", "Tune container pids limit (set -1 for").
					String("restart", "Restart policy to apply when a"),
			), ox.Sub(
				ox.Usage("wait", "Block until one or more containers stop, then print their exit codes"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("docker wait"),
				ox.Section(0),
			)), ox.Sub(
			ox.Usage("context", "Manage contexts"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker context COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("export", "Export a context to a tar archive FILE or a tar stream on STDOUT."),
				ox.Spec("[OPTIONS] CONTEXT [FILE|-]"),
				ox.Section(0),
				ox.Footer("Export a context to a tar archive FILE or a tar stream on STDOUT."),
			), ox.Sub(
				ox.Usage("import", "Import a context from a tar or zip file"),
				ox.Spec("CONTEXT FILE|-"),
				ox.Section(0),
				ox.Footer("Import a context from a tar or zip file"),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more contexts"),
				ox.Spec("[OPTIONS] [CONTEXT] [CONTEXT...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")),
			), ox.Sub(
				ox.Usage("ls", "List contexts"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker context list"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only show context names", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more contexts"),
				ox.Spec("CONTEXT [CONTEXT...]"),
				ox.Aliases("docker context remove"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the removal of a context in use", ox.Short("f")),
			), ox.Sub(
				ox.Usage("show", "Print the name of the current context"),
				ox.Section(0),
				ox.Footer("Print the name of the current context"),
			), ox.Sub(
				ox.Usage("update", "Update a context"),
				ox.Spec("[OPTIONS] CONTEXT"),
				ox.Example("\n\n$ docker context update my-context --description \"some description\" --docker \"host=tcp://myserver:2376,ca=~/ca-file,cert=~/cert-file,key=~/key-file\""),
				ox.Section(0),
				ox.Flags().
					String("description", "Description of the context").
					Map("docker", "set the docker endpoint", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("[]")),
			), ox.Sub(
				ox.Usage("use", "Set the current docker context"),
				ox.Spec("CONTEXT"),
				ox.Section(0),
				ox.Footer("Set the current docker context"),
			)), ox.Sub(
			ox.Usage("image", "Manage images"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker image COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("build", "Build an image from a Dockerfile"),
				ox.Spec("[OPTIONS] PATH | URL | -"),
				ox.Aliases("docker build", "docker builder build"),
				ox.Section(0),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Elem(ox.StringT)).
					Slice("build-arg", "Set build-time variables", ox.Elem(ox.StringT)).
					Slice("cache-from", "Images to consider as cache sources", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Set the parent cgroup for the \"RUN\"").
					Bool("compress", "Compress the build context using gzip").
					Int("cpu-period", "Limit the CPU CFS (Completely Fair").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair").
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)").
					Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
					String("file", "Name of the Dockerfile (Default is", ox.Short("f")).
					Bool("force-rm", "Always remove intermediate containers").
					String("iidfile", "Write the image ID to the file").
					String("isolation", "Container isolation technology").
					Slice("label", "Set metadata for an image", ox.Elem(ox.StringT)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
					Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT)).
					String("network", "Set the networking mode for the RUN").
					Bool("no-cache", "Do not use cache when building the image").
					String("platform", "Set platform if server is multi-platform").
					Bool("pull", "Always attempt to pull a newer version of").
					Bool("quiet", "Suppress the build output and print image", ox.Short("q")).
					Bool("rm", "Remove intermediate containers after a").
					Slice("security-opt", "Security options", ox.Elem(ox.StringT)).
					Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT)).
					Bool("squash", "Squash newly built layers into a single").
					Slice("tag", "Name and optionally a tag in the", ox.Elem(ox.StringT), ox.Short("t")).
					String("target", "Set the target build stage to build.").
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")),
			), ox.Sub(
				ox.Usage("history", "Show the history of an image"),
				ox.Spec("[OPTIONS] IMAGE"),
				ox.Aliases("docker history"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:").
					Bool("human", "Print sizes and dates in human readable format", ox.Short("H")).
					Bool("no-trunc", "Don't truncate output").
					Bool("quiet", "Only show image IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("import", "Import the contents from a tarball to create a filesystem image"),
				ox.Spec("[OPTIONS] file|URL|- [REPOSITORY[:TAG]]"),
				ox.Aliases("docker import"),
				ox.Section(0),
				ox.Flags().
					Slice("change", "Apply Dockerfile instruction to the created image", ox.Elem(ox.StringT), ox.Short("c")).
					String("message", "Set commit message for imported image", ox.Short("m")).
					String("platform", "Set platform if server is multi-platform capable"),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more images"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")),
			), ox.Sub(
				ox.Usage("load", "Load an image from a tar archive or STDIN"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker load"),
				ox.Section(0),
				ox.Flags().
					String("input", "Read from tar archive file, instead of STDIN", ox.Short("i")).
					Bool("quiet", "Suppress the load output", ox.Short("q")),
			), ox.Sub(
				ox.Usage("ls", "List images"),
				ox.Spec("[OPTIONS] [REPOSITORY[:TAG]]"),
				ox.Aliases("docker image list", "docker images"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a")).
					Bool("digests", "Show digests").
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("no-trunc", "Don't truncate output").
					Bool("quiet", "Only show image IDs", ox.Short("q")).
					Bool("tree", "List multi-platform images as a tree (EXPERIMENTAL)"),
			), ox.Sub(
				ox.Usage("prune", "Remove unused images"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Remove all unused images, not just dangling ones", ox.Short("a")).
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Usage("pull", "Download an image from a registry"),
				ox.Spec("[OPTIONS] NAME[:TAG|@DIGEST]"),
				ox.Aliases("docker pull"),
				ox.Section(0),
				ox.Flags().
					Bool("all-tags", "Download all tagged images in the repository", ox.Short("a")).
					Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
					String("platform", "Set platform if server is multi-platform").
					Bool("quiet", "Suppress verbose output", ox.Short("q")),
			), ox.Sub(
				ox.Usage("push", "Upload an image to a registry"),
				ox.Spec("[OPTIONS] NAME[:TAG]"),
				ox.Aliases("docker push"),
				ox.Section(0),
				ox.Flags().
					Bool("all-tags", "Push all tags of an image to the repository", ox.Short("a")).
					Bool("disable-content-trust", "Skip image signing", ox.Default("true")).
					String("platform", "Push a platform-specific manifest as a").
					Bool("quiet", "Suppress verbose output", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more images"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Aliases("docker image remove", "docker rmi"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force removal of the image", ox.Short("f")).
					Bool("no-prune", "Do not delete untagged parents"),
			), ox.Sub(
				ox.Usage("save", "Save one or more images to a tar archive (streamed to STDOUT by default)"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Aliases("docker save"),
				ox.Section(0),
				ox.Flags().
					String("output", "Write to a file, instead of STDOUT", ox.Short("o")),
			), ox.Sub(
				ox.Usage("tag", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
				ox.Spec("SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]"),
				ox.Aliases("docker tag"),
				ox.Section(0),
			)), ox.Sub(
			ox.Usage("manifest", "Manage Docker image manifests and manifest lists"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker manifest COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("annotate", "Add additional information to a local image manifest"),
				ox.Spec("[OPTIONS] MANIFEST_LIST MANIFEST"),
				ox.Section(0),
				ox.Flags().
					String("arch", "Set architecture").
					String("os", "Set operating system").
					Slice("os-features", "Set operating system feature", ox.Elem(ox.StringT)).
					String("os-version", "Set operating system version").
					String("variant", "Set architecture variant"),
			), ox.Sub(
				ox.Usage("create", "Create a local manifest list for annotating and pushing to a registry"),
				ox.Spec("MANIFEST_LIST MANIFEST [MANIFEST...]"),
				ox.Section(0),
				ox.Flags().
					Bool("amend", "Amend an existing manifest list", ox.Short("a")).
					Bool("insecure", "Allow communication with an insecure registry"),
			), ox.Sub(
				ox.Usage("inspect", "Display an image manifest, or manifest list"),
				ox.Spec("[OPTIONS] [MANIFEST_LIST] MANIFEST"),
				ox.Section(0),
				ox.Flags().
					Bool("insecure", "Allow communication with an insecure registry").
					Bool("verbose", "Output additional info including layers and platform", ox.Short("v")),
			), ox.Sub(
				ox.Usage("push", "Push a manifest list to a repository"),
				ox.Spec("[OPTIONS] MANIFEST_LIST"),
				ox.Section(0),
				ox.Flags().
					Bool("insecure", "Allow push to an insecure registry").
					Bool("purge", "Remove the local manifest list after push", ox.Short("p")),
			), ox.Sub(
				ox.Usage("rm", "Delete one or more manifest lists from local storage"),
				ox.Spec("MANIFEST_LIST [MANIFEST_LIST...]"),
				ox.Section(0),
			)), ox.Sub(
			ox.Usage("network", "Manage networks"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker network COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("connect", "Connect a container to a network"),
				ox.Spec("[OPTIONS] NETWORK CONTAINER"),
				ox.Section(0),
				ox.Flags().
					Slice("alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
					Slice("driver-opt", "driver options for the network", ox.Elem(ox.StringT)).
					String("ip", "IPv4 address (e.g., \"172.30.100.104\")").
					String("ip6", "IPv6 address (e.g., \"2001:db8::33\")").
					Slice("link", "Add link to another container", ox.Elem(ox.StringT)).
					Slice("link-local-ip", "Add a link-local address for the container", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Usage("create", "Create a network"),
				ox.Spec("[OPTIONS] NETWORK"),
				ox.Section(0),
				ox.Flags().
					Bool("attachable", "Enable manual container attachment").
					Map("aux-address", "Auxiliary IPv4 or IPv6 addresses used by", ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					String("config-from", "The network from which to copy the configuration").
					Bool("config-only", "Create a configuration only network").
					String("driver", "Driver to manage the Network", ox.Default("bridge"), ox.Short("d")).
					Slice("gateway", "IPv4 or IPv6 Gateway for the master subnet", ox.Elem(ox.StringT)).
					Bool("ingress", "Create swarm routing-mesh network").
					Bool("internal", "Restrict external access to the network").
					Slice("ip-range", "Allocate container ip from a sub-range", ox.Elem(ox.StringT)).
					String("ipam-driver", "IP Address Management Driver", ox.Default("default")).
					Map("ipam-opt", "Set IPAM driver specific options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
					Bool("ipv6", "Enable or disable IPv6 networking").
					Slice("label", "Set metadata on a network", ox.Elem(ox.StringT)).
					Map("opt", "Set driver specific options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]"), ox.Short("o")).
					String("scope", "Control the network's scope").
					Slice("subnet", "Subnet in CIDR format that represents a", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Usage("disconnect", "Disconnect a container from a network"),
				ox.Spec("[OPTIONS] NETWORK CONTAINER"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the container to disconnect from a network", ox.Short("f")),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more networks"),
				ox.Spec("[OPTIONS] NETWORK [NETWORK...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("verbose", "Verbose output for diagnostics", ox.Short("v")),
			), ox.Sub(
				ox.Usage("ls", "List networks"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker network list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"driver=bridge\")", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("no-trunc", "Do not truncate the output").
					Bool("quiet", "Only display network IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("prune", "Remove all unused networks"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more networks"),
				ox.Spec("NETWORK [NETWORK...]"),
				ox.Aliases("docker network remove"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Do not error if the network does not exist", ox.Short("f")),
			)), ox.Sub(
			ox.Usage("plugin", "Manage plugins"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker plugin COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory."),
				ox.Spec("[OPTIONS] PLUGIN PLUGIN-DATA-DIR"),
				ox.Section(0),
				ox.Flags().
					Bool("compress", "Compress the context using gzip"),
			), ox.Sub(
				ox.Usage("disable", "Disable a plugin"),
				ox.Spec("[OPTIONS] PLUGIN"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the disable of an active plugin", ox.Short("f")),
			), ox.Sub(
				ox.Usage("enable", "Enable a plugin"),
				ox.Spec("[OPTIONS] PLUGIN"),
				ox.Section(0),
				ox.Flags().
					Int("timeout", "HTTP client timeout (in seconds)", ox.Default("30")),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more plugins"),
				ox.Spec("[OPTIONS] PLUGIN [PLUGIN...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")),
			), ox.Sub(
				ox.Usage("install", "Install a plugin"),
				ox.Spec("[OPTIONS] PLUGIN [KEY=VALUE...]"),
				ox.Section(0),
				ox.Flags().
					String("alias", "Local name for plugin").
					Bool("disable", "Do not enable the plugin on install").
					Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
					Bool("grant-all-permissions", "Grant all permissions necessary to run"),
			), ox.Sub(
				ox.Usage("ls", "List plugins"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker plugin list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"enabled=true\")", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("no-trunc", "Don't truncate output").
					Bool("quiet", "Only display plugin IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("push", "Push a plugin to a registry"),
				ox.Spec("[OPTIONS] PLUGIN[:TAG]"),
				ox.Section(0),
				ox.Flags().
					Bool("disable-content-trust", "Skip image signing", ox.Default("true")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more plugins"),
				ox.Spec("[OPTIONS] PLUGIN [PLUGIN...]"),
				ox.Aliases("docker plugin remove"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the removal of an active plugin", ox.Short("f")),
			), ox.Sub(
				ox.Usage("set", "Change settings for a plugin"),
				ox.Spec("PLUGIN KEY=VALUE [KEY=VALUE...]"),
				ox.Section(0),
				ox.Footer("Change settings for a plugin"),
			), ox.Sub(
				ox.Usage("upgrade", "Upgrade an existing plugin"),
				ox.Spec("[OPTIONS] PLUGIN [REMOTE]"),
				ox.Section(0),
				ox.Flags().
					Bool("disable-content-trust", "Skip image verification", ox.Default("true")).
					Bool("grant-all-permissions", "Grant all permissions necessary to run").
					Bool("skip-remote-check", "Do not check if specified remote plugin"),
			)), ox.Sub(
			ox.Usage("system", "Manage Docker"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker system COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("df", "Show docker disk usage"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:").
					Bool("verbose", "Show detailed information on space usage", ox.Short("v")),
			), ox.Sub(
				ox.Usage("events", "Get real time events from the server"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker events"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					String("since", "Show all events created since timestamp").
					String("until", "Stream events until this timestamp"),
			), ox.Sub(
				ox.Usage("info", "Display system-wide information"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker info"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")),
			), ox.Sub(
				ox.Usage("prune", "Remove unused data"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Remove all unused images not just dangling ones", ox.Short("a")).
					String("filter", "Provide filter values (e.g. \"label=<key>=<value>\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")).
					Bool("volumes", "Prune anonymous volumes"),
			)), ox.Sub(
			ox.Usage("trust", "Manage trust on Docker images"),
			ox.Spec("COMMAND"),
			ox.Sections("Management Commands", "Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker trust COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("key", "Manage keys for signing Docker images"),
				ox.Spec("COMMAND"),
				ox.Sections("Commands"),
				ox.Section(0),
				ox.Footer("Run 'docker trust key COMMAND --help' for more information on a command."),
				ox.Sub(
					ox.Usage("generate", "Generate and load a signing key-pair"),
					ox.Spec("NAME"),
					ox.Section(0),
					ox.Flags().
						String("dir", "Directory to generate key in, defaults to current"),
				), ox.Sub(
					ox.Usage("load", "Load a private key file for signing"),
					ox.Spec("[OPTIONS] KEYFILE"),
					ox.Section(0),
					ox.Flags().
						String("name", "Name for the loaded key", ox.Default("signer")),
				)), ox.Sub(
				ox.Usage("signer", "Manage entities who can sign Docker images"),
				ox.Spec("COMMAND"),
				ox.Sections("Commands"),
				ox.Section(0),
				ox.Footer("Run 'docker trust signer COMMAND --help' for more information on a command."),
				ox.Sub(
					ox.Usage("add", "Add a signer"),
					ox.Spec("OPTIONS NAME REPOSITORY [REPOSITORY...]"),
					ox.Section(0),
					ox.Flags().
						Slice("key", "Path to the signer's public key file", ox.Elem(ox.StringT)),
				), ox.Sub(
					ox.Usage("remove", "Remove a signer"),
					ox.Spec("[OPTIONS] NAME REPOSITORY [REPOSITORY...]"),
					ox.Section(0),
					ox.Flags().
						Bool("force", "Do not prompt for confirmation before removing the most", ox.Short("f")),
				)), ox.Sub(
				ox.Usage("inspect", "Return low-level information about keys and signatures"),
				ox.Spec("IMAGE[:TAG] [IMAGE[:TAG]...]"),
				ox.Section(1),
				ox.Flags().
					Bool("pretty", "Print the information in a human friendly format"),
			), ox.Sub(
				ox.Usage("revoke", "Remove trust for an image"),
				ox.Spec("[OPTIONS] IMAGE[:TAG]"),
				ox.Section(1),
				ox.Flags().
					Bool("yes", "Do not prompt for confirmation", ox.Short("y")),
			), ox.Sub(
				ox.Usage("sign", "Sign an image"),
				ox.Spec("IMAGE:TAG"),
				ox.Section(1),
				ox.Flags().
					Bool("local", "Sign a locally tagged image"),
			)), ox.Sub(
			ox.Usage("volume", "Manage volumes"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker volume COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a volume"),
				ox.Spec("[OPTIONS] [VOLUME]"),
				ox.Section(0),
				ox.Flags().
					String("availability", "Cluster Volume availability (\"active\",").
					String("driver", "Specify volume driver name", ox.Default("local"), ox.Short("d")).
					String("group", "Cluster Volume group (cluster volumes)").
					Slice("label", "Set metadata for a volume", ox.Elem(ox.StringT)).
					Slice("limit-bytes", "Minimum size of the Cluster Volume in bytes", ox.Elem(ox.UintT)).
					Map("opt", "Set driver specific options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]"), ox.Short("o")).
					Slice("required-bytes", "Maximum size of the Cluster Volume in bytes", ox.Elem(ox.UintT)).
					String("scope", "Cluster Volume access scope (\"single\",").
					Map("secret", "Cluster Volume secrets", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
					String("sharing", "Cluster Volume access sharing (\"none\",").
					Slice("topology-preferred", "A topology that the Cluster Volume", ox.Elem(ox.StringT)).
					Slice("topology-required", "A topology that the Cluster Volume must", ox.Elem(ox.StringT)).
					String("type", "Cluster Volume access type (\"mount\","),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more volumes"),
				ox.Spec("[OPTIONS] VOLUME [VOLUME...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")),
			), ox.Sub(
				ox.Usage("ls", "List volumes"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker volume list"),
				ox.Section(0),
				ox.Flags().
					Bool("cluster", "Display only cluster volumes, and use cluster").
					String("filter", "Provide filter values (e.g. \"dangling=true\")", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display volume names", ox.Short("q")),
			), ox.Sub(
				ox.Usage("prune", "Remove unused local volumes"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("all", "Remove all unused volumes, not just anonymous ones", ox.Short("a")).
					String("filter", "Provide filter values (e.g. \"label=<label>\")", ox.Spec("filter")).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more volumes"),
				ox.Spec("[OPTIONS] VOLUME [VOLUME...]"),
				ox.Aliases("docker volume remove"),
				ox.Example("\n\n$ docker volume rm hello\nhello"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force the removal of one or more volumes", ox.Short("f")),
			), ox.Sub(
				ox.Usage("update", "Update a volume (cluster volumes only)"),
				ox.Spec("[OPTIONS] [VOLUME]"),
				ox.Section(0),
				ox.Flags().
					String("availability", "Cluster Volume availability (\"active\","),
			)), ox.Sub(
			ox.Usage("config", "Manage Swarm configs"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker config COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a config from a file or STDIN"),
				ox.Spec("[OPTIONS] CONFIG file|-"),
				ox.Section(0),
				ox.Flags().
					Slice("label", "Config labels", ox.Elem(ox.StringT), ox.Short("l")).
					String("template-driver", "Template driver"),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more configs"),
				ox.Spec("[OPTIONS] CONFIG [CONFIG...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("pretty", "Print the information in a human friendly format"),
			), ox.Sub(
				ox.Usage("ls", "List configs"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker config list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more configs"),
				ox.Spec("CONFIG [CONFIG...]"),
				ox.Aliases("docker config remove"),
				ox.Section(0),
			)), ox.Sub(
			ox.Usage("node", "Manage Swarm nodes"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker node COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("demote", "Demote one or more nodes from manager in the swarm"),
				ox.Spec("NODE [NODE...]"),
				ox.Section(0),
				ox.Footer("Demote one or more nodes from manager in the swarm"),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more nodes"),
				ox.Spec("[OPTIONS] self|NODE [NODE...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("pretty", "Print the information in a human friendly format"),
			), ox.Sub(
				ox.Usage("ls", "List nodes in the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker node list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("promote", "Promote one or more nodes to manager in the swarm"),
				ox.Spec("NODE [NODE...]"),
				ox.Section(0),
				ox.Footer("Promote one or more nodes to manager in the swarm"),
			), ox.Sub(
				ox.Usage("ps", "List tasks running on one or more nodes, defaults to current node"),
				ox.Spec("[OPTIONS] [NODE...]"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Pretty-print tasks using a Go template").
					Bool("no-resolve", "Do not map IDs to Names").
					Bool("no-trunc", "Do not truncate output").
					Bool("quiet", "Only display task IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more nodes from the swarm"),
				ox.Spec("[OPTIONS] NODE [NODE...]"),
				ox.Aliases("docker node remove"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force remove a node from the swarm", ox.Short("f")),
			), ox.Sub(
				ox.Usage("update", "Update a node"),
				ox.Spec("[OPTIONS] NODE"),
				ox.Section(0),
				ox.Flags().
					String("availability", "Availability of the node (\"active\",").
					Slice("label-add", "Add or update a node label (\"key=value\")", ox.Elem(ox.StringT)).
					Slice("label-rm", "Remove a node label if exists", ox.Elem(ox.StringT)).
					String("role", "Role of the node (\"worker\", \"manager\")"),
			)), ox.Sub(
			ox.Usage("secret", "Manage Swarm secrets"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker secret COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a secret from a file or STDIN as content"),
				ox.Spec("[OPTIONS] SECRET [file|-]"),
				ox.Section(0),
				ox.Flags().
					String("driver", "Secret driver", ox.Short("d")).
					Slice("label", "Secret labels", ox.Elem(ox.StringT), ox.Short("l")).
					String("template-driver", "Template driver"),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more secrets"),
				ox.Spec("[OPTIONS] SECRET [SECRET...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("pretty", "Print the information in a human friendly format"),
			), ox.Sub(
				ox.Usage("ls", "List secrets"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker secret list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more secrets"),
				ox.Spec("SECRET [SECRET...]"),
				ox.Aliases("docker secret remove"),
				ox.Section(0),
			)), ox.Sub(
			ox.Usage("service", "Manage Swarm services"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker service COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a new service"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Section(0),
				ox.Flags().
					Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
					String("config", "Specify configurations to", ox.Spec("config")).
					Slice("constraint", "Placement constraints", ox.Elem(ox.StringT)).
					Slice("container-label", "Container labels", ox.Elem(ox.StringT)).
					String("credential-spec", "Credential spec for managed", ox.Spec("credential-spec")).
					Bool("detach", "Exit immediately instead of", ox.Short("d")).
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("endpoint-mode", "Endpoint mode (vip or dnsrr)").
					String("entrypoint", "Overwrite the default", ox.Spec("command")).
					Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
					Slice("env-file", "Read in a file of environment", ox.Elem(ox.StringT)).
					Slice("generic-resource", "User defined resources", ox.Elem(ox.StringT)).
					Slice("group", "Set one or more supplementary", ox.Elem(ox.StringT)).
					String("health-cmd", "Command to run to check health").
					Duration("health-interval", "Time between running the check").
					Int("health-retries", "Consecutive failures needed to").
					Duration("health-start-interval", "Time between running the check").
					Duration("health-start-period", "Start period for the container").
					Duration("health-timeout", "Maximum time to allow one").
					Slice("host", "Set one or more custom", ox.Elem(ox.StringT)).
					String("hostname", "Container hostname").
					Bool("init", "Use an init inside each").
					String("isolation", "Service container isolation mode").
					Slice("label", "Service labels", ox.Elem(ox.StringT), ox.Short("l")).
					Float32("limit-cpu", "Limit CPUs", ox.Spec("decimal")).
					Slice("limit-memory", "Limit Memory", ox.Elem(ox.UintT)).
					Int("limit-pids", "Limit maximum number of").
					String("log-driver", "Logging driver for service").
					Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
					Uint("max-concurrent", "Number of job tasks to run").
					String("mode", "Service mode (\"replicated\",").
					String("mount", "Attach a filesystem mount to", ox.Spec("mount")).
					String("name", "Service name").
					String("network", "Network attachments", ox.Spec("network")).
					Bool("no-healthcheck", "Disable any").
					Bool("no-resolve-image", "Do not query the registry to").
					Int("oom-score-adj", "Tune host's OOM preferences").
					String("placement-pref", "Add a placement preference", ox.Spec("pref")).
					Uint("publish", "Publish a port as a node port", ox.Spec("port"), ox.Short("p")).
					Bool("quiet", "Suppress progress output", ox.Short("q")).
					Bool("read-only", "Mount the container's root").
					Uint("replicas", "Number of tasks").
					Uint("replicas-max-per-node", "Maximum number of tasks per").
					Float32("reserve-cpu", "Reserve CPUs", ox.Spec("decimal")).
					Slice("reserve-memory", "Reserve Memory", ox.Elem(ox.UintT)).
					String("restart-condition", "Restart when condition is met").
					Duration("restart-delay", "Delay between restart attempts").
					Uint("restart-max-attempts", "Maximum number of restarts").
					Duration("restart-window", "Window used to evaluate the").
					Duration("rollback-delay", "Delay between task rollbacks").
					String("rollback-failure-action", "Action on rollback failure").
					Float32("rollback-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float")).
					Duration("rollback-monitor", "Duration after each task").
					String("rollback-order", "Rollback order (\"start-first\",").
					Uint("rollback-parallelism", "Maximum number of tasks rolled").
					String("secret", "Specify secrets to expose to", ox.Spec("secret")).
					Duration("stop-grace-period", "Time to wait before force").
					String("stop-signal", "Signal to stop the container").
					Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")).
					Duration("update-delay", "Delay between updates").
					String("update-failure-action", "Action on update failure").
					Float32("update-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float")).
					Duration("update-monitor", "Duration after each task").
					String("update-order", "Update order (\"start-first\",").
					Uint("update-parallelism", "Maximum number of tasks").
					String("user", "Username or UID (format:", ox.Short("u")).
					Bool("with-registry-auth", "Send registry authentication").
					String("workdir", "Working directory inside the", ox.Short("w")),
			), ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more services"),
				ox.Spec("[OPTIONS] SERVICE [SERVICE...]"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f")).
					Bool("pretty", "Print the information in a human friendly format"),
			), ox.Sub(
				ox.Usage("logs", "Fetch the logs of a service or task"),
				ox.Spec("[OPTIONS] SERVICE|TASK"),
				ox.Section(0),
				ox.Flags().
					Bool("details", "Show extra details provided to logs").
					Bool("follow", "Follow log output", ox.Short("f")).
					Bool("no-resolve", "Do not map IDs to Names in output").
					Bool("no-task-ids", "Do not include task IDs in output").
					Bool("no-trunc", "Do not truncate output").
					Bool("raw", "Do not neatly format logs").
					String("since", "Show logs since timestamp (e.g.").
					String("tail", "Number of lines to show from the end of the logs", ox.Short("n")).
					Bool("timestamps", "Show timestamps", ox.Short("t")),
			), ox.Sub(
				ox.Usage("ls", "List services"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker service list"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("ps", "List the tasks of one or more services"),
				ox.Spec("[OPTIONS] SERVICE [SERVICE...]"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Pretty-print tasks using a Go template").
					Bool("no-resolve", "Do not map IDs to Names").
					Bool("no-trunc", "Do not truncate output").
					Bool("quiet", "Only display task IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more services"),
				ox.Spec("SERVICE [SERVICE...]"),
				ox.Aliases("docker service remove"),
				ox.Section(0),
			), ox.Sub(
				ox.Usage("rollback", "Revert changes to a service's configuration"),
				ox.Spec("[OPTIONS] SERVICE"),
				ox.Section(0),
				ox.Flags().
					Bool("detach", "Exit immediately instead of waiting for the service to", ox.Short("d")).
					Bool("quiet", "Suppress progress output", ox.Short("q")),
			), ox.Sub(
				ox.Usage("scale", "Scale one or multiple replicated services"),
				ox.Spec("SERVICE=REPLICAS [SERVICE=REPLICAS...]"),
				ox.Section(0),
				ox.Flags().
					Bool("detach", "Exit immediately instead of waiting for the service to", ox.Short("d")),
			), ox.Sub(
				ox.Usage("update", "Update a service"),
				ox.Spec("[OPTIONS] SERVICE"),
				ox.Section(0),
				ox.Flags().
					String("args", "Service command args", ox.Spec("command")).
					Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
					String("config-add", "Add or update a config file on", ox.Spec("config")).
					Slice("config-rm", "Remove a configuration file", ox.Elem(ox.StringT)).
					Slice("constraint-add", "Add or update a placement", ox.Elem(ox.StringT)).
					Slice("constraint-rm", "Remove a constraint", ox.Elem(ox.StringT)).
					Slice("container-label-add", "Add or update a container label", ox.Elem(ox.StringT)).
					Slice("container-label-rm", "Remove a container label by its key", ox.Elem(ox.StringT)).
					String("credential-spec", "Credential spec for managed", ox.Spec("credential-spec")).
					Bool("detach", "Exit immediately instead of", ox.Short("d")).
					Slice("dns-add", "Add or update a custom DNS server", ox.Elem(ox.StringT)).
					Slice("dns-option-add", "Add or update a DNS option", ox.Elem(ox.StringT)).
					Slice("dns-option-rm", "Remove a DNS option", ox.Elem(ox.StringT)).
					Slice("dns-rm", "Remove a custom DNS server", ox.Elem(ox.StringT)).
					Slice("dns-search-add", "Add or update a custom DNS", ox.Elem(ox.StringT)).
					Slice("dns-search-rm", "Remove a DNS search domain", ox.Elem(ox.StringT)).
					String("endpoint-mode", "Endpoint mode (vip or dnsrr)").
					String("entrypoint", "Overwrite the default", ox.Spec("command")).
					Slice("env-add", "Add or update an environment", ox.Elem(ox.StringT)).
					Slice("env-rm", "Remove an environment variable", ox.Elem(ox.StringT)).
					Bool("force", "Force update even if no").
					Slice("generic-resource-add", "Add a Generic resource", ox.Elem(ox.StringT)).
					Slice("generic-resource-rm", "Remove a Generic resource", ox.Elem(ox.StringT)).
					Slice("group-add", "Add an additional", ox.Elem(ox.StringT)).
					Slice("group-rm", "Remove a previously added", ox.Elem(ox.StringT)).
					String("health-cmd", "Command to run to check health").
					Duration("health-interval", "Time between running the check").
					Int("health-retries", "Consecutive failures needed to").
					Duration("health-start-interval", "Time between running the check").
					Duration("health-start-period", "Start period for the container").
					Duration("health-timeout", "Maximum time to allow one").
					Slice("host-add", "Add a custom host-to-IP", ox.Elem(ox.StringT)).
					Slice("host-rm", "Remove a custom host-to-IP", ox.Elem(ox.StringT)).
					String("hostname", "Container hostname").
					String("image", "Service image tag").
					Bool("init", "Use an init inside each").
					String("isolation", "Service container isolation mode").
					Slice("label-add", "Add or update a service label", ox.Elem(ox.StringT)).
					Slice("label-rm", "Remove a label by its key", ox.Elem(ox.StringT)).
					Float32("limit-cpu", "Limit CPUs", ox.Spec("decimal")).
					Slice("limit-memory", "Limit Memory", ox.Elem(ox.UintT)).
					Int("limit-pids", "Limit maximum number of").
					String("log-driver", "Logging driver for service").
					Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
					Uint("max-concurrent", "Number of job tasks to run").
					String("mount-add", "Add or update a mount on a service", ox.Spec("mount")).
					Slice("mount-rm", "Remove a mount by its target path", ox.Elem(ox.StringT)).
					String("network-add", "Add a network", ox.Spec("network")).
					Slice("network-rm", "Remove a network", ox.Elem(ox.StringT)).
					Bool("no-healthcheck", "Disable any").
					Bool("no-resolve-image", "Do not query the registry to").
					Int("oom-score-adj", "Tune host's OOM preferences").
					String("placement-pref-add", "Add a placement preference", ox.Spec("pref")).
					String("placement-pref-rm", "Remove a placement preference", ox.Spec("pref")).
					Uint("publish-add", "Add or update a published port", ox.Spec("port")).
					Uint("publish-rm", "Remove a published port by its", ox.Spec("port")).
					Bool("quiet", "Suppress progress output", ox.Short("q")).
					Bool("read-only", "Mount the container's root").
					Uint("replicas", "Number of tasks").
					Uint("replicas-max-per-node", "Maximum number of tasks per").
					Float32("reserve-cpu", "Reserve CPUs", ox.Spec("decimal")).
					Slice("reserve-memory", "Reserve Memory", ox.Elem(ox.UintT)).
					String("restart-condition", "Restart when condition is met").
					Duration("restart-delay", "Delay between restart attempts").
					Uint("restart-max-attempts", "Maximum number of restarts").
					Duration("restart-window", "Window used to evaluate the").
					Bool("rollback", "Rollback to previous specification").
					Duration("rollback-delay", "Delay between task rollbacks").
					String("rollback-failure-action", "Action on rollback failure").
					Float32("rollback-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float")).
					Duration("rollback-monitor", "Duration after each task").
					String("rollback-order", "Rollback order (\"start-first\",").
					Uint("rollback-parallelism", "Maximum number of tasks rolled").
					String("secret-add", "Add or update a secret on a service", ox.Spec("secret")).
					Slice("secret-rm", "Remove a secret", ox.Elem(ox.StringT)).
					Duration("stop-grace-period", "Time to wait before force").
					String("stop-signal", "Signal to stop the container").
					Slice("sysctl-add", "Add or update a Sysctl option", ox.Elem(ox.StringT)).
					Slice("sysctl-rm", "Remove a Sysctl option", ox.Elem(ox.StringT)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
					String("ulimit-add", "Add or update a ulimit option", ox.Spec("ulimit")).
					Slice("ulimit-rm", "Remove a ulimit option", ox.Elem(ox.StringT)).
					Duration("update-delay", "Delay between updates").
					String("update-failure-action", "Action on update failure").
					Float32("update-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float")).
					Duration("update-monitor", "Duration after each task").
					String("update-order", "Update order (\"start-first\",").
					Uint("update-parallelism", "Maximum number of tasks").
					String("user", "Username or UID (format:", ox.Short("u")).
					Bool("with-registry-auth", "Send registry authentication").
					String("workdir", "Working directory inside the", ox.Short("w")),
			)), ox.Sub(
			ox.Usage("stack", "Manage Swarm stacks"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker stack COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("config", "Outputs the final config file, after doing merges and interpolations"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Slice("compose-file", "Path to a Compose file, or \"-\" to read", ox.Elem(ox.StringT), ox.Short("c")).
					Bool("skip-interpolation", "Skip interpolation and output only merged"),
			), ox.Sub(
				ox.Usage("deploy", "Deploy a new stack or update an existing stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Aliases("docker stack up"),
				ox.Section(0),
				ox.Flags().
					Slice("compose-file", "Path to a Compose file, or \"-\" to read", ox.Elem(ox.StringT), ox.Short("c")).
					Bool("detach", "Exit immediately instead of waiting for", ox.Short("d")).
					Bool("prune", "Prune services that are no longer referenced").
					Bool("quiet", "Suppress progress output", ox.Short("q")).
					String("resolve-image", "Query the registry to resolve image digest").
					Bool("with-registry-auth", "Send registry authentication details to"),
			), ox.Sub(
				ox.Usage("ls", "List stacks"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("docker stack list"),
				ox.Section(0),
				ox.Flags().
					String("format", "Format output using a custom template:"),
			), ox.Sub(
				ox.Usage("ps", "List the tasks in the stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("no-resolve", "Do not map IDs to Names").
					Bool("no-trunc", "Do not truncate output").
					Bool("quiet", "Only display task IDs", ox.Short("q")),
			), ox.Sub(
				ox.Usage("rm", "Remove one or more stacks"),
				ox.Spec("[OPTIONS] STACK [STACK...]"),
				ox.Aliases("docker stack remove", "docker stack down"),
				ox.Section(0),
				ox.Flags().
					Bool("detach", "Do not wait for stack removal", ox.Default("true"), ox.Short("d")),
			), ox.Sub(
				ox.Usage("services", "List the services in the stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Section(0),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
					String("format", "Format output using a custom template:").
					Bool("quiet", "Only display IDs", ox.Short("q")),
			)), ox.Sub(
			ox.Usage("swarm", "Manage Swarm"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker swarm COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("ca", "Display and rotate the root CA"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					String("ca-cert", "Path to the PEM-formatted root CA", ox.Spec("pem-file")).
					String("ca-key", "Path to the PEM-formatted root CA key", ox.Spec("pem-file")).
					Duration("cert-expiry", "Validity period for node certificates").
					Bool("detach", "Exit immediately instead of waiting for", ox.Short("d")).
					String("external-ca", "Specifications of one or more", ox.Spec("external-ca")).
					Bool("quiet", "Suppress progress output", ox.Short("q")).
					Bool("rotate", "Rotate the swarm CA - if no certificate"),
			), ox.Sub(
				ox.Usage("init", "Initialize a swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					String("advertise-addr", "Advertised address").
					Bool("autolock", "Enable manager autolocking").
					String("availability", "Availability of the node").
					Duration("cert-expiry", "Validity period for node").
					String("data-path-addr", "Address or interface to").
					Uint32("data-path-port", "Port number to use for").
					Slice("default-addr-pool", "default address pool in", ox.Elem(ox.CIDRT)).
					Uint32("default-addr-pool-mask-length", "default address pool").
					Duration("dispatcher-heartbeat", "Dispatcher heartbeat").
					String("external-ca", "Specifications of one or", ox.Spec("external-ca")).
					Bool("force-new-cluster", "Force create a new cluster").
					AddrPort("listen-addr", "Listen address (format:").
					Uint("max-snapshots", "Number of additional Raft").
					Uint("snapshot-interval", "Number of log entries").
					Int("task-history-limit", "Task history retention"),
			), ox.Sub(
				ox.Usage("join", "Join a swarm as a node and/or manager"),
				ox.Spec("[OPTIONS] HOST:PORT"),
				ox.Section(0),
				ox.Flags().
					String("advertise-addr", "Advertised address (format:").
					String("availability", "Availability of the node (\"active\",").
					String("data-path-addr", "Address or interface to use for data path").
					AddrPort("listen-addr", "Listen address (format:").
					String("token", "Token for entry into the swarm"),
			), ox.Sub(
				ox.Usage("join-token", "Manage join tokens"),
				ox.Spec("[OPTIONS] (worker|manager)"),
				ox.Section(0),
				ox.Flags().
					Bool("quiet", "Only display token", ox.Short("q")).
					Bool("rotate", "Rotate join token"),
			), ox.Sub(
				ox.Usage("leave", "Leave the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("force", "Force this node to leave the swarm, ignoring warnings", ox.Short("f")),
			), ox.Sub(
				ox.Usage("unlock", "Unlock swarm"),
				ox.Section(0),
				ox.Footer("Unlock swarm"),
			), ox.Sub(
				ox.Usage("unlock-key", "Manage the unlock key"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("quiet", "Only display token", ox.Short("q")).
					Bool("rotate", "Rotate unlock key"),
			), ox.Sub(
				ox.Usage("update", "Update the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Flags().
					Bool("autolock", "Change manager autolocking").
					Duration("cert-expiry", "Validity period for node").
					Duration("dispatcher-heartbeat", "Dispatcher heartbeat period").
					String("external-ca", "Specifications of one or more", ox.Spec("external-ca")).
					Uint("max-snapshots", "Number of additional Raft").
					Uint("snapshot-interval", "Number of log entries between").
					Int("task-history-limit", "Task history retention limit"),
			)), ox.Sub(
			ox.Usage("attach", "Attach local standard input, output, and error streams to a running container"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("docker container attach", "docker attach"),
			ox.Section(3),
			ox.Flags().
				String("detach-keys", "Override the key sequence for detaching a").
				Bool("no-stdin", "Do not attach STDIN").
				Bool("sig-proxy", "Proxy all received signals to the process"),
		), ox.Sub(
			ox.Usage("commit", "Create a new image from a container's changes"),
			ox.Spec("[OPTIONS] CONTAINER [REPOSITORY[:TAG]]"),
			ox.Aliases("docker container commit", "docker commit"),
			ox.Section(3),
			ox.Flags().
				String("author", "Author (e.g., \"John Hannibal Smith", ox.Short("a")).
				Slice("change", "Apply Dockerfile instruction to the created image", ox.Elem(ox.StringT), ox.Short("c")).
				String("message", "Commit message", ox.Short("m")).
				Bool("pause", "Pause container during commit", ox.Default("true"), ox.Short("p")),
		), ox.Sub(
			ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
			ox.Spec("[OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-"),
			ox.Aliases("docker container cp", "docker cp"),
			ox.Section(3),
			ox.Flags().
				Bool("archive", "Archive mode (copy all uid/gid information)", ox.Short("a")).
				Bool("follow-link", "Always follow symbol link in SRC_PATH", ox.Short("L")).
				Bool("quiet", "Suppress progress output during copy. Progress", ox.Short("q")),
		), ox.Sub(
			ox.Usage("create", "Create a new container"),
			ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
			ox.Aliases("docker container create", "docker create"),
			ox.Section(3),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping", ox.Elem(ox.StringT)).
				Map("annotation", "Add an annotation to the", ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
				Uint16("blkio-weight", "Block IO (relative weight),").
				Slice("blkio-weight-device", "Block IO weight (relative device", ox.Elem(ox.StringT)).
				Slice("cap-add", "Add Linux capabilities", ox.Elem(ox.StringT)).
				Slice("cap-drop", "Drop Linux capabilities", ox.Elem(ox.StringT)).
				String("cgroup-parent", "Optional parent cgroup for the").
				String("cgroupns", "Cgroup namespace to use").
				String("cidfile", "Write the container ID to the file").
				Int("cpu-count", "CPU count (Windows only)").
				Int("cpu-percent", "CPU percent (Windows only)").
				Int("cpu-period", "Limit CPU CFS (Completely Fair").
				Int("cpu-quota", "Limit CPU CFS (Completely Fair").
				Int("cpu-rt-period", "Limit CPU real-time period in").
				Int("cpu-rt-runtime", "Limit CPU real-time runtime in").
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
				String("cpuset-cpus", "CPUs in which to allow execution").
				String("cpuset-mems", "MEMs in which to allow execution").
				Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Elem(ox.StringT)).
				Slice("device-read-bps", "Limit read rate (bytes per", ox.Elem(ox.StringT)).
				Slice("device-read-iops", "Limit read rate (IO per second)", ox.Elem(ox.StringT)).
				Slice("device-write-bps", "Limit write rate (bytes per", ox.Elem(ox.StringT)).
				Slice("device-write-iops", "Limit write rate (IO per second)", ox.Elem(ox.StringT)).
				Bool("disable-content-trust", "Skip image verification (default").
				Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
				Slice("dns-option", "Set DNS options", ox.Elem(ox.StringT)).
				Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
				String("domainname", "Container NIS domain name").
				String("entrypoint", "Overwrite the default ENTRYPOINT").
				Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
				Slice("env-file", "Read in a file of environment", ox.Elem(ox.StringT)).
				Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
				String("gpus", "GPU devices to add to the", ox.Spec("gpu-request")).
				Slice("group-add", "Add additional groups to join", ox.Elem(ox.StringT)).
				String("health-cmd", "Command to run to check health").
				Duration("health-interval", "Time between running the check").
				Int("health-retries", "Consecutive failures needed to").
				Duration("health-start-interval", "Time between running the check").
				Duration("health-start-period", "Start period for the container").
				Duration("health-timeout", "Maximum time to allow one check").
				String("hostname", "Container host name", ox.Short("h")).
				Bool("init", "Run an init inside the container").
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i")).
				Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT)).
				Uint("io-maxiops", "Maximum IOps limit for the").
				String("ip", "IPv4 address (e.g., 172.30.100.104)").
				String("ip6", "IPv6 address (e.g., 2001:db8::33)").
				String("ipc", "IPC mode to use").
				String("isolation", "Container isolation technology").
				Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT)).
				Slice("label", "Set meta data on a container", ox.Elem(ox.StringT), ox.Short("l")).
				Slice("label-file", "Read in a line delimited file of", ox.Elem(ox.StringT)).
				Slice("link", "Add link to another container", ox.Elem(ox.StringT)).
				Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Elem(ox.StringT)).
				String("log-driver", "Logging driver for the container").
				Slice("log-opt", "Log driver options", ox.Elem(ox.StringT)).
				String("mac-address", "Container MAC address (e.g.,").
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
				Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT)).
				Int("memory-swappiness", "Tune container memory swappiness").
				String("mount", "Attach a filesystem mount to the", ox.Spec("mount")).
				String("name", "Assign a name to the container").
				String("network", "Connect a container to a network", ox.Spec("network")).
				Slice("network-alias", "Add network-scoped alias for the", ox.Elem(ox.StringT)).
				Bool("no-healthcheck", "Disable any container-specified").
				Bool("oom-kill-disable", "Disable OOM Killer").
				Int("oom-score-adj", "Tune host's OOM preferences").
				String("pid", "PID namespace to use").
				Int("pids-limit", "Tune container pids limit (set").
				String("platform", "Set platform if server is").
				Bool("privileged", "Give extended privileges to this").
				Slice("publish", "Publish a container's port(s) to", ox.Elem(ox.StringT), ox.Short("p")).
				Bool("publish-all", "Publish all exposed ports to", ox.Short("P")).
				String("pull", "Pull image before creating").
				Bool("quiet", "Suppress the pull output", ox.Short("q")).
				Bool("read-only", "Mount the container's root").
				String("restart", "Restart policy to apply when a").
				Bool("rm", "Automatically remove the").
				String("runtime", "Runtime to use for this container").
				Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
				Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT)).
				String("stop-signal", "Signal to stop the container").
				Int("stop-timeout", "Timeout (in seconds) to stop a").
				Slice("storage-opt", "Storage driver options for the", ox.Elem(ox.StringT)).
				Map("sysctl", "Sysctl options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("map[]")).
				Slice("tmpfs", "Mount a tmpfs directory", ox.Elem(ox.StringT)).
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t")).
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]")).
				String("user", "Username or UID (format:", ox.Short("u")).
				String("userns", "User namespace to use").
				String("uts", "UTS namespace to use").
				Slice("volume", "Bind mount a volume", ox.Elem(ox.StringT), ox.Short("v")).
				String("volume-driver", "Optional volume driver for the").
				Slice("volumes-from", "Mount volumes from the specified", ox.Elem(ox.StringT)).
				String("workdir", "Working directory inside the", ox.Short("w")),
		), ox.Sub(
			ox.Usage("diff", "Inspect changes to files or directories on a container's filesystem"),
			ox.Spec("CONTAINER"),
			ox.Aliases("docker container diff", "docker diff"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("events", "Get real time events from the server"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("docker system events", "docker events"),
			ox.Section(3),
			ox.Flags().
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f")).
				String("format", "Format output using a custom template:").
				String("since", "Show all events created since timestamp").
				String("until", "Stream events until this timestamp"),
		), ox.Sub(
			ox.Usage("export", "Export a container's filesystem as a tar archive"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("docker container export", "docker export"),
			ox.Section(3),
			ox.Flags().
				String("output", "Write to a file, instead of STDOUT", ox.Short("o")),
		), ox.Sub(
			ox.Usage("history", "Show the history of an image"),
			ox.Spec("[OPTIONS] IMAGE"),
			ox.Aliases("docker image history", "docker history"),
			ox.Section(3),
			ox.Flags().
				String("format", "Format output using a custom template:").
				Bool("human", "Print sizes and dates in human readable format", ox.Short("H")).
				Bool("no-trunc", "Don't truncate output").
				Bool("quiet", "Only show image IDs", ox.Short("q")),
		), ox.Sub(
			ox.Usage("import", "Import the contents from a tarball to create a filesystem image"),
			ox.Spec("[OPTIONS] file|URL|- [REPOSITORY[:TAG]]"),
			ox.Aliases("docker image import", "docker import"),
			ox.Section(3),
			ox.Flags().
				Slice("change", "Apply Dockerfile instruction to the created image", ox.Elem(ox.StringT), ox.Short("c")).
				String("message", "Set commit message for imported image", ox.Short("m")).
				String("platform", "Set platform if server is multi-platform capable"),
		), ox.Sub(
			ox.Usage("inspect", "Return low-level information on Docker objects"),
			ox.Spec("[OPTIONS] NAME|ID [NAME|ID...]"),
			ox.Section(3),
			ox.Flags().
				String("format", "Format output using a custom template:", ox.Short("f")).
				Bool("size", "Display total file sizes if the type is container", ox.Short("s")).
				String("type", "Return JSON for specified type"),
		), ox.Sub(
			ox.Usage("kill", "Kill one or more running containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container kill", "docker kill"),
			ox.Section(3),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s")),
		), ox.Sub(
			ox.Usage("load", "Load an image from a tar archive or STDIN"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("docker image load", "docker load"),
			ox.Section(3),
			ox.Flags().
				String("input", "Read from tar archive file, instead of STDIN", ox.Short("i")).
				Bool("quiet", "Suppress the load output", ox.Short("q")),
		), ox.Sub(
			ox.Usage("logs", "Fetch the logs of a container"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("docker container logs", "docker logs"),
			ox.Section(3),
			ox.Flags().
				Bool("details", "Show extra details provided to logs").
				Bool("follow", "Follow log output", ox.Short("f")).
				String("since", "Show logs since timestamp (e.g.").
				String("tail", "Number of lines to show from the end of the logs", ox.Short("n")).
				Bool("timestamps", "Show timestamps", ox.Short("t")).
				String("until", "Show logs before a timestamp (e.g."),
		), ox.Sub(
			ox.Usage("pause", "Pause all processes within one or more containers"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container pause", "docker pause"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("port", "List port mappings or a specific mapping for the container"),
			ox.Spec("CONTAINER [PRIVATE_PORT[/PROTO]]"),
			ox.Aliases("docker container port", "docker port"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("rename", "Rename a container"),
			ox.Spec("CONTAINER NEW_NAME"),
			ox.Aliases("docker container rename", "docker rename"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("restart", "Restart one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container restart", "docker restart"),
			ox.Section(3),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s")).
				Int("time", "Seconds to wait before killing the container", ox.Short("t")),
		), ox.Sub(
			ox.Usage("rm", "Remove one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container rm", "docker container remove", "docker rm"),
			ox.Section(3),
			ox.Flags().
				Bool("force", "Force the removal of a running container (uses SIGKILL)", ox.Short("f")).
				Bool("link", "Remove the specified link", ox.Short("l")).
				Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v")),
		), ox.Sub(
			ox.Usage("rmi", "Remove one or more images"),
			ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
			ox.Aliases("docker image rm", "docker image remove", "docker rmi"),
			ox.Section(3),
			ox.Flags().
				Bool("force", "Force removal of the image", ox.Short("f")).
				Bool("no-prune", "Do not delete untagged parents"),
		), ox.Sub(
			ox.Usage("save", "Save one or more images to a tar archive (streamed to STDOUT by default)"),
			ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
			ox.Aliases("docker image save", "docker save"),
			ox.Section(3),
			ox.Flags().
				String("output", "Write to a file, instead of STDOUT", ox.Short("o")),
		), ox.Sub(
			ox.Usage("start", "Start one or more stopped containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container start", "docker start"),
			ox.Section(3),
			ox.Flags().
				Bool("attach", "Attach STDOUT/STDERR and forward signals", ox.Short("a")).
				String("checkpoint", "Restore from this checkpoint").
				String("checkpoint-dir", "Use a custom checkpoint storage directory").
				String("detach-keys", "Override the key sequence for detaching a").
				Bool("interactive", "Attach container's STDIN", ox.Short("i")),
		), ox.Sub(
			ox.Usage("stats", "Display a live stream of container(s) resource usage statistics"),
			ox.Spec("[OPTIONS] [CONTAINER...]"),
			ox.Aliases("docker container stats", "docker stats"),
			ox.Section(3),
			ox.Flags().
				Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a")).
				String("format", "Format output using a custom template:").
				Bool("no-stream", "Disable streaming stats and only pull the first result").
				Bool("no-trunc", "Do not truncate output"),
		), ox.Sub(
			ox.Usage("stop", "Stop one or more running containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container stop", "docker stop"),
			ox.Section(3),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s")).
				Int("time", "Seconds to wait before killing the container", ox.Short("t")),
		), ox.Sub(
			ox.Usage("tag", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
			ox.Spec("SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]"),
			ox.Aliases("docker image tag", "docker tag"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("top", "Display the running processes of a container"),
			ox.Spec("CONTAINER [ps OPTIONS]"),
			ox.Aliases("docker container top", "docker top"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("unpause", "Unpause all processes within one or more containers"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container unpause", "docker unpause"),
			ox.Section(3),
		), ox.Sub(
			ox.Usage("update", "Update configuration of one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container update", "docker update"),
			ox.Section(3),
			ox.Flags().
				Uint16("blkio-weight", "Block IO (relative weight), between 10").
				Int("cpu-period", "Limit CPU CFS (Completely Fair").
				Int("cpu-quota", "Limit CPU CFS (Completely Fair").
				Int("cpu-rt-period", "Limit the CPU real-time period in").
				Int("cpu-rt-runtime", "Limit the CPU real-time runtime in").
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)").
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m")).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT)).
				Slice("memory-swap", "Swap limit equal to memory plus swap:", ox.Elem(ox.UintT)).
				Int("pids-limit", "Tune container pids limit (set -1 for").
				String("restart", "Restart policy to apply when a"),
		), ox.Sub(
			ox.Usage("wait", "Block until one or more containers stop, then print their exit codes"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("docker container wait", "docker wait"),
			ox.Section(3),
		), ox.Flags().
			String("config", "Location of client config files (default").
			String("context", "Name of the context to use to connect to the", ox.Short("c")).
			Bool("debug", "Enable debug mode", ox.Short("D")).
			Slice("host", "Daemon socket to connect to", ox.Elem(ox.StringT), ox.Short("H")).
			String("log-level", "Set the logging level (\"debug\", \"info\",", ox.Short("l")).
			Bool("tls", "Use TLS; implied by --tlsverify").
			String("tlscacert", "Trust certs signed only by this CA (default").
			String("tlscert", "Path to TLS certificate file (default").
			String("tlskey", "Path to TLS key file (default").
			Bool("tlsverify", "Use TLS and verify the remote"),
	)
}
