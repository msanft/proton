package daemon

import (
	"errors"
	"fmt"
	"io"

	"github.com/msanft/proton/internal/primitive"
	"github.com/msanft/proton/internal/protocol/stderr"
	"github.com/msanft/proton/internal/pseudo"
)

// nixWire is an interface for types that can be (un)marshaled from/into
// the Nix wire format.
type nixWire interface {
	MarshalNix() ([]byte, error)
	UnmarshalNix([]byte) error
	Size() uint64
}

// ReceiveStderr tries to read all leftover error messages from the daemon, if any.
// Error messages are written into the writer.
func (c *Conn) ParseStderr() (rest []byte, retErr error) {
	data, err := c.readUntilLastStderr()
	if err != nil {
		retErr = fmt.Errorf("reading stderr: %w", err)
		return
	}

	for {
		var marker stderr.Marker
		if err := marker.UnmarshalNix(data); err != nil {
			retErr = errors.Join(retErr, fmt.Errorf("unmarshaling stderr marker: %w", err))
			return
		}
		data = data[marker.Size():]

		var body nixWire
		switch marker {
		case stderr.MarkerLast:
			// Parsing is done
			return
		case stderr.MarkerWrite, stderr.MarkerNext:
			body = new(pseudo.String)
		case stderr.MarkerError:
			body = &stderr.Error{}
		case stderr.MarkerResult:
			body = &stderr.Result{}
		case stderr.MarkerStartActivity:
			body = &stderr.StartActivity{}
		case stderr.MarkerStopActivity:
			body = &primitive.Int{}
		default:
			// we read into data
			rest, err = marker.MarshalNix()
			if err != nil {
				retErr = errors.Join(retErr, fmt.Errorf("marshaling stderr marker: %w", err))
			}
			return
		}

		if err := body.UnmarshalNix(data); err != nil {
			retErr = errors.Join(retErr, fmt.Errorf("unmarshaling stderr body: %w", err))
			return
		}
		data = data[body.Size():]

		// TODO: Actual error message
		c.stderr.Write([]byte(fmt.Sprintf("stderr: %q\n", body)))
	}
}

// readUntilLastStderr reads all data from the reader until the last stderr marker.
func (c *Conn) readUntilLastStderr() ([]byte, error) {
	var buf []byte
	var rawMarker [8]byte
	for {
		_, readErr := c.r.Read(rawMarker[:])
		if readErr != nil && !errors.Is(readErr, io.EOF) {
			return nil, fmt.Errorf("reading marker: %w", readErr)
		}
		buf = append(buf, rawMarker[:]...)

		var marker stderr.Marker
		if err := marker.UnmarshalNix(rawMarker[:]); err != nil {
			return nil, fmt.Errorf("unmarshaling marker: %w", err)
		}

		if marker == stderr.MarkerLast || errors.Is(readErr, io.EOF) {
			break
		}
	}

	return buf, nil
}
