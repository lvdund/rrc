// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BWP-UplinkCommon (line 5328).

var bWPUplinkCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "genericParameters"},
		{Name: "rach-ConfigCommon", Optional: true},
		{Name: "pusch-ConfigCommon", Optional: true},
		{Name: "pucch-ConfigCommon", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var bWP_UplinkCommonRachConfigCommonConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Rach_ConfigCommon_Release = 0
	BWP_UplinkCommon_Rach_ConfigCommon_Setup   = 1
)

var bWP_UplinkCommonPuschConfigCommonConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Pusch_ConfigCommon_Release = 0
	BWP_UplinkCommon_Pusch_ConfigCommon_Setup   = 1
)

var bWP_UplinkCommonPucchConfigCommonConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Pucch_ConfigCommon_Release = 0
	BWP_UplinkCommon_Pucch_ConfigCommon_Setup   = 1
)

var bWPUplinkCommonExtRachConfigCommonIABR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Release = 0
	BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Setup   = 1
)

const (
	BWP_UplinkCommon_Ext_UseInterlacePUCCH_PUSCH_r16_Enabled = 0
)

var bWPUplinkCommonExtUseInterlacePUCCHPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkCommon_Ext_UseInterlacePUCCH_PUSCH_r16_Enabled},
}

var bWPUplinkCommonExtMsgAConfigCommonR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Release = 0
	BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Setup   = 1
)

var bWPUplinkCommonExtAdditionalRACHConfigListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Release = 0
	BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Setup   = 1
)

var bWPUplinkCommonExtNumberOfMsg3RepetitionsListR17Constraints = per.FixedSize(4)

var bWPUplinkCommonExtMcsMsg3RepetitionsR17Constraints = per.FixedSize(8)

var bWPUplinkCommonExtAdditionalRACHPerPCIToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofAdditionalPRACHConfigs_r18)

var bWPUplinkCommonExtAdditionalRACHPerPCIToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofAdditionalPRACHConfigs_r18)

const (
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N1   = 0
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N2   = 1
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N4   = 2
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N6   = 3
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N8   = 4
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N10  = 5
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N20  = 6
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N50  = 7
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N100 = 8
	BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N200 = 9
)

var bWPUplinkCommonExtPreambleTransMaxMsg1RepetitionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N1, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N2, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N4, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N6, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N8, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N10, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N20, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N50, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N100, BWP_UplinkCommon_Ext_PreambleTransMax_Msg1_Repetition_r18_N200},
}

const (
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N1   = 0
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N2   = 1
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N4   = 2
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N6   = 3
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N8   = 4
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N10  = 5
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N20  = 6
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N50  = 7
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N100 = 8
	BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N200 = 9
)

var bWPUplinkCommonExtPreambleTransMaxROTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N1, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N2, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N4, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N6, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N8, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N10, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N20, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N50, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N100, BWP_UplinkCommon_Ext_PreambleTransMaxRO_Type_r19_N200},
}

const (
	BWP_UplinkCommon_Ext_Sbfd_RO_Type_r19_Sbfd     = 0
	BWP_UplinkCommon_Ext_Sbfd_RO_Type_r19_Non_Sbfd = 1
)

var bWPUplinkCommonExtSbfdROTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkCommon_Ext_Sbfd_RO_Type_r19_Sbfd, BWP_UplinkCommon_Ext_Sbfd_RO_Type_r19_Non_Sbfd},
}

const (
	BWP_UplinkCommon_Ext_Sbfd_RSRP_ThresholdRO_TypeUsage_r19_Above = 0
	BWP_UplinkCommon_Ext_Sbfd_RSRP_ThresholdRO_TypeUsage_r19_Below = 1
)

var bWPUplinkCommonExtSbfdRSRPThresholdROTypeUsageR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BWP_UplinkCommon_Ext_Sbfd_RSRP_ThresholdRO_TypeUsage_r19_Above, BWP_UplinkCommon_Ext_Sbfd_RSRP_ThresholdRO_TypeUsage_r19_Below},
}

var bWPUplinkCommonExtSbfdRACHConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "sbfd-RACH-SingleConfig-r19"},
		{Name: "sbfd-RACH-DualConfig-r19"},
	},
}

const (
	BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_SingleConfig_r19 = 0
	BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_DualConfig_r19   = 1
)

type BWP_UplinkCommon struct {
	GenericParameters BWP
	Rach_ConfigCommon *struct {
		Choice  int
		Release *struct{}
		Setup   *RACH_ConfigCommon
	}
	Pusch_ConfigCommon *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_ConfigCommon
	}
	Pucch_ConfigCommon *struct {
		Choice  int
		Release *struct{}
		Setup   *PUCCH_ConfigCommon
	}
	Rach_ConfigCommonIAB_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *RACH_ConfigCommon
	}
	UseInterlacePUCCH_PUSCH_r16 *int64
	MsgA_ConfigCommon_r16       *struct {
		Choice  int
		Release *struct{}
		Setup   *MsgA_ConfigCommon_r16
	}
	EnableRA_PrioritizationForSlicing_r17 *bool
	AdditionalRACH_ConfigList_r17         *struct {
		Choice  int
		Release *struct{}
		Setup   *AdditionalRACH_ConfigList_r17
	}
	Rsrp_ThresholdMsg3_r17                     *RSRP_Range
	NumberOfMsg3_RepetitionsList_r17           []NumberOfMsg3_Repetitions_r17
	Mcs_Msg3_Repetitions_r17                   []int64
	AdditionalRACH_PerPCI_ToAddModList_r18     []RACH_ConfigTwoTA_r18
	AdditionalRACH_PerPCI_ToReleaseList_r18    []AdditionalPCIIndex_r17
	Rsrp_ThresholdMsg1_RepetitionNum2_r18      *RSRP_Range
	Rsrp_ThresholdMsg1_RepetitionNum4_r18      *RSRP_Range
	Rsrp_ThresholdMsg1_RepetitionNum8_r18      *RSRP_Range
	PreambleTransMax_Msg1_Repetition_r18       *int64
	PreambleTransMaxRO_Type_r19                *int64
	Sbfd_RO_Type_r19                           *int64
	Sbfd_RSRP_ThresholdRO_Type_r19             *RSRP_Range
	Sbfd_RSRP_ThresholdRO_TypeUsage_r19        *int64
	Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19 *RSRP_Range
	Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19 *RSRP_Range
	Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19 *RSRP_Range
	Sbfd_RACH_Config_r19                       *struct {
		Choice                     int
		Sbfd_RACH_SingleConfig_r19 *struct{}
		Sbfd_RACH_DualConfig_r19   *SBFD_RACH_DualConfig_r19
	}
}

func (ie *BWP_UplinkCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bWPUplinkCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Rach_ConfigCommonIAB_r16 != nil || ie.UseInterlacePUCCH_PUSCH_r16 != nil || ie.MsgA_ConfigCommon_r16 != nil
	hasExtGroup1 := ie.EnableRA_PrioritizationForSlicing_r17 != nil || ie.AdditionalRACH_ConfigList_r17 != nil || ie.Rsrp_ThresholdMsg3_r17 != nil || ie.NumberOfMsg3_RepetitionsList_r17 != nil || ie.Mcs_Msg3_Repetitions_r17 != nil
	hasExtGroup2 := ie.AdditionalRACH_PerPCI_ToAddModList_r18 != nil || ie.AdditionalRACH_PerPCI_ToReleaseList_r18 != nil || ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18 != nil || ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18 != nil || ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18 != nil || ie.PreambleTransMax_Msg1_Repetition_r18 != nil
	hasExtGroup3 := ie.PreambleTransMaxRO_Type_r19 != nil || ie.Sbfd_RO_Type_r19 != nil || ie.Sbfd_RSRP_ThresholdRO_Type_r19 != nil || ie.Sbfd_RSRP_ThresholdRO_TypeUsage_r19 != nil || ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19 != nil || ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19 != nil || ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19 != nil || ie.Sbfd_RACH_Config_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Rach_ConfigCommon != nil, ie.Pusch_ConfigCommon != nil, ie.Pucch_ConfigCommon != nil}); err != nil {
		return err
	}

	// 3. genericParameters: ref
	{
		if err := ie.GenericParameters.Encode(e); err != nil {
			return err
		}
	}

	// 4. rach-ConfigCommon: choice
	{
		if ie.Rach_ConfigCommon != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkCommonRachConfigCommonConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Rach_ConfigCommon).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Rach_ConfigCommon).Choice {
			case BWP_UplinkCommon_Rach_ConfigCommon_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Rach_ConfigCommon_Setup:
				if err := (*ie.Rach_ConfigCommon).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Rach_ConfigCommon).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. pusch-ConfigCommon: choice
	{
		if ie.Pusch_ConfigCommon != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkCommonPuschConfigCommonConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_ConfigCommon).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pusch_ConfigCommon).Choice {
			case BWP_UplinkCommon_Pusch_ConfigCommon_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Pusch_ConfigCommon_Setup:
				if err := (*ie.Pusch_ConfigCommon).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pusch_ConfigCommon).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. pucch-ConfigCommon: choice
	{
		if ie.Pucch_ConfigCommon != nil {
			choiceEnc := e.NewChoiceEncoder(bWP_UplinkCommonPucchConfigCommonConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pucch_ConfigCommon).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pucch_ConfigCommon).Choice {
			case BWP_UplinkCommon_Pucch_ConfigCommon_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Pucch_ConfigCommon_Setup:
				if err := (*ie.Pucch_ConfigCommon).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pucch_ConfigCommon).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "rach-ConfigCommonIAB-r16", Optional: true},
					{Name: "useInterlacePUCCH-PUSCH-r16", Optional: true},
					{Name: "msgA-ConfigCommon-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Rach_ConfigCommonIAB_r16 != nil, ie.UseInterlacePUCCH_PUSCH_r16 != nil, ie.MsgA_ConfigCommon_r16 != nil}); err != nil {
				return err
			}

			if ie.Rach_ConfigCommonIAB_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkCommonExtRachConfigCommonIABR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Rach_ConfigCommonIAB_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Rach_ConfigCommonIAB_r16).Choice {
				case BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Setup:
					if err := (*ie.Rach_ConfigCommonIAB_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.UseInterlacePUCCH_PUSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UseInterlacePUCCH_PUSCH_r16, bWPUplinkCommonExtUseInterlacePUCCHPUSCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_ConfigCommon_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkCommonExtMsgAConfigCommonR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MsgA_ConfigCommon_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MsgA_ConfigCommon_r16).Choice {
				case BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Setup:
					if err := (*ie.MsgA_ConfigCommon_r16).Setup.Encode(ex); err != nil {
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
					{Name: "enableRA-PrioritizationForSlicing-r17", Optional: true},
					{Name: "additionalRACH-ConfigList-r17", Optional: true},
					{Name: "rsrp-ThresholdMsg3-r17", Optional: true},
					{Name: "numberOfMsg3-RepetitionsList-r17", Optional: true},
					{Name: "mcs-Msg3-Repetitions-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnableRA_PrioritizationForSlicing_r17 != nil, ie.AdditionalRACH_ConfigList_r17 != nil, ie.Rsrp_ThresholdMsg3_r17 != nil, ie.NumberOfMsg3_RepetitionsList_r17 != nil, ie.Mcs_Msg3_Repetitions_r17 != nil}); err != nil {
				return err
			}

			if ie.EnableRA_PrioritizationForSlicing_r17 != nil {
				if err := ex.EncodeBoolean(*ie.EnableRA_PrioritizationForSlicing_r17); err != nil {
					return err
				}
			}

			if ie.AdditionalRACH_ConfigList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkCommonExtAdditionalRACHConfigListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AdditionalRACH_ConfigList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AdditionalRACH_ConfigList_r17).Choice {
				case BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Setup:
					if err := (*ie.AdditionalRACH_ConfigList_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Rsrp_ThresholdMsg3_r17 != nil {
				if err := ie.Rsrp_ThresholdMsg3_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.NumberOfMsg3_RepetitionsList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkCommonExtNumberOfMsg3RepetitionsListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.NumberOfMsg3_RepetitionsList_r17))); err != nil {
					return err
				}
				for i := range ie.NumberOfMsg3_RepetitionsList_r17 {
					if err := ie.NumberOfMsg3_RepetitionsList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Mcs_Msg3_Repetitions_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkCommonExtMcsMsg3RepetitionsR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mcs_Msg3_Repetitions_r17))); err != nil {
					return err
				}
				for i := range ie.Mcs_Msg3_Repetitions_r17 {
					if err := ex.EncodeInteger(ie.Mcs_Msg3_Repetitions_r17[i], per.Constrained(0, 31)); err != nil {
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
					{Name: "additionalRACH-perPCI-ToAddModList-r18", Optional: true},
					{Name: "additionalRACH-perPCI-ToReleaseList-r18", Optional: true},
					{Name: "rsrp-ThresholdMsg1-RepetitionNum2-r18", Optional: true},
					{Name: "rsrp-ThresholdMsg1-RepetitionNum4-r18", Optional: true},
					{Name: "rsrp-ThresholdMsg1-RepetitionNum8-r18", Optional: true},
					{Name: "preambleTransMax-Msg1-Repetition-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdditionalRACH_PerPCI_ToAddModList_r18 != nil, ie.AdditionalRACH_PerPCI_ToReleaseList_r18 != nil, ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18 != nil, ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18 != nil, ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18 != nil, ie.PreambleTransMax_Msg1_Repetition_r18 != nil}); err != nil {
				return err
			}

			if ie.AdditionalRACH_PerPCI_ToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkCommonExtAdditionalRACHPerPCIToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AdditionalRACH_PerPCI_ToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.AdditionalRACH_PerPCI_ToAddModList_r18 {
					if err := ie.AdditionalRACH_PerPCI_ToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AdditionalRACH_PerPCI_ToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(bWPUplinkCommonExtAdditionalRACHPerPCIToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AdditionalRACH_PerPCI_ToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.AdditionalRACH_PerPCI_ToReleaseList_r18 {
					if err := ie.AdditionalRACH_PerPCI_ToReleaseList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18 != nil {
				if err := ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18 != nil {
				if err := ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18 != nil {
				if err := ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PreambleTransMax_Msg1_Repetition_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PreambleTransMax_Msg1_Repetition_r18, bWPUplinkCommonExtPreambleTransMaxMsg1RepetitionR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "preambleTransMaxRO-Type-r19", Optional: true},
					{Name: "sbfd-RO-Type-r19", Optional: true},
					{Name: "sbfd-RSRP-ThresholdRO-Type-r19", Optional: true},
					{Name: "sbfd-RSRP-ThresholdRO-TypeUsage-r19", Optional: true},
					{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum2-r19", Optional: true},
					{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum4-r19", Optional: true},
					{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum8-r19", Optional: true},
					{Name: "sbfd-RACH-Config-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PreambleTransMaxRO_Type_r19 != nil, ie.Sbfd_RO_Type_r19 != nil, ie.Sbfd_RSRP_ThresholdRO_Type_r19 != nil, ie.Sbfd_RSRP_ThresholdRO_TypeUsage_r19 != nil, ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19 != nil, ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19 != nil, ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19 != nil, ie.Sbfd_RACH_Config_r19 != nil}); err != nil {
				return err
			}

			if ie.PreambleTransMaxRO_Type_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PreambleTransMaxRO_Type_r19, bWPUplinkCommonExtPreambleTransMaxROTypeR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_RO_Type_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sbfd_RO_Type_r19, bWPUplinkCommonExtSbfdROTypeR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_RSRP_ThresholdRO_Type_r19 != nil {
				if err := ie.Sbfd_RSRP_ThresholdRO_Type_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sbfd_RSRP_ThresholdRO_TypeUsage_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sbfd_RSRP_ThresholdRO_TypeUsage_r19, bWPUplinkCommonExtSbfdRSRPThresholdROTypeUsageR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19 != nil {
				if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19 != nil {
				if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19 != nil {
				if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sbfd_RACH_Config_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(bWPUplinkCommonExtSbfdRACHConfigR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Sbfd_RACH_Config_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Sbfd_RACH_Config_r19).Choice {
				case BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_SingleConfig_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_DualConfig_r19:
					if err := (*ie.Sbfd_RACH_Config_r19).Sbfd_RACH_DualConfig_r19.Encode(ex); err != nil {
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

func (ie *BWP_UplinkCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bWPUplinkCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. genericParameters: ref
	{
		if err := ie.GenericParameters.Decode(d); err != nil {
			return err
		}
	}

	// 4. rach-ConfigCommon: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Rach_ConfigCommon = &struct {
				Choice  int
				Release *struct{}
				Setup   *RACH_ConfigCommon
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkCommonRachConfigCommonConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rach_ConfigCommon).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkCommon_Rach_ConfigCommon_Release:
				(*ie.Rach_ConfigCommon).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Rach_ConfigCommon_Setup:
				(*ie.Rach_ConfigCommon).Setup = new(RACH_ConfigCommon)
				if err := (*ie.Rach_ConfigCommon).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. pusch-ConfigCommon: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Pusch_ConfigCommon = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_ConfigCommon
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkCommonPuschConfigCommonConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_ConfigCommon).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkCommon_Pusch_ConfigCommon_Release:
				(*ie.Pusch_ConfigCommon).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Pusch_ConfigCommon_Setup:
				(*ie.Pusch_ConfigCommon).Setup = new(PUSCH_ConfigCommon)
				if err := (*ie.Pusch_ConfigCommon).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. pucch-ConfigCommon: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Pucch_ConfigCommon = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUCCH_ConfigCommon
			}{}
			choiceDec := d.NewChoiceDecoder(bWP_UplinkCommonPucchConfigCommonConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pucch_ConfigCommon).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BWP_UplinkCommon_Pucch_ConfigCommon_Release:
				(*ie.Pucch_ConfigCommon).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Pucch_ConfigCommon_Setup:
				(*ie.Pucch_ConfigCommon).Setup = new(PUCCH_ConfigCommon)
				if err := (*ie.Pucch_ConfigCommon).Setup.Decode(d); err != nil {
					return err
				}
			}
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
				{Name: "rach-ConfigCommonIAB-r16", Optional: true},
				{Name: "useInterlacePUCCH-PUSCH-r16", Optional: true},
				{Name: "msgA-ConfigCommon-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Rach_ConfigCommonIAB_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RACH_ConfigCommon
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkCommonExtRachConfigCommonIABR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Rach_ConfigCommonIAB_r16).Choice = int(index)
			switch index {
			case BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Release:
				(*ie.Rach_ConfigCommonIAB_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Ext_Rach_ConfigCommonIAB_r16_Setup:
				(*ie.Rach_ConfigCommonIAB_r16).Setup = new(RACH_ConfigCommon)
				if err := (*ie.Rach_ConfigCommonIAB_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bWPUplinkCommonExtUseInterlacePUCCHPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.UseInterlacePUCCH_PUSCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MsgA_ConfigCommon_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MsgA_ConfigCommon_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkCommonExtMsgAConfigCommonR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MsgA_ConfigCommon_r16).Choice = int(index)
			switch index {
			case BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Release:
				(*ie.MsgA_ConfigCommon_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Ext_MsgA_ConfigCommon_r16_Setup:
				(*ie.MsgA_ConfigCommon_r16).Setup = new(MsgA_ConfigCommon_r16)
				if err := (*ie.MsgA_ConfigCommon_r16).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "enableRA-PrioritizationForSlicing-r17", Optional: true},
				{Name: "additionalRACH-ConfigList-r17", Optional: true},
				{Name: "rsrp-ThresholdMsg3-r17", Optional: true},
				{Name: "numberOfMsg3-RepetitionsList-r17", Optional: true},
				{Name: "mcs-Msg3-Repetitions-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.EnableRA_PrioritizationForSlicing_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AdditionalRACH_ConfigList_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *AdditionalRACH_ConfigList_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkCommonExtAdditionalRACHConfigListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AdditionalRACH_ConfigList_r17).Choice = int(index)
			switch index {
			case BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Release:
				(*ie.AdditionalRACH_ConfigList_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Ext_AdditionalRACH_ConfigList_r17_Setup:
				(*ie.AdditionalRACH_ConfigList_r17).Setup = new(AdditionalRACH_ConfigList_r17)
				if err := (*ie.AdditionalRACH_ConfigList_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Rsrp_ThresholdMsg3_r17 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdMsg3_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkCommonExtNumberOfMsg3RepetitionsListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NumberOfMsg3_RepetitionsList_r17 = make([]NumberOfMsg3_Repetitions_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.NumberOfMsg3_RepetitionsList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkCommonExtMcsMsg3RepetitionsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mcs_Msg3_Repetitions_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(0, 31))
				if err != nil {
					return err
				}
				ie.Mcs_Msg3_Repetitions_r17[i] = v
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
				{Name: "additionalRACH-perPCI-ToAddModList-r18", Optional: true},
				{Name: "additionalRACH-perPCI-ToReleaseList-r18", Optional: true},
				{Name: "rsrp-ThresholdMsg1-RepetitionNum2-r18", Optional: true},
				{Name: "rsrp-ThresholdMsg1-RepetitionNum4-r18", Optional: true},
				{Name: "rsrp-ThresholdMsg1-RepetitionNum8-r18", Optional: true},
				{Name: "preambleTransMax-Msg1-Repetition-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkCommonExtAdditionalRACHPerPCIToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalRACH_PerPCI_ToAddModList_r18 = make([]RACH_ConfigTwoTA_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalRACH_PerPCI_ToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(bWPUplinkCommonExtAdditionalRACHPerPCIToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AdditionalRACH_PerPCI_ToReleaseList_r18 = make([]AdditionalPCIIndex_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AdditionalRACH_PerPCI_ToReleaseList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdMsg1_RepetitionNum2_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdMsg1_RepetitionNum4_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18 = new(RSRP_Range)
			if err := ie.Rsrp_ThresholdMsg1_RepetitionNum8_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bWPUplinkCommonExtPreambleTransMaxMsg1RepetitionR18Constraints)
			if err != nil {
				return err
			}
			ie.PreambleTransMax_Msg1_Repetition_r18 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "preambleTransMaxRO-Type-r19", Optional: true},
				{Name: "sbfd-RO-Type-r19", Optional: true},
				{Name: "sbfd-RSRP-ThresholdRO-Type-r19", Optional: true},
				{Name: "sbfd-RSRP-ThresholdRO-TypeUsage-r19", Optional: true},
				{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum2-r19", Optional: true},
				{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum4-r19", Optional: true},
				{Name: "sbfd-RSRP-ThresholdMsg1-RepetitionNum8-r19", Optional: true},
				{Name: "sbfd-RACH-Config-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bWPUplinkCommonExtPreambleTransMaxROTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.PreambleTransMaxRO_Type_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bWPUplinkCommonExtSbfdROTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_RO_Type_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Sbfd_RSRP_ThresholdRO_Type_r19 = new(RSRP_Range)
			if err := ie.Sbfd_RSRP_ThresholdRO_Type_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bWPUplinkCommonExtSbfdRSRPThresholdROTypeUsageR19Constraints)
			if err != nil {
				return err
			}
			ie.Sbfd_RSRP_ThresholdRO_TypeUsage_r19 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19 = new(RSRP_Range)
			if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum2_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19 = new(RSRP_Range)
			if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum4_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19 = new(RSRP_Range)
			if err := ie.Sbfd_RSRP_ThresholdMsg1_RepetitionNum8_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Sbfd_RACH_Config_r19 = &struct {
				Choice                     int
				Sbfd_RACH_SingleConfig_r19 *struct{}
				Sbfd_RACH_DualConfig_r19   *SBFD_RACH_DualConfig_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(bWPUplinkCommonExtSbfdRACHConfigR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sbfd_RACH_Config_r19).Choice = int(index)
			switch index {
			case BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_SingleConfig_r19:
				(*ie.Sbfd_RACH_Config_r19).Sbfd_RACH_SingleConfig_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case BWP_UplinkCommon_Ext_Sbfd_RACH_Config_r19_Sbfd_RACH_DualConfig_r19:
				(*ie.Sbfd_RACH_Config_r19).Sbfd_RACH_DualConfig_r19 = new(SBFD_RACH_DualConfig_r19)
				if err := (*ie.Sbfd_RACH_Config_r19).Sbfd_RACH_DualConfig_r19.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
