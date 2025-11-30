package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PreConfig_r17_type_r17_Enum_type1A aper.Enumerated = 0
	DL_PPW_PreConfig_r17_type_r17_Enum_type1B aper.Enumerated = 1
	DL_PPW_PreConfig_r17_type_r17_Enum_type2  aper.Enumerated = 2
)

type DL_PPW_PreConfig_r17_type_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DL_PPW_PreConfig_r17_type_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DL_PPW_PreConfig_r17_type_r17", err)
	}
	return nil
}

func (ie *DL_PPW_PreConfig_r17_type_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DL_PPW_PreConfig_r17_type_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
