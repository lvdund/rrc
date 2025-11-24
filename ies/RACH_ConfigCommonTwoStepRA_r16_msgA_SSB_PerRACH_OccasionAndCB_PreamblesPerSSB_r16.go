package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_nothing uint64 = iota
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneEighth
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneFourth
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneHalf
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_One
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Two
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Four
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Eight
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Sixteen
)

type RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 struct {
	Choice    uint64
	OneEighth *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneEighth
	OneFourth *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneFourth
	OneHalf   *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneHalf
	One       *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_one
	Two       *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two
	Four      int64 `lb:1,ub:16,madatory`
	Eight     int64 `lb:1,ub:8,madatory`
	Sixteen   int64 `lb:1,ub:4,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneEighth:
		if err = ie.OneEighth.Encode(w); err != nil {
			err = utils.WrapError("Encode OneEighth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneFourth:
		if err = ie.OneFourth.Encode(w); err != nil {
			err = utils.WrapError("Encode OneFourth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneHalf:
		if err = ie.OneHalf.Encode(w); err != nil {
			err = utils.WrapError("Encode OneHalf", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_One:
		if err = ie.One.Encode(w); err != nil {
			err = utils.WrapError("Encode One", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Two:
		if err = ie.Two.Encode(w); err != nil {
			err = utils.WrapError("Encode Two", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Four:
		if err = w.WriteInteger(int64(ie.Four), &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Four", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Eight:
		if err = w.WriteInteger(int64(ie.Eight), &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Eight", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Sixteen:
		if err = w.WriteInteger(int64(ie.Sixteen), &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sixteen", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneEighth:
		ie.OneEighth = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneEighth)
		if err = ie.OneEighth.Decode(r); err != nil {
			return utils.WrapError("Decode OneEighth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneFourth:
		ie.OneFourth = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneFourth)
		if err = ie.OneFourth.Decode(r); err != nil {
			return utils.WrapError("Decode OneFourth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_OneHalf:
		ie.OneHalf = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneHalf)
		if err = ie.OneHalf.Decode(r); err != nil {
			return utils.WrapError("Decode OneHalf", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_One:
		ie.One = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_one)
		if err = ie.One.Decode(r); err != nil {
			return utils.WrapError("Decode One", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Two:
		ie.Two = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two)
		if err = ie.Two.Decode(r); err != nil {
			return utils.WrapError("Decode Two", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Four:
		var tmp_int_Four int64
		if tmp_int_Four, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Four", err)
		}
		ie.Four = tmp_int_Four
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Eight:
		var tmp_int_Eight int64
		if tmp_int_Eight, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Eight", err)
		}
		ie.Eight = tmp_int_Eight
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_Sixteen:
		var tmp_int_Sixteen int64
		if tmp_int_Sixteen, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sixteen", err)
		}
		ie.Sixteen = tmp_int_Sixteen
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
