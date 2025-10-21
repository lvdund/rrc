package ies

import "rrc/utils"

// Tx-ConfigIndex-r14 ::= INTEGER (0..maxSL-V2X-TxConfig-1-r14)
type TxConfigindexR14 struct {
	Value utils.INTEGER
}
