package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n1  aper.Enumerated = 0
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n5  aper.Enumerated = 1
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n10 aper.Enumerated = 2
	SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16_Enum_n20 aper.Enumerated = 3
)

type SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16", err)
	}
	return nil
}

func (ie *SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode SL_SelectionWindowConfig_r16_sl_SelectionWindow_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
