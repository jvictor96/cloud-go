package transformer

import (
	"cloud/core"
)

type Vanisher struct {
	Period int
}

func (f *Vanisher) Resize(arts []core.ArtWork) []core.ArtWork {
	return arts
}

func (f *Vanisher) CalculateFrameCount(mapa []core.ArtWork) {
	for i := range mapa {
		mapa[i].FrameCount = f.Period * 2
	}
}

func (f *Vanisher) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		mapa[i].Snapshot = mapa[i].ArtWork.Content
	}
	return mapa
}
