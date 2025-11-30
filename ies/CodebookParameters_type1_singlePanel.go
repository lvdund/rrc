package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters_type1_singlePanel struct {
	SupportedCSI_RS_ResourceList   []SupportedCSI_RS_Resource                 `lb:1,ub:maxNrofCSI_RS_Resources,madatory`
	Modes                          CodebookParameters_type1_singlePanel_modes `madatory`
	MaxNumberCSI_RS_PerResourceSet int64                                      `lb:1,ub:8,madatory`
}

func (ie *CodebookParameters_type1_singlePanel) Encode(w *aper.AperWriter) error {
	var err error
	tmp_SupportedCSI_RS_ResourceList := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, aper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
	for _, i := range ie.SupportedCSI_RS_ResourceList {
		tmp_SupportedCSI_RS_ResourceList.Value = append(tmp_SupportedCSI_RS_ResourceList.Value, &i)
	}
	if err = tmp_SupportedCSI_RS_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedCSI_RS_ResourceList", err)
	}
	if err = ie.Modes.Encode(w); err != nil {
		return utils.WrapError("Encode Modes", err)
	}
	if err = w.WriteInteger(ie.MaxNumberCSI_RS_PerResourceSet, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	return nil
}

func (ie *CodebookParameters_type1_singlePanel) Decode(r *aper.AperReader) error {
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
	if err = ie.Modes.Decode(r); err != nil {
		return utils.WrapError("Decode Modes", err)
	}
	var tmp_int_MaxNumberCSI_RS_PerResourceSet int64
	if tmp_int_MaxNumberCSI_RS_PerResourceSet, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxNumberCSI_RS_PerResourceSet", err)
	}
	ie.MaxNumberCSI_RS_PerResourceSet = tmp_int_MaxNumberCSI_RS_PerResourceSet
	return nil
}
