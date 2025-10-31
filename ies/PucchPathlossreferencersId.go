package ies

import "rrc/utils"

// PUCCH-PathlossReferenceRS-Id ::= utils.INTEGER (0..maxNrofPUCCH-PathlossReferenceRSs-1)
type PucchPathlossreferencersId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUCCHPathlossreferencerss1`
}
