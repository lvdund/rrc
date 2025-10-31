package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteNbR13 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationcompleteNbR13IesNoncriticalextension
}
