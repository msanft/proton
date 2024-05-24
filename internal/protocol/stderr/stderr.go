// The stderr package implements the format the Nix daemon protocol uses to report
// error messages back to the clients.
package stderr

import "fmt"

// nixWire is an interface for types that can be (un)marshaled from/into
// the Nix wire format.
type nixWire interface {
	MarshalNix() ([]byte, error)
	UnmarshalNix([]byte) error
	Size() uint64
}

// Stderr is the error format that the Nix daemon sends back to the client.
type Stderr struct {
	Marker Marker
	Body   nixWire
}

// NewStderr creates a new stderr message with the given marker and body.
func NewStderr(marker Marker, body nixWire) *Stderr {
	return &Stderr{marker, body}
}

// MarshalNix serializes the stderr message to the Nix wire format.
func (s Stderr) MarshalNix() ([]byte, error) {
	marker, err := s.Marker.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling marker: %w", err)
	}

	body, err := s.Body.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling body: %w", err)
	}

	return append(marker, body...), nil
}

// UnmarshalNix deserializes the stderr message from the Nix wire format.
//
// Note that usually, as the body type will only be known after the marker has been read,
// this will not be very useful in most client cases, and rather is implemented to fulfill
// the common interface.
//
// Clients wanting to unmarshal a stderr message should first read the marker, and then,
// based on the marker, unmarshal the body.
func (s *Stderr) UnmarshalNix(raw []byte) error {
	var newStderr Stderr

	if err := newStderr.Marker.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling marker: %w", err)
	}
	raw = raw[newStderr.Marker.Size():]

	if err := newStderr.Body.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling body: %w", err)
	}

	*s = newStderr
	return nil
}

// Size returns the size of the stderr message in bytes.
func (s Stderr) Size() uint64 {
	return s.Marker.Size() + s.Body.Size()
}
