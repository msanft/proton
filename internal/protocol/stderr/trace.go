package stderr

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

// Trace is a trace message for backtracking errors.
type Trace struct {
	Position primitive.Int
	Message  primitive.ByteBuf
}

// MarshalNix serializes the trace to the Nix wire format.
func (t Trace) MarshalNix() ([]byte, error) {
	pos, err := t.Position.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling position: %w", err)
	}

	msg, err := t.Message.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling message: %w", err)
	}

	return append(pos, msg...), nil
}

// UnmarshalNix deserializes the trace from the Nix wire format.
func (t *Trace) UnmarshalNix(raw []byte) error {
	var newTrace Trace

	if err := newTrace.Position.UnmarshalNix(raw[:8]); err != nil {
		return fmt.Errorf("unmarshaling position: %w", err)
	}

	if err := newTrace.Message.UnmarshalNix(raw[8:]); err != nil {
		return fmt.Errorf("unmarshaling message: %w", err)
	}

	*t = newTrace
	return nil
}

// Size returns the size of the trace in bytes.
func (t Trace) Size() uint64 {
	return t.Position.Size() + t.Message.Size()
}
