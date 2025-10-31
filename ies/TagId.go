package ies

import "rrc/utils"

// TAG-Id ::= utils.INTEGER (0..maxNrofTAGs-1)
type TagId struct {
	Value utils.INTEGER `lb:0,ub:maxNrofTAGs1`
}
