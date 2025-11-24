package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_TwoStep_r16_resourcesTwoStep_r16 struct {
	Ssb_ResourceList         []CFRA_SSB_Resource `lb:1,ub:maxRA_SSB_Resources,madatory`
	Ra_ssb_OccasionMaskIndex int64               `lb:0,ub:15,madatory`
}

func (ie *CFRA_TwoStep_r16_resourcesTwoStep_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Ssb_ResourceList := utils.NewSequence[*CFRA_SSB_Resource]([]*CFRA_SSB_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_SSB_Resources}, false)
	for _, i := range ie.Ssb_ResourceList {
		tmp_Ssb_ResourceList.Value = append(tmp_Ssb_ResourceList.Value, &i)
	}
	if err = tmp_Ssb_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_ResourceList", err)
	}
	if err = w.WriteInteger(ie.Ra_ssb_OccasionMaskIndex, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger Ra_ssb_OccasionMaskIndex", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_resourcesTwoStep_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp_Ssb_ResourceList := utils.NewSequence[*CFRA_SSB_Resource]([]*CFRA_SSB_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_SSB_Resources}, false)
	fn_Ssb_ResourceList := func() *CFRA_SSB_Resource {
		return new(CFRA_SSB_Resource)
	}
	if err = tmp_Ssb_ResourceList.Decode(r, fn_Ssb_ResourceList); err != nil {
		return utils.WrapError("Decode Ssb_ResourceList", err)
	}
	ie.Ssb_ResourceList = []CFRA_SSB_Resource{}
	for _, i := range tmp_Ssb_ResourceList.Value {
		ie.Ssb_ResourceList = append(ie.Ssb_ResourceList, *i)
	}
	var tmp_int_Ra_ssb_OccasionMaskIndex int64
	if tmp_int_Ra_ssb_OccasionMaskIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger Ra_ssb_OccasionMaskIndex", err)
	}
	ie.Ra_ssb_OccasionMaskIndex = tmp_int_Ra_ssb_OccasionMaskIndex
	return nil
}
