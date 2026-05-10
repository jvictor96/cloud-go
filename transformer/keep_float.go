package transformer

import (
	"cloud/core"
)

type KeepFloat struct {
}

func (f *KeepFloat) Resize(arts []core.ArtWork) []core.ArtWork {
	for art := range arts {
		arts[art].Content = core.FillAbove(arts[art].Content, arts[art].Width, arts[art].Height+2)
		arts[art].Height += 2
	}
	return arts
}

func (f *KeepFloat) CalculateFrameCount(mapa []core.ArtWork) int {
	for i := range mapa {
		mapa[i].FrameCount = 5_000_000
	}
	return 5_000_000
}

func (f *KeepFloat) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		v_frame := (frame + mapa[i].FirstFrame) % 4
		if v_frame == 3 {
			v_frame = 1
		}
		mapa[i].Snapshot = mapa[i].ArtWork.Content[v_frame:]
		mapa[i].Snapshot = core.FillBelow(mapa[i].Snapshot, mapa[i].ArtWork.Width, mapa[i].ArtWork.Height+2)
	}
	return mapa
}
