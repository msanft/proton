package stderr

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
)

// is the format used to report new errors over the stderr format.
type LoggerFields struct {
	Fields []LoggerField
}

// MarshalNix serializes the logger fields to the Nix wire format.
func (f LoggerFields) MarshalNix() ([]byte, error) {
	var buf []byte

	for _, field := range f.Fields {
		fieldBuf, err := field.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling field: %w", err)
		}
		buf = append(buf, fieldBuf...)
	}

	return buf, nil
}

// UnmarshalNix deserializes the logger fields from the Nix wire format.
func (f *LoggerFields) UnmarshalNix(raw []byte) error {
	var newFields LoggerFields

	for len(raw) > 0 {
		var field LoggerField
		if err := field.UnmarshalNix(raw); err != nil {
			return fmt.Errorf("unmarshaling field: %w", err)
		}
		newFields.Fields = append(newFields.Fields, field)

		fieldSize := uint64(16)
		if field.Kind == KindString {
			fieldSize += field.ContentString.Len.Value
		}
		raw = raw[fieldSize:]
	}

	*f = newFields
	return nil
}

// Size returns the size of the logger fields in bytes.
func (f LoggerFields) Size() uint64 {
	var size uint64
	for _, field := range f.Fields {
		size += field.Size()
	}
	return size
}

// LoggerField is a field in a log message.
type LoggerField struct {
	Kind          Kind
	ContentInt    primitive.Int
	ContentString primitive.ByteBuf
}

// MarshalNix serializes the logger field to the Nix wire format.
func (f LoggerField) MarshalNix() ([]byte, error) {
	var buf []byte

	kind, err := f.Kind.MarshalNix()
	if err != nil {
		return nil, fmt.Errorf("marshaling kind: %w", err)
	}
	buf = append(buf, kind...)

	switch f.Kind {
	case KindInt:
		contentInt, err := f.ContentInt.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling integer content: %w", err)
		}
		buf = append(buf, contentInt...)
	case KindString:
		contentString, err := f.ContentString.MarshalNix()
		if err != nil {
			return nil, fmt.Errorf("marshaling string content: %w", err)
		}
		buf = append(buf, contentString...)
	default:
		return nil, fmt.Errorf("unknown kind: %d", f.Kind)
	}

	return buf, nil
}

// UnmarshalNix deserializes the logger field from the Nix wire format.
func (f *LoggerField) UnmarshalNix(raw []byte) error {
	var newLoggerField LoggerField

	if err := newLoggerField.Kind.UnmarshalNix(raw[:8]); err != nil {
		return fmt.Errorf("unmarshaling kind: %w", err)
	}

	switch newLoggerField.Kind {
	case KindInt:
		if err := newLoggerField.ContentInt.UnmarshalNix(raw[8:]); err != nil {
			return fmt.Errorf("unmarshaling integer content: %w", err)
		}
	case KindString:
		if err := newLoggerField.ContentString.UnmarshalNix(raw[8:]); err != nil {
			return fmt.Errorf("unmarshaling string content: %w", err)
		}
	default:
		return fmt.Errorf("unknown kind: %d", newLoggerField.Kind)
	}

	*f = newLoggerField
	return nil
}

// Size returns the size of the logger field in bytes.
func (f *LoggerField) Size() uint64 {
	size := f.Kind.Size()
	switch f.Kind {
	case KindInt:
		size += f.ContentInt.Size()
	case KindString:
		size += f.ContentString.Size()
	}
	return size
}

// Kind is the type of the field.
type Kind uint64

const (
	// KindString is a string field.
	KindString Kind = 1
	// KindInt is an integer field.
	KindInt Kind = 0
)

// MarshalNix serializes the kind to the Nix wire format.
func (k Kind) MarshalNix() ([]byte, error) {
	return primitive.NewInt(uint64(k)).MarshalNix()
}

// UnmarshalNix deserializes the kind from the Nix wire format.
func (k *Kind) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}

	*k = Kind(i.Value)
	return nil
}

// Size returns the size of the kind in bytes.
func (k *Kind) Size() uint64 {
	return primitive.NewInt(uint64(*k)).Size()
}
