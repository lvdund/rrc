package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17_Enum_sl8  aper.Enumerated = 0
	MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17_Enum_sl16 aper.Enumerated = 1
	MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17_Enum_sl32 aper.Enumerated = 2
	MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17_Enum_sl48 aper.Enumerated = 3
)

type MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17", err)
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17_preferredK0_SCS_480kHz_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
