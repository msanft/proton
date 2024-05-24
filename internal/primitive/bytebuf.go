package primitive

import (
	"fmt"
)

// ByteBuf is the primitive byte buffer type.
type ByteBuf struct {
	Len Int
	Buf []byte
}

/*
NewByteBuf creates a new ByteBuf.

It will pad the byte buffer with zeros if the length is not a multiple of 8.
*/
func NewByteBuf(buf []byte) ByteBuf {
	lenWithoutPadding := len(buf)
	if lenWithoutPadding%8 != 0 {
		buf = append(buf, make([]byte, 8-lenWithoutPadding%8)...)
	}
	return ByteBuf{Len: NewInt(uint64(lenWithoutPadding)), Buf: buf}
}

// MarshalNix serializes the byte buffer to the Nix wire format.
func (b ByteBuf) MarshalNix() ([]byte, error) {
	l, err := b.Len.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling buffer length: %w", err)
	}
	return append(l, b.Buf...), nil
}

// UnmarshalNix deserializes the byte buffer from the Nix wire format.
func (b *ByteBuf) UnmarshalNix(raw []byte) error {
	// If less than 8 bytes are given, we can't even read the length.
	if len(raw) < 8 {
		return fmt.Errorf("buffer length is less than 8 bytes")
	}

	if err := b.Len.UnmarshalNix(raw[:8]); err != nil {
		return fmt.Errorf("unmarshaling buffer length: %w", err)
	}
	b.Buf = raw[8 : 8+b.Len.Value]
	return nil
}

// Size returns the size of the byte buffer in bytes.
func (b ByteBuf) Size() uint64 {
	paddedLen := uint64(len(b.Buf))
	if paddedLen%8 != 0 {
		paddedLen += 8 - paddedLen%8
	}
	return 8 + paddedLen
}
