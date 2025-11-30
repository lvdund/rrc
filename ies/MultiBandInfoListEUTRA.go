package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MultiBandInfoListEUTRA struct {
	Value []FreqBandIndicatorEUTRA `lb:1,ub:maxMultiBands,madatory`
}

func (ie *MultiBandInfoListEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, aper.Constraint{Lb: 1, Ub: maxMultiBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MultiBandInfoListEUTRA", err)
	}
	return nil
}

func (ie *MultiBandInfoListEUTRA) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*FreqBandIndicatorEUTRA]([]*FreqBandIndicatorEUTRA{}, aper.Constraint{Lb: 1, Ub: maxMultiBands}, false)
	fn := func() *FreqBandIndicatorEUTRA {
		return new(FreqBandIndicatorEUTRA)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MultiBandInfoListEUTRA", err)
	}
	ie.Value = []FreqBandIndicatorEUTRA{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
