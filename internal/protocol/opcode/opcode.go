package opcode

import "github.com/msanft/proton/internal/primitive"

// Opcode represent an operation code used in the Nix daemon protocol.
type Opcode uint64

const (
	IsValidPath              Opcode = 1
	QueryReferrers           Opcode = 6
	AddToStore               Opcode = 7
	BuildPaths               Opcode = 9
	EnsurePath               Opcode = 10
	AddTempRoot              Opcode = 11
	FindRoots                Opcode = 14
	SetOptions               Opcode = 19
	CollectGarbage           Opcode = 20
	QueryAllValidPaths       Opcode = 23
	QueryPathInfo            Opcode = 26
	QueryPathFromHashPart    Opcode = 29
	QueryValidPaths          Opcode = 31
	QuerySubstitutablePaths  Opcode = 32
	QueryValidDerivers       Opcode = 33
	OptimiseStore            Opcode = 34
	VerifyStore              Opcode = 35
	BuildDerivation          Opcode = 36
	AddSignatures            Opcode = 37
	NarFromPath              Opcode = 38
	AddToStoreNar            Opcode = 39
	QueryMissing             Opcode = 40
	QueryDerivationOutputMap Opcode = 41
	RegisterDrvOutput        Opcode = 42
	QueryRealisation         Opcode = 43
	AddMultipleToStore       Opcode = 44
	AddBuildLog              Opcode = 45
	BuildPathsWithResults    Opcode = 46
)

// MarshalNix serializes the opcode to the Nix wire format.
func (o Opcode) MarshalNix() ([]byte, error) {
	return primitive.NewInt(uint64(o)).MarshalNix()
}

// UnmarshalNix deserializes the opcode from the Nix wire format.
func (o *Opcode) UnmarshalNix(raw []byte) error {
	var i primitive.Int
	if err := i.UnmarshalNix(raw); err != nil {
		return err
	}
	*o = Opcode(i.Value)
	return nil
}

// Size returns the size of the opcode in bytes.
func (o Opcode) Size() uint64 {
	return primitive.NewInt(uint64(o)).Size()
}
