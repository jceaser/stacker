/*******************************************************************************
*******************************************************************************/

package main

/** run with go test ./tests/... */

import (
	"testing"

	"github.com/jceaser/stacker/stacker"
)

/**************************************/
// MARK - Functions

func TestHostSet(t *testing.T) {
	config := stacker.Config{}
	ß := stacker.HostSet(config)
	AssertEqual(t, "has correct count", 6, len(ß))
}
