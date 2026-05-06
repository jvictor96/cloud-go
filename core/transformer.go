package core

import (
	"strings"
)

type Transformer interface {
	Transform(frame int, e *Engine)
	CalculateFrameCount(e *Engine) int
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
