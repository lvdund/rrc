package ies

import "rrc/utils"

// RRCConnectionSetup-v8a0-IEs ::= SEQUENCE
type RrcconnectionsetupV8a0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionsetupV1610
}
