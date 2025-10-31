package ies

// SL-TxPoolToReleaseList-r12 ::= SEQUENCE OF SL-TxPoolIdentity-r12
// SIZE (1..maxSL-TxPool-r12)
type SlTxpooltoreleaselistR12 struct {
	Value []SlTxpoolidentityR12 `lb:1,ub:maxSLTxpoolR12`
}
