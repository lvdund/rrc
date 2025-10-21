package ies

import "rrc/utils"

// RRCConnectionResumeComplete-r13 ::= SEQUENCE
type RrcconnectionresumecompleteR13 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
