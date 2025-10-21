package ies

import "rrc/utils"

// RRCConnectionReconfiguration-v920-IEs ::= SEQUENCE
type RrcconnectionreconfigurationV920Ies struct {
	OtherconfigR9        *OtherconfigR9
	FullconfigR9         *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreconfigurationV1020Ies
}
