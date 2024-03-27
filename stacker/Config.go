/*
Public functions to interact with the Stacker system

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"fmt"
	"os"
)

type Config struct {
	Name   string
	Path   string
	Encode bool
}

func (self Config) UserHomeDir() string {
	if 0 < len(self.Path) {
		return self.Path
	}
	return fmt.Sprintf("./%s.json", os.Args[0])
}
