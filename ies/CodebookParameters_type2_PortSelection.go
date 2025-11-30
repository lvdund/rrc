package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type2_PortSelection struct {
	SupportedCSI_RS_ResourceList []SupportedCSI_RS_Resource                                  `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	ParameterLx                  int64                                                       `lb:2,ub:4,madatory`
	AmplitudeScalingType         CodebookParameters_type2_PortSelection_amplitudeScalingType `madatory`
}

func (ie *CodebookParameters_type2_PortSelection) Encode(w *aper.AperWriter) error {
	var err error
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
	return nil
}

func (ie *CodebookParameters_type2_PortSelection) Decode(r *aper.AperReader) error {
	var err error
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
	return nil
}
