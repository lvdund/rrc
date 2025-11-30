package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17_Enum_typeC            aper.Enumerated = 0
	DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17_Enum_typeD            aper.Enumerated = 1
	DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17_Enum_typeC_plus_typeD aper.Enumerated = 2
)

type DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17 struct {
	Value aper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17", err)
	}
	return nil
}

func (ie *DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode DL_PRS_QCL_Info_r17_ssb_r17_rs_Type_r17", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
