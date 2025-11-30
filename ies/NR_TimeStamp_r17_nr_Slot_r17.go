package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_TimeStamp_r17_nr_Slot_r17_Choice_nothing uint64 = iota
	NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs15_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs30_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs60_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs120_r17
)

type NR_TimeStamp_r17_nr_Slot_r17 struct {
	Choice     uint64
	Scs15_r17  int64 `lb:0,ub:9,madatory`
	Scs30_r17  int64 `lb:0,ub:19,madatory`
	Scs60_r17  int64 `lb:0,ub:39,madatory`
	Scs120_r17 int64 `lb:0,ub:79,madatory`
}

func (ie *NR_TimeStamp_r17_nr_Slot_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs15_r17:
		if err = w.WriteInteger(int64(ie.Scs15_r17), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Scs15_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs30_r17:
		if err = w.WriteInteger(int64(ie.Scs30_r17), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Scs30_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs60_r17:
		if err = w.WriteInteger(int64(ie.Scs60_r17), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Scs60_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs120_r17:
		if err = w.WriteInteger(int64(ie.Scs120_r17), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Scs120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_TimeStamp_r17_nr_Slot_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs15_r17:
		var tmp_int_Scs15_r17 int64
		if tmp_int_Scs15_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Scs15_r17", err)
		}
		ie.Scs15_r17 = tmp_int_Scs15_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs30_r17:
		var tmp_int_Scs30_r17 int64
		if tmp_int_Scs30_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Scs30_r17", err)
		}
		ie.Scs30_r17 = tmp_int_Scs30_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs60_r17:
		var tmp_int_Scs60_r17 int64
		if tmp_int_Scs60_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Scs60_r17", err)
		}
		ie.Scs60_r17 = tmp_int_Scs60_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_Scs120_r17:
		var tmp_int_Scs120_r17 int64
		if tmp_int_Scs120_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Scs120_r17", err)
		}
		ie.Scs120_r17 = tmp_int_Scs120_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
