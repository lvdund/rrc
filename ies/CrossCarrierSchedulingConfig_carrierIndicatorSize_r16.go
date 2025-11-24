package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingConfig_carrierIndicatorSize_r16 struct {
	CarrierIndicatorSizeDCI_1_2_r16 int64 `lb:0,ub:3,madatory`
	CarrierIndicatorSizeDCI_0_2_r16 int64 `lb:0,ub:3,madatory`
}

func (ie *CrossCarrierSchedulingConfig_carrierIndicatorSize_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.CarrierIndicatorSizeDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger CarrierIndicatorSizeDCI_1_2_r16", err)
	}
	if err = w.WriteInteger(ie.CarrierIndicatorSizeDCI_0_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("WriteInteger CarrierIndicatorSizeDCI_0_2_r16", err)
	}
	return nil
}

func (ie *CrossCarrierSchedulingConfig_carrierIndicatorSize_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_CarrierIndicatorSizeDCI_1_2_r16 int64
	if tmp_int_CarrierIndicatorSizeDCI_1_2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger CarrierIndicatorSizeDCI_1_2_r16", err)
	}
	ie.CarrierIndicatorSizeDCI_1_2_r16 = tmp_int_CarrierIndicatorSizeDCI_1_2_r16
	var tmp_int_CarrierIndicatorSizeDCI_0_2_r16 int64
	if tmp_int_CarrierIndicatorSizeDCI_0_2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("ReadInteger CarrierIndicatorSizeDCI_0_2_r16", err)
	}
	ie.CarrierIndicatorSizeDCI_0_2_r16 = tmp_int_CarrierIndicatorSizeDCI_0_2_r16
	return nil
}
