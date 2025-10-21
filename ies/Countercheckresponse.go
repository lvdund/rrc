package ies

import "rrc/utils"

// CounterCheckResponse ::= SEQUENCE
type Countercheckresponse struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
