package ies

import "rrc/utils"

// SL-ThresPSSCH-RSRP-r14 ::= utils.INTEGER (0..66)
type SlThrespsschRsrpR14 struct {
	Value utils.INTEGER `lb:0,ub:66`
}
