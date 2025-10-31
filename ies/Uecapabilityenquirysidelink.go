package ies

// UECapabilityEnquirySidelink ::= SEQUENCE
type Uecapabilityenquirysidelink struct {
	RrcTransactionidentifierR16 RrcTransactionidentifier
	Criticalextensions          UecapabilityenquirysidelinkCriticalextensions
}
