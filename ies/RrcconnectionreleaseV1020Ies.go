package ies

import "rrc/utils"

// RRCConnectionRelease-v1020-IEs ::= SEQUENCE
type RrcconnectionreleaseV1020Ies struct {
	ExtendedwaittimeR10  *utils.INTEGER
	Noncriticalextension *RrcconnectionreleaseV1320Ies
}
