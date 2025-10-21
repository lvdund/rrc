package ies

import "rrc/utils"

// RRCConnectionReject-v8a0-IEs ::= SEQUENCE
type RrcconnectionrejectV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionrejectV1020Ies
}
