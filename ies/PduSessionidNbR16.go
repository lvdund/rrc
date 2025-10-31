package ies

import "rrc/utils"

// PDU-SessionID-NB-r16 ::= utils.INTEGER (0..255)
type PduSessionidNbR16 struct {
	Value utils.INTEGER `lb:0,ub:255`
}
