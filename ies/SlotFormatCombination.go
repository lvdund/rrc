package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatCombination struct {
	SlotFormatCombinationId SlotFormatCombinationId `madatory`
	SlotFormats             []int64                 `lb:1,ub:maxNrofSlotFormatsPerCombination,e_lb:0,e_ub:255,madatory`
}

func (ie *SlotFormatCombination) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.SlotFormatCombinationId.Encode(w); err != nil {
		return utils.WrapError("Encode SlotFormatCombinationId", err)
	}
	tmp_SlotFormats := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatsPerCombination}, false)
	for _, i := range ie.SlotFormats {
		tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 255}, false)
		tmp_SlotFormats.Value = append(tmp_SlotFormats.Value, &tmp_ie)
	}
	if err = tmp_SlotFormats.Encode(w); err != nil {
		return utils.WrapError("Encode SlotFormats", err)
	}
	return nil
}

func (ie *SlotFormatCombination) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.SlotFormatCombinationId.Decode(r); err != nil {
		return utils.WrapError("Decode SlotFormatCombinationId", err)
	}
	tmp_SlotFormats := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatsPerCombination}, false)
	fn_SlotFormats := func() *utils.INTEGER {
		ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 255}, false)
		return &ie
	}
	if err = tmp_SlotFormats.Decode(r, fn_SlotFormats); err != nil {
		return utils.WrapError("Decode SlotFormats", err)
	}
	ie.SlotFormats = []int64{}
	for _, i := range tmp_SlotFormats.Value {
		ie.SlotFormats = append(ie.SlotFormats, int64(i.Value))
	}
	return nil
}
