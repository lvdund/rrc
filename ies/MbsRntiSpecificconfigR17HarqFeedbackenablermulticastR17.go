package ies

import "rrc/utils"

// MBS-RNTI-SpecificConfig-r17-harq-FeedbackEnablerMulticast-r17 ::= ENUMERATED
type MbsRntiSpecificconfigR17HarqFeedbackenablermulticastR17 struct {
	Value utils.ENUMERATED
}

const (
	MbsRntiSpecificconfigR17HarqFeedbackenablermulticastR17EnumeratedNothing = iota
	MbsRntiSpecificconfigR17HarqFeedbackenablermulticastR17EnumeratedDci_Enabler
	MbsRntiSpecificconfigR17HarqFeedbackenablermulticastR17EnumeratedEnabled
)
