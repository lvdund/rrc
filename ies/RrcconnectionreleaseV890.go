package ies

import "rrc/utils"

// RRCConnectionRelease-v890-IEs ::= SEQUENCE
type RrcconnectionreleaseV890 struct {
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *RrcconnectionreleaseV920
}
