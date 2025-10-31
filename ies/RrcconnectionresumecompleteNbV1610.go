package ies

// RRCConnectionResumeComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionresumecompleteNbV1610 struct {
	RlfInfoavailableR16  *bool
	AnrInfoavailableR16  *bool
	Noncriticalextension *RrcconnectionresumecompleteNbV1610IesNoncriticalextension
}
