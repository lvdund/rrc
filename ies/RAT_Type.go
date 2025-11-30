package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAT_Type_Enum_nr             aper.Enumerated = 0
	RAT_Type_Enum_eutra_nr       aper.Enumerated = 1
	RAT_Type_Enum_eutra          aper.Enumerated = 2
	RAT_Type_Enum_utra_fdd_v1610 aper.Enumerated = 3
)

type RAT_Type struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RAT_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RAT_Type", err)
	}
	return nil
}

func (ie *RAT_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RAT_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
