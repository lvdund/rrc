// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandSidelink-r16 (line 25104).

var bandSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "freqBandSidelink-r16"},
		{Name: "sl-Reception-r16", Optional: true},
		{Name: "sl-TransmissionMode1-r16", Optional: true},
		{Name: "sync-Sidelink-r16", Optional: true},
		{Name: "sl-Tx-256QAM-r16", Optional: true},
		{Name: "psfch-FormatZeroSidelink-r16", Optional: true},
		{Name: "lowSE-64QAM-MCS-TableSidelink-r16", Optional: true},
		{Name: "enb-sync-Sidelink-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	BandSidelink_r16_Sl_Tx_256QAM_r16_Supported = 0
)

var bandSidelinkR16SlTx256QAMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_Tx_256QAM_r16_Supported},
}

const (
	BandSidelink_r16_LowSE_64QAM_MCS_TableSidelink_r16_Supported = 0
)

var bandSidelinkR16LowSE64QAMMCSTableSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_LowSE_64QAM_MCS_TableSidelink_r16_Supported},
}

const (
	BandSidelink_r16_Enb_Sync_Sidelink_r16_Supported = 0
)

var bandSidelinkR16EnbSyncSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Enb_Sync_Sidelink_r16_Supported},
}

const (
	BandSidelink_r16_Ext_FewerSymbolSlotSidelink_r16_Supported = 0
)

var bandSidelinkR16ExtFewerSymbolSlotSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_FewerSymbolSlotSidelink_r16_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_OpenLoopPC_RSRP_ReportSidelink_r16_Supported = 0
)

var bandSidelinkR16ExtSlOpenLoopPCRSRPReportSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_OpenLoopPC_RSRP_ReportSidelink_r16_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_Rx_256QAM_r16_Supported = 0
)

var bandSidelinkR16ExtSlRx256QAMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_Rx_256QAM_r16_Supported},
}

const (
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc2       = 0
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc3       = 1
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc5_v1820 = 2
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare5    = 3
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare4    = 4
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare3    = 5
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare2    = 6
	BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare1    = 7
)

var bandSidelinkR16ExtUePowerClassSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc2, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc3, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Pc5_v1820, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare5, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare4, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare3, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare2, BandSidelink_r16_Ext_Ue_PowerClassSidelink_r16_Spare1},
}

const (
	BandSidelink_r16_Ext_Enb_Sync_Sidelink_v1710_Supported = 0
)

var bandSidelinkR16ExtEnbSyncSidelinkV1710Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Enb_Sync_Sidelink_v1710_Supported},
}

const (
	BandSidelink_r16_Ext_Rx_IUC_Scheme1_PreferredMode2Sidelink_r17_Supported = 0
)

var bandSidelinkR16ExtRxIUCScheme1PreferredMode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Rx_IUC_Scheme1_PreferredMode2Sidelink_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17_Supported = 0
)

var bandSidelinkR16ExtRxIUCScheme1NonPreferredMode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N5  = 0
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N15 = 1
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N25 = 2
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N32 = 3
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N35 = 4
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N45 = 5
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N50 = 6
	BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N64 = 7
)

var bandSidelinkR16ExtRxIUCScheme2Mode2SidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N5, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N15, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N25, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N32, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N35, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N45, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N50, BandSidelink_r16_Ext_Rx_IUC_Scheme2_Mode2Sidelink_r17_N64},
}

const (
	BandSidelink_r16_Ext_Rx_IUC_Scheme1_SCI_r17_Supported = 0
)

var bandSidelinkR16ExtRxIUCScheme1SCIR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Rx_IUC_Scheme1_SCI_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Rx_IUC_Scheme1_SCI_ExplicitReq_r17_Supported = 0
)

var bandSidelinkR16ExtRxIUCScheme1SCIExplicitReqR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Rx_IUC_Scheme1_SCI_ExplicitReq_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_RxInSharedResourcePool_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSRxInSharedResourcePoolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_RxInSharedResourcePool_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxInSharedResourcePool_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSTxInSharedResourcePoolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxInSharedResourcePool_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxScheme1InDedicatedResourcePool_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSTxScheme1InDedicatedResourcePoolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxScheme1InDedicatedResourcePool_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxScheme2InDedicatedResourcePool_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSTxScheme2InDedicatedResourcePoolR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxScheme2InDedicatedResourcePool_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt1 = 0
	BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt2 = 1
	BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt3 = 2
)

var bandSidelinkR16ExtSlPRSCongestionCtrlR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt1, BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt2, BandSidelink_r16_Ext_Sl_PRS_CongestionCtrl_r18_Cpt3},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxRandomSelection_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSTxRandomSelectionR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxRandomSelection_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxUsingFullSensing_r18_Value1 = 0
	BandSidelink_r16_Ext_Sl_PRS_TxUsingFullSensing_r18_Value2 = 1
)

var bandSidelinkR16ExtSlPRSTxUsingFullSensingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxUsingFullSensing_r18_Value1, BandSidelink_r16_Ext_Sl_PRS_TxUsingFullSensing_r18_Value2},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_RxForBandWithSL_CA_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSRxForBandWithSLCAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_RxForBandWithSL_CA_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_TxForBandWithSL_CA_r18_Supported = 0
)

var bandSidelinkR16ExtSlPRSTxForBandWithSLCAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_TxForBandWithSL_CA_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_DynamicSharingTxRx_r18_Supported = 0
)

var bandSidelinkR16ExtSlDynamicSharingTxRxR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_DynamicSharingTxRx_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_CA_Synchronization_r18_Supported = 0
)

var bandSidelinkR16ExtSlCASynchronizationR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_CA_Synchronization_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_ReceptionIntraCarrierGuardBand_r18_Supported = 0
)

var bandSidelinkR16ExtSlReceptionIntraCarrierGuardBandR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_ReceptionIntraCarrierGuardBand_r18_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_PathlossBasedOLPC_SL_RSRP_Report_r18_Supported = 0
)

var bandSidelinkR16ExtSlPathlossBasedOLPCSLRSRPReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PathlossBasedOLPC_SL_RSRP_Report_r18_Supported},
}

var bandSidelinkR16SlReceptionR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-RxProcessSidelink-r16"},
		{Name: "pscch-RxSidelink-r16"},
		{Name: "scs-CP-PatternRxSidelink-r16", Optional: true},
		{Name: "extendedCP-RxSidelink-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N16 = 0
	BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N24 = 1
	BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N32 = 2
	BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N48 = 3
	BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N64 = 4
)

var bandSidelinkR16SlReceptionR16HarqRxProcessSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N16, BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N24, BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N32, BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N48, BandSidelink_r16_Sl_Reception_r16_Harq_RxProcessSidelink_r16_N64},
}

const (
	BandSidelink_r16_Sl_Reception_r16_Pscch_RxSidelink_r16_Value1 = 0
	BandSidelink_r16_Sl_Reception_r16_Pscch_RxSidelink_r16_Value2 = 1
)

var bandSidelinkR16SlReceptionR16PscchRxSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_Reception_r16_Pscch_RxSidelink_r16_Value1, BandSidelink_r16_Sl_Reception_r16_Pscch_RxSidelink_r16_Value2},
}

var bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r16"},
		{Name: "fr2-r16"},
	},
}

const (
	BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr1_r16 = 0
	BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr2_r16 = 1
)

var bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
	},
}

var bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Sl_Reception_r16_ExtendedCP_RxSidelink_r16_Supported = 0
)

var bandSidelinkR16SlReceptionR16ExtendedCPRxSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_Reception_r16_ExtendedCP_RxSidelink_r16_Supported},
}

var bandSidelinkR16SlTransmissionMode1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-TxProcessModeOneSidelink-r16"},
		{Name: "scs-CP-PatternTxSidelinkModeOne-r16"},
		{Name: "extendedCP-TxSidelink-r16", Optional: true},
		{Name: "harq-ReportOnPUCCH-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_TxProcessModeOneSidelink_r16_N8  = 0
	BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_TxProcessModeOneSidelink_r16_N16 = 1
)

var bandSidelinkR16SlTransmissionMode1R16HarqTxProcessModeOneSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_TxProcessModeOneSidelink_r16_N8, BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_TxProcessModeOneSidelink_r16_N16},
}

var bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r16"},
		{Name: "fr2-r16"},
	},
}

const (
	BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr1_r16 = 0
	BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr2_r16 = 1
)

var bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
	},
}

var bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Sl_TransmissionMode1_r16_ExtendedCP_TxSidelink_r16_Supported = 0
)

var bandSidelinkR16SlTransmissionMode1R16ExtendedCPTxSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_TransmissionMode1_r16_ExtendedCP_TxSidelink_r16_Supported},
}

const (
	BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_ReportOnPUCCH_r16_Supported = 0
)

var bandSidelinkR16SlTransmissionMode1R16HarqReportOnPUCCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sl_TransmissionMode1_r16_Harq_ReportOnPUCCH_r16_Supported},
}

var bandSidelinkR16SyncSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "gNB-Sync-r16", Optional: true},
		{Name: "gNB-GNSS-UE-SyncWithPriorityOnGNB-ENB-r16", Optional: true},
		{Name: "gNB-GNSS-UE-SyncWithPriorityOnGNSS-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Sync_Sidelink_r16_GNB_Sync_r16_Supported = 0
)

var bandSidelinkR16SyncSidelinkR16GNBSyncR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sync_Sidelink_r16_GNB_Sync_r16_Supported},
}

const (
	BandSidelink_r16_Sync_Sidelink_r16_GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16_Supported = 0
)

var bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNBENBR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sync_Sidelink_r16_GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16_Supported},
}

const (
	BandSidelink_r16_Sync_Sidelink_r16_GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16_Supported = 0
)

var bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNSSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Sync_Sidelink_r16_GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16_Supported},
}

const (
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N5  = 0
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N15 = 1
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N25 = 2
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N32 = 3
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N35 = 4
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N45 = 5
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N50 = 6
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N64 = 7
)

var bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchRxNumberConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N5, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N15, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N25, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N32, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N35, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N45, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N50, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_RxNumber_N64},
}

const (
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N4  = 0
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N8  = 1
	BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N16 = 2
)

var bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchTxNumberConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N4, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N8, BandSidelink_r16_Psfch_FormatZeroSidelink_r16_Psfch_TxNumber_N16},
}

var bandSidelinkR16ExtSlTransmissionMode2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-TxProcessModeTwoSidelink-r16"},
		{Name: "scs-CP-PatternTxSidelinkModeTwo-r16", Optional: true},
		{Name: "dl-openLoopPC-Sidelink-r16", Optional: true},
	},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Harq_TxProcessModeTwoSidelink_r16_N8  = 0
	BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Harq_TxProcessModeTwoSidelink_r16_N16 = 1
)

var bandSidelinkR16ExtSlTransmissionMode2R16HarqTxProcessModeTwoSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Harq_TxProcessModeTwoSidelink_r16_N8, BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Harq_TxProcessModeTwoSidelink_r16_N16},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Scs_CP_PatternTxSidelinkModeTwo_r16_Supported = 0
)

var bandSidelinkR16ExtSlTransmissionMode2R16ScsCPPatternTxSidelinkModeTwoR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Scs_CP_PatternTxSidelinkModeTwo_r16_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Dl_OpenLoopPC_Sidelink_r16_Supported = 0
)

var bandSidelinkR16ExtSlTransmissionMode2R16DlOpenLoopPCSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_r16_Dl_OpenLoopPC_Sidelink_r16_Supported},
}

var bandSidelinkR16ExtCongestionControlSidelinkR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cbr-ReportSidelink-r16", Optional: true},
		{Name: "cbr-CR-TimeLimitSidelink-r16"},
	},
}

const (
	BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_ReportSidelink_r16_Supported = 0
)

var bandSidelinkR16ExtCongestionControlSidelinkR16CbrReportSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_ReportSidelink_r16_Supported},
}

const (
	BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_CR_TimeLimitSidelink_r16_Time1 = 0
	BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_CR_TimeLimitSidelink_r16_Time2 = 1
)

var bandSidelinkR16ExtCongestionControlSidelinkR16CbrCRTimeLimitSidelinkR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_CR_TimeLimitSidelink_r16_Time1, BandSidelink_r16_Ext_CongestionControlSidelink_r16_Cbr_CR_TimeLimitSidelink_r16_Time2},
}

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "harq-TxProcessModeTwoSidelink-r17"},
		{Name: "scs-CP-PatternTxSidelinkModeTwo-r17", Optional: true},
		{Name: "extendedCP-Mode2Random-r17", Optional: true},
		{Name: "dl-openLoopPC-Sidelink-r17", Optional: true},
	},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Harq_TxProcessModeTwoSidelink_r17_N8  = 0
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Harq_TxProcessModeTwoSidelink_r17_N16 = 1
)

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17HarqTxProcessModeTwoSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Harq_TxProcessModeTwoSidelink_r17_N8, BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Harq_TxProcessModeTwoSidelink_r17_N16},
}

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r17"},
		{Name: "fr2-r17"},
	},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17 = 0
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17 = 1
)

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r17", Optional: true},
		{Name: "scs-30kHz-r17", Optional: true},
		{Name: "scs-60kHz-r17", Optional: true},
	},
}

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r17", Optional: true},
		{Name: "scs-120kHz-r17", Optional: true},
	},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_ExtendedCP_Mode2Random_r17_Supported = 0
)

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ExtendedCPMode2RandomR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_ExtendedCP_Mode2Random_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Dl_OpenLoopPC_Sidelink_r17_Supported = 0
)

var bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17DlOpenLoopPCSidelinkR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Dl_OpenLoopPC_Sidelink_r17_Supported},
}

var bandSidelinkR16ExtSyncSidelinkV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sync-GNSS-r17", Optional: true},
		{Name: "gNB-Sync-r17", Optional: true},
		{Name: "gNB-GNSS-UE-SyncWithPriorityOnGNB-ENB-r17", Optional: true},
		{Name: "gNB-GNSS-UE-SyncWithPriorityOnGNSS-r17", Optional: true},
	},
}

const (
	BandSidelink_r16_Ext_Sync_Sidelink_v1710_Sync_GNSS_r17_Supported = 0
)

var bandSidelinkR16ExtSyncSidelinkV1710SyncGNSSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sync_Sidelink_v1710_Sync_GNSS_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_Sync_r17_Supported = 0
)

var bandSidelinkR16ExtSyncSidelinkV1710GNBSyncR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_Sync_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17_Supported = 0
)

var bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNBENBR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17_Supported},
}

const (
	BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17_Supported = 0
)

var bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNSSR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sync_Sidelink_v1710_GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17_Supported},
}

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r18"},
		{Name: "fr2-r18"},
	},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18 = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18 = 1
)

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz5   = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz10  = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz20  = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz40  = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz50  = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz80  = 5
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz100 = 6
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz5, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz10, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz20, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz40, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz50, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz80, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18_Mhz100},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz50  = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz100 = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz200 = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz400 = 3
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz50, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz100, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz200, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18_Mhz400},
}

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r18"},
		{Name: "fr2-r18"},
	},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18 = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18 = 1
)

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N1  = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N2  = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N4  = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N6  = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N8  = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N12 = 5
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N16 = 6
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N24 = 7
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N1, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N2, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N4, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N6, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N8, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N12, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N16, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18_N24},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N1   = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N2   = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N4   = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N6   = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N8   = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N12  = 5
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N16  = 6
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N24  = 7
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N32  = 8
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N48  = 9
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N64  = 10
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N128 = 11
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N1, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N2, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N4, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N6, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N8, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N12, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N16, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N24, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N32, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N48, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N64, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18_N128},
}

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "fr1-r18"},
		{Name: "fr2-r18"},
	},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18 = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18 = 1
)

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N1 = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N2 = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N3 = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N4 = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N6 = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N8 = 5
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N1, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N2, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N3, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N4, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N6, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18_N8},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N1  = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N2  = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N4  = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N8  = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N12 = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N16 = 5
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N24 = 6
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N32 = 7
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N48 = 8
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N64 = 9
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N1, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N2, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N4, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N8, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N12, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N16, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N24, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N32, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N48, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18_N64},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms20  = 0
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms30  = 1
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms40  = 2
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms50  = 3
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms80  = 4
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms100 = 5
	BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms160 = 6
)

var bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MinTimeAfterEndofSlotCarryActiveSLPRSResourcesR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms20, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms30, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms40, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms50, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms80, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms100, BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18_Ms160},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_NumOfSupportedRxPSCCH_PerSlot_r18_Value1 = 0
	BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_NumOfSupportedRxPSCCH_PerSlot_r18_Value2 = 1
)

var bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18NumOfSupportedRxPSCCHPerSlotR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_NumOfSupportedRxPSCCH_PerSlot_r18_Value1, BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_NumOfSupportedRxPSCCH_PerSlot_r18_Value2},
}

const (
	BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_SupportedCP_TypeFor60kHzSCS_r18_Ncp       = 0
	BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_SupportedCP_TypeFor60kHzSCS_r18_NcpAndECP = 1
)

var bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18SupportedCPTypeFor60kHzSCSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_SupportedCP_TypeFor60kHzSCS_r18_Ncp, BandSidelink_r16_Ext_Sl_PRS_RxInDedicatedResourcePool_r18_SupportedCP_TypeFor60kHzSCS_r18_NcpAndECP},
}

const (
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz20 = 0
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz30 = 1
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz40 = 2
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz50 = 3
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz60 = 4
	BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz70 = 5
)

var bandSidelinkR16ExtSlCACommunicationR18TotalBandwidthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz20, BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz30, BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz40, BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz50, BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz60, BandSidelink_r16_Ext_Sl_CA_Communication_r18_TotalBandwidth_r18_Mhz70},
}

const (
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N5   = 0
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N15  = 1
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N25  = 2
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N32  = 3
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N35  = 4
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N45  = 5
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N50  = 6
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N64  = 7
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N100 = 8
)

var bandSidelinkR16ExtSlCAPSFCHR18RxPSFCHResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N5, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N15, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N25, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N32, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N35, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N45, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N50, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N64, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Rx_PSFCH_Resource_r18_N100},
}

const (
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N4  = 0
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N8  = 1
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N16 = 2
	BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N24 = 3
)

var bandSidelinkR16ExtSlCAPSFCHR18TxPSFCHResourceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N4, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N8, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N16, BandSidelink_r16_Ext_Sl_CA_PSFCH_r18_Tx_PSFCH_Resource_r18_N24},
}

type BandSidelink_r16 struct {
	FreqBandSidelink_r16 FreqBandIndicatorNR
	Sl_Reception_r16     *struct {
		Harq_RxProcessSidelink_r16   int64
		Pscch_RxSidelink_r16         int64
		Scs_CP_PatternRxSidelink_r16 *struct {
			Choice  int
			Fr1_r16 *struct {
				Scs_15kHz_r16 *per.BitString
				Scs_30kHz_r16 *per.BitString
				Scs_60kHz_r16 *per.BitString
			}
			Fr2_r16 *struct {
				Scs_60kHz_r16  *per.BitString
				Scs_120kHz_r16 *per.BitString
			}
		}
		ExtendedCP_RxSidelink_r16 *int64
	}
	Sl_TransmissionMode1_r16 *struct {
		Harq_TxProcessModeOneSidelink_r16   int64
		Scs_CP_PatternTxSidelinkModeOne_r16 struct {
			Choice  int
			Fr1_r16 *struct {
				Scs_15kHz_r16 *per.BitString
				Scs_30kHz_r16 *per.BitString
				Scs_60kHz_r16 *per.BitString
			}
			Fr2_r16 *struct {
				Scs_60kHz_r16  *per.BitString
				Scs_120kHz_r16 *per.BitString
			}
		}
		ExtendedCP_TxSidelink_r16 *int64
		Harq_ReportOnPUCCH_r16    *int64
	}
	Sync_Sidelink_r16 *struct {
		GNB_Sync_r16                              *int64
		GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 *int64
		GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    *int64
	}
	Sl_Tx_256QAM_r16             *int64
	Psfch_FormatZeroSidelink_r16 *struct {
		Psfch_RxNumber int64
		Psfch_TxNumber int64
	}
	LowSE_64QAM_MCS_TableSidelink_r16 *int64
	Enb_Sync_Sidelink_r16             *int64
	Sl_TransmissionMode2_r16          *struct {
		Harq_TxProcessModeTwoSidelink_r16   int64
		Scs_CP_PatternTxSidelinkModeTwo_r16 *int64
		Dl_OpenLoopPC_Sidelink_r16          *int64
	}
	CongestionControlSidelink_r16 *struct {
		Cbr_ReportSidelink_r16       *int64
		Cbr_CR_TimeLimitSidelink_r16 int64
	}
	FewerSymbolSlotSidelink_r16                      *int64
	Sl_OpenLoopPC_RSRP_ReportSidelink_r16            *int64
	Sl_Rx_256QAM_r16                                 *int64
	Ue_PowerClassSidelink_r16                        *int64
	Sl_TransmissionMode2_RandomResourceSelection_r17 *struct {
		Harq_TxProcessModeTwoSidelink_r17   int64
		Scs_CP_PatternTxSidelinkModeTwo_r17 *struct {
			Choice  int
			Fr1_r17 *struct {
				Scs_15kHz_r17 *per.BitString
				Scs_30kHz_r17 *per.BitString
				Scs_60kHz_r17 *per.BitString
			}
			Fr2_r17 *struct {
				Scs_60kHz_r17  *per.BitString
				Scs_120kHz_r17 *per.BitString
			}
		}
		ExtendedCP_Mode2Random_r17 *int64
		Dl_OpenLoopPC_Sidelink_r17 *int64
	}
	Sync_Sidelink_v1710 *struct {
		Sync_GNSS_r17                             *int64
		GNB_Sync_r17                              *int64
		GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 *int64
		GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    *int64
	}
	Enb_Sync_Sidelink_v1710                         *int64
	Rx_IUC_Scheme1_PreferredMode2Sidelink_r17       *int64
	Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17    *int64
	Rx_IUC_Scheme2_Mode2Sidelink_r17                *int64
	Rx_IUC_Scheme1_SCI_r17                          *int64
	Rx_IUC_Scheme1_SCI_ExplicitReq_r17              *int64
	SharedSpectrumChAccessParamsSidelinkPerBand_r18 *SharedSpectrumChAccessParamsSidelinkPerBand_r18
	Sl_PRS_CommonProcCapabilityPerBand_r18          *struct {
		MaxSL_PRS_Bandwidth_r18 struct {
			Choice  int
			Fr1_r18 *int64
			Fr2_r18 *int64
		}
		MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18 struct {
			Choice  int
			Fr1_r18 *int64
			Fr2_r18 *int64
		}
		MaxNumOfSlotsWithActiveSL_PRS_Resources_r18 struct {
			Choice  int
			Fr1_r18 *int64
			Fr2_r18 *int64
		}
		MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18 int64
	}
	Sl_PRS_RxInSharedResourcePool_r18    *int64
	Sl_PRS_RxInDedicatedResourcePool_r18 *struct {
		NumOfSupportedRxPSCCH_PerSlot_r18 int64
		SupportedCP_TypeFor60kHzSCS_r18   int64
	}
	Sl_PRS_TxInSharedResourcePool_r18           *int64
	Sl_PRS_TxScheme1InDedicatedResourcePool_r18 *int64
	Sl_PRS_TxScheme2InDedicatedResourcePool_r18 *int64
	Sl_PRS_CongestionCtrl_r18                   *int64
	Sl_PRS_TxRandomSelection_r18                *int64
	Sl_PRS_TxUsingFullSensing_r18               *int64
	Sl_PRS_RxForBandWithSL_CA_r18               *int64
	Sl_PRS_TxForBandWithSL_CA_r18               *int64
	Sl_DynamicSharingTxRx_r18                   *int64
	Sl_CA_Communication_r18                     *struct {
		NumberOfCarriers_r18           int64
		NumberOfPSCCH_DecodeValueZ_r18 int64
		TotalBandwidth_r18             int64
	}
	Sl_CA_Synchronization_r18 *int64
	Sl_CA_PSFCH_r18           *struct {
		Rx_PSFCH_Resource_r18 int64
		Tx_PSFCH_Resource_r18 int64
	}
	Sl_ReceptionIntraCarrierGuardBand_r18   *int64
	Sl_PathlossBasedOLPC_SL_RSRP_Report_r18 *int64
}

func (ie *BandSidelink_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandSidelinkR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Sl_TransmissionMode2_r16 != nil || ie.CongestionControlSidelink_r16 != nil || ie.FewerSymbolSlotSidelink_r16 != nil || ie.Sl_OpenLoopPC_RSRP_ReportSidelink_r16 != nil || ie.Sl_Rx_256QAM_r16 != nil
	hasExtGroup1 := ie.Ue_PowerClassSidelink_r16 != nil
	hasExtGroup2 := ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil || ie.Sync_Sidelink_v1710 != nil || ie.Enb_Sync_Sidelink_v1710 != nil || ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil || ie.Rx_IUC_Scheme1_SCI_r17 != nil || ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil
	hasExtGroup3 := ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18 != nil || ie.Sl_PRS_CommonProcCapabilityPerBand_r18 != nil || ie.Sl_PRS_RxInSharedResourcePool_r18 != nil || ie.Sl_PRS_RxInDedicatedResourcePool_r18 != nil || ie.Sl_PRS_TxInSharedResourcePool_r18 != nil || ie.Sl_PRS_TxScheme1InDedicatedResourcePool_r18 != nil || ie.Sl_PRS_TxScheme2InDedicatedResourcePool_r18 != nil || ie.Sl_PRS_CongestionCtrl_r18 != nil || ie.Sl_PRS_TxRandomSelection_r18 != nil || ie.Sl_PRS_TxUsingFullSensing_r18 != nil || ie.Sl_PRS_RxForBandWithSL_CA_r18 != nil || ie.Sl_PRS_TxForBandWithSL_CA_r18 != nil || ie.Sl_DynamicSharingTxRx_r18 != nil || ie.Sl_CA_Communication_r18 != nil || ie.Sl_CA_Synchronization_r18 != nil || ie.Sl_CA_PSFCH_r18 != nil || ie.Sl_ReceptionIntraCarrierGuardBand_r18 != nil
	hasExtGroup4 := ie.Sl_PathlossBasedOLPC_SL_RSRP_Report_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_Reception_r16 != nil, ie.Sl_TransmissionMode1_r16 != nil, ie.Sync_Sidelink_r16 != nil, ie.Sl_Tx_256QAM_r16 != nil, ie.Psfch_FormatZeroSidelink_r16 != nil, ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil, ie.Enb_Sync_Sidelink_r16 != nil}); err != nil {
		return err
	}

	// 3. freqBandSidelink-r16: ref
	{
		if err := ie.FreqBandSidelink_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. sl-Reception-r16: sequence
	{
		if ie.Sl_Reception_r16 != nil {
			c := ie.Sl_Reception_r16
			bandSidelinkR16SlReceptionR16Seq := e.NewSequenceEncoder(bandSidelinkR16SlReceptionR16Constraints)
			if err := bandSidelinkR16SlReceptionR16Seq.EncodePreamble([]bool{c.Scs_CP_PatternRxSidelink_r16 != nil, c.ExtendedCP_RxSidelink_r16 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Harq_RxProcessSidelink_r16, bandSidelinkR16SlReceptionR16HarqRxProcessSidelinkR16Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Pscch_RxSidelink_r16, bandSidelinkR16SlReceptionR16PscchRxSidelinkR16Constraints); err != nil {
				return err
			}
			if c.Scs_CP_PatternRxSidelink_r16 != nil {
				choiceEnc := e.NewChoiceEncoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Constraints)
				if err := choiceEnc.EncodeChoice(int64((*c.Scs_CP_PatternRxSidelink_r16).Choice), false, nil); err != nil {
					return err
				}
				switch (*c.Scs_CP_PatternRxSidelink_r16).Choice {
				case BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr1_r16:
					c := (*(*c.Scs_CP_PatternRxSidelink_r16).Fr1_r16)
					bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq := e.NewSequenceEncoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Constraints)
					if err := bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_15kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_30kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_30kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_60kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
				case BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr2_r16:
					c := (*(*c.Scs_CP_PatternRxSidelink_r16).Fr2_r16)
					bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq := e.NewSequenceEncoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Constraints)
					if err := bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq.EncodePreamble([]bool{c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_120kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_120kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
				}
			}
			if c.ExtendedCP_RxSidelink_r16 != nil {
				if err := e.EncodeEnumerated((*c.ExtendedCP_RxSidelink_r16), bandSidelinkR16SlReceptionR16ExtendedCPRxSidelinkR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 5. sl-TransmissionMode1-r16: sequence
	{
		if ie.Sl_TransmissionMode1_r16 != nil {
			c := ie.Sl_TransmissionMode1_r16
			bandSidelinkR16SlTransmissionMode1R16Seq := e.NewSequenceEncoder(bandSidelinkR16SlTransmissionMode1R16Constraints)
			if err := bandSidelinkR16SlTransmissionMode1R16Seq.EncodePreamble([]bool{c.ExtendedCP_TxSidelink_r16 != nil, c.Harq_ReportOnPUCCH_r16 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Harq_TxProcessModeOneSidelink_r16, bandSidelinkR16SlTransmissionMode1R16HarqTxProcessModeOneSidelinkR16Constraints); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.Scs_CP_PatternTxSidelinkModeOne_r16.Choice), false, nil); err != nil {
					return err
				}
				switch c.Scs_CP_PatternTxSidelinkModeOne_r16.Choice {
				case BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr1_r16:
					c := (*c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr1_r16)
					bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq := e.NewSequenceEncoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Constraints)
					if err := bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil}); err != nil {
						return err
					}
					if c.Scs_15kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_15kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_30kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_30kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_60kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
				case BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr2_r16:
					c := (*c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr2_r16)
					bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq := e.NewSequenceEncoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Constraints)
					if err := bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq.EncodePreamble([]bool{c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
						return err
					}
					if c.Scs_60kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_60kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
					if c.Scs_120kHz_r16 != nil {
						if err := e.EncodeBitString((*c.Scs_120kHz_r16), per.FixedSize(16)); err != nil {
							return err
						}
					}
				}
			}
			if c.ExtendedCP_TxSidelink_r16 != nil {
				if err := e.EncodeEnumerated((*c.ExtendedCP_TxSidelink_r16), bandSidelinkR16SlTransmissionMode1R16ExtendedCPTxSidelinkR16Constraints); err != nil {
					return err
				}
			}
			if c.Harq_ReportOnPUCCH_r16 != nil {
				if err := e.EncodeEnumerated((*c.Harq_ReportOnPUCCH_r16), bandSidelinkR16SlTransmissionMode1R16HarqReportOnPUCCHR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 6. sync-Sidelink-r16: sequence
	{
		if ie.Sync_Sidelink_r16 != nil {
			c := ie.Sync_Sidelink_r16
			bandSidelinkR16SyncSidelinkR16Seq := e.NewSequenceEncoder(bandSidelinkR16SyncSidelinkR16Constraints)
			if err := bandSidelinkR16SyncSidelinkR16Seq.EncodePreamble([]bool{c.GNB_Sync_r16 != nil, c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil, c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil}); err != nil {
				return err
			}
			if c.GNB_Sync_r16 != nil {
				if err := e.EncodeEnumerated((*c.GNB_Sync_r16), bandSidelinkR16SyncSidelinkR16GNBSyncR16Constraints); err != nil {
					return err
				}
			}
			if c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 != nil {
				if err := e.EncodeEnumerated((*c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16), bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNBENBR16Constraints); err != nil {
					return err
				}
			}
			if c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 != nil {
				if err := e.EncodeEnumerated((*c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16), bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNSSR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 7. sl-Tx-256QAM-r16: enumerated
	{
		if ie.Sl_Tx_256QAM_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_Tx_256QAM_r16, bandSidelinkR16SlTx256QAMR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. psfch-FormatZeroSidelink-r16: sequence
	{
		if ie.Psfch_FormatZeroSidelink_r16 != nil {
			c := ie.Psfch_FormatZeroSidelink_r16
			if err := e.EncodeEnumerated(c.Psfch_RxNumber, bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchRxNumberConstraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.Psfch_TxNumber, bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchTxNumberConstraints); err != nil {
				return err
			}
		}
	}

	// 9. lowSE-64QAM-MCS-TableSidelink-r16: enumerated
	{
		if ie.LowSE_64QAM_MCS_TableSidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.LowSE_64QAM_MCS_TableSidelink_r16, bandSidelinkR16LowSE64QAMMCSTableSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. enb-sync-Sidelink-r16: enumerated
	{
		if ie.Enb_Sync_Sidelink_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Enb_Sync_Sidelink_r16, bandSidelinkR16EnbSyncSidelinkR16Constraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sl-TransmissionMode2-r16", Optional: true},
					{Name: "congestionControlSidelink-r16", Optional: true},
					{Name: "fewerSymbolSlotSidelink-r16", Optional: true},
					{Name: "sl-openLoopPC-RSRP-ReportSidelink-r16", Optional: true},
					{Name: "sl-Rx-256QAM-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_TransmissionMode2_r16 != nil, ie.CongestionControlSidelink_r16 != nil, ie.FewerSymbolSlotSidelink_r16 != nil, ie.Sl_OpenLoopPC_RSRP_ReportSidelink_r16 != nil, ie.Sl_Rx_256QAM_r16 != nil}); err != nil {
				return err
			}

			if ie.Sl_TransmissionMode2_r16 != nil {
				c := ie.Sl_TransmissionMode2_r16
				bandSidelinkR16ExtSlTransmissionMode2R16Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtSlTransmissionMode2R16Constraints)
				if err := bandSidelinkR16ExtSlTransmissionMode2R16Seq.EncodePreamble([]bool{c.Scs_CP_PatternTxSidelinkModeTwo_r16 != nil, c.Dl_OpenLoopPC_Sidelink_r16 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Harq_TxProcessModeTwoSidelink_r16, bandSidelinkR16ExtSlTransmissionMode2R16HarqTxProcessModeTwoSidelinkR16Constraints); err != nil {
					return err
				}
				if c.Scs_CP_PatternTxSidelinkModeTwo_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Scs_CP_PatternTxSidelinkModeTwo_r16), bandSidelinkR16ExtSlTransmissionMode2R16ScsCPPatternTxSidelinkModeTwoR16Constraints); err != nil {
						return err
					}
				}
				if c.Dl_OpenLoopPC_Sidelink_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Dl_OpenLoopPC_Sidelink_r16), bandSidelinkR16ExtSlTransmissionMode2R16DlOpenLoopPCSidelinkR16Constraints); err != nil {
						return err
					}
				}
			}

			if ie.CongestionControlSidelink_r16 != nil {
				c := ie.CongestionControlSidelink_r16
				bandSidelinkR16ExtCongestionControlSidelinkR16Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtCongestionControlSidelinkR16Constraints)
				if err := bandSidelinkR16ExtCongestionControlSidelinkR16Seq.EncodePreamble([]bool{c.Cbr_ReportSidelink_r16 != nil}); err != nil {
					return err
				}
				if c.Cbr_ReportSidelink_r16 != nil {
					if err := ex.EncodeEnumerated((*c.Cbr_ReportSidelink_r16), bandSidelinkR16ExtCongestionControlSidelinkR16CbrReportSidelinkR16Constraints); err != nil {
						return err
					}
				}
				if err := ex.EncodeEnumerated(c.Cbr_CR_TimeLimitSidelink_r16, bandSidelinkR16ExtCongestionControlSidelinkR16CbrCRTimeLimitSidelinkR16Constraints); err != nil {
					return err
				}
			}

			if ie.FewerSymbolSlotSidelink_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.FewerSymbolSlotSidelink_r16, bandSidelinkR16ExtFewerSymbolSlotSidelinkR16Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_OpenLoopPC_RSRP_ReportSidelink_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_OpenLoopPC_RSRP_ReportSidelink_r16, bandSidelinkR16ExtSlOpenLoopPCRSRPReportSidelinkR16Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_Rx_256QAM_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_Rx_256QAM_r16, bandSidelinkR16ExtSlRx256QAMR16Constraints); err != nil {
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
					{Name: "ue-PowerClassSidelink-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ue_PowerClassSidelink_r16 != nil}); err != nil {
				return err
			}

			if ie.Ue_PowerClassSidelink_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Ue_PowerClassSidelink_r16, bandSidelinkR16ExtUePowerClassSidelinkR16Constraints); err != nil {
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
					{Name: "sl-TransmissionMode2-RandomResourceSelection-r17", Optional: true},
					{Name: "sync-Sidelink-v1710", Optional: true},
					{Name: "enb-sync-Sidelink-v1710", Optional: true},
					{Name: "rx-IUC-Scheme1-PreferredMode2Sidelink-r17", Optional: true},
					{Name: "rx-IUC-Scheme1-NonPreferredMode2Sidelink-r17", Optional: true},
					{Name: "rx-IUC-Scheme2-Mode2Sidelink-r17", Optional: true},
					{Name: "rx-IUC-Scheme1-SCI-r17", Optional: true},
					{Name: "rx-IUC-Scheme1-SCI-ExplicitReq-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil, ie.Sync_Sidelink_v1710 != nil, ie.Enb_Sync_Sidelink_v1710 != nil, ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil, ie.Rx_IUC_Scheme1_SCI_r17 != nil, ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil}); err != nil {
				return err
			}

			if ie.Sl_TransmissionMode2_RandomResourceSelection_r17 != nil {
				c := ie.Sl_TransmissionMode2_RandomResourceSelection_r17
				bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Constraints)
				if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq.EncodePreamble([]bool{c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil, c.ExtendedCP_Mode2Random_r17 != nil, c.Dl_OpenLoopPC_Sidelink_r17 != nil}); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Harq_TxProcessModeTwoSidelink_r17, bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17HarqTxProcessModeTwoSidelinkR17Constraints); err != nil {
					return err
				}
				if c.Scs_CP_PatternTxSidelinkModeTwo_r17 != nil {
					choiceEnc := ex.NewChoiceEncoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
					if err := choiceEnc.EncodeChoice(int64((*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice), false, nil); err != nil {
						return err
					}
					switch (*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice {
					case BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
						c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
						bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
						if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.EncodePreamble([]bool{c.Scs_15kHz_r17 != nil, c.Scs_30kHz_r17 != nil, c.Scs_60kHz_r17 != nil}); err != nil {
							return err
						}
						if c.Scs_15kHz_r17 != nil {
							if err := ex.EncodeBitString((*c.Scs_15kHz_r17), per.FixedSize(16)); err != nil {
								return err
							}
						}
						if c.Scs_30kHz_r17 != nil {
							if err := ex.EncodeBitString((*c.Scs_30kHz_r17), per.FixedSize(16)); err != nil {
								return err
							}
						}
						if c.Scs_60kHz_r17 != nil {
							if err := ex.EncodeBitString((*c.Scs_60kHz_r17), per.FixedSize(16)); err != nil {
								return err
							}
						}
					case BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
						c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
						bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
						if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.EncodePreamble([]bool{c.Scs_60kHz_r17 != nil, c.Scs_120kHz_r17 != nil}); err != nil {
							return err
						}
						if c.Scs_60kHz_r17 != nil {
							if err := ex.EncodeBitString((*c.Scs_60kHz_r17), per.FixedSize(16)); err != nil {
								return err
							}
						}
						if c.Scs_120kHz_r17 != nil {
							if err := ex.EncodeBitString((*c.Scs_120kHz_r17), per.FixedSize(16)); err != nil {
								return err
							}
						}
					}
				}
				if c.ExtendedCP_Mode2Random_r17 != nil {
					if err := ex.EncodeEnumerated((*c.ExtendedCP_Mode2Random_r17), bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ExtendedCPMode2RandomR17Constraints); err != nil {
						return err
					}
				}
				if c.Dl_OpenLoopPC_Sidelink_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Dl_OpenLoopPC_Sidelink_r17), bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17DlOpenLoopPCSidelinkR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Sync_Sidelink_v1710 != nil {
				c := ie.Sync_Sidelink_v1710
				bandSidelinkR16ExtSyncSidelinkV1710Seq := ex.NewSequenceEncoder(bandSidelinkR16ExtSyncSidelinkV1710Constraints)
				if err := bandSidelinkR16ExtSyncSidelinkV1710Seq.EncodePreamble([]bool{c.Sync_GNSS_r17 != nil, c.GNB_Sync_r17 != nil, c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil, c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil}); err != nil {
					return err
				}
				if c.Sync_GNSS_r17 != nil {
					if err := ex.EncodeEnumerated((*c.Sync_GNSS_r17), bandSidelinkR16ExtSyncSidelinkV1710SyncGNSSR17Constraints); err != nil {
						return err
					}
				}
				if c.GNB_Sync_r17 != nil {
					if err := ex.EncodeEnumerated((*c.GNB_Sync_r17), bandSidelinkR16ExtSyncSidelinkV1710GNBSyncR17Constraints); err != nil {
						return err
					}
				}
				if c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 != nil {
					if err := ex.EncodeEnumerated((*c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17), bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNBENBR17Constraints); err != nil {
						return err
					}
				}
				if c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 != nil {
					if err := ex.EncodeEnumerated((*c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17), bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNSSR17Constraints); err != nil {
						return err
					}
				}
			}

			if ie.Enb_Sync_Sidelink_v1710 != nil {
				if err := ex.EncodeEnumerated(*ie.Enb_Sync_Sidelink_v1710, bandSidelinkR16ExtEnbSyncSidelinkV1710Constraints); err != nil {
					return err
				}
			}

			if ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17, bandSidelinkR16ExtRxIUCScheme1PreferredMode2SidelinkR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17, bandSidelinkR16ExtRxIUCScheme1NonPreferredMode2SidelinkR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rx_IUC_Scheme2_Mode2Sidelink_r17, bandSidelinkR16ExtRxIUCScheme2Mode2SidelinkR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rx_IUC_Scheme1_SCI_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rx_IUC_Scheme1_SCI_r17, bandSidelinkR16ExtRxIUCScheme1SCIR17Constraints); err != nil {
					return err
				}
			}

			if ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17, bandSidelinkR16ExtRxIUCScheme1SCIExplicitReqR17Constraints); err != nil {
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
					{Name: "sharedSpectrumChAccessParamsSidelinkPerBand-r18", Optional: true},
					{Name: "sl-PRS-CommonProcCapabilityPerBand-r18", Optional: true},
					{Name: "sl-PRS-RxInSharedResourcePool-r18", Optional: true},
					{Name: "sl-PRS-RxInDedicatedResourcePool-r18", Optional: true},
					{Name: "sl-PRS-TxInSharedResourcePool-r18", Optional: true},
					{Name: "sl-PRS-TxScheme1InDedicatedResourcePool-r18", Optional: true},
					{Name: "sl-PRS-TxScheme2InDedicatedResourcePool-r18", Optional: true},
					{Name: "sl-PRS-CongestionCtrl-r18", Optional: true},
					{Name: "sl-PRS-TxRandomSelection-r18", Optional: true},
					{Name: "sl-PRS-TxUsingFullSensing-r18", Optional: true},
					{Name: "sl-PRS-RxForBandWithSL-CA-r18", Optional: true},
					{Name: "sl-PRS-TxForBandWithSL-CA-r18", Optional: true},
					{Name: "sl-DynamicSharingTxRx-r18", Optional: true},
					{Name: "sl-CA-Communication-r18", Optional: true},
					{Name: "sl-CA-Synchronization-r18", Optional: true},
					{Name: "sl-CA-PSFCH-r18", Optional: true},
					{Name: "sl-ReceptionIntraCarrierGuardBand-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18 != nil, ie.Sl_PRS_CommonProcCapabilityPerBand_r18 != nil, ie.Sl_PRS_RxInSharedResourcePool_r18 != nil, ie.Sl_PRS_RxInDedicatedResourcePool_r18 != nil, ie.Sl_PRS_TxInSharedResourcePool_r18 != nil, ie.Sl_PRS_TxScheme1InDedicatedResourcePool_r18 != nil, ie.Sl_PRS_TxScheme2InDedicatedResourcePool_r18 != nil, ie.Sl_PRS_CongestionCtrl_r18 != nil, ie.Sl_PRS_TxRandomSelection_r18 != nil, ie.Sl_PRS_TxUsingFullSensing_r18 != nil, ie.Sl_PRS_RxForBandWithSL_CA_r18 != nil, ie.Sl_PRS_TxForBandWithSL_CA_r18 != nil, ie.Sl_DynamicSharingTxRx_r18 != nil, ie.Sl_CA_Communication_r18 != nil, ie.Sl_CA_Synchronization_r18 != nil, ie.Sl_CA_PSFCH_r18 != nil, ie.Sl_ReceptionIntraCarrierGuardBand_r18 != nil}); err != nil {
				return err
			}

			if ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18 != nil {
				if err := ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18.Encode(ex); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_CommonProcCapabilityPerBand_r18 != nil {
				c := ie.Sl_PRS_CommonProcCapabilityPerBand_r18
				{
					choiceEnc := ex.NewChoiceEncoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MaxSL_PRS_Bandwidth_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.MaxSL_PRS_Bandwidth_r18.Choice {
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18:
						if err := ex.EncodeEnumerated((*c.MaxSL_PRS_Bandwidth_r18.Fr1_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr1R18Constraints); err != nil {
							return err
						}
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18:
						if err := ex.EncodeEnumerated((*c.MaxSL_PRS_Bandwidth_r18.Fr2_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr2R18Constraints); err != nil {
							return err
						}
					}
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Choice {
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18:
						if err := ex.EncodeEnumerated((*c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr1_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr1R18Constraints); err != nil {
							return err
						}
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18:
						if err := ex.EncodeEnumerated((*c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr2_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr2R18Constraints); err != nil {
							return err
						}
					}
				}
				{
					choiceEnc := ex.NewChoiceEncoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Constraints)
					if err := choiceEnc.EncodeChoice(int64(c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Choice), false, nil); err != nil {
						return err
					}
					switch c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Choice {
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18:
						if err := ex.EncodeEnumerated((*c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr1_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr1R18Constraints); err != nil {
							return err
						}
					case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18:
						if err := ex.EncodeEnumerated((*c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr2_r18), bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr2R18Constraints); err != nil {
							return err
						}
					}
				}
				if err := ex.EncodeEnumerated(c.MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18, bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MinTimeAfterEndofSlotCarryActiveSLPRSResourcesR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_RxInSharedResourcePool_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_RxInSharedResourcePool_r18, bandSidelinkR16ExtSlPRSRxInSharedResourcePoolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_RxInDedicatedResourcePool_r18 != nil {
				c := ie.Sl_PRS_RxInDedicatedResourcePool_r18
				if err := ex.EncodeEnumerated(c.NumOfSupportedRxPSCCH_PerSlot_r18, bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18NumOfSupportedRxPSCCHPerSlotR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.SupportedCP_TypeFor60kHzSCS_r18, bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18SupportedCPTypeFor60kHzSCSR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxInSharedResourcePool_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxInSharedResourcePool_r18, bandSidelinkR16ExtSlPRSTxInSharedResourcePoolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxScheme1InDedicatedResourcePool_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxScheme1InDedicatedResourcePool_r18, bandSidelinkR16ExtSlPRSTxScheme1InDedicatedResourcePoolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxScheme2InDedicatedResourcePool_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxScheme2InDedicatedResourcePool_r18, bandSidelinkR16ExtSlPRSTxScheme2InDedicatedResourcePoolR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_CongestionCtrl_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_CongestionCtrl_r18, bandSidelinkR16ExtSlPRSCongestionCtrlR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxRandomSelection_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxRandomSelection_r18, bandSidelinkR16ExtSlPRSTxRandomSelectionR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxUsingFullSensing_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxUsingFullSensing_r18, bandSidelinkR16ExtSlPRSTxUsingFullSensingR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_RxForBandWithSL_CA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_RxForBandWithSL_CA_r18, bandSidelinkR16ExtSlPRSRxForBandWithSLCAR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_PRS_TxForBandWithSL_CA_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PRS_TxForBandWithSL_CA_r18, bandSidelinkR16ExtSlPRSTxForBandWithSLCAR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_DynamicSharingTxRx_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_DynamicSharingTxRx_r18, bandSidelinkR16ExtSlDynamicSharingTxRxR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_CA_Communication_r18 != nil {
				c := ie.Sl_CA_Communication_r18
				if err := ex.EncodeInteger(c.NumberOfCarriers_r18, per.Constrained(2, 8)); err != nil {
					return err
				}
				if err := ex.EncodeInteger(c.NumberOfPSCCH_DecodeValueZ_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.TotalBandwidth_r18, bandSidelinkR16ExtSlCACommunicationR18TotalBandwidthR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_CA_Synchronization_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_CA_Synchronization_r18, bandSidelinkR16ExtSlCASynchronizationR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_CA_PSFCH_r18 != nil {
				c := ie.Sl_CA_PSFCH_r18
				if err := ex.EncodeEnumerated(c.Rx_PSFCH_Resource_r18, bandSidelinkR16ExtSlCAPSFCHR18RxPSFCHResourceR18Constraints); err != nil {
					return err
				}
				if err := ex.EncodeEnumerated(c.Tx_PSFCH_Resource_r18, bandSidelinkR16ExtSlCAPSFCHR18TxPSFCHResourceR18Constraints); err != nil {
					return err
				}
			}

			if ie.Sl_ReceptionIntraCarrierGuardBand_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_ReceptionIntraCarrierGuardBand_r18, bandSidelinkR16ExtSlReceptionIntraCarrierGuardBandR18Constraints); err != nil {
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
					{Name: "sl-PathlossBasedOLPC-SL-RSRP-Report-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sl_PathlossBasedOLPC_SL_RSRP_Report_r18 != nil}); err != nil {
				return err
			}

			if ie.Sl_PathlossBasedOLPC_SL_RSRP_Report_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Sl_PathlossBasedOLPC_SL_RSRP_Report_r18, bandSidelinkR16ExtSlPathlossBasedOLPCSLRSRPReportR18Constraints); err != nil {
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

func (ie *BandSidelink_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandSidelinkR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. freqBandSidelink-r16: ref
	{
		if err := ie.FreqBandSidelink_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. sl-Reception-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_Reception_r16 = &struct {
				Harq_RxProcessSidelink_r16   int64
				Pscch_RxSidelink_r16         int64
				Scs_CP_PatternRxSidelink_r16 *struct {
					Choice  int
					Fr1_r16 *struct {
						Scs_15kHz_r16 *per.BitString
						Scs_30kHz_r16 *per.BitString
						Scs_60kHz_r16 *per.BitString
					}
					Fr2_r16 *struct {
						Scs_60kHz_r16  *per.BitString
						Scs_120kHz_r16 *per.BitString
					}
				}
				ExtendedCP_RxSidelink_r16 *int64
			}{}
			c := ie.Sl_Reception_r16
			bandSidelinkR16SlReceptionR16Seq := d.NewSequenceDecoder(bandSidelinkR16SlReceptionR16Constraints)
			if err := bandSidelinkR16SlReceptionR16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(bandSidelinkR16SlReceptionR16HarqRxProcessSidelinkR16Constraints)
				if err != nil {
					return err
				}
				c.Harq_RxProcessSidelink_r16 = v
			}
			{
				v, err := d.DecodeEnumerated(bandSidelinkR16SlReceptionR16PscchRxSidelinkR16Constraints)
				if err != nil {
					return err
				}
				c.Pscch_RxSidelink_r16 = v
			}
			if bandSidelinkR16SlReceptionR16Seq.IsComponentPresent(2) {
				c.Scs_CP_PatternRxSidelink_r16 = &struct {
					Choice  int
					Fr1_r16 *struct {
						Scs_15kHz_r16 *per.BitString
						Scs_30kHz_r16 *per.BitString
						Scs_60kHz_r16 *per.BitString
					}
					Fr2_r16 *struct {
						Scs_60kHz_r16  *per.BitString
						Scs_120kHz_r16 *per.BitString
					}
				}{}
				choiceDec := d.NewChoiceDecoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Scs_CP_PatternRxSidelink_r16).Choice = int(index)
				switch index {
				case BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr1_r16:
					(*c.Scs_CP_PatternRxSidelink_r16).Fr1_r16 = &struct {
						Scs_15kHz_r16 *per.BitString
						Scs_30kHz_r16 *per.BitString
						Scs_60kHz_r16 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternRxSidelink_r16).Fr1_r16)
					bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq := d.NewSequenceDecoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Constraints)
					if err := bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r16) = v
					}
					if bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r16) = v
					}
					if bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr1R16Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r16) = v
					}
				case BandSidelink_r16_Sl_Reception_r16_Scs_CP_PatternRxSidelink_r16_Fr2_r16:
					(*c.Scs_CP_PatternRxSidelink_r16).Fr2_r16 = &struct {
						Scs_60kHz_r16  *per.BitString
						Scs_120kHz_r16 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternRxSidelink_r16).Fr2_r16)
					bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq := d.NewSequenceDecoder(bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Constraints)
					if err := bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r16) = v
					}
					if bandSidelinkR16SlReceptionR16ScsCPPatternRxSidelinkR16Fr2R16Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r16) = v
					}
				}
			}
			if bandSidelinkR16SlReceptionR16Seq.IsComponentPresent(3) {
				c.ExtendedCP_RxSidelink_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SlReceptionR16ExtendedCPRxSidelinkR16Constraints)
				if err != nil {
					return err
				}
				(*c.ExtendedCP_RxSidelink_r16) = v
			}
		}
	}

	// 5. sl-TransmissionMode1-r16: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_TransmissionMode1_r16 = &struct {
				Harq_TxProcessModeOneSidelink_r16   int64
				Scs_CP_PatternTxSidelinkModeOne_r16 struct {
					Choice  int
					Fr1_r16 *struct {
						Scs_15kHz_r16 *per.BitString
						Scs_30kHz_r16 *per.BitString
						Scs_60kHz_r16 *per.BitString
					}
					Fr2_r16 *struct {
						Scs_60kHz_r16  *per.BitString
						Scs_120kHz_r16 *per.BitString
					}
				}
				ExtendedCP_TxSidelink_r16 *int64
				Harq_ReportOnPUCCH_r16    *int64
			}{}
			c := ie.Sl_TransmissionMode1_r16
			bandSidelinkR16SlTransmissionMode1R16Seq := d.NewSequenceDecoder(bandSidelinkR16SlTransmissionMode1R16Constraints)
			if err := bandSidelinkR16SlTransmissionMode1R16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(bandSidelinkR16SlTransmissionMode1R16HarqTxProcessModeOneSidelinkR16Constraints)
				if err != nil {
					return err
				}
				c.Harq_TxProcessModeOneSidelink_r16 = v
			}
			{
				choiceDec := d.NewChoiceDecoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.Scs_CP_PatternTxSidelinkModeOne_r16.Choice = int(index)
				switch index {
				case BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr1_r16:
					c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr1_r16 = &struct {
						Scs_15kHz_r16 *per.BitString
						Scs_30kHz_r16 *per.BitString
						Scs_60kHz_r16 *per.BitString
					}{}
					c := (*c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr1_r16)
					bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq := d.NewSequenceDecoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Constraints)
					if err := bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r16) = v
					}
					if bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r16) = v
					}
					if bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr1R16Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r16) = v
					}
				case BandSidelink_r16_Sl_TransmissionMode1_r16_Scs_CP_PatternTxSidelinkModeOne_r16_Fr2_r16:
					c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr2_r16 = &struct {
						Scs_60kHz_r16  *per.BitString
						Scs_120kHz_r16 *per.BitString
					}{}
					c := (*c.Scs_CP_PatternTxSidelinkModeOne_r16.Fr2_r16)
					bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq := d.NewSequenceDecoder(bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Constraints)
					if err := bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r16) = v
					}
					if bandSidelinkR16SlTransmissionMode1R16ScsCPPatternTxSidelinkModeOneR16Fr2R16Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r16 = new(per.BitString)
						v, err := d.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r16) = v
					}
				}
			}
			if bandSidelinkR16SlTransmissionMode1R16Seq.IsComponentPresent(2) {
				c.ExtendedCP_TxSidelink_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SlTransmissionMode1R16ExtendedCPTxSidelinkR16Constraints)
				if err != nil {
					return err
				}
				(*c.ExtendedCP_TxSidelink_r16) = v
			}
			if bandSidelinkR16SlTransmissionMode1R16Seq.IsComponentPresent(3) {
				c.Harq_ReportOnPUCCH_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SlTransmissionMode1R16HarqReportOnPUCCHR16Constraints)
				if err != nil {
					return err
				}
				(*c.Harq_ReportOnPUCCH_r16) = v
			}
		}
	}

	// 6. sync-Sidelink-r16: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Sync_Sidelink_r16 = &struct {
				GNB_Sync_r16                              *int64
				GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 *int64
				GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16    *int64
			}{}
			c := ie.Sync_Sidelink_r16
			bandSidelinkR16SyncSidelinkR16Seq := d.NewSequenceDecoder(bandSidelinkR16SyncSidelinkR16Constraints)
			if err := bandSidelinkR16SyncSidelinkR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandSidelinkR16SyncSidelinkR16Seq.IsComponentPresent(0) {
				c.GNB_Sync_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SyncSidelinkR16GNBSyncR16Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_Sync_r16) = v
			}
			if bandSidelinkR16SyncSidelinkR16Seq.IsComponentPresent(1) {
				c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNBENBR16Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r16) = v
			}
			if bandSidelinkR16SyncSidelinkR16Seq.IsComponentPresent(2) {
				c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16 = new(int64)
				v, err := d.DecodeEnumerated(bandSidelinkR16SyncSidelinkR16GNBGNSSUESyncWithPriorityOnGNSSR16Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r16) = v
			}
		}
	}

	// 7. sl-Tx-256QAM-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(bandSidelinkR16SlTx256QAMR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Tx_256QAM_r16 = &idx
		}
	}

	// 8. psfch-FormatZeroSidelink-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Psfch_FormatZeroSidelink_r16 = &struct {
				Psfch_RxNumber int64
				Psfch_TxNumber int64
			}{}
			c := ie.Psfch_FormatZeroSidelink_r16
			{
				v, err := d.DecodeEnumerated(bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchRxNumberConstraints)
				if err != nil {
					return err
				}
				c.Psfch_RxNumber = v
			}
			{
				v, err := d.DecodeEnumerated(bandSidelinkR16PsfchFormatZeroSidelinkR16PsfchTxNumberConstraints)
				if err != nil {
					return err
				}
				c.Psfch_TxNumber = v
			}
		}
	}

	// 9. lowSE-64QAM-MCS-TableSidelink-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(bandSidelinkR16LowSE64QAMMCSTableSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.LowSE_64QAM_MCS_TableSidelink_r16 = &idx
		}
	}

	// 10. enb-sync-Sidelink-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(bandSidelinkR16EnbSyncSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Enb_Sync_Sidelink_r16 = &idx
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
				{Name: "sl-TransmissionMode2-r16", Optional: true},
				{Name: "congestionControlSidelink-r16", Optional: true},
				{Name: "fewerSymbolSlotSidelink-r16", Optional: true},
				{Name: "sl-openLoopPC-RSRP-ReportSidelink-r16", Optional: true},
				{Name: "sl-Rx-256QAM-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_TransmissionMode2_r16 = &struct {
				Harq_TxProcessModeTwoSidelink_r16   int64
				Scs_CP_PatternTxSidelinkModeTwo_r16 *int64
				Dl_OpenLoopPC_Sidelink_r16          *int64
			}{}
			c := ie.Sl_TransmissionMode2_r16
			bandSidelinkR16ExtSlTransmissionMode2R16Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtSlTransmissionMode2R16Constraints)
			if err := bandSidelinkR16ExtSlTransmissionMode2R16Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2R16HarqTxProcessModeTwoSidelinkR16Constraints)
				if err != nil {
					return err
				}
				c.Harq_TxProcessModeTwoSidelink_r16 = v
			}
			if bandSidelinkR16ExtSlTransmissionMode2R16Seq.IsComponentPresent(1) {
				c.Scs_CP_PatternTxSidelinkModeTwo_r16 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2R16ScsCPPatternTxSidelinkModeTwoR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_CP_PatternTxSidelinkModeTwo_r16) = v
			}
			if bandSidelinkR16ExtSlTransmissionMode2R16Seq.IsComponentPresent(2) {
				c.Dl_OpenLoopPC_Sidelink_r16 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2R16DlOpenLoopPCSidelinkR16Constraints)
				if err != nil {
					return err
				}
				(*c.Dl_OpenLoopPC_Sidelink_r16) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.CongestionControlSidelink_r16 = &struct {
				Cbr_ReportSidelink_r16       *int64
				Cbr_CR_TimeLimitSidelink_r16 int64
			}{}
			c := ie.CongestionControlSidelink_r16
			bandSidelinkR16ExtCongestionControlSidelinkR16Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtCongestionControlSidelinkR16Constraints)
			if err := bandSidelinkR16ExtCongestionControlSidelinkR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandSidelinkR16ExtCongestionControlSidelinkR16Seq.IsComponentPresent(0) {
				c.Cbr_ReportSidelink_r16 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtCongestionControlSidelinkR16CbrReportSidelinkR16Constraints)
				if err != nil {
					return err
				}
				(*c.Cbr_ReportSidelink_r16) = v
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtCongestionControlSidelinkR16CbrCRTimeLimitSidelinkR16Constraints)
				if err != nil {
					return err
				}
				c.Cbr_CR_TimeLimitSidelink_r16 = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtFewerSymbolSlotSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.FewerSymbolSlotSidelink_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlOpenLoopPCRSRPReportSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_OpenLoopPC_RSRP_ReportSidelink_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlRx256QAMR16Constraints)
			if err != nil {
				return err
			}
			ie.Sl_Rx_256QAM_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "ue-PowerClassSidelink-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtUePowerClassSidelinkR16Constraints)
			if err != nil {
				return err
			}
			ie.Ue_PowerClassSidelink_r16 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-TransmissionMode2-RandomResourceSelection-r17", Optional: true},
				{Name: "sync-Sidelink-v1710", Optional: true},
				{Name: "enb-sync-Sidelink-v1710", Optional: true},
				{Name: "rx-IUC-Scheme1-PreferredMode2Sidelink-r17", Optional: true},
				{Name: "rx-IUC-Scheme1-NonPreferredMode2Sidelink-r17", Optional: true},
				{Name: "rx-IUC-Scheme2-Mode2Sidelink-r17", Optional: true},
				{Name: "rx-IUC-Scheme1-SCI-r17", Optional: true},
				{Name: "rx-IUC-Scheme1-SCI-ExplicitReq-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Sl_TransmissionMode2_RandomResourceSelection_r17 = &struct {
				Harq_TxProcessModeTwoSidelink_r17   int64
				Scs_CP_PatternTxSidelinkModeTwo_r17 *struct {
					Choice  int
					Fr1_r17 *struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}
					Fr2_r17 *struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}
				}
				ExtendedCP_Mode2Random_r17 *int64
				Dl_OpenLoopPC_Sidelink_r17 *int64
			}{}
			c := ie.Sl_TransmissionMode2_RandomResourceSelection_r17
			bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Constraints)
			if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17HarqTxProcessModeTwoSidelinkR17Constraints)
				if err != nil {
					return err
				}
				c.Harq_TxProcessModeTwoSidelink_r17 = v
			}
			if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq.IsComponentPresent(1) {
				c.Scs_CP_PatternTxSidelinkModeTwo_r17 = &struct {
					Choice  int
					Fr1_r17 *struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}
					Fr2_r17 *struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}
				}{}
				choiceDec := dx.NewChoiceDecoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Choice = int(index)
				switch index {
				case BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr1_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17 = &struct {
						Scs_15kHz_r17 *per.BitString
						Scs_30kHz_r17 *per.BitString
						Scs_60kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr1_r17)
					bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Constraints)
					if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(0) {
						c.Scs_15kHz_r17 = new(per.BitString)
						v, err := dx.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_15kHz_r17) = v
					}
					if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(1) {
						c.Scs_30kHz_r17 = new(per.BitString)
						v, err := dx.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_30kHz_r17) = v
					}
					if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr1R17Seq.IsComponentPresent(2) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := dx.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
				case BandSidelink_r16_Ext_Sl_TransmissionMode2_RandomResourceSelection_r17_Scs_CP_PatternTxSidelinkModeTwo_r17_Fr2_r17:
					(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17 = &struct {
						Scs_60kHz_r17  *per.BitString
						Scs_120kHz_r17 *per.BitString
					}{}
					c := (*(*c.Scs_CP_PatternTxSidelinkModeTwo_r17).Fr2_r17)
					bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Constraints)
					if err := bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.DecodePreamble(); err != nil {
						return err
					}
					if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(0) {
						c.Scs_60kHz_r17 = new(per.BitString)
						v, err := dx.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_60kHz_r17) = v
					}
					if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ScsCPPatternTxSidelinkModeTwoR17Fr2R17Seq.IsComponentPresent(1) {
						c.Scs_120kHz_r17 = new(per.BitString)
						v, err := dx.DecodeBitString(per.FixedSize(16))
						if err != nil {
							return err
						}
						(*c.Scs_120kHz_r17) = v
					}
				}
			}
			if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq.IsComponentPresent(2) {
				c.ExtendedCP_Mode2Random_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17ExtendedCPMode2RandomR17Constraints)
				if err != nil {
					return err
				}
				(*c.ExtendedCP_Mode2Random_r17) = v
			}
			if bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17Seq.IsComponentPresent(3) {
				c.Dl_OpenLoopPC_Sidelink_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlTransmissionMode2RandomResourceSelectionR17DlOpenLoopPCSidelinkR17Constraints)
				if err != nil {
					return err
				}
				(*c.Dl_OpenLoopPC_Sidelink_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sync_Sidelink_v1710 = &struct {
				Sync_GNSS_r17                             *int64
				GNB_Sync_r17                              *int64
				GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 *int64
				GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17    *int64
			}{}
			c := ie.Sync_Sidelink_v1710
			bandSidelinkR16ExtSyncSidelinkV1710Seq := dx.NewSequenceDecoder(bandSidelinkR16ExtSyncSidelinkV1710Constraints)
			if err := bandSidelinkR16ExtSyncSidelinkV1710Seq.DecodePreamble(); err != nil {
				return err
			}
			if bandSidelinkR16ExtSyncSidelinkV1710Seq.IsComponentPresent(0) {
				c.Sync_GNSS_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSyncSidelinkV1710SyncGNSSR17Constraints)
				if err != nil {
					return err
				}
				(*c.Sync_GNSS_r17) = v
			}
			if bandSidelinkR16ExtSyncSidelinkV1710Seq.IsComponentPresent(1) {
				c.GNB_Sync_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSyncSidelinkV1710GNBSyncR17Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_Sync_r17) = v
			}
			if bandSidelinkR16ExtSyncSidelinkV1710Seq.IsComponentPresent(2) {
				c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNBENBR17Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_GNSS_UE_SyncWithPriorityOnGNB_ENB_r17) = v
			}
			if bandSidelinkR16ExtSyncSidelinkV1710Seq.IsComponentPresent(3) {
				c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17 = new(int64)
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSyncSidelinkV1710GNBGNSSUESyncWithPriorityOnGNSSR17Constraints)
				if err != nil {
					return err
				}
				(*c.GNB_GNSS_UE_SyncWithPriorityOnGNSS_r17) = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtEnbSyncSidelinkV1710Constraints)
			if err != nil {
				return err
			}
			ie.Enb_Sync_Sidelink_v1710 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtRxIUCScheme1PreferredMode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Rx_IUC_Scheme1_PreferredMode2Sidelink_r17 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtRxIUCScheme1NonPreferredMode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Rx_IUC_Scheme1_NonPreferredMode2Sidelink_r17 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtRxIUCScheme2Mode2SidelinkR17Constraints)
			if err != nil {
				return err
			}
			ie.Rx_IUC_Scheme2_Mode2Sidelink_r17 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtRxIUCScheme1SCIR17Constraints)
			if err != nil {
				return err
			}
			ie.Rx_IUC_Scheme1_SCI_r17 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtRxIUCScheme1SCIExplicitReqR17Constraints)
			if err != nil {
				return err
			}
			ie.Rx_IUC_Scheme1_SCI_ExplicitReq_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sharedSpectrumChAccessParamsSidelinkPerBand-r18", Optional: true},
				{Name: "sl-PRS-CommonProcCapabilityPerBand-r18", Optional: true},
				{Name: "sl-PRS-RxInSharedResourcePool-r18", Optional: true},
				{Name: "sl-PRS-RxInDedicatedResourcePool-r18", Optional: true},
				{Name: "sl-PRS-TxInSharedResourcePool-r18", Optional: true},
				{Name: "sl-PRS-TxScheme1InDedicatedResourcePool-r18", Optional: true},
				{Name: "sl-PRS-TxScheme2InDedicatedResourcePool-r18", Optional: true},
				{Name: "sl-PRS-CongestionCtrl-r18", Optional: true},
				{Name: "sl-PRS-TxRandomSelection-r18", Optional: true},
				{Name: "sl-PRS-TxUsingFullSensing-r18", Optional: true},
				{Name: "sl-PRS-RxForBandWithSL-CA-r18", Optional: true},
				{Name: "sl-PRS-TxForBandWithSL-CA-r18", Optional: true},
				{Name: "sl-DynamicSharingTxRx-r18", Optional: true},
				{Name: "sl-CA-Communication-r18", Optional: true},
				{Name: "sl-CA-Synchronization-r18", Optional: true},
				{Name: "sl-CA-PSFCH-r18", Optional: true},
				{Name: "sl-ReceptionIntraCarrierGuardBand-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18 = new(SharedSpectrumChAccessParamsSidelinkPerBand_r18)
			if err := ie.SharedSpectrumChAccessParamsSidelinkPerBand_r18.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Sl_PRS_CommonProcCapabilityPerBand_r18 = &struct {
				MaxSL_PRS_Bandwidth_r18 struct {
					Choice  int
					Fr1_r18 *int64
					Fr2_r18 *int64
				}
				MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18 struct {
					Choice  int
					Fr1_r18 *int64
					Fr2_r18 *int64
				}
				MaxNumOfSlotsWithActiveSL_PRS_Resources_r18 struct {
					Choice  int
					Fr1_r18 *int64
					Fr2_r18 *int64
				}
				MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18 int64
			}{}
			c := ie.Sl_PRS_CommonProcCapabilityPerBand_r18
			{
				choiceDec := dx.NewChoiceDecoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MaxSL_PRS_Bandwidth_r18.Choice = int(index)
				switch index {
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr1_r18:
					c.MaxSL_PRS_Bandwidth_r18.Fr1_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr1R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxSL_PRS_Bandwidth_r18.Fr1_r18) = v
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxSL_PRS_Bandwidth_r18_Fr2_r18:
					c.MaxSL_PRS_Bandwidth_r18.Fr2_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxSLPRSBandwidthR18Fr2R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxSL_PRS_Bandwidth_r18.Fr2_r18) = v
				}
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Choice = int(index)
				switch index {
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr1_r18:
					c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr1_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr1R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr1_r18) = v
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18_Fr2_r18:
					c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr2_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfActiveSLPRSResourcesInOneSlotR18Fr2R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxNumOfActiveSL_PRS_ResourcesInOneSlot_r18.Fr2_r18) = v
				}
			}
			{
				choiceDec := dx.NewChoiceDecoder(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Choice = int(index)
				switch index {
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr1_r18:
					c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr1_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr1R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr1_r18) = v
				case BandSidelink_r16_Ext_Sl_PRS_CommonProcCapabilityPerBand_r18_MaxNumOfSlotsWithActiveSL_PRS_Resources_r18_Fr2_r18:
					c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr2_r18 = new(int64)
					v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MaxNumOfSlotsWithActiveSLPRSResourcesR18Fr2R18Constraints)
					if err != nil {
						return err
					}
					(*c.MaxNumOfSlotsWithActiveSL_PRS_Resources_r18.Fr2_r18) = v
				}
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCommonProcCapabilityPerBandR18MinTimeAfterEndofSlotCarryActiveSLPRSResourcesR18Constraints)
				if err != nil {
					return err
				}
				c.MinTimeAfterEndofSlotCarryActiveSL_PRS_Resources_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSRxInSharedResourcePoolR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_RxInSharedResourcePool_r18 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Sl_PRS_RxInDedicatedResourcePool_r18 = &struct {
				NumOfSupportedRxPSCCH_PerSlot_r18 int64
				SupportedCP_TypeFor60kHzSCS_r18   int64
			}{}
			c := ie.Sl_PRS_RxInDedicatedResourcePool_r18
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18NumOfSupportedRxPSCCHPerSlotR18Constraints)
				if err != nil {
					return err
				}
				c.NumOfSupportedRxPSCCH_PerSlot_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSRxInDedicatedResourcePoolR18SupportedCPTypeFor60kHzSCSR18Constraints)
				if err != nil {
					return err
				}
				c.SupportedCP_TypeFor60kHzSCS_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxInSharedResourcePoolR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxInSharedResourcePool_r18 = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxScheme1InDedicatedResourcePoolR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxScheme1InDedicatedResourcePool_r18 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxScheme2InDedicatedResourcePoolR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxScheme2InDedicatedResourcePool_r18 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSCongestionCtrlR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_CongestionCtrl_r18 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxRandomSelectionR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxRandomSelection_r18 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxUsingFullSensingR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxUsingFullSensing_r18 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSRxForBandWithSLCAR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_RxForBandWithSL_CA_r18 = &v
		}

		if groupSeq.IsComponentPresent(11) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPRSTxForBandWithSLCAR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_TxForBandWithSL_CA_r18 = &v
		}

		if groupSeq.IsComponentPresent(12) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlDynamicSharingTxRxR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_DynamicSharingTxRx_r18 = &v
		}

		if groupSeq.IsComponentPresent(13) {
			ie.Sl_CA_Communication_r18 = &struct {
				NumberOfCarriers_r18           int64
				NumberOfPSCCH_DecodeValueZ_r18 int64
				TotalBandwidth_r18             int64
			}{}
			c := ie.Sl_CA_Communication_r18
			{
				v, err := dx.DecodeInteger(per.Constrained(2, 8))
				if err != nil {
					return err
				}
				c.NumberOfCarriers_r18 = v
			}
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.NumberOfPSCCH_DecodeValueZ_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlCACommunicationR18TotalBandwidthR18Constraints)
				if err != nil {
					return err
				}
				c.TotalBandwidth_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(14) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlCASynchronizationR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CA_Synchronization_r18 = &v
		}

		if groupSeq.IsComponentPresent(15) {
			ie.Sl_CA_PSFCH_r18 = &struct {
				Rx_PSFCH_Resource_r18 int64
				Tx_PSFCH_Resource_r18 int64
			}{}
			c := ie.Sl_CA_PSFCH_r18
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlCAPSFCHR18RxPSFCHResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Rx_PSFCH_Resource_r18 = v
			}
			{
				v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlCAPSFCHR18TxPSFCHResourceR18Constraints)
				if err != nil {
					return err
				}
				c.Tx_PSFCH_Resource_r18 = v
			}
		}

		if groupSeq.IsComponentPresent(16) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlReceptionIntraCarrierGuardBandR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_ReceptionIntraCarrierGuardBand_r18 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sl-PathlossBasedOLPC-SL-RSRP-Report-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(bandSidelinkR16ExtSlPathlossBasedOLPCSLRSRPReportR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PathlossBasedOLPC_SL_RSRP_Report_r18 = &v
		}
	}

	return nil
}
