package ies

import "rrc/utils"

// SecurityModeComplete ::= SEQUENCE
type Securitymodecomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
