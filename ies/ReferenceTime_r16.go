package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceTime_r16 struct {
	RefDays_r16           int64 `lb:0,ub:72999,madatory`
	RefSeconds_r16        int64 `lb:0,ub:86399,madatory`
	RefMilliSeconds_r16   int64 `lb:0,ub:999,madatory`
	RefTenNanoSeconds_r16 int64 `lb:0,ub:99999,madatory`
}

func (ie *ReferenceTime_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.RefDays_r16, &uper.Constraint{Lb: 0, Ub: 72999}, false); err != nil {
		return utils.WrapError("WriteInteger RefDays_r16", err)
	}
	if err = w.WriteInteger(ie.RefSeconds_r16, &uper.Constraint{Lb: 0, Ub: 86399}, false); err != nil {
		return utils.WrapError("WriteInteger RefSeconds_r16", err)
	}
	if err = w.WriteInteger(ie.RefMilliSeconds_r16, &uper.Constraint{Lb: 0, Ub: 999}, false); err != nil {
		return utils.WrapError("WriteInteger RefMilliSeconds_r16", err)
	}
	if err = w.WriteInteger(ie.RefTenNanoSeconds_r16, &uper.Constraint{Lb: 0, Ub: 99999}, false); err != nil {
		return utils.WrapError("WriteInteger RefTenNanoSeconds_r16", err)
	}
	return nil
}

func (ie *ReferenceTime_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_RefDays_r16 int64
	if tmp_int_RefDays_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 72999}, false); err != nil {
		return utils.WrapError("ReadInteger RefDays_r16", err)
	}
	ie.RefDays_r16 = tmp_int_RefDays_r16
	var tmp_int_RefSeconds_r16 int64
	if tmp_int_RefSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 86399}, false); err != nil {
		return utils.WrapError("ReadInteger RefSeconds_r16", err)
	}
	ie.RefSeconds_r16 = tmp_int_RefSeconds_r16
	var tmp_int_RefMilliSeconds_r16 int64
	if tmp_int_RefMilliSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 999}, false); err != nil {
		return utils.WrapError("ReadInteger RefMilliSeconds_r16", err)
	}
	ie.RefMilliSeconds_r16 = tmp_int_RefMilliSeconds_r16
	var tmp_int_RefTenNanoSeconds_r16 int64
	if tmp_int_RefTenNanoSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 99999}, false); err != nil {
		return utils.WrapError("ReadInteger RefTenNanoSeconds_r16", err)
	}
	ie.RefTenNanoSeconds_r16 = tmp_int_RefTenNanoSeconds_r16
	return nil
}
