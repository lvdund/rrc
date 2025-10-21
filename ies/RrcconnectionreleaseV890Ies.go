package ies

import "rrc/utils"

// RRCConnectionRelease-v890-IEs ::= SEQUENCE
type RrcconnectionreleaseV890Ies struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreleaseV920Ies
}
