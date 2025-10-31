package ies

// SL-PPPP-TxConfigIndex-v1530 ::= SEQUENCE
type SlPpppTxconfigindexV1530 struct {
	McsPsschRangelistR15 *[]McsPsschRangeR15 `lb:1,ub:maxCBRLevelR14`
}
