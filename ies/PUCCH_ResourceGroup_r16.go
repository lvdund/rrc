package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceGroup_r16 struct {
	Pucch_ResourceGroupId_r16 PUCCH_ResourceGroupId_r16 `madatory`
	ResourcePerGroupList_r16  []PUCCH_ResourceId        `lb:1,ub:maxNrofPUCCH_ResourcesPerGroup_r16,madatory`
}

func (ie *PUCCH_ResourceGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Pucch_ResourceGroupId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_ResourceGroupId_r16", err)
	}
	tmp_ResourcePerGroupList_r16 := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerGroup_r16}, false)
	for _, i := range ie.ResourcePerGroupList_r16 {
		tmp_ResourcePerGroupList_r16.Value = append(tmp_ResourcePerGroupList_r16.Value, &i)
	}
	if err = tmp_ResourcePerGroupList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ResourcePerGroupList_r16", err)
	}
	return nil
}

func (ie *PUCCH_ResourceGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Pucch_ResourceGroupId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_ResourceGroupId_r16", err)
	}
	tmp_ResourcePerGroupList_r16 := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerGroup_r16}, false)
	fn_ResourcePerGroupList_r16 := func() *PUCCH_ResourceId {
		return new(PUCCH_ResourceId)
	}
	if err = tmp_ResourcePerGroupList_r16.Decode(r, fn_ResourcePerGroupList_r16); err != nil {
		return utils.WrapError("Decode ResourcePerGroupList_r16", err)
	}
	ie.ResourcePerGroupList_r16 = []PUCCH_ResourceId{}
	for _, i := range tmp_ResourcePerGroupList_r16.Value {
		ie.ResourcePerGroupList_r16 = append(ie.ResourcePerGroupList_r16, *i)
	}
	return nil
}
