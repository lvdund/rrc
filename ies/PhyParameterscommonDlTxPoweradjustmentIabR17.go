package ies

import "rrc/utils"

// Phy-ParametersCommon-dl-tx-PowerAdjustment-IAB-r17 ::= ENUMERATED
type PhyParameterscommonDlTxPoweradjustmentIabR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDlTxPoweradjustmentIabR17EnumeratedNothing = iota
	PhyParameterscommonDlTxPoweradjustmentIabR17EnumeratedSupported
)
