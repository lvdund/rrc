package ies

// RRCConnectionSetup ::= SEQUENCE
type Rrcconnectionsetup struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionsetupCriticalextensions
}
