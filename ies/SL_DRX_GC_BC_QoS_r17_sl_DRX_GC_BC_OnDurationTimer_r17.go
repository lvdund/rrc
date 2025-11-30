package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_nothing uint64 = iota
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_SubMilliSeconds
	SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_MilliSeconds
)

type SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17 struct {
	Choice          uint64
	SubMilliSeconds int64 `lb:1,ub:31,madatory`
	MilliSeconds    *SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_milliSeconds
}

func (ie *SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_SubMilliSeconds:
		if err = w.WriteInteger(int64(ie.SubMilliSeconds), &aper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode SubMilliSeconds", err)
		}
	case SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_MilliSeconds:
		if err = ie.MilliSeconds.Encode(w); err != nil {
			err = utils.WrapError("Encode MilliSeconds", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_SubMilliSeconds:
		var tmp_int_SubMilliSeconds int64
		if tmp_int_SubMilliSeconds, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode SubMilliSeconds", err)
		}
		ie.SubMilliSeconds = tmp_int_SubMilliSeconds
	case SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_Choice_MilliSeconds:
		ie.MilliSeconds = new(SL_DRX_GC_BC_QoS_r17_sl_DRX_GC_BC_OnDurationTimer_r17_milliSeconds)
		if err = ie.MilliSeconds.Decode(r); err != nil {
			return utils.WrapError("Decode MilliSeconds", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
