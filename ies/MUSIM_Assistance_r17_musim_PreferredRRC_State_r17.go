package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_idle           aper.Enumerated = 0
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_inactive       aper.Enumerated = 1
	MUSIM_Assistance_r17_musim_PreferredRRC_State_r17_Enum_outOfConnected aper.Enumerated = 2
)

type MUSIM_Assistance_r17_musim_PreferredRRC_State_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MUSIM_Assistance_r17_musim_PreferredRRC_State_r17", err)
	}
	return nil
}

func (ie *MUSIM_Assistance_r17_musim_PreferredRRC_State_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MUSIM_Assistance_r17_musim_PreferredRRC_State_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
