package ies

// SIB-MappingInfo-MBMS-r14 ::= SEQUENCE OF SIB-Type-MBMS-r14
// SIZE (0..maxSIB-1)
type SibMappinginfoMbmsR14 struct {
	Value []SibTypeMbmsR14 `lb:0,ub:maxSIB1`
}
