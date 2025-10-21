package ies

import "rrc/utils"

// RRCConnectionReestablishment-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionreestablishmentNbV1430Ies struct {
	DlNasMac             *utils.BITSTRING
	Noncriticalextension *interface{}
}
