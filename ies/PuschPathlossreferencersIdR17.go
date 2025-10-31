package ies

import "rrc/utils"

// PUSCH-PathlossReferenceRS-Id-r17 ::= utils.INTEGER (0..maxNrofPUSCH-PathlossReferenceRSs-1-r16)
type PuschPathlossreferencersIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUSCHPathlossreferencerss1R16`
}
