package ticker

import (
	"time"
)

type Linear struct {
	Speed int
}

func (l *Linear) Sleep(i int) {
	time.Sleep(time.Duration(l.Speed) * time.Millisecond)
}

func (l *Linear) SetDuration(i int) {
}

type Accelerated struct {
	MinSpeed float32
	MaxSpeed float32
	Duration int
}

func (a *Accelerated) Sleep(frame int) {
	Diff := a.MaxSpeed - a.MinSpeed
	Speed := float32(a.Duration/2.0-frame) / float32(a.Duration/2.0)
	RawSpeed := a.MinSpeed + Diff*max(Speed, -Speed)
	time.Sleep(time.Duration(RawSpeed) * time.Millisecond)
}

func (a *Accelerated) SetDuration(d int) {
	a.Duration = d
}

type InvertedAccelerated struct {
	MinSpeed float32
	MaxSpeed float32
	Duration int
}

func (a *InvertedAccelerated) Sleep(frame int) {
	Diff := a.MinSpeed - a.MaxSpeed
	Speed := float32(a.Duration/2.0-frame) / float32(a.Duration/2.0)
	RawSpeed := a.MaxSpeed + Diff*max(Speed, -Speed)
	time.Sleep(time.Duration(RawSpeed) * time.Millisecond)
}

func (i *InvertedAccelerated) SetDuration(d int) {
	i.Duration = d
}
