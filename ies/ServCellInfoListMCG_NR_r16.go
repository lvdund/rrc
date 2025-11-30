package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoListMCG_NR_r16 struct {
	Value []ServCellInfoXCG_NR_r16 `lb:1,ub:maxNrofServingCells,madatory`
}

func (ie *ServCellInfoListMCG_NR_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_NR_r16]([]*ServCellInfoXCG_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ServCellInfoListMCG_NR_r16", err)
	}
	return nil
}

func (ie *ServCellInfoListMCG_NR_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*ServCellInfoXCG_NR_r16]([]*ServCellInfoXCG_NR_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofServingCells}, false)
	fn := func() *ServCellInfoXCG_NR_r16 {
		return new(ServCellInfoXCG_NR_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ServCellInfoListMCG_NR_r16", err)
	}
	ie.Value = []ServCellInfoXCG_NR_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
