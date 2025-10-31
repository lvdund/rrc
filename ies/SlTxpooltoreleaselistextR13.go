package ies

// SL-TxPoolToReleaseListExt-r13 ::= SEQUENCE OF SL-TxPoolIdentity-v1310
// SIZE (1..maxSL-TxPool-v1310)
type SlTxpooltoreleaselistextR13 struct {
	Value []SlTxpoolidentityV1310 `lb:1,ub:maxSLTxpoolV1310`
}
