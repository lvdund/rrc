package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s0       aper.Enumerated = 0
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s0dot5   aper.Enumerated = 1
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s1       aper.Enumerated = 2
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s2       aper.Enumerated = 3
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s3       aper.Enumerated = 4
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s4       aper.Enumerated = 5
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s5       aper.Enumerated = 6
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s6       aper.Enumerated = 7
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s7       aper.Enumerated = 8
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s8       aper.Enumerated = 9
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s9       aper.Enumerated = 10
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s10      aper.Enumerated = 11
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s20      aper.Enumerated = 12
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_s30      aper.Enumerated = 13
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_infinity aper.Enumerated = 14
	ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16_Enum_spare1   aper.Enumerated = 15
)

type ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode ReleasePreferenceConfig_r16_releasePreferenceProhibitTimer_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
