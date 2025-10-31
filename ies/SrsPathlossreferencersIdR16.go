package ies

import "rrc/utils"

// SRS-PathlossReferenceRS-Id-r16 ::= utils.INTEGER (0..maxNrofSRS-PathlossReferenceRS-1-r16)
type SrsPathlossreferencersIdR16 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSRSPathlossreferencers1R16`
}
