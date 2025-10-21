package ies

import "rrc/utils"

// Tx-ResourcePoolMeasList-r14 ::= SEQUENCE OF SL-V2X-TxPoolReportIdentity-r14
// SIZE (1..maxSL-PoolToMeasure-r14)
type TxResourcepoolmeaslistR14 struct {
	Value []SlV2xTxpoolreportidentityR14
}
