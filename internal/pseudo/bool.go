package pseudo

import (
	"github.com/msanft/proton/internal/primitive"
)

// Bool is a pseudo-type for a boolean value used by Nix. Under the
// hood, it is a integer primitive.
type Bool bool

// MarshalNix serializes the pseudo-boolean to the Nix wire format.
func (b Bool) MarshalNix() ([]byte, error) {
	var i uint64
	if b {
		i = 1
	}
	return primitive.NewInt(i).MarshalNix()
}

// UnmarshalNix deserializes the pseudo-boolean from the Nix wire format.
// A value of 0 is considered false, all other values are considered true.
func (b *Bool) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}
	*b = i.Value != 0
	return nil
}
