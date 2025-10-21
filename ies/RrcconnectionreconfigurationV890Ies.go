package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v890-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV890Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationV920Ies
}
