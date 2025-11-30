package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n1  aper.Enumerated = 0
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n2  aper.Enumerated = 1
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n4  aper.Enumerated = 2
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n8  aper.Enumerated = 3
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n16 aper.Enumerated = 4
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n32 aper.Enumerated = 5
	SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16_Enum_n64 aper.Enumerated = 6
)

type SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16 struct {
	Value aper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16", err)
	}
	return nil
}

func (ie *SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SL_SSB_TimeAllocation_r16_sl_NumSSB_WithinPeriod_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
