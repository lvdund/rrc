package ies

import "rrc/utils"

// SL-CBR-PPPP-TxPreconfigList-r14 ::= SEQUENCE OF SL-PPPP-TxPreconfigIndex-r14
// SIZE (1..8)
type SlCbrPpppTxpreconfiglistR14 struct {
	Value utils.Sequence[SlPpppTxpreconfigindexR14]
}
