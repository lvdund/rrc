package ies

import "rrc/utils"

// PhyLayerParameters-v1250-pusch-FeedbackMode-r12 ::= ENUMERATED
type PhylayerparametersV1250PuschFeedbackmodeR12 struct {
	Value utils.ENUMERATED
}

const (
	PhylayerparametersV1250PuschFeedbackmodeR12EnumeratedNothing = iota
	PhylayerparametersV1250PuschFeedbackmodeR12EnumeratedSupported
)
