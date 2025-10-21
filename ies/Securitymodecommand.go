package ies

import "rrc/utils"

// SecurityModeCommand ::= SEQUENCE
type Securitymodecommand struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
