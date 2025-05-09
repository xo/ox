// Command docker is a xo/ox version of `docker`.
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
		ox.Banner("A self-sufficient runtime for containers"),
		ox.Spec("[OPTIONS] COMMAND"),
		ox.Sections("Common Commands", "Management Commands", "Swarm Commands", "Commands"),
		ox.Help(ox.Sections(
			"Global Options",
		)),
		ox.Footer("Run 'docker COMMAND --help' for more information on a command.\n\nFor more help on how to use Docker, head to https://docs.docker.com/go/guides/"),
		ox.Sub(
			ox.Usage("run", "Create and run a new container from an image"),
			ox.Banner("Create and run a new container from an image"),
			ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
			ox.Aliases("container run"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping", ox.Section(0)).
				Map("annotation", "Add an annotation to the", ox.Section(0)).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Short("a"), ox.Section(0)).
				Uint16("blkio-weight", "Block IO (relative weight),", ox.Section(0)).
				Slice("blkio-weight-device", "Block IO weight (relative device", ox.Section(0)).
				Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
				Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
				String("cgroup-parent", "Optional parent cgroup for the", ox.Section(0)).
				String("cgroupns", "Cgroup namespace to use", ox.Section(0)).
				String("cidfile", "Write the container ID to the file", ox.Section(0)).
				Int("cpu-count", "CPU count (Windows only)", ox.Section(0)).
				Int("cpu-percent", "CPU percent (Windows only)", ox.Section(0)).
				Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-rt-period", "Limit CPU real-time period in", ox.Section(0)).
				Int("cpu-rt-runtime", "Limit CPU real-time runtime in", ox.Section(0)).
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
				String("cpuset-cpus", "CPUs in which to allow execution", ox.Section(0)).
				String("cpuset-mems", "MEMs in which to allow execution", ox.Section(0)).
				Bool("detach", "Run container in background and", ox.Short("d"), ox.Section(0)).
				String("detach-keys", "Override the key sequence for", ox.Section(0)).
				Slice("device", "Add a host device to the container", ox.Section(0)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Section(0)).
				Slice("device-read-bps", "Limit read rate (bytes per", ox.Section(0)).
				Slice("device-read-iops", "Limit read rate (IO per second)", ox.Section(0)).
				Slice("device-write-bps", "Limit write rate (bytes per", ox.Section(0)).
				Slice("device-write-iops", "Limit write rate (IO per second)", ox.Section(0)).
				Bool("disable-content-trust", "Skip image verification (default", ox.Section(0)).
				Slice("dns", "Set custom DNS servers", ox.Section(0)).
				Slice("dns-option", "Set DNS options", ox.Section(0)).
				Slice("dns-search", "Set custom DNS search domains", ox.Section(0)).
				String("domainname", "Container NIS domain name", ox.Section(0)).
				String("entrypoint", "Overwrite the default ENTRYPOINT", ox.Section(0)).
				Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
				Slice("env-file", "Read in a file of environment", ox.Section(0)).
				Slice("expose", "Expose a port or a range of ports", ox.Section(0)).
				String("gpus", "GPU devices to add to the", ox.Spec("gpu-request"), ox.Section(0)).
				Slice("group-add", "Add additional groups to join", ox.Section(0)).
				String("health-cmd", "Command to run to check health", ox.Section(0)).
				Duration("health-interval", "Time between running the check", ox.Section(0)).
				Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
				Duration("health-start-interval", "Time between running the check", ox.Section(0)).
				Duration("health-start-period", "Start period for the container", ox.Section(0)).
				Duration("health-timeout", "Maximum time to allow one check", ox.Section(0)).
				String("hostname", "Container host name", ox.Short("h"), ox.Section(0)).
				Bool("init", "Run an init inside the container", ox.Section(0)).
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
				Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT), ox.Section(0)).
				Uint("io-maxiops", "Maximum IOps limit for the", ox.Section(0)).
				String("ip", "IPv4 address (e.g., 172.30.100.104)", ox.Section(0)).
				String("ip6", "IPv6 address (e.g., 2001:db8::33)", ox.Section(0)).
				String("ipc", "IPC mode to use", ox.Section(0)).
				String("isolation", "Container isolation technology", ox.Section(0)).
				Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT), ox.Section(0)).
				Slice("label", "Set meta data on a container", ox.Short("l"), ox.Section(0)).
				Slice("label-file", "Read in a line delimited file of", ox.Section(0)).
				Slice("link", "Add link to another container", ox.Section(0)).
				Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Section(0)).
				String("log-driver", "Logging driver for the container", ox.Section(0)).
				Slice("log-opt", "Log driver options", ox.Section(0)).
				String("mac-address", "Container MAC address (e.g.,", ox.Section(0)).
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
				Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT), ox.Section(0)).
				Int("memory-swappiness", "Tune container memory swappiness", ox.Section(0)).
				String("mount", "Attach a filesystem mount to the", ox.Spec("mount"), ox.Section(0)).
				String("name", "Assign a name to the container", ox.Section(0)).
				String("network", "Connect a container to a network", ox.Spec("network"), ox.Section(0)).
				Slice("network-alias", "Add network-scoped alias for the", ox.Section(0)).
				Bool("no-healthcheck", "Disable any container-specified", ox.Section(0)).
				Bool("oom-kill-disable", "Disable OOM Killer", ox.Section(0)).
				Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
				String("pid", "PID namespace to use", ox.Section(0)).
				Int("pids-limit", "Tune container pids limit (set", ox.Section(0)).
				String("platform", "Set platform if server is", ox.Section(0)).
				Bool("privileged", "Give extended privileges to this", ox.Section(0)).
				Slice("publish", "Publish a container's port(s) to", ox.Short("p"), ox.Section(0)).
				Bool("publish-all", "Publish all exposed ports to", ox.Short("P"), ox.Section(0)).
				String("pull", "Pull image before running", ox.Section(0)).
				Bool("quiet", "Suppress the pull output", ox.Short("q"), ox.Section(0)).
				Bool("read-only", "Mount the container's root", ox.Section(0)).
				String("restart", "Restart policy to apply when a", ox.Section(0)).
				Bool("rm", "Automatically remove the", ox.Section(0)).
				String("runtime", "Runtime to use for this container", ox.Section(0)).
				Slice("security-opt", "Security Options", ox.Section(0)).
				Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT), ox.Section(0)).
				Bool("sig-proxy", "Proxy received signals to the", ox.Section(0)).
				String("stop-signal", "Signal to stop the container", ox.Section(0)).
				Int("stop-timeout", "Timeout (in seconds) to stop a", ox.Section(0)).
				Slice("storage-opt", "Storage driver options for the", ox.Section(0)).
				Map("sysctl", "Sysctl options", ox.Default("map[]"), ox.Section(0)).
				Slice("tmpfs", "Mount a tmpfs directory", ox.Section(0)).
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)).
				Bool("use-api-socket", "Bind mount Docker API socket and", ox.Section(0)).
				String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
				String("userns", "User namespace to use", ox.Section(0)).
				String("uts", "UTS namespace to use", ox.Section(0)).
				Slice("volume", "Bind mount a volume", ox.Short("v"), ox.Section(0)).
				String("volume-driver", "Optional volume driver for the", ox.Section(0)).
				Slice("volumes-from", "Mount volumes from the specified", ox.Section(0)).
				String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("exec", "Execute a command in a running container"),
			ox.Banner("Execute a command in a running container"),
			ox.Spec("[OPTIONS] CONTAINER COMMAND [ARG...]"),
			ox.Aliases("container exec"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("detach", "Detached mode: run command in the background", ox.Short("d"), ox.Section(0)).
				String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
				Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
				Slice("env-file", "Read in a file of environment variables", ox.Section(0)).
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
				Bool("privileged", "Give extended privileges to the command", ox.Section(0)).
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
				String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
				String("workdir", "Working directory inside the container", ox.Short("w"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("ps", "List containers"),
			ox.Banner("List containers"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("container ls", "container list", "container ps"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a"), ox.Section(0)).
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
				String("format", "Format output using a custom template:", ox.Section(0)).
				Int("last", "Show n last created containers (includes all", ox.Short("n"), ox.Section(0)).
				Bool("latest", "Show the latest created container (includes all", ox.Short("l"), ox.Section(0)).
				Bool("no-trunc", "Don't truncate output", ox.Section(0)).
				Bool("quiet", "Only display container IDs", ox.Short("q"), ox.Section(0)).
				Bool("size", "Display total file sizes", ox.Short("s"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("build", "Build an image from a Dockerfile"),
			ox.Banner("Build an image from a Dockerfile"),
			ox.Spec("[OPTIONS] PATH | URL | -"),
			ox.Aliases("image build", "builder build"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Section(0)).
				Slice("build-arg", "Set build-time variables", ox.Section(0)).
				Slice("cache-from", "Images to consider as cache sources", ox.Section(0)).
				String("cgroup-parent", "Set the parent cgroup for the \"RUN\"", ox.Section(0)).
				Bool("compress", "Compress the build context using gzip", ox.Section(0)).
				Int("cpu-period", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-quota", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)", ox.Section(0)).
				String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)", ox.Section(0)).
				Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
				String("file", "Name of the Dockerfile (Default is", ox.Short("f"), ox.Section(0)).
				Bool("force-rm", "Always remove intermediate containers", ox.Section(0)).
				String("iidfile", "Write the image ID to the file", ox.Section(0)).
				String("isolation", "Container isolation technology", ox.Section(0)).
				Slice("label", "Set metadata for an image", ox.Section(0)).
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
				Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT), ox.Section(0)).
				String("network", "Set the networking mode for the RUN", ox.Section(0)).
				Bool("no-cache", "Do not use cache when building the image", ox.Section(0)).
				String("platform", "Set platform if server is multi-platform", ox.Section(0)).
				Bool("pull", "Always attempt to pull a newer version of", ox.Section(0)).
				Bool("quiet", "Suppress the build output and print image", ox.Short("q"), ox.Section(0)).
				Bool("rm", "Remove intermediate containers after a", ox.Section(0)).
				Slice("security-opt", "Security options", ox.Section(0)).
				Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT), ox.Section(0)).
				Bool("squash", "Squash newly built layers into a single", ox.Section(0)).
				Slice("tag", "Name and optionally a tag in the", ox.Short("t"), ox.Section(0)).
				String("target", "Set the target build stage to build.", ox.Section(0)).
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("pull", "Download an image from a registry"),
			ox.Banner("Download an image from a registry"),
			ox.Spec("[OPTIONS] NAME[:TAG|@DIGEST]"),
			ox.Aliases("image pull"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("all-tags", "Download all tagged images in the repository", ox.Short("a"), ox.Section(0)).
				Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
				String("platform", "Set platform if server is multi-platform", ox.Section(0)).
				Bool("quiet", "Suppress verbose output", ox.Short("q"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("push", "Upload an image to a registry"),
			ox.Banner("Upload an image to a registry"),
			ox.Spec("[OPTIONS] NAME[:TAG]"),
			ox.Aliases("image push"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("all-tags", "Push all tags of an image to the repository", ox.Short("a"), ox.Section(0)).
				Bool("disable-content-trust", "Skip image signing", ox.Default("true"), ox.Section(0)).
				String("platform", "Push a platform-specific manifest as a", ox.Section(0)).
				Bool("quiet", "Suppress verbose output", ox.Short("q"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("images", "List images"),
			ox.Banner("List images"),
			ox.Spec("[OPTIONS] [REPOSITORY[:TAG]]"),
			ox.Aliases("image ls", "image list"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a"), ox.Section(0)).
				Bool("digests", "Show digests", ox.Section(0)).
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
				String("format", "Format output using a custom template:", ox.Section(0)).
				Bool("no-trunc", "Don't truncate output", ox.Section(0)).
				Bool("quiet", "Only show image IDs", ox.Short("q"), ox.Section(0)).
				Bool("tree", "List multi-platform images as a tree (EXPERIMENTAL)", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("login", "Authenticate to a registry"),
			ox.Banner("Authenticate to a registry.\nDefaults to Docker Hub if no server is specified."),
			ox.Spec("[OPTIONS] [SERVER]"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("password", "Password or Personal Access Token (PAT)", ox.Short("p"), ox.Section(0)).
				Bool("password-stdin", "Take the Password or Personal Access Token", ox.Section(0)).
				String("username", "Username", ox.Short("u"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("logout", "Log out from a registry"),
			ox.Spec("[SERVER]"),
			ox.Section(0),
			ox.Footer("Log out from a registry.\nIf no server is specified, the default is defined by the daemon."),
		),
		ox.Sub(
			ox.Usage("search", "Search Docker Hub for images"),
			ox.Banner("Search Docker Hub for images"),
			ox.Spec("[OPTIONS] TERM"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
				String("format", "Pretty-print search using a Go template", ox.Section(0)).
				Int("limit", "Max number of search results", ox.Section(0)).
				Bool("no-trunc", "Don't truncate output", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("info", "Display system-wide information"),
			ox.Banner("Display system-wide information"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("system info"),
			ox.Section(0),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("builder", "Manage builds"),
			ox.Banner("Manage builds"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker builder COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("build", "Build an image from a Dockerfile"),
				ox.Banner("Build an image from a Dockerfile"),
				ox.Spec("[OPTIONS] PATH | URL | -"),
				ox.Aliases("image build", "builder build"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Section(0)).
					Slice("build-arg", "Set build-time variables", ox.Section(0)).
					Slice("cache-from", "Images to consider as cache sources", ox.Section(0)).
					String("cgroup-parent", "Set the parent cgroup for the \"RUN\"", ox.Section(0)).
					Bool("compress", "Compress the build context using gzip", ox.Section(0)).
					Int("cpu-period", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
					String("file", "Name of the Dockerfile (Default is", ox.Short("f"), ox.Section(0)).
					Bool("force-rm", "Always remove intermediate containers", ox.Section(0)).
					String("iidfile", "Write the image ID to the file", ox.Section(0)).
					String("isolation", "Container isolation technology", ox.Section(0)).
					Slice("label", "Set metadata for an image", ox.Section(0)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
					Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT), ox.Section(0)).
					String("network", "Set the networking mode for the RUN", ox.Section(0)).
					Bool("no-cache", "Do not use cache when building the image", ox.Section(0)).
					String("platform", "Set platform if server is multi-platform", ox.Section(0)).
					Bool("pull", "Always attempt to pull a newer version of", ox.Section(0)).
					Bool("quiet", "Suppress the build output and print image", ox.Short("q"), ox.Section(0)).
					Bool("rm", "Remove intermediate containers after a", ox.Section(0)).
					Slice("security-opt", "Security options", ox.Section(0)).
					Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT), ox.Section(0)).
					Bool("squash", "Squash newly built layers into a single", ox.Section(0)).
					Slice("tag", "Name and optionally a tag in the", ox.Short("t"), ox.Section(0)).
					String("target", "Set the target build stage to build.", ox.Section(0)).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("prune", "Remove build cache"),
				ox.Banner("Remove build cache"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Remove all unused build cache, not just", ox.Short("a"), ox.Section(0)).
					String("filter", "Provide filter values (e.g. \"until=24h\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)).
					Slice("keep-storage", "Amount of disk space to keep for cache", ox.Elem(ox.UintT), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("checkpoint", "Manage checkpoints"),
			ox.Banner("Manage checkpoints"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker checkpoint COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a checkpoint from a running container"),
				ox.Banner("Create a checkpoint from a running container"),
				ox.Spec("[OPTIONS] CONTAINER CHECKPOINT"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory", ox.Section(0)).
					Bool("leave-running", "Leave the container running after checkpoint", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List checkpoints for a container"),
				ox.Banner("List checkpoints for a container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("checkpoint ls", "checkpoint list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove a checkpoint"),
				ox.Banner("Remove a checkpoint"),
				ox.Spec("[OPTIONS] CONTAINER CHECKPOINT"),
				ox.Aliases("checkpoint rm", "checkpoint remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("checkpoint-dir", "Use a custom checkpoint storage directory", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("container", "Manage containers"),
			ox.Banner("Manage containers"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker container COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("attach", "Attach local standard input, output, and error streams to a running container"),
				ox.Banner("Attach local standard input, output, and error streams to a running container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("container attach"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
					Bool("no-stdin", "Do not attach STDIN", ox.Section(0)).
					Bool("sig-proxy", "Proxy all received signals to the process", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("commit", "Create a new image from a container's changes"),
				ox.Banner("Create a new image from a container's changes"),
				ox.Spec("[OPTIONS] CONTAINER [REPOSITORY[:TAG]]"),
				ox.Aliases("container commit"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("author", "Author (e.g., \"John Hannibal Smith", ox.Short("a"), ox.Section(0)).
					Slice("change", "Apply Dockerfile instruction to the created image", ox.Short("c"), ox.Section(0)).
					String("message", "Commit message", ox.Short("m"), ox.Section(0)).
					Bool("pause", "Pause container during commit", ox.Default("true"), ox.Short("p"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
				ox.Banner("docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH\n\nCopy files/folders between a container and the local filesystem\n\nUse '-' as the source to read a tar archive from stdin\nand extract it to a directory destination in a container.\nUse '-' as the destination to stream a tar archive of a\ncontainer source to stdout."),
				ox.Spec("[OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-"),
				ox.Aliases("container cp"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("archive", "Archive mode (copy all uid/gid information)", ox.Short("a"), ox.Section(0)).
					Bool("follow-link", "Always follow symbol link in SRC_PATH", ox.Short("L"), ox.Section(0)).
					Bool("quiet", "Suppress progress output during copy. Progress", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("create", "Create a new container"),
				ox.Banner("Create a new container"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Aliases("container create"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping", ox.Section(0)).
					Map("annotation", "Add an annotation to the", ox.Section(0)).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Short("a"), ox.Section(0)).
					Uint16("blkio-weight", "Block IO (relative weight),", ox.Section(0)).
					Slice("blkio-weight-device", "Block IO weight (relative device", ox.Section(0)).
					Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
					String("cgroup-parent", "Optional parent cgroup for the", ox.Section(0)).
					String("cgroupns", "Cgroup namespace to use", ox.Section(0)).
					String("cidfile", "Write the container ID to the file", ox.Section(0)).
					Int("cpu-count", "CPU count (Windows only)", ox.Section(0)).
					Int("cpu-percent", "CPU percent (Windows only)", ox.Section(0)).
					Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-rt-period", "Limit CPU real-time period in", ox.Section(0)).
					Int("cpu-rt-runtime", "Limit CPU real-time runtime in", ox.Section(0)).
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
					String("cpuset-cpus", "CPUs in which to allow execution", ox.Section(0)).
					String("cpuset-mems", "MEMs in which to allow execution", ox.Section(0)).
					Slice("device", "Add a host device to the container", ox.Section(0)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Section(0)).
					Slice("device-read-bps", "Limit read rate (bytes per", ox.Section(0)).
					Slice("device-read-iops", "Limit read rate (IO per second)", ox.Section(0)).
					Slice("device-write-bps", "Limit write rate (bytes per", ox.Section(0)).
					Slice("device-write-iops", "Limit write rate (IO per second)", ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification (default", ox.Section(0)).
					Slice("dns", "Set custom DNS servers", ox.Section(0)).
					Slice("dns-option", "Set DNS options", ox.Section(0)).
					Slice("dns-search", "Set custom DNS search domains", ox.Section(0)).
					String("domainname", "Container NIS domain name", ox.Section(0)).
					String("entrypoint", "Overwrite the default ENTRYPOINT", ox.Section(0)).
					Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
					Slice("env-file", "Read in a file of environment", ox.Section(0)).
					Slice("expose", "Expose a port or a range of ports", ox.Section(0)).
					String("gpus", "GPU devices to add to the", ox.Spec("gpu-request"), ox.Section(0)).
					Slice("group-add", "Add additional groups to join", ox.Section(0)).
					String("health-cmd", "Command to run to check health", ox.Section(0)).
					Duration("health-interval", "Time between running the check", ox.Section(0)).
					Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
					Duration("health-start-interval", "Time between running the check", ox.Section(0)).
					Duration("health-start-period", "Start period for the container", ox.Section(0)).
					Duration("health-timeout", "Maximum time to allow one check", ox.Section(0)).
					String("hostname", "Container host name", ox.Short("h"), ox.Section(0)).
					Bool("init", "Run an init inside the container", ox.Section(0)).
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
					Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT), ox.Section(0)).
					Uint("io-maxiops", "Maximum IOps limit for the", ox.Section(0)).
					String("ip", "IPv4 address (e.g., 172.30.100.104)", ox.Section(0)).
					String("ip6", "IPv6 address (e.g., 2001:db8::33)", ox.Section(0)).
					String("ipc", "IPC mode to use", ox.Section(0)).
					String("isolation", "Container isolation technology", ox.Section(0)).
					Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT), ox.Section(0)).
					Slice("label", "Set meta data on a container", ox.Short("l"), ox.Section(0)).
					Slice("label-file", "Read in a line delimited file of", ox.Section(0)).
					Slice("link", "Add link to another container", ox.Section(0)).
					Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Section(0)).
					String("log-driver", "Logging driver for the container", ox.Section(0)).
					Slice("log-opt", "Log driver options", ox.Section(0)).
					String("mac-address", "Container MAC address (e.g.,", ox.Section(0)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
					Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT), ox.Section(0)).
					Int("memory-swappiness", "Tune container memory swappiness", ox.Section(0)).
					String("mount", "Attach a filesystem mount to the", ox.Spec("mount"), ox.Section(0)).
					String("name", "Assign a name to the container", ox.Section(0)).
					String("network", "Connect a container to a network", ox.Spec("network"), ox.Section(0)).
					Slice("network-alias", "Add network-scoped alias for the", ox.Section(0)).
					Bool("no-healthcheck", "Disable any container-specified", ox.Section(0)).
					Bool("oom-kill-disable", "Disable OOM Killer", ox.Section(0)).
					Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
					String("pid", "PID namespace to use", ox.Section(0)).
					Int("pids-limit", "Tune container pids limit (set", ox.Section(0)).
					String("platform", "Set platform if server is", ox.Section(0)).
					Bool("privileged", "Give extended privileges to this", ox.Section(0)).
					Slice("publish", "Publish a container's port(s) to", ox.Short("p"), ox.Section(0)).
					Bool("publish-all", "Publish all exposed ports to", ox.Short("P"), ox.Section(0)).
					String("pull", "Pull image before creating", ox.Section(0)).
					Bool("quiet", "Suppress the pull output", ox.Short("q"), ox.Section(0)).
					Bool("read-only", "Mount the container's root", ox.Section(0)).
					String("restart", "Restart policy to apply when a", ox.Section(0)).
					Bool("rm", "Automatically remove the", ox.Section(0)).
					String("runtime", "Runtime to use for this container", ox.Section(0)).
					Slice("security-opt", "Security Options", ox.Section(0)).
					Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT), ox.Section(0)).
					String("stop-signal", "Signal to stop the container", ox.Section(0)).
					Int("stop-timeout", "Timeout (in seconds) to stop a", ox.Section(0)).
					Slice("storage-opt", "Storage driver options for the", ox.Section(0)).
					Map("sysctl", "Sysctl options", ox.Default("map[]"), ox.Section(0)).
					Slice("tmpfs", "Mount a tmpfs directory", ox.Section(0)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)).
					Bool("use-api-socket", "Bind mount Docker API socket and", ox.Section(0)).
					String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
					String("userns", "User namespace to use", ox.Section(0)).
					String("uts", "UTS namespace to use", ox.Section(0)).
					Slice("volume", "Bind mount a volume", ox.Short("v"), ox.Section(0)).
					String("volume-driver", "Optional volume driver for the", ox.Section(0)).
					Slice("volumes-from", "Mount volumes from the specified", ox.Section(0)).
					String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("diff", "Inspect changes to files or directories on a container's filesystem"),
				ox.Banner("Inspect changes to files or directories on a container's filesystem"),
				ox.Spec("CONTAINER"),
				ox.Aliases("container diff"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("exec", "Execute a command in a running container"),
				ox.Banner("Execute a command in a running container"),
				ox.Spec("[OPTIONS] CONTAINER COMMAND [ARG...]"),
				ox.Aliases("container exec"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("detach", "Detached mode: run command in the background", ox.Short("d"), ox.Section(0)).
					String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
					Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
					Slice("env-file", "Read in a file of environment variables", ox.Section(0)).
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
					Bool("privileged", "Give extended privileges to the command", ox.Section(0)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
					String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
					String("workdir", "Working directory inside the container", ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("export", "Export a container's filesystem as a tar archive"),
				ox.Banner("Export a container's filesystem as a tar archive"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("container export"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("output", "Write to a file, instead of STDOUT", ox.Short("o"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more containers"),
				ox.Banner("Display detailed information on one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("size", "Display total file sizes", ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("kill", "Kill one or more running containers"),
				ox.Banner("Kill one or more running containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container kill"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("logs", "Fetch the logs of a container"),
				ox.Banner("Fetch the logs of a container"),
				ox.Spec("[OPTIONS] CONTAINER"),
				ox.Aliases("container logs"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("details", "Show extra details provided to logs", ox.Section(0)).
					Bool("follow", "Follow log output", ox.Short("f"), ox.Section(0)).
					String("since", "Show logs since timestamp (e.g.", ox.Section(0)).
					String("tail", "Number of lines to show from the end of the logs", ox.Short("n"), ox.Section(0)).
					Bool("timestamps", "Show timestamps", ox.Short("t"), ox.Section(0)).
					String("until", "Show logs before a timestamp (e.g.", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List containers"),
				ox.Banner("List containers"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("container ls", "container list", "container ps", "ps"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a"), ox.Section(0)).
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Int("last", "Show n last created containers (includes all", ox.Short("n"), ox.Section(0)).
					Bool("latest", "Show the latest created container (includes all", ox.Short("l"), ox.Section(0)).
					Bool("no-trunc", "Don't truncate output", ox.Section(0)).
					Bool("quiet", "Only display container IDs", ox.Short("q"), ox.Section(0)).
					Bool("size", "Display total file sizes", ox.Short("s"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("pause", "Pause all processes within one or more containers"),
				ox.Banner("Pause all processes within one or more containers"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("container pause"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("port", "List port mappings or a specific mapping for the container"),
				ox.Banner("List port mappings or a specific mapping for the container"),
				ox.Spec("CONTAINER [PRIVATE_PORT[/PROTO]]"),
				ox.Aliases("container port"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("prune", "Remove all stopped containers"),
				ox.Banner("Remove all stopped containers"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rename", "Rename a container"),
				ox.Banner("Rename a container"),
				ox.Spec("CONTAINER NEW_NAME"),
				ox.Aliases("container rename"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("restart", "Restart one or more containers"),
				ox.Banner("Restart one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container restart"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)).
					Int("timeout", "Seconds to wait before killing the container", ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more containers"),
				ox.Banner("Remove one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container rm", "container remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the removal of a running container (uses SIGKILL)", ox.Short("f"), ox.Section(0)).
					Bool("link", "Remove the specified link", ox.Short("l"), ox.Section(0)).
					Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("run", "Create and run a new container from an image"),
				ox.Banner("Create and run a new container from an image"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Aliases("container run"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping", ox.Section(0)).
					Map("annotation", "Add an annotation to the", ox.Section(0)).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Short("a"), ox.Section(0)).
					Uint16("blkio-weight", "Block IO (relative weight),", ox.Section(0)).
					Slice("blkio-weight-device", "Block IO weight (relative device", ox.Section(0)).
					Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
					String("cgroup-parent", "Optional parent cgroup for the", ox.Section(0)).
					String("cgroupns", "Cgroup namespace to use", ox.Section(0)).
					String("cidfile", "Write the container ID to the file", ox.Section(0)).
					Int("cpu-count", "CPU count (Windows only)", ox.Section(0)).
					Int("cpu-percent", "CPU percent (Windows only)", ox.Section(0)).
					Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-rt-period", "Limit CPU real-time period in", ox.Section(0)).
					Int("cpu-rt-runtime", "Limit CPU real-time runtime in", ox.Section(0)).
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
					String("cpuset-cpus", "CPUs in which to allow execution", ox.Section(0)).
					String("cpuset-mems", "MEMs in which to allow execution", ox.Section(0)).
					Bool("detach", "Run container in background and", ox.Short("d"), ox.Section(0)).
					String("detach-keys", "Override the key sequence for", ox.Section(0)).
					Slice("device", "Add a host device to the container", ox.Section(0)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Section(0)).
					Slice("device-read-bps", "Limit read rate (bytes per", ox.Section(0)).
					Slice("device-read-iops", "Limit read rate (IO per second)", ox.Section(0)).
					Slice("device-write-bps", "Limit write rate (bytes per", ox.Section(0)).
					Slice("device-write-iops", "Limit write rate (IO per second)", ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification (default", ox.Section(0)).
					Slice("dns", "Set custom DNS servers", ox.Section(0)).
					Slice("dns-option", "Set DNS options", ox.Section(0)).
					Slice("dns-search", "Set custom DNS search domains", ox.Section(0)).
					String("domainname", "Container NIS domain name", ox.Section(0)).
					String("entrypoint", "Overwrite the default ENTRYPOINT", ox.Section(0)).
					Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
					Slice("env-file", "Read in a file of environment", ox.Section(0)).
					Slice("expose", "Expose a port or a range of ports", ox.Section(0)).
					String("gpus", "GPU devices to add to the", ox.Spec("gpu-request"), ox.Section(0)).
					Slice("group-add", "Add additional groups to join", ox.Section(0)).
					String("health-cmd", "Command to run to check health", ox.Section(0)).
					Duration("health-interval", "Time between running the check", ox.Section(0)).
					Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
					Duration("health-start-interval", "Time between running the check", ox.Section(0)).
					Duration("health-start-period", "Start period for the container", ox.Section(0)).
					Duration("health-timeout", "Maximum time to allow one check", ox.Section(0)).
					String("hostname", "Container host name", ox.Short("h"), ox.Section(0)).
					Bool("init", "Run an init inside the container", ox.Section(0)).
					Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
					Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT), ox.Section(0)).
					Uint("io-maxiops", "Maximum IOps limit for the", ox.Section(0)).
					String("ip", "IPv4 address (e.g., 172.30.100.104)", ox.Section(0)).
					String("ip6", "IPv6 address (e.g., 2001:db8::33)", ox.Section(0)).
					String("ipc", "IPC mode to use", ox.Section(0)).
					String("isolation", "Container isolation technology", ox.Section(0)).
					Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT), ox.Section(0)).
					Slice("label", "Set meta data on a container", ox.Short("l"), ox.Section(0)).
					Slice("label-file", "Read in a line delimited file of", ox.Section(0)).
					Slice("link", "Add link to another container", ox.Section(0)).
					Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Section(0)).
					String("log-driver", "Logging driver for the container", ox.Section(0)).
					Slice("log-opt", "Log driver options", ox.Section(0)).
					String("mac-address", "Container MAC address (e.g.,", ox.Section(0)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
					Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT), ox.Section(0)).
					Int("memory-swappiness", "Tune container memory swappiness", ox.Section(0)).
					String("mount", "Attach a filesystem mount to the", ox.Spec("mount"), ox.Section(0)).
					String("name", "Assign a name to the container", ox.Section(0)).
					String("network", "Connect a container to a network", ox.Spec("network"), ox.Section(0)).
					Slice("network-alias", "Add network-scoped alias for the", ox.Section(0)).
					Bool("no-healthcheck", "Disable any container-specified", ox.Section(0)).
					Bool("oom-kill-disable", "Disable OOM Killer", ox.Section(0)).
					Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
					String("pid", "PID namespace to use", ox.Section(0)).
					Int("pids-limit", "Tune container pids limit (set", ox.Section(0)).
					String("platform", "Set platform if server is", ox.Section(0)).
					Bool("privileged", "Give extended privileges to this", ox.Section(0)).
					Slice("publish", "Publish a container's port(s) to", ox.Short("p"), ox.Section(0)).
					Bool("publish-all", "Publish all exposed ports to", ox.Short("P"), ox.Section(0)).
					String("pull", "Pull image before running", ox.Section(0)).
					Bool("quiet", "Suppress the pull output", ox.Short("q"), ox.Section(0)).
					Bool("read-only", "Mount the container's root", ox.Section(0)).
					String("restart", "Restart policy to apply when a", ox.Section(0)).
					Bool("rm", "Automatically remove the", ox.Section(0)).
					String("runtime", "Runtime to use for this container", ox.Section(0)).
					Slice("security-opt", "Security Options", ox.Section(0)).
					Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT), ox.Section(0)).
					Bool("sig-proxy", "Proxy received signals to the", ox.Section(0)).
					String("stop-signal", "Signal to stop the container", ox.Section(0)).
					Int("stop-timeout", "Timeout (in seconds) to stop a", ox.Section(0)).
					Slice("storage-opt", "Storage driver options for the", ox.Section(0)).
					Map("sysctl", "Sysctl options", ox.Default("map[]"), ox.Section(0)).
					Slice("tmpfs", "Mount a tmpfs directory", ox.Section(0)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)).
					Bool("use-api-socket", "Bind mount Docker API socket and", ox.Section(0)).
					String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
					String("userns", "User namespace to use", ox.Section(0)).
					String("uts", "UTS namespace to use", ox.Section(0)).
					Slice("volume", "Bind mount a volume", ox.Short("v"), ox.Section(0)).
					String("volume-driver", "Optional volume driver for the", ox.Section(0)).
					Slice("volumes-from", "Mount volumes from the specified", ox.Section(0)).
					String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("start", "Start one or more stopped containers"),
				ox.Banner("Start one or more stopped containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container start"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("attach", "Attach STDOUT/STDERR and forward signals", ox.Short("a"), ox.Section(0)).
					String("checkpoint", "Restore from this checkpoint", ox.Section(0)).
					String("checkpoint-dir", "Use a custom checkpoint storage directory", ox.Section(0)).
					String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
					Bool("interactive", "Attach container's STDIN", ox.Short("i"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("stats", "Display a live stream of container(s) resource usage statistics"),
				ox.Banner("Display a live stream of container(s) resource usage statistics"),
				ox.Spec("[OPTIONS] [CONTAINER...]"),
				ox.Aliases("container stats"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("no-stream", "Disable streaming stats and only pull the first result", ox.Section(0)).
					Bool("no-trunc", "Do not truncate output", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("stop", "Stop one or more running containers"),
				ox.Banner("Stop one or more running containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container stop"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)).
					Int("timeout", "Seconds to wait before killing the container", ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("top", "Display the running processes of a container"),
				ox.Banner("Display the running processes of a container"),
				ox.Spec("CONTAINER [ps OPTIONS]"),
				ox.Aliases("container top"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("unpause", "Unpause all processes within one or more containers"),
				ox.Banner("Unpause all processes within one or more containers"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("container unpause"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("update", "Update configuration of one or more containers"),
				ox.Banner("Update configuration of one or more containers"),
				ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
				ox.Aliases("container update"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Uint16("blkio-weight", "Block IO (relative weight), between 10", ox.Section(0)).
					Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-rt-period", "Limit the CPU real-time period in", ox.Section(0)).
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in", ox.Section(0)).
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
					Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
					Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
					Slice("memory-swap", "Swap limit equal to memory plus swap:", ox.Elem(ox.UintT), ox.Section(0)).
					Int("pids-limit", "Tune container pids limit (set -1 for", ox.Section(0)).
					String("restart", "Restart policy to apply when a", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("wait", "Block until one or more containers stop, then print their exit codes"),
				ox.Banner("Block until one or more containers stop, then print their exit codes"),
				ox.Spec("CONTAINER [CONTAINER...]"),
				ox.Aliases("container wait"),
				ox.Section(0),
			),
		),
		ox.Sub(
			ox.Usage("context", "Manage contexts"),
			ox.Banner("Manage contexts"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker context COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("export", "Export a context to a tar archive FILE or a tar stream on STDOUT."),
				ox.Spec("[OPTIONS] CONTEXT [FILE|-]"),
				ox.Section(0),
				ox.Footer("Export a context to a tar archive FILE or a tar stream on STDOUT."),
			),
			ox.Sub(
				ox.Usage("import", "Import a context from a tar or zip file"),
				ox.Spec("CONTEXT FILE|-"),
				ox.Section(0),
				ox.Footer("Import a context from a tar or zip file"),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more contexts"),
				ox.Banner("Display detailed information on one or more contexts"),
				ox.Spec("[OPTIONS] [CONTEXT] [CONTEXT...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List contexts"),
				ox.Banner("List contexts"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("context ls", "context list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only show context names", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more contexts"),
				ox.Banner("Remove one or more contexts"),
				ox.Spec("CONTEXT [CONTEXT...]"),
				ox.Aliases("context rm", "context remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the removal of a context in use", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("show", "Print the name of the current context"),
				ox.Section(0),
				ox.Footer("Print the name of the current context"),
			),
			ox.Sub(
				ox.Usage("update", "Update a context"),
				ox.Banner("Update a context"),
				ox.Spec("[OPTIONS] CONTEXT"),
				ox.Example("\n$ docker context update my-context --description \"some description\" --docker \"host=tcp://myserver:2376,ca=~/ca-file,cert=~/cert-file,key=~/key-file\""),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("description", "Description of the context", ox.Section(0)).
					Map("docker", "set the docker endpoint", ox.Default("[]"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("use", "Set the current docker context"),
				ox.Spec("CONTEXT"),
				ox.Section(0),
				ox.Footer("Set the current docker context"),
			),
		),
		ox.Sub(
			ox.Usage("image", "Manage images"),
			ox.Banner("Manage images"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker image COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("build", "Build an image from a Dockerfile"),
				ox.Banner("Build an image from a Dockerfile"),
				ox.Spec("[OPTIONS] PATH | URL | -"),
				ox.Aliases("image build", "builder build"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (\"host:ip\")", ox.Section(0)).
					Slice("build-arg", "Set build-time variables", ox.Section(0)).
					Slice("cache-from", "Images to consider as cache sources", ox.Section(0)).
					String("cgroup-parent", "Set the parent cgroup for the \"RUN\"", ox.Section(0)).
					Bool("compress", "Compress the build context using gzip", ox.Section(0)).
					Int("cpu-period", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair", ox.Section(0)).
					Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)", ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
					String("file", "Name of the Dockerfile (Default is", ox.Short("f"), ox.Section(0)).
					Bool("force-rm", "Always remove intermediate containers", ox.Section(0)).
					String("iidfile", "Write the image ID to the file", ox.Section(0)).
					String("isolation", "Container isolation technology", ox.Section(0)).
					Slice("label", "Set metadata for an image", ox.Section(0)).
					Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
					Slice("memory-swap", "Swap limit equal to memory plus swap: -1", ox.Elem(ox.UintT), ox.Section(0)).
					String("network", "Set the networking mode for the RUN", ox.Section(0)).
					Bool("no-cache", "Do not use cache when building the image", ox.Section(0)).
					String("platform", "Set platform if server is multi-platform", ox.Section(0)).
					Bool("pull", "Always attempt to pull a newer version of", ox.Section(0)).
					Bool("quiet", "Suppress the build output and print image", ox.Short("q"), ox.Section(0)).
					Bool("rm", "Remove intermediate containers after a", ox.Section(0)).
					Slice("security-opt", "Security options", ox.Section(0)).
					Slice("shm-size", "Size of \"/dev/shm\"", ox.Elem(ox.UintT), ox.Section(0)).
					Bool("squash", "Squash newly built layers into a single", ox.Section(0)).
					Slice("tag", "Name and optionally a tag in the", ox.Short("t"), ox.Section(0)).
					String("target", "Set the target build stage to build.", ox.Section(0)).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("history", "Show the history of an image"),
				ox.Banner("Show the history of an image"),
				ox.Spec("[OPTIONS] IMAGE"),
				ox.Aliases("image history"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("human", "Print sizes and dates in human readable format", ox.Short("H"), ox.Section(0)).
					Bool("no-trunc", "Don't truncate output", ox.Section(0)).
					String("platform", "Show history for the given platform. Formatted", ox.Section(0)).
					Bool("quiet", "Only show image IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("import", "Import the contents from a tarball to create a filesystem image"),
				ox.Banner("Import the contents from a tarball to create a filesystem image"),
				ox.Spec("[OPTIONS] file|URL|- [REPOSITORY[:TAG]]"),
				ox.Aliases("image import"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("change", "Apply Dockerfile instruction to the created image", ox.Short("c"), ox.Section(0)).
					String("message", "Set commit message for imported image", ox.Short("m"), ox.Section(0)).
					String("platform", "Set platform if server is multi-platform capable", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more images"),
				ox.Banner("Display detailed information on one or more images"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					String("platform", "Inspect a specific platform of the", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("load", "Load an image from a tar archive or STDIN"),
				ox.Banner("Load an image from a tar archive or STDIN"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("image load"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("input", "Read from tar archive file, instead of STDIN", ox.Short("i"), ox.Section(0)).
					String("platform", "Load only the given platform variant. Formatted", ox.Section(0)).
					Bool("quiet", "Suppress the load output", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List images"),
				ox.Banner("List images"),
				ox.Spec("[OPTIONS] [REPOSITORY[:TAG]]"),
				ox.Aliases("image ls", "image list", "images"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a"), ox.Section(0)).
					Bool("digests", "Show digests", ox.Section(0)).
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("no-trunc", "Don't truncate output", ox.Section(0)).
					Bool("quiet", "Only show image IDs", ox.Short("q"), ox.Section(0)).
					Bool("tree", "List multi-platform images as a tree (EXPERIMENTAL)", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("prune", "Remove unused images"),
				ox.Banner("Remove unused images"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Remove all unused images, not just dangling ones", ox.Short("a"), ox.Section(0)).
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("pull", "Download an image from a registry"),
				ox.Banner("Download an image from a registry"),
				ox.Spec("[OPTIONS] NAME[:TAG|@DIGEST]"),
				ox.Aliases("image pull"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all-tags", "Download all tagged images in the repository", ox.Short("a"), ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
					String("platform", "Set platform if server is multi-platform", ox.Section(0)).
					Bool("quiet", "Suppress verbose output", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("push", "Upload an image to a registry"),
				ox.Banner("Upload an image to a registry"),
				ox.Spec("[OPTIONS] NAME[:TAG]"),
				ox.Aliases("image push"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all-tags", "Push all tags of an image to the repository", ox.Short("a"), ox.Section(0)).
					Bool("disable-content-trust", "Skip image signing", ox.Default("true"), ox.Section(0)).
					String("platform", "Push a platform-specific manifest as a", ox.Section(0)).
					Bool("quiet", "Suppress verbose output", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more images"),
				ox.Banner("Remove one or more images"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Aliases("image rm", "image remove", "rmi"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force removal of the image", ox.Short("f"), ox.Section(0)).
					Bool("no-prune", "Do not delete untagged parents", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("save", "Save one or more images to a tar archive (streamed to STDOUT by default)"),
				ox.Banner("Save one or more images to a tar archive (streamed to STDOUT by default)"),
				ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
				ox.Aliases("image save"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("output", "Write to a file, instead of STDOUT", ox.Short("o"), ox.Section(0)).
					String("platform", "Save only the given platform variant. Formatted", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("tag", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
				ox.Banner("Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
				ox.Spec("SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]"),
				ox.Aliases("image tag"),
				ox.Section(0),
			),
		),
		ox.Sub(
			ox.Usage("manifest", "Manage Docker image manifests and manifest lists"),
			ox.Banner("The **docker manifest** command has subcommands for managing image manifests and\nmanifest lists. A manifest list allows you to use one name to refer to the same image\nbuilt for multiple architectures.\n\nTo see help for a subcommand, use:\n\n    docker manifest CMD --help\n\nFor full details on using docker manifest lists, see the registry v2 specification."),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker manifest COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("annotate", "Add additional information to a local image manifest"),
				ox.Banner("Add additional information to a local image manifest"),
				ox.Spec("[OPTIONS] MANIFEST_LIST MANIFEST"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("arch", "Set architecture", ox.Section(0)).
					String("os", "Set operating system", ox.Section(0)).
					Slice("os-features", "Set operating system feature", ox.Section(0)).
					String("os-version", "Set operating system version", ox.Section(0)).
					String("variant", "Set architecture variant", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("create", "Create a local manifest list for annotating and pushing to a registry"),
				ox.Banner("Create a local manifest list for annotating and pushing to a registry"),
				ox.Spec("MANIFEST_LIST MANIFEST [MANIFEST...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("amend", "Amend an existing manifest list", ox.Short("a"), ox.Section(0)).
					Bool("insecure", "Allow communication with an insecure registry", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display an image manifest, or manifest list"),
				ox.Banner("Display an image manifest, or manifest list"),
				ox.Spec("[OPTIONS] [MANIFEST_LIST] MANIFEST"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("insecure", "Allow communication with an insecure registry", ox.Section(0)).
					Bool("verbose", "Output additional info including layers and platform", ox.Short("v"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("push", "Push a manifest list to a repository"),
				ox.Banner("Push a manifest list to a repository"),
				ox.Spec("[OPTIONS] MANIFEST_LIST"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("insecure", "Allow push to an insecure registry", ox.Section(0)).
					Bool("purge", "Remove the local manifest list after push", ox.Short("p"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Delete one or more manifest lists from local storage"),
				ox.Banner("Delete one or more manifest lists from local storage"),
				ox.Spec("MANIFEST_LIST [MANIFEST_LIST...]"),
				ox.Section(0),
			),
		),
		ox.Sub(
			ox.Usage("network", "Manage networks"),
			ox.Banner("Manage networks"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker network COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("connect", "Connect a container to a network"),
				ox.Banner("Connect a container to a network"),
				ox.Spec("[OPTIONS] NETWORK CONTAINER"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("alias", "Add network-scoped alias for the container", ox.Section(0)).
					Slice("driver-opt", "driver options for the network", ox.Section(0)).
					Int("gw-priority", "Highest gw-priority provides the default", ox.Section(0)).
					String("ip", "IPv4 address (e.g., \"172.30.100.104\")", ox.Section(0)).
					String("ip6", "IPv6 address (e.g., \"2001:db8::33\")", ox.Section(0)).
					Slice("link", "Add link to another container", ox.Section(0)).
					Slice("link-local-ip", "Add a link-local address for the container", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("create", "Create a network"),
				ox.Banner("Create a network"),
				ox.Spec("[OPTIONS] NETWORK"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("attachable", "Enable manual container attachment", ox.Section(0)).
					Map("aux-address", "Auxiliary IPv4 or IPv6 addresses used by", ox.Section(0)).
					String("config-from", "The network from which to copy the configuration", ox.Section(0)).
					Bool("config-only", "Create a configuration only network", ox.Section(0)).
					String("driver", "Driver to manage the Network", ox.Default("bridge"), ox.Short("d"), ox.Section(0)).
					Slice("gateway", "IPv4 or IPv6 Gateway for the master subnet", ox.Section(0)).
					Bool("ingress", "Create swarm routing-mesh network", ox.Section(0)).
					Bool("internal", "Restrict external access to the network", ox.Section(0)).
					Slice("ip-range", "Allocate container ip from a sub-range", ox.Section(0)).
					String("ipam-driver", "IP Address Management Driver", ox.Default("default"), ox.Section(0)).
					Map("ipam-opt", "Set IPAM driver specific options", ox.Default("map[]"), ox.Section(0)).
					Bool("ipv4", "Enable or disable IPv4 address assignment", ox.Section(0)).
					Bool("ipv6", "Enable or disable IPv6 address assignment", ox.Section(0)).
					Slice("label", "Set metadata on a network", ox.Section(0)).
					Map("opt", "Set driver specific options", ox.Default("map[]"), ox.Short("o"), ox.Section(0)).
					String("scope", "Control the network's scope", ox.Section(0)).
					Slice("subnet", "Subnet in CIDR format that represents a", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("disconnect", "Disconnect a container from a network"),
				ox.Banner("Disconnect a container from a network"),
				ox.Spec("[OPTIONS] NETWORK CONTAINER"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the container to disconnect from a network", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more networks"),
				ox.Banner("Display detailed information on one or more networks"),
				ox.Spec("[OPTIONS] NETWORK [NETWORK...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("verbose", "Verbose output for diagnostics", ox.Short("v"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List networks"),
				ox.Banner("List networks"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("network ls", "network list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"driver=bridge\")", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("no-trunc", "Do not truncate the output", ox.Section(0)).
					Bool("quiet", "Only display network IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("prune", "Remove all unused networks"),
				ox.Banner("Remove all unused networks"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"until=<timestamp>\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more networks"),
				ox.Banner("Remove one or more networks"),
				ox.Spec("NETWORK [NETWORK...]"),
				ox.Aliases("network rm", "network remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Do not error if the network does not exist", ox.Short("f"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("plugin", "Manage plugins"),
			ox.Banner("Manage plugins"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker plugin COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory."),
				ox.Banner("Create a plugin from a rootfs and configuration. Plugin data directory must contain config.json and rootfs directory."),
				ox.Spec("[OPTIONS] PLUGIN PLUGIN-DATA-DIR"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("compress", "Compress the context using gzip", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("disable", "Disable a plugin"),
				ox.Banner("Disable a plugin"),
				ox.Spec("[OPTIONS] PLUGIN"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the disable of an active plugin", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("enable", "Enable a plugin"),
				ox.Banner("Enable a plugin"),
				ox.Spec("[OPTIONS] PLUGIN"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Int("timeout", "HTTP client timeout (in seconds)", ox.Default("30"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more plugins"),
				ox.Banner("Display detailed information on one or more plugins"),
				ox.Spec("[OPTIONS] PLUGIN [PLUGIN...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("install", "Install a plugin"),
				ox.Banner("Install a plugin"),
				ox.Spec("[OPTIONS] PLUGIN [KEY=VALUE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("alias", "Local name for plugin", ox.Section(0)).
					Bool("disable", "Do not enable the plugin on install", ox.Section(0)).
					Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
					Bool("grant-all-permissions", "Grant all permissions necessary to run", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List plugins"),
				ox.Banner("List plugins"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("plugin ls", "plugin list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Provide filter values (e.g. \"enabled=true\")", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("no-trunc", "Don't truncate output", ox.Section(0)).
					Bool("quiet", "Only display plugin IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("push", "Push a plugin to a registry"),
				ox.Banner("Push a plugin to a registry"),
				ox.Spec("[OPTIONS] PLUGIN[:TAG]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("disable-content-trust", "Skip image signing", ox.Default("true"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more plugins"),
				ox.Banner("Remove one or more plugins"),
				ox.Spec("[OPTIONS] PLUGIN [PLUGIN...]"),
				ox.Aliases("plugin rm", "plugin remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the removal of an active plugin", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("set", "Change settings for a plugin"),
				ox.Spec("PLUGIN KEY=VALUE [KEY=VALUE...]"),
				ox.Section(0),
				ox.Footer("Change settings for a plugin"),
			),
			ox.Sub(
				ox.Usage("upgrade", "Upgrade an existing plugin"),
				ox.Banner("Upgrade an existing plugin"),
				ox.Spec("[OPTIONS] PLUGIN [REMOTE]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("disable-content-trust", "Skip image verification", ox.Default("true"), ox.Section(0)).
					Bool("grant-all-permissions", "Grant all permissions necessary to run", ox.Section(0)).
					Bool("skip-remote-check", "Do not check if specified remote plugin", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("system", "Manage Docker"),
			ox.Banner("Manage Docker"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker system COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("df", "Show docker disk usage"),
				ox.Banner("Show docker disk usage"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("verbose", "Show detailed information on space usage", ox.Short("v"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("events", "Get real time events from the server"),
				ox.Banner("Get real time events from the server"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("system events"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					String("since", "Show all events created since timestamp", ox.Section(0)).
					String("until", "Stream events until this timestamp", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("info", "Display system-wide information"),
				ox.Banner("Display system-wide information"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("system info"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("prune", "Remove unused data"),
				ox.Banner("Remove unused data"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Remove all unused images not just dangling ones", ox.Short("a"), ox.Section(0)).
					String("filter", "Provide filter values (e.g. \"label=<key>=<value>\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)).
					Bool("volumes", "Prune anonymous volumes", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("trust", "Manage trust on Docker images"),
			ox.Banner("Manage trust on Docker images"),
			ox.Spec("COMMAND"),
			ox.Sections("Management Commands", "Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker trust COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("key", "Manage keys for signing Docker images"),
				ox.Banner("Manage keys for signing Docker images"),
				ox.Spec("COMMAND"),
				ox.Sections("Commands"),
				ox.Section(0),
				ox.Footer("Run 'docker trust key COMMAND --help' for more information on a command."),
				ox.Sub(
					ox.Usage("generate", "Generate and load a signing key-pair"),
					ox.Banner("Generate and load a signing key-pair"),
					ox.Spec("NAME"),
					ox.Section(0),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Flags().
						String("dir", "Directory to generate key in, defaults to current", ox.Section(0)),
				),
				ox.Sub(
					ox.Usage("load", "Load a private key file for signing"),
					ox.Banner("Load a private key file for signing"),
					ox.Spec("[OPTIONS] KEYFILE"),
					ox.Section(0),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Flags().
						String("name", "Name for the loaded key", ox.Default("signer"), ox.Section(0)),
				),
			),
			ox.Sub(
				ox.Usage("signer", "Manage entities who can sign Docker images"),
				ox.Banner("Manage entities who can sign Docker images"),
				ox.Spec("COMMAND"),
				ox.Sections("Commands"),
				ox.Section(0),
				ox.Footer("Run 'docker trust signer COMMAND --help' for more information on a command."),
				ox.Sub(
					ox.Usage("add", "Add a signer"),
					ox.Banner("Add a signer"),
					ox.Spec("OPTIONS NAME REPOSITORY [REPOSITORY...]"),
					ox.Section(0),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Flags().
						Slice("key", "Path to the signer's public key file", ox.Section(0)),
				),
				ox.Sub(
					ox.Usage("remove", "Remove a signer"),
					ox.Banner("Remove a signer"),
					ox.Spec("[OPTIONS] NAME REPOSITORY [REPOSITORY...]"),
					ox.Section(0),
					ox.Help(ox.Sections(
						"Options",
					)),
					ox.Flags().
						Bool("force", "Do not prompt for confirmation before removing the most", ox.Short("f"), ox.Section(0)),
				),
			),
			ox.Sub(
				ox.Usage("inspect", "Return low-level information about keys and signatures"),
				ox.Banner("Return low-level information about keys and signatures"),
				ox.Spec("IMAGE[:TAG] [IMAGE[:TAG]...]"),
				ox.Section(1),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("pretty", "Print the information in a human friendly format", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("revoke", "Remove trust for an image"),
				ox.Banner("Remove trust for an image"),
				ox.Spec("[OPTIONS] IMAGE[:TAG]"),
				ox.Section(1),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("yes", "Do not prompt for confirmation", ox.Short("y"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("sign", "Sign an image"),
				ox.Banner("Sign an image"),
				ox.Spec("IMAGE:TAG"),
				ox.Section(1),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("local", "Sign a locally tagged image", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("volume", "Manage volumes"),
			ox.Banner("Manage volumes"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(1),
			ox.Footer("Run 'docker volume COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a volume"),
				ox.Banner("Create a volume"),
				ox.Spec("[OPTIONS] [VOLUME]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("availability", "Cluster Volume availability (\"active\",", ox.Section(0)).
					String("driver", "Specify volume driver name", ox.Default("local"), ox.Short("d"), ox.Section(0)).
					String("group", "Cluster Volume group (cluster volumes)", ox.Section(0)).
					Slice("label", "Set metadata for a volume", ox.Section(0)).
					Slice("limit-bytes", "Minimum size of the Cluster Volume in bytes", ox.Elem(ox.UintT), ox.Section(0)).
					Map("opt", "Set driver specific options", ox.Default("map[]"), ox.Short("o"), ox.Section(0)).
					Slice("required-bytes", "Maximum size of the Cluster Volume in bytes", ox.Elem(ox.UintT), ox.Section(0)).
					String("scope", "Cluster Volume access scope (\"single\",", ox.Section(0)).
					Map("secret", "Cluster Volume secrets", ox.Default("map[]"), ox.Section(0)).
					String("sharing", "Cluster Volume access sharing (\"none\",", ox.Section(0)).
					Slice("topology-preferred", "A topology that the Cluster Volume", ox.Section(0)).
					Slice("topology-required", "A topology that the Cluster Volume must", ox.Section(0)).
					String("type", "Cluster Volume access type (\"mount\",", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more volumes"),
				ox.Banner("Display detailed information on one or more volumes"),
				ox.Spec("[OPTIONS] VOLUME [VOLUME...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List volumes"),
				ox.Banner("List volumes"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("volume ls", "volume list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("cluster", "Display only cluster volumes, and use cluster", ox.Section(0)).
					String("filter", "Provide filter values (e.g. \"dangling=true\")", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display volume names", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("prune", "Remove unused local volumes"),
				ox.Banner("Remove unused local volumes"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("all", "Remove all unused volumes, not just anonymous ones", ox.Short("a"), ox.Section(0)).
					String("filter", "Provide filter values (e.g. \"label=<label>\")", ox.Spec("filter"), ox.Section(0)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more volumes"),
				ox.Banner("Remove one or more volumes. You cannot remove a volume that is in use by a container."),
				ox.Spec("[OPTIONS] VOLUME [VOLUME...]"),
				ox.Aliases("volume rm", "volume remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force the removal of one or more volumes", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("update", "Update a volume (cluster volumes only)"),
				ox.Banner("Update a volume (cluster volumes only)"),
				ox.Spec("[OPTIONS] [VOLUME]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("availability", "Cluster Volume availability (\"active\",", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("config", "Manage Swarm configs"),
			ox.Banner("Manage Swarm configs"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker config COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a config from a file or STDIN"),
				ox.Banner("Create a config from a file or STDIN"),
				ox.Spec("[OPTIONS] CONFIG file|-"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("label", "Config labels", ox.Short("l"), ox.Section(0)).
					String("template-driver", "Template driver", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more configs"),
				ox.Banner("Display detailed information on one or more configs"),
				ox.Spec("[OPTIONS] CONFIG [CONFIG...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("pretty", "Print the information in a human friendly format", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List configs"),
				ox.Banner("List configs"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("config ls", "config list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more configs"),
				ox.Banner("Remove one or more configs"),
				ox.Spec("CONFIG [CONFIG...]"),
				ox.Aliases("config rm", "config remove"),
				ox.Section(0),
			),
		),
		ox.Sub(
			ox.Usage("node", "Manage Swarm nodes"),
			ox.Banner("Manage Swarm nodes"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker node COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("demote", "Demote one or more nodes from manager in the swarm"),
				ox.Spec("NODE [NODE...]"),
				ox.Section(0),
				ox.Footer("Demote one or more nodes from manager in the swarm"),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more nodes"),
				ox.Banner("Display detailed information on one or more nodes"),
				ox.Spec("[OPTIONS] self|NODE [NODE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("pretty", "Print the information in a human friendly format", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List nodes in the swarm"),
				ox.Banner("List nodes in the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("node ls", "node list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("promote", "Promote one or more nodes to manager in the swarm"),
				ox.Spec("NODE [NODE...]"),
				ox.Section(0),
				ox.Footer("Promote one or more nodes to manager in the swarm"),
			),
			ox.Sub(
				ox.Usage("ps", "List tasks running on one or more nodes, defaults to current node"),
				ox.Banner("List tasks running on one or more nodes, defaults to current node"),
				ox.Spec("[OPTIONS] [NODE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Pretty-print tasks using a Go template", ox.Section(0)).
					Bool("no-resolve", "Do not map IDs to Names", ox.Section(0)).
					Bool("no-trunc", "Do not truncate output", ox.Section(0)).
					Bool("quiet", "Only display task IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more nodes from the swarm"),
				ox.Banner("Remove one or more nodes from the swarm"),
				ox.Spec("[OPTIONS] NODE [NODE...]"),
				ox.Aliases("node rm", "node remove"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force remove a node from the swarm", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("update", "Update a node"),
				ox.Banner("Update a node"),
				ox.Spec("[OPTIONS] NODE"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("availability", "Availability of the node (\"active\",", ox.Section(0)).
					Slice("label-add", "Add or update a node label (\"key=value\")", ox.Section(0)).
					Slice("label-rm", "Remove a node label if exists", ox.Section(0)).
					String("role", "Role of the node (\"worker\", \"manager\")", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("secret", "Manage Swarm secrets"),
			ox.Banner("Manage Swarm secrets"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker secret COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a secret from a file or STDIN as content"),
				ox.Banner("Create a secret from a file or STDIN as content"),
				ox.Spec("[OPTIONS] SECRET [file|-]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("driver", "Secret driver", ox.Short("d"), ox.Section(0)).
					Slice("label", "Secret labels", ox.Short("l"), ox.Section(0)).
					String("template-driver", "Template driver", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more secrets"),
				ox.Banner("Display detailed information on one or more secrets"),
				ox.Spec("[OPTIONS] SECRET [SECRET...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("pretty", "Print the information in a human friendly format", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List secrets"),
				ox.Banner("List secrets"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("secret ls", "secret list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more secrets"),
				ox.Banner("Remove one or more secrets"),
				ox.Spec("SECRET [SECRET...]"),
				ox.Aliases("secret rm", "secret remove"),
				ox.Section(0),
			),
		),
		ox.Sub(
			ox.Usage("service", "Manage Swarm services"),
			ox.Banner("Manage Swarm services"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker service COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("create", "Create a new service"),
				ox.Banner("Create a new service"),
				ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
					String("config", "Specify configurations to", ox.Spec("config"), ox.Section(0)).
					Slice("constraint", "Placement constraints", ox.Section(0)).
					Slice("container-label", "Container labels", ox.Section(0)).
					String("credential-spec", "Credential spec for managed", ox.Spec("credential-spec"), ox.Section(0)).
					Bool("detach", "Exit immediately instead of", ox.Short("d"), ox.Section(0)).
					Slice("dns", "Set custom DNS servers", ox.Section(0)).
					Slice("dns-option", "Set DNS options", ox.Section(0)).
					Slice("dns-search", "Set custom DNS search domains", ox.Section(0)).
					String("endpoint-mode", "Endpoint mode (vip or dnsrr)", ox.Section(0)).
					String("entrypoint", "Overwrite the default", ox.Spec("command"), ox.Section(0)).
					Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
					Slice("env-file", "Read in a file of environment", ox.Section(0)).
					Slice("generic-resource", "User defined resources", ox.Section(0)).
					Slice("group", "Set one or more supplementary", ox.Section(0)).
					String("health-cmd", "Command to run to check health", ox.Section(0)).
					Duration("health-interval", "Time between running the check", ox.Section(0)).
					Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
					Duration("health-start-interval", "Time between running the check", ox.Section(0)).
					Duration("health-start-period", "Start period for the container", ox.Section(0)).
					Duration("health-timeout", "Maximum time to allow one", ox.Section(0)).
					Slice("host", "Set one or more custom", ox.Section(0)).
					String("hostname", "Container hostname", ox.Section(0)).
					Bool("init", "Use an init inside each", ox.Section(0)).
					String("isolation", "Service container isolation mode", ox.Section(0)).
					Slice("label", "Service labels", ox.Short("l"), ox.Section(0)).
					Float32("limit-cpu", "Limit CPUs", ox.Spec("decimal"), ox.Section(0)).
					Slice("limit-memory", "Limit Memory", ox.Elem(ox.UintT), ox.Section(0)).
					Int("limit-pids", "Limit maximum number of", ox.Section(0)).
					String("log-driver", "Logging driver for service", ox.Section(0)).
					Slice("log-opt", "Logging driver options", ox.Section(0)).
					Uint("max-concurrent", "Number of job tasks to run", ox.Section(0)).
					String("mode", "Service mode (\"replicated\",", ox.Section(0)).
					String("mount", "Attach a filesystem mount to", ox.Spec("mount"), ox.Section(0)).
					String("name", "Service name", ox.Section(0)).
					String("network", "Network attachments", ox.Spec("network"), ox.Section(0)).
					Bool("no-healthcheck", "Disable any", ox.Section(0)).
					Bool("no-resolve-image", "Do not query the registry to", ox.Section(0)).
					Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
					String("placement-pref", "Add a placement preference", ox.Spec("pref"), ox.Section(0)).
					Uint("publish", "Publish a port as a node port", ox.Spec("port"), ox.Short("p"), ox.Section(0)).
					Bool("quiet", "Suppress progress output", ox.Short("q"), ox.Section(0)).
					Bool("read-only", "Mount the container's root", ox.Section(0)).
					Uint("replicas", "Number of tasks", ox.Section(0)).
					Uint("replicas-max-per-node", "Maximum number of tasks per", ox.Section(0)).
					Float32("reserve-cpu", "Reserve CPUs", ox.Spec("decimal"), ox.Section(0)).
					Slice("reserve-memory", "Reserve Memory", ox.Elem(ox.UintT), ox.Section(0)).
					String("restart-condition", "Restart when condition is met", ox.Section(0)).
					Duration("restart-delay", "Delay between restart attempts", ox.Section(0)).
					Uint("restart-max-attempts", "Maximum number of restarts", ox.Section(0)).
					Duration("restart-window", "Window used to evaluate the", ox.Section(0)).
					Duration("rollback-delay", "Delay between task rollbacks", ox.Section(0)).
					String("rollback-failure-action", "Action on rollback failure", ox.Section(0)).
					Float32("rollback-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float"), ox.Section(0)).
					Duration("rollback-monitor", "Duration after each task", ox.Section(0)).
					String("rollback-order", "Rollback order (\"start-first\",", ox.Section(0)).
					Uint("rollback-parallelism", "Maximum number of tasks rolled", ox.Section(0)).
					String("secret", "Specify secrets to expose to", ox.Spec("secret"), ox.Section(0)).
					Duration("stop-grace-period", "Time to wait before force", ox.Section(0)).
					String("stop-signal", "Signal to stop the container", ox.Section(0)).
					Slice("sysctl", "Sysctl options", ox.Section(0)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
					String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)).
					Duration("update-delay", "Delay between updates", ox.Section(0)).
					String("update-failure-action", "Action on update failure", ox.Section(0)).
					Float32("update-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float"), ox.Section(0)).
					Duration("update-monitor", "Duration after each task", ox.Section(0)).
					String("update-order", "Update order (\"start-first\",", ox.Section(0)).
					Uint("update-parallelism", "Maximum number of tasks", ox.Section(0)).
					String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
					Bool("with-registry-auth", "Send registry authentication", ox.Section(0)).
					String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("inspect", "Display detailed information on one or more services"),
				ox.Banner("Display detailed information on one or more services"),
				ox.Spec("[OPTIONS] SERVICE [SERVICE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
					Bool("pretty", "Print the information in a human friendly format", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("logs", "Fetch the logs of a service or task"),
				ox.Banner("Fetch the logs of a service or task"),
				ox.Spec("[OPTIONS] SERVICE|TASK"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("details", "Show extra details provided to logs", ox.Section(0)).
					Bool("follow", "Follow log output", ox.Short("f"), ox.Section(0)).
					Bool("no-resolve", "Do not map IDs to Names in output", ox.Section(0)).
					Bool("no-task-ids", "Do not include task IDs in output", ox.Section(0)).
					Bool("no-trunc", "Do not truncate output", ox.Section(0)).
					Bool("raw", "Do not neatly format logs", ox.Section(0)).
					String("since", "Show logs since timestamp (e.g.", ox.Section(0)).
					String("tail", "Number of lines to show from the end of the logs", ox.Short("n"), ox.Section(0)).
					Bool("timestamps", "Show timestamps", ox.Short("t"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List services"),
				ox.Banner("List services"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("service ls", "service list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ps", "List the tasks of one or more services"),
				ox.Banner("List the tasks of one or more services"),
				ox.Spec("[OPTIONS] SERVICE [SERVICE...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Pretty-print tasks using a Go template", ox.Section(0)).
					Bool("no-resolve", "Do not map IDs to Names", ox.Section(0)).
					Bool("no-trunc", "Do not truncate output", ox.Section(0)).
					Bool("quiet", "Only display task IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more services"),
				ox.Banner("Remove one or more services"),
				ox.Spec("SERVICE [SERVICE...]"),
				ox.Aliases("service rm", "service remove"),
				ox.Section(0),
			),
			ox.Sub(
				ox.Usage("rollback", "Revert changes to a service's configuration"),
				ox.Banner("Revert changes to a service's configuration"),
				ox.Spec("[OPTIONS] SERVICE"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("detach", "Exit immediately instead of waiting for the service to", ox.Short("d"), ox.Section(0)).
					Bool("quiet", "Suppress progress output", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("scale", "Scale one or multiple replicated services"),
				ox.Banner("Scale one or multiple replicated services"),
				ox.Spec("SERVICE=REPLICAS [SERVICE=REPLICAS...]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("detach", "Exit immediately instead of waiting for the service to", ox.Short("d"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("update", "Update a service"),
				ox.Banner("Update a service"),
				ox.Spec("[OPTIONS] SERVICE"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("args", "Service command args", ox.Spec("command"), ox.Section(0)).
					Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
					Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
					String("config-add", "Add or update a config file on", ox.Spec("config"), ox.Section(0)).
					Slice("config-rm", "Remove a configuration file", ox.Section(0)).
					Slice("constraint-add", "Add or update a placement", ox.Section(0)).
					Slice("constraint-rm", "Remove a constraint", ox.Section(0)).
					Slice("container-label-add", "Add or update a container label", ox.Section(0)).
					Slice("container-label-rm", "Remove a container label by its key", ox.Section(0)).
					String("credential-spec", "Credential spec for managed", ox.Spec("credential-spec"), ox.Section(0)).
					Bool("detach", "Exit immediately instead of", ox.Short("d"), ox.Section(0)).
					Slice("dns-add", "Add or update a custom DNS server", ox.Section(0)).
					Slice("dns-option-add", "Add or update a DNS option", ox.Section(0)).
					Slice("dns-option-rm", "Remove a DNS option", ox.Section(0)).
					Slice("dns-rm", "Remove a custom DNS server", ox.Section(0)).
					Slice("dns-search-add", "Add or update a custom DNS", ox.Section(0)).
					Slice("dns-search-rm", "Remove a DNS search domain", ox.Section(0)).
					String("endpoint-mode", "Endpoint mode (vip or dnsrr)", ox.Section(0)).
					String("entrypoint", "Overwrite the default", ox.Spec("command"), ox.Section(0)).
					Slice("env-add", "Add or update an environment", ox.Section(0)).
					Slice("env-rm", "Remove an environment variable", ox.Section(0)).
					Bool("force", "Force update even if no", ox.Section(0)).
					Slice("generic-resource-add", "Add a Generic resource", ox.Section(0)).
					Slice("generic-resource-rm", "Remove a Generic resource", ox.Section(0)).
					Slice("group-add", "Add an additional", ox.Section(0)).
					Slice("group-rm", "Remove a previously added", ox.Section(0)).
					String("health-cmd", "Command to run to check health", ox.Section(0)).
					Duration("health-interval", "Time between running the check", ox.Section(0)).
					Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
					Duration("health-start-interval", "Time between running the check", ox.Section(0)).
					Duration("health-start-period", "Start period for the container", ox.Section(0)).
					Duration("health-timeout", "Maximum time to allow one", ox.Section(0)).
					Slice("host-add", "Add a custom host-to-IP", ox.Section(0)).
					Slice("host-rm", "Remove a custom host-to-IP", ox.Section(0)).
					String("hostname", "Container hostname", ox.Section(0)).
					String("image", "Service image tag", ox.Section(0)).
					Bool("init", "Use an init inside each", ox.Section(0)).
					String("isolation", "Service container isolation mode", ox.Section(0)).
					Slice("label-add", "Add or update a service label", ox.Section(0)).
					Slice("label-rm", "Remove a label by its key", ox.Section(0)).
					Float32("limit-cpu", "Limit CPUs", ox.Spec("decimal"), ox.Section(0)).
					Slice("limit-memory", "Limit Memory", ox.Elem(ox.UintT), ox.Section(0)).
					Int("limit-pids", "Limit maximum number of", ox.Section(0)).
					String("log-driver", "Logging driver for service", ox.Section(0)).
					Slice("log-opt", "Logging driver options", ox.Section(0)).
					Uint("max-concurrent", "Number of job tasks to run", ox.Section(0)).
					String("mount-add", "Add or update a mount on a service", ox.Spec("mount"), ox.Section(0)).
					Slice("mount-rm", "Remove a mount by its target path", ox.Section(0)).
					String("network-add", "Add a network", ox.Spec("network"), ox.Section(0)).
					Slice("network-rm", "Remove a network", ox.Section(0)).
					Bool("no-healthcheck", "Disable any", ox.Section(0)).
					Bool("no-resolve-image", "Do not query the registry to", ox.Section(0)).
					Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
					String("placement-pref-add", "Add a placement preference", ox.Spec("pref"), ox.Section(0)).
					String("placement-pref-rm", "Remove a placement preference", ox.Spec("pref"), ox.Section(0)).
					Uint("publish-add", "Add or update a published port", ox.Spec("port"), ox.Section(0)).
					Uint("publish-rm", "Remove a published port by its", ox.Spec("port"), ox.Section(0)).
					Bool("quiet", "Suppress progress output", ox.Short("q"), ox.Section(0)).
					Bool("read-only", "Mount the container's root", ox.Section(0)).
					Uint("replicas", "Number of tasks", ox.Section(0)).
					Uint("replicas-max-per-node", "Maximum number of tasks per", ox.Section(0)).
					Float32("reserve-cpu", "Reserve CPUs", ox.Spec("decimal"), ox.Section(0)).
					Slice("reserve-memory", "Reserve Memory", ox.Elem(ox.UintT), ox.Section(0)).
					String("restart-condition", "Restart when condition is met", ox.Section(0)).
					Duration("restart-delay", "Delay between restart attempts", ox.Section(0)).
					Uint("restart-max-attempts", "Maximum number of restarts", ox.Section(0)).
					Duration("restart-window", "Window used to evaluate the", ox.Section(0)).
					Bool("rollback", "Rollback to previous specification", ox.Section(0)).
					Duration("rollback-delay", "Delay between task rollbacks", ox.Section(0)).
					String("rollback-failure-action", "Action on rollback failure", ox.Section(0)).
					Float32("rollback-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float"), ox.Section(0)).
					Duration("rollback-monitor", "Duration after each task", ox.Section(0)).
					String("rollback-order", "Rollback order (\"start-first\",", ox.Section(0)).
					Uint("rollback-parallelism", "Maximum number of tasks rolled", ox.Section(0)).
					String("secret-add", "Add or update a secret on a service", ox.Spec("secret"), ox.Section(0)).
					Slice("secret-rm", "Remove a secret", ox.Section(0)).
					Duration("stop-grace-period", "Time to wait before force", ox.Section(0)).
					String("stop-signal", "Signal to stop the container", ox.Section(0)).
					Slice("sysctl-add", "Add or update a Sysctl option", ox.Section(0)).
					Slice("sysctl-rm", "Remove a Sysctl option", ox.Section(0)).
					Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
					String("ulimit-add", "Add or update a ulimit option", ox.Spec("ulimit"), ox.Section(0)).
					Slice("ulimit-rm", "Remove a ulimit option", ox.Section(0)).
					Duration("update-delay", "Delay between updates", ox.Section(0)).
					String("update-failure-action", "Action on update failure", ox.Section(0)).
					Float32("update-max-failure-ratio", "Failure rate to tolerate", ox.Spec("float"), ox.Section(0)).
					Duration("update-monitor", "Duration after each task", ox.Section(0)).
					String("update-order", "Update order (\"start-first\",", ox.Section(0)).
					Uint("update-parallelism", "Maximum number of tasks", ox.Section(0)).
					String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
					Bool("with-registry-auth", "Send registry authentication", ox.Section(0)).
					String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("stack", "Manage Swarm stacks"),
			ox.Banner("Manage Swarm stacks"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker stack COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("config", "Outputs the final config file, after doing merges and interpolations"),
				ox.Banner("Outputs the final config file, after doing merges and interpolations"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("compose-file", "Path to a Compose file, or \"-\" to read", ox.Short("c"), ox.Section(0)).
					Bool("skip-interpolation", "Skip interpolation and output only merged", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("deploy", "Deploy a new stack or update an existing stack"),
				ox.Banner("Deploy a new stack or update an existing stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Aliases("stack deploy", "stack up"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Slice("compose-file", "Path to a Compose file, or \"-\" to read", ox.Short("c"), ox.Section(0)).
					Bool("detach", "Exit immediately instead of waiting for", ox.Short("d"), ox.Section(0)).
					Bool("prune", "Prune services that are no longer referenced", ox.Section(0)).
					Bool("quiet", "Suppress progress output", ox.Short("q"), ox.Section(0)).
					String("resolve-image", "Query the registry to resolve image digest", ox.Section(0)).
					Bool("with-registry-auth", "Send registry authentication details to", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ls", "List stacks"),
				ox.Banner("List stacks"),
				ox.Spec("[OPTIONS]"),
				ox.Aliases("stack ls", "stack list"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("format", "Format output using a custom template:", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("ps", "List the tasks in the stack"),
				ox.Banner("List the tasks in the stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("no-resolve", "Do not map IDs to Names", ox.Section(0)).
					Bool("no-trunc", "Do not truncate output", ox.Section(0)).
					Bool("quiet", "Only display task IDs", ox.Short("q"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("rm", "Remove one or more stacks"),
				ox.Banner("Remove one or more stacks"),
				ox.Spec("[OPTIONS] STACK [STACK...]"),
				ox.Aliases("stack rm", "stack remove", "stack down"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("detach", "Do not wait for stack removal", ox.Default("true"), ox.Short("d"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("services", "List the services in the stack"),
				ox.Banner("List the services in the stack"),
				ox.Spec("[OPTIONS] STACK"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
					String("format", "Format output using a custom template:", ox.Section(0)).
					Bool("quiet", "Only display IDs", ox.Short("q"), ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("swarm", "Manage Swarm"),
			ox.Banner("Manage Swarm"),
			ox.Spec("COMMAND"),
			ox.Sections("Commands"),
			ox.Section(2),
			ox.Footer("Run 'docker swarm COMMAND --help' for more information on a command."),
			ox.Sub(
				ox.Usage("ca", "Display and rotate the root CA"),
				ox.Banner("Display and rotate the root CA"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("ca-cert", "Path to the PEM-formatted root CA", ox.Spec("pem-file"), ox.Section(0)).
					String("ca-key", "Path to the PEM-formatted root CA key", ox.Spec("pem-file"), ox.Section(0)).
					Duration("cert-expiry", "Validity period for node certificates", ox.Section(0)).
					Bool("detach", "Exit immediately instead of waiting for", ox.Short("d"), ox.Section(0)).
					String("external-ca", "Specifications of one or more", ox.Spec("external-ca"), ox.Section(0)).
					Bool("quiet", "Suppress progress output", ox.Short("q"), ox.Section(0)).
					Bool("rotate", "Rotate the swarm CA - if no certificate", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("init", "Initialize a swarm"),
				ox.Banner("Initialize a swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("advertise-addr", "Advertised address", ox.Section(0)).
					Bool("autolock", "Enable manager autolocking", ox.Section(0)).
					String("availability", "Availability of the node", ox.Section(0)).
					Duration("cert-expiry", "Validity period for node", ox.Section(0)).
					String("data-path-addr", "Address or interface to", ox.Section(0)).
					Uint32("data-path-port", "Port number to use for", ox.Section(0)).
					Slice("default-addr-pool", "default address pool in", ox.Elem(ox.CIDRT), ox.Section(0)).
					Uint32("default-addr-pool-mask-length", "default address pool", ox.Section(0)).
					Duration("dispatcher-heartbeat", "Dispatcher heartbeat", ox.Section(0)).
					String("external-ca", "Specifications of one or", ox.Spec("external-ca"), ox.Section(0)).
					Bool("force-new-cluster", "Force create a new cluster", ox.Section(0)).
					AddrPort("listen-addr", "Listen address (format:", ox.Section(0)).
					Uint("max-snapshots", "Number of additional Raft", ox.Section(0)).
					Uint("snapshot-interval", "Number of log entries", ox.Section(0)).
					Int("task-history-limit", "Task history retention", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("join", "Join a swarm as a node and/or manager"),
				ox.Banner("Join a swarm as a node and/or manager"),
				ox.Spec("[OPTIONS] HOST:PORT"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					String("advertise-addr", "Advertised address (format:", ox.Section(0)).
					String("availability", "Availability of the node (\"active\",", ox.Section(0)).
					String("data-path-addr", "Address or interface to use for data path", ox.Section(0)).
					AddrPort("listen-addr", "Listen address (format:", ox.Section(0)).
					String("token", "Token for entry into the swarm", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("join-token", "Manage join tokens"),
				ox.Banner("Manage join tokens"),
				ox.Spec("[OPTIONS] (worker|manager)"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("quiet", "Only display token", ox.Short("q"), ox.Section(0)).
					Bool("rotate", "Rotate join token", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("leave", "Leave the swarm"),
				ox.Banner("Leave the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("force", "Force this node to leave the swarm, ignoring warnings", ox.Short("f"), ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("unlock", "Unlock swarm"),
				ox.Section(0),
				ox.Footer("Unlock swarm"),
			),
			ox.Sub(
				ox.Usage("unlock-key", "Manage the unlock key"),
				ox.Banner("Manage the unlock key"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("quiet", "Only display token", ox.Short("q"), ox.Section(0)).
					Bool("rotate", "Rotate unlock key", ox.Section(0)),
			),
			ox.Sub(
				ox.Usage("update", "Update the swarm"),
				ox.Banner("Update the swarm"),
				ox.Spec("[OPTIONS]"),
				ox.Section(0),
				ox.Help(ox.Sections(
					"Options",
				)),
				ox.Flags().
					Bool("autolock", "Change manager autolocking", ox.Section(0)).
					Duration("cert-expiry", "Validity period for node", ox.Section(0)).
					Duration("dispatcher-heartbeat", "Dispatcher heartbeat period", ox.Section(0)).
					String("external-ca", "Specifications of one or more", ox.Spec("external-ca"), ox.Section(0)).
					Uint("max-snapshots", "Number of additional Raft", ox.Section(0)).
					Uint("snapshot-interval", "Number of log entries between", ox.Section(0)).
					Int("task-history-limit", "Task history retention limit", ox.Section(0)),
			),
		),
		ox.Sub(
			ox.Usage("attach", "Attach local standard input, output, and error streams to a running container"),
			ox.Banner("Attach local standard input, output, and error streams to a running container"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("container attach"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
				Bool("no-stdin", "Do not attach STDIN", ox.Section(0)).
				Bool("sig-proxy", "Proxy all received signals to the process", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("commit", "Create a new image from a container's changes"),
			ox.Banner("Create a new image from a container's changes"),
			ox.Spec("[OPTIONS] CONTAINER [REPOSITORY[:TAG]]"),
			ox.Aliases("container commit"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("author", "Author (e.g., \"John Hannibal Smith", ox.Short("a"), ox.Section(0)).
				Slice("change", "Apply Dockerfile instruction to the created image", ox.Short("c"), ox.Section(0)).
				String("message", "Commit message", ox.Short("m"), ox.Section(0)).
				Bool("pause", "Pause container during commit", ox.Default("true"), ox.Short("p"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
			ox.Banner("docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH\n\nCopy files/folders between a container and the local filesystem\n\nUse '-' as the source to read a tar archive from stdin\nand extract it to a directory destination in a container.\nUse '-' as the destination to stream a tar archive of a\ncontainer source to stdout."),
			ox.Spec("[OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-"),
			ox.Aliases("container cp"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("archive", "Archive mode (copy all uid/gid information)", ox.Short("a"), ox.Section(0)).
				Bool("follow-link", "Always follow symbol link in SRC_PATH", ox.Short("L"), ox.Section(0)).
				Bool("quiet", "Suppress progress output during copy. Progress", ox.Short("q"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("create", "Create a new container"),
			ox.Banner("Create a new container"),
			ox.Spec("[OPTIONS] IMAGE [COMMAND] [ARG...]"),
			ox.Aliases("container create"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping", ox.Section(0)).
				Map("annotation", "Add an annotation to the", ox.Section(0)).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Short("a"), ox.Section(0)).
				Uint16("blkio-weight", "Block IO (relative weight),", ox.Section(0)).
				Slice("blkio-weight-device", "Block IO weight (relative device", ox.Section(0)).
				Slice("cap-add", "Add Linux capabilities", ox.Section(0)).
				Slice("cap-drop", "Drop Linux capabilities", ox.Section(0)).
				String("cgroup-parent", "Optional parent cgroup for the", ox.Section(0)).
				String("cgroupns", "Cgroup namespace to use", ox.Section(0)).
				String("cidfile", "Write the container ID to the file", ox.Section(0)).
				Int("cpu-count", "CPU count (Windows only)", ox.Section(0)).
				Int("cpu-percent", "CPU percent (Windows only)", ox.Section(0)).
				Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-rt-period", "Limit CPU real-time period in", ox.Section(0)).
				Int("cpu-rt-runtime", "Limit CPU real-time runtime in", ox.Section(0)).
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
				String("cpuset-cpus", "CPUs in which to allow execution", ox.Section(0)).
				String("cpuset-mems", "MEMs in which to allow execution", ox.Section(0)).
				Slice("device", "Add a host device to the container", ox.Section(0)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed", ox.Section(0)).
				Slice("device-read-bps", "Limit read rate (bytes per", ox.Section(0)).
				Slice("device-read-iops", "Limit read rate (IO per second)", ox.Section(0)).
				Slice("device-write-bps", "Limit write rate (bytes per", ox.Section(0)).
				Slice("device-write-iops", "Limit write rate (IO per second)", ox.Section(0)).
				Bool("disable-content-trust", "Skip image verification (default", ox.Section(0)).
				Slice("dns", "Set custom DNS servers", ox.Section(0)).
				Slice("dns-option", "Set DNS options", ox.Section(0)).
				Slice("dns-search", "Set custom DNS search domains", ox.Section(0)).
				String("domainname", "Container NIS domain name", ox.Section(0)).
				String("entrypoint", "Overwrite the default ENTRYPOINT", ox.Section(0)).
				Slice("env", "Set environment variables", ox.Short("e"), ox.Section(0)).
				Slice("env-file", "Read in a file of environment", ox.Section(0)).
				Slice("expose", "Expose a port or a range of ports", ox.Section(0)).
				String("gpus", "GPU devices to add to the", ox.Spec("gpu-request"), ox.Section(0)).
				Slice("group-add", "Add additional groups to join", ox.Section(0)).
				String("health-cmd", "Command to run to check health", ox.Section(0)).
				Duration("health-interval", "Time between running the check", ox.Section(0)).
				Int("health-retries", "Consecutive failures needed to", ox.Section(0)).
				Duration("health-start-interval", "Time between running the check", ox.Section(0)).
				Duration("health-start-period", "Start period for the container", ox.Section(0)).
				Duration("health-timeout", "Maximum time to allow one check", ox.Section(0)).
				String("hostname", "Container host name", ox.Short("h"), ox.Section(0)).
				Bool("init", "Run an init inside the container", ox.Section(0)).
				Bool("interactive", "Keep STDIN open even if not attached", ox.Short("i"), ox.Section(0)).
				Slice("io-maxbandwidth", "Maximum IO bandwidth limit for", ox.Elem(ox.UintT), ox.Section(0)).
				Uint("io-maxiops", "Maximum IOps limit for the", ox.Section(0)).
				String("ip", "IPv4 address (e.g., 172.30.100.104)", ox.Section(0)).
				String("ip6", "IPv6 address (e.g., 2001:db8::33)", ox.Section(0)).
				String("ipc", "IPC mode to use", ox.Section(0)).
				String("isolation", "Container isolation technology", ox.Section(0)).
				Slice("kernel-memory", "Kernel memory limit", ox.Elem(ox.UintT), ox.Section(0)).
				Slice("label", "Set meta data on a container", ox.Short("l"), ox.Section(0)).
				Slice("label-file", "Read in a line delimited file of", ox.Section(0)).
				Slice("link", "Add link to another container", ox.Section(0)).
				Slice("link-local-ip", "Container IPv4/IPv6 link-local", ox.Section(0)).
				String("log-driver", "Logging driver for the container", ox.Section(0)).
				Slice("log-opt", "Log driver options", ox.Section(0)).
				String("mac-address", "Container MAC address (e.g.,", ox.Section(0)).
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
				Slice("memory-swap", "Swap limit equal to memory plus", ox.Elem(ox.UintT), ox.Section(0)).
				Int("memory-swappiness", "Tune container memory swappiness", ox.Section(0)).
				String("mount", "Attach a filesystem mount to the", ox.Spec("mount"), ox.Section(0)).
				String("name", "Assign a name to the container", ox.Section(0)).
				String("network", "Connect a container to a network", ox.Spec("network"), ox.Section(0)).
				Slice("network-alias", "Add network-scoped alias for the", ox.Section(0)).
				Bool("no-healthcheck", "Disable any container-specified", ox.Section(0)).
				Bool("oom-kill-disable", "Disable OOM Killer", ox.Section(0)).
				Int("oom-score-adj", "Tune host's OOM preferences", ox.Section(0)).
				String("pid", "PID namespace to use", ox.Section(0)).
				Int("pids-limit", "Tune container pids limit (set", ox.Section(0)).
				String("platform", "Set platform if server is", ox.Section(0)).
				Bool("privileged", "Give extended privileges to this", ox.Section(0)).
				Slice("publish", "Publish a container's port(s) to", ox.Short("p"), ox.Section(0)).
				Bool("publish-all", "Publish all exposed ports to", ox.Short("P"), ox.Section(0)).
				String("pull", "Pull image before creating", ox.Section(0)).
				Bool("quiet", "Suppress the pull output", ox.Short("q"), ox.Section(0)).
				Bool("read-only", "Mount the container's root", ox.Section(0)).
				String("restart", "Restart policy to apply when a", ox.Section(0)).
				Bool("rm", "Automatically remove the", ox.Section(0)).
				String("runtime", "Runtime to use for this container", ox.Section(0)).
				Slice("security-opt", "Security Options", ox.Section(0)).
				Slice("shm-size", "Size of /dev/shm", ox.Elem(ox.UintT), ox.Section(0)).
				String("stop-signal", "Signal to stop the container", ox.Section(0)).
				Int("stop-timeout", "Timeout (in seconds) to stop a", ox.Section(0)).
				Slice("storage-opt", "Storage driver options for the", ox.Section(0)).
				Map("sysctl", "Sysctl options", ox.Default("map[]"), ox.Section(0)).
				Slice("tmpfs", "Mount a tmpfs directory", ox.Section(0)).
				Bool("tty", "Allocate a pseudo-TTY", ox.Short("t"), ox.Section(0)).
				String("ulimit", "Ulimit options", ox.Spec("ulimit"), ox.Default("[]"), ox.Section(0)).
				Bool("use-api-socket", "Bind mount Docker API socket and", ox.Section(0)).
				String("user", "Username or UID (format:", ox.Short("u"), ox.Section(0)).
				String("userns", "User namespace to use", ox.Section(0)).
				String("uts", "UTS namespace to use", ox.Section(0)).
				Slice("volume", "Bind mount a volume", ox.Short("v"), ox.Section(0)).
				String("volume-driver", "Optional volume driver for the", ox.Section(0)).
				Slice("volumes-from", "Mount volumes from the specified", ox.Section(0)).
				String("workdir", "Working directory inside the", ox.Short("w"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("diff", "Inspect changes to files or directories on a container's filesystem"),
			ox.Banner("Inspect changes to files or directories on a container's filesystem"),
			ox.Spec("CONTAINER"),
			ox.Aliases("container diff"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("events", "Get real time events from the server"),
			ox.Banner("Get real time events from the server"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("system events"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("filter", "Filter output based on conditions provided", ox.Spec("filter"), ox.Short("f"), ox.Section(0)).
				String("format", "Format output using a custom template:", ox.Section(0)).
				String("since", "Show all events created since timestamp", ox.Section(0)).
				String("until", "Stream events until this timestamp", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("export", "Export a container's filesystem as a tar archive"),
			ox.Banner("Export a container's filesystem as a tar archive"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("container export"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("output", "Write to a file, instead of STDOUT", ox.Short("o"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("history", "Show the history of an image"),
			ox.Banner("Show the history of an image"),
			ox.Spec("[OPTIONS] IMAGE"),
			ox.Aliases("image history"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("format", "Format output using a custom template:", ox.Section(0)).
				Bool("human", "Print sizes and dates in human readable format", ox.Short("H"), ox.Section(0)).
				Bool("no-trunc", "Don't truncate output", ox.Section(0)).
				String("platform", "Show history for the given platform. Formatted", ox.Section(0)).
				Bool("quiet", "Only show image IDs", ox.Short("q"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("import", "Import the contents from a tarball to create a filesystem image"),
			ox.Banner("Import the contents from a tarball to create a filesystem image"),
			ox.Spec("[OPTIONS] file|URL|- [REPOSITORY[:TAG]]"),
			ox.Aliases("image import"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Slice("change", "Apply Dockerfile instruction to the created image", ox.Short("c"), ox.Section(0)).
				String("message", "Set commit message for imported image", ox.Short("m"), ox.Section(0)).
				String("platform", "Set platform if server is multi-platform capable", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("inspect", "Return low-level information on Docker objects"),
			ox.Banner("Return low-level information on Docker objects"),
			ox.Spec("[OPTIONS] NAME|ID [NAME|ID...]"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("format", "Format output using a custom template:", ox.Short("f"), ox.Section(0)).
				Bool("size", "Display total file sizes if the type is container", ox.Short("s"), ox.Section(0)).
				String("type", "Return JSON for specified type", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("kill", "Kill one or more running containers"),
			ox.Banner("Kill one or more running containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container kill"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("load", "Load an image from a tar archive or STDIN"),
			ox.Banner("Load an image from a tar archive or STDIN"),
			ox.Spec("[OPTIONS]"),
			ox.Aliases("image load"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("input", "Read from tar archive file, instead of STDIN", ox.Short("i"), ox.Section(0)).
				String("platform", "Load only the given platform variant. Formatted", ox.Section(0)).
				Bool("quiet", "Suppress the load output", ox.Short("q"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("logs", "Fetch the logs of a container"),
			ox.Banner("Fetch the logs of a container"),
			ox.Spec("[OPTIONS] CONTAINER"),
			ox.Aliases("container logs"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("details", "Show extra details provided to logs", ox.Section(0)).
				Bool("follow", "Follow log output", ox.Short("f"), ox.Section(0)).
				String("since", "Show logs since timestamp (e.g.", ox.Section(0)).
				String("tail", "Number of lines to show from the end of the logs", ox.Short("n"), ox.Section(0)).
				Bool("timestamps", "Show timestamps", ox.Short("t"), ox.Section(0)).
				String("until", "Show logs before a timestamp (e.g.", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("pause", "Pause all processes within one or more containers"),
			ox.Banner("Pause all processes within one or more containers"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("container pause"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("port", "List port mappings or a specific mapping for the container"),
			ox.Banner("List port mappings or a specific mapping for the container"),
			ox.Spec("CONTAINER [PRIVATE_PORT[/PROTO]]"),
			ox.Aliases("container port"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("rename", "Rename a container"),
			ox.Banner("Rename a container"),
			ox.Spec("CONTAINER NEW_NAME"),
			ox.Aliases("container rename"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("restart", "Restart one or more containers"),
			ox.Banner("Restart one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container restart"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)).
				Int("timeout", "Seconds to wait before killing the container", ox.Short("t"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("rm", "Remove one or more containers"),
			ox.Banner("Remove one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container rm", "container remove"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("force", "Force the removal of a running container (uses SIGKILL)", ox.Short("f"), ox.Section(0)).
				Bool("link", "Remove the specified link", ox.Short("l"), ox.Section(0)).
				Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("rmi", "Remove one or more images"),
			ox.Banner("Remove one or more images"),
			ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
			ox.Aliases("image rm", "image remove"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("force", "Force removal of the image", ox.Short("f"), ox.Section(0)).
				Bool("no-prune", "Do not delete untagged parents", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("save", "Save one or more images to a tar archive (streamed to STDOUT by default)"),
			ox.Banner("Save one or more images to a tar archive (streamed to STDOUT by default)"),
			ox.Spec("[OPTIONS] IMAGE [IMAGE...]"),
			ox.Aliases("image save"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("output", "Write to a file, instead of STDOUT", ox.Short("o"), ox.Section(0)).
				String("platform", "Save only the given platform variant. Formatted", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("start", "Start one or more stopped containers"),
			ox.Banner("Start one or more stopped containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container start"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("attach", "Attach STDOUT/STDERR and forward signals", ox.Short("a"), ox.Section(0)).
				String("checkpoint", "Restore from this checkpoint", ox.Section(0)).
				String("checkpoint-dir", "Use a custom checkpoint storage directory", ox.Section(0)).
				String("detach-keys", "Override the key sequence for detaching a", ox.Section(0)).
				Bool("interactive", "Attach container's STDIN", ox.Short("i"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("stats", "Display a live stream of container(s) resource usage statistics"),
			ox.Banner("Display a live stream of container(s) resource usage statistics"),
			ox.Spec("[OPTIONS] [CONTAINER...]"),
			ox.Aliases("container stats"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Bool("all", "Show all containers", ox.Default("shows just running"), ox.Short("a"), ox.Section(0)).
				String("format", "Format output using a custom template:", ox.Section(0)).
				Bool("no-stream", "Disable streaming stats and only pull the first result", ox.Section(0)).
				Bool("no-trunc", "Do not truncate output", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("stop", "Stop one or more running containers"),
			ox.Banner("Stop one or more running containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container stop"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				String("signal", "Signal to send to the container", ox.Short("s"), ox.Section(0)).
				Int("timeout", "Seconds to wait before killing the container", ox.Short("t"), ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("tag", "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
			ox.Banner("Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"),
			ox.Spec("SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]"),
			ox.Aliases("image tag"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("top", "Display the running processes of a container"),
			ox.Banner("Display the running processes of a container"),
			ox.Spec("CONTAINER [ps OPTIONS]"),
			ox.Aliases("container top"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("unpause", "Unpause all processes within one or more containers"),
			ox.Banner("Unpause all processes within one or more containers"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("container unpause"),
			ox.Section(3),
		),
		ox.Sub(
			ox.Usage("update", "Update configuration of one or more containers"),
			ox.Banner("Update configuration of one or more containers"),
			ox.Spec("[OPTIONS] CONTAINER [CONTAINER...]"),
			ox.Aliases("container update"),
			ox.Section(3),
			ox.Help(ox.Sections(
				"Options",
			)),
			ox.Flags().
				Uint16("blkio-weight", "Block IO (relative weight), between 10", ox.Section(0)).
				Int("cpu-period", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-quota", "Limit CPU CFS (Completely Fair", ox.Section(0)).
				Int("cpu-rt-period", "Limit the CPU real-time period in", ox.Section(0)).
				Int("cpu-rt-runtime", "Limit the CPU real-time runtime in", ox.Section(0)).
				Int("cpu-shares", "CPU shares (relative weight)", ox.Short("c"), ox.Section(0)).
				Float32("cpus", "Number of CPUs", ox.Spec("decimal"), ox.Section(0)).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)", ox.Section(0)).
				String("cpuset-mems", "MEMs in which to allow execution (0-3, 0,1)", ox.Section(0)).
				Slice("memory", "Memory limit", ox.Elem(ox.UintT), ox.Short("m"), ox.Section(0)).
				Slice("memory-reservation", "Memory soft limit", ox.Elem(ox.UintT), ox.Section(0)).
				Slice("memory-swap", "Swap limit equal to memory plus swap:", ox.Elem(ox.UintT), ox.Section(0)).
				Int("pids-limit", "Tune container pids limit (set -1 for", ox.Section(0)).
				String("restart", "Restart policy to apply when a", ox.Section(0)),
		),
		ox.Sub(
			ox.Usage("wait", "Block until one or more containers stop, then print their exit codes"),
			ox.Banner("Block until one or more containers stop, then print their exit codes"),
			ox.Spec("CONTAINER [CONTAINER...]"),
			ox.Aliases("container wait"),
			ox.Section(3),
		),
		ox.Flags().
			String("config", "Location of client config files (default", ox.Section(0)).
			String("context", "Name of the context to use to connect to the", ox.Short("c"), ox.Section(0)).
			Bool("debug", "Enable debug mode", ox.Short("D"), ox.Section(0)).
			Slice("host", "Daemon socket to connect to", ox.Short("H"), ox.Section(0)).
			String("log-level", "Set the logging level (\"debug\", \"info\",", ox.Short("l"), ox.Section(0)).
			Bool("tls", "Use TLS; implied by --tlsverify", ox.Section(0)).
			String("tlscacert", "Trust certs signed only by this CA (default", ox.Section(0)).
			String("tlscert", "Path to TLS certificate file (default", ox.Section(0)).
			String("tlskey", "Path to TLS key file (default", ox.Section(0)).
			Bool("tlsverify", "Use TLS and verify the remote", ox.Section(0)),
	)
}
