package daemon

import (
	"fmt"

	"github.com/msanft/proton/internal/primitive"
	"github.com/msanft/proton/internal/protocol/opcode"
	"github.com/msanft/proton/internal/protocol/operation"
	"github.com/msanft/proton/internal/pseudo"
)

// IsValidPath returns wheter the given store path is valid
// on the daemon's store.
func (c *Conn) IsValidPath(path string) (bool, error) {
	p := pseudo.String(path)
	if err := c.writeNix(operation.NewOperation(
		opcode.IsValidPath,
		&p,
	)); err != nil {
		return false, fmt.Errorf("write store path to connection: %w", err)
	}

	// TODO: This seems to be some kind of response identifier? need to check
	var i primitive.Int
	_ = c.readNNix(&i, 8)

	b, err := c.readNixBool()
	if err != nil {
		return false, fmt.Errorf("read validity: %w", err)
	}

	return bool(b), nil
}
