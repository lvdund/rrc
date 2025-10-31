package ies

// MobilityFromEUTRACommand-v1530-IEs ::= SEQUENCE
type MobilityfromeutracommandV1530 struct {
	SmtcR15              *MtcSsbNrR15
	Noncriticalextension *MobilityfromeutracommandV1530IesNoncriticalextension
}
