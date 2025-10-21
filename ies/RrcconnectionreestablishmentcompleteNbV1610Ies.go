package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-NB-v1610-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteNbV1610Ies struct {
	RlfInfoavailableR16  *utils.ENUMERATED
	AnrInfoavailableR16  *utils.ENUMERATED
	Noncriticalextension *interface{}
}
