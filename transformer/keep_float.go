package transformer

import (
	"cloud/core"
)

type KeepFloat struct {
	Amplitude int
}

func (f *KeepFloat) Resize(arts []core.ArtWork) []core.ArtWork {
	for art := range arts {
		arts[art].Content = core.FillAbove(arts[art].Content, arts[art].Width, arts[art].Height+f.Amplitude-1)
		arts[art].Height += f.Amplitude - 1
	}
	return arts
}

func (f *KeepFloat) CalculateFrameCount(mapa []core.ArtWork) {
	for i := range mapa {
		mapa[i].FrameCount = 5_000_000
	}
}

func (f *KeepFloat) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		v_frame := (frame + mapa[i].FirstFrame) % (f.Amplitude*2 - 2)
		if v_frame >= f.Amplitude {
			v_frame = f.Amplitude - 2 - v_frame%f.Amplitude
		}
		mapa[i].Snapshot = mapa[i].ArtWork.Content[v_frame:]
		mapa[i].Snapshot = core.FillBelow(mapa[i].Snapshot, mapa[i].ArtWork.Width, mapa[i].ArtWork.Height+f.Amplitude-1)
	}
	return mapa
}
