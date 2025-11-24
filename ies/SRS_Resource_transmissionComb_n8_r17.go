package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_Resource_transmissionComb_n8_r17 struct {
	CombOffset_n8_r17  int64 `lb:0,ub:7,madatory`
	CyclicShift_n8_r17 int64 `lb:0,ub:5,madatory`
}

func (ie *SRS_Resource_transmissionComb_n8_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.CombOffset_n8_r17, &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteInteger CombOffset_n8_r17", err)
	}
	if err = w.WriteInteger(ie.CyclicShift_n8_r17, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger CyclicShift_n8_r17", err)
	}
	return nil
}

func (ie *SRS_Resource_transmissionComb_n8_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_CombOffset_n8_r17 int64
	if tmp_int_CombOffset_n8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadInteger CombOffset_n8_r17", err)
	}
	ie.CombOffset_n8_r17 = tmp_int_CombOffset_n8_r17
	var tmp_int_CyclicShift_n8_r17 int64
	if tmp_int_CyclicShift_n8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger CyclicShift_n8_r17", err)
	}
	ie.CyclicShift_n8_r17 = tmp_int_CyclicShift_n8_r17
	return nil
}
