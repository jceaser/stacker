/*******************************************************************************
*******************************************************************************/

package main

import (
	"testing"
	"bytes"
)

/*
t.Errorf() needs to be called inside a helper function to ensure that multiple
calls are allowed to run in a single test. This version should work with any
comparable types
*/
func AssertEqual(t *testing.T, msg string, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Errorf("%s: [%v] vs [%v]", msg, expected, actual)
	}
}

/*
t.Errorf() needs to be called inside a helper function to ensure that multiple
calls are allowed to run in a single test. This version should work with any
comparable types
*/
func AssertNotEqual(t *testing.T, msg string, expected, actual any) {
	t.Helper()
	if expected == actual {
		t.Errorf("%s: [%v] vs [%v]", msg, expected, actual)
	}
}

/*
t.Errorf() needs to be called inside a helper function to ensure that multiple
calls are allowed to run in a single test. This version should work a single
boolean
*/
func AssertTrue(t *testing.T, msg string, actual bool) {
	t.Helper()
	if !actual {
		t.Errorf("%s: True vs [%v]", msg, actual)
	}
}

/*
t.Errorf() needs to be called inside a helper function to ensure that multiple
calls are allowed to run in a single test. This version should work a single
boolean
*/
func AssertBytesCompare(t *testing.T, msg string, expected, actual []byte) {
	t.Helper()

	result := bytes.Compare(expected, actual)

	if result != 0 {
		t.Errorf("%s: %v vs [%v]", msg, expected, actual)
	}
}

/*
t.Errorf() needs to be called inside a helper function to ensure that multiple
calls are allowed to run in a single test. This version does error and moves on
*/
func Error(t *testing.T, msg string) {
	t.Helper()
	t.Errorf(msg)
}
