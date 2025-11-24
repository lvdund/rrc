package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_DRX_ConfigUC_r17 struct {
	Sl_drx_onDurationTimer_r17     SL_DRX_ConfigUC_r17_sl_drx_onDurationTimer_r17     `lb:1,ub:31,madatory`
	Sl_drx_InactivityTimer_r17     SL_DRX_ConfigUC_r17_sl_drx_InactivityTimer_r17     `madatory`
	Sl_drx_HARQ_RTT_Timer1_r17     *SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer1_r17    `optional`
	Sl_drx_HARQ_RTT_Timer2_r17     *SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17    `optional`
	Sl_drx_RetransmissionTimer_r17 SL_DRX_ConfigUC_r17_sl_drx_RetransmissionTimer_r17 `madatory`
	Sl_drx_CycleStartOffset_r17    SL_DRX_ConfigUC_r17_sl_drx_CycleStartOffset_r17    `lb:0,ub:9,madatory`
	Sl_drx_SlotOffset              int64                                              `lb:0,ub:31,madatory`
}

func (ie *SL_DRX_ConfigUC_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_drx_HARQ_RTT_Timer1_r17 != nil, ie.Sl_drx_HARQ_RTT_Timer2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_drx_onDurationTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_onDurationTimer_r17", err)
	}
	if err = ie.Sl_drx_InactivityTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_InactivityTimer_r17", err)
	}
	if ie.Sl_drx_HARQ_RTT_Timer1_r17 != nil {
		if err = ie.Sl_drx_HARQ_RTT_Timer1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_drx_HARQ_RTT_Timer1_r17", err)
		}
	}
	if ie.Sl_drx_HARQ_RTT_Timer2_r17 != nil {
		if err = ie.Sl_drx_HARQ_RTT_Timer2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_drx_HARQ_RTT_Timer2_r17", err)
		}
	}
	if err = ie.Sl_drx_RetransmissionTimer_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_RetransmissionTimer_r17", err)
	}
	if err = ie.Sl_drx_CycleStartOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_drx_CycleStartOffset_r17", err)
	}
	if err = w.WriteInteger(ie.Sl_drx_SlotOffset, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_drx_SlotOffset", err)
	}
	return nil
}

func (ie *SL_DRX_ConfigUC_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_drx_HARQ_RTT_Timer1_r17Present bool
	if Sl_drx_HARQ_RTT_Timer1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_drx_HARQ_RTT_Timer2_r17Present bool
	if Sl_drx_HARQ_RTT_Timer2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_drx_onDurationTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_onDurationTimer_r17", err)
	}
	if err = ie.Sl_drx_InactivityTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_InactivityTimer_r17", err)
	}
	if Sl_drx_HARQ_RTT_Timer1_r17Present {
		ie.Sl_drx_HARQ_RTT_Timer1_r17 = new(SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer1_r17)
		if err = ie.Sl_drx_HARQ_RTT_Timer1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_drx_HARQ_RTT_Timer1_r17", err)
		}
	}
	if Sl_drx_HARQ_RTT_Timer2_r17Present {
		ie.Sl_drx_HARQ_RTT_Timer2_r17 = new(SL_DRX_ConfigUC_r17_sl_drx_HARQ_RTT_Timer2_r17)
		if err = ie.Sl_drx_HARQ_RTT_Timer2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_drx_HARQ_RTT_Timer2_r17", err)
		}
	}
	if err = ie.Sl_drx_RetransmissionTimer_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_RetransmissionTimer_r17", err)
	}
	if err = ie.Sl_drx_CycleStartOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_drx_CycleStartOffset_r17", err)
	}
	var tmp_int_Sl_drx_SlotOffset int64
	if tmp_int_Sl_drx_SlotOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_drx_SlotOffset", err)
	}
	ie.Sl_drx_SlotOffset = tmp_int_Sl_drx_SlotOffset
	return nil
}
