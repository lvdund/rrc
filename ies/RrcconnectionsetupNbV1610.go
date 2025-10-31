package ies

// RRCConnectionSetup-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionsetupNbV1610 struct {
	DedicatedinfonasR16  *Dedicatedinfonas
	Noncriticalextension *RrcconnectionsetupNbV1610IesNoncriticalextension
}
