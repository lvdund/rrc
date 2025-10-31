package ies

import "rrc/utils"

// DL-DataToUL-ACK-MulticastDCI-Format4-1-r17 ::= SEQUENCE OF utils.INTEGER // SIZE (1..8)
type DlDatatoulAckMulticastdciFormat41R17 struct {
	Value []utils.INTEGER `lb:1,ub:8`
}
