package ies

import "rrc/utils"

// DL-DataToUL-ACK-DCI-1-2-r16 ::= SEQUENCE OF utils.INTEGER // SIZE (1..8)
type DlDatatoulAckDci12R16 struct {
	Value []utils.INTEGER `lb:1,ub:8`
}
