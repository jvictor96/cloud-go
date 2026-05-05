package main

import (
	"strings"
	"time"
)

type Transformer interface {
	Transform(frame int, e *Engine)
	CalculateFrameCount(e *Engine) int
	Sleep()
}

type Float struct {
}

func (f *Float) CalculateFrameCount(e *Engine) int {
	frame_count := 0
	for _, art := range e.Map {
		if art.ArtWork.Height > frame_count {
			frame_count = art.ArtWork.Height
		}
	}
	return frame_count * 2
}

func (f *Float) Sleep() {
	time.Sleep(50 * time.Millisecond)
}

func (f *Float) Transform(frame int, e *Engine) {
	for i := range e.Map {
		e.Map[i].Snapshot = []string{}
		if frame < e.Map[i].ArtWork.Height {
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, e.Map[i].ArtWork.Content[:frame]...)
			e.Map[i].Snapshot = FillAbove(e.Map[i].Snapshot, e.Map[i].ArtWork.Width, e.Map[i].ArtWork.Height)
		} else {
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, e.Map[i].ArtWork.Content[frame-e.Map[i].ArtWork.Height:]...)
			e.Map[i].Snapshot = FillBelow(e.Map[i].Snapshot, e.Map[i].ArtWork.Width, e.Map[i].ArtWork.Height)
		}
	}
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
