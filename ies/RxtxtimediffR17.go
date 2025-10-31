package ies

import "rrc/utils"

// RxTxTimeDiff-r17 ::= SEQUENCE
// Extensible
type RxtxtimediffR17 struct {
	ResultK5R17 *utils.INTEGER `lb:0,ub:61565`
}
