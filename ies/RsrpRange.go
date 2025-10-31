package ies

import "rrc/utils"

// RSRP-Range ::= utils.INTEGER (0..127)
type RsrpRange struct {
	Value utils.INTEGER `lb:0,ub:127`
}
