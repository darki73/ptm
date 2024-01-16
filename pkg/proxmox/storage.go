package proxmox

import (
	"github.com/darki73/ptm/pkg/utils"
	"strings"
)

// Storage represents a Proxmox VE storage.
type Storage struct {
	// command is the command to execute.
	command string
	// arguments are the arguments to pass to the command.
	arguments []string
	// targets are the list of available storage targets.
	targets []*StorageTarget
}

// NewStorage creates a new Storage instance.
func NewStorage() (*Storage, error) {
	storage := &Storage{
		command:   "pvesm",
		arguments: []string{"status"},
		targets:   make([]*StorageTarget, 0),
	}
	if err := storage.listAvailableStorageTargets(); err != nil {
		return nil, err
	}

	return storage, nil
}

// GetTargets returns the list of available storage targets.
func (storage *Storage) GetTargets() []*StorageTarget {
	return storage.targets
}

// FindTargetByName returns the storage target with the specified name.
func (storage *Storage) FindTargetByName(name string) *StorageTarget {
	for _, target := range storage.targets {
		if target.GetName() == name {
			return target
		}
	}

	return nil
}

// listAvailableStorageTargets lists the available storage targets.
func (storage *Storage) listAvailableStorageTargets() error {
	output, err := utils.ExecuteCommand("pvesm", "status")
	if err != nil {
		return err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if len(line) != 0 && !strings.Contains(line, "Status") {
			storageTarget := NewStorageTarget(line)
			if storageTarget.IsValidTarget() {
				storage.targets = append(storage.targets, storageTarget)
			}
		}
	}

	return nil
}
