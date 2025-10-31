package ies

// DLInformationTransfer ::= SEQUENCE
type Dlinformationtransfer struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       DlinformationtransferCriticalextensions
}
