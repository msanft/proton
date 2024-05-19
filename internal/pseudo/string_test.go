package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringMarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	s := String("foobar")
	raw, err := s.MarshalNix()
	require.NoError(err)
	assert.Equal([]byte{6, 0, 0, 0, 0, 0, 0, 0, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0, 0}, raw)
}

func TestStringUnmarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	var s String
	err := s.UnmarshalNix([]byte{6, 0, 0, 0, 0, 0, 0, 0, 0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72, 0, 0})
	require.NoError(err)
	assert.Equal("foobar", string(s))
}
