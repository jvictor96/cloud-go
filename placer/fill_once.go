package placer

import (
	"cloud/core"
	"math/rand"
	"unicode/utf8"
)

type FillOnce struct {
	Spacing int
	Chance  float32
}

func (placer *FillOnce) PlaceArt(artWorks []core.ArtWork, terminal core.Terminal) []core.Placing {
	terminal.LastPrint = 0
	mapa := []core.Placing{}
	modified := true
	rand.Shuffle(len(artWorks), func(i, j int) { artWorks[i], artWorks[j] = artWorks[j], artWorks[i] })
	for modified {
		modified = false
		for i := range artWorks {
			art := &artWorks[i]
			height := 0
			pos := 0
			minDif := 0
			startingPoint := terminal.LastPrint + 1 + placer.Spacing

			for cursor, line := range terminal.Buffer {
				lineLen := utf8.RuneCountInString(line)
				if (terminal.Columns-lineLen > art.Width) && (cursor >= startingPoint) {
					if rand.Float32() > placer.Chance && height == 0 {
						continue
					}
					height++
					if minDif < lineLen {
						minDif = lineLen
					}
					if height > art.Height {
						terminal.LastPrint = pos + art.Height + placer.Spacing
						fuzz := rand.Intn(terminal.Columns - minDif - art.Width - 1)
						mapa = append(mapa, core.Placing{
							ArtWork: art,
							PosY:    pos,
							Padding: minDif + fuzz,
						})
						modified = true
						break
					}
				} else {
					pos = cursor
					minDif = lineLen
					height = 0
				}
			}
		}
	}
	return mapa
}
