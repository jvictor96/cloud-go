package main

import "time"

type Sleeper interface {
	Sleep(i int)
}

type Linear struct {
	Speed int
}

func (l *Linear) Sleep(i int) {
	time.Sleep(time.Duration(l.Speed) * time.Millisecond)
}

type Accelerated struct {
	Speed        int
	Acceleration int
}

func (a *Accelerated) Sleep(i int) {
	time.Sleep(time.Duration(a.Speed+a.Acceleration*i) * time.Millisecond)
}
