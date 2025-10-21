package ies

import "rrc/utils"

// UECapabilityEnquiry ::= SEQUENCE
type Uecapabilityenquiry struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
