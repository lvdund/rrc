package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigPTM_r17 struct {
	Drx_onDurationTimerPTM_r17        DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17         `lb:1,ub:31,madatory`
	Drx_InactivityTimerPTM_r17        DRX_ConfigPTM_r17_drx_InactivityTimerPTM_r17         `madatory`
	Drx_HARQ_RTT_TimerDL_PTM_r17      *int64                                               `lb:0,ub:56,optional`
	Drx_RetransmissionTimerDL_PTM_r17 *DRX_ConfigPTM_r17_drx_RetransmissionTimerDL_PTM_r17 `optional`
	Drx_LongCycleStartOffsetPTM_r17   DRX_ConfigPTM_r17_drx_LongCycleStartOffsetPTM_r17    `lb:0,ub:9,madatory`
	Drx_SlotOffsetPTM_r17             int64                                                `lb:0,ub:31,madatory`
}

func (ie *DRX_ConfigPTM_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Drx_HARQ_RTT_TimerDL_PTM_r17 != nil, ie.Drx_RetransmissionTimerDL_PTM_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Drx_onDurationTimerPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_onDurationTimerPTM_r17", err)
	}
	if err = ie.Drx_InactivityTimerPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_InactivityTimerPTM_r17", err)
	}
	if ie.Drx_HARQ_RTT_TimerDL_PTM_r17 != nil {
		if err = w.WriteInteger(*ie.Drx_HARQ_RTT_TimerDL_PTM_r17, &aper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
			return utils.WrapError("Encode Drx_HARQ_RTT_TimerDL_PTM_r17", err)
		}
	}
	if ie.Drx_RetransmissionTimerDL_PTM_r17 != nil {
		if err = ie.Drx_RetransmissionTimerDL_PTM_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Drx_RetransmissionTimerDL_PTM_r17", err)
		}
	}
	if err = ie.Drx_LongCycleStartOffsetPTM_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_LongCycleStartOffsetPTM_r17", err)
	}
	if err = w.WriteInteger(ie.Drx_SlotOffsetPTM_r17, &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_SlotOffsetPTM_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigPTM_r17) Decode(r *aper.AperReader) error {
	var err error
	var Drx_HARQ_RTT_TimerDL_PTM_r17Present bool
	if Drx_HARQ_RTT_TimerDL_PTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Drx_RetransmissionTimerDL_PTM_r17Present bool
	if Drx_RetransmissionTimerDL_PTM_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Drx_onDurationTimerPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_onDurationTimerPTM_r17", err)
	}
	if err = ie.Drx_InactivityTimerPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_InactivityTimerPTM_r17", err)
	}
	if Drx_HARQ_RTT_TimerDL_PTM_r17Present {
		var tmp_int_Drx_HARQ_RTT_TimerDL_PTM_r17 int64
		if tmp_int_Drx_HARQ_RTT_TimerDL_PTM_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
			return utils.WrapError("Decode Drx_HARQ_RTT_TimerDL_PTM_r17", err)
		}
		ie.Drx_HARQ_RTT_TimerDL_PTM_r17 = &tmp_int_Drx_HARQ_RTT_TimerDL_PTM_r17
	}
	if Drx_RetransmissionTimerDL_PTM_r17Present {
		ie.Drx_RetransmissionTimerDL_PTM_r17 = new(DRX_ConfigPTM_r17_drx_RetransmissionTimerDL_PTM_r17)
		if err = ie.Drx_RetransmissionTimerDL_PTM_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Drx_RetransmissionTimerDL_PTM_r17", err)
		}
	}
	if err = ie.Drx_LongCycleStartOffsetPTM_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_LongCycleStartOffsetPTM_r17", err)
	}
	var tmp_int_Drx_SlotOffsetPTM_r17 int64
	if tmp_int_Drx_SlotOffsetPTM_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_SlotOffsetPTM_r17", err)
	}
	ie.Drx_SlotOffsetPTM_r17 = tmp_int_Drx_SlotOffsetPTM_r17
	return nil
}
