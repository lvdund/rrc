package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16_Enum_n8   aper.Enumerated = 0
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16_Enum_n16  aper.Enumerated = 1
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16_Enum_n32  aper.Enumerated = 2
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16_Enum_n64  aper.Enumerated = 3
	MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16_Enum_n128 aper.Enumerated = 4
)

type MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16", err)
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16_maxNumberSSB_CSIRS_res_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
