package qemu

import (
	"fmt"
	"github.com/darki73/ptm/pkg/utils"
	"strconv"
	"strings"
)

// QemuResources is the QEMU resources configuration.
type QemuResources struct {
	// Cores is the number of cores to use.
	Cores int `json:"cores" yaml:"cores" toml:"cores" mapstructure:"cores"`
	// Memory is the amount of memory to use.
	Memory string `json:"memory" yaml:"memory" toml:"memory" mapstructure:"memory"`
	// CpuType is the CPU type to use.
	CpuType string `json:"cpu_type" yaml:"cpu_type" toml:"cpu_type" mapstructure:"cpu_type"`
}

// InitializeQemuResourcesWithDefaults initializes the QemuResources with default values.
func InitializeQemuResourcesWithDefaults() *QemuResources {
	return &QemuResources{
		Cores:   0,
		Memory:  "",
		CpuType: "host",
	}
}

// GetCores returns the number of cores to use.
func (qemuResources *QemuResources) GetCores() int {
	return qemuResources.Cores
}

// GetMemory returns the amount of memory to use.
func (qemuResources *QemuResources) GetMemory() (int64, error) {
	memory := strings.ToLower(qemuResources.Memory)

	if utils.IsNumeric(memory) {
		result, err := strconv.Atoi(memory)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory to integer: %s", err.Error())
		}
		return int64(result), nil
	}

	if strings.HasSuffix(memory, "t") {
		result, err := utils.ConvertToMegabytes(memory)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory to megabytes: %s", err.Error())
		}
		return result, nil
	}
	if strings.HasSuffix(memory, "g") {
		result, err := utils.ConvertToMegabytes(memory)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory to megabytes: %s", err.Error())
		}
		return result, nil
	}
	if strings.HasSuffix(memory, "m") {
		result, err := utils.ConvertToMegabytes(memory)
		if err != nil {
			return 0, fmt.Errorf("failed to convert memory to megabytes: %s", err.Error())
		}
		return result, nil
	}

	return 0, fmt.Errorf("failed to convert memory to megabytes: %s", "unknown suffix")
}

// GetCpuType returns the CPU type to use.
func (qemuResources *QemuResources) GetCpuType() string {
	return qemuResources.CpuType
}

// IsConfigured returns true if the configuration is configured.
func (qemuResources *QemuResources) IsConfigured() bool {
	if qemuResources.Cores == 0 {
		return false
	}

	if qemuResources.Memory == "" {
		return false
	}

	if qemuResources.CpuType == "" {
		return false
	}

	return true
}
