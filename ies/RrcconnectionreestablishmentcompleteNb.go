package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-NB ::= SEQUENCE
type RrcconnectionreestablishmentcompleteNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
