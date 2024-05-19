package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVersionMarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	v := NewVersion(1, 34)
	raw, err := v.MarshalNix()
	require.NoError(err)
	assert.Equal([]byte{34, 1, 0, 0, 0, 0, 0, 0}, raw)
}
