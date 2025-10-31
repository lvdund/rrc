package ies

// SL-TxPercentageList-r16 ::= SEQUENCE OF SL-TxPercentageConfig-r16
// SIZE (8)
type SlTxpercentagelistR16 struct {
	Value []SlTxpercentageconfigR16 `lb:8,ub:8`
}
