package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressConfigurationList_r16 struct {
	Iab_IP_AddressToAddModList_r16  []IAB_IP_AddressConfiguration_r16 `lb:1,ub:maxIAB_IP_Address_r16,optional`
	Iab_IP_AddressToReleaseList_r16 []IAB_IP_AddressIndex_r16         `lb:1,ub:maxIAB_IP_Address_r16,optional`
}

func (ie *IAB_IP_AddressConfigurationList_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Iab_IP_AddressToAddModList_r16) > 0, len(ie.Iab_IP_AddressToReleaseList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.Iab_IP_AddressToAddModList_r16) > 0 {
		tmp_Iab_IP_AddressToAddModList_r16 := utils.NewSequence[*IAB_IP_AddressConfiguration_r16]([]*IAB_IP_AddressConfiguration_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		for _, i := range ie.Iab_IP_AddressToAddModList_r16 {
			tmp_Iab_IP_AddressToAddModList_r16.Value = append(tmp_Iab_IP_AddressToAddModList_r16.Value, &i)
		}
		if err = tmp_Iab_IP_AddressToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_IP_AddressToAddModList_r16", err)
		}
	}
	if len(ie.Iab_IP_AddressToReleaseList_r16) > 0 {
		tmp_Iab_IP_AddressToReleaseList_r16 := utils.NewSequence[*IAB_IP_AddressIndex_r16]([]*IAB_IP_AddressIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		for _, i := range ie.Iab_IP_AddressToReleaseList_r16 {
			tmp_Iab_IP_AddressToReleaseList_r16.Value = append(tmp_Iab_IP_AddressToReleaseList_r16.Value, &i)
		}
		if err = tmp_Iab_IP_AddressToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_IP_AddressToReleaseList_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressConfigurationList_r16) Decode(r *uper.UperReader) error {
	var err error
	var Iab_IP_AddressToAddModList_r16Present bool
	if Iab_IP_AddressToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Iab_IP_AddressToReleaseList_r16Present bool
	if Iab_IP_AddressToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Iab_IP_AddressToAddModList_r16Present {
		tmp_Iab_IP_AddressToAddModList_r16 := utils.NewSequence[*IAB_IP_AddressConfiguration_r16]([]*IAB_IP_AddressConfiguration_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		fn_Iab_IP_AddressToAddModList_r16 := func() *IAB_IP_AddressConfiguration_r16 {
			return new(IAB_IP_AddressConfiguration_r16)
		}
		if err = tmp_Iab_IP_AddressToAddModList_r16.Decode(r, fn_Iab_IP_AddressToAddModList_r16); err != nil {
			return utils.WrapError("Decode Iab_IP_AddressToAddModList_r16", err)
		}
		ie.Iab_IP_AddressToAddModList_r16 = []IAB_IP_AddressConfiguration_r16{}
		for _, i := range tmp_Iab_IP_AddressToAddModList_r16.Value {
			ie.Iab_IP_AddressToAddModList_r16 = append(ie.Iab_IP_AddressToAddModList_r16, *i)
		}
	}
	if Iab_IP_AddressToReleaseList_r16Present {
		tmp_Iab_IP_AddressToReleaseList_r16 := utils.NewSequence[*IAB_IP_AddressIndex_r16]([]*IAB_IP_AddressIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxIAB_IP_Address_r16}, false)
		fn_Iab_IP_AddressToReleaseList_r16 := func() *IAB_IP_AddressIndex_r16 {
			return new(IAB_IP_AddressIndex_r16)
		}
		if err = tmp_Iab_IP_AddressToReleaseList_r16.Decode(r, fn_Iab_IP_AddressToReleaseList_r16); err != nil {
			return utils.WrapError("Decode Iab_IP_AddressToReleaseList_r16", err)
		}
		ie.Iab_IP_AddressToReleaseList_r16 = []IAB_IP_AddressIndex_r16{}
		for _, i := range tmp_Iab_IP_AddressToReleaseList_r16.Value {
			ie.Iab_IP_AddressToReleaseList_r16 = append(ie.Iab_IP_AddressToReleaseList_r16, *i)
		}
	}
	return nil
}
