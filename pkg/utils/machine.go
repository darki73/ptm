package utils

import (
	"github.com/pbnjay/memory"
	"runtime"
)

// GetCoreCount returns the number of cores on the machine
func GetCoreCount() int {
	return runtime.NumCPU()
}

// GetTotalMemory returns the total memory on the machine
func GetTotalMemory() uint64 {
	return memory.TotalMemory() / 1024 / 1024
}
