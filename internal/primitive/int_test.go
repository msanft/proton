package primitive

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntMarshalNix(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	i := NewInt(42)

	b, err := i.MarshalNix()
	require.NoError(err)

	assert.Equal([]byte{42, 0, 0, 0, 0, 0, 0, 0}, b)
}

func TestIntUnmarshalNix(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	var i Int

	err := i.UnmarshalNix([]byte{42, 0, 0, 0, 0, 0, 0, 0})
	require.NoError(err)

	assert.Equal(uint64(42), i.Value)
}
