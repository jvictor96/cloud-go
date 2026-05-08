package core

import (
	"strings"
)

type Transformer interface {
	Transform(frame int, art_work []Placing) []Placing
	CalculateFrameCount(art_work []Placing) int
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
