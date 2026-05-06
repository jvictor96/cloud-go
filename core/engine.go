package core

import (
	"fmt"
	"strings"
)

type Engine struct {
	Buffer      []string
	Galery      Galery
	Transformer Transformer
	Sleeper     Sleeper
	Placer      Placer
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

	e.Placer.PlaceArt(e)
	frame_count := e.Transformer.CalculateFrameCount(e)
	e.ManipulateBuffer(0)
	fmt.Print(strings.Join(e.FinalBuffer, "\n") + "\n")
	if !e.Dynamic {
		return
	}
	e.Sleeper.SetDuration(frame_count)
	for range 10 {
		e.Map = []Placing{}
		e.Placer.PlaceArt(e)
		for i := range frame_count {
			e.Sleeper.Sleep(i)
			fmt.Printf("\033[%dA", len(e.FinalBuffer))
			e.ManipulateBuffer(i)
			fmt.Print(strings.Join(e.FinalBuffer, "\n") + "\n")
		}
	}
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
