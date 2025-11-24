package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_nothing uint64 = iota
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ImplicitACBarringList
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ExplicitACBarringList
)

type UAC_BarringPerPLMN_uac_ACBarringListType struct {
	Choice                    uint64
	Uac_ImplicitACBarringList []UAC_BarringInfoSetIndex `lb:maxAccessCat_1,ub:maxAccessCat_1,madatory`
	Uac_ExplicitACBarringList *UAC_BarringPerCatList
}

func (ie *UAC_BarringPerPLMN_uac_ACBarringListType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ImplicitACBarringList:
		tmp := utils.NewSequence[*UAC_BarringInfoSetIndex]([]*UAC_BarringInfoSetIndex{}, uper.Constraint{Lb: maxAccessCat_1, Ub: maxAccessCat_1}, false)
		for _, i := range ie.Uac_ImplicitACBarringList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode Uac_ImplicitACBarringList", err)
		}
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ExplicitACBarringList:
		if err = ie.Uac_ExplicitACBarringList.Encode(w); err != nil {
			err = utils.WrapError("Encode Uac_ExplicitACBarringList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UAC_BarringPerPLMN_uac_ACBarringListType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ImplicitACBarringList:
		tmp := utils.NewSequence[*UAC_BarringInfoSetIndex]([]*UAC_BarringInfoSetIndex{}, uper.Constraint{Lb: maxAccessCat_1, Ub: maxAccessCat_1}, false)
		fn := func() *UAC_BarringInfoSetIndex {
			return new(UAC_BarringInfoSetIndex)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode Uac_ImplicitACBarringList", err)
		}
		ie.Uac_ImplicitACBarringList = []UAC_BarringInfoSetIndex{}
		for _, i := range tmp.Value {
			ie.Uac_ImplicitACBarringList = append(ie.Uac_ImplicitACBarringList, *i)
		}
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_Uac_ExplicitACBarringList:
		ie.Uac_ExplicitACBarringList = new(UAC_BarringPerCatList)
		if err = ie.Uac_ExplicitACBarringList.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_ExplicitACBarringList", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
