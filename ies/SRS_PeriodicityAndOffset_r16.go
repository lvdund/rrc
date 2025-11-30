package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PeriodicityAndOffset_r16_Choice_nothing uint64 = iota
	SRS_PeriodicityAndOffset_r16_Choice_Sl1
	SRS_PeriodicityAndOffset_r16_Choice_Sl2
	SRS_PeriodicityAndOffset_r16_Choice_Sl4
	SRS_PeriodicityAndOffset_r16_Choice_Sl5
	SRS_PeriodicityAndOffset_r16_Choice_Sl8
	SRS_PeriodicityAndOffset_r16_Choice_Sl10
	SRS_PeriodicityAndOffset_r16_Choice_Sl16
	SRS_PeriodicityAndOffset_r16_Choice_Sl20
	SRS_PeriodicityAndOffset_r16_Choice_Sl32
	SRS_PeriodicityAndOffset_r16_Choice_Sl40
	SRS_PeriodicityAndOffset_r16_Choice_Sl64
	SRS_PeriodicityAndOffset_r16_Choice_Sl80
	SRS_PeriodicityAndOffset_r16_Choice_Sl160
	SRS_PeriodicityAndOffset_r16_Choice_Sl320
	SRS_PeriodicityAndOffset_r16_Choice_Sl640
	SRS_PeriodicityAndOffset_r16_Choice_Sl1280
	SRS_PeriodicityAndOffset_r16_Choice_Sl2560
	SRS_PeriodicityAndOffset_r16_Choice_Sl5120
	SRS_PeriodicityAndOffset_r16_Choice_Sl10240
	SRS_PeriodicityAndOffset_r16_Choice_Sl40960
	SRS_PeriodicityAndOffset_r16_Choice_Sl81920
)

type SRS_PeriodicityAndOffset_r16 struct {
	Choice  uint64
	Sl1     aper.NULL `madatory`
	Sl2     int64     `lb:0,ub:1,madatory`
	Sl4     int64     `lb:0,ub:3,madatory`
	Sl5     int64     `lb:0,ub:4,madatory`
	Sl8     int64     `lb:0,ub:7,madatory`
	Sl10    int64     `lb:0,ub:9,madatory`
	Sl16    int64     `lb:0,ub:15,madatory`
	Sl20    int64     `lb:0,ub:19,madatory`
	Sl32    int64     `lb:0,ub:31,madatory`
	Sl40    int64     `lb:0,ub:39,madatory`
	Sl64    int64     `lb:0,ub:63,madatory`
	Sl80    int64     `lb:0,ub:79,madatory`
	Sl160   int64     `lb:0,ub:159,madatory`
	Sl320   int64     `lb:0,ub:319,madatory`
	Sl640   int64     `lb:0,ub:639,madatory`
	Sl1280  int64     `lb:0,ub:1279,madatory`
	Sl2560  int64     `lb:0,ub:2559,madatory`
	Sl5120  int64     `lb:0,ub:5119,madatory`
	Sl10240 int64     `lb:0,ub:10239,madatory`
	Sl40960 int64     `lb:0,ub:40959,madatory`
	Sl81920 int64     `lb:0,ub:81919,madatory`
}

func (ie *SRS_PeriodicityAndOffset_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 21, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_r16_Choice_Sl1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Sl1", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl2:
		if err = w.WriteInteger(int64(ie.Sl2), &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode Sl2", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl4:
		if err = w.WriteInteger(int64(ie.Sl4), &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Sl4", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl5:
		if err = w.WriteInteger(int64(ie.Sl5), &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sl5", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl8:
		if err = w.WriteInteger(int64(ie.Sl8), &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Sl8", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl10:
		if err = w.WriteInteger(int64(ie.Sl10), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sl10", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl16:
		if err = w.WriteInteger(int64(ie.Sl16), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Sl16", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl20:
		if err = w.WriteInteger(int64(ie.Sl20), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sl20", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl32:
		if err = w.WriteInteger(int64(ie.Sl32), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Sl32", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl40:
		if err = w.WriteInteger(int64(ie.Sl40), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sl40", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl64:
		if err = w.WriteInteger(int64(ie.Sl64), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Sl64", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl80:
		if err = w.WriteInteger(int64(ie.Sl80), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sl80", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl160:
		if err = w.WriteInteger(int64(ie.Sl160), &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sl160", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl320:
		if err = w.WriteInteger(int64(ie.Sl320), &aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Sl320", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl640:
		if err = w.WriteInteger(int64(ie.Sl640), &aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Sl640", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl1280:
		if err = w.WriteInteger(int64(ie.Sl1280), &aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Sl1280", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl2560:
		if err = w.WriteInteger(int64(ie.Sl2560), &aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			err = utils.WrapError("Encode Sl2560", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl5120:
		if err = w.WriteInteger(int64(ie.Sl5120), &aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			err = utils.WrapError("Encode Sl5120", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl10240:
		if err = w.WriteInteger(int64(ie.Sl10240), &aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			err = utils.WrapError("Encode Sl10240", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl40960:
		if err = w.WriteInteger(int64(ie.Sl40960), &aper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			err = utils.WrapError("Encode Sl40960", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl81920:
		if err = w.WriteInteger(int64(ie.Sl81920), &aper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			err = utils.WrapError("Encode Sl81920", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PeriodicityAndOffset_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(21, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffset_r16_Choice_Sl1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Sl1", err)
		}
	case SRS_PeriodicityAndOffset_r16_Choice_Sl2:
		var tmp_int_Sl2 int64
		if tmp_int_Sl2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl2", err)
		}
		ie.Sl2 = tmp_int_Sl2
	case SRS_PeriodicityAndOffset_r16_Choice_Sl4:
		var tmp_int_Sl4 int64
		if tmp_int_Sl4, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Sl4", err)
		}
		ie.Sl4 = tmp_int_Sl4
	case SRS_PeriodicityAndOffset_r16_Choice_Sl5:
		var tmp_int_Sl5 int64
		if tmp_int_Sl5, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sl5", err)
		}
		ie.Sl5 = tmp_int_Sl5
	case SRS_PeriodicityAndOffset_r16_Choice_Sl8:
		var tmp_int_Sl8 int64
		if tmp_int_Sl8, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Sl8", err)
		}
		ie.Sl8 = tmp_int_Sl8
	case SRS_PeriodicityAndOffset_r16_Choice_Sl10:
		var tmp_int_Sl10 int64
		if tmp_int_Sl10, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl10", err)
		}
		ie.Sl10 = tmp_int_Sl10
	case SRS_PeriodicityAndOffset_r16_Choice_Sl16:
		var tmp_int_Sl16 int64
		if tmp_int_Sl16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl16", err)
		}
		ie.Sl16 = tmp_int_Sl16
	case SRS_PeriodicityAndOffset_r16_Choice_Sl20:
		var tmp_int_Sl20 int64
		if tmp_int_Sl20, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sl20", err)
		}
		ie.Sl20 = tmp_int_Sl20
	case SRS_PeriodicityAndOffset_r16_Choice_Sl32:
		var tmp_int_Sl32 int64
		if tmp_int_Sl32, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Sl32", err)
		}
		ie.Sl32 = tmp_int_Sl32
	case SRS_PeriodicityAndOffset_r16_Choice_Sl40:
		var tmp_int_Sl40 int64
		if tmp_int_Sl40, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sl40", err)
		}
		ie.Sl40 = tmp_int_Sl40
	case SRS_PeriodicityAndOffset_r16_Choice_Sl64:
		var tmp_int_Sl64 int64
		if tmp_int_Sl64, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Sl64", err)
		}
		ie.Sl64 = tmp_int_Sl64
	case SRS_PeriodicityAndOffset_r16_Choice_Sl80:
		var tmp_int_Sl80 int64
		if tmp_int_Sl80, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sl80", err)
		}
		ie.Sl80 = tmp_int_Sl80
	case SRS_PeriodicityAndOffset_r16_Choice_Sl160:
		var tmp_int_Sl160 int64
		if tmp_int_Sl160, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sl160", err)
		}
		ie.Sl160 = tmp_int_Sl160
	case SRS_PeriodicityAndOffset_r16_Choice_Sl320:
		var tmp_int_Sl320 int64
		if tmp_int_Sl320, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Sl320", err)
		}
		ie.Sl320 = tmp_int_Sl320
	case SRS_PeriodicityAndOffset_r16_Choice_Sl640:
		var tmp_int_Sl640 int64
		if tmp_int_Sl640, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Sl640", err)
		}
		ie.Sl640 = tmp_int_Sl640
	case SRS_PeriodicityAndOffset_r16_Choice_Sl1280:
		var tmp_int_Sl1280 int64
		if tmp_int_Sl1280, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Sl1280", err)
		}
		ie.Sl1280 = tmp_int_Sl1280
	case SRS_PeriodicityAndOffset_r16_Choice_Sl2560:
		var tmp_int_Sl2560 int64
		if tmp_int_Sl2560, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode Sl2560", err)
		}
		ie.Sl2560 = tmp_int_Sl2560
	case SRS_PeriodicityAndOffset_r16_Choice_Sl5120:
		var tmp_int_Sl5120 int64
		if tmp_int_Sl5120, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 5119}, false); err != nil {
			return utils.WrapError("Decode Sl5120", err)
		}
		ie.Sl5120 = tmp_int_Sl5120
	case SRS_PeriodicityAndOffset_r16_Choice_Sl10240:
		var tmp_int_Sl10240 int64
		if tmp_int_Sl10240, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 10239}, false); err != nil {
			return utils.WrapError("Decode Sl10240", err)
		}
		ie.Sl10240 = tmp_int_Sl10240
	case SRS_PeriodicityAndOffset_r16_Choice_Sl40960:
		var tmp_int_Sl40960 int64
		if tmp_int_Sl40960, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 40959}, false); err != nil {
			return utils.WrapError("Decode Sl40960", err)
		}
		ie.Sl40960 = tmp_int_Sl40960
	case SRS_PeriodicityAndOffset_r16_Choice_Sl81920:
		var tmp_int_Sl81920 int64
		if tmp_int_Sl81920, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 81919}, false); err != nil {
			return utils.WrapError("Decode Sl81920", err)
		}
		ie.Sl81920 = tmp_int_Sl81920
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
