package ies

import "rrc/utils"

// RRCConnectionReject-v8a0-IEs ::= SEQUENCE
type RrcconnectionrejectV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionrejectV1020
}
