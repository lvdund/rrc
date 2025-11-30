package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16 struct {
	Pdcch_BlindDetectionMCG_UE1_r16 int64 `lb:0,ub:15,madatory`
	Pdcch_BlindDetectionMCG_UE2_r16 int64 `lb:0,ub:15,madatory`
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionMCG_UE1_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionMCG_UE1_r16", err)
	}
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionMCG_UE2_r16, &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionMCG_UE2_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1610_pdcch_BlindDetectionMCG_UE_Mixed_r16) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Pdcch_BlindDetectionMCG_UE1_r16 int64
	if tmp_int_Pdcch_BlindDetectionMCG_UE1_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionMCG_UE1_r16", err)
	}
	ie.Pdcch_BlindDetectionMCG_UE1_r16 = tmp_int_Pdcch_BlindDetectionMCG_UE1_r16
	var tmp_int_Pdcch_BlindDetectionMCG_UE2_r16 int64
	if tmp_int_Pdcch_BlindDetectionMCG_UE2_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionMCG_UE2_r16", err)
	}
	ie.Pdcch_BlindDetectionMCG_UE2_r16 = tmp_int_Pdcch_BlindDetectionMCG_UE2_r16
	return nil
}
