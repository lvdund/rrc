package ies

// SL-CBR-PPPP-TxPreconfigList-v1530 ::= SEQUENCE OF SL-PPPP-TxPreconfigIndex-v1530
// SIZE (1..8)
type SlCbrPpppTxpreconfiglistV1530 struct {
	Value []SlPpppTxpreconfigindexV1530 `lb:1,ub:8`
}
