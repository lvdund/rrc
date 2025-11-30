package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_nothing uint64 = iota
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L571
	RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L1151
)

type RACH_ConfigCommon_prach_RootSequenceIndex_r16 struct {
	Choice uint64
	L571   int64 `lb:0,ub:569,madatory`
	L1151  int64 `lb:0,ub:1149,madatory`
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L571:
		if err = w.WriteInteger(int64(ie.L571), &aper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			err = utils.WrapError("Encode L571", err)
		}
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L1151:
		if err = w.WriteInteger(int64(ie.L1151), &aper.Constraint{Lb: 0, Ub: 1149}, false); err != nil {
			err = utils.WrapError("Encode L1151", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommon_prach_RootSequenceIndex_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L571:
		var tmp_int_L571 int64
		if tmp_int_L571, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 569}, false); err != nil {
			return utils.WrapError("Decode L571", err)
		}
		ie.L571 = tmp_int_L571
	case RACH_ConfigCommon_prach_RootSequenceIndex_r16_Choice_L1151:
		var tmp_int_L1151 int64
		if tmp_int_L1151, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1149}, false); err != nil {
			return utils.WrapError("Decode L1151", err)
		}
		ie.L1151 = tmp_int_L1151
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
