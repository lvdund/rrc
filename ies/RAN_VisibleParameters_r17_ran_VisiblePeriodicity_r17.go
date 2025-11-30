package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms120  aper.Enumerated = 0
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms240  aper.Enumerated = 1
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms480  aper.Enumerated = 2
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms640  aper.Enumerated = 3
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms1024 aper.Enumerated = 4
)

type RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17", err)
	}
	return nil
}

func (ie *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
