package proxmox

import (
	"os"
	"path"
	"strings"
)

// ShellKeys represents the container of available ssh keys.
type ShellKeys struct {
	// path is the path to ssh keys.
	path string
	// keys is the list of available ssh keys.
	keys []*ShellKey
}

// NewShellKeys creates a new ShellKeys instance.
func NewShellKeys() (*ShellKeys, error) {
	shellKeys := &ShellKeys{
		path: "/root/.ssh",
		keys: make([]*ShellKey, 0),
	}

	if err := shellKeys.listAvailableShellKeys(); err != nil {
		return nil, err
	}

	return shellKeys, nil
}

// GetKeys returns the list of available ssh keys.
func (shellKeys *ShellKeys) GetKeys() []*ShellKey {
	return shellKeys.keys
}

// FindKeyByName returns the ssh key with the specified name.
func (shellKeys *ShellKeys) FindKeyByName(name string) *ShellKey {
	for _, key := range shellKeys.keys {
		if key.GetName() == name {
			return key
		}
	}

	return nil
}

// listAvailableShellKeys lists the available shell keys.
func (shellKeys *ShellKeys) listAvailableShellKeys() error {
	items, err := os.ReadDir(shellKeys.path)
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}

		if !strings.Contains(item.Name(), ".pub") {
			continue
		}

		bytes, err := os.ReadFile(path.Join(shellKeys.path, item.Name()))
		if err != nil {
			return err
		}

		shellKeys.keys = append(shellKeys.keys, NewShellKey(item.Name(), shellKeys.path, string(bytes)))
	}

	return nil
}
