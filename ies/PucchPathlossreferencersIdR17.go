package ies

import "rrc/utils"

// PUCCH-PathlossReferenceRS-Id-r17 ::= utils.INTEGER (0..maxNrofPUCCH-PathlossReferenceRSs-1-r17)
type PucchPathlossreferencersIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUCCHPathlossreferencerss1R17`
}
