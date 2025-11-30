package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_IdentityInfoList struct {
	Value []PLMN_IdentityInfo `lb:1,ub:maxPLMN,madatory`
}

func (ie *PLMN_IdentityInfoList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PLMN_IdentityInfo]([]*PLMN_IdentityInfo{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PLMN_IdentityInfoList", err)
	}
	return nil
}

func (ie *PLMN_IdentityInfoList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PLMN_IdentityInfo]([]*PLMN_IdentityInfo{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
	fn := func() *PLMN_IdentityInfo {
		return new(PLMN_IdentityInfo)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PLMN_IdentityInfoList", err)
	}
	ie.Value = []PLMN_IdentityInfo{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
