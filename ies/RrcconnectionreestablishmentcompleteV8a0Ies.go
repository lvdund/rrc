package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v8a0-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV8a0Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreestablishmentcompleteV1020Ies
}
