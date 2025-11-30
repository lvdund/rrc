package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf0    aper.Enumerated = 0
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf10   aper.Enumerated = 1
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf20   aper.Enumerated = 2
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf50   aper.Enumerated = 3
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf100  aper.Enumerated = 4
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf200  aper.Enumerated = 5
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf500  aper.Enumerated = 6
	MPE_Config_FR2_r17_mpe_ProhibitTimer_r17_Enum_sf1000 aper.Enumerated = 7
)

type MPE_Config_FR2_r17_mpe_ProhibitTimer_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MPE_Config_FR2_r17_mpe_ProhibitTimer_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MPE_Config_FR2_r17_mpe_ProhibitTimer_r17", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r17_mpe_ProhibitTimer_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MPE_Config_FR2_r17_mpe_ProhibitTimer_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
