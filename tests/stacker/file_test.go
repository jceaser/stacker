/*******************************************************************************
*******************************************************************************/

package main

/** run with go test ./tests/... */

import (
	"testing"

	"bytes"
	"os"

	"github.com/jceaser/stacker/stacker"
)

/**************************************/
// MARK - Functions

func TestExpandPath(t *testing.T) {
	actual := stacker.ExpandPath("~")
	AssertTrue(t, "Tildi expand test", len(actual) > 0 && actual != "~")
}

func TestFileExists(t *testing.T) {
	actual := stacker.FileExists("/bin/sh")
	AssertTrue(t, "Does sh exists", actual)
}

func TestFileCrud(t *testing.T) {
	input := []byte("This is text\nfile content.")

	file, err := os.CreateTemp("", "stacker-file-*.txt")
	if err != nil {
		t.Errorf("Error getting temp file: %v", err)
	}

	file.Close()
	file_name := file.Name()
	stacker.FileSave(file_name, input)
	actual := stacker.FileRead(file_name)
	stacker.FileDelete(file_name)
	result := bytes.Compare(input, actual)

	AssertTrue(t, "File was not deleted", !stacker.FileExists(file_name))
	AssertEqual(t, "Content did not round trip through file", result, 0)
}
