package ies

import "rrc/utils"

// UECapabilityInformation-NB ::= SEQUENCE
type UecapabilityinformationNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
