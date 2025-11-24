package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_nothing uint64 = iota
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_SubMilliSeconds
	DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_MilliSeconds
)

type DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16 struct {
	Choice          uint64
	SubMilliSeconds int64 `lb:1,ub:31,madatory`
	MilliSeconds    *DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds
}

func (ie *DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_SubMilliSeconds:
		if err = w.WriteInteger(int64(ie.SubMilliSeconds), &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode SubMilliSeconds", err)
		}
	case DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_MilliSeconds:
		if err = ie.MilliSeconds.Encode(w); err != nil {
			err = utils.WrapError("Encode MilliSeconds", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_SubMilliSeconds:
		var tmp_int_SubMilliSeconds int64
		if tmp_int_SubMilliSeconds, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode SubMilliSeconds", err)
		}
		ie.SubMilliSeconds = tmp_int_SubMilliSeconds
	case DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_Choice_MilliSeconds:
		ie.MilliSeconds = new(DRX_ConfigSecondaryGroup_r16_drx_onDurationTimer_r16_milliSeconds)
		if err = ie.MilliSeconds.Decode(r); err != nil {
			return utils.WrapError("Decode MilliSeconds", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
