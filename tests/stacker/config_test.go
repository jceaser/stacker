/*******************************************************************************
*******************************************************************************/

package main

/** run with go test ./tests/... */

import (
	"strings"
	"testing"

	"github.com/jceaser/stacker/stacker"
)

/**************************************/
// MARK - Functions

func TestConfigPath(t *testing.T) {
	config := stacker.Config{}
	ß := config.UserHomeDir()
	AssertEqual(t, "Has home dir", true, strings.HasSuffix(ß, "/stacker.test.json"))
	config.Path = "~/stacker.test2.json"
	AssertEqual(t, "Can set home", "~/stacker.test2.json", config.UserHomeDir())
}
