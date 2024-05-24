package pseudo

import "github.com/msanft/proton/internal/primitive"

// String is a pseudo-type for a string value used by Nix. Under the
// hood, it is a byte buffer primitive.
type String string

// MarshalNix serializes the pseudo-string to the Nix wire format.
func (s String) MarshalNix() ([]byte, error) {
	return primitive.NewByteBuf([]byte(s)).MarshalNix()
}

// UnmarshalNix deserializes the pseudo-string from the Nix wire format.
func (s *String) UnmarshalNix(raw []byte) error {
	var bbuf primitive.ByteBuf
	if err := bbuf.UnmarshalNix(raw); err != nil {
		return err
	}
	*s = String(bbuf.Buf)
	return nil
}

// Size returns the size of the pseudo-string in bytes.
func (s String) Size() uint64 {
	return primitive.NewByteBuf([]byte(s)).Size()
}
