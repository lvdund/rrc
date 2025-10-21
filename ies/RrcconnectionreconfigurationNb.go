package ies

import "rrc/utils"

// RRCConnectionReconfiguration-NB ::= SEQUENCE
type RrcconnectionreconfigurationNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
