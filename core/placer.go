package core

import (
	"math/rand"
	"unicode/utf8"
)

type Placer interface {
	PlaceArt(e *Engine)
}

func PlaceImages(e *Engine) bool {
	modified := false
	for i := range e.Galery.ArtWorks {
		art := &e.Galery.ArtWorks[i]
		height := 0
		pos := 0
		minDif := 0
		startingPoint := e.LastPrint + 1 + e.Spacing

		for cursor, line := range e.Buffer {
			lineLen := utf8.RuneCountInString(line)
			if (e.Columns-lineLen > art.Width) && (cursor >= startingPoint) {
				height++
				if minDif < lineLen {
					minDif = lineLen
				}
				if height > art.Height {
					e.LastPrint = pos + art.Height + e.Spacing
					fuzz := rand.Intn(e.Columns - minDif - art.Width - 1)
					e.Map = append(e.Map, Placing{
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
	return modified
}
