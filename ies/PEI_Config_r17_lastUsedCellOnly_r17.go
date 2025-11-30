package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PEI_Config_r17_lastUsedCellOnly_r17_Enum_true aper.Enumerated = 0
)

type PEI_Config_r17_lastUsedCellOnly_r17 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PEI_Config_r17_lastUsedCellOnly_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PEI_Config_r17_lastUsedCellOnly_r17", err)
	}
	return nil
}

func (ie *PEI_Config_r17_lastUsedCellOnly_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PEI_Config_r17_lastUsedCellOnly_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
