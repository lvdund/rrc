// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PhysicalCellGroupConfig (line 11596).

var physicalCellGroupConfigConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-ACK-SpatialBundlingPUCCH", Optional: true},
		{Name: "harq-ACK-SpatialBundlingPUSCH", Optional: true},
		{Name: "p-NR-FR1", Optional: true},
		{Name: "pdsch-HARQ-ACK-Codebook"},
		{Name: "tpc-SRS-RNTI", Optional: true},
		{Name: "tpc-PUCCH-RNTI", Optional: true},
		{Name: "tpc-PUSCH-RNTI", Optional: true},
		{Name: "sp-CSI-RNTI", Optional: true},
		{Name: "cs-RNTI", Optional: true},
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
	},
}

const (
	PhysicalCellGroupConfig_Harq_ACK_SpatialBundlingPUCCH_True = 0
)

var physicalCellGroupConfigHarqACKSpatialBundlingPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Harq_ACK_SpatialBundlingPUCCH_True},
}

const (
	PhysicalCellGroupConfig_Harq_ACK_SpatialBundlingPUSCH_True = 0
)

var physicalCellGroupConfigHarqACKSpatialBundlingPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Harq_ACK_SpatialBundlingPUSCH_True},
}

const (
	PhysicalCellGroupConfig_Pdsch_HARQ_ACK_Codebook_SemiStatic = 0
	PhysicalCellGroupConfig_Pdsch_HARQ_ACK_Codebook_Dynamic    = 1
)

var physicalCellGroupConfigPdschHARQACKCodebookConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Pdsch_HARQ_ACK_Codebook_SemiStatic, PhysicalCellGroupConfig_Pdsch_HARQ_ACK_Codebook_Dynamic},
}

var physicalCellGroupConfigCsRNTIConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Cs_RNTI_Release = 0
	PhysicalCellGroupConfig_Cs_RNTI_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_XScale_DB0    = 0
	PhysicalCellGroupConfig_Ext_XScale_DB6    = 1
	PhysicalCellGroupConfig_Ext_XScale_Spare2 = 2
	PhysicalCellGroupConfig_Ext_XScale_Spare1 = 3
)

var physicalCellGroupConfigExtXScaleConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_XScale_DB0, PhysicalCellGroupConfig_Ext_XScale_DB6, PhysicalCellGroupConfig_Ext_XScale_Spare2, PhysicalCellGroupConfig_Ext_XScale_Spare1},
}

var physicalCellGroupConfigExtPdcchBlindDetectionConstraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Setup   = 1
)

var physicalCellGroupConfigExtDcpConfigR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Release = 0
	PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16_Enabled  = 0
	PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16_Disabled = 1
)

var physicalCellGroupConfigExtHarqACKSpatialBundlingPUCCHSecondaryPUCCHgroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16_Enabled, PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16_Disabled},
}

const (
	PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16_Enabled  = 0
	PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16_Disabled = 1
)

var physicalCellGroupConfigExtHarqACKSpatialBundlingPUSCHSecondaryPUCCHgroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16_Enabled, PhysicalCellGroupConfig_Ext_Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16_Disabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16_SemiStatic = 0
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16_Dynamic    = 1
)

var physicalCellGroupConfigExtPdschHARQACKCodebookSecondaryPUCCHgroupR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16_SemiStatic, PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16_Dynamic},
}

const (
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Semi_Static_Mode1 = 0
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Semi_Static_Mode2 = 1
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Dynamic           = 2
)

var physicalCellGroupConfigExtNrdcPCmodeFR1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Semi_Static_Mode1, PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Semi_Static_Mode2, PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR1_r16_Dynamic},
}

const (
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Semi_Static_Mode1 = 0
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Semi_Static_Mode2 = 1
	PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Dynamic           = 2
)

var physicalCellGroupConfigExtNrdcPCmodeFR2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Semi_Static_Mode1, PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Semi_Static_Mode2, PhysicalCellGroupConfig_Ext_Nrdc_PCmode_FR2_r16_Dynamic},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_r16_EnhancedDynamic = 0
)

var physicalCellGroupConfigExtPdschHARQACKCodebookR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Codebook_r16_EnhancedDynamic},
}

const (
	PhysicalCellGroupConfig_Ext_Nfi_TotalDAI_Included_r16_True = 0
)

var physicalCellGroupConfigExtNfiTotalDAIIncludedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Nfi_TotalDAI_Included_r16_True},
}

const (
	PhysicalCellGroupConfig_Ext_Ul_TotalDAI_Included_r16_True = 0
)

var physicalCellGroupConfigExtUlTotalDAIIncludedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Ul_TotalDAI_Included_r16_True},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedback_r16_True = 0
)

var physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedback_r16_True},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16_True = 0
)

var physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackNDIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16_True},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16_True = 0
)

var physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackCBGR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16_True},
}

const (
	PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_0_2_r16_Enabled = 0
)

var physicalCellGroupConfigExtDownlinkAssignmentIndexDCI02R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_0_2_r16_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N1 = 0
	PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N2 = 1
	PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N4 = 2
)

var physicalCellGroupConfigExtDownlinkAssignmentIndexDCI12R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N1, PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N2, PhysicalCellGroupConfig_Ext_DownlinkAssignmentIndexDCI_1_2_r16_N4},
}

var physicalCellGroupConfigExtPdschHARQACKCodebookListR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Release = 0
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_AckNackFeedbackMode_r16_Joint    = 0
	PhysicalCellGroupConfig_Ext_AckNackFeedbackMode_r16_Separate = 1
)

var physicalCellGroupConfigExtAckNackFeedbackModeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_AckNackFeedbackMode_r16_Joint, PhysicalCellGroupConfig_Ext_AckNackFeedbackMode_r16_Separate},
}

var physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Setup   = 1
)

var physicalCellGroupConfigExtPdcchBlindDetection2R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Setup   = 1
)

var physicalCellGroupConfigExtPdcchBlindDetection3R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_BdFactorR_r16_N1 = 0
)

var physicalCellGroupConfigExtBdFactorRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_BdFactorR_r16_N1},
}

var physicalCellGroupConfigExtPdschHARQACKEnhType3ToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofEnhType3HARQ_ACK_r17)

var physicalCellGroupConfigExtPdschHARQACKEnhType3ToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofEnhType3HARQ_ACK_r17)

var physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToAddModListR17Constraints = per.SizeRange(1, common.MaxNrofEnhType3HARQ_ACK_r17)

var physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToReleaseListR17Constraints = per.SizeRange(1, common.MaxNrofEnhType3HARQ_ACK_r17)

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldSecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_EnhType3DCI_Field_r17_Enabled = 0
)

var physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_EnhType3DCI_Field_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Retx_r17_Enabled = 0
)

var physicalCellGroupConfigExtPdschHARQACKRetxR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_Retx_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtPdschHARQACKRetxSecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pucch_SSCellDyn_r17_Enabled = 0
)

var physicalCellGroupConfigExtPucchSSCellDynR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pucch_SSCellDyn_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Pucch_SSCellDynSecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtPucchSSCellDynSecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Pucch_SSCellDynSecondaryPUCCHgroup_r17_Enabled},
}

var physicalCellGroupConfigExtPucchSSCellPatternR17Constraints = per.SizeRange(1, common.MaxNrofSlots)

var physicalCellGroupConfigExtPucchSSCellPatternSecondaryPUCCHgroupR17Constraints = per.SizeRange(1, common.MaxNrofSlots)

const (
	PhysicalCellGroupConfig_Ext_Uci_MuxWithDiffPrio_r17_Enabled = 0
)

var physicalCellGroupConfigExtUciMuxWithDiffPrioR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Uci_MuxWithDiffPrio_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtUciMuxWithDiffPrioSecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_r17_Enabled = 0
)

var physicalCellGroupConfigExtSimultaneousPUCCHPUSCHR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_PrioLowDG_HighCG_r17_Enabled = 0
)

var physicalCellGroupConfigExtPrioLowDGHighCGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_PrioLowDG_HighCG_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_PrioHighDG_LowCG_r17_Enabled = 0
)

var physicalCellGroupConfigExtPrioHighDGLowCGR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_PrioHighDG_LowCG_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_TwoQCLTypeDforPDCCHRepetition_r17_Enabled = 0
)

var physicalCellGroupConfigExtTwoQCLTypeDforPDCCHRepetitionR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_TwoQCLTypeDforPDCCHRepetition_r17_Enabled},
}

var physicalCellGroupConfigExtMulticastConfigR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Release = 0
	PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Setup   = 1
)

var physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17_Enabled = 0
)

var physicalCellGroupConfigExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_IntraBandNC_PRACH_SimulTx_r17_Enabled = 0
)

var physicalCellGroupConfigExtIntraBandNCPRACHSimulTxR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_IntraBandNC_PRACH_SimulTx_r17_Enabled},
}

var physicalCellGroupConfigExtPdcchBlindDetection4R17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Release = 0
	PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SamePriority_r17_Enabled = 0
)

var physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePriorityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SamePriority_r17_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17_Enabled = 0
)

var physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePrioritySecondaryPUCCHgroupR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17_Enabled},
}

var physicalCellGroupConfigExtCellDTRXDCIConfigR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Release = 0
	PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Setup   = 1
)

const (
	PhysicalCellGroupConfig_Ext_TwoQCL_TypeD_ForMultiDCI_r18_Enabled = 0
)

var physicalCellGroupConfigExtTwoQCLTypeDForMultiDCIR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_TwoQCL_TypeD_ForMultiDCI_r18_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled = 0
)

var physicalCellGroupConfigExtEnableType1HARQACKMuxForDLAssignmentAfterULGrantR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled = 0
)

var physicalCellGroupConfigExtEnableType2HARQACKMuxForDLAssignmentAfterULGrantR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled = 0
)

var physicalCellGroupConfigExtEnableType3HARQACKMuxForDLAssignmentAfterULGrantR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_EnableDiffPUCCH_Resource_r18_Enabled = 0
)

var physicalCellGroupConfigExtEnableDiffPUCCHResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_EnableDiffPUCCH_Resource_r18_Enabled},
}

const (
	PhysicalCellGroupConfig_Ext_EnableDiffCB_Size_r18_Enabled = 0
)

var physicalCellGroupConfigExtEnableDiffCBSizeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PhysicalCellGroupConfig_Ext_EnableDiffCB_Size_r18_Enabled},
}

var physicalCellGroupConfigExtLpwusConfigR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Release = 0
	PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Setup   = 1
)

var physicalCellGroupConfigExtAdaptSSBPeriodicityIndicationRNTIR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Release = 0
	PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Setup   = 1
)

var physicalCellGroupConfigExtAdaptSSBSizeDCI29R19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Release = 0
	PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Setup   = 1
)

type PhysicalCellGroupConfig struct {
	Harq_ACK_SpatialBundlingPUCCH *int64
	Harq_ACK_SpatialBundlingPUSCH *int64
	P_NR_FR1                      *P_Max
	Pdsch_HARQ_ACK_Codebook       int64
	Tpc_SRS_RNTI                  *RNTI_Value
	Tpc_PUCCH_RNTI                *RNTI_Value
	Tpc_PUSCH_RNTI                *RNTI_Value
	Sp_CSI_RNTI                   *RNTI_Value
	Cs_RNTI                       *struct {
		Choice  int
		Release *struct{}
		Setup   *RNTI_Value
	}
	Mcs_C_RNTI           *RNTI_Value
	P_UE_FR1             *P_Max
	XScale               *int64
	Pdcch_BlindDetection *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetection
	}
	Dcp_Config_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *DCP_Config_r16
	}
	Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16 *int64
	Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16 *int64
	Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16       *int64
	P_NR_FR2_r16                                          *P_Max
	P_UE_FR2_r16                                          *P_Max
	Nrdc_PCmode_FR1_r16                                   *int64
	Nrdc_PCmode_FR2_r16                                   *int64
	Pdsch_HARQ_ACK_Codebook_r16                           *int64
	Nfi_TotalDAI_Included_r16                             *int64
	Ul_TotalDAI_Included_r16                              *int64
	Pdsch_HARQ_ACK_OneShotFeedback_r16                    *int64
	Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16                 *int64
	Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16                 *int64
	DownlinkAssignmentIndexDCI_0_2_r16                    *int64
	DownlinkAssignmentIndexDCI_1_2_r16                    *int64
	Pdsch_HARQ_ACK_CodebookList_r16                       *struct {
		Choice  int
		Release *struct{}
		Setup   *PDSCH_HARQ_ACK_CodebookList_r16
	}
	AckNackFeedbackMode_r16                  *int64
	Pdcch_BlindDetectionCA_CombIndicator_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetectionCA_CombIndicator_r16
	}
	Pdcch_BlindDetection2_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetection2_r16
	}
	Pdcch_BlindDetection3_r16 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetection3_r16
	}
	BdFactorR_r16                                           *int64
	Pdsch_HARQ_ACK_EnhType3ToAddModList_r17                 []PDSCH_HARQ_ACK_EnhType3_r17
	Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17                []PDSCH_HARQ_ACK_EnhType3Index_r17
	Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17        []PDSCH_HARQ_ACK_EnhType3_r17
	Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17       []PDSCH_HARQ_ACK_EnhType3Index_r17
	Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 *int64
	Pdsch_HARQ_ACK_EnhType3DCI_Field_r17                    *int64
	Pdsch_HARQ_ACK_Retx_r17                                 *int64
	Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17              *int64
	Pucch_SSCell_r17                                        *SCellIndex
	Pucch_SSCellSecondaryPUCCHgroup_r17                     *SCellIndex
	Pucch_SSCellDyn_r17                                     *int64
	Pucch_SSCellDynSecondaryPUCCHgroup_r17                  *int64
	Pucch_SSCellPattern_r17                                 []int64
	Pucch_SSCellPatternSecondaryPUCCHgroup_r17              []int64
	Uci_MuxWithDiffPrio_r17                                 *int64
	Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17              *int64
	SimultaneousPUCCH_PUSCH_r17                             *int64
	SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17         *int64
	PrioLowDG_HighCG_r17                                    *int64
	PrioHighDG_LowCG_r17                                    *int64
	TwoQCLTypeDforPDCCHRepetition_r17                       *int64
	MulticastConfig_r17                                     *struct {
		Choice  int
		Release *struct{}
		Setup   *MulticastConfig_r17
	}
	Pdcch_BlindDetectionCA_CombIndicator_r17 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetectionCA_CombIndicator_r17
	}
	SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 *int64
	IntraBandNC_PRACH_SimulTx_r17             *int64
	Pdcch_BlindDetection4_r17                 *struct {
		Choice  int
		Release *struct{}
		Setup   *PDCCH_BlindDetection4_r17
	}
	SimultaneousPUCCH_PUSCH_SamePriority_r17                     *int64
	SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17 *int64
	Ncr_RNTI_r18                                                 *RNTI_Value
	CellDTRX_DCI_Config_r18                                      *struct {
		Choice  int
		Release *struct{}
		Setup   *CellDTRX_DCI_Config_r18
	}
	TwoQCL_TypeD_ForMultiDCI_r18                             *int64
	EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 *int64
	EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 *int64
	EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 *int64
	EnableDiffPUCCH_Resource_r18                             *int64
	EnableDiffCB_Size_r18                                    *int64
	Lpwus_Config_r19                                         *struct {
		Choice  int
		Release *struct{}
		Setup   *LPWUS_Config_r19
	}
	AdaptSSB_PeriodicityIndication_RNTI_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *RNTI_Value
	}
	AdaptSSB_SizeDCI_2_9_r19 *struct {
		Choice  int
		Release *struct{}
		Setup   *SizeDCI_2_9_r19
	}
}

func (ie *PhysicalCellGroupConfig) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(physicalCellGroupConfigConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Mcs_C_RNTI != nil || ie.P_UE_FR1 != nil
	hasExtGroup1 := ie.XScale != nil
	hasExtGroup2 := ie.Pdcch_BlindDetection != nil
	hasExtGroup3 := ie.Dcp_Config_r16 != nil || ie.Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16 != nil || ie.Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16 != nil || ie.P_NR_FR2_r16 != nil || ie.P_UE_FR2_r16 != nil || ie.Nrdc_PCmode_FR1_r16 != nil || ie.Nrdc_PCmode_FR2_r16 != nil || ie.Pdsch_HARQ_ACK_Codebook_r16 != nil || ie.Nfi_TotalDAI_Included_r16 != nil || ie.Ul_TotalDAI_Included_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil || ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil || ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil || ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil || ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil || ie.AckNackFeedbackMode_r16 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil || ie.Pdcch_BlindDetection2_r16 != nil || ie.Pdcch_BlindDetection3_r16 != nil || ie.BdFactorR_r16 != nil
	hasExtGroup4 := ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil || ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil || ie.Pdsch_HARQ_ACK_Retx_r17 != nil || ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil || ie.Pucch_SSCell_r17 != nil || ie.Pucch_SSCellSecondaryPUCCHgroup_r17 != nil || ie.Pucch_SSCellDyn_r17 != nil || ie.Pucch_SSCellDynSecondaryPUCCHgroup_r17 != nil || ie.Pucch_SSCellPattern_r17 != nil || ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17 != nil || ie.Uci_MuxWithDiffPrio_r17 != nil || ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil || ie.SimultaneousPUCCH_PUSCH_r17 != nil || ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil || ie.PrioLowDG_HighCG_r17 != nil || ie.PrioHighDG_LowCG_r17 != nil || ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil || ie.MulticastConfig_r17 != nil || ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil
	hasExtGroup5 := ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil
	hasExtGroup6 := ie.IntraBandNC_PRACH_SimulTx_r17 != nil
	hasExtGroup7 := ie.Pdcch_BlindDetection4_r17 != nil
	hasExtGroup8 := ie.SimultaneousPUCCH_PUSCH_SamePriority_r17 != nil || ie.SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17 != nil
	hasExtGroup9 := ie.Ncr_RNTI_r18 != nil || ie.CellDTRX_DCI_Config_r18 != nil || ie.TwoQCL_TypeD_ForMultiDCI_r18 != nil || ie.EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil || ie.EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil || ie.EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil || ie.EnableDiffPUCCH_Resource_r18 != nil || ie.EnableDiffCB_Size_r18 != nil
	hasExtGroup10 := ie.Lpwus_Config_r19 != nil || ie.AdaptSSB_PeriodicityIndication_RNTI_r19 != nil || ie.AdaptSSB_SizeDCI_2_9_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6 || hasExtGroup7 || hasExtGroup8 || hasExtGroup9 || hasExtGroup10

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Harq_ACK_SpatialBundlingPUCCH != nil, ie.Harq_ACK_SpatialBundlingPUSCH != nil, ie.P_NR_FR1 != nil, ie.Tpc_SRS_RNTI != nil, ie.Tpc_PUCCH_RNTI != nil, ie.Tpc_PUSCH_RNTI != nil, ie.Sp_CSI_RNTI != nil, ie.Cs_RNTI != nil}); err != nil {
		return err
	}

	// 3. harq-ACK-SpatialBundlingPUCCH: enumerated
	{
		if ie.Harq_ACK_SpatialBundlingPUCCH != nil {
			if err := e.EncodeEnumerated(*ie.Harq_ACK_SpatialBundlingPUCCH, physicalCellGroupConfigHarqACKSpatialBundlingPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 4. harq-ACK-SpatialBundlingPUSCH: enumerated
	{
		if ie.Harq_ACK_SpatialBundlingPUSCH != nil {
			if err := e.EncodeEnumerated(*ie.Harq_ACK_SpatialBundlingPUSCH, physicalCellGroupConfigHarqACKSpatialBundlingPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 5. p-NR-FR1: ref
	{
		if ie.P_NR_FR1 != nil {
			if err := ie.P_NR_FR1.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. pdsch-HARQ-ACK-Codebook: enumerated
	{
		if err := e.EncodeEnumerated(ie.Pdsch_HARQ_ACK_Codebook, physicalCellGroupConfigPdschHARQACKCodebookConstraints); err != nil {
			return err
		}
	}

	// 7. tpc-SRS-RNTI: ref
	{
		if ie.Tpc_SRS_RNTI != nil {
			if err := ie.Tpc_SRS_RNTI.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. tpc-PUCCH-RNTI: ref
	{
		if ie.Tpc_PUCCH_RNTI != nil {
			if err := ie.Tpc_PUCCH_RNTI.Encode(e); err != nil {
				return err
			}
		}
	}

	// 9. tpc-PUSCH-RNTI: ref
	{
		if ie.Tpc_PUSCH_RNTI != nil {
			if err := ie.Tpc_PUSCH_RNTI.Encode(e); err != nil {
				return err
			}
		}
	}

	// 10. sp-CSI-RNTI: ref
	{
		if ie.Sp_CSI_RNTI != nil {
			if err := ie.Sp_CSI_RNTI.Encode(e); err != nil {
				return err
			}
		}
	}

	// 11. cs-RNTI: choice
	{
		if ie.Cs_RNTI != nil {
			choiceEnc := e.NewChoiceEncoder(physicalCellGroupConfigCsRNTIConstraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Cs_RNTI).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Cs_RNTI).Choice {
			case PhysicalCellGroupConfig_Cs_RNTI_Release:
				if err := e.EncodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Cs_RNTI_Setup:
				if err := (*ie.Cs_RNTI).Setup.Encode(e); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Cs_RNTI).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4, hasExtGroup5, hasExtGroup6, hasExtGroup7, hasExtGroup8, hasExtGroup9, hasExtGroup10}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "mcs-C-RNTI", Optional: true},
					{Name: "p-UE-FR1", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Mcs_C_RNTI != nil, ie.P_UE_FR1 != nil}); err != nil {
				return err
			}

			if ie.Mcs_C_RNTI != nil {
				if err := ie.Mcs_C_RNTI.Encode(ex); err != nil {
					return err
				}
			}

			if ie.P_UE_FR1 != nil {
				if err := ie.P_UE_FR1.Encode(ex); err != nil {
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
					{Name: "xScale", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.XScale != nil}); err != nil {
				return err
			}

			if ie.XScale != nil {
				if err := ex.EncodeEnumerated(*ie.XScale, physicalCellGroupConfigExtXScaleConstraints); err != nil {
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
					{Name: "pdcch-BlindDetection", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcch_BlindDetection != nil}); err != nil {
				return err
			}

			if ie.Pdcch_BlindDetection != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetectionConstraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetection).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetection).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Setup:
					if err := (*ie.Pdcch_BlindDetection).Setup.Encode(ex); err != nil {
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
					{Name: "dcp-Config-r16", Optional: true},
					{Name: "harq-ACK-SpatialBundlingPUCCH-secondaryPUCCHgroup-r16", Optional: true},
					{Name: "harq-ACK-SpatialBundlingPUSCH-secondaryPUCCHgroup-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-Codebook-secondaryPUCCHgroup-r16", Optional: true},
					{Name: "p-NR-FR2-r16", Optional: true},
					{Name: "p-UE-FR2-r16", Optional: true},
					{Name: "nrdc-PCmode-FR1-r16", Optional: true},
					{Name: "nrdc-PCmode-FR2-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-Codebook-r16", Optional: true},
					{Name: "nfi-TotalDAI-Included-r16", Optional: true},
					{Name: "ul-TotalDAI-Included-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-OneShotFeedback-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-OneShotFeedbackNDI-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-OneShotFeedbackCBG-r16", Optional: true},
					{Name: "downlinkAssignmentIndexDCI-0-2-r16", Optional: true},
					{Name: "downlinkAssignmentIndexDCI-1-2-r16", Optional: true},
					{Name: "pdsch-HARQ-ACK-CodebookList-r16", Optional: true},
					{Name: "ackNackFeedbackMode-r16", Optional: true},
					{Name: "pdcch-BlindDetectionCA-CombIndicator-r16", Optional: true},
					{Name: "pdcch-BlindDetection2-r16", Optional: true},
					{Name: "pdcch-BlindDetection3-r16", Optional: true},
					{Name: "bdFactorR-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dcp_Config_r16 != nil, ie.Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16 != nil, ie.Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16 != nil, ie.Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16 != nil, ie.P_NR_FR2_r16 != nil, ie.P_UE_FR2_r16 != nil, ie.Nrdc_PCmode_FR1_r16 != nil, ie.Nrdc_PCmode_FR2_r16 != nil, ie.Pdsch_HARQ_ACK_Codebook_r16 != nil, ie.Nfi_TotalDAI_Included_r16 != nil, ie.Ul_TotalDAI_Included_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil, ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil, ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil, ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil, ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil, ie.AckNackFeedbackMode_r16 != nil, ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil, ie.Pdcch_BlindDetection2_r16 != nil, ie.Pdcch_BlindDetection3_r16 != nil, ie.BdFactorR_r16 != nil}); err != nil {
				return err
			}

			if ie.Dcp_Config_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtDcpConfigR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dcp_Config_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dcp_Config_r16).Choice {
				case PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Setup:
					if err := (*ie.Dcp_Config_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16, physicalCellGroupConfigExtHarqACKSpatialBundlingPUCCHSecondaryPUCCHgroupR16Constraints); err != nil {
					return err
				}
			}

			if ie.Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16, physicalCellGroupConfigExtHarqACKSpatialBundlingPUSCHSecondaryPUCCHgroupR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16, physicalCellGroupConfigExtPdschHARQACKCodebookSecondaryPUCCHgroupR16Constraints); err != nil {
					return err
				}
			}

			if ie.P_NR_FR2_r16 != nil {
				if err := ie.P_NR_FR2_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.P_UE_FR2_r16 != nil {
				if err := ie.P_UE_FR2_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Nrdc_PCmode_FR1_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nrdc_PCmode_FR1_r16, physicalCellGroupConfigExtNrdcPCmodeFR1R16Constraints); err != nil {
					return err
				}
			}

			if ie.Nrdc_PCmode_FR2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nrdc_PCmode_FR2_r16, physicalCellGroupConfigExtNrdcPCmodeFR2R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_Codebook_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_Codebook_r16, physicalCellGroupConfigExtPdschHARQACKCodebookR16Constraints); err != nil {
					return err
				}
			}

			if ie.Nfi_TotalDAI_Included_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nfi_TotalDAI_Included_r16, physicalCellGroupConfigExtNfiTotalDAIIncludedR16Constraints); err != nil {
					return err
				}
			}

			if ie.Ul_TotalDAI_Included_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_TotalDAI_Included_r16, physicalCellGroupConfigExtUlTotalDAIIncludedR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_OneShotFeedback_r16, physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16, physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackNDIR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16, physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackCBGR16Constraints); err != nil {
					return err
				}
			}

			if ie.DownlinkAssignmentIndexDCI_0_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DownlinkAssignmentIndexDCI_0_2_r16, physicalCellGroupConfigExtDownlinkAssignmentIndexDCI02R16Constraints); err != nil {
					return err
				}
			}

			if ie.DownlinkAssignmentIndexDCI_1_2_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.DownlinkAssignmentIndexDCI_1_2_r16, physicalCellGroupConfigExtDownlinkAssignmentIndexDCI12R16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_CodebookList_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdschHARQACKCodebookListR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdsch_HARQ_ACK_CodebookList_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdsch_HARQ_ACK_CodebookList_r16).Choice {
				case PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Setup:
					if err := (*ie.Pdsch_HARQ_ACK_CodebookList_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AckNackFeedbackMode_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AckNackFeedbackMode_r16, physicalCellGroupConfigExtAckNackFeedbackModeR16Constraints); err != nil {
					return err
				}
			}

			if ie.Pdcch_BlindDetectionCA_CombIndicator_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Setup:
					if err := (*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdcch_BlindDetection2_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetection2R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetection2_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetection2_r16).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Setup:
					if err := (*ie.Pdcch_BlindDetection2_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdcch_BlindDetection3_r16 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetection3R16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetection3_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetection3_r16).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Setup:
					if err := (*ie.Pdcch_BlindDetection3_r16).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.BdFactorR_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.BdFactorR_r16, physicalCellGroupConfigExtBdFactorRR16Constraints); err != nil {
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
					{Name: "pdsch-HARQ-ACK-EnhType3ToAddModList-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3ToReleaseList-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3SecondaryToAddModList-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3SecondaryToReleaseList-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3DCI-FieldSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-EnhType3DCI-Field-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-Retx-r17", Optional: true},
					{Name: "pdsch-HARQ-ACK-RetxSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "pucch-sSCell-r17", Optional: true},
					{Name: "pucch-sSCellSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "pucch-sSCellDyn-r17", Optional: true},
					{Name: "pucch-sSCellDynSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "pucch-sSCellPattern-r17", Optional: true},
					{Name: "pucch-sSCellPatternSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "uci-MuxWithDiffPrio-r17", Optional: true},
					{Name: "uci-MuxWithDiffPrioSecondaryPUCCHgroup-r17", Optional: true},
					{Name: "simultaneousPUCCH-PUSCH-r17", Optional: true},
					{Name: "simultaneousPUCCH-PUSCH-SecondaryPUCCHgroup-r17", Optional: true},
					{Name: "prioLowDG-HighCG-r17", Optional: true},
					{Name: "prioHighDG-LowCG-r17", Optional: true},
					{Name: "twoQCLTypeDforPDCCHRepetition-r17", Optional: true},
					{Name: "multicastConfig-r17", Optional: true},
					{Name: "pdcch-BlindDetectionCA-CombIndicator-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil, ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil, ie.Pdsch_HARQ_ACK_Retx_r17 != nil, ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil, ie.Pucch_SSCell_r17 != nil, ie.Pucch_SSCellSecondaryPUCCHgroup_r17 != nil, ie.Pucch_SSCellDyn_r17 != nil, ie.Pucch_SSCellDynSecondaryPUCCHgroup_r17 != nil, ie.Pucch_SSCellPattern_r17 != nil, ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17 != nil, ie.Uci_MuxWithDiffPrio_r17 != nil, ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil, ie.SimultaneousPUCCH_PUSCH_r17 != nil, ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil, ie.PrioLowDG_HighCG_r17 != nil, ie.PrioHighDG_LowCG_r17 != nil, ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil, ie.MulticastConfig_r17 != nil, ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPdschHARQACKEnhType3ToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 {
					if err := ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPdschHARQACKEnhType3ToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 {
					if err := ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToAddModListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17))); err != nil {
					return err
				}
				for i := range ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 {
					if err := ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToReleaseListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17))); err != nil {
					return err
				}
				for i := range ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 {
					if err := ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17, physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldSecondaryPUCCHgroupR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17, physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_Retx_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_Retx_r17, physicalCellGroupConfigExtPdschHARQACKRetxR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17, physicalCellGroupConfigExtPdschHARQACKRetxSecondaryPUCCHgroupR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCell_r17 != nil {
				if err := ie.Pucch_SSCell_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCellSecondaryPUCCHgroup_r17 != nil {
				if err := ie.Pucch_SSCellSecondaryPUCCHgroup_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCellDyn_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_SSCellDyn_r17, physicalCellGroupConfigExtPucchSSCellDynR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCellDynSecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_SSCellDynSecondaryPUCCHgroup_r17, physicalCellGroupConfigExtPucchSSCellDynSecondaryPUCCHgroupR17Constraints); err != nil {
					return err
				}
			}

			if ie.Pucch_SSCellPattern_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPucchSSCellPatternR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pucch_SSCellPattern_r17))); err != nil {
					return err
				}
				for i := range ie.Pucch_SSCellPattern_r17 {
					if err := ex.EncodeInteger(ie.Pucch_SSCellPattern_r17[i], per.Constrained(0, 1)); err != nil {
						return err
					}
				}
			}

			if ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(physicalCellGroupConfigExtPucchSSCellPatternSecondaryPUCCHgroupR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17))); err != nil {
					return err
				}
				for i := range ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17 {
					if err := ex.EncodeInteger(ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17[i], per.Constrained(0, 1)); err != nil {
						return err
					}
				}
			}

			if ie.Uci_MuxWithDiffPrio_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Uci_MuxWithDiffPrio_r17, physicalCellGroupConfigExtUciMuxWithDiffPrioR17Constraints); err != nil {
					return err
				}
			}

			if ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17, physicalCellGroupConfigExtUciMuxWithDiffPrioSecondaryPUCCHgroupR17Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousPUCCH_PUSCH_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousPUCCH_PUSCH_r17, physicalCellGroupConfigExtSimultaneousPUCCHPUSCHR17Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17, physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSecondaryPUCCHgroupR17Constraints); err != nil {
					return err
				}
			}

			if ie.PrioLowDG_HighCG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PrioLowDG_HighCG_r17, physicalCellGroupConfigExtPrioLowDGHighCGR17Constraints); err != nil {
					return err
				}
			}

			if ie.PrioHighDG_LowCG_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.PrioHighDG_LowCG_r17, physicalCellGroupConfigExtPrioHighDGLowCGR17Constraints); err != nil {
					return err
				}
			}

			if ie.TwoQCLTypeDforPDCCHRepetition_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoQCLTypeDforPDCCHRepetition_r17, physicalCellGroupConfigExtTwoQCLTypeDforPDCCHRepetitionR17Constraints); err != nil {
					return err
				}
			}

			if ie.MulticastConfig_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtMulticastConfigR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.MulticastConfig_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.MulticastConfig_r17).Choice {
				case PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Setup:
					if err := (*ie.MulticastConfig_r17).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Pdcch_BlindDetectionCA_CombIndicator_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Setup:
					if err := (*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Setup.Encode(ex); err != nil {
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
					{Name: "simultaneousSR-PUSCH-diffPUCCH-Groups-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil}); err != nil {
				return err
			}

			if ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17, physicalCellGroupConfigExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints); err != nil {
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
					{Name: "intraBandNC-PRACH-simulTx-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IntraBandNC_PRACH_SimulTx_r17 != nil}); err != nil {
				return err
			}

			if ie.IntraBandNC_PRACH_SimulTx_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.IntraBandNC_PRACH_SimulTx_r17, physicalCellGroupConfigExtIntraBandNCPRACHSimulTxR17Constraints); err != nil {
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
					{Name: "pdcch-BlindDetection4-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcch_BlindDetection4_r17 != nil}); err != nil {
				return err
			}

			if ie.Pdcch_BlindDetection4_r17 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtPdcchBlindDetection4R17Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Pdcch_BlindDetection4_r17).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Pdcch_BlindDetection4_r17).Choice {
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Setup:
					if err := (*ie.Pdcch_BlindDetection4_r17).Setup.Encode(ex); err != nil {
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
					{Name: "simultaneousPUCCH-PUSCH-SamePriority-r17", Optional: true},
					{Name: "simultaneousPUCCH-PUSCH-SamePriority-SecondaryPUCCHgroup-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SimultaneousPUCCH_PUSCH_SamePriority_r17 != nil, ie.SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17 != nil}); err != nil {
				return err
			}

			if ie.SimultaneousPUCCH_PUSCH_SamePriority_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousPUCCH_PUSCH_SamePriority_r17, physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePriorityR17Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17, physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePrioritySecondaryPUCCHgroupR17Constraints); err != nil {
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
					{Name: "ncr-RNTI-r18", Optional: true},
					{Name: "cellDTRX-DCI-config-r18", Optional: true},
					{Name: "twoQCL-TypeD-ForMultiDCI-r18", Optional: true},
					{Name: "enableType1HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
					{Name: "enableType2HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
					{Name: "enableType3HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
					{Name: "enableDiffPUCCH-Resource-r18", Optional: true},
					{Name: "enableDiffCB-Size-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ncr_RNTI_r18 != nil, ie.CellDTRX_DCI_Config_r18 != nil, ie.TwoQCL_TypeD_ForMultiDCI_r18 != nil, ie.EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil, ie.EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil, ie.EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil, ie.EnableDiffPUCCH_Resource_r18 != nil, ie.EnableDiffCB_Size_r18 != nil}); err != nil {
				return err
			}

			if ie.Ncr_RNTI_r18 != nil {
				if err := ie.Ncr_RNTI_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.CellDTRX_DCI_Config_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtCellDTRXDCIConfigR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.CellDTRX_DCI_Config_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.CellDTRX_DCI_Config_r18).Choice {
				case PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Setup:
					if err := (*ie.CellDTRX_DCI_Config_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.TwoQCL_TypeD_ForMultiDCI_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoQCL_TypeD_ForMultiDCI_r18, physicalCellGroupConfigExtTwoQCLTypeDForMultiDCIR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18, physicalCellGroupConfigExtEnableType1HARQACKMuxForDLAssignmentAfterULGrantR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18, physicalCellGroupConfigExtEnableType2HARQACKMuxForDLAssignmentAfterULGrantR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18, physicalCellGroupConfigExtEnableType3HARQACKMuxForDLAssignmentAfterULGrantR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDiffPUCCH_Resource_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDiffPUCCH_Resource_r18, physicalCellGroupConfigExtEnableDiffPUCCHResourceR18Constraints); err != nil {
					return err
				}
			}

			if ie.EnableDiffCB_Size_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.EnableDiffCB_Size_r18, physicalCellGroupConfigExtEnableDiffCBSizeR18Constraints); err != nil {
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
					{Name: "lpwus-Config-r19", Optional: true},
					{Name: "adaptSSB-PeriodicityIndication-RNTI-r19", Optional: true},
					{Name: "adaptSSB-SizeDCI-2-9-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Lpwus_Config_r19 != nil, ie.AdaptSSB_PeriodicityIndication_RNTI_r19 != nil, ie.AdaptSSB_SizeDCI_2_9_r19 != nil}); err != nil {
				return err
			}

			if ie.Lpwus_Config_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtLpwusConfigR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Lpwus_Config_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Lpwus_Config_r19).Choice {
				case PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Setup:
					if err := (*ie.Lpwus_Config_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AdaptSSB_PeriodicityIndication_RNTI_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtAdaptSSBPeriodicityIndicationRNTIR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Choice {
				case PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Setup:
					if err := (*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.AdaptSSB_SizeDCI_2_9_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(physicalCellGroupConfigExtAdaptSSBSizeDCI29R19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.AdaptSSB_SizeDCI_2_9_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.AdaptSSB_SizeDCI_2_9_r19).Choice {
				case PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Setup:
					if err := (*ie.AdaptSSB_SizeDCI_2_9_r19).Setup.Encode(ex); err != nil {
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

func (ie *PhysicalCellGroupConfig) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(physicalCellGroupConfigConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. harq-ACK-SpatialBundlingPUCCH: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(physicalCellGroupConfigHarqACKSpatialBundlingPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.Harq_ACK_SpatialBundlingPUCCH = &idx
		}
	}

	// 4. harq-ACK-SpatialBundlingPUSCH: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(physicalCellGroupConfigHarqACKSpatialBundlingPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.Harq_ACK_SpatialBundlingPUSCH = &idx
		}
	}

	// 5. p-NR-FR1: ref
	{
		if seq.IsComponentPresent(2) {
			ie.P_NR_FR1 = new(P_Max)
			if err := ie.P_NR_FR1.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. pdsch-HARQ-ACK-Codebook: enumerated
	{
		v3, err := d.DecodeEnumerated(physicalCellGroupConfigPdschHARQACKCodebookConstraints)
		if err != nil {
			return err
		}
		ie.Pdsch_HARQ_ACK_Codebook = v3
	}

	// 7. tpc-SRS-RNTI: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Tpc_SRS_RNTI = new(RNTI_Value)
			if err := ie.Tpc_SRS_RNTI.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. tpc-PUCCH-RNTI: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Tpc_PUCCH_RNTI = new(RNTI_Value)
			if err := ie.Tpc_PUCCH_RNTI.Decode(d); err != nil {
				return err
			}
		}
	}

	// 9. tpc-PUSCH-RNTI: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Tpc_PUSCH_RNTI = new(RNTI_Value)
			if err := ie.Tpc_PUSCH_RNTI.Decode(d); err != nil {
				return err
			}
		}
	}

	// 10. sp-CSI-RNTI: ref
	{
		if seq.IsComponentPresent(7) {
			ie.Sp_CSI_RNTI = new(RNTI_Value)
			if err := ie.Sp_CSI_RNTI.Decode(d); err != nil {
				return err
			}
		}
	}

	// 11. cs-RNTI: choice
	{
		if seq.IsComponentPresent(8) {
			ie.Cs_RNTI = &struct {
				Choice  int
				Release *struct{}
				Setup   *RNTI_Value
			}{}
			choiceDec := d.NewChoiceDecoder(physicalCellGroupConfigCsRNTIConstraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Cs_RNTI).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PhysicalCellGroupConfig_Cs_RNTI_Release:
				(*ie.Cs_RNTI).Release = &struct{}{}
				if err := d.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Cs_RNTI_Setup:
				(*ie.Cs_RNTI).Setup = new(RNTI_Value)
				if err := (*ie.Cs_RNTI).Setup.Decode(d); err != nil {
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
				{Name: "mcs-C-RNTI", Optional: true},
				{Name: "p-UE-FR1", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Mcs_C_RNTI = new(RNTI_Value)
			if err := ie.Mcs_C_RNTI.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.P_UE_FR1 = new(P_Max)
			if err := ie.P_UE_FR1.Decode(dx); err != nil {
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
				{Name: "xScale", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtXScaleConstraints)
			if err != nil {
				return err
			}
			ie.XScale = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdcch-BlindDetection", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdcch_BlindDetection = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetection
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetectionConstraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetection).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Release:
				(*ie.Pdcch_BlindDetection).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection_Setup:
				(*ie.Pdcch_BlindDetection).Setup = new(PDCCH_BlindDetection)
				if err := (*ie.Pdcch_BlindDetection).Setup.Decode(dx); err != nil {
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
				{Name: "dcp-Config-r16", Optional: true},
				{Name: "harq-ACK-SpatialBundlingPUCCH-secondaryPUCCHgroup-r16", Optional: true},
				{Name: "harq-ACK-SpatialBundlingPUSCH-secondaryPUCCHgroup-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-Codebook-secondaryPUCCHgroup-r16", Optional: true},
				{Name: "p-NR-FR2-r16", Optional: true},
				{Name: "p-UE-FR2-r16", Optional: true},
				{Name: "nrdc-PCmode-FR1-r16", Optional: true},
				{Name: "nrdc-PCmode-FR2-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-Codebook-r16", Optional: true},
				{Name: "nfi-TotalDAI-Included-r16", Optional: true},
				{Name: "ul-TotalDAI-Included-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-OneShotFeedback-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-OneShotFeedbackNDI-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-OneShotFeedbackCBG-r16", Optional: true},
				{Name: "downlinkAssignmentIndexDCI-0-2-r16", Optional: true},
				{Name: "downlinkAssignmentIndexDCI-1-2-r16", Optional: true},
				{Name: "pdsch-HARQ-ACK-CodebookList-r16", Optional: true},
				{Name: "ackNackFeedbackMode-r16", Optional: true},
				{Name: "pdcch-BlindDetectionCA-CombIndicator-r16", Optional: true},
				{Name: "pdcch-BlindDetection2-r16", Optional: true},
				{Name: "pdcch-BlindDetection3-r16", Optional: true},
				{Name: "bdFactorR-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Dcp_Config_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DCP_Config_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtDcpConfigR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dcp_Config_r16).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Release:
				(*ie.Dcp_Config_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Dcp_Config_r16_Setup:
				(*ie.Dcp_Config_r16).Setup = new(DCP_Config_r16)
				if err := (*ie.Dcp_Config_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtHarqACKSpatialBundlingPUCCHSecondaryPUCCHgroupR16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ACK_SpatialBundlingPUCCH_SecondaryPUCCHgroup_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtHarqACKSpatialBundlingPUSCHSecondaryPUCCHgroupR16Constraints)
			if err != nil {
				return err
			}
			ie.Harq_ACK_SpatialBundlingPUSCH_SecondaryPUCCHgroup_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKCodebookSecondaryPUCCHgroupR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_Codebook_SecondaryPUCCHgroup_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			ie.P_NR_FR2_r16 = new(P_Max)
			if err := ie.P_NR_FR2_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(5) {
			ie.P_UE_FR2_r16 = new(P_Max)
			if err := ie.P_UE_FR2_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtNrdcPCmodeFR1R16Constraints)
			if err != nil {
				return err
			}
			ie.Nrdc_PCmode_FR1_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtNrdcPCmodeFR2R16Constraints)
			if err != nil {
				return err
			}
			ie.Nrdc_PCmode_FR2_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKCodebookR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_Codebook_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtNfiTotalDAIIncludedR16Constraints)
			if err != nil {
				return err
			}
			ie.Nfi_TotalDAI_Included_r16 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtUlTotalDAIIncludedR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_TotalDAI_Included_r16 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_OneShotFeedback_r16 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackNDIR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKOneShotFeedbackCBGR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtDownlinkAssignmentIndexDCI02R16Constraints)
			if err != nil {
				return err
			}
			ie.DownlinkAssignmentIndexDCI_0_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtDownlinkAssignmentIndexDCI12R16Constraints)
			if err != nil {
				return err
			}
			ie.DownlinkAssignmentIndexDCI_1_2_r16 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			ie.Pdsch_HARQ_ACK_CodebookList_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDSCH_HARQ_ACK_CodebookList_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdschHARQACKCodebookListR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdsch_HARQ_ACK_CodebookList_r16).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Release:
				(*ie.Pdsch_HARQ_ACK_CodebookList_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdsch_HARQ_ACK_CodebookList_r16_Setup:
				(*ie.Pdsch_HARQ_ACK_CodebookList_r16).Setup = new(PDSCH_HARQ_ACK_CodebookList_r16)
				if err := (*ie.Pdsch_HARQ_ACK_CodebookList_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtAckNackFeedbackModeR16Constraints)
			if err != nil {
				return err
			}
			ie.AckNackFeedbackMode_r16 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			ie.Pdcch_BlindDetectionCA_CombIndicator_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetectionCA_CombIndicator_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Release:
				(*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r16_Setup:
				(*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Setup = new(PDCCH_BlindDetectionCA_CombIndicator_r16)
				if err := (*ie.Pdcch_BlindDetectionCA_CombIndicator_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(19) {
			ie.Pdcch_BlindDetection2_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetection2_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetection2R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetection2_r16).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Release:
				(*ie.Pdcch_BlindDetection2_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection2_r16_Setup:
				(*ie.Pdcch_BlindDetection2_r16).Setup = new(PDCCH_BlindDetection2_r16)
				if err := (*ie.Pdcch_BlindDetection2_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(20) {
			ie.Pdcch_BlindDetection3_r16 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetection3_r16
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetection3R16Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetection3_r16).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Release:
				(*ie.Pdcch_BlindDetection3_r16).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection3_r16_Setup:
				(*ie.Pdcch_BlindDetection3_r16).Setup = new(PDCCH_BlindDetection3_r16)
				if err := (*ie.Pdcch_BlindDetection3_r16).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(21) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtBdFactorRR16Constraints)
			if err != nil {
				return err
			}
			ie.BdFactorR_r16 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdsch-HARQ-ACK-EnhType3ToAddModList-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3ToReleaseList-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3SecondaryToAddModList-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3SecondaryToReleaseList-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3DCI-FieldSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-EnhType3DCI-Field-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-Retx-r17", Optional: true},
				{Name: "pdsch-HARQ-ACK-RetxSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "pucch-sSCell-r17", Optional: true},
				{Name: "pucch-sSCellSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "pucch-sSCellDyn-r17", Optional: true},
				{Name: "pucch-sSCellDynSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "pucch-sSCellPattern-r17", Optional: true},
				{Name: "pucch-sSCellPatternSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "uci-MuxWithDiffPrio-r17", Optional: true},
				{Name: "uci-MuxWithDiffPrioSecondaryPUCCHgroup-r17", Optional: true},
				{Name: "simultaneousPUCCH-PUSCH-r17", Optional: true},
				{Name: "simultaneousPUCCH-PUSCH-SecondaryPUCCHgroup-r17", Optional: true},
				{Name: "prioLowDG-HighCG-r17", Optional: true},
				{Name: "prioHighDG-LowCG-r17", Optional: true},
				{Name: "twoQCLTypeDforPDCCHRepetition-r17", Optional: true},
				{Name: "multicastConfig-r17", Optional: true},
				{Name: "pdcch-BlindDetectionCA-CombIndicator-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPdschHARQACKEnhType3ToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17 = make([]PDSCH_HARQ_ACK_EnhType3_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdsch_HARQ_ACK_EnhType3ToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPdschHARQACKEnhType3ToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 = make([]PDSCH_HARQ_ACK_EnhType3Index_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdsch_HARQ_ACK_EnhType3ToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToAddModListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 = make([]PDSCH_HARQ_ACK_EnhType3_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPdschHARQACKEnhType3SecondaryToReleaseListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 = make([]PDSCH_HARQ_ACK_EnhType3Index_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKEnhType3DCIFieldR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_EnhType3DCI_Field_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKRetxR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_Retx_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPdschHARQACKRetxSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.Pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.Pucch_SSCell_r17 = new(SCellIndex)
			if err := ie.Pucch_SSCell_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Pucch_SSCellSecondaryPUCCHgroup_r17 = new(SCellIndex)
			if err := ie.Pucch_SSCellSecondaryPUCCHgroup_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPucchSSCellDynR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_SSCellDyn_r17 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPucchSSCellDynSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_SSCellDynSecondaryPUCCHgroup_r17 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPucchSSCellPatternR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pucch_SSCellPattern_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				ie.Pucch_SSCellPattern_r17[i] = v
			}
		}

		if groupSeq.IsComponentPresent(13) {
			seqOf := dx.NewSequenceOfDecoder(physicalCellGroupConfigExtPucchSSCellPatternSecondaryPUCCHgroupR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17 = make([]int64, n)
			for i := int64(0); i < n; i++ {
				v, err := dx.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				ie.Pucch_SSCellPatternSecondaryPUCCHgroup_r17[i] = v
			}
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtUciMuxWithDiffPrioR17Constraints)
			if err != nil {
				return err
			}
			ie.Uci_MuxWithDiffPrio_r17 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtUciMuxWithDiffPrioSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.Uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 = &v
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtSimultaneousPUCCHPUSCHR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousPUCCH_PUSCH_r17 = &v
		}

		if groupSeq.IsComponentPresent(17) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 = &v
		}

		if groupSeq.IsComponentPresent(18) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPrioLowDGHighCGR17Constraints)
			if err != nil {
				return err
			}
			ie.PrioLowDG_HighCG_r17 = &v
		}

		if groupSeq.IsComponentPresent(19) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtPrioHighDGLowCGR17Constraints)
			if err != nil {
				return err
			}
			ie.PrioHighDG_LowCG_r17 = &v
		}

		if groupSeq.IsComponentPresent(20) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtTwoQCLTypeDforPDCCHRepetitionR17Constraints)
			if err != nil {
				return err
			}
			ie.TwoQCLTypeDforPDCCHRepetition_r17 = &v
		}

		if groupSeq.IsComponentPresent(21) {
			ie.MulticastConfig_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *MulticastConfig_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtMulticastConfigR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.MulticastConfig_r17).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Release:
				(*ie.MulticastConfig_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_MulticastConfig_r17_Setup:
				(*ie.MulticastConfig_r17).Setup = new(MulticastConfig_r17)
				if err := (*ie.MulticastConfig_r17).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(22) {
			ie.Pdcch_BlindDetectionCA_CombIndicator_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetectionCA_CombIndicator_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetectionCACombIndicatorR17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Release:
				(*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetectionCA_CombIndicator_r17_Setup:
				(*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Setup = new(PDCCH_BlindDetectionCA_CombIndicator_r17)
				if err := (*ie.Pdcch_BlindDetectionCA_CombIndicator_r17).Setup.Decode(dx); err != nil {
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
				{Name: "simultaneousSR-PUSCH-diffPUCCH-Groups-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtSimultaneousSRPUSCHDiffPUCCHGroupsR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousSR_PUSCH_DiffPUCCH_Groups_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "intraBandNC-PRACH-simulTx-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtIntraBandNCPRACHSimulTxR17Constraints)
			if err != nil {
				return err
			}
			ie.IntraBandNC_PRACH_SimulTx_r17 = &v
		}
	}

	// Extension group 7:
	if seq.IsExtensionPresent(7) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdcch-BlindDetection4-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdcch_BlindDetection4_r17 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PDCCH_BlindDetection4_r17
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtPdcchBlindDetection4R17Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Pdcch_BlindDetection4_r17).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Release:
				(*ie.Pdcch_BlindDetection4_r17).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Pdcch_BlindDetection4_r17_Setup:
				(*ie.Pdcch_BlindDetection4_r17).Setup = new(PDCCH_BlindDetection4_r17)
				if err := (*ie.Pdcch_BlindDetection4_r17).Setup.Decode(dx); err != nil {
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
				{Name: "simultaneousPUCCH-PUSCH-SamePriority-r17", Optional: true},
				{Name: "simultaneousPUCCH-PUSCH-SamePriority-SecondaryPUCCHgroup-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePriorityR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousPUCCH_PUSCH_SamePriority_r17 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtSimultaneousPUCCHPUSCHSamePrioritySecondaryPUCCHgroupR17Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousPUCCH_PUSCH_SamePriority_SecondaryPUCCHgroup_r17 = &v
		}
	}

	// Extension group 9:
	if seq.IsExtensionPresent(9) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ncr-RNTI-r18", Optional: true},
				{Name: "cellDTRX-DCI-config-r18", Optional: true},
				{Name: "twoQCL-TypeD-ForMultiDCI-r18", Optional: true},
				{Name: "enableType1HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
				{Name: "enableType2HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
				{Name: "enableType3HARQ-ACK-MuxForDL-AssignmentAfterUL-Grant-r18", Optional: true},
				{Name: "enableDiffPUCCH-Resource-r18", Optional: true},
				{Name: "enableDiffCB-Size-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Ncr_RNTI_r18 = new(RNTI_Value)
			if err := ie.Ncr_RNTI_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CellDTRX_DCI_Config_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *CellDTRX_DCI_Config_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtCellDTRXDCIConfigR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.CellDTRX_DCI_Config_r18).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Release:
				(*ie.CellDTRX_DCI_Config_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_CellDTRX_DCI_Config_r18_Setup:
				(*ie.CellDTRX_DCI_Config_r18).Setup = new(CellDTRX_DCI_Config_r18)
				if err := (*ie.CellDTRX_DCI_Config_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtTwoQCLTypeDForMultiDCIR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoQCL_TypeD_ForMultiDCI_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtEnableType1HARQACKMuxForDLAssignmentAfterULGrantR18Constraints)
			if err != nil {
				return err
			}
			ie.EnableType1HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtEnableType2HARQACKMuxForDLAssignmentAfterULGrantR18Constraints)
			if err != nil {
				return err
			}
			ie.EnableType2HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtEnableType3HARQACKMuxForDLAssignmentAfterULGrantR18Constraints)
			if err != nil {
				return err
			}
			ie.EnableType3HARQ_ACK_MuxForDL_AssignmentAfterUL_Grant_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtEnableDiffPUCCHResourceR18Constraints)
			if err != nil {
				return err
			}
			ie.EnableDiffPUCCH_Resource_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(physicalCellGroupConfigExtEnableDiffCBSizeR18Constraints)
			if err != nil {
				return err
			}
			ie.EnableDiffCB_Size_r18 = &v
		}
	}

	// Extension group 10:
	if seq.IsExtensionPresent(10) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "lpwus-Config-r19", Optional: true},
				{Name: "adaptSSB-PeriodicityIndication-RNTI-r19", Optional: true},
				{Name: "adaptSSB-SizeDCI-2-9-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Lpwus_Config_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *LPWUS_Config_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtLpwusConfigR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Lpwus_Config_r19).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Release:
				(*ie.Lpwus_Config_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_Lpwus_Config_r19_Setup:
				(*ie.Lpwus_Config_r19).Setup = new(LPWUS_Config_r19)
				if err := (*ie.Lpwus_Config_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.AdaptSSB_PeriodicityIndication_RNTI_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *RNTI_Value
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtAdaptSSBPeriodicityIndicationRNTIR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Release:
				(*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_AdaptSSB_PeriodicityIndication_RNTI_r19_Setup:
				(*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Setup = new(RNTI_Value)
				if err := (*ie.AdaptSSB_PeriodicityIndication_RNTI_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.AdaptSSB_SizeDCI_2_9_r19 = &struct {
				Choice  int
				Release *struct{}
				Setup   *SizeDCI_2_9_r19
			}{}
			choiceDec := dx.NewChoiceDecoder(physicalCellGroupConfigExtAdaptSSBSizeDCI29R19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.AdaptSSB_SizeDCI_2_9_r19).Choice = int(index)
			switch index {
			case PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Release:
				(*ie.AdaptSSB_SizeDCI_2_9_r19).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PhysicalCellGroupConfig_Ext_AdaptSSB_SizeDCI_2_9_r19_Setup:
				(*ie.AdaptSSB_SizeDCI_2_9_r19).Setup = new(SizeDCI_2_9_r19)
				if err := (*ie.AdaptSSB_SizeDCI_2_9_r19).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
