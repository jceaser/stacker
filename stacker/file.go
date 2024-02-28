/*
These are very simple file management functions, mostly boiler plate code and
in no way perfect

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

/** resolve tilde to the full path for a user */
func ExpandPath(path string) string {
	usr, _ := user.Current()
	home := usr.HomeDir
	if path == "~" {
		path = home
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(home, path[2:])
	}
	return path
}

/** test if a file exists */
func FileExists(fileName string) bool {
	fileName = ExpandPath(fileName)
	_, err := os.Stat(fileName)
	return !errors.Is(err, os.ErrNotExist)
}

/** load a string from a file */
func FileRead(file string) []byte {
	file = ExpandPath(file)
	json_raw, err := os.Open(file)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil
	} else {
		defer json_raw.Close()
		bytes, err := ioutil.ReadAll(json_raw)
		if err != nil {
			log.Printf("Error: %s\n", err)
		} else {
			return bytes
		}
	}
	return nil
}

/** save the database to a file */
func FileSave(file string, json_text []byte) {
	var err error
	file = ExpandPath(file)
	err = ioutil.WriteFile(file, json_text, 0644)
	if err != nil {
		log.Printf("Write Error: %s - %s\n", file, err)
	}
}

func FileDelete(file string) {
	file = ExpandPath(file)
	err := os.Remove(file)
	if err != nil {
		log.Printf("Delete Error: %s - %s\n", file, err)
	}
}
