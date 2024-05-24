package stderr

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

// Error is the format used to report errors over the stderr format.
type Error struct {
	Kind        primitive.ByteBuf
	Level       primitive.Int
	Name        primitive.ByteBuf
	Message     primitive.ByteBuf
	HasPosition primitive.Int
	Traces      []Trace
}

// MarshalNix serializes the error to the Nix wire format.
func (e Error) MarshalNix() ([]byte, error) {
	var buf []byte

	kind, err := e.Kind.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling kind: %w", err)
	}
	buf = append(buf, kind...)

	level, err := e.Level.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling level: %w", err)
	}
	buf = append(buf, level...)

	name, err := e.Name.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling name: %w", err)
	}
	buf = append(buf, name...)

	msg, err := e.Message.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling message: %w", err)
	}
	buf = append(buf, msg...)

	hasPos, err := e.HasPosition.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling hasPosition: %w", err)
	}
	buf = append(buf, hasPos...)

	var traces []byte
	for _, trace := range e.Traces {
		t, err := trace.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling trace: %w", err)
		}
		traces = append(traces, t...)
	}
	tracesLen, err := primitive.NewInt(uint64(len(e.Traces))).MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling traces length: %w", err)
	}
	buf = append(buf, tracesLen...)
	buf = append(buf, traces...)

	return buf, nil
}

// UnmarshalNix deserializes the error from the Nix wire format.
func (e *Error) UnmarshalNix(raw []byte) error {
	var newError Error

	if err := newError.Kind.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling kind: %w", err)
	}
	raw = raw[newError.Kind.Size():]

	if err := newError.Level.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling level: %w", err)
	}
	raw = raw[newError.Level.Size():]

	if err := newError.Name.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling name: %w", err)
	}
	raw = raw[newError.Name.Size():]

	if err := newError.Message.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling message: %w", err)
	}
	raw = raw[newError.Message.Size():]

	if err := newError.HasPosition.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling hasPosition: %w", err)
	}
	raw = raw[newError.HasPosition.Size():]

	var lenTraces primitive.Int
	if err := lenTraces.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling lenTraces: %w", err)
	}
	var traces []Trace
	for i := uint64(0); i < lenTraces.Value; i++ {
		var trace Trace
		if err := trace.UnmarshalNix(raw); err != nil {
			return fmt.Errorf("unmarshaling trace: %w", err)
		}
		traces = append(traces, trace)
		raw = raw[trace.Size():]
	}
	newError.Traces = traces

	*e = newError
	return nil
}

// Size returns the size of the error in bytes.
func (e Error) Size() uint64 {
	size := e.Kind.Size() + e.Level.Size() + e.Name.Size() + e.Message.Size() + e.HasPosition.Size()
	for _, trace := range e.Traces {
		size += trace.Size()
	}
	return size
}
