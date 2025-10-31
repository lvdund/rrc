package ies

import "rrc/utils"

// SC-MTCH-Info-r13 ::= SEQUENCE
// Extensible
type ScMtchInfoR13 struct {
	MbmssessioninfoR13      MbmssessioninfoR13
	GRntiR13                utils.BITSTRING `lb:16,ub:16`
	ScMtchSchedulinginfoR13 *ScMtchSchedulinginfoR13
	ScMtchNeighbourcellR13  *utils.BITSTRING `lb:maxNeighCellScptmR13,ub:maxNeighCellScptmR13`
}
