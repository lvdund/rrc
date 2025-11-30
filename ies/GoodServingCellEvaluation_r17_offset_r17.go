package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GoodServingCellEvaluation_r17_offset_r17_Enum_db2 aper.Enumerated = 0
	GoodServingCellEvaluation_r17_offset_r17_Enum_db4 aper.Enumerated = 1
	GoodServingCellEvaluation_r17_offset_r17_Enum_db6 aper.Enumerated = 2
	GoodServingCellEvaluation_r17_offset_r17_Enum_db8 aper.Enumerated = 3
)

type GoodServingCellEvaluation_r17_offset_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *GoodServingCellEvaluation_r17_offset_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode GoodServingCellEvaluation_r17_offset_r17", err)
	}
	return nil
}

func (ie *GoodServingCellEvaluation_r17_offset_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode GoodServingCellEvaluation_r17_offset_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
