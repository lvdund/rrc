package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UL_DelayValueConfig_r16 struct {
	Delay_DRBlist_r16 []DRB_Identity `lb:1,ub:maxDRB,madatory`
}

func (ie *UL_DelayValueConfig_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp_Delay_DRBlist_r16 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	for _, i := range ie.Delay_DRBlist_r16 {
		tmp_Delay_DRBlist_r16.Value = append(tmp_Delay_DRBlist_r16.Value, &i)
	}
	if err = tmp_Delay_DRBlist_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Delay_DRBlist_r16", err)
	}
	return nil
}

func (ie *UL_DelayValueConfig_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp_Delay_DRBlist_r16 := utils.NewSequence[*DRB_Identity]([]*DRB_Identity{}, aper.Constraint{Lb: 1, Ub: maxDRB}, false)
	fn_Delay_DRBlist_r16 := func() *DRB_Identity {
		return new(DRB_Identity)
	}
	if err = tmp_Delay_DRBlist_r16.Decode(r, fn_Delay_DRBlist_r16); err != nil {
		return utils.WrapError("Decode Delay_DRBlist_r16", err)
	}
	ie.Delay_DRBlist_r16 = []DRB_Identity{}
	for _, i := range tmp_Delay_DRBlist_r16.Value {
		ie.Delay_DRBlist_r16 = append(ie.Delay_DRBlist_r16, *i)
	}
	return nil
}
