// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ServingCellConfigCommon (line 14897).

var servingCellConfigCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "physCellId", Optional: true},
		{Name: "downlinkConfigCommon", Optional: true},
		{Name: "uplinkConfigCommon", Optional: true},
		{Name: "supplementaryUplinkConfig", Optional: true},
		{Name: "n-TimingAdvanceOffset", Optional: true},
		{Name: "ssb-PositionsInBurst", Optional: true},
		{Name: "ssb-periodicityServingCell", Optional: true},
		{Name: "dmrs-TypeA-Position"},
		{Name: "lte-CRS-ToMatchAround", Optional: true},
		{Name: "rateMatchPatternToAddModList", Optional: true},
		{Name: "rateMatchPatternToReleaseList", Optional: true},
		{Name: "ssbSubcarrierSpacing", Optional: true},
		{Name: "tdd-UL-DL-ConfigurationCommon", Optional: true},
		{Name: "ss-PBCH-BlockPower"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	ServingCellConfigCommon_N_TimingAdvanceOffset_N0     = 0
	ServingCellConfigCommon_N_TimingAdvanceOffset_N25600 = 1
	ServingCellConfigCommon_N_TimingAdvanceOffset_N39936 = 2
)

var servingCellConfigCommonNTimingAdvanceOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_N_TimingAdvanceOffset_N0, ServingCellConfigCommon_N_TimingAdvanceOffset_N25600, ServingCellConfigCommon_N_TimingAdvanceOffset_N39936},
}

var servingCellConfigCommonSsbPositionsInBurstConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "shortBitmap"},
		{Name: "mediumBitmap"},
		{Name: "longBitmap"},
	},
}

const (
	ServingCellConfigCommon_Ssb_PositionsInBurst_ShortBitmap  = 0
	ServingCellConfigCommon_Ssb_PositionsInBurst_MediumBitmap = 1
	ServingCellConfigCommon_Ssb_PositionsInBurst_LongBitmap   = 2
)

const (
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms5    = 0
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms10   = 1
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms20   = 2
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms40   = 3
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms80   = 4
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms160  = 5
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Spare2 = 6
	ServingCellConfigCommon_Ssb_PeriodicityServingCell_Spare1 = 7
)

var servingCellConfigCommonSsbPeriodicityServingCellConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms5, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms10, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms20, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms40, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms80, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Ms160, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Spare2, ServingCellConfigCommon_Ssb_PeriodicityServingCell_Spare1},
}

const (
	ServingCellConfigCommon_Dmrs_TypeA_Position_Pos2 = 0
	ServingCellConfigCommon_Dmrs_TypeA_Position_Pos3 = 1
)

var servingCellConfigCommonDmrsTypeAPositionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Dmrs_TypeA_Position_Pos2, ServingCellConfigCommon_Dmrs_TypeA_Position_Pos3},
}

var servingCellConfigCommonLteCRSToMatchAroundConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfigCommon_Lte_CRS_ToMatchAround_Release = 0
	ServingCellConfigCommon_Lte_CRS_ToMatchAround_Setup   = 1
)

var servingCellConfigCommonRateMatchPatternToAddModListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

var servingCellConfigCommonRateMatchPatternToReleaseListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

var servingCellConfigCommonSsPBCHBlockPowerConstraints = per.Constrained(-60, 50)

var servingCellConfigCommonExtChannelAccessModeR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "dynamic"},
		{Name: "semiStatic"},
	},
}

const (
	ServingCellConfigCommon_Ext_ChannelAccessMode_r16_Dynamic    = 0
	ServingCellConfigCommon_Ext_ChannelAccessMode_r16_SemiStatic = 1
)

const (
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms0dot5 = 0
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms1     = 1
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms2     = 2
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms3     = 3
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms4     = 4
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms5     = 5
)

var servingCellConfigCommonExtDiscoveryBurstWindowLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms0dot5, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms1, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms2, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms3, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms4, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r16_Ms5},
}

const (
	ServingCellConfigCommon_Ext_ChannelAccessMode2_r17_Enabled = 0
)

var servingCellConfigCommonExtChannelAccessMode2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Ext_ChannelAccessMode2_r17_Enabled},
}

const (
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot125 = 0
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot25  = 1
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot5   = 2
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot75  = 3
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms1       = 4
	ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms1dot25  = 5
)

var servingCellConfigCommonExtDiscoveryBurstWindowLengthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot125, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot25, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot5, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms0dot75, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms1, ServingCellConfigCommon_Ext_DiscoveryBurstWindowLength_r17_Ms1dot25},
}

const (
	ServingCellConfigCommon_Ext_Ra_ChannelAccess_r17_Enabled = 0
)

var servingCellConfigCommonExtRaChannelAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfigCommon_Ext_Ra_ChannelAccess_r17_Enabled},
}

var servingCellConfigCommonExtFeaturePrioritiesR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "redCapPriority-r17", Optional: true},
		{Name: "slicingPriority-r17", Optional: true},
		{Name: "msg3-Repetitions-Priority-r17", Optional: true},
		{Name: "sdt-Priority-r17", Optional: true},
	},
}

var servingCellConfigCommonExtFeaturePrioritiesV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "msg1-Repetitions-Priority-r18", Optional: true},
		{Name: "eRedCapPriority-r18", Optional: true},
	},
}

type ServingCellConfigCommon struct {
	PhysCellId                *PhysCellId
	DownlinkConfigCommon      *DownlinkConfigCommon
	UplinkConfigCommon        *UplinkConfigCommon
	SupplementaryUplinkConfig *UplinkConfigCommon
	N_TimingAdvanceOffset     *int64
	Ssb_PositionsInBurst      *struct {
		Choice       int
		ShortBitmap  *per.BitString
		MediumBitmap *per.BitString
		LongBitmap   *per.BitString
	}
	Ssb_PeriodicityServingCell *int64
	Dmrs_TypeA_Position        int64
	Lte_CRS_ToMatchAround      *struct {
		Choice  int
		Release *struct{}
		Setup   *RateMatchPatternLTE_CRS
	}
	RateMatchPatternToAddModList  []RateMatchPattern
	RateMatchPatternToReleaseList []RateMatchPatternId
	SsbSubcarrierSpacing          *SubcarrierSpacing
	Tdd_UL_DL_ConfigurationCommon *TDD_UL_DL_ConfigCommon
	Ss_PBCH_BlockPower            int64
	ChannelAccessMode_r16         *struct {
		Choice     int
		Dynamic    *struct{}
		SemiStatic *SemiStaticChannelAccessConfig_r16
	}
	DiscoveryBurstWindowLength_r16 *int64
	Ssb_PositionQCL_r16            *SSB_PositionQCL_Relation_r16
	HighSpeedConfig_r16            *HighSpeedConfig_r16
	HighSpeedConfig_v1700          *HighSpeedConfig_v1700
	ChannelAccessMode2_r17         *int64
	DiscoveryBurstWindowLength_r17 *int64
	Ssb_PositionQCL_r17            *SSB_PositionQCL_Relation_r17
	HighSpeedConfigFR2_r17         *HighSpeedConfigFR2_r17
	UplinkConfigCommon_v1700       *UplinkConfigCommon_v1700
	Ntn_Config_r17                 *NTN_Config_r17
	FeaturePriorities_r17          *struct {
		RedCapPriority_r17            *FeaturePriority_r17
		SlicingPriority_r17           *FeaturePriority_r17
		Msg3_Repetitions_Priority_r17 *FeaturePriority_r17
		Sdt_Priority_r17              *FeaturePriority_r17
	}
	Ra_ChannelAccess_r17    *int64
	FeaturePriorities_v1800 *struct {
		Msg1_Repetitions_Priority_r18 *FeaturePriority_r17
		ERedCapPriority_r18           *FeaturePriority_r17
	}
	Atg_Config_r18 *ATG_Config_r18
}

func (ie *ServingCellConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servingCellConfigCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ChannelAccessMode_r16 != nil || ie.DiscoveryBurstWindowLength_r16 != nil || ie.Ssb_PositionQCL_r16 != nil || ie.HighSpeedConfig_r16 != nil
	hasExtGroup1 := ie.HighSpeedConfig_v1700 != nil || ie.ChannelAccessMode2_r17 != nil || ie.DiscoveryBurstWindowLength_r17 != nil || ie.Ssb_PositionQCL_r17 != nil || ie.HighSpeedConfigFR2_r17 != nil || ie.UplinkConfigCommon_v1700 != nil || ie.Ntn_Config_r17 != nil
	hasExtGroup2 := ie.FeaturePriorities_r17 != nil
	hasExtGroup3 := ie.Ra_ChannelAccess_r17 != nil
	hasExtGroup4 := ie.FeaturePriorities_v1800 != nil || ie.Atg_Config_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PhysCellId != nil, ie.DownlinkConfigCommon != nil, ie.UplinkConfigCommon != nil, ie.SupplementaryUplinkConfig != nil, ie.N_TimingAdvanceOffset != nil, ie.Ssb_PositionsInBurst != nil, ie.Ssb_PeriodicityServingCell != nil, ie.Lte_CRS_ToMatchAround != nil, ie.RateMatchPatternToAddModList != nil, ie.RateMatchPatternToReleaseList != nil, ie.SsbSubcarrierSpacing != nil, ie.Tdd_UL_DL_ConfigurationCommon != nil}); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if ie.PhysCellId != nil {
			if err := ie.PhysCellId.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. downlinkConfigCommon: ref
	{
		if ie.DownlinkConfigCommon != nil {
			if err := ie.DownlinkConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. uplinkConfigCommon: ref
	{
		if ie.UplinkConfigCommon != nil {
			if err := ie.UplinkConfigCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. supplementaryUplinkConfig: ref
	{
		if ie.SupplementaryUplinkConfig != nil {
			if err := ie.SupplementaryUplinkConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. n-TimingAdvanceOffset: enumerated
	{
		if ie.N_TimingAdvanceOffset != nil {
			if err := e.EncodeEnumerated(*ie.N_TimingAdvanceOffset, servingCellConfigCommonNTimingAdvanceOffsetConstraints); err != nil {
				return err
			}
		}
	}

	// 8. ssb-PositionsInBurst: choice
	{
		if ie.Ssb_PositionsInBurst != nil {
			choiceEnc := e.NewChoiceEncoder(servingCellConfigCommonSsbPositionsInBurstConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Ssb_PositionsInBurst).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Ssb_PositionsInBurst).Choice {
			case ServingCellConfigCommon_Ssb_PositionsInBurst_ShortBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst).ShortBitmap), per.FixedSize(4)); err != nil {
					return err
				}
			case ServingCellConfigCommon_Ssb_PositionsInBurst_MediumBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst).MediumBitmap), per.FixedSize(8)); err != nil {
					return err
				}
			case ServingCellConfigCommon_Ssb_PositionsInBurst_LongBitmap:
				if err := e.EncodeBitString((*(*ie.Ssb_PositionsInBurst).LongBitmap), per.FixedSize(64)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Ssb_PositionsInBurst).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 9. ssb-periodicityServingCell: enumerated
	{
		if ie.Ssb_PeriodicityServingCell != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_PeriodicityServingCell, servingCellConfigCommonSsbPeriodicityServingCellConstraints); err != nil {
				return err
			}
		}
	}

	// 10. dmrs-TypeA-Position: enumerated
	{
		if err := e.EncodeEnumerated(ie.Dmrs_TypeA_Position, servingCellConfigCommonDmrsTypeAPositionConstraints); err != nil {
			return err
		}
	}

	// 11. lte-CRS-ToMatchAround: choice
	{
		if ie.Lte_CRS_ToMatchAround != nil {
			choiceEnc := e.NewChoiceEncoder(servingCellConfigCommonLteCRSToMatchAroundConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_ToMatchAround).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Lte_CRS_ToMatchAround).Choice {
			case ServingCellConfigCommon_Lte_CRS_ToMatchAround_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ServingCellConfigCommon_Lte_CRS_ToMatchAround_Setup:
				if err := (*ie.Lte_CRS_ToMatchAround).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Lte_CRS_ToMatchAround).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. rateMatchPatternToAddModList: sequence-of
	{
		if ie.RateMatchPatternToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(servingCellConfigCommonRateMatchPatternToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToAddModList))); err != nil {
				return err
			}
			for i := range ie.RateMatchPatternToAddModList {
				if err := ie.RateMatchPatternToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. rateMatchPatternToReleaseList: sequence-of
	{
		if ie.RateMatchPatternToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(servingCellConfigCommonRateMatchPatternToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToReleaseList))); err != nil {
				return err
			}
			for i := range ie.RateMatchPatternToReleaseList {
				if err := ie.RateMatchPatternToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 14. ssbSubcarrierSpacing: ref
	{
		if ie.SsbSubcarrierSpacing != nil {
			if err := ie.SsbSubcarrierSpacing.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. tdd-UL-DL-ConfigurationCommon: ref
	{
		if ie.Tdd_UL_DL_ConfigurationCommon != nil {
			if err := ie.Tdd_UL_DL_ConfigurationCommon.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. ss-PBCH-BlockPower: integer
	{
		if err := e.EncodeInteger(ie.Ss_PBCH_BlockPower, servingCellConfigCommonSsPBCHBlockPowerConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "channelAccessMode-r16", Optional: true},
					{Name: "discoveryBurstWindowLength-r16", Optional: true},
					{Name: "ssb-PositionQCL-r16", Optional: true},
					{Name: "highSpeedConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelAccessMode_r16 != nil, ie.DiscoveryBurstWindowLength_r16 != nil, ie.Ssb_PositionQCL_r16 != nil, ie.HighSpeedConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.ChannelAccessMode_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigCommonExtChannelAccessModeR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelAccessMode_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelAccessMode_r16).Choice {
				case ServingCellConfigCommon_Ext_ChannelAccessMode_r16_Dynamic:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfigCommon_Ext_ChannelAccessMode_r16_SemiStatic:
					if err := (*ie.ChannelAccessMode_r16).SemiStatic.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.DiscoveryBurstWindowLength_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DiscoveryBurstWindowLength_r16, servingCellConfigCommonExtDiscoveryBurstWindowLengthR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_r16 != nil {
				if err := ie.Ssb_PositionQCL_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.HighSpeedConfig_r16 != nil {
				if err := ie.HighSpeedConfig_r16.Encode(ex); err != nil {
					return err
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
					{Name: "highSpeedConfig-v1700", Optional: true},
					{Name: "channelAccessMode2-r17", Optional: true},
					{Name: "discoveryBurstWindowLength-r17", Optional: true},
					{Name: "ssb-PositionQCL-r17", Optional: true},
					{Name: "highSpeedConfigFR2-r17", Optional: true},
					{Name: "uplinkConfigCommon-v1700", Optional: true},
					{Name: "ntn-Config-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.HighSpeedConfig_v1700 != nil, ie.ChannelAccessMode2_r17 != nil, ie.DiscoveryBurstWindowLength_r17 != nil, ie.Ssb_PositionQCL_r17 != nil, ie.HighSpeedConfigFR2_r17 != nil, ie.UplinkConfigCommon_v1700 != nil, ie.Ntn_Config_r17 != nil}); err != nil {
				return err
			}

			if ie.HighSpeedConfig_v1700 != nil {
				if err := ie.HighSpeedConfig_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ChannelAccessMode2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ChannelAccessMode2_r17, servingCellConfigCommonExtChannelAccessMode2R17Constraints); err != nil {
					return err
				}
			}

			if ie.DiscoveryBurstWindowLength_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DiscoveryBurstWindowLength_r17, servingCellConfigCommonExtDiscoveryBurstWindowLengthR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ssb_PositionQCL_r17 != nil {
				if err := ie.Ssb_PositionQCL_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.HighSpeedConfigFR2_r17 != nil {
				if err := ie.HighSpeedConfigFR2_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.UplinkConfigCommon_v1700 != nil {
				if err := ie.UplinkConfigCommon_v1700.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Ntn_Config_r17 != nil {
				if err := ie.Ntn_Config_r17.Encode(ex); err != nil {
					return err
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
					{Name: "featurePriorities-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeaturePriorities_r17 != nil}); err != nil {
				return err
			}

			if ie.FeaturePriorities_r17 != nil {
				c := ie.FeaturePriorities_r17
				servingCellConfigCommonExtFeaturePrioritiesR17Seq := ex.NewSequenceEncoder(servingCellConfigCommonExtFeaturePrioritiesR17Constraints)
				if err := servingCellConfigCommonExtFeaturePrioritiesR17Seq.EncodePreamble([]bool{c.RedCapPriority_r17 != nil, c.SlicingPriority_r17 != nil, c.Msg3_Repetitions_Priority_r17 != nil, c.Sdt_Priority_r17 != nil}); err != nil {
					return err
				}
				if c.RedCapPriority_r17 != nil {
					if err := c.RedCapPriority_r17.Encode(ex); err != nil {
						return err
					}
				}
				if c.SlicingPriority_r17 != nil {
					if err := c.SlicingPriority_r17.Encode(ex); err != nil {
						return err
					}
				}
				if c.Msg3_Repetitions_Priority_r17 != nil {
					if err := c.Msg3_Repetitions_Priority_r17.Encode(ex); err != nil {
						return err
					}
				}
				if c.Sdt_Priority_r17 != nil {
					if err := c.Sdt_Priority_r17.Encode(ex); err != nil {
						return err
					}
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
					{Name: "ra-ChannelAccess-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ra_ChannelAccess_r17 != nil}); err != nil {
				return err
			}

			if ie.Ra_ChannelAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ra_ChannelAccess_r17, servingCellConfigCommonExtRaChannelAccessR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "featurePriorities-v1800", Optional: true},
					{Name: "atg-Config-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.FeaturePriorities_v1800 != nil, ie.Atg_Config_r18 != nil}); err != nil {
				return err
			}

			if ie.FeaturePriorities_v1800 != nil {
				c := ie.FeaturePriorities_v1800
				servingCellConfigCommonExtFeaturePrioritiesV1800Seq := ex.NewSequenceEncoder(servingCellConfigCommonExtFeaturePrioritiesV1800Constraints)
				if err := servingCellConfigCommonExtFeaturePrioritiesV1800Seq.EncodePreamble([]bool{c.Msg1_Repetitions_Priority_r18 != nil, c.ERedCapPriority_r18 != nil}); err != nil {
					return err
				}
				if c.Msg1_Repetitions_Priority_r18 != nil {
					if err := c.Msg1_Repetitions_Priority_r18.Encode(ex); err != nil {
						return err
					}
				}
				if c.ERedCapPriority_r18 != nil {
					if err := c.ERedCapPriority_r18.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Atg_Config_r18 != nil {
				if err := ie.Atg_Config_r18.Encode(ex); err != nil {
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

func (ie *ServingCellConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servingCellConfigCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. physCellId: ref
	{
		if seq.IsComponentPresent(0) {
			ie.PhysCellId = new(PhysCellId)
			if err := ie.PhysCellId.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. downlinkConfigCommon: ref
	{
		if seq.IsComponentPresent(1) {
			ie.DownlinkConfigCommon = new(DownlinkConfigCommon)
			if err := ie.DownlinkConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. uplinkConfigCommon: ref
	{
		if seq.IsComponentPresent(2) {
			ie.UplinkConfigCommon = new(UplinkConfigCommon)
			if err := ie.UplinkConfigCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. supplementaryUplinkConfig: ref
	{
		if seq.IsComponentPresent(3) {
			ie.SupplementaryUplinkConfig = new(UplinkConfigCommon)
			if err := ie.SupplementaryUplinkConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. n-TimingAdvanceOffset: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(servingCellConfigCommonNTimingAdvanceOffsetConstraints)
			if err != nil {
				return err
			}
			ie.N_TimingAdvanceOffset = &idx
		}
	}

	// 8. ssb-PositionsInBurst: choice
	{
		if seq.IsComponentPresent(5) {
			ie.Ssb_PositionsInBurst = &struct {
				Choice       int
				ShortBitmap  *per.BitString
				MediumBitmap *per.BitString
				LongBitmap   *per.BitString
			}{}
			choiceDec := d.NewChoiceDecoder(servingCellConfigCommonSsbPositionsInBurstConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ssb_PositionsInBurst).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ServingCellConfigCommon_Ssb_PositionsInBurst_ShortBitmap:
				(*ie.Ssb_PositionsInBurst).ShortBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(4))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst).ShortBitmap) = v
			case ServingCellConfigCommon_Ssb_PositionsInBurst_MediumBitmap:
				(*ie.Ssb_PositionsInBurst).MediumBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(8))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst).MediumBitmap) = v
			case ServingCellConfigCommon_Ssb_PositionsInBurst_LongBitmap:
				(*ie.Ssb_PositionsInBurst).LongBitmap = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(64))
				if err != nil {
					return err
				}
				(*(*ie.Ssb_PositionsInBurst).LongBitmap) = v
			}
		}
	}

	// 9. ssb-periodicityServingCell: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(servingCellConfigCommonSsbPeriodicityServingCellConstraints)
			if err != nil {
				return err
			}
			ie.Ssb_PeriodicityServingCell = &idx
		}
	}

	// 10. dmrs-TypeA-Position: enumerated
	{
		v7, err := d.DecodeEnumerated(servingCellConfigCommonDmrsTypeAPositionConstraints)
		if err != nil {
			return err
		}
		ie.Dmrs_TypeA_Position = v7
	}

	// 11. lte-CRS-ToMatchAround: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Lte_CRS_ToMatchAround = &struct {
				Choice  int
				Release *struct{}
				Setup   *RateMatchPatternLTE_CRS
			}{}
			choiceDec := d.NewChoiceDecoder(servingCellConfigCommonLteCRSToMatchAroundConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_ToMatchAround).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ServingCellConfigCommon_Lte_CRS_ToMatchAround_Release:
				(*ie.Lte_CRS_ToMatchAround).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfigCommon_Lte_CRS_ToMatchAround_Setup:
				(*ie.Lte_CRS_ToMatchAround).Setup = new(RateMatchPatternLTE_CRS)
				if err := (*ie.Lte_CRS_ToMatchAround).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. rateMatchPatternToAddModList: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(servingCellConfigCommonRateMatchPatternToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToAddModList = make([]RateMatchPattern, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. rateMatchPatternToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(servingCellConfigCommonRateMatchPatternToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToReleaseList = make([]RateMatchPatternId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. ssbSubcarrierSpacing: ref
	{
		if seq.IsComponentPresent(11) {
			ie.SsbSubcarrierSpacing = new(SubcarrierSpacing)
			if err := ie.SsbSubcarrierSpacing.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. tdd-UL-DL-ConfigurationCommon: ref
	{
		if seq.IsComponentPresent(12) {
			ie.Tdd_UL_DL_ConfigurationCommon = new(TDD_UL_DL_ConfigCommon)
			if err := ie.Tdd_UL_DL_ConfigurationCommon.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. ss-PBCH-BlockPower: integer
	{
		v13, err := d.DecodeInteger(servingCellConfigCommonSsPBCHBlockPowerConstraints)
		if err != nil {
			return err
		}
		ie.Ss_PBCH_BlockPower = v13
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
				{Name: "channelAccessMode-r16", Optional: true},
				{Name: "discoveryBurstWindowLength-r16", Optional: true},
				{Name: "ssb-PositionQCL-r16", Optional: true},
				{Name: "highSpeedConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ChannelAccessMode_r16 = &struct {
				Choice     int
				Dynamic    *struct{}
				SemiStatic *SemiStaticChannelAccessConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigCommonExtChannelAccessModeR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelAccessMode_r16).Choice = int(index)
			switch index {
			case ServingCellConfigCommon_Ext_ChannelAccessMode_r16_Dynamic:
				(*ie.ChannelAccessMode_r16).Dynamic = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfigCommon_Ext_ChannelAccessMode_r16_SemiStatic:
				(*ie.ChannelAccessMode_r16).SemiStatic = new(SemiStaticChannelAccessConfig_r16)
				if err := (*ie.ChannelAccessMode_r16).SemiStatic.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonExtDiscoveryBurstWindowLengthR16Constraints)
			if err != nil {
				return err
			}
			ie.DiscoveryBurstWindowLength_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Ssb_PositionQCL_r16 = new(SSB_PositionQCL_Relation_r16)
			if err := ie.Ssb_PositionQCL_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.HighSpeedConfig_r16 = new(HighSpeedConfig_r16)
			if err := ie.HighSpeedConfig_r16.Decode(dx); err != nil {
				return err
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
				{Name: "highSpeedConfig-v1700", Optional: true},
				{Name: "channelAccessMode2-r17", Optional: true},
				{Name: "discoveryBurstWindowLength-r17", Optional: true},
				{Name: "ssb-PositionQCL-r17", Optional: true},
				{Name: "highSpeedConfigFR2-r17", Optional: true},
				{Name: "uplinkConfigCommon-v1700", Optional: true},
				{Name: "ntn-Config-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.HighSpeedConfig_v1700 = new(HighSpeedConfig_v1700)
			if err := ie.HighSpeedConfig_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonExtChannelAccessMode2R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessMode2_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonExtDiscoveryBurstWindowLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.DiscoveryBurstWindowLength_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Ssb_PositionQCL_r17 = new(SSB_PositionQCL_Relation_r17)
			if err := ie.Ssb_PositionQCL_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.HighSpeedConfigFR2_r17 = new(HighSpeedConfigFR2_r17)
			if err := ie.HighSpeedConfigFR2_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.UplinkConfigCommon_v1700 = new(UplinkConfigCommon_v1700)
			if err := ie.UplinkConfigCommon_v1700.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Ntn_Config_r17 = new(NTN_Config_r17)
			if err := ie.Ntn_Config_r17.Decode(dx); err != nil {
				return err
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
				{Name: "featurePriorities-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FeaturePriorities_r17 = &struct {
				RedCapPriority_r17            *FeaturePriority_r17
				SlicingPriority_r17           *FeaturePriority_r17
				Msg3_Repetitions_Priority_r17 *FeaturePriority_r17
				Sdt_Priority_r17              *FeaturePriority_r17
			}{}
			c := ie.FeaturePriorities_r17
			servingCellConfigCommonExtFeaturePrioritiesR17Seq := dx.NewSequenceDecoder(servingCellConfigCommonExtFeaturePrioritiesR17Constraints)
			if err := servingCellConfigCommonExtFeaturePrioritiesR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if servingCellConfigCommonExtFeaturePrioritiesR17Seq.IsComponentPresent(0) {
				c.RedCapPriority_r17 = new(FeaturePriority_r17)
				if err := (*c.RedCapPriority_r17).Decode(dx); err != nil {
					return err
				}
			}
			if servingCellConfigCommonExtFeaturePrioritiesR17Seq.IsComponentPresent(1) {
				c.SlicingPriority_r17 = new(FeaturePriority_r17)
				if err := (*c.SlicingPriority_r17).Decode(dx); err != nil {
					return err
				}
			}
			if servingCellConfigCommonExtFeaturePrioritiesR17Seq.IsComponentPresent(2) {
				c.Msg3_Repetitions_Priority_r17 = new(FeaturePriority_r17)
				if err := (*c.Msg3_Repetitions_Priority_r17).Decode(dx); err != nil {
					return err
				}
			}
			if servingCellConfigCommonExtFeaturePrioritiesR17Seq.IsComponentPresent(3) {
				c.Sdt_Priority_r17 = new(FeaturePriority_r17)
				if err := (*c.Sdt_Priority_r17).Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ra-ChannelAccess-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigCommonExtRaChannelAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Ra_ChannelAccess_r17 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "featurePriorities-v1800", Optional: true},
				{Name: "atg-Config-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.FeaturePriorities_v1800 = &struct {
				Msg1_Repetitions_Priority_r18 *FeaturePriority_r17
				ERedCapPriority_r18           *FeaturePriority_r17
			}{}
			c := ie.FeaturePriorities_v1800
			servingCellConfigCommonExtFeaturePrioritiesV1800Seq := dx.NewSequenceDecoder(servingCellConfigCommonExtFeaturePrioritiesV1800Constraints)
			if err := servingCellConfigCommonExtFeaturePrioritiesV1800Seq.DecodePreamble(); err != nil {
				return err
			}
			if servingCellConfigCommonExtFeaturePrioritiesV1800Seq.IsComponentPresent(0) {
				c.Msg1_Repetitions_Priority_r18 = new(FeaturePriority_r17)
				if err := (*c.Msg1_Repetitions_Priority_r18).Decode(dx); err != nil {
					return err
				}
			}
			if servingCellConfigCommonExtFeaturePrioritiesV1800Seq.IsComponentPresent(1) {
				c.ERedCapPriority_r18 = new(FeaturePriority_r17)
				if err := (*c.ERedCapPriority_r18).Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Atg_Config_r18 = new(ATG_Config_r18)
			if err := ie.Atg_Config_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
