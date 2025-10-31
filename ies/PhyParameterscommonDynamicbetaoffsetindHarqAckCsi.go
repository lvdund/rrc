package ies

import "rrc/utils"

// Phy-ParametersCommon-dynamicBetaOffsetInd-HARQ-ACK-CSI ::= ENUMERATED
type PhyParameterscommonDynamicbetaoffsetindHarqAckCsi struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonDynamicbetaoffsetindHarqAckCsiEnumeratedNothing = iota
	PhyParameterscommonDynamicbetaoffsetindHarqAckCsiEnumeratedSupported
)
