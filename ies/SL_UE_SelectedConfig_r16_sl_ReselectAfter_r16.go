package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n1 aper.Enumerated = 0
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n2 aper.Enumerated = 1
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n3 aper.Enumerated = 2
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n4 aper.Enumerated = 3
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n5 aper.Enumerated = 4
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n6 aper.Enumerated = 5
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n7 aper.Enumerated = 6
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n8 aper.Enumerated = 7
	SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16_Enum_n9 aper.Enumerated = 8
)

type SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16 struct {
	Value aper.Enumerated `lb:0,ub:8,madatory`
}

func (ie *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Encode SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16", err)
	}
	return nil
}

func (ie *SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 8}, false); err != nil {
		return utils.WrapError("Decode SL_UE_SelectedConfig_r16_sl_ReselectAfter_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
