package ies

import "rrc/utils"

// RACH-ConfigCommonTwoStepRA-r16 ::= SEQUENCE
// Extensible
type RachConfigcommontwostepraR16 struct {
	RachConfiggenerictwostepraR16                 RachConfiggenerictwostepraR16
	MsgaTotalnumberofraPreamblesR16               *utils.INTEGER `lb:0,ub:63`
	MsgaSsbPerrachOccasionandcbPreamblesperssbR16 *RachConfigcommontwostepraR16MsgaSsbPerrachOccasionandcbPreamblesperssbR16
	MsgaCbPreamblesperssbPersharedroR16           *utils.INTEGER `lb:0,ub:60`
	MsgaSsbSharedroMaskindexR16                   *utils.INTEGER `lb:0,ub:15`
	GroupbConfiguredtwostepraR16                  *GroupbConfiguredtwostepraR16
	MsgaPrachRootsequenceindexR16                 *RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16
	MsgaTransmaxR16                               *RachConfigcommontwostepraR16MsgaTransmaxR16
	MsgaRsrpThresholdR16                          *RsrpRange
	MsgaRsrpThresholdssbR16                       *RsrpRange
	MsgaSubcarrierspacingR16                      *Subcarrierspacing
	MsgaRestrictedsetconfigR16                    *RachConfigcommontwostepraR16MsgaRestrictedsetconfigR16
	RaPrioritizationforaccessidentitytwostepR16   *RachConfigcommontwostepraR16RaPrioritizationforaccessidentitytwostepR16
	RaContentionresolutiontimerR16                *RachConfigcommontwostepraR16RaContentionresolutiontimerR16
	RaPrioritizationforslicingtwostepR17          *RaPrioritizationforslicingR17    `ext`
	FeaturecombinationpreambleslistR17            *[]FeaturecombinationpreamblesR17 `lb:1,ub:maxFeatureCombPreamblesPerRACHResourceR17,ext`
}
