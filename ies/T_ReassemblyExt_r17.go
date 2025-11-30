package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	T_ReassemblyExt_r17_Enum_ms210  aper.Enumerated = 0
	T_ReassemblyExt_r17_Enum_ms220  aper.Enumerated = 1
	T_ReassemblyExt_r17_Enum_ms340  aper.Enumerated = 2
	T_ReassemblyExt_r17_Enum_ms350  aper.Enumerated = 3
	T_ReassemblyExt_r17_Enum_ms550  aper.Enumerated = 4
	T_ReassemblyExt_r17_Enum_ms1100 aper.Enumerated = 5
	T_ReassemblyExt_r17_Enum_ms1650 aper.Enumerated = 6
	T_ReassemblyExt_r17_Enum_ms2200 aper.Enumerated = 7
)

type T_ReassemblyExt_r17 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *T_ReassemblyExt_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode T_ReassemblyExt_r17", err)
	}
	return nil
}

func (ie *T_ReassemblyExt_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode T_ReassemblyExt_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
