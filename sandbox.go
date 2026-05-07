package main

import (
	"bufio"
	"cloud/core"
	"cloud/placer"
	"cloud/ticker"
	"cloud/transformer"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	f, _ := os.Open("/dev/tty")
	w, h, _ := term.GetSize(int(f.Fd()))
	galery := core.Galery{
		Path: fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".cloud/art/"),
	}
	galery.LoadArt()
	dynamic := false
	if term.IsTerminal(int(os.Stdout.Fd())) {
		dynamic = true
	}
	engine := core.Engine{
		Columns:     w, // Detectar via term.GetSize ou exec "tput cols"
		Lines:       h,
		Dynamic:     dynamic,
		Galery:      galery,
		Transformer: &transformer.Float{},
		Placer:      &placer.FillOnce{},
		Sleeper:     &ticker.Linear{Speed: 100},
	}

	var input []string
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Estamos recebendo dados via PIPE (ex: cat arquivo | engine)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		for line := range input {
			input[line] = strings.ReplaceAll(input[line], "\t", "    ")
		}

		engine.Route(input)
	} else {
		args := os.Args[1:]
		galery.Route(args)
	}
}
