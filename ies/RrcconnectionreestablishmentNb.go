package ies

import "rrc/utils"

// RRCConnectionReestablishment-NB ::= SEQUENCE
type RrcconnectionreestablishmentNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
