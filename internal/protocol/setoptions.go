package protocol

import "github.com/msanft/proton/internal/protocol/verbosity"

type SetOptions struct {
	KeepFailing     bool
	KeepGoing       bool
	TryFallback     bool
	Verbosity       verbosity.Verbosity
	MaxBuildJobs    uint64
	MaxSilentTime   uint64
	useBuildHook    bool
	BuildVerbosity  verbosity.Verbosity
	logType         uint64
	printBuildTrace uint64
	BuildCores      uint64
	UseSubstitutes  bool
	Options         map[string]string
}
