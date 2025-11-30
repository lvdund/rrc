package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandCombination_v1610_powerClass_v1610_Enum_pc1dot5 aper.Enumerated = 0
)

type BandCombination_v1610_powerClass_v1610 struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *BandCombination_v1610_powerClass_v1610) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode BandCombination_v1610_powerClass_v1610", err)
	}
	return nil
}

func (ie *BandCombination_v1610_powerClass_v1610) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode BandCombination_v1610_powerClass_v1610", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
