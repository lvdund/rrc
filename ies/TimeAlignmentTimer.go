package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	TimeAlignmentTimer_Enum_ms500    aper.Enumerated = 0
	TimeAlignmentTimer_Enum_ms750    aper.Enumerated = 1
	TimeAlignmentTimer_Enum_ms1280   aper.Enumerated = 2
	TimeAlignmentTimer_Enum_ms1920   aper.Enumerated = 3
	TimeAlignmentTimer_Enum_ms2560   aper.Enumerated = 4
	TimeAlignmentTimer_Enum_ms5120   aper.Enumerated = 5
	TimeAlignmentTimer_Enum_ms10240  aper.Enumerated = 6
	TimeAlignmentTimer_Enum_infinity aper.Enumerated = 7
)

type TimeAlignmentTimer struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *TimeAlignmentTimer) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode TimeAlignmentTimer", err)
	}
	return nil
}

func (ie *TimeAlignmentTimer) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode TimeAlignmentTimer", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
