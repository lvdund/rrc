package ies

import "rrc/utils"

// PDSCH-ConfigIndex-r17 ::= utils.INTEGER (0..maxNrofPDSCH-ConfigPTM-1-r17)
type PdschConfigindexR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPDSCHConfigptm1R17`
}
