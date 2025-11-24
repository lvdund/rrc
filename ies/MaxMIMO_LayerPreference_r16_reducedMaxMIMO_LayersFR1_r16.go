package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16 struct {
	ReducedMIMO_LayersFR1_DL_r16 int64 `lb:1,ub:8,madatory`
	ReducedMIMO_LayersFR1_UL_r16 int64 `lb:1,ub:4,madatory`
}

func (ie *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.ReducedMIMO_LayersFR1_DL_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedMIMO_LayersFR1_DL_r16", err)
	}
	if err = w.WriteInteger(ie.ReducedMIMO_LayersFR1_UL_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedMIMO_LayersFR1_UL_r16", err)
	}
	return nil
}

func (ie *MaxMIMO_LayerPreference_r16_reducedMaxMIMO_LayersFR1_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_ReducedMIMO_LayersFR1_DL_r16 int64
	if tmp_int_ReducedMIMO_LayersFR1_DL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedMIMO_LayersFR1_DL_r16", err)
	}
	ie.ReducedMIMO_LayersFR1_DL_r16 = tmp_int_ReducedMIMO_LayersFR1_DL_r16
	var tmp_int_ReducedMIMO_LayersFR1_UL_r16 int64
	if tmp_int_ReducedMIMO_LayersFR1_UL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedMIMO_LayersFR1_UL_r16", err)
	}
	ie.ReducedMIMO_LayersFR1_UL_r16 = tmp_int_ReducedMIMO_LayersFR1_UL_r16
	return nil
}
