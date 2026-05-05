package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	w, h, e := term.GetSize(0)
	if e != nil {
		fmt.Println(e)
	}
	galery := Galery{
		Path: fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".cloud/art/"),
	}
	galery.LoadArt()
	engine := Engine{
		Columns:  w, // Detectar via term.GetSize ou exec "tput cols"
		Lines:    h,
		Spacing:  2,
		MaxLines: 100,
		Galery:   galery,
	}

	command := os.Args[1]
	args := os.Args[2:]

	if command == "cfg" {
		galery.Route(args)
		return
	}
	engine.Route(command, args)
}
