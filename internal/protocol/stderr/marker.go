package stderr

import "github.com/msanft/proton/internal/primitive"

// Marker is a marker that (similar to the opcodes in the operation requests to the daemon)
// indicates the type of message that is being sent. It's usually followed by a body that
// contains the actual message.
type Marker uint64

const (
	MarkerWrite         Marker = 0x64617416
	MarkerError         Marker = 0x63787470
	MarkerNext          Marker = 0x6f6c6d67
	MarkerStartActivity Marker = 0x53545254
	MarkerStopActivity  Marker = 0x53544f50
	MarkerResult        Marker = 0x52534c54
	// MarkerLast indicates that there's no body following the marker.
	MarkerLast Marker = 0x616c7473
)

// MarshalNix serializes the marker to the Nix wire format.
func (m Marker) MarshalNix() ([]byte, error) {
	return primitive.NewInt(uint64(m)).MarshalNix()
}

// UnmarshalNix deserializes the marker from the Nix wire format.
func (m *Marker) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}
	*m = Marker(i.Value)
	return nil
}

// Size returns the size of the marker in bytes.
func (m Marker) Size() uint64 {
	return primitive.NewInt(uint64(m)).Size()
}
