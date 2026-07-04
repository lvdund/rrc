// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CSI-ReportConfig (line 7074).

var cSIReportConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "reportConfigId"},
		{Name: "carrier", Optional: true},
		{Name: "resourcesForChannelMeasurement"},
		{Name: "csi-IM-ResourcesForInterference", Optional: true},
		{Name: "nzp-CSI-RS-ResourcesForInterference", Optional: true},
		{Name: "reportConfigType"},
		{Name: "reportQuantity"},
		{Name: "reportFreqConfiguration", Optional: true},
		{Name: "timeRestrictionForChannelMeasurements"},
		{Name: "timeRestrictionForInterferenceMeasurements"},
		{Name: "codebookConfig", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "groupBasedBeamReporting"},
		{Name: "cqi-Table", Optional: true},
		{Name: "subbandSize"},
		{Name: "non-PMI-PortIndication", Optional: true},
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

var cSI_ReportConfigReportConfigTypeConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "periodic"},
		{Name: "semiPersistentOnPUCCH"},
		{Name: "semiPersistentOnPUSCH"},
		{Name: "aperiodic"},
	},
}

const (
	CSI_ReportConfig_ReportConfigType_Periodic              = 0
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUCCH = 1
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH = 2
	CSI_ReportConfig_ReportConfigType_Aperiodic             = 3
)

var cSI_ReportConfigReportQuantityConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "none"},
		{Name: "cri-RI-PMI-CQI"},
		{Name: "cri-RI-i1"},
		{Name: "cri-RI-i1-CQI"},
		{Name: "cri-RI-CQI"},
		{Name: "cri-RSRP"},
		{Name: "ssb-Index-RSRP"},
		{Name: "cri-RI-LI-PMI-CQI"},
	},
}

const (
	CSI_ReportConfig_ReportQuantity_None              = 0
	CSI_ReportConfig_ReportQuantity_Cri_RI_PMI_CQI    = 1
	CSI_ReportConfig_ReportQuantity_Cri_RI_I1         = 2
	CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI     = 3
	CSI_ReportConfig_ReportQuantity_Cri_RI_CQI        = 4
	CSI_ReportConfig_ReportQuantity_Cri_RSRP          = 5
	CSI_ReportConfig_ReportQuantity_Ssb_Index_RSRP    = 6
	CSI_ReportConfig_ReportQuantity_Cri_RI_LI_PMI_CQI = 7
)

const (
	CSI_ReportConfig_TimeRestrictionForChannelMeasurements_Configured    = 0
	CSI_ReportConfig_TimeRestrictionForChannelMeasurements_NotConfigured = 1
)

var cSIReportConfigTimeRestrictionForChannelMeasurementsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_TimeRestrictionForChannelMeasurements_Configured, CSI_ReportConfig_TimeRestrictionForChannelMeasurements_NotConfigured},
}

const (
	CSI_ReportConfig_TimeRestrictionForInterferenceMeasurements_Configured    = 0
	CSI_ReportConfig_TimeRestrictionForInterferenceMeasurements_NotConfigured = 1
)

var cSIReportConfigTimeRestrictionForInterferenceMeasurementsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_TimeRestrictionForInterferenceMeasurements_Configured, CSI_ReportConfig_TimeRestrictionForInterferenceMeasurements_NotConfigured},
}

const (
	CSI_ReportConfig_Dummy_N1 = 0
	CSI_ReportConfig_Dummy_N2 = 1
)

var cSIReportConfigDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Dummy_N1, CSI_ReportConfig_Dummy_N2},
}

var cSI_ReportConfigGroupBasedBeamReportingConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "enabled"},
		{Name: "disabled"},
	},
}

const (
	CSI_ReportConfig_GroupBasedBeamReporting_Enabled  = 0
	CSI_ReportConfig_GroupBasedBeamReporting_Disabled = 1
)

const (
	CSI_ReportConfig_SubbandSize_Value1 = 0
	CSI_ReportConfig_SubbandSize_Value2 = 1
)

var cSIReportConfigSubbandSizeConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_SubbandSize_Value1, CSI_ReportConfig_SubbandSize_Value2},
}

var cSIReportConfigNonPMIPortIndicationConstraints = per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerConfig)

var cSIReportConfigExtReportQuantityR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cri-SINR-r16"},
		{Name: "ssb-Index-SINR-r16"},
	},
}

const (
	CSI_ReportConfig_Ext_ReportQuantity_r16_Cri_SINR_r16       = 0
	CSI_ReportConfig_Ext_ReportQuantity_r16_Ssb_Index_SINR_r16 = 1
)

const (
	CSI_ReportConfig_Ext_Cqi_BitsPerSubband_r17_Bits4 = 0
)

var cSIReportConfigExtCqiBitsPerSubbandR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_Cqi_BitsPerSubband_r17_Bits4},
}

const (
	CSI_ReportConfig_Ext_SharedCMR_r17_Enable = 0
)

var cSIReportConfigExtSharedCMRR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_SharedCMR_r17_Enable},
}

const (
	CSI_ReportConfig_Ext_Csi_ReportMode_r17_Mode1 = 0
	CSI_ReportConfig_Ext_Csi_ReportMode_r17_Mode2 = 1
)

var cSIReportConfigExtCsiReportModeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_Csi_ReportMode_r17_Mode1, CSI_ReportConfig_Ext_Csi_ReportMode_r17_Mode2},
}

const (
	CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N0 = 0
	CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N1 = 1
	CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N2 = 2
)

var cSIReportConfigExtNumberOfSingleTRPCSIMode1R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N0, CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N1, CSI_ReportConfig_Ext_NumberOfSingleTRP_CSI_Mode1_r17_N2},
}

var cSIReportConfigExtReportQuantityR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "cri-RSRP-Index-r17"},
		{Name: "ssb-Index-RSRP-Index-r17"},
		{Name: "cri-SINR-Index-r17"},
		{Name: "ssb-Index-SINR-Index-r17"},
	},
}

const (
	CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_RSRP_Index_r17       = 0
	CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_RSRP_Index_r17 = 1
	CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_SINR_Index_r17       = 2
	CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_SINR_Index_r17 = 3
)

var cSIReportConfigExtCsiReportSubConfigToAddModListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_ReportSubconfigPerCSI_ReportConfig_r18)

var cSIReportConfigExtCsiReportSubConfigToReleaseListR18Constraints = per.SizeRange(1, common.MaxNrofCSI_ReportSubconfigPerCSI_ReportConfig_r18)

const (
	CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N1 = 0
	CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N2 = 1
	CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N3 = 2
	CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N4 = 3
)

var cSIReportConfigExtNrofReportedCLImeasResourcesR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N1, CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N2, CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N3, CSI_ReportConfig_Ext_NrofReportedCLImeasResources_r19_N4},
}

var cSIReportConfigExtPucchCSIResourceListExtR19Constraints = per.SizeRange(1, common.MaxNrofBWPs)

const (
	CSI_ReportConfig_Ext_SymbolType_r19_Sbfd     = 0
	CSI_ReportConfig_Ext_SymbolType_r19_Non_Sbfd = 1
)

var cSIReportConfigExtSymbolTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_SymbolType_r19_Sbfd, CSI_ReportConfig_Ext_SymbolType_r19_Non_Sbfd},
}

const (
	CSI_ReportConfig_Ext_NrofReportedRS_v1900_N6 = 0
	CSI_ReportConfig_Ext_NrofReportedRS_v1900_N8 = 1
)

var cSIReportConfigExtNrofReportedRSV1900Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_NrofReportedRS_v1900_N6, CSI_ReportConfig_Ext_NrofReportedRS_v1900_N8},
}

var cSIReportConfigExtPredictionConfigurationR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "csi-InferencePrediction-r19"},
		{Name: "configurationForBM-PredictionAndDataCollection-r19"},
		{Name: "configurationForBM-Monitoring-r19"},
		{Name: "configurationForCSI-Monitoring-r19"},
	},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_Csi_InferencePrediction_r19                        = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19 = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19                  = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19                 = 3
)

const (
	CSI_ReportConfig_Ext_PortMappingMethod_r19_Method1 = 0
	CSI_ReportConfig_Ext_PortMappingMethod_r19_Method2 = 1
)

var cSIReportConfigExtPortMappingMethodR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PortMappingMethod_r19_Method1, CSI_ReportConfig_Ext_PortMappingMethod_r19_Method2},
}

var cSIReportConfigValueOfMR19Constraints = per.Constrained(1, 4)

var cSIReportConfigExtReportQuantityR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "noneBM-r19"},
		{Name: "noneCSI-r19"},
		{Name: "p-CRI-r19"},
		{Name: "p-SSB-Index-r19"},
		{Name: "p-CRI-RSRP-r19"},
		{Name: "p-SSB-Index-RSRP-r19"},
		{Name: "rs-PAI-r19"},
		{Name: "csi-PAI-r19"},
		{Name: "cjtc-Dd-r19"},
		{Name: "cjtc-F-r19"},
		{Name: "cjtc-P-r19"},
		{Name: "cjtc-Dd-F-r19"},
		{Name: "cli-RSSI-r19"},
		{Name: "cli-SRS-RSRP-r19"},
	},
}

const (
	CSI_ReportConfig_Ext_ReportQuantity_r19_NoneBM_r19           = 0
	CSI_ReportConfig_Ext_ReportQuantity_r19_NoneCSI_r19          = 1
	CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_r19            = 2
	CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_r19      = 3
	CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_RSRP_r19       = 4
	CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_RSRP_r19 = 5
	CSI_ReportConfig_Ext_ReportQuantity_r19_Rs_PAI_r19           = 6
	CSI_ReportConfig_Ext_ReportQuantity_r19_Csi_PAI_r19          = 7
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_r19          = 8
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_F_r19           = 9
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_P_r19           = 10
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_F_r19        = 11
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_RSSI_r19         = 12
	CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_SRS_RSRP_r19     = 13
)

var cSIReportConfigExtCsiReportSubConfigToAddModListExtR19Constraints = per.SizeRange(1, common.MaxNrofCSI_ReportSubconfigPerCSI_ReportConfig_r18)

var cSIReportConfigReportConfigTypePeriodicPucchCSIResourceListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

var cSIReportConfigReportConfigTypeSemiPersistentOnPUCCHPucchCSIResourceListConstraints = per.SizeRange(1, common.MaxNrofBWPs)

const (
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl5   = 0
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl10  = 1
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl20  = 2
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl40  = 3
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl80  = 4
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl160 = 5
	CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl320 = 6
)

var cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotConfigConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl5, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl10, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl20, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl40, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl80, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl160, CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH_ReportSlotConfig_Sl320},
}

var cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotOffsetListConstraints = per.SizeRange(1, common.MaxNrofUL_Allocations)

var cSIReportConfigReportConfigTypeAperiodicReportSlotOffsetListConstraints = per.SizeRange(1, common.MaxNrofUL_Allocations)

var cSIReportConfigReportQuantityCriRII1CQIConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-BundleSizeForCSI", Optional: true},
	},
}

const (
	CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI_Pdsch_BundleSizeForCSI_N2 = 0
	CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI_Pdsch_BundleSizeForCSI_N4 = 1
)

var cSIReportConfigReportQuantityCriRII1CQIPdschBundleSizeForCSIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI_Pdsch_BundleSizeForCSI_N2, CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI_Pdsch_BundleSizeForCSI_N4},
}

var cSIReportConfigReportFreqConfigurationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cqi-FormatIndicator", Optional: true},
		{Name: "pmi-FormatIndicator", Optional: true},
		{Name: "csi-ReportingBand", Optional: true},
	},
}

const (
	CSI_ReportConfig_ReportFreqConfiguration_Cqi_FormatIndicator_WidebandCQI = 0
	CSI_ReportConfig_ReportFreqConfiguration_Cqi_FormatIndicator_SubbandCQI  = 1
)

var cSIReportConfigReportFreqConfigurationCqiFormatIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_ReportFreqConfiguration_Cqi_FormatIndicator_WidebandCQI, CSI_ReportConfig_ReportFreqConfiguration_Cqi_FormatIndicator_SubbandCQI},
}

const (
	CSI_ReportConfig_ReportFreqConfiguration_Pmi_FormatIndicator_WidebandPMI = 0
	CSI_ReportConfig_ReportFreqConfiguration_Pmi_FormatIndicator_SubbandPMI  = 1
)

var cSIReportConfigReportFreqConfigurationPmiFormatIndicatorConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_ReportFreqConfiguration_Pmi_FormatIndicator_WidebandPMI, CSI_ReportConfig_ReportFreqConfiguration_Pmi_FormatIndicator_SubbandPMI},
}

var cSIReportConfigReportFreqConfigurationCsiReportingBandConstraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "subbands3"},
		{Name: "subbands4"},
		{Name: "subbands5"},
		{Name: "subbands6"},
		{Name: "subbands7"},
		{Name: "subbands8"},
		{Name: "subbands9"},
		{Name: "subbands10"},
		{Name: "subbands11"},
		{Name: "subbands12"},
		{Name: "subbands13"},
		{Name: "subbands14"},
		{Name: "subbands15"},
		{Name: "subbands16"},
		{Name: "subbands17"},
		{Name: "subbands18"},
	},
	ExtAlternatives: []per.AlternativeInfo{
		{Name: "subbands19-v1530"},
	},
}

const (
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands3  = 0
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands4  = 1
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands5  = 2
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands6  = 3
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands7  = 4
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands8  = 5
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands9  = 6
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands10 = 7
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands11 = 8
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands12 = 9
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands13 = 10
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands14 = 11
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands15 = 12
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands16 = 13
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands17 = 14
	CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands18 = 15
)

var cSIReportConfigGroupBasedBeamReportingDisabledConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nrofReportedRS", Optional: true},
	},
}

const (
	CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N1 = 0
	CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N2 = 1
	CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N3 = 2
	CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N4 = 3
)

var cSIReportConfigGroupBasedBeamReportingDisabledNrofReportedRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N1, CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N2, CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N3, CSI_ReportConfig_GroupBasedBeamReporting_Disabled_NrofReportedRS_N4},
}

const (
	CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl4  = 0
	CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl8  = 1
	CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl16 = 2
)

var cSIReportConfigExtSemiPersistentOnPUSCHV1530ReportSlotConfigV1530Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl4, CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl8, CSI_ReportConfig_Ext_SemiPersistentOnPUSCH_v1530_ReportSlotConfig_v1530_Sl16},
}

var cSIReportConfigExtSemiPersistentOnPUSCHV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSlotOffsetListDCI-0-2-r16", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-1-r16", Optional: true},
	},
}

var cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI02R16Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI01R16Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtAperiodicV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSlotOffsetListDCI-0-2-r16", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-1-r16", Optional: true},
	},
}

var cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI02R16Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI01R16Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

const (
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N1 = 0
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N2 = 1
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N3 = 2
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N4 = 3
)

var cSIReportConfigExtGroupBasedBeamReportingV1710NrofReportedGroupsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N1, CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N2, CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N3, CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1710_NrofReportedGroups_r17_N4},
}

var cSIReportConfigExtSemiPersistentOnPUSCHV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSlotOffsetList-r17", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-2-r17", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-1-r17", Optional: true},
	},
}

var cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListR17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI02R17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI01R17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtAperiodicV1720Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "reportSlotOffsetList-r17", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-2-r17", Optional: true},
		{Name: "reportSlotOffsetListDCI-0-1-r17", Optional: true},
	},
}

var cSIReportConfigExtAperiodicV1720ReportSlotOffsetListR17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI02R17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

var cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI01R17Constraints = per.SizeRange(1, common.MaxNrofUL_Allocations_r16)

const (
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1800_ReportingMode_r18_JointULDL = 0
	CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1800_ReportingMode_r18_OnlyUL    = 1
)

var cSIReportConfigExtGroupBasedBeamReportingV1800ReportingModeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1800_ReportingMode_r18_JointULDL, CSI_ReportConfig_Ext_GroupBasedBeamReporting_v1800_ReportingMode_r18_OnlyUL},
}

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "resourcesForChannelPrediction-r19"},
		{Name: "associatedIdForChannelPrediction-r19", Optional: true},
		{Name: "associatedIdForChannelMeasurement-r19", Optional: true},
		{Name: "nrofReportedPredictedRS-r19", Optional: true},
		{Name: "nrofTimeInstance-r19", Optional: true},
		{Name: "timeGap-r19", Optional: true},
	},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N1 = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N2 = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N3 = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N4 = 3
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofReportedPredictedRSR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N2, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N3, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofReportedPredictedRS_r19_N4},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N1 = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N2 = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N4 = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N8 = 3
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofTimeInstanceR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N2, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N4, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_NrofTimeInstance_r19_N8},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms10   = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms20   = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms40   = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms80   = 3
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms160  = 4
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare3 = 5
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare2 = 6
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare1 = 7
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19TimeGapR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms10, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms20, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms40, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms80, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Ms160, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare3, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare2, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19_TimeGap_r19_Spare1},
}

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "refToPredictionConfig-r19"},
		{Name: "nrofBestBeamForMonitoring-r19", Optional: true},
		{Name: "nrofTransmissionOccasion-r19", Optional: true},
		{Name: "timeInstanceForRS-PAI-r19", Optional: true},
		{Name: "mappingToResourcesForChannelPrediction-r19", Optional: true},
	},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofBestBeamForMonitoring_r19_N1 = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofBestBeamForMonitoring_r19_N2 = 1
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofBestBeamForMonitoringR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofBestBeamForMonitoring_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofBestBeamForMonitoring_r19_N2},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N1  = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N3  = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N7  = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N15 = 3
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofTransmissionOccasionR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N3, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N7, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_NrofTransmissionOccasion_r19_N15},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N1 = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N2 = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N3 = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N4 = 3
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N5 = 4
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N6 = 5
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N7 = 6
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N8 = 7
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19TimeInstanceForRSPAIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N2, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N3, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N4, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N5, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N6, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N7, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19_TimeInstanceForRS_PAI_r19_N8},
}

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "refToPredictionConfig-r19"},
		{Name: "timeInstanceForCSI-PAI-r19", Optional: true},
	},
}

const (
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N1 = 0
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N2 = 1
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N3 = 2
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N4 = 3
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N5 = 4
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N6 = 5
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N7 = 6
	CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N8 = 7
)

var cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19TimeInstanceForCSIPAIR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N1, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N2, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N3, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N4, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N5, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N6, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N7, CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19_TimeInstanceForCSI_PAI_r19_N8},
}

type CSI_ReportConfig struct {
	ReportConfigId                      CSI_ReportConfigId
	Carrier                             *ServCellIndex
	ResourcesForChannelMeasurement      CSI_ResourceConfigId
	Csi_IM_ResourcesForInterference     *CSI_ResourceConfigId
	Nzp_CSI_RS_ResourcesForInterference *CSI_ResourceConfigId
	ReportConfigType                    struct {
		Choice   int
		Periodic *struct {
			ReportSlotConfig       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUCCH *struct {
			ReportSlotConfig       CSI_ReportPeriodicityAndOffset
			Pucch_CSI_ResourceList []PUCCH_CSI_Resource
		}
		SemiPersistentOnPUSCH *struct {
			ReportSlotConfig     int64
			ReportSlotOffsetList []int64
			P0alpha              P0_PUSCH_AlphaSetId
		}
		Aperiodic *struct{ ReportSlotOffsetList []int64 }
	}
	ReportQuantity struct {
		Choice            int
		None              *struct{}
		Cri_RI_PMI_CQI    *struct{}
		Cri_RI_I1         *struct{}
		Cri_RI_I1_CQI     *struct{ Pdsch_BundleSizeForCSI *int64 }
		Cri_RI_CQI        *struct{}
		Cri_RSRP          *struct{}
		Ssb_Index_RSRP    *struct{}
		Cri_RI_LI_PMI_CQI *struct{}
	}
	ReportFreqConfiguration *struct {
		Cqi_FormatIndicator *int64
		Pmi_FormatIndicator *int64
		Csi_ReportingBand   *struct {
			Choice     int
			Subbands3  *per.BitString
			Subbands4  *per.BitString
			Subbands5  *per.BitString
			Subbands6  *per.BitString
			Subbands7  *per.BitString
			Subbands8  *per.BitString
			Subbands9  *per.BitString
			Subbands10 *per.BitString
			Subbands11 *per.BitString
			Subbands12 *per.BitString
			Subbands13 *per.BitString
			Subbands14 *per.BitString
			Subbands15 *per.BitString
			Subbands16 *per.BitString
			Subbands17 *per.BitString
			Subbands18 *per.BitString
		}
	}
	TimeRestrictionForChannelMeasurements      int64
	TimeRestrictionForInterferenceMeasurements int64
	CodebookConfig                             *CodebookConfig
	Dummy                                      *int64
	GroupBasedBeamReporting                    struct {
		Choice   int
		Enabled  *struct{}
		Disabled *struct{ NrofReportedRS *int64 }
	}
	Cqi_Table                   *CQI_Table
	SubbandSize                 int64
	Non_PMI_PortIndication      []PortIndexFor8Ranks
	SemiPersistentOnPUSCH_v1530 *struct{ ReportSlotConfig_v1530 int64 }
	SemiPersistentOnPUSCH_v1610 *struct {
		ReportSlotOffsetListDCI_0_2_r16 []int64
		ReportSlotOffsetListDCI_0_1_r16 []int64
	}
	Aperiodic_v1610 *struct {
		ReportSlotOffsetListDCI_0_2_r16 []int64
		ReportSlotOffsetListDCI_0_1_r16 []int64
	}
	ReportQuantity_r16 *struct {
		Choice             int
		Cri_SINR_r16       *struct{}
		Ssb_Index_SINR_r16 *struct{}
	}
	CodebookConfig_r16              *CodebookConfig_r16
	Cqi_BitsPerSubband_r17          *int64
	GroupBasedBeamReporting_v1710   *struct{ NrofReportedGroups_r17 int64 }
	CodebookConfig_r17              *CodebookConfig_r17
	SharedCMR_r17                   *int64
	Csi_ReportMode_r17              *int64
	NumberOfSingleTRP_CSI_Mode1_r17 *int64
	ReportQuantity_r17              *struct {
		Choice                   int
		Cri_RSRP_Index_r17       *struct{}
		Ssb_Index_RSRP_Index_r17 *struct{}
		Cri_SINR_Index_r17       *struct{}
		Ssb_Index_SINR_Index_r17 *struct{}
	}
	SemiPersistentOnPUSCH_v1720 *struct {
		ReportSlotOffsetList_r17        []int64
		ReportSlotOffsetListDCI_0_2_r17 []int64
		ReportSlotOffsetListDCI_0_1_r17 []int64
	}
	Aperiodic_v1720 *struct {
		ReportSlotOffsetList_r17        []int64
		ReportSlotOffsetListDCI_0_2_r17 []int64
		ReportSlotOffsetListDCI_0_1_r17 []int64
	}
	CodebookConfig_v1730                 *CodebookConfig_v1730
	GroupBasedBeamReporting_v1800        *struct{ ReportingMode_r18 int64 }
	ReportQuantity_r18                   *TDCP_r18
	CodebookConfig_r18                   *CodebookConfig_r18
	Csi_ReportSubConfigToAddModList_r18  []CSI_ReportSubConfig_r18
	Csi_ReportSubConfigToReleaseList_r18 []CSI_ReportSubConfigId_r18
	NrofReportedCLImeasResources_r19     *int64
	Pucch_CSI_ResourceListExt_r19        []PUCCH_CSI_ResourceExt_v1900
	SymbolType_r19                       *int64
	NrofReportedRS_v1900                 *int64
	PredictionConfiguration_r19          *struct {
		Choice                                             int
		Csi_InferencePrediction_r19                        *struct{}
		ConfigurationForBM_PredictionAndDataCollection_r19 *struct {
			ResourcesForChannelPrediction_r19     CSI_ResourceConfigId
			AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
			AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
			NrofReportedPredictedRS_r19           *int64
			NrofTimeInstance_r19                  *int64
			TimeGap_r19                           *int64
		}
		ConfigurationForBM_Monitoring_r19 *struct {
			RefToPredictionConfig_r19                  CSI_ReportConfigId
			NrofBestBeamForMonitoring_r19              *int64
			NrofTransmissionOccasion_r19               *int64
			TimeInstanceForRS_PAI_r19                  *int64
			MappingToResourcesForChannelPrediction_r19 *per.BitString
		}
		ConfigurationForCSI_Monitoring_r19 *struct {
			RefToPredictionConfig_r19  CSI_ReportConfigId
			TimeInstanceForCSI_PAI_r19 *int64
		}
	}
	CodebookConfig_r19    *CodebookConfig_r19
	PortMappingMethod_r19 *int64
	ValueOfM_r19          *int64
	ReportQuantity_r19    *struct {
		Choice               int
		NoneBM_r19           *struct{}
		NoneCSI_r19          *struct{}
		P_CRI_r19            *struct{}
		P_SSB_Index_r19      *struct{}
		P_CRI_RSRP_r19       *struct{}
		P_SSB_Index_RSRP_r19 *struct{}
		Rs_PAI_r19           *struct{}
		Csi_PAI_r19          *struct{}
		Cjtc_Dd_r19          *struct{}
		Cjtc_F_r19           *struct{}
		Cjtc_P_r19           *struct{}
		Cjtc_Dd_F_r19        *struct{}
		Cli_RSSI_r19         *struct{}
		Cli_SRS_RSRP_r19     *struct{}
	}
	Csi_ReportCJTC_r19                     *CSI_ReportCJTC_r19
	Csi_ReportSubConfigToAddModListExt_r19 []CSI_ReportSubConfig_v1900
	Csi_ReportUE_Initiated_r19             *CSI_ReportUE_Initiated_r19
}

func (ie *CSI_ReportConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIReportConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.SemiPersistentOnPUSCH_v1530 != nil
	hasExtGroup1 := ie.SemiPersistentOnPUSCH_v1610 != nil || ie.Aperiodic_v1610 != nil || ie.ReportQuantity_r16 != nil || ie.CodebookConfig_r16 != nil
	hasExtGroup2 := ie.Cqi_BitsPerSubband_r17 != nil || ie.GroupBasedBeamReporting_v1710 != nil || ie.CodebookConfig_r17 != nil || ie.SharedCMR_r17 != nil || ie.Csi_ReportMode_r17 != nil || ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil || ie.ReportQuantity_r17 != nil
	hasExtGroup3 := ie.SemiPersistentOnPUSCH_v1720 != nil || ie.Aperiodic_v1720 != nil
	hasExtGroup4 := ie.CodebookConfig_v1730 != nil
	hasExtGroup5 := ie.GroupBasedBeamReporting_v1800 != nil || ie.ReportQuantity_r18 != nil || ie.CodebookConfig_r18 != nil || ie.Csi_ReportSubConfigToAddModList_r18 != nil || ie.Csi_ReportSubConfigToReleaseList_r18 != nil
	hasExtGroup6 := ie.NrofReportedCLImeasResources_r19 != nil || ie.Pucch_CSI_ResourceListExt_r19 != nil || ie.SymbolType_r19 != nil || ie.NrofReportedRS_v1900 != nil || ie.PredictionConfiguration_r19 != nil || ie.CodebookConfig_r19 != nil || ie.PortMappingMethod_r19 != nil || ie.ValueOfM_r19 != nil || ie.ReportQuantity_r19 != nil || ie.Csi_ReportCJTC_r19 != nil || ie.Csi_ReportSubConfigToAddModListExt_r19 != nil || ie.Csi_ReportUE_Initiated_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Carrier != nil, ie.Csi_IM_ResourcesForInterference != nil, ie.Nzp_CSI_RS_ResourcesForInterference != nil, ie.ReportFreqConfiguration != nil, ie.CodebookConfig != nil, ie.Dummy != nil, ie.Cqi_Table != nil, ie.Non_PMI_PortIndication != nil}); err != nil {
		return err
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Encode(e); err != nil {
			return err
		}
	}

	// 4. carrier: ref
	{
		if ie.Carrier != nil {
			if err := ie.Carrier.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. resourcesForChannelMeasurement: ref
	{
		if err := ie.ResourcesForChannelMeasurement.Encode(e); err != nil {
			return err
		}
	}

	// 6. csi-IM-ResourcesForInterference: ref
	{
		if ie.Csi_IM_ResourcesForInterference != nil {
			if err := ie.Csi_IM_ResourcesForInterference.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. nzp-CSI-RS-ResourcesForInterference: ref
	{
		if ie.Nzp_CSI_RS_ResourcesForInterference != nil {
			if err := ie.Nzp_CSI_RS_ResourcesForInterference.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. reportConfigType: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ReportConfigReportConfigTypeConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportConfigType.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportConfigType.Choice {
		case CSI_ReportConfig_ReportConfigType_Periodic:
			c := (*ie.ReportConfigType.Periodic)
			if err := c.ReportSlotConfig.Encode(e); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportConfigReportConfigTypePeriodicPucchCSIResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList))); err != nil {
					return err
				}
				for i := range c.Pucch_CSI_ResourceList {
					if err := c.Pucch_CSI_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUCCH:
			c := (*ie.ReportConfigType.SemiPersistentOnPUCCH)
			if err := c.ReportSlotConfig.Encode(e); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportConfigReportConfigTypeSemiPersistentOnPUCCHPucchCSIResourceListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.Pucch_CSI_ResourceList))); err != nil {
					return err
				}
				for i := range c.Pucch_CSI_ResourceList {
					if err := c.Pucch_CSI_ResourceList[i].Encode(e); err != nil {
						return err
					}
				}
			}
		case CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH:
			c := (*ie.ReportConfigType.SemiPersistentOnPUSCH)
			if err := e.EncodeEnumerated(c.ReportSlotConfig, cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotConfigConstraints); err != nil {
				return err
			}
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotOffsetListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetList {
					if err := e.EncodeInteger(c.ReportSlotOffsetList[i], per.Constrained(0, 32)); err != nil {
						return err
					}
				}
			}
			if err := c.P0alpha.Encode(e); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportConfigType_Aperiodic:
			c := (*ie.ReportConfigType.Aperiodic)
			{
				seqOf := e.NewSequenceOfEncoder(cSIReportConfigReportConfigTypeAperiodicReportSlotOffsetListConstraints)
				if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList))); err != nil {
					return err
				}
				for i := range c.ReportSlotOffsetList {
					if err := e.EncodeInteger(c.ReportSlotOffsetList[i], per.Constrained(0, 32)); err != nil {
						return err
					}
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportConfigType.Choice), Constraint: "index out of range"}
		}
	}

	// 9. reportQuantity: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ReportConfigReportQuantityConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.ReportQuantity.Choice), false, nil); err != nil {
			return err
		}
		switch ie.ReportQuantity.Choice {
		case CSI_ReportConfig_ReportQuantity_None:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_PMI_CQI:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_I1:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI:
			c := (*ie.ReportQuantity.Cri_RI_I1_CQI)
			cSIReportConfigReportQuantityCriRII1CQISeq := e.NewSequenceEncoder(cSIReportConfigReportQuantityCriRII1CQIConstraints)
			if err := cSIReportConfigReportQuantityCriRII1CQISeq.EncodePreamble([]bool{c.Pdsch_BundleSizeForCSI != nil}); err != nil {
				return err
			}
			if c.Pdsch_BundleSizeForCSI != nil {
				if err := e.EncodeEnumerated((*c.Pdsch_BundleSizeForCSI), cSIReportConfigReportQuantityCriRII1CQIPdschBundleSizeForCSIConstraints); err != nil {
					return err
				}
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_CQI:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RSRP:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Ssb_Index_RSRP:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_LI_PMI_CQI:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.ReportQuantity.Choice), Constraint: "index out of range"}
		}
	}

	// 10. reportFreqConfiguration: sequence
	{
		if ie.ReportFreqConfiguration != nil {
			c := ie.ReportFreqConfiguration
			cSIReportConfigReportFreqConfigurationSeq := e.NewSequenceEncoder(cSIReportConfigReportFreqConfigurationConstraints)
			if err := cSIReportConfigReportFreqConfigurationSeq.EncodePreamble([]bool{c.Cqi_FormatIndicator != nil, c.Pmi_FormatIndicator != nil, c.Csi_ReportingBand != nil}); err != nil {
				return err
			}
			if c.Cqi_FormatIndicator != nil {
				if err := e.EncodeEnumerated((*c.Cqi_FormatIndicator), cSIReportConfigReportFreqConfigurationCqiFormatIndicatorConstraints); err != nil {
					return err
				}
			}
			if c.Pmi_FormatIndicator != nil {
				if err := e.EncodeEnumerated((*c.Pmi_FormatIndicator), cSIReportConfigReportFreqConfigurationPmiFormatIndicatorConstraints); err != nil {
					return err
				}
			}
			if c.Csi_ReportingBand != nil {
				choiceEnc := e.NewChoiceEncoder(cSIReportConfigReportFreqConfigurationCsiReportingBandConstraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Csi_ReportingBand).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Csi_ReportingBand).Choice {
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands3:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands3), per.FixedSize(3)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands4:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands4), per.FixedSize(4)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands5:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands5), per.FixedSize(5)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands6:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands6), per.FixedSize(6)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands7:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands7), per.FixedSize(7)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands8:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands8), per.FixedSize(8)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands9:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands9), per.FixedSize(9)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands10:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands10), per.FixedSize(10)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands11:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands11), per.FixedSize(11)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands12:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands12), per.FixedSize(12)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands13:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands13), per.FixedSize(13)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands14:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands14), per.FixedSize(14)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands15:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands15), per.FixedSize(15)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands16:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands16), per.FixedSize(16)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands17:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands17), per.FixedSize(17)); err != nil {
						return err
					}
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands18:
					if err := e.EncodeBitString((*(*c.Csi_ReportingBand).Subbands18), per.FixedSize(18)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 11. timeRestrictionForChannelMeasurements: enumerated
	{
		if err := e.EncodeEnumerated(ie.TimeRestrictionForChannelMeasurements, cSIReportConfigTimeRestrictionForChannelMeasurementsConstraints); err != nil {
			return err
		}
	}

	// 12. timeRestrictionForInterferenceMeasurements: enumerated
	{
		if err := e.EncodeEnumerated(ie.TimeRestrictionForInterferenceMeasurements, cSIReportConfigTimeRestrictionForInterferenceMeasurementsConstraints); err != nil {
			return err
		}
	}

	// 13. codebookConfig: ref
	{
		if ie.CodebookConfig != nil {
			if err := ie.CodebookConfig.Encode(e); err != nil {
				return err
			}
		}
	}

	// 14. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, cSIReportConfigDummyConstraints); err != nil {
				return err
			}
		}
	}

	// 15. groupBasedBeamReporting: choice
	{
		choiceEnc := e.NewChoiceEncoder(cSI_ReportConfigGroupBasedBeamReportingConstraints)
		if err := choiceEnc.EncodeChoice(int64(ie.GroupBasedBeamReporting.Choice), false, nil); err != nil {
			return err
		}
		switch ie.GroupBasedBeamReporting.Choice {
		case CSI_ReportConfig_GroupBasedBeamReporting_Enabled:
			if err := e.EncodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_GroupBasedBeamReporting_Disabled:
			c := (*ie.GroupBasedBeamReporting.Disabled)
			cSIReportConfigGroupBasedBeamReportingDisabledSeq := e.NewSequenceEncoder(cSIReportConfigGroupBasedBeamReportingDisabledConstraints)
			if err := cSIReportConfigGroupBasedBeamReportingDisabledSeq.EncodePreamble([]bool{c.NrofReportedRS != nil}); err != nil {
				return err
			}
			if c.NrofReportedRS != nil {
				if err := e.EncodeEnumerated((*c.NrofReportedRS), cSIReportConfigGroupBasedBeamReportingDisabledNrofReportedRSConstraints); err != nil {
					return err
				}
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.GroupBasedBeamReporting.Choice), Constraint: "index out of range"}
		}
	}

	// 16. cqi-Table: ref
	{
		if ie.Cqi_Table != nil {
			if err := ie.Cqi_Table.Encode(e); err != nil {
				return err
			}
		}
	}

	// 17. subbandSize: enumerated
	{
		if err := e.EncodeEnumerated(ie.SubbandSize, cSIReportConfigSubbandSizeConstraints); err != nil {
			return err
		}
	}

	// 18. non-PMI-PortIndication: sequence-of
	{
		if ie.Non_PMI_PortIndication != nil {
			seqOf := e.NewSequenceOfEncoder(cSIReportConfigNonPMIPortIndicationConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.Non_PMI_PortIndication))); err != nil {
				return err
			}
			for i := range ie.Non_PMI_PortIndication {
				if err := ie.Non_PMI_PortIndication[i].Encode(e); err != nil {
					return err
				}
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
					{Name: "semiPersistentOnPUSCH-v1530", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SemiPersistentOnPUSCH_v1530 != nil}); err != nil {
				return err
			}

			if ie.SemiPersistentOnPUSCH_v1530 != nil {
				c := ie.SemiPersistentOnPUSCH_v1530
				if err := ex.EncodeEnumerated(c.ReportSlotConfig_v1530, cSIReportConfigExtSemiPersistentOnPUSCHV1530ReportSlotConfigV1530Constraints); err != nil {
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
					{Name: "semiPersistentOnPUSCH-v1610", Optional: true},
					{Name: "aperiodic-v1610", Optional: true},
					{Name: "reportQuantity-r16", Optional: true},
					{Name: "codebookConfig-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SemiPersistentOnPUSCH_v1610 != nil, ie.Aperiodic_v1610 != nil, ie.ReportQuantity_r16 != nil, ie.CodebookConfig_r16 != nil}); err != nil {
				return err
			}

			if ie.SemiPersistentOnPUSCH_v1610 != nil {
				c := ie.SemiPersistentOnPUSCH_v1610
				cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq := ex.NewSequenceEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610Constraints)
				if err := cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq.EncodePreamble([]bool{c.ReportSlotOffsetListDCI_0_2_r16 != nil, c.ReportSlotOffsetListDCI_0_1_r16 != nil}); err != nil {
					return err
				}
				if c.ReportSlotOffsetListDCI_0_2_r16 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI02R16Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r16))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_2_r16 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r16[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_1_r16 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI01R16Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r16))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_1_r16 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r16[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
			}

			if ie.Aperiodic_v1610 != nil {
				c := ie.Aperiodic_v1610
				cSIReportConfigExtAperiodicV1610Seq := ex.NewSequenceEncoder(cSIReportConfigExtAperiodicV1610Constraints)
				if err := cSIReportConfigExtAperiodicV1610Seq.EncodePreamble([]bool{c.ReportSlotOffsetListDCI_0_2_r16 != nil, c.ReportSlotOffsetListDCI_0_1_r16 != nil}); err != nil {
					return err
				}
				if c.ReportSlotOffsetListDCI_0_2_r16 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI02R16Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r16))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_2_r16 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r16[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_1_r16 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI01R16Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r16))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_1_r16 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r16[i], per.Constrained(0, 32)); err != nil {
							return err
						}
					}
				}
			}

			if ie.ReportQuantity_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIReportConfigExtReportQuantityR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ReportQuantity_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ReportQuantity_r16).Choice {
				case CSI_ReportConfig_Ext_ReportQuantity_r16_Cri_SINR_r16:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r16_Ssb_Index_SINR_r16:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				}
			}

			if ie.CodebookConfig_r16 != nil {
				if err := ie.CodebookConfig_r16.Encode(ex); err != nil {
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
					{Name: "cqi-BitsPerSubband-r17", Optional: true},
					{Name: "groupBasedBeamReporting-v1710", Optional: true},
					{Name: "codebookConfig-r17", Optional: true},
					{Name: "sharedCMR-r17", Optional: true},
					{Name: "csi-ReportMode-r17", Optional: true},
					{Name: "numberOfSingleTRP-CSI-Mode1-r17", Optional: true},
					{Name: "reportQuantity-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cqi_BitsPerSubband_r17 != nil, ie.GroupBasedBeamReporting_v1710 != nil, ie.CodebookConfig_r17 != nil, ie.SharedCMR_r17 != nil, ie.Csi_ReportMode_r17 != nil, ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil, ie.ReportQuantity_r17 != nil}); err != nil {
				return err
			}

			if ie.Cqi_BitsPerSubband_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cqi_BitsPerSubband_r17, cSIReportConfigExtCqiBitsPerSubbandR17Constraints); err != nil {
					return err
				}
			}

			if ie.GroupBasedBeamReporting_v1710 != nil {
				c := ie.GroupBasedBeamReporting_v1710
				if err := ex.EncodeEnumerated(c.NrofReportedGroups_r17, cSIReportConfigExtGroupBasedBeamReportingV1710NrofReportedGroupsR17Constraints); err != nil {
					return err
				}
			}

			if ie.CodebookConfig_r17 != nil {
				if err := ie.CodebookConfig_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.SharedCMR_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SharedCMR_r17, cSIReportConfigExtSharedCMRR17Constraints); err != nil {
					return err
				}
			}

			if ie.Csi_ReportMode_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Csi_ReportMode_r17, cSIReportConfigExtCsiReportModeR17Constraints); err != nil {
					return err
				}
			}

			if ie.NumberOfSingleTRP_CSI_Mode1_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.NumberOfSingleTRP_CSI_Mode1_r17, cSIReportConfigExtNumberOfSingleTRPCSIMode1R17Constraints); err != nil {
					return err
				}
			}

			if ie.ReportQuantity_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIReportConfigExtReportQuantityR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ReportQuantity_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ReportQuantity_r17).Choice {
				case CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_RSRP_Index_r17:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_RSRP_Index_r17:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_SINR_Index_r17:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_SINR_Index_r17:
					if err := ex.EncodeNull(); err != nil {
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
					{Name: "semiPersistentOnPUSCH-v1720", Optional: true},
					{Name: "aperiodic-v1720", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SemiPersistentOnPUSCH_v1720 != nil, ie.Aperiodic_v1720 != nil}); err != nil {
				return err
			}

			if ie.SemiPersistentOnPUSCH_v1720 != nil {
				c := ie.SemiPersistentOnPUSCH_v1720
				cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq := ex.NewSequenceEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720Constraints)
				if err := cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq.EncodePreamble([]bool{c.ReportSlotOffsetList_r17 != nil, c.ReportSlotOffsetListDCI_0_2_r17 != nil, c.ReportSlotOffsetListDCI_0_1_r17 != nil}); err != nil {
					return err
				}
				if c.ReportSlotOffsetList_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListR17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetList_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetList_r17[i], per.Constrained(0, 128)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_2_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI02R17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_2_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r17[i], per.Constrained(0, 128)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_1_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI01R17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_1_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r17[i], per.Constrained(0, 128)); err != nil {
							return err
						}
					}
				}
			}

			if ie.Aperiodic_v1720 != nil {
				c := ie.Aperiodic_v1720
				cSIReportConfigExtAperiodicV1720Seq := ex.NewSequenceEncoder(cSIReportConfigExtAperiodicV1720Constraints)
				if err := cSIReportConfigExtAperiodicV1720Seq.EncodePreamble([]bool{c.ReportSlotOffsetList_r17 != nil, c.ReportSlotOffsetListDCI_0_2_r17 != nil, c.ReportSlotOffsetListDCI_0_1_r17 != nil}); err != nil {
					return err
				}
				if c.ReportSlotOffsetList_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListR17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetList_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetList_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetList_r17[i], per.Constrained(0, 128)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_2_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI02R17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_2_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_2_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_2_r17[i], per.Constrained(0, 128)); err != nil {
							return err
						}
					}
				}
				if c.ReportSlotOffsetListDCI_0_1_r17 != nil {
					seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI01R17Constraints)
					if err := seqOf.EncodeLength(int64(len(c.ReportSlotOffsetListDCI_0_1_r17))); err != nil {
						return err
					}
					for i := range c.ReportSlotOffsetListDCI_0_1_r17 {
						if err := ex.EncodeInteger(c.ReportSlotOffsetListDCI_0_1_r17[i], per.Constrained(0, 128)); err != nil {
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
					{Name: "codebookConfig-v1730", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CodebookConfig_v1730 != nil}); err != nil {
				return err
			}

			if ie.CodebookConfig_v1730 != nil {
				if err := ie.CodebookConfig_v1730.Encode(ex); err != nil {
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
					{Name: "groupBasedBeamReporting-v1800", Optional: true},
					{Name: "reportQuantity-r18", Optional: true},
					{Name: "codebookConfig-r18", Optional: true},
					{Name: "csi-ReportSubConfigToAddModList-r18", Optional: true},
					{Name: "csi-ReportSubConfigToReleaseList-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GroupBasedBeamReporting_v1800 != nil, ie.ReportQuantity_r18 != nil, ie.CodebookConfig_r18 != nil, ie.Csi_ReportSubConfigToAddModList_r18 != nil, ie.Csi_ReportSubConfigToReleaseList_r18 != nil}); err != nil {
				return err
			}

			if ie.GroupBasedBeamReporting_v1800 != nil {
				c := ie.GroupBasedBeamReporting_v1800
				if err := ex.EncodeEnumerated(c.ReportingMode_r18, cSIReportConfigExtGroupBasedBeamReportingV1800ReportingModeR18Constraints); err != nil {
					return err
				}
			}

			if ie.ReportQuantity_r18 != nil {
				if err := ie.ReportQuantity_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CodebookConfig_r18 != nil {
				if err := ie.CodebookConfig_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_ReportSubConfigToAddModList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtCsiReportSubConfigToAddModListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Csi_ReportSubConfigToAddModList_r18))); err != nil {
					return err
				}
				for i := range ie.Csi_ReportSubConfigToAddModList_r18 {
					if err := ie.Csi_ReportSubConfigToAddModList_r18[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Csi_ReportSubConfigToReleaseList_r18 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtCsiReportSubConfigToReleaseListR18Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Csi_ReportSubConfigToReleaseList_r18))); err != nil {
					return err
				}
				for i := range ie.Csi_ReportSubConfigToReleaseList_r18 {
					if err := ie.Csi_ReportSubConfigToReleaseList_r18[i].Encode(ex); err != nil {
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
					{Name: "nrofReportedCLImeasResources-r19", Optional: true},
					{Name: "pucch-CSI-ResourceListExt-r19", Optional: true},
					{Name: "symbolType-r19", Optional: true},
					{Name: "nrofReportedRS-v1900", Optional: true},
					{Name: "predictionConfiguration-r19", Optional: true},
					{Name: "codebookConfig-r19", Optional: true},
					{Name: "portMappingMethod-r19", Optional: true},
					{Name: "valueOfM-r19", Optional: true},
					{Name: "reportQuantity-r19", Optional: true},
					{Name: "csi-ReportCJTC-r19", Optional: true},
					{Name: "csi-ReportSubConfigToAddModListExt-r19", Optional: true},
					{Name: "csi-ReportUE-Initiated-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.NrofReportedCLImeasResources_r19 != nil, ie.Pucch_CSI_ResourceListExt_r19 != nil, ie.SymbolType_r19 != nil, ie.NrofReportedRS_v1900 != nil, ie.PredictionConfiguration_r19 != nil, ie.CodebookConfig_r19 != nil, ie.PortMappingMethod_r19 != nil, ie.ValueOfM_r19 != nil, ie.ReportQuantity_r19 != nil, ie.Csi_ReportCJTC_r19 != nil, ie.Csi_ReportSubConfigToAddModListExt_r19 != nil, ie.Csi_ReportUE_Initiated_r19 != nil}); err != nil {
				return err
			}

			if ie.NrofReportedCLImeasResources_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofReportedCLImeasResources_r19, cSIReportConfigExtNrofReportedCLImeasResourcesR19Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_CSI_ResourceListExt_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtPucchCSIResourceListExtR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pucch_CSI_ResourceListExt_r19))); err != nil {
					return err
				}
				for i := range ie.Pucch_CSI_ResourceListExt_r19 {
					if err := ie.Pucch_CSI_ResourceListExt_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.SymbolType_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.SymbolType_r19, cSIReportConfigExtSymbolTypeR19Constraints); err != nil {
					return err
				}
			}

			if ie.NrofReportedRS_v1900 != nil {
				if err := ex.EncodeEnumerated(*ie.NrofReportedRS_v1900, cSIReportConfigExtNrofReportedRSV1900Constraints); err != nil {
					return err
				}
			}

			if ie.PredictionConfiguration_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIReportConfigExtPredictionConfigurationR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.PredictionConfiguration_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.PredictionConfiguration_r19).Choice {
				case CSI_ReportConfig_Ext_PredictionConfiguration_r19_Csi_InferencePrediction_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19:
					c := (*(*ie.PredictionConfiguration_r19).ConfigurationForBM_PredictionAndDataCollection_r19)
					cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq := ex.NewSequenceEncoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Constraints)
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.EncodePreamble([]bool{c.AssociatedIdForChannelPrediction_r19 != nil, c.AssociatedIdForChannelMeasurement_r19 != nil, c.NrofReportedPredictedRS_r19 != nil, c.NrofTimeInstance_r19 != nil, c.TimeGap_r19 != nil}); err != nil {
						return err
					}
					if err := c.ResourcesForChannelPrediction_r19.Encode(ex); err != nil {
						return err
					}
					if c.AssociatedIdForChannelPrediction_r19 != nil {
						if err := c.AssociatedIdForChannelPrediction_r19.Encode(ex); err != nil {
							return err
						}
					}
					if c.AssociatedIdForChannelMeasurement_r19 != nil {
						if err := c.AssociatedIdForChannelMeasurement_r19.Encode(ex); err != nil {
							return err
						}
					}
					if c.NrofReportedPredictedRS_r19 != nil {
						if err := ex.EncodeEnumerated((*c.NrofReportedPredictedRS_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofReportedPredictedRSR19Constraints); err != nil {
							return err
						}
					}
					if c.NrofTimeInstance_r19 != nil {
						if err := ex.EncodeEnumerated((*c.NrofTimeInstance_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofTimeInstanceR19Constraints); err != nil {
							return err
						}
					}
					if c.TimeGap_r19 != nil {
						if err := ex.EncodeEnumerated((*c.TimeGap_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19TimeGapR19Constraints); err != nil {
							return err
						}
					}
				case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19:
					c := (*(*ie.PredictionConfiguration_r19).ConfigurationForBM_Monitoring_r19)
					cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq := ex.NewSequenceEncoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Constraints)
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.EncodePreamble([]bool{c.NrofBestBeamForMonitoring_r19 != nil, c.NrofTransmissionOccasion_r19 != nil, c.TimeInstanceForRS_PAI_r19 != nil, c.MappingToResourcesForChannelPrediction_r19 != nil}); err != nil {
						return err
					}
					if err := c.RefToPredictionConfig_r19.Encode(ex); err != nil {
						return err
					}
					if c.NrofBestBeamForMonitoring_r19 != nil {
						if err := ex.EncodeEnumerated((*c.NrofBestBeamForMonitoring_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofBestBeamForMonitoringR19Constraints); err != nil {
							return err
						}
					}
					if c.NrofTransmissionOccasion_r19 != nil {
						if err := ex.EncodeEnumerated((*c.NrofTransmissionOccasion_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofTransmissionOccasionR19Constraints); err != nil {
							return err
						}
					}
					if c.TimeInstanceForRS_PAI_r19 != nil {
						if err := ex.EncodeEnumerated((*c.TimeInstanceForRS_PAI_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19TimeInstanceForRSPAIR19Constraints); err != nil {
							return err
						}
					}
					if c.MappingToResourcesForChannelPrediction_r19 != nil {
						if err := ex.EncodeBitString((*c.MappingToResourcesForChannelPrediction_r19), per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerSet)); err != nil {
							return err
						}
					}
				case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19:
					c := (*(*ie.PredictionConfiguration_r19).ConfigurationForCSI_Monitoring_r19)
					cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq := ex.NewSequenceEncoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Constraints)
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq.EncodeExtensionBit(false); err != nil {
						return err
					}
					if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq.EncodePreamble([]bool{c.TimeInstanceForCSI_PAI_r19 != nil}); err != nil {
						return err
					}
					if err := c.RefToPredictionConfig_r19.Encode(ex); err != nil {
						return err
					}
					if c.TimeInstanceForCSI_PAI_r19 != nil {
						if err := ex.EncodeEnumerated((*c.TimeInstanceForCSI_PAI_r19), cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19TimeInstanceForCSIPAIR19Constraints); err != nil {
							return err
						}
					}
				}
			}

			if ie.CodebookConfig_r19 != nil {
				if err := ie.CodebookConfig_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.PortMappingMethod_r19 != nil {
				if err := ex.EncodeEnumerated(*ie.PortMappingMethod_r19, cSIReportConfigExtPortMappingMethodR19Constraints); err != nil {
					return err
				}
			}

			if ie.ValueOfM_r19 != nil {
				if err := ex.EncodeInteger(*ie.ValueOfM_r19, cSIReportConfigValueOfMR19Constraints); err != nil {
					return err
				}
			}

			if ie.ReportQuantity_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(cSIReportConfigExtReportQuantityR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.ReportQuantity_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.ReportQuantity_r19).Choice {
				case CSI_ReportConfig_Ext_ReportQuantity_r19_NoneBM_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_NoneCSI_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_RSRP_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_RSRP_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Rs_PAI_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Csi_PAI_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_F_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_P_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_F_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_RSSI_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_SRS_RSRP_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				}
			}

			if ie.Csi_ReportCJTC_r19 != nil {
				if err := ie.Csi_ReportCJTC_r19.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Csi_ReportSubConfigToAddModListExt_r19 != nil {
				seqOf := ex.NewSequenceOfEncoder(cSIReportConfigExtCsiReportSubConfigToAddModListExtR19Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Csi_ReportSubConfigToAddModListExt_r19))); err != nil {
					return err
				}
				for i := range ie.Csi_ReportSubConfigToAddModListExt_r19 {
					if err := ie.Csi_ReportSubConfigToAddModListExt_r19[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Csi_ReportUE_Initiated_r19 != nil {
				if err := ie.Csi_ReportUE_Initiated_r19.Encode(ex); err != nil {
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

func (ie *CSI_ReportConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIReportConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. reportConfigId: ref
	{
		if err := ie.ReportConfigId.Decode(d); err != nil {
			return err
		}
	}

	// 4. carrier: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Carrier = new(ServCellIndex)
			if err := ie.Carrier.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. resourcesForChannelMeasurement: ref
	{
		if err := ie.ResourcesForChannelMeasurement.Decode(d); err != nil {
			return err
		}
	}

	// 6. csi-IM-ResourcesForInterference: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Csi_IM_ResourcesForInterference = new(CSI_ResourceConfigId)
			if err := ie.Csi_IM_ResourcesForInterference.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. nzp-CSI-RS-ResourcesForInterference: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Nzp_CSI_RS_ResourcesForInterference = new(CSI_ResourceConfigId)
			if err := ie.Nzp_CSI_RS_ResourcesForInterference.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. reportConfigType: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ReportConfigReportConfigTypeConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportConfigType.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ReportConfig_ReportConfigType_Periodic:
			ie.ReportConfigType.Periodic = &struct {
				ReportSlotConfig       CSI_ReportPeriodicityAndOffset
				Pucch_CSI_ResourceList []PUCCH_CSI_Resource
			}{}
			c := (*ie.ReportConfigType.Periodic)
			{
				if err := c.ReportSlotConfig.Decode(d); err != nil {
					return err
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportConfigReportConfigTypePeriodicPucchCSIResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Pucch_CSI_ResourceList = make([]PUCCH_CSI_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Pucch_CSI_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUCCH:
			ie.ReportConfigType.SemiPersistentOnPUCCH = &struct {
				ReportSlotConfig       CSI_ReportPeriodicityAndOffset
				Pucch_CSI_ResourceList []PUCCH_CSI_Resource
			}{}
			c := (*ie.ReportConfigType.SemiPersistentOnPUCCH)
			{
				if err := c.ReportSlotConfig.Decode(d); err != nil {
					return err
				}
			}
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportConfigReportConfigTypeSemiPersistentOnPUCCHPucchCSIResourceListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.Pucch_CSI_ResourceList = make([]PUCCH_CSI_Resource, n)
				for i := int64(0); i < n; i++ {
					if err := c.Pucch_CSI_ResourceList[i].Decode(d); err != nil {
						return err
					}
				}
			}
		case CSI_ReportConfig_ReportConfigType_SemiPersistentOnPUSCH:
			ie.ReportConfigType.SemiPersistentOnPUSCH = &struct {
				ReportSlotConfig     int64
				ReportSlotOffsetList []int64
				P0alpha              P0_PUSCH_AlphaSetId
			}{}
			c := (*ie.ReportConfigType.SemiPersistentOnPUSCH)
			{
				v, err := d.DecodeEnumerated(cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotConfigConstraints)
				if err != nil {
					return err
				}
				c.ReportSlotConfig = v
			}
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportConfigReportConfigTypeSemiPersistentOnPUSCHReportSlotOffsetListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList[i] = v
				}
			}
			{
				if err := c.P0alpha.Decode(d); err != nil {
					return err
				}
			}
		case CSI_ReportConfig_ReportConfigType_Aperiodic:
			ie.ReportConfigType.Aperiodic = &struct{ ReportSlotOffsetList []int64 }{}
			c := (*ie.ReportConfigType.Aperiodic)
			{
				seqOf := d.NewSequenceOfDecoder(cSIReportConfigReportConfigTypeAperiodicReportSlotOffsetListConstraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := d.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList[i] = v
				}
			}
		}
	}

	// 9. reportQuantity: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ReportConfigReportQuantityConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.ReportQuantity.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ReportConfig_ReportQuantity_None:
			ie.ReportQuantity.None = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_PMI_CQI:
			ie.ReportQuantity.Cri_RI_PMI_CQI = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_I1:
			ie.ReportQuantity.Cri_RI_I1 = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_I1_CQI:
			ie.ReportQuantity.Cri_RI_I1_CQI = &struct{ Pdsch_BundleSizeForCSI *int64 }{}
			c := (*ie.ReportQuantity.Cri_RI_I1_CQI)
			cSIReportConfigReportQuantityCriRII1CQISeq := d.NewSequenceDecoder(cSIReportConfigReportQuantityCriRII1CQIConstraints)
			if err := cSIReportConfigReportQuantityCriRII1CQISeq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigReportQuantityCriRII1CQISeq.IsComponentPresent(0) {
				c.Pdsch_BundleSizeForCSI = new(int64)
				v, err := d.DecodeEnumerated(cSIReportConfigReportQuantityCriRII1CQIPdschBundleSizeForCSIConstraints)
				if err != nil {
					return err
				}
				(*c.Pdsch_BundleSizeForCSI) = v
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_CQI:
			ie.ReportQuantity.Cri_RI_CQI = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RSRP:
			ie.ReportQuantity.Cri_RSRP = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Ssb_Index_RSRP:
			ie.ReportQuantity.Ssb_Index_RSRP = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_ReportQuantity_Cri_RI_LI_PMI_CQI:
			ie.ReportQuantity.Cri_RI_LI_PMI_CQI = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		}
	}

	// 10. reportFreqConfiguration: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.ReportFreqConfiguration = &struct {
				Cqi_FormatIndicator *int64
				Pmi_FormatIndicator *int64
				Csi_ReportingBand   *struct {
					Choice     int
					Subbands3  *per.BitString
					Subbands4  *per.BitString
					Subbands5  *per.BitString
					Subbands6  *per.BitString
					Subbands7  *per.BitString
					Subbands8  *per.BitString
					Subbands9  *per.BitString
					Subbands10 *per.BitString
					Subbands11 *per.BitString
					Subbands12 *per.BitString
					Subbands13 *per.BitString
					Subbands14 *per.BitString
					Subbands15 *per.BitString
					Subbands16 *per.BitString
					Subbands17 *per.BitString
					Subbands18 *per.BitString
				}
			}{}
			c := ie.ReportFreqConfiguration
			cSIReportConfigReportFreqConfigurationSeq := d.NewSequenceDecoder(cSIReportConfigReportFreqConfigurationConstraints)
			if err := cSIReportConfigReportFreqConfigurationSeq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigReportFreqConfigurationSeq.IsComponentPresent(0) {
				c.Cqi_FormatIndicator = new(int64)
				v, err := d.DecodeEnumerated(cSIReportConfigReportFreqConfigurationCqiFormatIndicatorConstraints)
				if err != nil {
					return err
				}
				(*c.Cqi_FormatIndicator) = v
			}
			if cSIReportConfigReportFreqConfigurationSeq.IsComponentPresent(1) {
				c.Pmi_FormatIndicator = new(int64)
				v, err := d.DecodeEnumerated(cSIReportConfigReportFreqConfigurationPmiFormatIndicatorConstraints)
				if err != nil {
					return err
				}
				(*c.Pmi_FormatIndicator) = v
			}
			if cSIReportConfigReportFreqConfigurationSeq.IsComponentPresent(2) {
				c.Csi_ReportingBand = &struct {
					Choice     int
					Subbands3  *per.BitString
					Subbands4  *per.BitString
					Subbands5  *per.BitString
					Subbands6  *per.BitString
					Subbands7  *per.BitString
					Subbands8  *per.BitString
					Subbands9  *per.BitString
					Subbands10 *per.BitString
					Subbands11 *per.BitString
					Subbands12 *per.BitString
					Subbands13 *per.BitString
					Subbands14 *per.BitString
					Subbands15 *per.BitString
					Subbands16 *per.BitString
					Subbands17 *per.BitString
					Subbands18 *per.BitString
				}{}
				choiceDec := d.NewChoiceDecoder(cSIReportConfigReportFreqConfigurationCsiReportingBandConstraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Csi_ReportingBand).Choice = int(index)
				switch index {
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands3:
					(*c.Csi_ReportingBand).Subbands3 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(3))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands3) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands4:
					(*c.Csi_ReportingBand).Subbands4 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(4))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands4) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands5:
					(*c.Csi_ReportingBand).Subbands5 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(5))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands5) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands6:
					(*c.Csi_ReportingBand).Subbands6 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(6))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands6) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands7:
					(*c.Csi_ReportingBand).Subbands7 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(7))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands7) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands8:
					(*c.Csi_ReportingBand).Subbands8 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(8))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands8) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands9:
					(*c.Csi_ReportingBand).Subbands9 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(9))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands9) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands10:
					(*c.Csi_ReportingBand).Subbands10 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(10))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands10) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands11:
					(*c.Csi_ReportingBand).Subbands11 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(11))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands11) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands12:
					(*c.Csi_ReportingBand).Subbands12 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(12))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands12) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands13:
					(*c.Csi_ReportingBand).Subbands13 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(13))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands13) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands14:
					(*c.Csi_ReportingBand).Subbands14 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(14))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands14) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands15:
					(*c.Csi_ReportingBand).Subbands15 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(15))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands15) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands16:
					(*c.Csi_ReportingBand).Subbands16 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(16))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands16) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands17:
					(*c.Csi_ReportingBand).Subbands17 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(17))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands17) = v
				case CSI_ReportConfig_ReportFreqConfiguration_Csi_ReportingBand_Subbands18:
					(*c.Csi_ReportingBand).Subbands18 = new(per.BitString)
					v, err := d.DecodeBitString(per.FixedSize(18))
					if err != nil {
						return err
					}
					(*(*c.Csi_ReportingBand).Subbands18) = v
				}
			}
		}
	}

	// 11. timeRestrictionForChannelMeasurements: enumerated
	{
		v8, err := d.DecodeEnumerated(cSIReportConfigTimeRestrictionForChannelMeasurementsConstraints)
		if err != nil {
			return err
		}
		ie.TimeRestrictionForChannelMeasurements = v8
	}

	// 12. timeRestrictionForInterferenceMeasurements: enumerated
	{
		v9, err := d.DecodeEnumerated(cSIReportConfigTimeRestrictionForInterferenceMeasurementsConstraints)
		if err != nil {
			return err
		}
		ie.TimeRestrictionForInterferenceMeasurements = v9
	}

	// 13. codebookConfig: ref
	{
		if seq.IsComponentPresent(10) {
			ie.CodebookConfig = new(CodebookConfig)
			if err := ie.CodebookConfig.Decode(d); err != nil {
				return err
			}
		}
	}

	// 14. dummy: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(cSIReportConfigDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 15. groupBasedBeamReporting: choice
	{
		choiceDec := d.NewChoiceDecoder(cSI_ReportConfigGroupBasedBeamReportingConstraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.GroupBasedBeamReporting.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case CSI_ReportConfig_GroupBasedBeamReporting_Enabled:
			ie.GroupBasedBeamReporting.Enabled = &struct{}{}
			if err := d.DecodeNull(); err != nil {
				return err
			}
		case CSI_ReportConfig_GroupBasedBeamReporting_Disabled:
			ie.GroupBasedBeamReporting.Disabled = &struct{ NrofReportedRS *int64 }{}
			c := (*ie.GroupBasedBeamReporting.Disabled)
			cSIReportConfigGroupBasedBeamReportingDisabledSeq := d.NewSequenceDecoder(cSIReportConfigGroupBasedBeamReportingDisabledConstraints)
			if err := cSIReportConfigGroupBasedBeamReportingDisabledSeq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigGroupBasedBeamReportingDisabledSeq.IsComponentPresent(0) {
				c.NrofReportedRS = new(int64)
				v, err := d.DecodeEnumerated(cSIReportConfigGroupBasedBeamReportingDisabledNrofReportedRSConstraints)
				if err != nil {
					return err
				}
				(*c.NrofReportedRS) = v
			}
		}
	}

	// 16. cqi-Table: ref
	{
		if seq.IsComponentPresent(13) {
			ie.Cqi_Table = new(CQI_Table)
			if err := ie.Cqi_Table.Decode(d); err != nil {
				return err
			}
		}
	}

	// 17. subbandSize: enumerated
	{
		v14, err := d.DecodeEnumerated(cSIReportConfigSubbandSizeConstraints)
		if err != nil {
			return err
		}
		ie.SubbandSize = v14
	}

	// 18. non-PMI-PortIndication: sequence-of
	{
		if seq.IsComponentPresent(15) {
			seqOf := d.NewSequenceOfDecoder(cSIReportConfigNonPMIPortIndicationConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Non_PMI_PortIndication = make([]PortIndexFor8Ranks, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Non_PMI_PortIndication[i].Decode(d); err != nil {
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
				{Name: "semiPersistentOnPUSCH-v1530", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SemiPersistentOnPUSCH_v1530 = &struct{ ReportSlotConfig_v1530 int64 }{}
			c := ie.SemiPersistentOnPUSCH_v1530
			{
				v, err := dx.DecodeEnumerated(cSIReportConfigExtSemiPersistentOnPUSCHV1530ReportSlotConfigV1530Constraints)
				if err != nil {
					return err
				}
				c.ReportSlotConfig_v1530 = v
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
				{Name: "semiPersistentOnPUSCH-v1610", Optional: true},
				{Name: "aperiodic-v1610", Optional: true},
				{Name: "reportQuantity-r16", Optional: true},
				{Name: "codebookConfig-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SemiPersistentOnPUSCH_v1610 = &struct {
				ReportSlotOffsetListDCI_0_2_r16 []int64
				ReportSlotOffsetListDCI_0_1_r16 []int64
			}{}
			c := ie.SemiPersistentOnPUSCH_v1610
			cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq := dx.NewSequenceDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610Constraints)
			if err := cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq.IsComponentPresent(0) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI02R16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r16[i] = v
				}
			}
			if cSIReportConfigExtSemiPersistentOnPUSCHV1610Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1610ReportSlotOffsetListDCI01R16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r16[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Aperiodic_v1610 = &struct {
				ReportSlotOffsetListDCI_0_2_r16 []int64
				ReportSlotOffsetListDCI_0_1_r16 []int64
			}{}
			c := ie.Aperiodic_v1610
			cSIReportConfigExtAperiodicV1610Seq := dx.NewSequenceDecoder(cSIReportConfigExtAperiodicV1610Constraints)
			if err := cSIReportConfigExtAperiodicV1610Seq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigExtAperiodicV1610Seq.IsComponentPresent(0) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI02R16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r16[i] = v
				}
			}
			if cSIReportConfigExtAperiodicV1610Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtAperiodicV1610ReportSlotOffsetListDCI01R16Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r16 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 32))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r16[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.ReportQuantity_r16 = &struct {
				Choice             int
				Cri_SINR_r16       *struct{}
				Ssb_Index_SINR_r16 *struct{}
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIReportConfigExtReportQuantityR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportQuantity_r16).Choice = int(index)
			switch index {
			case CSI_ReportConfig_Ext_ReportQuantity_r16_Cri_SINR_r16:
				(*ie.ReportQuantity_r16).Cri_SINR_r16 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r16_Ssb_Index_SINR_r16:
				(*ie.ReportQuantity_r16).Ssb_Index_SINR_r16 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.CodebookConfig_r16 = new(CodebookConfig_r16)
			if err := ie.CodebookConfig_r16.Decode(dx); err != nil {
				return err
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
				{Name: "cqi-BitsPerSubband-r17", Optional: true},
				{Name: "groupBasedBeamReporting-v1710", Optional: true},
				{Name: "codebookConfig-r17", Optional: true},
				{Name: "sharedCMR-r17", Optional: true},
				{Name: "csi-ReportMode-r17", Optional: true},
				{Name: "numberOfSingleTRP-CSI-Mode1-r17", Optional: true},
				{Name: "reportQuantity-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtCqiBitsPerSubbandR17Constraints)
			if err != nil {
				return err
			}
			ie.Cqi_BitsPerSubband_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.GroupBasedBeamReporting_v1710 = &struct{ NrofReportedGroups_r17 int64 }{}
			c := ie.GroupBasedBeamReporting_v1710
			{
				v, err := dx.DecodeEnumerated(cSIReportConfigExtGroupBasedBeamReportingV1710NrofReportedGroupsR17Constraints)
				if err != nil {
					return err
				}
				c.NrofReportedGroups_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CodebookConfig_r17 = new(CodebookConfig_r17)
			if err := ie.CodebookConfig_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtSharedCMRR17Constraints)
			if err != nil {
				return err
			}
			ie.SharedCMR_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtCsiReportModeR17Constraints)
			if err != nil {
				return err
			}
			ie.Csi_ReportMode_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtNumberOfSingleTRPCSIMode1R17Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfSingleTRP_CSI_Mode1_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			ie.ReportQuantity_r17 = &struct {
				Choice                   int
				Cri_RSRP_Index_r17       *struct{}
				Ssb_Index_RSRP_Index_r17 *struct{}
				Cri_SINR_Index_r17       *struct{}
				Ssb_Index_SINR_Index_r17 *struct{}
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIReportConfigExtReportQuantityR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportQuantity_r17).Choice = int(index)
			switch index {
			case CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_RSRP_Index_r17:
				(*ie.ReportQuantity_r17).Cri_RSRP_Index_r17 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_RSRP_Index_r17:
				(*ie.ReportQuantity_r17).Ssb_Index_RSRP_Index_r17 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r17_Cri_SINR_Index_r17:
				(*ie.ReportQuantity_r17).Cri_SINR_Index_r17 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r17_Ssb_Index_SINR_Index_r17:
				(*ie.ReportQuantity_r17).Ssb_Index_SINR_Index_r17 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
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
				{Name: "semiPersistentOnPUSCH-v1720", Optional: true},
				{Name: "aperiodic-v1720", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SemiPersistentOnPUSCH_v1720 = &struct {
				ReportSlotOffsetList_r17        []int64
				ReportSlotOffsetListDCI_0_2_r17 []int64
				ReportSlotOffsetListDCI_0_1_r17 []int64
			}{}
			c := ie.SemiPersistentOnPUSCH_v1720
			cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq := dx.NewSequenceDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720Constraints)
			if err := cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq.IsComponentPresent(0) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListR17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r17[i] = v
				}
			}
			if cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI02R17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r17[i] = v
				}
			}
			if cSIReportConfigExtSemiPersistentOnPUSCHV1720Seq.IsComponentPresent(2) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtSemiPersistentOnPUSCHV1720ReportSlotOffsetListDCI01R17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r17[i] = v
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Aperiodic_v1720 = &struct {
				ReportSlotOffsetList_r17        []int64
				ReportSlotOffsetListDCI_0_2_r17 []int64
				ReportSlotOffsetListDCI_0_1_r17 []int64
			}{}
			c := ie.Aperiodic_v1720
			cSIReportConfigExtAperiodicV1720Seq := dx.NewSequenceDecoder(cSIReportConfigExtAperiodicV1720Constraints)
			if err := cSIReportConfigExtAperiodicV1720Seq.DecodePreamble(); err != nil {
				return err
			}
			if cSIReportConfigExtAperiodicV1720Seq.IsComponentPresent(0) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListR17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetList_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetList_r17[i] = v
				}
			}
			if cSIReportConfigExtAperiodicV1720Seq.IsComponentPresent(1) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI02R17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_2_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_2_r17[i] = v
				}
			}
			if cSIReportConfigExtAperiodicV1720Seq.IsComponentPresent(2) {
				seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtAperiodicV1720ReportSlotOffsetListDCI01R17Constraints)
				n, err := seqOf.DecodeLength()
				if err != nil {
					return err
				}
				c.ReportSlotOffsetListDCI_0_1_r17 = make([]int64, n)
				for i := int64(0); i < n; i++ {
					v, err := dx.DecodeInteger(per.Constrained(0, 128))
					if err != nil {
						return err
					}
					c.ReportSlotOffsetListDCI_0_1_r17[i] = v
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
				{Name: "codebookConfig-v1730", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.CodebookConfig_v1730 = new(CodebookConfig_v1730)
			if err := ie.CodebookConfig_v1730.Decode(dx); err != nil {
				return err
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
				{Name: "groupBasedBeamReporting-v1800", Optional: true},
				{Name: "reportQuantity-r18", Optional: true},
				{Name: "codebookConfig-r18", Optional: true},
				{Name: "csi-ReportSubConfigToAddModList-r18", Optional: true},
				{Name: "csi-ReportSubConfigToReleaseList-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.GroupBasedBeamReporting_v1800 = &struct{ ReportingMode_r18 int64 }{}
			c := ie.GroupBasedBeamReporting_v1800
			{
				v, err := dx.DecodeEnumerated(cSIReportConfigExtGroupBasedBeamReportingV1800ReportingModeR18Constraints)
				if err != nil {
					return err
				}
				c.ReportingMode_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.ReportQuantity_r18 = new(TDCP_r18)
			if err := ie.ReportQuantity_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.CodebookConfig_r18 = new(CodebookConfig_r18)
			if err := ie.CodebookConfig_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtCsiReportSubConfigToAddModListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ReportSubConfigToAddModList_r18 = make([]CSI_ReportSubConfig_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ReportSubConfigToAddModList_r18[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtCsiReportSubConfigToReleaseListR18Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ReportSubConfigToReleaseList_r18 = make([]CSI_ReportSubConfigId_r18, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ReportSubConfigToReleaseList_r18[i].Decode(dx); err != nil {
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
				{Name: "nrofReportedCLImeasResources-r19", Optional: true},
				{Name: "pucch-CSI-ResourceListExt-r19", Optional: true},
				{Name: "symbolType-r19", Optional: true},
				{Name: "nrofReportedRS-v1900", Optional: true},
				{Name: "predictionConfiguration-r19", Optional: true},
				{Name: "codebookConfig-r19", Optional: true},
				{Name: "portMappingMethod-r19", Optional: true},
				{Name: "valueOfM-r19", Optional: true},
				{Name: "reportQuantity-r19", Optional: true},
				{Name: "csi-ReportCJTC-r19", Optional: true},
				{Name: "csi-ReportSubConfigToAddModListExt-r19", Optional: true},
				{Name: "csi-ReportUE-Initiated-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtNrofReportedCLImeasResourcesR19Constraints)
			if err != nil {
				return err
			}
			ie.NrofReportedCLImeasResources_r19 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtPucchCSIResourceListExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pucch_CSI_ResourceListExt_r19 = make([]PUCCH_CSI_ResourceExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pucch_CSI_ResourceListExt_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtSymbolTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.SymbolType_r19 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtNrofReportedRSV1900Constraints)
			if err != nil {
				return err
			}
			ie.NrofReportedRS_v1900 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.PredictionConfiguration_r19 = &struct {
				Choice                                             int
				Csi_InferencePrediction_r19                        *struct{}
				ConfigurationForBM_PredictionAndDataCollection_r19 *struct {
					ResourcesForChannelPrediction_r19     CSI_ResourceConfigId
					AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
					AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
					NrofReportedPredictedRS_r19           *int64
					NrofTimeInstance_r19                  *int64
					TimeGap_r19                           *int64
				}
				ConfigurationForBM_Monitoring_r19 *struct {
					RefToPredictionConfig_r19                  CSI_ReportConfigId
					NrofBestBeamForMonitoring_r19              *int64
					NrofTransmissionOccasion_r19               *int64
					TimeInstanceForRS_PAI_r19                  *int64
					MappingToResourcesForChannelPrediction_r19 *per.BitString
				}
				ConfigurationForCSI_Monitoring_r19 *struct {
					RefToPredictionConfig_r19  CSI_ReportConfigId
					TimeInstanceForCSI_PAI_r19 *int64
				}
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIReportConfigExtPredictionConfigurationR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PredictionConfiguration_r19).Choice = int(index)
			switch index {
			case CSI_ReportConfig_Ext_PredictionConfiguration_r19_Csi_InferencePrediction_r19:
				(*ie.PredictionConfiguration_r19).Csi_InferencePrediction_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_PredictionAndDataCollection_r19:
				(*ie.PredictionConfiguration_r19).ConfigurationForBM_PredictionAndDataCollection_r19 = &struct {
					ResourcesForChannelPrediction_r19     CSI_ResourceConfigId
					AssociatedIdForChannelPrediction_r19  *AssociatedId_r19
					AssociatedIdForChannelMeasurement_r19 *AssociatedId_r19
					NrofReportedPredictedRS_r19           *int64
					NrofTimeInstance_r19                  *int64
					TimeGap_r19                           *int64
				}{}
				c := (*(*ie.PredictionConfiguration_r19).ConfigurationForBM_PredictionAndDataCollection_r19)
				cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq := dx.NewSequenceDecoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Constraints)
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.ResourcesForChannelPrediction_r19.Decode(dx); err != nil {
						return err
					}
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.IsComponentPresent(1) {
					c.AssociatedIdForChannelPrediction_r19 = new(AssociatedId_r19)
					if err := (*c.AssociatedIdForChannelPrediction_r19).Decode(dx); err != nil {
						return err
					}
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.IsComponentPresent(2) {
					c.AssociatedIdForChannelMeasurement_r19 = new(AssociatedId_r19)
					if err := (*c.AssociatedIdForChannelMeasurement_r19).Decode(dx); err != nil {
						return err
					}
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.IsComponentPresent(3) {
					c.NrofReportedPredictedRS_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofReportedPredictedRSR19Constraints)
					if err != nil {
						return err
					}
					(*c.NrofReportedPredictedRS_r19) = v
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.IsComponentPresent(4) {
					c.NrofTimeInstance_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19NrofTimeInstanceR19Constraints)
					if err != nil {
						return err
					}
					(*c.NrofTimeInstance_r19) = v
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19Seq.IsComponentPresent(5) {
					c.TimeGap_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMPredictionAndDataCollectionR19TimeGapR19Constraints)
					if err != nil {
						return err
					}
					(*c.TimeGap_r19) = v
				}
			case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForBM_Monitoring_r19:
				(*ie.PredictionConfiguration_r19).ConfigurationForBM_Monitoring_r19 = &struct {
					RefToPredictionConfig_r19                  CSI_ReportConfigId
					NrofBestBeamForMonitoring_r19              *int64
					NrofTransmissionOccasion_r19               *int64
					TimeInstanceForRS_PAI_r19                  *int64
					MappingToResourcesForChannelPrediction_r19 *per.BitString
				}{}
				c := (*(*ie.PredictionConfiguration_r19).ConfigurationForBM_Monitoring_r19)
				cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq := dx.NewSequenceDecoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Constraints)
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RefToPredictionConfig_r19.Decode(dx); err != nil {
						return err
					}
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.IsComponentPresent(1) {
					c.NrofBestBeamForMonitoring_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofBestBeamForMonitoringR19Constraints)
					if err != nil {
						return err
					}
					(*c.NrofBestBeamForMonitoring_r19) = v
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.IsComponentPresent(2) {
					c.NrofTransmissionOccasion_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19NrofTransmissionOccasionR19Constraints)
					if err != nil {
						return err
					}
					(*c.NrofTransmissionOccasion_r19) = v
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.IsComponentPresent(3) {
					c.TimeInstanceForRS_PAI_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19TimeInstanceForRSPAIR19Constraints)
					if err != nil {
						return err
					}
					(*c.TimeInstanceForRS_PAI_r19) = v
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForBMMonitoringR19Seq.IsComponentPresent(4) {
					c.MappingToResourcesForChannelPrediction_r19 = new(per.BitString)
					v, err := dx.DecodeBitString(per.SizeRange(1, common.MaxNrofNZP_CSI_RS_ResourcesPerSet))
					if err != nil {
						return err
					}
					(*c.MappingToResourcesForChannelPrediction_r19) = v
				}
			case CSI_ReportConfig_Ext_PredictionConfiguration_r19_ConfigurationForCSI_Monitoring_r19:
				(*ie.PredictionConfiguration_r19).ConfigurationForCSI_Monitoring_r19 = &struct {
					RefToPredictionConfig_r19  CSI_ReportConfigId
					TimeInstanceForCSI_PAI_r19 *int64
				}{}
				c := (*(*ie.PredictionConfiguration_r19).ConfigurationForCSI_Monitoring_r19)
				cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq := dx.NewSequenceDecoder(cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Constraints)
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq.DecodeExtensionBit(); err != nil {
					return err
				}
				if err := cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq.DecodePreamble(); err != nil {
					return err
				}
				{
					if err := c.RefToPredictionConfig_r19.Decode(dx); err != nil {
						return err
					}
				}
				if cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19Seq.IsComponentPresent(1) {
					c.TimeInstanceForCSI_PAI_r19 = new(int64)
					v, err := dx.DecodeEnumerated(cSIReportConfigExtPredictionConfigurationR19ConfigurationForCSIMonitoringR19TimeInstanceForCSIPAIR19Constraints)
					if err != nil {
						return err
					}
					(*c.TimeInstanceForCSI_PAI_r19) = v
				}
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.CodebookConfig_r19 = new(CodebookConfig_r19)
			if err := ie.CodebookConfig_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(cSIReportConfigExtPortMappingMethodR19Constraints)
			if err != nil {
				return err
			}
			ie.PortMappingMethod_r19 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeInteger(cSIReportConfigValueOfMR19Constraints)
			if err != nil {
				return err
			}
			ie.ValueOfM_r19 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.ReportQuantity_r19 = &struct {
				Choice               int
				NoneBM_r19           *struct{}
				NoneCSI_r19          *struct{}
				P_CRI_r19            *struct{}
				P_SSB_Index_r19      *struct{}
				P_CRI_RSRP_r19       *struct{}
				P_SSB_Index_RSRP_r19 *struct{}
				Rs_PAI_r19           *struct{}
				Csi_PAI_r19          *struct{}
				Cjtc_Dd_r19          *struct{}
				Cjtc_F_r19           *struct{}
				Cjtc_P_r19           *struct{}
				Cjtc_Dd_F_r19        *struct{}
				Cli_RSSI_r19         *struct{}
				Cli_SRS_RSRP_r19     *struct{}
			}{}
			choiceDec := dx.NewChoiceDecoder(cSIReportConfigExtReportQuantityR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.ReportQuantity_r19).Choice = int(index)
			switch index {
			case CSI_ReportConfig_Ext_ReportQuantity_r19_NoneBM_r19:
				(*ie.ReportQuantity_r19).NoneBM_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_NoneCSI_r19:
				(*ie.ReportQuantity_r19).NoneCSI_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_r19:
				(*ie.ReportQuantity_r19).P_CRI_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_r19:
				(*ie.ReportQuantity_r19).P_SSB_Index_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_P_CRI_RSRP_r19:
				(*ie.ReportQuantity_r19).P_CRI_RSRP_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_P_SSB_Index_RSRP_r19:
				(*ie.ReportQuantity_r19).P_SSB_Index_RSRP_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Rs_PAI_r19:
				(*ie.ReportQuantity_r19).Rs_PAI_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Csi_PAI_r19:
				(*ie.ReportQuantity_r19).Csi_PAI_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_r19:
				(*ie.ReportQuantity_r19).Cjtc_Dd_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_F_r19:
				(*ie.ReportQuantity_r19).Cjtc_F_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_P_r19:
				(*ie.ReportQuantity_r19).Cjtc_P_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cjtc_Dd_F_r19:
				(*ie.ReportQuantity_r19).Cjtc_Dd_F_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_RSSI_r19:
				(*ie.ReportQuantity_r19).Cli_RSSI_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case CSI_ReportConfig_Ext_ReportQuantity_r19_Cli_SRS_RSRP_r19:
				(*ie.ReportQuantity_r19).Cli_SRS_RSRP_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Csi_ReportCJTC_r19 = new(CSI_ReportCJTC_r19)
			if err := ie.Csi_ReportCJTC_r19.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(10) {
			seqOf := dx.NewSequenceOfDecoder(cSIReportConfigExtCsiReportSubConfigToAddModListExtR19Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Csi_ReportSubConfigToAddModListExt_r19 = make([]CSI_ReportSubConfig_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Csi_ReportSubConfigToAddModListExt_r19[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(11) {
			ie.Csi_ReportUE_Initiated_r19 = new(CSI_ReportUE_Initiated_r19)
			if err := ie.Csi_ReportUE_Initiated_r19.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
