package ies

import "rrc/utils"

// ZeroTxPowerCSI-RS-r12 ::= SEQUENCE
type ZerotxpowercsiRsR12 struct {
	ZerotxpowerresourceconfiglistR12 utils.BITSTRING
	ZerotxpowersubframeconfigR12     utils.INTEGER
}
