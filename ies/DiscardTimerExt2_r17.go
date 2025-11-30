package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DiscardTimerExt2_r17_Enum_ms2000 aper.Enumerated = 0
	DiscardTimerExt2_r17_Enum_spare3 aper.Enumerated = 1
	DiscardTimerExt2_r17_Enum_spare2 aper.Enumerated = 2
	DiscardTimerExt2_r17_Enum_spare1 aper.Enumerated = 3
)

type DiscardTimerExt2_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *DiscardTimerExt2_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode DiscardTimerExt2_r17", err)
	}
	return nil
}

func (ie *DiscardTimerExt2_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode DiscardTimerExt2_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
