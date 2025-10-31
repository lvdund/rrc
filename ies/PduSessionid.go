package ies

import "rrc/utils"

// PDU-SessionID ::= utils.INTEGER (0..255)
type PduSessionid struct {
	Value utils.INTEGER `lb:0,ub:255`
}
