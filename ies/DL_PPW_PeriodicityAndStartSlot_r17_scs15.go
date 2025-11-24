package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_nothing uint64 = iota
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N4
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N8
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N16
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N20
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N32
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N40
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N64
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N80
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N160
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N320
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N640
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N1280
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N2560
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5120
	DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10240
)

type DL_PPW_PeriodicityAndStartSlot_r17_scs15 struct {
	Choice uint64
	N4     int64 `lb:0,ub:3,madatory`
	N5     int64 `lb:0,ub:4,madatory`
	N8     int64 `lb:0,ub:7,madatory`
	N10    int64 `lb:0,ub:9,madatory`
	N16    int64 `lb:0,ub:15,madatory`
	N20    int64 `lb:0,ub:19,madatory`
	N32    int64 `lb:0,ub:31,madatory`
	N40    int64 `lb:0,ub:39,madatory`
	N64    int64 `lb:0,ub:63,madatory`
	N80    int64 `lb:0,ub:79,madatory`
	N160   int64 `lb:0,ub:159,madatory`
	N320   int64 `lb:0,ub:319,madatory`
	N640   int64 `lb:0,ub:639,madatory`
	N1280  int64 `lb:0,ub:1279,madatory`
	N2560  int64 `lb:0,ub:2559,madatory`
	N5120  int64 `lb:0,ub:5119,madatory`
	N10240 int64 `lb:0,ub:10239,madatory`
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs15) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N4:
		if err = w.WriteInteger(int64(ie.N4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode N4", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5:
		if err = w.WriteInteger(int64(ie.N5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode N5", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N8:
		if err = w.WriteInteger(int64(ie.N8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode N8", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10:
		if err = w.WriteInteger(int64(ie.N10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode N10", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N16:
		if err = w.WriteInteger(int64(ie.N16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode N16", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N20:
		if err = w.WriteInteger(int64(ie.N20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode N20", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N32:
		if err = w.WriteInteger(int64(ie.N32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode N32", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N40:
		if err = w.WriteInteger(int64(ie.N40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode N40", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N64:
		if err = w.WriteInteger(int64(ie.N64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode N64", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N80:
		if err = w.WriteInteger(int64(ie.N80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode N80", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N160:
		if err = w.WriteInteger(int64(ie.N160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode N160", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N320:
		if err = w.WriteInteger(int64(ie.N320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode N320", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N640:
		if err = w.WriteInteger(int64(ie.N640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode N640", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N1280:
		if err = w.WriteInteger(int64(ie.N1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode N1280", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N2560:
		if err = w.WriteInteger(int64(ie.N2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode N2560", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5120:
		if err = w.WriteInteger(int64(ie.N5120), &uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode N5120", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10240:
		if err = w.WriteInteger(int64(ie.N10240), &uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode N10240", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17_scs15) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N4:
		var tmp_int_N4 int64
		if tmp_int_N4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode N4", err)
		}
		ie.N4 = tmp_int_N4
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5:
		var tmp_int_N5 int64
		if tmp_int_N5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode N5", err)
		}
		ie.N5 = tmp_int_N5
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N8:
		var tmp_int_N8 int64
		if tmp_int_N8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode N8", err)
		}
		ie.N8 = tmp_int_N8
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10:
		var tmp_int_N10 int64
		if tmp_int_N10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode N10", err)
		}
		ie.N10 = tmp_int_N10
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N16:
		var tmp_int_N16 int64
		if tmp_int_N16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode N16", err)
		}
		ie.N16 = tmp_int_N16
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N20:
		var tmp_int_N20 int64
		if tmp_int_N20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode N20", err)
		}
		ie.N20 = tmp_int_N20
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N32:
		var tmp_int_N32 int64
		if tmp_int_N32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode N32", err)
		}
		ie.N32 = tmp_int_N32
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N40:
		var tmp_int_N40 int64
		if tmp_int_N40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode N40", err)
		}
		ie.N40 = tmp_int_N40
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N64:
		var tmp_int_N64 int64
		if tmp_int_N64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode N64", err)
		}
		ie.N64 = tmp_int_N64
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N80:
		var tmp_int_N80 int64
		if tmp_int_N80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode N80", err)
		}
		ie.N80 = tmp_int_N80
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N160:
		var tmp_int_N160 int64
		if tmp_int_N160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode N160", err)
		}
		ie.N160 = tmp_int_N160
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N320:
		var tmp_int_N320 int64
		if tmp_int_N320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode N320", err)
		}
		ie.N320 = tmp_int_N320
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N640:
		var tmp_int_N640 int64
		if tmp_int_N640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode N640", err)
		}
		ie.N640 = tmp_int_N640
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N1280:
		var tmp_int_N1280 int64
		if tmp_int_N1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode N1280", err)
		}
		ie.N1280 = tmp_int_N1280
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N2560:
		var tmp_int_N2560 int64
		if tmp_int_N2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode N2560", err)
		}
		ie.N2560 = tmp_int_N2560
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N5120:
		var tmp_int_N5120 int64
		if tmp_int_N5120, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode N5120", err)
		}
		ie.N5120 = tmp_int_N5120
	case DL_PPW_PeriodicityAndStartSlot_r17_scs15_Choice_N10240:
		var tmp_int_N10240 int64
		if tmp_int_N10240, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode N10240", err)
		}
		ie.N10240 = tmp_int_N10240
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
