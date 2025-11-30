package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_nothing uint64 = iota
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneEighth
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneFourth
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneHalf
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_One
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Two
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Four
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Eight
	RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Sixteen
)

type RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB struct {
	Choice    uint64
	OneEighth *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth
	OneFourth *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneFourth
	OneHalf   *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneHalf
	One       *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one
	Two       *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two
	Four      int64 `lb:1,ub:16,madatory`
	Eight     int64 `lb:1,ub:8,madatory`
	Sixteen   int64 `lb:1,ub:4,madatory`
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneEighth:
		if err = ie.OneEighth.Encode(w); err != nil {
			err = utils.WrapError("Encode OneEighth", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneFourth:
		if err = ie.OneFourth.Encode(w); err != nil {
			err = utils.WrapError("Encode OneFourth", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneHalf:
		if err = ie.OneHalf.Encode(w); err != nil {
			err = utils.WrapError("Encode OneHalf", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_One:
		if err = ie.One.Encode(w); err != nil {
			err = utils.WrapError("Encode One", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Two:
		if err = ie.Two.Encode(w); err != nil {
			err = utils.WrapError("Encode Two", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Four:
		if err = w.WriteInteger(int64(ie.Four), &aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode Four", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Eight:
		if err = w.WriteInteger(int64(ie.Eight), &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode Eight", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Sixteen:
		if err = w.WriteInteger(int64(ie.Sixteen), &aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Sixteen", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneEighth:
		ie.OneEighth = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneEighth)
		if err = ie.OneEighth.Decode(r); err != nil {
			return utils.WrapError("Decode OneEighth", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneFourth:
		ie.OneFourth = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneFourth)
		if err = ie.OneFourth.Decode(r); err != nil {
			return utils.WrapError("Decode OneFourth", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_OneHalf:
		ie.OneHalf = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_oneHalf)
		if err = ie.OneHalf.Decode(r); err != nil {
			return utils.WrapError("Decode OneHalf", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_One:
		ie.One = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_one)
		if err = ie.One.Decode(r); err != nil {
			return utils.WrapError("Decode One", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Two:
		ie.Two = new(RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_two)
		if err = ie.Two.Decode(r); err != nil {
			return utils.WrapError("Decode Two", err)
		}
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Four:
		var tmp_int_Four int64
		if tmp_int_Four, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Four", err)
		}
		ie.Four = tmp_int_Four
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Eight:
		var tmp_int_Eight int64
		if tmp_int_Eight, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Eight", err)
		}
		ie.Eight = tmp_int_Eight
	case RACH_ConfigCommon_ssb_perRACH_OccasionAndCB_PreamblesPerSSB_Choice_Sixteen:
		var tmp_int_Sixteen int64
		if tmp_int_Sixteen, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Sixteen", err)
		}
		ie.Sixteen = tmp_int_Sixteen
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
