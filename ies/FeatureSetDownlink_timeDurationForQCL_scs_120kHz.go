package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetDownlink_timeDurationForQCL_scs_120kHz_Enum_s14 aper.Enumerated = 0
	FeatureSetDownlink_timeDurationForQCL_scs_120kHz_Enum_s28 aper.Enumerated = 1
)

type FeatureSetDownlink_timeDurationForQCL_scs_120kHz struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_120kHz) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlink_timeDurationForQCL_scs_120kHz", err)
	}
	return nil
}

func (ie *FeatureSetDownlink_timeDurationForQCL_scs_120kHz) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlink_timeDurationForQCL_scs_120kHz", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
