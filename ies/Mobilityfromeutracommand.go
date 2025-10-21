package ies

import "rrc/utils"

// MobilityFromEUTRACommand ::= SEQUENCE
type Mobilityfromeutracommand struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
