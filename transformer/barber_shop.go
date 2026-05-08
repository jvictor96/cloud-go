package transformer

import (
	"cloud/core"
	"strings"
)

type BarberShop struct {
}

func (f *BarberShop) CalculateFrameCount(mapa []core.Placing) int {
	return 5_000_000
}

func (f *BarberShop) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		for _, line := range mapa[i].ArtWork.Content {
			runes := []rune(line)
			width := len(runes)
			padding := strings.Repeat(" ", 3)
			runes = append(runes, []rune(padding)...)
			frame = frame % len(runes)
			runes = append(runes[frame:], runes[:frame]...)
			runes = runes[:width]
			mapa[i].Snapshot = append(mapa[i].Snapshot, string(runes))
		}
	}
	return mapa
}
