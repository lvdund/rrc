package ies

import "rrc/utils"

// NZP-CSI-RS-Pairing-r17 ::= SEQUENCE
type NzpCsiRsPairingR17 struct {
	NzpCsiRsResourceid1R17 utils.INTEGER `lb:0,ub:7`
	NzpCsiRsResourceid2R17 utils.INTEGER `lb:0,ub:7`
}
