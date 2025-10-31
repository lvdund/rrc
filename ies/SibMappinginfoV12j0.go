package ies

// SIB-MappingInfo-v12j0 ::= SEQUENCE OF SIB-Type-v12j0
// SIZE (1..maxSIB-1)
type SibMappinginfoV12j0 struct {
	Value []SibTypeV12j0 `lb:1,ub:maxSIB1`
}
