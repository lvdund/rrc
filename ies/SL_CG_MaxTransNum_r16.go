package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_CG_MaxTransNum_r16 struct {
	Sl_Priority_r16    int64 `lb:1,ub:8,madatory`
	Sl_MaxTransNum_r16 int64 `lb:1,ub:32,madatory`
}

func (ie *SL_CG_MaxTransNum_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Sl_Priority_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_Priority_r16", err)
	}
	if err = w.WriteInteger(ie.Sl_MaxTransNum_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_MaxTransNum_r16", err)
	}
	return nil
}

func (ie *SL_CG_MaxTransNum_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Sl_Priority_r16 int64
	if tmp_int_Sl_Priority_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_Priority_r16", err)
	}
	ie.Sl_Priority_r16 = tmp_int_Sl_Priority_r16
	var tmp_int_Sl_MaxTransNum_r16 int64
	if tmp_int_Sl_MaxTransNum_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_MaxTransNum_r16", err)
	}
	ie.Sl_MaxTransNum_r16 = tmp_int_Sl_MaxTransNum_r16
	return nil
}
