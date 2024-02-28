/*
Public functions to interact with the Stacker system

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"path/filepath"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Name string
	Path string
}

func (self Config) UserHomeDir() string {
    if len(self.Path)<1 {
        return self.Path
    }
	return fmt.Sprintf("./%s.json", os.Args[0])
}

/* ************************************************************************** */
//MARK: - Functions

/* Create the storage path for  */
func createUserStoreDir(cxt Config) {
	config := ExpandPath(filepath.Dir(cxt.UserHomeDir()))
	if _, err := os.Stat(config) ; err != nil {
		os.MkdirAll(config, 0750)
	}
}

/*
Load a storage file and return a stacker struct and a bool. The bool is true if
config already existing, false if it is assumed.
*/
func loadItemsFromDisk(cxt Config) (Stacker, bool) {
	createUserStoreDir(cxt)
	fileName := cxt.UserHomeDir()
	var data Stacker
	if FileExists(fileName) {
		raw_data := FileRead(fileName)
		data = *FromBytes(raw_data)
		return data, true
	}
	data.Name = "default"
	data.Clear()
	return data, false
}

/* Save a Stacker image to disk */
func saveItemsToDisk(cxt Config, data Stacker) {
	FileSave(cxt.UserHomeDir(), data.ToJson())
}

/*****************************************************************************/
//MARK: - Public CRUD actions

/* Create an Item in the Stack */
func CreateItem(cxt Config, text string) {
	var data Stacker
	var exists bool
	if data, exists = loadItemsFromDisk(cxt) ; exists {
		if 0<len(data.Items) {
			if data.Peek().Data == text {
				return
			}
		}
	}
	data.Push(MakeItem(text))
	saveItemsToDisk(cxt, data)
}

/* Read an Item from the Stack */
func ReadItem(cxt Config) string {
	var value string
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		item := data.Pop()
		saveItemsToDisk(cxt, data)
		value = item.Data
	}
	return value
}

/* Update the top Item on the Stack */
func UpdateItem(cxt Config, text string) string {
	var old string
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		item := data.Pop()
		old = item.Data
		data.Push(MakeItem(text))
		saveItemsToDisk(cxt, data)
	}
	return old
}

/* Delete the top Item in the Stack */
func DeleteItem(cxt Config) string {
	var old string
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		item := data.Pop()
		old = item.Data
		saveItemsToDisk(cxt, data)
	}
	return old
}

/*****************************************************************************/
//MARK: - Other Public actions

/* Return the first item off the stack without doing a Pop */
func PeekItem(ctx Config) string {
	var value string
	if data, exists := loadItemsFromDisk(ctx) ; exists {
		value = data.Peek().Data
	}
	return value
}

/* Return a formated list of stack items */
func ListItems(cxt Config) string {
	var everything string
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		list := []string{}
		for i, item := range data.Items {
		    text := strings.TrimSpace(item.Data)
		    id := color(32, fmt.Sprintf("%d", i))
			list = append(list, fmt.Sprintf("%s\t%s", id, text))
		}
		everything = strings.Join(list, "\n")
	}
	return everything
}

/* Rotate the stack up, taking the last item and making it the first */
func RotateUp(cxt Config) string {
	var top string
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		data.RotateUp()
		saveItemsToDisk(cxt, data)
		top = data.Peek().Data
	}
	return top
}

/* Clear all items off the stack */
func ClearAll(cxt Config) string {
	if data, exists := loadItemsFromDisk(cxt) ; exists {
		data.Clear()
		saveItemsToDisk(cxt, data)
	}
	return ""
}

/* Remove items older then the limit */
func RemoveOld(cxt Config, limit int) {
	if limit > 0 {
		if data, exists := loadItemsFromDisk(cxt) ; exists {
			for i:=len(data.Items)-1 ; i>=0 ; i-- {
				item := data.Items[i]
				if item.FromNow() > limit {
					fmt.Printf("%s %d-%s\n", color(33, "found an old one:"),
					    item.Time, item.Data)
				}
			}
		}
	}
}
