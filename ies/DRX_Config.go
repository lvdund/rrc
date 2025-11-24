package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Config struct {
	Drx_onDurationTimer       DRX_Config_drx_onDurationTimer       `lb:1,ub:31,madatory`
	Drx_InactivityTimer       DRX_Config_drx_InactivityTimer       `madatory`
	Drx_HARQ_RTT_TimerDL      int64                                `lb:0,ub:56,madatory`
	Drx_HARQ_RTT_TimerUL      int64                                `lb:0,ub:56,madatory`
	Drx_RetransmissionTimerDL DRX_Config_drx_RetransmissionTimerDL `madatory`
	Drx_RetransmissionTimerUL DRX_Config_drx_RetransmissionTimerUL `madatory`
	Drx_LongCycleStartOffset  DRX_Config_drx_LongCycleStartOffset  `lb:0,ub:9,madatory`
	ShortDRX                  *DRX_Config_shortDRX                 `lb:1,ub:16,optional`
	Drx_SlotOffset            int64                                `lb:0,ub:31,madatory`
}

func (ie *DRX_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ShortDRX != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Drx_onDurationTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_onDurationTimer", err)
	}
	if err = ie.Drx_InactivityTimer.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_InactivityTimer", err)
	}
	if err = w.WriteInteger(ie.Drx_HARQ_RTT_TimerDL, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_HARQ_RTT_TimerDL", err)
	}
	if err = w.WriteInteger(ie.Drx_HARQ_RTT_TimerUL, &uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_HARQ_RTT_TimerUL", err)
	}
	if err = ie.Drx_RetransmissionTimerDL.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_RetransmissionTimerDL", err)
	}
	if err = ie.Drx_RetransmissionTimerUL.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_RetransmissionTimerUL", err)
	}
	if err = ie.Drx_LongCycleStartOffset.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_LongCycleStartOffset", err)
	}
	if ie.ShortDRX != nil {
		if err = ie.ShortDRX.Encode(w); err != nil {
			return utils.WrapError("Encode ShortDRX", err)
		}
	}
	if err = w.WriteInteger(ie.Drx_SlotOffset, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_SlotOffset", err)
	}
	return nil
}

func (ie *DRX_Config) Decode(r *uper.UperReader) error {
	var err error
	var ShortDRXPresent bool
	if ShortDRXPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Drx_onDurationTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_onDurationTimer", err)
	}
	if err = ie.Drx_InactivityTimer.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_InactivityTimer", err)
	}
	var tmp_int_Drx_HARQ_RTT_TimerDL int64
	if tmp_int_Drx_HARQ_RTT_TimerDL, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_HARQ_RTT_TimerDL", err)
	}
	ie.Drx_HARQ_RTT_TimerDL = tmp_int_Drx_HARQ_RTT_TimerDL
	var tmp_int_Drx_HARQ_RTT_TimerUL int64
	if tmp_int_Drx_HARQ_RTT_TimerUL, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_HARQ_RTT_TimerUL", err)
	}
	ie.Drx_HARQ_RTT_TimerUL = tmp_int_Drx_HARQ_RTT_TimerUL
	if err = ie.Drx_RetransmissionTimerDL.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_RetransmissionTimerDL", err)
	}
	if err = ie.Drx_RetransmissionTimerUL.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_RetransmissionTimerUL", err)
	}
	if err = ie.Drx_LongCycleStartOffset.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_LongCycleStartOffset", err)
	}
	if ShortDRXPresent {
		ie.ShortDRX = new(DRX_Config_shortDRX)
		if err = ie.ShortDRX.Decode(r); err != nil {
			return utils.WrapError("Decode ShortDRX", err)
		}
	}
	var tmp_int_Drx_SlotOffset int64
	if tmp_int_Drx_SlotOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_SlotOffset", err)
	}
	ie.Drx_SlotOffset = tmp_int_Drx_SlotOffset
	return nil
}
