package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_scalingFactor_Enum_f0p4  aper.Enumerated = 0
	FeatureSetDownlink_scalingFactor_Enum_f0p75 aper.Enumerated = 1
	FeatureSetDownlink_scalingFactor_Enum_f0p8  aper.Enumerated = 2
)

type FeatureSetDownlink_scalingFactor struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *FeatureSetDownlink_scalingFactor) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_scalingFactor", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_scalingFactor) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_scalingFactor", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
