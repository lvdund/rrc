package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16_Enum_semi_static_mode1 aper.Enumerated = 0
	ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16_Enum_semi_static_mode2 aper.Enumerated = 1
	ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16_Enum_dynamic           aper.Enumerated = 2
)

type ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16", err)
	}
	return nil
}

func (ie *ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
