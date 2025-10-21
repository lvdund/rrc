package ies

import "rrc/utils"

// RRCConnectionRelease ::= SEQUENCE
type Rrcconnectionrelease struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
