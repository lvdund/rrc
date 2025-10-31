package ies

// RRCConnectionRelease-r8-IEs ::= SEQUENCE
type RrcconnectionreleaseR8 struct {
	Releasecause                Releasecause
	Redirectedcarrierinfo       *Redirectedcarrierinfo
	Idlemodemobilitycontrolinfo *Idlemodemobilitycontrolinfo
	Noncriticalextension        *RrcconnectionreleaseV890
}
