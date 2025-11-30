package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms10   aper.Enumerated = 0
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms20   aper.Enumerated = 1
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms40   aper.Enumerated = 2
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms60   aper.Enumerated = 3
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms80   aper.Enumerated = 4
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_ms100  aper.Enumerated = 5
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_spare2 aper.Enumerated = 6
	MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17_Enum_spare1 aper.Enumerated = 7
)

type MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17", err)
	}
	return nil
}

func (ie *MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MUSIM_LeaveAssistanceConfig_r17_musim_LeaveWithoutResponseTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
