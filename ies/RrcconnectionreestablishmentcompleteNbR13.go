package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteNbR13 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreestablishmentcompleteNbV1470
}
