// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplink-v1720 (line 20282).

var featureSetUplinkV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pucch-Repetition-F0-1-2-3-4-RRC-Config-r17", Optional: true},
		{Name: "pucch-Repetition-F0-1-2-3-4-DynamicIndication-r17", Optional: true},
		{Name: "interSubslotFreqHopping-PUCCH-r17", Optional: true},
		{Name: "semiStaticHARQ-ACK-CodebookSub-SlotPUCCH-r17", Optional: true},
		{Name: "phy-PrioritizationLowPriorityDG-HighPriorityCG-r17", Optional: true},
		{Name: "phy-PrioritizationHighPriorityDG-LowPriorityCG-r17", Optional: true},
		{Name: "extendedDC-LocationReport-r17", Optional: true},
	},
}

const (
	FeatureSetUplink_v1720_Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17_Supported = 0
)

var featureSetUplinkV1720PucchRepetitionF01234RRCConfigR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17_Supported},
}

const (
	FeatureSetUplink_v1720_Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17_Supported = 0
)

var featureSetUplinkV1720PucchRepetitionF01234DynamicIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17_Supported},
}

const (
	FeatureSetUplink_v1720_InterSubslotFreqHopping_PUCCH_r17_Supported = 0
)

var featureSetUplinkV1720InterSubslotFreqHoppingPUCCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_InterSubslotFreqHopping_PUCCH_r17_Supported},
}

const (
	FeatureSetUplink_v1720_SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17_Supported = 0
)

var featureSetUplinkV1720SemiStaticHARQACKCodebookSubSlotPUCCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17_Supported},
}

var featureSetUplinkV1720PhyPrioritizationLowPriorityDGHighPriorityCGR17Constraints = per.Constrained(1, 16)

const (
	FeatureSetUplink_v1720_ExtendedDC_LocationReport_r17_Supported = 0
)

var featureSetUplinkV1720ExtendedDCLocationReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_ExtendedDC_LocationReport_r17_Supported},
}

const (
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym0 = 0
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym1 = 1
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym2 = 2
)

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17PuschPreparationLowPriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym0, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym1, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_Pusch_PreparationLowPriority_r17_Sym2},
}

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

const (
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym0 = 0
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym1 = 1
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym2 = 2
)

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs15kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym0, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym1, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_15kHz_r17_Sym2},
}

const (
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym0 = 0
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym1 = 1
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym2 = 2
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym3 = 3
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym4 = 4
)

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs30kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym0, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym1, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym2, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym3, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_30kHz_r17_Sym4},
}

const (
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym0 = 0
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym1 = 1
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym2 = 2
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym3 = 3
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym4 = 4
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym5 = 5
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym6 = 6
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym7 = 7
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym8 = 8
)

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs60kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym0, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym1, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym2, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym3, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym4, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym5, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym6, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym7, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_60kHz_r17_Sym8},
}

const (
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym0  = 0
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym1  = 1
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym2  = 2
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym3  = 3
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym4  = 4
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym5  = 5
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym6  = 6
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym7  = 7
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym8  = 8
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym9  = 9
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym10 = 10
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym11 = 11
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym12 = 12
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym13 = 13
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym14 = 14
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym15 = 15
	FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym16 = 16
)

var featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym0, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym1, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym2, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym3, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym4, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym5, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym6, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym7, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym8, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym9, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym10, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym11, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym12, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym13, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym14, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym15, FeatureSetUplink_v1720_Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17_AdditionalCancellationTime_r17_Scs_120kHz_r17_Sym16},
}

type FeatureSetUplink_v1720 struct {
	Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17         *int64
	Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17  *int64
	InterSubslotFreqHopping_PUCCH_r17                  *int64
	SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17       *int64
	Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 *int64
	Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 *struct {
		Pusch_PreparationLowPriority_r17 int64
		AdditionalCancellationTime_r17   struct {
			Scs_15kHz_r17  *int64
			Scs_30kHz_r17  *int64
			Scs_60kHz_r17  *int64
			Scs_120kHz_r17 *int64
		}
		MaxNumberCarriers_r17 int64
	}
	ExtendedDC_LocationReport_r17 *int64
}

func (ie *FeatureSetUplink_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil, ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil, ie.InterSubslotFreqHopping_PUCCH_r17 != nil, ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil, ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil, ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil, ie.ExtendedDC_LocationReport_r17 != nil}); err != nil {
		return err
	}

	// 2. pucch-Repetition-F0-1-2-3-4-RRC-Config-r17: enumerated
	{
		if ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17, featureSetUplinkV1720PucchRepetitionF01234RRCConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. pucch-Repetition-F0-1-2-3-4-DynamicIndication-r17: enumerated
	{
		if ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17, featureSetUplinkV1720PucchRepetitionF01234DynamicIndicationR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. interSubslotFreqHopping-PUCCH-r17: enumerated
	{
		if ie.InterSubslotFreqHopping_PUCCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.InterSubslotFreqHopping_PUCCH_r17, featureSetUplinkV1720InterSubslotFreqHoppingPUCCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. semiStaticHARQ-ACK-CodebookSub-SlotPUCCH-r17: enumerated
	{
		if ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 != nil {
			if err := e.EncodeEnumerated(*ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17, featureSetUplinkV1720SemiStaticHARQACKCodebookSubSlotPUCCHR17Constraints); err != nil {
				return err
			}
		}
	}

	// 6. phy-PrioritizationLowPriorityDG-HighPriorityCG-r17: integer
	{
		if ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 != nil {
			if err := e.EncodeInteger(*ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17, featureSetUplinkV1720PhyPrioritizationLowPriorityDGHighPriorityCGR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. phy-PrioritizationHighPriorityDG-LowPriorityCG-r17: sequence
	{
		if ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 != nil {
			c := ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17
			if err := e.EncodeEnumerated(c.Pusch_PreparationLowPriority_r17, featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17PuschPreparationLowPriorityR17Constraints); err != nil {
				return err
			}
			{
				c := &c.AdditionalCancellationTime_r17
				featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq := e.NewSequenceEncoder(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Constraints)
				if err := featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_15kHz_r17), featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs15kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_30kHz_r17), featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs30kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_60kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz_r17), featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs60kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r17 != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz_r17), featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs120kHzR17Constraints); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxNumberCarriers_r17, per.Constrained(1, 16)); err != nil {
				return err
			}
		}
	}

	// 8. extendedDC-LocationReport-r17: enumerated
	{
		if ie.ExtendedDC_LocationReport_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ExtendedDC_LocationReport_r17, featureSetUplinkV1720ExtendedDCLocationReportR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplink_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. pucch-Repetition-F0-1-2-3-4-RRC-Config-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1720PucchRepetitionF01234RRCConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_Repetition_F0_1_2_3_4_RRC_Config_r17 = &idx
		}
	}

	// 3. pucch-Repetition-F0-1-2-3-4-DynamicIndication-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1720PucchRepetitionF01234DynamicIndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_Repetition_F0_1_2_3_4_DynamicIndication_r17 = &idx
		}
	}

	// 4. interSubslotFreqHopping-PUCCH-r17: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1720InterSubslotFreqHoppingPUCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.InterSubslotFreqHopping_PUCCH_r17 = &idx
		}
	}

	// 5. semiStaticHARQ-ACK-CodebookSub-SlotPUCCH-r17: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1720SemiStaticHARQACKCodebookSubSlotPUCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.SemiStaticHARQ_ACK_CodebookSub_SlotPUCCH_r17 = &idx
		}
	}

	// 6. phy-PrioritizationLowPriorityDG-HighPriorityCG-r17: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(featureSetUplinkV1720PhyPrioritizationLowPriorityDGHighPriorityCGR17Constraints)
			if err != nil {
				return err
			}
			ie.Phy_PrioritizationLowPriorityDG_HighPriorityCG_r17 = &v
		}
	}

	// 7. phy-PrioritizationHighPriorityDG-LowPriorityCG-r17: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17 = &struct {
				Pusch_PreparationLowPriority_r17 int64
				AdditionalCancellationTime_r17   struct {
					Scs_15kHz_r17  *int64
					Scs_30kHz_r17  *int64
					Scs_60kHz_r17  *int64
					Scs_120kHz_r17 *int64
				}
				MaxNumberCarriers_r17 int64
			}{}
			c := ie.Phy_PrioritizationHighPriorityDG_LowPriorityCG_r17
			{
				v, err := d.DecodeEnumerated(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17PuschPreparationLowPriorityR17Constraints)
				if err != nil {
					return err
				}
				c.Pusch_PreparationLowPriority_r17 = v
			}
			{
				c := &c.AdditionalCancellationTime_r17
				featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq := d.NewSequenceDecoder(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Constraints)
				if err := featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs15kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_r17) = v
				}
				if featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs30kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_r17) = v
				}
				if featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.IsComponentPresent(2) {
					c.Scs_60kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs60kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r17) = v
				}
				if featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Seq.IsComponentPresent(3) {
					c.Scs_120kHz_r17 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkV1720PhyPrioritizationHighPriorityDGLowPriorityCGR17AdditionalCancellationTimeR17Scs120kHzR17Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_r17) = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxNumberCarriers_r17 = v
			}
		}
	}

	// 8. extendedDC-LocationReport-r17: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetUplinkV1720ExtendedDCLocationReportR17Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedDC_LocationReport_r17 = &idx
		}
	}

	return nil
}
