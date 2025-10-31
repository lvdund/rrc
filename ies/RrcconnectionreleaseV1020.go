package ies

import "rrc/utils"

// RRCConnectionRelease-v1020-IEs ::= SEQUENCE
type RrcconnectionreleaseV1020 struct {
	ExtendedwaittimeR10  *utils.INTEGER `lb:0,ub:1800`
	Noncriticalextension *RrcconnectionreleaseV1320
}
