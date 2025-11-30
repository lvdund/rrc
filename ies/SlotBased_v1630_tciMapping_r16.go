package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SlotBased_v1630_tciMapping_r16_Enum_cyclicMapping     aper.Enumerated = 0
	SlotBased_v1630_tciMapping_r16_Enum_sequentialMapping aper.Enumerated = 1
)

type SlotBased_v1630_tciMapping_r16 struct {
	Value aper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SlotBased_v1630_tciMapping_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SlotBased_v1630_tciMapping_r16", err)
	}
	return nil
}

func (ie *SlotBased_v1630_tciMapping_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SlotBased_v1630_tciMapping_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
