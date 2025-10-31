package ies

// PosSIB-MappingInfo-r16 ::= SEQUENCE OF PosSIB-Type-r16
// SIZE (1..maxSIB)
type PossibMappinginfoR16 struct {
	Value []PossibTypeR16 `lb:1,ub:maxSIB`
}
