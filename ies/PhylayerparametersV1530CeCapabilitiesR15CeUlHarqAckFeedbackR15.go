package ies

import "rrc/utils"

// PhyLayerParameters-v1530-ce-Capabilities-r15-ce-UL-HARQ-ACK-Feedback-r15 ::= ENUMERATED
type PhylayerparametersV1530CeCapabilitiesR15CeUlHarqAckFeedbackR15 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1530CeCapabilitiesR15CeUlHarqAckFeedbackR15EnumeratedNothing = iota
	PhylayerparametersV1530CeCapabilitiesR15CeUlHarqAckFeedbackR15EnumeratedSupported
)
