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

	if err := newOpts.KeepFailing.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling keepFailing: %w", err)
	}
	raw = raw[newOpts.KeepFailing.Size():]

	if err := newOpts.KeepGoing.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling keepGoing: %w", err)
	}
	raw = raw[newOpts.KeepGoing.Size():]

	if err := newOpts.TryFallback.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling tryFallback: %w", err)
	}
	raw = raw[newOpts.TryFallback.Size():]

	if err := newOpts.Verbosity.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling verbosity: %w", err)
	}
	raw = raw[newOpts.Verbosity.Size():]

	if err := newOpts.MaxBuildJobs.UnmarshalNix(raw); err != nil {

		return fmt.Errorf("unmarshaling maxBuildJobs: %w", err)
	}
	raw = raw[newOpts.MaxBuildJobs.Size():]

	if err := newOpts.MaxSilentTime.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling maxSilentTime: %w", err)
	}
	raw = raw[newOpts.MaxSilentTime.Size():]

	if err := newOpts.useBuildHook.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling useBuildHook: %w", err)
	}
	raw = raw[newOpts.useBuildHook.Size():]

	if err := newOpts.BuildVerbosity.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling buildVerbosity: %w", err)
	}
	raw = raw[newOpts.BuildVerbosity.Size():]

	if err := newOpts.logType.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling logType: %w", err)
	}
	raw = raw[newOpts.logType.Size():]

	if err := newOpts.printBuildTrace.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling printBuildTrace: %w", err)
	}
	raw = raw[newOpts.printBuildTrace.Size():]

	if err := newOpts.BuildCores.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling buildCores: %w", err)
	}
	raw = raw[newOpts.BuildCores.Size():]

	if err := newOpts.UseSubstitutes.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling useSubstitutes: %w", err)
	}
	raw = raw[newOpts.UseSubstitutes.Size():]

	newOpts.Options = make(map[pseudo.String]pseudo.String)

	for len(raw) > 0 {
		var key pseudo.String
		if err := key.UnmarshalNix(raw); err != nil {
			return fmt.Errorf("unmarshaling key: %w", err)
		}
		raw = raw[key.Size():]

		var value pseudo.String
		if err := value.UnmarshalNix(raw); err != nil {
			return fmt.Errorf("unmarshaling value: %w", err)
		}
		raw = raw[value.Size():]

		newOpts.Options[key] = value
	}

	*o = newOpts
	return nil
}

// Size returns the size of the set-options message in bytes.
func (o *SetOptions) Size() uint64 {
	size := o.KeepFailing.Size() +
		o.KeepGoing.Size() +
		o.TryFallback.Size() +
		o.Verbosity.Size() +
		o.MaxBuildJobs.Size() +
		o.MaxSilentTime.Size() +
		o.useBuildHook.Size() +
		o.BuildVerbosity.Size() +
		o.logType.Size() +
		o.printBuildTrace.Size() +
		o.BuildCores.Size() +
		o.UseSubstitutes.Size()

	for k, v := range o.Options {
		size += k.Size() + v.Size()
	}

	return size
}
