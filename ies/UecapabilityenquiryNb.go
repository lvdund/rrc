package ies

import "rrc/utils"

// UECapabilityEnquiry-NB ::= SEQUENCE
type UecapabilityenquiryNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
