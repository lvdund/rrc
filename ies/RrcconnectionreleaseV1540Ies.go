package ies

import "rrc/utils"

// RRCConnectionRelease-v1540-IEs ::= SEQUENCE
type RrcconnectionreleaseV1540Ies struct {
	Waittime             *utils.INTEGER
	Noncriticalextension *RrcconnectionreleaseV15b0Ies
}
