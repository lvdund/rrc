package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DRX_ConfigSL_r17 struct {
	Drx_HARQ_RTT_TimerSL_r17      int64                                          `lb:0,ub:56,madatory`
	Drx_RetransmissionTimerSL_r17 DRX_ConfigSL_r17_drx_RetransmissionTimerSL_r17 `madatory`
}

func (ie *DRX_ConfigSL_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.Drx_HARQ_RTT_TimerSL_r17, &aper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("WriteInteger Drx_HARQ_RTT_TimerSL_r17", err)
	}
	if err = ie.Drx_RetransmissionTimerSL_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}

func (ie *DRX_ConfigSL_r17) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_Drx_HARQ_RTT_TimerSL_r17 int64
	if tmp_int_Drx_HARQ_RTT_TimerSL_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 56}, false); err != nil {
		return utils.WrapError("ReadInteger Drx_HARQ_RTT_TimerSL_r17", err)
	}
	ie.Drx_HARQ_RTT_TimerSL_r17 = tmp_int_Drx_HARQ_RTT_TimerSL_r17
	if err = ie.Drx_RetransmissionTimerSL_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_RetransmissionTimerSL_r17", err)
	}
	return nil
}
