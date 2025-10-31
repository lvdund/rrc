package ies

// RRCConnectionResumeComplete-v1530-IEs ::= SEQUENCE
type RrcconnectionresumecompleteV1530 struct {
	LogmeasavailablebtR15      *bool
	LogmeasavailablewlanR15    *bool
	IdlemeasavailableR15       *bool
	FlightpathinfoavailableR15 *bool
	Noncriticalextension       *RrcconnectionresumecompleteV1610
}
