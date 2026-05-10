package placer

import (
	"cloud/core"
	"math/rand"
	"unicode/utf8"
)

type FillOnce struct {
	Spacing int
	MaxFuzz int
	Chance  float32
}

func (placer *FillOnce) PlaceArt(artWorks []core.ArtWork, terminal core.Terminal) []core.Placing {
	if placer.MaxFuzz <= 0 {
		placer.MaxFuzz = terminal.Columns
	}
	lastPrint := 0
	mapa := []core.Placing{}
	modified := true
	//rand.Shuffle(len(artWorks), func(i, j int) { artWorks[i], artWorks[j] = artWorks[j], artWorks[i] })
	for modified {
		modified = false
		for i := range artWorks {
			art := &artWorks[i]
			height := 0
			pos := 0
			maxLineLen := 0

			for cursor, line := range terminal.Buffer {
				lineLen := utf8.RuneCountInString(line)
				if (terminal.Columns-lineLen-art.Width > 0) && (cursor >= lastPrint) {
					if height == 0 {
						if rand.Float32() > placer.Chance {
							continue
						}
						pos = cursor
						maxLineLen = lineLen
					}
					height++
					if maxLineLen < lineLen {
						maxLineLen = lineLen
					}
					if height == art.Height {
						lastPrint = pos + art.Height + placer.Spacing
						mapa = append(mapa, core.Placing{
							ArtWork: art,
							PosY:    pos,
							Padding: maxLineLen,
						})
						modified = true
						break
					}
				} else {
					height = 0
				}
			}
		}
	}
	return mapa
}
