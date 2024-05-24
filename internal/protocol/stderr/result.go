package stderr

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

// Result is the format used to report results over the stderr format.
type Result struct {
	Activity primitive.Int
	Kind     primitive.Int
	Fields   LoggerFields
}

// MarshalNix serializes the result to the Nix wire format.
func (r Result) MarshalNix() ([]byte, error) {
	var buf []byte

	activity, err := r.Activity.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling activity: %w", err)
	}
	buf = append(buf, activity...)

	kind, err := r.Kind.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling kind: %w", err)
	}
	buf = append(buf, kind...)

	fields, err := r.Fields.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling fields: %w", err)
	}
	buf = append(buf, fields...)

	return buf, nil
}

// UnmarshalNix deserializes the result from the Nix wire format.
func (r *Result) UnmarshalNix(raw []byte) error {
	var newResult Result

	if err := newResult.Activity.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling activity: %w", err)
	}
	raw = raw[newResult.Activity.Size():]

	if err := newResult.Kind.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling kind: %w", err)
	}
	raw = raw[newResult.Kind.Size():]

	if err := newResult.Fields.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling fields: %w", err)
	}

	*r = newResult
	return nil
}

// Size returns the size of the result in bytes.
func (r Result) Size() uint64 {
	return r.Activity.Size() + r.Kind.Size() + r.Fields.Size()
}
