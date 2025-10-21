package ies

import "rrc/utils"

// RRCConnectionReestablishment ::= SEQUENCE
type Rrcconnectionreestablishment struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       interface{}
}
