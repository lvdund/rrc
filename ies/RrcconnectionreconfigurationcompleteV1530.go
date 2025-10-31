package ies

// RRCConnectionReconfigurationComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1530 struct {
	LogmeasavailablebtR15      *bool
	LogmeasavailablewlanR15    *bool
	FlightpathinfoavailableR15 *bool
	Noncriticalextension       *RrcconnectionreconfigurationcompleteV1530IesNoncriticalextension
}
