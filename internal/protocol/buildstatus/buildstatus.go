package buildstatus

type BuildStatus uint64

const (
	Built BuildStatus = iota
	Substituted
	AlreadyValid
	PermanentFailure
	InputRejected
	OutputRejected
	TransientFailure
	CachedFailure
	TimedOut
	MiscFailure
	DependencyFailed
	LogLimitExceeded
	NotDeterministic
	ResolvesToAlreadyValid
	NoSubstituters
)
