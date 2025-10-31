package ies

// RRCConnectionReconfiguration ::= SEQUENCE
type Rrcconnectionreconfiguration struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionreconfigurationCriticalextensions
}
