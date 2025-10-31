package ies

import "rrc/utils"

// TimeUntilReconnection-r16 ::= utils.INTEGER (0..172800)
type TimeuntilreconnectionR16 struct {
	Value utils.INTEGER `lb:0,ub:172800`
}
