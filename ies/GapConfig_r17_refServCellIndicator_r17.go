package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	GapConfig_r17_refServCellIndicator_r17_Enum_pCell   aper.Enumerated = 0
	GapConfig_r17_refServCellIndicator_r17_Enum_pSCell  aper.Enumerated = 1
	GapConfig_r17_refServCellIndicator_r17_Enum_mcg_FR2 aper.Enumerated = 2
)

type GapConfig_r17_refServCellIndicator_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *GapConfig_r17_refServCellIndicator_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode GapConfig_r17_refServCellIndicator_r17", err)
	}
	return nil
}

func (ie *GapConfig_r17_refServCellIndicator_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode GapConfig_r17_refServCellIndicator_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
