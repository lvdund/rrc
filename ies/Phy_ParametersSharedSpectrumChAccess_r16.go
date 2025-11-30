package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersSharedSpectrumChAccess_r16 struct {
	Ss_SINR_Meas_r16                           *Phy_ParametersSharedSpectrumChAccess_r16_ss_SINR_Meas_r16                           `optional`
	Sp_CSI_ReportPUCCH_r16                     *Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUCCH_r16                     `optional`
	Sp_CSI_ReportPUSCH_r16                     *Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUSCH_r16                     `optional`
	DynamicSFI_r16                             *Phy_ParametersSharedSpectrumChAccess_r16_dynamicSFI_r16                             `optional`
	Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16  *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16  `optional`
	Mux_SR_HARQ_ACK_PUCCH_r16                  *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_PUCCH_r16                  `optional`
	Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 *Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 `optional`
	Mux_HARQ_ACK_PUSCH_DiffSymbol_r16          *Phy_ParametersSharedSpectrumChAccess_r16_mux_HARQ_ACK_PUSCH_DiffSymbol_r16          `optional`
	Pucch_Repetition_F1_3_4_r16                *Phy_ParametersSharedSpectrumChAccess_r16_pucch_Repetition_F1_3_4_r16                `optional`
	Type1_PUSCH_RepetitionMultiSlots_r16       *Phy_ParametersSharedSpectrumChAccess_r16_type1_PUSCH_RepetitionMultiSlots_r16       `optional`
	Type2_PUSCH_RepetitionMultiSlots_r16       *Phy_ParametersSharedSpectrumChAccess_r16_type2_PUSCH_RepetitionMultiSlots_r16       `optional`
	Pusch_RepetitionMultiSlots_r16             *Phy_ParametersSharedSpectrumChAccess_r16_pusch_RepetitionMultiSlots_r16             `optional`
	Pdsch_RepetitionMultiSlots_r16             *Phy_ParametersSharedSpectrumChAccess_r16_pdsch_RepetitionMultiSlots_r16             `optional`
	DownlinkSPS_r16                            *Phy_ParametersSharedSpectrumChAccess_r16_downlinkSPS_r16                            `optional`
	ConfiguredUL_GrantType1_r16                *Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType1_r16                `optional`
	ConfiguredUL_GrantType2_r16                *Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType2_r16                `optional`
	Pre_EmptIndication_DL_r16                  *Phy_ParametersSharedSpectrumChAccess_r16_pre_EmptIndication_DL_r16                  `optional`
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ss_SINR_Meas_r16 != nil, ie.Sp_CSI_ReportPUCCH_r16 != nil, ie.Sp_CSI_ReportPUSCH_r16 != nil, ie.DynamicSFI_r16 != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil, ie.Mux_SR_HARQ_ACK_PUCCH_r16 != nil, ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil, ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil, ie.Pucch_Repetition_F1_3_4_r16 != nil, ie.Type1_PUSCH_RepetitionMultiSlots_r16 != nil, ie.Type2_PUSCH_RepetitionMultiSlots_r16 != nil, ie.Pusch_RepetitionMultiSlots_r16 != nil, ie.Pdsch_RepetitionMultiSlots_r16 != nil, ie.DownlinkSPS_r16 != nil, ie.ConfiguredUL_GrantType1_r16 != nil, ie.ConfiguredUL_GrantType2_r16 != nil, ie.Pre_EmptIndication_DL_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ss_SINR_Meas_r16 != nil {
		if err = ie.Ss_SINR_Meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_SINR_Meas_r16", err)
		}
	}
	if ie.Sp_CSI_ReportPUCCH_r16 != nil {
		if err = ie.Sp_CSI_ReportPUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportPUCCH_r16", err)
		}
	}
	if ie.Sp_CSI_ReportPUSCH_r16 != nil {
		if err = ie.Sp_CSI_ReportPUSCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sp_CSI_ReportPUSCH_r16", err)
		}
	}
	if ie.DynamicSFI_r16 != nil {
		if err = ie.DynamicSFI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSFI_r16", err)
		}
	}
	if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 != nil {
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16", err)
		}
	}
	if ie.Mux_SR_HARQ_ACK_PUCCH_r16 != nil {
		if err = ie.Mux_SR_HARQ_ACK_PUCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_SR_HARQ_ACK_PUCCH_r16", err)
		}
	}
	if ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 != nil {
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16", err)
		}
	}
	if ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 != nil {
		if err = ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Mux_HARQ_ACK_PUSCH_DiffSymbol_r16", err)
		}
	}
	if ie.Pucch_Repetition_F1_3_4_r16 != nil {
		if err = ie.Pucch_Repetition_F1_3_4_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_Repetition_F1_3_4_r16", err)
		}
	}
	if ie.Type1_PUSCH_RepetitionMultiSlots_r16 != nil {
		if err = ie.Type1_PUSCH_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type1_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.Type2_PUSCH_RepetitionMultiSlots_r16 != nil {
		if err = ie.Type2_PUSCH_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.Pusch_RepetitionMultiSlots_r16 != nil {
		if err = ie.Pusch_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.Pdsch_RepetitionMultiSlots_r16 != nil {
		if err = ie.Pdsch_RepetitionMultiSlots_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_RepetitionMultiSlots_r16", err)
		}
	}
	if ie.DownlinkSPS_r16 != nil {
		if err = ie.DownlinkSPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DownlinkSPS_r16", err)
		}
	}
	if ie.ConfiguredUL_GrantType1_r16 != nil {
		if err = ie.ConfiguredUL_GrantType1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredUL_GrantType1_r16", err)
		}
	}
	if ie.ConfiguredUL_GrantType2_r16 != nil {
		if err = ie.ConfiguredUL_GrantType2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredUL_GrantType2_r16", err)
		}
	}
	if ie.Pre_EmptIndication_DL_r16 != nil {
		if err = ie.Pre_EmptIndication_DL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pre_EmptIndication_DL_r16", err)
		}
	}
	return nil
}

func (ie *Phy_ParametersSharedSpectrumChAccess_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ss_SINR_Meas_r16Present bool
	if Ss_SINR_Meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportPUCCH_r16Present bool
	if Sp_CSI_ReportPUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sp_CSI_ReportPUSCH_r16Present bool
	if Sp_CSI_ReportPUSCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSFI_r16Present bool
	if DynamicSFI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present bool
	if Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_SR_HARQ_ACK_PUCCH_r16Present bool
	if Mux_SR_HARQ_ACK_PUCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present bool
	if Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present bool
	if Mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_Repetition_F1_3_4_r16Present bool
	if Pucch_Repetition_F1_3_4_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type1_PUSCH_RepetitionMultiSlots_r16Present bool
	if Type1_PUSCH_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_PUSCH_RepetitionMultiSlots_r16Present bool
	if Type2_PUSCH_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_RepetitionMultiSlots_r16Present bool
	if Pusch_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_RepetitionMultiSlots_r16Present bool
	if Pdsch_RepetitionMultiSlots_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DownlinkSPS_r16Present bool
	if DownlinkSPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredUL_GrantType1_r16Present bool
	if ConfiguredUL_GrantType1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ConfiguredUL_GrantType2_r16Present bool
	if ConfiguredUL_GrantType2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pre_EmptIndication_DL_r16Present bool
	if Pre_EmptIndication_DL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ss_SINR_Meas_r16Present {
		ie.Ss_SINR_Meas_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_ss_SINR_Meas_r16)
		if err = ie.Ss_SINR_Meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_SINR_Meas_r16", err)
		}
	}
	if Sp_CSI_ReportPUCCH_r16Present {
		ie.Sp_CSI_ReportPUCCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUCCH_r16)
		if err = ie.Sp_CSI_ReportPUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportPUCCH_r16", err)
		}
	}
	if Sp_CSI_ReportPUSCH_r16Present {
		ie.Sp_CSI_ReportPUSCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_sp_CSI_ReportPUSCH_r16)
		if err = ie.Sp_CSI_ReportPUSCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sp_CSI_ReportPUSCH_r16", err)
		}
	}
	if DynamicSFI_r16Present {
		ie.DynamicSFI_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_dynamicSFI_r16)
		if err = ie.DynamicSFI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSFI_r16", err)
		}
	}
	if Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16Present {
		ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16)
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot_r16", err)
		}
	}
	if Mux_SR_HARQ_ACK_PUCCH_r16Present {
		ie.Mux_SR_HARQ_ACK_PUCCH_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_PUCCH_r16)
		if err = ie.Mux_SR_HARQ_ACK_PUCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_SR_HARQ_ACK_PUCCH_r16", err)
		}
	}
	if Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16Present {
		ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16)
		if err = ie.Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot_r16", err)
		}
	}
	if Mux_HARQ_ACK_PUSCH_DiffSymbol_r16Present {
		ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_mux_HARQ_ACK_PUSCH_DiffSymbol_r16)
		if err = ie.Mux_HARQ_ACK_PUSCH_DiffSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Mux_HARQ_ACK_PUSCH_DiffSymbol_r16", err)
		}
	}
	if Pucch_Repetition_F1_3_4_r16Present {
		ie.Pucch_Repetition_F1_3_4_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pucch_Repetition_F1_3_4_r16)
		if err = ie.Pucch_Repetition_F1_3_4_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_Repetition_F1_3_4_r16", err)
		}
	}
	if Type1_PUSCH_RepetitionMultiSlots_r16Present {
		ie.Type1_PUSCH_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_type1_PUSCH_RepetitionMultiSlots_r16)
		if err = ie.Type1_PUSCH_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type1_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if Type2_PUSCH_RepetitionMultiSlots_r16Present {
		ie.Type2_PUSCH_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_type2_PUSCH_RepetitionMultiSlots_r16)
		if err = ie.Type2_PUSCH_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_PUSCH_RepetitionMultiSlots_r16", err)
		}
	}
	if Pusch_RepetitionMultiSlots_r16Present {
		ie.Pusch_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pusch_RepetitionMultiSlots_r16)
		if err = ie.Pusch_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_RepetitionMultiSlots_r16", err)
		}
	}
	if Pdsch_RepetitionMultiSlots_r16Present {
		ie.Pdsch_RepetitionMultiSlots_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pdsch_RepetitionMultiSlots_r16)
		if err = ie.Pdsch_RepetitionMultiSlots_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_RepetitionMultiSlots_r16", err)
		}
	}
	if DownlinkSPS_r16Present {
		ie.DownlinkSPS_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_downlinkSPS_r16)
		if err = ie.DownlinkSPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DownlinkSPS_r16", err)
		}
	}
	if ConfiguredUL_GrantType1_r16Present {
		ie.ConfiguredUL_GrantType1_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType1_r16)
		if err = ie.ConfiguredUL_GrantType1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredUL_GrantType1_r16", err)
		}
	}
	if ConfiguredUL_GrantType2_r16Present {
		ie.ConfiguredUL_GrantType2_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_configuredUL_GrantType2_r16)
		if err = ie.ConfiguredUL_GrantType2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredUL_GrantType2_r16", err)
		}
	}
	if Pre_EmptIndication_DL_r16Present {
		ie.Pre_EmptIndication_DL_r16 = new(Phy_ParametersSharedSpectrumChAccess_r16_pre_EmptIndication_DL_r16)
		if err = ie.Pre_EmptIndication_DL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Pre_EmptIndication_DL_r16", err)
		}
	}
	return nil
}
