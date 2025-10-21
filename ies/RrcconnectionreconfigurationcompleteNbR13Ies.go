package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-NB-r13-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteNbR13Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
