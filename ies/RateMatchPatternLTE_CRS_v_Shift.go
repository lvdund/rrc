package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternLTE_CRS_v_Shift_Enum_n0 aper.Enumerated = 0
	RateMatchPatternLTE_CRS_v_Shift_Enum_n1 aper.Enumerated = 1
	RateMatchPatternLTE_CRS_v_Shift_Enum_n2 aper.Enumerated = 2
	RateMatchPatternLTE_CRS_v_Shift_Enum_n3 aper.Enumerated = 3
	RateMatchPatternLTE_CRS_v_Shift_Enum_n4 aper.Enumerated = 4
	RateMatchPatternLTE_CRS_v_Shift_Enum_n5 aper.Enumerated = 5
)

type RateMatchPatternLTE_CRS_v_Shift struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *RateMatchPatternLTE_CRS_v_Shift) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternLTE_CRS_v_Shift", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS_v_Shift) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternLTE_CRS_v_Shift", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
