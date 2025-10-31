package ies

import "rrc/utils"

// FeatureSetUplink-v1720 ::= SEQUENCE
type FeaturesetuplinkV1720 struct {
	PucchRepetitionF01234RrcConfigR17               *FeaturesetuplinkV1720PucchRepetitionF01234RrcConfigR17
	PucchRepetitionF01234DynamicindicationR17       *FeaturesetuplinkV1720PucchRepetitionF01234DynamicindicationR17
	IntersubslotfreqhoppingPucchR17                 *FeaturesetuplinkV1720IntersubslotfreqhoppingPucchR17
	SemistaticharqAckCodebooksubSlotpucchR17        *FeaturesetuplinkV1720SemistaticharqAckCodebooksubSlotpucchR17
	PhyPrioritizationlowprioritydgHighprioritycgR17 *utils.INTEGER `lb:0,ub:16`
	PhyPrioritizationhighprioritydgLowprioritycgR17 *FeaturesetuplinkV1720PhyPrioritizationhighprioritydgLowprioritycgR17
	ExtendeddcLocationreportR17                     *FeaturesetuplinkV1720ExtendeddcLocationreportR17
}
