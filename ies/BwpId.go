package ies

import "rrc/utils"

// BWP-Id ::= utils.INTEGER (0..maxNrofBWPs)
type BwpId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofBWPs`
}
