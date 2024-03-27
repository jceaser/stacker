/*
This is the stacker command.

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package main

import (
	"flag"
	"fmt"
	"github.com/jceaser/stacker/stacker"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/* ************************************************************************** */
//MARK: - Structs

/*
	Holds all application configuration and state, can be passed to lower

functions if needed
*/
type AppData struct {
	peekMode    *bool
	listMode    *bool
	updateMode  *bool
	deleteMode  *bool
	rotateMode  *bool
	clearMode   *bool
	versionMode *bool
	http        *bool
	cleanLimit  *int

	encode *bool
	name   *string
	path   *string
}

func (self AppData) Context() stacker.Config {
	return stacker.Config{Name: *self.name, Path: *self.path, Encode: *self.encode}
}

/* ************************************************************************** */
//MARK: - Server

func host(config stacker.Config) {
	stacker.Host(config)
}

/* ************************************************************************** */
//MARK: - Functions

/*
export XDG_CONFIG_HOME=~/.config/test ; go run stacker.go
*/
func FindConfigPath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")                //standard location
	appConfigPath := filepath.Base(os.Args[0]) + "/data.json" //sub directory
	if len(configHome) < 1 {
		configHome = "~/.config" //assume things
	}
	relativeConfigPath := fmt.Sprintf("%s/%s", configHome, appConfigPath)
	configPath := stacker.ExpandPath(relativeConfigPath)
	return configPath
}

/* create app data from the command line flags */
func Setup() AppData {
	app_data := AppData{}
	app_data.peekMode = flag.Bool("peek", false, "peek mode")
	app_data.listMode = flag.Bool("ls", false, "List Everything")
	app_data.updateMode = flag.Bool("update", false, "Update an item")
	app_data.deleteMode = flag.Bool("delete", false, "Delete item")
	app_data.rotateMode = flag.Bool("rotate", false, "Rotate up")
	app_data.clearMode = flag.Bool("clear", false, "Clear all data")
	app_data.versionMode = flag.Bool("version", false, "Clear all data")

	app_data.http = flag.Bool("http", false, "Create HTTP host")
	app_data.cleanLimit = flag.Int("clean", -1, "Clean old data")
	app_data.encode = flag.Bool("encode", false, "Mask sensitive data in list and file")
	app_data.name = flag.String("name", "default", "Stack Name")
	app_data.path = flag.String("path", "", "full Path to stack file")
	flag.Parse()

	if len(*app_data.path) < 1 {
		default_path := FindConfigPath()
		app_data.path = &default_path
	}

	return app_data
}

/*
Primary action for the app, if a stream exists, then push it to the stack, if on
stream then pop from the stack.
*/
func StreamAction(app_data AppData) {
	/*
		Default action, if standard in has text, then assume we are in write mode
		otherwise assume read mode.
	*/
	config := app_data.Context()
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		everything, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		text := strings.TrimSpace(string(everything))
		if len(text) > 0 {
			//Write mode
			if *app_data.updateMode {
				fmt.Println(stacker.UpdateItem(config, text))
			} else {
				stacker.CreateItem(config, text)
			}
		} else {
			//no data, so fall back to a read
			fmt.Println(stacker.ReadItem(config))
		}
	} else {
		fmt.Println(stacker.ReadItem(config))
	}
}

func main() {
	app_data := Setup()

	if *app_data.versionMode {
		fmt.Println("stacker 1.0 by thomas.cherry@gmail.com")
		fmt.Printf("Config file: [%s]\n", *app_data.path)
		os.Exit(0)
	}

	config := app_data.Context()

	if *app_data.http {
		host(config)
	} else if *app_data.listMode {
		fmt.Print(stacker.ListItems(config))
	} else if *app_data.peekMode {
		fmt.Println(stacker.PeekItem(config))
	} else if *app_data.deleteMode {
		fmt.Println(stacker.DeleteItem(config))
	} else if *app_data.clearMode {
		stacker.ClearAll(config)
	} else if *app_data.rotateMode {
		fmt.Println(stacker.RotateUp(config))
	} else if *app_data.cleanLimit > 0 {
		stacker.RemoveOld(config, *app_data.cleanLimit)
	} else {
		StreamAction(app_data)
	}
}
