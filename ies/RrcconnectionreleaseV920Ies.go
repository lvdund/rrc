package ies

import "rrc/utils"

// RRCConnectionRelease-v920-IEs ::= SEQUENCE
// Extensible
type RrcconnectionreleaseV920Ies struct {
	CellinfolistR9       *interface{}
	Noncriticalextension *RrcconnectionreleaseV1020Ies
}
