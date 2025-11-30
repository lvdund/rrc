package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList struct {
	Value []BandCombination `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination]([]*BandCombination{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList", err)
	}
	return nil
}

func (ie *BandCombinationList) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination]([]*BandCombination{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination {
		return new(BandCombination)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList", err)
	}
	ie.Value = []BandCombination{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
