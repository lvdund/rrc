package ies

import "rrc/utils"

// RRCConnectionRelease-v15b0-IEs ::= SEQUENCE
type RrcconnectionreleaseV15b0Ies struct {
	NolastcellupdateR15  *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreleaseV1610Ies
}
