package ies

import "rrc/utils"

// RRCConnectionReestablishmentReject-v8a0-IEs ::= SEQUENCE
type RrcconnectionreestablishmentrejectV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreestablishmentrejectV8a0IesNoncriticalextension
}
