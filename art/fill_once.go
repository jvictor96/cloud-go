package art

import "cloud/core"

type FillOnce struct {
}

func (f *FillOnce) PlaceArt(e *core.Engine) {
	e.LastPrint = 0
	for core.PlaceImages(e) {
	}
}
