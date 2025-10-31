package ies

// RRCConnectionReestablishmentComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV1530 struct {
	LogmeasavailablebtR15      *bool
	LogmeasavailablewlanR15    *bool
	FlightpathinfoavailableR15 *bool
	Noncriticalextension       *RrcconnectionreestablishmentcompleteV1530IesNoncriticalextension
}
