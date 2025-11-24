package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigExt_v1700 struct {
	Drx_HARQ_RTT_TimerDL_r17 int64 `lb:0,ub:448,madatory`
	Drx_HARQ_RTT_TimerUL_r17 int64 `lb:0,ub:448,madatory`
}

func (ie *DRX_ConfigExt_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Drx_HARQ_RTT_TimerDL_r17, &uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_HARQ_RTT_TimerDL_r17", err)
	}
	if err = w.WriteInteger(ie.Drx_HARQ_RTT_TimerUL_r17, &uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_HARQ_RTT_TimerUL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigExt_v1700) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_Drx_HARQ_RTT_TimerDL_r17 int64
	if tmp_int_Drx_HARQ_RTT_TimerDL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_HARQ_RTT_TimerDL_r17", err)
	}
	ie.Drx_HARQ_RTT_TimerDL_r17 = tmp_int_Drx_HARQ_RTT_TimerDL_r17
	var tmp_int_Drx_HARQ_RTT_TimerUL_r17 int64
	if tmp_int_Drx_HARQ_RTT_TimerUL_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 448}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_HARQ_RTT_TimerUL_r17", err)
	}
	ie.Drx_HARQ_RTT_TimerUL_r17 = tmp_int_Drx_HARQ_RTT_TimerUL_r17
	return nil
}
