package ies

import "rrc/utils"

// RRCConnectionReestablishment-v8a0-IEs ::= SEQUENCE
type RrcconnectionreestablishmentV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreestablishmentV8a0IesNoncriticalextension
}
