package ies

import "rrc/utils"

// ZeroTxPowerCSI-RS-r12 ::= SEQUENCE
type ZerotxpowercsiRsR12 struct {
	ZerotxpowerresourceconfiglistR12 utils.BITSTRING `lb:16,ub:16`
	ZerotxpowersubframeconfigR12     utils.INTEGER   `lb:0,ub:154`
}
