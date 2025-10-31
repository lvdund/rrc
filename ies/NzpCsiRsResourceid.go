package ies

import "rrc/utils"

// NZP-CSI-RS-ResourceId ::= utils.INTEGER (0..maxNrofNZP-CSI-RS-Resources-1)
type NzpCsiRsResourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofNZPCsiRsResources1`
}
