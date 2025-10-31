package ies

import "rrc/utils"

// RRCConnectionRelease-v1540-IEs ::= SEQUENCE
type RrcconnectionreleaseV1540 struct {
	Waittime             *utils.INTEGER `lb:0,ub:16`
	Noncriticalextension *RrcconnectionreleaseV15b0
}
