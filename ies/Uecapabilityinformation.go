package ies

// UECapabilityInformation ::= SEQUENCE
type Uecapabilityinformation struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       UecapabilityinformationCriticalextensions
}
