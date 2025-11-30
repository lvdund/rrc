package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ScalingFactorSidelink_r16_Enum_f0p4  aper.Enumerated = 0
	ScalingFactorSidelink_r16_Enum_f0p75 aper.Enumerated = 1
	ScalingFactorSidelink_r16_Enum_f0p8  aper.Enumerated = 2
	ScalingFactorSidelink_r16_Enum_f1    aper.Enumerated = 3
)

type ScalingFactorSidelink_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ScalingFactorSidelink_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ScalingFactorSidelink_r16", err)
	}
	return nil
}

func (ie *ScalingFactorSidelink_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ScalingFactorSidelink_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
