package ies

import "rrc/utils"

// RRCConnectionSetupComplete-NB ::= SEQUENCE
type RrcconnectionsetupcompleteNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
