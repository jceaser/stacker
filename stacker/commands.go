package stacker

import (
	"fmt"
	"os"
	"strings"
)

/*****************************************************************************/
//mark - Functions

func UserHomeDir() string {
	return "~/.config/stacker/data.json"
}

func CreateUserStoreDir() {
	config := fixPath("~/.config/stacker")
	if _, err := os.Stat(config) ; err != nil {
		os.MkdirAll(config, 0750)
	}
}

func LoadItemsFromDisk() (Stacker, bool) {
	CreateUserStoreDir()
	fileName := "~/.config/stacker/data.json"
	var data Stacker
	if FileExists(fileName) {
		raw_data := Read(fileName)
		data = *FromBytes(raw_data)
		return data, true
	}
	data.Name = "default"
	data.Clear()
	return data, false
}

/*****************************************************************************/
//mark - CRUD actions

func CreateItem(text string) {
	var data Stacker
	var exists bool
	if data, exists = LoadItemsFromDisk() ; exists {
		if 0<len(data.Items) {
			if data.Peek().Data == text {
				return
			}
		}
	}
	
	data.Push(MakeItem(text))
	Save(UserHomeDir(), data.ToJson())
}

func ReadItem() string {
	if data, exists := LoadItemsFromDisk() ; exists {
		item := data.Pop()
		Save(UserHomeDir(), data.ToJson())
		return item.Data
	}
	return ""
}

func updateItem(text string) {
	if data, exists := LoadItemsFromDisk() ; exists {
		data.Pop()
		data.Push(MakeItem(text))
		Save(UserHomeDir(), data.ToJson())
	}
}

func DeleteItem() {
	if data, exists := LoadItemsFromDisk() ; exists {
		data.Pop()
		Save(UserHomeDir(), data.ToJson())
	}
}

/*****************************************************************************/
//mark - Other actions

func PeekItem() string {
	if data, exists := LoadItemsFromDisk() ; exists {
		return data.Peek().Data
	}
	return ""
}

func ListItems() string {
	if data, exists := LoadItemsFromDisk() ; exists {
		list := []string{}
		for i, item := range data.Items {
			list = append(list, fmt.Sprintf("%d\t%s", i, item.Data))
		}
		return strings.Join(list, "\n")
	}
	return ""
}

func RotateUp() {
	if data, exists := LoadItemsFromDisk() ; exists {
		data.RotateUp()
		Save(UserHomeDir(), data.ToJson())
	}
}

func ClearAll() {
	if data, exists := LoadItemsFromDisk() ; exists {
		data.Clear()
		Save(UserHomeDir(), data.ToJson())
	}
}
