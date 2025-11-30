package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CellsToAddModExt_v1710_ntn_PolarizationDL_r17_Enum_rhcp   aper.Enumerated = 0
	CellsToAddModExt_v1710_ntn_PolarizationDL_r17_Enum_lhcp   aper.Enumerated = 1
	CellsToAddModExt_v1710_ntn_PolarizationDL_r17_Enum_linear aper.Enumerated = 2
)

type CellsToAddModExt_v1710_ntn_PolarizationDL_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CellsToAddModExt_v1710_ntn_PolarizationDL_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CellsToAddModExt_v1710_ntn_PolarizationDL_r17", err)
	}
	return nil
}

func (ie *CellsToAddModExt_v1710_ntn_PolarizationDL_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CellsToAddModExt_v1710_ntn_PolarizationDL_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
