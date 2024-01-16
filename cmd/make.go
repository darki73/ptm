package cmd

import (
	"fmt"
	config "github.com/darki73/ptm/pkg/configuration"
	"github.com/darki73/ptm/pkg/maker"
	"github.com/darki73/ptm/pkg/qemu"
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
	"github.com/darki73/ptm/pkg/utils"
	"github.com/spf13/cobra"
)

// makeCommand represents the make command.
var makeCommand = &cobra.Command{
	Use:   "make",
	Short: "Creates template for Proxmox VE",
	Long:  "Asks user for input relevant to the template creation and then creates it based on the input.",
	Run: func(cmd *cobra.Command, args []string) {
		ensureRoot()
		if !ensurePackageAvailable("proxmox-ve") {
			printAndErrorOut("this application is only supported on Proxmox VE")
		}

		initializeConfiguration()

		var qemuConfiguration *qemu.Qemu

		if getConfiguration().GetQemu().IsConfigured() {
			qemuConfigurationFromConfigurationFile, err := createQemuConfigurationFromConfigurationFile(getConfiguration())
			if err == nil {
				qemuConfiguration = qemuConfigurationFromConfigurationFile
			}
		}

		qemuConfigurationFromFlags, err := createQemuConfigurationFromFlags()
		if err == nil {
			qemuConfiguration = qemuConfigurationFromFlags
		}

		handler, err := maker.NewMaker(getConfiguration(), qemuConfiguration)

		if err != nil {
			printAndErrorOut(err.Error())
		}

		if err := handler.Run(); err != nil {
			printAndErrorOut(err.Error())
		}
	},
}

// createQemuConfigurationFromFlags creates a QEMU configuration from the flags.
func createQemuConfigurationFromFlags() (*qemu.Qemu, error) {
	qemuConfiguration := qemu.NewQemuConfiguration()
	qemuConfiguration.
		SetIdentifier(identifier).
		SetName(name).
		SetCores(cores)

	if memory != "" {
		memoryValue, err := utils.ConvertToMegabytes(memory)
		if err != nil {
			return nil, err
		}
		qemuConfiguration.SetMemory(int(memoryValue))
	}

	qemuConfiguration.SetCpuType(cpuType)
	qemuConfiguration.SetNetworkDriver(networkDriver)
	qemuConfiguration.SetNetworkBridge(networkBridge)
	qemuConfiguration.SetStorage(storage)
	qemuConfiguration.SetImage(image)
	qemuConfiguration.SetNewImageSizeAsString(imageNewSize)
	qemuConfiguration.SetConfigurationSource(qemu.ConfigurationSourceFlags)

	cloudInitConfiguration := ci.NewCloudInitConfiguration()
	cloudInitConfiguration.SetUsername(ciUsername)
	cloudInitConfiguration.SetPassword(ciPassword)
	cloudInitConfiguration.SetKeys(ciSSHKeys)
	cloudInitConfiguration.SetIPv4(ciIPv4Address)
	cloudInitConfiguration.SetIPv4Gateway(ciIPv4Gateway)
	cloudInitConfiguration.SetIPv6(ciIPv6Address)
	cloudInitConfiguration.SetIPv6Gateway(ciIPv6Gateway)
	if ciIPv4Auto {
		cloudInitConfiguration.AutoConfigureIPv4()
	}
	if ciIPv6Auto {
		cloudInitConfiguration.AutoConfigureIPv6()
	}
	cloudInitConfiguration.SetConfigurationSource(ci.ConfigurationSourceFlags)

	qemuConfiguration.SetCloudInit(cloudInitConfiguration)

	isValid, err := qemuConfiguration.IsConfigurationValid()
	if err != nil {
		return nil, err
	}

	if !isValid {
		return nil, fmt.Errorf("flags provided qemu configuration is not valid")
	}

	return qemuConfiguration, nil
}

// createQemuConfigurationFromConfigurationFile creates a QEMU configuration from the configuration file.
func createQemuConfigurationFromConfigurationFile(configuration *config.Configuration) (*qemu.Qemu, error) {
	qc := configuration.GetQemu()

	qemuConfiguration := qemu.NewQemuConfiguration()
	qemuConfiguration.SetIdentifier(qc.GetIdentifier())
	qemuConfiguration.SetName(qc.GetName())
	resources := qc.GetResources()
	qemuConfiguration.SetCores(resources.GetCores())
	memory, err := resources.GetMemory()
	if err != nil {
		return nil, err
	}
	qemuConfiguration.SetImage(qc.GetImage())
	qemuConfiguration.SetStorage(qc.GetStorage().GetStorage())
	qemuConfiguration.SetNewImageSizeAsString(qc.GetStorage().GetResize())
	qemuConfiguration.SetMemory(int(memory))
	qemuConfiguration.SetCpuType(resources.GetCpuType())
	qemuConfiguration.SetNetworkDriver(qc.GetNetwork().GetDriver())
	qemuConfiguration.SetNetworkBridge(qc.GetNetwork().GetBridge())
	qemuConfiguration.SetConfigurationSource(qemu.ConfigurationSourceConfigurationFile)

	cic := configuration.GetCloudInit()

	if cic.GetEnabled() {
		cloudInitConfiguration := ci.NewCloudInitConfiguration()
		cloudInitConfiguration.SetUsername(cic.GetUsername())
		cloudInitConfiguration.SetPassword(cic.GetPassword())
		cloudInitConfiguration.SetKeys(cic.GetKeys())

		cicNetwork := cic.GetNetwork()
		if cicNetwork != nil {
			ipv4 := cicNetwork.GetIPv4()
			if ipv4 != nil {
				if ipv4.GetAutoConfigure() {
					cloudInitConfiguration.AutoConfigureIPv4()
				} else {
					cloudInitConfiguration.SetIPv4(ipv4.GetAddress())
					cloudInitConfiguration.SetIPv4Gateway(ipv4.GetGateway())
				}
			}

			ipv6 := cicNetwork.GetIPv6()
			if ipv6 != nil {
				if ipv6.GetAutoConfigure() {
					cloudInitConfiguration.AutoConfigureIPv6()
				} else {
					cloudInitConfiguration.SetIPv6(ipv6.GetAddress())
					cloudInitConfiguration.SetIPv6Gateway(ipv6.GetGateway())
				}
			}
		}

		cloudInitConfiguration.SetConfigurationSource(ci.ConfigurationSourceConfigurationFile)
		qemuConfiguration.SetCloudInit(cloudInitConfiguration)
	}

	isValid, err := qemuConfiguration.IsConfigurationValid()
	if !isValid {
		if err != nil && err.Error() == "missing image for virtual machine" {
			return qemuConfiguration, nil
		}
		return nil, err
	}

	return qemuConfiguration, nil
}

var (
	// identifier is an integer that is used as an identifier for the virtual machine template.
	identifier int
	// name is a string that is used as a name for the virtual machine template.
	name string
	// cores is an integer that is used as the number of cores to allocate to the virtual machine template.
	cores int
	// cpuType is a string that is used as the type of CPU to allocate to the virtual machine template.
	cpuType string
	// memory is an integer that is used as the amount of memory to allocate to the virtual machine template.
	memory string
	// storage is a string that is used to define the storage used for the virtual machine template.
	storage string
	// image is a string that is used to define the path to the image used for the virtual machine template creation.
	image string
	// imageNewSize is a string that is used to define the new size of the image for the virtual machine template.
	imageNewSize string
	// networkDriver is a string that is used to define the network driver for the virtual machine template.
	networkDriver string
	// networkBridge is a string that is used to define the network bridge for the virtual machine template.
	networkBridge string
	// ciUsername is a string that is used to define the cloud-init user for the virtual machine template.
	ciUsername string
	// ciPassword is a string that is used to define the cloud-init password for the virtual machine template.
	ciPassword string
	// ciSSHKeys is a string that is used to define the cloud-init SSH keys for the virtual machine template.
	ciSSHKeys []string
	// ciIPv4Auto is a flag that indicates whether IPv4 autoconfiguration is enabled.
	ciIPv4Auto bool
	// ciIPv6Auto is a flag that indicates whether IPv6 autoconfiguration is enabled.
	ciIPv6Auto bool
	// ciIPv4Address is a string that contains the IPv4 address.
	ciIPv4Address string
	// ciIPv6Address is a string that contains the IPv6 address.
	ciIPv6Address string
	// ciIPv4Gateway is a string that contains the IPv4 gateway.
	ciIPv4Gateway string
	// ciIPv6Gateway is a string that contains the IPv6 gateway.
	ciIPv6Gateway string
)

// init initializes the make command.
func init() {
	rootCmd.AddCommand(makeCommand)

	makeCommand.Flags().IntVar(&identifier, "identifier", 0, "Identifier of template")
	makeCommand.Flags().StringVar(&name, "name", "", "Name of the template")
	makeCommand.Flags().IntVar(&cores, "cores", 0, "Number of cpu cores")
	makeCommand.Flags().StringVar(&cpuType, "cpu-type", "", "Desired cpu type (host / kvm64 / etc)")
	makeCommand.Flags().StringVar(&memory, "memory", "", "Amount of memory (example: 1024 / 1024M / 1G)")
	makeCommand.Flags().StringVar(&storage, "storage", "", "Disk storage (local-lvm / local / etc)")
	makeCommand.Flags().StringVar(&image, "image", "", "Path to the image (/etc/ptm/images/image.qcow2)")
	makeCommand.Flags().StringVar(&imageNewSize, "image-new-size", "", "Size to which the image should be resized (example: 4G)")
	makeCommand.Flags().StringVar(&networkDriver, "network-driver", "", "Network driver (virtio / e1000 / etc)")
	makeCommand.Flags().StringVar(&networkBridge, "network-bridge", "", "Network bridge (vmbr0 / vmbr1 / etc)")
	makeCommand.Flags().StringVar(&ciUsername, "ci-username", "", "Username for cloud-init")
	makeCommand.Flags().StringVar(&ciPassword, "ci-password", "", "Password for cloud-init")
	makeCommand.Flags().StringArrayVar(&ciSSHKeys, "ci-ssh-keys", []string{}, "Comma-separated list of SSH keys for cloud-init")
	makeCommand.Flags().BoolVar(&ciIPv4Auto, "ci-ipv4-auto", true, "Automatically configure IPv4 for cloud-init")
	makeCommand.Flags().BoolVar(&ciIPv6Auto, "ci-ipv6-auto", true, "Automatically configure IPv6 for cloud-init")
	makeCommand.Flags().StringVar(&ciIPv4Address, "ci-ipv4-address", "", "Manually set IPv4 address for cloud-init (example: 10.10.10.10/24)")
	makeCommand.Flags().StringVar(&ciIPv6Address, "ci-ipv6-address", "", "Manually set IPv6 address for cloud-init (example: 2001:db8::1/64)")
	makeCommand.Flags().StringVar(&ciIPv4Gateway, "ci-ipv4-gateway", "", "Manually set IPv4 gateway for cloud-init (example: 10.10.10.1)")
	makeCommand.Flags().StringVar(&ciIPv6Gateway, "ci-ipv6-gateway", "", "Manually set IPv6 gateway for cloud-init (example: 2001:db8::1)")
}
