package ies

import "rrc/utils"

// CSI-RS-Index ::= utils.INTEGER (0..maxNrofCSI-RS-ResourcesRRM-1)
type CsiRsIndex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSIRsResourcesrrm1`
}
