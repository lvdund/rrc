package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16_Enum_offset01 aper.Enumerated = 0
	SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16_Enum_offset10 aper.Enumerated = 1
	SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16_Enum_offset11 aper.Enumerated = 2
)

type SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16", err)
	}
	return nil
}

func (ie *SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode SL_PTRS_Config_r16_sl_PTRS_RE_Offset_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
