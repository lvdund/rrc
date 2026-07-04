// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RACH-ConfigCommon (line 12816).

var rACHConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-ConfigGeneric"},
		{Name: "totalNumberOfRA-Preambles", Optional: true},
		{Name: "ssb-perRACH-OccasionAndCB-PreamblesPerSSB", Optional: true},
		{Name: "groupBconfigured", Optional: true},
		{Name: "ra-ContentionResolutionTimer"},
		{Name: "rsrp-ThresholdSSB", Optional: true},
		{Name: "rsrp-ThresholdSSB-SUL", Optional: true},
		{Name: "prach-RootSequenceIndex"},
		{Name: "msg1-SubcarrierSpacing", Optional: true},
		{Name: "restrictedSetConfig"},
		{Name: "msg3-transformPrecoder", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var rACHConfigCommonTotalNumberOfRAPreamblesConstraints = per.Constrained(1, 63)

var rACH_ConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "oneEighth"},
		{Name: "oneFourth"},
		{Name: "oneHalf"},
		{Name: "one"},
		{Name: "two"},
		{Name: "four"},
		{Name: "eight"},
		{Name: "sixteen"},
	},
}

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf   = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One       = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two       = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Four      = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Eight     = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Sixteen   = 7
)

const (
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf8  = 0
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf16 = 1
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf24 = 2
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf32 = 3
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf40 = 4
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf48 = 5
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf56 = 6
	RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf64 = 7
)

var rACHConfigCommonRaContentionResolutionTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf8, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf16, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf24, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf32, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf40, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf48, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf56, RACH_ConfigCommon_Ra_ContentionResolutionTimer_Sf64},
}

var rACH_ConfigCommonPrachRootSequenceIndexConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l839"},
		{Name: "l139"},
	},
}

const (
	RACH_ConfigCommon_Prach_RootSequenceIndex_L839 = 0
	RACH_ConfigCommon_Prach_RootSequenceIndex_L139 = 1
)

const (
	RACH_ConfigCommon_RestrictedSetConfig_UnrestrictedSet    = 0
	RACH_ConfigCommon_RestrictedSetConfig_RestrictedSetTypeA = 1
	RACH_ConfigCommon_RestrictedSetConfig_RestrictedSetTypeB = 2
)

var rACHConfigCommonRestrictedSetConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_RestrictedSetConfig_UnrestrictedSet, RACH_ConfigCommon_RestrictedSetConfig_RestrictedSetTypeA, RACH_ConfigCommon_RestrictedSetConfig_RestrictedSetTypeB},
}

const (
	RACH_ConfigCommon_Msg3_TransformPrecoder_Enabled = 0
)

var rACHConfigCommonMsg3TransformPrecoderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Msg3_TransformPrecoder_Enabled},
}

var rACHConfigCommonExtPrachRootSequenceIndexR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l571"},
		{Name: "l1151"},
	},
}

const (
	RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L571  = 0
	RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L1151 = 1
)

var rACHConfigCommonExtFeatureCombinationPreamblesListR17Constraints = per.SizeRange(1, common.MaxFeatureCombPreamblesPerRACHResource_r17)

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N4  = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N8  = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N12 = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N16 = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N20 = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N24 = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N28 = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N32 = 7
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N36 = 8
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N40 = 9
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N44 = 10
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N48 = 11
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N52 = 12
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N56 = 13
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N60 = 14
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N64 = 15
)

var rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneEighthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N4, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N8, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N12, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N16, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N20, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N24, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N28, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N32, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N36, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N40, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N44, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N48, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N52, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N56, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N60, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth_N64},
}

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N4  = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N8  = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N12 = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N16 = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N20 = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N24 = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N28 = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N32 = 7
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N36 = 8
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N40 = 9
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N44 = 10
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N48 = 11
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N52 = 12
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N56 = 13
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N60 = 14
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N64 = 15
)

var rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneFourthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N4, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N8, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N12, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N16, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N20, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N24, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N28, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N32, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N36, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N40, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N44, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N48, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N52, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N56, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N60, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth_N64},
}

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N4  = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N8  = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N12 = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N16 = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N20 = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N24 = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N28 = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N32 = 7
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N36 = 8
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N40 = 9
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N44 = 10
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N48 = 11
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N52 = 12
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N56 = 13
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N60 = 14
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N64 = 15
)

var rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneHalfConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N4, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N8, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N12, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N16, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N20, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N24, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N28, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N32, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N36, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N40, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N44, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N48, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N52, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N56, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N60, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf_N64},
}

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N4  = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N8  = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N12 = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N16 = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N20 = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N24 = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N28 = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N32 = 7
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N36 = 8
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N40 = 9
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N44 = 10
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N48 = 11
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N52 = 12
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N56 = 13
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N60 = 14
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N64 = 15
)

var rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N4, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N8, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N12, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N16, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N20, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N24, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N28, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N32, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N36, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N40, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N44, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N48, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N52, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N56, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N60, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One_N64},
}

const (
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N4  = 0
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N8  = 1
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N12 = 2
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N16 = 3
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N20 = 4
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N24 = 5
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N28 = 6
	RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N32 = 7
)

var rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBTwoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N4, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N8, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N12, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N16, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N20, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N24, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N28, RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two_N32},
}

const (
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B56    = 0
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B144   = 1
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B208   = 2
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B256   = 3
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B282   = 4
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B480   = 5
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B640   = 6
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B800   = 7
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B1000  = 8
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B72    = 9
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare6 = 10
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare5 = 11
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare4 = 12
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare3 = 13
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare2 = 14
	RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare1 = 15
)

var rACHConfigCommonGroupBconfiguredRaMsg3SizeGroupAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B56, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B144, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B208, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B256, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B282, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B480, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B640, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B800, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B1000, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_B72, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare6, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare5, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare4, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare3, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare2, RACH_ConfigCommon_GroupBconfigured_Ra_Msg3SizeGroupA_Spare1},
}

const (
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_Minusinfinity = 0
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB0           = 1
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB5           = 2
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB8           = 3
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB10          = 4
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB12          = 5
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB15          = 6
	RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB18          = 7
)

var rACHConfigCommonGroupBconfiguredMessagePowerOffsetGroupBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_Minusinfinity, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB0, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB5, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB8, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB10, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB12, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB15, RACH_ConfigCommon_GroupBconfigured_MessagePowerOffsetGroupB_DB18},
}

type RACH_ConfigCommon struct {
	Rach_ConfigGeneric                        RACH_ConfigGeneric
	TotalNumberOfRA_Preambles                 *int64
	Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB *struct {
		Choice    int
		OneEighth *int64
		OneFourth *int64
		OneHalf   *int64
		One       *int64
		Two       *int64
		Four      *int64
		Eight     *int64
		Sixteen   *int64
	}
	GroupBconfigured *struct {
		Ra_Msg3SizeGroupA          int64
		MessagePowerOffsetGroupB   int64
		NumberOfRA_PreamblesGroupA int64
	}
	Ra_ContentionResolutionTimer int64
	Rsrp_ThresholdSSB            *RSRP_Range
	Rsrp_ThresholdSSB_SUL        *RSRP_Range
	Prach_RootSequenceIndex      struct {
		Choice int
		L839   *int64
		L139   *int64
	}
	Msg1_SubcarrierSpacing                 *SubcarrierSpacing
	RestrictedSetConfig                    int64
	Msg3_TransformPrecoder                 *int64
	Ra_PrioritizationForAccessIdentity_r16 *struct {
		Ra_Prioritization_r16      RA_Prioritization
		Ra_PrioritizationForAI_r16 per.BitString
	}
	Prach_RootSequenceIndex_r16 *struct {
		Choice int
		L571   *int64
		L1151  *int64
	}
	Ra_PrioritizationForSlicing_r17     *RA_PrioritizationForSlicing_r17
	FeatureCombinationPreamblesList_r17 []FeatureCombinationPreambles_r17
	Rach_ConfigAdapt_r19                *RA_AdaptConfig_r19
}

func (ie *RACH_ConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ra_PrioritizationForAccessIdentity_r16 != nil || ie.Prach_RootSequenceIndex_r16 != nil
	hasExtGroup1 := ie.Ra_PrioritizationForSlicing_r17 != nil || ie.FeatureCombinationPreamblesList_r17 != nil
	hasExtGroup2 := ie.Rach_ConfigAdapt_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TotalNumberOfRA_Preambles != nil, ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB != nil, ie.GroupBconfigured != nil, ie.Rsrp_ThresholdSSB != nil, ie.Rsrp_ThresholdSSB_SUL != nil, ie.Msg1_SubcarrierSpacing != nil, ie.Msg3_TransformPrecoder != nil}); err != nil {
		return err
	}

	// 3. rach-ConfigGeneric: ref
	{
		if err := ie.Rach_ConfigGeneric.Encode(e); err != nil {
			return err
		}
	}

	// 4. totalNumberOfRA-Preambles: integer
	{
		if ie.TotalNumberOfRA_Preambles != nil {
			if err := e.EncodeInteger(*ie.TotalNumberOfRA_Preambles, rACHConfigCommonTotalNumberOfRAPreamblesConstraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-perRACH-OccasionAndCB-PreamblesPerSSB: choice
	{
		if ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB != nil {
			choiceEnc := e.NewChoiceEncoder(rACH_ConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Choice {
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneEighth), rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneEighthConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneFourth), rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneFourthConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneHalf), rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneHalfConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).One), rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Two), rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBTwoConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Four:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Four), per.Constrained(1, 16)); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Eight:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Eight), per.Constrained(1, 8)); err != nil {
					return err
				}
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Sixteen:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Sixteen), per.Constrained(1, 4)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. groupBconfigured: sequence
	{
		if ie.GroupBconfigured != nil {
			c := ie.GroupBconfigured
			if err := e.EncodeEnumerated(c.Ra_Msg3SizeGroupA, rACHConfigCommonGroupBconfiguredRaMsg3SizeGroupAConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MessagePowerOffsetGroupB, rACHConfigCommonGroupBconfiguredMessagePowerOffsetGroupBConstraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.NumberOfRA_PreamblesGroupA, per.Constrained(1, 64)); err != nil {
				return err
			}
		}
	}

	// 7. ra-ContentionResolutionTimer: enumerated
	{
		if err := e.EncodeEnumerated(ie.Ra_ContentionResolutionTimer, rACHConfigCommonRaContentionResolutionTimerConstraints); err != nil {
			return err
		}
	}

	// 8. rsrp-ThresholdSSB: ref
	{
		if ie.Rsrp_ThresholdSSB != nil {
			if err := ie.Rsrp_ThresholdSSB.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. rsrp-ThresholdSSB-SUL: ref
	{
		if ie.Rsrp_ThresholdSSB_SUL != nil {
			if err := ie.Rsrp_ThresholdSSB_SUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. prach-RootSequenceIndex: choice
	{
		choiceEnc := e.NewChoiceEncoder(rACH_ConfigCommonPrachRootSequenceIndexConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Prach_RootSequenceIndex.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Prach_RootSequenceIndex.Choice {
		case RACH_ConfigCommon_Prach_RootSequenceIndex_L839:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex.L839), per.Constrained(0, 837)); err != nil {
				return err
			}
		case RACH_ConfigCommon_Prach_RootSequenceIndex_L139:
			if err := e.EncodeInteger((*ie.Prach_RootSequenceIndex.L139), per.Constrained(0, 137)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Prach_RootSequenceIndex.Choice), Constraint: "index out of range"}
		}
	}

	// 11. msg1-SubcarrierSpacing: ref
	{
		if ie.Msg1_SubcarrierSpacing != nil {
			if err := ie.Msg1_SubcarrierSpacing.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. restrictedSetConfig: enumerated
	{
		if err := e.EncodeEnumerated(ie.RestrictedSetConfig, rACHConfigCommonRestrictedSetConfigConstraints); err != nil {
			return err
		}
	}

	// 13. msg3-transformPrecoder: enumerated
	{
		if ie.Msg3_TransformPrecoder != nil {
			if err := e.EncodeEnumerated(*ie.Msg3_TransformPrecoder, rACHConfigCommonMsg3TransformPrecoderConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-PrioritizationForAccessIdentity-r16", Optional: true},
					{Name: "prach-RootSequenceIndex-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_PrioritizationForAccessIdentity_r16 != nil, ie.Prach_RootSequenceIndex_r16 != nil}); err != nil {
				return err
			}

			if ie.Ra_PrioritizationForAccessIdentity_r16 != nil {
				c := ie.Ra_PrioritizationForAccessIdentity_r16
				if err := c.Ra_Prioritization_r16.Encode(ex); err != nil {
					return err
				}
				if err := ex.EncodeBitString(c.Ra_PrioritizationForAI_r16, per.FixedSize(2)); err != nil {
					return err
				}
			}

			if ie.Prach_RootSequenceIndex_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(rACHConfigCommonExtPrachRootSequenceIndexR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Prach_RootSequenceIndex_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Prach_RootSequenceIndex_r16).Choice {
				case RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L571:
					if err := ex.EncodeInteger((*(*ie.Prach_RootSequenceIndex_r16).L571), per.Constrained(0, 569)); err != nil {
						return err
					}
				case RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L1151:
					if err := ex.EncodeInteger((*(*ie.Prach_RootSequenceIndex_r16).L1151), per.Constrained(0, 1149)); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-PrioritizationForSlicing-r17", Optional: true},
					{Name: "featureCombinationPreamblesList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_PrioritizationForSlicing_r17 != nil, ie.FeatureCombinationPreamblesList_r17 != nil}); err != nil {
				return err
			}

			if ie.Ra_PrioritizationForSlicing_r17 != nil {
				if err := ie.Ra_PrioritizationForSlicing_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.FeatureCombinationPreamblesList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(rACHConfigCommonExtFeatureCombinationPreamblesListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FeatureCombinationPreamblesList_r17))); err != nil {
					return err
				}
				for i := range ie.FeatureCombinationPreamblesList_r17 {
					if err := ie.FeatureCombinationPreamblesList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rach-ConfigAdapt-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rach_ConfigAdapt_r19 != nil}); err != nil {
				return err
			}

			if ie.Rach_ConfigAdapt_r19 != nil {
				if err := ie.Rach_ConfigAdapt_r19.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RACH_ConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rach-ConfigGeneric: ref
	{
		if err := ie.Rach_ConfigGeneric.Decode(d); err != nil {
			return err
		}
	}

	// 4. totalNumberOfRA-Preambles: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rACHConfigCommonTotalNumberOfRAPreamblesConstraints)
			if err != nil {
				return err
			}
			ie.TotalNumberOfRA_Preambles = &v
		}
	}

	// 5. ssb-perRACH-OccasionAndCB-PreamblesPerSSB: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB = &struct {
				Choice    int
				OneEighth *int64
				OneFourth *int64
				OneHalf   *int64
				One       *int64
				Two       *int64
				Four      *int64
				Eight     *int64
				Sixteen   *int64
			}{}
			choiceDec := d.NewChoiceDecoder(rACH_ConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneEighth:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneEighth = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneEighthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneEighth) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneFourth:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneFourth = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneFourthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneFourth) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_OneHalf:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneHalf = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneHalfConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).OneHalf) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_One:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).One = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBOneConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).One) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Two:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Two = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonSsbPerRACHOccasionAndCBPreamblesPerSSBTwoConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Two) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Four:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Four = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Four) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Eight:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Eight = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Eight) = v
			case RACH_ConfigCommon_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_Sixteen:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Sixteen = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB).Sixteen) = v
			}
		}
	}

	// 6. groupBconfigured: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.GroupBconfigured = &struct {
				Ra_Msg3SizeGroupA          int64
				MessagePowerOffsetGroupB   int64
				NumberOfRA_PreamblesGroupA int64
			}{}
			c := ie.GroupBconfigured
			{
				v, err := d.DecodeEnumerated(rACHConfigCommonGroupBconfiguredRaMsg3SizeGroupAConstraints)
				if err != nil {
					return err
				}
				c.Ra_Msg3SizeGroupA = v
			}
			{
				v, err := d.DecodeEnumerated(rACHConfigCommonGroupBconfiguredMessagePowerOffsetGroupBConstraints)
				if err != nil {
					return err
				}
				c.MessagePowerOffsetGroupB = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 64))
				if err != nil {
					return err
				}
				c.NumberOfRA_PreamblesGroupA = v
			}
		}
	}

	// 7. ra-ContentionResolutionTimer: enumerated
	{
		v4, err := d.DecodeEnumerated(rACHConfigCommonRaContentionResolutionTimerConstraints)
		if err != nil {
			return err
		}
		ie.Ra_ContentionResolutionTimer = v4
	}

	// 8. rsrp-ThresholdSSB: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Rsrp_ThresholdSSB = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdSSB.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. rsrp-ThresholdSSB-SUL: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Rsrp_ThresholdSSB_SUL = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdSSB_SUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. prach-RootSequenceIndex: choice
	{
		choiceDec := d.NewChoiceDecoder(rACH_ConfigCommonPrachRootSequenceIndexConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Prach_RootSequenceIndex.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case RACH_ConfigCommon_Prach_RootSequenceIndex_L839:
			ie.Prach_RootSequenceIndex.L839 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 837))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex.L839) = v
		case RACH_ConfigCommon_Prach_RootSequenceIndex_L139:
			ie.Prach_RootSequenceIndex.L139 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 137))
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex.L139) = v
		}
	}

	// 11. msg1-SubcarrierSpacing: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Msg1_SubcarrierSpacing = new(SubcarrierSpacing)
			if err := ie.Msg1_SubcarrierSpacing.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. restrictedSetConfig: enumerated
	{
		v9, err := d.DecodeEnumerated(rACHConfigCommonRestrictedSetConfigConstraints)
		if err != nil {
			return err
		}
		ie.RestrictedSetConfig = v9
	}

	// 13. msg3-transformPrecoder: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(rACHConfigCommonMsg3TransformPrecoderConstraints)
			if err != nil {
				return err
			}
			ie.Msg3_TransformPrecoder = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-PrioritizationForAccessIdentity-r16", Optional: true},
				{Name: "prach-RootSequenceIndex-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ra_PrioritizationForAccessIdentity_r16 = &struct {
				Ra_Prioritization_r16      RA_Prioritization
				Ra_PrioritizationForAI_r16 per.BitString
			}{}
			c := ie.Ra_PrioritizationForAccessIdentity_r16
			{
				if err := c.Ra_Prioritization_r16.Decode(dx); err != nil {
					return err
				}
			}
			{
				v, err := dx.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				c.Ra_PrioritizationForAI_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Prach_RootSequenceIndex_r16 = &struct {
				Choice int
				L571   *int64
				L1151  *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(rACHConfigCommonExtPrachRootSequenceIndexR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Prach_RootSequenceIndex_r16).Choice = int(index)
			switch index {
			case RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L571:
				(*ie.Prach_RootSequenceIndex_r16).L571 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 569))
				if err != nil {
					return err
				}
				(*(*ie.Prach_RootSequenceIndex_r16).L571) = v
			case RACH_ConfigCommon_Ext_Prach_RootSequenceIndex_r16_L1151:
				(*ie.Prach_RootSequenceIndex_r16).L1151 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1149))
				if err != nil {
					return err
				}
				(*(*ie.Prach_RootSequenceIndex_r16).L1151) = v
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-PrioritizationForSlicing-r17", Optional: true},
				{Name: "featureCombinationPreamblesList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ra_PrioritizationForSlicing_r17 = new(RA_PrioritizationForSlicing_r17)
			if err := ie.Ra_PrioritizationForSlicing_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(rACHConfigCommonExtFeatureCombinationPreamblesListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FeatureCombinationPreamblesList_r17 = make([]FeatureCombinationPreambles_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.FeatureCombinationPreamblesList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "rach-ConfigAdapt-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rach_ConfigAdapt_r19 = new(RA_AdaptConfig_r19)
			if err := ie.Rach_ConfigAdapt_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
