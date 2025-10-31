package ies

// SecurityModeComplete ::= SEQUENCE
type Securitymodecomplete struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       SecuritymodecompleteCriticalextensions
}
