package ies

import "rrc/utils"

// SecurityModeFailure ::= SEQUENCE
type Securitymodefailure struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
