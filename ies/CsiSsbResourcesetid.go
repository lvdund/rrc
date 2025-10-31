package ies

import "rrc/utils"

// CSI-SSB-ResourceSetId ::= utils.INTEGER (0..maxNrofCSI-SSB-ResourceSets-1)
type CsiSsbResourcesetid struct {
	Value utils.INTEGER `lb:0,ub:maxNrofCSISsbResourcesets1`
}
