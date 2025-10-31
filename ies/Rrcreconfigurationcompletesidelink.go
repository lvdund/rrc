package ies

// RRCReconfigurationCompleteSidelink ::= SEQUENCE
type Rrcreconfigurationcompletesidelink struct {
	RrcTransactionidentifierR16 RrcTransactionidentifier
	Criticalextensions          RrcreconfigurationcompletesidelinkCriticalextensions
}
