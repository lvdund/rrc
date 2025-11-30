package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s0    aper.Enumerated = 0
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s1    aper.Enumerated = 1
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s2    aper.Enumerated = 2
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s4    aper.Enumerated = 3
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s8    aper.Enumerated = 4
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s10   aper.Enumerated = 5
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s15   aper.Enumerated = 6
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s30   aper.Enumerated = 7
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s60   aper.Enumerated = 8
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s120  aper.Enumerated = 9
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s180  aper.Enumerated = 10
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s240  aper.Enumerated = 11
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s300  aper.Enumerated = 12
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s600  aper.Enumerated = 13
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s900  aper.Enumerated = 14
	SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17_Enum_s1800 aper.Enumerated = 15
)

type SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	return nil
}

func (ie *SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SCG_DeactivationPreferenceConfig_r17_scg_DeactivationPreferenceProhibitTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
