package ies

import "rrc/utils"

// Phy-ParametersCommon-desired-ul-tx-PowerAdjustment-r17 ::= ENUMERATED
type PhyParameterscommonDesiredUlTxPoweradjustmentR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDesiredUlTxPoweradjustmentR17EnumeratedNothing = iota
	PhyParameterscommonDesiredUlTxPoweradjustmentR17EnumeratedSupported
)
