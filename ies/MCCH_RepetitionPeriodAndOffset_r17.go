package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_RepetitionPeriodAndOffset_r17_Choice_nothing uint64 = iota
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf1_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf2_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf4_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf8_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf16_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf32_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf64_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf128_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf256_r17
)

type MCCH_RepetitionPeriodAndOffset_r17 struct {
	Choice    uint64
	Rf1_r17   int64 `lb:0,ub:0,madatory`
	Rf2_r17   int64 `lb:0,ub:1,madatory`
	Rf4_r17   int64 `lb:0,ub:3,madatory`
	Rf8_r17   int64 `lb:0,ub:7,madatory`
	Rf16_r17  int64 `lb:0,ub:15,madatory`
	Rf32_r17  int64 `lb:0,ub:31,madatory`
	Rf64_r17  int64 `lb:0,ub:63,madatory`
	Rf128_r17 int64 `lb:0,ub:127,madatory`
	Rf256_r17 int64 `lb:0,ub:255,madatory`
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf1_r17:
		if err = w.WriteInteger(int64(ie.Rf1_r17), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode Rf1_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf2_r17:
		if err = w.WriteInteger(int64(ie.Rf2_r17), &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode Rf2_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf4_r17:
		if err = w.WriteInteger(int64(ie.Rf4_r17), &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Rf4_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf8_r17:
		if err = w.WriteInteger(int64(ie.Rf8_r17), &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Rf8_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf16_r17:
		if err = w.WriteInteger(int64(ie.Rf16_r17), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Rf16_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf32_r17:
		if err = w.WriteInteger(int64(ie.Rf32_r17), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Rf32_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf64_r17:
		if err = w.WriteInteger(int64(ie.Rf64_r17), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Rf64_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf128_r17:
		if err = w.WriteInteger(int64(ie.Rf128_r17), &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode Rf128_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf256_r17:
		if err = w.WriteInteger(int64(ie.Rf256_r17), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode Rf256_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf1_r17:
		var tmp_int_Rf1_r17 int64
		if tmp_int_Rf1_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Rf1_r17", err)
		}
		ie.Rf1_r17 = tmp_int_Rf1_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf2_r17:
		var tmp_int_Rf2_r17 int64
		if tmp_int_Rf2_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Rf2_r17", err)
		}
		ie.Rf2_r17 = tmp_int_Rf2_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf4_r17:
		var tmp_int_Rf4_r17 int64
		if tmp_int_Rf4_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Rf4_r17", err)
		}
		ie.Rf4_r17 = tmp_int_Rf4_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf8_r17:
		var tmp_int_Rf8_r17 int64
		if tmp_int_Rf8_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Rf8_r17", err)
		}
		ie.Rf8_r17 = tmp_int_Rf8_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf16_r17:
		var tmp_int_Rf16_r17 int64
		if tmp_int_Rf16_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Rf16_r17", err)
		}
		ie.Rf16_r17 = tmp_int_Rf16_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf32_r17:
		var tmp_int_Rf32_r17 int64
		if tmp_int_Rf32_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Rf32_r17", err)
		}
		ie.Rf32_r17 = tmp_int_Rf32_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf64_r17:
		var tmp_int_Rf64_r17 int64
		if tmp_int_Rf64_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Rf64_r17", err)
		}
		ie.Rf64_r17 = tmp_int_Rf64_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf128_r17:
		var tmp_int_Rf128_r17 int64
		if tmp_int_Rf128_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Rf128_r17", err)
		}
		ie.Rf128_r17 = tmp_int_Rf128_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_Rf256_r17:
		var tmp_int_Rf256_r17 int64
		if tmp_int_Rf256_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode Rf256_r17", err)
		}
		ie.Rf256_r17 = tmp_int_Rf256_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
