package ies

import "rrc/utils"

// Tx-PreconfigIndex-r14 ::= INTEGER (0..maxSL-V2X-TxConfig2-1-r14)
type TxPreconfigindexR14 struct {
	Value utils.INTEGER
}
