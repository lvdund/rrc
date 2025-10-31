package ies

import "rrc/utils"

// RRCConnectionReject-v1020-IEs ::= SEQUENCE
type RrcconnectionrejectV1020 struct {
	ExtendedwaittimeR10  *utils.INTEGER `lb:0,ub:1800`
	Noncriticalextension *RrcconnectionrejectV1130
}
