package core

import (
	"strings"
)

type Placer interface {
	PlaceArt(artWorks []ArtWork, terminal Terminal) []Placing
}

type Sleeper interface {
	Sleep(frame int)
}

type Transformer interface {
	Transform(frame int, art_works []Placing) []Placing
	CalculateFrameCount(art_work []ArtWork) int
	Resize(art []ArtWork) []ArtWork
}

func FillAbove(buffer []string, width int, height int) []string {
	for len(buffer) < height {
		buffer = append([]string{strings.Repeat(" ", width)}, buffer...)
	}
	return buffer
}

func FillBelow(buffer []string, width int, height int) []string {
	for len(buffer) < height {
		buffer = append(buffer, strings.Repeat(" ", width))
	}
	return buffer
}
