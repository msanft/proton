package gcaction

type GCAction uint64

const (
	ReturnLive GCAction = iota
	ReturnDead
	DeleteDead
	DeleteSpecific
)
