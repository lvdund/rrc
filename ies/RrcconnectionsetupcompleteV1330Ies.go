package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1330-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1330Ies struct {
	UeCeNeedulgapsR13    *utils.ENUMERATED
	Noncriticalextension *RrcconnectionsetupcompleteV1430Ies
}
