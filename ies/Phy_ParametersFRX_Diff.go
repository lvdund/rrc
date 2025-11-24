package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff struct {
	DynamicSFI                                        *Phy_ParametersFRX_Diff_dynamicSFI                                        `optional`
	Dummy1                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	TwoFL_DMRS                                        *uper.BitString                                                           `lb:2,ub:2,optional`
	Dummy2                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	Dummy3                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	SupportedDMRS_TypeDL                              *Phy_ParametersFRX_Diff_supportedDMRS_TypeDL                              `optional`
	SupportedDMRS_TypeUL                              *Phy_ParametersFRX_Diff_supportedDMRS_TypeUL                              `optional`
	SemiOpenLoopCSI                                   *Phy_ParametersFRX_Diff_semiOpenLoopCSI                                   `optional`
	Csi_ReportWithoutPMI                              *Phy_ParametersFRX_Diff_csi_ReportWithoutPMI                              `optional`
	Csi_ReportWithoutCQI                              *Phy_ParametersFRX_Diff_csi_ReportWithoutCQI                              `optional`
	OnePortsPTRS                                      *uper.BitString                                                           `lb:2,ub:2,optional`
	TwoPUCCH_F0_2_ConsecSymbols                       *Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols                       `optional`
	Pucch_F2_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F2_WithFH                                   `optional`
	Pucch_F3_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F3_WithFH                                   `optional`
	Pucch_F4_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F4_WithFH                                   `optional`
	Pucch_F0_2WithoutFH                               *Phy_ParametersFRX_Diff_pucch_F0_2WithoutFH                               `optional`
	Pucch_F1_3_4WithoutFH                             *Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH                             `optional`
	Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot            *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot            `optional`
	Uci_CodeBlockSegmentation                         *Phy_ParametersFRX_Diff_uci_CodeBlockSegmentation                         `optional`
	OnePUCCH_LongAndShortFormat                       *Phy_ParametersFRX_Diff_onePUCCH_LongAndShortFormat                       `optional`
	TwoPUCCH_AnyOthersInSlot                          *Phy_ParametersFRX_Diff_twoPUCCH_AnyOthersInSlot                          `optional`
	IntraSlotFreqHopping_PUSCH                        *Phy_ParametersFRX_Diff_intraSlotFreqHopping_PUSCH                        `optional`
	Pusch_LBRM                                        *Phy_ParametersFRX_Diff_pusch_LBRM                                        `optional`
	Pdcch_BlindDetectionCA                            *int64                                                                    `lb:4,ub:16,optional`
	Tpc_PUSCH_RNTI                                    *Phy_ParametersFRX_Diff_tpc_PUSCH_RNTI                                    `optional`
	Tpc_PUCCH_RNTI                                    *Phy_ParametersFRX_Diff_tpc_PUCCH_RNTI                                    `optional`
	Tpc_SRS_RNTI                                      *Phy_ParametersFRX_Diff_tpc_SRS_RNTI                                      `optional`
	AbsoluteTPC_Command                               *Phy_ParametersFRX_Diff_absoluteTPC_Command                               `optional`
	TwoDifferentTPC_Loop_PUSCH                        *Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUSCH                        `optional`
	TwoDifferentTPC_Loop_PUCCH                        *Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUCCH                        `optional`
	Pusch_HalfPi_BPSK                                 *Phy_ParametersFRX_Diff_pusch_HalfPi_BPSK                                 `optional`
	Pucch_F3_4_HalfPi_BPSK                            *Phy_ParametersFRX_Diff_pucch_F3_4_HalfPi_BPSK                            `optional`
	AlmostContiguousCP_OFDM_UL                        *Phy_ParametersFRX_Diff_almostContiguousCP_OFDM_UL                        `optional`
	Sp_CSI_RS                                         *Phy_ParametersFRX_Diff_sp_CSI_RS                                         `optional`
	Sp_CSI_IM                                         *Phy_ParametersFRX_Diff_sp_CSI_IM                                         `optional`
	Tdd_MultiDL_UL_SwitchPerSlot                      *Phy_ParametersFRX_Diff_tdd_MultiDL_UL_SwitchPerSlot                      `optional`
	MultipleCORESET                                   *Phy_ParametersFRX_Diff_multipleCORESET                                   `optional`
	Csi_RS_IM_ReceptionForFeedback                    *CSI_RS_IM_ReceptionForFeedback                                           `optional,ext-1`
	Csi_RS_ProcFrameworkForSRS                        *CSI_RS_ProcFrameworkForSRS                                               `optional,ext-1`
	Csi_ReportFramework                               *CSI_ReportFramework                                                      `optional,ext-1`
	Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot             *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot             `optional,ext-1`
	Mux_SR_HARQ_ACK_PUCCH                             *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_PUCCH                             `optional,ext-1`
	Mux_MultipleGroupCtrlCH_Overlap                   *Phy_ParametersFRX_Diff_mux_MultipleGroupCtrlCH_Overlap                   `optional,ext-1`
	Dl_SchedulingOffset_PDSCH_TypeA                   *Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeA                   `optional,ext-1`
	Dl_SchedulingOffset_PDSCH_TypeB                   *Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeB                   `optional,ext-1`
	Ul_SchedulingOffset                               *Phy_ParametersFRX_Diff_ul_SchedulingOffset                               `optional,ext-1`
	Dl_64QAM_MCS_TableAlt                             *Phy_ParametersFRX_Diff_dl_64QAM_MCS_TableAlt                             `optional,ext-1`
	Ul_64QAM_MCS_TableAlt                             *Phy_ParametersFRX_Diff_ul_64QAM_MCS_TableAlt                             `optional,ext-1`
	Cqi_TableAlt                                      *Phy_ParametersFRX_Diff_cqi_TableAlt                                      `optional,ext-1`
	OneFL_DMRS_TwoAdditionalDMRS_UL                   *Phy_ParametersFRX_Diff_oneFL_DMRS_TwoAdditionalDMRS_UL                   `optional,ext-1`
	TwoFL_DMRS_TwoAdditionalDMRS_UL                   *Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL                   `optional,ext-1`
	OneFL_DMRS_ThreeAdditionalDMRS_UL                 *Phy_ParametersFRX_Diff_oneFL_DMRS_ThreeAdditionalDMRS_UL                 `optional,ext-1`
	Pdcch_BlindDetectionNRDC                          *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC                          `lb:1,ub:15,optional,ext-2`
	Mux_HARQ_ACK_PUSCH_DiffSymbol                     *Phy_ParametersFRX_Diff_mux_HARQ_ACK_PUSCH_DiffSymbol                     `optional,ext-2`
	Type1_HARQ_ACK_Codebook_r16                       *Phy_ParametersFRX_Diff_type1_HARQ_ACK_Codebook_r16                       `optional,ext-3`
	EnhancedPowerControl_r16                          *Phy_ParametersFRX_Diff_enhancedPowerControl_r16                          `optional,ext-3`
	SimultaneousTCI_ActMultipleCC_r16                 *Phy_ParametersFRX_Diff_simultaneousTCI_ActMultipleCC_r16                 `optional,ext-3`
	SimultaneousSpatialRelationMultipleCC_r16         *Phy_ParametersFRX_Diff_simultaneousSpatialRelationMultipleCC_r16         `optional,ext-3`
	Cli_RSSI_FDM_DL_r16                               *Phy_ParametersFRX_Diff_cli_RSSI_FDM_DL_r16                               `optional,ext-3`
	Cli_SRS_RSRP_FDM_DL_r16                           *Phy_ParametersFRX_Diff_cli_SRS_RSRP_FDM_DL_r16                           `optional,ext-3`
	MaxLayersMIMO_Adaptation_r16                      *Phy_ParametersFRX_Diff_maxLayersMIMO_Adaptation_r16                      `optional,ext-3`
	AggregationFactorSPS_DL_r16                       *Phy_ParametersFRX_Diff_aggregationFactorSPS_DL_r16                       `optional,ext-3`
	MaxTotalResourcesForOneFreqRange_r16              *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16              `optional,ext-3`
	Csi_ReportFrameworkExt_r16                        *CSI_ReportFrameworkExt_r16                                               `optional,ext-3`
	TwoTCI_Act_servingCellInCC_List_r16               *Phy_ParametersFRX_Diff_twoTCI_Act_servingCellInCC_List_r16               `optional,ext-4`
	Cri_RI_CQI_WithoutNon_PMI_PortInd_r16             *Phy_ParametersFRX_Diff_cri_RI_CQI_WithoutNon_PMI_PortInd_r16             `optional,ext-5`
	Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 *Phy_ParametersFRX_Diff_cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 `optional,ext-6`
}

func (ie *Phy_ParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil || ie.Mux_SR_HARQ_ACK_PUCCH != nil || ie.Mux_MultipleGroupCtrlCH_Overlap != nil || ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil || ie.Dl_64QAM_MCS_TableAlt != nil || ie.Ul_64QAM_MCS_TableAlt != nil || ie.Cqi_TableAlt != nil || ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil || ie.Pdcch_BlindDetectionNRDC != nil || ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil || ie.Type1_HARQ_ACK_Codebook_r16 != nil || ie.EnhancedPowerControl_r16 != nil || ie.SimultaneousTCI_ActMultipleCC_r16 != nil || ie.SimultaneousSpatialRelationMultipleCC_r16 != nil || ie.Cli_RSSI_FDM_DL_r16 != nil || ie.Cli_SRS_RSRP_FDM_DL_r16 != nil || ie.MaxLayersMIMO_Adaptation_r16 != nil || ie.AggregationFactorSPS_DL_r16 != nil || ie.MaxTotalResourcesForOneFreqRange_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil || ie.TwoTCI_Act_servingCellInCC_List_r16 != nil || ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil || ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil
	preambleBits := []bool{hasExtensions, ie.DynamicSFI != nil, ie.Dummy1 != nil, ie.TwoFL_DMRS != nil, ie.Dummy2 != nil, ie.Dummy3 != nil, ie.SupportedDMRS_TypeDL != nil, ie.SupportedDMRS_TypeUL != nil, ie.SemiOpenLoopCSI != nil, ie.Csi_ReportWithoutPMI != nil, ie.Csi_ReportWithoutCQI != nil, ie.OnePortsPTRS != nil, ie.TwoPUCCH_F0_2_ConsecSymbols != nil, ie.Pucch_F2_WithFH != nil, ie.Pucch_F3_WithFH != nil, ie.Pucch_F4_WithFH != nil, ie.Pucch_F0_2WithoutFH != nil, ie.Pucch_F1_3_4WithoutFH != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil, ie.Uci_CodeBlockSegmentation != nil, ie.OnePUCCH_LongAndShortFormat != nil, ie.TwoPUCCH_AnyOthersInSlot != nil, ie.IntraSlotFreqHopping_PUSCH != nil, ie.Pusch_LBRM != nil, ie.Pdcch_BlindDetectionCA != nil, ie.Tpc_PUSCH_RNTI != nil, ie.Tpc_PUCCH_RNTI != nil, ie.Tpc_SRS_RNTI != nil, ie.AbsoluteTPC_Command != nil, ie.TwoDifferentTPC_Loop_PUSCH != nil, ie.TwoDifferentTPC_Loop_PUCCH != nil, ie.Pusch_HalfPi_BPSK != nil, ie.Pucch_F3_4_HalfPi_BPSK != nil, ie.AlmostContiguousCP_OFDM_UL != nil, ie.Sp_CSI_RS != nil, ie.Sp_CSI_IM != nil, ie.Tdd_MultiDL_UL_SwitchPerSlot != nil, ie.MultipleCORESET != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DynamicSFI != nil {
		if err = ie.DynamicSFI.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSFI", err)
		}
	}
	if ie.Dummy1 != nil {
		if err = w.WriteBitString(ie.Dummy1.Bytes, uint(ie.Dummy1.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode Dummy1", err)
		}
	}
	if ie.TwoFL_DMRS != nil {
		if err = w.WriteBitString(ie.TwoFL_DMRS.Bytes, uint(ie.TwoFL_DMRS.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode TwoFL_DMRS", err)
		}
	}
	if ie.Dummy2 != nil {
		if err = w.WriteBitString(ie.Dummy2.Bytes, uint(ie.Dummy2.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode Dummy2", err)
		}
	}
	if ie.Dummy3 != nil {
		if err = w.WriteBitString(ie.Dummy3.Bytes, uint(ie.Dummy3.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode Dummy3", err)
		}
	}
	if ie.SupportedDMRS_TypeDL != nil {
		if err = ie.SupportedDMRS_TypeDL.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedDMRS_TypeDL", err)
		}
	}
	if ie.SupportedDMRS_TypeUL != nil {
		if err = ie.SupportedDMRS_TypeUL.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedDMRS_TypeUL", err)
		}
	}
	if ie.SemiOpenLoopCSI != nil {
		if err = ie.SemiOpenLoopCSI.Encode(w); err != nil {
			return utils.WrapError("Encode SemiOpenLoopCSI", err)
		}
	}
	if ie.Csi_ReportWithoutPMI != nil {
		if err = ie.Csi_ReportWithoutPMI.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportWithoutPMI", err)
		}
	}
	if ie.Csi_ReportWithoutCQI != nil {
		if err = ie.Csi_ReportWithoutCQI.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_ReportWithoutCQI", err)
		}
	}
	if ie.OnePortsPTRS != nil {
		if err = w.WriteBitString(ie.OnePortsPTRS.Bytes, uint(ie.OnePortsPTRS.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode OnePortsPTRS", err)
		}
	}
	if ie.TwoPUCCH_F0_2_ConsecSymbols != nil {
		if err = ie.TwoPUCCH_F0_2_ConsecSymbols.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if ie.Pucch_F2_WithFH != nil {
		if err = ie.Pucch_F2_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F2_WithFH", err)
		}
	}
	if ie.Pucch_F3_WithFH != nil {
		if err = ie.Pucch_F3_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F3_WithFH", err)
		}
	}
	if ie.Pucch_F4_WithFH != nil {
		if err = ie.Pucch_F4_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F4_WithFH", err)
		}
	}
	if ie.Pucch_F0_2WithoutFH != nil {
		if err = ie.Pucch_F0_2WithoutFH.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F0_2WithoutFH", err)
		}
	}
	if ie.Pucch_F1_3_4WithoutFH != nil {
		if err = ie.Pucch_F1_3_4WithoutFH.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F1_3_4WithoutFH", err)
		}
	}
	if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil {
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot", err)
		}
	}
	if ie.Uci_CodeBlockSegmentation != nil {
		if err = ie.Uci_CodeBlockSegmentation.Encode(w); err != nil {
			return utils.WrapError("Encode Uci_CodeBlockSegmentation", err)
		}
	}
	if ie.OnePUCCH_LongAndShortFormat != nil {
		if err = ie.OnePUCCH_LongAndShortFormat.Encode(w); err != nil {
			return utils.WrapError("Encode OnePUCCH_LongAndShortFormat", err)
		}
	}
	if ie.TwoPUCCH_AnyOthersInSlot != nil {
		if err = ie.TwoPUCCH_AnyOthersInSlot.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_AnyOthersInSlot", err)
		}
	}
	if ie.IntraSlotFreqHopping_PUSCH != nil {
		if err = ie.IntraSlotFreqHopping_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode IntraSlotFreqHopping_PUSCH", err)
		}
	}
	if ie.Pusch_LBRM != nil {
		if err = ie.Pusch_LBRM.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_LBRM", err)
		}
	}
	if ie.Pdcch_BlindDetectionCA != nil {
		if err = w.WriteInteger(*ie.Pdcch_BlindDetectionCA, &uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA", err)
		}
	}
	if ie.Tpc_PUSCH_RNTI != nil {
		if err = ie.Tpc_PUSCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUSCH_RNTI", err)
		}
	}
	if ie.Tpc_PUCCH_RNTI != nil {
		if err = ie.Tpc_PUCCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_PUCCH_RNTI", err)
		}
	}
	if ie.Tpc_SRS_RNTI != nil {
		if err = ie.Tpc_SRS_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode Tpc_SRS_RNTI", err)
		}
	}
	if ie.AbsoluteTPC_Command != nil {
		if err = ie.AbsoluteTPC_Command.Encode(w); err != nil {
			return utils.WrapError("Encode AbsoluteTPC_Command", err)
		}
	}
	if ie.TwoDifferentTPC_Loop_PUSCH != nil {
		if err = ie.TwoDifferentTPC_Loop_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode TwoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if ie.TwoDifferentTPC_Loop_PUCCH != nil {
		if err = ie.TwoDifferentTPC_Loop_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode TwoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if ie.Pusch_HalfPi_BPSK != nil {
		if err = ie.Pusch_HalfPi_BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_HalfPi_BPSK", err)
		}
	}
	if ie.Pucch_F3_4_HalfPi_BPSK != nil {
		if err = ie.Pucch_F3_4_HalfPi_BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_F3_4_HalfPi_BPSK", err)
		}
	}
	if ie.AlmostContiguousCP_OFDM_UL != nil {
		if err = ie.AlmostContiguousCP_OFDM_UL.Encode(w); err != nil {
			return utils.WrapError("Encode AlmostContiguousCP_OFDM_UL", err)
		}
	}
	if ie.Sp_CSI_RS != nil {
		if err = ie.Sp_CSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_RS", err)
		}
	}
	if ie.Sp_CSI_IM != nil {
		if err = ie.Sp_CSI_IM.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_IM", err)
		}
	}
	if ie.Tdd_MultiDL_UL_SwitchPerSlot != nil {
		if err = ie.Tdd_MultiDL_UL_SwitchPerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_MultiDL_UL_SwitchPerSlot", err)
		}
	}
	if ie.MultipleCORESET != nil {
		if err = ie.MultipleCORESET.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleCORESET", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 6 bits for 6 extension groups
		extBitmap := []bool{ie.Csi_RS_IM_ReceptionForFeedback != nil || ie.Csi_RS_ProcFrameworkForSRS != nil || ie.Csi_ReportFramework != nil || ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil || ie.Mux_SR_HARQ_ACK_PUCCH != nil || ie.Mux_MultipleGroupCtrlCH_Overlap != nil || ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil || ie.Dl_64QAM_MCS_TableAlt != nil || ie.Ul_64QAM_MCS_TableAlt != nil || ie.Cqi_TableAlt != nil || ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil, ie.Pdcch_BlindDetectionNRDC != nil || ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil, ie.Type1_HARQ_ACK_Codebook_r16 != nil || ie.EnhancedPowerControl_r16 != nil || ie.SimultaneousTCI_ActMultipleCC_r16 != nil || ie.SimultaneousSpatialRelationMultipleCC_r16 != nil || ie.Cli_RSSI_FDM_DL_r16 != nil || ie.Cli_SRS_RSRP_FDM_DL_r16 != nil || ie.MaxLayersMIMO_Adaptation_r16 != nil || ie.AggregationFactorSPS_DL_r16 != nil || ie.MaxTotalResourcesForOneFreqRange_r16 != nil || ie.Csi_ReportFrameworkExt_r16 != nil, ie.TwoTCI_Act_servingCellInCC_List_r16 != nil, ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil, ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersFRX_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Csi_RS_IM_ReceptionForFeedback != nil, ie.Csi_RS_ProcFrameworkForSRS != nil, ie.Csi_ReportFramework != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil, ie.Mux_SR_HARQ_ACK_PUCCH != nil, ie.Mux_MultipleGroupCtrlCH_Overlap != nil, ie.Dl_SchedulingOffset_PDSCH_TypeA != nil, ie.Dl_SchedulingOffset_PDSCH_TypeB != nil, ie.Ul_SchedulingOffset != nil, ie.Dl_64QAM_MCS_TableAlt != nil, ie.Ul_64QAM_MCS_TableAlt != nil, ie.Cqi_TableAlt != nil, ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Csi_RS_IM_ReceptionForFeedback optional
			if ie.Csi_RS_IM_ReceptionForFeedback != nil {
				if err = ie.Csi_RS_IM_ReceptionForFeedback.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// encode Csi_RS_ProcFrameworkForSRS optional
			if ie.Csi_RS_ProcFrameworkForSRS != nil {
				if err = ie.Csi_RS_ProcFrameworkForSRS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// encode Csi_ReportFramework optional
			if ie.Csi_ReportFramework != nil {
				if err = ie.Csi_ReportFramework.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportFramework", err)
				}
			}
			// encode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot optional
			if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil {
				if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot", err)
				}
			}
			// encode Mux_SR_HARQ_ACK_PUCCH optional
			if ie.Mux_SR_HARQ_ACK_PUCCH != nil {
				if err = ie.Mux_SR_HARQ_ACK_PUCCH.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_SR_HARQ_ACK_PUCCH", err)
				}
			}
			// encode Mux_MultipleGroupCtrlCH_Overlap optional
			if ie.Mux_MultipleGroupCtrlCH_Overlap != nil {
				if err = ie.Mux_MultipleGroupCtrlCH_Overlap.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_MultipleGroupCtrlCH_Overlap", err)
				}
			}
			// encode Dl_SchedulingOffset_PDSCH_TypeA optional
			if ie.Dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// encode Dl_SchedulingOffset_PDSCH_TypeB optional
			if ie.Dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeB.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// encode Ul_SchedulingOffset optional
			if ie.Ul_SchedulingOffset != nil {
				if err = ie.Ul_SchedulingOffset.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_SchedulingOffset", err)
				}
			}
			// encode Dl_64QAM_MCS_TableAlt optional
			if ie.Dl_64QAM_MCS_TableAlt != nil {
				if err = ie.Dl_64QAM_MCS_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_64QAM_MCS_TableAlt", err)
				}
			}
			// encode Ul_64QAM_MCS_TableAlt optional
			if ie.Ul_64QAM_MCS_TableAlt != nil {
				if err = ie.Ul_64QAM_MCS_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_64QAM_MCS_TableAlt", err)
				}
			}
			// encode Cqi_TableAlt optional
			if ie.Cqi_TableAlt != nil {
				if err = ie.Cqi_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cqi_TableAlt", err)
				}
			}
			// encode OneFL_DMRS_TwoAdditionalDMRS_UL optional
			if ie.OneFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err = ie.OneFL_DMRS_TwoAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OneFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// encode TwoFL_DMRS_TwoAdditionalDMRS_UL optional
			if ie.TwoFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err = ie.TwoFL_DMRS_TwoAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// encode OneFL_DMRS_ThreeAdditionalDMRS_UL optional
			if ie.OneFL_DMRS_ThreeAdditionalDMRS_UL != nil {
				if err = ie.OneFL_DMRS_ThreeAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode OneFL_DMRS_ThreeAdditionalDMRS_UL", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.Pdcch_BlindDetectionNRDC != nil, ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdcch_BlindDetectionNRDC optional
			if ie.Pdcch_BlindDetectionNRDC != nil {
				if err = ie.Pdcch_BlindDetectionNRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdcch_BlindDetectionNRDC", err)
				}
			}
			// encode Mux_HARQ_ACK_PUSCH_DiffSymbol optional
			if ie.Mux_HARQ_ACK_PUSCH_DiffSymbol != nil {
				if err = ie.Mux_HARQ_ACK_PUSCH_DiffSymbol.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mux_HARQ_ACK_PUSCH_DiffSymbol", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Type1_HARQ_ACK_Codebook_r16 != nil, ie.EnhancedPowerControl_r16 != nil, ie.SimultaneousTCI_ActMultipleCC_r16 != nil, ie.SimultaneousSpatialRelationMultipleCC_r16 != nil, ie.Cli_RSSI_FDM_DL_r16 != nil, ie.Cli_SRS_RSRP_FDM_DL_r16 != nil, ie.MaxLayersMIMO_Adaptation_r16 != nil, ie.AggregationFactorSPS_DL_r16 != nil, ie.MaxTotalResourcesForOneFreqRange_r16 != nil, ie.Csi_ReportFrameworkExt_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Type1_HARQ_ACK_Codebook_r16 optional
			if ie.Type1_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.Type1_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Type1_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode EnhancedPowerControl_r16 optional
			if ie.EnhancedPowerControl_r16 != nil {
				if err = ie.EnhancedPowerControl_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EnhancedPowerControl_r16", err)
				}
			}
			// encode SimultaneousTCI_ActMultipleCC_r16 optional
			if ie.SimultaneousTCI_ActMultipleCC_r16 != nil {
				if err = ie.SimultaneousTCI_ActMultipleCC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousTCI_ActMultipleCC_r16", err)
				}
			}
			// encode SimultaneousSpatialRelationMultipleCC_r16 optional
			if ie.SimultaneousSpatialRelationMultipleCC_r16 != nil {
				if err = ie.SimultaneousSpatialRelationMultipleCC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousSpatialRelationMultipleCC_r16", err)
				}
			}
			// encode Cli_RSSI_FDM_DL_r16 optional
			if ie.Cli_RSSI_FDM_DL_r16 != nil {
				if err = ie.Cli_RSSI_FDM_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cli_RSSI_FDM_DL_r16", err)
				}
			}
			// encode Cli_SRS_RSRP_FDM_DL_r16 optional
			if ie.Cli_SRS_RSRP_FDM_DL_r16 != nil {
				if err = ie.Cli_SRS_RSRP_FDM_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cli_SRS_RSRP_FDM_DL_r16", err)
				}
			}
			// encode MaxLayersMIMO_Adaptation_r16 optional
			if ie.MaxLayersMIMO_Adaptation_r16 != nil {
				if err = ie.MaxLayersMIMO_Adaptation_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxLayersMIMO_Adaptation_r16", err)
				}
			}
			// encode AggregationFactorSPS_DL_r16 optional
			if ie.AggregationFactorSPS_DL_r16 != nil {
				if err = ie.AggregationFactorSPS_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AggregationFactorSPS_DL_r16", err)
				}
			}
			// encode MaxTotalResourcesForOneFreqRange_r16 optional
			if ie.MaxTotalResourcesForOneFreqRange_r16 != nil {
				if err = ie.MaxTotalResourcesForOneFreqRange_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxTotalResourcesForOneFreqRange_r16", err)
				}
			}
			// encode Csi_ReportFrameworkExt_r16 optional
			if ie.Csi_ReportFrameworkExt_r16 != nil {
				if err = ie.Csi_ReportFrameworkExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Csi_ReportFrameworkExt_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.TwoTCI_Act_servingCellInCC_List_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode TwoTCI_Act_servingCellInCC_List_r16 optional
			if ie.TwoTCI_Act_servingCellInCC_List_r16 != nil {
				if err = ie.TwoTCI_Act_servingCellInCC_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TwoTCI_Act_servingCellInCC_List_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 optional
			if ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil {
				if err = ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cri_RI_CQI_WithoutNon_PMI_PortInd_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 optional
			if ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *Phy_ParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSFIPresent bool
	if DynamicSFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy1Present bool
	if Dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoFL_DMRSPresent bool
	if TwoFL_DMRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy2Present bool
	if Dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Dummy3Present bool
	if Dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedDMRS_TypeDLPresent bool
	if SupportedDMRS_TypeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedDMRS_TypeULPresent bool
	if SupportedDMRS_TypeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SemiOpenLoopCSIPresent bool
	if SemiOpenLoopCSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ReportWithoutPMIPresent bool
	if Csi_ReportWithoutPMIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_ReportWithoutCQIPresent bool
	if Csi_ReportWithoutCQIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OnePortsPTRSPresent bool
	if OnePortsPTRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_F0_2_ConsecSymbolsPresent bool
	if TwoPUCCH_F0_2_ConsecSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F2_WithFHPresent bool
	if Pucch_F2_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F3_WithFHPresent bool
	if Pucch_F3_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F4_WithFHPresent bool
	if Pucch_F4_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F0_2WithoutFHPresent bool
	if Pucch_F0_2WithoutFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F1_3_4WithoutFHPresent bool
	if Pucch_F1_3_4WithoutFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent bool
	if Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Uci_CodeBlockSegmentationPresent bool
	if Uci_CodeBlockSegmentationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var OnePUCCH_LongAndShortFormatPresent bool
	if OnePUCCH_LongAndShortFormatPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_AnyOthersInSlotPresent bool
	if TwoPUCCH_AnyOthersInSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var IntraSlotFreqHopping_PUSCHPresent bool
	if IntraSlotFreqHopping_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_LBRMPresent bool
	if Pusch_LBRMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionCAPresent bool
	if Pdcch_BlindDetectionCAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUSCH_RNTIPresent bool
	if Tpc_PUSCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_PUCCH_RNTIPresent bool
	if Tpc_PUCCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tpc_SRS_RNTIPresent bool
	if Tpc_SRS_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsoluteTPC_CommandPresent bool
	if AbsoluteTPC_CommandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoDifferentTPC_Loop_PUSCHPresent bool
	if TwoDifferentTPC_Loop_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoDifferentTPC_Loop_PUCCHPresent bool
	if TwoDifferentTPC_Loop_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_HalfPi_BPSKPresent bool
	if Pusch_HalfPi_BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_F3_4_HalfPi_BPSKPresent bool
	if Pucch_F3_4_HalfPi_BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AlmostContiguousCP_OFDM_ULPresent bool
	if AlmostContiguousCP_OFDM_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_RSPresent bool
	if Sp_CSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_IMPresent bool
	if Sp_CSI_IMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_MultiDL_UL_SwitchPerSlotPresent bool
	if Tdd_MultiDL_UL_SwitchPerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MultipleCORESETPresent bool
	if MultipleCORESETPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DynamicSFIPresent {
		ie.DynamicSFI = new(Phy_ParametersFRX_Diff_dynamicSFI)
		if err = ie.DynamicSFI.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSFI", err)
		}
	}
	if Dummy1Present {
		var tmp_bs_Dummy1 []byte
		var n_Dummy1 uint
		if tmp_bs_Dummy1, n_Dummy1, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode Dummy1", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Dummy1,
			NumBits: uint64(n_Dummy1),
		}
		ie.Dummy1 = &tmp_bitstring
	}
	if TwoFL_DMRSPresent {
		var tmp_bs_TwoFL_DMRS []byte
		var n_TwoFL_DMRS uint
		if tmp_bs_TwoFL_DMRS, n_TwoFL_DMRS, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode TwoFL_DMRS", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_TwoFL_DMRS,
			NumBits: uint64(n_TwoFL_DMRS),
		}
		ie.TwoFL_DMRS = &tmp_bitstring
	}
	if Dummy2Present {
		var tmp_bs_Dummy2 []byte
		var n_Dummy2 uint
		if tmp_bs_Dummy2, n_Dummy2, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode Dummy2", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Dummy2,
			NumBits: uint64(n_Dummy2),
		}
		ie.Dummy2 = &tmp_bitstring
	}
	if Dummy3Present {
		var tmp_bs_Dummy3 []byte
		var n_Dummy3 uint
		if tmp_bs_Dummy3, n_Dummy3, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode Dummy3", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Dummy3,
			NumBits: uint64(n_Dummy3),
		}
		ie.Dummy3 = &tmp_bitstring
	}
	if SupportedDMRS_TypeDLPresent {
		ie.SupportedDMRS_TypeDL = new(Phy_ParametersFRX_Diff_supportedDMRS_TypeDL)
		if err = ie.SupportedDMRS_TypeDL.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedDMRS_TypeDL", err)
		}
	}
	if SupportedDMRS_TypeULPresent {
		ie.SupportedDMRS_TypeUL = new(Phy_ParametersFRX_Diff_supportedDMRS_TypeUL)
		if err = ie.SupportedDMRS_TypeUL.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedDMRS_TypeUL", err)
		}
	}
	if SemiOpenLoopCSIPresent {
		ie.SemiOpenLoopCSI = new(Phy_ParametersFRX_Diff_semiOpenLoopCSI)
		if err = ie.SemiOpenLoopCSI.Decode(r); err != nil {
			return utils.WrapError("Decode SemiOpenLoopCSI", err)
		}
	}
	if Csi_ReportWithoutPMIPresent {
		ie.Csi_ReportWithoutPMI = new(Phy_ParametersFRX_Diff_csi_ReportWithoutPMI)
		if err = ie.Csi_ReportWithoutPMI.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_ReportWithoutPMI", err)
		}
	}
	if Csi_ReportWithoutCQIPresent {
		ie.Csi_ReportWithoutCQI = new(Phy_ParametersFRX_Diff_csi_ReportWithoutCQI)
		if err = ie.Csi_ReportWithoutCQI.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_ReportWithoutCQI", err)
		}
	}
	if OnePortsPTRSPresent {
		var tmp_bs_OnePortsPTRS []byte
		var n_OnePortsPTRS uint
		if tmp_bs_OnePortsPTRS, n_OnePortsPTRS, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode OnePortsPTRS", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_OnePortsPTRS,
			NumBits: uint64(n_OnePortsPTRS),
		}
		ie.OnePortsPTRS = &tmp_bitstring
	}
	if TwoPUCCH_F0_2_ConsecSymbolsPresent {
		ie.TwoPUCCH_F0_2_ConsecSymbols = new(Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols)
		if err = ie.TwoPUCCH_F0_2_ConsecSymbols.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if Pucch_F2_WithFHPresent {
		ie.Pucch_F2_WithFH = new(Phy_ParametersFRX_Diff_pucch_F2_WithFH)
		if err = ie.Pucch_F2_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F2_WithFH", err)
		}
	}
	if Pucch_F3_WithFHPresent {
		ie.Pucch_F3_WithFH = new(Phy_ParametersFRX_Diff_pucch_F3_WithFH)
		if err = ie.Pucch_F3_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F3_WithFH", err)
		}
	}
	if Pucch_F4_WithFHPresent {
		ie.Pucch_F4_WithFH = new(Phy_ParametersFRX_Diff_pucch_F4_WithFH)
		if err = ie.Pucch_F4_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F4_WithFH", err)
		}
	}
	if Pucch_F0_2WithoutFHPresent {
		ie.Pucch_F0_2WithoutFH = new(Phy_ParametersFRX_Diff_pucch_F0_2WithoutFH)
		if err = ie.Pucch_F0_2WithoutFH.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F0_2WithoutFH", err)
		}
	}
	if Pucch_F1_3_4WithoutFHPresent {
		ie.Pucch_F1_3_4WithoutFH = new(Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH)
		if err = ie.Pucch_F1_3_4WithoutFH.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F1_3_4WithoutFH", err)
		}
	}
	if Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent {
		ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot)
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot", err)
		}
	}
	if Uci_CodeBlockSegmentationPresent {
		ie.Uci_CodeBlockSegmentation = new(Phy_ParametersFRX_Diff_uci_CodeBlockSegmentation)
		if err = ie.Uci_CodeBlockSegmentation.Decode(r); err != nil {
			return utils.WrapError("Decode Uci_CodeBlockSegmentation", err)
		}
	}
	if OnePUCCH_LongAndShortFormatPresent {
		ie.OnePUCCH_LongAndShortFormat = new(Phy_ParametersFRX_Diff_onePUCCH_LongAndShortFormat)
		if err = ie.OnePUCCH_LongAndShortFormat.Decode(r); err != nil {
			return utils.WrapError("Decode OnePUCCH_LongAndShortFormat", err)
		}
	}
	if TwoPUCCH_AnyOthersInSlotPresent {
		ie.TwoPUCCH_AnyOthersInSlot = new(Phy_ParametersFRX_Diff_twoPUCCH_AnyOthersInSlot)
		if err = ie.TwoPUCCH_AnyOthersInSlot.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_AnyOthersInSlot", err)
		}
	}
	if IntraSlotFreqHopping_PUSCHPresent {
		ie.IntraSlotFreqHopping_PUSCH = new(Phy_ParametersFRX_Diff_intraSlotFreqHopping_PUSCH)
		if err = ie.IntraSlotFreqHopping_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode IntraSlotFreqHopping_PUSCH", err)
		}
	}
	if Pusch_LBRMPresent {
		ie.Pusch_LBRM = new(Phy_ParametersFRX_Diff_pusch_LBRM)
		if err = ie.Pusch_LBRM.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_LBRM", err)
		}
	}
	if Pdcch_BlindDetectionCAPresent {
		var tmp_int_Pdcch_BlindDetectionCA int64
		if tmp_int_Pdcch_BlindDetectionCA, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA", err)
		}
		ie.Pdcch_BlindDetectionCA = &tmp_int_Pdcch_BlindDetectionCA
	}
	if Tpc_PUSCH_RNTIPresent {
		ie.Tpc_PUSCH_RNTI = new(Phy_ParametersFRX_Diff_tpc_PUSCH_RNTI)
		if err = ie.Tpc_PUSCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUSCH_RNTI", err)
		}
	}
	if Tpc_PUCCH_RNTIPresent {
		ie.Tpc_PUCCH_RNTI = new(Phy_ParametersFRX_Diff_tpc_PUCCH_RNTI)
		if err = ie.Tpc_PUCCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_PUCCH_RNTI", err)
		}
	}
	if Tpc_SRS_RNTIPresent {
		ie.Tpc_SRS_RNTI = new(Phy_ParametersFRX_Diff_tpc_SRS_RNTI)
		if err = ie.Tpc_SRS_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode Tpc_SRS_RNTI", err)
		}
	}
	if AbsoluteTPC_CommandPresent {
		ie.AbsoluteTPC_Command = new(Phy_ParametersFRX_Diff_absoluteTPC_Command)
		if err = ie.AbsoluteTPC_Command.Decode(r); err != nil {
			return utils.WrapError("Decode AbsoluteTPC_Command", err)
		}
	}
	if TwoDifferentTPC_Loop_PUSCHPresent {
		ie.TwoDifferentTPC_Loop_PUSCH = new(Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUSCH)
		if err = ie.TwoDifferentTPC_Loop_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode TwoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if TwoDifferentTPC_Loop_PUCCHPresent {
		ie.TwoDifferentTPC_Loop_PUCCH = new(Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUCCH)
		if err = ie.TwoDifferentTPC_Loop_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode TwoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if Pusch_HalfPi_BPSKPresent {
		ie.Pusch_HalfPi_BPSK = new(Phy_ParametersFRX_Diff_pusch_HalfPi_BPSK)
		if err = ie.Pusch_HalfPi_BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_HalfPi_BPSK", err)
		}
	}
	if Pucch_F3_4_HalfPi_BPSKPresent {
		ie.Pucch_F3_4_HalfPi_BPSK = new(Phy_ParametersFRX_Diff_pucch_F3_4_HalfPi_BPSK)
		if err = ie.Pucch_F3_4_HalfPi_BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_F3_4_HalfPi_BPSK", err)
		}
	}
	if AlmostContiguousCP_OFDM_ULPresent {
		ie.AlmostContiguousCP_OFDM_UL = new(Phy_ParametersFRX_Diff_almostContiguousCP_OFDM_UL)
		if err = ie.AlmostContiguousCP_OFDM_UL.Decode(r); err != nil {
			return utils.WrapError("Decode AlmostContiguousCP_OFDM_UL", err)
		}
	}
	if Sp_CSI_RSPresent {
		ie.Sp_CSI_RS = new(Phy_ParametersFRX_Diff_sp_CSI_RS)
		if err = ie.Sp_CSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_RS", err)
		}
	}
	if Sp_CSI_IMPresent {
		ie.Sp_CSI_IM = new(Phy_ParametersFRX_Diff_sp_CSI_IM)
		if err = ie.Sp_CSI_IM.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_IM", err)
		}
	}
	if Tdd_MultiDL_UL_SwitchPerSlotPresent {
		ie.Tdd_MultiDL_UL_SwitchPerSlot = new(Phy_ParametersFRX_Diff_tdd_MultiDL_UL_SwitchPerSlot)
		if err = ie.Tdd_MultiDL_UL_SwitchPerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_MultiDL_UL_SwitchPerSlot", err)
		}
	}
	if MultipleCORESETPresent {
		ie.MultipleCORESET = new(Phy_ParametersFRX_Diff_multipleCORESET)
		if err = ie.MultipleCORESET.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleCORESET", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 6 bits for 6 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Csi_RS_IM_ReceptionForFeedbackPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_RS_ProcFrameworkForSRSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_ReportFrameworkPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlotPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mux_SR_HARQ_ACK_PUCCHPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mux_MultipleGroupCtrlCH_OverlapPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_SchedulingOffset_PDSCH_TypeAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_SchedulingOffset_PDSCH_TypeBPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_SchedulingOffsetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_64QAM_MCS_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_64QAM_MCS_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cqi_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OneFL_DMRS_TwoAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TwoFL_DMRS_TwoAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			OneFL_DMRS_ThreeAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Csi_RS_IM_ReceptionForFeedback optional
			if Csi_RS_IM_ReceptionForFeedbackPresent {
				ie.Csi_RS_IM_ReceptionForFeedback = new(CSI_RS_IM_ReceptionForFeedback)
				if err = ie.Csi_RS_IM_ReceptionForFeedback.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// decode Csi_RS_ProcFrameworkForSRS optional
			if Csi_RS_ProcFrameworkForSRSPresent {
				ie.Csi_RS_ProcFrameworkForSRS = new(CSI_RS_ProcFrameworkForSRS)
				if err = ie.Csi_RS_ProcFrameworkForSRS.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// decode Csi_ReportFramework optional
			if Csi_ReportFrameworkPresent {
				ie.Csi_ReportFramework = new(CSI_ReportFramework)
				if err = ie.Csi_ReportFramework.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportFramework", err)
				}
			}
			// decode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot optional
			if Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlotPresent {
				ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot)
				if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot", err)
				}
			}
			// decode Mux_SR_HARQ_ACK_PUCCH optional
			if Mux_SR_HARQ_ACK_PUCCHPresent {
				ie.Mux_SR_HARQ_ACK_PUCCH = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_PUCCH)
				if err = ie.Mux_SR_HARQ_ACK_PUCCH.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_SR_HARQ_ACK_PUCCH", err)
				}
			}
			// decode Mux_MultipleGroupCtrlCH_Overlap optional
			if Mux_MultipleGroupCtrlCH_OverlapPresent {
				ie.Mux_MultipleGroupCtrlCH_Overlap = new(Phy_ParametersFRX_Diff_mux_MultipleGroupCtrlCH_Overlap)
				if err = ie.Mux_MultipleGroupCtrlCH_Overlap.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_MultipleGroupCtrlCH_Overlap", err)
				}
			}
			// decode Dl_SchedulingOffset_PDSCH_TypeA optional
			if Dl_SchedulingOffset_PDSCH_TypeAPresent {
				ie.Dl_SchedulingOffset_PDSCH_TypeA = new(Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeA)
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeA.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// decode Dl_SchedulingOffset_PDSCH_TypeB optional
			if Dl_SchedulingOffset_PDSCH_TypeBPresent {
				ie.Dl_SchedulingOffset_PDSCH_TypeB = new(Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeB)
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeB.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// decode Ul_SchedulingOffset optional
			if Ul_SchedulingOffsetPresent {
				ie.Ul_SchedulingOffset = new(Phy_ParametersFRX_Diff_ul_SchedulingOffset)
				if err = ie.Ul_SchedulingOffset.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_SchedulingOffset", err)
				}
			}
			// decode Dl_64QAM_MCS_TableAlt optional
			if Dl_64QAM_MCS_TableAltPresent {
				ie.Dl_64QAM_MCS_TableAlt = new(Phy_ParametersFRX_Diff_dl_64QAM_MCS_TableAlt)
				if err = ie.Dl_64QAM_MCS_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_64QAM_MCS_TableAlt", err)
				}
			}
			// decode Ul_64QAM_MCS_TableAlt optional
			if Ul_64QAM_MCS_TableAltPresent {
				ie.Ul_64QAM_MCS_TableAlt = new(Phy_ParametersFRX_Diff_ul_64QAM_MCS_TableAlt)
				if err = ie.Ul_64QAM_MCS_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_64QAM_MCS_TableAlt", err)
				}
			}
			// decode Cqi_TableAlt optional
			if Cqi_TableAltPresent {
				ie.Cqi_TableAlt = new(Phy_ParametersFRX_Diff_cqi_TableAlt)
				if err = ie.Cqi_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cqi_TableAlt", err)
				}
			}
			// decode OneFL_DMRS_TwoAdditionalDMRS_UL optional
			if OneFL_DMRS_TwoAdditionalDMRS_ULPresent {
				ie.OneFL_DMRS_TwoAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_oneFL_DMRS_TwoAdditionalDMRS_UL)
				if err = ie.OneFL_DMRS_TwoAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode OneFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// decode TwoFL_DMRS_TwoAdditionalDMRS_UL optional
			if TwoFL_DMRS_TwoAdditionalDMRS_ULPresent {
				ie.TwoFL_DMRS_TwoAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL)
				if err = ie.TwoFL_DMRS_TwoAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// decode OneFL_DMRS_ThreeAdditionalDMRS_UL optional
			if OneFL_DMRS_ThreeAdditionalDMRS_ULPresent {
				ie.OneFL_DMRS_ThreeAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_oneFL_DMRS_ThreeAdditionalDMRS_UL)
				if err = ie.OneFL_DMRS_ThreeAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode OneFL_DMRS_ThreeAdditionalDMRS_UL", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdcch_BlindDetectionNRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mux_HARQ_ACK_PUSCH_DiffSymbolPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdcch_BlindDetectionNRDC optional
			if Pdcch_BlindDetectionNRDCPresent {
				ie.Pdcch_BlindDetectionNRDC = new(Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC)
				if err = ie.Pdcch_BlindDetectionNRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdcch_BlindDetectionNRDC", err)
				}
			}
			// decode Mux_HARQ_ACK_PUSCH_DiffSymbol optional
			if Mux_HARQ_ACK_PUSCH_DiffSymbolPresent {
				ie.Mux_HARQ_ACK_PUSCH_DiffSymbol = new(Phy_ParametersFRX_Diff_mux_HARQ_ACK_PUSCH_DiffSymbol)
				if err = ie.Mux_HARQ_ACK_PUSCH_DiffSymbol.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mux_HARQ_ACK_PUSCH_DiffSymbol", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Type1_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			EnhancedPowerControl_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousTCI_ActMultipleCC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousSpatialRelationMultipleCC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cli_RSSI_FDM_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cli_SRS_RSRP_FDM_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxLayersMIMO_Adaptation_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AggregationFactorSPS_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxTotalResourcesForOneFreqRange_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Csi_ReportFrameworkExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Type1_HARQ_ACK_Codebook_r16 optional
			if Type1_HARQ_ACK_Codebook_r16Present {
				ie.Type1_HARQ_ACK_Codebook_r16 = new(Phy_ParametersFRX_Diff_type1_HARQ_ACK_Codebook_r16)
				if err = ie.Type1_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Type1_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode EnhancedPowerControl_r16 optional
			if EnhancedPowerControl_r16Present {
				ie.EnhancedPowerControl_r16 = new(Phy_ParametersFRX_Diff_enhancedPowerControl_r16)
				if err = ie.EnhancedPowerControl_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode EnhancedPowerControl_r16", err)
				}
			}
			// decode SimultaneousTCI_ActMultipleCC_r16 optional
			if SimultaneousTCI_ActMultipleCC_r16Present {
				ie.SimultaneousTCI_ActMultipleCC_r16 = new(Phy_ParametersFRX_Diff_simultaneousTCI_ActMultipleCC_r16)
				if err = ie.SimultaneousTCI_ActMultipleCC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousTCI_ActMultipleCC_r16", err)
				}
			}
			// decode SimultaneousSpatialRelationMultipleCC_r16 optional
			if SimultaneousSpatialRelationMultipleCC_r16Present {
				ie.SimultaneousSpatialRelationMultipleCC_r16 = new(Phy_ParametersFRX_Diff_simultaneousSpatialRelationMultipleCC_r16)
				if err = ie.SimultaneousSpatialRelationMultipleCC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousSpatialRelationMultipleCC_r16", err)
				}
			}
			// decode Cli_RSSI_FDM_DL_r16 optional
			if Cli_RSSI_FDM_DL_r16Present {
				ie.Cli_RSSI_FDM_DL_r16 = new(Phy_ParametersFRX_Diff_cli_RSSI_FDM_DL_r16)
				if err = ie.Cli_RSSI_FDM_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cli_RSSI_FDM_DL_r16", err)
				}
			}
			// decode Cli_SRS_RSRP_FDM_DL_r16 optional
			if Cli_SRS_RSRP_FDM_DL_r16Present {
				ie.Cli_SRS_RSRP_FDM_DL_r16 = new(Phy_ParametersFRX_Diff_cli_SRS_RSRP_FDM_DL_r16)
				if err = ie.Cli_SRS_RSRP_FDM_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cli_SRS_RSRP_FDM_DL_r16", err)
				}
			}
			// decode MaxLayersMIMO_Adaptation_r16 optional
			if MaxLayersMIMO_Adaptation_r16Present {
				ie.MaxLayersMIMO_Adaptation_r16 = new(Phy_ParametersFRX_Diff_maxLayersMIMO_Adaptation_r16)
				if err = ie.MaxLayersMIMO_Adaptation_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxLayersMIMO_Adaptation_r16", err)
				}
			}
			// decode AggregationFactorSPS_DL_r16 optional
			if AggregationFactorSPS_DL_r16Present {
				ie.AggregationFactorSPS_DL_r16 = new(Phy_ParametersFRX_Diff_aggregationFactorSPS_DL_r16)
				if err = ie.AggregationFactorSPS_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode AggregationFactorSPS_DL_r16", err)
				}
			}
			// decode MaxTotalResourcesForOneFreqRange_r16 optional
			if MaxTotalResourcesForOneFreqRange_r16Present {
				ie.MaxTotalResourcesForOneFreqRange_r16 = new(Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16)
				if err = ie.MaxTotalResourcesForOneFreqRange_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxTotalResourcesForOneFreqRange_r16", err)
				}
			}
			// decode Csi_ReportFrameworkExt_r16 optional
			if Csi_ReportFrameworkExt_r16Present {
				ie.Csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
				if err = ie.Csi_ReportFrameworkExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Csi_ReportFrameworkExt_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			TwoTCI_Act_servingCellInCC_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode TwoTCI_Act_servingCellInCC_List_r16 optional
			if TwoTCI_Act_servingCellInCC_List_r16Present {
				ie.TwoTCI_Act_servingCellInCC_List_r16 = new(Phy_ParametersFRX_Diff_twoTCI_Act_servingCellInCC_List_r16)
				if err = ie.TwoTCI_Act_servingCellInCC_List_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode TwoTCI_Act_servingCellInCC_List_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Cri_RI_CQI_WithoutNon_PMI_PortInd_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 optional
			if Cri_RI_CQI_WithoutNon_PMI_PortInd_r16Present {
				ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16 = new(Phy_ParametersFRX_Diff_cri_RI_CQI_WithoutNon_PMI_PortInd_r16)
				if err = ie.Cri_RI_CQI_WithoutNon_PMI_PortInd_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cri_RI_CQI_WithoutNon_PMI_PortInd_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 optional
			if Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17Present {
				ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 = new(Phy_ParametersFRX_Diff_cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17)
				if err = ie.Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
		}
	}
	return nil
}
