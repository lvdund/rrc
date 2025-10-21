package ies

import "rrc/utils"

// SL-TxPoolToReleaseListV2X-r14 ::= SEQUENCE OF SL-V2X-TxPoolIdentity-r14
// SIZE (1.. maxSL-V2X-TxPool-r14)
type SlTxpooltoreleaselistv2xR14 struct {
	Value utils.Sequence[SlV2xTxpoolidentityR14]
}
