package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NPN_Identity_r16_snpn_r16 struct {
	Plmn_Identity_r16 PLMN_Identity `madatory`
	Nid_List_r16      []NID_r16     `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *NPN_Identity_r16_snpn_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Plmn_Identity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r16", err)
	}
	tmp_Nid_List_r16 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.Nid_List_r16 {
		tmp_Nid_List_r16.Value = append(tmp_Nid_List_r16.Value, &i)
	}
	if err = tmp_Nid_List_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Nid_List_r16", err)
	}
	return nil
}

func (ie *NPN_Identity_r16_snpn_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Plmn_Identity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r16", err)
	}
	tmp_Nid_List_r16 := utils.NewSequence[*NID_r16]([]*NID_r16{}, uper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn_Nid_List_r16 := func() *NID_r16 {
		return new(NID_r16)
	}
	if err = tmp_Nid_List_r16.Decode(r, fn_Nid_List_r16); err != nil {
		return utils.WrapError("Decode Nid_List_r16", err)
	}
	ie.Nid_List_r16 = []NID_r16{}
	for _, i := range tmp_Nid_List_r16.Value {
		ie.Nid_List_r16 = append(ie.Nid_List_r16, *i)
	}
	return nil
}
