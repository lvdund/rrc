package ies

import "rrc/utils"

// Q-RxLevMin ::= utils.INTEGER (-70..-22)
type QRxlevmin struct {
	Value utils.INTEGER `lb:0,ub:-22`
}
