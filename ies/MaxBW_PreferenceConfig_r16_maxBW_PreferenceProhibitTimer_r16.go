package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s0     aper.Enumerated = 0
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s0dot5 aper.Enumerated = 1
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s1     aper.Enumerated = 2
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s2     aper.Enumerated = 3
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s3     aper.Enumerated = 4
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s4     aper.Enumerated = 5
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s5     aper.Enumerated = 6
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s6     aper.Enumerated = 7
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s7     aper.Enumerated = 8
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s8     aper.Enumerated = 9
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s9     aper.Enumerated = 10
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s10    aper.Enumerated = 11
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s20    aper.Enumerated = 12
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_s30    aper.Enumerated = 13
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_spare2 aper.Enumerated = 14
	MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16_Enum_spare1 aper.Enumerated = 15
)

type MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MaxBW_PreferenceConfig_r16_maxBW_PreferenceProhibitTimer_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
