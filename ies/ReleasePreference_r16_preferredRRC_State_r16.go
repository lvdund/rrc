package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReleasePreference_r16_preferredRRC_State_r16_Enum_idle           aper.Enumerated = 0
	ReleasePreference_r16_preferredRRC_State_r16_Enum_inactive       aper.Enumerated = 1
	ReleasePreference_r16_preferredRRC_State_r16_Enum_connected      aper.Enumerated = 2
	ReleasePreference_r16_preferredRRC_State_r16_Enum_outOfConnected aper.Enumerated = 3
)

type ReleasePreference_r16_preferredRRC_State_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ReleasePreference_r16_preferredRRC_State_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ReleasePreference_r16_preferredRRC_State_r16", err)
	}
	return nil
}

func (ie *ReleasePreference_r16_preferredRRC_State_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ReleasePreference_r16_preferredRRC_State_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
