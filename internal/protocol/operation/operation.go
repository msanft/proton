package operation

import (
	"fmt"

	"github.com/msanft/proton/internal/protocol/opcode"
)

// nixWire is an interface for types that can be (un)marshaled from/into
// the Nix wire format.
type nixWire interface {
	MarshalNix() ([]byte, error)
	UnmarshalNix([]byte) error
}

// Operation represents an operation on the Nix daemon.
type Operation struct {
	Opcode opcode.Opcode
	Body   nixWire
}

// NewOperation creates a new operation with the given opcode and body.
func NewOperation(opcode opcode.Opcode, body nixWire) *Operation {
	return &Operation{opcode, body}
}

// MarshalNix serializes the operation to the Nix wire format.
func (o Operation) MarshalNix() ([]byte, error) {
	opcode, err := o.Opcode.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling opcode: %w", err)
	}

	body, err := o.Body.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling body: %w", err)
	}

	return append(opcode, body...), nil
}

// UnmarshalNix deserializes the operation from the Nix wire format.
func (o *Operation) UnmarshalNix(raw []byte) error {
	if err := o.Opcode.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling opcode: %w", err)
	}

	if err := o.Body.UnmarshalNix(raw[8:]); err != nil {
		return fmt.Errorf("unmarshaling body: %w", err)
	}

	return nil
}
