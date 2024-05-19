package primitives

import "encoding/binary"

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
func (i *Int) UnmarshalNix(raw []byte) error {
	i.Value = binary.LittleEndian.Uint64(raw)
	return nil
}
