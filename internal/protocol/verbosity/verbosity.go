package verbosity

import "github.com/msanft/proton/internal/primitive"

type Verbosity uint64

const (
	Error Verbosity = iota
	Warn
	Notice
	Info
	Talkative
	Chatty
	Debug
	Vomit
)

// MarshalNix serializes a verbosity level to the Nix wire format.
func (v Verbosity) MarshalNix() ([]byte, error) {
	return primitive.NewInt(uint64(v)).MarshalNix()
}

// UnmarshalNix deserializes a verbosity level from the Nix wire format.
func (v *Verbosity) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}
	*v = Verbosity(i.Value)
	return nil
}

// Size returns the size of the verbosity level in bytes.
func (v Verbosity) Size() uint64 {
	return primitive.NewInt(uint64(v)).Size()
}
