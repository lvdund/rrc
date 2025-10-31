package ies

import "rrc/utils"

// RRCSetupComplete-v1690-IEs ::= SEQUENCE
type RrcsetupcompleteV1690 struct {
	UlRrcSegmentationR16 *utils.BOOLEAN
	Noncriticalextension *RrcsetupcompleteV1700
}
