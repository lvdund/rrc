package ies

// RRCConnectionRelease ::= SEQUENCE
type Rrcconnectionrelease struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionreleaseCriticalextensions
}
