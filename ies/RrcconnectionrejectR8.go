package ies

import "rrc/utils"

// RRCConnectionReject-r8-IEs ::= SEQUENCE
type RrcconnectionrejectR8 struct {
	Waittime             utils.INTEGER `lb:0,ub:16`
	Noncriticalextension *RrcconnectionrejectV8a0
}
