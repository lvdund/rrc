package ies

// RRCConnectionReconfigurationComplete-v1020-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1020 struct {
	RlfInfoavailableR10  *bool
	LogmeasavailableR10  *bool
	Noncriticalextension *RrcconnectionreconfigurationcompleteV1130
}
