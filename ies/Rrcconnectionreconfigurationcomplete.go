package ies

// RRCConnectionReconfigurationComplete ::= SEQUENCE
type Rrcconnectionreconfigurationcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionreconfigurationcompleteCriticalextensions
}
