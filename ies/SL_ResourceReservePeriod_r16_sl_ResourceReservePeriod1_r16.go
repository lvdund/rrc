package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms0    aper.Enumerated = 0
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms100  aper.Enumerated = 1
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms200  aper.Enumerated = 2
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms300  aper.Enumerated = 3
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms400  aper.Enumerated = 4
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms500  aper.Enumerated = 5
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms600  aper.Enumerated = 6
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms700  aper.Enumerated = 7
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms800  aper.Enumerated = 8
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms900  aper.Enumerated = 9
	SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16_Enum_ms1000 aper.Enumerated = 10
)

type SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16 struct {
	Value aper.Enumerated `lb:0,ub:10,madatory`
}

func (ie *SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Encode SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16", err)
	}
	return nil
}

func (ie *SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 10}, false); err != nil {
		return utils.WrapError("Decode SL_ResourceReservePeriod_r16_sl_ResourceReservePeriod1_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
