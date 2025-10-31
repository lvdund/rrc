package ies

import "rrc/utils"

// Phy-ParametersCommon-updated-T-DeltaRangeRecption-r17 ::= ENUMERATED
type PhyParameterscommonUpdatedTDeltarangerecptionR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonUpdatedTDeltarangerecptionR17EnumeratedNothing = iota
	PhyParameterscommonUpdatedTDeltarangerecptionR17EnumeratedSupported
)
