package ies

// SecurityModeCommand ::= SEQUENCE
type Securitymodecommand struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       SecuritymodecommandCriticalextensions
}
