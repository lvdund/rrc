package ies

import "rrc/utils"

// SIB-MappingInfo ::= SEQUENCE OF SIB-Type
// SIZE (0..maxSIB-1)
type SibMappinginfo struct {
	Value utils.Sequence[SibType]
}
