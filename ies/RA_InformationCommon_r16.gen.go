// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RA-InformationCommon-r16 (line 3087).

var rAInformationCommonR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "absoluteFrequencyPointA-r16"},
		{Name: "locationAndBandwidth-r16"},
		{Name: "subcarrierSpacing-r16"},
		{Name: "msg1-FrequencyStart-r16", Optional: true},
		{Name: "msg1-FrequencyStartCFRA-r16", Optional: true},
		{Name: "msg1-SubcarrierSpacing-r16", Optional: true},
		{Name: "msg1-SubcarrierSpacingCFRA-r16", Optional: true},
		{Name: "msg1-FDM-r16", Optional: true},
		{Name: "msg1-FDMCFRA-r16", Optional: true},
		{Name: "perRAInfoList-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

var rAInformationCommonR16LocationAndBandwidthR16Constraints = per.Constrained(0, 37949)

var rAInformationCommonR16Msg1FrequencyStartR16Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var rAInformationCommonR16Msg1FrequencyStartCFRAR16Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

const (
	RA_InformationCommon_r16_Msg1_FDM_r16_One   = 0
	RA_InformationCommon_r16_Msg1_FDM_r16_Two   = 1
	RA_InformationCommon_r16_Msg1_FDM_r16_Four  = 2
	RA_InformationCommon_r16_Msg1_FDM_r16_Eight = 3
)

var rAInformationCommonR16Msg1FDMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Msg1_FDM_r16_One, RA_InformationCommon_r16_Msg1_FDM_r16_Two, RA_InformationCommon_r16_Msg1_FDM_r16_Four, RA_InformationCommon_r16_Msg1_FDM_r16_Eight},
}

const (
	RA_InformationCommon_r16_Msg1_FDMCFRA_r16_One   = 0
	RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Two   = 1
	RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Four  = 2
	RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Eight = 3
)

var rAInformationCommonR16Msg1FDMCFRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Msg1_FDMCFRA_r16_One, RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Two, RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Four, RA_InformationCommon_r16_Msg1_FDMCFRA_r16_Eight},
}

const (
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_KHz1dot25 = 0
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_KHz5      = 1
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_Spare2    = 2
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_Spare1    = 3
)

var rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_KHz1dot25, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_KHz5, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_Spare2, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndex_r16_Spare1},
}

const (
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_KHz1dot25 = 0
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_KHz5      = 1
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_Spare2    = 2
	RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_Spare1    = 3
)

var rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexCFRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_KHz1dot25, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_KHz5, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_Spare2, RA_InformationCommon_r16_Ext_Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16_Spare1},
}

var rAInformationCommonR16MsgAROFrequencyStartR17Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

var rAInformationCommonR16MsgAROFrequencyStartCFRAR17Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

const (
	RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_One   = 0
	RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Two   = 1
	RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Four  = 2
	RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Eight = 3
)

var rAInformationCommonR16ExtMsgAROFDMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_One, RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Two, RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Four, RA_InformationCommon_r16_Ext_MsgA_RO_FDM_r17_Eight},
}

const (
	RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_One   = 0
	RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Two   = 1
	RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Four  = 2
	RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Eight = 3
)

var rAInformationCommonR16ExtMsgAROFDMCFRAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_One, RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Two, RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Four, RA_InformationCommon_r16_Ext_MsgA_RO_FDMCFRA_r17_Eight},
}

const (
	RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_KHz1dot25 = 0
	RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_KHz5      = 1
	RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_Spare2    = 2
	RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_Spare1    = 3
)

var rAInformationCommonR16ExtMsgASCSFromPrachConfigurationIndexR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_KHz1dot25, RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_KHz5, RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_Spare2, RA_InformationCommon_r16_Ext_MsgA_SCS_From_Prach_ConfigurationIndex_r17_Spare1},
}

const (
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N1   = 0
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N2   = 1
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N4   = 2
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N6   = 3
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N8   = 4
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N10  = 5
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N20  = 6
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N50  = 7
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N100 = 8
	RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N200 = 9
)

var rAInformationCommonR16ExtMsgATransMaxR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N1, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N2, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N4, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N6, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N8, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N10, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N20, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N50, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N100, RA_InformationCommon_r16_Ext_MsgA_TransMax_r17_N200},
}

var rAInformationCommonR16MsgAMCSR17Constraints = per.Constrained(0, 15)

var rAInformationCommonR16NrofPRBsPerMsgAPOR17Constraints = per.Constrained(1, 32)

var rAInformationCommonR16MsgAPUSCHTimeDomainAllocationR17Constraints = per.Constrained(1, common.MaxNrofUL_Allocations)

var rAInformationCommonR16FrequencyStartMsgAPUSCHR17Constraints = per.Constrained(0, common.MaxNrofPhysicalResourceBlocks_1)

const (
	RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_One   = 0
	RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Two   = 1
	RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Four  = 2
	RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Eight = 3
)

var rAInformationCommonR16ExtNrofMsgAPOFDMR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_One, RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Two, RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Four, RA_InformationCommon_r16_Ext_NrofMsgA_PO_FDM_r17_Eight},
}

var rAInformationCommonR16ExtIntendedSIBsR17Constraints = per.SizeRange(1, common.MaxSIB)

var rAInformationCommonR16ExtSsbsForSIAcquisitionR17Constraints = per.SizeRange(1, common.MaxNrofSSBs_r16)

var rAInformationCommonR16MsgAPUSCHPayloadSizeR17Constraints = per.FixedSize(5)

const (
	RA_InformationCommon_r16_Ext_OnDemandSISuccess_r17_True = 0
)

var rAInformationCommonR16ExtOnDemandSISuccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RA_InformationCommon_r16_Ext_OnDemandSISuccess_r17_True},
}

var rAInformationCommonR16StartPreambleForThisPartitionR18Constraints = per.Constrained(0, 63)

var rAInformationCommonR16NumberOfPreamblesPerSSBForThisPartitionR18Constraints = per.Constrained(1, 64)

var rAInformationCommonR16ExtAttemptedBWPInfoListR18Constraints = per.SizeRange(1, common.MaxNrofBWPs)

var rAInformationCommonR16NumberOfLBTFailuresR18Constraints = per.Constrained(1, 128)

var rAInformationCommonR16ExtIntendedSIBsR18Constraints = per.SizeRange(1, common.MaxSIB)

type RA_InformationCommon_r16 struct {
	AbsoluteFrequencyPointA_r16                    ARFCN_ValueNR
	LocationAndBandwidth_r16                       int64
	SubcarrierSpacing_r16                          SubcarrierSpacing
	Msg1_FrequencyStart_r16                        *int64
	Msg1_FrequencyStartCFRA_r16                    *int64
	Msg1_SubcarrierSpacing_r16                     *SubcarrierSpacing
	Msg1_SubcarrierSpacingCFRA_r16                 *SubcarrierSpacing
	Msg1_FDM_r16                                   *int64
	Msg1_FDMCFRA_r16                               *int64
	PerRAInfoList_r16                              PerRAInfoList_r16
	PerRAInfoList_v1660                            *PerRAInfoList_v1660
	Msg1_SCS_From_Prach_ConfigurationIndex_r16     *int64
	Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16 *int64
	MsgA_RO_FrequencyStart_r17                     *int64
	MsgA_RO_FrequencyStartCFRA_r17                 *int64
	MsgA_SubcarrierSpacing_r17                     *SubcarrierSpacing
	MsgA_RO_FDM_r17                                *int64
	MsgA_RO_FDMCFRA_r17                            *int64
	MsgA_SCS_From_Prach_ConfigurationIndex_r17     *int64
	MsgA_TransMax_r17                              *int64
	MsgA_MCS_r17                                   *int64
	NrofPRBs_PerMsgA_PO_r17                        *int64
	MsgA_PUSCH_TimeDomainAllocation_r17            *int64
	FrequencyStartMsgA_PUSCH_r17                   *int64
	NrofMsgA_PO_FDM_r17                            *int64
	DlPathlossRSRP_r17                             *RSRP_Range
	IntendedSIBs_r17                               []SIB_Type_r17
	SsbsForSI_Acquisition_r17                      []SSB_Index
	MsgA_PUSCH_PayloadSize_r17                     *per.BitString
	OnDemandSISuccess_r17                          *int64
	UsedFeatureCombination_r18                     *ReportedFeatureCombination_r18
	TriggeredFeatureCombination_r18                *ReportedFeatureCombination_r18
	StartPreambleForThisPartition_r18              *int64
	NumberOfPreamblesPerSSB_ForThisPartition_r18   *int64
	AttemptedBWP_InfoList_r18                      []AttemptedBWP_Info_r18
	NumberOfLBT_Failures_r18                       *int64
	PerRAInfoList_v1800                            *PerRAInfoList_v1800
	IntendedSIBs_r18                               []SIB_Type_r18
}

func (ie *RA_InformationCommon_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rAInformationCommonR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PerRAInfoList_v1660 != nil
	hasExtGroup1 := ie.Msg1_SCS_From_Prach_ConfigurationIndex_r16 != nil
	hasExtGroup2 := ie.Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16 != nil
	hasExtGroup3 := ie.MsgA_RO_FrequencyStart_r17 != nil || ie.MsgA_RO_FrequencyStartCFRA_r17 != nil || ie.MsgA_SubcarrierSpacing_r17 != nil || ie.MsgA_RO_FDM_r17 != nil || ie.MsgA_RO_FDMCFRA_r17 != nil || ie.MsgA_SCS_From_Prach_ConfigurationIndex_r17 != nil || ie.MsgA_TransMax_r17 != nil || ie.MsgA_MCS_r17 != nil || ie.NrofPRBs_PerMsgA_PO_r17 != nil || ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil || ie.FrequencyStartMsgA_PUSCH_r17 != nil || ie.NrofMsgA_PO_FDM_r17 != nil || ie.DlPathlossRSRP_r17 != nil || ie.IntendedSIBs_r17 != nil || ie.SsbsForSI_Acquisition_r17 != nil || ie.MsgA_PUSCH_PayloadSize_r17 != nil || ie.OnDemandSISuccess_r17 != nil
	hasExtGroup4 := ie.UsedFeatureCombination_r18 != nil || ie.TriggeredFeatureCombination_r18 != nil || ie.StartPreambleForThisPartition_r18 != nil || ie.NumberOfPreamblesPerSSB_ForThisPartition_r18 != nil || ie.AttemptedBWP_InfoList_r18 != nil || ie.NumberOfLBT_Failures_r18 != nil || ie.PerRAInfoList_v1800 != nil || ie.IntendedSIBs_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Msg1_FrequencyStart_r16 != nil, ie.Msg1_FrequencyStartCFRA_r16 != nil, ie.Msg1_SubcarrierSpacing_r16 != nil, ie.Msg1_SubcarrierSpacingCFRA_r16 != nil, ie.Msg1_FDM_r16 != nil, ie.Msg1_FDMCFRA_r16 != nil}); err != nil {
		return err
	}

	// 3. absoluteFrequencyPointA-r16: ref
	{
		if err := ie.AbsoluteFrequencyPointA_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. locationAndBandwidth-r16: integer
	{
		if err := e.EncodeInteger(ie.LocationAndBandwidth_r16, rAInformationCommonR16LocationAndBandwidthR16Constraints); err != nil {
			return err
		}
	}

	// 5. subcarrierSpacing-r16: ref
	{
		if err := ie.SubcarrierSpacing_r16.Encode(e); err != nil {
			return err
		}
	}

	// 6. msg1-FrequencyStart-r16: integer
	{
		if ie.Msg1_FrequencyStart_r16 != nil {
			if err := e.EncodeInteger(*ie.Msg1_FrequencyStart_r16, rAInformationCommonR16Msg1FrequencyStartR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. msg1-FrequencyStartCFRA-r16: integer
	{
		if ie.Msg1_FrequencyStartCFRA_r16 != nil {
			if err := e.EncodeInteger(*ie.Msg1_FrequencyStartCFRA_r16, rAInformationCommonR16Msg1FrequencyStartCFRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. msg1-SubcarrierSpacing-r16: ref
	{
		if ie.Msg1_SubcarrierSpacing_r16 != nil {
			if err := ie.Msg1_SubcarrierSpacing_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. msg1-SubcarrierSpacingCFRA-r16: ref
	{
		if ie.Msg1_SubcarrierSpacingCFRA_r16 != nil {
			if err := ie.Msg1_SubcarrierSpacingCFRA_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. msg1-FDM-r16: enumerated
	{
		if ie.Msg1_FDM_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Msg1_FDM_r16, rAInformationCommonR16Msg1FDMR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. msg1-FDMCFRA-r16: enumerated
	{
		if ie.Msg1_FDMCFRA_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Msg1_FDMCFRA_r16, rAInformationCommonR16Msg1FDMCFRAR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. perRAInfoList-r16: ref
	{
		if err := ie.PerRAInfoList_r16.Encode(e); err != nil {
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
					{Name: "perRAInfoList-v1660", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PerRAInfoList_v1660 != nil}); err != nil {
				return err
			}

			if ie.PerRAInfoList_v1660 != nil {
				if err := ie.PerRAInfoList_v1660.Encode(ex); err != nil {
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
					{Name: "msg1-SCS-From-prach-ConfigurationIndex-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg1_SCS_From_Prach_ConfigurationIndex_r16 != nil}); err != nil {
				return err
			}

			if ie.Msg1_SCS_From_Prach_ConfigurationIndex_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Msg1_SCS_From_Prach_ConfigurationIndex_r16, rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexR16Constraints); err != nil {
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
					{Name: "msg1-SCS-From-prach-ConfigurationIndexCFRA-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16 != nil}); err != nil {
				return err
			}

			if ie.Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16, rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexCFRAR16Constraints); err != nil {
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
					{Name: "msgA-RO-FrequencyStart-r17", Optional: true},
					{Name: "msgA-RO-FrequencyStartCFRA-r17", Optional: true},
					{Name: "msgA-SubcarrierSpacing-r17", Optional: true},
					{Name: "msgA-RO-FDM-r17", Optional: true},
					{Name: "msgA-RO-FDMCFRA-r17", Optional: true},
					{Name: "msgA-SCS-From-prach-ConfigurationIndex-r17", Optional: true},
					{Name: "msgA-TransMax-r17", Optional: true},
					{Name: "msgA-MCS-r17", Optional: true},
					{Name: "nrofPRBs-PerMsgA-PO-r17", Optional: true},
					{Name: "msgA-PUSCH-TimeDomainAllocation-r17", Optional: true},
					{Name: "frequencyStartMsgA-PUSCH-r17", Optional: true},
					{Name: "nrofMsgA-PO-FDM-r17", Optional: true},
					{Name: "dlPathlossRSRP-r17", Optional: true},
					{Name: "intendedSIBs-r17", Optional: true},
					{Name: "ssbsForSI-Acquisition-r17", Optional: true},
					{Name: "msgA-PUSCH-PayloadSize-r17", Optional: true},
					{Name: "onDemandSISuccess-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MsgA_RO_FrequencyStart_r17 != nil, ie.MsgA_RO_FrequencyStartCFRA_r17 != nil, ie.MsgA_SubcarrierSpacing_r17 != nil, ie.MsgA_RO_FDM_r17 != nil, ie.MsgA_RO_FDMCFRA_r17 != nil, ie.MsgA_SCS_From_Prach_ConfigurationIndex_r17 != nil, ie.MsgA_TransMax_r17 != nil, ie.MsgA_MCS_r17 != nil, ie.NrofPRBs_PerMsgA_PO_r17 != nil, ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil, ie.FrequencyStartMsgA_PUSCH_r17 != nil, ie.NrofMsgA_PO_FDM_r17 != nil, ie.DlPathlossRSRP_r17 != nil, ie.IntendedSIBs_r17 != nil, ie.SsbsForSI_Acquisition_r17 != nil, ie.MsgA_PUSCH_PayloadSize_r17 != nil, ie.OnDemandSISuccess_r17 != nil}); err != nil {
				return err
			}

			if ie.MsgA_RO_FrequencyStart_r17 != nil {
				if err := ex.EncodeInteger(*ie.MsgA_RO_FrequencyStart_r17, rAInformationCommonR16MsgAROFrequencyStartR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_RO_FrequencyStartCFRA_r17 != nil {
				if err := ex.EncodeInteger(*ie.MsgA_RO_FrequencyStartCFRA_r17, rAInformationCommonR16MsgAROFrequencyStartCFRAR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_SubcarrierSpacing_r17 != nil {
				if err := ie.MsgA_SubcarrierSpacing_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MsgA_RO_FDM_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MsgA_RO_FDM_r17, rAInformationCommonR16ExtMsgAROFDMR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_RO_FDMCFRA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MsgA_RO_FDMCFRA_r17, rAInformationCommonR16ExtMsgAROFDMCFRAR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_SCS_From_Prach_ConfigurationIndex_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MsgA_SCS_From_Prach_ConfigurationIndex_r17, rAInformationCommonR16ExtMsgASCSFromPrachConfigurationIndexR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_TransMax_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MsgA_TransMax_r17, rAInformationCommonR16ExtMsgATransMaxR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_MCS_r17 != nil {
				if err := ex.EncodeInteger(*ie.MsgA_MCS_r17, rAInformationCommonR16MsgAMCSR17Constraints); err != nil {
					return err
				}
			}

			if ie.NrofPRBs_PerMsgA_PO_r17 != nil {
				if err := ex.EncodeInteger(*ie.NrofPRBs_PerMsgA_PO_r17, rAInformationCommonR16NrofPRBsPerMsgAPOR17Constraints); err != nil {
					return err
				}
			}

			if ie.MsgA_PUSCH_TimeDomainAllocation_r17 != nil {
				if err := ex.EncodeInteger(*ie.MsgA_PUSCH_TimeDomainAllocation_r17, rAInformationCommonR16MsgAPUSCHTimeDomainAllocationR17Constraints); err != nil {
					return err
				}
			}

			if ie.FrequencyStartMsgA_PUSCH_r17 != nil {
				if err := ex.EncodeInteger(*ie.FrequencyStartMsgA_PUSCH_r17, rAInformationCommonR16FrequencyStartMsgAPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.NrofMsgA_PO_FDM_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofMsgA_PO_FDM_r17, rAInformationCommonR16ExtNrofMsgAPOFDMR17Constraints); err != nil {
					return err
				}
			}

			if ie.DlPathlossRSRP_r17 != nil {
				if err := ie.DlPathlossRSRP_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IntendedSIBs_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(rAInformationCommonR16ExtIntendedSIBsR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.IntendedSIBs_r17))); err != nil {
					return err
				}
				for i := range ie.IntendedSIBs_r17 {
					if err := ie.IntendedSIBs_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SsbsForSI_Acquisition_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(rAInformationCommonR16ExtSsbsForSIAcquisitionR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.SsbsForSI_Acquisition_r17))); err != nil {
					return err
				}
				for i := range ie.SsbsForSI_Acquisition_r17 {
					if err := ie.SsbsForSI_Acquisition_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MsgA_PUSCH_PayloadSize_r17 != nil {
				if err := ex.EncodeBitString(*ie.MsgA_PUSCH_PayloadSize_r17, rAInformationCommonR16MsgAPUSCHPayloadSizeR17Constraints); err != nil {
					return err
				}
			}

			if ie.OnDemandSISuccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.OnDemandSISuccess_r17, rAInformationCommonR16ExtOnDemandSISuccessR17Constraints); err != nil {
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
					{Name: "usedFeatureCombination-r18", Optional: true},
					{Name: "triggeredFeatureCombination-r18", Optional: true},
					{Name: "startPreambleForThisPartition-r18", Optional: true},
					{Name: "numberOfPreamblesPerSSB-ForThisPartition-r18", Optional: true},
					{Name: "attemptedBWP-InfoList-r18", Optional: true},
					{Name: "numberOfLBT-Failures-r18", Optional: true},
					{Name: "perRAInfoList-v1800", Optional: true},
					{Name: "intendedSIBs-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UsedFeatureCombination_r18 != nil, ie.TriggeredFeatureCombination_r18 != nil, ie.StartPreambleForThisPartition_r18 != nil, ie.NumberOfPreamblesPerSSB_ForThisPartition_r18 != nil, ie.AttemptedBWP_InfoList_r18 != nil, ie.NumberOfLBT_Failures_r18 != nil, ie.PerRAInfoList_v1800 != nil, ie.IntendedSIBs_r18 != nil}); err != nil {
				return err
			}

			if ie.UsedFeatureCombination_r18 != nil {
				if err := ie.UsedFeatureCombination_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.TriggeredFeatureCombination_r18 != nil {
				if err := ie.TriggeredFeatureCombination_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.StartPreambleForThisPartition_r18 != nil {
				if err := ex.EncodeInteger(*ie.StartPreambleForThisPartition_r18, rAInformationCommonR16StartPreambleForThisPartitionR18Constraints); err != nil {
					return err
				}
			}

			if ie.NumberOfPreamblesPerSSB_ForThisPartition_r18 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfPreamblesPerSSB_ForThisPartition_r18, rAInformationCommonR16NumberOfPreamblesPerSSBForThisPartitionR18Constraints); err != nil {
					return err
				}
			}

			if ie.AttemptedBWP_InfoList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(rAInformationCommonR16ExtAttemptedBWPInfoListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AttemptedBWP_InfoList_r18))); err != nil {
					return err
				}
				for i := range ie.AttemptedBWP_InfoList_r18 {
					if err := ie.AttemptedBWP_InfoList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NumberOfLBT_Failures_r18 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfLBT_Failures_r18, rAInformationCommonR16NumberOfLBTFailuresR18Constraints); err != nil {
					return err
				}
			}

			if ie.PerRAInfoList_v1800 != nil {
				if err := ie.PerRAInfoList_v1800.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IntendedSIBs_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(rAInformationCommonR16ExtIntendedSIBsR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.IntendedSIBs_r18))); err != nil {
					return err
				}
				for i := range ie.IntendedSIBs_r18 {
					if err := ie.IntendedSIBs_r18[i].Encode(ex); err != nil {
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

func (ie *RA_InformationCommon_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rAInformationCommonR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. absoluteFrequencyPointA-r16: ref
	{
		if err := ie.AbsoluteFrequencyPointA_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. locationAndBandwidth-r16: integer
	{
		v1, err := d.DecodeInteger(rAInformationCommonR16LocationAndBandwidthR16Constraints)
		if err != nil {
			return err
		}
		ie.LocationAndBandwidth_r16 = v1
	}

	// 5. subcarrierSpacing-r16: ref
	{
		if err := ie.SubcarrierSpacing_r16.Decode(d); err != nil {
			return err
		}
	}

	// 6. msg1-FrequencyStart-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(rAInformationCommonR16Msg1FrequencyStartR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_FrequencyStart_r16 = &v
		}
	}

	// 7. msg1-FrequencyStartCFRA-r16: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(rAInformationCommonR16Msg1FrequencyStartCFRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_FrequencyStartCFRA_r16 = &v
		}
	}

	// 8. msg1-SubcarrierSpacing-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Msg1_SubcarrierSpacing_r16 = new(SubcarrierSpacing)
			if err := ie.Msg1_SubcarrierSpacing_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. msg1-SubcarrierSpacingCFRA-r16: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Msg1_SubcarrierSpacingCFRA_r16 = new(SubcarrierSpacing)
			if err := ie.Msg1_SubcarrierSpacingCFRA_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. msg1-FDM-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(rAInformationCommonR16Msg1FDMR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_FDM_r16 = &idx
		}
	}

	// 11. msg1-FDMCFRA-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(rAInformationCommonR16Msg1FDMCFRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_FDMCFRA_r16 = &idx
		}
	}

	// 12. perRAInfoList-r16: ref
	{
		if err := ie.PerRAInfoList_r16.Decode(d); err != nil {
			return err
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
				{Name: "perRAInfoList-v1660", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PerRAInfoList_v1660 = new(PerRAInfoList_v1660)
			if err := ie.PerRAInfoList_v1660.Decode(dx); err != nil {
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
				{Name: "msg1-SCS-From-prach-ConfigurationIndex-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_SCS_From_Prach_ConfigurationIndex_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "msg1-SCS-From-prach-ConfigurationIndexCFRA-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsg1SCSFromPrachConfigurationIndexCFRAR16Constraints)
			if err != nil {
				return err
			}
			ie.Msg1_SCS_From_Prach_ConfigurationIndexCFRA_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "msgA-RO-FrequencyStart-r17", Optional: true},
				{Name: "msgA-RO-FrequencyStartCFRA-r17", Optional: true},
				{Name: "msgA-SubcarrierSpacing-r17", Optional: true},
				{Name: "msgA-RO-FDM-r17", Optional: true},
				{Name: "msgA-RO-FDMCFRA-r17", Optional: true},
				{Name: "msgA-SCS-From-prach-ConfigurationIndex-r17", Optional: true},
				{Name: "msgA-TransMax-r17", Optional: true},
				{Name: "msgA-MCS-r17", Optional: true},
				{Name: "nrofPRBs-PerMsgA-PO-r17", Optional: true},
				{Name: "msgA-PUSCH-TimeDomainAllocation-r17", Optional: true},
				{Name: "frequencyStartMsgA-PUSCH-r17", Optional: true},
				{Name: "nrofMsgA-PO-FDM-r17", Optional: true},
				{Name: "dlPathlossRSRP-r17", Optional: true},
				{Name: "intendedSIBs-r17", Optional: true},
				{Name: "ssbsForSI-Acquisition-r17", Optional: true},
				{Name: "msgA-PUSCH-PayloadSize-r17", Optional: true},
				{Name: "onDemandSISuccess-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(rAInformationCommonR16MsgAROFrequencyStartR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FrequencyStart_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(rAInformationCommonR16MsgAROFrequencyStartCFRAR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FrequencyStartCFRA_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MsgA_SubcarrierSpacing_r17 = new(SubcarrierSpacing)
			if err := ie.MsgA_SubcarrierSpacing_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsgAROFDMR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FDM_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsgAROFDMCFRAR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_RO_FDMCFRA_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsgASCSFromPrachConfigurationIndexR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_SCS_From_Prach_ConfigurationIndex_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtMsgATransMaxR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_TransMax_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeInteger(rAInformationCommonR16MsgAMCSR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_MCS_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeInteger(rAInformationCommonR16NrofPRBsPerMsgAPOR17Constraints)
			if err != nil {
				return err
			}
			ie.NrofPRBs_PerMsgA_PO_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeInteger(rAInformationCommonR16MsgAPUSCHTimeDomainAllocationR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_TimeDomainAllocation_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeInteger(rAInformationCommonR16FrequencyStartMsgAPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.FrequencyStartMsgA_PUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtNrofMsgAPOFDMR17Constraints)
			if err != nil {
				return err
			}
			ie.NrofMsgA_PO_FDM_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			ie.DlPathlossRSRP_r17 = new(RSRP_Range)
			if err := ie.DlPathlossRSRP_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(13) {
			seqOf := dx.NewSequenceOfDecoder(rAInformationCommonR16ExtIntendedSIBsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntendedSIBs_r17 = make([]SIB_Type_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntendedSIBs_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(14) {
			seqOf := dx.NewSequenceOfDecoder(rAInformationCommonR16ExtSsbsForSIAcquisitionR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SsbsForSI_Acquisition_r17 = make([]SSB_Index, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SsbsForSI_Acquisition_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeBitString(rAInformationCommonR16MsgAPUSCHPayloadSizeR17Constraints)
			if err != nil {
				return err
			}
			ie.MsgA_PUSCH_PayloadSize_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(rAInformationCommonR16ExtOnDemandSISuccessR17Constraints)
			if err != nil {
				return err
			}
			ie.OnDemandSISuccess_r17 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "usedFeatureCombination-r18", Optional: true},
				{Name: "triggeredFeatureCombination-r18", Optional: true},
				{Name: "startPreambleForThisPartition-r18", Optional: true},
				{Name: "numberOfPreamblesPerSSB-ForThisPartition-r18", Optional: true},
				{Name: "attemptedBWP-InfoList-r18", Optional: true},
				{Name: "numberOfLBT-Failures-r18", Optional: true},
				{Name: "perRAInfoList-v1800", Optional: true},
				{Name: "intendedSIBs-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.UsedFeatureCombination_r18 = new(ReportedFeatureCombination_r18)
			if err := ie.UsedFeatureCombination_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.TriggeredFeatureCombination_r18 = new(ReportedFeatureCombination_r18)
			if err := ie.TriggeredFeatureCombination_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(rAInformationCommonR16StartPreambleForThisPartitionR18Constraints)
			if err != nil {
				return err
			}
			ie.StartPreambleForThisPartition_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(rAInformationCommonR16NumberOfPreamblesPerSSBForThisPartitionR18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfPreamblesPerSSB_ForThisPartition_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(rAInformationCommonR16ExtAttemptedBWPInfoListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AttemptedBWP_InfoList_r18 = make([]AttemptedBWP_Info_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AttemptedBWP_InfoList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeInteger(rAInformationCommonR16NumberOfLBTFailuresR18Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfLBT_Failures_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.PerRAInfoList_v1800 = new(PerRAInfoList_v1800)
			if err := ie.PerRAInfoList_v1800.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(rAInformationCommonR16ExtIntendedSIBsR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntendedSIBs_r18 = make([]SIB_Type_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntendedSIBs_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
