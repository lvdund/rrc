// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandNR (line 23677).

var bandNRConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "bandNR"},
		{Name: "modifiedMPR-Behaviour", Optional: true},
		{Name: "mimo-ParametersPerBand", Optional: true},
		{Name: "extendedCP", Optional: true},
		{Name: "multipleTCI", Optional: true},
		{Name: "bwp-WithoutRestriction", Optional: true},
		{Name: "bwp-SameNumerology", Optional: true},
		{Name: "bwp-DiffNumerology", Optional: true},
		{Name: "crossCarrierScheduling-SameSCS", Optional: true},
		{Name: "pdsch-256QAM-FR2", Optional: true},
		{Name: "pusch-256QAM", Optional: true},
		{Name: "ue-PowerClass", Optional: true},
		{Name: "rateMatchingLTE-CRS", Optional: true},
		{Name: "channelBWs-DL", Optional: true},
		{Name: "channelBWs-UL", Optional: true},
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
		{Name: "extension-addition-15", Optional: true},
		{Name: "extension-addition-16", Optional: true},
		{Name: "extension-addition-17", Optional: true},
		{Name: "extension-addition-18", Optional: true},
		{Name: "extension-addition-19", Optional: true},
		{Name: "extension-addition-20", Optional: true},
		{Name: "extension-addition-21", Optional: true},
	},
}

var bandNRModifiedMPRBehaviourConstraints = per.FixedSize(8)

const (
	BandNR_ExtendedCP_Supported = 0
)

var bandNRExtendedCPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_ExtendedCP_Supported},
}

const (
	BandNR_MultipleTCI_Supported = 0
)

var bandNRMultipleTCIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_MultipleTCI_Supported},
}

const (
	BandNR_Bwp_WithoutRestriction_Supported = 0
)

var bandNRBwpWithoutRestrictionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Bwp_WithoutRestriction_Supported},
}

const (
	BandNR_Bwp_SameNumerology_Upto2 = 0
	BandNR_Bwp_SameNumerology_Upto4 = 1
)

var bandNRBwpSameNumerologyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Bwp_SameNumerology_Upto2, BandNR_Bwp_SameNumerology_Upto4},
}

const (
	BandNR_Bwp_DiffNumerology_Upto4 = 0
)

var bandNRBwpDiffNumerologyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Bwp_DiffNumerology_Upto4},
}

const (
	BandNR_CrossCarrierScheduling_SameSCS_Supported = 0
)

var bandNRCrossCarrierSchedulingSameSCSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_CrossCarrierScheduling_SameSCS_Supported},
}

const (
	BandNR_Pdsch_256QAM_FR2_Supported = 0
)

var bandNRPdsch256QAMFR2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Pdsch_256QAM_FR2_Supported},
}

const (
	BandNR_Pusch_256QAM_Supported = 0
)

var bandNRPusch256QAMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Pusch_256QAM_Supported},
}

const (
	BandNR_Ue_PowerClass_Pc1 = 0
	BandNR_Ue_PowerClass_Pc2 = 1
	BandNR_Ue_PowerClass_Pc3 = 2
	BandNR_Ue_PowerClass_Pc4 = 3
)

var bandNRUePowerClassConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ue_PowerClass_Pc1, BandNR_Ue_PowerClass_Pc2, BandNR_Ue_PowerClass_Pc3, BandNR_Ue_PowerClass_Pc4},
}

const (
	BandNR_RateMatchingLTE_CRS_Supported = 0
)

var bandNRRateMatchingLTECRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_RateMatchingLTE_CRS_Supported},
}

var bandNRChannelBWsDLConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1"},
		{Name: "fr2"},
	},
}

const (
	BandNR_ChannelBWs_DL_Fr1 = 0
	BandNR_ChannelBWs_DL_Fr2 = 1
)

var bandNRChannelBWsULConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1"},
		{Name: "fr2"},
	},
}

const (
	BandNR_ChannelBWs_UL_Fr1 = 0
	BandNR_ChannelBWs_UL_Fr2 = 1
)

const (
	BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N60  = 0
	BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N70  = 1
	BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N80  = 2
	BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N90  = 3
	BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N100 = 4
)

var bandNRExtMaxUplinkDutyCyclePC2FR1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N60, BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N70, BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N80, BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N90, BandNR_Ext_MaxUplinkDutyCycle_PC2_FR1_N100},
}

const (
	BandNR_Ext_Pucch_SpatialRelInfoMAC_CE_Supported = 0
)

var bandNRExtPucchSpatialRelInfoMACCEConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pucch_SpatialRelInfoMAC_CE_Supported},
}

const (
	BandNR_Ext_PowerBoosting_Pi2BPSK_Supported = 0
)

var bandNRExtPowerBoostingPi2BPSKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PowerBoosting_Pi2BPSK_Supported},
}

const (
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N15  = 0
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N20  = 1
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N25  = 2
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N30  = 3
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N40  = 4
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N50  = 5
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N60  = 6
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N70  = 7
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N80  = 8
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N90  = 9
	BandNR_Ext_MaxUplinkDutyCycle_FR2_N100 = 10
)

var bandNRExtMaxUplinkDutyCycleFR2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxUplinkDutyCycle_FR2_N15, BandNR_Ext_MaxUplinkDutyCycle_FR2_N20, BandNR_Ext_MaxUplinkDutyCycle_FR2_N25, BandNR_Ext_MaxUplinkDutyCycle_FR2_N30, BandNR_Ext_MaxUplinkDutyCycle_FR2_N40, BandNR_Ext_MaxUplinkDutyCycle_FR2_N50, BandNR_Ext_MaxUplinkDutyCycle_FR2_N60, BandNR_Ext_MaxUplinkDutyCycle_FR2_N70, BandNR_Ext_MaxUplinkDutyCycle_FR2_N80, BandNR_Ext_MaxUplinkDutyCycle_FR2_N90, BandNR_Ext_MaxUplinkDutyCycle_FR2_N100},
}

var bandNRExtChannelBWsDLV1590Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1"},
		{Name: "fr2"},
	},
}

const (
	BandNR_Ext_ChannelBWs_DL_v1590_Fr1 = 0
	BandNR_Ext_ChannelBWs_DL_v1590_Fr2 = 1
)

var bandNRExtChannelBWsULV1590Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1"},
		{Name: "fr2"},
	},
}

const (
	BandNR_Ext_ChannelBWs_UL_v1590_Fr1 = 0
	BandNR_Ext_ChannelBWs_UL_v1590_Fr2 = 1
)

var bandNRAsymmetricBandwidthCombinationSetConstraints = per.SizeRange(1, 32)

const (
	BandNR_Ext_CancelOverlappingPUSCH_r16_Supported = 0
)

var bandNRExtCancelOverlappingPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CancelOverlappingPUSCH_r16_Supported},
}

const (
	BandNR_Ext_OverlapRateMatchingEUTRA_CRS_r16_Supported = 0
)

var bandNRExtOverlapRateMatchingEUTRACRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_OverlapRateMatchingEUTRA_CRS_r16_Supported},
}

const (
	BandNR_Ext_Pdsch_MappingTypeB_Alt_r16_Supported = 0
)

var bandNRExtPdschMappingTypeBAltR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdsch_MappingTypeB_Alt_r16_Supported},
}

const (
	BandNR_Ext_OneSlotPeriodicTRS_r16_Supported = 0
)

var bandNRExtOneSlotPeriodicTRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_OneSlotPeriodicTRS_r16_Supported},
}

const (
	BandNR_Ext_SimulSRS_MIMO_TransWithinBand_r16_N2 = 0
)

var bandNRExtSimulSRSMIMOTransWithinBandR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SimulSRS_MIMO_TransWithinBand_r16_N2},
}

var bandNRExtChannelBWDLIABR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-100mhz"},
		{Name: "fr2-200mhz"},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz = 0
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz = 1
)

var bandNRExtChannelBWULIABR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-100mhz"},
		{Name: "fr2-200mhz"},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz = 0
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz = 1
)

const (
	BandNR_Ext_RasterShift7dot5_IAB_r16_Supported = 0
)

var bandNRExtRasterShift7dot5IABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_RasterShift7dot5_IAB_r16_Supported},
}

const (
	BandNR_Ext_Ue_PowerClass_v1610_Pc1dot5 = 0
)

var bandNRExtUePowerClassV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ue_PowerClass_v1610_Pc1dot5},
}

const (
	BandNR_Ext_CondHandover_r16_Supported = 0
)

var bandNRExtCondHandoverR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondHandover_r16_Supported},
}

const (
	BandNR_Ext_CondHandoverFailure_r16_Supported = 0
)

var bandNRExtCondHandoverFailureR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondHandoverFailure_r16_Supported},
}

const (
	BandNR_Ext_CondHandoverTwoTriggerEvents_r16_Supported = 0
)

var bandNRExtCondHandoverTwoTriggerEventsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondHandoverTwoTriggerEvents_r16_Supported},
}

const (
	BandNR_Ext_CondPSCellChange_r16_Supported = 0
)

var bandNRExtCondPSCellChangeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondPSCellChange_r16_Supported},
}

const (
	BandNR_Ext_CondPSCellChangeTwoTriggerEvents_r16_Supported = 0
)

var bandNRExtCondPSCellChangeTwoTriggerEventsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondPSCellChangeTwoTriggerEvents_r16_Supported},
}

const (
	BandNR_Ext_Mpr_PowerBoost_FR2_r16_Supported = 0
)

var bandNRExtMprPowerBoostFR2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mpr_PowerBoost_FR2_r16_Supported},
}

const (
	BandNR_Ext_JointReleaseConfiguredGrantType2_r16_Supported = 0
)

var bandNRExtJointReleaseConfiguredGrantType2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_JointReleaseConfiguredGrantType2_r16_Supported},
}

const (
	BandNR_Ext_JointReleaseSPS_r16_Supported = 0
)

var bandNRExtJointReleaseSPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_JointReleaseSPS_r16_Supported},
}

const (
	BandNR_Ext_SimulSRS_TransWithinBand_r16_N2 = 0
)

var bandNRExtSimulSRSTransWithinBandR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SimulSRS_TransWithinBand_r16_N2},
}

const (
	BandNR_Ext_Trs_AdditionalBandwidth_r16_Trs_AddBW_Set1 = 0
	BandNR_Ext_Trs_AdditionalBandwidth_r16_Trs_AddBW_Set2 = 1
)

var bandNRExtTrsAdditionalBandwidthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Trs_AdditionalBandwidth_r16_Trs_AddBW_Set1, BandNR_Ext_Trs_AdditionalBandwidth_r16_Trs_AddBW_Set2},
}

const (
	BandNR_Ext_HandoverIntraF_IAB_r16_Supported = 0
)

var bandNRExtHandoverIntraFIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_HandoverIntraF_IAB_r16_Supported},
}

const (
	BandNR_Ext_HandoverUTRA_FDD_r16_Supported = 0
)

var bandNRExtHandoverUTRAFDDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_HandoverUTRA_FDD_r16_Supported},
}

const (
	BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us2 = 0
	BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us4 = 1
	BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us7 = 2
)

var bandNRExtEnhancedULTransientPeriodR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us2, BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us4, BandNR_Ext_EnhancedUL_TransientPeriod_r16_Us7},
}

const (
	BandNR_Ext_Type1_PUSCH_RepetitionMultiSlots_v1650_Supported = 0
)

var bandNRExtType1PUSCHRepetitionMultiSlotsV1650Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Type1_PUSCH_RepetitionMultiSlots_v1650_Supported},
}

const (
	BandNR_Ext_Type2_PUSCH_RepetitionMultiSlots_v1650_Supported = 0
)

var bandNRExtType2PUSCHRepetitionMultiSlotsV1650Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Type2_PUSCH_RepetitionMultiSlots_v1650_Supported},
}

const (
	BandNR_Ext_Pusch_RepetitionMultiSlots_v1650_Supported = 0
)

var bandNRExtPuschRepetitionMultiSlotsV1650Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pusch_RepetitionMultiSlots_v1650_Supported},
}

const (
	BandNR_Ext_ConfiguredUL_GrantType1_v1650_Supported = 0
)

var bandNRExtConfiguredULGrantType1V1650Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ConfiguredUL_GrantType1_v1650_Supported},
}

const (
	BandNR_Ext_ConfiguredUL_GrantType2_v1650_Supported = 0
)

var bandNRExtConfiguredULGrantType2V1650Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ConfiguredUL_GrantType2_v1650_Supported},
}

const (
	BandNR_Ext_EnhancedSkipUplinkTxConfigured_v1660_Supported = 0
)

var bandNRExtEnhancedSkipUplinkTxConfiguredV1660Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedSkipUplinkTxConfigured_v1660_Supported},
}

const (
	BandNR_Ext_EnhancedSkipUplinkTxDynamic_v1660_Supported = 0
)

var bandNRExtEnhancedSkipUplinkTxDynamicV1660Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedSkipUplinkTxDynamic_v1660_Supported},
}

const (
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N10  = 0
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N15  = 1
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N20  = 2
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N25  = 3
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N30  = 4
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N40  = 5
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N50  = 6
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N60  = 7
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N70  = 8
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N80  = 9
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N90  = 10
	BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N100 = 11
)

var bandNRExtMaxUplinkDutyCyclePC1dot5MPEFR1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N10, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N15, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N20, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N25, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N30, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N40, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N50, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N60, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N70, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N80, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N90, BandNR_Ext_MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16_N100},
}

const (
	BandNR_Ext_TxDiversity_r16_Supported = 0
)

var bandNRExtTxDiversityR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TxDiversity_r16_Supported},
}

const (
	BandNR_Ext_Pdsch_1024QAM_FR1_r17_Supported = 0
)

var bandNRExtPdsch1024QAMFR1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdsch_1024QAM_FR1_r17_Supported},
}

const (
	BandNR_Ext_Ue_PowerClass_v1700_Pc5 = 0
	BandNR_Ext_Ue_PowerClass_v1700_Pc6 = 1
	BandNR_Ext_Ue_PowerClass_v1700_Pc7 = 2
)

var bandNRExtUePowerClassV1700Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ue_PowerClass_v1700_Pc5, BandNR_Ext_Ue_PowerClass_v1700_Pc6, BandNR_Ext_Ue_PowerClass_v1700_Pc7},
}

const (
	BandNR_Ext_Rlm_Relaxation_r17_Supported = 0
)

var bandNRExtRlmRelaxationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Rlm_Relaxation_r17_Supported},
}

const (
	BandNR_Ext_Bfd_Relaxation_r17_Supported = 0
)

var bandNRExtBfdRelaxationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Bfd_Relaxation_r17_Supported},
}

const (
	BandNR_Ext_Cg_SDT_r17_Supported = 0
)

var bandNRExtCgSDTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Cg_SDT_r17_Supported},
}

const (
	BandNR_Ext_LocationBasedCondHandover_r17_Supported = 0
)

var bandNRExtLocationBasedCondHandoverR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_LocationBasedCondHandover_r17_Supported},
}

const (
	BandNR_Ext_TimeBasedCondHandover_r17_Supported = 0
)

var bandNRExtTimeBasedCondHandoverR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TimeBasedCondHandover_r17_Supported},
}

const (
	BandNR_Ext_EventA4BasedCondHandover_r17_Supported = 0
)

var bandNRExtEventA4BasedCondHandoverR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EventA4BasedCondHandover_r17_Supported},
}

const (
	BandNR_Ext_Mn_InitiatedCondPSCellChangeNRDC_r17_Supported = 0
)

var bandNRExtMnInitiatedCondPSCellChangeNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mn_InitiatedCondPSCellChangeNRDC_r17_Supported},
}

const (
	BandNR_Ext_Sn_InitiatedCondPSCellChangeNRDC_r17_Supported = 0
)

var bandNRExtSnInitiatedCondPSCellChangeNRDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sn_InitiatedCondPSCellChangeNRDC_r17_Supported},
}

const (
	BandNR_Ext_Pdcch_SkippingWithoutSSSG_r17_Supported = 0
)

var bandNRExtPdcchSkippingWithoutSSSGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdcch_SkippingWithoutSSSG_r17_Supported},
}

const (
	BandNR_Ext_Sssg_Switching_1BitInd_r17_Supported = 0
)

var bandNRExtSssgSwitching1BitIndR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sssg_Switching_1BitInd_r17_Supported},
}

const (
	BandNR_Ext_Sssg_Switching_2BitInd_r17_Supported = 0
)

var bandNRExtSssgSwitching2BitIndR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sssg_Switching_2BitInd_r17_Supported},
}

const (
	BandNR_Ext_Pdcch_SkippingWithSSSG_r17_Supported = 0
)

var bandNRExtPdcchSkippingWithSSSGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdcch_SkippingWithSSSG_r17_Supported},
}

const (
	BandNR_Ext_SearchSpaceSetGrp_SwitchCap2_r17_Supported = 0
)

var bandNRExtSearchSpaceSetGrpSwitchCap2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SearchSpaceSetGrp_SwitchCap2_r17_Supported},
}

const (
	BandNR_Ext_UplinkPreCompensation_r17_Supported = 0
)

var bandNRExtUplinkPreCompensationR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UplinkPreCompensation_r17_Supported},
}

const (
	BandNR_Ext_Uplink_TA_Reporting_r17_Supported = 0
)

var bandNRExtUplinkTAReportingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Uplink_TA_Reporting_r17_Supported},
}

const (
	BandNR_Ext_Max_HARQ_ProcessNumber_r17_U16d32 = 0
	BandNR_Ext_Max_HARQ_ProcessNumber_r17_U32d16 = 1
	BandNR_Ext_Max_HARQ_ProcessNumber_r17_U32d32 = 2
)

var bandNRExtMaxHARQProcessNumberR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Max_HARQ_ProcessNumber_r17_U16d32, BandNR_Ext_Max_HARQ_ProcessNumber_r17_U32d16, BandNR_Ext_Max_HARQ_ProcessNumber_r17_U32d32},
}

const (
	BandNR_Ext_Type2_HARQ_Codebook_r17_Supported = 0
)

var bandNRExtType2HARQCodebookR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Type2_HARQ_Codebook_r17_Supported},
}

const (
	BandNR_Ext_Type1_HARQ_Codebook_r17_Supported = 0
)

var bandNRExtType1HARQCodebookR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Type1_HARQ_Codebook_r17_Supported},
}

const (
	BandNR_Ext_Type3_HARQ_Codebook_r17_Supported = 0
)

var bandNRExtType3HARQCodebookR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Type3_HARQ_Codebook_r17_Supported},
}

const (
	BandNR_Ext_Ue_Specific_K_Offset_r17_Supported = 0
)

var bandNRExtUeSpecificKOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ue_Specific_K_Offset_r17_Supported},
}

const (
	BandNR_Ext_MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17_Supported = 0
)

var bandNRExtMultiPDSCHSingleDCIFR21SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17_Supported},
}

const (
	BandNR_Ext_MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17_Supported = 0
)

var bandNRExtMultiPUSCHSingleDCIFR21SCS120kHzR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17_Supported},
}

const (
	BandNR_Ext_ParallelPRS_MeasRRC_Inactive_r17_Supported = 0
)

var bandNRExtParallelPRSMeasRRCInactiveR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ParallelPRS_MeasRRC_Inactive_r17_Supported},
}

const (
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N1 = 0
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N2 = 1
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N3 = 2
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N4 = 3
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N6 = 4
	BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N8 = 5
)

var bandNRExtNrUETxTEGIDMaxSupportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N1, BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N2, BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N3, BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N4, BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N6, BandNR_Ext_Nr_UE_TxTEG_ID_MaxSupport_r17_N8},
}

const (
	BandNR_Ext_Prs_ProcessingRRC_Inactive_r17_Supported = 0
)

var bandNRExtPrsProcessingRRCInactiveR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prs_ProcessingRRC_Inactive_r17_Supported},
}

const (
	BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option1 = 0
	BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option2 = 1
	BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option3 = 2
)

var bandNRExtPrsProcessingWindowType1AR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option1, BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option2, BandNR_Ext_Prs_ProcessingWindowType1A_r17_Option3},
}

const (
	BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option1 = 0
	BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option2 = 1
	BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option3 = 2
)

var bandNRExtPrsProcessingWindowType1BR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option1, BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option2, BandNR_Ext_Prs_ProcessingWindowType1B_r17_Option3},
}

const (
	BandNR_Ext_Prs_ProcessingWindowType2_r17_Option1 = 0
	BandNR_Ext_Prs_ProcessingWindowType2_r17_Option2 = 1
	BandNR_Ext_Prs_ProcessingWindowType2_r17_Option3 = 2
)

var bandNRExtPrsProcessingWindowType2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prs_ProcessingWindowType2_r17_Option1, BandNR_Ext_Prs_ProcessingWindowType2_r17_Option2, BandNR_Ext_Prs_ProcessingWindowType2_r17_Option3},
}

const (
	BandNR_Ext_MaxNumberPUSCH_TypeA_Repetition_r17_Supported = 0
)

var bandNRExtMaxNumberPUSCHTypeARepetitionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxNumberPUSCH_TypeA_Repetition_r17_Supported},
}

const (
	BandNR_Ext_PuschTypeA_RepetitionsAvailSlot_r17_Supported = 0
)

var bandNRExtPuschTypeARepetitionsAvailSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PuschTypeA_RepetitionsAvailSlot_r17_Supported},
}

const (
	BandNR_Ext_Tb_ProcessingMultiSlotPUSCH_r17_Supported = 0
)

var bandNRExtTbProcessingMultiSlotPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Tb_ProcessingMultiSlotPUSCH_r17_Supported},
}

const (
	BandNR_Ext_Tb_ProcessingRepMultiSlotPUSCH_r17_Supported = 0
)

var bandNRExtTbProcessingRepMultiSlotPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Tb_ProcessingRepMultiSlotPUSCH_r17_Supported},
}

const (
	BandNR_Ext_Pusch_RepetitionMsg3_r17_Supported = 0
)

var bandNRExtPuschRepetitionMsg3R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pusch_RepetitionMsg3_r17_Supported},
}

const (
	BandNR_Ext_ParallelMeasurementWithoutRestriction_r17_Supported = 0
)

var bandNRExtParallelMeasurementWithoutRestrictionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ParallelMeasurementWithoutRestriction_r17_Supported},
}

const (
	BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N1 = 0
	BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N2 = 1
	BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N3 = 2
	BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N4 = 3
)

var bandNRExtMaxNumberNGSOSatellitesWithinOneSMTCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N1, BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N2, BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N3, BandNR_Ext_MaxNumber_NGSO_SatellitesWithinOneSMTC_r17_N4},
}

const (
	BandNR_Ext_K1_RangeExtension_r17_Supported = 0
)

var bandNRExtK1RangeExtensionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_K1_RangeExtension_r17_Supported},
}

const (
	BandNR_Ext_AperiodicCSI_RS_AdditionalBandwidth_r17_AddBW_Set1 = 0
	BandNR_Ext_AperiodicCSI_RS_AdditionalBandwidth_r17_AddBW_Set2 = 1
)

var bandNRExtAperiodicCSIRSAdditionalBandwidthR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_AperiodicCSI_RS_AdditionalBandwidth_r17_AddBW_Set1, BandNR_Ext_AperiodicCSI_RS_AdditionalBandwidth_r17_AddBW_Set2},
}

const (
	BandNR_Ext_Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17_Supported = 0
)

var bandNRExtBwpWithoutCDSSBOrNCDSSBRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17_Supported},
}

const (
	BandNR_Ext_HalfDuplexFDD_TypeA_RedCap_r17_Supported = 0
)

var bandNRExtHalfDuplexFDDTypeARedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_HalfDuplexFDD_TypeA_RedCap_r17_Supported},
}

var bandNRChannelBWsDLSCS480kHzFR22R17Constraints = per.FixedSize(8)

var bandNRChannelBWsULSCS480kHzFR22R17Constraints = per.FixedSize(8)

var bandNRChannelBWsDLSCS960kHzFR22R17Constraints = per.FixedSize(8)

var bandNRChannelBWsULSCS960kHzFR22R17Constraints = per.FixedSize(8)

const (
	BandNR_Ext_Ul_GapFR2_r17_Supported = 0
)

var bandNRExtUlGapFR2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ul_GapFR2_r17_Supported},
}

const (
	BandNR_Ext_OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17_Supported = 0
)

var bandNRExtOneShotHARQFeedbackTriggeredByDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17_Supported},
}

const (
	BandNR_Ext_OneShotHARQ_FeedbackPhy_Priority_r17_Supported = 0
)

var bandNRExtOneShotHARQFeedbackPhyPriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_OneShotHARQ_FeedbackPhy_Priority_r17_Supported},
}

const (
	BandNR_Ext_Ue_OneShotUL_TimingAdj_r17_Supported = 0
)

var bandNRExtUeOneShotULTimingAdjR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ue_OneShotUL_TimingAdj_r17_Supported},
}

const (
	BandNR_Ext_Pucch_Repetition_F0_2_r17_Supported = 0
)

var bandNRExtPucchRepetitionF02R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pucch_Repetition_F0_2_r17_Supported},
}

const (
	BandNR_Ext_Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17_Supported = 0
)

var bandNRExtCqi4BitsSubbandNTNSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17_Supported},
}

const (
	BandNR_Ext_Mux_HARQ_ACK_DiffPriorities_r17_Supported = 0
)

var bandNRExtMuxHARQACKDiffPrioritiesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mux_HARQ_ACK_DiffPriorities_r17_Supported},
}

const (
	BandNR_Ext_Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17_Supported = 0
)

var bandNRExtTaBasedPDCNTNSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17_Supported},
}

const (
	BandNR_Ext_Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17_Supported = 0
)

var bandNRExtAckNACKFeedbackForMulticastWithDCIEnablerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17_Supported},
}

var bandNRMaxNumberGRNTIR17Constraints = per.Constrained(2, 8)

const (
	BandNR_Ext_DynamicMulticastDCI_Format4_2_r17_Supported = 0
)

var bandNRExtDynamicMulticastDCIFormat42R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_DynamicMulticastDCI_Format4_2_r17_Supported},
}

var bandNRExtMaxModulationOrderForMulticastR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17 = 0
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17 = 1
)

const (
	BandNR_Ext_DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17_N8  = 0
	BandNR_Ext_DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17_N16 = 1
)

var bandNRExtDynamicSlotRepetitionMulticastTNNonSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17_N8, BandNR_Ext_DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17_N16},
}

const (
	BandNR_Ext_DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17_N8  = 0
	BandNR_Ext_DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17_N16 = 1
)

var bandNRExtDynamicSlotRepetitionMulticastNTNSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17_N8, BandNR_Ext_DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17_N16},
}

const (
	BandNR_Ext_Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17_Supported = 0
)

var bandNRExtNackOnlyFeedbackForMulticastWithDCIEnablerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17_Supported},
}

const (
	BandNR_Ext_Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17_Supported = 0
)

var bandNRExtAckNACKFeedbackForSPSMulticastWithDCIEnablerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17_Supported},
}

var bandNRMaxNumberGCSRNTIR17Constraints = per.Constrained(2, 8)

const (
	BandNR_Ext_Re_LevelRateMatchingForMulticast_r17_Supported = 0
)

var bandNRExtReLevelRateMatchingForMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Re_LevelRateMatchingForMulticast_r17_Supported},
}

const (
	BandNR_Ext_Pdsch_1024QAM_2MIMO_FR1_r17_Supported = 0
)

var bandNRExtPdsch1024QAM2MIMOFR1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdsch_1024QAM_2MIMO_FR1_r17_Supported},
}

const (
	BandNR_Ext_Prs_MeasurementWithoutMG_r17_CpLength      = 0
	BandNR_Ext_Prs_MeasurementWithoutMG_r17_QuarterSymbol = 1
	BandNR_Ext_Prs_MeasurementWithoutMG_r17_HalfSymbol    = 2
	BandNR_Ext_Prs_MeasurementWithoutMG_r17_HalfSlot      = 3
)

var bandNRExtPrsMeasurementWithoutMGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prs_MeasurementWithoutMG_r17_CpLength, BandNR_Ext_Prs_MeasurementWithoutMG_r17_QuarterSymbol, BandNR_Ext_Prs_MeasurementWithoutMG_r17_HalfSymbol, BandNR_Ext_Prs_MeasurementWithoutMG_r17_HalfSlot},
}

var bandNRMaxNumberNGSOSatellitesPerCarrierR17Constraints = per.Constrained(3, 4)

var bandNRExtPrsProcessingCapabilityOutsideMGinPPWR17Constraints = per.SizeRange(1, 3)

var bandNRChannelBWsDLSCS120kHzFR22R17Constraints = per.FixedSize(8)

var bandNRChannelBWsULSCS120kHzFR22R17Constraints = per.FixedSize(8)

const (
	BandNR_Ext_Dmrs_BundlingPUSCH_RepTypeA_r17_Supported = 0
)

var bandNRExtDmrsBundlingPUSCHRepTypeAR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingPUSCH_RepTypeA_r17_Supported},
}

const (
	BandNR_Ext_Dmrs_BundlingPUSCH_RepTypeB_r17_Supported = 0
)

var bandNRExtDmrsBundlingPUSCHRepTypeBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingPUSCH_RepTypeB_r17_Supported},
}

const (
	BandNR_Ext_Dmrs_BundlingPUSCH_MultiSlot_r17_Supported = 0
)

var bandNRExtDmrsBundlingPUSCHMultiSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingPUSCH_MultiSlot_r17_Supported},
}

const (
	BandNR_Ext_Dmrs_BundlingPUCCH_Rep_r17_Supported = 0
)

var bandNRExtDmrsBundlingPUCCHRepR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingPUCCH_Rep_r17_Supported},
}

const (
	BandNR_Ext_InterSlotFreqHopInterSlotBundlingPUSCH_r17_Supported = 0
)

var bandNRExtInterSlotFreqHopInterSlotBundlingPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_InterSlotFreqHopInterSlotBundlingPUSCH_r17_Supported},
}

const (
	BandNR_Ext_InterSlotFreqHopPUCCH_r17_Supported = 0
)

var bandNRExtInterSlotFreqHopPUCCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_InterSlotFreqHopPUCCH_r17_Supported},
}

const (
	BandNR_Ext_Dmrs_BundlingRestart_r17_Supported = 0
)

var bandNRExtDmrsBundlingRestartR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingRestart_r17_Supported},
}

const (
	BandNR_Ext_Dmrs_BundlingNonBackToBackTX_r17_Supported = 0
)

var bandNRExtDmrsBundlingNonBackToBackTXR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dmrs_BundlingNonBackToBackTX_r17_Supported},
}

const (
	BandNR_Ext_MaxDynamicSlotRepetitionForSPS_Multicast_r17_N8  = 0
	BandNR_Ext_MaxDynamicSlotRepetitionForSPS_Multicast_r17_N16 = 1
)

var bandNRExtMaxDynamicSlotRepetitionForSPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxDynamicSlotRepetitionForSPS_Multicast_r17_N8, BandNR_Ext_MaxDynamicSlotRepetitionForSPS_Multicast_r17_N16},
}

const (
	BandNR_Ext_Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17_Supported = 0
)

var bandNRExtNackOnlyFeedbackForSPSMulticastWithDCIEnablerR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17_Supported},
}

const (
	BandNR_Ext_Sps_MulticastDCI_Format4_2_r17_Supported = 0
)

var bandNRExtSpsMulticastDCIFormat42R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sps_MulticastDCI_Format4_2_r17_Supported},
}

var bandNRSpsMulticastMultiConfigR17Constraints = per.Constrained(1, 8)

const (
	BandNR_Ext_PriorityIndicatorInDCI_Multicast_r17_Supported = 0
)

var bandNRExtPriorityIndicatorInDCIMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PriorityIndicatorInDCI_Multicast_r17_Supported},
}

const (
	BandNR_Ext_PriorityIndicatorInDCI_SPS_Multicast_r17_Supported = 0
)

var bandNRExtPriorityIndicatorInDCISPSMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PriorityIndicatorInDCI_SPS_Multicast_r17_Supported},
}

const (
	BandNR_Ext_TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17_Supported = 0
)

var bandNRExtTwoHARQACKCodebookForUnicastAndMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17_Supported},
}

const (
	BandNR_Ext_MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17_Supported = 0
)

var bandNRExtMultiPUCCHHARQACKForMulticastUnicastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17_Supported},
}

const (
	BandNR_Ext_ReleaseSPS_MulticastWithCS_RNTI_r17_Supported = 0
)

var bandNRExtReleaseSPSMulticastWithCSRNTIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ReleaseSPS_MulticastWithCS_RNTI_r17_Supported},
}

const (
	BandNR_Ext_PosUE_TA_AutoAdjustment_r18_Supported = 0
)

var bandNRExtPosUETAAutoAdjustmentR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosUE_TA_AutoAdjustment_r18_Supported},
}

const (
	BandNR_Ext_PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18_Supported = 0
)

var bandNRExtPosSRSValidityAreaRRCInactiveInitialULBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18_Supported},
}

const (
	BandNR_Ext_PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18_Supported = 0
)

var bandNRExtPosSRSValidityAreaRRCInactiveOutsideInitialULBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18_Supported},
}

const (
	BandNR_Ext_PosJointTriggerBySingleDCI_RRC_Connected_r18_Supported = 0
)

var bandNRExtPosJointTriggerBySingleDCIRRCConnectedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosJointTriggerBySingleDCI_RRC_Connected_r18_Supported},
}

const (
	BandNR_Ext_Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18_Supported = 0
)

var bandNRExtDlPRSMeasurementWithRxFHRRCInactiveforRedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18_Supported},
}

const (
	BandNR_Ext_Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18_Supported = 0
)

var bandNRExtDlPRSMeasurementWithRxFHRRCIdleforRedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18_Supported},
}

const (
	BandNR_Ext_Nes_CellDTX_DRX_r18_CellDTXonly = 0
	BandNR_Ext_Nes_CellDTX_DRX_r18_CellDRXonly = 1
	BandNR_Ext_Nes_CellDTX_DRX_r18_Both        = 2
)

var bandNRExtNesCellDTXDRXR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nes_CellDTX_DRX_r18_CellDTXonly, BandNR_Ext_Nes_CellDTX_DRX_r18_CellDRXonly, BandNR_Ext_Nes_CellDTX_DRX_r18_Both},
}

const (
	BandNR_Ext_Nes_CellDTX_DRX_DCI2_9_r18_Supported = 0
)

var bandNRExtNesCellDTXDRXDCI29R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nes_CellDTX_DRX_DCI2_9_r18_Supported},
}

const (
	BandNR_Ext_MixCodeBookSpatialAdaptation_r18_Supported = 0
)

var bandNRExtMixCodeBookSpatialAdaptationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MixCodeBookSpatialAdaptation_r18_Supported},
}

var bandNRSimultaneousCSISubReportsPerCCR18Constraints = per.Constrained(1, 8)

const (
	BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N4  = 0
	BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N8  = 1
	BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N16 = 2
	BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N32 = 3
)

var bandNRExtNtnDMRSBundlingNGSOR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N4, BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N8, BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N16, BandNR_Ext_Ntn_DMRS_BundlingNGSO_r18_N32},
}

var bandNRRachEarlyTAMeasurementR18Constraints = per.Constrained(1, 8)

var bandNRUeTAMeasurementR18Constraints = per.Constrained(1, 8)

const (
	BandNR_Ext_Ta_IndicationCellSwitch_r18_Supported = 0
)

var bandNRExtTaIndicationCellSwitchR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ta_IndicationCellSwitch_r18_Supported},
}

const (
	BandNR_Ext_MultiPUSCH_CG_r18_N16 = 0
	BandNR_Ext_MultiPUSCH_CG_r18_N32 = 1
)

var bandNRExtMultiPUSCHCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPUSCH_CG_r18_N16, BandNR_Ext_MultiPUSCH_CG_r18_N32},
}

const (
	BandNR_Ext_JointReleaseDCI_r18_Supported = 0
)

var bandNRExtJointReleaseDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_JointReleaseDCI_r18_Supported},
}

const (
	BandNR_Ext_Cg_PUSCH_UTO_UCI_Ind_r18_Supported = 0
)

var bandNRExtCgPUSCHUTOUCIIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Cg_PUSCH_UTO_UCI_Ind_r18_Supported},
}

const (
	BandNR_Ext_Pdcch_MonitoringResumptionAfterUL_NACK_r18_Supported = 0
)

var bandNRExtPdcchMonitoringResumptionAfterULNACKR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdcch_MonitoringResumptionAfterUL_NACK_r18_Supported},
}

const (
	BandNR_Ext_Support3MHz_ChannelBW_Symmetric_r18_Supported = 0
)

var bandNRExtSupport3MHzChannelBWSymmetricR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Support3MHz_ChannelBW_Symmetric_r18_Supported},
}

const (
	BandNR_Ext_Support3MHz_ChannelBW_Asymmetric_r18_Supported = 0
)

var bandNRExtSupport3MHzChannelBWAsymmetricR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Support3MHz_ChannelBW_Asymmetric_r18_Supported},
}

const (
	BandNR_Ext_Support12PRB_CORESET0_r18_Supported = 0
)

var bandNRExtSupport12PRBCORESET0R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Support12PRB_CORESET0_r18_Supported},
}

const (
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18_Supported = 0
)

var bandNRExtNrPDCCHOverlapLTECRSREMultiPatternsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18_Supported},
}

const (
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18_Supported = 0
)

var bandNRExtNrPDCCHOverlapLTECRSRESpan34R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18_Supported},
}

const (
	BandNR_Ext_OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18_Supported = 0
)

var bandNRExtOverlapRateMatchingEUTRACRSPatterns34DiffCSPoolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18_Supported},
}

const (
	BandNR_Ext_Ncd_SSB_BWP_Wor_r18_Supported = 0
)

var bandNRExtNcdSSBBWPWorR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ncd_SSB_BWP_Wor_r18_Supported},
}

const (
	BandNR_Ext_Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18_Supported = 0
)

var bandNRExtRlmBMBFDCSIRSOutsideActiveBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18_Supported},
}

const (
	BandNR_Ext_Prach_CoverageEnh_r18_Supported = 0
)

var bandNRExtPrachCoverageEnhR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prach_CoverageEnh_r18_Supported},
}

const (
	BandNR_Ext_Prach_Repetition_r18_Supported = 0
)

var bandNRExtPrachRepetitionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Prach_Repetition_r18_Supported},
}

const (
	BandNR_Ext_DynamicWaveformSwitch_r18_Supported = 0
)

var bandNRExtDynamicWaveformSwitchR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_DynamicWaveformSwitch_r18_Supported},
}

const (
	BandNR_Ext_DynamicWaveformSwitchPHR_r18_Supported = 0
)

var bandNRExtDynamicWaveformSwitchPHRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_DynamicWaveformSwitchPHR_r18_Supported},
}

var bandNRDynamicWaveformSwitchIntraCAR18Constraints = per.Constrained(2, 8)

const (
	BandNR_Ext_MultiPUSCH_SingleDCI_NonConsSlots_r18_Supported = 0
)

var bandNRExtMultiPUSCHSingleDCINonConsSlotsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPUSCH_SingleDCI_NonConsSlots_r18_Supported},
}

const (
	BandNR_Ext_MulticastInactive_r18_Supported = 0
)

var bandNRExtMulticastInactiveR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MulticastInactive_r18_Supported},
}

const (
	BandNR_Ext_ThresholdBasedMulticastResume_r18_Supported = 0
)

var bandNRExtThresholdBasedMulticastResumeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ThresholdBasedMulticastResume_r18_Supported},
}

var bandNRExtLowerMSDR18Constraints = per.SizeRange(1, common.MaxLowerMSD_r18)

var bandNRExtLowerMSDENDCR18Constraints = per.SizeRange(1, common.MaxLowerMSD_r18)

const (
	BandNR_Ext_EnhancedChannelRaster_r18_Supported = 0
)

var bandNRExtEnhancedChannelRasterR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedChannelRaster_r18_Supported},
}

const (
	BandNR_Ext_FastBeamSweepingMultiRx_r18_N2 = 0
	BandNR_Ext_FastBeamSweepingMultiRx_r18_N4 = 1
	BandNR_Ext_FastBeamSweepingMultiRx_r18_N6 = 2
)

var bandNRExtFastBeamSweepingMultiRxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_FastBeamSweepingMultiRx_r18_N2, BandNR_Ext_FastBeamSweepingMultiRx_r18_N4, BandNR_Ext_FastBeamSweepingMultiRx_r18_N6},
}

const (
	BandNR_Ext_SimultaneousReceptionTwoQCL_r18_Supported = 0
)

var bandNRExtSimultaneousReceptionTwoQCLR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SimultaneousReceptionTwoQCL_r18_Supported},
}

const (
	BandNR_Ext_MeasEnhCAInterFreqFR2_r18_Supported = 0
)

var bandNRExtMeasEnhCAInterFreqFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MeasEnhCAInterFreqFR2_r18_Supported},
}

const (
	BandNR_Ext_Tci_StateSwitchInd_r18_Supported = 0
)

var bandNRExtTciStateSwitchIndR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Tci_StateSwitchInd_r18_Supported},
}

const (
	BandNR_Ext_AntennaArrayType_r18_Supported = 0
)

var bandNRExtAntennaArrayTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_AntennaArrayType_r18_Supported},
}

const (
	BandNR_Ext_LocationBasedCondHandoverATG_r18_Supported = 0
)

var bandNRExtLocationBasedCondHandoverATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_LocationBasedCondHandoverATG_r18_Supported},
}

var bandNRMaxOutputPowerATGR18Constraints = per.Constrained(1, 18)

const (
	BandNR_Ext_MeasValidationReportEMR_r18_Supported = 0
)

var bandNRExtMeasValidationReportEMRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MeasValidationReportEMR_r18_Supported},
}

const (
	BandNR_Ext_MeasValidationReportReselectionMeasurements_r18_Supported = 0
)

var bandNRExtMeasValidationReportReselectionMeasurementsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MeasValidationReportReselectionMeasurements_r18_Supported},
}

const (
	BandNR_Ext_EventA4BasedCondHandoverNES_r18_Supported = 0
)

var bandNRExtEventA4BasedCondHandoverNESR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EventA4BasedCondHandoverNES_r18_Supported},
}

const (
	BandNR_Ext_NesBasedCondHandoverWithDCI_r18_Supported = 0
)

var bandNRExtNesBasedCondHandoverWithDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_NesBasedCondHandoverWithDCI_r18_Supported},
}

const (
	BandNR_Ext_Rach_LessHandoverCG_r18_Supported = 0
)

var bandNRExtRachLessHandoverCGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Rach_LessHandoverCG_r18_Supported},
}

const (
	BandNR_Ext_Rach_LessHandoverDG_r18_Supported = 0
)

var bandNRExtRachLessHandoverDGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Rach_LessHandoverDG_r18_Supported},
}

const (
	BandNR_Ext_LocationBasedCondHandoverEMC_r18_Supported = 0
)

var bandNRExtLocationBasedCondHandoverEMCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_LocationBasedCondHandoverEMC_r18_Supported},
}

const (
	BandNR_Ext_Mt_CG_SDT_r18_Supported = 0
)

var bandNRExtMtCGSDTR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mt_CG_SDT_r18_Supported},
}

const (
	BandNR_Ext_PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18_Supported = 0
)

var bandNRExtPosSRSPreconfigureRRCInactiveInitialULBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18_Supported},
}

const (
	BandNR_Ext_PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18_Supported = 0
)

var bandNRExtPosSRSPreconfigureRRCInactiveOutsideInitialULBWPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18_Supported},
}

const (
	BandNR_Ext_Cg_SDT_PeriodicityExt_r18_Supported = 0
)

var bandNRExtCgSDTPeriodicityExtR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Cg_SDT_PeriodicityExt_r18_Supported},
}

const (
	BandNR_Ext_SupportOf2RxXR_r18_Supported = 0
)

var bandNRExtSupportOf2RxXRR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SupportOf2RxXR_r18_Supported},
}

const (
	BandNR_Ext_CondHandoverWithCandSCG_Change_r18_Supported = 0
)

var bandNRExtCondHandoverWithCandSCGChangeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_CondHandoverWithCandSCG_Change_r18_Supported},
}

var bandNRExtChannelBWDLNCRR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-100mhz"},
		{Name: "fr2-200mhz"},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz = 0
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz = 1
)

var bandNRExtChannelBWULNCRR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-100mhz"},
		{Name: "fr2-200mhz"},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz = 0
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz = 1
)

const (
	BandNR_Ext_Ncr_PDSCH_64QAM_FR2_r18_Supported = 0
)

var bandNRExtNcrPDSCH64QAMFR2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ncr_PDSCH_64QAM_FR2_r18_Supported},
}

const (
	BandNR_Ext_Ltm_MCG_IntraFreq_r18_Supported = 0
)

var bandNRExtLtmMCGIntraFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_MCG_IntraFreq_r18_Supported},
}

const (
	BandNR_Ext_Ltm_SCG_IntraFreq_r18_Supported = 0
)

var bandNRExtLtmSCGIntraFreqR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_SCG_IntraFreq_r18_Supported},
}

const (
	BandNR_Ext_EventA4BasedCondHandoverATG_r18_Supported = 0
)

var bandNRExtEventA4BasedCondHandoverATGR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EventA4BasedCondHandoverATG_r18_Supported},
}

const (
	BandNR_Ext_PosSRS_TxFH_WithTimeWindowForRedCap_r18_Supported = 0
)

var bandNRExtPosSRSTxFHWithTimeWindowForRedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_TxFH_WithTimeWindowForRedCap_r18_Supported},
}

const (
	BandNR_Ext_Ul_ResourceMutingCP_OFDM_r19_Supported = 0
)

var bandNRExtUlResourceMutingCPOFDMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ul_ResourceMutingCP_OFDM_r19_Supported},
}

const (
	BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_r19_Supported = 0
)

var bandNRExtUlResourceMutingDFTSOFDMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_r19_Supported},
}

const (
	BandNR_Ext_Ul_ResourceMutingCP_OFDM_CG_Type1_r19_N2 = 0
	BandNR_Ext_Ul_ResourceMutingCP_OFDM_CG_Type1_r19_N4 = 1
)

var bandNRExtUlResourceMutingCPOFDMCGType1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ul_ResourceMutingCP_OFDM_CG_Type1_r19_N2, BandNR_Ext_Ul_ResourceMutingCP_OFDM_CG_Type1_r19_N4},
}

const (
	BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19_N2 = 0
	BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19_N4 = 1
)

var bandNRExtUlResourceMutingDFTSOFDMCGType1R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19_N2, BandNR_Ext_Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19_N4},
}

const (
	BandNR_Ext_Od_SSB_NoAlwaysOn_RRC_r19_Supported = 0
)

var bandNRExtOdSSBNoAlwaysOnRRCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_NoAlwaysOn_RRC_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19_Supported = 0
)

var bandNRExtOdSSBNoAlwaysOnRRCMACCER19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_RRC_r19_Timec1    = 0
	BandNR_Ext_Od_SSB_AlwaysOn_RRC_r19_Timec1nc2 = 1
)

var bandNRExtOdSSBAlwaysOnRRCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_RRC_r19_Timec1, BandNR_Ext_Od_SSB_AlwaysOn_RRC_r19_Timec1nc2},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_RRC_Diff_r19_Supported = 0
)

var bandNRExtOdSSBAlwaysOnRRCDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_RRC_Diff_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_RRC_MAC_CE_r19_Supported = 0
)

var bandNRExtOdSSBAlwaysOnRRCMACCER19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_RRC_MAC_CE_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19_Supported = 0
)

var bandNRExtOdSSBAlwaysOnRRCMACCEDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_NoAlwaysOn_MAC_CE_r19_Explicit = 0
	BandNR_Ext_Od_SSB_NoAlwaysOn_MAC_CE_r19_Both     = 1
)

var bandNRExtOdSSBNoAlwaysOnMACCER19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_NoAlwaysOn_MAC_CE_r19_Explicit, BandNR_Ext_Od_SSB_NoAlwaysOn_MAC_CE_r19_Both},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_Diff_r19_Explicit = 0
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_Diff_r19_Both     = 1
)

var bandNRExtOdSSBAlwaysOnMACCEDiffR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_Diff_r19_Explicit, BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_Diff_r19_Both},
}

const (
	BandNR_Ext_Ssb_BurstPeriodicityAdaptation_r19_Supported = 0
)

var bandNRExtSsbBurstPeriodicityAdaptationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ssb_BurstPeriodicityAdaptation_r19_Supported},
}

const (
	BandNR_Ext_Rach_AdaptationTimeDomain_r19_Supported = 0
)

var bandNRExtRachAdaptationTimeDomainR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Rach_AdaptationTimeDomain_r19_Supported},
}

const (
	BandNR_Ext_Ltm_BeamIndicationJointTCI_CSI_RS_r19_Supported = 0
)

var bandNRExtLtmBeamIndicationJointTCICSIRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationJointTCI_CSI_RS_r19_Supported},
}

const (
	BandNR_Ext_Ltm_MAC_CE_JointTCI_CSI_RS_r19_Supported = 0
)

var bandNRExtLtmMACCEJointTCICSIRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_MAC_CE_JointTCI_CSI_RS_r19_Supported},
}

const (
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_CSI_RS_r19_Supported = 0
)

var bandNRExtLtmBeamIndicationSeparateTCICSIRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationSeparateTCI_CSI_RS_r19_Supported},
}

const (
	BandNR_Ext_Ltm_MAC_CE_SeparateTCI_CSI_RS_r19_Supported = 0
)

var bandNRExtLtmMACCESeparateTCICSIRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_MAC_CE_SeparateTCI_CSI_RS_r19_Supported},
}

const (
	BandNR_Ext_Ltm_WithoutCSI_IM_r19_Supported = 0
)

var bandNRExtLtmWithoutCSIIMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_WithoutCSI_IM_r19_Supported},
}

const (
	BandNR_Ext_Ssb_PeriodicityAddition_r19_Supported = 0
)

var bandNRExtSsbPeriodicityAdditionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ssb_PeriodicityAddition_r19_Supported},
}

const (
	BandNR_Ext_Pdcch_RepetitionType0_r19_Supported = 0
)

var bandNRExtPdcchRepetitionType0R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdcch_RepetitionType0_r19_Supported},
}

var bandNRPdcchRepetitionTypeOthersR19Constraints = per.Constrained(1, 2)

const (
	BandNR_Ext_Pdsch_RepetitionMsg4_r19_Supported = 0
)

var bandNRExtPdschRepetitionMsg4R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdsch_RepetitionMsg4_r19_Supported},
}

const (
	BandNR_Ext_Ntn_Collision_RedCap_r19_Supported = 0
)

var bandNRExtNtnCollisionRedCapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ntn_Collision_RedCap_r19_Supported},
}

const (
	BandNR_Ext_PosSRS_TxFH_WithTimeWindowForNonRedCap_r19_Supported = 0
)

var bandNRExtPosSRSTxFHWithTimeWindowForNonRedCapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PosSRS_TxFH_WithTimeWindowForNonRedCap_r19_Supported},
}

const (
	BandNR_Ext_Sr_TriggeredSSSG_Switching_r19_Supported = 0
)

var bandNRExtSrTriggeredSSSGSwitchingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sr_TriggeredSSSG_Switching_r19_Supported},
}

const (
	BandNR_Ext_Pdcch_RepetitionType0_TN_r19_Supported = 0
)

var bandNRExtPdcchRepetitionType0TNR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdcch_RepetitionType0_TN_r19_Supported},
}

var bandNRPdcchRepetitionTypeOthersTNR19Constraints = per.Constrained(1, 2)

var bandNRExtMprSingleCCR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "mpr-SingleCC-SingleValue-r19"},
		{Name: "mpr-SingleCC-MultipleValue-r19"},
	},
}

const (
	BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_SingleValue_r19   = 0
	BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_MultipleValue_r19 = 1
)

const (
	BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N2 = 0
	BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N4 = 1
	BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N6 = 2
)

var bandNRExtFastRxBSFMeasDelayReductionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N2, BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N4, BandNR_Ext_FastRx_BSF_MeasDelayReduction_r19_N6},
}

const (
	BandNR_Ext_FastSCellActivationEarlyMeas_r19_Supported = 0
)

var bandNRExtFastSCellActivationEarlyMeasR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_FastSCellActivationEarlyMeas_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_AdditionalProcessingTime_r19_Supported = 0
)

var bandNRExtOdSSBAdditionalProcessingTimeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AdditionalProcessingTime_r19_Supported},
}

const (
	BandNR_Ext_Aiml_BM_Case1_KnownRxBeam_r19_Supported = 0
)

var bandNRExtAimlBMCase1KnownRxBeamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Aiml_BM_Case1_KnownRxBeam_r19_Supported},
}

const (
	BandNR_Ext_Aiml_BM_Case2_KnownRxBeam_r19_Supported = 0
)

var bandNRExtAimlBMCase2KnownRxBeamR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Aiml_BM_Case2_KnownRxBeam_r19_Supported},
}

const (
	BandNR_Ext_Ntn_PowerBoosting_ERedCap_r19_Supported = 0
)

var bandNRExtNtnPowerBoostingERedCapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ntn_PowerBoosting_ERedCap_r19_Supported},
}

const (
	BandNR_Ext_EarlyCSI_AcquisitionWithoutCSI_IM_r19_Supported = 0
)

var bandNRExtEarlyCSIAcquisitionWithoutCSIIMR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EarlyCSI_AcquisitionWithoutCSI_IM_r19_Supported},
}

var bandNRChannelBWsDLFr1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

var bandNRChannelBWsDLFr2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var bandNRChannelBWsULFr1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

var bandNRChannelBWsULFr2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var bandNRExtChannelBWsDLV1590Fr1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

var bandNRExtChannelBWsDLV1590Fr2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var bandNRExtChannelBWsULV1590Fr1Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

var bandNRExtChannelBWsULV1590Fr2Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

var bandNRExtChannelBWDLIABR16Fr1100mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_15kHz_Supported = 0
)

var bandNRExtChannelBWDLIABR16Fr1100mhzScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_15kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_30kHz_Supported = 0
)

var bandNRExtChannelBWDLIABR16Fr1100mhzScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_30kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWDLIABR16Fr1100mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz_Scs_60kHz_Supported},
}

var bandNRExtChannelBWDLIABR16Fr2200mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWDLIABR16Fr2200mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz_Scs_60kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz_Scs_120kHz_Supported = 0
)

var bandNRExtChannelBWDLIABR16Fr2200mhzScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz_Scs_120kHz_Supported},
}

var bandNRExtChannelBWULIABR16Fr1100mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_15kHz_Supported = 0
)

var bandNRExtChannelBWULIABR16Fr1100mhzScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_15kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_30kHz_Supported = 0
)

var bandNRExtChannelBWULIABR16Fr1100mhzScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_30kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWULIABR16Fr1100mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz_Scs_60kHz_Supported},
}

var bandNRExtChannelBWULIABR16Fr2200mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWULIABR16Fr2200mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz_Scs_60kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz_Scs_120kHz_Supported = 0
)

var bandNRExtChannelBWULIABR16Fr2200mhzScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz_Scs_120kHz_Supported},
}

const (
	BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N1  = 0
	BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N2  = 1
	BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N4  = 2
	BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N8  = 3
	BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N12 = 4
)

var bandNRExtActiveConfiguredGrantR16MaxNumberConfigsPerBWPR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N1, BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N2, BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N4, BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N8, BandNR_Ext_ActiveConfiguredGrant_r16_MaxNumberConfigsPerBWP_r16_N12},
}

var bandNRExtMaxDurationDMRSBundlingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "fdd-r17", Optional: true},
		{Name: "tdd-r17", Optional: true},
	},
}

const (
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N4  = 0
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N8  = 1
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N16 = 2
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N32 = 3
)

var bandNRExtMaxDurationDMRSBundlingR17FddR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N4, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N8, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N16, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Fdd_r17_N32},
}

const (
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N2  = 0
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N4  = 1
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N8  = 2
	BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N16 = 3
)

var bandNRExtMaxDurationDMRSBundlingR17TddR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N2, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N4, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N8, BandNR_Ext_MaxDurationDMRS_Bundling_r17_Tdd_r17_N16},
}

const (
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N8   = 0
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N16  = 1
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N32  = 2
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N48  = 3
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N64  = 4
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N128 = 5
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N255 = 6
)

var bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSPerCCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N8, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N16, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N32, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N48, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N64, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N128, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_PerCC_r17_N255},
}

const (
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N8    = 0
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N16   = 1
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N32   = 2
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N64   = 3
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N128  = 4
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N256  = 5
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N512  = 6
	BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N1024 = 7
)

var bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSAcrossCCsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N8, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N16, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N32, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N64, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N128, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N256, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N512, BandNR_Ext_AperiodicCSI_RS_FastScellActivation_r17_MaxNumberAperiodicCSI_RS_AcrossCCs_r17_N1024},
}

const (
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N1 = 0
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N2 = 1
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N4 = 2
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N8 = 3
)

var bandNRExtEnhancedType3HARQCodebookFeedbackR17EnhancedType3HARQCodebooksR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N1, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N2, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N4, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_EnhancedType3_HARQ_Codebooks_r17_N8},
}

const (
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N1 = 0
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N2 = 1
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N3 = 2
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N4 = 3
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N5 = 4
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N6 = 5
	BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N7 = 6
)

var bandNRExtEnhancedType3HARQCodebookFeedbackR17MaxNumberPUCCHTransmissionsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N1, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N2, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N3, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N4, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N5, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N6, BandNR_Ext_EnhancedType3_HARQ_CodebookFeedback_r17_MaxNumberPUCCH_Transmissions_r17_N7},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_7 = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_5 = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_3 = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_1 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N1  = 4
)

var bandNRExtTriggeredHARQCodebookRetxR17MinHARQRetxOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_7, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_5, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_3, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N_1, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MinHARQ_Retx_Offset_r17_N1},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N4  = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N6  = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N8  = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N10 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N12 = 4
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N14 = 5
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N16 = 6
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N18 = 7
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N20 = 8
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N22 = 9
	BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N24 = 10
)

var bandNRExtTriggeredHARQCodebookRetxR17MaxHARQRetxOffsetR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N4, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N6, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N8, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N10, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N12, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N14, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N16, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N18, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N20, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N22, BandNR_Ext_TriggeredHARQ_CodebookRetx_r17_MaxHARQ_Retx_Offset_r17_N24},
}

const (
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17_Qam256  = 0
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17_Qam1024 = 1
)

var bandNRExtMaxModulationOrderForMulticastR17Fr1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17_Qam256, BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17_Qam1024},
}

const (
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17_Qam64  = 0
	BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17_Qam256 = 1
)

var bandNRExtMaxModulationOrderForMulticastR17Fr2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17_Qam64, BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17_Qam256},
}

const (
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N1  = 0
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N2  = 1
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N4  = 2
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N8  = 3
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N16 = 4
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N32 = 5
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N64 = 6
)

var bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N1, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N2, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N4, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N8, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N16, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N32, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResources_r17_N64},
}

const (
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N1  = 0
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N2  = 1
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N3  = 2
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N4  = 3
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N5  = 4
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N6  = 5
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N8  = 6
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N10 = 7
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N12 = 8
	BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N14 = 9
)

var bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N1, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N2, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N3, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N4, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N5, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N6, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N8, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N10, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N12, BandNR_Ext_Srs_SemiPersistent_PosResourcesRRC_Inactive_r17_MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17_N14},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_SdType1 = 0
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_SdType2 = 1
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_Both    = 2
)

var bandNRExtSpatialAdaptationCSIFeedbackR18CsiFeedbackTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_SdType1, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_SdType2, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_CsiFeedbackType_r18_Both},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N128},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N128},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_SdType1 = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_SdType2 = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_Both    = 2
)

var bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18CsiFeedbackTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_SdType1, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_SdType2, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_CsiFeedbackType_r18_Both},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_SdType1 = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_SdType2 = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_Both    = 2
)

var bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18CsiFeedbackTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_SdType1, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_SdType2, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_CsiFeedbackType_r18_Both},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType1_Resource_r18_N128},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_SdType2_Resource_r18_N128},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_SdType1 = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_SdType2 = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_Both    = 2
)

var bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18CsiFeedbackTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_SdType1, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_SdType2, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_CsiFeedbackType_r18_Both},
}

const (
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_SpatialAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtPowerAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_PowerAdaptation_CSI_Feedback_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtPowerAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUSCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtPowerAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_PowerAdaptation_CSI_FeedbackAperiodic_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8   = 0
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16  = 1
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24  = 2
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32  = 3
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64  = 4
	BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128 = 5
)

var bandNRExtPowerAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N8, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N16, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N24, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N32, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N64, BandNR_Ext_PowerAdaptation_CSI_FeedbackPUCCH_r18_MaxNumberTotalCSI_ResourcePerCC_r18_N128},
}

const (
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N8   = 0
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N12  = 1
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N16  = 2
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N24  = 3
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N32  = 4
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N48  = 5
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N64  = 6
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N128 = 7
)

var bandNRExtLtmBeamIndicationJointTCIR18MaxNumberJointTCIPerCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N8, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N12, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N16, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N24, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N32, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N48, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N64, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_MaxNumberJointTCI_PerCell_r18_N128},
}

const (
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtLtmBeamIndicationJointTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Ltm_BeamIndicationJointTCI_r18_Qcl_Resource_r18_Both},
}

const (
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtDummyLtmMACCEJointTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Both},
}

const (
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N1  = 0
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N2  = 1
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N3  = 2
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N4  = 3
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N8  = 4
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N16 = 5
	BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N32 = 6
)

var bandNRExtDummyLtmMACCEJointTCIR18MaxNumberJointTCIAcrossCellsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N1, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N2, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N3, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N4, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N8, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N16, BandNR_Ext_Dummy_Ltm_MAC_CE_JointTCI_r18_MaxNumberJointTCI_AcrossCells_r18_N32},
}

const (
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N4   = 0
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N8   = 1
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N12  = 2
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N16  = 3
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N24  = 4
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N32  = 5
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N48  = 6
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N64  = 7
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N128 = 8
)

var bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberDLTCIPerCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N4, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N8, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N12, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N16, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N24, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N32, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N48, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N64, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberDL_TCI_PerCell_r18_N128},
}

const (
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N4  = 0
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N8  = 1
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N12 = 2
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N16 = 3
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N24 = 4
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N32 = 5
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N48 = 6
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N64 = 7
)

var bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberULTCIPerCellR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N4, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N8, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N12, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N16, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N24, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N32, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N48, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_MaxNumberUL_TCI_PerCell_r18_N64},
}

const (
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtLtmBeamIndicationSeparateTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Ltm_BeamIndicationSeparateTCI_r18_Qcl_Resource_r18_Both},
}

const (
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtDummyLtmMACCESeparateTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Both},
}

const (
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N1  = 0
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N2  = 1
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N4  = 2
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N8  = 3
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N16 = 4
)

var bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberDLTCIAcrossCellsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N1, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N2, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N4, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N8, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberDL_TCI_AcrossCells_r18_N16},
}

const (
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N1  = 0
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N2  = 1
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N4  = 2
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N8  = 3
	BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N16 = 4
)

var bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberULTCIAcrossCellsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N1, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N2, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N4, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N8, BandNR_Ext_Dummy_Ltm_MAC_CE_SeparateTCI_r18_MaxNumberUL_TCI_AcrossCells_r18_N16},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_7 = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_5 = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_3 = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_1 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N1  = 4
)

var bandNRExtTriggeredHARQCodebookRetxDCI13R18MinHARQRetxOffsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_7, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_5, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_3, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N_1, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MinHARQ_Retx_Offset_r18_N1},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N4  = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N6  = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N8  = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N10 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N12 = 4
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N14 = 5
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N16 = 6
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N18 = 7
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N20 = 8
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N22 = 9
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N24 = 10
)

var bandNRExtTriggeredHARQCodebookRetxDCI13R18MaxHARQRetxOffsetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N4, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N6, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N8, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N10, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N12, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N14, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N16, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N18, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N20, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N22, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_r18_MaxHARQ_Retx_Offset_r18_N24},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "minBeamApplicationTime-r18"},
		{Name: "maxActivatedTCI-PerCC-r18", Optional: true},
	},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r18"},
		{Name: "fr2-r18"},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18 = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18 = 1
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym70},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym70},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym70},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym70, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym84, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym98, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym112, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym224, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym336},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym70, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym84, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym98, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym112, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym224, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym336},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "minBeamApplicationTime-r18"},
		{Name: "maxActivatedDL-TCI-PerCC-r18", Optional: true},
		{Name: "maxActivatedUL-TCI-PerCC-r18", Optional: true},
	},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r18"},
		{Name: "fr2-r18"},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18 = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18 = 1
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_15kHz_r18_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_30kHz_r18_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18_Scs_60kHz_r18_Sym336},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_60kHz_r18_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18_Scs_120kHz_r18_Sym336},
}

const (
	BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N1  = 0
	BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N2  = 1
	BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N4  = 2
	BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N8  = 3
	BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N12 = 4
)

var bandNRExtMultiPUSCHActiveConfiguredGrantR18MaxNumberConfigsPerBWPConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N1, BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N2, BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N4, BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N8, BandNR_Ext_MultiPUSCH_ActiveConfiguredGrant_r18_MaxNumberConfigsPerBWP_N12},
}

const (
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInRE_r18_OneSymbolNoOverlap  = 0
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInRE_r18_SomeOrAllSymOverlap = 1
)

var bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInRER18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInRE_r18_OneSymbolNoOverlap, BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInRE_r18_SomeOrAllSymOverlap},
}

const (
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInSymbol_r18_Symbol2     = 0
	BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInSymbol_r18_Symbol1And2 = 1
)

var bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInSymbolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInSymbol_r18_Symbol2, BandNR_Ext_Nr_PDCCH_OverlapLTE_CRS_RE_r18_OverlapInSymbol_r18_Symbol1And2},
}

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r18", Optional: true},
		{Name: "scs-30kHz-r18", Optional: true},
		{Name: "scs-60kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N1  = 0
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N2  = 1
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N4  = 2
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N6  = 3
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N8  = 4
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N12 = 5
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N16 = 6
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N24 = 7
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N32 = 8
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N48 = 9
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N64 = 10
)

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs15kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N1, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N2, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N4, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N6, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N8, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N12, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N16, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N24, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N32, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N48, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_15kHz_r18_N64},
}

const (
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N1  = 0
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N2  = 1
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N4  = 2
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N6  = 3
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N8  = 4
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N12 = 5
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N16 = 6
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N24 = 7
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N32 = 8
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N48 = 9
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N64 = 10
)

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs30kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N1, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N2, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N4, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N6, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N8, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N12, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N16, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N24, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N32, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N48, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_30kHz_r18_N64},
}

const (
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N1  = 0
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N2  = 1
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N4  = 2
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N6  = 3
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N8  = 4
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N12 = 5
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N16 = 6
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N24 = 7
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N32 = 8
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N48 = 9
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N64 = 10
)

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N1, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N2, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N4, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N6, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N8, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N12, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N16, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N24, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N32, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N48, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr1_r18_Scs_60kHz_r18_N64},
}

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N1  = 0
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N2  = 1
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N4  = 2
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N6  = 3
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N8  = 4
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N12 = 5
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N16 = 6
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N24 = 7
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N32 = 8
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N48 = 9
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N64 = 10
)

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N1, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N2, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N4, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N6, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N8, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N12, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N16, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N24, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N32, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N48, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_60kHz_r18_N64},
}

const (
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N1  = 0
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N2  = 1
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N4  = 2
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N6  = 3
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N8  = 4
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N12 = 5
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N16 = 6
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N24 = 7
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N32 = 8
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N48 = 9
	BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N64 = 10
)

var bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N1, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N2, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N4, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N6, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N8, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N12, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N16, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N24, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N32, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N48, BandNR_Ext_Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18_Fr2_r18_Scs_120kHz_r18_N64},
}

const (
	BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N1 = 0
	BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N2 = 1
	BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N4 = 2
	BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N6 = 3
)

var bandNRExtBeamSweepingFactorReductionR18ReduceForCellDetectionConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N1, BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N2, BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N4, BandNR_Ext_BeamSweepingFactorReduction_r18_ReduceForCellDetection_N6},
}

const (
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N2  = 0
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N3  = 1
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N4  = 2
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N5  = 3
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N6  = 4
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N7  = 5
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N8  = 6
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N9  = 7
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N10 = 8
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N11 = 9
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N12 = 10
	BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N16 = 11
)

var bandNRExtLtmFastProcessingConfigR18MaxNumberStoredConfigCellsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N2, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N3, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N4, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N5, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N6, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N7, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N8, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N9, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N10, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N11, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N12, BandNR_Ext_Ltm_FastProcessingConfig_r18_MaxNumberStoredConfigCells_r18_N16},
}

var bandNRExtChannelBWDLNCRR18Fr1100mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_15kHz_Supported = 0
)

var bandNRExtChannelBWDLNCRR18Fr1100mhzScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_15kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_30kHz_Supported = 0
)

var bandNRExtChannelBWDLNCRR18Fr1100mhzScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_30kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWDLNCRR18Fr1100mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz_Scs_60kHz_Supported},
}

var bandNRExtChannelBWDLNCRR18Fr2200mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWDLNCRR18Fr2200mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz_Scs_60kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz_Scs_120kHz_Supported = 0
)

var bandNRExtChannelBWDLNCRR18Fr2200mhzScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz_Scs_120kHz_Supported},
}

var bandNRExtChannelBWULNCRR18Fr1100mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz", Optional: true},
		{Name: "scs-30kHz", Optional: true},
		{Name: "scs-60kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_15kHz_Supported = 0
)

var bandNRExtChannelBWULNCRR18Fr1100mhzScs15kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_15kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_30kHz_Supported = 0
)

var bandNRExtChannelBWULNCRR18Fr1100mhzScs30kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_30kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWULNCRR18Fr1100mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz_Scs_60kHz_Supported},
}

var bandNRExtChannelBWULNCRR18Fr2200mhzConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz", Optional: true},
		{Name: "scs-120kHz", Optional: true},
	},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz_Scs_60kHz_Supported = 0
)

var bandNRExtChannelBWULNCRR18Fr2200mhzScs60kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz_Scs_60kHz_Supported},
}

const (
	BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz_Scs_120kHz_Supported = 0
)

var bandNRExtChannelBWULNCRR18Fr2200mhzScs120kHzConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz_Scs_120kHz_Supported},
}

const (
	BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtLtmMACCEJointTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Ltm_MAC_CE_JointTCI_r18_Qcl_Resource_r18_Both},
}

const (
	BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Ssb  = 0
	BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Trs  = 1
	BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Both = 2
)

var bandNRExtLtmMACCESeparateTCIR18QclResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Ssb, BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Trs, BandNR_Ext_Ltm_MAC_CE_SeparateTCI_r18_Qcl_Resource_r18_Both},
}

var bandNRExtSbfdAwareR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "numOfPartialPRG-r19"},
		{Name: "timeline-r19"},
		{Name: "sbfd-Rx-r19", Optional: true},
		{Name: "sbfd-Tx-r19", Optional: true},
		{Name: "sbfd-OneRACH-r19", Optional: true},
		{Name: "sbfd-TwoRACH-r19", Optional: true},
		{Name: "preambleRepetitionOneRACH-r19", Optional: true},
		{Name: "preambleRepetitionTwoRACH-r19", Optional: true},
	},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_NumOfPartialPRG_r19_N2 = 0
	BandNR_Ext_Sbfd_Aware_r19_NumOfPartialPRG_r19_N4 = 1
)

var bandNRExtSbfdAwareR19NumOfPartialPRGR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_NumOfPartialPRG_r19_N2, BandNR_Ext_Sbfd_Aware_r19_NumOfPartialPRG_r19_N4},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_Full    = 0
	BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_Partial = 1
	BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_No      = 2
)

var bandNRExtSbfdAwareR19TimelineR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_Full, BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_Partial, BandNR_Ext_Sbfd_Aware_r19_Timeline_r19_No},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_Sbfd_Rx_r19_Supported = 0
)

var bandNRExtSbfdAwareR19SbfdRxR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_Sbfd_Rx_r19_Supported},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_Sbfd_Tx_r19_Supported = 0
)

var bandNRExtSbfdAwareR19SbfdTxR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_Sbfd_Tx_r19_Supported},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_Sbfd_OneRACH_r19_Supported = 0
)

var bandNRExtSbfdAwareR19SbfdOneRACHR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_Sbfd_OneRACH_r19_Supported},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_Sbfd_TwoRACH_r19_Supported = 0
)

var bandNRExtSbfdAwareR19SbfdTwoRACHR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_Sbfd_TwoRACH_r19_Supported},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_PreambleRepetitionOneRACH_r19_Supported = 0
)

var bandNRExtSbfdAwareR19PreambleRepetitionOneRACHR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_PreambleRepetitionOneRACH_r19_Supported},
}

const (
	BandNR_Ext_Sbfd_Aware_r19_PreambleRepetitionTwoRACH_r19_Supported = 0
)

var bandNRExtSbfdAwareR19PreambleRepetitionTwoRACHR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Sbfd_Aware_r19_PreambleRepetitionTwoRACH_r19_Supported},
}

var bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberOfConfiguredMeasResourceAcrossCC-r19"},
		{Name: "maxNumberOfAperiodicReportSettingPerBWP-r19"},
		{Name: "maxNumberOfConfiguredMeasResourcePerCC-r19"},
		{Name: "maxNumberOfSimultaneousMeasResourcePerCC-r19"},
		{Name: "periodicReporting-r19", Optional: true},
		{Name: "semiPersistentMeasResource-r19", Optional: true},
	},
}

const (
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N8  = 0
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N16 = 1
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N32 = 2
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N64 = 3
)

var bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N8, BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N16, BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N32, BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N64},
}

const (
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_PeriodicReporting_r19_Supported = 0
)

var bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19PeriodicReportingR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_PeriodicReporting_r19_Supported},
}

const (
	BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_SemiPersistentMeasResource_r19_Supported = 0
)

var bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_CLI_RSSI_MeasAndAperiodicReporting_r19_SemiPersistentMeasResource_r19_Supported},
}

var bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberOfConfiguredMeasResourceAcrossCC-r19"},
		{Name: "maxNumberOfConfiguredMeasResourcePerCC-r19"},
		{Name: "maxNumberOfSimultaneousMeasResourcePerCC-r19"},
		{Name: "maxNumberOfAperiodicReportSettingPerBWP-r19"},
		{Name: "maxNumberOfMeasResourceAcrossCCWithinSlot-r19"},
		{Name: "maxNumberOfAperiodicMeasResourceAcrossCC-r19"},
		{Name: "fdm-Reception-r19", Optional: true},
		{Name: "semiPersistentMeasResource-r19", Optional: true},
	},
}

const (
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N4  = 0
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N8  = 1
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N16 = 2
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N32 = 3
)

var bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N4, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N8, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N16, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfConfiguredMeasResourceAcrossCC_r19_N32},
}

const (
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N1 = 0
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N2 = 1
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N4 = 2
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N8 = 3
)

var bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfMeasResourceAcrossCCWithinSlotR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N1, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N2, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N4, BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_MaxNumberOfMeasResourceAcrossCCWithinSlot_r19_N8},
}

const (
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_Fdm_Reception_r19_Supported = 0
)

var bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19FdmReceptionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_Fdm_Reception_r19_Supported},
}

const (
	BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_SemiPersistentMeasResource_r19_Supported = 0
)

var bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_L1_SRS_RSRP_MeasAndAperiodicReporting_r19_SemiPersistentMeasResource_r19_Supported},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_TimeRelation_r19_Timec1    = 0
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_TimeRelation_r19_Timec1nc2 = 1
)

var bandNRExtOdSSBAlwaysOnMACCER19TimeRelationR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_TimeRelation_r19_Timec1, BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_TimeRelation_r19_Timec1nc2},
}

const (
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_DeactivationScheme_r19_Explicit = 0
	BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_DeactivationScheme_r19_Both     = 1
)

var bandNRExtOdSSBAlwaysOnMACCER19DeactivationSchemeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_DeactivationScheme_r19_Explicit, BandNR_Ext_Od_SSB_AlwaysOn_MAC_CE_r19_DeactivationScheme_r19_Both},
}

const (
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N1   = 0
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N2   = 1
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N4   = 2
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N8   = 3
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N12  = 4
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N16  = 5
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N24  = 6
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N32  = 7
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N48  = 8
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N64  = 9
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N128 = 10
)

var bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N1, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N2, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N4, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N8, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N12, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N16, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N24, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N32, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N48, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N64, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfCSI_RS_Ports_r19_N128},
}

const (
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1  = 0
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2  = 1
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4  = 2
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8  = 3
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12 = 4
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16 = 5
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24 = 6
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32 = 7
)

var bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfNZPCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24, BandNR_Ext_Ltm_CSI_RS_CSI_IM_Periodic_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32},
}

const (
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N1   = 0
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N2   = 1
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N4   = 2
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N8   = 3
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N12  = 4
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N16  = 5
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N24  = 6
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N32  = 7
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N48  = 8
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N64  = 9
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N128 = 10
)

var bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N1, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N2, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N4, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N8, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N12, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N16, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N24, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N32, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N48, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N64, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfCSI_RS_Ports_r19_N128},
}

const (
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1  = 0
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2  = 1
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4  = 2
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8  = 3
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12 = 4
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16 = 5
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24 = 6
	BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32 = 7
)

var bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfNZPCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24, BandNR_Ext_Ltm_CSI_RS_CSI_IM_SP_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32},
}

const (
	BandNR_Ext_Pusch_InterSlotOCC_r19_OccLength_r19_N2     = 0
	BandNR_Ext_Pusch_InterSlotOCC_r19_OccLength_r19_N2And4 = 1
)

var bandNRExtPuschInterSlotOCCR19OccLengthR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pusch_InterSlotOCC_r19_OccLength_r19_N2, BandNR_Ext_Pusch_InterSlotOCC_r19_OccLength_r19_N2And4},
}

const (
	BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Gso  = 0
	BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Ngso = 1
	BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Both = 2
)

var bandNRExtPuschInterSlotOCCR19SatelliteOrbitR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Gso, BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Ngso, BandNR_Ext_Pusch_InterSlotOCC_r19_SatelliteOrbit_r19_Both},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_7 = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_5 = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_3 = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_1 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N1  = 4
)

var bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MinHARQRetxOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_7, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_5, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_3, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N_1, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MinHARQ_Retx_Offset_r19_N1},
}

const (
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N4  = 0
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N6  = 1
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N8  = 2
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N10 = 3
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N12 = 4
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N14 = 5
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N16 = 6
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N18 = 7
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N20 = 8
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N22 = 9
	BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N24 = 10
)

var bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MaxHARQRetxOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N4, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N6, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N8, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N10, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N12, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N14, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N16, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N18, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N20, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N22, BandNR_Ext_TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19_MaxHARQ_Retx_Offset_r19_N24},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "minBeamApplicationTime-r19"},
		{Name: "maxActivatedTCI-PerCC-r19", Optional: true},
	},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r19"},
		{Name: "fr2-r19"},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19 = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19 = 1
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r19", Optional: true},
		{Name: "scs-30kHz-r19", Optional: true},
		{Name: "scs-60kHz-r19", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym70},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym70},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym1  = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym2  = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym4  = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym7  = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym14 = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym28 = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym42 = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym56 = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym70 = 8
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym70},
}

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r19", Optional: true},
		{Name: "scs-120kHz-r19", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym70, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym84, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym98, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym112, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym224, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym336},
}

const (
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym1, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym2, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym4, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym7, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym14, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym28, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym42, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym56, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym70, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym84, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym98, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym112, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym224, BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym336},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "minBeamApplicationTime-r19"},
		{Name: "maxActivatedDL-TCI-PerCC-r19", Optional: true},
		{Name: "maxActivatedUL-TCI-PerCC-r19", Optional: true},
	},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r19"},
		{Name: "fr2-r19"},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19 = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19 = 1
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r19", Optional: true},
		{Name: "scs-30kHz-r19", Optional: true},
		{Name: "scs-60kHz-r19", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_15kHz_r19_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_30kHz_r19_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19_Scs_60kHz_r19_Sym336},
}

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r19", Optional: true},
		{Name: "scs-120kHz-r19", Optional: true},
	},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_60kHz_r19_Sym336},
}

const (
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym1   = 0
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym2   = 1
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym4   = 2
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym7   = 3
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym14  = 4
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym28  = 5
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym42  = 6
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym56  = 7
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym70  = 8
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym84  = 9
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym98  = 10
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym112 = 11
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym224 = 12
	BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym336 = 13
)

var bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym1, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym2, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym4, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym7, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym14, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym28, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym42, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym56, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym70, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym84, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym98, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym112, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym224, BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19_Scs_120kHz_r19_Sym336},
}

const (
	BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_SingleValue_r19_Supported = 0
)

var bandNRExtMprSingleCCR19MprSingleCCSingleValueR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_SingleValue_r19_Supported},
}

const (
	BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_MultipleValue_r19_Supported = 0
)

var bandNRExtMprSingleCCR19MprSingleCCMultipleValueR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_MultipleValue_r19_Supported},
}

const (
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N1   = 0
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N2   = 1
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N4   = 2
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N8   = 3
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N12  = 4
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N16  = 5
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N24  = 6
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N32  = 7
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N48  = 8
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N64  = 9
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N128 = 10
)

var bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N1, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N2, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N4, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N8, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N12, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N16, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N24, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N32, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N48, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N64, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfCSI_RS_Ports_r19_N128},
}

const (
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1  = 0
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2  = 1
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4  = 2
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8  = 3
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12 = 4
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16 = 5
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24 = 6
	BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32 = 7
)

var bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfNZPCSIRSPortsR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N1, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N2, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N4, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N8, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N12, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N16, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N24, BandNR_Ext_EarlyCSI_AcquisitionWithCSI_IM_r19_MaxNumOfNZP_CSI_RS_Ports_r19_N32},
}

const (
	BandNR_Ext_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option1 = 0
	BandNR_Ext_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option2 = 1
)

var bandNRExtEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option1, BandNR_Ext_EnableTx_RxDuringMeasGap_r19_IndicationField_r19_Option2},
}

const (
	BandNR_Ext_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms5 = 0
	BandNR_Ext_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms3 = 1
)

var bandNRExtEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandNR_Ext_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms5, BandNR_Ext_EnableTx_RxDuringMeasGap_r19_MinimumTimeOffset_r19_Ms3},
}

type BandNR struct {
	BandNR                         FreqBandIndicatorNR
	ModifiedMPR_Behaviour          *per.BitString
	Mimo_ParametersPerBand         *MIMO_ParametersPerBand
	ExtendedCP                     *int64
	MultipleTCI                    *int64
	Bwp_WithoutRestriction         *int64
	Bwp_SameNumerology             *int64
	Bwp_DiffNumerology             *int64
	CrossCarrierScheduling_SameSCS *int64
	Pdsch_256QAM_FR2               *int64
	Pusch_256QAM                   *int64
	Ue_PowerClass                  *int64
	RateMatchingLTE_CRS            *int64
	ChannelBWs_DL                  *struct {
		Choice int
		Fr1    *struct {
			Scs_15kHz *per.BitString
			Scs_30kHz *per.BitString
			Scs_60kHz *per.BitString
		}
		Fr2 *struct {
			Scs_60kHz  *per.BitString
			Scs_120kHz *per.BitString
		}
	}
	ChannelBWs_UL *struct {
		Choice int
		Fr1    *struct {
			Scs_15kHz *per.BitString
			Scs_30kHz *per.BitString
			Scs_60kHz *per.BitString
		}
		Fr2 *struct {
			Scs_60kHz  *per.BitString
			Scs_120kHz *per.BitString
		}
	}
	MaxUplinkDutyCycle_PC2_FR1 *int64
	Pucch_SpatialRelInfoMAC_CE *int64
	PowerBoosting_Pi2BPSK      *int64
	MaxUplinkDutyCycle_FR2     *int64
	ChannelBWs_DL_v1590        *struct {
		Choice int
		Fr1    *struct {
			Scs_15kHz *per.BitString
			Scs_30kHz *per.BitString
			Scs_60kHz *per.BitString
		}
		Fr2 *struct {
			Scs_60kHz  *per.BitString
			Scs_120kHz *per.BitString
		}
	}
	ChannelBWs_UL_v1590 *struct {
		Choice int
		Fr1    *struct {
			Scs_15kHz *per.BitString
			Scs_30kHz *per.BitString
			Scs_60kHz *per.BitString
		}
		Fr2 *struct {
			Scs_60kHz  *per.BitString
			Scs_120kHz *per.BitString
		}
	}
	AsymmetricBandwidthCombinationSet       *per.BitString
	SharedSpectrumChAccessParamsPerBand_r16 *SharedSpectrumChAccessParamsPerBand_r16
	CancelOverlappingPUSCH_r16              *int64
	MultipleRateMatchingEUTRA_CRS_r16       *struct {
		MaxNumberPatterns_r16            int64
		MaxNumberNon_OverlapPatterns_r16 int64
	}
	OverlapRateMatchingEUTRA_CRS_r16  *int64
	Pdsch_MappingTypeB_Alt_r16        *int64
	OneSlotPeriodicTRS_r16            *int64
	Olpc_SRS_Pos_r16                  *OLPC_SRS_Pos_r16
	SpatialRelationsSRS_Pos_r16       *SpatialRelationsSRS_Pos_r16
	SimulSRS_MIMO_TransWithinBand_r16 *int64
	ChannelBW_DL_IAB_r16              *struct {
		Choice     int
		Fr1_100mhz *struct {
			Scs_15kHz *int64
			Scs_30kHz *int64
			Scs_60kHz *int64
		}
		Fr2_200mhz *struct {
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
	}
	ChannelBW_UL_IAB_r16 *struct {
		Choice     int
		Fr1_100mhz *struct {
			Scs_15kHz *int64
			Scs_30kHz *int64
			Scs_60kHz *int64
		}
		Fr2_200mhz *struct {
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
	}
	RasterShift7dot5_IAB_r16             *int64
	Ue_PowerClass_v1610                  *int64
	CondHandover_r16                     *int64
	CondHandoverFailure_r16              *int64
	CondHandoverTwoTriggerEvents_r16     *int64
	CondPSCellChange_r16                 *int64
	CondPSCellChangeTwoTriggerEvents_r16 *int64
	Mpr_PowerBoost_FR2_r16               *int64
	ActiveConfiguredGrant_r16            *struct {
		MaxNumberConfigsPerBWP_r16 int64
		MaxNumberConfigsAllCC_r16  int64
	}
	JointReleaseConfiguredGrantType2_r16 *int64
	Sps_r16                              *struct {
		MaxNumberConfigsPerBWP_r16 int64
		MaxNumberConfigsAllCC_r16  int64
	}
	JointReleaseSPS_r16                        *int64
	SimulSRS_TransWithinBand_r16               *int64
	Trs_AdditionalBandwidth_r16                *int64
	HandoverIntraF_IAB_r16                     *int64
	SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 *SimulSRS_ForAntennaSwitching_r16
	SharedSpectrumChAccessParamsPerBand_v1630  *SharedSpectrumChAccessParamsPerBand_v1630
	HandoverUTRA_FDD_r16                       *int64
	EnhancedUL_TransientPeriod_r16             *int64
	SharedSpectrumChAccessParamsPerBand_v1640  *SharedSpectrumChAccessParamsPerBand_v1640
	Type1_PUSCH_RepetitionMultiSlots_v1650     *int64
	Type2_PUSCH_RepetitionMultiSlots_v1650     *int64
	Pusch_RepetitionMultiSlots_v1650           *int64
	ConfiguredUL_GrantType1_v1650              *int64
	ConfiguredUL_GrantType2_v1650              *int64
	SharedSpectrumChAccessParamsPerBand_v1650  *SharedSpectrumChAccessParamsPerBand_v1650
	EnhancedSkipUplinkTxConfigured_v1660       *int64
	EnhancedSkipUplinkTxDynamic_v1660          *int64
	MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16     *int64
	TxDiversity_r16                            *int64
	Pdsch_1024QAM_FR1_r17                      *int64
	Ue_PowerClass_v1700                        *int64
	Fr2_2_AccessParamsPerBand_r17              *FR2_2_AccessParamsPerBand_r17
	Rlm_Relaxation_r17                         *int64
	Bfd_Relaxation_r17                         *int64
	Cg_SDT_r17                                 *int64
	LocationBasedCondHandover_r17              *int64
	TimeBasedCondHandover_r17                  *int64
	EventA4BasedCondHandover_r17               *int64
	Mn_InitiatedCondPSCellChangeNRDC_r17       *int64
	Sn_InitiatedCondPSCellChangeNRDC_r17       *int64
	Pdcch_SkippingWithoutSSSG_r17              *int64
	Sssg_Switching_1BitInd_r17                 *int64
	Sssg_Switching_2BitInd_r17                 *int64
	Pdcch_SkippingWithSSSG_r17                 *int64
	SearchSpaceSetGrp_SwitchCap2_r17           *int64
	UplinkPreCompensation_r17                  *int64
	Uplink_TA_Reporting_r17                    *int64
	Max_HARQ_ProcessNumber_r17                 *int64
	Type2_HARQ_Codebook_r17                    *int64
	Type1_HARQ_Codebook_r17                    *int64
	Type3_HARQ_Codebook_r17                    *int64
	Ue_Specific_K_Offset_r17                   *int64
	MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17  *int64
	MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17  *int64
	ParallelPRS_MeasRRC_Inactive_r17           *int64
	Nr_UE_TxTEG_ID_MaxSupport_r17              *int64
	Prs_ProcessingRRC_Inactive_r17             *int64
	Prs_ProcessingWindowType1A_r17             *int64
	Prs_ProcessingWindowType1B_r17             *int64
	Prs_ProcessingWindowType2_r17              *int64
	Srs_AllPosResourcesRRC_Inactive_r17        *SRS_AllPosResourcesRRC_Inactive_r17
	Olpc_SRS_PosRRC_Inactive_r17               *OLPC_SRS_Pos_r16
	SpatialRelationsSRS_PosRRC_Inactive_r17    *SpatialRelationsSRS_Pos_r16
	MaxNumberPUSCH_TypeA_Repetition_r17        *int64
	PuschTypeA_RepetitionsAvailSlot_r17        *int64
	Tb_ProcessingMultiSlotPUSCH_r17            *int64
	Tb_ProcessingRepMultiSlotPUSCH_r17         *int64
	MaxDurationDMRS_Bundling_r17               *struct {
		Fdd_r17 *int64
		Tdd_r17 *int64
	}
	Pusch_RepetitionMsg3_r17                   *int64
	SharedSpectrumChAccessParamsPerBand_v1710  *SharedSpectrumChAccessParamsPerBand_v1710
	ParallelMeasurementWithoutRestriction_r17  *int64
	MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 *int64
	K1_RangeExtension_r17                      *int64
	AperiodicCSI_RS_FastScellActivation_r17    *struct {
		MaxNumberAperiodicCSI_RS_PerCC_r17     int64
		MaxNumberAperiodicCSI_RS_AcrossCCs_r17 int64
	}
	AperiodicCSI_RS_AdditionalBandwidth_r17      *int64
	Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17       *int64
	HalfDuplexFDD_TypeA_RedCap_r17               *int64
	PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17
	ChannelBWs_DL_SCS_480kHz_FR2_2_r17           *per.BitString
	ChannelBWs_UL_SCS_480kHz_FR2_2_r17           *per.BitString
	ChannelBWs_DL_SCS_960kHz_FR2_2_r17           *per.BitString
	ChannelBWs_UL_SCS_960kHz_FR2_2_r17           *per.BitString
	Ul_GapFR2_r17                                *int64
	OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17   *int64
	OneShotHARQ_FeedbackPhy_Priority_r17         *int64
	EnhancedType3_HARQ_CodebookFeedback_r17      *struct {
		EnhancedType3_HARQ_Codebooks_r17 int64
		MaxNumberPUCCH_Transmissions_r17 int64
	}
	TriggeredHARQ_CodebookRetx_r17 *struct {
		MinHARQ_Retx_Offset_r17 int64
		MaxHARQ_Retx_Offset_r17 int64
	}
	Ue_OneShotUL_TimingAdj_r17                       *int64
	Pucch_Repetition_F0_2_r17                        *int64
	Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17  *int64
	Mux_HARQ_ACK_DiffPriorities_r17                  *int64
	Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17       *int64
	Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 *int64
	MaxNumberG_RNTI_r17                              *int64
	DynamicMulticastDCI_Format4_2_r17                *int64
	MaxModulationOrderForMulticast_r17               *struct {
		Choice  int
		Fr1_r17 *int64
		Fr2_r17 *int64
	}
	DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 *int64
	DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17   *int64
	Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17               *int64
	Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17           *int64
	MaxNumberG_CS_RNTI_r17                                         *int64
	Re_LevelRateMatchingForMulticast_r17                           *int64
	Pdsch_1024QAM_2MIMO_FR1_r17                                    *int64
	Prs_MeasurementWithoutMG_r17                                   *int64
	MaxNumber_NGSO_SatellitesPerCarrier_r17                        *int64
	Prs_ProcessingCapabilityOutsideMGinPPW_r17                     []PRS_ProcessingCapabilityOutsideMGinPPWperType_r17
	Srs_SemiPersistent_PosResourcesRRC_Inactive_r17                *struct {
		MaxNumOfSemiPersistentSRSposResources_r17        int64
		MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 int64
	}
	ChannelBWs_DL_SCS_120kHz_FR2_2_r17                      *per.BitString
	ChannelBWs_UL_SCS_120kHz_FR2_2_r17                      *per.BitString
	Dmrs_BundlingPUSCH_RepTypeA_r17                         *int64
	Dmrs_BundlingPUSCH_RepTypeB_r17                         *int64
	Dmrs_BundlingPUSCH_MultiSlot_r17                        *int64
	Dmrs_BundlingPUCCH_Rep_r17                              *int64
	InterSlotFreqHopInterSlotBundlingPUSCH_r17              *int64
	InterSlotFreqHopPUCCH_r17                               *int64
	Dmrs_BundlingRestart_r17                                *int64
	Dmrs_BundlingNonBackToBackTX_r17                        *int64
	MaxDynamicSlotRepetitionForSPS_Multicast_r17            *int64
	Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17    *int64
	Sps_MulticastDCI_Format4_2_r17                          *int64
	Sps_MulticastMultiConfig_r17                            *int64
	PriorityIndicatorInDCI_Multicast_r17                    *int64
	PriorityIndicatorInDCI_SPS_Multicast_r17                *int64
	TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17          *int64
	MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17             *int64
	ReleaseSPS_MulticastWithCS_RNTI_r17                     *int64
	PosUE_TA_AutoAdjustment_r18                             *int64
	PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18        *int64
	PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18 *int64
	Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18   *DL_PRS_MeasurementWithRxFH_RRC_Connected_r18
	PosSRS_TxFH_RRC_ConnectedForRedCap_r18                  *PosSRS_TxFrequencyHoppingRRC_Connected_r18
	PosSRS_TxFH_RRC_InactiveForRedCap_r18                   *PosSRS_TxFrequencyHoppingRRC_Inactive_r18
	PosSRS_BWA_RRC_Inactive_r18                             *PosSRS_BWA_RRC_Inactive_r18
	PosJointTriggerBySingleDCI_RRC_Connected_r18            *int64
	Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18    *int64
	Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18        *int64
	SpatialAdaptation_CSI_Feedback_r18                      *struct {
		CsiFeedbackType_r18            int64
		MaxNumberLmax_r18              int64
		MaxNumberCSI_ResourcePerCC_r18 struct {
			SdType1_Resource_r18 int64
			SdType2_Resource_r18 int64
		}
		MaxNumberTotalCSI_ResourcePerCC_r18 struct {
			SdType1_Resource_r18 int64
			SdType2_Resource_r18 int64
		}
		TotalNumberCSI_Reporting_r18 int64
	}
	SpatialAdaptation_CSI_FeedbackPUSCH_r18 *struct {
		CsiFeedbackType_r18                 int64
		MaxNumberLmax_r18                   int64
		SubReportCSI_r18                    int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	SpatialAdaptation_CSI_FeedbackAperiodic_r18 *struct {
		CsiFeedbackType_r18            int64
		MaxNumberLmax_r18              int64
		SubReportCSI_r18               int64
		MaxNumberCSI_ResourcePerCC_r18 struct {
			SdType1_Resource_r18 int64
			SdType2_Resource_r18 int64
		}
		MaxNumberTotalCSI_ResourcePerCC_r18 struct {
			SdType1_Resource_r18 int64
			SdType2_Resource_r18 int64
		}
		TotalNumberCSI_Reporting_r18 int64
	}
	SpatialAdaptation_CSI_FeedbackPUCCH_r18 *struct {
		CsiFeedbackType_r18                 int64
		MaxNumberLmax_r18                   int64
		SubReportCSI_r18                    int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	PowerAdaptation_CSI_Feedback_r18 *struct {
		MaxNumberLmax_r18                   int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	PowerAdaptation_CSI_FeedbackPUSCH_r18 *struct {
		MaxNumberLmax_r18                   int64
		SubReportCSI_r18                    int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	PowerAdaptation_CSI_FeedbackAperiodic_r18 *struct {
		MaxNumberLmax_r18                   int64
		SubReportCSI_r18                    int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	PowerAdaptation_CSI_FeedbackPUCCH_r18 *struct {
		MaxNumberLmax_r18                   int64
		SubReportCSI_r18                    int64
		MaxNumberCSI_ResourcePerCC_r18      int64
		MaxNumberTotalCSI_ResourcePerCC_r18 int64
		TotalNumberCSI_Reporting_r18        int64
	}
	Nes_CellDTX_DRX_r18                 *int64
	Nes_CellDTX_DRX_DCI2_9_r18          *int64
	MixCodeBookSpatialAdaptation_r18    *int64
	SimultaneousCSI_SubReportsPerCC_r18 *int64
	Ntn_DMRS_BundlingNGSO_r18           *int64
	Ltm_BeamIndicationJointTCI_r18      *struct {
		MaxNumberJointTCI_PerCell_r18     int64
		Qcl_Resource_r18                  int64
		MaxNumberJointTCI_AcrossCells_r18 int64
		MaxNumberCells_r18                int64
	}
	Dummy_Ltm_MAC_CE_JointTCI_r18 *struct {
		Qcl_Resource_r18                  int64
		MaxNumberJointTCI_PerCell_r18     int64
		MaxNumberJointTCI_AcrossCells_r18 int64
	}
	Ltm_BeamIndicationSeparateTCI_r18 *struct {
		MaxNumberDL_TCI_PerCell_r18     int64
		MaxNumberUL_TCI_PerCell_r18     int64
		Qcl_Resource_r18                int64
		MaxNumberDL_TCI_AcrossCells_r18 int64
		MaxNumberUL_TCI_AcrossCells_r18 int64
		MaxNumberCells_r18              int64
	}
	Dummy_Ltm_MAC_CE_SeparateTCI_r18 *struct {
		Qcl_Resource_r18                int64
		MaxNumberDL_TCI_PerCell_r18     int64
		MaxNumberUL_TCI_PerCell_r18     int64
		MaxNumberDL_TCI_AcrossCells_r18 int64
		MaxNumberUL_TCI_AcrossCells_r18 int64
	}
	Rach_EarlyTA_Measurement_r18          *int64
	Ue_TA_Measurement_r18                 *int64
	Ta_IndicationCellSwitch_r18           *int64
	TriggeredHARQ_CodebookRetxDCI_1_3_r18 *struct {
		MinHARQ_Retx_Offset_r18 int64
		MaxHARQ_Retx_Offset_r18 int64
	}
	UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18 *struct {
		MinBeamApplicationTime_r18 struct {
			Choice  int
			Fr1_r18 *struct {
				Scs_15kHz_r18 *int64
				Scs_30kHz_r18 *int64
				Scs_60kHz_r18 *int64
			}
			Fr2_r18 *struct {
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
			}
		}
		MaxActivatedTCI_PerCC_r18 *int64
	}
	UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18 *struct {
		MinBeamApplicationTime_r18 struct {
			Choice  int
			Fr1_r18 *struct {
				Scs_15kHz_r18 *int64
				Scs_30kHz_r18 *int64
				Scs_60kHz_r18 *int64
			}
			Fr2_r18 *struct {
				Scs_60kHz_r18  *int64
				Scs_120kHz_r18 *int64
			}
		}
		MaxActivatedDL_TCI_PerCC_r18 *int64
		MaxActivatedUL_TCI_PerCC_r18 *int64
	}
	MultiPUSCH_CG_r18                    *int64
	MultiPUSCH_ActiveConfiguredGrant_r18 *struct {
		MaxNumberConfigsPerBWP    int64
		MaxNumberConfigsAllCC_FR1 int64
		MaxNumberConfigsAllCC_FR2 int64
	}
	JointReleaseDCI_r18                        *int64
	Cg_PUSCH_UTO_UCI_Ind_r18                   *int64
	Pdcch_MonitoringResumptionAfterUL_NACK_r18 *int64
	Support3MHz_ChannelBW_Symmetric_r18        *int64
	Support3MHz_ChannelBW_Asymmetric_r18       *int64
	Support12PRB_CORESET0_r18                  *int64
	Nr_PDCCH_OverlapLTE_CRS_RE_r18             *struct {
		OverlapInRE_r18     int64
		OverlapInSymbol_r18 int64
	}
	Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18 *int64
	Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18      *int64
	TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18    *struct {
		MaxNumberPatterns_r18            int64
		MaxNumberNon_OverlapPatterns_r18 int64
	}
	OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18 *int64
	Ncd_SSB_BWP_Wor_r18                                        *int64
	Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18                     *int64
	Prach_CoverageEnh_r18                                      *int64
	Prach_Repetition_r18                                       *int64
	DynamicWaveformSwitch_r18                                  *int64
	DynamicWaveformSwitchPHR_r18                               *int64
	DynamicWaveformSwitchIntraCA_r18                           *int64
	MultiPUSCH_SingleDCI_NonConsSlots_r18                      *int64
	Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18              *struct {
		Fr1_r18 struct {
			Scs_15kHz_r18 *int64
			Scs_30kHz_r18 *int64
			Scs_60kHz_r18 *int64
		}
		Fr2_r18 struct {
			Scs_60kHz_r18  *int64
			Scs_120kHz_r18 *int64
		}
	}
	IntraSlot_PDSCH_MulticastInactive_r18 *bool
	MulticastInactive_r18                 *int64
	ThresholdBasedMulticastResume_r18     *int64
	LowerMSD_r18                          []LowerMSD_r18
	LowerMSD_ENDC_r18                     []LowerMSD_r18
	EnhancedChannelRaster_r18             *int64
	FastBeamSweepingMultiRx_r18           *int64
	BeamSweepingFactorReduction_r18       *struct {
		ReduceForCellDetection    int64
		ReduceForSSB_L1_RSRP_Meas int64
	}
	SimultaneousReceptionTwoQCL_r18  *int64
	MeasEnhCAInterFreqFR2_r18        *int64
	Tci_StateSwitchInd_r18           *int64
	AntennaArrayType_r18             *int64
	LocationBasedCondHandoverATG_r18 *int64
	MaxOutputPowerATG_r18            *int64
	Ltm_FastProcessingConfig_r18     *struct {
		MaxNumberStoredConfigCells_r18 int64
		MaxNumberConfigs_r18           int64
	}
	MeasValidationReportEMR_r18                             *int64
	MeasValidationReportReselectionMeasurements_r18         *int64
	EventA4BasedCondHandoverNES_r18                         *int64
	NesBasedCondHandoverWithDCI_r18                         *int64
	Rach_LessHandoverCG_r18                                 *int64
	Rach_LessHandoverDG_r18                                 *int64
	LocationBasedCondHandoverEMC_r18                        *int64
	Mt_CG_SDT_r18                                           *int64
	PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18        *int64
	PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18 *int64
	Cg_SDT_PeriodicityExt_r18                               *int64
	SupportOf2RxXR_r18                                      *int64
	CondHandoverWithCandSCG_Change_r18                      *int64
	Mac_ParametersPerBand_r18                               *MAC_ParametersPerBand_r18
	ChannelBW_DL_NCR_r18                                    *struct {
		Choice     int
		Fr1_100mhz *struct {
			Scs_15kHz *int64
			Scs_30kHz *int64
			Scs_60kHz *int64
		}
		Fr2_200mhz *struct {
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
	}
	ChannelBW_UL_NCR_r18 *struct {
		Choice     int
		Fr1_100mhz *struct {
			Scs_15kHz *int64
			Scs_30kHz *int64
			Scs_60kHz *int64
		}
		Fr2_200mhz *struct {
			Scs_60kHz  *int64
			Scs_120kHz *int64
		}
	}
	Ncr_PDSCH_64QAM_FR2_r18 *int64
	Ltm_MCG_IntraFreq_r18   *int64
	Ltm_SCG_IntraFreq_r18   *int64
	Ltm_MAC_CE_JointTCI_r18 *struct {
		Qcl_Resource_r18                  int64
		MaxNumberJointTCI_PerCell_r18     int64
		MaxNumberJointTCI_AcrossCells_r18 int64
	}
	Ltm_MAC_CE_SeparateTCI_r18 *struct {
		Qcl_Resource_r18                int64
		MaxNumberDL_TCI_PerCell_r18     int64
		MaxNumberUL_TCI_PerCell_r18     int64
		MaxNumberDL_TCI_AcrossCells_r18 int64
		MaxNumberUL_TCI_AcrossCells_r18 int64
	}
	EventA4BasedCondHandoverATG_r18         *int64
	PosSRS_TxFH_WithTimeWindowForRedCap_r18 *int64
	Sbfd_Aware_r19                          *struct {
		NumOfPartialPRG_r19           int64
		Timeline_r19                  int64
		Sbfd_Rx_r19                   *int64
		Sbfd_Tx_r19                   *int64
		Sbfd_OneRACH_r19              *int64
		Sbfd_TwoRACH_r19              *int64
		PreambleRepetitionOneRACH_r19 *int64
		PreambleRepetitionTwoRACH_r19 *int64
	}
	Ul_ResourceMutingCP_OFDM_r19              *int64
	Ul_ResourceMutingDFTS_OFDM_r19            *int64
	Ul_ResourceMutingCP_OFDM_CG_Type1_r19     *int64
	Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19   *int64
	L1_CLI_RSSI_MeasAndAperiodicReporting_r19 *struct {
		MaxNumberOfConfiguredMeasResourceAcrossCC_r19 int64
		MaxNumberOfAperiodicReportSettingPerBWP_r19   int64
		MaxNumberOfConfiguredMeasResourcePerCC_r19    int64
		MaxNumberOfSimultaneousMeasResourcePerCC_r19  int64
		PeriodicReporting_r19                         *int64
		SemiPersistentMeasResource_r19                *int64
	}
	L1_SRS_RSRP_MeasAndAperiodicReporting_r19 *struct {
		MaxNumberOfConfiguredMeasResourceAcrossCC_r19 int64
		MaxNumberOfConfiguredMeasResourcePerCC_r19    int64
		MaxNumberOfSimultaneousMeasResourcePerCC_r19  int64
		MaxNumberOfAperiodicReportSettingPerBWP_r19   int64
		MaxNumberOfMeasResourceAcrossCCWithinSlot_r19 int64
		MaxNumberOfAperiodicMeasResourceAcrossCC_r19  int64
		Fdm_Reception_r19                             *int64
		SemiPersistentMeasResource_r19                *int64
	}
	Od_SSB_NoAlwaysOn_RRC_r19           *int64
	Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19    *int64
	Od_SSB_AlwaysOn_RRC_r19             *int64
	Od_SSB_AlwaysOn_RRC_Diff_r19        *int64
	Od_SSB_AlwaysOn_RRC_MAC_CE_r19      *int64
	Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19 *int64
	Od_SSB_NoAlwaysOn_MAC_CE_r19        *int64
	Od_SSB_AlwaysOn_MAC_CE_r19          *struct {
		TimeRelation_r19       int64
		DeactivationScheme_r19 int64
	}
	Od_SSB_AlwaysOn_MAC_CE_Diff_r19          *int64
	Ssb_BurstPeriodicityAdaptation_r19       *int64
	Rach_AdaptationTimeDomain_r19            *int64
	Ltm_BeamIndicationJointTCI_CSI_RS_r19    *int64
	Ltm_MAC_CE_JointTCI_CSI_RS_r19           *int64
	Ltm_BeamIndicationSeparateTCI_CSI_RS_r19 *int64
	Ltm_MAC_CE_SeparateTCI_CSI_RS_r19        *int64
	Ltm_CSI_RS_CSI_IM_Periodic_r19           *struct {
		MaxNumOfCSI_RS_Resource_r19  int64
		MaxNumOfCSI_RS_Ports_r19     int64
		MaxNumOfNZP_CSI_RS_Ports_r19 int64
		MaxRank_r19                  int64
		MaxNumOfCSI_IM_Resource_r19  int64
	}
	Ltm_CSI_RS_CSI_IM_SP_r19 *struct {
		MaxNumOfCSI_RS_Resource_r19  int64
		MaxNumOfCSI_RS_Ports_r19     int64
		MaxNumOfNZP_CSI_RS_Ports_r19 int64
		MaxRank_r19                  int64
		MaxNumOfCSI_IM_Resource_r19  int64
	}
	Ltm_WithoutCSI_IM_r19          *int64
	Ssb_PeriodicityAddition_r19    *int64
	Pdcch_RepetitionType0_r19      *int64
	Pdcch_RepetitionTypeOthers_r19 *int64
	Pdsch_RepetitionMsg4_r19       *int64
	Ntn_Collision_RedCap_r19       *int64
	Pusch_InterSlotOCC_r19         *struct {
		OccLength_r19      int64
		SatelliteOrbit_r19 int64
	}
	TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19 *struct {
		MinHARQ_Retx_Offset_r19 int64
		MaxHARQ_Retx_Offset_r19 int64
	}
	UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19 *struct {
		MinBeamApplicationTime_r19 struct {
			Choice  int
			Fr1_r19 *struct {
				Scs_15kHz_r19 *int64
				Scs_30kHz_r19 *int64
				Scs_60kHz_r19 *int64
			}
			Fr2_r19 *struct {
				Scs_60kHz_r19  *int64
				Scs_120kHz_r19 *int64
			}
		}
		MaxActivatedTCI_PerCC_r19 *int64
	}
	UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19 *struct {
		MinBeamApplicationTime_r19 struct {
			Choice  int
			Fr1_r19 *struct {
				Scs_15kHz_r19 *int64
				Scs_30kHz_r19 *int64
				Scs_60kHz_r19 *int64
			}
			Fr2_r19 *struct {
				Scs_60kHz_r19  *int64
				Scs_120kHz_r19 *int64
			}
		}
		MaxActivatedDL_TCI_PerCC_r19 *int64
		MaxActivatedUL_TCI_PerCC_r19 *int64
	}
	PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19  *PosSRS_TxFrequencyHoppingRRC_ConnectedNonRedCap_r19
	PosSRS_TxFH_RRC_InactiveForNonRedCap_r19   *PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19
	PosSRS_TxFH_WithTimeWindowForNonRedCap_r19 *int64
	Sr_TriggeredSSSG_Switching_r19             *int64
	Pdcch_RepetitionType0_TN_r19               *int64
	Pdcch_RepetitionTypeOthersTN_r19           *int64
	Mpr_SingleCC_r19                           *struct {
		Choice                         int
		Mpr_SingleCC_SingleValue_r19   *int64
		Mpr_SingleCC_MultipleValue_r19 *int64
	}
	FastRx_BSF_MeasDelayReduction_r19   *int64
	FastSCellActivationEarlyMeas_r19    *int64
	Od_SSB_AdditionalProcessingTime_r19 *int64
	Aiml_BM_Case1_KnownRxBeam_r19       *int64
	Aiml_BM_Case2_KnownRxBeam_r19       *int64
	Ntn_PowerBoosting_ERedCap_r19       *int64
	EarlyCSI_AcquisitionWithCSI_IM_r19  *struct {
		MaxNumOfCSI_RS_Resource_r19  int64
		MaxNumOfCSI_RS_Ports_r19     int64
		MaxNumOfNZP_CSI_RS_Ports_r19 int64
		MaxRank_r19                  int64
		MaxNumOfCSI_IM_Resource_r19  int64
	}
	EarlyCSI_AcquisitionWithoutCSI_IM_r19 *int64
	EnableTx_RxDuringMeasGap_r19          *struct {
		AdditionalDCI_r19     per.BitString
		IndicationField_r19   int64
		MinimumTimeOffset_r19 int64
	}
}

func (ie *BandNR) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandNRConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxUplinkDutyCycle_PC2_FR1 != nil
	hasExtGroup1 := ie.Pucch_SpatialRelInfoMAC_CE != nil || ie.PowerBoosting_Pi2BPSK != nil
	hasExtGroup2 := ie.MaxUplinkDutyCycle_FR2 != nil
	hasExtGroup3 := ie.ChannelBWs_DL_v1590 != nil || ie.ChannelBWs_UL_v1590 != nil
	hasExtGroup4 := ie.AsymmetricBandwidthCombinationSet != nil
	hasExtGroup5 := ie.SharedSpectrumChAccessParamsPerBand_r16 != nil || ie.CancelOverlappingPUSCH_r16 != nil || ie.MultipleRateMatchingEUTRA_CRS_r16 != nil || ie.OverlapRateMatchingEUTRA_CRS_r16 != nil || ie.Pdsch_MappingTypeB_Alt_r16 != nil || ie.OneSlotPeriodicTRS_r16 != nil || ie.Olpc_SRS_Pos_r16 != nil || ie.SpatialRelationsSRS_Pos_r16 != nil || ie.SimulSRS_MIMO_TransWithinBand_r16 != nil || ie.ChannelBW_DL_IAB_r16 != nil || ie.ChannelBW_UL_IAB_r16 != nil || ie.RasterShift7dot5_IAB_r16 != nil || ie.Ue_PowerClass_v1610 != nil || ie.CondHandover_r16 != nil || ie.CondHandoverFailure_r16 != nil || ie.CondHandoverTwoTriggerEvents_r16 != nil || ie.CondPSCellChange_r16 != nil || ie.CondPSCellChangeTwoTriggerEvents_r16 != nil || ie.Mpr_PowerBoost_FR2_r16 != nil || ie.ActiveConfiguredGrant_r16 != nil || ie.JointReleaseConfiguredGrantType2_r16 != nil || ie.Sps_r16 != nil || ie.JointReleaseSPS_r16 != nil || ie.SimulSRS_TransWithinBand_r16 != nil || ie.Trs_AdditionalBandwidth_r16 != nil || ie.HandoverIntraF_IAB_r16 != nil
	hasExtGroup6 := ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil
	hasExtGroup7 := ie.HandoverUTRA_FDD_r16 != nil || ie.EnhancedUL_TransientPeriod_r16 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil
	hasExtGroup8 := ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.Pusch_RepetitionMultiSlots_v1650 != nil || ie.ConfiguredUL_GrantType1_v1650 != nil || ie.ConfiguredUL_GrantType2_v1650 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil
	hasExtGroup9 := ie.EnhancedSkipUplinkTxConfigured_v1660 != nil || ie.EnhancedSkipUplinkTxDynamic_v1660 != nil
	hasExtGroup10 := ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil || ie.TxDiversity_r16 != nil
	hasExtGroup11 := ie.Pdsch_1024QAM_FR1_r17 != nil || ie.Ue_PowerClass_v1700 != nil || ie.Fr2_2_AccessParamsPerBand_r17 != nil || ie.Rlm_Relaxation_r17 != nil || ie.Bfd_Relaxation_r17 != nil || ie.Cg_SDT_r17 != nil || ie.LocationBasedCondHandover_r17 != nil || ie.TimeBasedCondHandover_r17 != nil || ie.EventA4BasedCondHandover_r17 != nil || ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.Pdcch_SkippingWithoutSSSG_r17 != nil || ie.Sssg_Switching_1BitInd_r17 != nil || ie.Sssg_Switching_2BitInd_r17 != nil || ie.Pdcch_SkippingWithSSSG_r17 != nil || ie.SearchSpaceSetGrp_SwitchCap2_r17 != nil || ie.UplinkPreCompensation_r17 != nil || ie.Uplink_TA_Reporting_r17 != nil || ie.Max_HARQ_ProcessNumber_r17 != nil || ie.Type2_HARQ_Codebook_r17 != nil || ie.Type1_HARQ_Codebook_r17 != nil || ie.Type3_HARQ_Codebook_r17 != nil || ie.Ue_Specific_K_Offset_r17 != nil || ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.ParallelPRS_MeasRRC_Inactive_r17 != nil || ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil || ie.Prs_ProcessingRRC_Inactive_r17 != nil || ie.Prs_ProcessingWindowType1A_r17 != nil || ie.Prs_ProcessingWindowType1B_r17 != nil || ie.Prs_ProcessingWindowType2_r17 != nil || ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil || ie.Olpc_SRS_PosRRC_Inactive_r17 != nil || ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil || ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil || ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil || ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil || ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil || ie.MaxDurationDMRS_Bundling_r17 != nil || ie.Pusch_RepetitionMsg3_r17 != nil || ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil || ie.ParallelMeasurementWithoutRestriction_r17 != nil || ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil || ie.K1_RangeExtension_r17 != nil || ie.AperiodicCSI_RS_FastScellActivation_r17 != nil || ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil || ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil || ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil || ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil || ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil || ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil || ie.Ul_GapFR2_r17 != nil || ie.OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17 != nil || ie.OneShotHARQ_FeedbackPhy_Priority_r17 != nil || ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil || ie.TriggeredHARQ_CodebookRetx_r17 != nil
	hasExtGroup12 := ie.Ue_OneShotUL_TimingAdj_r17 != nil || ie.Pucch_Repetition_F0_2_r17 != nil || ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil || ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil || ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil || ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_RNTI_r17 != nil || ie.DynamicMulticastDCI_Format4_2_r17 != nil || ie.MaxModulationOrderForMulticast_r17 != nil || ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil || ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil || ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.MaxNumberG_CS_RNTI_r17 != nil || ie.Re_LevelRateMatchingForMulticast_r17 != nil || ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil || ie.Prs_MeasurementWithoutMG_r17 != nil || ie.MaxNumber_NGSO_SatellitesPerCarrier_r17 != nil || ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 != nil || ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil || ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil || ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil
	hasExtGroup13 := ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil || ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil || ie.Dmrs_BundlingPUSCH_MultiSlot_r17 != nil || ie.Dmrs_BundlingPUCCH_Rep_r17 != nil || ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil || ie.InterSlotFreqHopPUCCH_r17 != nil || ie.Dmrs_BundlingRestart_r17 != nil || ie.Dmrs_BundlingNonBackToBackTX_r17 != nil
	hasExtGroup14 := ie.MaxDynamicSlotRepetitionForSPS_Multicast_r17 != nil || ie.Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.Sps_MulticastDCI_Format4_2_r17 != nil || ie.Sps_MulticastMultiConfig_r17 != nil || ie.PriorityIndicatorInDCI_Multicast_r17 != nil || ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil || ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil || ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil || ie.ReleaseSPS_MulticastWithCS_RNTI_r17 != nil
	hasExtGroup15 := ie.PosUE_TA_AutoAdjustment_r18 != nil || ie.PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18 != nil || ie.PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18 != nil || ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18 != nil || ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18 != nil || ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18 != nil || ie.PosSRS_BWA_RRC_Inactive_r18 != nil || ie.PosJointTriggerBySingleDCI_RRC_Connected_r18 != nil || ie.Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18 != nil || ie.Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18 != nil || ie.SpatialAdaptation_CSI_Feedback_r18 != nil || ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18 != nil || ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18 != nil || ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18 != nil || ie.PowerAdaptation_CSI_Feedback_r18 != nil || ie.PowerAdaptation_CSI_FeedbackPUSCH_r18 != nil || ie.PowerAdaptation_CSI_FeedbackAperiodic_r18 != nil || ie.PowerAdaptation_CSI_FeedbackPUCCH_r18 != nil || ie.Nes_CellDTX_DRX_r18 != nil || ie.Nes_CellDTX_DRX_DCI2_9_r18 != nil || ie.MixCodeBookSpatialAdaptation_r18 != nil || ie.SimultaneousCSI_SubReportsPerCC_r18 != nil || ie.Ntn_DMRS_BundlingNGSO_r18 != nil || ie.Ltm_BeamIndicationJointTCI_r18 != nil || ie.Dummy_Ltm_MAC_CE_JointTCI_r18 != nil || ie.Ltm_BeamIndicationSeparateTCI_r18 != nil || ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18 != nil || ie.Rach_EarlyTA_Measurement_r18 != nil || ie.Ue_TA_Measurement_r18 != nil || ie.Ta_IndicationCellSwitch_r18 != nil || ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18 != nil || ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18 != nil || ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18 != nil || ie.MultiPUSCH_CG_r18 != nil || ie.MultiPUSCH_ActiveConfiguredGrant_r18 != nil || ie.JointReleaseDCI_r18 != nil || ie.Cg_PUSCH_UTO_UCI_Ind_r18 != nil || ie.Pdcch_MonitoringResumptionAfterUL_NACK_r18 != nil || ie.Support3MHz_ChannelBW_Symmetric_r18 != nil || ie.Support3MHz_ChannelBW_Asymmetric_r18 != nil || ie.Support12PRB_CORESET0_r18 != nil || ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18 != nil || ie.Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18 != nil || ie.Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18 != nil || ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18 != nil || ie.OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18 != nil || ie.Ncd_SSB_BWP_Wor_r18 != nil || ie.Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18 != nil || ie.Prach_CoverageEnh_r18 != nil || ie.Prach_Repetition_r18 != nil || ie.DynamicWaveformSwitch_r18 != nil || ie.DynamicWaveformSwitchPHR_r18 != nil || ie.DynamicWaveformSwitchIntraCA_r18 != nil || ie.MultiPUSCH_SingleDCI_NonConsSlots_r18 != nil || ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18 != nil || ie.IntraSlot_PDSCH_MulticastInactive_r18 != nil || ie.MulticastInactive_r18 != nil || ie.ThresholdBasedMulticastResume_r18 != nil || ie.LowerMSD_r18 != nil || ie.LowerMSD_ENDC_r18 != nil || ie.EnhancedChannelRaster_r18 != nil || ie.FastBeamSweepingMultiRx_r18 != nil || ie.BeamSweepingFactorReduction_r18 != nil || ie.SimultaneousReceptionTwoQCL_r18 != nil || ie.MeasEnhCAInterFreqFR2_r18 != nil || ie.Tci_StateSwitchInd_r18 != nil || ie.AntennaArrayType_r18 != nil || ie.LocationBasedCondHandoverATG_r18 != nil || ie.MaxOutputPowerATG_r18 != nil || ie.Ltm_FastProcessingConfig_r18 != nil || ie.MeasValidationReportEMR_r18 != nil || ie.MeasValidationReportReselectionMeasurements_r18 != nil || ie.EventA4BasedCondHandoverNES_r18 != nil || ie.NesBasedCondHandoverWithDCI_r18 != nil || ie.Rach_LessHandoverCG_r18 != nil || ie.Rach_LessHandoverDG_r18 != nil || ie.LocationBasedCondHandoverEMC_r18 != nil || ie.Mt_CG_SDT_r18 != nil || ie.PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18 != nil || ie.PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18 != nil || ie.Cg_SDT_PeriodicityExt_r18 != nil || ie.SupportOf2RxXR_r18 != nil || ie.CondHandoverWithCandSCG_Change_r18 != nil
	hasExtGroup16 := ie.Mac_ParametersPerBand_r18 != nil || ie.ChannelBW_DL_NCR_r18 != nil || ie.ChannelBW_UL_NCR_r18 != nil || ie.Ncr_PDSCH_64QAM_FR2_r18 != nil || ie.Ltm_MCG_IntraFreq_r18 != nil || ie.Ltm_SCG_IntraFreq_r18 != nil
	hasExtGroup17 := ie.Ltm_MAC_CE_JointTCI_r18 != nil || ie.Ltm_MAC_CE_SeparateTCI_r18 != nil
	hasExtGroup18 := ie.EventA4BasedCondHandoverATG_r18 != nil
	hasExtGroup19 := ie.PosSRS_TxFH_WithTimeWindowForRedCap_r18 != nil
	hasExtGroup20 := ie.Sbfd_Aware_r19 != nil || ie.Ul_ResourceMutingCP_OFDM_r19 != nil || ie.Ul_ResourceMutingDFTS_OFDM_r19 != nil || ie.Ul_ResourceMutingCP_OFDM_CG_Type1_r19 != nil || ie.Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19 != nil || ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19 != nil || ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19 != nil || ie.Od_SSB_NoAlwaysOn_RRC_r19 != nil || ie.Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19 != nil || ie.Od_SSB_AlwaysOn_RRC_r19 != nil || ie.Od_SSB_AlwaysOn_RRC_Diff_r19 != nil || ie.Od_SSB_AlwaysOn_RRC_MAC_CE_r19 != nil || ie.Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19 != nil || ie.Od_SSB_NoAlwaysOn_MAC_CE_r19 != nil || ie.Od_SSB_AlwaysOn_MAC_CE_r19 != nil || ie.Od_SSB_AlwaysOn_MAC_CE_Diff_r19 != nil || ie.Ssb_BurstPeriodicityAdaptation_r19 != nil || ie.Rach_AdaptationTimeDomain_r19 != nil || ie.Ltm_BeamIndicationJointTCI_CSI_RS_r19 != nil || ie.Ltm_MAC_CE_JointTCI_CSI_RS_r19 != nil || ie.Ltm_BeamIndicationSeparateTCI_CSI_RS_r19 != nil || ie.Ltm_MAC_CE_SeparateTCI_CSI_RS_r19 != nil || ie.Ltm_CSI_RS_CSI_IM_Periodic_r19 != nil || ie.Ltm_CSI_RS_CSI_IM_SP_r19 != nil || ie.Ltm_WithoutCSI_IM_r19 != nil || ie.Ssb_PeriodicityAddition_r19 != nil || ie.Pdcch_RepetitionType0_r19 != nil || ie.Pdcch_RepetitionTypeOthers_r19 != nil || ie.Pdsch_RepetitionMsg4_r19 != nil || ie.Ntn_Collision_RedCap_r19 != nil || ie.Pusch_InterSlotOCC_r19 != nil || ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19 != nil || ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19 != nil || ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19 != nil || ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19 != nil || ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19 != nil || ie.PosSRS_TxFH_WithTimeWindowForNonRedCap_r19 != nil || ie.Sr_TriggeredSSSG_Switching_r19 != nil || ie.Pdcch_RepetitionType0_TN_r19 != nil || ie.Pdcch_RepetitionTypeOthersTN_r19 != nil || ie.Mpr_SingleCC_r19 != nil || ie.FastRx_BSF_MeasDelayReduction_r19 != nil || ie.FastSCellActivationEarlyMeas_r19 != nil || ie.Od_SSB_AdditionalProcessingTime_r19 != nil || ie.Aiml_BM_Case1_KnownRxBeam_r19 != nil || ie.Aiml_BM_Case2_KnownRxBeam_r19 != nil || ie.Ntn_PowerBoosting_ERedCap_r19 != nil || ie.EarlyCSI_AcquisitionWithCSI_IM_r19 != nil || ie.EarlyCSI_AcquisitionWithoutCSI_IM_r19 != nil
	hasExtGroup21 := ie.EnableTx_RxDuringMeasGap_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10 || hasExtGroup11 || hasExtGroup12 || hasExtGroup13 || hasExtGroup14 || hasExtGroup15 || hasExtGroup16 || hasExtGroup17 || hasExtGroup18 || hasExtGroup19 || hasExtGroup20 || hasExtGroup21

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ModifiedMPR_Behaviour != nil, ie.Mimo_ParametersPerBand != nil, ie.ExtendedCP != nil, ie.MultipleTCI != nil, ie.Bwp_WithoutRestriction != nil, ie.Bwp_SameNumerology != nil, ie.Bwp_DiffNumerology != nil, ie.CrossCarrierScheduling_SameSCS != nil, ie.Pdsch_256QAM_FR2 != nil, ie.Pusch_256QAM != nil, ie.Ue_PowerClass != nil, ie.RateMatchingLTE_CRS != nil, ie.ChannelBWs_DL != nil, ie.ChannelBWs_UL != nil}); err != nil {
		return err
	}

	// 3. bandNR: ref
	{
		if err := ie.BandNR.Encode(e); err != nil {
			return err
		}
	}

	// 4. modifiedMPR-Behaviour: bit-string
	{
		if ie.ModifiedMPR_Behaviour != nil {
			if err := e.EncodeBitString(*ie.ModifiedMPR_Behaviour, bandNRModifiedMPRBehaviourConstraints); err != nil {
				return err
			}
		}
	}

	// 5. mimo-ParametersPerBand: ref
	{
		if ie.Mimo_ParametersPerBand != nil {
			if err := ie.Mimo_ParametersPerBand.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. extendedCP: enumerated
	{
		if ie.ExtendedCP != nil {
			if err := e.EncodeEnumerated(*ie.ExtendedCP, bandNRExtendedCPConstraints); err != nil {
				return err
			}
		}
	}

	// 7. multipleTCI: enumerated
	{
		if ie.MultipleTCI != nil {
			if err := e.EncodeEnumerated(*ie.MultipleTCI, bandNRMultipleTCIConstraints); err != nil {
				return err
			}
		}
	}

	// 8. bwp-WithoutRestriction: enumerated
	{
		if ie.Bwp_WithoutRestriction != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_WithoutRestriction, bandNRBwpWithoutRestrictionConstraints); err != nil {
				return err
			}
		}
	}

	// 9. bwp-SameNumerology: enumerated
	{
		if ie.Bwp_SameNumerology != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_SameNumerology, bandNRBwpSameNumerologyConstraints); err != nil {
				return err
			}
		}
	}

	// 10. bwp-DiffNumerology: enumerated
	{
		if ie.Bwp_DiffNumerology != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_DiffNumerology, bandNRBwpDiffNumerologyConstraints); err != nil {
				return err
			}
		}
	}

	// 11. crossCarrierScheduling-SameSCS: enumerated
	{
		if ie.CrossCarrierScheduling_SameSCS != nil {
			if err := e.EncodeEnumerated(*ie.CrossCarrierScheduling_SameSCS, bandNRCrossCarrierSchedulingSameSCSConstraints); err != nil {
				return err
			}
		}
	}

	// 12. pdsch-256QAM-FR2: enumerated
	{
		if ie.Pdsch_256QAM_FR2 != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_256QAM_FR2, bandNRPdsch256QAMFR2Constraints); err != nil {
				return err
			}
		}
	}

	// 13. pusch-256QAM: enumerated
	{
		if ie.Pusch_256QAM != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_256QAM, bandNRPusch256QAMConstraints); err != nil {
				return err
			}
		}
	}

	// 14. ue-PowerClass: enumerated
	{
		if ie.Ue_PowerClass != nil {
			if err := e.EncodeEnumerated(*ie.Ue_PowerClass, bandNRUePowerClassConstraints); err != nil {
				return err
			}
		}
	}

	// 15. rateMatchingLTE-CRS: enumerated
	{
		if ie.RateMatchingLTE_CRS != nil {
			if err := e.EncodeEnumerated(*ie.RateMatchingLTE_CRS, bandNRRateMatchingLTECRSConstraints); err != nil {
				return err
			}
		}
	}

	// 16. channelBWs-DL: choice
	{
		if ie.ChannelBWs_DL != nil {
			choiceEnc := e.NewChoiceEncoder(bandNRChannelBWsDLConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBWs_DL).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ChannelBWs_DL).Choice {
			case BandNR_ChannelBWs_DL_Fr1:
				c := (*(*ie.ChannelBWs_DL).Fr1)
				bandNRChannelBWsDLFr1Seq := e.NewSequenceEncoder(bandNRChannelBWsDLFr1Constraints)
				if err := bandNRChannelBWsDLFr1Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz != nil {
					if err := e.EncodeBitString((*c.Scs_15kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
				if c.Scs_30kHz != nil {
					if err := e.EncodeBitString((*c.Scs_30kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeBitString((*c.Scs_60kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
			case BandNR_ChannelBWs_DL_Fr2:
				c := (*(*ie.ChannelBWs_DL).Fr2)
				bandNRChannelBWsDLFr2Seq := e.NewSequenceEncoder(bandNRChannelBWsDLFr2Constraints)
				if err := bandNRChannelBWsDLFr2Seq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeBitString((*c.Scs_60kHz), per.FixedSize(3)); err != nil {
						return err
					}
				}
				if c.Scs_120kHz != nil {
					if err := e.EncodeBitString((*c.Scs_120kHz), per.FixedSize(3)); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ChannelBWs_DL).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 17. channelBWs-UL: choice
	{
		if ie.ChannelBWs_UL != nil {
			choiceEnc := e.NewChoiceEncoder(bandNRChannelBWsULConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBWs_UL).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.ChannelBWs_UL).Choice {
			case BandNR_ChannelBWs_UL_Fr1:
				c := (*(*ie.ChannelBWs_UL).Fr1)
				bandNRChannelBWsULFr1Seq := e.NewSequenceEncoder(bandNRChannelBWsULFr1Constraints)
				if err := bandNRChannelBWsULFr1Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz != nil {
					if err := e.EncodeBitString((*c.Scs_15kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
				if c.Scs_30kHz != nil {
					if err := e.EncodeBitString((*c.Scs_30kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeBitString((*c.Scs_60kHz), per.FixedSize(10)); err != nil {
						return err
					}
				}
			case BandNR_ChannelBWs_UL_Fr2:
				c := (*(*ie.ChannelBWs_UL).Fr2)
				bandNRChannelBWsULFr2Seq := e.NewSequenceEncoder(bandNRChannelBWsULFr2Constraints)
				if err := bandNRChannelBWsULFr2Seq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz != nil {
					if err := e.EncodeBitString((*c.Scs_60kHz), per.FixedSize(3)); err != nil {
						return err
					}
				}
				if c.Scs_120kHz != nil {
					if err := e.EncodeBitString((*c.Scs_120kHz), per.FixedSize(3)); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.ChannelBWs_UL).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10, hasExtGroup11, hasExtGroup12, hasExtGroup13, hasExtGroup14, hasExtGroup15, hasExtGroup16, hasExtGroup17, hasExtGroup18, hasExtGroup19, hasExtGroup20, hasExtGroup21}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "maxUplinkDutyCycle-PC2-FR1", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxUplinkDutyCycle_PC2_FR1 != nil}); err != nil {
				return err
			}

			if ie.MaxUplinkDutyCycle_PC2_FR1 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxUplinkDutyCycle_PC2_FR1, bandNRExtMaxUplinkDutyCyclePC2FR1Constraints); err != nil {
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
					{Name: "pucch-SpatialRelInfoMAC-CE", Optional: true},
					{Name: "powerBoosting-pi2BPSK", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pucch_SpatialRelInfoMAC_CE != nil, ie.PowerBoosting_Pi2BPSK != nil}); err != nil {
				return err
			}

			if ie.Pucch_SpatialRelInfoMAC_CE != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_SpatialRelInfoMAC_CE, bandNRExtPucchSpatialRelInfoMACCEConstraints); err != nil {
					return err
				}
			}

			if ie.PowerBoosting_Pi2BPSK != nil {
				if err := ex.EncodeEnumerated(*ie.PowerBoosting_Pi2BPSK, bandNRExtPowerBoostingPi2BPSKConstraints); err != nil {
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
					{Name: "maxUplinkDutyCycle-FR2", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxUplinkDutyCycle_FR2 != nil}); err != nil {
				return err
			}

			if ie.MaxUplinkDutyCycle_FR2 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxUplinkDutyCycle_FR2, bandNRExtMaxUplinkDutyCycleFR2Constraints); err != nil {
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
					{Name: "channelBWs-DL-v1590", Optional: true},
					{Name: "channelBWs-UL-v1590", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelBWs_DL_v1590 != nil, ie.ChannelBWs_UL_v1590 != nil}); err != nil {
				return err
			}

			if ie.ChannelBWs_DL_v1590 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWsDLV1590Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBWs_DL_v1590).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBWs_DL_v1590).Choice {
				case BandNR_Ext_ChannelBWs_DL_v1590_Fr1:
					c := (*(*ie.ChannelBWs_DL_v1590).Fr1)
					bandNRExtChannelBWsDLV1590Fr1Seq := ex.NewSequenceEncoder(bandNRExtChannelBWsDLV1590Fr1Constraints)
					if err := bandNRExtChannelBWsDLV1590Fr1Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_15kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_30kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_60kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBWs_DL_v1590_Fr2:
					c := (*(*ie.ChannelBWs_DL_v1590).Fr2)
					bandNRExtChannelBWsDLV1590Fr2Seq := ex.NewSequenceEncoder(bandNRExtChannelBWsDLV1590Fr2Constraints)
					if err := bandNRExtChannelBWsDLV1590Fr2Seq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_60kHz), per.FixedSize(8)); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_120kHz), per.FixedSize(8)); err != nil {
							return err
						}
					}
				}
			}

			if ie.ChannelBWs_UL_v1590 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWsULV1590Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBWs_UL_v1590).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBWs_UL_v1590).Choice {
				case BandNR_Ext_ChannelBWs_UL_v1590_Fr1:
					c := (*(*ie.ChannelBWs_UL_v1590).Fr1)
					bandNRExtChannelBWsULV1590Fr1Seq := ex.NewSequenceEncoder(bandNRExtChannelBWsULV1590Fr1Constraints)
					if err := bandNRExtChannelBWsULV1590Fr1Seq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_15kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_30kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_60kHz), per.FixedSize(16)); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBWs_UL_v1590_Fr2:
					c := (*(*ie.ChannelBWs_UL_v1590).Fr2)
					bandNRExtChannelBWsULV1590Fr2Seq := ex.NewSequenceEncoder(bandNRExtChannelBWsULV1590Fr2Constraints)
					if err := bandNRExtChannelBWsULV1590Fr2Seq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_60kHz), per.FixedSize(8)); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeBitString((*c.Scs_120kHz), per.FixedSize(8)); err != nil {
							return err
						}
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
					{Name: "asymmetricBandwidthCombinationSet", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AsymmetricBandwidthCombinationSet != nil}); err != nil {
				return err
			}

			if ie.AsymmetricBandwidthCombinationSet != nil {
				if err := ex.EncodeBitString(*ie.AsymmetricBandwidthCombinationSet, bandNRAsymmetricBandwidthCombinationSetConstraints); err != nil {
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
					{Name: "sharedSpectrumChAccessParamsPerBand-r16", Optional: true},
					{Name: "cancelOverlappingPUSCH-r16", Optional: true},
					{Name: "multipleRateMatchingEUTRA-CRS-r16", Optional: true},
					{Name: "overlapRateMatchingEUTRA-CRS-r16", Optional: true},
					{Name: "pdsch-MappingTypeB-Alt-r16", Optional: true},
					{Name: "oneSlotPeriodicTRS-r16", Optional: true},
					{Name: "olpc-SRS-Pos-r16", Optional: true},
					{Name: "spatialRelationsSRS-Pos-r16", Optional: true},
					{Name: "simulSRS-MIMO-TransWithinBand-r16", Optional: true},
					{Name: "channelBW-DL-IAB-r16", Optional: true},
					{Name: "channelBW-UL-IAB-r16", Optional: true},
					{Name: "rasterShift7dot5-IAB-r16", Optional: true},
					{Name: "ue-PowerClass-v1610", Optional: true},
					{Name: "condHandover-r16", Optional: true},
					{Name: "condHandoverFailure-r16", Optional: true},
					{Name: "condHandoverTwoTriggerEvents-r16", Optional: true},
					{Name: "condPSCellChange-r16", Optional: true},
					{Name: "condPSCellChangeTwoTriggerEvents-r16", Optional: true},
					{Name: "mpr-PowerBoost-FR2-r16", Optional: true},
					{Name: "activeConfiguredGrant-r16", Optional: true},
					{Name: "jointReleaseConfiguredGrantType2-r16", Optional: true},
					{Name: "sps-r16", Optional: true},
					{Name: "jointReleaseSPS-r16", Optional: true},
					{Name: "simulSRS-TransWithinBand-r16", Optional: true},
					{Name: "trs-AdditionalBandwidth-r16", Optional: true},
					{Name: "handoverIntraF-IAB-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SharedSpectrumChAccessParamsPerBand_r16 != nil, ie.CancelOverlappingPUSCH_r16 != nil, ie.MultipleRateMatchingEUTRA_CRS_r16 != nil, ie.OverlapRateMatchingEUTRA_CRS_r16 != nil, ie.Pdsch_MappingTypeB_Alt_r16 != nil, ie.OneSlotPeriodicTRS_r16 != nil, ie.Olpc_SRS_Pos_r16 != nil, ie.SpatialRelationsSRS_Pos_r16 != nil, ie.SimulSRS_MIMO_TransWithinBand_r16 != nil, ie.ChannelBW_DL_IAB_r16 != nil, ie.ChannelBW_UL_IAB_r16 != nil, ie.RasterShift7dot5_IAB_r16 != nil, ie.Ue_PowerClass_v1610 != nil, ie.CondHandover_r16 != nil, ie.CondHandoverFailure_r16 != nil, ie.CondHandoverTwoTriggerEvents_r16 != nil, ie.CondPSCellChange_r16 != nil, ie.CondPSCellChangeTwoTriggerEvents_r16 != nil, ie.Mpr_PowerBoost_FR2_r16 != nil, ie.ActiveConfiguredGrant_r16 != nil, ie.JointReleaseConfiguredGrantType2_r16 != nil, ie.Sps_r16 != nil, ie.JointReleaseSPS_r16 != nil, ie.SimulSRS_TransWithinBand_r16 != nil, ie.Trs_AdditionalBandwidth_r16 != nil, ie.HandoverIntraF_IAB_r16 != nil}); err != nil {
				return err
			}

			if ie.SharedSpectrumChAccessParamsPerBand_r16 != nil {
				if err := ie.SharedSpectrumChAccessParamsPerBand_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CancelOverlappingPUSCH_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CancelOverlappingPUSCH_r16, bandNRExtCancelOverlappingPUSCHR16Constraints); err != nil {
					return err
				}
			}

			if ie.MultipleRateMatchingEUTRA_CRS_r16 != nil {
				c := ie.MultipleRateMatchingEUTRA_CRS_r16
				if err := ex.EncodeInteger(c.MaxNumberPatterns_r16, per.Constrained(2, 6)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberNon_OverlapPatterns_r16, per.Constrained(1, 3)); err != nil {
					return err
				}
			}

			if ie.OverlapRateMatchingEUTRA_CRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.OverlapRateMatchingEUTRA_CRS_r16, bandNRExtOverlapRateMatchingEUTRACRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_MappingTypeB_Alt_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_MappingTypeB_Alt_r16, bandNRExtPdschMappingTypeBAltR16Constraints); err != nil {
					return err
				}
			}

			if ie.OneSlotPeriodicTRS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.OneSlotPeriodicTRS_r16, bandNRExtOneSlotPeriodicTRSR16Constraints); err != nil {
					return err
				}
			}

			if ie.Olpc_SRS_Pos_r16 != nil {
				if err := ie.Olpc_SRS_Pos_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SpatialRelationsSRS_Pos_r16 != nil {
				if err := ie.SpatialRelationsSRS_Pos_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SimulSRS_MIMO_TransWithinBand_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimulSRS_MIMO_TransWithinBand_r16, bandNRExtSimulSRSMIMOTransWithinBandR16Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBW_DL_IAB_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWDLIABR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBW_DL_IAB_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBW_DL_IAB_r16).Choice {
				case BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz:
					c := (*(*ie.ChannelBW_DL_IAB_r16).Fr1_100mhz)
					bandNRExtChannelBWDLIABR16Fr1100mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWDLIABR16Fr1100mhzConstraints)
					if err := bandNRExtChannelBWDLIABR16Fr1100mhzSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_15kHz), bandNRExtChannelBWDLIABR16Fr1100mhzScs15kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_30kHz), bandNRExtChannelBWDLIABR16Fr1100mhzScs30kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWDLIABR16Fr1100mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz:
					c := (*(*ie.ChannelBW_DL_IAB_r16).Fr2_200mhz)
					bandNRExtChannelBWDLIABR16Fr2200mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWDLIABR16Fr2200mhzConstraints)
					if err := bandNRExtChannelBWDLIABR16Fr2200mhzSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWDLIABR16Fr2200mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_120kHz), bandNRExtChannelBWDLIABR16Fr2200mhzScs120kHzConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.ChannelBW_UL_IAB_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWULIABR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBW_UL_IAB_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBW_UL_IAB_r16).Choice {
				case BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz:
					c := (*(*ie.ChannelBW_UL_IAB_r16).Fr1_100mhz)
					bandNRExtChannelBWULIABR16Fr1100mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWULIABR16Fr1100mhzConstraints)
					if err := bandNRExtChannelBWULIABR16Fr1100mhzSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_15kHz), bandNRExtChannelBWULIABR16Fr1100mhzScs15kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_30kHz), bandNRExtChannelBWULIABR16Fr1100mhzScs30kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWULIABR16Fr1100mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz:
					c := (*(*ie.ChannelBW_UL_IAB_r16).Fr2_200mhz)
					bandNRExtChannelBWULIABR16Fr2200mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWULIABR16Fr2200mhzConstraints)
					if err := bandNRExtChannelBWULIABR16Fr2200mhzSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWULIABR16Fr2200mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_120kHz), bandNRExtChannelBWULIABR16Fr2200mhzScs120kHzConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.RasterShift7dot5_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.RasterShift7dot5_IAB_r16, bandNRExtRasterShift7dot5IABR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ue_PowerClass_v1610 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_PowerClass_v1610, bandNRExtUePowerClassV1610Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandover_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CondHandover_r16, bandNRExtCondHandoverR16Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandoverFailure_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CondHandoverFailure_r16, bandNRExtCondHandoverFailureR16Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandoverTwoTriggerEvents_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CondHandoverTwoTriggerEvents_r16, bandNRExtCondHandoverTwoTriggerEventsR16Constraints); err != nil {
					return err
				}
			}

			if ie.CondPSCellChange_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CondPSCellChange_r16, bandNRExtCondPSCellChangeR16Constraints); err != nil {
					return err
				}
			}

			if ie.CondPSCellChangeTwoTriggerEvents_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CondPSCellChangeTwoTriggerEvents_r16, bandNRExtCondPSCellChangeTwoTriggerEventsR16Constraints); err != nil {
					return err
				}
			}

			if ie.Mpr_PowerBoost_FR2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mpr_PowerBoost_FR2_r16, bandNRExtMprPowerBoostFR2R16Constraints); err != nil {
					return err
				}
			}

			if ie.ActiveConfiguredGrant_r16 != nil {
				c := ie.ActiveConfiguredGrant_r16
				if err := ex.EncodeEnumerated(c.MaxNumberConfigsPerBWP_r16, bandNRExtActiveConfiguredGrantR16MaxNumberConfigsPerBWPR16Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigsAllCC_r16, per.Constrained(2, 32)); err != nil {
					return err
				}
			}

			if ie.JointReleaseConfiguredGrantType2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.JointReleaseConfiguredGrantType2_r16, bandNRExtJointReleaseConfiguredGrantType2R16Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_r16 != nil {
				c := ie.Sps_r16
				if err := ex.EncodeInteger(c.MaxNumberConfigsPerBWP_r16, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigsAllCC_r16, per.Constrained(2, 32)); err != nil {
					return err
				}
			}

			if ie.JointReleaseSPS_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.JointReleaseSPS_r16, bandNRExtJointReleaseSPSR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimulSRS_TransWithinBand_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimulSRS_TransWithinBand_r16, bandNRExtSimulSRSTransWithinBandR16Constraints); err != nil {
					return err
				}
			}

			if ie.Trs_AdditionalBandwidth_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Trs_AdditionalBandwidth_r16, bandNRExtTrsAdditionalBandwidthR16Constraints); err != nil {
					return err
				}
			}

			if ie.HandoverIntraF_IAB_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverIntraF_IAB_r16, bandNRExtHandoverIntraFIABR16Constraints); err != nil {
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
					{Name: "simulTX-SRS-AntSwitchingIntraBandUL-CA-r16", Optional: true},
					{Name: "sharedSpectrumChAccessParamsPerBand-v1630", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil}); err != nil {
				return err
			}

			if ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil {
				if err := ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SharedSpectrumChAccessParamsPerBand_v1630 != nil {
				if err := ie.SharedSpectrumChAccessParamsPerBand_v1630.Encode(ex); err != nil {
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
					{Name: "handoverUTRA-FDD-r16", Optional: true},
					{Name: "enhancedUL-TransientPeriod-r16", Optional: true},
					{Name: "sharedSpectrumChAccessParamsPerBand-v1640", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.HandoverUTRA_FDD_r16 != nil, ie.EnhancedUL_TransientPeriod_r16 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil}); err != nil {
				return err
			}

			if ie.HandoverUTRA_FDD_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverUTRA_FDD_r16, bandNRExtHandoverUTRAFDDR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedUL_TransientPeriod_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedUL_TransientPeriod_r16, bandNRExtEnhancedULTransientPeriodR16Constraints); err != nil {
					return err
				}
			}

			if ie.SharedSpectrumChAccessParamsPerBand_v1640 != nil {
				if err := ie.SharedSpectrumChAccessParamsPerBand_v1640.Encode(ex); err != nil {
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
					{Name: "type1-PUSCH-RepetitionMultiSlots-v1650", Optional: true},
					{Name: "type2-PUSCH-RepetitionMultiSlots-v1650", Optional: true},
					{Name: "pusch-RepetitionMultiSlots-v1650", Optional: true},
					{Name: "configuredUL-GrantType1-v1650", Optional: true},
					{Name: "configuredUL-GrantType2-v1650", Optional: true},
					{Name: "sharedSpectrumChAccessParamsPerBand-v1650", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.Pusch_RepetitionMultiSlots_v1650 != nil, ie.ConfiguredUL_GrantType1_v1650 != nil, ie.ConfiguredUL_GrantType2_v1650 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil}); err != nil {
				return err
			}

			if ie.Type1_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err := ex.EncodeEnumerated(*ie.Type1_PUSCH_RepetitionMultiSlots_v1650, bandNRExtType1PUSCHRepetitionMultiSlotsV1650Constraints); err != nil {
					return err
				}
			}

			if ie.Type2_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err := ex.EncodeEnumerated(*ie.Type2_PUSCH_RepetitionMultiSlots_v1650, bandNRExtType2PUSCHRepetitionMultiSlotsV1650Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_RepetitionMultiSlots_v1650 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_RepetitionMultiSlots_v1650, bandNRExtPuschRepetitionMultiSlotsV1650Constraints); err != nil {
					return err
				}
			}

			if ie.ConfiguredUL_GrantType1_v1650 != nil {
				if err := ex.EncodeEnumerated(*ie.ConfiguredUL_GrantType1_v1650, bandNRExtConfiguredULGrantType1V1650Constraints); err != nil {
					return err
				}
			}

			if ie.ConfiguredUL_GrantType2_v1650 != nil {
				if err := ex.EncodeEnumerated(*ie.ConfiguredUL_GrantType2_v1650, bandNRExtConfiguredULGrantType2V1650Constraints); err != nil {
					return err
				}
			}

			if ie.SharedSpectrumChAccessParamsPerBand_v1650 != nil {
				if err := ie.SharedSpectrumChAccessParamsPerBand_v1650.Encode(ex); err != nil {
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
					{Name: "enhancedSkipUplinkTxConfigured-v1660", Optional: true},
					{Name: "enhancedSkipUplinkTxDynamic-v1660", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnhancedSkipUplinkTxConfigured_v1660 != nil, ie.EnhancedSkipUplinkTxDynamic_v1660 != nil}); err != nil {
				return err
			}

			if ie.EnhancedSkipUplinkTxConfigured_v1660 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxConfigured_v1660, bandNRExtEnhancedSkipUplinkTxConfiguredV1660Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedSkipUplinkTxDynamic_v1660 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedSkipUplinkTxDynamic_v1660, bandNRExtEnhancedSkipUplinkTxDynamicV1660Constraints); err != nil {
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
					{Name: "maxUplinkDutyCycle-PC1dot5-MPE-FR1-r16", Optional: true},
					{Name: "txDiversity-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil, ie.TxDiversity_r16 != nil}); err != nil {
				return err
			}

			if ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16, bandNRExtMaxUplinkDutyCyclePC1dot5MPEFR1R16Constraints); err != nil {
					return err
				}
			}

			if ie.TxDiversity_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.TxDiversity_r16, bandNRExtTxDiversityR16Constraints); err != nil {
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
					{Name: "pdsch-1024QAM-FR1-r17", Optional: true},
					{Name: "ue-PowerClass-v1700", Optional: true},
					{Name: "fr2-2-AccessParamsPerBand-r17", Optional: true},
					{Name: "rlm-Relaxation-r17", Optional: true},
					{Name: "bfd-Relaxation-r17", Optional: true},
					{Name: "cg-SDT-r17", Optional: true},
					{Name: "locationBasedCondHandover-r17", Optional: true},
					{Name: "timeBasedCondHandover-r17", Optional: true},
					{Name: "eventA4BasedCondHandover-r17", Optional: true},
					{Name: "mn-InitiatedCondPSCellChangeNRDC-r17", Optional: true},
					{Name: "sn-InitiatedCondPSCellChangeNRDC-r17", Optional: true},
					{Name: "pdcch-SkippingWithoutSSSG-r17", Optional: true},
					{Name: "sssg-Switching-1BitInd-r17", Optional: true},
					{Name: "sssg-Switching-2BitInd-r17", Optional: true},
					{Name: "pdcch-SkippingWithSSSG-r17", Optional: true},
					{Name: "searchSpaceSetGrp-switchCap2-r17", Optional: true},
					{Name: "uplinkPreCompensation-r17", Optional: true},
					{Name: "uplink-TA-Reporting-r17", Optional: true},
					{Name: "max-HARQ-ProcessNumber-r17", Optional: true},
					{Name: "type2-HARQ-Codebook-r17", Optional: true},
					{Name: "type1-HARQ-Codebook-r17", Optional: true},
					{Name: "type3-HARQ-Codebook-r17", Optional: true},
					{Name: "ue-specific-K-Offset-r17", Optional: true},
					{Name: "multiPDSCH-SingleDCI-FR2-1-SCS-120kHz-r17", Optional: true},
					{Name: "multiPUSCH-SingleDCI-FR2-1-SCS-120kHz-r17", Optional: true},
					{Name: "parallelPRS-MeasRRC-Inactive-r17", Optional: true},
					{Name: "nr-UE-TxTEG-ID-MaxSupport-r17", Optional: true},
					{Name: "prs-ProcessingRRC-Inactive-r17", Optional: true},
					{Name: "prs-ProcessingWindowType1A-r17", Optional: true},
					{Name: "prs-ProcessingWindowType1B-r17", Optional: true},
					{Name: "prs-ProcessingWindowType2-r17", Optional: true},
					{Name: "srs-AllPosResourcesRRC-Inactive-r17", Optional: true},
					{Name: "olpc-SRS-PosRRC-Inactive-r17", Optional: true},
					{Name: "spatialRelationsSRS-PosRRC-Inactive-r17", Optional: true},
					{Name: "maxNumberPUSCH-TypeA-Repetition-r17", Optional: true},
					{Name: "puschTypeA-RepetitionsAvailSlot-r17", Optional: true},
					{Name: "tb-ProcessingMultiSlotPUSCH-r17", Optional: true},
					{Name: "tb-ProcessingRepMultiSlotPUSCH-r17", Optional: true},
					{Name: "maxDurationDMRS-Bundling-r17", Optional: true},
					{Name: "pusch-RepetitionMsg3-r17", Optional: true},
					{Name: "sharedSpectrumChAccessParamsPerBand-v1710", Optional: true},
					{Name: "parallelMeasurementWithoutRestriction-r17", Optional: true},
					{Name: "maxNumber-NGSO-SatellitesWithinOneSMTC-r17", Optional: true},
					{Name: "k1-RangeExtension-r17", Optional: true},
					{Name: "aperiodicCSI-RS-FastScellActivation-r17", Optional: true},
					{Name: "aperiodicCSI-RS-AdditionalBandwidth-r17", Optional: true},
					{Name: "bwp-WithoutCD-SSB-OrNCD-SSB-RedCap-r17", Optional: true},
					{Name: "halfDuplexFDD-TypeA-RedCap-r17", Optional: true},
					{Name: "posSRS-RRC-Inactive-OutsideInitialUL-BWP-r17", Optional: true},
					{Name: "channelBWs-DL-SCS-480kHz-FR2-2-r17", Optional: true},
					{Name: "channelBWs-UL-SCS-480kHz-FR2-2-r17", Optional: true},
					{Name: "channelBWs-DL-SCS-960kHz-FR2-2-r17", Optional: true},
					{Name: "channelBWs-UL-SCS-960kHz-FR2-2-r17", Optional: true},
					{Name: "ul-GapFR2-r17", Optional: true},
					{Name: "oneShotHARQ-feedbackTriggeredByDCI-1-2-r17", Optional: true},
					{Name: "oneShotHARQ-feedbackPhy-Priority-r17", Optional: true},
					{Name: "enhancedType3-HARQ-CodebookFeedback-r17", Optional: true},
					{Name: "triggeredHARQ-CodebookRetx-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_1024QAM_FR1_r17 != nil, ie.Ue_PowerClass_v1700 != nil, ie.Fr2_2_AccessParamsPerBand_r17 != nil, ie.Rlm_Relaxation_r17 != nil, ie.Bfd_Relaxation_r17 != nil, ie.Cg_SDT_r17 != nil, ie.LocationBasedCondHandover_r17 != nil, ie.TimeBasedCondHandover_r17 != nil, ie.EventA4BasedCondHandover_r17 != nil, ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.Pdcch_SkippingWithoutSSSG_r17 != nil, ie.Sssg_Switching_1BitInd_r17 != nil, ie.Sssg_Switching_2BitInd_r17 != nil, ie.Pdcch_SkippingWithSSSG_r17 != nil, ie.SearchSpaceSetGrp_SwitchCap2_r17 != nil, ie.UplinkPreCompensation_r17 != nil, ie.Uplink_TA_Reporting_r17 != nil, ie.Max_HARQ_ProcessNumber_r17 != nil, ie.Type2_HARQ_Codebook_r17 != nil, ie.Type1_HARQ_Codebook_r17 != nil, ie.Type3_HARQ_Codebook_r17 != nil, ie.Ue_Specific_K_Offset_r17 != nil, ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.ParallelPRS_MeasRRC_Inactive_r17 != nil, ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil, ie.Prs_ProcessingRRC_Inactive_r17 != nil, ie.Prs_ProcessingWindowType1A_r17 != nil, ie.Prs_ProcessingWindowType1B_r17 != nil, ie.Prs_ProcessingWindowType2_r17 != nil, ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil, ie.Olpc_SRS_PosRRC_Inactive_r17 != nil, ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil, ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil, ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil, ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil, ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil, ie.MaxDurationDMRS_Bundling_r17 != nil, ie.Pusch_RepetitionMsg3_r17 != nil, ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil, ie.ParallelMeasurementWithoutRestriction_r17 != nil, ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil, ie.K1_RangeExtension_r17 != nil, ie.AperiodicCSI_RS_FastScellActivation_r17 != nil, ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil, ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil, ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil, ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil, ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil, ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil, ie.Ul_GapFR2_r17 != nil, ie.OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17 != nil, ie.OneShotHARQ_FeedbackPhy_Priority_r17 != nil, ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil, ie.TriggeredHARQ_CodebookRetx_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_1024QAM_FR1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_1024QAM_FR1_r17, bandNRExtPdsch1024QAMFR1R17Constraints); err != nil {
					return err
				}
			}

			if ie.Ue_PowerClass_v1700 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_PowerClass_v1700, bandNRExtUePowerClassV1700Constraints); err != nil {
					return err
				}
			}

			if ie.Fr2_2_AccessParamsPerBand_r17 != nil {
				if err := ie.Fr2_2_AccessParamsPerBand_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Rlm_Relaxation_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rlm_Relaxation_r17, bandNRExtRlmRelaxationR17Constraints); err != nil {
					return err
				}
			}

			if ie.Bfd_Relaxation_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Bfd_Relaxation_r17, bandNRExtBfdRelaxationR17Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_SDT_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_SDT_r17, bandNRExtCgSDTR17Constraints); err != nil {
					return err
				}
			}

			if ie.LocationBasedCondHandover_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.LocationBasedCondHandover_r17, bandNRExtLocationBasedCondHandoverR17Constraints); err != nil {
					return err
				}
			}

			if ie.TimeBasedCondHandover_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TimeBasedCondHandover_r17, bandNRExtTimeBasedCondHandoverR17Constraints); err != nil {
					return err
				}
			}

			if ie.EventA4BasedCondHandover_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.EventA4BasedCondHandover_r17, bandNRExtEventA4BasedCondHandoverR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mn_InitiatedCondPSCellChangeNRDC_r17, bandNRExtMnInitiatedCondPSCellChangeNRDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sn_InitiatedCondPSCellChangeNRDC_r17, bandNRExtSnInitiatedCondPSCellChangeNRDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_SkippingWithoutSSSG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_SkippingWithoutSSSG_r17, bandNRExtPdcchSkippingWithoutSSSGR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sssg_Switching_1BitInd_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sssg_Switching_1BitInd_r17, bandNRExtSssgSwitching1BitIndR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sssg_Switching_2BitInd_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sssg_Switching_2BitInd_r17, bandNRExtSssgSwitching2BitIndR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_SkippingWithSSSG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_SkippingWithSSSG_r17, bandNRExtPdcchSkippingWithSSSGR17Constraints); err != nil {
					return err
				}
			}

			if ie.SearchSpaceSetGrp_SwitchCap2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SearchSpaceSetGrp_SwitchCap2_r17, bandNRExtSearchSpaceSetGrpSwitchCap2R17Constraints); err != nil {
					return err
				}
			}

			if ie.UplinkPreCompensation_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.UplinkPreCompensation_r17, bandNRExtUplinkPreCompensationR17Constraints); err != nil {
					return err
				}
			}

			if ie.Uplink_TA_Reporting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Uplink_TA_Reporting_r17, bandNRExtUplinkTAReportingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Max_HARQ_ProcessNumber_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Max_HARQ_ProcessNumber_r17, bandNRExtMaxHARQProcessNumberR17Constraints); err != nil {
					return err
				}
			}

			if ie.Type2_HARQ_Codebook_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Type2_HARQ_Codebook_r17, bandNRExtType2HARQCodebookR17Constraints); err != nil {
					return err
				}
			}

			if ie.Type1_HARQ_Codebook_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Type1_HARQ_Codebook_r17, bandNRExtType1HARQCodebookR17Constraints); err != nil {
					return err
				}
			}

			if ie.Type3_HARQ_Codebook_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Type3_HARQ_Codebook_r17, bandNRExtType3HARQCodebookR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ue_Specific_K_Offset_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_Specific_K_Offset_r17, bandNRExtUeSpecificKOffsetR17Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17, bandNRExtMultiPDSCHSingleDCIFR21SCS120kHzR17Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17, bandNRExtMultiPUSCHSingleDCIFR21SCS120kHzR17Constraints); err != nil {
					return err
				}
			}

			if ie.ParallelPRS_MeasRRC_Inactive_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ParallelPRS_MeasRRC_Inactive_r17, bandNRExtParallelPRSMeasRRCInactiveR17Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_UE_TxTEG_ID_MaxSupport_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_UE_TxTEG_ID_MaxSupport_r17, bandNRExtNrUETxTEGIDMaxSupportR17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_ProcessingRRC_Inactive_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Prs_ProcessingRRC_Inactive_r17, bandNRExtPrsProcessingRRCInactiveR17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_ProcessingWindowType1A_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Prs_ProcessingWindowType1A_r17, bandNRExtPrsProcessingWindowType1AR17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_ProcessingWindowType1B_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Prs_ProcessingWindowType1B_r17, bandNRExtPrsProcessingWindowType1BR17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_ProcessingWindowType2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Prs_ProcessingWindowType2_r17, bandNRExtPrsProcessingWindowType2R17Constraints); err != nil {
					return err
				}
			}

			if ie.Srs_AllPosResourcesRRC_Inactive_r17 != nil {
				if err := ie.Srs_AllPosResourcesRRC_Inactive_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Olpc_SRS_PosRRC_Inactive_r17 != nil {
				if err := ie.Olpc_SRS_PosRRC_Inactive_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SpatialRelationsSRS_PosRRC_Inactive_r17 != nil {
				if err := ie.SpatialRelationsSRS_PosRRC_Inactive_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MaxNumberPUSCH_TypeA_Repetition_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberPUSCH_TypeA_Repetition_r17, bandNRExtMaxNumberPUSCHTypeARepetitionR17Constraints); err != nil {
					return err
				}
			}

			if ie.PuschTypeA_RepetitionsAvailSlot_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PuschTypeA_RepetitionsAvailSlot_r17, bandNRExtPuschTypeARepetitionsAvailSlotR17Constraints); err != nil {
					return err
				}
			}

			if ie.Tb_ProcessingMultiSlotPUSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Tb_ProcessingMultiSlotPUSCH_r17, bandNRExtTbProcessingMultiSlotPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.Tb_ProcessingRepMultiSlotPUSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Tb_ProcessingRepMultiSlotPUSCH_r17, bandNRExtTbProcessingRepMultiSlotPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxDurationDMRS_Bundling_r17 != nil {
				c := ie.MaxDurationDMRS_Bundling_r17
				bandNRExtMaxDurationDMRSBundlingR17Seq := ex.NewSequenceEncoder(bandNRExtMaxDurationDMRSBundlingR17Constraints)
				if err := bandNRExtMaxDurationDMRSBundlingR17Seq.EncodePreamble([]bool{c.Fdd_r17 != nil, c.Tdd_r17 != nil}); err != nil {
					return err
				}
				if c.Fdd_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Fdd_r17), bandNRExtMaxDurationDMRSBundlingR17FddR17Constraints); err != nil {
						return err
					}
				}
				if c.Tdd_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Tdd_r17), bandNRExtMaxDurationDMRSBundlingR17TddR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Pusch_RepetitionMsg3_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_RepetitionMsg3_r17, bandNRExtPuschRepetitionMsg3R17Constraints); err != nil {
					return err
				}
			}

			if ie.SharedSpectrumChAccessParamsPerBand_v1710 != nil {
				if err := ie.SharedSpectrumChAccessParamsPerBand_v1710.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ParallelMeasurementWithoutRestriction_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ParallelMeasurementWithoutRestriction_r17, bandNRExtParallelMeasurementWithoutRestrictionR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17, bandNRExtMaxNumberNGSOSatellitesWithinOneSMTCR17Constraints); err != nil {
					return err
				}
			}

			if ie.K1_RangeExtension_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.K1_RangeExtension_r17, bandNRExtK1RangeExtensionR17Constraints); err != nil {
					return err
				}
			}

			if ie.AperiodicCSI_RS_FastScellActivation_r17 != nil {
				c := ie.AperiodicCSI_RS_FastScellActivation_r17
				if err := ex.EncodeEnumerated(c.MaxNumberAperiodicCSI_RS_PerCC_r17, bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSPerCCR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberAperiodicCSI_RS_AcrossCCs_r17, bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSAcrossCCsR17Constraints); err != nil {
					return err
				}
			}

			if ie.AperiodicCSI_RS_AdditionalBandwidth_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.AperiodicCSI_RS_AdditionalBandwidth_r17, bandNRExtAperiodicCSIRSAdditionalBandwidthR17Constraints); err != nil {
					return err
				}
			}

			if ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17, bandNRExtBwpWithoutCDSSBOrNCDSSBRedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.HalfDuplexFDD_TypeA_RedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.HalfDuplexFDD_TypeA_RedCap_r17, bandNRExtHalfDuplexFDDTypeARedCapR17Constraints); err != nil {
					return err
				}
			}

			if ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil {
				if err := ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17, bandNRChannelBWsDLSCS480kHzFR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17, bandNRChannelBWsULSCS480kHzFR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17, bandNRChannelBWsDLSCS960kHzFR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17, bandNRChannelBWsULSCS960kHzFR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_GapFR2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_GapFR2_r17, bandNRExtUlGapFR2R17Constraints); err != nil {
					return err
				}
			}

			if ie.OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17, bandNRExtOneShotHARQFeedbackTriggeredByDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.OneShotHARQ_FeedbackPhy_Priority_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.OneShotHARQ_FeedbackPhy_Priority_r17, bandNRExtOneShotHARQFeedbackPhyPriorityR17Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedType3_HARQ_CodebookFeedback_r17 != nil {
				c := ie.EnhancedType3_HARQ_CodebookFeedback_r17
				if err := ex.EncodeEnumerated(c.EnhancedType3_HARQ_Codebooks_r17, bandNRExtEnhancedType3HARQCodebookFeedbackR17EnhancedType3HARQCodebooksR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberPUCCH_Transmissions_r17, bandNRExtEnhancedType3HARQCodebookFeedbackR17MaxNumberPUCCHTransmissionsR17Constraints); err != nil {
					return err
				}
			}

			if ie.TriggeredHARQ_CodebookRetx_r17 != nil {
				c := ie.TriggeredHARQ_CodebookRetx_r17
				if err := ex.EncodeEnumerated(c.MinHARQ_Retx_Offset_r17, bandNRExtTriggeredHARQCodebookRetxR17MinHARQRetxOffsetR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxHARQ_Retx_Offset_r17, bandNRExtTriggeredHARQCodebookRetxR17MaxHARQRetxOffsetR17Constraints); err != nil {
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
					{Name: "ue-OneShotUL-TimingAdj-r17", Optional: true},
					{Name: "pucch-Repetition-F0-2-r17", Optional: true},
					{Name: "cqi-4-BitsSubbandNTN-SharedSpectrumChAccess-r17", Optional: true},
					{Name: "mux-HARQ-ACK-DiffPriorities-r17", Optional: true},
					{Name: "ta-BasedPDC-NTN-SharedSpectrumChAccess-r17", Optional: true},
					{Name: "ack-NACK-FeedbackForMulticastWithDCI-Enabler-r17", Optional: true},
					{Name: "maxNumberG-RNTI-r17", Optional: true},
					{Name: "dynamicMulticastDCI-Format4-2-r17", Optional: true},
					{Name: "maxModulationOrderForMulticast-r17", Optional: true},
					{Name: "dynamicSlotRepetitionMulticastTN-NonSharedSpectrumChAccess-r17", Optional: true},
					{Name: "dynamicSlotRepetitionMulticastNTN-SharedSpectrumChAccess-r17", Optional: true},
					{Name: "nack-OnlyFeedbackForMulticastWithDCI-Enabler-r17", Optional: true},
					{Name: "ack-NACK-FeedbackForSPS-MulticastWithDCI-Enabler-r17", Optional: true},
					{Name: "maxNumberG-CS-RNTI-r17", Optional: true},
					{Name: "re-LevelRateMatchingForMulticast-r17", Optional: true},
					{Name: "pdsch-1024QAM-2MIMO-FR1-r17", Optional: true},
					{Name: "prs-MeasurementWithoutMG-r17", Optional: true},
					{Name: "maxNumber-NGSO-SatellitesPerCarrier-r17", Optional: true},
					{Name: "prs-ProcessingCapabilityOutsideMGinPPW-r17", Optional: true},
					{Name: "srs-SemiPersistent-PosResourcesRRC-Inactive-r17", Optional: true},
					{Name: "channelBWs-DL-SCS-120kHz-FR2-2-r17", Optional: true},
					{Name: "channelBWs-UL-SCS-120kHz-FR2-2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ue_OneShotUL_TimingAdj_r17 != nil, ie.Pucch_Repetition_F0_2_r17 != nil, ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil, ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil, ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil, ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.MaxNumberG_RNTI_r17 != nil, ie.DynamicMulticastDCI_Format4_2_r17 != nil, ie.MaxModulationOrderForMulticast_r17 != nil, ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil, ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil, ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil, ie.MaxNumberG_CS_RNTI_r17 != nil, ie.Re_LevelRateMatchingForMulticast_r17 != nil, ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil, ie.Prs_MeasurementWithoutMG_r17 != nil, ie.MaxNumber_NGSO_SatellitesPerCarrier_r17 != nil, ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 != nil, ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil, ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil, ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil}); err != nil {
				return err
			}

			if ie.Ue_OneShotUL_TimingAdj_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_OneShotUL_TimingAdj_r17, bandNRExtUeOneShotULTimingAdjR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_Repetition_F0_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_Repetition_F0_2_r17, bandNRExtPucchRepetitionF02R17Constraints); err != nil {
					return err
				}
			}

			if ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17, bandNRExtCqi4BitsSubbandNTNSharedSpectrumChAccessR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mux_HARQ_ACK_DiffPriorities_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mux_HARQ_ACK_DiffPriorities_r17, bandNRExtMuxHARQACKDiffPrioritiesR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17, bandNRExtTaBasedPDCNTNSharedSpectrumChAccessR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17, bandNRExtAckNACKFeedbackForMulticastWithDCIEnablerR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberG_RNTI_r17 != nil {
				if err := ex.EncodeInteger(*ie.MaxNumberG_RNTI_r17, bandNRMaxNumberGRNTIR17Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicMulticastDCI_Format4_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicMulticastDCI_Format4_2_r17, bandNRExtDynamicMulticastDCIFormat42R17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxModulationOrderForMulticast_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtMaxModulationOrderForMulticastR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MaxModulationOrderForMulticast_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MaxModulationOrderForMulticast_r17).Choice {
				case BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17:
					if err := ex.EncodeEnumerated((*(*ie.MaxModulationOrderForMulticast_r17).Fr1_r17), bandNRExtMaxModulationOrderForMulticastR17Fr1R17Constraints); err != nil {
						return err
					}
				case BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17:
					if err := ex.EncodeEnumerated((*(*ie.MaxModulationOrderForMulticast_r17).Fr2_r17), bandNRExtMaxModulationOrderForMulticastR17Fr2R17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17, bandNRExtDynamicSlotRepetitionMulticastTNNonSharedSpectrumChAccessR17Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17, bandNRExtDynamicSlotRepetitionMulticastNTNSharedSpectrumChAccessR17Constraints); err != nil {
					return err
				}
			}

			if ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17, bandNRExtNackOnlyFeedbackForMulticastWithDCIEnablerR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17, bandNRExtAckNACKFeedbackForSPSMulticastWithDCIEnablerR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumberG_CS_RNTI_r17 != nil {
				if err := ex.EncodeInteger(*ie.MaxNumberG_CS_RNTI_r17, bandNRMaxNumberGCSRNTIR17Constraints); err != nil {
					return err
				}
			}

			if ie.Re_LevelRateMatchingForMulticast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Re_LevelRateMatchingForMulticast_r17, bandNRExtReLevelRateMatchingForMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_1024QAM_2MIMO_FR1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_1024QAM_2MIMO_FR1_r17, bandNRExtPdsch1024QAM2MIMOFR1R17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_MeasurementWithoutMG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Prs_MeasurementWithoutMG_r17, bandNRExtPrsMeasurementWithoutMGR17Constraints); err != nil {
					return err
				}
			}

			if ie.MaxNumber_NGSO_SatellitesPerCarrier_r17 != nil {
				if err := ex.EncodeInteger(*ie.MaxNumber_NGSO_SatellitesPerCarrier_r17, bandNRMaxNumberNGSOSatellitesPerCarrierR17Constraints); err != nil {
					return err
				}
			}

			if ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(bandNRExtPrsProcessingCapabilityOutsideMGinPPWR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17))); err != nil {
					return err
				}
				for i := range ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 {
					if err := ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil {
				c := ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17
				if err := ex.EncodeEnumerated(c.MaxNumOfSemiPersistentSRSposResources_r17, bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17, bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17, bandNRChannelBWsDLSCS120kHzFR22R17Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 != nil {
				if err := ex.EncodeBitString(*ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17, bandNRChannelBWsULSCS120kHzFR22R17Constraints); err != nil {
					return err
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
					{Name: "dmrs-BundlingPUSCH-RepTypeA-r17", Optional: true},
					{Name: "dmrs-BundlingPUSCH-RepTypeB-r17", Optional: true},
					{Name: "dmrs-BundlingPUSCH-multiSlot-r17", Optional: true},
					{Name: "dmrs-BundlingPUCCH-Rep-r17", Optional: true},
					{Name: "interSlotFreqHopInterSlotBundlingPUSCH-r17", Optional: true},
					{Name: "interSlotFreqHopPUCCH-r17", Optional: true},
					{Name: "dmrs-BundlingRestart-r17", Optional: true},
					{Name: "dmrs-BundlingNonBackToBackTX-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil, ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil, ie.Dmrs_BundlingPUSCH_MultiSlot_r17 != nil, ie.Dmrs_BundlingPUCCH_Rep_r17 != nil, ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil, ie.InterSlotFreqHopPUCCH_r17 != nil, ie.Dmrs_BundlingRestart_r17 != nil, ie.Dmrs_BundlingNonBackToBackTX_r17 != nil}); err != nil {
				return err
			}

			if ie.Dmrs_BundlingPUSCH_RepTypeA_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_RepTypeA_r17, bandNRExtDmrsBundlingPUSCHRepTypeAR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingPUSCH_RepTypeB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_RepTypeB_r17, bandNRExtDmrsBundlingPUSCHRepTypeBR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingPUSCH_MultiSlot_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingPUSCH_MultiSlot_r17, bandNRExtDmrsBundlingPUSCHMultiSlotR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingPUCCH_Rep_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingPUCCH_Rep_r17, bandNRExtDmrsBundlingPUCCHRepR17Constraints); err != nil {
					return err
				}
			}

			if ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17, bandNRExtInterSlotFreqHopInterSlotBundlingPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.InterSlotFreqHopPUCCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.InterSlotFreqHopPUCCH_r17, bandNRExtInterSlotFreqHopPUCCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingRestart_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingRestart_r17, bandNRExtDmrsBundlingRestartR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingNonBackToBackTX_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_BundlingNonBackToBackTX_r17, bandNRExtDmrsBundlingNonBackToBackTXR17Constraints); err != nil {
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
					{Name: "maxDynamicSlotRepetitionForSPS-Multicast-r17", Optional: true},
					{Name: "nack-OnlyFeedbackForSPS-MulticastWithDCI-Enabler-r17", Optional: true},
					{Name: "sps-MulticastDCI-Format4-2-r17", Optional: true},
					{Name: "sps-MulticastMultiConfig-r17", Optional: true},
					{Name: "priorityIndicatorInDCI-Multicast-r17", Optional: true},
					{Name: "priorityIndicatorInDCI-SPS-Multicast-r17", Optional: true},
					{Name: "twoHARQ-ACK-CodebookForUnicastAndMulticast-r17", Optional: true},
					{Name: "multiPUCCH-HARQ-ACK-ForMulticastUnicast-r17", Optional: true},
					{Name: "releaseSPS-MulticastWithCS-RNTI-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxDynamicSlotRepetitionForSPS_Multicast_r17 != nil, ie.Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil, ie.Sps_MulticastDCI_Format4_2_r17 != nil, ie.Sps_MulticastMultiConfig_r17 != nil, ie.PriorityIndicatorInDCI_Multicast_r17 != nil, ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil, ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil, ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil, ie.ReleaseSPS_MulticastWithCS_RNTI_r17 != nil}); err != nil {
				return err
			}

			if ie.MaxDynamicSlotRepetitionForSPS_Multicast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxDynamicSlotRepetitionForSPS_Multicast_r17, bandNRExtMaxDynamicSlotRepetitionForSPSMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17, bandNRExtNackOnlyFeedbackForSPSMulticastWithDCIEnablerR17Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_MulticastDCI_Format4_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Sps_MulticastDCI_Format4_2_r17, bandNRExtSpsMulticastDCIFormat42R17Constraints); err != nil {
					return err
				}
			}

			if ie.Sps_MulticastMultiConfig_r17 != nil {
				if err := ex.EncodeInteger(*ie.Sps_MulticastMultiConfig_r17, bandNRSpsMulticastMultiConfigR17Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorInDCI_Multicast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorInDCI_Multicast_r17, bandNRExtPriorityIndicatorInDCIMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorInDCI_SPS_Multicast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorInDCI_SPS_Multicast_r17, bandNRExtPriorityIndicatorInDCISPSMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17, bandNRExtTwoHARQACKCodebookForUnicastAndMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17, bandNRExtMultiPUCCHHARQACKForMulticastUnicastR17Constraints); err != nil {
					return err
				}
			}

			if ie.ReleaseSPS_MulticastWithCS_RNTI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ReleaseSPS_MulticastWithCS_RNTI_r17, bandNRExtReleaseSPSMulticastWithCSRNTIR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 15:
		if hasExtGroup15 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "posUE-TA-AutoAdjustment-r18", Optional: true},
					{Name: "posSRS-ValidityAreaRRC-InactiveInitialUL-BWP-r18", Optional: true},
					{Name: "posSRS-ValidityAreaRRC-InactiveOutsideInitialUL-BWP-r18", Optional: true},
					{Name: "dl-PRS-MeasurementWithRxFH-RRC-ConnectedForRedCap-r18", Optional: true},
					{Name: "posSRS-TxFH-RRC-ConnectedForRedCap-r18", Optional: true},
					{Name: "posSRS-TxFH-RRC-InactiveForRedCap-r18", Optional: true},
					{Name: "posSRS-BWA-RRC-Inactive-r18", Optional: true},
					{Name: "posJointTriggerBySingleDCI-RRC-Connected-r18", Optional: true},
					{Name: "dl-PRS-MeasurementWithRxFH-RRC-InactiveforRedCap-r18", Optional: true},
					{Name: "dl-PRS-MeasurementWithRxFH-RRC-IdleforRedCap-r18", Optional: true},
					{Name: "spatialAdaptation-CSI-Feedback-r18", Optional: true},
					{Name: "spatialAdaptation-CSI-FeedbackPUSCH-r18", Optional: true},
					{Name: "spatialAdaptation-CSI-FeedbackAperiodic-r18", Optional: true},
					{Name: "spatialAdaptation-CSI-FeedbackPUCCH-r18", Optional: true},
					{Name: "powerAdaptation-CSI-Feedback-r18", Optional: true},
					{Name: "powerAdaptation-CSI-FeedbackPUSCH-r18", Optional: true},
					{Name: "powerAdaptation-CSI-FeedbackAperiodic-r18", Optional: true},
					{Name: "powerAdaptation-CSI-FeedbackPUCCH-r18", Optional: true},
					{Name: "nes-CellDTX-DRX-r18", Optional: true},
					{Name: "nes-CellDTX-DRX-DCI2-9-r18", Optional: true},
					{Name: "mixCodeBookSpatialAdaptation-r18", Optional: true},
					{Name: "simultaneousCSI-SubReportsPerCC-r18", Optional: true},
					{Name: "ntn-DMRS-BundlingNGSO-r18", Optional: true},
					{Name: "ltm-BeamIndicationJointTCI-r18", Optional: true},
					{Name: "dummy-ltm-MAC-CE-JointTCI-r18", Optional: true},
					{Name: "ltm-BeamIndicationSeparateTCI-r18", Optional: true},
					{Name: "dummy-ltm-MAC-CE-SeparateTCI-r18", Optional: true},
					{Name: "rach-EarlyTA-Measurement-r18", Optional: true},
					{Name: "ue-TA-Measurement-r18", Optional: true},
					{Name: "ta-IndicationCellSwitch-r18", Optional: true},
					{Name: "triggeredHARQ-CodebookRetxDCI-1-3-r18", Optional: true},
					{Name: "unifiedJointTCI-MultiMAC-CE-DCI-1-3-r18", Optional: true},
					{Name: "unifiedSeparateTCI-MultiMAC-CE-IntraCell-r18", Optional: true},
					{Name: "multiPUSCH-CG-r18", Optional: true},
					{Name: "multiPUSCH-ActiveConfiguredGrant-r18", Optional: true},
					{Name: "jointReleaseDCI-r18", Optional: true},
					{Name: "cg-PUSCH-UTO-UCI-Ind-r18", Optional: true},
					{Name: "pdcch-MonitoringResumptionAfterUL-NACK-r18", Optional: true},
					{Name: "support3MHz-ChannelBW-Symmetric-r18", Optional: true},
					{Name: "support3MHz-ChannelBW-Asymmetric-r18", Optional: true},
					{Name: "support12PRB-CORESET0-r18", Optional: true},
					{Name: "nr-PDCCH-OverlapLTE-CRS-RE-r18", Optional: true},
					{Name: "nr-PDCCH-OverlapLTE-CRS-RE-MultiPatterns-r18", Optional: true},
					{Name: "nr-PDCCH-OverlapLTE-CRS-RE-Span-3-4-r18", Optional: true},
					{Name: "twoRateMatchingEUTRA-CRS-patterns-3-4-r18", Optional: true},
					{Name: "overlapRateMatchingEUTRA-CRS-Patterns-3-4-Diff-CS-Pool-r18", Optional: true},
					{Name: "ncd-SSB-BWP-Wor-r18", Optional: true},
					{Name: "rlm-BM-BFD-CSI-RS-OutsideActiveBWP-r18", Optional: true},
					{Name: "prach-CoverageEnh-r18", Optional: true},
					{Name: "prach-Repetition-r18", Optional: true},
					{Name: "dynamicWaveformSwitch-r18", Optional: true},
					{Name: "dynamicWaveformSwitchPHR-r18", Optional: true},
					{Name: "dynamicWaveformSwitchIntraCA-r18", Optional: true},
					{Name: "multiPUSCH-SingleDCI-NonConsSlots-r18", Optional: true},
					{Name: "pdc-maxNumberPRS-ResourceProcessedPerSlot-r18", Optional: true},
					{Name: "intraSlot-PDSCH-MulticastInactive-r18", Optional: true},
					{Name: "multicastInactive-r18", Optional: true},
					{Name: "thresholdBasedMulticastResume-r18", Optional: true},
					{Name: "lowerMSD-r18", Optional: true},
					{Name: "lowerMSD-ENDC-r18", Optional: true},
					{Name: "enhancedChannelRaster-r18", Optional: true},
					{Name: "fastBeamSweepingMultiRx-r18", Optional: true},
					{Name: "beamSweepingFactorReduction-r18", Optional: true},
					{Name: "simultaneousReceptionTwoQCL-r18", Optional: true},
					{Name: "measEnhCAInterFreqFR2-r18", Optional: true},
					{Name: "tci-StateSwitchInd-r18", Optional: true},
					{Name: "antennaArrayType-r18", Optional: true},
					{Name: "locationBasedCondHandoverATG-r18", Optional: true},
					{Name: "maxOutputPowerATG-r18", Optional: true},
					{Name: "ltm-FastProcessingConfig-r18", Optional: true},
					{Name: "measValidationReportEMR-r18", Optional: true},
					{Name: "measValidationReportReselectionMeasurements-r18", Optional: true},
					{Name: "eventA4BasedCondHandoverNES-r18", Optional: true},
					{Name: "nesBasedCondHandoverWithDCI-r18", Optional: true},
					{Name: "rach-LessHandoverCG-r18", Optional: true},
					{Name: "rach-LessHandoverDG-r18", Optional: true},
					{Name: "locationBasedCondHandoverEMC-r18", Optional: true},
					{Name: "mt-CG-SDT-r18", Optional: true},
					{Name: "posSRS-PreconfigureRRC-InactiveInitialUL-BWP-r18", Optional: true},
					{Name: "posSRS-PreconfigureRRC-InactiveOutsideInitialUL-BWP-r18", Optional: true},
					{Name: "cg-SDT-PeriodicityExt-r18", Optional: true},
					{Name: "supportOf2RxXR-r18", Optional: true},
					{Name: "condHandoverWithCandSCG-change-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PosUE_TA_AutoAdjustment_r18 != nil, ie.PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18 != nil, ie.PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18 != nil, ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18 != nil, ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18 != nil, ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18 != nil, ie.PosSRS_BWA_RRC_Inactive_r18 != nil, ie.PosJointTriggerBySingleDCI_RRC_Connected_r18 != nil, ie.Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18 != nil, ie.Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18 != nil, ie.SpatialAdaptation_CSI_Feedback_r18 != nil, ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18 != nil, ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18 != nil, ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18 != nil, ie.PowerAdaptation_CSI_Feedback_r18 != nil, ie.PowerAdaptation_CSI_FeedbackPUSCH_r18 != nil, ie.PowerAdaptation_CSI_FeedbackAperiodic_r18 != nil, ie.PowerAdaptation_CSI_FeedbackPUCCH_r18 != nil, ie.Nes_CellDTX_DRX_r18 != nil, ie.Nes_CellDTX_DRX_DCI2_9_r18 != nil, ie.MixCodeBookSpatialAdaptation_r18 != nil, ie.SimultaneousCSI_SubReportsPerCC_r18 != nil, ie.Ntn_DMRS_BundlingNGSO_r18 != nil, ie.Ltm_BeamIndicationJointTCI_r18 != nil, ie.Dummy_Ltm_MAC_CE_JointTCI_r18 != nil, ie.Ltm_BeamIndicationSeparateTCI_r18 != nil, ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18 != nil, ie.Rach_EarlyTA_Measurement_r18 != nil, ie.Ue_TA_Measurement_r18 != nil, ie.Ta_IndicationCellSwitch_r18 != nil, ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18 != nil, ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18 != nil, ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18 != nil, ie.MultiPUSCH_CG_r18 != nil, ie.MultiPUSCH_ActiveConfiguredGrant_r18 != nil, ie.JointReleaseDCI_r18 != nil, ie.Cg_PUSCH_UTO_UCI_Ind_r18 != nil, ie.Pdcch_MonitoringResumptionAfterUL_NACK_r18 != nil, ie.Support3MHz_ChannelBW_Symmetric_r18 != nil, ie.Support3MHz_ChannelBW_Asymmetric_r18 != nil, ie.Support12PRB_CORESET0_r18 != nil, ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18 != nil, ie.Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18 != nil, ie.Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18 != nil, ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18 != nil, ie.OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18 != nil, ie.Ncd_SSB_BWP_Wor_r18 != nil, ie.Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18 != nil, ie.Prach_CoverageEnh_r18 != nil, ie.Prach_Repetition_r18 != nil, ie.DynamicWaveformSwitch_r18 != nil, ie.DynamicWaveformSwitchPHR_r18 != nil, ie.DynamicWaveformSwitchIntraCA_r18 != nil, ie.MultiPUSCH_SingleDCI_NonConsSlots_r18 != nil, ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18 != nil, ie.IntraSlot_PDSCH_MulticastInactive_r18 != nil, ie.MulticastInactive_r18 != nil, ie.ThresholdBasedMulticastResume_r18 != nil, ie.LowerMSD_r18 != nil, ie.LowerMSD_ENDC_r18 != nil, ie.EnhancedChannelRaster_r18 != nil, ie.FastBeamSweepingMultiRx_r18 != nil, ie.BeamSweepingFactorReduction_r18 != nil, ie.SimultaneousReceptionTwoQCL_r18 != nil, ie.MeasEnhCAInterFreqFR2_r18 != nil, ie.Tci_StateSwitchInd_r18 != nil, ie.AntennaArrayType_r18 != nil, ie.LocationBasedCondHandoverATG_r18 != nil, ie.MaxOutputPowerATG_r18 != nil, ie.Ltm_FastProcessingConfig_r18 != nil, ie.MeasValidationReportEMR_r18 != nil, ie.MeasValidationReportReselectionMeasurements_r18 != nil, ie.EventA4BasedCondHandoverNES_r18 != nil, ie.NesBasedCondHandoverWithDCI_r18 != nil, ie.Rach_LessHandoverCG_r18 != nil, ie.Rach_LessHandoverDG_r18 != nil, ie.LocationBasedCondHandoverEMC_r18 != nil, ie.Mt_CG_SDT_r18 != nil, ie.PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18 != nil, ie.PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18 != nil, ie.Cg_SDT_PeriodicityExt_r18 != nil, ie.SupportOf2RxXR_r18 != nil, ie.CondHandoverWithCandSCG_Change_r18 != nil}); err != nil {
				return err
			}

			if ie.PosUE_TA_AutoAdjustment_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosUE_TA_AutoAdjustment_r18, bandNRExtPosUETAAutoAdjustmentR18Constraints); err != nil {
					return err
				}
			}

			if ie.PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18, bandNRExtPosSRSValidityAreaRRCInactiveInitialULBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18, bandNRExtPosSRSValidityAreaRRCInactiveOutsideInitialULBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18 != nil {
				if err := ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18 != nil {
				if err := ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18 != nil {
				if err := ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSRS_BWA_RRC_Inactive_r18 != nil {
				if err := ie.PosSRS_BWA_RRC_Inactive_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosJointTriggerBySingleDCI_RRC_Connected_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosJointTriggerBySingleDCI_RRC_Connected_r18, bandNRExtPosJointTriggerBySingleDCIRRCConnectedR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18, bandNRExtDlPRSMeasurementWithRxFHRRCInactiveforRedCapR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18, bandNRExtDlPRSMeasurementWithRxFHRRCIdleforRedCapR18Constraints); err != nil {
					return err
				}
			}

			if ie.SpatialAdaptation_CSI_Feedback_r18 != nil {
				c := ie.SpatialAdaptation_CSI_Feedback_r18
				if err := ex.EncodeEnumerated(c.CsiFeedbackType_r18, bandNRExtSpatialAdaptationCSIFeedbackR18CsiFeedbackTypeR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				{
					c := &c.MaxNumberCSI_ResourcePerCC_r18
					if err := ex.EncodeInteger(c.SdType1_Resource_r18, per.Constrained(1, 32)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.SdType2_Resource_r18, per.Constrained(1, 32)); err != nil {
						return err
					}
				}
				{
					c := &c.MaxNumberTotalCSI_ResourcePerCC_r18
					if err := ex.EncodeEnumerated(c.SdType1_Resource_r18, bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.SdType2_Resource_r18, bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
			}

			if ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18 != nil {
				c := ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18
				if err := ex.EncodeEnumerated(c.CsiFeedbackType_r18, bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18CsiFeedbackTypeR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 12)); err != nil {
					return err
				}
			}

			if ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18 != nil {
				c := ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18
				if err := ex.EncodeEnumerated(c.CsiFeedbackType_r18, bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18CsiFeedbackTypeR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				{
					c := &c.MaxNumberCSI_ResourcePerCC_r18
					if err := ex.EncodeInteger(c.SdType1_Resource_r18, per.Constrained(1, 32)); err != nil {
						return err
					}
					if err := ex.EncodeInteger(c.SdType2_Resource_r18, per.Constrained(1, 32)); err != nil {
						return err
					}
				}
				{
					c := &c.MaxNumberTotalCSI_ResourcePerCC_r18
					if err := ex.EncodeEnumerated(c.SdType1_Resource_r18, bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints); err != nil {
						return err
					}
					if err := ex.EncodeEnumerated(c.SdType2_Resource_r18, bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 12)); err != nil {
					return err
				}
			}

			if ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18 != nil {
				c := ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18
				if err := ex.EncodeEnumerated(c.CsiFeedbackType_r18, bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18CsiFeedbackTypeR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
			}

			if ie.PowerAdaptation_CSI_Feedback_r18 != nil {
				c := ie.PowerAdaptation_CSI_Feedback_r18
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtPowerAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
			}

			if ie.PowerAdaptation_CSI_FeedbackPUSCH_r18 != nil {
				c := ie.PowerAdaptation_CSI_FeedbackPUSCH_r18
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtPowerAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 12)); err != nil {
					return err
				}
			}

			if ie.PowerAdaptation_CSI_FeedbackAperiodic_r18 != nil {
				c := ie.PowerAdaptation_CSI_FeedbackAperiodic_r18
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtPowerAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 12)); err != nil {
					return err
				}
			}

			if ie.PowerAdaptation_CSI_FeedbackPUCCH_r18 != nil {
				c := ie.PowerAdaptation_CSI_FeedbackPUCCH_r18
				if err := ex.EncodeInteger(c.MaxNumberLmax_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.SubReportCSI_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCSI_ResourcePerCC_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberTotalCSI_ResourcePerCC_r18, bandNRExtPowerAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.TotalNumberCSI_Reporting_r18, per.Constrained(2, 4)); err != nil {
					return err
				}
			}

			if ie.Nes_CellDTX_DRX_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Nes_CellDTX_DRX_r18, bandNRExtNesCellDTXDRXR18Constraints); err != nil {
					return err
				}
			}

			if ie.Nes_CellDTX_DRX_DCI2_9_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Nes_CellDTX_DRX_DCI2_9_r18, bandNRExtNesCellDTXDRXDCI29R18Constraints); err != nil {
					return err
				}
			}

			if ie.MixCodeBookSpatialAdaptation_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MixCodeBookSpatialAdaptation_r18, bandNRExtMixCodeBookSpatialAdaptationR18Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousCSI_SubReportsPerCC_r18 != nil {
				if err := ex.EncodeInteger(*ie.SimultaneousCSI_SubReportsPerCC_r18, bandNRSimultaneousCSISubReportsPerCCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ntn_DMRS_BundlingNGSO_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_DMRS_BundlingNGSO_r18, bandNRExtNtnDMRSBundlingNGSOR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_BeamIndicationJointTCI_r18 != nil {
				c := ie.Ltm_BeamIndicationJointTCI_r18
				if err := ex.EncodeEnumerated(c.MaxNumberJointTCI_PerCell_r18, bandNRExtLtmBeamIndicationJointTCIR18MaxNumberJointTCIPerCellR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtLtmBeamIndicationJointTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberJointTCI_AcrossCells_r18, per.Constrained(1, 128)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCells_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.Dummy_Ltm_MAC_CE_JointTCI_r18 != nil {
				c := ie.Dummy_Ltm_MAC_CE_JointTCI_r18
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtDummyLtmMACCEJointTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberJointTCI_PerCell_r18, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberJointTCI_AcrossCells_r18, bandNRExtDummyLtmMACCEJointTCIR18MaxNumberJointTCIAcrossCellsR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_BeamIndicationSeparateTCI_r18 != nil {
				c := ie.Ltm_BeamIndicationSeparateTCI_r18
				if err := ex.EncodeEnumerated(c.MaxNumberDL_TCI_PerCell_r18, bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberDLTCIPerCellR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberUL_TCI_PerCell_r18, bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberULTCIPerCellR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtLtmBeamIndicationSeparateTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberDL_TCI_AcrossCells_r18, per.Constrained(1, 128)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberUL_TCI_AcrossCells_r18, per.Constrained(1, 64)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberCells_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18 != nil {
				c := ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtDummyLtmMACCESeparateTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberDL_TCI_PerCell_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberUL_TCI_PerCell_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberDL_TCI_AcrossCells_r18, bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberDLTCIAcrossCellsR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberUL_TCI_AcrossCells_r18, bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberULTCIAcrossCellsR18Constraints); err != nil {
					return err
				}
			}

			if ie.Rach_EarlyTA_Measurement_r18 != nil {
				if err := ex.EncodeInteger(*ie.Rach_EarlyTA_Measurement_r18, bandNRRachEarlyTAMeasurementR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ue_TA_Measurement_r18 != nil {
				if err := ex.EncodeInteger(*ie.Ue_TA_Measurement_r18, bandNRUeTAMeasurementR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ta_IndicationCellSwitch_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ta_IndicationCellSwitch_r18, bandNRExtTaIndicationCellSwitchR18Constraints); err != nil {
					return err
				}
			}

			if ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18 != nil {
				c := ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18
				if err := ex.EncodeEnumerated(c.MinHARQ_Retx_Offset_r18, bandNRExtTriggeredHARQCodebookRetxDCI13R18MinHARQRetxOffsetR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxHARQ_Retx_Offset_r18, bandNRExtTriggeredHARQCodebookRetxDCI13R18MaxHARQRetxOffsetR18Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18 != nil {
				c := ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18
				bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Constraints)
				if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Seq.EncodePreamble([]bool{c.MaxActivatedTCI_PerCC_r18 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MinBeamApplicationTime_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.MinBeamApplicationTime_r18.Choice {
					case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18:
						c := (*c.MinBeamApplicationTime_r18.Fr1_r18)
						bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Constraints)
						if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil}); err != nil {
							return err
						}
						if c.Scs_15kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_15kHz_r18), bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_30kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_30kHz_r18), bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_60kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints); err != nil {
								return err
							}
						}
					case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18:
						c := (*c.MinBeamApplicationTime_r18.Fr2_r18)
						bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Constraints)
						if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq.EncodePreamble([]bool{c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
							return err
						}
						if c.Scs_60kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_120kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_120kHz_r18), bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.MaxActivatedTCI_PerCC_r18 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedTCI_PerCC_r18), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
			}

			if ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18 != nil {
				c := ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18
				bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Constraints)
				if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq.EncodePreamble([]bool{c.MaxActivatedDL_TCI_PerCC_r18 != nil, c.MaxActivatedUL_TCI_PerCC_r18 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MinBeamApplicationTime_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.MinBeamApplicationTime_r18.Choice {
					case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18:
						c := (*c.MinBeamApplicationTime_r18.Fr1_r18)
						bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Constraints)
						if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil}); err != nil {
							return err
						}
						if c.Scs_15kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_15kHz_r18), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_30kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_30kHz_r18), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_60kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints); err != nil {
								return err
							}
						}
					case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18:
						c := (*c.MinBeamApplicationTime_r18.Fr2_r18)
						bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Constraints)
						if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq.EncodePreamble([]bool{c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
							return err
						}
						if c.Scs_60kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints); err != nil {
								return err
							}
						}
						if c.Scs_120kHz_r18 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_120kHz_r18), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.MaxActivatedDL_TCI_PerCC_r18 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedDL_TCI_PerCC_r18), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
				if c.MaxActivatedUL_TCI_PerCC_r18 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedUL_TCI_PerCC_r18), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
			}

			if ie.MultiPUSCH_CG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUSCH_CG_r18, bandNRExtMultiPUSCHCGR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUSCH_ActiveConfiguredGrant_r18 != nil {
				c := ie.MultiPUSCH_ActiveConfiguredGrant_r18
				if err := ex.EncodeEnumerated(c.MaxNumberConfigsPerBWP, bandNRExtMultiPUSCHActiveConfiguredGrantR18MaxNumberConfigsPerBWPConstraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigsAllCC_FR1, per.Constrained(2, 32)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigsAllCC_FR2, per.Constrained(2, 32)); err != nil {
					return err
				}
			}

			if ie.JointReleaseDCI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.JointReleaseDCI_r18, bandNRExtJointReleaseDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_PUSCH_UTO_UCI_Ind_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_PUSCH_UTO_UCI_Ind_r18, bandNRExtCgPUSCHUTOUCIIndR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_MonitoringResumptionAfterUL_NACK_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_MonitoringResumptionAfterUL_NACK_r18, bandNRExtPdcchMonitoringResumptionAfterULNACKR18Constraints); err != nil {
					return err
				}
			}

			if ie.Support3MHz_ChannelBW_Symmetric_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Support3MHz_ChannelBW_Symmetric_r18, bandNRExtSupport3MHzChannelBWSymmetricR18Constraints); err != nil {
					return err
				}
			}

			if ie.Support3MHz_ChannelBW_Asymmetric_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Support3MHz_ChannelBW_Asymmetric_r18, bandNRExtSupport3MHzChannelBWAsymmetricR18Constraints); err != nil {
					return err
				}
			}

			if ie.Support12PRB_CORESET0_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Support12PRB_CORESET0_r18, bandNRExtSupport12PRBCORESET0R18Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18 != nil {
				c := ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18
				if err := ex.EncodeEnumerated(c.OverlapInRE_r18, bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInRER18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.OverlapInSymbol_r18, bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInSymbolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18, bandNRExtNrPDCCHOverlapLTECRSREMultiPatternsR18Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18, bandNRExtNrPDCCHOverlapLTECRSRESpan34R18Constraints); err != nil {
					return err
				}
			}

			if ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18 != nil {
				c := ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18
				if err := ex.EncodeInteger(c.MaxNumberPatterns_r18, per.Constrained(2, 6)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberNon_OverlapPatterns_r18, per.Constrained(1, 3)); err != nil {
					return err
				}
			}

			if ie.OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18, bandNRExtOverlapRateMatchingEUTRACRSPatterns34DiffCSPoolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ncd_SSB_BWP_Wor_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncd_SSB_BWP_Wor_r18, bandNRExtNcdSSBBWPWorR18Constraints); err != nil {
					return err
				}
			}

			if ie.Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18, bandNRExtRlmBMBFDCSIRSOutsideActiveBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.Prach_CoverageEnh_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Prach_CoverageEnh_r18, bandNRExtPrachCoverageEnhR18Constraints); err != nil {
					return err
				}
			}

			if ie.Prach_Repetition_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Prach_Repetition_r18, bandNRExtPrachRepetitionR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicWaveformSwitch_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicWaveformSwitch_r18, bandNRExtDynamicWaveformSwitchR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicWaveformSwitchPHR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicWaveformSwitchPHR_r18, bandNRExtDynamicWaveformSwitchPHRR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicWaveformSwitchIntraCA_r18 != nil {
				if err := ex.EncodeInteger(*ie.DynamicWaveformSwitchIntraCA_r18, bandNRDynamicWaveformSwitchIntraCAR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultiPUSCH_SingleDCI_NonConsSlots_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPUSCH_SingleDCI_NonConsSlots_r18, bandNRExtMultiPUSCHSingleDCINonConsSlotsR18Constraints); err != nil {
					return err
				}
			}

			if ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18 != nil {
				c := ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18
				{
					c := &c.Fr1_r18
					bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq := ex.NewSequenceEncoder(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Constraints)
					if err := bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq.EncodePreamble([]bool{c.Scs_15kHz_r18 != nil, c.Scs_30kHz_r18 != nil, c.Scs_60kHz_r18 != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Scs_15kHz_r18), bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs15kHzR18Constraints); err != nil {
							return err
						}
					}
					if c.Scs_30kHz_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Scs_30kHz_r18), bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs30kHzR18Constraints); err != nil {
							return err
						}
					}
					if c.Scs_60kHz_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs60kHzR18Constraints); err != nil {
							return err
						}
					}
				}
				{
					c := &c.Fr2_r18
					bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq := ex.NewSequenceEncoder(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Constraints)
					if err := bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq.EncodePreamble([]bool{c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz_r18), bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs60kHzR18Constraints); err != nil {
							return err
						}
					}
					if c.Scs_120kHz_r18 != nil {
						if err := ex.EncodeEnumerated((*c.Scs_120kHz_r18), bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs120kHzR18Constraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.IntraSlot_PDSCH_MulticastInactive_r18 != nil {
				if err := ex.EncodeBoolean(*ie.IntraSlot_PDSCH_MulticastInactive_r18); err != nil {
					return err
				}
			}

			if ie.MulticastInactive_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MulticastInactive_r18, bandNRExtMulticastInactiveR18Constraints); err != nil {
					return err
				}
			}

			if ie.ThresholdBasedMulticastResume_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ThresholdBasedMulticastResume_r18, bandNRExtThresholdBasedMulticastResumeR18Constraints); err != nil {
					return err
				}
			}

			if ie.LowerMSD_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(bandNRExtLowerMSDR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.LowerMSD_r18))); err != nil {
					return err
				}
				for i := range ie.LowerMSD_r18 {
					if err := ie.LowerMSD_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.LowerMSD_ENDC_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(bandNRExtLowerMSDENDCR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.LowerMSD_ENDC_r18))); err != nil {
					return err
				}
				for i := range ie.LowerMSD_ENDC_r18 {
					if err := ie.LowerMSD_ENDC_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.EnhancedChannelRaster_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedChannelRaster_r18, bandNRExtEnhancedChannelRasterR18Constraints); err != nil {
					return err
				}
			}

			if ie.FastBeamSweepingMultiRx_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.FastBeamSweepingMultiRx_r18, bandNRExtFastBeamSweepingMultiRxR18Constraints); err != nil {
					return err
				}
			}

			if ie.BeamSweepingFactorReduction_r18 != nil {
				c := ie.BeamSweepingFactorReduction_r18
				if err := ex.EncodeEnumerated(c.ReduceForCellDetection, bandNRExtBeamSweepingFactorReductionR18ReduceForCellDetectionConstraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.ReduceForSSB_L1_RSRP_Meas, per.Constrained(0, 7)); err != nil {
					return err
				}
			}

			if ie.SimultaneousReceptionTwoQCL_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousReceptionTwoQCL_r18, bandNRExtSimultaneousReceptionTwoQCLR18Constraints); err != nil {
					return err
				}
			}

			if ie.MeasEnhCAInterFreqFR2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasEnhCAInterFreqFR2_r18, bandNRExtMeasEnhCAInterFreqFR2R18Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_StateSwitchInd_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Tci_StateSwitchInd_r18, bandNRExtTciStateSwitchIndR18Constraints); err != nil {
					return err
				}
			}

			if ie.AntennaArrayType_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.AntennaArrayType_r18, bandNRExtAntennaArrayTypeR18Constraints); err != nil {
					return err
				}
			}

			if ie.LocationBasedCondHandoverATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.LocationBasedCondHandoverATG_r18, bandNRExtLocationBasedCondHandoverATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.MaxOutputPowerATG_r18 != nil {
				if err := ex.EncodeInteger(*ie.MaxOutputPowerATG_r18, bandNRMaxOutputPowerATGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_FastProcessingConfig_r18 != nil {
				c := ie.Ltm_FastProcessingConfig_r18
				if err := ex.EncodeEnumerated(c.MaxNumberStoredConfigCells_r18, bandNRExtLtmFastProcessingConfigR18MaxNumberStoredConfigCellsR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberConfigs_r18, per.Constrained(1, 4)); err != nil {
					return err
				}
			}

			if ie.MeasValidationReportEMR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasValidationReportEMR_r18, bandNRExtMeasValidationReportEMRR18Constraints); err != nil {
					return err
				}
			}

			if ie.MeasValidationReportReselectionMeasurements_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MeasValidationReportReselectionMeasurements_r18, bandNRExtMeasValidationReportReselectionMeasurementsR18Constraints); err != nil {
					return err
				}
			}

			if ie.EventA4BasedCondHandoverNES_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EventA4BasedCondHandoverNES_r18, bandNRExtEventA4BasedCondHandoverNESR18Constraints); err != nil {
					return err
				}
			}

			if ie.NesBasedCondHandoverWithDCI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.NesBasedCondHandoverWithDCI_r18, bandNRExtNesBasedCondHandoverWithDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.Rach_LessHandoverCG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Rach_LessHandoverCG_r18, bandNRExtRachLessHandoverCGR18Constraints); err != nil {
					return err
				}
			}

			if ie.Rach_LessHandoverDG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Rach_LessHandoverDG_r18, bandNRExtRachLessHandoverDGR18Constraints); err != nil {
					return err
				}
			}

			if ie.LocationBasedCondHandoverEMC_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.LocationBasedCondHandoverEMC_r18, bandNRExtLocationBasedCondHandoverEMCR18Constraints); err != nil {
					return err
				}
			}

			if ie.Mt_CG_SDT_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Mt_CG_SDT_r18, bandNRExtMtCGSDTR18Constraints); err != nil {
					return err
				}
			}

			if ie.PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18, bandNRExtPosSRSPreconfigureRRCInactiveInitialULBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18, bandNRExtPosSRSPreconfigureRRCInactiveOutsideInitialULBWPR18Constraints); err != nil {
					return err
				}
			}

			if ie.Cg_SDT_PeriodicityExt_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Cg_SDT_PeriodicityExt_r18, bandNRExtCgSDTPeriodicityExtR18Constraints); err != nil {
					return err
				}
			}

			if ie.SupportOf2RxXR_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.SupportOf2RxXR_r18, bandNRExtSupportOf2RxXRR18Constraints); err != nil {
					return err
				}
			}

			if ie.CondHandoverWithCandSCG_Change_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.CondHandoverWithCandSCG_Change_r18, bandNRExtCondHandoverWithCandSCGChangeR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 16:
		if hasExtGroup16 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "mac-ParametersPerBand-r18", Optional: true},
					{Name: "channelBW-DL-NCR-r18", Optional: true},
					{Name: "channelBW-UL-NCR-r18", Optional: true},
					{Name: "ncr-PDSCH-64QAM-FR2-r18", Optional: true},
					{Name: "ltm-MCG-IntraFreq-r18", Optional: true},
					{Name: "ltm-SCG-IntraFreq-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mac_ParametersPerBand_r18 != nil, ie.ChannelBW_DL_NCR_r18 != nil, ie.ChannelBW_UL_NCR_r18 != nil, ie.Ncr_PDSCH_64QAM_FR2_r18 != nil, ie.Ltm_MCG_IntraFreq_r18 != nil, ie.Ltm_SCG_IntraFreq_r18 != nil}); err != nil {
				return err
			}

			if ie.Mac_ParametersPerBand_r18 != nil {
				if err := ie.Mac_ParametersPerBand_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ChannelBW_DL_NCR_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWDLNCRR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBW_DL_NCR_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBW_DL_NCR_r18).Choice {
				case BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz:
					c := (*(*ie.ChannelBW_DL_NCR_r18).Fr1_100mhz)
					bandNRExtChannelBWDLNCRR18Fr1100mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWDLNCRR18Fr1100mhzConstraints)
					if err := bandNRExtChannelBWDLNCRR18Fr1100mhzSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_15kHz), bandNRExtChannelBWDLNCRR18Fr1100mhzScs15kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_30kHz), bandNRExtChannelBWDLNCRR18Fr1100mhzScs30kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWDLNCRR18Fr1100mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz:
					c := (*(*ie.ChannelBW_DL_NCR_r18).Fr2_200mhz)
					bandNRExtChannelBWDLNCRR18Fr2200mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWDLNCRR18Fr2200mhzConstraints)
					if err := bandNRExtChannelBWDLNCRR18Fr2200mhzSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWDLNCRR18Fr2200mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_120kHz), bandNRExtChannelBWDLNCRR18Fr2200mhzScs120kHzConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.ChannelBW_UL_NCR_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtChannelBWULNCRR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelBW_UL_NCR_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelBW_UL_NCR_r18).Choice {
				case BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz:
					c := (*(*ie.ChannelBW_UL_NCR_r18).Fr1_100mhz)
					bandNRExtChannelBWULNCRR18Fr1100mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWULNCRR18Fr1100mhzConstraints)
					if err := bandNRExtChannelBWULNCRR18Fr1100mhzSeq.EncodePreamble([]bool{c.Scs_15kHz != nil, c.Scs_30kHz != nil, c.Scs_60kHz != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_15kHz), bandNRExtChannelBWULNCRR18Fr1100mhzScs15kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_30kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_30kHz), bandNRExtChannelBWULNCRR18Fr1100mhzScs30kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWULNCRR18Fr1100mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
				case BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz:
					c := (*(*ie.ChannelBW_UL_NCR_r18).Fr2_200mhz)
					bandNRExtChannelBWULNCRR18Fr2200mhzSeq := ex.NewSequenceEncoder(bandNRExtChannelBWULNCRR18Fr2200mhzConstraints)
					if err := bandNRExtChannelBWULNCRR18Fr2200mhzSeq.EncodePreamble([]bool{c.Scs_60kHz != nil, c.Scs_120kHz != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_60kHz), bandNRExtChannelBWULNCRR18Fr2200mhzScs60kHzConstraints); err != nil {
							return err
						}
					}
					if c.Scs_120kHz != nil {
						if err := ex.EncodeEnumerated((*c.Scs_120kHz), bandNRExtChannelBWULNCRR18Fr2200mhzScs120kHzConstraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.Ncr_PDSCH_64QAM_FR2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ncr_PDSCH_64QAM_FR2_r18, bandNRExtNcrPDSCH64QAMFR2R18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MCG_IntraFreq_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MCG_IntraFreq_r18, bandNRExtLtmMCGIntraFreqR18Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_SCG_IntraFreq_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_SCG_IntraFreq_r18, bandNRExtLtmSCGIntraFreqR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 17:
		if hasExtGroup17 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "ltm-MAC-CE-JointTCI-r18", Optional: true},
					{Name: "ltm-MAC-CE-SeparateTCI-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ltm_MAC_CE_JointTCI_r18 != nil, ie.Ltm_MAC_CE_SeparateTCI_r18 != nil}); err != nil {
				return err
			}

			if ie.Ltm_MAC_CE_JointTCI_r18 != nil {
				c := ie.Ltm_MAC_CE_JointTCI_r18
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtLtmMACCEJointTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberJointTCI_PerCell_r18, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberJointTCI_AcrossCells_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
			}

			if ie.Ltm_MAC_CE_SeparateTCI_r18 != nil {
				c := ie.Ltm_MAC_CE_SeparateTCI_r18
				if err := ex.EncodeEnumerated(c.Qcl_Resource_r18, bandNRExtLtmMACCESeparateTCIR18QclResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberDL_TCI_PerCell_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberUL_TCI_PerCell_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberDL_TCI_AcrossCells_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberUL_TCI_AcrossCells_r18, per.Constrained(1, 32)); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 18:
		if hasExtGroup18 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "eventA4BasedCondHandoverATG-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EventA4BasedCondHandoverATG_r18 != nil}); err != nil {
				return err
			}

			if ie.EventA4BasedCondHandoverATG_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EventA4BasedCondHandoverATG_r18, bandNRExtEventA4BasedCondHandoverATGR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 19:
		if hasExtGroup19 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "posSRS-TxFH-WithTimeWindowForRedCap-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PosSRS_TxFH_WithTimeWindowForRedCap_r18 != nil}); err != nil {
				return err
			}

			if ie.PosSRS_TxFH_WithTimeWindowForRedCap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_TxFH_WithTimeWindowForRedCap_r18, bandNRExtPosSRSTxFHWithTimeWindowForRedCapR18Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 20:
		if hasExtGroup20 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sbfd-Aware-r19", Optional: true},
					{Name: "ul-ResourceMutingCP-OFDM-r19", Optional: true},
					{Name: "ul-ResourceMutingDFTS-OFDM-r19", Optional: true},
					{Name: "ul-ResourceMutingCP-OFDM-CG-Type1-r19", Optional: true},
					{Name: "ul-ResourceMutingDFTS-OFDM-CG-Type1-r19", Optional: true},
					{Name: "l1-CLI-RSSI-MeasAndAperiodicReporting-r19", Optional: true},
					{Name: "l1-SRS-RSRP-MeasAndAperiodicReporting-r19", Optional: true},
					{Name: "od-SSB-NoAlwaysOn-RRC-r19", Optional: true},
					{Name: "od-SSB-NoAlwaysOn-RRC-MAC-CE-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-RRC-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-RRC-Diff-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-RRC-MAC-CE-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-RRC-MAC-CE-Diff-r19", Optional: true},
					{Name: "od-SSB-NoAlwaysOn-MAC-CE-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-MAC-CE-r19", Optional: true},
					{Name: "od-SSB-AlwaysOn-MAC-CE-Diff-r19", Optional: true},
					{Name: "ssb-BurstPeriodicityAdaptation-r19", Optional: true},
					{Name: "rach-AdaptationTimeDomain-r19", Optional: true},
					{Name: "ltm-BeamIndicationJointTCI-CSI-RS-r19", Optional: true},
					{Name: "ltm-MAC-CE-JointTCI-CSI-RS-r19", Optional: true},
					{Name: "ltm-BeamIndicationSeparateTCI-CSI-RS-r19", Optional: true},
					{Name: "ltm-MAC-CE-SeparateTCI-CSI-RS-r19", Optional: true},
					{Name: "ltm-CSI-RS-CSI-IM-Periodic-r19", Optional: true},
					{Name: "ltm-CSI-RS-CSI-IM-SP-r19", Optional: true},
					{Name: "ltm-WithoutCSI-IM-r19", Optional: true},
					{Name: "ssb-PeriodicityAddition-r19", Optional: true},
					{Name: "pdcch-RepetitionType0-r19", Optional: true},
					{Name: "pdcch-RepetitionTypeOthers-r19", Optional: true},
					{Name: "pdsch-RepetitionMsg4-r19", Optional: true},
					{Name: "ntn-Collision-RedCap-r19", Optional: true},
					{Name: "pusch-InterSlotOCC-r19", Optional: true},
					{Name: "triggeredHARQ-CodebookRetxDCI-1-3-Diff-r19", Optional: true},
					{Name: "unifiedJointTCI-MultiMAC-CE-DCI-1-3-Diff-r19", Optional: true},
					{Name: "unifiedSeparateTCI-MultiMAC-CE-IntraCell-Diff-r19", Optional: true},
					{Name: "posSRS-TxFH-RRC-ConnectedForNonRedCap-r19", Optional: true},
					{Name: "posSRS-TxFH-RRC-InactiveForNonRedCap-r19", Optional: true},
					{Name: "posSRS-TxFH-WithTimeWindowForNonRedCap-r19", Optional: true},
					{Name: "sr-TriggeredSSSG-Switching-r19", Optional: true},
					{Name: "pdcch-RepetitionType0-TN-r19", Optional: true},
					{Name: "pdcch-RepetitionTypeOthersTN-r19", Optional: true},
					{Name: "mpr-SingleCC-r19", Optional: true},
					{Name: "fastRx-BSF-MeasDelayReduction-r19", Optional: true},
					{Name: "fastSCellActivationEarlyMeas-r19", Optional: true},
					{Name: "od-SSB-AdditionalProcessingTime-r19", Optional: true},
					{Name: "aiml-BM-Case1-KnownRxBeam-r19", Optional: true},
					{Name: "aiml-BM-Case2-KnownRxBeam-r19", Optional: true},
					{Name: "ntn-PowerBoosting-ERedCap-r19", Optional: true},
					{Name: "earlyCSI-AcquisitionWithCSI-IM-r19", Optional: true},
					{Name: "earlyCSI-AcquisitionWithoutCSI-IM-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sbfd_Aware_r19 != nil, ie.Ul_ResourceMutingCP_OFDM_r19 != nil, ie.Ul_ResourceMutingDFTS_OFDM_r19 != nil, ie.Ul_ResourceMutingCP_OFDM_CG_Type1_r19 != nil, ie.Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19 != nil, ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19 != nil, ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19 != nil, ie.Od_SSB_NoAlwaysOn_RRC_r19 != nil, ie.Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19 != nil, ie.Od_SSB_AlwaysOn_RRC_r19 != nil, ie.Od_SSB_AlwaysOn_RRC_Diff_r19 != nil, ie.Od_SSB_AlwaysOn_RRC_MAC_CE_r19 != nil, ie.Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19 != nil, ie.Od_SSB_NoAlwaysOn_MAC_CE_r19 != nil, ie.Od_SSB_AlwaysOn_MAC_CE_r19 != nil, ie.Od_SSB_AlwaysOn_MAC_CE_Diff_r19 != nil, ie.Ssb_BurstPeriodicityAdaptation_r19 != nil, ie.Rach_AdaptationTimeDomain_r19 != nil, ie.Ltm_BeamIndicationJointTCI_CSI_RS_r19 != nil, ie.Ltm_MAC_CE_JointTCI_CSI_RS_r19 != nil, ie.Ltm_BeamIndicationSeparateTCI_CSI_RS_r19 != nil, ie.Ltm_MAC_CE_SeparateTCI_CSI_RS_r19 != nil, ie.Ltm_CSI_RS_CSI_IM_Periodic_r19 != nil, ie.Ltm_CSI_RS_CSI_IM_SP_r19 != nil, ie.Ltm_WithoutCSI_IM_r19 != nil, ie.Ssb_PeriodicityAddition_r19 != nil, ie.Pdcch_RepetitionType0_r19 != nil, ie.Pdcch_RepetitionTypeOthers_r19 != nil, ie.Pdsch_RepetitionMsg4_r19 != nil, ie.Ntn_Collision_RedCap_r19 != nil, ie.Pusch_InterSlotOCC_r19 != nil, ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19 != nil, ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19 != nil, ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19 != nil, ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19 != nil, ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19 != nil, ie.PosSRS_TxFH_WithTimeWindowForNonRedCap_r19 != nil, ie.Sr_TriggeredSSSG_Switching_r19 != nil, ie.Pdcch_RepetitionType0_TN_r19 != nil, ie.Pdcch_RepetitionTypeOthersTN_r19 != nil, ie.Mpr_SingleCC_r19 != nil, ie.FastRx_BSF_MeasDelayReduction_r19 != nil, ie.FastSCellActivationEarlyMeas_r19 != nil, ie.Od_SSB_AdditionalProcessingTime_r19 != nil, ie.Aiml_BM_Case1_KnownRxBeam_r19 != nil, ie.Aiml_BM_Case2_KnownRxBeam_r19 != nil, ie.Ntn_PowerBoosting_ERedCap_r19 != nil, ie.EarlyCSI_AcquisitionWithCSI_IM_r19 != nil, ie.EarlyCSI_AcquisitionWithoutCSI_IM_r19 != nil}); err != nil {
				return err
			}

			if ie.Sbfd_Aware_r19 != nil {
				c := ie.Sbfd_Aware_r19
				bandNRExtSbfdAwareR19Seq := ex.NewSequenceEncoder(bandNRExtSbfdAwareR19Constraints)
				if err := bandNRExtSbfdAwareR19Seq.EncodePreamble([]bool{c.Sbfd_Rx_r19 != nil, c.Sbfd_Tx_r19 != nil, c.Sbfd_OneRACH_r19 != nil, c.Sbfd_TwoRACH_r19 != nil, c.PreambleRepetitionOneRACH_r19 != nil, c.PreambleRepetitionTwoRACH_r19 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.NumOfPartialPRG_r19, bandNRExtSbfdAwareR19NumOfPartialPRGR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Timeline_r19, bandNRExtSbfdAwareR19TimelineR19Constraints); err != nil {
					return err
				}
				if c.Sbfd_Rx_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Sbfd_Rx_r19), bandNRExtSbfdAwareR19SbfdRxR19Constraints); err != nil {
						return err
					}
				}
				if c.Sbfd_Tx_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Sbfd_Tx_r19), bandNRExtSbfdAwareR19SbfdTxR19Constraints); err != nil {
						return err
					}
				}
				if c.Sbfd_OneRACH_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Sbfd_OneRACH_r19), bandNRExtSbfdAwareR19SbfdOneRACHR19Constraints); err != nil {
						return err
					}
				}
				if c.Sbfd_TwoRACH_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Sbfd_TwoRACH_r19), bandNRExtSbfdAwareR19SbfdTwoRACHR19Constraints); err != nil {
						return err
					}
				}
				if c.PreambleRepetitionOneRACH_r19 != nil {
					if err := ex.EncodeEnumerated((*c.PreambleRepetitionOneRACH_r19), bandNRExtSbfdAwareR19PreambleRepetitionOneRACHR19Constraints); err != nil {
						return err
					}
				}
				if c.PreambleRepetitionTwoRACH_r19 != nil {
					if err := ex.EncodeEnumerated((*c.PreambleRepetitionTwoRACH_r19), bandNRExtSbfdAwareR19PreambleRepetitionTwoRACHR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Ul_ResourceMutingCP_OFDM_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_ResourceMutingCP_OFDM_r19, bandNRExtUlResourceMutingCPOFDMR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_ResourceMutingDFTS_OFDM_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_ResourceMutingDFTS_OFDM_r19, bandNRExtUlResourceMutingDFTSOFDMR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_ResourceMutingCP_OFDM_CG_Type1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_ResourceMutingCP_OFDM_CG_Type1_r19, bandNRExtUlResourceMutingCPOFDMCGType1R19Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19, bandNRExtUlResourceMutingDFTSOFDMCGType1R19Constraints); err != nil {
					return err
				}
			}

			if ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19 != nil {
				c := ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19
				bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq := ex.NewSequenceEncoder(bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Constraints)
				if err := bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq.EncodePreamble([]bool{c.PeriodicReporting_r19 != nil, c.SemiPersistentMeasResource_r19 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfConfiguredMeasResourceAcrossCC_r19, bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfAperiodicReportSettingPerBWP_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfConfiguredMeasResourcePerCC_r19, per.Constrained(1, 32)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfSimultaneousMeasResourcePerCC_r19, per.Constrained(1, 32)); err != nil {
					return err
				}
				if c.PeriodicReporting_r19 != nil {
					if err := ex.EncodeEnumerated((*c.PeriodicReporting_r19), bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19PeriodicReportingR19Constraints); err != nil {
						return err
					}
				}
				if c.SemiPersistentMeasResource_r19 != nil {
					if err := ex.EncodeEnumerated((*c.SemiPersistentMeasResource_r19), bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19 != nil {
				c := ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19
				bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq := ex.NewSequenceEncoder(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Constraints)
				if err := bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq.EncodePreamble([]bool{c.Fdm_Reception_r19 != nil, c.SemiPersistentMeasResource_r19 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfConfiguredMeasResourceAcrossCC_r19, bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfConfiguredMeasResourcePerCC_r19, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfSimultaneousMeasResourcePerCC_r19, per.Constrained(1, 16)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfAperiodicReportSettingPerBWP_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumberOfMeasResourceAcrossCCWithinSlot_r19, bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfMeasResourceAcrossCCWithinSlotR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumberOfAperiodicMeasResourceAcrossCC_r19, per.Constrained(1, 4)); err != nil {
					return err
				}
				if c.Fdm_Reception_r19 != nil {
					if err := ex.EncodeEnumerated((*c.Fdm_Reception_r19), bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19FdmReceptionR19Constraints); err != nil {
						return err
					}
				}
				if c.SemiPersistentMeasResource_r19 != nil {
					if err := ex.EncodeEnumerated((*c.SemiPersistentMeasResource_r19), bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Od_SSB_NoAlwaysOn_RRC_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_NoAlwaysOn_RRC_r19, bandNRExtOdSSBNoAlwaysOnRRCR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19, bandNRExtOdSSBNoAlwaysOnRRCMACCER19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_RRC_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AlwaysOn_RRC_r19, bandNRExtOdSSBAlwaysOnRRCR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_RRC_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AlwaysOn_RRC_Diff_r19, bandNRExtOdSSBAlwaysOnRRCDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_RRC_MAC_CE_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AlwaysOn_RRC_MAC_CE_r19, bandNRExtOdSSBAlwaysOnRRCMACCER19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19, bandNRExtOdSSBAlwaysOnRRCMACCEDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_NoAlwaysOn_MAC_CE_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_NoAlwaysOn_MAC_CE_r19, bandNRExtOdSSBNoAlwaysOnMACCER19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_MAC_CE_r19 != nil {
				c := ie.Od_SSB_AlwaysOn_MAC_CE_r19
				if err := ex.EncodeEnumerated(c.TimeRelation_r19, bandNRExtOdSSBAlwaysOnMACCER19TimeRelationR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.DeactivationScheme_r19, bandNRExtOdSSBAlwaysOnMACCER19DeactivationSchemeR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AlwaysOn_MAC_CE_Diff_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AlwaysOn_MAC_CE_Diff_r19, bandNRExtOdSSBAlwaysOnMACCEDiffR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ssb_BurstPeriodicityAdaptation_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ssb_BurstPeriodicityAdaptation_r19, bandNRExtSsbBurstPeriodicityAdaptationR19Constraints); err != nil {
					return err
				}
			}

			if ie.Rach_AdaptationTimeDomain_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Rach_AdaptationTimeDomain_r19, bandNRExtRachAdaptationTimeDomainR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_BeamIndicationJointTCI_CSI_RS_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_BeamIndicationJointTCI_CSI_RS_r19, bandNRExtLtmBeamIndicationJointTCICSIRSR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MAC_CE_JointTCI_CSI_RS_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MAC_CE_JointTCI_CSI_RS_r19, bandNRExtLtmMACCEJointTCICSIRSR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_BeamIndicationSeparateTCI_CSI_RS_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_BeamIndicationSeparateTCI_CSI_RS_r19, bandNRExtLtmBeamIndicationSeparateTCICSIRSR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_MAC_CE_SeparateTCI_CSI_RS_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_MAC_CE_SeparateTCI_CSI_RS_r19, bandNRExtLtmMACCESeparateTCICSIRSR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_RS_CSI_IM_Periodic_r19 != nil {
				c := ie.Ltm_CSI_RS_CSI_IM_Periodic_r19
				if err := ex.EncodeInteger(c.MaxNumOfCSI_RS_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfCSI_RS_Ports_r19, bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfNZP_CSI_RS_Ports_r19, bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfNZPCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxRank_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumOfCSI_IM_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.Ltm_CSI_RS_CSI_IM_SP_r19 != nil {
				c := ie.Ltm_CSI_RS_CSI_IM_SP_r19
				if err := ex.EncodeInteger(c.MaxNumOfCSI_RS_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfCSI_RS_Ports_r19, bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfNZP_CSI_RS_Ports_r19, bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfNZPCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxRank_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumOfCSI_IM_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.Ltm_WithoutCSI_IM_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ltm_WithoutCSI_IM_r19, bandNRExtLtmWithoutCSIIMR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ssb_PeriodicityAddition_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ssb_PeriodicityAddition_r19, bandNRExtSsbPeriodicityAdditionR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_RepetitionType0_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_RepetitionType0_r19, bandNRExtPdcchRepetitionType0R19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_RepetitionTypeOthers_r19 != nil {
				if err := ex.EncodeInteger(*ie.Pdcch_RepetitionTypeOthers_r19, bandNRPdcchRepetitionTypeOthersR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_RepetitionMsg4_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_RepetitionMsg4_r19, bandNRExtPdschRepetitionMsg4R19Constraints); err != nil {
					return err
				}
			}

			if ie.Ntn_Collision_RedCap_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_Collision_RedCap_r19, bandNRExtNtnCollisionRedCapR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_InterSlotOCC_r19 != nil {
				c := ie.Pusch_InterSlotOCC_r19
				if err := ex.EncodeEnumerated(c.OccLength_r19, bandNRExtPuschInterSlotOCCR19OccLengthR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.SatelliteOrbit_r19, bandNRExtPuschInterSlotOCCR19SatelliteOrbitR19Constraints); err != nil {
					return err
				}
			}

			if ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19 != nil {
				c := ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19
				if err := ex.EncodeEnumerated(c.MinHARQ_Retx_Offset_r19, bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MinHARQRetxOffsetR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxHARQ_Retx_Offset_r19, bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MaxHARQRetxOffsetR19Constraints); err != nil {
					return err
				}
			}

			if ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19 != nil {
				c := ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19
				bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Constraints)
				if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Seq.EncodePreamble([]bool{c.MaxActivatedTCI_PerCC_r19 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MinBeamApplicationTime_r19.Choice), false, nil); err != nil {
						return err
					}
					switch c.MinBeamApplicationTime_r19.Choice {
					case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19:
						c := (*c.MinBeamApplicationTime_r19.Fr1_r19)
						bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Constraints)
						if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq.EncodePreamble([]bool{c.Scs_15kHz_r19 != nil, c.Scs_30kHz_r19 != nil, c.Scs_60kHz_r19 != nil}); err != nil {
							return err
						}
						if c.Scs_15kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_15kHz_r19), bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_30kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_30kHz_r19), bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_60kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r19), bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints); err != nil {
								return err
							}
						}
					case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19:
						c := (*c.MinBeamApplicationTime_r19.Fr2_r19)
						bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Constraints)
						if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq.EncodePreamble([]bool{c.Scs_60kHz_r19 != nil, c.Scs_120kHz_r19 != nil}); err != nil {
							return err
						}
						if c.Scs_60kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r19), bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_120kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_120kHz_r19), bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.MaxActivatedTCI_PerCC_r19 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedTCI_PerCC_r19), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
			}

			if ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19 != nil {
				c := ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19
				bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Constraints)
				if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq.EncodePreamble([]bool{c.MaxActivatedDL_TCI_PerCC_r19 != nil, c.MaxActivatedUL_TCI_PerCC_r19 != nil}); err != nil {
					return err
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MinBeamApplicationTime_r19.Choice), false, nil); err != nil {
						return err
					}
					switch c.MinBeamApplicationTime_r19.Choice {
					case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19:
						c := (*c.MinBeamApplicationTime_r19.Fr1_r19)
						bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Constraints)
						if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq.EncodePreamble([]bool{c.Scs_15kHz_r19 != nil, c.Scs_30kHz_r19 != nil, c.Scs_60kHz_r19 != nil}); err != nil {
							return err
						}
						if c.Scs_15kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_15kHz_r19), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_30kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_30kHz_r19), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_60kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r19), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints); err != nil {
								return err
							}
						}
					case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19:
						c := (*c.MinBeamApplicationTime_r19.Fr2_r19)
						bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq := ex.NewSequenceEncoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Constraints)
						if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq.EncodePreamble([]bool{c.Scs_60kHz_r19 != nil, c.Scs_120kHz_r19 != nil}); err != nil {
							return err
						}
						if c.Scs_60kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_60kHz_r19), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints); err != nil {
								return err
							}
						}
						if c.Scs_120kHz_r19 != nil {
							if err := ex.EncodeEnumerated((*c.Scs_120kHz_r19), bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints); err != nil {
								return err
							}
						}
					}
				}
				if c.MaxActivatedDL_TCI_PerCC_r19 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedDL_TCI_PerCC_r19), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
				if c.MaxActivatedUL_TCI_PerCC_r19 != nil {
					if err := ex.EncodeInteger((*c.MaxActivatedUL_TCI_PerCC_r19), per.Constrained(2, 8)); err != nil {
						return err
					}
				}
			}

			if ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19 != nil {
				if err := ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19 != nil {
				if err := ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PosSRS_TxFH_WithTimeWindowForNonRedCap_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PosSRS_TxFH_WithTimeWindowForNonRedCap_r19, bandNRExtPosSRSTxFHWithTimeWindowForNonRedCapR19Constraints); err != nil {
					return err
				}
			}

			if ie.Sr_TriggeredSSSG_Switching_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Sr_TriggeredSSSG_Switching_r19, bandNRExtSrTriggeredSSSGSwitchingR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_RepetitionType0_TN_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_RepetitionType0_TN_r19, bandNRExtPdcchRepetitionType0TNR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_RepetitionTypeOthersTN_r19 != nil {
				if err := ex.EncodeInteger(*ie.Pdcch_RepetitionTypeOthersTN_r19, bandNRPdcchRepetitionTypeOthersTNR19Constraints); err != nil {
					return err
				}
			}

			if ie.Mpr_SingleCC_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(bandNRExtMprSingleCCR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Mpr_SingleCC_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Mpr_SingleCC_r19).Choice {
				case BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_SingleValue_r19:
					if err := ex.EncodeEnumerated((*(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_SingleValue_r19), bandNRExtMprSingleCCR19MprSingleCCSingleValueR19Constraints); err != nil {
						return err
					}
				case BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_MultipleValue_r19:
					if err := ex.EncodeEnumerated((*(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_MultipleValue_r19), bandNRExtMprSingleCCR19MprSingleCCMultipleValueR19Constraints); err != nil {
						return err
					}
				}
			}

			if ie.FastRx_BSF_MeasDelayReduction_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.FastRx_BSF_MeasDelayReduction_r19, bandNRExtFastRxBSFMeasDelayReductionR19Constraints); err != nil {
					return err
				}
			}

			if ie.FastSCellActivationEarlyMeas_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.FastSCellActivationEarlyMeas_r19, bandNRExtFastSCellActivationEarlyMeasR19Constraints); err != nil {
					return err
				}
			}

			if ie.Od_SSB_AdditionalProcessingTime_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Od_SSB_AdditionalProcessingTime_r19, bandNRExtOdSSBAdditionalProcessingTimeR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case1_KnownRxBeam_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Aiml_BM_Case1_KnownRxBeam_r19, bandNRExtAimlBMCase1KnownRxBeamR19Constraints); err != nil {
					return err
				}
			}

			if ie.Aiml_BM_Case2_KnownRxBeam_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Aiml_BM_Case2_KnownRxBeam_r19, bandNRExtAimlBMCase2KnownRxBeamR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ntn_PowerBoosting_ERedCap_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_PowerBoosting_ERedCap_r19, bandNRExtNtnPowerBoostingERedCapR19Constraints); err != nil {
					return err
				}
			}

			if ie.EarlyCSI_AcquisitionWithCSI_IM_r19 != nil {
				c := ie.EarlyCSI_AcquisitionWithCSI_IM_r19
				if err := ex.EncodeInteger(c.MaxNumOfCSI_RS_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfCSI_RS_Ports_r19, bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MaxNumOfNZP_CSI_RS_Ports_r19, bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfNZPCSIRSPortsR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxRank_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.MaxNumOfCSI_IM_Resource_r19, per.Constrained(1, 8)); err != nil {
					return err
				}
			}

			if ie.EarlyCSI_AcquisitionWithoutCSI_IM_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.EarlyCSI_AcquisitionWithoutCSI_IM_r19, bandNRExtEarlyCSIAcquisitionWithoutCSIIMR19Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 21:
		if hasExtGroup21 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "enableTx-RxDuringMeasGap-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.EnableTx_RxDuringMeasGap_r19 != nil}); err != nil {
				return err
			}

			if ie.EnableTx_RxDuringMeasGap_r19 != nil {
				c := ie.EnableTx_RxDuringMeasGap_r19
				if err := ex.EncodeBitString(c.AdditionalDCI_r19, per.FixedSize(3)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.IndicationField_r19, bandNRExtEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.MinimumTimeOffset_r19, bandNRExtEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints); err != nil {
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

func (ie *BandNR) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandNRConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. bandNR: ref
	{
		if err := ie.BandNR.Decode(d); err != nil {
			return err
		}
	}

	// 4. modifiedMPR-Behaviour: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(bandNRModifiedMPRBehaviourConstraints)
			if err != nil {
				return err
			}
			ie.ModifiedMPR_Behaviour = &v
		}
	}

	// 5. mimo-ParametersPerBand: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mimo_ParametersPerBand = new(MIMO_ParametersPerBand)
			if err := ie.Mimo_ParametersPerBand.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. extendedCP: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(bandNRExtendedCPConstraints)
			if err != nil {
				return err
			}
			ie.ExtendedCP = &idx
		}
	}

	// 7. multipleTCI: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(bandNRMultipleTCIConstraints)
			if err != nil {
				return err
			}
			ie.MultipleTCI = &idx
		}
	}

	// 8. bwp-WithoutRestriction: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(bandNRBwpWithoutRestrictionConstraints)
			if err != nil {
				return err
			}
			ie.Bwp_WithoutRestriction = &idx
		}
	}

	// 9. bwp-SameNumerology: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(bandNRBwpSameNumerologyConstraints)
			if err != nil {
				return err
			}
			ie.Bwp_SameNumerology = &idx
		}
	}

	// 10. bwp-DiffNumerology: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(bandNRBwpDiffNumerologyConstraints)
			if err != nil {
				return err
			}
			ie.Bwp_DiffNumerology = &idx
		}
	}

	// 11. crossCarrierScheduling-SameSCS: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(bandNRCrossCarrierSchedulingSameSCSConstraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierScheduling_SameSCS = &idx
		}
	}

	// 12. pdsch-256QAM-FR2: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(bandNRPdsch256QAMFR2Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_256QAM_FR2 = &idx
		}
	}

	// 13. pusch-256QAM: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(bandNRPusch256QAMConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_256QAM = &idx
		}
	}

	// 14. ue-PowerClass: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(bandNRUePowerClassConstraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClass = &idx
		}
	}

	// 15. rateMatchingLTE-CRS: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(bandNRRateMatchingLTECRSConstraints)
			if err != nil {
				return err
			}
			ie.RateMatchingLTE_CRS = &idx
		}
	}

	// 16. channelBWs-DL: choice
	{
		if seq.IsComponentPresent(13) {
			ie.ChannelBWs_DL = &struct {
				Choice int
				Fr1    *struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}
				Fr2 *struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}
			}{}
			choiceDec := d.NewChoiceDecoder(bandNRChannelBWsDLConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBWs_DL).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BandNR_ChannelBWs_DL_Fr1:
				(*ie.ChannelBWs_DL).Fr1 = &struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_DL).Fr1)
				bandNRChannelBWsDLFr1Seq := d.NewSequenceDecoder(bandNRChannelBWsDLFr1Constraints)
				if err := bandNRChannelBWsDLFr1Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRChannelBWsDLFr1Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRChannelBWsDLFr1Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRChannelBWsDLFr1Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_ChannelBWs_DL_Fr2:
				(*ie.ChannelBWs_DL).Fr2 = &struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_DL).Fr2)
				bandNRChannelBWsDLFr2Seq := d.NewSequenceDecoder(bandNRChannelBWsDLFr2Constraints)
				if err := bandNRChannelBWsDLFr2Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRChannelBWsDLFr2Seq.IsComponentPresent(0) {
					c.Scs_60kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(3))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRChannelBWsDLFr2Seq.IsComponentPresent(1) {
					c.Scs_120kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(3))
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}
	}

	// 17. channelBWs-UL: choice
	{
		if seq.IsComponentPresent(14) {
			ie.ChannelBWs_UL = &struct {
				Choice int
				Fr1    *struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}
				Fr2 *struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}
			}{}
			choiceDec := d.NewChoiceDecoder(bandNRChannelBWsULConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBWs_UL).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case BandNR_ChannelBWs_UL_Fr1:
				(*ie.ChannelBWs_UL).Fr1 = &struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_UL).Fr1)
				bandNRChannelBWsULFr1Seq := d.NewSequenceDecoder(bandNRChannelBWsULFr1Constraints)
				if err := bandNRChannelBWsULFr1Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRChannelBWsULFr1Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRChannelBWsULFr1Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRChannelBWsULFr1Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_ChannelBWs_UL_Fr2:
				(*ie.ChannelBWs_UL).Fr2 = &struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_UL).Fr2)
				bandNRChannelBWsULFr2Seq := d.NewSequenceDecoder(bandNRChannelBWsULFr2Constraints)
				if err := bandNRChannelBWsULFr2Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRChannelBWsULFr2Seq.IsComponentPresent(0) {
					c.Scs_60kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(3))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRChannelBWsULFr2Seq.IsComponentPresent(1) {
					c.Scs_120kHz = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(3))
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
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
				{Name: "maxUplinkDutyCycle-PC2-FR1", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxUplinkDutyCyclePC2FR1Constraints)
			if err != nil {
				return err
			}
			ie.MaxUplinkDutyCycle_PC2_FR1 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pucch-SpatialRelInfoMAC-CE", Optional: true},
				{Name: "powerBoosting-pi2BPSK", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtPucchSpatialRelInfoMACCEConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_SpatialRelInfoMAC_CE = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtPowerBoostingPi2BPSKConstraints)
			if err != nil {
				return err
			}
			ie.PowerBoosting_Pi2BPSK = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxUplinkDutyCycle-FR2", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxUplinkDutyCycleFR2Constraints)
			if err != nil {
				return err
			}
			ie.MaxUplinkDutyCycle_FR2 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "channelBWs-DL-v1590", Optional: true},
				{Name: "channelBWs-UL-v1590", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ChannelBWs_DL_v1590 = &struct {
				Choice int
				Fr1    *struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}
				Fr2 *struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWsDLV1590Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBWs_DL_v1590).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBWs_DL_v1590_Fr1:
				(*ie.ChannelBWs_DL_v1590).Fr1 = &struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_DL_v1590).Fr1)
				bandNRExtChannelBWsDLV1590Fr1Seq := dx.NewSequenceDecoder(bandNRExtChannelBWsDLV1590Fr1Constraints)
				if err := bandNRExtChannelBWsDLV1590Fr1Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWsDLV1590Fr1Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWsDLV1590Fr1Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWsDLV1590Fr1Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBWs_DL_v1590_Fr2:
				(*ie.ChannelBWs_DL_v1590).Fr2 = &struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_DL_v1590).Fr2)
				bandNRExtChannelBWsDLV1590Fr2Seq := dx.NewSequenceDecoder(bandNRExtChannelBWsDLV1590Fr2Constraints)
				if err := bandNRExtChannelBWsDLV1590Fr2Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWsDLV1590Fr2Seq.IsComponentPresent(0) {
					c.Scs_60kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWsDLV1590Fr2Seq.IsComponentPresent(1) {
					c.Scs_120kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ChannelBWs_UL_v1590 = &struct {
				Choice int
				Fr1    *struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}
				Fr2 *struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWsULV1590Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBWs_UL_v1590).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBWs_UL_v1590_Fr1:
				(*ie.ChannelBWs_UL_v1590).Fr1 = &struct {
					Scs_15kHz *per.BitString
					Scs_30kHz *per.BitString
					Scs_60kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_UL_v1590).Fr1)
				bandNRExtChannelBWsULV1590Fr1Seq := dx.NewSequenceDecoder(bandNRExtChannelBWsULV1590Fr1Constraints)
				if err := bandNRExtChannelBWsULV1590Fr1Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWsULV1590Fr1Seq.IsComponentPresent(0) {
					c.Scs_15kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWsULV1590Fr1Seq.IsComponentPresent(1) {
					c.Scs_30kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWsULV1590Fr1Seq.IsComponentPresent(2) {
					c.Scs_60kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBWs_UL_v1590_Fr2:
				(*ie.ChannelBWs_UL_v1590).Fr2 = &struct {
					Scs_60kHz  *per.BitString
					Scs_120kHz *per.BitString
				}{}
				c := (*(*ie.ChannelBWs_UL_v1590).Fr2)
				bandNRExtChannelBWsULV1590Fr2Seq := dx.NewSequenceDecoder(bandNRExtChannelBWsULV1590Fr2Constraints)
				if err := bandNRExtChannelBWsULV1590Fr2Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWsULV1590Fr2Seq.IsComponentPresent(0) {
					c.Scs_60kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWsULV1590Fr2Seq.IsComponentPresent(1) {
					c.Scs_120kHz = new(per.BitString)
					v, err := dx.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
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
				{Name: "asymmetricBandwidthCombinationSet", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeBitString(bandNRAsymmetricBandwidthCombinationSetConstraints)
			if err != nil {
				return err
			}
			ie.AsymmetricBandwidthCombinationSet = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sharedSpectrumChAccessParamsPerBand-r16", Optional: true},
				{Name: "cancelOverlappingPUSCH-r16", Optional: true},
				{Name: "multipleRateMatchingEUTRA-CRS-r16", Optional: true},
				{Name: "overlapRateMatchingEUTRA-CRS-r16", Optional: true},
				{Name: "pdsch-MappingTypeB-Alt-r16", Optional: true},
				{Name: "oneSlotPeriodicTRS-r16", Optional: true},
				{Name: "olpc-SRS-Pos-r16", Optional: true},
				{Name: "spatialRelationsSRS-Pos-r16", Optional: true},
				{Name: "simulSRS-MIMO-TransWithinBand-r16", Optional: true},
				{Name: "channelBW-DL-IAB-r16", Optional: true},
				{Name: "channelBW-UL-IAB-r16", Optional: true},
				{Name: "rasterShift7dot5-IAB-r16", Optional: true},
				{Name: "ue-PowerClass-v1610", Optional: true},
				{Name: "condHandover-r16", Optional: true},
				{Name: "condHandoverFailure-r16", Optional: true},
				{Name: "condHandoverTwoTriggerEvents-r16", Optional: true},
				{Name: "condPSCellChange-r16", Optional: true},
				{Name: "condPSCellChangeTwoTriggerEvents-r16", Optional: true},
				{Name: "mpr-PowerBoost-FR2-r16", Optional: true},
				{Name: "activeConfiguredGrant-r16", Optional: true},
				{Name: "jointReleaseConfiguredGrantType2-r16", Optional: true},
				{Name: "sps-r16", Optional: true},
				{Name: "jointReleaseSPS-r16", Optional: true},
				{Name: "simulSRS-TransWithinBand-r16", Optional: true},
				{Name: "trs-AdditionalBandwidth-r16", Optional: true},
				{Name: "handoverIntraF-IAB-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SharedSpectrumChAccessParamsPerBand_r16 = new(SharedSpectrumChAccessParamsPerBand_r16)
			if err := ie.SharedSpectrumChAccessParamsPerBand_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtCancelOverlappingPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.CancelOverlappingPUSCH_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MultipleRateMatchingEUTRA_CRS_r16 = &struct {
				MaxNumberPatterns_r16            int64
				MaxNumberNon_OverlapPatterns_r16 int64
			}{}
			c := ie.MultipleRateMatchingEUTRA_CRS_r16
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 6))
				if err != nil {
					return err
				}
				c.MaxNumberPatterns_r16 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberNon_OverlapPatterns_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtOverlapRateMatchingEUTRACRSR16Constraints)
			if err != nil {
				return err
			}
			ie.OverlapRateMatchingEUTRA_CRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtPdschMappingTypeBAltR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_MappingTypeB_Alt_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtOneSlotPeriodicTRSR16Constraints)
			if err != nil {
				return err
			}
			ie.OneSlotPeriodicTRS_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Olpc_SRS_Pos_r16 = new(OLPC_SRS_Pos_r16)
			if err := ie.Olpc_SRS_Pos_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.SpatialRelationsSRS_Pos_r16 = new(SpatialRelationsSRS_Pos_r16)
			if err := ie.SpatialRelationsSRS_Pos_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandNRExtSimulSRSMIMOTransWithinBandR16Constraints)
			if err != nil {
				return err
			}
			ie.SimulSRS_MIMO_TransWithinBand_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			ie.ChannelBW_DL_IAB_r16 = &struct {
				Choice     int
				Fr1_100mhz *struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}
				Fr2_200mhz *struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWDLIABR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBW_DL_IAB_r16).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBW_DL_IAB_r16_Fr1_100mhz:
				(*ie.ChannelBW_DL_IAB_r16).Fr1_100mhz = &struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}{}
				c := (*(*ie.ChannelBW_DL_IAB_r16).Fr1_100mhz)
				bandNRExtChannelBWDLIABR16Fr1100mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWDLIABR16Fr1100mhzConstraints)
				if err := bandNRExtChannelBWDLIABR16Fr1100mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWDLIABR16Fr1100mhzSeq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLIABR16Fr1100mhzScs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWDLIABR16Fr1100mhzSeq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLIABR16Fr1100mhzScs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWDLIABR16Fr1100mhzSeq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLIABR16Fr1100mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBW_DL_IAB_r16_Fr2_200mhz:
				(*ie.ChannelBW_DL_IAB_r16).Fr2_200mhz = &struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}{}
				c := (*(*ie.ChannelBW_DL_IAB_r16).Fr2_200mhz)
				bandNRExtChannelBWDLIABR16Fr2200mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWDLIABR16Fr2200mhzConstraints)
				if err := bandNRExtChannelBWDLIABR16Fr2200mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWDLIABR16Fr2200mhzSeq.IsComponentPresent(0) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLIABR16Fr2200mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWDLIABR16Fr2200mhzSeq.IsComponentPresent(1) {
					c.Scs_120kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLIABR16Fr2200mhzScs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			ie.ChannelBW_UL_IAB_r16 = &struct {
				Choice     int
				Fr1_100mhz *struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}
				Fr2_200mhz *struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWULIABR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBW_UL_IAB_r16).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBW_UL_IAB_r16_Fr1_100mhz:
				(*ie.ChannelBW_UL_IAB_r16).Fr1_100mhz = &struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}{}
				c := (*(*ie.ChannelBW_UL_IAB_r16).Fr1_100mhz)
				bandNRExtChannelBWULIABR16Fr1100mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWULIABR16Fr1100mhzConstraints)
				if err := bandNRExtChannelBWULIABR16Fr1100mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWULIABR16Fr1100mhzSeq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULIABR16Fr1100mhzScs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWULIABR16Fr1100mhzSeq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULIABR16Fr1100mhzScs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWULIABR16Fr1100mhzSeq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULIABR16Fr1100mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBW_UL_IAB_r16_Fr2_200mhz:
				(*ie.ChannelBW_UL_IAB_r16).Fr2_200mhz = &struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}{}
				c := (*(*ie.ChannelBW_UL_IAB_r16).Fr2_200mhz)
				bandNRExtChannelBWULIABR16Fr2200mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWULIABR16Fr2200mhzConstraints)
				if err := bandNRExtChannelBWULIABR16Fr2200mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWULIABR16Fr2200mhzSeq.IsComponentPresent(0) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULIABR16Fr2200mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWULIABR16Fr2200mhzSeq.IsComponentPresent(1) {
					c.Scs_120kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULIABR16Fr2200mhzScs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(bandNRExtRasterShift7dot5IABR16Constraints)
			if err != nil {
				return err
			}
			ie.RasterShift7dot5_IAB_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(bandNRExtUePowerClassV1610Constraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClass_v1610 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(bandNRExtCondHandoverR16Constraints)
			if err != nil {
				return err
			}
			ie.CondHandover_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(bandNRExtCondHandoverFailureR16Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverFailure_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(bandNRExtCondHandoverTwoTriggerEventsR16Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverTwoTriggerEvents_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(bandNRExtCondPSCellChangeR16Constraints)
			if err != nil {
				return err
			}
			ie.CondPSCellChange_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(bandNRExtCondPSCellChangeTwoTriggerEventsR16Constraints)
			if err != nil {
				return err
			}
			ie.CondPSCellChangeTwoTriggerEvents_r16 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(bandNRExtMprPowerBoostFR2R16Constraints)
			if err != nil {
				return err
			}
			ie.Mpr_PowerBoost_FR2_r16 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			ie.ActiveConfiguredGrant_r16 = &struct {
				MaxNumberConfigsPerBWP_r16 int64
				MaxNumberConfigsAllCC_r16  int64
			}{}
			c := ie.ActiveConfiguredGrant_r16
			{
				v, err := dx.DecodeEnumerated(bandNRExtActiveConfiguredGrantR16MaxNumberConfigsPerBWPR16Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfigsPerBWP_r16 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 32))
				if err != nil {
					return err
				}
				c.MaxNumberConfigsAllCC_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(bandNRExtJointReleaseConfiguredGrantType2R16Constraints)
			if err != nil {
				return err
			}
			ie.JointReleaseConfiguredGrantType2_r16 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			ie.Sps_r16 = &struct {
				MaxNumberConfigsPerBWP_r16 int64
				MaxNumberConfigsAllCC_r16  int64
			}{}
			c := ie.Sps_r16
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberConfigsPerBWP_r16 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 32))
				if err != nil {
					return err
				}
				c.MaxNumberConfigsAllCC_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(bandNRExtJointReleaseSPSR16Constraints)
			if err != nil {
				return err
			}
			ie.JointReleaseSPS_r16 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(bandNRExtSimulSRSTransWithinBandR16Constraints)
			if err != nil {
				return err
			}
			ie.SimulSRS_TransWithinBand_r16 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(bandNRExtTrsAdditionalBandwidthR16Constraints)
			if err != nil {
				return err
			}
			ie.Trs_AdditionalBandwidth_r16 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(bandNRExtHandoverIntraFIABR16Constraints)
			if err != nil {
				return err
			}
			ie.HandoverIntraF_IAB_r16 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "simulTX-SRS-AntSwitchingIntraBandUL-CA-r16", Optional: true},
				{Name: "sharedSpectrumChAccessParamsPerBand-v1630", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
			if err := ie.SimulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SharedSpectrumChAccessParamsPerBand_v1630 = new(SharedSpectrumChAccessParamsPerBand_v1630)
			if err := ie.SharedSpectrumChAccessParamsPerBand_v1630.Decode(dx); err != nil {
				return err
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
				{Name: "handoverUTRA-FDD-r16", Optional: true},
				{Name: "enhancedUL-TransientPeriod-r16", Optional: true},
				{Name: "sharedSpectrumChAccessParamsPerBand-v1640", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtHandoverUTRAFDDR16Constraints)
			if err != nil {
				return err
			}
			ie.HandoverUTRA_FDD_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtEnhancedULTransientPeriodR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedUL_TransientPeriod_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.SharedSpectrumChAccessParamsPerBand_v1640 = new(SharedSpectrumChAccessParamsPerBand_v1640)
			if err := ie.SharedSpectrumChAccessParamsPerBand_v1640.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 8:
	if seq.IsExtensionPresent(8) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "type1-PUSCH-RepetitionMultiSlots-v1650", Optional: true},
				{Name: "type2-PUSCH-RepetitionMultiSlots-v1650", Optional: true},
				{Name: "pusch-RepetitionMultiSlots-v1650", Optional: true},
				{Name: "configuredUL-GrantType1-v1650", Optional: true},
				{Name: "configuredUL-GrantType2-v1650", Optional: true},
				{Name: "sharedSpectrumChAccessParamsPerBand-v1650", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtType1PUSCHRepetitionMultiSlotsV1650Constraints)
			if err != nil {
				return err
			}
			ie.Type1_PUSCH_RepetitionMultiSlots_v1650 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtType2PUSCHRepetitionMultiSlotsV1650Constraints)
			if err != nil {
				return err
			}
			ie.Type2_PUSCH_RepetitionMultiSlots_v1650 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtPuschRepetitionMultiSlotsV1650Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepetitionMultiSlots_v1650 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtConfiguredULGrantType1V1650Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType1_v1650 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtConfiguredULGrantType2V1650Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_GrantType2_v1650 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.SharedSpectrumChAccessParamsPerBand_v1650 = new(SharedSpectrumChAccessParamsPerBand_v1650)
			if err := ie.SharedSpectrumChAccessParamsPerBand_v1650.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enhancedSkipUplinkTxConfigured-v1660", Optional: true},
				{Name: "enhancedSkipUplinkTxDynamic-v1660", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtEnhancedSkipUplinkTxConfiguredV1660Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxConfigured_v1660 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtEnhancedSkipUplinkTxDynamicV1660Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedSkipUplinkTxDynamic_v1660 = &v
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxUplinkDutyCycle-PC1dot5-MPE-FR1-r16", Optional: true},
				{Name: "txDiversity-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxUplinkDutyCyclePC1dot5MPEFR1R16Constraints)
			if err != nil {
				return err
			}
			ie.MaxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtTxDiversityR16Constraints)
			if err != nil {
				return err
			}
			ie.TxDiversity_r16 = &v
		}
	}

	// Extension group 11:
	if seq.IsExtensionPresent(11) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdsch-1024QAM-FR1-r17", Optional: true},
				{Name: "ue-PowerClass-v1700", Optional: true},
				{Name: "fr2-2-AccessParamsPerBand-r17", Optional: true},
				{Name: "rlm-Relaxation-r17", Optional: true},
				{Name: "bfd-Relaxation-r17", Optional: true},
				{Name: "cg-SDT-r17", Optional: true},
				{Name: "locationBasedCondHandover-r17", Optional: true},
				{Name: "timeBasedCondHandover-r17", Optional: true},
				{Name: "eventA4BasedCondHandover-r17", Optional: true},
				{Name: "mn-InitiatedCondPSCellChangeNRDC-r17", Optional: true},
				{Name: "sn-InitiatedCondPSCellChangeNRDC-r17", Optional: true},
				{Name: "pdcch-SkippingWithoutSSSG-r17", Optional: true},
				{Name: "sssg-Switching-1BitInd-r17", Optional: true},
				{Name: "sssg-Switching-2BitInd-r17", Optional: true},
				{Name: "pdcch-SkippingWithSSSG-r17", Optional: true},
				{Name: "searchSpaceSetGrp-switchCap2-r17", Optional: true},
				{Name: "uplinkPreCompensation-r17", Optional: true},
				{Name: "uplink-TA-Reporting-r17", Optional: true},
				{Name: "max-HARQ-ProcessNumber-r17", Optional: true},
				{Name: "type2-HARQ-Codebook-r17", Optional: true},
				{Name: "type1-HARQ-Codebook-r17", Optional: true},
				{Name: "type3-HARQ-Codebook-r17", Optional: true},
				{Name: "ue-specific-K-Offset-r17", Optional: true},
				{Name: "multiPDSCH-SingleDCI-FR2-1-SCS-120kHz-r17", Optional: true},
				{Name: "multiPUSCH-SingleDCI-FR2-1-SCS-120kHz-r17", Optional: true},
				{Name: "parallelPRS-MeasRRC-Inactive-r17", Optional: true},
				{Name: "nr-UE-TxTEG-ID-MaxSupport-r17", Optional: true},
				{Name: "prs-ProcessingRRC-Inactive-r17", Optional: true},
				{Name: "prs-ProcessingWindowType1A-r17", Optional: true},
				{Name: "prs-ProcessingWindowType1B-r17", Optional: true},
				{Name: "prs-ProcessingWindowType2-r17", Optional: true},
				{Name: "srs-AllPosResourcesRRC-Inactive-r17", Optional: true},
				{Name: "olpc-SRS-PosRRC-Inactive-r17", Optional: true},
				{Name: "spatialRelationsSRS-PosRRC-Inactive-r17", Optional: true},
				{Name: "maxNumberPUSCH-TypeA-Repetition-r17", Optional: true},
				{Name: "puschTypeA-RepetitionsAvailSlot-r17", Optional: true},
				{Name: "tb-ProcessingMultiSlotPUSCH-r17", Optional: true},
				{Name: "tb-ProcessingRepMultiSlotPUSCH-r17", Optional: true},
				{Name: "maxDurationDMRS-Bundling-r17", Optional: true},
				{Name: "pusch-RepetitionMsg3-r17", Optional: true},
				{Name: "sharedSpectrumChAccessParamsPerBand-v1710", Optional: true},
				{Name: "parallelMeasurementWithoutRestriction-r17", Optional: true},
				{Name: "maxNumber-NGSO-SatellitesWithinOneSMTC-r17", Optional: true},
				{Name: "k1-RangeExtension-r17", Optional: true},
				{Name: "aperiodicCSI-RS-FastScellActivation-r17", Optional: true},
				{Name: "aperiodicCSI-RS-AdditionalBandwidth-r17", Optional: true},
				{Name: "bwp-WithoutCD-SSB-OrNCD-SSB-RedCap-r17", Optional: true},
				{Name: "halfDuplexFDD-TypeA-RedCap-r17", Optional: true},
				{Name: "posSRS-RRC-Inactive-OutsideInitialUL-BWP-r17", Optional: true},
				{Name: "channelBWs-DL-SCS-480kHz-FR2-2-r17", Optional: true},
				{Name: "channelBWs-UL-SCS-480kHz-FR2-2-r17", Optional: true},
				{Name: "channelBWs-DL-SCS-960kHz-FR2-2-r17", Optional: true},
				{Name: "channelBWs-UL-SCS-960kHz-FR2-2-r17", Optional: true},
				{Name: "ul-GapFR2-r17", Optional: true},
				{Name: "oneShotHARQ-feedbackTriggeredByDCI-1-2-r17", Optional: true},
				{Name: "oneShotHARQ-feedbackPhy-Priority-r17", Optional: true},
				{Name: "enhancedType3-HARQ-CodebookFeedback-r17", Optional: true},
				{Name: "triggeredHARQ-CodebookRetx-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtPdsch1024QAMFR1R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_1024QAM_FR1_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtUePowerClassV1700Constraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClass_v1700 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Fr2_2_AccessParamsPerBand_r17 = new(FR2_2_AccessParamsPerBand_r17)
			if err := ie.Fr2_2_AccessParamsPerBand_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtRlmRelaxationR17Constraints)
			if err != nil {
				return err
			}
			ie.Rlm_Relaxation_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtBfdRelaxationR17Constraints)
			if err != nil {
				return err
			}
			ie.Bfd_Relaxation_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtCgSDTR17Constraints)
			if err != nil {
				return err
			}
			ie.Cg_SDT_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(bandNRExtLocationBasedCondHandoverR17Constraints)
			if err != nil {
				return err
			}
			ie.LocationBasedCondHandover_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtTimeBasedCondHandoverR17Constraints)
			if err != nil {
				return err
			}
			ie.TimeBasedCondHandover_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandNRExtEventA4BasedCondHandoverR17Constraints)
			if err != nil {
				return err
			}
			ie.EventA4BasedCondHandover_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(bandNRExtMnInitiatedCondPSCellChangeNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Mn_InitiatedCondPSCellChangeNRDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(bandNRExtSnInitiatedCondPSCellChangeNRDCR17Constraints)
			if err != nil {
				return err
			}
			ie.Sn_InitiatedCondPSCellChangeNRDC_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(bandNRExtPdcchSkippingWithoutSSSGR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_SkippingWithoutSSSG_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(bandNRExtSssgSwitching1BitIndR17Constraints)
			if err != nil {
				return err
			}
			ie.Sssg_Switching_1BitInd_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(bandNRExtSssgSwitching2BitIndR17Constraints)
			if err != nil {
				return err
			}
			ie.Sssg_Switching_2BitInd_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(bandNRExtPdcchSkippingWithSSSGR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_SkippingWithSSSG_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(bandNRExtSearchSpaceSetGrpSwitchCap2R17Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSetGrp_SwitchCap2_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(bandNRExtUplinkPreCompensationR17Constraints)
			if err != nil {
				return err
			}
			ie.UplinkPreCompensation_r17 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(bandNRExtUplinkTAReportingR17Constraints)
			if err != nil {
				return err
			}
			ie.Uplink_TA_Reporting_r17 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxHARQProcessNumberR17Constraints)
			if err != nil {
				return err
			}
			ie.Max_HARQ_ProcessNumber_r17 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(bandNRExtType2HARQCodebookR17Constraints)
			if err != nil {
				return err
			}
			ie.Type2_HARQ_Codebook_r17 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(bandNRExtType1HARQCodebookR17Constraints)
			if err != nil {
				return err
			}
			ie.Type1_HARQ_Codebook_r17 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(bandNRExtType3HARQCodebookR17Constraints)
			if err != nil {
				return err
			}
			ie.Type3_HARQ_Codebook_r17 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(bandNRExtUeSpecificKOffsetR17Constraints)
			if err != nil {
				return err
			}
			ie.Ue_Specific_K_Offset_r17 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(bandNRExtMultiPDSCHSingleDCIFR21SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(bandNRExtMultiPUSCHSingleDCIFR21SCS120kHzR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(bandNRExtParallelPRSMeasRRCInactiveR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelPRS_MeasRRC_Inactive_r17 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeEnumerated(bandNRExtNrUETxTEGIDMaxSupportR17Constraints)
			if err != nil {
				return err
			}
			ie.Nr_UE_TxTEG_ID_MaxSupport_r17 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			v, err := dx.DecodeEnumerated(bandNRExtPrsProcessingRRCInactiveR17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_ProcessingRRC_Inactive_r17 = &v
		}

		if groupSeq.IsComponentPresent(28) {
			v, err := dx.DecodeEnumerated(bandNRExtPrsProcessingWindowType1AR17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_ProcessingWindowType1A_r17 = &v
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(bandNRExtPrsProcessingWindowType1BR17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_ProcessingWindowType1B_r17 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			v, err := dx.DecodeEnumerated(bandNRExtPrsProcessingWindowType2R17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_ProcessingWindowType2_r17 = &v
		}

		if groupSeq.IsComponentPresent(31) {
			ie.Srs_AllPosResourcesRRC_Inactive_r17 = new(SRS_AllPosResourcesRRC_Inactive_r17)
			if err := ie.Srs_AllPosResourcesRRC_Inactive_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(32) {
			ie.Olpc_SRS_PosRRC_Inactive_r17 = new(OLPC_SRS_Pos_r16)
			if err := ie.Olpc_SRS_PosRRC_Inactive_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(33) {
			ie.SpatialRelationsSRS_PosRRC_Inactive_r17 = new(SpatialRelationsSRS_Pos_r16)
			if err := ie.SpatialRelationsSRS_PosRRC_Inactive_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(34) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxNumberPUSCHTypeARepetitionR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberPUSCH_TypeA_Repetition_r17 = &v
		}

		if groupSeq.IsComponentPresent(35) {
			v, err := dx.DecodeEnumerated(bandNRExtPuschTypeARepetitionsAvailSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.PuschTypeA_RepetitionsAvailSlot_r17 = &v
		}

		if groupSeq.IsComponentPresent(36) {
			v, err := dx.DecodeEnumerated(bandNRExtTbProcessingMultiSlotPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Tb_ProcessingMultiSlotPUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(37) {
			v, err := dx.DecodeEnumerated(bandNRExtTbProcessingRepMultiSlotPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Tb_ProcessingRepMultiSlotPUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(38) {
			ie.MaxDurationDMRS_Bundling_r17 = &struct {
				Fdd_r17 *int64
				Tdd_r17 *int64
			}{}
			c := ie.MaxDurationDMRS_Bundling_r17
			bandNRExtMaxDurationDMRSBundlingR17Seq := dx.NewSequenceDecoder(bandNRExtMaxDurationDMRSBundlingR17Constraints)
			if err := bandNRExtMaxDurationDMRSBundlingR17Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandNRExtMaxDurationDMRSBundlingR17Seq.IsComponentPresent(0) {
				c.Fdd_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMaxDurationDMRSBundlingR17FddR17Constraints)
				if err != nil {
					return err
				}
				(*c.Fdd_r17) = v
			}
			if bandNRExtMaxDurationDMRSBundlingR17Seq.IsComponentPresent(1) {
				c.Tdd_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMaxDurationDMRSBundlingR17TddR17Constraints)
				if err != nil {
					return err
				}
				(*c.Tdd_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(39) {
			v, err := dx.DecodeEnumerated(bandNRExtPuschRepetitionMsg3R17Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepetitionMsg3_r17 = &v
		}

		if groupSeq.IsComponentPresent(40) {
			ie.SharedSpectrumChAccessParamsPerBand_v1710 = new(SharedSpectrumChAccessParamsPerBand_v1710)
			if err := ie.SharedSpectrumChAccessParamsPerBand_v1710.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(41) {
			v, err := dx.DecodeEnumerated(bandNRExtParallelMeasurementWithoutRestrictionR17Constraints)
			if err != nil {
				return err
			}
			ie.ParallelMeasurementWithoutRestriction_r17 = &v
		}

		if groupSeq.IsComponentPresent(42) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxNumberNGSOSatellitesWithinOneSMTCR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumber_NGSO_SatellitesWithinOneSMTC_r17 = &v
		}

		if groupSeq.IsComponentPresent(43) {
			v, err := dx.DecodeEnumerated(bandNRExtK1RangeExtensionR17Constraints)
			if err != nil {
				return err
			}
			ie.K1_RangeExtension_r17 = &v
		}

		if groupSeq.IsComponentPresent(44) {
			ie.AperiodicCSI_RS_FastScellActivation_r17 = &struct {
				MaxNumberAperiodicCSI_RS_PerCC_r17     int64
				MaxNumberAperiodicCSI_RS_AcrossCCs_r17 int64
			}{}
			c := ie.AperiodicCSI_RS_FastScellActivation_r17
			{
				v, err := dx.DecodeEnumerated(bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSPerCCR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicCSI_RS_PerCC_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtAperiodicCSIRSFastScellActivationR17MaxNumberAperiodicCSIRSAcrossCCsR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberAperiodicCSI_RS_AcrossCCs_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(bandNRExtAperiodicCSIRSAdditionalBandwidthR17Constraints)
			if err != nil {
				return err
			}
			ie.AperiodicCSI_RS_AdditionalBandwidth_r17 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(bandNRExtBwpWithoutCDSSBOrNCDSSBRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.Bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			v, err := dx.DecodeEnumerated(bandNRExtHalfDuplexFDDTypeARedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.HalfDuplexFDD_TypeA_RedCap_r17 = &v
		}

		if groupSeq.IsComponentPresent(48) {
			ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17)
			if err := ie.PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(49) {
			v, err := dx.DecodeBitString(bandNRChannelBWsDLSCS480kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_DL_SCS_480kHz_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(50) {
			v, err := dx.DecodeBitString(bandNRChannelBWsULSCS480kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_UL_SCS_480kHz_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(51) {
			v, err := dx.DecodeBitString(bandNRChannelBWsDLSCS960kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_DL_SCS_960kHz_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(52) {
			v, err := dx.DecodeBitString(bandNRChannelBWsULSCS960kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_UL_SCS_960kHz_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(53) {
			v, err := dx.DecodeEnumerated(bandNRExtUlGapFR2R17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_GapFR2_r17 = &v
		}

		if groupSeq.IsComponentPresent(54) {
			v, err := dx.DecodeEnumerated(bandNRExtOneShotHARQFeedbackTriggeredByDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.OneShotHARQ_FeedbackTriggeredByDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(55) {
			v, err := dx.DecodeEnumerated(bandNRExtOneShotHARQFeedbackPhyPriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.OneShotHARQ_FeedbackPhy_Priority_r17 = &v
		}

		if groupSeq.IsComponentPresent(56) {
			ie.EnhancedType3_HARQ_CodebookFeedback_r17 = &struct {
				EnhancedType3_HARQ_Codebooks_r17 int64
				MaxNumberPUCCH_Transmissions_r17 int64
			}{}
			c := ie.EnhancedType3_HARQ_CodebookFeedback_r17
			{
				v, err := dx.DecodeEnumerated(bandNRExtEnhancedType3HARQCodebookFeedbackR17EnhancedType3HARQCodebooksR17Constraints)
				if err != nil {
					return err
				}
				c.EnhancedType3_HARQ_Codebooks_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtEnhancedType3HARQCodebookFeedbackR17MaxNumberPUCCHTransmissionsR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberPUCCH_Transmissions_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(57) {
			ie.TriggeredHARQ_CodebookRetx_r17 = &struct {
				MinHARQ_Retx_Offset_r17 int64
				MaxHARQ_Retx_Offset_r17 int64
			}{}
			c := ie.TriggeredHARQ_CodebookRetx_r17
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxR17MinHARQRetxOffsetR17Constraints)
				if err != nil {
					return err
				}
				c.MinHARQ_Retx_Offset_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxR17MaxHARQRetxOffsetR17Constraints)
				if err != nil {
					return err
				}
				c.MaxHARQ_Retx_Offset_r17 = v
			}
		}
	}

	// Extension group 12:
	if seq.IsExtensionPresent(12) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ue-OneShotUL-TimingAdj-r17", Optional: true},
				{Name: "pucch-Repetition-F0-2-r17", Optional: true},
				{Name: "cqi-4-BitsSubbandNTN-SharedSpectrumChAccess-r17", Optional: true},
				{Name: "mux-HARQ-ACK-DiffPriorities-r17", Optional: true},
				{Name: "ta-BasedPDC-NTN-SharedSpectrumChAccess-r17", Optional: true},
				{Name: "ack-NACK-FeedbackForMulticastWithDCI-Enabler-r17", Optional: true},
				{Name: "maxNumberG-RNTI-r17", Optional: true},
				{Name: "dynamicMulticastDCI-Format4-2-r17", Optional: true},
				{Name: "maxModulationOrderForMulticast-r17", Optional: true},
				{Name: "dynamicSlotRepetitionMulticastTN-NonSharedSpectrumChAccess-r17", Optional: true},
				{Name: "dynamicSlotRepetitionMulticastNTN-SharedSpectrumChAccess-r17", Optional: true},
				{Name: "nack-OnlyFeedbackForMulticastWithDCI-Enabler-r17", Optional: true},
				{Name: "ack-NACK-FeedbackForSPS-MulticastWithDCI-Enabler-r17", Optional: true},
				{Name: "maxNumberG-CS-RNTI-r17", Optional: true},
				{Name: "re-LevelRateMatchingForMulticast-r17", Optional: true},
				{Name: "pdsch-1024QAM-2MIMO-FR1-r17", Optional: true},
				{Name: "prs-MeasurementWithoutMG-r17", Optional: true},
				{Name: "maxNumber-NGSO-SatellitesPerCarrier-r17", Optional: true},
				{Name: "prs-ProcessingCapabilityOutsideMGinPPW-r17", Optional: true},
				{Name: "srs-SemiPersistent-PosResourcesRRC-Inactive-r17", Optional: true},
				{Name: "channelBWs-DL-SCS-120kHz-FR2-2-r17", Optional: true},
				{Name: "channelBWs-UL-SCS-120kHz-FR2-2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtUeOneShotULTimingAdjR17Constraints)
			if err != nil {
				return err
			}
			ie.Ue_OneShotUL_TimingAdj_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtPucchRepetitionF02R17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_Repetition_F0_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtCqi4BitsSubbandNTNSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtMuxHARQACKDiffPrioritiesR17Constraints)
			if err != nil {
				return err
			}
			ie.Mux_HARQ_ACK_DiffPriorities_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtTaBasedPDCNTNSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtAckNACKFeedbackForMulticastWithDCIEnablerR17Constraints)
			if err != nil {
				return err
			}
			ie.Ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeInteger(bandNRMaxNumberGRNTIR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberG_RNTI_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtDynamicMulticastDCIFormat42R17Constraints)
			if err != nil {
				return err
			}
			ie.DynamicMulticastDCI_Format4_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.MaxModulationOrderForMulticast_r17 = &struct {
				Choice  int
				Fr1_r17 *int64
				Fr2_r17 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtMaxModulationOrderForMulticastR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxModulationOrderForMulticast_r17).Choice = int(index)
			switch index {
			case BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr1_r17:
				(*ie.MaxModulationOrderForMulticast_r17).Fr1_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMaxModulationOrderForMulticastR17Fr1R17Constraints)
				if err != nil {
					return err
				}
				(*(*ie.MaxModulationOrderForMulticast_r17).Fr1_r17) = v
			case BandNR_Ext_MaxModulationOrderForMulticast_r17_Fr2_r17:
				(*ie.MaxModulationOrderForMulticast_r17).Fr2_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMaxModulationOrderForMulticastR17Fr2R17Constraints)
				if err != nil {
					return err
				}
				(*(*ie.MaxModulationOrderForMulticast_r17).Fr2_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(bandNRExtDynamicSlotRepetitionMulticastTNNonSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.DynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(bandNRExtDynamicSlotRepetitionMulticastNTNSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.DynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(bandNRExtNackOnlyFeedbackForMulticastWithDCIEnablerR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(bandNRExtAckNACKFeedbackForSPSMulticastWithDCIEnablerR17Constraints)
			if err != nil {
				return err
			}
			ie.Ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeInteger(bandNRMaxNumberGCSRNTIR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumberG_CS_RNTI_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(bandNRExtReLevelRateMatchingForMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Re_LevelRateMatchingForMulticast_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(bandNRExtPdsch1024QAM2MIMOFR1R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_1024QAM_2MIMO_FR1_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(bandNRExtPrsMeasurementWithoutMGR17Constraints)
			if err != nil {
				return err
			}
			ie.Prs_MeasurementWithoutMG_r17 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeInteger(bandNRMaxNumberNGSOSatellitesPerCarrierR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxNumber_NGSO_SatellitesPerCarrier_r17 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			seqOf := dx.NewSequenceOfDecoder(bandNRExtPrsProcessingCapabilityOutsideMGinPPWR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17 = make([]PRS_ProcessingCapabilityOutsideMGinPPWperType_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Prs_ProcessingCapabilityOutsideMGinPPW_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(19) {
			ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17 = &struct {
				MaxNumOfSemiPersistentSRSposResources_r17        int64
				MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 int64
			}{}
			c := ie.Srs_SemiPersistent_PosResourcesRRC_Inactive_r17
			{
				v, err := dx.DecodeEnumerated(bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfSemiPersistentSRSposResources_r17 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtSrsSemiPersistentPosResourcesRRCInactiveR17MaxNumOfSemiPersistentSRSposResourcesPerSlotR17Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfSemiPersistentSRSposResourcesPerSlot_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeBitString(bandNRChannelBWsDLSCS120kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_DL_SCS_120kHz_FR2_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeBitString(bandNRChannelBWsULSCS120kHzFR22R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelBWs_UL_SCS_120kHz_FR2_2_r17 = &v
		}
	}

	// Extension group 13:
	if seq.IsExtensionPresent(13) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dmrs-BundlingPUSCH-RepTypeA-r17", Optional: true},
				{Name: "dmrs-BundlingPUSCH-RepTypeB-r17", Optional: true},
				{Name: "dmrs-BundlingPUSCH-multiSlot-r17", Optional: true},
				{Name: "dmrs-BundlingPUCCH-Rep-r17", Optional: true},
				{Name: "interSlotFreqHopInterSlotBundlingPUSCH-r17", Optional: true},
				{Name: "interSlotFreqHopPUCCH-r17", Optional: true},
				{Name: "dmrs-BundlingRestart-r17", Optional: true},
				{Name: "dmrs-BundlingNonBackToBackTX-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingPUSCHRepTypeAR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_RepTypeA_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingPUSCHRepTypeBR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_RepTypeB_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingPUSCHMultiSlotR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUSCH_MultiSlot_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingPUCCHRepR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingPUCCH_Rep_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtInterSlotFreqHopInterSlotBundlingPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.InterSlotFreqHopInterSlotBundlingPUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtInterSlotFreqHopPUCCHR17Constraints)
			if err != nil {
				return err
			}
			ie.InterSlotFreqHopPUCCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingRestartR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingRestart_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtDmrsBundlingNonBackToBackTXR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_BundlingNonBackToBackTX_r17 = &v
		}
	}

	// Extension group 14:
	if seq.IsExtensionPresent(14) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "maxDynamicSlotRepetitionForSPS-Multicast-r17", Optional: true},
				{Name: "nack-OnlyFeedbackForSPS-MulticastWithDCI-Enabler-r17", Optional: true},
				{Name: "sps-MulticastDCI-Format4-2-r17", Optional: true},
				{Name: "sps-MulticastMultiConfig-r17", Optional: true},
				{Name: "priorityIndicatorInDCI-Multicast-r17", Optional: true},
				{Name: "priorityIndicatorInDCI-SPS-Multicast-r17", Optional: true},
				{Name: "twoHARQ-ACK-CodebookForUnicastAndMulticast-r17", Optional: true},
				{Name: "multiPUCCH-HARQ-ACK-ForMulticastUnicast-r17", Optional: true},
				{Name: "releaseSPS-MulticastWithCS-RNTI-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtMaxDynamicSlotRepetitionForSPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.MaxDynamicSlotRepetitionForSPS_Multicast_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtNackOnlyFeedbackForSPSMulticastWithDCIEnablerR17Constraints)
			if err != nil {
				return err
			}
			ie.Nack_OnlyFeedbackForSPS_MulticastWithDCI_Enabler_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtSpsMulticastDCIFormat42R17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_MulticastDCI_Format4_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeInteger(bandNRSpsMulticastMultiConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Sps_MulticastMultiConfig_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtPriorityIndicatorInDCIMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorInDCI_Multicast_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtPriorityIndicatorInDCISPSMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorInDCI_SPS_Multicast_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(bandNRExtTwoHARQACKCodebookForUnicastAndMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.TwoHARQ_ACK_CodebookForUnicastAndMulticast_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtMultiPUCCHHARQACKForMulticastUnicastR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandNRExtReleaseSPSMulticastWithCSRNTIR17Constraints)
			if err != nil {
				return err
			}
			ie.ReleaseSPS_MulticastWithCS_RNTI_r17 = &v
		}
	}

	// Extension group 15:
	if seq.IsExtensionPresent(15) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "posUE-TA-AutoAdjustment-r18", Optional: true},
				{Name: "posSRS-ValidityAreaRRC-InactiveInitialUL-BWP-r18", Optional: true},
				{Name: "posSRS-ValidityAreaRRC-InactiveOutsideInitialUL-BWP-r18", Optional: true},
				{Name: "dl-PRS-MeasurementWithRxFH-RRC-ConnectedForRedCap-r18", Optional: true},
				{Name: "posSRS-TxFH-RRC-ConnectedForRedCap-r18", Optional: true},
				{Name: "posSRS-TxFH-RRC-InactiveForRedCap-r18", Optional: true},
				{Name: "posSRS-BWA-RRC-Inactive-r18", Optional: true},
				{Name: "posJointTriggerBySingleDCI-RRC-Connected-r18", Optional: true},
				{Name: "dl-PRS-MeasurementWithRxFH-RRC-InactiveforRedCap-r18", Optional: true},
				{Name: "dl-PRS-MeasurementWithRxFH-RRC-IdleforRedCap-r18", Optional: true},
				{Name: "spatialAdaptation-CSI-Feedback-r18", Optional: true},
				{Name: "spatialAdaptation-CSI-FeedbackPUSCH-r18", Optional: true},
				{Name: "spatialAdaptation-CSI-FeedbackAperiodic-r18", Optional: true},
				{Name: "spatialAdaptation-CSI-FeedbackPUCCH-r18", Optional: true},
				{Name: "powerAdaptation-CSI-Feedback-r18", Optional: true},
				{Name: "powerAdaptation-CSI-FeedbackPUSCH-r18", Optional: true},
				{Name: "powerAdaptation-CSI-FeedbackAperiodic-r18", Optional: true},
				{Name: "powerAdaptation-CSI-FeedbackPUCCH-r18", Optional: true},
				{Name: "nes-CellDTX-DRX-r18", Optional: true},
				{Name: "nes-CellDTX-DRX-DCI2-9-r18", Optional: true},
				{Name: "mixCodeBookSpatialAdaptation-r18", Optional: true},
				{Name: "simultaneousCSI-SubReportsPerCC-r18", Optional: true},
				{Name: "ntn-DMRS-BundlingNGSO-r18", Optional: true},
				{Name: "ltm-BeamIndicationJointTCI-r18", Optional: true},
				{Name: "dummy-ltm-MAC-CE-JointTCI-r18", Optional: true},
				{Name: "ltm-BeamIndicationSeparateTCI-r18", Optional: true},
				{Name: "dummy-ltm-MAC-CE-SeparateTCI-r18", Optional: true},
				{Name: "rach-EarlyTA-Measurement-r18", Optional: true},
				{Name: "ue-TA-Measurement-r18", Optional: true},
				{Name: "ta-IndicationCellSwitch-r18", Optional: true},
				{Name: "triggeredHARQ-CodebookRetxDCI-1-3-r18", Optional: true},
				{Name: "unifiedJointTCI-MultiMAC-CE-DCI-1-3-r18", Optional: true},
				{Name: "unifiedSeparateTCI-MultiMAC-CE-IntraCell-r18", Optional: true},
				{Name: "multiPUSCH-CG-r18", Optional: true},
				{Name: "multiPUSCH-ActiveConfiguredGrant-r18", Optional: true},
				{Name: "jointReleaseDCI-r18", Optional: true},
				{Name: "cg-PUSCH-UTO-UCI-Ind-r18", Optional: true},
				{Name: "pdcch-MonitoringResumptionAfterUL-NACK-r18", Optional: true},
				{Name: "support3MHz-ChannelBW-Symmetric-r18", Optional: true},
				{Name: "support3MHz-ChannelBW-Asymmetric-r18", Optional: true},
				{Name: "support12PRB-CORESET0-r18", Optional: true},
				{Name: "nr-PDCCH-OverlapLTE-CRS-RE-r18", Optional: true},
				{Name: "nr-PDCCH-OverlapLTE-CRS-RE-MultiPatterns-r18", Optional: true},
				{Name: "nr-PDCCH-OverlapLTE-CRS-RE-Span-3-4-r18", Optional: true},
				{Name: "twoRateMatchingEUTRA-CRS-patterns-3-4-r18", Optional: true},
				{Name: "overlapRateMatchingEUTRA-CRS-Patterns-3-4-Diff-CS-Pool-r18", Optional: true},
				{Name: "ncd-SSB-BWP-Wor-r18", Optional: true},
				{Name: "rlm-BM-BFD-CSI-RS-OutsideActiveBWP-r18", Optional: true},
				{Name: "prach-CoverageEnh-r18", Optional: true},
				{Name: "prach-Repetition-r18", Optional: true},
				{Name: "dynamicWaveformSwitch-r18", Optional: true},
				{Name: "dynamicWaveformSwitchPHR-r18", Optional: true},
				{Name: "dynamicWaveformSwitchIntraCA-r18", Optional: true},
				{Name: "multiPUSCH-SingleDCI-NonConsSlots-r18", Optional: true},
				{Name: "pdc-maxNumberPRS-ResourceProcessedPerSlot-r18", Optional: true},
				{Name: "intraSlot-PDSCH-MulticastInactive-r18", Optional: true},
				{Name: "multicastInactive-r18", Optional: true},
				{Name: "thresholdBasedMulticastResume-r18", Optional: true},
				{Name: "lowerMSD-r18", Optional: true},
				{Name: "lowerMSD-ENDC-r18", Optional: true},
				{Name: "enhancedChannelRaster-r18", Optional: true},
				{Name: "fastBeamSweepingMultiRx-r18", Optional: true},
				{Name: "beamSweepingFactorReduction-r18", Optional: true},
				{Name: "simultaneousReceptionTwoQCL-r18", Optional: true},
				{Name: "measEnhCAInterFreqFR2-r18", Optional: true},
				{Name: "tci-StateSwitchInd-r18", Optional: true},
				{Name: "antennaArrayType-r18", Optional: true},
				{Name: "locationBasedCondHandoverATG-r18", Optional: true},
				{Name: "maxOutputPowerATG-r18", Optional: true},
				{Name: "ltm-FastProcessingConfig-r18", Optional: true},
				{Name: "measValidationReportEMR-r18", Optional: true},
				{Name: "measValidationReportReselectionMeasurements-r18", Optional: true},
				{Name: "eventA4BasedCondHandoverNES-r18", Optional: true},
				{Name: "nesBasedCondHandoverWithDCI-r18", Optional: true},
				{Name: "rach-LessHandoverCG-r18", Optional: true},
				{Name: "rach-LessHandoverDG-r18", Optional: true},
				{Name: "locationBasedCondHandoverEMC-r18", Optional: true},
				{Name: "mt-CG-SDT-r18", Optional: true},
				{Name: "posSRS-PreconfigureRRC-InactiveInitialUL-BWP-r18", Optional: true},
				{Name: "posSRS-PreconfigureRRC-InactiveOutsideInitialUL-BWP-r18", Optional: true},
				{Name: "cg-SDT-PeriodicityExt-r18", Optional: true},
				{Name: "supportOf2RxXR-r18", Optional: true},
				{Name: "condHandoverWithCandSCG-change-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtPosUETAAutoAdjustmentR18Constraints)
			if err != nil {
				return err
			}
			ie.PosUE_TA_AutoAdjustment_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSValidityAreaRRCInactiveInitialULBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_ValidityAreaRRC_InactiveInitialUL_BWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSValidityAreaRRCInactiveOutsideInitialULBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_ValidityAreaRRC_InactiveOutsideInitialUL_BWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18 = new(DL_PRS_MeasurementWithRxFH_RRC_Connected_r18)
			if err := ie.Dl_PRS_MeasurementWithRxFH_RRC_ConnectedForRedCap_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18 = new(PosSRS_TxFrequencyHoppingRRC_Connected_r18)
			if err := ie.PosSRS_TxFH_RRC_ConnectedForRedCap_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18 = new(PosSRS_TxFrequencyHoppingRRC_Inactive_r18)
			if err := ie.PosSRS_TxFH_RRC_InactiveForRedCap_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.PosSRS_BWA_RRC_Inactive_r18 = new(PosSRS_BWA_RRC_Inactive_r18)
			if err := ie.PosSRS_BWA_RRC_Inactive_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtPosJointTriggerBySingleDCIRRCConnectedR18Constraints)
			if err != nil {
				return err
			}
			ie.PosJointTriggerBySingleDCI_RRC_Connected_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandNRExtDlPRSMeasurementWithRxFHRRCInactiveforRedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.Dl_PRS_MeasurementWithRxFH_RRC_InactiveforRedCap_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(bandNRExtDlPRSMeasurementWithRxFHRRCIdleforRedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.Dl_PRS_MeasurementWithRxFH_RRC_IdleforRedCap_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			ie.SpatialAdaptation_CSI_Feedback_r18 = &struct {
				CsiFeedbackType_r18            int64
				MaxNumberLmax_r18              int64
				MaxNumberCSI_ResourcePerCC_r18 struct {
					SdType1_Resource_r18 int64
					SdType2_Resource_r18 int64
				}
				MaxNumberTotalCSI_ResourcePerCC_r18 struct {
					SdType1_Resource_r18 int64
					SdType2_Resource_r18 int64
				}
				TotalNumberCSI_Reporting_r18 int64
			}{}
			c := ie.SpatialAdaptation_CSI_Feedback_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackR18CsiFeedbackTypeR18Constraints)
				if err != nil {
					return err
				}
				c.CsiFeedbackType_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				c := &c.MaxNumberCSI_ResourcePerCC_r18
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 32))
					if err != nil {
						return err
					}
					c.SdType1_Resource_r18 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 32))
					if err != nil {
						return err
					}
					c.SdType2_Resource_r18 = v
				}
			}
			{
				c := &c.MaxNumberTotalCSI_ResourcePerCC_r18
				{
					v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints)
					if err != nil {
						return err
					}
					c.SdType1_Resource_r18 = v
				}
				{
					v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints)
					if err != nil {
						return err
					}
					c.SdType2_Resource_r18 = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(11) {
			ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18 = &struct {
				CsiFeedbackType_r18                 int64
				MaxNumberLmax_r18                   int64
				SubReportCSI_r18                    int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.SpatialAdaptation_CSI_FeedbackPUSCH_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18CsiFeedbackTypeR18Constraints)
				if err != nil {
					return err
				}
				c.CsiFeedbackType_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 12))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(12) {
			ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18 = &struct {
				CsiFeedbackType_r18            int64
				MaxNumberLmax_r18              int64
				SubReportCSI_r18               int64
				MaxNumberCSI_ResourcePerCC_r18 struct {
					SdType1_Resource_r18 int64
					SdType2_Resource_r18 int64
				}
				MaxNumberTotalCSI_ResourcePerCC_r18 struct {
					SdType1_Resource_r18 int64
					SdType2_Resource_r18 int64
				}
				TotalNumberCSI_Reporting_r18 int64
			}{}
			c := ie.SpatialAdaptation_CSI_FeedbackAperiodic_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18CsiFeedbackTypeR18Constraints)
				if err != nil {
					return err
				}
				c.CsiFeedbackType_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				c := &c.MaxNumberCSI_ResourcePerCC_r18
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 32))
					if err != nil {
						return err
					}
					c.SdType1_Resource_r18 = v
				}
				{
					v, err := dx.DecodeInteger(per.Constrained(1, 32))
					if err != nil {
						return err
					}
					c.SdType2_Resource_r18 = v
				}
			}
			{
				c := &c.MaxNumberTotalCSI_ResourcePerCC_r18
				{
					v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType1ResourceR18Constraints)
					if err != nil {
						return err
					}
					c.SdType1_Resource_r18 = v
				}
				{
					v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18SdType2ResourceR18Constraints)
					if err != nil {
						return err
					}
					c.SdType2_Resource_r18 = v
				}
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 12))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(13) {
			ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18 = &struct {
				CsiFeedbackType_r18                 int64
				MaxNumberLmax_r18                   int64
				SubReportCSI_r18                    int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.SpatialAdaptation_CSI_FeedbackPUCCH_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18CsiFeedbackTypeR18Constraints)
				if err != nil {
					return err
				}
				c.CsiFeedbackType_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtSpatialAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(14) {
			ie.PowerAdaptation_CSI_Feedback_r18 = &struct {
				MaxNumberLmax_r18                   int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.PowerAdaptation_CSI_Feedback_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtPowerAdaptationCSIFeedbackR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			ie.PowerAdaptation_CSI_FeedbackPUSCH_r18 = &struct {
				MaxNumberLmax_r18                   int64
				SubReportCSI_r18                    int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.PowerAdaptation_CSI_FeedbackPUSCH_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtPowerAdaptationCSIFeedbackPUSCHR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 12))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(16) {
			ie.PowerAdaptation_CSI_FeedbackAperiodic_r18 = &struct {
				MaxNumberLmax_r18                   int64
				SubReportCSI_r18                    int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.PowerAdaptation_CSI_FeedbackAperiodic_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtPowerAdaptationCSIFeedbackAperiodicR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 12))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(17) {
			ie.PowerAdaptation_CSI_FeedbackPUCCH_r18 = &struct {
				MaxNumberLmax_r18                   int64
				SubReportCSI_r18                    int64
				MaxNumberCSI_ResourcePerCC_r18      int64
				MaxNumberTotalCSI_ResourcePerCC_r18 int64
				TotalNumberCSI_Reporting_r18        int64
			}{}
			c := ie.PowerAdaptation_CSI_FeedbackPUCCH_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberLmax_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.SubReportCSI_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtPowerAdaptationCSIFeedbackPUCCHR18MaxNumberTotalCSIResourcePerCCR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberTotalCSI_ResourcePerCC_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.TotalNumberCSI_Reporting_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(bandNRExtNesCellDTXDRXR18Constraints)
			if err != nil {
				return err
			}
			ie.Nes_CellDTX_DRX_r18 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(bandNRExtNesCellDTXDRXDCI29R18Constraints)
			if err != nil {
				return err
			}
			ie.Nes_CellDTX_DRX_DCI2_9_r18 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(bandNRExtMixCodeBookSpatialAdaptationR18Constraints)
			if err != nil {
				return err
			}
			ie.MixCodeBookSpatialAdaptation_r18 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeInteger(bandNRSimultaneousCSISubReportsPerCCR18Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousCSI_SubReportsPerCC_r18 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(bandNRExtNtnDMRSBundlingNGSOR18Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_DMRS_BundlingNGSO_r18 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			ie.Ltm_BeamIndicationJointTCI_r18 = &struct {
				MaxNumberJointTCI_PerCell_r18     int64
				Qcl_Resource_r18                  int64
				MaxNumberJointTCI_AcrossCells_r18 int64
				MaxNumberCells_r18                int64
			}{}
			c := ie.Ltm_BeamIndicationJointTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationJointTCIR18MaxNumberJointTCIPerCellR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationJointTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 128))
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_AcrossCells_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberCells_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(24) {
			ie.Dummy_Ltm_MAC_CE_JointTCI_r18 = &struct {
				Qcl_Resource_r18                  int64
				MaxNumberJointTCI_PerCell_r18     int64
				MaxNumberJointTCI_AcrossCells_r18 int64
			}{}
			c := ie.Dummy_Ltm_MAC_CE_JointTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtDummyLtmMACCEJointTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtDummyLtmMACCEJointTCIR18MaxNumberJointTCIAcrossCellsR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_AcrossCells_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(25) {
			ie.Ltm_BeamIndicationSeparateTCI_r18 = &struct {
				MaxNumberDL_TCI_PerCell_r18     int64
				MaxNumberUL_TCI_PerCell_r18     int64
				Qcl_Resource_r18                int64
				MaxNumberDL_TCI_AcrossCells_r18 int64
				MaxNumberUL_TCI_AcrossCells_r18 int64
				MaxNumberCells_r18              int64
			}{}
			c := ie.Ltm_BeamIndicationSeparateTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberDLTCIPerCellR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationSeparateTCIR18MaxNumberULTCIPerCellR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationSeparateTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 128))
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_AcrossCells_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 64))
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_AcrossCells_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberCells_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(26) {
			ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18 = &struct {
				Qcl_Resource_r18                int64
				MaxNumberDL_TCI_PerCell_r18     int64
				MaxNumberUL_TCI_PerCell_r18     int64
				MaxNumberDL_TCI_AcrossCells_r18 int64
				MaxNumberUL_TCI_AcrossCells_r18 int64
			}{}
			c := ie.Dummy_Ltm_MAC_CE_SeparateTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtDummyLtmMACCESeparateTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberDLTCIAcrossCellsR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_AcrossCells_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtDummyLtmMACCESeparateTCIR18MaxNumberULTCIAcrossCellsR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_AcrossCells_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(27) {
			v, err := dx.DecodeInteger(bandNRRachEarlyTAMeasurementR18Constraints)
			if err != nil {
				return err
			}
			ie.Rach_EarlyTA_Measurement_r18 = &v
		}

		if groupSeq.IsComponentPresent(28) {
			v, err := dx.DecodeInteger(bandNRUeTAMeasurementR18Constraints)
			if err != nil {
				return err
			}
			ie.Ue_TA_Measurement_r18 = &v
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(bandNRExtTaIndicationCellSwitchR18Constraints)
			if err != nil {
				return err
			}
			ie.Ta_IndicationCellSwitch_r18 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18 = &struct {
				MinHARQ_Retx_Offset_r18 int64
				MaxHARQ_Retx_Offset_r18 int64
			}{}
			c := ie.TriggeredHARQ_CodebookRetxDCI_1_3_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxDCI13R18MinHARQRetxOffsetR18Constraints)
				if err != nil {
					return err
				}
				c.MinHARQ_Retx_Offset_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxDCI13R18MaxHARQRetxOffsetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxHARQ_Retx_Offset_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(31) {
			ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18 = &struct {
				MinBeamApplicationTime_r18 struct {
					Choice  int
					Fr1_r18 *struct {
						Scs_15kHz_r18 *int64
						Scs_30kHz_r18 *int64
						Scs_60kHz_r18 *int64
					}
					Fr2_r18 *struct {
						Scs_60kHz_r18  *int64
						Scs_120kHz_r18 *int64
					}
				}
				MaxActivatedTCI_PerCC_r18 *int64
			}{}
			c := ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18
			bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Constraints)
			if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MinBeamApplicationTime_r18.Choice = int(index)
				switch index {
				case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr1_r18:
					c.MinBeamApplicationTime_r18.Fr1_r18 = &struct {
						Scs_15kHz_r18 *int64
						Scs_30kHz_r18 *int64
						Scs_60kHz_r18 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r18.Fr1_r18)
					bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Constraints)
					if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r18) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r18) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r18) = v
					}
				case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_r18_MinBeamApplicationTime_r18_Fr2_r18:
					c.MinBeamApplicationTime_r18.Fr2_r18 = &struct {
						Scs_60kHz_r18  *int64
						Scs_120kHz_r18 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r18.Fr2_r18)
					bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Constraints)
					if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r18) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13R18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r18) = v
					}
				}
			}
			if bandNRExtUnifiedJointTCIMultiMACCEDCI13R18Seq.IsComponentPresent(1) {
				c.MaxActivatedTCI_PerCC_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedTCI_PerCC_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(32) {
			ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18 = &struct {
				MinBeamApplicationTime_r18 struct {
					Choice  int
					Fr1_r18 *struct {
						Scs_15kHz_r18 *int64
						Scs_30kHz_r18 *int64
						Scs_60kHz_r18 *int64
					}
					Fr2_r18 *struct {
						Scs_60kHz_r18  *int64
						Scs_120kHz_r18 *int64
					}
				}
				MaxActivatedDL_TCI_PerCC_r18 *int64
				MaxActivatedUL_TCI_PerCC_r18 *int64
			}{}
			c := ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18
			bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Constraints)
			if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MinBeamApplicationTime_r18.Choice = int(index)
				switch index {
				case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr1_r18:
					c.MinBeamApplicationTime_r18.Fr1_r18 = &struct {
						Scs_15kHz_r18 *int64
						Scs_30kHz_r18 *int64
						Scs_60kHz_r18 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r18.Fr1_r18)
					bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Constraints)
					if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs15kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r18) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs30kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r18) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr1R18Scs60kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r18) = v
					}
				case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_r18_MinBeamApplicationTime_r18_Fr2_r18:
					c.MinBeamApplicationTime_r18.Fr2_r18 = &struct {
						Scs_60kHz_r18  *int64
						Scs_120kHz_r18 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r18.Fr2_r18)
					bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Constraints)
					if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs60kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r18) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r18 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18MinBeamApplicationTimeR18Fr2R18Scs120kHzR18Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r18) = v
					}
				}
			}
			if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq.IsComponentPresent(1) {
				c.MaxActivatedDL_TCI_PerCC_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedDL_TCI_PerCC_r18) = v
			}
			if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellR18Seq.IsComponentPresent(2) {
				c.MaxActivatedUL_TCI_PerCC_r18 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedUL_TCI_PerCC_r18) = v
			}
		}

		if groupSeq.IsComponentPresent(33) {
			v, err := dx.DecodeEnumerated(bandNRExtMultiPUSCHCGR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_CG_r18 = &v
		}

		if groupSeq.IsComponentPresent(34) {
			ie.MultiPUSCH_ActiveConfiguredGrant_r18 = &struct {
				MaxNumberConfigsPerBWP    int64
				MaxNumberConfigsAllCC_FR1 int64
				MaxNumberConfigsAllCC_FR2 int64
			}{}
			c := ie.MultiPUSCH_ActiveConfiguredGrant_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtMultiPUSCHActiveConfiguredGrantR18MaxNumberConfigsPerBWPConstraints)
				if err != nil {
					return err
				}
				c.MaxNumberConfigsPerBWP = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 32))
				if err != nil {
					return err
				}
				c.MaxNumberConfigsAllCC_FR1 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 32))
				if err != nil {
					return err
				}
				c.MaxNumberConfigsAllCC_FR2 = v
			}
		}

		if groupSeq.IsComponentPresent(35) {
			v, err := dx.DecodeEnumerated(bandNRExtJointReleaseDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.JointReleaseDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(36) {
			v, err := dx.DecodeEnumerated(bandNRExtCgPUSCHUTOUCIIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_PUSCH_UTO_UCI_Ind_r18 = &v
		}

		if groupSeq.IsComponentPresent(37) {
			v, err := dx.DecodeEnumerated(bandNRExtPdcchMonitoringResumptionAfterULNACKR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringResumptionAfterUL_NACK_r18 = &v
		}

		if groupSeq.IsComponentPresent(38) {
			v, err := dx.DecodeEnumerated(bandNRExtSupport3MHzChannelBWSymmetricR18Constraints)
			if err != nil {
				return err
			}
			ie.Support3MHz_ChannelBW_Symmetric_r18 = &v
		}

		if groupSeq.IsComponentPresent(39) {
			v, err := dx.DecodeEnumerated(bandNRExtSupport3MHzChannelBWAsymmetricR18Constraints)
			if err != nil {
				return err
			}
			ie.Support3MHz_ChannelBW_Asymmetric_r18 = &v
		}

		if groupSeq.IsComponentPresent(40) {
			v, err := dx.DecodeEnumerated(bandNRExtSupport12PRBCORESET0R18Constraints)
			if err != nil {
				return err
			}
			ie.Support12PRB_CORESET0_r18 = &v
		}

		if groupSeq.IsComponentPresent(41) {
			ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18 = &struct {
				OverlapInRE_r18     int64
				OverlapInSymbol_r18 int64
			}{}
			c := ie.Nr_PDCCH_OverlapLTE_CRS_RE_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInRER18Constraints)
				if err != nil {
					return err
				}
				c.OverlapInRE_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtNrPDCCHOverlapLTECRSRER18OverlapInSymbolR18Constraints)
				if err != nil {
					return err
				}
				c.OverlapInSymbol_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(42) {
			v, err := dx.DecodeEnumerated(bandNRExtNrPDCCHOverlapLTECRSREMultiPatternsR18Constraints)
			if err != nil {
				return err
			}
			ie.Nr_PDCCH_OverlapLTE_CRS_RE_MultiPatterns_r18 = &v
		}

		if groupSeq.IsComponentPresent(43) {
			v, err := dx.DecodeEnumerated(bandNRExtNrPDCCHOverlapLTECRSRESpan34R18Constraints)
			if err != nil {
				return err
			}
			ie.Nr_PDCCH_OverlapLTE_CRS_RE_Span_3_4_r18 = &v
		}

		if groupSeq.IsComponentPresent(44) {
			ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18 = &struct {
				MaxNumberPatterns_r18            int64
				MaxNumberNon_OverlapPatterns_r18 int64
			}{}
			c := ie.TwoRateMatchingEUTRA_CRS_Patterns_3_4_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 6))
				if err != nil {
					return err
				}
				c.MaxNumberPatterns_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 3))
				if err != nil {
					return err
				}
				c.MaxNumberNon_OverlapPatterns_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(bandNRExtOverlapRateMatchingEUTRACRSPatterns34DiffCSPoolR18Constraints)
			if err != nil {
				return err
			}
			ie.OverlapRateMatchingEUTRA_CRS_Patterns_3_4_Diff_CS_Pool_r18 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(bandNRExtNcdSSBBWPWorR18Constraints)
			if err != nil {
				return err
			}
			ie.Ncd_SSB_BWP_Wor_r18 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			v, err := dx.DecodeEnumerated(bandNRExtRlmBMBFDCSIRSOutsideActiveBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.Rlm_BM_BFD_CSI_RS_OutsideActiveBWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(48) {
			v, err := dx.DecodeEnumerated(bandNRExtPrachCoverageEnhR18Constraints)
			if err != nil {
				return err
			}
			ie.Prach_CoverageEnh_r18 = &v
		}

		if groupSeq.IsComponentPresent(49) {
			v, err := dx.DecodeEnumerated(bandNRExtPrachRepetitionR18Constraints)
			if err != nil {
				return err
			}
			ie.Prach_Repetition_r18 = &v
		}

		if groupSeq.IsComponentPresent(50) {
			v, err := dx.DecodeEnumerated(bandNRExtDynamicWaveformSwitchR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicWaveformSwitch_r18 = &v
		}

		if groupSeq.IsComponentPresent(51) {
			v, err := dx.DecodeEnumerated(bandNRExtDynamicWaveformSwitchPHRR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicWaveformSwitchPHR_r18 = &v
		}

		if groupSeq.IsComponentPresent(52) {
			v, err := dx.DecodeInteger(bandNRDynamicWaveformSwitchIntraCAR18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicWaveformSwitchIntraCA_r18 = &v
		}

		if groupSeq.IsComponentPresent(53) {
			v, err := dx.DecodeEnumerated(bandNRExtMultiPUSCHSingleDCINonConsSlotsR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_SingleDCI_NonConsSlots_r18 = &v
		}

		if groupSeq.IsComponentPresent(54) {
			ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18 = &struct {
				Fr1_r18 struct {
					Scs_15kHz_r18 *int64
					Scs_30kHz_r18 *int64
					Scs_60kHz_r18 *int64
				}
				Fr2_r18 struct {
					Scs_60kHz_r18  *int64
					Scs_120kHz_r18 *int64
				}
			}{}
			c := ie.Pdc_MaxNumberPRS_ResourceProcessedPerSlot_r18
			{
				c := &c.Fr1_r18
				bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq := dx.NewSequenceDecoder(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Constraints)
				if err := bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs15kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz_r18) = v
				}
				if bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs30kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz_r18) = v
				}
				if bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Seq.IsComponentPresent(2) {
					c.Scs_60kHz_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr1R18Scs60kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r18) = v
				}
			}
			{
				c := &c.Fr2_r18
				bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq := dx.NewSequenceDecoder(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Constraints)
				if err := bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq.IsComponentPresent(0) {
					c.Scs_60kHz_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs60kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r18) = v
				}
				if bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Seq.IsComponentPresent(1) {
					c.Scs_120kHz_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtPdcMaxNumberPRSResourceProcessedPerSlotR18Fr2R18Scs120kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_r18) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(55) {
			v, err := dx.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.IntraSlot_PDSCH_MulticastInactive_r18 = &v
		}

		if groupSeq.IsComponentPresent(56) {
			v, err := dx.DecodeEnumerated(bandNRExtMulticastInactiveR18Constraints)
			if err != nil {
				return err
			}
			ie.MulticastInactive_r18 = &v
		}

		if groupSeq.IsComponentPresent(57) {
			v, err := dx.DecodeEnumerated(bandNRExtThresholdBasedMulticastResumeR18Constraints)
			if err != nil {
				return err
			}
			ie.ThresholdBasedMulticastResume_r18 = &v
		}

		if groupSeq.IsComponentPresent(58) {
			seqOf := dx.NewSequenceOfDecoder(bandNRExtLowerMSDR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.LowerMSD_r18 = make([]LowerMSD_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.LowerMSD_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(59) {
			seqOf := dx.NewSequenceOfDecoder(bandNRExtLowerMSDENDCR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.LowerMSD_ENDC_r18 = make([]LowerMSD_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.LowerMSD_ENDC_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(60) {
			v, err := dx.DecodeEnumerated(bandNRExtEnhancedChannelRasterR18Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedChannelRaster_r18 = &v
		}

		if groupSeq.IsComponentPresent(61) {
			v, err := dx.DecodeEnumerated(bandNRExtFastBeamSweepingMultiRxR18Constraints)
			if err != nil {
				return err
			}
			ie.FastBeamSweepingMultiRx_r18 = &v
		}

		if groupSeq.IsComponentPresent(62) {
			ie.BeamSweepingFactorReduction_r18 = &struct {
				ReduceForCellDetection    int64
				ReduceForSSB_L1_RSRP_Meas int64
			}{}
			c := ie.BeamSweepingFactorReduction_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtBeamSweepingFactorReductionR18ReduceForCellDetectionConstraints)
				if err != nil {
					return err
				}
				c.ReduceForCellDetection = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(0, 7))
				if err != nil {
					return err
				}
				c.ReduceForSSB_L1_RSRP_Meas = v
			}
		}

		if groupSeq.IsComponentPresent(63) {
			v, err := dx.DecodeEnumerated(bandNRExtSimultaneousReceptionTwoQCLR18Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousReceptionTwoQCL_r18 = &v
		}

		if groupSeq.IsComponentPresent(64) {
			v, err := dx.DecodeEnumerated(bandNRExtMeasEnhCAInterFreqFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.MeasEnhCAInterFreqFR2_r18 = &v
		}

		if groupSeq.IsComponentPresent(65) {
			v, err := dx.DecodeEnumerated(bandNRExtTciStateSwitchIndR18Constraints)
			if err != nil {
				return err
			}
			ie.Tci_StateSwitchInd_r18 = &v
		}

		if groupSeq.IsComponentPresent(66) {
			v, err := dx.DecodeEnumerated(bandNRExtAntennaArrayTypeR18Constraints)
			if err != nil {
				return err
			}
			ie.AntennaArrayType_r18 = &v
		}

		if groupSeq.IsComponentPresent(67) {
			v, err := dx.DecodeEnumerated(bandNRExtLocationBasedCondHandoverATGR18Constraints)
			if err != nil {
				return err
			}
			ie.LocationBasedCondHandoverATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(68) {
			v, err := dx.DecodeInteger(bandNRMaxOutputPowerATGR18Constraints)
			if err != nil {
				return err
			}
			ie.MaxOutputPowerATG_r18 = &v
		}

		if groupSeq.IsComponentPresent(69) {
			ie.Ltm_FastProcessingConfig_r18 = &struct {
				MaxNumberStoredConfigCells_r18 int64
				MaxNumberConfigs_r18           int64
			}{}
			c := ie.Ltm_FastProcessingConfig_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmFastProcessingConfigR18MaxNumberStoredConfigCellsR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberStoredConfigCells_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberConfigs_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(70) {
			v, err := dx.DecodeEnumerated(bandNRExtMeasValidationReportEMRR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasValidationReportEMR_r18 = &v
		}

		if groupSeq.IsComponentPresent(71) {
			v, err := dx.DecodeEnumerated(bandNRExtMeasValidationReportReselectionMeasurementsR18Constraints)
			if err != nil {
				return err
			}
			ie.MeasValidationReportReselectionMeasurements_r18 = &v
		}

		if groupSeq.IsComponentPresent(72) {
			v, err := dx.DecodeEnumerated(bandNRExtEventA4BasedCondHandoverNESR18Constraints)
			if err != nil {
				return err
			}
			ie.EventA4BasedCondHandoverNES_r18 = &v
		}

		if groupSeq.IsComponentPresent(73) {
			v, err := dx.DecodeEnumerated(bandNRExtNesBasedCondHandoverWithDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.NesBasedCondHandoverWithDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(74) {
			v, err := dx.DecodeEnumerated(bandNRExtRachLessHandoverCGR18Constraints)
			if err != nil {
				return err
			}
			ie.Rach_LessHandoverCG_r18 = &v
		}

		if groupSeq.IsComponentPresent(75) {
			v, err := dx.DecodeEnumerated(bandNRExtRachLessHandoverDGR18Constraints)
			if err != nil {
				return err
			}
			ie.Rach_LessHandoverDG_r18 = &v
		}

		if groupSeq.IsComponentPresent(76) {
			v, err := dx.DecodeEnumerated(bandNRExtLocationBasedCondHandoverEMCR18Constraints)
			if err != nil {
				return err
			}
			ie.LocationBasedCondHandoverEMC_r18 = &v
		}

		if groupSeq.IsComponentPresent(77) {
			v, err := dx.DecodeEnumerated(bandNRExtMtCGSDTR18Constraints)
			if err != nil {
				return err
			}
			ie.Mt_CG_SDT_r18 = &v
		}

		if groupSeq.IsComponentPresent(78) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSPreconfigureRRCInactiveInitialULBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_PreconfigureRRC_InactiveInitialUL_BWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(79) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSPreconfigureRRCInactiveOutsideInitialULBWPR18Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_PreconfigureRRC_InactiveOutsideInitialUL_BWP_r18 = &v
		}

		if groupSeq.IsComponentPresent(80) {
			v, err := dx.DecodeEnumerated(bandNRExtCgSDTPeriodicityExtR18Constraints)
			if err != nil {
				return err
			}
			ie.Cg_SDT_PeriodicityExt_r18 = &v
		}

		if groupSeq.IsComponentPresent(81) {
			v, err := dx.DecodeEnumerated(bandNRExtSupportOf2RxXRR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportOf2RxXR_r18 = &v
		}

		if groupSeq.IsComponentPresent(82) {
			v, err := dx.DecodeEnumerated(bandNRExtCondHandoverWithCandSCGChangeR18Constraints)
			if err != nil {
				return err
			}
			ie.CondHandoverWithCandSCG_Change_r18 = &v
		}
	}

	// Extension group 16:
	if seq.IsExtensionPresent(16) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "mac-ParametersPerBand-r18", Optional: true},
				{Name: "channelBW-DL-NCR-r18", Optional: true},
				{Name: "channelBW-UL-NCR-r18", Optional: true},
				{Name: "ncr-PDSCH-64QAM-FR2-r18", Optional: true},
				{Name: "ltm-MCG-IntraFreq-r18", Optional: true},
				{Name: "ltm-SCG-IntraFreq-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Mac_ParametersPerBand_r18 = new(MAC_ParametersPerBand_r18)
			if err := ie.Mac_ParametersPerBand_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ChannelBW_DL_NCR_r18 = &struct {
				Choice     int
				Fr1_100mhz *struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}
				Fr2_200mhz *struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWDLNCRR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBW_DL_NCR_r18).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBW_DL_NCR_r18_Fr1_100mhz:
				(*ie.ChannelBW_DL_NCR_r18).Fr1_100mhz = &struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}{}
				c := (*(*ie.ChannelBW_DL_NCR_r18).Fr1_100mhz)
				bandNRExtChannelBWDLNCRR18Fr1100mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWDLNCRR18Fr1100mhzConstraints)
				if err := bandNRExtChannelBWDLNCRR18Fr1100mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWDLNCRR18Fr1100mhzSeq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLNCRR18Fr1100mhzScs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWDLNCRR18Fr1100mhzSeq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLNCRR18Fr1100mhzScs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWDLNCRR18Fr1100mhzSeq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLNCRR18Fr1100mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBW_DL_NCR_r18_Fr2_200mhz:
				(*ie.ChannelBW_DL_NCR_r18).Fr2_200mhz = &struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}{}
				c := (*(*ie.ChannelBW_DL_NCR_r18).Fr2_200mhz)
				bandNRExtChannelBWDLNCRR18Fr2200mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWDLNCRR18Fr2200mhzConstraints)
				if err := bandNRExtChannelBWDLNCRR18Fr2200mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWDLNCRR18Fr2200mhzSeq.IsComponentPresent(0) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLNCRR18Fr2200mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWDLNCRR18Fr2200mhzSeq.IsComponentPresent(1) {
					c.Scs_120kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWDLNCRR18Fr2200mhzScs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.ChannelBW_UL_NCR_r18 = &struct {
				Choice     int
				Fr1_100mhz *struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}
				Fr2_200mhz *struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtChannelBWULNCRR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelBW_UL_NCR_r18).Choice = int(index)
			switch index {
			case BandNR_Ext_ChannelBW_UL_NCR_r18_Fr1_100mhz:
				(*ie.ChannelBW_UL_NCR_r18).Fr1_100mhz = &struct {
					Scs_15kHz *int64
					Scs_30kHz *int64
					Scs_60kHz *int64
				}{}
				c := (*(*ie.ChannelBW_UL_NCR_r18).Fr1_100mhz)
				bandNRExtChannelBWULNCRR18Fr1100mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWULNCRR18Fr1100mhzConstraints)
				if err := bandNRExtChannelBWULNCRR18Fr1100mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWULNCRR18Fr1100mhzSeq.IsComponentPresent(0) {
					c.Scs_15kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULNCRR18Fr1100mhzScs15kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_15kHz) = v
				}
				if bandNRExtChannelBWULNCRR18Fr1100mhzSeq.IsComponentPresent(1) {
					c.Scs_30kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULNCRR18Fr1100mhzScs30kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_30kHz) = v
				}
				if bandNRExtChannelBWULNCRR18Fr1100mhzSeq.IsComponentPresent(2) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULNCRR18Fr1100mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
			case BandNR_Ext_ChannelBW_UL_NCR_r18_Fr2_200mhz:
				(*ie.ChannelBW_UL_NCR_r18).Fr2_200mhz = &struct {
					Scs_60kHz  *int64
					Scs_120kHz *int64
				}{}
				c := (*(*ie.ChannelBW_UL_NCR_r18).Fr2_200mhz)
				bandNRExtChannelBWULNCRR18Fr2200mhzSeq := dx.NewSequenceDecoder(bandNRExtChannelBWULNCRR18Fr2200mhzConstraints)
				if err := bandNRExtChannelBWULNCRR18Fr2200mhzSeq.DecodePreamble(); err != nil {
					return err
				}
				if bandNRExtChannelBWULNCRR18Fr2200mhzSeq.IsComponentPresent(0) {
					c.Scs_60kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULNCRR18Fr2200mhzScs60kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz) = v
				}
				if bandNRExtChannelBWULNCRR18Fr2200mhzSeq.IsComponentPresent(1) {
					c.Scs_120kHz = new(int64)
					v, err := dx.DecodeEnumerated(bandNRExtChannelBWULNCRR18Fr2200mhzScs120kHzConstraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtNcrPDSCH64QAMFR2R18Constraints)
			if err != nil {
				return err
			}
			ie.Ncr_PDSCH_64QAM_FR2_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmMCGIntraFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MCG_IntraFreq_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmSCGIntraFreqR18Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_SCG_IntraFreq_r18 = &v
		}
	}

	// Extension group 17:
	if seq.IsExtensionPresent(17) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ltm-MAC-CE-JointTCI-r18", Optional: true},
				{Name: "ltm-MAC-CE-SeparateTCI-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ltm_MAC_CE_JointTCI_r18 = &struct {
				Qcl_Resource_r18                  int64
				MaxNumberJointTCI_PerCell_r18     int64
				MaxNumberJointTCI_AcrossCells_r18 int64
			}{}
			c := ie.Ltm_MAC_CE_JointTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmMACCEJointTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberJointTCI_AcrossCells_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ltm_MAC_CE_SeparateTCI_r18 = &struct {
				Qcl_Resource_r18                int64
				MaxNumberDL_TCI_PerCell_r18     int64
				MaxNumberUL_TCI_PerCell_r18     int64
				MaxNumberDL_TCI_AcrossCells_r18 int64
				MaxNumberUL_TCI_AcrossCells_r18 int64
			}{}
			c := ie.Ltm_MAC_CE_SeparateTCI_r18
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmMACCESeparateTCIR18QclResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Qcl_Resource_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_PerCell_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberDL_TCI_AcrossCells_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberUL_TCI_AcrossCells_r18 = v
			}
		}
	}

	// Extension group 18:
	if seq.IsExtensionPresent(18) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "eventA4BasedCondHandoverATG-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtEventA4BasedCondHandoverATGR18Constraints)
			if err != nil {
				return err
			}
			ie.EventA4BasedCondHandoverATG_r18 = &v
		}
	}

	// Extension group 19:
	if seq.IsExtensionPresent(19) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "posSRS-TxFH-WithTimeWindowForRedCap-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSTxFHWithTimeWindowForRedCapR18Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_TxFH_WithTimeWindowForRedCap_r18 = &v
		}
	}

	// Extension group 20:
	if seq.IsExtensionPresent(20) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sbfd-Aware-r19", Optional: true},
				{Name: "ul-ResourceMutingCP-OFDM-r19", Optional: true},
				{Name: "ul-ResourceMutingDFTS-OFDM-r19", Optional: true},
				{Name: "ul-ResourceMutingCP-OFDM-CG-Type1-r19", Optional: true},
				{Name: "ul-ResourceMutingDFTS-OFDM-CG-Type1-r19", Optional: true},
				{Name: "l1-CLI-RSSI-MeasAndAperiodicReporting-r19", Optional: true},
				{Name: "l1-SRS-RSRP-MeasAndAperiodicReporting-r19", Optional: true},
				{Name: "od-SSB-NoAlwaysOn-RRC-r19", Optional: true},
				{Name: "od-SSB-NoAlwaysOn-RRC-MAC-CE-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-RRC-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-RRC-Diff-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-RRC-MAC-CE-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-RRC-MAC-CE-Diff-r19", Optional: true},
				{Name: "od-SSB-NoAlwaysOn-MAC-CE-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-MAC-CE-r19", Optional: true},
				{Name: "od-SSB-AlwaysOn-MAC-CE-Diff-r19", Optional: true},
				{Name: "ssb-BurstPeriodicityAdaptation-r19", Optional: true},
				{Name: "rach-AdaptationTimeDomain-r19", Optional: true},
				{Name: "ltm-BeamIndicationJointTCI-CSI-RS-r19", Optional: true},
				{Name: "ltm-MAC-CE-JointTCI-CSI-RS-r19", Optional: true},
				{Name: "ltm-BeamIndicationSeparateTCI-CSI-RS-r19", Optional: true},
				{Name: "ltm-MAC-CE-SeparateTCI-CSI-RS-r19", Optional: true},
				{Name: "ltm-CSI-RS-CSI-IM-Periodic-r19", Optional: true},
				{Name: "ltm-CSI-RS-CSI-IM-SP-r19", Optional: true},
				{Name: "ltm-WithoutCSI-IM-r19", Optional: true},
				{Name: "ssb-PeriodicityAddition-r19", Optional: true},
				{Name: "pdcch-RepetitionType0-r19", Optional: true},
				{Name: "pdcch-RepetitionTypeOthers-r19", Optional: true},
				{Name: "pdsch-RepetitionMsg4-r19", Optional: true},
				{Name: "ntn-Collision-RedCap-r19", Optional: true},
				{Name: "pusch-InterSlotOCC-r19", Optional: true},
				{Name: "triggeredHARQ-CodebookRetxDCI-1-3-Diff-r19", Optional: true},
				{Name: "unifiedJointTCI-MultiMAC-CE-DCI-1-3-Diff-r19", Optional: true},
				{Name: "unifiedSeparateTCI-MultiMAC-CE-IntraCell-Diff-r19", Optional: true},
				{Name: "posSRS-TxFH-RRC-ConnectedForNonRedCap-r19", Optional: true},
				{Name: "posSRS-TxFH-RRC-InactiveForNonRedCap-r19", Optional: true},
				{Name: "posSRS-TxFH-WithTimeWindowForNonRedCap-r19", Optional: true},
				{Name: "sr-TriggeredSSSG-Switching-r19", Optional: true},
				{Name: "pdcch-RepetitionType0-TN-r19", Optional: true},
				{Name: "pdcch-RepetitionTypeOthersTN-r19", Optional: true},
				{Name: "mpr-SingleCC-r19", Optional: true},
				{Name: "fastRx-BSF-MeasDelayReduction-r19", Optional: true},
				{Name: "fastSCellActivationEarlyMeas-r19", Optional: true},
				{Name: "od-SSB-AdditionalProcessingTime-r19", Optional: true},
				{Name: "aiml-BM-Case1-KnownRxBeam-r19", Optional: true},
				{Name: "aiml-BM-Case2-KnownRxBeam-r19", Optional: true},
				{Name: "ntn-PowerBoosting-ERedCap-r19", Optional: true},
				{Name: "earlyCSI-AcquisitionWithCSI-IM-r19", Optional: true},
				{Name: "earlyCSI-AcquisitionWithoutCSI-IM-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sbfd_Aware_r19 = &struct {
				NumOfPartialPRG_r19           int64
				Timeline_r19                  int64
				Sbfd_Rx_r19                   *int64
				Sbfd_Tx_r19                   *int64
				Sbfd_OneRACH_r19              *int64
				Sbfd_TwoRACH_r19              *int64
				PreambleRepetitionOneRACH_r19 *int64
				PreambleRepetitionTwoRACH_r19 *int64
			}{}
			c := ie.Sbfd_Aware_r19
			bandNRExtSbfdAwareR19Seq := dx.NewSequenceDecoder(bandNRExtSbfdAwareR19Constraints)
			if err := bandNRExtSbfdAwareR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19NumOfPartialPRGR19Constraints)
				if err != nil {
					return err
				}
				c.NumOfPartialPRG_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19TimelineR19Constraints)
				if err != nil {
					return err
				}
				c.Timeline_r19 = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(2) {
				c.Sbfd_Rx_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19SbfdRxR19Constraints)
				if err != nil {
					return err
				}
				(*c.Sbfd_Rx_r19) = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(3) {
				c.Sbfd_Tx_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19SbfdTxR19Constraints)
				if err != nil {
					return err
				}
				(*c.Sbfd_Tx_r19) = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(4) {
				c.Sbfd_OneRACH_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19SbfdOneRACHR19Constraints)
				if err != nil {
					return err
				}
				(*c.Sbfd_OneRACH_r19) = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(5) {
				c.Sbfd_TwoRACH_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19SbfdTwoRACHR19Constraints)
				if err != nil {
					return err
				}
				(*c.Sbfd_TwoRACH_r19) = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(6) {
				c.PreambleRepetitionOneRACH_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19PreambleRepetitionOneRACHR19Constraints)
				if err != nil {
					return err
				}
				(*c.PreambleRepetitionOneRACH_r19) = v
			}
			if bandNRExtSbfdAwareR19Seq.IsComponentPresent(7) {
				c.PreambleRepetitionTwoRACH_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtSbfdAwareR19PreambleRepetitionTwoRACHR19Constraints)
				if err != nil {
					return err
				}
				(*c.PreambleRepetitionTwoRACH_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(bandNRExtUlResourceMutingCPOFDMR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ResourceMutingCP_OFDM_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandNRExtUlResourceMutingDFTSOFDMR19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ResourceMutingDFTS_OFDM_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandNRExtUlResourceMutingCPOFDMCGType1R19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ResourceMutingCP_OFDM_CG_Type1_r19 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandNRExtUlResourceMutingDFTSOFDMCGType1R19Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ResourceMutingDFTS_OFDM_CG_Type1_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19 = &struct {
				MaxNumberOfConfiguredMeasResourceAcrossCC_r19 int64
				MaxNumberOfAperiodicReportSettingPerBWP_r19   int64
				MaxNumberOfConfiguredMeasResourcePerCC_r19    int64
				MaxNumberOfSimultaneousMeasResourcePerCC_r19  int64
				PeriodicReporting_r19                         *int64
				SemiPersistentMeasResource_r19                *int64
			}{}
			c := ie.L1_CLI_RSSI_MeasAndAperiodicReporting_r19
			bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq := dx.NewSequenceDecoder(bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Constraints)
			if err := bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfConfiguredMeasResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberOfAperiodicReportSettingPerBWP_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberOfConfiguredMeasResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 32))
				if err != nil {
					return err
				}
				c.MaxNumberOfSimultaneousMeasResourcePerCC_r19 = v
			}
			if bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq.IsComponentPresent(4) {
				c.PeriodicReporting_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19PeriodicReportingR19Constraints)
				if err != nil {
					return err
				}
				(*c.PeriodicReporting_r19) = v
			}
			if bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19Seq.IsComponentPresent(5) {
				c.SemiPersistentMeasResource_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtL1CLIRSSIMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints)
				if err != nil {
					return err
				}
				(*c.SemiPersistentMeasResource_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19 = &struct {
				MaxNumberOfConfiguredMeasResourceAcrossCC_r19 int64
				MaxNumberOfConfiguredMeasResourcePerCC_r19    int64
				MaxNumberOfSimultaneousMeasResourcePerCC_r19  int64
				MaxNumberOfAperiodicReportSettingPerBWP_r19   int64
				MaxNumberOfMeasResourceAcrossCCWithinSlot_r19 int64
				MaxNumberOfAperiodicMeasResourceAcrossCC_r19  int64
				Fdm_Reception_r19                             *int64
				SemiPersistentMeasResource_r19                *int64
			}{}
			c := ie.L1_SRS_RSRP_MeasAndAperiodicReporting_r19
			bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq := dx.NewSequenceDecoder(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Constraints)
			if err := bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfConfiguredMeasResourceAcrossCCR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfConfiguredMeasResourceAcrossCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxNumberOfConfiguredMeasResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.MaxNumberOfSimultaneousMeasResourcePerCC_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberOfAperiodicReportSettingPerBWP_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19MaxNumberOfMeasResourceAcrossCCWithinSlotR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberOfMeasResourceAcrossCCWithinSlot_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberOfAperiodicMeasResourceAcrossCC_r19 = v
			}
			if bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq.IsComponentPresent(6) {
				c.Fdm_Reception_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19FdmReceptionR19Constraints)
				if err != nil {
					return err
				}
				(*c.Fdm_Reception_r19) = v
			}
			if bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19Seq.IsComponentPresent(7) {
				c.SemiPersistentMeasResource_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtL1SRSRSRPMeasAndAperiodicReportingR19SemiPersistentMeasResourceR19Constraints)
				if err != nil {
					return err
				}
				(*c.SemiPersistentMeasResource_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBNoAlwaysOnRRCR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_NoAlwaysOn_RRC_r19 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBNoAlwaysOnRRCMACCER19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_NoAlwaysOn_RRC_MAC_CE_r19 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnRRCR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AlwaysOn_RRC_r19 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnRRCDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AlwaysOn_RRC_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnRRCMACCER19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AlwaysOn_RRC_MAC_CE_r19 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnRRCMACCEDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AlwaysOn_RRC_MAC_CE_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBNoAlwaysOnMACCER19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_NoAlwaysOn_MAC_CE_r19 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			ie.Od_SSB_AlwaysOn_MAC_CE_r19 = &struct {
				TimeRelation_r19       int64
				DeactivationScheme_r19 int64
			}{}
			c := ie.Od_SSB_AlwaysOn_MAC_CE_r19
			{
				v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnMACCER19TimeRelationR19Constraints)
				if err != nil {
					return err
				}
				c.TimeRelation_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnMACCER19DeactivationSchemeR19Constraints)
				if err != nil {
					return err
				}
				c.DeactivationScheme_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAlwaysOnMACCEDiffR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AlwaysOn_MAC_CE_Diff_r19 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(bandNRExtSsbBurstPeriodicityAdaptationR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_BurstPeriodicityAdaptation_r19 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(bandNRExtRachAdaptationTimeDomainR19Constraints)
			if err != nil {
				return err
			}
			ie.Rach_AdaptationTimeDomain_r19 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationJointTCICSIRSR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_BeamIndicationJointTCI_CSI_RS_r19 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmMACCEJointTCICSIRSR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MAC_CE_JointTCI_CSI_RS_r19 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmBeamIndicationSeparateTCICSIRSR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_BeamIndicationSeparateTCI_CSI_RS_r19 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmMACCESeparateTCICSIRSR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_MAC_CE_SeparateTCI_CSI_RS_r19 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			ie.Ltm_CSI_RS_CSI_IM_Periodic_r19 = &struct {
				MaxNumOfCSI_RS_Resource_r19  int64
				MaxNumOfCSI_RS_Ports_r19     int64
				MaxNumOfNZP_CSI_RS_Ports_r19 int64
				MaxRank_r19                  int64
				MaxNumOfCSI_IM_Resource_r19  int64
			}{}
			c := ie.Ltm_CSI_RS_CSI_IM_Periodic_r19
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmCSIRSCSIIMPeriodicR19MaxNumOfNZPCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfNZP_CSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_IM_Resource_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(23) {
			ie.Ltm_CSI_RS_CSI_IM_SP_r19 = &struct {
				MaxNumOfCSI_RS_Resource_r19  int64
				MaxNumOfCSI_RS_Ports_r19     int64
				MaxNumOfNZP_CSI_RS_Ports_r19 int64
				MaxRank_r19                  int64
				MaxNumOfCSI_IM_Resource_r19  int64
			}{}
			c := ie.Ltm_CSI_RS_CSI_IM_SP_r19
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtLtmCSIRSCSIIMSPR19MaxNumOfNZPCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfNZP_CSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_IM_Resource_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(bandNRExtLtmWithoutCSIIMR19Constraints)
			if err != nil {
				return err
			}
			ie.Ltm_WithoutCSI_IM_r19 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(bandNRExtSsbPeriodicityAdditionR19Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_PeriodicityAddition_r19 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			v, err := dx.DecodeEnumerated(bandNRExtPdcchRepetitionType0R19Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_RepetitionType0_r19 = &v
		}

		if groupSeq.IsComponentPresent(27) {
			v, err := dx.DecodeInteger(bandNRPdcchRepetitionTypeOthersR19Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_RepetitionTypeOthers_r19 = &v
		}

		if groupSeq.IsComponentPresent(28) {
			v, err := dx.DecodeEnumerated(bandNRExtPdschRepetitionMsg4R19Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_RepetitionMsg4_r19 = &v
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(bandNRExtNtnCollisionRedCapR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_Collision_RedCap_r19 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			ie.Pusch_InterSlotOCC_r19 = &struct {
				OccLength_r19      int64
				SatelliteOrbit_r19 int64
			}{}
			c := ie.Pusch_InterSlotOCC_r19
			{
				v, err := dx.DecodeEnumerated(bandNRExtPuschInterSlotOCCR19OccLengthR19Constraints)
				if err != nil {
					return err
				}
				c.OccLength_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtPuschInterSlotOCCR19SatelliteOrbitR19Constraints)
				if err != nil {
					return err
				}
				c.SatelliteOrbit_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(31) {
			ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19 = &struct {
				MinHARQ_Retx_Offset_r19 int64
				MaxHARQ_Retx_Offset_r19 int64
			}{}
			c := ie.TriggeredHARQ_CodebookRetxDCI_1_3_Diff_r19
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MinHARQRetxOffsetR19Constraints)
				if err != nil {
					return err
				}
				c.MinHARQ_Retx_Offset_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtTriggeredHARQCodebookRetxDCI13DiffR19MaxHARQRetxOffsetR19Constraints)
				if err != nil {
					return err
				}
				c.MaxHARQ_Retx_Offset_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(32) {
			ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19 = &struct {
				MinBeamApplicationTime_r19 struct {
					Choice  int
					Fr1_r19 *struct {
						Scs_15kHz_r19 *int64
						Scs_30kHz_r19 *int64
						Scs_60kHz_r19 *int64
					}
					Fr2_r19 *struct {
						Scs_60kHz_r19  *int64
						Scs_120kHz_r19 *int64
					}
				}
				MaxActivatedTCI_PerCC_r19 *int64
			}{}
			c := ie.UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19
			bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Constraints)
			if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MinBeamApplicationTime_r19.Choice = int(index)
				switch index {
				case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19:
					c.MinBeamApplicationTime_r19.Fr1_r19 = &struct {
						Scs_15kHz_r19 *int64
						Scs_30kHz_r19 *int64
						Scs_60kHz_r19 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r19.Fr1_r19)
					bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Constraints)
					if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r19) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r19) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r19) = v
					}
				case BandNR_Ext_UnifiedJointTCI_MultiMAC_CE_DCI_1_3_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19:
					c.MinBeamApplicationTime_r19.Fr2_r19 = &struct {
						Scs_60kHz_r19  *int64
						Scs_120kHz_r19 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r19.Fr2_r19)
					bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Constraints)
					if err := bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r19) = v
					}
					if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r19) = v
					}
				}
			}
			if bandNRExtUnifiedJointTCIMultiMACCEDCI13DiffR19Seq.IsComponentPresent(1) {
				c.MaxActivatedTCI_PerCC_r19 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedTCI_PerCC_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(33) {
			ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19 = &struct {
				MinBeamApplicationTime_r19 struct {
					Choice  int
					Fr1_r19 *struct {
						Scs_15kHz_r19 *int64
						Scs_30kHz_r19 *int64
						Scs_60kHz_r19 *int64
					}
					Fr2_r19 *struct {
						Scs_60kHz_r19  *int64
						Scs_120kHz_r19 *int64
					}
				}
				MaxActivatedDL_TCI_PerCC_r19 *int64
				MaxActivatedUL_TCI_PerCC_r19 *int64
			}{}
			c := ie.UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19
			bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Constraints)
			if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MinBeamApplicationTime_r19.Choice = int(index)
				switch index {
				case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr1_r19:
					c.MinBeamApplicationTime_r19.Fr1_r19 = &struct {
						Scs_15kHz_r19 *int64
						Scs_30kHz_r19 *int64
						Scs_60kHz_r19 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r19.Fr1_r19)
					bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Constraints)
					if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs15kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r19) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs30kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r19) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr1R19Scs60kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r19) = v
					}
				case BandNR_Ext_UnifiedSeparateTCI_MultiMAC_CE_IntraCell_Diff_r19_MinBeamApplicationTime_r19_Fr2_r19:
					c.MinBeamApplicationTime_r19.Fr2_r19 = &struct {
						Scs_60kHz_r19  *int64
						Scs_120kHz_r19 *int64
					}{}
					c := (*c.MinBeamApplicationTime_r19.Fr2_r19)
					bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq := dx.NewSequenceDecoder(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Constraints)
					if err := bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs60kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r19) = v
					}
					if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r19 = new(int64)
						v, err := dx.DecodeEnumerated(bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19MinBeamApplicationTimeR19Fr2R19Scs120kHzR19Constraints)
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r19) = v
					}
				}
			}
			if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq.IsComponentPresent(1) {
				c.MaxActivatedDL_TCI_PerCC_r19 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedDL_TCI_PerCC_r19) = v
			}
			if bandNRExtUnifiedSeparateTCIMultiMACCEIntraCellDiffR19Seq.IsComponentPresent(2) {
				c.MaxActivatedUL_TCI_PerCC_r19 = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				(*c.MaxActivatedUL_TCI_PerCC_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(34) {
			ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19 = new(PosSRS_TxFrequencyHoppingRRC_ConnectedNonRedCap_r19)
			if err := ie.PosSRS_TxFH_RRC_ConnectedForNonRedCap_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(35) {
			ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19 = new(PosSRS_TxFrequencyHoppingRRC_InactiveNonRedCap_r19)
			if err := ie.PosSRS_TxFH_RRC_InactiveForNonRedCap_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(36) {
			v, err := dx.DecodeEnumerated(bandNRExtPosSRSTxFHWithTimeWindowForNonRedCapR19Constraints)
			if err != nil {
				return err
			}
			ie.PosSRS_TxFH_WithTimeWindowForNonRedCap_r19 = &v
		}

		if groupSeq.IsComponentPresent(37) {
			v, err := dx.DecodeEnumerated(bandNRExtSrTriggeredSSSGSwitchingR19Constraints)
			if err != nil {
				return err
			}
			ie.Sr_TriggeredSSSG_Switching_r19 = &v
		}

		if groupSeq.IsComponentPresent(38) {
			v, err := dx.DecodeEnumerated(bandNRExtPdcchRepetitionType0TNR19Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_RepetitionType0_TN_r19 = &v
		}

		if groupSeq.IsComponentPresent(39) {
			v, err := dx.DecodeInteger(bandNRPdcchRepetitionTypeOthersTNR19Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_RepetitionTypeOthersTN_r19 = &v
		}

		if groupSeq.IsComponentPresent(40) {
			ie.Mpr_SingleCC_r19 = &struct {
				Choice                         int
				Mpr_SingleCC_SingleValue_r19   *int64
				Mpr_SingleCC_MultipleValue_r19 *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(bandNRExtMprSingleCCR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Mpr_SingleCC_r19).Choice = int(index)
			switch index {
			case BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_SingleValue_r19:
				(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_SingleValue_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMprSingleCCR19MprSingleCCSingleValueR19Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_SingleValue_r19) = v
			case BandNR_Ext_Mpr_SingleCC_r19_Mpr_SingleCC_MultipleValue_r19:
				(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_MultipleValue_r19 = new(int64)
				v, err := dx.DecodeEnumerated(bandNRExtMprSingleCCR19MprSingleCCMultipleValueR19Constraints)
				if err != nil {
					return err
				}
				(*(*ie.Mpr_SingleCC_r19).Mpr_SingleCC_MultipleValue_r19) = v
			}
		}

		if groupSeq.IsComponentPresent(41) {
			v, err := dx.DecodeEnumerated(bandNRExtFastRxBSFMeasDelayReductionR19Constraints)
			if err != nil {
				return err
			}
			ie.FastRx_BSF_MeasDelayReduction_r19 = &v
		}

		if groupSeq.IsComponentPresent(42) {
			v, err := dx.DecodeEnumerated(bandNRExtFastSCellActivationEarlyMeasR19Constraints)
			if err != nil {
				return err
			}
			ie.FastSCellActivationEarlyMeas_r19 = &v
		}

		if groupSeq.IsComponentPresent(43) {
			v, err := dx.DecodeEnumerated(bandNRExtOdSSBAdditionalProcessingTimeR19Constraints)
			if err != nil {
				return err
			}
			ie.Od_SSB_AdditionalProcessingTime_r19 = &v
		}

		if groupSeq.IsComponentPresent(44) {
			v, err := dx.DecodeEnumerated(bandNRExtAimlBMCase1KnownRxBeamR19Constraints)
			if err != nil {
				return err
			}
			ie.Aiml_BM_Case1_KnownRxBeam_r19 = &v
		}

		if groupSeq.IsComponentPresent(45) {
			v, err := dx.DecodeEnumerated(bandNRExtAimlBMCase2KnownRxBeamR19Constraints)
			if err != nil {
				return err
			}
			ie.Aiml_BM_Case2_KnownRxBeam_r19 = &v
		}

		if groupSeq.IsComponentPresent(46) {
			v, err := dx.DecodeEnumerated(bandNRExtNtnPowerBoostingERedCapR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_PowerBoosting_ERedCap_r19 = &v
		}

		if groupSeq.IsComponentPresent(47) {
			ie.EarlyCSI_AcquisitionWithCSI_IM_r19 = &struct {
				MaxNumOfCSI_RS_Resource_r19  int64
				MaxNumOfCSI_RS_Ports_r19     int64
				MaxNumOfNZP_CSI_RS_Ports_r19 int64
				MaxRank_r19                  int64
				MaxNumOfCSI_IM_Resource_r19  int64
			}{}
			c := ie.EarlyCSI_AcquisitionWithCSI_IM_r19
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Resource_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtEarlyCSIAcquisitionWithCSIIMR19MaxNumOfNZPCSIRSPortsR19Constraints)
				if err != nil {
					return err
				}
				c.MaxNumOfNZP_CSI_RS_Ports_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxRank_r19 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumOfCSI_IM_Resource_r19 = v
			}
		}

		if groupSeq.IsComponentPresent(48) {
			v, err := dx.DecodeEnumerated(bandNRExtEarlyCSIAcquisitionWithoutCSIIMR19Constraints)
			if err != nil {
				return err
			}
			ie.EarlyCSI_AcquisitionWithoutCSI_IM_r19 = &v
		}
	}

	// Extension group 21:
	if seq.IsExtensionPresent(21) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "enableTx-RxDuringMeasGap-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.EnableTx_RxDuringMeasGap_r19 = &struct {
				AdditionalDCI_r19     per.BitString
				IndicationField_r19   int64
				MinimumTimeOffset_r19 int64
			}{}
			c := ie.EnableTx_RxDuringMeasGap_r19
			{
				v, err := dx.DecodeBitString(per.FixedSize(3))
				if err != nil {
					return err
				}
				c.AdditionalDCI_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtEnableTxRxDuringMeasGapR19IndicationFieldR19Constraints)
				if err != nil {
					return err
				}
				c.IndicationField_r19 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandNRExtEnableTxRxDuringMeasGapR19MinimumTimeOffsetR19Constraints)
				if err != nil {
					return err
				}
				c.MinimumTimeOffset_r19 = v
			}
		}
	}

	return nil
}
