package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyE_amplitudeScalingType_Enum_wideband           aper.Enumerated = 0
	DummyE_amplitudeScalingType_Enum_widebandAndSubband aper.Enumerated = 1
)

type DummyE_amplitudeScalingType struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *DummyE_amplitudeScalingType) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode DummyE_amplitudeScalingType", err)
	}
	return nil
}

func (ie *DummyE_amplitudeScalingType) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode DummyE_amplitudeScalingType", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
