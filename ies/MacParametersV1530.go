package ies

import "rrc/utils"

// MAC-Parameters-v1530 ::= SEQUENCE
type MacParametersV1530 struct {
	MinProcTimelinesubslotR15  *interface{}
	SkipsubframeprocessingR15  *SkipsubframeprocessingR15
	EarlydataUpR15             *utils.ENUMERATED
	DormantscellstateR15       *utils.ENUMERATED
	DirectscellactivationR15   *utils.ENUMERATED
	DirectscellhibernationR15  *utils.ENUMERATED
	ExtendedlcidDuplicationR15 *utils.ENUMERATED
	SpsServingcellR15          *utils.ENUMERATED
}
