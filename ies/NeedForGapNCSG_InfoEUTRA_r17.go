package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapNCSG_InfoEUTRA_r17 struct {
	NeedForNCSG_EUTRA_r17 []NeedForNCSG_EUTRA_r17 `lb:1,ub:maxBandsEUTRA,madatory`
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp_NeedForNCSG_EUTRA_r17 := utils.NewSequence[*NeedForNCSG_EUTRA_r17]([]*NeedForNCSG_EUTRA_r17{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	for _, i := range ie.NeedForNCSG_EUTRA_r17 {
		tmp_NeedForNCSG_EUTRA_r17.Value = append(tmp_NeedForNCSG_EUTRA_r17.Value, &i)
	}
	if err = tmp_NeedForNCSG_EUTRA_r17.Encode(w); err != nil {
		return utils.WrapError("Encode NeedForNCSG_EUTRA_r17", err)
	}
	return nil
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp_NeedForNCSG_EUTRA_r17 := utils.NewSequence[*NeedForNCSG_EUTRA_r17]([]*NeedForNCSG_EUTRA_r17{}, aper.Constraint{Lb: 1, Ub: maxBandsEUTRA}, false)
	fn_NeedForNCSG_EUTRA_r17 := func() *NeedForNCSG_EUTRA_r17 {
		return new(NeedForNCSG_EUTRA_r17)
	}
	if err = tmp_NeedForNCSG_EUTRA_r17.Decode(r, fn_NeedForNCSG_EUTRA_r17); err != nil {
		return utils.WrapError("Decode NeedForNCSG_EUTRA_r17", err)
	}
	ie.NeedForNCSG_EUTRA_r17 = []NeedForNCSG_EUTRA_r17{}
	for _, i := range tmp_NeedForNCSG_EUTRA_r17.Value {
		ie.NeedForNCSG_EUTRA_r17 = append(ie.NeedForNCSG_EUTRA_r17, *i)
	}
	return nil
}
