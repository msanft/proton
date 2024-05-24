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
	Size() uint64
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
	var newOperation Operation

	if err := newOperation.Opcode.UnmarshalNix(raw[:8]); err != nil {
		return fmt.Errorf("unmarshaling opcode: %w", err)
	}

	if err := newOperation.Body.UnmarshalNix(raw[8:]); err != nil {
		return fmt.Errorf("unmarshaling body: %w", err)
	}

	*o = newOperation
	return nil
}

// Size returns the size of the operation in bytes.
func (o *Operation) Size() uint64 {
	return o.Opcode.Size() + o.Body.Size()
}
