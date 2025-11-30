package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16_Enum_n10 aper.Enumerated = 0
	SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16_Enum_n12 aper.Enumerated = 1
	SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16_Enum_n15 aper.Enumerated = 2
	SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16_Enum_n20 aper.Enumerated = 3
	SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16_Enum_n25 aper.Enumerated = 4
)

type SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16 struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16", err)
	}
	return nil
}

func (ie *SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
