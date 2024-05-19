package daemon

import (
	"fmt"
	"io"

	"github.com/msanft/proton/internal/primitives"
	"github.com/msanft/proton/internal/protocol"
)

// Conn represents a connection to a Nix daemon.
type Conn struct {
	r io.Reader
	w io.Writer
}

// NewConn establishes a connection to a Nix daemon through the given connection.
func NewConn(r io.Reader, w io.Writer) (*Conn, error) {
	c := &Conn{r, w}
	err := c.shakeHands()
	if err != nil {
		return nil, fmt.Errorf("performing handshake: %w", err)
	}
	return c, nil
}

// NixMarshalable is an interface for types that can be marshaled into
// the Nix wire format.
type NixMarshalable interface {
	MarshalNix() ([]byte, error)
}

// WriteNix marshals the given data to the Nix wire format
// and writes it to the connection.
func (c *Conn) WriteNix(data NixMarshalable) error {
	b, err := data.MarshalNix()
	if err != nil {
		return fmt.Errorf("marshaling data: %w", err)
	}
	if _, err = c.w.Write(b); err != nil {
		return fmt.Errorf("writing data: %w", err)
	}
	return nil
}

// NixUnmarshalable is an interface for types that can be unmarshaled from
// the Nix wire format.
type NixUnmarshalable interface {
	UnmarshalNix([]byte) error
}

// ReadNix reads n bytes from the connection and unmarshals the received data
// into the given data structure.
// Callers should rather use the type-specific implementations like ReadNixInt.
func (c *Conn) ReadNNix(data NixUnmarshalable, n int) error {
	buf := make([]byte, n)
	_, err := c.r.Read(buf)
	if err != nil {
		return fmt.Errorf("reading data: %w", err)
	}
	if err = data.UnmarshalNix(buf); err != nil {
		return fmt.Errorf("unmarshaling data: %w", err)
	}
	return nil
}

// ReadNixInt reads the nix wire format encoded integer from the connection.
func (c *Conn) ReadNixInt() (primitives.Int, error) {
	var i primitives.Int
	if err := c.ReadNNix(&i, 8); err != nil {
		return primitives.Int{}, fmt.Errorf("reading integer: %w", err)
	}
	return i, nil
}

// ReadNixByteBuf reads a byte buffer from the connection.
func (c *Conn) ReadNixByteBuf() (primitives.ByteBuf, error) {
	l, err := c.ReadNixInt()
	if err != nil {
		return primitives.ByteBuf{}, fmt.Errorf("reading buffer length: %w", err)
	}
	buf := make([]byte, l.Value)
	if _, err = c.r.Read(buf); err != nil {
		return primitives.ByteBuf{}, fmt.Errorf("reading buffer: %w", err)
	}
	return primitives.ByteBuf{Len: l, Buf: buf}, nil
}

// shakeHands performs a handshake with the Nix daemon to
// establish a connection.
func (c *Conn) shakeHands() error {
	// Send our own magic number
	magic := primitives.NewInt(protocol.ClientMagic)
	if err := c.WriteNix(magic); err != nil {
		return fmt.Errorf("writing client magic: %w", err)
	}

	// Read the server's magic number
	magicResp, err := c.ReadNixInt()
	if err != nil {
		return fmt.Errorf("reading server magic: %w", err)
	}
	if magicResp.Value != protocol.ServerMagic {
		return fmt.Errorf("unexpected server magic: got %x, want %x", magicResp.Value, protocol.ServerMagic)
	}

	// Read the server's protocol version
	var serverVersion protocol.Version
	if err := c.ReadNNix(&serverVersion, 8); err != nil {
		return fmt.Errorf("reading server protocol version: %w", err)
	}

	// Send our protocol version
	if err := c.WriteNix(protocol.OwnVersion()); err != nil {
		return fmt.Errorf("writing client protocol version: %w", err)
	}

	// Write the now obsolete CPU Affinity..
	if err := c.WriteNix(primitives.NewInt(0)); err != nil {
		return fmt.Errorf("writing CPU affinity: %w", err)
	}
	// ..and reserved field
	if err := c.WriteNix(primitives.NewInt(0)); err != nil {
		return fmt.Errorf("writing reserved field: %w", err)
	}

	return nil
}
