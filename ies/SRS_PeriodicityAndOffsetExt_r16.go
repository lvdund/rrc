package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PeriodicityAndOffsetExt_r16_Choice_nothing uint64 = iota
	SRS_PeriodicityAndOffsetExt_r16_Choice_Sl128
	SRS_PeriodicityAndOffsetExt_r16_Choice_Sl256
	SRS_PeriodicityAndOffsetExt_r16_Choice_Sl512
	SRS_PeriodicityAndOffsetExt_r16_Choice_Sl20480
)

type SRS_PeriodicityAndOffsetExt_r16 struct {
	Choice  uint64
	Sl128   int64 `lb:0,ub:127,madatory`
	Sl256   int64 `lb:0,ub:255,madatory`
	Sl512   int64 `lb:0,ub:511,madatory`
	Sl20480 int64 `lb:0,ub:20479,madatory`
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl128:
		if err = w.WriteInteger(int64(ie.Sl128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode Sl128", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl256:
		if err = w.WriteInteger(int64(ie.Sl256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode Sl256", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl512:
		if err = w.WriteInteger(int64(ie.Sl512), &uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode Sl512", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl20480:
		if err = w.WriteInteger(int64(ie.Sl20480), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode Sl20480", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl128:
		var tmp_int_Sl128 int64
		if tmp_int_Sl128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Sl128", err)
		}
		ie.Sl128 = tmp_int_Sl128
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl256:
		var tmp_int_Sl256 int64
		if tmp_int_Sl256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode Sl256", err)
		}
		ie.Sl256 = tmp_int_Sl256
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl512:
		var tmp_int_Sl512 int64
		if tmp_int_Sl512, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode Sl512", err)
		}
		ie.Sl512 = tmp_int_Sl512
	case SRS_PeriodicityAndOffsetExt_r16_Choice_Sl20480:
		var tmp_int_Sl20480 int64
		if tmp_int_Sl20480, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode Sl20480", err)
		}
		ie.Sl20480 = tmp_int_Sl20480
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
