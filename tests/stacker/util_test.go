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

func TestColor(t *testing.T) {
	ß := stacker.Color(32, "test")
	AssertEqual(t, "has a name", "\033[0;32mtest\033[0m", ß)
}

func TestTime_now_unix(t *testing.T) {
	stacker.TimeFrezer(1024)
	ß := stacker.Time_now_unix()
	AssertEqual(t, "Timecheck", int64(1024), ß)
}
