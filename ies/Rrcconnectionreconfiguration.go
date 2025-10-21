package ies

import "rrc/utils"

// RRCConnectionReconfiguration ::= SEQUENCE
type Rrcconnectionreconfiguration struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
