package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ResourcePeriodicityAndOffset_Choice_nothing uint64 = iota
	CSI_ResourcePeriodicityAndOffset_Choice_Slots4
	CSI_ResourcePeriodicityAndOffset_Choice_Slots5
	CSI_ResourcePeriodicityAndOffset_Choice_Slots8
	CSI_ResourcePeriodicityAndOffset_Choice_Slots10
	CSI_ResourcePeriodicityAndOffset_Choice_Slots16
	CSI_ResourcePeriodicityAndOffset_Choice_Slots20
	CSI_ResourcePeriodicityAndOffset_Choice_Slots32
	CSI_ResourcePeriodicityAndOffset_Choice_Slots40
	CSI_ResourcePeriodicityAndOffset_Choice_Slots64
	CSI_ResourcePeriodicityAndOffset_Choice_Slots80
	CSI_ResourcePeriodicityAndOffset_Choice_Slots160
	CSI_ResourcePeriodicityAndOffset_Choice_Slots320
	CSI_ResourcePeriodicityAndOffset_Choice_Slots640
)

type CSI_ResourcePeriodicityAndOffset struct {
	Choice   uint64
	Slots4   int64 `lb:0,ub:3,madatory`
	Slots5   int64 `lb:0,ub:4,madatory`
	Slots8   int64 `lb:0,ub:7,madatory`
	Slots10  int64 `lb:0,ub:9,madatory`
	Slots16  int64 `lb:0,ub:15,madatory`
	Slots20  int64 `lb:0,ub:19,madatory`
	Slots32  int64 `lb:0,ub:31,madatory`
	Slots40  int64 `lb:0,ub:39,madatory`
	Slots64  int64 `lb:0,ub:63,madatory`
	Slots80  int64 `lb:0,ub:79,madatory`
	Slots160 int64 `lb:0,ub:159,madatory`
	Slots320 int64 `lb:0,ub:319,madatory`
	Slots640 int64 `lb:0,ub:639,madatory`
}

func (ie *CSI_ResourcePeriodicityAndOffset) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots4:
		if err = w.WriteInteger(int64(ie.Slots4), &aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode Slots4", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots5:
		if err = w.WriteInteger(int64(ie.Slots5), &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Slots5", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots8:
		if err = w.WriteInteger(int64(ie.Slots8), &aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode Slots8", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots10:
		if err = w.WriteInteger(int64(ie.Slots10), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Slots10", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots16:
		if err = w.WriteInteger(int64(ie.Slots16), &aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode Slots16", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots20:
		if err = w.WriteInteger(int64(ie.Slots20), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Slots20", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots32:
		if err = w.WriteInteger(int64(ie.Slots32), &aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode Slots32", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots40:
		if err = w.WriteInteger(int64(ie.Slots40), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Slots40", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots64:
		if err = w.WriteInteger(int64(ie.Slots64), &aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode Slots64", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots80:
		if err = w.WriteInteger(int64(ie.Slots80), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Slots80", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots160:
		if err = w.WriteInteger(int64(ie.Slots160), &aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode Slots160", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots320:
		if err = w.WriteInteger(int64(ie.Slots320), &aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			err = utils.WrapError("Encode Slots320", err)
		}
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots640:
		if err = w.WriteInteger(int64(ie.Slots640), &aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			err = utils.WrapError("Encode Slots640", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_ResourcePeriodicityAndOffset) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(13, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots4:
		var tmp_int_Slots4 int64
		if tmp_int_Slots4, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode Slots4", err)
		}
		ie.Slots4 = tmp_int_Slots4
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots5:
		var tmp_int_Slots5 int64
		if tmp_int_Slots5, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Slots5", err)
		}
		ie.Slots5 = tmp_int_Slots5
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots8:
		var tmp_int_Slots8 int64
		if tmp_int_Slots8, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode Slots8", err)
		}
		ie.Slots8 = tmp_int_Slots8
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots10:
		var tmp_int_Slots10 int64
		if tmp_int_Slots10, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Slots10", err)
		}
		ie.Slots10 = tmp_int_Slots10
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots16:
		var tmp_int_Slots16 int64
		if tmp_int_Slots16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Slots16", err)
		}
		ie.Slots16 = tmp_int_Slots16
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots20:
		var tmp_int_Slots20 int64
		if tmp_int_Slots20, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Slots20", err)
		}
		ie.Slots20 = tmp_int_Slots20
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots32:
		var tmp_int_Slots32 int64
		if tmp_int_Slots32, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode Slots32", err)
		}
		ie.Slots32 = tmp_int_Slots32
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots40:
		var tmp_int_Slots40 int64
		if tmp_int_Slots40, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Slots40", err)
		}
		ie.Slots40 = tmp_int_Slots40
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots64:
		var tmp_int_Slots64 int64
		if tmp_int_Slots64, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode Slots64", err)
		}
		ie.Slots64 = tmp_int_Slots64
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots80:
		var tmp_int_Slots80 int64
		if tmp_int_Slots80, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Slots80", err)
		}
		ie.Slots80 = tmp_int_Slots80
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots160:
		var tmp_int_Slots160 int64
		if tmp_int_Slots160, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode Slots160", err)
		}
		ie.Slots160 = tmp_int_Slots160
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots320:
		var tmp_int_Slots320 int64
		if tmp_int_Slots320, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 319}, false); err != nil {
			return utils.WrapError("Decode Slots320", err)
		}
		ie.Slots320 = tmp_int_Slots320
	case CSI_ResourcePeriodicityAndOffset_Choice_Slots640:
		var tmp_int_Slots640 int64
		if tmp_int_Slots640, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode Slots640", err)
		}
		ie.Slots640 = tmp_int_Slots640
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
