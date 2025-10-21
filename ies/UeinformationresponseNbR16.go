package ies

import "rrc/utils"

// UEInformationResponse-NB-r16 ::= SEQUENCE
type UeinformationresponseNbR16 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
