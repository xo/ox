// Command podman is a xo/ox version of `+podman`.
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
		ox.Banner("Manage pods, containers and images"),
		ox.Usage("podman", ""),
		ox.Spec("[options] [command]"),
		ox.Sub(
			ox.Banner("Attach to a running container\n\nDescription:\n  The podman attach command allows you to attach to a running container using the container's ID or name, either to view its ongoing output or to control it interactively."),
			ox.Usage("attach", "Attach to a running container"),
			ox.Spec("[options] CONTAINER"),
			ox.Example("\n  podman attach ctrID\n  podman attach 1234\n  podman attach --no-stdin foobar"),
			ox.Flags().
				String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-z`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("no-stdin", "Do not attach STDIN. The default is false").
				Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true")),
		), ox.Sub(
			ox.Banner("Auto update containers according to their auto-update policy\n\nDescription:\n  Auto update containers according to their auto-update policy.\n\n  Auto-update policies are specified with the \"io.containers.autoupdate\" label.\n  Containers are expected to run in systemd units created with \"podman-generate-systemd --new\",\n  or similar units that create new containers in order to run the updated images.\n  Please refer to the podman-auto-update(1) man page for details."),
			ox.Usage("auto-update", "Auto update containers according to their auto-update policy"),
			ox.Spec("[options]"),
			ox.Example("\n  podman auto-update\n  podman auto-update --authfile ~/authfile.json"),
			ox.Flags().
				String("authfile", "Path to the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				Bool("dry-run", "Check for pending updates").
				String("format", "Change the output format to JSON or a Go template").
				Bool("rollback", "Rollback to previous image if update fails", ox.Default("true")).
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
		), ox.Sub(
			ox.Banner("Build an image using instructions from Containerfiles\n\nDescription:\n  Builds an OCI or Docker image using instructions from one or more Containerfiles and a specified build context directory."),
			ox.Usage("build", "Build an image using instructions from Containerfiles"),
			ox.Spec("[options] [CONTEXT]"),
			ox.Example("\n  podman build .\n  podman build --creds=username:password -t imageName -f Containerfile.simple .\n  podman build --layers --force-rm --tag imageName ."),
			ox.Flags().
				String("add-host", "add a custom host-to-IP mapping (host:ip)", ox.Spec("host:ip"), ox.Default("[]")).
				Bool("all-platforms", "attempt to build for all base image platforms").
				Slice("annotation", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
				String("arch", "set the ARCH of the image to the provided value instead of the architecture of the host", ox.Default("amd64")).
				String("authfile", "path of the authentication file.").
				Map("build-arg", "argument=value to supply to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				String("build-arg-file", "argfile.conf containing lines of argument=value to supply to the builder", ox.Spec("argfile.conf")).
				Map("build-context", "argument=value to supply additional build context to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
				Slice("cache-from", "remote repository list to utilise as potential cache source.", ox.Elem(ox.StringT)).
				Slice("cache-to", "remote repository list to utilise as potential cache destination.", ox.Elem(ox.StringT)).
				String("cache-ttl", "only consider cache images under specified duration.").
				Slice("cap-add", "add the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
				Slice("cap-drop", "drop the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
				String("cert-dir", "use certificates at the specified path to access the registry").
				String("cgroup-parent", "optional parent cgroup for the container").
				String("cgroupns", "'private', or 'host'").
				Bool("compat-volumes", "preserve the contents of VOLUMEs during RUN instructions").
				Bool("compress", "this is a legacy option, which has no effect on the image").
				Slice("cpp-flag", "set additional flag to pass to C preprocessor (cpp)", ox.Elem(ox.StringT)).
				Uint("cpu-period", "limit the CPU CFS (Completely Fair Scheduler) period").
				Int("cpu-quota", "limit the CPU CFS (Completely Fair Scheduler) quota").
				Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
				String("creds", "use [username[:password]] for accessing the registry", ox.Spec("[username[:password]]")).
				Slice("cw", "confidential workload options", ox.Elem(ox.StringT)).
				Slice("decryption-key", "key needed to decrypt the image", ox.Elem(ox.StringT)).
				Slice("device", "additional devices to provide", ox.Elem(ox.StringT)).
				Bool("disable-compression", "don't compress layers by default", ox.Default("true"), ox.Short("D")).
				Bool("disable-content-trust", "this is a Docker specific option and is a NOOP").
				String("dns", "set custom DNS servers or disable it completely by setting it to 'none', which prevents the automatic creation of /etc/resolv.conf.", ox.Spec("path"), ox.Default("/etc/resolv.conf")).
				Slice("dns-option", "set custom DNS options", ox.Elem(ox.StringT)).
				Slice("dns-search", "set custom DNS search domains", ox.Elem(ox.StringT)).
				Slice("env", "set environment variable for the image", ox.Elem(ox.StringT)).
				String("file", "or URL                         pathname or URL of a Dockerfile", ox.Spec("pathname"), ox.Short("f")).
				Bool("force-rm", "always remove intermediate containers after a build, even if the build is unsuccessful.", ox.Default("true")).
				String("format", "format of the built image's manifest and metadata. Use BUILDAH_FORMAT environment variable to override.", ox.Spec("format"), ox.Default("oci")).
				String("from", "image name used to replace the value in the first FROM instruction in the Containerfile").
				Slice("group-add", "add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
				Slice("hooks-dir", "set the OCI hooks directory path (may be set multiple times)", ox.Elem(ox.StringT)).
				Bool("http-proxy", "pass through HTTP Proxy environment variables", ox.Default("true")).
				Bool("identity-label", "add default identity label", ox.Default("true")).
				String("ignorefile", "path to an alternate .dockerignore file").
				String("iidfile", "file to write the image ID to", ox.Spec("file")).
				String("ipc", "'private', path of IPC namespace to join, or 'host'", ox.Spec("path")).
				String("isolation", "type of process isolation to use. Use BUILDAH_ISOLATION environment variable to override.", ox.Spec("type"), ox.Default("rootless")).
				Int("jobs", "how many stages to run in parallel", ox.Default("1")).
				Slice("label", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
				Slice("layer-label", "set metadata for an intermediate image", ox.Elem(ox.StringT), ox.Default("[]")).
				Bool("layers", "use intermediate layers during build. Use BUILDAH_LAYERS environment variable to override.", ox.Default("true")).
				String("logfile", "log to file instead of stdout/stderr", ox.Spec("file")).
				Bool("logsplit", "split logfile to different files for each platform").
				String("manifest", "add the image to the specified manifest list. Creates manifest list if it does not exist").
				String("memory", "memory limit (format: <number>[<unit>], where unit = b, k, m or g)", ox.Short("m")).
				String("memory-swap", "swap limit equal to memory plus swap: '-1' to enable unlimited swap").
				String("network", "'private', 'none', 'ns:path' of network namespace to join, or 'host'").
				Bool("no-cache", "do not use existing cached images for the container build. Build from the start with a new set of cached layers.").
				Bool("no-hostname", "do not create new /etc/hostname file for RUN instructions, use the one from the base image.").
				Bool("no-hosts", "do not create new /etc/hosts file for RUN instructions, use the one from the base image.").
				Bool("omit-history", "omit build history information from built image").
				String("os", "set the OS to the provided value instead of the current operating system of the host", ox.Default("linux")).
				String("os-feature", "set required OS feature for the target image in addition to values from the base image", ox.Spec("feature")).
				String("os-version", "set required OS version for the target image instead of the value from the base image", ox.Spec("version")).
				String("output", "output destination (format: type=local,dest=path)", ox.Short("o")).
				String("pid", "private, path of PID namespace to join, or 'host'", ox.Spec("path")).
				String("platform", "set the OS/ARCH[/VARIANT] of the image to the provided value instead of the current operating system and architecture of the host (for example \"linux/arm\")", ox.Spec("OS/ARCH[/VARIANT]"), ox.Default("[linux/amd64]")).
				Map("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Spec("string[=\"missing\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("missing")).
				Bool("quiet", "refrain from announcing build instructions and image read/write progress", ox.Short("q")).
				Int("retry", "number of times to retry in case of failure when performing push/pull", ox.Default("3")).
				String("retry-delay", "delay between retries in case of push/pull failures").
				Bool("rm", "remove intermediate containers after a successful build", ox.Default("true")).
				Slice("runtime-flag", "add global flags for the container runtime", ox.Elem(ox.StringT)).
				String("sbom", "scan working container using preset configuration", ox.Spec("preset")).
				String("sbom-image-output", "add scan results to image as path", ox.Spec("path")).
				String("sbom-image-purl-output", "add scan results to image as path", ox.Spec("path")).
				String("sbom-merge-strategy", "merge scan results using strategy", ox.Spec("strategy")).
				String("sbom-output", "save scan results to file", ox.Spec("file")).
				String("sbom-purl-output", "save scan results to file`", ox.Spec("file")).
				String("sbom-scanner-command", "scan working container using command in scanner image", ox.Spec("command")).
				String("sbom-scanner-image", "scan working container using scanner command from image", ox.Spec("image")).
				Slice("secret", "secret file to expose to the build", ox.Elem(ox.StringT)).
				Slice("security-opt", "security options", ox.Elem(ox.StringT), ox.Default("[]")).
				String("shm-size", "size of '/dev/shm'. The format is <number><unit>.", ox.Spec("<number><unit>"), ox.Default("65536k")).
				String("sign-by", "sign the image using a GPG key with the specified FINGERPRINT", ox.Spec("FINGERPRINT")).
				Bool("skip-unused-stages", "skips stages in multi-stage builds which do not affect the final target", ox.Default("true")).
				Bool("squash", "squash all image layers into a single layer").
				Bool("squash-all", "Squash all layers into a single layer").
				Slice("ssh", "SSH agent socket or keys to expose to the build. (format: default|<id>[=<socket>|<key>[,<key>]])", ox.Elem(ox.StringT)).
				Bool("stdin", "pass stdin into containers").
				String("tag", "tagged name to apply to the built image", ox.Spec("name"), ox.Short("t")).
				String("target", "set the target build stage to build").
				Int("timestamp", "set created timestamp to the specified epoch seconds to allow for deterministic builds, defaults to current time").
				Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")).
				Slice("ulimit", "ulimit options", ox.Elem(ox.StringT)).
				Slice("unsetenv", "unset environment variable from final image", ox.Elem(ox.StringT)).
				Slice("unsetlabel", "unset label when inheriting labels from base image", ox.Elem(ox.StringT)).
				String("userns", "'container', path of user namespace to join, or 'host'", ox.Spec("path")).
				String("userns-gid-map", "containerGID:hostGID:length GID mapping to use in user namespace", ox.Spec("containerGID:hostGID:length")).
				String("userns-gid-map-group", "name of entries from /etc/subgid to use to set user namespace GID mapping", ox.Spec("name")).
				String("userns-uid-map", "containerUID:hostUID:length UID mapping to use in user namespace", ox.Spec("containerUID:hostUID:length")).
				String("userns-uid-map-user", "name of entries from /etc/subuid to use to set user namespace UID mapping", ox.Spec("name")).
				String("uts", "private, :path of UTS namespace to join, or 'host'", ox.Spec("path")).
				String("variant", "override the variant of the specified image", ox.Spec("variant")).
				Slice("volume", "bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")),
		), ox.Sub(
			ox.Banner("Create new image based on the changed container\n\nDescription:\n  Create an image from a container's changes. Optionally tag the image created, set the author with the --author flag, set the commit message with the --message flag, and make changes to the instructions with the --change flag."),
			ox.Usage("commit", "Create new image based on the changed container"),
			ox.Spec("[options] CONTAINER [IMAGE]"),
			ox.Example("\n  podman commit -q --message \"committing container to image\" reverent_golick image-committed\n  podman commit -q --author \"firstName lastName\" reverent_golick image-committed\n  podman commit -q --pause=false containerID image-committed\n  podman commit containerID"),
			ox.Flags().
				String("author", "Set the author for the image committed", ox.Short("a")).
				Slice("change", "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR", ox.Elem(ox.StringT), ox.Short("c")).
				String("config", "file containing a container configuration to merge into the image", ox.Spec("file")).
				String("format", "Format of the image manifest and metadata", ox.Spec("Format"), ox.Default("oci"), ox.Short("f")).
				String("iidfile", "file to write the image ID to", ox.Spec("file")).
				Bool("include-volumes", "Include container volumes as image volumes").
				String("message", "Set commit message for imported image", ox.Short("m")).
				Bool("pause", "Pause container during commit", ox.Short("p")).
				Bool("quiet", "Suppress output", ox.Short("q")).
				Bool("squash", "squash newly built layers into a single new layer", ox.Short("s")),
		), ox.Sub(
			ox.Banner("Run compose workloads via an external provider such as docker-compose or podman-compose\n\nDescription:\n  This command is a thin wrapper around an external compose provider such as docker-compose or podman-compose.  This means that podman compose is executing another tool that implements the compose functionality but sets up the environment in a way to let the compose provider communicate transparently with the local Podman socket.  The specified options as well the command and argument are passed directly to the compose provider.\n\nThe default compose providers are docker-compose and podman-compose.  If installed, docker-compose takes precedence since it is the original implementation of the Compose specification and is widely used on the supported platforms (i.e., Linux, Mac OS, Windows).\n\nIf you want to change the default behavior or have a custom installation path for your provider of choice, please change the compose_provider field in containers.conf(5).  You may also set PODMAN_COMPOSE_PROVIDER environment variable."),
			ox.Usage("compose", "Run compose workloads via an external provider such as docker-compose or podman-compose"),
			ox.Spec("[options]"),
			ox.Example("\n  podman compose -f nginx.yaml up --detach\n  podman --log-level=debug compose -f many-images.yaml pull"),
		), ox.Sub(
			ox.Banner("Manage containers\n\nDescription:\n  Manage containers"),
			ox.Usage("container", "Manage containers"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Attach to a running container\n\nDescription:\n  The podman attach command allows you to attach to a running container using the container's ID or name, either to view its ongoing output or to control it interactively."),
				ox.Usage("attach", "Attach to a running container"),
				ox.Spec("[options] CONTAINER"),
				ox.Example("\n  podman container attach ctrID\n\tpodman container attach 1234\n\tpodman container attach --no-stdin foobar"),
				ox.Flags().
					String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-z`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-stdin", "Do not attach STDIN. The default is false").
					Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Checkpoint one or more containers\n\nDescription:\n  \n   podman container checkpoint\n\n   Checkpoints one or more running containers. The container name or ID can be used."),
				ox.Usage("checkpoint", "Checkpoint one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container checkpoint --keep ctrID\n  podman container checkpoint --all\n  podman container checkpoint --leave-running ctrID"),
				ox.Flags().
					Bool("all", "Checkpoint all running containers", ox.Short("a")).
					String("compress", "Select compression algorithm (gzip, none, zstd) for checkpoint archive.", ox.Default("zstd"), ox.Short("c")).
					String("create-image", "Create checkpoint image with specified name").
					String("export", "Export the checkpoint image to a tar.gz", ox.Short("e")).
					Bool("file-locks", "Checkpoint a container with file locks").
					Bool("ignore-rootfs", "Do not include root file-system changes when exporting").
					Bool("ignore-volumes", "Do not export volumes associated with container").
					Bool("keep", "Keep all temporary checkpoint files", ox.Short("k")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("leave-running", "Leave the container running after writing checkpoint to disk", ox.Short("R")).
					Bool("pre-checkpoint", "Dump container's memory information only, leave the container running", ox.Short("P")).
					Bool("print-stats", "Display checkpoint statistics").
					Bool("tcp-established", "Checkpoint a container with established TCP connections").
					Bool("with-previous", "Checkpoint container with pre-checkpoint images"),
			), ox.Sub(
				ox.Banner("Clean up network and mountpoints of one or more containers\n\nDescription:\n  \n   podman container cleanup\n\n   Cleans up mount points and network stacks on one or more containers from the host. The container name or ID can be used. This command is used internally when running containers, but can also be used if container cleanup has failed when a container exits."),
				ox.Usage("cleanup", "Clean up network and mountpoints of one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container cleanup ctrID1 ctrID2 ctrID3\n  podman container cleanup --all"),
				ox.Flags().
					Bool("all", "Cleans up all containers", ox.Short("a")).
					String("exec", "Clean up the given exec session instead of the container").
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("rm", "After cleanup, remove the container entirely").
					Bool("rmi", "After cleanup, remove the image entirely"),
			), ox.Sub(
				ox.Banner("Clone an existing container\n\nDescription:\n  Creates a copy of an existing container."),
				ox.Usage("clone", "Clone an existing container"),
				ox.Spec("[options] CONTAINER NAME IMAGE"),
				ox.Example("\n  podman container clone container_name new_name image_name"),
				ox.Flags().
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Bool("destroy", "destroy the original container").
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Bool("force", "force the existing container to be destroyed", ox.Short("f")).
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
					String("name", "Assign a name to the container").
					String("pod", "Run container in an existing pod").
					Bool("run", "run the new container"),
			), ox.Sub(
				ox.Banner("Create new image based on the changed container\n\nDescription:\n  Create an image from a container's changes. Optionally tag the image created, set the author with the --author flag, set the commit message with the --message flag, and make changes to the instructions with the --change flag."),
				ox.Usage("commit", "Create new image based on the changed container"),
				ox.Spec("[options] CONTAINER [IMAGE]"),
				ox.Example("\n  podman container commit -q --message \"committing container to image\" reverent_golick image-committed\n  podman container commit -q --author \"firstName lastName\" reverent_golick image-committed\n  podman container commit -q --pause=false containerID image-committed\n  podman container commit containerID"),
				ox.Flags().
					String("author", "Set the author for the image committed", ox.Short("a")).
					Slice("change", "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR", ox.Elem(ox.StringT), ox.Short("c")).
					String("config", "file containing a container configuration to merge into the image", ox.Spec("file")).
					String("format", "Format of the image manifest and metadata", ox.Spec("Format"), ox.Default("oci"), ox.Short("f")).
					String("iidfile", "file to write the image ID to", ox.Spec("file")).
					Bool("include-volumes", "Include container volumes as image volumes").
					String("message", "Set commit message for imported image", ox.Short("m")).
					Bool("pause", "Pause container during commit", ox.Short("p")).
					Bool("quiet", "Suppress output", ox.Short("q")).
					Bool("squash", "squash newly built layers into a single new layer", ox.Short("s")),
			), ox.Sub(
				ox.Banner("Copy files/folders between a container and the local filesystem\n\nDescription:\n  Copy the contents of SRC_PATH to the DEST_PATH.\n\n  You can copy from the container's file system to the local machine or the reverse, from the local filesystem to the container. If \"-\" is specified for either the SRC_PATH or DEST_PATH, you can also stream a tar archive from STDIN or to STDOUT. The CONTAINER can be a running or stopped container. The SRC_PATH or DEST_PATH can be a file or a directory."),
				ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
				ox.Spec("[options] [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH"),
				ox.Example("\n  podman container cp [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH"),
				ox.Flags().
					Bool("archive", "Chown copied files to the primary uid/gid of the destination container.", ox.Default("true"), ox.Short("a")).
					Bool("overwrite", "Allow to overwrite directories with non-directories and vice versa"),
			), ox.Sub(
				ox.Banner("Create but do not start a container\n\nDescription:\n  Creates a new container from the given image or storage and prepares it for running the specified command.\n\n  The container ID is then printed to stdout. You can then start it at any time with the podman start <container_id> command. The container will be created with the initial state 'created'."),
				ox.Usage("create", "Create but do not start a container"),
				ox.Spec("[options] IMAGE [COMMAND [ARG...]]"),
				ox.Example("\n  podman container create alpine ls\n  podman container create --annotation HELLO=WORLD alpine ls\n  podman container create -t -i --name myctr alpine ls"),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (host:ip)", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("annotation", "Add annotations to container (key=value)", ox.Elem(ox.StringT)).
					String("arch", "use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					Slice("cap-add", "Add capabilities to the container", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop capabilities from the container", ox.Elem(ox.StringT)).
					Slice("cgroup-conf", "Configure cgroup v2 (key=value)", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Optional parent cgroup for the container").
					String("cgroupns", "cgroup namespace to use").
					String("cgroups", "control container cgroup configuration (\"enabled\"|\"disabled\"|\"no-conmon\"|\"split\")", ox.Default("enabled")).
					Slice("chrootdirs", "Chroot directories inside the container", ox.Elem(ox.StringT)).
					String("cidfile", "Write the container ID to the file").
					String("conmon-pidfile", "Path to the file that will receive the PID of conmon").
					Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed devices list", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set custom DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("entrypoint", "Overwrite the default ENTRYPOINT of the image").
					Slice("env", "Set environment variables in container", ox.Elem(ox.StringT), ox.Default("[PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]"), ox.Short("e")).
					Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
					Bool("env-host", "Use all current host environment variables in container").
					Slice("env-merge", "Preprocess environment variables from image before injecting them into the container", ox.Elem(ox.StringT)).
					Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
					Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
					Slice("group-add", "Add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
					String("group-entry", "Entry to write to /etc/group").
					String("health-cmd", "set a healthcheck command for the container ('none' disables the existing healthcheck)").
					String("health-interval", "set an interval for the healthcheck (a value of disable results in no automatic timer setup)", ox.Default("30s")).
					String("health-log-destination", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file)", ox.Default("local")).
					Uint("health-max-log-count", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)", ox.Default("5")).
					Uint("health-max-log-size", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)", ox.Default("500")).
					String("health-on-failure", "action to take once the container turns unhealthy", ox.Default("none")).
					Uint("health-retries", "the number of retries allowed before a healthcheck is considered to be unhealthy", ox.Default("3")).
					String("health-start-period", "the initialization time needed for a container to bootstrap", ox.Default("0s")).
					String("health-startup-cmd", "Set a startup healthcheck command for the container").
					String("health-startup-interval", "Set an interval for the startup healthcheck", ox.Default("30s")).
					Uint("health-startup-retries", "Set the maximum number of retries before the startup healthcheck will restart the container").
					Uint("health-startup-success", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)").
					String("health-startup-timeout", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed", ox.Default("30s")).
					String("health-timeout", "the maximum time allowed to complete the healthcheck before an interval is considered failed", ox.Default("30s")).
					Slice("hostuser", "Host user account to add to /etc/passwd within container", ox.Elem(ox.StringT)).
					Bool("http-proxy", "Set proxy environment variables in the container based on the host proxy vars", ox.Default("true")).
					String("image-volume", "Tells podman how to handle the builtin image volumes (\"bind\"|\"tmpfs\"|\"ignore\")", ox.Default("anonymous")).
					Bool("init", "Run an init binary inside the container that forwards signals and reaps processes").
					String("init-ctr", "Make this a pod init container.").
					String("init-path", "Path to the container-init binary").
					Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
					String("ip", "Specify a static IPv4 address for the container").
					String("ip6", "Specify a static IPv6 address for the container").
					String("ipc", "IPC namespace to use").
					Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
					String("log-driver", "Logging driver for the container", ox.Default("journald")).
					Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
					String("mac-address", "Container MAC address (e.g. 92:d0:c6:0a:29:33)").
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
					Slice("mount", "Attach a filesystem mount to the container", ox.Elem(ox.StringT)).
					String("name", "Assign a name to the container").
					Slice("network", "Connect a container to a network", ox.Elem(ox.StringT)).
					Slice("network-alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
					Bool("no-healthcheck", "Disable healthchecks on container").
					Bool("no-hosts", "Do not create /etc/hosts within the container, instead use the version from the image").
					Bool("oom-kill-disable", "Disable OOM Killer").
					Int("oom-score-adj", "Tune the host's OOM preferences (-1000 to 1000)").
					String("os", "use OS instead of the running OS for choosing images", ox.Spec("OS")).
					String("passwd-entry", "Entry to write to /etc/passwd").
					String("personality", "Configure execution domain using personality (e.g., LINUX/LINUX32)").
					String("pid", "PID namespace to use").
					String("pidfile", "Write the container process ID to the file").
					Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
					String("platform", "Specify the platform for selecting the image.  (Conflicts with --arch and --os)").
					String("pod", "Run container in an existing pod").
					String("pod-id-file", "Read the pod ID from the file").
					Bool("privileged", "Give extended privileges to container").
					Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
					Bool("publish-all", "Publish all exposed ports to random ports on the host interface", ox.Short("P")).
					String("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Default("missing")).
					Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
					String("rdt-class", "Class of Service (COS) that the container should be assigned to").
					Bool("read-only", "Make containers root filesystem read-only").
					Bool("read-only-tmpfs", "When running --read-only containers mount read-write tmpfs on /dev, /dev/shm, /run, /tmp and /var/tmp", ox.Default("true")).
					Bool("replace", "If a container with the same name exists, replace it").
					Slice("requires", "Add one or more requirement containers that must be started before this container will start", ox.Elem(ox.StringT)).
					String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
					Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
					String("retry-delay", "delay between retries in case of pull failures").
					Bool("rm", "Remove container and any anonymous unnamed volume associated with the container after exit").
					Bool("rootfs", "The first argument is not an image but the rootfs to the exploded container").
					String("sdnotify", "control sd-notify behavior (\"container\"|\"conmon\"|\"healthy\"|\"ignore\")", ox.Default("container")).
					String("seccomp-policy", "Policy for selecting a seccomp profile (experimental)", ox.Default("default")).
					Slice("secret", "Add secret to container", ox.Elem(ox.StringT)).
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
					String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("stop-signal", "Signal to stop a container. Default is SIGTERM").
					Uint("stop-timeout", "Timeout (in seconds) that containers stopped by user command have to exit. If exceeded, the container will be forcibly stopped via SIGKILL.", ox.Default("10")).
					String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
					String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
					Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
					String("systemd", "Run container in systemd mode (\"true\"|\"false\"|\"always\")", ox.Default("true")).
					Uint("timeout", "Maximum length of time a container is allowed to run. The container will be killed automatically after the time expires.").
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries for pulling images").
					String("tmpfs", "Mount a temporary filesystem (tmpfs) into a container", ox.Spec("tmpfs")).
					Bool("tty", "Allocate a pseudo-TTY for container", ox.Short("t")).
					String("tz", "Set timezone in container").
					Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("ulimit", "Ulimit options", ox.Elem(ox.StringT)).
					String("umask", "Set umask in container", ox.Default("0022")).
					Slice("unsetenv", "Unset environment default variables in container", ox.Elem(ox.StringT)).
					Bool("unsetenv-all", "Unset all default environment variables in container").
					String("user", "Username or UID (format: <name|uid>[:<group|gid>])", ox.Short("u")).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					String("variant", "Use VARIANT instead of the running architecture variant for choosing images", ox.Spec("VARIANT")).
					Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
					Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)).
					String("workdir", "Working directory inside the container", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Inspect changes to the container's file systems\n\nDescription:\n  Displays changes to the container filesystem's'.  The container will be compared to its parent layer or the second argument when given."),
				ox.Usage("diff", "Inspect changes to the container's file systems"),
				ox.Spec("[options] CONTAINER [CONTAINER]"),
				ox.Example("\n  podman container diff myCtr\n  podman container diff -l --format json myCtr"),
				ox.Flags().
					String("format", "Change the output format (json)").
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Run a process in a running container\n\nDescription:\n  Execute the specified command inside a running container."),
				ox.Usage("exec", "Run a process in a running container"),
				ox.Spec("[options] CONTAINER COMMAND [ARG...]"),
				ox.Example("\n  podman container exec -it ctrID ls\n  podman container exec -it -w /tmp myCtr pwd\n  podman container exec --user root ctrID ls"),
				ox.Flags().
					Bool("detach", "Run the exec session in detached mode (backgrounded)", ox.Short("d")).
					String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _", ox.Default("ctrl-p,ctrl-q")).
					Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
					Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
					Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Slice("preserve-fd", "Pass a list of additional file descriptors to the container", ox.Elem(ox.UintT), ox.Default("[]")).
					Uint("preserve-fds", "Pass N additional file descriptors to the container").
					Bool("privileged", "Give the process extended Linux capabilities inside the container.  The default is false").
					Bool("tty", "Allocate a pseudo-TTY. The default is false", ox.Short("t")).
					String("user", "Sets the username or UID used and optionally the groupname or GID for the specified command", ox.Short("u")).
					String("workdir", "Working directory inside the container", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Check if a container exists in local storage\n\nDescription:\n  If the named container exists in local storage, podman container exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if a container exists in local storage"),
				ox.Spec("[options] CONTAINER"),
				ox.Example("\n  podman container exists --external containerID\n  podman container exists myctr || podman run --name myctr [etc...]"),
				ox.Flags().
					Bool("external", "Check external storage containers as well as Podman containers"),
			), ox.Sub(
				ox.Banner("Export container's filesystem contents as a tar archive\n\nDescription:\n  Exports container's filesystem contents as a tar archive and saves it on the local machine."),
				ox.Usage("export", "Export container's filesystem contents as a tar archive"),
				ox.Spec("[options] CONTAINER"),
				ox.Example("\n  podman container export ctrID > myCtr.tar\n  podman container export --output=\"myCtr.tar\" ctrID"),
				ox.Flags().
					String("output", "Write to a specified file (default: stdout, which must be redirected)", ox.Short("o")),
			), ox.Sub(
				ox.Banner("Initialize one or more containers\n\nDescription:\n  Initialize one or more containers, creating the OCI spec and mounts for inspection. Container names or IDs can be used."),
				ox.Usage("init", "Initialize one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container init 3c45ef19d893\n  podman container init test1"),
				ox.Flags().
					Bool("all", "Initialize all containers", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Display the configuration of a container\n\nDescription:\n  Displays the low-level information on a container identified by name or ID."),
				ox.Usage("inspect", "Display the configuration of a container"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container inspect myCtr\n  podman container inspect -l --format '{{.Id}} {{.Config.Labels}}'"),
				ox.Flags().
					String("format", "Format the output to a Go template or json", ox.Default("json"), ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("size", "Display total file size", ox.Short("s")),
			), ox.Sub(
				ox.Banner("Kill one or more running containers with a specific signal\n\nDescription:\n  The main process inside each container specified will be sent SIGKILL, or any signal specified with option --signal."),
				ox.Usage("kill", "Kill one or more running containers with a specific signal"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container kill mywebserver\n  podman container kill 860a4b23\n  podman container kill --signal TERM ctrID"),
				ox.Flags().
					Bool("all", "Signal all running containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					String("signal", "Signal to send to the container", ox.Default("KILL"), ox.Short("s")),
			), ox.Sub(
				ox.Banner("List containers\n\nDescription:\n  Prints out information about the containers"),
				ox.Usage("list", "List containers"),
				ox.Spec("[options]"),
				ox.Aliases("list", "ls"),
				ox.Example("\n  podman container list -a\n  podman container list -a --format \"{{.ID}}  {{.Image}}  {{.Labels}}  {{.Mounts}}\"\n  podman container list --size --sort names"),
				ox.Flags().
					Bool("all", "Show all the containers, default is only running containers", ox.Short("a")).
					Bool("external", "Show containers in storage not controlled by Podman").
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Pretty-print containers to JSON or using a Go template").
					Int("last", "Print the n last created containers (all states)", ox.Default("-1"), ox.Short("n")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-trunc", "Display the extended information").
					Bool("noheading", "Do not print headers").
					Bool("ns", "Display namespace information").
					Bool("pod", "Print the ID and name of the pod the containers are associated with", ox.Short("p")).
					Bool("quiet", "Print the numeric IDs of the containers only", ox.Short("q")).
					Bool("size", "Display the total file sizes", ox.Short("s")).
					String("sort", "Sort output by: command, created, id, image, names, runningfor, size, status", ox.Spec("choice")).
					Bool("sync", "Sync container state with OCI runtime").
					Uint("watch", "Watch the ps output on an interval in seconds", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Fetch the logs of one or more containers\n\nDescription:\n  Retrieves logs for one or more containers.\n\n  This does not guarantee execution order when combined with podman run (i.e., your run may not have generated any logs at the time you execute podman logs)."),
				ox.Usage("logs", "Fetch the logs of one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container logs ctrID\n\t\tpodman container logs --names ctrID1 ctrID2\n\t\tpodman container logs --color --names ctrID1 ctrID2\n\t\tpodman container logs --tail 2 mywebserver\n\t\tpodman container logs --follow=true --since 10m ctrID\n\t\tpodman container logs mywebserver mydbserver"),
				ox.Flags().
					Bool("color", "Output the containers with different colors in the log.").
					Bool("follow", "Follow log output.  The default is false", ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("names", "Output the container name in the log", ox.Short("n")).
					String("since", "Show logs since TIMESTAMP").
					Int("tail", "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines", ox.Default("-1")).
					Bool("timestamps", "Output the timestamps in the log", ox.Short("t")).
					String("until", "Show logs until TIMESTAMP"),
			), ox.Sub(
				ox.Banner("Mount a working container's root filesystem\n\nDescription:\n  podman mount\n    Lists all mounted containers mount points if no container is specified\n\n  podman mount CONTAINER-NAME-OR-ID\n    Mounts the specified container and outputs the mountpoint"),
				ox.Usage("mount", "Mount a working container's root filesystem"),
				ox.Spec("[options] [CONTAINER...]"),
				ox.Flags().
					Bool("all", "Mount all containers", ox.Short("a")).
					String("format", "Print the mounted containers in specified format (json)").
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-trunc", "Do not truncate output"),
			), ox.Sub(
				ox.Banner("Pause all the processes in one or more containers\n\nDescription:\n  Pauses one or more running containers.  The container name or ID can be used."),
				ox.Usage("pause", "Pause all the processes in one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container pause mywebserver\n  podman container pause 860a4b23\n  podman container pause --all"),
				ox.Flags().
					Bool("all", "Pause all running containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("List port mappings or a specific mapping for the container\n\nDescription:\n  List port mappings for the CONTAINER, or look up the public-facing port that is NAT-ed to the PRIVATE_PORT"),
				ox.Usage("port", "List port mappings or a specific mapping for the container"),
				ox.Spec("[options] CONTAINER [PORT]"),
				ox.Example("\n  podman container port --all\n  podman container port CTRID 80"),
				ox.Flags().
					Bool("all", "Display port information for all containers", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Remove all non running containers\n\nDescription:\n  podman container prune\n\n\tRemoves all non running containers"),
				ox.Usage("prune", "Remove all non running containers"),
				ox.Spec("[options]"),
				ox.Example("\n  podman container prune"),
				ox.Flags().
					Slice("filter", "Provide filter values (e.g. 'label=<key>=<value>')", ox.Elem(ox.StringT)).
					Bool("force", "Do not prompt for confirmation.  The default is false", ox.Short("f")),
			), ox.Sub(
				ox.Banner("List containers\n\nDescription:\n  Prints out information about the containers"),
				ox.Usage("ps", "List containers"),
				ox.Spec("[options]"),
				ox.Example("\n  podman container ps -a\n  podman container ps -a --format \"{{.ID}}  {{.Image}}  {{.Labels}}  {{.Mounts}}\"\n  podman container ps --size --sort names"),
				ox.Flags().
					Bool("all", "Show all the containers, default is only running containers", ox.Short("a")).
					Bool("external", "Show containers in storage not controlled by Podman").
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Pretty-print containers to JSON or using a Go template").
					Int("last", "Print the n last created containers (all states)", ox.Default("-1"), ox.Short("n")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-trunc", "Display the extended information").
					Bool("noheading", "Do not print headers").
					Bool("ns", "Display namespace information").
					Bool("pod", "Print the ID and name of the pod the containers are associated with", ox.Short("p")).
					Bool("quiet", "Print the numeric IDs of the containers only", ox.Short("q")).
					Bool("size", "Display the total file sizes", ox.Short("s")).
					String("sort", "Sort output by: command, created, id, image, names, runningfor, size, status", ox.Spec("choice")).
					Bool("sync", "Sync container state with OCI runtime").
					Uint("watch", "Watch the ps output on an interval in seconds", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Rename an existing container\n\nDescription:\n  The podman rename command allows you to rename an existing container"),
				ox.Usage("rename", "Rename an existing container"),
				ox.Spec("CONTAINER NAME"),
				ox.Example("\n  podman container rename containerA newName"),
			), ox.Sub(
				ox.Banner("Restart one or more containers\n\nDescription:\n  Restarts one or more running containers. The container ID or name can be used.\n\n  A timeout before forcibly stopping can be set, but defaults to 10 seconds."),
				ox.Usage("restart", "Restart one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container restart ctrID\n  podman container restart ctrID1 ctrID2"),
				ox.Flags().
					Bool("all", "Restart all non-running containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("running", "Restart only running containers").
					Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Restore one or more containers from a checkpoint\n\nDescription:\n  \n   podman container restore\n\n   Restores a container from a checkpoint. The container name or ID can be used."),
				ox.Usage("restore", "Restore one or more containers from a checkpoint"),
				ox.Spec("[options] CONTAINER|IMAGE [CONTAINER|IMAGE...]"),
				ox.Example("\n  podman container restore ctrID\n  podman container restore imageID\n  podman container restore --all"),
				ox.Flags().
					Bool("all", "Restore all checkpointed containers", ox.Short("a")).
					Bool("file-locks", "Restore a container with file locks").
					Bool("ignore-rootfs", "Do not apply root file-system changes when importing from exported checkpoint").
					Bool("ignore-static-ip", "Ignore IP address set via --static-ip").
					Bool("ignore-static-mac", "Ignore MAC address set via --mac-address").
					Bool("ignore-volumes", "Do not export volumes associated with container").
					String("import", "Restore from exported checkpoint archive (tar.gz)", ox.Short("i")).
					String("import-previous", "Restore from exported pre-checkpoint archive (tar.gz)").
					Bool("keep", "Keep all temporary checkpoint files", ox.Short("k")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					String("name", "Specify new name for container restored from exported checkpoint (only works with image or --import)", ox.Short("n")).
					String("pod", "Restore container into existing Pod (only works with image or --import)").
					Bool("print-stats", "Display restore statistics").
					Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
					Bool("tcp-established", "Restore a container with established TCP connections"),
			), ox.Sub(
				ox.Banner("Remove one or more containers\n\nDescription:\n  Removes one or more containers from the host. The container name or ID can be used.\n\n  Command does not remove images. Running or unusable containers will not be removed without the -f option."),
				ox.Usage("rm", "Remove one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container rm imageID\n  podman container rm mywebserver myflaskserver 860a4b23\n  podman container rm --force --all\n  podman container rm -f c684f0d469f2"),
				ox.Flags().
					Bool("all", "Remove all containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Bool("depend", "Remove container and all containers that depend on the selected container").
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT)).
					Bool("force", "Force removal of a running or unusable container", ox.Short("f")).
					Bool("ignore", "Ignore errors when a specified container is missing", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")).
					Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v")),
			), ox.Sub(
				ox.Banner("Run a command in a new container\n\nDescription:\n  Runs a command in a new container from the given image"),
				ox.Usage("run", "Run a command in a new container"),
				ox.Spec("[options] IMAGE [COMMAND [ARG...]]"),
				ox.Example("\n  podman container run imageID ls -alF /etc\n\tpodman container run --network=host imageID dnf -y install java\n\tpodman container run --volume /var/hostdir:/var/ctrdir -i -t fedora /bin/bash"),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (host:ip)", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("annotation", "Add annotations to container (key=value)", ox.Elem(ox.StringT)).
					String("arch", "use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
					Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					Slice("cap-add", "Add capabilities to the container", ox.Elem(ox.StringT)).
					Slice("cap-drop", "Drop capabilities from the container", ox.Elem(ox.StringT)).
					Slice("cgroup-conf", "Configure cgroup v2 (key=value)", ox.Elem(ox.StringT)).
					String("cgroup-parent", "Optional parent cgroup for the container").
					String("cgroupns", "cgroup namespace to use").
					String("cgroups", "control container cgroup configuration (\"enabled\"|\"disabled\"|\"no-conmon\"|\"split\")", ox.Default("enabled")).
					Slice("chrootdirs", "Chroot directories inside the container", ox.Elem(ox.StringT)).
					String("cidfile", "Write the container ID to the file").
					String("conmon-pidfile", "Path to the file that will receive the PID of conmon").
					Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
					Bool("detach", "Run container in background and print container ID", ox.Short("d")).
					String("detach-keys", "Override the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-cf`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-cgroup-rule", "Add a rule to the cgroup allowed devices list", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set custom DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("entrypoint", "Overwrite the default ENTRYPOINT of the image").
					Slice("env", "Set environment variables in container", ox.Elem(ox.StringT), ox.Default("[PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]"), ox.Short("e")).
					Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
					Bool("env-host", "Use all current host environment variables in container").
					Slice("env-merge", "Preprocess environment variables from image before injecting them into the container", ox.Elem(ox.StringT)).
					Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
					Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
					Slice("group-add", "Add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
					String("group-entry", "Entry to write to /etc/group").
					String("health-cmd", "set a healthcheck command for the container ('none' disables the existing healthcheck)").
					String("health-interval", "set an interval for the healthcheck (a value of disable results in no automatic timer setup)", ox.Default("30s")).
					String("health-log-destination", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file)", ox.Default("local")).
					Uint("health-max-log-count", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)", ox.Default("5")).
					Uint("health-max-log-size", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)", ox.Default("500")).
					String("health-on-failure", "action to take once the container turns unhealthy", ox.Default("none")).
					Uint("health-retries", "the number of retries allowed before a healthcheck is considered to be unhealthy", ox.Default("3")).
					String("health-start-period", "the initialization time needed for a container to bootstrap", ox.Default("0s")).
					String("health-startup-cmd", "Set a startup healthcheck command for the container").
					String("health-startup-interval", "Set an interval for the startup healthcheck", ox.Default("30s")).
					Uint("health-startup-retries", "Set the maximum number of retries before the startup healthcheck will restart the container").
					Uint("health-startup-success", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)").
					String("health-startup-timeout", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed", ox.Default("30s")).
					String("health-timeout", "the maximum time allowed to complete the healthcheck before an interval is considered failed", ox.Default("30s")).
					Slice("hostuser", "Host user account to add to /etc/passwd within container", ox.Elem(ox.StringT)).
					Bool("http-proxy", "Set proxy environment variables in the container based on the host proxy vars", ox.Default("true")).
					String("image-volume", "Tells podman how to handle the builtin image volumes (\"bind\"|\"tmpfs\"|\"ignore\")", ox.Default("anonymous")).
					Bool("init", "Run an init binary inside the container that forwards signals and reaps processes").
					String("init-path", "Path to the container-init binary").
					Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
					String("ip", "Specify a static IPv4 address for the container").
					String("ip6", "Specify a static IPv6 address for the container").
					String("ipc", "IPC namespace to use").
					Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
					String("log-driver", "Logging driver for the container", ox.Default("journald")).
					Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
					String("mac-address", "Container MAC address (e.g. 92:d0:c6:0a:29:33)").
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
					Slice("mount", "Attach a filesystem mount to the container", ox.Elem(ox.StringT)).
					String("name", "Assign a name to the container").
					Slice("network", "Connect a container to a network", ox.Elem(ox.StringT)).
					Slice("network-alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
					Bool("no-healthcheck", "Disable healthchecks on container").
					Bool("no-hosts", "Do not create /etc/hosts within the container, instead use the version from the image").
					Bool("oom-kill-disable", "Disable OOM Killer").
					Int("oom-score-adj", "Tune the host's OOM preferences (-1000 to 1000)").
					String("os", "use OS instead of the running OS for choosing images", ox.Spec("OS")).
					Bool("passwd", "add entries to /etc/passwd and /etc/group", ox.Default("true")).
					String("passwd-entry", "Entry to write to /etc/passwd").
					String("personality", "Configure execution domain using personality (e.g., LINUX/LINUX32)").
					String("pid", "PID namespace to use").
					String("pidfile", "Write the container process ID to the file").
					Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
					String("platform", "Specify the platform for selecting the image.  (Conflicts with --arch and --os)").
					String("pod", "Run container in an existing pod").
					String("pod-id-file", "Read the pod ID from the file").
					Slice("preserve-fd", "Pass a file descriptor into the container", ox.Elem(ox.UintT), ox.Default("[]")).
					Uint("preserve-fds", "Pass a number of additional file descriptors into the container").
					Bool("privileged", "Give extended privileges to container").
					Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
					Bool("publish-all", "Publish all exposed ports to random ports on the host interface", ox.Short("P")).
					String("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Default("missing")).
					Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
					String("rdt-class", "Class of Service (COS) that the container should be assigned to").
					Bool("read-only", "Make containers root filesystem read-only").
					Bool("read-only-tmpfs", "When running --read-only containers mount read-write tmpfs on /dev, /dev/shm, /run, /tmp and /var/tmp", ox.Default("true")).
					Bool("replace", "If a container with the same name exists, replace it").
					Slice("requires", "Add one or more requirement containers that must be started before this container will start", ox.Elem(ox.StringT)).
					String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
					Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
					String("retry-delay", "delay between retries in case of pull failures").
					Bool("rm", "Remove container and any anonymous unnamed volume associated with the container after exit").
					Bool("rmi", "Remove image unless used by other containers, implies --rm").
					Bool("rootfs", "The first argument is not an image but the rootfs to the exploded container").
					String("sdnotify", "control sd-notify behavior (\"container\"|\"conmon\"|\"healthy\"|\"ignore\")", ox.Default("container")).
					String("seccomp-policy", "Policy for selecting a seccomp profile (experimental)", ox.Default("default")).
					Slice("secret", "Add secret to container", ox.Elem(ox.StringT)).
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
					String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true")).
					String("stop-signal", "Signal to stop a container. Default is SIGTERM").
					Uint("stop-timeout", "Timeout (in seconds) that containers stopped by user command have to exit. If exceeded, the container will be forcibly stopped via SIGKILL.", ox.Default("10")).
					String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
					String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
					Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
					String("systemd", "Run container in systemd mode (\"true\"|\"false\"|\"always\")", ox.Default("true")).
					Uint("timeout", "Maximum length of time a container is allowed to run. The container will be killed automatically after the time expires.").
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries for pulling images").
					String("tmpfs", "Mount a temporary filesystem (tmpfs) into a container", ox.Spec("tmpfs")).
					Bool("tty", "Allocate a pseudo-TTY for container", ox.Short("t")).
					String("tz", "Set timezone in container").
					Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("ulimit", "Ulimit options", ox.Elem(ox.StringT)).
					String("umask", "Set umask in container", ox.Default("0022")).
					Slice("unsetenv", "Unset environment default variables in container", ox.Elem(ox.StringT)).
					Bool("unsetenv-all", "Unset all default environment variables in container").
					String("user", "Username or UID (format: <name|uid>[:<group|gid>])", ox.Short("u")).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					String("variant", "Use VARIANT instead of the running architecture variant for choosing images", ox.Spec("VARIANT")).
					Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
					Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)).
					String("workdir", "Working directory inside the container", ox.Short("w")),
			), ox.Sub(
				ox.Banner("Execute the command described by an image label\n\nDescription:\n  Executes a command as described by a container image label."),
				ox.Usage("runlabel", "Execute the command described by an image label"),
				ox.Spec("[options] LABEL IMAGE [ARG...]"),
				ox.Example("\n  podman container runlabel run imageID\n  podman container runlabel install imageID arg1 arg2\n  podman container runlabel --display run myImage"),
				ox.Flags().
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
					Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
					Bool("display", "Preview the command that the label would run").
					String("name", "Assign a name to the container", ox.Short("n")).
					Bool("quiet", "Suppress output information when installing images", ox.Short("q")).
					Bool("replace", "Replace existing container with a new one from the image").
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Start one or more containers\n\nDescription:\n  Starts one or more containers.  The container name or ID can be used."),
				ox.Usage("start", "Start one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container start 860a4b231279 5421ab43b45\n  podman container start --interactive --attach imageID"),
				ox.Flags().
					Bool("all", "Start all containers regardless of their state or configuration").
					Bool("attach", "Attach container's STDOUT and STDERR", ox.Short("a")).
					String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-z`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true if attaching, false otherwise")),
			), ox.Sub(
				ox.Banner("Display a live stream of container resource usage statistics\n\nDescription:\n  Display percentage of CPU, memory, network I/O, block I/O and PIDs for one or more containers."),
				ox.Usage("stats", "Display a live stream of container resource usage statistics"),
				ox.Spec("[options] [CONTAINER...]"),
				ox.Example("\n  podman container stats --all --no-stream\n  podman container stats ctrID\n  podman container stats --no-stream --format \"table {{.ID}} {{.Name}} {{.MemUsage}}\" ctrID"),
				ox.Flags().
					Bool("all", "Show all containers. Only running containers are shown by default. The default is false", ox.Short("a")).
					String("format", "Pretty-print container statistics to JSON or using a Go template").
					Int("interval", "Time in seconds between stats reports", ox.Default("5"), ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-reset", "Disable resetting the screen between intervals").
					Bool("no-stream", "Disable streaming stats and only pull the first result, default setting is false").
					Bool("no-trunc", "Do not truncate output"),
			), ox.Sub(
				ox.Banner("Stop one or more containers\n\nDescription:\n  Stops one or more running containers.  The container name or ID can be used.\n\n  A timeout to forcibly stop the container can also be set but defaults to 10 seconds otherwise."),
				ox.Usage("stop", "Stop one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container stop ctrID\n  podman container stop --time 2 mywebserver 6e534f14da9d"),
				ox.Flags().
					Bool("all", "Stop all running containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("ignore", "Ignore errors when a specified container is missing", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Display the running processes of a container\n\nDescription:\n  Display the running processes of a container.\n\n  The top command extends the ps(1) compatible AIX descriptors with container-specific ones as shown below.  In the presence of ps(1) specific flags (e.g, -eo), Podman will execute ps(1) inside the container."),
				ox.Usage("top", "Display the running processes of a container"),
				ox.Spec("[options] CONTAINER [FORMAT-DESCRIPTORS|ARGS...]"),
				ox.Example("\n  podman container top ctrID\npodman container top ctrID pid seccomp args %C\npodman container top ctrID -eo user,pid,comm"),
				ox.Flags().
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Unmount working container's root filesystem\n\nDescription:\n  Container storage increments a mount counter each time a container is mounted.\n\n  When a container is unmounted, the mount counter is decremented. The container's root filesystem is physically unmounted only when the mount counter reaches zero indicating no other processes are using the mount.\n\n  An unmount can be forced with the --force flag."),
				ox.Usage("unmount", "Unmount working container's root filesystem"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Aliases("unmount", "umount"),
				ox.Example("\n  podman container unmount ctrID\n  podman container unmount ctrID1 ctrID2 ctrID3\n  podman container unmount --all"),
				ox.Flags().
					Bool("all", "Unmount all of the currently mounted containers", ox.Short("a")).
					Bool("force", "Force the complete unmount of the specified mounted containers", ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Unpause the processes in one or more containers\n\nDescription:\n  Unpauses one or more previously paused containers.  The container name or ID can be used."),
				ox.Usage("unpause", "Unpause the processes in one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container unpause ctrID\n  podman container unpause --all"),
				ox.Flags().
					Bool("all", "Unpause all paused containers", ox.Short("a")).
					Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Update an existing container\n\nDescription:\n  Updates the cgroup configuration of a given container"),
				ox.Usage("update", "Update an existing container"),
				ox.Spec("[options] CONTAINER"),
				ox.Example("\n  podman container update --cpus=5 foobar_container"),
				ox.Flags().
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
					Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
					Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
					String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")"),
			), ox.Sub(
				ox.Banner("Block on one or more containers\n\nDescription:\n  Block until one or more containers stop and then print their exit codes."),
				ox.Usage("wait", "Block on one or more containers"),
				ox.Spec("[options] CONTAINER [CONTAINER...]"),
				ox.Example("\n  podman container wait --interval 5s ctrID\n  podman container wait ctrID1 ctrID2"),
				ox.Flags().
					Slice("condition", "Condition to wait on", ox.Elem(ox.StringT)).
					Bool("ignore", "Ignore if a container does not exist").
					String("interval", "Time Interval to wait before polling for completion", ox.Default("250ms"), ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			)), ox.Sub(
			ox.Banner("Copy files/folders between a container and the local filesystem\n\nDescription:\n  Copy the contents of SRC_PATH to the DEST_PATH.\n\n  You can copy from the container's file system to the local machine or the reverse, from the local filesystem to the container. If \"-\" is specified for either the SRC_PATH or DEST_PATH, you can also stream a tar archive from STDIN or to STDOUT. The CONTAINER can be a running or stopped container. The SRC_PATH or DEST_PATH can be a file or a directory."),
			ox.Usage("cp", "Copy files/folders between a container and the local filesystem"),
			ox.Spec("[options] [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH"),
			ox.Example("\n  podman cp [options] [CONTAINER:]SRC_PATH [CONTAINER:]DEST_PATH"),
			ox.Flags().
				Bool("archive", "Chown copied files to the primary uid/gid of the destination container.", ox.Default("true"), ox.Short("a")).
				Bool("overwrite", "Allow to overwrite directories with non-directories and vice versa"),
		), ox.Sub(
			ox.Banner("Create but do not start a container\n\nDescription:\n  Creates a new container from the given image or storage and prepares it for running the specified command.\n\n  The container ID is then printed to stdout. You can then start it at any time with the podman start <container_id> command. The container will be created with the initial state 'created'."),
			ox.Usage("create", "Create but do not start a container"),
			ox.Spec("[options] IMAGE [COMMAND [ARG...]]"),
			ox.Example("\n  podman create alpine ls\n  podman create --annotation HELLO=WORLD alpine ls\n  podman create -t -i --name myctr alpine ls"),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping (host:ip)", ox.Elem(ox.StringT), ox.Default("[]")).
				Slice("annotation", "Add annotations to container (key=value)", ox.Elem(ox.StringT)).
				String("arch", "use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
				String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
				String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
				Slice("cap-add", "Add capabilities to the container", ox.Elem(ox.StringT)).
				Slice("cap-drop", "Drop capabilities from the container", ox.Elem(ox.StringT)).
				Slice("cgroup-conf", "Configure cgroup v2 (key=value)", ox.Elem(ox.StringT)).
				String("cgroup-parent", "Optional parent cgroup for the container").
				String("cgroupns", "cgroup namespace to use").
				String("cgroups", "control container cgroup configuration (\"enabled\"|\"disabled\"|\"no-conmon\"|\"split\")", ox.Default("enabled")).
				Slice("chrootdirs", "Chroot directories inside the container", ox.Elem(ox.StringT)).
				String("cidfile", "Write the container ID to the file").
				String("conmon-pidfile", "Path to the file that will receive the PID of conmon").
				Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
				Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
				Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
				Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
				Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
				Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
				Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed devices list", ox.Elem(ox.StringT)).
				Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
				Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
				Slice("dns-option", "Set custom DNS options", ox.Elem(ox.StringT)).
				Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
				String("entrypoint", "Overwrite the default ENTRYPOINT of the image").
				Slice("env", "Set environment variables in container", ox.Elem(ox.StringT), ox.Default("[PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]"), ox.Short("e")).
				Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
				Bool("env-host", "Use all current host environment variables in container").
				Slice("env-merge", "Preprocess environment variables from image before injecting them into the container", ox.Elem(ox.StringT)).
				Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
				Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
				Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
				Slice("group-add", "Add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
				String("group-entry", "Entry to write to /etc/group").
				String("health-cmd", "set a healthcheck command for the container ('none' disables the existing healthcheck)").
				String("health-interval", "set an interval for the healthcheck (a value of disable results in no automatic timer setup)", ox.Default("30s")).
				String("health-log-destination", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file)", ox.Default("local")).
				Uint("health-max-log-count", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)", ox.Default("5")).
				Uint("health-max-log-size", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)", ox.Default("500")).
				String("health-on-failure", "action to take once the container turns unhealthy", ox.Default("none")).
				Uint("health-retries", "the number of retries allowed before a healthcheck is considered to be unhealthy", ox.Default("3")).
				String("health-start-period", "the initialization time needed for a container to bootstrap", ox.Default("0s")).
				String("health-startup-cmd", "Set a startup healthcheck command for the container").
				String("health-startup-interval", "Set an interval for the startup healthcheck", ox.Default("30s")).
				Uint("health-startup-retries", "Set the maximum number of retries before the startup healthcheck will restart the container").
				Uint("health-startup-success", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)").
				String("health-startup-timeout", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed", ox.Default("30s")).
				String("health-timeout", "the maximum time allowed to complete the healthcheck before an interval is considered failed", ox.Default("30s")).
				Slice("hostuser", "Host user account to add to /etc/passwd within container", ox.Elem(ox.StringT)).
				Bool("http-proxy", "Set proxy environment variables in the container based on the host proxy vars", ox.Default("true")).
				String("image-volume", "Tells podman how to handle the builtin image volumes (\"bind\"|\"tmpfs\"|\"ignore\")", ox.Default("anonymous")).
				Bool("init", "Run an init binary inside the container that forwards signals and reaps processes").
				String("init-ctr", "Make this a pod init container.").
				String("init-path", "Path to the container-init binary").
				Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
				String("ip", "Specify a static IPv4 address for the container").
				String("ip6", "Specify a static IPv6 address for the container").
				String("ipc", "IPC namespace to use").
				Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
				Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
				String("log-driver", "Logging driver for the container", ox.Default("journald")).
				Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
				String("mac-address", "Container MAC address (e.g. 92:d0:c6:0a:29:33)").
				String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
				String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
				String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
				Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
				Slice("mount", "Attach a filesystem mount to the container", ox.Elem(ox.StringT)).
				String("name", "Assign a name to the container").
				Slice("network", "Connect a container to a network", ox.Elem(ox.StringT)).
				Slice("network-alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
				Bool("no-healthcheck", "Disable healthchecks on container").
				Bool("no-hosts", "Do not create /etc/hosts within the container, instead use the version from the image").
				Bool("oom-kill-disable", "Disable OOM Killer").
				Int("oom-score-adj", "Tune the host's OOM preferences (-1000 to 1000)").
				String("os", "use OS instead of the running OS for choosing images", ox.Spec("OS")).
				String("passwd-entry", "Entry to write to /etc/passwd").
				String("personality", "Configure execution domain using personality (e.g., LINUX/LINUX32)").
				String("pid", "PID namespace to use").
				String("pidfile", "Write the container process ID to the file").
				Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
				String("platform", "Specify the platform for selecting the image.  (Conflicts with --arch and --os)").
				String("pod", "Run container in an existing pod").
				String("pod-id-file", "Read the pod ID from the file").
				Bool("privileged", "Give extended privileges to container").
				Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
				Bool("publish-all", "Publish all exposed ports to random ports on the host interface", ox.Short("P")).
				String("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Default("missing")).
				Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
				String("rdt-class", "Class of Service (COS) that the container should be assigned to").
				Bool("read-only", "Make containers root filesystem read-only").
				Bool("read-only-tmpfs", "When running --read-only containers mount read-write tmpfs on /dev, /dev/shm, /run, /tmp and /var/tmp", ox.Default("true")).
				Bool("replace", "If a container with the same name exists, replace it").
				Slice("requires", "Add one or more requirement containers that must be started before this container will start", ox.Elem(ox.StringT)).
				String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
				Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
				String("retry-delay", "delay between retries in case of pull failures").
				Bool("rm", "Remove container and any anonymous unnamed volume associated with the container after exit").
				Bool("rootfs", "The first argument is not an image but the rootfs to the exploded container").
				String("sdnotify", "control sd-notify behavior (\"container\"|\"conmon\"|\"healthy\"|\"ignore\")", ox.Default("container")).
				String("seccomp-policy", "Policy for selecting a seccomp profile (experimental)", ox.Default("default")).
				Slice("secret", "Add secret to container", ox.Elem(ox.StringT)).
				Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
				String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
				String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
				String("stop-signal", "Signal to stop a container. Default is SIGTERM").
				Uint("stop-timeout", "Timeout (in seconds) that containers stopped by user command have to exit. If exceeded, the container will be forcibly stopped via SIGKILL.", ox.Default("10")).
				String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
				String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
				Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
				String("systemd", "Run container in systemd mode (\"true\"|\"false\"|\"always\")", ox.Default("true")).
				Uint("timeout", "Maximum length of time a container is allowed to run. The container will be killed automatically after the time expires.").
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries for pulling images").
				String("tmpfs", "Mount a temporary filesystem (tmpfs) into a container", ox.Spec("tmpfs")).
				Bool("tty", "Allocate a pseudo-TTY for container", ox.Short("t")).
				String("tz", "Set timezone in container").
				Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
				Slice("ulimit", "Ulimit options", ox.Elem(ox.StringT)).
				String("umask", "Set umask in container", ox.Default("0022")).
				Slice("unsetenv", "Unset environment default variables in container", ox.Elem(ox.StringT)).
				Bool("unsetenv-all", "Unset all default environment variables in container").
				String("user", "Username or UID (format: <name|uid>[:<group|gid>])", ox.Short("u")).
				String("userns", "User namespace to use").
				String("uts", "UTS namespace to use").
				String("variant", "Use VARIANT instead of the running architecture variant for choosing images", ox.Spec("VARIANT")).
				Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
				Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)).
				String("workdir", "Working directory inside the container", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Display the changes to the object's file system\n\nDescription:\n  Displays changes on a container or image's filesystem.  The container or image will be compared to its parent layer or the second argument when given."),
			ox.Usage("diff", "Display the changes to the object's file system"),
			ox.Spec("[options] {CONTAINER|IMAGE} [{CONTAINER|IMAGE}]"),
			ox.Example("\n  podman diff imageID\n  podman diff ctrID\n  podman diff --format json redis:alpine"),
			ox.Flags().
				String("format", "Change the output format (json)").
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Show podman system events\n\nDescription:\n  Monitor podman system events.\n\n  By default, streaming mode is used, printing new events as they occur.  Previous events can be listed via --since and --until."),
			ox.Usage("events", "Show podman system events"),
			ox.Spec("[options]"),
			ox.Example("\n  podman events\n  podman events --filter event=create\n  podman events --format {{.Image}}\n  podman events --since 1h30s"),
			ox.Flags().
				Slice("filter", "filter output", ox.Elem(ox.StringT), ox.Short("f")).
				String("format", "format the output using a Go template").
				Bool("no-trunc", "do not truncate the output", ox.Default("true")).
				String("since", "show all events created since timestamp").
				Bool("stream", "stream events and do not exit when returning the last known event", ox.Default("true")).
				String("until", "show all events until timestamp"),
		), ox.Sub(
			ox.Banner("Run a process in a running container\n\nDescription:\n  Execute the specified command inside a running container."),
			ox.Usage("exec", "Run a process in a running container"),
			ox.Spec("[options] CONTAINER COMMAND [ARG...]"),
			ox.Example("\n  podman exec -it ctrID ls\n  podman exec -it -w /tmp myCtr pwd\n  podman exec --user root ctrID ls"),
			ox.Flags().
				Bool("detach", "Run the exec session in detached mode (backgrounded)", ox.Short("d")).
				String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-<value> where <value> is one of: a-z, @, ^, [, , or _", ox.Default("ctrl-p,ctrl-q")).
				Slice("env", "Set environment variables", ox.Elem(ox.StringT), ox.Short("e")).
				Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
				Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Slice("preserve-fd", "Pass a list of additional file descriptors to the container", ox.Elem(ox.UintT), ox.Default("[]")).
				Uint("preserve-fds", "Pass N additional file descriptors to the container").
				Bool("privileged", "Give the process extended Linux capabilities inside the container.  The default is false").
				Bool("tty", "Allocate a pseudo-TTY. The default is false", ox.Short("t")).
				String("user", "Sets the username or UID used and optionally the groupname or GID for the specified command", ox.Short("u")).
				String("workdir", "Working directory inside the container", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Export container's filesystem contents as a tar archive\n\nDescription:\n  Exports container's filesystem contents as a tar archive and saves it on the local machine."),
			ox.Usage("export", "Export container's filesystem contents as a tar archive"),
			ox.Spec("[options] CONTAINER"),
			ox.Example("\n  podman export ctrID > myCtr.tar\n  podman export --output=\"myCtr.tar\" ctrID"),
			ox.Flags().
				String("output", "Write to a specified file (default: stdout, which must be redirected)", ox.Short("o")),
		), ox.Sub(
			ox.Banner("Farm out builds to remote machines\n\nDescription:\n  Farm out builds to remote machines that podman can connect to via podman system connection"),
			ox.Usage("farm", "Farm out builds to remote machines"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Build a container image for multiple architectures\n\nDescription:\n  Build images on farm nodes, then bundle them into a manifest list"),
				ox.Usage("build", "Build a container image for multiple architectures"),
				ox.Spec("[options] [CONTEXT]"),
				ox.Example("\n  podman farm build [flags] buildContextDirectory"),
				ox.Flags().
					String("add-host", "add a custom host-to-IP mapping (host:ip)", ox.Spec("host:ip"), ox.Default("[]")).
					Slice("annotation", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
					String("authfile", "path of the authentication file.").
					Map("build-arg", "argument=value to supply to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					String("build-arg-file", "argfile.conf containing lines of argument=value to supply to the builder", ox.Spec("argfile.conf")).
					Map("build-context", "argument=value to supply additional build context to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					Slice("cache-from", "remote repository list to utilise as potential cache source.", ox.Elem(ox.StringT)).
					Slice("cache-to", "remote repository list to utilise as potential cache destination.", ox.Elem(ox.StringT)).
					String("cache-ttl", "only consider cache images under specified duration.").
					Slice("cap-add", "add the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("cap-drop", "drop the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
					String("cert-dir", "use certificates at the specified path to access the registry").
					String("cgroup-parent", "optional parent cgroup for the container").
					String("cgroupns", "'private', or 'host'").
					Bool("cleanup", "Remove built images from farm nodes on success").
					Bool("compat-volumes", "preserve the contents of VOLUMEs during RUN instructions").
					Slice("cpp-flag", "set additional flag to pass to C preprocessor (cpp)", ox.Elem(ox.StringT)).
					Uint("cpu-period", "limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					String("creds", "use [username[:password]] for accessing the registry", ox.Spec("[username[:password]]")).
					Slice("decryption-key", "key needed to decrypt the image", ox.Elem(ox.StringT)).
					Slice("device", "additional devices to provide", ox.Elem(ox.StringT)).
					Bool("disable-compression", "don't compress layers by default", ox.Default("true"), ox.Short("D")).
					String("dns", "set custom DNS servers or disable it completely by setting it to 'none', which prevents the automatic creation of /etc/resolv.conf.", ox.Spec("path"), ox.Default("/etc/resolv.conf")).
					Slice("dns-option", "set custom DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "set custom DNS search domains", ox.Elem(ox.StringT)).
					Slice("env", "set environment variable for the image", ox.Elem(ox.StringT)).
					String("farm", "Farm to use for builds").
					String("file", "or URL                         pathname or URL of a Dockerfile", ox.Spec("pathname"), ox.Short("f")).
					Bool("force-rm", "always remove intermediate containers after a build, even if the build is unsuccessful.", ox.Default("true")).
					String("format", "format of the built image's manifest and metadata. Use BUILDAH_FORMAT environment variable to override.", ox.Spec("format"), ox.Default("oci")).
					String("from", "image name used to replace the value in the first FROM instruction in the Containerfile").
					Slice("group-add", "add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
					Slice("hooks-dir", "set the OCI hooks directory path (may be set multiple times)", ox.Elem(ox.StringT)).
					Bool("http-proxy", "pass through HTTP Proxy environment variables", ox.Default("true")).
					Bool("identity-label", "add default identity label", ox.Default("true")).
					String("ignorefile", "path to an alternate .dockerignore file").
					String("iidfile", "file to write the image ID to", ox.Spec("file")).
					String("ipc", "'private', path of IPC namespace to join, or 'host'", ox.Spec("path")).
					String("isolation", "type of process isolation to use. Use BUILDAH_ISOLATION environment variable to override.", ox.Spec("type"), ox.Default("rootless")).
					Int("jobs", "how many stages to run in parallel", ox.Default("1")).
					Slice("label", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("layer-label", "set metadata for an intermediate image", ox.Elem(ox.StringT), ox.Default("[]")).
					Bool("layers", "use intermediate layers during build. Use BUILDAH_LAYERS environment variable to override.", ox.Default("true")).
					Bool("local", "Build image on local machine as well as on farm nodes", ox.Default("true"), ox.Short("l")).
					String("logfile", "log to file instead of stdout/stderr", ox.Spec("file")).
					String("memory", "memory limit (format: <number>[<unit>], where unit = b, k, m or g)", ox.Short("m")).
					String("memory-swap", "swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					String("network", "'private', 'none', 'ns:path' of network namespace to join, or 'host'").
					Bool("no-cache", "do not use existing cached images for the container build. Build from the start with a new set of cached layers.").
					Bool("no-hostname", "do not create new /etc/hostname file for RUN instructions, use the one from the base image.").
					Bool("no-hosts", "do not create new /etc/hosts file for RUN instructions, use the one from the base image.").
					Bool("omit-history", "omit build history information from built image").
					String("os-feature", "set required OS feature for the target image in addition to values from the base image", ox.Spec("feature")).
					String("os-version", "set required OS version for the target image instead of the value from the base image", ox.Spec("version")).
					String("pid", "private, path of PID namespace to join, or 'host'", ox.Spec("path")).
					Slice("platforms", "Build only on farm nodes that match the given platforms", ox.Elem(ox.StringT)).
					Map("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Spec("string[=\"missing\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("missing")).
					Bool("quiet", "refrain from announcing build instructions and image read/write progress", ox.Short("q")).
					Int("retry", "number of times to retry in case of failure when performing push/pull", ox.Default("3")).
					String("retry-delay", "delay between retries in case of push/pull failures").
					Bool("rm", "remove intermediate containers after a successful build", ox.Default("true")).
					Slice("runtime-flag", "add global flags for the container runtime", ox.Elem(ox.StringT)).
					String("sbom", "scan working container using preset configuration", ox.Spec("preset")).
					String("sbom-image-output", "add scan results to image as path", ox.Spec("path")).
					String("sbom-image-purl-output", "add scan results to image as path", ox.Spec("path")).
					String("sbom-merge-strategy", "merge scan results using strategy", ox.Spec("strategy")).
					String("sbom-output", "save scan results to file", ox.Spec("file")).
					String("sbom-purl-output", "save scan results to file`", ox.Spec("file")).
					String("sbom-scanner-command", "scan working container using command in scanner image", ox.Spec("command")).
					String("sbom-scanner-image", "scan working container using scanner command from image", ox.Spec("image")).
					Slice("secret", "secret file to expose to the build", ox.Elem(ox.StringT)).
					Slice("security-opt", "security options", ox.Elem(ox.StringT), ox.Default("[]")).
					String("shm-size", "size of '/dev/shm'. The format is <number><unit>.", ox.Spec("<number><unit>"), ox.Default("65536k")).
					Bool("skip-unused-stages", "skips stages in multi-stage builds which do not affect the final target", ox.Default("true")).
					Bool("squash", "squash all image layers into a single layer").
					Bool("squash-all", "Squash all layers into a single layer").
					Slice("ssh", "SSH agent socket or keys to expose to the build. (format: default|<id>[=<socket>|<key>[,<key>]])", ox.Elem(ox.StringT)).
					String("tag", "tagged name to apply to the built image", ox.Spec("name"), ox.Short("t")).
					String("target", "set the target build stage to build").
					Int("timestamp", "set created timestamp to the specified epoch seconds to allow for deterministic builds, defaults to current time").
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")).
					Slice("ulimit", "ulimit options", ox.Elem(ox.StringT)).
					Slice("unsetenv", "unset environment variable from final image", ox.Elem(ox.StringT)).
					Slice("unsetlabel", "unset label when inheriting labels from base image", ox.Elem(ox.StringT)).
					String("userns", "'container', path of user namespace to join, or 'host'", ox.Spec("path")).
					String("userns-gid-map", "containerGID:hostGID:length GID mapping to use in user namespace", ox.Spec("containerGID:hostGID:length")).
					String("userns-gid-map-group", "name of entries from /etc/subgid to use to set user namespace GID mapping", ox.Spec("name")).
					String("userns-uid-map", "containerUID:hostUID:length UID mapping to use in user namespace", ox.Spec("containerUID:hostUID:length")).
					String("userns-uid-map-user", "name of entries from /etc/subuid to use to set user namespace UID mapping", ox.Spec("name")).
					String("uts", "private, :path of UTS namespace to join, or 'host'", ox.Spec("path")).
					Slice("volume", "bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")),
			), ox.Sub(
				ox.Banner("Create a new farm\n\nDescription:\n  Create a new farm with connections added via podman system connection add.\n\n\tThe \"podman system connection add --farm\" command can be used to add a new connection to a new or existing farm."),
				ox.Usage("create", "Create a new farm"),
				ox.Spec("NAME [CONNECTIONS...]"),
				ox.Example("\n  podman farm create myfarm connection1\n  podman farm create myfarm"),
			), ox.Sub(
				ox.Banner("List all existing farms\n\nDescription:\n  podman farm ls\n\nList all available farms. The output of the farms can be filtered\nand the output format can be changed to JSON or a user specified Go template."),
				ox.Usage("list", "List all existing farms"),
				ox.Spec("[options]"),
				ox.Aliases("list", "ls"),
				ox.Flags().
					String("format", "Format farm output using Go template"),
			), ox.Sub(
				ox.Banner("Remove one or more farms\n\nDescription:\n  Remove one or more existing farms."),
				ox.Usage("remove", "Remove one or more farms"),
				ox.Spec("[options] [FARM...]"),
				ox.Aliases("remove", "rm"),
				ox.Example("\n  podman farm rm myfarm1 myfarm2\n  podman farm rm --all"),
				ox.Flags().
					Bool("all", "Remove all farms", ox.Short("a")),
			), ox.Sub(
				ox.Banner("Update an existing farm\n\nDescription:\n  Update an existing farm by adding a connection, removing a connection, or changing it to the default farm."),
				ox.Usage("update", "Update an existing farm"),
				ox.Spec("[options] FARM"),
				ox.Example("\n  podman farm update --add con1 farm1\n\tpodman farm update --remove con2 farm2\n\tpodman farm update --default farm3"),
				ox.Flags().
					Slice("add", "add system connection(s) to farm", ox.Elem(ox.StringT), ox.Short("a")).
					Bool("default", "set the given farm as the default farm", ox.Short("d")).
					Slice("remove", "remove system connection(s) from farm", ox.Elem(ox.StringT), ox.Short("r")),
			)), ox.Sub(
			ox.Banner("Generate structured data based on containers, pods or volumes\n\nDescription:\n  Generate structured data (e.g., Kubernetes YAML or systemd units) based on containers, pods or volumes."),
			ox.Usage("generate", "Generate structured data based on containers, pods or volumes"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Generate Kubernetes YAML from containers, pods or volumes.\n\nDescription:\n  Command generates Kubernetes Pod, Service or PersistentVolumeClaim YAML (v1 specification) from Podman containers, pods or volumes.\n\n  Whether the input is for a container or pod, Podman will always generate the specification as a pod."),
				ox.Usage("kube", "Generate Kubernetes YAML from containers, pods or volumes."),
				ox.Spec("[options] {CONTAINER...|POD...|VOLUME...}"),
				ox.Example("\n  podman kube generate ctrID\n  podman kube generate podID\n  podman kube generate --service podID\n  podman kube generate volumeName\n  podman kube generate ctrID podID volumeName --service"),
				ox.Flags().
					String("filename", "Write output to the specified path", ox.Short("f")).
					Bool("podman-only", "Add podman-only reserved annotations to the generated YAML file (Cannot be used by Kubernetes)").
					Int32("replicas", "Set the replicas number for Deployment kind", ox.Default("1"), ox.Short("r")).
					Bool("service", "Generate YAML for a Kubernetes service object", ox.Short("s")).
					String("type", "Generate YAML for the given Kubernetes kind", ox.Default("pod"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Generate Specgen JSON based on containers or pods\n\nDescription:\n  Generate Specgen JSON based on containers or pods"),
				ox.Usage("spec", "Generate Specgen JSON based on containers or pods"),
				ox.Spec("[options] {CONTAINER|POD}"),
				ox.Example("\n  podman generate spec ctrID"),
				ox.Flags().
					Bool("compact", "Print the json in a compact format for consumption", ox.Short("c")).
					String("filename", "Write output to the specified path", ox.Short("f")).
					Bool("name", "Specify a new name for the generated spec", ox.Default("true"), ox.Short("n")),
			), ox.Sub(
				ox.Banner("[DEPRECATED] Generate systemd units\n\nDescription:\n  Generate systemd units for a pod or container.\n  The generated units can later be controlled via systemctl(1).\n\nDEPRECATED command:\nIt is recommended to use Quadlets for running containers and pods under systemd.\n\nPlease refer to podman-systemd.unit(5) for details."),
				ox.Usage("systemd", "[DEPRECATED] Generate systemd units"),
				ox.Spec("[options] {CONTAINER|POD}"),
				ox.Example("\n  podman generate systemd CTR\n  podman generate systemd --new --time 10 CTR\n  podman generate systemd --files --name POD"),
				ox.Flags().
					Slice("after", "Add dependencies order to the generated unit file", ox.Elem(ox.StringT)).
					String("container-prefix", "Systemd unit name prefix for containers", ox.Default("container")).
					Slice("env", "Set environment variables to the systemd unit files", ox.Elem(ox.StringT), ox.Short("e")).
					Bool("files", "Generate .service files instead of printing to stdout", ox.Short("f")).
					String("format", "Print the created units in specified format (json)").
					Bool("name", "Use container/pod names instead of IDs", ox.Short("n")).
					Bool("new", "Create a new container or pod instead of starting an existing one").
					Bool("no-header", "Skip header generation").
					String("pod-prefix", "Systemd unit name prefix for pods", ox.Default("pod")).
					Slice("requires", "Similar to wants, but declares stronger requirement dependencies", ox.Elem(ox.StringT)).
					String("restart-policy", "Systemd restart-policy", ox.Default("on-failure")).
					Uint("restart-sec", "Systemd restart-sec").
					String("separator", "Systemd unit name separator between name/id and prefix", ox.Default("-")).
					Uint("start-timeout", "Start timeout override").
					Uint("stop-timeout", "Stop timeout override", ox.Default("10")).
					Bool("template", "Make it a template file and use %i and %I specifiers. Working only for containers").
					Slice("wants", "Add (weak) requirement dependencies to the generated unit file", ox.Elem(ox.StringT)),
			)), ox.Sub(
			ox.Banner("Manage health checks on containers\n\nDescription:\n  Run health checks on containers"),
			ox.Usage("healthcheck", "Manage health checks on containers"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Run the health check of a container\n\nDescription:\n  Run the health check of a container"),
				ox.Usage("run", "Run the health check of a container"),
				ox.Spec("CONTAINER"),
				ox.Example("\n  podman healthcheck run mywebapp"),
			)), ox.Sub(
			ox.Banner("Show history of a specified image\n\nDescription:\n  Displays the history of an image.\n\n  The information can be printed out in an easy to read, or user specified format, and can be truncated."),
			ox.Usage("history", "Show history of a specified image"),
			ox.Spec("[options] IMAGE"),
			ox.Example("\n  podman history quay.io/fedora/fedora"),
			ox.Flags().
				String("format", "Change the output to JSON or a Go template").
				Bool("human", "Display sizes and dates in human readable format", ox.Default("true"), ox.Short("H")).
				Bool("no-trunc", "Do not truncate the output").
				Bool("quiet", "Display the numeric IDs only", ox.Short("q")),
		), ox.Sub(
			ox.Banner("Manage images\n\nDescription:\n  Manage images"),
			ox.Usage("image", "Manage images"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Build an image using instructions from Containerfiles\n\nDescription:\n  Builds an OCI or Docker image using instructions from one or more Containerfiles and a specified build context directory."),
				ox.Usage("build", "Build an image using instructions from Containerfiles"),
				ox.Spec("[options] [CONTEXT]"),
				ox.Example("\n  podman image build .\n  podman image build --creds=username:password -t imageName -f Containerfile.simple .\n  podman image build --layers --force-rm --tag imageName ."),
				ox.Flags().
					String("add-host", "add a custom host-to-IP mapping (host:ip)", ox.Spec("host:ip"), ox.Default("[]")).
					Bool("all-platforms", "attempt to build for all base image platforms").
					Slice("annotation", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
					String("arch", "set the ARCH of the image to the provided value instead of the architecture of the host", ox.Default("amd64")).
					String("authfile", "path of the authentication file.").
					Map("build-arg", "argument=value to supply to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					String("build-arg-file", "argfile.conf containing lines of argument=value to supply to the builder", ox.Spec("argfile.conf")).
					Map("build-context", "argument=value to supply additional build context to the builder", ox.Spec("argument=value"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT)).
					Slice("cache-from", "remote repository list to utilise as potential cache source.", ox.Elem(ox.StringT)).
					Slice("cache-to", "remote repository list to utilise as potential cache destination.", ox.Elem(ox.StringT)).
					String("cache-ttl", "only consider cache images under specified duration.").
					Slice("cap-add", "add the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("cap-drop", "drop the specified capability when running", ox.Elem(ox.StringT), ox.Default("[]")).
					String("cert-dir", "use certificates at the specified path to access the registry").
					String("cgroup-parent", "optional parent cgroup for the container").
					String("cgroupns", "'private', or 'host'").
					Bool("compat-volumes", "preserve the contents of VOLUMEs during RUN instructions").
					Bool("compress", "this is a legacy option, which has no effect on the image").
					Slice("cpp-flag", "set additional flag to pass to C preprocessor (cpp)", ox.Elem(ox.StringT)).
					Uint("cpu-period", "limit the CPU CFS (Completely Fair Scheduler) period").
					Int("cpu-quota", "limit the CPU CFS (Completely Fair Scheduler) quota").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					String("creds", "use [username[:password]] for accessing the registry", ox.Spec("[username[:password]]")).
					Slice("cw", "confidential workload options", ox.Elem(ox.StringT)).
					Slice("decryption-key", "key needed to decrypt the image", ox.Elem(ox.StringT)).
					Slice("device", "additional devices to provide", ox.Elem(ox.StringT)).
					Bool("disable-compression", "don't compress layers by default", ox.Default("true"), ox.Short("D")).
					Bool("disable-content-trust", "this is a Docker specific option and is a NOOP").
					String("dns", "set custom DNS servers or disable it completely by setting it to 'none', which prevents the automatic creation of /etc/resolv.conf.", ox.Spec("path"), ox.Default("/etc/resolv.conf")).
					Slice("dns-option", "set custom DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "set custom DNS search domains", ox.Elem(ox.StringT)).
					Slice("env", "set environment variable for the image", ox.Elem(ox.StringT)).
					String("file", "or URL                         pathname or URL of a Dockerfile", ox.Spec("pathname"), ox.Short("f")).
					Bool("force-rm", "always remove intermediate containers after a build, even if the build is unsuccessful.", ox.Default("true")).
					String("format", "format of the built image's manifest and metadata. Use BUILDAH_FORMAT environment variable to override.", ox.Spec("format"), ox.Default("oci")).
					String("from", "image name used to replace the value in the first FROM instruction in the Containerfile").
					Slice("group-add", "add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
					Slice("hooks-dir", "set the OCI hooks directory path (may be set multiple times)", ox.Elem(ox.StringT)).
					Bool("http-proxy", "pass through HTTP Proxy environment variables", ox.Default("true")).
					Bool("identity-label", "add default identity label", ox.Default("true")).
					String("ignorefile", "path to an alternate .dockerignore file").
					String("iidfile", "file to write the image ID to", ox.Spec("file")).
					String("ipc", "'private', path of IPC namespace to join, or 'host'", ox.Spec("path")).
					String("isolation", "type of process isolation to use. Use BUILDAH_ISOLATION environment variable to override.", ox.Spec("type"), ox.Default("rootless")).
					Int("jobs", "how many stages to run in parallel", ox.Default("1")).
					Slice("label", "set metadata for an image", ox.Elem(ox.StringT), ox.Default("[]")).
					Slice("layer-label", "set metadata for an intermediate image", ox.Elem(ox.StringT), ox.Default("[]")).
					Bool("layers", "use intermediate layers during build. Use BUILDAH_LAYERS environment variable to override.", ox.Default("true")).
					String("logfile", "log to file instead of stdout/stderr", ox.Spec("file")).
					Bool("logsplit", "split logfile to different files for each platform").
					String("manifest", "add the image to the specified manifest list. Creates manifest list if it does not exist").
					String("memory", "memory limit (format: <number>[<unit>], where unit = b, k, m or g)", ox.Short("m")).
					String("memory-swap", "swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					String("network", "'private', 'none', 'ns:path' of network namespace to join, or 'host'").
					Bool("no-cache", "do not use existing cached images for the container build. Build from the start with a new set of cached layers.").
					Bool("no-hostname", "do not create new /etc/hostname file for RUN instructions, use the one from the base image.").
					Bool("no-hosts", "do not create new /etc/hosts file for RUN instructions, use the one from the base image.").
					Bool("omit-history", "omit build history information from built image").
					String("os", "set the OS to the provided value instead of the current operating system of the host", ox.Default("linux")).
					String("os-feature", "set required OS feature for the target image in addition to values from the base image", ox.Spec("feature")).
					String("os-version", "set required OS version for the target image instead of the value from the base image", ox.Spec("version")).
					String("output", "output destination (format: type=local,dest=path)", ox.Short("o")).
					String("pid", "private, path of PID namespace to join, or 'host'", ox.Spec("path")).
					String("platform", "set the OS/ARCH[/VARIANT] of the image to the provided value instead of the current operating system and architecture of the host (for example \"linux/arm\")", ox.Spec("OS/ARCH[/VARIANT]"), ox.Default("[linux/amd64]")).
					Map("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Spec("string[=\"missing\"]"), ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("missing")).
					Bool("quiet", "refrain from announcing build instructions and image read/write progress", ox.Short("q")).
					Int("retry", "number of times to retry in case of failure when performing push/pull", ox.Default("3")).
					String("retry-delay", "delay between retries in case of push/pull failures").
					Bool("rm", "remove intermediate containers after a successful build", ox.Default("true")).
					Slice("runtime-flag", "add global flags for the container runtime", ox.Elem(ox.StringT)).
					String("sbom", "scan working container using preset configuration", ox.Spec("preset")).
					String("sbom-image-output", "add scan results to image as path", ox.Spec("path")).
					String("sbom-image-purl-output", "add scan results to image as path", ox.Spec("path")).
					String("sbom-merge-strategy", "merge scan results using strategy", ox.Spec("strategy")).
					String("sbom-output", "save scan results to file", ox.Spec("file")).
					String("sbom-purl-output", "save scan results to file`", ox.Spec("file")).
					String("sbom-scanner-command", "scan working container using command in scanner image", ox.Spec("command")).
					String("sbom-scanner-image", "scan working container using scanner command from image", ox.Spec("image")).
					Slice("secret", "secret file to expose to the build", ox.Elem(ox.StringT)).
					Slice("security-opt", "security options", ox.Elem(ox.StringT), ox.Default("[]")).
					String("shm-size", "size of '/dev/shm'. The format is <number><unit>.", ox.Spec("<number><unit>"), ox.Default("65536k")).
					String("sign-by", "sign the image using a GPG key with the specified FINGERPRINT", ox.Spec("FINGERPRINT")).
					Bool("skip-unused-stages", "skips stages in multi-stage builds which do not affect the final target", ox.Default("true")).
					Bool("squash", "squash all image layers into a single layer").
					Bool("squash-all", "Squash all layers into a single layer").
					Slice("ssh", "SSH agent socket or keys to expose to the build. (format: default|<id>[=<socket>|<key>[,<key>]])", ox.Elem(ox.StringT)).
					Bool("stdin", "pass stdin into containers").
					String("tag", "tagged name to apply to the built image", ox.Spec("name"), ox.Short("t")).
					String("target", "set the target build stage to build").
					Int("timestamp", "set created timestamp to the specified epoch seconds to allow for deterministic builds, defaults to current time").
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")).
					Slice("ulimit", "ulimit options", ox.Elem(ox.StringT)).
					Slice("unsetenv", "unset environment variable from final image", ox.Elem(ox.StringT)).
					Slice("unsetlabel", "unset label when inheriting labels from base image", ox.Elem(ox.StringT)).
					String("userns", "'container', path of user namespace to join, or 'host'", ox.Spec("path")).
					String("userns-gid-map", "containerGID:hostGID:length GID mapping to use in user namespace", ox.Spec("containerGID:hostGID:length")).
					String("userns-gid-map-group", "name of entries from /etc/subgid to use to set user namespace GID mapping", ox.Spec("name")).
					String("userns-uid-map", "containerUID:hostUID:length UID mapping to use in user namespace", ox.Spec("containerUID:hostUID:length")).
					String("userns-uid-map-user", "name of entries from /etc/subuid to use to set user namespace UID mapping", ox.Spec("name")).
					String("uts", "private, :path of UTS namespace to join, or 'host'", ox.Spec("path")).
					String("variant", "override the variant of the specified image", ox.Spec("variant")).
					Slice("volume", "bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")),
			), ox.Sub(
				ox.Banner("Inspect changes to the image's file systems\n\nDescription:\n  Displays changes to the image's filesystem.  The image will be compared to its parent layer or the second argument when given."),
				ox.Usage("diff", "Inspect changes to the image's file systems"),
				ox.Spec("[options] IMAGE [IMAGE]"),
				ox.Example("\n  podman image diff myImage\n  podman image diff --format json redis:alpine"),
				ox.Flags().
					String("format", "Change the output format (json)"),
			), ox.Sub(
				ox.Banner("Check if an image exists in local storage\n\nDescription:\n  If the named image exists in local storage, podman image exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if an image exists in local storage"),
				ox.Spec("IMAGE"),
				ox.Example("\n  podman image exists ID\n  podman image exists IMAGE && podman pull IMAGE"),
			), ox.Sub(
				ox.Banner("Show history of a specified image\n\nDescription:\n  Displays the history of an image.\n\n  The information can be printed out in an easy to read, or user specified format, and can be truncated."),
				ox.Usage("history", "Show history of a specified image"),
				ox.Spec("[options] IMAGE"),
				ox.Example("\n  podman image history quay.io/fedora/fedora"),
				ox.Flags().
					String("format", "Change the output to JSON or a Go template").
					Bool("human", "Display sizes and dates in human readable format", ox.Default("true"), ox.Short("H")).
					Bool("no-trunc", "Do not truncate the output").
					Bool("quiet", "Display the numeric IDs only", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Import a tarball to create a filesystem image\n\nDescription:\n  Create a container image from the contents of the specified tarball (.tar, .tar.gz, .tgz, .bzip, .tar.xz, .txz).\n\n  Note remote tar balls can be specified, via web address.\n  Optionally tag the image. You can specify the instructions using the --change option."),
				ox.Usage("import", "Import a tarball to create a filesystem image"),
				ox.Spec("[options] PATH [REFERENCE]"),
				ox.Example("\n  podman image import https://example.com/ctr.tar url-image\n  cat ctr.tar | podman -q image import --message \"importing the ctr.tar tarball\" - image-imported\n  cat ctr.tar | podman image import -"),
				ox.Flags().
					String("arch", "Set the architecture of the imported image").
					Slice("change", "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR", ox.Elem(ox.StringT), ox.Short("c")).
					String("message", "Set commit message for imported image", ox.Short("m")).
					String("os", "Set the OS of the imported image").
					Bool("quiet", "Suppress output", ox.Short("q")).
					String("variant", "Set the variant of the imported image"),
			), ox.Sub(
				ox.Banner("Display the configuration of an image\n\nDescription:\n  Displays the low-level information of an image identified by name or ID."),
				ox.Usage("inspect", "Display the configuration of an image"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Example("\n  podman image inspect alpine\n  podman image inspect --format \"imageId: {{.Id}} size: {{.Size}}\" alpine\n  podman image inspect --format \"image: {{.ImageName}} driver: {{.Driver}}\" myctr"),
				ox.Flags().
					String("format", "Format the output to a Go template or json", ox.Default("json"), ox.Short("f")),
			), ox.Sub(
				ox.Banner("List images in local storage\n\nDescription:\n  Lists images previously pulled to the system or created on the system."),
				ox.Usage("list", "List images in local storage"),
				ox.Spec("[options] [IMAGE]"),
				ox.Aliases("list", "ls"),
				ox.Example("\n  podman image list --format json\n  podman image list --sort repository --format \"table {{.ID}} {{.Repository}} {{.Tag}}\"\n  podman image list --filter dangling=true"),
				ox.Flags().
					Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a")).
					Bool("digests", "Show digests").
					Slice("filter", "Filter output based on conditions provided", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("f")).
					String("format", "Change the output format to JSON or a Go template").
					Bool("history", "Display the image name history").
					Bool("no-trunc", "Do not truncate output").
					Bool("noheading", "Do not print column headings", ox.Short("n")).
					Bool("quiet", "Display only image IDs", ox.Short("q")).
					String("sort", "Sort by size, tag, created, id, repository", ox.Default("created")),
			), ox.Sub(
				ox.Banner("Load image(s) from a tar archive\n\nDescription:\n  Loads an image from a locally stored archive (tar file) into container storage."),
				ox.Usage("load", "Load image(s) from a tar archive"),
				ox.Spec("[options]"),
				ox.Flags().
					String("input", "Read from specified archive file (default: stdin)", ox.Short("i")).
					Bool("quiet", "Suppress the output", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Mount an image's root filesystem\n\nDescription:\n  podman image mount\n    Lists all mounted images mount points if no images is specified\n\n  podman image mount IMAGE-NAME-OR-ID\n    Mounts the specified image and prints the mountpoint"),
				ox.Usage("mount", "Mount an image's root filesystem"),
				ox.Spec("[options] [IMAGE...]"),
				ox.Example("\n  podman image mount imgID\n  podman image mount imgID1 imgID2 imgID3\n  podman image mount\n  podman image mount --all"),
				ox.Flags().
					Bool("all", "Mount all images", ox.Short("a")).
					String("format", "Print the mounted images in specified format (json)"),
			), ox.Sub(
				ox.Banner("Remove unused images\n\nDescription:\n  Removes dangling or unused images from local storage."),
				ox.Usage("prune", "Remove unused images"),
				ox.Spec("[options]"),
				ox.Example("\n  podman image prune"),
				ox.Flags().
					Bool("all", "Remove all images not in use by containers, not just dangling ones", ox.Short("a")).
					Bool("build-cache", "Remove persistent build cache created by --mount=type=cache").
					Bool("external", "Remove images even when they are used by external containers (e.g., by build containers)").
					Slice("filter", "Provide filter values (e.g. 'label=<key>=<value>')", ox.Elem(ox.StringT)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Pull an image from a registry\n\nDescription:\n  Pulls an image from a registry and stores it locally.\n\n  An image can be pulled by tag or digest. If a tag is not specified, the image with the 'latest' tag is pulled."),
				ox.Usage("pull", "Pull an image from a registry"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Example("\n  podman image pull imageName\n  podman image pull fedora:latest"),
				ox.Flags().
					Bool("all-tags", "All tagged images in the repository will be pulled", ox.Short("a")).
					String("arch", "Use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
					Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
					Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
					Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
					String("os", "Use OS instead of the running OS for choosing images", ox.Spec("OS")).
					String("platform", "Specify the platform for selecting the image.  (Conflicts with arch and os)").
					Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
					Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
					String("retry-delay", "delay between retries in case of pull failures").
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")).
					String("variant", "Use VARIANT instead of the running architecture variant for choosing images"),
			), ox.Sub(
				ox.Banner("Push an image to a specified destination\n\nDescription:\n  Pushes a source image to a specified destination.\n\n\tThe Image \"DESTINATION\" uses a \"transport\":\"details\" format. See podman-push(1) section \"DESTINATION\" for the expected format."),
				ox.Usage("push", "Push an image to a specified destination"),
				ox.Spec("[options] IMAGE [DESTINATION]"),
				ox.Example("\n  podman image push imageID docker://registry.example.com/repository:tag\n\t\tpodman image push imageID oci-archive:/path/to/layout:image:tag"),
				ox.Flags().
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "Path to a directory containing TLS certificates and keys").
					Bool("compress", "Compress tarball image layers when pushing to a directory using the 'dir' transport.", ox.Default("is same compression type as source")).
					String("compression-format", "compression format to use", ox.Default("gzip")).
					Int("compression-level", "compression level to use").
					Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
					String("digestfile", "Write the digest of the pushed image to the specified file").
					Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
					Slice("encrypt-layer", "Layers to encrypt, 0-indexed layer indices with support for negative indexing (e.g. 0 is the first layer, -1 is the last layer). If not defined, will encrypt all layers if encryption-key flag is specified", ox.Elem(ox.IntT)).
					Slice("encryption-key", "Key with the encryption protocol to use to encrypt the image (e.g. jwe:/path/to/key.pem)", ox.Elem(ox.StringT)).
					Bool("force-compression", "Use the specified compression algorithm even if the destination contains a differently-compressed variant already").
					String("format", "Manifest type (oci, v2s2, or v2s1) to use in the destination", ox.Default("is manifest type of source, with fallbacks"), ox.Short("f")).
					Bool("quiet", "Suppress output information when pushing images", ox.Short("q")).
					Bool("remove-signatures", "Discard any pre-existing signatures in the image").
					Uint("retry", "number of times to retry in case of failure when performing push", ox.Default("3")).
					String("retry-delay", "delay between retries in case of push failures").
					String("sign-by", "Add a signature at the destination using the specified key").
					String("sign-by-sigstore", "Sign the image using a sigstore parameter file at PATH", ox.Spec("PATH")).
					String("sign-by-sigstore-private-key", "Sign the image using a sigstore private key at PATH", ox.Spec("PATH")).
					String("sign-passphrase-file", "Read a passphrase for signing an image from PATH", ox.Spec("PATH")).
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Remove one or more images from local storage\n\nDescription:\n  Removes one or more previously pulled or locally created images."),
				ox.Usage("rm", "Remove one or more images from local storage"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Example("\n  podman image rm imageID\n  podman image rm --force alpine\n  podman image rm c4dfb1609ee2 93fd78260bd1 c0ed59d05ff7"),
				ox.Flags().
					Bool("all", "Remove all images", ox.Short("a")).
					Bool("force", "Force Removal of the image", ox.Short("f")).
					Bool("ignore", "Ignore errors if a specified image does not exist", ox.Short("i")).
					Bool("no-prune", "Do not remove dangling images"),
			), ox.Sub(
				ox.Banner("Save image(s) to an archive\n\nDescription:\n  Save an image to docker-archive or oci-archive on the local machine. Default is docker-archive."),
				ox.Usage("save", "Save image(s) to an archive"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Example("\n  podman image save --quiet -o myimage.tar imageID\n  podman image save --format docker-dir -o ubuntu-dir ubuntu\n  podman image save > alpine-all.tar alpine:latest"),
				ox.Flags().
					Bool("compress", "Compress tarball image layers when saving to a directory using the 'dir' transport.", ox.Default("is same compression type as source")).
					String("format", "Save image to oci-archive, oci-dir (directory with oci manifest type), docker-archive, docker-dir (directory with v2s2 manifest type)", ox.Default("docker-archive")).
					Bool("multi-image-archive", "Interpret additional arguments as images not tags and create a multi-image-archive (only for docker-archive)", ox.Short("m")).
					String("output", "Write to a specified file (default: stdout, which must be redirected)", ox.Short("o")).
					Bool("quiet", "Suppress the output", ox.Short("q")).
					Bool("uncompressed", "Accept uncompressed layers when copying OCI images"),
			), ox.Sub(
				ox.Banner("Securely copy images\n\nDescription:\n  Securely copy an image from one host to another."),
				ox.Usage("scp", "Securely copy images"),
				ox.Spec("[options] IMAGE [HOST::]"),
				ox.Example("\n  podman image scp myimage:latest otherhost::"),
				ox.Flags().
					Bool("quiet", "Suppress the output", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Search registry for image\n\nDescription:\n  Search registries for a given image. Can search all the default registries or a specific registry.\n\n\tUsers can limit the number of results, and filter the output based on certain conditions."),
				ox.Usage("search", "Search registry for image"),
				ox.Spec("[options] TERM"),
				ox.Example("\n  podman image search --filter=is-official --limit 3 alpine\n  podman image search registry.fedoraproject.org/  # only works with v2 registries\n  podman image search --format \"table {{.Index}} {{.Name}}\" registry.fedoraproject.org/fedora"),
				ox.Flags().
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
					Bool("compatible", "List stars, official and automated columns (Docker compatibility)").
					Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
					Slice("filter", "Filter output based on conditions provided", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("f")).
					String("format", "Change the output format to JSON or a Go template").
					Int("limit", "Limit the number of results").
					Bool("list-tags", "List the tags of the input registry").
					Bool("no-trunc", "Do not truncate the output").
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Sign an image\n\nDescription:\n  Create a signature file that can be used later to verify the image."),
				ox.Usage("sign", "Sign an image"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Example("\n  podman image sign --sign-by mykey imageID\n  podman image sign --sign-by mykey --directory ./mykeydir imageID"),
				ox.Flags().
					Bool("all", "Sign all the manifests of the multi-architecture image", ox.Short("a")).
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
					String("directory", "Define an alternate directory to store signatures", ox.Short("d")).
					String("sign-by", "Name of the signing key"),
			), ox.Sub(
				ox.Banner("Add an additional name to a local image\n\nDescription:\n  Adds one or more additional names to locally-stored image."),
				ox.Usage("tag", "Add an additional name to a local image"),
				ox.Spec("IMAGE TARGET_NAME [TARGET_NAME...]"),
				ox.Example("\n  podman image tag 0e3bbc2 fedora:latest\n  podman image tag imageID:latest myNewImage:newTag\n  podman image tag httpd myregistryhost:5000/fedora/httpd:v2"),
			), ox.Sub(
				ox.Banner("Print layer hierarchy of an image in a tree format\n\nDescription:\n  Print layer hierarchy of an image in a tree format"),
				ox.Usage("tree", "Print layer hierarchy of an image in a tree format"),
				ox.Spec("[options] IMAGE"),
				ox.Example("\n  podman image tree alpine:latest"),
				ox.Flags().
					Bool("whatrequires", "Show all child images and layers of the specified image"),
			), ox.Sub(
				ox.Banner("Manage container image trust policy\n\nDescription:\n  Manages which registries you trust as a source of container images based on their location.\n  The location is determined by the transport and the registry host of the image.  Using this container image docker://quay.io/podman/stable as an example, docker is the transport and quay.io is the registry host."),
				ox.Usage("trust", "Manage container image trust policy"),
				ox.Spec("[command]"),
				ox.Sub(
					ox.Banner("Set default trust policy or a new trust policy for a registry\n\nDescription:\n  Set default trust policy or add a new trust policy for a registry"),
					ox.Usage("set", "Set default trust policy or a new trust policy for a registry"),
					ox.Spec("[options] REGISTRY"),
					ox.Flags().
						Slice("pubkeysfile", "Path of installed public key(s) to trust for TARGET.", ox.Elem(ox.StringT), ox.Short("f")).
						String("type", "Trust type, accept values: signedBy(default), accept, reject", ox.Default("signedBy"), ox.Short("t")),
				), ox.Sub(
					ox.Banner("Display trust policy for the system\n\nDescription:\n  Display trust policy for the system"),
					ox.Usage("show", "Display trust policy for the system"),
					ox.Spec("[options] [REGISTRY]"),
					ox.Flags().
						Bool("json", "Output as json", ox.Short("j")).
						Bool("noheading", "Do not print column headings", ox.Short("n")).
						Bool("raw", "Output raw policy file"),
				)), ox.Sub(
				ox.Banner("Unmount an image's root filesystem\n\nDescription:\n  Image storage increments a mount counter each time an image is mounted.\n\n  When an image is unmounted, the mount counter is decremented. The image's root filesystem is physically unmounted only when the mount counter reaches zero indicating no other processes are using the mount.\n\n  An unmount can be forced with the --force flag."),
				ox.Usage("unmount", "Unmount an image's root filesystem"),
				ox.Spec("[options] IMAGE [IMAGE...]"),
				ox.Aliases("unmount", "umount"),
				ox.Example("\n  podman unmount imgID\n  podman unmount imgID1 imgID2 imgID3\n  podman unmount --all"),
				ox.Flags().
					Bool("all", "Unmount all of the currently mounted images", ox.Short("a")).
					Bool("force", "Force the complete unmount of the specified mounted images", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Remove a name from a local image\n\nDescription:\n  Removes one or more names from a locally-stored image."),
				ox.Usage("untag", "Remove a name from a local image"),
				ox.Spec("IMAGE [IMAGE...]"),
				ox.Example("\n  podman image untag 0e3bbc2\n  podman image untag imageID:latest otherImageName:latest\n  podman image untag httpd myregistryhost:5000/fedora/httpd:v2"),
			)), ox.Sub(
			ox.Banner("List images in local storage\n\nDescription:\n  Lists images previously pulled to the system or created on the system."),
			ox.Usage("images", "List images in local storage"),
			ox.Spec("[options] [IMAGE]"),
			ox.Example("\n  podman images --format json\n  podman images --sort repository --format \"table {{.ID}} {{.Repository}} {{.Tag}}\"\n  podman images --filter dangling=true"),
			ox.Flags().
				Bool("all", "Show all images", ox.Default("hides intermediate images"), ox.Short("a")).
				Bool("digests", "Show digests").
				Slice("filter", "Filter output based on conditions provided", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("f")).
				String("format", "Change the output format to JSON or a Go template").
				Bool("history", "Display the image name history").
				Bool("no-trunc", "Do not truncate output").
				Bool("noheading", "Do not print column headings", ox.Short("n")).
				Bool("quiet", "Display only image IDs", ox.Short("q")).
				String("sort", "Sort by size, tag, created, id, repository", ox.Default("created")),
		), ox.Sub(
			ox.Banner("Import a tarball to create a filesystem image\n\nDescription:\n  Create a container image from the contents of the specified tarball (.tar, .tar.gz, .tgz, .bzip, .tar.xz, .txz).\n\n  Note remote tar balls can be specified, via web address.\n  Optionally tag the image. You can specify the instructions using the --change option."),
			ox.Usage("import", "Import a tarball to create a filesystem image"),
			ox.Spec("[options] PATH [REFERENCE]"),
			ox.Example("\n  podman import https://example.com/ctr.tar url-image\n  cat ctr.tar | podman -q import --message \"importing the ctr.tar tarball\" - image-imported\n  cat ctr.tar | podman import -"),
			ox.Flags().
				String("arch", "Set the architecture of the imported image").
				Slice("change", "Apply the following possible instructions to the created image (default []): CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | ONBUILD | STOPSIGNAL | USER | VOLUME | WORKDIR", ox.Elem(ox.StringT), ox.Short("c")).
				String("message", "Set commit message for imported image", ox.Short("m")).
				String("os", "Set the OS of the imported image").
				Bool("quiet", "Suppress output", ox.Short("q")).
				String("variant", "Set the variant of the imported image"),
		), ox.Sub(
			ox.Banner("Display podman system information\n\nDescription:\n  Display information pertaining to the host, current storage stats, and build of podman.\n\n  Useful for the user and when reporting issues."),
			ox.Usage("info", "Display podman system information"),
			ox.Spec("[options]"),
			ox.Example("\n  podman info"),
			ox.Flags().
				String("format", "Change the output format to JSON or a Go template", ox.Short("f")),
		), ox.Sub(
			ox.Banner("Initialize one or more containers\n\nDescription:\n  Initialize one or more containers, creating the OCI spec and mounts for inspection. Container names or IDs can be used."),
			ox.Usage("init", "Initialize one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman init 3c45ef19d893\n  podman init test1"),
			ox.Flags().
				Bool("all", "Initialize all containers", ox.Short("a")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Display the configuration of object denoted by ID\n\nDescription:\n  Displays the low-level information on an object identified by name or ID.\n  For more inspection options, see:\n\n      podman container inspect\n      podman image inspect\n      podman network inspect\n      podman pod inspect\n      podman volume inspect"),
			ox.Usage("inspect", "Display the configuration of object denoted by ID"),
			ox.Spec("[options] {CONTAINER|IMAGE|POD|NETWORK|VOLUME} [...]"),
			ox.Example("\n  podman inspect fedora\n  podman inspect --type image fedora\n  podman inspect CtrID ImgID\n  podman inspect --format \"imageId: {{.Id}} size: {{.Size}}\" fedora"),
			ox.Flags().
				String("format", "Format the output to a Go template or json", ox.Default("json"), ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("size", "Display total file size", ox.Short("s")).
				String("type", "Specify inspect-object type", ox.Default("all"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Kill one or more running containers with a specific signal\n\nDescription:\n  The main process inside each container specified will be sent SIGKILL, or any signal specified with option --signal."),
			ox.Usage("kill", "Kill one or more running containers with a specific signal"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman kill mywebserver\n  podman kill 860a4b23\n  podman kill --signal TERM ctrID"),
			ox.Flags().
				Bool("all", "Signal all running containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				String("signal", "Signal to send to the container", ox.Default("KILL"), ox.Short("s")),
		), ox.Sub(
			ox.Banner("Play containers, pods or volumes from a structured file\n\nDescription:\n  Play structured data (e.g., Kubernetes YAML) based on containers, pods or volumes."),
			ox.Usage("kube", "Play containers, pods or volumes from a structured file"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Deploy a podman container, pod, volume, or Kubernetes yaml to a Kubernetes cluster\n\nDescription:\n  Command applies a podman container, pod, volume, or kube yaml to a Kubernetes cluster when a kubeconfig file is given."),
				ox.Usage("apply", "Deploy a podman container, pod, volume, or Kubernetes yaml to a Kubernetes cluster"),
				ox.Spec("[options] [CONTAINER...|POD...|VOLUME...]"),
				ox.Example("\n  podman kube apply ctrName volName\n  podman kube apply --namespace project -f fileName"),
				ox.Flags().
					String("ca-cert-file", "Path to the CA cert file for the Kubernetes cluster.").
					String("file", "Path to the Kubernetes yaml file to deploy.", ox.Short("f")).
					String("kubeconfig", "Path to the kubeconfig file for the Kubernetes cluster", ox.Short("k")).
					String("ns", "The namespace to deploy the workload to on the Kubernetes cluster").
					Bool("service", "Create a service object for the container being deployed.", ox.Short("s")),
			), ox.Sub(
				ox.Banner("Remove pods based on Kubernetes YAML\n\nDescription:\n  Reads in a structured file of Kubernetes YAML.\n\n  Removes pods that have been based on the Kubernetes kind described in the YAML."),
				ox.Usage("down", "Remove pods based on Kubernetes YAML"),
				ox.Spec("[options] KUBEFILE|-"),
				ox.Example("\n  podman kube down nginx.yml\n   cat nginx.yml | podman kube down -\n   podman kube down https://example.com/nginx.yml"),
				ox.Flags().
					Bool("force", "remove volumes"),
			), ox.Sub(
				ox.Banner("Generate Kubernetes YAML from containers, pods or volumes.\n\nDescription:\n  Command generates Kubernetes Pod, Service or PersistentVolumeClaim YAML (v1 specification) from Podman containers, pods or volumes.\n\n  Whether the input is for a container or pod, Podman will always generate the specification as a pod."),
				ox.Usage("generate", "Generate Kubernetes YAML from containers, pods or volumes."),
				ox.Spec("[options] {CONTAINER...|POD...|VOLUME...}"),
				ox.Example("\n  podman kube generate ctrID\n  podman kube generate podID\n  podman kube generate --service podID\n  podman kube generate volumeName\n  podman kube generate ctrID podID volumeName --service"),
				ox.Flags().
					String("filename", "Write output to the specified path", ox.Short("f")).
					Bool("podman-only", "Add podman-only reserved annotations to the generated YAML file (Cannot be used by Kubernetes)").
					Int32("replicas", "Set the replicas number for Deployment kind", ox.Default("1"), ox.Short("r")).
					Bool("service", "Generate YAML for a Kubernetes service object", ox.Short("s")).
					String("type", "Generate YAML for the given Kubernetes kind", ox.Default("pod"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Play a pod or volume based on Kubernetes YAML\n\nDescription:\n  Reads in a structured file of Kubernetes YAML.\n\n  Creates pods or volumes based on the Kubernetes kind described in the YAML. Supported kinds are Pods, Deployments, DaemonSets, Jobs, and PersistentVolumeClaims."),
				ox.Usage("play", "Play a pod or volume based on Kubernetes YAML"),
				ox.Spec("[options] KUBEFILE|-"),
				ox.Example("\n  podman kube play nginx.yml\n  cat nginx.yml | podman kube play -\n  podman kube play --creds user:password --seccomp-profile-root /custom/path apache.yml\n  podman kube play https://example.com/nginx.yml"),
				ox.Flags().
					Slice("annotation", "Add annotations to pods (key=value)", ox.Elem(ox.StringT)).
					String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					Bool("build", "Build all images in a YAML (given Containerfiles exist)").
					String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
					String("configmap", "Pathname of a YAML file containing a kubernetes configmap", ox.Spec("Pathname")).
					String("context-dir", "Path to top level of context directory").
					Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
					Bool("force", "Remove volumes as part of --down").
					Slice("ip", "Static IP addresses to assign to the pods", ox.Elem(ox.AddrT), ox.Default("[]")).
					String("log-driver", "Logging driver for the container", ox.Default("journald")).
					Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
					Slice("mac-address", "Static MAC addresses to assign to the pods", ox.Elem(ox.StringT)).
					Slice("network", "Connect pod to network(s) or network mode", ox.Elem(ox.StringT)).
					Bool("no-hosts", "Do not create /etc/hosts within the pod's containers, instead use the version from the image").
					Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT)).
					Bool("publish-all", "Whether to publish all ports defined in the K8S YAML file (containerPort, hostPort), if false only hostPort will be published").
					Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
					Bool("replace", "Delete and recreate pods defined in the YAML file").
					String("seccomp-profile-root", "Directory path for seccomp profiles", ox.Default("/var/lib/kubelet/seccomp")).
					Bool("start", "Start the pod after creating it", ox.Default("true")).
					Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")).
					String("userns", "User namespace to use").
					Bool("wait", "Clean up all objects created when a SIGTERM is received or pods exit", ox.Short("w")),
			)), ox.Sub(
			ox.Banner("Load image(s) from a tar archive\n\nDescription:\n  Loads an image from a locally stored archive (tar file) into container storage."),
			ox.Usage("load", "Load image(s) from a tar archive"),
			ox.Spec("[options]"),
			ox.Flags().
				String("input", "Read from specified archive file (default: stdin)", ox.Short("i")).
				Bool("quiet", "Suppress the output", ox.Short("q")),
		), ox.Sub(
			ox.Banner("Log in to a container registry\n\nDescription:\n  Log in to a container registry on a specified server."),
			ox.Usage("login", "Log in to a container registry"),
			ox.Spec("[options] [REGISTRY]"),
			ox.Example("\n  podman login quay.io\n  podman login --username ... --password ... quay.io\n  podman login --authfile dir/auth.json quay.io"),
			ox.Flags().
				String("authfile", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("cert-dir", "use certificates at the specified path to access the registry").
				String("compat-auth-file", "path of a Docker-compatible config file to update instead").
				Bool("get-login", "Return the current login user for the registry").
				String("password", "Password for registry", ox.Short("p")).
				Bool("password-stdin", "Take the password from stdin").
				String("secret", "Retrieve password from a podman secret").
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries").
				String("username", "Username for registry", ox.Short("u")).
				Bool("verbose", "Write more detailed information to stdout", ox.Short("v")),
		), ox.Sub(
			ox.Banner("Log out of a container registry\n\nDescription:\n  Remove the cached username and password for the registry."),
			ox.Usage("logout", "Log out of a container registry"),
			ox.Spec("[options] [REGISTRY]"),
			ox.Example("\n  podman logout quay.io\n  podman logout --authfile dir/auth.json quay.io\n  podman logout --all"),
			ox.Flags().
				Bool("all", "Remove the cached credentials for all registries in the auth file", ox.Short("a")).
				String("authfile", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("compat-auth-file", "path of a Docker-compatible config file to update instead"),
		), ox.Sub(
			ox.Banner("Fetch the logs of one or more containers\n\nDescription:\n  Retrieves logs for one or more containers.\n\n  This does not guarantee execution order when combined with podman run (i.e., your run may not have generated any logs at the time you execute podman logs)."),
			ox.Usage("logs", "Fetch the logs of one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman logs ctrID\n  podman logs --names ctrID1 ctrID2\n  podman logs --tail 2 mywebserver\n  podman logs --follow=true --since 10m ctrID\n  podman logs mywebserver mydbserver"),
			ox.Flags().
				Bool("color", "Output the containers with different colors in the log.").
				Bool("follow", "Follow log output.  The default is false", ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("names", "Output the container name in the log", ox.Short("n")).
				String("since", "Show logs since TIMESTAMP").
				Int("tail", "Output the specified number of LINES at the end of the logs.  Defaults to -1, which prints all lines", ox.Default("-1")).
				Bool("timestamps", "Output the timestamps in the log", ox.Short("t")).
				String("until", "Show logs until TIMESTAMP"),
		), ox.Sub(
			ox.Banner("Manage a virtual machine\n\nDescription:\n  Manage a virtual machine. Virtual machines are used to run Podman."),
			ox.Usage("machine", "Manage a virtual machine"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Display machine host info\n\nDescription:\n  Display information pertaining to the machine host."),
				ox.Usage("info", "Display machine host info"),
				ox.Spec("[options]"),
				ox.Example("\n  podman machine info"),
				ox.Flags().
					String("format", "Change the output format to JSON or a Go template", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Initialize a virtual machine\n\nDescription:\n  Initialize a virtual machine"),
				ox.Usage("init", "Initialize a virtual machine"),
				ox.Spec("[options] [NAME]"),
				ox.Example("\n  podman machine init podman-machine-default"),
				ox.Flags().
					Uint("cpus", "Number of CPUs", ox.Default("16")).
					Uint("disk-size", "Disk size in GiB", ox.Default("100")).
					String("ignition-path", "Path to ignition file").
					String("image", "Bootable image for machine").
					Uint("memory", "Memory in MiB", ox.Default("2048"), ox.Short("m")).
					Bool("now", "Start machine now").
					Bool("rootful", "Whether this machine should prefer rootful container execution").
					String("timezone", "Set timezone", ox.Default("local")).
					Slice("usb", "USB Host passthrough: bus=$1,devnum=$2 or vendor=$1,product=$2", ox.Elem(ox.StringT)).
					Bool("user-mode-networking", "Whether this machine should use user-mode networking, routing traffic through a host user-space process").
					String("username", "Username used in image", ox.Default("core")).
					Slice("volume", "Volumes to mount, source:target", ox.Elem(ox.StringT), ox.Default("[$HOME:$HOME]"), ox.Short("v")),
			), ox.Sub(
				ox.Banner("Inspect an existing machine\n\nDescription:\n  Provide details on a managed virtual machine"),
				ox.Usage("inspect", "Inspect an existing machine"),
				ox.Spec("[options] [MACHINE...]"),
				ox.Example("\n  podman machine inspect myvm"),
				ox.Flags().
					String("format", "Format volume output using JSON or a Go template"),
			), ox.Sub(
				ox.Banner("List machines\n\nDescription:\n  List managed virtual machines."),
				ox.Usage("list", "List machines"),
				ox.Spec("[options]"),
				ox.Aliases("list", "ls"),
				ox.Example("\n  podman machine list,\n  podman machine list --format json\n  podman machine ls"),
				ox.Flags().
					Bool("all-providers", "Show machines from all providers").
					String("format", "Format volume output using JSON or a Go template", ox.Default("{{range .}}{{.Name}}\\t{{.VMType}}\\t{{.Created}}\\t{{.LastUp}}\\t{{.CPUs}}\\t{{.Memory}}\\t{{.DiskSize}}\\n{{end -}}")).
					Bool("noheading", "Do not print headers", ox.Short("n")).
					Bool("quiet", "Show only machine names", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Manage a Podman virtual machine's OS\n\nDescription:\n  Manage a Podman virtual machine's operating system"),
				ox.Usage("os", "Manage a Podman virtual machine's OS"),
				ox.Spec("[command]"),
				ox.Sub(
					ox.Banner("Apply an OCI image to a Podman Machine's OS\n\nDescription:\n  Apply custom layers from a containerized Fedora CoreOS OCI image on top of an existing VM"),
					ox.Usage("apply", "Apply an OCI image to a Podman Machine's OS"),
					ox.Spec("[options] IMAGE [NAME]"),
					ox.Example("\n  podman machine os apply myimage"),
					ox.Flags().
						Bool("restart", "Restart VM to apply changes"),
				)), ox.Sub(
				ox.Banner("Remove all machines\n\nDescription:\n  Remove all machines, configurations, data, and cached images"),
				ox.Usage("reset", "Remove all machines"),
				ox.Spec("[options]"),
				ox.Example("\n  podman machine reset"),
				ox.Flags().
					Bool("force", "Do not prompt before reset", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Remove an existing machine\n\nDescription:\n  Remove a managed virtual machine"),
				ox.Usage("rm", "Remove an existing machine"),
				ox.Spec("[options] [MACHINE]"),
				ox.Example("\n  podman machine rm podman-machine-default"),
				ox.Flags().
					Bool("force", "Stop and do not prompt before rming", ox.Short("f")).
					Bool("save-ignition", "Do not delete ignition file").
					Bool("save-image", "Do not delete the image file"),
			), ox.Sub(
				ox.Banner("Set a virtual machine setting\n\nDescription:\n  Set an updatable virtual machine setting"),
				ox.Usage("set", "Set a virtual machine setting"),
				ox.Spec("[options] [NAME]"),
				ox.Example("\n  podman machine set --rootful=false"),
				ox.Flags().
					Uint("cpus", "Number of CPUs").
					Uint("disk-size", "Disk size in GiB").
					Uint("memory", "Memory in MiB", ox.Short("m")).
					Bool("rootful", "Whether this machine should prefer rootful container execution").
					Slice("usb", "USBs bus=$1,devnum=$2 or vendor=$1,product=$2", ox.Elem(ox.StringT)).
					Bool("user-mode-networking", "Whether this machine should use user-mode networking, routing traffic through a host user-space process"),
			), ox.Sub(
				ox.Banner("SSH into an existing machine\n\nDescription:\n  SSH into a managed virtual machine"),
				ox.Usage("ssh", "SSH into an existing machine"),
				ox.Spec("[options] [NAME] [COMMAND [ARG ...]]"),
				ox.Example("\n  podman machine ssh podman-machine-default\n  podman machine ssh myvm echo hello"),
				ox.Flags().
					String("username", "Username to use when ssh-ing into the VM."),
			), ox.Sub(
				ox.Banner("Start an existing machine\n\nDescription:\n  Start a managed virtual machine"),
				ox.Usage("start", "Start an existing machine"),
				ox.Spec("[options] [MACHINE]"),
				ox.Example("\n  podman machine start podman-machine-default"),
				ox.Flags().
					Bool("no-info", "Suppress informational tips").
					Bool("quiet", "Suppress machine starting status output", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Stop an existing machine\n\nDescription:\n  Stop a managed virtual machine"),
				ox.Usage("stop", "Stop an existing machine"),
				ox.Spec("[MACHINE]"),
				ox.Example("\n  podman machine stop podman-machine-default"),
			)), ox.Sub(
			ox.Banner("Manipulate manifest lists and image indexes\n\nDescription:\n  Creates, modifies, and pushes manifest lists and image indexes."),
			ox.Usage("manifest", "Manipulate manifest lists and image indexes"),
			ox.Spec("[command]"),
			ox.Example("\n  podman manifest add mylist:v1.11 image:v1.11-amd64\n  podman manifest create localhost/list\n  podman manifest inspect localhost/list\n  podman manifest annotate --annotation left=right mylist:v1.11 sha256:15352d97781ffdf357bf3459c037be3efac4133dc9070c2dce7eca7c05c3e736\n  podman manifest push mylist:v1.11 docker://quay.io/myuser/image:v1.11\n  podman manifest remove mylist:v1.11 sha256:15352d97781ffdf357bf3459c037be3efac4133dc9070c2dce7eca7c05c3e736\n  podman manifest rm mylist:v1.11"),
			ox.Sub(
				ox.Banner("Add images or artifacts to a manifest list or image index\n\nDescription:\n  Adds an image or artifact to a manifest list or image index."),
				ox.Usage("add", "Add images or artifacts to a manifest list or image index"),
				ox.Spec("[options] LIST IMAGEORARTIFACT [IMAGEORARTIFACT...]"),
				ox.Example("\n  podman manifest add mylist:v1.11 image:v1.11-amd64\n  podman manifest add mylist:v1.11 transport:imageName"),
				ox.Flags().
					Bool("all", "add all of the list's images if the image is a list").
					String("annotation", "set an annotation for the specified image", ox.Spec("annotation")).
					String("arch", "override the architecture of the specified image", ox.Spec("architecture")).
					Bool("artifact", "add all arguments as artifact files rather than as images").
					String("artifact-config", "artifact configuration file").
					String("artifact-config-type", "artifact configuration media type").
					Bool("artifact-exclude-titles", "refrain from setting \"org.opencontainers.image.title\" annotations on \"layers\"").
					String("artifact-layer-type", "artifact layer media type").
					String("artifact-subject", "artifact subject reference").
					String("artifact-type", "override the artifactType value").
					String("authfile", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "use certificates at the specified path to access the registry").
					String("creds", "use [username[:password]] for accessing the registry", ox.Spec("[username[:password]]")).
					Slice("features", "override the features of the specified image", ox.Elem(ox.StringT)).
					String("os", "override the OS of the specified image", ox.Spec("OS")).
					String("os-version", "override the OS version of the specified image", ox.Spec("version")).
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")).
					String("variant", "override the Variant of the specified image", ox.Spec("Variant")),
			), ox.Sub(
				ox.Banner("Add or update information about an entry in a manifest list or image index\n\nDescription:\n  Adds or updates information about an entry in a manifest list or image index."),
				ox.Usage("annotate", "Add or update information about an entry in a manifest list or image index"),
				ox.Spec("[options] LIST IMAGEORARTIFACT"),
				ox.Example("\n  podman manifest annotate --annotation left=right mylist:v1.11 sha256:15352d97781ffdf357bf3459c037be3efac4133dc9070c2dce7eca7c05c3e736"),
				ox.Flags().
					String("annotation", "set an annotation for the specified image or artifact", ox.Spec("annotation")).
					String("arch", "override the architecture of the specified image or artifact", ox.Spec("architecture")).
					Slice("features", "override the features of the specified image or artifact", ox.Elem(ox.StringT)).
					Bool("index", "apply --annotation values to the image index itself").
					String("os", "override the OS of the specified image or artifact", ox.Spec("OS")).
					Slice("os-features", "override the OS features of the specified image or artifact", ox.Elem(ox.StringT)).
					String("os-version", "override the OS version of the specified image or artifact", ox.Spec("version")).
					String("subject", "set the subject to which the image index refers", ox.Spec("subject")).
					String("variant", "override the Variant of the specified image or artifact", ox.Spec("Variant")),
			), ox.Sub(
				ox.Banner("Create manifest list or image index\n\nDescription:\n  Creates manifest lists or image indexes."),
				ox.Usage("create", "Create manifest list or image index"),
				ox.Spec("[options] LIST [IMAGE...]"),
				ox.Example("\n  podman manifest create mylist:v1.11\n  podman manifest create mylist:v1.11 arch-specific-image-to-add\n  podman manifest create mylist:v1.11 arch-specific-image-to-add another-arch-specific-image-to-add\n  podman manifest create --all mylist:v1.11 transport:tagged-image-to-add"),
				ox.Flags().
					Bool("all", "add all of the lists' images if the images to add are lists").
					Bool("amend", "modify an existing list if one with the desired name already exists", ox.Short("a")).
					Slice("annotation", "set annotations on the new list", ox.Elem(ox.StringT)).
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Check if a manifest list exists in local storage\n\nDescription:\n  If the manifest list exists in local storage, podman manifest exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if a manifest list exists in local storage"),
				ox.Spec("MANIFEST"),
				ox.Example("\n  podman manifest exists mylist"),
			), ox.Sub(
				ox.Banner("Display the contents of a manifest list or image index\n\nDescription:\n  Display the contents of a manifest list or image index."),
				ox.Usage("inspect", "Display the contents of a manifest list or image index"),
				ox.Spec("[options] IMAGE"),
				ox.Example("\n  podman manifest inspect localhost/list"),
				ox.Flags().
					String("authfile", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Push a manifest list or image index to a registry\n\nDescription:\n  Pushes manifest lists and image indexes to registries."),
				ox.Usage("push", "Push a manifest list or image index to a registry"),
				ox.Spec("[options] LIST DESTINATION"),
				ox.Example("\n  podman manifest push mylist:v1.11 docker://quay.io/myuser/image:v1.11"),
				ox.Flags().
					Slice("add-compression", "add instances with selected compression while pushing", ox.Elem(ox.StringT)).
					Bool("all", "also push the images in the list", ox.Default("true")).
					String("authfile", "path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
					String("cert-dir", "use certificates at the specified path to access the registry").
					String("compression-format", "compression format to use").
					Int("compression-level", "compression level to use").
					String("creds", "use [username[:password]] for accessing the registry", ox.Spec("[username[:password]]")).
					String("digestfile", "after copying the image, write the digest of the resulting digest to the file").
					Bool("force-compression", "Use the specified compression algorithm even if the destination contains a differently-compressed variant already").
					String("format", "manifest type (oci or v2s2) to attempt to use when pushing the manifest list", ox.Default("is manifest type of source"), ox.Short("f")).
					Bool("quiet", "don't output progress information when pushing lists", ox.Short("q")).
					Bool("remove-signatures", "don't copy signatures when pushing images").
					Bool("rm", "remove the manifest list if push succeeds").
					String("sign-by", "sign the image using a GPG key with the specified FINGERPRINT", ox.Spec("FINGERPRINT")).
					String("sign-by-sigstore", "Sign the image using a sigstore parameter file at PATH", ox.Spec("PATH")).
					String("sign-by-sigstore-private-key", "Sign the image using a sigstore private key at PATH", ox.Spec("PATH")).
					String("sign-passphrase-file", "Read a passphrase for signing an image from PATH", ox.Spec("PATH")).
					Bool("tls-verify", "require HTTPS and verify certificates when accessing the registry", ox.Default("true")),
			), ox.Sub(
				ox.Banner("Remove an item from a manifest list or image index\n\nDescription:\n  Removes an item from a manifest list or image index."),
				ox.Usage("remove", "Remove an item from a manifest list or image index"),
				ox.Spec("LIST DIGEST"),
				ox.Example("\n  podman manifest remove mylist:v1.11 sha256:15352d97781ffdf357bf3459c037be3efac4133dc9070c2dce7eca7c05c3e736"),
			), ox.Sub(
				ox.Banner("Remove manifest list or image index from local storage\n\nDescription:\n  Remove manifest list or image index from local storage."),
				ox.Usage("rm", "Remove manifest list or image index from local storage"),
				ox.Spec("LIST [LIST...]"),
				ox.Example("\n  podman manifest rm mylist:v1.11"),
			)), ox.Sub(
			ox.Banner("Mount a working container's root filesystem\n\nDescription:\n  podman mount\n    Lists all mounted containers mount points if no container is specified\n\n  podman mount CONTAINER-NAME-OR-ID\n    Mounts the specified container and outputs the mountpoint"),
			ox.Usage("mount", "Mount a working container's root filesystem"),
			ox.Spec("[options] [CONTAINER...]"),
			ox.Flags().
				Bool("all", "Mount all containers", ox.Short("a")).
				String("format", "Print the mounted containers in specified format (json)").
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("no-trunc", "Do not truncate output"),
		), ox.Sub(
			ox.Banner("Manage networks\n\nDescription:\n  Manage networks"),
			ox.Usage("network", "Manage networks"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Add container to a network\n\nDescription:\n  Add container to a network"),
				ox.Usage("connect", "Add container to a network"),
				ox.Spec("[options] NETWORK CONTAINER"),
				ox.Example("\n  podman network connect web secondary"),
				ox.Flags().
					Slice("alias", "network scoped alias for container", ox.Elem(ox.StringT)).
					Addr("ip", "set a static ipv4 address for this container network").
					Addr("ip6", "set a static ipv6 address for this container network").
					String("mac-address", "set a static mac address for this container network"),
			), ox.Sub(
				ox.Banner("Create networks for containers and pods\n\nDescription:\n  Create networks for containers and pods"),
				ox.Usage("create", "Create networks for containers and pods"),
				ox.Spec("[options] [NAME]"),
				ox.Example("\n  podman network create podman1"),
				ox.Flags().
					Bool("disable-dns", "disable dns plugin").
					Slice("dns", "DNS servers this network will use", ox.Elem(ox.StringT)).
					String("driver", "driver to manage the network", ox.Default("bridge"), ox.Short("d")).
					Slice("gateway", "IPv4 or IPv6 gateway for the subnet", ox.Elem(ox.AddrT), ox.Default("[]")).
					Bool("ignore", "Don't fail if network already exists").
					String("interface-name", "interface name which is used by the driver").
					Bool("internal", "restrict external access from this network").
					Slice("ip-range", "allocate container IP from range", ox.Elem(ox.StringT)).
					String("ipam-driver", "IP Address Management Driver").
					Bool("ipv6", "enable IPv6 networking").
					Slice("label", "set metadata on a network", ox.Elem(ox.StringT)).
					Slice("opt", "Set driver specific options", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("o")).
					Slice("route", "static routes", ox.Elem(ox.StringT)).
					Slice("subnet", "subnets in CIDR format", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Disconnect a container from a network\n\nDescription:\n  Remove container from a network"),
				ox.Usage("disconnect", "Disconnect a container from a network"),
				ox.Spec("[options] NETWORK CONTAINER"),
				ox.Example("\n  podman network disconnect web secondary"),
				ox.Flags().
					Bool("force", "force removal of container from network", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Check if network exists\n\nDescription:\n  If the named network exists, podman network exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if network exists"),
				ox.Spec("NETWORK"),
				ox.Example("\n  podman network exists net1"),
			), ox.Sub(
				ox.Banner("Inspect network\n\nDescription:\n  Displays the network configuration for one or more networks."),
				ox.Usage("inspect", "Inspect network"),
				ox.Spec("[options] NETWORK [NETWORK...]"),
				ox.Example("\n  podman network inspect podman"),
				ox.Flags().
					String("format", "Pretty-print network to JSON or using a Go template", ox.Short("f")),
			), ox.Sub(
				ox.Banner("List networks\n\nDescription:\n  List networks"),
				ox.Usage("ls", "List networks"),
				ox.Spec("[options]"),
				ox.Aliases("ls", "list"),
				ox.Example("\n  podman network list"),
				ox.Flags().
					Slice("filter", "Provide filter values (e.g. 'name=podman')", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Pretty-print networks to JSON or using a Go template").
					Bool("no-trunc", "Do not truncate the network ID").
					Bool("noheading", "Do not print headers", ox.Short("n")).
					Bool("quiet", "display only names", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Prune unused networks\n\nDescription:\n  Prune unused networks"),
				ox.Usage("prune", "Prune unused networks"),
				ox.Spec("[options]"),
				ox.Example("\n  podman network prune"),
				ox.Flags().
					Slice("filter", "Provide filter values (e.g. 'label=<key>=<value>')", ox.Elem(ox.StringT)).
					Bool("force", "do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Reload firewall rules for one or more containers\n\nDescription:\n  Reload container networks, recreating firewall rules"),
				ox.Usage("reload", "Reload firewall rules for one or more containers"),
				ox.Spec("[options] [CONTAINER...]"),
				ox.Example("\n  podman network reload 3c13ef6dd843\n  podman network reload test1 test2"),
				ox.Flags().
					Bool("all", "Reload network configuration of all containers", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Remove networks\n\nDescription:\n  Remove networks"),
				ox.Usage("rm", "Remove networks"),
				ox.Spec("[options] NETWORK [NETWORK...]"),
				ox.Aliases("rm", "remove"),
				ox.Example("\n  podman network rm podman"),
				ox.Flags().
					Bool("force", "remove any containers using network", ox.Short("f")).
					Int("time", "Seconds to wait for running containers to stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Update an existing podman network\n\nDescription:\n  Update an existing podman network"),
				ox.Usage("update", "Update an existing podman network"),
				ox.Spec("[options] NETWORK"),
				ox.Example("\n  podman network update podman1"),
				ox.Flags().
					Slice("dns-add", "add network level nameservers", ox.Elem(ox.StringT)).
					Slice("dns-drop", "remove network level nameservers", ox.Elem(ox.StringT)),
			)), ox.Sub(
			ox.Banner("Pause all the processes in one or more containers\n\nDescription:\n  Pauses one or more running containers.  The container name or ID can be used."),
			ox.Usage("pause", "Pause all the processes in one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman pause mywebserver\n  podman pause 860a4b23\n  podman pause --all"),
			ox.Flags().
				Bool("all", "Pause all running containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Manage pods\n\nDescription:\n  Pods are a group of one or more containers sharing the same network, pid and ipc namespaces."),
			ox.Usage("pod", "Manage pods"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Clone an existing pod\n\nDescription:\n  Create an exact copy of a pod and the containers within it"),
				ox.Usage("clone", "Clone an existing pod"),
				ox.Spec("[options] POD NAME"),
				ox.Example("\n  podman pod clone pod_name new_name"),
				ox.Flags().
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					String("cgroup-parent", "Optional parent cgroup for the container").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Bool("destroy", "destroy the original pod").
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
					String("infra-command", "Overwrite the default ENTRYPOINT of the image").
					String("infra-conmon-pidfile", "Path to the file that will receive the PID of conmon").
					String("infra-name", "Assign a name to the container").
					Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					String("name", "name the new pod", ox.Short("n")).
					String("pid", "PID namespace to use").
					String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
					String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					Bool("start", "start the new pod").
					String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
					String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
					Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
					Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
					Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Create a new empty pod\n\nDescription:\n  After creating the pod, the pod ID is printed to stdout.\n\n  You can then start it at any time with the  podman pod start <pod_id> command. The pod will be created with the initial state 'created'."),
				ox.Usage("create", "Create a new empty pod"),
				ox.Spec("[options] [NAME]"),
				ox.Example("\n  podman pod create\n  podman pod create --label foo=bar mypod"),
				ox.Flags().
					Slice("add-host", "Add a custom host-to-IP mapping (host:ip)", ox.Elem(ox.StringT), ox.Default("[]")).
					String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
					String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
					String("cgroup-parent", "Optional parent cgroup for the container").
					Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
					Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
					String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
					String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
					Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
					Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
					Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
					Slice("dns-option", "Set custom DNS options", ox.Elem(ox.StringT)).
					Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
					String("exit-policy", "Behaviour when the last container exits", ox.Default("continue")).
					Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
					Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
					Bool("infra", "Create an infra container associated with the pod to share namespaces with", ox.Default("true")).
					String("infra-command", "Overwrite the default ENTRYPOINT of the image").
					String("infra-conmon-pidfile", "Path to the file that will receive the PID of conmon").
					String("infra-image", "Image to use to override builtin infra container").
					String("infra-name", "Assign a name to the container").
					String("ip", "Specify a static IPv4 address for the container").
					String("ip6", "Specify a static IPv6 address for the container").
					Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
					Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
					String("mac-address", "Container MAC address (e.g. 92:d0:c6:0a:29:33)").
					String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
					String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
					String("name", "Assign a name to the pod", ox.Short("n")).
					Slice("network", "Connect a container to a network", ox.Elem(ox.StringT)).
					Slice("network-alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
					Bool("no-hosts", "Do not create /etc/hosts within the container, instead use the version from the image").
					String("pid", "PID namespace to use").
					String("pod-id-file", "Write the pod ID to the file").
					Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
					Bool("replace", "If a pod with the same name exists, replace it").
					String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
					Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
					String("share", "A comma delimited list of kernel namespaces the pod will share", ox.Default("ipc,net,uts")).
					Bool("share-parent", "Set the pod's cgroup as the cgroup parent for all containers joining the pod", ox.Default("true")).
					String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
					String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
					String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
					String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
					Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
					Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
					String("userns", "User namespace to use").
					String("uts", "UTS namespace to use").
					Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
					Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Check if a pod exists in local storage\n\nDescription:\n  If the named pod exists in local storage, podman pod exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if a pod exists in local storage"),
				ox.Spec("POD"),
				ox.Example("\n  podman pod exists podID\n  podman pod exists mypod || podman pod create --name mypod"),
			), ox.Sub(
				ox.Banner("Display a pod configuration\n\nDescription:\n  Display the configuration for a pod by name or id\n\n\tBy default, this will render all results in a JSON array."),
				ox.Usage("inspect", "Display a pod configuration"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod inspect podID"),
				ox.Flags().
					String("format", "Format the output to a Go template or json", ox.Default("json"), ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Send the specified signal or SIGKILL to containers in pod\n\nDescription:\n  Signals are sent to the main process of each container inside the specified pod.\n\n  The default signal is SIGKILL, or any signal specified with option --signal."),
				ox.Usage("kill", "Send the specified signal or SIGKILL to containers in pod"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod kill podID\n  podman pod kill --signal TERM mywebserver"),
				ox.Flags().
					Bool("all", "Kill all containers in all pods", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					String("signal", "Signal to send to the containers in the pod", ox.Default("KILL"), ox.Short("s")),
			), ox.Sub(
				ox.Banner("Fetch logs for pod with one or more containers\n\nDescription:\n  Displays logs for pod with one or more containers."),
				ox.Usage("logs", "Fetch logs for pod with one or more containers"),
				ox.Spec("[options] POD"),
				ox.Example("\n  podman pod logs podID\n\t\tpodman pod logs -c ctrname podName\n\t\tpodman pod logs --tail 2 mywebserver\n\t\tpodman pod logs --follow=true --since 10m podID\n\t\tpodman pod logs mywebserver"),
				ox.Flags().
					Bool("color", "Output the containers within a pod with different colors in the log").
					String("container", "Filter logs by container name or id which belongs to pod", ox.Short("c")).
					Bool("follow", "Follow log output.", ox.Short("f")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("names", "Output container names instead of container IDs in the log", ox.Short("n")).
					String("since", "Show logs since TIMESTAMP").
					Int("tail", "Output the specified number of LINES at the end of the logs.", ox.Default("-1")).
					Bool("timestamps", "Output the timestamps in the log", ox.Short("t")).
					String("until", "Show logs until TIMESTAMP"),
			), ox.Sub(
				ox.Banner("Pause one or more pods\n\nDescription:\n  The pod name or ID can be used.\n\n  All running containers within each specified pod will then be paused."),
				ox.Usage("pause", "Pause one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod pause podID1 podID2\n  podman pod pause --all"),
				ox.Flags().
					Bool("all", "Pause all running pods", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Remove all stopped pods and their containers\n\nDescription:\n  podman pod prune Removes all exited pods"),
				ox.Usage("prune", "Remove all stopped pods and their containers"),
				ox.Spec("[options]"),
				ox.Example("\n  podman pod prune"),
				ox.Flags().
					Bool("force", "Do not prompt for confirmation.  The default is false", ox.Short("f")),
			), ox.Sub(
				ox.Banner("List pods\n\nDescription:\n  List all pods on system including their names, ids and current state."),
				ox.Usage("ps", "List pods"),
				ox.Spec("[options]"),
				ox.Aliases("ps", "ls", "list"),
				ox.Flags().
					Bool("ctr-ids", "Display the container UUIDs. If no-trunc is not set they will be truncated").
					Bool("ctr-names", "Display the container names").
					Bool("ctr-status", "Display the container status").
					Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Pretty-print pods to JSON or using a Go template").
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-trunc", "Do not truncate pod and container IDs").
					Bool("noheading", "Do not print headers", ox.Short("n")).
					Bool("ns", "Display namespace information of the pod").
					Bool("quiet", "Print the numeric IDs of the pods only", ox.Short("q")).
					String("sort", "Sort output by created, id, name, or number", ox.Default("created")),
			), ox.Sub(
				ox.Banner("Restart one or more pods\n\nDescription:\n  The pod ID or name can be used.\n\n  All of the containers within each of the specified pods will be restarted. If a container in a pod is not currently running it will be started."),
				ox.Usage("restart", "Restart one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod restart podID1 podID2\n  podman pod restart --all"),
				ox.Flags().
					Bool("all", "Restart all running pods", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Remove one or more pods\n\nDescription:\n  podman rm will remove one or more stopped pods and their containers from the host.\n\n  The pod name or ID can be used.  A pod with containers will not be removed without --force. If --force is specified, all containers will be stopped, then removed."),
				ox.Usage("rm", "Remove one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod rm mywebserverpod\n  podman pod rm -f 860a4b23\n  podman pod rm -f -a"),
				ox.Flags().
					Bool("all", "Remove all running pods", ox.Short("a")).
					Bool("force", "Force removal of a running pod by first stopping all containers, then removing all containers in the pod.  The default is false", ox.Short("f")).
					Bool("ignore", "Ignore errors when a specified pod is missing", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Slice("pod-id-file", "Read the pod ID from the file", ox.Elem(ox.StringT)).
					Int("time", "Seconds to wait for pod stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Start one or more pods\n\nDescription:\n  The pod name or ID can be used.\n\n  All containers defined in the pod will be started."),
				ox.Usage("start", "Start one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod start podID\n  podman pod start --all"),
				ox.Flags().
					Bool("all", "Restart all running pods", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Slice("pod-id-file", "Read the pod ID from the file", ox.Elem(ox.StringT)),
			), ox.Sub(
				ox.Banner("Display a live stream of resource usage statistics for the containers in one or more pods\n\nDescription:\n  Display the containers' resource-usage statistics of one or more running pod"),
				ox.Usage("stats", "Display a live stream of resource usage statistics for the containers in one or more pods"),
				ox.Spec("[options] [POD...]"),
				ox.Example("\n  podman pod stats\n  podman pod stats a69b23034235 named-pod\n  podman pod stats --all"),
				ox.Flags().
					Bool("all", "Provide stats for all pods", ox.Short("a")).
					String("format", "Pretty-print container statistics to JSON or using a Go template").
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Bool("no-reset", "Disable resetting the screen when streaming").
					Bool("no-stream", "Disable streaming stats and only pull the first result"),
			), ox.Sub(
				ox.Banner("Stop one or more pods\n\nDescription:\n  The pod name or ID can be used.\n\n  This command will stop all running containers in each of the specified pods."),
				ox.Usage("stop", "Stop one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod stop mywebserverpod\n  podman pod stop --time 0 490eb 3557fb"),
				ox.Flags().
					Bool("all", "Stop all running pods", ox.Short("a")).
					Bool("ignore", "Ignore errors when a specified pod is missing", ox.Short("i")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
					Slice("pod-id-file", "Write the pod ID to the file", ox.Elem(ox.StringT)).
					Int("time", "Seconds to wait for pod stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Display the running processes of containers in a pod\n\nDescription:\n  Specify format descriptors to alter the output.\n\n  You may run \"podman pod top -l pid pcpu seccomp\" to print the process ID, the CPU percentage and the seccomp mode of each process of the latest pod.\n\n  Format Descriptors:\n    args,capamb,capbnd,capeff,capinh,capprm,comm,etime,group,groups,hgroup,hgroups,hpid,huid,huser,label,nice,pcpu,pgid,pid,ppid,rgroup,rss,ruser,seccomp,state,stime,time,tty,uid,user,vsz"),
				ox.Usage("top", "Display the running processes of containers in a pod"),
				ox.Spec("[options] POD [FORMAT-DESCRIPTORS|ARGS...]"),
				ox.Example("\n  podman pod top podID\npodman pod top podID pid seccomp args %C\npodman pod top podID -eo user,pid,comm"),
				ox.Flags().
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			), ox.Sub(
				ox.Banner("Unpause one or more pods\n\nDescription:\n  The podman unpause command will unpause all \"paused\" containers assigned to the pod.\n\n  The pod name or ID can be used."),
				ox.Usage("unpause", "Unpause one or more pods"),
				ox.Spec("[options] POD [POD...]"),
				ox.Example("\n  podman pod unpause podID1 podID2\n  podman pod unpause --all"),
				ox.Flags().
					Bool("all", "Unpause all running pods", ox.Short("a")).
					Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
			)), ox.Sub(
			ox.Banner("List port mappings or a specific mapping for the container\n\nDescription:\n  List port mappings for the CONTAINER, or look up the public-facing port that is NAT-ed to the PRIVATE_PORT"),
			ox.Usage("port", "List port mappings or a specific mapping for the container"),
			ox.Spec("[options] CONTAINER [PORT]"),
			ox.Example("\n  podman port --all\n  podman port ctrID 80/tcp"),
			ox.Flags().
				Bool("all", "Display port information for all containers", ox.Short("a")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("List containers\n\nDescription:\n  Prints out information about the containers"),
			ox.Usage("ps", "List containers"),
			ox.Spec("[options]"),
			ox.Example("\n  podman ps -a\n  podman ps -a --format \"{{.ID}}  {{.Image}}  {{.Labels}}  {{.Mounts}}\"\n  podman ps --size --sort names"),
			ox.Flags().
				Bool("all", "Show all the containers, default is only running containers", ox.Short("a")).
				Bool("external", "Show containers in storage not controlled by Podman").
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				String("format", "Pretty-print containers to JSON or using a Go template").
				Int("last", "Print the n last created containers (all states)", ox.Default("-1"), ox.Short("n")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("no-trunc", "Display the extended information").
				Bool("noheading", "Do not print headers").
				Bool("ns", "Display namespace information").
				Bool("pod", "Print the ID and name of the pod the containers are associated with", ox.Short("p")).
				Bool("quiet", "Print the numeric IDs of the containers only", ox.Short("q")).
				Bool("size", "Display the total file sizes", ox.Short("s")).
				String("sort", "Sort output by: command, created, id, image, names, runningfor, size, status", ox.Spec("choice")).
				Bool("sync", "Sync container state with OCI runtime").
				Uint("watch", "Watch the ps output on an interval in seconds", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Pull an image from a registry\n\nDescription:\n  Pulls an image from a registry and stores it locally.\n\n  An image can be pulled by tag or digest. If a tag is not specified, the image with the 'latest' tag is pulled."),
			ox.Usage("pull", "Pull an image from a registry"),
			ox.Spec("[options] IMAGE [IMAGE...]"),
			ox.Example("\n  podman pull imageName\n  podman pull fedora:latest"),
			ox.Flags().
				Bool("all-tags", "All tagged images in the repository will be pulled", ox.Short("a")).
				String("arch", "Use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
				String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
				Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
				Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
				Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
				String("os", "Use OS instead of the running OS for choosing images", ox.Spec("OS")).
				String("platform", "Specify the platform for selecting the image.  (Conflicts with arch and os)").
				Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
				Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
				String("retry-delay", "delay between retries in case of pull failures").
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")).
				String("variant", "Use VARIANT instead of the running architecture variant for choosing images"),
		), ox.Sub(
			ox.Banner("Push an image to a specified destination\n\nDescription:\n  Pushes a source image to a specified destination.\n\n\tThe Image \"DESTINATION\" uses a \"transport\":\"details\" format. See podman-push(1) section \"DESTINATION\" for the expected format."),
			ox.Usage("push", "Push an image to a specified destination"),
			ox.Spec("[options] IMAGE [DESTINATION]"),
			ox.Example("\n  podman push imageID docker://registry.example.com/repository:tag\n\t\tpodman push imageID oci-archive:/path/to/layout:image:tag"),
			ox.Flags().
				String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("cert-dir", "Path to a directory containing TLS certificates and keys").
				Bool("compress", "Compress tarball image layers when pushing to a directory using the 'dir' transport.", ox.Default("is same compression type as source")).
				String("compression-format", "compression format to use", ox.Default("gzip")).
				Int("compression-level", "compression level to use").
				Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
				String("digestfile", "Write the digest of the pushed image to the specified file").
				Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
				Slice("encrypt-layer", "Layers to encrypt, 0-indexed layer indices with support for negative indexing (e.g. 0 is the first layer, -1 is the last layer). If not defined, will encrypt all layers if encryption-key flag is specified", ox.Elem(ox.IntT)).
				Slice("encryption-key", "Key with the encryption protocol to use to encrypt the image (e.g. jwe:/path/to/key.pem)", ox.Elem(ox.StringT)).
				Bool("force-compression", "Use the specified compression algorithm even if the destination contains a differently-compressed variant already").
				String("format", "Manifest type (oci, v2s2, or v2s1) to use in the destination", ox.Default("is manifest type of source, with fallbacks"), ox.Short("f")).
				Bool("quiet", "Suppress output information when pushing images", ox.Short("q")).
				Bool("remove-signatures", "Discard any pre-existing signatures in the image").
				Uint("retry", "number of times to retry in case of failure when performing push", ox.Default("3")).
				String("retry-delay", "delay between retries in case of push failures").
				String("sign-by", "Add a signature at the destination using the specified key").
				String("sign-by-sigstore", "Sign the image using a sigstore parameter file at PATH", ox.Spec("PATH")).
				String("sign-by-sigstore-private-key", "Sign the image using a sigstore private key at PATH", ox.Spec("PATH")).
				String("sign-passphrase-file", "Read a passphrase for signing an image from PATH", ox.Spec("PATH")).
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
		), ox.Sub(
			ox.Banner("Rename an existing container\n\nDescription:\n  The podman rename command allows you to rename an existing container"),
			ox.Usage("rename", "Rename an existing container"),
			ox.Spec("CONTAINER NAME"),
			ox.Example("\n  podman rename containerA newName"),
		), ox.Sub(
			ox.Banner("Restart one or more containers\n\nDescription:\n  Restarts one or more running containers. The container ID or name can be used.\n\n  A timeout before forcibly stopping can be set, but defaults to 10 seconds."),
			ox.Usage("restart", "Restart one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman restart ctrID\n  podman restart ctrID1 ctrID2"),
			ox.Flags().
				Bool("all", "Restart all non-running containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("running", "Restart only running containers").
				Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Remove one or more containers\n\nDescription:\n  Removes one or more containers from the host. The container name or ID can be used.\n\n  Command does not remove images. Running or unusable containers will not be removed without the -f option."),
			ox.Usage("rm", "Remove one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman rm imageID\n  podman rm mywebserver myflaskserver 860a4b23\n  podman rm --force --all\n  podman rm -f c684f0d469f2"),
			ox.Flags().
				Bool("all", "Remove all containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Bool("depend", "Remove container and all containers that depend on the selected container").
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT)).
				Bool("force", "Force removal of a running or unusable container", ox.Short("f")).
				Bool("ignore", "Ignore errors when a specified container is missing", ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")).
				Bool("volumes", "Remove anonymous volumes associated with the container", ox.Short("v")),
		), ox.Sub(
			ox.Banner("Remove one or more images from local storage\n\nDescription:\n  Removes one or more previously pulled or locally created images."),
			ox.Usage("rmi", "Remove one or more images from local storage"),
			ox.Spec("[options] IMAGE [IMAGE...]"),
			ox.Example("\n  podman rmi imageID\n  podman rmi --force alpine\n  podman rmi c4dfb1609ee2 93fd78260bd1 c0ed59d05ff7"),
			ox.Flags().
				Bool("all", "Remove all images", ox.Short("a")).
				Bool("force", "Force Removal of the image", ox.Short("f")).
				Bool("ignore", "Ignore errors if a specified image does not exist", ox.Short("i")).
				Bool("no-prune", "Do not remove dangling images"),
		), ox.Sub(
			ox.Banner("Run a command in a new container\n\nDescription:\n  Runs a command in a new container from the given image"),
			ox.Usage("run", "Run a command in a new container"),
			ox.Spec("[options] IMAGE [COMMAND [ARG...]]"),
			ox.Example("\n  podman run imageID ls -alF /etc\n  podman run --network=host imageID dnf -y install java\n  podman run --volume /var/hostdir:/var/ctrdir -i -t fedora /bin/bash"),
			ox.Flags().
				Slice("add-host", "Add a custom host-to-IP mapping (host:ip)", ox.Elem(ox.StringT), ox.Default("[]")).
				Slice("annotation", "Add annotations to container (key=value)", ox.Elem(ox.StringT)).
				String("arch", "use ARCH instead of the architecture of the machine for choosing images", ox.Spec("ARCH")).
				Slice("attach", "Attach to STDIN, STDOUT or STDERR", ox.Elem(ox.StringT), ox.Short("a")).
				String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
				String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
				Slice("cap-add", "Add capabilities to the container", ox.Elem(ox.StringT)).
				Slice("cap-drop", "Drop capabilities from the container", ox.Elem(ox.StringT)).
				Slice("cgroup-conf", "Configure cgroup v2 (key=value)", ox.Elem(ox.StringT)).
				String("cgroup-parent", "Optional parent cgroup for the container").
				String("cgroupns", "cgroup namespace to use").
				String("cgroups", "control container cgroup configuration (\"enabled\"|\"disabled\"|\"no-conmon\"|\"split\")", ox.Default("enabled")).
				Slice("chrootdirs", "Chroot directories inside the container", ox.Elem(ox.StringT)).
				String("cidfile", "Write the container ID to the file").
				String("conmon-pidfile", "Path to the file that will receive the PID of conmon").
				Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
				Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
				Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
				Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
				Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
				Slice("decryption-key", "Key needed to decrypt the image (e.g. /path/to/key.pem)", ox.Elem(ox.StringT)).
				Bool("detach", "Run container in background and print container ID", ox.Short("d")).
				String("detach-keys", "Override the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-cf`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
				Slice("device", "Add a host device to the container", ox.Elem(ox.StringT)).
				Slice("device-cgroup-rule", "Add a rule to the cgroup allowed devices list", ox.Elem(ox.StringT)).
				Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				Bool("disable-content-trust", "This is a Docker specific option and is a NOOP").
				Slice("dns", "Set custom DNS servers", ox.Elem(ox.StringT)).
				Slice("dns-option", "Set custom DNS options", ox.Elem(ox.StringT)).
				Slice("dns-search", "Set custom DNS search domains", ox.Elem(ox.StringT)).
				String("entrypoint", "Overwrite the default ENTRYPOINT of the image").
				Slice("env", "Set environment variables in container", ox.Elem(ox.StringT), ox.Default("[PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]"), ox.Short("e")).
				Slice("env-file", "Read in a file of environment variables", ox.Elem(ox.StringT)).
				Bool("env-host", "Use all current host environment variables in container").
				Slice("env-merge", "Preprocess environment variables from image before injecting them into the container", ox.Elem(ox.StringT)).
				Slice("expose", "Expose a port or a range of ports", ox.Elem(ox.StringT)).
				Slice("gidmap", "GID map to use for the user namespace", ox.Elem(ox.StringT)).
				Slice("gpus", "GPU devices to add to the container ('all' to pass all GPUs)", ox.Elem(ox.StringT)).
				Slice("group-add", "Add additional groups to the primary container process. 'keep-groups' allows container processes to use supplementary groups.", ox.Elem(ox.StringT)).
				String("group-entry", "Entry to write to /etc/group").
				String("health-cmd", "set a healthcheck command for the container ('none' disables the existing healthcheck)").
				String("health-interval", "set an interval for the healthcheck (a value of disable results in no automatic timer setup)", ox.Default("30s")).
				String("health-log-destination", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file)", ox.Default("local")).
				Uint("health-max-log-count", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)", ox.Default("5")).
				Uint("health-max-log-size", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)", ox.Default("500")).
				String("health-on-failure", "action to take once the container turns unhealthy", ox.Default("none")).
				Uint("health-retries", "the number of retries allowed before a healthcheck is considered to be unhealthy", ox.Default("3")).
				String("health-start-period", "the initialization time needed for a container to bootstrap", ox.Default("0s")).
				String("health-startup-cmd", "Set a startup healthcheck command for the container").
				String("health-startup-interval", "Set an interval for the startup healthcheck", ox.Default("30s")).
				Uint("health-startup-retries", "Set the maximum number of retries before the startup healthcheck will restart the container").
				Uint("health-startup-success", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)").
				String("health-startup-timeout", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed", ox.Default("30s")).
				String("health-timeout", "the maximum time allowed to complete the healthcheck before an interval is considered failed", ox.Default("30s")).
				Slice("hostuser", "Host user account to add to /etc/passwd within container", ox.Elem(ox.StringT)).
				Bool("http-proxy", "Set proxy environment variables in the container based on the host proxy vars", ox.Default("true")).
				String("image-volume", "Tells podman how to handle the builtin image volumes (\"bind\"|\"tmpfs\"|\"ignore\")", ox.Default("anonymous")).
				Bool("init", "Run an init binary inside the container that forwards signals and reaps processes").
				String("init-path", "Path to the container-init binary").
				Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
				String("ip", "Specify a static IPv4 address for the container").
				String("ip6", "Specify a static IPv6 address for the container").
				String("ipc", "IPC namespace to use").
				Slice("label", "Set metadata on container", ox.Elem(ox.StringT), ox.Short("l")).
				Slice("label-file", "Read in a line delimited file of labels", ox.Elem(ox.StringT)).
				String("log-driver", "Logging driver for the container", ox.Default("journald")).
				Slice("log-opt", "Logging driver options", ox.Elem(ox.StringT)).
				String("mac-address", "Container MAC address (e.g. 92:d0:c6:0a:29:33)").
				String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
				String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
				String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
				Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
				Slice("mount", "Attach a filesystem mount to the container", ox.Elem(ox.StringT)).
				String("name", "Assign a name to the container").
				Slice("network", "Connect a container to a network", ox.Elem(ox.StringT)).
				Slice("network-alias", "Add network-scoped alias for the container", ox.Elem(ox.StringT)).
				Bool("no-healthcheck", "Disable healthchecks on container").
				Bool("no-hosts", "Do not create /etc/hosts within the container, instead use the version from the image").
				Bool("oom-kill-disable", "Disable OOM Killer").
				Int("oom-score-adj", "Tune the host's OOM preferences (-1000 to 1000)").
				String("os", "use OS instead of the running OS for choosing images", ox.Spec("OS")).
				Bool("passwd", "add entries to /etc/passwd and /etc/group", ox.Default("true")).
				String("passwd-entry", "Entry to write to /etc/passwd").
				String("personality", "Configure execution domain using personality (e.g., LINUX/LINUX32)").
				String("pid", "PID namespace to use").
				String("pidfile", "Write the container process ID to the file").
				Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
				String("platform", "Specify the platform for selecting the image.  (Conflicts with --arch and --os)").
				String("pod", "Run container in an existing pod").
				String("pod-id-file", "Read the pod ID from the file").
				Slice("preserve-fd", "Pass a file descriptor into the container", ox.Elem(ox.UintT), ox.Default("[]")).
				Uint("preserve-fds", "Pass a number of additional file descriptors into the container").
				Bool("privileged", "Give extended privileges to container").
				Slice("publish", "Publish a container's port, or a range of ports, to the host", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("p")).
				Bool("publish-all", "Publish all exposed ports to random ports on the host interface", ox.Short("P")).
				String("pull", "Pull image policy (\"always\"|\"missing\"|\"never\"|\"newer\")", ox.Default("missing")).
				Bool("quiet", "Suppress output information when pulling images", ox.Short("q")).
				String("rdt-class", "Class of Service (COS) that the container should be assigned to").
				Bool("read-only", "Make containers root filesystem read-only").
				Bool("read-only-tmpfs", "When running --read-only containers mount read-write tmpfs on /dev, /dev/shm, /run, /tmp and /var/tmp", ox.Default("true")).
				Bool("replace", "If a container with the same name exists, replace it").
				Slice("requires", "Add one or more requirement containers that must be started before this container will start", ox.Elem(ox.StringT)).
				String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")").
				Uint("retry", "number of times to retry in case of failure when performing pull", ox.Default("3")).
				String("retry-delay", "delay between retries in case of pull failures").
				Bool("rm", "Remove container and any anonymous unnamed volume associated with the container after exit").
				Bool("rmi", "Remove image unless used by other containers, implies --rm").
				Bool("rootfs", "The first argument is not an image but the rootfs to the exploded container").
				String("sdnotify", "control sd-notify behavior (\"container\"|\"conmon\"|\"healthy\"|\"ignore\")", ox.Default("container")).
				String("seccomp-policy", "Policy for selecting a seccomp profile (experimental)", ox.Default("default")).
				Slice("secret", "Add secret to container", ox.Elem(ox.StringT)).
				Slice("security-opt", "Security Options", ox.Elem(ox.StringT)).
				String("shm-size", "Size of /dev/shm (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Default("65536k")).
				String("shm-size-systemd", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
				Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true")).
				String("stop-signal", "Signal to stop a container. Default is SIGTERM").
				Uint("stop-timeout", "Timeout (in seconds) that containers stopped by user command have to exit. If exceeded, the container will be forcibly stopped via SIGKILL.", ox.Default("10")).
				String("subgidname", "Name of range listed in /etc/subgid for use in user namespace").
				String("subuidname", "Name of range listed in /etc/subuid for use in user namespace").
				Slice("sysctl", "Sysctl options", ox.Elem(ox.StringT)).
				String("systemd", "Run container in systemd mode (\"true\"|\"false\"|\"always\")", ox.Default("true")).
				Uint("timeout", "Maximum length of time a container is allowed to run. The container will be killed automatically after the time expires.").
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries for pulling images").
				String("tmpfs", "Mount a temporary filesystem (tmpfs) into a container", ox.Spec("tmpfs")).
				Bool("tty", "Allocate a pseudo-TTY for container", ox.Short("t")).
				String("tz", "Set timezone in container").
				Slice("uidmap", "UID map to use for the user namespace", ox.Elem(ox.StringT)).
				Slice("ulimit", "Ulimit options", ox.Elem(ox.StringT)).
				String("umask", "Set umask in container", ox.Default("0022")).
				Slice("unsetenv", "Unset environment default variables in container", ox.Elem(ox.StringT)).
				Bool("unsetenv-all", "Unset all default environment variables in container").
				String("user", "Username or UID (format: <name|uid>[:<group|gid>])", ox.Short("u")).
				String("userns", "User namespace to use").
				String("uts", "UTS namespace to use").
				String("variant", "Use VARIANT instead of the running architecture variant for choosing images", ox.Spec("VARIANT")).
				Slice("volume", "Bind mount a volume into the container", ox.Elem(ox.StringT), ox.Short("v")).
				Slice("volumes-from", "Mount volumes from the specified container(s)", ox.Elem(ox.StringT)).
				String("workdir", "Working directory inside the container", ox.Short("w")),
		), ox.Sub(
			ox.Banner("Save image(s) to an archive\n\nDescription:\n  Save an image to docker-archive or oci-archive on the local machine. Default is docker-archive."),
			ox.Usage("save", "Save image(s) to an archive"),
			ox.Spec("[options] IMAGE [IMAGE...]"),
			ox.Example("\n  podman save --quiet -o myimage.tar imageID\n  podman save --format docker-dir -o ubuntu-dir ubuntu\n  podman save > alpine-all.tar alpine:latest"),
			ox.Flags().
				Bool("compress", "Compress tarball image layers when saving to a directory using the 'dir' transport.", ox.Default("is same compression type as source")).
				String("format", "Save image to oci-archive, oci-dir (directory with oci manifest type), docker-archive, docker-dir (directory with v2s2 manifest type)", ox.Default("docker-archive")).
				Bool("multi-image-archive", "Interpret additional arguments as images not tags and create a multi-image-archive (only for docker-archive)", ox.Short("m")).
				String("output", "Write to a specified file (default: stdout, which must be redirected)", ox.Short("o")).
				Bool("quiet", "Suppress the output", ox.Short("q")).
				Bool("uncompressed", "Accept uncompressed layers when copying OCI images"),
		), ox.Sub(
			ox.Banner("Search registry for image\n\nDescription:\n  Search registries for a given image. Can search all the default registries or a specific registry.\n\n\tUsers can limit the number of results, and filter the output based on certain conditions."),
			ox.Usage("search", "Search registry for image"),
			ox.Spec("[options] TERM"),
			ox.Example("\n  podman search --filter=is-official --limit 3 alpine\n  podman search registry.fedoraproject.org/  # only works with v2 registries\n  podman search --format \"table {{.Index}} {{.Name}}\" registry.fedoraproject.org/fedora"),
			ox.Flags().
				String("authfile", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override").
				String("cert-dir", "Pathname of a directory containing TLS certificates and keys", ox.Spec("Pathname")).
				Bool("compatible", "List stars, official and automated columns (Docker compatibility)").
				Slice("creds", "Credentials (USERNAME:PASSWORD) to use for authenticating to a registry", ox.Elem(ox.StringT)).
				Slice("filter", "Filter output based on conditions provided", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("f")).
				String("format", "Change the output format to JSON or a Go template").
				Int("limit", "Limit the number of results").
				Bool("list-tags", "List the tags of the input registry").
				Bool("no-trunc", "Do not truncate the output").
				Bool("tls-verify", "Require HTTPS and verify certificates when contacting registries", ox.Default("true")),
		), ox.Sub(
			ox.Banner("Manage secrets\n\nDescription:\n  Manage secrets"),
			ox.Usage("secret", "Manage secrets"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Create a new secret\n\nDescription:\n  Create a secret. Input can be a path to a file or \"-\" (read from stdin). Secret drivers \"file\" (default), \"pass\", and \"shell\" are available."),
				ox.Usage("create", "Create a new secret"),
				ox.Spec("[options] NAME FILE|-"),
				ox.Example("\n  podman secret create mysecret /path/to/secret\n\t\tprintf \"secretdata\" | podman secret create mysecret -"),
				ox.Flags().
					String("driver", "Specify secret driver", ox.Default("file"), ox.Short("d")).
					Map("driver-opts", "Specify driver specific options", ox.MapKey(ox.StringT), ox.Elem(ox.StringT), ox.Default("[]")).
					Bool("env", "Read secret data from environment variable").
					Slice("label", "Specify labels on the secret", ox.Elem(ox.StringT), ox.Short("l")).
					Bool("replace", "If a secret with the same name exists, replace it"),
			), ox.Sub(
				ox.Banner("Check if a secret exists in local storage\n\nDescription:\n  If the named secret exists in local storage, podman secret exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if a secret exists in local storage"),
				ox.Spec("SECRET"),
				ox.Example("\n  podman secret exists ID\n  podman secret exists SECRET || podman secret create SECRET <secret source>"),
			), ox.Sub(
				ox.Banner("Inspect a secret\n\nDescription:\n  Display detail information on one or more secrets"),
				ox.Usage("inspect", "Inspect a secret"),
				ox.Spec("[options] SECRET [SECRET...]"),
				ox.Example("\n  podman secret inspect MYSECRET"),
				ox.Flags().
					String("format", "Format inspect output using Go template", ox.Short("f")).
					Bool("pretty", "Print inspect output in human-readable format").
					Bool("showsecret", "Display the secret"),
			), ox.Sub(
				ox.Banner("List secrets\n\nDescription:"),
				ox.Usage("ls", "List secrets"),
				ox.Spec("[options]"),
				ox.Aliases("ls", "list"),
				ox.Example("\n  podman secret ls"),
				ox.Flags().
					Slice("filter", "Filter secret output", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Format volume output using Go template", ox.Default("{{range .}}{{.ID}}\\t{{.Name}}\\t{{.Driver}}\\t{{.CreatedAt}}\\t{{.UpdatedAt}}\\n{{end -}}")).
					Bool("noheading", "Do not print headers", ox.Short("n")).
					Bool("quiet", "Print secret IDs only", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Remove one or more secrets\n\nDescription:"),
				ox.Usage("rm", "Remove one or more secrets"),
				ox.Spec("[options] SECRET [SECRET...]"),
				ox.Example("\n  podman secret rm mysecret1 mysecret2"),
				ox.Flags().
					Bool("all", "Remove all secrets", ox.Short("a")).
					Bool("ignore", "Ignore errors when a specified secret is missing", ox.Short("i")),
			)), ox.Sub(
			ox.Banner("Start one or more containers\n\nDescription:\n  Starts one or more containers.  The container name or ID can be used."),
			ox.Usage("start", "Start one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman start 860a4b231279 5421ab43b45\n  podman start --interactive --attach imageID"),
			ox.Flags().
				Bool("all", "Start all containers regardless of their state or configuration").
				Bool("attach", "Attach container's STDOUT and STDERR", ox.Short("a")).
				String("detach-keys", "Select the key sequence for detaching a container. Format is a single character [a-Z] or a comma separated sequence of `ctrl-<value>`, where `<value>` is one of: `a-z`, `@`, `^`, `[`, `\\`, `]`, `^` or `_`", ox.Spec("glob"), ox.Default("ctrl-p,ctrl-q")).
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("interactive", "Make STDIN available to the contained process", ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("sig-proxy", "Proxy received signals to the process", ox.Default("true if attaching, false otherwise")),
		), ox.Sub(
			ox.Banner("Display a live stream of container resource usage statistics\n\nDescription:\n  Display percentage of CPU, memory, network I/O, block I/O and PIDs for one or more containers."),
			ox.Usage("stats", "Display a live stream of container resource usage statistics"),
			ox.Spec("[options] [CONTAINER...]"),
			ox.Example("\n  podman stats --all --no-stream\n  podman stats ctrID\n  podman stats --no-stream --format \"table {{.ID}} {{.Name}} {{.MemUsage}}\" ctrID"),
			ox.Flags().
				Bool("all", "Show all containers. Only running containers are shown by default. The default is false", ox.Short("a")).
				String("format", "Pretty-print container statistics to JSON or using a Go template").
				Int("interval", "Time in seconds between stats reports", ox.Default("5"), ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Bool("no-reset", "Disable resetting the screen between intervals").
				Bool("no-stream", "Disable streaming stats and only pull the first result, default setting is false").
				Bool("no-trunc", "Do not truncate output"),
		), ox.Sub(
			ox.Banner("Stop one or more containers\n\nDescription:\n  Stops one or more running containers.  The container name or ID can be used.\n\n  A timeout to forcibly stop the container can also be set but defaults to 10 seconds otherwise."),
			ox.Usage("stop", "Stop one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman stop ctrID\n  podman stop --time 2 mywebserver 6e534f14da9d"),
			ox.Flags().
				Bool("all", "Stop all running containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("ignore", "Ignore errors when a specified container is missing", ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")).
				Int("time", "Seconds to wait for stop before killing the container", ox.Default("10"), ox.Short("t")),
		), ox.Sub(
			ox.Banner("Manage podman\n\nDescription:\n  Manage podman"),
			ox.Usage("system", "Manage podman"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Check storage consistency\n\nDescription:\n  \n\tpodman system check\n\n        Check storage for consistency and remove anything that looks damaged"),
				ox.Usage("check", "Check storage consistency"),
				ox.Spec("[options]"),
				ox.Example("\n  podman system check"),
				ox.Flags().
					Bool("force", "Remove inconsistent images and containers", ox.Short("f")).
					Duration("max", "Maximum allowed age of unreferenced layers", ox.Default("24h0m0s"), ox.Short("m")).
					Bool("quick", "Skip time-consuming checks. The default is to include time-consuming checks", ox.Short("q")).
					Bool("repair", "Remove inconsistent images", ox.Short("r")),
			), ox.Sub(
				ox.Banner("Manage remote API service destinations\n\nDescription:\n  Manage remote API service destination information in podman configuration"),
				ox.Usage("connection", "Manage remote API service destinations"),
				ox.Spec("[command]"),
				ox.Sub(
					ox.Banner("Record destination for the Podman service\n\nDescription:\n  Add destination to podman configuration.\n  \"destination\" is one of the form:\n    [user@]hostname (will default to ssh)\n    ssh://[user@]hostname[:port][/path] (will obtain socket path from service, if not given.)\n    tcp://hostname:port (not secured)\n    unix://path (absolute path required)"),
					ox.Usage("add", "Record destination for the Podman service"),
					ox.Spec("[options] NAME DESTINATION"),
					ox.Example("\n  podman system connection add laptop server.fubar.com\n  podman system connection add --identity ~/.ssh/dev_rsa testing ssh://root@server.fubar.com:2222\n  podman system connection add --identity ~/.ssh/dev_rsa --port 22 production root@server.fubar.com\n  podman system connection add debug tcp://localhost:8080"),
					ox.Flags().
						Bool("default", "Set connection to be default", ox.Short("d")).
						String("identity", "path to SSH identity file").
						Int("port", "SSH port number for destination", ox.Default("22"), ox.Short("p")).
						String("socket-path", "path to podman socket on remote host.", ox.Default("'/run/$APPNAME/$APPNAME.sock' or '/run/user/{uid}/$APPNAME/$APPNAME.sock")),
				), ox.Sub(
					ox.Banner("Set named destination as default\n\nDescription:\n  Set named destination as default for the Podman service"),
					ox.Usage("default", "Set named destination as default"),
					ox.Spec("NAME"),
					ox.Example("\n  podman system connection default testing"),
				), ox.Sub(
					ox.Banner("List destination for the Podman service(s)\n\nDescription:\n  List destination information for the Podman service(s) in podman configuration"),
					ox.Usage("list", "List destination for the Podman service(s)"),
					ox.Spec("[options]"),
					ox.Aliases("list", "ls"),
					ox.Example("\n  podman system connection list\n  podman system connection ls\n  podman system connection ls --format=json"),
					ox.Flags().
						String("format", "Custom Go template for printing connections", ox.Short("f")).
						Bool("quiet", "Custom Go template for printing connections", ox.Short("q")),
				), ox.Sub(
					ox.Banner("Delete named destination\n\nDescription:\n  Delete named destination from podman configuration"),
					ox.Usage("remove", "Delete named destination"),
					ox.Spec("[options] NAME"),
					ox.Aliases("remove", "rm"),
					ox.Example("\n  podman system connection remove devl\n  podman system connection rm devl"),
					ox.Flags().
						Bool("all", "Remove all connections", ox.Short("a")),
				), ox.Sub(
					ox.Banner("Rename \"old\" to \"new\"\n\nDescription:\n  Rename destination for the Podman service from \"old\" to \"new\""),
					ox.Usage("rename", "Rename \"old\" to \"new\""),
					ox.Spec("OLD NEW"),
					ox.Aliases("rename", "mv"),
					ox.Example("\n  podman system connection rename laptop devl,\n  podman system connection mv laptop devl"),
				)), ox.Sub(
				ox.Banner("Show podman disk usage\n\nDescription:\n  \n\tpodman system df\n\n\tShow podman disk usage"),
				ox.Usage("df", "Show podman disk usage"),
				ox.Spec("[options]"),
				ox.Flags().
					String("format", "Pretty-print images using a Go template").
					Bool("verbose", "Show detailed information on disk usage", ox.Short("v")),
			), ox.Sub(
				ox.Banner("Show podman system events\n\nDescription:\n  Monitor podman system events.\n\n  By default, streaming mode is used, printing new events as they occur.  Previous events can be listed via --since and --until."),
				ox.Usage("events", "Show podman system events"),
				ox.Spec("[options]"),
				ox.Example("\n  podman system events"),
				ox.Flags().
					Slice("filter", "filter output", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "format the output using a Go template").
					Bool("no-trunc", "do not truncate the output", ox.Default("true")).
					String("since", "show all events created since timestamp").
					Bool("stream", "stream events and do not exit when returning the last known event", ox.Default("true")).
					String("until", "show all events until timestamp"),
			), ox.Sub(
				ox.Banner("Display podman system information\n\nDescription:\n  Display information pertaining to the host, current storage stats, and build of podman.\n\n  Useful for the user and when reporting issues."),
				ox.Usage("info", "Display podman system information"),
				ox.Spec("[options]"),
				ox.Example("\n  podman system info"),
				ox.Flags().
					String("format", "Change the output format to JSON or a Go template", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Migrate containers\n\nDescription:\n  \n        podman system migrate\n\n        Migrate existing containers to a new version of Podman."),
				ox.Usage("migrate", "Migrate containers"),
				ox.Spec("[options]"),
				ox.Flags().
					String("new-runtime", "Specify a new runtime for all containers"),
			), ox.Sub(
				ox.Banner("Remove unused data\n\nDescription:\n  \n\tpodman system prune\n\n        Remove unused data"),
				ox.Usage("prune", "Remove unused data"),
				ox.Spec("[options]"),
				ox.Example("\n  podman system prune"),
				ox.Flags().
					Bool("all", "Remove all unused data", ox.Short("a")).
					Bool("external", "Remove container data in storage not controlled by podman").
					Slice("filter", "Provide filter values (e.g. 'label=<key>=<value>')", ox.Elem(ox.StringT)).
					Bool("force", "Do not prompt for confirmation.  The default is false", ox.Short("f")).
					Bool("volumes", "Prune volumes"),
			), ox.Sub(
				ox.Banner("Migrate lock numbers\n\nDescription:\n  \n        podman system renumber\n\n        Migrate lock numbers to handle a change in maximum number of locks.\n        Mandatory after the number of locks in containers.conf is changed."),
				ox.Usage("renumber", "Migrate lock numbers"),
			), ox.Sub(
				ox.Banner("Reset podman storage\n\nDescription:\n  Reset podman storage back to default state\n\n  All containers will be stopped and removed, and all images, volumes, networks and container content will be removed."),
				ox.Usage("reset", "Reset podman storage"),
				ox.Spec("[options]"),
				ox.Flags().
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Run API service\n\nDescription:\n  Run an API service\n\nEnable a listening service for API access to Podman commands."),
				ox.Usage("service", "Run API service"),
				ox.Spec("[options] [URI]"),
				ox.Example("\n  podman system service --time=0 unix:///tmp/podman.sock\n  podman system service --time=0 tcp://localhost:8888"),
				ox.Flags().
					String("cors", "Set CORS Headers").
					Uint("time", "Time until the service session expires in seconds.  Use 0 to disable the timeout", ox.Default("5"), ox.Short("t")),
			)), ox.Sub(
			ox.Banner("Add an additional name to a local image\n\nDescription:\n  Adds one or more additional names to locally-stored image."),
			ox.Usage("tag", "Add an additional name to a local image"),
			ox.Spec("IMAGE TARGET_NAME [TARGET_NAME...]"),
			ox.Example("\n  podman tag 0e3bbc2 fedora:latest\n  podman tag imageID:latest myNewImage:newTag\n  podman tag httpd myregistryhost:5000/fedora/httpd:v2"),
		), ox.Sub(
			ox.Banner("Display the running processes of a container\n\nDescription:\n  Display the running processes of a container.\n\n  The top command extends the ps(1) compatible AIX descriptors with container-specific ones as shown below.  In the presence of ps(1) specific flags (e.g, -eo), Podman will execute ps(1) inside the container.\n\n\n  Format Descriptors:\n    args,capamb,capbnd,capeff,capinh,capprm,comm,etime,group,groups,hgroup,hgroups,hpid,huid,huser,label,nice,pcpu,pgid,pid,ppid,rgroup,rss,ruser,seccomp,state,stime,time,tty,uid,user,vsz"),
			ox.Usage("top", "Display the running processes of a container"),
			ox.Spec("[options] CONTAINER [FORMAT-DESCRIPTORS|ARGS...]"),
			ox.Example("\n  podman top ctrID\npodman top ctrID pid seccomp args %C\npodman top ctrID -eo user,pid,comm"),
			ox.Flags().
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Unmount working container's root filesystem\n\nDescription:\n  Container storage increments a mount counter each time a container is mounted.\n\n  When a container is unmounted, the mount counter is decremented. The container's root filesystem is physically unmounted only when the mount counter reaches zero indicating no other processes are using the mount.\n\n  An unmount can be forced with the --force flag."),
			ox.Usage("unmount", "Unmount working container's root filesystem"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Aliases("unmount", "umount"),
			ox.Example("\n  podman unmount ctrID\n  podman unmount ctrID1 ctrID2 ctrID3\n  podman unmount --all"),
			ox.Flags().
				Bool("all", "Unmount all of the currently mounted containers", ox.Short("a")).
				Bool("force", "Force the complete unmount of the specified mounted containers", ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Unpause the processes in one or more containers\n\nDescription:\n  Unpauses one or more previously paused containers.  The container name or ID can be used."),
			ox.Usage("unpause", "Unpause the processes in one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman unpause ctrID\n  podman unpause --all"),
			ox.Flags().
				Bool("all", "Unpause all paused containers", ox.Short("a")).
				Slice("cidfile", "Read the container ID from the file", ox.Elem(ox.StringT)).
				Slice("filter", "Filter output based on conditions given", ox.Elem(ox.StringT), ox.Short("f")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Sub(
			ox.Banner("Run a command in a modified user namespace\n\nDescription:\n  Runs a command in a modified user namespace."),
			ox.Usage("unshare", "Run a command in a modified user namespace"),
			ox.Spec("[options] [COMMAND [ARG...]]"),
			ox.Example("\n  podman unshare id\n  podman unshare cat /proc/self/uid_map\n  podman unshare podman-script.sh"),
			ox.Flags().
				Bool("rootless-netns", "Join the rootless network namespace used for CNI and netavark networking"),
		), ox.Sub(
			ox.Banner("Remove a name from a local image\n\nDescription:\n  Removes one or more names from a locally-stored image."),
			ox.Usage("untag", "Remove a name from a local image"),
			ox.Spec("IMAGE [IMAGE...]"),
			ox.Example("\n  podman untag 0e3bbc2\n  podman untag imageID:latest otherImageName:latest\n  podman untag httpd myregistryhost:5000/fedora/httpd:v2"),
		), ox.Sub(
			ox.Banner("Update an existing container\n\nDescription:\n  Updates the cgroup configuration of a given container"),
			ox.Usage("update", "Update an existing container"),
			ox.Spec("[options] CONTAINER"),
			ox.Example("\n  podman update --cpus=5 foobar_container"),
			ox.Flags().
				String("blkio-weight", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.").
				String("blkio-weight-device", "Block IO weight (relative device weight, format: DEVICE_NAME:WEIGHT)", ox.Spec("DEVICE_NAME:WEIGHT")).
				Uint("cpu-period", "Limit the CPU CFS (Completely Fair Scheduler) period").
				Int("cpu-quota", "Limit the CPU CFS (Completely Fair Scheduler) quota").
				Uint("cpu-rt-period", "Limit the CPU real-time period in microseconds").
				Int("cpu-rt-runtime", "Limit the CPU real-time runtime in microseconds").
				Uint("cpu-shares", "CPU shares (relative weight)", ox.Short("c")).
				Float32("cpus", "Number of CPUs. The default is 0.000 which means no limit", ox.Spec("float")).
				String("cpuset-cpus", "CPUs in which to allow execution (0-3, 0,1)").
				String("cpuset-mems", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.").
				Slice("device-read-bps", "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-read-iops", "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				Slice("device-write-bps", "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)", ox.Elem(ox.StringT)).
				Slice("device-write-iops", "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)", ox.Elem(ox.StringT)).
				String("memory", "Memory limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]"), ox.Short("m")).
				String("memory-reservation", "Memory soft limit (format: <number>[<unit>], where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))", ox.Spec("<number>[<unit>]")).
				String("memory-swap", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap").
				Int("memory-swappiness", "Tune container memory swappiness (0 to 100, or -1 for system default)", ox.Default("-1")).
				Int("pids-limit", "Tune container pids limit (set -1 for unlimited)", ox.Default("2048")).
				String("restart", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")"),
		), ox.Sub(
			ox.Banner("Manage volumes\n\nDescription:\n  Volumes are created in and can be shared between containers"),
			ox.Usage("volume", "Manage volumes"),
			ox.Spec("[command]"),
			ox.Sub(
				ox.Banner("Create a new volume\n\nDescription:\n  If using the default driver, \"local\", the volume will be created on the host in the volumes directory under container storage."),
				ox.Usage("create", "Create a new volume"),
				ox.Spec("[options] [NAME]"),
				ox.Example("\n  podman volume create myvol\n  podman volume create\n  podman volume create --label foo=bar myvol"),
				ox.Flags().
					String("driver", "Specify volume driver name", ox.Default("local"), ox.Short("d")).
					Bool("ignore", "Don't fail if volume already exists").
					Slice("label", "Set metadata for a volume", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("l")).
					Slice("opt", "Set driver specific options", ox.Elem(ox.StringT), ox.Default("[]"), ox.Short("o")),
			), ox.Sub(
				ox.Banner("Check if volume exists\n\nDescription:\n  If the given volume exists, podman volume exists exits with 0, otherwise the exit code will be 1."),
				ox.Usage("exists", "Check if volume exists"),
				ox.Spec("VOLUME"),
				ox.Example("\n  podman volume exists myvol"),
			), ox.Sub(
				ox.Banner("Export volumes\n\nDescription:\n  \npodman volume export\n\nAllow content of volume to be exported into external tar."),
				ox.Usage("export", "Export volumes"),
				ox.Spec("[options] VOLUME"),
				ox.Flags().
					String("output", "Write to a specified file (default: stdout, which must be redirected)", ox.Default("/dev/stdout"), ox.Short("o")),
			), ox.Sub(
				ox.Banner("Import a tarball contents into a podman volume\n\nDescription:\n  Imports contents into a podman volume from specified tarball (.tar, .tar.gz, .tgz, .bzip, .tar.xz, .txz)."),
				ox.Usage("import", "Import a tarball contents into a podman volume"),
				ox.Spec("VOLUME [SOURCE]"),
				ox.Example("\n  podman volume import my_vol /home/user/import.tar\n  cat ctr.tar | podman volume import my_vol -"),
			), ox.Sub(
				ox.Banner("Display detailed information on one or more volumes\n\nDescription:\n  Display detailed information on one or more volumes.\n\n  Use a Go template to change the format from JSON."),
				ox.Usage("inspect", "Display detailed information on one or more volumes"),
				ox.Spec("[options] VOLUME [VOLUME...]"),
				ox.Example("\n  podman volume inspect myvol\n  podman volume inspect --all\n  podman volume inspect --format \"{{.Driver}} {{.Scope}}\" myvol"),
				ox.Flags().
					Bool("all", "Inspect all volumes", ox.Short("a")).
					String("format", "Format volume output using Go template", ox.Default("json"), ox.Short("f")),
			), ox.Sub(
				ox.Banner("List volumes\n\nDescription:\n  \npodman volume ls\n\nList all available volumes. The output of the volumes can be filtered\nand the output format can be changed to JSON or a user specified Go template."),
				ox.Usage("ls", "List volumes"),
				ox.Spec("[options]"),
				ox.Aliases("ls", "list"),
				ox.Flags().
					Slice("filter", "Filter volume output", ox.Elem(ox.StringT), ox.Short("f")).
					String("format", "Format volume output using Go template", ox.Default("{{range .}}{{.Driver}}\\t{{.Name}}\\n{{end -}}")).
					Bool("noheading", "Do not print headers", ox.Short("n")).
					Bool("quiet", "Print volume output in quiet mode", ox.Short("q")),
			), ox.Sub(
				ox.Banner("Mount volume\n\nDescription:\n  Mount a volume and return the mountpoint"),
				ox.Usage("mount", "Mount volume"),
				ox.Spec("NAME"),
				ox.Example("\n  podman volume mount myvol"),
			), ox.Sub(
				ox.Banner("Remove all unused volumes\n\nDescription:\n  Volumes that are not currently owned by a container will be removed.\n\n  The command prompts for confirmation which can be overridden with the --force flag.\n  Note all data will be destroyed."),
				ox.Usage("prune", "Remove all unused volumes"),
				ox.Spec("[options]"),
				ox.Flags().
					Slice("filter", "Provide filter values (e.g. 'label=<key>=<value>')", ox.Elem(ox.StringT)).
					Bool("force", "Do not prompt for confirmation", ox.Short("f")),
			), ox.Sub(
				ox.Banner("Reload all volumes from volume plugins\n\nDescription:\n  Check all configured volume plugins and update the libpod database with all available volumes.\n\n  Existing volumes are also removed from the database when they are no longer present in the plugin."),
				ox.Usage("reload", "Reload all volumes from volume plugins"),
			), ox.Sub(
				ox.Banner("Remove one or more volumes\n\nDescription:\n  Remove one or more existing volumes.\n\n  By default only volumes that are not being used by any containers will be removed. To remove the volumes anyways, use the --force flag."),
				ox.Usage("rm", "Remove one or more volumes"),
				ox.Spec("[options] VOLUME [VOLUME...]"),
				ox.Aliases("rm", "remove"),
				ox.Example("\n  podman volume rm myvol1 myvol2\n  podman volume rm --all\n  podman volume rm --force myvol"),
				ox.Flags().
					Bool("all", "Remove all volumes", ox.Short("a")).
					Bool("force", "Remove a volume by force, even if it is being used by a container", ox.Short("f")).
					Int("time", "Seconds to wait for running containers to stop before killing the container", ox.Default("10"), ox.Short("t")),
			), ox.Sub(
				ox.Banner("Unmount volume\n\nDescription:\n  Unmount a volume"),
				ox.Usage("unmount", "Unmount volume"),
				ox.Spec("NAME"),
				ox.Example("\n  podman volume unmount myvol"),
			)), ox.Sub(
			ox.Banner("Block on one or more containers\n\nDescription:\n  Block until one or more containers stop and then print their exit codes."),
			ox.Usage("wait", "Block on one or more containers"),
			ox.Spec("[options] CONTAINER [CONTAINER...]"),
			ox.Example("\n  podman wait --interval 5s ctrID\n  podman wait ctrID1 ctrID2"),
			ox.Flags().
				Slice("condition", "Condition to wait on", ox.Elem(ox.StringT)).
				Bool("ignore", "Ignore if a container does not exist").
				String("interval", "Time Interval to wait before polling for completion", ox.Default("250ms"), ox.Short("i")).
				Bool("latest", "Act on the latest container podman is aware of", ox.Short("l")),
		), ox.Flags().
			String("cgroup-manager", "Cgroup manager to use (\"cgroupfs\"|\"systemd\")", ox.Default("systemd")).
			String("config", "Location of authentication config file").
			String("conmon", "Path of the conmon binary").
			String("connection", "Connection to use for remote Podman service (CONTAINER_CONNECTION)", ox.Short("c")).
			String("events-backend", "Events backend to use (\"file\"|\"journald\"|\"none\")", ox.Default("journald")).
			Slice("hooks-dir", "Set the OCI hooks directory path (may be set multiple times)", ox.Elem(ox.StringT), ox.Default("[/usr/share/containers/oci/hooks.d]")).
			String("identity", "path to SSH identity file, (CONTAINER_SSHKEY)").
			String("imagestore", "Path to the 'image store', different from 'graph root', use this to split storing the image into a separate 'image store', see 'man containers-storage.conf' for details").
			String("log-level", "Log messages above specified level (trace, debug, info, warn, warning, error, fatal, panic)", ox.Default("warn")).
			Slice("module", "Load the containers.conf(5) module", ox.Elem(ox.StringT)).
			String("network-cmd-path", "Path to the command for configuring the network").
			String("network-config-dir", "Path of the configuration directory for networks").
			String("out", "Send output (stdout) from podman to a file").
			Bool("remote", "Access remote Podman service", ox.Short("r")).
			String("root", "Path to the graph root directory where images, containers, etc. are stored").
			String("runroot", "Path to the 'run directory' where all state information is stored").
			String("runtime", "Path to the OCI-compatible binary used to run containers.", ox.Default("crun")).
			Slice("runtime-flag", "add global flags for the container runtime", ox.Elem(ox.StringT)).
			String("ssh", "define the ssh mode", ox.Default("golang")).
			String("storage-driver", "Select which storage driver is used to manage storage of images and containers").
			Slice("storage-opt", "Used to pass an option to the storage driver", ox.Elem(ox.StringT)).
			Bool("syslog", "Output logging information to syslog as well as the console", ox.Default("false")).
			String("tmpdir", "Path to the tmp directory for libpod state content.").
			Bool("transient-store", "Enable transient container storage").
			String("url", "URL to access Podman service (CONTAINER_HOST)", ox.Default("unix:///run/user/1000/$APPNAME/$APPNAME.sock")).
			String("volumepath", "Path to the volume directory in which volume data is stored"),
	)
}
