package ies

// SL-CBR-PPPP-TxPreconfigList-r14 ::= SEQUENCE OF SL-PPPP-TxPreconfigIndex-r14
// SIZE (1..8)
type SlCbrPpppTxpreconfiglistR14 struct {
	Value []SlPpppTxpreconfigindexR14 `lb:1,ub:8`
}
