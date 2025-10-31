package ies

// SIB-MappingInfo-NB-v1530 ::= SEQUENCE OF SIB-Type-NB-v1530
// SIZE (1..8)
type SibMappinginfoNbV1530 struct {
	Value []SibTypeNbV1530 `lb:1,ub:8`
}
