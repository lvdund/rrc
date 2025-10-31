package ies

import "rrc/utils"

// ServingCellConfigCommonSIB-ssb-PositionsInBurst ::= SEQUENCE
type ServingcellconfigcommonsibSsbPositionsinburst struct {
	Inonegroup    utils.BITSTRING  `lb:8,ub:8`
	Grouppresence *utils.BITSTRING `lb:8,ub:8`
}
