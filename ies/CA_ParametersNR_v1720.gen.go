// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CA-ParametersNR-v1720 (line 17484).

var cAParametersNRV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "parallelTxSRS-PUCCH-PUSCH-intraBand-r17", Optional: true},
		{Name: "parallelTxPRACH-SRS-PUCCH-PUSCH-intraBand-r17", Optional: true},
		{Name: "semiStaticPUCCH-CellSwitchSingleGroup-r17", Optional: true},
		{Name: "semiStaticPUCCH-CellSwitchTwoGroups-r17", Optional: true},
		{Name: "dynamicPUCCH-CellSwitchSameLengthSingleGroup-r17", Optional: true},
		{Name: "dynamicPUCCH-CellSwitchDiffLengthSingleGroup-r17", Optional: true},
		{Name: "dynamicPUCCH-CellSwitchSameLengthTwoGroups-r17", Optional: true},
		{Name: "dynamicPUCCH-CellSwitchDiffLengthTwoGroups-r17", Optional: true},
		{Name: "ack-NACK-FeedbackForMulticast-r17", Optional: true},
		{Name: "ptp-Retx-Multicast-r17", Optional: true},
		{Name: "nack-OnlyFeedbackForMulticast-r17", Optional: true},
		{Name: "nack-OnlyFeedbackSpecificResourceForMulticast-r17", Optional: true},
		{Name: "ack-NACK-FeedbackForSPS-Multicast-r17", Optional: true},
		{Name: "ptp-Retx-SPS-Multicast-r17", Optional: true},
		{Name: "higherPowerLimit-r17", Optional: true},
		{Name: "parallelTxMsgA-SRS-PUCCH-PUSCH-intraBand-r17", Optional: true},
		{Name: "pdcch-MonitoringCA-r17", Optional: true},
		{Name: "pdcch-BlindDetectionMCG-SCG-List-r17", Optional: true},
		{Name: "pdcch-BlindDetectionMixedList1-r17", Optional: true},
		{Name: "pdcch-BlindDetectionMixedList2-r17", Optional: true},
		{Name: "pdcch-BlindDetectionMixedList3-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1720_ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17_Supported = 0
)

var cAParametersNRV1720ParallelTxSRSPUCCHPUSCHIntraBandR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17_Supported},
}

const (
	CA_ParametersNR_v1720_ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17_Supported = 0
)

var cAParametersNRV1720ParallelTxPRACHSRSPUCCHPUSCHIntraBandR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17_Supported},
}

var cAParametersNRV1720SemiStaticPUCCHCellSwitchTwoGroupsR17Constraints = per.SizeRange(1, common.MaxTwoPUCCH_Grp_ConfigList_r17)

var cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthTwoGroupsR17Constraints = per.SizeRange(1, common.MaxTwoPUCCH_Grp_ConfigList_r17)

var cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthTwoGroupsR17Constraints = per.SizeRange(1, common.MaxTwoPUCCH_Grp_ConfigList_r17)

const (
	CA_ParametersNR_v1720_Ack_NACK_FeedbackForMulticast_r17_Supported = 0
)

var cAParametersNRV1720AckNACKFeedbackForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Ack_NACK_FeedbackForMulticast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_Ptp_Retx_Multicast_r17_Supported = 0
)

var cAParametersNRV1720PtpRetxMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Ptp_Retx_Multicast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_Nack_OnlyFeedbackForMulticast_r17_Supported = 0
)

var cAParametersNRV1720NackOnlyFeedbackForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Nack_OnlyFeedbackForMulticast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_Nack_OnlyFeedbackSpecificResourceForMulticast_r17_Supported = 0
)

var cAParametersNRV1720NackOnlyFeedbackSpecificResourceForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Nack_OnlyFeedbackSpecificResourceForMulticast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_Ack_NACK_FeedbackForSPS_Multicast_r17_Supported = 0
)

var cAParametersNRV1720AckNACKFeedbackForSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Ack_NACK_FeedbackForSPS_Multicast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_Ptp_Retx_SPS_Multicast_r17_Supported = 0
)

var cAParametersNRV1720PtpRetxSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_Ptp_Retx_SPS_Multicast_r17_Supported},
}

const (
	CA_ParametersNR_v1720_HigherPowerLimit_r17_Supported = 0
)

var cAParametersNRV1720HigherPowerLimitR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_HigherPowerLimit_r17_Supported},
}

const (
	CA_ParametersNR_v1720_ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17_Supported = 0
)

var cAParametersNRV1720ParallelTxMsgASRSPUCCHPUSCHIntraBandR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17_Supported},
}

var cAParametersNRV1720PdcchMonitoringCAR17Constraints = per.Constrained(4, 16)

var cAParametersNRV1720PdcchBlindDetectionMCGSCGListR17Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetection_r17)

var cAParametersNRV1720PdcchBlindDetectionMixedList1R17Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetection_r17)

var cAParametersNRV1720PdcchBlindDetectionMixedList2R17Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetection_r17)

var cAParametersNRV1720PdcchBlindDetectionMixedList3R17Constraints = per.SizeRange(1, common.MaxNrofPdcch_BlindDetection_r17)

const (
	CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly              = 0
	CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly            = 1
	CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup = 2
)

var cAParametersNRV1720SemiStaticPUCCHCellSwitchSingleGroupR17PucchGroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly, CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly, CA_ParametersNR_v1720_SemiStaticPUCCH_CellSwitchSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup},
}

const (
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly              = 0
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly            = 1
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup = 2
)

var cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthSingleGroupR17PucchGroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly, CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly, CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup},
}

const (
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly              = 0
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly            = 1
	CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup = 2
)

var cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthSingleGroupR17PucchGroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_PrimaryGroupOnly, CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_SecondaryGroupOnly, CA_ParametersNR_v1720_DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17_Pucch_Group_r17_EitherPrimaryOrSecondaryGroup},
}

type CA_ParametersNR_v1720 struct {
	ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17       *int64
	ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17 *int64
	SemiStaticPUCCH_CellSwitchSingleGroup_r17     *struct {
		Pucch_Group_r17        int64
		Pucch_Group_Config_r17 PUCCH_Group_Config_r17
	}
	SemiStaticPUCCH_CellSwitchTwoGroups_r17          []TwoPUCCH_Grp_Configurations_r17
	DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 *struct {
		Pucch_Group_r17        int64
		Pucch_Group_Config_r17 PUCCH_Group_Config_r17
	}
	DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 *struct {
		Pucch_Group_r17        int64
		Pucch_Group_Config_r17 PUCCH_Group_Config_r17
	}
	DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17
	DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17    []TwoPUCCH_Grp_Configurations_r17
	Ack_NACK_FeedbackForMulticast_r17                 *int64
	Ptp_Retx_Multicast_r17                            *int64
	Nack_OnlyFeedbackForMulticast_r17                 *int64
	Nack_OnlyFeedbackSpecificResourceForMulticast_r17 *int64
	Ack_NACK_FeedbackForSPS_Multicast_r17             *int64
	Ptp_Retx_SPS_Multicast_r17                        *int64
	HigherPowerLimit_r17                              *int64
	ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17      *int64
	Pdcch_MonitoringCA_r17                            *int64
	Pdcch_BlindDetectionMCG_SCG_List_r17              []PDCCH_BlindDetectionMCG_SCG_r17
	Pdcch_BlindDetectionMixedList1_r17                []PDCCH_BlindDetectionMixed_r17
	Pdcch_BlindDetectionMixedList2_r17                []PDCCH_BlindDetectionMixed_r17
	Pdcch_BlindDetectionMixedList3_r17                []PDCCH_BlindDetectionMixed1_r17
}

func (ie *CA_ParametersNR_v1720) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1720Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17 != nil, ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17 != nil, ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 != nil, ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 != nil, ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil, ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil, ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 != nil, ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 != nil, ie.Ack_NACK_FeedbackForMulticast_r17 != nil, ie.Ptp_Retx_Multicast_r17 != nil, ie.Nack_OnlyFeedbackForMulticast_r17 != nil, ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil, ie.Ack_NACK_FeedbackForSPS_Multicast_r17 != nil, ie.Ptp_Retx_SPS_Multicast_r17 != nil, ie.HigherPowerLimit_r17 != nil, ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17 != nil, ie.Pdcch_MonitoringCA_r17 != nil, ie.Pdcch_BlindDetectionMCG_SCG_List_r17 != nil, ie.Pdcch_BlindDetectionMixedList1_r17 != nil, ie.Pdcch_BlindDetectionMixedList2_r17 != nil, ie.Pdcch_BlindDetectionMixedList3_r17 != nil}); err != nil {
		return err
	}

	// 2. parallelTxSRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if ie.ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17, cAParametersNRV1720ParallelTxSRSPUCCHPUSCHIntraBandR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. parallelTxPRACH-SRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17, cAParametersNRV1720ParallelTxPRACHSRSPUCCHPUSCHIntraBandR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. semiStaticPUCCH-CellSwitchSingleGroup-r17: sequence
	{
		if ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 != nil {
			c := ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17
			if err := e.EncodeEnumerated(c.Pucch_Group_r17, cAParametersNRV1720SemiStaticPUCCHCellSwitchSingleGroupR17PucchGroupR17Constraints); err != nil {
				return err
			}
			if err := c.Pucch_Group_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. semiStaticPUCCH-CellSwitchTwoGroups-r17: sequence-of
	{
		if ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720SemiStaticPUCCHCellSwitchTwoGroupsR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17))); err != nil {
				return err
			}
			for i := range ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 {
				if err := ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. dynamicPUCCH-CellSwitchSameLengthSingleGroup-r17: sequence
	{
		if ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 != nil {
			c := ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17
			if err := e.EncodeEnumerated(c.Pucch_Group_r17, cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthSingleGroupR17PucchGroupR17Constraints); err != nil {
				return err
			}
			if err := c.Pucch_Group_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. dynamicPUCCH-CellSwitchDiffLengthSingleGroup-r17: sequence
	{
		if ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 != nil {
			c := ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17
			if err := e.EncodeEnumerated(c.Pucch_Group_r17, cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthSingleGroupR17PucchGroupR17Constraints); err != nil {
				return err
			}
			if err := c.Pucch_Group_Config_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. dynamicPUCCH-CellSwitchSameLengthTwoGroups-r17: sequence-of
	{
		if ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthTwoGroupsR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17))); err != nil {
				return err
			}
			for i := range ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 {
				if err := ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. dynamicPUCCH-CellSwitchDiffLengthTwoGroups-r17: sequence-of
	{
		if ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthTwoGroupsR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17))); err != nil {
				return err
			}
			for i := range ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 {
				if err := ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. ack-NACK-FeedbackForMulticast-r17: enumerated
	{
		if ie.Ack_NACK_FeedbackForMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ack_NACK_FeedbackForMulticast_r17, cAParametersNRV1720AckNACKFeedbackForMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 11. ptp-Retx-Multicast-r17: enumerated
	{
		if ie.Ptp_Retx_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ptp_Retx_Multicast_r17, cAParametersNRV1720PtpRetxMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 12. nack-OnlyFeedbackForMulticast-r17: enumerated
	{
		if ie.Nack_OnlyFeedbackForMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Nack_OnlyFeedbackForMulticast_r17, cAParametersNRV1720NackOnlyFeedbackForMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 13. nack-OnlyFeedbackSpecificResourceForMulticast-r17: enumerated
	{
		if ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17, cAParametersNRV1720NackOnlyFeedbackSpecificResourceForMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 14. ack-NACK-FeedbackForSPS-Multicast-r17: enumerated
	{
		if ie.Ack_NACK_FeedbackForSPS_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ack_NACK_FeedbackForSPS_Multicast_r17, cAParametersNRV1720AckNACKFeedbackForSPSMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 15. ptp-Retx-SPS-Multicast-r17: enumerated
	{
		if ie.Ptp_Retx_SPS_Multicast_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ptp_Retx_SPS_Multicast_r17, cAParametersNRV1720PtpRetxSPSMulticastR17Constraints); err != nil {
				return err
			}
		}
	}

	// 16. higherPowerLimit-r17: enumerated
	{
		if ie.HigherPowerLimit_r17 != nil {
			if err := e.EncodeEnumerated(*ie.HigherPowerLimit_r17, cAParametersNRV1720HigherPowerLimitR17Constraints); err != nil {
				return err
			}
		}
	}

	// 17. parallelTxMsgA-SRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17, cAParametersNRV1720ParallelTxMsgASRSPUCCHPUSCHIntraBandR17Constraints); err != nil {
				return err
			}
		}
	}

	// 18. pdcch-MonitoringCA-r17: integer
	{
		if ie.Pdcch_MonitoringCA_r17 != nil {
			if err := e.EncodeInteger(*ie.Pdcch_MonitoringCA_r17, cAParametersNRV1720PdcchMonitoringCAR17Constraints); err != nil {
				return err
			}
		}
	}

	// 19. pdcch-BlindDetectionMCG-SCG-List-r17: sequence-of
	{
		if ie.Pdcch_BlindDetectionMCG_SCG_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720PdcchBlindDetectionMCGSCGListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionMCG_SCG_List_r17))); err != nil {
				return err
			}
			for i := range ie.Pdcch_BlindDetectionMCG_SCG_List_r17 {
				if err := ie.Pdcch_BlindDetectionMCG_SCG_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 20. pdcch-BlindDetectionMixedList1-r17: sequence-of
	{
		if ie.Pdcch_BlindDetectionMixedList1_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720PdcchBlindDetectionMixedList1R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionMixedList1_r17))); err != nil {
				return err
			}
			for i := range ie.Pdcch_BlindDetectionMixedList1_r17 {
				if err := ie.Pdcch_BlindDetectionMixedList1_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 21. pdcch-BlindDetectionMixedList2-r17: sequence-of
	{
		if ie.Pdcch_BlindDetectionMixedList2_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720PdcchBlindDetectionMixedList2R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionMixedList2_r17))); err != nil {
				return err
			}
			for i := range ie.Pdcch_BlindDetectionMixedList2_r17 {
				if err := ie.Pdcch_BlindDetectionMixedList2_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 22. pdcch-BlindDetectionMixedList3-r17: sequence-of
	{
		if ie.Pdcch_BlindDetectionMixedList3_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cAParametersNRV1720PdcchBlindDetectionMixedList3R17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Pdcch_BlindDetectionMixedList3_r17))); err != nil {
				return err
			}
			for i := range ie.Pdcch_BlindDetectionMixedList3_r17 {
				if err := ie.Pdcch_BlindDetectionMixedList3_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1720) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1720Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. parallelTxSRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720ParallelTxSRSPUCCHPUSCHIntraBandR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxSRS_PUCCH_PUSCH_IntraBand_r17 = &idx
		}
	}

	// 3. parallelTxPRACH-SRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720ParallelTxPRACHSRSPUCCHPUSCHIntraBandR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxPRACH_SRS_PUCCH_PUSCH_IntraBand_r17 = &idx
		}
	}

	// 4. semiStaticPUCCH-CellSwitchSingleGroup-r17: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17 = &struct {
				Pucch_Group_r17        int64
				Pucch_Group_Config_r17 PUCCH_Group_Config_r17
			}{}
			c := ie.SemiStaticPUCCH_CellSwitchSingleGroup_r17
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1720SemiStaticPUCCHCellSwitchSingleGroupR17PucchGroupR17Constraints)
				if err != nil {
					return err
				}
				c.Pucch_Group_r17 = v
			}
			{
				if err := c.Pucch_Group_Config_r17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. semiStaticPUCCH-CellSwitchTwoGroups-r17: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720SemiStaticPUCCHCellSwitchTwoGroupsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17 = make([]TwoPUCCH_Grp_Configurations_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SemiStaticPUCCH_CellSwitchTwoGroups_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. dynamicPUCCH-CellSwitchSameLengthSingleGroup-r17: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17 = &struct {
				Pucch_Group_r17        int64
				Pucch_Group_Config_r17 PUCCH_Group_Config_r17
			}{}
			c := ie.DynamicPUCCH_CellSwitchSameLengthSingleGroup_r17
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthSingleGroupR17PucchGroupR17Constraints)
				if err != nil {
					return err
				}
				c.Pucch_Group_r17 = v
			}
			{
				if err := c.Pucch_Group_Config_r17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. dynamicPUCCH-CellSwitchDiffLengthSingleGroup-r17: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17 = &struct {
				Pucch_Group_r17        int64
				Pucch_Group_Config_r17 PUCCH_Group_Config_r17
			}{}
			c := ie.DynamicPUCCH_CellSwitchDiffLengthSingleGroup_r17
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthSingleGroupR17PucchGroupR17Constraints)
				if err != nil {
					return err
				}
				c.Pucch_Group_r17 = v
			}
			{
				if err := c.Pucch_Group_Config_r17.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. dynamicPUCCH-CellSwitchSameLengthTwoGroups-r17: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720DynamicPUCCHCellSwitchSameLengthTwoGroupsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17 = make([]TwoPUCCH_Grp_Configurations_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DynamicPUCCH_CellSwitchSameLengthTwoGroups_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. dynamicPUCCH-CellSwitchDiffLengthTwoGroups-r17: sequence-of
	{
		if seq.IsComponentPresent(7) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720DynamicPUCCHCellSwitchDiffLengthTwoGroupsR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17 = make([]TwoPUCCH_Grp_Configurations_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DynamicPUCCH_CellSwitchDiffLengthTwoGroups_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. ack-NACK-FeedbackForMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720AckNACKFeedbackForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Ack_NACK_FeedbackForMulticast_r17 = &idx
		}
	}

	// 11. ptp-Retx-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720PtpRetxMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Ptp_Retx_Multicast_r17 = &idx
		}
	}

	// 12. nack-OnlyFeedbackForMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720NackOnlyFeedbackForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackForMulticast_r17 = &idx
		}
	}

	// 13. nack-OnlyFeedbackSpecificResourceForMulticast-r17: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720NackOnlyFeedbackSpecificResourceForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackSpecificResourceForMulticast_r17 = &idx
		}
	}

	// 14. ack-NACK-FeedbackForSPS-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720AckNACKFeedbackForSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Ack_NACK_FeedbackForSPS_Multicast_r17 = &idx
		}
	}

	// 15. ptp-Retx-SPS-Multicast-r17: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720PtpRetxSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Ptp_Retx_SPS_Multicast_r17 = &idx
		}
	}

	// 16. higherPowerLimit-r17: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720HigherPowerLimitR17Constraints)
			if err != nil {
				return err
			}
			ie.HigherPowerLimit_r17 = &idx
		}
	}

	// 17. parallelTxMsgA-SRS-PUCCH-PUSCH-intraBand-r17: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1720ParallelTxMsgASRSPUCCHPUSCHIntraBandR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelTxMsgA_SRS_PUCCH_PUSCH_IntraBand_r17 = &idx
		}
	}

	// 18. pdcch-MonitoringCA-r17: integer
	{
		if seq.IsComponentPresent(16) {
			v, err := d.DecodeInteger(cAParametersNRV1720PdcchMonitoringCAR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringCA_r17 = &v
		}
	}

	// 19. pdcch-BlindDetectionMCG-SCG-List-r17: sequence-of
	{
		if seq.IsComponentPresent(17) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720PdcchBlindDetectionMCGSCGListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionMCG_SCG_List_r17 = make([]PDCCH_BlindDetectionMCG_SCG_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_BlindDetectionMCG_SCG_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 20. pdcch-BlindDetectionMixedList1-r17: sequence-of
	{
		if seq.IsComponentPresent(18) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720PdcchBlindDetectionMixedList1R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionMixedList1_r17 = make([]PDCCH_BlindDetectionMixed_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_BlindDetectionMixedList1_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 21. pdcch-BlindDetectionMixedList2-r17: sequence-of
	{
		if seq.IsComponentPresent(19) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720PdcchBlindDetectionMixedList2R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionMixedList2_r17 = make([]PDCCH_BlindDetectionMixed_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_BlindDetectionMixedList2_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 22. pdcch-BlindDetectionMixedList3-r17: sequence-of
	{
		if seq.IsComponentPresent(20) {
			seqOf := d.NewSequenceOfDecoder(cAParametersNRV1720PdcchBlindDetectionMixedList3R17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionMixedList3_r17 = make([]PDCCH_BlindDetectionMixed1_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdcch_BlindDetectionMixedList3_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
