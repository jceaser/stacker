/*
Public functions to interact with the Stacker system

Author: thomas.cherry@gmail.com
Copyright 2024, all rights reserved
*/

package stacker

import (
	"fmt"
	"log"
	"io"
	"io/ioutil"
	"net/http"
)

type WebEntry struct {
	Name   string
	Method   string
	Path string
	Help string
	Action func (w http.ResponseWriter, r *http.Request)
}

func send (w http.ResponseWriter, out string) {
	io.WriteString(w, fmt.Sprintf("%s\n", out))
}

func nop(w http.ResponseWriter, r *http.Request) {
}

func HostSet(config Config) []WebEntry {
	webCommands := []WebEntry{
		WebEntry{"List", "GET", "/list", "List stack items",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				send(w, ListItems(config))
			},
		},

		WebEntry{"Peek", "GET", "/", "Return top item on stack",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				send(w, PeekItem(config))
			},
		},

		WebEntry{"Rotate", "POST", "/rot", "Rotate Stack",
			func (w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				send(w, RotateUp(config))
			},
		},

		WebEntry{"Push", "PUT", "/",
			"Push value onto stack",
			func(w http.ResponseWriter, r *http.Request) {
				body, err := ioutil.ReadAll(r.Body)
		    	if err != nil {
        			return
    			}
				//CreateItem(config, r.FormValue("value"))
				CreateItem(config, string(body))
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				send(w, "")
			},
		},

		WebEntry{"Pop", "DELETE", "/", "Pop the stack",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				send(w, PopItem(config))
			},
		},

		WebEntry{"Summary", "HEAD", "/*", "Info on holdings",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Generate", "thomas.cherry@gmail.com")
				w.Header().Set("X-SHA-1", HashItems(config))
			},
		},
	}
	return webCommands
}

func Host(config Config) {
	handlers := HostSet(config)

	help := func (w http.ResponseWriter, r *http.Request) {
		for _, item := range handlers {
			line := "%6s - %-18s : %s\n"
			io.WriteString(w, fmt.Sprintf(line, item.Method, item.Path, item.Help))
		}
	}
	for _, item := range handlers {
		http.HandleFunc(fmt.Sprintf("%s %s", item.Method, item.Path), item.Action)
	}

	http.HandleFunc("GET /help", help)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("error:", err)
	}
}

