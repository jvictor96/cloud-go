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
	w, _, _ := term.GetSize(int(f.Fd()))
	var transformer_impl core.Transformer
	transformer_impl = &transformer.Static{}
	reps := 1
	if term.IsTerminal(int(os.Stdout.Fd())) {
		transformer_impl = &transformer.KeepFloat{Amplitude: 5}
		reps = 3
	}
	engine := core.Engine{
		Terminal:    core.Terminal{Columns: w},
		Galery:      core.Galery{Path: fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".cloud/art/")},
		Transformer: transformer_impl,
		Placer:      &placer.DoFill{FillOnce: placer.FillOnce{Spacing: 3, Chance: 0.25}},
		Sleeper:     &ticker.Linear{Speed: 400},
		Repetitions: reps,
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
