package utils

import (
	"os/user"
	"path/filepath"
	"testing"
)

// TestExpandHomeDir tests the expandHomeDir function.
func TestExpandHomeDir(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}
	homeDir := usr.HomeDir

	// Define test cases
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "Path with tilde",
			input:   "~/testpath",
			want:    filepath.Join(homeDir, "testpath"),
			wantErr: false,
		},
		{
			name:    "Path without tilde",
			input:   "/tmp/testpath",
			want:    "/tmp/testpath",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandHomeDir(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandHomeDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("expandHomeDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
