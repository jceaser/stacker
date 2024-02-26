package stacker

/*
These are very simple file management functions, mostly boiler plate code and
in no way perfect

Author: thomas.cherry@gmail.com
Copyright 2023, all rights reserved
*/

import (
	"io/ioutil"
	"os/user"
	"strings"
	"path/filepath"
	"os"
	"log"
)

/** resolve tildi to the full path for a user */
func fixPath (path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if path == "~" {
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}
	return path
}

/** test if a file exists */
func FileExists(fileName string) bool {
	fileName = fixPath(fileName)
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err) 
}

/** load a string from a file */
func Read(file string) []byte {
	file = fixPath(file)
    json_raw, err := os.Open(file)
    if err != nil {
        log.Printf("Error: %s\n", err)
        return nil
    } else {
        defer json_raw.Close()
        bytes, err := ioutil.ReadAll(json_raw)
        if err!=nil {
            log.Printf("Error: %s\n", err)
        } else {
            return bytes
        }
    }
    return nil
}

/** save the database to a file */
func Save(file string, json_text []byte) {
    var err error
    file = fixPath(file)
    err = ioutil.WriteFile(file, json_text, 0644)
    if err!=nil {
        log.Printf("Write Error: %s - %s\n", file, err)
    } else {
        //log.Printf("File %s has been saved\n", file)
    }
}
