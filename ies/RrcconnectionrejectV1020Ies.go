package ies

import "rrc/utils"

// RRCConnectionReject-v1020-IEs ::= SEQUENCE
type RrcconnectionrejectV1020Ies struct {
	ExtendedwaittimeR10  *utils.INTEGER
	Noncriticalextension *RrcconnectionrejectV1130Ies
}
