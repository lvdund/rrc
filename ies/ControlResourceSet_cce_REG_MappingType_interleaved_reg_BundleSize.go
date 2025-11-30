package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize_Enum_n2 aper.Enumerated = 0
	ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize_Enum_n3 aper.Enumerated = 1
	ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize_Enum_n6 aper.Enumerated = 2
)

type ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize", err)
	}
	return nil
}

func (ie *ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSet_cce_REG_MappingType_interleaved_reg_BundleSize", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
