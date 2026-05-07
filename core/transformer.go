package core

import (
	"strings"
)

type Transformer interface {
	Transform(frame int, e *Engine)
	CalculateFrameCount(art_work ArtWork) int
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
