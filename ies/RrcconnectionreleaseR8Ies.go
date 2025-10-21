package ies

import "rrc/utils"

// RRCConnectionRelease-r8-IEs ::= SEQUENCE
type RrcconnectionreleaseR8Ies struct {
	Releasecause                Releasecause
	Redirectedcarrierinfo       *Redirectedcarrierinfo
	Idlemodemobilitycontrolinfo *Idlemodemobilitycontrolinfo
	Noncriticalextension        *RrcconnectionreleaseV890Ies
}
