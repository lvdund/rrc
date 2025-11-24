package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PeriodicityAndOffset_Choice_nothing uint64 = iota
	SRS_PeriodicityAndOffset_Choice_Sl1
	SRS_PeriodicityAndOffset_Choice_Sl2
	SRS_PeriodicityAndOffset_Choice_Sl4
	SRS_PeriodicityAndOffset_Choice_Sl5
	SRS_PeriodicityAndOffset_Choice_Sl8
	SRS_PeriodicityAndOffset_Choice_Sl10
	SRS_PeriodicityAndOffset_Choice_Sl16
	SRS_PeriodicityAndOffset_Choice_Sl20
	SRS_PeriodicityAndOffset_Choice_Sl32
	SRS_PeriodicityAndOffset_Choice_Sl40
	SRS_PeriodicityAndOffset_Choice_Sl64
	SRS_PeriodicityAndOffset_Choice_Sl80
	SRS_PeriodicityAndOffset_Choice_Sl160
	SRS_PeriodicityAndOffset_Choice_Sl320
	SRS_PeriodicityAndOffset_Choice_Sl640
	SRS_PeriodicityAndOffset_Choice_Sl1280
	SRS_PeriodicityAndOffset_Choice_Sl2560
)

type SRS_PeriodicityAndOffset struct {
	Choice uint64
	Sl1    uper.NULL `madatory`
	Sl2    int64     `lb:0,ub:1,madatory`
	Sl4    int64     `lb:0,ub:3,madatory`
	Sl5    int64     `lb:0,ub:4,madatory`
	Sl8    int64     `lb:0,ub:7,madatory`
	Sl10   int64     `lb:0,ub:9,madatory`
	Sl16   int64     `lb:0,ub:15,madatory`
	Sl20   int64     `lb:0,ub:19,madatory`
	Sl32   int64     `lb:0,ub:31,madatory`
	Sl40   int64     `lb:0,ub:39,madatory`
	Sl64   int64     `lb:0,ub:63,madatory`
	Sl80   int64     `lb:0,ub:79,madatory`
	Sl160  int64     `lb:0,ub:159,madatory`
	Sl320  int64     `lb:0,ub:319,madatory`
	Sl640  int64     `lb:0,ub:639,madatory`
	Sl1280 int64     `lb:0,ub:1279,madatory`
	Sl2560 int64     `lb:0,ub:2559,madatory`
}

func (ie *SRS_PeriodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_Choice_Sl1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Sl1", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl2:
		if err = w.WriteInteger(int64(ie.Sl2), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode Sl2", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl4:
		if err = w.WriteInteger(int64(ie.Sl4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Sl4", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl5:
		if err = w.WriteInteger(int64(ie.Sl5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sl5", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl8:
		if err = w.WriteInteger(int64(ie.Sl8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Sl8", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl10:
		if err = w.WriteInteger(int64(ie.Sl10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sl10", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl16:
		if err = w.WriteInteger(int64(ie.Sl16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Sl16", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl20:
		if err = w.WriteInteger(int64(ie.Sl20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sl20", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl32:
		if err = w.WriteInteger(int64(ie.Sl32), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Sl32", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl40:
		if err = w.WriteInteger(int64(ie.Sl40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sl40", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl64:
		if err = w.WriteInteger(int64(ie.Sl64), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Sl64", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl80:
		if err = w.WriteInteger(int64(ie.Sl80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sl80", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl160:
		if err = w.WriteInteger(int64(ie.Sl160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sl160", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl320:
		if err = w.WriteInteger(int64(ie.Sl320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Sl320", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl640:
		if err = w.WriteInteger(int64(ie.Sl640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Sl640", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl1280:
		if err = w.WriteInteger(int64(ie.Sl1280), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Sl1280", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl2560:
		if err = w.WriteInteger(int64(ie.Sl2560), &uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode Sl2560", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PeriodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(17, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_Choice_Sl1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Sl1", err)
		}
	case SRS_PeriodicityAndOffset_Choice_Sl2:
		var tmp_int_Sl2 int64
		if tmp_int_Sl2, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl2", err)
		}
		ie.Sl2 = tmp_int_Sl2
	case SRS_PeriodicityAndOffset_Choice_Sl4:
		var tmp_int_Sl4 int64
		if tmp_int_Sl4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Sl4", err)
		}
		ie.Sl4 = tmp_int_Sl4
	case SRS_PeriodicityAndOffset_Choice_Sl5:
		var tmp_int_Sl5 int64
		if tmp_int_Sl5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sl5", err)
		}
		ie.Sl5 = tmp_int_Sl5
	case SRS_PeriodicityAndOffset_Choice_Sl8:
		var tmp_int_Sl8 int64
		if tmp_int_Sl8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Sl8", err)
		}
		ie.Sl8 = tmp_int_Sl8
	case SRS_PeriodicityAndOffset_Choice_Sl10:
		var tmp_int_Sl10 int64
		if tmp_int_Sl10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl10", err)
		}
		ie.Sl10 = tmp_int_Sl10
	case SRS_PeriodicityAndOffset_Choice_Sl16:
		var tmp_int_Sl16 int64
		if tmp_int_Sl16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl16", err)
		}
		ie.Sl16 = tmp_int_Sl16
	case SRS_PeriodicityAndOffset_Choice_Sl20:
		var tmp_int_Sl20 int64
		if tmp_int_Sl20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sl20", err)
		}
		ie.Sl20 = tmp_int_Sl20
	case SRS_PeriodicityAndOffset_Choice_Sl32:
		var tmp_int_Sl32 int64
		if tmp_int_Sl32, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Sl32", err)
		}
		ie.Sl32 = tmp_int_Sl32
	case SRS_PeriodicityAndOffset_Choice_Sl40:
		var tmp_int_Sl40 int64
		if tmp_int_Sl40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sl40", err)
		}
		ie.Sl40 = tmp_int_Sl40
	case SRS_PeriodicityAndOffset_Choice_Sl64:
		var tmp_int_Sl64 int64
		if tmp_int_Sl64, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Sl64", err)
		}
		ie.Sl64 = tmp_int_Sl64
	case SRS_PeriodicityAndOffset_Choice_Sl80:
		var tmp_int_Sl80 int64
		if tmp_int_Sl80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sl80", err)
		}
		ie.Sl80 = tmp_int_Sl80
	case SRS_PeriodicityAndOffset_Choice_Sl160:
		var tmp_int_Sl160 int64
		if tmp_int_Sl160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sl160", err)
		}
		ie.Sl160 = tmp_int_Sl160
	case SRS_PeriodicityAndOffset_Choice_Sl320:
		var tmp_int_Sl320 int64
		if tmp_int_Sl320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Sl320", err)
		}
		ie.Sl320 = tmp_int_Sl320
	case SRS_PeriodicityAndOffset_Choice_Sl640:
		var tmp_int_Sl640 int64
		if tmp_int_Sl640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Sl640", err)
		}
		ie.Sl640 = tmp_int_Sl640
	case SRS_PeriodicityAndOffset_Choice_Sl1280:
		var tmp_int_Sl1280 int64
		if tmp_int_Sl1280, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Sl1280", err)
		}
		ie.Sl1280 = tmp_int_Sl1280
	case SRS_PeriodicityAndOffset_Choice_Sl2560:
		var tmp_int_Sl2560 int64
		if tmp_int_Sl2560, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Sl2560", err)
		}
		ie.Sl2560 = tmp_int_Sl2560
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
