package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16_Enum_us100 aper.Enumerated = 0
	Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16_Enum_us200 aper.Enumerated = 1
)

type Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16", err)
	}
	return nil
}

func (ie *Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16_type1_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
