package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReducedMaxCCs_r16 struct {
	ReducedCCsDL_r16 int64 `lb:0,ub:31,madatory`
	ReducedCCsUL_r16 int64 `lb:0,ub:31,madatory`
}

func (ie *ReducedMaxCCs_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.ReducedCCsDL_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedCCsDL_r16", err)
	}
	if err = w.WriteInteger(ie.ReducedCCsUL_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger ReducedCCsUL_r16", err)
	}
	return nil
}

func (ie *ReducedMaxCCs_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_ReducedCCsDL_r16 int64
	if tmp_int_ReducedCCsDL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedCCsDL_r16", err)
	}
	ie.ReducedCCsDL_r16 = tmp_int_ReducedCCsDL_r16
	var tmp_int_ReducedCCsUL_r16 int64
	if tmp_int_ReducedCCsUL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger ReducedCCsUL_r16", err)
	}
	ie.ReducedCCsUL_r16 = tmp_int_ReducedCCsUL_r16
	return nil
}
