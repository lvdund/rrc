package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandNR struct {
	BandNR                                                         FreqBandIndicatorNR                                                    `madatory`
	ModifiedMPR_Behaviour                                          *aper.BitString                                                        `lb:8,ub:8,optional`
	Mimo_ParametersPerBand                                         *MIMO_ParametersPerBand                                                `optional`
	ExtendedCP                                                     *BandNR_extendedCP                                                     `optional`
	MultipleTCI                                                    *BandNR_multipleTCI                                                    `optional`
	Bwp_WithoutRestriction                                         *BandNR_bwp_WithoutRestriction                                         `optional`
	Bwp_SameNumerology                                             *BandNR_bwp_SameNumerology                                             `optional`
	Bwp_DiffNumerology                                             *BandNR_bwp_DiffNumerology                                             `optional`
	CrossCarrierScheduling_SameSCS                                 *BandNR_crossCarrierScheduling_SameSCS                                 `optional`
	Pdsch_256QAM_FR2                                               *BandNR_pdsch_256QAM_FR2                                               `optional`
	Pusch_256QAM                                                   *BandNR_pusch_256QAM                                                   `optional`
	Ue_PowerClass                                                  *BandNR_ue_PowerClass                                                  `optional`
	RateMatchingLTE_CRS                                            *BandNR_rateMatchingLTE_CRS                                            `optional`
	ChannelBWs_DL                                                  *BandNR_channelBWs_DL                                                  `lb:10,ub:10,optional`
	ChannelBWs_UL                                                  *BandNR_channelBWs_UL                                                  `lb:10,ub:10,optional`
	MaxUplinkDutyCycle_PC2_FR1                                     *BandNR_maxUplinkDutyCycle_PC2_FR1                                     `optional,ext-1`
	Pucch_SpatialRelInfoMAC_CE                                     *BandNR_pucch_SpatialRelInfoMAC_CE                                     `optional,ext-2`
	PowerBoosting_pi2BPSK                                          *BandNR_powerBoosting_pi2BPSK                                          `optional,ext-2`
	MaxUplinkDutyCycle_FR2                                         *BandNR_maxUplinkDutyCycle_FR2                                         `optional,ext-3`
	ChannelBWs_DL_v1590                                            *BandNR_channelBWs_DL_v1590                                            `lb:16,ub:16,optional,ext-4`
	ChannelBWs_UL_v1590                                            *BandNR_channelBWs_UL_v1590                                            `lb:16,ub:16,optional,ext-4`
	AsymmetricBandwidthCombinationSet                              *aper.BitString                                                        `lb:1,ub:32,optional,ext-5`
	SharedSpectrumChAccessParamsPerBand_r16                        *SharedSpectrumChAccessParamsPerBand_r16                               `optional,ext-6`
	CancelOverlappingPUSCH_r16                                     *BandNR_cancelOverlappingPUSCH_r16                                     `optional,ext-6`
	MultipleRateMatchingEUTRA_CRS_r16                              *BandNR_multipleRateMatchingEUTRA_CRS_r16                              `lb:2,ub:6,optional,ext-6`
	OverlapRateMatchingEUTRA_CRS_r16                               *BandNR_overlapRateMatchingEUTRA_CRS_r16                               `optional,ext-6`
	Pdsch_MappingTypeB_Alt_r16                                     *BandNR_pdsch_MappingTypeB_Alt_r16                                     `optional,ext-6`
	OneSlotPeriodicTRS_r16                                         *BandNR_oneSlotPeriodicTRS_r16                                         `optional,ext-6`
	Olpc_SRS_Pos_r16                                               *OLPC_SRS_Pos_r16                                                      `optional,ext-6`
	SpatialRelationsSRS_Pos_r16                                    *SpatialRelationsSRS_Pos_r16                                           `optional,ext-6`
	SimulSRS_MIMO_TransWithinBand_r16                              *BandNR_simulSRS_MIMO_TransWithinBand_r16                              `optional,ext-6`
	ChannelBW_DL_IAB_r16                                           *BandNR_channelBW_DL_IAB_r16                                           `optional,ext-6`
	ChannelBW_UL_IAB_r16                                           *BandNR_channelBW_UL_IAB_r16                                           `optional,ext-6`
	RasterShift7dot5_IAB_r16                                       *BandNR_rasterShift7dot5_IAB_r16                                       `optional,ext-6`
	Ue_PowerClass_v1610                                            *BandNR_ue_PowerClass_v1610                                            `optional,ext-6`
	CondHandover_r16                                               *BandNR_condHandover_r16                                               `optional,ext-6`
	CondHandoverFailure_r16                                        *BandNR_condHandoverFailure_r16                                        `optional,ext-6`
	CondHandoverTwoTriggerEvents_r16                               *BandNR_condHandoverTwoTriggerEvents_r16                               `optional,ext-6`
	CondPSCellChange_r16                                           *BandNR_condPSCellChange_r16                                           `optional,ext-6`
	CondPSCellChangeTwoTriggerEvents_r16                           *BandNR_condPSCellChangeTwoTriggerEvents_r16                           `optional,ext-6`
	Mpr_PowerBoost_FR2_r16                                         *BandNR_mpr_PowerBoost_FR2_r16                                         `optional,ext-6`
	ActiveConfiguredGrant_r16                                      *BandNR_activeConfiguredGrant_r16                                      `lb:2,ub:32,optional,ext-6`
	JointReleaseConfiguredGrantType2_r16                           *BandNR_jointReleaseConfiguredGrantType2_r16                           `optional,ext-6`
	Sps_r16                                                        *BandNR_sps_r16                                                        `lb:1,ub:8,optional,ext-6`
	JointReleaseSPS_r16                                            *BandNR_jointReleaseSPS_r16                                            `optional,ext-6`
	SimulSRS_TransWithinBand_r16                                   *BandNR_simulSRS_TransWithinBand_r16                                   `optional,ext-6`
	Trs_AdditionalBandwidth_r16                                    *BandNR_trs_AdditionalBandwidth_r16                                    `optional,ext-6`
	HandoverIntraF_IAB_r16                                         *BandNR_handoverIntraF_IAB_r16                                         `optional,ext-6`
	SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16                     *SimulSRS_ForAntennaSwitching_r16                                      `optional,ext-7`
	SharedSpectrumChAccessParamsPerBand_v1630                      *SharedSpectrumChAccessParamsPerBand_v1630                             `optional,ext-7`
	HandoverUTRA_FDD_r16                                           *BandNR_handoverUTRA_FDD_r16                                           `optional,ext-8`
	EnhancedUL_TransientPeriod_r16                                 *BandNR_enhancedUL_TransientPeriod_r16                                 `optional,ext-8`
	SharedSpectrumChAccessParamsPerBand_v1640                      *SharedSpectrumChAccessParamsPerBand_v1640                             `optional,ext-8`
	Type1_PUSCH_RepetitionMultiSlots_v1650                         *BandNR_type1_PUSCH_RepetitionMultiSlots_v1650                         `optional,ext-9`
	Type2_PUSCH_RepetitionMultiSlots_v1650                         *BandNR_type2_PUSCH_RepetitionMultiSlots_v1650                         `optional,ext-9`
	Pusch_RepetitionMultiSlots_v1650                               *BandNR_pusch_RepetitionMultiSlots_v1650                               `optional,ext-9`
	ConfiguredUL_GrantType1_v1650                                  *BandNR_configuredUL_GrantType1_v1650                                  `optional,ext-9`
	ConfiguredUL_GrantType2_v1650                                  *BandNR_configuredUL_GrantType2_v1650                                  `optional,ext-9`
	SharedSpectrumChAccessParamsPerBand_v1650                      *SharedSpectrumChAccessParamsPerBand_v1650                             `optional,ext-9`
	EnhancedSkipUplinkTxConfigured_v1660                           *BandNR_enhancedSkipUplinkTxConfigured_v1660                           `optional,ext-10`
	EnhancedSkipUplinkTxDynamic_v1660                              *BandNR_enhancedSkipUplinkTxDynamic_v1660                              `optional,ext-10`
	MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16                         *BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16                         `optional,ext-11`
	TxDiversity_r16                                                *BandNR_txDiversity_r16                                                `optional,ext-11`
	Pdsch_1024QAM_FR1_r17                                          *BandNR_pdsch_1024QAM_FR1_r17                                          `optional,ext-12`
	Ue_PowerClass_v1700                                            *BandNR_ue_PowerClass_v1700                                            `optional,ext-12`
	Fr2_2_AccessParamsPerBand_r17                                  *FR2_2_AccessParamsPerBand_r17                                         `optional,ext-12`
	Rlm_Relaxation_r17                                             *BandNR_rlm_Relaxation_r17                                             `optional,ext-12`
	Bfd_Relaxation_r17                                             *BandNR_bfd_Relaxation_r17                                             `optional,ext-12`
	Cg_SDT_r17                                                     *BandNR_cg_SDT_r17                                                     `optional,ext-12`
	LocationBasedCondHandover_r17                                  *BandNR_locationBasedCondHandover_r17                                  `optional,ext-12`
	TimeBasedCondHandover_r17                                      *BandNR_timeBasedCondHandover_r17                                      `optional,ext-12`
	EventA4BasedCondHandover_r17                                   *BandNR_eventA4BasedCondHandover_r17                                   `optional,ext-12`
	Mn_InitiatedCondPSCellChangeNRDC_r17                           *BandNR_mn_InitiatedCondPSCellChangeNRDC_r17                           `optional,ext-12`
	Sn_InitiatedCondPSCellChangeNRDC_r17                           *BandNR_sn_InitiatedCondPSCellChangeNRDC_r17                           `optional,ext-12`
	Pdcch_SkippingWithoutSSSG_r17                                  *BandNR_pdcch_SkippingWithoutSSSG_r17                                  `optional,ext-12`
	Sssg_Switching_1BitInd_r17                                     *BandNR_sssg_Switching_1BitInd_r17                                     `optional,ext-12`
	Sssg_Switching_2BitInd_r17                                     *BandNR_sssg_Switching_2BitInd_r17                                     `optional,ext-12`
	Pdcch_SkippingWithSSSG_r17                                     *BandNR_pdcch_SkippingWithSSSG_r17                                     `optional,ext-12`
	SearchSpaceSetGrp_switchCap2_r17                               *BandNR_searchSpaceSetGrp_switchCap2_r17                               `optional,ext-12`
	UplinkPreCompensation_r17                                      *BandNR_uplinkPreCompensation_r17                                      `optional,ext-12`
	Uplink_TA_Reporting_r17                                        *BandNR_uplink_TA_Reporting_r17                                        `optional,ext-12`
	Max_HARQ_ProcessNumber_r17                                     *BandNR_max_HARQ_ProcessNumber_r17                                     `optional,ext-12`
	Type2_HARQ_Codebook_r17                                        *BandNR_type2_HARQ_Codebook_r17                                        `optional,ext-12`
	Type1_HARQ_Codebook_r17                                        *BandNR_type1_HARQ_Codebook_r17                                        `optional,ext-12`
	Type3_HARQ_Codebook_r17                                        *BandNR_type3_HARQ_Codebook_r17                                        `optional,ext-12`
	Ue_specific_K_Offset_r17                                       *BandNR_ue_specific_K_Offset_r17                                       `optional,ext-12`
	MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      *BandNR_multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      `optional,ext-12`
	MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      *BandNR_multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      `optional,ext-12`
	ParallelPRS_MeasRRC_Inactive_r17                               *BandNR_parallelPRS_MeasRRC_Inactive_r17                               `optional,ext-12`
	Nr_UE_TxTEG_ID_MaxSupport_r17                                  *BandNR_nr_UE_TxTEG_ID_MaxSupport_r17                                  `optional,ext-12`
	Prs_ProcessingRRC_Inactive_r17                                 *BandNR_prs_ProcessingRRC_Inactive_r17                                 `optional,ext-12`
	Prs_ProcessingWindowType1A_r17                                 *BandNR_prs_ProcessingWindowType1A_r17                                 `optional,ext-12`
	Prs_ProcessingWindowType1B_r17                                 *BandNR_prs_ProcessingWindowType1B_r17                                 `optional,ext-12`
	Prs_ProcessingWindowType2_r17                                  *BandNR_prs_ProcessingWindowType2_r17                                  `optional,ext-12`
	Srs_AllPosResourcesRRC_Inactive_r17                            *SRS_AllPosResourcesRRC_Inactive_r17                                   `optional,ext-12`
	Olpc_SRS_PosRRC_Inactive_r17                                   *OLPC_SRS_Pos_r16                                                      `optional,ext-12`
	SpatialRelationsSRS_PosRRC_Inactive_r17                        *SpatialRelationsSRS_Pos_r16                                           `optional,ext-12`
	MaxNumberPUSCH_TypeA_Repetition_r17                            *BandNR_maxNumberPUSCH_TypeA_Repetition_r17                            `optional,ext-12`
	PuschTypeA_RepetitionsAvailSlot_r17                            *BandNR_puschTypeA_RepetitionsAvailSlot_r17                            `optional,ext-12`
	Tb_ProcessingMultiSlotPUSCH_r17                                *BandNR_tb_ProcessingMultiSlotPUSCH_r17                                `optional,ext-12`
	Tb_ProcessingRepMultiSlotPUSCH_r17                             *BandNR_tb_ProcessingRepMultiSlotPUSCH_r17                             `optional,ext-12`
	MaxDurationDMRS_Bundling_r17                                   *BandNR_maxDurationDMRS_Bundling_r17                                   `optional,ext-12`
	Pusch_RepetitionMsg3_r17                                       *BandNR_pusch_RepetitionMsg3_r17                                       `optional,ext-12`
	SharedSpectrumChAccessParamsPerBand_v1710                      *SharedSpectrumChAccessParamsPerBand_v1710                             `optional,ext-12`
	ParallelMeasurementWithoutRestriction_r17                      *BandNR_parallelMeasurementWithoutRestriction_r17                      `optional,ext-12`
	MaxNumber_NGSO_SatellitesWithinOneSMTC_r17                     *BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17                     `optional,ext-12`
	K1_RangeExtension_r17                                          *BandNR_k1_RangeExtension_r17                                          `optional,ext-12`
	AperiodicCSI_RS_FastScellActivation_r17                        *BandNR_aperiodicCSI_RS_FastScellActivation_r17                        `optional,ext-12`
	AperiodicCSI_RS_AdditionalBandwidth_r17                        *BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17                        `optional,ext-12`
	Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17                         *BandNR_bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17                         `optional,ext-12`
	HalfDuplexFDD_TypeA_RedCap_r17                                 *BandNR_halfDuplexFDD_TypeA_RedCap_r17                                 `optional,ext-12`
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17                   *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17                          `optional,ext-12`
	ChannelBWs_DL_SCS_480kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	ChannelBWs_UL_SCS_480kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	ChannelBWs_DL_SCS_960kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	ChannelBWs_UL_SCS_960kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	Ul_GapFR2_r17                                                  *BandNR_ul_GapFR2_r17                                                  `optional,ext-12`
	OneShotHARQ_feedbackTriggeredByDCI_1_2_r17                     *BandNR_oneShotHARQ_feedbackTriggeredByDCI_1_2_r17                     `optional,ext-12`
	OneShotHARQ_feedbackPhy_Priority_r17                           *BandNR_oneShotHARQ_feedbackPhy_Priority_r17                           `optional,ext-12`
	EnhancedType3_HARQ_CodebookFeedback_r17                        *BandNR_enhancedType3_HARQ_CodebookFeedback_r17                        `optional,ext-12`
	TriggeredHARQ_CodebookRetx_r17                                 *BandNR_triggeredHARQ_CodebookRetx_r17                                 `optional,ext-12`
	Ue_OneShotUL_TimingAdj_r17                                     *BandNR_ue_OneShotUL_TimingAdj_r17                                     `optional,ext-13`
	Pucch_Repetition_F0_2_r17                                      *BandNR_pucch_Repetition_F0_2_r17                                      `optional,ext-13`
	Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17                *BandNR_cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17                `optional,ext-13`
	Mux_HARQ_ACK_DiffPriorities_r17                                *BandNR_mux_HARQ_ACK_DiffPriorities_r17                                `optional,ext-13`
	Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17                     *BandNR_ta_BasedPDC_NTN_SharedSpectrumChAccess_r17                     `optional,ext-13`
	Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17               *BandNR_ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17               `optional,ext-13`
	MaxNumberG_RNTI_r17                                            *int64                                                                 `lb:2,ub:8,optional,ext-13`
	DynamicMulticastDCI_Format4_2_r17                              *BandNR_dynamicMulticastDCI_Format4_2_r17                              `optional,ext-13`
	MaxModulationOrderForMulticast_r17                             *BandNR_maxModulationOrderForMulticast_r17                             `optional,ext-13`
	DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 *BandNR_dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 `optional,ext-13`
	DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17   *BandNR_dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17   `optional,ext-13`
	Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17               *BandNR_nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17               `optional,ext-13`
	Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17           *BandNR_ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17           `optional,ext-13`
	MaxNumberG_CS_RNTI_r17                                         *int64                                                                 `lb:2,ub:8,optional,ext-13`
	Re_LevelRateMatchingForMulticast_r17                           *BandNR_re_LevelRateMatchingForMulticast_r17                           `optional,ext-13`
	Pdsch_1024QAM_2MIMO_FR1_r17                                    *BandNR_pdsch_1024QAM_2MIMO_FR1_r17                                    `optional,ext-13`
	Prs_MeasurementWithoutMG_r17                                   *BandNR_prs_MeasurementWithoutMG_r17                                   `optional,ext-13`
	MaxNumber_LEO_SatellitesPerCarrier_r17                         *int64                                                                 `lb:3,ub:4,optional,ext-13`
	Prs_ProcessingCapabilityOutsideMGinPPW_r17                     []PRS_ProcessingCapabilityOutsideMGinPPWperType_r17                    `lb:1,ub:3,optional,ext-13`
	Srs_SemiPersistent_PosResourcesRRC_Inactive_r17                *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17                `optional,ext-13`
	ChannelBWs_DL_SCS_120kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-13`
	ChannelBWs_UL_SCS_120kHz_FR2_2_r17                             *aper.BitString                                                        `lb:8,ub:8,optional,ext-13`
	Dmrs_BundlingPUSCH_RepTypeA_r17                                *BandNR_dmrs_BundlingPUSCH_RepTypeA_r17                                `optional,ext-14`
	Dmrs_BundlingPUSCH_RepTypeB_r17                                *BandNR_dmrs_BundlingPUSCH_RepTypeB_r17                                `optional,ext-14`
	Dmrs_BundlingPUSCH_multiSlot_r17                               *BandNR_dmrs_BundlingPUSCH_multiSlot_r17                               `optional,ext-14`
	Dmrs_BundlingPUCCH_Rep_r17                                     *BandNR_dmrs_BundlingPUCCH_Rep_r17                                     `optional,ext-14`
	InterSlotFreqHopInterSlotBundlingPUSCH_r17                     *BandNR_interSlotFreqHopInterSlotBundlingPUSCH_r17                     `optional,ext-14`
	InterSlotFreqHopPUCCH_r17                                      *BandNR_interSlotFreqHopPUCCH_r17                                      `optional,ext-14`
	Dmrs_BundlingRestart_r17                                       *BandNR_dmrs_BundlingRestart_r17                                       `optional,ext-14`
	Dmrs_BundlingNonBackToBackTX_r17                               *BandNR_dmrs_BundlingNonBackToBackTX_r17                               `optional,ext-14`
}

func (ie *BandNR) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.MaxUplinkDutyCycle_PC2_FR1 != nil || ie.Pucch_SpatialRelInfoMAC_CE != nil || ie.PowerBoosting_pi2BPSK != nil || ie.MaxUplinkDutyCycle_FR2 != nil || ie.ChannelBWs_DL_v1590 != nil || ie.ChannelBWs_UL_v1590 != nil || ie.AsymmetricBandwidthCombinationSet != nil || ie.SharedSpectrumChAccessParamsPerBand_r16 != nil || ie.CancelOverlappingPUSCH_r16 != nil || ie.MultipleRateMatchingEUTRA_CRS_r16 != nil || ie.OverlapRateMatchingEUTRA_CRS_r16 != nil || ie.Pdsch_MappingTypeB_Alt_r16 != nil || ie.OneSlotPeriodicTRS_r16 != nil || ie.Olpc_SRS_Pos_r16 != nil || ie.SpatialRelationsSRS_Pos_r16 != nil || ie.SimulSRS_MIMO_TransWithinBand_r16 != nil || ie.ChannelBW_DL_IAB_r16 != nil || ie.ChannelBW_UL_IAB_r16 != nil || ie.RasterShift7dot5_IAB_r16 != nil || ie.Ue_PowerClass_v1610 != nil || ie.CondHandover_r16 != nil || ie.CondHandoverFailure_r16 != nil || ie.CondHandoverTwoTriggerEvents_r16 != nil || ie.CondPSCellChange_r16 != nil || ie.CondPSCellChangeTwoTriggerEvents_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil || ie.ActiveConfiguredGrant_r16 != nil || ie.JointReleaseConfiguredGrantType2_r16 != nil || ie.Sps_r16 != nil || ie.JointReleaseSPS_r16 != nil || ie.SimulSRS_TransWithinBand_r16 != nil || ie.Trs_AdditionalBandwidth_r16 != nil || ie.HandoverIntraF_IAB_r16 != nil || ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil || ie.HandoverUTRA_FDD_r16 != nil || ie.EnhancedUL_TransientPeriod_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil || ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Pusch_RepetitionMultiSlots_v1650 != nil || ie.ConfiguredUL_GrantType1_v1650 != nil || ie.ConfiguredUL_GrantType2_v1650 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil || ie.EnhancedSkipUplinkTxConfigured_v1660 != nil || ie.EnhancedSkipUplinkTxDynamic_v1660 != nil || ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil || ie.TxDiversity_r16 != nil || ie.Pdsch_1024QAM_FR1_r17 != nil || ie.Ue_PowerClass_v1700 != nil || ie.Fr2_2_AccessParamsPerBand_r17 != nil || ie.Rlm_Relaxation_r17 != nil || ie.Bfd_Relaxation_r17 != nil || ie.Cg_SDT_r17 != nil || ie.LocationBasedCondHandover_r17 != nil || ie.TimeBasedCondHandover_r17 != nil || ie.EventA4BasedCondHandover_r17 != nil || ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Pdcch_SkippingWithoutSSSG_r17 != nil || ie.Sssg_Switching_1BitInd_r17 != nil || ie.Sssg_Switching_2BitInd_r17 != nil || ie.Pdcch_SkippingWithSSSG_r17 != nil || ie.SearchSpaceSetGrp_switchCap2_r17 != nil || ie.UplinkPreCompensation_r17 != nil || ie.Uplink_TA_Reporting_r17 != nil || ie.Max_HARQ_ProcessNumber_r17 != nil || ie.Type2_HARQ_Codebook_r17 != nil || ie.Type1_HARQ_Codebook_r17 != nil || ie.Type3_HARQ_Codebook_r17 != nil || ie.Ue_specific_K_Offset_r17 != nil || ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.ParallelPRS_MeasRRC_Inactive_r17 != nil || ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil || ie.Prs_ProcessingRRC_Inactive_r17 != nil || ie.Prs_ProcessingWindowType1A_r17 != nil || ie.Prs_ProcessingWindowType1B_r17 != nil || ie.Prs_ProcessingWindowType2_r17 != nil || ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil || ie.Olpc_SRS_PosRRC_Inactive_r17 != nil || ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil || ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil || ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil || ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil || ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil || ie.MaxDurationDMRS_Bundling_r17 != nil || ie.Pusch_RepetitionMsg3_r17 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil || ie.ParallelMeasurementWithoutRestriction_r17 != nil || ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil || ie.K1_RangeExtension_r17 != nil || ie.AperiodicCSI_RS_FastScellActivation_r17 != nil || ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil || ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil || ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil || ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil || ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil || ie.Ul_GapFR2_r17 != nil || ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil || ie.OneShotHARQ_feedbackPhy_Priority_r17 != nil || ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil || ie.TriggeredHARQ_CodebookRetx_r17 != nil || ie.Ue_OneShotUL_TimingAdj_r17 != nil || ie.Pucch_Repetition_F0_2_r17 != nil || ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil || ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil || ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil || ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_RNTI_r17 != nil || ie.DynamicMulticastDCI_Format4_2_r17 != nil || ie.MaxModulationOrderForMulticast_r17 != nil || ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil || ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil || ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_CS_RNTI_r17 != nil || ie.Re_LevelRateMatchingForMulticast_r17 != nil || ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil || ie.Prs_MeasurementWithoutMG_r17 != nil || ie.MaxNumber_LEO_SatellitesPerCarrier_r17 != nil || len(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 || ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil || ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil || ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil || ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil || ie.Dmrs_BundlingPUSCH_multiSlot_r17 != nil || ie.Dmrs_BundlingPUCCH_Rep_r17 != nil || ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil || ie.InterSlotFreqHopPUCCH_r17 != nil || ie.Dmrs_BundlingRestart_r17 != nil || ie.Dmrs_BundlingNonBackToBackTX_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ModifiedMPR_Behaviour != nil, ie.Mimo_ParametersPerBand != nil, ie.ExtendedCP != nil, ie.MultipleTCI != nil, ie.Bwp_WithoutRestriction != nil, ie.Bwp_SameNumerology != nil, ie.Bwp_DiffNumerology != nil, ie.CrossCarrierScheduling_SameSCS != nil, ie.Pdsch_256QAM_FR2 != nil, ie.Pusch_256QAM != nil, ie.Ue_PowerClass != nil, ie.RateMatchingLTE_CRS != nil, ie.ChannelBWs_DL != nil, ie.ChannelBWs_UL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.BandNR.Encode(w); err != nil {
		return utils.WrapError("Encode BandNR", err)
	}
	if ie.ModifiedMPR_Behaviour != nil {
		if err = w.WriteBitString(ie.ModifiedMPR_Behaviour.Bytes, uint(ie.ModifiedMPR_Behaviour.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode ModifiedMPR_Behaviour", err)
		}
	}
	if ie.Mimo_ParametersPerBand != nil {
		if err = ie.Mimo_ParametersPerBand.Encode(w); err != nil {
			return utils.WrapError("Encode Mimo_ParametersPerBand", err)
		}
	}
	if ie.ExtendedCP != nil {
		if err = ie.ExtendedCP.Encode(w); err != nil {
			return utils.WrapError("Encode ExtendedCP", err)
		}
	}
	if ie.MultipleTCI != nil {
		if err = ie.MultipleTCI.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleTCI", err)
		}
	}
	if ie.Bwp_WithoutRestriction != nil {
		if err = ie.Bwp_WithoutRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_WithoutRestriction", err)
		}
	}
	if ie.Bwp_SameNumerology != nil {
		if err = ie.Bwp_SameNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_SameNumerology", err)
		}
	}
	if ie.Bwp_DiffNumerology != nil {
		if err = ie.Bwp_DiffNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode Bwp_DiffNumerology", err)
		}
	}
	if ie.CrossCarrierScheduling_SameSCS != nil {
		if err = ie.CrossCarrierScheduling_SameSCS.Encode(w); err != nil {
			return utils.WrapError("Encode CrossCarrierScheduling_SameSCS", err)
		}
	}
	if ie.Pdsch_256QAM_FR2 != nil {
		if err = ie.Pdsch_256QAM_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_256QAM_FR2", err)
		}
	}
	if ie.Pusch_256QAM != nil {
		if err = ie.Pusch_256QAM.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_256QAM", err)
		}
	}
	if ie.Ue_PowerClass != nil {
		if err = ie.Ue_PowerClass.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_PowerClass", err)
		}
	}
	if ie.RateMatchingLTE_CRS != nil {
		if err = ie.RateMatchingLTE_CRS.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchingLTE_CRS", err)
		}
	}
	if ie.ChannelBWs_DL != nil {
		if err = ie.ChannelBWs_DL.Encode(w); err != nil {
			return utils.WrapError("Encode ChannelBWs_DL", err)
		}
	}
	if ie.ChannelBWs_UL != nil {
		if err = ie.ChannelBWs_UL.Encode(w); err != nil {
			return utils.WrapError("Encode ChannelBWs_UL", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 14 bits for 14 extension groups
		extBitmap := []bool{ie.MaxUplinkDutyCycle_PC2_FR1 != nil, ie.Pucch_SpatialRelInfoMAC_CE != nil || ie.PowerBoosting_pi2BPSK != nil, ie.MaxUplinkDutyCycle_FR2 != nil, ie.ChannelBWs_DL_v1590 != nil || ie.ChannelBWs_UL_v1590 != nil, ie.AsymmetricBandwidthCombinationSet != nil, ie.SharedSpectrumChAccessParamsPerBand_r16 != nil || ie.CancelOverlappingPUSCH_r16 != nil || ie.MultipleRateMatchingEUTRA_CRS_r16 != nil || ie.OverlapRateMatchingEUTRA_CRS_r16 != nil || ie.Pdsch_MappingTypeB_Alt_r16 != nil || ie.OneSlotPeriodicTRS_r16 != nil || ie.Olpc_SRS_Pos_r16 != nil || ie.SpatialRelationsSRS_Pos_r16 != nil || ie.SimulSRS_MIMO_TransWithinBand_r16 != nil || ie.ChannelBW_DL_IAB_r16 != nil || ie.ChannelBW_UL_IAB_r16 != nil || ie.RasterShift7dot5_IAB_r16 != nil || ie.Ue_PowerClass_v1610 != nil || ie.CondHandover_r16 != nil || ie.CondHandoverFailure_r16 != nil || ie.CondHandoverTwoTriggerEvents_r16 != nil || ie.CondPSCellChange_r16 != nil || ie.CondPSCellChangeTwoTriggerEvents_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil || ie.ActiveConfiguredGrant_r16 != nil || ie.JointReleaseConfiguredGrantType2_r16 != nil || ie.Sps_r16 != nil || ie.JointReleaseSPS_r16 != nil || ie.SimulSRS_TransWithinBand_r16 != nil || ie.Trs_AdditionalBandwidth_r16 != nil || ie.HandoverIntraF_IAB_r16 != nil, ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil, ie.HandoverUTRA_FDD_r16 != nil || ie.EnhancedUL_TransientPeriod_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil, ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Pusch_RepetitionMultiSlots_v1650 != nil || ie.ConfiguredUL_GrantType1_v1650 != nil || ie.ConfiguredUL_GrantType2_v1650 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil, ie.EnhancedSkipUplinkTxConfigured_v1660 != nil || ie.EnhancedSkipUplinkTxDynamic_v1660 != nil, ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil || ie.TxDiversity_r16 != nil, ie.Pdsch_1024QAM_FR1_r17 != nil || ie.Ue_PowerClass_v1700 != nil || ie.Fr2_2_AccessParamsPerBand_r17 != nil || ie.Rlm_Relaxation_r17 != nil || ie.Bfd_Relaxation_r17 != nil || ie.Cg_SDT_r17 != nil || ie.LocationBasedCondHandover_r17 != nil || ie.TimeBasedCondHandover_r17 != nil || ie.EventA4BasedCondHandover_r17 != nil || ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Pdcch_SkippingWithoutSSSG_r17 != nil || ie.Sssg_Switching_1BitInd_r17 != nil || ie.Sssg_Switching_2BitInd_r17 != nil || ie.Pdcch_SkippingWithSSSG_r17 != nil || ie.SearchSpaceSetGrp_switchCap2_r17 != nil || ie.UplinkPreCompensation_r17 != nil || ie.Uplink_TA_Reporting_r17 != nil || ie.Max_HARQ_ProcessNumber_r17 != nil || ie.Type2_HARQ_Codebook_r17 != nil || ie.Type1_HARQ_Codebook_r17 != nil || ie.Type3_HARQ_Codebook_r17 != nil || ie.Ue_specific_K_Offset_r17 != nil || ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.ParallelPRS_MeasRRC_Inactive_r17 != nil || ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil || ie.Prs_ProcessingRRC_Inactive_r17 != nil || ie.Prs_ProcessingWindowType1A_r17 != nil || ie.Prs_ProcessingWindowType1B_r17 != nil || ie.Prs_ProcessingWindowType2_r17 != nil || ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil || ie.Olpc_SRS_PosRRC_Inactive_r17 != nil || ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil || ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil || ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil || ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil || ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil || ie.MaxDurationDMRS_Bundling_r17 != nil || ie.Pusch_RepetitionMsg3_r17 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil || ie.ParallelMeasurementWithoutRestriction_r17 != nil || ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil || ie.K1_RangeExtension_r17 != nil || ie.AperiodicCSI_RS_FastScellActivation_r17 != nil || ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil || ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil || ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil || ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil || ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil || ie.Ul_GapFR2_r17 != nil || ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil || ie.OneShotHARQ_feedbackPhy_Priority_r17 != nil || ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil || ie.TriggeredHARQ_CodebookRetx_r17 != nil, ie.Ue_OneShotUL_TimingAdj_r17 != nil || ie.Pucch_Repetition_F0_2_r17 != nil || ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil || ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil || ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil || ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_RNTI_r17 != nil || ie.DynamicMulticastDCI_Format4_2_r17 != nil || ie.MaxModulationOrderForMulticast_r17 != nil || ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil || ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil || ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_CS_RNTI_r17 != nil || ie.Re_LevelRateMatchingForMulticast_r17 != nil || ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil || ie.Prs_MeasurementWithoutMG_r17 != nil || ie.MaxNumber_LEO_SatellitesPerCarrier_r17 != nil || len(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 || ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil || ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil, ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil || ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil || ie.Dmrs_BundlingPUSCH_multiSlot_r17 != nil || ie.Dmrs_BundlingPUCCH_Rep_r17 != nil || ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil || ie.InterSlotFreqHopPUCCH_r17 != nil || ie.Dmrs_BundlingRestart_r17 != nil || ie.Dmrs_BundlingNonBackToBackTX_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.MaxUplinkDutyCycle_PC2_FR1 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxUplinkDutyCycle_PC2_FR1 optional
			if ie.MaxUplinkDutyCycle_PC2_FR1 != nil {
				if err = ie.MaxUplinkDutyCycle_PC2_FR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxUplinkDutyCycle_PC2_FR1", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Pucch_SpatialRelInfoMAC_CE != nil, ie.PowerBoosting_pi2BPSK != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pucch_SpatialRelInfoMAC_CE optional
			if ie.Pucch_SpatialRelInfoMAC_CE != nil {
				if err = ie.Pucch_SpatialRelInfoMAC_CE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_SpatialRelInfoMAC_CE", err)
				}
			}
			// encode PowerBoosting_pi2BPSK optional
			if ie.PowerBoosting_pi2BPSK != nil {
				if err = ie.PowerBoosting_pi2BPSK.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PowerBoosting_pi2BPSK", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.MaxUplinkDutyCycle_FR2 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxUplinkDutyCycle_FR2 optional
			if ie.MaxUplinkDutyCycle_FR2 != nil {
				if err = ie.MaxUplinkDutyCycle_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxUplinkDutyCycle_FR2", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.ChannelBWs_DL_v1590 != nil, ie.ChannelBWs_UL_v1590 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChannelBWs_DL_v1590 optional
			if ie.ChannelBWs_DL_v1590 != nil {
				if err = ie.ChannelBWs_DL_v1590.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelBWs_DL_v1590", err)
				}
			}
			// encode ChannelBWs_UL_v1590 optional
			if ie.ChannelBWs_UL_v1590 != nil {
				if err = ie.ChannelBWs_UL_v1590.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelBWs_UL_v1590", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.AsymmetricBandwidthCombinationSet != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AsymmetricBandwidthCombinationSet optional
			if ie.AsymmetricBandwidthCombinationSet != nil {
				if err = extWriter.WriteBitString(ie.AsymmetricBandwidthCombinationSet.Bytes, uint(ie.AsymmetricBandwidthCombinationSet.NumBits), &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode AsymmetricBandwidthCombinationSet", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.SharedSpectrumChAccessParamsPerBand_r16 != nil, ie.CancelOverlappingPUSCH_r16 != nil, ie.MultipleRateMatchingEUTRA_CRS_r16 != nil, ie.OverlapRateMatchingEUTRA_CRS_r16 != nil, ie.Pdsch_MappingTypeB_Alt_r16 != nil, ie.OneSlotPeriodicTRS_r16 != nil, ie.Olpc_SRS_Pos_r16 != nil, ie.SpatialRelationsSRS_Pos_r16 != nil, ie.SimulSRS_MIMO_TransWithinBand_r16 != nil, ie.ChannelBW_DL_IAB_r16 != nil, ie.ChannelBW_UL_IAB_r16 != nil, ie.RasterShift7dot5_IAB_r16 != nil, ie.Ue_PowerClass_v1610 != nil, ie.CondHandover_r16 != nil, ie.CondHandoverFailure_r16 != nil, ie.CondHandoverTwoTriggerEvents_r16 != nil, ie.CondPSCellChange_r16 != nil, ie.CondPSCellChangeTwoTriggerEvents_r16 != nil, ie.Mpr_PowerBoost_FR2_r16 != nil, ie.ActiveConfiguredGrant_r16 != nil, ie.JointReleaseConfiguredGrantType2_r16 != nil, ie.Sps_r16 != nil, ie.JointReleaseSPS_r16 != nil, ie.SimulSRS_TransWithinBand_r16 != nil, ie.Trs_AdditionalBandwidth_r16 != nil, ie.HandoverIntraF_IAB_r16 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SharedSpectrumChAccessParamsPerBand_r16 optional
			if ie.SharedSpectrumChAccessParamsPerBand_r16 != nil {
				if err = ie.SharedSpectrumChAccessParamsPerBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_r16", err)
				}
			}
			// encode CancelOverlappingPUSCH_r16 optional
			if ie.CancelOverlappingPUSCH_r16 != nil {
				if err = ie.CancelOverlappingPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CancelOverlappingPUSCH_r16", err)
				}
			}
			// encode MultipleRateMatchingEUTRA_CRS_r16 optional
			if ie.MultipleRateMatchingEUTRA_CRS_r16 != nil {
				if err = ie.MultipleRateMatchingEUTRA_CRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultipleRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// encode OverlapRateMatchingEUTRA_CRS_r16 optional
			if ie.OverlapRateMatchingEUTRA_CRS_r16 != nil {
				if err = ie.OverlapRateMatchingEUTRA_CRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OverlapRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// encode Pdsch_MappingTypeB_Alt_r16 optional
			if ie.Pdsch_MappingTypeB_Alt_r16 != nil {
				if err = ie.Pdsch_MappingTypeB_Alt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_MappingTypeB_Alt_r16", err)
				}
			}
			// encode OneSlotPeriodicTRS_r16 optional
			if ie.OneSlotPeriodicTRS_r16 != nil {
				if err = ie.OneSlotPeriodicTRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OneSlotPeriodicTRS_r16", err)
				}
			}
			// encode Olpc_SRS_Pos_r16 optional
			if ie.Olpc_SRS_Pos_r16 != nil {
				if err = ie.Olpc_SRS_Pos_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Olpc_SRS_Pos_r16", err)
				}
			}
			// encode SpatialRelationsSRS_Pos_r16 optional
			if ie.SpatialRelationsSRS_Pos_r16 != nil {
				if err = ie.SpatialRelationsSRS_Pos_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationsSRS_Pos_r16", err)
				}
			}
			// encode SimulSRS_MIMO_TransWithinBand_r16 optional
			if ie.SimulSRS_MIMO_TransWithinBand_r16 != nil {
				if err = ie.SimulSRS_MIMO_TransWithinBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimulSRS_MIMO_TransWithinBand_r16", err)
				}
			}
			// encode ChannelBW_DL_IAB_r16 optional
			if ie.ChannelBW_DL_IAB_r16 != nil {
				if err = ie.ChannelBW_DL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelBW_DL_IAB_r16", err)
				}
			}
			// encode ChannelBW_UL_IAB_r16 optional
			if ie.ChannelBW_UL_IAB_r16 != nil {
				if err = ie.ChannelBW_UL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChannelBW_UL_IAB_r16", err)
				}
			}
			// encode RasterShift7dot5_IAB_r16 optional
			if ie.RasterShift7dot5_IAB_r16 != nil {
				if err = ie.RasterShift7dot5_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RasterShift7dot5_IAB_r16", err)
				}
			}
			// encode Ue_PowerClass_v1610 optional
			if ie.Ue_PowerClass_v1610 != nil {
				if err = ie.Ue_PowerClass_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ue_PowerClass_v1610", err)
				}
			}
			// encode CondHandover_r16 optional
			if ie.CondHandover_r16 != nil {
				if err = ie.CondHandover_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondHandover_r16", err)
				}
			}
			// encode CondHandoverFailure_r16 optional
			if ie.CondHandoverFailure_r16 != nil {
				if err = ie.CondHandoverFailure_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondHandoverFailure_r16", err)
				}
			}
			// encode CondHandoverTwoTriggerEvents_r16 optional
			if ie.CondHandoverTwoTriggerEvents_r16 != nil {
				if err = ie.CondHandoverTwoTriggerEvents_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondHandoverTwoTriggerEvents_r16", err)
				}
			}
			// encode CondPSCellChange_r16 optional
			if ie.CondPSCellChange_r16 != nil {
				if err = ie.CondPSCellChange_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondPSCellChange_r16", err)
				}
			}
			// encode CondPSCellChangeTwoTriggerEvents_r16 optional
			if ie.CondPSCellChangeTwoTriggerEvents_r16 != nil {
				if err = ie.CondPSCellChangeTwoTriggerEvents_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondPSCellChangeTwoTriggerEvents_r16", err)
				}
			}
			// encode Mpr_PowerBoost_FR2_r16 optional
			if ie.Mpr_PowerBoost_FR2_r16 != nil {
				if err = ie.Mpr_PowerBoost_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpr_PowerBoost_FR2_r16", err)
				}
			}
			// encode ActiveConfiguredGrant_r16 optional
			if ie.ActiveConfiguredGrant_r16 != nil {
				if err = ie.ActiveConfiguredGrant_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ActiveConfiguredGrant_r16", err)
				}
			}
			// encode JointReleaseConfiguredGrantType2_r16 optional
			if ie.JointReleaseConfiguredGrantType2_r16 != nil {
				if err = ie.JointReleaseConfiguredGrantType2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode JointReleaseConfiguredGrantType2_r16", err)
				}
			}
			// encode Sps_r16 optional
			if ie.Sps_r16 != nil {
				if err = ie.Sps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sps_r16", err)
				}
			}
			// encode JointReleaseSPS_r16 optional
			if ie.JointReleaseSPS_r16 != nil {
				if err = ie.JointReleaseSPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode JointReleaseSPS_r16", err)
				}
			}
			// encode SimulSRS_TransWithinBand_r16 optional
			if ie.SimulSRS_TransWithinBand_r16 != nil {
				if err = ie.SimulSRS_TransWithinBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimulSRS_TransWithinBand_r16", err)
				}
			}
			// encode Trs_AdditionalBandwidth_r16 optional
			if ie.Trs_AdditionalBandwidth_r16 != nil {
				if err = ie.Trs_AdditionalBandwidth_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Trs_AdditionalBandwidth_r16", err)
				}
			}
			// encode HandoverIntraF_IAB_r16 optional
			if ie.HandoverIntraF_IAB_r16 != nil {
				if err = ie.HandoverIntraF_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverIntraF_IAB_r16", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 optional
			if ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil {
				if err = ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16", err)
				}
			}
			// encode SharedSpectrumChAccessParamsPerBand_v1630 optional
			if ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil {
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_v1630", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.HandoverUTRA_FDD_r16 != nil, ie.EnhancedUL_TransientPeriod_r16 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode HandoverUTRA_FDD_r16 optional
			if ie.HandoverUTRA_FDD_r16 != nil {
				if err = ie.HandoverUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverUTRA_FDD_r16", err)
				}
			}
			// encode EnhancedUL_TransientPeriod_r16 optional
			if ie.EnhancedUL_TransientPeriod_r16 != nil {
				if err = ie.EnhancedUL_TransientPeriod_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedUL_TransientPeriod_r16", err)
				}
			}
			// encode SharedSpectrumChAccessParamsPerBand_v1640 optional
			if ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil {
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_v1640", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.Pusch_RepetitionMultiSlots_v1650 != nil, ie.ConfiguredUL_GrantType1_v1650 != nil, ie.ConfiguredUL_GrantType2_v1650 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Type1_PUSCH_RepetitionMultiSlots_v1650 optional
			if ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err = ie.Type1_PUSCH_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type1_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode Type2_PUSCH_RepetitionMultiSlots_v1650 optional
			if ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err = ie.Type2_PUSCH_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type2_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode Pusch_RepetitionMultiSlots_v1650 optional
			if ie.Pusch_RepetitionMultiSlots_v1650 != nil {
				if err = ie.Pusch_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode ConfiguredUL_GrantType1_v1650 optional
			if ie.ConfiguredUL_GrantType1_v1650 != nil {
				if err = ie.ConfiguredUL_GrantType1_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredUL_GrantType1_v1650", err)
				}
			}
			// encode ConfiguredUL_GrantType2_v1650 optional
			if ie.ConfiguredUL_GrantType2_v1650 != nil {
				if err = ie.ConfiguredUL_GrantType2_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConfiguredUL_GrantType2_v1650", err)
				}
			}
			// encode SharedSpectrumChAccessParamsPerBand_v1650 optional
			if ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil {
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_v1650", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.EnhancedSkipUplinkTxConfigured_v1660 != nil, ie.EnhancedSkipUplinkTxDynamic_v1660 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EnhancedSkipUplinkTxConfigured_v1660 optional
			if ie.EnhancedSkipUplinkTxConfigured_v1660 != nil {
				if err = ie.EnhancedSkipUplinkTxConfigured_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxConfigured_v1660", err)
				}
			}
			// encode EnhancedSkipUplinkTxDynamic_v1660 optional
			if ie.EnhancedSkipUplinkTxDynamic_v1660 != nil {
				if err = ie.EnhancedSkipUplinkTxDynamic_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedSkipUplinkTxDynamic_v1660", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil, ie.TxDiversity_r16 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 optional
			if ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil {
				if err = ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
				}
			}
			// encode TxDiversity_r16 optional
			if ie.TxDiversity_r16 != nil {
				if err = ie.TxDiversity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TxDiversity_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 12
		if extBitmap[11] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 12
			optionals_ext_12 := []bool{ie.Pdsch_1024QAM_FR1_r17 != nil, ie.Ue_PowerClass_v1700 != nil, ie.Fr2_2_AccessParamsPerBand_r17 != nil, ie.Rlm_Relaxation_r17 != nil, ie.Bfd_Relaxation_r17 != nil, ie.Cg_SDT_r17 != nil, ie.LocationBasedCondHandover_r17 != nil, ie.TimeBasedCondHandover_r17 != nil, ie.EventA4BasedCondHandover_r17 != nil, ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.Pdcch_SkippingWithoutSSSG_r17 != nil, ie.Sssg_Switching_1BitInd_r17 != nil, ie.Sssg_Switching_2BitInd_r17 != nil, ie.Pdcch_SkippingWithSSSG_r17 != nil, ie.SearchSpaceSetGrp_switchCap2_r17 != nil, ie.UplinkPreCompensation_r17 != nil, ie.Uplink_TA_Reporting_r17 != nil, ie.Max_HARQ_ProcessNumber_r17 != nil, ie.Type2_HARQ_Codebook_r17 != nil, ie.Type1_HARQ_Codebook_r17 != nil, ie.Type3_HARQ_Codebook_r17 != nil, ie.Ue_specific_K_Offset_r17 != nil, ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.ParallelPRS_MeasRRC_Inactive_r17 != nil, ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil, ie.Prs_ProcessingRRC_Inactive_r17 != nil, ie.Prs_ProcessingWindowType1A_r17 != nil, ie.Prs_ProcessingWindowType1B_r17 != nil, ie.Prs_ProcessingWindowType2_r17 != nil, ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil, ie.Olpc_SRS_PosRRC_Inactive_r17 != nil, ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil, ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil, ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil, ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil, ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil, ie.MaxDurationDMRS_Bundling_r17 != nil, ie.Pusch_RepetitionMsg3_r17 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil, ie.ParallelMeasurementWithoutRestriction_r17 != nil, ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil, ie.K1_RangeExtension_r17 != nil, ie.AperiodicCSI_RS_FastScellActivation_r17 != nil, ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil, ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil, ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil, ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil, ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil, ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil, ie.Ul_GapFR2_r17 != nil, ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil, ie.OneShotHARQ_feedbackPhy_Priority_r17 != nil, ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil, ie.TriggeredHARQ_CodebookRetx_r17 != nil}
			for _, bit := range optionals_ext_12 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdsch_1024QAM_FR1_r17 optional
			if ie.Pdsch_1024QAM_FR1_r17 != nil {
				if err = ie.Pdsch_1024QAM_FR1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_1024QAM_FR1_r17", err)
				}
			}
			// encode Ue_PowerClass_v1700 optional
			if ie.Ue_PowerClass_v1700 != nil {
				if err = ie.Ue_PowerClass_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ue_PowerClass_v1700", err)
				}
			}
			// encode Fr2_2_AccessParamsPerBand_r17 optional
			if ie.Fr2_2_AccessParamsPerBand_r17 != nil {
				if err = ie.Fr2_2_AccessParamsPerBand_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Fr2_2_AccessParamsPerBand_r17", err)
				}
			}
			// encode Rlm_Relaxation_r17 optional
			if ie.Rlm_Relaxation_r17 != nil {
				if err = ie.Rlm_Relaxation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rlm_Relaxation_r17", err)
				}
			}
			// encode Bfd_Relaxation_r17 optional
			if ie.Bfd_Relaxation_r17 != nil {
				if err = ie.Bfd_Relaxation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bfd_Relaxation_r17", err)
				}
			}
			// encode Cg_SDT_r17 optional
			if ie.Cg_SDT_r17 != nil {
				if err = ie.Cg_SDT_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cg_SDT_r17", err)
				}
			}
			// encode LocationBasedCondHandover_r17 optional
			if ie.LocationBasedCondHandover_r17 != nil {
				if err = ie.LocationBasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LocationBasedCondHandover_r17", err)
				}
			}
			// encode TimeBasedCondHandover_r17 optional
			if ie.TimeBasedCondHandover_r17 != nil {
				if err = ie.TimeBasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TimeBasedCondHandover_r17", err)
				}
			}
			// encode EventA4BasedCondHandover_r17 optional
			if ie.EventA4BasedCondHandover_r17 != nil {
				if err = ie.EventA4BasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EventA4BasedCondHandover_r17", err)
				}
			}
			// encode Mn_InitiatedCondPSCellChangeNRDC_r17 optional
			if ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err = ie.Mn_InitiatedCondPSCellChangeNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// encode Sn_InitiatedCondPSCellChangeNRDC_r17 optional
			if ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err = ie.Sn_InitiatedCondPSCellChangeNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// encode Pdcch_SkippingWithoutSSSG_r17 optional
			if ie.Pdcch_SkippingWithoutSSSG_r17 != nil {
				if err = ie.Pdcch_SkippingWithoutSSSG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_SkippingWithoutSSSG_r17", err)
				}
			}
			// encode Sssg_Switching_1BitInd_r17 optional
			if ie.Sssg_Switching_1BitInd_r17 != nil {
				if err = ie.Sssg_Switching_1BitInd_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sssg_Switching_1BitInd_r17", err)
				}
			}
			// encode Sssg_Switching_2BitInd_r17 optional
			if ie.Sssg_Switching_2BitInd_r17 != nil {
				if err = ie.Sssg_Switching_2BitInd_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sssg_Switching_2BitInd_r17", err)
				}
			}
			// encode Pdcch_SkippingWithSSSG_r17 optional
			if ie.Pdcch_SkippingWithSSSG_r17 != nil {
				if err = ie.Pdcch_SkippingWithSSSG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_SkippingWithSSSG_r17", err)
				}
			}
			// encode SearchSpaceSetGrp_switchCap2_r17 optional
			if ie.SearchSpaceSetGrp_switchCap2_r17 != nil {
				if err = ie.SearchSpaceSetGrp_switchCap2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SearchSpaceSetGrp_switchCap2_r17", err)
				}
			}
			// encode UplinkPreCompensation_r17 optional
			if ie.UplinkPreCompensation_r17 != nil {
				if err = ie.UplinkPreCompensation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkPreCompensation_r17", err)
				}
			}
			// encode Uplink_TA_Reporting_r17 optional
			if ie.Uplink_TA_Reporting_r17 != nil {
				if err = ie.Uplink_TA_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uplink_TA_Reporting_r17", err)
				}
			}
			// encode Max_HARQ_ProcessNumber_r17 optional
			if ie.Max_HARQ_ProcessNumber_r17 != nil {
				if err = ie.Max_HARQ_ProcessNumber_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Max_HARQ_ProcessNumber_r17", err)
				}
			}
			// encode Type2_HARQ_Codebook_r17 optional
			if ie.Type2_HARQ_Codebook_r17 != nil {
				if err = ie.Type2_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type2_HARQ_Codebook_r17", err)
				}
			}
			// encode Type1_HARQ_Codebook_r17 optional
			if ie.Type1_HARQ_Codebook_r17 != nil {
				if err = ie.Type1_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type1_HARQ_Codebook_r17", err)
				}
			}
			// encode Type3_HARQ_Codebook_r17 optional
			if ie.Type3_HARQ_Codebook_r17 != nil {
				if err = ie.Type3_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type3_HARQ_Codebook_r17", err)
				}
			}
			// encode Ue_specific_K_Offset_r17 optional
			if ie.Ue_specific_K_Offset_r17 != nil {
				if err = ie.Ue_specific_K_Offset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ue_specific_K_Offset_r17", err)
				}
			}
			// encode MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err = ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// encode MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err = ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// encode ParallelPRS_MeasRRC_Inactive_r17 optional
			if ie.ParallelPRS_MeasRRC_Inactive_r17 != nil {
				if err = ie.ParallelPRS_MeasRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ParallelPRS_MeasRRC_Inactive_r17", err)
				}
			}
			// encode Nr_UE_TxTEG_ID_MaxSupport_r17 optional
			if ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil {
				if err = ie.Nr_UE_TxTEG_ID_MaxSupport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_UE_TxTEG_ID_MaxSupport_r17", err)
				}
			}
			// encode Prs_ProcessingRRC_Inactive_r17 optional
			if ie.Prs_ProcessingRRC_Inactive_r17 != nil {
				if err = ie.Prs_ProcessingRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_ProcessingRRC_Inactive_r17", err)
				}
			}
			// encode Prs_ProcessingWindowType1A_r17 optional
			if ie.Prs_ProcessingWindowType1A_r17 != nil {
				if err = ie.Prs_ProcessingWindowType1A_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_ProcessingWindowType1A_r17", err)
				}
			}
			// encode Prs_ProcessingWindowType1B_r17 optional
			if ie.Prs_ProcessingWindowType1B_r17 != nil {
				if err = ie.Prs_ProcessingWindowType1B_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_ProcessingWindowType1B_r17", err)
				}
			}
			// encode Prs_ProcessingWindowType2_r17 optional
			if ie.Prs_ProcessingWindowType2_r17 != nil {
				if err = ie.Prs_ProcessingWindowType2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_ProcessingWindowType2_r17", err)
				}
			}
			// encode Srs_AllPosResourcesRRC_Inactive_r17 optional
			if ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil {
				if err = ie.Srs_AllPosResourcesRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_AllPosResourcesRRC_Inactive_r17", err)
				}
			}
			// encode Olpc_SRS_PosRRC_Inactive_r17 optional
			if ie.Olpc_SRS_PosRRC_Inactive_r17 != nil {
				if err = ie.Olpc_SRS_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Olpc_SRS_PosRRC_Inactive_r17", err)
				}
			}
			// encode SpatialRelationsSRS_PosRRC_Inactive_r17 optional
			if ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil {
				if err = ie.SpatialRelationsSRS_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationsSRS_PosRRC_Inactive_r17", err)
				}
			}
			// encode MaxNumberPUSCH_TypeA_Repetition_r17 optional
			if ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil {
				if err = ie.MaxNumberPUSCH_TypeA_Repetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberPUSCH_TypeA_Repetition_r17", err)
				}
			}
			// encode PuschTypeA_RepetitionsAvailSlot_r17 optional
			if ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil {
				if err = ie.PuschTypeA_RepetitionsAvailSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PuschTypeA_RepetitionsAvailSlot_r17", err)
				}
			}
			// encode Tb_ProcessingMultiSlotPUSCH_r17 optional
			if ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil {
				if err = ie.Tb_ProcessingMultiSlotPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tb_ProcessingMultiSlotPUSCH_r17", err)
				}
			}
			// encode Tb_ProcessingRepMultiSlotPUSCH_r17 optional
			if ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil {
				if err = ie.Tb_ProcessingRepMultiSlotPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Tb_ProcessingRepMultiSlotPUSCH_r17", err)
				}
			}
			// encode MaxDurationDMRS_Bundling_r17 optional
			if ie.MaxDurationDMRS_Bundling_r17 != nil {
				if err = ie.MaxDurationDMRS_Bundling_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxDurationDMRS_Bundling_r17", err)
				}
			}
			// encode Pusch_RepetitionMsg3_r17 optional
			if ie.Pusch_RepetitionMsg3_r17 != nil {
				if err = ie.Pusch_RepetitionMsg3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pusch_RepetitionMsg3_r17", err)
				}
			}
			// encode SharedSpectrumChAccessParamsPerBand_v1710 optional
			if ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil {
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SharedSpectrumChAccessParamsPerBand_v1710", err)
				}
			}
			// encode ParallelMeasurementWithoutRestriction_r17 optional
			if ie.ParallelMeasurementWithoutRestriction_r17 != nil {
				if err = ie.ParallelMeasurementWithoutRestriction_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ParallelMeasurementWithoutRestriction_r17", err)
				}
			}
			// encode MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 optional
			if ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil {
				if err = ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
				}
			}
			// encode K1_RangeExtension_r17 optional
			if ie.K1_RangeExtension_r17 != nil {
				if err = ie.K1_RangeExtension_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode K1_RangeExtension_r17", err)
				}
			}
			// encode AperiodicCSI_RS_FastScellActivation_r17 optional
			if ie.AperiodicCSI_RS_FastScellActivation_r17 != nil {
				if err = ie.AperiodicCSI_RS_FastScellActivation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AperiodicCSI_RS_FastScellActivation_r17", err)
				}
			}
			// encode AperiodicCSI_RS_AdditionalBandwidth_r17 optional
			if ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil {
				if err = ie.AperiodicCSI_RS_AdditionalBandwidth_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AperiodicCSI_RS_AdditionalBandwidth_r17", err)
				}
			}
			// encode Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 optional
			if ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil {
				if err = ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17", err)
				}
			}
			// encode HalfDuplexFDD_TypeA_RedCap_r17 optional
			if ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil {
				if err = ie.HalfDuplexFDD_TypeA_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HalfDuplexFDD_TypeA_RedCap_r17", err)
				}
			}
			// encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 optional
			if ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil {
				if err = ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17", err)
				}
			}
			// encode ChannelBWs_DL_SCS_480kHz_FR2_2_r17 optional
			if ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_DL_SCS_480kHz_FR2_2_r17", err)
				}
			}
			// encode ChannelBWs_UL_SCS_480kHz_FR2_2_r17 optional
			if ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_UL_SCS_480kHz_FR2_2_r17", err)
				}
			}
			// encode ChannelBWs_DL_SCS_960kHz_FR2_2_r17 optional
			if ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_DL_SCS_960kHz_FR2_2_r17", err)
				}
			}
			// encode ChannelBWs_UL_SCS_960kHz_FR2_2_r17 optional
			if ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_UL_SCS_960kHz_FR2_2_r17", err)
				}
			}
			// encode Ul_GapFR2_r17 optional
			if ie.Ul_GapFR2_r17 != nil {
				if err = ie.Ul_GapFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_GapFR2_r17", err)
				}
			}
			// encode OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 optional
			if ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil {
				if err = ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OneShotHARQ_feedbackTriggeredByDCI_1_2_r17", err)
				}
			}
			// encode OneShotHARQ_feedbackPhy_Priority_r17 optional
			if ie.OneShotHARQ_feedbackPhy_Priority_r17 != nil {
				if err = ie.OneShotHARQ_feedbackPhy_Priority_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OneShotHARQ_feedbackPhy_Priority_r17", err)
				}
			}
			// encode EnhancedType3_HARQ_CodebookFeedback_r17 optional
			if ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil {
				if err = ie.EnhancedType3_HARQ_CodebookFeedback_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedType3_HARQ_CodebookFeedback_r17", err)
				}
			}
			// encode TriggeredHARQ_CodebookRetx_r17 optional
			if ie.TriggeredHARQ_CodebookRetx_r17 != nil {
				if err = ie.TriggeredHARQ_CodebookRetx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TriggeredHARQ_CodebookRetx_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 13
		if extBitmap[12] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 13
			optionals_ext_13 := []bool{ie.Ue_OneShotUL_TimingAdj_r17 != nil, ie.Pucch_Repetition_F0_2_r17 != nil, ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil, ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil, ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil, ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.MaxNumberG_RNTI_r17 != nil, ie.DynamicMulticastDCI_Format4_2_r17 != nil, ie.MaxModulationOrderForMulticast_r17 != nil, ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil, ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil, ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil, ie.MaxNumberG_CS_RNTI_r17 != nil, ie.Re_LevelRateMatchingForMulticast_r17 != nil, ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil, ie.Prs_MeasurementWithoutMG_r17 != nil, ie.MaxNumber_LEO_SatellitesPerCarrier_r17 != nil, len(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0, ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil, ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil}
			for _, bit := range optionals_ext_13 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Ue_OneShotUL_TimingAdj_r17 optional
			if ie.Ue_OneShotUL_TimingAdj_r17 != nil {
				if err = ie.Ue_OneShotUL_TimingAdj_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ue_OneShotUL_TimingAdj_r17", err)
				}
			}
			// encode Pucch_Repetition_F0_2_r17 optional
			if ie.Pucch_Repetition_F0_2_r17 != nil {
				if err = ie.Pucch_Repetition_F0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pucch_Repetition_F0_2_r17", err)
				}
			}
			// encode Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 optional
			if ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode Mux_HARQ_ACK_DiffPriorities_r17 optional
			if ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil {
				if err = ie.Mux_HARQ_ACK_DiffPriorities_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_HARQ_ACK_DiffPriorities_r17", err)
				}
			}
			// encode Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 optional
			if ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 optional
			if ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err = ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode MaxNumberG_RNTI_r17 optional
			if ie.MaxNumberG_RNTI_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberG_RNTI_r17, &aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode MaxNumberG_RNTI_r17", err)
				}
			}
			// encode DynamicMulticastDCI_Format4_2_r17 optional
			if ie.DynamicMulticastDCI_Format4_2_r17 != nil {
				if err = ie.DynamicMulticastDCI_Format4_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DynamicMulticastDCI_Format4_2_r17", err)
				}
			}
			// encode MaxModulationOrderForMulticast_r17 optional
			if ie.MaxModulationOrderForMulticast_r17 != nil {
				if err = ie.MaxModulationOrderForMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxModulationOrderForMulticast_r17", err)
				}
			}
			// encode DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 optional
			if ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// encode DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 optional
			if ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 optional
			if ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err = ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 optional
			if ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil {
				if err = ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode MaxNumberG_CS_RNTI_r17 optional
			if ie.MaxNumberG_CS_RNTI_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberG_CS_RNTI_r17, &aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode MaxNumberG_CS_RNTI_r17", err)
				}
			}
			// encode Re_LevelRateMatchingForMulticast_r17 optional
			if ie.Re_LevelRateMatchingForMulticast_r17 != nil {
				if err = ie.Re_LevelRateMatchingForMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Re_LevelRateMatchingForMulticast_r17", err)
				}
			}
			// encode Pdsch_1024QAM_2MIMO_FR1_r17 optional
			if ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil {
				if err = ie.Pdsch_1024QAM_2MIMO_FR1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_1024QAM_2MIMO_FR1_r17", err)
				}
			}
			// encode Prs_MeasurementWithoutMG_r17 optional
			if ie.Prs_MeasurementWithoutMG_r17 != nil {
				if err = ie.Prs_MeasurementWithoutMG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_MeasurementWithoutMG_r17", err)
				}
			}
			// encode MaxNumber_LEO_SatellitesPerCarrier_r17 optional
			if ie.MaxNumber_LEO_SatellitesPerCarrier_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumber_LEO_SatellitesPerCarrier_r17, &aper.Constraint{Lb: 3, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode MaxNumber_LEO_SatellitesPerCarrier_r17", err)
				}
			}
			// encode Prs_ProcessingCapabilityOutsideMGinPPW_r17 optional
			if len(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 {
				tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17 := utils.NewSequence[*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17]([]*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}, aper.Constraint{Lb: 1, Ub: 3}, false)
				for _, i := range ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 {
					tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17.Value = append(tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17.Value, &i)
				}
				if err = tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Prs_ProcessingCapabilityOutsideMGinPPW_r17", err)
				}
			}
			// encode Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 optional
			if ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil {
				if err = ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_SemiPersistent_PosResourcesRRC_Inactive_r17", err)
				}
			}
			// encode ChannelBWs_DL_SCS_120kHz_FR2_2_r17 optional
			if ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_DL_SCS_120kHz_FR2_2_r17", err)
				}
			}
			// encode ChannelBWs_UL_SCS_120kHz_FR2_2_r17 optional
			if ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17.Bytes, uint(ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode ChannelBWs_UL_SCS_120kHz_FR2_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 14
		if extBitmap[13] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 14
			optionals_ext_14 := []bool{ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil, ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil, ie.Dmrs_BundlingPUSCH_multiSlot_r17 != nil, ie.Dmrs_BundlingPUCCH_Rep_r17 != nil, ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil, ie.InterSlotFreqHopPUCCH_r17 != nil, ie.Dmrs_BundlingRestart_r17 != nil, ie.Dmrs_BundlingNonBackToBackTX_r17 != nil}
			for _, bit := range optionals_ext_14 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dmrs_BundlingPUSCH_RepTypeA_r17 optional
			if ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil {
				if err = ie.Dmrs_BundlingPUSCH_RepTypeA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUSCH_RepTypeA_r17", err)
				}
			}
			// encode Dmrs_BundlingPUSCH_RepTypeB_r17 optional
			if ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil {
				if err = ie.Dmrs_BundlingPUSCH_RepTypeB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUSCH_RepTypeB_r17", err)
				}
			}
			// encode Dmrs_BundlingPUSCH_multiSlot_r17 optional
			if ie.Dmrs_BundlingPUSCH_multiSlot_r17 != nil {
				if err = ie.Dmrs_BundlingPUSCH_multiSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUSCH_multiSlot_r17", err)
				}
			}
			// encode Dmrs_BundlingPUCCH_Rep_r17 optional
			if ie.Dmrs_BundlingPUCCH_Rep_r17 != nil {
				if err = ie.Dmrs_BundlingPUCCH_Rep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingPUCCH_Rep_r17", err)
				}
			}
			// encode InterSlotFreqHopInterSlotBundlingPUSCH_r17 optional
			if ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil {
				if err = ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterSlotFreqHopInterSlotBundlingPUSCH_r17", err)
				}
			}
			// encode InterSlotFreqHopPUCCH_r17 optional
			if ie.InterSlotFreqHopPUCCH_r17 != nil {
				if err = ie.InterSlotFreqHopPUCCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterSlotFreqHopPUCCH_r17", err)
				}
			}
			// encode Dmrs_BundlingRestart_r17 optional
			if ie.Dmrs_BundlingRestart_r17 != nil {
				if err = ie.Dmrs_BundlingRestart_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingRestart_r17", err)
				}
			}
			// encode Dmrs_BundlingNonBackToBackTX_r17 optional
			if ie.Dmrs_BundlingNonBackToBackTX_r17 != nil {
				if err = ie.Dmrs_BundlingNonBackToBackTX_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dmrs_BundlingNonBackToBackTX_r17", err)
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

func (ie *BandNR) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ModifiedMPR_BehaviourPresent bool
	if ModifiedMPR_BehaviourPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mimo_ParametersPerBandPresent bool
	if Mimo_ParametersPerBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ExtendedCPPresent bool
	if ExtendedCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultipleTCIPresent bool
	if MultipleTCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_WithoutRestrictionPresent bool
	if Bwp_WithoutRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_SameNumerologyPresent bool
	if Bwp_SameNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Bwp_DiffNumerologyPresent bool
	if Bwp_DiffNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CrossCarrierScheduling_SameSCSPresent bool
	if CrossCarrierScheduling_SameSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_256QAM_FR2Present bool
	if Pdsch_256QAM_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_256QAMPresent bool
	if Pusch_256QAMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_PowerClassPresent bool
	if Ue_PowerClassPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchingLTE_CRSPresent bool
	if RateMatchingLTE_CRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ChannelBWs_DLPresent bool
	if ChannelBWs_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ChannelBWs_ULPresent bool
	if ChannelBWs_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.BandNR.Decode(r); err != nil {
		return utils.WrapError("Decode BandNR", err)
	}
	if ModifiedMPR_BehaviourPresent {
		var tmp_bs_ModifiedMPR_Behaviour []byte
		var n_ModifiedMPR_Behaviour uint
		if tmp_bs_ModifiedMPR_Behaviour, n_ModifiedMPR_Behaviour, err = r.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode ModifiedMPR_Behaviour", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_ModifiedMPR_Behaviour,
			NumBits: uint64(n_ModifiedMPR_Behaviour),
		}
		ie.ModifiedMPR_Behaviour = &tmp_bitstring
	}
	if Mimo_ParametersPerBandPresent {
		ie.Mimo_ParametersPerBand = new(MIMO_ParametersPerBand)
		if err = ie.Mimo_ParametersPerBand.Decode(r); err != nil {
			return utils.WrapError("Decode Mimo_ParametersPerBand", err)
		}
	}
	if ExtendedCPPresent {
		ie.ExtendedCP = new(BandNR_extendedCP)
		if err = ie.ExtendedCP.Decode(r); err != nil {
			return utils.WrapError("Decode ExtendedCP", err)
		}
	}
	if MultipleTCIPresent {
		ie.MultipleTCI = new(BandNR_multipleTCI)
		if err = ie.MultipleTCI.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleTCI", err)
		}
	}
	if Bwp_WithoutRestrictionPresent {
		ie.Bwp_WithoutRestriction = new(BandNR_bwp_WithoutRestriction)
		if err = ie.Bwp_WithoutRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_WithoutRestriction", err)
		}
	}
	if Bwp_SameNumerologyPresent {
		ie.Bwp_SameNumerology = new(BandNR_bwp_SameNumerology)
		if err = ie.Bwp_SameNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_SameNumerology", err)
		}
	}
	if Bwp_DiffNumerologyPresent {
		ie.Bwp_DiffNumerology = new(BandNR_bwp_DiffNumerology)
		if err = ie.Bwp_DiffNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode Bwp_DiffNumerology", err)
		}
	}
	if CrossCarrierScheduling_SameSCSPresent {
		ie.CrossCarrierScheduling_SameSCS = new(BandNR_crossCarrierScheduling_SameSCS)
		if err = ie.CrossCarrierScheduling_SameSCS.Decode(r); err != nil {
			return utils.WrapError("Decode CrossCarrierScheduling_SameSCS", err)
		}
	}
	if Pdsch_256QAM_FR2Present {
		ie.Pdsch_256QAM_FR2 = new(BandNR_pdsch_256QAM_FR2)
		if err = ie.Pdsch_256QAM_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_256QAM_FR2", err)
		}
	}
	if Pusch_256QAMPresent {
		ie.Pusch_256QAM = new(BandNR_pusch_256QAM)
		if err = ie.Pusch_256QAM.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_256QAM", err)
		}
	}
	if Ue_PowerClassPresent {
		ie.Ue_PowerClass = new(BandNR_ue_PowerClass)
		if err = ie.Ue_PowerClass.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_PowerClass", err)
		}
	}
	if RateMatchingLTE_CRSPresent {
		ie.RateMatchingLTE_CRS = new(BandNR_rateMatchingLTE_CRS)
		if err = ie.RateMatchingLTE_CRS.Decode(r); err != nil {
			return utils.WrapError("Decode RateMatchingLTE_CRS", err)
		}
	}
	if ChannelBWs_DLPresent {
		ie.ChannelBWs_DL = new(BandNR_channelBWs_DL)
		if err = ie.ChannelBWs_DL.Decode(r); err != nil {
			return utils.WrapError("Decode ChannelBWs_DL", err)
		}
	}
	if ChannelBWs_ULPresent {
		ie.ChannelBWs_UL = new(BandNR_channelBWs_UL)
		if err = ie.ChannelBWs_UL.Decode(r); err != nil {
			return utils.WrapError("Decode ChannelBWs_UL", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 14 bits for 14 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxUplinkDutyCycle_PC2_FR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxUplinkDutyCycle_PC2_FR1 optional
			if MaxUplinkDutyCycle_PC2_FR1Present {
				ie.MaxUplinkDutyCycle_PC2_FR1 = new(BandNR_maxUplinkDutyCycle_PC2_FR1)
				if err = ie.MaxUplinkDutyCycle_PC2_FR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxUplinkDutyCycle_PC2_FR1", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Pucch_SpatialRelInfoMAC_CEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PowerBoosting_pi2BPSKPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pucch_SpatialRelInfoMAC_CE optional
			if Pucch_SpatialRelInfoMAC_CEPresent {
				ie.Pucch_SpatialRelInfoMAC_CE = new(BandNR_pucch_SpatialRelInfoMAC_CE)
				if err = ie.Pucch_SpatialRelInfoMAC_CE.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_SpatialRelInfoMAC_CE", err)
				}
			}
			// decode PowerBoosting_pi2BPSK optional
			if PowerBoosting_pi2BPSKPresent {
				ie.PowerBoosting_pi2BPSK = new(BandNR_powerBoosting_pi2BPSK)
				if err = ie.PowerBoosting_pi2BPSK.Decode(extReader); err != nil {
					return utils.WrapError("Decode PowerBoosting_pi2BPSK", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxUplinkDutyCycle_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxUplinkDutyCycle_FR2 optional
			if MaxUplinkDutyCycle_FR2Present {
				ie.MaxUplinkDutyCycle_FR2 = new(BandNR_maxUplinkDutyCycle_FR2)
				if err = ie.MaxUplinkDutyCycle_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxUplinkDutyCycle_FR2", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			ChannelBWs_DL_v1590Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_UL_v1590Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChannelBWs_DL_v1590 optional
			if ChannelBWs_DL_v1590Present {
				ie.ChannelBWs_DL_v1590 = new(BandNR_channelBWs_DL_v1590)
				if err = ie.ChannelBWs_DL_v1590.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelBWs_DL_v1590", err)
				}
			}
			// decode ChannelBWs_UL_v1590 optional
			if ChannelBWs_UL_v1590Present {
				ie.ChannelBWs_UL_v1590 = new(BandNR_channelBWs_UL_v1590)
				if err = ie.ChannelBWs_UL_v1590.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelBWs_UL_v1590", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			AsymmetricBandwidthCombinationSetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AsymmetricBandwidthCombinationSet optional
			if AsymmetricBandwidthCombinationSetPresent {
				var tmp_bs_AsymmetricBandwidthCombinationSet []byte
				var n_AsymmetricBandwidthCombinationSet uint
				if tmp_bs_AsymmetricBandwidthCombinationSet, n_AsymmetricBandwidthCombinationSet, err = extReader.ReadBitString(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode AsymmetricBandwidthCombinationSet", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_AsymmetricBandwidthCombinationSet,
					NumBits: uint64(n_AsymmetricBandwidthCombinationSet),
				}
				ie.AsymmetricBandwidthCombinationSet = &tmp_bitstring
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SharedSpectrumChAccessParamsPerBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CancelOverlappingPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultipleRateMatchingEUTRA_CRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OverlapRateMatchingEUTRA_CRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_MappingTypeB_Alt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OneSlotPeriodicTRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Olpc_SRS_Pos_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationsSRS_Pos_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimulSRS_MIMO_TransWithinBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBW_DL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBW_UL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			RasterShift7dot5_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ue_PowerClass_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondHandover_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondHandoverFailure_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondHandoverTwoTriggerEvents_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondPSCellChange_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondPSCellChangeTwoTriggerEvents_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mpr_PowerBoost_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ActiveConfiguredGrant_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			JointReleaseConfiguredGrantType2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			JointReleaseSPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimulSRS_TransWithinBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Trs_AdditionalBandwidth_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverIntraF_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SharedSpectrumChAccessParamsPerBand_r16 optional
			if SharedSpectrumChAccessParamsPerBand_r16Present {
				ie.SharedSpectrumChAccessParamsPerBand_r16 = new(SharedSpectrumChAccessParamsPerBand_r16)
				if err = ie.SharedSpectrumChAccessParamsPerBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_r16", err)
				}
			}
			// decode CancelOverlappingPUSCH_r16 optional
			if CancelOverlappingPUSCH_r16Present {
				ie.CancelOverlappingPUSCH_r16 = new(BandNR_cancelOverlappingPUSCH_r16)
				if err = ie.CancelOverlappingPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CancelOverlappingPUSCH_r16", err)
				}
			}
			// decode MultipleRateMatchingEUTRA_CRS_r16 optional
			if MultipleRateMatchingEUTRA_CRS_r16Present {
				ie.MultipleRateMatchingEUTRA_CRS_r16 = new(BandNR_multipleRateMatchingEUTRA_CRS_r16)
				if err = ie.MultipleRateMatchingEUTRA_CRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultipleRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// decode OverlapRateMatchingEUTRA_CRS_r16 optional
			if OverlapRateMatchingEUTRA_CRS_r16Present {
				ie.OverlapRateMatchingEUTRA_CRS_r16 = new(BandNR_overlapRateMatchingEUTRA_CRS_r16)
				if err = ie.OverlapRateMatchingEUTRA_CRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode OverlapRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// decode Pdsch_MappingTypeB_Alt_r16 optional
			if Pdsch_MappingTypeB_Alt_r16Present {
				ie.Pdsch_MappingTypeB_Alt_r16 = new(BandNR_pdsch_MappingTypeB_Alt_r16)
				if err = ie.Pdsch_MappingTypeB_Alt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_MappingTypeB_Alt_r16", err)
				}
			}
			// decode OneSlotPeriodicTRS_r16 optional
			if OneSlotPeriodicTRS_r16Present {
				ie.OneSlotPeriodicTRS_r16 = new(BandNR_oneSlotPeriodicTRS_r16)
				if err = ie.OneSlotPeriodicTRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode OneSlotPeriodicTRS_r16", err)
				}
			}
			// decode Olpc_SRS_Pos_r16 optional
			if Olpc_SRS_Pos_r16Present {
				ie.Olpc_SRS_Pos_r16 = new(OLPC_SRS_Pos_r16)
				if err = ie.Olpc_SRS_Pos_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Olpc_SRS_Pos_r16", err)
				}
			}
			// decode SpatialRelationsSRS_Pos_r16 optional
			if SpatialRelationsSRS_Pos_r16Present {
				ie.SpatialRelationsSRS_Pos_r16 = new(SpatialRelationsSRS_Pos_r16)
				if err = ie.SpatialRelationsSRS_Pos_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelationsSRS_Pos_r16", err)
				}
			}
			// decode SimulSRS_MIMO_TransWithinBand_r16 optional
			if SimulSRS_MIMO_TransWithinBand_r16Present {
				ie.SimulSRS_MIMO_TransWithinBand_r16 = new(BandNR_simulSRS_MIMO_TransWithinBand_r16)
				if err = ie.SimulSRS_MIMO_TransWithinBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimulSRS_MIMO_TransWithinBand_r16", err)
				}
			}
			// decode ChannelBW_DL_IAB_r16 optional
			if ChannelBW_DL_IAB_r16Present {
				ie.ChannelBW_DL_IAB_r16 = new(BandNR_channelBW_DL_IAB_r16)
				if err = ie.ChannelBW_DL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelBW_DL_IAB_r16", err)
				}
			}
			// decode ChannelBW_UL_IAB_r16 optional
			if ChannelBW_UL_IAB_r16Present {
				ie.ChannelBW_UL_IAB_r16 = new(BandNR_channelBW_UL_IAB_r16)
				if err = ie.ChannelBW_UL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChannelBW_UL_IAB_r16", err)
				}
			}
			// decode RasterShift7dot5_IAB_r16 optional
			if RasterShift7dot5_IAB_r16Present {
				ie.RasterShift7dot5_IAB_r16 = new(BandNR_rasterShift7dot5_IAB_r16)
				if err = ie.RasterShift7dot5_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode RasterShift7dot5_IAB_r16", err)
				}
			}
			// decode Ue_PowerClass_v1610 optional
			if Ue_PowerClass_v1610Present {
				ie.Ue_PowerClass_v1610 = new(BandNR_ue_PowerClass_v1610)
				if err = ie.Ue_PowerClass_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ue_PowerClass_v1610", err)
				}
			}
			// decode CondHandover_r16 optional
			if CondHandover_r16Present {
				ie.CondHandover_r16 = new(BandNR_condHandover_r16)
				if err = ie.CondHandover_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondHandover_r16", err)
				}
			}
			// decode CondHandoverFailure_r16 optional
			if CondHandoverFailure_r16Present {
				ie.CondHandoverFailure_r16 = new(BandNR_condHandoverFailure_r16)
				if err = ie.CondHandoverFailure_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondHandoverFailure_r16", err)
				}
			}
			// decode CondHandoverTwoTriggerEvents_r16 optional
			if CondHandoverTwoTriggerEvents_r16Present {
				ie.CondHandoverTwoTriggerEvents_r16 = new(BandNR_condHandoverTwoTriggerEvents_r16)
				if err = ie.CondHandoverTwoTriggerEvents_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondHandoverTwoTriggerEvents_r16", err)
				}
			}
			// decode CondPSCellChange_r16 optional
			if CondPSCellChange_r16Present {
				ie.CondPSCellChange_r16 = new(BandNR_condPSCellChange_r16)
				if err = ie.CondPSCellChange_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondPSCellChange_r16", err)
				}
			}
			// decode CondPSCellChangeTwoTriggerEvents_r16 optional
			if CondPSCellChangeTwoTriggerEvents_r16Present {
				ie.CondPSCellChangeTwoTriggerEvents_r16 = new(BandNR_condPSCellChangeTwoTriggerEvents_r16)
				if err = ie.CondPSCellChangeTwoTriggerEvents_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondPSCellChangeTwoTriggerEvents_r16", err)
				}
			}
			// decode Mpr_PowerBoost_FR2_r16 optional
			if Mpr_PowerBoost_FR2_r16Present {
				ie.Mpr_PowerBoost_FR2_r16 = new(BandNR_mpr_PowerBoost_FR2_r16)
				if err = ie.Mpr_PowerBoost_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mpr_PowerBoost_FR2_r16", err)
				}
			}
			// decode ActiveConfiguredGrant_r16 optional
			if ActiveConfiguredGrant_r16Present {
				ie.ActiveConfiguredGrant_r16 = new(BandNR_activeConfiguredGrant_r16)
				if err = ie.ActiveConfiguredGrant_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ActiveConfiguredGrant_r16", err)
				}
			}
			// decode JointReleaseConfiguredGrantType2_r16 optional
			if JointReleaseConfiguredGrantType2_r16Present {
				ie.JointReleaseConfiguredGrantType2_r16 = new(BandNR_jointReleaseConfiguredGrantType2_r16)
				if err = ie.JointReleaseConfiguredGrantType2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode JointReleaseConfiguredGrantType2_r16", err)
				}
			}
			// decode Sps_r16 optional
			if Sps_r16Present {
				ie.Sps_r16 = new(BandNR_sps_r16)
				if err = ie.Sps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sps_r16", err)
				}
			}
			// decode JointReleaseSPS_r16 optional
			if JointReleaseSPS_r16Present {
				ie.JointReleaseSPS_r16 = new(BandNR_jointReleaseSPS_r16)
				if err = ie.JointReleaseSPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode JointReleaseSPS_r16", err)
				}
			}
			// decode SimulSRS_TransWithinBand_r16 optional
			if SimulSRS_TransWithinBand_r16Present {
				ie.SimulSRS_TransWithinBand_r16 = new(BandNR_simulSRS_TransWithinBand_r16)
				if err = ie.SimulSRS_TransWithinBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimulSRS_TransWithinBand_r16", err)
				}
			}
			// decode Trs_AdditionalBandwidth_r16 optional
			if Trs_AdditionalBandwidth_r16Present {
				ie.Trs_AdditionalBandwidth_r16 = new(BandNR_trs_AdditionalBandwidth_r16)
				if err = ie.Trs_AdditionalBandwidth_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Trs_AdditionalBandwidth_r16", err)
				}
			}
			// decode HandoverIntraF_IAB_r16 optional
			if HandoverIntraF_IAB_r16Present {
				ie.HandoverIntraF_IAB_r16 = new(BandNR_handoverIntraF_IAB_r16)
				if err = ie.HandoverIntraF_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverIntraF_IAB_r16", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SharedSpectrumChAccessParamsPerBand_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 optional
			if SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16Present {
				ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
				if err = ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16", err)
				}
			}
			// decode SharedSpectrumChAccessParamsPerBand_v1630 optional
			if SharedSpectrumChAccessParamsPerBand_v1630Present {
				ie.SharedSpectrumChAccessParamsPerBand_v1630 = new(SharedSpectrumChAccessParamsPerBand_v1630)
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_v1630", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			HandoverUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedUL_TransientPeriod_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SharedSpectrumChAccessParamsPerBand_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode HandoverUTRA_FDD_r16 optional
			if HandoverUTRA_FDD_r16Present {
				ie.HandoverUTRA_FDD_r16 = new(BandNR_handoverUTRA_FDD_r16)
				if err = ie.HandoverUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverUTRA_FDD_r16", err)
				}
			}
			// decode EnhancedUL_TransientPeriod_r16 optional
			if EnhancedUL_TransientPeriod_r16Present {
				ie.EnhancedUL_TransientPeriod_r16 = new(BandNR_enhancedUL_TransientPeriod_r16)
				if err = ie.EnhancedUL_TransientPeriod_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedUL_TransientPeriod_r16", err)
				}
			}
			// decode SharedSpectrumChAccessParamsPerBand_v1640 optional
			if SharedSpectrumChAccessParamsPerBand_v1640Present {
				ie.SharedSpectrumChAccessParamsPerBand_v1640 = new(SharedSpectrumChAccessParamsPerBand_v1640)
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_v1640", err)
				}
			}
		}
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Type1_PUSCH_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type2_PUSCH_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredUL_GrantType1_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConfiguredUL_GrantType2_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SharedSpectrumChAccessParamsPerBand_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Type1_PUSCH_RepetitionMultiSlots_v1650 optional
			if Type1_PUSCH_RepetitionMultiSlots_v1650Present {
				ie.Type1_PUSCH_RepetitionMultiSlots_v1650 = new(BandNR_type1_PUSCH_RepetitionMultiSlots_v1650)
				if err = ie.Type1_PUSCH_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type1_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode Type2_PUSCH_RepetitionMultiSlots_v1650 optional
			if Type2_PUSCH_RepetitionMultiSlots_v1650Present {
				ie.Type2_PUSCH_RepetitionMultiSlots_v1650 = new(BandNR_type2_PUSCH_RepetitionMultiSlots_v1650)
				if err = ie.Type2_PUSCH_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type2_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode Pusch_RepetitionMultiSlots_v1650 optional
			if Pusch_RepetitionMultiSlots_v1650Present {
				ie.Pusch_RepetitionMultiSlots_v1650 = new(BandNR_pusch_RepetitionMultiSlots_v1650)
				if err = ie.Pusch_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode ConfiguredUL_GrantType1_v1650 optional
			if ConfiguredUL_GrantType1_v1650Present {
				ie.ConfiguredUL_GrantType1_v1650 = new(BandNR_configuredUL_GrantType1_v1650)
				if err = ie.ConfiguredUL_GrantType1_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredUL_GrantType1_v1650", err)
				}
			}
			// decode ConfiguredUL_GrantType2_v1650 optional
			if ConfiguredUL_GrantType2_v1650Present {
				ie.ConfiguredUL_GrantType2_v1650 = new(BandNR_configuredUL_GrantType2_v1650)
				if err = ie.ConfiguredUL_GrantType2_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConfiguredUL_GrantType2_v1650", err)
				}
			}
			// decode SharedSpectrumChAccessParamsPerBand_v1650 optional
			if SharedSpectrumChAccessParamsPerBand_v1650Present {
				ie.SharedSpectrumChAccessParamsPerBand_v1650 = new(SharedSpectrumChAccessParamsPerBand_v1650)
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_v1650", err)
				}
			}
		}
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			EnhancedSkipUplinkTxConfigured_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedSkipUplinkTxDynamic_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EnhancedSkipUplinkTxConfigured_v1660 optional
			if EnhancedSkipUplinkTxConfigured_v1660Present {
				ie.EnhancedSkipUplinkTxConfigured_v1660 = new(BandNR_enhancedSkipUplinkTxConfigured_v1660)
				if err = ie.EnhancedSkipUplinkTxConfigured_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxConfigured_v1660", err)
				}
			}
			// decode EnhancedSkipUplinkTxDynamic_v1660 optional
			if EnhancedSkipUplinkTxDynamic_v1660Present {
				ie.EnhancedSkipUplinkTxDynamic_v1660 = new(BandNR_enhancedSkipUplinkTxDynamic_v1660)
				if err = ie.EnhancedSkipUplinkTxDynamic_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedSkipUplinkTxDynamic_v1660", err)
				}
			}
		}
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TxDiversity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 optional
			if MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16Present {
				ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 = new(BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16)
				if err = ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
				}
			}
			// decode TxDiversity_r16 optional
			if TxDiversity_r16Present {
				ie.TxDiversity_r16 = new(BandNR_txDiversity_r16)
				if err = ie.TxDiversity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode TxDiversity_r16", err)
				}
			}
		}
		// decode extension group 12
		if len(extBitmap) > 11 && extBitmap[11] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Pdsch_1024QAM_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ue_PowerClass_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Fr2_2_AccessParamsPerBand_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rlm_Relaxation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Bfd_Relaxation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cg_SDT_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LocationBasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TimeBasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EventA4BasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mn_InitiatedCondPSCellChangeNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sn_InitiatedCondPSCellChangeNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_SkippingWithoutSSSG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sssg_Switching_1BitInd_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sssg_Switching_2BitInd_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdcch_SkippingWithSSSG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SearchSpaceSetGrp_switchCap2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkPreCompensation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uplink_TA_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Max_HARQ_ProcessNumber_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type2_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type1_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Type3_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ue_specific_K_Offset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ParallelPRS_MeasRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_UE_TxTEG_ID_MaxSupport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_ProcessingRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_ProcessingWindowType1A_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_ProcessingWindowType1B_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_ProcessingWindowType2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_AllPosResourcesRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Olpc_SRS_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationsSRS_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberPUSCH_TypeA_Repetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PuschTypeA_RepetitionsAvailSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tb_ProcessingMultiSlotPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Tb_ProcessingRepMultiSlotPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxDurationDMRS_Bundling_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pusch_RepetitionMsg3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SharedSpectrumChAccessParamsPerBand_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ParallelMeasurementWithoutRestriction_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumber_NGSO_SatellitesWithinOneSMTC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			K1_RangeExtension_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicCSI_RS_FastScellActivation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicCSI_RS_AdditionalBandwidth_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HalfDuplexFDD_TypeA_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_DL_SCS_480kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_UL_SCS_480kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_DL_SCS_960kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_UL_SCS_960kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_GapFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OneShotHARQ_feedbackTriggeredByDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OneShotHARQ_feedbackPhy_Priority_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedType3_HARQ_CodebookFeedback_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TriggeredHARQ_CodebookRetx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdsch_1024QAM_FR1_r17 optional
			if Pdsch_1024QAM_FR1_r17Present {
				ie.Pdsch_1024QAM_FR1_r17 = new(BandNR_pdsch_1024QAM_FR1_r17)
				if err = ie.Pdsch_1024QAM_FR1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_1024QAM_FR1_r17", err)
				}
			}
			// decode Ue_PowerClass_v1700 optional
			if Ue_PowerClass_v1700Present {
				ie.Ue_PowerClass_v1700 = new(BandNR_ue_PowerClass_v1700)
				if err = ie.Ue_PowerClass_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ue_PowerClass_v1700", err)
				}
			}
			// decode Fr2_2_AccessParamsPerBand_r17 optional
			if Fr2_2_AccessParamsPerBand_r17Present {
				ie.Fr2_2_AccessParamsPerBand_r17 = new(FR2_2_AccessParamsPerBand_r17)
				if err = ie.Fr2_2_AccessParamsPerBand_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Fr2_2_AccessParamsPerBand_r17", err)
				}
			}
			// decode Rlm_Relaxation_r17 optional
			if Rlm_Relaxation_r17Present {
				ie.Rlm_Relaxation_r17 = new(BandNR_rlm_Relaxation_r17)
				if err = ie.Rlm_Relaxation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rlm_Relaxation_r17", err)
				}
			}
			// decode Bfd_Relaxation_r17 optional
			if Bfd_Relaxation_r17Present {
				ie.Bfd_Relaxation_r17 = new(BandNR_bfd_Relaxation_r17)
				if err = ie.Bfd_Relaxation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Bfd_Relaxation_r17", err)
				}
			}
			// decode Cg_SDT_r17 optional
			if Cg_SDT_r17Present {
				ie.Cg_SDT_r17 = new(BandNR_cg_SDT_r17)
				if err = ie.Cg_SDT_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cg_SDT_r17", err)
				}
			}
			// decode LocationBasedCondHandover_r17 optional
			if LocationBasedCondHandover_r17Present {
				ie.LocationBasedCondHandover_r17 = new(BandNR_locationBasedCondHandover_r17)
				if err = ie.LocationBasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode LocationBasedCondHandover_r17", err)
				}
			}
			// decode TimeBasedCondHandover_r17 optional
			if TimeBasedCondHandover_r17Present {
				ie.TimeBasedCondHandover_r17 = new(BandNR_timeBasedCondHandover_r17)
				if err = ie.TimeBasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TimeBasedCondHandover_r17", err)
				}
			}
			// decode EventA4BasedCondHandover_r17 optional
			if EventA4BasedCondHandover_r17Present {
				ie.EventA4BasedCondHandover_r17 = new(BandNR_eventA4BasedCondHandover_r17)
				if err = ie.EventA4BasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EventA4BasedCondHandover_r17", err)
				}
			}
			// decode Mn_InitiatedCondPSCellChangeNRDC_r17 optional
			if Mn_InitiatedCondPSCellChangeNRDC_r17Present {
				ie.Mn_InitiatedCondPSCellChangeNRDC_r17 = new(BandNR_mn_InitiatedCondPSCellChangeNRDC_r17)
				if err = ie.Mn_InitiatedCondPSCellChangeNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// decode Sn_InitiatedCondPSCellChangeNRDC_r17 optional
			if Sn_InitiatedCondPSCellChangeNRDC_r17Present {
				ie.Sn_InitiatedCondPSCellChangeNRDC_r17 = new(BandNR_sn_InitiatedCondPSCellChangeNRDC_r17)
				if err = ie.Sn_InitiatedCondPSCellChangeNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// decode Pdcch_SkippingWithoutSSSG_r17 optional
			if Pdcch_SkippingWithoutSSSG_r17Present {
				ie.Pdcch_SkippingWithoutSSSG_r17 = new(BandNR_pdcch_SkippingWithoutSSSG_r17)
				if err = ie.Pdcch_SkippingWithoutSSSG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_SkippingWithoutSSSG_r17", err)
				}
			}
			// decode Sssg_Switching_1BitInd_r17 optional
			if Sssg_Switching_1BitInd_r17Present {
				ie.Sssg_Switching_1BitInd_r17 = new(BandNR_sssg_Switching_1BitInd_r17)
				if err = ie.Sssg_Switching_1BitInd_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sssg_Switching_1BitInd_r17", err)
				}
			}
			// decode Sssg_Switching_2BitInd_r17 optional
			if Sssg_Switching_2BitInd_r17Present {
				ie.Sssg_Switching_2BitInd_r17 = new(BandNR_sssg_Switching_2BitInd_r17)
				if err = ie.Sssg_Switching_2BitInd_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sssg_Switching_2BitInd_r17", err)
				}
			}
			// decode Pdcch_SkippingWithSSSG_r17 optional
			if Pdcch_SkippingWithSSSG_r17Present {
				ie.Pdcch_SkippingWithSSSG_r17 = new(BandNR_pdcch_SkippingWithSSSG_r17)
				if err = ie.Pdcch_SkippingWithSSSG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_SkippingWithSSSG_r17", err)
				}
			}
			// decode SearchSpaceSetGrp_switchCap2_r17 optional
			if SearchSpaceSetGrp_switchCap2_r17Present {
				ie.SearchSpaceSetGrp_switchCap2_r17 = new(BandNR_searchSpaceSetGrp_switchCap2_r17)
				if err = ie.SearchSpaceSetGrp_switchCap2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SearchSpaceSetGrp_switchCap2_r17", err)
				}
			}
			// decode UplinkPreCompensation_r17 optional
			if UplinkPreCompensation_r17Present {
				ie.UplinkPreCompensation_r17 = new(BandNR_uplinkPreCompensation_r17)
				if err = ie.UplinkPreCompensation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkPreCompensation_r17", err)
				}
			}
			// decode Uplink_TA_Reporting_r17 optional
			if Uplink_TA_Reporting_r17Present {
				ie.Uplink_TA_Reporting_r17 = new(BandNR_uplink_TA_Reporting_r17)
				if err = ie.Uplink_TA_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Uplink_TA_Reporting_r17", err)
				}
			}
			// decode Max_HARQ_ProcessNumber_r17 optional
			if Max_HARQ_ProcessNumber_r17Present {
				ie.Max_HARQ_ProcessNumber_r17 = new(BandNR_max_HARQ_ProcessNumber_r17)
				if err = ie.Max_HARQ_ProcessNumber_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Max_HARQ_ProcessNumber_r17", err)
				}
			}
			// decode Type2_HARQ_Codebook_r17 optional
			if Type2_HARQ_Codebook_r17Present {
				ie.Type2_HARQ_Codebook_r17 = new(BandNR_type2_HARQ_Codebook_r17)
				if err = ie.Type2_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type2_HARQ_Codebook_r17", err)
				}
			}
			// decode Type1_HARQ_Codebook_r17 optional
			if Type1_HARQ_Codebook_r17Present {
				ie.Type1_HARQ_Codebook_r17 = new(BandNR_type1_HARQ_Codebook_r17)
				if err = ie.Type1_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type1_HARQ_Codebook_r17", err)
				}
			}
			// decode Type3_HARQ_Codebook_r17 optional
			if Type3_HARQ_Codebook_r17Present {
				ie.Type3_HARQ_Codebook_r17 = new(BandNR_type3_HARQ_Codebook_r17)
				if err = ie.Type3_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type3_HARQ_Codebook_r17", err)
				}
			}
			// decode Ue_specific_K_Offset_r17 optional
			if Ue_specific_K_Offset_r17Present {
				ie.Ue_specific_K_Offset_r17 = new(BandNR_ue_specific_K_Offset_r17)
				if err = ie.Ue_specific_K_Offset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ue_specific_K_Offset_r17", err)
				}
			}
			// decode MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present {
				ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = new(BandNR_multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17)
				if err = ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// decode MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present {
				ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = new(BandNR_multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17)
				if err = ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// decode ParallelPRS_MeasRRC_Inactive_r17 optional
			if ParallelPRS_MeasRRC_Inactive_r17Present {
				ie.ParallelPRS_MeasRRC_Inactive_r17 = new(BandNR_parallelPRS_MeasRRC_Inactive_r17)
				if err = ie.ParallelPRS_MeasRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ParallelPRS_MeasRRC_Inactive_r17", err)
				}
			}
			// decode Nr_UE_TxTEG_ID_MaxSupport_r17 optional
			if Nr_UE_TxTEG_ID_MaxSupport_r17Present {
				ie.Nr_UE_TxTEG_ID_MaxSupport_r17 = new(BandNR_nr_UE_TxTEG_ID_MaxSupport_r17)
				if err = ie.Nr_UE_TxTEG_ID_MaxSupport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_UE_TxTEG_ID_MaxSupport_r17", err)
				}
			}
			// decode Prs_ProcessingRRC_Inactive_r17 optional
			if Prs_ProcessingRRC_Inactive_r17Present {
				ie.Prs_ProcessingRRC_Inactive_r17 = new(BandNR_prs_ProcessingRRC_Inactive_r17)
				if err = ie.Prs_ProcessingRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prs_ProcessingRRC_Inactive_r17", err)
				}
			}
			// decode Prs_ProcessingWindowType1A_r17 optional
			if Prs_ProcessingWindowType1A_r17Present {
				ie.Prs_ProcessingWindowType1A_r17 = new(BandNR_prs_ProcessingWindowType1A_r17)
				if err = ie.Prs_ProcessingWindowType1A_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prs_ProcessingWindowType1A_r17", err)
				}
			}
			// decode Prs_ProcessingWindowType1B_r17 optional
			if Prs_ProcessingWindowType1B_r17Present {
				ie.Prs_ProcessingWindowType1B_r17 = new(BandNR_prs_ProcessingWindowType1B_r17)
				if err = ie.Prs_ProcessingWindowType1B_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prs_ProcessingWindowType1B_r17", err)
				}
			}
			// decode Prs_ProcessingWindowType2_r17 optional
			if Prs_ProcessingWindowType2_r17Present {
				ie.Prs_ProcessingWindowType2_r17 = new(BandNR_prs_ProcessingWindowType2_r17)
				if err = ie.Prs_ProcessingWindowType2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prs_ProcessingWindowType2_r17", err)
				}
			}
			// decode Srs_AllPosResourcesRRC_Inactive_r17 optional
			if Srs_AllPosResourcesRRC_Inactive_r17Present {
				ie.Srs_AllPosResourcesRRC_Inactive_r17 = new(SRS_AllPosResourcesRRC_Inactive_r17)
				if err = ie.Srs_AllPosResourcesRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_AllPosResourcesRRC_Inactive_r17", err)
				}
			}
			// decode Olpc_SRS_PosRRC_Inactive_r17 optional
			if Olpc_SRS_PosRRC_Inactive_r17Present {
				ie.Olpc_SRS_PosRRC_Inactive_r17 = new(OLPC_SRS_Pos_r16)
				if err = ie.Olpc_SRS_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Olpc_SRS_PosRRC_Inactive_r17", err)
				}
			}
			// decode SpatialRelationsSRS_PosRRC_Inactive_r17 optional
			if SpatialRelationsSRS_PosRRC_Inactive_r17Present {
				ie.SpatialRelationsSRS_PosRRC_Inactive_r17 = new(SpatialRelationsSRS_Pos_r16)
				if err = ie.SpatialRelationsSRS_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelationsSRS_PosRRC_Inactive_r17", err)
				}
			}
			// decode MaxNumberPUSCH_TypeA_Repetition_r17 optional
			if MaxNumberPUSCH_TypeA_Repetition_r17Present {
				ie.MaxNumberPUSCH_TypeA_Repetition_r17 = new(BandNR_maxNumberPUSCH_TypeA_Repetition_r17)
				if err = ie.MaxNumberPUSCH_TypeA_Repetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberPUSCH_TypeA_Repetition_r17", err)
				}
			}
			// decode PuschTypeA_RepetitionsAvailSlot_r17 optional
			if PuschTypeA_RepetitionsAvailSlot_r17Present {
				ie.PuschTypeA_RepetitionsAvailSlot_r17 = new(BandNR_puschTypeA_RepetitionsAvailSlot_r17)
				if err = ie.PuschTypeA_RepetitionsAvailSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PuschTypeA_RepetitionsAvailSlot_r17", err)
				}
			}
			// decode Tb_ProcessingMultiSlotPUSCH_r17 optional
			if Tb_ProcessingMultiSlotPUSCH_r17Present {
				ie.Tb_ProcessingMultiSlotPUSCH_r17 = new(BandNR_tb_ProcessingMultiSlotPUSCH_r17)
				if err = ie.Tb_ProcessingMultiSlotPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tb_ProcessingMultiSlotPUSCH_r17", err)
				}
			}
			// decode Tb_ProcessingRepMultiSlotPUSCH_r17 optional
			if Tb_ProcessingRepMultiSlotPUSCH_r17Present {
				ie.Tb_ProcessingRepMultiSlotPUSCH_r17 = new(BandNR_tb_ProcessingRepMultiSlotPUSCH_r17)
				if err = ie.Tb_ProcessingRepMultiSlotPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Tb_ProcessingRepMultiSlotPUSCH_r17", err)
				}
			}
			// decode MaxDurationDMRS_Bundling_r17 optional
			if MaxDurationDMRS_Bundling_r17Present {
				ie.MaxDurationDMRS_Bundling_r17 = new(BandNR_maxDurationDMRS_Bundling_r17)
				if err = ie.MaxDurationDMRS_Bundling_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxDurationDMRS_Bundling_r17", err)
				}
			}
			// decode Pusch_RepetitionMsg3_r17 optional
			if Pusch_RepetitionMsg3_r17Present {
				ie.Pusch_RepetitionMsg3_r17 = new(BandNR_pusch_RepetitionMsg3_r17)
				if err = ie.Pusch_RepetitionMsg3_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pusch_RepetitionMsg3_r17", err)
				}
			}
			// decode SharedSpectrumChAccessParamsPerBand_v1710 optional
			if SharedSpectrumChAccessParamsPerBand_v1710Present {
				ie.SharedSpectrumChAccessParamsPerBand_v1710 = new(SharedSpectrumChAccessParamsPerBand_v1710)
				if err = ie.SharedSpectrumChAccessParamsPerBand_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode SharedSpectrumChAccessParamsPerBand_v1710", err)
				}
			}
			// decode ParallelMeasurementWithoutRestriction_r17 optional
			if ParallelMeasurementWithoutRestriction_r17Present {
				ie.ParallelMeasurementWithoutRestriction_r17 = new(BandNR_parallelMeasurementWithoutRestriction_r17)
				if err = ie.ParallelMeasurementWithoutRestriction_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ParallelMeasurementWithoutRestriction_r17", err)
				}
			}
			// decode MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 optional
			if MaxNumber_NGSO_SatellitesWithinOneSMTC_r17Present {
				ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 = new(BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17)
				if err = ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
				}
			}
			// decode K1_RangeExtension_r17 optional
			if K1_RangeExtension_r17Present {
				ie.K1_RangeExtension_r17 = new(BandNR_k1_RangeExtension_r17)
				if err = ie.K1_RangeExtension_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode K1_RangeExtension_r17", err)
				}
			}
			// decode AperiodicCSI_RS_FastScellActivation_r17 optional
			if AperiodicCSI_RS_FastScellActivation_r17Present {
				ie.AperiodicCSI_RS_FastScellActivation_r17 = new(BandNR_aperiodicCSI_RS_FastScellActivation_r17)
				if err = ie.AperiodicCSI_RS_FastScellActivation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AperiodicCSI_RS_FastScellActivation_r17", err)
				}
			}
			// decode AperiodicCSI_RS_AdditionalBandwidth_r17 optional
			if AperiodicCSI_RS_AdditionalBandwidth_r17Present {
				ie.AperiodicCSI_RS_AdditionalBandwidth_r17 = new(BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17)
				if err = ie.AperiodicCSI_RS_AdditionalBandwidth_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode AperiodicCSI_RS_AdditionalBandwidth_r17", err)
				}
			}
			// decode Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 optional
			if Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17Present {
				ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 = new(BandNR_bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17)
				if err = ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17", err)
				}
			}
			// decode HalfDuplexFDD_TypeA_RedCap_r17 optional
			if HalfDuplexFDD_TypeA_RedCap_r17Present {
				ie.HalfDuplexFDD_TypeA_RedCap_r17 = new(BandNR_halfDuplexFDD_TypeA_RedCap_r17)
				if err = ie.HalfDuplexFDD_TypeA_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode HalfDuplexFDD_TypeA_RedCap_r17", err)
				}
			}
			// decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 optional
			if PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17Present {
				ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17)
				if err = ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17", err)
				}
			}
			// decode ChannelBWs_DL_SCS_480kHz_FR2_2_r17 optional
			if ChannelBWs_DL_SCS_480kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_DL_SCS_480kHz_FR2_2_r17 []byte
				var n_ChannelBWs_DL_SCS_480kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_DL_SCS_480kHz_FR2_2_r17, n_ChannelBWs_DL_SCS_480kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_DL_SCS_480kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_DL_SCS_480kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_DL_SCS_480kHz_FR2_2_r17),
				}
				ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode ChannelBWs_UL_SCS_480kHz_FR2_2_r17 optional
			if ChannelBWs_UL_SCS_480kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_UL_SCS_480kHz_FR2_2_r17 []byte
				var n_ChannelBWs_UL_SCS_480kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_UL_SCS_480kHz_FR2_2_r17, n_ChannelBWs_UL_SCS_480kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_UL_SCS_480kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_UL_SCS_480kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_UL_SCS_480kHz_FR2_2_r17),
				}
				ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode ChannelBWs_DL_SCS_960kHz_FR2_2_r17 optional
			if ChannelBWs_DL_SCS_960kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_DL_SCS_960kHz_FR2_2_r17 []byte
				var n_ChannelBWs_DL_SCS_960kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_DL_SCS_960kHz_FR2_2_r17, n_ChannelBWs_DL_SCS_960kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_DL_SCS_960kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_DL_SCS_960kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_DL_SCS_960kHz_FR2_2_r17),
				}
				ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode ChannelBWs_UL_SCS_960kHz_FR2_2_r17 optional
			if ChannelBWs_UL_SCS_960kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_UL_SCS_960kHz_FR2_2_r17 []byte
				var n_ChannelBWs_UL_SCS_960kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_UL_SCS_960kHz_FR2_2_r17, n_ChannelBWs_UL_SCS_960kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_UL_SCS_960kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_UL_SCS_960kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_UL_SCS_960kHz_FR2_2_r17),
				}
				ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode Ul_GapFR2_r17 optional
			if Ul_GapFR2_r17Present {
				ie.Ul_GapFR2_r17 = new(BandNR_ul_GapFR2_r17)
				if err = ie.Ul_GapFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_GapFR2_r17", err)
				}
			}
			// decode OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 optional
			if OneShotHARQ_feedbackTriggeredByDCI_1_2_r17Present {
				ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17 = new(BandNR_oneShotHARQ_feedbackTriggeredByDCI_1_2_r17)
				if err = ie.OneShotHARQ_feedbackTriggeredByDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode OneShotHARQ_feedbackTriggeredByDCI_1_2_r17", err)
				}
			}
			// decode OneShotHARQ_feedbackPhy_Priority_r17 optional
			if OneShotHARQ_feedbackPhy_Priority_r17Present {
				ie.OneShotHARQ_feedbackPhy_Priority_r17 = new(BandNR_oneShotHARQ_feedbackPhy_Priority_r17)
				if err = ie.OneShotHARQ_feedbackPhy_Priority_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode OneShotHARQ_feedbackPhy_Priority_r17", err)
				}
			}
			// decode EnhancedType3_HARQ_CodebookFeedback_r17 optional
			if EnhancedType3_HARQ_CodebookFeedback_r17Present {
				ie.EnhancedType3_HARQ_CodebookFeedback_r17 = new(BandNR_enhancedType3_HARQ_CodebookFeedback_r17)
				if err = ie.EnhancedType3_HARQ_CodebookFeedback_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedType3_HARQ_CodebookFeedback_r17", err)
				}
			}
			// decode TriggeredHARQ_CodebookRetx_r17 optional
			if TriggeredHARQ_CodebookRetx_r17Present {
				ie.TriggeredHARQ_CodebookRetx_r17 = new(BandNR_triggeredHARQ_CodebookRetx_r17)
				if err = ie.TriggeredHARQ_CodebookRetx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TriggeredHARQ_CodebookRetx_r17", err)
				}
			}
		}
		// decode extension group 13
		if len(extBitmap) > 12 && extBitmap[12] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Ue_OneShotUL_TimingAdj_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pucch_Repetition_F0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mux_HARQ_ACK_DiffPriorities_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberG_RNTI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DynamicMulticastDCI_Format4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxModulationOrderForMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberG_CS_RNTI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Re_LevelRateMatchingForMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_1024QAM_2MIMO_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_MeasurementWithoutMG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumber_LEO_SatellitesPerCarrier_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Prs_ProcessingCapabilityOutsideMGinPPW_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_SemiPersistent_PosResourcesRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_DL_SCS_120kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChannelBWs_UL_SCS_120kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Ue_OneShotUL_TimingAdj_r17 optional
			if Ue_OneShotUL_TimingAdj_r17Present {
				ie.Ue_OneShotUL_TimingAdj_r17 = new(BandNR_ue_OneShotUL_TimingAdj_r17)
				if err = ie.Ue_OneShotUL_TimingAdj_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ue_OneShotUL_TimingAdj_r17", err)
				}
			}
			// decode Pucch_Repetition_F0_2_r17 optional
			if Pucch_Repetition_F0_2_r17Present {
				ie.Pucch_Repetition_F0_2_r17 = new(BandNR_pucch_Repetition_F0_2_r17)
				if err = ie.Pucch_Repetition_F0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pucch_Repetition_F0_2_r17", err)
				}
			}
			// decode Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 optional
			if Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17Present {
				ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 = new(BandNR_cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17)
				if err = ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode Mux_HARQ_ACK_DiffPriorities_r17 optional
			if Mux_HARQ_ACK_DiffPriorities_r17Present {
				ie.Mux_HARQ_ACK_DiffPriorities_r17 = new(BandNR_mux_HARQ_ACK_DiffPriorities_r17)
				if err = ie.Mux_HARQ_ACK_DiffPriorities_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_HARQ_ACK_DiffPriorities_r17", err)
				}
			}
			// decode Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 optional
			if Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17Present {
				ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 = new(BandNR_ta_BasedPDC_NTN_SharedSpectrumChAccess_r17)
				if err = ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 optional
			if Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17Present {
				ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 = new(BandNR_ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17)
				if err = ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode MaxNumberG_RNTI_r17 optional
			if MaxNumberG_RNTI_r17Present {
				var tmp_int_MaxNumberG_RNTI_r17 int64
				if tmp_int_MaxNumberG_RNTI_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode MaxNumberG_RNTI_r17", err)
				}
				ie.MaxNumberG_RNTI_r17 = &tmp_int_MaxNumberG_RNTI_r17
			}
			// decode DynamicMulticastDCI_Format4_2_r17 optional
			if DynamicMulticastDCI_Format4_2_r17Present {
				ie.DynamicMulticastDCI_Format4_2_r17 = new(BandNR_dynamicMulticastDCI_Format4_2_r17)
				if err = ie.DynamicMulticastDCI_Format4_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DynamicMulticastDCI_Format4_2_r17", err)
				}
			}
			// decode MaxModulationOrderForMulticast_r17 optional
			if MaxModulationOrderForMulticast_r17Present {
				ie.MaxModulationOrderForMulticast_r17 = new(BandNR_maxModulationOrderForMulticast_r17)
				if err = ie.MaxModulationOrderForMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxModulationOrderForMulticast_r17", err)
				}
			}
			// decode DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 optional
			if DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17Present {
				ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 = new(BandNR_dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17)
				if err = ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// decode DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 optional
			if DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17Present {
				ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 = new(BandNR_dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17)
				if err = ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 optional
			if Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17Present {
				ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 = new(BandNR_nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17)
				if err = ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 optional
			if Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17Present {
				ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 = new(BandNR_ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17)
				if err = ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode MaxNumberG_CS_RNTI_r17 optional
			if MaxNumberG_CS_RNTI_r17Present {
				var tmp_int_MaxNumberG_CS_RNTI_r17 int64
				if tmp_int_MaxNumberG_CS_RNTI_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode MaxNumberG_CS_RNTI_r17", err)
				}
				ie.MaxNumberG_CS_RNTI_r17 = &tmp_int_MaxNumberG_CS_RNTI_r17
			}
			// decode Re_LevelRateMatchingForMulticast_r17 optional
			if Re_LevelRateMatchingForMulticast_r17Present {
				ie.Re_LevelRateMatchingForMulticast_r17 = new(BandNR_re_LevelRateMatchingForMulticast_r17)
				if err = ie.Re_LevelRateMatchingForMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Re_LevelRateMatchingForMulticast_r17", err)
				}
			}
			// decode Pdsch_1024QAM_2MIMO_FR1_r17 optional
			if Pdsch_1024QAM_2MIMO_FR1_r17Present {
				ie.Pdsch_1024QAM_2MIMO_FR1_r17 = new(BandNR_pdsch_1024QAM_2MIMO_FR1_r17)
				if err = ie.Pdsch_1024QAM_2MIMO_FR1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_1024QAM_2MIMO_FR1_r17", err)
				}
			}
			// decode Prs_MeasurementWithoutMG_r17 optional
			if Prs_MeasurementWithoutMG_r17Present {
				ie.Prs_MeasurementWithoutMG_r17 = new(BandNR_prs_MeasurementWithoutMG_r17)
				if err = ie.Prs_MeasurementWithoutMG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Prs_MeasurementWithoutMG_r17", err)
				}
			}
			// decode MaxNumber_LEO_SatellitesPerCarrier_r17 optional
			if MaxNumber_LEO_SatellitesPerCarrier_r17Present {
				var tmp_int_MaxNumber_LEO_SatellitesPerCarrier_r17 int64
				if tmp_int_MaxNumber_LEO_SatellitesPerCarrier_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 3, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode MaxNumber_LEO_SatellitesPerCarrier_r17", err)
				}
				ie.MaxNumber_LEO_SatellitesPerCarrier_r17 = &tmp_int_MaxNumber_LEO_SatellitesPerCarrier_r17
			}
			// decode Prs_ProcessingCapabilityOutsideMGinPPW_r17 optional
			if Prs_ProcessingCapabilityOutsideMGinPPW_r17Present {
				tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17 := utils.NewSequence[*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17]([]*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}, aper.Constraint{Lb: 1, Ub: 3}, false)
				fn_Prs_ProcessingCapabilityOutsideMGinPPW_r17 := func() *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17 {
					return new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17)
				}
				if err = tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17.Decode(extReader, fn_Prs_ProcessingCapabilityOutsideMGinPPW_r17); err != nil {
					return utils.WrapError("Decode Prs_ProcessingCapabilityOutsideMGinPPW_r17", err)
				}
				ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 = []PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}
				for _, i := range tmp_Prs_ProcessingCapabilityOutsideMGinPPW_r17.Value {
					ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 = append(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17, *i)
				}
			}
			// decode Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 optional
			if Srs_SemiPersistent_PosResourcesRRC_Inactive_r17Present {
				ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 = new(BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17)
				if err = ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_SemiPersistent_PosResourcesRRC_Inactive_r17", err)
				}
			}
			// decode ChannelBWs_DL_SCS_120kHz_FR2_2_r17 optional
			if ChannelBWs_DL_SCS_120kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_DL_SCS_120kHz_FR2_2_r17 []byte
				var n_ChannelBWs_DL_SCS_120kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_DL_SCS_120kHz_FR2_2_r17, n_ChannelBWs_DL_SCS_120kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_DL_SCS_120kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_DL_SCS_120kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_DL_SCS_120kHz_FR2_2_r17),
				}
				ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode ChannelBWs_UL_SCS_120kHz_FR2_2_r17 optional
			if ChannelBWs_UL_SCS_120kHz_FR2_2_r17Present {
				var tmp_bs_ChannelBWs_UL_SCS_120kHz_FR2_2_r17 []byte
				var n_ChannelBWs_UL_SCS_120kHz_FR2_2_r17 uint
				if tmp_bs_ChannelBWs_UL_SCS_120kHz_FR2_2_r17, n_ChannelBWs_UL_SCS_120kHz_FR2_2_r17, err = extReader.ReadBitString(&aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode ChannelBWs_UL_SCS_120kHz_FR2_2_r17", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_ChannelBWs_UL_SCS_120kHz_FR2_2_r17,
					NumBits: uint64(n_ChannelBWs_UL_SCS_120kHz_FR2_2_r17),
				}
				ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 = &tmp_bitstring
			}
		}
		// decode extension group 14
		if len(extBitmap) > 13 && extBitmap[13] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Dmrs_BundlingPUSCH_RepTypeA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingPUSCH_RepTypeB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingPUSCH_multiSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingPUCCH_Rep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InterSlotFreqHopInterSlotBundlingPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InterSlotFreqHopPUCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingRestart_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dmrs_BundlingNonBackToBackTX_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dmrs_BundlingPUSCH_RepTypeA_r17 optional
			if Dmrs_BundlingPUSCH_RepTypeA_r17Present {
				ie.Dmrs_BundlingPUSCH_RepTypeA_r17 = new(BandNR_dmrs_BundlingPUSCH_RepTypeA_r17)
				if err = ie.Dmrs_BundlingPUSCH_RepTypeA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUSCH_RepTypeA_r17", err)
				}
			}
			// decode Dmrs_BundlingPUSCH_RepTypeB_r17 optional
			if Dmrs_BundlingPUSCH_RepTypeB_r17Present {
				ie.Dmrs_BundlingPUSCH_RepTypeB_r17 = new(BandNR_dmrs_BundlingPUSCH_RepTypeB_r17)
				if err = ie.Dmrs_BundlingPUSCH_RepTypeB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUSCH_RepTypeB_r17", err)
				}
			}
			// decode Dmrs_BundlingPUSCH_multiSlot_r17 optional
			if Dmrs_BundlingPUSCH_multiSlot_r17Present {
				ie.Dmrs_BundlingPUSCH_multiSlot_r17 = new(BandNR_dmrs_BundlingPUSCH_multiSlot_r17)
				if err = ie.Dmrs_BundlingPUSCH_multiSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUSCH_multiSlot_r17", err)
				}
			}
			// decode Dmrs_BundlingPUCCH_Rep_r17 optional
			if Dmrs_BundlingPUCCH_Rep_r17Present {
				ie.Dmrs_BundlingPUCCH_Rep_r17 = new(BandNR_dmrs_BundlingPUCCH_Rep_r17)
				if err = ie.Dmrs_BundlingPUCCH_Rep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingPUCCH_Rep_r17", err)
				}
			}
			// decode InterSlotFreqHopInterSlotBundlingPUSCH_r17 optional
			if InterSlotFreqHopInterSlotBundlingPUSCH_r17Present {
				ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 = new(BandNR_interSlotFreqHopInterSlotBundlingPUSCH_r17)
				if err = ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterSlotFreqHopInterSlotBundlingPUSCH_r17", err)
				}
			}
			// decode InterSlotFreqHopPUCCH_r17 optional
			if InterSlotFreqHopPUCCH_r17Present {
				ie.InterSlotFreqHopPUCCH_r17 = new(BandNR_interSlotFreqHopPUCCH_r17)
				if err = ie.InterSlotFreqHopPUCCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterSlotFreqHopPUCCH_r17", err)
				}
			}
			// decode Dmrs_BundlingRestart_r17 optional
			if Dmrs_BundlingRestart_r17Present {
				ie.Dmrs_BundlingRestart_r17 = new(BandNR_dmrs_BundlingRestart_r17)
				if err = ie.Dmrs_BundlingRestart_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingRestart_r17", err)
				}
			}
			// decode Dmrs_BundlingNonBackToBackTX_r17 optional
			if Dmrs_BundlingNonBackToBackTX_r17Present {
				ie.Dmrs_BundlingNonBackToBackTX_r17 = new(BandNR_dmrs_BundlingNonBackToBackTX_r17)
				if err = ie.Dmrs_BundlingNonBackToBackTX_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dmrs_BundlingNonBackToBackTX_r17", err)
				}
			}
		}
	}
	return nil
}
