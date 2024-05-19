package protocol

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

const (
	// Major is the major version number of the Nix daemon protocol supported by this implementation.
	Major = 1
	// Minor is the minor version number of the Nix daemon protocol supported by this implementation.
	Minor = 37
)

// Version is a version of the Nix daemon protocol.
type Version struct {
	major uint8
	minor uint8
}

// NewVersion creates a new protocol version with the given major and minor version numbers.
func NewVersion(major, minor uint8) Version {
	return Version{major: major, minor: minor}
}

// OwnVersion returns the version of the protocol supported by this implementation.
func OwnVersion() Version {
	return NewVersion(Major, Minor)
}

// MarshalNix serializes the protocol version to the Nix wire format.
func (v Version) MarshalNix() ([]byte, error) {
	i := uint64(v.major)<<8 | uint64(v.minor)
	return primitive.NewInt(i).MarshalNix()
}

// UnmarshalNix deserializes the protocol version from the Nix wire format.
func (v *Version) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}
	v.major = uint8(i.Value >> 8)
	v.minor = uint8(i.Value)
	return nil
}

// String returns a string representation of the protocol version.
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d", v.major, v.minor)
}
