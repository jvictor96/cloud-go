package core

import (
	"fmt"
	"strings"
)

type Engine struct {
	Galery      Galery
	Transformer Transformer
	Sleeper     Sleeper
	Placer      Placer
	Terminal    Terminal
	Map         []Placing
	Repetitions int
}

type Terminal struct {
	Columns     int
	Lines       int
	Buffer      []string
	FinalBuffer []string
	Dynamic     bool
	LastPrint   int
}

type Placing struct {
	ArtWork    *ArtWork
	Snapshot   []string
	PosY       int
	Padding    int
	FrameCount int
}

func (e *Engine) Route(input []string) {
	e.Terminal.Buffer = input
	e.Map = []Placing{}

	e.Galery.LoadArt()
	e.PlaceArt()
	e.ManipulateBuffer(0)
	fmt.Print(strings.Join(e.Terminal.FinalBuffer, "\n") + "\n")
	if !e.Terminal.Dynamic {
		return
	}
	for range e.Repetitions {
		e.Map = []Placing{}
		frame_count := e.PlaceArt()
		for i := range frame_count {
			e.Sleeper.Sleep(i)
			fmt.Printf("\033[%dA", len(e.Terminal.FinalBuffer))
			e.ManipulateBuffer(i)
			fmt.Print(strings.Join(e.Terminal.FinalBuffer, "\n") + "\n")
		}
	}
}

func (e *Engine) PlaceArt() int {
	e.Placer.PlaceArt(e)
	max_frame_count := 0
	for art := range e.Map {
		frame_count := e.Transformer.CalculateFrameCount(*e.Map[art].ArtWork)
		e.Map[art].FrameCount = frame_count
		max_frame_count = max(frame_count, max_frame_count)
	}
	return max_frame_count
}

func (e *Engine) ManipulateBuffer(frame int) {
	e.Terminal.FinalBuffer = []string{}
	cursor := 0
	if e.Terminal.Dynamic {
		e.Transformer.Transform(frame, e)
	} else {
		for art := range e.Map {
			e.Map[art].Snapshot = e.Map[art].ArtWork.Content
		}
	}

	for _, art := range e.Map {

		for cursor < art.PosY {
			e.Terminal.FinalBuffer = append(e.Terminal.FinalBuffer, e.Terminal.Buffer[cursor])
			cursor++
		}

		for cursor < (art.PosY + art.ArtWork.Height) {
			line := fmt.Sprintf("%-*s%s", art.Padding, e.Terminal.Buffer[cursor], art.Snapshot[cursor-art.PosY])
			e.Terminal.FinalBuffer = append(e.Terminal.FinalBuffer, line)
			cursor++
		}
	}

	for cursor < len(e.Terminal.Buffer) {
		e.Terminal.FinalBuffer = append(e.Terminal.FinalBuffer, e.Terminal.Buffer[cursor])
		cursor++
	}
}
