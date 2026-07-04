// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: Phy-ParametersFRX-Diff (line 23092).

var phyParametersFRXDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "dynamicSFI", Optional: true},
		{Name: "dummy1", Optional: true},
		{Name: "twoFL-DMRS", Optional: true},
		{Name: "dummy2", Optional: true},
		{Name: "dummy3", Optional: true},
		{Name: "supportedDMRS-TypeDL", Optional: true},
		{Name: "supportedDMRS-TypeUL", Optional: true},
		{Name: "semiOpenLoopCSI", Optional: true},
		{Name: "csi-ReportWithoutPMI", Optional: true},
		{Name: "csi-ReportWithoutCQI", Optional: true},
		{Name: "onePortsPTRS", Optional: true},
		{Name: "twoPUCCH-F0-2-ConsecSymbols", Optional: true},
		{Name: "pucch-F2-WithFH", Optional: true},
		{Name: "pucch-F3-WithFH", Optional: true},
		{Name: "pucch-F4-WithFH", Optional: true},
		{Name: "pucch-F0-2WithoutFH", Optional: true},
		{Name: "pucch-F1-3-4WithoutFH", Optional: true},
		{Name: "mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot", Optional: true},
		{Name: "uci-CodeBlockSegmentation", Optional: true},
		{Name: "onePUCCH-LongAndShortFormat", Optional: true},
		{Name: "twoPUCCH-AnyOthersInSlot", Optional: true},
		{Name: "intraSlotFreqHopping-PUSCH", Optional: true},
		{Name: "pusch-LBRM", Optional: true},
		{Name: "pdcch-BlindDetectionCA", Optional: true},
		{Name: "tpc-PUSCH-RNTI", Optional: true},
		{Name: "tpc-PUCCH-RNTI", Optional: true},
		{Name: "tpc-SRS-RNTI", Optional: true},
		{Name: "absoluteTPC-Command", Optional: true},
		{Name: "twoDifferentTPC-Loop-PUSCH", Optional: true},
		{Name: "twoDifferentTPC-Loop-PUCCH", Optional: true},
		{Name: "pusch-HalfPi-BPSK", Optional: true},
		{Name: "pucch-F3-4-HalfPi-BPSK", Optional: true},
		{Name: "almostContiguousCP-OFDM-UL", Optional: true},
		{Name: "sp-CSI-RS", Optional: true},
		{Name: "sp-CSI-IM", Optional: true},
		{Name: "tdd-MultiDL-UL-SwitchPerSlot", Optional: true},
		{Name: "multipleCORESET", Optional: true},
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

const (
	Phy_ParametersFRX_Diff_DynamicSFI_Supported = 0
)

var phyParametersFRXDiffDynamicSFIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_DynamicSFI_Supported},
}

var phyParametersFRXDiffDummy1Constraints = per.FixedSize(2)

var phyParametersFRXDiffTwoFLDMRSConstraints = per.FixedSize(2)

var phyParametersFRXDiffDummy2Constraints = per.FixedSize(2)

var phyParametersFRXDiffDummy3Constraints = per.FixedSize(2)

const (
	Phy_ParametersFRX_Diff_SupportedDMRS_TypeDL_Type1     = 0
	Phy_ParametersFRX_Diff_SupportedDMRS_TypeDL_Type1And2 = 1
)

var phyParametersFRXDiffSupportedDMRSTypeDLConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_SupportedDMRS_TypeDL_Type1, Phy_ParametersFRX_Diff_SupportedDMRS_TypeDL_Type1And2},
}

const (
	Phy_ParametersFRX_Diff_SupportedDMRS_TypeUL_Type1     = 0
	Phy_ParametersFRX_Diff_SupportedDMRS_TypeUL_Type1And2 = 1
)

var phyParametersFRXDiffSupportedDMRSTypeULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_SupportedDMRS_TypeUL_Type1, Phy_ParametersFRX_Diff_SupportedDMRS_TypeUL_Type1And2},
}

const (
	Phy_ParametersFRX_Diff_SemiOpenLoopCSI_Supported = 0
)

var phyParametersFRXDiffSemiOpenLoopCSIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_SemiOpenLoopCSI_Supported},
}

const (
	Phy_ParametersFRX_Diff_Csi_ReportWithoutPMI_Supported = 0
)

var phyParametersFRXDiffCsiReportWithoutPMIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Csi_ReportWithoutPMI_Supported},
}

const (
	Phy_ParametersFRX_Diff_Csi_ReportWithoutCQI_Supported = 0
)

var phyParametersFRXDiffCsiReportWithoutCQIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Csi_ReportWithoutCQI_Supported},
}

var phyParametersFRXDiffOnePortsPTRSConstraints = per.FixedSize(2)

const (
	Phy_ParametersFRX_Diff_TwoPUCCH_F0_2_ConsecSymbols_Supported = 0
)

var phyParametersFRXDiffTwoPUCCHF02ConsecSymbolsConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_TwoPUCCH_F0_2_ConsecSymbols_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F2_WithFH_Supported = 0
)

var phyParametersFRXDiffPucchF2WithFHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F2_WithFH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F3_WithFH_Supported = 0
)

var phyParametersFRXDiffPucchF3WithFHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F3_WithFH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F4_WithFH_Supported = 0
)

var phyParametersFRXDiffPucchF4WithFHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F4_WithFH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F0_2WithoutFH_NotSupported = 0
)

var phyParametersFRXDiffPucchF02WithoutFHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F0_2WithoutFH_NotSupported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F1_3_4WithoutFH_NotSupported = 0
)

var phyParametersFRXDiffPucchF134WithoutFHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F1_3_4WithoutFH_NotSupported},
}

const (
	Phy_ParametersFRX_Diff_Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_Supported = 0
)

var phyParametersFRXDiffMuxSRHARQACKCSIPUCCHMultiPerSlotConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_Supported},
}

const (
	Phy_ParametersFRX_Diff_Uci_CodeBlockSegmentation_Supported = 0
)

var phyParametersFRXDiffUciCodeBlockSegmentationConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Uci_CodeBlockSegmentation_Supported},
}

const (
	Phy_ParametersFRX_Diff_OnePUCCH_LongAndShortFormat_Supported = 0
)

var phyParametersFRXDiffOnePUCCHLongAndShortFormatConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_OnePUCCH_LongAndShortFormat_Supported},
}

const (
	Phy_ParametersFRX_Diff_TwoPUCCH_AnyOthersInSlot_Supported = 0
)

var phyParametersFRXDiffTwoPUCCHAnyOthersInSlotConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_TwoPUCCH_AnyOthersInSlot_Supported},
}

const (
	Phy_ParametersFRX_Diff_IntraSlotFreqHopping_PUSCH_Supported = 0
)

var phyParametersFRXDiffIntraSlotFreqHoppingPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_IntraSlotFreqHopping_PUSCH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pusch_LBRM_Supported = 0
)

var phyParametersFRXDiffPuschLBRMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pusch_LBRM_Supported},
}

var phyParametersFRXDiffPdcchBlindDetectionCAConstraints = per.Constrained(4, 16)

const (
	Phy_ParametersFRX_Diff_Tpc_PUSCH_RNTI_Supported = 0
)

var phyParametersFRXDiffTpcPUSCHRNTIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Tpc_PUSCH_RNTI_Supported},
}

const (
	Phy_ParametersFRX_Diff_Tpc_PUCCH_RNTI_Supported = 0
)

var phyParametersFRXDiffTpcPUCCHRNTIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Tpc_PUCCH_RNTI_Supported},
}

const (
	Phy_ParametersFRX_Diff_Tpc_SRS_RNTI_Supported = 0
)

var phyParametersFRXDiffTpcSRSRNTIConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Tpc_SRS_RNTI_Supported},
}

const (
	Phy_ParametersFRX_Diff_AbsoluteTPC_Command_Supported = 0
)

var phyParametersFRXDiffAbsoluteTPCCommandConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_AbsoluteTPC_Command_Supported},
}

const (
	Phy_ParametersFRX_Diff_TwoDifferentTPC_Loop_PUSCH_Supported = 0
)

var phyParametersFRXDiffTwoDifferentTPCLoopPUSCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_TwoDifferentTPC_Loop_PUSCH_Supported},
}

const (
	Phy_ParametersFRX_Diff_TwoDifferentTPC_Loop_PUCCH_Supported = 0
)

var phyParametersFRXDiffTwoDifferentTPCLoopPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_TwoDifferentTPC_Loop_PUCCH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pusch_HalfPi_BPSK_Supported = 0
)

var phyParametersFRXDiffPuschHalfPiBPSKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pusch_HalfPi_BPSK_Supported},
}

const (
	Phy_ParametersFRX_Diff_Pucch_F3_4_HalfPi_BPSK_Supported = 0
)

var phyParametersFRXDiffPucchF34HalfPiBPSKConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Pucch_F3_4_HalfPi_BPSK_Supported},
}

const (
	Phy_ParametersFRX_Diff_AlmostContiguousCP_OFDM_UL_Supported = 0
)

var phyParametersFRXDiffAlmostContiguousCPOFDMULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_AlmostContiguousCP_OFDM_UL_Supported},
}

const (
	Phy_ParametersFRX_Diff_Sp_CSI_RS_Supported = 0
)

var phyParametersFRXDiffSpCSIRSConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Sp_CSI_RS_Supported},
}

const (
	Phy_ParametersFRX_Diff_Sp_CSI_IM_Supported = 0
)

var phyParametersFRXDiffSpCSIIMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Sp_CSI_IM_Supported},
}

const (
	Phy_ParametersFRX_Diff_Tdd_MultiDL_UL_SwitchPerSlot_Supported = 0
)

var phyParametersFRXDiffTddMultiDLULSwitchPerSlotConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Tdd_MultiDL_UL_SwitchPerSlot_Supported},
}

const (
	Phy_ParametersFRX_Diff_MultipleCORESET_Supported = 0
)

var phyParametersFRXDiffMultipleCORESETConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_MultipleCORESET_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_PUCCH_Supported = 0
)

var phyParametersFRXDiffExtMuxSRHARQACKPUCCHConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_PUCCH_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Mux_MultipleGroupCtrlCH_Overlap_Supported = 0
)

var phyParametersFRXDiffExtMuxMultipleGroupCtrlCHOverlapConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Mux_MultipleGroupCtrlCH_Overlap_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeA_Supported = 0
)

var phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeAConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeA_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeB_Supported = 0
)

var phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Dl_SchedulingOffset_PDSCH_TypeB_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Ul_SchedulingOffset_Supported = 0
)

var phyParametersFRXDiffExtUlSchedulingOffsetConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Ul_SchedulingOffset_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Dl_64QAM_MCS_TableAlt_Supported = 0
)

var phyParametersFRXDiffExtDl64QAMMCSTableAltConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Dl_64QAM_MCS_TableAlt_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Ul_64QAM_MCS_TableAlt_Supported = 0
)

var phyParametersFRXDiffExtUl64QAMMCSTableAltConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Ul_64QAM_MCS_TableAlt_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Cqi_TableAlt_Supported = 0
)

var phyParametersFRXDiffExtCqiTableAltConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Cqi_TableAlt_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_OneFL_DMRS_TwoAdditionalDMRS_UL_Supported = 0
)

var phyParametersFRXDiffExtOneFLDMRSTwoAdditionalDMRSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_OneFL_DMRS_TwoAdditionalDMRS_UL_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_TwoFL_DMRS_TwoAdditionalDMRS_UL_Supported = 0
)

var phyParametersFRXDiffExtTwoFLDMRSTwoAdditionalDMRSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_TwoFL_DMRS_TwoAdditionalDMRS_UL_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_OneFL_DMRS_ThreeAdditionalDMRS_UL_Supported = 0
)

var phyParametersFRXDiffExtOneFLDMRSThreeAdditionalDMRSULConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_OneFL_DMRS_ThreeAdditionalDMRS_UL_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Mux_HARQ_ACK_PUSCH_DiffSymbol_Supported = 0
)

var phyParametersFRXDiffExtMuxHARQACKPUSCHDiffSymbolConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Mux_HARQ_ACK_PUSCH_DiffSymbol_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Type1_HARQ_ACK_Codebook_r16_Supported = 0
)

var phyParametersFRXDiffExtType1HARQACKCodebookR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Type1_HARQ_ACK_Codebook_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_EnhancedPowerControl_r16_Supported = 0
)

var phyParametersFRXDiffExtEnhancedPowerControlR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_EnhancedPowerControl_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_SimultaneousTCI_ActMultipleCC_r16_Supported = 0
)

var phyParametersFRXDiffExtSimultaneousTCIActMultipleCCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_SimultaneousTCI_ActMultipleCC_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_SimultaneousSpatialRelationMultipleCC_r16_Supported = 0
)

var phyParametersFRXDiffExtSimultaneousSpatialRelationMultipleCCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_SimultaneousSpatialRelationMultipleCC_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Cli_RSSI_FDM_DL_r16_Supported = 0
)

var phyParametersFRXDiffExtCliRSSIFDMDLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Cli_RSSI_FDM_DL_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Cli_SRS_RSRP_FDM_DL_r16_Supported = 0
)

var phyParametersFRXDiffExtCliSRSRSRPFDMDLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Cli_SRS_RSRP_FDM_DL_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_MaxLayersMIMO_Adaptation_r16_Supported = 0
)

var phyParametersFRXDiffExtMaxLayersMIMOAdaptationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_MaxLayersMIMO_Adaptation_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_AggregationFactorSPS_DL_r16_Supported = 0
)

var phyParametersFRXDiffExtAggregationFactorSPSDLR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_AggregationFactorSPS_DL_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_TwoTCI_Act_ServingCellInCC_List_r16_Supported = 0
)

var phyParametersFRXDiffExtTwoTCIActServingCellInCCListR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_TwoTCI_Act_ServingCellInCC_List_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Cri_RI_CQI_WithoutNon_PMI_PortInd_r16_Supported = 0
)

var phyParametersFRXDiffExtCriRICQIWithoutNonPMIPortIndR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Cri_RI_CQI_WithoutNon_PMI_PortInd_r16_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17_Supported = 0
)

var phyParametersFRXDiffExtCqi4BitsSubbandTNNonSharedSpectrumChAccessR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_MultipleCORESET_RedCap_r17_Supported = 0
)

var phyParametersFRXDiffExtMultipleCORESETRedCapR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_MultipleCORESET_RedCap_r17_Supported},
}

var phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sameSymbol", Optional: true},
		{Name: "diffSymbol", Optional: true},
	},
}

const (
	Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_SameSymbol_Supported = 0
)

var phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSameSymbolConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_SameSymbol_Supported},
}

const (
	Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_DiffSymbol_Supported = 0
)

var phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotDiffSymbolConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_DiffSymbol_Supported},
}

var phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberResWithinSlotAcrossCC-OneFR-r16", Optional: true},
		{Name: "maxNumberResAcrossCC-OneFR-r16", Optional: true},
	},
}

const (
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N2   = 0
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N4   = 1
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N8   = 2
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N12  = 3
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N16  = 4
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N32  = 5
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N64  = 6
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N128 = 7
)

var phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResWithinSlotAcrossCCOneFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N2, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N4, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N8, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N12, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N16, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N32, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N64, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResWithinSlotAcrossCC_OneFR_r16_N128},
}

const (
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N2   = 0
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N4   = 1
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N8   = 2
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N12  = 3
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N16  = 4
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N32  = 5
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N40  = 6
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N48  = 7
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N64  = 8
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N72  = 9
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N80  = 10
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N96  = 11
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N128 = 12
	Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N256 = 13
)

var phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResAcrossCCOneFRR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N2, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N4, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N8, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N12, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N16, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N32, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N40, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N48, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N64, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N72, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N80, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N96, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N128, Phy_ParametersFRX_Diff_Ext_MaxTotalResourcesForOneFreqRange_r16_MaxNumberResAcrossCC_OneFR_r16_N256},
}

type Phy_ParametersFRX_Diff struct {
	DynamicSFI                             *int64
	Dummy1                                 *per.BitString
	TwoFL_DMRS                             *per.BitString
	Dummy2                                 *per.BitString
	Dummy3                                 *per.BitString
	SupportedDMRS_TypeDL                   *int64
	SupportedDMRS_TypeUL                   *int64
	SemiOpenLoopCSI                        *int64
	Csi_ReportWithoutPMI                   *int64
	Csi_ReportWithoutCQI                   *int64
	OnePortsPTRS                           *per.BitString
	TwoPUCCH_F0_2_ConsecSymbols            *int64
	Pucch_F2_WithFH                        *int64
	Pucch_F3_WithFH                        *int64
	Pucch_F4_WithFH                        *int64
	Pucch_F0_2WithoutFH                    *int64
	Pucch_F1_3_4WithoutFH                  *int64
	Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot *int64
	Uci_CodeBlockSegmentation              *int64
	OnePUCCH_LongAndShortFormat            *int64
	TwoPUCCH_AnyOthersInSlot               *int64
	IntraSlotFreqHopping_PUSCH             *int64
	Pusch_LBRM                             *int64
	Pdcch_BlindDetectionCA                 *int64
	Tpc_PUSCH_RNTI                         *int64
	Tpc_PUCCH_RNTI                         *int64
	Tpc_SRS_RNTI                           *int64
	AbsoluteTPC_Command                    *int64
	TwoDifferentTPC_Loop_PUSCH             *int64
	TwoDifferentTPC_Loop_PUCCH             *int64
	Pusch_HalfPi_BPSK                      *int64
	Pucch_F3_4_HalfPi_BPSK                 *int64
	AlmostContiguousCP_OFDM_UL             *int64
	Sp_CSI_RS                              *int64
	Sp_CSI_IM                              *int64
	Tdd_MultiDL_UL_SwitchPerSlot           *int64
	MultipleCORESET                        *int64
	Csi_RS_IM_ReceptionForFeedback         *CSI_RS_IM_ReceptionForFeedback
	Csi_RS_ProcFrameworkForSRS             *CSI_RS_ProcFrameworkForSRS
	Csi_ReportFramework                    *CSI_ReportFramework
	Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot  *struct {
		SameSymbol *int64
		DiffSymbol *int64
	}
	Mux_SR_HARQ_ACK_PUCCH             *int64
	Mux_MultipleGroupCtrlCH_Overlap   *int64
	Dl_SchedulingOffset_PDSCH_TypeA   *int64
	Dl_SchedulingOffset_PDSCH_TypeB   *int64
	Ul_SchedulingOffset               *int64
	Dl_64QAM_MCS_TableAlt             *int64
	Ul_64QAM_MCS_TableAlt             *int64
	Cqi_TableAlt                      *int64
	OneFL_DMRS_TwoAdditionalDMRS_UL   *int64
	TwoFL_DMRS_TwoAdditionalDMRS_UL   *int64
	OneFL_DMRS_ThreeAdditionalDMRS_UL *int64
	Pdcch_BlindDetectionNRDC          *struct {
		Pdcch_BlindDetectionMCG_UE int64
		Pdcch_BlindDetectionSCG_UE int64
	}
	Mux_HARQ_ACK_PUSCH_DiffSymbol             *int64
	Type1_HARQ_ACK_Codebook_r16               *int64
	EnhancedPowerControl_r16                  *int64
	SimultaneousTCI_ActMultipleCC_r16         *int64
	SimultaneousSpatialRelationMultipleCC_r16 *int64
	Cli_RSSI_FDM_DL_r16                       *int64
	Cli_SRS_RSRP_FDM_DL_r16                   *int64
	MaxLayersMIMO_Adaptation_r16              *int64
	AggregationFactorSPS_DL_r16               *int64
	MaxTotalResourcesForOneFreqRange_r16      *struct {
		MaxNumberResWithinSlotAcrossCC_OneFR_r16 *int64
		MaxNumberResAcrossCC_OneFR_r16           *int64
	}
	Csi_ReportFrameworkExt_r16                        *CSI_ReportFrameworkExt_r16
	TwoTCI_Act_ServingCellInCC_List_r16               *int64
	Cri_RI_CQI_WithoutNon_PMI_PortInd_r16             *int64
	Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 *int64
	MultipleCORESET_RedCap_r17                        *int64
}

func (ie *Phy_ParametersFRX_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(phyParametersFRXDiffConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil || ie.Mux_SR_HARQ_ACK_PUCCH != nil || ie.Mux_MultipleGroupCtrlCH_Overlap != nil || ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil || ie.Dl_64QAM_MCS_TableAlt != nil || ie.Ul_64QAM_MCS_TableAlt != nil || ie.Cqi_TableAlt != nil || ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil
	hasExtGroup1 := ie.Pdcch_BlindDetectionNRDC != nil || ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil
	hasExtGroup2 := ie.Type1_HARQ_ACK_Codebook_r16 != nil || ie.EnhancedPowerControl_r16 != nil || ie.SimultaneousTCI_ActMultipleCC_r16 != nil || ie.SimultaneousSpatialRelationMultipleCC_r16 != nil || ie.Cli_RSSI_FDM_DL_r16 != nil || ie.Cli_SRS_RSRP_FDM_DL_r16 != nil || ie.MaxLayersMIMO_Adaptation_r16 != nil || ie.AggregationFactorSPS_DL_r16 != nil || ie.MaxTotalResourcesForOneFreqRange_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil
	hasExtGroup3 := ie.TwoTCI_Act_ServingCellInCC_List_r16 != nil
	hasExtGroup4 := ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil
	hasExtGroup5 := ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil
	hasExtGroup6 := ie.MultipleCORESET_RedCap_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4 || hasExtGroup5 || hasExtGroup6

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DynamicSFI != nil, ie.Dummy1 != nil, ie.TwoFL_DMRS != nil, ie.Dummy2 != nil, ie.Dummy3 != nil, ie.SupportedDMRS_TypeDL != nil, ie.SupportedDMRS_TypeUL != nil, ie.SemiOpenLoopCSI != nil, ie.Csi_ReportWithoutPMI != nil, ie.Csi_ReportWithoutCQI != nil, ie.OnePortsPTRS != nil, ie.TwoPUCCH_F0_2_ConsecSymbols != nil, ie.Pucch_F2_WithFH != nil, ie.Pucch_F3_WithFH != nil, ie.Pucch_F4_WithFH != nil, ie.Pucch_F0_2WithoutFH != nil, ie.Pucch_F1_3_4WithoutFH != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil, ie.Uci_CodeBlockSegmentation != nil, ie.OnePUCCH_LongAndShortFormat != nil, ie.TwoPUCCH_AnyOthersInSlot != nil, ie.IntraSlotFreqHopping_PUSCH != nil, ie.Pusch_LBRM != nil, ie.Pdcch_BlindDetectionCA != nil, ie.Tpc_PUSCH_RNTI != nil, ie.Tpc_PUCCH_RNTI != nil, ie.Tpc_SRS_RNTI != nil, ie.AbsoluteTPC_Command != nil, ie.TwoDifferentTPC_Loop_PUSCH != nil, ie.TwoDifferentTPC_Loop_PUCCH != nil, ie.Pusch_HalfPi_BPSK != nil, ie.Pucch_F3_4_HalfPi_BPSK != nil, ie.AlmostContiguousCP_OFDM_UL != nil, ie.Sp_CSI_RS != nil, ie.Sp_CSI_IM != nil, ie.Tdd_MultiDL_UL_SwitchPerSlot != nil, ie.MultipleCORESET != nil}); err != nil {
		return err
	}

	// 3. dynamicSFI: enumerated
	{
		if ie.DynamicSFI != nil {
			if err := e.EncodeEnumerated(*ie.DynamicSFI, phyParametersFRXDiffDynamicSFIConstraints); err != nil {
				return err
			}
		}
	}

	// 4. dummy1: bit-string
	{
		if ie.Dummy1 != nil {
			if err := e.EncodeBitString(*ie.Dummy1, phyParametersFRXDiffDummy1Constraints); err != nil {
				return err
			}
		}
	}

	// 5. twoFL-DMRS: bit-string
	{
		if ie.TwoFL_DMRS != nil {
			if err := e.EncodeBitString(*ie.TwoFL_DMRS, phyParametersFRXDiffTwoFLDMRSConstraints); err != nil {
				return err
			}
		}
	}

	// 6. dummy2: bit-string
	{
		if ie.Dummy2 != nil {
			if err := e.EncodeBitString(*ie.Dummy2, phyParametersFRXDiffDummy2Constraints); err != nil {
				return err
			}
		}
	}

	// 7. dummy3: bit-string
	{
		if ie.Dummy3 != nil {
			if err := e.EncodeBitString(*ie.Dummy3, phyParametersFRXDiffDummy3Constraints); err != nil {
				return err
			}
		}
	}

	// 8. supportedDMRS-TypeDL: enumerated
	{
		if ie.SupportedDMRS_TypeDL != nil {
			if err := e.EncodeEnumerated(*ie.SupportedDMRS_TypeDL, phyParametersFRXDiffSupportedDMRSTypeDLConstraints); err != nil {
				return err
			}
		}
	}

	// 9. supportedDMRS-TypeUL: enumerated
	{
		if ie.SupportedDMRS_TypeUL != nil {
			if err := e.EncodeEnumerated(*ie.SupportedDMRS_TypeUL, phyParametersFRXDiffSupportedDMRSTypeULConstraints); err != nil {
				return err
			}
		}
	}

	// 10. semiOpenLoopCSI: enumerated
	{
		if ie.SemiOpenLoopCSI != nil {
			if err := e.EncodeEnumerated(*ie.SemiOpenLoopCSI, phyParametersFRXDiffSemiOpenLoopCSIConstraints); err != nil {
				return err
			}
		}
	}

	// 11. csi-ReportWithoutPMI: enumerated
	{
		if ie.Csi_ReportWithoutPMI != nil {
			if err := e.EncodeEnumerated(*ie.Csi_ReportWithoutPMI, phyParametersFRXDiffCsiReportWithoutPMIConstraints); err != nil {
				return err
			}
		}
	}

	// 12. csi-ReportWithoutCQI: enumerated
	{
		if ie.Csi_ReportWithoutCQI != nil {
			if err := e.EncodeEnumerated(*ie.Csi_ReportWithoutCQI, phyParametersFRXDiffCsiReportWithoutCQIConstraints); err != nil {
				return err
			}
		}
	}

	// 13. onePortsPTRS: bit-string
	{
		if ie.OnePortsPTRS != nil {
			if err := e.EncodeBitString(*ie.OnePortsPTRS, phyParametersFRXDiffOnePortsPTRSConstraints); err != nil {
				return err
			}
		}
	}

	// 14. twoPUCCH-F0-2-ConsecSymbols: enumerated
	{
		if ie.TwoPUCCH_F0_2_ConsecSymbols != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_F0_2_ConsecSymbols, phyParametersFRXDiffTwoPUCCHF02ConsecSymbolsConstraints); err != nil {
				return err
			}
		}
	}

	// 15. pucch-F2-WithFH: enumerated
	{
		if ie.Pucch_F2_WithFH != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F2_WithFH, phyParametersFRXDiffPucchF2WithFHConstraints); err != nil {
				return err
			}
		}
	}

	// 16. pucch-F3-WithFH: enumerated
	{
		if ie.Pucch_F3_WithFH != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F3_WithFH, phyParametersFRXDiffPucchF3WithFHConstraints); err != nil {
				return err
			}
		}
	}

	// 17. pucch-F4-WithFH: enumerated
	{
		if ie.Pucch_F4_WithFH != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F4_WithFH, phyParametersFRXDiffPucchF4WithFHConstraints); err != nil {
				return err
			}
		}
	}

	// 18. pucch-F0-2WithoutFH: enumerated
	{
		if ie.Pucch_F0_2WithoutFH != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F0_2WithoutFH, phyParametersFRXDiffPucchF02WithoutFHConstraints); err != nil {
				return err
			}
		}
	}

	// 19. pucch-F1-3-4WithoutFH: enumerated
	{
		if ie.Pucch_F1_3_4WithoutFH != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F1_3_4WithoutFH, phyParametersFRXDiffPucchF134WithoutFHConstraints); err != nil {
				return err
			}
		}
	}

	// 20. mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot: enumerated
	{
		if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil {
			if err := e.EncodeEnumerated(*ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot, phyParametersFRXDiffMuxSRHARQACKCSIPUCCHMultiPerSlotConstraints); err != nil {
				return err
			}
		}
	}

	// 21. uci-CodeBlockSegmentation: enumerated
	{
		if ie.Uci_CodeBlockSegmentation != nil {
			if err := e.EncodeEnumerated(*ie.Uci_CodeBlockSegmentation, phyParametersFRXDiffUciCodeBlockSegmentationConstraints); err != nil {
				return err
			}
		}
	}

	// 22. onePUCCH-LongAndShortFormat: enumerated
	{
		if ie.OnePUCCH_LongAndShortFormat != nil {
			if err := e.EncodeEnumerated(*ie.OnePUCCH_LongAndShortFormat, phyParametersFRXDiffOnePUCCHLongAndShortFormatConstraints); err != nil {
				return err
			}
		}
	}

	// 23. twoPUCCH-AnyOthersInSlot: enumerated
	{
		if ie.TwoPUCCH_AnyOthersInSlot != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_AnyOthersInSlot, phyParametersFRXDiffTwoPUCCHAnyOthersInSlotConstraints); err != nil {
				return err
			}
		}
	}

	// 24. intraSlotFreqHopping-PUSCH: enumerated
	{
		if ie.IntraSlotFreqHopping_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.IntraSlotFreqHopping_PUSCH, phyParametersFRXDiffIntraSlotFreqHoppingPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 25. pusch-LBRM: enumerated
	{
		if ie.Pusch_LBRM != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_LBRM, phyParametersFRXDiffPuschLBRMConstraints); err != nil {
				return err
			}
		}
	}

	// 26. pdcch-BlindDetectionCA: integer
	{
		if ie.Pdcch_BlindDetectionCA != nil {
			if err := e.EncodeInteger(*ie.Pdcch_BlindDetectionCA, phyParametersFRXDiffPdcchBlindDetectionCAConstraints); err != nil {
				return err
			}
		}
	}

	// 27. tpc-PUSCH-RNTI: enumerated
	{
		if ie.Tpc_PUSCH_RNTI != nil {
			if err := e.EncodeEnumerated(*ie.Tpc_PUSCH_RNTI, phyParametersFRXDiffTpcPUSCHRNTIConstraints); err != nil {
				return err
			}
		}
	}

	// 28. tpc-PUCCH-RNTI: enumerated
	{
		if ie.Tpc_PUCCH_RNTI != nil {
			if err := e.EncodeEnumerated(*ie.Tpc_PUCCH_RNTI, phyParametersFRXDiffTpcPUCCHRNTIConstraints); err != nil {
				return err
			}
		}
	}

	// 29. tpc-SRS-RNTI: enumerated
	{
		if ie.Tpc_SRS_RNTI != nil {
			if err := e.EncodeEnumerated(*ie.Tpc_SRS_RNTI, phyParametersFRXDiffTpcSRSRNTIConstraints); err != nil {
				return err
			}
		}
	}

	// 30. absoluteTPC-Command: enumerated
	{
		if ie.AbsoluteTPC_Command != nil {
			if err := e.EncodeEnumerated(*ie.AbsoluteTPC_Command, phyParametersFRXDiffAbsoluteTPCCommandConstraints); err != nil {
				return err
			}
		}
	}

	// 31. twoDifferentTPC-Loop-PUSCH: enumerated
	{
		if ie.TwoDifferentTPC_Loop_PUSCH != nil {
			if err := e.EncodeEnumerated(*ie.TwoDifferentTPC_Loop_PUSCH, phyParametersFRXDiffTwoDifferentTPCLoopPUSCHConstraints); err != nil {
				return err
			}
		}
	}

	// 32. twoDifferentTPC-Loop-PUCCH: enumerated
	{
		if ie.TwoDifferentTPC_Loop_PUCCH != nil {
			if err := e.EncodeEnumerated(*ie.TwoDifferentTPC_Loop_PUCCH, phyParametersFRXDiffTwoDifferentTPCLoopPUCCHConstraints); err != nil {
				return err
			}
		}
	}

	// 33. pusch-HalfPi-BPSK: enumerated
	{
		if ie.Pusch_HalfPi_BPSK != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_HalfPi_BPSK, phyParametersFRXDiffPuschHalfPiBPSKConstraints); err != nil {
				return err
			}
		}
	}

	// 34. pucch-F3-4-HalfPi-BPSK: enumerated
	{
		if ie.Pucch_F3_4_HalfPi_BPSK != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F3_4_HalfPi_BPSK, phyParametersFRXDiffPucchF34HalfPiBPSKConstraints); err != nil {
				return err
			}
		}
	}

	// 35. almostContiguousCP-OFDM-UL: enumerated
	{
		if ie.AlmostContiguousCP_OFDM_UL != nil {
			if err := e.EncodeEnumerated(*ie.AlmostContiguousCP_OFDM_UL, phyParametersFRXDiffAlmostContiguousCPOFDMULConstraints); err != nil {
				return err
			}
		}
	}

	// 36. sp-CSI-RS: enumerated
	{
		if ie.Sp_CSI_RS != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_RS, phyParametersFRXDiffSpCSIRSConstraints); err != nil {
				return err
			}
		}
	}

	// 37. sp-CSI-IM: enumerated
	{
		if ie.Sp_CSI_IM != nil {
			if err := e.EncodeEnumerated(*ie.Sp_CSI_IM, phyParametersFRXDiffSpCSIIMConstraints); err != nil {
				return err
			}
		}
	}

	// 38. tdd-MultiDL-UL-SwitchPerSlot: enumerated
	{
		if ie.Tdd_MultiDL_UL_SwitchPerSlot != nil {
			if err := e.EncodeEnumerated(*ie.Tdd_MultiDL_UL_SwitchPerSlot, phyParametersFRXDiffTddMultiDLULSwitchPerSlotConstraints); err != nil {
				return err
			}
		}
	}

	// 39. multipleCORESET: enumerated
	{
		if ie.MultipleCORESET != nil {
			if err := e.EncodeEnumerated(*ie.MultipleCORESET, phyParametersFRXDiffMultipleCORESETConstraints); err != nil {
				return err
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
					{Name: "csi-RS-IM-ReceptionForFeedback", Optional: true},
					{Name: "csi-RS-ProcFrameworkForSRS", Optional: true},
					{Name: "csi-ReportFramework", Optional: true},
					{Name: "mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot", Optional: true},
					{Name: "mux-SR-HARQ-ACK-PUCCH", Optional: true},
					{Name: "mux-MultipleGroupCtrlCH-Overlap", Optional: true},
					{Name: "dl-SchedulingOffset-PDSCH-TypeA", Optional: true},
					{Name: "dl-SchedulingOffset-PDSCH-TypeB", Optional: true},
					{Name: "ul-SchedulingOffset", Optional: true},
					{Name: "dl-64QAM-MCS-TableAlt", Optional: true},
					{Name: "ul-64QAM-MCS-TableAlt", Optional: true},
					{Name: "cqi-TableAlt", Optional: true},
					{Name: "oneFL-DMRS-TwoAdditionalDMRS-UL", Optional: true},
					{Name: "twoFL-DMRS-TwoAdditionalDMRS-UL", Optional: true},
					{Name: "oneFL-DMRS-ThreeAdditionalDMRS-UL", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Csi_RS_IM_ReceptionForFeedback != nil, ie.Csi_RS_ProcFrameworkForSRS != nil, ie.Csi_ReportFramework != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil, ie.Mux_SR_HARQ_ACK_PUCCH != nil, ie.Mux_MultipleGroupCtrlCH_Overlap != nil, ie.Dl_SchedulingOffset_PDSCH_TypeA != nil, ie.Dl_SchedulingOffset_PDSCH_TypeB != nil, ie.Ul_SchedulingOffset != nil, ie.Dl_64QAM_MCS_TableAlt != nil, ie.Ul_64QAM_MCS_TableAlt != nil, ie.Cqi_TableAlt != nil, ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil}); err != nil {
				return err
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

			if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil {
				c := ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot
				phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq := ex.NewSequenceEncoder(phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotConstraints)
				if err := phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq.EncodePreamble([]bool{c.SameSymbol != nil, c.DiffSymbol != nil}); err != nil {
					return err
				}
				if c.SameSymbol != nil {
					if err := ex.EncodeEnumerated((*c.SameSymbol), phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSameSymbolConstraints); err != nil {
						return err
					}
				}
				if c.DiffSymbol != nil {
					if err := ex.EncodeEnumerated((*c.DiffSymbol), phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotDiffSymbolConstraints); err != nil {
						return err
					}
				}
			}

			if ie.Mux_SR_HARQ_ACK_PUCCH != nil {
				if err := ex.EncodeEnumerated(*ie.Mux_SR_HARQ_ACK_PUCCH, phyParametersFRXDiffExtMuxSRHARQACKPUCCHConstraints); err != nil {
					return err
				}
			}

			if ie.Mux_MultipleGroupCtrlCH_Overlap != nil {
				if err := ex.EncodeEnumerated(*ie.Mux_MultipleGroupCtrlCH_Overlap, phyParametersFRXDiffExtMuxMultipleGroupCtrlCHOverlapConstraints); err != nil {
					return err
				}
			}

			if ie.Dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_SchedulingOffset_PDSCH_TypeA, phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeAConstraints); err != nil {
					return err
				}
			}

			if ie.Dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_SchedulingOffset_PDSCH_TypeB, phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeBConstraints); err != nil {
					return err
				}
			}

			if ie.Ul_SchedulingOffset != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_SchedulingOffset, phyParametersFRXDiffExtUlSchedulingOffsetConstraints); err != nil {
					return err
				}
			}

			if ie.Dl_64QAM_MCS_TableAlt != nil {
				if err := ex.EncodeEnumerated(*ie.Dl_64QAM_MCS_TableAlt, phyParametersFRXDiffExtDl64QAMMCSTableAltConstraints); err != nil {
					return err
				}
			}

			if ie.Ul_64QAM_MCS_TableAlt != nil {
				if err := ex.EncodeEnumerated(*ie.Ul_64QAM_MCS_TableAlt, phyParametersFRXDiffExtUl64QAMMCSTableAltConstraints); err != nil {
					return err
				}
			}

			if ie.Cqi_TableAlt != nil {
				if err := ex.EncodeEnumerated(*ie.Cqi_TableAlt, phyParametersFRXDiffExtCqiTableAltConstraints); err != nil {
					return err
				}
			}

			if ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err := ex.EncodeEnumerated(*ie.OneFL_DMRS_TwoAdditionalDMRS_UL, phyParametersFRXDiffExtOneFLDMRSTwoAdditionalDMRSULConstraints); err != nil {
					return err
				}
			}

			if ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err := ex.EncodeEnumerated(*ie.TwoFL_DMRS_TwoAdditionalDMRS_UL, phyParametersFRXDiffExtTwoFLDMRSTwoAdditionalDMRSULConstraints); err != nil {
					return err
				}
			}

			if ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil {
				if err := ex.EncodeEnumerated(*ie.OneFL_DMRS_ThreeAdditionalDMRS_UL, phyParametersFRXDiffExtOneFLDMRSThreeAdditionalDMRSULConstraints); err != nil {
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
					{Name: "pdcch-BlindDetectionNRDC", Optional: true},
					{Name: "mux-HARQ-ACK-PUSCH-DiffSymbol", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Pdcch_BlindDetectionNRDC != nil, ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil}); err != nil {
				return err
			}

			if ie.Pdcch_BlindDetectionNRDC != nil {
				c := ie.Pdcch_BlindDetectionNRDC
				if err := ex.EncodeInteger(c.Pdcch_BlindDetectionMCG_UE, per.Constrained(1, 15)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.Pdcch_BlindDetectionSCG_UE, per.Constrained(1, 15)); err != nil {
					return err
				}
			}

			if ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil {
				if err := ex.EncodeEnumerated(*ie.Mux_HARQ_ACK_PUSCH_DiffSymbol, phyParametersFRXDiffExtMuxHARQACKPUSCHDiffSymbolConstraints); err != nil {
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
					{Name: "type1-HARQ-ACK-Codebook-r16", Optional: true},
					{Name: "enhancedPowerControl-r16", Optional: true},
					{Name: "simultaneousTCI-ActMultipleCC-r16", Optional: true},
					{Name: "simultaneousSpatialRelationMultipleCC-r16", Optional: true},
					{Name: "cli-RSSI-FDM-DL-r16", Optional: true},
					{Name: "cli-SRS-RSRP-FDM-DL-r16", Optional: true},
					{Name: "maxLayersMIMO-Adaptation-r16", Optional: true},
					{Name: "aggregationFactorSPS-DL-r16", Optional: true},
					{Name: "maxTotalResourcesForOneFreqRange-r16", Optional: true},
					{Name: "csi-ReportFrameworkExt-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Type1_HARQ_ACK_Codebook_r16 != nil, ie.EnhancedPowerControl_r16 != nil, ie.SimultaneousTCI_ActMultipleCC_r16 != nil, ie.SimultaneousSpatialRelationMultipleCC_r16 != nil, ie.Cli_RSSI_FDM_DL_r16 != nil, ie.Cli_SRS_RSRP_FDM_DL_r16 != nil, ie.MaxLayersMIMO_Adaptation_r16 != nil, ie.AggregationFactorSPS_DL_r16 != nil, ie.MaxTotalResourcesForOneFreqRange_r16 != nil, ie.Csi_ReportFrameworkExt_r16 != nil}); err != nil {
				return err
			}

			if ie.Type1_HARQ_ACK_Codebook_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Type1_HARQ_ACK_Codebook_r16, phyParametersFRXDiffExtType1HARQACKCodebookR16Constraints); err != nil {
					return err
				}
			}

			if ie.EnhancedPowerControl_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.EnhancedPowerControl_r16, phyParametersFRXDiffExtEnhancedPowerControlR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousTCI_ActMultipleCC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousTCI_ActMultipleCC_r16, phyParametersFRXDiffExtSimultaneousTCIActMultipleCCR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousSpatialRelationMultipleCC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousSpatialRelationMultipleCC_r16, phyParametersFRXDiffExtSimultaneousSpatialRelationMultipleCCR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cli_RSSI_FDM_DL_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cli_RSSI_FDM_DL_r16, phyParametersFRXDiffExtCliRSSIFDMDLR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cli_SRS_RSRP_FDM_DL_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cli_SRS_RSRP_FDM_DL_r16, phyParametersFRXDiffExtCliSRSRSRPFDMDLR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxLayersMIMO_Adaptation_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.MaxLayersMIMO_Adaptation_r16, phyParametersFRXDiffExtMaxLayersMIMOAdaptationR16Constraints); err != nil {
					return err
				}
			}

			if ie.AggregationFactorSPS_DL_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.AggregationFactorSPS_DL_r16, phyParametersFRXDiffExtAggregationFactorSPSDLR16Constraints); err != nil {
					return err
				}
			}

			if ie.MaxTotalResourcesForOneFreqRange_r16 != nil {
				c := ie.MaxTotalResourcesForOneFreqRange_r16
				phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq := ex.NewSequenceEncoder(phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Constraints)
				if err := phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq.EncodePreamble([]bool{c.MaxNumberResWithinSlotAcrossCC_OneFR_r16 != nil, c.MaxNumberResAcrossCC_OneFR_r16 != nil}); err != nil {
					return err
				}
				if c.MaxNumberResWithinSlotAcrossCC_OneFR_r16 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumberResWithinSlotAcrossCC_OneFR_r16), phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResWithinSlotAcrossCCOneFRR16Constraints); err != nil {
						return err
					}
				}
				if c.MaxNumberResAcrossCC_OneFR_r16 != nil {
					if err := ex.EncodeEnumerated((*c.MaxNumberResAcrossCC_OneFR_r16), phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResAcrossCCOneFRR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Csi_ReportFrameworkExt_r16 != nil {
				if err := ie.Csi_ReportFrameworkExt_r16.Encode(ex); err != nil {
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
					{Name: "twoTCI-Act-servingCellInCC-List-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TwoTCI_Act_ServingCellInCC_List_r16 != nil}); err != nil {
				return err
			}

			if ie.TwoTCI_Act_ServingCellInCC_List_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.TwoTCI_Act_ServingCellInCC_List_r16, phyParametersFRXDiffExtTwoTCIActServingCellInCCListR16Constraints); err != nil {
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
					{Name: "cri-RI-CQI-WithoutNon-PMI-PortInd-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil}); err != nil {
				return err
			}

			if ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16, phyParametersFRXDiffExtCriRICQIWithoutNonPMIPortIndR16Constraints); err != nil {
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
					{Name: "cqi-4-BitsSubbandTN-NonSharedSpectrumChAccess-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil}); err != nil {
				return err
			}

			if ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17, phyParametersFRXDiffExtCqi4BitsSubbandTNNonSharedSpectrumChAccessR17Constraints); err != nil {
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
					{Name: "multipleCORESET-RedCap-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MultipleCORESET_RedCap_r17 != nil}); err != nil {
				return err
			}

			if ie.MultipleCORESET_RedCap_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipleCORESET_RedCap_r17, phyParametersFRXDiffExtMultipleCORESETRedCapR17Constraints); err != nil {
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

func (ie *Phy_ParametersFRX_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(phyParametersFRXDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. dynamicSFI: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffDynamicSFIConstraints)
			if err != nil {
				return err
			}
			ie.DynamicSFI = &idx
		}
	}

	// 4. dummy1: bit-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeBitString(phyParametersFRXDiffDummy1Constraints)
			if err != nil {
				return err
			}
			ie.Dummy1 = &v
		}
	}

	// 5. twoFL-DMRS: bit-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBitString(phyParametersFRXDiffTwoFLDMRSConstraints)
			if err != nil {
				return err
			}
			ie.TwoFL_DMRS = &v
		}
	}

	// 6. dummy2: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(phyParametersFRXDiffDummy2Constraints)
			if err != nil {
				return err
			}
			ie.Dummy2 = &v
		}
	}

	// 7. dummy3: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(phyParametersFRXDiffDummy3Constraints)
			if err != nil {
				return err
			}
			ie.Dummy3 = &v
		}
	}

	// 8. supportedDMRS-TypeDL: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffSupportedDMRSTypeDLConstraints)
			if err != nil {
				return err
			}
			ie.SupportedDMRS_TypeDL = &idx
		}
	}

	// 9. supportedDMRS-TypeUL: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffSupportedDMRSTypeULConstraints)
			if err != nil {
				return err
			}
			ie.SupportedDMRS_TypeUL = &idx
		}
	}

	// 10. semiOpenLoopCSI: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffSemiOpenLoopCSIConstraints)
			if err != nil {
				return err
			}
			ie.SemiOpenLoopCSI = &idx
		}
	}

	// 11. csi-ReportWithoutPMI: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffCsiReportWithoutPMIConstraints)
			if err != nil {
				return err
			}
			ie.Csi_ReportWithoutPMI = &idx
		}
	}

	// 12. csi-ReportWithoutCQI: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffCsiReportWithoutCQIConstraints)
			if err != nil {
				return err
			}
			ie.Csi_ReportWithoutCQI = &idx
		}
	}

	// 13. onePortsPTRS: bit-string
	{
		if seq.IsComponentPresent(10) {
			v, err := d.DecodeBitString(phyParametersFRXDiffOnePortsPTRSConstraints)
			if err != nil {
				return err
			}
			ie.OnePortsPTRS = &v
		}
	}

	// 14. twoPUCCH-F0-2-ConsecSymbols: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTwoPUCCHF02ConsecSymbolsConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_F0_2_ConsecSymbols = &idx
		}
	}

	// 15. pucch-F2-WithFH: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF2WithFHConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F2_WithFH = &idx
		}
	}

	// 16. pucch-F3-WithFH: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF3WithFHConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F3_WithFH = &idx
		}
	}

	// 17. pucch-F4-WithFH: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF4WithFHConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F4_WithFH = &idx
		}
	}

	// 18. pucch-F0-2WithoutFH: enumerated
	{
		if seq.IsComponentPresent(15) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF02WithoutFHConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F0_2WithoutFH = &idx
		}
	}

	// 19. pucch-F1-3-4WithoutFH: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF134WithoutFHConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F1_3_4WithoutFH = &idx
		}
	}

	// 20. mux-SR-HARQ-ACK-CSI-PUCCH-MultiPerSlot: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffMuxSRHARQACKCSIPUCCHMultiPerSlotConstraints)
			if err != nil {
				return err
			}
			ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot = &idx
		}
	}

	// 21. uci-CodeBlockSegmentation: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffUciCodeBlockSegmentationConstraints)
			if err != nil {
				return err
			}
			ie.Uci_CodeBlockSegmentation = &idx
		}
	}

	// 22. onePUCCH-LongAndShortFormat: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffOnePUCCHLongAndShortFormatConstraints)
			if err != nil {
				return err
			}
			ie.OnePUCCH_LongAndShortFormat = &idx
		}
	}

	// 23. twoPUCCH-AnyOthersInSlot: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTwoPUCCHAnyOthersInSlotConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_AnyOthersInSlot = &idx
		}
	}

	// 24. intraSlotFreqHopping-PUSCH: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffIntraSlotFreqHoppingPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.IntraSlotFreqHopping_PUSCH = &idx
		}
	}

	// 25. pusch-LBRM: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPuschLBRMConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_LBRM = &idx
		}
	}

	// 26. pdcch-BlindDetectionCA: integer
	{
		if seq.IsComponentPresent(23) {
			v, err := d.DecodeInteger(phyParametersFRXDiffPdcchBlindDetectionCAConstraints)
			if err != nil {
				return err
			}
			ie.Pdcch_BlindDetectionCA = &v
		}
	}

	// 27. tpc-PUSCH-RNTI: enumerated
	{
		if seq.IsComponentPresent(24) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTpcPUSCHRNTIConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_PUSCH_RNTI = &idx
		}
	}

	// 28. tpc-PUCCH-RNTI: enumerated
	{
		if seq.IsComponentPresent(25) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTpcPUCCHRNTIConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_PUCCH_RNTI = &idx
		}
	}

	// 29. tpc-SRS-RNTI: enumerated
	{
		if seq.IsComponentPresent(26) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTpcSRSRNTIConstraints)
			if err != nil {
				return err
			}
			ie.Tpc_SRS_RNTI = &idx
		}
	}

	// 30. absoluteTPC-Command: enumerated
	{
		if seq.IsComponentPresent(27) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffAbsoluteTPCCommandConstraints)
			if err != nil {
				return err
			}
			ie.AbsoluteTPC_Command = &idx
		}
	}

	// 31. twoDifferentTPC-Loop-PUSCH: enumerated
	{
		if seq.IsComponentPresent(28) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTwoDifferentTPCLoopPUSCHConstraints)
			if err != nil {
				return err
			}
			ie.TwoDifferentTPC_Loop_PUSCH = &idx
		}
	}

	// 32. twoDifferentTPC-Loop-PUCCH: enumerated
	{
		if seq.IsComponentPresent(29) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTwoDifferentTPCLoopPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.TwoDifferentTPC_Loop_PUCCH = &idx
		}
	}

	// 33. pusch-HalfPi-BPSK: enumerated
	{
		if seq.IsComponentPresent(30) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPuschHalfPiBPSKConstraints)
			if err != nil {
				return err
			}
			ie.Pusch_HalfPi_BPSK = &idx
		}
	}

	// 34. pucch-F3-4-HalfPi-BPSK: enumerated
	{
		if seq.IsComponentPresent(31) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffPucchF34HalfPiBPSKConstraints)
			if err != nil {
				return err
			}
			ie.Pucch_F3_4_HalfPi_BPSK = &idx
		}
	}

	// 35. almostContiguousCP-OFDM-UL: enumerated
	{
		if seq.IsComponentPresent(32) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffAlmostContiguousCPOFDMULConstraints)
			if err != nil {
				return err
			}
			ie.AlmostContiguousCP_OFDM_UL = &idx
		}
	}

	// 36. sp-CSI-RS: enumerated
	{
		if seq.IsComponentPresent(33) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffSpCSIRSConstraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_RS = &idx
		}
	}

	// 37. sp-CSI-IM: enumerated
	{
		if seq.IsComponentPresent(34) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffSpCSIIMConstraints)
			if err != nil {
				return err
			}
			ie.Sp_CSI_IM = &idx
		}
	}

	// 38. tdd-MultiDL-UL-SwitchPerSlot: enumerated
	{
		if seq.IsComponentPresent(35) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffTddMultiDLULSwitchPerSlotConstraints)
			if err != nil {
				return err
			}
			ie.Tdd_MultiDL_UL_SwitchPerSlot = &idx
		}
	}

	// 39. multipleCORESET: enumerated
	{
		if seq.IsComponentPresent(36) {
			idx, err := d.DecodeEnumerated(phyParametersFRXDiffMultipleCORESETConstraints)
			if err != nil {
				return err
			}
			ie.MultipleCORESET = &idx
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
				{Name: "csi-RS-IM-ReceptionForFeedback", Optional: true},
				{Name: "csi-RS-ProcFrameworkForSRS", Optional: true},
				{Name: "csi-ReportFramework", Optional: true},
				{Name: "mux-SR-HARQ-ACK-CSI-PUCCH-OncePerSlot", Optional: true},
				{Name: "mux-SR-HARQ-ACK-PUCCH", Optional: true},
				{Name: "mux-MultipleGroupCtrlCH-Overlap", Optional: true},
				{Name: "dl-SchedulingOffset-PDSCH-TypeA", Optional: true},
				{Name: "dl-SchedulingOffset-PDSCH-TypeB", Optional: true},
				{Name: "ul-SchedulingOffset", Optional: true},
				{Name: "dl-64QAM-MCS-TableAlt", Optional: true},
				{Name: "ul-64QAM-MCS-TableAlt", Optional: true},
				{Name: "cqi-TableAlt", Optional: true},
				{Name: "oneFL-DMRS-TwoAdditionalDMRS-UL", Optional: true},
				{Name: "twoFL-DMRS-TwoAdditionalDMRS-UL", Optional: true},
				{Name: "oneFL-DMRS-ThreeAdditionalDMRS-UL", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Csi_RS_IM_ReceptionForFeedback = new(CSI_RS_IM_ReceptionForFeedback)
			if err := ie.Csi_RS_IM_ReceptionForFeedback.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Csi_RS_ProcFrameworkForSRS = new(CSI_RS_ProcFrameworkForSRS)
			if err := ie.Csi_RS_ProcFrameworkForSRS.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Csi_ReportFramework = new(CSI_ReportFramework)
			if err := ie.Csi_ReportFramework.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot = &struct {
				SameSymbol *int64
				DiffSymbol *int64
			}{}
			c := ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot
			phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq := dx.NewSequenceDecoder(phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotConstraints)
			if err := phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq.IsComponentPresent(0) {
				c.SameSymbol = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSameSymbolConstraints)
				if err != nil {
					return err
				}
				(*c.SameSymbol) = v
			}
			if phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotSeq.IsComponentPresent(1) {
				c.DiffSymbol = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMuxSRHARQACKCSIPUCCHOncePerSlotDiffSymbolConstraints)
				if err != nil {
					return err
				}
				(*c.DiffSymbol) = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMuxSRHARQACKPUCCHConstraints)
			if err != nil {
				return err
			}
			ie.Mux_SR_HARQ_ACK_PUCCH = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMuxMultipleGroupCtrlCHOverlapConstraints)
			if err != nil {
				return err
			}
			ie.Mux_MultipleGroupCtrlCH_Overlap = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeAConstraints)
			if err != nil {
				return err
			}
			ie.Dl_SchedulingOffset_PDSCH_TypeA = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtDlSchedulingOffsetPDSCHTypeBConstraints)
			if err != nil {
				return err
			}
			ie.Dl_SchedulingOffset_PDSCH_TypeB = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtUlSchedulingOffsetConstraints)
			if err != nil {
				return err
			}
			ie.Ul_SchedulingOffset = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtDl64QAMMCSTableAltConstraints)
			if err != nil {
				return err
			}
			ie.Dl_64QAM_MCS_TableAlt = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtUl64QAMMCSTableAltConstraints)
			if err != nil {
				return err
			}
			ie.Ul_64QAM_MCS_TableAlt = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtCqiTableAltConstraints)
			if err != nil {
				return err
			}
			ie.Cqi_TableAlt = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtOneFLDMRSTwoAdditionalDMRSULConstraints)
			if err != nil {
				return err
			}
			ie.OneFL_DMRS_TwoAdditionalDMRS_UL = &v
		}

		if groupSeq.IsComponentPresent(13) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtTwoFLDMRSTwoAdditionalDMRSULConstraints)
			if err != nil {
				return err
			}
			ie.TwoFL_DMRS_TwoAdditionalDMRS_UL = &v
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtOneFLDMRSThreeAdditionalDMRSULConstraints)
			if err != nil {
				return err
			}
			ie.OneFL_DMRS_ThreeAdditionalDMRS_UL = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pdcch-BlindDetectionNRDC", Optional: true},
				{Name: "mux-HARQ-ACK-PUSCH-DiffSymbol", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Pdcch_BlindDetectionNRDC = &struct {
				Pdcch_BlindDetectionMCG_UE int64
				Pdcch_BlindDetectionSCG_UE int64
			}{}
			c := ie.Pdcch_BlindDetectionNRDC
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionMCG_UE = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionSCG_UE = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMuxHARQACKPUSCHDiffSymbolConstraints)
			if err != nil {
				return err
			}
			ie.Mux_HARQ_ACK_PUSCH_DiffSymbol = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "type1-HARQ-ACK-Codebook-r16", Optional: true},
				{Name: "enhancedPowerControl-r16", Optional: true},
				{Name: "simultaneousTCI-ActMultipleCC-r16", Optional: true},
				{Name: "simultaneousSpatialRelationMultipleCC-r16", Optional: true},
				{Name: "cli-RSSI-FDM-DL-r16", Optional: true},
				{Name: "cli-SRS-RSRP-FDM-DL-r16", Optional: true},
				{Name: "maxLayersMIMO-Adaptation-r16", Optional: true},
				{Name: "aggregationFactorSPS-DL-r16", Optional: true},
				{Name: "maxTotalResourcesForOneFreqRange-r16", Optional: true},
				{Name: "csi-ReportFrameworkExt-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtType1HARQACKCodebookR16Constraints)
			if err != nil {
				return err
			}
			ie.Type1_HARQ_ACK_Codebook_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtEnhancedPowerControlR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedPowerControl_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtSimultaneousTCIActMultipleCCR16Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousTCI_ActMultipleCC_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtSimultaneousSpatialRelationMultipleCCR16Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousSpatialRelationMultipleCC_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtCliRSSIFDMDLR16Constraints)
			if err != nil {
				return err
			}
			ie.Cli_RSSI_FDM_DL_r16 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtCliSRSRSRPFDMDLR16Constraints)
			if err != nil {
				return err
			}
			ie.Cli_SRS_RSRP_FDM_DL_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMaxLayersMIMOAdaptationR16Constraints)
			if err != nil {
				return err
			}
			ie.MaxLayersMIMO_Adaptation_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtAggregationFactorSPSDLR16Constraints)
			if err != nil {
				return err
			}
			ie.AggregationFactorSPS_DL_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			ie.MaxTotalResourcesForOneFreqRange_r16 = &struct {
				MaxNumberResWithinSlotAcrossCC_OneFR_r16 *int64
				MaxNumberResAcrossCC_OneFR_r16           *int64
			}{}
			c := ie.MaxTotalResourcesForOneFreqRange_r16
			phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq := dx.NewSequenceDecoder(phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Constraints)
			if err := phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq.IsComponentPresent(0) {
				c.MaxNumberResWithinSlotAcrossCC_OneFR_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResWithinSlotAcrossCCOneFRR16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberResWithinSlotAcrossCC_OneFR_r16) = v
			}
			if phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16Seq.IsComponentPresent(1) {
				c.MaxNumberResAcrossCC_OneFR_r16 = new(int64)
				v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMaxTotalResourcesForOneFreqRangeR16MaxNumberResAcrossCCOneFRR16Constraints)
				if err != nil {
					return err
				}
				(*c.MaxNumberResAcrossCC_OneFR_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(9) {
			ie.Csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
			if err := ie.Csi_ReportFrameworkExt_r16.Decode(dx); err != nil {
				return err
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
				{Name: "twoTCI-Act-servingCellInCC-List-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtTwoTCIActServingCellInCCListR16Constraints)
			if err != nil {
				return err
			}
			ie.TwoTCI_Act_ServingCellInCC_List_r16 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cri-RI-CQI-WithoutNon-PMI-PortInd-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtCriRICQIWithoutNonPMIPortIndR16Constraints)
			if err != nil {
				return err
			}
			ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 = &v
		}
	}

	// Extension group 5:
	if seq.IsExtensionPresent(5) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "cqi-4-BitsSubbandTN-NonSharedSpectrumChAccess-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtCqi4BitsSubbandTNNonSharedSpectrumChAccessR17Constraints)
			if err != nil {
				return err
			}
			ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 = &v
		}
	}

	// Extension group 6:
	if seq.IsExtensionPresent(6) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "multipleCORESET-RedCap-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(phyParametersFRXDiffExtMultipleCORESETRedCapR17Constraints)
			if err != nil {
				return err
			}
			ie.MultipleCORESET_RedCap_r17 = &v
		}
	}

	return nil
}
