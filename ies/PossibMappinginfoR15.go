package ies

// PosSIB-MappingInfo-r15 ::= SEQUENCE OF PosSIB-Type-r15
// SIZE (1..maxSIB)
type PossibMappinginfoR15 struct {
	Value []PossibTypeR15 `lb:1,ub:maxSIB`
}
