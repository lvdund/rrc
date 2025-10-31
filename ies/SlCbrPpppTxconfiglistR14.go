package ies

// SL-CBR-PPPP-TxConfigList-r14 ::= SEQUENCE OF SL-PPPP-TxConfigIndex-r14
// SIZE (1..8)
type SlCbrPpppTxconfiglistR14 struct {
	Value []SlPpppTxconfigindexR14 `lb:1,ub:8`
}
