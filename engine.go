package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

type Engine struct {
	Buffer      []string
	Galery      Galery
	Map         []Placing
	FinalBuffer []string
	Columns     int
	Lines       int
	MaxLines    int
	Spacing     int
	LastPrint   int
}

type Placing struct {
	ArtWork *ArtWork
	PosY    int
	MinDif  int
	Fuzz    int
}

func (e *Engine) Route(command string, args []string) {
	cmd := exec.Command(command, args...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	e.Buffer = strings.Split(string(out), "\n")
	e.Map = []Placing{}

	for e.PlaceImages() {
	}
	frame_count := 0
	for _, art := range e.Map {
		if art.ArtWork.Height > frame_count {
			frame_count = art.ArtWork.Height
		}
	}
	e.ManipulateBuffer(0)
	fmt.Print(strings.Join(e.FinalBuffer, "\n"))
	for i := range frame_count {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("\033[%dA", len(e.FinalBuffer)-1)
		e.ManipulateBuffer(i)
		fmt.Print(strings.Join(e.FinalBuffer, "\n"))
	}
}

func (e *Engine) PlaceImages() bool {
	modified := false
	for i := range e.Galery.ArtWorks {
		art := &e.Galery.ArtWorks[i]
		height := 0
		pos := 0
		minDif := 0
		startingPoint := e.LastPrint + 1 + e.Spacing

		for cursor, line := range e.Buffer {
			// Em Go, len(line) lida com caracteres, não precisamos de ghost_bytes do wc
			lineLen := len(line)
			if (e.Columns-lineLen > art.Width) && (cursor >= startingPoint) {
				height++
				if minDif < lineLen {
					minDif = lineLen
				}
				if height > art.Height {
					e.LastPrint = pos + art.Height + e.Spacing
					fuzz := rand.Intn(e.Columns - minDif - art.Width - 1)
					e.Map = append(e.Map, Placing{
						ArtWork: art,
						PosY:    pos,
						MinDif:  minDif,
						Fuzz:    fuzz,
					})
					modified = true
					break
				}
			} else {
				pos = cursor
				minDif = lineLen
				height = 0
			}
		}
	}
	return modified
}

func (e *Engine) ManipulateBuffer(frame int) {
	e.FinalBuffer = []string{}
	cursor := 0

	for _, art := range e.Map {

		// Preenche as linhas antes da imagem
		for cursor < art.PosY {
			e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
			cursor++
		}

		// Desenha a imagem com o efeito de frame (scroll-in)
		for cursor < (art.PosY + art.ArtWork.Height) {
			relativeLine := cursor - art.PosY
			threshold := art.ArtWork.Height - frame
			if threshold < 0 {
				threshold = 0
			}

			artLine := ""
			if relativeLine >= threshold {
				artLine = art.ArtWork.Content[relativeLine-threshold]
			}

			// Montagem da linha: Buffer original + Padding + Arte
			padding := art.MinDif + art.Fuzz
			line := fmt.Sprintf("%-*s%s", padding, e.Buffer[cursor], artLine)
			e.FinalBuffer = append(e.FinalBuffer, line)
			cursor++
		}
	}

	// Adiciona o restante do buffer original
	for cursor < len(e.Buffer) {
		e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
		cursor++
	}
}
