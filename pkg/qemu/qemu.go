package qemu

import (
	"fmt"
	ci "github.com/darki73/ptm/pkg/qemu/cloud-init"
)

const (
	// ConfigurationSourcePrompt indicates that the configuration source is the prompt.
	ConfigurationSourcePrompt = "prompt"
	// ConfigurationSourceFlags indicates that the configuration source is the flags.
	ConfigurationSourceFlags = "flags"
	// ConfigurationSourceConfigurationFile indicates that the configuration source is the configuration file.
	ConfigurationSourceConfigurationFile = "configuration-file"
)

// Qemu is a structure that holds information for QEMU configurator.
type Qemu struct {
	// identifier is the QEMU VM identifier.
	identifier int
	// name is the QEMU VM name.
	name string
	// cores is the number of cores to use.
	cores int
	// memory is the amount of memory to use.
	memory int
	// cpuType is the CPU type to use.
	cpuType string
	// networkDriver is the network driver to use.
	networkDriver string
	// networkBridge is the network bridge to use.
	networkBridge string
	// storage is the storage to use.
	storage string
	// storageSize is the storage size (in MB).
	storageSize int64
	// image is the image to use.
	image string
	// imageSize is the original image size.
	imageSize int64
	// newImageSize is the image size to use (will resize if the image size is smaller than set size).
	newImageSize int64
	// newImageSizeAsString is the original image size as string.
	newImageSizeAsString string
	// cloudInit is the cloud-init configuration to use.
	cloudInit *ci.CloudInit
	// configurationSource is the configuration source.
	configurationSource string
}

// NewQemuConfiguration creates a new QEMU configuration.
func NewQemuConfiguration() *Qemu {
	return &Qemu{
		identifier:           0,
		name:                 "",
		cores:                0,
		memory:               0,
		cpuType:              "",
		networkDriver:        "",
		networkBridge:        "",
		storage:              "",
		storageSize:          0,
		image:                "",
		imageSize:            0,
		newImageSizeAsString: "",
		newImageSize:         0,
		cloudInit:            nil,
		configurationSource:  ConfigurationSourcePrompt,
	}
}

// GetIdentifier returns the QEMU VM identifier.
func (qemu *Qemu) GetIdentifier() int {
	return qemu.identifier
}

// SetIdentifier sets the QEMU VM identifier.
func (qemu *Qemu) SetIdentifier(identifier int) *Qemu {
	qemu.identifier = identifier
	return qemu
}

// GetName returns the QEMU VM name.
func (qemu *Qemu) GetName() string {
	return qemu.name
}

// SetName sets the QEMU VM name.
func (qemu *Qemu) SetName(name string) *Qemu {
	qemu.name = name
	return qemu
}

// GetCores returns the number of cores to use.
func (qemu *Qemu) GetCores() int {
	return qemu.cores
}

// SetCores sets the number of cores to use.
func (qemu *Qemu) SetCores(cores int) *Qemu {
	qemu.cores = cores
	return qemu
}

// GetMemory returns the amount of memory to use.
func (qemu *Qemu) GetMemory() int {
	return qemu.memory
}

// SetMemory sets the amount of memory to use.
func (qemu *Qemu) SetMemory(memory int) *Qemu {
	qemu.memory = memory
	return qemu
}

// GetCpuType returns the CPU type to use.
func (qemu *Qemu) GetCpuType() string {
	return qemu.cpuType
}

// GetNetworkDriver returns the network driver to use.
func (qemu *Qemu) GetNetworkDriver() string {
	return qemu.networkDriver
}

// SetNetworkDriver sets the network driver to use.
func (qemu *Qemu) SetNetworkDriver(networkDriver string) *Qemu {
	qemu.networkDriver = networkDriver
	return qemu
}

// GetNetworkBridge returns the network bridge to use.
func (qemu *Qemu) GetNetworkBridge() string {
	return qemu.networkBridge
}

// SetNetworkBridge sets the network bridge to use.
func (qemu *Qemu) SetNetworkBridge(networkBridge string) *Qemu {
	qemu.networkBridge = networkBridge
	return qemu
}

// SetCpuType sets the CPU type to use.
func (qemu *Qemu) SetCpuType(cpuType string) *Qemu {
	qemu.cpuType = cpuType
	return qemu
}

// GetStorage returns the storage to use.
func (qemu *Qemu) GetStorage() string {
	return qemu.storage
}

// SetStorage sets the storage to use.
func (qemu *Qemu) SetStorage(storage string) *Qemu {
	qemu.storage = storage
	return qemu
}

// GetStorageSize returns the storage size (in MB).
func (qemu *Qemu) GetStorageSize() int64 {
	return qemu.storageSize
}

// SetStorageSize sets the storage size (in MB).
func (qemu *Qemu) SetStorageSize(storageSize int64) *Qemu {
	qemu.storageSize = storageSize
	return qemu
}

// GetImage returns the image to use.
func (qemu *Qemu) GetImage() string {
	return qemu.image
}

// SetImage sets the image to use.
func (qemu *Qemu) SetImage(image string) *Qemu {
	qemu.image = image
	return qemu
}

// GetImageSize returns the original image size.
func (qemu *Qemu) GetImageSize() int64 {
	return qemu.imageSize
}

// SetImageSize sets the original image size.
func (qemu *Qemu) SetImageSize(imageSize int64) *Qemu {
	qemu.imageSize = imageSize
	return qemu
}

// GetNewImageSize returns the image size to use (will resize if the image size is smaller than set size).
func (qemu *Qemu) GetNewImageSize() int64 {
	return qemu.newImageSize
}

// SetNewImageSize sets the image size to use (will resize if the image size is smaller than set size).
func (qemu *Qemu) SetNewImageSize(newImageSize int64) *Qemu {
	qemu.newImageSize = newImageSize
	return qemu
}

// GetNewImageSizeAsString returns the original image size as string.
func (qemu *Qemu) GetNewImageSizeAsString() string {
	return qemu.newImageSizeAsString
}

// SetNewImageSizeAsString sets the original image size as string.
func (qemu *Qemu) SetNewImageSizeAsString(newImageSizeAsString string) *Qemu {
	qemu.newImageSizeAsString = newImageSizeAsString
	return qemu
}

// IsResizingRequired returns true if the image size is smaller than set size.
func (qemu *Qemu) IsResizingRequired() bool {
	return qemu.newImageSize != 0 && qemu.GetNewImageSize() > qemu.GetImageSize()
}

// GetCloudInit returns the cloud-init configuration to use.
func (qemu *Qemu) GetCloudInit() *ci.CloudInit {
	return qemu.cloudInit
}

// SetCloudInit sets the cloud-init configuration to use.
func (qemu *Qemu) SetCloudInit(cloudInit *ci.CloudInit) *Qemu {
	qemu.cloudInit = cloudInit
	return qemu
}

// GetConfigurationSource returns the configuration source.
func (qemu *Qemu) GetConfigurationSource() string {
	return qemu.configurationSource
}

// SetConfigurationSource sets the configuration source.
func (qemu *Qemu) SetConfigurationSource(configurationSource string) *Qemu {
	qemu.configurationSource = configurationSource
	return qemu
}

// IsConfigurationValid returns true if the configuration is valid.
func (qemu *Qemu) IsConfigurationValid() (bool, error) {
	if qemu.identifier == 0 {
		return false, fmt.Errorf("missing virtual machine identifier")
	}

	if qemu.name == "" {
		return false, fmt.Errorf("missing virtual machine name")
	}

	if qemu.cores == 0 {
		return false, fmt.Errorf("invalid number of cores assigned to virtual machine")
	}

	if qemu.memory == 0 {
		return false, fmt.Errorf("invalid amount of memory assigned to virtual machine")
	}

	if qemu.cpuType == "" {
		return false, fmt.Errorf("missing cpu type for virtual machine")
	}

	if qemu.networkDriver == "" {
		return false, fmt.Errorf("missing network driver for virtual machine")
	}

	if qemu.networkBridge == "" {
		return false, fmt.Errorf("missing network bridge for virtual machine")
	}

	if qemu.storage == "" {
		return false, fmt.Errorf("missing storage for virtual machine")
	}

	if qemu.image == "" {
		return false, fmt.Errorf("missing image for virtual machine")
	}

	if qemu.GetCloudInit() != nil {
		valid, err := qemu.GetCloudInit().IsConfigurationValid()
		if err != nil {
			return false, err
		}
		if !valid {
			return false, fmt.Errorf("invalid cloud-init configuration")
		}
	}

	return true, nil
}
