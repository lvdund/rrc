package ies

// SL-TxPoolToReleaseListV2X-r14 ::= SEQUENCE OF SL-V2X-TxPoolIdentity-r14
// SIZE (1.. maxSL-V2X-TxPool-r14)
type SlTxpooltoreleaselistv2xR14 struct {
	Value []SlV2xTxpoolidentityR14 `lb:1,ub:maxSLV2xTxpoolR14`
}
