package ies

import "rrc/utils"

// SL-CBR-PPPP-TxConfigList-r15 ::= SEQUENCE OF SL-PPPP-TxConfigIndex-r15
// SIZE (1..8)
type SlCbrPpppTxconfiglistR15 struct {
	Value []SlPpppTxconfigindexR15
}
