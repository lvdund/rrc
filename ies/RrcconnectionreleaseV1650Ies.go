package ies

import "rrc/utils"

// RRCConnectionRelease-v1650-IEs ::= SEQUENCE
type RrcconnectionreleaseV1650Ies struct {
	MpspriorityindicationR16 *utils.ENUMERATED
	Noncriticalextension     *interface{}
}
