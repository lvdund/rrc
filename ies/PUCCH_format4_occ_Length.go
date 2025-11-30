package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_format4_occ_Length_Enum_n2 aper.Enumerated = 0
	PUCCH_format4_occ_Length_Enum_n4 aper.Enumerated = 1
)

type PUCCH_format4_occ_Length struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PUCCH_format4_occ_Length) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PUCCH_format4_occ_Length", err)
	}
	return nil
}

func (ie *PUCCH_format4_occ_Length) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PUCCH_format4_occ_Length", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
