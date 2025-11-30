package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_prach_RootSequenceIndex_Choice_nothing uint64 = iota
	RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L839
	RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L139
)

type RACH_ConfigCommon_prach_RootSequenceIndex struct {
	Choice uint64
	L839   int64 `lb:0,ub:837,madatory`
	L139   int64 `lb:0,ub:137,madatory`
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L839:
		if err = w.WriteInteger(int64(ie.L839), &aper.Constraint{Lb: 0, Ub: 837}, false); err != nil {
			err = utils.WrapError("Encode L839", err)
		}
	case RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L139:
		if err = w.WriteInteger(int64(ie.L139), &aper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			err = utils.WrapError("Encode L139", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L839:
		var tmp_int_L839 int64
		if tmp_int_L839, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 837}, false); err != nil {
			return utils.WrapError("Decode L839", err)
		}
		ie.L839 = tmp_int_L839
	case RACH_ConfigCommon_prach_RootSequenceIndex_Choice_L139:
		var tmp_int_L139 int64
		if tmp_int_L139, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 137}, false); err != nil {
			return utils.WrapError("Decode L139", err)
		}
		ie.L139 = tmp_int_L139
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
