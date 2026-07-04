// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MAC-ParametersCommon (line 20965).

var mACParametersCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "lcp-Restriction", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "lch-ToSCellRestriction", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
		{Name: "extension-addition-7", Optional: true},
	},
}

const (
	MAC_ParametersCommon_Lcp_Restriction_Supported = 0
)

var mACParametersCommonLcpRestrictionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Lcp_Restriction_Supported},
}

const (
	MAC_ParametersCommon_Dummy_Supported = 0
)

var mACParametersCommonDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Dummy_Supported},
}

const (
	MAC_ParametersCommon_Lch_ToSCellRestriction_Supported = 0
)

var mACParametersCommonLchToSCellRestrictionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Lch_ToSCellRestriction_Supported},
}

const (
	MAC_ParametersCommon_Ext_RecommendedBitRate_Supported = 0
)

var mACParametersCommonExtRecommendedBitRateConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_RecommendedBitRate_Supported},
}

const (
	MAC_ParametersCommon_Ext_RecommendedBitRateQuery_Supported = 0
)

var mACParametersCommonExtRecommendedBitRateQueryConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_RecommendedBitRateQuery_Supported},
}

const (
	MAC_ParametersCommon_Ext_RecommendedBitRateMultiplier_r16_Supported = 0
)

var mACParametersCommonExtRecommendedBitRateMultiplierR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_RecommendedBitRateMultiplier_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_PreEmptiveBSR_r16_Supported = 0
)

var mACParametersCommonExtPreEmptiveBSRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_PreEmptiveBSR_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_AutonomousTransmission_r16_Supported = 0
)

var mACParametersCommonExtAutonomousTransmissionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_AutonomousTransmission_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lch_PriorityBasedPrioritization_r16_Supported = 0
)

var mACParametersCommonExtLchPriorityBasedPrioritizationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lch_PriorityBasedPrioritization_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lch_ToConfiguredGrantMapping_r16_Supported = 0
)

var mACParametersCommonExtLchToConfiguredGrantMappingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lch_ToConfiguredGrantMapping_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lch_ToGrantPriorityRestriction_r16_Supported = 0
)

var mACParametersCommonExtLchToGrantPriorityRestrictionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lch_ToGrantPriorityRestriction_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_SinglePHR_P_r16_Supported = 0
)

var mACParametersCommonExtSinglePHRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_SinglePHR_P_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Ul_LBT_FailureDetectionRecovery_r16_Supported = 0
)

var mACParametersCommonExtUlLBTFailureDetectionRecoveryR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Ul_LBT_FailureDetectionRecovery_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Tdd_MPE_P_MPR_Reporting_r16_Supported = 0
)

var mACParametersCommonExtTddMPEPMPRReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Tdd_MPE_P_MPR_Reporting_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lcid_ExtensionIAB_r16_Supported = 0
)

var mACParametersCommonExtLcidExtensionIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lcid_ExtensionIAB_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_SpCell_BFR_CBRA_r16_Supported = 0
)

var mACParametersCommonExtSpCellBFRCBRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_SpCell_BFR_CBRA_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_Srs_ResourceId_Ext_r16_Supported = 0
)

var mACParametersCommonExtSrsResourceIdExtR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Srs_ResourceId_Ext_r16_Supported},
}

const (
	MAC_ParametersCommon_Ext_EnhancedUuDRX_ForSidelink_r17_Supported = 0
)

var mACParametersCommonExtEnhancedUuDRXForSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_EnhancedUuDRX_ForSidelink_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Mg_ActivationRequestPRS_Meas_r17_Supported = 0
)

var mACParametersCommonExtMgActivationRequestPRSMeasR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Mg_ActivationRequestPRS_Meas_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Mg_ActivationCommPRS_Meas_r17_Supported = 0
)

var mACParametersCommonExtMgActivationCommPRSMeasR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Mg_ActivationCommPRS_Meas_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_IntraCG_Prioritization_r17_Supported = 0
)

var mACParametersCommonExtIntraCGPrioritizationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_IntraCG_Prioritization_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_JointPrioritizationCG_Retx_Timer_r17_Supported = 0
)

var mACParametersCommonExtJointPrioritizationCGRetxTimerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_JointPrioritizationCG_Retx_Timer_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_SurvivalTime_r17_Supported = 0
)

var mACParametersCommonExtSurvivalTimeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_SurvivalTime_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lcg_ExtensionIAB_r17_Supported = 0
)

var mACParametersCommonExtLcgExtensionIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lcg_ExtensionIAB_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Harq_FeedbackDisabled_r17_Supported = 0
)

var mACParametersCommonExtHarqFeedbackDisabledR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Harq_FeedbackDisabled_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Uplink_Harq_ModeB_r17_Supported = 0
)

var mACParametersCommonExtUplinkHarqModeBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Uplink_Harq_ModeB_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Sr_TriggeredBy_TA_Report_r17_Supported = 0
)

var mACParametersCommonExtSrTriggeredByTAReportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Sr_TriggeredBy_TA_Report_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_ExtendedDRX_CycleInactive_r17_Supported = 0
)

var mACParametersCommonExtExtendedDRXCycleInactiveR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_ExtendedDRX_CycleInactive_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17_Supported = 0
)

var mACParametersCommonExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_LastTransmissionUL_r17_Supported = 0
)

var mACParametersCommonExtLastTransmissionULR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_LastTransmissionUL_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17_Supported = 0
)

var mACParametersCommonExtHarqRTTTimerDLForNTNMulticastMBSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17_Supported},
}

const (
	MAC_ParametersCommon_Ext_Sr_TriggeredByTA_ReportATG_r18_Supported = 0
)

var mACParametersCommonExtSrTriggeredByTAReportATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Sr_TriggeredByTA_ReportATG_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_ExtendedDRX_CycleInactive_r18_Supported = 0
)

var mACParametersCommonExtExtendedDRXCycleInactiveR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_ExtendedDRX_CycleInactive_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_AdditionalBS_Table_r18_Supported = 0
)

var mACParametersCommonExtAdditionalBSTableR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_AdditionalBS_Table_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_DelayStatusReport_r18_Supported = 0
)

var mACParametersCommonExtDelayStatusReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_DelayStatusReport_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_Cg_RetransmissionMonitoringDisabling_r18_Supported = 0
)

var mACParametersCommonExtCgRetransmissionMonitoringDisablingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Cg_RetransmissionMonitoringDisabling_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_Non_IntegerDRX_r18_Supported = 0
)

var mACParametersCommonExtNonIntegerDRXR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Non_IntegerDRX_r18_Supported},
}

const (
	MAC_ParametersCommon_Ext_MultipleEntryDelayStatusReport_r19_Supported = 0
)

var mACParametersCommonExtMultipleEntryDelayStatusReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_MultipleEntryDelayStatusReport_r19_Supported},
}

const (
	MAC_ParametersCommon_Ext_Lcp_PriorityAdjustment_r19_Supported = 0
)

var mACParametersCommonExtLcpPriorityAdjustmentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Lcp_PriorityAdjustment_r19_Supported},
}

const (
	MAC_ParametersCommon_Ext_Ul_RateControl_r19_Supported = 0
)

var mACParametersCommonExtUlRateControlR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Ul_RateControl_r19_Supported},
}

const (
	MAC_ParametersCommon_Ext_Ul_RateQuery_r19_Supported = 0
)

var mACParametersCommonExtUlRateQueryR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_Ul_RateQuery_r19_Supported},
}

const (
	MAC_ParametersCommon_Ext_DelayStatusReportNonDelayReportingData_r19_Supported = 0
)

var mACParametersCommonExtDelayStatusReportNonDelayReportingDataR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MAC_ParametersCommon_Ext_DelayStatusReportNonDelayReportingData_r19_Supported},
}

type MAC_ParametersCommon struct {
	Lcp_Restriction                            *int64
	Dummy                                      *int64
	Lch_ToSCellRestriction                     *int64
	RecommendedBitRate                         *int64
	RecommendedBitRateQuery                    *int64
	RecommendedBitRateMultiplier_r16           *int64
	PreEmptiveBSR_r16                          *int64
	AutonomousTransmission_r16                 *int64
	Lch_PriorityBasedPrioritization_r16        *int64
	Lch_ToConfiguredGrantMapping_r16           *int64
	Lch_ToGrantPriorityRestriction_r16         *int64
	SinglePHR_P_r16                            *int64
	Ul_LBT_FailureDetectionRecovery_r16        *int64
	Tdd_MPE_P_MPR_Reporting_r16                *int64
	Lcid_ExtensionIAB_r16                      *int64
	SpCell_BFR_CBRA_r16                        *int64
	Srs_ResourceId_Ext_r16                     *int64
	EnhancedUuDRX_ForSidelink_r17              *int64
	Mg_ActivationRequestPRS_Meas_r17           *int64
	Mg_ActivationCommPRS_Meas_r17              *int64
	IntraCG_Prioritization_r17                 *int64
	JointPrioritizationCG_Retx_Timer_r17       *int64
	SurvivalTime_r17                           *int64
	Lcg_ExtensionIAB_r17                       *int64
	Harq_FeedbackDisabled_r17                  *int64
	Uplink_Harq_ModeB_r17                      *int64
	Sr_TriggeredBy_TA_Report_r17               *int64
	ExtendedDRX_CycleInactive_r17              *int64
	SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17  *int64
	LastTransmissionUL_r17                     *int64
	Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17   *int64
	Sr_TriggeredByTA_ReportATG_r18             *int64
	ExtendedDRX_CycleInactive_r18              *int64
	AdditionalBS_Table_r18                     *int64
	DelayStatusReport_r18                      *int64
	Cg_RetransmissionMonitoringDisabling_r18   *int64
	Non_IntegerDRX_r18                         *int64
	MultipleEntryDelayStatusReport_r19         *int64
	Lcp_PriorityAdjustment_r19                 *int64
	Ul_RateControl_r19                         *int64
	Ul_RateQuery_r19                           *int64
	DelayStatusReportNonDelayReportingData_r19 *int64
}

func (ie *MAC_ParametersCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mACParametersCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.RecommendedBitRate != nil || ie.RecommendedBitRateQuery != nil
	hasExtGroup1 := ie.RecommendedBitRateMultiplier_r16 != nil || ie.PreEmptiveBSR_r16 != nil || ie.AutonomousTransmission_r16 != nil || ie.Lch_PriorityBasedPrioritization_r16 != nil || ie.Lch_ToConfiguredGrantMapping_r16 != nil || ie.Lch_ToGrantPriorityRestriction_r16 != nil || ie.SinglePHR_P_r16 != nil || ie.Ul_LBT_FailureDetectionRecovery_r16 != nil || ie.Tdd_MPE_P_MPR_Reporting_r16 != nil || ie.Lcid_ExtensionIAB_r16 != nil
	hasExtGroup2 := ie.SpCell_BFR_CBRA_r16 != nil
	hasExtGroup3 := ie.Srs_ResourceId_Ext_r16 != nil
	hasExtGroup4 := ie.EnhancedUuDRX_ForSidelink_r17 != nil || ie.Mg_ActivationRequestPRS_Meas_r17 != nil || ie.Mg_ActivationCommPRS_Meas_r17 != nil || ie.IntraCG_Prioritization_r17 != nil || ie.JointPrioritizationCG_Retx_Timer_r17 != nil || ie.SurvivalTime_r17 != nil || ie.Lcg_ExtensionIAB_r17 != nil || ie.Harq_FeedbackDisabled_r17 != nil || ie.Uplink_Harq_ModeB_r17 != nil || ie.Sr_TriggeredBy_TA_Report_r17 != nil || ie.ExtendedDRX_CycleInactive_r17 != nil || ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil || ie.LastTransmissionUL_r17 != nil
	hasExtGroup5 := ie.Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17 != nil
	hasExtGroup6 := ie.Sr_TriggeredByTA_ReportATG_r18 != nil || ie.ExtendedDRX_CycleInactive_r18 != nil || ie.AdditionalBS_Table_r18 != nil || ie.DelayStatusReport_r18 != nil || ie.Cg_RetransmissionMonitoringDisabling_r18 != nil || ie.Non_IntegerDRX_r18 != nil
	hasExtGroup7 := ie.MultipleEntryDelayStatusReport_r19 != nil || ie.Lcp_PriorityAdjustment_r19 != nil || ie.Ul_RateControl_r19 != nil || ie.Ul_RateQuery_r19 != nil || ie.DelayStatusReportNonDelayReportingData_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Lcp_Restriction != nil, ie.Dummy != nil, ie.Lch_ToSCellRestriction != nil}); err != nil {
		return err
	}

	// 3. lcp-Restriction: enumerated
	{
		if ie.Lcp_Restriction != nil {
			if err := e.EncodeEnumerated(*ie.Lcp_Restriction, mACParametersCommonLcpRestrictionConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, mACParametersCommonDummyConstraints); err != nil {
				return err
			}
		}
	}

	// 5. lch-ToSCellRestriction: enumerated
	{
		if ie.Lch_ToSCellRestriction != nil {
			if err := e.EncodeEnumerated(*ie.Lch_ToSCellRestriction, mACParametersCommonLchToSCellRestrictionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "recommendedBitRate", Optional: true},
					{Name: "recommendedBitRateQuery", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RecommendedBitRate != nil, ie.RecommendedBitRateQuery != nil}); err != nil {
				return err
			}

			if ie.RecommendedBitRate != nil {
				if err := ex.EncodeEnumerated(*ie.RecommendedBitRate, mACParametersCommonExtRecommendedBitRateConstraints); err != nil {
					return err
				}
			}

			if ie.RecommendedBitRateQuery != nil {
				if err := ex.EncodeEnumerated(*ie.RecommendedBitRateQuery, mACParametersCommonExtRecommendedBitRateQueryConstraints); err != nil {
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
					{Name: "recommendedBitRateMultiplier-r16", Optional: true},
					{Name: "preEmptiveBSR-r16", Optional: true},
					{Name: "autonomousTransmission-r16", Optional: true},
					{Name: "lch-PriorityBasedPrioritization-r16", Optional: true},
					{Name: "lch-ToConfiguredGrantMapping-r16", Optional: true},
					{Name: "lch-ToGrantPriorityRestriction-r16", Optional: true},
					{Name: "singlePHR-P-r16", Optional: true},
					{Name: "ul-LBT-FailureDetectionRecovery-r16", Optional: true},
					{Name: "tdd-MPE-P-MPR-Reporting-r16", Optional: true},
					{Name: "lcid-ExtensionIAB-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RecommendedBitRateMultiplier_r16 != nil, ie.PreEmptiveBSR_r16 != nil, ie.AutonomousTransmission_r16 != nil, ie.Lch_PriorityBasedPrioritization_r16 != nil, ie.Lch_ToConfiguredGrantMapping_r16 != nil, ie.Lch_ToGrantPriorityRestriction_r16 != nil, ie.SinglePHR_P_r16 != nil, ie.Ul_LBT_FailureDetectionRecovery_r16 != nil, ie.Tdd_MPE_P_MPR_Reporting_r16 != nil, ie.Lcid_ExtensionIAB_r16 != nil}); err != nil {
				return err
			}

			if ie.RecommendedBitRateMultiplier_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.RecommendedBitRateMultiplier_r16, mACParametersCommonExtRecommendedBitRateMultiplierR16Constraints); err != nil {
					return err
				}
			}

			if ie.PreEmptiveBSR_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PreEmptiveBSR_r16, mACParametersCommonExtPreEmptiveBSRR16Constraints); err != nil {
					return err
				}
			}

			if ie.AutonomousTransmission_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AutonomousTransmission_r16, mACParametersCommonExtAutonomousTransmissionR16Constraints); err != nil {
					return err
				}
			}

			if ie.Lch_PriorityBasedPrioritization_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Lch_PriorityBasedPrioritization_r16, mACParametersCommonExtLchPriorityBasedPrioritizationR16Constraints); err != nil {
					return err
				}
			}

			if ie.Lch_ToConfiguredGrantMapping_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Lch_ToConfiguredGrantMapping_r16, mACParametersCommonExtLchToConfiguredGrantMappingR16Constraints); err != nil {
					return err
				}
			}

			if ie.Lch_ToGrantPriorityRestriction_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Lch_ToGrantPriorityRestriction_r16, mACParametersCommonExtLchToGrantPriorityRestrictionR16Constraints); err != nil {
					return err
				}
			}

			if ie.SinglePHR_P_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SinglePHR_P_r16, mACParametersCommonExtSinglePHRPR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_LBT_FailureDetectionRecovery_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_LBT_FailureDetectionRecovery_r16, mACParametersCommonExtUlLBTFailureDetectionRecoveryR16Constraints); err != nil {
					return err
				}
			}

			if ie.Tdd_MPE_P_MPR_Reporting_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Tdd_MPE_P_MPR_Reporting_r16, mACParametersCommonExtTddMPEPMPRReportingR16Constraints); err != nil {
					return err
				}
			}

			if ie.Lcid_ExtensionIAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Lcid_ExtensionIAB_r16, mACParametersCommonExtLcidExtensionIABR16Constraints); err != nil {
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
					{Name: "spCell-BFR-CBRA-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpCell_BFR_CBRA_r16 != nil}); err != nil {
				return err
			}

			if ie.SpCell_BFR_CBRA_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SpCell_BFR_CBRA_r16, mACParametersCommonExtSpCellBFRCBRAR16Constraints); err != nil {
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
					{Name: "srs-ResourceId-Ext-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_ResourceId_Ext_r16 != nil}); err != nil {
				return err
			}

			if ie.Srs_ResourceId_Ext_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_ResourceId_Ext_r16, mACParametersCommonExtSrsResourceIdExtR16Constraints); err != nil {
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
					{Name: "enhancedUuDRX-forSidelink-r17", Optional: true},
					{Name: "mg-ActivationRequestPRS-Meas-r17", Optional: true},
					{Name: "mg-ActivationCommPRS-Meas-r17", Optional: true},
					{Name: "intraCG-Prioritization-r17", Optional: true},
					{Name: "jointPrioritizationCG-Retx-Timer-r17", Optional: true},
					{Name: "survivalTime-r17", Optional: true},
					{Name: "lcg-ExtensionIAB-r17", Optional: true},
					{Name: "harq-FeedbackDisabled-r17", Optional: true},
					{Name: "uplink-Harq-ModeB-r17", Optional: true},
					{Name: "sr-TriggeredBy-TA-Report-r17", Optional: true},
					{Name: "extendedDRX-CycleInactive-r17", Optional: true},
					{Name: "simultaneousSR-PUSCH-DiffPUCCH-groups-r17", Optional: true},
					{Name: "lastTransmissionUL-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnhancedUuDRX_ForSidelink_r17 != nil, ie.Mg_ActivationRequestPRS_Meas_r17 != nil, ie.Mg_ActivationCommPRS_Meas_r17 != nil, ie.IntraCG_Prioritization_r17 != nil, ie.JointPrioritizationCG_Retx_Timer_r17 != nil, ie.SurvivalTime_r17 != nil, ie.Lcg_ExtensionIAB_r17 != nil, ie.Harq_FeedbackDisabled_r17 != nil, ie.Uplink_Harq_ModeB_r17 != nil, ie.Sr_TriggeredBy_TA_Report_r17 != nil, ie.ExtendedDRX_CycleInactive_r17 != nil, ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil, ie.LastTransmissionUL_r17 != nil}); err != nil {
				return err
			}

			if ie.EnhancedUuDRX_ForSidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedUuDRX_ForSidelink_r17, mACParametersCommonExtEnhancedUuDRXForSidelinkR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_ActivationRequestPRS_Meas_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_ActivationRequestPRS_Meas_r17, mACParametersCommonExtMgActivationRequestPRSMeasR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_ActivationCommPRS_Meas_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_ActivationCommPRS_Meas_r17, mACParametersCommonExtMgActivationCommPRSMeasR17Constraints); err != nil {
					return err
				}
			}

			if ie.IntraCG_Prioritization_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.IntraCG_Prioritization_r17, mACParametersCommonExtIntraCGPrioritizationR17Constraints); err != nil {
					return err
				}
			}

			if ie.JointPrioritizationCG_Retx_Timer_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.JointPrioritizationCG_Retx_Timer_r17, mACParametersCommonExtJointPrioritizationCGRetxTimerR17Constraints); err != nil {
					return err
				}
			}

			if ie.SurvivalTime_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SurvivalTime_r17, mACParametersCommonExtSurvivalTimeR17Constraints); err != nil {
					return err
				}
			}

			if ie.Lcg_ExtensionIAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Lcg_ExtensionIAB_r17, mACParametersCommonExtLcgExtensionIABR17Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_FeedbackDisabled_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Harq_FeedbackDisabled_r17, mACParametersCommonExtHarqFeedbackDisabledR17Constraints); err != nil {
					return err
				}
			}

			if ie.Uplink_Harq_ModeB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Uplink_Harq_ModeB_r17, mACParametersCommonExtUplinkHarqModeBR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sr_TriggeredBy_TA_Report_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sr_TriggeredBy_TA_Report_r17, mACParametersCommonExtSrTriggeredByTAReportR17Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedDRX_CycleInactive_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedDRX_CycleInactive_r17, mACParametersCommonExtExtendedDRXCycleInactiveR17Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17, mACParametersCommonExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints); err != nil {
					return err
				}
			}

			if ie.LastTransmissionUL_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.LastTransmissionUL_r17, mACParametersCommonExtLastTransmissionULR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 5:
		if hasExtGroup5 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "harq-RTT-TimerDL-ForNTN-MulticastMBS-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17 != nil}); err != nil {
				return err
			}

			if ie.Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17, mACParametersCommonExtHarqRTTTimerDLForNTNMulticastMBSR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 6:
		if hasExtGroup6 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sr-TriggeredByTA-ReportATG-r18", Optional: true},
					{Name: "extendedDRX-CycleInactive-r18", Optional: true},
					{Name: "additionalBS-Table-r18", Optional: true},
					{Name: "delayStatusReport-r18", Optional: true},
					{Name: "cg-RetransmissionMonitoringDisabling-r18", Optional: true},
					{Name: "non-IntegerDRX-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sr_TriggeredByTA_ReportATG_r18 != nil, ie.ExtendedDRX_CycleInactive_r18 != nil, ie.AdditionalBS_Table_r18 != nil, ie.DelayStatusReport_r18 != nil, ie.Cg_RetransmissionMonitoringDisabling_r18 != nil, ie.Non_IntegerDRX_r18 != nil}); err != nil {
				return err
			}

			if ie.Sr_TriggeredByTA_ReportATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sr_TriggeredByTA_ReportATG_r18, mACParametersCommonExtSrTriggeredByTAReportATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedDRX_CycleInactive_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedDRX_CycleInactive_r18, mACParametersCommonExtExtendedDRXCycleInactiveR18Constraints); err != nil {
					return err
				}
			}

			if ie.AdditionalBS_Table_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.AdditionalBS_Table_r18, mACParametersCommonExtAdditionalBSTableR18Constraints); err != nil {
					return err
				}
			}

			if ie.DelayStatusReport_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DelayStatusReport_r18, mACParametersCommonExtDelayStatusReportR18Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_RetransmissionMonitoringDisabling_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_RetransmissionMonitoringDisabling_r18, mACParametersCommonExtCgRetransmissionMonitoringDisablingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Non_IntegerDRX_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Non_IntegerDRX_r18, mACParametersCommonExtNonIntegerDRXR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 7:
		if hasExtGroup7 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "multipleEntryDelayStatusReport-r19", Optional: true},
					{Name: "lcp-PriorityAdjustment-r19", Optional: true},
					{Name: "ul-RateControl-r19", Optional: true},
					{Name: "ul-RateQuery-r19", Optional: true},
					{Name: "delayStatusReportNonDelayReportingData-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultipleEntryDelayStatusReport_r19 != nil, ie.Lcp_PriorityAdjustment_r19 != nil, ie.Ul_RateControl_r19 != nil, ie.Ul_RateQuery_r19 != nil, ie.DelayStatusReportNonDelayReportingData_r19 != nil}); err != nil {
				return err
			}

			if ie.MultipleEntryDelayStatusReport_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipleEntryDelayStatusReport_r19, mACParametersCommonExtMultipleEntryDelayStatusReportR19Constraints); err != nil {
					return err
				}
			}

			if ie.Lcp_PriorityAdjustment_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Lcp_PriorityAdjustment_r19, mACParametersCommonExtLcpPriorityAdjustmentR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_RateControl_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_RateControl_r19, mACParametersCommonExtUlRateControlR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_RateQuery_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_RateQuery_r19, mACParametersCommonExtUlRateQueryR19Constraints); err != nil {
					return err
				}
			}

			if ie.DelayStatusReportNonDelayReportingData_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.DelayStatusReportNonDelayReportingData_r19, mACParametersCommonExtDelayStatusReportNonDelayReportingDataR19Constraints); err != nil {
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

func (ie *MAC_ParametersCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mACParametersCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. lcp-Restriction: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mACParametersCommonLcpRestrictionConstraints)
			if err != nil {
				return err
			}
			ie.Lcp_Restriction = &idx
		}
	}

	// 4. dummy: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mACParametersCommonDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 5. lch-ToSCellRestriction: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mACParametersCommonLchToSCellRestrictionConstraints)
			if err != nil {
				return err
			}
			ie.Lch_ToSCellRestriction = &idx
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
				{Name: "recommendedBitRate", Optional: true},
				{Name: "recommendedBitRateQuery", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtRecommendedBitRateConstraints)
			if err != nil {
				return err
			}
			ie.RecommendedBitRate = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtRecommendedBitRateQueryConstraints)
			if err != nil {
				return err
			}
			ie.RecommendedBitRateQuery = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "recommendedBitRateMultiplier-r16", Optional: true},
				{Name: "preEmptiveBSR-r16", Optional: true},
				{Name: "autonomousTransmission-r16", Optional: true},
				{Name: "lch-PriorityBasedPrioritization-r16", Optional: true},
				{Name: "lch-ToConfiguredGrantMapping-r16", Optional: true},
				{Name: "lch-ToGrantPriorityRestriction-r16", Optional: true},
				{Name: "singlePHR-P-r16", Optional: true},
				{Name: "ul-LBT-FailureDetectionRecovery-r16", Optional: true},
				{Name: "tdd-MPE-P-MPR-Reporting-r16", Optional: true},
				{Name: "lcid-ExtensionIAB-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtRecommendedBitRateMultiplierR16Constraints)
			if err != nil {
				return err
			}
			ie.RecommendedBitRateMultiplier_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtPreEmptiveBSRR16Constraints)
			if err != nil {
				return err
			}
			ie.PreEmptiveBSR_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtAutonomousTransmissionR16Constraints)
			if err != nil {
				return err
			}
			ie.AutonomousTransmission_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLchPriorityBasedPrioritizationR16Constraints)
			if err != nil {
				return err
			}
			ie.Lch_PriorityBasedPrioritization_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLchToConfiguredGrantMappingR16Constraints)
			if err != nil {
				return err
			}
			ie.Lch_ToConfiguredGrantMapping_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLchToGrantPriorityRestrictionR16Constraints)
			if err != nil {
				return err
			}
			ie.Lch_ToGrantPriorityRestriction_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSinglePHRPR16Constraints)
			if err != nil {
				return err
			}
			ie.SinglePHR_P_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtUlLBTFailureDetectionRecoveryR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_LBT_FailureDetectionRecovery_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtTddMPEPMPRReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.Tdd_MPE_P_MPR_Reporting_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLcidExtensionIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Lcid_ExtensionIAB_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spCell-BFR-CBRA-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSpCellBFRCBRAR16Constraints)
			if err != nil {
				return err
			}
			ie.SpCell_BFR_CBRA_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srs-ResourceId-Ext-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSrsResourceIdExtR16Constraints)
			if err != nil {
				return err
			}
			ie.Srs_ResourceId_Ext_r16 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enhancedUuDRX-forSidelink-r17", Optional: true},
				{Name: "mg-ActivationRequestPRS-Meas-r17", Optional: true},
				{Name: "mg-ActivationCommPRS-Meas-r17", Optional: true},
				{Name: "intraCG-Prioritization-r17", Optional: true},
				{Name: "jointPrioritizationCG-Retx-Timer-r17", Optional: true},
				{Name: "survivalTime-r17", Optional: true},
				{Name: "lcg-ExtensionIAB-r17", Optional: true},
				{Name: "harq-FeedbackDisabled-r17", Optional: true},
				{Name: "uplink-Harq-ModeB-r17", Optional: true},
				{Name: "sr-TriggeredBy-TA-Report-r17", Optional: true},
				{Name: "extendedDRX-CycleInactive-r17", Optional: true},
				{Name: "simultaneousSR-PUSCH-DiffPUCCH-groups-r17", Optional: true},
				{Name: "lastTransmissionUL-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtEnhancedUuDRXForSidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedUuDRX_ForSidelink_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtMgActivationRequestPRSMeasR17Constraints)
			if err != nil {
				return err
			}
			ie.Mg_ActivationRequestPRS_Meas_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtMgActivationCommPRSMeasR17Constraints)
			if err != nil {
				return err
			}
			ie.Mg_ActivationCommPRS_Meas_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtIntraCGPrioritizationR17Constraints)
			if err != nil {
				return err
			}
			ie.IntraCG_Prioritization_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtJointPrioritizationCGRetxTimerR17Constraints)
			if err != nil {
				return err
			}
			ie.JointPrioritizationCG_Retx_Timer_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSurvivalTimeR17Constraints)
			if err != nil {
				return err
			}
			ie.SurvivalTime_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLcgExtensionIABR17Constraints)
			if err != nil {
				return err
			}
			ie.Lcg_ExtensionIAB_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtHarqFeedbackDisabledR17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_FeedbackDisabled_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtUplinkHarqModeBR17Constraints)
			if err != nil {
				return err
			}
			ie.Uplink_Harq_ModeB_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSrTriggeredByTAReportR17Constraints)
			if err != nil {
				return err
			}
			ie.Sr_TriggeredBy_TA_Report_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtExtendedDRXCycleInactiveR17Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedDRX_CycleInactive_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLastTransmissionULR17Constraints)
			if err != nil {
				return err
			}
			ie.LastTransmissionUL_r17 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "harq-RTT-TimerDL-ForNTN-MulticastMBS-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtHarqRTTTimerDLForNTNMulticastMBSR17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_RTT_TimerDL_ForNTN_MulticastMBS_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sr-TriggeredByTA-ReportATG-r18", Optional: true},
				{Name: "extendedDRX-CycleInactive-r18", Optional: true},
				{Name: "additionalBS-Table-r18", Optional: true},
				{Name: "delayStatusReport-r18", Optional: true},
				{Name: "cg-RetransmissionMonitoringDisabling-r18", Optional: true},
				{Name: "non-IntegerDRX-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtSrTriggeredByTAReportATGR18Constraints)
			if err != nil {
				return err
			}
			ie.Sr_TriggeredByTA_ReportATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtExtendedDRXCycleInactiveR18Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedDRX_CycleInactive_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtAdditionalBSTableR18Constraints)
			if err != nil {
				return err
			}
			ie.AdditionalBS_Table_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtDelayStatusReportR18Constraints)
			if err != nil {
				return err
			}
			ie.DelayStatusReport_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtCgRetransmissionMonitoringDisablingR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_RetransmissionMonitoringDisabling_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtNonIntegerDRXR18Constraints)
			if err != nil {
				return err
			}
			ie.Non_IntegerDRX_r18 = &v
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multipleEntryDelayStatusReport-r19", Optional: true},
				{Name: "lcp-PriorityAdjustment-r19", Optional: true},
				{Name: "ul-RateControl-r19", Optional: true},
				{Name: "ul-RateQuery-r19", Optional: true},
				{Name: "delayStatusReportNonDelayReportingData-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtMultipleEntryDelayStatusReportR19Constraints)
			if err != nil {
				return err
			}
			ie.MultipleEntryDelayStatusReport_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtLcpPriorityAdjustmentR19Constraints)
			if err != nil {
				return err
			}
			ie.Lcp_PriorityAdjustment_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtUlRateControlR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RateControl_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtUlRateQueryR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_RateQuery_r19 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mACParametersCommonExtDelayStatusReportNonDelayReportingDataR19Constraints)
			if err != nil {
				return err
			}
			ie.DelayStatusReportNonDelayReportingData_r19 = &v
		}
	}

	return nil
}
