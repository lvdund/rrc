package ies

import "rrc/utils"

// RNReconfigurationComplete-r10 ::= SEQUENCE
type RnreconfigurationcompleteR10 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
