package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DummyA_maxNumberCS_IM_PerCC_Enum_n1  aper.Enumerated = 0
	DummyA_maxNumberCS_IM_PerCC_Enum_n2  aper.Enumerated = 1
	DummyA_maxNumberCS_IM_PerCC_Enum_n4  aper.Enumerated = 2
	DummyA_maxNumberCS_IM_PerCC_Enum_n8  aper.Enumerated = 3
	DummyA_maxNumberCS_IM_PerCC_Enum_n16 aper.Enumerated = 4
	DummyA_maxNumberCS_IM_PerCC_Enum_n32 aper.Enumerated = 5
)

type DummyA_maxNumberCS_IM_PerCC struct {
	Value aper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *DummyA_maxNumberCS_IM_PerCC) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode DummyA_maxNumberCS_IM_PerCC", err)
	}
	return nil
}

func (ie *DummyA_maxNumberCS_IM_PerCC) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode DummyA_maxNumberCS_IM_PerCC", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
