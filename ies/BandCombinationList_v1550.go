package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationList_v1550 struct {
	Value []BandCombination_v1550 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationList_v1550) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v1550]([]*BandCombination_v1550{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationList_v1550", err)
	}
	return nil
}

func (ie *BandCombinationList_v1550) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombination_v1550]([]*BandCombination_v1550{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombination_v1550 {
		return new(BandCombination_v1550)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationList_v1550", err)
	}
	ie.Value = []BandCombination_v1550{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
