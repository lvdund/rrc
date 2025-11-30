package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_ExcessDelayConfig_r17 struct {
	ExcessDelay_DRBlist_r17 []ExcessDelay_DRB_IdentityInfo_r17 `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_ExcessDelayConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp_ExcessDelay_DRBlist_r17 := utils.NewSequence[*ExcessDelay_DRB_IdentityInfo_r17]([]*ExcessDelay_DRB_IdentityInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.ExcessDelay_DRBlist_r17 {
		tmp_ExcessDelay_DRBlist_r17.Value = append(tmp_ExcessDelay_DRBlist_r17.Value, &i)
	}
	if err = tmp_ExcessDelay_DRBlist_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ExcessDelay_DRBlist_r17", err)
	}
	return nil
}

func (ie *UL_ExcessDelayConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp_ExcessDelay_DRBlist_r17 := utils.NewSequence[*ExcessDelay_DRB_IdentityInfo_r17]([]*ExcessDelay_DRB_IdentityInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_ExcessDelay_DRBlist_r17 := func() *ExcessDelay_DRB_IdentityInfo_r17 {
		return new(ExcessDelay_DRB_IdentityInfo_r17)
	}
	if err = tmp_ExcessDelay_DRBlist_r17.Decode(r, fn_ExcessDelay_DRBlist_r17); err != nil {
		return utils.WrapError("Decode ExcessDelay_DRBlist_r17", err)
	}
	ie.ExcessDelay_DRBlist_r17 = []ExcessDelay_DRB_IdentityInfo_r17{}
	for _, i := range tmp_ExcessDelay_DRBlist_r17.Value {
		ie.ExcessDelay_DRBlist_r17 = append(ie.ExcessDelay_DRBlist_r17, *i)
	}
	return nil
}
