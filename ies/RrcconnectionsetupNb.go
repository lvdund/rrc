package ies

import "rrc/utils"

// RRCConnectionSetup-NB ::= SEQUENCE
type RrcconnectionsetupNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
