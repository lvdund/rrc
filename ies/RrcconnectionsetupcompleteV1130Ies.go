package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1130-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1130Ies struct {
	ConnestfailinfoavailableR11 *utils.ENUMERATED
	Noncriticalextension        *RrcconnectionsetupcompleteV1250Ies
}
