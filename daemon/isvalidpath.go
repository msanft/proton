package daemon

import (
	"fmt"

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

	rest, err := c.ParseStderr()
	if err != nil {
		return false, fmt.Errorf("receive stderr: %w", err)
	}

	var validity pseudo.Bool
	if rest == nil {
		validity, err = c.readNixBool()
		if err != nil {
			return false, fmt.Errorf("read validity: %w", err)
		}
	} else {
		if err := validity.UnmarshalNix(rest); err != nil {
			return false, fmt.Errorf("unmarshal validity: %w", err)
		}
	}

	return bool(validity), nil
}
