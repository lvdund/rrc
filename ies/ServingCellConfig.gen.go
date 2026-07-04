// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ServingCellConfig (line 14616).

var servingCellConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "tdd-UL-DL-ConfigurationDedicated", Optional: true},
		{Name: "initialDownlinkBWP", Optional: true},
		{Name: "downlinkBWP-ToReleaseList", Optional: true},
		{Name: "downlinkBWP-ToAddModList", Optional: true},
		{Name: "firstActiveDownlinkBWP-Id", Optional: true},
		{Name: "bwp-InactivityTimer", Optional: true},
		{Name: "defaultDownlinkBWP-Id", Optional: true},
		{Name: "uplinkConfig", Optional: true},
		{Name: "supplementaryUplink", Optional: true},
		{Name: "pdcch-ServingCellConfig", Optional: true},
		{Name: "pdsch-ServingCellConfig", Optional: true},
		{Name: "csi-MeasConfig", Optional: true},
		{Name: "sCellDeactivationTimer", Optional: true},
		{Name: "crossCarrierSchedulingConfig", Optional: true},
		{Name: "tag-Id"},
		{Name: "dummy1", Optional: true},
		{Name: "pathlossReferenceLinking", Optional: true},
		{Name: "servingCellMO", Optional: true},
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
	},
}

var servingCellConfigDownlinkBWPToReleaseListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

var servingCellConfigDownlinkBWPToAddModListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

const (
	ServingCellConfig_Bwp_InactivityTimer_Ms2     = 0
	ServingCellConfig_Bwp_InactivityTimer_Ms3     = 1
	ServingCellConfig_Bwp_InactivityTimer_Ms4     = 2
	ServingCellConfig_Bwp_InactivityTimer_Ms5     = 3
	ServingCellConfig_Bwp_InactivityTimer_Ms6     = 4
	ServingCellConfig_Bwp_InactivityTimer_Ms8     = 5
	ServingCellConfig_Bwp_InactivityTimer_Ms10    = 6
	ServingCellConfig_Bwp_InactivityTimer_Ms20    = 7
	ServingCellConfig_Bwp_InactivityTimer_Ms30    = 8
	ServingCellConfig_Bwp_InactivityTimer_Ms40    = 9
	ServingCellConfig_Bwp_InactivityTimer_Ms50    = 10
	ServingCellConfig_Bwp_InactivityTimer_Ms60    = 11
	ServingCellConfig_Bwp_InactivityTimer_Ms80    = 12
	ServingCellConfig_Bwp_InactivityTimer_Ms100   = 13
	ServingCellConfig_Bwp_InactivityTimer_Ms200   = 14
	ServingCellConfig_Bwp_InactivityTimer_Ms300   = 15
	ServingCellConfig_Bwp_InactivityTimer_Ms500   = 16
	ServingCellConfig_Bwp_InactivityTimer_Ms750   = 17
	ServingCellConfig_Bwp_InactivityTimer_Ms1280  = 18
	ServingCellConfig_Bwp_InactivityTimer_Ms1920  = 19
	ServingCellConfig_Bwp_InactivityTimer_Ms2560  = 20
	ServingCellConfig_Bwp_InactivityTimer_Spare10 = 21
	ServingCellConfig_Bwp_InactivityTimer_Spare9  = 22
	ServingCellConfig_Bwp_InactivityTimer_Spare8  = 23
	ServingCellConfig_Bwp_InactivityTimer_Spare7  = 24
	ServingCellConfig_Bwp_InactivityTimer_Spare6  = 25
	ServingCellConfig_Bwp_InactivityTimer_Spare5  = 26
	ServingCellConfig_Bwp_InactivityTimer_Spare4  = 27
	ServingCellConfig_Bwp_InactivityTimer_Spare3  = 28
	ServingCellConfig_Bwp_InactivityTimer_Spare2  = 29
	ServingCellConfig_Bwp_InactivityTimer_Spare1  = 30
)

var servingCellConfigBwpInactivityTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Bwp_InactivityTimer_Ms2, ServingCellConfig_Bwp_InactivityTimer_Ms3, ServingCellConfig_Bwp_InactivityTimer_Ms4, ServingCellConfig_Bwp_InactivityTimer_Ms5, ServingCellConfig_Bwp_InactivityTimer_Ms6, ServingCellConfig_Bwp_InactivityTimer_Ms8, ServingCellConfig_Bwp_InactivityTimer_Ms10, ServingCellConfig_Bwp_InactivityTimer_Ms20, ServingCellConfig_Bwp_InactivityTimer_Ms30, ServingCellConfig_Bwp_InactivityTimer_Ms40, ServingCellConfig_Bwp_InactivityTimer_Ms50, ServingCellConfig_Bwp_InactivityTimer_Ms60, ServingCellConfig_Bwp_InactivityTimer_Ms80, ServingCellConfig_Bwp_InactivityTimer_Ms100, ServingCellConfig_Bwp_InactivityTimer_Ms200, ServingCellConfig_Bwp_InactivityTimer_Ms300, ServingCellConfig_Bwp_InactivityTimer_Ms500, ServingCellConfig_Bwp_InactivityTimer_Ms750, ServingCellConfig_Bwp_InactivityTimer_Ms1280, ServingCellConfig_Bwp_InactivityTimer_Ms1920, ServingCellConfig_Bwp_InactivityTimer_Ms2560, ServingCellConfig_Bwp_InactivityTimer_Spare10, ServingCellConfig_Bwp_InactivityTimer_Spare9, ServingCellConfig_Bwp_InactivityTimer_Spare8, ServingCellConfig_Bwp_InactivityTimer_Spare7, ServingCellConfig_Bwp_InactivityTimer_Spare6, ServingCellConfig_Bwp_InactivityTimer_Spare5, ServingCellConfig_Bwp_InactivityTimer_Spare4, ServingCellConfig_Bwp_InactivityTimer_Spare3, ServingCellConfig_Bwp_InactivityTimer_Spare2, ServingCellConfig_Bwp_InactivityTimer_Spare1},
}

var servingCellConfigPdcchServingCellConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Pdcch_ServingCellConfig_Release = 0
	ServingCellConfig_Pdcch_ServingCellConfig_Setup   = 1
)

var servingCellConfigPdschServingCellConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Pdsch_ServingCellConfig_Release = 0
	ServingCellConfig_Pdsch_ServingCellConfig_Setup   = 1
)

var servingCellConfigCsiMeasConfigConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Csi_MeasConfig_Release = 0
	ServingCellConfig_Csi_MeasConfig_Setup   = 1
)

const (
	ServingCellConfig_SCellDeactivationTimer_Ms20   = 0
	ServingCellConfig_SCellDeactivationTimer_Ms40   = 1
	ServingCellConfig_SCellDeactivationTimer_Ms80   = 2
	ServingCellConfig_SCellDeactivationTimer_Ms160  = 3
	ServingCellConfig_SCellDeactivationTimer_Ms200  = 4
	ServingCellConfig_SCellDeactivationTimer_Ms240  = 5
	ServingCellConfig_SCellDeactivationTimer_Ms320  = 6
	ServingCellConfig_SCellDeactivationTimer_Ms400  = 7
	ServingCellConfig_SCellDeactivationTimer_Ms480  = 8
	ServingCellConfig_SCellDeactivationTimer_Ms520  = 9
	ServingCellConfig_SCellDeactivationTimer_Ms640  = 10
	ServingCellConfig_SCellDeactivationTimer_Ms720  = 11
	ServingCellConfig_SCellDeactivationTimer_Ms840  = 12
	ServingCellConfig_SCellDeactivationTimer_Ms1280 = 13
	ServingCellConfig_SCellDeactivationTimer_Spare2 = 14
	ServingCellConfig_SCellDeactivationTimer_Spare1 = 15
)

var servingCellConfigSCellDeactivationTimerConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_SCellDeactivationTimer_Ms20, ServingCellConfig_SCellDeactivationTimer_Ms40, ServingCellConfig_SCellDeactivationTimer_Ms80, ServingCellConfig_SCellDeactivationTimer_Ms160, ServingCellConfig_SCellDeactivationTimer_Ms200, ServingCellConfig_SCellDeactivationTimer_Ms240, ServingCellConfig_SCellDeactivationTimer_Ms320, ServingCellConfig_SCellDeactivationTimer_Ms400, ServingCellConfig_SCellDeactivationTimer_Ms480, ServingCellConfig_SCellDeactivationTimer_Ms520, ServingCellConfig_SCellDeactivationTimer_Ms640, ServingCellConfig_SCellDeactivationTimer_Ms720, ServingCellConfig_SCellDeactivationTimer_Ms840, ServingCellConfig_SCellDeactivationTimer_Ms1280, ServingCellConfig_SCellDeactivationTimer_Spare2, ServingCellConfig_SCellDeactivationTimer_Spare1},
}

const (
	ServingCellConfig_Dummy1_Enabled = 0
)

var servingCellConfigDummy1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Dummy1_Enabled},
}

const (
	ServingCellConfig_PathlossReferenceLinking_SpCell = 0
	ServingCellConfig_PathlossReferenceLinking_SCell  = 1
)

var servingCellConfigPathlossReferenceLinkingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_PathlossReferenceLinking_SpCell, ServingCellConfig_PathlossReferenceLinking_SCell},
}

var servingCellConfigExtLteCRSToMatchAroundConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Release = 0
	ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Setup   = 1
)

var servingCellConfigExtRateMatchPatternToAddModListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

var servingCellConfigExtRateMatchPatternToReleaseListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

var servingCellConfigExtDownlinkChannelBWPerSCSListConstraints = per.SizeRange(1, common.MaxSCSs)

const (
	ServingCellConfig_Ext_SupplementaryUplinkRelease_r16_True = 0
)

var servingCellConfigExtSupplementaryUplinkReleaseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_SupplementaryUplinkRelease_r16_True},
}

var servingCellConfigExtDormantBWPConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_DormantBWP_Config_r16_Release = 0
	ServingCellConfig_Ext_DormantBWP_Config_r16_Setup   = 1
)

var servingCellConfigExtCaSlotOffsetR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "refSCS15kHz"},
		{Name: "refSCS30KHz"},
		{Name: "refSCS60KHz"},
		{Name: "refSCS120KHz"},
	},
}

const (
	ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS15kHz  = 0
	ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS30KHz  = 1
	ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS60KHz  = 2
	ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS120KHz = 3
)

var servingCellConfigExtDummy2Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Dummy2_Release = 0
	ServingCellConfig_Ext_Dummy2_Setup   = 1
)

var servingCellConfigExtIntraCellGuardBandsDLListR16Constraints = per.SizeRange(1, common.MaxSCSs)

var servingCellConfigExtIntraCellGuardBandsULListR16Constraints = per.SizeRange(1, common.MaxSCSs)

const (
	ServingCellConfig_Ext_Csi_RS_ValidationWithDCI_r16_Enabled = 0
)

var servingCellConfigExtCsiRSValidationWithDCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Csi_RS_ValidationWithDCI_r16_Enabled},
}

var servingCellConfigExtLteCRSPatternList1R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Release = 0
	ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Setup   = 1
)

var servingCellConfigExtLteCRSPatternList2R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Release = 0
	ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Setup   = 1
)

const (
	ServingCellConfig_Ext_Crs_RateMatch_PerCORESETPoolIndex_r16_Enabled = 0
)

var servingCellConfigExtCrsRateMatchPerCORESETPoolIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Crs_RateMatch_PerCORESETPoolIndex_r16_Enabled},
}

const (
	ServingCellConfig_Ext_EnableTwoDefaultTCI_States_r16_Enabled = 0
)

var servingCellConfigExtEnableTwoDefaultTCIStatesR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_EnableTwoDefaultTCI_States_r16_Enabled},
}

const (
	ServingCellConfig_Ext_EnableDefaultTCI_StatePerCoresetPoolIndex_r16_Enabled = 0
)

var servingCellConfigExtEnableDefaultTCIStatePerCoresetPoolIndexR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_EnableDefaultTCI_StatePerCoresetPoolIndex_r16_Enabled},
}

const (
	ServingCellConfig_Ext_EnableBeamSwitchTiming_r16_True = 0
)

var servingCellConfigExtEnableBeamSwitchTimingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_EnableBeamSwitchTiming_r16_True},
}

const (
	ServingCellConfig_Ext_Cbg_TxDiffTBsProcessingType1_r16_Enabled = 0
)

var servingCellConfigExtCbgTxDiffTBsProcessingType1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Cbg_TxDiffTBsProcessingType1_r16_Enabled},
}

const (
	ServingCellConfig_Ext_Cbg_TxDiffTBsProcessingType2_r16_Enabled = 0
)

var servingCellConfigExtCbgTxDiffTBsProcessingType2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Cbg_TxDiffTBsProcessingType2_r16_Enabled},
}

const (
	ServingCellConfig_Ext_DirectionalCollisionHandling_r16_Enabled = 0
)

var servingCellConfigExtDirectionalCollisionHandlingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_DirectionalCollisionHandling_r16_Enabled},
}

var servingCellConfigExtChannelAccessConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_ChannelAccessConfig_r16_Release = 0
	ServingCellConfig_Ext_ChannelAccessConfig_r16_Setup   = 1
)

var servingCellConfigExtNrDlPRSPDCInfoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Release = 0
	ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Setup   = 1
)

var servingCellConfigExtSemiStaticChannelAccessConfigUER17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Release = 0
	ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Setup   = 1
)

var servingCellConfigExtMimoParamR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_MimoParam_r17_Release = 0
	ServingCellConfig_Ext_MimoParam_r17_Setup   = 1
)

const (
	ServingCellConfig_Ext_ChannelAccessMode2_r17_Enabled = 0
)

var servingCellConfigExtChannelAccessMode2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_ChannelAccessMode2_r17_Enabled},
}

const (
	ServingCellConfig_Ext_TimeDomainHARQ_BundlingType1_r17_Enabled = 0
)

var servingCellConfigExtTimeDomainHARQBundlingType1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_TimeDomainHARQ_BundlingType1_r17_Enabled},
}

const (
	ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N1 = 0
	ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N2 = 1
	ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N4 = 2
)

var servingCellConfigExtNrofHARQBundlingGroupsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N1, ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N2, ServingCellConfig_Ext_NrofHARQ_BundlingGroups_r17_N4},
}

const (
	ServingCellConfig_Ext_Fdmed_ReceptionMulticast_r17_True = 0
)

var servingCellConfigExtFdmedReceptionMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Fdmed_ReceptionMulticast_r17_True},
}

const (
	ServingCellConfig_Ext_MoreThanOneNackOnlyMode_r17_Mode2 = 0
)

var servingCellConfigExtMoreThanOneNackOnlyModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_MoreThanOneNackOnlyMode_r17_Mode2},
}

const (
	ServingCellConfig_Ext_DirectionalCollisionHandling_DC_r17_Enabled = 0
)

var servingCellConfigExtDirectionalCollisionHandlingDCR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_DirectionalCollisionHandling_DC_r17_Enabled},
}

var servingCellConfigExtLteNeighCellsCRSAssistInfoListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Release = 0
	ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Setup   = 1
)

const (
	ServingCellConfig_Ext_Lte_NeighCellsCRS_Assumptions_r17_False = 0
)

var servingCellConfigExtLteNeighCellsCRSAssumptionsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Lte_NeighCellsCRS_Assumptions_r17_False},
}

const (
	ServingCellConfig_Ext_CrossCarrierSchedulingConfigRelease_r17_True = 0
)

var servingCellConfigExtCrossCarrierSchedulingConfigReleaseR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_CrossCarrierSchedulingConfigRelease_r17_True},
}

const (
	ServingCellConfig_Ext_MultiPDSCH_PerSlotType1_CB_r17_Enabled  = 0
	ServingCellConfig_Ext_MultiPDSCH_PerSlotType1_CB_r17_Disabled = 1
)

var servingCellConfigExtMultiPDSCHPerSlotType1CBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_MultiPDSCH_PerSlotType1_CB_r17_Enabled, ServingCellConfig_Ext_MultiPDSCH_PerSlotType1_CB_r17_Disabled},
}

var servingCellConfigExtLteCRSPatternList3R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Release = 0
	ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Setup   = 1
)

var servingCellConfigExtLteCRSPatternList4R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Release = 0
	ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Setup   = 1
)

const (
	ServingCellConfig_Ext_Pdcch_CandidateReceptionWithCRS_Overlap_r18_Enabled = 0
)

var servingCellConfigExtPdcchCandidateReceptionWithCRSOverlapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Pdcch_CandidateReceptionWithCRS_Overlap_r18_Enabled},
}

const (
	ServingCellConfig_Ext_Cjt_Scheme_PDSCH_r18_CjtSchemeA = 0
	ServingCellConfig_Ext_Cjt_Scheme_PDSCH_r18_CjtSchemeB = 1
)

var servingCellConfigExtCjtSchemePDSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Cjt_Scheme_PDSCH_r18_CjtSchemeA, ServingCellConfig_Ext_Cjt_Scheme_PDSCH_r18_CjtSchemeB},
}

var servingCellConfigExtCellDTXDRXConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Release = 0
	ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Setup   = 1
)

var servingCellConfigPositionInDCICellDTRXR18Constraints = per.Constrained(0, common.MaxDCI_2_9_Size_1_r18)

const (
	ServingCellConfig_Ext_CellDTX_DRX_L1activation_r18_Enabled = 0
)

var servingCellConfigExtCellDTXDRXL1activationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_CellDTX_DRX_L1activation_r18_Enabled},
}

var servingCellConfigExtMcDCISetOfCellsToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofSetsOfCells_r18)

var servingCellConfigExtMcDCISetOfCellsToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofSetsOfCells_r18)

var servingCellConfigExtMimoParamV1850Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_MimoParam_v1850_Release = 0
	ServingCellConfig_Ext_MimoParam_v1850_Setup   = 1
)

var servingCellConfigExtMcDCISetOfCellsToAddModListExtV1900Constraints = per.SizeRange(1, common.MaxNrofSetsOfCells_r18)

var servingCellConfigExtMimoParamV1900Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	ServingCellConfig_Ext_MimoParam_v1900_Release = 0
	ServingCellConfig_Ext_MimoParam_v1900_Setup   = 1
)

const (
	ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeC = 0
	ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeD = 1
	ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeE = 2
)

var servingCellConfigExtCjtSchemePDSCHV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeC, ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeD, ServingCellConfig_Ext_Cjt_Scheme_PDSCH_v1900_Cjt_SchemeE},
}

const (
	ServingCellConfig_Ext_Ntn_RedcapPrioritizeUL_Dynamic_r19_Enabled = 0
)

var servingCellConfigExtNtnRedcapPrioritizeULDynamicR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Ntn_RedcapPrioritizeUL_Dynamic_r19_Enabled},
}

const (
	ServingCellConfig_Ext_Ntn_RedcapPrioritizeUL_Semistatic_r19_Enabled = 0
)

var servingCellConfigExtNtnRedcapPrioritizeULSemistaticR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ServingCellConfig_Ext_Ntn_RedcapPrioritizeUL_Semistatic_r19_Enabled},
}

type ServingCellConfig struct {
	Tdd_UL_DL_ConfigurationDedicated *TDD_UL_DL_ConfigDedicated
	InitialDownlinkBWP               *BWP_DownlinkDedicated
	DownlinkBWP_ToReleaseList        []BWP_Id
	DownlinkBWP_ToAddModList         []BWP_Downlink
	FirstActiveDownlinkBWP_Id        *BWP_Id
	Bwp_InactivityTimer              *int64
	DefaultDownlinkBWP_Id            *BWP_Id
	UplinkConfig                     *UplinkConfig
	SupplementaryUplink              *UplinkConfig
	Pdcch_ServingCellConfig          *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_ServingCellConfig
	}
	Pdsch_ServingCellConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_ServingCellConfig
	}
	Csi_MeasConfig *struct {
		Choice  int
		Release *struct{}
		Setup   *CSI_MeasConfig
	}
	SCellDeactivationTimer       *int64
	CrossCarrierSchedulingConfig *CrossCarrierSchedulingConfig
	Tag_Id                       TAG_Id
	Dummy1                       *int64
	PathlossReferenceLinking     *int64
	ServingCellMO                *MeasObjectId
	Lte_CRS_ToMatchAround        *struct {
		Choice  int
		Release *struct{}
		Setup   *RateMatchPatternLTE_CRS
	}
	RateMatchPatternToAddModList                []RateMatchPattern
	RateMatchPatternToReleaseList               []RateMatchPatternId
	DownlinkChannelBW_PerSCS_List               []SCS_SpecificCarrier
	SupplementaryUplinkRelease_r16              *int64
	Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 *TDD_UL_DL_ConfigDedicated_IAB_MT_r16
	DormantBWP_Config_r16                       *struct {
		Choice  int
		Release *struct{}
		Setup   *DormantBWP_Config_r16
	}
	Ca_SlotOffset_r16 *struct {
		Choice       int
		RefSCS15kHz  *int64
		RefSCS30KHz  *int64
		RefSCS60KHz  *int64
		RefSCS120KHz *int64
	}
	Dummy2 *struct {
		Choice  int
		Release *struct{}
		Setup   *DummyJ
	}
	IntraCellGuardBandsDL_List_r16 []IntraCellGuardBandsPerSCS_r16
	IntraCellGuardBandsUL_List_r16 []IntraCellGuardBandsPerSCS_r16
	Csi_RS_ValidationWithDCI_r16   *int64
	Lte_CRS_PatternList1_r16       *struct {
		Choice  int
		Release *struct{}
		Setup   *LTE_CRS_PatternList_r16
	}
	Lte_CRS_PatternList2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *LTE_CRS_PatternList_r16
	}
	Crs_RateMatch_PerCORESETPoolIndex_r16         *int64
	EnableTwoDefaultTCI_States_r16                *int64
	EnableDefaultTCI_StatePerCoresetPoolIndex_r16 *int64
	EnableBeamSwitchTiming_r16                    *int64
	Cbg_TxDiffTBsProcessingType1_r16              *int64
	Cbg_TxDiffTBsProcessingType2_r16              *int64
	DirectionalCollisionHandling_r16              *int64
	ChannelAccessConfig_r16                       *struct {
		Choice  int
		Release *struct{}
		Setup   *ChannelAccessConfig_r16
	}
	Nr_Dl_PRS_PDC_Info_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *NR_DL_PRS_PDC_Info_r17
	}
	SemiStaticChannelAccessConfigUE_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *SemiStaticChannelAccessConfigUE_r17
	}
	MimoParam_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MIMOParam_r17
	}
	ChannelAccessMode2_r17               *int64
	TimeDomainHARQ_BundlingType1_r17     *int64
	NrofHARQ_BundlingGroups_r17          *int64
	Fdmed_ReceptionMulticast_r17         *int64
	MoreThanOneNackOnlyMode_r17          *int64
	Tci_ActivatedConfig_r17              *TCI_ActivatedConfig_r17
	DirectionalCollisionHandling_DC_r17  *int64
	Lte_NeighCellsCRS_AssistInfoList_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *LTE_NeighCellsCRS_AssistInfoList_r17
	}
	Lte_NeighCellsCRS_Assumptions_r17       *int64
	CrossCarrierSchedulingConfigRelease_r17 *int64
	MultiPDSCH_PerSlotType1_CB_r17          *int64
	Lte_CRS_PatternList3_r18                *struct {
		Choice  int
		Release *struct{}
		Setup   *LTE_CRS_PatternList_r16
	}
	Lte_CRS_PatternList4_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *LTE_CRS_PatternList_r16
	}
	Pdcch_CandidateReceptionWithCRS_Overlap_r18 *int64
	Cjt_Scheme_PDSCH_r18                        *int64
	Tag2_r18                                    *Tag2_r18
	CellDTX_DRX_Config_r18                      *struct {
		Choice  int
		Release *struct{}
		Setup   *CellDTX_DRX_Config_r18
	}
	PositionInDCI_CellDTRX_r18         *int64
	CellDTX_DRX_L1activation_r18       *int64
	Mc_DCI_SetOfCellsToAddModList_r18  []MC_DCI_SetOfCells_r18
	Mc_DCI_SetOfCellsToReleaseList_r18 []SetOfCellsId_r18
	MimoParam_v1850                    *struct {
		Choice  int
		Release *struct{}
		Setup   *MIMOParam_v1850
	}
	Mc_DCI_SetOfCellsToAddModListExt_v1900 []MC_DCI_SetOfCellsExt_v1900
	MimoParam_v1900                        *struct {
		Choice  int
		Release *struct{}
		Setup   *MIMOParam_v1900
	}
	ServingCellMO_OD_r19                  *MeasObjectId
	Cjt_Scheme_PDSCH_v1900                *int64
	Ntn_RedcapPrioritizeUL_Dynamic_r19    *int64
	Ntn_RedcapPrioritizeUL_Semistatic_r19 *int64
}

func (ie *ServingCellConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(servingCellConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Lte_CRS_ToMatchAround != nil || ie.RateMatchPatternToAddModList != nil || ie.RateMatchPatternToReleaseList != nil || ie.DownlinkChannelBW_PerSCS_List != nil
	hasExtGroup1 := ie.SupplementaryUplinkRelease_r16 != nil || ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil || ie.DormantBWP_Config_r16 != nil || ie.Ca_SlotOffset_r16 != nil || ie.Dummy2 != nil || ie.IntraCellGuardBandsDL_List_r16 != nil || ie.IntraCellGuardBandsUL_List_r16 != nil || ie.Csi_RS_ValidationWithDCI_r16 != nil || ie.Lte_CRS_PatternList1_r16 != nil || ie.Lte_CRS_PatternList2_r16 != nil || ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil || ie.EnableTwoDefaultTCI_States_r16 != nil || ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil || ie.EnableBeamSwitchTiming_r16 != nil || ie.Cbg_TxDiffTBsProcessingType1_r16 != nil || ie.Cbg_TxDiffTBsProcessingType2_r16 != nil
	hasExtGroup2 := ie.DirectionalCollisionHandling_r16 != nil || ie.ChannelAccessConfig_r16 != nil
	hasExtGroup3 := ie.Nr_Dl_PRS_PDC_Info_r17 != nil || ie.SemiStaticChannelAccessConfigUE_r17 != nil || ie.MimoParam_r17 != nil || ie.ChannelAccessMode2_r17 != nil || ie.TimeDomainHARQ_BundlingType1_r17 != nil || ie.NrofHARQ_BundlingGroups_r17 != nil || ie.Fdmed_ReceptionMulticast_r17 != nil || ie.MoreThanOneNackOnlyMode_r17 != nil || ie.Tci_ActivatedConfig_r17 != nil || ie.DirectionalCollisionHandling_DC_r17 != nil || ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil
	hasExtGroup4 := ie.Lte_NeighCellsCRS_Assumptions_r17 != nil
	hasExtGroup5 := ie.CrossCarrierSchedulingConfigRelease_r17 != nil
	hasExtGroup6 := ie.MultiPDSCH_PerSlotType1_CB_r17 != nil
	hasExtGroup7 := ie.Lte_CRS_PatternList3_r18 != nil || ie.Lte_CRS_PatternList4_r18 != nil || ie.Pdcch_CandidateReceptionWithCRS_Overlap_r18 != nil || ie.Cjt_Scheme_PDSCH_r18 != nil || ie.Tag2_r18 != nil || ie.CellDTX_DRX_Config_r18 != nil || ie.PositionInDCI_CellDTRX_r18 != nil || ie.CellDTX_DRX_L1activation_r18 != nil || ie.Mc_DCI_SetOfCellsToAddModList_r18 != nil || ie.Mc_DCI_SetOfCellsToReleaseList_r18 != nil
	hasExtGroup8 := ie.MimoParam_v1850 != nil
	hasExtGroup9 := ie.Mc_DCI_SetOfCellsToAddModListExt_v1900 != nil || ie.MimoParam_v1900 != nil || ie.ServingCellMO_OD_r19 != nil || ie.Cjt_Scheme_PDSCH_v1900 != nil || ie.Ntn_RedcapPrioritizeUL_Dynamic_r19 != nil || ie.Ntn_RedcapPrioritizeUL_Semistatic_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Tdd_UL_DL_ConfigurationDedicated != nil, ie.InitialDownlinkBWP != nil, ie.DownlinkBWP_ToReleaseList != nil, ie.DownlinkBWP_ToAddModList != nil, ie.FirstActiveDownlinkBWP_Id != nil, ie.Bwp_InactivityTimer != nil, ie.DefaultDownlinkBWP_Id != nil, ie.UplinkConfig != nil, ie.SupplementaryUplink != nil, ie.Pdcch_ServingCellConfig != nil, ie.Pdsch_ServingCellConfig != nil, ie.Csi_MeasConfig != nil, ie.SCellDeactivationTimer != nil, ie.CrossCarrierSchedulingConfig != nil, ie.Dummy1 != nil, ie.PathlossReferenceLinking != nil, ie.ServingCellMO != nil}); err != nil {
		return err
	}

	// 3. tdd-UL-DL-ConfigurationDedicated: ref
	{
		if ie.Tdd_UL_DL_ConfigurationDedicated != nil {
			if err := ie.Tdd_UL_DL_ConfigurationDedicated.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. initialDownlinkBWP: ref
	{
		if ie.InitialDownlinkBWP != nil {
			if err := ie.InitialDownlinkBWP.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. downlinkBWP-ToReleaseList: sequence-of
	{
		if ie.DownlinkBWP_ToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(servingCellConfigDownlinkBWPToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.DownlinkBWP_ToReleaseList))); err != nil {
				return err
			}
			for i := range ie.DownlinkBWP_ToReleaseList {
				if err := ie.DownlinkBWP_ToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. downlinkBWP-ToAddModList: sequence-of
	{
		if ie.DownlinkBWP_ToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(servingCellConfigDownlinkBWPToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.DownlinkBWP_ToAddModList))); err != nil {
				return err
			}
			for i := range ie.DownlinkBWP_ToAddModList {
				if err := ie.DownlinkBWP_ToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. firstActiveDownlinkBWP-Id: ref
	{
		if ie.FirstActiveDownlinkBWP_Id != nil {
			if err := ie.FirstActiveDownlinkBWP_Id.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. bwp-InactivityTimer: enumerated
	{
		if ie.Bwp_InactivityTimer != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_InactivityTimer, servingCellConfigBwpInactivityTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 9. defaultDownlinkBWP-Id: ref
	{
		if ie.DefaultDownlinkBWP_Id != nil {
			if err := ie.DefaultDownlinkBWP_Id.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. uplinkConfig: ref
	{
		if ie.UplinkConfig != nil {
			if err := ie.UplinkConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. supplementaryUplink: ref
	{
		if ie.SupplementaryUplink != nil {
			if err := ie.SupplementaryUplink.Encode(e); err != nil {
				return err
			}
		}
	}

	// 12. pdcch-ServingCellConfig: choice
	{
		if ie.Pdcch_ServingCellConfig != nil {
			choiceEnc := e.NewChoiceEncoder(servingCellConfigPdcchServingCellConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_ServingCellConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdcch_ServingCellConfig).Choice {
			case ServingCellConfig_Pdcch_ServingCellConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Pdcch_ServingCellConfig_Setup:
				if err := (*ie.Pdcch_ServingCellConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdcch_ServingCellConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 13. pdsch-ServingCellConfig: choice
	{
		if ie.Pdsch_ServingCellConfig != nil {
			choiceEnc := e.NewChoiceEncoder(servingCellConfigPdschServingCellConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_ServingCellConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_ServingCellConfig).Choice {
			case ServingCellConfig_Pdsch_ServingCellConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Pdsch_ServingCellConfig_Setup:
				if err := (*ie.Pdsch_ServingCellConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_ServingCellConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 14. csi-MeasConfig: choice
	{
		if ie.Csi_MeasConfig != nil {
			choiceEnc := e.NewChoiceEncoder(servingCellConfigCsiMeasConfigConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Csi_MeasConfig).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Csi_MeasConfig).Choice {
			case ServingCellConfig_Csi_MeasConfig_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Csi_MeasConfig_Setup:
				if err := (*ie.Csi_MeasConfig).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Csi_MeasConfig).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 15. sCellDeactivationTimer: enumerated
	{
		if ie.SCellDeactivationTimer != nil {
			if err := e.EncodeEnumerated(*ie.SCellDeactivationTimer, servingCellConfigSCellDeactivationTimerConstraints); err != nil {
				return err
			}
		}
	}

	// 16. crossCarrierSchedulingConfig: ref
	{
		if ie.CrossCarrierSchedulingConfig != nil {
			if err := ie.CrossCarrierSchedulingConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. tag-Id: ref
	{
		if err := ie.Tag_Id.Encode(e); err != nil {
			return err
		}
	}

	// 18. dummy1: enumerated
	{
		if ie.Dummy1 != nil {
			if err := e.EncodeEnumerated(*ie.Dummy1, servingCellConfigDummy1Constraints); err != nil {
				return err
			}
		}
	}

	// 19. pathlossReferenceLinking: enumerated
	{
		if ie.PathlossReferenceLinking != nil {
			if err := e.EncodeEnumerated(*ie.PathlossReferenceLinking, servingCellConfigPathlossReferenceLinkingConstraints); err != nil {
				return err
			}
		}
	}

	// 20. servingCellMO: ref
	{
		if ie.ServingCellMO != nil {
			if err := ie.ServingCellMO.Encode(e); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "lte-CRS-ToMatchAround", Optional: true},
					{Name: "rateMatchPatternToAddModList", Optional: true},
					{Name: "rateMatchPatternToReleaseList", Optional: true},
					{Name: "downlinkChannelBW-PerSCS-List", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Lte_CRS_ToMatchAround != nil, ie.RateMatchPatternToAddModList != nil, ie.RateMatchPatternToReleaseList != nil, ie.DownlinkChannelBW_PerSCS_List != nil}); err != nil {
				return err
			}

			if ie.Lte_CRS_ToMatchAround != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteCRSToMatchAroundConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_ToMatchAround).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_CRS_ToMatchAround).Choice {
				case ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Setup:
					if err := (*ie.Lte_CRS_ToMatchAround).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.RateMatchPatternToAddModList != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtRateMatchPatternToAddModListConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToAddModList))); err != nil {
					return err
				}
				for i := range ie.RateMatchPatternToAddModList {
					if err := ie.RateMatchPatternToAddModList[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.RateMatchPatternToReleaseList != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtRateMatchPatternToReleaseListConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToReleaseList))); err != nil {
					return err
				}
				for i := range ie.RateMatchPatternToReleaseList {
					if err := ie.RateMatchPatternToReleaseList[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.DownlinkChannelBW_PerSCS_List != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtDownlinkChannelBWPerSCSListConstraints)
				if err := seqOf.EncodeLength(int64(len(ie.DownlinkChannelBW_PerSCS_List))); err != nil {
					return err
				}
				for i := range ie.DownlinkChannelBW_PerSCS_List {
					if err := ie.DownlinkChannelBW_PerSCS_List[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "supplementaryUplinkRelease-r16", Optional: true},
					{Name: "tdd-UL-DL-ConfigurationDedicated-IAB-MT-r16", Optional: true},
					{Name: "dormantBWP-Config-r16", Optional: true},
					{Name: "ca-SlotOffset-r16", Optional: true},
					{Name: "dummy2", Optional: true},
					{Name: "intraCellGuardBandsDL-List-r16", Optional: true},
					{Name: "intraCellGuardBandsUL-List-r16", Optional: true},
					{Name: "csi-RS-ValidationWithDCI-r16", Optional: true},
					{Name: "lte-CRS-PatternList1-r16", Optional: true},
					{Name: "lte-CRS-PatternList2-r16", Optional: true},
					{Name: "crs-RateMatch-PerCORESETPoolIndex-r16", Optional: true},
					{Name: "enableTwoDefaultTCI-States-r16", Optional: true},
					{Name: "enableDefaultTCI-StatePerCoresetPoolIndex-r16", Optional: true},
					{Name: "enableBeamSwitchTiming-r16", Optional: true},
					{Name: "cbg-TxDiffTBsProcessingType1-r16", Optional: true},
					{Name: "cbg-TxDiffTBsProcessingType2-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SupplementaryUplinkRelease_r16 != nil, ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil, ie.DormantBWP_Config_r16 != nil, ie.Ca_SlotOffset_r16 != nil, ie.Dummy2 != nil, ie.IntraCellGuardBandsDL_List_r16 != nil, ie.IntraCellGuardBandsUL_List_r16 != nil, ie.Csi_RS_ValidationWithDCI_r16 != nil, ie.Lte_CRS_PatternList1_r16 != nil, ie.Lte_CRS_PatternList2_r16 != nil, ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil, ie.EnableTwoDefaultTCI_States_r16 != nil, ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil, ie.EnableBeamSwitchTiming_r16 != nil, ie.Cbg_TxDiffTBsProcessingType1_r16 != nil, ie.Cbg_TxDiffTBsProcessingType2_r16 != nil}); err != nil {
				return err
			}

			if ie.SupplementaryUplinkRelease_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SupplementaryUplinkRelease_r16, servingCellConfigExtSupplementaryUplinkReleaseR16Constraints); err != nil {
					return err
				}
			}

			if ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 != nil {
				if err := ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.DormantBWP_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtDormantBWPConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.DormantBWP_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.DormantBWP_Config_r16).Choice {
				case ServingCellConfig_Ext_DormantBWP_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_DormantBWP_Config_r16_Setup:
					if err := (*ie.DormantBWP_Config_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ca_SlotOffset_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtCaSlotOffsetR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ca_SlotOffset_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ca_SlotOffset_r16).Choice {
				case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS15kHz:
					if err := ex.EncodeInteger((*(*ie.Ca_SlotOffset_r16).RefSCS15kHz), per.Constrained(-2, 2)); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS30KHz:
					if err := ex.EncodeInteger((*(*ie.Ca_SlotOffset_r16).RefSCS30KHz), per.Constrained(-5, 5)); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS60KHz:
					if err := ex.EncodeInteger((*(*ie.Ca_SlotOffset_r16).RefSCS60KHz), per.Constrained(-10, 10)); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS120KHz:
					if err := ex.EncodeInteger((*(*ie.Ca_SlotOffset_r16).RefSCS120KHz), per.Constrained(-20, 20)); err != nil {
						return err
					}
				}
			}

			if ie.Dummy2 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtDummy2Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dummy2).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dummy2).Choice {
				case ServingCellConfig_Ext_Dummy2_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Dummy2_Setup:
					if err := (*ie.Dummy2).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IntraCellGuardBandsDL_List_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtIntraCellGuardBandsDLListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.IntraCellGuardBandsDL_List_r16))); err != nil {
					return err
				}
				for i := range ie.IntraCellGuardBandsDL_List_r16 {
					if err := ie.IntraCellGuardBandsDL_List_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.IntraCellGuardBandsUL_List_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtIntraCellGuardBandsULListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.IntraCellGuardBandsUL_List_r16))); err != nil {
					return err
				}
				for i := range ie.IntraCellGuardBandsUL_List_r16 {
					if err := ie.IntraCellGuardBandsUL_List_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Csi_RS_ValidationWithDCI_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Csi_RS_ValidationWithDCI_r16, servingCellConfigExtCsiRSValidationWithDCIR16Constraints); err != nil {
					return err
				}
			}

			if ie.Lte_CRS_PatternList1_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteCRSPatternList1R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_PatternList1_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_CRS_PatternList1_r16).Choice {
				case ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Setup:
					if err := (*ie.Lte_CRS_PatternList1_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Lte_CRS_PatternList2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteCRSPatternList2R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_PatternList2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_CRS_PatternList2_r16).Choice {
				case ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Setup:
					if err := (*ie.Lte_CRS_PatternList2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Crs_RateMatch_PerCORESETPoolIndex_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Crs_RateMatch_PerCORESETPoolIndex_r16, servingCellConfigExtCrsRateMatchPerCORESETPoolIndexR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableTwoDefaultTCI_States_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableTwoDefaultTCI_States_r16, servingCellConfigExtEnableTwoDefaultTCIStatesR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16, servingCellConfigExtEnableDefaultTCIStatePerCoresetPoolIndexR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnableBeamSwitchTiming_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableBeamSwitchTiming_r16, servingCellConfigExtEnableBeamSwitchTimingR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cbg_TxDiffTBsProcessingType1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cbg_TxDiffTBsProcessingType1_r16, servingCellConfigExtCbgTxDiffTBsProcessingType1R16Constraints); err != nil {
					return err
				}
			}

			if ie.Cbg_TxDiffTBsProcessingType2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cbg_TxDiffTBsProcessingType2_r16, servingCellConfigExtCbgTxDiffTBsProcessingType2R16Constraints); err != nil {
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
					{Name: "directionalCollisionHandling-r16", Optional: true},
					{Name: "channelAccessConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.DirectionalCollisionHandling_r16 != nil, ie.ChannelAccessConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.DirectionalCollisionHandling_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DirectionalCollisionHandling_r16, servingCellConfigExtDirectionalCollisionHandlingR16Constraints); err != nil {
					return err
				}
			}

			if ie.ChannelAccessConfig_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtChannelAccessConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ChannelAccessConfig_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ChannelAccessConfig_r16).Choice {
				case ServingCellConfig_Ext_ChannelAccessConfig_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_ChannelAccessConfig_r16_Setup:
					if err := (*ie.ChannelAccessConfig_r16).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "nr-dl-PRS-PDC-Info-r17", Optional: true},
					{Name: "semiStaticChannelAccessConfigUE-r17", Optional: true},
					{Name: "mimoParam-r17", Optional: true},
					{Name: "channelAccessMode2-r17", Optional: true},
					{Name: "timeDomainHARQ-BundlingType1-r17", Optional: true},
					{Name: "nrofHARQ-BundlingGroups-r17", Optional: true},
					{Name: "fdmed-ReceptionMulticast-r17", Optional: true},
					{Name: "moreThanOneNackOnlyMode-r17", Optional: true},
					{Name: "tci-ActivatedConfig-r17", Optional: true},
					{Name: "directionalCollisionHandling-DC-r17", Optional: true},
					{Name: "lte-NeighCellsCRS-AssistInfoList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nr_Dl_PRS_PDC_Info_r17 != nil, ie.SemiStaticChannelAccessConfigUE_r17 != nil, ie.MimoParam_r17 != nil, ie.ChannelAccessMode2_r17 != nil, ie.TimeDomainHARQ_BundlingType1_r17 != nil, ie.NrofHARQ_BundlingGroups_r17 != nil, ie.Fdmed_ReceptionMulticast_r17 != nil, ie.MoreThanOneNackOnlyMode_r17 != nil, ie.Tci_ActivatedConfig_r17 != nil, ie.DirectionalCollisionHandling_DC_r17 != nil, ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil}); err != nil {
				return err
			}

			if ie.Nr_Dl_PRS_PDC_Info_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtNrDlPRSPDCInfoR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Nr_Dl_PRS_PDC_Info_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Nr_Dl_PRS_PDC_Info_r17).Choice {
				case ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Setup:
					if err := (*ie.Nr_Dl_PRS_PDC_Info_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SemiStaticChannelAccessConfigUE_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtSemiStaticChannelAccessConfigUER17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.SemiStaticChannelAccessConfigUE_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.SemiStaticChannelAccessConfigUE_r17).Choice {
				case ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Setup:
					if err := (*ie.SemiStaticChannelAccessConfigUE_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MimoParam_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtMimoParamR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MimoParam_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MimoParam_r17).Choice {
				case ServingCellConfig_Ext_MimoParam_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_MimoParam_r17_Setup:
					if err := (*ie.MimoParam_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ChannelAccessMode2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ChannelAccessMode2_r17, servingCellConfigExtChannelAccessMode2R17Constraints); err != nil {
					return err
				}
			}

			if ie.TimeDomainHARQ_BundlingType1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TimeDomainHARQ_BundlingType1_r17, servingCellConfigExtTimeDomainHARQBundlingType1R17Constraints); err != nil {
					return err
				}
			}

			if ie.NrofHARQ_BundlingGroups_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofHARQ_BundlingGroups_r17, servingCellConfigExtNrofHARQBundlingGroupsR17Constraints); err != nil {
					return err
				}
			}

			if ie.Fdmed_ReceptionMulticast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Fdmed_ReceptionMulticast_r17, servingCellConfigExtFdmedReceptionMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.MoreThanOneNackOnlyMode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MoreThanOneNackOnlyMode_r17, servingCellConfigExtMoreThanOneNackOnlyModeR17Constraints); err != nil {
					return err
				}
			}

			if ie.Tci_ActivatedConfig_r17 != nil {
				if err := ie.Tci_ActivatedConfig_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.DirectionalCollisionHandling_DC_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.DirectionalCollisionHandling_DC_r17, servingCellConfigExtDirectionalCollisionHandlingDCR17Constraints); err != nil {
					return err
				}
			}

			if ie.Lte_NeighCellsCRS_AssistInfoList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteNeighCellsCRSAssistInfoListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Choice {
				case ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Setup:
					if err := (*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Setup.Encode(ex); err != nil {
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
					{Name: "lte-NeighCellsCRS-Assumptions-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Lte_NeighCellsCRS_Assumptions_r17 != nil}); err != nil {
				return err
			}

			if ie.Lte_NeighCellsCRS_Assumptions_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Lte_NeighCellsCRS_Assumptions_r17, servingCellConfigExtLteNeighCellsCRSAssumptionsR17Constraints); err != nil {
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
					{Name: "crossCarrierSchedulingConfigRelease-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CrossCarrierSchedulingConfigRelease_r17 != nil}); err != nil {
				return err
			}

			if ie.CrossCarrierSchedulingConfigRelease_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.CrossCarrierSchedulingConfigRelease_r17, servingCellConfigExtCrossCarrierSchedulingConfigReleaseR17Constraints); err != nil {
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
					{Name: "multiPDSCH-PerSlotType1-CB-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultiPDSCH_PerSlotType1_CB_r17 != nil}); err != nil {
				return err
			}

			if ie.MultiPDSCH_PerSlotType1_CB_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultiPDSCH_PerSlotType1_CB_r17, servingCellConfigExtMultiPDSCHPerSlotType1CBR17Constraints); err != nil {
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
					{Name: "lte-CRS-PatternList3-r18", Optional: true},
					{Name: "lte-CRS-PatternList4-r18", Optional: true},
					{Name: "pdcch-CandidateReceptionWithCRS-Overlap-r18", Optional: true},
					{Name: "cjt-Scheme-PDSCH-r18", Optional: true},
					{Name: "tag2-r18", Optional: true},
					{Name: "cellDTX-DRX-Config-r18", Optional: true},
					{Name: "positionInDCI-cellDTRX-r18", Optional: true},
					{Name: "cellDTX-DRX-L1activation-r18", Optional: true},
					{Name: "mc-DCI-SetOfCellsToAddModList-r18", Optional: true},
					{Name: "mc-DCI-SetOfCellsToReleaseList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Lte_CRS_PatternList3_r18 != nil, ie.Lte_CRS_PatternList4_r18 != nil, ie.Pdcch_CandidateReceptionWithCRS_Overlap_r18 != nil, ie.Cjt_Scheme_PDSCH_r18 != nil, ie.Tag2_r18 != nil, ie.CellDTX_DRX_Config_r18 != nil, ie.PositionInDCI_CellDTRX_r18 != nil, ie.CellDTX_DRX_L1activation_r18 != nil, ie.Mc_DCI_SetOfCellsToAddModList_r18 != nil, ie.Mc_DCI_SetOfCellsToReleaseList_r18 != nil}); err != nil {
				return err
			}

			if ie.Lte_CRS_PatternList3_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteCRSPatternList3R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_PatternList3_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_CRS_PatternList3_r18).Choice {
				case ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Setup:
					if err := (*ie.Lte_CRS_PatternList3_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Lte_CRS_PatternList4_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtLteCRSPatternList4R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lte_CRS_PatternList4_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lte_CRS_PatternList4_r18).Choice {
				case ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Setup:
					if err := (*ie.Lte_CRS_PatternList4_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdcch_CandidateReceptionWithCRS_Overlap_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdcch_CandidateReceptionWithCRS_Overlap_r18, servingCellConfigExtPdcchCandidateReceptionWithCRSOverlapR18Constraints); err != nil {
					return err
				}
			}

			if ie.Cjt_Scheme_PDSCH_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjt_Scheme_PDSCH_r18, servingCellConfigExtCjtSchemePDSCHR18Constraints); err != nil {
					return err
				}
			}

			if ie.Tag2_r18 != nil {
				if err := ie.Tag2_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CellDTX_DRX_Config_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtCellDTXDRXConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.CellDTX_DRX_Config_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.CellDTX_DRX_Config_r18).Choice {
				case ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Setup:
					if err := (*ie.CellDTX_DRX_Config_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.PositionInDCI_CellDTRX_r18 != nil {
				if err := ex.EncodeInteger(*ie.PositionInDCI_CellDTRX_r18, servingCellConfigPositionInDCICellDTRXR18Constraints); err != nil {
					return err
				}
			}

			if ie.CellDTX_DRX_L1activation_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.CellDTX_DRX_L1activation_r18, servingCellConfigExtCellDTXDRXL1activationR18Constraints); err != nil {
					return err
				}
			}

			if ie.Mc_DCI_SetOfCellsToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtMcDCISetOfCellsToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mc_DCI_SetOfCellsToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.Mc_DCI_SetOfCellsToAddModList_r18 {
					if err := ie.Mc_DCI_SetOfCellsToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Mc_DCI_SetOfCellsToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtMcDCISetOfCellsToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mc_DCI_SetOfCellsToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.Mc_DCI_SetOfCellsToReleaseList_r18 {
					if err := ie.Mc_DCI_SetOfCellsToReleaseList_r18[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "mimoParam-v1850", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MimoParam_v1850 != nil}); err != nil {
				return err
			}

			if ie.MimoParam_v1850 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtMimoParamV1850Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MimoParam_v1850).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MimoParam_v1850).Choice {
				case ServingCellConfig_Ext_MimoParam_v1850_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_MimoParam_v1850_Setup:
					if err := (*ie.MimoParam_v1850).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "mc-DCI-SetOfCellsToAddModListExt-v1900", Optional: true},
					{Name: "mimoParam-v1900", Optional: true},
					{Name: "servingCellMO-OD-r19", Optional: true},
					{Name: "cjt-Scheme-PDSCH-v1900", Optional: true},
					{Name: "ntn-RedcapPrioritizeUL-Dynamic-r19", Optional: true},
					{Name: "ntn-RedcapPrioritizeUL-Semistatic-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mc_DCI_SetOfCellsToAddModListExt_v1900 != nil, ie.MimoParam_v1900 != nil, ie.ServingCellMO_OD_r19 != nil, ie.Cjt_Scheme_PDSCH_v1900 != nil, ie.Ntn_RedcapPrioritizeUL_Dynamic_r19 != nil, ie.Ntn_RedcapPrioritizeUL_Semistatic_r19 != nil}); err != nil {
				return err
			}

			if ie.Mc_DCI_SetOfCellsToAddModListExt_v1900 != nil {
				seqOf := ex.NewSequenceOfEncoder(servingCellConfigExtMcDCISetOfCellsToAddModListExtV1900Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mc_DCI_SetOfCellsToAddModListExt_v1900))); err != nil {
					return err
				}
				for i := range ie.Mc_DCI_SetOfCellsToAddModListExt_v1900 {
					if err := ie.Mc_DCI_SetOfCellsToAddModListExt_v1900[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MimoParam_v1900 != nil {
				choiceEnc := ex.NewChoiceEncoder(servingCellConfigExtMimoParamV1900Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MimoParam_v1900).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MimoParam_v1900).Choice {
				case ServingCellConfig_Ext_MimoParam_v1900_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case ServingCellConfig_Ext_MimoParam_v1900_Setup:
					if err := (*ie.MimoParam_v1900).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ServingCellMO_OD_r19 != nil {
				if err := ie.ServingCellMO_OD_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Cjt_Scheme_PDSCH_v1900 != nil {
				if err := ex.EncodeEnumerated(*ie.Cjt_Scheme_PDSCH_v1900, servingCellConfigExtCjtSchemePDSCHV1900Constraints); err != nil {
					return err
				}
			}

			if ie.Ntn_RedcapPrioritizeUL_Dynamic_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_RedcapPrioritizeUL_Dynamic_r19, servingCellConfigExtNtnRedcapPrioritizeULDynamicR19Constraints); err != nil {
					return err
				}
			}

			if ie.Ntn_RedcapPrioritizeUL_Semistatic_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Ntn_RedcapPrioritizeUL_Semistatic_r19, servingCellConfigExtNtnRedcapPrioritizeULSemistaticR19Constraints); err != nil {
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

func (ie *ServingCellConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(servingCellConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. tdd-UL-DL-ConfigurationDedicated: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Tdd_UL_DL_ConfigurationDedicated = new(TDD_UL_DL_ConfigDedicated)
			if err := ie.Tdd_UL_DL_ConfigurationDedicated.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. initialDownlinkBWP: ref
	{
		if seq.IsComponentPresent(1) {
			ie.InitialDownlinkBWP = new(BWP_DownlinkDedicated)
			if err := ie.InitialDownlinkBWP.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. downlinkBWP-ToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(servingCellConfigDownlinkBWPToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DownlinkBWP_ToReleaseList = make([]BWP_Id, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DownlinkBWP_ToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. downlinkBWP-ToAddModList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(servingCellConfigDownlinkBWPToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DownlinkBWP_ToAddModList = make([]BWP_Downlink, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DownlinkBWP_ToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. firstActiveDownlinkBWP-Id: ref
	{
		if seq.IsComponentPresent(4) {
			ie.FirstActiveDownlinkBWP_Id = new(BWP_Id)
			if err := ie.FirstActiveDownlinkBWP_Id.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. bwp-InactivityTimer: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(servingCellConfigBwpInactivityTimerConstraints)
			if err != nil {
				return err
			}
			ie.Bwp_InactivityTimer = &idx
		}
	}

	// 9. defaultDownlinkBWP-Id: ref
	{
		if seq.IsComponentPresent(6) {
			ie.DefaultDownlinkBWP_Id = new(BWP_Id)
			if err := ie.DefaultDownlinkBWP_Id.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. uplinkConfig: ref
	{
		if seq.IsComponentPresent(7) {
			ie.UplinkConfig = new(UplinkConfig)
			if err := ie.UplinkConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. supplementaryUplink: ref
	{
		if seq.IsComponentPresent(8) {
			ie.SupplementaryUplink = new(UplinkConfig)
			if err := ie.SupplementaryUplink.Decode(d); err != nil {
				return err
			}
		}
	}

	// 12. pdcch-ServingCellConfig: choice
	{
		if seq.IsComponentPresent(9) {
			ie.Pdcch_ServingCellConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_ServingCellConfig
			}{}
			choiceDec := d.NewChoiceDecoder(servingCellConfigPdcchServingCellConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_ServingCellConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ServingCellConfig_Pdcch_ServingCellConfig_Release:
				(*ie.Pdcch_ServingCellConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Pdcch_ServingCellConfig_Setup:
				(*ie.Pdcch_ServingCellConfig).Setup = new(PDCCH_ServingCellConfig)
				if err := (*ie.Pdcch_ServingCellConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. pdsch-ServingCellConfig: choice
	{
		if seq.IsComponentPresent(10) {
			ie.Pdsch_ServingCellConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_ServingCellConfig
			}{}
			choiceDec := d.NewChoiceDecoder(servingCellConfigPdschServingCellConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_ServingCellConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ServingCellConfig_Pdsch_ServingCellConfig_Release:
				(*ie.Pdsch_ServingCellConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Pdsch_ServingCellConfig_Setup:
				(*ie.Pdsch_ServingCellConfig).Setup = new(PDSCH_ServingCellConfig)
				if err := (*ie.Pdsch_ServingCellConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. csi-MeasConfig: choice
	{
		if seq.IsComponentPresent(11) {
			ie.Csi_MeasConfig = &struct {
				Choice  int
				Release *struct{}
				Setup   *CSI_MeasConfig
			}{}
			choiceDec := d.NewChoiceDecoder(servingCellConfigCsiMeasConfigConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Csi_MeasConfig).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case ServingCellConfig_Csi_MeasConfig_Release:
				(*ie.Csi_MeasConfig).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Csi_MeasConfig_Setup:
				(*ie.Csi_MeasConfig).Setup = new(CSI_MeasConfig)
				if err := (*ie.Csi_MeasConfig).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 15. sCellDeactivationTimer: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(servingCellConfigSCellDeactivationTimerConstraints)
			if err != nil {
				return err
			}
			ie.SCellDeactivationTimer = &idx
		}
	}

	// 16. crossCarrierSchedulingConfig: ref
	{
		if seq.IsComponentPresent(13) {
			ie.CrossCarrierSchedulingConfig = new(CrossCarrierSchedulingConfig)
			if err := ie.CrossCarrierSchedulingConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. tag-Id: ref
	{
		if err := ie.Tag_Id.Decode(d); err != nil {
			return err
		}
	}

	// 18. dummy1: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(servingCellConfigDummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &idx
		}
	}

	// 19. pathlossReferenceLinking: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(servingCellConfigPathlossReferenceLinkingConstraints)
			if err != nil {
				return err
			}
			ie.PathlossReferenceLinking = &idx
		}
	}

	// 20. servingCellMO: ref
	{
		if seq.IsComponentPresent(17) {
			ie.ServingCellMO = new(MeasObjectId)
			if err := ie.ServingCellMO.Decode(d); err != nil {
				return err
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
				{Name: "lte-CRS-ToMatchAround", Optional: true},
				{Name: "rateMatchPatternToAddModList", Optional: true},
				{Name: "rateMatchPatternToReleaseList", Optional: true},
				{Name: "downlinkChannelBW-PerSCS-List", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Lte_CRS_ToMatchAround = &struct {
				Choice  int
				Release *struct{}
				Setup   *RateMatchPatternLTE_CRS
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteCRSToMatchAroundConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_ToMatchAround).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Release:
				(*ie.Lte_CRS_ToMatchAround).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_CRS_ToMatchAround_Setup:
				(*ie.Lte_CRS_ToMatchAround).Setup = new(RateMatchPatternLTE_CRS)
				if err := (*ie.Lte_CRS_ToMatchAround).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtRateMatchPatternToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToAddModList = make([]RateMatchPattern, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToAddModList[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtRateMatchPatternToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToReleaseList = make([]RateMatchPatternId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToReleaseList[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtDownlinkChannelBWPerSCSListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.DownlinkChannelBW_PerSCS_List = make([]SCS_SpecificCarrier, n)
			for i := int64(0); i < n; i++ {
				if err := ie.DownlinkChannelBW_PerSCS_List[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "supplementaryUplinkRelease-r16", Optional: true},
				{Name: "tdd-UL-DL-ConfigurationDedicated-IAB-MT-r16", Optional: true},
				{Name: "dormantBWP-Config-r16", Optional: true},
				{Name: "ca-SlotOffset-r16", Optional: true},
				{Name: "dummy2", Optional: true},
				{Name: "intraCellGuardBandsDL-List-r16", Optional: true},
				{Name: "intraCellGuardBandsUL-List-r16", Optional: true},
				{Name: "csi-RS-ValidationWithDCI-r16", Optional: true},
				{Name: "lte-CRS-PatternList1-r16", Optional: true},
				{Name: "lte-CRS-PatternList2-r16", Optional: true},
				{Name: "crs-RateMatch-PerCORESETPoolIndex-r16", Optional: true},
				{Name: "enableTwoDefaultTCI-States-r16", Optional: true},
				{Name: "enableDefaultTCI-StatePerCoresetPoolIndex-r16", Optional: true},
				{Name: "enableBeamSwitchTiming-r16", Optional: true},
				{Name: "cbg-TxDiffTBsProcessingType1-r16", Optional: true},
				{Name: "cbg-TxDiffTBsProcessingType2-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtSupplementaryUplinkReleaseR16Constraints)
			if err != nil {
				return err
			}
			ie.SupplementaryUplinkRelease_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16 = new(TDD_UL_DL_ConfigDedicated_IAB_MT_r16)
			if err := ie.Tdd_UL_DL_ConfigurationDedicated_IAB_MT_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.DormantBWP_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DormantBWP_Config_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtDormantBWPConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.DormantBWP_Config_r16).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_DormantBWP_Config_r16_Release:
				(*ie.DormantBWP_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_DormantBWP_Config_r16_Setup:
				(*ie.DormantBWP_Config_r16).Setup = new(DormantBWP_Config_r16)
				if err := (*ie.DormantBWP_Config_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Ca_SlotOffset_r16 = &struct {
				Choice       int
				RefSCS15kHz  *int64
				RefSCS30KHz  *int64
				RefSCS60KHz  *int64
				RefSCS120KHz *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtCaSlotOffsetR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ca_SlotOffset_r16).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS15kHz:
				(*ie.Ca_SlotOffset_r16).RefSCS15kHz = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(-2, 2))
				if err != nil {
					return err
				}
				(*(*ie.Ca_SlotOffset_r16).RefSCS15kHz) = v
			case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS30KHz:
				(*ie.Ca_SlotOffset_r16).RefSCS30KHz = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(-5, 5))
				if err != nil {
					return err
				}
				(*(*ie.Ca_SlotOffset_r16).RefSCS30KHz) = v
			case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS60KHz:
				(*ie.Ca_SlotOffset_r16).RefSCS60KHz = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(-10, 10))
				if err != nil {
					return err
				}
				(*(*ie.Ca_SlotOffset_r16).RefSCS60KHz) = v
			case ServingCellConfig_Ext_Ca_SlotOffset_r16_RefSCS120KHz:
				(*ie.Ca_SlotOffset_r16).RefSCS120KHz = new(int64)
				v, err := dx.DecodeInteger(per.Constrained(-20, 20))
				if err != nil {
					return err
				}
				(*(*ie.Ca_SlotOffset_r16).RefSCS120KHz) = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Dummy2 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DummyJ
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtDummy2Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dummy2).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Dummy2_Release:
				(*ie.Dummy2).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Dummy2_Setup:
				(*ie.Dummy2).Setup = new(DummyJ)
				if err := (*ie.Dummy2).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtIntraCellGuardBandsDLListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntraCellGuardBandsDL_List_r16 = make([]IntraCellGuardBandsPerSCS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntraCellGuardBandsDL_List_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtIntraCellGuardBandsULListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntraCellGuardBandsUL_List_r16 = make([]IntraCellGuardBandsPerSCS_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntraCellGuardBandsUL_List_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCsiRSValidationWithDCIR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_ValidationWithDCI_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Lte_CRS_PatternList1_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTE_CRS_PatternList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteCRSPatternList1R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_PatternList1_r16).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Release:
				(*ie.Lte_CRS_PatternList1_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_CRS_PatternList1_r16_Setup:
				(*ie.Lte_CRS_PatternList1_r16).Setup = new(LTE_CRS_PatternList_r16)
				if err := (*ie.Lte_CRS_PatternList1_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Lte_CRS_PatternList2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTE_CRS_PatternList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteCRSPatternList2R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_PatternList2_r16).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Release:
				(*ie.Lte_CRS_PatternList2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_CRS_PatternList2_r16_Setup:
				(*ie.Lte_CRS_PatternList2_r16).Setup = new(LTE_CRS_PatternList_r16)
				if err := (*ie.Lte_CRS_PatternList2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCrsRateMatchPerCORESETPoolIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.Crs_RateMatch_PerCORESETPoolIndex_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtEnableTwoDefaultTCIStatesR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableTwoDefaultTCI_States_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtEnableDefaultTCIStatePerCoresetPoolIndexR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableDefaultTCI_StatePerCoresetPoolIndex_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtEnableBeamSwitchTimingR16Constraints)
			if err != nil {
				return err
			}
			ie.EnableBeamSwitchTiming_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCbgTxDiffTBsProcessingType1R16Constraints)
			if err != nil {
				return err
			}
			ie.Cbg_TxDiffTBsProcessingType1_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCbgTxDiffTBsProcessingType2R16Constraints)
			if err != nil {
				return err
			}
			ie.Cbg_TxDiffTBsProcessingType2_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "directionalCollisionHandling-r16", Optional: true},
				{Name: "channelAccessConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtDirectionalCollisionHandlingR16Constraints)
			if err != nil {
				return err
			}
			ie.DirectionalCollisionHandling_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ChannelAccessConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *ChannelAccessConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtChannelAccessConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ChannelAccessConfig_r16).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_ChannelAccessConfig_r16_Release:
				(*ie.ChannelAccessConfig_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_ChannelAccessConfig_r16_Setup:
				(*ie.ChannelAccessConfig_r16).Setup = new(ChannelAccessConfig_r16)
				if err := (*ie.ChannelAccessConfig_r16).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "nr-dl-PRS-PDC-Info-r17", Optional: true},
				{Name: "semiStaticChannelAccessConfigUE-r17", Optional: true},
				{Name: "mimoParam-r17", Optional: true},
				{Name: "channelAccessMode2-r17", Optional: true},
				{Name: "timeDomainHARQ-BundlingType1-r17", Optional: true},
				{Name: "nrofHARQ-BundlingGroups-r17", Optional: true},
				{Name: "fdmed-ReceptionMulticast-r17", Optional: true},
				{Name: "moreThanOneNackOnlyMode-r17", Optional: true},
				{Name: "tci-ActivatedConfig-r17", Optional: true},
				{Name: "directionalCollisionHandling-DC-r17", Optional: true},
				{Name: "lte-NeighCellsCRS-AssistInfoList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Nr_Dl_PRS_PDC_Info_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *NR_DL_PRS_PDC_Info_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtNrDlPRSPDCInfoR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Nr_Dl_PRS_PDC_Info_r17).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Release:
				(*ie.Nr_Dl_PRS_PDC_Info_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Nr_Dl_PRS_PDC_Info_r17_Setup:
				(*ie.Nr_Dl_PRS_PDC_Info_r17).Setup = new(NR_DL_PRS_PDC_Info_r17)
				if err := (*ie.Nr_Dl_PRS_PDC_Info_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.SemiStaticChannelAccessConfigUE_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SemiStaticChannelAccessConfigUE_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtSemiStaticChannelAccessConfigUER17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.SemiStaticChannelAccessConfigUE_r17).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Release:
				(*ie.SemiStaticChannelAccessConfigUE_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_SemiStaticChannelAccessConfigUE_r17_Setup:
				(*ie.SemiStaticChannelAccessConfigUE_r17).Setup = new(SemiStaticChannelAccessConfigUE_r17)
				if err := (*ie.SemiStaticChannelAccessConfigUE_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MimoParam_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MIMOParam_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtMimoParamR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MimoParam_r17).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_MimoParam_r17_Release:
				(*ie.MimoParam_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_MimoParam_r17_Setup:
				(*ie.MimoParam_r17).Setup = new(MIMOParam_r17)
				if err := (*ie.MimoParam_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtChannelAccessMode2R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessMode2_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtTimeDomainHARQBundlingType1R17Constraints)
			if err != nil {
				return err
			}
			ie.TimeDomainHARQ_BundlingType1_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtNrofHARQBundlingGroupsR17Constraints)
			if err != nil {
				return err
			}
			ie.NrofHARQ_BundlingGroups_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtFdmedReceptionMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.Fdmed_ReceptionMulticast_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtMoreThanOneNackOnlyModeR17Constraints)
			if err != nil {
				return err
			}
			ie.MoreThanOneNackOnlyMode_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Tci_ActivatedConfig_r17 = new(TCI_ActivatedConfig_r17)
			if err := ie.Tci_ActivatedConfig_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtDirectionalCollisionHandlingDCR17Constraints)
			if err != nil {
				return err
			}
			ie.DirectionalCollisionHandling_DC_r17 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			ie.Lte_NeighCellsCRS_AssistInfoList_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTE_NeighCellsCRS_AssistInfoList_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteNeighCellsCRSAssistInfoListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Release:
				(*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_NeighCellsCRS_AssistInfoList_r17_Setup:
				(*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Setup = new(LTE_NeighCellsCRS_AssistInfoList_r17)
				if err := (*ie.Lte_NeighCellsCRS_AssistInfoList_r17).Setup.Decode(dx); err != nil {
					return err
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
				{Name: "lte-NeighCellsCRS-Assumptions-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtLteNeighCellsCRSAssumptionsR17Constraints)
			if err != nil {
				return err
			}
			ie.Lte_NeighCellsCRS_Assumptions_r17 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "crossCarrierSchedulingConfigRelease-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCrossCarrierSchedulingConfigReleaseR17Constraints)
			if err != nil {
				return err
			}
			ie.CrossCarrierSchedulingConfigRelease_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multiPDSCH-PerSlotType1-CB-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtMultiPDSCHPerSlotType1CBR17Constraints)
			if err != nil {
				return err
			}
			ie.MultiPDSCH_PerSlotType1_CB_r17 = &v
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "lte-CRS-PatternList3-r18", Optional: true},
				{Name: "lte-CRS-PatternList4-r18", Optional: true},
				{Name: "pdcch-CandidateReceptionWithCRS-Overlap-r18", Optional: true},
				{Name: "cjt-Scheme-PDSCH-r18", Optional: true},
				{Name: "tag2-r18", Optional: true},
				{Name: "cellDTX-DRX-Config-r18", Optional: true},
				{Name: "positionInDCI-cellDTRX-r18", Optional: true},
				{Name: "cellDTX-DRX-L1activation-r18", Optional: true},
				{Name: "mc-DCI-SetOfCellsToAddModList-r18", Optional: true},
				{Name: "mc-DCI-SetOfCellsToReleaseList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Lte_CRS_PatternList3_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTE_CRS_PatternList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteCRSPatternList3R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_PatternList3_r18).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Release:
				(*ie.Lte_CRS_PatternList3_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_CRS_PatternList3_r18_Setup:
				(*ie.Lte_CRS_PatternList3_r18).Setup = new(LTE_CRS_PatternList_r16)
				if err := (*ie.Lte_CRS_PatternList3_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Lte_CRS_PatternList4_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LTE_CRS_PatternList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtLteCRSPatternList4R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lte_CRS_PatternList4_r18).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Release:
				(*ie.Lte_CRS_PatternList4_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_Lte_CRS_PatternList4_r18_Setup:
				(*ie.Lte_CRS_PatternList4_r18).Setup = new(LTE_CRS_PatternList_r16)
				if err := (*ie.Lte_CRS_PatternList4_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtPdcchCandidateReceptionWithCRSOverlapR18Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_CandidateReceptionWithCRS_Overlap_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCjtSchemePDSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Cjt_Scheme_PDSCH_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.Tag2_r18 = new(Tag2_r18)
			if err := ie.Tag2_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.CellDTX_DRX_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CellDTX_DRX_Config_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtCellDTXDRXConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CellDTX_DRX_Config_r18).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Release:
				(*ie.CellDTX_DRX_Config_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_CellDTX_DRX_Config_r18_Setup:
				(*ie.CellDTX_DRX_Config_r18).Setup = new(CellDTX_DRX_Config_r18)
				if err := (*ie.CellDTX_DRX_Config_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeInteger(servingCellConfigPositionInDCICellDTRXR18Constraints)
			if err != nil {
				return err
			}
			ie.PositionInDCI_CellDTRX_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCellDTXDRXL1activationR18Constraints)
			if err != nil {
				return err
			}
			ie.CellDTX_DRX_L1activation_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtMcDCISetOfCellsToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mc_DCI_SetOfCellsToAddModList_r18 = make([]MC_DCI_SetOfCells_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mc_DCI_SetOfCellsToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtMcDCISetOfCellsToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mc_DCI_SetOfCellsToReleaseList_r18 = make([]SetOfCellsId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mc_DCI_SetOfCellsToReleaseList_r18[i].Decode(dx); err != nil {
					return err
				}
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
				{Name: "mimoParam-v1850", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MimoParam_v1850 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MIMOParam_v1850
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtMimoParamV1850Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MimoParam_v1850).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_MimoParam_v1850_Release:
				(*ie.MimoParam_v1850).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_MimoParam_v1850_Setup:
				(*ie.MimoParam_v1850).Setup = new(MIMOParam_v1850)
				if err := (*ie.MimoParam_v1850).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "mc-DCI-SetOfCellsToAddModListExt-v1900", Optional: true},
				{Name: "mimoParam-v1900", Optional: true},
				{Name: "servingCellMO-OD-r19", Optional: true},
				{Name: "cjt-Scheme-PDSCH-v1900", Optional: true},
				{Name: "ntn-RedcapPrioritizeUL-Dynamic-r19", Optional: true},
				{Name: "ntn-RedcapPrioritizeUL-Semistatic-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(servingCellConfigExtMcDCISetOfCellsToAddModListExtV1900Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mc_DCI_SetOfCellsToAddModListExt_v1900 = make([]MC_DCI_SetOfCellsExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mc_DCI_SetOfCellsToAddModListExt_v1900[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MimoParam_v1900 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MIMOParam_v1900
			}{}
			choiceDec := dx.NewChoiceDecoder(servingCellConfigExtMimoParamV1900Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MimoParam_v1900).Choice = int(index)
			switch index {
			case ServingCellConfig_Ext_MimoParam_v1900_Release:
				(*ie.MimoParam_v1900).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case ServingCellConfig_Ext_MimoParam_v1900_Setup:
				(*ie.MimoParam_v1900).Setup = new(MIMOParam_v1900)
				if err := (*ie.MimoParam_v1900).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.ServingCellMO_OD_r19 = new(MeasObjectId)
			if err := ie.ServingCellMO_OD_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtCjtSchemePDSCHV1900Constraints)
			if err != nil {
				return err
			}
			ie.Cjt_Scheme_PDSCH_v1900 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtNtnRedcapPrioritizeULDynamicR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_RedcapPrioritizeUL_Dynamic_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(servingCellConfigExtNtnRedcapPrioritizeULSemistaticR19Constraints)
			if err != nil {
				return err
			}
			ie.Ntn_RedcapPrioritizeUL_Semistatic_r19 = &v
		}
	}

	return nil
}
