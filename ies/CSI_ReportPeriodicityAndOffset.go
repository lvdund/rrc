package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportPeriodicityAndOffset_Choice_nothing uint64 = iota
	CSI_ReportPeriodicityAndOffset_Choice_Slots4
	CSI_ReportPeriodicityAndOffset_Choice_Slots5
	CSI_ReportPeriodicityAndOffset_Choice_Slots8
	CSI_ReportPeriodicityAndOffset_Choice_Slots10
	CSI_ReportPeriodicityAndOffset_Choice_Slots16
	CSI_ReportPeriodicityAndOffset_Choice_Slots20
	CSI_ReportPeriodicityAndOffset_Choice_Slots40
	CSI_ReportPeriodicityAndOffset_Choice_Slots80
	CSI_ReportPeriodicityAndOffset_Choice_Slots160
	CSI_ReportPeriodicityAndOffset_Choice_Slots320
)

type CSI_ReportPeriodicityAndOffset struct {
	Choice   uint64
	Slots4   int64 `lb:0,ub:3,madatory`
	Slots5   int64 `lb:0,ub:4,madatory`
	Slots8   int64 `lb:0,ub:7,madatory`
	Slots10  int64 `lb:0,ub:9,madatory`
	Slots16  int64 `lb:0,ub:15,madatory`
	Slots20  int64 `lb:0,ub:19,madatory`
	Slots40  int64 `lb:0,ub:39,madatory`
	Slots80  int64 `lb:0,ub:79,madatory`
	Slots160 int64 `lb:0,ub:159,madatory`
	Slots320 int64 `lb:0,ub:319,madatory`
}

func (ie *CSI_ReportPeriodicityAndOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 10, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportPeriodicityAndOffset_Choice_Slots4:
		if err = w.WriteInteger(int64(ie.Slots4), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Slots4", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots5:
		if err = w.WriteInteger(int64(ie.Slots5), &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Slots5", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots8:
		if err = w.WriteInteger(int64(ie.Slots8), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Slots8", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots10:
		if err = w.WriteInteger(int64(ie.Slots10), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Slots10", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots16:
		if err = w.WriteInteger(int64(ie.Slots16), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Slots16", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots20:
		if err = w.WriteInteger(int64(ie.Slots20), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Slots20", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots40:
		if err = w.WriteInteger(int64(ie.Slots40), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Slots40", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots80:
		if err = w.WriteInteger(int64(ie.Slots80), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Slots80", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots160:
		if err = w.WriteInteger(int64(ie.Slots160), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Slots160", err)
		}
	case CSI_ReportPeriodicityAndOffset_Choice_Slots320:
		if err = w.WriteInteger(int64(ie.Slots320), &uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Slots320", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ReportPeriodicityAndOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(10, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ReportPeriodicityAndOffset_Choice_Slots4:
		var tmp_int_Slots4 int64
		if tmp_int_Slots4, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Slots4", err)
		}
		ie.Slots4 = tmp_int_Slots4
	case CSI_ReportPeriodicityAndOffset_Choice_Slots5:
		var tmp_int_Slots5 int64
		if tmp_int_Slots5, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Slots5", err)
		}
		ie.Slots5 = tmp_int_Slots5
	case CSI_ReportPeriodicityAndOffset_Choice_Slots8:
		var tmp_int_Slots8 int64
		if tmp_int_Slots8, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Slots8", err)
		}
		ie.Slots8 = tmp_int_Slots8
	case CSI_ReportPeriodicityAndOffset_Choice_Slots10:
		var tmp_int_Slots10 int64
		if tmp_int_Slots10, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Slots10", err)
		}
		ie.Slots10 = tmp_int_Slots10
	case CSI_ReportPeriodicityAndOffset_Choice_Slots16:
		var tmp_int_Slots16 int64
		if tmp_int_Slots16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Slots16", err)
		}
		ie.Slots16 = tmp_int_Slots16
	case CSI_ReportPeriodicityAndOffset_Choice_Slots20:
		var tmp_int_Slots20 int64
		if tmp_int_Slots20, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Slots20", err)
		}
		ie.Slots20 = tmp_int_Slots20
	case CSI_ReportPeriodicityAndOffset_Choice_Slots40:
		var tmp_int_Slots40 int64
		if tmp_int_Slots40, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Slots40", err)
		}
		ie.Slots40 = tmp_int_Slots40
	case CSI_ReportPeriodicityAndOffset_Choice_Slots80:
		var tmp_int_Slots80 int64
		if tmp_int_Slots80, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Slots80", err)
		}
		ie.Slots80 = tmp_int_Slots80
	case CSI_ReportPeriodicityAndOffset_Choice_Slots160:
		var tmp_int_Slots160 int64
		if tmp_int_Slots160, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Slots160", err)
		}
		ie.Slots160 = tmp_int_Slots160
	case CSI_ReportPeriodicityAndOffset_Choice_Slots320:
		var tmp_int_Slots320 int64
		if tmp_int_Slots320, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Slots320", err)
		}
		ie.Slots320 = tmp_int_Slots320
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
