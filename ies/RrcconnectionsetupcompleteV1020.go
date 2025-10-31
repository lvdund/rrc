package ies

// RRCConnectionSetupComplete-v1020-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1020 struct {
	GummeiTypeR10          *RrcconnectionsetupcompleteV1020IesGummeiTypeR10
	RlfInfoavailableR10    *bool
	LogmeasavailableR10    *bool
	RnSubframeconfigreqR10 *RrcconnectionsetupcompleteV1020IesRnSubframeconfigreqR10
	Noncriticalextension   *RrcconnectionsetupcompleteV1130
}
