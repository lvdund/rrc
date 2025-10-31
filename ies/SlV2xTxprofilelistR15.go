package ies

// SL-V2X-TxProfileList-r15 ::= SEQUENCE OF SL-V2X-TxProfile-r15
// SIZE (1..256)
type SlV2xTxprofilelistR15 struct {
	Value []SlV2xTxprofileR15 `lb:1,ub:256`
}
