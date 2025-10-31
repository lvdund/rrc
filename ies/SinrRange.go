package ies

import "rrc/utils"

// SINR-Range ::= utils.INTEGER (0..127)
type SinrRange struct {
	Value utils.INTEGER `lb:0,ub:127`
}
