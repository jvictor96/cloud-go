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
	LastPrint int
}

type Placing struct {
	ArtWork     *ArtWork
	Snapshot    []string
	AlphaFilter []string
	PosY        int
	Padding     int
	FirstFrame  int
}

func (e *Engine) Route(input []string) {
	e.Terminal.Buffer = input
	e.Galery.LoadArt()
	e.Galery.ArtWorks = e.Transformer.Resize(e.Galery.ArtWorks)
	e.Transformer.CalculateFrameCount(e.Galery.ArtWorks)
	for range e.Repetitions {
		placing := e.Placer.PlaceArt(e.Galery.ArtWorks, e.Terminal)
		frame_count := MaxFrameCount(placing)
		for i := range frame_count {
			e.Sleeper.Sleep(i)
			final_buffer := ManipulateBuffer(i, placing, e.Transformer, e.Terminal)
			fmt.Print(strings.Join(final_buffer, "\n") + "\n")
			fmt.Printf("\033[%dA", len(final_buffer))
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
	transformer.Transform(frame, mapa)

	for index, art := range mapa {

		for cursor < art.PosY {
			final_buffer = append(final_buffer, terminal.Buffer[cursor])
			cursor++
		}

		for cursor < (art.PosY + art.ArtWork.Height) {
			buffer := []rune(terminal.Buffer[cursor])
			artLine := []rune(strings.Repeat("f", art.Padding) + art.Snapshot[cursor-art.PosY])
			l := max(len(artLine), len(buffer))
			var line strings.Builder
			for i := range l {
				if i < len(artLine) && artLine[i] != 'f' {
					line.WriteRune(artLine[i])
				} else if i < len(buffer) {
					line.WriteRune(buffer[i])
				} else {
					line.WriteByte(' ')
				}
			}
			final_buffer = append(final_buffer, line.String())
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
