package transformer

import (
	"cloud/core"
	"strings"
)

type BarberShop struct {
}

func (f *BarberShop) CalculateFrameCount(e *core.Engine) int {
	return 5_000_000
}

func (f *BarberShop) Transform(frame int, e *core.Engine) {
	for i := range e.Map {
		e.Map[i].Snapshot = []string{}
		for _, line := range e.Map[i].ArtWork.Content {
			runes := []rune(line)
			width := len(runes)
			padding := strings.Repeat(" ", 3)
			runes = append(runes, []rune(padding)...)
			frame = frame % len(runes)
			runes = append(runes[frame:], runes[:frame]...)
			runes = runes[:width]
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, string(runes))
		}
	}
}
