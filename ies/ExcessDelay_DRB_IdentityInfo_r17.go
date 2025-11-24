package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ExcessDelay_DRB_IdentityInfo_r17 struct {
	Drb_IdentityList []DRB_Identity                                  `lb:1,ub:maxDRB,madatory`
	DelayThreshold   ExcessDelay_DRB_IdentityInfo_r17_delayThreshold `madatory`
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_Drb_IdentityList := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Drb_IdentityList {
		tmp_Drb_IdentityList.Value = append(tmp_Drb_IdentityList.Value, &i)
	}
	if err = tmp_Drb_IdentityList.Encode(w); err != nil {
		return utils.WrapError("Encode Drb_IdentityList", err)
	}
	if err = ie.DelayThreshold.Encode(w); err != nil {
		return utils.WrapError("Encode DelayThreshold", err)
	}
	return nil
}

func (ie *ExcessDelay_DRB_IdentityInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_Drb_IdentityList := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, uper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_Drb_IdentityList := func() *DRB_Identity {
		return new(DRB_Identity)
	}
	if err = tmp_Drb_IdentityList.Decode(r, fn_Drb_IdentityList); err != nil {
		return utils.WrapError("Decode Drb_IdentityList", err)
	}
	ie.Drb_IdentityList = []DRB_Identity{}
	for _, i := range tmp_Drb_IdentityList.Value {
		ie.Drb_IdentityList = append(ie.Drb_IdentityList, *i)
	}
	if err = ie.DelayThreshold.Decode(r); err != nil {
		return utils.WrapError("Decode DelayThreshold", err)
	}
	return nil
}
