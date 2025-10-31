package ies

import "rrc/utils"

// RSRP-Range-v1360 ::= utils.INTEGER (-17..-1)
type RsrpRangeV1360 struct {
	Value utils.INTEGER `lb:0,ub:-1`
}
