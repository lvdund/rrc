package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourceReservePeriod_r16_Choice_nothing uint64 = iota
	SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod1_r16
	SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod2_r16
)

type SL_ResourceReservePeriod_r16 struct {
	Choice                        uint64
	Sl_ResourceReservePeriod1_r16 *SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16
	Sl_ResourceReservePeriod2_r16 int64 `lb:1,ub:99,madatory`
}

func (ie *SL_ResourceReservePeriod_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod1_r16:
		if err = ie.Sl_ResourceReservePeriod1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode Sl_ResourceReservePeriod1_r16", err)
		}
	case SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod2_r16:
		if err = w.WriteInteger(int64(ie.Sl_ResourceReservePeriod2_r16), &aper.Constraint{Lb: 1, Ub: 99}, false); err != nil {
			err = utils.WrapError("Encode Sl_ResourceReservePeriod2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_ResourceReservePeriod_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod1_r16:
		ie.Sl_ResourceReservePeriod1_r16 = new(SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16)
		if err = ie.Sl_ResourceReservePeriod1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ResourceReservePeriod1_r16", err)
		}
	case SL_ResourceReservePeriod_r16_Choice_Sl_ResourceReservePeriod2_r16:
		var tmp_int_Sl_ResourceReservePeriod2_r16 int64
		if tmp_int_Sl_ResourceReservePeriod2_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 99}, false); err != nil {
			return utils.WrapError("Decode Sl_ResourceReservePeriod2_r16", err)
		}
		ie.Sl_ResourceReservePeriod2_r16 = tmp_int_Sl_ResourceReservePeriod2_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
