package ies

import "rrc/utils"

// UEInformationRequest-NB-r16 ::= SEQUENCE
type UeinformationrequestNbR16 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
