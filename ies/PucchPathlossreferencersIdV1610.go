package ies

import "rrc/utils"

// PUCCH-PathlossReferenceRS-Id-v1610 ::= utils.INTEGER (maxNrofPUCCH-PathlossReferenceRSs..maxNrofPUCCH-PathlossReferenceRSs-1-r16)
type PucchPathlossreferencersIdV1610 struct {
	Value utils.INTEGER `lb:maxNrofPUCCHPathlossreferencerss,ub:maxNrofPUCCHPathlossreferencerss1R16`
}
