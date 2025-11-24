package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RSSI_PeriodicityAndOffset_r16_Choice_nothing uint64 = iota
	RSSI_PeriodicityAndOffset_r16_Choice_Sl10
	RSSI_PeriodicityAndOffset_r16_Choice_Sl20
	RSSI_PeriodicityAndOffset_r16_Choice_Sl40
	RSSI_PeriodicityAndOffset_r16_Choice_Sl80
	RSSI_PeriodicityAndOffset_r16_Choice_Sl160
	RSSI_PeriodicityAndOffset_r16_Choice_Sl320
	RSSI_PeriodicityAndOffset_r16_Choice_S1640
)

type RSSI_PeriodicityAndOffset_r16 struct {
	Choice uint64
	Sl10   int64 `lb:0,ub:9,madatory`
	Sl20   int64 `lb:0,ub:19,madatory`
	Sl40   int64 `lb:0,ub:39,madatory`
	Sl80   int64 `lb:0,ub:79,madatory`
	Sl160  int64 `lb:0,ub:159,madatory`
	Sl320  int64 `lb:0,ub:319,madatory`
	S1640  int64 `lb:0,ub:639,madatory`
}

func (ie *RSSI_PeriodicityAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl10:
		if err = w.WriteInteger(int64(ie.Sl10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sl10", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl20:
		if err = w.WriteInteger(int64(ie.Sl20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sl20", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl40:
		if err = w.WriteInteger(int64(ie.Sl40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sl40", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl80:
		if err = w.WriteInteger(int64(ie.Sl80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sl80", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl160:
		if err = w.WriteInteger(int64(ie.Sl160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sl160", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl320:
		if err = w.WriteInteger(int64(ie.Sl320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Sl320", err)
		}
	case RSSI_PeriodicityAndOffset_r16_Choice_S1640:
		if err = w.WriteInteger(int64(ie.S1640), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode S1640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RSSI_PeriodicityAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(7, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl10:
		var tmp_int_Sl10 int64
		if tmp_int_Sl10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sl10", err)
		}
		ie.Sl10 = tmp_int_Sl10
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl20:
		var tmp_int_Sl20 int64
		if tmp_int_Sl20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sl20", err)
		}
		ie.Sl20 = tmp_int_Sl20
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl40:
		var tmp_int_Sl40 int64
		if tmp_int_Sl40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sl40", err)
		}
		ie.Sl40 = tmp_int_Sl40
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl80:
		var tmp_int_Sl80 int64
		if tmp_int_Sl80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sl80", err)
		}
		ie.Sl80 = tmp_int_Sl80
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl160:
		var tmp_int_Sl160 int64
		if tmp_int_Sl160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sl160", err)
		}
		ie.Sl160 = tmp_int_Sl160
	case RSSI_PeriodicityAndOffset_r16_Choice_Sl320:
		var tmp_int_Sl320 int64
		if tmp_int_Sl320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Sl320", err)
		}
		ie.Sl320 = tmp_int_Sl320
	case RSSI_PeriodicityAndOffset_r16_Choice_S1640:
		var tmp_int_S1640 int64
		if tmp_int_S1640, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode S1640", err)
		}
		ie.S1640 = tmp_int_S1640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
