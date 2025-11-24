package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_Config_nAndPagingFrameOffset_Choice_nothing uint64 = iota
	PCCH_Config_nAndPagingFrameOffset_Choice_OneT
	PCCH_Config_nAndPagingFrameOffset_Choice_HalfT
	PCCH_Config_nAndPagingFrameOffset_Choice_QuarterT
	PCCH_Config_nAndPagingFrameOffset_Choice_OneEighthT
	PCCH_Config_nAndPagingFrameOffset_Choice_OneSixteenthT
)

type PCCH_Config_nAndPagingFrameOffset struct {
	Choice        uint64
	OneT          uper.NULL `madatory`
	HalfT         int64     `lb:0,ub:1,madatory`
	QuarterT      int64     `lb:0,ub:3,madatory`
	OneEighthT    int64     `lb:0,ub:7,madatory`
	OneSixteenthT int64     `lb:0,ub:15,madatory`
}

func (ie *PCCH_Config_nAndPagingFrameOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneT:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode OneT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_HalfT:
		if err = w.WriteInteger(int64(ie.HalfT), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode HalfT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_QuarterT:
		if err = w.WriteInteger(int64(ie.QuarterT), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode QuarterT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneEighthT:
		if err = w.WriteInteger(int64(ie.OneEighthT), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode OneEighthT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneSixteenthT:
		if err = w.WriteInteger(int64(ie.OneSixteenthT), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode OneSixteenthT", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PCCH_Config_nAndPagingFrameOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneT:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode OneT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_HalfT:
		var tmp_int_HalfT int64
		if tmp_int_HalfT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode HalfT", err)
		}
		ie.HalfT = tmp_int_HalfT
	case PCCH_Config_nAndPagingFrameOffset_Choice_QuarterT:
		var tmp_int_QuarterT int64
		if tmp_int_QuarterT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode QuarterT", err)
		}
		ie.QuarterT = tmp_int_QuarterT
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneEighthT:
		var tmp_int_OneEighthT int64
		if tmp_int_OneEighthT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode OneEighthT", err)
		}
		ie.OneEighthT = tmp_int_OneEighthT
	case PCCH_Config_nAndPagingFrameOffset_Choice_OneSixteenthT:
		var tmp_int_OneSixteenthT int64
		if tmp_int_OneSixteenthT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode OneSixteenthT", err)
		}
		ie.OneSixteenthT = tmp_int_OneSixteenthT
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
