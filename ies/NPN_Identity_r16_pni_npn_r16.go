package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NPN_Identity_r16_pni_npn_r16 struct {
	Plmn_Identity_r16    PLMN_Identity          `madatory`
	Cag_IdentityList_r16 []CAG_IdentityInfo_r16 `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *NPN_Identity_r16_pni_npn_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r16", err)
	}
	tmp_Cag_IdentityList_r16 := utils.NewSequence[*CAG_IdentityInfo_r16]([]*CAG_IdentityInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.Cag_IdentityList_r16 {
		tmp_Cag_IdentityList_r16.Value = append(tmp_Cag_IdentityList_r16.Value, &i)
	}
	if err = tmp_Cag_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Cag_IdentityList_r16", err)
	}
	return nil
}

func (ie *NPN_Identity_r16_pni_npn_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r16", err)
	}
	tmp_Cag_IdentityList_r16 := utils.NewSequence[*CAG_IdentityInfo_r16]([]*CAG_IdentityInfo_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_Cag_IdentityList_r16 := func() *CAG_IdentityInfo_r16 {
		return new(CAG_IdentityInfo_r16)
	}
	if err = tmp_Cag_IdentityList_r16.Decode(r, fn_Cag_IdentityList_r16); err != nil {
		return utils.WrapError("Decode Cag_IdentityList_r16", err)
	}
	ie.Cag_IdentityList_r16 = []CAG_IdentityInfo_r16{}
	for _, i := range tmp_Cag_IdentityList_r16.Value {
		ie.Cag_IdentityList_r16 = append(ie.Cag_IdentityList_r16, *i)
	}
	return nil
}
