package placer

import "cloud/core"

type FillOnce struct {
}

func (f *FillOnce) PlaceArt(e *core.Engine) int {
	e.LastPrint = 0
	for core.PlaceImages(e) {
	}
	max_frame_count := 0
	for art := range e.Map {
		frame_count := e.Transformer.CalculateFrameCount(*e.Map[art].ArtWork)
		e.Map[art].FrameCount = frame_count
		max_frame_count = max(frame_count, max_frame_count)
	}
	return max_frame_count
}
