package ies

import "rrc/utils"

// SRS-ResourceId ::= utils.INTEGER (0..maxNrofSRS-Resources-1)
type SrsResourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRSResources1`
}
