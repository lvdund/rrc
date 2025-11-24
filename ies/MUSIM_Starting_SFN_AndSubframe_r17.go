package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MUSIM_Starting_SFN_AndSubframe_r17 struct {
	Starting_SFN_r17     int64 `lb:0,ub:1023,madatory`
	StartingSubframe_r17 int64 `lb:0,ub:9,madatory`
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Starting_SFN_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger Starting_SFN_r17", err)
	}
	if err = w.WriteInteger(ie.StartingSubframe_r17, &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("WriteInteger StartingSubframe_r17", err)
	}
	return nil
}

func (ie *MUSIM_Starting_SFN_AndSubframe_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Starting_SFN_r17 int64
	if tmp_int_Starting_SFN_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger Starting_SFN_r17", err)
	}
	ie.Starting_SFN_r17 = tmp_int_Starting_SFN_r17
	var tmp_int_StartingSubframe_r17 int64
	if tmp_int_StartingSubframe_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("ReadInteger StartingSubframe_r17", err)
	}
	ie.StartingSubframe_r17 = tmp_int_StartingSubframe_r17
	return nil
}
