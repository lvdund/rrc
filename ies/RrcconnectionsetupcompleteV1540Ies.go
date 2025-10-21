package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1540-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1540Ies struct {
	GummeiTypeV1540      *utils.ENUMERATED
	GuamiTypeR15         *utils.ENUMERATED
	Noncriticalextension *RrcconnectionsetupcompleteV1610Ies
}
