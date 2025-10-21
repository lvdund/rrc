package ies

import "rrc/utils"

// RRCConnectionSetupComplete ::= SEQUENCE
type Rrcconnectionsetupcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
