package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CG_COT_Sharing_r17_Choice_nothing uint64 = iota
	CG_COT_Sharing_r17_Choice_NoCOT_Sharing_r17
	CG_COT_Sharing_r17_Choice_Cot_Sharing_r17
)

type CG_COT_Sharing_r17 struct {
	Choice            uint64
	NoCOT_Sharing_r17 aper.NULL `madatory`
	Cot_Sharing_r17   *CG_COT_Sharing_r17_cot_Sharing_r17
}

func (ie *CG_COT_Sharing_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r17_Choice_NoCOT_Sharing_r17:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode NoCOT_Sharing_r17", err)
		}
	case CG_COT_Sharing_r17_Choice_Cot_Sharing_r17:
		if err = ie.Cot_Sharing_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode Cot_Sharing_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CG_COT_Sharing_r17) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CG_COT_Sharing_r17_Choice_NoCOT_Sharing_r17:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode NoCOT_Sharing_r17", err)
		}
	case CG_COT_Sharing_r17_Choice_Cot_Sharing_r17:
		ie.Cot_Sharing_r17 = new(CG_COT_Sharing_r17_cot_Sharing_r17)
		if err = ie.Cot_Sharing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cot_Sharing_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
