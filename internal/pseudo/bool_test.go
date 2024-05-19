package pseudo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolMarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	t.Run("true", func(t *testing.T) {
		b := Bool(true)
		raw, err := b.MarshalNix()
		require.NoError(err)
		assert.Equal([]byte{1, 0, 0, 0, 0, 0, 0, 0}, raw)
	})
	t.Run("true", func(t *testing.T) {
		b := Bool(false)
		raw, err := b.MarshalNix()
		require.NoError(err)
		assert.Equal([]byte{0, 0, 0, 0, 0, 0, 0, 0}, raw)
	})
}

func TestBoolUnmarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	t.Run("true", func(t *testing.T) {
		var b Bool
		err := b.UnmarshalNix([]byte{1, 0, 0, 0, 0, 0, 0, 0})
		require.NoError(err)
		assert.True(bool(b))
	})

	t.Run("true, non-1", func(t *testing.T) {
		var b Bool
		err := b.UnmarshalNix([]byte{42, 13, 37, 0, 0, 0, 0, 0})
		require.NoError(err)
		assert.True(bool(b))
	})

	t.Run("false", func(t *testing.T) {
		var b Bool
		err := b.UnmarshalNix([]byte{0, 0, 0, 0, 0, 0, 0, 0})
		require.NoError(err)
		assert.False(bool(b))
	})
}
