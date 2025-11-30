package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec10  aper.Enumerated = 0
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec30  aper.Enumerated = 1
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec60  aper.Enumerated = 2
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec120 aper.Enumerated = 3
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec180 aper.Enumerated = 4
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec240 aper.Enumerated = 5
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_sec300 aper.Enumerated = 6
	VarMeasIdleConfig_r16_measIdleDuration_r16_Enum_spare  aper.Enumerated = 7
)

type VarMeasIdleConfig_r16_measIdleDuration_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *VarMeasIdleConfig_r16_measIdleDuration_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode VarMeasIdleConfig_r16_measIdleDuration_r16", err)
	}
	return nil
}

func (ie *VarMeasIdleConfig_r16_measIdleDuration_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode VarMeasIdleConfig_r16_measIdleDuration_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
