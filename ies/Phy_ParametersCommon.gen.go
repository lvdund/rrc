// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersCommon (line 22787).

var phyParametersCommonConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "csi-RS-CFRA-ForHO", Optional: true},
		{Name: "dynamicPRB-BundlingDL", Optional: true},
		{Name: "sp-CSI-ReportPUCCH", Optional: true},
		{Name: "sp-CSI-ReportPUSCH", Optional: true},
		{Name: "nzp-CSI-RS-IntefMgmt", Optional: true},
		{Name: "type2-SP-CSI-Feedback-LongPUCCH", Optional: true},
		{Name: "precoderGranularityCORESET", Optional: true},
		{Name: "dynamicHARQ-ACK-Codebook", Optional: true},
		{Name: "semiStaticHARQ-ACK-Codebook", Optional: true},
		{Name: "spatialBundlingHARQ-ACK", Optional: true},
		{Name: "dynamicBetaOffsetInd-HARQ-ACK-CSI", Optional: true},
		{Name: "pucch-Repetition-F1-3-4", Optional: true},
		{Name: "ra-Type0-PUSCH", Optional: true},
		{Name: "dynamicSwitchRA-Type0-1-PDSCH", Optional: true},
		{Name: "dynamicSwitchRA-Type0-1-PUSCH", Optional: true},
		{Name: "pdsch-MappingTypeA", Optional: true},
		{Name: "pdsch-MappingTypeB", Optional: true},
		{Name: "interleavingVRB-ToPRB-PDSCH", Optional: true},
		{Name: "interSlotFreqHopping-PUSCH", Optional: true},
		{Name: "type1-PUSCH-RepetitionMultiSlots", Optional: true},
		{Name: "type2-PUSCH-RepetitionMultiSlots", Optional: true},
		{Name: "pusch-RepetitionMultiSlots", Optional: true},
		{Name: "pdsch-RepetitionMultiSlots", Optional: true},
		{Name: "downlinkSPS", Optional: true},
		{Name: "configuredUL-GrantType1", Optional: true},
		{Name: "configuredUL-GrantType2", Optional: true},
		{Name: "pre-EmptIndication-DL", Optional: true},
		{Name: "cbg-TransIndication-DL", Optional: true},
		{Name: "cbg-TransIndication-UL", Optional: true},
		{Name: "cbg-FlushIndication-DL", Optional: true},
		{Name: "dynamicHARQ-ACK-CodeB-CBG-Retx-DL", Optional: true},
		{Name: "rateMatchingResrcSetSemi-Static", Optional: true},
		{Name: "rateMatchingResrcSetDynamic", Optional: true},
		{Name: "bwp-SwitchingDelay", Optional: true},
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
		{Name: "extension-addition-8", Optional: true},
		{Name: "extension-addition-9", Optional: true},
		{Name: "extension-addition-10", Optional: true},
		{Name: "extension-addition-11", Optional: true},
		{Name: "extension-addition-12", Optional: true},
		{Name: "extension-addition-13", Optional: true},
		{Name: "extension-addition-14", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Csi_RS_CFRA_ForHO_Supported = 0
)

var phyParametersCommonCsiRSCFRAForHOConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Csi_RS_CFRA_ForHO_Supported},
}

const (
	Phy_ParametersCommon_DynamicPRB_BundlingDL_Supported = 0
)

var phyParametersCommonDynamicPRBBundlingDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicPRB_BundlingDL_Supported},
}

const (
	Phy_ParametersCommon_Sp_CSI_ReportPUCCH_Supported = 0
)

var phyParametersCommonSpCSIReportPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Sp_CSI_ReportPUCCH_Supported},
}

const (
	Phy_ParametersCommon_Sp_CSI_ReportPUSCH_Supported = 0
)

var phyParametersCommonSpCSIReportPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Sp_CSI_ReportPUSCH_Supported},
}

const (
	Phy_ParametersCommon_Nzp_CSI_RS_IntefMgmt_Supported = 0
)

var phyParametersCommonNzpCSIRSIntefMgmtConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Nzp_CSI_RS_IntefMgmt_Supported},
}

const (
	Phy_ParametersCommon_Type2_SP_CSI_Feedback_LongPUCCH_Supported = 0
)

var phyParametersCommonType2SPCSIFeedbackLongPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Type2_SP_CSI_Feedback_LongPUCCH_Supported},
}

const (
	Phy_ParametersCommon_PrecoderGranularityCORESET_Supported = 0
)

var phyParametersCommonPrecoderGranularityCORESETConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_PrecoderGranularityCORESET_Supported},
}

const (
	Phy_ParametersCommon_DynamicHARQ_ACK_Codebook_Supported = 0
)

var phyParametersCommonDynamicHARQACKCodebookConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicHARQ_ACK_Codebook_Supported},
}

const (
	Phy_ParametersCommon_SemiStaticHARQ_ACK_Codebook_Supported = 0
)

var phyParametersCommonSemiStaticHARQACKCodebookConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_SemiStaticHARQ_ACK_Codebook_Supported},
}

const (
	Phy_ParametersCommon_SpatialBundlingHARQ_ACK_Supported = 0
)

var phyParametersCommonSpatialBundlingHARQACKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_SpatialBundlingHARQ_ACK_Supported},
}

const (
	Phy_ParametersCommon_DynamicBetaOffsetInd_HARQ_ACK_CSI_Supported = 0
)

var phyParametersCommonDynamicBetaOffsetIndHARQACKCSIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicBetaOffsetInd_HARQ_ACK_CSI_Supported},
}

const (
	Phy_ParametersCommon_Pucch_Repetition_F1_3_4_Supported = 0
)

var phyParametersCommonPucchRepetitionF134Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pucch_Repetition_F1_3_4_Supported},
}

const (
	Phy_ParametersCommon_Ra_Type0_PUSCH_Supported = 0
)

var phyParametersCommonRaType0PUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ra_Type0_PUSCH_Supported},
}

const (
	Phy_ParametersCommon_DynamicSwitchRA_Type0_1_PDSCH_Supported = 0
)

var phyParametersCommonDynamicSwitchRAType01PDSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicSwitchRA_Type0_1_PDSCH_Supported},
}

const (
	Phy_ParametersCommon_DynamicSwitchRA_Type0_1_PUSCH_Supported = 0
)

var phyParametersCommonDynamicSwitchRAType01PUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicSwitchRA_Type0_1_PUSCH_Supported},
}

const (
	Phy_ParametersCommon_Pdsch_MappingTypeA_Supported = 0
)

var phyParametersCommonPdschMappingTypeAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pdsch_MappingTypeA_Supported},
}

const (
	Phy_ParametersCommon_Pdsch_MappingTypeB_Supported = 0
)

var phyParametersCommonPdschMappingTypeBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pdsch_MappingTypeB_Supported},
}

const (
	Phy_ParametersCommon_InterleavingVRB_ToPRB_PDSCH_Supported = 0
)

var phyParametersCommonInterleavingVRBToPRBPDSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_InterleavingVRB_ToPRB_PDSCH_Supported},
}

const (
	Phy_ParametersCommon_InterSlotFreqHopping_PUSCH_Supported = 0
)

var phyParametersCommonInterSlotFreqHoppingPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_InterSlotFreqHopping_PUSCH_Supported},
}

const (
	Phy_ParametersCommon_Type1_PUSCH_RepetitionMultiSlots_Supported = 0
)

var phyParametersCommonType1PUSCHRepetitionMultiSlotsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Type1_PUSCH_RepetitionMultiSlots_Supported},
}

const (
	Phy_ParametersCommon_Type2_PUSCH_RepetitionMultiSlots_Supported = 0
)

var phyParametersCommonType2PUSCHRepetitionMultiSlotsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Type2_PUSCH_RepetitionMultiSlots_Supported},
}

const (
	Phy_ParametersCommon_Pusch_RepetitionMultiSlots_Supported = 0
)

var phyParametersCommonPuschRepetitionMultiSlotsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pusch_RepetitionMultiSlots_Supported},
}

const (
	Phy_ParametersCommon_Pdsch_RepetitionMultiSlots_Supported = 0
)

var phyParametersCommonPdschRepetitionMultiSlotsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pdsch_RepetitionMultiSlots_Supported},
}

const (
	Phy_ParametersCommon_DownlinkSPS_Supported = 0
)

var phyParametersCommonDownlinkSPSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DownlinkSPS_Supported},
}

const (
	Phy_ParametersCommon_ConfiguredUL_GrantType1_Supported = 0
)

var phyParametersCommonConfiguredULGrantType1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_ConfiguredUL_GrantType1_Supported},
}

const (
	Phy_ParametersCommon_ConfiguredUL_GrantType2_Supported = 0
)

var phyParametersCommonConfiguredULGrantType2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_ConfiguredUL_GrantType2_Supported},
}

const (
	Phy_ParametersCommon_Pre_EmptIndication_DL_Supported = 0
)

var phyParametersCommonPreEmptIndicationDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Pre_EmptIndication_DL_Supported},
}

const (
	Phy_ParametersCommon_Cbg_TransIndication_DL_Supported = 0
)

var phyParametersCommonCbgTransIndicationDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Cbg_TransIndication_DL_Supported},
}

const (
	Phy_ParametersCommon_Cbg_TransIndication_UL_Supported = 0
)

var phyParametersCommonCbgTransIndicationULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Cbg_TransIndication_UL_Supported},
}

const (
	Phy_ParametersCommon_Cbg_FlushIndication_DL_Supported = 0
)

var phyParametersCommonCbgFlushIndicationDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Cbg_FlushIndication_DL_Supported},
}

const (
	Phy_ParametersCommon_DynamicHARQ_ACK_CodeB_CBG_Retx_DL_Supported = 0
)

var phyParametersCommonDynamicHARQACKCodeBCBGRetxDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_DynamicHARQ_ACK_CodeB_CBG_Retx_DL_Supported},
}

const (
	Phy_ParametersCommon_RateMatchingResrcSetSemi_Static_Supported = 0
)

var phyParametersCommonRateMatchingResrcSetSemiStaticConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_RateMatchingResrcSetSemi_Static_Supported},
}

const (
	Phy_ParametersCommon_RateMatchingResrcSetDynamic_Supported = 0
)

var phyParametersCommonRateMatchingResrcSetDynamicConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_RateMatchingResrcSetDynamic_Supported},
}

const (
	Phy_ParametersCommon_Bwp_SwitchingDelay_Type1 = 0
	Phy_ParametersCommon_Bwp_SwitchingDelay_Type2 = 1
)

var phyParametersCommonBwpSwitchingDelayConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Bwp_SwitchingDelay_Type1, Phy_ParametersCommon_Bwp_SwitchingDelay_Type2},
}

const (
	Phy_ParametersCommon_Ext_Dummy_Supported = 0
)

var phyParametersCommonExtDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy_Supported},
}

const (
	Phy_ParametersCommon_Ext_MaxNumberSearchSpaces_N10 = 0
)

var phyParametersCommonExtMaxNumberSearchSpacesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxNumberSearchSpaces_N10},
}

const (
	Phy_ParametersCommon_Ext_RateMatchingCtrlResrcSetDynamic_Supported = 0
)

var phyParametersCommonExtRateMatchingCtrlResrcSetDynamicConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_RateMatchingCtrlResrcSetDynamic_Supported},
}

const (
	Phy_ParametersCommon_Ext_MaxLayersMIMO_Indication_Supported = 0
)

var phyParametersCommonExtMaxLayersMIMOIndicationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxLayersMIMO_Indication_Supported},
}

const (
	Phy_ParametersCommon_Ext_TwoStepRACH_r16_Supported = 0
)

var phyParametersCommonExtTwoStepRACHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_TwoStepRACH_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dci_Format1_2And0_2_r16_Supported = 0
)

var phyParametersCommonExtDciFormat12And02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dci_Format1_2And0_2_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_MonitoringDCI_SameSearchSpace_r16_Supported = 0
)

var phyParametersCommonExtMonitoringDCISameSearchSpaceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MonitoringDCI_SameSearchSpace_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Type2_CG_ReleaseDCI_0_1_r16_Supported = 0
)

var phyParametersCommonExtType2CGReleaseDCI01R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Type2_CG_ReleaseDCI_0_1_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Type2_CG_ReleaseDCI_0_2_r16_Supported = 0
)

var phyParametersCommonExtType2CGReleaseDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Type2_CG_ReleaseDCI_0_2_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Sps_ReleaseDCI_1_1_r16_Supported = 0
)

var phyParametersCommonExtSpsReleaseDCI11R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Sps_ReleaseDCI_1_1_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Sps_ReleaseDCI_1_2_r16_Supported = 0
)

var phyParametersCommonExtSpsReleaseDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Sps_ReleaseDCI_1_2_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Csi_TriggerStateNon_ActiveBWP_r16_Supported = 0
)

var phyParametersCommonExtCsiTriggerStateNonActiveBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Csi_TriggerStateNon_ActiveBWP_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_SeparateSMTC_InterIAB_Support_r16_Supported = 0
)

var phyParametersCommonExtSeparateSMTCInterIABSupportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SeparateSMTC_InterIAB_Support_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_SeparateRACH_IAB_Support_r16_Supported = 0
)

var phyParametersCommonExtSeparateRACHIABSupportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SeparateRACH_IAB_Support_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16_Supported = 0
)

var phyParametersCommonExtUlFlexibleDLSlotFormatSemiStaticIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ul_FlexibleDL_SlotFormatDynamics_IAB_r16_Supported = 0
)

var phyParametersCommonExtUlFlexibleDLSlotFormatDynamicsIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ul_FlexibleDL_SlotFormatDynamics_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dft_S_OFDM_WaveformUL_IAB_r16_Supported = 0
)

var phyParametersCommonExtDftSOFDMWaveformULIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dft_S_OFDM_WaveformUL_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dci_25_AI_RNTI_Support_IAB_r16_Supported = 0
)

var phyParametersCommonExtDci25AIRNTISupportIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dci_25_AI_RNTI_Support_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_T_DeltaReceptionSupport_IAB_r16_Supported = 0
)

var phyParametersCommonExtTDeltaReceptionSupportIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_T_DeltaReceptionSupport_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_GuardSymbolReportReception_IAB_r16_Supported = 0
)

var phyParametersCommonExtGuardSymbolReportReceptionIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_GuardSymbolReportReception_IAB_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_HarqACK_CB_SpatialBundlingPUCCH_Group_r16_Supported = 0
)

var phyParametersCommonExtHarqACKCBSpatialBundlingPUCCHGroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_HarqACK_CB_SpatialBundlingPUCCH_Group_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N1  = 0
	Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N4  = 1
	Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N8  = 2
	Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N16 = 3
)

var phyParametersCommonExtMaxNumberSRSPosPathLossEstimateAllServingCellsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N1, Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N4, Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N8, Phy_ParametersCommon_Ext_MaxNumberSRS_PosPathLossEstimateAllServingCells_r16_N16},
}

const (
	Phy_ParametersCommon_Ext_ExtendedCG_Periodicities_r16_Supported = 0
)

var phyParametersCommonExtExtendedCGPeriodicitiesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_ExtendedCG_Periodicities_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_ExtendedSPS_Periodicities_r16_Supported = 0
)

var phyParametersCommonExtExtendedSPSPeriodicitiesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_ExtendedSPS_Periodicities_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dci_DL_PriorityIndicator_r16_Supported = 0
)

var phyParametersCommonExtDciDLPriorityIndicatorR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dci_DL_PriorityIndicator_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dci_UL_PriorityIndicator_r16_Supported = 0
)

var phyParametersCommonExtDciULPriorityIndicatorR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dci_UL_PriorityIndicator_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N4  = 0
	Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N8  = 1
	Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N16 = 2
	Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N32 = 3
	Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N64 = 4
)

var phyParametersCommonExtMaxNumberPathlossRSUpdateR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N4, Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N8, Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N16, Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N32, Phy_ParametersCommon_Ext_MaxNumberPathlossRS_Update_r16_N64},
}

const (
	Phy_ParametersCommon_Ext_Type2_HARQ_ACK_Codebook_r16_Supported = 0
)

var phyParametersCommonExtType2HARQACKCodebookR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Type2_HARQ_ACK_Codebook_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_HarqACK_JointMultiDCI_MultiTRP_r16_Supported = 0
)

var phyParametersCommonExtHarqACKJointMultiDCIMultiTRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_HarqACK_JointMultiDCI_MultiTRP_r16_Supported},
}

var phyParametersCommonExtBwpSwitchingMultiCCsR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1-r16"},
		{Name: "type2-r16"},
	},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16 = 1
)

const (
	Phy_ParametersCommon_Ext_TargetSMTC_SCG_r16_Supported = 0
)

var phyParametersCommonExtTargetSMTCSCGR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_TargetSMTC_SCG_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_SupportRepetitionZeroOffsetRV_r16_Supported = 0
)

var phyParametersCommonExtSupportRepetitionZeroOffsetRVR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SupportRepetitionZeroOffsetRV_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Cbg_TransInOrderPUSCH_UL_r16_Supported = 0
)

var phyParametersCommonExtCbgTransInOrderPUSCHULR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Cbg_TransInOrderPUSCH_UL_r16_Supported},
}

var phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1-r16"},
		{Name: "type2-r16"},
	},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16 = 1
)

const (
	Phy_ParametersCommon_Ext_SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16_NotSupported = 0
)

var phyParametersCommonExtSupportRetxDiffCoresetPoolMultiDCITRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16_NotSupported},
}

const (
	Phy_ParametersCommon_Ext_Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16_Mode2 = 0
	Phy_ParametersCommon_Ext_Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16_Mode3 = 1
)

var phyParametersCommonExtPdcchMonitoringAnyOccasionsWithSpanGapCrossCarrierSchR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16_Mode2, Phy_ParametersCommon_Ext_Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16_Mode3},
}

const (
	Phy_ParametersCommon_Ext_NewBeamIdentifications2PortCSI_RS_r16_Supported = 0
)

var phyParametersCommonExtNewBeamIdentifications2PortCSIRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_NewBeamIdentifications2PortCSI_RS_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_PathlossEstimation2PortCSI_RS_r16_Supported = 0
)

var phyParametersCommonExtPathlossEstimation2PortCSIRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PathlossEstimation2PortCSI_RS_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16_Supported = 0
)

var phyParametersCommonExtMuxHARQACKWithoutPUCCHOnPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_GuardSymbolReportReception_IAB_r17_Supported = 0
)

var phyParametersCommonExtGuardSymbolReportReceptionIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_GuardSymbolReportReception_IAB_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Restricted_IAB_DU_BeamReception_r17_Supported = 0
)

var phyParametersCommonExtRestrictedIABDUBeamReceptionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Restricted_IAB_DU_BeamReception_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Recommended_IAB_MT_BeamTransmission_r17_Supported = 0
)

var phyParametersCommonExtRecommendedIABMTBeamTransmissionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Recommended_IAB_MT_BeamTransmission_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Case6_TimingAlignmentReception_IAB_r17_Supported = 0
)

var phyParametersCommonExtCase6TimingAlignmentReceptionIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Case6_TimingAlignmentReception_IAB_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Case7_TimingAlignmentReception_IAB_r17_Supported = 0
)

var phyParametersCommonExtCase7TimingAlignmentReceptionIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Case7_TimingAlignmentReception_IAB_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dl_Tx_PowerAdjustment_IAB_r17_Supported = 0
)

var phyParametersCommonExtDlTxPowerAdjustmentIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dl_Tx_PowerAdjustment_IAB_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Desired_Ul_Tx_PowerAdjustment_r17_Supported = 0
)

var phyParametersCommonExtDesiredUlTxPowerAdjustmentR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Desired_Ul_Tx_PowerAdjustment_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Fdm_SoftResourceAvailability_DynamicIndication_r17_Supported = 0
)

var phyParametersCommonExtFdmSoftResourceAvailabilityDynamicIndicationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Fdm_SoftResourceAvailability_DynamicIndication_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Updated_T_DeltaRangeReception_r17_Supported = 0
)

var phyParametersCommonExtUpdatedTDeltaRangeReceptionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Updated_T_DeltaRangeReception_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_SlotBasedDynamicPUCCH_Rep_r17_Supported = 0
)

var phyParametersCommonExtSlotBasedDynamicPUCCHRepR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SlotBasedDynamicPUCCH_Rep_r17_Supported},
}

var phyParametersCommonUnifiedJointTCICommonUpdateR17Constraints = per.Constrained(1, 4)

const (
	Phy_ParametersCommon_Ext_MTRP_PDCCH_SingleSpan_r17_Supported = 0
)

var phyParametersCommonExtMTRPPDCCHSingleSpanR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MTRP_PDCCH_SingleSpan_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N2 = 0
	Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N3 = 1
	Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N4 = 2
)

var phyParametersCommonExtSupportedActivatedPRSProcessingWindowR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N2, Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N3, Phy_ParametersCommon_Ext_SupportedActivatedPRS_ProcessingWindow_r17_N4},
}

const (
	Phy_ParametersCommon_Ext_Cg_TimeDomainAllocationExtension_r17_Supported = 0
)

var phyParametersCommonExtCgTimeDomainAllocationExtensionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Cg_TimeDomainAllocationExtension_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17_Supported = 0
)

var phyParametersCommonExtTaBasedPDCTNNonSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_DirectionalCollisionDC_IAB_r17_Supported = 0
)

var phyParametersCommonExtDirectionalCollisionDCIABR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_DirectionalCollisionDC_IAB_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dummy1_Supported = 0
)

var phyParametersCommonExtDummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy1_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dummy2_Supported = 0
)

var phyParametersCommonExtDummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy2_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dummy3_Supported = 0
)

var phyParametersCommonExtDummy3Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy3_Supported},
}

const (
	Phy_ParametersCommon_Ext_Dummy4_Supported = 0
)

var phyParametersCommonExtDummy4Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy4_Supported},
}

const (
	Phy_ParametersCommon_Ext_Srs_AdditionalRepetition_r17_Supported = 0
)

var phyParametersCommonExtSrsAdditionalRepetitionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Srs_AdditionalRepetition_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Pusch_Repetition_CG_SDT_r17_Supported = 0
)

var phyParametersCommonExtPuschRepetitionCGSDTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Pusch_Repetition_CG_SDT_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_MultiPDSCH_PerSlotType1_CB_Support_r17_Supported = 0
)

var phyParametersCommonExtMultiPDSCHPerSlotType1CBSupportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MultiPDSCH_PerSlotType1_CB_Support_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_JointPowerSpatialAdaptation_r18_Supported = 0
)

var phyParametersCommonExtJointPowerSpatialAdaptationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_JointPowerSpatialAdaptation_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ncr_Semi_PersistentBeamInd_AccessLink_r18_Supported = 0
)

var phyParametersCommonExtNcrSemiPersistentBeamIndAccessLinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ncr_Semi_PersistentBeamInd_AccessLink_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ncr_SimultaneousUL_BackhaulAndC_Link_r18_Supported = 0
)

var phyParametersCommonExtNcrSimultaneousULBackhaulAndCLinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ncr_SimultaneousUL_BackhaulAndC_Link_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_NonUnifiedTCI = 0
	Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_UnifiedTCI    = 1
	Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_Both          = 2
)

var phyParametersCommonExtNcrBackhaulBeamIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_NonUnifiedTCI, Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_UnifiedTCI, Phy_ParametersCommon_Ext_Ncr_BackhaulBeamInd_r18_Both},
}

const (
	Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_NonUnifiedTCI = 0
	Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_UnifiedTCI    = 1
	Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_Both          = 2
)

var phyParametersCommonExtNcrAdaptiveBeamBackhaulAndCLinkR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_NonUnifiedTCI, Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_UnifiedTCI, Phy_ParametersCommon_Ext_Ncr_AdaptiveBeamBackhaulAndC_Link_r18_Both},
}

const (
	Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18_Supported = 0
)

var phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18_Supported = 0
)

var phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18_Supported = 0
)

var phyParametersCommonExtConfigurableType1AFieldsForDCI03And13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18_Supported = 0
)

var phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationDL_r18_Supported = 0
)

var phyParametersCommonExtPriorityIndicationDLR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationDL_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationUL_r18_Supported = 0
)

var phyParametersCommonExtPriorityIndicationULR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationUL_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_DynamicIndicationSchedulingRestriction_r18_Supported = 0
)

var phyParametersCommonExtDynamicIndicationSchedulingRestrictionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_DynamicIndicationSchedulingRestriction_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationOneSlotHARQ_r18_Supported = 0
)

var phyParametersCommonExtPriorityIndicationOneSlotHARQR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationOneSlotHARQ_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_MultiPUSCH_DCI_0_1_r18_Supported = 0
)

var phyParametersCommonExtMultiPUSCHDCI01R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MultiPUSCH_DCI_0_1_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_MultiPUSCH_DCI_0_2_r18_Supported = 0
)

var phyParametersCommonExtMultiPUSCHDCI02R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MultiPUSCH_DCI_0_2_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_PathlossRS_UpdateForType1CG_PUSCH_r18_Supported = 0
)

var phyParametersCommonExtPathlossRSUpdateForType1CGPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PathlossRS_UpdateForType1CG_PUSCH_r18_Supported},
}

var phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "type1-r18"},
		{Name: "type2-r18"},
	},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18 = 1
)

const (
	Phy_ParametersCommon_Ext_Ncr_Dft_S_OFDM_WaveformUL_r18_Supported = 0
)

var phyParametersCommonExtNcrDftSOFDMWaveformULR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Ncr_Dft_S_OFDM_WaveformUL_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19_Supported = 0
)

var phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13DiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19_Supported = 0
)

var phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03DiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19_Supported = 0
)

var phyParametersCommonExtConfigurableType1AFieldsForDCI03And13DiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19_Supported = 0
)

var phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03DiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationDL_Diff_r19_Supported = 0
)

var phyParametersCommonExtPriorityIndicationDLDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationDL_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationUL_Diff_r19_Supported = 0
)

var phyParametersCommonExtPriorityIndicationULDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationUL_Diff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_DynamicIndicationSchedulingRestrictionDiff_r19_Supported = 0
)

var phyParametersCommonExtDynamicIndicationSchedulingRestrictionDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_DynamicIndicationSchedulingRestrictionDiff_r19_Supported},
}

const (
	Phy_ParametersCommon_Ext_PriorityIndicationOneSlotHARQ_Diff_r19_Supported = 0
)

var phyParametersCommonExtPriorityIndicationOneSlotHARQDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_PriorityIndicationOneSlotHARQ_Diff_r19_Supported},
}

var phyParametersCommonExtCrossSlotSchedulingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "non-SharedSpectrumChAccess-r16", Optional: true},
		{Name: "sharedSpectrumChAccess-r16", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_CrossSlotScheduling_r16_Non_SharedSpectrumChAccess_r16_Supported = 0
)

var phyParametersCommonExtCrossSlotSchedulingR16NonSharedSpectrumChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_CrossSlotScheduling_r16_Non_SharedSpectrumChAccess_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_CrossSlotScheduling_r16_SharedSpectrumChAccess_r16_Supported = 0
)

var phyParametersCommonExtCrossSlotSchedulingR16SharedSpectrumChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_CrossSlotScheduling_r16_SharedSpectrumChAccess_r16_Supported},
}

var phyParametersCommonExtPuschRepetitionTypeAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sharedSpectrumChAccess-r16", Optional: true},
		{Name: "non-SharedSpectrumChAccess-r16", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_Pusch_RepetitionTypeA_r16_SharedSpectrumChAccess_r16_Supported = 0
)

var phyParametersCommonExtPuschRepetitionTypeAR16SharedSpectrumChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Pusch_RepetitionTypeA_r16_SharedSpectrumChAccess_r16_Supported},
}

const (
	Phy_ParametersCommon_Ext_Pusch_RepetitionTypeA_r16_Non_SharedSpectrumChAccess_r16_Supported = 0
)

var phyParametersCommonExtPuschRepetitionTypeAR16NonSharedSpectrumChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Pusch_RepetitionTypeA_r16_Non_SharedSpectrumChAccess_r16_Supported},
}

var phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberResWithinSlotAcrossCC-AcrossFR-r16", Optional: true},
		{Name: "maxNumberResAcrossCC-AcrossFR-r16", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N2   = 0
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N4   = 1
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N8   = 2
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N12  = 3
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N16  = 4
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N32  = 5
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N64  = 6
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N128 = 7
)

var phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResWithinSlotAcrossCCAcrossFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N2, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N4, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N8, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N12, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N16, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N32, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N64, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResWithinSlotAcrossCC_AcrossFR_r16_N128},
}

const (
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N2   = 0
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N4   = 1
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N8   = 2
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N12  = 3
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N16  = 4
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N32  = 5
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N40  = 6
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N48  = 7
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N64  = 8
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N72  = 9
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N80  = 10
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N96  = 11
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N128 = 12
	Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N256 = 13
)

var phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResAcrossCCAcrossFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N2, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N4, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N8, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N12, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N16, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N32, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N40, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N48, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N64, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N72, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N80, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N96, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N128, Phy_ParametersCommon_Ext_MaxTotalResourcesForAcrossFreqRanges_r16_MaxNumberResAcrossCC_AcrossFR_r16_N256},
}

var phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberLongPUCCHs-r16", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_LongAndLong   = 0
	Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_LongAndShort  = 1
	Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_ShortAndShort = 2
)

var phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16MaxNumberLongPUCCHsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_LongAndLong, Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_LongAndShort, Phy_ParametersCommon_Ext_HarqACK_SeparateMultiDCI_MultiTRP_r16_MaxNumberLongPUCCHs_r16_ShortAndShort},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16_Us100 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16_Us200 = 1
)

var phyParametersCommonExtBwpSwitchingMultiCCsR16Type1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16_Us100, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16_Us200},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us200  = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us400  = 1
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us800  = 2
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us1000 = 3
)

var phyParametersCommonExtBwpSwitchingMultiCCsR16Type2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us200, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us400, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us800, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16_Us1000},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16_Us100 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16_Us200 = 1
)

var phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16_Us100, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16_Us200},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us200  = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us400  = 1
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us800  = 2
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us1000 = 3
)

var phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us200, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us400, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us800, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16_Us1000},
}

var phyParametersCommonExtSpsHARQACKDeferralR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "non-SharedSpectrumChAccess-r17", Optional: true},
		{Name: "sharedSpectrumChAccess-r17", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_Sps_HARQ_ACK_Deferral_r17_Non_SharedSpectrumChAccess_r17_Supported = 0
)

var phyParametersCommonExtSpsHARQACKDeferralR17NonSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Sps_HARQ_ACK_Deferral_r17_Non_SharedSpectrumChAccess_r17_Supported},
}

const (
	Phy_ParametersCommon_Ext_Sps_HARQ_ACK_Deferral_r17_SharedSpectrumChAccess_r17_Supported = 0
)

var phyParametersCommonExtSpsHARQACKDeferralR17SharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Sps_HARQ_ACK_Deferral_r17_SharedSpectrumChAccess_r17_Supported},
}

var phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

var phyParametersCommonExtAdditionalSRPeriodicitiesR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	Phy_ParametersCommon_Ext_AdditionalSR_Periodicities_r18_Scs_30kHz_r18_Supported = 0
)

var phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_AdditionalSR_Periodicities_r18_Scs_30kHz_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_AdditionalSR_Periodicities_r18_Scs_120kHz_r18_Supported = 0
)

var phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_AdditionalSR_Periodicities_r18_Scs_120kHz_r18_Supported},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18_Us100 = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18_Us200 = 1
)

var phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18_Us100, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18_Us200},
}

const (
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us200  = 0
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us400  = 1
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us800  = 2
	Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us1000 = 3
)

var phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us200, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us400, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us800, Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18_Us1000},
}

const (
	Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option1 = 0
	Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option2 = 1
)

var phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option1, Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option2},
}

const (
	Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms5 = 0
	Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms3 = 1
)

var phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms5, Phy_ParametersCommon_Ext_Dummy_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms3},
}

type Phy_ParametersCommon struct {
	Csi_RS_CFRA_ForHO                          *int64
	DynamicPRB_BundlingDL                      *int64
	Sp_CSI_ReportPUCCH                         *int64
	Sp_CSI_ReportPUSCH                         *int64
	Nzp_CSI_RS_IntefMgmt                       *int64
	Type2_SP_CSI_Feedback_LongPUCCH            *int64
	PrecoderGranularityCORESET                 *int64
	DynamicHARQ_ACK_Codebook                   *int64
	SemiStaticHARQ_ACK_Codebook                *int64
	SpatialBundlingHARQ_ACK                    *int64
	DynamicBetaOffsetInd_HARQ_ACK_CSI          *int64
	Pucch_Repetition_F1_3_4                    *int64
	Ra_Type0_PUSCH                             *int64
	DynamicSwitchRA_Type0_1_PDSCH              *int64
	DynamicSwitchRA_Type0_1_PUSCH              *int64
	Pdsch_MappingTypeA                         *int64
	Pdsch_MappingTypeB                         *int64
	InterleavingVRB_ToPRB_PDSCH                *int64
	InterSlotFreqHopping_PUSCH                 *int64
	Type1_PUSCH_RepetitionMultiSlots           *int64
	Type2_PUSCH_RepetitionMultiSlots           *int64
	Pusch_RepetitionMultiSlots                 *int64
	Pdsch_RepetitionMultiSlots                 *int64
	DownlinkSPS                                *int64
	ConfiguredUL_GrantType1                    *int64
	ConfiguredUL_GrantType2                    *int64
	Pre_EmptIndication_DL                      *int64
	Cbg_TransIndication_DL                     *int64
	Cbg_TransIndication_UL                     *int64
	Cbg_FlushIndication_DL                     *int64
	DynamicHARQ_ACK_CodeB_CBG_Retx_DL          *int64
	RateMatchingResrcSetSemi_Static            *int64
	RateMatchingResrcSetDynamic                *int64
	Bwp_SwitchingDelay                         *int64
	Dummy                                      *int64
	MaxNumberSearchSpaces                      *int64
	RateMatchingCtrlResrcSetDynamic            *int64
	MaxLayersMIMO_Indication                   *int64
	SpCellPlacement                            *CarrierAggregationVariant
	TwoStepRACH_r16                            *int64
	Dci_Format1_2And0_2_r16                    *int64
	MonitoringDCI_SameSearchSpace_r16          *int64
	Type2_CG_ReleaseDCI_0_1_r16                *int64
	Type2_CG_ReleaseDCI_0_2_r16                *int64
	Sps_ReleaseDCI_1_1_r16                     *int64
	Sps_ReleaseDCI_1_2_r16                     *int64
	Csi_TriggerStateNon_ActiveBWP_r16          *int64
	SeparateSMTC_InterIAB_Support_r16          *int64
	SeparateRACH_IAB_Support_r16               *int64
	Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16 *int64
	Ul_FlexibleDL_SlotFormatDynamics_IAB_r16   *int64
	Dft_S_OFDM_WaveformUL_IAB_r16              *int64
	Dci_25_AI_RNTI_Support_IAB_r16             *int64
	T_DeltaReceptionSupport_IAB_r16            *int64
	GuardSymbolReportReception_IAB_r16         *int64
	HarqACK_CB_SpatialBundlingPUCCH_Group_r16  *int64
	CrossSlotScheduling_r16                    *struct {
		Non_SharedSpectrumChAccess_r16 *int64
		SharedSpectrumChAccess_r16     *int64
	}
	MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 *int64
	ExtendedCG_Periodicities_r16                        *int64
	ExtendedSPS_Periodicities_r16                       *int64
	CodebookVariantsList_r16                            *CodebookVariantsList_r16
	Pusch_RepetitionTypeA_r16                           *struct {
		SharedSpectrumChAccess_r16     *int64
		Non_SharedSpectrumChAccess_r16 *int64
	}
	Dci_DL_PriorityIndicator_r16             *int64
	Dci_UL_PriorityIndicator_r16             *int64
	MaxNumberPathlossRS_Update_r16           *int64
	Type2_HARQ_ACK_Codebook_r16              *int64
	MaxTotalResourcesForAcrossFreqRanges_r16 *struct {
		MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 *int64
		MaxNumberResAcrossCC_AcrossFR_r16           *int64
	}
	HarqACK_SeparateMultiDCI_MultiTRP_r16 *struct{ MaxNumberLongPUCCHs_r16 *int64 }
	HarqACK_JointMultiDCI_MultiTRP_r16    *int64
	Bwp_SwitchingMultiCCs_r16             *struct {
		Choice    int
		Type1_r16 *int64
		Type2_r16 *int64
	}
	TargetSMTC_SCG_r16                *int64
	SupportRepetitionZeroOffsetRV_r16 *int64
	Cbg_TransInOrderPUSCH_UL_r16      *int64
	Bwp_SwitchingMultiDormancyCCs_r16 *struct {
		Choice    int
		Type1_r16 *int64
		Type2_r16 *int64
	}
	SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16             *int64
	Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 *int64
	NewBeamIdentifications2PortCSI_RS_r16                      *int64
	PathlossEstimation2PortCSI_RS_r16                          *int64
	Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16                      *int64
	GuardSymbolReportReception_IAB_r17                         *int64
	Restricted_IAB_DU_BeamReception_r17                        *int64
	Recommended_IAB_MT_BeamTransmission_r17                    *int64
	Case6_TimingAlignmentReception_IAB_r17                     *int64
	Case7_TimingAlignmentReception_IAB_r17                     *int64
	Dl_Tx_PowerAdjustment_IAB_r17                              *int64
	Desired_Ul_Tx_PowerAdjustment_r17                          *int64
	Fdm_SoftResourceAvailability_DynamicIndication_r17         *int64
	Updated_T_DeltaRangeReception_r17                          *int64
	SlotBasedDynamicPUCCH_Rep_r17                              *int64
	Sps_HARQ_ACK_Deferral_r17                                  *struct {
		Non_SharedSpectrumChAccess_r17 *int64
		SharedSpectrumChAccess_r17     *int64
	}
	UnifiedJointTCI_CommonUpdate_r17             *int64
	MTRP_PDCCH_SingleSpan_r17                    *int64
	SupportedActivatedPRS_ProcessingWindow_r17   *int64
	Cg_TimeDomainAllocationExtension_r17         *int64
	Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 *int64
	DirectionalCollisionDC_IAB_r17               *int64
	Dummy1                                       *int64
	Dummy2                                       *int64
	Dummy3                                       *int64
	Dummy4                                       *int64
	Srs_AdditionalRepetition_r17                 *int64
	Pusch_Repetition_CG_SDT_r17                  *int64
	MultiPDSCH_PerSlotType1_CB_Support_r17       *int64
	JointPowerSpatialAdaptation_r18              *int64
	Ncr_AperiodicBeamInd_AccessLink_r18          *struct {
		Scs_15kHz_r18  *int64
		Scs_30kHz_r18  *int64
		Scs_60kHz_r18  *int64
		Scs_120kHz_r18 *int64
	}
	Ncr_Semi_PersistentBeamInd_AccessLink_r18            *int64
	Ncr_SimultaneousUL_BackhaulAndC_Link_r18             *int64
	Ncr_BackhaulBeamInd_r18                              *int64
	Ncr_AdaptiveBeamBackhaulAndC_Link_r18                *int64
	NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18    *int64
	NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18    *int64
	ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18     *int64
	Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18 *int64
	PriorityIndicationDL_r18                             *int64
	PriorityIndicationUL_r18                             *int64
	DynamicIndicationSchedulingRestriction_r18           *int64
	PriorityIndicationOneSlotHARQ_r18                    *int64
	MultiPUSCH_DCI_0_1_r18                               *int64
	MultiPUSCH_DCI_0_2_r18                               *int64
	AdditionalSR_Periodicities_r18                       *struct {
		Scs_30kHz_r18  *int64
		Scs_120kHz_r18 *int64
	}
	PathlossRS_UpdateForType1CG_PUSCH_r18            *int64
	Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18 *struct {
		Choice    int
		Type1_r18 *int64
		Type2_r18 *int64
	}
	Ncr_Dft_S_OFDM_WaveformUL_r18      *int64
	CodebookVariantsListExt_r19        *CodebookVariantsListExt_r19
	CodebookVariantsListAggregate_r19  *CodebookVariantsListAggregate_r19
	CodebookVariantsListHybrid_r19     *CodebookVariantsListHybrid_r19
	Dummy_EnableTx_RxDuringMeasGap_r19 *struct {
		AdditionalDCI_r19     per.BitString
		IndicationField_r19   int64
		MinimumTimeOffset_r19 int64
	}
	NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19    *int64
	NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19    *int64
	ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19     *int64
	Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19 *int64
	PriorityIndicationDL_Diff_r19                             *int64
	PriorityIndicationUL_Diff_r19                             *int64
	DynamicIndicationSchedulingRestrictionDiff_r19            *int64
	PriorityIndicationOneSlotHARQ_Diff_r19                    *int64
}

func (ie *Phy_ParametersCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersCommonConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dummy != nil
	hasExtGroup1 := ie.MaxNumberSearchSpaces != nil || ie.RateMatchingCtrlResrcSetDynamic != nil || ie.MaxLayersMIMO_Indication != nil
	hasExtGroup2 := ie.SpCellPlacement != nil
	hasExtGroup3 := ie.TwoStepRACH_r16 != nil || ie.Dci_Format1_2And0_2_r16 != nil || ie.MonitoringDCI_SameSearchSpace_r16 != nil || ie.Type2_CG_ReleaseDCI_0_1_r16 != nil || ie.Type2_CG_ReleaseDCI_0_2_r16 != nil || ie.Sps_ReleaseDCI_1_1_r16 != nil || ie.Sps_ReleaseDCI_1_2_r16 != nil || ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil || ie.SeparateSMTC_InterIAB_Support_r16 != nil || ie.SeparateRACH_IAB_Support_r16 != nil || ie.Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16 != nil || ie.Ul_FlexibleDL_SlotFormatDynamics_IAB_r16 != nil || ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil || ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil || ie.T_DeltaReceptionSupport_IAB_r16 != nil || ie.GuardSymbolReportReception_IAB_r16 != nil || ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil || ie.CrossSlotScheduling_r16 != nil || ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil || ie.ExtendedCG_Periodicities_r16 != nil || ie.ExtendedSPS_Periodicities_r16 != nil || ie.CodebookVariantsList_r16 != nil || ie.Pusch_RepetitionTypeA_r16 != nil || ie.Dci_DL_PriorityIndicator_r16 != nil || ie.Dci_UL_PriorityIndicator_r16 != nil || ie.MaxNumberPathlossRS_Update_r16 != nil || ie.Type2_HARQ_ACK_Codebook_r16 != nil || ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil || ie.HarqACK_SeparateMultiDCI_MultiTRP_r16 != nil || ie.HarqACK_JointMultiDCI_MultiTRP_r16 != nil || ie.Bwp_SwitchingMultiCCs_r16 != nil
	hasExtGroup4 := ie.TargetSMTC_SCG_r16 != nil || ie.SupportRepetitionZeroOffsetRV_r16 != nil || ie.Cbg_TransInOrderPUSCH_UL_r16 != nil
	hasExtGroup5 := ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil || ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil || ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil
	hasExtGroup6 := ie.NewBeamIdentifications2PortCSI_RS_r16 != nil || ie.PathlossEstimation2PortCSI_RS_r16 != nil
	hasExtGroup7 := ie.Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16 != nil
	hasExtGroup8 := ie.GuardSymbolReportReception_IAB_r17 != nil || ie.Restricted_IAB_DU_BeamReception_r17 != nil || ie.Recommended_IAB_MT_BeamTransmission_r17 != nil || ie.Case6_TimingAlignmentReception_IAB_r17 != nil || ie.Case7_TimingAlignmentReception_IAB_r17 != nil || ie.Dl_Tx_PowerAdjustment_IAB_r17 != nil || ie.Desired_Ul_Tx_PowerAdjustment_r17 != nil || ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil || ie.Updated_T_DeltaRangeReception_r17 != nil || ie.SlotBasedDynamicPUCCH_Rep_r17 != nil || ie.Sps_HARQ_ACK_Deferral_r17 != nil || ie.UnifiedJointTCI_CommonUpdate_r17 != nil || ie.MTRP_PDCCH_SingleSpan_r17 != nil || ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil || ie.Cg_TimeDomainAllocationExtension_r17 != nil
	hasExtGroup9 := ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil || ie.DirectionalCollisionDC_IAB_r17 != nil
	hasExtGroup10 := ie.Dummy1 != nil || ie.Dummy2 != nil || ie.Dummy3 != nil || ie.Dummy4 != nil || ie.Srs_AdditionalRepetition_r17 != nil || ie.Pusch_Repetition_CG_SDT_r17 != nil
	hasExtGroup11 := ie.MultiPDSCH_PerSlotType1_CB_Support_r17 != nil
	hasExtGroup12 := ie.JointPowerSpatialAdaptation_r18 != nil || ie.Ncr_AperiodicBeamInd_AccessLink_r18 != nil || ie.Ncr_Semi_PersistentBeamInd_AccessLink_r18 != nil || ie.Ncr_SimultaneousUL_BackhaulAndC_Link_r18 != nil || ie.Ncr_BackhaulBeamInd_r18 != nil || ie.Ncr_AdaptiveBeamBackhaulAndC_Link_r18 != nil || ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18 != nil || ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18 != nil || ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18 != nil || ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18 != nil || ie.PriorityIndicationDL_r18 != nil || ie.PriorityIndicationUL_r18 != nil || ie.DynamicIndicationSchedulingRestriction_r18 != nil || ie.PriorityIndicationOneSlotHARQ_r18 != nil || ie.MultiPUSCH_DCI_0_1_r18 != nil || ie.MultiPUSCH_DCI_0_2_r18 != nil || ie.AdditionalSR_Periodicities_r18 != nil || ie.PathlossRS_UpdateForType1CG_PUSCH_r18 != nil || ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18 != nil
	hasExtGroup13 := ie.Ncr_Dft_S_OFDM_WaveformUL_r18 != nil
	hasExtGroup14 := ie.CodebookVariantsListExt_r19 != nil || ie.CodebookVariantsListAggregate_r19 != nil || ie.CodebookVariantsListHybrid_r19 != nil || ie.Dummy_EnableTx_RxDuringMeasGap_r19 != nil || ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19 != nil || ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19 != nil || ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19 != nil || ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19 != nil || ie.PriorityIndicationDL_Diff_r19 != nil || ie.PriorityIndicationUL_Diff_r19 != nil || ie.DynamicIndicationSchedulingRestrictionDiff_r19 != nil || ie.PriorityIndicationOneSlotHARQ_Diff_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12 || hasExtGroup13 || hasExtGroup14

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Csi_RS_CFRA_ForHO != nil, ie.DynamicPRB_BundlingDL != nil, ie.Sp_CSI_ReportPUCCH != nil, ie.Sp_CSI_ReportPUSCH != nil, ie.Nzp_CSI_RS_IntefMgmt != nil, ie.Type2_SP_CSI_Feedback_LongPUCCH != nil, ie.PrecoderGranularityCORESET != nil, ie.DynamicHARQ_ACK_Codebook != nil, ie.SemiStaticHARQ_ACK_Codebook != nil, ie.SpatialBundlingHARQ_ACK != nil, ie.DynamicBetaOffsetInd_HARQ_ACK_CSI != nil, ie.Pucch_Repetition_F1_3_4 != nil, ie.Ra_Type0_PUSCH != nil, ie.DynamicSwitchRA_Type0_1_PDSCH != nil, ie.DynamicSwitchRA_Type0_1_PUSCH != nil, ie.Pdsch_MappingTypeA != nil, ie.Pdsch_MappingTypeB != nil, ie.InterleavingVRB_ToPRB_PDSCH != nil, ie.InterSlotFreqHopping_PUSCH != nil, ie.Type1_PUSCH_RepetitionMultiSlots != nil, ie.Type2_PUSCH_RepetitionMultiSlots != nil, ie.Pusch_RepetitionMultiSlots != nil, ie.Pdsch_RepetitionMultiSlots != nil, ie.DownlinkSPS != nil, ie.ConfiguredUL_GrantType1 != nil, ie.ConfiguredUL_GrantType2 != nil, ie.Pre_EmptIndication_DL != nil, ie.Cbg_TransIndication_DL != nil, ie.Cbg_TransIndication_UL != nil, ie.Cbg_FlushIndication_DL != nil, ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil, ie.RateMatchingResrcSetSemi_Static != nil, ie.RateMatchingResrcSetDynamic != nil, ie.Bwp_SwitchingDelay != nil}); err != nil {
		return err
	}

	// 3. csi-RS-CFRA-ForHO: enumerated
	{
		if ie.Csi_RS_CFRA_ForHO != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_CFRA_ForHO, phyParametersCommonCsiRSCFRAForHOConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dynamicPRB-BundlingDL: enumerated
	{
		if ie.DynamicPRB_BundlingDL != nil {
			if err := e.EncodeEnumerated(*ie.DynamicPRB_BundlingDL, phyParametersCommonDynamicPRBBundlingDLConstraints); err != nil {
				return err
			}
		}
	}

	// 5. sp-CSI-ReportPUCCH: enumerated
	{
		if ie.Sp_CSI_ReportPUCCH != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_ReportPUCCH, phyParametersCommonSpCSIReportPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 6. sp-CSI-ReportPUSCH: enumerated
	{
		if ie.Sp_CSI_ReportPUSCH != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_ReportPUSCH, phyParametersCommonSpCSIReportPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 7. nzp-CSI-RS-IntefMgmt: enumerated
	{
		if ie.Nzp_CSI_RS_IntefMgmt != nil {
			if err := e.EncodeEnumerated(*ie.Nzp_CSI_RS_IntefMgmt, phyParametersCommonNzpCSIRSIntefMgmtConstraints); err != nil {
				return err
			}
		}
	}

	// 8. type2-SP-CSI-Feedback-LongPUCCH: enumerated
	{
		if ie.Type2_SP_CSI_Feedback_LongPUCCH != nil {
			if err := e.EncodeEnumerated(*ie.Type2_SP_CSI_Feedback_LongPUCCH, phyParametersCommonType2SPCSIFeedbackLongPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 9. precoderGranularityCORESET: enumerated
	{
		if ie.PrecoderGranularityCORESET != nil {
			if err := e.EncodeEnumerated(*ie.PrecoderGranularityCORESET, phyParametersCommonPrecoderGranularityCORESETConstraints); err != nil {
				return err
			}
		}
	}

	// 10. dynamicHARQ-ACK-Codebook: enumerated
	{
		if ie.DynamicHARQ_ACK_Codebook != nil {
			if err := e.EncodeEnumerated(*ie.DynamicHARQ_ACK_Codebook, phyParametersCommonDynamicHARQACKCodebookConstraints); err != nil {
				return err
			}
		}
	}

	// 11. semiStaticHARQ-ACK-Codebook: enumerated
	{
		if ie.SemiStaticHARQ_ACK_Codebook != nil {
			if err := e.EncodeEnumerated(*ie.SemiStaticHARQ_ACK_Codebook, phyParametersCommonSemiStaticHARQACKCodebookConstraints); err != nil {
				return err
			}
		}
	}

	// 12. spatialBundlingHARQ-ACK: enumerated
	{
		if ie.SpatialBundlingHARQ_ACK != nil {
			if err := e.EncodeEnumerated(*ie.SpatialBundlingHARQ_ACK, phyParametersCommonSpatialBundlingHARQACKConstraints); err != nil {
				return err
			}
		}
	}

	// 13. dynamicBetaOffsetInd-HARQ-ACK-CSI: enumerated
	{
		if ie.DynamicBetaOffsetInd_HARQ_ACK_CSI != nil {
			if err := e.EncodeEnumerated(*ie.DynamicBetaOffsetInd_HARQ_ACK_CSI, phyParametersCommonDynamicBetaOffsetIndHARQACKCSIConstraints); err != nil {
				return err
			}
		}
	}

	// 14. pucch-Repetition-F1-3-4: enumerated
	{
		if ie.Pucch_Repetition_F1_3_4 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_Repetition_F1_3_4, phyParametersCommonPucchRepetitionF134Constraints); err != nil {
				return err
			}
		}
	}

	// 15. ra-Type0-PUSCH: enumerated
	{
		if ie.Ra_Type0_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.Ra_Type0_PUSCH, phyParametersCommonRaType0PUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 16. dynamicSwitchRA-Type0-1-PDSCH: enumerated
	{
		if ie.DynamicSwitchRA_Type0_1_PDSCH != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSwitchRA_Type0_1_PDSCH, phyParametersCommonDynamicSwitchRAType01PDSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 17. dynamicSwitchRA-Type0-1-PUSCH: enumerated
	{
		if ie.DynamicSwitchRA_Type0_1_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSwitchRA_Type0_1_PUSCH, phyParametersCommonDynamicSwitchRAType01PUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 18. pdsch-MappingTypeA: enumerated
	{
		if ie.Pdsch_MappingTypeA != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_MappingTypeA, phyParametersCommonPdschMappingTypeAConstraints); err != nil {
				return err
			}
		}
	}

	// 19. pdsch-MappingTypeB: enumerated
	{
		if ie.Pdsch_MappingTypeB != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_MappingTypeB, phyParametersCommonPdschMappingTypeBConstraints); err != nil {
				return err
			}
		}
	}

	// 20. interleavingVRB-ToPRB-PDSCH: enumerated
	{
		if ie.InterleavingVRB_ToPRB_PDSCH != nil {
			if err := e.EncodeEnumerated(*ie.InterleavingVRB_ToPRB_PDSCH, phyParametersCommonInterleavingVRBToPRBPDSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 21. interSlotFreqHopping-PUSCH: enumerated
	{
		if ie.InterSlotFreqHopping_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.InterSlotFreqHopping_PUSCH, phyParametersCommonInterSlotFreqHoppingPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 22. type1-PUSCH-RepetitionMultiSlots: enumerated
	{
		if ie.Type1_PUSCH_RepetitionMultiSlots != nil {
			if err := e.EncodeEnumerated(*ie.Type1_PUSCH_RepetitionMultiSlots, phyParametersCommonType1PUSCHRepetitionMultiSlotsConstraints); err != nil {
				return err
			}
		}
	}

	// 23. type2-PUSCH-RepetitionMultiSlots: enumerated
	{
		if ie.Type2_PUSCH_RepetitionMultiSlots != nil {
			if err := e.EncodeEnumerated(*ie.Type2_PUSCH_RepetitionMultiSlots, phyParametersCommonType2PUSCHRepetitionMultiSlotsConstraints); err != nil {
				return err
			}
		}
	}

	// 24. pusch-RepetitionMultiSlots: enumerated
	{
		if ie.Pusch_RepetitionMultiSlots != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_RepetitionMultiSlots, phyParametersCommonPuschRepetitionMultiSlotsConstraints); err != nil {
				return err
			}
		}
	}

	// 25. pdsch-RepetitionMultiSlots: enumerated
	{
		if ie.Pdsch_RepetitionMultiSlots != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_RepetitionMultiSlots, phyParametersCommonPdschRepetitionMultiSlotsConstraints); err != nil {
				return err
			}
		}
	}

	// 26. downlinkSPS: enumerated
	{
		if ie.DownlinkSPS != nil {
			if err := e.EncodeEnumerated(*ie.DownlinkSPS, phyParametersCommonDownlinkSPSConstraints); err != nil {
				return err
			}
		}
	}

	// 27. configuredUL-GrantType1: enumerated
	{
		if ie.ConfiguredUL_GrantType1 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredUL_GrantType1, phyParametersCommonConfiguredULGrantType1Constraints); err != nil {
				return err
			}
		}
	}

	// 28. configuredUL-GrantType2: enumerated
	{
		if ie.ConfiguredUL_GrantType2 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredUL_GrantType2, phyParametersCommonConfiguredULGrantType2Constraints); err != nil {
				return err
			}
		}
	}

	// 29. pre-EmptIndication-DL: enumerated
	{
		if ie.Pre_EmptIndication_DL != nil {
			if err := e.EncodeEnumerated(*ie.Pre_EmptIndication_DL, phyParametersCommonPreEmptIndicationDLConstraints); err != nil {
				return err
			}
		}
	}

	// 30. cbg-TransIndication-DL: enumerated
	{
		if ie.Cbg_TransIndication_DL != nil {
			if err := e.EncodeEnumerated(*ie.Cbg_TransIndication_DL, phyParametersCommonCbgTransIndicationDLConstraints); err != nil {
				return err
			}
		}
	}

	// 31. cbg-TransIndication-UL: enumerated
	{
		if ie.Cbg_TransIndication_UL != nil {
			if err := e.EncodeEnumerated(*ie.Cbg_TransIndication_UL, phyParametersCommonCbgTransIndicationULConstraints); err != nil {
				return err
			}
		}
	}

	// 32. cbg-FlushIndication-DL: enumerated
	{
		if ie.Cbg_FlushIndication_DL != nil {
			if err := e.EncodeEnumerated(*ie.Cbg_FlushIndication_DL, phyParametersCommonCbgFlushIndicationDLConstraints); err != nil {
				return err
			}
		}
	}

	// 33. dynamicHARQ-ACK-CodeB-CBG-Retx-DL: enumerated
	{
		if ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil {
			if err := e.EncodeEnumerated(*ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL, phyParametersCommonDynamicHARQACKCodeBCBGRetxDLConstraints); err != nil {
				return err
			}
		}
	}

	// 34. rateMatchingResrcSetSemi-Static: enumerated
	{
		if ie.RateMatchingResrcSetSemi_Static != nil {
			if err := e.EncodeEnumerated(*ie.RateMatchingResrcSetSemi_Static, phyParametersCommonRateMatchingResrcSetSemiStaticConstraints); err != nil {
				return err
			}
		}
	}

	// 35. rateMatchingResrcSetDynamic: enumerated
	{
		if ie.RateMatchingResrcSetDynamic != nil {
			if err := e.EncodeEnumerated(*ie.RateMatchingResrcSetDynamic, phyParametersCommonRateMatchingResrcSetDynamicConstraints); err != nil {
				return err
			}
		}
	}

	// 36. bwp-SwitchingDelay: enumerated
	{
		if ie.Bwp_SwitchingDelay != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_SwitchingDelay, phyParametersCommonBwpSwitchingDelayConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12, hasExtGroup13, hasExtGroup14}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dummy", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy != nil}); err != nil {
				return err
			}

			if ie.Dummy != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy, phyParametersCommonExtDummyConstraints); err != nil {
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
					{Name: "maxNumberSearchSpaces", Optional: true},
					{Name: "rateMatchingCtrlResrcSetDynamic", Optional: true},
					{Name: "maxLayersMIMO-Indication", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxNumberSearchSpaces != nil, ie.RateMatchingCtrlResrcSetDynamic != nil, ie.MaxLayersMIMO_Indication != nil}); err != nil {
				return err
			}

			if ie.MaxNumberSearchSpaces != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberSearchSpaces, phyParametersCommonExtMaxNumberSearchSpacesConstraints); err != nil {
					return err
				}
			}

			if ie.RateMatchingCtrlResrcSetDynamic != nil {
				if err := ex.EncodeEnumerated(*ie.RateMatchingCtrlResrcSetDynamic, phyParametersCommonExtRateMatchingCtrlResrcSetDynamicConstraints); err != nil {
					return err
				}
			}

			if ie.MaxLayersMIMO_Indication != nil {
				if err := ex.EncodeEnumerated(*ie.MaxLayersMIMO_Indication, phyParametersCommonExtMaxLayersMIMOIndicationConstraints); err != nil {
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
					{Name: "spCellPlacement", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpCellPlacement != nil}); err != nil {
				return err
			}

			if ie.SpCellPlacement != nil {
				if err := ie.SpCellPlacement.Encode(ex); err != nil {
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
					{Name: "twoStepRACH-r16", Optional: true},
					{Name: "dci-Format1-2And0-2-r16", Optional: true},
					{Name: "monitoringDCI-SameSearchSpace-r16", Optional: true},
					{Name: "type2-CG-ReleaseDCI-0-1-r16", Optional: true},
					{Name: "type2-CG-ReleaseDCI-0-2-r16", Optional: true},
					{Name: "sps-ReleaseDCI-1-1-r16", Optional: true},
					{Name: "sps-ReleaseDCI-1-2-r16", Optional: true},
					{Name: "csi-TriggerStateNon-ActiveBWP-r16", Optional: true},
					{Name: "separateSMTC-InterIAB-Support-r16", Optional: true},
					{Name: "separateRACH-IAB-Support-r16", Optional: true},
					{Name: "ul-flexibleDL-SlotFormatSemiStatic-IAB-r16", Optional: true},
					{Name: "ul-flexibleDL-SlotFormatDynamics-IAB-r16", Optional: true},
					{Name: "dft-S-OFDM-WaveformUL-IAB-r16", Optional: true},
					{Name: "dci-25-AI-RNTI-Support-IAB-r16", Optional: true},
					{Name: "t-DeltaReceptionSupport-IAB-r16", Optional: true},
					{Name: "guardSymbolReportReception-IAB-r16", Optional: true},
					{Name: "harqACK-CB-SpatialBundlingPUCCH-Group-r16", Optional: true},
					{Name: "crossSlotScheduling-r16", Optional: true},
					{Name: "maxNumberSRS-PosPathLossEstimateAllServingCells-r16", Optional: true},
					{Name: "extendedCG-Periodicities-r16", Optional: true},
					{Name: "extendedSPS-Periodicities-r16", Optional: true},
					{Name: "codebookVariantsList-r16", Optional: true},
					{Name: "pusch-RepetitionTypeA-r16", Optional: true},
					{Name: "dci-DL-PriorityIndicator-r16", Optional: true},
					{Name: "dci-UL-PriorityIndicator-r16", Optional: true},
					{Name: "maxNumberPathlossRS-Update-r16", Optional: true},
					{Name: "type2-HARQ-ACK-Codebook-r16", Optional: true},
					{Name: "maxTotalResourcesForAcrossFreqRanges-r16", Optional: true},
					{Name: "harqACK-separateMultiDCI-MultiTRP-r16", Optional: true},
					{Name: "harqACK-jointMultiDCI-MultiTRP-r16", Optional: true},
					{Name: "bwp-SwitchingMultiCCs-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TwoStepRACH_r16 != nil, ie.Dci_Format1_2And0_2_r16 != nil, ie.MonitoringDCI_SameSearchSpace_r16 != nil, ie.Type2_CG_ReleaseDCI_0_1_r16 != nil, ie.Type2_CG_ReleaseDCI_0_2_r16 != nil, ie.Sps_ReleaseDCI_1_1_r16 != nil, ie.Sps_ReleaseDCI_1_2_r16 != nil, ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil, ie.SeparateSMTC_InterIAB_Support_r16 != nil, ie.SeparateRACH_IAB_Support_r16 != nil, ie.Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16 != nil, ie.Ul_FlexibleDL_SlotFormatDynamics_IAB_r16 != nil, ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil, ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil, ie.T_DeltaReceptionSupport_IAB_r16 != nil, ie.GuardSymbolReportReception_IAB_r16 != nil, ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil, ie.CrossSlotScheduling_r16 != nil, ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil, ie.ExtendedCG_Periodicities_r16 != nil, ie.ExtendedSPS_Periodicities_r16 != nil, ie.CodebookVariantsList_r16 != nil, ie.Pusch_RepetitionTypeA_r16 != nil, ie.Dci_DL_PriorityIndicator_r16 != nil, ie.Dci_UL_PriorityIndicator_r16 != nil, ie.MaxNumberPathlossRS_Update_r16 != nil, ie.Type2_HARQ_ACK_Codebook_r16 != nil, ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil, ie.HarqACK_SeparateMultiDCI_MultiTRP_r16 != nil, ie.HarqACK_JointMultiDCI_MultiTRP_r16 != nil, ie.Bwp_SwitchingMultiCCs_r16 != nil}); err != nil {
				return err
			}

			if ie.TwoStepRACH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoStepRACH_r16, phyParametersCommonExtTwoStepRACHR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dci_Format1_2And0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dci_Format1_2And0_2_r16, phyParametersCommonExtDciFormat12And02R16Constraints); err != nil {
					return err
				}
			}

			if ie.MonitoringDCI_SameSearchSpace_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MonitoringDCI_SameSearchSpace_r16, phyParametersCommonExtMonitoringDCISameSearchSpaceR16Constraints); err != nil {
					return err
				}
			}

			if ie.Type2_CG_ReleaseDCI_0_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Type2_CG_ReleaseDCI_0_1_r16, phyParametersCommonExtType2CGReleaseDCI01R16Constraints); err != nil {
					return err
				}
			}

			if ie.Type2_CG_ReleaseDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Type2_CG_ReleaseDCI_0_2_r16, phyParametersCommonExtType2CGReleaseDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_ReleaseDCI_1_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Sps_ReleaseDCI_1_1_r16, phyParametersCommonExtSpsReleaseDCI11R16Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_ReleaseDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Sps_ReleaseDCI_1_2_r16, phyParametersCommonExtSpsReleaseDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Csi_TriggerStateNon_ActiveBWP_r16, phyParametersCommonExtCsiTriggerStateNonActiveBWPR16Constraints); err != nil {
					return err
				}
			}

			if ie.SeparateSMTC_InterIAB_Support_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SeparateSMTC_InterIAB_Support_r16, phyParametersCommonExtSeparateSMTCInterIABSupportR16Constraints); err != nil {
					return err
				}
			}

			if ie.SeparateRACH_IAB_Support_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SeparateRACH_IAB_Support_r16, phyParametersCommonExtSeparateRACHIABSupportR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16, phyParametersCommonExtUlFlexibleDLSlotFormatSemiStaticIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_FlexibleDL_SlotFormatDynamics_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_FlexibleDL_SlotFormatDynamics_IAB_r16, phyParametersCommonExtUlFlexibleDLSlotFormatDynamicsIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dft_S_OFDM_WaveformUL_IAB_r16, phyParametersCommonExtDftSOFDMWaveformULIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dci_25_AI_RNTI_Support_IAB_r16, phyParametersCommonExtDci25AIRNTISupportIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.T_DeltaReceptionSupport_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.T_DeltaReceptionSupport_IAB_r16, phyParametersCommonExtTDeltaReceptionSupportIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.GuardSymbolReportReception_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.GuardSymbolReportReception_IAB_r16, phyParametersCommonExtGuardSymbolReportReceptionIABR16Constraints); err != nil {
					return err
				}
			}

			if ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16, phyParametersCommonExtHarqACKCBSpatialBundlingPUCCHGroupR16Constraints); err != nil {
					return err
				}
			}

			if ie.CrossSlotScheduling_r16 != nil {
				c := ie.CrossSlotScheduling_r16
				phyParametersCommonExtCrossSlotSchedulingR16Seq := ex.NewSequenceEncoder(phyParametersCommonExtCrossSlotSchedulingR16Constraints)
				if err := phyParametersCommonExtCrossSlotSchedulingR16Seq.EncodePreamble([]bool{c.Non_SharedSpectrumChAccess_r16 != nil, c.SharedSpectrumChAccess_r16 != nil}); err != nil {
					return err
				}
				if c.Non_SharedSpectrumChAccess_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Non_SharedSpectrumChAccess_r16), phyParametersCommonExtCrossSlotSchedulingR16NonSharedSpectrumChAccessR16Constraints); err != nil {
						return err
					}
				}
				if c.SharedSpectrumChAccess_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SharedSpectrumChAccess_r16), phyParametersCommonExtCrossSlotSchedulingR16SharedSpectrumChAccessR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16, phyParametersCommonExtMaxNumberSRSPosPathLossEstimateAllServingCellsR16Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedCG_Periodicities_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedCG_Periodicities_r16, phyParametersCommonExtExtendedCGPeriodicitiesR16Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedSPS_Periodicities_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedSPS_Periodicities_r16, phyParametersCommonExtExtendedSPSPeriodicitiesR16Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookVariantsList_r16 != nil {
				if err := ie.CodebookVariantsList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pusch_RepetitionTypeA_r16 != nil {
				c := ie.Pusch_RepetitionTypeA_r16
				phyParametersCommonExtPuschRepetitionTypeAR16Seq := ex.NewSequenceEncoder(phyParametersCommonExtPuschRepetitionTypeAR16Constraints)
				if err := phyParametersCommonExtPuschRepetitionTypeAR16Seq.EncodePreamble([]bool{c.SharedSpectrumChAccess_r16 != nil, c.Non_SharedSpectrumChAccess_r16 != nil}); err != nil {
					return err
				}
				if c.SharedSpectrumChAccess_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SharedSpectrumChAccess_r16), phyParametersCommonExtPuschRepetitionTypeAR16SharedSpectrumChAccessR16Constraints); err != nil {
						return err
					}
				}
				if c.Non_SharedSpectrumChAccess_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Non_SharedSpectrumChAccess_r16), phyParametersCommonExtPuschRepetitionTypeAR16NonSharedSpectrumChAccessR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Dci_DL_PriorityIndicator_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dci_DL_PriorityIndicator_r16, phyParametersCommonExtDciDLPriorityIndicatorR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dci_UL_PriorityIndicator_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dci_UL_PriorityIndicator_r16, phyParametersCommonExtDciULPriorityIndicatorR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberPathlossRS_Update_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberPathlossRS_Update_r16, phyParametersCommonExtMaxNumberPathlossRSUpdateR16Constraints); err != nil {
					return err
				}
			}

			if ie.Type2_HARQ_ACK_Codebook_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Type2_HARQ_ACK_Codebook_r16, phyParametersCommonExtType2HARQACKCodebookR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil {
				c := ie.MaxTotalResourcesForAcrossFreqRanges_r16
				phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq := ex.NewSequenceEncoder(phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Constraints)
				if err := phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq.EncodePreamble([]bool{c.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil, c.MaxNumberResAcrossCC_AcrossFR_r16 != nil}); err != nil {
					return err
				}
				if c.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16), phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResWithinSlotAcrossCCAcrossFRR16Constraints); err != nil {
						return err
					}
				}
				if c.MaxNumberResAcrossCC_AcrossFR_r16 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumberResAcrossCC_AcrossFR_r16), phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResAcrossCCAcrossFRR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.HarqACK_SeparateMultiDCI_MultiTRP_r16 != nil {
				c := ie.HarqACK_SeparateMultiDCI_MultiTRP_r16
				phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Seq := ex.NewSequenceEncoder(phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Constraints)
				if err := phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Seq.EncodePreamble([]bool{c.MaxNumberLongPUCCHs_r16 != nil}); err != nil {
					return err
				}
				if c.MaxNumberLongPUCCHs_r16 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumberLongPUCCHs_r16), phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16MaxNumberLongPUCCHsR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.HarqACK_JointMultiDCI_MultiTRP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.HarqACK_JointMultiDCI_MultiTRP_r16, phyParametersCommonExtHarqACKJointMultiDCIMultiTRPR16Constraints); err != nil {
					return err
				}
			}

			if ie.Bwp_SwitchingMultiCCs_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(phyParametersCommonExtBwpSwitchingMultiCCsR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Bwp_SwitchingMultiCCs_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Bwp_SwitchingMultiCCs_r16).Choice {
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiCCs_r16).Type1_r16), phyParametersCommonExtBwpSwitchingMultiCCsR16Type1R16Constraints); err != nil {
						return err
					}
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiCCs_r16).Type2_r16), phyParametersCommonExtBwpSwitchingMultiCCsR16Type2R16Constraints); err != nil {
						return err
					}
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
					{Name: "targetSMTC-SCG-r16", Optional: true},
					{Name: "supportRepetitionZeroOffsetRV-r16", Optional: true},
					{Name: "cbg-TransInOrderPUSCH-UL-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TargetSMTC_SCG_r16 != nil, ie.SupportRepetitionZeroOffsetRV_r16 != nil, ie.Cbg_TransInOrderPUSCH_UL_r16 != nil}); err != nil {
				return err
			}

			if ie.TargetSMTC_SCG_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.TargetSMTC_SCG_r16, phyParametersCommonExtTargetSMTCSCGR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportRepetitionZeroOffsetRV_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportRepetitionZeroOffsetRV_r16, phyParametersCommonExtSupportRepetitionZeroOffsetRVR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cbg_TransInOrderPUSCH_UL_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cbg_TransInOrderPUSCH_UL_r16, phyParametersCommonExtCbgTransInOrderPUSCHULR16Constraints); err != nil {
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
					{Name: "bwp-SwitchingMultiDormancyCCs-r16", Optional: true},
					{Name: "supportRetx-Diff-CoresetPool-Multi-DCI-TRP-r16", Optional: true},
					{Name: "pdcch-MonitoringAnyOccasionsWithSpanGapCrossCarrierSch-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil, ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil, ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil}); err != nil {
				return err
			}

			if ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Bwp_SwitchingMultiDormancyCCs_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Bwp_SwitchingMultiDormancyCCs_r16).Choice {
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type1_r16), phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type1R16Constraints); err != nil {
						return err
					}
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type2_r16), phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type2R16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16, phyParametersCommonExtSupportRetxDiffCoresetPoolMultiDCITRPR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16, phyParametersCommonExtPdcchMonitoringAnyOccasionsWithSpanGapCrossCarrierSchR16Constraints); err != nil {
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
					{Name: "newBeamIdentifications2PortCSI-RS-r16", Optional: true},
					{Name: "pathlossEstimation2PortCSI-RS-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NewBeamIdentifications2PortCSI_RS_r16 != nil, ie.PathlossEstimation2PortCSI_RS_r16 != nil}); err != nil {
				return err
			}

			if ie.NewBeamIdentifications2PortCSI_RS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.NewBeamIdentifications2PortCSI_RS_r16, phyParametersCommonExtNewBeamIdentifications2PortCSIRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossEstimation2PortCSI_RS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossEstimation2PortCSI_RS_r16, phyParametersCommonExtPathlossEstimation2PortCSIRSR16Constraints); err != nil {
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
					{Name: "mux-HARQ-ACK-withoutPUCCH-onPUSCH-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16 != nil}); err != nil {
				return err
			}

			if ie.Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16, phyParametersCommonExtMuxHARQACKWithoutPUCCHOnPUSCHR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 8:
		if hasExtGroup8 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "guardSymbolReportReception-IAB-r17", Optional: true},
					{Name: "restricted-IAB-DU-BeamReception-r17", Optional: true},
					{Name: "recommended-IAB-MT-BeamTransmission-r17", Optional: true},
					{Name: "case6-TimingAlignmentReception-IAB-r17", Optional: true},
					{Name: "case7-TimingAlignmentReception-IAB-r17", Optional: true},
					{Name: "dl-tx-PowerAdjustment-IAB-r17", Optional: true},
					{Name: "desired-ul-tx-PowerAdjustment-r17", Optional: true},
					{Name: "fdm-SoftResourceAvailability-DynamicIndication-r17", Optional: true},
					{Name: "updated-T-DeltaRangeReception-r17", Optional: true},
					{Name: "slotBasedDynamicPUCCH-Rep-r17", Optional: true},
					{Name: "sps-HARQ-ACK-Deferral-r17", Optional: true},
					{Name: "unifiedJointTCI-commonUpdate-r17", Optional: true},
					{Name: "mTRP-PDCCH-singleSpan-r17", Optional: true},
					{Name: "supportedActivatedPRS-ProcessingWindow-r17", Optional: true},
					{Name: "cg-TimeDomainAllocationExtension-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GuardSymbolReportReception_IAB_r17 != nil, ie.Restricted_IAB_DU_BeamReception_r17 != nil, ie.Recommended_IAB_MT_BeamTransmission_r17 != nil, ie.Case6_TimingAlignmentReception_IAB_r17 != nil, ie.Case7_TimingAlignmentReception_IAB_r17 != nil, ie.Dl_Tx_PowerAdjustment_IAB_r17 != nil, ie.Desired_Ul_Tx_PowerAdjustment_r17 != nil, ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil, ie.Updated_T_DeltaRangeReception_r17 != nil, ie.SlotBasedDynamicPUCCH_Rep_r17 != nil, ie.Sps_HARQ_ACK_Deferral_r17 != nil, ie.UnifiedJointTCI_CommonUpdate_r17 != nil, ie.MTRP_PDCCH_SingleSpan_r17 != nil, ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil, ie.Cg_TimeDomainAllocationExtension_r17 != nil}); err != nil {
				return err
			}

			if ie.GuardSymbolReportReception_IAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.GuardSymbolReportReception_IAB_r17, phyParametersCommonExtGuardSymbolReportReceptionIABR17Constraints); err != nil {
					return err
				}
			}

			if ie.Restricted_IAB_DU_BeamReception_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Restricted_IAB_DU_BeamReception_r17, phyParametersCommonExtRestrictedIABDUBeamReceptionR17Constraints); err != nil {
					return err
				}
			}

			if ie.Recommended_IAB_MT_BeamTransmission_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Recommended_IAB_MT_BeamTransmission_r17, phyParametersCommonExtRecommendedIABMTBeamTransmissionR17Constraints); err != nil {
					return err
				}
			}

			if ie.Case6_TimingAlignmentReception_IAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Case6_TimingAlignmentReception_IAB_r17, phyParametersCommonExtCase6TimingAlignmentReceptionIABR17Constraints); err != nil {
					return err
				}
			}

			if ie.Case7_TimingAlignmentReception_IAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Case7_TimingAlignmentReception_IAB_r17, phyParametersCommonExtCase7TimingAlignmentReceptionIABR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_Tx_PowerAdjustment_IAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_Tx_PowerAdjustment_IAB_r17, phyParametersCommonExtDlTxPowerAdjustmentIABR17Constraints); err != nil {
					return err
				}
			}

			if ie.Desired_Ul_Tx_PowerAdjustment_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Desired_Ul_Tx_PowerAdjustment_r17, phyParametersCommonExtDesiredUlTxPowerAdjustmentR17Constraints); err != nil {
					return err
				}
			}

			if ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Fdm_SoftResourceAvailability_DynamicIndication_r17, phyParametersCommonExtFdmSoftResourceAvailabilityDynamicIndicationR17Constraints); err != nil {
					return err
				}
			}

			if ie.Updated_T_DeltaRangeReception_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Updated_T_DeltaRangeReception_r17, phyParametersCommonExtUpdatedTDeltaRangeReceptionR17Constraints); err != nil {
					return err
				}
			}

			if ie.SlotBasedDynamicPUCCH_Rep_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SlotBasedDynamicPUCCH_Rep_r17, phyParametersCommonExtSlotBasedDynamicPUCCHRepR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_HARQ_ACK_Deferral_r17 != nil {
				c := ie.Sps_HARQ_ACK_Deferral_r17
				phyParametersCommonExtSpsHARQACKDeferralR17Seq := ex.NewSequenceEncoder(phyParametersCommonExtSpsHARQACKDeferralR17Constraints)
				if err := phyParametersCommonExtSpsHARQACKDeferralR17Seq.EncodePreamble([]bool{c.Non_SharedSpectrumChAccess_r17 != nil, c.SharedSpectrumChAccess_r17 != nil}); err != nil {
					return err
				}
				if c.Non_SharedSpectrumChAccess_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Non_SharedSpectrumChAccess_r17), phyParametersCommonExtSpsHARQACKDeferralR17NonSharedSpectrumChAccessR17Constraints); err != nil {
						return err
					}
				}
				if c.SharedSpectrumChAccess_r17 != nil {
					if err := ex.EncodeEnumerated((*c.SharedSpectrumChAccess_r17), phyParametersCommonExtSpsHARQACKDeferralR17SharedSpectrumChAccessR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.UnifiedJointTCI_CommonUpdate_r17 != nil {
				if err := ex.EncodeInteger(*ie.UnifiedJointTCI_CommonUpdate_r17, phyParametersCommonUnifiedJointTCICommonUpdateR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PDCCH_SingleSpan_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PDCCH_SingleSpan_r17, phyParametersCommonExtMTRPPDCCHSingleSpanR17Constraints); err != nil {
					return err
				}
			}

			if ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportedActivatedPRS_ProcessingWindow_r17, phyParametersCommonExtSupportedActivatedPRSProcessingWindowR17Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_TimeDomainAllocationExtension_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_TimeDomainAllocationExtension_r17, phyParametersCommonExtCgTimeDomainAllocationExtensionR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 9:
		if hasExtGroup9 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ta-BasedPDC-TN-NonSharedSpectrumChAccess-r17", Optional: true},
					{Name: "directionalCollisionDC-IAB-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil, ie.DirectionalCollisionDC_IAB_r17 != nil}); err != nil {
				return err
			}

			if ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17, phyParametersCommonExtTaBasedPDCTNNonSharedSpectrumChAccessR17Constraints); err != nil {
					return err
				}
			}

			if ie.DirectionalCollisionDC_IAB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DirectionalCollisionDC_IAB_r17, phyParametersCommonExtDirectionalCollisionDCIABR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 10:
		if hasExtGroup10 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dummy1", Optional: true},
					{Name: "dummy2", Optional: true},
					{Name: "dummy3", Optional: true},
					{Name: "dummy4", Optional: true},
					{Name: "srs-AdditionalRepetition-r17", Optional: true},
					{Name: "pusch-Repetition-CG-SDT-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy1 != nil, ie.Dummy2 != nil, ie.Dummy3 != nil, ie.Dummy4 != nil, ie.Srs_AdditionalRepetition_r17 != nil, ie.Pusch_Repetition_CG_SDT_r17 != nil}); err != nil {
				return err
			}

			if ie.Dummy1 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy1, phyParametersCommonExtDummy1Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy2 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy2, phyParametersCommonExtDummy2Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy3 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy3, phyParametersCommonExtDummy3Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy4 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy4, phyParametersCommonExtDummy4Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_AdditionalRepetition_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_AdditionalRepetition_r17, phyParametersCommonExtSrsAdditionalRepetitionR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_Repetition_CG_SDT_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_Repetition_CG_SDT_r17, phyParametersCommonExtPuschRepetitionCGSDTR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 11:
		if hasExtGroup11 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "multiPDSCH-PerSlotType1-CB-Support-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultiPDSCH_PerSlotType1_CB_Support_r17 != nil}); err != nil {
				return err
			}

			if ie.MultiPDSCH_PerSlotType1_CB_Support_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPDSCH_PerSlotType1_CB_Support_r17, phyParametersCommonExtMultiPDSCHPerSlotType1CBSupportR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 12:
		if hasExtGroup12 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "jointPowerSpatialAdaptation-r18", Optional: true},
					{Name: "ncr-AperiodicBeamInd-AccessLink-r18", Optional: true},
					{Name: "ncr-Semi-PersistentBeamInd-AccessLink-r18", Optional: true},
					{Name: "ncr-SimultaneousUL-BackhaulAndC-Link-r18", Optional: true},
					{Name: "ncr-BackhaulBeamInd-r18", Optional: true},
					{Name: "ncr-AdaptiveBeamBackhaulAndC-Link-r18", Optional: true},
					{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-1-3-r18", Optional: true},
					{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-0-3-r18", Optional: true},
					{Name: "configurableType-1A-FieldsForDCI-0-3-And-1-3-r18", Optional: true},
					{Name: "fdra-Type-1-Gty-2-4-8-16-RBs-RIV-DCI-1-3-And-0-3-r18", Optional: true},
					{Name: "priorityIndicationDL-r18", Optional: true},
					{Name: "priorityIndicationUL-r18", Optional: true},
					{Name: "dynamicIndicationSchedulingRestriction-r18", Optional: true},
					{Name: "priorityIndicationOneSlotHARQ-r18", Optional: true},
					{Name: "multiPUSCH-DCI-0-1-r18", Optional: true},
					{Name: "multiPUSCH-DCI-0-2-r18", Optional: true},
					{Name: "additionalSR-Periodicities-r18", Optional: true},
					{Name: "pathlossRS-UpdateForType1CG-PUSCH-r18", Optional: true},
					{Name: "bwp-SwitchingMultiDormancyCC-DCI-0-3-And-1-3-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.JointPowerSpatialAdaptation_r18 != nil, ie.Ncr_AperiodicBeamInd_AccessLink_r18 != nil, ie.Ncr_Semi_PersistentBeamInd_AccessLink_r18 != nil, ie.Ncr_SimultaneousUL_BackhaulAndC_Link_r18 != nil, ie.Ncr_BackhaulBeamInd_r18 != nil, ie.Ncr_AdaptiveBeamBackhaulAndC_Link_r18 != nil, ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18 != nil, ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18 != nil, ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18 != nil, ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18 != nil, ie.PriorityIndicationDL_r18 != nil, ie.PriorityIndicationUL_r18 != nil, ie.DynamicIndicationSchedulingRestriction_r18 != nil, ie.PriorityIndicationOneSlotHARQ_r18 != nil, ie.MultiPUSCH_DCI_0_1_r18 != nil, ie.MultiPUSCH_DCI_0_2_r18 != nil, ie.AdditionalSR_Periodicities_r18 != nil, ie.PathlossRS_UpdateForType1CG_PUSCH_r18 != nil, ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18 != nil}); err != nil {
				return err
			}

			if ie.JointPowerSpatialAdaptation_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.JointPowerSpatialAdaptation_r18, phyParametersCommonExtJointPowerSpatialAdaptationR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ncr_AperiodicBeamInd_AccessLink_r18 != nil {
				c := ie.Ncr_AperiodicBeamInd_AccessLink_r18
				phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq := ex.NewSequenceEncoder(phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Constraints)
				if err := phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r18 != nil {
					if err := ex.EncodeInteger((*c.Scs_15kHz_r18), per.Constrained(0, 1)); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r18 != nil {
					if err := ex.EncodeInteger((*c.Scs_30kHz_r18), per.Constrained(0, 1)); err != nil {
						return err
					}
				}
				if c.Scs_60kHz_r18 != nil {
					if err := ex.EncodeInteger((*c.Scs_60kHz_r18), per.Constrained(0, 2)); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r18 != nil {
					if err := ex.EncodeInteger((*c.Scs_120kHz_r18), per.Constrained(0, 2)); err != nil {
						return err
					}
				}
			}

			if ie.Ncr_Semi_PersistentBeamInd_AccessLink_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_Semi_PersistentBeamInd_AccessLink_r18, phyParametersCommonExtNcrSemiPersistentBeamIndAccessLinkR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ncr_SimultaneousUL_BackhaulAndC_Link_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_SimultaneousUL_BackhaulAndC_Link_r18, phyParametersCommonExtNcrSimultaneousULBackhaulAndCLinkR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ncr_BackhaulBeamInd_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_BackhaulBeamInd_r18, phyParametersCommonExtNcrBackhaulBeamIndR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ncr_AdaptiveBeamBackhaulAndC_Link_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_AdaptiveBeamBackhaulAndC_Link_r18, phyParametersCommonExtNcrAdaptiveBeamBackhaulAndCLinkR18Constraints); err != nil {
					return err
				}
			}

			if ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18, phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13R18Constraints); err != nil {
					return err
				}
			}

			if ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18, phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03R18Constraints); err != nil {
					return err
				}
			}

			if ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18, phyParametersCommonExtConfigurableType1AFieldsForDCI03And13R18Constraints); err != nil {
					return err
				}
			}

			if ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18, phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03R18Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationDL_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationDL_r18, phyParametersCommonExtPriorityIndicationDLR18Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationUL_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationUL_r18, phyParametersCommonExtPriorityIndicationULR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicIndicationSchedulingRestriction_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicIndicationSchedulingRestriction_r18, phyParametersCommonExtDynamicIndicationSchedulingRestrictionR18Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationOneSlotHARQ_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationOneSlotHARQ_r18, phyParametersCommonExtPriorityIndicationOneSlotHARQR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUSCH_DCI_0_1_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUSCH_DCI_0_1_r18, phyParametersCommonExtMultiPUSCHDCI01R18Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUSCH_DCI_0_2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUSCH_DCI_0_2_r18, phyParametersCommonExtMultiPUSCHDCI02R18Constraints); err != nil {
					return err
				}
			}

			if ie.AdditionalSR_Periodicities_r18 != nil {
				c := ie.AdditionalSR_Periodicities_r18
				phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq := ex.NewSequenceEncoder(phyParametersCommonExtAdditionalSRPeriodicitiesR18Constraints)
				if err := phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq.EncodePreamble([]bool{c.Scs_30kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_30kHz_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_30kHz_r18), phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs30kHzR18Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r18 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_120kHz_r18), phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs120kHzR18Constraints); err != nil {
						return err
					}
				}
			}

			if ie.PathlossRS_UpdateForType1CG_PUSCH_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossRS_UpdateForType1CG_PUSCH_r18, phyParametersCommonExtPathlossRSUpdateForType1CGPUSCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Choice {
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type1_r18), phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type1R18Constraints); err != nil {
						return err
					}
				case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18:
					if err := ex.EncodeEnumerated((*(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type2_r18), phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type2R18Constraints); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 13:
		if hasExtGroup13 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ncr-dft-S-OFDM-WaveformUL-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ncr_Dft_S_OFDM_WaveformUL_r18 != nil}); err != nil {
				return err
			}

			if ie.Ncr_Dft_S_OFDM_WaveformUL_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_Dft_S_OFDM_WaveformUL_r18, phyParametersCommonExtNcrDftSOFDMWaveformULR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 14:
		if hasExtGroup14 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "codebookVariantsListExt-r19", Optional: true},
					{Name: "codebookVariantsListAggregate-r19", Optional: true},
					{Name: "codebookVariantsListHybrid-r19", Optional: true},
					{Name: "dummy-EnableTx-RxDuringMeasGap-r19", Optional: true},
					{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-1-3-Diff-r19", Optional: true},
					{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-0-3-Diff-r19", Optional: true},
					{Name: "configurableType-1A-FieldsForDCI-0-3-And-1-3-Diff-r19", Optional: true},
					{Name: "fdra-Type-1-Gty-2-4-8-16-RBs-RIV-DCI-1-3-And-0-3-Diff-r19", Optional: true},
					{Name: "priorityIndicationDL-Diff-r19", Optional: true},
					{Name: "priorityIndicationUL-Diff-r19", Optional: true},
					{Name: "dynamicIndicationSchedulingRestrictionDiff-r19", Optional: true},
					{Name: "priorityIndicationOneSlotHARQ-Diff-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CodebookVariantsListExt_r19 != nil, ie.CodebookVariantsListAggregate_r19 != nil, ie.CodebookVariantsListHybrid_r19 != nil, ie.Dummy_EnableTx_RxDuringMeasGap_r19 != nil, ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19 != nil, ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19 != nil, ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19 != nil, ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19 != nil, ie.PriorityIndicationDL_Diff_r19 != nil, ie.PriorityIndicationUL_Diff_r19 != nil, ie.DynamicIndicationSchedulingRestrictionDiff_r19 != nil, ie.PriorityIndicationOneSlotHARQ_Diff_r19 != nil}); err != nil {
				return err
			}

			if ie.CodebookVariantsListExt_r19 != nil {
				if err := ie.CodebookVariantsListExt_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookVariantsListAggregate_r19 != nil {
				if err := ie.CodebookVariantsListAggregate_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookVariantsListHybrid_r19 != nil {
				if err := ie.CodebookVariantsListHybrid_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Dummy_EnableTx_RxDuringMeasGap_r19 != nil {
				c := ie.Dummy_EnableTx_RxDuringMeasGap_r19
				if err := ex.EncodeBitString(c.AdditionalDCI_r19, per.FixedSize(3)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.IndicationField_r19, phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MinimumTimeOffset_r19, phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints); err != nil {
					return err
				}
			}

			if ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19, phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13DiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19, phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03DiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19, phyParametersCommonExtConfigurableType1AFieldsForDCI03And13DiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19, phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03DiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationDL_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationDL_Diff_r19, phyParametersCommonExtPriorityIndicationDLDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationUL_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationUL_Diff_r19, phyParametersCommonExtPriorityIndicationULDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicIndicationSchedulingRestrictionDiff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicIndicationSchedulingRestrictionDiff_r19, phyParametersCommonExtDynamicIndicationSchedulingRestrictionDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicationOneSlotHARQ_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicationOneSlotHARQ_Diff_r19, phyParametersCommonExtPriorityIndicationOneSlotHARQDiffR19Constraints); err != nil {
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

func (ie *Phy_ParametersCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersCommonConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. csi-RS-CFRA-ForHO: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersCommonCsiRSCFRAForHOConstraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_CFRA_ForHO = &idx
		}
	}

	// 4. dynamicPRB-BundlingDL: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicPRBBundlingDLConstraints)
			if err != nil {
				return err
			}
			ie.DynamicPRB_BundlingDL = &idx
		}
	}

	// 5. sp-CSI-ReportPUCCH: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(phyParametersCommonSpCSIReportPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_ReportPUCCH = &idx
		}
	}

	// 6. sp-CSI-ReportPUSCH: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(phyParametersCommonSpCSIReportPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_ReportPUSCH = &idx
		}
	}

	// 7. nzp-CSI-RS-IntefMgmt: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(phyParametersCommonNzpCSIRSIntefMgmtConstraints)
			if err != nil {
				return err
			}
			ie.Nzp_CSI_RS_IntefMgmt = &idx
		}
	}

	// 8. type2-SP-CSI-Feedback-LongPUCCH: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(phyParametersCommonType2SPCSIFeedbackLongPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.Type2_SP_CSI_Feedback_LongPUCCH = &idx
		}
	}

	// 9. precoderGranularityCORESET: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPrecoderGranularityCORESETConstraints)
			if err != nil {
				return err
			}
			ie.PrecoderGranularityCORESET = &idx
		}
	}

	// 10. dynamicHARQ-ACK-Codebook: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicHARQACKCodebookConstraints)
			if err != nil {
				return err
			}
			ie.DynamicHARQ_ACK_Codebook = &idx
		}
	}

	// 11. semiStaticHARQ-ACK-Codebook: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(phyParametersCommonSemiStaticHARQACKCodebookConstraints)
			if err != nil {
				return err
			}
			ie.SemiStaticHARQ_ACK_Codebook = &idx
		}
	}

	// 12. spatialBundlingHARQ-ACK: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(phyParametersCommonSpatialBundlingHARQACKConstraints)
			if err != nil {
				return err
			}
			ie.SpatialBundlingHARQ_ACK = &idx
		}
	}

	// 13. dynamicBetaOffsetInd-HARQ-ACK-CSI: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicBetaOffsetIndHARQACKCSIConstraints)
			if err != nil {
				return err
			}
			ie.DynamicBetaOffsetInd_HARQ_ACK_CSI = &idx
		}
	}

	// 14. pucch-Repetition-F1-3-4: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPucchRepetitionF134Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_Repetition_F1_3_4 = &idx
		}
	}

	// 15. ra-Type0-PUSCH: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(phyParametersCommonRaType0PUSCHConstraints)
			if err != nil {
				return err
			}
			ie.Ra_Type0_PUSCH = &idx
		}
	}

	// 16. dynamicSwitchRA-Type0-1-PDSCH: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicSwitchRAType01PDSCHConstraints)
			if err != nil {
				return err
			}
			ie.DynamicSwitchRA_Type0_1_PDSCH = &idx
		}
	}

	// 17. dynamicSwitchRA-Type0-1-PUSCH: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicSwitchRAType01PUSCHConstraints)
			if err != nil {
				return err
			}
			ie.DynamicSwitchRA_Type0_1_PUSCH = &idx
		}
	}

	// 18. pdsch-MappingTypeA: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPdschMappingTypeAConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_MappingTypeA = &idx
		}
	}

	// 19. pdsch-MappingTypeB: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPdschMappingTypeBConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_MappingTypeB = &idx
		}
	}

	// 20. interleavingVRB-ToPRB-PDSCH: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(phyParametersCommonInterleavingVRBToPRBPDSCHConstraints)
			if err != nil {
				return err
			}
			ie.InterleavingVRB_ToPRB_PDSCH = &idx
		}
	}

	// 21. interSlotFreqHopping-PUSCH: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(phyParametersCommonInterSlotFreqHoppingPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.InterSlotFreqHopping_PUSCH = &idx
		}
	}

	// 22. type1-PUSCH-RepetitionMultiSlots: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(phyParametersCommonType1PUSCHRepetitionMultiSlotsConstraints)
			if err != nil {
				return err
			}
			ie.Type1_PUSCH_RepetitionMultiSlots = &idx
		}
	}

	// 23. type2-PUSCH-RepetitionMultiSlots: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(phyParametersCommonType2PUSCHRepetitionMultiSlotsConstraints)
			if err != nil {
				return err
			}
			ie.Type2_PUSCH_RepetitionMultiSlots = &idx
		}
	}

	// 24. pusch-RepetitionMultiSlots: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPuschRepetitionMultiSlotsConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepetitionMultiSlots = &idx
		}
	}

	// 25. pdsch-RepetitionMultiSlots: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPdschRepetitionMultiSlotsConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RepetitionMultiSlots = &idx
		}
	}

	// 26. downlinkSPS: enumerated
	{
		if seq.IsComponentPresent(23) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDownlinkSPSConstraints)
			if err != nil {
				return err
			}
			ie.DownlinkSPS = &idx
		}
	}

	// 27. configuredUL-GrantType1: enumerated
	{
		if seq.IsComponentPresent(24) {
			idx, err := d.DecodeEnumerated(phyParametersCommonConfiguredULGrantType1Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType1 = &idx
		}
	}

	// 28. configuredUL-GrantType2: enumerated
	{
		if seq.IsComponentPresent(25) {
			idx, err := d.DecodeEnumerated(phyParametersCommonConfiguredULGrantType2Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType2 = &idx
		}
	}

	// 29. pre-EmptIndication-DL: enumerated
	{
		if seq.IsComponentPresent(26) {
			idx, err := d.DecodeEnumerated(phyParametersCommonPreEmptIndicationDLConstraints)
			if err != nil {
				return err
			}
			ie.Pre_EmptIndication_DL = &idx
		}
	}

	// 30. cbg-TransIndication-DL: enumerated
	{
		if seq.IsComponentPresent(27) {
			idx, err := d.DecodeEnumerated(phyParametersCommonCbgTransIndicationDLConstraints)
			if err != nil {
				return err
			}
			ie.Cbg_TransIndication_DL = &idx
		}
	}

	// 31. cbg-TransIndication-UL: enumerated
	{
		if seq.IsComponentPresent(28) {
			idx, err := d.DecodeEnumerated(phyParametersCommonCbgTransIndicationULConstraints)
			if err != nil {
				return err
			}
			ie.Cbg_TransIndication_UL = &idx
		}
	}

	// 32. cbg-FlushIndication-DL: enumerated
	{
		if seq.IsComponentPresent(29) {
			idx, err := d.DecodeEnumerated(phyParametersCommonCbgFlushIndicationDLConstraints)
			if err != nil {
				return err
			}
			ie.Cbg_FlushIndication_DL = &idx
		}
	}

	// 33. dynamicHARQ-ACK-CodeB-CBG-Retx-DL: enumerated
	{
		if seq.IsComponentPresent(30) {
			idx, err := d.DecodeEnumerated(phyParametersCommonDynamicHARQACKCodeBCBGRetxDLConstraints)
			if err != nil {
				return err
			}
			ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL = &idx
		}
	}

	// 34. rateMatchingResrcSetSemi-Static: enumerated
	{
		if seq.IsComponentPresent(31) {
			idx, err := d.DecodeEnumerated(phyParametersCommonRateMatchingResrcSetSemiStaticConstraints)
			if err != nil {
				return err
			}
			ie.RateMatchingResrcSetSemi_Static = &idx
		}
	}

	// 35. rateMatchingResrcSetDynamic: enumerated
	{
		if seq.IsComponentPresent(32) {
			idx, err := d.DecodeEnumerated(phyParametersCommonRateMatchingResrcSetDynamicConstraints)
			if err != nil {
				return err
			}
			ie.RateMatchingResrcSetDynamic = &idx
		}
	}

	// 36. bwp-SwitchingDelay: enumerated
	{
		if seq.IsComponentPresent(33) {
			idx, err := d.DecodeEnumerated(phyParametersCommonBwpSwitchingDelayConstraints)
			if err != nil {
				return err
			}
			ie.Bwp_SwitchingDelay = &idx
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
				{Name: "dummy", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxNumberSearchSpaces", Optional: true},
				{Name: "rateMatchingCtrlResrcSetDynamic", Optional: true},
				{Name: "maxLayersMIMO-Indication", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxNumberSearchSpacesConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberSearchSpaces = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtRateMatchingCtrlResrcSetDynamicConstraints)
			if err != nil {
				return err
			}
			ie.RateMatchingCtrlResrcSetDynamic = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxLayersMIMOIndicationConstraints)
			if err != nil {
				return err
			}
			ie.MaxLayersMIMO_Indication = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spCellPlacement", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SpCellPlacement = new(CarrierAggregationVariant)
			if err := ie.SpCellPlacement.Decode(dx); err != nil {
				return err
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
				{Name: "twoStepRACH-r16", Optional: true},
				{Name: "dci-Format1-2And0-2-r16", Optional: true},
				{Name: "monitoringDCI-SameSearchSpace-r16", Optional: true},
				{Name: "type2-CG-ReleaseDCI-0-1-r16", Optional: true},
				{Name: "type2-CG-ReleaseDCI-0-2-r16", Optional: true},
				{Name: "sps-ReleaseDCI-1-1-r16", Optional: true},
				{Name: "sps-ReleaseDCI-1-2-r16", Optional: true},
				{Name: "csi-TriggerStateNon-ActiveBWP-r16", Optional: true},
				{Name: "separateSMTC-InterIAB-Support-r16", Optional: true},
				{Name: "separateRACH-IAB-Support-r16", Optional: true},
				{Name: "ul-flexibleDL-SlotFormatSemiStatic-IAB-r16", Optional: true},
				{Name: "ul-flexibleDL-SlotFormatDynamics-IAB-r16", Optional: true},
				{Name: "dft-S-OFDM-WaveformUL-IAB-r16", Optional: true},
				{Name: "dci-25-AI-RNTI-Support-IAB-r16", Optional: true},
				{Name: "t-DeltaReceptionSupport-IAB-r16", Optional: true},
				{Name: "guardSymbolReportReception-IAB-r16", Optional: true},
				{Name: "harqACK-CB-SpatialBundlingPUCCH-Group-r16", Optional: true},
				{Name: "crossSlotScheduling-r16", Optional: true},
				{Name: "maxNumberSRS-PosPathLossEstimateAllServingCells-r16", Optional: true},
				{Name: "extendedCG-Periodicities-r16", Optional: true},
				{Name: "extendedSPS-Periodicities-r16", Optional: true},
				{Name: "codebookVariantsList-r16", Optional: true},
				{Name: "pusch-RepetitionTypeA-r16", Optional: true},
				{Name: "dci-DL-PriorityIndicator-r16", Optional: true},
				{Name: "dci-UL-PriorityIndicator-r16", Optional: true},
				{Name: "maxNumberPathlossRS-Update-r16", Optional: true},
				{Name: "type2-HARQ-ACK-Codebook-r16", Optional: true},
				{Name: "maxTotalResourcesForAcrossFreqRanges-r16", Optional: true},
				{Name: "harqACK-separateMultiDCI-MultiTRP-r16", Optional: true},
				{Name: "harqACK-jointMultiDCI-MultiTRP-r16", Optional: true},
				{Name: "bwp-SwitchingMultiCCs-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtTwoStepRACHR16Constraints)
			if err != nil {
				return err
			}
			ie.TwoStepRACH_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDciFormat12And02R16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_Format1_2And0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMonitoringDCISameSearchSpaceR16Constraints)
			if err != nil {
				return err
			}
			ie.MonitoringDCI_SameSearchSpace_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtType2CGReleaseDCI01R16Constraints)
			if err != nil {
				return err
			}
			ie.Type2_CG_ReleaseDCI_0_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtType2CGReleaseDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Type2_CG_ReleaseDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSpsReleaseDCI11R16Constraints)
			if err != nil {
				return err
			}
			ie.Sps_ReleaseDCI_1_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSpsReleaseDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Sps_ReleaseDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtCsiTriggerStateNonActiveBWPR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_TriggerStateNon_ActiveBWP_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSeparateSMTCInterIABSupportR16Constraints)
			if err != nil {
				return err
			}
			ie.SeparateSMTC_InterIAB_Support_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSeparateRACHIABSupportR16Constraints)
			if err != nil {
				return err
			}
			ie.SeparateRACH_IAB_Support_r16 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtUlFlexibleDLSlotFormatSemiStaticIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FlexibleDL_SlotFormatSemiStatic_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtUlFlexibleDLSlotFormatDynamicsIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FlexibleDL_SlotFormatDynamics_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDftSOFDMWaveformULIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Dft_S_OFDM_WaveformUL_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDci25AIRNTISupportIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_25_AI_RNTI_Support_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtTDeltaReceptionSupportIABR16Constraints)
			if err != nil {
				return err
			}
			ie.T_DeltaReceptionSupport_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtGuardSymbolReportReceptionIABR16Constraints)
			if err != nil {
				return err
			}
			ie.GuardSymbolReportReception_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtHarqACKCBSpatialBundlingPUCCHGroupR16Constraints)
			if err != nil {
				return err
			}
			ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			ie.CrossSlotScheduling_r16 = &struct {
				Non_SharedSpectrumChAccess_r16 *int64
				SharedSpectrumChAccess_r16     *int64
			}{}
			c := ie.CrossSlotScheduling_r16
			phyParametersCommonExtCrossSlotSchedulingR16Seq := dx.NewSequenceDecoder(phyParametersCommonExtCrossSlotSchedulingR16Constraints)
			if err := phyParametersCommonExtCrossSlotSchedulingR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtCrossSlotSchedulingR16Seq.IsComponentPresent(0) {
				c.Non_SharedSpectrumChAccess_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtCrossSlotSchedulingR16NonSharedSpectrumChAccessR16Constraints)
				if err != nil {
					return err
				}
				(*c.Non_SharedSpectrumChAccess_r16) = v
			}
			if phyParametersCommonExtCrossSlotSchedulingR16Seq.IsComponentPresent(1) {
				c.SharedSpectrumChAccess_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtCrossSlotSchedulingR16SharedSpectrumChAccessR16Constraints)
				if err != nil {
					return err
				}
				(*c.SharedSpectrumChAccess_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxNumberSRSPosPathLossEstimateAllServingCellsR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtExtendedCGPeriodicitiesR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedCG_Periodicities_r16 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtExtendedSPSPeriodicitiesR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedSPS_Periodicities_r16 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			ie.CodebookVariantsList_r16 = new(CodebookVariantsList_r16)
			if err := ie.CodebookVariantsList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(22) {
			ie.Pusch_RepetitionTypeA_r16 = &struct {
				SharedSpectrumChAccess_r16     *int64
				Non_SharedSpectrumChAccess_r16 *int64
			}{}
			c := ie.Pusch_RepetitionTypeA_r16
			phyParametersCommonExtPuschRepetitionTypeAR16Seq := dx.NewSequenceDecoder(phyParametersCommonExtPuschRepetitionTypeAR16Constraints)
			if err := phyParametersCommonExtPuschRepetitionTypeAR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtPuschRepetitionTypeAR16Seq.IsComponentPresent(0) {
				c.SharedSpectrumChAccess_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtPuschRepetitionTypeAR16SharedSpectrumChAccessR16Constraints)
				if err != nil {
					return err
				}
				(*c.SharedSpectrumChAccess_r16) = v
			}
			if phyParametersCommonExtPuschRepetitionTypeAR16Seq.IsComponentPresent(1) {
				c.Non_SharedSpectrumChAccess_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtPuschRepetitionTypeAR16NonSharedSpectrumChAccessR16Constraints)
				if err != nil {
					return err
				}
				(*c.Non_SharedSpectrumChAccess_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDciDLPriorityIndicatorR16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_DL_PriorityIndicator_r16 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDciULPriorityIndicatorR16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_UL_PriorityIndicator_r16 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxNumberPathlossRSUpdateR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberPathlossRS_Update_r16 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtType2HARQACKCodebookR16Constraints)
			if err != nil {
				return err
			}
			ie.Type2_HARQ_ACK_Codebook_r16 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			ie.MaxTotalResourcesForAcrossFreqRanges_r16 = &struct {
				MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 *int64
				MaxNumberResAcrossCC_AcrossFR_r16           *int64
			}{}
			c := ie.MaxTotalResourcesForAcrossFreqRanges_r16
			phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq := dx.NewSequenceDecoder(phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Constraints)
			if err := phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq.IsComponentPresent(0) {
				c.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResWithinSlotAcrossCCAcrossFRR16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberResWithinSlotAcrossCC_AcrossFR_r16) = v
			}
			if phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16Seq.IsComponentPresent(1) {
				c.MaxNumberResAcrossCC_AcrossFR_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtMaxTotalResourcesForAcrossFreqRangesR16MaxNumberResAcrossCCAcrossFRR16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberResAcrossCC_AcrossFR_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(28) {
			ie.HarqACK_SeparateMultiDCI_MultiTRP_r16 = &struct{ MaxNumberLongPUCCHs_r16 *int64 }{}
			c := ie.HarqACK_SeparateMultiDCI_MultiTRP_r16
			phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Seq := dx.NewSequenceDecoder(phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Constraints)
			if err := phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16Seq.IsComponentPresent(0) {
				c.MaxNumberLongPUCCHs_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtHarqACKSeparateMultiDCIMultiTRPR16MaxNumberLongPUCCHsR16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberLongPUCCHs_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtHarqACKJointMultiDCIMultiTRPR16Constraints)
			if err != nil {
				return err
			}
			ie.HarqACK_JointMultiDCI_MultiTRP_r16 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			ie.Bwp_SwitchingMultiCCs_r16 = &struct {
				Choice    int
				Type1_r16 *int64
				Type2_r16 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(phyParametersCommonExtBwpSwitchingMultiCCsR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bwp_SwitchingMultiCCs_r16).Choice = int(index)
			switch index {
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type1_r16:
				(*ie.Bwp_SwitchingMultiCCs_r16).Type1_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiCCsR16Type1R16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiCCs_r16).Type1_r16) = v
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiCCs_r16_Type2_r16:
				(*ie.Bwp_SwitchingMultiCCs_r16).Type2_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiCCsR16Type2R16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiCCs_r16).Type2_r16) = v
			}
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "targetSMTC-SCG-r16", Optional: true},
				{Name: "supportRepetitionZeroOffsetRV-r16", Optional: true},
				{Name: "cbg-TransInOrderPUSCH-UL-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtTargetSMTCSCGR16Constraints)
			if err != nil {
				return err
			}
			ie.TargetSMTC_SCG_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSupportRepetitionZeroOffsetRVR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportRepetitionZeroOffsetRV_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtCbgTransInOrderPUSCHULR16Constraints)
			if err != nil {
				return err
			}
			ie.Cbg_TransInOrderPUSCH_UL_r16 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "bwp-SwitchingMultiDormancyCCs-r16", Optional: true},
				{Name: "supportRetx-Diff-CoresetPool-Multi-DCI-TRP-r16", Optional: true},
				{Name: "pdcch-MonitoringAnyOccasionsWithSpanGapCrossCarrierSch-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Bwp_SwitchingMultiDormancyCCs_r16 = &struct {
				Choice    int
				Type1_r16 *int64
				Type2_r16 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Choice = int(index)
			switch index {
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type1_r16:
				(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type1_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type1R16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type1_r16) = v
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCCs_r16_Type2_r16:
				(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type2_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiDormancyCCsR16Type2R16Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiDormancyCCs_r16).Type2_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSupportRetxDiffCoresetPoolMultiDCITRPR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPdcchMonitoringAnyOccasionsWithSpanGapCrossCarrierSchR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "newBeamIdentifications2PortCSI-RS-r16", Optional: true},
				{Name: "pathlossEstimation2PortCSI-RS-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNewBeamIdentifications2PortCSIRSR16Constraints)
			if err != nil {
				return err
			}
			ie.NewBeamIdentifications2PortCSI_RS_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPathlossEstimation2PortCSIRSR16Constraints)
			if err != nil {
				return err
			}
			ie.PathlossEstimation2PortCSI_RS_r16 = &v
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "mux-HARQ-ACK-withoutPUCCH-onPUSCH-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMuxHARQACKWithoutPUCCHOnPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_HARQ_ACK_WithoutPUCCH_OnPUSCH_r16 = &v
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "guardSymbolReportReception-IAB-r17", Optional: true},
				{Name: "restricted-IAB-DU-BeamReception-r17", Optional: true},
				{Name: "recommended-IAB-MT-BeamTransmission-r17", Optional: true},
				{Name: "case6-TimingAlignmentReception-IAB-r17", Optional: true},
				{Name: "case7-TimingAlignmentReception-IAB-r17", Optional: true},
				{Name: "dl-tx-PowerAdjustment-IAB-r17", Optional: true},
				{Name: "desired-ul-tx-PowerAdjustment-r17", Optional: true},
				{Name: "fdm-SoftResourceAvailability-DynamicIndication-r17", Optional: true},
				{Name: "updated-T-DeltaRangeReception-r17", Optional: true},
				{Name: "slotBasedDynamicPUCCH-Rep-r17", Optional: true},
				{Name: "sps-HARQ-ACK-Deferral-r17", Optional: true},
				{Name: "unifiedJointTCI-commonUpdate-r17", Optional: true},
				{Name: "mTRP-PDCCH-singleSpan-r17", Optional: true},
				{Name: "supportedActivatedPRS-ProcessingWindow-r17", Optional: true},
				{Name: "cg-TimeDomainAllocationExtension-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtGuardSymbolReportReceptionIABR17Constraints)
			if err != nil {
				return err
			}
			ie.GuardSymbolReportReception_IAB_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtRestrictedIABDUBeamReceptionR17Constraints)
			if err != nil {
				return err
			}
			ie.Restricted_IAB_DU_BeamReception_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtRecommendedIABMTBeamTransmissionR17Constraints)
			if err != nil {
				return err
			}
			ie.Recommended_IAB_MT_BeamTransmission_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtCase6TimingAlignmentReceptionIABR17Constraints)
			if err != nil {
				return err
			}
			ie.Case6_TimingAlignmentReception_IAB_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtCase7TimingAlignmentReceptionIABR17Constraints)
			if err != nil {
				return err
			}
			ie.Case7_TimingAlignmentReception_IAB_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDlTxPowerAdjustmentIABR17Constraints)
			if err != nil {
				return err
			}
			ie.Dl_Tx_PowerAdjustment_IAB_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDesiredUlTxPowerAdjustmentR17Constraints)
			if err != nil {
				return err
			}
			ie.Desired_Ul_Tx_PowerAdjustment_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtFdmSoftResourceAvailabilityDynamicIndicationR17Constraints)
			if err != nil {
				return err
			}
			ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtUpdatedTDeltaRangeReceptionR17Constraints)
			if err != nil {
				return err
			}
			ie.Updated_T_DeltaRangeReception_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSlotBasedDynamicPUCCHRepR17Constraints)
			if err != nil {
				return err
			}
			ie.SlotBasedDynamicPUCCH_Rep_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			ie.Sps_HARQ_ACK_Deferral_r17 = &struct {
				Non_SharedSpectrumChAccess_r17 *int64
				SharedSpectrumChAccess_r17     *int64
			}{}
			c := ie.Sps_HARQ_ACK_Deferral_r17
			phyParametersCommonExtSpsHARQACKDeferralR17Seq := dx.NewSequenceDecoder(phyParametersCommonExtSpsHARQACKDeferralR17Constraints)
			if err := phyParametersCommonExtSpsHARQACKDeferralR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtSpsHARQACKDeferralR17Seq.IsComponentPresent(0) {
				c.Non_SharedSpectrumChAccess_r17 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtSpsHARQACKDeferralR17NonSharedSpectrumChAccessR17Constraints)
				if err != nil {
					return err
				}
				(*c.Non_SharedSpectrumChAccess_r17) = v
			}
			if phyParametersCommonExtSpsHARQACKDeferralR17Seq.IsComponentPresent(1) {
				c.SharedSpectrumChAccess_r17 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtSpsHARQACKDeferralR17SharedSpectrumChAccessR17Constraints)
				if err != nil {
					return err
				}
				(*c.SharedSpectrumChAccess_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeInteger(phyParametersCommonUnifiedJointTCICommonUpdateR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_CommonUpdate_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMTRPPDCCHSingleSpanR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PDCCH_SingleSpan_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSupportedActivatedPRSProcessingWindowR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedActivatedPRS_ProcessingWindow_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtCgTimeDomainAllocationExtensionR17Constraints)
			if err != nil {
				return err
			}
			ie.Cg_TimeDomainAllocationExtension_r17 = &v
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ta-BasedPDC-TN-NonSharedSpectrumChAccess-r17", Optional: true},
				{Name: "directionalCollisionDC-IAB-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtTaBasedPDCTNNonSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDirectionalCollisionDCIABR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectionalCollisionDC_IAB_r17 = &v
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dummy1", Optional: true},
				{Name: "dummy2", Optional: true},
				{Name: "dummy3", Optional: true},
				{Name: "dummy4", Optional: true},
				{Name: "srs-AdditionalRepetition-r17", Optional: true},
				{Name: "pusch-Repetition-CG-SDT-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDummy3Constraints)
			if err != nil {
				return err
			}
			ie.Dummy3 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDummy4Constraints)
			if err != nil {
				return err
			}
			ie.Dummy4 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtSrsAdditionalRepetitionR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_AdditionalRepetition_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPuschRepetitionCGSDTR17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_Repetition_CG_SDT_r17 = &v
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multiPDSCH-PerSlotType1-CB-Support-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMultiPDSCHPerSlotType1CBSupportR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPDSCH_PerSlotType1_CB_Support_r17 = &v
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "jointPowerSpatialAdaptation-r18", Optional: true},
				{Name: "ncr-AperiodicBeamInd-AccessLink-r18", Optional: true},
				{Name: "ncr-Semi-PersistentBeamInd-AccessLink-r18", Optional: true},
				{Name: "ncr-SimultaneousUL-BackhaulAndC-Link-r18", Optional: true},
				{Name: "ncr-BackhaulBeamInd-r18", Optional: true},
				{Name: "ncr-AdaptiveBeamBackhaulAndC-Link-r18", Optional: true},
				{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-1-3-r18", Optional: true},
				{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-0-3-r18", Optional: true},
				{Name: "configurableType-1A-FieldsForDCI-0-3-And-1-3-r18", Optional: true},
				{Name: "fdra-Type-1-Gty-2-4-8-16-RBs-RIV-DCI-1-3-And-0-3-r18", Optional: true},
				{Name: "priorityIndicationDL-r18", Optional: true},
				{Name: "priorityIndicationUL-r18", Optional: true},
				{Name: "dynamicIndicationSchedulingRestriction-r18", Optional: true},
				{Name: "priorityIndicationOneSlotHARQ-r18", Optional: true},
				{Name: "multiPUSCH-DCI-0-1-r18", Optional: true},
				{Name: "multiPUSCH-DCI-0-2-r18", Optional: true},
				{Name: "additionalSR-Periodicities-r18", Optional: true},
				{Name: "pathlossRS-UpdateForType1CG-PUSCH-r18", Optional: true},
				{Name: "bwp-SwitchingMultiDormancyCC-DCI-0-3-And-1-3-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtJointPowerSpatialAdaptationR18Constraints)
			if err != nil {
				return err
			}
			ie.JointPowerSpatialAdaptation_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ncr_AperiodicBeamInd_AccessLink_r18 = &struct {
				Scs_15kHz_r18  *int64
				Scs_30kHz_r18  *int64
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
			}{}
			c := ie.Ncr_AperiodicBeamInd_AccessLink_r18
			phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq := dx.NewSequenceDecoder(phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Constraints)
			if err := phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r18) = v
			}
			if phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r18) = v
			}
			if phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 2))
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r18) = v
			}
			if phyParametersCommonExtNcrAperiodicBeamIndAccessLinkR18Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(0, 2))
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNcrSemiPersistentBeamIndAccessLinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_Semi_PersistentBeamInd_AccessLink_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNcrSimultaneousULBackhaulAndCLinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_SimultaneousUL_BackhaulAndC_Link_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNcrBackhaulBeamIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_BackhaulBeamInd_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNcrAdaptiveBeamBackhaulAndCLinkR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_AdaptiveBeamBackhaulAndC_Link_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03R18Constraints)
			if err != nil {
				return err
			}
			ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtConfigurableType1AFieldsForDCI03And13R18Constraints)
			if err != nil {
				return err
			}
			ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03R18Constraints)
			if err != nil {
				return err
			}
			ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationDLR18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationDL_r18 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationULR18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationUL_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDynamicIndicationSchedulingRestrictionR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicIndicationSchedulingRestriction_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationOneSlotHARQR18Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationOneSlotHARQ_r18 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMultiPUSCHDCI01R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_DCI_0_1_r18 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtMultiPUSCHDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_DCI_0_2_r18 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			ie.AdditionalSR_Periodicities_r18 = &struct {
				Scs_30kHz_r18  *int64
				Scs_120kHz_r18 *int64
			}{}
			c := ie.AdditionalSR_Periodicities_r18
			phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq := dx.NewSequenceDecoder(phyParametersCommonExtAdditionalSRPeriodicitiesR18Constraints)
			if err := phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq.IsComponentPresent(0) {
				c.Scs_30kHz_r18 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs30kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r18) = v
			}
			if phyParametersCommonExtAdditionalSRPeriodicitiesR18Seq.IsComponentPresent(1) {
				c.Scs_120kHz_r18 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtAdditionalSRPeriodicitiesR18Scs120kHzR18Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPathlossRSUpdateForType1CGPUSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.PathlossRS_UpdateForType1CG_PUSCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18 = &struct {
				Choice    int
				Type1_r18 *int64
				Type2_r18 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Choice = int(index)
			switch index {
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type1_r18:
				(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type1_r18 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type1R18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type1_r18) = v
			case Phy_ParametersCommon_Ext_Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18_Type2_r18:
				(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type2_r18 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersCommonExtBwpSwitchingMultiDormancyCCDCI03And13R18Type2R18Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Bwp_SwitchingMultiDormancyCC_DCI_0_3_And_1_3_r18).Type2_r18) = v
			}
		}
	}

	// Extension group 13:
	if seq.IsExtensionPresent(13) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ncr-dft-S-OFDM-WaveformUL-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNcrDftSOFDMWaveformULR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_Dft_S_OFDM_WaveformUL_r18 = &v
		}
	}

	// Extension group 14:
	if seq.IsExtensionPresent(14) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "codebookVariantsListExt-r19", Optional: true},
				{Name: "codebookVariantsListAggregate-r19", Optional: true},
				{Name: "codebookVariantsListHybrid-r19", Optional: true},
				{Name: "dummy-EnableTx-RxDuringMeasGap-r19", Optional: true},
				{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-1-3-Diff-r19", Optional: true},
				{Name: "nominalRBG-SizeOfConfig-3-FDRA-Type-0-DCI-0-3-Diff-r19", Optional: true},
				{Name: "configurableType-1A-FieldsForDCI-0-3-And-1-3-Diff-r19", Optional: true},
				{Name: "fdra-Type-1-Gty-2-4-8-16-RBs-RIV-DCI-1-3-And-0-3-Diff-r19", Optional: true},
				{Name: "priorityIndicationDL-Diff-r19", Optional: true},
				{Name: "priorityIndicationUL-Diff-r19", Optional: true},
				{Name: "dynamicIndicationSchedulingRestrictionDiff-r19", Optional: true},
				{Name: "priorityIndicationOneSlotHARQ-Diff-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CodebookVariantsListExt_r19 = new(CodebookVariantsListExt_r19)
			if err := ie.CodebookVariantsListExt_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CodebookVariantsListAggregate_r19 = new(CodebookVariantsListAggregate_r19)
			if err := ie.CodebookVariantsListAggregate_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CodebookVariantsListHybrid_r19 = new(CodebookVariantsListHybrid_r19)
			if err := ie.CodebookVariantsListHybrid_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Dummy_EnableTx_RxDuringMeasGap_r19 = &struct {
				AdditionalDCI_r19     per.BitString
				IndicationField_r19   int64
				MinimumTimeOffset_r19 int64
			}{}
			c := ie.Dummy_EnableTx_RxDuringMeasGap_r19
			{
				v, err := dx.DecodeBitString(per.FixedSize(3))
				if err != nil {
					return err
				}
				c.AdditionalDCI_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints)
				if err != nil {
					return err
				}
				c.IndicationField_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(phyParametersCommonExtDummyEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints)
				if err != nil {
					return err
				}
				c.MinimumTimeOffset_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI13DiffR19Constraints)
			if err != nil {
				return err
			}
			ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_1_3_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtNominalRBGSizeOfConfig3FDRAType0DCI03DiffR19Constraints)
			if err != nil {
				return err
			}
			ie.NominalRBG_SizeOfConfig_3_FDRA_Type_0_DCI_0_3_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtConfigurableType1AFieldsForDCI03And13DiffR19Constraints)
			if err != nil {
				return err
			}
			ie.ConfigurableType_1A_FieldsForDCI_0_3_And_1_3_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtFdraType1Gty24816RBsRIVDCI13And03DiffR19Constraints)
			if err != nil {
				return err
			}
			ie.Fdra_Type_1_Gty_2_4_8_16_RBs_RIV_DCI_1_3_And_0_3_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationDLDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationDL_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationULDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationUL_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtDynamicIndicationSchedulingRestrictionDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.DynamicIndicationSchedulingRestrictionDiff_r19 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(phyParametersCommonExtPriorityIndicationOneSlotHARQDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicationOneSlotHARQ_Diff_r19 = &v
		}
	}

	return nil
}
