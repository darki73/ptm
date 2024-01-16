package proxmox

import (
	"github.com/darki73/ptm/pkg/log"
	"github.com/darki73/ptm/pkg/utils"
	"strconv"
	"strings"
)

// StorageTarget represents a storage target.
type StorageTarget struct {
	// name is the name of the storage target.
	name string
	// type is the type of the storage target.
	storageType string
	// status is the status of the storage target.
	status string
	// total is the total size of the storage target (in kilobytes).
	total int64
	// used is the used size of the storage target (in kilobytes).
	used int64
	// available is the available size of the storage target (in kilobytes).
	available int64
	// percentUsed is the percentage of used space on the storage target.
	percentUsed string
}

// NewStorageTarget creates a new storage target.
func NewStorageTarget(rawStorageTargetLine string) *StorageTarget {
	storageTarget := &StorageTarget{}
	storageTarget.parseRawStorageTargetLine(rawStorageTargetLine)
	return storageTarget
}

// GetName returns the name of the storage target.
func (storageTarget *StorageTarget) GetName() string {
	return storageTarget.name
}

// GetType returns the type of the storage target.
func (storageTarget *StorageTarget) GetType() string {
	return storageTarget.storageType
}

// GetStatus returns the status of the storage target.
func (storageTarget *StorageTarget) GetStatus() string {
	return storageTarget.status
}

// IsActive returns true if the storage target is active.
func (storageTarget *StorageTarget) IsActive() bool {
	return storageTarget.status == "active"
}

// GetTotal returns the total size of the storage target (in kilobytes).
func (storageTarget *StorageTarget) GetTotal() int64 {
	return storageTarget.total
}

// GetUsed returns the used size of the storage target (in kilobytes).
func (storageTarget *StorageTarget) GetUsed() int64 {
	return storageTarget.used
}

// GetAvailable returns the available size of the storage target (in kilobytes).
func (storageTarget *StorageTarget) GetAvailable() int64 {
	return storageTarget.available
}

// HasEnoughSpace returns true if the storage target has enough space.
func (storageTarget *StorageTarget) HasEnoughSpace(requiredSpace int64) bool {
	return storageTarget.available >= requiredSpace
}

// GetPercentUsed returns the percentage of used space on the storage target.
func (storageTarget *StorageTarget) GetPercentUsed() string {
	return storageTarget.percentUsed
}

// IsValidTarget returns true if the storage target is valid.
func (storageTarget *StorageTarget) IsValidTarget() bool {
	return storageTarget.IsActive() && storageTarget.GetType() != "dir"
}

// parseRawStorageTargetLine parses a raw storage target line.
func (storageTarget *StorageTarget) parseRawStorageTargetLine(rawStorageTargetLine string) {
	storageTargetLine := strings.Split(rawStorageTargetLine, " ")
	storageTargetLine = utils.RemoveEmptyStringsFromSlice(storageTargetLine)

	storageTarget.name = storageTargetLine[0]
	storageTarget.storageType = storageTargetLine[1]
	storageTarget.status = storageTargetLine[2]

	total, err := strconv.Atoi(storageTargetLine[3])
	if err != nil {
		log.Fatal(err.Error())
	}
	storageTarget.total = int64(total)

	used, err := strconv.Atoi(storageTargetLine[4])
	if err != nil {
		log.Fatal(err.Error())
	}
	storageTarget.used = int64(used)

	available, err := strconv.Atoi(storageTargetLine[5])
	if err != nil {
		log.Fatal(err.Error())
	}
	storageTarget.available = int64(available)

	storageTarget.percentUsed = storageTargetLine[6]
}
