package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_nothing uint64 = iota
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl1280
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl2560
	SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl5120
)

type SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17 struct {
	Choice uint64
	Sl1280 int64 `lb:0,ub:1279,madatory`
	Sl2560 int64 `lb:0,ub:2559,madatory`
	Sl5120 int64 `lb:0,ub:5119,madatory`
}

func (ie *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl1280:
		if err = w.WriteInteger(int64(ie.Sl1280), &aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Sl1280", err)
		}
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl2560:
		if err = w.WriteInteger(int64(ie.Sl2560), &aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode Sl2560", err)
		}
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl5120:
		if err = w.WriteInteger(int64(ie.Sl5120), &aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode Sl5120", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl1280:
		var tmp_int_Sl1280 int64
		if tmp_int_Sl1280, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Sl1280", err)
		}
		ie.Sl1280 = tmp_int_Sl1280
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl2560:
		var tmp_int_Sl2560 int64
		if tmp_int_Sl2560, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Sl2560", err)
		}
		ie.Sl2560 = tmp_int_Sl2560
	case SchedulingRequestResourceConfigExt_v1700_periodicityAndOffset_r17_Choice_Sl5120:
		var tmp_int_Sl5120 int64
		if tmp_int_Sl5120, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode Sl5120", err)
		}
		ie.Sl5120 = tmp_int_Sl5120
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
