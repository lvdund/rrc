package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type2 struct {
	SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource                           `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	ParameterLx                  int64                                                `lb:2,ub:4,madatory`
	AmplitudeScalingType         CodebookParameters_type2_amplitudeScalingType        `madatory`
	AmplitudeSubsetRestriction   *CodebookParameters_type2_amplitudeSubsetRestriction `optional`
}

func (ie *CodebookParameters_type2) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AmplitudeSubsetRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_SupportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	for _, i := range ie.SupportedCSI_RS_ResourceList {
		tmp_SupportedCSI_RS_ResourceList.Value = append(tmp_SupportedCSI_RS_ResourceList.Value, &i)
	}
	if err = tmp_SupportedCSI_RS_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedCSI_RS_ResourceList", err)
	}
	if err = w.WriteInteger(ie.ParameterLx, &aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger ParameterLx", err)
	}
	if err = ie.AmplitudeScalingType.Encode(w); err != nil {
		return utils.WrapError("Encode AmplitudeScalingType", err)
	}
	if ie.AmplitudeSubsetRestriction != nil {
		if err = ie.AmplitudeSubsetRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode AmplitudeSubsetRestriction", err)
		}
	}
	return nil
}

func (ie *CodebookParameters_type2) Decode(r *aper.AperReader) error {
	var err error
	var AmplitudeSubsetRestrictionPresent bool
	if AmplitudeSubsetRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_SupportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	fn_SupportedCSI_RS_ResourceList := func() *SupportedCSI_RS_Resource {
		return new(SupportedCSI_RS_Resource)
	}
	if err = tmp_SupportedCSI_RS_ResourceList.Decode(r, fn_SupportedCSI_RS_ResourceList); err != nil {
		return utils.WrapError("Decode SupportedCSI_RS_ResourceList", err)
	}
	ie.SupportedCSI_RS_ResourceList = []SupportedCSI_RS_Resource{}
	for _, i := range tmp_SupportedCSI_RS_ResourceList.Value {
		ie.SupportedCSI_RS_ResourceList = append(ie.SupportedCSI_RS_ResourceList, *i)
	}
	var tmp_int_ParameterLx int64
	if tmp_int_ParameterLx, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger ParameterLx", err)
	}
	ie.ParameterLx = tmp_int_ParameterLx
	if err = ie.AmplitudeScalingType.Decode(r); err != nil {
		return utils.WrapError("Decode AmplitudeScalingType", err)
	}
	if AmplitudeSubsetRestrictionPresent {
		ie.AmplitudeSubsetRestriction = new(CodebookParameters_type2_amplitudeSubsetRestriction)
		if err = ie.AmplitudeSubsetRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode AmplitudeSubsetRestriction", err)
		}
	}
	return nil
}
