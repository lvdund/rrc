package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v8a0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationcompleteV1020Ies
}
