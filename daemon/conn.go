package daemon

import (
	"errors"
	"fmt"
	"io"

	"github.com/msanft/proton/internal/primitive"
	"github.com/msanft/proton/internal/protocol"
	"github.com/msanft/proton/internal/pseudo"
)

// Conn represents a connection to a Nix daemon.
type Conn struct {
	r io.Reader
	w io.Writer

	stderr io.Writer

	// PeerVersion is the Nix version used by the other end of the connection.
	PeerVersion string
}

// NewConn establishes a connection to a Nix daemon through the given connection.
func NewConn(r io.Reader, w io.Writer, stderr io.Writer) (*Conn, error) {
	c := &Conn{r: r, w: w, stderr: stderr}
	err := c.shakeHands()
	if err != nil {
		return nil, fmt.Errorf("performing handshake: %w", err)
	}
	return c, nil
}

// nixMarshalable is an interface for types that can be marshaled into
// the Nix wire format.
type nixMarshalable interface {
	MarshalNix() ([]byte, error)
}

// writeNix marshals the given data to the Nix wire format
// and writes it to the connection.
func (c *Conn) writeNix(data nixMarshalable) error {
	b, err := data.MarshalNix()
	if err != nil {
		return fmt.Errorf("marshaling data: %w", err)
	}
	// fmt.Printf("Write %x\n", b)
	if _, err = c.w.Write(b); err != nil {
		return fmt.Errorf("writing data: %w", err)
	}
	return nil
}

// nixUnmarshalable is an interface for types that can be unmarshaled from
// the Nix wire format.
type nixUnmarshalable interface {
	UnmarshalNix([]byte) error
}

// ReadNix reads n bytes from the connection and unmarshals the received data
// into the given data structure.
// Callers should rather use the type-specific implementations like ReadNixInt.
func (c *Conn) readNNix(data nixUnmarshalable, n int) error {
	buf := make([]byte, n)
	_, err := c.r.Read(buf)
	if err != nil {
		return fmt.Errorf("reading data: %w", err)
	}
	//fmt.Printf("Read %x\n", buf)
	if err = data.UnmarshalNix(buf); err != nil {
		return fmt.Errorf("unmarshaling data: %w", err)
	}
	return nil
}

// readNixInt reads the nix wire format encoded integer from the connection.
func (c *Conn) readNixInt() (primitive.Int, error) {
	var i primitive.Int
	if err := c.readNNix(&i, 8); err != nil {
		return primitive.Int{}, fmt.Errorf("reading integer: %w", err)
	}
	return i, nil
}

// readNixByteBuf reads a byte buffer from the connection.
func (c *Conn) readNixByteBuf() (primitive.ByteBuf, error) {
	l, err := c.readNixInt()
	if err != nil {
		return primitive.ByteBuf{}, fmt.Errorf("reading buffer length: %w", err)
	}

	// Read the actual buffer
	buf := make([]byte, l.Value)
	if _, err = c.r.Read(buf); err != nil {
		return primitive.ByteBuf{}, fmt.Errorf("reading buffer: %w", err)
	}

	// Discard the padding
	pad := 8 - l.Value%8
	if pad > 0 {
		if _, err := io.CopyN(io.Discard, c.r, int64(pad)); err != nil {
			return primitive.ByteBuf{}, fmt.Errorf("discarding padding: %w", err)
		}
	}

	return primitive.ByteBuf{Len: l, Buf: buf}, nil
}

// readNixBool reads a boolean from the connection.
func (c *Conn) readNixBool() (pseudo.Bool, error) {
	var b pseudo.Bool
	if err := c.readNNix(&b, 8); err != nil {
		return false, fmt.Errorf("reading boolean: %w", err)
	}
	return b, nil
}

// shakeHands performs a handshake with the Nix daemon to
// establish a connection.
//
// See https://github.com/NixOS/nix/blob/7cb3c80bb53f4ddf602d41eb708b3a32c04f2620/src/libstore/daemon.cc#L1028.
func (c *Conn) shakeHands() error {
	// Send our own magic number
	magic := primitive.NewInt(protocol.ClientMagic)
	if err := c.writeNix(magic); err != nil {
		return fmt.Errorf("writing client magic: %w", err)
	}

	// Read the server's magic number
	magicResp, err := c.readNixInt()
	if err != nil {
		return fmt.Errorf("reading server magic: %w", err)
	}
	if magicResp.Value != protocol.ServerMagic {
		return fmt.Errorf("unexpected server magic: got %x, want %x", magicResp.Value, protocol.ServerMagic)
	}

	// Read the server's protocol version
	var serverVersion protocol.Version
	if err := c.readNNix(&serverVersion, 8); err != nil {
		return fmt.Errorf("reading server protocol version: %w", err)
	}
	// Check whether the server's protocol version is compatible with ours.
	// TODO: Make this more flexible.
	clientVersion := protocol.OwnVersion()
	if serverVersion != clientVersion {
		return fmt.Errorf(
			"protocol version mismatch: server is %s, client is %s",
			serverVersion, clientVersion)
	}

	// Send our protocol version
	if err := c.writeNix(clientVersion); err != nil {
		return fmt.Errorf("writing client protocol version: %w", err)
	}

	// Write the now obsolete CPU Affinity..
	if err := c.writeNix(primitive.NewInt(0)); err != nil {
		return fmt.Errorf("writing CPU affinity: %w", err)
	}
	// ..and reserved field
	if err := c.writeNix(primitive.NewInt(0)); err != nil {
		return fmt.Errorf("writing reserved field: %w", err)
	}

	// Read the server's Nix version
	v, err := c.readNixByteBuf()
	if err != nil {
		return fmt.Errorf("reading server Nix version: %w", err)
	}
	c.PeerVersion = string(v.Buf)

	// Read whether we're trusted
	trusted, err := c.readNixInt()
	if err != nil {
		return fmt.Errorf("reading trusted status: %w", err)
	}
	switch trusted.Value {
	case 0: // Unset
	case 1: // Trusted
	case 2: // Untrusted
		return errors.New("connection is untrusted")
	default:
		return fmt.Errorf("unexpected trusted status: %d", trusted.Value)
	}

	// TODO: Properly read stderr
	stderrMarker, err := c.readNixInt()
	if err != nil {
		return fmt.Errorf("reading stderr marker: %w", err)
	}
	if stderrMarker.Value != 0x616c7473 {
		return fmt.Errorf("unexpected stderr marker: got %x, want 0x616c7473", stderrMarker.Value)
	}

	return nil
}
