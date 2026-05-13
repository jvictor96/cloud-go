package transformer

import (
	"cloud/core"
)

type Static struct {
}

func (f *Static) Resize(arts []core.ArtWork) []core.ArtWork {
	return arts
}

func (f *Static) CalculateFrameCount(mapa []core.ArtWork) int {
	for i := range mapa {
		mapa[i].FrameCount = 1
	}
	return 1
}

func (f *Static) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		mapa[i].Snapshot = mapa[i].ArtWork.Content
	}
	return mapa
}
