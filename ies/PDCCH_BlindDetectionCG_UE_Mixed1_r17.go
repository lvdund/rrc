package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionCG_UE_Mixed1_r17 struct {
	Pdcch_BlindDetectionCG_UE1_r17 int64 `lb:0,ub:15,madatory`
	Pdcch_BlindDetectionCG_UE2_r17 int64 `lb:0,ub:15,madatory`
	Pdcch_BlindDetectionCG_UE3_r17 int64 `lb:0,ub:15,madatory`
}

func (ie *PDCCH_BlindDetectionCG_UE_Mixed1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionCG_UE1_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionCG_UE1_r17", err)
	}
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionCG_UE2_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionCG_UE2_r17", err)
	}
	if err = w.WriteInteger(ie.Pdcch_BlindDetectionCG_UE3_r17, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Pdcch_BlindDetectionCG_UE3_r17", err)
	}
	return nil
}

func (ie *PDCCH_BlindDetectionCG_UE_Mixed1_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Pdcch_BlindDetectionCG_UE1_r17 int64
	if tmp_int_Pdcch_BlindDetectionCG_UE1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionCG_UE1_r17", err)
	}
	ie.Pdcch_BlindDetectionCG_UE1_r17 = tmp_int_Pdcch_BlindDetectionCG_UE1_r17
	var tmp_int_Pdcch_BlindDetectionCG_UE2_r17 int64
	if tmp_int_Pdcch_BlindDetectionCG_UE2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionCG_UE2_r17", err)
	}
	ie.Pdcch_BlindDetectionCG_UE2_r17 = tmp_int_Pdcch_BlindDetectionCG_UE2_r17
	var tmp_int_Pdcch_BlindDetectionCG_UE3_r17 int64
	if tmp_int_Pdcch_BlindDetectionCG_UE3_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Pdcch_BlindDetectionCG_UE3_r17", err)
	}
	ie.Pdcch_BlindDetectionCG_UE3_r17 = tmp_int_Pdcch_BlindDetectionCG_UE3_r17
	return nil
}
