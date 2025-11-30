package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m5     aper.Enumerated = 0
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m10    aper.Enumerated = 1
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m20    aper.Enumerated = 2
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m30    aper.Enumerated = 3
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m40    aper.Enumerated = 4
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_m50    aper.Enumerated = 5
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_spare2 aper.Enumerated = 6
	SL_ZoneConfig_r16_sl_ZoneLength_r16_Enum_spare1 aper.Enumerated = 7
)

type SL_ZoneConfig_r16_sl_ZoneLength_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_ZoneConfig_r16_sl_ZoneLength_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_ZoneConfig_r16_sl_ZoneLength_r16", err)
	}
	return nil
}

func (ie *SL_ZoneConfig_r16_sl_ZoneLength_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_ZoneConfig_r16_sl_ZoneLength_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
