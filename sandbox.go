package main

import (
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/term"
)

func main() {
	w, h, e := term.GetSize(0)
	if e != nil {
		fmt.Println(e)
	}
	engine := Engine{
		Columns:  w, // Detectar via term.GetSize ou exec "tput cols"
		Lines:    h,
		Spacing:  2,
		MaxLines: 100,
	}
	fmt.Println(engine.Columns, engine.Lines)

	args := os.Args[2:]
	cmd := exec.Command(os.Args[1], args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
