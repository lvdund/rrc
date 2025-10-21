package ies

import "rrc/utils"

// RRCConnectionReestablishmentReject-v8a0-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrejectV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
