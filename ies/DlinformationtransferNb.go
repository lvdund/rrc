package ies

import "rrc/utils"

// DLInformationTransfer-NB ::= SEQUENCE
type DlinformationtransferNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
