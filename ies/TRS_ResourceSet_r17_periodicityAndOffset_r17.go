package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_nothing uint64 = iota
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots10
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots20
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots40
	TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots80
)

type TRS_ResourceSet_r17_periodicityAndOffset_r17 struct {
	Choice  uint64
	Slots10 int64 `lb:0,ub:9,madatory`
	Slots20 int64 `lb:0,ub:19,madatory`
	Slots40 int64 `lb:0,ub:39,madatory`
	Slots80 int64 `lb:0,ub:79,madatory`
}

func (ie *TRS_ResourceSet_r17_periodicityAndOffset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots10:
		if err = w.WriteInteger(int64(ie.Slots10), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Slots10", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots20:
		if err = w.WriteInteger(int64(ie.Slots20), &aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode Slots20", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots40:
		if err = w.WriteInteger(int64(ie.Slots40), &aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode Slots40", err)
		}
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots80:
		if err = w.WriteInteger(int64(ie.Slots80), &aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode Slots80", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TRS_ResourceSet_r17_periodicityAndOffset_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots10:
		var tmp_int_Slots10 int64
		if tmp_int_Slots10, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Slots10", err)
		}
		ie.Slots10 = tmp_int_Slots10
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots20:
		var tmp_int_Slots20 int64
		if tmp_int_Slots20, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode Slots20", err)
		}
		ie.Slots20 = tmp_int_Slots20
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots40:
		var tmp_int_Slots40 int64
		if tmp_int_Slots40, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode Slots40", err)
		}
		ie.Slots40 = tmp_int_Slots40
	case TRS_ResourceSet_r17_periodicityAndOffset_r17_Choice_Slots80:
		var tmp_int_Slots80 int64
		if tmp_int_Slots80, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode Slots80", err)
		}
		ie.Slots80 = tmp_int_Slots80
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
