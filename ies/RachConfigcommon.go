package ies

import "rrc/utils"

// RACH-ConfigCommon ::= SEQUENCE
// Extensible
type RachConfigcommon struct {
	RachConfiggeneric                      RachConfiggeneric
	TotalnumberofraPreambles               *utils.INTEGER `lb:0,ub:63`
	SsbPerrachOccasionandcbPreamblesperssb *RachConfigcommonSsbPerrachOccasionandcbPreamblesperssb
	Groupbconfigured                       *RachConfigcommonGroupbconfigured
	RaContentionresolutiontimer            RachConfigcommonRaContentionresolutiontimer
	RsrpThresholdssb                       *RsrpRange
	RsrpThresholdssbSul                    *RsrpRange
	PrachRootsequenceindex                 RachConfigcommonPrachRootsequenceindex
	Msg1Subcarrierspacing                  *Subcarrierspacing
	Restrictedsetconfig                    RachConfigcommonRestrictedsetconfig
	Msg3Transformprecoder                  *RachConfigcommonMsg3Transformprecoder
	RaPrioritizationforaccessidentityR16   *RachConfigcommonRaPrioritizationforaccessidentityR16 `ext`
	PrachRootsequenceindexR16              *RachConfigcommonPrachRootsequenceindexR16            `ext`
	RaPrioritizationforslicingR17          *RaPrioritizationforslicingR17                        `ext`
	FeaturecombinationpreambleslistR17     *[]FeaturecombinationpreamblesR17                     `lb:1,ub:maxFeatureCombPreamblesPerRACHResourceR17,ext`
}
