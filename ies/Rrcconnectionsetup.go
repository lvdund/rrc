package ies

import "rrc/utils"

// RRCConnectionSetup ::= SEQUENCE
type Rrcconnectionsetup struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
