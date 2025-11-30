package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	QCL_Info_qcl_Type_Enum_typeA aper.Enumerated = 0
	QCL_Info_qcl_Type_Enum_typeB aper.Enumerated = 1
	QCL_Info_qcl_Type_Enum_typeC aper.Enumerated = 2
	QCL_Info_qcl_Type_Enum_typeD aper.Enumerated = 3
)

type QCL_Info_qcl_Type struct {
	Value aper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *QCL_Info_qcl_Type) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode QCL_Info_qcl_Type", err)
	}
	return nil
}

func (ie *QCL_Info_qcl_Type) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode QCL_Info_qcl_Type", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
