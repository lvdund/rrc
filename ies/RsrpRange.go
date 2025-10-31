package ies

import "rrc/utils"

// RSRP-Range ::= utils.INTEGER (0..97)
type RsrpRange struct {
	Value utils.INTEGER `lb:0,ub:97`
}
