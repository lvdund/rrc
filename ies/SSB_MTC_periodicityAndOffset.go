package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC_periodicityAndOffset_Choice_nothing uint64 = iota
	SSB_MTC_periodicityAndOffset_Choice_Sf5
	SSB_MTC_periodicityAndOffset_Choice_Sf10
	SSB_MTC_periodicityAndOffset_Choice_Sf20
	SSB_MTC_periodicityAndOffset_Choice_Sf40
	SSB_MTC_periodicityAndOffset_Choice_Sf80
	SSB_MTC_periodicityAndOffset_Choice_Sf160
)

type SSB_MTC_periodicityAndOffset struct {
	Choice uint64
	Sf5    int64 `lb:0,ub:4,madatory`
	Sf10   int64 `lb:0,ub:9,madatory`
	Sf20   int64 `lb:0,ub:19,madatory`
	Sf40   int64 `lb:0,ub:39,madatory`
	Sf80   int64 `lb:0,ub:79,madatory`
	Sf160  int64 `lb:0,ub:159,madatory`
}

func (ie *SSB_MTC_periodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC_periodicityAndOffset_Choice_Sf5:
		if err = w.WriteInteger(int64(ie.Sf5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sf5", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_Sf10:
		if err = w.WriteInteger(int64(ie.Sf10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Sf10", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_Sf20:
		if err = w.WriteInteger(int64(ie.Sf20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Sf20", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_Sf40:
		if err = w.WriteInteger(int64(ie.Sf40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Sf40", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_Sf80:
		if err = w.WriteInteger(int64(ie.Sf80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Sf80", err)
		}
	case SSB_MTC_periodicityAndOffset_Choice_Sf160:
		if err = w.WriteInteger(int64(ie.Sf160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Sf160", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SSB_MTC_periodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SSB_MTC_periodicityAndOffset_Choice_Sf5:
		var tmp_int_Sf5 int64
		if tmp_int_Sf5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sf5", err)
		}
		ie.Sf5 = tmp_int_Sf5
	case SSB_MTC_periodicityAndOffset_Choice_Sf10:
		var tmp_int_Sf10 int64
		if tmp_int_Sf10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Sf10", err)
		}
		ie.Sf10 = tmp_int_Sf10
	case SSB_MTC_periodicityAndOffset_Choice_Sf20:
		var tmp_int_Sf20 int64
		if tmp_int_Sf20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Sf20", err)
		}
		ie.Sf20 = tmp_int_Sf20
	case SSB_MTC_periodicityAndOffset_Choice_Sf40:
		var tmp_int_Sf40 int64
		if tmp_int_Sf40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Sf40", err)
		}
		ie.Sf40 = tmp_int_Sf40
	case SSB_MTC_periodicityAndOffset_Choice_Sf80:
		var tmp_int_Sf80 int64
		if tmp_int_Sf80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Sf80", err)
		}
		ie.Sf80 = tmp_int_Sf80
	case SSB_MTC_periodicityAndOffset_Choice_Sf160:
		var tmp_int_Sf160 int64
		if tmp_int_Sf160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Sf160", err)
		}
		ie.Sf160 = tmp_int_Sf160
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
