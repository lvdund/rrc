package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms3      aper.Enumerated = 0
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms10     aper.Enumerated = 1
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms20     aper.Enumerated = 2
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms25     aper.Enumerated = 3
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms30     aper.Enumerated = 4
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms40     aper.Enumerated = 5
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms50     aper.Enumerated = 6
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms60     aper.Enumerated = 7
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms75     aper.Enumerated = 8
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms100    aper.Enumerated = 9
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms150    aper.Enumerated = 10
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms200    aper.Enumerated = 11
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms250    aper.Enumerated = 12
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms300    aper.Enumerated = 13
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms500    aper.Enumerated = 14
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms750    aper.Enumerated = 15
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_ms1500   aper.Enumerated = 16
	SL_PDCP_Config_r16_sl_DiscardTimer_r16_Enum_infinity aper.Enumerated = 17
)

type SL_PDCP_Config_r16_sl_DiscardTimer_r16 struct {
	Value aper.Enumerated `lb:0,ub:17,madatory`
}

func (ie *SL_PDCP_Config_r16_sl_DiscardTimer_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Encode SL_PDCP_Config_r16_sl_DiscardTimer_r16", err)
	}
	return nil
}

func (ie *SL_PDCP_Config_r16_sl_DiscardTimer_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Decode SL_PDCP_Config_r16_sl_DiscardTimer_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
