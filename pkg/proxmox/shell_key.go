package proxmox

import "path"

// ShellKey represents the shell key.
type ShellKey struct {
	// name is the name of the key.
	name string
	// path is the path of the key.
	path string
	// content is the content of the key.
	content string
}

// NewShellKey creates a new ShellKey instance.
func NewShellKey(name string, path string, content string) *ShellKey {
	return &ShellKey{
		name:    name,
		path:    path,
		content: content,
	}
}

// GetName returns the name of the key.
func (key *ShellKey) GetName() string {
	return key.name
}

// GetPath returns the path of the key.
func (key *ShellKey) GetPath() string {
	return key.path
}

// GetFullPath returns the full path of the key.
func (key *ShellKey) GetFullPath() string {
	return path.Join(key.GetPath(), key.GetName())
}

// GetContent returns the content of the key.
func (key *ShellKey) GetContent() string {
	return key.content
}
