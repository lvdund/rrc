package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms10    aper.Enumerated = 0
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms20    aper.Enumerated = 1
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms32    aper.Enumerated = 2
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms40    aper.Enumerated = 3
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms60    aper.Enumerated = 4
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms64    aper.Enumerated = 5
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms70    aper.Enumerated = 6
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms80    aper.Enumerated = 7
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms128   aper.Enumerated = 8
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms160   aper.Enumerated = 9
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms256   aper.Enumerated = 10
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms320   aper.Enumerated = 11
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms512   aper.Enumerated = 12
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms640   aper.Enumerated = 13
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms1024  aper.Enumerated = 14
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms1280  aper.Enumerated = 15
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms2048  aper.Enumerated = 16
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms2560  aper.Enumerated = 17
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms5120  aper.Enumerated = 18
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_ms10240 aper.Enumerated = 19
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare12 aper.Enumerated = 20
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare11 aper.Enumerated = 21
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare10 aper.Enumerated = 22
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare9  aper.Enumerated = 23
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare8  aper.Enumerated = 24
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare7  aper.Enumerated = 25
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare6  aper.Enumerated = 26
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare5  aper.Enumerated = 27
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare4  aper.Enumerated = 28
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare3  aper.Enumerated = 29
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare2  aper.Enumerated = 30
	DRX_Preference_r16_preferredDRX_LongCycle_r16_Enum_spare1  aper.Enumerated = 31
)

type DRX_Preference_r16_preferredDRX_LongCycle_r16 struct {
	Value aper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *DRX_Preference_r16_preferredDRX_LongCycle_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode DRX_Preference_r16_preferredDRX_LongCycle_r16", err)
	}
	return nil
}

func (ie *DRX_Preference_r16_preferredDRX_LongCycle_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode DRX_Preference_r16_preferredDRX_LongCycle_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
