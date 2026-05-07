package placer

import "cloud/core"

type FillOnce struct {
	Spacing int
}

func (f *FillOnce) PlaceArt(e *core.Engine) {
	e.Terminal.LastPrint = 0
	for core.PlaceImages(e) {
	}
}

func (f *FillOnce) GetSpacing() int {
	return f.Spacing
}
