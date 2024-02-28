/*
Stacker is a simple stack for storing clipboard strings held in Item

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"encoding/json"
	"log"
	"time"
)

/******************************************************************************/
//MARK: - Structures

/*
Stacker is the entire data structure for the file system, it contains a list of
Items
*/
type Stacker struct {
	Name string
	Items []Item
}

/*
One node in the Structure, has a time stamp for judging how long item is been
stored.
*/
type Item struct {
	Time int64
	Data string
}

/* Creates a storable Json record for Stacker */
func (data Stacker) ToJson() []byte {
	data_as_json, err := json.MarshalIndent (data, "", "  ")
	if err != nil {
		log.Println (err)
	}
	return data_as_json
}

/* Make sure that the Items is initialized and an empty list */
func (data *Stacker) ensure() {
	if data.Items == nil {
		data.Items = []Item{}
	}
}

/* Take the difference of the Item time stamp from now */
func (self Item) FromNow() int {
	return int(time.Now().Unix() - self.Time)
}

/******************************************************************************/
//MARK: - Stack Operations

/* Reset the stack back to an empty state */
func (data *Stacker) Clear() {
	data.Items = []Item{}
}

/* Put an item onto the stack */
func (data *Stacker) Push(item Item) {
	data.ensure()
	data.Items = append(data.Items, item)
}

/* Remove an item from the stack and return it */
func (data *Stacker) Pop() Item {
	l := len(data.Items)
	var found Item
	if 0<l {
		n := l-1
		found = data.Items[n]
		data.Items = data.Items[:n]
	}
	return found
}

/* Remove the top item from the stack */
func (data *Stacker) PopQueue() Item {
	var ret Item
	if 1 < len(data.Items) {
		last := data.Items[0]
		data.Items = data.Items[1:]
		return last
	}
	return ret
}

/* Return the top item from the stack without removing it */
func (data Stacker) Peek() Item {
    l := len(data.Items)
    var found Item
    if 0<l {
    	found = data.Items[l-1]
    }
	return found
}

/* replace the current top of the stack and replace it */
func (data *Stacker) Replace(item Item) Item {
	l := len(data.Items)
	found := data.Items[l-1]
	data.Items[l-1] = item
	return found
}

/* Swap the top two item on the stack */
func (data *Stacker) Swap() {
	l := len(data.Items)
	if 1<l {
		was_last := data.Pop()
		was_next := data.Pop()
		data.Push(was_last)
		data.Push(was_next)
	}
}

/* take the last item on the stack and make it the first */
func (data *Stacker) RotateUp() {
	top := data.PopQueue()
	data.Push(top)
}

/******************************************************************************/
//MARK: - Other Operations

/* Create a new Item struct using the current time as the time stamp */
func MakeItem(data string) Item {
	item := Item{Time_now_unix(), data}
	return item
}

/* Create a stacker struct from a byte array (json) */
func FromBytes(raw []byte) *Stacker {
	data := &Stacker{}
	err := json.Unmarshal(raw, data)
	if err !=nil {
		log.Println(err)
	}
	data.ensure()
	return data
}
