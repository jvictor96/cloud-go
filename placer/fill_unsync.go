package placer

import (
	"cloud/core"
	"math/rand"
)

type FillUnsync struct {
	FillOnce FillOnce
}

func (placer *FillUnsync) PlaceArt(artWorks []core.ArtWork, terminal core.Terminal) []core.Placing {
	placing := placer.FillOnce.PlaceArt(artWorks, terminal)
	for placemnt := range placing {
		placing[placemnt].FirstFrame = rand.Int() % (placing[placemnt].ArtWork.FrameCount / 4)
	}
	return placing
}
