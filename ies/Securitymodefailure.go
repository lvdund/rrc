package ies

// SecurityModeFailure ::= SEQUENCE
type Securitymodefailure struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       SecuritymodefailureCriticalextensions
}
