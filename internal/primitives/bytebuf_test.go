package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalNix(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	t.Run("multiple of 8", func(t *testing.T) {
		b := NewByteBuf([]byte{1, 2, 3, 4, 5, 6, 7, 8})

		raw, err := b.MarshalNix()
		require.NoError(err)

		assert.Equal([]byte{8, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8}, raw)
	})

	t.Run("not multiple of 8", func(t *testing.T) {
		b := NewByteBuf([]byte{1, 2, 3, 4, 5, 6, 7})

		raw, err := b.MarshalNix()
		require.NoError(err)

		assert.Equal([]byte{7, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 0}, raw)
	})
}

func TestUnmarshalNix(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		b := ByteBuf{}

		err := b.UnmarshalNix([]byte{8, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8})
		require.NoError(err)

		assert.Equal([]byte{1, 2, 3, 4, 5, 6, 7, 8}, b.Buf)
		assert.Equal(uint64(8), b.Len.Value)
	})

	t.Run("less than 8 bytes", func(t *testing.T) {
		b := ByteBuf{}

		err := b.UnmarshalNix([]byte{42})
		require.Error(err)
	})
}
