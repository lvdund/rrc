package ies

// SL-PPPP-TxPreconfigIndex-v1530 ::= SEQUENCE
type SlPpppTxpreconfigindexV1530 struct {
	McsPsschRangeR15 *[]McsPsschRangeR15 `lb:1,ub:maxCBRLevelR14`
}
