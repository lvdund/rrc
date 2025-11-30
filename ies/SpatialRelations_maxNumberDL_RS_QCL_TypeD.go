package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	SpatialRelations_maxNumberDL_RS_QCL_TypeD_Enum_n1  aper.Enumerated = 0
	SpatialRelations_maxNumberDL_RS_QCL_TypeD_Enum_n2  aper.Enumerated = 1
	SpatialRelations_maxNumberDL_RS_QCL_TypeD_Enum_n4  aper.Enumerated = 2
	SpatialRelations_maxNumberDL_RS_QCL_TypeD_Enum_n8  aper.Enumerated = 3
	SpatialRelations_maxNumberDL_RS_QCL_TypeD_Enum_n14 aper.Enumerated = 4
)

type SpatialRelations_maxNumberDL_RS_QCL_TypeD struct {
	Value aper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SpatialRelations_maxNumberDL_RS_QCL_TypeD) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SpatialRelations_maxNumberDL_RS_QCL_TypeD", err)
	}
	return nil
}

func (ie *SpatialRelations_maxNumberDL_RS_QCL_TypeD) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SpatialRelations_maxNumberDL_RS_QCL_TypeD", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
