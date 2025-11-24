package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventT1_r17 struct {
	T1_Threshold_r17 int64 `lb:0,ub:549755813887,madatory`
	Duration_r17     int64 `lb:1,ub:6000,madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventT1_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.T1_Threshold_r17, &uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("WriteInteger T1_Threshold_r17", err)
	}
	if err = w.WriteInteger(ie.Duration_r17, &uper.Constraint{Lb: 1, Ub: 6000}, false); err != nil {
		return utils.WrapError("WriteInteger Duration_r17", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventT1_r17) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_T1_Threshold_r17 int64
	if tmp_int_T1_Threshold_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 549755813887}, false); err != nil {
		return utils.WrapError("ReadInteger T1_Threshold_r17", err)
	}
	ie.T1_Threshold_r17 = tmp_int_T1_Threshold_r17
	var tmp_int_Duration_r17 int64
	if tmp_int_Duration_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 6000}, false); err != nil {
		return utils.WrapError("ReadInteger Duration_r17", err)
	}
	ie.Duration_r17 = tmp_int_Duration_r17
	return nil
}
