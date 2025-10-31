package ies

import "rrc/utils"

// PathlossReferenceRS-Id-r17 ::= utils.INTEGER (0..maxNrofPathlossReferenceRSs-1-r17)
type PathlossreferencersIdR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPathlossReferenceRSs1R17`
}
