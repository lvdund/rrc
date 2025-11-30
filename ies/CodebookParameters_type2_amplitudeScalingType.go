package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookParameters_type2_amplitudeScalingType_Enum_wideband           aper.Enumerated = 0
	CodebookParameters_type2_amplitudeScalingType_Enum_widebandAndSubband aper.Enumerated = 1
)

type CodebookParameters_type2_amplitudeScalingType struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *CodebookParameters_type2_amplitudeScalingType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode CodebookParameters_type2_amplitudeScalingType", err)
	}
	return nil
}

func (ie *CodebookParameters_type2_amplitudeScalingType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode CodebookParameters_type2_amplitudeScalingType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
