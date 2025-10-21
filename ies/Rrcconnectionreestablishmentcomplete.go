package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete ::= SEQUENCE
type Rrcconnectionreestablishmentcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
