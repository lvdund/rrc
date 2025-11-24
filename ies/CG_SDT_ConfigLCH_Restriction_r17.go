package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_ConfigLCH_Restriction_r17 struct {
	LogicalChannelIdentity_r17      LogicalChannelIdentity                                            `madatory`
	ConfiguredGrantType1Allowed_r17 *CG_SDT_ConfigLCH_Restriction_r17_configuredGrantType1Allowed_r17 `optional`
	AllowedCG_List_r17              []ConfiguredGrantConfigIndexMAC_r16                               `lb:0,ub:maxNrofConfiguredGrantConfigMAC_1_r16,optional`
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ConfiguredGrantType1Allowed_r17 != nil, len(ie.AllowedCG_List_r17) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.LogicalChannelIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode LogicalChannelIdentity_r17", err)
	}
	if ie.ConfiguredGrantType1Allowed_r17 != nil {
		if err = ie.ConfiguredGrantType1Allowed_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ConfiguredGrantType1Allowed_r17", err)
		}
	}
	if len(ie.AllowedCG_List_r17) > 0 {
		tmp_AllowedCG_List_r17 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		for _, i := range ie.AllowedCG_List_r17 {
			tmp_AllowedCG_List_r17.Value = append(tmp_AllowedCG_List_r17.Value, &i)
		}
		if err = tmp_AllowedCG_List_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AllowedCG_List_r17", err)
		}
	}
	return nil
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Decode(r *uper.UperReader) error {
	var err error
	var ConfiguredGrantType1Allowed_r17Present bool
	if ConfiguredGrantType1Allowed_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AllowedCG_List_r17Present bool
	if AllowedCG_List_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.LogicalChannelIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode LogicalChannelIdentity_r17", err)
	}
	if ConfiguredGrantType1Allowed_r17Present {
		ie.ConfiguredGrantType1Allowed_r17 = new(CG_SDT_ConfigLCH_Restriction_r17_configuredGrantType1Allowed_r17)
		if err = ie.ConfiguredGrantType1Allowed_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ConfiguredGrantType1Allowed_r17", err)
		}
	}
	if AllowedCG_List_r17Present {
		tmp_AllowedCG_List_r17 := utils.NewSequence[*ConfiguredGrantConfigIndexMAC_r16]([]*ConfiguredGrantConfigIndexMAC_r16{}, uper.Constraint{Lb: 0, Ub: maxNrofConfiguredGrantConfigMAC_1_r16}, false)
		fn_AllowedCG_List_r17 := func() *ConfiguredGrantConfigIndexMAC_r16 {
			return new(ConfiguredGrantConfigIndexMAC_r16)
		}
		if err = tmp_AllowedCG_List_r17.Decode(r, fn_AllowedCG_List_r17); err != nil {
			return utils.WrapError("Decode AllowedCG_List_r17", err)
		}
		ie.AllowedCG_List_r17 = []ConfiguredGrantConfigIndexMAC_r16{}
		for _, i := range tmp_AllowedCG_List_r17.Value {
			ie.AllowedCG_List_r17 = append(ie.AllowedCG_List_r17, *i)
		}
	}
	return nil
}
