package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Config_FR2_r16_mpe_Threshold_r16_Enum_dB3  aper.Enumerated = 0
	MPE_Config_FR2_r16_mpe_Threshold_r16_Enum_dB6  aper.Enumerated = 1
	MPE_Config_FR2_r16_mpe_Threshold_r16_Enum_dB9  aper.Enumerated = 2
	MPE_Config_FR2_r16_mpe_Threshold_r16_Enum_dB12 aper.Enumerated = 3
)

type MPE_Config_FR2_r16_mpe_Threshold_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MPE_Config_FR2_r16_mpe_Threshold_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MPE_Config_FR2_r16_mpe_Threshold_r16", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r16_mpe_Threshold_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MPE_Config_FR2_r16_mpe_Threshold_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
