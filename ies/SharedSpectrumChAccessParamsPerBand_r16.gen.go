// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsPerBand-r16 (line 24857).

var sharedSpectrumChAccessParamsPerBandR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-DynamicChAccess-r16", Optional: true},
		{Name: "ul-Semi-StaticChAccess-r16", Optional: true},
		{Name: "ssb-RRM-DynamicChAccess-r16", Optional: true},
		{Name: "ssb-RRM-Semi-StaticChAccess-r16", Optional: true},
		{Name: "mib-Acquisition-r16", Optional: true},
		{Name: "ssb-RLM-DynamicChAccess-r16", Optional: true},
		{Name: "ssb-RLM-Semi-StaticChAccess-r16", Optional: true},
		{Name: "sib1-Acquisition-r16", Optional: true},
		{Name: "extRA-ResponseWindow-r16", Optional: true},
		{Name: "ssb-BFD-CBD-dynamicChannelAccess-r16", Optional: true},
		{Name: "ssb-BFD-CBD-semi-staticChannelAccess-r16", Optional: true},
		{Name: "csi-RS-BFD-CBD-r16", Optional: true},
		{Name: "ul-ChannelBW-SCell-10mhz-r16", Optional: true},
		{Name: "rssi-ChannelOccupancyReporting-r16", Optional: true},
		{Name: "srs-StartAnyOFDM-Symbol-r16", Optional: true},
		{Name: "searchSpaceFreqMonitorLocation-r16", Optional: true},
		{Name: "coreset-RB-Offset-r16", Optional: true},
		{Name: "cgi-Acquisition-r16", Optional: true},
		{Name: "configuredUL-Tx-r16", Optional: true},
		{Name: "prach-Wideband-r16", Optional: true},
		{Name: "dci-AvailableRB-Set-r16", Optional: true},
		{Name: "dci-ChOccupancyDuration-r16", Optional: true},
		{Name: "typeB-PDSCH-length-r16", Optional: true},
		{Name: "searchSpaceSwitchWithDCI-r16", Optional: true},
		{Name: "searchSpaceSwitchWithoutDCI-r16", Optional: true},
		{Name: "searchSpaceSwitchCapability2-r16", Optional: true},
		{Name: "non-numericalPDSCH-HARQ-timing-r16", Optional: true},
		{Name: "enhancedDynamicHARQ-codebook-r16", Optional: true},
		{Name: "oneShotHARQ-feedback-r16", Optional: true},
		{Name: "multiPUSCH-UL-grant-r16", Optional: true},
		{Name: "csi-RS-RLM-r16", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "periodicAndSemi-PersistentCSI-RS-r16", Optional: true},
		{Name: "pusch-PRB-interlace-r16", Optional: true},
		{Name: "pucch-F0-F1-PRB-Interlace-r16", Optional: true},
		{Name: "occ-PRB-PF2-PF3-r16", Optional: true},
		{Name: "extCP-rangeCG-PUSCH-r16", Optional: true},
		{Name: "configuredGrantWithReTx-r16", Optional: true},
		{Name: "ed-Threshold-r16", Optional: true},
		{Name: "ul-DL-COT-Sharing-r16", Optional: true},
		{Name: "mux-CG-UCI-HARQ-ACK-r16", Optional: true},
		{Name: "cg-resourceConfig-r16", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ul_DynamicChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16UlDynamicChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ul_DynamicChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ul_Semi_StaticChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16UlSemiStaticChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ul_Semi_StaticChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_RRM_DynamicChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbRRMDynamicChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_RRM_DynamicChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_RRM_Semi_StaticChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbRRMSemiStaticChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_RRM_Semi_StaticChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Mib_Acquisition_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16MibAcquisitionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Mib_Acquisition_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_RLM_DynamicChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbRLMDynamicChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_RLM_DynamicChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_RLM_Semi_StaticChAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbRLMSemiStaticChAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_RLM_Semi_StaticChAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Sib1_Acquisition_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16Sib1AcquisitionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Sib1_Acquisition_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_ExtRA_ResponseWindow_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16ExtRAResponseWindowR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_ExtRA_ResponseWindow_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_BFD_CBD_DynamicChannelAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDDynamicChannelAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_BFD_CBD_DynamicChannelAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ssb_BFD_CBD_Semi_StaticChannelAccess_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDSemiStaticChannelAccessR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ssb_BFD_CBD_Semi_StaticChannelAccess_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Csi_RS_BFD_CBD_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16CsiRSBFDCBDR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Csi_RS_BFD_CBD_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ul_ChannelBW_SCell_10mhz_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16UlChannelBWSCell10mhzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ul_ChannelBW_SCell_10mhz_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Rssi_ChannelOccupancyReporting_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16RssiChannelOccupancyReportingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Rssi_ChannelOccupancyReporting_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Srs_StartAnyOFDM_Symbol_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SrsStartAnyOFDMSymbolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Srs_StartAnyOFDM_Symbol_r16_Supported},
}

var sharedSpectrumChAccessParamsPerBandR16SearchSpaceFreqMonitorLocationR16Constraints = per.Constrained(1, 5)

const (
	SharedSpectrumChAccessParamsPerBand_r16_Coreset_RB_Offset_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16CoresetRBOffsetR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Coreset_RB_Offset_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Cgi_Acquisition_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16CgiAcquisitionR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Cgi_Acquisition_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_ConfiguredUL_Tx_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16ConfiguredULTxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_ConfiguredUL_Tx_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Prach_Wideband_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16PrachWidebandR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Prach_Wideband_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Dci_AvailableRB_Set_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16DciAvailableRBSetR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Dci_AvailableRB_Set_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Dci_ChOccupancyDuration_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16DciChOccupancyDurationR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Dci_ChOccupancyDuration_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_TypeB_PDSCH_Length_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16TypeBPDSCHLengthR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_TypeB_PDSCH_Length_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchWithDCI_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithDCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchWithDCI_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchWithoutDCI_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithoutDCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchWithoutDCI_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchCapability2_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchCapability2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_SearchSpaceSwitchCapability2_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Non_NumericalPDSCH_HARQ_Timing_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16NonNumericalPDSCHHARQTimingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Non_NumericalPDSCH_HARQ_Timing_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_EnhancedDynamicHARQ_Codebook_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16EnhancedDynamicHARQCodebookR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_EnhancedDynamicHARQ_Codebook_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_OneShotHARQ_Feedback_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16OneShotHARQFeedbackR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_OneShotHARQ_Feedback_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_MultiPUSCH_UL_Grant_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16MultiPUSCHULGrantR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_MultiPUSCH_UL_Grant_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Csi_RS_RLM_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16CsiRSRLMR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Csi_RS_RLM_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Dummy_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16DummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Dummy_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_PeriodicAndSemi_PersistentCSI_RS_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16PeriodicAndSemiPersistentCSIRSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_PeriodicAndSemi_PersistentCSI_RS_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Pusch_PRB_Interlace_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16PuschPRBInterlaceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Pusch_PRB_Interlace_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Pucch_F0_F1_PRB_Interlace_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16PucchF0F1PRBInterlaceR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Pucch_F0_F1_PRB_Interlace_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Occ_PRB_PF2_PF3_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16OccPRBPF2PF3R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Occ_PRB_PF2_PF3_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_ExtCP_RangeCG_PUSCH_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16ExtCPRangeCGPUSCHR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_ExtCP_RangeCG_PUSCH_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_ConfiguredGrantWithReTx_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16ConfiguredGrantWithReTxR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_ConfiguredGrantWithReTx_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ed_Threshold_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16EdThresholdR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ed_Threshold_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Ul_DL_COT_Sharing_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16UlDLCOTSharingR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Ul_DL_COT_Sharing_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Mux_CG_UCI_HARQ_ACK_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16MuxCGUCIHARQACKR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Mux_CG_UCI_HARQ_ACK_r16_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_r16_Cg_ResourceConfig_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandR16CgResourceConfigR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_r16_Cg_ResourceConfig_r16_Supported},
}

type SharedSpectrumChAccessParamsPerBand_r16 struct {
	Ul_DynamicChAccess_r16                   *int64
	Ul_Semi_StaticChAccess_r16               *int64
	Ssb_RRM_DynamicChAccess_r16              *int64
	Ssb_RRM_Semi_StaticChAccess_r16          *int64
	Mib_Acquisition_r16                      *int64
	Ssb_RLM_DynamicChAccess_r16              *int64
	Ssb_RLM_Semi_StaticChAccess_r16          *int64
	Sib1_Acquisition_r16                     *int64
	ExtRA_ResponseWindow_r16                 *int64
	Ssb_BFD_CBD_DynamicChannelAccess_r16     *int64
	Ssb_BFD_CBD_Semi_StaticChannelAccess_r16 *int64
	Csi_RS_BFD_CBD_r16                       *int64
	Ul_ChannelBW_SCell_10mhz_r16             *int64
	Rssi_ChannelOccupancyReporting_r16       *int64
	Srs_StartAnyOFDM_Symbol_r16              *int64
	SearchSpaceFreqMonitorLocation_r16       *int64
	Coreset_RB_Offset_r16                    *int64
	Cgi_Acquisition_r16                      *int64
	ConfiguredUL_Tx_r16                      *int64
	Prach_Wideband_r16                       *int64
	Dci_AvailableRB_Set_r16                  *int64
	Dci_ChOccupancyDuration_r16              *int64
	TypeB_PDSCH_Length_r16                   *int64
	SearchSpaceSwitchWithDCI_r16             *int64
	SearchSpaceSwitchWithoutDCI_r16          *int64
	SearchSpaceSwitchCapability2_r16         *int64
	Non_NumericalPDSCH_HARQ_Timing_r16       *int64
	EnhancedDynamicHARQ_Codebook_r16         *int64
	OneShotHARQ_Feedback_r16                 *int64
	MultiPUSCH_UL_Grant_r16                  *int64
	Csi_RS_RLM_r16                           *int64
	Dummy                                    *int64
	PeriodicAndSemi_PersistentCSI_RS_r16     *int64
	Pusch_PRB_Interlace_r16                  *int64
	Pucch_F0_F1_PRB_Interlace_r16            *int64
	Occ_PRB_PF2_PF3_r16                      *int64
	ExtCP_RangeCG_PUSCH_r16                  *int64
	ConfiguredGrantWithReTx_r16              *int64
	Ed_Threshold_r16                         *int64
	Ul_DL_COT_Sharing_r16                    *int64
	Mux_CG_UCI_HARQ_ACK_r16                  *int64
	Cg_ResourceConfig_r16                    *int64
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsPerBandR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_DynamicChAccess_r16 != nil, ie.Ul_Semi_StaticChAccess_r16 != nil, ie.Ssb_RRM_DynamicChAccess_r16 != nil, ie.Ssb_RRM_Semi_StaticChAccess_r16 != nil, ie.Mib_Acquisition_r16 != nil, ie.Ssb_RLM_DynamicChAccess_r16 != nil, ie.Ssb_RLM_Semi_StaticChAccess_r16 != nil, ie.Sib1_Acquisition_r16 != nil, ie.ExtRA_ResponseWindow_r16 != nil, ie.Ssb_BFD_CBD_DynamicChannelAccess_r16 != nil, ie.Ssb_BFD_CBD_Semi_StaticChannelAccess_r16 != nil, ie.Csi_RS_BFD_CBD_r16 != nil, ie.Ul_ChannelBW_SCell_10mhz_r16 != nil, ie.Rssi_ChannelOccupancyReporting_r16 != nil, ie.Srs_StartAnyOFDM_Symbol_r16 != nil, ie.SearchSpaceFreqMonitorLocation_r16 != nil, ie.Coreset_RB_Offset_r16 != nil, ie.Cgi_Acquisition_r16 != nil, ie.ConfiguredUL_Tx_r16 != nil, ie.Prach_Wideband_r16 != nil, ie.Dci_AvailableRB_Set_r16 != nil, ie.Dci_ChOccupancyDuration_r16 != nil, ie.TypeB_PDSCH_Length_r16 != nil, ie.SearchSpaceSwitchWithDCI_r16 != nil, ie.SearchSpaceSwitchWithoutDCI_r16 != nil, ie.SearchSpaceSwitchCapability2_r16 != nil, ie.Non_NumericalPDSCH_HARQ_Timing_r16 != nil, ie.EnhancedDynamicHARQ_Codebook_r16 != nil, ie.OneShotHARQ_Feedback_r16 != nil, ie.MultiPUSCH_UL_Grant_r16 != nil, ie.Csi_RS_RLM_r16 != nil, ie.Dummy != nil, ie.PeriodicAndSemi_PersistentCSI_RS_r16 != nil, ie.Pusch_PRB_Interlace_r16 != nil, ie.Pucch_F0_F1_PRB_Interlace_r16 != nil, ie.Occ_PRB_PF2_PF3_r16 != nil, ie.ExtCP_RangeCG_PUSCH_r16 != nil, ie.ConfiguredGrantWithReTx_r16 != nil, ie.Ed_Threshold_r16 != nil, ie.Ul_DL_COT_Sharing_r16 != nil, ie.Mux_CG_UCI_HARQ_ACK_r16 != nil, ie.Cg_ResourceConfig_r16 != nil}); err != nil {
		return err
	}

	// 2. ul-DynamicChAccess-r16: enumerated
	{
		if ie.Ul_DynamicChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_DynamicChAccess_r16, sharedSpectrumChAccessParamsPerBandR16UlDynamicChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ul-Semi-StaticChAccess-r16: enumerated
	{
		if ie.Ul_Semi_StaticChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_Semi_StaticChAccess_r16, sharedSpectrumChAccessParamsPerBandR16UlSemiStaticChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. ssb-RRM-DynamicChAccess-r16: enumerated
	{
		if ie.Ssb_RRM_DynamicChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_RRM_DynamicChAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbRRMDynamicChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. ssb-RRM-Semi-StaticChAccess-r16: enumerated
	{
		if ie.Ssb_RRM_Semi_StaticChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_RRM_Semi_StaticChAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbRRMSemiStaticChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. mib-Acquisition-r16: enumerated
	{
		if ie.Mib_Acquisition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mib_Acquisition_r16, sharedSpectrumChAccessParamsPerBandR16MibAcquisitionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. ssb-RLM-DynamicChAccess-r16: enumerated
	{
		if ie.Ssb_RLM_DynamicChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_RLM_DynamicChAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbRLMDynamicChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 8. ssb-RLM-Semi-StaticChAccess-r16: enumerated
	{
		if ie.Ssb_RLM_Semi_StaticChAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_RLM_Semi_StaticChAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbRLMSemiStaticChAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. sib1-Acquisition-r16: enumerated
	{
		if ie.Sib1_Acquisition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Sib1_Acquisition_r16, sharedSpectrumChAccessParamsPerBandR16Sib1AcquisitionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 10. extRA-ResponseWindow-r16: enumerated
	{
		if ie.ExtRA_ResponseWindow_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ExtRA_ResponseWindow_r16, sharedSpectrumChAccessParamsPerBandR16ExtRAResponseWindowR16Constraints); err != nil {
				return err
			}
		}
	}

	// 11. ssb-BFD-CBD-dynamicChannelAccess-r16: enumerated
	{
		if ie.Ssb_BFD_CBD_DynamicChannelAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_BFD_CBD_DynamicChannelAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDDynamicChannelAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 12. ssb-BFD-CBD-semi-staticChannelAccess-r16: enumerated
	{
		if ie.Ssb_BFD_CBD_Semi_StaticChannelAccess_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ssb_BFD_CBD_Semi_StaticChannelAccess_r16, sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDSemiStaticChannelAccessR16Constraints); err != nil {
				return err
			}
		}
	}

	// 13. csi-RS-BFD-CBD-r16: enumerated
	{
		if ie.Csi_RS_BFD_CBD_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_BFD_CBD_r16, sharedSpectrumChAccessParamsPerBandR16CsiRSBFDCBDR16Constraints); err != nil {
				return err
			}
		}
	}

	// 14. ul-ChannelBW-SCell-10mhz-r16: enumerated
	{
		if ie.Ul_ChannelBW_SCell_10mhz_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_ChannelBW_SCell_10mhz_r16, sharedSpectrumChAccessParamsPerBandR16UlChannelBWSCell10mhzR16Constraints); err != nil {
				return err
			}
		}
	}

	// 15. rssi-ChannelOccupancyReporting-r16: enumerated
	{
		if ie.Rssi_ChannelOccupancyReporting_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Rssi_ChannelOccupancyReporting_r16, sharedSpectrumChAccessParamsPerBandR16RssiChannelOccupancyReportingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 16. srs-StartAnyOFDM-Symbol-r16: enumerated
	{
		if ie.Srs_StartAnyOFDM_Symbol_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Srs_StartAnyOFDM_Symbol_r16, sharedSpectrumChAccessParamsPerBandR16SrsStartAnyOFDMSymbolR16Constraints); err != nil {
				return err
			}
		}
	}

	// 17. searchSpaceFreqMonitorLocation-r16: integer
	{
		if ie.SearchSpaceFreqMonitorLocation_r16 != nil {
			if err := e.EncodeInteger(*ie.SearchSpaceFreqMonitorLocation_r16, sharedSpectrumChAccessParamsPerBandR16SearchSpaceFreqMonitorLocationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 18. coreset-RB-Offset-r16: enumerated
	{
		if ie.Coreset_RB_Offset_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Coreset_RB_Offset_r16, sharedSpectrumChAccessParamsPerBandR16CoresetRBOffsetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 19. cgi-Acquisition-r16: enumerated
	{
		if ie.Cgi_Acquisition_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Cgi_Acquisition_r16, sharedSpectrumChAccessParamsPerBandR16CgiAcquisitionR16Constraints); err != nil {
				return err
			}
		}
	}

	// 20. configuredUL-Tx-r16: enumerated
	{
		if ie.ConfiguredUL_Tx_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredUL_Tx_r16, sharedSpectrumChAccessParamsPerBandR16ConfiguredULTxR16Constraints); err != nil {
				return err
			}
		}
	}

	// 21. prach-Wideband-r16: enumerated
	{
		if ie.Prach_Wideband_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Prach_Wideband_r16, sharedSpectrumChAccessParamsPerBandR16PrachWidebandR16Constraints); err != nil {
				return err
			}
		}
	}

	// 22. dci-AvailableRB-Set-r16: enumerated
	{
		if ie.Dci_AvailableRB_Set_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dci_AvailableRB_Set_r16, sharedSpectrumChAccessParamsPerBandR16DciAvailableRBSetR16Constraints); err != nil {
				return err
			}
		}
	}

	// 23. dci-ChOccupancyDuration-r16: enumerated
	{
		if ie.Dci_ChOccupancyDuration_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Dci_ChOccupancyDuration_r16, sharedSpectrumChAccessParamsPerBandR16DciChOccupancyDurationR16Constraints); err != nil {
				return err
			}
		}
	}

	// 24. typeB-PDSCH-length-r16: enumerated
	{
		if ie.TypeB_PDSCH_Length_r16 != nil {
			if err := e.EncodeEnumerated(*ie.TypeB_PDSCH_Length_r16, sharedSpectrumChAccessParamsPerBandR16TypeBPDSCHLengthR16Constraints); err != nil {
				return err
			}
		}
	}

	// 25. searchSpaceSwitchWithDCI-r16: enumerated
	{
		if ie.SearchSpaceSwitchWithDCI_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SearchSpaceSwitchWithDCI_r16, sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithDCIR16Constraints); err != nil {
				return err
			}
		}
	}

	// 26. searchSpaceSwitchWithoutDCI-r16: enumerated
	{
		if ie.SearchSpaceSwitchWithoutDCI_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SearchSpaceSwitchWithoutDCI_r16, sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithoutDCIR16Constraints); err != nil {
				return err
			}
		}
	}

	// 27. searchSpaceSwitchCapability2-r16: enumerated
	{
		if ie.SearchSpaceSwitchCapability2_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SearchSpaceSwitchCapability2_r16, sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchCapability2R16Constraints); err != nil {
				return err
			}
		}
	}

	// 28. non-numericalPDSCH-HARQ-timing-r16: enumerated
	{
		if ie.Non_NumericalPDSCH_HARQ_Timing_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Non_NumericalPDSCH_HARQ_Timing_r16, sharedSpectrumChAccessParamsPerBandR16NonNumericalPDSCHHARQTimingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 29. enhancedDynamicHARQ-codebook-r16: enumerated
	{
		if ie.EnhancedDynamicHARQ_Codebook_r16 != nil {
			if err := e.EncodeEnumerated(*ie.EnhancedDynamicHARQ_Codebook_r16, sharedSpectrumChAccessParamsPerBandR16EnhancedDynamicHARQCodebookR16Constraints); err != nil {
				return err
			}
		}
	}

	// 30. oneShotHARQ-feedback-r16: enumerated
	{
		if ie.OneShotHARQ_Feedback_r16 != nil {
			if err := e.EncodeEnumerated(*ie.OneShotHARQ_Feedback_r16, sharedSpectrumChAccessParamsPerBandR16OneShotHARQFeedbackR16Constraints); err != nil {
				return err
			}
		}
	}

	// 31. multiPUSCH-UL-grant-r16: enumerated
	{
		if ie.MultiPUSCH_UL_Grant_r16 != nil {
			if err := e.EncodeEnumerated(*ie.MultiPUSCH_UL_Grant_r16, sharedSpectrumChAccessParamsPerBandR16MultiPUSCHULGrantR16Constraints); err != nil {
				return err
			}
		}
	}

	// 32. csi-RS-RLM-r16: enumerated
	{
		if ie.Csi_RS_RLM_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_RLM_r16, sharedSpectrumChAccessParamsPerBandR16CsiRSRLMR16Constraints); err != nil {
				return err
			}
		}
	}

	// 33. dummy: enumerated
	{
		if ie.Dummy != nil {
			if err := e.EncodeEnumerated(*ie.Dummy, sharedSpectrumChAccessParamsPerBandR16DummyConstraints); err != nil {
				return err
			}
		}
	}

	// 34. periodicAndSemi-PersistentCSI-RS-r16: enumerated
	{
		if ie.PeriodicAndSemi_PersistentCSI_RS_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PeriodicAndSemi_PersistentCSI_RS_r16, sharedSpectrumChAccessParamsPerBandR16PeriodicAndSemiPersistentCSIRSR16Constraints); err != nil {
				return err
			}
		}
	}

	// 35. pusch-PRB-interlace-r16: enumerated
	{
		if ie.Pusch_PRB_Interlace_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pusch_PRB_Interlace_r16, sharedSpectrumChAccessParamsPerBandR16PuschPRBInterlaceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 36. pucch-F0-F1-PRB-Interlace-r16: enumerated
	{
		if ie.Pucch_F0_F1_PRB_Interlace_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pucch_F0_F1_PRB_Interlace_r16, sharedSpectrumChAccessParamsPerBandR16PucchF0F1PRBInterlaceR16Constraints); err != nil {
				return err
			}
		}
	}

	// 37. occ-PRB-PF2-PF3-r16: enumerated
	{
		if ie.Occ_PRB_PF2_PF3_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Occ_PRB_PF2_PF3_r16, sharedSpectrumChAccessParamsPerBandR16OccPRBPF2PF3R16Constraints); err != nil {
				return err
			}
		}
	}

	// 38. extCP-rangeCG-PUSCH-r16: enumerated
	{
		if ie.ExtCP_RangeCG_PUSCH_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ExtCP_RangeCG_PUSCH_r16, sharedSpectrumChAccessParamsPerBandR16ExtCPRangeCGPUSCHR16Constraints); err != nil {
				return err
			}
		}
	}

	// 39. configuredGrantWithReTx-r16: enumerated
	{
		if ie.ConfiguredGrantWithReTx_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredGrantWithReTx_r16, sharedSpectrumChAccessParamsPerBandR16ConfiguredGrantWithReTxR16Constraints); err != nil {
				return err
			}
		}
	}

	// 40. ed-Threshold-r16: enumerated
	{
		if ie.Ed_Threshold_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ed_Threshold_r16, sharedSpectrumChAccessParamsPerBandR16EdThresholdR16Constraints); err != nil {
				return err
			}
		}
	}

	// 41. ul-DL-COT-Sharing-r16: enumerated
	{
		if ie.Ul_DL_COT_Sharing_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_DL_COT_Sharing_r16, sharedSpectrumChAccessParamsPerBandR16UlDLCOTSharingR16Constraints); err != nil {
				return err
			}
		}
	}

	// 42. mux-CG-UCI-HARQ-ACK-r16: enumerated
	{
		if ie.Mux_CG_UCI_HARQ_ACK_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Mux_CG_UCI_HARQ_ACK_r16, sharedSpectrumChAccessParamsPerBandR16MuxCGUCIHARQACKR16Constraints); err != nil {
				return err
			}
		}
	}

	// 43. cg-resourceConfig-r16: enumerated
	{
		if ie.Cg_ResourceConfig_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Cg_ResourceConfig_r16, sharedSpectrumChAccessParamsPerBandR16CgResourceConfigR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsPerBandR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-DynamicChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16UlDynamicChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_DynamicChAccess_r16 = &idx
		}
	}

	// 3. ul-Semi-StaticChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16UlSemiStaticChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_Semi_StaticChAccess_r16 = &idx
		}
	}

	// 4. ssb-RRM-DynamicChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbRRMDynamicChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_RRM_DynamicChAccess_r16 = &idx
		}
	}

	// 5. ssb-RRM-Semi-StaticChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbRRMSemiStaticChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_RRM_Semi_StaticChAccess_r16 = &idx
		}
	}

	// 6. mib-Acquisition-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16MibAcquisitionR16Constraints)
			if err != nil {
				return err
			}
			ie.Mib_Acquisition_r16 = &idx
		}
	}

	// 7. ssb-RLM-DynamicChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbRLMDynamicChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_RLM_DynamicChAccess_r16 = &idx
		}
	}

	// 8. ssb-RLM-Semi-StaticChAccess-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbRLMSemiStaticChAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_RLM_Semi_StaticChAccess_r16 = &idx
		}
	}

	// 9. sib1-Acquisition-r16: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16Sib1AcquisitionR16Constraints)
			if err != nil {
				return err
			}
			ie.Sib1_Acquisition_r16 = &idx
		}
	}

	// 10. extRA-ResponseWindow-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16ExtRAResponseWindowR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtRA_ResponseWindow_r16 = &idx
		}
	}

	// 11. ssb-BFD-CBD-dynamicChannelAccess-r16: enumerated
	{
		if seq.IsComponentPresent(9) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDDynamicChannelAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_BFD_CBD_DynamicChannelAccess_r16 = &idx
		}
	}

	// 12. ssb-BFD-CBD-semi-staticChannelAccess-r16: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SsbBFDCBDSemiStaticChannelAccessR16Constraints)
			if err != nil {
				return err
			}
			ie.Ssb_BFD_CBD_Semi_StaticChannelAccess_r16 = &idx
		}
	}

	// 13. csi-RS-BFD-CBD-r16: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16CsiRSBFDCBDR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_BFD_CBD_r16 = &idx
		}
	}

	// 14. ul-ChannelBW-SCell-10mhz-r16: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16UlChannelBWSCell10mhzR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_ChannelBW_SCell_10mhz_r16 = &idx
		}
	}

	// 15. rssi-ChannelOccupancyReporting-r16: enumerated
	{
		if seq.IsComponentPresent(13) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16RssiChannelOccupancyReportingR16Constraints)
			if err != nil {
				return err
			}
			ie.Rssi_ChannelOccupancyReporting_r16 = &idx
		}
	}

	// 16. srs-StartAnyOFDM-Symbol-r16: enumerated
	{
		if seq.IsComponentPresent(14) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SrsStartAnyOFDMSymbolR16Constraints)
			if err != nil {
				return err
			}
			ie.Srs_StartAnyOFDM_Symbol_r16 = &idx
		}
	}

	// 17. searchSpaceFreqMonitorLocation-r16: integer
	{
		if seq.IsComponentPresent(15) {
			v, err := d.DecodeInteger(sharedSpectrumChAccessParamsPerBandR16SearchSpaceFreqMonitorLocationR16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceFreqMonitorLocation_r16 = &v
		}
	}

	// 18. coreset-RB-Offset-r16: enumerated
	{
		if seq.IsComponentPresent(16) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16CoresetRBOffsetR16Constraints)
			if err != nil {
				return err
			}
			ie.Coreset_RB_Offset_r16 = &idx
		}
	}

	// 19. cgi-Acquisition-r16: enumerated
	{
		if seq.IsComponentPresent(17) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16CgiAcquisitionR16Constraints)
			if err != nil {
				return err
			}
			ie.Cgi_Acquisition_r16 = &idx
		}
	}

	// 20. configuredUL-Tx-r16: enumerated
	{
		if seq.IsComponentPresent(18) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16ConfiguredULTxR16Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredUL_Tx_r16 = &idx
		}
	}

	// 21. prach-Wideband-r16: enumerated
	{
		if seq.IsComponentPresent(19) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16PrachWidebandR16Constraints)
			if err != nil {
				return err
			}
			ie.Prach_Wideband_r16 = &idx
		}
	}

	// 22. dci-AvailableRB-Set-r16: enumerated
	{
		if seq.IsComponentPresent(20) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16DciAvailableRBSetR16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_AvailableRB_Set_r16 = &idx
		}
	}

	// 23. dci-ChOccupancyDuration-r16: enumerated
	{
		if seq.IsComponentPresent(21) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16DciChOccupancyDurationR16Constraints)
			if err != nil {
				return err
			}
			ie.Dci_ChOccupancyDuration_r16 = &idx
		}
	}

	// 24. typeB-PDSCH-length-r16: enumerated
	{
		if seq.IsComponentPresent(22) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16TypeBPDSCHLengthR16Constraints)
			if err != nil {
				return err
			}
			ie.TypeB_PDSCH_Length_r16 = &idx
		}
	}

	// 25. searchSpaceSwitchWithDCI-r16: enumerated
	{
		if seq.IsComponentPresent(23) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithDCIR16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchWithDCI_r16 = &idx
		}
	}

	// 26. searchSpaceSwitchWithoutDCI-r16: enumerated
	{
		if seq.IsComponentPresent(24) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchWithoutDCIR16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchWithoutDCI_r16 = &idx
		}
	}

	// 27. searchSpaceSwitchCapability2-r16: enumerated
	{
		if seq.IsComponentPresent(25) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16SearchSpaceSwitchCapability2R16Constraints)
			if err != nil {
				return err
			}
			ie.SearchSpaceSwitchCapability2_r16 = &idx
		}
	}

	// 28. non-numericalPDSCH-HARQ-timing-r16: enumerated
	{
		if seq.IsComponentPresent(26) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16NonNumericalPDSCHHARQTimingR16Constraints)
			if err != nil {
				return err
			}
			ie.Non_NumericalPDSCH_HARQ_Timing_r16 = &idx
		}
	}

	// 29. enhancedDynamicHARQ-codebook-r16: enumerated
	{
		if seq.IsComponentPresent(27) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16EnhancedDynamicHARQCodebookR16Constraints)
			if err != nil {
				return err
			}
			ie.EnhancedDynamicHARQ_Codebook_r16 = &idx
		}
	}

	// 30. oneShotHARQ-feedback-r16: enumerated
	{
		if seq.IsComponentPresent(28) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16OneShotHARQFeedbackR16Constraints)
			if err != nil {
				return err
			}
			ie.OneShotHARQ_Feedback_r16 = &idx
		}
	}

	// 31. multiPUSCH-UL-grant-r16: enumerated
	{
		if seq.IsComponentPresent(29) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16MultiPUSCHULGrantR16Constraints)
			if err != nil {
				return err
			}
			ie.MultiPUSCH_UL_Grant_r16 = &idx
		}
	}

	// 32. csi-RS-RLM-r16: enumerated
	{
		if seq.IsComponentPresent(30) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16CsiRSRLMR16Constraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_RLM_r16 = &idx
		}
	}

	// 33. dummy: enumerated
	{
		if seq.IsComponentPresent(31) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16DummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &idx
		}
	}

	// 34. periodicAndSemi-PersistentCSI-RS-r16: enumerated
	{
		if seq.IsComponentPresent(32) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16PeriodicAndSemiPersistentCSIRSR16Constraints)
			if err != nil {
				return err
			}
			ie.PeriodicAndSemi_PersistentCSI_RS_r16 = &idx
		}
	}

	// 35. pusch-PRB-interlace-r16: enumerated
	{
		if seq.IsComponentPresent(33) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16PuschPRBInterlaceR16Constraints)
			if err != nil {
				return err
			}
			ie.Pusch_PRB_Interlace_r16 = &idx
		}
	}

	// 36. pucch-F0-F1-PRB-Interlace-r16: enumerated
	{
		if seq.IsComponentPresent(34) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16PucchF0F1PRBInterlaceR16Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_F0_F1_PRB_Interlace_r16 = &idx
		}
	}

	// 37. occ-PRB-PF2-PF3-r16: enumerated
	{
		if seq.IsComponentPresent(35) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16OccPRBPF2PF3R16Constraints)
			if err != nil {
				return err
			}
			ie.Occ_PRB_PF2_PF3_r16 = &idx
		}
	}

	// 38. extCP-rangeCG-PUSCH-r16: enumerated
	{
		if seq.IsComponentPresent(36) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16ExtCPRangeCGPUSCHR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtCP_RangeCG_PUSCH_r16 = &idx
		}
	}

	// 39. configuredGrantWithReTx-r16: enumerated
	{
		if seq.IsComponentPresent(37) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16ConfiguredGrantWithReTxR16Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredGrantWithReTx_r16 = &idx
		}
	}

	// 40. ed-Threshold-r16: enumerated
	{
		if seq.IsComponentPresent(38) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16EdThresholdR16Constraints)
			if err != nil {
				return err
			}
			ie.Ed_Threshold_r16 = &idx
		}
	}

	// 41. ul-DL-COT-Sharing-r16: enumerated
	{
		if seq.IsComponentPresent(39) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16UlDLCOTSharingR16Constraints)
			if err != nil {
				return err
			}
			ie.Ul_DL_COT_Sharing_r16 = &idx
		}
	}

	// 42. mux-CG-UCI-HARQ-ACK-r16: enumerated
	{
		if seq.IsComponentPresent(40) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16MuxCGUCIHARQACKR16Constraints)
			if err != nil {
				return err
			}
			ie.Mux_CG_UCI_HARQ_ACK_r16 = &idx
		}
	}

	// 43. cg-resourceConfig-r16: enumerated
	{
		if seq.IsComponentPresent(41) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandR16CgResourceConfigR16Constraints)
			if err != nil {
				return err
			}
			ie.Cg_ResourceConfig_r16 = &idx
		}
	}

	return nil
}
