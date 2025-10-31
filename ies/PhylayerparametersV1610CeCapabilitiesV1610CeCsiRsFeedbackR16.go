package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-ce-CSI-RS-Feedback-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackR16EnumeratedSupported
)
