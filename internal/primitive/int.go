package primitive

import (
	"encoding/binary"
	"errors"
)

// Int is the unsigned 64-bit integer primitive type.
type Int struct {
	Value uint64
}

// NewInt creates a new Int.
func NewInt(v uint64) Int {
	return Int{Value: v}
}

// MarshalNix serializes the integer to the Nix wire format.
func (i Int) MarshalNix() ([]byte, error) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i.Value)
	return b, nil
}

// UnmarshalNix deserializes the integer from the Nix wire format.
// If a buffer with more than 8 bytes is given, only the first 8 bytes are read.
func (i *Int) UnmarshalNix(raw []byte) error {
	if len(raw) < 8 {
		return errors.New("buffer is smaller than 8 bytes")
	}
	i.Value = binary.LittleEndian.Uint64(raw[:8])
	return nil
}

// Size returns the size of the integer in bytes.
func (i Int) Size() uint64 {
	return 8
}
