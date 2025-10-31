package ies

import "rrc/utils"

// SRS-ResourceSetId ::= utils.INTEGER (0..maxNrofSRS-ResourceSets-1)
type SrsResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRSResourcesets1`
}
