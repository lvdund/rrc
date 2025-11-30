package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c1 aper.Enumerated = 0
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c2 aper.Enumerated = 1
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c3 aper.Enumerated = 2
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c4 aper.Enumerated = 3
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c5 aper.Enumerated = 4
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c6 aper.Enumerated = 5
	SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17_Enum_c7 aper.Enumerated = 6
)

type SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17", err)
	}
	return nil
}

func (ie *SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SL_PBPS_CPS_Config_r17_sl_AllowedResourceSelectionConfig_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
