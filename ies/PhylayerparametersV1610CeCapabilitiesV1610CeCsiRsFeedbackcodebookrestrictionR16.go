package ies

import "rrc/utils"

// PhyLayerParameters-v1610-ce-Capabilities-v1610-ce-CSI-RS-FeedbackCodebookRestriction-r16 ::= ENUMERATED
type PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackcodebookrestrictionR16 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackcodebookrestrictionR16EnumeratedNothing = iota
	PhylayerparametersV1610CeCapabilitiesV1610CeCsiRsFeedbackcodebookrestrictionR16EnumeratedSupported
)
