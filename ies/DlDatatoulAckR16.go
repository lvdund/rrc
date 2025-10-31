package ies

import "rrc/utils"

// DL-DataToUL-ACK-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..8)
type DlDatatoulAckR16 struct {
	Value []utils.INTEGER `lb:1,ub:8`
}
