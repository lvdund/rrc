package ies

import "rrc/utils"

// MBMS-CarrierType-r14 ::= SEQUENCE
type MbmsCarriertypeR14 struct {
	CarriertypeR14 MbmsCarriertypeR14CarriertypeR14
	FrameoffsetR14 *utils.INTEGER `lb:0,ub:3`
}
