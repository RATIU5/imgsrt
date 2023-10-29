package utils

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"~", userHomeDir()},
		{"./test", filepath.Join(userHomeDir(), "test")},
		// ... add more test cases
	}

	for _, test := range tests {
		got, err := NormalizePath(test.input)
		if err != nil {
			t.Errorf("NormalizePath(%q) returned error: %v", test.input, err)
		}
		if got != test.want {
			t.Errorf("NormalizePath(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func userHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

func TestEnsureDir(t *testing.T) {
	testDir := filepath.Join(userHomeDir(), "testDir")
	defer os.RemoveAll(testDir) // cleanup

	err := EnsureDir(testDir)
	if err != nil {
		t.Fatalf("EnsureDir() returned error: %v", err)
	}

	_, err = os.Stat(testDir)
	if os.IsNotExist(err) {
		t.Errorf("Directory %q was not created", testDir)
	}
}

func TestIsDirEmpty(t *testing.T) {
	testDir := filepath.Join(userHomeDir(), "testDir")
	defer os.RemoveAll(testDir) // cleanup

	os.Mkdir(testDir, 0755)
	isEmpty, err := IsDirEmpty(testDir)
	if err != nil {
		t.Fatalf("IsDirEmpty() returned error: %v", err)
	}
	if !isEmpty {
		t.Errorf("IsDirEmpty() returned false, want true")
	}

	// Test case: non-empty directory
	err = ioutil.WriteFile(filepath.Join(testDir, "test.txt"), []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	isEmpty, err = IsDirEmpty(testDir)
	if err != nil {
		t.Fatalf("IsDirEmpty() returned error: %v", err)
	}
	if isEmpty {
		t.Errorf("IsDirEmpty() returned true, want false")
	}

	// Test case: directory does not exist
	nonExistentDir := filepath.Join(userHomeDir(), "nonExistentDir")
	isEmpty, err = IsDirEmpty(nonExistentDir)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if isEmpty {
		t.Errorf("IsDirEmpty() returned true, want false")
	}
}
