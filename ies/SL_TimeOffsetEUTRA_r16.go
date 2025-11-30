package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TimeOffsetEUTRA_r16_Enum_ms0       aper.Enumerated = 0
	SL_TimeOffsetEUTRA_r16_Enum_ms0dot25  aper.Enumerated = 1
	SL_TimeOffsetEUTRA_r16_Enum_ms0dot5   aper.Enumerated = 2
	SL_TimeOffsetEUTRA_r16_Enum_ms0dot625 aper.Enumerated = 3
	SL_TimeOffsetEUTRA_r16_Enum_ms0dot75  aper.Enumerated = 4
	SL_TimeOffsetEUTRA_r16_Enum_ms1       aper.Enumerated = 5
	SL_TimeOffsetEUTRA_r16_Enum_ms1dot25  aper.Enumerated = 6
	SL_TimeOffsetEUTRA_r16_Enum_ms1dot5   aper.Enumerated = 7
	SL_TimeOffsetEUTRA_r16_Enum_ms1dot75  aper.Enumerated = 8
	SL_TimeOffsetEUTRA_r16_Enum_ms2       aper.Enumerated = 9
	SL_TimeOffsetEUTRA_r16_Enum_ms2dot5   aper.Enumerated = 10
	SL_TimeOffsetEUTRA_r16_Enum_ms3       aper.Enumerated = 11
	SL_TimeOffsetEUTRA_r16_Enum_ms4       aper.Enumerated = 12
	SL_TimeOffsetEUTRA_r16_Enum_ms5       aper.Enumerated = 13
	SL_TimeOffsetEUTRA_r16_Enum_ms6       aper.Enumerated = 14
	SL_TimeOffsetEUTRA_r16_Enum_ms8       aper.Enumerated = 15
	SL_TimeOffsetEUTRA_r16_Enum_ms10      aper.Enumerated = 16
	SL_TimeOffsetEUTRA_r16_Enum_ms20      aper.Enumerated = 17
)

type SL_TimeOffsetEUTRA_r16 struct {
	Value aper.Enumerated `lb:0,ub:17,madatory`
}

func (ie *SL_TimeOffsetEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Encode SL_TimeOffsetEUTRA_r16", err)
	}
	return nil
}

func (ie *SL_TimeOffsetEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 17}, false); err != nil {
		return utils.WrapError("Decode SL_TimeOffsetEUTRA_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
