package ies

import "rrc/utils"

// ZP-CSI-RS-ResourceSetId ::= utils.INTEGER (0..maxNrofZP-CSI-RS-ResourceSets-1)
type ZpCsiRsResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofZPCsiRsResourcesets1`
}
