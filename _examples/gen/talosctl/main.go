// Command talosctl is a xo/ox version of `talosctl`.
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
		ox.Usage(`talosctl`, ``),
		ox.Banner(`A CLI for out-of-band management of Kubernetes nodes created by Talos`),
		ox.Spec(`[command]`),
		ox.Footer(`Use "talosctl [command] --help" for more information about a command.`),
		ox.Flags().
			String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`).
			String(`context`, `Context to be used in command`).
			Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e")).
			Slice(`nodes`, `target the specified nodes`, ox.Short("n")).
			String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`),
		ox.Sub( // talosctl apply-config
			ox.Usage(`apply-config`, `Apply a new configuration to a node`),
			ox.Banner(`Apply a new configuration to a node`),
			ox.Spec(`[flags]`),
			ox.Aliases("apply"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Slice(`cert-fingerprint`, `list of server certificate fingeprints to accept (defaults to no check)`).
				Array(`config-patch`, `the list of config patches to apply to the local config file before sending it to the node`, ox.Short("p")).
				Bool(`dry-run`, `check how the config change will be applied in dry-run mode`).
				String(`file`, `the filename of the updated configuration`, ox.Short("f")).
				Bool(`insecure`, `apply the config using the insecure (encrypted with no auth) maintenance service`, ox.Short("i")).
				String(`mode`, `interactive, no-reboot, reboot, staged, try   apply config mode`, ox.Spec(`auto,`), ox.Default("auto"), ox.Short("m")).
				Duration(`timeout`, `the config will be rolled back after specified timeout (if try mode is selected)`, ox.Default("1m0s")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl bootstrap
			ox.Usage(`bootstrap`, `Bootstrap the etcd cluster on the specified node.`),
			ox.Banner(`When Talos cluster is created etcd service on control plane nodes enter the join loop waiting
to join etcd peers from other control plane nodes. One node should be picked as the bootstrap node.
When bootstrap command is issued, the node aborts join process and bootstraps etcd cluster as a single node cluster.
Other control plane nodes will join etcd cluster once Kubernetes is bootstrapped on the bootstrap node.

This command should not be used when "init" type node are used.

Talos etcd cluster can be recovered from a known snapshot with '--recover-from=' flag.`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`recover-from`, `recover etcd cluster from the snapshot`).
				Bool(`recover-skip-hash-check`, `skip integrity check when recovering etcd (use when recovering from data directory copy)`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl cgroups
			ox.Usage(`cgroups`, `Retrieve cgroups usage information`),
			ox.Banner(`The cgroups command fetches control group v2 (cgroupv2) usage details from the machine.
Several presets are available to focus on specific cgroup subsystems:

* cpu
* cpuset
* io
* memory
* process
* swap

You can specify the preset using the --preset flag.

Alternatively, a custom schema can be provided using the --schema-file flag.
To see schema examples, refer to https://github.com/siderolabs/talos/tree/main/cmd/talosctl/cmd/talos/cgroupsprinter/schemas.`),
			ox.Spec(`[flags]`),
			ox.Aliases("cg"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`preset`, `preset name (one of: [cpu cpuset io memory process swap])`).
				String(`schema-file`, `path to the columns schema file`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl cluster
			ox.Usage(`cluster`, `A collection of commands for managing local docker-based or QEMU-based clusters`),
			ox.Banner(`A collection of commands for managing local docker-based or QEMU-based clusters`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl cluster [command] --help" for more information about a command.`),
			ox.Flags().
				String(`name`, `the name of the cluster`, ox.Default("talos-default")).
				String(`provisioner`, `Talos cluster provisioner to use`, ox.Default("docker")).
				String(`state`, `directory path to store cluster state`, ox.Default("$HOME/.talos/clusters")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl cluster create
				ox.Usage(`create`, `Creates a local docker-based or QEMU-based kubernetes cluster`),
				ox.Banner(`Creates a local docker-based or QEMU-based kubernetes cluster`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`arch`, `cluster architecture`, ox.Default("amd64")).
					Bool(`bad-rtc`, `launch VM with bad RTC state (QEMU only)`).
					String(`cidr`, `CIDR of the cluster network (IPv4, ULA network for IPv6 is derived in automated way)`, ox.Default("10.5.0.0/24")).
					Slice(`cni-bin-path`, `search path for CNI binaries (VM only)`, ox.Default("[$HOME/.talos/cni/bin]")).
					String(`cni-bundle-url`, `URL to download CNI bundle from (VM only)`, ox.Default("https://github.com/siderolabs/talos/releases/download/v1.10.4/$APPNAME-cni-bundle-${ARCH}.tar.gz")).
					String(`cni-cache-dir`, `CNI cache directory path (VM only)`, ox.Default("$HOME/.talos/cni/cache")).
					String(`cni-conf-dir`, `CNI config directory path (VM only)`, ox.Default("$HOME/.talos/cni/conf.d")).
					String(`config-injection-method`, `a method to inject machine config: default is HTTP server, 'metal-iso' to mount an ISO (QEMU only)`).
					Array(`config-patch`, `patch generated machineconfigs (applied to all node types), use @file to read a patch from file`).
					Array(`config-patch-control-plane`, `patch generated machineconfigs (applied to 'init' and 'controlplane' types)`).
					Array(`config-patch-worker`, `patch generated machineconfigs (applied to 'worker' type)`).
					Int(`control-plane-port`, `control plane port (load balancer and local API port, QEMU only)`, ox.Default("6443")).
					Int(`controlplanes`, `the number of controlplanes to create`, ox.Default("1")).
					String(`cpus`, `the share of CPUs as fraction (each control plane/VM)`, ox.Default("2.0")).
					String(`cpus-workers`, `the share of CPUs as fraction (each worker/VM)`, ox.Default("2.0")).
					String(`custom-cni-url`, `install custom CNI from the URL (Talos cluster)`).
					Bool(`disable-dhcp-hostname`, `skip announcing hostname via DHCP (QEMU only)`).
					Int(`disk`, `default limit on disk size in MB (each VM)`, ox.Default("6144")).
					Uint(`disk-block-size`, `disk block size (VM only)`, ox.Default("512")).
					Array(`disk-encryption-key-types`, `encryption key types to use for disk encryption (uuid, kms)`, ox.Default("[uuid]")).
					String(`disk-image-path`, `disk image to use`).
					Bool(`disk-preallocate`, `whether disk space should be preallocated`, ox.Default("true")).
					String(`dns-domain`, `the dns domain to use for cluster`, ox.Default("cluster.local")).
					Bool(`docker-disable-ipv6`, `skip enabling IPv6 in containers (Docker only)`).
					String(`docker-host-ip`, `Host IP to forward exposed ports to (Docker provisioner only)`, ox.Default("0.0.0.0")).
					Bool(`encrypt-ephemeral`, `enable ephemeral partition encryption`).
					Bool(`encrypt-state`, `enable state partition encryption`).
					Bool(`encrypt-user-volumes`, `enable ephemeral partition encryption`).
					String(`endpoint`, `use endpoint instead of provider defaults`).
					String(`exposed-ports`, `Comma-separated list of ports/protocols to expose on init node. Ex -p <hostPort>:<containerPort>/<protocol (tcp or udp)> (Docker provisioner only)`, ox.Short("p")).
					String(`extra-boot-kernel-args`, `add extra kernel args to the initial boot from vmlinuz and initramfs (QEMU only)`).
					Int(`extra-disks`, `number of extra disks to create for each worker VM`).
					Slice(`extra-disks-drivers`, `driver for each extra disk (virtio, ide, ahci, scsi, nvme, megaraid)`).
					Int(`extra-disks-size`, `default limit on disk size in MB (each VM)`, ox.Default("5120")).
					Slice(`extra-uefi-search-paths`, `additional search paths for UEFI firmware (only applies when UEFI is enabled)`).
					String(`image`, `the image to use`, ox.Default("ghcr.io/siderolabs/talos:v1.10.4")).
					Bool(`init-node-as-endpoint`, `use init node as endpoint instead of any load balancer endpoint`).
					String(`initrd-path`, `initramfs image to use`, ox.Default("_out/initramfs-${ARCH}.xz")).
					String(`input-dir`, `location of pre-generated config files`, ox.Short("i")).
					String(`install-image`, `the installer image to use`, ox.Default("ghcr.io/siderolabs/installer:v1.10.4")).
					Bool(`ipv4`, `enable IPv4 network in the cluster`, ox.Default("true")).
					Bool(`ipv6`, `enable IPv6 network in the cluster (QEMU provisioner only)`).
					String(`ipxe-boot-script`, `iPXE boot script (URL) to use`).
					String(`iso-path`, `the ISO path to use for the initial boot (VM only)`).
					Int(`kubeprism-port`, `KubePrism port (set to 0 to disable)`, ox.Default("7445")).
					String(`kubernetes-version`, `desired kubernetes version to run`, ox.Default("1.33.1")).
					Int(`memory`, `the limit on memory usage in MB (each control plane/VM)`, ox.Default("2048")).
					Int(`memory-workers`, `the limit on memory usage in MB (each worker/VM)`, ox.Default("2048")).
					String(`mount`, `attach a mount to the container (Docker only)`, ox.Spec(`mount`)).
					Int(`mtu`, `MTU of the cluster network`, ox.Default("1500")).
					Slice(`nameservers`, `list of nameservers to use`, ox.Default("[8.8.8.8,1.1.1.1,2001:4860:4860::8888,2606:4700:4700::1111]")).
					Slice(`no-masquerade-cidrs`, `list of CIDRs to exclude from NAT (QEMU provisioner only)`).
					Slice(`registry-insecure-skip-verify`, `list of registry hostnames to skip TLS verification for`).
					Slice(`registry-mirror`, `list of registry mirrors to use in format: <registry host>=<mirror URL>`).
					Bool(`skip-injecting-config`, `skip injecting config from embedded metadata server, write config files to current directory`).
					Bool(`skip-k8s-node-readiness-check`, `skip k8s node readiness checks`).
					Bool(`skip-kubeconfig`, `skip merging kubeconfig from the created cluster`).
					String(`talos-version`, `the desired Talos version to generate config for (if not set, defaults to image version)`).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`).
					String(`uki-path`, `the UKI image path to use for the initial boot (VM only)`).
					String(`usb-path`, `the USB stick image path to use for the initial boot (VM only)`).
					Bool(`use-vip`, `use a virtual IP for the controlplane endpoint instead of the loadbalancer`).
					Slice(`user-volumes`, `list of user volumes to create for each VM in format: <name1>:<size1>:<name2>:<size2>`).
					String(`vmlinuz-path`, `the compressed kernel image to use`, ox.Default("_out/vmlinuz-${ARCH}")).
					Bool(`wait`, `wait for the cluster to be ready before returning`, ox.Default("true")).
					Duration(`wait-timeout`, `timeout to wait for the cluster to be ready`, ox.Default("20m0s")).
					String(`wireguard-cidr`, `CIDR of the wireguard network`).
					Bool(`with-apply-config`, `enable apply config when the VM is starting in maintenance mode`).
					Bool(`with-bootloader`, `enable bootloader to load kernel and initramfs from disk image after install`, ox.Default("true")).
					Bool(`with-cluster-discovery`, `enable cluster discovery`, ox.Default("true")).
					Bool(`with-debug`, `enable debug in Talos config to send service logs to the console`).
					String(`with-firewall`, `inject firewall rules into the cluster, value is default policy - accept/block (QEMU only)`).
					Bool(`with-init-node`, `create the cluster with an init node`).
					Bool(`with-iommu`, `enable IOMMU support, this also add a new PCI root port and an interface attached to it (qemu only)`).
					Bool(`with-json-logs`, `enable JSON logs receiver and configure Talos to send logs there`).
					Bool(`with-kubespan`, `enable KubeSpan system`).
					Int(`with-network-bandwidth`, `specify bandwidth restriction (in kbps) on the bridge interface when creating a qemu cluster`).
					Bool(`with-network-chaos`, `enable to use network chaos parameters when creating a qemu cluster`).
					Duration(`with-network-jitter`, `specify jitter on the bridge interface when creating a qemu cluster`).
					Duration(`with-network-latency`, `specify latency on the bridge interface when creating a qemu cluster`).
					Float32(`with-network-packet-corrupt`, `specify percent of corrupt packets on the bridge interface when creating a qemu cluster. e.g. 50% = 0.50 (default: 0.0)`, ox.Spec(`float`)).
					Float32(`with-network-packet-loss`, `specify percent of packet loss on the bridge interface when creating a qemu cluster. e.g. 50% = 0.50 (default: 0.0)`, ox.Spec(`float`)).
					Float32(`with-network-packet-reorder`, `specify percent of reordered packets on the bridge interface when creating a qemu cluster. e.g. 50% = 0.50 (default: 0.0)`, ox.Spec(`float`)).
					Bool(`with-siderolink`, `enables the use of siderolink agent as configuration apply mechanism. true or `+"`"+`wireguard`+"`"+` enables the agent, `+"`"+`tunnel`+"`"+` enables the agent with grpc tunneling`, ox.Spec(`true`), ox.Default("none")).
					Bool(`with-tpm1_2`, `enable TPM 1.2 emulation support using swtpm`).
					Bool(`with-tpm2`, `enable TPM 2.0 emulation support using swtpm`).
					Bool(`with-uefi`, `enable UEFI on x86_64 architecture`, ox.Default("true")).
					Bool(`with-uuid-hostnames`, `use machine UUIDs as default hostnames (QEMU only)`).
					Int(`workers`, `the number of workers to create`, ox.Default("1")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`name`, `the name of the cluster`, ox.Default("talos-default"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`provisioner`, `Talos cluster provisioner to use`, ox.Default("docker"), ox.Section(0)).
					String(`state`, `directory path to store cluster state`, ox.Default("$HOME/.talos/clusters"), ox.Section(0)),
			),
			ox.Sub( // talosctl cluster destroy
				ox.Usage(`destroy`, `Destroys a local docker-based or firecracker-based kubernetes cluster`),
				ox.Banner(`Destroys a local docker-based or firecracker-based kubernetes cluster`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `force deletion of cluster directory if there were errors`, ox.Short("f")).
					String(`save-cluster-logs-archive-path`, `save cluster logs archive to the specified file on destroy`).
					String(`save-support-archive-path`, `save support archive to the specified file on destroy`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`name`, `the name of the cluster`, ox.Default("talos-default"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`provisioner`, `Talos cluster provisioner to use`, ox.Default("docker"), ox.Section(0)).
					String(`state`, `directory path to store cluster state`, ox.Default("$HOME/.talos/clusters"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl cluster show
				ox.Usage(`show`, `Shows info about a local provisioned kubernetes cluster`),
				ox.Banner(`Shows info about a local provisioned kubernetes cluster`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`name`, `the name of the cluster`, ox.Default("talos-default"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`provisioner`, `Talos cluster provisioner to use`, ox.Default("docker"), ox.Section(0)).
					String(`state`, `directory path to store cluster state`, ox.Default("$HOME/.talos/clusters"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl config
			ox.Usage(`config`, `Manage the client configuration file (talosconfig)`),
			ox.Banner(`Manage the client configuration file (talosconfig)`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl config [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl config add
				ox.Usage(`add`, `Add a new context`),
				ox.Banner(`Add a new context`),
				ox.Spec(`<context> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca`, `the path to the CA certificate`).
					String(`crt`, `the path to the certificate`).
					String(`key`, `the path to the key`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config context
				ox.Usage(`context`, `Set the current context`),
				ox.Banner(`Set the current context`),
				ox.Spec(`<context> [flags]`),
				ox.Aliases("use-context"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config contexts
				ox.Usage(`contexts`, `List defined contexts`),
				ox.Banner(`List defined contexts`),
				ox.Spec(`[flags]`),
				ox.Aliases("get-contexts"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config endpoint
				ox.Usage(`endpoint`, `Set the endpoint(s) for the current context`),
				ox.Banner(`Set the endpoint(s) for the current context`),
				ox.Spec(`<endpoint>... [flags]`),
				ox.Aliases("endpoints"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config info
				ox.Usage(`info`, `Show information about the current context`),
				ox.Banner(`Show information about the current context`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`output`, `output format (json|yaml|text). Default text.`, ox.Default("text"), ox.Short("o")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config merge
				ox.Usage(`merge`, `Merge additional contexts from another client configuration file`),
				ox.Banner(`Contexts with the same name are renamed while merging configs.`),
				ox.Spec(`<from> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config new
				ox.Usage(`new`, `Generate a new client configuration file`),
				ox.Banner(`Generate a new client configuration file`),
				ox.Spec(`[<path>] [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Duration(`crt-ttl`, `certificate TTL`, ox.Default("8760h0m0s")).
					Slice(`roles`, `roles`, ox.Default("[os:admin]")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config node
				ox.Usage(`node`, `Set the node(s) for the current context`),
				ox.Banner(`Set the node(s) for the current context`),
				ox.Spec(`<endpoint>... [flags]`),
				ox.Aliases("nodes"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl config remove
				ox.Usage(`remove`, `Remove contexts`),
				ox.Banner(`Remove contexts`),
				ox.Spec(`<context> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`dry-run`, `dry run`).
					Bool(`noconfirm`, `do not ask for confirmation`, ox.Short("y")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl conformance
			ox.Usage(`conformance`, `Run conformance tests`),
			ox.Banner(`Run conformance tests`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl conformance [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl conformance kubernetes
				ox.Usage(`kubernetes`, `Run Kubernetes conformance tests`),
				ox.Banner(`Run Kubernetes conformance tests`),
				ox.Spec(`[flags]`),
				ox.Aliases("k8s"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`mode`, `conformance test mode: [fast, certified]`, ox.Default("fast")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl containers
			ox.Usage(`containers`, `List containers`),
			ox.Banner(`List containers`),
			ox.Spec(`[flags]`),
			ox.Aliases("c"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`kubernetes`, `use the k8s.io containerd namespace`, ox.Short("k")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl copy
			ox.Usage(`copy`, `Copy data out from the node`),
			ox.Banner(`Creates an .tar.gz archive at the node starting at <src-path> and
streams it back to the client.

If '-' is given for <local-path>, archive is written to stdout.
Otherwise archive is extracted to <local-path> which should be an empty directory or
talosctl creates a directory if <local-path> doesn't exist. Command doesn't preserve
ownership and access mode for the files in extract mode, while  streamed .tar archive
captures ownership and permission bits.`),
			ox.Spec(`<src-path> -|<local-path> [flags]`),
			ox.Aliases("cp"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl dashboard
			ox.Usage(`dashboard`, `Cluster dashboard with node overview, logs and real-time metrics`),
			ox.Banner(`Provide a text-based UI to navigate node overview, logs and real-time metrics.

Keyboard shortcuts:

 - h, &lt;Left&gt; - switch one node to the left
 - l, &lt;Right&gt; - switch one node to the right
 - j, &lt;Down&gt; - scroll logs/process list down
 - k, &lt;Up&gt; - scroll logs/process list up
 - &lt;C-d&gt; - scroll logs/process list half page down
 - &lt;C-u&gt; - scroll logs/process list half page up
 - &lt;C-f&gt; - scroll logs/process list one page down
 - &lt;C-b&gt; - scroll logs/process list one page up`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Duration(`update-interval`, `interval between updates`, ox.Default("3s"), ox.Short("d")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl dmesg
			ox.Usage(`dmesg`, `Retrieve kernel logs`),
			ox.Banner(`Retrieve kernel logs`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`follow`, `specify if the kernel log should be streamed`, ox.Short("f")).
				Bool(`tail`, `specify if only new messages should be sent (makes sense only when combined with --follow)`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl edit
			ox.Usage(`edit`, `Edit a resource from the default editor.`),
			ox.Banner(`The edit command allows you to directly edit any API resource
you can retrieve via the command line tools.

It will open the editor defined by your TALOS_EDITOR,
or EDITOR environment variables, or fall back to 'vi' for Linux
or 'notepad' for Windows.`),
			ox.Spec(`<type> [<id>] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`dry-run`, `do not apply the change after editing and print the change summary instead`).
				String(`mode`, `no-reboot, reboot, staged, try   apply config mode`, ox.Spec(`auto,`), ox.Default("auto"), ox.Short("m")).
				String(`namespace`, `resource namespace`, ox.Default("is to use default namespace per resource")).
				Duration(`timeout`, `the config will be rolled back after specified timeout (if try mode is selected)`, ox.Default("1m0s")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl etcd
			ox.Usage(`etcd`, `Manage etcd`),
			ox.Banner(`Manage etcd`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl etcd [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl etcd alarm
				ox.Usage(`alarm`, `Manage etcd alarms`),
				ox.Banner(`Manage etcd alarms`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "talosctl etcd alarm [command] --help" for more information about a command.`),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				ox.Sub( // talosctl etcd alarm disarm
					ox.Usage(`disarm`, `Disarm the etcd alarms for the node.`),
					ox.Banner(`Disarm the etcd alarms for the node.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
						String(`context`, `Context to be used in command`, ox.Section(0)).
						Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
						Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
						String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				),
				ox.Sub( // talosctl etcd alarm list
					ox.Usage(`list`, `List the etcd alarms for the node.`),
					ox.Banner(`List the etcd alarms for the node.`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
						String(`context`, `Context to be used in command`, ox.Section(0)).
						Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
						Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
						String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				),
			),
			ox.Sub( // talosctl etcd defrag
				ox.Usage(`defrag`, `Defragment etcd database on the node`),
				ox.Banner(`Defragmentation is a maintenance operation that releases unused space from the etcd database file.
Defragmentation is a resource heavy operation and should be performed only when necessary on a single node at a time.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd forfeit-leadership
				ox.Usage(`forfeit-leadership`, `Tell node to forfeit etcd cluster leadership`),
				ox.Banner(`Tell node to forfeit etcd cluster leadership`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd leave
				ox.Usage(`leave`, `Tell nodes to leave etcd cluster`),
				ox.Banner(`Tell nodes to leave etcd cluster`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd members
				ox.Usage(`members`, `Get the list of etcd cluster members`),
				ox.Banner(`Get the list of etcd cluster members`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd remove-member
				ox.Usage(`remove-member`, `Remove the node from etcd cluster`),
				ox.Banner(`Use this command only if you want to remove a member which is in broken state.
If there is no access to the node, or the node can't access etcd to call etcd leave.
Always prefer etcd leave over this command.`),
				ox.Spec(`<member ID> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd snapshot
				ox.Usage(`snapshot`, `Stream snapshot of the etcd node to the path.`),
				ox.Banner(`Stream snapshot of the etcd node to the path.`),
				ox.Spec(`<path> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl etcd status
				ox.Usage(`status`, `Get the status of etcd cluster member`),
				ox.Banner(`Returns the status of etcd member on the node, use multiple nodes to get status of all members.`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl events
			ox.Usage(`events`, `Stream runtime events`),
			ox.Banner(`Stream runtime events`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`actor-id`, `filter events by the specified actor ID`, ox.Default("is no filter")).
				Duration(`duration`, `show events for the past duration interval (one second resolution, default is to show no history)`).
				String(`since`, `show events after the specified event ID`, ox.Default("is to show no history")).
				Int32(`tail`, `show specified number of past events (use -1 to show full history, default is to show no history)`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl gen
			ox.Usage(`gen`, `Generate CAs, certificates, and private keys`),
			ox.Banner(`Generate CAs, certificates, and private keys`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl gen [command] --help" for more information about a command.`),
			ox.Flags().
				Bool(`force`, `will overwrite existing files`, ox.Short("f")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl gen ca
				ox.Usage(`ca`, `Generates a self-signed X.509 certificate authority`),
				ox.Banner(`Generates a self-signed X.509 certificate authority`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Int(`hours`, `the hours from now on which the certificate validity period ends`, ox.Default("87600")).
					String(`organization`, `X.509 distinguished name for the Organization`).
					Bool(`rsa`, `generate in RSA format`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen config
				ox.Usage(`config`, `Generates a set of configuration files for Talos cluster`),
				ox.Banner(`The cluster endpoint is the URL for the Kubernetes API. If you decide to use
a control plane node, common in a single node control plane setup, use port 6443 as
this is the port that the API server binds to on every control plane node. For an HA
setup, usually involving a load balancer, use the IP and port of the load balancer.`),
				ox.Spec(`<cluster name> <cluster endpoint> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Slice(`additional-sans`, `additional Subject-Alt-Names for the APIServer certificate`).
					Array(`config-patch`, `patch generated machineconfigs (applied to all node types), use @file to read a patch from file`).
					Array(`config-patch-control-plane`, `patch generated machineconfigs (applied to 'init' and 'controlplane' types)`).
					Array(`config-patch-worker`, `patch generated machineconfigs (applied to 'worker' type)`).
					String(`dns-domain`, `the dns domain to use for cluster`, ox.Default("cluster.local")).
					String(`install-disk`, `the disk to install to`, ox.Default("/dev/sda")).
					String(`install-image`, `the image used to perform an installation`, ox.Default("ghcr.io/siderolabs/installer:v1.10.4")).
					String(`kubernetes-version`, `desired kubernetes version to run`, ox.Default("1.33.1")).
					String(`output`, `destination to output generated files. when multiple output types are specified, it must be a directory. for a single output type, it must either be a file path, or "-" for stdout`, ox.Short("o")).
					Slice(`output-types`, `types of outputs to be generated. valid types are: ["controlplane" "worker" "talosconfig"]`, ox.Default("[controlplane,worker,talosconfig]"), ox.Short("t")).
					Bool(`persist`, `the desired persist value for configs`, ox.Default("true"), ox.Short("p")).
					Slice(`registry-mirror`, `list of registry mirrors to use in format: <registry host>=<mirror URL>`).
					String(`talos-version`, `the desired Talos version to generate config for (backwards compatibility, e.g. v0.8)`).
					Bool(`with-cluster-discovery`, `enable cluster discovery feature`, ox.Default("true")).
					Bool(`with-docs`, `renders all machine configs adding the documentation for each field`, ox.Default("true")).
					Bool(`with-examples`, `renders all machine configs with the commented examples`, ox.Default("true")).
					Bool(`with-kubespan`, `enable KubeSpan feature`).
					String(`with-secrets`, `use a secrets file generated using 'gen secrets'`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen crt
				ox.Usage(`crt`, `Generates an X.509 Ed25519 certificate`),
				ox.Banner(`Generates an X.509 Ed25519 certificate`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ca`, `path to the PEM encoded CERTIFICATE`).
					String(`csr`, `path to the PEM encoded CERTIFICATE REQUEST`).
					Int(`hours`, `the hours from now on which the certificate validity period ends`, ox.Default("24")).
					String(`name`, `the basename of the generated file`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen csr
				ox.Usage(`csr`, `Generates a CSR using an Ed25519 private key`),
				ox.Banner(`Generates a CSR using an Ed25519 private key`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ip`, `generate the certificate for this IP address`).
					String(`key`, `path to the PEM encoded EC or RSA PRIVATE KEY`).
					Slice(`roles`, `roles`, ox.Default("[os:admin]")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen key
				ox.Usage(`key`, `Generates an Ed25519 private key`),
				ox.Banner(`Generates an Ed25519 private key`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`name`, `the basename of the generated file`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen keypair
				ox.Usage(`keypair`, `Generates an X.509 Ed25519 key pair`),
				ox.Banner(`Generates an X.509 Ed25519 key pair`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`ip`, `generate the certificate for this IP address`).
					String(`organization`, `X.509 distinguished name for the Organization`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen secrets
				ox.Usage(`secrets`, `Generates a secrets bundle file which can later be used to generate a config`),
				ox.Banner(`Generates a secrets bundle file which can later be used to generate a config`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`from-controlplane-config`, `use the provided controlplane Talos machine configuration as input`).
					String(`from-kubernetes-pki`, `use a Kubernetes PKI directory (e.g. /etc/kubernetes/pki) as input`, ox.Short("p")).
					String(`kubernetes-bootstrap-token`, `use the provided bootstrap token as input`, ox.Short("t")).
					String(`output-file`, `path of the output file`, ox.Default("secrets.yaml"), ox.Short("o")).
					String(`talos-version`, `the desired Talos version to generate secrets bundle for (backwards compatibility, e.g. v0.8)`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl gen secureboot
				ox.Usage(`secureboot`, `Generates secrets for the SecureBoot process`),
				ox.Banner(`Generates secrets for the SecureBoot process`),
				ox.Spec(`[command]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Footer(`Use "talosctl gen secureboot [command] --help" for more information about a command.`),
				ox.Flags().
					String(`output`, `path to the directory storing the generated files`, ox.Default("_out"), ox.Short("o")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				ox.Sub( // talosctl gen secureboot database
					ox.Usage(`database`, `Generates a UEFI database to enroll the signing certificate`),
					ox.Banner(`Generates a UEFI database to enroll the signing certificate`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`enrolled-certificate`, `path to the certificate to enroll`, ox.Default("_out/uki-signing-cert.pem")).
						Bool(`include-well-known-uefi-certs`, `include well-known UEFI (Microsoft) certificates in the database`).
						String(`signing-certificate`, `path to the certificate used to sign the database`, ox.Default("_out/uki-signing-cert.pem")).
						String(`signing-key`, `path to the key used to sign the database`, ox.Default("_out/uki-signing-key.pem")).
						String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
						String(`context`, `Context to be used in command`, ox.Section(0)).
						Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
						Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
						Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
						String(`output`, `path to the directory storing the generated files`, ox.Default("_out"), ox.Short("o"), ox.Section(0)).
						String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				),
				ox.Sub( // talosctl gen secureboot pcr
					ox.Usage(`pcr`, `Generates a key which is used to sign TPM PCR values`),
					ox.Banner(`Generates a key which is used to sign TPM PCR values`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
						String(`context`, `Context to be used in command`, ox.Section(0)).
						Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
						Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
						Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
						String(`output`, `path to the directory storing the generated files`, ox.Default("_out"), ox.Short("o"), ox.Section(0)).
						String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				),
				ox.Sub( // talosctl gen secureboot uki
					ox.Usage(`uki`, `Generates a certificate which is used to sign boot assets (UKI)`),
					ox.Banner(`Generates a certificate which is used to sign boot assets (UKI)`),
					ox.Spec(`[flags]`),
					ox.Help(ox.Sections(
						"Global Flags",
					)),
					ox.Flags().
						String(`common-name`, `common name for the certificate`, ox.Default("Test UKI Signing Key")).
						String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
						String(`context`, `Context to be used in command`, ox.Section(0)).
						Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
						Bool(`force`, `will overwrite existing files`, ox.Short("f"), ox.Section(0)).
						Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
						String(`output`, `path to the directory storing the generated files`, ox.Default("_out"), ox.Short("o"), ox.Section(0)).
						String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
				),
			),
		),
		ox.Sub( // talosctl get
			ox.Usage(`get`, `Get a specific resource or list of resources (use 'talosctl get rd' to see all available resource types).`),
			ox.Banner(`Similar to 'kubectl get', 'talosctl get' returns a set of resources from the OS.
To get a list of all available resource definitions, issue 'talosctl get rd'`),
			ox.Spec(`<type> [<id>] [flags]`),
			ox.Aliases("g"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`insecure`, `get resources using the insecure (encrypted with no auth) maintenance service`, ox.Short("i")).
				String(`namespace`, `resource namespace`, ox.Default("is to use default namespace per resource")).
				String(`output`, `output mode (json, table, yaml, jsonpath)`, ox.Default("table"), ox.Short("o")).
				Bool(`watch`, `watch resource changes`, ox.Short("w")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl health
			ox.Usage(`health`, `Check cluster health`),
			ox.Banner(`Check cluster health`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Slice(`control-plane-nodes`, `specify IPs of control plane nodes`).
				String(`init-node`, `specify IPs of init node`).
				String(`k8s-endpoint`, `use endpoint instead of kubeconfig default`).
				Bool(`run-e2e`, `run Kubernetes e2e test`).
				Bool(`server`, `run server-side check`, ox.Default("true")).
				Duration(`wait-timeout`, `timeout to wait for the cluster to be ready`, ox.Default("20m0s")).
				Slice(`worker-nodes`, `specify IPs of worker nodes`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl image
			ox.Usage(`image`, `Manage CRI container images`),
			ox.Banner(`Manage CRI container images`),
			ox.Spec(`[command]`),
			ox.Aliases("images"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl image [command] --help" for more information about a command.`),
			ox.Flags().
				String(`namespace`, `namespace to use: system (etcd and kubelet images) or `+"`"+`cri`+"`"+` for all Kubernetes workloads`, ox.Spec(`system`), ox.Default("cri")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl image cache-create
				ox.Usage(`cache-create`, `Create a cache of images in OCI format into a directory`),
				ox.Banner(`Create a cache of images in OCI format into a directory`),
				ox.Spec(`[flags]`),
				ox.Example(`
talosctl images cache-create --images=ghcr.io/siderolabs/kubelet:v1.33.1 --image-cache-path=/tmp/talos-image-cache

Alternatively, stdin can be piped to the command:
talosctl images default | talosctl images cache-create --image-cache-path=/tmp/talos-image-cache --images=-`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`force`, `force overwrite of existing image cache`).
					String(`image-cache-path`, `directory to save the image cache in OCI format`).
					String(`image-layer-cache-path`, `directory to save the image layer cache`).
					Slice(`images`, `images to cache`).
					Bool(`insecure`, `allow insecure registries`).
					String(`platform`, `platform to use for the cache`, ox.Default("linux/amd64")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`namespace`, `namespace to use: system (etcd and kubelet images) or `+"`"+`cri`+"`"+` for all Kubernetes workloads`, ox.Spec(`system`), ox.Default("cri"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl image default
				ox.Usage(`default`, `List the default images used by Talos`),
				ox.Banner(`List the default images used by Talos`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`namespace`, `namespace to use: system (etcd and kubelet images) or `+"`"+`cri`+"`"+` for all Kubernetes workloads`, ox.Spec(`system`), ox.Default("cri"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl image list
				ox.Usage(`list`, `List CRI images`),
				ox.Banner(`List CRI images`),
				ox.Spec(`[flags]`),
				ox.Aliases("l", "ls"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`namespace`, `namespace to use: system (etcd and kubelet images) or `+"`"+`cri`+"`"+` for all Kubernetes workloads`, ox.Spec(`system`), ox.Default("cri"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl image pull
				ox.Usage(`pull`, `Pull an image into CRI`),
				ox.Banner(`Pull an image into CRI`),
				ox.Spec(`<image> [flags]`),
				ox.Aliases("p"),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					String(`namespace`, `namespace to use: system (etcd and kubelet images) or `+"`"+`cri`+"`"+` for all Kubernetes workloads`, ox.Spec(`system`), ox.Default("cri"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl inject
			ox.Usage(`inject`, `Inject Talos API resources into Kubernetes manifests`),
			ox.Banner(`Inject Talos API resources into Kubernetes manifests`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl inject [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl inject serviceaccount
				ox.Usage(`serviceaccount`, `Inject Talos API ServiceAccount into Kubernetes manifests`),
				ox.Banner(`Inject Talos API ServiceAccount into Kubernetes manifests`),
				ox.Spec(`[--roles='<ROLE_1>,<ROLE_2>'] -f <manifest.yaml> [flags]`),
				ox.Aliases("tsa"),
				ox.Example(`
talosctl inject serviceaccount --roles="os:admin" -f deployment.yaml > deployment-injected.yaml

Alternatively, stdin can be piped to the command:
cat deployment.yaml | talosctl inject serviceaccount --roles="os:admin" -f - > deployment-injected.yaml`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`file`, `file with Kubernetes manifests to be injected with ServiceAccount`, ox.Short("f")).
					Slice(`roles`, `roles to add to the generated ServiceAccount manifests`, ox.Default("[os:reader]"), ox.Short("r")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl inspect
			ox.Usage(`inspect`, `Inspect internals of Talos`),
			ox.Banner(`Inspect internals of Talos`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl inspect [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl inspect dependencies
				ox.Usage(`dependencies`, `Inspect controller-resource dependencies as graphviz graph.`),
				ox.Banner(`Inspect controller-resource dependencies as graphviz graph.

Pipe the output of the command through the "dot" program (part of graphviz package)
to render the graph:

    talosctl inspect dependencies | dot -Tpng > graph.png`),
				ox.Spec(`[flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`with-resources`, `display live resource information with dependencies`).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl kubeconfig
			ox.Usage(`kubeconfig`, `Download the admin kubeconfig from the node`),
			ox.Banner(`Download the admin kubeconfig from the node.
If merge flag is true, config will be merged with ~/.kube/config or [local-path] if specified.
Otherwise, kubeconfig will be written to PWD or [local-path] if specified.

If merge flag is false and [local-path] is "-", config will be written to stdout.`),
			ox.Spec(`[local-path] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`force`, `Force overwrite of kubeconfig if already present, force overwrite on kubeconfig merge`, ox.Short("f")).
				String(`force-context-name`, `Force context name for kubeconfig merge`).
				Bool(`merge`, `Merge with existing kubeconfig`, ox.Default("true"), ox.Short("m")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl list
			ox.Usage(`list`, `Retrieve a directory listing`),
			ox.Banner(`Retrieve a directory listing`),
			ox.Spec(`[path] [flags]`),
			ox.Aliases("ls"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Int32(`depth`, `maximum recursion depth`, ox.Default("1"), ox.Short("d")).
				Bool(`humanize`, `humanize size and time in the output`, ox.Short("H")).
				Bool(`long`, `display additional file details`, ox.Short("l")).
				Bool(`recurse`, `recurse into subdirectories`, ox.Short("r")).
				Slice(`type`, `filter by specified types:`, ox.Short("t")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl logs
			ox.Usage(`logs`, `Retrieve logs for a service`),
			ox.Banner(`Retrieve logs for a service`),
			ox.Spec(`<service name> [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`follow`, `specify if the logs should be streamed`, ox.Short("f")).
				Bool(`kubernetes`, `use the k8s.io containerd namespace`, ox.Short("k")).
				Int32(`tail`, `lines of log file to display (default is to show from the beginning)`, ox.Default("-1")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl machineconfig
			ox.Usage(`machineconfig`, `Machine config related commands`),
			ox.Banner(`Machine config related commands`),
			ox.Spec(`[command]`),
			ox.Aliases("mc"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl machineconfig [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl machineconfig gen
				ox.Usage(`gen`, `Generates a set of configuration files for Talos cluster`),
				ox.Banner(`The cluster endpoint is the URL for the Kubernetes API. If you decide to use
a control plane node, common in a single node control plane setup, use port 6443 as
this is the port that the API server binds to on every control plane node. For an HA
setup, usually involving a load balancer, use the IP and port of the load balancer.`),
				ox.Spec(`<cluster name> <cluster endpoint> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl machineconfig patch
				ox.Usage(`patch`, `Patch a machine config`),
				ox.Banner(`Patch a machine config`),
				ox.Spec(`<machineconfig-file> [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`output`, `output destination. if not specified, output will be printed to stdout`, ox.Short("o")).
					Array(`patch`, `patch generated machineconfigs (applied to all node types), use @file to read a patch from file`, ox.Short("p")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl memory
			ox.Usage(`memory`, `Show memory usage`),
			ox.Banner(`Show memory usage`),
			ox.Spec(`[flags]`),
			ox.Aliases("m", "free"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`verbose`, `display extended memory statistics`, ox.Short("v")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl meta
			ox.Usage(`meta`, `Write and delete keys in the META partition`),
			ox.Banner(`Write and delete keys in the META partition`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl meta [command] --help" for more information about a command.`),
			ox.Flags().
				Bool(`insecure`, `write|delete meta using the insecure (encrypted with no auth) maintenance service`, ox.Short("i")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl meta delete
				ox.Usage(`delete`, `Delete a key from the META partition.`),
				ox.Banner(`Delete a key from the META partition.`),
				ox.Spec(`key [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`insecure`, `write|delete meta using the insecure (encrypted with no auth) maintenance service`, ox.Short("i"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
			ox.Sub( // talosctl meta write
				ox.Usage(`write`, `Write a key-value pair to the META partition.`),
				ox.Banner(`Write a key-value pair to the META partition.`),
				ox.Spec(`key value [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Bool(`insecure`, `write|delete meta using the insecure (encrypted with no auth) maintenance service`, ox.Short("i"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
		ox.Sub( // talosctl mounts
			ox.Usage(`mounts`, `List mounts`),
			ox.Banner(`List mounts`),
			ox.Spec(`[flags]`),
			ox.Aliases("mount"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl netstat
			ox.Usage(`netstat`, `Show network connections and sockets`),
			ox.Banner(`Show network connections and sockets.

You can pass an optional argument to view a specific pod's connections.
To do this, format the argument as "namespace/pod".
Note that only pods with a pod network namespace are allowed.
If you don't pass an argument, the command will show host connections.`),
			ox.Spec(`[flags]`),
			ox.Aliases("ss"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`all`, `display all sockets states (default: connected)`, ox.Short("a")).
				Bool(`extend`, `show detailed socket information`, ox.Short("x")).
				Bool(`ipv4`, `display only ipv4 sockets`, ox.Short("4")).
				Bool(`ipv6`, `display only ipv6 sockets`, ox.Short("6")).
				Bool(`listening`, `display listening server sockets`, ox.Short("l")).
				Bool(`pods`, `show sockets used by Kubernetes pods`, ox.Short("k")).
				Bool(`programs`, `show process using socket`, ox.Short("p")).
				Bool(`raw`, `display only RAW sockets`, ox.Short("w")).
				Bool(`tcp`, `display only TCP sockets`, ox.Short("t")).
				Bool(`timers`, `display timers`, ox.Short("o")).
				Bool(`udp`, `display only UDP sockets`, ox.Short("u")).
				Bool(`udplite`, `display only UDPLite sockets`, ox.Short("U")).
				Bool(`verbose`, `display sockets of all supported transport protocols`, ox.Short("v")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl patch
			ox.Usage(`patch`, `Update field(s) of a resource using a JSON patch.`),
			ox.Banner(`Update field(s) of a resource using a JSON patch.`),
			ox.Spec(`<type> [<id>] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`dry-run`, `print the change summary and patch preview without applying the changes`).
				String(`mode`, `no-reboot, reboot, staged, try   apply config mode`, ox.Spec(`auto,`), ox.Default("auto"), ox.Short("m")).
				String(`namespace`, `resource namespace`, ox.Default("is to use default namespace per resource")).
				Array(`patch`, `the patch to be applied to the resource file, use @file to read a patch from file.`, ox.Short("p")).
				String(`patch-file`, `a file containing a patch to be applied to the resource.`).
				Duration(`timeout`, `the config will be rolled back after specified timeout (if try mode is selected)`, ox.Default("1m0s")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl pcap
			ox.Usage(`pcap`, `Capture the network packets from the node.`),
			ox.Banner(`The command launches packet capture on the node and streams back the packets as raw pcap file.

Default behavior is to decode the packets with internal decoder to stdout:

    talosctl pcap -i eth0

Raw pcap file can be saved with `+"`"+`--output`+"`"+` flag:

    talosctl pcap -i eth0 --output eth0.pcap

Output can be piped to tcpdump:

    talosctl pcap -i eth0 -o - | tcpdump -vvv -r -

BPF filter can be applied, but it has to compiled to BPF instructions first using tcpdump.
Correct link type should be specified for the tcpdump: EN10MB for Ethernet links and RAW
for e.g. Wireguard tunnels:

    talosctl pcap -i eth0 --bpf-filter "$(tcpdump -dd -y EN10MB 'tcp and dst port 80')"

    talosctl pcap -i kubespan --bpf-filter "$(tcpdump -dd -y RAW 'port 50000')"

As packet capture is transmitted over the network, it is recommended to filter out the Talos API traffic,
e.g. by excluding packets with the port 50000.`),
			ox.Spec(`[flags]`),
			ox.Aliases("tcpdump"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`bpf-filter`, `bpf filter to apply, tcpdump -dd format`).
				Duration(`duration`, `duration of the capture`).
				String(`interface`, `interface name to capture packets on`, ox.Default("eth0"), ox.Short("i")).
				String(`output`, `if not set, decode packets to stdout; if set write raw pcap data to a file, use '-' for stdout`, ox.Short("o")).
				Bool(`promiscuous`, `put interface into promiscuous mode`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl processes
			ox.Usage(`processes`, `List running processes`),
			ox.Banner(`List running processes`),
			ox.Spec(`[flags]`),
			ox.Aliases("p", "ps"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`sort`, `Column to sort output by. [rss|cpu]`, ox.Default("rss"), ox.Short("s")).
				Bool(`watch`, `Stream running processes`, ox.Short("w")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl read
			ox.Usage(`read`, `Read a file on the machine`),
			ox.Banner(`Read a file on the machine`),
			ox.Spec(`<path> [flags]`),
			ox.Aliases("cat"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl reboot
			ox.Usage(`reboot`, `Reboot a node`),
			ox.Banner(`Reboot a node`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`debug`, `debug operation from kernel logs. --wait is set to true when this flag is set`).
				String(`mode`, `select the reboot mode: "default", "powercycle" (skips kexec)`, ox.Default("default"), ox.Short("m")).
				Duration(`timeout`, `time to wait for the operation is complete if --debug or --wait is set`, ox.Default("30m0s")).
				Bool(`wait`, `wait for the operation to complete, tracking its progress. always set to true when --debug is set`, ox.Default("true")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl reset
			ox.Usage(`reset`, `Reset a node`),
			ox.Banner(`Reset a node`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`debug`, `debug operation from kernel logs. --wait is set to true when this flag is set`).
				Bool(`graceful`, `if true, attempt to cordon/drain node and leave etcd (if applicable)`, ox.Default("true")).
				Bool(`insecure`, `reset using the insecure (encrypted with no auth) maintenance service`).
				Bool(`reboot`, `if true, reboot the node after resetting instead of shutting down`).
				Slice(`system-labels-to-wipe`, `if set, just wipe selected system disk partitions by label but keep other partitions intact`).
				Duration(`timeout`, `time to wait for the operation is complete if --debug or --wait is set`, ox.Default("30m0s")).
				Slice(`user-disks-to-wipe`, `if set, wipes defined devices in the list`).
				Bool(`wait`, `wait for the operation to complete, tracking its progress. always set to true when --debug is set`, ox.Default("true")).
				String(`wipe-mode`, `system-disk, user-disks   disk reset mode`, ox.Spec(`all,`), ox.Default("all")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl restart
			ox.Usage(`restart`, `Restart a process`),
			ox.Banner(`Restart a process`),
			ox.Spec(`<id> [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`kubernetes`, `use the k8s.io containerd namespace`, ox.Short("k")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl rollback
			ox.Usage(`rollback`, `Rollback a node to the previous installation`),
			ox.Banner(`Rollback a node to the previous installation`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl rotate-ca
			ox.Usage(`rotate-ca`, `Rotate cluster CAs (Talos and Kubernetes APIs).`),
			ox.Banner(`The command can rotate both Talos and Kubernetes root CAs (for the API).
By default both CAs are rotated, but you can choose to rotate just one or another.
The command starts by generating new CAs, and gracefully applying it to the cluster.

For Kubernetes, the command only rotates the API server issuing CA, and other Kubernetes
PKI can be rotated by applying machine config changes to the controlplane nodes.`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Slice(`control-plane-nodes`, `specify IPs of control plane nodes`).
				Bool(`dry-run`, `dry-run mode (no changes to the cluster)`, ox.Default("true")).
				String(`init-node`, `specify IPs of init node`).
				String(`k8s-endpoint`, `use endpoint instead of kubeconfig default`).
				Bool(`kubernetes`, `rotate Kubernetes API CA`, ox.Default("true")).
				String(`output`, `path to the output new talosconfig`, ox.Spec(`talosconfig`), ox.Default("talosconfig"), ox.Short("o")).
				Bool(`talos`, `rotate Talos API CA`, ox.Default("true")).
				Bool(`with-docs`, `patch all machine configs adding the documentation for each field`, ox.Default("true")).
				Bool(`with-examples`, `patch all machine configs with the commented examples`, ox.Default("true")).
				Slice(`worker-nodes`, `specify IPs of worker nodes`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl service
			ox.Usage(`service`, `Retrieve the state of a service (or all services), control service state`),
			ox.Banner(`Service control command. If run without arguments, lists all the services and their state.
If service ID is specified, default action 'status' is executed which shows status of a single list service.
With actions 'start', 'stop', 'restart', service state is updated respectively.`),
			ox.Spec(`[<id> [start|stop|restart|status]] [flags]`),
			ox.Aliases("services"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl shutdown
			ox.Usage(`shutdown`, `Shutdown a node`),
			ox.Banner(`Shutdown a node`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`debug`, `debug operation from kernel logs. --wait is set to true when this flag is set`).
				Bool(`force`, `if true, force a node to shutdown without a cordon/drain`).
				Duration(`timeout`, `time to wait for the operation is complete if --debug or --wait is set`, ox.Default("30m0s")).
				Bool(`wait`, `wait for the operation to complete, tracking its progress. always set to true when --debug is set`, ox.Default("true")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl stats
			ox.Usage(`stats`, `Get container stats`),
			ox.Banner(`Get container stats`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`kubernetes`, `use the k8s.io containerd namespace`, ox.Short("k")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl support
			ox.Usage(`support`, `Dump debug information about the cluster`),
			ox.Banner(`Generated bundle contains the following debug information:

- For each node:

	- Kernel logs.
	- All Talos internal services logs.
	- All kube-system pods logs.
	- Talos COSI resources without secrets.
	- COSI runtime state graph.
	- Processes snapshot.
	- IO pressure snapshot.
	- Mounts list.
	- PCI devices info.
	- Talos version.

- For the cluster:

	- Kubernetes nodes and kube-system pods manifests.`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Int(`num-workers`, `number of workers per node`, ox.Default("1"), ox.Short("w")).
				String(`output`, `output file to write support archive to`, ox.Short("O")).
				Bool(`verbose`, `verbose output`, ox.Short("v")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl time
			ox.Usage(`time`, `Gets current server time`),
			ox.Banner(`Gets current server time`),
			ox.Spec(`[--check server] [flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`check`, `checks server time against specified ntp server`, ox.Short("c")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl upgrade
			ox.Usage(`upgrade`, `Upgrade Talos on the target node`),
			ox.Banner(`Upgrade Talos on the target node`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`debug`, `debug operation from kernel logs. --wait is set to true when this flag is set`).
				Bool(`force`, `force the upgrade (skip checks on etcd health and members, might lead to data loss)`, ox.Short("f")).
				String(`image`, `the container image to use for performing the install`, ox.Default("ghcr.io/siderolabs/installer:v1.10.4"), ox.Short("i")).
				Bool(`insecure`, `upgrade using the insecure (encrypted with no auth) maintenance service`).
				String(`reboot-mode`, `select the reboot mode during upgrade. Mode "powercycle" bypasses kexec. Valid values are: ["default" "powercycle"].`, ox.Default("default"), ox.Short("m")).
				Bool(`stage`, `stage the upgrade to perform it after a reboot`, ox.Short("s")).
				Duration(`timeout`, `time to wait for the operation is complete if --debug or --wait is set`, ox.Default("30m0s")).
				Bool(`wait`, `wait for the operation to complete, tracking its progress. always set to true when --debug is set`, ox.Default("true")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl upgrade-k8s
			ox.Usage(`upgrade-k8s`, `Upgrade Kubernetes control plane in the Talos cluster.`),
			ox.Banner(`Command runs upgrade of Kubernetes control plane components between specified versions.`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`apiserver-image`, `kube-apiserver image to use`, ox.Default("registry.k8s.io/kube-apiserver")).
				String(`controller-manager-image`, `kube-controller-manager image to use`, ox.Default("registry.k8s.io/kube-controller-manager")).
				Bool(`dry-run`, `skip the actual upgrade and show the upgrade plan instead`).
				String(`endpoint`, `the cluster control plane endpoint`).
				String(`from`, `the Kubernetes control plane version to upgrade from`).
				String(`kubelet-image`, `kubelet image to use`, ox.Default("ghcr.io/siderolabs/kubelet")).
				Bool(`pre-pull-images`, `pre-pull images before upgrade`, ox.Default("true")).
				String(`proxy-image`, `kube-proxy image to use`, ox.Default("registry.k8s.io/kube-proxy")).
				String(`scheduler-image`, `kube-scheduler image to use`, ox.Default("registry.k8s.io/kube-scheduler")).
				String(`to`, `the Kubernetes control plane version to upgrade to`, ox.Default("1.33.1")).
				Bool(`upgrade-kubelet`, `upgrade kubelet service`, ox.Default("true")).
				Bool(`with-docs`, `patch all machine configs adding the documentation for each field`, ox.Default("true")).
				Bool(`with-examples`, `patch all machine configs with the commented examples`, ox.Default("true")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl usage
			ox.Usage(`usage`, `Retrieve a disk usage`),
			ox.Banner(`Retrieve a disk usage`),
			ox.Spec(`[path1] [path2] ... [pathN] [flags]`),
			ox.Aliases("du"),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				Bool(`all`, `write counts for all files, not just directories`, ox.Short("a")).
				Int32(`depth`, `maximum recursion depth`, ox.Short("d")).
				Bool(`humanize`, `humanize size and time in the output`, ox.Short("H")).
				Int(`threshold`, `threshold exclude entries smaller than SIZE if positive, or entries greater than SIZE if negative`, ox.Short("t")).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl validate
			ox.Usage(`validate`, `Validate config`),
			ox.Banner(`Validate config`),
			ox.Spec(`[flags]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Flags().
				String(`config`, `the path of the config file`, ox.Short("c")).
				String(`mode`, `the mode to validate the config for (valid values are metal, cloud, and container)`, ox.Short("m")).
				Bool(`strict`, `treat validation warnings as errors`).
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
		),
		ox.Sub( // talosctl wipe
			ox.Usage(`wipe`, `Wipe block device or volumes`),
			ox.Banner(`Wipe block device or volumes`),
			ox.Spec(`[command]`),
			ox.Help(ox.Sections(
				"Global Flags",
			)),
			ox.Footer(`Use "talosctl wipe [command] --help" for more information about a command.`),
			ox.Flags().
				String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
				String(`context`, `Context to be used in command`, ox.Section(0)).
				Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
				Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
				String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			ox.Sub( // talosctl wipe disk
				ox.Usage(`disk`, `Wipe a block device (disk or partition) which is not used as a volume`),
				ox.Banner(`Wipe a block device (disk or partition) which is not used as a volume.

Use device names as arguments, for example: vda or sda5.`),
				ox.Spec(`<device names>... [flags]`),
				ox.Help(ox.Sections(
					"Global Flags",
				)),
				ox.Flags().
					Bool(`drop-partition`, `drop partition after wipe (if applicable)`).
					String(`method`, `wipe method to use [FAST ZEROES]`, ox.Default("FAST")).
					String(`cluster`, `Cluster to connect to if a proxy endpoint is used.`, ox.Section(0)).
					String(`context`, `Context to be used in command`, ox.Section(0)).
					Slice(`endpoints`, `override default endpoints in Talos configuration`, ox.Short("e"), ox.Section(0)).
					Slice(`nodes`, `target the specified nodes`, ox.Short("n"), ox.Section(0)).
					String(`talosconfig`, `The path to the Talos configuration file. Defaults to 'TALOSCONFIG' env variable if set, otherwise '$HOME/.talos/config' and '/var/run/secrets/talos.dev/config' in order.`, ox.Section(0)),
			),
		),
	)
}
