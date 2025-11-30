package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot125 aper.Enumerated = 0
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot25  aper.Enumerated = 1
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms0dot5   aper.Enumerated = 2
	UL_GapFR2_Config_r17_ugl_r17_Enum_ms1       aper.Enumerated = 3
)

type UL_GapFR2_Config_r17_ugl_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *UL_GapFR2_Config_r17_ugl_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode UL_GapFR2_Config_r17_ugl_r17", err)
	}
	return nil
}

func (ie *UL_GapFR2_Config_r17_ugl_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode UL_GapFR2_Config_r17_ugl_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
