package ies

// MobilityFromEUTRACommand ::= SEQUENCE
type Mobilityfromeutracommand struct {
	RrcTransactionidentifier RrcTransactionidentifier
	Criticalextensions       MobilityfromeutracommandCriticalextensions
}
