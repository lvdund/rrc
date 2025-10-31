package ies

import "rrc/utils"

// SL-V2X-TxPoolReportIdentity-r14 ::= utils.INTEGER (1..maxSL-PoolToMeasure-r14)
type SlV2xTxpoolreportidentityR14 struct {
	Value utils.INTEGER `lb:0,ub:maxSLPooltomeasureR14`
}
