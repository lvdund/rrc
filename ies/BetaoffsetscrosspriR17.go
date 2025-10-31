package ies

import "rrc/utils"

// BetaOffsetsCrossPri-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (3)
type BetaoffsetscrosspriR17 struct {
	Value []utils.INTEGER `lb:3,ub:3`
}
