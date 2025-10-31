package ies

// RRCConnectionSetup-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupV1610 struct {
	DedicatedinfonasR16  *Dedicatedinfonas
	Noncriticalextension *RrcconnectionsetupV1610IesNoncriticalextension
}
