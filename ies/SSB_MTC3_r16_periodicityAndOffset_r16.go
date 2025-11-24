package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_nothing uint64 = iota
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf5_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf10_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf20_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf40_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf80_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf160_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf320_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf640_r16
	SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf1280_r16
)

type SSB_MTC3_r16_periodicityAndOffset_r16 struct {
	Choice     uint64
	Sf5_r16    int64 `lb:0,ub:4,madatory`
	Sf10_r16   int64 `lb:0,ub:9,madatory`
	Sf20_r16   int64 `lb:0,ub:19,madatory`
	Sf40_r16   int64 `lb:0,ub:39,madatory`
	Sf80_r16   int64 `lb:0,ub:79,madatory`
	Sf160_r16  int64 `lb:0,ub:159,madatory`
	Sf320_r16  int64 `lb:0,ub:319,madatory`
	Sf640_r16  int64 `lb:0,ub:639,madatory`
	Sf1280_r16 int64 `lb:0,ub:1279,madatory`
}

func (ie *SSB_MTC3_r16_periodicityAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf5_r16:
		if err = w.WriteInteger(int64(ie.Sf5_r16), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sf5_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf10_r16:
		if err = w.WriteInteger(int64(ie.Sf10_r16), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sf10_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf20_r16:
		if err = w.WriteInteger(int64(ie.Sf20_r16), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sf20_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf40_r16:
		if err = w.WriteInteger(int64(ie.Sf40_r16), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sf40_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf80_r16:
		if err = w.WriteInteger(int64(ie.Sf80_r16), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sf80_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf160_r16:
		if err = w.WriteInteger(int64(ie.Sf160_r16), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sf160_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf320_r16:
		if err = w.WriteInteger(int64(ie.Sf320_r16), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Sf320_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf640_r16:
		if err = w.WriteInteger(int64(ie.Sf640_r16), &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Sf640_r16", err)
		}
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf1280_r16:
		if err = w.WriteInteger(int64(ie.Sf1280_r16), &uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			err = utils.WrapError("Encode Sf1280_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_MTC3_r16_periodicityAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf5_r16:
		var tmp_int_Sf5_r16 int64
		if tmp_int_Sf5_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sf5_r16", err)
		}
		ie.Sf5_r16 = tmp_int_Sf5_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf10_r16:
		var tmp_int_Sf10_r16 int64
		if tmp_int_Sf10_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sf10_r16", err)
		}
		ie.Sf10_r16 = tmp_int_Sf10_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf20_r16:
		var tmp_int_Sf20_r16 int64
		if tmp_int_Sf20_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sf20_r16", err)
		}
		ie.Sf20_r16 = tmp_int_Sf20_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf40_r16:
		var tmp_int_Sf40_r16 int64
		if tmp_int_Sf40_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sf40_r16", err)
		}
		ie.Sf40_r16 = tmp_int_Sf40_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf80_r16:
		var tmp_int_Sf80_r16 int64
		if tmp_int_Sf80_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sf80_r16", err)
		}
		ie.Sf80_r16 = tmp_int_Sf80_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf160_r16:
		var tmp_int_Sf160_r16 int64
		if tmp_int_Sf160_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sf160_r16", err)
		}
		ie.Sf160_r16 = tmp_int_Sf160_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf320_r16:
		var tmp_int_Sf320_r16 int64
		if tmp_int_Sf320_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Sf320_r16", err)
		}
		ie.Sf320_r16 = tmp_int_Sf320_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf640_r16:
		var tmp_int_Sf640_r16 int64
		if tmp_int_Sf640_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Sf640_r16", err)
		}
		ie.Sf640_r16 = tmp_int_Sf640_r16
	case SSB_MTC3_r16_periodicityAndOffset_r16_Choice_Sf1280_r16:
		var tmp_int_Sf1280_r16 int64
		if tmp_int_Sf1280_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1279}, false); err != nil {
			return utils.WrapError("Decode Sf1280_r16", err)
		}
		ie.Sf1280_r16 = tmp_int_Sf1280_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
