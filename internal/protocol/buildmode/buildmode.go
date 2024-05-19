package buildmode

type BuildMode uint64

const (
	Normal BuildMode = iota
	Repair
	Check
)
