package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type GuardBand_r16 struct {
	StartCRB_r16 int64 `lb:0,ub:274,madatory`
	NrofCRBs_r16 int64 `lb:0,ub:15,madatory`
}

func (ie *GuardBand_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.StartCRB_r16, &uper.Constraint{Lb: 0, Ub: 274}, false); err != nil {
		return utils.WrapError("WriteInteger StartCRB_r16", err)
	}
	if err = w.WriteInteger(ie.NrofCRBs_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger NrofCRBs_r16", err)
	}
	return nil
}

func (ie *GuardBand_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_StartCRB_r16 int64
	if tmp_int_StartCRB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 274}, false); err != nil {
		return utils.WrapError("ReadInteger StartCRB_r16", err)
	}
	ie.StartCRB_r16 = tmp_int_StartCRB_r16
	var tmp_int_NrofCRBs_r16 int64
	if tmp_int_NrofCRBs_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger NrofCRBs_r16", err)
	}
	ie.NrofCRBs_r16 = tmp_int_NrofCRBs_r16
	return nil
}
