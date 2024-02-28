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

var sample = []byte(`{
  "Name": "default",
  "Items": [
	{
	  "Time": 1709054366,
	  "Data": "d1"
	},
	{
	  "Time": 1709054371,
	  "Data": "d2"
	},
	{
	  "Time": 1709054380,
	  "Data": "d3"
	}
  ]
}`)

func TestBasicStack(t *testing.T) {
	s := stacker.FromBytes(sample)
	AssertEqual(t, "has a name", "default", s.Name)
}

func TestStackRoundTrip(t *testing.T) {
	obj := stacker.FromBytes(sample)
	result := obj.ToJson()
	AssertBytesCompare(t, "round trip", sample, result)
}

func TestStackRotateUp(t *testing.T) {
	obj := stacker.FromBytes(sample)

	obj.RotateUp()
	item := obj.Peek()
	AssertEqual(t, "was rotated", "d1", item.Data)
	AssertEqual(t, "should not drop any", 3, len(obj.Items))
}

func TestStackSwap(t *testing.T) {
	obj := stacker.FromBytes(sample)

	obj.Swap()
	item := obj.Peek()
	AssertEqual(t, "was swap", "d2", item.Data)
	AssertEqual(t, "should not drop any", 3, len(obj.Items))
}

func TestStackReplace(t *testing.T) {
	obj := stacker.FromBytes(sample)

	item := obj.Replace(stacker.MakeItem("Replace Item"))
	AssertEqual(t, "old value", "d3", item.Data)
	AssertEqual(t, "new value", "Replace Item", obj.Peek().Data)
	AssertEqual(t, "should not drop any", 3, len(obj.Items))
}

func TestStackPopQueue(t *testing.T) {
	obj := stacker.FromBytes(sample)

	item := obj.PopQueue()
	AssertEqual(t, "old value", "d1", item.Data)
	AssertEqual(t, "should drop one", 2, len(obj.Items))
}

func TestStackPop(t *testing.T) {
	obj := stacker.FromBytes(sample)

	item := obj.Pop()
	AssertEqual(t, "old value", "d3", item.Data)
	AssertEqual(t, "should drop one", 2, len(obj.Items))
}

func TestStackClear(t *testing.T) {
	obj := stacker.FromBytes(sample)

	obj.Clear()
	AssertEqual(t, "should drop all", 0, len(obj.Items))
}
