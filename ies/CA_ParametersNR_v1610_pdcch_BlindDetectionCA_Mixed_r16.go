package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16 struct {
	Pdcch_BlindDetectionCA1_r16  int64                                                                               `lb:1,ub:15,madatory`
	Pdcch_BlindDetectionCA2_r16  int64                                                                               `lb:1,ub:15,madatory`
	SupportedSpanArrangement_r16 CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16_supportedSpanArrangement_r16 `madatory`
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionCA1_r16, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionCA1_r16", err)
	}
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionCA2_r16, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionCA2_r16", err)
	}
	if err = ie.SupportedSpanArrangement_r16.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedSpanArrangement_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionCA_Mixed_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Pdcch_BlindDetectionCA1_r16 int64
	if tmp_int_Pdcch_BlindDetectionCA1_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionCA1_r16", err)
	}
	ie.Pdcch_BlindDetectionCA1_r16 = tmp_int_Pdcch_BlindDetectionCA1_r16
	var tmp_int_Pdcch_BlindDetectionCA2_r16 int64
	if tmp_int_Pdcch_BlindDetectionCA2_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionCA2_r16", err)
	}
	ie.Pdcch_BlindDetectionCA2_r16 = tmp_int_Pdcch_BlindDetectionCA2_r16
	if err = ie.SupportedSpanArrangement_r16.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedSpanArrangement_r16", err)
	}
	return nil
}
