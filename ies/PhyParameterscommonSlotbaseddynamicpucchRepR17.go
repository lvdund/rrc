package ies

import "rrc/utils"

// Phy-ParametersCommon-slotBasedDynamicPUCCH-Rep-r17 ::= ENUMERATED
type PhyParameterscommonSlotbaseddynamicpucchRepR17 struct {
	Value utils.ENUMERATED
}

const (
	PhyParameterscommonSlotbaseddynamicpucchRepR17EnumeratedNothing = iota
	PhyParameterscommonSlotbaseddynamicpucchRepR17EnumeratedSupported
)
