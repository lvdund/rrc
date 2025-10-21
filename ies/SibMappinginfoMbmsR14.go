package ies

import "rrc/utils"

// SIB-MappingInfo-MBMS-r14 ::= SEQUENCE OF SIB-Type-MBMS-r14
// SIZE (0..maxSIB-1)
type SibMappinginfoMbmsR14 struct {
	Value utils.Sequence[SibTypeMbmsR14]
}
