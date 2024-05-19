package pseudo

import "github.com/msanft/proton/internal/primitives"

// String is a pseudo-type for a string value used by Nix. Under the
// hood, it is a byte buffer primitive.
type String string

// MarshalNix serializes the pseudo-string to the Nix wire format.
func (s String) MarshalNix() ([]byte, error) {
	return primitives.NewByteBuf([]byte(s)).MarshalNix()
}

// UnmarshalNix deserializes the pseudo-string from the Nix wire format.
func (s *String) UnmarshalNix(raw []byte) error {
	var bbuf primitives.ByteBuf
	if err := bbuf.UnmarshalNix(raw); err != nil {
		return err
	}
	*s = String(bbuf.Buf)
	return nil
}
