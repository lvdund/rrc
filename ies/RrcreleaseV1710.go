package ies

import "rrc/utils"

// RRCRelease-v1710-IEs ::= SEQUENCE
type RrcreleaseV1710 struct {
	NolastcellupdateR17  *utils.BOOLEAN
	Noncriticalextension *RrcreleaseV1710IesNoncriticalextension
}
