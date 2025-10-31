package ies

// SIB-Mapping-v1700 ::= SEQUENCE OF SIB-TypeInfo-v1700
// SIZE (1..maxSIB)
type SibMappingV1700 struct {
	Value []SibTypeinfoV1700 `lb:1,ub:maxSIB`
}
