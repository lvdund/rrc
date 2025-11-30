package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DiscardTimerExt_r16_Enum_ms0dot5 aper.Enumerated = 0
	DiscardTimerExt_r16_Enum_ms1     aper.Enumerated = 1
	DiscardTimerExt_r16_Enum_ms2     aper.Enumerated = 2
	DiscardTimerExt_r16_Enum_ms4     aper.Enumerated = 3
	DiscardTimerExt_r16_Enum_ms6     aper.Enumerated = 4
	DiscardTimerExt_r16_Enum_ms8     aper.Enumerated = 5
	DiscardTimerExt_r16_Enum_spare2  aper.Enumerated = 6
	DiscardTimerExt_r16_Enum_spare1  aper.Enumerated = 7
)

type DiscardTimerExt_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *DiscardTimerExt_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode DiscardTimerExt_r16", err)
	}
	return nil
}

func (ie *DiscardTimerExt_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode DiscardTimerExt_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
