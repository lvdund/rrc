package ies

import "rrc/utils"

// HandoverFromEUTRAPreparationRequest ::= SEQUENCE
type Handoverfromeutrapreparationrequest struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
