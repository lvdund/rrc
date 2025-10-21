package ies

import "rrc/utils"

// RRCConnectionSetupComplete-v1690-IEs ::= SEQUENCE
type RrcconnectionsetupcompleteV1690Ies struct {
	UlRrcSegmentationR16 *utils.ENUMERATED
	Noncriticalextension *interface{}
}
