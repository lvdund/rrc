package ies

import "rrc/utils"

// CSI-IM-ResourceId ::= utils.INTEGER (0..maxNrofCSI-IM-Resources-1)
type CsiImResourceid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSIImResources1`
}
