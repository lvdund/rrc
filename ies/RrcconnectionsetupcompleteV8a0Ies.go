package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v8a0-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionsetupcompleteV1020Ies
}
