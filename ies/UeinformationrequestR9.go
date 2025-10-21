package ies

import "rrc/utils"

// UEInformationRequest-r9 ::= SEQUENCE
type UeinformationrequestR9 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
