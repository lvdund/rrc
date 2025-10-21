package ies

import "rrc/utils"

// RRCConnectionReconfigurationComplete-v1130-IEs ::= SEQUENCE
type RrcconnectionreconfigurationcompleteV1130Ies struct {
	ConnestfailinfoavailableR11 *utils.ENUMERATED
	Noncriticalextension        *RrcconnectionreconfigurationcompleteV1250Ies
}
