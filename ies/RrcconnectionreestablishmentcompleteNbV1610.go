package ies

// RRCConnectionReestablishmentComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteNbV1610 struct {
	RlfInfoavailableR16  *bool
	AnrInfoavailableR16  *bool
	Noncriticalextension *RrcconnectionreestablishmentcompleteNbV1610IesNoncriticalextension
}
