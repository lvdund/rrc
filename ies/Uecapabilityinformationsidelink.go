package ies

// UECapabilityInformationSidelink ::= SEQUENCE
type Uecapabilityinformationsidelink struct {
	RrcTransactionidentifierR16 RrcTransactionidentifier
	Criticalextensions          UecapabilityinformationsidelinkCriticalextensions
}
