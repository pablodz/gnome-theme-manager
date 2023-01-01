package utils

import (
	"os"
	"testing"
)

func TestGetDirectories(t *testing.T) {
	// Set up a test directory with some subdirectories and files
	path := "test_dir"
	os.Mkdir(path, 0755)
	os.Mkdir(path+"/dir1", 0755)
	os.Mkdir(path+"/dir2", 0755)
	os.Create(path + "/file1.txt")
	os.Create(path + "/file2.txt")

	// Call the GetDirectories function with the test directory
	directories, err := GetDirectories(path)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check that the returned slice of directories has the expected length
	if len(directories) != 2 {
		t.Errorf("Expected 2 directories, got %d", len(directories))
	}

	// Check that the returned slice of directories contains the expected names
	expectedDirectories := []string{"dir1", "dir2"}
	for _, expected := range expectedDirectories {
		found := false
		for _, actual := range directories {
			if actual == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected directory '%s' not found in returned slice", expected)
		}
	}

	// Clean up the test directory
	os.RemoveAll(path)
}
