package ies

import "rrc/utils"

// RNReconfiguration-r10 ::= SEQUENCE
type RnreconfigurationR10 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
