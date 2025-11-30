package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0     aper.Enumerated = 0
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot1 aper.Enumerated = 1
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot2 aper.Enumerated = 2
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot3 aper.Enumerated = 3
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot4 aper.Enumerated = 4
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s0dot5 aper.Enumerated = 5
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s1     aper.Enumerated = 6
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s2     aper.Enumerated = 7
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s3     aper.Enumerated = 8
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s4     aper.Enumerated = 9
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s5     aper.Enumerated = 10
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s6     aper.Enumerated = 11
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s7     aper.Enumerated = 12
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s8     aper.Enumerated = 13
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s9     aper.Enumerated = 14
	MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17_Enum_s10    aper.Enumerated = 15
)

type MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode MUSIM_GapAssistanceConfig_r17_musim_GapProhibitTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
