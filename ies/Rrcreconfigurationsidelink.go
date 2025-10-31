package ies

// RRCReconfigurationSidelink ::= SEQUENCE
type Rrcreconfigurationsidelink struct {
	RrcTransactionidentifierR16 RrcTransactionidentifier
	Criticalextensions          RrcreconfigurationsidelinkCriticalextensions
}
