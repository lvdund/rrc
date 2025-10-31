package ies

// HandoverFromEUTRAPreparationRequest ::= SEQUENCE
type Handoverfromeutrapreparationrequest struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       HandoverfromeutrapreparationrequestCriticalextensions
}
