package stacker

/*
Stacker is a simple stack for storing clipboard strings held in Item

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

import (
	"encoding/json"
	"log"
	//"fmt"
	"time"
)

/******************************************************************************/
// #mark - Structures

type Stacker struct {
	Name string
	Items []Item
}

type Item struct {
	Time int64
	Data string
}

func MakeItem(data string) Item {
	item := Item{time.Now().Unix(), data}
	return item
}

/******************************************************************************/
//mark - Stack Operations */

func (data *Stacker) ensure() {
	if data.Items == nil {
		data.Clear()
	}
}

func (data *Stacker) Clear() {
	data.Items = []Item{}
}

func (data *Stacker) Push(item Item) {
	data.ensure()
	data.Items = append(data.Items, item)
}

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

func (data *Stacker) PopQueue() Item {
	var ret Item
	if 1 < len(data.Items) {
		last := data.Items[0]
		data.Items = data.Items[1:]
		return last
	}
	return ret
}

func (data Stacker) Peek() Item {
    l := len(data.Items)
    var found Item
    if 0<l {
    	found = data.Items[l-1]
    }
	return found
}

func (data *Stacker) Replace(item Item) Item {
	l := len(data.Items)
	found := data.Items[l-1]
	data.Items[l-1] = item
	return found
}

func (data *Stacker) Swap() {
	l := len(data.Items)
	if 1<l {
		was_last := data.Pop()
		was_next := data.Pop()
		data.Push(was_last)
		data.Push(was_next)
	}
}

func (data *Stacker) RotateUp() {
	top := data.PopQueue()
	data.Push(top)
}

/******************************************************************************/
// #mark - Stack Operations */

func (data Stacker) ToJson() []byte {
	data_as_json, err := json.MarshalIndent (data, "", "  ")
	if err != nil {
		log.Println (err)
	}
	return data_as_json
}

func FromBytes(raw []byte) *Stacker {
	data := &Stacker{}
	err := json.Unmarshal(raw, data)
	if err !=nil {
		log.Println(err)
	}
	return data
}
