package core

type Sleeper interface {
	Sleep(frame int)
	SetDuration(t int)
}
