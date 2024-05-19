package verbosity

type Verbosity uint64

const (
	Error Verbosity = iota
	Warn
	Notice
	Info
	Talkative
	Chatty
	Debug
	Vomit
)
