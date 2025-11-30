package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand struct {
	Tci_StatePDSCH                               *MIMO_ParametersPerBand_tci_StatePDSCH                               `optional`
	AdditionalActiveTCI_StatePDCCH               *MIMO_ParametersPerBand_additionalActiveTCI_StatePDCCH               `optional`
	Pusch_TransCoherence                         *MIMO_ParametersPerBand_pusch_TransCoherence                         `optional`
	BeamCorrespondenceWithoutUL_BeamSweeping     *MIMO_ParametersPerBand_beamCorrespondenceWithoutUL_BeamSweeping     `optional`
	PeriodicBeamReport                           *MIMO_ParametersPerBand_periodicBeamReport                           `optional`
	AperiodicBeamReport                          *MIMO_ParametersPerBand_aperiodicBeamReport                          `optional`
	Sp_BeamReportPUCCH                           *MIMO_ParametersPerBand_sp_BeamReportPUCCH                           `optional`
	Sp_BeamReportPUSCH                           *MIMO_ParametersPerBand_sp_BeamReportPUSCH                           `optional`
	Dummy1                                       *DummyG                                                              `optional`
	MaxNumberRxBeam                              *int64                                                               `lb:2,ub:8,optional`
	MaxNumberRxTxBeamSwitchDL                    *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL                    `optional`
	MaxNumberNonGroupBeamReporting               *MIMO_ParametersPerBand_maxNumberNonGroupBeamReporting               `optional`
	GroupBeamReporting                           *MIMO_ParametersPerBand_groupBeamReporting                           `optional`
	UplinkBeamManagement                         *MIMO_ParametersPerBand_uplinkBeamManagement                         `lb:1,ub:8,optional`
	MaxNumberCSI_RS_BFD                          *int64                                                               `lb:1,ub:64,optional`
	MaxNumberSSB_BFD                             *int64                                                               `lb:1,ub:64,optional`
	MaxNumberCSI_RS_SSB_CBD                      *int64                                                               `lb:1,ub:256,optional`
	Dummy2                                       *MIMO_ParametersPerBand_dummy2                                       `optional`
	TwoPortsPTRS_UL                              *MIMO_ParametersPerBand_twoPortsPTRS_UL                              `optional`
	Dummy5                                       *SRS_Resources                                                       `optional`
	Dummy3                                       *int64                                                               `lb:1,ub:4,optional`
	BeamReportTiming                             *MIMO_ParametersPerBand_beamReportTiming                             `optional`
	Ptrs_DensityRecommendationSetDL              *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL              `optional`
	Ptrs_DensityRecommendationSetUL              *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetUL              `optional`
	Dummy4                                       *DummyH                                                              `optional`
	AperiodicTRS                                 *MIMO_ParametersPerBand_aperiodicTRS                                 `optional`
	Dummy6                                       *MIMO_ParametersPerBand_dummy6                                       `optional,ext-1`
	BeamManagementSSB_CSI_RS                     *BeamManagementSSB_CSI_RS                                            `optional,ext-1`
	BeamSwitchTiming                             *MIMO_ParametersPerBand_beamSwitchTiming                             `optional,ext-1`
	CodebookParameters                           *CodebookParameters                                                  `optional,ext-1`
	Csi_RS_IM_ReceptionForFeedback               *CSI_RS_IM_ReceptionForFeedback                                      `optional,ext-1`
	Csi_RS_ProcFrameworkForSRS                   *CSI_RS_ProcFrameworkForSRS                                          `optional,ext-1`
	Csi_ReportFramework                          *CSI_ReportFramework                                                 `optional,ext-1`
	Csi_RS_ForTracking                           *CSI_RS_ForTracking                                                  `optional,ext-1`
	Srs_AssocCSI_RS                              []SupportedCSI_RS_Resource                                           `lb:1,ub:maxNrofCSI_RS_Resources,optional,ext-1`
	SpatialRelations                             *SpatialRelations                                                    `optional,ext-1`
	DefaultQCL_TwoTCI_r16                        *MIMO_ParametersPerBand_defaultQCL_TwoTCI_r16                        `optional,ext-2`
	CodebookParametersPerBand_r16                *CodebookParameters_v1610                                            `optional,ext-2`
	Simul_SpatialRelationUpdatePUCCHResGroup_r16 *MIMO_ParametersPerBand_simul_SpatialRelationUpdatePUCCHResGroup_r16 `optional,ext-2`
	MaxNumberSCellBFR_r16                        *MIMO_ParametersPerBand_maxNumberSCellBFR_r16                        `optional,ext-2`
	SimultaneousReceptionDiffTypeD_r16           *MIMO_ParametersPerBand_simultaneousReceptionDiffTypeD_r16           `optional,ext-2`
	Ssb_csirs_SINR_measurement_r16               *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16               `optional,ext-2`
	NonGroupSINR_reporting_r16                   *MIMO_ParametersPerBand_nonGroupSINR_reporting_r16                   `optional,ext-2`
	GroupSINR_reporting_r16                      *MIMO_ParametersPerBand_groupSINR_reporting_r16                      `optional,ext-2`
	MultiDCI_multiTRP_Parameters_r16             *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16             `lb:1,ub:2,optional,ext-2`
	SingleDCI_SDM_scheme_Parameters_r16          *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16          `optional,ext-2`
	SupportFDM_SchemeA_r16                       *MIMO_ParametersPerBand_supportFDM_SchemeA_r16                       `optional,ext-2`
	SupportCodeWordSoftCombining_r16             *MIMO_ParametersPerBand_supportCodeWordSoftCombining_r16             `optional,ext-2`
	SupportTDM_SchemeA_r16                       *MIMO_ParametersPerBand_supportTDM_SchemeA_r16                       `optional,ext-2`
	SupportInter_slotTDM_r16                     *MIMO_ParametersPerBand_supportInter_slotTDM_r16                     `lb:1,ub:2,optional,ext-2`
	LowPAPR_DMRS_PDSCH_r16                       *MIMO_ParametersPerBand_lowPAPR_DMRS_PDSCH_r16                       `optional,ext-2`
	LowPAPR_DMRS_PUSCHwithoutPrecoding_r16       *MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithoutPrecoding_r16       `optional,ext-2`
	LowPAPR_DMRS_PUCCH_r16                       *MIMO_ParametersPerBand_lowPAPR_DMRS_PUCCH_r16                       `optional,ext-2`
	LowPAPR_DMRS_PUSCHwithPrecoding_r16          *MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithPrecoding_r16          `optional,ext-2`
	Csi_ReportFrameworkExt_r16                   *CSI_ReportFrameworkExt_r16                                          `optional,ext-2`
	CodebookParametersAddition_r16               *CodebookParametersAddition_r16                                      `optional,ext-2`
	CodebookComboParametersAddition_r16          *CodebookComboParametersAddition_r16                                 `optional,ext-2`
	BeamCorrespondenceSSB_based_r16              *MIMO_ParametersPerBand_beamCorrespondenceSSB_based_r16              `optional,ext-2`
	BeamCorrespondenceCSI_RS_based_r16           *MIMO_ParametersPerBand_beamCorrespondenceCSI_RS_based_r16           `optional,ext-2`
	BeamSwitchTiming_r16                         *MIMO_ParametersPerBand_beamSwitchTiming_r16                         `optional,ext-2`
	Semi_PersistentL1_SINR_Report_PUCCH_r16      *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16      `optional,ext-3`
	Semi_PersistentL1_SINR_Report_PUSCH_r16      *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUSCH_r16      `optional,ext-3`
	SpatialRelations_v1640                       *MIMO_ParametersPerBand_spatialRelations_v1640                       `optional,ext-4`
	Support64CandidateBeamRS_BFR_r16             *MIMO_ParametersPerBand_support64CandidateBeamRS_BFR_r16             `optional,ext-4`
	MaxMIMO_LayersForMulti_DCI_mTRP_r16          *MIMO_ParametersPerBand_maxMIMO_LayersForMulti_DCI_mTRP_r16          `optional,ext-5`
	SupportedSINR_meas_v1670                     *aper.BitString                                                      `lb:4,ub:4,optional,ext-6`
	Srs_increasedRepetition_r17                  *MIMO_ParametersPerBand_srs_increasedRepetition_r17                  `optional,ext-7`
	Srs_partialFrequencySounding_r17             *MIMO_ParametersPerBand_srs_partialFrequencySounding_r17             `optional,ext-7`
	Srs_startRB_locationHoppingPartial_r17       *MIMO_ParametersPerBand_srs_startRB_locationHoppingPartial_r17       `optional,ext-7`
	Srs_combEight_r17                            *MIMO_ParametersPerBand_srs_combEight_r17                            `optional,ext-7`
	CodebookParametersfetype2_r17                *CodebookParametersfetype2_r17                                       `optional,ext-7`
	MTRP_PUSCH_twoCSI_RS_r17                     *MIMO_ParametersPerBand_mTRP_PUSCH_twoCSI_RS_r17                     `optional,ext-7`
	MTRP_PUCCH_InterSlot_r17                     *MIMO_ParametersPerBand_mTRP_PUCCH_InterSlot_r17                     `optional,ext-7`
	MTRP_PUCCH_CyclicMapping_r17                 *MIMO_ParametersPerBand_mTRP_PUCCH_CyclicMapping_r17                 `optional,ext-7`
	MTRP_PUCCH_SecondTPC_r17                     *MIMO_ParametersPerBand_mTRP_PUCCH_SecondTPC_r17                     `optional,ext-7`
	MTRP_BFR_twoBFD_RS_Set_r17                   *MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17                   `lb:1,ub:9,optional,ext-7`
	MTRP_BFR_PUCCH_SR_perCG_r17                  *MIMO_ParametersPerBand_mTRP_BFR_PUCCH_SR_perCG_r17                  `optional,ext-7`
	MTRP_BFR_association_PUCCH_SR_r17            *MIMO_ParametersPerBand_mTRP_BFR_association_PUCCH_SR_r17            `optional,ext-7`
	Sfn_SimulTwoTCI_AcrossMultiCC_r17            *MIMO_ParametersPerBand_sfn_SimulTwoTCI_AcrossMultiCC_r17            `optional,ext-7`
	Sfn_DefaultDL_BeamSetup_r17                  *MIMO_ParametersPerBand_sfn_DefaultDL_BeamSetup_r17                  `optional,ext-7`
	Sfn_DefaultUL_BeamSetup_r17                  *MIMO_ParametersPerBand_sfn_DefaultUL_BeamSetup_r17                  `optional,ext-7`
	Srs_TriggeringOffset_r17                     *MIMO_ParametersPerBand_srs_TriggeringOffset_r17                     `optional,ext-7`
	Srs_TriggeringDCI_r17                        *MIMO_ParametersPerBand_srs_TriggeringDCI_r17                        `optional,ext-7`
	CodebookComboParameterMixedType_r17          *CodebookComboParameterMixedType_r17                                 `optional,ext-7`
	UnifiedJointTCI_r17                          *MIMO_ParametersPerBand_unifiedJointTCI_r17                          `optional,ext-7`
	UnifiedJointTCI_multiMAC_CE_r17              *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17              `optional,ext-7`
	UnifiedJointTCI_perBWP_CA_r17                *MIMO_ParametersPerBand_unifiedJointTCI_perBWP_CA_r17                `optional,ext-7`
	UnifiedJointTCI_ListSharingCA_r17            *MIMO_ParametersPerBand_unifiedJointTCI_ListSharingCA_r17            `optional,ext-7`
	UnifiedJointTCI_commonMultiCC_r17            *MIMO_ParametersPerBand_unifiedJointTCI_commonMultiCC_r17            `optional,ext-7`
	UnifiedJointTCI_BeamAlignDLRS_r17            *MIMO_ParametersPerBand_unifiedJointTCI_BeamAlignDLRS_r17            `optional,ext-7`
	UnifiedJointTCI_PC_association_r17           *MIMO_ParametersPerBand_unifiedJointTCI_PC_association_r17           `optional,ext-7`
	UnifiedJointTCI_Legacy_r17                   *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_r17                   `optional,ext-7`
	UnifiedJointTCI_Legacy_SRS_r17               *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_SRS_r17               `optional,ext-7`
	UnifiedJointTCI_Legacy_CORESET0_r17          *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_CORESET0_r17          `optional,ext-7`
	UnifiedJointTCI_SCellBFR_r17                 *MIMO_ParametersPerBand_unifiedJointTCI_SCellBFR_r17                 `optional,ext-7`
	UnifiedJointTCI_InterCell_r17                *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17                `optional,ext-7`
	UnifiedSeparateTCI_r17                       *MIMO_ParametersPerBand_unifiedSeparateTCI_r17                       `optional,ext-7`
	UnifiedSeparateTCI_multiMAC_CE_r17           *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17           `lb:2,ub:8,optional,ext-7`
	UnifiedSeparateTCI_perBWP_CA_r17             *MIMO_ParametersPerBand_unifiedSeparateTCI_perBWP_CA_r17             `optional,ext-7`
	UnifiedSeparateTCI_ListSharingCA_r17         *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17         `optional,ext-7`
	UnifiedSeparateTCI_commonMultiCC_r17         *MIMO_ParametersPerBand_unifiedSeparateTCI_commonMultiCC_r17         `optional,ext-7`
	UnifiedSeparateTCI_InterCell_r17             *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17             `optional,ext-7`
	UnifiedJointTCI_mTRP_InterCell_BM_r17        *MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17        `lb:1,ub:7,optional,ext-7`
	Mpe_Mitigation_r17                           *MIMO_ParametersPerBand_mpe_Mitigation_r17                           `lb:1,ub:4,optional,ext-7`
	Srs_PortReport_r17                           *MIMO_ParametersPerBand_srs_PortReport_r17                           `optional,ext-7`
	MTRP_PDCCH_individual_r17                    *MIMO_ParametersPerBand_mTRP_PDCCH_individual_r17                    `optional,ext-7`
	MTRP_PDCCH_anySpan_3Symbols_r17              *MIMO_ParametersPerBand_mTRP_PDCCH_anySpan_3Symbols_r17              `optional,ext-7`
	MTRP_PDCCH_TwoQCL_TypeD_r17                  *MIMO_ParametersPerBand_mTRP_PDCCH_TwoQCL_TypeD_r17                  `optional,ext-7`
	MTRP_PUSCH_CSI_RS_r17                        *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17                        `lb:1,ub:8,optional,ext-7`
	MTRP_PUSCH_cyclicMapping_r17                 *MIMO_ParametersPerBand_mTRP_PUSCH_cyclicMapping_r17                 `optional,ext-7`
	MTRP_PUSCH_secondTPC_r17                     *MIMO_ParametersPerBand_mTRP_PUSCH_secondTPC_r17                     `optional,ext-7`
	MTRP_PUSCH_twoPHR_Reporting_r17              *MIMO_ParametersPerBand_mTRP_PUSCH_twoPHR_Reporting_r17              `optional,ext-7`
	MTRP_PUSCH_A_CSI_r17                         *MIMO_ParametersPerBand_mTRP_PUSCH_A_CSI_r17                         `optional,ext-7`
	MTRP_PUSCH_SP_CSI_r17                        *MIMO_ParametersPerBand_mTRP_PUSCH_SP_CSI_r17                        `optional,ext-7`
	MTRP_PUSCH_CG_r17                            *MIMO_ParametersPerBand_mTRP_PUSCH_CG_r17                            `optional,ext-7`
	MTRP_PUCCH_MAC_CE_r17                        *MIMO_ParametersPerBand_mTRP_PUCCH_MAC_CE_r17                        `optional,ext-7`
	MTRP_PUCCH_maxNum_PC_FR1_r17                 *int64                                                               `lb:3,ub:8,optional,ext-7`
	MTRP_inter_Cell_r17                          *MIMO_ParametersPerBand_mTRP_inter_Cell_r17                          `lb:1,ub:7,optional,ext-7`
	MTRP_GroupBasedL1_RSRP_r17                   *MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17                   `lb:1,ub:4,optional,ext-7`
	MTRP_BFD_RS_MAC_CE_r17                       *MIMO_ParametersPerBand_mTRP_BFD_RS_MAC_CE_r17                       `optional,ext-7`
	MTRP_CSI_EnhancementPerBand_r17              *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17              `lb:1,ub:16,optional,ext-7`
	CodebookComboParameterMultiTRP_r17           *CodebookComboParameterMultiTRP_r17                                  `optional,ext-7`
	MTRP_CSI_additionalCSI_r17                   *MIMO_ParametersPerBand_mTRP_CSI_additionalCSI_r17                   `optional,ext-7`
	MTRP_CSI_N_Max2_r17                          *MIMO_ParametersPerBand_mTRP_CSI_N_Max2_r17                          `optional,ext-7`
	MTRP_CSI_CMR_r17                             *MIMO_ParametersPerBand_mTRP_CSI_CMR_r17                             `optional,ext-7`
	Srs_partialFreqSounding_r17                  *MIMO_ParametersPerBand_srs_partialFreqSounding_r17                  `optional,ext-7`
	BeamSwitchTiming_v1710                       *MIMO_ParametersPerBand_beamSwitchTiming_v1710                       `optional,ext-7`
	BeamSwitchTiming_r17                         *MIMO_ParametersPerBand_beamSwitchTiming_r17                         `optional,ext-7`
	BeamReportTiming_v1710                       *MIMO_ParametersPerBand_beamReportTiming_v1710                       `optional,ext-7`
	MaxNumberRxTxBeamSwitchDL_v1710              *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710              `optional,ext-7`
	Srs_PortReportSP_AP_r17                      *MIMO_ParametersPerBand_srs_PortReportSP_AP_r17                      `optional,ext-8`
	MaxNumberRxBeam_v1720                        *int64                                                               `lb:9,ub:12,optional,ext-8`
	Sfn_ImplicitRS_twoTCI_r17                    *MIMO_ParametersPerBand_sfn_ImplicitRS_twoTCI_r17                    `optional,ext-8`
	Sfn_QCL_TypeD_Collision_twoTCI_r17           *MIMO_ParametersPerBand_sfn_QCL_TypeD_Collision_twoTCI_r17           `optional,ext-8`
	MTRP_CSI_numCPU_r17                          *MIMO_ParametersPerBand_mTRP_CSI_numCPU_r17                          `optional,ext-8`
	SupportRepNumPDSCH_TDRA_DCI_1_2_r17          *MIMO_ParametersPerBand_supportRepNumPDSCH_TDRA_DCI_1_2_r17          `optional,ext-9`
}

func (ie *MIMO_ParametersPerBand) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Dummy6 != nil || ie.BeamManagementSSB_CSI_RS != nil || ie.BeamSwitchTiming != nil || ie.CodebookParameters != nil || ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Csi_RS_ForTracking != nil || len(ie.Srs_AssocCSI_RS) > 0 || ie.SpatialRelations != nil || ie.DefaultQCL_TwoTCI_r16 != nil || ie.CodebookParametersPerBand_r16 != nil || ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil || ie.MaxNumberSCellBFR_r16 != nil || ie.SimultaneousReceptionDiffTypeD_r16 != nil || ie.Ssb_csirs_SINR_measurement_r16 != nil || ie.NonGroupSINR_reporting_r16 != nil || ie.GroupSINR_reporting_r16 != nil || ie.MultiDCI_multiTRP_Parameters_r16 != nil || ie.SingleDCI_SDM_scheme_Parameters_r16 != nil || ie.SupportFDM_SchemeA_r16 != nil || ie.SupportCodeWordSoftCombining_r16 != nil || ie.SupportTDM_SchemeA_r16 != nil || ie.SupportInter_slotTDM_r16 != nil || ie.LowPAPR_DMRS_PDSCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil || ie.LowPAPR_DMRS_PUCCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil || ie.CodebookParametersAddition_r16 != nil || ie.CodebookComboParametersAddition_r16 != nil || ie.BeamCorrespondenceSSB_based_r16 != nil || ie.BeamCorrespondenceCSI_RS_based_r16 != nil || ie.BeamSwitchTiming_r16 != nil || ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil || ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil || ie.SpatialRelations_v1640 != nil || ie.Support64CandidateBeamRS_BFR_r16 != nil || ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16 != nil || ie.SupportedSINR_meas_v1670 != nil || ie.Srs_increasedRepetition_r17 != nil || ie.Srs_partialFrequencySounding_r17 != nil || ie.Srs_startRB_locationHoppingPartial_r17 != nil || ie.Srs_combEight_r17 != nil || ie.CodebookParametersfetype2_r17 != nil || ie.MTRP_PUSCH_twoCSI_RS_r17 != nil || ie.MTRP_PUCCH_InterSlot_r17 != nil || ie.MTRP_PUCCH_CyclicMapping_r17 != nil || ie.MTRP_PUCCH_SecondTPC_r17 != nil || ie.MTRP_BFR_twoBFD_RS_Set_r17 != nil || ie.MTRP_BFR_PUCCH_SR_perCG_r17 != nil || ie.MTRP_BFR_association_PUCCH_SR_r17 != nil || ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil || ie.Sfn_DefaultDL_BeamSetup_r17 != nil || ie.Sfn_DefaultUL_BeamSetup_r17 != nil || ie.Srs_TriggeringOffset_r17 != nil || ie.Srs_TriggeringDCI_r17 != nil || ie.CodebookComboParameterMixedType_r17 != nil || ie.UnifiedJointTCI_r17 != nil || ie.UnifiedJointTCI_multiMAC_CE_r17 != nil || ie.UnifiedJointTCI_perBWP_CA_r17 != nil || ie.UnifiedJointTCI_ListSharingCA_r17 != nil || ie.UnifiedJointTCI_commonMultiCC_r17 != nil || ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil || ie.UnifiedJointTCI_PC_association_r17 != nil || ie.UnifiedJointTCI_Legacy_r17 != nil || ie.UnifiedJointTCI_Legacy_SRS_r17 != nil || ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil || ie.UnifiedJointTCI_SCellBFR_r17 != nil || ie.UnifiedJointTCI_InterCell_r17 != nil || ie.UnifiedSeparateTCI_r17 != nil || ie.UnifiedSeparateTCI_multiMAC_CE_r17 != nil || ie.UnifiedSeparateTCI_perBWP_CA_r17 != nil || ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil || ie.UnifiedSeparateTCI_commonMultiCC_r17 != nil || ie.UnifiedSeparateTCI_InterCell_r17 != nil || ie.UnifiedJointTCI_mTRP_InterCell_BM_r17 != nil || ie.Mpe_Mitigation_r17 != nil || ie.Srs_PortReport_r17 != nil || ie.MTRP_PDCCH_individual_r17 != nil || ie.MTRP_PDCCH_anySpan_3Symbols_r17 != nil || ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil || ie.MTRP_PUSCH_CSI_RS_r17 != nil || ie.MTRP_PUSCH_cyclicMapping_r17 != nil || ie.MTRP_PUSCH_secondTPC_r17 != nil || ie.MTRP_PUSCH_twoPHR_Reporting_r17 != nil || ie.MTRP_PUSCH_A_CSI_r17 != nil || ie.MTRP_PUSCH_SP_CSI_r17 != nil || ie.MTRP_PUSCH_CG_r17 != nil || ie.MTRP_PUCCH_MAC_CE_r17 != nil || ie.MTRP_PUCCH_maxNum_PC_FR1_r17 != nil || ie.MTRP_inter_Cell_r17 != nil || ie.MTRP_GroupBasedL1_RSRP_r17 != nil || ie.MTRP_BFD_RS_MAC_CE_r17 != nil || ie.MTRP_CSI_EnhancementPerBand_r17 != nil || ie.CodebookComboParameterMultiTRP_r17 != nil || ie.MTRP_CSI_additionalCSI_r17 != nil || ie.MTRP_CSI_N_Max2_r17 != nil || ie.MTRP_CSI_CMR_r17 != nil || ie.Srs_partialFreqSounding_r17 != nil || ie.BeamSwitchTiming_v1710 != nil || ie.BeamSwitchTiming_r17 != nil || ie.BeamReportTiming_v1710 != nil || ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil || ie.Srs_PortReportSP_AP_r17 != nil || ie.MaxNumberRxBeam_v1720 != nil || ie.Sfn_ImplicitRS_twoTCI_r17 != nil || ie.Sfn_QCL_TypeD_Collision_twoTCI_r17 != nil || ie.MTRP_CSI_numCPU_r17 != nil || ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Tci_StatePDSCH != nil, ie.AdditionalActiveTCI_StatePDCCH != nil, ie.Pusch_TransCoherence != nil, ie.BeamCorrespondenceWithoutUL_BeamSweeping != nil, ie.PeriodicBeamReport != nil, ie.AperiodicBeamReport != nil, ie.Sp_BeamReportPUCCH != nil, ie.Sp_BeamReportPUSCH != nil, ie.Dummy1 != nil, ie.MaxNumberRxBeam != nil, ie.MaxNumberRxTxBeamSwitchDL != nil, ie.MaxNumberNonGroupBeamReporting != nil, ie.GroupBeamReporting != nil, ie.UplinkBeamManagement != nil, ie.MaxNumberCSI_RS_BFD != nil, ie.MaxNumberSSB_BFD != nil, ie.MaxNumberCSI_RS_SSB_CBD != nil, ie.Dummy2 != nil, ie.TwoPortsPTRS_UL != nil, ie.Dummy5 != nil, ie.Dummy3 != nil, ie.BeamReportTiming != nil, ie.Ptrs_DensityRecommendationSetDL != nil, ie.Ptrs_DensityRecommendationSetUL != nil, ie.Dummy4 != nil, ie.AperiodicTRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Tci_StatePDSCH != nil {
		if err = ie.Tci_StatePDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Tci_StatePDSCH", err)
		}
	}
	if ie.AdditionalActiveTCI_StatePDCCH != nil {
		if err = ie.AdditionalActiveTCI_StatePDCCH.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalActiveTCI_StatePDCCH", err)
		}
	}
	if ie.Pusch_TransCoherence != nil {
		if err = ie.Pusch_TransCoherence.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_TransCoherence", err)
		}
	}
	if ie.BeamCorrespondenceWithoutUL_BeamSweeping != nil {
		if err = ie.BeamCorrespondenceWithoutUL_BeamSweeping.Encode(w); err != nil {
			return utils.WrapError("Encode BeamCorrespondenceWithoutUL_BeamSweeping", err)
		}
	}
	if ie.PeriodicBeamReport != nil {
		if err = ie.PeriodicBeamReport.Encode(w); err != nil {
			return utils.WrapError("Encode PeriodicBeamReport", err)
		}
	}
	if ie.AperiodicBeamReport != nil {
		if err = ie.AperiodicBeamReport.Encode(w); err != nil {
			return utils.WrapError("Encode AperiodicBeamReport", err)
		}
	}
	if ie.Sp_BeamReportPUCCH != nil {
		if err = ie.Sp_BeamReportPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_BeamReportPUCCH", err)
		}
	}
	if ie.Sp_BeamReportPUSCH != nil {
		if err = ie.Sp_BeamReportPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_BeamReportPUSCH", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = ie.Dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.MaxNumberRxBeam != nil {
		if err = w.WriteInteger(*ie.MaxNumberRxBeam, &aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode MaxNumberRxBeam", err)
		}
	}
	if ie.MaxNumberRxTxBeamSwitchDL != nil {
		if err = ie.MaxNumberRxTxBeamSwitchDL.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberRxTxBeamSwitchDL", err)
		}
	}
	if ie.MaxNumberNonGroupBeamReporting != nil {
		if err = ie.MaxNumberNonGroupBeamReporting.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberNonGroupBeamReporting", err)
		}
	}
	if ie.GroupBeamReporting != nil {
		if err = ie.GroupBeamReporting.Encode(w); err != nil {
			return utils.WrapError("Encode GroupBeamReporting", err)
		}
	}
	if ie.UplinkBeamManagement != nil {
		if err = ie.UplinkBeamManagement.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkBeamManagement", err)
		}
	}
	if ie.MaxNumberCSI_RS_BFD != nil {
		if err = w.WriteInteger(*ie.MaxNumberCSI_RS_BFD, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode MaxNumberCSI_RS_BFD", err)
		}
	}
	if ie.MaxNumberSSB_BFD != nil {
		if err = w.WriteInteger(*ie.MaxNumberSSB_BFD, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode MaxNumberSSB_BFD", err)
		}
	}
	if ie.MaxNumberCSI_RS_SSB_CBD != nil {
		if err = w.WriteInteger(*ie.MaxNumberCSI_RS_SSB_CBD, &aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode MaxNumberCSI_RS_SSB_CBD", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = ie.Dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	if ie.TwoPortsPTRS_UL != nil {
		if err = ie.TwoPortsPTRS_UL.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPortsPTRS_UL", err)
		}
	}
	if ie.Dummy5 != nil {
		if err = ie.Dummy5.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy5", err)
		}
	}
	if ie.Dummy3 != nil {
		if err = w.WriteInteger(*ie.Dummy3, &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode Dummy3", err)
		}
	}
	if ie.BeamReportTiming != nil {
		if err = ie.BeamReportTiming.Encode(w); err != nil {
			return utils.WrapError("Encode BeamReportTiming", err)
		}
	}
	if ie.Ptrs_DensityRecommendationSetDL != nil {
		if err = ie.Ptrs_DensityRecommendationSetDL.Encode(w); err != nil {
			return utils.WrapError("Encode Ptrs_DensityRecommendationSetDL", err)
		}
	}
	if ie.Ptrs_DensityRecommendationSetUL != nil {
		if err = ie.Ptrs_DensityRecommendationSetUL.Encode(w); err != nil {
			return utils.WrapError("Encode Ptrs_DensityRecommendationSetUL", err)
		}
	}
	if ie.Dummy4 != nil {
		if err = ie.Dummy4.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy4", err)
		}
	}
	if ie.AperiodicTRS != nil {
		if err = ie.AperiodicTRS.Encode(w); err != nil {
			return utils.WrapError("Encode AperiodicTRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 9 bits for 9 extension groups
		extBitmap := []bool{ie.Dummy6 != nil || ie.BeamManagementSSB_CSI_RS != nil || ie.BeamSwitchTiming != nil || ie.CodebookParameters != nil || ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Csi_RS_ForTracking != nil || len(ie.Srs_AssocCSI_RS) > 0 || ie.SpatialRelations != nil, ie.DefaultQCL_TwoTCI_r16 != nil || ie.CodebookParametersPerBand_r16 != nil || ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil || ie.MaxNumberSCellBFR_r16 != nil || ie.SimultaneousReceptionDiffTypeD_r16 != nil || ie.Ssb_csirs_SINR_measurement_r16 != nil || ie.NonGroupSINR_reporting_r16 != nil || ie.GroupSINR_reporting_r16 != nil || ie.MultiDCI_multiTRP_Parameters_r16 != nil || ie.SingleDCI_SDM_scheme_Parameters_r16 != nil || ie.SupportFDM_SchemeA_r16 != nil || ie.SupportCodeWordSoftCombining_r16 != nil || ie.SupportTDM_SchemeA_r16 != nil || ie.SupportInter_slotTDM_r16 != nil || ie.LowPAPR_DMRS_PDSCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil || ie.LowPAPR_DMRS_PUCCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil || ie.CodebookParametersAddition_r16 != nil || ie.CodebookComboParametersAddition_r16 != nil || ie.BeamCorrespondenceSSB_based_r16 != nil || ie.BeamCorrespondenceCSI_RS_based_r16 != nil || ie.BeamSwitchTiming_r16 != nil, ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil || ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil, ie.SpatialRelations_v1640 != nil || ie.Support64CandidateBeamRS_BFR_r16 != nil, ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16 != nil, ie.SupportedSINR_meas_v1670 != nil, ie.Srs_increasedRepetition_r17 != nil || ie.Srs_partialFrequencySounding_r17 != nil || ie.Srs_startRB_locationHoppingPartial_r17 != nil || ie.Srs_combEight_r17 != nil || ie.CodebookParametersfetype2_r17 != nil || ie.MTRP_PUSCH_twoCSI_RS_r17 != nil || ie.MTRP_PUCCH_InterSlot_r17 != nil || ie.MTRP_PUCCH_CyclicMapping_r17 != nil || ie.MTRP_PUCCH_SecondTPC_r17 != nil || ie.MTRP_BFR_twoBFD_RS_Set_r17 != nil || ie.MTRP_BFR_PUCCH_SR_perCG_r17 != nil || ie.MTRP_BFR_association_PUCCH_SR_r17 != nil || ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil || ie.Sfn_DefaultDL_BeamSetup_r17 != nil || ie.Sfn_DefaultUL_BeamSetup_r17 != nil || ie.Srs_TriggeringOffset_r17 != nil || ie.Srs_TriggeringDCI_r17 != nil || ie.CodebookComboParameterMixedType_r17 != nil || ie.UnifiedJointTCI_r17 != nil || ie.UnifiedJointTCI_multiMAC_CE_r17 != nil || ie.UnifiedJointTCI_perBWP_CA_r17 != nil || ie.UnifiedJointTCI_ListSharingCA_r17 != nil || ie.UnifiedJointTCI_commonMultiCC_r17 != nil || ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil || ie.UnifiedJointTCI_PC_association_r17 != nil || ie.UnifiedJointTCI_Legacy_r17 != nil || ie.UnifiedJointTCI_Legacy_SRS_r17 != nil || ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil || ie.UnifiedJointTCI_SCellBFR_r17 != nil || ie.UnifiedJointTCI_InterCell_r17 != nil || ie.UnifiedSeparateTCI_r17 != nil || ie.UnifiedSeparateTCI_multiMAC_CE_r17 != nil || ie.UnifiedSeparateTCI_perBWP_CA_r17 != nil || ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil || ie.UnifiedSeparateTCI_commonMultiCC_r17 != nil || ie.UnifiedSeparateTCI_InterCell_r17 != nil || ie.UnifiedJointTCI_mTRP_InterCell_BM_r17 != nil || ie.Mpe_Mitigation_r17 != nil || ie.Srs_PortReport_r17 != nil || ie.MTRP_PDCCH_individual_r17 != nil || ie.MTRP_PDCCH_anySpan_3Symbols_r17 != nil || ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil || ie.MTRP_PUSCH_CSI_RS_r17 != nil || ie.MTRP_PUSCH_cyclicMapping_r17 != nil || ie.MTRP_PUSCH_secondTPC_r17 != nil || ie.MTRP_PUSCH_twoPHR_Reporting_r17 != nil || ie.MTRP_PUSCH_A_CSI_r17 != nil || ie.MTRP_PUSCH_SP_CSI_r17 != nil || ie.MTRP_PUSCH_CG_r17 != nil || ie.MTRP_PUCCH_MAC_CE_r17 != nil || ie.MTRP_PUCCH_maxNum_PC_FR1_r17 != nil || ie.MTRP_inter_Cell_r17 != nil || ie.MTRP_GroupBasedL1_RSRP_r17 != nil || ie.MTRP_BFD_RS_MAC_CE_r17 != nil || ie.MTRP_CSI_EnhancementPerBand_r17 != nil || ie.CodebookComboParameterMultiTRP_r17 != nil || ie.MTRP_CSI_additionalCSI_r17 != nil || ie.MTRP_CSI_N_Max2_r17 != nil || ie.MTRP_CSI_CMR_r17 != nil || ie.Srs_partialFreqSounding_r17 != nil || ie.BeamSwitchTiming_v1710 != nil || ie.BeamSwitchTiming_r17 != nil || ie.BeamReportTiming_v1710 != nil || ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil, ie.Srs_PortReportSP_AP_r17 != nil || ie.MaxNumberRxBeam_v1720 != nil || ie.Sfn_ImplicitRS_twoTCI_r17 != nil || ie.Sfn_QCL_TypeD_Collision_twoTCI_r17 != nil || ie.MTRP_CSI_numCPU_r17 != nil, ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MIMO_ParametersPerBand", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dummy6 != nil, ie.BeamManagementSSB_CSI_RS != nil, ie.BeamSwitchTiming != nil, ie.CodebookParameters != nil, ie.Csi_RS_IM_ReceptionForFeedback != nil, ie.Csi_RS_ProcFrameworkForSRS != nil, ie.Csi_ReportFramework != nil, ie.Csi_RS_ForTracking != nil, len(ie.Srs_AssocCSI_RS) > 0, ie.SpatialRelations != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dummy6 optional
			if ie.Dummy6 != nil {
				if err = ie.Dummy6.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy6", err)
				}
			}
			// encode BeamManagementSSB_CSI_RS optional
			if ie.BeamManagementSSB_CSI_RS != nil {
				if err = ie.BeamManagementSSB_CSI_RS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamManagementSSB_CSI_RS", err)
				}
			}
			// encode BeamSwitchTiming optional
			if ie.BeamSwitchTiming != nil {
				if err = ie.BeamSwitchTiming.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamSwitchTiming", err)
				}
			}
			// encode CodebookParameters optional
			if ie.CodebookParameters != nil {
				if err = ie.CodebookParameters.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookParameters", err)
				}
			}
			// encode Csi_RS_IM_ReceptionForFeedback optional
			if ie.Csi_RS_IM_ReceptionForFeedback != nil {
				if err = ie.Csi_RS_IM_ReceptionForFeedback.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// encode Csi_RS_ProcFrameworkForSRS optional
			if ie.Csi_RS_ProcFrameworkForSRS != nil {
				if err = ie.Csi_RS_ProcFrameworkForSRS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// encode Csi_ReportFramework optional
			if ie.Csi_ReportFramework != nil {
				if err = ie.Csi_ReportFramework.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportFramework", err)
				}
			}
			// encode Csi_RS_ForTracking optional
			if ie.Csi_RS_ForTracking != nil {
				if err = ie.Csi_RS_ForTracking.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_ForTracking", err)
				}
			}
			// encode Srs_AssocCSI_RS optional
			if len(ie.Srs_AssocCSI_RS) > 0 {
				tmp_Srs_AssocCSI_RS := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
				for _, i := range ie.Srs_AssocCSI_RS {
					tmp_Srs_AssocCSI_RS.Value = append(tmp_Srs_AssocCSI_RS.Value, &i)
				}
				if err = tmp_Srs_AssocCSI_RS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_AssocCSI_RS", err)
				}
			}
			// encode SpatialRelations optional
			if ie.SpatialRelations != nil {
				if err = ie.SpatialRelations.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelations", err)
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
			optionals_ext_2 := []bool{ie.DefaultQCL_TwoTCI_r16 != nil, ie.CodebookParametersPerBand_r16 != nil, ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil, ie.MaxNumberSCellBFR_r16 != nil, ie.SimultaneousReceptionDiffTypeD_r16 != nil, ie.Ssb_csirs_SINR_measurement_r16 != nil, ie.NonGroupSINR_reporting_r16 != nil, ie.GroupSINR_reporting_r16 != nil, ie.MultiDCI_multiTRP_Parameters_r16 != nil, ie.SingleDCI_SDM_scheme_Parameters_r16 != nil, ie.SupportFDM_SchemeA_r16 != nil, ie.SupportCodeWordSoftCombining_r16 != nil, ie.SupportTDM_SchemeA_r16 != nil, ie.SupportInter_slotTDM_r16 != nil, ie.LowPAPR_DMRS_PDSCH_r16 != nil, ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil, ie.LowPAPR_DMRS_PUCCH_r16 != nil, ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil, ie.Csi_ReportFrameworkExt_r16 != nil, ie.CodebookParametersAddition_r16 != nil, ie.CodebookComboParametersAddition_r16 != nil, ie.BeamCorrespondenceSSB_based_r16 != nil, ie.BeamCorrespondenceCSI_RS_based_r16 != nil, ie.BeamSwitchTiming_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DefaultQCL_TwoTCI_r16 optional
			if ie.DefaultQCL_TwoTCI_r16 != nil {
				if err = ie.DefaultQCL_TwoTCI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DefaultQCL_TwoTCI_r16", err)
				}
			}
			// encode CodebookParametersPerBand_r16 optional
			if ie.CodebookParametersPerBand_r16 != nil {
				if err = ie.CodebookParametersPerBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookParametersPerBand_r16", err)
				}
			}
			// encode Simul_SpatialRelationUpdatePUCCHResGroup_r16 optional
			if ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil {
				if err = ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Simul_SpatialRelationUpdatePUCCHResGroup_r16", err)
				}
			}
			// encode MaxNumberSCellBFR_r16 optional
			if ie.MaxNumberSCellBFR_r16 != nil {
				if err = ie.MaxNumberSCellBFR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberSCellBFR_r16", err)
				}
			}
			// encode SimultaneousReceptionDiffTypeD_r16 optional
			if ie.SimultaneousReceptionDiffTypeD_r16 != nil {
				if err = ie.SimultaneousReceptionDiffTypeD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousReceptionDiffTypeD_r16", err)
				}
			}
			// encode Ssb_csirs_SINR_measurement_r16 optional
			if ie.Ssb_csirs_SINR_measurement_r16 != nil {
				if err = ie.Ssb_csirs_SINR_measurement_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ssb_csirs_SINR_measurement_r16", err)
				}
			}
			// encode NonGroupSINR_reporting_r16 optional
			if ie.NonGroupSINR_reporting_r16 != nil {
				if err = ie.NonGroupSINR_reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode NonGroupSINR_reporting_r16", err)
				}
			}
			// encode GroupSINR_reporting_r16 optional
			if ie.GroupSINR_reporting_r16 != nil {
				if err = ie.GroupSINR_reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GroupSINR_reporting_r16", err)
				}
			}
			// encode MultiDCI_multiTRP_Parameters_r16 optional
			if ie.MultiDCI_multiTRP_Parameters_r16 != nil {
				if err = ie.MultiDCI_multiTRP_Parameters_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MultiDCI_multiTRP_Parameters_r16", err)
				}
			}
			// encode SingleDCI_SDM_scheme_Parameters_r16 optional
			if ie.SingleDCI_SDM_scheme_Parameters_r16 != nil {
				if err = ie.SingleDCI_SDM_scheme_Parameters_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SingleDCI_SDM_scheme_Parameters_r16", err)
				}
			}
			// encode SupportFDM_SchemeA_r16 optional
			if ie.SupportFDM_SchemeA_r16 != nil {
				if err = ie.SupportFDM_SchemeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportFDM_SchemeA_r16", err)
				}
			}
			// encode SupportCodeWordSoftCombining_r16 optional
			if ie.SupportCodeWordSoftCombining_r16 != nil {
				if err = ie.SupportCodeWordSoftCombining_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportCodeWordSoftCombining_r16", err)
				}
			}
			// encode SupportTDM_SchemeA_r16 optional
			if ie.SupportTDM_SchemeA_r16 != nil {
				if err = ie.SupportTDM_SchemeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportTDM_SchemeA_r16", err)
				}
			}
			// encode SupportInter_slotTDM_r16 optional
			if ie.SupportInter_slotTDM_r16 != nil {
				if err = ie.SupportInter_slotTDM_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportInter_slotTDM_r16", err)
				}
			}
			// encode LowPAPR_DMRS_PDSCH_r16 optional
			if ie.LowPAPR_DMRS_PDSCH_r16 != nil {
				if err = ie.LowPAPR_DMRS_PDSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LowPAPR_DMRS_PDSCH_r16", err)
				}
			}
			// encode LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 optional
			if ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil {
				if err = ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LowPAPR_DMRS_PUSCHwithoutPrecoding_r16", err)
				}
			}
			// encode LowPAPR_DMRS_PUCCH_r16 optional
			if ie.LowPAPR_DMRS_PUCCH_r16 != nil {
				if err = ie.LowPAPR_DMRS_PUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LowPAPR_DMRS_PUCCH_r16", err)
				}
			}
			// encode LowPAPR_DMRS_PUSCHwithPrecoding_r16 optional
			if ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil {
				if err = ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode LowPAPR_DMRS_PUSCHwithPrecoding_r16", err)
				}
			}
			// encode Csi_ReportFrameworkExt_r16 optional
			if ie.Csi_ReportFrameworkExt_r16 != nil {
				if err = ie.Csi_ReportFrameworkExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportFrameworkExt_r16", err)
				}
			}
			// encode CodebookParametersAddition_r16 optional
			if ie.CodebookParametersAddition_r16 != nil {
				if err = ie.CodebookParametersAddition_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookParametersAddition_r16", err)
				}
			}
			// encode CodebookComboParametersAddition_r16 optional
			if ie.CodebookComboParametersAddition_r16 != nil {
				if err = ie.CodebookComboParametersAddition_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookComboParametersAddition_r16", err)
				}
			}
			// encode BeamCorrespondenceSSB_based_r16 optional
			if ie.BeamCorrespondenceSSB_based_r16 != nil {
				if err = ie.BeamCorrespondenceSSB_based_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamCorrespondenceSSB_based_r16", err)
				}
			}
			// encode BeamCorrespondenceCSI_RS_based_r16 optional
			if ie.BeamCorrespondenceCSI_RS_based_r16 != nil {
				if err = ie.BeamCorrespondenceCSI_RS_based_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamCorrespondenceCSI_RS_based_r16", err)
				}
			}
			// encode BeamSwitchTiming_r16 optional
			if ie.BeamSwitchTiming_r16 != nil {
				if err = ie.BeamSwitchTiming_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamSwitchTiming_r16", err)
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
			optionals_ext_3 := []bool{ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil, ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Semi_PersistentL1_SINR_Report_PUCCH_r16 optional
			if ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil {
				if err = ie.Semi_PersistentL1_SINR_Report_PUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Semi_PersistentL1_SINR_Report_PUCCH_r16", err)
				}
			}
			// encode Semi_PersistentL1_SINR_Report_PUSCH_r16 optional
			if ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil {
				if err = ie.Semi_PersistentL1_SINR_Report_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Semi_PersistentL1_SINR_Report_PUSCH_r16", err)
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
			optionals_ext_4 := []bool{ie.SpatialRelations_v1640 != nil, ie.Support64CandidateBeamRS_BFR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SpatialRelations_v1640 optional
			if ie.SpatialRelations_v1640 != nil {
				if err = ie.SpatialRelations_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelations_v1640", err)
				}
			}
			// encode Support64CandidateBeamRS_BFR_r16 optional
			if ie.Support64CandidateBeamRS_BFR_r16 != nil {
				if err = ie.Support64CandidateBeamRS_BFR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Support64CandidateBeamRS_BFR_r16", err)
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
			optionals_ext_5 := []bool{ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxMIMO_LayersForMulti_DCI_mTRP_r16 optional
			if ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16 != nil {
				if err = ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxMIMO_LayersForMulti_DCI_mTRP_r16", err)
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
			optionals_ext_6 := []bool{ie.SupportedSINR_meas_v1670 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportedSINR_meas_v1670 optional
			if ie.SupportedSINR_meas_v1670 != nil {
				if err = extWriter.WriteBitString(ie.SupportedSINR_meas_v1670.Bytes, uint(ie.SupportedSINR_meas_v1670.NumBits), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode SupportedSINR_meas_v1670", err)
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
			optionals_ext_7 := []bool{ie.Srs_increasedRepetition_r17 != nil, ie.Srs_partialFrequencySounding_r17 != nil, ie.Srs_startRB_locationHoppingPartial_r17 != nil, ie.Srs_combEight_r17 != nil, ie.CodebookParametersfetype2_r17 != nil, ie.MTRP_PUSCH_twoCSI_RS_r17 != nil, ie.MTRP_PUCCH_InterSlot_r17 != nil, ie.MTRP_PUCCH_CyclicMapping_r17 != nil, ie.MTRP_PUCCH_SecondTPC_r17 != nil, ie.MTRP_BFR_twoBFD_RS_Set_r17 != nil, ie.MTRP_BFR_PUCCH_SR_perCG_r17 != nil, ie.MTRP_BFR_association_PUCCH_SR_r17 != nil, ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil, ie.Sfn_DefaultDL_BeamSetup_r17 != nil, ie.Sfn_DefaultUL_BeamSetup_r17 != nil, ie.Srs_TriggeringOffset_r17 != nil, ie.Srs_TriggeringDCI_r17 != nil, ie.CodebookComboParameterMixedType_r17 != nil, ie.UnifiedJointTCI_r17 != nil, ie.UnifiedJointTCI_multiMAC_CE_r17 != nil, ie.UnifiedJointTCI_perBWP_CA_r17 != nil, ie.UnifiedJointTCI_ListSharingCA_r17 != nil, ie.UnifiedJointTCI_commonMultiCC_r17 != nil, ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil, ie.UnifiedJointTCI_PC_association_r17 != nil, ie.UnifiedJointTCI_Legacy_r17 != nil, ie.UnifiedJointTCI_Legacy_SRS_r17 != nil, ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil, ie.UnifiedJointTCI_SCellBFR_r17 != nil, ie.UnifiedJointTCI_InterCell_r17 != nil, ie.UnifiedSeparateTCI_r17 != nil, ie.UnifiedSeparateTCI_multiMAC_CE_r17 != nil, ie.UnifiedSeparateTCI_perBWP_CA_r17 != nil, ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil, ie.UnifiedSeparateTCI_commonMultiCC_r17 != nil, ie.UnifiedSeparateTCI_InterCell_r17 != nil, ie.UnifiedJointTCI_mTRP_InterCell_BM_r17 != nil, ie.Mpe_Mitigation_r17 != nil, ie.Srs_PortReport_r17 != nil, ie.MTRP_PDCCH_individual_r17 != nil, ie.MTRP_PDCCH_anySpan_3Symbols_r17 != nil, ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil, ie.MTRP_PUSCH_CSI_RS_r17 != nil, ie.MTRP_PUSCH_cyclicMapping_r17 != nil, ie.MTRP_PUSCH_secondTPC_r17 != nil, ie.MTRP_PUSCH_twoPHR_Reporting_r17 != nil, ie.MTRP_PUSCH_A_CSI_r17 != nil, ie.MTRP_PUSCH_SP_CSI_r17 != nil, ie.MTRP_PUSCH_CG_r17 != nil, ie.MTRP_PUCCH_MAC_CE_r17 != nil, ie.MTRP_PUCCH_maxNum_PC_FR1_r17 != nil, ie.MTRP_inter_Cell_r17 != nil, ie.MTRP_GroupBasedL1_RSRP_r17 != nil, ie.MTRP_BFD_RS_MAC_CE_r17 != nil, ie.MTRP_CSI_EnhancementPerBand_r17 != nil, ie.CodebookComboParameterMultiTRP_r17 != nil, ie.MTRP_CSI_additionalCSI_r17 != nil, ie.MTRP_CSI_N_Max2_r17 != nil, ie.MTRP_CSI_CMR_r17 != nil, ie.Srs_partialFreqSounding_r17 != nil, ie.BeamSwitchTiming_v1710 != nil, ie.BeamSwitchTiming_r17 != nil, ie.BeamReportTiming_v1710 != nil, ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Srs_increasedRepetition_r17 optional
			if ie.Srs_increasedRepetition_r17 != nil {
				if err = ie.Srs_increasedRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_increasedRepetition_r17", err)
				}
			}
			// encode Srs_partialFrequencySounding_r17 optional
			if ie.Srs_partialFrequencySounding_r17 != nil {
				if err = ie.Srs_partialFrequencySounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_partialFrequencySounding_r17", err)
				}
			}
			// encode Srs_startRB_locationHoppingPartial_r17 optional
			if ie.Srs_startRB_locationHoppingPartial_r17 != nil {
				if err = ie.Srs_startRB_locationHoppingPartial_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_startRB_locationHoppingPartial_r17", err)
				}
			}
			// encode Srs_combEight_r17 optional
			if ie.Srs_combEight_r17 != nil {
				if err = ie.Srs_combEight_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_combEight_r17", err)
				}
			}
			// encode CodebookParametersfetype2_r17 optional
			if ie.CodebookParametersfetype2_r17 != nil {
				if err = ie.CodebookParametersfetype2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookParametersfetype2_r17", err)
				}
			}
			// encode MTRP_PUSCH_twoCSI_RS_r17 optional
			if ie.MTRP_PUSCH_twoCSI_RS_r17 != nil {
				if err = ie.MTRP_PUSCH_twoCSI_RS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_twoCSI_RS_r17", err)
				}
			}
			// encode MTRP_PUCCH_InterSlot_r17 optional
			if ie.MTRP_PUCCH_InterSlot_r17 != nil {
				if err = ie.MTRP_PUCCH_InterSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUCCH_InterSlot_r17", err)
				}
			}
			// encode MTRP_PUCCH_CyclicMapping_r17 optional
			if ie.MTRP_PUCCH_CyclicMapping_r17 != nil {
				if err = ie.MTRP_PUCCH_CyclicMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUCCH_CyclicMapping_r17", err)
				}
			}
			// encode MTRP_PUCCH_SecondTPC_r17 optional
			if ie.MTRP_PUCCH_SecondTPC_r17 != nil {
				if err = ie.MTRP_PUCCH_SecondTPC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUCCH_SecondTPC_r17", err)
				}
			}
			// encode MTRP_BFR_twoBFD_RS_Set_r17 optional
			if ie.MTRP_BFR_twoBFD_RS_Set_r17 != nil {
				if err = ie.MTRP_BFR_twoBFD_RS_Set_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_BFR_twoBFD_RS_Set_r17", err)
				}
			}
			// encode MTRP_BFR_PUCCH_SR_perCG_r17 optional
			if ie.MTRP_BFR_PUCCH_SR_perCG_r17 != nil {
				if err = ie.MTRP_BFR_PUCCH_SR_perCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_BFR_PUCCH_SR_perCG_r17", err)
				}
			}
			// encode MTRP_BFR_association_PUCCH_SR_r17 optional
			if ie.MTRP_BFR_association_PUCCH_SR_r17 != nil {
				if err = ie.MTRP_BFR_association_PUCCH_SR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_BFR_association_PUCCH_SR_r17", err)
				}
			}
			// encode Sfn_SimulTwoTCI_AcrossMultiCC_r17 optional
			if ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil {
				if err = ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sfn_SimulTwoTCI_AcrossMultiCC_r17", err)
				}
			}
			// encode Sfn_DefaultDL_BeamSetup_r17 optional
			if ie.Sfn_DefaultDL_BeamSetup_r17 != nil {
				if err = ie.Sfn_DefaultDL_BeamSetup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sfn_DefaultDL_BeamSetup_r17", err)
				}
			}
			// encode Sfn_DefaultUL_BeamSetup_r17 optional
			if ie.Sfn_DefaultUL_BeamSetup_r17 != nil {
				if err = ie.Sfn_DefaultUL_BeamSetup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sfn_DefaultUL_BeamSetup_r17", err)
				}
			}
			// encode Srs_TriggeringOffset_r17 optional
			if ie.Srs_TriggeringOffset_r17 != nil {
				if err = ie.Srs_TriggeringOffset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_TriggeringOffset_r17", err)
				}
			}
			// encode Srs_TriggeringDCI_r17 optional
			if ie.Srs_TriggeringDCI_r17 != nil {
				if err = ie.Srs_TriggeringDCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_TriggeringDCI_r17", err)
				}
			}
			// encode CodebookComboParameterMixedType_r17 optional
			if ie.CodebookComboParameterMixedType_r17 != nil {
				if err = ie.CodebookComboParameterMixedType_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookComboParameterMixedType_r17", err)
				}
			}
			// encode UnifiedJointTCI_r17 optional
			if ie.UnifiedJointTCI_r17 != nil {
				if err = ie.UnifiedJointTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_r17", err)
				}
			}
			// encode UnifiedJointTCI_multiMAC_CE_r17 optional
			if ie.UnifiedJointTCI_multiMAC_CE_r17 != nil {
				if err = ie.UnifiedJointTCI_multiMAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_multiMAC_CE_r17", err)
				}
			}
			// encode UnifiedJointTCI_perBWP_CA_r17 optional
			if ie.UnifiedJointTCI_perBWP_CA_r17 != nil {
				if err = ie.UnifiedJointTCI_perBWP_CA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_perBWP_CA_r17", err)
				}
			}
			// encode UnifiedJointTCI_ListSharingCA_r17 optional
			if ie.UnifiedJointTCI_ListSharingCA_r17 != nil {
				if err = ie.UnifiedJointTCI_ListSharingCA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_ListSharingCA_r17", err)
				}
			}
			// encode UnifiedJointTCI_commonMultiCC_r17 optional
			if ie.UnifiedJointTCI_commonMultiCC_r17 != nil {
				if err = ie.UnifiedJointTCI_commonMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_commonMultiCC_r17", err)
				}
			}
			// encode UnifiedJointTCI_BeamAlignDLRS_r17 optional
			if ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil {
				if err = ie.UnifiedJointTCI_BeamAlignDLRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_BeamAlignDLRS_r17", err)
				}
			}
			// encode UnifiedJointTCI_PC_association_r17 optional
			if ie.UnifiedJointTCI_PC_association_r17 != nil {
				if err = ie.UnifiedJointTCI_PC_association_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_PC_association_r17", err)
				}
			}
			// encode UnifiedJointTCI_Legacy_r17 optional
			if ie.UnifiedJointTCI_Legacy_r17 != nil {
				if err = ie.UnifiedJointTCI_Legacy_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_Legacy_r17", err)
				}
			}
			// encode UnifiedJointTCI_Legacy_SRS_r17 optional
			if ie.UnifiedJointTCI_Legacy_SRS_r17 != nil {
				if err = ie.UnifiedJointTCI_Legacy_SRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_Legacy_SRS_r17", err)
				}
			}
			// encode UnifiedJointTCI_Legacy_CORESET0_r17 optional
			if ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil {
				if err = ie.UnifiedJointTCI_Legacy_CORESET0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_Legacy_CORESET0_r17", err)
				}
			}
			// encode UnifiedJointTCI_SCellBFR_r17 optional
			if ie.UnifiedJointTCI_SCellBFR_r17 != nil {
				if err = ie.UnifiedJointTCI_SCellBFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_SCellBFR_r17", err)
				}
			}
			// encode UnifiedJointTCI_InterCell_r17 optional
			if ie.UnifiedJointTCI_InterCell_r17 != nil {
				if err = ie.UnifiedJointTCI_InterCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_InterCell_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_r17 optional
			if ie.UnifiedSeparateTCI_r17 != nil {
				if err = ie.UnifiedSeparateTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_multiMAC_CE_r17 optional
			if ie.UnifiedSeparateTCI_multiMAC_CE_r17 != nil {
				if err = ie.UnifiedSeparateTCI_multiMAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_multiMAC_CE_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_perBWP_CA_r17 optional
			if ie.UnifiedSeparateTCI_perBWP_CA_r17 != nil {
				if err = ie.UnifiedSeparateTCI_perBWP_CA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_perBWP_CA_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_ListSharingCA_r17 optional
			if ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil {
				if err = ie.UnifiedSeparateTCI_ListSharingCA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_ListSharingCA_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_commonMultiCC_r17 optional
			if ie.UnifiedSeparateTCI_commonMultiCC_r17 != nil {
				if err = ie.UnifiedSeparateTCI_commonMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_commonMultiCC_r17", err)
				}
			}
			// encode UnifiedSeparateTCI_InterCell_r17 optional
			if ie.UnifiedSeparateTCI_InterCell_r17 != nil {
				if err = ie.UnifiedSeparateTCI_InterCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedSeparateTCI_InterCell_r17", err)
				}
			}
			// encode UnifiedJointTCI_mTRP_InterCell_BM_r17 optional
			if ie.UnifiedJointTCI_mTRP_InterCell_BM_r17 != nil {
				if err = ie.UnifiedJointTCI_mTRP_InterCell_BM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UnifiedJointTCI_mTRP_InterCell_BM_r17", err)
				}
			}
			// encode Mpe_Mitigation_r17 optional
			if ie.Mpe_Mitigation_r17 != nil {
				if err = ie.Mpe_Mitigation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mpe_Mitigation_r17", err)
				}
			}
			// encode Srs_PortReport_r17 optional
			if ie.Srs_PortReport_r17 != nil {
				if err = ie.Srs_PortReport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PortReport_r17", err)
				}
			}
			// encode MTRP_PDCCH_individual_r17 optional
			if ie.MTRP_PDCCH_individual_r17 != nil {
				if err = ie.MTRP_PDCCH_individual_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PDCCH_individual_r17", err)
				}
			}
			// encode MTRP_PDCCH_anySpan_3Symbols_r17 optional
			if ie.MTRP_PDCCH_anySpan_3Symbols_r17 != nil {
				if err = ie.MTRP_PDCCH_anySpan_3Symbols_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PDCCH_anySpan_3Symbols_r17", err)
				}
			}
			// encode MTRP_PDCCH_TwoQCL_TypeD_r17 optional
			if ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil {
				if err = ie.MTRP_PDCCH_TwoQCL_TypeD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PDCCH_TwoQCL_TypeD_r17", err)
				}
			}
			// encode MTRP_PUSCH_CSI_RS_r17 optional
			if ie.MTRP_PUSCH_CSI_RS_r17 != nil {
				if err = ie.MTRP_PUSCH_CSI_RS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_CSI_RS_r17", err)
				}
			}
			// encode MTRP_PUSCH_cyclicMapping_r17 optional
			if ie.MTRP_PUSCH_cyclicMapping_r17 != nil {
				if err = ie.MTRP_PUSCH_cyclicMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_cyclicMapping_r17", err)
				}
			}
			// encode MTRP_PUSCH_secondTPC_r17 optional
			if ie.MTRP_PUSCH_secondTPC_r17 != nil {
				if err = ie.MTRP_PUSCH_secondTPC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_secondTPC_r17", err)
				}
			}
			// encode MTRP_PUSCH_twoPHR_Reporting_r17 optional
			if ie.MTRP_PUSCH_twoPHR_Reporting_r17 != nil {
				if err = ie.MTRP_PUSCH_twoPHR_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_twoPHR_Reporting_r17", err)
				}
			}
			// encode MTRP_PUSCH_A_CSI_r17 optional
			if ie.MTRP_PUSCH_A_CSI_r17 != nil {
				if err = ie.MTRP_PUSCH_A_CSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_A_CSI_r17", err)
				}
			}
			// encode MTRP_PUSCH_SP_CSI_r17 optional
			if ie.MTRP_PUSCH_SP_CSI_r17 != nil {
				if err = ie.MTRP_PUSCH_SP_CSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_SP_CSI_r17", err)
				}
			}
			// encode MTRP_PUSCH_CG_r17 optional
			if ie.MTRP_PUSCH_CG_r17 != nil {
				if err = ie.MTRP_PUSCH_CG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUSCH_CG_r17", err)
				}
			}
			// encode MTRP_PUCCH_MAC_CE_r17 optional
			if ie.MTRP_PUCCH_MAC_CE_r17 != nil {
				if err = ie.MTRP_PUCCH_MAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_PUCCH_MAC_CE_r17", err)
				}
			}
			// encode MTRP_PUCCH_maxNum_PC_FR1_r17 optional
			if ie.MTRP_PUCCH_maxNum_PC_FR1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.MTRP_PUCCH_maxNum_PC_FR1_r17, &aper.Constraint{Lb: 3, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode MTRP_PUCCH_maxNum_PC_FR1_r17", err)
				}
			}
			// encode MTRP_inter_Cell_r17 optional
			if ie.MTRP_inter_Cell_r17 != nil {
				if err = ie.MTRP_inter_Cell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_inter_Cell_r17", err)
				}
			}
			// encode MTRP_GroupBasedL1_RSRP_r17 optional
			if ie.MTRP_GroupBasedL1_RSRP_r17 != nil {
				if err = ie.MTRP_GroupBasedL1_RSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_GroupBasedL1_RSRP_r17", err)
				}
			}
			// encode MTRP_BFD_RS_MAC_CE_r17 optional
			if ie.MTRP_BFD_RS_MAC_CE_r17 != nil {
				if err = ie.MTRP_BFD_RS_MAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_BFD_RS_MAC_CE_r17", err)
				}
			}
			// encode MTRP_CSI_EnhancementPerBand_r17 optional
			if ie.MTRP_CSI_EnhancementPerBand_r17 != nil {
				if err = ie.MTRP_CSI_EnhancementPerBand_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_CSI_EnhancementPerBand_r17", err)
				}
			}
			// encode CodebookComboParameterMultiTRP_r17 optional
			if ie.CodebookComboParameterMultiTRP_r17 != nil {
				if err = ie.CodebookComboParameterMultiTRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookComboParameterMultiTRP_r17", err)
				}
			}
			// encode MTRP_CSI_additionalCSI_r17 optional
			if ie.MTRP_CSI_additionalCSI_r17 != nil {
				if err = ie.MTRP_CSI_additionalCSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_CSI_additionalCSI_r17", err)
				}
			}
			// encode MTRP_CSI_N_Max2_r17 optional
			if ie.MTRP_CSI_N_Max2_r17 != nil {
				if err = ie.MTRP_CSI_N_Max2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_CSI_N_Max2_r17", err)
				}
			}
			// encode MTRP_CSI_CMR_r17 optional
			if ie.MTRP_CSI_CMR_r17 != nil {
				if err = ie.MTRP_CSI_CMR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_CSI_CMR_r17", err)
				}
			}
			// encode Srs_partialFreqSounding_r17 optional
			if ie.Srs_partialFreqSounding_r17 != nil {
				if err = ie.Srs_partialFreqSounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_partialFreqSounding_r17", err)
				}
			}
			// encode BeamSwitchTiming_v1710 optional
			if ie.BeamSwitchTiming_v1710 != nil {
				if err = ie.BeamSwitchTiming_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamSwitchTiming_v1710", err)
				}
			}
			// encode BeamSwitchTiming_r17 optional
			if ie.BeamSwitchTiming_r17 != nil {
				if err = ie.BeamSwitchTiming_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamSwitchTiming_r17", err)
				}
			}
			// encode BeamReportTiming_v1710 optional
			if ie.BeamReportTiming_v1710 != nil {
				if err = ie.BeamReportTiming_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode BeamReportTiming_v1710", err)
				}
			}
			// encode MaxNumberRxTxBeamSwitchDL_v1710 optional
			if ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil {
				if err = ie.MaxNumberRxTxBeamSwitchDL_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberRxTxBeamSwitchDL_v1710", err)
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
			optionals_ext_8 := []bool{ie.Srs_PortReportSP_AP_r17 != nil, ie.MaxNumberRxBeam_v1720 != nil, ie.Sfn_ImplicitRS_twoTCI_r17 != nil, ie.Sfn_QCL_TypeD_Collision_twoTCI_r17 != nil, ie.MTRP_CSI_numCPU_r17 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Srs_PortReportSP_AP_r17 optional
			if ie.Srs_PortReportSP_AP_r17 != nil {
				if err = ie.Srs_PortReportSP_AP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Srs_PortReportSP_AP_r17", err)
				}
			}
			// encode MaxNumberRxBeam_v1720 optional
			if ie.MaxNumberRxBeam_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.MaxNumberRxBeam_v1720, &aper.Constraint{Lb: 9, Ub: 12}, false); err != nil {
					return utils.WrapError("Encode MaxNumberRxBeam_v1720", err)
				}
			}
			// encode Sfn_ImplicitRS_twoTCI_r17 optional
			if ie.Sfn_ImplicitRS_twoTCI_r17 != nil {
				if err = ie.Sfn_ImplicitRS_twoTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sfn_ImplicitRS_twoTCI_r17", err)
				}
			}
			// encode Sfn_QCL_TypeD_Collision_twoTCI_r17 optional
			if ie.Sfn_QCL_TypeD_Collision_twoTCI_r17 != nil {
				if err = ie.Sfn_QCL_TypeD_Collision_twoTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Sfn_QCL_TypeD_Collision_twoTCI_r17", err)
				}
			}
			// encode MTRP_CSI_numCPU_r17 optional
			if ie.MTRP_CSI_numCPU_r17 != nil {
				if err = ie.MTRP_CSI_numCPU_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MTRP_CSI_numCPU_r17", err)
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
			optionals_ext_9 := []bool{ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SupportRepNumPDSCH_TDRA_DCI_1_2_r17 optional
			if ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil {
				if err = ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportRepNumPDSCH_TDRA_DCI_1_2_r17", err)
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

func (ie *MIMO_ParametersPerBand) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Tci_StatePDSCHPresent bool
	if Tci_StatePDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalActiveTCI_StatePDCCHPresent bool
	if AdditionalActiveTCI_StatePDCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_TransCoherencePresent bool
	if Pusch_TransCoherencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamCorrespondenceWithoutUL_BeamSweepingPresent bool
	if BeamCorrespondenceWithoutUL_BeamSweepingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PeriodicBeamReportPresent bool
	if PeriodicBeamReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AperiodicBeamReportPresent bool
	if AperiodicBeamReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_BeamReportPUCCHPresent bool
	if Sp_BeamReportPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_BeamReportPUSCHPresent bool
	if Sp_BeamReportPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberRxBeamPresent bool
	if MaxNumberRxBeamPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberRxTxBeamSwitchDLPresent bool
	if MaxNumberRxTxBeamSwitchDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberNonGroupBeamReportingPresent bool
	if MaxNumberNonGroupBeamReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var GroupBeamReportingPresent bool
	if GroupBeamReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkBeamManagementPresent bool
	if UplinkBeamManagementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberCSI_RS_BFDPresent bool
	if MaxNumberCSI_RS_BFDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberSSB_BFDPresent bool
	if MaxNumberSSB_BFDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberCSI_RS_SSB_CBDPresent bool
	if MaxNumberCSI_RS_SSB_CBDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPortsPTRS_ULPresent bool
	if TwoPortsPTRS_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy5Present bool
	if Dummy5Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy3Present bool
	if Dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BeamReportTimingPresent bool
	if BeamReportTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ptrs_DensityRecommendationSetDLPresent bool
	if Ptrs_DensityRecommendationSetDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ptrs_DensityRecommendationSetULPresent bool
	if Ptrs_DensityRecommendationSetULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy4Present bool
	if Dummy4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AperiodicTRSPresent bool
	if AperiodicTRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Tci_StatePDSCHPresent {
		ie.Tci_StatePDSCH = new(MIMO_ParametersPerBand_tci_StatePDSCH)
		if err = ie.Tci_StatePDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Tci_StatePDSCH", err)
		}
	}
	if AdditionalActiveTCI_StatePDCCHPresent {
		ie.AdditionalActiveTCI_StatePDCCH = new(MIMO_ParametersPerBand_additionalActiveTCI_StatePDCCH)
		if err = ie.AdditionalActiveTCI_StatePDCCH.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalActiveTCI_StatePDCCH", err)
		}
	}
	if Pusch_TransCoherencePresent {
		ie.Pusch_TransCoherence = new(MIMO_ParametersPerBand_pusch_TransCoherence)
		if err = ie.Pusch_TransCoherence.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_TransCoherence", err)
		}
	}
	if BeamCorrespondenceWithoutUL_BeamSweepingPresent {
		ie.BeamCorrespondenceWithoutUL_BeamSweeping = new(MIMO_ParametersPerBand_beamCorrespondenceWithoutUL_BeamSweeping)
		if err = ie.BeamCorrespondenceWithoutUL_BeamSweeping.Decode(r); err != nil {
			return utils.WrapError("Decode BeamCorrespondenceWithoutUL_BeamSweeping", err)
		}
	}
	if PeriodicBeamReportPresent {
		ie.PeriodicBeamReport = new(MIMO_ParametersPerBand_periodicBeamReport)
		if err = ie.PeriodicBeamReport.Decode(r); err != nil {
			return utils.WrapError("Decode PeriodicBeamReport", err)
		}
	}
	if AperiodicBeamReportPresent {
		ie.AperiodicBeamReport = new(MIMO_ParametersPerBand_aperiodicBeamReport)
		if err = ie.AperiodicBeamReport.Decode(r); err != nil {
			return utils.WrapError("Decode AperiodicBeamReport", err)
		}
	}
	if Sp_BeamReportPUCCHPresent {
		ie.Sp_BeamReportPUCCH = new(MIMO_ParametersPerBand_sp_BeamReportPUCCH)
		if err = ie.Sp_BeamReportPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_BeamReportPUCCH", err)
		}
	}
	if Sp_BeamReportPUSCHPresent {
		ie.Sp_BeamReportPUSCH = new(MIMO_ParametersPerBand_sp_BeamReportPUSCH)
		if err = ie.Sp_BeamReportPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_BeamReportPUSCH", err)
		}
	}
	if Dummy1Present {
		ie.Dummy1 = new(DummyG)
		if err = ie.Dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
	}
	if MaxNumberRxBeamPresent {
		var tmp_int_MaxNumberRxBeam int64
		if tmp_int_MaxNumberRxBeam, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode MaxNumberRxBeam", err)
		}
		ie.MaxNumberRxBeam = &tmp_int_MaxNumberRxBeam
	}
	if MaxNumberRxTxBeamSwitchDLPresent {
		ie.MaxNumberRxTxBeamSwitchDL = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL)
		if err = ie.MaxNumberRxTxBeamSwitchDL.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberRxTxBeamSwitchDL", err)
		}
	}
	if MaxNumberNonGroupBeamReportingPresent {
		ie.MaxNumberNonGroupBeamReporting = new(MIMO_ParametersPerBand_maxNumberNonGroupBeamReporting)
		if err = ie.MaxNumberNonGroupBeamReporting.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberNonGroupBeamReporting", err)
		}
	}
	if GroupBeamReportingPresent {
		ie.GroupBeamReporting = new(MIMO_ParametersPerBand_groupBeamReporting)
		if err = ie.GroupBeamReporting.Decode(r); err != nil {
			return utils.WrapError("Decode GroupBeamReporting", err)
		}
	}
	if UplinkBeamManagementPresent {
		ie.UplinkBeamManagement = new(MIMO_ParametersPerBand_uplinkBeamManagement)
		if err = ie.UplinkBeamManagement.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkBeamManagement", err)
		}
	}
	if MaxNumberCSI_RS_BFDPresent {
		var tmp_int_MaxNumberCSI_RS_BFD int64
		if tmp_int_MaxNumberCSI_RS_BFD, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode MaxNumberCSI_RS_BFD", err)
		}
		ie.MaxNumberCSI_RS_BFD = &tmp_int_MaxNumberCSI_RS_BFD
	}
	if MaxNumberSSB_BFDPresent {
		var tmp_int_MaxNumberSSB_BFD int64
		if tmp_int_MaxNumberSSB_BFD, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode MaxNumberSSB_BFD", err)
		}
		ie.MaxNumberSSB_BFD = &tmp_int_MaxNumberSSB_BFD
	}
	if MaxNumberCSI_RS_SSB_CBDPresent {
		var tmp_int_MaxNumberCSI_RS_SSB_CBD int64
		if tmp_int_MaxNumberCSI_RS_SSB_CBD, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode MaxNumberCSI_RS_SSB_CBD", err)
		}
		ie.MaxNumberCSI_RS_SSB_CBD = &tmp_int_MaxNumberCSI_RS_SSB_CBD
	}
	if Dummy2Present {
		ie.Dummy2 = new(MIMO_ParametersPerBand_dummy2)
		if err = ie.Dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
	}
	if TwoPortsPTRS_ULPresent {
		ie.TwoPortsPTRS_UL = new(MIMO_ParametersPerBand_twoPortsPTRS_UL)
		if err = ie.TwoPortsPTRS_UL.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPortsPTRS_UL", err)
		}
	}
	if Dummy5Present {
		ie.Dummy5 = new(SRS_Resources)
		if err = ie.Dummy5.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy5", err)
		}
	}
	if Dummy3Present {
		var tmp_int_Dummy3 int64
		if tmp_int_Dummy3, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Dummy3", err)
		}
		ie.Dummy3 = &tmp_int_Dummy3
	}
	if BeamReportTimingPresent {
		ie.BeamReportTiming = new(MIMO_ParametersPerBand_beamReportTiming)
		if err = ie.BeamReportTiming.Decode(r); err != nil {
			return utils.WrapError("Decode BeamReportTiming", err)
		}
	}
	if Ptrs_DensityRecommendationSetDLPresent {
		ie.Ptrs_DensityRecommendationSetDL = new(MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL)
		if err = ie.Ptrs_DensityRecommendationSetDL.Decode(r); err != nil {
			return utils.WrapError("Decode Ptrs_DensityRecommendationSetDL", err)
		}
	}
	if Ptrs_DensityRecommendationSetULPresent {
		ie.Ptrs_DensityRecommendationSetUL = new(MIMO_ParametersPerBand_ptrs_DensityRecommendationSetUL)
		if err = ie.Ptrs_DensityRecommendationSetUL.Decode(r); err != nil {
			return utils.WrapError("Decode Ptrs_DensityRecommendationSetUL", err)
		}
	}
	if Dummy4Present {
		ie.Dummy4 = new(DummyH)
		if err = ie.Dummy4.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy4", err)
		}
	}
	if AperiodicTRSPresent {
		ie.AperiodicTRS = new(MIMO_ParametersPerBand_aperiodicTRS)
		if err = ie.AperiodicTRS.Decode(r); err != nil {
			return utils.WrapError("Decode AperiodicTRS", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 9 bits for 9 extension groups
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

			Dummy6Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamManagementSSB_CSI_RSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamSwitchTimingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookParametersPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_RS_IM_ReceptionForFeedbackPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_RS_ProcFrameworkForSRSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_ReportFrameworkPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_RS_ForTrackingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_AssocCSI_RSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationsPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dummy6 optional
			if Dummy6Present {
				ie.Dummy6 = new(MIMO_ParametersPerBand_dummy6)
				if err = ie.Dummy6.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy6", err)
				}
			}
			// decode BeamManagementSSB_CSI_RS optional
			if BeamManagementSSB_CSI_RSPresent {
				ie.BeamManagementSSB_CSI_RS = new(BeamManagementSSB_CSI_RS)
				if err = ie.BeamManagementSSB_CSI_RS.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamManagementSSB_CSI_RS", err)
				}
			}
			// decode BeamSwitchTiming optional
			if BeamSwitchTimingPresent {
				ie.BeamSwitchTiming = new(MIMO_ParametersPerBand_beamSwitchTiming)
				if err = ie.BeamSwitchTiming.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamSwitchTiming", err)
				}
			}
			// decode CodebookParameters optional
			if CodebookParametersPresent {
				ie.CodebookParameters = new(CodebookParameters)
				if err = ie.CodebookParameters.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookParameters", err)
				}
			}
			// decode Csi_RS_IM_ReceptionForFeedback optional
			if Csi_RS_IM_ReceptionForFeedbackPresent {
				ie.Csi_RS_IM_ReceptionForFeedback = new(CSI_RS_IM_ReceptionForFeedback)
				if err = ie.Csi_RS_IM_ReceptionForFeedback.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// decode Csi_RS_ProcFrameworkForSRS optional
			if Csi_RS_ProcFrameworkForSRSPresent {
				ie.Csi_RS_ProcFrameworkForSRS = new(CSI_RS_ProcFrameworkForSRS)
				if err = ie.Csi_RS_ProcFrameworkForSRS.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// decode Csi_ReportFramework optional
			if Csi_ReportFrameworkPresent {
				ie.Csi_ReportFramework = new(CSI_ReportFramework)
				if err = ie.Csi_ReportFramework.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportFramework", err)
				}
			}
			// decode Csi_RS_ForTracking optional
			if Csi_RS_ForTrackingPresent {
				ie.Csi_RS_ForTracking = new(CSI_RS_ForTracking)
				if err = ie.Csi_RS_ForTracking.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_ForTracking", err)
				}
			}
			// decode Srs_AssocCSI_RS optional
			if Srs_AssocCSI_RSPresent {
				tmp_Srs_AssocCSI_RS := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
				fn_Srs_AssocCSI_RS := func() *SupportedCSI_RS_Resource {
					return new(SupportedCSI_RS_Resource)
				}
				if err = tmp_Srs_AssocCSI_RS.Decode(extReader, fn_Srs_AssocCSI_RS); err != nil {
					return utils.WrapError("Decode Srs_AssocCSI_RS", err)
				}
				ie.Srs_AssocCSI_RS = []SupportedCSI_RS_Resource{}
				for _, i := range tmp_Srs_AssocCSI_RS.Value {
					ie.Srs_AssocCSI_RS = append(ie.Srs_AssocCSI_RS, *i)
				}
			}
			// decode SpatialRelations optional
			if SpatialRelationsPresent {
				ie.SpatialRelations = new(SpatialRelations)
				if err = ie.SpatialRelations.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelations", err)
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

			DefaultQCL_TwoTCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookParametersPerBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Simul_SpatialRelationUpdatePUCCHResGroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberSCellBFR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousReceptionDiffTypeD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ssb_csirs_SINR_measurement_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			NonGroupSINR_reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GroupSINR_reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MultiDCI_multiTRP_Parameters_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SingleDCI_SDM_scheme_Parameters_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportFDM_SchemeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportCodeWordSoftCombining_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportTDM_SchemeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportInter_slotTDM_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LowPAPR_DMRS_PDSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LowPAPR_DMRS_PUSCHwithoutPrecoding_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LowPAPR_DMRS_PUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			LowPAPR_DMRS_PUSCHwithPrecoding_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_ReportFrameworkExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookParametersAddition_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookComboParametersAddition_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamCorrespondenceSSB_based_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamCorrespondenceCSI_RS_based_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamSwitchTiming_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DefaultQCL_TwoTCI_r16 optional
			if DefaultQCL_TwoTCI_r16Present {
				ie.DefaultQCL_TwoTCI_r16 = new(MIMO_ParametersPerBand_defaultQCL_TwoTCI_r16)
				if err = ie.DefaultQCL_TwoTCI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DefaultQCL_TwoTCI_r16", err)
				}
			}
			// decode CodebookParametersPerBand_r16 optional
			if CodebookParametersPerBand_r16Present {
				ie.CodebookParametersPerBand_r16 = new(CodebookParameters_v1610)
				if err = ie.CodebookParametersPerBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookParametersPerBand_r16", err)
				}
			}
			// decode Simul_SpatialRelationUpdatePUCCHResGroup_r16 optional
			if Simul_SpatialRelationUpdatePUCCHResGroup_r16Present {
				ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 = new(MIMO_ParametersPerBand_simul_SpatialRelationUpdatePUCCHResGroup_r16)
				if err = ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Simul_SpatialRelationUpdatePUCCHResGroup_r16", err)
				}
			}
			// decode MaxNumberSCellBFR_r16 optional
			if MaxNumberSCellBFR_r16Present {
				ie.MaxNumberSCellBFR_r16 = new(MIMO_ParametersPerBand_maxNumberSCellBFR_r16)
				if err = ie.MaxNumberSCellBFR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberSCellBFR_r16", err)
				}
			}
			// decode SimultaneousReceptionDiffTypeD_r16 optional
			if SimultaneousReceptionDiffTypeD_r16Present {
				ie.SimultaneousReceptionDiffTypeD_r16 = new(MIMO_ParametersPerBand_simultaneousReceptionDiffTypeD_r16)
				if err = ie.SimultaneousReceptionDiffTypeD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousReceptionDiffTypeD_r16", err)
				}
			}
			// decode Ssb_csirs_SINR_measurement_r16 optional
			if Ssb_csirs_SINR_measurement_r16Present {
				ie.Ssb_csirs_SINR_measurement_r16 = new(MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16)
				if err = ie.Ssb_csirs_SINR_measurement_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ssb_csirs_SINR_measurement_r16", err)
				}
			}
			// decode NonGroupSINR_reporting_r16 optional
			if NonGroupSINR_reporting_r16Present {
				ie.NonGroupSINR_reporting_r16 = new(MIMO_ParametersPerBand_nonGroupSINR_reporting_r16)
				if err = ie.NonGroupSINR_reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode NonGroupSINR_reporting_r16", err)
				}
			}
			// decode GroupSINR_reporting_r16 optional
			if GroupSINR_reporting_r16Present {
				ie.GroupSINR_reporting_r16 = new(MIMO_ParametersPerBand_groupSINR_reporting_r16)
				if err = ie.GroupSINR_reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode GroupSINR_reporting_r16", err)
				}
			}
			// decode MultiDCI_multiTRP_Parameters_r16 optional
			if MultiDCI_multiTRP_Parameters_r16Present {
				ie.MultiDCI_multiTRP_Parameters_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16)
				if err = ie.MultiDCI_multiTRP_Parameters_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MultiDCI_multiTRP_Parameters_r16", err)
				}
			}
			// decode SingleDCI_SDM_scheme_Parameters_r16 optional
			if SingleDCI_SDM_scheme_Parameters_r16Present {
				ie.SingleDCI_SDM_scheme_Parameters_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16)
				if err = ie.SingleDCI_SDM_scheme_Parameters_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SingleDCI_SDM_scheme_Parameters_r16", err)
				}
			}
			// decode SupportFDM_SchemeA_r16 optional
			if SupportFDM_SchemeA_r16Present {
				ie.SupportFDM_SchemeA_r16 = new(MIMO_ParametersPerBand_supportFDM_SchemeA_r16)
				if err = ie.SupportFDM_SchemeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportFDM_SchemeA_r16", err)
				}
			}
			// decode SupportCodeWordSoftCombining_r16 optional
			if SupportCodeWordSoftCombining_r16Present {
				ie.SupportCodeWordSoftCombining_r16 = new(MIMO_ParametersPerBand_supportCodeWordSoftCombining_r16)
				if err = ie.SupportCodeWordSoftCombining_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportCodeWordSoftCombining_r16", err)
				}
			}
			// decode SupportTDM_SchemeA_r16 optional
			if SupportTDM_SchemeA_r16Present {
				ie.SupportTDM_SchemeA_r16 = new(MIMO_ParametersPerBand_supportTDM_SchemeA_r16)
				if err = ie.SupportTDM_SchemeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportTDM_SchemeA_r16", err)
				}
			}
			// decode SupportInter_slotTDM_r16 optional
			if SupportInter_slotTDM_r16Present {
				ie.SupportInter_slotTDM_r16 = new(MIMO_ParametersPerBand_supportInter_slotTDM_r16)
				if err = ie.SupportInter_slotTDM_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportInter_slotTDM_r16", err)
				}
			}
			// decode LowPAPR_DMRS_PDSCH_r16 optional
			if LowPAPR_DMRS_PDSCH_r16Present {
				ie.LowPAPR_DMRS_PDSCH_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PDSCH_r16)
				if err = ie.LowPAPR_DMRS_PDSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LowPAPR_DMRS_PDSCH_r16", err)
				}
			}
			// decode LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 optional
			if LowPAPR_DMRS_PUSCHwithoutPrecoding_r16Present {
				ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithoutPrecoding_r16)
				if err = ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LowPAPR_DMRS_PUSCHwithoutPrecoding_r16", err)
				}
			}
			// decode LowPAPR_DMRS_PUCCH_r16 optional
			if LowPAPR_DMRS_PUCCH_r16Present {
				ie.LowPAPR_DMRS_PUCCH_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUCCH_r16)
				if err = ie.LowPAPR_DMRS_PUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LowPAPR_DMRS_PUCCH_r16", err)
				}
			}
			// decode LowPAPR_DMRS_PUSCHwithPrecoding_r16 optional
			if LowPAPR_DMRS_PUSCHwithPrecoding_r16Present {
				ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithPrecoding_r16)
				if err = ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode LowPAPR_DMRS_PUSCHwithPrecoding_r16", err)
				}
			}
			// decode Csi_ReportFrameworkExt_r16 optional
			if Csi_ReportFrameworkExt_r16Present {
				ie.Csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
				if err = ie.Csi_ReportFrameworkExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportFrameworkExt_r16", err)
				}
			}
			// decode CodebookParametersAddition_r16 optional
			if CodebookParametersAddition_r16Present {
				ie.CodebookParametersAddition_r16 = new(CodebookParametersAddition_r16)
				if err = ie.CodebookParametersAddition_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookParametersAddition_r16", err)
				}
			}
			// decode CodebookComboParametersAddition_r16 optional
			if CodebookComboParametersAddition_r16Present {
				ie.CodebookComboParametersAddition_r16 = new(CodebookComboParametersAddition_r16)
				if err = ie.CodebookComboParametersAddition_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookComboParametersAddition_r16", err)
				}
			}
			// decode BeamCorrespondenceSSB_based_r16 optional
			if BeamCorrespondenceSSB_based_r16Present {
				ie.BeamCorrespondenceSSB_based_r16 = new(MIMO_ParametersPerBand_beamCorrespondenceSSB_based_r16)
				if err = ie.BeamCorrespondenceSSB_based_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamCorrespondenceSSB_based_r16", err)
				}
			}
			// decode BeamCorrespondenceCSI_RS_based_r16 optional
			if BeamCorrespondenceCSI_RS_based_r16Present {
				ie.BeamCorrespondenceCSI_RS_based_r16 = new(MIMO_ParametersPerBand_beamCorrespondenceCSI_RS_based_r16)
				if err = ie.BeamCorrespondenceCSI_RS_based_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamCorrespondenceCSI_RS_based_r16", err)
				}
			}
			// decode BeamSwitchTiming_r16 optional
			if BeamSwitchTiming_r16Present {
				ie.BeamSwitchTiming_r16 = new(MIMO_ParametersPerBand_beamSwitchTiming_r16)
				if err = ie.BeamSwitchTiming_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamSwitchTiming_r16", err)
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

			Semi_PersistentL1_SINR_Report_PUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Semi_PersistentL1_SINR_Report_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Semi_PersistentL1_SINR_Report_PUCCH_r16 optional
			if Semi_PersistentL1_SINR_Report_PUCCH_r16Present {
				ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16)
				if err = ie.Semi_PersistentL1_SINR_Report_PUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Semi_PersistentL1_SINR_Report_PUCCH_r16", err)
				}
			}
			// decode Semi_PersistentL1_SINR_Report_PUSCH_r16 optional
			if Semi_PersistentL1_SINR_Report_PUSCH_r16Present {
				ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUSCH_r16)
				if err = ie.Semi_PersistentL1_SINR_Report_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Semi_PersistentL1_SINR_Report_PUSCH_r16", err)
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

			SpatialRelations_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Support64CandidateBeamRS_BFR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SpatialRelations_v1640 optional
			if SpatialRelations_v1640Present {
				ie.SpatialRelations_v1640 = new(MIMO_ParametersPerBand_spatialRelations_v1640)
				if err = ie.SpatialRelations_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelations_v1640", err)
				}
			}
			// decode Support64CandidateBeamRS_BFR_r16 optional
			if Support64CandidateBeamRS_BFR_r16Present {
				ie.Support64CandidateBeamRS_BFR_r16 = new(MIMO_ParametersPerBand_support64CandidateBeamRS_BFR_r16)
				if err = ie.Support64CandidateBeamRS_BFR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Support64CandidateBeamRS_BFR_r16", err)
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

			MaxMIMO_LayersForMulti_DCI_mTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxMIMO_LayersForMulti_DCI_mTRP_r16 optional
			if MaxMIMO_LayersForMulti_DCI_mTRP_r16Present {
				ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16 = new(MIMO_ParametersPerBand_maxMIMO_LayersForMulti_DCI_mTRP_r16)
				if err = ie.MaxMIMO_LayersForMulti_DCI_mTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxMIMO_LayersForMulti_DCI_mTRP_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SupportedSINR_meas_v1670Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportedSINR_meas_v1670 optional
			if SupportedSINR_meas_v1670Present {
				var tmp_bs_SupportedSINR_meas_v1670 []byte
				var n_SupportedSINR_meas_v1670 uint
				if tmp_bs_SupportedSINR_meas_v1670, n_SupportedSINR_meas_v1670, err = extReader.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode SupportedSINR_meas_v1670", err)
				}
				tmp_bitstring := aper.BitString{
					Bytes:   tmp_bs_SupportedSINR_meas_v1670,
					NumBits: uint64(n_SupportedSINR_meas_v1670),
				}
				ie.SupportedSINR_meas_v1670 = &tmp_bitstring
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Srs_increasedRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_partialFrequencySounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_startRB_locationHoppingPartial_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_combEight_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookParametersfetype2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_twoCSI_RS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUCCH_InterSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUCCH_CyclicMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUCCH_SecondTPC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_BFR_twoBFD_RS_Set_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_BFR_PUCCH_SR_perCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_BFR_association_PUCCH_SR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sfn_SimulTwoTCI_AcrossMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sfn_DefaultDL_BeamSetup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sfn_DefaultUL_BeamSetup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_TriggeringOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_TriggeringDCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookComboParameterMixedType_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_multiMAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_perBWP_CA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_ListSharingCA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_commonMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_BeamAlignDLRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_PC_association_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_Legacy_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_Legacy_SRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_Legacy_CORESET0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_SCellBFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_InterCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_multiMAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_perBWP_CA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_ListSharingCA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_commonMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedSeparateTCI_InterCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UnifiedJointTCI_mTRP_InterCell_BM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mpe_Mitigation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_PortReport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PDCCH_individual_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PDCCH_anySpan_3Symbols_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PDCCH_TwoQCL_TypeD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_CSI_RS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_cyclicMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_secondTPC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_twoPHR_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_A_CSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_SP_CSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUSCH_CG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUCCH_MAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_PUCCH_maxNum_PC_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_inter_Cell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_GroupBasedL1_RSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_BFD_RS_MAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_CSI_EnhancementPerBand_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CodebookComboParameterMultiTRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_CSI_additionalCSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_CSI_N_Max2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_CSI_CMR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Srs_partialFreqSounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamSwitchTiming_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamSwitchTiming_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			BeamReportTiming_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberRxTxBeamSwitchDL_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Srs_increasedRepetition_r17 optional
			if Srs_increasedRepetition_r17Present {
				ie.Srs_increasedRepetition_r17 = new(MIMO_ParametersPerBand_srs_increasedRepetition_r17)
				if err = ie.Srs_increasedRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_increasedRepetition_r17", err)
				}
			}
			// decode Srs_partialFrequencySounding_r17 optional
			if Srs_partialFrequencySounding_r17Present {
				ie.Srs_partialFrequencySounding_r17 = new(MIMO_ParametersPerBand_srs_partialFrequencySounding_r17)
				if err = ie.Srs_partialFrequencySounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_partialFrequencySounding_r17", err)
				}
			}
			// decode Srs_startRB_locationHoppingPartial_r17 optional
			if Srs_startRB_locationHoppingPartial_r17Present {
				ie.Srs_startRB_locationHoppingPartial_r17 = new(MIMO_ParametersPerBand_srs_startRB_locationHoppingPartial_r17)
				if err = ie.Srs_startRB_locationHoppingPartial_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_startRB_locationHoppingPartial_r17", err)
				}
			}
			// decode Srs_combEight_r17 optional
			if Srs_combEight_r17Present {
				ie.Srs_combEight_r17 = new(MIMO_ParametersPerBand_srs_combEight_r17)
				if err = ie.Srs_combEight_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_combEight_r17", err)
				}
			}
			// decode CodebookParametersfetype2_r17 optional
			if CodebookParametersfetype2_r17Present {
				ie.CodebookParametersfetype2_r17 = new(CodebookParametersfetype2_r17)
				if err = ie.CodebookParametersfetype2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookParametersfetype2_r17", err)
				}
			}
			// decode MTRP_PUSCH_twoCSI_RS_r17 optional
			if MTRP_PUSCH_twoCSI_RS_r17Present {
				ie.MTRP_PUSCH_twoCSI_RS_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_twoCSI_RS_r17)
				if err = ie.MTRP_PUSCH_twoCSI_RS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_twoCSI_RS_r17", err)
				}
			}
			// decode MTRP_PUCCH_InterSlot_r17 optional
			if MTRP_PUCCH_InterSlot_r17Present {
				ie.MTRP_PUCCH_InterSlot_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_InterSlot_r17)
				if err = ie.MTRP_PUCCH_InterSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUCCH_InterSlot_r17", err)
				}
			}
			// decode MTRP_PUCCH_CyclicMapping_r17 optional
			if MTRP_PUCCH_CyclicMapping_r17Present {
				ie.MTRP_PUCCH_CyclicMapping_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_CyclicMapping_r17)
				if err = ie.MTRP_PUCCH_CyclicMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUCCH_CyclicMapping_r17", err)
				}
			}
			// decode MTRP_PUCCH_SecondTPC_r17 optional
			if MTRP_PUCCH_SecondTPC_r17Present {
				ie.MTRP_PUCCH_SecondTPC_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_SecondTPC_r17)
				if err = ie.MTRP_PUCCH_SecondTPC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUCCH_SecondTPC_r17", err)
				}
			}
			// decode MTRP_BFR_twoBFD_RS_Set_r17 optional
			if MTRP_BFR_twoBFD_RS_Set_r17Present {
				ie.MTRP_BFR_twoBFD_RS_Set_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17)
				if err = ie.MTRP_BFR_twoBFD_RS_Set_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_BFR_twoBFD_RS_Set_r17", err)
				}
			}
			// decode MTRP_BFR_PUCCH_SR_perCG_r17 optional
			if MTRP_BFR_PUCCH_SR_perCG_r17Present {
				ie.MTRP_BFR_PUCCH_SR_perCG_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_PUCCH_SR_perCG_r17)
				if err = ie.MTRP_BFR_PUCCH_SR_perCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_BFR_PUCCH_SR_perCG_r17", err)
				}
			}
			// decode MTRP_BFR_association_PUCCH_SR_r17 optional
			if MTRP_BFR_association_PUCCH_SR_r17Present {
				ie.MTRP_BFR_association_PUCCH_SR_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_association_PUCCH_SR_r17)
				if err = ie.MTRP_BFR_association_PUCCH_SR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_BFR_association_PUCCH_SR_r17", err)
				}
			}
			// decode Sfn_SimulTwoTCI_AcrossMultiCC_r17 optional
			if Sfn_SimulTwoTCI_AcrossMultiCC_r17Present {
				ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 = new(MIMO_ParametersPerBand_sfn_SimulTwoTCI_AcrossMultiCC_r17)
				if err = ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sfn_SimulTwoTCI_AcrossMultiCC_r17", err)
				}
			}
			// decode Sfn_DefaultDL_BeamSetup_r17 optional
			if Sfn_DefaultDL_BeamSetup_r17Present {
				ie.Sfn_DefaultDL_BeamSetup_r17 = new(MIMO_ParametersPerBand_sfn_DefaultDL_BeamSetup_r17)
				if err = ie.Sfn_DefaultDL_BeamSetup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sfn_DefaultDL_BeamSetup_r17", err)
				}
			}
			// decode Sfn_DefaultUL_BeamSetup_r17 optional
			if Sfn_DefaultUL_BeamSetup_r17Present {
				ie.Sfn_DefaultUL_BeamSetup_r17 = new(MIMO_ParametersPerBand_sfn_DefaultUL_BeamSetup_r17)
				if err = ie.Sfn_DefaultUL_BeamSetup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sfn_DefaultUL_BeamSetup_r17", err)
				}
			}
			// decode Srs_TriggeringOffset_r17 optional
			if Srs_TriggeringOffset_r17Present {
				ie.Srs_TriggeringOffset_r17 = new(MIMO_ParametersPerBand_srs_TriggeringOffset_r17)
				if err = ie.Srs_TriggeringOffset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_TriggeringOffset_r17", err)
				}
			}
			// decode Srs_TriggeringDCI_r17 optional
			if Srs_TriggeringDCI_r17Present {
				ie.Srs_TriggeringDCI_r17 = new(MIMO_ParametersPerBand_srs_TriggeringDCI_r17)
				if err = ie.Srs_TriggeringDCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_TriggeringDCI_r17", err)
				}
			}
			// decode CodebookComboParameterMixedType_r17 optional
			if CodebookComboParameterMixedType_r17Present {
				ie.CodebookComboParameterMixedType_r17 = new(CodebookComboParameterMixedType_r17)
				if err = ie.CodebookComboParameterMixedType_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookComboParameterMixedType_r17", err)
				}
			}
			// decode UnifiedJointTCI_r17 optional
			if UnifiedJointTCI_r17Present {
				ie.UnifiedJointTCI_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_r17)
				if err = ie.UnifiedJointTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_r17", err)
				}
			}
			// decode UnifiedJointTCI_multiMAC_CE_r17 optional
			if UnifiedJointTCI_multiMAC_CE_r17Present {
				ie.UnifiedJointTCI_multiMAC_CE_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17)
				if err = ie.UnifiedJointTCI_multiMAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_multiMAC_CE_r17", err)
				}
			}
			// decode UnifiedJointTCI_perBWP_CA_r17 optional
			if UnifiedJointTCI_perBWP_CA_r17Present {
				ie.UnifiedJointTCI_perBWP_CA_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_perBWP_CA_r17)
				if err = ie.UnifiedJointTCI_perBWP_CA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_perBWP_CA_r17", err)
				}
			}
			// decode UnifiedJointTCI_ListSharingCA_r17 optional
			if UnifiedJointTCI_ListSharingCA_r17Present {
				ie.UnifiedJointTCI_ListSharingCA_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_ListSharingCA_r17)
				if err = ie.UnifiedJointTCI_ListSharingCA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_ListSharingCA_r17", err)
				}
			}
			// decode UnifiedJointTCI_commonMultiCC_r17 optional
			if UnifiedJointTCI_commonMultiCC_r17Present {
				ie.UnifiedJointTCI_commonMultiCC_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_commonMultiCC_r17)
				if err = ie.UnifiedJointTCI_commonMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_commonMultiCC_r17", err)
				}
			}
			// decode UnifiedJointTCI_BeamAlignDLRS_r17 optional
			if UnifiedJointTCI_BeamAlignDLRS_r17Present {
				ie.UnifiedJointTCI_BeamAlignDLRS_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_BeamAlignDLRS_r17)
				if err = ie.UnifiedJointTCI_BeamAlignDLRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_BeamAlignDLRS_r17", err)
				}
			}
			// decode UnifiedJointTCI_PC_association_r17 optional
			if UnifiedJointTCI_PC_association_r17Present {
				ie.UnifiedJointTCI_PC_association_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_PC_association_r17)
				if err = ie.UnifiedJointTCI_PC_association_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_PC_association_r17", err)
				}
			}
			// decode UnifiedJointTCI_Legacy_r17 optional
			if UnifiedJointTCI_Legacy_r17Present {
				ie.UnifiedJointTCI_Legacy_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_r17)
				if err = ie.UnifiedJointTCI_Legacy_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_Legacy_r17", err)
				}
			}
			// decode UnifiedJointTCI_Legacy_SRS_r17 optional
			if UnifiedJointTCI_Legacy_SRS_r17Present {
				ie.UnifiedJointTCI_Legacy_SRS_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_SRS_r17)
				if err = ie.UnifiedJointTCI_Legacy_SRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_Legacy_SRS_r17", err)
				}
			}
			// decode UnifiedJointTCI_Legacy_CORESET0_r17 optional
			if UnifiedJointTCI_Legacy_CORESET0_r17Present {
				ie.UnifiedJointTCI_Legacy_CORESET0_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_CORESET0_r17)
				if err = ie.UnifiedJointTCI_Legacy_CORESET0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_Legacy_CORESET0_r17", err)
				}
			}
			// decode UnifiedJointTCI_SCellBFR_r17 optional
			if UnifiedJointTCI_SCellBFR_r17Present {
				ie.UnifiedJointTCI_SCellBFR_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_SCellBFR_r17)
				if err = ie.UnifiedJointTCI_SCellBFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_SCellBFR_r17", err)
				}
			}
			// decode UnifiedJointTCI_InterCell_r17 optional
			if UnifiedJointTCI_InterCell_r17Present {
				ie.UnifiedJointTCI_InterCell_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17)
				if err = ie.UnifiedJointTCI_InterCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_InterCell_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_r17 optional
			if UnifiedSeparateTCI_r17Present {
				ie.UnifiedSeparateTCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_r17)
				if err = ie.UnifiedSeparateTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_multiMAC_CE_r17 optional
			if UnifiedSeparateTCI_multiMAC_CE_r17Present {
				ie.UnifiedSeparateTCI_multiMAC_CE_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17)
				if err = ie.UnifiedSeparateTCI_multiMAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_multiMAC_CE_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_perBWP_CA_r17 optional
			if UnifiedSeparateTCI_perBWP_CA_r17Present {
				ie.UnifiedSeparateTCI_perBWP_CA_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_perBWP_CA_r17)
				if err = ie.UnifiedSeparateTCI_perBWP_CA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_perBWP_CA_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_ListSharingCA_r17 optional
			if UnifiedSeparateTCI_ListSharingCA_r17Present {
				ie.UnifiedSeparateTCI_ListSharingCA_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17)
				if err = ie.UnifiedSeparateTCI_ListSharingCA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_ListSharingCA_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_commonMultiCC_r17 optional
			if UnifiedSeparateTCI_commonMultiCC_r17Present {
				ie.UnifiedSeparateTCI_commonMultiCC_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_commonMultiCC_r17)
				if err = ie.UnifiedSeparateTCI_commonMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_commonMultiCC_r17", err)
				}
			}
			// decode UnifiedSeparateTCI_InterCell_r17 optional
			if UnifiedSeparateTCI_InterCell_r17Present {
				ie.UnifiedSeparateTCI_InterCell_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17)
				if err = ie.UnifiedSeparateTCI_InterCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedSeparateTCI_InterCell_r17", err)
				}
			}
			// decode UnifiedJointTCI_mTRP_InterCell_BM_r17 optional
			if UnifiedJointTCI_mTRP_InterCell_BM_r17Present {
				ie.UnifiedJointTCI_mTRP_InterCell_BM_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17)
				if err = ie.UnifiedJointTCI_mTRP_InterCell_BM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UnifiedJointTCI_mTRP_InterCell_BM_r17", err)
				}
			}
			// decode Mpe_Mitigation_r17 optional
			if Mpe_Mitigation_r17Present {
				ie.Mpe_Mitigation_r17 = new(MIMO_ParametersPerBand_mpe_Mitigation_r17)
				if err = ie.Mpe_Mitigation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mpe_Mitigation_r17", err)
				}
			}
			// decode Srs_PortReport_r17 optional
			if Srs_PortReport_r17Present {
				ie.Srs_PortReport_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17)
				if err = ie.Srs_PortReport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_PortReport_r17", err)
				}
			}
			// decode MTRP_PDCCH_individual_r17 optional
			if MTRP_PDCCH_individual_r17Present {
				ie.MTRP_PDCCH_individual_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_individual_r17)
				if err = ie.MTRP_PDCCH_individual_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PDCCH_individual_r17", err)
				}
			}
			// decode MTRP_PDCCH_anySpan_3Symbols_r17 optional
			if MTRP_PDCCH_anySpan_3Symbols_r17Present {
				ie.MTRP_PDCCH_anySpan_3Symbols_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_anySpan_3Symbols_r17)
				if err = ie.MTRP_PDCCH_anySpan_3Symbols_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PDCCH_anySpan_3Symbols_r17", err)
				}
			}
			// decode MTRP_PDCCH_TwoQCL_TypeD_r17 optional
			if MTRP_PDCCH_TwoQCL_TypeD_r17Present {
				ie.MTRP_PDCCH_TwoQCL_TypeD_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_TwoQCL_TypeD_r17)
				if err = ie.MTRP_PDCCH_TwoQCL_TypeD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PDCCH_TwoQCL_TypeD_r17", err)
				}
			}
			// decode MTRP_PUSCH_CSI_RS_r17 optional
			if MTRP_PUSCH_CSI_RS_r17Present {
				ie.MTRP_PUSCH_CSI_RS_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17)
				if err = ie.MTRP_PUSCH_CSI_RS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_CSI_RS_r17", err)
				}
			}
			// decode MTRP_PUSCH_cyclicMapping_r17 optional
			if MTRP_PUSCH_cyclicMapping_r17Present {
				ie.MTRP_PUSCH_cyclicMapping_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_cyclicMapping_r17)
				if err = ie.MTRP_PUSCH_cyclicMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_cyclicMapping_r17", err)
				}
			}
			// decode MTRP_PUSCH_secondTPC_r17 optional
			if MTRP_PUSCH_secondTPC_r17Present {
				ie.MTRP_PUSCH_secondTPC_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_secondTPC_r17)
				if err = ie.MTRP_PUSCH_secondTPC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_secondTPC_r17", err)
				}
			}
			// decode MTRP_PUSCH_twoPHR_Reporting_r17 optional
			if MTRP_PUSCH_twoPHR_Reporting_r17Present {
				ie.MTRP_PUSCH_twoPHR_Reporting_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_twoPHR_Reporting_r17)
				if err = ie.MTRP_PUSCH_twoPHR_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_twoPHR_Reporting_r17", err)
				}
			}
			// decode MTRP_PUSCH_A_CSI_r17 optional
			if MTRP_PUSCH_A_CSI_r17Present {
				ie.MTRP_PUSCH_A_CSI_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_A_CSI_r17)
				if err = ie.MTRP_PUSCH_A_CSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_A_CSI_r17", err)
				}
			}
			// decode MTRP_PUSCH_SP_CSI_r17 optional
			if MTRP_PUSCH_SP_CSI_r17Present {
				ie.MTRP_PUSCH_SP_CSI_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_SP_CSI_r17)
				if err = ie.MTRP_PUSCH_SP_CSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_SP_CSI_r17", err)
				}
			}
			// decode MTRP_PUSCH_CG_r17 optional
			if MTRP_PUSCH_CG_r17Present {
				ie.MTRP_PUSCH_CG_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_CG_r17)
				if err = ie.MTRP_PUSCH_CG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUSCH_CG_r17", err)
				}
			}
			// decode MTRP_PUCCH_MAC_CE_r17 optional
			if MTRP_PUCCH_MAC_CE_r17Present {
				ie.MTRP_PUCCH_MAC_CE_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_MAC_CE_r17)
				if err = ie.MTRP_PUCCH_MAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_PUCCH_MAC_CE_r17", err)
				}
			}
			// decode MTRP_PUCCH_maxNum_PC_FR1_r17 optional
			if MTRP_PUCCH_maxNum_PC_FR1_r17Present {
				var tmp_int_MTRP_PUCCH_maxNum_PC_FR1_r17 int64
				if tmp_int_MTRP_PUCCH_maxNum_PC_FR1_r17, err = extReader.ReadInteger(&aper.Constraint{Lb: 3, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode MTRP_PUCCH_maxNum_PC_FR1_r17", err)
				}
				ie.MTRP_PUCCH_maxNum_PC_FR1_r17 = &tmp_int_MTRP_PUCCH_maxNum_PC_FR1_r17
			}
			// decode MTRP_inter_Cell_r17 optional
			if MTRP_inter_Cell_r17Present {
				ie.MTRP_inter_Cell_r17 = new(MIMO_ParametersPerBand_mTRP_inter_Cell_r17)
				if err = ie.MTRP_inter_Cell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_inter_Cell_r17", err)
				}
			}
			// decode MTRP_GroupBasedL1_RSRP_r17 optional
			if MTRP_GroupBasedL1_RSRP_r17Present {
				ie.MTRP_GroupBasedL1_RSRP_r17 = new(MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17)
				if err = ie.MTRP_GroupBasedL1_RSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_GroupBasedL1_RSRP_r17", err)
				}
			}
			// decode MTRP_BFD_RS_MAC_CE_r17 optional
			if MTRP_BFD_RS_MAC_CE_r17Present {
				ie.MTRP_BFD_RS_MAC_CE_r17 = new(MIMO_ParametersPerBand_mTRP_BFD_RS_MAC_CE_r17)
				if err = ie.MTRP_BFD_RS_MAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_BFD_RS_MAC_CE_r17", err)
				}
			}
			// decode MTRP_CSI_EnhancementPerBand_r17 optional
			if MTRP_CSI_EnhancementPerBand_r17Present {
				ie.MTRP_CSI_EnhancementPerBand_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17)
				if err = ie.MTRP_CSI_EnhancementPerBand_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_CSI_EnhancementPerBand_r17", err)
				}
			}
			// decode CodebookComboParameterMultiTRP_r17 optional
			if CodebookComboParameterMultiTRP_r17Present {
				ie.CodebookComboParameterMultiTRP_r17 = new(CodebookComboParameterMultiTRP_r17)
				if err = ie.CodebookComboParameterMultiTRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookComboParameterMultiTRP_r17", err)
				}
			}
			// decode MTRP_CSI_additionalCSI_r17 optional
			if MTRP_CSI_additionalCSI_r17Present {
				ie.MTRP_CSI_additionalCSI_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_additionalCSI_r17)
				if err = ie.MTRP_CSI_additionalCSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_CSI_additionalCSI_r17", err)
				}
			}
			// decode MTRP_CSI_N_Max2_r17 optional
			if MTRP_CSI_N_Max2_r17Present {
				ie.MTRP_CSI_N_Max2_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_N_Max2_r17)
				if err = ie.MTRP_CSI_N_Max2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_CSI_N_Max2_r17", err)
				}
			}
			// decode MTRP_CSI_CMR_r17 optional
			if MTRP_CSI_CMR_r17Present {
				ie.MTRP_CSI_CMR_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_CMR_r17)
				if err = ie.MTRP_CSI_CMR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_CSI_CMR_r17", err)
				}
			}
			// decode Srs_partialFreqSounding_r17 optional
			if Srs_partialFreqSounding_r17Present {
				ie.Srs_partialFreqSounding_r17 = new(MIMO_ParametersPerBand_srs_partialFreqSounding_r17)
				if err = ie.Srs_partialFreqSounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_partialFreqSounding_r17", err)
				}
			}
			// decode BeamSwitchTiming_v1710 optional
			if BeamSwitchTiming_v1710Present {
				ie.BeamSwitchTiming_v1710 = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710)
				if err = ie.BeamSwitchTiming_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamSwitchTiming_v1710", err)
				}
			}
			// decode BeamSwitchTiming_r17 optional
			if BeamSwitchTiming_r17Present {
				ie.BeamSwitchTiming_r17 = new(MIMO_ParametersPerBand_beamSwitchTiming_r17)
				if err = ie.BeamSwitchTiming_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamSwitchTiming_r17", err)
				}
			}
			// decode BeamReportTiming_v1710 optional
			if BeamReportTiming_v1710Present {
				ie.BeamReportTiming_v1710 = new(MIMO_ParametersPerBand_beamReportTiming_v1710)
				if err = ie.BeamReportTiming_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode BeamReportTiming_v1710", err)
				}
			}
			// decode MaxNumberRxTxBeamSwitchDL_v1710 optional
			if MaxNumberRxTxBeamSwitchDL_v1710Present {
				ie.MaxNumberRxTxBeamSwitchDL_v1710 = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710)
				if err = ie.MaxNumberRxTxBeamSwitchDL_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberRxTxBeamSwitchDL_v1710", err)
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

			Srs_PortReportSP_AP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberRxBeam_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sfn_ImplicitRS_twoTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Sfn_QCL_TypeD_Collision_twoTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MTRP_CSI_numCPU_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Srs_PortReportSP_AP_r17 optional
			if Srs_PortReportSP_AP_r17Present {
				ie.Srs_PortReportSP_AP_r17 = new(MIMO_ParametersPerBand_srs_PortReportSP_AP_r17)
				if err = ie.Srs_PortReportSP_AP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Srs_PortReportSP_AP_r17", err)
				}
			}
			// decode MaxNumberRxBeam_v1720 optional
			if MaxNumberRxBeam_v1720Present {
				var tmp_int_MaxNumberRxBeam_v1720 int64
				if tmp_int_MaxNumberRxBeam_v1720, err = extReader.ReadInteger(&aper.Constraint{Lb: 9, Ub: 12}, false); err != nil {
					return utils.WrapError("Decode MaxNumberRxBeam_v1720", err)
				}
				ie.MaxNumberRxBeam_v1720 = &tmp_int_MaxNumberRxBeam_v1720
			}
			// decode Sfn_ImplicitRS_twoTCI_r17 optional
			if Sfn_ImplicitRS_twoTCI_r17Present {
				ie.Sfn_ImplicitRS_twoTCI_r17 = new(MIMO_ParametersPerBand_sfn_ImplicitRS_twoTCI_r17)
				if err = ie.Sfn_ImplicitRS_twoTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sfn_ImplicitRS_twoTCI_r17", err)
				}
			}
			// decode Sfn_QCL_TypeD_Collision_twoTCI_r17 optional
			if Sfn_QCL_TypeD_Collision_twoTCI_r17Present {
				ie.Sfn_QCL_TypeD_Collision_twoTCI_r17 = new(MIMO_ParametersPerBand_sfn_QCL_TypeD_Collision_twoTCI_r17)
				if err = ie.Sfn_QCL_TypeD_Collision_twoTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Sfn_QCL_TypeD_Collision_twoTCI_r17", err)
				}
			}
			// decode MTRP_CSI_numCPU_r17 optional
			if MTRP_CSI_numCPU_r17Present {
				ie.MTRP_CSI_numCPU_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_numCPU_r17)
				if err = ie.MTRP_CSI_numCPU_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode MTRP_CSI_numCPU_r17", err)
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

			SupportRepNumPDSCH_TDRA_DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SupportRepNumPDSCH_TDRA_DCI_1_2_r17 optional
			if SupportRepNumPDSCH_TDRA_DCI_1_2_r17Present {
				ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 = new(MIMO_ParametersPerBand_supportRepNumPDSCH_TDRA_DCI_1_2_r17)
				if err = ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportRepNumPDSCH_TDRA_DCI_1_2_r17", err)
				}
			}
		}
	}
	return nil
}
