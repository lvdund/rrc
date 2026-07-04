// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-ResourcePool-r18 (line 27616).

var sLPRSResourcePoolR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-PSCCH-Config-r18", Optional: true},
		{Name: "sl-StartRB-SubchannelDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-FilterCoefficient-r18", Optional: true},
		{Name: "sl-ThreshS-RSSI-PRS-CBR-r18", Optional: true},
		{Name: "sl-RB-Number-r18", Optional: true},
		{Name: "sl-TimeResource-r18", Optional: true},
		{Name: "sl-PosAllowedResourceSelectionConfig-r18", Optional: true},
		{Name: "sl-PRS-ResourceReservePeriodList-r18", Optional: true},
		{Name: "sl-PRS-ResourcesDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-PRS-PowerControl-r18", Optional: true},
		{Name: "sl-SensingWindowDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-TxPercentageDedicatedSL-PRS-RP-List-r18", Optional: true},
		{Name: "sl-SCI-basedSL-PRS-TxTriggerSCI1-B-r18", Optional: true},
		{Name: "sl-NumSubchannelDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-SubchannelSizeDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-MaxNumPerReserveDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-NumReservedBitsSCI1B-DedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-SRC-ID-LenDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-CBR-PriorityTxConfigDedicatedSL-PRS-RP-List-r18", Optional: true},
		{Name: "sl-TimeWindowSizeCBR-DedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-TimeWindowSizeCR-DedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-CBR-CommonTxDedicatedSL-PRS-RP-List-r18", Optional: true},
		{Name: "sl-PriorityThreshold-UL-URLLC-r18", Optional: true},
		{Name: "sl-PriorityThreshold-r18", Optional: true},
		{Name: "sl-SelectionWindowListDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-Thres-RSRP-ListDedicatedSL-PRS-RP-r18", Optional: true},
		{Name: "sl-PreemptionEnableDedicatedSL-PRS-RP-r18", Optional: true},
	},
}

var sL_PRS_ResourcePool_r18SlPRSPSCCHConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Release = 0
	SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Setup   = 1
)

var sLPRSResourcePoolR18SlStartRBSubchannelDedicatedSLPRSRPR18Constraints = per.Constrained(0, 265)

var sLPRSResourcePoolR18SlThreshSRSSIPRSCBRR18Constraints = per.Constrained(0, 45)

var sLPRSResourcePoolR18SlRBNumberR18Constraints = per.Constrained(10, 275)

var sLPRSResourcePoolR18SlTimeResourceR18Constraints = per.SizeRange(10, 160)

const (
	SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C1 = 0
	SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C2 = 1
	SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C3 = 2
)

var sLPRSResourcePoolR18SlPosAllowedResourceSelectionConfigR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C1, SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C2, SL_PRS_ResourcePool_r18_Sl_PosAllowedResourceSelectionConfig_r18_C3},
}

var sLPRSResourcePoolR18SlPRSResourceReservePeriodListR18Constraints = per.SizeRange(1, 16)

var sLPRSResourcePoolR18SlPRSResourcesDedicatedSLPRSRPR18Constraints = per.SizeRange(1, 12)

const (
	SL_PRS_ResourcePool_r18_Sl_SensingWindowDedicatedSL_PRS_RP_r18_Ms100  = 0
	SL_PRS_ResourcePool_r18_Sl_SensingWindowDedicatedSL_PRS_RP_r18_Ms1100 = 1
)

var sLPRSResourcePoolR18SlSensingWindowDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_SensingWindowDedicatedSL_PRS_RP_r18_Ms100, SL_PRS_ResourcePool_r18_Sl_SensingWindowDedicatedSL_PRS_RP_r18_Ms1100},
}

var sLPRSResourcePoolR18SlTxPercentageDedicatedSLPRSRPListR18Constraints = per.FixedSize(8)

var sLPRSResourcePoolR18SlNumSubchannelDedicatedSLPRSRPR18Constraints = per.Constrained(1, 27)

const (
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N10  = 0
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N12  = 1
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N15  = 2
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N20  = 3
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N25  = 4
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N50  = 5
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N75  = 6
	SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N100 = 7
)

var sLPRSResourcePoolR18SlSubchannelSizeDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N10, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N12, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N15, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N20, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N25, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N50, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N75, SL_PRS_ResourcePool_r18_Sl_SubchannelSizeDedicatedSL_PRS_RP_r18_N100},
}

const (
	SL_PRS_ResourcePool_r18_Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18_N2 = 0
	SL_PRS_ResourcePool_r18_Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18_N3 = 1
)

var sLPRSResourcePoolR18SlMaxNumPerReserveDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18_N2, SL_PRS_ResourcePool_r18_Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18_N3},
}

var sLPRSResourcePoolR18SlNumReservedBitsSCI1BDedicatedSLPRSRPR18Constraints = per.Constrained(0, 20)

const (
	SL_PRS_ResourcePool_r18_Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18_N12 = 0
	SL_PRS_ResourcePool_r18_Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18_N24 = 1
)

var sLPRSResourcePoolR18SlSRCIDLenDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18_N12, SL_PRS_ResourcePool_r18_Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18_N24},
}

var sLPRSResourcePoolR18SlCBRPriorityTxConfigDedicatedSLPRSRPListR18Constraints = per.SizeRange(1, 8)

const (
	SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18_Ms100   = 0
	SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18_Slot100 = 1
)

var sLPRSResourcePoolR18SlTimeWindowSizeCBRDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18_Ms100, SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18_Slot100},
}

const (
	SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18_Ms1000   = 0
	SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18_Slot1000 = 1
)

var sLPRSResourcePoolR18SlTimeWindowSizeCRDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18_Ms1000, SL_PRS_ResourcePool_r18_Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18_Slot1000},
}

var sLPRSResourcePoolR18SlPriorityThresholdULURLLCR18Constraints = per.Constrained(1, 9)

var sLPRSResourcePoolR18SlPriorityThresholdR18Constraints = per.Constrained(1, 9)

var sLPRSResourcePoolR18SlSelectionWindowListDedicatedSLPRSRPR18Constraints = per.FixedSize(8)

var sLPRSResourcePoolR18SlThresRSRPListDedicatedSLPRSRPR18Constraints = per.FixedSize(64)

const (
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Enabled = 0
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl1     = 1
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl2     = 2
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl3     = 3
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl4     = 4
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl5     = 5
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl6     = 6
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl7     = 7
	SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl8     = 8
)

var sLPRSResourcePoolR18SlPreemptionEnableDedicatedSLPRSRPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Enabled, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl1, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl2, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl3, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl4, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl5, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl6, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl7, SL_PRS_ResourcePool_r18_Sl_PreemptionEnableDedicatedSL_PRS_RP_r18_Pl8},
}

type SL_PRS_ResourcePool_r18 struct {
	Sl_PRS_PSCCH_Config_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18
	}
	Sl_StartRB_SubchannelDedicatedSL_PRS_RP_r18        *int64
	Sl_FilterCoefficient_r18                           *FilterCoefficient
	Sl_ThreshS_RSSI_PRS_CBR_r18                        *int64
	Sl_RB_Number_r18                                   *int64
	Sl_TimeResource_r18                                *per.BitString
	Sl_PosAllowedResourceSelectionConfig_r18           *int64
	Sl_PRS_ResourceReservePeriodList_r18               []SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18
	Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18             []SL_PRS_ResourceDedicatedSL_PRS_RP_r18
	Sl_PRS_PowerControl_r18                            *SL_PRS_PowerControl_r18
	Sl_SensingWindowDedicatedSL_PRS_RP_r18             *int64
	Sl_TxPercentageDedicatedSL_PRS_RP_List_r18         []SL_TxPercentageDedicatedSL_PRS_RP_Config_r18
	Sl_SCI_BasedSL_PRS_TxTriggerSCI1_B_r18             *bool
	Sl_NumSubchannelDedicatedSL_PRS_RP_r18             *int64
	Sl_SubchannelSizeDedicatedSL_PRS_RP_r18            *int64
	Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18          *int64
	Sl_NumReservedBitsSCI1B_DedicatedSL_PRS_RP_r18     *int64
	Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18                *int64
	Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18 []SL_PriorityTxConfigIndexDedicatedSL_PRS_RP_r18
	Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18        *int64
	Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18         *int64
	Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18         *SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18
	Sl_PriorityThreshold_UL_URLLC_r18                  *int64
	Sl_PriorityThreshold_r18                           *int64
	Sl_SelectionWindowListDedicatedSL_PRS_RP_r18       []SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18
	Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18           []SL_PRS_ThresRSRP_r18
	Sl_PreemptionEnableDedicatedSL_PRS_RP_r18          *int64
}

func (ie *SL_PRS_ResourcePool_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSResourcePoolR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_PSCCH_Config_r18 != nil, ie.Sl_StartRB_SubchannelDedicatedSL_PRS_RP_r18 != nil, ie.Sl_FilterCoefficient_r18 != nil, ie.Sl_ThreshS_RSSI_PRS_CBR_r18 != nil, ie.Sl_RB_Number_r18 != nil, ie.Sl_TimeResource_r18 != nil, ie.Sl_PosAllowedResourceSelectionConfig_r18 != nil, ie.Sl_PRS_ResourceReservePeriodList_r18 != nil, ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18 != nil, ie.Sl_PRS_PowerControl_r18 != nil, ie.Sl_SensingWindowDedicatedSL_PRS_RP_r18 != nil, ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18 != nil, ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI1_B_r18 != nil, ie.Sl_NumSubchannelDedicatedSL_PRS_RP_r18 != nil, ie.Sl_SubchannelSizeDedicatedSL_PRS_RP_r18 != nil, ie.Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18 != nil, ie.Sl_NumReservedBitsSCI1B_DedicatedSL_PRS_RP_r18 != nil, ie.Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18 != nil, ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18 != nil, ie.Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18 != nil, ie.Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18 != nil, ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 != nil, ie.Sl_PriorityThreshold_UL_URLLC_r18 != nil, ie.Sl_PriorityThreshold_r18 != nil, ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18 != nil, ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18 != nil, ie.Sl_PreemptionEnableDedicatedSL_PRS_RP_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-PSCCH-Config-r18: choice
	{
		if ie.Sl_PRS_PSCCH_Config_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_PRS_ResourcePool_r18SlPRSPSCCHConfigR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PRS_PSCCH_Config_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PRS_PSCCH_Config_r18).Choice {
			case SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Setup:
				if err := (*ie.Sl_PRS_PSCCH_Config_r18).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PRS_PSCCH_Config_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 3. sl-StartRB-SubchannelDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_StartRB_SubchannelDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_StartRB_SubchannelDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlStartRBSubchannelDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-FilterCoefficient-r18: ref
	{
		if ie.Sl_FilterCoefficient_r18 != nil {
			if err := ie.Sl_FilterCoefficient_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. sl-ThreshS-RSSI-PRS-CBR-r18: integer
	{
		if ie.Sl_ThreshS_RSSI_PRS_CBR_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_ThreshS_RSSI_PRS_CBR_r18, sLPRSResourcePoolR18SlThreshSRSSIPRSCBRR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-RB-Number-r18: integer
	{
		if ie.Sl_RB_Number_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_RB_Number_r18, sLPRSResourcePoolR18SlRBNumberR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. sl-TimeResource-r18: bit-string
	{
		if ie.Sl_TimeResource_r18 != nil {
			if err := e.EncodeBitString(*ie.Sl_TimeResource_r18, sLPRSResourcePoolR18SlTimeResourceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. sl-PosAllowedResourceSelectionConfig-r18: enumerated
	{
		if ie.Sl_PosAllowedResourceSelectionConfig_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PosAllowedResourceSelectionConfig_r18, sLPRSResourcePoolR18SlPosAllowedResourceSelectionConfigR18Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sl-PRS-ResourceReservePeriodList-r18: sequence-of
	{
		if ie.Sl_PRS_ResourceReservePeriodList_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlPRSResourceReservePeriodListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_ResourceReservePeriodList_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_ResourceReservePeriodList_r18 {
				if err := ie.Sl_PRS_ResourceReservePeriodList_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-PRS-ResourcesDedicatedSL-PRS-RP-r18: sequence-of
	{
		if ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlPRSResourcesDedicatedSLPRSRPR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18 {
				if err := ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 11. sl-PRS-PowerControl-r18: ref
	{
		if ie.Sl_PRS_PowerControl_r18 != nil {
			if err := ie.Sl_PRS_PowerControl_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. sl-SensingWindowDedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_SensingWindowDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SensingWindowDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlSensingWindowDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. sl-TxPercentageDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlTxPercentageDedicatedSLPRSRPListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18 {
				if err := ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 14. sl-SCI-basedSL-PRS-TxTriggerSCI1-B-r18: boolean
	{
		if ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI1_B_r18 != nil {
			if err := e.EncodeBoolean(*ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI1_B_r18); err != nil {
				return err
			}
		}
	}

	// 15. sl-NumSubchannelDedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_NumSubchannelDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumSubchannelDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlNumSubchannelDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 16. sl-SubchannelSizeDedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_SubchannelSizeDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SubchannelSizeDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlSubchannelSizeDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 17. sl-MaxNumPerReserveDedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlMaxNumPerReserveDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 18. sl-NumReservedBitsSCI1B-DedicatedSL-PRS-RP-r18: integer
	{
		if ie.Sl_NumReservedBitsSCI1B_DedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumReservedBitsSCI1B_DedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlNumReservedBitsSCI1BDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 19. sl-SRC-ID-LenDedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlSRCIDLenDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 20. sl-CBR-PriorityTxConfigDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlCBRPriorityTxConfigDedicatedSLPRSRPListR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18 {
				if err := ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 21. sl-TimeWindowSizeCBR-DedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlTimeWindowSizeCBRDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 22. sl-TimeWindowSizeCR-DedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlTimeWindowSizeCRDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 23. sl-CBR-CommonTxDedicatedSL-PRS-RP-List-r18: ref
	{
		if ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 != nil {
			if err := ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 24. sl-PriorityThreshold-UL-URLLC-r18: integer
	{
		if ie.Sl_PriorityThreshold_UL_URLLC_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThreshold_UL_URLLC_r18, sLPRSResourcePoolR18SlPriorityThresholdULURLLCR18Constraints); err != nil {
				return err
			}
		}
	}

	// 25. sl-PriorityThreshold-r18: integer
	{
		if ie.Sl_PriorityThreshold_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PriorityThreshold_r18, sLPRSResourcePoolR18SlPriorityThresholdR18Constraints); err != nil {
				return err
			}
		}
	}

	// 26. sl-SelectionWindowListDedicatedSL-PRS-RP-r18: sequence-of
	{
		if ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlSelectionWindowListDedicatedSLPRSRPR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18 {
				if err := ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 27. sl-Thres-RSRP-ListDedicatedSL-PRS-RP-r18: sequence-of
	{
		if ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18 != nil {
			seqOf := e.NewSequenceOfEncoder(sLPRSResourcePoolR18SlThresRSRPListDedicatedSLPRSRPR18Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18))); err != nil {
				return err
			}
			for i := range ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18 {
				if err := ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 28. sl-PreemptionEnableDedicatedSL-PRS-RP-r18: enumerated
	{
		if ie.Sl_PreemptionEnableDedicatedSL_PRS_RP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_PreemptionEnableDedicatedSL_PRS_RP_r18, sLPRSResourcePoolR18SlPreemptionEnableDedicatedSLPRSRPR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PRS_ResourcePool_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSResourcePoolR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-PSCCH-Config-r18: choice
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_PRS_PSCCH_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18
			}{}
			choiceDec := d.NewChoiceDecoder(sL_PRS_ResourcePool_r18SlPRSPSCCHConfigR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PRS_PSCCH_Config_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Release:
				(*ie.Sl_PRS_PSCCH_Config_r18).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case SL_PRS_ResourcePool_r18_Sl_PRS_PSCCH_Config_r18_Setup:
				(*ie.Sl_PRS_PSCCH_Config_r18).Setup = new(SL_PSCCH_ConfigDedicatedSL_PRS_RP_r18)
				if err := (*ie.Sl_PRS_PSCCH_Config_r18).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. sl-StartRB-SubchannelDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlStartRBSubchannelDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_StartRB_SubchannelDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 4. sl-FilterCoefficient-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_FilterCoefficient_r18 = new(FilterCoefficient)
			if err := ie.Sl_FilterCoefficient_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. sl-ThreshS-RSSI-PRS-CBR-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlThreshSRSSIPRSCBRR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ThreshS_RSSI_PRS_CBR_r18 = &v
		}
	}

	// 6. sl-RB-Number-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlRBNumberR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_RB_Number_r18 = &v
		}
	}

	// 7. sl-TimeResource-r18: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(sLPRSResourcePoolR18SlTimeResourceR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeResource_r18 = &v
		}
	}

	// 8. sl-PosAllowedResourceSelectionConfig-r18: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlPosAllowedResourceSelectionConfigR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PosAllowedResourceSelectionConfig_r18 = &idx
		}
	}

	// 9. sl-PRS-ResourceReservePeriodList-r18: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlPRSResourceReservePeriodListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_ResourceReservePeriodList_r18 = make([]SL_ReservationPeriodAllowedDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_ResourceReservePeriodList_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. sl-PRS-ResourcesDedicatedSL-PRS-RP-r18: sequence-of
	{
		if seq.IsComponentPresent(8) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlPRSResourcesDedicatedSLPRSRPR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18 = make([]SL_PRS_ResourceDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_PRS_ResourcesDedicatedSL_PRS_RP_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. sl-PRS-PowerControl-r18: ref
	{
		if seq.IsComponentPresent(9) {
			ie.Sl_PRS_PowerControl_r18 = new(SL_PRS_PowerControl_r18)
			if err := ie.Sl_PRS_PowerControl_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. sl-SensingWindowDedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlSensingWindowDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SensingWindowDedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 13. sl-TxPercentageDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if seq.IsComponentPresent(11) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlTxPercentageDedicatedSLPRSRPListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18 = make([]SL_TxPercentageDedicatedSL_PRS_RP_Config_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_TxPercentageDedicatedSL_PRS_RP_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. sl-SCI-basedSL-PRS-TxTriggerSCI1-B-r18: boolean
	{
		if seq.IsComponentPresent(12) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.Sl_SCI_BasedSL_PRS_TxTriggerSCI1_B_r18 = &v
		}
	}

	// 15. sl-NumSubchannelDedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(13) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlNumSubchannelDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumSubchannelDedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 16. sl-SubchannelSizeDedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlSubchannelSizeDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SubchannelSizeDedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 17. sl-MaxNumPerReserveDedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlMaxNumPerReserveDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_MaxNumPerReserveDedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 18. sl-NumReservedBitsSCI1B-DedicatedSL-PRS-RP-r18: integer
	{
		if seq.IsComponentPresent(16) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlNumReservedBitsSCI1BDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumReservedBitsSCI1B_DedicatedSL_PRS_RP_r18 = &v
		}
	}

	// 19. sl-SRC-ID-LenDedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlSRCIDLenDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_SRC_ID_LenDedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 20. sl-CBR-PriorityTxConfigDedicatedSL-PRS-RP-List-r18: sequence-of
	{
		if seq.IsComponentPresent(18) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlCBRPriorityTxConfigDedicatedSLPRSRPListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18 = make([]SL_PriorityTxConfigIndexDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_CBR_PriorityTxConfigDedicatedSL_PRS_RP_List_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 21. sl-TimeWindowSizeCBR-DedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlTimeWindowSizeCBRDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeWindowSizeCBR_DedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 22. sl-TimeWindowSizeCR-DedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlTimeWindowSizeCRDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_TimeWindowSizeCR_DedicatedSL_PRS_RP_r18 = &idx
		}
	}

	// 23. sl-CBR-CommonTxDedicatedSL-PRS-RP-List-r18: ref
	{
		if seq.IsComponentPresent(21) {
			ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18 = new(SL_CBR_CommonTxDedicatedSL_PRS_RP_List_r18)
			if err := ie.Sl_CBR_CommonTxDedicatedSL_PRS_RP_List_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 24. sl-PriorityThreshold-UL-URLLC-r18: integer
	{
		if seq.IsComponentPresent(22) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlPriorityThresholdULURLLCR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThreshold_UL_URLLC_r18 = &v
		}
	}

	// 25. sl-PriorityThreshold-r18: integer
	{
		if seq.IsComponentPresent(23) {
			v, err := d.DecodeInteger(sLPRSResourcePoolR18SlPriorityThresholdR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PriorityThreshold_r18 = &v
		}
	}

	// 26. sl-SelectionWindowListDedicatedSL-PRS-RP-r18: sequence-of
	{
		if seq.IsComponentPresent(24) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlSelectionWindowListDedicatedSLPRSRPR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18 = make([]SL_SelectionWindowConfigDedicatedSL_PRS_RP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_SelectionWindowListDedicatedSL_PRS_RP_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 27. sl-Thres-RSRP-ListDedicatedSL-PRS-RP-r18: sequence-of
	{
		if seq.IsComponentPresent(25) {
			seqOf := d.NewSequenceOfDecoder(sLPRSResourcePoolR18SlThresRSRPListDedicatedSLPRSRPR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18 = make([]SL_PRS_ThresRSRP_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sl_Thres_RSRP_ListDedicatedSL_PRS_RP_r18[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 28. sl-PreemptionEnableDedicatedSL-PRS-RP-r18: enumerated
	{
		if seq.IsComponentPresent(26) {
			idx, err := d.DecodeEnumerated(sLPRSResourcePoolR18SlPreemptionEnableDedicatedSLPRSRPR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PreemptionEnableDedicatedSL_PRS_RP_r18 = &idx
		}
	}

	return nil
}
