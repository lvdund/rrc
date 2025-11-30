package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g0 aper.Enumerated = 0
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g1 aper.Enumerated = 1
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g2 aper.Enumerated = 2
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g3 aper.Enumerated = 3
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g4 aper.Enumerated = 4
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g5 aper.Enumerated = 5
	FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16_Enum_g6 aper.Enumerated = 6
)

type FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
