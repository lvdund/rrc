package ies

import "rrc/utils"

// PUSCH-PathlossReferenceRS-Id ::= utils.INTEGER (0..maxNrofPUSCH-PathlossReferenceRSs-1)
type PuschPathlossreferencersId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPUSCHPathlossreferencerss1`
}
