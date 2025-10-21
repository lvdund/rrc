package ies

import "rrc/utils"

// RRCConnectionSetup-v8a0-IEs ::= SEQUENCE
type RrcconnectionsetupV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionsetupV1610Ies
}
