package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityCombinationsPerCellIndex_r16 struct {
	Value uint64 `lb:0,ub:maxNrofDUCells_r16,madatory`
}

func (ie *AvailabilityCombinationsPerCellIndex_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: maxNrofDUCells_r16}, false); err != nil {
		return utils.WrapError("Encode AvailabilityCombinationsPerCellIndex_r16", err)
	}
	return nil
}

func (ie *AvailabilityCombinationsPerCellIndex_r16) Decode(r *aper.AperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofDUCells_r16}, false); err != nil {
		return utils.WrapError("Decode AvailabilityCombinationsPerCellIndex_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
