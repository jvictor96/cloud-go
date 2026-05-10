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
	e.Galery.ArtWorks = e.Transformer.Resize(e.Galery.ArtWorks)
	e.Transformer.CalculateFrameCount(e.Galery.ArtWorks)
	if !e.Terminal.Dynamic {
		placing := e.Placer.PlaceArt(e.Galery.ArtWorks, e.Terminal)
		final_buffer := ManipulateBuffer(0, placing, e.Transformer, e.Terminal)
		fmt.Print(strings.Join(final_buffer, "\n") + "\n")
		return
	}
	final_buffer := ManipulateBuffer(0, []Placing{}, e.Transformer, e.Terminal)
	fmt.Print(strings.Join(final_buffer, "\n") + "\n")
	for range e.Repetitions {
		placing := e.Placer.PlaceArt(e.Galery.ArtWorks, e.Terminal)
		frame_count := MaxFrameCount(placing)
		for i := range frame_count {
			e.Sleeper.Sleep(i)
			final_buffer := ManipulateBuffer(i, placing, e.Transformer, e.Terminal)
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

func ManipulateBuffer(frame int, mapa []Placing, transformer Transformer, terminal Terminal) []string {
	cursor := 0
	final_buffer := []string{}
	if terminal.Dynamic {
		for i := range mapa {
			mapa[i].Snapshot = []string{}
		}
		transformer.Transform(frame, mapa)
	} else {
		for art := range mapa {
			mapa[art].Snapshot = mapa[art].ArtWork.Content
		}
	}

	for index, art := range mapa {

		for cursor < art.PosY {
			final_buffer = append(final_buffer, terminal.Buffer[cursor])
			cursor++
		}

		for cursor < (art.PosY + art.ArtWork.Height) {
			buffer := terminal.Buffer[cursor]
			artLine := art.Snapshot[cursor-art.PosY]
			line := fmt.Sprintf("%-*s%s", art.Padding, buffer, artLine)
			final_buffer = append(final_buffer, line)
			cursor++
		}
		if index+1 < len(mapa) && mapa[index+1].PosY < cursor {
			for cursor < len(terminal.Buffer) {
				final_buffer = append(final_buffer, terminal.Buffer[cursor])
				cursor++
			}
			cursor = 0
			terminal.Buffer = final_buffer
			final_buffer = []string{}
		}
	}

	for cursor < len(terminal.Buffer) {
		final_buffer = append(final_buffer, terminal.Buffer[cursor])
		cursor++
	}
	return final_buffer
}
