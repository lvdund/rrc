package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_format4_occ_Index_Enum_n0 aper.Enumerated = 0
	PUCCH_format4_occ_Index_Enum_n1 aper.Enumerated = 1
	PUCCH_format4_occ_Index_Enum_n2 aper.Enumerated = 2
	PUCCH_format4_occ_Index_Enum_n3 aper.Enumerated = 3
)

type PUCCH_format4_occ_Index struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *PUCCH_format4_occ_Index) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PUCCH_format4_occ_Index", err)
	}
	return nil
}

func (ie *PUCCH_format4_occ_Index) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PUCCH_format4_occ_Index", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
