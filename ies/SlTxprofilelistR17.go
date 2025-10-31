package ies

// SL-TxProfileList-r17 ::= SEQUENCE OF SL-TxProfile-r17
// SIZE (1..256)
type SlTxprofilelistR17 struct {
	Value []SlTxprofileR17 `lb:1,ub:256`
}
