package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_resources_csirs struct {
	Csirs_ResourceList   []CFRA_CSIRS_Resource `lb:1,ub:maxRA_CSIRS_Resources,madatory`
	Rsrp_ThresholdCSI_RS RSRP_Range            `madatory`
}

func (ie *CFRA_resources_csirs) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Csirs_ResourceList := utils.NewSequence[*CFRA_CSIRS_Resource]([]*CFRA_CSIRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_CSIRS_Resources}, false)
	for _, i := range ie.Csirs_ResourceList {
		tmp_Csirs_ResourceList.Value = append(tmp_Csirs_ResourceList.Value, &i)
	}
	if err = tmp_Csirs_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode Csirs_ResourceList", err)
	}
	if err = ie.Rsrp_ThresholdCSI_RS.Encode(w); err != nil {
		return utils.WrapError("Encode Rsrp_ThresholdCSI_RS", err)
	}
	return nil
}

func (ie *CFRA_resources_csirs) Decode(r *uper.UperReader) error {
	var err error
	tmp_Csirs_ResourceList := utils.NewSequence[*CFRA_CSIRS_Resource]([]*CFRA_CSIRS_Resource{}, uper.Constraint{Lb: 1, Ub: maxRA_CSIRS_Resources}, false)
	fn_Csirs_ResourceList := func() *CFRA_CSIRS_Resource {
		return new(CFRA_CSIRS_Resource)
	}
	if err = tmp_Csirs_ResourceList.Decode(r, fn_Csirs_ResourceList); err != nil {
		return utils.WrapError("Decode Csirs_ResourceList", err)
	}
	ie.Csirs_ResourceList = []CFRA_CSIRS_Resource{}
	for _, i := range tmp_Csirs_ResourceList.Value {
		ie.Csirs_ResourceList = append(ie.Csirs_ResourceList, *i)
	}
	if err = ie.Rsrp_ThresholdCSI_RS.Decode(r); err != nil {
		return utils.WrapError("Decode Rsrp_ThresholdCSI_RS", err)
	}
	return nil
}
