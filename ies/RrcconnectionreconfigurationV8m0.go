package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v8m0-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV8m0 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreconfigurationV10i0
}
