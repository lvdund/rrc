package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CellsToAddModListUTRA_FDD_r16 struct {
	Value []CellsToAddModUTRA_FDD_r16 `lb:1,ub:maxCellMeasUTRA_FDD_r16,madatory`
}

func (ie *CellsToAddModListUTRA_FDD_r16) Encode(w *aper.AperWriter) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddModUTRA_FDD_r16]([]*CellsToAddModUTRA_FDD_r16{}, aper.Constraint{Lb: 1, Ub: maxCellMeasUTRA_FDD_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode CellsToAddModListUTRA_FDD_r16", err)
	}
	return nil
}

func (ie *CellsToAddModListUTRA_FDD_r16) Decode(r *aper.AperReader) error {
	var err error
	tmp := utils.NewSequence[*CellsToAddModUTRA_FDD_r16]([]*CellsToAddModUTRA_FDD_r16{}, aper.Constraint{Lb: 1, Ub: maxCellMeasUTRA_FDD_r16}, false)
	fn := func() *CellsToAddModUTRA_FDD_r16 {
		return new(CellsToAddModUTRA_FDD_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode CellsToAddModListUTRA_FDD_r16", err)
	}
	ie.Value = []CellsToAddModUTRA_FDD_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
