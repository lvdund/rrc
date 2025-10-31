package ies

import "rrc/utils"

// ZP-CSI-RS-ResourceId ::= utils.INTEGER (0..maxNrofZP-CSI-RS-Resources-1)
type ZpCsiRsResourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofZPCsiRsResources1`
}
