package ies

import "rrc/utils"

// PUSCH-PathlossReferenceRS-Id-v1610 ::= utils.INTEGER (maxNrofPUSCH-PathlossReferenceRSs..maxNrofPUSCH-PathlossReferenceRSs-1-r16)
type PuschPathlossreferencersIdV1610 struct {
	Value utils.INTEGER `lb:maxNrofPUSCHPathlossreferencerss,ub:maxNrofPUSCHPathlossreferencerss1R16`
}
