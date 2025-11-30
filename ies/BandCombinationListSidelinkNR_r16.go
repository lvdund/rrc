package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationListSidelinkNR_r16 struct {
	Value []BandCombinationParametersSidelinkNR_r16 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationListSidelinkNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationParametersSidelinkNR_r16]([]*BandCombinationParametersSidelinkNR_r16{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationListSidelinkNR_r16", err)
	}
	return nil
}

func (ie *BandCombinationListSidelinkNR_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationParametersSidelinkNR_r16]([]*BandCombinationParametersSidelinkNR_r16{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombinationParametersSidelinkNR_r16 {
		return new(BandCombinationParametersSidelinkNR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationListSidelinkNR_r16", err)
	}
	ie.Value = []BandCombinationParametersSidelinkNR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
