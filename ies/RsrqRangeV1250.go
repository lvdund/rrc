package ies

import "rrc/utils"

// RSRQ-Range-v1250 ::= utils.INTEGER (-30..46)
type RsrqRangeV1250 struct {
	Value utils.INTEGER `lb:0,ub:46`
}
