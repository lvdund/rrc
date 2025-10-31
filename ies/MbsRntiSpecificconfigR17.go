package ies

import "rrc/utils"

// MBS-RNTI-SpecificConfig-r17 ::= SEQUENCE
type MbsRntiSpecificconfigR17 struct {
	MbsRntiSpecificconfigidR17      MbsRntiSpecificconfigidR17
	GroupcommonRntiR17              MbsRntiSpecificconfigR17GroupcommonRntiR17
	DrxConfigptmR17                 *utils.Setuprelease[DrxConfigptmR17]
	HarqFeedbackenablermulticastR17 *MbsRntiSpecificconfigR17HarqFeedbackenablermulticastR17
	HarqFeedbackoptionmulticastR17  *MbsRntiSpecificconfigR17HarqFeedbackoptionmulticastR17
	PdschAggregationfactorR17       *MbsRntiSpecificconfigR17PdschAggregationfactorR17
}
