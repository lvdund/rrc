package ies

import "rrc/utils"

// RRCConnectionReestablishment-v8a0-IEs ::= SEQUENCE
type RrcconnectionreestablishmentV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
