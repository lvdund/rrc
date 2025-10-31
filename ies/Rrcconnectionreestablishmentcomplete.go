package ies

// RRCConnectionReestablishmentComplete ::= SEQUENCE
type Rrcconnectionreestablishmentcomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionreestablishmentcompleteCriticalextensions
}
