package ies

// SL-CBR-PPPP-TxConfigList-v1530 ::= SEQUENCE OF SL-PPPP-TxConfigIndex-v1530
// SIZE (1..8)
type SlCbrPpppTxconfiglistV1530 struct {
	Value []SlPpppTxconfigindexV1530 `lb:1,ub:8`
}
