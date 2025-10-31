package ies

import "rrc/utils"

// SSB-Index ::= utils.INTEGER (0..maxNrofSSBs-1)
type SsbIndex struct {
	Value utils.INTEGER `lb:0,ub:maxNrofSSBs1`
}
