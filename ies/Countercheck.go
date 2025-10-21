package ies

import "rrc/utils"

// CounterCheck ::= SEQUENCE
type Countercheck struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
