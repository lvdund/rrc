package ies

import "rrc/utils"

// RRCConnectionResume-r13 ::= SEQUENCE
type RrcconnectionresumeR13 struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
