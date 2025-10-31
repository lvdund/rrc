package ies

// RRCConnectionReestablishment ::= SEQUENCE
type Rrcconnectionreestablishment struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       RrcconnectionreestablishmentCriticalextensions
}
