package main

/*
This is the stacker command.

Author: thomas.cherry@gmail.com
Copyright 2023, all rights reserved
*/


import (
    "io"
	"flag"
	"fmt"
	"os"
	"github.com/jceaser/stacker/stacker"
)

/* Holds all application configuration and state, can be passed to lower
functions if needed */
type AppData struct {
	peekMode *bool
	listMode *bool
	updateMode *bool
	deleteMode *bool
	rotateMode *bool
	clearMode *bool
}

func setup() AppData {
	app_data := AppData{}
    app_data.peekMode = flag.Bool("peek", false, "peek mode")
    app_data.listMode = flag.Bool("ls", false, "List Everything")
    app_data.updateMode = flag.Bool("update", false, "Update an item")
    app_data.deleteMode = flag.Bool("delete", false, "Delete item")
    app_data.rotateMode = flag.Bool("rotate", false, "Rotate up")
    app_data.clearMode = flag.Bool("clear", false, "Clear all data")
    flag.Parse()

	return app_data
}

func main() {
	app_data := setup()
	action_taken := false

	if *app_data.peekMode {
		fmt.Println(stacker.PeekItem())
		action_taken = true
	}

	if *app_data.listMode {
		fmt.Print(stacker.ListItems())
		action_taken = true
	}

	if *app_data.updateMode {
		fmt.Println("updateMode not implimented")
	}

	if *app_data.deleteMode {
		stacker.DeleteItem()
		action_taken = true
	}

	if *app_data.clearMode {
		stacker.ClearAll()
		action_taken = true
	}

	if *app_data.rotateMode {
		stacker.RotateUp()
		action_taken = true
	}

	if action_taken {
		os.Exit(0)
	}

	/*
	Default action, if standard in has text, then assume we are in write mode
	otherwise assume read mode.
	*/
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		everything, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		text := string (everything)
		if len(text)>0 {
			//Write mode
			stacker.CreateItem(text)
		} else {
			fmt.Println(stacker.ReadItem())
		}
	} else {
		fmt.Println(stacker.ReadItem())
	}
}