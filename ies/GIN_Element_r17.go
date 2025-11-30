package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type GIN_Element_r17 struct {
	Plmn_Identity_r17 PLMN_Identity `madatory`
	Nid_List_r17      []NID_r16     `lb:1,ub:maxGIN_r17,madatory`
}

func (ie *GIN_Element_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Plmn_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_Identity_r17", err)
	}
	tmp_Nid_List_r17 := utils.NewSequence[*NID_r16]([]*NID_r16{}, aper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
	for _, i := range ie.Nid_List_r17 {
		tmp_Nid_List_r17.Value = append(tmp_Nid_List_r17.Value, &i)
	}
	if err = tmp_Nid_List_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nid_List_r17", err)
	}
	return nil
}

func (ie *GIN_Element_r17) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Plmn_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_Identity_r17", err)
	}
	tmp_Nid_List_r17 := utils.NewSequence[*NID_r16]([]*NID_r16{}, aper.Constraint{Lb: 1, Ub: maxGIN_r17}, false)
	fn_Nid_List_r17 := func() *NID_r16 {
		return new(NID_r16)
	}
	if err = tmp_Nid_List_r17.Decode(r, fn_Nid_List_r17); err != nil {
		return utils.WrapError("Decode Nid_List_r17", err)
	}
	ie.Nid_List_r17 = []NID_r16{}
	for _, i := range tmp_Nid_List_r17.Value {
		ie.Nid_List_r17 = append(ie.Nid_List_r17, *i)
	}
	return nil
}
