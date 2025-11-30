package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PLMN_IdentityList2_r16 struct {
	Value []PLMN_Identity `lb:1,ub:16,madatory`
}

func (ie *PLMN_IdentityList2_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: 16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode PLMN_IdentityList2_r16", err)
	}
	return nil
}

func (ie *PLMN_IdentityList2_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: 16}, false)
	fn := func() *PLMN_Identity {
		return new(PLMN_Identity)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode PLMN_IdentityList2_r16", err)
	}
	ie.Value = []PLMN_Identity{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
