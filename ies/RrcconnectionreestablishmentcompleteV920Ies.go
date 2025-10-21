package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v920-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV920Ies struct {
	RlfInfoavailableR9   *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreestablishmentcompleteV8a0Ies
}
