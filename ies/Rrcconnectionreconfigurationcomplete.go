package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete ::= SEQUENCE
type Rrcconnectionreconfigurationcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
