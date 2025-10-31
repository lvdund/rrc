package ies

import "rrc/utils"

// RRCRelease-v1650-IEs ::= SEQUENCE
type RrcreleaseV1650 struct {
	MpspriorityindicationR16 *utils.BOOLEAN
	Noncriticalextension     *RrcreleaseV1710
}
