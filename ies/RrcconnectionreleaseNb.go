package ies

import "rrc/utils"

// RRCConnectionRelease-NB ::= SEQUENCE
type RrcconnectionreleaseNb struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
