package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"unicode/utf8"
)

type Engine struct {
	Buffer      []string
	Galery      Galery
	Transformer Transformer
	Map         []Placing
	FinalBuffer []string
	Columns     int
	Lines       int
	MaxLines    int
	Spacing     int
	LastPrint   int
}

type Placing struct {
	ArtWork  *ArtWork
	Snapshot []string
	PosY     int
	Padding  int
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
	frame_count := e.Transformer.CalculateFrameCount(e)
	e.ManipulateBuffer(0)
	fmt.Print(strings.Join(e.FinalBuffer, "\n"))
	for i := range frame_count {
		e.Transformer.Sleep()
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
						Padding: minDif + fuzz,
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

	e.Transformer.Transform(frame, e)

	for _, art := range e.Map {

		for cursor < art.PosY {
			e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
			cursor++
		}

		for cursor < (art.PosY + art.ArtWork.Height) {
			line := fmt.Sprintf("%-*s%s", art.Padding, e.Buffer[cursor], art.Snapshot[cursor-art.PosY])
			e.FinalBuffer = append(e.FinalBuffer, line)
			cursor++
		}
	}

	for cursor < len(e.Buffer) {
		e.FinalBuffer = append(e.FinalBuffer, e.Buffer[cursor])
		cursor++
	}
}

func (e *Engine) ManipulateArts2(frame int) {
	for i := range e.Map {
		e.Map[i].Snapshot = []string{}
		for art_index, line := range e.Map[i].ArtWork.Content {
			if art_index < frame {
				e.Map[i].Snapshot = append(e.Map[i].Snapshot, line)
			} else {
				e.Map[i].Snapshot = append(e.Map[i].Snapshot, strings.Repeat(" ", utf8.RuneCountInString(line)))
			}
		}
	}
}

func (e *Engine) ManipulateArts3(frame int) {
	for i := range e.Map {
		e.Map[i].Snapshot = []string{}
		for _, line := range e.Map[i].ArtWork.Content {
			runes := []rune(line)
			frame = frame % len(runes)
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, string(runes[frame:])+string(runes[:frame]))
		}
	}
}
