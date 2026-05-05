package main

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode/utf8"
)

type Engine struct {
	Buffer      []string
	Galery      Galery
	Transformer Transformer
	Sleeper     Sleeper
	Map         []Placing
	FinalBuffer []string
	Columns     int
	Lines       int
	Dynamic     bool
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

func (e *Engine) Route(input []string) {
	e.Buffer = input
	e.Map = []Placing{}

	for e.PlaceImages() {
	}
	frame_count := e.Transformer.CalculateFrameCount(e)
	e.ManipulateBuffer(0)
	fmt.Print(strings.Join(e.FinalBuffer, "\n"))
	if !e.Dynamic {
		return
	}
	for i := range frame_count {
		e.Sleeper.Sleep(i)
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
			lineLen := utf8.RuneCountInString(line)
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
	if e.Dynamic {
		e.Transformer.Transform(frame, e)
	} else {
		for art := range e.Map {
			e.Map[art].Snapshot = e.Map[art].ArtWork.Content
		}
	}

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
