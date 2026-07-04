// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RACH-ConfigCommonTwoStepRA-r16 (line 12869).

var rACHConfigCommonTwoStepRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "rach-ConfigGenericTwoStepRA-r16"},
		{Name: "msgA-TotalNumberOfRA-Preambles-r16", Optional: true},
		{Name: "msgA-SSB-PerRACH-OccasionAndCB-PreamblesPerSSB-r16", Optional: true},
		{Name: "msgA-CB-PreamblesPerSSB-PerSharedRO-r16", Optional: true},
		{Name: "msgA-SSB-SharedRO-MaskIndex-r16", Optional: true},
		{Name: "groupB-ConfiguredTwoStepRA-r16", Optional: true},
		{Name: "msgA-PRACH-RootSequenceIndex-r16", Optional: true},
		{Name: "msgA-TransMax-r16", Optional: true},
		{Name: "msgA-RSRP-Threshold-r16", Optional: true},
		{Name: "msgA-RSRP-ThresholdSSB-r16", Optional: true},
		{Name: "msgA-SubcarrierSpacing-r16", Optional: true},
		{Name: "msgA-RestrictedSetConfig-r16", Optional: true},
		{Name: "ra-PrioritizationForAccessIdentityTwoStep-r16", Optional: true},
		{Name: "ra-ContentionResolutionTimer-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var rACHConfigCommonTwoStepRAR16MsgATotalNumberOfRAPreamblesR16Constraints = per.Constrained(1, 63)

var rACH_ConfigCommonTwoStepRA_r16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16Constraints = per.ChoiceConstraints{
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
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf   = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One       = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two       = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Four      = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Eight     = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Sixteen   = 7
)

var rACHConfigCommonTwoStepRAR16MsgACBPreamblesPerSSBPerSharedROR16Constraints = per.Constrained(1, 60)

var rACHConfigCommonTwoStepRAR16MsgASSBSharedROMaskIndexR16Constraints = per.Constrained(1, 15)

var rACH_ConfigCommonTwoStepRA_r16MsgAPRACHRootSequenceIndexR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "l839"},
		{Name: "l139"},
		{Name: "l571"},
		{Name: "l1151"},
	},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L839  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L139  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L571  = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L1151 = 3
)

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N1   = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N2   = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N4   = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N6   = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N8   = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N10  = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N20  = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N50  = 7
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N100 = 8
	RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N200 = 9
)

var rACHConfigCommonTwoStepRAR16MsgATransMaxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N1, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N2, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N6, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N10, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N50, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N100, RACH_ConfigCommonTwoStepRA_r16_MsgA_TransMax_r16_N200},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_UnrestrictedSet    = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_RestrictedSetTypeA = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_RestrictedSetTypeB = 2
)

var rACHConfigCommonTwoStepRAR16MsgARestrictedSetConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_UnrestrictedSet, RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_RestrictedSetTypeA, RACH_ConfigCommonTwoStepRA_r16_MsgA_RestrictedSetConfig_r16_RestrictedSetTypeB},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf8  = 0
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf16 = 1
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf24 = 2
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf32 = 3
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf40 = 4
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf48 = 5
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf56 = 6
	RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf64 = 7
)

var rACHConfigCommonTwoStepRAR16RaContentionResolutionTimerR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf8, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf16, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf24, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf32, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf40, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf48, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf56, RACH_ConfigCommonTwoStepRA_r16_Ra_ContentionResolutionTimer_r16_Sf64},
}

var rACHConfigCommonTwoStepRAR16ExtFeatureCombinationPreamblesListR17Constraints = per.SizeRange(1, common.MaxFeatureCombPreamblesPerRACHResource_r17)

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N4  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N8  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N12 = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N16 = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N20 = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N24 = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N28 = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N32 = 7
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N36 = 8
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N40 = 9
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N44 = 10
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N48 = 11
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N52 = 12
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N56 = 13
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N60 = 14
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N64 = 15
)

var rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneEighthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N12, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N16, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N24, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N28, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N32, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N36, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N40, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N44, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N48, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N52, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N56, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N60, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth_N64},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N4  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N8  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N12 = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N16 = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N20 = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N24 = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N28 = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N32 = 7
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N36 = 8
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N40 = 9
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N44 = 10
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N48 = 11
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N52 = 12
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N56 = 13
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N60 = 14
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N64 = 15
)

var rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneFourthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N12, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N16, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N24, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N28, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N32, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N36, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N40, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N44, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N48, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N52, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N56, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N60, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth_N64},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N4  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N8  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N12 = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N16 = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N20 = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N24 = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N28 = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N32 = 7
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N36 = 8
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N40 = 9
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N44 = 10
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N48 = 11
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N52 = 12
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N56 = 13
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N60 = 14
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N64 = 15
)

var rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneHalfConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N12, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N16, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N24, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N28, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N32, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N36, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N40, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N44, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N48, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N52, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N56, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N60, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf_N64},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N4  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N8  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N12 = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N16 = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N20 = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N24 = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N28 = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N32 = 7
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N36 = 8
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N40 = 9
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N44 = 10
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N48 = 11
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N52 = 12
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N56 = 13
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N60 = 14
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N64 = 15
)

var rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N12, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N16, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N24, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N28, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N32, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N36, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N40, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N44, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N48, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N52, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N56, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N60, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One_N64},
}

const (
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N4  = 0
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N8  = 1
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N12 = 2
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N16 = 3
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N20 = 4
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N24 = 5
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N28 = 6
	RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N32 = 7
)

var rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16TwoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N4, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N8, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N12, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N16, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N20, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N24, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N28, RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two_N32},
}

type RACH_ConfigCommonTwoStepRA_r16 struct {
	Rach_ConfigGenericTwoStepRA_r16                    RACH_ConfigGenericTwoStepRA_r16
	MsgA_TotalNumberOfRA_Preambles_r16                 *int64
	MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 *struct {
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
	MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 *int64
	MsgA_SSB_SharedRO_MaskIndex_r16         *int64
	GroupB_ConfiguredTwoStepRA_r16          *GroupB_ConfiguredTwoStepRA_r16
	MsgA_PRACH_RootSequenceIndex_r16        *struct {
		Choice int
		L839   *int64
		L139   *int64
		L571   *int64
		L1151  *int64
	}
	MsgA_TransMax_r16                             *int64
	MsgA_RSRP_Threshold_r16                       *RSRP_Range
	MsgA_RSRP_ThresholdSSB_r16                    *RSRP_Range
	MsgA_SubcarrierSpacing_r16                    *SubcarrierSpacing
	MsgA_RestrictedSetConfig_r16                  *int64
	Ra_PrioritizationForAccessIdentityTwoStep_r16 *struct {
		Ra_Prioritization_r16      RA_Prioritization
		Ra_PrioritizationForAI_r16 per.BitString
	}
	Ra_ContentionResolutionTimer_r16       *int64
	Ra_PrioritizationForSlicingTwoStep_r17 *RA_PrioritizationForSlicing_r17
	FeatureCombinationPreamblesList_r17    []FeatureCombinationPreambles_r17
}

func (ie *RACH_ConfigCommonTwoStepRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rACHConfigCommonTwoStepRAR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil || ie.FeatureCombinationPreamblesList_r17 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MsgA_TotalNumberOfRA_Preambles_r16 != nil, ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil, ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil, ie.MsgA_SSB_SharedRO_MaskIndex_r16 != nil, ie.GroupB_ConfiguredTwoStepRA_r16 != nil, ie.MsgA_PRACH_RootSequenceIndex_r16 != nil, ie.MsgA_TransMax_r16 != nil, ie.MsgA_RSRP_Threshold_r16 != nil, ie.MsgA_RSRP_ThresholdSSB_r16 != nil, ie.MsgA_SubcarrierSpacing_r16 != nil, ie.MsgA_RestrictedSetConfig_r16 != nil, ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 != nil, ie.Ra_ContentionResolutionTimer_r16 != nil}); err != nil {
		return err
	}

	// 3. rach-ConfigGenericTwoStepRA-r16: ref
	{
		if err := ie.Rach_ConfigGenericTwoStepRA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. msgA-TotalNumberOfRA-Preambles-r16: integer
	{
		if ie.MsgA_TotalNumberOfRA_Preambles_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_TotalNumberOfRA_Preambles_r16, rACHConfigCommonTwoStepRAR16MsgATotalNumberOfRAPreamblesR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. msgA-SSB-PerRACH-OccasionAndCB-PreamblesPerSSB-r16: choice
	{
		if ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rACH_ConfigCommonTwoStepRA_r16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Choice {
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth:
				if err := e.EncodeEnumerated((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneEighth), rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneEighthConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth:
				if err := e.EncodeEnumerated((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneFourth), rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneFourthConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf:
				if err := e.EncodeEnumerated((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneHalf), rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneHalfConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One:
				if err := e.EncodeEnumerated((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).One), rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two:
				if err := e.EncodeEnumerated((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Two), rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16TwoConstraints); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Four:
				if err := e.EncodeInteger((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Four), per.Constrained(1, 16)); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Eight:
				if err := e.EncodeInteger((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Eight), per.Constrained(1, 8)); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Sixteen:
				if err := e.EncodeInteger((*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Sixteen), per.Constrained(1, 4)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. msgA-CB-PreamblesPerSSB-PerSharedRO-r16: integer
	{
		if ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16, rACHConfigCommonTwoStepRAR16MsgACBPreamblesPerSSBPerSharedROR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. msgA-SSB-SharedRO-MaskIndex-r16: integer
	{
		if ie.MsgA_SSB_SharedRO_MaskIndex_r16 != nil {
			if err := e.EncodeInteger(*ie.MsgA_SSB_SharedRO_MaskIndex_r16, rACHConfigCommonTwoStepRAR16MsgASSBSharedROMaskIndexR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. groupB-ConfiguredTwoStepRA-r16: ref
	{
		if ie.GroupB_ConfiguredTwoStepRA_r16 != nil {
			if err := ie.GroupB_ConfiguredTwoStepRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. msgA-PRACH-RootSequenceIndex-r16: choice
	{
		if ie.MsgA_PRACH_RootSequenceIndex_r16 != nil {
			choiceEnc := e.NewChoiceEncoder(rACH_ConfigCommonTwoStepRA_r16MsgAPRACHRootSequenceIndexR16Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.MsgA_PRACH_RootSequenceIndex_r16).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.MsgA_PRACH_RootSequenceIndex_r16).Choice {
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L839:
				if err := e.EncodeInteger((*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L839), per.Constrained(0, 837)); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L139:
				if err := e.EncodeInteger((*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L139), per.Constrained(0, 137)); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L571:
				if err := e.EncodeInteger((*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L571), per.Constrained(0, 569)); err != nil {
					return err
				}
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L1151:
				if err := e.EncodeInteger((*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L1151), per.Constrained(0, 1149)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.MsgA_PRACH_RootSequenceIndex_r16).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 10. msgA-TransMax-r16: enumerated
	{
		if ie.MsgA_TransMax_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_TransMax_r16, rACHConfigCommonTwoStepRAR16MsgATransMaxR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. msgA-RSRP-Threshold-r16: ref
	{
		if ie.MsgA_RSRP_Threshold_r16 != nil {
			if err := ie.MsgA_RSRP_Threshold_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. msgA-RSRP-ThresholdSSB-r16: ref
	{
		if ie.MsgA_RSRP_ThresholdSSB_r16 != nil {
			if err := ie.MsgA_RSRP_ThresholdSSB_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 13. msgA-SubcarrierSpacing-r16: ref
	{
		if ie.MsgA_SubcarrierSpacing_r16 != nil {
			if err := ie.MsgA_SubcarrierSpacing_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. msgA-RestrictedSetConfig-r16: enumerated
	{
		if ie.MsgA_RestrictedSetConfig_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MsgA_RestrictedSetConfig_r16, rACHConfigCommonTwoStepRAR16MsgARestrictedSetConfigR16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. ra-PrioritizationForAccessIdentityTwoStep-r16: sequence
	{
		if ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 != nil {
			c := ie.Ra_PrioritizationForAccessIdentityTwoStep_r16
			if err := c.Ra_Prioritization_r16.Encode(e); err != nil {
				return err
			}
			if err := e.EncodeBitString(c.Ra_PrioritizationForAI_r16, per.FixedSize(2)); err != nil {
				return err
			}
		}
	}

	// 16. ra-ContentionResolutionTimer-r16: enumerated
	{
		if ie.Ra_ContentionResolutionTimer_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ra_ContentionResolutionTimer_r16, rACHConfigCommonTwoStepRAR16RaContentionResolutionTimerR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ra-PrioritizationForSlicingTwoStep-r17", Optional: true},
					{Name: "featureCombinationPreamblesList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil, ie.FeatureCombinationPreamblesList_r17 != nil}); err != nil {
				return err
			}

			if ie.Ra_PrioritizationForSlicingTwoStep_r17 != nil {
				if err := ie.Ra_PrioritizationForSlicingTwoStep_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.FeatureCombinationPreamblesList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(rACHConfigCommonTwoStepRAR16ExtFeatureCombinationPreamblesListR17Constraints)
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

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rACHConfigCommonTwoStepRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. rach-ConfigGenericTwoStepRA-r16: ref
	{
		if err := ie.Rach_ConfigGenericTwoStepRA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. msgA-TotalNumberOfRA-Preambles-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rACHConfigCommonTwoStepRAR16MsgATotalNumberOfRAPreamblesR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_TotalNumberOfRA_Preambles_r16 = &v
		}
	}

	// 5. msgA-SSB-PerRACH-OccasionAndCB-PreamblesPerSSB-r16: choice
	{
		if seq.IsComponentPresent(2) {
			ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 = &struct {
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
			choiceDec := d.NewChoiceDecoder(rACH_ConfigCommonTwoStepRA_r16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneEighth:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneEighth = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneEighthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneEighth) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneFourth:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneFourth = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneFourthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneFourth) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_OneHalf:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneHalf = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneHalfConstraints)
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).OneHalf) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_One:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).One = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16OneConstraints)
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).One) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Two:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Two = new(int64)
				v, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgASSBPerRACHOccasionAndCBPreamblesPerSSBR16TwoConstraints)
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Two) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Four:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Four = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Four) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Eight:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Eight = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Eight) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Sixteen:
				(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Sixteen = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16).Sixteen) = v
			}
		}
	}

	// 6. msgA-CB-PreamblesPerSSB-PerSharedRO-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(rACHConfigCommonTwoStepRAR16MsgACBPreamblesPerSSBPerSharedROR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_CB_PreamblesPerSSB_PerSharedRO_r16 = &v
		}
	}

	// 7. msgA-SSB-SharedRO-MaskIndex-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(rACHConfigCommonTwoStepRAR16MsgASSBSharedROMaskIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_SSB_SharedRO_MaskIndex_r16 = &v
		}
	}

	// 8. groupB-ConfiguredTwoStepRA-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.GroupB_ConfiguredTwoStepRA_r16 = new(GroupB_ConfiguredTwoStepRA_r16)
			if err := ie.GroupB_ConfiguredTwoStepRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. msgA-PRACH-RootSequenceIndex-r16: choice
	{
		if seq.IsComponentPresent(6) {
			ie.MsgA_PRACH_RootSequenceIndex_r16 = &struct {
				Choice int
				L839   *int64
				L139   *int64
				L571   *int64
				L1151  *int64
			}{}
			choiceDec := d.NewChoiceDecoder(rACH_ConfigCommonTwoStepRA_r16MsgAPRACHRootSequenceIndexR16Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MsgA_PRACH_RootSequenceIndex_r16).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L839:
				(*ie.MsgA_PRACH_RootSequenceIndex_r16).L839 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 837))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L839) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L139:
				(*ie.MsgA_PRACH_RootSequenceIndex_r16).L139 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 137))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L139) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L571:
				(*ie.MsgA_PRACH_RootSequenceIndex_r16).L571 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 569))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L571) = v
			case RACH_ConfigCommonTwoStepRA_r16_MsgA_PRACH_RootSequenceIndex_r16_L1151:
				(*ie.MsgA_PRACH_RootSequenceIndex_r16).L1151 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1149))
				if err != nil {
					return err
				}
				(*(*ie.MsgA_PRACH_RootSequenceIndex_r16).L1151) = v
			}
		}
	}

	// 10. msgA-TransMax-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgATransMaxR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_TransMax_r16 = &idx
		}
	}

	// 11. msgA-RSRP-Threshold-r16: ref
	{
		if seq.IsComponentPresent(8) {
			ie.MsgA_RSRP_Threshold_r16 = new(RSRP_Range)
			if err := ie.MsgA_RSRP_Threshold_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. msgA-RSRP-ThresholdSSB-r16: ref
	{
		if seq.IsComponentPresent(9) {
			ie.MsgA_RSRP_ThresholdSSB_r16 = new(RSRP_Range)
			if err := ie.MsgA_RSRP_ThresholdSSB_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 13. msgA-SubcarrierSpacing-r16: ref
	{
		if seq.IsComponentPresent(10) {
			ie.MsgA_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
			if err := ie.MsgA_SubcarrierSpacing_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. msgA-RestrictedSetConfig-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16MsgARestrictedSetConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RestrictedSetConfig_r16 = &idx
		}
	}

	// 15. ra-PrioritizationForAccessIdentityTwoStep-r16: sequence
	{
		if seq.IsComponentPresent(12) {
			ie.Ra_PrioritizationForAccessIdentityTwoStep_r16 = &struct {
				Ra_Prioritization_r16      RA_Prioritization
				Ra_PrioritizationForAI_r16 per.BitString
			}{}
			c := ie.Ra_PrioritizationForAccessIdentityTwoStep_r16
			{
				if err := c.Ra_Prioritization_r16.Decode(d); err != nil {
					return err
				}
			}
			{
				v, err := d.DecodeBitString(per.FixedSize(2))
				if err != nil {
					return err
				}
				c.Ra_PrioritizationForAI_r16 = v
			}
		}
	}

	// 16. ra-ContentionResolutionTimer-r16: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(rACHConfigCommonTwoStepRAR16RaContentionResolutionTimerR16Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ContentionResolutionTimer_r16 = &idx
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
				{Name: "ra-PrioritizationForSlicingTwoStep-r17", Optional: true},
				{Name: "featureCombinationPreamblesList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ra_PrioritizationForSlicingTwoStep_r17 = new(RA_PrioritizationForSlicing_r17)
			if err := ie.Ra_PrioritizationForSlicingTwoStep_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(rACHConfigCommonTwoStepRAR16ExtFeatureCombinationPreamblesListR17Constraints)
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

	return nil
}
