package ies

// MAC-Parameters-v1530 ::= SEQUENCE
type MacParametersV1530 struct {
	MinProcTimelinesubslotR15  *[]ProcessingtimelinesetR15 `lb:1,ub:3`
	SkipsubframeprocessingR15  *SkipsubframeprocessingR15
	EarlydataUpR15             *MacParametersV1530EarlydataUpR15
	DormantscellstateR15       *MacParametersV1530DormantscellstateR15
	DirectscellactivationR15   *MacParametersV1530DirectscellactivationR15
	DirectscellhibernationR15  *MacParametersV1530DirectscellhibernationR15
	ExtendedlcidDuplicationR15 *MacParametersV1530ExtendedlcidDuplicationR15
	SpsServingcellR15          *MacParametersV1530SpsServingcellR15
}
