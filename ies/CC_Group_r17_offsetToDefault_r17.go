package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CC_Group_r17_offsetToDefault_r17_Choice_nothing uint64 = iota
	CC_Group_r17_offsetToDefault_r17_Choice_OffsetValue
	CC_Group_r17_offsetToDefault_r17_Choice_Offsetlist
)

type CC_Group_r17_offsetToDefault_r17 struct {
	Choice      uint64
	OffsetValue *OffsetValue_r17
	Offsetlist  []OffsetValue_r17 `lb:1,ub:maxNrofReqComDC_Location_r17,madatory`
}

func (ie *CC_Group_r17_offsetToDefault_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CC_Group_r17_offsetToDefault_r17_Choice_OffsetValue:
		if err = ie.OffsetValue.Encode(w); err != nil {
			err = utils.WrapError("Encode OffsetValue", err)
		}
	case CC_Group_r17_offsetToDefault_r17_Choice_Offsetlist:
		tmp := utils.NewSequence[*OffsetValue_r17]([]*OffsetValue_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
		for _, i := range ie.Offsetlist {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Offsetlist", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CC_Group_r17_offsetToDefault_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CC_Group_r17_offsetToDefault_r17_Choice_OffsetValue:
		ie.OffsetValue = new(OffsetValue_r17)
		if err = ie.OffsetValue.Decode(r); err != nil {
			return utils.WrapError("Decode OffsetValue", err)
		}
	case CC_Group_r17_offsetToDefault_r17_Choice_Offsetlist:
		tmp := utils.NewSequence[*OffsetValue_r17]([]*OffsetValue_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofReqComDC_Location_r17}, false)
		fn := func() *OffsetValue_r17 {
			return new(OffsetValue_r17)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode Offsetlist", err)
		}
		ie.Offsetlist = []OffsetValue_r17{}
		for _, i := range tmp.Value {
			ie.Offsetlist = append(ie.Offsetlist, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
