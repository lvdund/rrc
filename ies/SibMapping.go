package ies

// SIB-Mapping ::= SEQUENCE OF SIB-TypeInfo
// SIZE (1..maxSIB)
type SibMapping struct {
	Value []SibTypeinfo `lb:1,ub:maxSIB`
}
