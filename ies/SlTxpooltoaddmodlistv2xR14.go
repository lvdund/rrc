package ies

// SL-TxPoolToAddModListV2X-r14 ::= SEQUENCE OF SL-TxPoolToAddMod-r14
// SIZE (1.. maxSL-V2X-TxPool-r14)
type SlTxpooltoaddmodlistv2xR14 struct {
	Value []SlTxpooltoaddmodR14 `lb:1,ub:maxSLV2xTxpoolR14`
}
