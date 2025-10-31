package ies

// SIB-MappingInfo ::= SEQUENCE OF SIB-Type
// SIZE (0..maxSIB-1)
type SibMappinginfo struct {
	Value []SibType `lb:0,ub:maxSIB1`
}
