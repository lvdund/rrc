package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16_Enum_sl2  aper.Enumerated = 0
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16_Enum_sl4  aper.Enumerated = 1
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16_Enum_sl8  aper.Enumerated = 2
	MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16_Enum_sl12 aper.Enumerated = 3
)

type MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
