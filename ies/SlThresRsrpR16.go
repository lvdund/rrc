package ies

import "rrc/utils"

// SL-Thres-RSRP-r16 ::= utils.INTEGER (0..66)
type SlThresRsrpR16 struct {
	Value utils.INTEGER `lb:0,ub:66`
}
