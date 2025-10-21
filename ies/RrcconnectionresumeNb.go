package ies

import "rrc/utils"

// RRCConnectionResume-NB ::= SEQUENCE
type RrcconnectionresumeNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
