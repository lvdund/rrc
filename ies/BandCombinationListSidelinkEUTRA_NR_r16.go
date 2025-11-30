package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationListSidelinkEUTRA_NR_r16 struct {
	Value []BandCombinationParametersSidelinkEUTRA_NR_r16 `lb:1,ub:maxBandComb,madatory`
}

func (ie *BandCombinationListSidelinkEUTRA_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationParametersSidelinkEUTRA_NR_r16]([]*BandCombinationParametersSidelinkEUTRA_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode BandCombinationListSidelinkEUTRA_NR_r16", err)
	}
	return nil
}

func (ie *BandCombinationListSidelinkEUTRA_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*BandCombinationParametersSidelinkEUTRA_NR_r16]([]*BandCombinationParametersSidelinkEUTRA_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxBandComb}, false)
	fn := func() *BandCombinationParametersSidelinkEUTRA_NR_r16 {
		return new(BandCombinationParametersSidelinkEUTRA_NR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode BandCombinationListSidelinkEUTRA_NR_r16", err)
	}
	ie.Value = []BandCombinationParametersSidelinkEUTRA_NR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
