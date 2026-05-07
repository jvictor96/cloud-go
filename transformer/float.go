package transformer

import (
	"cloud/core"
)

type Float struct {
}

func (f *Float) CalculateFrameCount(a core.ArtWork) int {
	return a.Height * 2
}

func (f *Float) Transform(frame int, e *core.Engine) {
	for i := range e.Map {
		if frame > e.Map[i].FrameCount {
			continue
		}
		e.Map[i].Snapshot = []string{}
		if frame < e.Map[i].ArtWork.Height {
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, e.Map[i].ArtWork.Content[:frame]...)
			e.Map[i].Snapshot = core.FillAbove(e.Map[i].Snapshot, e.Map[i].ArtWork.Width, e.Map[i].ArtWork.Height)
		} else {
			e.Map[i].Snapshot = append(e.Map[i].Snapshot, e.Map[i].ArtWork.Content[frame-e.Map[i].ArtWork.Height:]...)
			e.Map[i].Snapshot = core.FillBelow(e.Map[i].Snapshot, e.Map[i].ArtWork.Width, e.Map[i].ArtWork.Height)
		}
	}
}
