package ies

// SIB-MappingInfo-NB-r13 ::= SEQUENCE OF SIB-Type-NB-r13
// SIZE (0..maxSIB-1)
type SibMappinginfoNbR13 struct {
	Value []SibTypeNbR13 `lb:0,ub:maxSIB1`
}
