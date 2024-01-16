package utils

import (
	"os/user"
	"path/filepath"
	"strings"
)

// ExpandHomeDir expands the home directory.
func ExpandHomeDir(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		dir := usr.HomeDir
		path = filepath.Join(dir, path[2:])
	}
	return path, nil
}
