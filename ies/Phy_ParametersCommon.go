package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon struct {
	Csi_RS_CFRA_ForHO                                          *Phy_ParametersCommon_csi_RS_CFRA_ForHO                                          `optional`
	DynamicPRB_BundlingDL                                      *Phy_ParametersCommon_dynamicPRB_BundlingDL                                      `optional`
	Sp_CSI_ReportPUCCH                                         *Phy_ParametersCommon_sp_CSI_ReportPUCCH                                         `optional`
	Sp_CSI_ReportPUSCH                                         *Phy_ParametersCommon_sp_CSI_ReportPUSCH                                         `optional`
	Nzp_CSI_RS_IntefMgmt                                       *Phy_ParametersCommon_nzp_CSI_RS_IntefMgmt                                       `optional`
	Type2_SP_CSI_Feedback_LongPUCCH                            *Phy_ParametersCommon_type2_SP_CSI_Feedback_LongPUCCH                            `optional`
	PrecoderGranularityCORESET                                 *Phy_ParametersCommon_precoderGranularityCORESET                                 `optional`
	DynamicHARQ_ACK_Codebook                                   *Phy_ParametersCommon_dynamicHARQ_ACK_Codebook                                   `optional`
	SemiStaticHARQ_ACK_Codebook                                *Phy_ParametersCommon_semiStaticHARQ_ACK_Codebook                                `optional`
	SpatialBundlingHARQ_ACK                                    *Phy_ParametersCommon_spatialBundlingHARQ_ACK                                    `optional`
	DynamicBetaOffsetInd_HARQ_ACK_CSI                          *Phy_ParametersCommon_dynamicBetaOffsetInd_HARQ_ACK_CSI                          `optional`
	Pucch_Repetition_F1_3_4                                    *Phy_ParametersCommon_pucch_Repetition_F1_3_4                                    `optional`
	Ra_Type0_PUSCH                                             *Phy_ParametersCommon_ra_Type0_PUSCH                                             `optional`
	DynamicSwitchRA_Type0_1_PDSCH                              *Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PDSCH                              `optional`
	DynamicSwitchRA_Type0_1_PUSCH                              *Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PUSCH                              `optional`
	Pdsch_MappingTypeA                                         *Phy_ParametersCommon_pdsch_MappingTypeA                                         `optional`
	Pdsch_MappingTypeB                                         *Phy_ParametersCommon_pdsch_MappingTypeB                                         `optional`
	InterleavingVRB_ToPRB_PDSCH                                *Phy_ParametersCommon_interleavingVRB_ToPRB_PDSCH                                `optional`
	InterSlotFreqHopping_PUSCH                                 *Phy_ParametersCommon_interSlotFreqHopping_PUSCH                                 `optional`
	Type1_PUSCH_RepetitionMultiSlots                           *Phy_ParametersCommon_type1_PUSCH_RepetitionMultiSlots                           `optional`
	Type2_PUSCH_RepetitionMultiSlots                           *Phy_ParametersCommon_type2_PUSCH_RepetitionMultiSlots                           `optional`
	Pusch_RepetitionMultiSlots                                 *Phy_ParametersCommon_pusch_RepetitionMultiSlots                                 `optional`
	Pdsch_RepetitionMultiSlots                                 *Phy_ParametersCommon_pdsch_RepetitionMultiSlots                                 `optional`
	DownlinkSPS                                                *Phy_ParametersCommon_downlinkSPS                                                `optional`
	ConfiguredUL_GrantType1                                    *Phy_ParametersCommon_configuredUL_GrantType1                                    `optional`
	ConfiguredUL_GrantType2                                    *Phy_ParametersCommon_configuredUL_GrantType2                                    `optional`
	Pre_EmptIndication_DL                                      *Phy_ParametersCommon_pre_EmptIndication_DL                                      `optional`
	Cbg_TransIndication_DL                                     *Phy_ParametersCommon_cbg_TransIndication_DL                                     `optional`
	Cbg_TransIndication_UL                                     *Phy_ParametersCommon_cbg_TransIndication_UL                                     `optional`
	Cbg_FlushIndication_DL                                     *Phy_ParametersCommon_cbg_FlushIndication_DL                                     `optional`
	DynamicHARQ_ACK_CodeB_CBG_Retx_DL                          *Phy_ParametersCommon_dynamicHARQ_ACK_CodeB_CBG_Retx_DL                          `optional`
	RateMatchingResrcSetSemi_Static                            *Phy_ParametersCommon_rateMatchingResrcSetSemi_Static                            `optional`
	RateMatchingResrcSetDynamic                                *Phy_ParametersCommon_rateMatchingResrcSetDynamic                                `optional`
	Bwp_SwitchingDelay                                         *Phy_ParametersCommon_bwp_SwitchingDelay                                         `optional`
	Dummy                                                      *Phy_ParametersCommon_dummy                                                      `optional,ext-1`
	MaxNumberSearchSpaces                                      *Phy_ParametersCommon_maxNumberSearchSpaces                                      `optional,ext-2`
	RateMatchingCtrlResrcSetDynamic                            *Phy_ParametersCommon_rateMatchingCtrlResrcSetDynamic                            `optional,ext-2`
	MaxLayersMIMO_Indication                                   *Phy_ParametersCommon_maxLayersMIMO_Indication                                   `optional,ext-2`
	SpCellPlacement                                            *CarrierAggregationVariant                                                       `optional,ext-3`
	TwoStepRACH_r16                                            *Phy_ParametersCommon_twoStepRACH_r16                                            `optional,ext-4`
	Dci_Format1_2And0_2_r16                                    *Phy_ParametersCommon_dci_Format1_2And0_2_r16                                    `optional,ext-4`
	MonitoringDCI_SameSearchSpace_r16                          *Phy_ParametersCommon_monitoringDCI_SameSearchSpace_r16                          `optional,ext-4`
	Type2_CG_ReleaseDCI_0_1_r16                                *Phy_ParametersCommon_type2_CG_ReleaseDCI_0_1_r16                                `optional,ext-4`
	Type2_CG_ReleaseDCI_0_2_r16                                *Phy_ParametersCommon_type2_CG_ReleaseDCI_0_2_r16                                `optional,ext-4`
	Sps_ReleaseDCI_1_1_r16                                     *Phy_ParametersCommon_sps_ReleaseDCI_1_1_r16                                     `optional,ext-4`
	Sps_ReleaseDCI_1_2_r16                                     *Phy_ParametersCommon_sps_ReleaseDCI_1_2_r16                                     `optional,ext-4`
	Csi_TriggerStateNon_ActiveBWP_r16                          *Phy_ParametersCommon_csi_TriggerStateNon_ActiveBWP_r16                          `optional,ext-4`
	SeparateSMTC_InterIAB_Support_r16                          *Phy_ParametersCommon_separateSMTC_InterIAB_Support_r16                          `optional,ext-4`
	SeparateRACH_IAB_Support_r16                               *Phy_ParametersCommon_separateRACH_IAB_Support_r16                               `optional,ext-4`
	Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16                 *Phy_ParametersCommon_ul_flexibleDL_SlotFormatSemiStatic_IAB_r16                 `optional,ext-4`
	Ul_flexibleDL_SlotFormatDynamics_IAB_r16                   *Phy_ParametersCommon_ul_flexibleDL_SlotFormatDynamics_IAB_r16                   `optional,ext-4`
	Dft_S_OFDM_WaveformUL_IAB_r16                              *Phy_ParametersCommon_dft_S_OFDM_WaveformUL_IAB_r16                              `optional,ext-4`
	Dci_25_AI_RNTI_Support_IAB_r16                             *Phy_ParametersCommon_dci_25_AI_RNTI_Support_IAB_r16                             `optional,ext-4`
	T_DeltaReceptionSupport_IAB_r16                            *Phy_ParametersCommon_t_DeltaReceptionSupport_IAB_r16                            `optional,ext-4`
	GuardSymbolReportReception_IAB_r16                         *Phy_ParametersCommon_guardSymbolReportReception_IAB_r16                         `optional,ext-4`
	HarqACK_CB_SpatialBundlingPUCCH_Group_r16                  *Phy_ParametersCommon_harqACK_CB_SpatialBundlingPUCCH_Group_r16                  `optional,ext-4`
	CrossSlotScheduling_r16                                    *Phy_ParametersCommon_crossSlotScheduling_r16                                    `optional,ext-4`
	MaxNumberSRS_PosPathLossEstimateAllServingCells_r16        *Phy_ParametersCommon_maxNumberSRS_PosPathLossEstimateAllServingCells_r16        `optional,ext-4`
	ExtendedCG_Periodicities_r16                               *Phy_ParametersCommon_extendedCG_Periodicities_r16                               `optional,ext-4`
	ExtendedSPS_Periodicities_r16                              *Phy_ParametersCommon_extendedSPS_Periodicities_r16                              `optional,ext-4`
	CodebookVariantsList_r16                                   *CodebookVariantsList_r16                                                        `optional,ext-4`
	Pusch_RepetitionTypeA_r16                                  *Phy_ParametersCommon_pusch_RepetitionTypeA_r16                                  `optional,ext-4`
	Dci_DL_PriorityIndicator_r16                               *Phy_ParametersCommon_dci_DL_PriorityIndicator_r16                               `optional,ext-4`
	Dci_UL_PriorityIndicator_r16                               *Phy_ParametersCommon_dci_UL_PriorityIndicator_r16                               `optional,ext-4`
	MaxNumberPathlossRS_Update_r16                             *Phy_ParametersCommon_maxNumberPathlossRS_Update_r16                             `optional,ext-4`
	Type2_HARQ_ACK_Codebook_r16                                *Phy_ParametersCommon_type2_HARQ_ACK_Codebook_r16                                `optional,ext-4`
	MaxTotalResourcesForAcrossFreqRanges_r16                   *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16                   `optional,ext-4`
	HarqACK_separateMultiDCI_MultiTRP_r16                      *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16                      `optional,ext-4`
	HarqACK_jointMultiDCI_MultiTRP_r16                         *Phy_ParametersCommon_harqACK_jointMultiDCI_MultiTRP_r16                         `optional,ext-4`
	Bwp_SwitchingMultiCCs_r16                                  *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16                                  `optional,ext-4`
	TargetSMTC_SCG_r16                                         *Phy_ParametersCommon_targetSMTC_SCG_r16                                         `optional,ext-5`
	SupportRepetitionZeroOffsetRV_r16                          *Phy_ParametersCommon_supportRepetitionZeroOffsetRV_r16                          `optional,ext-5`
	Cbg_TransInOrderPUSCH_UL_r16                               *Phy_ParametersCommon_cbg_TransInOrderPUSCH_UL_r16                               `optional,ext-5`
	Bwp_SwitchingMultiDormancyCCs_r16                          *Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16                          `optional,ext-6`
	SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16             *Phy_ParametersCommon_supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16             `optional,ext-6`
	Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 *Phy_ParametersCommon_pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 `optional,ext-6`
	NewBeamIdentifications2PortCSI_RS_r16                      *Phy_ParametersCommon_newBeamIdentifications2PortCSI_RS_r16                      `optional,ext-7`
	PathlossEstimation2PortCSI_RS_r16                          *Phy_ParametersCommon_pathlossEstimation2PortCSI_RS_r16                          `optional,ext-7`
	Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16                      *Phy_ParametersCommon_mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16                      `optional,ext-8`
	GuardSymbolReportReception_IAB_r17                         *Phy_ParametersCommon_guardSymbolReportReception_IAB_r17                         `optional,ext-9`
	Restricted_IAB_DU_BeamReception_r17                        *Phy_ParametersCommon_restricted_IAB_DU_BeamReception_r17                        `optional,ext-9`
	Recommended_IAB_MT_BeamTransmission_r17                    *Phy_ParametersCommon_recommended_IAB_MT_BeamTransmission_r17                    `optional,ext-9`
	Case6_TimingAlignmentReception_IAB_r17                     *Phy_ParametersCommon_case6_TimingAlignmentReception_IAB_r17                     `optional,ext-9`
	Case7_TimingAlignmentReception_IAB_r17                     *Phy_ParametersCommon_case7_TimingAlignmentReception_IAB_r17                     `optional,ext-9`
	Dl_tx_PowerAdjustment_IAB_r17                              *Phy_ParametersCommon_dl_tx_PowerAdjustment_IAB_r17                              `optional,ext-9`
	Desired_ul_tx_PowerAdjustment_r17                          *Phy_ParametersCommon_desired_ul_tx_PowerAdjustment_r17                          `optional,ext-9`
	Fdm_SoftResourceAvailability_DynamicIndication_r17         *Phy_ParametersCommon_fdm_SoftResourceAvailability_DynamicIndication_r17         `optional,ext-9`
	Updated_T_DeltaRangeRecption_r17                           *Phy_ParametersCommon_updated_T_DeltaRangeRecption_r17                           `optional,ext-9`
	SlotBasedDynamicPUCCH_Rep_r17                              *Phy_ParametersCommon_slotBasedDynamicPUCCH_Rep_r17                              `optional,ext-9`
	Sps_HARQ_ACK_Deferral_r17                                  *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17                                  `optional,ext-9`
	UnifiedJointTCI_commonUpdate_r17                           *int64                                                                           `lb:1,ub:4,optional,ext-9`
	MTRP_PDCCH_singleSpan_r17                                  *Phy_ParametersCommon_mTRP_PDCCH_singleSpan_r17                                  `optional,ext-9`
	SupportedActivatedPRS_ProcessingWindow_r17                 *Phy_ParametersCommon_supportedActivatedPRS_ProcessingWindow_r17                 `optional,ext-9`
	Cg_TimeDomainAllocationExtension_r17                       *Phy_ParametersCommon_cg_TimeDomainAllocationExtension_r17                       `optional,ext-9`
	Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17               *Phy_ParametersCommon_ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17               `optional,ext-10`
	DirectionalCollisionDC_IAB_r17                             *Phy_ParametersCommon_directionalCollisionDC_IAB_r17                             `optional,ext-10`
	PriorityIndicatorInDCI_Multicast_r17                       *Phy_ParametersCommon_priorityIndicatorInDCI_Multicast_r17                       `optional,ext-11`
	PriorityIndicatorInDCI_SPS_Multicast_r17                   *Phy_ParametersCommon_priorityIndicatorInDCI_SPS_Multicast_r17                   `optional,ext-11`
	TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17             *Phy_ParametersCommon_twoHARQ_ACK_CodebookForUnicastAndMulticast_r17             `optional,ext-11`
	MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17                *Phy_ParametersCommon_multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17                `optional,ext-11`
	Srs_AdditionalRepetition_r17                               *Phy_ParametersCommon_srs_AdditionalRepetition_r17                               `optional,ext-11`
	Pusch_Repetition_CG_SDT_r17                                *Phy_ParametersCommon_pusch_Repetition_CG_SDT_r17                                `optional,ext-11`
}

func (ie *Phy_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Dummy != nil || ie.MaxNumberSearchSpaces != nil || ie.RateMatchingCtrlResrcSetDynamic != nil || ie.MaxLayersMIMO_Indication != nil || ie.SpCellPlacement != nil || ie.TwoStepRACH_r16 != nil || ie.Dci_Format1_2And0_2_r16 != nil || ie.MonitoringDCI_SameSearchSpace_r16 != nil || ie.Type2_CG_ReleaseDCI_0_1_r16 != nil || ie.Type2_CG_ReleaseDCI_0_2_r16 != nil || ie.Sps_ReleaseDCI_1_1_r16 != nil || ie.Sps_ReleaseDCI_1_2_r16 != nil || ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil || ie.SeparateSMTC_InterIAB_Support_r16 != nil || ie.SeparateRACH_IAB_Support_r16 != nil || ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil || ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil || ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil || ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil || ie.T_DeltaReceptionSupport_IAB_r16 != nil || ie.GuardSymbolReportReception_IAB_r16 != nil || ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil || ie.CrossSlotScheduling_r16 != nil || ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil || ie.ExtendedCG_Periodicities_r16 != nil || ie.ExtendedSPS_Periodicities_r16 != nil || ie.CodebookVariantsList_r16 != nil || ie.Pusch_RepetitionTypeA_r16 != nil || ie.Dci_DL_PriorityIndicator_r16 != nil || ie.Dci_UL_PriorityIndicator_r16 != nil || ie.MaxNumberPathlossRS_Update_r16 != nil || ie.Type2_HARQ_ACK_Codebook_r16 != nil || ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil || ie.HarqACK_separateMultiDCI_MultiTRP_r16 != nil || ie.HarqACK_jointMultiDCI_MultiTRP_r16 != nil || ie.Bwp_SwitchingMultiCCs_r16 != nil || ie.TargetSMTC_SCG_r16 != nil || ie.SupportRepetitionZeroOffsetRV_r16 != nil || ie.Cbg_TransInOrderPUSCH_UL_r16 != nil || ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil || ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil || ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil || ie.NewBeamIdentifications2PortCSI_RS_r16 != nil || ie.PathlossEstimation2PortCSI_RS_r16 != nil || ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil || ie.GuardSymbolReportReception_IAB_r17 != nil || ie.Restricted_IAB_DU_BeamReception_r17 != nil || ie.Recommended_IAB_MT_BeamTransmission_r17 != nil || ie.Case6_TimingAlignmentReception_IAB_r17 != nil || ie.Case7_TimingAlignmentReception_IAB_r17 != nil || ie.Dl_tx_PowerAdjustment_IAB_r17 != nil || ie.Desired_ul_tx_PowerAdjustment_r17 != nil || ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil || ie.Updated_T_DeltaRangeRecption_r17 != nil || ie.SlotBasedDynamicPUCCH_Rep_r17 != nil || ie.Sps_HARQ_ACK_Deferral_r17 != nil || ie.UnifiedJointTCI_commonUpdate_r17 != nil || ie.MTRP_PDCCH_singleSpan_r17 != nil || ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil || ie.Cg_TimeDomainAllocationExtension_r17 != nil || ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil || ie.DirectionalCollisionDC_IAB_r17 != nil || ie.PriorityIndicatorInDCI_Multicast_r17 != nil || ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil || ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil || ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil || ie.Srs_AdditionalRepetition_r17 != nil || ie.Pusch_Repetition_CG_SDT_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Csi_RS_CFRA_ForHO != nil, ie.DynamicPRB_BundlingDL != nil, ie.Sp_CSI_ReportPUCCH != nil, ie.Sp_CSI_ReportPUSCH != nil, ie.Nzp_CSI_RS_IntefMgmt != nil, ie.Type2_SP_CSI_Feedback_LongPUCCH != nil, ie.PrecoderGranularityCORESET != nil, ie.DynamicHARQ_ACK_Codebook != nil, ie.SemiStaticHARQ_ACK_Codebook != nil, ie.SpatialBundlingHARQ_ACK != nil, ie.DynamicBetaOffsetInd_HARQ_ACK_CSI != nil, ie.Pucch_Repetition_F1_3_4 != nil, ie.Ra_Type0_PUSCH != nil, ie.DynamicSwitchRA_Type0_1_PDSCH != nil, ie.DynamicSwitchRA_Type0_1_PUSCH != nil, ie.Pdsch_MappingTypeA != nil, ie.Pdsch_MappingTypeB != nil, ie.InterleavingVRB_ToPRB_PDSCH != nil, ie.InterSlotFreqHopping_PUSCH != nil, ie.Type1_PUSCH_RepetitionMultiSlots != nil, ie.Type2_PUSCH_RepetitionMultiSlots != nil, ie.Pusch_RepetitionMultiSlots != nil, ie.Pdsch_RepetitionMultiSlots != nil, ie.DownlinkSPS != nil, ie.ConfiguredUL_GrantType1 != nil, ie.ConfiguredUL_GrantType2 != nil, ie.Pre_EmptIndication_DL != nil, ie.Cbg_TransIndication_DL != nil, ie.Cbg_TransIndication_UL != nil, ie.Cbg_FlushIndication_DL != nil, ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil, ie.RateMatchingResrcSetSemi_Static != nil, ie.RateMatchingResrcSetDynamic != nil, ie.Bwp_SwitchingDelay != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Csi_RS_CFRA_ForHO != nil {
		if err = ie.Csi_RS_CFRA_ForHO.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_CFRA_ForHO", err)
		}
	}
	if ie.DynamicPRB_BundlingDL != nil {
		if err = ie.DynamicPRB_BundlingDL.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPRB_BundlingDL", err)
		}
	}
	if ie.Sp_CSI_ReportPUCCH != nil {
		if err = ie.Sp_CSI_ReportPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportPUCCH", err)
		}
	}
	if ie.Sp_CSI_ReportPUSCH != nil {
		if err = ie.Sp_CSI_ReportPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportPUSCH", err)
		}
	}
	if ie.Nzp_CSI_RS_IntefMgmt != nil {
		if err = ie.Nzp_CSI_RS_IntefMgmt.Encode(w); err != nil {
			return utils.WrapError("Encode Nzp_CSI_RS_IntefMgmt", err)
		}
	}
	if ie.Type2_SP_CSI_Feedback_LongPUCCH != nil {
		if err = ie.Type2_SP_CSI_Feedback_LongPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_SP_CSI_Feedback_LongPUCCH", err)
		}
	}
	if ie.PrecoderGranularityCORESET != nil {
		if err = ie.PrecoderGranularityCORESET.Encode(w); err != nil {
			return utils.WrapError("Encode PrecoderGranularityCORESET", err)
		}
	}
	if ie.DynamicHARQ_ACK_Codebook != nil {
		if err = ie.DynamicHARQ_ACK_Codebook.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicHARQ_ACK_Codebook", err)
		}
	}
	if ie.SemiStaticHARQ_ACK_Codebook != nil {
		if err = ie.SemiStaticHARQ_ACK_Codebook.Encode(w); err != nil {
			return utils.WrapError("Encode SemiStaticHARQ_ACK_Codebook", err)
		}
	}
	if ie.SpatialBundlingHARQ_ACK != nil {
		if err = ie.SpatialBundlingHARQ_ACK.Encode(w); err != nil {
			return utils.WrapError("Encode SpatialBundlingHARQ_ACK", err)
		}
	}
	if ie.DynamicBetaOffsetInd_HARQ_ACK_CSI != nil {
		if err = ie.DynamicBetaOffsetInd_HARQ_ACK_CSI.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicBetaOffsetInd_HARQ_ACK_CSI", err)
		}
	}
	if ie.Pucch_Repetition_F1_3_4 != nil {
		if err = ie.Pucch_Repetition_F1_3_4.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Repetition_F1_3_4", err)
		}
	}
	if ie.Ra_Type0_PUSCH != nil {
		if err = ie.Ra_Type0_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_Type0_PUSCH", err)
		}
	}
	if ie.DynamicSwitchRA_Type0_1_PDSCH != nil {
		if err = ie.DynamicSwitchRA_Type0_1_PDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSwitchRA_Type0_1_PDSCH", err)
		}
	}
	if ie.DynamicSwitchRA_Type0_1_PUSCH != nil {
		if err = ie.DynamicSwitchRA_Type0_1_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSwitchRA_Type0_1_PUSCH", err)
		}
	}
	if ie.Pdsch_MappingTypeA != nil {
		if err = ie.Pdsch_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_MappingTypeA", err)
		}
	}
	if ie.Pdsch_MappingTypeB != nil {
		if err = ie.Pdsch_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_MappingTypeB", err)
		}
	}
	if ie.InterleavingVRB_ToPRB_PDSCH != nil {
		if err = ie.InterleavingVRB_ToPRB_PDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode InterleavingVRB_ToPRB_PDSCH", err)
		}
	}
	if ie.InterSlotFreqHopping_PUSCH != nil {
		if err = ie.InterSlotFreqHopping_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode InterSlotFreqHopping_PUSCH", err)
		}
	}
	if ie.Type1_PUSCH_RepetitionMultiSlots != nil {
		if err = ie.Type1_PUSCH_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if ie.Type2_PUSCH_RepetitionMultiSlots != nil {
		if err = ie.Type2_PUSCH_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if ie.Pusch_RepetitionMultiSlots != nil {
		if err = ie.Pusch_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_RepetitionMultiSlots", err)
		}
	}
	if ie.Pdsch_RepetitionMultiSlots != nil {
		if err = ie.Pdsch_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_RepetitionMultiSlots", err)
		}
	}
	if ie.DownlinkSPS != nil {
		if err = ie.DownlinkSPS.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkSPS", err)
		}
	}
	if ie.ConfiguredUL_GrantType1 != nil {
		if err = ie.ConfiguredUL_GrantType1.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredUL_GrantType1", err)
		}
	}
	if ie.ConfiguredUL_GrantType2 != nil {
		if err = ie.ConfiguredUL_GrantType2.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredUL_GrantType2", err)
		}
	}
	if ie.Pre_EmptIndication_DL != nil {
		if err = ie.Pre_EmptIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode Pre_EmptIndication_DL", err)
		}
	}
	if ie.Cbg_TransIndication_DL != nil {
		if err = ie.Cbg_TransIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode Cbg_TransIndication_DL", err)
		}
	}
	if ie.Cbg_TransIndication_UL != nil {
		if err = ie.Cbg_TransIndication_UL.Encode(w); err != nil {
			return utils.WrapError("Encode Cbg_TransIndication_UL", err)
		}
	}
	if ie.Cbg_FlushIndication_DL != nil {
		if err = ie.Cbg_FlushIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode Cbg_FlushIndication_DL", err)
		}
	}
	if ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil {
		if err = ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicHARQ_ACK_CodeB_CBG_Retx_DL", err)
		}
	}
	if ie.RateMatchingResrcSetSemi_Static != nil {
		if err = ie.RateMatchingResrcSetSemi_Static.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchingResrcSetSemi_Static", err)
		}
	}
	if ie.RateMatchingResrcSetDynamic != nil {
		if err = ie.RateMatchingResrcSetDynamic.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchingResrcSetDynamic", err)
		}
	}
	if ie.Bwp_SwitchingDelay != nil {
		if err = ie.Bwp_SwitchingDelay.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_SwitchingDelay", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 11 bits for 11 extension groups
		extBitmap := []bool{ie.Dummy != nil, ie.MaxNumberSearchSpaces != nil || ie.RateMatchingCtrlResrcSetDynamic != nil || ie.MaxLayersMIMO_Indication != nil, ie.SpCellPlacement != nil, ie.TwoStepRACH_r16 != nil || ie.Dci_Format1_2And0_2_r16 != nil || ie.MonitoringDCI_SameSearchSpace_r16 != nil || ie.Type2_CG_ReleaseDCI_0_1_r16 != nil || ie.Type2_CG_ReleaseDCI_0_2_r16 != nil || ie.Sps_ReleaseDCI_1_1_r16 != nil || ie.Sps_ReleaseDCI_1_2_r16 != nil || ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil || ie.SeparateSMTC_InterIAB_Support_r16 != nil || ie.SeparateRACH_IAB_Support_r16 != nil || ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil || ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil || ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil || ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil || ie.T_DeltaReceptionSupport_IAB_r16 != nil || ie.GuardSymbolReportReception_IAB_r16 != nil || ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil || ie.CrossSlotScheduling_r16 != nil || ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil || ie.ExtendedCG_Periodicities_r16 != nil || ie.ExtendedSPS_Periodicities_r16 != nil || ie.CodebookVariantsList_r16 != nil || ie.Pusch_RepetitionTypeA_r16 != nil || ie.Dci_DL_PriorityIndicator_r16 != nil || ie.Dci_UL_PriorityIndicator_r16 != nil || ie.MaxNumberPathlossRS_Update_r16 != nil || ie.Type2_HARQ_ACK_Codebook_r16 != nil || ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil || ie.HarqACK_separateMultiDCI_MultiTRP_r16 != nil || ie.HarqACK_jointMultiDCI_MultiTRP_r16 != nil || ie.Bwp_SwitchingMultiCCs_r16 != nil, ie.TargetSMTC_SCG_r16 != nil || ie.SupportRepetitionZeroOffsetRV_r16 != nil || ie.Cbg_TransInOrderPUSCH_UL_r16 != nil, ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil || ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil || ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil, ie.NewBeamIdentifications2PortCSI_RS_r16 != nil || ie.PathlossEstimation2PortCSI_RS_r16 != nil, ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil, ie.GuardSymbolReportReception_IAB_r17 != nil || ie.Restricted_IAB_DU_BeamReception_r17 != nil || ie.Recommended_IAB_MT_BeamTransmission_r17 != nil || ie.Case6_TimingAlignmentReception_IAB_r17 != nil || ie.Case7_TimingAlignmentReception_IAB_r17 != nil || ie.Dl_tx_PowerAdjustment_IAB_r17 != nil || ie.Desired_ul_tx_PowerAdjustment_r17 != nil || ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil || ie.Updated_T_DeltaRangeRecption_r17 != nil || ie.SlotBasedDynamicPUCCH_Rep_r17 != nil || ie.Sps_HARQ_ACK_Deferral_r17 != nil || ie.UnifiedJointTCI_commonUpdate_r17 != nil || ie.MTRP_PDCCH_singleSpan_r17 != nil || ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil || ie.Cg_TimeDomainAllocationExtension_r17 != nil, ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil || ie.DirectionalCollisionDC_IAB_r17 != nil, ie.PriorityIndicatorInDCI_Multicast_r17 != nil || ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil || ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil || ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil || ie.Srs_AdditionalRepetition_r17 != nil || ie.Pusch_Repetition_CG_SDT_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dummy != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dummy optional
			if ie.Dummy != nil {
				if err = ie.Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.MaxNumberSearchSpaces != nil, ie.RateMatchingCtrlResrcSetDynamic != nil, ie.MaxLayersMIMO_Indication != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxNumberSearchSpaces optional
			if ie.MaxNumberSearchSpaces != nil {
				if err = ie.MaxNumberSearchSpaces.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberSearchSpaces", err)
				}
			}
			// encode RateMatchingCtrlResrcSetDynamic optional
			if ie.RateMatchingCtrlResrcSetDynamic != nil {
				if err = ie.RateMatchingCtrlResrcSetDynamic.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RateMatchingCtrlResrcSetDynamic", err)
				}
			}
			// encode MaxLayersMIMO_Indication optional
			if ie.MaxLayersMIMO_Indication != nil {
				if err = ie.MaxLayersMIMO_Indication.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxLayersMIMO_Indication", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.SpCellPlacement != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpCellPlacement optional
			if ie.SpCellPlacement != nil {
				if err = ie.SpCellPlacement.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpCellPlacement", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.TwoStepRACH_r16 != nil, ie.Dci_Format1_2And0_2_r16 != nil, ie.MonitoringDCI_SameSearchSpace_r16 != nil, ie.Type2_CG_ReleaseDCI_0_1_r16 != nil, ie.Type2_CG_ReleaseDCI_0_2_r16 != nil, ie.Sps_ReleaseDCI_1_1_r16 != nil, ie.Sps_ReleaseDCI_1_2_r16 != nil, ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil, ie.SeparateSMTC_InterIAB_Support_r16 != nil, ie.SeparateRACH_IAB_Support_r16 != nil, ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil, ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil, ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil, ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil, ie.T_DeltaReceptionSupport_IAB_r16 != nil, ie.GuardSymbolReportReception_IAB_r16 != nil, ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil, ie.CrossSlotScheduling_r16 != nil, ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil, ie.ExtendedCG_Periodicities_r16 != nil, ie.ExtendedSPS_Periodicities_r16 != nil, ie.CodebookVariantsList_r16 != nil, ie.Pusch_RepetitionTypeA_r16 != nil, ie.Dci_DL_PriorityIndicator_r16 != nil, ie.Dci_UL_PriorityIndicator_r16 != nil, ie.MaxNumberPathlossRS_Update_r16 != nil, ie.Type2_HARQ_ACK_Codebook_r16 != nil, ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil, ie.HarqACK_separateMultiDCI_MultiTRP_r16 != nil, ie.HarqACK_jointMultiDCI_MultiTRP_r16 != nil, ie.Bwp_SwitchingMultiCCs_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TwoStepRACH_r16 optional
			if ie.TwoStepRACH_r16 != nil {
				if err = ie.TwoStepRACH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoStepRACH_r16", err)
				}
			}
			// encode Dci_Format1_2And0_2_r16 optional
			if ie.Dci_Format1_2And0_2_r16 != nil {
				if err = ie.Dci_Format1_2And0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dci_Format1_2And0_2_r16", err)
				}
			}
			// encode MonitoringDCI_SameSearchSpace_r16 optional
			if ie.MonitoringDCI_SameSearchSpace_r16 != nil {
				if err = ie.MonitoringDCI_SameSearchSpace_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MonitoringDCI_SameSearchSpace_r16", err)
				}
			}
			// encode Type2_CG_ReleaseDCI_0_1_r16 optional
			if ie.Type2_CG_ReleaseDCI_0_1_r16 != nil {
				if err = ie.Type2_CG_ReleaseDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type2_CG_ReleaseDCI_0_1_r16", err)
				}
			}
			// encode Type2_CG_ReleaseDCI_0_2_r16 optional
			if ie.Type2_CG_ReleaseDCI_0_2_r16 != nil {
				if err = ie.Type2_CG_ReleaseDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type2_CG_ReleaseDCI_0_2_r16", err)
				}
			}
			// encode Sps_ReleaseDCI_1_1_r16 optional
			if ie.Sps_ReleaseDCI_1_1_r16 != nil {
				if err = ie.Sps_ReleaseDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ReleaseDCI_1_1_r16", err)
				}
			}
			// encode Sps_ReleaseDCI_1_2_r16 optional
			if ie.Sps_ReleaseDCI_1_2_r16 != nil {
				if err = ie.Sps_ReleaseDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_ReleaseDCI_1_2_r16", err)
				}
			}
			// encode Csi_TriggerStateNon_ActiveBWP_r16 optional
			if ie.Csi_TriggerStateNon_ActiveBWP_r16 != nil {
				if err = ie.Csi_TriggerStateNon_ActiveBWP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_TriggerStateNon_ActiveBWP_r16", err)
				}
			}
			// encode SeparateSMTC_InterIAB_Support_r16 optional
			if ie.SeparateSMTC_InterIAB_Support_r16 != nil {
				if err = ie.SeparateSMTC_InterIAB_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SeparateSMTC_InterIAB_Support_r16", err)
				}
			}
			// encode SeparateRACH_IAB_Support_r16 optional
			if ie.SeparateRACH_IAB_Support_r16 != nil {
				if err = ie.SeparateRACH_IAB_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SeparateRACH_IAB_Support_r16", err)
				}
			}
			// encode Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 optional
			if ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil {
				if err = ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16", err)
				}
			}
			// encode Ul_flexibleDL_SlotFormatDynamics_IAB_r16 optional
			if ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil {
				if err = ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_flexibleDL_SlotFormatDynamics_IAB_r16", err)
				}
			}
			// encode Dft_S_OFDM_WaveformUL_IAB_r16 optional
			if ie.Dft_S_OFDM_WaveformUL_IAB_r16 != nil {
				if err = ie.Dft_S_OFDM_WaveformUL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dft_S_OFDM_WaveformUL_IAB_r16", err)
				}
			}
			// encode Dci_25_AI_RNTI_Support_IAB_r16 optional
			if ie.Dci_25_AI_RNTI_Support_IAB_r16 != nil {
				if err = ie.Dci_25_AI_RNTI_Support_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dci_25_AI_RNTI_Support_IAB_r16", err)
				}
			}
			// encode T_DeltaReceptionSupport_IAB_r16 optional
			if ie.T_DeltaReceptionSupport_IAB_r16 != nil {
				if err = ie.T_DeltaReceptionSupport_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode T_DeltaReceptionSupport_IAB_r16", err)
				}
			}
			// encode GuardSymbolReportReception_IAB_r16 optional
			if ie.GuardSymbolReportReception_IAB_r16 != nil {
				if err = ie.GuardSymbolReportReception_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GuardSymbolReportReception_IAB_r16", err)
				}
			}
			// encode HarqACK_CB_SpatialBundlingPUCCH_Group_r16 optional
			if ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil {
				if err = ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HarqACK_CB_SpatialBundlingPUCCH_Group_r16", err)
				}
			}
			// encode CrossSlotScheduling_r16 optional
			if ie.CrossSlotScheduling_r16 != nil {
				if err = ie.CrossSlotScheduling_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CrossSlotScheduling_r16", err)
				}
			}
			// encode MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 optional
			if ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil {
				if err = ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberSRS_PosPathLossEstimateAllServingCells_r16", err)
				}
			}
			// encode ExtendedCG_Periodicities_r16 optional
			if ie.ExtendedCG_Periodicities_r16 != nil {
				if err = ie.ExtendedCG_Periodicities_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedCG_Periodicities_r16", err)
				}
			}
			// encode ExtendedSPS_Periodicities_r16 optional
			if ie.ExtendedSPS_Periodicities_r16 != nil {
				if err = ie.ExtendedSPS_Periodicities_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ExtendedSPS_Periodicities_r16", err)
				}
			}
			// encode CodebookVariantsList_r16 optional
			if ie.CodebookVariantsList_r16 != nil {
				if err = ie.CodebookVariantsList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookVariantsList_r16", err)
				}
			}
			// encode Pusch_RepetitionTypeA_r16 optional
			if ie.Pusch_RepetitionTypeA_r16 != nil {
				if err = ie.Pusch_RepetitionTypeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_RepetitionTypeA_r16", err)
				}
			}
			// encode Dci_DL_PriorityIndicator_r16 optional
			if ie.Dci_DL_PriorityIndicator_r16 != nil {
				if err = ie.Dci_DL_PriorityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dci_DL_PriorityIndicator_r16", err)
				}
			}
			// encode Dci_UL_PriorityIndicator_r16 optional
			if ie.Dci_UL_PriorityIndicator_r16 != nil {
				if err = ie.Dci_UL_PriorityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dci_UL_PriorityIndicator_r16", err)
				}
			}
			// encode MaxNumberPathlossRS_Update_r16 optional
			if ie.MaxNumberPathlossRS_Update_r16 != nil {
				if err = ie.MaxNumberPathlossRS_Update_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberPathlossRS_Update_r16", err)
				}
			}
			// encode Type2_HARQ_ACK_Codebook_r16 optional
			if ie.Type2_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.Type2_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type2_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode MaxTotalResourcesForAcrossFreqRanges_r16 optional
			if ie.MaxTotalResourcesForAcrossFreqRanges_r16 != nil {
				if err = ie.MaxTotalResourcesForAcrossFreqRanges_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxTotalResourcesForAcrossFreqRanges_r16", err)
				}
			}
			// encode HarqACK_separateMultiDCI_MultiTRP_r16 optional
			if ie.HarqACK_separateMultiDCI_MultiTRP_r16 != nil {
				if err = ie.HarqACK_separateMultiDCI_MultiTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HarqACK_separateMultiDCI_MultiTRP_r16", err)
				}
			}
			// encode HarqACK_jointMultiDCI_MultiTRP_r16 optional
			if ie.HarqACK_jointMultiDCI_MultiTRP_r16 != nil {
				if err = ie.HarqACK_jointMultiDCI_MultiTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HarqACK_jointMultiDCI_MultiTRP_r16", err)
				}
			}
			// encode Bwp_SwitchingMultiCCs_r16 optional
			if ie.Bwp_SwitchingMultiCCs_r16 != nil {
				if err = ie.Bwp_SwitchingMultiCCs_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bwp_SwitchingMultiCCs_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.TargetSMTC_SCG_r16 != nil, ie.SupportRepetitionZeroOffsetRV_r16 != nil, ie.Cbg_TransInOrderPUSCH_UL_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TargetSMTC_SCG_r16 optional
			if ie.TargetSMTC_SCG_r16 != nil {
				if err = ie.TargetSMTC_SCG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TargetSMTC_SCG_r16", err)
				}
			}
			// encode SupportRepetitionZeroOffsetRV_r16 optional
			if ie.SupportRepetitionZeroOffsetRV_r16 != nil {
				if err = ie.SupportRepetitionZeroOffsetRV_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportRepetitionZeroOffsetRV_r16", err)
				}
			}
			// encode Cbg_TransInOrderPUSCH_UL_r16 optional
			if ie.Cbg_TransInOrderPUSCH_UL_r16 != nil {
				if err = ie.Cbg_TransInOrderPUSCH_UL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cbg_TransInOrderPUSCH_UL_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil, ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil, ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Bwp_SwitchingMultiDormancyCCs_r16 optional
			if ie.Bwp_SwitchingMultiDormancyCCs_r16 != nil {
				if err = ie.Bwp_SwitchingMultiDormancyCCs_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bwp_SwitchingMultiDormancyCCs_r16", err)
				}
			}
			// encode SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 optional
			if ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil {
				if err = ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16", err)
				}
			}
			// encode Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 optional
			if ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil {
				if err = ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.NewBeamIdentifications2PortCSI_RS_r16 != nil, ie.PathlossEstimation2PortCSI_RS_r16 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode NewBeamIdentifications2PortCSI_RS_r16 optional
			if ie.NewBeamIdentifications2PortCSI_RS_r16 != nil {
				if err = ie.NewBeamIdentifications2PortCSI_RS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NewBeamIdentifications2PortCSI_RS_r16", err)
				}
			}
			// encode PathlossEstimation2PortCSI_RS_r16 optional
			if ie.PathlossEstimation2PortCSI_RS_r16 != nil {
				if err = ie.PathlossEstimation2PortCSI_RS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PathlossEstimation2PortCSI_RS_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 optional
			if ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil {
				if err = ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 9
		if extBitmap[8] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.GuardSymbolReportReception_IAB_r17 != nil, ie.Restricted_IAB_DU_BeamReception_r17 != nil, ie.Recommended_IAB_MT_BeamTransmission_r17 != nil, ie.Case6_TimingAlignmentReception_IAB_r17 != nil, ie.Case7_TimingAlignmentReception_IAB_r17 != nil, ie.Dl_tx_PowerAdjustment_IAB_r17 != nil, ie.Desired_ul_tx_PowerAdjustment_r17 != nil, ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil, ie.Updated_T_DeltaRangeRecption_r17 != nil, ie.SlotBasedDynamicPUCCH_Rep_r17 != nil, ie.Sps_HARQ_ACK_Deferral_r17 != nil, ie.UnifiedJointTCI_commonUpdate_r17 != nil, ie.MTRP_PDCCH_singleSpan_r17 != nil, ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil, ie.Cg_TimeDomainAllocationExtension_r17 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode GuardSymbolReportReception_IAB_r17 optional
			if ie.GuardSymbolReportReception_IAB_r17 != nil {
				if err = ie.GuardSymbolReportReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GuardSymbolReportReception_IAB_r17", err)
				}
			}
			// encode Restricted_IAB_DU_BeamReception_r17 optional
			if ie.Restricted_IAB_DU_BeamReception_r17 != nil {
				if err = ie.Restricted_IAB_DU_BeamReception_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Restricted_IAB_DU_BeamReception_r17", err)
				}
			}
			// encode Recommended_IAB_MT_BeamTransmission_r17 optional
			if ie.Recommended_IAB_MT_BeamTransmission_r17 != nil {
				if err = ie.Recommended_IAB_MT_BeamTransmission_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Recommended_IAB_MT_BeamTransmission_r17", err)
				}
			}
			// encode Case6_TimingAlignmentReception_IAB_r17 optional
			if ie.Case6_TimingAlignmentReception_IAB_r17 != nil {
				if err = ie.Case6_TimingAlignmentReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Case6_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// encode Case7_TimingAlignmentReception_IAB_r17 optional
			if ie.Case7_TimingAlignmentReception_IAB_r17 != nil {
				if err = ie.Case7_TimingAlignmentReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Case7_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// encode Dl_tx_PowerAdjustment_IAB_r17 optional
			if ie.Dl_tx_PowerAdjustment_IAB_r17 != nil {
				if err = ie.Dl_tx_PowerAdjustment_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_tx_PowerAdjustment_IAB_r17", err)
				}
			}
			// encode Desired_ul_tx_PowerAdjustment_r17 optional
			if ie.Desired_ul_tx_PowerAdjustment_r17 != nil {
				if err = ie.Desired_ul_tx_PowerAdjustment_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Desired_ul_tx_PowerAdjustment_r17", err)
				}
			}
			// encode Fdm_SoftResourceAvailability_DynamicIndication_r17 optional
			if ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 != nil {
				if err = ie.Fdm_SoftResourceAvailability_DynamicIndication_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Fdm_SoftResourceAvailability_DynamicIndication_r17", err)
				}
			}
			// encode Updated_T_DeltaRangeRecption_r17 optional
			if ie.Updated_T_DeltaRangeRecption_r17 != nil {
				if err = ie.Updated_T_DeltaRangeRecption_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Updated_T_DeltaRangeRecption_r17", err)
				}
			}
			// encode SlotBasedDynamicPUCCH_Rep_r17 optional
			if ie.SlotBasedDynamicPUCCH_Rep_r17 != nil {
				if err = ie.SlotBasedDynamicPUCCH_Rep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SlotBasedDynamicPUCCH_Rep_r17", err)
				}
			}
			// encode Sps_HARQ_ACK_Deferral_r17 optional
			if ie.Sps_HARQ_ACK_Deferral_r17 != nil {
				if err = ie.Sps_HARQ_ACK_Deferral_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_HARQ_ACK_Deferral_r17", err)
				}
			}
			// encode UnifiedJointTCI_commonUpdate_r17 optional
			if ie.UnifiedJointTCI_commonUpdate_r17 != nil {
				if err = extWriter.WriteInteger(*ie.UnifiedJointTCI_commonUpdate_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_commonUpdate_r17", err)
				}
			}
			// encode MTRP_PDCCH_singleSpan_r17 optional
			if ie.MTRP_PDCCH_singleSpan_r17 != nil {
				if err = ie.MTRP_PDCCH_singleSpan_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PDCCH_singleSpan_r17", err)
				}
			}
			// encode SupportedActivatedPRS_ProcessingWindow_r17 optional
			if ie.SupportedActivatedPRS_ProcessingWindow_r17 != nil {
				if err = ie.SupportedActivatedPRS_ProcessingWindow_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedActivatedPRS_ProcessingWindow_r17", err)
				}
			}
			// encode Cg_TimeDomainAllocationExtension_r17 optional
			if ie.Cg_TimeDomainAllocationExtension_r17 != nil {
				if err = ie.Cg_TimeDomainAllocationExtension_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_TimeDomainAllocationExtension_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 10
		if extBitmap[9] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil, ie.DirectionalCollisionDC_IAB_r17 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 optional
			if ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// encode DirectionalCollisionDC_IAB_r17 optional
			if ie.DirectionalCollisionDC_IAB_r17 != nil {
				if err = ie.DirectionalCollisionDC_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DirectionalCollisionDC_IAB_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 11
		if extBitmap[10] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.PriorityIndicatorInDCI_Multicast_r17 != nil, ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil, ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil, ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil, ie.Srs_AdditionalRepetition_r17 != nil, ie.Pusch_Repetition_CG_SDT_r17 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PriorityIndicatorInDCI_Multicast_r17 optional
			if ie.PriorityIndicatorInDCI_Multicast_r17 != nil {
				if err = ie.PriorityIndicatorInDCI_Multicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorInDCI_Multicast_r17", err)
				}
			}
			// encode PriorityIndicatorInDCI_SPS_Multicast_r17 optional
			if ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil {
				if err = ie.PriorityIndicatorInDCI_SPS_Multicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PriorityIndicatorInDCI_SPS_Multicast_r17", err)
				}
			}
			// encode TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 optional
			if ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil {
				if err = ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17", err)
				}
			}
			// encode MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 optional
			if ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil {
				if err = ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17", err)
				}
			}
			// encode Srs_AdditionalRepetition_r17 optional
			if ie.Srs_AdditionalRepetition_r17 != nil {
				if err = ie.Srs_AdditionalRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_AdditionalRepetition_r17", err)
				}
			}
			// encode Pusch_Repetition_CG_SDT_r17 optional
			if ie.Pusch_Repetition_CG_SDT_r17 != nil {
				if err = ie.Pusch_Repetition_CG_SDT_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_Repetition_CG_SDT_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *Phy_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_CFRA_ForHOPresent bool
	if Csi_RS_CFRA_ForHOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicPRB_BundlingDLPresent bool
	if DynamicPRB_BundlingDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportPUCCHPresent bool
	if Sp_CSI_ReportPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportPUSCHPresent bool
	if Sp_CSI_ReportPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Nzp_CSI_RS_IntefMgmtPresent bool
	if Nzp_CSI_RS_IntefMgmtPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_SP_CSI_Feedback_LongPUCCHPresent bool
	if Type2_SP_CSI_Feedback_LongPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PrecoderGranularityCORESETPresent bool
	if PrecoderGranularityCORESETPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicHARQ_ACK_CodebookPresent bool
	if DynamicHARQ_ACK_CodebookPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiStaticHARQ_ACK_CodebookPresent bool
	if SemiStaticHARQ_ACK_CodebookPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpatialBundlingHARQ_ACKPresent bool
	if SpatialBundlingHARQ_ACKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicBetaOffsetInd_HARQ_ACK_CSIPresent bool
	if DynamicBetaOffsetInd_HARQ_ACK_CSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_Repetition_F1_3_4Present bool
	if Pucch_Repetition_F1_3_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_Type0_PUSCHPresent bool
	if Ra_Type0_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSwitchRA_Type0_1_PDSCHPresent bool
	if DynamicSwitchRA_Type0_1_PDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSwitchRA_Type0_1_PUSCHPresent bool
	if DynamicSwitchRA_Type0_1_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_MappingTypeAPresent bool
	if Pdsch_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_MappingTypeBPresent bool
	if Pdsch_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InterleavingVRB_ToPRB_PDSCHPresent bool
	if InterleavingVRB_ToPRB_PDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InterSlotFreqHopping_PUSCHPresent bool
	if InterSlotFreqHopping_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_PUSCH_RepetitionMultiSlotsPresent bool
	if Type1_PUSCH_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_PUSCH_RepetitionMultiSlotsPresent bool
	if Type2_PUSCH_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_RepetitionMultiSlotsPresent bool
	if Pusch_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_RepetitionMultiSlotsPresent bool
	if Pdsch_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkSPSPresent bool
	if DownlinkSPSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredUL_GrantType1Present bool
	if ConfiguredUL_GrantType1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredUL_GrantType2Present bool
	if ConfiguredUL_GrantType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pre_EmptIndication_DLPresent bool
	if Pre_EmptIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cbg_TransIndication_DLPresent bool
	if Cbg_TransIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cbg_TransIndication_ULPresent bool
	if Cbg_TransIndication_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Cbg_FlushIndication_DLPresent bool
	if Cbg_FlushIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent bool
	if DynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchingResrcSetSemi_StaticPresent bool
	if RateMatchingResrcSetSemi_StaticPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchingResrcSetDynamicPresent bool
	if RateMatchingResrcSetDynamicPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_SwitchingDelayPresent bool
	if Bwp_SwitchingDelayPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Csi_RS_CFRA_ForHOPresent {
		ie.Csi_RS_CFRA_ForHO = new(Phy_ParametersCommon_csi_RS_CFRA_ForHO)
		if err = ie.Csi_RS_CFRA_ForHO.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_CFRA_ForHO", err)
		}
	}
	if DynamicPRB_BundlingDLPresent {
		ie.DynamicPRB_BundlingDL = new(Phy_ParametersCommon_dynamicPRB_BundlingDL)
		if err = ie.DynamicPRB_BundlingDL.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicPRB_BundlingDL", err)
		}
	}
	if Sp_CSI_ReportPUCCHPresent {
		ie.Sp_CSI_ReportPUCCH = new(Phy_ParametersCommon_sp_CSI_ReportPUCCH)
		if err = ie.Sp_CSI_ReportPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportPUCCH", err)
		}
	}
	if Sp_CSI_ReportPUSCHPresent {
		ie.Sp_CSI_ReportPUSCH = new(Phy_ParametersCommon_sp_CSI_ReportPUSCH)
		if err = ie.Sp_CSI_ReportPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportPUSCH", err)
		}
	}
	if Nzp_CSI_RS_IntefMgmtPresent {
		ie.Nzp_CSI_RS_IntefMgmt = new(Phy_ParametersCommon_nzp_CSI_RS_IntefMgmt)
		if err = ie.Nzp_CSI_RS_IntefMgmt.Decode(r); err != nil {
			return utils.WrapError("Decode Nzp_CSI_RS_IntefMgmt", err)
		}
	}
	if Type2_SP_CSI_Feedback_LongPUCCHPresent {
		ie.Type2_SP_CSI_Feedback_LongPUCCH = new(Phy_ParametersCommon_type2_SP_CSI_Feedback_LongPUCCH)
		if err = ie.Type2_SP_CSI_Feedback_LongPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_SP_CSI_Feedback_LongPUCCH", err)
		}
	}
	if PrecoderGranularityCORESETPresent {
		ie.PrecoderGranularityCORESET = new(Phy_ParametersCommon_precoderGranularityCORESET)
		if err = ie.PrecoderGranularityCORESET.Decode(r); err != nil {
			return utils.WrapError("Decode PrecoderGranularityCORESET", err)
		}
	}
	if DynamicHARQ_ACK_CodebookPresent {
		ie.DynamicHARQ_ACK_Codebook = new(Phy_ParametersCommon_dynamicHARQ_ACK_Codebook)
		if err = ie.DynamicHARQ_ACK_Codebook.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicHARQ_ACK_Codebook", err)
		}
	}
	if SemiStaticHARQ_ACK_CodebookPresent {
		ie.SemiStaticHARQ_ACK_Codebook = new(Phy_ParametersCommon_semiStaticHARQ_ACK_Codebook)
		if err = ie.SemiStaticHARQ_ACK_Codebook.Decode(r); err != nil {
			return utils.WrapError("Decode SemiStaticHARQ_ACK_Codebook", err)
		}
	}
	if SpatialBundlingHARQ_ACKPresent {
		ie.SpatialBundlingHARQ_ACK = new(Phy_ParametersCommon_spatialBundlingHARQ_ACK)
		if err = ie.SpatialBundlingHARQ_ACK.Decode(r); err != nil {
			return utils.WrapError("Decode SpatialBundlingHARQ_ACK", err)
		}
	}
	if DynamicBetaOffsetInd_HARQ_ACK_CSIPresent {
		ie.DynamicBetaOffsetInd_HARQ_ACK_CSI = new(Phy_ParametersCommon_dynamicBetaOffsetInd_HARQ_ACK_CSI)
		if err = ie.DynamicBetaOffsetInd_HARQ_ACK_CSI.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicBetaOffsetInd_HARQ_ACK_CSI", err)
		}
	}
	if Pucch_Repetition_F1_3_4Present {
		ie.Pucch_Repetition_F1_3_4 = new(Phy_ParametersCommon_pucch_Repetition_F1_3_4)
		if err = ie.Pucch_Repetition_F1_3_4.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Repetition_F1_3_4", err)
		}
	}
	if Ra_Type0_PUSCHPresent {
		ie.Ra_Type0_PUSCH = new(Phy_ParametersCommon_ra_Type0_PUSCH)
		if err = ie.Ra_Type0_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_Type0_PUSCH", err)
		}
	}
	if DynamicSwitchRA_Type0_1_PDSCHPresent {
		ie.DynamicSwitchRA_Type0_1_PDSCH = new(Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PDSCH)
		if err = ie.DynamicSwitchRA_Type0_1_PDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSwitchRA_Type0_1_PDSCH", err)
		}
	}
	if DynamicSwitchRA_Type0_1_PUSCHPresent {
		ie.DynamicSwitchRA_Type0_1_PUSCH = new(Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PUSCH)
		if err = ie.DynamicSwitchRA_Type0_1_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSwitchRA_Type0_1_PUSCH", err)
		}
	}
	if Pdsch_MappingTypeAPresent {
		ie.Pdsch_MappingTypeA = new(Phy_ParametersCommon_pdsch_MappingTypeA)
		if err = ie.Pdsch_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_MappingTypeA", err)
		}
	}
	if Pdsch_MappingTypeBPresent {
		ie.Pdsch_MappingTypeB = new(Phy_ParametersCommon_pdsch_MappingTypeB)
		if err = ie.Pdsch_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_MappingTypeB", err)
		}
	}
	if InterleavingVRB_ToPRB_PDSCHPresent {
		ie.InterleavingVRB_ToPRB_PDSCH = new(Phy_ParametersCommon_interleavingVRB_ToPRB_PDSCH)
		if err = ie.InterleavingVRB_ToPRB_PDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode InterleavingVRB_ToPRB_PDSCH", err)
		}
	}
	if InterSlotFreqHopping_PUSCHPresent {
		ie.InterSlotFreqHopping_PUSCH = new(Phy_ParametersCommon_interSlotFreqHopping_PUSCH)
		if err = ie.InterSlotFreqHopping_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode InterSlotFreqHopping_PUSCH", err)
		}
	}
	if Type1_PUSCH_RepetitionMultiSlotsPresent {
		ie.Type1_PUSCH_RepetitionMultiSlots = new(Phy_ParametersCommon_type1_PUSCH_RepetitionMultiSlots)
		if err = ie.Type1_PUSCH_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if Type2_PUSCH_RepetitionMultiSlotsPresent {
		ie.Type2_PUSCH_RepetitionMultiSlots = new(Phy_ParametersCommon_type2_PUSCH_RepetitionMultiSlots)
		if err = ie.Type2_PUSCH_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if Pusch_RepetitionMultiSlotsPresent {
		ie.Pusch_RepetitionMultiSlots = new(Phy_ParametersCommon_pusch_RepetitionMultiSlots)
		if err = ie.Pusch_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_RepetitionMultiSlots", err)
		}
	}
	if Pdsch_RepetitionMultiSlotsPresent {
		ie.Pdsch_RepetitionMultiSlots = new(Phy_ParametersCommon_pdsch_RepetitionMultiSlots)
		if err = ie.Pdsch_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_RepetitionMultiSlots", err)
		}
	}
	if DownlinkSPSPresent {
		ie.DownlinkSPS = new(Phy_ParametersCommon_downlinkSPS)
		if err = ie.DownlinkSPS.Decode(r); err != nil {
			return utils.WrapError("Decode DownlinkSPS", err)
		}
	}
	if ConfiguredUL_GrantType1Present {
		ie.ConfiguredUL_GrantType1 = new(Phy_ParametersCommon_configuredUL_GrantType1)
		if err = ie.ConfiguredUL_GrantType1.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredUL_GrantType1", err)
		}
	}
	if ConfiguredUL_GrantType2Present {
		ie.ConfiguredUL_GrantType2 = new(Phy_ParametersCommon_configuredUL_GrantType2)
		if err = ie.ConfiguredUL_GrantType2.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredUL_GrantType2", err)
		}
	}
	if Pre_EmptIndication_DLPresent {
		ie.Pre_EmptIndication_DL = new(Phy_ParametersCommon_pre_EmptIndication_DL)
		if err = ie.Pre_EmptIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode Pre_EmptIndication_DL", err)
		}
	}
	if Cbg_TransIndication_DLPresent {
		ie.Cbg_TransIndication_DL = new(Phy_ParametersCommon_cbg_TransIndication_DL)
		if err = ie.Cbg_TransIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode Cbg_TransIndication_DL", err)
		}
	}
	if Cbg_TransIndication_ULPresent {
		ie.Cbg_TransIndication_UL = new(Phy_ParametersCommon_cbg_TransIndication_UL)
		if err = ie.Cbg_TransIndication_UL.Decode(r); err != nil {
			return utils.WrapError("Decode Cbg_TransIndication_UL", err)
		}
	}
	if Cbg_FlushIndication_DLPresent {
		ie.Cbg_FlushIndication_DL = new(Phy_ParametersCommon_cbg_FlushIndication_DL)
		if err = ie.Cbg_FlushIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode Cbg_FlushIndication_DL", err)
		}
	}
	if DynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent {
		ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL = new(Phy_ParametersCommon_dynamicHARQ_ACK_CodeB_CBG_Retx_DL)
		if err = ie.DynamicHARQ_ACK_CodeB_CBG_Retx_DL.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicHARQ_ACK_CodeB_CBG_Retx_DL", err)
		}
	}
	if RateMatchingResrcSetSemi_StaticPresent {
		ie.RateMatchingResrcSetSemi_Static = new(Phy_ParametersCommon_rateMatchingResrcSetSemi_Static)
		if err = ie.RateMatchingResrcSetSemi_Static.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatchingResrcSetSemi_Static", err)
		}
	}
	if RateMatchingResrcSetDynamicPresent {
		ie.RateMatchingResrcSetDynamic = new(Phy_ParametersCommon_rateMatchingResrcSetDynamic)
		if err = ie.RateMatchingResrcSetDynamic.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatchingResrcSetDynamic", err)
		}
	}
	if Bwp_SwitchingDelayPresent {
		ie.Bwp_SwitchingDelay = new(Phy_ParametersCommon_bwp_SwitchingDelay)
		if err = ie.Bwp_SwitchingDelay.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_SwitchingDelay", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 11 bits for 11 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dummy optional
			if DummyPresent {
				ie.Dummy = new(Phy_ParametersCommon_dummy)
				if err = ie.Dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			MaxNumberSearchSpacesPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RateMatchingCtrlResrcSetDynamicPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxLayersMIMO_IndicationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxNumberSearchSpaces optional
			if MaxNumberSearchSpacesPresent {
				ie.MaxNumberSearchSpaces = new(Phy_ParametersCommon_maxNumberSearchSpaces)
				if err = ie.MaxNumberSearchSpaces.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberSearchSpaces", err)
				}
			}
			// decode RateMatchingCtrlResrcSetDynamic optional
			if RateMatchingCtrlResrcSetDynamicPresent {
				ie.RateMatchingCtrlResrcSetDynamic = new(Phy_ParametersCommon_rateMatchingCtrlResrcSetDynamic)
				if err = ie.RateMatchingCtrlResrcSetDynamic.Decode(extReader); err != nil {
					return utils.WrapError("Decode RateMatchingCtrlResrcSetDynamic", err)
				}
			}
			// decode MaxLayersMIMO_Indication optional
			if MaxLayersMIMO_IndicationPresent {
				ie.MaxLayersMIMO_Indication = new(Phy_ParametersCommon_maxLayersMIMO_Indication)
				if err = ie.MaxLayersMIMO_Indication.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxLayersMIMO_Indication", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			SpCellPlacementPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpCellPlacement optional
			if SpCellPlacementPresent {
				ie.SpCellPlacement = new(CarrierAggregationVariant)
				if err = ie.SpCellPlacement.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpCellPlacement", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			TwoStepRACH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dci_Format1_2And0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MonitoringDCI_SameSearchSpace_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type2_CG_ReleaseDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type2_CG_ReleaseDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_ReleaseDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_ReleaseDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_TriggerStateNon_ActiveBWP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SeparateSMTC_InterIAB_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SeparateRACH_IAB_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_flexibleDL_SlotFormatDynamics_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dft_S_OFDM_WaveformUL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dci_25_AI_RNTI_Support_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			T_DeltaReceptionSupport_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GuardSymbolReportReception_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HarqACK_CB_SpatialBundlingPUCCH_Group_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CrossSlotScheduling_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberSRS_PosPathLossEstimateAllServingCells_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedCG_Periodicities_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ExtendedSPS_Periodicities_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookVariantsList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_RepetitionTypeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dci_DL_PriorityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dci_UL_PriorityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberPathlossRS_Update_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type2_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxTotalResourcesForAcrossFreqRanges_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HarqACK_separateMultiDCI_MultiTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HarqACK_jointMultiDCI_MultiTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Bwp_SwitchingMultiCCs_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TwoStepRACH_r16 optional
			if TwoStepRACH_r16Present {
				ie.TwoStepRACH_r16 = new(Phy_ParametersCommon_twoStepRACH_r16)
				if err = ie.TwoStepRACH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoStepRACH_r16", err)
				}
			}
			// decode Dci_Format1_2And0_2_r16 optional
			if Dci_Format1_2And0_2_r16Present {
				ie.Dci_Format1_2And0_2_r16 = new(Phy_ParametersCommon_dci_Format1_2And0_2_r16)
				if err = ie.Dci_Format1_2And0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dci_Format1_2And0_2_r16", err)
				}
			}
			// decode MonitoringDCI_SameSearchSpace_r16 optional
			if MonitoringDCI_SameSearchSpace_r16Present {
				ie.MonitoringDCI_SameSearchSpace_r16 = new(Phy_ParametersCommon_monitoringDCI_SameSearchSpace_r16)
				if err = ie.MonitoringDCI_SameSearchSpace_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MonitoringDCI_SameSearchSpace_r16", err)
				}
			}
			// decode Type2_CG_ReleaseDCI_0_1_r16 optional
			if Type2_CG_ReleaseDCI_0_1_r16Present {
				ie.Type2_CG_ReleaseDCI_0_1_r16 = new(Phy_ParametersCommon_type2_CG_ReleaseDCI_0_1_r16)
				if err = ie.Type2_CG_ReleaseDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type2_CG_ReleaseDCI_0_1_r16", err)
				}
			}
			// decode Type2_CG_ReleaseDCI_0_2_r16 optional
			if Type2_CG_ReleaseDCI_0_2_r16Present {
				ie.Type2_CG_ReleaseDCI_0_2_r16 = new(Phy_ParametersCommon_type2_CG_ReleaseDCI_0_2_r16)
				if err = ie.Type2_CG_ReleaseDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type2_CG_ReleaseDCI_0_2_r16", err)
				}
			}
			// decode Sps_ReleaseDCI_1_1_r16 optional
			if Sps_ReleaseDCI_1_1_r16Present {
				ie.Sps_ReleaseDCI_1_1_r16 = new(Phy_ParametersCommon_sps_ReleaseDCI_1_1_r16)
				if err = ie.Sps_ReleaseDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ReleaseDCI_1_1_r16", err)
				}
			}
			// decode Sps_ReleaseDCI_1_2_r16 optional
			if Sps_ReleaseDCI_1_2_r16Present {
				ie.Sps_ReleaseDCI_1_2_r16 = new(Phy_ParametersCommon_sps_ReleaseDCI_1_2_r16)
				if err = ie.Sps_ReleaseDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_ReleaseDCI_1_2_r16", err)
				}
			}
			// decode Csi_TriggerStateNon_ActiveBWP_r16 optional
			if Csi_TriggerStateNon_ActiveBWP_r16Present {
				ie.Csi_TriggerStateNon_ActiveBWP_r16 = new(Phy_ParametersCommon_csi_TriggerStateNon_ActiveBWP_r16)
				if err = ie.Csi_TriggerStateNon_ActiveBWP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_TriggerStateNon_ActiveBWP_r16", err)
				}
			}
			// decode SeparateSMTC_InterIAB_Support_r16 optional
			if SeparateSMTC_InterIAB_Support_r16Present {
				ie.SeparateSMTC_InterIAB_Support_r16 = new(Phy_ParametersCommon_separateSMTC_InterIAB_Support_r16)
				if err = ie.SeparateSMTC_InterIAB_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SeparateSMTC_InterIAB_Support_r16", err)
				}
			}
			// decode SeparateRACH_IAB_Support_r16 optional
			if SeparateRACH_IAB_Support_r16Present {
				ie.SeparateRACH_IAB_Support_r16 = new(Phy_ParametersCommon_separateRACH_IAB_Support_r16)
				if err = ie.SeparateRACH_IAB_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SeparateRACH_IAB_Support_r16", err)
				}
			}
			// decode Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 optional
			if Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16Present {
				ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 = new(Phy_ParametersCommon_ul_flexibleDL_SlotFormatSemiStatic_IAB_r16)
				if err = ie.Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_flexibleDL_SlotFormatSemiStatic_IAB_r16", err)
				}
			}
			// decode Ul_flexibleDL_SlotFormatDynamics_IAB_r16 optional
			if Ul_flexibleDL_SlotFormatDynamics_IAB_r16Present {
				ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16 = new(Phy_ParametersCommon_ul_flexibleDL_SlotFormatDynamics_IAB_r16)
				if err = ie.Ul_flexibleDL_SlotFormatDynamics_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_flexibleDL_SlotFormatDynamics_IAB_r16", err)
				}
			}
			// decode Dft_S_OFDM_WaveformUL_IAB_r16 optional
			if Dft_S_OFDM_WaveformUL_IAB_r16Present {
				ie.Dft_S_OFDM_WaveformUL_IAB_r16 = new(Phy_ParametersCommon_dft_S_OFDM_WaveformUL_IAB_r16)
				if err = ie.Dft_S_OFDM_WaveformUL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dft_S_OFDM_WaveformUL_IAB_r16", err)
				}
			}
			// decode Dci_25_AI_RNTI_Support_IAB_r16 optional
			if Dci_25_AI_RNTI_Support_IAB_r16Present {
				ie.Dci_25_AI_RNTI_Support_IAB_r16 = new(Phy_ParametersCommon_dci_25_AI_RNTI_Support_IAB_r16)
				if err = ie.Dci_25_AI_RNTI_Support_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dci_25_AI_RNTI_Support_IAB_r16", err)
				}
			}
			// decode T_DeltaReceptionSupport_IAB_r16 optional
			if T_DeltaReceptionSupport_IAB_r16Present {
				ie.T_DeltaReceptionSupport_IAB_r16 = new(Phy_ParametersCommon_t_DeltaReceptionSupport_IAB_r16)
				if err = ie.T_DeltaReceptionSupport_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode T_DeltaReceptionSupport_IAB_r16", err)
				}
			}
			// decode GuardSymbolReportReception_IAB_r16 optional
			if GuardSymbolReportReception_IAB_r16Present {
				ie.GuardSymbolReportReception_IAB_r16 = new(Phy_ParametersCommon_guardSymbolReportReception_IAB_r16)
				if err = ie.GuardSymbolReportReception_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode GuardSymbolReportReception_IAB_r16", err)
				}
			}
			// decode HarqACK_CB_SpatialBundlingPUCCH_Group_r16 optional
			if HarqACK_CB_SpatialBundlingPUCCH_Group_r16Present {
				ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16 = new(Phy_ParametersCommon_harqACK_CB_SpatialBundlingPUCCH_Group_r16)
				if err = ie.HarqACK_CB_SpatialBundlingPUCCH_Group_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HarqACK_CB_SpatialBundlingPUCCH_Group_r16", err)
				}
			}
			// decode CrossSlotScheduling_r16 optional
			if CrossSlotScheduling_r16Present {
				ie.CrossSlotScheduling_r16 = new(Phy_ParametersCommon_crossSlotScheduling_r16)
				if err = ie.CrossSlotScheduling_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CrossSlotScheduling_r16", err)
				}
			}
			// decode MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 optional
			if MaxNumberSRS_PosPathLossEstimateAllServingCells_r16Present {
				ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16 = new(Phy_ParametersCommon_maxNumberSRS_PosPathLossEstimateAllServingCells_r16)
				if err = ie.MaxNumberSRS_PosPathLossEstimateAllServingCells_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberSRS_PosPathLossEstimateAllServingCells_r16", err)
				}
			}
			// decode ExtendedCG_Periodicities_r16 optional
			if ExtendedCG_Periodicities_r16Present {
				ie.ExtendedCG_Periodicities_r16 = new(Phy_ParametersCommon_extendedCG_Periodicities_r16)
				if err = ie.ExtendedCG_Periodicities_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedCG_Periodicities_r16", err)
				}
			}
			// decode ExtendedSPS_Periodicities_r16 optional
			if ExtendedSPS_Periodicities_r16Present {
				ie.ExtendedSPS_Periodicities_r16 = new(Phy_ParametersCommon_extendedSPS_Periodicities_r16)
				if err = ie.ExtendedSPS_Periodicities_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ExtendedSPS_Periodicities_r16", err)
				}
			}
			// decode CodebookVariantsList_r16 optional
			if CodebookVariantsList_r16Present {
				ie.CodebookVariantsList_r16 = new(CodebookVariantsList_r16)
				if err = ie.CodebookVariantsList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookVariantsList_r16", err)
				}
			}
			// decode Pusch_RepetitionTypeA_r16 optional
			if Pusch_RepetitionTypeA_r16Present {
				ie.Pusch_RepetitionTypeA_r16 = new(Phy_ParametersCommon_pusch_RepetitionTypeA_r16)
				if err = ie.Pusch_RepetitionTypeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_RepetitionTypeA_r16", err)
				}
			}
			// decode Dci_DL_PriorityIndicator_r16 optional
			if Dci_DL_PriorityIndicator_r16Present {
				ie.Dci_DL_PriorityIndicator_r16 = new(Phy_ParametersCommon_dci_DL_PriorityIndicator_r16)
				if err = ie.Dci_DL_PriorityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dci_DL_PriorityIndicator_r16", err)
				}
			}
			// decode Dci_UL_PriorityIndicator_r16 optional
			if Dci_UL_PriorityIndicator_r16Present {
				ie.Dci_UL_PriorityIndicator_r16 = new(Phy_ParametersCommon_dci_UL_PriorityIndicator_r16)
				if err = ie.Dci_UL_PriorityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dci_UL_PriorityIndicator_r16", err)
				}
			}
			// decode MaxNumberPathlossRS_Update_r16 optional
			if MaxNumberPathlossRS_Update_r16Present {
				ie.MaxNumberPathlossRS_Update_r16 = new(Phy_ParametersCommon_maxNumberPathlossRS_Update_r16)
				if err = ie.MaxNumberPathlossRS_Update_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberPathlossRS_Update_r16", err)
				}
			}
			// decode Type2_HARQ_ACK_Codebook_r16 optional
			if Type2_HARQ_ACK_Codebook_r16Present {
				ie.Type2_HARQ_ACK_Codebook_r16 = new(Phy_ParametersCommon_type2_HARQ_ACK_Codebook_r16)
				if err = ie.Type2_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type2_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode MaxTotalResourcesForAcrossFreqRanges_r16 optional
			if MaxTotalResourcesForAcrossFreqRanges_r16Present {
				ie.MaxTotalResourcesForAcrossFreqRanges_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16)
				if err = ie.MaxTotalResourcesForAcrossFreqRanges_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxTotalResourcesForAcrossFreqRanges_r16", err)
				}
			}
			// decode HarqACK_separateMultiDCI_MultiTRP_r16 optional
			if HarqACK_separateMultiDCI_MultiTRP_r16Present {
				ie.HarqACK_separateMultiDCI_MultiTRP_r16 = new(Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16)
				if err = ie.HarqACK_separateMultiDCI_MultiTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HarqACK_separateMultiDCI_MultiTRP_r16", err)
				}
			}
			// decode HarqACK_jointMultiDCI_MultiTRP_r16 optional
			if HarqACK_jointMultiDCI_MultiTRP_r16Present {
				ie.HarqACK_jointMultiDCI_MultiTRP_r16 = new(Phy_ParametersCommon_harqACK_jointMultiDCI_MultiTRP_r16)
				if err = ie.HarqACK_jointMultiDCI_MultiTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HarqACK_jointMultiDCI_MultiTRP_r16", err)
				}
			}
			// decode Bwp_SwitchingMultiCCs_r16 optional
			if Bwp_SwitchingMultiCCs_r16Present {
				ie.Bwp_SwitchingMultiCCs_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16)
				if err = ie.Bwp_SwitchingMultiCCs_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Bwp_SwitchingMultiCCs_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			TargetSMTC_SCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportRepetitionZeroOffsetRV_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cbg_TransInOrderPUSCH_UL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TargetSMTC_SCG_r16 optional
			if TargetSMTC_SCG_r16Present {
				ie.TargetSMTC_SCG_r16 = new(Phy_ParametersCommon_targetSMTC_SCG_r16)
				if err = ie.TargetSMTC_SCG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode TargetSMTC_SCG_r16", err)
				}
			}
			// decode SupportRepetitionZeroOffsetRV_r16 optional
			if SupportRepetitionZeroOffsetRV_r16Present {
				ie.SupportRepetitionZeroOffsetRV_r16 = new(Phy_ParametersCommon_supportRepetitionZeroOffsetRV_r16)
				if err = ie.SupportRepetitionZeroOffsetRV_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportRepetitionZeroOffsetRV_r16", err)
				}
			}
			// decode Cbg_TransInOrderPUSCH_UL_r16 optional
			if Cbg_TransInOrderPUSCH_UL_r16Present {
				ie.Cbg_TransInOrderPUSCH_UL_r16 = new(Phy_ParametersCommon_cbg_TransInOrderPUSCH_UL_r16)
				if err = ie.Cbg_TransInOrderPUSCH_UL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cbg_TransInOrderPUSCH_UL_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Bwp_SwitchingMultiDormancyCCs_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Bwp_SwitchingMultiDormancyCCs_r16 optional
			if Bwp_SwitchingMultiDormancyCCs_r16Present {
				ie.Bwp_SwitchingMultiDormancyCCs_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16)
				if err = ie.Bwp_SwitchingMultiDormancyCCs_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Bwp_SwitchingMultiDormancyCCs_r16", err)
				}
			}
			// decode SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 optional
			if SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16Present {
				ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 = new(Phy_ParametersCommon_supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16)
				if err = ie.SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16", err)
				}
			}
			// decode Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 optional
			if Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16Present {
				ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 = new(Phy_ParametersCommon_pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16)
				if err = ie.Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			NewBeamIdentifications2PortCSI_RS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PathlossEstimation2PortCSI_RS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode NewBeamIdentifications2PortCSI_RS_r16 optional
			if NewBeamIdentifications2PortCSI_RS_r16Present {
				ie.NewBeamIdentifications2PortCSI_RS_r16 = new(Phy_ParametersCommon_newBeamIdentifications2PortCSI_RS_r16)
				if err = ie.NewBeamIdentifications2PortCSI_RS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode NewBeamIdentifications2PortCSI_RS_r16", err)
				}
			}
			// decode PathlossEstimation2PortCSI_RS_r16 optional
			if PathlossEstimation2PortCSI_RS_r16Present {
				ie.PathlossEstimation2PortCSI_RS_r16 = new(Phy_ParametersCommon_pathlossEstimation2PortCSI_RS_r16)
				if err = ie.PathlossEstimation2PortCSI_RS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PathlossEstimation2PortCSI_RS_r16", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 optional
			if Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16Present {
				ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 = new(Phy_ParametersCommon_mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16)
				if err = ie.Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16", err)
				}
			}
		}
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			GuardSymbolReportReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Restricted_IAB_DU_BeamReception_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Recommended_IAB_MT_BeamTransmission_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Case6_TimingAlignmentReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Case7_TimingAlignmentReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_tx_PowerAdjustment_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Desired_ul_tx_PowerAdjustment_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Fdm_SoftResourceAvailability_DynamicIndication_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Updated_T_DeltaRangeRecption_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SlotBasedDynamicPUCCH_Rep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_HARQ_ACK_Deferral_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_commonUpdate_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PDCCH_singleSpan_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedActivatedPRS_ProcessingWindow_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_TimeDomainAllocationExtension_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode GuardSymbolReportReception_IAB_r17 optional
			if GuardSymbolReportReception_IAB_r17Present {
				ie.GuardSymbolReportReception_IAB_r17 = new(Phy_ParametersCommon_guardSymbolReportReception_IAB_r17)
				if err = ie.GuardSymbolReportReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GuardSymbolReportReception_IAB_r17", err)
				}
			}
			// decode Restricted_IAB_DU_BeamReception_r17 optional
			if Restricted_IAB_DU_BeamReception_r17Present {
				ie.Restricted_IAB_DU_BeamReception_r17 = new(Phy_ParametersCommon_restricted_IAB_DU_BeamReception_r17)
				if err = ie.Restricted_IAB_DU_BeamReception_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Restricted_IAB_DU_BeamReception_r17", err)
				}
			}
			// decode Recommended_IAB_MT_BeamTransmission_r17 optional
			if Recommended_IAB_MT_BeamTransmission_r17Present {
				ie.Recommended_IAB_MT_BeamTransmission_r17 = new(Phy_ParametersCommon_recommended_IAB_MT_BeamTransmission_r17)
				if err = ie.Recommended_IAB_MT_BeamTransmission_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Recommended_IAB_MT_BeamTransmission_r17", err)
				}
			}
			// decode Case6_TimingAlignmentReception_IAB_r17 optional
			if Case6_TimingAlignmentReception_IAB_r17Present {
				ie.Case6_TimingAlignmentReception_IAB_r17 = new(Phy_ParametersCommon_case6_TimingAlignmentReception_IAB_r17)
				if err = ie.Case6_TimingAlignmentReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Case6_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// decode Case7_TimingAlignmentReception_IAB_r17 optional
			if Case7_TimingAlignmentReception_IAB_r17Present {
				ie.Case7_TimingAlignmentReception_IAB_r17 = new(Phy_ParametersCommon_case7_TimingAlignmentReception_IAB_r17)
				if err = ie.Case7_TimingAlignmentReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Case7_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// decode Dl_tx_PowerAdjustment_IAB_r17 optional
			if Dl_tx_PowerAdjustment_IAB_r17Present {
				ie.Dl_tx_PowerAdjustment_IAB_r17 = new(Phy_ParametersCommon_dl_tx_PowerAdjustment_IAB_r17)
				if err = ie.Dl_tx_PowerAdjustment_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_tx_PowerAdjustment_IAB_r17", err)
				}
			}
			// decode Desired_ul_tx_PowerAdjustment_r17 optional
			if Desired_ul_tx_PowerAdjustment_r17Present {
				ie.Desired_ul_tx_PowerAdjustment_r17 = new(Phy_ParametersCommon_desired_ul_tx_PowerAdjustment_r17)
				if err = ie.Desired_ul_tx_PowerAdjustment_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Desired_ul_tx_PowerAdjustment_r17", err)
				}
			}
			// decode Fdm_SoftResourceAvailability_DynamicIndication_r17 optional
			if Fdm_SoftResourceAvailability_DynamicIndication_r17Present {
				ie.Fdm_SoftResourceAvailability_DynamicIndication_r17 = new(Phy_ParametersCommon_fdm_SoftResourceAvailability_DynamicIndication_r17)
				if err = ie.Fdm_SoftResourceAvailability_DynamicIndication_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Fdm_SoftResourceAvailability_DynamicIndication_r17", err)
				}
			}
			// decode Updated_T_DeltaRangeRecption_r17 optional
			if Updated_T_DeltaRangeRecption_r17Present {
				ie.Updated_T_DeltaRangeRecption_r17 = new(Phy_ParametersCommon_updated_T_DeltaRangeRecption_r17)
				if err = ie.Updated_T_DeltaRangeRecption_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Updated_T_DeltaRangeRecption_r17", err)
				}
			}
			// decode SlotBasedDynamicPUCCH_Rep_r17 optional
			if SlotBasedDynamicPUCCH_Rep_r17Present {
				ie.SlotBasedDynamicPUCCH_Rep_r17 = new(Phy_ParametersCommon_slotBasedDynamicPUCCH_Rep_r17)
				if err = ie.SlotBasedDynamicPUCCH_Rep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SlotBasedDynamicPUCCH_Rep_r17", err)
				}
			}
			// decode Sps_HARQ_ACK_Deferral_r17 optional
			if Sps_HARQ_ACK_Deferral_r17Present {
				ie.Sps_HARQ_ACK_Deferral_r17 = new(Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17)
				if err = ie.Sps_HARQ_ACK_Deferral_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_HARQ_ACK_Deferral_r17", err)
				}
			}
			// decode UnifiedJointTCI_commonUpdate_r17 optional
			if UnifiedJointTCI_commonUpdate_r17Present {
				var tmp_int_UnifiedJointTCI_commonUpdate_r17 int64
				if tmp_int_UnifiedJointTCI_commonUpdate_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_commonUpdate_r17", err)
				}
				ie.UnifiedJointTCI_commonUpdate_r17 = &tmp_int_UnifiedJointTCI_commonUpdate_r17
			}
			// decode MTRP_PDCCH_singleSpan_r17 optional
			if MTRP_PDCCH_singleSpan_r17Present {
				ie.MTRP_PDCCH_singleSpan_r17 = new(Phy_ParametersCommon_mTRP_PDCCH_singleSpan_r17)
				if err = ie.MTRP_PDCCH_singleSpan_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PDCCH_singleSpan_r17", err)
				}
			}
			// decode SupportedActivatedPRS_ProcessingWindow_r17 optional
			if SupportedActivatedPRS_ProcessingWindow_r17Present {
				ie.SupportedActivatedPRS_ProcessingWindow_r17 = new(Phy_ParametersCommon_supportedActivatedPRS_ProcessingWindow_r17)
				if err = ie.SupportedActivatedPRS_ProcessingWindow_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedActivatedPRS_ProcessingWindow_r17", err)
				}
			}
			// decode Cg_TimeDomainAllocationExtension_r17 optional
			if Cg_TimeDomainAllocationExtension_r17Present {
				ie.Cg_TimeDomainAllocationExtension_r17 = new(Phy_ParametersCommon_cg_TimeDomainAllocationExtension_r17)
				if err = ie.Cg_TimeDomainAllocationExtension_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_TimeDomainAllocationExtension_r17", err)
				}
			}
		}
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DirectionalCollisionDC_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 optional
			if Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17Present {
				ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 = new(Phy_ParametersCommon_ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17)
				if err = ie.Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// decode DirectionalCollisionDC_IAB_r17 optional
			if DirectionalCollisionDC_IAB_r17Present {
				ie.DirectionalCollisionDC_IAB_r17 = new(Phy_ParametersCommon_directionalCollisionDC_IAB_r17)
				if err = ie.DirectionalCollisionDC_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DirectionalCollisionDC_IAB_r17", err)
				}
			}
		}
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PriorityIndicatorInDCI_Multicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PriorityIndicatorInDCI_SPS_Multicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_AdditionalRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_Repetition_CG_SDT_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PriorityIndicatorInDCI_Multicast_r17 optional
			if PriorityIndicatorInDCI_Multicast_r17Present {
				ie.PriorityIndicatorInDCI_Multicast_r17 = new(Phy_ParametersCommon_priorityIndicatorInDCI_Multicast_r17)
				if err = ie.PriorityIndicatorInDCI_Multicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorInDCI_Multicast_r17", err)
				}
			}
			// decode PriorityIndicatorInDCI_SPS_Multicast_r17 optional
			if PriorityIndicatorInDCI_SPS_Multicast_r17Present {
				ie.PriorityIndicatorInDCI_SPS_Multicast_r17 = new(Phy_ParametersCommon_priorityIndicatorInDCI_SPS_Multicast_r17)
				if err = ie.PriorityIndicatorInDCI_SPS_Multicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PriorityIndicatorInDCI_SPS_Multicast_r17", err)
				}
			}
			// decode TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 optional
			if TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17Present {
				ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 = new(Phy_ParametersCommon_twoHARQ_ACK_CodebookForUnicastAndMulticast_r17)
				if err = ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17", err)
				}
			}
			// decode MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 optional
			if MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17Present {
				ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 = new(Phy_ParametersCommon_multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17)
				if err = ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17", err)
				}
			}
			// decode Srs_AdditionalRepetition_r17 optional
			if Srs_AdditionalRepetition_r17Present {
				ie.Srs_AdditionalRepetition_r17 = new(Phy_ParametersCommon_srs_AdditionalRepetition_r17)
				if err = ie.Srs_AdditionalRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_AdditionalRepetition_r17", err)
				}
			}
			// decode Pusch_Repetition_CG_SDT_r17 optional
			if Pusch_Repetition_CG_SDT_r17Present {
				ie.Pusch_Repetition_CG_SDT_r17 = new(Phy_ParametersCommon_pusch_Repetition_CG_SDT_r17)
				if err = ie.Pusch_Repetition_CG_SDT_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_Repetition_CG_SDT_r17", err)
				}
			}
		}
	}
	return nil
}
