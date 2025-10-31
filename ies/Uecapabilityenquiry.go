package ies

// UECapabilityEnquiry ::= SEQUENCE
type Uecapabilityenquiry struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       UecapabilityenquiryCriticalextensions
}
