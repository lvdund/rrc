package ies

import "rrc/utils"

// RRCConnectionReject-r8-IEs ::= SEQUENCE
type RrcconnectionrejectR8Ies struct {
	Waittime             utils.INTEGER
	Noncriticalextension *RrcconnectionrejectV8a0Ies
}
