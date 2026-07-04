// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MIMO-ParametersPerBand (line 21479).

var mIMOParametersPerBandConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tci-StatePDSCH", Optional: true},
		{Name: "additionalActiveTCI-StatePDCCH", Optional: true},
		{Name: "pusch-TransCoherence", Optional: true},
		{Name: "beamCorrespondenceWithoutUL-BeamSweeping", Optional: true},
		{Name: "periodicBeamReport", Optional: true},
		{Name: "aperiodicBeamReport", Optional: true},
		{Name: "sp-BeamReportPUCCH", Optional: true},
		{Name: "sp-BeamReportPUSCH", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "maxNumberRxBeam", Optional: true},
		{Name: "maxNumberRxTxBeamSwitchDL", Optional: true},
		{Name: "maxNumberNonGroupBeamReporting", Optional: true},
		{Name: "groupBeamReporting", Optional: true},
		{Name: "uplinkBeamManagement", Optional: true},
		{Name: "maxNumberCSI-RS-BFD", Optional: true},
		{Name: "maxNumberSSB-BFD", Optional: true},
		{Name: "maxNumberCSI-RS-SSB-CBD", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "twoPortsPTRS-UL", Optional: true},
		{Name: "dummy5", Optional: true},
		{Name: "dummy3", Optional: true},
		{Name: "beamReportTiming", Optional: true},
		{Name: "ptrs-DensityRecommendationSetDL", Optional: true},
		{Name: "ptrs-DensityRecommendationSetUL", Optional: true},
		{Name: "dummy4", Optional: true},
		{Name: "aperiodicTRS", Optional: true},
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
	},
}

const (
	MIMO_ParametersPerBand_AdditionalActiveTCI_StatePDCCH_Supported = 0
)

var mIMOParametersPerBandAdditionalActiveTCIStatePDCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_AdditionalActiveTCI_StatePDCCH_Supported},
}

const (
	MIMO_ParametersPerBand_Pusch_TransCoherence_NonCoherent     = 0
	MIMO_ParametersPerBand_Pusch_TransCoherence_PartialCoherent = 1
	MIMO_ParametersPerBand_Pusch_TransCoherence_FullCoherent    = 2
)

var mIMOParametersPerBandPuschTransCoherenceConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Pusch_TransCoherence_NonCoherent, MIMO_ParametersPerBand_Pusch_TransCoherence_PartialCoherent, MIMO_ParametersPerBand_Pusch_TransCoherence_FullCoherent},
}

const (
	MIMO_ParametersPerBand_BeamCorrespondenceWithoutUL_BeamSweeping_Supported = 0
)

var mIMOParametersPerBandBeamCorrespondenceWithoutULBeamSweepingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_BeamCorrespondenceWithoutUL_BeamSweeping_Supported},
}

const (
	MIMO_ParametersPerBand_PeriodicBeamReport_Supported = 0
)

var mIMOParametersPerBandPeriodicBeamReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_PeriodicBeamReport_Supported},
}

const (
	MIMO_ParametersPerBand_AperiodicBeamReport_Supported = 0
)

var mIMOParametersPerBandAperiodicBeamReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_AperiodicBeamReport_Supported},
}

const (
	MIMO_ParametersPerBand_Sp_BeamReportPUCCH_Supported = 0
)

var mIMOParametersPerBandSpBeamReportPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Sp_BeamReportPUCCH_Supported},
}

const (
	MIMO_ParametersPerBand_Sp_BeamReportPUSCH_Supported = 0
)

var mIMOParametersPerBandSpBeamReportPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Sp_BeamReportPUSCH_Supported},
}

var mIMOParametersPerBandMaxNumberRxBeamConstraints = per.Constrained(2, 8)

const (
	MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N1 = 0
	MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N2 = 1
	MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N4 = 2
)

var mIMOParametersPerBandMaxNumberNonGroupBeamReportingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N1, MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N2, MIMO_ParametersPerBand_MaxNumberNonGroupBeamReporting_N4},
}

const (
	MIMO_ParametersPerBand_GroupBeamReporting_Supported = 0
)

var mIMOParametersPerBandGroupBeamReportingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_GroupBeamReporting_Supported},
}

var mIMOParametersPerBandMaxNumberCSIRSBFDConstraints = per.Constrained(1, 64)

var mIMOParametersPerBandMaxNumberSSBBFDConstraints = per.Constrained(1, 64)

var mIMOParametersPerBandMaxNumberCSIRSSSBCBDConstraints = per.Constrained(1, 256)

const (
	MIMO_ParametersPerBand_Dummy2_Supported = 0
)

var mIMOParametersPerBandDummy2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Dummy2_Supported},
}

const (
	MIMO_ParametersPerBand_TwoPortsPTRS_UL_Supported = 0
)

var mIMOParametersPerBandTwoPortsPTRSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_TwoPortsPTRS_UL_Supported},
}

var mIMOParametersPerBandDummy3Constraints = per.Constrained(1, 4)

const (
	MIMO_ParametersPerBand_AperiodicTRS_Supported = 0
)

var mIMOParametersPerBandAperiodicTRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_AperiodicTRS_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Dummy6_True = 0
)

var mIMOParametersPerBandExtDummy6Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Dummy6_True},
}

var mIMOParametersPerBandExtSrsAssocCSIRSConstraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	MIMO_ParametersPerBand_Ext_DefaultQCL_TwoTCI_r16_Supported = 0
)

var mIMOParametersPerBandExtDefaultQCLTwoTCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_DefaultQCL_TwoTCI_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Simul_SpatialRelationUpdatePUCCHResGroup_r16_Supported = 0
)

var mIMOParametersPerBandExtSimulSpatialRelationUpdatePUCCHResGroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Simul_SpatialRelationUpdatePUCCHResGroup_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N1 = 0
	MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N2 = 1
	MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N4 = 2
	MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N8 = 3
)

var mIMOParametersPerBandExtMaxNumberSCellBFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N1, MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N2, MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N4, MIMO_ParametersPerBand_Ext_MaxNumberSCellBFR_r16_N8},
}

const (
	MIMO_ParametersPerBand_Ext_SimultaneousReceptionDiffTypeD_r16_Supported = 0
)

var mIMOParametersPerBandExtSimultaneousReceptionDiffTypeDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SimultaneousReceptionDiffTypeD_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N1 = 0
	MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N2 = 1
	MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N4 = 2
)

var mIMOParametersPerBandExtNonGroupSINRReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N1, MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N2, MIMO_ParametersPerBand_Ext_NonGroupSINR_Reporting_r16_N4},
}

const (
	MIMO_ParametersPerBand_Ext_GroupSINR_Reporting_r16_Supported = 0
)

var mIMOParametersPerBandExtGroupSINRReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_GroupSINR_Reporting_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SupportFDM_SchemeA_r16_Supported = 0
)

var mIMOParametersPerBandExtSupportFDMSchemeAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportFDM_SchemeA_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SupportCodeWordSoftCombining_r16_Supported = 0
)

var mIMOParametersPerBandExtSupportCodeWordSoftCombiningR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportCodeWordSoftCombining_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb3           = 0
	MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb5           = 1
	MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb10          = 2
	MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb20          = 3
	MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_NoRestriction = 4
)

var mIMOParametersPerBandExtSupportTDMSchemeAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb3, MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb5, MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb10, MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_Kb20, MIMO_ParametersPerBand_Ext_SupportTDM_SchemeA_r16_NoRestriction},
}

const (
	MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PDSCH_r16_Supported = 0
)

var mIMOParametersPerBandExtLowPAPRDMRSPDSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PDSCH_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUSCHwithoutPrecoding_r16_Supported = 0
)

var mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithoutPrecodingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUSCHwithoutPrecoding_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUCCH_r16_Supported = 0
)

var mIMOParametersPerBandExtLowPAPRDMRSPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUCCH_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUSCHwithPrecoding_r16_Supported = 0
)

var mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithPrecodingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_LowPAPR_DMRS_PUSCHwithPrecoding_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_BeamCorrespondenceSSB_Based_r16_Supported = 0
)

var mIMOParametersPerBandExtBeamCorrespondenceSSBBasedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamCorrespondenceSSB_Based_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_BeamCorrespondenceCSI_RS_Based_r16_Supported = 0
)

var mIMOParametersPerBandExtBeamCorrespondenceCSIRSBasedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamCorrespondenceCSI_RS_Based_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUSCH_r16_Supported = 0
)

var mIMOParametersPerBandExtSemiPersistentL1SINRReportPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUSCH_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Support64CandidateBeamRS_BFR_r16_Supported = 0
)

var mIMOParametersPerBandExtSupport64CandidateBeamRSBFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Support64CandidateBeamRS_BFR_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MaxMIMO_LayersForMulti_DCI_MTRP_r16_Supported = 0
)

var mIMOParametersPerBandExtMaxMIMOLayersForMultiDCIMTRPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MaxMIMO_LayersForMulti_DCI_MTRP_r16_Supported},
}

var mIMOParametersPerBandSupportedSINRMeasV1670Constraints = per.FixedSize(4)

const (
	MIMO_ParametersPerBand_Ext_Srs_IncreasedRepetition_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsIncreasedRepetitionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_IncreasedRepetition_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PartialFrequencySounding_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsPartialFrequencySoundingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PartialFrequencySounding_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_StartRB_LocationHoppingPartial_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsStartRBLocationHoppingPartialR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_StartRB_LocationHoppingPartial_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CombEight_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsCombEightR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CombEight_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_TwoCSI_RS_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHTwoCSIRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_TwoCSI_RS_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf0_2   = 0
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf1_3_4 = 1
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf0_4   = 2
)

var mIMOParametersPerBandExtMTRPPUCCHInterSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf0_2, MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf1_3_4, MIMO_ParametersPerBand_Ext_MTRP_PUCCH_InterSlot_r17_Pf0_4},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_CyclicMapping_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUCCHCyclicMappingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUCCH_CyclicMapping_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_SecondTPC_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUCCHSecondTPCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUCCH_SecondTPC_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_BFR_PUCCH_SR_PerCG_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_MTRP_BFR_PUCCH_SR_PerCG_r17_N2 = 1
)

var mIMOParametersPerBandExtMTRPBFRPUCCHSRPerCGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_BFR_PUCCH_SR_PerCG_r17_N1, MIMO_ParametersPerBand_Ext_MTRP_BFR_PUCCH_SR_PerCG_r17_N2},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_BFR_Association_PUCCH_SR_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPBFRAssociationPUCCHSRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_BFR_Association_PUCCH_SR_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Sfn_SimulTwoTCI_AcrossMultiCC_r17_Supported = 0
)

var mIMOParametersPerBandExtSfnSimulTwoTCIAcrossMultiCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Sfn_SimulTwoTCI_AcrossMultiCC_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Sfn_DefaultDL_BeamSetup_r17_Supported = 0
)

var mIMOParametersPerBandExtSfnDefaultDLBeamSetupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Sfn_DefaultDL_BeamSetup_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Sfn_DefaultUL_BeamSetup_r17_Supported = 0
)

var mIMOParametersPerBandExtSfnDefaultULBeamSetupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Sfn_DefaultUL_BeamSetup_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N4 = 2
)

var mIMOParametersPerBandExtSrsTriggeringOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N1, MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N2, MIMO_ParametersPerBand_Ext_Srs_TriggeringOffset_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_TriggeringDCI_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsTriggeringDCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_TriggeringDCI_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_PerBWP_CA_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCIPerBWPCAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_PerBWP_CA_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N4 = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N8 = 3
)

var mIMOParametersPerBandExtUnifiedJointTCIListSharingCAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_ListSharingCA_r17_N8},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_CommonMultiCC_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCICommonMultiCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_CommonMultiCC_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_BeamAlignDLRS_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCIBeamAlignDLRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_BeamAlignDLRS_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_PC_Association_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCIPCAssociationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_PC_Association_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCILegacyR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_SRS_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCILegacySRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_SRS_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_CORESET0_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCILegacyCORESET0R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_Legacy_CORESET0_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_SCellBFR_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedJointTCISCellBFRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_SCellBFR_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_PerBWP_CA_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedSeparateTCIPerBWPCAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_PerBWP_CA_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_CommonMultiCC_r17_Supported = 0
)

var mIMOParametersPerBandExtUnifiedSeparateTCICommonMultiCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_CommonMultiCC_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PDCCH_Individual_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPDCCHIndividualR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PDCCH_Individual_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PDCCH_AnySpan_3Symbols_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPDCCHAnySpan3SymbolsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PDCCH_AnySpan_3Symbols_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PDCCH_TwoQCL_TypeD_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPDCCHTwoQCLTypeDR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PDCCH_TwoQCL_TypeD_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_TypeA = 0
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_TypeB = 1
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_Both  = 2
)

var mIMOParametersPerBandExtMTRPPUSCHCyclicMappingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_TypeA, MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_TypeB, MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CyclicMapping_r17_Both},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_SecondTPC_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHSecondTPCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_SecondTPC_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_TwoPHR_Reporting_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHTwoPHRReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_TwoPHR_Reporting_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_A_CSI_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHACSIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_A_CSI_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_SP_CSI_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHSPCSIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_SP_CSI_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CG_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUSCHCGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUSCH_CG_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_PUCCH_MAC_CE_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPPUCCHMACCER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_PUCCH_MAC_CE_r17_Supported},
}

var mIMOParametersPerBandMTRPPUCCHMaxNumPCFR1R17Constraints = per.Constrained(3, 8)

const (
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N4  = 0
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N8  = 1
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N12 = 2
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N16 = 3
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N32 = 4
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N48 = 5
	MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N64 = 6
)

var mIMOParametersPerBandExtMTRPBFDRSMACCER17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N4, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N8, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N12, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N16, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N32, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N48, MIMO_ParametersPerBand_Ext_MTRP_BFD_RS_MAC_CE_r17_N64},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_AdditionalCSI_r17_X1 = 0
	MIMO_ParametersPerBand_Ext_MTRP_CSI_AdditionalCSI_r17_X2 = 1
)

var mIMOParametersPerBandExtMTRPCSIAdditionalCSIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_AdditionalCSI_r17_X1, MIMO_ParametersPerBand_Ext_MTRP_CSI_AdditionalCSI_r17_X2},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_N_Max2_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPCSINMax2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_N_Max2_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_CMR_r17_Supported = 0
)

var mIMOParametersPerBandExtMTRPCSICMRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_CMR_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PartialFreqSounding_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsPartialFreqSoundingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PartialFreqSounding_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PortReportSP_AP_r17_Supported = 0
)

var mIMOParametersPerBandExtSrsPortReportSPAPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PortReportSP_AP_r17_Supported},
}

var mIMOParametersPerBandMaxNumberRxBeamV1720Constraints = per.Constrained(9, 12)

const (
	MIMO_ParametersPerBand_Ext_Sfn_ImplicitRS_TwoTCI_r17_Supported = 0
)

var mIMOParametersPerBandExtSfnImplicitRSTwoTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Sfn_ImplicitRS_TwoTCI_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Sfn_QCL_TypeD_Collision_TwoTCI_r17_Supported = 0
)

var mIMOParametersPerBandExtSfnQCLTypeDCollisionTwoTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Sfn_QCL_TypeD_Collision_TwoTCI_r17_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N2 = 0
	MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N3 = 1
	MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N4 = 2
)

var mIMOParametersPerBandExtMTRPCSINumCPUR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N2, MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N3, MIMO_ParametersPerBand_Ext_MTRP_CSI_NumCPU_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N2  = 0
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N3  = 1
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N5  = 3
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N6  = 4
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N7  = 5
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N8  = 6
	MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N16 = 7
)

var mIMOParametersPerBandExtSupportRepNumPDSCHTDRADCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N2, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N3, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N4, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N5, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N6, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N7, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N8, MIMO_ParametersPerBand_Ext_SupportRepNumPDSCH_TDRA_DCI_1_2_r17_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SelectionDCI_r18_Supported = 0
)

var mIMOParametersPerBandExtTciSelectionDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SelectionDCI_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_PerResource    = 0
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_PerResourceSet = 1
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_Both           = 2
)

var mIMOParametersPerBandExtTciSelectionAperiodicCSIRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_PerResource, MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_PerResourceSet, MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_r18_Both},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_PerResource    = 0
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_PerResourceSet = 1
	MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_Both           = 2
)

var mIMOParametersPerBandExtTciSelectionAperiodicCSIRSMDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_PerResource, MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_PerResourceSet, MIMO_ParametersPerBand_Ext_Tci_SelectionAperiodicCSI_RS_M_DCI_r18_Both},
}

const (
	MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_CjtSchemeA = 0
	MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_CjtSchemeB = 1
	MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_Both       = 2
)

var mIMOParametersPerBandExtTwoTCIStatePDSCHCJTTxSchemeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_CjtSchemeA, MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_CjtSchemeB, MIMO_ParametersPerBand_Ext_TwoTCI_StatePDSCH_CJT_TxScheme_r18_Both},
}

var mIMOParametersPerBandTciJointTCIUpdateMultiActiveTCIPerCCPerCORESETR18Constraints = per.Constrained(2, 8)

const (
	MIMO_ParametersPerBand_Ext_Tci_TRP_BFR_r18_Supported = 0
)

var mIMOParametersPerBandExtTciTRPBFRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_TRP_BFR_r18_Supported},
}

var mIMOParametersPerBandCommonTCISingleDCIR18Constraints = per.Constrained(1, 4)

var mIMOParametersPerBandCommonTCIMultiDCIR18Constraints = per.Constrained(1, 4)

const (
	MIMO_ParametersPerBand_Ext_TwoPHR_Reporting_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPHRReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPHR_Reporting_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SpCell_TAG_Ind_r18_Supported = 0
)

var mIMOParametersPerBandExtSpCellTAGIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SpCell_TAG_Ind_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_InterCellCrossTRP_PDCCH_OrderCFRA_r18_Supported = 0
)

var mIMOParametersPerBandExtInterCellCrossTRPPDCCHOrderCFRAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_InterCellCrossTRP_PDCCH_OrderCFRA_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_IntraCellCrossTRP_PDCCH_OrderCFRA_r18_Supported = 0
)

var mIMOParametersPerBandExtIntraCellCrossTRPPDCCHOrderCFRAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_IntraCellCrossTRP_PDCCH_OrderCFRA_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_OverlapUL_TransReduction_r18_Supported = 0
)

var mIMOParametersPerBandExtOverlapULTransReductionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_OverlapUL_TransReduction_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl4  = 0
	MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl5  = 1
	MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl8  = 2
	MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl10 = 3
	MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl20 = 4
)

var mIMOParametersPerBandExtMaxPeriodicityCMRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl4, MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl5, MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl8, MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl10, MIMO_ParametersPerBand_Ext_MaxPeriodicityCMR_r18_Sl20},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJT_CSI_r18_N0 = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJT_CSI_r18_N2 = 1
)

var mIMOParametersPerBandExtTimelineRelaxCJTCSIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJT_CSI_r18_N0, MIMO_ParametersPerBand_Ext_TimelineRelax_CJT_CSI_r18_N2},
}

const (
	MIMO_ParametersPerBand_Ext_JointConfigDMRSPortDynamicSwitching_r18_Supported = 0
)

var mIMOParametersPerBandExtJointConfigDMRSPortDynamicSwitchingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_JointConfigDMRSPortDynamicSwitching_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetHopping_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCombOffsetHoppingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CombOffsetHopping_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Srs  = 0
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Rsrs = 1
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Both = 2
)

var mIMOParametersPerBandExtSrsCombOffsetInTimeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Srs, MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Rsrs, MIMO_ParametersPerBand_Ext_Srs_CombOffsetInTime_r18_Both},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetCombinedGroupSequence_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCombOffsetCombinedGroupSequenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CombOffsetCombinedGroupSequence_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CombOffsetHoppingWithinSubset_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCombOffsetHoppingWithinSubsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CombOffsetHoppingWithinSubset_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CyclicShiftHopping_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCyclicShiftHoppingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CyclicShiftHopping_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CyclicShiftHoppingSmallGranularity_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCyclicShiftHoppingSmallGranularityR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CyclicShiftHoppingSmallGranularity_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CyclicShiftCombinedGroupSequence_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCyclicShiftCombinedGroupSequenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CyclicShiftCombinedGroupSequence_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_CyclicShiftHoppingWithinSubset_r18_Supported = 0
)

var mIMOParametersPerBandExtCyclicShiftHoppingWithinSubsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_CyclicShiftHoppingWithinSubset_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_CyclicShiftCombinedCombOffset_r18_Supported = 0
)

var mIMOParametersPerBandExtSrsCyclicShiftCombinedCombOffsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_CyclicShiftCombinedCombOffset_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18_Supported = 0
)

var mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18_Supported = 0
)

var mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Dmrs_PortEntrySingleDCI_SDM_r18_Supported = 0
)

var mIMOParametersPerBandExtDmrsPortEntrySingleDCISDMR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Dmrs_PortEntrySingleDCI_SDM_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18_Supported = 0
)

var mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSFNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18_Supported = 0
)

var mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSFNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGDGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18_Supported = 0
)

var mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGDGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Pucch_RepetitionDynamicIndicationSFN_r18_Supported = 0
)

var mIMOParametersPerBandExtPucchRepetitionDynamicIndicationSFNR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Pucch_RepetitionDynamicIndicationSFN_r18_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SimulConfigDMRS_DCI_1_3_r18_Supported = 0
)

var mIMOParametersPerBandExtSimulConfigDMRSDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SimulConfigDMRS_DCI_1_3_r18_Supported},
}

var mIMOParametersPerBandExtAimlCSIReportR19Constraints = per.SizeRange(1, 2)

const (
	MIMO_ParametersPerBand_Ext_IncreasedNumOfReportedRS_r19_N6 = 0
	MIMO_ParametersPerBand_Ext_IncreasedNumOfReportedRS_r19_N8 = 1
)

var mIMOParametersPerBandExtIncreasedNumOfReportedRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_IncreasedNumOfReportedRS_r19_N6, MIMO_ParametersPerBand_Ext_IncreasedNumOfReportedRS_r19_N8},
}

var mIMOParametersPerBandAimlBMCase1PredictedRSRPR19Constraints = per.Constrained(1, 4)

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_PredictionUnitDurationDD_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlCSIPredictionUnitDurationDDR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_PredictionUnitDurationDD_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_PredictionUE_DataCollection_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlCSIPredictionUEDataCollectionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_PredictionUE_DataCollection_r19_Supported},
}

var mIMOParametersPerBandUeiModeAEvent2R19Constraints = per.Constrained(1, 64)

var mIMOParametersPerBandUeiTriggerEventDeterminationR19Constraints = per.Constrained(1, 64)

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeA_Event1_r19_Supported = 0
)

var mIMOParametersPerBandExtUeiModeAEvent1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeA_Event1_r19_Supported},
}

var mIMOParametersPerBandUeiModeAEvent7R19Constraints = per.FixedSize(8)

const (
	MIMO_ParametersPerBand_Ext_Event2ConditionIndication_r19_Supported = 0
)

var mIMOParametersPerBandExtEvent2ConditionIndicationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Event2ConditionIndication_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TimeRestriction128Port_r19_Supported = 0
)

var mIMOParametersPerBandExtTimeRestriction128PortR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimeRestriction128Port_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_GroupScalingFactor_r19_Rank1     = 0
	MIMO_ParametersPerBand_Ext_GroupScalingFactor_r19_Rank1and2 = 1
)

var mIMOParametersPerBandExtGroupScalingFactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_GroupScalingFactor_r19_Rank1, MIMO_ParametersPerBand_Ext_GroupScalingFactor_r19_Rank1and2},
}

const (
	MIMO_ParametersPerBand_Ext_Mr_AlwaysReportedType1SP_r19_Supported = 0
)

var mIMOParametersPerBandExtMrAlwaysReportedType1SPR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Mr_AlwaysReportedType1SP_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Mr_AlwaysReported_EType2_r19_Supported = 0
)

var mIMOParametersPerBandExtMrAlwaysReportedEType2R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Mr_AlwaysReported_EType2_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeC_r19_Supported = 0
)

var mIMOParametersPerBandExtCjtQCLPDSCHSchemeCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeC_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeD_r19_Supported = 0
)

var mIMOParametersPerBandExtCjtQCLPDSCHSchemeDR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeD_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeE_r19_Supported = 0
)

var mIMOParametersPerBandExtCjtQCLPDSCHSchemeER19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjt_QCL_PDSCH_SchemeE_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_Joint_r19_Supported = 0
)

var mIMOParametersPerBandExtLinkedCJTCDdEType2CJTJointR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_Joint_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_Separate_r19_Supported = 0
)

var mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_Separate_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19_Supported = 0
)

var mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparatePerStateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19_Supported},
}

var mIMOParametersPerBandExtNonCodebookCSIRSSRS3TxPUSCHR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_Resources)

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSJointTCIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSSeparateTCIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetPRACH_JointTCI_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetPRACHJointTCIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetPRACH_JointTCI_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetPRACH_SeparateTCI_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetPRACHSeparateTCIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetPRACH_SeparateTCI_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_ExtendedStartBitDCI_2_3_r19_Supported = 0
)

var mIMOParametersPerBandExtExtendedStartBitDCI23R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_ExtendedStartBitDCI_2_3_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoSRS_PwrControlAdjust_r19_Supported = 0
)

var mIMOParametersPerBandExtTwoSRSPwrControlAdjustR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoSRS_PwrControlAdjust_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_OverlapUL_TransReductionEnh_r19_Supported = 0
)

var mIMOParametersPerBandExtOverlapULTransReductionEnhR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_OverlapUL_TransReductionEnh_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetUpdate_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetUpdateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetUpdate_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoSRS_TPC_DCI_2_3_r19_Supported = 0
)

var mIMOParametersPerBandExtTwoSRSTPCDCI23R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoSRS_TPC_DCI_2_3_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_TPC_CLPC_AdjustmentState_r19_Supported = 0
)

var mIMOParametersPerBandExtSrsTPCCLPCAdjustmentStateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_TPC_CLPC_AdjustmentState_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoSRS_DCI_1_1_Separate_r19_Supported = 0
)

var mIMOParametersPerBandExtTwoSRSDCI11SeparateR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoSRS_DCI_1_1_Separate_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_TwoSRS_DCI_1_1_Joint_r19_Supported = 0
)

var mIMOParametersPerBandExtTwoSRSDCI11JointR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TwoSRS_DCI_1_1_Joint_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_PathlossOffsetPHR_r19_Supported = 0
)

var mIMOParametersPerBandExtPathlossOffsetPHRR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_PathlossOffsetPHR_r19_Supported},
}

var mIMOParametersPerBandTciStatePDSCHConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberConfiguredTCI-StatesPerCC", Optional: true},
		{Name: "maxNumberActiveTCI-PerBWP", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N4   = 0
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N8   = 1
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N16  = 2
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N32  = 3
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N64  = 4
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N128 = 5
)

var mIMOParametersPerBandTciStatePDSCHMaxNumberConfiguredTCIStatesPerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N4, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N8, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N16, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N32, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N64, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberConfiguredTCI_StatesPerCC_N128},
}

const (
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N1 = 0
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N2 = 1
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N4 = 2
	MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N8 = 3
)

var mIMOParametersPerBandTciStatePDSCHMaxNumberActiveTCIPerBWPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N1, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N2, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N4, MIMO_ParametersPerBand_Tci_StatePDSCH_MaxNumberActiveTCI_PerBWP_N8},
}

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
		{Name: "scs-240kHz", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N4  = 0
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N7  = 1
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N14 = 2
)

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N4, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N7, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_15kHz_N14},
}

const (
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N4  = 0
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N7  = 1
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N14 = 2
)

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N4, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N7, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_30kHz_N14},
}

const (
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N4  = 0
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N7  = 1
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N14 = 2
)

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N4, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N7, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_60kHz_N14},
}

const (
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N4  = 0
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N7  = 1
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N14 = 2
)

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N4, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N7, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_120kHz_N14},
}

const (
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N4  = 0
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N7  = 1
	MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N14 = 2
)

var mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs240kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N4, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N7, MIMO_ParametersPerBand_MaxNumberRxTxBeamSwitchDL_Scs_240kHz_N14},
}

const (
	MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N2  = 0
	MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N4  = 1
	MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N8  = 2
	MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N16 = 3
)

var mIMOParametersPerBandUplinkBeamManagementMaxNumberSRSResourcePerSetBMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N2, MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N4, MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N8, MIMO_ParametersPerBand_UplinkBeamManagement_MaxNumberSRS_ResourcePerSet_BM_N16},
}

var mIMOParametersPerBandBeamReportTimingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym2 = 0
	MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym4 = 1
	MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym8 = 2
)

var mIMOParametersPerBandBeamReportTimingScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym2, MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym4, MIMO_ParametersPerBand_BeamReportTiming_Scs_15kHz_Sym8},
}

const (
	MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym4  = 0
	MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym8  = 1
	MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym14 = 2
	MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym28 = 3
)

var mIMOParametersPerBandBeamReportTimingScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym4, MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym8, MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym14, MIMO_ParametersPerBand_BeamReportTiming_Scs_30kHz_Sym28},
}

const (
	MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym8  = 0
	MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym14 = 1
	MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym28 = 2
)

var mIMOParametersPerBandBeamReportTimingScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym8, MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym14, MIMO_ParametersPerBand_BeamReportTiming_Scs_60kHz_Sym28},
}

const (
	MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym14 = 0
	MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym28 = 1
	MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym56 = 2
)

var mIMOParametersPerBandBeamReportTimingScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym14, MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym28, MIMO_ParametersPerBand_BeamReportTiming_Scs_120kHz_Sym56},
}

var mIMOParametersPerBandPtrsDensityRecommendationSetDLConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var mIMOParametersPerBandPtrsDensityRecommendationSetULConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var mIMOParametersPerBandExtBeamSwitchTimingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym14  = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym28  = 1
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym48  = 2
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym224 = 3
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym336 = 4
)

var mIMOParametersPerBandExtBeamSwitchTimingScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym14, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym28, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym48, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym224, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_60kHz_Sym336},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym14  = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym28  = 1
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym48  = 2
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym224 = 3
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym336 = 4
)

var mIMOParametersPerBandExtBeamSwitchTimingScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym14, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym28, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym48, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym224, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_Scs_120kHz_Sym336},
}

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSSB-CSIRS-OneTx-CMR-r16"},
		{Name: "maxNumberCSI-IM-NZP-IMR-res-r16"},
		{Name: "maxNumberCSIRS-2Tx-res-r16"},
		{Name: "maxNumberSSB-CSIRS-res-r16"},
		{Name: "maxNumberCSI-IM-NZP-IMR-res-mem-r16"},
		{Name: "supportedCSI-RS-Density-CMR-r16"},
		{Name: "maxNumberAperiodicCSI-RS-Res-r16"},
		{Name: "supportedSINR-meas-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N8  = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N16 = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N32 = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N64 = 3
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSOneTxCMRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_OneTx_CMR_r16_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N8  = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N16 = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N32 = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N64 = 3
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_r16_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N0  = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N4  = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N8  = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N16 = 3
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N32 = 4
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N64 = 5
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIRS2TxResR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N0, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N4, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSIRS_2Tx_Res_r16_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N8   = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N16  = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N32  = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N64  = 3
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N128 = 4
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSResR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N64, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberSSB_CSIRS_Res_r16_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N8   = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N16  = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N32  = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N64  = 3
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N128 = 4
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResMemR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N64, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_One         = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_Three       = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_OneAndThree = 2
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedCSIRSDensityCMRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_One, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_Three, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedCSI_RS_Density_CMR_r16_OneAndThree},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N2  = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N4  = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N8  = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N16 = 3
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N32 = 4
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N64 = 5
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberAperiodicCSIRSResR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N2, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N4, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N8, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N16, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N32, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_MaxNumberAperiodicCSI_RS_Res_r16_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_SsbWithCSI_IM    = 0
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_SsbWithNZP_IMR   = 1
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_CsirsWithNZP_IMR = 2
	MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_Csi_RSWithoutIMR = 3
)

var mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedSINRMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_SsbWithCSI_IM, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_SsbWithNZP_IMR, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_CsirsWithNZP_IMR, MIMO_ParametersPerBand_Ext_Ssb_Csirs_SINR_Measurement_r16_SupportedSINR_Meas_r16_Csi_RSWithoutIMR},
}

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "overlapPDSCHsFullyFreqTime-r16", Optional: true},
		{Name: "overlapPDSCHsInTimePartiallyFreq-r16", Optional: true},
		{Name: "outOfOrderOperationDL-r16", Optional: true},
		{Name: "outOfOrderOperationUL-r16", Optional: true},
		{Name: "separateCRS-RateMatching-r16", Optional: true},
		{Name: "defaultQCL-PerCORESETPoolIndex-r16", Optional: true},
		{Name: "maxNumberActivatedTCI-States-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OverlapPDSCHsInTimePartiallyFreq_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OverlapPDSCHsInTimePartiallyFreqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OverlapPDSCHsInTimePartiallyFreq_r16_Supported},
}

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportPDCCH-ToPDSCH-r16", Optional: true},
		{Name: "supportPDSCH-ToHARQ-ACK-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationDL_r16_SupportPDCCH_ToPDSCH_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDCCHToPDSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationDL_r16_SupportPDCCH_ToPDSCH_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationDL_r16_SupportPDSCH_ToHARQ_ACK_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDSCHToHARQACKR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationDL_r16_SupportPDSCH_ToHARQ_ACK_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationUL_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationULR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_OutOfOrderOperationUL_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_SeparateCRS_RateMatching_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16SeparateCRSRateMatchingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_SeparateCRS_RateMatching_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_DefaultQCL_PerCORESETPoolIndex_r16_Supported = 0
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16DefaultQCLPerCORESETPoolIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_DefaultQCL_PerCORESETPoolIndex_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N1 = 0
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N2 = 1
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N4 = 2
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N8 = 3
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxNumberPerCORESETPoolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N1, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N2, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N4, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxNumberPerCORESET_Pool_r16_N8},
}

const (
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N2  = 0
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N4  = 1
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N8  = 2
	MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N16 = 3
)

var mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxTotalNumberAcrossCORESETPoolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N2, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N4, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N8, MIMO_ParametersPerBand_Ext_MultiDCI_MultiTRP_Parameters_r16_MaxNumberActivatedTCI_States_r16_MaxTotalNumberAcrossCORESET_Pool_r16_N16},
}

var mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportNewDMRS-Port-r16", Optional: true},
		{Name: "supportTwoPortDL-PTRS-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported1 = 0
	MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported2 = 1
	MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported3 = 2
)

var mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportNewDMRSPortR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported1, MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported2, MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportNewDMRS_Port_r16_Supported3},
}

const (
	MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportTwoPortDL_PTRS_r16_Supported = 0
)

var mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportTwoPortDLPTRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SingleDCI_SDM_Scheme_Parameters_r16_SupportTwoPortDL_PTRS_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N2  = 0
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N3  = 1
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N4  = 2
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N5  = 3
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N6  = 4
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N7  = 5
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N8  = 6
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N16 = 7
)

var mIMOParametersPerBandExtSupportInterSlotTDMR16SupportRepNumPDSCHTDRAR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N2, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N3, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N4, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N5, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N6, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N7, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N8, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_SupportRepNumPDSCH_TDRA_r16_N16},
}

const (
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb3           = 0
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb5           = 1
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb10          = 2
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb20          = 3
	MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_NoRestriction = 4
)

var mIMOParametersPerBandExtSupportInterSlotTDMR16MaxTBSSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb3, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb5, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb10, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_Kb20, MIMO_ParametersPerBand_Ext_SupportInter_SlotTDM_r16_MaxTBS_Size_r16_NoRestriction},
}

var mIMOParametersPerBandExtBeamSwitchTimingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_60kHz_r16_Sym224 = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_60kHz_r16_Sym336 = 1
)

var mIMOParametersPerBandExtBeamSwitchTimingR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_60kHz_r16_Sym224, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_60kHz_r16_Sym336},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_120kHz_r16_Sym224 = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_120kHz_r16_Sym336 = 1
)

var mIMOParametersPerBandExtBeamSwitchTimingR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_120kHz_r16_Sym224, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r16_Scs_120kHz_r16_Sym336},
}

var mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportReportFormat1-2OFDM-syms-r16", Optional: true},
		{Name: "supportReportFormat4-14OFDM-syms-r16", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUCCH_r16_SupportReportFormat1_2OFDM_Syms_r16_Supported = 0
)

var mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat12OFDMSymsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUCCH_r16_SupportReportFormat1_2OFDM_Syms_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUCCH_r16_SupportReportFormat4_14OFDM_Syms_r16_Supported = 0
)

var mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat414OFDMSymsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Semi_PersistentL1_SINR_Report_PUCCH_r16_SupportReportFormat4_14OFDM_Syms_r16_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N96  = 0
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N128 = 1
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N160 = 2
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N192 = 3
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N224 = 4
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N256 = 5
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N288 = 6
	MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N320 = 7
)

var mIMOParametersPerBandExtSpatialRelationsV1640MaxNumberConfiguredSpatialRelationsV1640Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N96, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N128, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N160, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N192, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N224, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N256, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N288, MIMO_ParametersPerBand_Ext_SpatialRelations_v1640_MaxNumberConfiguredSpatialRelations_v1640_N320},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesPerSetPerBWP_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesPerSetPerBWP_r17_N2 = 1
)

var mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesPerSetPerBWPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesPerSetPerBWP_r17_N1, MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesPerSetPerBWP_r17_N2},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N2 = 0
	MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N3 = 1
	MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N4 = 2
)

var mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesAcrossSetsPerBWPR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N2, MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N3, MIMO_ParametersPerBand_Ext_MTRP_BFR_TwoBFD_RS_Set_r17_MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N8   = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N12  = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N16  = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N24  = 3
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N32  = 4
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N48  = 5
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N64  = 6
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N128 = 7
)

var mIMOParametersPerBandExtUnifiedJointTCIR17MaxConfiguredJointTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N12, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N16, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N24, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N32, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N48, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N64, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxConfiguredJointTCI_r17_N128},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N1  = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N2  = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N8  = 3
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N16 = 4
)

var mIMOParametersPerBandExtUnifiedJointTCIR17MaxActivatedTCIAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_r17_MaxActivatedTCIAcrossCC_r17_N16},
}

var mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "minBeamApplicationTime-r17", Optional: true},
		{Name: "maxNumMAC-CE-PerCC"},
	},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N1   = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N2   = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N4   = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N7   = 3
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N14  = 4
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N28  = 5
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N42  = 6
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N56  = 7
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N70  = 8
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N84  = 9
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N98  = 10
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N112 = 11
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N224 = 12
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N336 = 13
)

var mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MinBeamApplicationTimeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N7, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N14, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N28, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N42, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N56, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N70, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N84, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N98, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N112, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N224, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N336},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N2 = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N3 = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N4 = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N5 = 3
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N6 = 4
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N7 = 5
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N8 = 6
)

var mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MaxNumMACCEPerCCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N3, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N4, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N5, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N6, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N7, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MultiMAC_CE_r17_MaxNumMAC_CE_PerCC_N8},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEPerCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_PerCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_InterCell_r17_AdditionalMAC_CE_AcrossCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N4   = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N8   = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N12  = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N16  = 3
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N24  = 4
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N32  = 5
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N48  = 6
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N64  = 7
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N128 = 8
)

var mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredDLTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N12, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N16, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N24, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N32, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N48, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N64, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredDL_TCI_r17_N128},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N4  = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N8  = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N12 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N16 = 3
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N24 = 4
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N32 = 5
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N48 = 6
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N64 = 7
)

var mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredULTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N12, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N16, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N24, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N32, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N48, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxConfiguredUL_TCI_r17_N64},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N1  = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N2  = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N8  = 3
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N16 = 4
)

var mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedDLTCIAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedDL_TCIAcrossCC_r17_N16},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N1  = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N2  = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N8  = 3
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N16 = 4
)

var mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedULTCIAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N8, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_r17_MaxActivatedUL_TCIAcrossCC_r17_N16},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N1   = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N2   = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N4   = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N7   = 3
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N14  = 4
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N28  = 5
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N42  = 6
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N56  = 7
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N70  = 8
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N84  = 9
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N98  = 10
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N112 = 11
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N224 = 12
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N336 = 13
)

var mIMOParametersPerBandExtUnifiedSeparateTCIMultiMACCER17MinBeamApplicationTimeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N7, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N14, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N28, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N42, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N56, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N70, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N84, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N98, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N112, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N224, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_MultiMAC_CE_r17_MinBeamApplicationTime_r17_N336},
}

var mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumListDL-TCI-r17", Optional: true},
		{Name: "maxNumListUL-TCI-r17", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N4 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N8 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListDLTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListDL_TCI_r17_N8},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N4 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N8 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListULTCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_ListSharingCA_r17_MaxNumListUL_TCI_r17_N8},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLPerCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_PerCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULPerCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_PerCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_DL_AcrossCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N0 = 0
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N1 = 1
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N2 = 2
	MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N4 = 3
)

var mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N0, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedSeparateTCI_InterCell_r17_K_UL_AcrossCC_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N4 = 2
	MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N8 = 3
)

var mIMOParametersPerBandExtUnifiedJointTCIMTRPInterCellBMR17MaxNumSSBResourceL1RSRPAcrossCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N1, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N2, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N4, MIMO_ParametersPerBand_Ext_UnifiedJointTCI_MTRP_InterCell_BM_r17_MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17_N8},
}

const (
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N1  = 0
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N2  = 1
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N8  = 3
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N12 = 4
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N16 = 5
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N28 = 6
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N32 = 7
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N48 = 8
	MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N64 = 9
)

var mIMOParametersPerBandExtMpeMitigationR17MaxNumConfRSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N1, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N2, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N4, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N8, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N12, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N16, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N28, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N32, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N48, MIMO_ParametersPerBand_Ext_Mpe_Mitigation_r17_MaxNumConfRS_r17_N64},
}

var mIMOParametersPerBandExtSrsPortReportR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "capVal1-r17", Optional: true},
		{Name: "capVal2-r17", Optional: true},
		{Name: "capVal3-r17", Optional: true},
		{Name: "capVal4-r17", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N4 = 2
)

var mIMOParametersPerBandExtSrsPortReportR17CapVal1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N1, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N2, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal1_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N4 = 2
)

var mIMOParametersPerBandExtSrsPortReportR17CapVal2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N1, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N2, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal2_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N4 = 2
)

var mIMOParametersPerBandExtSrsPortReportR17CapVal3R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N1, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N2, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal3_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N4 = 2
)

var mIMOParametersPerBandExtSrsPortReportR17CapVal4R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N1, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N2, MIMO_ParametersPerBand_Ext_Srs_PortReport_r17_CapVal4_r17_N4},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N2  = 0
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N3  = 1
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N4  = 2
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N8  = 3
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N16 = 4
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N32 = 5
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N64 = 6
)

var mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSWithinSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N2, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N3, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N4, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N8, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N16, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N32, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_WithinSlot_r17_N64},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N8   = 0
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N16  = 1
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N32  = 2
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N64  = 3
	MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N128 = 4
)

var mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSAcrossSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N8, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N16, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N32, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N64, MIMO_ParametersPerBand_Ext_MTRP_GroupBasedL1_RSRP_r17_MaxNumRS_AcrossSlot_r17_N128},
}

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Mode1 = 0
	MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Mode2 = 1
	MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Both  = 2
)

var mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CSIReportModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Mode1, MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Mode2, MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CSI_Report_Mode_r17_Both},
}

var mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17SupportedComboAcrossCCsR17Constraints = per.SizeRange(1, 16)

const (
	MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CodebookModeNCJT_r17_Mode1     = 0
	MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CodebookModeNCJT_r17_Mode1And2 = 1
)

var mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CodebookModeNCJTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CodebookModeNCJT_r17_Mode1, MIMO_ParametersPerBand_Ext_MTRP_CSI_EnhancementPerBand_r17_CodebookModeNCJT_r17_Mode1And2},
}

var mIMOParametersPerBandExtBeamSwitchTimingV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-480kHz", Optional: true},
		{Name: "scs-960kHz", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym56   = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym112  = 1
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym192  = 2
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym896  = 3
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym1344 = 4
)

var mIMOParametersPerBandExtBeamSwitchTimingV1710Scs480kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym56, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym112, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym192, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym896, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_480kHz_Sym1344},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym112  = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym224  = 1
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym384  = 2
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym1792 = 3
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym2688 = 4
)

var mIMOParametersPerBandExtBeamSwitchTimingV1710Scs960kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym112, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym224, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym384, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym1792, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_v1710_Scs_960kHz_Sym2688},
}

var mIMOParametersPerBandExtBeamSwitchTimingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_480kHz_r17_Sym896  = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_480kHz_r17_Sym1344 = 1
)

var mIMOParametersPerBandExtBeamSwitchTimingR17Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_480kHz_r17_Sym896, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_480kHz_r17_Sym1344},
}

const (
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_960kHz_r17_Sym1792 = 0
	MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_960kHz_r17_Sym2688 = 1
)

var mIMOParametersPerBandExtBeamSwitchTimingR17Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_960kHz_r17_Sym1792, MIMO_ParametersPerBand_Ext_BeamSwitchTiming_r17_Scs_960kHz_r17_Sym2688},
}

var mIMOParametersPerBandExtBeamReportTimingV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym56  = 0
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym112 = 1
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym224 = 2
)

var mIMOParametersPerBandExtBeamReportTimingV1710Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym56, MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym112, MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_480kHz_r17_Sym224},
}

const (
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym112 = 0
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym224 = 1
	MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym448 = 2
)

var mIMOParametersPerBandExtBeamReportTimingV1710Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym112, MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym224, MIMO_ParametersPerBand_Ext_BeamReportTiming_v1710_Scs_960kHz_r17_Sym448},
}

var mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-480kHz-r17", Optional: true},
		{Name: "scs-960kHz-r17", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N2 = 0
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N4 = 1
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N7 = 2
)

var mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs480kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N2, MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N4, MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_480kHz_r17_N7},
}

const (
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N1 = 0
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N2 = 1
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N4 = 2
	MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N7 = 3
)

var mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs960kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N1, MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N2, MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N4, MIMO_ParametersPerBand_Ext_MaxNumberRxTxBeamSwitchDL_v1710_Scs_960kHz_r17_N7},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N8   = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N12  = 1
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N16  = 2
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N24  = 3
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N32  = 4
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N48  = 5
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N64  = 6
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N128 = 7
)

var mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N64, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N6  = 2
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N16 = 4
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N32 = 5
)

var mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberActiveJointTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N6, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N16, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumberActiveJointTCI_AcrossCC_r18_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18_Tci_StateInd_r18_WithAssignment    = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18_Tci_StateInd_r18_WithoutAssignment = 1
)

var mIMOParametersPerBandExtTciJointTCIUpdateMultiActiveTCIPerCCR18TciStateIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18_Tci_StateInd_r18_WithAssignment, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18_Tci_StateInd_r18_WithoutAssignment},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N4   = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N8   = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N12  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N16  = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N24  = 4
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N32  = 5
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N48  = 6
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N64  = 7
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N128 = 8
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N64, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N4  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N8  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N12 = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N16 = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N24 = 4
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N32 = 5
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N48 = 6
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N64 = 7
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigULTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16 = 3
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16 = 3
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16 = 3
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16 = 3
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCell             = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCellAndInterCell = 1
)

var mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCell, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCellAndInterCell},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N8   = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N12  = 1
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N16  = 2
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N24  = 3
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N32  = 4
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N48  = 5
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N64  = 6
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N128 = 7
)

var mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N64, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberConfigJointTCIPerCC_PerBWP_r18_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N1  = 0
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N2  = 1
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N4  = 2
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N16 = 4
)

var mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberActiveJointTCIAcrossCCPerCORESETR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N1, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N2, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N4, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N8, MIMO_ParametersPerBand_Ext_Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCell             = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCellAndInterCell = 1
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCell, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MTRP_Operation_r18_IntraCellAndInterCell},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N8   = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N12  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N16  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N24  = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N32  = 4
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N48  = 5
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N64  = 6
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N128 = 7
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N64, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigDL_TCI_PerCC_PerBWP_r18_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N8  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N12 = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N16 = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N24 = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N32 = 4
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N48 = 5
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N64 = 6
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigULTCIPerCCPerBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N12, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N16, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N24, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N32, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N48, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumConfigUL_TCI_PerCC_PerBWP_r18_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N1  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16 = 4
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveDLTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N1, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveDL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N1  = 0
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2  = 1
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4  = 2
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16 = 4
)

var mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveULTCIAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N1, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18_MaxNumActiveUL_TCI_AcrossCC_r18_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N6  = 2
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N10 = 4
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N12 = 5
)

var mIMOParametersPerBandExtTdcpResourceR18MaxNumberConfigPerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N2, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N4, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N6, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N8, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N10, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberConfigPerCC_r18_N12},
}

const (
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N4  = 1
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N6  = 2
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N12 = 4
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N16 = 5
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N20 = 6
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N24 = 7
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N28 = 8
	MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N32 = 9
)

var mIMOParametersPerBandExtTdcpResourceR18MaxNumberSimultaneousPerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N2, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N4, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N6, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N8, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N12, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N16, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N20, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N24, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N28, MIMO_ParametersPerBand_Ext_Tdcp_Resource_r18_MaxNumberSimultaneousPerCC_r18_N32},
}

const (
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_JointULandDL = 0
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_UlOnly       = 1
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_Both         = 2
)

var mIMOParametersPerBandExtGroupBeamReportingSTx2PR18GroupL1RSRPReportingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_JointULandDL, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_UlOnly, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_GroupL1_RSRP_Reporting_r18_Both},
}

const (
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N2  = 0
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N3  = 1
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N4  = 2
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N8  = 3
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N16 = 4
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N32 = 5
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N64 = 6
)

var mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResWithinSlotAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N2, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N3, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N4, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N16, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N32, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResWithinSlotAcrossCC_r18_N64},
}

const (
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N8   = 0
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N16  = 1
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N32  = 2
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N64  = 3
	MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N128 = 4
)

var mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResAcrossCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N8, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N16, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N32, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N64, MIMO_ParametersPerBand_Ext_GroupBeamReporting_STx2P_r18_MaxNumberResAcrossCC_r18_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N3  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N4  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N8  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N10 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N12 = 6
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N16 = 7
)

var mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N3, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N10, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N12, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N16 = 2
)

var mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetBR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetB_r19_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N8  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N16 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N32 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N64 = 3
)

var mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetAR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_MaxNumberOfResourceSetA_r19_N64},
}

var mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodic-r19", Optional: true},
		{Name: "aperiodic-r19", Optional: true},
		{Name: "semiPersistent-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_Periodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_Periodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_Aperiodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19AperiodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_Aperiodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_SemiPersistent_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_ResourceTypeSetB_CSI_RS_r19_SemiPersistent_r19_Supported},
}

var mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodic-r19", Optional: true},
		{Name: "aperiodic-r19", Optional: true},
		{Name: "semiPersistent-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_Periodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_Periodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_Aperiodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19AperiodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_Aperiodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_SemiPersistent_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19SemiPersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_InferenceReportType_r19_SemiPersistent_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_Subset  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_DiffSet = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_Both    = 2
)

var mIMOParametersPerBandExtAimlBMCase1R19SubUseCasesR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_Subset, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_DiffSet, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_SubUseCases_r19_Both},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N7  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N14 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N21 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N28 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N35 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N42 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N56 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N7, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N21, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N35, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N56},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N14  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N28  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N42  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N56  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N70  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N84  = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N112 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N70, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N112},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N28  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N56  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N84  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N112 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N140 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N168 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N224 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N140, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N224},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N168 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N224 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N280 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N336 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N448 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N280, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N224  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N448  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N672  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N896  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1120 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1344 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1792 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N672, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1120, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N448  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N896  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1344 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1792 = 3
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N7  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N14 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N21 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N28 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N35 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N42 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N56 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N7, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N21, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N35, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N56},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N14  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N28  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N42  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N56  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N70  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N84  = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N112 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N70, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N112},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N28  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N56  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N84  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N112 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N140 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N168 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N224 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N140, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N224},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N168 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N224 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N280 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N336 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N448 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N280, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N224  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N448  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N672  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N896  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1120 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1344 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1792 = 6
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N672, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1120, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N448  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N896  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1344 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1792 = 3
)

var mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case1_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N3  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N4  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N8  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N10 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N12 = 6
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N16 = 7
)

var mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N3, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N10, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N12, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfInferenceReportAcrossAllCC_r19_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N16 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N32 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N64 = 4
)

var mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetBR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetB_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N16 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N32 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N64 = 4
)

var mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetAR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfResourceSetA_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N2 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N4 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N8 = 2
)

var mIMOParametersPerBandExtAimlBMCase2R19MinNumberOfKBMSetBR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MinNumberOfKBM_SetB_r19_N8},
}

var mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodic-r19", Optional: true},
		{Name: "semiPersistent-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_ResourceTypeSetB_CSI_RS_r19_Periodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_ResourceTypeSetB_CSI_RS_r19_Periodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_ResourceTypeSetB_CSI_RS_r19_SemiPersistent_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_ResourceTypeSetB_CSI_RS_r19_SemiPersistent_r19_Supported},
}

var mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "periodic-r19", Optional: true},
		{Name: "aperiodic-r19", Optional: true},
		{Name: "semiPersistent-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_Periodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19PeriodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_Periodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_Aperiodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19AperiodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_Aperiodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_SemiPersistent_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19SemiPersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_InferenceReportType_r19_SemiPersistent_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N1 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N2 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N4 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N8 = 3
)

var mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfPredictedTimeInstanceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxNumberOfPredictedTimeInstance_r19_N8},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N4  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N6  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N8  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N12 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N16 = 6
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N32 = 7
)

var mIMOParametersPerBandExtAimlBMCase2R19MaxTotalNumberOfPredictedBeamPerReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N6, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N12, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_MaxTotalNumberOfPredictedBeamPerReport_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms10  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms20  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms40  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms80  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms160 = 4
)

var mIMOParametersPerBandExtAimlBMCase2R19DummyTimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms10, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms20, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms40, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms80, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_Dummy_TimeGap_r19_Ms160},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N14 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N28 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N42 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N56 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N70 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N84 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N98 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N70, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs15kHz_r19_N98},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N28  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N56  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N84  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N112 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N140 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N168 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N196 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N140, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs30kHz_r19_N196},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N168 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N224 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N280 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N336 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N392 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N280, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs60kHz_r19_N392},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N112 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N224 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N336 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N448 = 3
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs120kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N448  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N896  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1344 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1792 = 3
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs480kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N896  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1792 = 1
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD_r19_Scs960kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N14 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N28 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N42 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N56 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N70 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N84 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N98 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N42, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N70, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs15kHz_r19_N98},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N28  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N56  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N84  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N112 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N140 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N168 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N196 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N84, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N140, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs30kHz_r19_N196},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N168 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N224 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N280 = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N336 = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N392 = 6
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N168, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N280, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs60kHz_r19_N392},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N112 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N224 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N336 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N448 = 3
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N336, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs120kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N448  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N896  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1344 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1792 = 3
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1344, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs480kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N896  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1792 = 1
)

var mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_r19_RelaxationTimelineD1_r19_Scs960kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N1 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N2 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N4 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N8 = 3
)

var mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxNumPredictedTimeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxNumPredictedTime_r19_N8},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N3  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N4  = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N6  = 4
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N8  = 5
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N12 = 6
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N16 = 7
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N24 = 8
	MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N32 = 9
)

var mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxTotalNumPredictedBeamInOneReportR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N3, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N6, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N12, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N24, MIMO_ParametersPerBand_Ext_Aiml_BM_Case2_PredictedRSRP_r19_MaxTotalNumPredictedBeamInOneReport_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N16 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N32 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N64 = 4
)

var mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumTotalResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumTotalResource_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N1 = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N2 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N4 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N8 = 3
)

var mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumReportAcrossAllCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumReportAcrossAllCC_r19_N8},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N3  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N7  = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N15 = 3
)

var mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumOccasionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N3, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N7, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MaxNumOccasion_r19_N15},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringResourceType_r19_Periodic       = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringResourceType_r19_Semipersistent = 1
)

var mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringResourceTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringResourceType_r19_Periodic, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringResourceType_r19_Semipersistent},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Periodic       = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Aperiodic      = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Semipersistent = 2
)

var mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringReportTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Periodic, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Aperiodic, MIMO_ParametersPerBand_Ext_Aiml_BM_Monitoring_r19_MonitoringReportType_r19_Semipersistent},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_Equal     = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_Subset    = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_NotSubset = 2
)

var mIMOParametersPerBandExtAimlBMUEDataCollectionR19SubCaseR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_Equal, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_Subset, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_SubCase_r19_NotSubset},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N16 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N32 = 3
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N64 = 4
)

var mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetBR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N4, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetB_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N8  = 0
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N16 = 1
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N32 = 2
	MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N64 = 3
)

var mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetAR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N8, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N16, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N32, MIMO_ParametersPerBand_Ext_Aiml_BM_UE_DataCollection_r19_MaxNumResourceSetA_r19_N64},
}

var mIMOParametersPerBandExtAimlCSIPredictionR19SupportedCSIRSResourceListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N1 = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N2 = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N4 = 2
)

var mIMOParametersPerBandExtAimlCSIPredictionR19ScalingFactorR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N1, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N2, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_ScalingFactor_r19_N4},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N14  = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N28  = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N56  = 2
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N112 = 3
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N14, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs15kHz_r19_N112},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N28  = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N56  = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N112 = 2
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N224 = 3
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N28, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs30kHz_r19_N224},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N224 = 2
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N448 = 3
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N56, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs60kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N112 = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N224 = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N448 = 2
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N112, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N224, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs120kHz_r19_N448},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N448  = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N896  = 1
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N1792 = 2
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N448, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs480kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs960kHz_r19_N896  = 0
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs960kHz_r19_N1792 = 1
)

var mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs960kHz_r19_N896, MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_RelaxationTimelineT_r19_Scs960kHz_r19_N1792},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_InferenceReportType_r19_Aperiodic_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19AperiodicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_InferenceReportType_r19_Aperiodic_r19_Supported},
}

const (
	MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_InferenceReportType_r19_SemiPersistent_r19_Supported = 0
)

var mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19SemiPersistentR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Aiml_CSI_Prediction_r19_InferenceReportType_r19_SemiPersistent_r19_Supported},
}

var mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossAllCCR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossOneReportR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var mIMOParametersPerBandExtAimlCSIPredictionMonitoringR19SupportedCSIRSResourceListR19Constraints = per.SizeRange(1, common.MaxNrofCSI_RS_ResourcesExt_r16)

var mIMOParametersPerBandExtUeiModeBR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs15kHz-r19", Optional: true},
		{Name: "scs30kHz-r19", Optional: true},
		{Name: "scs60kHz-r19", Optional: true},
		{Name: "scs120kHz-r19", Optional: true},
		{Name: "scs480kHz-r19", Optional: true},
		{Name: "scs960kHz-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N0  = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N1  = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N2  = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N4  = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N8  = 4
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N16 = 5
)

var mIMOParametersPerBandExtUeiModeBR19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N1, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N2, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N4, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N8, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs15kHz_r19_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N0  = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N4  = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N16 = 4
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N32 = 5
)

var mIMOParametersPerBandExtUeiModeBR19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N2, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N4, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N8, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N16, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs30kHz_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N0  = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N8  = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N32 = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N64 = 4
)

var mIMOParametersPerBandExtUeiModeBR19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N4, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N8, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N32, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs60kHz_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N0   = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N8   = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N16  = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N32  = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N64  = 4
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N128 = 5
)

var mIMOParametersPerBandExtUeiModeBR19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N8, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N16, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N32, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N64, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs120kHz_r19_N128},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N0   = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N32  = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N64  = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N128 = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N256 = 4
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N512 = 5
)

var mIMOParametersPerBandExtUeiModeBR19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N32, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N64, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N128, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N256, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs480kHz_r19_N512},
}

const (
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N0   = 0
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N64  = 1
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N128 = 2
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N256 = 3
	MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N512 = 4
)

var mIMOParametersPerBandExtUeiModeBR19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N0, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N64, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N128, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N256, MIMO_ParametersPerBand_Ext_Uei_ModeB_r19_Scs960kHz_r19_N512},
}

const (
	MIMO_ParametersPerBand_Ext_Nes_SD_Type1_SP_r19_Timeline_r19_Cap1 = 0
	MIMO_ParametersPerBand_Ext_Nes_SD_Type1_SP_r19_Timeline_r19_Cap2 = 1
)

var mIMOParametersPerBandExtNesSDType1SPR19TimelineR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Nes_SD_Type1_SP_r19_Timeline_r19_Cap1, MIMO_ParametersPerBand_Ext_Nes_SD_Type1_SP_r19_Timeline_r19_Cap2},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MinRangeDdInCyclicPrefix_r19_Half = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MinRangeDdInCyclicPrefix_r19_Full = 1
)

var mIMOParametersPerBandExtCjtcDdReportR19MinRangeDdInCyclicPrefixR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MinRangeDdInCyclicPrefix_r19_Half, MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MinRangeDdInCyclicPrefix_r19_Full},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N32  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N64  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N128 = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N256 = 3
)

var mIMOParametersPerBandExtCjtcDdReportR19MaxResolutionDdR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N64, MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N128, MIMO_ParametersPerBand_Ext_Cjtc_DdReport_r19_MaxResolutionDd_r19_N256},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N10 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N12 = 5
)

var mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N10, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_Resource_r19_N12},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64 = 5
)

var mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32 = 9
)

var mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32 = 9
	MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64 = 10
)

var mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_DdReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MinRangeFO_r19_PpmDot1 = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MinRangeFO_r19_PpmDot2 = 1
)

var mIMOParametersPerBandExtCjtcFOReportR19MinRangeFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MinRangeFO_r19_PpmDot1, MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MinRangeFO_r19_PpmDot2},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N16  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N32  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N256 = 2
)

var mIMOParametersPerBandExtCjtcFOReportR19MaxResolutionFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_FO_Report_r19_MaxResolutionFO_r19_N256},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N10 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N12 = 5
)

var mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N10, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N12},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64 = 5
)

var mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32 = 9
)

var mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32 = 9
	MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64 = 10
)

var mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_FO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWideband_r19_MaxResolution_r19_N16 = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWideband_r19_MaxResolution_r19_N32 = 1
)

var mIMOParametersPerBandExtCjtcPOReportWidebandR19MaxResolutionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWideband_r19_MaxResolution_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWideband_r19_MaxResolution_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N10 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N12 = 5
)

var mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N10, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_Configured_r19_N12},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N64 = 5
)

var mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ConfiguredAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32 = 9
)

var mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32 = 9
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64 = 10
)

var mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportWidebandProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MaxResolution_r19_N16 = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MaxResolution_r19_N32 = 1
)

var mIMOParametersPerBandExtCjtcPOReportSubbandR19MaxResolutionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MaxResolution_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MaxResolution_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N1  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N2  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N4  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N16 = 4
)

var mIMOParametersPerBandExtCjtcPOReportSubbandR19MinSubbandSizeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N1, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_PO_ReportSubband_r19_MinSubbandSize_r19_N16},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeDdInCyclicPrefix_r19_Half = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeDdInCyclicPrefix_r19_Full = 1
)

var mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeDdInCyclicPrefixR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeDdInCyclicPrefix_r19_Half, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeDdInCyclicPrefix_r19_Full},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N32  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N64  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N128 = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N256 = 3
)

var mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionDdR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N64, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N128, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionDd_r19_N256},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeFO_r19_PpmDot1 = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeFO_r19_PpmDot2 = 1
)

var mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeFO_r19_PpmDot1, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MinRangeFO_r19_PpmDot2},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N16  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N32  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N256 = 2
)

var mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionFOR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_Report_r19_MaxResolutionFO_r19_N256},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N10 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N12 = 5
)

var mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N10, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_Resource_r19_N12},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64 = 5
)

var mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberTRS_ResourceAcrossCC_r19_N64},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32 = 9
)

var mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourcePerCC_r19_N32},
}

const (
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2  = 0
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4  = 1
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6  = 2
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8  = 3
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12 = 4
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16 = 5
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20 = 6
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24 = 7
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28 = 8
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32 = 9
	MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64 = 10
)

var mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N2, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N4, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N6, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N8, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N12, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N16, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N20, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N24, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N28, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N32, MIMO_ParametersPerBand_Ext_Cjtc_DdFO_ReportProcessing_r19_MaxNumberCSI_RS_ResourceAcrossCC_r19_N64},
}

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs15kHz-r19", Optional: true},
		{Name: "scs30kHz-r19", Optional: true},
		{Name: "scs60kHz-r19", Optional: true},
		{Name: "scs120kHz-r19", Optional: true},
		{Name: "scs480kHz-r19", Optional: true},
		{Name: "scs960kHz-r19", Optional: true},
	},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N2 = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N4 = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N8 = 2
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N2, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N4, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs15kHz_r19_N8},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N4  = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N8  = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N14 = 2
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N28 = 3
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N4, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N8, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N14, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs30kHz_r19_N28},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N8  = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N14 = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N28 = 2
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N8, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N14, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs60kHz_r19_N28},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N14 = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N28 = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N56 = 2
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N14, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N28, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs120kHz_r19_N56},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N56  = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N112 = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N224 = 2
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs480kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N56, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N112, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs480kHz_r19_N224},
}

const (
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N112 = 0
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N224 = 1
	MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N448 = 2
)

var mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs960kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N112, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N224, MIMO_ParametersPerBand_Ext_TimelineRelax_CJTC_Dd_EType2CJT_r19_Scs960kHz_r19_N448},
}

type MIMO_ParametersPerBand struct {
	Tci_StatePDSCH *struct {
		MaxNumberConfiguredTCI_StatesPerCC *int64
		MaxNumberActiveTCI_PerBWP          *int64
	}
	AdditionalActiveTCI_StatePDCCH           *int64
	Pusch_TransCoherence                     *int64
	BeamCorrespondenceWithoutUL_BeamSweeping *int64
	PeriodicBeamReport                       *int64
	AperiodicBeamReport                      *int64
	Sp_BeamReportPUCCH                       *int64
	Sp_BeamReportPUSCH                       *int64
	Dummy1                                   *DummyG
	MaxNumberRxBeam                          *int64
	MaxNumberRxTxBeamSwitchDL                *struct {
		Scs_15kHz  *int64
		Scs_30kHz  *int64
		Scs_60kHz  *int64
		Scs_120kHz *int64
		Scs_240kHz *int64
	}
	MaxNumberNonGroupBeamReporting *int64
	GroupBeamReporting             *int64
	UplinkBeamManagement           *struct {
		MaxNumberSRS_ResourcePerSet_BM int64
		MaxNumberSRS_ResourceSet       int64
	}
	MaxNumberCSI_RS_BFD     *int64
	MaxNumberSSB_BFD        *int64
	MaxNumberCSI_RS_SSB_CBD *int64
	Dummy2                  *int64
	TwoPortsPTRS_UL         *int64
	Dummy5                  *SRS_Resources
	Dummy3                  *int64
	BeamReportTiming        *struct {
		Scs_15kHz  *int64
		Scs_30kHz  *int64
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	Ptrs_DensityRecommendationSetDL *struct {
		Scs_15kHz  *PTRS_DensityRecommendationDL
		Scs_30kHz  *PTRS_DensityRecommendationDL
		Scs_60kHz  *PTRS_DensityRecommendationDL
		Scs_120kHz *PTRS_DensityRecommendationDL
	}
	Ptrs_DensityRecommendationSetUL *struct {
		Scs_15kHz  *PTRS_DensityRecommendationUL
		Scs_30kHz  *PTRS_DensityRecommendationUL
		Scs_60kHz  *PTRS_DensityRecommendationUL
		Scs_120kHz *PTRS_DensityRecommendationUL
	}
	Dummy4                   *DummyH
	AperiodicTRS             *int64
	Dummy6                   *int64
	BeamManagementSSB_CSI_RS *BeamManagementSSB_CSI_RS
	BeamSwitchTiming         *struct {
		Scs_60kHz  *int64
		Scs_120kHz *int64
	}
	CodebookParameters                           *CodebookParameters
	Csi_RS_IM_ReceptionForFeedback               *CSI_RS_IM_ReceptionForFeedback
	Csi_RS_ProcFrameworkForSRS                   *CSI_RS_ProcFrameworkForSRS
	Csi_ReportFramework                          *CSI_ReportFramework
	Csi_RS_ForTracking                           *CSI_RS_ForTracking
	Srs_AssocCSI_RS                              []SupportedCSI_RS_Resource
	SpatialRelations                             *SpatialRelations
	DefaultQCL_TwoTCI_r16                        *int64
	CodebookParametersPerBand_r16                *CodebookParameters_v1610
	Simul_SpatialRelationUpdatePUCCHResGroup_r16 *int64
	MaxNumberSCellBFR_r16                        *int64
	SimultaneousReceptionDiffTypeD_r16           *int64
	Ssb_Csirs_SINR_Measurement_r16               *struct {
		MaxNumberSSB_CSIRS_OneTx_CMR_r16    int64
		MaxNumberCSI_IM_NZP_IMR_Res_r16     int64
		MaxNumberCSIRS_2Tx_Res_r16          int64
		MaxNumberSSB_CSIRS_Res_r16          int64
		MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16 int64
		SupportedCSI_RS_Density_CMR_r16     int64
		MaxNumberAperiodicCSI_RS_Res_r16    int64
		SupportedSINR_Meas_r16              *int64
	}
	NonGroupSINR_Reporting_r16       *int64
	GroupSINR_Reporting_r16          *int64
	MultiDCI_MultiTRP_Parameters_r16 *struct {
		OverlapPDSCHsFullyFreqTime_r16       *int64
		OverlapPDSCHsInTimePartiallyFreq_r16 *int64
		OutOfOrderOperationDL_r16            *struct {
			SupportPDCCH_ToPDSCH_r16    *int64
			SupportPDSCH_ToHARQ_ACK_r16 *int64
		}
		OutOfOrderOperationUL_r16          *int64
		SeparateCRS_RateMatching_r16       *int64
		DefaultQCL_PerCORESETPoolIndex_r16 *int64
		MaxNumberActivatedTCI_States_r16   *struct {
			MaxNumberPerCORESET_Pool_r16         int64
			MaxTotalNumberAcrossCORESET_Pool_r16 int64
		}
	}
	SingleDCI_SDM_Scheme_Parameters_r16 *struct {
		SupportNewDMRS_Port_r16   *int64
		SupportTwoPortDL_PTRS_r16 *int64
	}
	SupportFDM_SchemeA_r16           *int64
	SupportCodeWordSoftCombining_r16 *int64
	SupportTDM_SchemeA_r16           *int64
	SupportInter_SlotTDM_r16         *struct {
		SupportRepNumPDSCH_TDRA_r16 int64
		MaxTBS_Size_r16             int64
		MaxNumberTCI_States_r16     int64
	}
	LowPAPR_DMRS_PDSCH_r16                 *int64
	LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 *int64
	LowPAPR_DMRS_PUCCH_r16                 *int64
	LowPAPR_DMRS_PUSCHwithPrecoding_r16    *int64
	Csi_ReportFrameworkExt_r16             *CSI_ReportFrameworkExt_r16
	CodebookParametersAddition_r16         *CodebookParametersAddition_r16
	CodebookComboParametersAddition_r16    *CodebookComboParametersAddition_r16
	BeamCorrespondenceSSB_Based_r16        *int64
	BeamCorrespondenceCSI_RS_Based_r16     *int64
	BeamSwitchTiming_r16                   *struct {
		Scs_60kHz_r16  *int64
		Scs_120kHz_r16 *int64
	}
	Semi_PersistentL1_SINR_Report_PUCCH_r16 *struct {
		SupportReportFormat1_2OFDM_Syms_r16  *int64
		SupportReportFormat4_14OFDM_Syms_r16 *int64
	}
	Semi_PersistentL1_SINR_Report_PUSCH_r16 *int64
	SpatialRelations_v1640                  *struct{ MaxNumberConfiguredSpatialRelations_v1640 int64 }
	Support64CandidateBeamRS_BFR_r16        *int64
	MaxMIMO_LayersForMulti_DCI_MTRP_r16     *int64
	SupportedSINR_Meas_v1670                *per.BitString
	Srs_IncreasedRepetition_r17             *int64
	Srs_PartialFrequencySounding_r17        *int64
	Srs_StartRB_LocationHoppingPartial_r17  *int64
	Srs_CombEight_r17                       *int64
	CodebookParametersfetype2_r17           *CodebookParametersfetype2_r17
	MTRP_PUSCH_TwoCSI_RS_r17                *int64
	MTRP_PUCCH_InterSlot_r17                *int64
	MTRP_PUCCH_CyclicMapping_r17            *int64
	MTRP_PUCCH_SecondTPC_r17                *int64
	MTRP_BFR_TwoBFD_RS_Set_r17              *struct {
		MaxBFD_RS_ResourcesPerSetPerBWP_r17     int64
		MaxBFR_r17                              int64
		MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17 int64
	}
	MTRP_BFR_PUCCH_SR_PerCG_r17         *int64
	MTRP_BFR_Association_PUCCH_SR_r17   *int64
	Sfn_SimulTwoTCI_AcrossMultiCC_r17   *int64
	Sfn_DefaultDL_BeamSetup_r17         *int64
	Sfn_DefaultUL_BeamSetup_r17         *int64
	Srs_TriggeringOffset_r17            *int64
	Srs_TriggeringDCI_r17               *int64
	CodebookComboParameterMixedType_r17 *CodebookComboParameterMixedType_r17
	UnifiedJointTCI_r17                 *struct {
		MaxConfiguredJointTCI_r17   int64
		MaxActivatedTCIAcrossCC_r17 int64
	}
	UnifiedJointTCI_MultiMAC_CE_r17 *struct {
		MinBeamApplicationTime_r17 *int64
		MaxNumMAC_CE_PerCC         int64
	}
	UnifiedJointTCI_PerBWP_CA_r17       *int64
	UnifiedJointTCI_ListSharingCA_r17   *int64
	UnifiedJointTCI_CommonMultiCC_r17   *int64
	UnifiedJointTCI_BeamAlignDLRS_r17   *int64
	UnifiedJointTCI_PC_Association_r17  *int64
	UnifiedJointTCI_Legacy_r17          *int64
	UnifiedJointTCI_Legacy_SRS_r17      *int64
	UnifiedJointTCI_Legacy_CORESET0_r17 *int64
	UnifiedJointTCI_SCellBFR_r17        *int64
	UnifiedJointTCI_InterCell_r17       *struct {
		AdditionalMAC_CE_PerCC_r17    int64
		AdditionalMAC_CE_AcrossCC_r17 int64
	}
	UnifiedSeparateTCI_r17 *struct {
		MaxConfiguredDL_TCI_r17        int64
		MaxConfiguredUL_TCI_r17        int64
		MaxActivatedDL_TCIAcrossCC_r17 int64
		MaxActivatedUL_TCIAcrossCC_r17 int64
	}
	UnifiedSeparateTCI_MultiMAC_CE_r17 *struct {
		MinBeamApplicationTime_r17  int64
		MaxActivatedDL_TCIPerCC_r17 int64
		MaxActivatedUL_TCIPerCC_r17 int64
	}
	UnifiedSeparateTCI_PerBWP_CA_r17     *int64
	UnifiedSeparateTCI_ListSharingCA_r17 *struct {
		MaxNumListDL_TCI_r17 *int64
		MaxNumListUL_TCI_r17 *int64
	}
	UnifiedSeparateTCI_CommonMultiCC_r17 *int64
	UnifiedSeparateTCI_InterCell_r17     *struct {
		K_DL_PerCC_r17    int64
		K_UL_PerCC_r17    int64
		K_DL_AcrossCC_r17 int64
		K_UL_AcrossCC_r17 int64
	}
	UnifiedJointTCI_MTRP_InterCell_BM_r17 *struct {
		MaxNumAdditionalPCI_L1_RSRP_r17        int64
		MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17 int64
	}
	Mpe_Mitigation_r17 *struct {
		MaxNumP_MPR_RI_Pairs_r17 int64
		MaxNumConfRS_r17         int64
	}
	Srs_PortReport_r17 *struct {
		CapVal1_r17 *int64
		CapVal2_r17 *int64
		CapVal3_r17 *int64
		CapVal4_r17 *int64
	}
	MTRP_PDCCH_Individual_r17       *int64
	MTRP_PDCCH_AnySpan_3Symbols_r17 *int64
	MTRP_PDCCH_TwoQCL_TypeD_r17     *int64
	MTRP_PUSCH_CSI_RS_r17           *struct {
		MaxNumPeriodicSRS_r17          int64
		MaxNumAperiodicSRS_r17         int64
		MaxNumSP_SRS_r17               int64
		NumSRS_ResourcePerCC_r17       int64
		NumSRS_ResourceNonCodebook_r17 int64
	}
	MTRP_PUSCH_CyclicMapping_r17    *int64
	MTRP_PUSCH_SecondTPC_r17        *int64
	MTRP_PUSCH_TwoPHR_Reporting_r17 *int64
	MTRP_PUSCH_A_CSI_r17            *int64
	MTRP_PUSCH_SP_CSI_r17           *int64
	MTRP_PUSCH_CG_r17               *int64
	MTRP_PUCCH_MAC_CE_r17           *int64
	MTRP_PUCCH_MaxNum_PC_FR1_r17    *int64
	MTRP_Inter_Cell_r17             *struct {
		MaxNumAdditionalPCI_Case1_r17 int64
		MaxNumAdditionalPCI_Case2_r17 int64
	}
	MTRP_GroupBasedL1_RSRP_r17 *struct {
		MaxNumBeamGroups_r17    int64
		MaxNumRS_WithinSlot_r17 int64
		MaxNumRS_AcrossSlot_r17 int64
	}
	MTRP_BFD_RS_MAC_CE_r17          *int64
	MTRP_CSI_EnhancementPerBand_r17 *struct {
		MaxNumNZP_CSI_RS_r17        int64
		CSI_Report_Mode_r17         int64
		SupportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17
		CodebookModeNCJT_r17        int64
	}
	CodebookComboParameterMultiTRP_r17 *CodebookComboParameterMultiTRP_r17
	MTRP_CSI_AdditionalCSI_r17         *int64
	MTRP_CSI_N_Max2_r17                *int64
	MTRP_CSI_CMR_r17                   *int64
	Srs_PartialFreqSounding_r17        *int64
	BeamSwitchTiming_v1710             *struct {
		Scs_480kHz *int64
		Scs_960kHz *int64
	}
	BeamSwitchTiming_r17 *struct {
		Scs_480kHz_r17 *int64
		Scs_960kHz_r17 *int64
	}
	BeamReportTiming_v1710 *struct {
		Scs_480kHz_r17 *int64
		Scs_960kHz_r17 *int64
	}
	MaxNumberRxTxBeamSwitchDL_v1710 *struct {
		Scs_480kHz_r17 *int64
		Scs_960kHz_r17 *int64
	}
	Srs_PortReportSP_AP_r17                      *int64
	MaxNumberRxBeam_v1720                        *int64
	Sfn_ImplicitRS_TwoTCI_r17                    *int64
	Sfn_QCL_TypeD_Collision_TwoTCI_r17           *int64
	MTRP_CSI_NumCPU_r17                          *int64
	SupportRepNumPDSCH_TDRA_DCI_1_2_r17          *int64
	CodebookParametersetype2DopplerCSI_r18       *CodebookParametersetype2DopplerCSI_r18
	CodebookParametersfetype2DopplerCSI_r18      *CodebookParametersfetype2DopplerCSI_r18
	CodebookParametersetype2CJT_r18              *CodebookParametersetype2CJT_r18
	CodebookParametersfetype2CJT_r18             *CodebookParametersfetype2CJT_r18
	CodebookComboParametersCJT_r18               *CodebookComboParametersCJT_r18
	CodebookParametersHARQ_ACK_PUSCH_r18         *CodebookParametersHARQ_ACK_PUSCH_r18
	Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18 *struct {
		MaxNumberConfigJointTCIPerCC_PerBWP_r18 int64
		MaxNumberActiveJointTCI_AcrossCC_r18    int64
	}
	Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18 *struct {
		Tci_StateInd_r18                  int64
		MaxNumberActiveJointTCI_PerCC_r18 int64
	}
	Tci_SelectionDCI_r18                            *int64
	Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18 *struct {
		MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
		MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
		MaxNumActiveDL_TCI_AcrossCC_r18     int64
		MaxNumActiveUL_TCI_AcrossCC_r18     int64
	}
	Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18 *struct {
		MaxNumActiveDL_TCI_AcrossCC_r18 int64
		MaxNumActiveUL_TCI_AcrossCC_r18 int64
	}
	Tci_SelectionAperiodicCSI_RS_r18                        *int64
	Tci_SelectionAperiodicCSI_RS_M_DCI_r18                  *int64
	TwoTCI_StatePDSCH_CJT_TxScheme_r18                      *int64
	Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 *struct {
		MTRP_Operation_r18                             int64
		MaxNumberConfigJointTCIPerCC_PerBWP_r18        int64
		MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18 int64
	}
	Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18     *int64
	Tci_TRP_BFR_r18                                            *int64
	Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 *struct {
		MTRP_Operation_r18                  int64
		MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
		MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
		MaxNumActiveDL_TCI_AcrossCC_r18     int64
		MaxNumActiveUL_TCI_AcrossCC_r18     int64
	}
	Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 *struct {
		MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
		MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
	}
	CommonTCI_SingleDCI_r18               *int64
	CommonTCI_MultiDCI_r18                *int64
	TwoPHR_Reporting_r18                  *int64
	SpCell_TAG_Ind_r18                    *int64
	InterCellCrossTRP_PDCCH_OrderCFRA_r18 *int64
	IntraCellCrossTRP_PDCCH_OrderCFRA_r18 *int64
	OverlapUL_TransReduction_r18          *int64
	MaxPeriodicityCMR_r18                 *int64
	Tdcp_Report_r18                       *struct {
		ValueX_r18                  int64
		MaxNumberActiveResource_r18 int64
	}
	Tdcp_Resource_r18 *struct {
		MaxNumberConfigPerCC_r18       int64
		MaxNumberConfigAcrossCC_r18    int64
		MaxNumberSimultaneousPerCC_r18 int64
	}
	TimelineRelax_CJT_CSI_r18                      *int64
	JointConfigDMRSPortDynamicSwitching_r18        *int64
	Srs_CombOffsetHopping_r18                      *int64
	Srs_CombOffsetInTime_r18                       *int64
	Srs_CombOffsetCombinedGroupSequence_r18        *int64
	Srs_CombOffsetHoppingWithinSubset_r18          *int64
	Srs_CyclicShiftHopping_r18                     *int64
	Srs_CyclicShiftHoppingSmallGranularity_r18     *int64
	Srs_CyclicShiftCombinedGroupSequence_r18       *int64
	CyclicShiftHoppingWithinSubset_r18             *int64
	Srs_CyclicShiftCombinedCombOffset_r18          *int64
	Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18         *int64
	Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18      *int64
	Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18 *struct {
		MaxNumberPeriodicSRS_Resource_PerBWP_r18      int64
		MaxNumberAperiodicSRS_Resource_PerBWP_r18     int64
		MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 int64
		ValueY_SRS_ResourceAssociate_r18              int64
		ValueX_CSI_RS_ResourceAssociate_r18           int64
	}
	TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18 *struct {
		MaxNumberPeriodicSRS_r18       int64
		MaxNumberAperiodicSRS_r18      int64
		MaxNumberSemiPersistentSRS_r18 int64
		SimultaneousSRS_PerCC_r18      int64
		SimultaneousCSI_RS_NonCB_r18   int64
	}
	Dmrs_PortEntrySingleDCI_SDM_r18                *int64
	Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18         *int64
	Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18      *int64
	Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18 *struct {
		MaxNumberPeriodicSRS_Resource_PerBWP_r18      int64
		MaxNumberAperiodicSRS_Resource_PerBWP_r18     int64
		MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 int64
		ValueY_SRS_ResourceAssociate_r18              int64
		ValueX_CSI_RS_ResourceAssociate_r18           int64
	}
	TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18          *int64
	TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18       *int64
	TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18       *int64
	TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18    *int64
	TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18        *int64
	TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18                            *int64
	TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18                            *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18       *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18    *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18    *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18     *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18                         *int64
	TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18                         *int64
	Pucch_RepetitionDynamicIndicationSFN_r18                        *int64
	GroupBeamReporting_STx2P_r18                                    *struct {
		GroupL1_RSRP_Reporting_r18         int64
		MaxNumberBeamGroups_r18            int64
		MaxNumberResWithinSlotAcrossCC_r18 int64
		MaxNumberResAcrossCC_r18           int64
	}
	SimulConfigDMRS_DCI_1_3_r18            *int64
	CodebookParametersType1SP_SchemeA_r19  *CodebookParametersType1SP_SchemeA_r19
	CodebookParametersType1SP_SchemeB_r19  *CodebookParametersType1SP_SchemeB_r19
	CodebookParametersType1MP_r19          *CodebookParametersType1MP_r19
	CodebookParameterseType2Ext_r19        *CodebookParameterseType2Ext_r19
	CodebookParametersfeType2Ext_r19       *CodebookParametersfeType2Ext_r19
	CodebookParameterseType2DopplerExt_r19 *CodebookParameterseType2DopplerExt_r19
	CodebookParametersHybridBF_Type1SP_r19 *CodebookParametersHybridBF_Type1SP_r19
	CodebookParametersHybridBF_EType2_r19  *CodebookParametersHybridBF_EType2_r19
	Aiml_CSI_PredictionDoppler_r19         *CodebookParametersCSI_PredictionDoppler_r19
	Aiml_CSI_Report_r19                    []CPU_PoolInfo_r19
	IncreasedNumOfReportedRS_r19           *int64
	Aiml_BM_Case1_r19                      *struct {
		MaxNumberOfInferenceReportPerBWP_r19 struct {
			PeriodicReporting_r19       int64
			AperiodicReporting_r19      int64
			SemiPersistentReporting_r19 int64
		}
		MaxNumberOfInferenceReportAcrossAllCC_r19 int64
		MaxNumberOfResourceSetB_r19               int64
		MaxNumberOfResourceSetA_r19               int64
		ResourceTypeSetB_CSI_RS_r19               struct {
			Periodic_r19       *int64
			Aperiodic_r19      *int64
			SemiPersistent_r19 *int64
		}
		InferenceReportType_r19 struct {
			Periodic_r19       *int64
			Aperiodic_r19      *int64
			SemiPersistent_r19 *int64
		}
		SubUseCases_r19                                  int64
		MaxNumberOfPredictedBeamPerReportingInstance_r19 int64
		NumberOfOccupiedCPU_r19                          int64
		NumberOfOccupiedCPUx_r19                         int64
		RelaxationTimelineD_r19                          struct {
			Scs15kHz_r19  int64
			Scs30kHz_r19  int64
			Scs60kHz_r19  int64
			Scs120kHz_r19 int64
			Scs480kHz_r19 int64
			Scs960kHz_r19 int64
		}
		RelaxationTimelineD1_r19 struct {
			Scs15kHz_r19  int64
			Scs30kHz_r19  int64
			Scs60kHz_r19  int64
			Scs120kHz_r19 int64
			Scs480kHz_r19 int64
			Scs960kHz_r19 int64
		}
		OccupiedResourcePool_r19 int64
	}
	Aiml_BM_Case1_PredictedRSRP_r19 *int64
	Aiml_BM_Case2_r19               *struct {
		MaxNumberOfInferenceReportPerBWP_r19 struct {
			PeriodicReporting_r19       int64
			AperiodicReporting_r19      int64
			SemiPersistentReporting_r19 int64
		}
		MaxNumberOfInferenceReportAcrossAllCC_r19 int64
		MaxNumberOfResourceSetB_r19               int64
		MaxNumberOfResourceSetA_r19               int64
		MinNumberOfKBM_SetB_r19                   int64
		ResourceTypeSetB_CSI_RS_r19               struct {
			Periodic_r19       *int64
			SemiPersistent_r19 *int64
		}
		InferenceReportType_r19 struct {
			Periodic_r19       *int64
			Aperiodic_r19      *int64
			SemiPersistent_r19 *int64
		}
		MaxNumberOfPredictedBeamPerPerTimeInstance_r19 int64
		MaxNumberOfPredictedTimeInstance_r19           int64
		MaxTotalNumberOfPredictedBeamPerReport_r19     int64
		Dummy_TimeGap_r19                              int64
		NumberOfOccupiedCPU_r19                        int64
		NumberOfOccupiedCPUx_r19                       int64
		RelaxationTimelineD_r19                        struct {
			Scs15kHz_r19  int64
			Scs30kHz_r19  int64
			Scs60kHz_r19  int64
			Scs120kHz_r19 int64
			Scs480kHz_r19 int64
			Scs960kHz_r19 int64
		}
		RelaxationTimelineD1_r19 struct {
			Scs15kHz_r19  int64
			Scs30kHz_r19  int64
			Scs60kHz_r19  int64
			Scs120kHz_r19 int64
			Scs480kHz_r19 int64
			Scs960kHz_r19 int64
		}
		OccupiedResourcePool_r19 int64
	}
	Aiml_BM_Case2_PredictedRSRP_r19 *struct {
		MaxNumPredictedBeamPerInstance_r19      int64
		MaxNumPredictedTime_r19                 int64
		MaxTotalNumPredictedBeamInOneReport_r19 int64
	}
	Aiml_BM_Monitoring_r19 *struct {
		MaxNumTotalResource_r19          int64
		MaxNumReportPerBWP_Periodic_r19  int64
		MaxNumReportPerBWP_Aperiodic_r19 int64
		MaxNumReportPerBWP_SP_r19        int64
		MaxNumReportAcrossAllCC_r19      int64
		MaxNumOccasion_r19               int64
		MonitoringResourceType_r19       int64
		MonitoringReportType_r19         int64
	}
	Aiml_BM_UE_DataCollection_r19 *struct {
		SubCase_r19            int64
		MaxNumResourceSetB_r19 int64
		MaxNumResourceSetA_r19 int64
	}
	Aiml_CSI_Prediction_r19 *struct {
		SupportedCSI_RS_ResourceList_r19 []int64
		ScalingFactor_r19                int64
		NumberOfOccupiedCPU_r19          int64
		NumberOfOccupiedCPUx_r19         int64
		RelaxationTimelineT_r19          struct {
			Scs15kHz_r19  int64
			Scs30kHz_r19  int64
			Scs60kHz_r19  int64
			Scs120kHz_r19 int64
			Scs480kHz_r19 int64
			Scs960kHz_r19 int64
		}
		OccupiedResourcePool_r19 int64
		InferenceReportType_r19  struct {
			Aperiodic_r19      int64
			SemiPersistent_r19 int64
		}
	}
	Aiml_CSI_PredictionUnitDurationDD_r19 *int64
	Aiml_CSI_PredictionN4_r19             *struct {
		SupportedCSI_RS_ReportSettingAcrossAllCC_r19     []SupportedCSI_RS_ReportSetting_r18
		SupportedCSI_RS_ReportSettingAcrossOneReport_r19 []SupportedCSI_RS_ReportSetting_r18
		NumOccupiedCPU_r19                               int64
		NumOccupiedCPUx_r19                              int64
		OccupiedPool_r19                                 int64
	}
	Aiml_CSI_PredictionUE_DataCollection_r19 *int64
	Aiml_CSI_PredictionMonitoring_r19        *struct {
		SupportedCSI_RS_ResourceList_r19 []int64
		NumOccupiedCPU_r19               int64
	}
	Uei_ModeA_Event2_r19 *int64
	Uei_ModeB_r19        *struct {
		Scs15kHz_r19  *int64
		Scs30kHz_r19  *int64
		Scs60kHz_r19  *int64
		Scs120kHz_r19 *int64
		Scs480kHz_r19 *int64
		Scs960kHz_r19 *int64
	}
	Uei_TriggerEventDetermination_r19 *int64
	Uei_ModeA_Event1_r19              *int64
	Uei_ModeA_Event7_r19              *per.BitString
	Event2ConditionIndication_r19     *int64
	TimeRestriction128Port_r19        *int64
	GroupScalingFactor_r19            *int64
	Nes_SD_Type1_SP_r19               *struct {
		Timeline_r19            int64
		NumOfPortCSI_Report_r19 per.BitString
	}
	Mr_AlwaysReportedType1SP_r19 *int64
	Mr_AlwaysReported_EType2_r19 *int64
	Cjtc_DdReport_r19            *struct {
		MinRangeDdInCyclicPrefix_r19 int64
		MaxResolutionDd_r19          int64
		ScalingFactor_r19            int64
	}
	Cjtc_DdReportProcessing_r19 *struct {
		MaxNumberTRS_Resource_r19            int64
		MaxNumberTRS_ResourceAcrossCC_r19    int64
		MaxNumberCSI_RS_ResourcePerCC_r19    int64
		MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
		ValueX_r19                           int64
	}
	Cjtc_FO_Report_r19 *struct {
		MinRangeFO_r19      int64
		MaxResolutionFO_r19 int64
		ScalingFactor_r19   int64
	}
	Cjtc_FO_ReportProcessing_r19 *struct {
		MaxNumberTRS_Resource_r19            int64
		MaxNumberTRS_ResourceAcrossCC_r19    int64
		MaxNumberCSI_RS_ResourcePerCC_r19    int64
		MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
		ValueX_r19                           int64
	}
	Cjtc_PO_ReportWideband_r19 *struct {
		MaxResolution_r19   int64
		ScalingFactor_r19   int64
		MaxSlotDuration_r19 int64
	}
	Cjtc_PO_ReportWidebandProcessing_r19 *struct {
		MaxNumberCSI_RS_Configured_r19         int64
		MaxNumberCSI_RS_ConfiguredAcrossCC_r19 int64
		MaxNumberCSI_RS_ResourcePerCC_r19      int64
		MaxNumberCSI_RS_ResourceAcrossCC_r19   int64
		ValueX_r19                             int64
	}
	Cjtc_PO_ReportSubband_r19 *struct {
		MaxResolution_r19   int64
		MinSubbandSize_r19  int64
		ScalingFactor_r19   int64
		MaxSlotDuration_r19 int64
	}
	Cjtc_DdFO_Report_r19 *struct {
		MinRangeDdInCyclicPrefix_r19 int64
		MaxResolutionDd_r19          int64
		MinRangeFO_r19               int64
		MaxResolutionFO_r19          int64
		ScalingFactor_r19            int64
	}
	Cjtc_DdFO_ReportProcessing_r19 *struct {
		MaxNumberTRS_Resource_r19            int64
		MaxNumberTRS_ResourceAcrossCC_r19    int64
		MaxNumberCSI_RS_ResourcePerCC_r19    int64
		MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
		ValueX_r19                           int64
	}
	Cjt_QCL_PDSCH_SchemeC_r19                     *int64
	Cjt_QCL_PDSCH_SchemeD_r19                     *int64
	Cjt_QCL_PDSCH_SchemeE_r19                     *int64
	Linked_CJTC_Dd_EType2CJT_Joint_r19            *int64
	Linked_CJTC_Dd_EType2CJT_Separate_r19         *int64
	Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19 *int64
	TimelineRelax_CJTC_Dd_EType2CJT_r19           *struct {
		Scs15kHz_r19  *int64
		Scs30kHz_r19  *int64
		Scs60kHz_r19  *int64
		Scs120kHz_r19 *int64
		Scs480kHz_r19 *int64
		Scs960kHz_r19 *int64
	}
	NonCodebook_CSI_RS_SRS_3TxPUSCH_r19           []SupportedCSI_RS_Resource
	PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19    *int64
	PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19 *int64
	PathlossOffsetPRACH_JointTCI_r19              *int64
	PathlossOffsetPRACH_SeparateTCI_r19           *int64
	ExtendedStartBitDCI_2_3_r19                   *int64
	TwoSRS_PwrControlAdjust_r19                   *int64
	OverlapUL_TransReductionEnh_r19               *int64
	PathlossOffsetUpdate_r19                      *int64
	TwoSRS_TPC_DCI_2_3_r19                        *int64
	Srs_TPC_CLPC_AdjustmentState_r19              *int64
	TwoSRS_DCI_1_1_Separate_r19                   *int64
	TwoSRS_DCI_1_1_Joint_r19                      *int64
	PathlossOffsetPHR_r19                         *int64
	Aiml_BM_Case2_v1920                           *struct{ TimeGap_r19 per.BitString }
}

func (ie *MIMO_ParametersPerBand) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIMOParametersPerBandConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dummy6 != nil || ie.BeamManagementSSB_CSI_RS != nil || ie.BeamSwitchTiming != nil || ie.CodebookParameters != nil || ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Csi_RS_ForTracking != nil || ie.Srs_AssocCSI_RS != nil || ie.SpatialRelations != nil
	hasExtGroup1 := ie.DefaultQCL_TwoTCI_r16 != nil || ie.CodebookParametersPerBand_r16 != nil || ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil || ie.MaxNumberSCellBFR_r16 != nil || ie.SimultaneousReceptionDiffTypeD_r16 != nil || ie.Ssb_Csirs_SINR_Measurement_r16 != nil || ie.NonGroupSINR_Reporting_r16 != nil || ie.GroupSINR_Reporting_r16 != nil || ie.MultiDCI_MultiTRP_Parameters_r16 != nil || ie.SingleDCI_SDM_Scheme_Parameters_r16 != nil || ie.SupportFDM_SchemeA_r16 != nil || ie.SupportCodeWordSoftCombining_r16 != nil || ie.SupportTDM_SchemeA_r16 != nil || ie.SupportInter_SlotTDM_r16 != nil || ie.LowPAPR_DMRS_PDSCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil || ie.LowPAPR_DMRS_PUCCH_r16 != nil || ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil || ie.CodebookParametersAddition_r16 != nil || ie.CodebookComboParametersAddition_r16 != nil || ie.BeamCorrespondenceSSB_Based_r16 != nil || ie.BeamCorrespondenceCSI_RS_Based_r16 != nil || ie.BeamSwitchTiming_r16 != nil
	hasExtGroup2 := ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil || ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil
	hasExtGroup3 := ie.SpatialRelations_v1640 != nil || ie.Support64CandidateBeamRS_BFR_r16 != nil
	hasExtGroup4 := ie.MaxMIMO_LayersForMulti_DCI_MTRP_r16 != nil
	hasExtGroup5 := ie.SupportedSINR_Meas_v1670 != nil
	hasExtGroup6 := ie.Srs_IncreasedRepetition_r17 != nil || ie.Srs_PartialFrequencySounding_r17 != nil || ie.Srs_StartRB_LocationHoppingPartial_r17 != nil || ie.Srs_CombEight_r17 != nil || ie.CodebookParametersfetype2_r17 != nil || ie.MTRP_PUSCH_TwoCSI_RS_r17 != nil || ie.MTRP_PUCCH_InterSlot_r17 != nil || ie.MTRP_PUCCH_CyclicMapping_r17 != nil || ie.MTRP_PUCCH_SecondTPC_r17 != nil || ie.MTRP_BFR_TwoBFD_RS_Set_r17 != nil || ie.MTRP_BFR_PUCCH_SR_PerCG_r17 != nil || ie.MTRP_BFR_Association_PUCCH_SR_r17 != nil || ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil || ie.Sfn_DefaultDL_BeamSetup_r17 != nil || ie.Sfn_DefaultUL_BeamSetup_r17 != nil || ie.Srs_TriggeringOffset_r17 != nil || ie.Srs_TriggeringDCI_r17 != nil || ie.CodebookComboParameterMixedType_r17 != nil || ie.UnifiedJointTCI_r17 != nil || ie.UnifiedJointTCI_MultiMAC_CE_r17 != nil || ie.UnifiedJointTCI_PerBWP_CA_r17 != nil || ie.UnifiedJointTCI_ListSharingCA_r17 != nil || ie.UnifiedJointTCI_CommonMultiCC_r17 != nil || ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil || ie.UnifiedJointTCI_PC_Association_r17 != nil || ie.UnifiedJointTCI_Legacy_r17 != nil || ie.UnifiedJointTCI_Legacy_SRS_r17 != nil || ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil || ie.UnifiedJointTCI_SCellBFR_r17 != nil || ie.UnifiedJointTCI_InterCell_r17 != nil || ie.UnifiedSeparateTCI_r17 != nil || ie.UnifiedSeparateTCI_MultiMAC_CE_r17 != nil || ie.UnifiedSeparateTCI_PerBWP_CA_r17 != nil || ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil || ie.UnifiedSeparateTCI_CommonMultiCC_r17 != nil || ie.UnifiedSeparateTCI_InterCell_r17 != nil || ie.UnifiedJointTCI_MTRP_InterCell_BM_r17 != nil || ie.Mpe_Mitigation_r17 != nil || ie.Srs_PortReport_r17 != nil || ie.MTRP_PDCCH_Individual_r17 != nil || ie.MTRP_PDCCH_AnySpan_3Symbols_r17 != nil || ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil || ie.MTRP_PUSCH_CSI_RS_r17 != nil || ie.MTRP_PUSCH_CyclicMapping_r17 != nil || ie.MTRP_PUSCH_SecondTPC_r17 != nil || ie.MTRP_PUSCH_TwoPHR_Reporting_r17 != nil || ie.MTRP_PUSCH_A_CSI_r17 != nil || ie.MTRP_PUSCH_SP_CSI_r17 != nil || ie.MTRP_PUSCH_CG_r17 != nil || ie.MTRP_PUCCH_MAC_CE_r17 != nil || ie.MTRP_PUCCH_MaxNum_PC_FR1_r17 != nil || ie.MTRP_Inter_Cell_r17 != nil || ie.MTRP_GroupBasedL1_RSRP_r17 != nil || ie.MTRP_BFD_RS_MAC_CE_r17 != nil || ie.MTRP_CSI_EnhancementPerBand_r17 != nil || ie.CodebookComboParameterMultiTRP_r17 != nil || ie.MTRP_CSI_AdditionalCSI_r17 != nil || ie.MTRP_CSI_N_Max2_r17 != nil || ie.MTRP_CSI_CMR_r17 != nil || ie.Srs_PartialFreqSounding_r17 != nil || ie.BeamSwitchTiming_v1710 != nil || ie.BeamSwitchTiming_r17 != nil || ie.BeamReportTiming_v1710 != nil || ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil
	hasExtGroup7 := ie.Srs_PortReportSP_AP_r17 != nil || ie.MaxNumberRxBeam_v1720 != nil || ie.Sfn_ImplicitRS_TwoTCI_r17 != nil || ie.Sfn_QCL_TypeD_Collision_TwoTCI_r17 != nil || ie.MTRP_CSI_NumCPU_r17 != nil
	hasExtGroup8 := ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil
	hasExtGroup9 := ie.CodebookParametersetype2DopplerCSI_r18 != nil || ie.CodebookParametersfetype2DopplerCSI_r18 != nil || ie.CodebookParametersetype2CJT_r18 != nil || ie.CodebookParametersfetype2CJT_r18 != nil || ie.CodebookComboParametersCJT_r18 != nil || ie.CodebookParametersHARQ_ACK_PUSCH_r18 != nil || ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18 != nil || ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18 != nil || ie.Tci_SelectionDCI_r18 != nil || ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18 != nil || ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18 != nil || ie.Tci_SelectionAperiodicCSI_RS_r18 != nil || ie.Tci_SelectionAperiodicCSI_RS_M_DCI_r18 != nil || ie.TwoTCI_StatePDSCH_CJT_TxScheme_r18 != nil || ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil || ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil || ie.Tci_TRP_BFR_r18 != nil || ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil || ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil || ie.CommonTCI_SingleDCI_r18 != nil || ie.CommonTCI_MultiDCI_r18 != nil || ie.TwoPHR_Reporting_r18 != nil || ie.SpCell_TAG_Ind_r18 != nil || ie.InterCellCrossTRP_PDCCH_OrderCFRA_r18 != nil || ie.IntraCellCrossTRP_PDCCH_OrderCFRA_r18 != nil || ie.OverlapUL_TransReduction_r18 != nil || ie.MaxPeriodicityCMR_r18 != nil || ie.Tdcp_Report_r18 != nil || ie.Tdcp_Resource_r18 != nil || ie.TimelineRelax_CJT_CSI_r18 != nil || ie.JointConfigDMRSPortDynamicSwitching_r18 != nil || ie.Srs_CombOffsetHopping_r18 != nil || ie.Srs_CombOffsetInTime_r18 != nil || ie.Srs_CombOffsetCombinedGroupSequence_r18 != nil || ie.Srs_CombOffsetHoppingWithinSubset_r18 != nil || ie.Srs_CyclicShiftHopping_r18 != nil || ie.Srs_CyclicShiftHoppingSmallGranularity_r18 != nil || ie.Srs_CyclicShiftCombinedGroupSequence_r18 != nil || ie.CyclicShiftHoppingWithinSubset_r18 != nil || ie.Srs_CyclicShiftCombinedCombOffset_r18 != nil || ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil || ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil || ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18 != nil || ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18 != nil || ie.Dmrs_PortEntrySingleDCI_SDM_r18 != nil || ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil || ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil || ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18 != nil || ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18 != nil || ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18 != nil || ie.Pucch_RepetitionDynamicIndicationSFN_r18 != nil || ie.GroupBeamReporting_STx2P_r18 != nil
	hasExtGroup10 := ie.SimulConfigDMRS_DCI_1_3_r18 != nil
	hasExtGroup11 := ie.CodebookParametersType1SP_SchemeA_r19 != nil || ie.CodebookParametersType1SP_SchemeB_r19 != nil || ie.CodebookParametersType1MP_r19 != nil || ie.CodebookParameterseType2Ext_r19 != nil || ie.CodebookParametersfeType2Ext_r19 != nil || ie.CodebookParameterseType2DopplerExt_r19 != nil || ie.CodebookParametersHybridBF_Type1SP_r19 != nil || ie.CodebookParametersHybridBF_EType2_r19 != nil || ie.Aiml_CSI_PredictionDoppler_r19 != nil || ie.Aiml_CSI_Report_r19 != nil || ie.IncreasedNumOfReportedRS_r19 != nil || ie.Aiml_BM_Case1_r19 != nil || ie.Aiml_BM_Case1_PredictedRSRP_r19 != nil || ie.Aiml_BM_Case2_r19 != nil || ie.Aiml_BM_Case2_PredictedRSRP_r19 != nil || ie.Aiml_BM_Monitoring_r19 != nil || ie.Aiml_BM_UE_DataCollection_r19 != nil || ie.Aiml_CSI_Prediction_r19 != nil || ie.Aiml_CSI_PredictionUnitDurationDD_r19 != nil || ie.Aiml_CSI_PredictionN4_r19 != nil || ie.Aiml_CSI_PredictionUE_DataCollection_r19 != nil || ie.Aiml_CSI_PredictionMonitoring_r19 != nil || ie.Uei_ModeA_Event2_r19 != nil || ie.Uei_ModeB_r19 != nil || ie.Uei_TriggerEventDetermination_r19 != nil || ie.Uei_ModeA_Event1_r19 != nil || ie.Uei_ModeA_Event7_r19 != nil || ie.Event2ConditionIndication_r19 != nil || ie.TimeRestriction128Port_r19 != nil || ie.GroupScalingFactor_r19 != nil || ie.Nes_SD_Type1_SP_r19 != nil || ie.Mr_AlwaysReportedType1SP_r19 != nil || ie.Mr_AlwaysReported_EType2_r19 != nil || ie.Cjtc_DdReport_r19 != nil || ie.Cjtc_DdReportProcessing_r19 != nil || ie.Cjtc_FO_Report_r19 != nil || ie.Cjtc_FO_ReportProcessing_r19 != nil || ie.Cjtc_PO_ReportWideband_r19 != nil || ie.Cjtc_PO_ReportWidebandProcessing_r19 != nil || ie.Cjtc_PO_ReportSubband_r19 != nil || ie.Cjtc_DdFO_Report_r19 != nil || ie.Cjtc_DdFO_ReportProcessing_r19 != nil || ie.Cjt_QCL_PDSCH_SchemeC_r19 != nil || ie.Cjt_QCL_PDSCH_SchemeD_r19 != nil || ie.Cjt_QCL_PDSCH_SchemeE_r19 != nil || ie.Linked_CJTC_Dd_EType2CJT_Joint_r19 != nil || ie.Linked_CJTC_Dd_EType2CJT_Separate_r19 != nil || ie.Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19 != nil || ie.TimelineRelax_CJTC_Dd_EType2CJT_r19 != nil || ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19 != nil || ie.PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19 != nil || ie.PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19 != nil || ie.PathlossOffsetPRACH_JointTCI_r19 != nil || ie.PathlossOffsetPRACH_SeparateTCI_r19 != nil || ie.ExtendedStartBitDCI_2_3_r19 != nil || ie.TwoSRS_PwrControlAdjust_r19 != nil || ie.OverlapUL_TransReductionEnh_r19 != nil || ie.PathlossOffsetUpdate_r19 != nil || ie.TwoSRS_TPC_DCI_2_3_r19 != nil || ie.Srs_TPC_CLPC_AdjustmentState_r19 != nil || ie.TwoSRS_DCI_1_1_Separate_r19 != nil || ie.TwoSRS_DCI_1_1_Joint_r19 != nil || ie.PathlossOffsetPHR_r19 != nil
	hasExtGroup12 := ie.Aiml_BM_Case2_v1920 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tci_StatePDSCH != nil, ie.AdditionalActiveTCI_StatePDCCH != nil, ie.Pusch_TransCoherence != nil, ie.BeamCorrespondenceWithoutUL_BeamSweeping != nil, ie.PeriodicBeamReport != nil, ie.AperiodicBeamReport != nil, ie.Sp_BeamReportPUCCH != nil, ie.Sp_BeamReportPUSCH != nil, ie.Dummy1 != nil, ie.MaxNumberRxBeam != nil, ie.MaxNumberRxTxBeamSwitchDL != nil, ie.MaxNumberNonGroupBeamReporting != nil, ie.GroupBeamReporting != nil, ie.UplinkBeamManagement != nil, ie.MaxNumberCSI_RS_BFD != nil, ie.MaxNumberSSB_BFD != nil, ie.MaxNumberCSI_RS_SSB_CBD != nil, ie.Dummy2 != nil, ie.TwoPortsPTRS_UL != nil, ie.Dummy5 != nil, ie.Dummy3 != nil, ie.BeamReportTiming != nil, ie.Ptrs_DensityRecommendationSetDL != nil, ie.Ptrs_DensityRecommendationSetUL != nil, ie.Dummy4 != nil, ie.AperiodicTRS != nil}); err != nil {
		return err
	}

	// 3. tci-StatePDSCH: sequence
	{
		if ie.Tci_StatePDSCH != nil {
			c := ie.Tci_StatePDSCH
			mIMOParametersPerBandTciStatePDSCHSeq := e.NewSequenceEncoder(mIMOParametersPerBandTciStatePDSCHConstraints)
			if err := mIMOParametersPerBandTciStatePDSCHSeq.EncodePreamble([]bool{c.MaxNumberConfiguredTCI_StatesPerCC != nil, c.MaxNumberActiveTCI_PerBWP != nil}); err != nil {
				return err
			}
			if c.MaxNumberConfiguredTCI_StatesPerCC != nil {
				if err := e.EncodeEnumerated((*c.MaxNumberConfiguredTCI_StatesPerCC), mIMOParametersPerBandTciStatePDSCHMaxNumberConfiguredTCIStatesPerCCConstraints); err != nil {
					return err
				}
			}
			if c.MaxNumberActiveTCI_PerBWP != nil {
				if err := e.EncodeEnumerated((*c.MaxNumberActiveTCI_PerBWP), mIMOParametersPerBandTciStatePDSCHMaxNumberActiveTCIPerBWPConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. additionalActiveTCI-StatePDCCH: enumerated
	{
		if ie.AdditionalActiveTCI_StatePDCCH != nil {
			if err := e.EncodeEnumerated(*ie.AdditionalActiveTCI_StatePDCCH, mIMOParametersPerBandAdditionalActiveTCIStatePDCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 5. pusch-TransCoherence: enumerated
	{
		if ie.Pusch_TransCoherence != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_TransCoherence, mIMOParametersPerBandPuschTransCoherenceConstraints); err != nil {
				return err
			}
		}
	}

	// 6. beamCorrespondenceWithoutUL-BeamSweeping: enumerated
	{
		if ie.BeamCorrespondenceWithoutUL_BeamSweeping != nil {
			if err := e.EncodeEnumerated(*ie.BeamCorrespondenceWithoutUL_BeamSweeping, mIMOParametersPerBandBeamCorrespondenceWithoutULBeamSweepingConstraints); err != nil {
				return err
			}
		}
	}

	// 7. periodicBeamReport: enumerated
	{
		if ie.PeriodicBeamReport != nil {
			if err := e.EncodeEnumerated(*ie.PeriodicBeamReport, mIMOParametersPerBandPeriodicBeamReportConstraints); err != nil {
				return err
			}
		}
	}

	// 8. aperiodicBeamReport: enumerated
	{
		if ie.AperiodicBeamReport != nil {
			if err := e.EncodeEnumerated(*ie.AperiodicBeamReport, mIMOParametersPerBandAperiodicBeamReportConstraints); err != nil {
				return err
			}
		}
	}

	// 9. sp-BeamReportPUCCH: enumerated
	{
		if ie.Sp_BeamReportPUCCH != nil {
			if err := e.EncodeEnumerated(*ie.Sp_BeamReportPUCCH, mIMOParametersPerBandSpBeamReportPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 10. sp-BeamReportPUSCH: enumerated
	{
		if ie.Sp_BeamReportPUSCH != nil {
			if err := e.EncodeEnumerated(*ie.Sp_BeamReportPUSCH, mIMOParametersPerBandSpBeamReportPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 11. dummy1: ref
	{
		if ie.Dummy1 != nil {
			if err := ie.Dummy1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. maxNumberRxBeam: integer
	{
		if ie.MaxNumberRxBeam != nil {
			if err := e.EncodeInteger(*ie.MaxNumberRxBeam, mIMOParametersPerBandMaxNumberRxBeamConstraints); err != nil {
				return err
			}
		}
	}

	// 13. maxNumberRxTxBeamSwitchDL: sequence
	{
		if ie.MaxNumberRxTxBeamSwitchDL != nil {
			c := ie.MaxNumberRxTxBeamSwitchDL
			mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq := e.NewSequenceEncoder(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLConstraints)
			if err := mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil, c.Scs_240kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz), mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs15kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz), mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs30kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs120kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_240kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_240kHz), mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs240kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 14. maxNumberNonGroupBeamReporting: enumerated
	{
		if ie.MaxNumberNonGroupBeamReporting != nil {
			if err := e.EncodeEnumerated(*ie.MaxNumberNonGroupBeamReporting, mIMOParametersPerBandMaxNumberNonGroupBeamReportingConstraints); err != nil {
				return err
			}
		}
	}

	// 15. groupBeamReporting: enumerated
	{
		if ie.GroupBeamReporting != nil {
			if err := e.EncodeEnumerated(*ie.GroupBeamReporting, mIMOParametersPerBandGroupBeamReportingConstraints); err != nil {
				return err
			}
		}
	}

	// 16. uplinkBeamManagement: sequence
	{
		if ie.UplinkBeamManagement != nil {
			c := ie.UplinkBeamManagement
			if err := e.EncodeEnumerated(c.MaxNumberSRS_ResourcePerSet_BM, mIMOParametersPerBandUplinkBeamManagementMaxNumberSRSResourcePerSetBMConstraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourceSet, per.Constrained(1, 8)); err != nil {
				return err
			}
		}
	}

	// 17. maxNumberCSI-RS-BFD: integer
	{
		if ie.MaxNumberCSI_RS_BFD != nil {
			if err := e.EncodeInteger(*ie.MaxNumberCSI_RS_BFD, mIMOParametersPerBandMaxNumberCSIRSBFDConstraints); err != nil {
				return err
			}
		}
	}

	// 18. maxNumberSSB-BFD: integer
	{
		if ie.MaxNumberSSB_BFD != nil {
			if err := e.EncodeInteger(*ie.MaxNumberSSB_BFD, mIMOParametersPerBandMaxNumberSSBBFDConstraints); err != nil {
				return err
			}
		}
	}

	// 19. maxNumberCSI-RS-SSB-CBD: integer
	{
		if ie.MaxNumberCSI_RS_SSB_CBD != nil {
			if err := e.EncodeInteger(*ie.MaxNumberCSI_RS_SSB_CBD, mIMOParametersPerBandMaxNumberCSIRSSSBCBDConstraints); err != nil {
				return err
			}
		}
	}

	// 20. dummy2: enumerated
	{
		if ie.Dummy2 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy2, mIMOParametersPerBandDummy2Constraints); err != nil {
				return err
			}
		}
	}

	// 21. twoPortsPTRS-UL: enumerated
	{
		if ie.TwoPortsPTRS_UL != nil {
			if err := e.EncodeEnumerated(*ie.TwoPortsPTRS_UL, mIMOParametersPerBandTwoPortsPTRSULConstraints); err != nil {
				return err
			}
		}
	}

	// 22. dummy5: ref
	{
		if ie.Dummy5 != nil {
			if err := ie.Dummy5.Encode(e); err != nil {
				return err
			}
		}
	}

	// 23. dummy3: integer
	{
		if ie.Dummy3 != nil {
			if err := e.EncodeInteger(*ie.Dummy3, mIMOParametersPerBandDummy3Constraints); err != nil {
				return err
			}
		}
	}

	// 24. beamReportTiming: sequence
	{
		if ie.BeamReportTiming != nil {
			c := ie.BeamReportTiming
			mIMOParametersPerBandBeamReportTimingSeq := e.NewSequenceEncoder(mIMOParametersPerBandBeamReportTimingConstraints)
			if err := mIMOParametersPerBandBeamReportTimingSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz), mIMOParametersPerBandBeamReportTimingScs15kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz), mIMOParametersPerBandBeamReportTimingScs30kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz), mIMOParametersPerBandBeamReportTimingScs60kHzConstraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz), mIMOParametersPerBandBeamReportTimingScs120kHzConstraints); err != nil {
					return err
				}
			}
		}
	}

	// 25. ptrs-DensityRecommendationSetDL: sequence
	{
		if ie.Ptrs_DensityRecommendationSetDL != nil {
			c := ie.Ptrs_DensityRecommendationSetDL
			mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq := e.NewSequenceEncoder(mIMOParametersPerBandPtrsDensityRecommendationSetDLConstraints)
			if err := mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := c.Scs_15kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := c.Scs_30kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := c.Scs_60kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := c.Scs_120kHz.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 26. ptrs-DensityRecommendationSetUL: sequence
	{
		if ie.Ptrs_DensityRecommendationSetUL != nil {
			c := ie.Ptrs_DensityRecommendationSetUL
			mIMOParametersPerBandPtrsDensityRecommendationSetULSeq := e.NewSequenceEncoder(mIMOParametersPerBandPtrsDensityRecommendationSetULConstraints)
			if err := mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz != nil {
				if err := c.Scs_15kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_30kHz != nil {
				if err := c.Scs_30kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_60kHz != nil {
				if err := c.Scs_60kHz.Encode(e); err != nil {
					return err
				}
			}
			if c.Scs_120kHz != nil {
				if err := c.Scs_120kHz.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 27. dummy4: ref
	{
		if ie.Dummy4 != nil {
			if err := ie.Dummy4.Encode(e); err != nil {
				return err
			}
		}
	}

	// 28. aperiodicTRS: enumerated
	{
		if ie.AperiodicTRS != nil {
			if err := e.EncodeEnumerated(*ie.AperiodicTRS, mIMOParametersPerBandAperiodicTRSConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dummy6", Optional: true},
					{Name: "beamManagementSSB-CSI-RS", Optional: true},
					{Name: "beamSwitchTiming", Optional: true},
					{Name: "codebookParameters", Optional: true},
					{Name: "csi-RS-IM-ReceptionForFeedback", Optional: true},
					{Name: "csi-RS-ProcFrameworkForSRS", Optional: true},
					{Name: "csi-ReportFramework", Optional: true},
					{Name: "csi-RS-ForTracking", Optional: true},
					{Name: "srs-AssocCSI-RS", Optional: true},
					{Name: "spatialRelations", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy6 != nil, ie.BeamManagementSSB_CSI_RS != nil, ie.BeamSwitchTiming != nil, ie.CodebookParameters != nil, ie.Csi_RS_IM_ReceptionForFeedback != nil, ie.Csi_RS_ProcFrameworkForSRS != nil, ie.Csi_ReportFramework != nil, ie.Csi_RS_ForTracking != nil, ie.Srs_AssocCSI_RS != nil, ie.SpatialRelations != nil}); err != nil {
				return err
			}

			if ie.Dummy6 != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy6, mIMOParametersPerBandExtDummy6Constraints); err != nil {
					return err
				}
			}

			if ie.BeamManagementSSB_CSI_RS != nil {
				if err := ie.BeamManagementSSB_CSI_RS.Encode(ex); err != nil {
					return err
				}
			}

			if ie.BeamSwitchTiming != nil {
				c := ie.BeamSwitchTiming
				mIMOParametersPerBandExtBeamSwitchTimingSeq := ex.NewSequenceEncoder(mIMOParametersPerBandExtBeamSwitchTimingConstraints)
				if err := mIMOParametersPerBandExtBeamSwitchTimingSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz != nil {
					if err := ex.EncodeEnumerated((*c.Scs_60kHz), mIMOParametersPerBandExtBeamSwitchTimingScs60kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz != nil {
					if err := ex.EncodeEnumerated((*c.Scs_120kHz), mIMOParametersPerBandExtBeamSwitchTimingScs120kHzConstraints); err != nil {
						return err
					}
				}
			}

			if ie.CodebookParameters != nil {
				if err := ie.CodebookParameters.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_RS_IM_ReceptionForFeedback != nil {
				if err := ie.Csi_RS_IM_ReceptionForFeedback.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_RS_ProcFrameworkForSRS != nil {
				if err := ie.Csi_RS_ProcFrameworkForSRS.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_ReportFramework != nil {
				if err := ie.Csi_ReportFramework.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_RS_ForTracking != nil {
				if err := ie.Csi_RS_ForTracking.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Srs_AssocCSI_RS != nil {
				seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtSrsAssocCSIRSConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.Srs_AssocCSI_RS))); err != nil {
					return err
				}
				for i := range ie.Srs_AssocCSI_RS {
					if err := ie.Srs_AssocCSI_RS[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SpatialRelations != nil {
				if err := ie.SpatialRelations.Encode(ex); err != nil {
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
					{Name: "defaultQCL-TwoTCI-r16", Optional: true},
					{Name: "codebookParametersPerBand-r16", Optional: true},
					{Name: "simul-SpatialRelationUpdatePUCCHResGroup-r16", Optional: true},
					{Name: "maxNumberSCellBFR-r16", Optional: true},
					{Name: "simultaneousReceptionDiffTypeD-r16", Optional: true},
					{Name: "ssb-csirs-SINR-measurement-r16", Optional: true},
					{Name: "nonGroupSINR-reporting-r16", Optional: true},
					{Name: "groupSINR-reporting-r16", Optional: true},
					{Name: "multiDCI-multiTRP-Parameters-r16", Optional: true},
					{Name: "singleDCI-SDM-scheme-Parameters-r16", Optional: true},
					{Name: "supportFDM-SchemeA-r16", Optional: true},
					{Name: "supportCodeWordSoftCombining-r16", Optional: true},
					{Name: "supportTDM-SchemeA-r16", Optional: true},
					{Name: "supportInter-slotTDM-r16", Optional: true},
					{Name: "lowPAPR-DMRS-PDSCH-r16", Optional: true},
					{Name: "lowPAPR-DMRS-PUSCHwithoutPrecoding-r16", Optional: true},
					{Name: "lowPAPR-DMRS-PUCCH-r16", Optional: true},
					{Name: "lowPAPR-DMRS-PUSCHwithPrecoding-r16", Optional: true},
					{Name: "csi-ReportFrameworkExt-r16", Optional: true},
					{Name: "codebookParametersAddition-r16", Optional: true},
					{Name: "codebookComboParametersAddition-r16", Optional: true},
					{Name: "beamCorrespondenceSSB-based-r16", Optional: true},
					{Name: "beamCorrespondenceCSI-RS-based-r16", Optional: true},
					{Name: "beamSwitchTiming-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DefaultQCL_TwoTCI_r16 != nil, ie.CodebookParametersPerBand_r16 != nil, ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil, ie.MaxNumberSCellBFR_r16 != nil, ie.SimultaneousReceptionDiffTypeD_r16 != nil, ie.Ssb_Csirs_SINR_Measurement_r16 != nil, ie.NonGroupSINR_Reporting_r16 != nil, ie.GroupSINR_Reporting_r16 != nil, ie.MultiDCI_MultiTRP_Parameters_r16 != nil, ie.SingleDCI_SDM_Scheme_Parameters_r16 != nil, ie.SupportFDM_SchemeA_r16 != nil, ie.SupportCodeWordSoftCombining_r16 != nil, ie.SupportTDM_SchemeA_r16 != nil, ie.SupportInter_SlotTDM_r16 != nil, ie.LowPAPR_DMRS_PDSCH_r16 != nil, ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil, ie.LowPAPR_DMRS_PUCCH_r16 != nil, ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil, ie.Csi_ReportFrameworkExt_r16 != nil, ie.CodebookParametersAddition_r16 != nil, ie.CodebookComboParametersAddition_r16 != nil, ie.BeamCorrespondenceSSB_Based_r16 != nil, ie.BeamCorrespondenceCSI_RS_Based_r16 != nil, ie.BeamSwitchTiming_r16 != nil}); err != nil {
				return err
			}

			if ie.DefaultQCL_TwoTCI_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DefaultQCL_TwoTCI_r16, mIMOParametersPerBandExtDefaultQCLTwoTCIR16Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookParametersPerBand_r16 != nil {
				if err := ie.CodebookParametersPerBand_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16, mIMOParametersPerBandExtSimulSpatialRelationUpdatePUCCHResGroupR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberSCellBFR_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberSCellBFR_r16, mIMOParametersPerBandExtMaxNumberSCellBFRR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousReceptionDiffTypeD_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousReceptionDiffTypeD_r16, mIMOParametersPerBandExtSimultaneousReceptionDiffTypeDR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ssb_Csirs_SINR_Measurement_r16 != nil {
				c := ie.Ssb_Csirs_SINR_Measurement_r16
				mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Constraints)
				if err := mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Seq.EncodePreamble([]bool{c.SupportedSINR_Meas_r16 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberSSB_CSIRS_OneTx_CMR_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSOneTxCMRR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_IM_NZP_IMR_Res_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSIRS_2Tx_Res_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIRS2TxResR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberSSB_CSIRS_Res_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSResR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResMemR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.SupportedCSI_RS_Density_CMR_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedCSIRSDensityCMRR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberAperiodicCSI_RS_Res_r16, mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberAperiodicCSIRSResR16Constraints); err != nil {
					return err
				}
				if c.SupportedSINR_Meas_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SupportedSINR_Meas_r16), mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedSINRMeasR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.NonGroupSINR_Reporting_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.NonGroupSINR_Reporting_r16, mIMOParametersPerBandExtNonGroupSINRReportingR16Constraints); err != nil {
					return err
				}
			}

			if ie.GroupSINR_Reporting_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.GroupSINR_Reporting_r16, mIMOParametersPerBandExtGroupSINRReportingR16Constraints); err != nil {
					return err
				}
			}

			if ie.MultiDCI_MultiTRP_Parameters_r16 != nil {
				c := ie.MultiDCI_MultiTRP_Parameters_r16
				mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Constraints)
				if err := mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.EncodePreamble([]bool{c.OverlapPDSCHsFullyFreqTime_r16 != nil, c.OverlapPDSCHsInTimePartiallyFreq_r16 != nil, c.OutOfOrderOperationDL_r16 != nil, c.OutOfOrderOperationUL_r16 != nil, c.SeparateCRS_RateMatching_r16 != nil, c.DefaultQCL_PerCORESETPoolIndex_r16 != nil, c.MaxNumberActivatedTCI_States_r16 != nil}); err != nil {
					return err
				}
				if c.OverlapPDSCHsFullyFreqTime_r16 != nil {
					if err := ex.EncodeInteger((*c.OverlapPDSCHsFullyFreqTime_r16), per.Constrained(1, 2)); err != nil {
						return err
					}
				}
				if c.OverlapPDSCHsInTimePartiallyFreq_r16 != nil {
					if err := ex.EncodeEnumerated((*c.OverlapPDSCHsInTimePartiallyFreq_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OverlapPDSCHsInTimePartiallyFreqR16Constraints); err != nil {
						return err
					}
				}
				if c.OutOfOrderOperationDL_r16 != nil {
					c := (*c.OutOfOrderOperationDL_r16)
					mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Constraints)
					if err := mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq.EncodePreamble([]bool{c.SupportPDCCH_ToPDSCH_r16 != nil, c.SupportPDSCH_ToHARQ_ACK_r16 != nil}); err != nil {
						return err
					}
					if c.SupportPDCCH_ToPDSCH_r16 != nil {
						if err := ex.EncodeEnumerated((*c.SupportPDCCH_ToPDSCH_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDCCHToPDSCHR16Constraints); err != nil {
							return err
						}
					}
					if c.SupportPDSCH_ToHARQ_ACK_r16 != nil {
						if err := ex.EncodeEnumerated((*c.SupportPDSCH_ToHARQ_ACK_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDSCHToHARQACKR16Constraints); err != nil {
							return err
						}
					}
				}
				if c.OutOfOrderOperationUL_r16 != nil {
					if err := ex.EncodeEnumerated((*c.OutOfOrderOperationUL_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationULR16Constraints); err != nil {
						return err
					}
				}
				if c.SeparateCRS_RateMatching_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SeparateCRS_RateMatching_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16SeparateCRSRateMatchingR16Constraints); err != nil {
						return err
					}
				}
				if c.DefaultQCL_PerCORESETPoolIndex_r16 != nil {
					if err := ex.EncodeEnumerated((*c.DefaultQCL_PerCORESETPoolIndex_r16), mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16DefaultQCLPerCORESETPoolIndexR16Constraints); err != nil {
						return err
					}
				}
				if c.MaxNumberActivatedTCI_States_r16 != nil {
					c := (*c.MaxNumberActivatedTCI_States_r16)
					if err := ex.EncodeEnumerated(c.MaxNumberPerCORESET_Pool_r16, mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxNumberPerCORESETPoolR16Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.MaxTotalNumberAcrossCORESET_Pool_r16, mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxTotalNumberAcrossCORESETPoolR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.SingleDCI_SDM_Scheme_Parameters_r16 != nil {
				c := ie.SingleDCI_SDM_Scheme_Parameters_r16
				mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Constraints)
				if err := mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq.EncodePreamble([]bool{c.SupportNewDMRS_Port_r16 != nil, c.SupportTwoPortDL_PTRS_r16 != nil}); err != nil {
					return err
				}
				if c.SupportNewDMRS_Port_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SupportNewDMRS_Port_r16), mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportNewDMRSPortR16Constraints); err != nil {
						return err
					}
				}
				if c.SupportTwoPortDL_PTRS_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SupportTwoPortDL_PTRS_r16), mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportTwoPortDLPTRSR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.SupportFDM_SchemeA_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportFDM_SchemeA_r16, mIMOParametersPerBandExtSupportFDMSchemeAR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportCodeWordSoftCombining_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportCodeWordSoftCombining_r16, mIMOParametersPerBandExtSupportCodeWordSoftCombiningR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportTDM_SchemeA_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportTDM_SchemeA_r16, mIMOParametersPerBandExtSupportTDMSchemeAR16Constraints); err != nil {
					return err
				}
			}

			if ie.SupportInter_SlotTDM_r16 != nil {
				c := ie.SupportInter_SlotTDM_r16
				if err := ex.EncodeEnumerated(c.SupportRepNumPDSCH_TDRA_r16, mIMOParametersPerBandExtSupportInterSlotTDMR16SupportRepNumPDSCHTDRAR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxTBS_Size_r16, mIMOParametersPerBandExtSupportInterSlotTDMR16MaxTBSSizeR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberTCI_States_r16, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.LowPAPR_DMRS_PDSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.LowPAPR_DMRS_PDSCH_r16, mIMOParametersPerBandExtLowPAPRDMRSPDSCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16, mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithoutPrecodingR16Constraints); err != nil {
					return err
				}
			}

			if ie.LowPAPR_DMRS_PUCCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.LowPAPR_DMRS_PUCCH_r16, mIMOParametersPerBandExtLowPAPRDMRSPUCCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16, mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithPrecodingR16Constraints); err != nil {
					return err
				}
			}

			if ie.Csi_ReportFrameworkExt_r16 != nil {
				if err := ie.Csi_ReportFrameworkExt_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersAddition_r16 != nil {
				if err := ie.CodebookParametersAddition_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookComboParametersAddition_r16 != nil {
				if err := ie.CodebookComboParametersAddition_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.BeamCorrespondenceSSB_Based_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.BeamCorrespondenceSSB_Based_r16, mIMOParametersPerBandExtBeamCorrespondenceSSBBasedR16Constraints); err != nil {
					return err
				}
			}

			if ie.BeamCorrespondenceCSI_RS_Based_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.BeamCorrespondenceCSI_RS_Based_r16, mIMOParametersPerBandExtBeamCorrespondenceCSIRSBasedR16Constraints); err != nil {
					return err
				}
			}

			if ie.BeamSwitchTiming_r16 != nil {
				c := ie.BeamSwitchTiming_r16
				mIMOParametersPerBandExtBeamSwitchTimingR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtBeamSwitchTimingR16Constraints)
				if err := mIMOParametersPerBandExtBeamSwitchTimingR16Seq.EncodePreamble([]bool{c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_60kHz_r16), mIMOParametersPerBandExtBeamSwitchTimingR16Scs60kHzR16Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_120kHz_r16), mIMOParametersPerBandExtBeamSwitchTimingR16Scs120kHzR16Constraints); err != nil {
						return err
					}
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
					{Name: "semi-PersistentL1-SINR-Report-PUCCH-r16", Optional: true},
					{Name: "semi-PersistentL1-SINR-Report-PUSCH-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil, ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil}); err != nil {
				return err
			}

			if ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 != nil {
				c := ie.Semi_PersistentL1_SINR_Report_PUCCH_r16
				mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Constraints)
				if err := mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq.EncodePreamble([]bool{c.SupportReportFormat1_2OFDM_Syms_r16 != nil, c.SupportReportFormat4_14OFDM_Syms_r16 != nil}); err != nil {
					return err
				}
				if c.SupportReportFormat1_2OFDM_Syms_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SupportReportFormat1_2OFDM_Syms_r16), mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat12OFDMSymsR16Constraints); err != nil {
						return err
					}
				}
				if c.SupportReportFormat4_14OFDM_Syms_r16 != nil {
					if err := ex.EncodeEnumerated((*c.SupportReportFormat4_14OFDM_Syms_r16), mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat414OFDMSymsR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Semi_PersistentL1_SINR_Report_PUSCH_r16, mIMOParametersPerBandExtSemiPersistentL1SINRReportPUSCHR16Constraints); err != nil {
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
					{Name: "spatialRelations-v1640", Optional: true},
					{Name: "support64CandidateBeamRS-BFR-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SpatialRelations_v1640 != nil, ie.Support64CandidateBeamRS_BFR_r16 != nil}); err != nil {
				return err
			}

			if ie.SpatialRelations_v1640 != nil {
				c := ie.SpatialRelations_v1640
				if err := ex.EncodeEnumerated(c.MaxNumberConfiguredSpatialRelations_v1640, mIMOParametersPerBandExtSpatialRelationsV1640MaxNumberConfiguredSpatialRelationsV1640Constraints); err != nil {
					return err
				}
			}

			if ie.Support64CandidateBeamRS_BFR_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Support64CandidateBeamRS_BFR_r16, mIMOParametersPerBandExtSupport64CandidateBeamRSBFRR16Constraints); err != nil {
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
					{Name: "maxMIMO-LayersForMulti-DCI-mTRP-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_LayersForMulti_DCI_MTRP_r16 != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_LayersForMulti_DCI_MTRP_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxMIMO_LayersForMulti_DCI_MTRP_r16, mIMOParametersPerBandExtMaxMIMOLayersForMultiDCIMTRPR16Constraints); err != nil {
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
					{Name: "supportedSINR-meas-v1670", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportedSINR_Meas_v1670 != nil}); err != nil {
				return err
			}

			if ie.SupportedSINR_Meas_v1670 != nil {
				if err := ex.EncodeBitString(*ie.SupportedSINR_Meas_v1670, mIMOParametersPerBandSupportedSINRMeasV1670Constraints); err != nil {
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
					{Name: "srs-increasedRepetition-r17", Optional: true},
					{Name: "srs-partialFrequencySounding-r17", Optional: true},
					{Name: "srs-startRB-locationHoppingPartial-r17", Optional: true},
					{Name: "srs-combEight-r17", Optional: true},
					{Name: "codebookParametersfetype2-r17", Optional: true},
					{Name: "mTRP-PUSCH-twoCSI-RS-r17", Optional: true},
					{Name: "mTRP-PUCCH-InterSlot-r17", Optional: true},
					{Name: "mTRP-PUCCH-CyclicMapping-r17", Optional: true},
					{Name: "mTRP-PUCCH-SecondTPC-r17", Optional: true},
					{Name: "mTRP-BFR-twoBFD-RS-Set-r17", Optional: true},
					{Name: "mTRP-BFR-PUCCH-SR-perCG-r17", Optional: true},
					{Name: "mTRP-BFR-association-PUCCH-SR-r17", Optional: true},
					{Name: "sfn-SimulTwoTCI-AcrossMultiCC-r17", Optional: true},
					{Name: "sfn-DefaultDL-BeamSetup-r17", Optional: true},
					{Name: "sfn-DefaultUL-BeamSetup-r17", Optional: true},
					{Name: "srs-TriggeringOffset-r17", Optional: true},
					{Name: "srs-TriggeringDCI-r17", Optional: true},
					{Name: "codebookComboParameterMixedType-r17", Optional: true},
					{Name: "unifiedJointTCI-r17", Optional: true},
					{Name: "unifiedJointTCI-multiMAC-CE-r17", Optional: true},
					{Name: "unifiedJointTCI-perBWP-CA-r17", Optional: true},
					{Name: "unifiedJointTCI-ListSharingCA-r17", Optional: true},
					{Name: "unifiedJointTCI-commonMultiCC-r17", Optional: true},
					{Name: "unifiedJointTCI-BeamAlignDLRS-r17", Optional: true},
					{Name: "unifiedJointTCI-PC-association-r17", Optional: true},
					{Name: "unifiedJointTCI-Legacy-r17", Optional: true},
					{Name: "unifiedJointTCI-Legacy-SRS-r17", Optional: true},
					{Name: "unifiedJointTCI-Legacy-CORESET0-r17", Optional: true},
					{Name: "unifiedJointTCI-SCellBFR-r17", Optional: true},
					{Name: "unifiedJointTCI-InterCell-r17", Optional: true},
					{Name: "unifiedSeparateTCI-r17", Optional: true},
					{Name: "unifiedSeparateTCI-multiMAC-CE-r17", Optional: true},
					{Name: "unifiedSeparateTCI-perBWP-CA-r17", Optional: true},
					{Name: "unifiedSeparateTCI-ListSharingCA-r17", Optional: true},
					{Name: "unifiedSeparateTCI-commonMultiCC-r17", Optional: true},
					{Name: "unifiedSeparateTCI-InterCell-r17", Optional: true},
					{Name: "unifiedJointTCI-mTRP-InterCell-BM-r17", Optional: true},
					{Name: "mpe-Mitigation-r17", Optional: true},
					{Name: "srs-PortReport-r17", Optional: true},
					{Name: "mTRP-PDCCH-individual-r17", Optional: true},
					{Name: "mTRP-PDCCH-anySpan-3Symbols-r17", Optional: true},
					{Name: "mTRP-PDCCH-TwoQCL-TypeD-r17", Optional: true},
					{Name: "mTRP-PUSCH-CSI-RS-r17", Optional: true},
					{Name: "mTRP-PUSCH-cyclicMapping-r17", Optional: true},
					{Name: "mTRP-PUSCH-secondTPC-r17", Optional: true},
					{Name: "mTRP-PUSCH-twoPHR-Reporting-r17", Optional: true},
					{Name: "mTRP-PUSCH-A-CSI-r17", Optional: true},
					{Name: "mTRP-PUSCH-SP-CSI-r17", Optional: true},
					{Name: "mTRP-PUSCH-CG-r17", Optional: true},
					{Name: "mTRP-PUCCH-MAC-CE-r17", Optional: true},
					{Name: "mTRP-PUCCH-maxNum-PC-FR1-r17", Optional: true},
					{Name: "mTRP-inter-Cell-r17", Optional: true},
					{Name: "mTRP-GroupBasedL1-RSRP-r17", Optional: true},
					{Name: "mTRP-BFD-RS-MAC-CE-r17", Optional: true},
					{Name: "mTRP-CSI-EnhancementPerBand-r17", Optional: true},
					{Name: "codebookComboParameterMultiTRP-r17", Optional: true},
					{Name: "mTRP-CSI-additionalCSI-r17", Optional: true},
					{Name: "mTRP-CSI-N-Max2-r17", Optional: true},
					{Name: "mTRP-CSI-CMR-r17", Optional: true},
					{Name: "srs-partialFreqSounding-r17", Optional: true},
					{Name: "beamSwitchTiming-v1710", Optional: true},
					{Name: "beamSwitchTiming-r17", Optional: true},
					{Name: "beamReportTiming-v1710", Optional: true},
					{Name: "maxNumberRxTxBeamSwitchDL-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_IncreasedRepetition_r17 != nil, ie.Srs_PartialFrequencySounding_r17 != nil, ie.Srs_StartRB_LocationHoppingPartial_r17 != nil, ie.Srs_CombEight_r17 != nil, ie.CodebookParametersfetype2_r17 != nil, ie.MTRP_PUSCH_TwoCSI_RS_r17 != nil, ie.MTRP_PUCCH_InterSlot_r17 != nil, ie.MTRP_PUCCH_CyclicMapping_r17 != nil, ie.MTRP_PUCCH_SecondTPC_r17 != nil, ie.MTRP_BFR_TwoBFD_RS_Set_r17 != nil, ie.MTRP_BFR_PUCCH_SR_PerCG_r17 != nil, ie.MTRP_BFR_Association_PUCCH_SR_r17 != nil, ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil, ie.Sfn_DefaultDL_BeamSetup_r17 != nil, ie.Sfn_DefaultUL_BeamSetup_r17 != nil, ie.Srs_TriggeringOffset_r17 != nil, ie.Srs_TriggeringDCI_r17 != nil, ie.CodebookComboParameterMixedType_r17 != nil, ie.UnifiedJointTCI_r17 != nil, ie.UnifiedJointTCI_MultiMAC_CE_r17 != nil, ie.UnifiedJointTCI_PerBWP_CA_r17 != nil, ie.UnifiedJointTCI_ListSharingCA_r17 != nil, ie.UnifiedJointTCI_CommonMultiCC_r17 != nil, ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil, ie.UnifiedJointTCI_PC_Association_r17 != nil, ie.UnifiedJointTCI_Legacy_r17 != nil, ie.UnifiedJointTCI_Legacy_SRS_r17 != nil, ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil, ie.UnifiedJointTCI_SCellBFR_r17 != nil, ie.UnifiedJointTCI_InterCell_r17 != nil, ie.UnifiedSeparateTCI_r17 != nil, ie.UnifiedSeparateTCI_MultiMAC_CE_r17 != nil, ie.UnifiedSeparateTCI_PerBWP_CA_r17 != nil, ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil, ie.UnifiedSeparateTCI_CommonMultiCC_r17 != nil, ie.UnifiedSeparateTCI_InterCell_r17 != nil, ie.UnifiedJointTCI_MTRP_InterCell_BM_r17 != nil, ie.Mpe_Mitigation_r17 != nil, ie.Srs_PortReport_r17 != nil, ie.MTRP_PDCCH_Individual_r17 != nil, ie.MTRP_PDCCH_AnySpan_3Symbols_r17 != nil, ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil, ie.MTRP_PUSCH_CSI_RS_r17 != nil, ie.MTRP_PUSCH_CyclicMapping_r17 != nil, ie.MTRP_PUSCH_SecondTPC_r17 != nil, ie.MTRP_PUSCH_TwoPHR_Reporting_r17 != nil, ie.MTRP_PUSCH_A_CSI_r17 != nil, ie.MTRP_PUSCH_SP_CSI_r17 != nil, ie.MTRP_PUSCH_CG_r17 != nil, ie.MTRP_PUCCH_MAC_CE_r17 != nil, ie.MTRP_PUCCH_MaxNum_PC_FR1_r17 != nil, ie.MTRP_Inter_Cell_r17 != nil, ie.MTRP_GroupBasedL1_RSRP_r17 != nil, ie.MTRP_BFD_RS_MAC_CE_r17 != nil, ie.MTRP_CSI_EnhancementPerBand_r17 != nil, ie.CodebookComboParameterMultiTRP_r17 != nil, ie.MTRP_CSI_AdditionalCSI_r17 != nil, ie.MTRP_CSI_N_Max2_r17 != nil, ie.MTRP_CSI_CMR_r17 != nil, ie.Srs_PartialFreqSounding_r17 != nil, ie.BeamSwitchTiming_v1710 != nil, ie.BeamSwitchTiming_r17 != nil, ie.BeamReportTiming_v1710 != nil, ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil}); err != nil {
				return err
			}

			if ie.Srs_IncreasedRepetition_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_IncreasedRepetition_r17, mIMOParametersPerBandExtSrsIncreasedRepetitionR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_PartialFrequencySounding_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_PartialFrequencySounding_r17, mIMOParametersPerBandExtSrsPartialFrequencySoundingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_StartRB_LocationHoppingPartial_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_StartRB_LocationHoppingPartial_r17, mIMOParametersPerBandExtSrsStartRBLocationHoppingPartialR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CombEight_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CombEight_r17, mIMOParametersPerBandExtSrsCombEightR17Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookParametersfetype2_r17 != nil {
				if err := ie.CodebookParametersfetype2_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_TwoCSI_RS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_TwoCSI_RS_r17, mIMOParametersPerBandExtMTRPPUSCHTwoCSIRSR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUCCH_InterSlot_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUCCH_InterSlot_r17, mIMOParametersPerBandExtMTRPPUCCHInterSlotR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUCCH_CyclicMapping_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUCCH_CyclicMapping_r17, mIMOParametersPerBandExtMTRPPUCCHCyclicMappingR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUCCH_SecondTPC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUCCH_SecondTPC_r17, mIMOParametersPerBandExtMTRPPUCCHSecondTPCR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_BFR_TwoBFD_RS_Set_r17 != nil {
				c := ie.MTRP_BFR_TwoBFD_RS_Set_r17
				if err := ex.EncodeEnumerated(c.MaxBFD_RS_ResourcesPerSetPerBWP_r17, mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesPerSetPerBWPR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxBFR_r17, per.Constrained(1, 9)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17, mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesAcrossSetsPerBWPR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_BFR_PUCCH_SR_PerCG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_BFR_PUCCH_SR_PerCG_r17, mIMOParametersPerBandExtMTRPBFRPUCCHSRPerCGR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_BFR_Association_PUCCH_SR_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_BFR_Association_PUCCH_SR_r17, mIMOParametersPerBandExtMTRPBFRAssociationPUCCHSRR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17, mIMOParametersPerBandExtSfnSimulTwoTCIAcrossMultiCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sfn_DefaultDL_BeamSetup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sfn_DefaultDL_BeamSetup_r17, mIMOParametersPerBandExtSfnDefaultDLBeamSetupR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sfn_DefaultUL_BeamSetup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sfn_DefaultUL_BeamSetup_r17, mIMOParametersPerBandExtSfnDefaultULBeamSetupR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_TriggeringOffset_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_TriggeringOffset_r17, mIMOParametersPerBandExtSrsTriggeringOffsetR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_TriggeringDCI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_TriggeringDCI_r17, mIMOParametersPerBandExtSrsTriggeringDCIR17Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookComboParameterMixedType_r17 != nil {
				if err := ie.CodebookComboParameterMixedType_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_r17 != nil {
				c := ie.UnifiedJointTCI_r17
				if err := ex.EncodeEnumerated(c.MaxConfiguredJointTCI_r17, mIMOParametersPerBandExtUnifiedJointTCIR17MaxConfiguredJointTCIR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxActivatedTCIAcrossCC_r17, mIMOParametersPerBandExtUnifiedJointTCIR17MaxActivatedTCIAcrossCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_MultiMAC_CE_r17 != nil {
				c := ie.UnifiedJointTCI_MultiMAC_CE_r17
				mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Constraints)
				if err := mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Seq.EncodePreamble([]bool{c.MinBeamApplicationTime_r17 != nil}); err != nil {
					return err
				}
				if c.MinBeamApplicationTime_r17 != nil {
					if err := ex.EncodeEnumerated((*c.MinBeamApplicationTime_r17), mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MinBeamApplicationTimeR17Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeEnumerated(c.MaxNumMAC_CE_PerCC, mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MaxNumMACCEPerCCConstraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_PerBWP_CA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_PerBWP_CA_r17, mIMOParametersPerBandExtUnifiedJointTCIPerBWPCAR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_ListSharingCA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_ListSharingCA_r17, mIMOParametersPerBandExtUnifiedJointTCIListSharingCAR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_CommonMultiCC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_CommonMultiCC_r17, mIMOParametersPerBandExtUnifiedJointTCICommonMultiCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_BeamAlignDLRS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_BeamAlignDLRS_r17, mIMOParametersPerBandExtUnifiedJointTCIBeamAlignDLRSR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_PC_Association_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_PC_Association_r17, mIMOParametersPerBandExtUnifiedJointTCIPCAssociationR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_Legacy_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_Legacy_r17, mIMOParametersPerBandExtUnifiedJointTCILegacyR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_Legacy_SRS_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_Legacy_SRS_r17, mIMOParametersPerBandExtUnifiedJointTCILegacySRSR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_Legacy_CORESET0_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_Legacy_CORESET0_r17, mIMOParametersPerBandExtUnifiedJointTCILegacyCORESET0R17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_SCellBFR_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedJointTCI_SCellBFR_r17, mIMOParametersPerBandExtUnifiedJointTCISCellBFRR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_InterCell_r17 != nil {
				c := ie.UnifiedJointTCI_InterCell_r17
				if err := ex.EncodeEnumerated(c.AdditionalMAC_CE_PerCC_r17, mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEPerCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.AdditionalMAC_CE_AcrossCC_r17, mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEAcrossCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedSeparateTCI_r17 != nil {
				c := ie.UnifiedSeparateTCI_r17
				if err := ex.EncodeEnumerated(c.MaxConfiguredDL_TCI_r17, mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredDLTCIR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxConfiguredUL_TCI_r17, mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredULTCIR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxActivatedDL_TCIAcrossCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedDLTCIAcrossCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxActivatedUL_TCIAcrossCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedULTCIAcrossCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedSeparateTCI_MultiMAC_CE_r17 != nil {
				c := ie.UnifiedSeparateTCI_MultiMAC_CE_r17
				if err := ex.EncodeEnumerated(c.MinBeamApplicationTime_r17, mIMOParametersPerBandExtUnifiedSeparateTCIMultiMACCER17MinBeamApplicationTimeR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxActivatedDL_TCIPerCC_r17, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxActivatedUL_TCIPerCC_r17, per.Constrained(2, 8)); err != nil {
					return err
				}
			}

			if ie.UnifiedSeparateTCI_PerBWP_CA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedSeparateTCI_PerBWP_CA_r17, mIMOParametersPerBandExtUnifiedSeparateTCIPerBWPCAR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedSeparateTCI_ListSharingCA_r17 != nil {
				c := ie.UnifiedSeparateTCI_ListSharingCA_r17
				mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Constraints)
				if err := mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq.EncodePreamble([]bool{c.MaxNumListDL_TCI_r17 != nil, c.MaxNumListUL_TCI_r17 != nil}); err != nil {
					return err
				}
				if c.MaxNumListDL_TCI_r17 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumListDL_TCI_r17), mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListDLTCIR17Constraints); err != nil {
						return err
					}
				}
				if c.MaxNumListUL_TCI_r17 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumListUL_TCI_r17), mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListULTCIR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.UnifiedSeparateTCI_CommonMultiCC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UnifiedSeparateTCI_CommonMultiCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCICommonMultiCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedSeparateTCI_InterCell_r17 != nil {
				c := ie.UnifiedSeparateTCI_InterCell_r17
				if err := ex.EncodeEnumerated(c.K_DL_PerCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLPerCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.K_UL_PerCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULPerCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.K_DL_AcrossCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLAcrossCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.K_UL_AcrossCC_r17, mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULAcrossCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_MTRP_InterCell_BM_r17 != nil {
				c := ie.UnifiedJointTCI_MTRP_InterCell_BM_r17
				if err := ex.EncodeInteger(c.MaxNumAdditionalPCI_L1_RSRP_r17, per.Constrained(1, 7)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17, mIMOParametersPerBandExtUnifiedJointTCIMTRPInterCellBMR17MaxNumSSBResourceL1RSRPAcrossCCR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mpe_Mitigation_r17 != nil {
				c := ie.Mpe_Mitigation_r17
				if err := ex.EncodeInteger(c.MaxNumP_MPR_RI_Pairs_r17, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumConfRS_r17, mIMOParametersPerBandExtMpeMitigationR17MaxNumConfRSR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_PortReport_r17 != nil {
				c := ie.Srs_PortReport_r17
				mIMOParametersPerBandExtSrsPortReportR17Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtSrsPortReportR17Constraints)
				if err := mIMOParametersPerBandExtSrsPortReportR17Seq.EncodePreamble([]bool{c.CapVal1_r17 != nil, c.CapVal2_r17 != nil, c.CapVal3_r17 != nil, c.CapVal4_r17 != nil}); err != nil {
					return err
				}
				if c.CapVal1_r17 != nil {
					if err := ex.EncodeEnumerated((*c.CapVal1_r17), mIMOParametersPerBandExtSrsPortReportR17CapVal1R17Constraints); err != nil {
						return err
					}
				}
				if c.CapVal2_r17 != nil {
					if err := ex.EncodeEnumerated((*c.CapVal2_r17), mIMOParametersPerBandExtSrsPortReportR17CapVal2R17Constraints); err != nil {
						return err
					}
				}
				if c.CapVal3_r17 != nil {
					if err := ex.EncodeEnumerated((*c.CapVal3_r17), mIMOParametersPerBandExtSrsPortReportR17CapVal3R17Constraints); err != nil {
						return err
					}
				}
				if c.CapVal4_r17 != nil {
					if err := ex.EncodeEnumerated((*c.CapVal4_r17), mIMOParametersPerBandExtSrsPortReportR17CapVal4R17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.MTRP_PDCCH_Individual_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PDCCH_Individual_r17, mIMOParametersPerBandExtMTRPPDCCHIndividualR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PDCCH_AnySpan_3Symbols_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PDCCH_AnySpan_3Symbols_r17, mIMOParametersPerBandExtMTRPPDCCHAnySpan3SymbolsR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PDCCH_TwoQCL_TypeD_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PDCCH_TwoQCL_TypeD_r17, mIMOParametersPerBandExtMTRPPDCCHTwoQCLTypeDR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_CSI_RS_r17 != nil {
				c := ie.MTRP_PUSCH_CSI_RS_r17
				if err := ex.EncodeInteger(c.MaxNumPeriodicSRS_r17, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumAperiodicSRS_r17, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumSP_SRS_r17, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumSRS_ResourcePerCC_r17, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumSRS_ResourceNonCodebook_r17, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_CyclicMapping_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_CyclicMapping_r17, mIMOParametersPerBandExtMTRPPUSCHCyclicMappingR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_SecondTPC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_SecondTPC_r17, mIMOParametersPerBandExtMTRPPUSCHSecondTPCR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_TwoPHR_Reporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_TwoPHR_Reporting_r17, mIMOParametersPerBandExtMTRPPUSCHTwoPHRReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_A_CSI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_A_CSI_r17, mIMOParametersPerBandExtMTRPPUSCHACSIR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_SP_CSI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_SP_CSI_r17, mIMOParametersPerBandExtMTRPPUSCHSPCSIR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUSCH_CG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUSCH_CG_r17, mIMOParametersPerBandExtMTRPPUSCHCGR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUCCH_MAC_CE_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_PUCCH_MAC_CE_r17, mIMOParametersPerBandExtMTRPPUCCHMACCER17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_PUCCH_MaxNum_PC_FR1_r17 != nil {
				if err := ex.EncodeInteger(*ie.MTRP_PUCCH_MaxNum_PC_FR1_r17, mIMOParametersPerBandMTRPPUCCHMaxNumPCFR1R17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_Inter_Cell_r17 != nil {
				c := ie.MTRP_Inter_Cell_r17
				if err := ex.EncodeInteger(c.MaxNumAdditionalPCI_Case1_r17, per.Constrained(1, 7)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumAdditionalPCI_Case2_r17, per.Constrained(0, 7)); err != nil {
					return err
				}
			}

			if ie.MTRP_GroupBasedL1_RSRP_r17 != nil {
				c := ie.MTRP_GroupBasedL1_RSRP_r17
				if err := ex.EncodeInteger(c.MaxNumBeamGroups_r17, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumRS_WithinSlot_r17, mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSWithinSlotR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumRS_AcrossSlot_r17, mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSAcrossSlotR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_BFD_RS_MAC_CE_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_BFD_RS_MAC_CE_r17, mIMOParametersPerBandExtMTRPBFDRSMACCER17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_CSI_EnhancementPerBand_r17 != nil {
				c := ie.MTRP_CSI_EnhancementPerBand_r17
				if err := ex.EncodeInteger(c.MaxNumNZP_CSI_RS_r17, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.CSI_Report_Mode_r17, mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CSIReportModeR17Constraints); err != nil {
					return err
				}
				{
					seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17SupportedComboAcrossCCsR17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedComboAcrossCCs_r17))); err != nil {
						return err
					}
					for i := range c.SupportedComboAcrossCCs_r17 {
						if err := c.SupportedComboAcrossCCs_r17[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeEnumerated(c.CodebookModeNCJT_r17, mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CodebookModeNCJTR17Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookComboParameterMultiTRP_r17 != nil {
				if err := ie.CodebookComboParameterMultiTRP_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MTRP_CSI_AdditionalCSI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_CSI_AdditionalCSI_r17, mIMOParametersPerBandExtMTRPCSIAdditionalCSIR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_CSI_N_Max2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_CSI_N_Max2_r17, mIMOParametersPerBandExtMTRPCSINMax2R17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_CSI_CMR_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_CSI_CMR_r17, mIMOParametersPerBandExtMTRPCSICMRR17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_PartialFreqSounding_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_PartialFreqSounding_r17, mIMOParametersPerBandExtSrsPartialFreqSoundingR17Constraints); err != nil {
					return err
				}
			}

			if ie.BeamSwitchTiming_v1710 != nil {
				c := ie.BeamSwitchTiming_v1710
				mIMOParametersPerBandExtBeamSwitchTimingV1710Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtBeamSwitchTimingV1710Constraints)
				if err := mIMOParametersPerBandExtBeamSwitchTimingV1710Seq.EncodePreamble([]bool{c.Scs_480kHz != nil, c.Scs_960kHz != nil}); err != nil {
					return err
				}
				if c.Scs_480kHz != nil {
					if err := ex.EncodeEnumerated((*c.Scs_480kHz), mIMOParametersPerBandExtBeamSwitchTimingV1710Scs480kHzConstraints); err != nil {
						return err
					}
				}
				if c.Scs_960kHz != nil {
					if err := ex.EncodeEnumerated((*c.Scs_960kHz), mIMOParametersPerBandExtBeamSwitchTimingV1710Scs960kHzConstraints); err != nil {
						return err
					}
				}
			}

			if ie.BeamSwitchTiming_r17 != nil {
				c := ie.BeamSwitchTiming_r17
				mIMOParametersPerBandExtBeamSwitchTimingR17Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtBeamSwitchTimingR17Constraints)
				if err := mIMOParametersPerBandExtBeamSwitchTimingR17Seq.EncodePreamble([]bool{c.Scs_480kHz_r17 != nil, c.Scs_960kHz_r17 != nil}); err != nil {
					return err
				}
				if c.Scs_480kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_480kHz_r17), mIMOParametersPerBandExtBeamSwitchTimingR17Scs480kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_960kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_960kHz_r17), mIMOParametersPerBandExtBeamSwitchTimingR17Scs960kHzR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.BeamReportTiming_v1710 != nil {
				c := ie.BeamReportTiming_v1710
				mIMOParametersPerBandExtBeamReportTimingV1710Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtBeamReportTimingV1710Constraints)
				if err := mIMOParametersPerBandExtBeamReportTimingV1710Seq.EncodePreamble([]bool{c.Scs_480kHz_r17 != nil, c.Scs_960kHz_r17 != nil}); err != nil {
					return err
				}
				if c.Scs_480kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_480kHz_r17), mIMOParametersPerBandExtBeamReportTimingV1710Scs480kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_960kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_960kHz_r17), mIMOParametersPerBandExtBeamReportTimingV1710Scs960kHzR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.MaxNumberRxTxBeamSwitchDL_v1710 != nil {
				c := ie.MaxNumberRxTxBeamSwitchDL_v1710
				mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Constraints)
				if err := mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq.EncodePreamble([]bool{c.Scs_480kHz_r17 != nil, c.Scs_960kHz_r17 != nil}); err != nil {
					return err
				}
				if c.Scs_480kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_480kHz_r17), mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs480kHzR17Constraints); err != nil {
						return err
					}
				}
				if c.Scs_960kHz_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_960kHz_r17), mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs960kHzR17Constraints); err != nil {
						return err
					}
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
					{Name: "srs-PortReportSP-AP-r17", Optional: true},
					{Name: "maxNumberRxBeam-v1720", Optional: true},
					{Name: "sfn-ImplicitRS-twoTCI-r17", Optional: true},
					{Name: "sfn-QCL-TypeD-Collision-twoTCI-r17", Optional: true},
					{Name: "mTRP-CSI-numCPU-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Srs_PortReportSP_AP_r17 != nil, ie.MaxNumberRxBeam_v1720 != nil, ie.Sfn_ImplicitRS_TwoTCI_r17 != nil, ie.Sfn_QCL_TypeD_Collision_TwoTCI_r17 != nil, ie.MTRP_CSI_NumCPU_r17 != nil}); err != nil {
				return err
			}

			if ie.Srs_PortReportSP_AP_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_PortReportSP_AP_r17, mIMOParametersPerBandExtSrsPortReportSPAPR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberRxBeam_v1720 != nil {
				if err := ex.EncodeInteger(*ie.MaxNumberRxBeam_v1720, mIMOParametersPerBandMaxNumberRxBeamV1720Constraints); err != nil {
					return err
				}
			}

			if ie.Sfn_ImplicitRS_TwoTCI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sfn_ImplicitRS_TwoTCI_r17, mIMOParametersPerBandExtSfnImplicitRSTwoTCIR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sfn_QCL_TypeD_Collision_TwoTCI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sfn_QCL_TypeD_Collision_TwoTCI_r17, mIMOParametersPerBandExtSfnQCLTypeDCollisionTwoTCIR17Constraints); err != nil {
					return err
				}
			}

			if ie.MTRP_CSI_NumCPU_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MTRP_CSI_NumCPU_r17, mIMOParametersPerBandExtMTRPCSINumCPUR17Constraints); err != nil {
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
					{Name: "supportRepNumPDSCH-TDRA-DCI-1-2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil}); err != nil {
				return err
			}

			if ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17, mIMOParametersPerBandExtSupportRepNumPDSCHTDRADCI12R17Constraints); err != nil {
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
					{Name: "codebookParametersetype2DopplerCSI-r18", Optional: true},
					{Name: "codebookParametersfetype2DopplerCSI-r18", Optional: true},
					{Name: "codebookParametersetype2CJT-r18", Optional: true},
					{Name: "codebookParametersfetype2CJT-r18", Optional: true},
					{Name: "codebookComboParametersCJT-r18", Optional: true},
					{Name: "codebookParametersHARQ-ACK-PUSCH-r18", Optional: true},
					{Name: "tci-JointTCI-UpdateSingleActiveTCI-PerCC-r18", Optional: true},
					{Name: "tci-JointTCI-UpdateMultiActiveTCI-PerCC-r18", Optional: true},
					{Name: "tci-SelectionDCI-r18", Optional: true},
					{Name: "tci-SeparateTCI-UpdateSingleActiveTCI-PerCC-r18", Optional: true},
					{Name: "tci-SeparateTCI-UpdateMultiActiveTCI-PerCC-r18", Optional: true},
					{Name: "tci-SelectionAperiodicCSI-RS-r18", Optional: true},
					{Name: "tci-SelectionAperiodicCSI-RS-M-DCI-r18", Optional: true},
					{Name: "twoTCI-StatePDSCH-CJT-TxScheme-r18", Optional: true},
					{Name: "tci-JointTCI-UpdateSingleActiveTCI-PerCC-PerCORESET-r18", Optional: true},
					{Name: "tci-JointTCI-UpdateMultiActiveTCI-PerCC-PerCORESET-r18", Optional: true},
					{Name: "tci-TRP-BFR-r18", Optional: true},
					{Name: "tci-SeparateTCI-UpdateSingleActiveTCI-PerCC-PerCORESET-r18", Optional: true},
					{Name: "tci-SeparateTCI-UpdateMultiActiveTCI-PerCC-PerCORESET-r18", Optional: true},
					{Name: "commonTCI-SingleDCI-r18", Optional: true},
					{Name: "commonTCI-MultiDCI-r18", Optional: true},
					{Name: "twoPHR-Reporting-r18", Optional: true},
					{Name: "spCell-TAG-Ind-r18", Optional: true},
					{Name: "interCellCrossTRP-PDCCH-OrderCFRA-r18", Optional: true},
					{Name: "intraCellCrossTRP-PDCCH-OrderCFRA-r18", Optional: true},
					{Name: "overlapUL-TransReduction-r18", Optional: true},
					{Name: "maxPeriodicityCMR-r18", Optional: true},
					{Name: "tdcp-Report-r18", Optional: true},
					{Name: "tdcp-Resource-r18", Optional: true},
					{Name: "timelineRelax-CJT-CSI-r18", Optional: true},
					{Name: "jointConfigDMRSPortDynamicSwitching-r18", Optional: true},
					{Name: "srs-combOffsetHopping-r18", Optional: true},
					{Name: "srs-combOffsetInTime-r18", Optional: true},
					{Name: "srs-combOffsetCombinedGroupSequence-r18", Optional: true},
					{Name: "srs-combOffsetHoppingWithinSubset-r18", Optional: true},
					{Name: "srs-cyclicShiftHopping-r18", Optional: true},
					{Name: "srs-cyclicShiftHoppingSmallGranularity-r18", Optional: true},
					{Name: "srs-cyclicShiftCombinedGroupSequence-r18", Optional: true},
					{Name: "cyclicShiftHoppingWithinSubset-r18", Optional: true},
					{Name: "srs-cyclicShiftCombinedCombOffset-r18", Optional: true},
					{Name: "pusch-CB-2PTRS-SingleDCI-STx2P-SDM-r18", Optional: true},
					{Name: "pusch-NonCB-2PTRS-SingleDCI-STx2P-SDM-r18", Optional: true},
					{Name: "pusch-NonCB-SingleDCI-STx2P-SDM-CSI-RS-SRS-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-Multi-DCI-STx2P-CSI-RS-Resource-r18", Optional: true},
					{Name: "dmrs-PortEntrySingleDCI-SDM-r18", Optional: true},
					{Name: "pusch-CB-2PTRS-SingleDCI-STx2P-SFN-r18", Optional: true},
					{Name: "pusch-NonCB-2PTRS-SingleDCI-STx2P-SFN-r18", Optional: true},
					{Name: "pusch-NonCB-SingleDCI-STx2P-SFN-CSI-RS-SRS-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-FullTimeFullFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-FullTimePartialFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimeFullFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimePartialFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimeNonFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-CG-CG-r18", Optional: true},
					{Name: "twoPUSCH-CB-MultiDCI-STx2P-CG-DG-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-FullTimeFullFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-FullTimePartialFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimeFullFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimePartialFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimeNonFreqOverlap-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-CG-CG-r18", Optional: true},
					{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-CG-DG-r18", Optional: true},
					{Name: "pucch-RepetitionDynamicIndicationSFN-r18", Optional: true},
					{Name: "groupBeamReporting-STx2P-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CodebookParametersetype2DopplerCSI_r18 != nil, ie.CodebookParametersfetype2DopplerCSI_r18 != nil, ie.CodebookParametersetype2CJT_r18 != nil, ie.CodebookParametersfetype2CJT_r18 != nil, ie.CodebookComboParametersCJT_r18 != nil, ie.CodebookParametersHARQ_ACK_PUSCH_r18 != nil, ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18 != nil, ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18 != nil, ie.Tci_SelectionDCI_r18 != nil, ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18 != nil, ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18 != nil, ie.Tci_SelectionAperiodicCSI_RS_r18 != nil, ie.Tci_SelectionAperiodicCSI_RS_M_DCI_r18 != nil, ie.TwoTCI_StatePDSCH_CJT_TxScheme_r18 != nil, ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil, ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil, ie.Tci_TRP_BFR_r18 != nil, ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil, ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil, ie.CommonTCI_SingleDCI_r18 != nil, ie.CommonTCI_MultiDCI_r18 != nil, ie.TwoPHR_Reporting_r18 != nil, ie.SpCell_TAG_Ind_r18 != nil, ie.InterCellCrossTRP_PDCCH_OrderCFRA_r18 != nil, ie.IntraCellCrossTRP_PDCCH_OrderCFRA_r18 != nil, ie.OverlapUL_TransReduction_r18 != nil, ie.MaxPeriodicityCMR_r18 != nil, ie.Tdcp_Report_r18 != nil, ie.Tdcp_Resource_r18 != nil, ie.TimelineRelax_CJT_CSI_r18 != nil, ie.JointConfigDMRSPortDynamicSwitching_r18 != nil, ie.Srs_CombOffsetHopping_r18 != nil, ie.Srs_CombOffsetInTime_r18 != nil, ie.Srs_CombOffsetCombinedGroupSequence_r18 != nil, ie.Srs_CombOffsetHoppingWithinSubset_r18 != nil, ie.Srs_CyclicShiftHopping_r18 != nil, ie.Srs_CyclicShiftHoppingSmallGranularity_r18 != nil, ie.Srs_CyclicShiftCombinedGroupSequence_r18 != nil, ie.CyclicShiftHoppingWithinSubset_r18 != nil, ie.Srs_CyclicShiftCombinedCombOffset_r18 != nil, ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil, ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil, ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18 != nil, ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18 != nil, ie.Dmrs_PortEntrySingleDCI_SDM_r18 != nil, ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil, ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil, ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18 != nil, ie.Pucch_RepetitionDynamicIndicationSFN_r18 != nil, ie.GroupBeamReporting_STx2P_r18 != nil}); err != nil {
				return err
			}

			if ie.CodebookParametersetype2DopplerCSI_r18 != nil {
				if err := ie.CodebookParametersetype2DopplerCSI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersfetype2DopplerCSI_r18 != nil {
				if err := ie.CodebookParametersfetype2DopplerCSI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersetype2CJT_r18 != nil {
				if err := ie.CodebookParametersetype2CJT_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersfetype2CJT_r18 != nil {
				if err := ie.CodebookParametersfetype2CJT_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookComboParametersCJT_r18 != nil {
				if err := ie.CodebookComboParametersCJT_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersHARQ_ACK_PUSCH_r18 != nil {
				if err := ie.CodebookParametersHARQ_ACK_PUSCH_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18 != nil {
				c := ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18
				if err := ex.EncodeEnumerated(c.MaxNumberConfigJointTCIPerCC_PerBWP_r18, mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberActiveJointTCI_AcrossCC_r18, mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberActiveJointTCIAcrossCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18 != nil {
				c := ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18
				if err := ex.EncodeEnumerated(c.Tci_StateInd_r18, mIMOParametersPerBandExtTciJointTCIUpdateMultiActiveTCIPerCCR18TciStateIndR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberActiveJointTCI_PerCC_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
			}

			if ie.Tci_SelectionDCI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tci_SelectionDCI_r18, mIMOParametersPerBandExtTciSelectionDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18 != nil {
				c := ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18
				if err := ex.EncodeEnumerated(c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigULTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumActiveDL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumActiveUL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18 != nil {
				c := ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18
				if err := ex.EncodeEnumerated(c.MaxNumActiveDL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumActiveUL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SelectionAperiodicCSI_RS_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tci_SelectionAperiodicCSI_RS_r18, mIMOParametersPerBandExtTciSelectionAperiodicCSIRSR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SelectionAperiodicCSI_RS_M_DCI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tci_SelectionAperiodicCSI_RS_M_DCI_r18, mIMOParametersPerBandExtTciSelectionAperiodicCSIRSMDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoTCI_StatePDSCH_CJT_TxScheme_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoTCI_StatePDSCH_CJT_TxScheme_r18, mIMOParametersPerBandExtTwoTCIStatePDSCHCJTTxSchemeR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil {
				c := ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18
				if err := ex.EncodeEnumerated(c.MTRP_Operation_r18, mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberConfigJointTCIPerCC_PerBWP_r18, mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18, mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberActiveJointTCIAcrossCCPerCORESETR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil {
				if err := ex.EncodeInteger(*ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18, mIMOParametersPerBandTciJointTCIUpdateMultiActiveTCIPerCCPerCORESETR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_TRP_BFR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tci_TRP_BFR_r18, mIMOParametersPerBandExtTciTRPBFRR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 != nil {
				c := ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18
				if err := ex.EncodeEnumerated(c.MTRP_Operation_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigULTCIPerCCPerBWPR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumActiveDL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveDLTCIAcrossCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumActiveUL_TCI_AcrossCC_r18, mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveULTCIAcrossCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 != nil {
				c := ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18
				if err := ex.EncodeInteger(c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.CommonTCI_SingleDCI_r18 != nil {
				if err := ex.EncodeInteger(*ie.CommonTCI_SingleDCI_r18, mIMOParametersPerBandCommonTCISingleDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.CommonTCI_MultiDCI_r18 != nil {
				if err := ex.EncodeInteger(*ie.CommonTCI_MultiDCI_r18, mIMOParametersPerBandCommonTCIMultiDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPHR_Reporting_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPHR_Reporting_r18, mIMOParametersPerBandExtTwoPHRReportingR18Constraints); err != nil {
					return err
				}
			}

			if ie.SpCell_TAG_Ind_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SpCell_TAG_Ind_r18, mIMOParametersPerBandExtSpCellTAGIndR18Constraints); err != nil {
					return err
				}
			}

			if ie.InterCellCrossTRP_PDCCH_OrderCFRA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.InterCellCrossTRP_PDCCH_OrderCFRA_r18, mIMOParametersPerBandExtInterCellCrossTRPPDCCHOrderCFRAR18Constraints); err != nil {
					return err
				}
			}

			if ie.IntraCellCrossTRP_PDCCH_OrderCFRA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.IntraCellCrossTRP_PDCCH_OrderCFRA_r18, mIMOParametersPerBandExtIntraCellCrossTRPPDCCHOrderCFRAR18Constraints); err != nil {
					return err
				}
			}

			if ie.OverlapUL_TransReduction_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.OverlapUL_TransReduction_r18, mIMOParametersPerBandExtOverlapULTransReductionR18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxPeriodicityCMR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxPeriodicityCMR_r18, mIMOParametersPerBandExtMaxPeriodicityCMRR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tdcp_Report_r18 != nil {
				c := ie.Tdcp_Report_r18
				if err := ex.EncodeInteger(c.ValueX_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberActiveResource_r18, per.Constrained(2, 32)); err != nil {
					return err
				}
			}

			if ie.Tdcp_Resource_r18 != nil {
				c := ie.Tdcp_Resource_r18
				if err := ex.EncodeEnumerated(c.MaxNumberConfigPerCC_r18, mIMOParametersPerBandExtTdcpResourceR18MaxNumberConfigPerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigAcrossCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberSimultaneousPerCC_r18, mIMOParametersPerBandExtTdcpResourceR18MaxNumberSimultaneousPerCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.TimelineRelax_CJT_CSI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TimelineRelax_CJT_CSI_r18, mIMOParametersPerBandExtTimelineRelaxCJTCSIR18Constraints); err != nil {
					return err
				}
			}

			if ie.JointConfigDMRSPortDynamicSwitching_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.JointConfigDMRSPortDynamicSwitching_r18, mIMOParametersPerBandExtJointConfigDMRSPortDynamicSwitchingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CombOffsetHopping_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CombOffsetHopping_r18, mIMOParametersPerBandExtSrsCombOffsetHoppingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CombOffsetInTime_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CombOffsetInTime_r18, mIMOParametersPerBandExtSrsCombOffsetInTimeR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CombOffsetCombinedGroupSequence_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CombOffsetCombinedGroupSequence_r18, mIMOParametersPerBandExtSrsCombOffsetCombinedGroupSequenceR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CombOffsetHoppingWithinSubset_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CombOffsetHoppingWithinSubset_r18, mIMOParametersPerBandExtSrsCombOffsetHoppingWithinSubsetR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CyclicShiftHopping_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CyclicShiftHopping_r18, mIMOParametersPerBandExtSrsCyclicShiftHoppingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CyclicShiftHoppingSmallGranularity_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CyclicShiftHoppingSmallGranularity_r18, mIMOParametersPerBandExtSrsCyclicShiftHoppingSmallGranularityR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CyclicShiftCombinedGroupSequence_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CyclicShiftCombinedGroupSequence_r18, mIMOParametersPerBandExtSrsCyclicShiftCombinedGroupSequenceR18Constraints); err != nil {
					return err
				}
			}

			if ie.CyclicShiftHoppingWithinSubset_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.CyclicShiftHoppingWithinSubset_r18, mIMOParametersPerBandExtCyclicShiftHoppingWithinSubsetR18Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_CyclicShiftCombinedCombOffset_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_CyclicShiftCombinedCombOffset_r18, mIMOParametersPerBandExtSrsCyclicShiftCombinedCombOffsetR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18, mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSDMR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18, mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSDMR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18 != nil {
				c := ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18
				if err := ex.EncodeInteger(c.MaxNumberPeriodicSRS_Resource_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberAperiodicSRS_Resource_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberSemiPersistentSRS_ResourcePerBWP_r18, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueY_SRS_ResourceAssociate_r18, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_CSI_RS_ResourceAssociate_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18 != nil {
				c := ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18
				if err := ex.EncodeInteger(c.MaxNumberPeriodicSRS_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberAperiodicSRS_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberSemiPersistentSRS_r18, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SimultaneousSRS_PerCC_r18, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SimultaneousCSI_RS_NonCB_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Dmrs_PortEntrySingleDCI_SDM_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_PortEntrySingleDCI_SDM_r18, mIMOParametersPerBandExtDmrsPortEntrySingleDCISDMR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18, mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSFNR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18, mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSFNR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18 != nil {
				c := ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18
				if err := ex.EncodeInteger(c.MaxNumberPeriodicSRS_Resource_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberAperiodicSRS_Resource_PerBWP_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberSemiPersistentSRS_ResourcePerBWP_r18, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueY_SRS_ResourceAssociate_r18, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_CSI_RS_ResourceAssociate_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGCGR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18, mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGDGR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGCGR18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18, mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGDGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_RepetitionDynamicIndicationSFN_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_RepetitionDynamicIndicationSFN_r18, mIMOParametersPerBandExtPucchRepetitionDynamicIndicationSFNR18Constraints); err != nil {
					return err
				}
			}

			if ie.GroupBeamReporting_STx2P_r18 != nil {
				c := ie.GroupBeamReporting_STx2P_r18
				if err := ex.EncodeEnumerated(c.GroupL1_RSRP_Reporting_r18, mIMOParametersPerBandExtGroupBeamReportingSTx2PR18GroupL1RSRPReportingR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberBeamGroups_r18, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberResWithinSlotAcrossCC_r18, mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResWithinSlotAcrossCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberResAcrossCC_r18, mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResAcrossCCR18Constraints); err != nil {
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
					{Name: "simulConfigDMRS-DCI-1-3-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SimulConfigDMRS_DCI_1_3_r18 != nil}); err != nil {
				return err
			}

			if ie.SimulConfigDMRS_DCI_1_3_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SimulConfigDMRS_DCI_1_3_r18, mIMOParametersPerBandExtSimulConfigDMRSDCI13R18Constraints); err != nil {
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
					{Name: "codebookParametersType1SP-SchemeA-r19", Optional: true},
					{Name: "codebookParametersType1SP-SchemeB-r19", Optional: true},
					{Name: "codebookParametersType1MP-r19", Optional: true},
					{Name: "codebookParameterseType2Ext-r19", Optional: true},
					{Name: "codebookParametersfeType2Ext-r19", Optional: true},
					{Name: "codebookParameterseType2DopplerExt-r19", Optional: true},
					{Name: "codebookParametersHybridBF-Type1SP-r19", Optional: true},
					{Name: "codebookParametersHybridBF-eType2-r19", Optional: true},
					{Name: "aiml-CSI-PredictionDoppler-r19", Optional: true},
					{Name: "aiml-CSI-Report-r19", Optional: true},
					{Name: "increasedNumOfReportedRS-r19", Optional: true},
					{Name: "aiml-BM-Case1-r19", Optional: true},
					{Name: "aiml-BM-Case1-PredictedRSRP-r19", Optional: true},
					{Name: "aiml-BM-Case2-r19", Optional: true},
					{Name: "aiml-BM-Case2-PredictedRSRP-r19", Optional: true},
					{Name: "aiml-BM-Monitoring-r19", Optional: true},
					{Name: "aiml-BM-UE-DataCollection-r19", Optional: true},
					{Name: "aiml-CSI-Prediction-r19", Optional: true},
					{Name: "aiml-CSI-PredictionUnitDurationDD-r19", Optional: true},
					{Name: "aiml-CSI-PredictionN4-r19", Optional: true},
					{Name: "aiml-CSI-PredictionUE-DataCollection-r19", Optional: true},
					{Name: "aiml-CSI-PredictionMonitoring-r19", Optional: true},
					{Name: "uei-ModeA-Event2-r19", Optional: true},
					{Name: "uei-ModeB-r19", Optional: true},
					{Name: "uei-TriggerEventDetermination-r19", Optional: true},
					{Name: "uei-ModeA-Event1-r19", Optional: true},
					{Name: "uei-ModeA-Event7-r19", Optional: true},
					{Name: "event2ConditionIndication-r19", Optional: true},
					{Name: "timeRestriction128Port-r19", Optional: true},
					{Name: "groupScalingFactor-r19", Optional: true},
					{Name: "nes-SD-Type1-SP-r19", Optional: true},
					{Name: "mr-AlwaysReportedType1SP-r19", Optional: true},
					{Name: "mr-AlwaysReported-eType2-r19", Optional: true},
					{Name: "cjtc-DdReport-r19", Optional: true},
					{Name: "cjtc-DdReportProcessing-r19", Optional: true},
					{Name: "cjtc-FO-Report-r19", Optional: true},
					{Name: "cjtc-FO-ReportProcessing-r19", Optional: true},
					{Name: "cjtc-PO-ReportWideband-r19", Optional: true},
					{Name: "cjtc-PO-ReportWidebandProcessing-r19", Optional: true},
					{Name: "cjtc-PO-ReportSubband-r19", Optional: true},
					{Name: "cjtc-DdFO-Report-r19", Optional: true},
					{Name: "cjtc-DdFO-ReportProcessing-r19", Optional: true},
					{Name: "cjt-QCL-PDSCH-SchemeC-r19", Optional: true},
					{Name: "cjt-QCL-PDSCH-SchemeD-r19", Optional: true},
					{Name: "cjt-QCL-PDSCH-SchemeE-r19", Optional: true},
					{Name: "linked-CJTC-Dd-eType2CJT-Joint-r19", Optional: true},
					{Name: "linked-CJTC-Dd-eType2CJT-Separate-r19", Optional: true},
					{Name: "linked-CJTC-Dd-eType2CJT-SeparatePerState-r19", Optional: true},
					{Name: "timelineRelax-CJTC-Dd-eType2CJT-r19", Optional: true},
					{Name: "nonCodebook-CSI-RS-SRS-3TxPUSCH-r19", Optional: true},
					{Name: "pathlossOffsetPUCCH-PUSCH-SRS-JointTCI-r19", Optional: true},
					{Name: "pathlossOffsetPUCCH-PUSCH-SRS-SeparateTCI-r19", Optional: true},
					{Name: "pathlossOffsetPRACH-JointTCI-r19", Optional: true},
					{Name: "pathlossOffsetPRACH-SeparateTCI-r19", Optional: true},
					{Name: "extendedStartBitDCI-2-3-r19", Optional: true},
					{Name: "twoSRS-PwrControlAdjust-r19", Optional: true},
					{Name: "overlapUL-TransReductionEnh-r19", Optional: true},
					{Name: "pathlossOffsetUpdate-r19", Optional: true},
					{Name: "twoSRS-TPC-DCI-2-3-r19", Optional: true},
					{Name: "srs-TPC-CLPC-AdjustmentState-r19", Optional: true},
					{Name: "twoSRS-DCI-1-1-Separate-r19", Optional: true},
					{Name: "twoSRS-DCI-1-1-Joint-r19", Optional: true},
					{Name: "pathlossOffsetPHR-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CodebookParametersType1SP_SchemeA_r19 != nil, ie.CodebookParametersType1SP_SchemeB_r19 != nil, ie.CodebookParametersType1MP_r19 != nil, ie.CodebookParameterseType2Ext_r19 != nil, ie.CodebookParametersfeType2Ext_r19 != nil, ie.CodebookParameterseType2DopplerExt_r19 != nil, ie.CodebookParametersHybridBF_Type1SP_r19 != nil, ie.CodebookParametersHybridBF_EType2_r19 != nil, ie.Aiml_CSI_PredictionDoppler_r19 != nil, ie.Aiml_CSI_Report_r19 != nil, ie.IncreasedNumOfReportedRS_r19 != nil, ie.Aiml_BM_Case1_r19 != nil, ie.Aiml_BM_Case1_PredictedRSRP_r19 != nil, ie.Aiml_BM_Case2_r19 != nil, ie.Aiml_BM_Case2_PredictedRSRP_r19 != nil, ie.Aiml_BM_Monitoring_r19 != nil, ie.Aiml_BM_UE_DataCollection_r19 != nil, ie.Aiml_CSI_Prediction_r19 != nil, ie.Aiml_CSI_PredictionUnitDurationDD_r19 != nil, ie.Aiml_CSI_PredictionN4_r19 != nil, ie.Aiml_CSI_PredictionUE_DataCollection_r19 != nil, ie.Aiml_CSI_PredictionMonitoring_r19 != nil, ie.Uei_ModeA_Event2_r19 != nil, ie.Uei_ModeB_r19 != nil, ie.Uei_TriggerEventDetermination_r19 != nil, ie.Uei_ModeA_Event1_r19 != nil, ie.Uei_ModeA_Event7_r19 != nil, ie.Event2ConditionIndication_r19 != nil, ie.TimeRestriction128Port_r19 != nil, ie.GroupScalingFactor_r19 != nil, ie.Nes_SD_Type1_SP_r19 != nil, ie.Mr_AlwaysReportedType1SP_r19 != nil, ie.Mr_AlwaysReported_EType2_r19 != nil, ie.Cjtc_DdReport_r19 != nil, ie.Cjtc_DdReportProcessing_r19 != nil, ie.Cjtc_FO_Report_r19 != nil, ie.Cjtc_FO_ReportProcessing_r19 != nil, ie.Cjtc_PO_ReportWideband_r19 != nil, ie.Cjtc_PO_ReportWidebandProcessing_r19 != nil, ie.Cjtc_PO_ReportSubband_r19 != nil, ie.Cjtc_DdFO_Report_r19 != nil, ie.Cjtc_DdFO_ReportProcessing_r19 != nil, ie.Cjt_QCL_PDSCH_SchemeC_r19 != nil, ie.Cjt_QCL_PDSCH_SchemeD_r19 != nil, ie.Cjt_QCL_PDSCH_SchemeE_r19 != nil, ie.Linked_CJTC_Dd_EType2CJT_Joint_r19 != nil, ie.Linked_CJTC_Dd_EType2CJT_Separate_r19 != nil, ie.Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19 != nil, ie.TimelineRelax_CJTC_Dd_EType2CJT_r19 != nil, ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19 != nil, ie.PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19 != nil, ie.PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19 != nil, ie.PathlossOffsetPRACH_JointTCI_r19 != nil, ie.PathlossOffsetPRACH_SeparateTCI_r19 != nil, ie.ExtendedStartBitDCI_2_3_r19 != nil, ie.TwoSRS_PwrControlAdjust_r19 != nil, ie.OverlapUL_TransReductionEnh_r19 != nil, ie.PathlossOffsetUpdate_r19 != nil, ie.TwoSRS_TPC_DCI_2_3_r19 != nil, ie.Srs_TPC_CLPC_AdjustmentState_r19 != nil, ie.TwoSRS_DCI_1_1_Separate_r19 != nil, ie.TwoSRS_DCI_1_1_Joint_r19 != nil, ie.PathlossOffsetPHR_r19 != nil}); err != nil {
				return err
			}

			if ie.CodebookParametersType1SP_SchemeA_r19 != nil {
				if err := ie.CodebookParametersType1SP_SchemeA_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersType1SP_SchemeB_r19 != nil {
				if err := ie.CodebookParametersType1SP_SchemeB_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersType1MP_r19 != nil {
				if err := ie.CodebookParametersType1MP_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParameterseType2Ext_r19 != nil {
				if err := ie.CodebookParameterseType2Ext_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersfeType2Ext_r19 != nil {
				if err := ie.CodebookParametersfeType2Ext_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParameterseType2DopplerExt_r19 != nil {
				if err := ie.CodebookParameterseType2DopplerExt_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersHybridBF_Type1SP_r19 != nil {
				if err := ie.CodebookParametersHybridBF_Type1SP_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookParametersHybridBF_EType2_r19 != nil {
				if err := ie.CodebookParametersHybridBF_EType2_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_PredictionDoppler_r19 != nil {
				if err := ie.Aiml_CSI_PredictionDoppler_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_Report_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtAimlCSIReportR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Aiml_CSI_Report_r19))); err != nil {
					return err
				}
				for i := range ie.Aiml_CSI_Report_r19 {
					if err := ie.Aiml_CSI_Report_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IncreasedNumOfReportedRS_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.IncreasedNumOfReportedRS_r19, mIMOParametersPerBandExtIncreasedNumOfReportedRSR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case1_r19 != nil {
				c := ie.Aiml_BM_Case1_r19
				{
					c := &c.MaxNumberOfInferenceReportPerBWP_r19
					if err := ex.EncodeInteger(c.PeriodicReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.AperiodicReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.SemiPersistentReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfInferenceReportAcrossAllCC_r19, mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfResourceSetB_r19, mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetBR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfResourceSetA_r19, mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetAR19Constraints); err != nil {
					return err
				}
				{
					c := &c.ResourceTypeSetB_CSI_RS_r19
					mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Constraints)
					if err := mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq.EncodePreamble([]bool{c.Periodic_r19 != nil, c.Aperiodic_r19 != nil, c.SemiPersistent_r19 != nil}); err != nil {
						return err
					}
					if c.Periodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Periodic_r19), mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints); err != nil {
							return err
						}
					}
					if c.Aperiodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Aperiodic_r19), mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19AperiodicR19Constraints); err != nil {
							return err
						}
					}
					if c.SemiPersistent_r19 != nil {
						if err := ex.EncodeEnumerated((*c.SemiPersistent_r19), mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints); err != nil {
							return err
						}
					}
				}
				{
					c := &c.InferenceReportType_r19
					mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Constraints)
					if err := mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq.EncodePreamble([]bool{c.Periodic_r19 != nil, c.Aperiodic_r19 != nil, c.SemiPersistent_r19 != nil}); err != nil {
						return err
					}
					if c.Periodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Periodic_r19), mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19PeriodicR19Constraints); err != nil {
							return err
						}
					}
					if c.Aperiodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Aperiodic_r19), mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19AperiodicR19Constraints); err != nil {
							return err
						}
					}
					if c.SemiPersistent_r19 != nil {
						if err := ex.EncodeEnumerated((*c.SemiPersistent_r19), mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19SemiPersistentR19Constraints); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeEnumerated(c.SubUseCases_r19, mIMOParametersPerBandExtAimlBMCase1R19SubUseCasesR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfPredictedBeamPerReportingInstance_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPU_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPUx_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				{
					c := &c.RelaxationTimelineD_r19
					if err := ex.EncodeEnumerated(c.Scs15kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs15kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs30kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs30kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs60kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs60kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs120kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs120kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs480kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs480kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs960kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
				{
					c := &c.RelaxationTimelineD1_r19
					if err := ex.EncodeEnumerated(c.Scs15kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs15kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs30kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs30kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs60kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs60kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs120kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs120kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs480kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs480kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs960kHz_r19, mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeInteger(c.OccupiedResourcePool_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case1_PredictedRSRP_r19 != nil {
				if err := ex.EncodeInteger(*ie.Aiml_BM_Case1_PredictedRSRP_r19, mIMOParametersPerBandAimlBMCase1PredictedRSRPR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case2_r19 != nil {
				c := ie.Aiml_BM_Case2_r19
				{
					c := &c.MaxNumberOfInferenceReportPerBWP_r19
					if err := ex.EncodeInteger(c.PeriodicReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.AperiodicReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.SemiPersistentReporting_r19, per.Constrained(1, 4)); err != nil {
						return err
					}
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfInferenceReportAcrossAllCC_r19, mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfResourceSetB_r19, mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetBR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfResourceSetA_r19, mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetAR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MinNumberOfKBM_SetB_r19, mIMOParametersPerBandExtAimlBMCase2R19MinNumberOfKBMSetBR19Constraints); err != nil {
					return err
				}
				{
					c := &c.ResourceTypeSetB_CSI_RS_r19
					mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Constraints)
					if err := mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq.EncodePreamble([]bool{c.Periodic_r19 != nil, c.SemiPersistent_r19 != nil}); err != nil {
						return err
					}
					if c.Periodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Periodic_r19), mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints); err != nil {
							return err
						}
					}
					if c.SemiPersistent_r19 != nil {
						if err := ex.EncodeEnumerated((*c.SemiPersistent_r19), mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints); err != nil {
							return err
						}
					}
				}
				{
					c := &c.InferenceReportType_r19
					mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Constraints)
					if err := mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq.EncodePreamble([]bool{c.Periodic_r19 != nil, c.Aperiodic_r19 != nil, c.SemiPersistent_r19 != nil}); err != nil {
						return err
					}
					if c.Periodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Periodic_r19), mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19PeriodicR19Constraints); err != nil {
							return err
						}
					}
					if c.Aperiodic_r19 != nil {
						if err := ex.EncodeEnumerated((*c.Aperiodic_r19), mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19AperiodicR19Constraints); err != nil {
							return err
						}
					}
					if c.SemiPersistent_r19 != nil {
						if err := ex.EncodeEnumerated((*c.SemiPersistent_r19), mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19SemiPersistentR19Constraints); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeInteger(c.MaxNumberOfPredictedBeamPerPerTimeInstance_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfPredictedTimeInstance_r19, mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfPredictedTimeInstanceR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxTotalNumberOfPredictedBeamPerReport_r19, mIMOParametersPerBandExtAimlBMCase2R19MaxTotalNumberOfPredictedBeamPerReportR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Dummy_TimeGap_r19, mIMOParametersPerBandExtAimlBMCase2R19DummyTimeGapR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPU_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPUx_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				{
					c := &c.RelaxationTimelineD_r19
					if err := ex.EncodeEnumerated(c.Scs15kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs15kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs30kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs30kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs60kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs60kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs120kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs120kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs480kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs480kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs960kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
				{
					c := &c.RelaxationTimelineD1_r19
					if err := ex.EncodeEnumerated(c.Scs15kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs15kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs30kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs30kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs60kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs60kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs120kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs120kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs480kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs480kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs960kHz_r19, mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeInteger(c.OccupiedResourcePool_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case2_PredictedRSRP_r19 != nil {
				c := ie.Aiml_BM_Case2_PredictedRSRP_r19
				if err := ex.EncodeInteger(c.MaxNumPredictedBeamPerInstance_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumPredictedTime_r19, mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxNumPredictedTimeR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxTotalNumPredictedBeamInOneReport_r19, mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxTotalNumPredictedBeamInOneReportR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Monitoring_r19 != nil {
				c := ie.Aiml_BM_Monitoring_r19
				if err := ex.EncodeEnumerated(c.MaxNumTotalResource_r19, mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumTotalResourceR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumReportPerBWP_Periodic_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumReportPerBWP_Aperiodic_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumReportPerBWP_SP_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumReportAcrossAllCC_r19, mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumReportAcrossAllCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOccasion_r19, mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumOccasionR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MonitoringResourceType_r19, mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringResourceTypeR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MonitoringReportType_r19, mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringReportTypeR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_UE_DataCollection_r19 != nil {
				c := ie.Aiml_BM_UE_DataCollection_r19
				if err := ex.EncodeEnumerated(c.SubCase_r19, mIMOParametersPerBandExtAimlBMUEDataCollectionR19SubCaseR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumResourceSetB_r19, mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetBR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumResourceSetA_r19, mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetAR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_Prediction_r19 != nil {
				c := ie.Aiml_CSI_Prediction_r19
				{
					seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtAimlCSIPredictionR19SupportedCSIRSResourceListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r19))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ResourceList_r19 {
						if err := ex.EncodeInteger(c.SupportedCSI_RS_ResourceList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeEnumerated(c.ScalingFactor_r19, mIMOParametersPerBandExtAimlCSIPredictionR19ScalingFactorR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPU_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfOccupiedCPUx_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				{
					c := &c.RelaxationTimelineT_r19
					if err := ex.EncodeEnumerated(c.Scs15kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs15kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs30kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs30kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs60kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs60kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs120kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs120kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs480kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs480kHzR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.Scs960kHz_r19, mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeInteger(c.OccupiedResourcePool_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
				{
					c := &c.InferenceReportType_r19
					if err := ex.EncodeEnumerated(c.Aperiodic_r19, mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19AperiodicR19Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.SemiPersistent_r19, mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19SemiPersistentR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Aiml_CSI_PredictionUnitDurationDD_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Aiml_CSI_PredictionUnitDurationDD_r19, mIMOParametersPerBandExtAimlCSIPredictionUnitDurationDDR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_PredictionN4_r19 != nil {
				c := ie.Aiml_CSI_PredictionN4_r19
				{
					seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossAllCCR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingAcrossAllCC_r19))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ReportSettingAcrossAllCC_r19 {
						if err := c.SupportedCSI_RS_ReportSettingAcrossAllCC_r19[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				{
					seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossOneReportR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ReportSettingAcrossOneReport_r19))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ReportSettingAcrossOneReport_r19 {
						if err := c.SupportedCSI_RS_ReportSettingAcrossOneReport_r19[i].Encode(ex); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeInteger(c.NumOccupiedCPU_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumOccupiedCPUx_r19, per.Constrained(0, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.OccupiedPool_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_PredictionUE_DataCollection_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Aiml_CSI_PredictionUE_DataCollection_r19, mIMOParametersPerBandExtAimlCSIPredictionUEDataCollectionR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_CSI_PredictionMonitoring_r19 != nil {
				c := ie.Aiml_CSI_PredictionMonitoring_r19
				{
					seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtAimlCSIPredictionMonitoringR19SupportedCSIRSResourceListR19Constraints)
					if err := seqOf.EncodeLength(int64(len(c.SupportedCSI_RS_ResourceList_r19))); err != nil {
						return err
					}
					for i := range c.SupportedCSI_RS_ResourceList_r19 {
						if err := ex.EncodeInteger(c.SupportedCSI_RS_ResourceList_r19[i], per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16)); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeInteger(c.NumOccupiedCPU_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Uei_ModeA_Event2_r19 != nil {
				if err := ex.EncodeInteger(*ie.Uei_ModeA_Event2_r19, mIMOParametersPerBandUeiModeAEvent2R19Constraints); err != nil {
					return err
				}
			}

			if ie.Uei_ModeB_r19 != nil {
				c := ie.Uei_ModeB_r19
				mIMOParametersPerBandExtUeiModeBR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtUeiModeBR19Constraints)
				if err := mIMOParametersPerBandExtUeiModeBR19Seq.EncodePreamble([]bool{c.Scs15kHz_r19 != nil, c.Scs30kHz_r19 != nil, c.Scs60kHz_r19 != nil, c.Scs120kHz_r19 != nil, c.Scs480kHz_r19 != nil, c.Scs960kHz_r19 != nil}); err != nil {
					return err
				}
				if c.Scs15kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs15kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs15kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs30kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs30kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs30kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs60kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs60kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs60kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs120kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs120kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs120kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs480kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs480kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs480kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs960kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs960kHz_r19), mIMOParametersPerBandExtUeiModeBR19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Uei_TriggerEventDetermination_r19 != nil {
				if err := ex.EncodeInteger(*ie.Uei_TriggerEventDetermination_r19, mIMOParametersPerBandUeiTriggerEventDeterminationR19Constraints); err != nil {
					return err
				}
			}

			if ie.Uei_ModeA_Event1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Uei_ModeA_Event1_r19, mIMOParametersPerBandExtUeiModeAEvent1R19Constraints); err != nil {
					return err
				}
			}

			if ie.Uei_ModeA_Event7_r19 != nil {
				if err := ex.EncodeBitString(*ie.Uei_ModeA_Event7_r19, mIMOParametersPerBandUeiModeAEvent7R19Constraints); err != nil {
					return err
				}
			}

			if ie.Event2ConditionIndication_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Event2ConditionIndication_r19, mIMOParametersPerBandExtEvent2ConditionIndicationR19Constraints); err != nil {
					return err
				}
			}

			if ie.TimeRestriction128Port_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TimeRestriction128Port_r19, mIMOParametersPerBandExtTimeRestriction128PortR19Constraints); err != nil {
					return err
				}
			}

			if ie.GroupScalingFactor_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.GroupScalingFactor_r19, mIMOParametersPerBandExtGroupScalingFactorR19Constraints); err != nil {
					return err
				}
			}

			if ie.Nes_SD_Type1_SP_r19 != nil {
				c := ie.Nes_SD_Type1_SP_r19
				if err := ex.EncodeEnumerated(c.Timeline_r19, mIMOParametersPerBandExtNesSDType1SPR19TimelineR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeBitString(c.NumOfPortCSI_Report_r19, per.FixedSize(10)); err != nil {
					return err
				}
			}

			if ie.Mr_AlwaysReportedType1SP_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mr_AlwaysReportedType1SP_r19, mIMOParametersPerBandExtMrAlwaysReportedType1SPR19Constraints); err != nil {
					return err
				}
			}

			if ie.Mr_AlwaysReported_EType2_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mr_AlwaysReported_EType2_r19, mIMOParametersPerBandExtMrAlwaysReportedEType2R19Constraints); err != nil {
					return err
				}
			}

			if ie.Cjtc_DdReport_r19 != nil {
				c := ie.Cjtc_DdReport_r19
				if err := ex.EncodeEnumerated(c.MinRangeDdInCyclicPrefix_r19, mIMOParametersPerBandExtCjtcDdReportR19MinRangeDdInCyclicPrefixR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxResolutionDd_r19, mIMOParametersPerBandExtCjtcDdReportR19MaxResolutionDdR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ScalingFactor_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_DdReportProcessing_r19 != nil {
				c := ie.Cjtc_DdReportProcessing_r19
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_Resource_r19, mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourcePerCC_r19, mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_FO_Report_r19 != nil {
				c := ie.Cjtc_FO_Report_r19
				if err := ex.EncodeEnumerated(c.MinRangeFO_r19, mIMOParametersPerBandExtCjtcFOReportR19MinRangeFOR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxResolutionFO_r19, mIMOParametersPerBandExtCjtcFOReportR19MaxResolutionFOR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ScalingFactor_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_FO_ReportProcessing_r19 != nil {
				c := ie.Cjtc_FO_ReportProcessing_r19
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_Resource_r19, mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourcePerCC_r19, mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_PO_ReportWideband_r19 != nil {
				c := ie.Cjtc_PO_ReportWideband_r19
				if err := ex.EncodeEnumerated(c.MaxResolution_r19, mIMOParametersPerBandExtCjtcPOReportWidebandR19MaxResolutionR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ScalingFactor_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxSlotDuration_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_PO_ReportWidebandProcessing_r19 != nil {
				c := ie.Cjtc_PO_ReportWidebandProcessing_r19
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_Configured_r19, mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ConfiguredAcrossCC_r19, mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourcePerCC_r19, mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_PO_ReportSubband_r19 != nil {
				c := ie.Cjtc_PO_ReportSubband_r19
				if err := ex.EncodeEnumerated(c.MaxResolution_r19, mIMOParametersPerBandExtCjtcPOReportSubbandR19MaxResolutionR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MinSubbandSize_r19, mIMOParametersPerBandExtCjtcPOReportSubbandR19MinSubbandSizeR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ScalingFactor_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxSlotDuration_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_DdFO_Report_r19 != nil {
				c := ie.Cjtc_DdFO_Report_r19
				if err := ex.EncodeEnumerated(c.MinRangeDdInCyclicPrefix_r19, mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeDdInCyclicPrefixR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxResolutionDd_r19, mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionDdR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MinRangeFO_r19, mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeFOR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxResolutionFO_r19, mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionFOR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ScalingFactor_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjtc_DdFO_ReportProcessing_r19 != nil {
				c := ie.Cjtc_DdFO_ReportProcessing_r19
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_Resource_r19, mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTRS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourcePerCC_r19, mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberCSI_RS_ResourceAcrossCC_r19, mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ValueX_r19, per.Constrained(1, 2)); err != nil {
					return err
				}
			}

			if ie.Cjt_QCL_PDSCH_SchemeC_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjt_QCL_PDSCH_SchemeC_r19, mIMOParametersPerBandExtCjtQCLPDSCHSchemeCR19Constraints); err != nil {
					return err
				}
			}

			if ie.Cjt_QCL_PDSCH_SchemeD_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjt_QCL_PDSCH_SchemeD_r19, mIMOParametersPerBandExtCjtQCLPDSCHSchemeDR19Constraints); err != nil {
					return err
				}
			}

			if ie.Cjt_QCL_PDSCH_SchemeE_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjt_QCL_PDSCH_SchemeE_r19, mIMOParametersPerBandExtCjtQCLPDSCHSchemeER19Constraints); err != nil {
					return err
				}
			}

			if ie.Linked_CJTC_Dd_EType2CJT_Joint_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Linked_CJTC_Dd_EType2CJT_Joint_r19, mIMOParametersPerBandExtLinkedCJTCDdEType2CJTJointR19Constraints); err != nil {
					return err
				}
			}

			if ie.Linked_CJTC_Dd_EType2CJT_Separate_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Linked_CJTC_Dd_EType2CJT_Separate_r19, mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparateR19Constraints); err != nil {
					return err
				}
			}

			if ie.Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19, mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparatePerStateR19Constraints); err != nil {
					return err
				}
			}

			if ie.TimelineRelax_CJTC_Dd_EType2CJT_r19 != nil {
				c := ie.TimelineRelax_CJTC_Dd_EType2CJT_r19
				mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq := ex.NewSequenceEncoder(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Constraints)
				if err := mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.EncodePreamble([]bool{c.Scs15kHz_r19 != nil, c.Scs30kHz_r19 != nil, c.Scs60kHz_r19 != nil, c.Scs120kHz_r19 != nil, c.Scs480kHz_r19 != nil, c.Scs960kHz_r19 != nil}); err != nil {
					return err
				}
				if c.Scs15kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs15kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs15kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs30kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs30kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs30kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs60kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs60kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs60kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs120kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs120kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs120kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs480kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs480kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs480kHzR19Constraints); err != nil {
						return err
					}
				}
				if c.Scs960kHz_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Scs960kHz_r19), mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs960kHzR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(mIMOParametersPerBandExtNonCodebookCSIRSSRS3TxPUSCHR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19))); err != nil {
					return err
				}
				for i := range ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19 {
					if err := ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19, mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSJointTCIR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19, mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSSeparateTCIR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetPRACH_JointTCI_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPRACH_JointTCI_r19, mIMOParametersPerBandExtPathlossOffsetPRACHJointTCIR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetPRACH_SeparateTCI_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPRACH_SeparateTCI_r19, mIMOParametersPerBandExtPathlossOffsetPRACHSeparateTCIR19Constraints); err != nil {
					return err
				}
			}

			if ie.ExtendedStartBitDCI_2_3_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.ExtendedStartBitDCI_2_3_r19, mIMOParametersPerBandExtExtendedStartBitDCI23R19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSRS_PwrControlAdjust_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSRS_PwrControlAdjust_r19, mIMOParametersPerBandExtTwoSRSPwrControlAdjustR19Constraints); err != nil {
					return err
				}
			}

			if ie.OverlapUL_TransReductionEnh_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.OverlapUL_TransReductionEnh_r19, mIMOParametersPerBandExtOverlapULTransReductionEnhR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetUpdate_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetUpdate_r19, mIMOParametersPerBandExtPathlossOffsetUpdateR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSRS_TPC_DCI_2_3_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSRS_TPC_DCI_2_3_r19, mIMOParametersPerBandExtTwoSRSTPCDCI23R19Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_TPC_CLPC_AdjustmentState_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Srs_TPC_CLPC_AdjustmentState_r19, mIMOParametersPerBandExtSrsTPCCLPCAdjustmentStateR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSRS_DCI_1_1_Separate_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSRS_DCI_1_1_Separate_r19, mIMOParametersPerBandExtTwoSRSDCI11SeparateR19Constraints); err != nil {
					return err
				}
			}

			if ie.TwoSRS_DCI_1_1_Joint_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoSRS_DCI_1_1_Joint_r19, mIMOParametersPerBandExtTwoSRSDCI11JointR19Constraints); err != nil {
					return err
				}
			}

			if ie.PathlossOffsetPHR_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PathlossOffsetPHR_r19, mIMOParametersPerBandExtPathlossOffsetPHRR19Constraints); err != nil {
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
					{Name: "aiml-BM-Case2-v1920", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Aiml_BM_Case2_v1920 != nil}); err != nil {
				return err
			}

			if ie.Aiml_BM_Case2_v1920 != nil {
				c := ie.Aiml_BM_Case2_v1920
				if err := ex.EncodeBitString(c.TimeGap_r19, per.FixedSize(5)); err != nil {
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

func (ie *MIMO_ParametersPerBand) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIMOParametersPerBandConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tci-StatePDSCH: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Tci_StatePDSCH = &struct {
				MaxNumberConfiguredTCI_StatesPerCC *int64
				MaxNumberActiveTCI_PerBWP          *int64
			}{}
			c := ie.Tci_StatePDSCH
			mIMOParametersPerBandTciStatePDSCHSeq := d.NewSequenceDecoder(mIMOParametersPerBandTciStatePDSCHConstraints)
			if err := mIMOParametersPerBandTciStatePDSCHSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandTciStatePDSCHSeq.IsComponentPresent(0) {
				c.MaxNumberConfiguredTCI_StatesPerCC = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandTciStatePDSCHMaxNumberConfiguredTCIStatesPerCCConstraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberConfiguredTCI_StatesPerCC) = v
			}
			if mIMOParametersPerBandTciStatePDSCHSeq.IsComponentPresent(1) {
				c.MaxNumberActiveTCI_PerBWP = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandTciStatePDSCHMaxNumberActiveTCIPerBWPConstraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberActiveTCI_PerBWP) = v
			}
		}
	}

	// 4. additionalActiveTCI-StatePDCCH: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandAdditionalActiveTCIStatePDCCHConstraints)
			if err != nil {
				return err
			}
			ie.AdditionalActiveTCI_StatePDCCH = &idx
		}
	}

	// 5. pusch-TransCoherence: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandPuschTransCoherenceConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_TransCoherence = &idx
		}
	}

	// 6. beamCorrespondenceWithoutUL-BeamSweeping: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandBeamCorrespondenceWithoutULBeamSweepingConstraints)
			if err != nil {
				return err
			}
			ie.BeamCorrespondenceWithoutUL_BeamSweeping = &idx
		}
	}

	// 7. periodicBeamReport: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandPeriodicBeamReportConstraints)
			if err != nil {
				return err
			}
			ie.PeriodicBeamReport = &idx
		}
	}

	// 8. aperiodicBeamReport: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandAperiodicBeamReportConstraints)
			if err != nil {
				return err
			}
			ie.AperiodicBeamReport = &idx
		}
	}

	// 9. sp-BeamReportPUCCH: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandSpBeamReportPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.Sp_BeamReportPUCCH = &idx
		}
	}

	// 10. sp-BeamReportPUSCH: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandSpBeamReportPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.Sp_BeamReportPUSCH = &idx
		}
	}

	// 11. dummy1: ref
	{
		if seq.IsComponentPresent(8) {
			ie.Dummy1 = new(DummyG)
			if err := ie.Dummy1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. maxNumberRxBeam: integer
	{
		if seq.IsComponentPresent(9) {
			v, err := d.DecodeInteger(mIMOParametersPerBandMaxNumberRxBeamConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberRxBeam = &v
		}
	}

	// 13. maxNumberRxTxBeamSwitchDL: sequence
	{
		if seq.IsComponentPresent(10) {
			ie.MaxNumberRxTxBeamSwitchDL = &struct {
				Scs_15kHz  *int64
				Scs_30kHz  *int64
				Scs_60kHz  *int64
				Scs_120kHz *int64
				Scs_240kHz *int64
			}{}
			c := ie.MaxNumberRxTxBeamSwitchDL
			mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq := d.NewSequenceDecoder(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLConstraints)
			if err := mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs15kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz) = v
			}
			if mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs30kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz) = v
			}
			if mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
			if mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLSeq.IsComponentPresent(4) {
				c.Scs_240kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberRxTxBeamSwitchDLScs240kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_240kHz) = v
			}
		}
	}

	// 14. maxNumberNonGroupBeamReporting: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandMaxNumberNonGroupBeamReportingConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberNonGroupBeamReporting = &idx
		}
	}

	// 15. groupBeamReporting: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandGroupBeamReportingConstraints)
			if err != nil {
				return err
			}
			ie.GroupBeamReporting = &idx
		}
	}

	// 16. uplinkBeamManagement: sequence
	{
		if seq.IsComponentPresent(13) {
			ie.UplinkBeamManagement = &struct {
				MaxNumberSRS_ResourcePerSet_BM int64
				MaxNumberSRS_ResourceSet       int64
			}{}
			c := ie.UplinkBeamManagement
			{
				v, err := d.DecodeEnumerated(mIMOParametersPerBandUplinkBeamManagementMaxNumberSRSResourcePerSetBMConstraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_BM = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourceSet = v
			}
		}
	}

	// 17. maxNumberCSI-RS-BFD: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(mIMOParametersPerBandMaxNumberCSIRSBFDConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberCSI_RS_BFD = &v
		}
	}

	// 18. maxNumberSSB-BFD: integer
	{
		if seq.IsComponentPresent(15) {
			v, err := d.DecodeInteger(mIMOParametersPerBandMaxNumberSSBBFDConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberSSB_BFD = &v
		}
	}

	// 19. maxNumberCSI-RS-SSB-CBD: integer
	{
		if seq.IsComponentPresent(16) {
			v, err := d.DecodeInteger(mIMOParametersPerBandMaxNumberCSIRSSSBCBDConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberCSI_RS_SSB_CBD = &v
		}
	}

	// 20. dummy2: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &idx
		}
	}

	// 21. twoPortsPTRS-UL: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandTwoPortsPTRSULConstraints)
			if err != nil {
				return err
			}
			ie.TwoPortsPTRS_UL = &idx
		}
	}

	// 22. dummy5: ref
	{
		if seq.IsComponentPresent(19) {
			ie.Dummy5 = new(SRS_Resources)
			if err := ie.Dummy5.Decode(d); err != nil {
				return err
			}
		}
	}

	// 23. dummy3: integer
	{
		if seq.IsComponentPresent(20) {
			v, err := d.DecodeInteger(mIMOParametersPerBandDummy3Constraints)
			if err != nil {
				return err
			}
			ie.Dummy3 = &v
		}
	}

	// 24. beamReportTiming: sequence
	{
		if seq.IsComponentPresent(21) {
			ie.BeamReportTiming = &struct {
				Scs_15kHz  *int64
				Scs_30kHz  *int64
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.BeamReportTiming
			mIMOParametersPerBandBeamReportTimingSeq := d.NewSequenceDecoder(mIMOParametersPerBandBeamReportTimingConstraints)
			if err := mIMOParametersPerBandBeamReportTimingSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandBeamReportTimingSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandBeamReportTimingScs15kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz) = v
			}
			if mIMOParametersPerBandBeamReportTimingSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandBeamReportTimingScs30kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz) = v
			}
			if mIMOParametersPerBandBeamReportTimingSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandBeamReportTimingScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if mIMOParametersPerBandBeamReportTimingSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(int64)
				v, err := d.DecodeEnumerated(mIMOParametersPerBandBeamReportTimingScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}
	}

	// 25. ptrs-DensityRecommendationSetDL: sequence
	{
		if seq.IsComponentPresent(22) {
			ie.Ptrs_DensityRecommendationSetDL = &struct {
				Scs_15kHz  *PTRS_DensityRecommendationDL
				Scs_30kHz  *PTRS_DensityRecommendationDL
				Scs_60kHz  *PTRS_DensityRecommendationDL
				Scs_120kHz *PTRS_DensityRecommendationDL
			}{}
			c := ie.Ptrs_DensityRecommendationSetDL
			mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq := d.NewSequenceDecoder(mIMOParametersPerBandPtrsDensityRecommendationSetDLConstraints)
			if err := mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(PTRS_DensityRecommendationDL)
				if err := (*c.Scs_15kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(PTRS_DensityRecommendationDL)
				if err := (*c.Scs_30kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(PTRS_DensityRecommendationDL)
				if err := (*c.Scs_60kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetDLSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(PTRS_DensityRecommendationDL)
				if err := (*c.Scs_120kHz).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 26. ptrs-DensityRecommendationSetUL: sequence
	{
		if seq.IsComponentPresent(23) {
			ie.Ptrs_DensityRecommendationSetUL = &struct {
				Scs_15kHz  *PTRS_DensityRecommendationUL
				Scs_30kHz  *PTRS_DensityRecommendationUL
				Scs_60kHz  *PTRS_DensityRecommendationUL
				Scs_120kHz *PTRS_DensityRecommendationUL
			}{}
			c := ie.Ptrs_DensityRecommendationSetUL
			mIMOParametersPerBandPtrsDensityRecommendationSetULSeq := d.NewSequenceDecoder(mIMOParametersPerBandPtrsDensityRecommendationSetULConstraints)
			if err := mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.IsComponentPresent(0) {
				c.Scs_15kHz = new(PTRS_DensityRecommendationUL)
				if err := (*c.Scs_15kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.IsComponentPresent(1) {
				c.Scs_30kHz = new(PTRS_DensityRecommendationUL)
				if err := (*c.Scs_30kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.IsComponentPresent(2) {
				c.Scs_60kHz = new(PTRS_DensityRecommendationUL)
				if err := (*c.Scs_60kHz).Decode(d); err != nil {
					return err
				}
			}
			if mIMOParametersPerBandPtrsDensityRecommendationSetULSeq.IsComponentPresent(3) {
				c.Scs_120kHz = new(PTRS_DensityRecommendationUL)
				if err := (*c.Scs_120kHz).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 27. dummy4: ref
	{
		if seq.IsComponentPresent(24) {
			ie.Dummy4 = new(DummyH)
			if err := ie.Dummy4.Decode(d); err != nil {
				return err
			}
		}
	}

	// 28. aperiodicTRS: enumerated
	{
		if seq.IsComponentPresent(25) {
			idx, err := d.DecodeEnumerated(mIMOParametersPerBandAperiodicTRSConstraints)
			if err != nil {
				return err
			}
			ie.AperiodicTRS = &idx
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
				{Name: "dummy6", Optional: true},
				{Name: "beamManagementSSB-CSI-RS", Optional: true},
				{Name: "beamSwitchTiming", Optional: true},
				{Name: "codebookParameters", Optional: true},
				{Name: "csi-RS-IM-ReceptionForFeedback", Optional: true},
				{Name: "csi-RS-ProcFrameworkForSRS", Optional: true},
				{Name: "csi-ReportFramework", Optional: true},
				{Name: "csi-RS-ForTracking", Optional: true},
				{Name: "srs-AssocCSI-RS", Optional: true},
				{Name: "spatialRelations", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtDummy6Constraints)
			if err != nil {
				return err
			}
			ie.Dummy6 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.BeamManagementSSB_CSI_RS = new(BeamManagementSSB_CSI_RS)
			if err := ie.BeamManagementSSB_CSI_RS.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.BeamSwitchTiming = &struct {
				Scs_60kHz  *int64
				Scs_120kHz *int64
			}{}
			c := ie.BeamSwitchTiming
			mIMOParametersPerBandExtBeamSwitchTimingSeq := dx.NewSequenceDecoder(mIMOParametersPerBandExtBeamSwitchTimingConstraints)
			if err := mIMOParametersPerBandExtBeamSwitchTimingSeq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtBeamSwitchTimingSeq.IsComponentPresent(0) {
				c.Scs_60kHz = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingScs60kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz) = v
			}
			if mIMOParametersPerBandExtBeamSwitchTimingSeq.IsComponentPresent(1) {
				c.Scs_120kHz = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingScs120kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz) = v
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.CodebookParameters = new(CodebookParameters)
			if err := ie.CodebookParameters.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Csi_RS_IM_ReceptionForFeedback = new(CSI_RS_IM_ReceptionForFeedback)
			if err := ie.Csi_RS_IM_ReceptionForFeedback.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Csi_RS_ProcFrameworkForSRS = new(CSI_RS_ProcFrameworkForSRS)
			if err := ie.Csi_RS_ProcFrameworkForSRS.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Csi_ReportFramework = new(CSI_ReportFramework)
			if err := ie.Csi_ReportFramework.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Csi_RS_ForTracking = new(CSI_RS_ForTracking)
			if err := ie.Csi_RS_ForTracking.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtSrsAssocCSIRSConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Srs_AssocCSI_RS = make([]SupportedCSI_RS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Srs_AssocCSI_RS[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.SpatialRelations = new(SpatialRelations)
			if err := ie.SpatialRelations.Decode(dx); err != nil {
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
				{Name: "defaultQCL-TwoTCI-r16", Optional: true},
				{Name: "codebookParametersPerBand-r16", Optional: true},
				{Name: "simul-SpatialRelationUpdatePUCCHResGroup-r16", Optional: true},
				{Name: "maxNumberSCellBFR-r16", Optional: true},
				{Name: "simultaneousReceptionDiffTypeD-r16", Optional: true},
				{Name: "ssb-csirs-SINR-measurement-r16", Optional: true},
				{Name: "nonGroupSINR-reporting-r16", Optional: true},
				{Name: "groupSINR-reporting-r16", Optional: true},
				{Name: "multiDCI-multiTRP-Parameters-r16", Optional: true},
				{Name: "singleDCI-SDM-scheme-Parameters-r16", Optional: true},
				{Name: "supportFDM-SchemeA-r16", Optional: true},
				{Name: "supportCodeWordSoftCombining-r16", Optional: true},
				{Name: "supportTDM-SchemeA-r16", Optional: true},
				{Name: "supportInter-slotTDM-r16", Optional: true},
				{Name: "lowPAPR-DMRS-PDSCH-r16", Optional: true},
				{Name: "lowPAPR-DMRS-PUSCHwithoutPrecoding-r16", Optional: true},
				{Name: "lowPAPR-DMRS-PUCCH-r16", Optional: true},
				{Name: "lowPAPR-DMRS-PUSCHwithPrecoding-r16", Optional: true},
				{Name: "csi-ReportFrameworkExt-r16", Optional: true},
				{Name: "codebookParametersAddition-r16", Optional: true},
				{Name: "codebookComboParametersAddition-r16", Optional: true},
				{Name: "beamCorrespondenceSSB-based-r16", Optional: true},
				{Name: "beamCorrespondenceCSI-RS-based-r16", Optional: true},
				{Name: "beamSwitchTiming-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtDefaultQCLTwoTCIR16Constraints)
			if err != nil {
				return err
			}
			ie.DefaultQCL_TwoTCI_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CodebookParametersPerBand_r16 = new(CodebookParameters_v1610)
			if err := ie.CodebookParametersPerBand_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSimulSpatialRelationUpdatePUCCHResGroupR16Constraints)
			if err != nil {
				return err
			}
			ie.Simul_SpatialRelationUpdatePUCCHResGroup_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMaxNumberSCellBFRR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberSCellBFR_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSimultaneousReceptionDiffTypeDR16Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousReceptionDiffTypeD_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Ssb_Csirs_SINR_Measurement_r16 = &struct {
				MaxNumberSSB_CSIRS_OneTx_CMR_r16    int64
				MaxNumberCSI_IM_NZP_IMR_Res_r16     int64
				MaxNumberCSIRS_2Tx_Res_r16          int64
				MaxNumberSSB_CSIRS_Res_r16          int64
				MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16 int64
				SupportedCSI_RS_Density_CMR_r16     int64
				MaxNumberAperiodicCSI_RS_Res_r16    int64
				SupportedSINR_Meas_r16              *int64
			}{}
			c := ie.Ssb_Csirs_SINR_Measurement_r16
			mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Constraints)
			if err := mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSOneTxCMRR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSSB_CSIRS_OneTx_CMR_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_IM_NZP_IMR_Res_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIRS2TxResR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSIRS_2Tx_Res_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberSSBCSIRSResR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSSB_CSIRS_Res_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberCSIIMNZPIMRResMemR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_IM_NZP_IMR_Res_Mem_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedCSIRSDensityCMRR16Constraints)
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_Density_CMR_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16MaxNumberAperiodicCSIRSResR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicCSI_RS_Res_r16 = v
			}
			if mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16Seq.IsComponentPresent(7) {
				c.SupportedSINR_Meas_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSsbCsirsSINRMeasurementR16SupportedSINRMeasR16Constraints)
				if err != nil {
					return err
				}
				(*c.SupportedSINR_Meas_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtNonGroupSINRReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.NonGroupSINR_Reporting_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtGroupSINRReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.GroupSINR_Reporting_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.MultiDCI_MultiTRP_Parameters_r16 = &struct {
				OverlapPDSCHsFullyFreqTime_r16       *int64
				OverlapPDSCHsInTimePartiallyFreq_r16 *int64
				OutOfOrderOperationDL_r16            *struct {
					SupportPDCCH_ToPDSCH_r16    *int64
					SupportPDSCH_ToHARQ_ACK_r16 *int64
				}
				OutOfOrderOperationUL_r16          *int64
				SeparateCRS_RateMatching_r16       *int64
				DefaultQCL_PerCORESETPoolIndex_r16 *int64
				MaxNumberActivatedTCI_States_r16   *struct {
					MaxNumberPerCORESET_Pool_r16         int64
					MaxTotalNumberAcrossCORESET_Pool_r16 int64
				}
			}{}
			c := ie.MultiDCI_MultiTRP_Parameters_r16
			mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Constraints)
			if err := mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(0) {
				c.OverlapPDSCHsFullyFreqTime_r16 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				(*c.OverlapPDSCHsFullyFreqTime_r16) = v
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(1) {
				c.OverlapPDSCHsInTimePartiallyFreq_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OverlapPDSCHsInTimePartiallyFreqR16Constraints)
				if err != nil {
					return err
				}
				(*c.OverlapPDSCHsInTimePartiallyFreq_r16) = v
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(2) {
				c.OutOfOrderOperationDL_r16 = &struct {
					SupportPDCCH_ToPDSCH_r16    *int64
					SupportPDSCH_ToHARQ_ACK_r16 *int64
				}{}
				c := (*c.OutOfOrderOperationDL_r16)
				mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Constraints)
				if err := mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq.IsComponentPresent(0) {
					c.SupportPDCCH_ToPDSCH_r16 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDCCHToPDSCHR16Constraints)
					if err != nil {
						return err
					}
					(*c.SupportPDCCH_ToPDSCH_r16) = v
				}
				if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16Seq.IsComponentPresent(1) {
					c.SupportPDSCH_ToHARQ_ACK_r16 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationDLR16SupportPDSCHToHARQACKR16Constraints)
					if err != nil {
						return err
					}
					(*c.SupportPDSCH_ToHARQ_ACK_r16) = v
				}
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(3) {
				c.OutOfOrderOperationUL_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16OutOfOrderOperationULR16Constraints)
				if err != nil {
					return err
				}
				(*c.OutOfOrderOperationUL_r16) = v
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(4) {
				c.SeparateCRS_RateMatching_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16SeparateCRSRateMatchingR16Constraints)
				if err != nil {
					return err
				}
				(*c.SeparateCRS_RateMatching_r16) = v
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(5) {
				c.DefaultQCL_PerCORESETPoolIndex_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16DefaultQCLPerCORESETPoolIndexR16Constraints)
				if err != nil {
					return err
				}
				(*c.DefaultQCL_PerCORESETPoolIndex_r16) = v
			}
			if mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16Seq.IsComponentPresent(6) {
				c.MaxNumberActivatedTCI_States_r16 = &struct {
					MaxNumberPerCORESET_Pool_r16         int64
					MaxTotalNumberAcrossCORESET_Pool_r16 int64
				}{}
				c := (*c.MaxNumberActivatedTCI_States_r16)
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxNumberPerCORESETPoolR16Constraints)
					if err != nil {
						return err
					}
					c.MaxNumberPerCORESET_Pool_r16 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMultiDCIMultiTRPParametersR16MaxNumberActivatedTCIStatesR16MaxTotalNumberAcrossCORESETPoolR16Constraints)
					if err != nil {
						return err
					}
					c.MaxTotalNumberAcrossCORESET_Pool_r16 = v
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.SingleDCI_SDM_Scheme_Parameters_r16 = &struct {
				SupportNewDMRS_Port_r16   *int64
				SupportTwoPortDL_PTRS_r16 *int64
			}{}
			c := ie.SingleDCI_SDM_Scheme_Parameters_r16
			mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Constraints)
			if err := mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq.IsComponentPresent(0) {
				c.SupportNewDMRS_Port_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportNewDMRSPortR16Constraints)
				if err != nil {
					return err
				}
				(*c.SupportNewDMRS_Port_r16) = v
			}
			if mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16Seq.IsComponentPresent(1) {
				c.SupportTwoPortDL_PTRS_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSingleDCISDMSchemeParametersR16SupportTwoPortDLPTRSR16Constraints)
				if err != nil {
					return err
				}
				(*c.SupportTwoPortDL_PTRS_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportFDMSchemeAR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportFDM_SchemeA_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportCodeWordSoftCombiningR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportCodeWordSoftCombining_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportTDMSchemeAR16Constraints)
			if err != nil {
				return err
			}
			ie.SupportTDM_SchemeA_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			ie.SupportInter_SlotTDM_r16 = &struct {
				SupportRepNumPDSCH_TDRA_r16 int64
				MaxTBS_Size_r16             int64
				MaxNumberTCI_States_r16     int64
			}{}
			c := ie.SupportInter_SlotTDM_r16
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportInterSlotTDMR16SupportRepNumPDSCHTDRAR16Constraints)
				if err != nil {
					return err
				}
				c.SupportRepNumPDSCH_TDRA_r16 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportInterSlotTDMR16MaxTBSSizeR16Constraints)
				if err != nil {
					return err
				}
				c.MaxTBS_Size_r16 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberTCI_States_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLowPAPRDMRSPDSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.LowPAPR_DMRS_PDSCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithoutPrecodingR16Constraints)
			if err != nil {
				return err
			}
			ie.LowPAPR_DMRS_PUSCHwithoutPrecoding_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLowPAPRDMRSPUCCHR16Constraints)
			if err != nil {
				return err
			}
			ie.LowPAPR_DMRS_PUCCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLowPAPRDMRSPUSCHwithPrecodingR16Constraints)
			if err != nil {
				return err
			}
			ie.LowPAPR_DMRS_PUSCHwithPrecoding_r16 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			ie.Csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
			if err := ie.Csi_ReportFrameworkExt_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(19) {
			ie.CodebookParametersAddition_r16 = new(CodebookParametersAddition_r16)
			if err := ie.CodebookParametersAddition_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(20) {
			ie.CodebookComboParametersAddition_r16 = new(CodebookComboParametersAddition_r16)
			if err := ie.CodebookComboParametersAddition_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamCorrespondenceSSBBasedR16Constraints)
			if err != nil {
				return err
			}
			ie.BeamCorrespondenceSSB_Based_r16 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamCorrespondenceCSIRSBasedR16Constraints)
			if err != nil {
				return err
			}
			ie.BeamCorrespondenceCSI_RS_Based_r16 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			ie.BeamSwitchTiming_r16 = &struct {
				Scs_60kHz_r16  *int64
				Scs_120kHz_r16 *int64
			}{}
			c := ie.BeamSwitchTiming_r16
			mIMOParametersPerBandExtBeamSwitchTimingR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtBeamSwitchTimingR16Constraints)
			if err := mIMOParametersPerBandExtBeamSwitchTimingR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtBeamSwitchTimingR16Seq.IsComponentPresent(0) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingR16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
			if mIMOParametersPerBandExtBeamSwitchTimingR16Seq.IsComponentPresent(1) {
				c.Scs_120kHz_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingR16Scs120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r16) = v
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
				{Name: "semi-PersistentL1-SINR-Report-PUCCH-r16", Optional: true},
				{Name: "semi-PersistentL1-SINR-Report-PUSCH-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Semi_PersistentL1_SINR_Report_PUCCH_r16 = &struct {
				SupportReportFormat1_2OFDM_Syms_r16  *int64
				SupportReportFormat4_14OFDM_Syms_r16 *int64
			}{}
			c := ie.Semi_PersistentL1_SINR_Report_PUCCH_r16
			mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Constraints)
			if err := mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq.IsComponentPresent(0) {
				c.SupportReportFormat1_2OFDM_Syms_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat12OFDMSymsR16Constraints)
				if err != nil {
					return err
				}
				(*c.SupportReportFormat1_2OFDM_Syms_r16) = v
			}
			if mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16Seq.IsComponentPresent(1) {
				c.SupportReportFormat4_14OFDM_Syms_r16 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSemiPersistentL1SINRReportPUCCHR16SupportReportFormat414OFDMSymsR16Constraints)
				if err != nil {
					return err
				}
				(*c.SupportReportFormat4_14OFDM_Syms_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSemiPersistentL1SINRReportPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.Semi_PersistentL1_SINR_Report_PUSCH_r16 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "spatialRelations-v1640", Optional: true},
				{Name: "support64CandidateBeamRS-BFR-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SpatialRelations_v1640 = &struct{ MaxNumberConfiguredSpatialRelations_v1640 int64 }{}
			c := ie.SpatialRelations_v1640
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSpatialRelationsV1640MaxNumberConfiguredSpatialRelationsV1640Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfiguredSpatialRelations_v1640 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupport64CandidateBeamRSBFRR16Constraints)
			if err != nil {
				return err
			}
			ie.Support64CandidateBeamRS_BFR_r16 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxMIMO-LayersForMulti-DCI-mTRP-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMaxMIMOLayersForMultiDCIMTRPR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxMIMO_LayersForMulti_DCI_MTRP_r16 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportedSINR-meas-v1670", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(mIMOParametersPerBandSupportedSINRMeasV1670Constraints)
			if err != nil {
				return err
			}
			ie.SupportedSINR_Meas_v1670 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srs-increasedRepetition-r17", Optional: true},
				{Name: "srs-partialFrequencySounding-r17", Optional: true},
				{Name: "srs-startRB-locationHoppingPartial-r17", Optional: true},
				{Name: "srs-combEight-r17", Optional: true},
				{Name: "codebookParametersfetype2-r17", Optional: true},
				{Name: "mTRP-PUSCH-twoCSI-RS-r17", Optional: true},
				{Name: "mTRP-PUCCH-InterSlot-r17", Optional: true},
				{Name: "mTRP-PUCCH-CyclicMapping-r17", Optional: true},
				{Name: "mTRP-PUCCH-SecondTPC-r17", Optional: true},
				{Name: "mTRP-BFR-twoBFD-RS-Set-r17", Optional: true},
				{Name: "mTRP-BFR-PUCCH-SR-perCG-r17", Optional: true},
				{Name: "mTRP-BFR-association-PUCCH-SR-r17", Optional: true},
				{Name: "sfn-SimulTwoTCI-AcrossMultiCC-r17", Optional: true},
				{Name: "sfn-DefaultDL-BeamSetup-r17", Optional: true},
				{Name: "sfn-DefaultUL-BeamSetup-r17", Optional: true},
				{Name: "srs-TriggeringOffset-r17", Optional: true},
				{Name: "srs-TriggeringDCI-r17", Optional: true},
				{Name: "codebookComboParameterMixedType-r17", Optional: true},
				{Name: "unifiedJointTCI-r17", Optional: true},
				{Name: "unifiedJointTCI-multiMAC-CE-r17", Optional: true},
				{Name: "unifiedJointTCI-perBWP-CA-r17", Optional: true},
				{Name: "unifiedJointTCI-ListSharingCA-r17", Optional: true},
				{Name: "unifiedJointTCI-commonMultiCC-r17", Optional: true},
				{Name: "unifiedJointTCI-BeamAlignDLRS-r17", Optional: true},
				{Name: "unifiedJointTCI-PC-association-r17", Optional: true},
				{Name: "unifiedJointTCI-Legacy-r17", Optional: true},
				{Name: "unifiedJointTCI-Legacy-SRS-r17", Optional: true},
				{Name: "unifiedJointTCI-Legacy-CORESET0-r17", Optional: true},
				{Name: "unifiedJointTCI-SCellBFR-r17", Optional: true},
				{Name: "unifiedJointTCI-InterCell-r17", Optional: true},
				{Name: "unifiedSeparateTCI-r17", Optional: true},
				{Name: "unifiedSeparateTCI-multiMAC-CE-r17", Optional: true},
				{Name: "unifiedSeparateTCI-perBWP-CA-r17", Optional: true},
				{Name: "unifiedSeparateTCI-ListSharingCA-r17", Optional: true},
				{Name: "unifiedSeparateTCI-commonMultiCC-r17", Optional: true},
				{Name: "unifiedSeparateTCI-InterCell-r17", Optional: true},
				{Name: "unifiedJointTCI-mTRP-InterCell-BM-r17", Optional: true},
				{Name: "mpe-Mitigation-r17", Optional: true},
				{Name: "srs-PortReport-r17", Optional: true},
				{Name: "mTRP-PDCCH-individual-r17", Optional: true},
				{Name: "mTRP-PDCCH-anySpan-3Symbols-r17", Optional: true},
				{Name: "mTRP-PDCCH-TwoQCL-TypeD-r17", Optional: true},
				{Name: "mTRP-PUSCH-CSI-RS-r17", Optional: true},
				{Name: "mTRP-PUSCH-cyclicMapping-r17", Optional: true},
				{Name: "mTRP-PUSCH-secondTPC-r17", Optional: true},
				{Name: "mTRP-PUSCH-twoPHR-Reporting-r17", Optional: true},
				{Name: "mTRP-PUSCH-A-CSI-r17", Optional: true},
				{Name: "mTRP-PUSCH-SP-CSI-r17", Optional: true},
				{Name: "mTRP-PUSCH-CG-r17", Optional: true},
				{Name: "mTRP-PUCCH-MAC-CE-r17", Optional: true},
				{Name: "mTRP-PUCCH-maxNum-PC-FR1-r17", Optional: true},
				{Name: "mTRP-inter-Cell-r17", Optional: true},
				{Name: "mTRP-GroupBasedL1-RSRP-r17", Optional: true},
				{Name: "mTRP-BFD-RS-MAC-CE-r17", Optional: true},
				{Name: "mTRP-CSI-EnhancementPerBand-r17", Optional: true},
				{Name: "codebookComboParameterMultiTRP-r17", Optional: true},
				{Name: "mTRP-CSI-additionalCSI-r17", Optional: true},
				{Name: "mTRP-CSI-N-Max2-r17", Optional: true},
				{Name: "mTRP-CSI-CMR-r17", Optional: true},
				{Name: "srs-partialFreqSounding-r17", Optional: true},
				{Name: "beamSwitchTiming-v1710", Optional: true},
				{Name: "beamSwitchTiming-r17", Optional: true},
				{Name: "beamReportTiming-v1710", Optional: true},
				{Name: "maxNumberRxTxBeamSwitchDL-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsIncreasedRepetitionR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_IncreasedRepetition_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPartialFrequencySoundingR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PartialFrequencySounding_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsStartRBLocationHoppingPartialR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_StartRB_LocationHoppingPartial_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCombEightR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CombEight_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.CodebookParametersfetype2_r17 = new(CodebookParametersfetype2_r17)
			if err := ie.CodebookParametersfetype2_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHTwoCSIRSR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_TwoCSI_RS_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUCCHInterSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_InterSlot_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUCCHCyclicMappingR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_CyclicMapping_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUCCHSecondTPCR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_SecondTPC_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			ie.MTRP_BFR_TwoBFD_RS_Set_r17 = &struct {
				MaxBFD_RS_ResourcesPerSetPerBWP_r17     int64
				MaxBFR_r17                              int64
				MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17 int64
			}{}
			c := ie.MTRP_BFR_TwoBFD_RS_Set_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesPerSetPerBWPR17Constraints)
				if err != nil {
					return err
				}
				c.MaxBFD_RS_ResourcesPerSetPerBWP_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 9))
				if err != nil {
					return err
				}
				c.MaxBFR_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPBFRTwoBFDRSSetR17MaxBFDRSResourcesAcrossSetsPerBWPR17Constraints)
				if err != nil {
					return err
				}
				c.MaxBFD_RS_ResourcesAcrossSetsPerBWP_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPBFRPUCCHSRPerCGR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_BFR_PUCCH_SR_PerCG_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPBFRAssociationPUCCHSRR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_BFR_Association_PUCCH_SR_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSfnSimulTwoTCIAcrossMultiCCR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_SimulTwoTCI_AcrossMultiCC_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSfnDefaultDLBeamSetupR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_DefaultDL_BeamSetup_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSfnDefaultULBeamSetupR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_DefaultUL_BeamSetup_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsTriggeringOffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_TriggeringOffset_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsTriggeringDCIR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_TriggeringDCI_r17 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			ie.CodebookComboParameterMixedType_r17 = new(CodebookComboParameterMixedType_r17)
			if err := ie.CodebookComboParameterMixedType_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(18) {
			ie.UnifiedJointTCI_r17 = &struct {
				MaxConfiguredJointTCI_r17   int64
				MaxActivatedTCIAcrossCC_r17 int64
			}{}
			c := ie.UnifiedJointTCI_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIR17MaxConfiguredJointTCIR17Constraints)
				if err != nil {
					return err
				}
				c.MaxConfiguredJointTCI_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIR17MaxActivatedTCIAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.MaxActivatedTCIAcrossCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(19) {
			ie.UnifiedJointTCI_MultiMAC_CE_r17 = &struct {
				MinBeamApplicationTime_r17 *int64
				MaxNumMAC_CE_PerCC         int64
			}{}
			c := ie.UnifiedJointTCI_MultiMAC_CE_r17
			mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Constraints)
			if err := mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17Seq.IsComponentPresent(0) {
				c.MinBeamApplicationTime_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MinBeamApplicationTimeR17Constraints)
				if err != nil {
					return err
				}
				(*c.MinBeamApplicationTime_r17) = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIMultiMACCER17MaxNumMACCEPerCCConstraints)
				if err != nil {
					return err
				}
				c.MaxNumMAC_CE_PerCC = v
			}
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIPerBWPCAR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_PerBWP_CA_r17 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIListSharingCAR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_ListSharingCA_r17 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCICommonMultiCCR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_CommonMultiCC_r17 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIBeamAlignDLRSR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_BeamAlignDLRS_r17 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIPCAssociationR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_PC_Association_r17 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCILegacyR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_Legacy_r17 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCILegacySRSR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_Legacy_SRS_r17 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCILegacyCORESET0R17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_Legacy_CORESET0_r17 = &v
		}

		if groupSeq.IsComponentPresent(28) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCISCellBFRR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedJointTCI_SCellBFR_r17 = &v
		}

		if groupSeq.IsComponentPresent(29) {
			ie.UnifiedJointTCI_InterCell_r17 = &struct {
				AdditionalMAC_CE_PerCC_r17    int64
				AdditionalMAC_CE_AcrossCC_r17 int64
			}{}
			c := ie.UnifiedJointTCI_InterCell_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEPerCCR17Constraints)
				if err != nil {
					return err
				}
				c.AdditionalMAC_CE_PerCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIInterCellR17AdditionalMACCEAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.AdditionalMAC_CE_AcrossCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(30) {
			ie.UnifiedSeparateTCI_r17 = &struct {
				MaxConfiguredDL_TCI_r17        int64
				MaxConfiguredUL_TCI_r17        int64
				MaxActivatedDL_TCIAcrossCC_r17 int64
				MaxActivatedUL_TCIAcrossCC_r17 int64
			}{}
			c := ie.UnifiedSeparateTCI_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredDLTCIR17Constraints)
				if err != nil {
					return err
				}
				c.MaxConfiguredDL_TCI_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxConfiguredULTCIR17Constraints)
				if err != nil {
					return err
				}
				c.MaxConfiguredUL_TCI_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedDLTCIAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.MaxActivatedDL_TCIAcrossCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIR17MaxActivatedULTCIAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.MaxActivatedUL_TCIAcrossCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(31) {
			ie.UnifiedSeparateTCI_MultiMAC_CE_r17 = &struct {
				MinBeamApplicationTime_r17  int64
				MaxActivatedDL_TCIPerCC_r17 int64
				MaxActivatedUL_TCIPerCC_r17 int64
			}{}
			c := ie.UnifiedSeparateTCI_MultiMAC_CE_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIMultiMACCER17MinBeamApplicationTimeR17Constraints)
				if err != nil {
					return err
				}
				c.MinBeamApplicationTime_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxActivatedDL_TCIPerCC_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxActivatedUL_TCIPerCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(32) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIPerBWPCAR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedSeparateTCI_PerBWP_CA_r17 = &v
		}

		if groupSeq.IsComponentPresent(33) {
			ie.UnifiedSeparateTCI_ListSharingCA_r17 = &struct {
				MaxNumListDL_TCI_r17 *int64
				MaxNumListUL_TCI_r17 *int64
			}{}
			c := ie.UnifiedSeparateTCI_ListSharingCA_r17
			mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Constraints)
			if err := mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq.IsComponentPresent(0) {
				c.MaxNumListDL_TCI_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListDLTCIR17Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumListDL_TCI_r17) = v
			}
			if mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17Seq.IsComponentPresent(1) {
				c.MaxNumListUL_TCI_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIListSharingCAR17MaxNumListULTCIR17Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumListUL_TCI_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(34) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCICommonMultiCCR17Constraints)
			if err != nil {
				return err
			}
			ie.UnifiedSeparateTCI_CommonMultiCC_r17 = &v
		}

		if groupSeq.IsComponentPresent(35) {
			ie.UnifiedSeparateTCI_InterCell_r17 = &struct {
				K_DL_PerCC_r17    int64
				K_UL_PerCC_r17    int64
				K_DL_AcrossCC_r17 int64
				K_UL_AcrossCC_r17 int64
			}{}
			c := ie.UnifiedSeparateTCI_InterCell_r17
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLPerCCR17Constraints)
				if err != nil {
					return err
				}
				c.K_DL_PerCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULPerCCR17Constraints)
				if err != nil {
					return err
				}
				c.K_UL_PerCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KDLAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.K_DL_AcrossCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedSeparateTCIInterCellR17KULAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.K_UL_AcrossCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(36) {
			ie.UnifiedJointTCI_MTRP_InterCell_BM_r17 = &struct {
				MaxNumAdditionalPCI_L1_RSRP_r17        int64
				MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17 int64
			}{}
			c := ie.UnifiedJointTCI_MTRP_InterCell_BM_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 7))
				if err != nil {
					return err
				}
				c.MaxNumAdditionalPCI_L1_RSRP_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUnifiedJointTCIMTRPInterCellBMR17MaxNumSSBResourceL1RSRPAcrossCCR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumSSB_ResourceL1_RSRP_AcrossCC_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(37) {
			ie.Mpe_Mitigation_r17 = &struct {
				MaxNumP_MPR_RI_Pairs_r17 int64
				MaxNumConfRS_r17         int64
			}{}
			c := ie.Mpe_Mitigation_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumP_MPR_RI_Pairs_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMpeMitigationR17MaxNumConfRSR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumConfRS_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(38) {
			ie.Srs_PortReport_r17 = &struct {
				CapVal1_r17 *int64
				CapVal2_r17 *int64
				CapVal3_r17 *int64
				CapVal4_r17 *int64
			}{}
			c := ie.Srs_PortReport_r17
			mIMOParametersPerBandExtSrsPortReportR17Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtSrsPortReportR17Constraints)
			if err := mIMOParametersPerBandExtSrsPortReportR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtSrsPortReportR17Seq.IsComponentPresent(0) {
				c.CapVal1_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPortReportR17CapVal1R17Constraints)
				if err != nil {
					return err
				}
				(*c.CapVal1_r17) = v
			}
			if mIMOParametersPerBandExtSrsPortReportR17Seq.IsComponentPresent(1) {
				c.CapVal2_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPortReportR17CapVal2R17Constraints)
				if err != nil {
					return err
				}
				(*c.CapVal2_r17) = v
			}
			if mIMOParametersPerBandExtSrsPortReportR17Seq.IsComponentPresent(2) {
				c.CapVal3_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPortReportR17CapVal3R17Constraints)
				if err != nil {
					return err
				}
				(*c.CapVal3_r17) = v
			}
			if mIMOParametersPerBandExtSrsPortReportR17Seq.IsComponentPresent(3) {
				c.CapVal4_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPortReportR17CapVal4R17Constraints)
				if err != nil {
					return err
				}
				(*c.CapVal4_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(39) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPDCCHIndividualR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PDCCH_Individual_r17 = &v
		}

		if groupSeq.IsComponentPresent(40) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPDCCHAnySpan3SymbolsR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PDCCH_AnySpan_3Symbols_r17 = &v
		}

		if groupSeq.IsComponentPresent(41) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPDCCHTwoQCLTypeDR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PDCCH_TwoQCL_TypeD_r17 = &v
		}

		if groupSeq.IsComponentPresent(42) {
			ie.MTRP_PUSCH_CSI_RS_r17 = &struct {
				MaxNumPeriodicSRS_r17          int64
				MaxNumAperiodicSRS_r17         int64
				MaxNumSP_SRS_r17               int64
				NumSRS_ResourcePerCC_r17       int64
				NumSRS_ResourceNonCodebook_r17 int64
			}{}
			c := ie.MTRP_PUSCH_CSI_RS_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumPeriodicSRS_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumAperiodicSRS_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.MaxNumSP_SRS_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.NumSRS_ResourcePerCC_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.NumSRS_ResourceNonCodebook_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(43) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHCyclicMappingR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_CyclicMapping_r17 = &v
		}

		if groupSeq.IsComponentPresent(44) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHSecondTPCR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_SecondTPC_r17 = &v
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHTwoPHRReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_TwoPHR_Reporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHACSIR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_A_CSI_r17 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHSPCSIR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_SP_CSI_r17 = &v
		}

		if groupSeq.IsComponentPresent(48) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUSCHCGR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUSCH_CG_r17 = &v
		}

		if groupSeq.IsComponentPresent(49) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPPUCCHMACCER17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_MAC_CE_r17 = &v
		}

		if groupSeq.IsComponentPresent(50) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandMTRPPUCCHMaxNumPCFR1R17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_PUCCH_MaxNum_PC_FR1_r17 = &v
		}

		if groupSeq.IsComponentPresent(51) {
			ie.MTRP_Inter_Cell_r17 = &struct {
				MaxNumAdditionalPCI_Case1_r17 int64
				MaxNumAdditionalPCI_Case2_r17 int64
			}{}
			c := ie.MTRP_Inter_Cell_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 7))
				if err != nil {
					return err
				}
				c.MaxNumAdditionalPCI_Case1_r17 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.MaxNumAdditionalPCI_Case2_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(52) {
			ie.MTRP_GroupBasedL1_RSRP_r17 = &struct {
				MaxNumBeamGroups_r17    int64
				MaxNumRS_WithinSlot_r17 int64
				MaxNumRS_AcrossSlot_r17 int64
			}{}
			c := ie.MTRP_GroupBasedL1_RSRP_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumBeamGroups_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSWithinSlotR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumRS_WithinSlot_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPGroupBasedL1RSRPR17MaxNumRSAcrossSlotR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumRS_AcrossSlot_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(53) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPBFDRSMACCER17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_BFD_RS_MAC_CE_r17 = &v
		}

		if groupSeq.IsComponentPresent(54) {
			ie.MTRP_CSI_EnhancementPerBand_r17 = &struct {
				MaxNumNZP_CSI_RS_r17        int64
				CSI_Report_Mode_r17         int64
				SupportedComboAcrossCCs_r17 []CSI_MultiTRP_SupportedCombinations_r17
				CodebookModeNCJT_r17        int64
			}{}
			c := ie.MTRP_CSI_EnhancementPerBand_r17
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumNZP_CSI_RS_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CSIReportModeR17Constraints)
				if err != nil {
					return err
				}
				c.CSI_Report_Mode_r17 = v
			}
			{
				seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17SupportedComboAcrossCCsR17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedComboAcrossCCs_r17 = make([]CSI_MultiTRP_SupportedCombinations_r17, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedComboAcrossCCs_r17[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSIEnhancementPerBandR17CodebookModeNCJTR17Constraints)
				if err != nil {
					return err
				}
				c.CodebookModeNCJT_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(55) {
			ie.CodebookComboParameterMultiTRP_r17 = new(CodebookComboParameterMultiTRP_r17)
			if err := ie.CodebookComboParameterMultiTRP_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(56) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSIAdditionalCSIR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_CSI_AdditionalCSI_r17 = &v
		}

		if groupSeq.IsComponentPresent(57) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSINMax2R17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_CSI_N_Max2_r17 = &v
		}

		if groupSeq.IsComponentPresent(58) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSICMRR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_CSI_CMR_r17 = &v
		}

		if groupSeq.IsComponentPresent(59) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPartialFreqSoundingR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PartialFreqSounding_r17 = &v
		}

		if groupSeq.IsComponentPresent(60) {
			ie.BeamSwitchTiming_v1710 = &struct {
				Scs_480kHz *int64
				Scs_960kHz *int64
			}{}
			c := ie.BeamSwitchTiming_v1710
			mIMOParametersPerBandExtBeamSwitchTimingV1710Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtBeamSwitchTimingV1710Constraints)
			if err := mIMOParametersPerBandExtBeamSwitchTimingV1710Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtBeamSwitchTimingV1710Seq.IsComponentPresent(0) {
				c.Scs_480kHz = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingV1710Scs480kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz) = v
			}
			if mIMOParametersPerBandExtBeamSwitchTimingV1710Seq.IsComponentPresent(1) {
				c.Scs_960kHz = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingV1710Scs960kHzConstraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz) = v
			}
		}

		if groupSeq.IsComponentPresent(61) {
			ie.BeamSwitchTiming_r17 = &struct {
				Scs_480kHz_r17 *int64
				Scs_960kHz_r17 *int64
			}{}
			c := ie.BeamSwitchTiming_r17
			mIMOParametersPerBandExtBeamSwitchTimingR17Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtBeamSwitchTimingR17Constraints)
			if err := mIMOParametersPerBandExtBeamSwitchTimingR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtBeamSwitchTimingR17Seq.IsComponentPresent(0) {
				c.Scs_480kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingR17Scs480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r17) = v
			}
			if mIMOParametersPerBandExtBeamSwitchTimingR17Seq.IsComponentPresent(1) {
				c.Scs_960kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamSwitchTimingR17Scs960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(62) {
			ie.BeamReportTiming_v1710 = &struct {
				Scs_480kHz_r17 *int64
				Scs_960kHz_r17 *int64
			}{}
			c := ie.BeamReportTiming_v1710
			mIMOParametersPerBandExtBeamReportTimingV1710Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtBeamReportTimingV1710Constraints)
			if err := mIMOParametersPerBandExtBeamReportTimingV1710Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtBeamReportTimingV1710Seq.IsComponentPresent(0) {
				c.Scs_480kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamReportTimingV1710Scs480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r17) = v
			}
			if mIMOParametersPerBandExtBeamReportTimingV1710Seq.IsComponentPresent(1) {
				c.Scs_960kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtBeamReportTimingV1710Scs960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(63) {
			ie.MaxNumberRxTxBeamSwitchDL_v1710 = &struct {
				Scs_480kHz_r17 *int64
				Scs_960kHz_r17 *int64
			}{}
			c := ie.MaxNumberRxTxBeamSwitchDL_v1710
			mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Constraints)
			if err := mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq.IsComponentPresent(0) {
				c.Scs_480kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs480kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_480kHz_r17) = v
			}
			if mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Seq.IsComponentPresent(1) {
				c.Scs_960kHz_r17 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMaxNumberRxTxBeamSwitchDLV1710Scs960kHzR17Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_960kHz_r17) = v
			}
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "srs-PortReportSP-AP-r17", Optional: true},
				{Name: "maxNumberRxBeam-v1720", Optional: true},
				{Name: "sfn-ImplicitRS-twoTCI-r17", Optional: true},
				{Name: "sfn-QCL-TypeD-Collision-twoTCI-r17", Optional: true},
				{Name: "mTRP-CSI-numCPU-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsPortReportSPAPR17Constraints)
			if err != nil {
				return err
			}
			ie.Srs_PortReportSP_AP_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandMaxNumberRxBeamV1720Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberRxBeam_v1720 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSfnImplicitRSTwoTCIR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_ImplicitRS_TwoTCI_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSfnQCLTypeDCollisionTwoTCIR17Constraints)
			if err != nil {
				return err
			}
			ie.Sfn_QCL_TypeD_Collision_TwoTCI_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMTRPCSINumCPUR17Constraints)
			if err != nil {
				return err
			}
			ie.MTRP_CSI_NumCPU_r17 = &v
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "supportRepNumPDSCH-TDRA-DCI-1-2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSupportRepNumPDSCHTDRADCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.SupportRepNumPDSCH_TDRA_DCI_1_2_r17 = &v
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "codebookParametersetype2DopplerCSI-r18", Optional: true},
				{Name: "codebookParametersfetype2DopplerCSI-r18", Optional: true},
				{Name: "codebookParametersetype2CJT-r18", Optional: true},
				{Name: "codebookParametersfetype2CJT-r18", Optional: true},
				{Name: "codebookComboParametersCJT-r18", Optional: true},
				{Name: "codebookParametersHARQ-ACK-PUSCH-r18", Optional: true},
				{Name: "tci-JointTCI-UpdateSingleActiveTCI-PerCC-r18", Optional: true},
				{Name: "tci-JointTCI-UpdateMultiActiveTCI-PerCC-r18", Optional: true},
				{Name: "tci-SelectionDCI-r18", Optional: true},
				{Name: "tci-SeparateTCI-UpdateSingleActiveTCI-PerCC-r18", Optional: true},
				{Name: "tci-SeparateTCI-UpdateMultiActiveTCI-PerCC-r18", Optional: true},
				{Name: "tci-SelectionAperiodicCSI-RS-r18", Optional: true},
				{Name: "tci-SelectionAperiodicCSI-RS-M-DCI-r18", Optional: true},
				{Name: "twoTCI-StatePDSCH-CJT-TxScheme-r18", Optional: true},
				{Name: "tci-JointTCI-UpdateSingleActiveTCI-PerCC-PerCORESET-r18", Optional: true},
				{Name: "tci-JointTCI-UpdateMultiActiveTCI-PerCC-PerCORESET-r18", Optional: true},
				{Name: "tci-TRP-BFR-r18", Optional: true},
				{Name: "tci-SeparateTCI-UpdateSingleActiveTCI-PerCC-PerCORESET-r18", Optional: true},
				{Name: "tci-SeparateTCI-UpdateMultiActiveTCI-PerCC-PerCORESET-r18", Optional: true},
				{Name: "commonTCI-SingleDCI-r18", Optional: true},
				{Name: "commonTCI-MultiDCI-r18", Optional: true},
				{Name: "twoPHR-Reporting-r18", Optional: true},
				{Name: "spCell-TAG-Ind-r18", Optional: true},
				{Name: "interCellCrossTRP-PDCCH-OrderCFRA-r18", Optional: true},
				{Name: "intraCellCrossTRP-PDCCH-OrderCFRA-r18", Optional: true},
				{Name: "overlapUL-TransReduction-r18", Optional: true},
				{Name: "maxPeriodicityCMR-r18", Optional: true},
				{Name: "tdcp-Report-r18", Optional: true},
				{Name: "tdcp-Resource-r18", Optional: true},
				{Name: "timelineRelax-CJT-CSI-r18", Optional: true},
				{Name: "jointConfigDMRSPortDynamicSwitching-r18", Optional: true},
				{Name: "srs-combOffsetHopping-r18", Optional: true},
				{Name: "srs-combOffsetInTime-r18", Optional: true},
				{Name: "srs-combOffsetCombinedGroupSequence-r18", Optional: true},
				{Name: "srs-combOffsetHoppingWithinSubset-r18", Optional: true},
				{Name: "srs-cyclicShiftHopping-r18", Optional: true},
				{Name: "srs-cyclicShiftHoppingSmallGranularity-r18", Optional: true},
				{Name: "srs-cyclicShiftCombinedGroupSequence-r18", Optional: true},
				{Name: "cyclicShiftHoppingWithinSubset-r18", Optional: true},
				{Name: "srs-cyclicShiftCombinedCombOffset-r18", Optional: true},
				{Name: "pusch-CB-2PTRS-SingleDCI-STx2P-SDM-r18", Optional: true},
				{Name: "pusch-NonCB-2PTRS-SingleDCI-STx2P-SDM-r18", Optional: true},
				{Name: "pusch-NonCB-SingleDCI-STx2P-SDM-CSI-RS-SRS-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-Multi-DCI-STx2P-CSI-RS-Resource-r18", Optional: true},
				{Name: "dmrs-PortEntrySingleDCI-SDM-r18", Optional: true},
				{Name: "pusch-CB-2PTRS-SingleDCI-STx2P-SFN-r18", Optional: true},
				{Name: "pusch-NonCB-2PTRS-SingleDCI-STx2P-SFN-r18", Optional: true},
				{Name: "pusch-NonCB-SingleDCI-STx2P-SFN-CSI-RS-SRS-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-FullTimeFullFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-FullTimePartialFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimeFullFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimePartialFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-PartialTimeNonFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-CG-CG-r18", Optional: true},
				{Name: "twoPUSCH-CB-MultiDCI-STx2P-CG-DG-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-FullTimeFullFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-FullTimePartialFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimeFullFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimePartialFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-PartialTimeNonFreqOverlap-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-CG-CG-r18", Optional: true},
				{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-CG-DG-r18", Optional: true},
				{Name: "pucch-RepetitionDynamicIndicationSFN-r18", Optional: true},
				{Name: "groupBeamReporting-STx2P-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CodebookParametersetype2DopplerCSI_r18 = new(CodebookParametersetype2DopplerCSI_r18)
			if err := ie.CodebookParametersetype2DopplerCSI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CodebookParametersfetype2DopplerCSI_r18 = new(CodebookParametersfetype2DopplerCSI_r18)
			if err := ie.CodebookParametersfetype2DopplerCSI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CodebookParametersetype2CJT_r18 = new(CodebookParametersetype2CJT_r18)
			if err := ie.CodebookParametersetype2CJT_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.CodebookParametersfetype2CJT_r18 = new(CodebookParametersfetype2CJT_r18)
			if err := ie.CodebookParametersfetype2CJT_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.CodebookComboParametersCJT_r18 = new(CodebookComboParametersCJT_r18)
			if err := ie.CodebookComboParametersCJT_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.CodebookParametersHARQ_ACK_PUSCH_r18 = new(CodebookParametersHARQ_ACK_PUSCH_r18)
			if err := ie.CodebookParametersHARQ_ACK_PUSCH_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18 = &struct {
				MaxNumberConfigJointTCIPerCC_PerBWP_r18 int64
				MaxNumberActiveJointTCI_AcrossCC_r18    int64
			}{}
			c := ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfigJointTCIPerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCR18MaxNumberActiveJointTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberActiveJointTCI_AcrossCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18 = &struct {
				Tci_StateInd_r18                  int64
				MaxNumberActiveJointTCI_PerCC_r18 int64
			}{}
			c := ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateMultiActiveTCIPerCCR18TciStateIndR18Constraints)
				if err != nil {
					return err
				}
				c.Tci_StateInd_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumberActiveJointTCI_PerCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSelectionDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_SelectionDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18 = &struct {
				MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
				MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
				MaxNumActiveDL_TCI_AcrossCC_r18     int64
				MaxNumActiveUL_TCI_AcrossCC_r18     int64
			}{}
			c := ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumConfigULTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveDL_TCI_AcrossCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveUL_TCI_AcrossCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(10) {
			ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18 = &struct {
				MaxNumActiveDL_TCI_AcrossCC_r18 int64
				MaxNumActiveUL_TCI_AcrossCC_r18 int64
			}{}
			c := ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveDLTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveDL_TCI_AcrossCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateMultiActiveTCIPerCCR18MaxNumActiveULTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveUL_TCI_AcrossCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSelectionAperiodicCSIRSR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_SelectionAperiodicCSI_RS_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSelectionAperiodicCSIRSMDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_SelectionAperiodicCSI_RS_M_DCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoTCIStatePDSCHCJTTxSchemeR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoTCI_StatePDSCH_CJT_TxScheme_r18 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 = &struct {
				MTRP_Operation_r18                             int64
				MaxNumberConfigJointTCIPerCC_PerBWP_r18        int64
				MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18 int64
			}{}
			c := ie.Tci_JointTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints)
				if err != nil {
					return err
				}
				c.MTRP_Operation_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberConfigJointTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfigJointTCIPerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciJointTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumberActiveJointTCIAcrossCCPerCORESETR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberActiveJointTCIAcrossCC_PerCORESET_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandTciJointTCIUpdateMultiActiveTCIPerCCPerCORESETR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_JointTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciTRPBFRR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_TRP_BFR_r18 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18 = &struct {
				MTRP_Operation_r18                  int64
				MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
				MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
				MaxNumActiveDL_TCI_AcrossCC_r18     int64
				MaxNumActiveUL_TCI_AcrossCC_r18     int64
			}{}
			c := ie.Tci_SeparateTCI_UpdateSingleActiveTCI_PerCC_PerCORESET_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MTRPOperationR18Constraints)
				if err != nil {
					return err
				}
				c.MTRP_Operation_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigDLTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumConfigULTCIPerCCPerBWPR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveDLTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveDL_TCI_AcrossCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTciSeparateTCIUpdateSingleActiveTCIPerCCPerCORESETR18MaxNumActiveULTCIAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumActiveUL_TCI_AcrossCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(18) {
			ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18 = &struct {
				MaxNumConfigDL_TCI_PerCC_PerBWP_r18 int64
				MaxNumConfigUL_TCI_PerCC_PerBWP_r18 int64
			}{}
			c := ie.Tci_SeparateTCI_UpdateMultiActiveTCI_PerCC_PerCORESET_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumConfigDL_TCI_PerCC_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumConfigUL_TCI_PerCC_PerBWP_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandCommonTCISingleDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.CommonTCI_SingleDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandCommonTCIMultiDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.CommonTCI_MultiDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPHRReportingR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPHR_Reporting_r18 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSpCellTAGIndR18Constraints)
			if err != nil {
				return err
			}
			ie.SpCell_TAG_Ind_r18 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtInterCellCrossTRPPDCCHOrderCFRAR18Constraints)
			if err != nil {
				return err
			}
			ie.InterCellCrossTRP_PDCCH_OrderCFRA_r18 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtIntraCellCrossTRPPDCCHOrderCFRAR18Constraints)
			if err != nil {
				return err
			}
			ie.IntraCellCrossTRP_PDCCH_OrderCFRA_r18 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtOverlapULTransReductionR18Constraints)
			if err != nil {
				return err
			}
			ie.OverlapUL_TransReduction_r18 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMaxPeriodicityCMRR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxPeriodicityCMR_r18 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			ie.Tdcp_Report_r18 = &struct {
				ValueX_r18                  int64
				MaxNumberActiveResource_r18 int64
			}{}
			c := ie.Tdcp_Report_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 32))
				if err != nil {
					return err
				}
				c.MaxNumberActiveResource_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(28) {
			ie.Tdcp_Resource_r18 = &struct {
				MaxNumberConfigPerCC_r18       int64
				MaxNumberConfigAcrossCC_r18    int64
				MaxNumberSimultaneousPerCC_r18 int64
			}{}
			c := ie.Tdcp_Resource_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTdcpResourceR18MaxNumberConfigPerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfigPerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberConfigAcrossCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTdcpResourceR18MaxNumberSimultaneousPerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSimultaneousPerCC_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCSIR18Constraints)
			if err != nil {
				return err
			}
			ie.TimelineRelax_CJT_CSI_r18 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtJointConfigDMRSPortDynamicSwitchingR18Constraints)
			if err != nil {
				return err
			}
			ie.JointConfigDMRSPortDynamicSwitching_r18 = &v
		}

		if groupSeq.IsComponentPresent(31) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCombOffsetHoppingR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CombOffsetHopping_r18 = &v
		}

		if groupSeq.IsComponentPresent(32) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCombOffsetInTimeR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CombOffsetInTime_r18 = &v
		}

		if groupSeq.IsComponentPresent(33) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCombOffsetCombinedGroupSequenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CombOffsetCombinedGroupSequence_r18 = &v
		}

		if groupSeq.IsComponentPresent(34) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCombOffsetHoppingWithinSubsetR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CombOffsetHoppingWithinSubset_r18 = &v
		}

		if groupSeq.IsComponentPresent(35) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCyclicShiftHoppingR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CyclicShiftHopping_r18 = &v
		}

		if groupSeq.IsComponentPresent(36) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCyclicShiftHoppingSmallGranularityR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CyclicShiftHoppingSmallGranularity_r18 = &v
		}

		if groupSeq.IsComponentPresent(37) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCyclicShiftCombinedGroupSequenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CyclicShiftCombinedGroupSequence_r18 = &v
		}

		if groupSeq.IsComponentPresent(38) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCyclicShiftHoppingWithinSubsetR18Constraints)
			if err != nil {
				return err
			}
			ie.CyclicShiftHoppingWithinSubset_r18 = &v
		}

		if groupSeq.IsComponentPresent(39) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsCyclicShiftCombinedCombOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.Srs_CyclicShiftCombinedCombOffset_r18 = &v
		}

		if groupSeq.IsComponentPresent(40) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSDMR18Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SDM_r18 = &v
		}

		if groupSeq.IsComponentPresent(41) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSDMR18Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SDM_r18 = &v
		}

		if groupSeq.IsComponentPresent(42) {
			ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18 = &struct {
				MaxNumberPeriodicSRS_Resource_PerBWP_r18      int64
				MaxNumberAperiodicSRS_Resource_PerBWP_r18     int64
				MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 int64
				ValueY_SRS_ResourceAssociate_r18              int64
				ValueX_CSI_RS_ResourceAssociate_r18           int64
			}{}
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SDM_CSI_RS_SRS_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberPeriodicSRS_Resource_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicSRS_Resource_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.ValueY_SRS_ResourceAssociate_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_CSI_RS_ResourceAssociate_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(43) {
			ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18 = &struct {
				MaxNumberPeriodicSRS_r18       int64
				MaxNumberAperiodicSRS_r18      int64
				MaxNumberSemiPersistentSRS_r18 int64
				SimultaneousSRS_PerCC_r18      int64
				SimultaneousCSI_RS_NonCB_r18   int64
			}{}
			c := ie.TwoPUSCH_NonCB_Multi_DCI_STx2P_CSI_RS_Resource_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberPeriodicSRS_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicSRS_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSemiPersistentSRS_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.SimultaneousSRS_PerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.SimultaneousCSI_RS_NonCB_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(44) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtDmrsPortEntrySingleDCISDMR18Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_PortEntrySingleDCI_SDM_r18 = &v
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPuschCB2PTRSSingleDCISTx2PSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_CB_2PTRS_SingleDCI_STx2P_SFN_r18 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPuschNonCB2PTRSSingleDCISTx2PSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_NonCB_2PTRS_SingleDCI_STx2P_SFN_r18 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18 = &struct {
				MaxNumberPeriodicSRS_Resource_PerBWP_r18      int64
				MaxNumberAperiodicSRS_Resource_PerBWP_r18     int64
				MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 int64
				ValueY_SRS_ResourceAssociate_r18              int64
				ValueX_CSI_RS_ResourceAssociate_r18           int64
			}{}
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SFN_CSI_RS_SRS_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberPeriodicSRS_Resource_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicSRS_Resource_PerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSemiPersistentSRS_ResourcePerBWP_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.ValueY_SRS_ResourceAssociate_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_CSI_RS_ResourceAssociate_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(48) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(49) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(50) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(51) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(52) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(53) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGCGR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_CG_r18 = &v
		}

		if groupSeq.IsComponentPresent(54) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHCBMultiDCISTx2PCGDGR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_CB_MultiDCI_STx2P_CG_DG_r18 = &v
		}

		if groupSeq.IsComponentPresent(55) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimeFullFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimeFullFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(56) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PFullTimePartialFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_FullTimePartialFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(57) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeFullFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeFullFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(58) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimePartialFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimePartialFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(59) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PPartialTimeNonFreqOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_PartialTimeNonFreqOverlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(60) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGCGR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_CG_r18 = &v
		}

		if groupSeq.IsComponentPresent(61) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoPUSCHNonCBMultiDCISTx2PCGDGR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_CG_DG_r18 = &v
		}

		if groupSeq.IsComponentPresent(62) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPucchRepetitionDynamicIndicationSFNR18Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_RepetitionDynamicIndicationSFN_r18 = &v
		}

		if groupSeq.IsComponentPresent(63) {
			ie.GroupBeamReporting_STx2P_r18 = &struct {
				GroupL1_RSRP_Reporting_r18         int64
				MaxNumberBeamGroups_r18            int64
				MaxNumberResWithinSlotAcrossCC_r18 int64
				MaxNumberResAcrossCC_r18           int64
			}{}
			c := ie.GroupBeamReporting_STx2P_r18
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtGroupBeamReportingSTx2PR18GroupL1RSRPReportingR18Constraints)
				if err != nil {
					return err
				}
				c.GroupL1_RSRP_Reporting_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberBeamGroups_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResWithinSlotAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResWithinSlotAcrossCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtGroupBeamReportingSTx2PR18MaxNumberResAcrossCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberResAcrossCC_r18 = v
			}
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "simulConfigDMRS-DCI-1-3-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSimulConfigDMRSDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.SimulConfigDMRS_DCI_1_3_r18 = &v
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "codebookParametersType1SP-SchemeA-r19", Optional: true},
				{Name: "codebookParametersType1SP-SchemeB-r19", Optional: true},
				{Name: "codebookParametersType1MP-r19", Optional: true},
				{Name: "codebookParameterseType2Ext-r19", Optional: true},
				{Name: "codebookParametersfeType2Ext-r19", Optional: true},
				{Name: "codebookParameterseType2DopplerExt-r19", Optional: true},
				{Name: "codebookParametersHybridBF-Type1SP-r19", Optional: true},
				{Name: "codebookParametersHybridBF-eType2-r19", Optional: true},
				{Name: "aiml-CSI-PredictionDoppler-r19", Optional: true},
				{Name: "aiml-CSI-Report-r19", Optional: true},
				{Name: "increasedNumOfReportedRS-r19", Optional: true},
				{Name: "aiml-BM-Case1-r19", Optional: true},
				{Name: "aiml-BM-Case1-PredictedRSRP-r19", Optional: true},
				{Name: "aiml-BM-Case2-r19", Optional: true},
				{Name: "aiml-BM-Case2-PredictedRSRP-r19", Optional: true},
				{Name: "aiml-BM-Monitoring-r19", Optional: true},
				{Name: "aiml-BM-UE-DataCollection-r19", Optional: true},
				{Name: "aiml-CSI-Prediction-r19", Optional: true},
				{Name: "aiml-CSI-PredictionUnitDurationDD-r19", Optional: true},
				{Name: "aiml-CSI-PredictionN4-r19", Optional: true},
				{Name: "aiml-CSI-PredictionUE-DataCollection-r19", Optional: true},
				{Name: "aiml-CSI-PredictionMonitoring-r19", Optional: true},
				{Name: "uei-ModeA-Event2-r19", Optional: true},
				{Name: "uei-ModeB-r19", Optional: true},
				{Name: "uei-TriggerEventDetermination-r19", Optional: true},
				{Name: "uei-ModeA-Event1-r19", Optional: true},
				{Name: "uei-ModeA-Event7-r19", Optional: true},
				{Name: "event2ConditionIndication-r19", Optional: true},
				{Name: "timeRestriction128Port-r19", Optional: true},
				{Name: "groupScalingFactor-r19", Optional: true},
				{Name: "nes-SD-Type1-SP-r19", Optional: true},
				{Name: "mr-AlwaysReportedType1SP-r19", Optional: true},
				{Name: "mr-AlwaysReported-eType2-r19", Optional: true},
				{Name: "cjtc-DdReport-r19", Optional: true},
				{Name: "cjtc-DdReportProcessing-r19", Optional: true},
				{Name: "cjtc-FO-Report-r19", Optional: true},
				{Name: "cjtc-FO-ReportProcessing-r19", Optional: true},
				{Name: "cjtc-PO-ReportWideband-r19", Optional: true},
				{Name: "cjtc-PO-ReportWidebandProcessing-r19", Optional: true},
				{Name: "cjtc-PO-ReportSubband-r19", Optional: true},
				{Name: "cjtc-DdFO-Report-r19", Optional: true},
				{Name: "cjtc-DdFO-ReportProcessing-r19", Optional: true},
				{Name: "cjt-QCL-PDSCH-SchemeC-r19", Optional: true},
				{Name: "cjt-QCL-PDSCH-SchemeD-r19", Optional: true},
				{Name: "cjt-QCL-PDSCH-SchemeE-r19", Optional: true},
				{Name: "linked-CJTC-Dd-eType2CJT-Joint-r19", Optional: true},
				{Name: "linked-CJTC-Dd-eType2CJT-Separate-r19", Optional: true},
				{Name: "linked-CJTC-Dd-eType2CJT-SeparatePerState-r19", Optional: true},
				{Name: "timelineRelax-CJTC-Dd-eType2CJT-r19", Optional: true},
				{Name: "nonCodebook-CSI-RS-SRS-3TxPUSCH-r19", Optional: true},
				{Name: "pathlossOffsetPUCCH-PUSCH-SRS-JointTCI-r19", Optional: true},
				{Name: "pathlossOffsetPUCCH-PUSCH-SRS-SeparateTCI-r19", Optional: true},
				{Name: "pathlossOffsetPRACH-JointTCI-r19", Optional: true},
				{Name: "pathlossOffsetPRACH-SeparateTCI-r19", Optional: true},
				{Name: "extendedStartBitDCI-2-3-r19", Optional: true},
				{Name: "twoSRS-PwrControlAdjust-r19", Optional: true},
				{Name: "overlapUL-TransReductionEnh-r19", Optional: true},
				{Name: "pathlossOffsetUpdate-r19", Optional: true},
				{Name: "twoSRS-TPC-DCI-2-3-r19", Optional: true},
				{Name: "srs-TPC-CLPC-AdjustmentState-r19", Optional: true},
				{Name: "twoSRS-DCI-1-1-Separate-r19", Optional: true},
				{Name: "twoSRS-DCI-1-1-Joint-r19", Optional: true},
				{Name: "pathlossOffsetPHR-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CodebookParametersType1SP_SchemeA_r19 = new(CodebookParametersType1SP_SchemeA_r19)
			if err := ie.CodebookParametersType1SP_SchemeA_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CodebookParametersType1SP_SchemeB_r19 = new(CodebookParametersType1SP_SchemeB_r19)
			if err := ie.CodebookParametersType1SP_SchemeB_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CodebookParametersType1MP_r19 = new(CodebookParametersType1MP_r19)
			if err := ie.CodebookParametersType1MP_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.CodebookParameterseType2Ext_r19 = new(CodebookParameterseType2Ext_r19)
			if err := ie.CodebookParameterseType2Ext_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.CodebookParametersfeType2Ext_r19 = new(CodebookParametersfeType2Ext_r19)
			if err := ie.CodebookParametersfeType2Ext_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.CodebookParameterseType2DopplerExt_r19 = new(CodebookParameterseType2DopplerExt_r19)
			if err := ie.CodebookParameterseType2DopplerExt_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.CodebookParametersHybridBF_Type1SP_r19 = new(CodebookParametersHybridBF_Type1SP_r19)
			if err := ie.CodebookParametersHybridBF_Type1SP_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.CodebookParametersHybridBF_EType2_r19 = new(CodebookParametersHybridBF_EType2_r19)
			if err := ie.CodebookParametersHybridBF_EType2_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Aiml_CSI_PredictionDoppler_r19 = new(CodebookParametersCSI_PredictionDoppler_r19)
			if err := ie.Aiml_CSI_PredictionDoppler_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtAimlCSIReportR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Aiml_CSI_Report_r19 = make([]CPU_PoolInfo_r19, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Aiml_CSI_Report_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtIncreasedNumOfReportedRSR19Constraints)
			if err != nil {
				return err
			}
			ie.IncreasedNumOfReportedRS_r19 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			ie.Aiml_BM_Case1_r19 = &struct {
				MaxNumberOfInferenceReportPerBWP_r19 struct {
					PeriodicReporting_r19       int64
					AperiodicReporting_r19      int64
					SemiPersistentReporting_r19 int64
				}
				MaxNumberOfInferenceReportAcrossAllCC_r19 int64
				MaxNumberOfResourceSetB_r19               int64
				MaxNumberOfResourceSetA_r19               int64
				ResourceTypeSetB_CSI_RS_r19               struct {
					Periodic_r19       *int64
					Aperiodic_r19      *int64
					SemiPersistent_r19 *int64
				}
				InferenceReportType_r19 struct {
					Periodic_r19       *int64
					Aperiodic_r19      *int64
					SemiPersistent_r19 *int64
				}
				SubUseCases_r19                                  int64
				MaxNumberOfPredictedBeamPerReportingInstance_r19 int64
				NumberOfOccupiedCPU_r19                          int64
				NumberOfOccupiedCPUx_r19                         int64
				RelaxationTimelineD_r19                          struct {
					Scs15kHz_r19  int64
					Scs30kHz_r19  int64
					Scs60kHz_r19  int64
					Scs120kHz_r19 int64
					Scs480kHz_r19 int64
					Scs960kHz_r19 int64
				}
				RelaxationTimelineD1_r19 struct {
					Scs15kHz_r19  int64
					Scs30kHz_r19  int64
					Scs60kHz_r19  int64
					Scs120kHz_r19 int64
					Scs480kHz_r19 int64
					Scs960kHz_r19 int64
				}
				OccupiedResourcePool_r19 int64
			}{}
			c := ie.Aiml_BM_Case1_r19
			{
				c := &c.MaxNumberOfInferenceReportPerBWP_r19
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.PeriodicReporting_r19 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.AperiodicReporting_r19 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.SemiPersistentReporting_r19 = v
				}
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfInferenceReportAcrossAllCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetBR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfResourceSetB_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19MaxNumberOfResourceSetAR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfResourceSetA_r19 = v
			}
			{
				c := &c.ResourceTypeSetB_CSI_RS_r19
				mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Constraints)
				if err := mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq.IsComponentPresent(0) {
					c.Periodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Periodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq.IsComponentPresent(1) {
					c.Aperiodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19AperiodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Aperiodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19Seq.IsComponentPresent(2) {
					c.SemiPersistent_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints)
					if err != nil {
						return err
					}
					(*c.SemiPersistent_r19) = v
				}
			}
			{
				c := &c.InferenceReportType_r19
				mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Constraints)
				if err := mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq.IsComponentPresent(0) {
					c.Periodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19PeriodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Periodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq.IsComponentPresent(1) {
					c.Aperiodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19AperiodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Aperiodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19Seq.IsComponentPresent(2) {
					c.SemiPersistent_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19InferenceReportTypeR19SemiPersistentR19Constraints)
					if err != nil {
						return err
					}
					(*c.SemiPersistent_r19) = v
				}
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19SubUseCasesR19Constraints)
				if err != nil {
					return err
				}
				c.SubUseCases_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberOfPredictedBeamPerReportingInstance_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPU_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPUx_r19 = v
			}
			{
				c := &c.RelaxationTimelineD_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs15kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs15kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs30kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs30kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs60kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs60kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs120kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs120kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs480kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs480kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineDR19Scs960kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs960kHz_r19 = v
				}
			}
			{
				c := &c.RelaxationTimelineD1_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs15kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs15kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs30kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs30kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs60kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs60kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs120kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs120kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs480kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs480kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase1R19RelaxationTimelineD1R19Scs960kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs960kHz_r19 = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.OccupiedResourcePool_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandAimlBMCase1PredictedRSRPR19Constraints)
			if err != nil {
				return err
			}
			ie.Aiml_BM_Case1_PredictedRSRP_r19 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			ie.Aiml_BM_Case2_r19 = &struct {
				MaxNumberOfInferenceReportPerBWP_r19 struct {
					PeriodicReporting_r19       int64
					AperiodicReporting_r19      int64
					SemiPersistentReporting_r19 int64
				}
				MaxNumberOfInferenceReportAcrossAllCC_r19 int64
				MaxNumberOfResourceSetB_r19               int64
				MaxNumberOfResourceSetA_r19               int64
				MinNumberOfKBM_SetB_r19                   int64
				ResourceTypeSetB_CSI_RS_r19               struct {
					Periodic_r19       *int64
					SemiPersistent_r19 *int64
				}
				InferenceReportType_r19 struct {
					Periodic_r19       *int64
					Aperiodic_r19      *int64
					SemiPersistent_r19 *int64
				}
				MaxNumberOfPredictedBeamPerPerTimeInstance_r19 int64
				MaxNumberOfPredictedTimeInstance_r19           int64
				MaxTotalNumberOfPredictedBeamPerReport_r19     int64
				Dummy_TimeGap_r19                              int64
				NumberOfOccupiedCPU_r19                        int64
				NumberOfOccupiedCPUx_r19                       int64
				RelaxationTimelineD_r19                        struct {
					Scs15kHz_r19  int64
					Scs30kHz_r19  int64
					Scs60kHz_r19  int64
					Scs120kHz_r19 int64
					Scs480kHz_r19 int64
					Scs960kHz_r19 int64
				}
				RelaxationTimelineD1_r19 struct {
					Scs15kHz_r19  int64
					Scs30kHz_r19  int64
					Scs60kHz_r19  int64
					Scs120kHz_r19 int64
					Scs480kHz_r19 int64
					Scs960kHz_r19 int64
				}
				OccupiedResourcePool_r19 int64
			}{}
			c := ie.Aiml_BM_Case2_r19
			{
				c := &c.MaxNumberOfInferenceReportPerBWP_r19
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.PeriodicReporting_r19 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.AperiodicReporting_r19 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 4))
					if err != nil {
						return err
					}
					c.SemiPersistentReporting_r19 = v
				}
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfInferenceReportAcrossAllCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfInferenceReportAcrossAllCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetBR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfResourceSetB_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfResourceSetAR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfResourceSetA_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MinNumberOfKBMSetBR19Constraints)
				if err != nil {
					return err
				}
				c.MinNumberOfKBM_SetB_r19 = v
			}
			{
				c := &c.ResourceTypeSetB_CSI_RS_r19
				mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Constraints)
				if err := mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq.IsComponentPresent(0) {
					c.Periodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19PeriodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Periodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19Seq.IsComponentPresent(1) {
					c.SemiPersistent_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19ResourceTypeSetBCSIRSR19SemiPersistentR19Constraints)
					if err != nil {
						return err
					}
					(*c.SemiPersistent_r19) = v
				}
			}
			{
				c := &c.InferenceReportType_r19
				mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Constraints)
				if err := mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq.DecodePreamble(); err != nil {
					return err
				}
				if mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq.IsComponentPresent(0) {
					c.Periodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19PeriodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Periodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq.IsComponentPresent(1) {
					c.Aperiodic_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19AperiodicR19Constraints)
					if err != nil {
						return err
					}
					(*c.Aperiodic_r19) = v
				}
				if mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19Seq.IsComponentPresent(2) {
					c.SemiPersistent_r19 = new(int64)
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19InferenceReportTypeR19SemiPersistentR19Constraints)
					if err != nil {
						return err
					}
					(*c.SemiPersistent_r19) = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberOfPredictedBeamPerPerTimeInstance_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MaxNumberOfPredictedTimeInstanceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfPredictedTimeInstance_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19MaxTotalNumberOfPredictedBeamPerReportR19Constraints)
				if err != nil {
					return err
				}
				c.MaxTotalNumberOfPredictedBeamPerReport_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19DummyTimeGapR19Constraints)
				if err != nil {
					return err
				}
				c.Dummy_TimeGap_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPU_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPUx_r19 = v
			}
			{
				c := &c.RelaxationTimelineD_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs15kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs15kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs30kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs30kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs60kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs60kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs120kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs120kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs480kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs480kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineDR19Scs960kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs960kHz_r19 = v
				}
			}
			{
				c := &c.RelaxationTimelineD1_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs15kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs15kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs30kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs30kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs60kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs60kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs120kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs120kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs480kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs480kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2R19RelaxationTimelineD1R19Scs960kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs960kHz_r19 = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.OccupiedResourcePool_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(14) {
			ie.Aiml_BM_Case2_PredictedRSRP_r19 = &struct {
				MaxNumPredictedBeamPerInstance_r19      int64
				MaxNumPredictedTime_r19                 int64
				MaxTotalNumPredictedBeamInOneReport_r19 int64
			}{}
			c := ie.Aiml_BM_Case2_PredictedRSRP_r19
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumPredictedBeamPerInstance_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxNumPredictedTimeR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumPredictedTime_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMCase2PredictedRSRPR19MaxTotalNumPredictedBeamInOneReportR19Constraints)
				if err != nil {
					return err
				}
				c.MaxTotalNumPredictedBeamInOneReport_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			ie.Aiml_BM_Monitoring_r19 = &struct {
				MaxNumTotalResource_r19          int64
				MaxNumReportPerBWP_Periodic_r19  int64
				MaxNumReportPerBWP_Aperiodic_r19 int64
				MaxNumReportPerBWP_SP_r19        int64
				MaxNumReportAcrossAllCC_r19      int64
				MaxNumOccasion_r19               int64
				MonitoringResourceType_r19       int64
				MonitoringReportType_r19         int64
			}{}
			c := ie.Aiml_BM_Monitoring_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumTotalResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumTotalResource_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumReportPerBWP_Periodic_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumReportPerBWP_Aperiodic_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumReportPerBWP_SP_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumReportAcrossAllCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumReportAcrossAllCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMMonitoringR19MaxNumOccasionR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOccasion_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringResourceTypeR19Constraints)
				if err != nil {
					return err
				}
				c.MonitoringResourceType_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMMonitoringR19MonitoringReportTypeR19Constraints)
				if err != nil {
					return err
				}
				c.MonitoringReportType_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(16) {
			ie.Aiml_BM_UE_DataCollection_r19 = &struct {
				SubCase_r19            int64
				MaxNumResourceSetB_r19 int64
				MaxNumResourceSetA_r19 int64
			}{}
			c := ie.Aiml_BM_UE_DataCollection_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMUEDataCollectionR19SubCaseR19Constraints)
				if err != nil {
					return err
				}
				c.SubCase_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetBR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumResourceSetB_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlBMUEDataCollectionR19MaxNumResourceSetAR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumResourceSetA_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(17) {
			ie.Aiml_CSI_Prediction_r19 = &struct {
				SupportedCSI_RS_ResourceList_r19 []int64
				ScalingFactor_r19                int64
				NumberOfOccupiedCPU_r19          int64
				NumberOfOccupiedCPUx_r19         int64
				RelaxationTimelineT_r19          struct {
					Scs15kHz_r19  int64
					Scs30kHz_r19  int64
					Scs60kHz_r19  int64
					Scs120kHz_r19 int64
					Scs480kHz_r19 int64
					Scs960kHz_r19 int64
				}
				OccupiedResourcePool_r19 int64
				InferenceReportType_r19  struct {
					Aperiodic_r19      int64
					SemiPersistent_r19 int64
				}
			}{}
			c := ie.Aiml_CSI_Prediction_r19
			{
				seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtAimlCSIPredictionR19SupportedCSIRSResourceListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList_r19[i] = v
				}
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19ScalingFactorR19Constraints)
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPU_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumberOfOccupiedCPUx_r19 = v
			}
			{
				c := &c.RelaxationTimelineT_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs15kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs15kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs30kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs30kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs60kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs60kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs120kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs120kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs480kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs480kHz_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19RelaxationTimelineTR19Scs960kHzR19Constraints)
					if err != nil {
						return err
					}
					c.Scs960kHz_r19 = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.OccupiedResourcePool_r19 = v
			}
			{
				c := &c.InferenceReportType_r19
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19AperiodicR19Constraints)
					if err != nil {
						return err
					}
					c.Aperiodic_r19 = v
				}
				{
					v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionR19InferenceReportTypeR19SemiPersistentR19Constraints)
					if err != nil {
						return err
					}
					c.SemiPersistent_r19 = v
				}
			}
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionUnitDurationDDR19Constraints)
			if err != nil {
				return err
			}
			ie.Aiml_CSI_PredictionUnitDurationDD_r19 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			ie.Aiml_CSI_PredictionN4_r19 = &struct {
				SupportedCSI_RS_ReportSettingAcrossAllCC_r19     []SupportedCSI_RS_ReportSetting_r18
				SupportedCSI_RS_ReportSettingAcrossOneReport_r19 []SupportedCSI_RS_ReportSetting_r18
				NumOccupiedCPU_r19                               int64
				NumOccupiedCPUx_r19                              int64
				OccupiedPool_r19                                 int64
			}{}
			c := ie.Aiml_CSI_PredictionN4_r19
			{
				seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossAllCCR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingAcrossAllCC_r19 = make([]SupportedCSI_RS_ReportSetting_r18, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingAcrossAllCC_r19[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			{
				seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtAimlCSIPredictionN4R19SupportedCSIRSReportSettingAcrossOneReportR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ReportSettingAcrossOneReport_r19 = make([]SupportedCSI_RS_ReportSetting_r18, n)
				for i := int64(0); i < n; i++ {
					if err := c.SupportedCSI_RS_ReportSettingAcrossOneReport_r19[i].Decode(dx); err != nil {
						return err
					}
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumOccupiedCPU_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 8))
				if err != nil {
					return err
				}
				c.NumOccupiedCPUx_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.OccupiedPool_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtAimlCSIPredictionUEDataCollectionR19Constraints)
			if err != nil {
				return err
			}
			ie.Aiml_CSI_PredictionUE_DataCollection_r19 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			ie.Aiml_CSI_PredictionMonitoring_r19 = &struct {
				SupportedCSI_RS_ResourceList_r19 []int64
				NumOccupiedCPU_r19               int64
			}{}
			c := ie.Aiml_CSI_PredictionMonitoring_r19
			{
				seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtAimlCSIPredictionMonitoringR19SupportedCSIRSResourceListR19Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.SupportedCSI_RS_ResourceList_r19 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, common.MaxNrofCSI_RS_ResourcesAlt_1_r16))
					if err != nil {
						return err
					}
					c.SupportedCSI_RS_ResourceList_r19[i] = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.NumOccupiedCPU_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandUeiModeAEvent2R19Constraints)
			if err != nil {
				return err
			}
			ie.Uei_ModeA_Event2_r19 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			ie.Uei_ModeB_r19 = &struct {
				Scs15kHz_r19  *int64
				Scs30kHz_r19  *int64
				Scs60kHz_r19  *int64
				Scs120kHz_r19 *int64
				Scs480kHz_r19 *int64
				Scs960kHz_r19 *int64
			}{}
			c := ie.Uei_ModeB_r19
			mIMOParametersPerBandExtUeiModeBR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtUeiModeBR19Constraints)
			if err := mIMOParametersPerBandExtUeiModeBR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(0) {
				c.Scs15kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs15kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs15kHz_r19) = v
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(1) {
				c.Scs30kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs30kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs30kHz_r19) = v
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(2) {
				c.Scs60kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs60kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs60kHz_r19) = v
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(3) {
				c.Scs120kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs120kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs120kHz_r19) = v
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(4) {
				c.Scs480kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs480kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs480kHz_r19) = v
			}
			if mIMOParametersPerBandExtUeiModeBR19Seq.IsComponentPresent(5) {
				c.Scs960kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeBR19Scs960kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs960kHz_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeInteger(mIMOParametersPerBandUeiTriggerEventDeterminationR19Constraints)
			if err != nil {
				return err
			}
			ie.Uei_TriggerEventDetermination_r19 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtUeiModeAEvent1R19Constraints)
			if err != nil {
				return err
			}
			ie.Uei_ModeA_Event1_r19 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeBitString(mIMOParametersPerBandUeiModeAEvent7R19Constraints)
			if err != nil {
				return err
			}
			ie.Uei_ModeA_Event7_r19 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtEvent2ConditionIndicationR19Constraints)
			if err != nil {
				return err
			}
			ie.Event2ConditionIndication_r19 = &v
		}

		if groupSeq.IsComponentPresent(28) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimeRestriction128PortR19Constraints)
			if err != nil {
				return err
			}
			ie.TimeRestriction128Port_r19 = &v
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtGroupScalingFactorR19Constraints)
			if err != nil {
				return err
			}
			ie.GroupScalingFactor_r19 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			ie.Nes_SD_Type1_SP_r19 = &struct {
				Timeline_r19            int64
				NumOfPortCSI_Report_r19 per.BitString
			}{}
			c := ie.Nes_SD_Type1_SP_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtNesSDType1SPR19TimelineR19Constraints)
				if err != nil {
					return err
				}
				c.Timeline_r19 = v
			}
			{
				v, err := dx.DecodeBitString(per.FixedSize(10))
				if err != nil {
					return err
				}
				c.NumOfPortCSI_Report_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(31) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMrAlwaysReportedType1SPR19Constraints)
			if err != nil {
				return err
			}
			ie.Mr_AlwaysReportedType1SP_r19 = &v
		}

		if groupSeq.IsComponentPresent(32) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtMrAlwaysReportedEType2R19Constraints)
			if err != nil {
				return err
			}
			ie.Mr_AlwaysReported_EType2_r19 = &v
		}

		if groupSeq.IsComponentPresent(33) {
			ie.Cjtc_DdReport_r19 = &struct {
				MinRangeDdInCyclicPrefix_r19 int64
				MaxResolutionDd_r19          int64
				ScalingFactor_r19            int64
			}{}
			c := ie.Cjtc_DdReport_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportR19MinRangeDdInCyclicPrefixR19Constraints)
				if err != nil {
					return err
				}
				c.MinRangeDdInCyclicPrefix_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportR19MaxResolutionDdR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolutionDd_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(34) {
			ie.Cjtc_DdReportProcessing_r19 = &struct {
				MaxNumberTRS_Resource_r19            int64
				MaxNumberTRS_ResourceAcrossCC_r19    int64
				MaxNumberCSI_RS_ResourcePerCC_r19    int64
				MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
				ValueX_r19                           int64
			}{}
			c := ie.Cjtc_DdReportProcessing_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(35) {
			ie.Cjtc_FO_Report_r19 = &struct {
				MinRangeFO_r19      int64
				MaxResolutionFO_r19 int64
				ScalingFactor_r19   int64
			}{}
			c := ie.Cjtc_FO_Report_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportR19MinRangeFOR19Constraints)
				if err != nil {
					return err
				}
				c.MinRangeFO_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportR19MaxResolutionFOR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolutionFO_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(36) {
			ie.Cjtc_FO_ReportProcessing_r19 = &struct {
				MaxNumberTRS_Resource_r19            int64
				MaxNumberTRS_ResourceAcrossCC_r19    int64
				MaxNumberCSI_RS_ResourcePerCC_r19    int64
				MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
				ValueX_r19                           int64
			}{}
			c := ie.Cjtc_FO_ReportProcessing_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(37) {
			ie.Cjtc_PO_ReportWideband_r19 = &struct {
				MaxResolution_r19   int64
				ScalingFactor_r19   int64
				MaxSlotDuration_r19 int64
			}{}
			c := ie.Cjtc_PO_ReportWideband_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportWidebandR19MaxResolutionR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolution_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxSlotDuration_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(38) {
			ie.Cjtc_PO_ReportWidebandProcessing_r19 = &struct {
				MaxNumberCSI_RS_Configured_r19         int64
				MaxNumberCSI_RS_ConfiguredAcrossCC_r19 int64
				MaxNumberCSI_RS_ResourcePerCC_r19      int64
				MaxNumberCSI_RS_ResourceAcrossCC_r19   int64
				ValueX_r19                             int64
			}{}
			c := ie.Cjtc_PO_ReportWidebandProcessing_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_Configured_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSConfiguredAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ConfiguredAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportWidebandProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(39) {
			ie.Cjtc_PO_ReportSubband_r19 = &struct {
				MaxResolution_r19   int64
				MinSubbandSize_r19  int64
				ScalingFactor_r19   int64
				MaxSlotDuration_r19 int64
			}{}
			c := ie.Cjtc_PO_ReportSubband_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportSubbandR19MaxResolutionR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolution_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcPOReportSubbandR19MinSubbandSizeR19Constraints)
				if err != nil {
					return err
				}
				c.MinSubbandSize_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxSlotDuration_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(40) {
			ie.Cjtc_DdFO_Report_r19 = &struct {
				MinRangeDdInCyclicPrefix_r19 int64
				MaxResolutionDd_r19          int64
				MinRangeFO_r19               int64
				MaxResolutionFO_r19          int64
				ScalingFactor_r19            int64
			}{}
			c := ie.Cjtc_DdFO_Report_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeDdInCyclicPrefixR19Constraints)
				if err != nil {
					return err
				}
				c.MinRangeDdInCyclicPrefix_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionDdR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolutionDd_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportR19MinRangeFOR19Constraints)
				if err != nil {
					return err
				}
				c.MinRangeFO_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportR19MaxResolutionFOR19Constraints)
				if err != nil {
					return err
				}
				c.MaxResolutionFO_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ScalingFactor_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(41) {
			ie.Cjtc_DdFO_ReportProcessing_r19 = &struct {
				MaxNumberTRS_Resource_r19            int64
				MaxNumberTRS_ResourceAcrossCC_r19    int64
				MaxNumberCSI_RS_ResourcePerCC_r19    int64
				MaxNumberCSI_RS_ResourceAcrossCC_r19 int64
				ValueX_r19                           int64
			}{}
			c := ie.Cjtc_DdFO_ReportProcessing_r19
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberTRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTRS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourcePerCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtcDdFOReportProcessingR19MaxNumberCSIRSResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberCSI_RS_ResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.ValueX_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(42) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtQCLPDSCHSchemeCR19Constraints)
			if err != nil {
				return err
			}
			ie.Cjt_QCL_PDSCH_SchemeC_r19 = &v
		}

		if groupSeq.IsComponentPresent(43) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtQCLPDSCHSchemeDR19Constraints)
			if err != nil {
				return err
			}
			ie.Cjt_QCL_PDSCH_SchemeD_r19 = &v
		}

		if groupSeq.IsComponentPresent(44) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtCjtQCLPDSCHSchemeER19Constraints)
			if err != nil {
				return err
			}
			ie.Cjt_QCL_PDSCH_SchemeE_r19 = &v
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLinkedCJTCDdEType2CJTJointR19Constraints)
			if err != nil {
				return err
			}
			ie.Linked_CJTC_Dd_EType2CJT_Joint_r19 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparateR19Constraints)
			if err != nil {
				return err
			}
			ie.Linked_CJTC_Dd_EType2CJT_Separate_r19 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtLinkedCJTCDdEType2CJTSeparatePerStateR19Constraints)
			if err != nil {
				return err
			}
			ie.Linked_CJTC_Dd_EType2CJT_SeparatePerState_r19 = &v
		}

		if groupSeq.IsComponentPresent(48) {
			ie.TimelineRelax_CJTC_Dd_EType2CJT_r19 = &struct {
				Scs15kHz_r19  *int64
				Scs30kHz_r19  *int64
				Scs60kHz_r19  *int64
				Scs120kHz_r19 *int64
				Scs480kHz_r19 *int64
				Scs960kHz_r19 *int64
			}{}
			c := ie.TimelineRelax_CJTC_Dd_EType2CJT_r19
			mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq := dx.NewSequenceDecoder(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Constraints)
			if err := mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.DecodePreamble(); err != nil {
				return err
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(0) {
				c.Scs15kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs15kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs15kHz_r19) = v
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(1) {
				c.Scs30kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs30kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs30kHz_r19) = v
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(2) {
				c.Scs60kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs60kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs60kHz_r19) = v
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(3) {
				c.Scs120kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs120kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs120kHz_r19) = v
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(4) {
				c.Scs480kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs480kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs480kHz_r19) = v
			}
			if mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Seq.IsComponentPresent(5) {
				c.Scs960kHz_r19 = new(int64)
				v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTimelineRelaxCJTCDdEType2CJTR19Scs960kHzR19Constraints)
				if err != nil {
					return err
				}
				(*c.Scs960kHz_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(49) {
			seqOf := dx.NewSequenceOfDecoder(mIMOParametersPerBandExtNonCodebookCSIRSSRS3TxPUSCHR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19 = make([]SupportedCSI_RS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.NonCodebook_CSI_RS_SRS_3TxPUSCH_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(50) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSJointTCIR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPUCCH_PUSCH_SRS_JointTCI_r19 = &v
		}

		if groupSeq.IsComponentPresent(51) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetPUCCHPUSCHSRSSeparateTCIR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPUCCH_PUSCH_SRS_SeparateTCI_r19 = &v
		}

		if groupSeq.IsComponentPresent(52) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetPRACHJointTCIR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPRACH_JointTCI_r19 = &v
		}

		if groupSeq.IsComponentPresent(53) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetPRACHSeparateTCIR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPRACH_SeparateTCI_r19 = &v
		}

		if groupSeq.IsComponentPresent(54) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtExtendedStartBitDCI23R19Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedStartBitDCI_2_3_r19 = &v
		}

		if groupSeq.IsComponentPresent(55) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoSRSPwrControlAdjustR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSRS_PwrControlAdjust_r19 = &v
		}

		if groupSeq.IsComponentPresent(56) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtOverlapULTransReductionEnhR19Constraints)
			if err != nil {
				return err
			}
			ie.OverlapUL_TransReductionEnh_r19 = &v
		}

		if groupSeq.IsComponentPresent(57) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetUpdateR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetUpdate_r19 = &v
		}

		if groupSeq.IsComponentPresent(58) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoSRSTPCDCI23R19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSRS_TPC_DCI_2_3_r19 = &v
		}

		if groupSeq.IsComponentPresent(59) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtSrsTPCCLPCAdjustmentStateR19Constraints)
			if err != nil {
				return err
			}
			ie.Srs_TPC_CLPC_AdjustmentState_r19 = &v
		}

		if groupSeq.IsComponentPresent(60) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoSRSDCI11SeparateR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSRS_DCI_1_1_Separate_r19 = &v
		}

		if groupSeq.IsComponentPresent(61) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtTwoSRSDCI11JointR19Constraints)
			if err != nil {
				return err
			}
			ie.TwoSRS_DCI_1_1_Joint_r19 = &v
		}

		if groupSeq.IsComponentPresent(62) {
			v, err := dx.DecodeEnumerated(mIMOParametersPerBandExtPathlossOffsetPHRR19Constraints)
			if err != nil {
				return err
			}
			ie.PathlossOffsetPHR_r19 = &v
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "aiml-BM-Case2-v1920", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Aiml_BM_Case2_v1920 = &struct{ TimeGap_r19 per.BitString }{}
			c := ie.Aiml_BM_Case2_v1920
			{
				v, err := dx.DecodeBitString(per.FixedSize(5))
				if err != nil {
					return err
				}
				c.TimeGap_r19 = v
			}
		}
	}

	return nil
}
