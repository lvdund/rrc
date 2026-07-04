// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUSCH-Config (line 12336).

var pUSCHConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dataScramblingIdentityPUSCH", Optional: true},
		{Name: "txConfig", Optional: true},
		{Name: "dmrs-UplinkForPUSCH-MappingTypeA", Optional: true},
		{Name: "dmrs-UplinkForPUSCH-MappingTypeB", Optional: true},
		{Name: "pusch-PowerControl", Optional: true},
		{Name: "frequencyHopping", Optional: true},
		{Name: "frequencyHoppingOffsetLists", Optional: true},
		{Name: "resourceAllocation"},
		{Name: "pusch-TimeDomainAllocationList", Optional: true},
		{Name: "pusch-AggregationFactor", Optional: true},
		{Name: "mcs-Table", Optional: true},
		{Name: "mcs-TableTransformPrecoder", Optional: true},
		{Name: "transformPrecoder", Optional: true},
		{Name: "codebookSubset", Optional: true},
		{Name: "maxRank", Optional: true},
		{Name: "rbg-Size", Optional: true},
		{Name: "uci-OnPUSCH", Optional: true},
		{Name: "tp-pi2BPSK", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var pUSCHConfigDataScramblingIdentityPUSCHConstraints = per.Constrained(0, 1023)

const (
	PUSCH_Config_TxConfig_Codebook    = 0
	PUSCH_Config_TxConfig_NonCodebook = 1
)

var pUSCHConfigTxConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_TxConfig_Codebook, PUSCH_Config_TxConfig_NonCodebook},
}

var pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeAConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Release = 0
	PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Setup   = 1
)

var pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeBConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Release = 0
	PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Setup   = 1
)

const (
	PUSCH_Config_FrequencyHopping_IntraSlot = 0
	PUSCH_Config_FrequencyHopping_InterSlot = 1
)

var pUSCHConfigFrequencyHoppingConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_FrequencyHopping_IntraSlot, PUSCH_Config_FrequencyHopping_InterSlot},
}

var pUSCHConfigFrequencyHoppingOffsetListsConstraints = per.SizeRange(1, 4)

const (
	PUSCH_Config_ResourceAllocation_ResourceAllocationType0 = 0
	PUSCH_Config_ResourceAllocation_ResourceAllocationType1 = 1
	PUSCH_Config_ResourceAllocation_DynamicSwitch           = 2
)

var pUSCHConfigResourceAllocationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_ResourceAllocation_ResourceAllocationType0, PUSCH_Config_ResourceAllocation_ResourceAllocationType1, PUSCH_Config_ResourceAllocation_DynamicSwitch},
}

var pUSCH_ConfigPuschTimeDomainAllocationListConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Pusch_TimeDomainAllocationList_Release = 0
	PUSCH_Config_Pusch_TimeDomainAllocationList_Setup   = 1
)

const (
	PUSCH_Config_Pusch_AggregationFactor_N2 = 0
	PUSCH_Config_Pusch_AggregationFactor_N4 = 1
	PUSCH_Config_Pusch_AggregationFactor_N8 = 2
)

var pUSCHConfigPuschAggregationFactorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Pusch_AggregationFactor_N2, PUSCH_Config_Pusch_AggregationFactor_N4, PUSCH_Config_Pusch_AggregationFactor_N8},
}

const (
	PUSCH_Config_Mcs_Table_Qam256     = 0
	PUSCH_Config_Mcs_Table_Qam64LowSE = 1
)

var pUSCHConfigMcsTableConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Mcs_Table_Qam256, PUSCH_Config_Mcs_Table_Qam64LowSE},
}

const (
	PUSCH_Config_Mcs_TableTransformPrecoder_Qam256     = 0
	PUSCH_Config_Mcs_TableTransformPrecoder_Qam64LowSE = 1
)

var pUSCHConfigMcsTableTransformPrecoderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Mcs_TableTransformPrecoder_Qam256, PUSCH_Config_Mcs_TableTransformPrecoder_Qam64LowSE},
}

const (
	PUSCH_Config_TransformPrecoder_Enabled  = 0
	PUSCH_Config_TransformPrecoder_Disabled = 1
)

var pUSCHConfigTransformPrecoderConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_TransformPrecoder_Enabled, PUSCH_Config_TransformPrecoder_Disabled},
}

const (
	PUSCH_Config_CodebookSubset_FullyAndPartialAndNonCoherent = 0
	PUSCH_Config_CodebookSubset_PartialAndNonCoherent         = 1
	PUSCH_Config_CodebookSubset_NonCoherent                   = 2
)

var pUSCHConfigCodebookSubsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_CodebookSubset_FullyAndPartialAndNonCoherent, PUSCH_Config_CodebookSubset_PartialAndNonCoherent, PUSCH_Config_CodebookSubset_NonCoherent},
}

var pUSCHConfigMaxRankConstraints = per.Constrained(1, 4)

const (
	PUSCH_Config_Rbg_Size_Config2 = 0
)

var pUSCHConfigRbgSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Rbg_Size_Config2},
}

var pUSCH_ConfigUciOnPUSCHConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Uci_OnPUSCH_Release = 0
	PUSCH_Config_Uci_OnPUSCH_Setup   = 1
)

const (
	PUSCH_Config_Tp_Pi2BPSK_Enabled = 0
)

var pUSCHConfigTpPi2BPSKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Tp_Pi2BPSK_Enabled},
}

var pUSCHConfigExtMinimumSchedulingOffsetK2R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Release = 0
	PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Setup   = 1
)

var pUSCHConfigExtUlAccessConfigListDCI01R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Release = 0
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Setup   = 1
)

var pUSCHConfigHarqProcessNumberSizeDCI02R16Constraints = per.Constrained(0, 4)

const (
	PUSCH_Config_Ext_Dmrs_SequenceInitializationDCI_0_2_r16_Enabled = 0
)

var pUSCHConfigExtDmrsSequenceInitializationDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Dmrs_SequenceInitializationDCI_0_2_r16_Enabled},
}

var pUSCHConfigNumberOfBitsForRVDCI02R16Constraints = per.Constrained(0, 2)

const (
	PUSCH_Config_Ext_AntennaPortsFieldPresenceDCI_0_2_r16_Enabled = 0
)

var pUSCHConfigExtAntennaPortsFieldPresenceDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_AntennaPortsFieldPresenceDCI_0_2_r16_Enabled},
}

var pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeADCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Release = 0
	PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Setup   = 1
)

var pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeBDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Release = 0
	PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Setup   = 1
)

var pUSCHConfigExtFrequencyHoppingDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pusch-RepTypeA"},
		{Name: "pusch-RepTypeB"},
	},
}

const (
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA = 0
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB = 1
)

var pUSCHConfigExtFrequencyHoppingOffsetListsDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Release = 0
	PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Setup   = 1
)

const (
	PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_FullyAndPartialAndNonCoherent = 0
	PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_PartialAndNonCoherent         = 1
	PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_NonCoherent                   = 2
)

var pUSCHConfigExtCodebookSubsetDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_FullyAndPartialAndNonCoherent, PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_PartialAndNonCoherent, PUSCH_Config_Ext_CodebookSubsetDCI_0_2_r16_NonCoherent},
}

const (
	PUSCH_Config_Ext_InvalidSymbolPatternIndicatorDCI_0_2_r16_Enabled = 0
)

var pUSCHConfigExtInvalidSymbolPatternIndicatorDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_InvalidSymbolPatternIndicatorDCI_0_2_r16_Enabled},
}

var pUSCHConfigMaxRankDCI02R16Constraints = per.Constrained(1, 4)

const (
	PUSCH_Config_Ext_Mcs_TableDCI_0_2_r16_Qam256     = 0
	PUSCH_Config_Ext_Mcs_TableDCI_0_2_r16_Qam64LowSE = 1
)

var pUSCHConfigExtMcsTableDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Mcs_TableDCI_0_2_r16_Qam256, PUSCH_Config_Ext_Mcs_TableDCI_0_2_r16_Qam64LowSE},
}

const (
	PUSCH_Config_Ext_Mcs_TableTransformPrecoderDCI_0_2_r16_Qam256     = 0
	PUSCH_Config_Ext_Mcs_TableTransformPrecoderDCI_0_2_r16_Qam64LowSE = 1
)

var pUSCHConfigExtMcsTableTransformPrecoderDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Mcs_TableTransformPrecoderDCI_0_2_r16_Qam256, PUSCH_Config_Ext_Mcs_TableTransformPrecoderDCI_0_2_r16_Qam64LowSE},
}

const (
	PUSCH_Config_Ext_PriorityIndicatorDCI_0_2_r16_Enabled = 0
)

var pUSCHConfigExtPriorityIndicatorDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_PriorityIndicatorDCI_0_2_r16_Enabled},
}

const (
	PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_2_r16_Pusch_RepTypeA = 0
	PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_2_r16_Pusch_RepTypeB = 1
)

var pUSCHConfigExtPuschRepTypeIndicatorDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_2_r16_Pusch_RepTypeA, PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_2_r16_Pusch_RepTypeB},
}

const (
	PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_ResourceAllocationType0 = 0
	PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_ResourceAllocationType1 = 1
	PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_DynamicSwitch           = 2
)

var pUSCHConfigExtResourceAllocationDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_ResourceAllocationType0, PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_ResourceAllocationType1, PUSCH_Config_Ext_ResourceAllocationDCI_0_2_r16_DynamicSwitch},
}

const (
	PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N2  = 0
	PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N4  = 1
	PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N8  = 2
	PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N16 = 3
)

var pUSCHConfigExtResourceAllocationType1GranularityDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N2, PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N4, PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N8, PUSCH_Config_Ext_ResourceAllocationType1GranularityDCI_0_2_r16_N16},
}

var pUSCHConfigExtUciOnPUSCHListDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Release = 0
	PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Setup   = 1
)

var pUSCHConfigExtPuschTimeDomainAllocationListDCI02R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Release = 0
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Setup   = 1
)

var pUSCHConfigExtPuschTimeDomainAllocationListDCI01R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Release = 0
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Setup   = 1
)

const (
	PUSCH_Config_Ext_InvalidSymbolPatternIndicatorDCI_0_1_r16_Enabled = 0
)

var pUSCHConfigExtInvalidSymbolPatternIndicatorDCI01R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_InvalidSymbolPatternIndicatorDCI_0_1_r16_Enabled},
}

const (
	PUSCH_Config_Ext_PriorityIndicatorDCI_0_1_r16_Enabled = 0
)

var pUSCHConfigExtPriorityIndicatorDCI01R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_PriorityIndicatorDCI_0_1_r16_Enabled},
}

const (
	PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_1_r16_Pusch_RepTypeA = 0
	PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_1_r16_Pusch_RepTypeB = 1
)

var pUSCHConfigExtPuschRepTypeIndicatorDCI01R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_1_r16_Pusch_RepTypeA, PUSCH_Config_Ext_Pusch_RepTypeIndicatorDCI_0_1_r16_Pusch_RepTypeB},
}

const (
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_1_r16_InterRepetition = 0
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_1_r16_InterSlot       = 1
)

var pUSCHConfigExtFrequencyHoppingDCI01R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_FrequencyHoppingDCI_0_1_r16_InterRepetition, PUSCH_Config_Ext_FrequencyHoppingDCI_0_1_r16_InterSlot},
}

var pUSCHConfigExtUciOnPUSCHListDCI01R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Release = 0
	PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Setup   = 1
)

var pUSCHConfigExtPuschPowerControlV1610Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_PowerControl_v1610_Release = 0
	PUSCH_Config_Ext_Pusch_PowerControl_v1610_Setup   = 1
)

const (
	PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_Fullpower      = 0
	PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_FullpowerMode1 = 1
	PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_FullpowerMode2 = 2
)

var pUSCHConfigExtUlFullPowerTransmissionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_Fullpower, PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_FullpowerMode1, PUSCH_Config_Ext_Ul_FullPowerTransmission_r16_FullpowerMode2},
}

var pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Release = 0
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Setup   = 1
)

var pUSCHConfigNumberOfInvalidSymbolsForDLULSwitchingR16Constraints = per.Constrained(1, 4)

var pUSCHConfigExtUlAccessConfigListDCI02R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Release = 0
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Setup   = 1
)

var pUSCHConfigExtBetaOffsetsCrossPri0R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Release = 0
	PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Setup   = 1
)

var pUSCHConfigExtBetaOffsetsCrossPri1R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Release = 0
	PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Setup   = 1
)

var pUSCHConfigExtBetaOffsetsCrossPri0DCI02R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Release = 0
	PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Setup   = 1
)

var pUSCHConfigExtBetaOffsetsCrossPri1DCI02R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Release = 0
	PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Setup   = 1
)

const (
	PUSCH_Config_Ext_MappingPattern_r17_CyclicMapping     = 0
	PUSCH_Config_Ext_MappingPattern_r17_SequentialMapping = 1
)

var pUSCHConfigExtMappingPatternR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_MappingPattern_r17_CyclicMapping, PUSCH_Config_Ext_MappingPattern_r17_SequentialMapping},
}

const (
	PUSCH_Config_Ext_SecondTPCFieldDCI_0_1_r17_Enabled = 0
)

var pUSCHConfigExtSecondTPCFieldDCI01R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_SecondTPCFieldDCI_0_1_r17_Enabled},
}

const (
	PUSCH_Config_Ext_SecondTPCFieldDCI_0_2_r17_Enabled = 0
)

var pUSCHConfigExtSecondTPCFieldDCI02R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_SecondTPCFieldDCI_0_2_r17_Enabled},
}

var pUSCHConfigSequenceOffsetForRVR17Constraints = per.Constrained(0, 3)

var pUSCHConfigExtUlAccessConfigListDCI01R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Release = 0
	PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Setup   = 1
)

var pUSCHConfigExtMinimumSchedulingOffsetK2R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Release = 0
	PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Setup   = 1
)

const (
	PUSCH_Config_Ext_AvailableSlotCounting_r17_Enabled = 0
)

var pUSCHConfigExtAvailableSlotCountingR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_AvailableSlotCounting_r17_Enabled},
}

var pUSCHConfigExtDmrsBundlingPUSCHConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Release = 0
	PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Setup   = 1
)

var pUSCHConfigHarqProcessNumberSizeDCI02V1700Constraints = per.Constrained(5, 5)

var pUSCHConfigHarqProcessNumberSizeDCI01R17Constraints = per.Constrained(5, 5)

var pUSCHConfigExtMpeResourcePoolToAddModListR17Constraints = per.SizeRange(1, common.MaxMPE_Resources_r17)

var pUSCHConfigExtMpeResourcePoolToReleaseListR17Constraints = per.SizeRange(1, common.MaxMPE_Resources_r17)

var pUSCHConfigMaxRankV1810Constraints = per.Constrained(5, 8)

const (
	PUSCH_Config_Ext_STx_2Panel_r18_Enabled = 0
)

var pUSCHConfigExtSTx2PanelR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_STx_2Panel_r18_Enabled},
}

var pUSCHConfigExtCodebookTypeULR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_CodebookTypeUL_r18_Release = 0
	PUSCH_Config_Ext_CodebookTypeUL_r18_Setup   = 1
)

const (
	PUSCH_Config_Ext_ApplyIndicatedTCI_State_r18_First  = 0
	PUSCH_Config_Ext_ApplyIndicatedTCI_State_r18_Second = 1
)

var pUSCHConfigExtApplyIndicatedTCIStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_ApplyIndicatedTCI_State_r18_First, PUSCH_Config_Ext_ApplyIndicatedTCI_State_r18_Second},
}

const (
	PUSCH_Config_Ext_DynamicTransformPrecoderFieldPresenceDCI_0_1_r18_Enabled = 0
)

var pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI01R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_DynamicTransformPrecoderFieldPresenceDCI_0_1_r18_Enabled},
}

const (
	PUSCH_Config_Ext_DynamicTransformPrecoderFieldPresenceDCI_0_2_r18_Enabled = 0
)

var pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI02R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_DynamicTransformPrecoderFieldPresenceDCI_0_2_r18_Enabled},
}

var pUSCHConfigExtPuschConfigDCI03R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Release = 0
	PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Setup   = 1
)

const (
	PUSCH_Config_Ext_Mg_CancellationDCI_0_1_r19_Enabled = 0
)

var pUSCHConfigExtMgCancellationDCI01R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Mg_CancellationDCI_0_1_r19_Enabled},
}

const (
	PUSCH_Config_Ext_Mg_CancellationDCI_0_2_r19_Enabled = 0
)

var pUSCHConfigExtMgCancellationDCI02R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Mg_CancellationDCI_0_2_r19_Enabled},
}

const (
	PUSCH_Config_Ext_Mg_CancellationDCI_0_3_r19_Enabled = 0
)

var pUSCHConfigExtMgCancellationDCI03R19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_Mg_CancellationDCI_0_3_r19_Enabled},
}

var pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHDCI03R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Release = 0
	PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Setup   = 1
)

var pUSCHConfigHarqProcessNumberSizeDCI01ExtR19Constraints = per.Constrained(5, 5)

var pUSCHConfigHarqProcessNumberSizeDCI02ExtR19Constraints = per.Constrained(0, 5)

var pUSCHConfigExtPuschConfigDCI03V1900Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Release = 0
	PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Setup   = 1
)

var pUSCHConfigExtFrequencyHoppingOffsetListsSBFDR19Constraints = per.SizeRange(1, 4)

var pUSCHConfigExtFrequencyHoppingOffsetListsDCI02SBFDR19Constraints = per.SizeRange(1, 4)

const (
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA_IntraSlot = 0
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA_InterSlot = 1
)

var pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA_IntraSlot, PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA_InterSlot},
}

const (
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB_InterRepetition = 0
	PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB_InterSlot       = 1
)

var pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB_InterRepetition, PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB_InterSlot},
}

type PUSCH_Config struct {
	DataScramblingIdentityPUSCH      *int64
	TxConfig                         *int64
	Dmrs_UplinkForPUSCH_MappingTypeA *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_UplinkConfig
	}
	Dmrs_UplinkForPUSCH_MappingTypeB *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_UplinkConfig
	}
	Pusch_PowerControl             *PUSCH_PowerControl
	FrequencyHopping               *int64
	FrequencyHoppingOffsetLists    []int64
	ResourceAllocation             int64
	Pusch_TimeDomainAllocationList *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TimeDomainResourceAllocationList
	}
	Pusch_AggregationFactor    *int64
	Mcs_Table                  *int64
	Mcs_TableTransformPrecoder *int64
	TransformPrecoder          *int64
	CodebookSubset             *int64
	MaxRank                    *int64
	Rbg_Size                   *int64
	Uci_OnPUSCH                *struct {
		Choice  int
		Release *struct{}
		Setup   *UCI_OnPUSCH
	}
	Tp_Pi2BPSK                    *int64
	MinimumSchedulingOffsetK2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *MinSchedulingOffsetK2_Values_r16
	}
	Ul_AccessConfigListDCI_0_1_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_0_1_r16
	}
	Harq_ProcessNumberSizeDCI_0_2_r16            *int64
	Dmrs_SequenceInitializationDCI_0_2_r16       *int64
	NumberOfBitsForRV_DCI_0_2_r16                *int64
	AntennaPortsFieldPresenceDCI_0_2_r16         *int64
	Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_UplinkConfig
	}
	Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_UplinkConfig
	}
	FrequencyHoppingDCI_0_2_r16 *struct {
		Choice         int
		Pusch_RepTypeA *int64
		Pusch_RepTypeB *int64
	}
	FrequencyHoppingOffsetListsDCI_0_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *FrequencyHoppingOffsetListsDCI_0_2_r16
	}
	CodebookSubsetDCI_0_2_r16                     *int64
	InvalidSymbolPatternIndicatorDCI_0_2_r16      *int64
	MaxRankDCI_0_2_r16                            *int64
	Mcs_TableDCI_0_2_r16                          *int64
	Mcs_TableTransformPrecoderDCI_0_2_r16         *int64
	PriorityIndicatorDCI_0_2_r16                  *int64
	Pusch_RepTypeIndicatorDCI_0_2_r16             *int64
	ResourceAllocationDCI_0_2_r16                 *int64
	ResourceAllocationType1GranularityDCI_0_2_r16 *int64
	Uci_OnPUSCH_ListDCI_0_2_r16                   *struct {
		Choice  int
		Release *struct{}
		Setup   *UCI_OnPUSCH_ListDCI_0_2_r16
	}
	Pusch_TimeDomainAllocationListDCI_0_2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TimeDomainResourceAllocationList_r16
	}
	Pusch_TimeDomainAllocationListDCI_0_1_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TimeDomainResourceAllocationList_r16
	}
	InvalidSymbolPatternIndicatorDCI_0_1_r16 *int64
	PriorityIndicatorDCI_0_1_r16             *int64
	Pusch_RepTypeIndicatorDCI_0_1_r16        *int64
	FrequencyHoppingDCI_0_1_r16              *int64
	Uci_OnPUSCH_ListDCI_0_1_r16              *struct {
		Choice  int
		Release *struct{}
		Setup   *UCI_OnPUSCH_ListDCI_0_1_r16
	}
	InvalidSymbolPattern_r16 *InvalidSymbolPattern_r16
	Pusch_PowerControl_v1610 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_PowerControl_v1610
	}
	Ul_FullPowerTransmission_r16                    *int64
	Pusch_TimeDomainAllocationListForMultiPUSCH_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TimeDomainResourceAllocationList_r16
	}
	NumberOfInvalidSymbolsForDL_UL_Switching_r16 *int64
	Ul_AccessConfigListDCI_0_2_r17               *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_0_2_r17
	}
	BetaOffsetsCrossPri0_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSel_r17
	}
	BetaOffsetsCrossPri1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSel_r17
	}
	BetaOffsetsCrossPri0DCI_0_2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSelDCI_0_2_r17
	}
	BetaOffsetsCrossPri1DCI_0_2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *BetaOffsetsCrossPriSelDCI_0_2_r17
	}
	MappingPattern_r17             *int64
	SecondTPCFieldDCI_0_1_r17      *int64
	SecondTPCFieldDCI_0_2_r17      *int64
	SequenceOffsetForRV_r17        *int64
	Ul_AccessConfigListDCI_0_1_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *UL_AccessConfigListDCI_0_1_r17
	}
	MinimumSchedulingOffsetK2_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *MinSchedulingOffsetK2_Values_r17
	}
	AvailableSlotCounting_r17     *int64
	Dmrs_BundlingPUSCH_Config_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *DMRS_BundlingPUSCH_Config_r17
	}
	Harq_ProcessNumberSizeDCI_0_2_v1700 *int64
	Harq_ProcessNumberSizeDCI_0_1_r17   *int64
	Mpe_ResourcePoolToAddModList_r17    []MPE_Resource_r17
	Mpe_ResourcePoolToReleaseList_r17   []MPE_ResourceId_r17
	MaxRank_v1810                       *int64
	STx_2Panel_r18                      *int64
	MultipanelSchemeSDM_r18             *SDM_Scheme_r18
	MultipanelSchemeSFN_r18             *SFN_Scheme_r18
	CodebookTypeUL_r18                  *struct {
		Choice  int
		Release *struct{}
		Setup   *CodebookTypeUL_r18
	}
	ApplyIndicatedTCI_State_r18                      *int64
	DynamicTransformPrecoderFieldPresenceDCI_0_1_r18 *int64
	DynamicTransformPrecoderFieldPresenceDCI_0_2_r18 *int64
	Pusch_ConfigDCI_0_3_r18                          *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_ConfigDCI_0_3_r18
	}
	Mg_CancellationDCI_0_1_r19                              *int64
	Mg_CancellationDCI_0_2_r19                              *int64
	Mg_CancellationDCI_0_3_r19                              *int64
	Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_TimeDomainResourceAllocationList_r16
	}
	Harq_ProcessNumberSizeDCI_0_1_Ext_r19 *int64
	Harq_ProcessNumberSizeDCI_0_2_Ext_r19 *int64
	Pusch_ConfigDCI_0_3_v1900             *struct {
		Choice  int
		Release *struct{}
		Setup   *PUSCH_ConfigDCI_0_3_v1900
	}
	FrequencyHoppingOffsetLists_SBFD_r19        []int64
	FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 []int64
	Pusch_MutingResources_r19                   *PUSCH_MutingResources_r19
}

func (ie *PUSCH_Config) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUSCHConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.MinimumSchedulingOffsetK2_r16 != nil || ie.Ul_AccessConfigListDCI_0_1_r16 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil || ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil || ie.NumberOfBitsForRV_DCI_0_2_r16 != nil || ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil || ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil || ie.FrequencyHoppingDCI_0_2_r16 != nil || ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil || ie.CodebookSubsetDCI_0_2_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil || ie.MaxRankDCI_0_2_r16 != nil || ie.Mcs_TableDCI_0_2_r16 != nil || ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil || ie.PriorityIndicatorDCI_0_2_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil || ie.ResourceAllocationDCI_0_2_r16 != nil || ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil || ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil || ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil || ie.PriorityIndicatorDCI_0_1_r16 != nil || ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil || ie.FrequencyHoppingDCI_0_1_r16 != nil || ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil || ie.InvalidSymbolPattern_r16 != nil || ie.Pusch_PowerControl_v1610 != nil || ie.Ul_FullPowerTransmission_r16 != nil || ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil || ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil
	hasExtGroup1 := ie.Ul_AccessConfigListDCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri0_r17 != nil || ie.BetaOffsetsCrossPri1_r17 != nil || ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil || ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil || ie.MappingPattern_r17 != nil || ie.SecondTPCFieldDCI_0_1_r17 != nil || ie.SecondTPCFieldDCI_0_2_r17 != nil || ie.SequenceOffsetForRV_r17 != nil || ie.Ul_AccessConfigListDCI_0_1_r17 != nil || ie.MinimumSchedulingOffsetK2_r17 != nil || ie.AvailableSlotCounting_r17 != nil || ie.Dmrs_BundlingPUSCH_Config_r17 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil || ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil || ie.Mpe_ResourcePoolToAddModList_r17 != nil || ie.Mpe_ResourcePoolToReleaseList_r17 != nil
	hasExtGroup2 := ie.MaxRank_v1810 != nil || ie.STx_2Panel_r18 != nil || ie.MultipanelSchemeSDM_r18 != nil || ie.MultipanelSchemeSFN_r18 != nil || ie.CodebookTypeUL_r18 != nil || ie.ApplyIndicatedTCI_State_r18 != nil || ie.DynamicTransformPrecoderFieldPresenceDCI_0_1_r18 != nil || ie.DynamicTransformPrecoderFieldPresenceDCI_0_2_r18 != nil || ie.Pusch_ConfigDCI_0_3_r18 != nil
	hasExtGroup3 := ie.Mg_CancellationDCI_0_1_r19 != nil || ie.Mg_CancellationDCI_0_2_r19 != nil || ie.Mg_CancellationDCI_0_3_r19 != nil || ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19 != nil || ie.Harq_ProcessNumberSizeDCI_0_1_Ext_r19 != nil || ie.Harq_ProcessNumberSizeDCI_0_2_Ext_r19 != nil || ie.Pusch_ConfigDCI_0_3_v1900 != nil || ie.FrequencyHoppingOffsetLists_SBFD_r19 != nil || ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 != nil || ie.Pusch_MutingResources_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DataScramblingIdentityPUSCH != nil, ie.TxConfig != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeA != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeB != nil, ie.Pusch_PowerControl != nil, ie.FrequencyHopping != nil, ie.FrequencyHoppingOffsetLists != nil, ie.Pusch_TimeDomainAllocationList != nil, ie.Pusch_AggregationFactor != nil, ie.Mcs_Table != nil, ie.Mcs_TableTransformPrecoder != nil, ie.TransformPrecoder != nil, ie.CodebookSubset != nil, ie.MaxRank != nil, ie.Rbg_Size != nil, ie.Uci_OnPUSCH != nil, ie.Tp_Pi2BPSK != nil}); err != nil {
		return err
	}

	// 3. dataScramblingIdentityPUSCH: integer
	{
		if ie.DataScramblingIdentityPUSCH != nil {
			if err := e.EncodeInteger(*ie.DataScramblingIdentityPUSCH, pUSCHConfigDataScramblingIdentityPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 4. txConfig: enumerated
	{
		if ie.TxConfig != nil {
			if err := e.EncodeEnumerated(*ie.TxConfig, pUSCHConfigTxConfigConstraints); err != nil {
				return err
			}
		}
	}

	// 5. dmrs-UplinkForPUSCH-MappingTypeA: choice
	{
		if ie.Dmrs_UplinkForPUSCH_MappingTypeA != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeAConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Choice {
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Setup:
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 6. dmrs-UplinkForPUSCH-MappingTypeB: choice
	{
		if ie.Dmrs_UplinkForPUSCH_MappingTypeB != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeBConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Choice {
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Setup:
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 7. pusch-PowerControl: ref
	{
		if ie.Pusch_PowerControl != nil {
			if err := ie.Pusch_PowerControl.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. frequencyHopping: enumerated
	{
		if ie.FrequencyHopping != nil {
			if err := e.EncodeEnumerated(*ie.FrequencyHopping, pUSCHConfigFrequencyHoppingConstraints); err != nil {
				return err
			}
		}
	}

	// 9. frequencyHoppingOffsetLists: sequence-of
	{
		if ie.FrequencyHoppingOffsetLists != nil {
			seqOf := e.NewSequenceOfEncoder(pUSCHConfigFrequencyHoppingOffsetListsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.FrequencyHoppingOffsetLists))); err != nil {
				return err
			}
			for i := range ie.FrequencyHoppingOffsetLists {
				if err := e.EncodeInteger(ie.FrequencyHoppingOffsetLists[i], per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
					return err
				}
			}
		}
	}

	// 10. resourceAllocation: enumerated
	{
		if err := e.EncodeEnumerated(ie.ResourceAllocation, pUSCHConfigResourceAllocationConstraints); err != nil {
			return err
		}
	}

	// 11. pusch-TimeDomainAllocationList: choice
	{
		if ie.Pusch_TimeDomainAllocationList != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ConfigPuschTimeDomainAllocationListConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_TimeDomainAllocationList).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Pusch_TimeDomainAllocationList).Choice {
			case PUSCH_Config_Pusch_TimeDomainAllocationList_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Pusch_TimeDomainAllocationList_Setup:
				if err := (*ie.Pusch_TimeDomainAllocationList).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Pusch_TimeDomainAllocationList).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 12. pusch-AggregationFactor: enumerated
	{
		if ie.Pusch_AggregationFactor != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_AggregationFactor, pUSCHConfigPuschAggregationFactorConstraints); err != nil {
				return err
			}
		}
	}

	// 13. mcs-Table: enumerated
	{
		if ie.Mcs_Table != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_Table, pUSCHConfigMcsTableConstraints); err != nil {
				return err
			}
		}
	}

	// 14. mcs-TableTransformPrecoder: enumerated
	{
		if ie.Mcs_TableTransformPrecoder != nil {
			if err := e.EncodeEnumerated(*ie.Mcs_TableTransformPrecoder, pUSCHConfigMcsTableTransformPrecoderConstraints); err != nil {
				return err
			}
		}
	}

	// 15. transformPrecoder: enumerated
	{
		if ie.TransformPrecoder != nil {
			if err := e.EncodeEnumerated(*ie.TransformPrecoder, pUSCHConfigTransformPrecoderConstraints); err != nil {
				return err
			}
		}
	}

	// 16. codebookSubset: enumerated
	{
		if ie.CodebookSubset != nil {
			if err := e.EncodeEnumerated(*ie.CodebookSubset, pUSCHConfigCodebookSubsetConstraints); err != nil {
				return err
			}
		}
	}

	// 17. maxRank: integer
	{
		if ie.MaxRank != nil {
			if err := e.EncodeInteger(*ie.MaxRank, pUSCHConfigMaxRankConstraints); err != nil {
				return err
			}
		}
	}

	// 18. rbg-Size: enumerated
	{
		if ie.Rbg_Size != nil {
			if err := e.EncodeEnumerated(*ie.Rbg_Size, pUSCHConfigRbgSizeConstraints); err != nil {
				return err
			}
		}
	}

	// 19. uci-OnPUSCH: choice
	{
		if ie.Uci_OnPUSCH != nil {
			choiceEnc := e.NewChoiceEncoder(pUSCH_ConfigUciOnPUSCHConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Uci_OnPUSCH).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Uci_OnPUSCH).Choice {
			case PUSCH_Config_Uci_OnPUSCH_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Uci_OnPUSCH_Setup:
				if err := (*ie.Uci_OnPUSCH).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Uci_OnPUSCH).Choice), Constraint: "index out of range"}
			}
		}
	}

	// 20. tp-pi2BPSK: enumerated
	{
		if ie.Tp_Pi2BPSK != nil {
			if err := e.EncodeEnumerated(*ie.Tp_Pi2BPSK, pUSCHConfigTpPi2BPSKConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "minimumSchedulingOffsetK2-r16", Optional: true},
					{Name: "ul-AccessConfigListDCI-0-1-r16", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-0-2-r16", Optional: true},
					{Name: "dmrs-SequenceInitializationDCI-0-2-r16", Optional: true},
					{Name: "numberOfBitsForRV-DCI-0-2-r16", Optional: true},
					{Name: "antennaPortsFieldPresenceDCI-0-2-r16", Optional: true},
					{Name: "dmrs-UplinkForPUSCH-MappingTypeA-DCI-0-2-r16", Optional: true},
					{Name: "dmrs-UplinkForPUSCH-MappingTypeB-DCI-0-2-r16", Optional: true},
					{Name: "frequencyHoppingDCI-0-2-r16", Optional: true},
					{Name: "frequencyHoppingOffsetListsDCI-0-2-r16", Optional: true},
					{Name: "codebookSubsetDCI-0-2-r16", Optional: true},
					{Name: "invalidSymbolPatternIndicatorDCI-0-2-r16", Optional: true},
					{Name: "maxRankDCI-0-2-r16", Optional: true},
					{Name: "mcs-TableDCI-0-2-r16", Optional: true},
					{Name: "mcs-TableTransformPrecoderDCI-0-2-r16", Optional: true},
					{Name: "priorityIndicatorDCI-0-2-r16", Optional: true},
					{Name: "pusch-RepTypeIndicatorDCI-0-2-r16", Optional: true},
					{Name: "resourceAllocationDCI-0-2-r16", Optional: true},
					{Name: "resourceAllocationType1GranularityDCI-0-2-r16", Optional: true},
					{Name: "uci-OnPUSCH-ListDCI-0-2-r16", Optional: true},
					{Name: "pusch-TimeDomainAllocationListDCI-0-2-r16", Optional: true},
					{Name: "pusch-TimeDomainAllocationListDCI-0-1-r16", Optional: true},
					{Name: "invalidSymbolPatternIndicatorDCI-0-1-r16", Optional: true},
					{Name: "priorityIndicatorDCI-0-1-r16", Optional: true},
					{Name: "pusch-RepTypeIndicatorDCI-0-1-r16", Optional: true},
					{Name: "frequencyHoppingDCI-0-1-r16", Optional: true},
					{Name: "uci-OnPUSCH-ListDCI-0-1-r16", Optional: true},
					{Name: "invalidSymbolPattern-r16", Optional: true},
					{Name: "pusch-PowerControl-v1610", Optional: true},
					{Name: "ul-FullPowerTransmission-r16", Optional: true},
					{Name: "pusch-TimeDomainAllocationListForMultiPUSCH-r16", Optional: true},
					{Name: "numberOfInvalidSymbolsForDL-UL-Switching-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MinimumSchedulingOffsetK2_r16 != nil, ie.Ul_AccessConfigListDCI_0_1_r16 != nil, ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil, ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil, ie.NumberOfBitsForRV_DCI_0_2_r16 != nil, ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil, ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil, ie.FrequencyHoppingDCI_0_2_r16 != nil, ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil, ie.CodebookSubsetDCI_0_2_r16 != nil, ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil, ie.MaxRankDCI_0_2_r16 != nil, ie.Mcs_TableDCI_0_2_r16 != nil, ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil, ie.PriorityIndicatorDCI_0_2_r16 != nil, ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil, ie.ResourceAllocationDCI_0_2_r16 != nil, ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil, ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil, ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil, ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil, ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil, ie.PriorityIndicatorDCI_0_1_r16 != nil, ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil, ie.FrequencyHoppingDCI_0_1_r16 != nil, ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil, ie.InvalidSymbolPattern_r16 != nil, ie.Pusch_PowerControl_v1610 != nil, ie.Ul_FullPowerTransmission_r16 != nil, ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil, ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil}); err != nil {
				return err
			}

			if ie.MinimumSchedulingOffsetK2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtMinimumSchedulingOffsetK2R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MinimumSchedulingOffsetK2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MinimumSchedulingOffsetK2_r16).Choice {
				case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Setup:
					if err := (*ie.MinimumSchedulingOffsetK2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_AccessConfigListDCI_0_1_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtUlAccessConfigListDCI01R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_0_1_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_0_1_r16).Choice {
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_0_1_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_0_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_2_r16, pUSCHConfigHarqProcessNumberSizeDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_SequenceInitializationDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Dmrs_SequenceInitializationDCI_0_2_r16, pUSCHConfigExtDmrsSequenceInitializationDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.NumberOfBitsForRV_DCI_0_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfBitsForRV_DCI_0_2_r16, pUSCHConfigNumberOfBitsForRVDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.AntennaPortsFieldPresenceDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AntennaPortsFieldPresenceDCI_0_2_r16, pUSCHConfigExtAntennaPortsFieldPresenceDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeADCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Setup:
					if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeBDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Setup:
					if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FrequencyHoppingDCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtFrequencyHoppingDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.FrequencyHoppingDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.FrequencyHoppingDCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA:
					if err := ex.EncodeEnumerated((*(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeA), pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeAConstraints); err != nil {
						return err
					}
				case PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB:
					if err := ex.EncodeEnumerated((*(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeB), pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeBConstraints); err != nil {
						return err
					}
				}
			}

			if ie.FrequencyHoppingOffsetListsDCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtFrequencyHoppingOffsetListsDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Setup:
					if err := (*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.CodebookSubsetDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CodebookSubsetDCI_0_2_r16, pUSCHConfigExtCodebookSubsetDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.InvalidSymbolPatternIndicatorDCI_0_2_r16, pUSCHConfigExtInvalidSymbolPatternIndicatorDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxRankDCI_0_2_r16 != nil {
				if err := ex.EncodeInteger(*ie.MaxRankDCI_0_2_r16, pUSCHConfigMaxRankDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Mcs_TableDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mcs_TableDCI_0_2_r16, pUSCHConfigExtMcsTableDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Mcs_TableTransformPrecoderDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Mcs_TableTransformPrecoderDCI_0_2_r16, pUSCHConfigExtMcsTableTransformPrecoderDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorDCI_0_2_r16, pUSCHConfigExtPriorityIndicatorDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_RepTypeIndicatorDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_RepTypeIndicatorDCI_0_2_r16, pUSCHConfigExtPuschRepTypeIndicatorDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.ResourceAllocationDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceAllocationDCI_0_2_r16, pUSCHConfigExtResourceAllocationDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.ResourceAllocationType1GranularityDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.ResourceAllocationType1GranularityDCI_0_2_r16, pUSCHConfigExtResourceAllocationType1GranularityDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.Uci_OnPUSCH_ListDCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtUciOnPUSCHListDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Setup:
					if err := (*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschTimeDomainAllocationListDCI02R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Choice {
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Setup:
					if err := (*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschTimeDomainAllocationListDCI01R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Choice {
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Setup:
					if err := (*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.InvalidSymbolPatternIndicatorDCI_0_1_r16, pUSCHConfigExtInvalidSymbolPatternIndicatorDCI01R16Constraints); err != nil {
					return err
				}
			}

			if ie.PriorityIndicatorDCI_0_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.PriorityIndicatorDCI_0_1_r16, pUSCHConfigExtPriorityIndicatorDCI01R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_RepTypeIndicatorDCI_0_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pusch_RepTypeIndicatorDCI_0_1_r16, pUSCHConfigExtPuschRepTypeIndicatorDCI01R16Constraints); err != nil {
					return err
				}
			}

			if ie.FrequencyHoppingDCI_0_1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.FrequencyHoppingDCI_0_1_r16, pUSCHConfigExtFrequencyHoppingDCI01R16Constraints); err != nil {
					return err
				}
			}

			if ie.Uci_OnPUSCH_ListDCI_0_1_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtUciOnPUSCHListDCI01R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Choice {
				case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Setup:
					if err := (*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.InvalidSymbolPattern_r16 != nil {
				if err := ie.InvalidSymbolPattern_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pusch_PowerControl_v1610 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschPowerControlV1610Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_PowerControl_v1610).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_PowerControl_v1610).Choice {
				case PUSCH_Config_Ext_Pusch_PowerControl_v1610_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_PowerControl_v1610_Setup:
					if err := (*ie.Pusch_PowerControl_v1610).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Ul_FullPowerTransmission_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_FullPowerTransmission_r16, pUSCHConfigExtUlFullPowerTransmissionR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Choice {
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Setup:
					if err := (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 != nil {
				if err := ex.EncodeInteger(*ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16, pUSCHConfigNumberOfInvalidSymbolsForDLULSwitchingR16Constraints); err != nil {
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
					{Name: "ul-AccessConfigListDCI-0-2-r17", Optional: true},
					{Name: "betaOffsetsCrossPri0-r17", Optional: true},
					{Name: "betaOffsetsCrossPri1-r17", Optional: true},
					{Name: "betaOffsetsCrossPri0DCI-0-2-r17", Optional: true},
					{Name: "betaOffsetsCrossPri1DCI-0-2-r17", Optional: true},
					{Name: "mappingPattern-r17", Optional: true},
					{Name: "secondTPCFieldDCI-0-1-r17", Optional: true},
					{Name: "secondTPCFieldDCI-0-2-r17", Optional: true},
					{Name: "sequenceOffsetForRV-r17", Optional: true},
					{Name: "ul-AccessConfigListDCI-0-1-r17", Optional: true},
					{Name: "minimumSchedulingOffsetK2-r17", Optional: true},
					{Name: "availableSlotCounting-r17", Optional: true},
					{Name: "dmrs-BundlingPUSCH-Config-r17", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-0-2-v1700", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-0-1-r17", Optional: true},
					{Name: "mpe-ResourcePoolToAddModList-r17", Optional: true},
					{Name: "mpe-ResourcePoolToReleaseList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ul_AccessConfigListDCI_0_2_r17 != nil, ie.BetaOffsetsCrossPri0_r17 != nil, ie.BetaOffsetsCrossPri1_r17 != nil, ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil, ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil, ie.MappingPattern_r17 != nil, ie.SecondTPCFieldDCI_0_1_r17 != nil, ie.SecondTPCFieldDCI_0_2_r17 != nil, ie.SequenceOffsetForRV_r17 != nil, ie.Ul_AccessConfigListDCI_0_1_r17 != nil, ie.MinimumSchedulingOffsetK2_r17 != nil, ie.AvailableSlotCounting_r17 != nil, ie.Dmrs_BundlingPUSCH_Config_r17 != nil, ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil, ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil, ie.Mpe_ResourcePoolToAddModList_r17 != nil, ie.Mpe_ResourcePoolToReleaseList_r17 != nil}); err != nil {
				return err
			}

			if ie.Ul_AccessConfigListDCI_0_2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtUlAccessConfigListDCI02R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_0_2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_0_2_r17).Choice {
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_0_2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BetaOffsetsCrossPri0_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtBetaOffsetsCrossPri0R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsetsCrossPri0_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BetaOffsetsCrossPri0_r17).Choice {
				case PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Setup:
					if err := (*ie.BetaOffsetsCrossPri0_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BetaOffsetsCrossPri1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtBetaOffsetsCrossPri1R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsetsCrossPri1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BetaOffsetsCrossPri1_r17).Choice {
				case PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Setup:
					if err := (*ie.BetaOffsetsCrossPri1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BetaOffsetsCrossPri0DCI_0_2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtBetaOffsetsCrossPri0DCI02R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Choice {
				case PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Setup:
					if err := (*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BetaOffsetsCrossPri1DCI_0_2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtBetaOffsetsCrossPri1DCI02R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Choice {
				case PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Setup:
					if err := (*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MappingPattern_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MappingPattern_r17, pUSCHConfigExtMappingPatternR17Constraints); err != nil {
					return err
				}
			}

			if ie.SecondTPCFieldDCI_0_1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondTPCFieldDCI_0_1_r17, pUSCHConfigExtSecondTPCFieldDCI01R17Constraints); err != nil {
					return err
				}
			}

			if ie.SecondTPCFieldDCI_0_2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SecondTPCFieldDCI_0_2_r17, pUSCHConfigExtSecondTPCFieldDCI02R17Constraints); err != nil {
					return err
				}
			}

			if ie.SequenceOffsetForRV_r17 != nil {
				if err := ex.EncodeInteger(*ie.SequenceOffsetForRV_r17, pUSCHConfigSequenceOffsetForRVR17Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_AccessConfigListDCI_0_1_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtUlAccessConfigListDCI01R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Ul_AccessConfigListDCI_0_1_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Ul_AccessConfigListDCI_0_1_r17).Choice {
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Setup:
					if err := (*ie.Ul_AccessConfigListDCI_0_1_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.MinimumSchedulingOffsetK2_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtMinimumSchedulingOffsetK2R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MinimumSchedulingOffsetK2_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MinimumSchedulingOffsetK2_r17).Choice {
				case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Setup:
					if err := (*ie.MinimumSchedulingOffsetK2_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AvailableSlotCounting_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.AvailableSlotCounting_r17, pUSCHConfigExtAvailableSlotCountingR17Constraints); err != nil {
					return err
				}
			}

			if ie.Dmrs_BundlingPUSCH_Config_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtDmrsBundlingPUSCHConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dmrs_BundlingPUSCH_Config_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dmrs_BundlingPUSCH_Config_r17).Choice {
				case PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Setup:
					if err := (*ie.Dmrs_BundlingPUSCH_Config_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_0_2_v1700 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_2_v1700, pUSCHConfigHarqProcessNumberSizeDCI02V1700Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_0_1_r17 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_1_r17, pUSCHConfigHarqProcessNumberSizeDCI01R17Constraints); err != nil {
					return err
				}
			}

			if ie.Mpe_ResourcePoolToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHConfigExtMpeResourcePoolToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mpe_ResourcePoolToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Mpe_ResourcePoolToAddModList_r17 {
					if err := ie.Mpe_ResourcePoolToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Mpe_ResourcePoolToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHConfigExtMpeResourcePoolToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Mpe_ResourcePoolToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.Mpe_ResourcePoolToReleaseList_r17 {
					if err := ie.Mpe_ResourcePoolToReleaseList_r17[i].Encode(ex); err != nil {
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
					{Name: "maxRank-v1810", Optional: true},
					{Name: "sTx-2Panel-r18", Optional: true},
					{Name: "multipanelSchemeSDM-r18", Optional: true},
					{Name: "multipanelSchemeSFN-r18", Optional: true},
					{Name: "codebookTypeUL-r18", Optional: true},
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
					{Name: "dynamicTransformPrecoderFieldPresenceDCI-0-1-r18", Optional: true},
					{Name: "dynamicTransformPrecoderFieldPresenceDCI-0-2-r18", Optional: true},
					{Name: "pusch-ConfigDCI-0-3-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxRank_v1810 != nil, ie.STx_2Panel_r18 != nil, ie.MultipanelSchemeSDM_r18 != nil, ie.MultipanelSchemeSFN_r18 != nil, ie.CodebookTypeUL_r18 != nil, ie.ApplyIndicatedTCI_State_r18 != nil, ie.DynamicTransformPrecoderFieldPresenceDCI_0_1_r18 != nil, ie.DynamicTransformPrecoderFieldPresenceDCI_0_2_r18 != nil, ie.Pusch_ConfigDCI_0_3_r18 != nil}); err != nil {
				return err
			}

			if ie.MaxRank_v1810 != nil {
				if err := ex.EncodeInteger(*ie.MaxRank_v1810, pUSCHConfigMaxRankV1810Constraints); err != nil {
					return err
				}
			}

			if ie.STx_2Panel_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.STx_2Panel_r18, pUSCHConfigExtSTx2PanelR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultipanelSchemeSDM_r18 != nil {
				if err := ie.MultipanelSchemeSDM_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.MultipanelSchemeSFN_r18 != nil {
				if err := ie.MultipanelSchemeSFN_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookTypeUL_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtCodebookTypeULR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.CodebookTypeUL_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.CodebookTypeUL_r18).Choice {
				case PUSCH_Config_Ext_CodebookTypeUL_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_CodebookTypeUL_r18_Setup:
					if err := (*ie.CodebookTypeUL_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ApplyIndicatedTCI_State_r18, pUSCHConfigExtApplyIndicatedTCIStateR18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicTransformPrecoderFieldPresenceDCI_0_1_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicTransformPrecoderFieldPresenceDCI_0_1_r18, pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI01R18Constraints); err != nil {
					return err
				}
			}

			if ie.DynamicTransformPrecoderFieldPresenceDCI_0_2_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.DynamicTransformPrecoderFieldPresenceDCI_0_2_r18, pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI02R18Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_ConfigDCI_0_3_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschConfigDCI03R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_ConfigDCI_0_3_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_ConfigDCI_0_3_r18).Choice {
				case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Setup:
					if err := (*ie.Pusch_ConfigDCI_0_3_r18).Setup.Encode(ex); err != nil {
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
					{Name: "mg-CancellationDCI-0-1-r19", Optional: true},
					{Name: "mg-CancellationDCI-0-2-r19", Optional: true},
					{Name: "mg-CancellationDCI-0-3-r19", Optional: true},
					{Name: "pusch-TimeDomainAllocationListForMultiPUSCH-DCI-0-3-r19", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-0-1-Ext-r19", Optional: true},
					{Name: "harq-ProcessNumberSizeDCI-0-2-Ext-r19", Optional: true},
					{Name: "pusch-ConfigDCI-0-3-v1900", Optional: true},
					{Name: "frequencyHoppingOffsetLists-SBFD-r19", Optional: true},
					{Name: "frequencyHoppingOffsetListsDCI-0-2-SBFD-r19", Optional: true},
					{Name: "pusch-MutingResources-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mg_CancellationDCI_0_1_r19 != nil, ie.Mg_CancellationDCI_0_2_r19 != nil, ie.Mg_CancellationDCI_0_3_r19 != nil, ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19 != nil, ie.Harq_ProcessNumberSizeDCI_0_1_Ext_r19 != nil, ie.Harq_ProcessNumberSizeDCI_0_2_Ext_r19 != nil, ie.Pusch_ConfigDCI_0_3_v1900 != nil, ie.FrequencyHoppingOffsetLists_SBFD_r19 != nil, ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 != nil, ie.Pusch_MutingResources_r19 != nil}); err != nil {
				return err
			}

			if ie.Mg_CancellationDCI_0_1_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_0_1_r19, pUSCHConfigExtMgCancellationDCI01R19Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_CancellationDCI_0_2_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_0_2_r19, pUSCHConfigExtMgCancellationDCI02R19Constraints); err != nil {
					return err
				}
			}

			if ie.Mg_CancellationDCI_0_3_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.Mg_CancellationDCI_0_3_r19, pUSCHConfigExtMgCancellationDCI03R19Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHDCI03R19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Choice {
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Setup:
					if err := (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_0_1_Ext_r19 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_1_Ext_r19, pUSCHConfigHarqProcessNumberSizeDCI01ExtR19Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ProcessNumberSizeDCI_0_2_Ext_r19 != nil {
				if err := ex.EncodeInteger(*ie.Harq_ProcessNumberSizeDCI_0_2_Ext_r19, pUSCHConfigHarqProcessNumberSizeDCI02ExtR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pusch_ConfigDCI_0_3_v1900 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUSCHConfigExtPuschConfigDCI03V1900Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pusch_ConfigDCI_0_3_v1900).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pusch_ConfigDCI_0_3_v1900).Choice {
				case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Setup:
					if err := (*ie.Pusch_ConfigDCI_0_3_v1900).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.FrequencyHoppingOffsetLists_SBFD_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHConfigExtFrequencyHoppingOffsetListsSBFDR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FrequencyHoppingOffsetLists_SBFD_r19))); err != nil {
					return err
				}
				for i := range ie.FrequencyHoppingOffsetLists_SBFD_r19 {
					if err := ex.EncodeInteger(ie.FrequencyHoppingOffsetLists_SBFD_r19[i], per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
						return err
					}
				}
			}

			if ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(pUSCHConfigExtFrequencyHoppingOffsetListsDCI02SBFDR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19))); err != nil {
					return err
				}
				for i := range ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 {
					if err := ex.EncodeInteger(ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19[i], per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1)); err != nil {
						return err
					}
				}
			}

			if ie.Pusch_MutingResources_r19 != nil {
				if err := ie.Pusch_MutingResources_r19.Encode(ex); err != nil {
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

func (ie *PUSCH_Config) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUSCHConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dataScramblingIdentityPUSCH: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUSCHConfigDataScramblingIdentityPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.DataScramblingIdentityPUSCH = &v
		}
	}

	// 4. txConfig: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(pUSCHConfigTxConfigConstraints)
			if err != nil {
				return err
			}
			ie.TxConfig = &idx
		}
	}

	// 5. dmrs-UplinkForPUSCH-MappingTypeA: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Dmrs_UplinkForPUSCH_MappingTypeA = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_UplinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeAConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Release:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeA_Setup:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Setup = new(DMRS_UplinkConfig)
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeA).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. dmrs-UplinkForPUSCH-MappingTypeB: choice
	{
		if seq.IsComponentPresent(3) {
			ie.Dmrs_UplinkForPUSCH_MappingTypeB = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_UplinkConfig
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ConfigDmrsUplinkForPUSCHMappingTypeBConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Release:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Dmrs_UplinkForPUSCH_MappingTypeB_Setup:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Setup = new(DMRS_UplinkConfig)
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeB).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 7. pusch-PowerControl: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Pusch_PowerControl = new(PUSCH_PowerControl)
			if err := ie.Pusch_PowerControl.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. frequencyHopping: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(pUSCHConfigFrequencyHoppingConstraints)
			if err != nil {
				return err
			}
			ie.FrequencyHopping = &idx
		}
	}

	// 9. frequencyHoppingOffsetLists: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(pUSCHConfigFrequencyHoppingOffsetListsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FrequencyHoppingOffsetLists = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := d.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1))
				if err != nil {
					return err
				}
				ie.FrequencyHoppingOffsetLists[i] = v
			}
		}
	}

	// 10. resourceAllocation: enumerated
	{
		v7, err := d.DecodeEnumerated(pUSCHConfigResourceAllocationConstraints)
		if err != nil {
			return err
		}
		ie.ResourceAllocation = v7
	}

	// 11. pusch-TimeDomainAllocationList: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Pusch_TimeDomainAllocationList = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TimeDomainResourceAllocationList
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ConfigPuschTimeDomainAllocationListConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_TimeDomainAllocationList).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_Config_Pusch_TimeDomainAllocationList_Release:
				(*ie.Pusch_TimeDomainAllocationList).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Pusch_TimeDomainAllocationList_Setup:
				(*ie.Pusch_TimeDomainAllocationList).Setup = new(PUSCH_TimeDomainResourceAllocationList)
				if err := (*ie.Pusch_TimeDomainAllocationList).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 12. pusch-AggregationFactor: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(pUSCHConfigPuschAggregationFactorConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_AggregationFactor = &idx
		}
	}

	// 13. mcs-Table: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(pUSCHConfigMcsTableConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_Table = &idx
		}
	}

	// 14. mcs-TableTransformPrecoder: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(pUSCHConfigMcsTableTransformPrecoderConstraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableTransformPrecoder = &idx
		}
	}

	// 15. transformPrecoder: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(pUSCHConfigTransformPrecoderConstraints)
			if err != nil {
				return err
			}
			ie.TransformPrecoder = &idx
		}
	}

	// 16. codebookSubset: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(pUSCHConfigCodebookSubsetConstraints)
			if err != nil {
				return err
			}
			ie.CodebookSubset = &idx
		}
	}

	// 17. maxRank: integer
	{
		if seq.IsComponentPresent(14) {
			v, err := d.DecodeInteger(pUSCHConfigMaxRankConstraints)
			if err != nil {
				return err
			}
			ie.MaxRank = &v
		}
	}

	// 18. rbg-Size: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(pUSCHConfigRbgSizeConstraints)
			if err != nil {
				return err
			}
			ie.Rbg_Size = &idx
		}
	}

	// 19. uci-OnPUSCH: choice
	{
		if seq.IsComponentPresent(16) {
			ie.Uci_OnPUSCH = &struct {
				Choice  int
				Release *struct{}
				Setup   *UCI_OnPUSCH
			}{}
			choiceDec := d.NewChoiceDecoder(pUSCH_ConfigUciOnPUSCHConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uci_OnPUSCH).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUSCH_Config_Uci_OnPUSCH_Release:
				(*ie.Uci_OnPUSCH).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Uci_OnPUSCH_Setup:
				(*ie.Uci_OnPUSCH).Setup = new(UCI_OnPUSCH)
				if err := (*ie.Uci_OnPUSCH).Setup.Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 20. tp-pi2BPSK: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(pUSCHConfigTpPi2BPSKConstraints)
			if err != nil {
				return err
			}
			ie.Tp_Pi2BPSK = &idx
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
				{Name: "minimumSchedulingOffsetK2-r16", Optional: true},
				{Name: "ul-AccessConfigListDCI-0-1-r16", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-0-2-r16", Optional: true},
				{Name: "dmrs-SequenceInitializationDCI-0-2-r16", Optional: true},
				{Name: "numberOfBitsForRV-DCI-0-2-r16", Optional: true},
				{Name: "antennaPortsFieldPresenceDCI-0-2-r16", Optional: true},
				{Name: "dmrs-UplinkForPUSCH-MappingTypeA-DCI-0-2-r16", Optional: true},
				{Name: "dmrs-UplinkForPUSCH-MappingTypeB-DCI-0-2-r16", Optional: true},
				{Name: "frequencyHoppingDCI-0-2-r16", Optional: true},
				{Name: "frequencyHoppingOffsetListsDCI-0-2-r16", Optional: true},
				{Name: "codebookSubsetDCI-0-2-r16", Optional: true},
				{Name: "invalidSymbolPatternIndicatorDCI-0-2-r16", Optional: true},
				{Name: "maxRankDCI-0-2-r16", Optional: true},
				{Name: "mcs-TableDCI-0-2-r16", Optional: true},
				{Name: "mcs-TableTransformPrecoderDCI-0-2-r16", Optional: true},
				{Name: "priorityIndicatorDCI-0-2-r16", Optional: true},
				{Name: "pusch-RepTypeIndicatorDCI-0-2-r16", Optional: true},
				{Name: "resourceAllocationDCI-0-2-r16", Optional: true},
				{Name: "resourceAllocationType1GranularityDCI-0-2-r16", Optional: true},
				{Name: "uci-OnPUSCH-ListDCI-0-2-r16", Optional: true},
				{Name: "pusch-TimeDomainAllocationListDCI-0-2-r16", Optional: true},
				{Name: "pusch-TimeDomainAllocationListDCI-0-1-r16", Optional: true},
				{Name: "invalidSymbolPatternIndicatorDCI-0-1-r16", Optional: true},
				{Name: "priorityIndicatorDCI-0-1-r16", Optional: true},
				{Name: "pusch-RepTypeIndicatorDCI-0-1-r16", Optional: true},
				{Name: "frequencyHoppingDCI-0-1-r16", Optional: true},
				{Name: "uci-OnPUSCH-ListDCI-0-1-r16", Optional: true},
				{Name: "invalidSymbolPattern-r16", Optional: true},
				{Name: "pusch-PowerControl-v1610", Optional: true},
				{Name: "ul-FullPowerTransmission-r16", Optional: true},
				{Name: "pusch-TimeDomainAllocationListForMultiPUSCH-r16", Optional: true},
				{Name: "numberOfInvalidSymbolsForDL-UL-Switching-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.MinimumSchedulingOffsetK2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MinSchedulingOffsetK2_Values_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtMinimumSchedulingOffsetK2R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MinimumSchedulingOffsetK2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Release:
				(*ie.MinimumSchedulingOffsetK2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r16_Setup:
				(*ie.MinimumSchedulingOffsetK2_r16).Setup = new(MinSchedulingOffsetK2_Values_r16)
				if err := (*ie.MinimumSchedulingOffsetK2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Ul_AccessConfigListDCI_0_1_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_0_1_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtUlAccessConfigListDCI01R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_0_1_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Release:
				(*ie.Ul_AccessConfigListDCI_0_1_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r16_Setup:
				(*ie.Ul_AccessConfigListDCI_0_1_r16).Setup = new(UL_AccessConfigListDCI_0_1_r16)
				if err := (*ie.Ul_AccessConfigListDCI_0_1_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeInteger(pUSCHConfigHarqProcessNumberSizeDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtDmrsSequenceInitializationDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Dmrs_SequenceInitializationDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(pUSCHConfigNumberOfBitsForRVDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBitsForRV_DCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtAntennaPortsFieldPresenceDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.AntennaPortsFieldPresenceDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_UplinkConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeADCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Release:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16_Setup:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Setup = new(DMRS_UplinkConfig)
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeA_DCI_0_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_UplinkConfig
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtDmrsUplinkForPUSCHMappingTypeBDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Release:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16_Setup:
				(*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Setup = new(DMRS_UplinkConfig)
				if err := (*ie.Dmrs_UplinkForPUSCH_MappingTypeB_DCI_0_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(8) {
			ie.FrequencyHoppingDCI_0_2_r16 = &struct {
				Choice         int
				Pusch_RepTypeA *int64
				Pusch_RepTypeB *int64
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtFrequencyHoppingDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FrequencyHoppingDCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeA:
				(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeA = new(int64)
				v, err := dx.DecodeEnumerated(pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeAConstraints)
				if err != nil {
					return err
				}
				(*(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeA) = v
			case PUSCH_Config_Ext_FrequencyHoppingDCI_0_2_r16_Pusch_RepTypeB:
				(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeB = new(int64)
				v, err := dx.DecodeEnumerated(pUSCHConfigExtFrequencyHoppingDCI02R16PuschRepTypeBConstraints)
				if err != nil {
					return err
				}
				(*(*ie.FrequencyHoppingDCI_0_2_r16).Pusch_RepTypeB) = v
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.FrequencyHoppingOffsetListsDCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *FrequencyHoppingOffsetListsDCI_0_2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtFrequencyHoppingOffsetListsDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Release:
				(*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_FrequencyHoppingOffsetListsDCI_0_2_r16_Setup:
				(*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Setup = new(FrequencyHoppingOffsetListsDCI_0_2_r16)
				if err := (*ie.FrequencyHoppingOffsetListsDCI_0_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtCodebookSubsetDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.CodebookSubsetDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtInvalidSymbolPatternIndicatorDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.InvalidSymbolPatternIndicatorDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeInteger(pUSCHConfigMaxRankDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.MaxRankDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMcsTableDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMcsTableTransformPrecoderDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Mcs_TableTransformPrecoderDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtPriorityIndicatorDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtPuschRepTypeIndicatorDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepTypeIndicatorDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtResourceAllocationDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtResourceAllocationType1GranularityDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.ResourceAllocationType1GranularityDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			ie.Uci_OnPUSCH_ListDCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UCI_OnPUSCH_ListDCI_0_2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtUciOnPUSCHListDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Release:
				(*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_2_r16_Setup:
				(*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Setup = new(UCI_OnPUSCH_ListDCI_0_2_r16)
				if err := (*ie.Uci_OnPUSCH_ListDCI_0_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(20) {
			ie.Pusch_TimeDomainAllocationListDCI_0_2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschTimeDomainAllocationListDCI02R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Release:
				(*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_2_r16_Setup:
				(*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Setup = new(PUSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pusch_TimeDomainAllocationListDCI_0_2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(21) {
			ie.Pusch_TimeDomainAllocationListDCI_0_1_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschTimeDomainAllocationListDCI01R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Release:
				(*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListDCI_0_1_r16_Setup:
				(*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Setup = new(PUSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pusch_TimeDomainAllocationListDCI_0_1_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(22) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtInvalidSymbolPatternIndicatorDCI01R16Constraints)
			if err != nil {
				return err
			}
			ie.InvalidSymbolPatternIndicatorDCI_0_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(23) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtPriorityIndicatorDCI01R16Constraints)
			if err != nil {
				return err
			}
			ie.PriorityIndicatorDCI_0_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(24) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtPuschRepTypeIndicatorDCI01R16Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_RepTypeIndicatorDCI_0_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(25) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtFrequencyHoppingDCI01R16Constraints)
			if err != nil {
				return err
			}
			ie.FrequencyHoppingDCI_0_1_r16 = &v
		}

		if groupSeq.IsComponentPresent(26) {
			ie.Uci_OnPUSCH_ListDCI_0_1_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UCI_OnPUSCH_ListDCI_0_1_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtUciOnPUSCHListDCI01R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Release:
				(*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Uci_OnPUSCH_ListDCI_0_1_r16_Setup:
				(*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Setup = new(UCI_OnPUSCH_ListDCI_0_1_r16)
				if err := (*ie.Uci_OnPUSCH_ListDCI_0_1_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(27) {
			ie.InvalidSymbolPattern_r16 = new(InvalidSymbolPattern_r16)
			if err := ie.InvalidSymbolPattern_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(28) {
			ie.Pusch_PowerControl_v1610 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_PowerControl_v1610
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschPowerControlV1610Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_PowerControl_v1610).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_PowerControl_v1610_Release:
				(*ie.Pusch_PowerControl_v1610).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_PowerControl_v1610_Setup:
				(*ie.Pusch_PowerControl_v1610).Setup = new(PUSCH_PowerControl_v1610)
				if err := (*ie.Pusch_PowerControl_v1610).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(29) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtUlFullPowerTransmissionR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_FullPowerTransmission_r16 = &v
		}

		if groupSeq.IsComponentPresent(30) {
			ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Release:
				(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_r16_Setup:
				(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Setup = new(PUSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(31) {
			v, err := dx.DecodeInteger(pUSCHConfigNumberOfInvalidSymbolsForDLULSwitchingR16Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfInvalidSymbolsForDL_UL_Switching_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ul-AccessConfigListDCI-0-2-r17", Optional: true},
				{Name: "betaOffsetsCrossPri0-r17", Optional: true},
				{Name: "betaOffsetsCrossPri1-r17", Optional: true},
				{Name: "betaOffsetsCrossPri0DCI-0-2-r17", Optional: true},
				{Name: "betaOffsetsCrossPri1DCI-0-2-r17", Optional: true},
				{Name: "mappingPattern-r17", Optional: true},
				{Name: "secondTPCFieldDCI-0-1-r17", Optional: true},
				{Name: "secondTPCFieldDCI-0-2-r17", Optional: true},
				{Name: "sequenceOffsetForRV-r17", Optional: true},
				{Name: "ul-AccessConfigListDCI-0-1-r17", Optional: true},
				{Name: "minimumSchedulingOffsetK2-r17", Optional: true},
				{Name: "availableSlotCounting-r17", Optional: true},
				{Name: "dmrs-BundlingPUSCH-Config-r17", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-0-2-v1700", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-0-1-r17", Optional: true},
				{Name: "mpe-ResourcePoolToAddModList-r17", Optional: true},
				{Name: "mpe-ResourcePoolToReleaseList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ul_AccessConfigListDCI_0_2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_0_2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtUlAccessConfigListDCI02R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_0_2_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Release:
				(*ie.Ul_AccessConfigListDCI_0_2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_2_r17_Setup:
				(*ie.Ul_AccessConfigListDCI_0_2_r17).Setup = new(UL_AccessConfigListDCI_0_2_r17)
				if err := (*ie.Ul_AccessConfigListDCI_0_2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.BetaOffsetsCrossPri0_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSel_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtBetaOffsetsCrossPri0R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsetsCrossPri0_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Release:
				(*ie.BetaOffsetsCrossPri0_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_BetaOffsetsCrossPri0_r17_Setup:
				(*ie.BetaOffsetsCrossPri0_r17).Setup = new(BetaOffsetsCrossPriSel_r17)
				if err := (*ie.BetaOffsetsCrossPri0_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.BetaOffsetsCrossPri1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSel_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtBetaOffsetsCrossPri1R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsetsCrossPri1_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Release:
				(*ie.BetaOffsetsCrossPri1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_BetaOffsetsCrossPri1_r17_Setup:
				(*ie.BetaOffsetsCrossPri1_r17).Setup = new(BetaOffsetsCrossPriSel_r17)
				if err := (*ie.BetaOffsetsCrossPri1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.BetaOffsetsCrossPri0DCI_0_2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSelDCI_0_2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtBetaOffsetsCrossPri0DCI02R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Release:
				(*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_BetaOffsetsCrossPri0DCI_0_2_r17_Setup:
				(*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Setup = new(BetaOffsetsCrossPriSelDCI_0_2_r17)
				if err := (*ie.BetaOffsetsCrossPri0DCI_0_2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.BetaOffsetsCrossPri1DCI_0_2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *BetaOffsetsCrossPriSelDCI_0_2_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtBetaOffsetsCrossPri1DCI02R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Release:
				(*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_BetaOffsetsCrossPri1DCI_0_2_r17_Setup:
				(*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Setup = new(BetaOffsetsCrossPriSelDCI_0_2_r17)
				if err := (*ie.BetaOffsetsCrossPri1DCI_0_2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMappingPatternR17Constraints)
			if err != nil {
				return err
			}
			ie.MappingPattern_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtSecondTPCFieldDCI01R17Constraints)
			if err != nil {
				return err
			}
			ie.SecondTPCFieldDCI_0_1_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtSecondTPCFieldDCI02R17Constraints)
			if err != nil {
				return err
			}
			ie.SecondTPCFieldDCI_0_2_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeInteger(pUSCHConfigSequenceOffsetForRVR17Constraints)
			if err != nil {
				return err
			}
			ie.SequenceOffsetForRV_r17 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Ul_AccessConfigListDCI_0_1_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *UL_AccessConfigListDCI_0_1_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtUlAccessConfigListDCI01R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Ul_AccessConfigListDCI_0_1_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Release:
				(*ie.Ul_AccessConfigListDCI_0_1_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Ul_AccessConfigListDCI_0_1_r17_Setup:
				(*ie.Ul_AccessConfigListDCI_0_1_r17).Setup = new(UL_AccessConfigListDCI_0_1_r17)
				if err := (*ie.Ul_AccessConfigListDCI_0_1_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(10) {
			ie.MinimumSchedulingOffsetK2_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MinSchedulingOffsetK2_Values_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtMinimumSchedulingOffsetK2R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MinimumSchedulingOffsetK2_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Release:
				(*ie.MinimumSchedulingOffsetK2_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_MinimumSchedulingOffsetK2_r17_Setup:
				(*ie.MinimumSchedulingOffsetK2_r17).Setup = new(MinSchedulingOffsetK2_Values_r17)
				if err := (*ie.MinimumSchedulingOffsetK2_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtAvailableSlotCountingR17Constraints)
			if err != nil {
				return err
			}
			ie.AvailableSlotCounting_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			ie.Dmrs_BundlingPUSCH_Config_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DMRS_BundlingPUSCH_Config_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtDmrsBundlingPUSCHConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dmrs_BundlingPUSCH_Config_r17).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Release:
				(*ie.Dmrs_BundlingPUSCH_Config_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Dmrs_BundlingPUSCH_Config_r17_Setup:
				(*ie.Dmrs_BundlingPUSCH_Config_r17).Setup = new(DMRS_BundlingPUSCH_Config_r17)
				if err := (*ie.Dmrs_BundlingPUSCH_Config_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeInteger(pUSCHConfigHarqProcessNumberSizeDCI02V1700Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_2_v1700 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeInteger(pUSCHConfigHarqProcessNumberSizeDCI01R17Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_1_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHConfigExtMpeResourcePoolToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mpe_ResourcePoolToAddModList_r17 = make([]MPE_Resource_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mpe_ResourcePoolToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(16) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHConfigExtMpeResourcePoolToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Mpe_ResourcePoolToReleaseList_r17 = make([]MPE_ResourceId_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Mpe_ResourcePoolToReleaseList_r17[i].Decode(dx); err != nil {
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
				{Name: "maxRank-v1810", Optional: true},
				{Name: "sTx-2Panel-r18", Optional: true},
				{Name: "multipanelSchemeSDM-r18", Optional: true},
				{Name: "multipanelSchemeSFN-r18", Optional: true},
				{Name: "codebookTypeUL-r18", Optional: true},
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
				{Name: "dynamicTransformPrecoderFieldPresenceDCI-0-1-r18", Optional: true},
				{Name: "dynamicTransformPrecoderFieldPresenceDCI-0-2-r18", Optional: true},
				{Name: "pusch-ConfigDCI-0-3-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(pUSCHConfigMaxRankV1810Constraints)
			if err != nil {
				return err
			}
			ie.MaxRank_v1810 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtSTx2PanelR18Constraints)
			if err != nil {
				return err
			}
			ie.STx_2Panel_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.MultipanelSchemeSDM_r18 = new(SDM_Scheme_r18)
			if err := ie.MultipanelSchemeSDM_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.MultipanelSchemeSFN_r18 = new(SFN_Scheme_r18)
			if err := ie.MultipanelSchemeSFN_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(4) {
			ie.CodebookTypeUL_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CodebookTypeUL_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtCodebookTypeULR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CodebookTypeUL_r18).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_CodebookTypeUL_r18_Release:
				(*ie.CodebookTypeUL_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_CodebookTypeUL_r18_Setup:
				(*ie.CodebookTypeUL_r18).Setup = new(CodebookTypeUL_r18)
				if err := (*ie.CodebookTypeUL_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtApplyIndicatedTCIStateR18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_State_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI01R18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicTransformPrecoderFieldPresenceDCI_0_1_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtDynamicTransformPrecoderFieldPresenceDCI02R18Constraints)
			if err != nil {
				return err
			}
			ie.DynamicTransformPrecoderFieldPresenceDCI_0_2_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Pusch_ConfigDCI_0_3_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_ConfigDCI_0_3_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschConfigDCI03R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_ConfigDCI_0_3_r18).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Release:
				(*ie.Pusch_ConfigDCI_0_3_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_r18_Setup:
				(*ie.Pusch_ConfigDCI_0_3_r18).Setup = new(PUSCH_ConfigDCI_0_3_r18)
				if err := (*ie.Pusch_ConfigDCI_0_3_r18).Setup.Decode(dx); err != nil {
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
				{Name: "mg-CancellationDCI-0-1-r19", Optional: true},
				{Name: "mg-CancellationDCI-0-2-r19", Optional: true},
				{Name: "mg-CancellationDCI-0-3-r19", Optional: true},
				{Name: "pusch-TimeDomainAllocationListForMultiPUSCH-DCI-0-3-r19", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-0-1-Ext-r19", Optional: true},
				{Name: "harq-ProcessNumberSizeDCI-0-2-Ext-r19", Optional: true},
				{Name: "pusch-ConfigDCI-0-3-v1900", Optional: true},
				{Name: "frequencyHoppingOffsetLists-SBFD-r19", Optional: true},
				{Name: "frequencyHoppingOffsetListsDCI-0-2-SBFD-r19", Optional: true},
				{Name: "pusch-MutingResources-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMgCancellationDCI01R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_0_1_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMgCancellationDCI02R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_0_2_r19 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(pUSCHConfigExtMgCancellationDCI03R19Constraints)
			if err != nil {
				return err
			}
			ie.Mg_CancellationDCI_0_3_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_TimeDomainResourceAllocationList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschTimeDomainAllocationListForMultiPUSCHDCI03R19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Release:
				(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19_Setup:
				(*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Setup = new(PUSCH_TimeDomainResourceAllocationList_r16)
				if err := (*ie.Pusch_TimeDomainAllocationListForMultiPUSCH_DCI_0_3_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeInteger(pUSCHConfigHarqProcessNumberSizeDCI01ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_1_Ext_r19 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeInteger(pUSCHConfigHarqProcessNumberSizeDCI02ExtR19Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ProcessNumberSizeDCI_0_2_Ext_r19 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.Pusch_ConfigDCI_0_3_v1900 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PUSCH_ConfigDCI_0_3_v1900
			}{}
			choiceDec := dx.NewChoiceDecoder(pUSCHConfigExtPuschConfigDCI03V1900Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pusch_ConfigDCI_0_3_v1900).Choice = int(index)
			switch index {
			case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Release:
				(*ie.Pusch_ConfigDCI_0_3_v1900).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUSCH_Config_Ext_Pusch_ConfigDCI_0_3_v1900_Setup:
				(*ie.Pusch_ConfigDCI_0_3_v1900).Setup = new(PUSCH_ConfigDCI_0_3_v1900)
				if err := (*ie.Pusch_ConfigDCI_0_3_v1900).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(7) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHConfigExtFrequencyHoppingOffsetListsSBFDR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FrequencyHoppingOffsetLists_SBFD_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1))
				if err != nil {
					return err
				}
				ie.FrequencyHoppingOffsetLists_SBFD_r19[i] = v
			}
		}

		if groupSeq.IsComponentPresent(8) {
			seqOf := dx.NewSequenceOfDecoder(pUSCHConfigExtFrequencyHoppingOffsetListsDCI02SBFDR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(1, common.MaxNrofPhysicalResourceBlocks_1))
				if err != nil {
					return err
				}
				ie.FrequencyHoppingOffsetListsDCI_0_2_SBFD_r19[i] = v
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Pusch_MutingResources_r19 = new(PUSCH_MutingResources_r19)
			if err := ie.Pusch_MutingResources_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
