package ies

import "rrc/utils"

// PosSIB-MappingInfo-r15 ::= SEQUENCE OF PosSIB-Type-r15
// SIZE (1..maxSIB)
type PossibMappinginfoR15 struct {
	Value utils.Sequence[PossibTypeR15]
}
