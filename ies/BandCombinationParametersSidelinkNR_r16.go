package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationParametersSidelinkNR_r16 struct {
	Value []BandParametersSidelink_r16 `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *BandCombinationParametersSidelinkNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelink_r16]([]*BandParametersSidelink_r16{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationParametersSidelinkNR_r16", err)
	}
	return nil
}

func (ie *BandCombinationParametersSidelinkNR_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandParametersSidelink_r16]([]*BandParametersSidelink_r16{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *BandParametersSidelink_r16 {
		return new(BandParametersSidelink_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationParametersSidelinkNR_r16", err)
	}
	ie.Value = []BandParametersSidelink_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
