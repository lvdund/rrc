package ies

import "rrc/utils"

// NZP-CSI-RS-ResourceSetId ::= utils.INTEGER (0..maxNrofNZP-CSI-RS-ResourceSets-1)
type NzpCsiRsResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofNZPCsiRsResourcesets1`
}
