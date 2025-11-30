package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ServingCellConfig_ca_SlotOffset_r16_Choice_nothing uint64 = iota
	ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS15kHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS30KHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS60KHz
	ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS120KHz
)

type ServingCellConfig_ca_SlotOffset_r16 struct {
	Choice       uint64
	RefSCS15kHz  int64 `lb:-2,ub:2,madatory`
	RefSCS30KHz  int64 `lb:-5,ub:5,madatory`
	RefSCS60KHz  int64 `lb:-10,ub:10,madatory`
	RefSCS120KHz int64 `lb:-20,ub:20,madatory`
}

func (ie *ServingCellConfig_ca_SlotOffset_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS15kHz:
		if err = w.WriteInteger(int64(ie.RefSCS15kHz), &aper.Constraint{Lb: -2, Ub: 2}, false); err != nil {
			err = utils.WrapError("Encode RefSCS15kHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS30KHz:
		if err = w.WriteInteger(int64(ie.RefSCS30KHz), &aper.Constraint{Lb: -5, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode RefSCS30KHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS60KHz:
		if err = w.WriteInteger(int64(ie.RefSCS60KHz), &aper.Constraint{Lb: -10, Ub: 10}, false); err != nil {
			err = utils.WrapError("Encode RefSCS60KHz", err)
		}
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS120KHz:
		if err = w.WriteInteger(int64(ie.RefSCS120KHz), &aper.Constraint{Lb: -20, Ub: 20}, false); err != nil {
			err = utils.WrapError("Encode RefSCS120KHz", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ServingCellConfig_ca_SlotOffset_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS15kHz:
		var tmp_int_RefSCS15kHz int64
		if tmp_int_RefSCS15kHz, err = r.ReadInteger(&aper.Constraint{Lb: -2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode RefSCS15kHz", err)
		}
		ie.RefSCS15kHz = tmp_int_RefSCS15kHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS30KHz:
		var tmp_int_RefSCS30KHz int64
		if tmp_int_RefSCS30KHz, err = r.ReadInteger(&aper.Constraint{Lb: -5, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode RefSCS30KHz", err)
		}
		ie.RefSCS30KHz = tmp_int_RefSCS30KHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS60KHz:
		var tmp_int_RefSCS60KHz int64
		if tmp_int_RefSCS60KHz, err = r.ReadInteger(&aper.Constraint{Lb: -10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode RefSCS60KHz", err)
		}
		ie.RefSCS60KHz = tmp_int_RefSCS60KHz
	case ServingCellConfig_ca_SlotOffset_r16_Choice_RefSCS120KHz:
		var tmp_int_RefSCS120KHz int64
		if tmp_int_RefSCS120KHz, err = r.ReadInteger(&aper.Constraint{Lb: -20, Ub: 20}, false); err != nil {
			return utils.WrapError("Decode RefSCS120KHz", err)
		}
		ie.RefSCS120KHz = tmp_int_RefSCS120KHz
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
