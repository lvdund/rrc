package ies

import "rrc/utils"

// MBS-RNTI-SpecificConfig-r17-harq-FeedbackOptionMulticast-r17 ::= ENUMERATED
type MbsRntiSpecificconfigR17HarqFeedbackoptionmulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	MbsRntiSpecificconfigR17HarqFeedbackoptionmulticastR17EnumeratedNothing = iota
	MbsRntiSpecificconfigR17HarqFeedbackoptionmulticastR17EnumeratedAck_Nack
	MbsRntiSpecificconfigR17HarqFeedbackoptionmulticastR17EnumeratedNack_Only
)
