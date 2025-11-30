package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NeedForNCSG_BandListNR_r17 struct {
	Value []NeedForNCSG_NR_r17 `lb:1,ub:maxBands,madatory`
}

func (ie *NeedForNCSG_BandListNR_r17) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*NeedForNCSG_NR_r17]([]*NeedForNCSG_NR_r17{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode NeedForNCSG_BandListNR_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_BandListNR_r17) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*NeedForNCSG_NR_r17]([]*NeedForNCSG_NR_r17{}, aper.Constraint{Lb: 1, Ub: maxBands}, false)
	fn := func() *NeedForNCSG_NR_r17 {
		return new(NeedForNCSG_NR_r17)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode NeedForNCSG_BandListNR_r17", err)
	}
	ie.Value = []NeedForNCSG_NR_r17{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
