package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigUC_SemiStatic_r17 struct {
	Sl_drx_onDurationTimer_r17  SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17  `lb:1,ub:31,madatory`
	Sl_drx_CycleStartOffset_r17 SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_CycleStartOffset_r17 `lb:0,ub:9,madatory`
	Sl_drx_SlotOffset_r17       int64                                                      `lb:0,ub:31,madatory`
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_drx_onDurationTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_onDurationTimer_r17", err)
	}
	if err = ie.Sl_drx_CycleStartOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_CycleStartOffset_r17", err)
	}
	if err = w.WriteInteger(ie.Sl_drx_SlotOffset_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_drx_SlotOffset_r17", err)
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_drx_onDurationTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_onDurationTimer_r17", err)
	}
	if err = ie.Sl_drx_CycleStartOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_CycleStartOffset_r17", err)
	}
	var tmp_int_Sl_drx_SlotOffset_r17 int64
	if tmp_int_Sl_drx_SlotOffset_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_drx_SlotOffset_r17", err)
	}
	ie.Sl_drx_SlotOffset_r17 = tmp_int_Sl_drx_SlotOffset_r17
	return nil
}
