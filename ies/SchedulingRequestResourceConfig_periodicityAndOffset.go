package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_nothing uint64 = iota
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym2
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym6or7
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl1
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl2
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl4
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl5
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl8
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl10
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl16
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl20
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl40
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl80
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl160
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl320
	SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl640
)

type SchedulingRequestResourceConfig_periodicityAndOffset struct {
	Choice  uint64
	Sym2    aper.NULL `madatory`
	Sym6or7 aper.NULL `madatory`
	Sl1     aper.NULL `madatory`
	Sl2     int64     `lb:0,ub:1,madatory`
	Sl4     int64     `lb:0,ub:3,madatory`
	Sl5     int64     `lb:0,ub:4,madatory`
	Sl8     int64     `lb:0,ub:7,madatory`
	Sl10    int64     `lb:0,ub:9,madatory`
	Sl16    int64     `lb:0,ub:15,madatory`
	Sl20    int64     `lb:0,ub:19,madatory`
	Sl40    int64     `lb:0,ub:39,madatory`
	Sl80    int64     `lb:0,ub:79,madatory`
	Sl160   int64     `lb:0,ub:159,madatory`
	Sl320   int64     `lb:0,ub:319,madatory`
	Sl640   int64     `lb:0,ub:639,madatory`
}

func (ie *SchedulingRequestResourceConfig_periodicityAndOffset) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 15, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym2:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Sym2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym6or7:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Sym6or7", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl1:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode Sl1", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl2:
		if err = w.WriteInteger(int64(ie.Sl2), &aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode Sl2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl4:
		if err = w.WriteInteger(int64(ie.Sl4), &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Sl4", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl5:
		if err = w.WriteInteger(int64(ie.Sl5), &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sl5", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl8:
		if err = w.WriteInteger(int64(ie.Sl8), &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Sl8", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl10:
		if err = w.WriteInteger(int64(ie.Sl10), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sl10", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl16:
		if err = w.WriteInteger(int64(ie.Sl16), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Sl16", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl20:
		if err = w.WriteInteger(int64(ie.Sl20), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sl20", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl40:
		if err = w.WriteInteger(int64(ie.Sl40), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sl40", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl80:
		if err = w.WriteInteger(int64(ie.Sl80), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sl80", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl160:
		if err = w.WriteInteger(int64(ie.Sl160), &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sl160", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl320:
		if err = w.WriteInteger(int64(ie.Sl320), &aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Sl320", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl640:
		if err = w.WriteInteger(int64(ie.Sl640), &aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Sl640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SchedulingRequestResourceConfig_periodicityAndOffset) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(15, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym2:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Sym2", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sym6or7:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Sym6or7", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl1:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode Sl1", err)
		}
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl2:
		var tmp_int_Sl2 int64
		if tmp_int_Sl2, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode Sl2", err)
		}
		ie.Sl2 = tmp_int_Sl2
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl4:
		var tmp_int_Sl4 int64
		if tmp_int_Sl4, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Sl4", err)
		}
		ie.Sl4 = tmp_int_Sl4
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl5:
		var tmp_int_Sl5 int64
		if tmp_int_Sl5, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sl5", err)
		}
		ie.Sl5 = tmp_int_Sl5
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl8:
		var tmp_int_Sl8 int64
		if tmp_int_Sl8, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Sl8", err)
		}
		ie.Sl8 = tmp_int_Sl8
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl10:
		var tmp_int_Sl10 int64
		if tmp_int_Sl10, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl10", err)
		}
		ie.Sl10 = tmp_int_Sl10
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl16:
		var tmp_int_Sl16 int64
		if tmp_int_Sl16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sl16", err)
		}
		ie.Sl16 = tmp_int_Sl16
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl20:
		var tmp_int_Sl20 int64
		if tmp_int_Sl20, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sl20", err)
		}
		ie.Sl20 = tmp_int_Sl20
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl40:
		var tmp_int_Sl40 int64
		if tmp_int_Sl40, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sl40", err)
		}
		ie.Sl40 = tmp_int_Sl40
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl80:
		var tmp_int_Sl80 int64
		if tmp_int_Sl80, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sl80", err)
		}
		ie.Sl80 = tmp_int_Sl80
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl160:
		var tmp_int_Sl160 int64
		if tmp_int_Sl160, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sl160", err)
		}
		ie.Sl160 = tmp_int_Sl160
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl320:
		var tmp_int_Sl320 int64
		if tmp_int_Sl320, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Sl320", err)
		}
		ie.Sl320 = tmp_int_Sl320
	case SchedulingRequestResourceConfig_periodicityAndOffset_Choice_Sl640:
		var tmp_int_Sl640 int64
		if tmp_int_Sl640, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Sl640", err)
		}
		ie.Sl640 = tmp_int_Sl640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
