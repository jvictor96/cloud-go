package art

import "cloud/core"

type FillOnce struct {
}

func (f *FillOnce) PlaceArt(e *core.Engine) {
	for core.PlaceImages(e) {
	}
}
