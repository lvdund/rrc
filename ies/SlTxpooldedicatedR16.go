package ies

// SL-TxPoolDedicated-r16 ::= SEQUENCE
type SlTxpooldedicatedR16 struct {
	SlPooltoreleaselistR16 *[]SlResourcepoolidR16     `lb:1,ub:maxNrofTXPoolR16`
	SlPooltoaddmodlistR16  *[]SlResourcepoolconfigR16 `lb:1,ub:maxNrofTXPoolR16`
}
