package ies

import "rrc/utils"

// RRCConnectionReestablishment-NB-v1430-IEs ::= SEQUENCE
type RrcconnectionreestablishmentNbV1430 struct {
	DlNasMac             *utils.BITSTRING `lb:16,ub:16`
	Noncriticalextension *RrcconnectionreestablishmentNbV1430IesNoncriticalextension
}
