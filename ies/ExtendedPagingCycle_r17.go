package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ExtendedPagingCycle_r17_Enum_rf256  aper.Enumerated = 0
	ExtendedPagingCycle_r17_Enum_rf512  aper.Enumerated = 1
	ExtendedPagingCycle_r17_Enum_rf1024 aper.Enumerated = 2
	ExtendedPagingCycle_r17_Enum_spare1 aper.Enumerated = 3
)

type ExtendedPagingCycle_r17 struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ExtendedPagingCycle_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ExtendedPagingCycle_r17", err)
	}
	return nil
}

func (ie *ExtendedPagingCycle_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ExtendedPagingCycle_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
