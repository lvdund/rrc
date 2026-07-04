// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RA-AdaptConfig-r19 (line 13060).

var rAAdaptConfigR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "prach-ConfigurationIndex-r19"},
		{Name: "msg1-FDM-r19"},
		{Name: "msg1-FrequencyStart-r19"},
		{Name: "prach-SubsetMaskIndexAdaptation-r19", Optional: true},
		{Name: "ssb-PerRACH-OccasionAndCB-PreamblesPerSSB-r19", Optional: true},
		{Name: "validityDurationForRACH-Adaptation-r19"},
		{Name: "valueKforAssociationPatternPeriodsForPRACH-SubsetMask-r19", Optional: true},
	},
}

var rAAdaptConfigR19PrachConfigurationIndexR19Constraints = per.Constrained(0, 255)

const (
	RA_AdaptConfig_r19_Msg1_FDM_r19_One   = 0
	RA_AdaptConfig_r19_Msg1_FDM_r19_Two   = 1
	RA_AdaptConfig_r19_Msg1_FDM_r19_Four  = 2
	RA_AdaptConfig_r19_Msg1_FDM_r19_Eight = 3
)

var rAAdaptConfigR19Msg1FDMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Msg1_FDM_r19_One, RA_AdaptConfig_r19_Msg1_FDM_r19_Two, RA_AdaptConfig_r19_Msg1_FDM_r19_Four, RA_AdaptConfig_r19_Msg1_FDM_r19_Eight},
}

var rAAdaptConfigR19Msg1FrequencyStartR19Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

const (
	RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Zero  = 0
	RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_One   = 1
	RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Two   = 2
	RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Three = 3
)

var rAAdaptConfigR19PrachSubsetMaskIndexAdaptationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Zero, RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_One, RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Two, RA_AdaptConfig_r19_Prach_SubsetMaskIndexAdaptation_r19_Three},
}

var rA_AdaptConfig_r19SsbPerRACHOccasionAndCBPreamblesPerSSBR19Constraints = per.ChoiceConstraints{
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
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf   = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One       = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two       = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Four      = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Eight     = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Sixteen   = 7
)

const (
	RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N2  = 0
	RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N4  = 1
	RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N8  = 2
	RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N16 = 3
)

var rAAdaptConfigR19ValidityDurationForRACHAdaptationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N2, RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N4, RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N8, RA_AdaptConfig_r19_ValidityDurationForRACH_Adaptation_r19_N16},
}

const (
	RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N2  = 0
	RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N4  = 1
	RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N8  = 2
	RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N16 = 3
)

var rAAdaptConfigR19ValueKforAssociationPatternPeriodsForPRACHSubsetMaskR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N2, RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N4, RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N8, RA_AdaptConfig_r19_ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19_N16},
}

const (
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N4  = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N8  = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N12 = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N16 = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N20 = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N24 = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N28 = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N32 = 7
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N36 = 8
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N40 = 9
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N44 = 10
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N48 = 11
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N52 = 12
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N56 = 13
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N60 = 14
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N64 = 15
)

var rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneEighthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N4, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N8, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N12, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N16, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N20, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N24, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N28, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N32, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N36, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N40, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N44, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N48, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N52, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N56, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N60, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth_N64},
}

const (
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N4  = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N8  = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N12 = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N16 = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N20 = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N24 = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N28 = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N32 = 7
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N36 = 8
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N40 = 9
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N44 = 10
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N48 = 11
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N52 = 12
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N56 = 13
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N60 = 14
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N64 = 15
)

var rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneFourthConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N4, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N8, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N12, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N16, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N20, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N24, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N28, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N32, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N36, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N40, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N44, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N48, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N52, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N56, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N60, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth_N64},
}

const (
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N4  = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N8  = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N12 = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N16 = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N20 = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N24 = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N28 = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N32 = 7
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N36 = 8
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N40 = 9
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N44 = 10
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N48 = 11
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N52 = 12
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N56 = 13
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N60 = 14
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N64 = 15
)

var rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneHalfConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N4, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N8, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N12, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N16, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N20, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N24, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N28, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N32, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N36, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N40, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N44, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N48, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N52, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N56, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N60, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf_N64},
}

const (
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N4  = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N8  = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N12 = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N16 = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N20 = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N24 = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N28 = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N32 = 7
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N36 = 8
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N40 = 9
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N44 = 10
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N48 = 11
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N52 = 12
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N56 = 13
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N60 = 14
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N64 = 15
)

var rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N4, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N8, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N12, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N16, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N20, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N24, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N28, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N32, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N36, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N40, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N44, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N48, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N52, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N56, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N60, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One_N64},
}

const (
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N4  = 0
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N8  = 1
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N12 = 2
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N16 = 3
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N20 = 4
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N24 = 5
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N28 = 6
	RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N32 = 7
)

var rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19TwoConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N4, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N8, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N12, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N16, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N20, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N24, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N28, RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two_N32},
}

type RA_AdaptConfig_r19 struct {
	Prach_ConfigurationIndex_r19                  int64
	Msg1_FDM_r19                                  int64
	Msg1_FrequencyStart_r19                       int64
	Prach_SubsetMaskIndexAdaptation_r19           *int64
	Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19 *struct {
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
	ValidityDurationForRACH_Adaptation_r19                    int64
	ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19 *int64
}

func (ie *RA_AdaptConfig_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAAdaptConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Prach_SubsetMaskIndexAdaptation_r19 != nil, ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19 != nil, ie.ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19 != nil}); err != nil {
		return err
	}

	// 3. prach-ConfigurationIndex-r19: integer
	{
		if err := e.EncodeInteger(ie.Prach_ConfigurationIndex_r19, rAAdaptConfigR19PrachConfigurationIndexR19Constraints); err != nil {
			return err
		}
	}

	// 4. msg1-FDM-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.Msg1_FDM_r19, rAAdaptConfigR19Msg1FDMR19Constraints); err != nil {
			return err
		}
	}

	// 5. msg1-FrequencyStart-r19: integer
	{
		if err := e.EncodeInteger(ie.Msg1_FrequencyStart_r19, rAAdaptConfigR19Msg1FrequencyStartR19Constraints); err != nil {
			return err
		}
	}

	// 6. prach-SubsetMaskIndexAdaptation-r19: enumerated
	{
		if ie.Prach_SubsetMaskIndexAdaptation_r19 != nil {
			if err := e.EncodeEnumerated(*ie.Prach_SubsetMaskIndexAdaptation_r19, rAAdaptConfigR19PrachSubsetMaskIndexAdaptationR19Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ssb-PerRACH-OccasionAndCB-PreamblesPerSSB-r19: choice
	{
		if ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19 != nil {
			choiceEnc := e.NewChoiceEncoder(rA_AdaptConfig_r19SsbPerRACHOccasionAndCBPreamblesPerSSBR19Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Choice {
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneEighth), rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneEighthConstraints); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneFourth), rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneFourthConstraints); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneHalf), rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneHalfConstraints); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).One), rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneConstraints); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two:
				if err := e.EncodeEnumerated((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Two), rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19TwoConstraints); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Four:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Four), per.Constrained(1, 16)); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Eight:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Eight), per.Constrained(1, 8)); err != nil {
					return err
				}
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Sixteen:
				if err := e.EncodeInteger((*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Sixteen), per.Constrained(1, 4)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 8. validityDurationForRACH-Adaptation-r19: enumerated
	{
		if err := e.EncodeEnumerated(ie.ValidityDurationForRACH_Adaptation_r19, rAAdaptConfigR19ValidityDurationForRACHAdaptationR19Constraints); err != nil {
			return err
		}
	}

	// 9. valueKforAssociationPatternPeriodsForPRACH-SubsetMask-r19: enumerated
	{
		if ie.ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19 != nil {
			if err := e.EncodeEnumerated(*ie.ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19, rAAdaptConfigR19ValueKforAssociationPatternPeriodsForPRACHSubsetMaskR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RA_AdaptConfig_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAAdaptConfigR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. prach-ConfigurationIndex-r19: integer
	{
		v0, err := d.DecodeInteger(rAAdaptConfigR19PrachConfigurationIndexR19Constraints)
		if err != nil {
			return err
		}
		ie.Prach_ConfigurationIndex_r19 = v0
	}

	// 4. msg1-FDM-r19: enumerated
	{
		v1, err := d.DecodeEnumerated(rAAdaptConfigR19Msg1FDMR19Constraints)
		if err != nil {
			return err
		}
		ie.Msg1_FDM_r19 = v1
	}

	// 5. msg1-FrequencyStart-r19: integer
	{
		v2, err := d.DecodeInteger(rAAdaptConfigR19Msg1FrequencyStartR19Constraints)
		if err != nil {
			return err
		}
		ie.Msg1_FrequencyStart_r19 = v2
	}

	// 6. prach-SubsetMaskIndexAdaptation-r19: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(rAAdaptConfigR19PrachSubsetMaskIndexAdaptationR19Constraints)
			if err != nil {
				return err
			}
			ie.Prach_SubsetMaskIndexAdaptation_r19 = &idx
		}
	}

	// 7. ssb-PerRACH-OccasionAndCB-PreamblesPerSSB-r19: choice
	{
		if seq.IsComponentPresent(4) {
			ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19 = &struct {
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
			choiceDec := d.NewChoiceDecoder(rA_AdaptConfig_r19SsbPerRACHOccasionAndCBPreamblesPerSSBR19Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneEighth:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneEighth = new(int64)
				v, err := d.DecodeEnumerated(rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneEighthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneEighth) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneFourth:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneFourth = new(int64)
				v, err := d.DecodeEnumerated(rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneFourthConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneFourth) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_OneHalf:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneHalf = new(int64)
				v, err := d.DecodeEnumerated(rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneHalfConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).OneHalf) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_One:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).One = new(int64)
				v, err := d.DecodeEnumerated(rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19OneConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).One) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Two:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Two = new(int64)
				v, err := d.DecodeEnumerated(rAAdaptConfigR19SsbPerRACHOccasionAndCBPreamblesPerSSBR19TwoConstraints)
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Two) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Four:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Four = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Four) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Eight:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Eight = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Eight) = v
			case RA_AdaptConfig_r19_Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19_Sixteen:
				(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Sixteen = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PerRACH_OccasionAndCB_PreamblesPerSSB_r19).Sixteen) = v
			}
		}
	}

	// 8. validityDurationForRACH-Adaptation-r19: enumerated
	{
		v5, err := d.DecodeEnumerated(rAAdaptConfigR19ValidityDurationForRACHAdaptationR19Constraints)
		if err != nil {
			return err
		}
		ie.ValidityDurationForRACH_Adaptation_r19 = v5
	}

	// 9. valueKforAssociationPatternPeriodsForPRACH-SubsetMask-r19: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(rAAdaptConfigR19ValueKforAssociationPatternPeriodsForPRACHSubsetMaskR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueKforAssociationPatternPeriodsForPRACH_SubsetMask_r19 = &idx
		}
	}

	return nil
}
