package ies

import "rrc/utils"

// RRCConnectionReestablishmentComplete-v1130-IEs ::= SEQUENCE
type RrcconnectionreestablishmentcompleteV1130Ies struct {
	ConnestfailinfoavailableR11 *utils.ENUMERATED
	Noncriticalextension        *RrcconnectionreestablishmentcompleteV1250Ies
}
