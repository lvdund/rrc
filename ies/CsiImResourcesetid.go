package ies

import "rrc/utils"

// CSI-IM-ResourceSetId ::= utils.INTEGER (0..maxNrofCSI-IM-ResourceSets-1)
type CsiImResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSIImResourcesets1`
}
