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
	Repetitions int
}

type Terminal struct {
	Columns   int
	Lines     int
	Buffer    []string
	Dynamic   bool
	LastPrint int
}

type Placing struct {
	ArtWork    *ArtWork
	Snapshot   []string
	PosY       int
	Padding    int
	FirstFrame int
}

func (e *Engine) Route(input []string) {
	e.Terminal.Buffer = input
	e.Galery.LoadArt()
	e.Transformer.CalculateFrameCount(e.Galery.ArtWorks)
	placing := e.Placer.PlaceArt(e.Galery.ArtWorks, e.Terminal)
	final_buffer := e.ManipulateBuffer(0, placing)
	fmt.Print(strings.Join(final_buffer, "\n") + "\n")
	if !e.Terminal.Dynamic {
		return
	}
	for range e.Repetitions {
		placing := e.Placer.PlaceArt(e.Galery.ArtWorks, e.Terminal)
		frame_count := MaxFrameCount(placing)
		for i := range frame_count {
			e.Sleeper.Sleep(i)
			final_buffer := e.ManipulateBuffer(i, placing)
			fmt.Printf("\033[%dA", len(final_buffer))
			fmt.Print(strings.Join(final_buffer, "\n") + "\n")
		}
	}
}

func MaxFrameCount(placing []Placing) int {
	fc := 0
	for _, p := range placing {
		fc = max(fc, p.FirstFrame+p.ArtWork.FrameCount)
	}
	return fc
}

func (e *Engine) ManipulateBuffer(frame int, mapa []Placing) []string {
	cursor := 0
	final_buffer := []string{}
	if e.Terminal.Dynamic {
		for i := range mapa {
			mapa[i].Snapshot = []string{}
		}
		e.Transformer.Transform(frame, mapa)
	} else {
		for art := range mapa {
			mapa[art].Snapshot = mapa[art].ArtWork.Content
		}
	}

	for _, art := range mapa {

		for cursor < art.PosY {
			final_buffer = append(final_buffer, e.Terminal.Buffer[cursor])
			cursor++
		}

		for cursor < (art.PosY + art.ArtWork.Height) {
			line := fmt.Sprintf("%-*s%s", art.Padding, e.Terminal.Buffer[cursor], art.Snapshot[cursor-art.PosY])
			final_buffer = append(final_buffer, line)
			cursor++
		}
	}

	for cursor < len(e.Terminal.Buffer) {
		final_buffer = append(final_buffer, e.Terminal.Buffer[cursor])
		cursor++
	}
	return final_buffer
}
