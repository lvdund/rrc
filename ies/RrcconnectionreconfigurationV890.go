package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v890-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV890 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationV920
}
