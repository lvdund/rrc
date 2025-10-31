package ies

// RRCConnectionSetupComplete ::= SEQUENCE
type Rrcconnectionsetupcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionsetupcompleteCriticalextensions
}
