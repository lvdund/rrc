package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_enabled aper.Enumerated = 0
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl1     aper.Enumerated = 1
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl2     aper.Enumerated = 2
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl3     aper.Enumerated = 3
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl4     aper.Enumerated = 4
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl5     aper.Enumerated = 5
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl6     aper.Enumerated = 6
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl7     aper.Enumerated = 7
	SL_ResourcePool_r16_sl_PreemptionEnable_r16_Enum_pl8     aper.Enumerated = 8
)

type SL_ResourcePool_r16_sl_PreemptionEnable_r16 struct {
	Value aper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *SL_ResourcePool_r16_sl_PreemptionEnable_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode SL_ResourcePool_r16_sl_PreemptionEnable_r16", err)
	}
	return nil
}

func (ie *SL_ResourcePool_r16_sl_PreemptionEnable_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode SL_ResourcePool_r16_sl_PreemptionEnable_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
