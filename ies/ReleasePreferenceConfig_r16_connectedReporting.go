package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReleasePreferenceConfig_r16_connectedReporting_Enum_true aper.Enumerated = 0
)

type ReleasePreferenceConfig_r16_connectedReporting struct {
	Value aper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *ReleasePreferenceConfig_r16_connectedReporting) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode ReleasePreferenceConfig_r16_connectedReporting", err)
	}
	return nil
}

func (ie *ReleasePreferenceConfig_r16_connectedReporting) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode ReleasePreferenceConfig_r16_connectedReporting", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
