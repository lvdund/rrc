package ies

import "rrc/utils"

// UEInformationResponse-r9 ::= SEQUENCE
type UeinformationresponseR9 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
