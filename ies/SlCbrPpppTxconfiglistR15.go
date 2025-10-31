package ies

// SL-CBR-PPPP-TxConfigList-r15 ::= SEQUENCE OF SL-PPPP-TxConfigIndex-r15
// SIZE (1..8)
type SlCbrPpppTxconfiglistR15 struct {
	Value []SlPpppTxconfigindexR15 `lb:1,ub:8`
}
