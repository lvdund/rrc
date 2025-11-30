package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_nothing uint64 = iota
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N32
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N64
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N80
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N128
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N160
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N256
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N320
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N512
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N640
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N1280
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N2560
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N5120
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N10240
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N20480
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40960
	DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N81920
)

type DL_PPW_PeriodicityAndStartSlot_r17_scs120 struct {
	Choice uint64
	N32    int64 `lb:0,ub:31,madatory`
	N40    int64 `lb:0,ub:39,madatory`
	N64    int64 `lb:0,ub:63,madatory`
	N80    int64 `lb:0,ub:79,madatory`
	N128   int64 `lb:0,ub:127,madatory`
	N160   int64 `lb:0,ub:159,madatory`
	N256   int64 `lb:0,ub:255,madatory`
	N320   int64 `lb:0,ub:319,madatory`
	N512   int64 `lb:0,ub:511,madatory`
	N640   int64 `lb:0,ub:639,madatory`
	N1280  int64 `lb:0,ub:1279,madatory`
	N2560  int64 `lb:0,ub:2559,madatory`
	N5120  int64 `lb:0,ub:5119,madatory`
	N10240 int64 `lb:0,ub:10239,madatory`
	N20480 int64 `lb:0,ub:20479,madatory`
	N40960 int64 `lb:0,ub:40959,madatory`
	N81920 int64 `lb:0,ub:81919,madatory`
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs120) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N32:
		if err = w.WriteInteger(int64(ie.N32), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode N32", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40:
		if err = w.WriteInteger(int64(ie.N40), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode N40", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N64:
		if err = w.WriteInteger(int64(ie.N64), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode N64", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N80:
		if err = w.WriteInteger(int64(ie.N80), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode N80", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N128:
		if err = w.WriteInteger(int64(ie.N128), &aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode N128", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N160:
		if err = w.WriteInteger(int64(ie.N160), &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode N160", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N256:
		if err = w.WriteInteger(int64(ie.N256), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode N256", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N320:
		if err = w.WriteInteger(int64(ie.N320), &aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode N320", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N512:
		if err = w.WriteInteger(int64(ie.N512), &aper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode N512", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N640:
		if err = w.WriteInteger(int64(ie.N640), &aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode N640", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N1280:
		if err = w.WriteInteger(int64(ie.N1280), &aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode N1280", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N2560:
		if err = w.WriteInteger(int64(ie.N2560), &aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode N2560", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N5120:
		if err = w.WriteInteger(int64(ie.N5120), &aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode N5120", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N10240:
		if err = w.WriteInteger(int64(ie.N10240), &aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode N10240", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N20480:
		if err = w.WriteInteger(int64(ie.N20480), &aper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode N20480", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40960:
		if err = w.WriteInteger(int64(ie.N40960), &aper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			err = utils.WrapError("Encode N40960", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N81920:
		if err = w.WriteInteger(int64(ie.N81920), &aper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			err = utils.WrapError("Encode N81920", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs120) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N32:
		var tmp_int_N32 int64
		if tmp_int_N32, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode N32", err)
		}
		ie.N32 = tmp_int_N32
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40:
		var tmp_int_N40 int64
		if tmp_int_N40, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode N40", err)
		}
		ie.N40 = tmp_int_N40
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N64:
		var tmp_int_N64 int64
		if tmp_int_N64, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode N64", err)
		}
		ie.N64 = tmp_int_N64
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N80:
		var tmp_int_N80 int64
		if tmp_int_N80, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode N80", err)
		}
		ie.N80 = tmp_int_N80
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N128:
		var tmp_int_N128 int64
		if tmp_int_N128, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode N128", err)
		}
		ie.N128 = tmp_int_N128
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N160:
		var tmp_int_N160 int64
		if tmp_int_N160, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode N160", err)
		}
		ie.N160 = tmp_int_N160
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N256:
		var tmp_int_N256 int64
		if tmp_int_N256, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode N256", err)
		}
		ie.N256 = tmp_int_N256
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N320:
		var tmp_int_N320 int64
		if tmp_int_N320, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode N320", err)
		}
		ie.N320 = tmp_int_N320
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N512:
		var tmp_int_N512 int64
		if tmp_int_N512, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode N512", err)
		}
		ie.N512 = tmp_int_N512
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N640:
		var tmp_int_N640 int64
		if tmp_int_N640, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode N640", err)
		}
		ie.N640 = tmp_int_N640
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N1280:
		var tmp_int_N1280 int64
		if tmp_int_N1280, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode N1280", err)
		}
		ie.N1280 = tmp_int_N1280
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N2560:
		var tmp_int_N2560 int64
		if tmp_int_N2560, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode N2560", err)
		}
		ie.N2560 = tmp_int_N2560
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N5120:
		var tmp_int_N5120 int64
		if tmp_int_N5120, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode N5120", err)
		}
		ie.N5120 = tmp_int_N5120
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N10240:
		var tmp_int_N10240 int64
		if tmp_int_N10240, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode N10240", err)
		}
		ie.N10240 = tmp_int_N10240
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N20480:
		var tmp_int_N20480 int64
		if tmp_int_N20480, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode N20480", err)
		}
		ie.N20480 = tmp_int_N20480
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N40960:
		var tmp_int_N40960 int64
		if tmp_int_N40960, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode N40960", err)
		}
		ie.N40960 = tmp_int_N40960
	case DL_PPW_PeriodicityAndStartSlot_r17_scs120_Choice_N81920:
		var tmp_int_N81920 int64
		if tmp_int_N81920, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			return utils.WrapError("Decode N81920", err)
		}
		ie.N81920 = tmp_int_N81920
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
