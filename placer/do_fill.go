package placer

import (
	"cloud/core"
	"cloud/transformer"
	"math/rand"
)

type DoFill struct {
	FillOnce FillOnce
}

func (placer *DoFill) PlaceArt(artWorks []core.ArtWork, terminal core.Terminal) []core.Placing {
	transformer_impl := &transformer.Static{}
	placing := placer.FillOnce.PlaceArt(artWorks, terminal)
	terminal.Buffer = core.ManipulateBuffer(0, placing, transformer_impl, terminal)
	step := placer.FillOnce.PlaceArt(artWorks, terminal)
	placing = append(placing, step...)
	for len(step) > 0 {
		terminal.Buffer = core.ManipulateBuffer(0, placing, transformer_impl, terminal)
		step = placer.FillOnce.PlaceArt(artWorks, terminal)
		placing = append(placing, step...)
	}
	for placemnt := range placing {
		placing[placemnt].FirstFrame = rand.Int() % (placing[placemnt].ArtWork.FrameCount)
	}
	return placing
}
