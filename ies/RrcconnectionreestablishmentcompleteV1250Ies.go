package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v1250-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV1250Ies struct {
	LogmeasavailablembsfnR12 *utils.ENUMERATED
	Noncriticalextension     *RrcconnectionreestablishmentcompleteV1530Ies
}
