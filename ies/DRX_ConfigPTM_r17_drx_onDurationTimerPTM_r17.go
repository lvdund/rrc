package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_nothing uint64 = iota
	DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_SubMilliSeconds
	DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_MilliSeconds
)

type DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17 struct {
	Choice          uint64
	SubMilliSeconds int64 `lb:1,ub:31,madatory`
	MilliSeconds    *DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_milliSeconds
}

func (ie *DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_SubMilliSeconds:
		if err = w.WriteInteger(int64(ie.SubMilliSeconds), &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode SubMilliSeconds", err)
		}
	case DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_MilliSeconds:
		if err = ie.MilliSeconds.Encode(w); err != nil {
			err = utils.WrapError("Encode MilliSeconds", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_SubMilliSeconds:
		var tmp_int_SubMilliSeconds int64
		if tmp_int_SubMilliSeconds, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode SubMilliSeconds", err)
		}
		ie.SubMilliSeconds = tmp_int_SubMilliSeconds
	case DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_Choice_MilliSeconds:
		ie.MilliSeconds = new(DRX_ConfigPTM_r17_drx_onDurationTimerPTM_r17_milliSeconds)
		if err = ie.MilliSeconds.Decode(r); err != nil {
			return utils.WrapError("Decode MilliSeconds", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
