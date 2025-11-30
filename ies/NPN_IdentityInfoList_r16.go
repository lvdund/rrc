package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NPN_IdentityInfoList_r16 struct {
	Value []NPN_IdentityInfo_r16 `lb:1,ub:maxNPN_r16,madatory`
}

func (ie *NPN_IdentityInfoList_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*NPN_IdentityInfo_r16]([]*NPN_IdentityInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NPN_IdentityInfoList_r16", err)
	}
	return nil
}

func (ie *NPN_IdentityInfoList_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*NPN_IdentityInfo_r16]([]*NPN_IdentityInfo_r16{}, aper.Constraint{Lb: 1, Ub: maxNPN_r16}, false)
	fn := func() *NPN_IdentityInfo_r16 {
		return new(NPN_IdentityInfo_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NPN_IdentityInfoList_r16", err)
	}
	ie.Value = []NPN_IdentityInfo_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
