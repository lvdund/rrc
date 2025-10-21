package ies

import "rrc/utils"

// RRCConnectionResumeComplete-NB ::= SEQUENCE
type RrcconnectionresumecompleteNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
