package stderr

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

// StartActivity is the format used to report new errors over the stderr format.
type StartActivity struct {
	Activity primitive.Int
	Level    primitive.Int
	Kind     primitive.Int
	Message  primitive.ByteBuf
	Fields   LoggerFields
	Parent   primitive.Int
}

// MarshalNix serializes the startActivity message to the Nix wire format.
func (s StartActivity) MarshalNix() ([]byte, error) {
	var buf []byte

	activity, err := s.Activity.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling activity: %w", err)
	}
	buf = append(buf, activity...)

	level, err := s.Level.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling level: %w", err)
	}
	buf = append(buf, level...)

	kind, err := s.Kind.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling kind: %w", err)
	}
	buf = append(buf, kind...)

	msg, err := s.Message.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling message: %w", err)
	}
	buf = append(buf, msg...)

	fields, err := s.Fields.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling fields: %w", err)
	}
	buf = append(buf, fields...)

	parent, err := s.Parent.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling parent: %w", err)
	}
	buf = append(buf, parent...)

	return buf, nil
}

// UnmarshalNix deserializes the startActivity message from the Nix wire format.
func (s *StartActivity) UnmarshalNix(raw []byte) error {
	var newStartActivity StartActivity

	if err := newStartActivity.Activity.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling activity: %w", err)
	}
	raw = raw[newStartActivity.Activity.Size():]

	if err := newStartActivity.Level.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling level: %w", err)
	}
	raw = raw[newStartActivity.Level.Size():]

	if err := newStartActivity.Kind.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling kind: %w", err)
	}
	raw = raw[newStartActivity.Kind.Size():]

	if err := newStartActivity.Message.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling message: %w", err)
	}
	raw = raw[8+newStartActivity.Message.Len.Value:]

	if err := newStartActivity.Fields.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling fields: %w", err)
	}
	raw = raw[newStartActivity.Fields.Size():]

	if err := newStartActivity.Parent.UnmarshalNix(raw); err != nil {
		return fmt.Errorf("unmarshaling parent: %w", err)
	}

	*s = newStartActivity
	return nil
}

// Size returns the size of the startActivity message in bytes.
func (s StartActivity) Size() uint64 {
	return s.Activity.Size() + s.Level.Size() + s.Kind.Size() + s.Message.Size() + s.Fields.Size() + s.Parent.Size()
}
