package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v15b0-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV15b0Ies struct {
	NolastcellupdateR15  *utils.ENUMERATED
	Noncriticalextension *RrcconnectionreleaseNbV1610Ies
}
