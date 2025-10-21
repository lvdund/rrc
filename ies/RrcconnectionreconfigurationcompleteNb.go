package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-NB ::= SEQUENCE
type RrcconnectionreconfigurationcompleteNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
