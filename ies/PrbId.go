package ies

import "rrc/utils"

// PRB-Id ::= utils.INTEGER (0..maxNrofPhysicalResourceBlocks-1)
type PrbId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
}
