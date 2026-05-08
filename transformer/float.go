package transformer

import (
	"cloud/core"
)

type Float struct {
}

func (f *Float) CalculateFrameCount(mapa []core.ArtWork) int {
	fc := 0
	for i := range mapa {
		art_fc := mapa[i].Height * 2
		mapa[i].FrameCount = art_fc
		fc = max(art_fc, fc)
	}
	return fc
}

func (f *Float) Transform(frame int, mapa []core.Placing) []core.Placing {
	for i := range mapa {
		if frame > mapa[i].ArtWork.FrameCount+mapa[i].FirstFrame || frame < mapa[i].FirstFrame {
			mapa[i].Snapshot = core.FillAbove([]string{}, mapa[i].ArtWork.Width, mapa[i].ArtWork.Height)
			continue
		}
		v_frame := frame - mapa[i].FirstFrame
		if v_frame < mapa[i].ArtWork.Height {
			mapa[i].Snapshot = core.FillAbove(mapa[i].ArtWork.Content[:v_frame], mapa[i].ArtWork.Width, mapa[i].ArtWork.Height)
		} else {
			mapa[i].Snapshot = core.FillBelow(mapa[i].ArtWork.Content[v_frame-mapa[i].ArtWork.Height:], mapa[i].ArtWork.Width, mapa[i].ArtWork.Height)
		}
	}
	return mapa
}
