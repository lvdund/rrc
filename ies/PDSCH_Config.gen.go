// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PDSCH-Config (line 11300).

var pDSCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataScramblingIdentityPDSCH", Optional: true},
		{Name: "dmrs-DownlinkForPDSCH-MappingTypeA", Optional: true},
		{Name: "dmrs-DownlinkForPDSCH-MappingTypeB", Optional: true},
		{Name: "tci-StatesToAddModList", Optional: true},
		{Name: "tci-StatesToReleaseList", Optional: true},
		{Name: "vrb-ToPRB-Interleaver", Optional: true},
		{Name: "resourceAllocation"},
		{Name: "pdsch-TimeDomainAllocationList", Optional: true},
		{Name: "pdsch-AggregationFactor", Optional: true},
		{Name: "rateMatchPatternToAddModList", Optional: true},
		{Name: "rateMatchPatternToReleaseList", Optional: true},
		{Name: "rateMatchPatternGroup1", Optional: true},
		{Name: "rateMatchPatternGroup2", Optional: true},
		{Name: "rbg-Size"},
		{Name: "mcs-Table", Optional: true},
		{Name: "maxNrofCodeWordsScheduledByDCI", Optional: true},
		{Name: "prb-BundlingType"},
		{Name: "zp-CSI-RS-ResourceToAddModList", Optional: true},
		{Name: "zp-CSI-RS-ResourceToReleaseList", Optional: true},
		{Name: "aperiodic-ZP-CSI-RS-ResourceSetsToAddModList", Optional: true},
		{Name: "aperiodic-ZP-CSI-RS-ResourceSetsToReleaseList", Optional: true},
		{Name: "sp-ZP-CSI-RS-ResourceSetsToAddModList", Optional: true},
		{Name: "sp-ZP-CSI-RS-ResourceSetsToReleaseList", Optional: true},
		{Name: "p-ZP-CSI-RS-ResourceSet", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
		{Name: "extension-addition-5", Optional: true},
		{Name: "extension-addition-6", Optional: true},
	},
}

var pDSCHConfigDataScramblingIdentityPDSCHConstraints = per.Constrained(0, 1023)

var pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeAConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Release = 0
	PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Setup   = 1
)

var pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeBConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Release = 0
	PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Setup   = 1
)

var pDSCHConfigTciStatesToAddModListConstraints = per.SizeRange(1, common.MaxNrofTCI_States)

var pDSCHConfigTciStatesToReleaseListConstraints = per.SizeRange(1, common.MaxNrofTCI_States)

const (
	PDSCH_Config_Vrb_ToPRB_Interleaver_N2 = 0
	PDSCH_Config_Vrb_ToPRB_Interleaver_N4 = 1
)

var pDSCHConfigVrbToPRBInterleaverConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Vrb_ToPRB_Interleaver_N2, PDSCH_Config_Vrb_ToPRB_Interleaver_N4},
}

const (
	PDSCH_Config_ResourceAllocation_ResourceAllocationType0 = 0
	PDSCH_Config_ResourceAllocation_ResourceAllocationType1 = 1
	PDSCH_Config_ResourceAllocation_DynamicSwitch           = 2
)

var pDSCHConfigResourceAllocationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_ResourceAllocation_ResourceAllocationType0, PDSCH_Config_ResourceAllocation_ResourceAllocationType1, PDSCH_Config_ResourceAllocation_DynamicSwitch},
}

var pDSCH_ConfigPdschTimeDomainAllocationListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Pdsch_TimeDomainAllocationList_Release = 0
	PDSCH_Config_Pdsch_TimeDomainAllocationList_Setup   = 1
)

const (
	PDSCH_Config_Pdsch_AggregationFactor_N2 = 0
	PDSCH_Config_Pdsch_AggregationFactor_N4 = 1
	PDSCH_Config_Pdsch_AggregationFactor_N8 = 2
)

var pDSCHConfigPdschAggregationFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Pdsch_AggregationFactor_N2, PDSCH_Config_Pdsch_AggregationFactor_N4, PDSCH_Config_Pdsch_AggregationFactor_N8},
}

var pDSCHConfigRateMatchPatternToAddModListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

var pDSCHConfigRateMatchPatternToReleaseListConstraints = per.SizeRange(1, common.MaxNrofRateMatchPatterns)

const (
	PDSCH_Config_Rbg_Size_Config1 = 0
	PDSCH_Config_Rbg_Size_Config2 = 1
)

var pDSCHConfigRbgSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Rbg_Size_Config1, PDSCH_Config_Rbg_Size_Config2},
}

const (
	PDSCH_Config_Mcs_Table_Qam256     = 0
	PDSCH_Config_Mcs_Table_Qam64LowSE = 1
)

var pDSCHConfigMcsTableConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Mcs_Table_Qam256, PDSCH_Config_Mcs_Table_Qam64LowSE},
}

const (
	PDSCH_Config_MaxNrofCodeWordsScheduledByDCI_N1 = 0
	PDSCH_Config_MaxNrofCodeWordsScheduledByDCI_N2 = 1
)

var pDSCHConfigMaxNrofCodeWordsScheduledByDCIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_MaxNrofCodeWordsScheduledByDCI_N1, PDSCH_Config_MaxNrofCodeWordsScheduledByDCI_N2},
}

var pDSCH_ConfigPrbBundlingTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "staticBundling"},
		{Name: "dynamicBundling"},
	},
}

const (
	PDSCH_Config_Prb_BundlingType_StaticBundling  = 0
	PDSCH_Config_Prb_BundlingType_DynamicBundling = 1
)

var pDSCHConfigZpCSIRSResourceToAddModListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_Resources)

var pDSCHConfigZpCSIRSResourceToReleaseListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_Resources)

var pDSCHConfigAperiodicZPCSIRSResourceSetsToAddModListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCHConfigAperiodicZPCSIRSResourceSetsToReleaseListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCHConfigSpZPCSIRSResourceSetsToAddModListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCHConfigSpZPCSIRSResourceSetsToReleaseListConstraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCH_ConfigPZPCSIRSResourceSetConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Release = 0
	PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Setup   = 1
)

var pDSCHConfigExtMaxMIMOLayersR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_MaxMIMO_Layers_r16_Release = 0
	PDSCH_Config_Ext_MaxMIMO_Layers_r16_Setup   = 1
)

var pDSCHConfigExtMinimumSchedulingOffsetK0R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Release = 0
	PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Setup   = 1
)

const (
	PDSCH_Config_Ext_AntennaPortsFieldPresenceDCI_1_2_r16_Enabled = 0
)

var pDSCHConfigExtAntennaPortsFieldPresenceDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_AntennaPortsFieldPresenceDCI_1_2_r16_Enabled},
}

var pDSCHConfigExtAperiodicZPCSIRSResourceSetsToAddModListDCI12R16Constraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCHConfigExtAperiodicZPCSIRSResourceSetsToReleaseListDCI12R16Constraints = per.SizeRange(1, common.MaxNrofZP_CSI_RS_ResourceSets)

var pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeADCI12R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Release = 0
	PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Setup   = 1
)

var pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeBDCI12R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Release = 0
	PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Setup   = 1
)

const (
	PDSCH_Config_Ext_Dmrs_SequenceInitializationDCI_1_2_r16_Enabled = 0
)

var pDSCHConfigExtDmrsSequenceInitializationDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Dmrs_SequenceInitializationDCI_1_2_r16_Enabled},
}

var pDSCHConfigHarqProcessNumberSizeDCI12R16Constraints = per.Constrained(0, 4)

const (
	PDSCH_Config_Ext_Mcs_TableDCI_1_2_r16_Qam256     = 0
	PDSCH_Config_Ext_Mcs_TableDCI_1_2_r16_Qam64LowSE = 1
)

var pDSCHConfigExtMcsTableDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mcs_TableDCI_1_2_r16_Qam256, PDSCH_Config_Ext_Mcs_TableDCI_1_2_r16_Qam64LowSE},
}

var pDSCHConfigNumberOfBitsForRVDCI12R16Constraints = per.Constrained(0, 2)

var pDSCHConfigExtPdschTimeDomainAllocationListDCI12R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Release = 0
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Setup   = 1
)

var pDSCHConfigExtPrbBundlingTypeDCI12R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "staticBundling-r16"},
		{Name: "dynamicBundling-r16"},
	},
}

const (
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16  = 0
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16 = 1
)

const (
	PDSCH_Config_Ext_PriorityIndicatorDCI_1_2_r16_Enabled = 0
)

var pDSCHConfigExtPriorityIndicatorDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_PriorityIndicatorDCI_1_2_r16_Enabled},
}

const (
	PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N2  = 0
	PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N4  = 1
	PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N8  = 2
	PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N16 = 3
)

var pDSCHConfigExtResourceAllocationType1GranularityDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N2, PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N4, PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N8, PDSCH_Config_Ext_ResourceAllocationType1GranularityDCI_1_2_r16_N16},
}

const (
	PDSCH_Config_Ext_Vrb_ToPRB_InterleaverDCI_1_2_r16_N2 = 0
	PDSCH_Config_Ext_Vrb_ToPRB_InterleaverDCI_1_2_r16_N4 = 1
)

var pDSCHConfigExtVrbToPRBInterleaverDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Vrb_ToPRB_InterleaverDCI_1_2_r16_N2, PDSCH_Config_Ext_Vrb_ToPRB_InterleaverDCI_1_2_r16_N4},
}

const (
	PDSCH_Config_Ext_ReferenceOfSLIVDCI_1_2_r16_Enabled = 0
)

var pDSCHConfigExtReferenceOfSLIVDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_ReferenceOfSLIVDCI_1_2_r16_Enabled},
}

const (
	PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_ResourceAllocationType0 = 0
	PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_ResourceAllocationType1 = 1
	PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_DynamicSwitch           = 2
)

var pDSCHConfigExtResourceAllocationDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_ResourceAllocationType0, PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_ResourceAllocationType1, PDSCH_Config_Ext_ResourceAllocationDCI_1_2_r16_DynamicSwitch},
}

const (
	PDSCH_Config_Ext_PriorityIndicatorDCI_1_1_r16_Enabled = 0
)

var pDSCHConfigExtPriorityIndicatorDCI11R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_PriorityIndicatorDCI_1_1_r16_Enabled},
}

var pDSCHConfigDataScramblingIdentityPDSCH2R16Constraints = per.Constrained(0, 1023)

var pDSCHConfigExtPdschTimeDomainAllocationListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Release = 0
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Setup   = 1
)

var pDSCHConfigExtRepetitionSchemeConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Release = 0
	PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Setup   = 1
)

var pDSCHConfigExtRepetitionSchemeConfigV1630Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Release = 0
	PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Setup   = 1
)

const (
	PDSCH_Config_Ext_Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17_Enabled = 0
)

var pDSCHConfigExtPdschHARQACKOneShotFeedbackDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17_Enabled},
}

const (
	PDSCH_Config_Ext_Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17_Enabled = 0
)

var pDSCHConfigExtPdschHARQACKEnhType3DCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17_Enabled},
}

const (
	PDSCH_Config_Ext_Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17_Enabled = 0
)

var pDSCHConfigExtPdschHARQACKEnhType3DCIField12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17_Enabled},
}

const (
	PDSCH_Config_Ext_Pdsch_HARQ_ACK_RetxDCI_1_2_r17_Enabled = 0
)

var pDSCHConfigExtPdschHARQACKRetxDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Pdsch_HARQ_ACK_RetxDCI_1_2_r17_Enabled},
}

const (
	PDSCH_Config_Ext_Pucch_SSCellDynDCI_1_2_r17_Enabled = 0
)

var pDSCHConfigExtPucchSSCellDynDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Pucch_SSCellDynDCI_1_2_r17_Enabled},
}

var pDSCHConfigExtDlOrJointTCIStateListR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "explicitlist"},
		{Name: "unifiedTCI-StateRef-r17"},
	},
}

const (
	PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_Explicitlist            = 0
	PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_UnifiedTCI_StateRef_r17 = 1
)

const (
	PDSCH_Config_Ext_BeamAppTime_r17_N1     = 0
	PDSCH_Config_Ext_BeamAppTime_r17_N2     = 1
	PDSCH_Config_Ext_BeamAppTime_r17_N4     = 2
	PDSCH_Config_Ext_BeamAppTime_r17_N7     = 3
	PDSCH_Config_Ext_BeamAppTime_r17_N14    = 4
	PDSCH_Config_Ext_BeamAppTime_r17_N28    = 5
	PDSCH_Config_Ext_BeamAppTime_r17_N42    = 6
	PDSCH_Config_Ext_BeamAppTime_r17_N56    = 7
	PDSCH_Config_Ext_BeamAppTime_r17_N70    = 8
	PDSCH_Config_Ext_BeamAppTime_r17_N84    = 9
	PDSCH_Config_Ext_BeamAppTime_r17_N98    = 10
	PDSCH_Config_Ext_BeamAppTime_r17_N112   = 11
	PDSCH_Config_Ext_BeamAppTime_r17_N224   = 12
	PDSCH_Config_Ext_BeamAppTime_r17_N336   = 13
	PDSCH_Config_Ext_BeamAppTime_r17_Spare2 = 14
	PDSCH_Config_Ext_BeamAppTime_r17_Spare1 = 15
)

var pDSCHConfigExtBeamAppTimeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_BeamAppTime_r17_N1, PDSCH_Config_Ext_BeamAppTime_r17_N2, PDSCH_Config_Ext_BeamAppTime_r17_N4, PDSCH_Config_Ext_BeamAppTime_r17_N7, PDSCH_Config_Ext_BeamAppTime_r17_N14, PDSCH_Config_Ext_BeamAppTime_r17_N28, PDSCH_Config_Ext_BeamAppTime_r17_N42, PDSCH_Config_Ext_BeamAppTime_r17_N56, PDSCH_Config_Ext_BeamAppTime_r17_N70, PDSCH_Config_Ext_BeamAppTime_r17_N84, PDSCH_Config_Ext_BeamAppTime_r17_N98, PDSCH_Config_Ext_BeamAppTime_r17_N112, PDSCH_Config_Ext_BeamAppTime_r17_N224, PDSCH_Config_Ext_BeamAppTime_r17_N336, PDSCH_Config_Ext_BeamAppTime_r17_Spare2, PDSCH_Config_Ext_BeamAppTime_r17_Spare1},
}

var pDSCHConfigExtDummyConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Dummy_Release = 0
	PDSCH_Config_Ext_Dummy_Setup   = 1
)

const (
	PDSCH_Config_Ext_Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17_True = 0
)

var pDSCHConfigExtDmrsFDOCCDisabledForRank1PDSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17_True},
}

var pDSCHConfigExtMinimumSchedulingOffsetK0R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Release = 0
	PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Setup   = 1
)

var pDSCHConfigHarqProcessNumberSizeDCI12V1700Constraints = per.Constrained(0, 5)

var pDSCHConfigHarqProcessNumberSizeDCI11R17Constraints = per.Constrained(5, 5)

const (
	PDSCH_Config_Ext_Mcs_Table_r17_Qam1024 = 0
)

var pDSCHConfigExtMcsTableR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mcs_Table_r17_Qam1024},
}

const (
	PDSCH_Config_Ext_Mcs_TableDCI_1_2_r17_Qam1024 = 0
)

var pDSCHConfigExtMcsTableDCI12R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mcs_TableDCI_1_2_r17_Qam1024},
}

const (
	PDSCH_Config_Ext_XOverheadMulticast_r17_XOh6  = 0
	PDSCH_Config_Ext_XOverheadMulticast_r17_XOh12 = 1
	PDSCH_Config_Ext_XOverheadMulticast_r17_XOh18 = 2
)

var pDSCHConfigExtXOverheadMulticastR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_XOverheadMulticast_r17_XOh6, PDSCH_Config_Ext_XOverheadMulticast_r17_XOh12, PDSCH_Config_Ext_XOverheadMulticast_r17_XOh18},
}

const (
	PDSCH_Config_Ext_PriorityIndicatorDCI_4_2_r17_Enabled = 0
)

var pDSCHConfigExtPriorityIndicatorDCI42R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_PriorityIndicatorDCI_4_2_r17_Enabled},
}

var pDSCHConfigSizeDCI42R17Constraints = per.Constrained(20, common.MaxDCI_4_2_Size_r17)

var pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Release = 0
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Setup   = 1
)

var pDSCHConfigExtAdvancedReceiverMUMIMOR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Release = 0
	PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Setup   = 1
)

var pDSCHConfigExtPdschConfigDCI13R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Release = 0
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Setup   = 1
)

var pDSCHConfigExtPdschConfigDCI13V1860Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Release = 0
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Setup   = 1
)

const (
	PDSCH_Config_Ext_Mg_CancellationDCI_1_1_r19_Enabled = 0
)

var pDSCHConfigExtMgCancellationDCI11R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mg_CancellationDCI_1_1_r19_Enabled},
}

const (
	PDSCH_Config_Ext_Mg_CancellationDCI_1_2_r19_Enabled = 0
)

var pDSCHConfigExtMgCancellationDCI12R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mg_CancellationDCI_1_2_r19_Enabled},
}

const (
	PDSCH_Config_Ext_Mg_CancellationDCI_1_3_r19_Enabled = 0
)

var pDSCHConfigExtMgCancellationDCI13R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Mg_CancellationDCI_1_3_r19_Enabled},
}

var pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHDCI13R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Release = 0
	PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Setup   = 1
)

var pDSCHConfigHarqProcessNumberSizeDCI11ExtR19Constraints = per.Constrained(5, 5)

var pDSCHConfigHarqProcessNumberSizeDCI12ExtR19Constraints = per.Constrained(0, 5)

var pDSCHConfigExtPdschConfigDCI13V1900Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Release = 0
	PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Setup   = 1
)

var pDSCHConfigPrbBundlingTypeStaticBundlingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bundleSize", Optional: true},
	},
}

const (
	PDSCH_Config_Prb_BundlingType_StaticBundling_BundleSize_N4       = 0
	PDSCH_Config_Prb_BundlingType_StaticBundling_BundleSize_Wideband = 1
)

var pDSCHConfigPrbBundlingTypeStaticBundlingBundleSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Prb_BundlingType_StaticBundling_BundleSize_N4, PDSCH_Config_Prb_BundlingType_StaticBundling_BundleSize_Wideband},
}

var pDSCHConfigPrbBundlingTypeDynamicBundlingConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bundleSizeSet1", Optional: true},
		{Name: "bundleSizeSet2", Optional: true},
	},
}

const (
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N4          = 0
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_Wideband    = 1
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N2_Wideband = 2
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N4_Wideband = 3
)

var pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet1Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N4, PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_Wideband, PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N2_Wideband, PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet1_N4_Wideband},
}

const (
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet2_N4       = 0
	PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet2_Wideband = 1
)

var pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet2Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet2_N4, PDSCH_Config_Prb_BundlingType_DynamicBundling_BundleSizeSet2_Wideband},
}

var pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bundleSize-r16", Optional: true},
	},
}

const (
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16_BundleSize_r16_N4       = 0
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16_BundleSize_r16_Wideband = 1
)

var pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16BundleSizeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16_BundleSize_r16_N4, PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16_BundleSize_r16_Wideband},
}

var pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bundleSizeSet1-r16", Optional: true},
		{Name: "bundleSizeSet2-r16", Optional: true},
	},
}

const (
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N4          = 0
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_Wideband    = 1
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N2_Wideband = 2
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N4_Wideband = 3
)

var pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N4, PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_Wideband, PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N2_Wideband, PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet1_r16_N4_Wideband},
}

const (
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet2_r16_N4       = 0
	PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet2_r16_Wideband = 1
)

var pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet2_r16_N4, PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16_BundleSizeSet2_r16_Wideband},
}

var pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "dl-OrJointTCI-StateToAddModList-r17", Optional: true},
		{Name: "dl-OrJointTCI-StateToReleaseList-r17", Optional: true},
	},
}

var pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofTCI_States)

var pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofTCI_States)

type PDSCH_Config struct {
	DataScramblingIdentityPDSCH        *int64
	Dmrs_DownlinkForPDSCH_MappingTypeA *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_DownlinkConfig
	}
	Dmrs_DownlinkForPDSCH_MappingTypeB *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_DownlinkConfig
	}
	Tci_StatesToAddModList         []TCI_State
	Tci_StatesToReleaseList        []TCI_StateId
	Vrb_ToPRB_Interleaver          *int64
	ResourceAllocation             int64
	Pdsch_TimeDomainAllocationList *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_TimeDomainResourceAllocationList
	}
	Pdsch_AggregationFactor        *int64
	RateMatchPatternToAddModList   []RateMatchPattern
	RateMatchPatternToReleaseList  []RateMatchPatternId
	RateMatchPatternGroup1         *RateMatchPatternGroup
	RateMatchPatternGroup2         *RateMatchPatternGroup
	Rbg_Size                       int64
	Mcs_Table                      *int64
	MaxNrofCodeWordsScheduledByDCI *int64
	Prb_BundlingType               struct {
		Choice          int
		StaticBundling  *struct{ BundleSize *int64 }
		DynamicBundling *struct {
			BundleSizeSet1 *int64
			BundleSizeSet2 *int64
		}
	}
	Zp_CSI_RS_ResourceToAddModList                []ZP_CSI_RS_Resource
	Zp_CSI_RS_ResourceToReleaseList               []ZP_CSI_RS_ResourceId
	Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList  []ZP_CSI_RS_ResourceSet
	Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList []ZP_CSI_RS_ResourceSetId
	Sp_ZP_CSI_RS_ResourceSetsToAddModList         []ZP_CSI_RS_ResourceSet
	Sp_ZP_CSI_RS_ResourceSetsToReleaseList        []ZP_CSI_RS_ResourceSetId
	P_ZP_CSI_RS_ResourceSet                       *struct {
		Choice  int
		Release *struct{}
		Setup   *ZP_CSI_RS_ResourceSet
	}
	MaxMIMO_Layers_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MaxMIMO_LayersDL_r16
	}
	MinimumSchedulingOffsetK0_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MinSchedulingOffsetK0_Values_r16
	}
	AntennaPortsFieldPresenceDCI_1_2_r16                    *int64
	AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16  []ZP_CSI_RS_ResourceSet
	AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 []ZP_CSI_RS_ResourceSetId
	Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16          *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_DownlinkConfig
	}
	Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_DownlinkConfig
	}
	Dmrs_SequenceInitializationDCI_1_2_r16    *int64
	Harq_ProcessNumberSizeDCI_1_2_r16         *int64
	Mcs_TableDCI_1_2_r16                      *int64
	NumberOfBitsForRV_DCI_1_2_r16             *int64
	Pdsch_TimeDomainAllocationListDCI_1_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_TimeDomainResourceAllocationList_r16
	}
	Prb_BundlingTypeDCI_1_2_r16 *struct {
		Choice              int
		StaticBundling_r16  *struct{ BundleSize_r16 *int64 }
		DynamicBundling_r16 *struct {
			BundleSizeSet1_r16 *int64
			BundleSizeSet2_r16 *int64
		}
	}
	PriorityIndicatorDCI_1_2_r16                  *int64
	RateMatchPatternGroup1DCI_1_2_r16             *RateMatchPatternGroup
	RateMatchPatternGroup2DCI_1_2_r16             *RateMatchPatternGroup
	ResourceAllocationType1GranularityDCI_1_2_r16 *int64
	Vrb_ToPRB_InterleaverDCI_1_2_r16              *int64
	ReferenceOfSLIVDCI_1_2_r16                    *int64
	ResourceAllocationDCI_1_2_r16                 *int64
	PriorityIndicatorDCI_1_1_r16                  *int64
	DataScramblingIdentityPDSCH2_r16              *int64
	Pdsch_TimeDomainAllocationList_r16            *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_TimeDomainResourceAllocationList_r16
	}
	RepetitionSchemeConfig_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *RepetitionSchemeConfig_r16
	}
	RepetitionSchemeConfig_v1630 *struct {
		Choice  int
		Release *struct{}
		Setup   *RepetitionSchemeConfig_v1630
	}
	Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 *int64
	Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17        *int64
	Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17  *int64
	Pdsch_HARQ_ACK_RetxDCI_1_2_r17            *int64
	Pucch_SSCellDynDCI_1_2_r17                *int64
	Dl_OrJointTCI_StateList_r17               *struct {
		Choice       int
		Explicitlist *struct {
			Dl_OrJointTCI_StateToAddModList_r17  []TCI_State
			Dl_OrJointTCI_StateToReleaseList_r17 []TCI_StateId
		}
		UnifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
	}
	BeamAppTime_r17 *int64
	Dummy           *struct {
		Choice  int
		Release *struct{}
		Setup   *Dummy_TDRA_List
	}
	Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 *int64
	MinimumSchedulingOffsetK0_r17          *struct {
		Choice  int
		Release *struct{}
		Setup   *MinSchedulingOffsetK0_Values_r17
	}
	Harq_ProcessNumberSizeDCI_1_2_v1700             *int64
	Harq_ProcessNumberSizeDCI_1_1_r17               *int64
	Mcs_Table_r17                                   *int64
	Mcs_TableDCI_1_2_r17                            *int64
	XOverheadMulticast_r17                          *int64
	PriorityIndicatorDCI_4_2_r17                    *int64
	SizeDCI_4_2_r17                                 *int64
	Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MultiPDSCH_TDRA_List_r17
	}
	AdvancedReceiver_MU_MIMO_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *AdvancedReceiver_MU_MIMO_r18
	}
	Pdsch_ConfigDCI_1_3_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_ConfigDCI_1_3_r18
	}
	Pdsch_ConfigDCI_1_3_v1860 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_ConfigDCI_1_3_v1860
	}
	Mg_CancellationDCI_1_1_r19                              *int64
	Mg_CancellationDCI_1_2_r19                              *int64
	Mg_CancellationDCI_1_3_r19                              *int64
	Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *MultiPDSCH_TDRA_List_r17
	}
	Harq_ProcessNumberSizeDCI_1_1_Ext_r19 *int64
	Harq_ProcessNumberSizeDCI_1_2_Ext_r19 *int64
	Pdsch_ConfigDCI_1_3_v1900             *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_ConfigDCI_1_3_v1900
	}
}

func (ie *PDSCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDSCHConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MaxMIMO_Layers_r16 != nil || ie.MinimumSchedulingOffsetK0_r16 != nil || ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil || ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 != nil || ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 != nil || ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil || ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil || ie.Mcs_TableDCI_1_2_r16 != nil || ie.NumberOfBitsForRV_DCI_1_2_r16 != nil || ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil || ie.Prb_BundlingTypeDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_2_r16 != nil || ie.RateMatchPatternGroup1DCI_1_2_r16 != nil || ie.RateMatchPatternGroup2DCI_1_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil || ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil || ie.ReferenceOfSLIVDCI_1_2_r16 != nil || ie.ResourceAllocationDCI_1_2_r16 != nil || ie.PriorityIndicatorDCI_1_1_r16 != nil || ie.DataScramblingIdentityPDSCH2_r16 != nil || ie.Pdsch_TimeDomainAllocationList_r16 != nil || ie.RepetitionSchemeConfig_r16 != nil
	hasExtGroup1 := ie.RepetitionSchemeConfig_v1630 != nil
	hasExtGroup2 := ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil || ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil || ie.Pucch_SSCellDynDCI_1_2_r17 != nil || ie.Dl_OrJointTCI_StateList_r17 != nil || ie.BeamAppTime_r17 != nil || ie.Dummy != nil || ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil || ie.MinimumSchedulingOffsetK0_r17 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil || ie.Mcs_Table_r17 != nil || ie.Mcs_TableDCI_1_2_r17 != nil || ie.XOverheadMulticast_r17 != nil || ie.PriorityIndicatorDCI_4_2_r17 != nil || ie.SizeDCI_4_2_r17 != nil
	hasExtGroup3 := ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil
	hasExtGroup4 := ie.AdvancedReceiver_MU_MIMO_r18 != nil || ie.Pdsch_ConfigDCI_1_3_r18 != nil
	hasExtGroup5 := ie.Pdsch_ConfigDCI_1_3_v1860 != nil
	hasExtGroup6 := ie.Mg_CancellationDCI_1_1_r19 != nil || ie.Mg_CancellationDCI_1_2_r19 != nil || ie.Mg_CancellationDCI_1_3_r19 != nil || ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19 != nil || ie.Harq_ProcessNumberSizeDCI_1_1_Ext_r19 != nil || ie.Harq_ProcessNumberSizeDCI_1_2_Ext_r19 != nil || ie.Pdsch_ConfigDCI_1_3_v1900 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataScramblingIdentityPDSCH != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeA != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeB != nil, ie.Tci_StatesToAddModList != nil, ie.Tci_StatesToReleaseList != nil, ie.Vrb_ToPRB_Interleaver != nil, ie.Pdsch_TimeDomainAllocationList != nil, ie.Pdsch_AggregationFactor != nil, ie.RateMatchPatternToAddModList != nil, ie.RateMatchPatternToReleaseList != nil, ie.RateMatchPatternGroup1 != nil, ie.RateMatchPatternGroup2 != nil, ie.Mcs_Table != nil, ie.MaxNrofCodeWordsScheduledByDCI != nil, ie.Zp_CSI_RS_ResourceToAddModList != nil, ie.Zp_CSI_RS_ResourceToReleaseList != nil, ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList != nil, ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList != nil, ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList != nil, ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList != nil, ie.P_ZP_CSI_RS_ResourceSet != nil}); err != nil {
		return err
	}

	// 3. dataScramblingIdentityPDSCH: integer
	{
		if ie.DataScramblingIdentityPDSCH != nil {
			if err := e.EncodeInteger(*ie.DataScramblingIdentityPDSCH, pDSCHConfigDataScramblingIdentityPDSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dmrs-DownlinkForPDSCH-MappingTypeA: choice
	{
		if ie.Dmrs_DownlinkForPDSCH_MappingTypeA != nil {
			choiceEnc := e.NewChoiceEncoder(pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeAConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Choice {
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Setup:
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 5. dmrs-DownlinkForPDSCH-MappingTypeB: choice
	{
		if ie.Dmrs_DownlinkForPDSCH_MappingTypeB != nil {
			choiceEnc := e.NewChoiceEncoder(pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeBConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Choice {
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Setup:
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. tci-StatesToAddModList: sequence-of
	{
		if ie.Tci_StatesToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigTciStatesToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tci_StatesToAddModList))); err != nil {
				return err
			}
			for i := range ie.Tci_StatesToAddModList {
				if err := ie.Tci_StatesToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 7. tci-StatesToReleaseList: sequence-of
	{
		if ie.Tci_StatesToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigTciStatesToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Tci_StatesToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Tci_StatesToReleaseList {
				if err := ie.Tci_StatesToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 8. vrb-ToPRB-Interleaver: enumerated
	{
		if ie.Vrb_ToPRB_Interleaver != nil {
			if err := e.EncodeEnumerated(*ie.Vrb_ToPRB_Interleaver, pDSCHConfigVrbToPRBInterleaverConstraints); err != nil {
				return err
			}
		}
	}

	// 9. resourceAllocation: enumerated
	{
		if err := e.EncodeEnumerated(ie.ResourceAllocation, pDSCHConfigResourceAllocationConstraints); err != nil {
			return err
		}
	}

	// 10. pdsch-TimeDomainAllocationList: choice
	{
		if ie.Pdsch_TimeDomainAllocationList != nil {
			choiceEnc := e.NewChoiceEncoder(pDSCH_ConfigPdschTimeDomainAllocationListConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_TimeDomainAllocationList).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pdsch_TimeDomainAllocationList).Choice {
			case PDSCH_Config_Pdsch_TimeDomainAllocationList_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Pdsch_TimeDomainAllocationList_Setup:
				if err := (*ie.Pdsch_TimeDomainAllocationList).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pdsch_TimeDomainAllocationList).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 11. pdsch-AggregationFactor: enumerated
	{
		if ie.Pdsch_AggregationFactor != nil {
			if err := e.EncodeEnumerated(*ie.Pdsch_AggregationFactor, pDSCHConfigPdschAggregationFactorConstraints); err != nil {
				return err
			}
		}
	}

	// 12. rateMatchPatternToAddModList: sequence-of
	{
		if ie.RateMatchPatternToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigRateMatchPatternToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToAddModList))); err != nil {
				return err
			}
			for i := range ie.RateMatchPatternToAddModList {
				if err := ie.RateMatchPatternToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 13. rateMatchPatternToReleaseList: sequence-of
	{
		if ie.RateMatchPatternToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigRateMatchPatternToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.RateMatchPatternToReleaseList))); err != nil {
				return err
			}
			for i := range ie.RateMatchPatternToReleaseList {
				if err := ie.RateMatchPatternToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 14. rateMatchPatternGroup1: ref
	{
		if ie.RateMatchPatternGroup1 != nil {
			if err := ie.RateMatchPatternGroup1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 15. rateMatchPatternGroup2: ref
	{
		if ie.RateMatchPatternGroup2 != nil {
			if err := ie.RateMatchPatternGroup2.Encode(e); err != nil {
				return err
			}
		}
	}

	// 16. rbg-Size: enumerated
	{
		if err := e.EncodeEnumerated(ie.Rbg_Size, pDSCHConfigRbgSizeConstraints); err != nil {
			return err
		}
	}

	// 17. mcs-Table: enumerated
	{
		if ie.Mcs_Table != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table, pDSCHConfigMcsTableConstraints); err != nil {
				return err
			}
		}
	}

	// 18. maxNrofCodeWordsScheduledByDCI: enumerated
	{
		if ie.MaxNrofCodeWordsScheduledByDCI != nil {
			if err := e.EncodeEnumerated(*ie.MaxNrofCodeWordsScheduledByDCI, pDSCHConfigMaxNrofCodeWordsScheduledByDCIConstraints); err != nil {
				return err
			}
		}
	}

	// 19. prb-BundlingType: choice
	{
		choiceEnc := e.NewChoiceEncoder(pDSCH_ConfigPrbBundlingTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Prb_BundlingType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Prb_BundlingType.Choice {
		case PDSCH_Config_Prb_BundlingType_StaticBundling:
			c := (*ie.Prb_BundlingType.StaticBundling)
			pDSCHConfigPrbBundlingTypeStaticBundlingSeq := e.NewSequenceEncoder(pDSCHConfigPrbBundlingTypeStaticBundlingConstraints)
			if err := pDSCHConfigPrbBundlingTypeStaticBundlingSeq.EncodePreamble([]bool{c.BundleSize != nil}); err != nil {
				return err
			}
			if c.BundleSize != nil {
				if err := e.EncodeEnumerated((*c.BundleSize), pDSCHConfigPrbBundlingTypeStaticBundlingBundleSizeConstraints); err != nil {
					return err
				}
			}
		case PDSCH_Config_Prb_BundlingType_DynamicBundling:
			c := (*ie.Prb_BundlingType.DynamicBundling)
			pDSCHConfigPrbBundlingTypeDynamicBundlingSeq := e.NewSequenceEncoder(pDSCHConfigPrbBundlingTypeDynamicBundlingConstraints)
			if err := pDSCHConfigPrbBundlingTypeDynamicBundlingSeq.EncodePreamble([]bool{c.BundleSizeSet1 != nil, c.BundleSizeSet2 != nil}); err != nil {
				return err
			}
			if c.BundleSizeSet1 != nil {
				if err := e.EncodeEnumerated((*c.BundleSizeSet1), pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet1Constraints); err != nil {
					return err
				}
			}
			if c.BundleSizeSet2 != nil {
				if err := e.EncodeEnumerated((*c.BundleSizeSet2), pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet2Constraints); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Prb_BundlingType.Choice), Constraint: "index out of range"}
		}
	}

	// 20. zp-CSI-RS-ResourceToAddModList: sequence-of
	{
		if ie.Zp_CSI_RS_ResourceToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigZpCSIRSResourceToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Zp_CSI_RS_ResourceToAddModList))); err != nil {
				return err
			}
			for i := range ie.Zp_CSI_RS_ResourceToAddModList {
				if err := ie.Zp_CSI_RS_ResourceToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 21. zp-CSI-RS-ResourceToReleaseList: sequence-of
	{
		if ie.Zp_CSI_RS_ResourceToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigZpCSIRSResourceToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Zp_CSI_RS_ResourceToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Zp_CSI_RS_ResourceToReleaseList {
				if err := ie.Zp_CSI_RS_ResourceToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 22. aperiodic-ZP-CSI-RS-ResourceSetsToAddModList: sequence-of
	{
		if ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigAperiodicZPCSIRSResourceSetsToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList))); err != nil {
				return err
			}
			for i := range ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList {
				if err := ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 23. aperiodic-ZP-CSI-RS-ResourceSetsToReleaseList: sequence-of
	{
		if ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigAperiodicZPCSIRSResourceSetsToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList {
				if err := ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 24. sp-ZP-CSI-RS-ResourceSetsToAddModList: sequence-of
	{
		if ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigSpZPCSIRSResourceSetsToAddModListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList))); err != nil {
				return err
			}
			for i := range ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList {
				if err := ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 25. sp-ZP-CSI-RS-ResourceSetsToReleaseList: sequence-of
	{
		if ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList != nil {
			seqOf := e.NewSequenceOfEncoder(pDSCHConfigSpZPCSIRSResourceSetsToReleaseListConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList))); err != nil {
				return err
			}
			for i := range ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList {
				if err := ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 26. p-ZP-CSI-RS-ResourceSet: choice
	{
		if ie.P_ZP_CSI_RS_ResourceSet != nil {
			choiceEnc := e.NewChoiceEncoder(pDSCH_ConfigPZPCSIRSResourceSetConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.P_ZP_CSI_RS_ResourceSet).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.P_ZP_CSI_RS_ResourceSet).Choice {
			case PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Setup:
				if err := (*ie.P_ZP_CSI_RS_ResourceSet).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.P_ZP_CSI_RS_ResourceSet).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "maxMIMO-Layers-r16", Optional: true},
					{Name: "minimumSchedulingOffsetK0-r16", Optional: true},
					{Name: "antennaPortsFieldPresenceDCI-1-2-r16", Optional: true},
					{Name: "aperiodicZP-CSI-RS-ResourceSetsToAddModListDCI-1-2-r16", Optional: true},
					{Name: "aperiodicZP-CSI-RS-ResourceSetsToReleaseListDCI-1-2-r16", Optional: true},
					{Name: "dmrs-DownlinkForPDSCH-MappingTypeA-DCI-1-2-r16", Optional: true},
					{Name: "dmrs-DownlinkForPDSCH-MappingTypeB-DCI-1-2-r16", Optional: true},
					{Name: "dmrs-SequenceInitializationDCI-1-2-r16", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-1-2-r16", Optional: true},
					{Name: "mcs-TableDCI-1-2-r16", Optional: true},
					{Name: "numberOfBitsForRV-DCI-1-2-r16", Optional: true},
					{Name: "pdsch-TimeDomainAllocationListDCI-1-2-r16", Optional: true},
					{Name: "prb-BundlingTypeDCI-1-2-r16", Optional: true},
					{Name: "priorityIndicatorDCI-1-2-r16", Optional: true},
					{Name: "rateMatchPatternGroup1DCI-1-2-r16", Optional: true},
					{Name: "rateMatchPatternGroup2DCI-1-2-r16", Optional: true},
					{Name: "resourceAllocationType1GranularityDCI-1-2-r16", Optional: true},
					{Name: "vrb-ToPRB-InterleaverDCI-1-2-r16", Optional: true},
					{Name: "referenceOfSLIVDCI-1-2-r16", Optional: true},
					{Name: "resourceAllocationDCI-1-2-r16", Optional: true},
					{Name: "priorityIndicatorDCI-1-1-r16", Optional: true},
					{Name: "dataScramblingIdentityPDSCH2-r16", Optional: true},
					{Name: "pdsch-TimeDomainAllocationList-r16", Optional: true},
					{Name: "repetitionSchemeConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxMIMO_Layers_r16 != nil, ie.MinimumSchedulingOffsetK0_r16 != nil, ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil, ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 != nil, ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil, ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil, ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil, ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil, ie.Mcs_TableDCI_1_2_r16 != nil, ie.NumberOfBitsForRV_DCI_1_2_r16 != nil, ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil, ie.Prb_BundlingTypeDCI_1_2_r16 != nil, ie.PriorityIndicatorDCI_1_2_r16 != nil, ie.RateMatchPatternGroup1DCI_1_2_r16 != nil, ie.RateMatchPatternGroup2DCI_1_2_r16 != nil, ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil, ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil, ie.ReferenceOfSLIVDCI_1_2_r16 != nil, ie.ResourceAllocationDCI_1_2_r16 != nil, ie.PriorityIndicatorDCI_1_1_r16 != nil, ie.DataScramblingIdentityPDSCH2_r16 != nil, ie.Pdsch_TimeDomainAllocationList_r16 != nil, ie.RepetitionSchemeConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.MaxMIMO_Layers_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtMaxMIMOLayersR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MaxMIMO_Layers_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MaxMIMO_Layers_r16).Choice {
				case PDSCH_Config_Ext_MaxMIMO_Layers_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_MaxMIMO_Layers_r16_Setup:
					if err := (*ie.MaxMIMO_Layers_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MinimumSchedulingOffsetK0_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtMinimumSchedulingOffsetK0R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MinimumSchedulingOffsetK0_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MinimumSchedulingOffsetK0_r16).Choice {
				case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Setup:
					if err := (*ie.MinimumSchedulingOffsetK0_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AntennaPortsFieldPresenceDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AntennaPortsFieldPresenceDCI_1_2_r16, pDSCHConfigExtAntennaPortsFieldPresenceDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDSCHConfigExtAperiodicZPCSIRSResourceSetsToAddModListDCI12R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16))); err != nil {
					return err
				}
				for i := range ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 {
					if err := ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(pDSCHConfigExtAperiodicZPCSIRSResourceSetsToReleaseListDCI12R16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16))); err != nil {
					return err
				}
				for i := range ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 {
					if err := ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeADCI12R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Choice {
				case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Setup:
					if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeBDCI12R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Choice {
				case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Setup:
					if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_SequenceInitializationDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_SequenceInitializationDCI_1_2_r16, pDSCHConfigExtDmrsSequenceInitializationDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_1_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_2_r16, pDSCHConfigHarqProcessNumberSizeDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Mcs_TableDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mcs_TableDCI_1_2_r16, pDSCHConfigExtMcsTableDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.NumberOfBitsForRV_DCI_1_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfBitsForRV_DCI_1_2_r16, pDSCHConfigNumberOfBitsForRVDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschTimeDomainAllocationListDCI12R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Choice {
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Setup:
					if err := (*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Prb_BundlingTypeDCI_1_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPrbBundlingTypeDCI12R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Prb_BundlingTypeDCI_1_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Prb_BundlingTypeDCI_1_2_r16).Choice {
				case PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16:
					c := (*(*ie.Prb_BundlingTypeDCI_1_2_r16).StaticBundling_r16)
					pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Seq := ex.NewSequenceEncoder(pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Constraints)
					if err := pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Seq.EncodePreamble([]bool{c.BundleSize_r16 != nil}); err != nil {
						return err
					}
					if c.BundleSize_r16 != nil {
						if err := ex.EncodeEnumerated((*c.BundleSize_r16), pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16BundleSizeR16Constraints); err != nil {
							return err
						}
					}
				case PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16:
					c := (*(*ie.Prb_BundlingTypeDCI_1_2_r16).DynamicBundling_r16)
					pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq := ex.NewSequenceEncoder(pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Constraints)
					if err := pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq.EncodePreamble([]bool{c.BundleSizeSet1_r16 != nil, c.BundleSizeSet2_r16 != nil}); err != nil {
						return err
					}
					if c.BundleSizeSet1_r16 != nil {
						if err := ex.EncodeEnumerated((*c.BundleSizeSet1_r16), pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet1R16Constraints); err != nil {
							return err
						}
					}
					if c.BundleSizeSet2_r16 != nil {
						if err := ex.EncodeEnumerated((*c.BundleSizeSet2_r16), pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet2R16Constraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.PriorityIndicatorDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorDCI_1_2_r16, pDSCHConfigExtPriorityIndicatorDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.RateMatchPatternGroup1DCI_1_2_r16 != nil {
				if err := ie.RateMatchPatternGroup1DCI_1_2_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.RateMatchPatternGroup2DCI_1_2_r16 != nil {
				if err := ie.RateMatchPatternGroup2DCI_1_2_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.ResourceAllocationType1GranularityDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceAllocationType1GranularityDCI_1_2_r16, pDSCHConfigExtResourceAllocationType1GranularityDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Vrb_ToPRB_InterleaverDCI_1_2_r16, pDSCHConfigExtVrbToPRBInterleaverDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.ReferenceOfSLIVDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ReferenceOfSLIVDCI_1_2_r16, pDSCHConfigExtReferenceOfSLIVDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.ResourceAllocationDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceAllocationDCI_1_2_r16, pDSCHConfigExtResourceAllocationDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorDCI_1_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorDCI_1_1_r16, pDSCHConfigExtPriorityIndicatorDCI11R16Constraints); err != nil {
					return err
				}
			}

			if ie.DataScramblingIdentityPDSCH2_r16 != nil {
				if err := ex.EncodeInteger(*ie.DataScramblingIdentityPDSCH2_r16, pDSCHConfigDataScramblingIdentityPDSCH2R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_TimeDomainAllocationList_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschTimeDomainAllocationListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_TimeDomainAllocationList_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_TimeDomainAllocationList_r16).Choice {
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Setup:
					if err := (*ie.Pdsch_TimeDomainAllocationList_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.RepetitionSchemeConfig_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtRepetitionSchemeConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.RepetitionSchemeConfig_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.RepetitionSchemeConfig_r16).Choice {
				case PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Setup:
					if err := (*ie.RepetitionSchemeConfig_r16).Setup.Encode(ex); err != nil {
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
					{Name: "repetitionSchemeConfig-v1630", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.RepetitionSchemeConfig_v1630 != nil}); err != nil {
				return err
			}

			if ie.RepetitionSchemeConfig_v1630 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtRepetitionSchemeConfigV1630Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.RepetitionSchemeConfig_v1630).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.RepetitionSchemeConfig_v1630).Choice {
				case PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Setup:
					if err := (*ie.RepetitionSchemeConfig_v1630).Setup.Encode(ex); err != nil {
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
					{Name: "pdsch-HARQ-ACK-OneShotFeedbackDCI-1-2-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3DCI-1-2-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3DCI-Field-1-2-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-RetxDCI-1-2-r17", Optional: true},
					{Name: "pucch-sSCellDynDCI-1-2-r17", Optional: true},
					{Name: "dl-OrJointTCI-StateList-r17", Optional: true},
					{Name: "beamAppTime-r17", Optional: true},
					{Name: "dummy", Optional: true},
					{Name: "dmrs-FD-OCC-DisabledForRank1-PDSCH-r17", Optional: true},
					{Name: "minimumSchedulingOffsetK0-r17", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-1-2-v1700", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-1-1-r17", Optional: true},
					{Name: "mcs-Table-r17", Optional: true},
					{Name: "mcs-TableDCI-1-2-r17", Optional: true},
					{Name: "xOverheadMulticast-r17", Optional: true},
					{Name: "priorityIndicatorDCI-4-2-r17", Optional: true},
					{Name: "sizeDCI-4-2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil, ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil, ie.Pucch_SSCellDynDCI_1_2_r17 != nil, ie.Dl_OrJointTCI_StateList_r17 != nil, ie.BeamAppTime_r17 != nil, ie.Dummy != nil, ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil, ie.MinimumSchedulingOffsetK0_r17 != nil, ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil, ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil, ie.Mcs_Table_r17 != nil, ie.Mcs_TableDCI_1_2_r17 != nil, ie.XOverheadMulticast_r17 != nil, ie.PriorityIndicatorDCI_4_2_r17 != nil, ie.SizeDCI_4_2_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17, pDSCHConfigExtPdschHARQACKOneShotFeedbackDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17, pDSCHConfigExtPdschHARQACKEnhType3DCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17, pDSCHConfigExtPdschHARQACKEnhType3DCIField12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17, pDSCHConfigExtPdschHARQACKRetxDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCellDynDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_SSCellDynDCI_1_2_r17, pDSCHConfigExtPucchSSCellDynDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_OrJointTCI_StateList_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtDlOrJointTCIStateListR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_OrJointTCI_StateList_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_OrJointTCI_StateList_r17).Choice {
				case PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_Explicitlist:
					c := (*(*ie.Dl_OrJointTCI_StateList_r17).Explicitlist)
					pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq := ex.NewSequenceEncoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistConstraints)
					if err := pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq.EncodePreamble([]bool{c.Dl_OrJointTCI_StateToAddModList_r17 != nil, c.Dl_OrJointTCI_StateToReleaseList_r17 != nil}); err != nil {
						return err
					}
					if c.Dl_OrJointTCI_StateToAddModList_r17 != nil {
						seqOf := ex.NewSequenceOfEncoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToAddModListR17Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Dl_OrJointTCI_StateToAddModList_r17))); err != nil {
							return err
						}
						for i := range c.Dl_OrJointTCI_StateToAddModList_r17 {
							if err := c.Dl_OrJointTCI_StateToAddModList_r17[i].Encode(ex); err != nil {
								return err
							}
						}
					}
					if c.Dl_OrJointTCI_StateToReleaseList_r17 != nil {
						seqOf := ex.NewSequenceOfEncoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToReleaseListR17Constraints)
						if err := seqOf.EncodeLength(int64(len(c.Dl_OrJointTCI_StateToReleaseList_r17))); err != nil {
							return err
						}
						for i := range c.Dl_OrJointTCI_StateToReleaseList_r17 {
							if err := c.Dl_OrJointTCI_StateToReleaseList_r17[i].Encode(ex); err != nil {
								return err
							}
						}
					}
				case PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_UnifiedTCI_StateRef_r17:
					if err := (*ie.Dl_OrJointTCI_StateList_r17).UnifiedTCI_StateRef_r17.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BeamAppTime_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.BeamAppTime_r17, pDSCHConfigExtBeamAppTimeR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtDummyConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dummy).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dummy).Choice {
				case PDSCH_Config_Ext_Dummy_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Dummy_Setup:
					if err := (*ie.Dummy).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17, pDSCHConfigExtDmrsFDOCCDisabledForRank1PDSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.MinimumSchedulingOffsetK0_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtMinimumSchedulingOffsetK0R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MinimumSchedulingOffsetK0_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MinimumSchedulingOffsetK0_r17).Choice {
				case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Setup:
					if err := (*ie.MinimumSchedulingOffsetK0_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_1_2_v1700 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_2_v1700, pDSCHConfigHarqProcessNumberSizeDCI12V1700Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_1_1_r17 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_1_r17, pDSCHConfigHarqProcessNumberSizeDCI11R17Constraints); err != nil {
					return err
				}
			}

			if ie.Mcs_Table_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mcs_Table_r17, pDSCHConfigExtMcsTableR17Constraints); err != nil {
					return err
				}
			}

			if ie.Mcs_TableDCI_1_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Mcs_TableDCI_1_2_r17, pDSCHConfigExtMcsTableDCI12R17Constraints); err != nil {
					return err
				}
			}

			if ie.XOverheadMulticast_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.XOverheadMulticast_r17, pDSCHConfigExtXOverheadMulticastR17Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorDCI_4_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorDCI_4_2_r17, pDSCHConfigExtPriorityIndicatorDCI42R17Constraints); err != nil {
					return err
				}
			}

			if ie.SizeDCI_4_2_r17 != nil {
				if err := ex.EncodeInteger(*ie.SizeDCI_4_2_r17, pDSCHConfigSizeDCI42R17Constraints); err != nil {
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
					{Name: "pdsch-TimeDomainAllocationListForMultiPDSCH-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Choice {
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Setup:
					if err := (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Setup.Encode(ex); err != nil {
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
					{Name: "advancedReceiver-MU-MIMO-r18", Optional: true},
					{Name: "pdsch-ConfigDCI-1-3-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.AdvancedReceiver_MU_MIMO_r18 != nil, ie.Pdsch_ConfigDCI_1_3_r18 != nil}); err != nil {
				return err
			}

			if ie.AdvancedReceiver_MU_MIMO_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtAdvancedReceiverMUMIMOR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AdvancedReceiver_MU_MIMO_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AdvancedReceiver_MU_MIMO_r18).Choice {
				case PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Setup:
					if err := (*ie.AdvancedReceiver_MU_MIMO_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdsch_ConfigDCI_1_3_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschConfigDCI13R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_ConfigDCI_1_3_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_ConfigDCI_1_3_r18).Choice {
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Setup:
					if err := (*ie.Pdsch_ConfigDCI_1_3_r18).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "pdsch-ConfigDCI-1-3-v1860", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_ConfigDCI_1_3_v1860 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_ConfigDCI_1_3_v1860 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschConfigDCI13V1860Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_ConfigDCI_1_3_v1860).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_ConfigDCI_1_3_v1860).Choice {
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Setup:
					if err := (*ie.Pdsch_ConfigDCI_1_3_v1860).Setup.Encode(ex); err != nil {
						return err
					}
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
					{Name: "mg-CancellationDCI-1-1-r19", Optional: true},
					{Name: "mg-CancellationDCI-1-2-r19", Optional: true},
					{Name: "mg-CancellationDCI-1-3-r19", Optional: true},
					{Name: "pdsch-TimeDomainAllocationListForMultiPDSCH-DCI-1-3-r19", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-1-1-Ext-r19", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-1-2-Ext-r19", Optional: true},
					{Name: "pdsch-ConfigDCI-1-3-v1900", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mg_CancellationDCI_1_1_r19 != nil, ie.Mg_CancellationDCI_1_2_r19 != nil, ie.Mg_CancellationDCI_1_3_r19 != nil, ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19 != nil, ie.Harq_ProcessNumberSizeDCI_1_1_Ext_r19 != nil, ie.Harq_ProcessNumberSizeDCI_1_2_Ext_r19 != nil, ie.Pdsch_ConfigDCI_1_3_v1900 != nil}); err != nil {
				return err
			}

			if ie.Mg_CancellationDCI_1_1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_1_1_r19, pDSCHConfigExtMgCancellationDCI11R19Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_CancellationDCI_1_2_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_1_2_r19, pDSCHConfigExtMgCancellationDCI12R19Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_CancellationDCI_1_3_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_1_3_r19, pDSCHConfigExtMgCancellationDCI13R19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHDCI13R19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Choice {
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Setup:
					if err := (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_1_1_Ext_r19 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_1_Ext_r19, pDSCHConfigHarqProcessNumberSizeDCI11ExtR19Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_1_2_Ext_r19 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_1_2_Ext_r19, pDSCHConfigHarqProcessNumberSizeDCI12ExtR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_ConfigDCI_1_3_v1900 != nil {
				choiceEnc := ex.NewChoiceEncoder(pDSCHConfigExtPdschConfigDCI13V1900Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_ConfigDCI_1_3_v1900).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_ConfigDCI_1_3_v1900).Choice {
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Setup:
					if err := (*ie.Pdsch_ConfigDCI_1_3_v1900).Setup.Encode(ex); err != nil {
						return err
					}
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

func (ie *PDSCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDSCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dataScramblingIdentityPDSCH: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pDSCHConfigDataScramblingIdentityPDSCHConstraints)
			if err != nil {
				return err
			}
			ie.DataScramblingIdentityPDSCH = &v
		}
	}

	// 4. dmrs-DownlinkForPDSCH-MappingTypeA: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Dmrs_DownlinkForPDSCH_MappingTypeA = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_DownlinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeAConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Release:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeA_Setup:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Setup = new(DMRS_DownlinkConfig)
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. dmrs-DownlinkForPDSCH-MappingTypeB: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Dmrs_DownlinkForPDSCH_MappingTypeB = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_DownlinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pDSCH_ConfigDmrsDownlinkForPDSCHMappingTypeBConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Release:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Dmrs_DownlinkForPDSCH_MappingTypeB_Setup:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Setup = new(DMRS_DownlinkConfig)
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. tci-StatesToAddModList: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigTciStatesToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tci_StatesToAddModList = make([]TCI_State, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tci_StatesToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. tci-StatesToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(4) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigTciStatesToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Tci_StatesToReleaseList = make([]TCI_StateId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Tci_StatesToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 8. vrb-ToPRB-Interleaver: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(pDSCHConfigVrbToPRBInterleaverConstraints)
			if err != nil {
				return err
			}
			ie.Vrb_ToPRB_Interleaver = &idx
		}
	}

	// 9. resourceAllocation: enumerated
	{
		v6, err := d.DecodeEnumerated(pDSCHConfigResourceAllocationConstraints)
		if err != nil {
			return err
		}
		ie.ResourceAllocation = v6
	}

	// 10. pdsch-TimeDomainAllocationList: choice
	{
		if seq.IsComponentPresent(7) {
			ie.Pdsch_TimeDomainAllocationList = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_TimeDomainResourceAllocationList
			}{}
			choiceDec := d.NewChoiceDecoder(pDSCH_ConfigPdschTimeDomainAllocationListConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_TimeDomainAllocationList).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDSCH_Config_Pdsch_TimeDomainAllocationList_Release:
				(*ie.Pdsch_TimeDomainAllocationList).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Pdsch_TimeDomainAllocationList_Setup:
				(*ie.Pdsch_TimeDomainAllocationList).Setup = new(PDSCH_TimeDomainResourceAllocationList)
				if err := (*ie.Pdsch_TimeDomainAllocationList).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 11. pdsch-AggregationFactor: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(pDSCHConfigPdschAggregationFactorConstraints)
			if err != nil {
				return err
			}
			ie.Pdsch_AggregationFactor = &idx
		}
	}

	// 12. rateMatchPatternToAddModList: sequence-of
	{
		if seq.IsComponentPresent(9) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigRateMatchPatternToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToAddModList = make([]RateMatchPattern, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 13. rateMatchPatternToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(10) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigRateMatchPatternToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.RateMatchPatternToReleaseList = make([]RateMatchPatternId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.RateMatchPatternToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 14. rateMatchPatternGroup1: ref
	{
		if seq.IsComponentPresent(11) {
			ie.RateMatchPatternGroup1 = new(RateMatchPatternGroup)
			if err := ie.RateMatchPatternGroup1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 15. rateMatchPatternGroup2: ref
	{
		if seq.IsComponentPresent(12) {
			ie.RateMatchPatternGroup2 = new(RateMatchPatternGroup)
			if err := ie.RateMatchPatternGroup2.Decode(d); err != nil {
				return err
			}
		}
	}

	// 16. rbg-Size: enumerated
	{
		v13, err := d.DecodeEnumerated(pDSCHConfigRbgSizeConstraints)
		if err != nil {
			return err
		}
		ie.Rbg_Size = v13
	}

	// 17. mcs-Table: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(pDSCHConfigMcsTableConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table = &idx
		}
	}

	// 18. maxNrofCodeWordsScheduledByDCI: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(pDSCHConfigMaxNrofCodeWordsScheduledByDCIConstraints)
			if err != nil {
				return err
			}
			ie.MaxNrofCodeWordsScheduledByDCI = &idx
		}
	}

	// 19. prb-BundlingType: choice
	{
		choiceDec := d.NewChoiceDecoder(pDSCH_ConfigPrbBundlingTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Prb_BundlingType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case PDSCH_Config_Prb_BundlingType_StaticBundling:
			ie.Prb_BundlingType.StaticBundling = &struct{ BundleSize *int64 }{}
			c := (*ie.Prb_BundlingType.StaticBundling)
			pDSCHConfigPrbBundlingTypeStaticBundlingSeq := d.NewSequenceDecoder(pDSCHConfigPrbBundlingTypeStaticBundlingConstraints)
			if err := pDSCHConfigPrbBundlingTypeStaticBundlingSeq.DecodePreamble(); err != nil {
				return err
			}
			if pDSCHConfigPrbBundlingTypeStaticBundlingSeq.IsComponentPresent(0) {
				c.BundleSize = new(int64)
				v, err := d.DecodeEnumerated(pDSCHConfigPrbBundlingTypeStaticBundlingBundleSizeConstraints)
				if err != nil {
					return err
				}
				(*c.BundleSize) = v
			}
		case PDSCH_Config_Prb_BundlingType_DynamicBundling:
			ie.Prb_BundlingType.DynamicBundling = &struct {
				BundleSizeSet1 *int64
				BundleSizeSet2 *int64
			}{}
			c := (*ie.Prb_BundlingType.DynamicBundling)
			pDSCHConfigPrbBundlingTypeDynamicBundlingSeq := d.NewSequenceDecoder(pDSCHConfigPrbBundlingTypeDynamicBundlingConstraints)
			if err := pDSCHConfigPrbBundlingTypeDynamicBundlingSeq.DecodePreamble(); err != nil {
				return err
			}
			if pDSCHConfigPrbBundlingTypeDynamicBundlingSeq.IsComponentPresent(0) {
				c.BundleSizeSet1 = new(int64)
				v, err := d.DecodeEnumerated(pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet1Constraints)
				if err != nil {
					return err
				}
				(*c.BundleSizeSet1) = v
			}
			if pDSCHConfigPrbBundlingTypeDynamicBundlingSeq.IsComponentPresent(1) {
				c.BundleSizeSet2 = new(int64)
				v, err := d.DecodeEnumerated(pDSCHConfigPrbBundlingTypeDynamicBundlingBundleSizeSet2Constraints)
				if err != nil {
					return err
				}
				(*c.BundleSizeSet2) = v
			}
		}
	}

	// 20. zp-CSI-RS-ResourceToAddModList: sequence-of
	{
		if seq.IsComponentPresent(17) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigZpCSIRSResourceToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Zp_CSI_RS_ResourceToAddModList = make([]ZP_CSI_RS_Resource, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Zp_CSI_RS_ResourceToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 21. zp-CSI-RS-ResourceToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(18) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigZpCSIRSResourceToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Zp_CSI_RS_ResourceToReleaseList = make([]ZP_CSI_RS_ResourceId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Zp_CSI_RS_ResourceToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 22. aperiodic-ZP-CSI-RS-ResourceSetsToAddModList: sequence-of
	{
		if seq.IsComponentPresent(19) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigAperiodicZPCSIRSResourceSetsToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList = make([]ZP_CSI_RS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Aperiodic_ZP_CSI_RS_ResourceSetsToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 23. aperiodic-ZP-CSI-RS-ResourceSetsToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(20) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigAperiodicZPCSIRSResourceSetsToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList = make([]ZP_CSI_RS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 24. sp-ZP-CSI-RS-ResourceSetsToAddModList: sequence-of
	{
		if seq.IsComponentPresent(21) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigSpZPCSIRSResourceSetsToAddModListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList = make([]ZP_CSI_RS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sp_ZP_CSI_RS_ResourceSetsToAddModList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 25. sp-ZP-CSI-RS-ResourceSetsToReleaseList: sequence-of
	{
		if seq.IsComponentPresent(22) {
			seqOf := d.NewSequenceOfDecoder(pDSCHConfigSpZPCSIRSResourceSetsToReleaseListConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList = make([]ZP_CSI_RS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Sp_ZP_CSI_RS_ResourceSetsToReleaseList[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 26. p-ZP-CSI-RS-ResourceSet: choice
	{
		if seq.IsComponentPresent(23) {
			ie.P_ZP_CSI_RS_ResourceSet = &struct {
				Choice  int
				Release *struct{}
				Setup   *ZP_CSI_RS_ResourceSet
			}{}
			choiceDec := d.NewChoiceDecoder(pDSCH_ConfigPZPCSIRSResourceSetConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.P_ZP_CSI_RS_ResourceSet).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Release:
				(*ie.P_ZP_CSI_RS_ResourceSet).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_P_ZP_CSI_RS_ResourceSet_Setup:
				(*ie.P_ZP_CSI_RS_ResourceSet).Setup = new(ZP_CSI_RS_ResourceSet)
				if err := (*ie.P_ZP_CSI_RS_ResourceSet).Setup.Decode(d); err != nil {
					return err
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
				{Name: "maxMIMO-Layers-r16", Optional: true},
				{Name: "minimumSchedulingOffsetK0-r16", Optional: true},
				{Name: "antennaPortsFieldPresenceDCI-1-2-r16", Optional: true},
				{Name: "aperiodicZP-CSI-RS-ResourceSetsToAddModListDCI-1-2-r16", Optional: true},
				{Name: "aperiodicZP-CSI-RS-ResourceSetsToReleaseListDCI-1-2-r16", Optional: true},
				{Name: "dmrs-DownlinkForPDSCH-MappingTypeA-DCI-1-2-r16", Optional: true},
				{Name: "dmrs-DownlinkForPDSCH-MappingTypeB-DCI-1-2-r16", Optional: true},
				{Name: "dmrs-SequenceInitializationDCI-1-2-r16", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-1-2-r16", Optional: true},
				{Name: "mcs-TableDCI-1-2-r16", Optional: true},
				{Name: "numberOfBitsForRV-DCI-1-2-r16", Optional: true},
				{Name: "pdsch-TimeDomainAllocationListDCI-1-2-r16", Optional: true},
				{Name: "prb-BundlingTypeDCI-1-2-r16", Optional: true},
				{Name: "priorityIndicatorDCI-1-2-r16", Optional: true},
				{Name: "rateMatchPatternGroup1DCI-1-2-r16", Optional: true},
				{Name: "rateMatchPatternGroup2DCI-1-2-r16", Optional: true},
				{Name: "resourceAllocationType1GranularityDCI-1-2-r16", Optional: true},
				{Name: "vrb-ToPRB-InterleaverDCI-1-2-r16", Optional: true},
				{Name: "referenceOfSLIVDCI-1-2-r16", Optional: true},
				{Name: "resourceAllocationDCI-1-2-r16", Optional: true},
				{Name: "priorityIndicatorDCI-1-1-r16", Optional: true},
				{Name: "dataScramblingIdentityPDSCH2-r16", Optional: true},
				{Name: "pdsch-TimeDomainAllocationList-r16", Optional: true},
				{Name: "repetitionSchemeConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MaxMIMO_Layers_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MaxMIMO_LayersDL_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtMaxMIMOLayersR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MaxMIMO_Layers_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_MaxMIMO_Layers_r16_Release:
				(*ie.MaxMIMO_Layers_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_MaxMIMO_Layers_r16_Setup:
				(*ie.MaxMIMO_Layers_r16).Setup = new(MaxMIMO_LayersDL_r16)
				if err := (*ie.MaxMIMO_Layers_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.MinimumSchedulingOffsetK0_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MinSchedulingOffsetK0_Values_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtMinimumSchedulingOffsetK0R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MinimumSchedulingOffsetK0_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Release:
				(*ie.MinimumSchedulingOffsetK0_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r16_Setup:
				(*ie.MinimumSchedulingOffsetK0_r16).Setup = new(MinSchedulingOffsetK0_Values_r16)
				if err := (*ie.MinimumSchedulingOffsetK0_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtAntennaPortsFieldPresenceDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.AntennaPortsFieldPresenceDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(pDSCHConfigExtAperiodicZPCSIRSResourceSetsToAddModListDCI12R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 = make([]ZP_CSI_RS_ResourceSet, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(pDSCHConfigExtAperiodicZPCSIRSResourceSetsToReleaseListDCI12R16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 = make([]ZP_CSI_RS_ResourceSetId, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_DownlinkConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeADCI12R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Release:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16_Setup:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Setup = new(DMRS_DownlinkConfig)
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_DownlinkConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtDmrsDownlinkForPDSCHMappingTypeBDCI12R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Release:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16_Setup:
				(*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Setup = new(DMRS_DownlinkConfig)
				if err := (*ie.Dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtDmrsSequenceInitializationDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_SequenceInitializationDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeInteger(pDSCHConfigHarqProcessNumberSizeDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMcsTableDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeInteger(pDSCHConfigNumberOfBitsForRVDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBitsForRV_DCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschTimeDomainAllocationListDCI12R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Release:
				(*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListDCI_1_2_r16_Setup:
				(*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Setup = new(PDSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pdsch_TimeDomainAllocationListDCI_1_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(12) {
			ie.Prb_BundlingTypeDCI_1_2_r16 = &struct {
				Choice              int
				StaticBundling_r16  *struct{ BundleSize_r16 *int64 }
				DynamicBundling_r16 *struct {
					BundleSizeSet1_r16 *int64
					BundleSizeSet2_r16 *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPrbBundlingTypeDCI12R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Prb_BundlingTypeDCI_1_2_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_StaticBundling_r16:
				(*ie.Prb_BundlingTypeDCI_1_2_r16).StaticBundling_r16 = &struct{ BundleSize_r16 *int64 }{}
				c := (*(*ie.Prb_BundlingTypeDCI_1_2_r16).StaticBundling_r16)
				pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Seq := dx.NewSequenceDecoder(pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Constraints)
				if err := pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16Seq.IsComponentPresent(0) {
					c.BundleSize_r16 = new(int64)
					v, err := dx.DecodeEnumerated(pDSCHConfigExtPrbBundlingTypeDCI12R16StaticBundlingR16BundleSizeR16Constraints)
					if err != nil {
						return err
					}
					(*c.BundleSize_r16) = v
				}
			case PDSCH_Config_Ext_Prb_BundlingTypeDCI_1_2_r16_DynamicBundling_r16:
				(*ie.Prb_BundlingTypeDCI_1_2_r16).DynamicBundling_r16 = &struct {
					BundleSizeSet1_r16 *int64
					BundleSizeSet2_r16 *int64
				}{}
				c := (*(*ie.Prb_BundlingTypeDCI_1_2_r16).DynamicBundling_r16)
				pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq := dx.NewSequenceDecoder(pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Constraints)
				if err := pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq.DecodePreamble(); err != nil {
					return err
				}
				if pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq.IsComponentPresent(0) {
					c.BundleSizeSet1_r16 = new(int64)
					v, err := dx.DecodeEnumerated(pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet1R16Constraints)
					if err != nil {
						return err
					}
					(*c.BundleSizeSet1_r16) = v
				}
				if pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16Seq.IsComponentPresent(1) {
					c.BundleSizeSet2_r16 = new(int64)
					v, err := dx.DecodeEnumerated(pDSCHConfigExtPrbBundlingTypeDCI12R16DynamicBundlingR16BundleSizeSet2R16Constraints)
					if err != nil {
						return err
					}
					(*c.BundleSizeSet2_r16) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPriorityIndicatorDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			ie.RateMatchPatternGroup1DCI_1_2_r16 = new(RateMatchPatternGroup)
			if err := ie.RateMatchPatternGroup1DCI_1_2_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(15) {
			ie.RateMatchPatternGroup2DCI_1_2_r16 = new(RateMatchPatternGroup)
			if err := ie.RateMatchPatternGroup2DCI_1_2_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtResourceAllocationType1GranularityDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationType1GranularityDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtVrbToPRBInterleaverDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.Vrb_ToPRB_InterleaverDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtReferenceOfSLIVDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.ReferenceOfSLIVDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtResourceAllocationDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPriorityIndicatorDCI11R16Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_1_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeInteger(pDSCHConfigDataScramblingIdentityPDSCH2R16Constraints)
			if err != nil {
				return err
			}
			ie.DataScramblingIdentityPDSCH2_r16 = &v
		}

		if groupSeq.IsComponentPresent(22) {
			ie.Pdsch_TimeDomainAllocationList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschTimeDomainAllocationListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_TimeDomainAllocationList_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Release:
				(*ie.Pdsch_TimeDomainAllocationList_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationList_r16_Setup:
				(*ie.Pdsch_TimeDomainAllocationList_r16).Setup = new(PDSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pdsch_TimeDomainAllocationList_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(23) {
			ie.RepetitionSchemeConfig_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RepetitionSchemeConfig_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtRepetitionSchemeConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.RepetitionSchemeConfig_r16).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Release:
				(*ie.RepetitionSchemeConfig_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_RepetitionSchemeConfig_r16_Setup:
				(*ie.RepetitionSchemeConfig_r16).Setup = new(RepetitionSchemeConfig_r16)
				if err := (*ie.RepetitionSchemeConfig_r16).Setup.Decode(dx); err != nil {
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
				{Name: "repetitionSchemeConfig-v1630", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.RepetitionSchemeConfig_v1630 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RepetitionSchemeConfig_v1630
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtRepetitionSchemeConfigV1630Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.RepetitionSchemeConfig_v1630).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Release:
				(*ie.RepetitionSchemeConfig_v1630).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_RepetitionSchemeConfig_v1630_Setup:
				(*ie.RepetitionSchemeConfig_v1630).Setup = new(RepetitionSchemeConfig_v1630)
				if err := (*ie.RepetitionSchemeConfig_v1630).Setup.Decode(dx); err != nil {
					return err
				}
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
				{Name: "pdsch-HARQ-ACK-OneShotFeedbackDCI-1-2-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3DCI-1-2-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3DCI-Field-1-2-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-RetxDCI-1-2-r17", Optional: true},
				{Name: "pucch-sSCellDynDCI-1-2-r17", Optional: true},
				{Name: "dl-OrJointTCI-StateList-r17", Optional: true},
				{Name: "beamAppTime-r17", Optional: true},
				{Name: "dummy", Optional: true},
				{Name: "dmrs-FD-OCC-DisabledForRank1-PDSCH-r17", Optional: true},
				{Name: "minimumSchedulingOffsetK0-r17", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-1-2-v1700", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-1-1-r17", Optional: true},
				{Name: "mcs-Table-r17", Optional: true},
				{Name: "mcs-TableDCI-1-2-r17", Optional: true},
				{Name: "xOverheadMulticast-r17", Optional: true},
				{Name: "priorityIndicatorDCI-4-2-r17", Optional: true},
				{Name: "sizeDCI-4-2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPdschHARQACKOneShotFeedbackDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPdschHARQACKEnhType3DCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPdschHARQACKEnhType3DCIField12R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPdschHARQACKRetxDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_RetxDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPucchSSCellDynDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_SSCellDynDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			ie.Dl_OrJointTCI_StateList_r17 = &struct {
				Choice       int
				Explicitlist *struct {
					Dl_OrJointTCI_StateToAddModList_r17  []TCI_State
					Dl_OrJointTCI_StateToReleaseList_r17 []TCI_StateId
				}
				UnifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtDlOrJointTCIStateListR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_OrJointTCI_StateList_r17).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_Explicitlist:
				(*ie.Dl_OrJointTCI_StateList_r17).Explicitlist = &struct {
					Dl_OrJointTCI_StateToAddModList_r17  []TCI_State
					Dl_OrJointTCI_StateToReleaseList_r17 []TCI_StateId
				}{}
				c := (*(*ie.Dl_OrJointTCI_StateList_r17).Explicitlist)
				pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq := dx.NewSequenceDecoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistConstraints)
				if err := pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq.DecodePreamble(); err != nil {
					return err
				}
				if pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq.IsComponentPresent(0) {
					seqOf := dx.NewSequenceOfDecoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToAddModListR17Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Dl_OrJointTCI_StateToAddModList_r17 = make([]TCI_State, n)
					for i := int64(0); i < n; i++ {
						if err := c.Dl_OrJointTCI_StateToAddModList_r17[i].Decode(dx); err != nil {
							return err
						}
					}
				}
				if pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistSeq.IsComponentPresent(1) {
					seqOf := dx.NewSequenceOfDecoder(pDSCHConfigExtDlOrJointTCIStateListR17ExplicitlistDlOrJointTCIStateToReleaseListR17Constraints)
					n, err := seqOf.DecodeLength()
					if err != nil {
						return err
					}
					c.Dl_OrJointTCI_StateToReleaseList_r17 = make([]TCI_StateId, n)
					for i := int64(0); i < n; i++ {
						if err := c.Dl_OrJointTCI_StateToReleaseList_r17[i].Decode(dx); err != nil {
							return err
						}
					}
				}
			case PDSCH_Config_Ext_Dl_OrJointTCI_StateList_r17_UnifiedTCI_StateRef_r17:
				(*ie.Dl_OrJointTCI_StateList_r17).UnifiedTCI_StateRef_r17 = new(ServingCellAndBWP_Id_r17)
				if err := (*ie.Dl_OrJointTCI_StateList_r17).UnifiedTCI_StateRef_r17.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtBeamAppTimeR17Constraints)
			if err != nil {
				return err
			}
			ie.BeamAppTime_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Dummy = &struct {
				Choice  int
				Release *struct{}
				Setup   *Dummy_TDRA_List
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtDummyConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dummy).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Dummy_Release:
				(*ie.Dummy).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Dummy_Setup:
				(*ie.Dummy).Setup = new(Dummy_TDRA_List)
				if err := (*ie.Dummy).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtDmrsFDOCCDisabledForRank1PDSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			ie.MinimumSchedulingOffsetK0_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MinSchedulingOffsetK0_Values_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtMinimumSchedulingOffsetK0R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MinimumSchedulingOffsetK0_r17).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Release:
				(*ie.MinimumSchedulingOffsetK0_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_MinimumSchedulingOffsetK0_r17_Setup:
				(*ie.MinimumSchedulingOffsetK0_r17).Setup = new(MinSchedulingOffsetK0_Values_r17)
				if err := (*ie.MinimumSchedulingOffsetK0_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeInteger(pDSCHConfigHarqProcessNumberSizeDCI12V1700Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_2_v1700 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeInteger(pDSCHConfigHarqProcessNumberSizeDCI11R17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_1_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMcsTableR17Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table_r17 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMcsTableDCI12R17Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableDCI_1_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtXOverheadMulticastR17Constraints)
			if err != nil {
				return err
			}
			ie.XOverheadMulticast_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtPriorityIndicatorDCI42R17Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_4_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeInteger(pDSCHConfigSizeDCI42R17Constraints)
			if err != nil {
				return err
			}
			ie.SizeDCI_4_2_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdsch-TimeDomainAllocationListForMultiPDSCH-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MultiPDSCH_TDRA_List_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Release:
				(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_r17_Setup:
				(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Setup = new(MultiPDSCH_TDRA_List_r17)
				if err := (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_r17).Setup.Decode(dx); err != nil {
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
				{Name: "advancedReceiver-MU-MIMO-r18", Optional: true},
				{Name: "pdsch-ConfigDCI-1-3-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.AdvancedReceiver_MU_MIMO_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *AdvancedReceiver_MU_MIMO_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtAdvancedReceiverMUMIMOR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AdvancedReceiver_MU_MIMO_r18).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Release:
				(*ie.AdvancedReceiver_MU_MIMO_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_AdvancedReceiver_MU_MIMO_r18_Setup:
				(*ie.AdvancedReceiver_MU_MIMO_r18).Setup = new(AdvancedReceiver_MU_MIMO_r18)
				if err := (*ie.AdvancedReceiver_MU_MIMO_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Pdsch_ConfigDCI_1_3_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_ConfigDCI_1_3_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschConfigDCI13R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_ConfigDCI_1_3_r18).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Release:
				(*ie.Pdsch_ConfigDCI_1_3_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_r18_Setup:
				(*ie.Pdsch_ConfigDCI_1_3_r18).Setup = new(PDSCH_ConfigDCI_1_3_r18)
				if err := (*ie.Pdsch_ConfigDCI_1_3_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdsch-ConfigDCI-1-3-v1860", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdsch_ConfigDCI_1_3_v1860 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_ConfigDCI_1_3_v1860
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschConfigDCI13V1860Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_ConfigDCI_1_3_v1860).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Release:
				(*ie.Pdsch_ConfigDCI_1_3_v1860).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1860_Setup:
				(*ie.Pdsch_ConfigDCI_1_3_v1860).Setup = new(PDSCH_ConfigDCI_1_3_v1860)
				if err := (*ie.Pdsch_ConfigDCI_1_3_v1860).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "mg-CancellationDCI-1-1-r19", Optional: true},
				{Name: "mg-CancellationDCI-1-2-r19", Optional: true},
				{Name: "mg-CancellationDCI-1-3-r19", Optional: true},
				{Name: "pdsch-TimeDomainAllocationListForMultiPDSCH-DCI-1-3-r19", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-1-1-Ext-r19", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-1-2-Ext-r19", Optional: true},
				{Name: "pdsch-ConfigDCI-1-3-v1900", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMgCancellationDCI11R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_1_1_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMgCancellationDCI12R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_1_2_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pDSCHConfigExtMgCancellationDCI13R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_1_3_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MultiPDSCH_TDRA_List_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschTimeDomainAllocationListForMultiPDSCHDCI13R19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Release:
				(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19_Setup:
				(*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Setup = new(MultiPDSCH_TDRA_List_r17)
				if err := (*ie.Pdsch_TimeDomainAllocationListForMultiPDSCH_DCI_1_3_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(pDSCHConfigHarqProcessNumberSizeDCI11ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_1_Ext_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeInteger(pDSCHConfigHarqProcessNumberSizeDCI12ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_1_2_Ext_r19 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Pdsch_ConfigDCI_1_3_v1900 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_ConfigDCI_1_3_v1900
			}{}
			choiceDec := dx.NewChoiceDecoder(pDSCHConfigExtPdschConfigDCI13V1900Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_ConfigDCI_1_3_v1900).Choice = int(index)
			switch index {
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Release:
				(*ie.Pdsch_ConfigDCI_1_3_v1900).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PDSCH_Config_Ext_Pdsch_ConfigDCI_1_3_v1900_Setup:
				(*ie.Pdsch_ConfigDCI_1_3_v1900).Setup = new(PDSCH_ConfigDCI_1_3_v1900)
				if err := (*ie.Pdsch_ConfigDCI_1_3_v1900).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
