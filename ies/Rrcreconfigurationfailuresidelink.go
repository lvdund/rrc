package ies

// RRCReconfigurationFailureSidelink ::= SEQUENCE
type Rrcreconfigurationfailuresidelink struct {
	RrcTransactionidentifierR16 RrcTransactionidentifier
	Criticalextensions          RrcreconfigurationfailuresidelinkCriticalextensions
}
