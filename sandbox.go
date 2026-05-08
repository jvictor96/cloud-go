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
	dynamic := false
	if term.IsTerminal(int(os.Stdout.Fd())) {
		dynamic = true
	}
	engine := core.Engine{
		Terminal:    core.Terminal{Columns: w, Lines: h, Dynamic: dynamic},
		Galery:      core.Galery{Path: fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".cloud/art/")},
		Transformer: &transformer.Float{},
		Placer:      &placer.FillUnsync{FillOnce: placer.FillOnce{Spacing: 3, Chance: 0.25}},
		Sleeper:     &ticker.Linear{Speed: 100},
		Repetitions: 4,
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
		engine.Galery.Route(args)
	}
}
