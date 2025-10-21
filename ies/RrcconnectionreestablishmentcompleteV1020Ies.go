package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v1020-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV1020Ies struct {
	LogmeasavailableR10  *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreestablishmentcompleteV1130Ies
}
