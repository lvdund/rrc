package ies

import "rrc/utils"

// DL-DataToUL-ACK-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..8)
type DlDatatoulAckR17 struct {
	Value []utils.INTEGER `lb:1,ub:8`
}
