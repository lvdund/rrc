package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasTriggerQuantityOffset_Choice_nothing uint64 = iota
	MeasTriggerQuantityOffset_Choice_Rsrp
	MeasTriggerQuantityOffset_Choice_Rsrq
	MeasTriggerQuantityOffset_Choice_Sinr
)

type MeasTriggerQuantityOffset struct {
	Choice uint64
	Rsrp   int64 `lb:-30,ub:30,madatory`
	Rsrq   int64 `lb:-30,ub:30,madatory`
	Sinr   int64 `lb:-30,ub:30,madatory`
}

func (ie *MeasTriggerQuantityOffset) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityOffset_Choice_Rsrp:
		if err = w.WriteInteger(int64(ie.Rsrp), &aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode Rsrp", err)
		}
	case MeasTriggerQuantityOffset_Choice_Rsrq:
		if err = w.WriteInteger(int64(ie.Rsrq), &aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode Rsrq", err)
		}
	case MeasTriggerQuantityOffset_Choice_Sinr:
		if err = w.WriteInteger(int64(ie.Sinr), &aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			err = utils.WrapError("Encode Sinr", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MeasTriggerQuantityOffset) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MeasTriggerQuantityOffset_Choice_Rsrp:
		var tmp_int_Rsrp int64
		if tmp_int_Rsrp, err = r.ReadInteger(&aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode Rsrp", err)
		}
		ie.Rsrp = tmp_int_Rsrp
	case MeasTriggerQuantityOffset_Choice_Rsrq:
		var tmp_int_Rsrq int64
		if tmp_int_Rsrq, err = r.ReadInteger(&aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode Rsrq", err)
		}
		ie.Rsrq = tmp_int_Rsrq
	case MeasTriggerQuantityOffset_Choice_Sinr:
		var tmp_int_Sinr int64
		if tmp_int_Sinr, err = r.ReadInteger(&aper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
			return utils.WrapError("Decode Sinr", err)
		}
		ie.Sinr = tmp_int_Sinr
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
