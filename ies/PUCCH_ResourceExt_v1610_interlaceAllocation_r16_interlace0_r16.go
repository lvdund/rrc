package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_nothing uint64 = iota
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs15
	PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs30
)

type PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16 struct {
	Choice uint64
	Scs15  int64 `lb:0,ub:9,madatory`
	Scs30  int64 `lb:0,ub:4,madatory`
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs15:
		if err = w.WriteInteger(int64(ie.Scs15), &aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode Scs15", err)
		}
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs30:
		if err = w.WriteInteger(int64(ie.Scs30), &aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode Scs30", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs15:
		var tmp_int_Scs15 int64
		if tmp_int_Scs15, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode Scs15", err)
		}
		ie.Scs15 = tmp_int_Scs15
	case PUCCH_ResourceExt_v1610_interlaceAllocation_r16_interlace0_r16_Choice_Scs30:
		var tmp_int_Scs30 int64
		if tmp_int_Scs30, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode Scs30", err)
		}
		ie.Scs30 = tmp_int_Scs30
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
