package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_nothing uint64 = iota
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms10
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms20
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms32
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms64
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms128
	MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms256
)

type MTCH_SSB_MappingWindowCycleOffset_r17 struct {
	Choice uint64
	Ms10   int64 `lb:0,ub:9,madatory`
	Ms20   int64 `lb:0,ub:19,madatory`
	Ms32   int64 `lb:0,ub:31,madatory`
	Ms64   int64 `lb:0,ub:63,madatory`
	Ms128  int64 `lb:0,ub:127,madatory`
	Ms256  int64 `lb:0,ub:255,madatory`
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms10:
		if err = w.WriteInteger(int64(ie.Ms10), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Ms10", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms20:
		if err = w.WriteInteger(int64(ie.Ms20), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Ms20", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms32:
		if err = w.WriteInteger(int64(ie.Ms32), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Ms32", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms64:
		if err = w.WriteInteger(int64(ie.Ms64), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Ms64", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms128:
		if err = w.WriteInteger(int64(ie.Ms128), &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode Ms128", err)
		}
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms256:
		if err = w.WriteInteger(int64(ie.Ms256), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode Ms256", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MTCH_SSB_MappingWindowCycleOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms10:
		var tmp_int_Ms10 int64
		if tmp_int_Ms10, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Ms10", err)
		}
		ie.Ms10 = tmp_int_Ms10
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms20:
		var tmp_int_Ms20 int64
		if tmp_int_Ms20, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Ms20", err)
		}
		ie.Ms20 = tmp_int_Ms20
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms32:
		var tmp_int_Ms32 int64
		if tmp_int_Ms32, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Ms32", err)
		}
		ie.Ms32 = tmp_int_Ms32
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms64:
		var tmp_int_Ms64 int64
		if tmp_int_Ms64, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Ms64", err)
		}
		ie.Ms64 = tmp_int_Ms64
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms128:
		var tmp_int_Ms128 int64
		if tmp_int_Ms128, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode Ms128", err)
		}
		ie.Ms128 = tmp_int_Ms128
	case MTCH_SSB_MappingWindowCycleOffset_r17_Choice_Ms256:
		var tmp_int_Ms256 int64
		if tmp_int_Ms256, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode Ms256", err)
		}
		ie.Ms256 = tmp_int_Ms256
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
