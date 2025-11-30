package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpeedStateScaleFactors_sf_High_Enum_oDot25 aper.Enumerated = 0
	SpeedStateScaleFactors_sf_High_Enum_oDot5  aper.Enumerated = 1
	SpeedStateScaleFactors_sf_High_Enum_oDot75 aper.Enumerated = 2
	SpeedStateScaleFactors_sf_High_Enum_lDot0  aper.Enumerated = 3
)

type SpeedStateScaleFactors_sf_High struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SpeedStateScaleFactors_sf_High) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SpeedStateScaleFactors_sf_High", err)
	}
	return nil
}

func (ie *SpeedStateScaleFactors_sf_High) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SpeedStateScaleFactors_sf_High", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
