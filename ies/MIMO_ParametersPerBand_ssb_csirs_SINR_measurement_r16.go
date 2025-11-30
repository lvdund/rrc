package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16 struct {
	MaxNumberSSB_CSIRS_OneTx_CMR_r16    MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_OneTx_CMR_r16    `madatory`
	MaxNumberCSI_IM_NZP_IMR_res_r16     MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSI_IM_NZP_IMR_res_r16     `madatory`
	MaxNumberCSIRS_2Tx_res_r16          MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSIRS_2Tx_res_r16          `madatory`
	MaxNumberSSB_CSIRS_res_r16          MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16          `madatory`
	MaxNumberCSI_IM_NZP_IMR_res_mem_r16 MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberCSI_IM_NZP_IMR_res_mem_r16 `madatory`
	SupportedCSI_RS_Density_CMR_r16     MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedCSI_RS_Density_CMR_r16     `madatory`
	MaxNumberAperiodicCSI_RS_Res_r16    MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberAperiodicCSI_RS_Res_r16    `madatory`
	SupportedSINR_meas_r16              *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16             `optional`
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedSINR_meas_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MaxNumberSSB_CSIRS_OneTx_CMR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSSB_CSIRS_OneTx_CMR_r16", err)
	}
	if err = ie.MaxNumberCSI_IM_NZP_IMR_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCSI_IM_NZP_IMR_res_r16", err)
	}
	if err = ie.MaxNumberCSIRS_2Tx_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCSIRS_2Tx_res_r16", err)
	}
	if err = ie.MaxNumberSSB_CSIRS_res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberSSB_CSIRS_res_r16", err)
	}
	if err = ie.MaxNumberCSI_IM_NZP_IMR_res_mem_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberCSI_IM_NZP_IMR_res_mem_r16", err)
	}
	if err = ie.SupportedCSI_RS_Density_CMR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedCSI_RS_Density_CMR_r16", err)
	}
	if err = ie.MaxNumberAperiodicCSI_RS_Res_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MaxNumberAperiodicCSI_RS_Res_r16", err)
	}
	if ie.SupportedSINR_meas_r16 != nil {
		if err = ie.SupportedSINR_meas_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedSINR_meas_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16) Decode(r *aper.AperReader) error {
	var err error
	var SupportedSINR_meas_r16Present bool
	if SupportedSINR_meas_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MaxNumberSSB_CSIRS_OneTx_CMR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSSB_CSIRS_OneTx_CMR_r16", err)
	}
	if err = ie.MaxNumberCSI_IM_NZP_IMR_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCSI_IM_NZP_IMR_res_r16", err)
	}
	if err = ie.MaxNumberCSIRS_2Tx_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCSIRS_2Tx_res_r16", err)
	}
	if err = ie.MaxNumberSSB_CSIRS_res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberSSB_CSIRS_res_r16", err)
	}
	if err = ie.MaxNumberCSI_IM_NZP_IMR_res_mem_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberCSI_IM_NZP_IMR_res_mem_r16", err)
	}
	if err = ie.SupportedCSI_RS_Density_CMR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedCSI_RS_Density_CMR_r16", err)
	}
	if err = ie.MaxNumberAperiodicCSI_RS_Res_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MaxNumberAperiodicCSI_RS_Res_r16", err)
	}
	if SupportedSINR_meas_r16Present {
		ie.SupportedSINR_meas_r16 = new(MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_supportedSINR_meas_r16)
		if err = ie.SupportedSINR_meas_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedSINR_meas_r16", err)
		}
	}
	return nil
}
