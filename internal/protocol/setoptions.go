package protocol

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
	"github.com/msanft/proton/internal/protocol/verbosity"
	"github.com/msanft/proton/internal/pseudo"
)

// SetOptions are the options that can be set on the daemon.
type SetOptions struct {
	KeepFailing     pseudo.Bool
	KeepGoing       pseudo.Bool
	TryFallback     pseudo.Bool
	Verbosity       verbosity.Verbosity
	MaxBuildJobs    primitive.Int
	MaxSilentTime   primitive.Int
	useBuildHook    pseudo.Bool
	BuildVerbosity  verbosity.Verbosity
	logType         primitive.Int
	printBuildTrace primitive.Int
	BuildCores      primitive.Int
	UseSubstitutes  pseudo.Bool
	Options         map[pseudo.String]pseudo.String
}

// MarshalNix serializes the set-options message to the Nix wire format.
func (o SetOptions) MarshalNix() ([]byte, error) {
	var buf []byte

	keepFailing, err := o.KeepFailing.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling KeepFailing: %w", err)
	}
	buf = append(buf, keepFailing...)

	keepGoing, err := o.KeepGoing.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling KeepGoing: %w", err)
	}
	buf = append(buf, keepGoing...)

	tryFallback, err := o.TryFallback.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling TryFallback: %w", err)
	}
	buf = append(buf, tryFallback...)

	verbosity, err := o.Verbosity.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling Verbosity: %w", err)
	}
	buf = append(buf, verbosity...)

	maxBuildJobs, err := o.MaxBuildJobs.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling MaxBuildJobs: %w", err)
	}
	buf = append(buf, maxBuildJobs...)

	maxSilentTime, err := o.MaxSilentTime.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling MaxSilentTime: %w", err)
	}
	buf = append(buf, maxSilentTime...)

	useBuildHook, err := o.useBuildHook.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling useBuildHook: %w", err)
	}
	buf = append(buf, useBuildHook...)

	buildVerbosity, err := o.BuildVerbosity.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling BuildVerbosity: %w", err)
	}
	buf = append(buf, buildVerbosity...)

	logType, err := o.logType.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling logType: %w", err)
	}
	buf = append(buf, logType...)

	printBuildTrace, err := o.printBuildTrace.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling printBuildTrace: %w", err)
	}
	buf = append(buf, printBuildTrace...)

	buildCores, err := o.BuildCores.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling BuildCores: %w", err)
	}
	buf = append(buf, buildCores...)

	useSubstitutes, err := o.UseSubstitutes.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling UseSubstitutes: %w", err)
	}
	buf = append(buf, useSubstitutes...)

	for k, v := range o.Options {
		key, err := k.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling key %s: %w", k, err)
		}
		buf = append(buf, key...)

		value, err := v.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling value %s: %w", v, err)
		}
		buf = append(buf, value...)
	}

	return buf, nil
}

// UnmarshalNix deserializes the set-options message from the Nix wire format.
func (o *SetOptions) UnmarshalNix(raw []byte) error {
	var newOpts SetOptions
	if err := newOpts.KeepFailing.UnmarshalNix(raw[:8]); err != nil {
		return fmt.Errorf("unmarshaling KeepFailing: %w", err)
	}

	if err := newOpts.KeepGoing.UnmarshalNix(raw[8:16]); err != nil {
		return fmt.Errorf("unmarshaling KeepGoing: %w", err)
	}

	if err := newOpts.TryFallback.UnmarshalNix(raw[16:24]); err != nil {
		return fmt.Errorf("unmarshaling TryFallback: %w", err)
	}

	if err := newOpts.Verbosity.UnmarshalNix(raw[24:32]); err != nil {
		return fmt.Errorf("unmarshaling Verbosity: %w", err)
	}

	if err := newOpts.MaxBuildJobs.UnmarshalNix(raw[32:40]); err != nil {
		return fmt.Errorf("unmarshaling MaxBuildJobs: %w", err)
	}

	if err := newOpts.MaxSilentTime.UnmarshalNix(raw[40:48]); err != nil {
		return fmt.Errorf("unmarshaling MaxSilentTime: %w", err)
	}

	if err := newOpts.useBuildHook.UnmarshalNix(raw[48:56]); err != nil {
		return fmt.Errorf("unmarshaling useBuildHook: %w", err)
	}

	if err := newOpts.BuildVerbosity.UnmarshalNix(raw[56:64]); err != nil {
		return fmt.Errorf("unmarshaling BuildVerbosity: %w", err)
	}

	if err := newOpts.logType.UnmarshalNix(raw[64:72]); err != nil {
		return fmt.Errorf("unmarshaling logType: %w", err)
	}

	if err := newOpts.printBuildTrace.UnmarshalNix(raw[72:80]); err != nil {
		return fmt.Errorf("unmarshaling printBuildTrace: %w", err)
	}

	if err := newOpts.BuildCores.UnmarshalNix(raw[80:88]); err != nil {
		return fmt.Errorf("unmarshaling BuildCores: %w", err)
	}

	if err := newOpts.UseSubstitutes.UnmarshalNix(raw[88:96]); err != nil {
		return fmt.Errorf("unmarshaling UseSubstitutes: %w", err)
	}

	newOpts.Options = make(map[pseudo.String]pseudo.String)
	i := 96
	for i < len(raw) {
		var key pseudo.String
		if err := key.UnmarshalNix(raw[i:]); err != nil {
			return fmt.Errorf("unmarshaling key: %w", err)
		}
		i += 8 + len(key)

		var value pseudo.String
		if err := value.UnmarshalNix(raw[i:]); err != nil {
			return fmt.Errorf("unmarshaling value: %w", err)
		}
		i += 8 + len(value)

		newOpts.Options[key] = value
	}

	*o = newOpts
	return nil
}
